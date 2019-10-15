$global:LogPath = "c:\k\windowsnodecleanup.log"
$global:HNSModule = "c:\k\hns.psm1"

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

Write-Log "Cleaning up persisted HNS policy lists"
Get-HnsPolicyList | Remove-HnsPolicyList

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

$hnsNetwork = Get-HnsNetwork | Where-Object Name -EQ l2bridge
if ($hnsNetwork) {
    Write-Log "cleaning up HNS network 'l2bridge'"

    Write-Log "cleaning up containers"
    docker ps -q | ForEach-Object { docker rm $_ -f }

    Write-Log "removing old HNS network"
    Remove-HnsNetwork $hnsNetwork

    Start-Sleep 10 
}

# TODO: if network plugin is kubenet create l2bridge network

#
# Start Services
#
Write-Log "Starting kubelet service"
Start-Service kubelet

Write-Log "Starting kubeproxy service"
Start-Service kubeproxy

Write-Log "Exiting windowsnodecleanup.ps1"