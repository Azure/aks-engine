$global:LogPath = "c:\k\on-restart.log"
$global:HNSModule = "c:\k\hns.psm1"

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log ($message) {
    $message | Timestamp | Tee-Object -FilePath $global:LogPath -Append
}

Write-Log "Entering on-restart.ps1"

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
Import-Module $global:HNSModule

Write-Log "Cleaning up persisted HNS policy lists"
Get-HnsPolicyList | Remove-HnsPolicyList

$hnsNetwork = Get-HnsNetwork | Where-Object Name -EQ azure
if ($hnsNetwork){
    Write-Log "Cleaning up HNS network 'azure'..."

    Write-Log "Cleaning up containers"
    docker ps -q | ForEach-Object {docker rm $_ -f}

    Write-Log "Removing old HNS network"
    Remove-HnsNetwork $hnsNetwork

    taskkill /IM azure-vnet.exe /f
    taskkill /IM azure-vnet-ipam.exe /f

    $cnijson = [io.path]::Combine("c:\k", "azure-vnet-ipam.json")
    if ((Test-Path $cnijson))
    {
        Remove-Item $cnijson
    }

    $cnilock = [io.path]::Combine("c:\k", "azure-vnet-ipam.json.lock")
    if ((Test-Path $cnilock))
    {
        Remove-Item $cnilock
    }

    $cnijson = [io.path]::Combine("c:\k", "azure-vnet.json")
    if ((Test-Path $cnijson))
    {
        Remove-Item $cnijson
    }

    $cnilock = [io.path]::Combine("c:\k", "azure-vnet.json.lock")
    if ((Test-Path $cnilock))
    {
        Remove-Item $cnilock
    }
}

$hnsNetwork = Get-HnsNetwork | Where-Object Name -EQ l2bridge
if ($hnsNetwork)
{
    Write-Log "cleaning up HNS network 'l2bridge'"

    Write-Log "cleaning up containers"
    docker ps -q | ForEach-Object {docker rm $_ -f}

    Write-Log "removing old HNS network"
    Remove-HnsNetwork $hnsNetwork

    Start-Sleep 10 

    Write-Log "Creating HNS network 'l2bridge'"
    # TODO: read values from cni config on disk and create network
}

#
# Start Services
#
Write-Log "Starting kubelet service"
Start-Service kubelet

Write-Log "Starting kubeproxy service"
Start-Service kubeproxy

Write-Log "Exiting on-restart.ps1"