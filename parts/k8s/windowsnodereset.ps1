<#
.DESCRIPTION
    This script is intended to be run each time a windows nodes is restarted and performs
    cleanup actions to help ensure the node comes up cleanly.
#>

$global:LogPath = "c:\k\windowsnodereset.log"
$global:HNSModule = "c:\k\hns.psm1"

# Note: the following templated values are expanded kuberneteswindowsfunctions.ps1/Register-NodeResetScriptTask() not during template generation!
$global:MasterSubnet = "{{MasterSubnet}}"
$global:NetworkMode = "{{NetworkMode}}"
$global:NetworkPlugin = "{{NetworkPlugin}}"

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log ($message) {
    $message | Timestamp | Tee-Object -FilePath $global:LogPath -Append
}

Write-Log "Entering windowsnodecleanup.ps1"

Import-Module $global:HNSModule

#
# Stop services
#
Write-Log "Stopping kubeproxy service"
Stop-Service kubeproxy

Write-Log "Stopping kubelet service"
Stop-Service kubelet

#
# Perform cleanup
#

$hnsNetwork = Get-HnsNetwork | Where-Object Name -EQ azure
if ($hnsNetwork) {
    Write-Log "Cleaning up HNS network 'azure'..."

    Write-Log "Cleaning up containers"
    docker ps -q | ForEach-Object { docker rm $_ -f }

    Write-Log "Removing old HNS network"
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

#
# Create required networks
#

Write-Log "Cleaning up persisted HNS policy lists"
# Workaround for https://github.com/kubernetes/kubernetes/pull/68923 in < 1.14,
# and https://github.com/kubernetes/kubernetes/pull/78612 for <= 1.15
Get-HnsPolicyList | Remove-HnsPolicyList

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
Write-Log "Starting kubelet service"
Start-Service kubelet

Write-Log "Starting kubeproxy service"
Start-Service kubeproxy

Write-Log "Exiting windowsnodecleanup.ps1"
