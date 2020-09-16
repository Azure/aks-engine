$ProgressPreference = "SilentlyContinue"

$lockedFiles = "kubelet.err.log", "kubelet.log", "kubeproxy.log", "kubeproxy.err.log", "azure-vnet-telemetry.log", "azure-vnet.log", "network-interfaces.json", "interfaces.json"

$timeStamp = get-date -format 'yyyyMMdd-hhmmss'
$zipName = "$env:computername-$($timeStamp)_logs.zip"

Write-Host "Collecting logs for various Kubernetes components"
$paths = @()
get-childitem c:\k\*.log* -Exclude $lockedFiles | Foreach-Object {
  $paths += $_
}

$lockedTemp = Join-Path ([System.IO.Path]::GetTempPath()) ([System.IO.Path]::GetRandomFileName())
New-Item -Type Directory $lockedTemp
$lockedFiles | Foreach-Object {
  Write-Host "Copying $_ to temp"
  $src = "c:\k\$_"
  if (Test-Path $src) {
    $tempfile = Copy-Item $src $lockedTemp -Passthru -ErrorAction Ignore
    if ($tempFile) {
      $paths += $tempFile
    }
  }
}

# Containerd log is outside the c:\k folder 
$containerd = "C:\ProgramData\containerd\root\panic.log"
if (Test-Path $containerd) {
  $tempfile = Copy-Item $containerd $lockedTemp -Passthru -ErrorAction Ignore
  if ($tempFile) {
    $paths += $tempFile
  }
}

Write-Host "Exporting ETW events to CSV files"
$scm = Get-WinEvent -FilterHashtable @{logname = 'System'; ProviderName = 'Service Control Manager' } | Where-Object { $_.Message -Like "*docker*" -or $_.Message -Like "*kub*" } | Select-Object -Property TimeCreated, Id, LevelDisplayName, Message
# 2004 = resource exhaustion, other 5 events related to reboots
$reboots = Get-WinEvent -ErrorAction Ignore -FilterHashtable @{logname = 'System'; id = 1074, 1076, 2004, 6005, 6006, 6008 } | Select-Object -Property TimeCreated, Id, LevelDisplayName, Message
$crashes = Get-WinEvent -ErrorAction Ignore -FilterHashtable @{logname = 'Application'; ProviderName = 'Windows Error Reporting' } | Select-Object -Property TimeCreated, Id, LevelDisplayName, Message
$scm + $reboots + $crashes | Sort-Object TimeCreated | Export-CSV -Path "$ENV:TEMP\\$($timeStamp)_services.csv"
$paths += "$ENV:TEMP\\$($timeStamp)_services.csv"
Get-WinEvent -LogName Microsoft-Windows-Hyper-V-Compute-Operational | Select-Object -Property TimeCreated, Id, LevelDisplayName, Message | Sort-Object TimeCreated | Export-Csv -Path "$ENV:TEMP\\$($timeStamp)_hyper-v-compute-operational.csv"
$paths += "$ENV:TEMP\\$($timeStamp)_hyper-v-compute-operational.csv"
get-eventlog -LogName Application -Source Docker | Select-Object Index, TimeGenerated, EntryType, Message | Sort-Object Index | Export-CSV -Path "$ENV:TEMP\\$($timeStamp)_docker.csv"
$paths += "$ENV:TEMP\\$($timeStamp)_docker.csv"
Get-CimInstance win32_pagefileusage | Format-List * | Out-File -Append "$ENV:TEMP\\$($timeStamp)_pagefile.txt"
Get-CimInstance win32_computersystem | Format-List AutomaticManagedPagefile | Out-File -Append "$ENV:TEMP\\$($timeStamp)_pagefile.txt"
$paths += "$ENV:TEMP\\$($timeStamp)_pagefile.txt"
mkdir 'c:\k\debug' -ErrorAction Ignore | Out-Null

Write-Host "Collecting networking related logs"
if (-not (Test-Path 'c:\k\debug\collectlogs.ps1')) {
  Invoke-WebRequest -UseBasicParsing https://raw.githubusercontent.com/Microsoft/SDN/master/Kubernetes/windows/debug/collectlogs.ps1 -OutFile 'c:\k\debug\collectlogs.ps1'
}
& 'c:\k\debug\collectlogs.ps1' | write-Host
$netLogs = Get-ChildItem (Get-ChildItem -Path c:\k\debug -Directory | Sort-Object LastWriteTime -Descending | Select-Object -First 1).FullName | Select-Object -ExpandProperty FullName
$paths += $netLogs
$paths += "c:\AzureData\CustomDataSetupScript.log"

Write-Host "Collecting containerd hyperv logs"
if ((Test-Path "$Env:ProgramFiles\containerd\diag.ps1") -And (Test-Path "$Env:ProgramFiles\containerd\ContainerPlatform.wprp")) {
  $tempHyperv = Join-Path ([System.IO.Path]::GetTempPath()) ([System.IO.Path]::GetRandomFileName())
  New-Item -Type Directory $tempHyperv
  $persistedLogs = "c:\logs"
  # there will either be an error collecting "bootlogs" or "trace profiles" as only one will be active at time. This will be fixed in future release of the script 
  & $Env:ProgramFiles\containerd\diag.ps1 -Snap -ProfilePath "$Env:ProgramFiles\containerd\ContainerPlatform.wprp!ContainerPlatformPersistent" -TraceDirPath "$tempHyperv" -TempPath $persistedLogs
  $hypervlogs = (Get-ChildItem -Path $tempHyperv | Sort-Object LastWriteTime -Descending | Select-Object -First 1).FullName
  $paths += $hypervlogs
}
else {
  Write-Host "Containerd hyperv logs not avalaible"
}

Write-Host "Compressing all logs to $zipName"
$paths | Format-Table FullName, Length -AutoSize
Compress-Archive -LiteralPath $paths -DestinationPath $zipName
Get-ChildItem $zipName # this puts a FileInfo on the pipeline so that another script can get it on the pipeline
