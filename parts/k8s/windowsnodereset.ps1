<#
.DESCRIPTION
    This script is intended to be run each time a windows nodes is restarted and performs
    cleanup actions to help ensure the node comes up cleanly.
#>

$global:LogPath = "c:\k\windowsnodereset.log"
$global:HNSModule = "c:\k\hns.psm1"

$Global:ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string)

$global:CsiProxyEnabled = [System.Convert]::ToBoolean($Global:ClusterConfiguration.Csi.EnableProxy)
$global:MasterSubnet = $Global:ClusterConfiguration.Kubernetes.ControlPlane.MasterSubnet
$global:NetworkMode = "L2Bridge"
$global:NetworkPlugin = $Global:ClusterConfiguration.Cni.Name

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log ($message) {
    $message | Timestamp | Tee-Object -FilePath $global:LogPath -Append
}

Write-Log "Entering windowsnodereset.ps1"

Import-Module $global:HNSModule

#
# Stop services
#
Write-Log "Stopping kubeproxy service"
Stop-Service kubeproxy

Write-Log "Stopping kubelet service"
Stop-Service kubelet

if ($global:CsiProxyEnabled) {
    Write-Log "Stopping csi-proxy service"
    Stop-Service csi-proxy
}

#
# Perform cleanup
#

$hnsNetwork = Get-HnsNetwork | Where-Object Name -EQ azure
if ($hnsNetwork) {
    Write-Log "Cleaning up containers"
    docker ps -q | ForEach-Object { docker rm $_ -f }

    Write-Log "Removing old HNS network 'azure'"
    Remove-HnsNetwork $hnsNetwork

    taskkill /IM azure-vnet.exe /f
    taskkill /IM azure-vnet-ipam.exe /f

    $filesToRemove = @(
        "c:\k\azure-vnet.json",
        "c:\k\azure-vnet.json.lock",
        "c:\k\azure-vnet-ipam.json",
        "c:\k\azure-vnet-ipam.json.lock"
    )

    foreach ($file in $filesToRemove) {
        if (Test-Path $file) {
            Write-Log "Deleting stale file at $file"
            Remove-Item $file
        }
    }
}

Write-Log "Cleaning up persisted HNS policy lists"
# Workaround for https://github.com/kubernetes/kubernetes/pull/68923 in < 1.14,
# and https://github.com/kubernetes/kubernetes/pull/78612 for <= 1.15
Get-HnsPolicyList | Remove-HnsPolicyList

#
# Create required networks
#

# If using kubenet create the HSN network here.
# (The kubelet creates the HSN network when using azure-cni + azure cloud provider)
if ($global:NetworkPlugin -eq 'kubenet') {
    Write-Log "Creating new hns network: $($global:NetworkMode.ToLower())"
    $podCIDR = Get-PodCIDR
    $masterSubnetGW = Get-DefaultGateway $global:MasterSubnet
    New-HNSNetwork -Type $global:NetworkMode -AddressPrefix $podCIDR -Gateway $masterSubnetGW -Name $global:NetworkMode.ToLower() -Verbose
    Start-sleep 10
}

#
# Start Services
#

if ($global:CsiProxyEnabled) {
    Write-Log "Starting csi-proxy service"
    Start-Service csi-proxy
}

Write-Log "Starting kubelet service"
Start-Service kubelet

Write-Log "Starting kubeproxy service"
Start-Service kubeproxy

Write-Log "Exiting windowsnodereset.ps1"
