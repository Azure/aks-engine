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
$global:ContainerRuntime = $Global:ClusterConfiguration.Cri.Name
$UseContainerD = ($global:ContainerRuntime -eq "containerd")
$IsDualStackEnabled = $Global:ClusterConfiguration.Kubernetes.Kubeproxy.FeatureGates -contains "IPv6DualStack=true"

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log ($message) {
    $message | Timestamp | Tee-Object -FilePath $global:LogPath -Append
}

Write-Log "Entering windowsnodereset.ps1"

Import-Module $global:HNSModule

Unregister-HNSRemediatorScriptTask

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

if ($global:EnableHostsConfigAgent) {
    Write-Log "Stopping hosts-config-agent service"
    Stop-Service hosts-config-agent
}

function Register-HNSRemediatorScriptTask {
    # Hardcoding RepetitionInterval to 1 Minute
    # Making it variable would need a new parameter to be added under windowsProfile
    Write-Log "Creating a scheduled task to run hnsremediator.ps1"

    $action = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-File `"c:\k\hnsremediator.ps1`""
    $principal = New-ScheduledTaskPrincipal -UserId SYSTEM -LogonType ServiceAccount -RunLevel Highest
    $trigger = New-JobTrigger -Once -At (Get-Date).Date -RepeatIndefinitely -RepetitionInterval (New-TimeSpan -Minutes 1)
    $definition = New-ScheduledTask -Action $action -Principal $principal -Trigger $trigger -Description "hns-remediator-task"
    Register-ScheduledTask -TaskName "hns-remediator-task" -InputObject $definition
}

function Unregister-HNSRemediatorScriptTask {
    if (Get-ScheduledTask -TaskName "hns-remediator-task" -ErrorAction Ignore) {
        Write-Log "Deleting the scheduled task hns-remediator-task"
        Unregister-ScheduledTask -TaskName "hns-remediator-task" -Confirm:$false
    }

    $hnsPIDFile="C:\k\hns.pid"
    if (Test-Path $hnsPIDFile) {
        # Remove this file since PID of HNS service may have been changed after node crashes or is rebooted
        # It should not always fail since hns-remediator-task is unregistered.
        # We set the max retry count to 20 to avoid dead loop for unknown issues.
        $maxRetries=20
        $retryCount=0
        while ($retryCount -lt $maxRetries) {
            Write-Log "Deleting $hnsPIDFile"
            Remove-Item -Path $hnsPIDFile -Force -Confirm:$false -ErrorAction Ignore

            # The file may not be deleted successfully because hnsremediator.ps1 is still writing the logs
            if (Test-Path $hnsPIDFile) {
                # Do not log the failure to reduce log
                Start-Sleep -Milliseconds 500
                $retryCount=$retryCount+1
            } else {
                Write-Log "$hnsPIDFile is deleted"
                break
            }
        }
    }
}

# Due to a bug in hns there is a race where it picks up the incorrect IPv6 address from the node in some cases.
# Hns service has to be restarted after the node internal IPv6 address is available when dual-stack is enabled.
# TODO Remove this once the bug is fixed in hns.
function Restart-HnsService {
    do {
        Start-Sleep -Seconds 1
        $nodeInternalIPv6Address = (Get-NetIPAddress | Where-Object {$_.PrefixOrigin -eq "Dhcp" -and $_.AddressFamily -eq "IPv6"}).IPAddress 
    } while ($nodeInternalIPv6Address -eq $null)
    Write-Log "Got node internal IPv6 address: $nodeInternalIPv6Address"
    
    $hnsManagementIPv6Address = (Get-HnsNetwork | Where-Object {$_.IPv6 -eq $true}).ManagementIPv6
    Write-Log "Got hns ManagementIPv6: $hnsManagementIPv6Address"

    if ($hnsManagementIPv6Address -ne $nodeInternalIPv6Address) {
        Restart-Service hns
        Write-Log "Restarted hns service"

        $hnsManagementIPv6Address = (Get-HnsNetwork | Where-Object {$_.IPv6 -eq $true}).ManagementIPv6
        Write-Log "Got hns ManagementIPv6: $hnsManagementIPv6Address after restart"
    }
    else {
        Write-Log "Hns network has correct IPv6 address, not restarting"
    }
}

if ($IsDualStackEnabled) {
    Restart-HnsService
}

#
# Perform cleanup
#

& "c:\k\cleanupnetwork.ps1"

#
# Create required networks
#

# If using kubenet create the HNS network here.
# (The kubelet creates the HNS network when using azure-cni + azure cloud provider)
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

Write-Log "Do not start kubeproxy service since kubelet will restart kubeproxy"

Register-HNSRemediatorScriptTask

Write-Log "Exiting windowsnodereset.ps1"
