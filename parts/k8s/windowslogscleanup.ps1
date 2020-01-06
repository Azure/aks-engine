<#
.DESCRIPTION
    This script cleans old rotated logs for various kubernetes components.
#>

$global:LogPath = "c:\k\windowslogscleanup.log"

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log ($message) {
    $message | Timestamp | Tee-Object -FilePath $global:LogPath -Append
}

Write-Log "Entering windowslogscleanup.ps1"

$logFilePrefixes = @("kubelet", "kubelet.err", "kubeproxy", "kubeproxy.err")

foreach ($logFilePrefix in $logFilePrefixes) {
    $oldLogs = [IO.Directory]::GetFiles("c:\k", "$($logFilePrefix)-*.log")
    $oldLogs = $oldLogs | Sort-Object | Select-Object -SkipLast 5
    foreach ($oldLog in $oldLogs) {
        Write-Log "Removing $oldLog"
        Remove-Item $oldLog
    }
}