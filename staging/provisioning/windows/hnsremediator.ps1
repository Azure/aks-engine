<#
.DESCRIPTION
    HNS service may crash and HNS policies will be purged after it is restarted.
    We use this script to restart kubeproxy to recover the node from the hns crash.
    Start sequence:
    1. windowsnodereset.ps1 deletes hns-remediator-task if hns-remediator-task exists
    2. windowsnodereset.ps1 deletes "C:\k\hns.pid" if "C:\k\hns.pid" exists
    3. windowsnodereset.ps1 resets all services, hns, csi, kubeproxy, kubelet, etc.
    4. windowsnodereset.ps1 creates hns-remediator-task with $Global:ClusterConfiguration.Services.HNSRemediator.IntervalInMinutes
       in c:\k\kubeclusterconfig.json when the value is not 0
    NOTES:
    1. We cannot run hns-remediator-task with an interval less than 1 minute since the RepetitionInterval parameter value in New-JobTrigger must be greater than 1 minute.
    2. When the node crashes or is rebooted, hns-remediator-task may restart kubeproxy before windowsnodereset.ps1 is executed.
       It should have no impact since windowsnodereset.ps1 always deletes hns-remediator-task and then deletes "C:\k\hns.pid" before stopping kubeproxy
#>

$LogPath = "c:\k\hnsremediator.log"
$hnsPIDFilePath="C:\k\hns.pid"
$isInitialized=$False

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log ($message) {
    $message | Timestamp | Tee-Object -FilePath $LogPath -Append
}

if (Test-Path -Path $hnsPIDFilePath) {
    $isInitialized=$True
}

$id = Get-WmiObject -Class Win32_Service -Filter "name='hns'" | Select-Object -ExpandProperty ProcessId
if (!$isInitialized) {
    Write-Log "Initializing with creating $hnsPIDFilePath. PID of HNS service is $id"
    echo $id > $hnsPIDFilePath
    $isInitialized=$True
}

$lastId=Get-Content $hnsPIDFilePath
if ($lastId -ne $id) {
    Write-Log "The PID of HNS service was changed from $lastId to $id"
    echo $id > $hnsPIDFilePath

    Write-Log "Restarting kubeproxy service"
    Restart-Service kubeproxy
    Write-Log "Restarted kubeproxy service"

    $calicoService = Get-Service -Name CalicoFelix -ErrorAction Ignore
    if ($calicoService) {
        Write-Log "Restarting Calico services"
        # CalicoFelix depends on CalicoNode
        # https://github.com/projectcalico/calico/blob/master/node/windows-packaging/CalicoWindows/start-calico.ps1#L20
        # https://github.com/projectcalico/calico/blob/35b0c499dc0b01d228cf70ba942afe4eb1b6a961/node/windows-packaging/CalicoWindows/felix/felix-service.ps1#L21
        Restart-Service CalicoNode -ErrorAction Ignore
        Restart-Service CalicoFelix -ErrorAction Ignore
        Write-Log "Restarted Calico services"
    }
}