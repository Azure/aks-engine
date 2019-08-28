<#
    .SYNOPSIS
        Produces a release notes file for a Windows VHD

    .DESCRIPTION
        Produces a release notes file for a Windows VHD
#>

$ErrorActionPreference = "Stop"

function Log($Message) {
    # Write-Output $Message
    $Message | Tee-Object -FilePath "c:\release-notes.txt" -Append
}

Log "System Info"
$systemInfo = Get-ItemProperty -Path 'HKLM:SOFTWARE\Microsoft\Windows NT\CurrentVersion'
Log ("`t{0,-14} : {1}" -f "OS Name", $systemInfo.ProductName)
LOG ("`t{0,-14} : {1}" -f "OS Version", "$($systemInfo.CurrentBuildNumber).$($systemInfo.UBR)")
LOG ("`t{0,-14} : {1}" -f "OS InstallType", $systemInfo.InstallationType)
Log ""
Log "Installed Features"
if ($systemInfo.InstallationType -ne 'client') {
    Log (Get-WindowsFeature | Where-Object Installed)
}
else {
    LOG "`t<Cannot enumerate installed features on client skus>"
}
Log ""

Log "Installed Packages"
$packages = Get-WindowsCapability -Online | Where-Object { $_.State -eq 'Installed' }
foreach ($package in $packages) {
    Log ("`t{0}" -f $package.Name)
}
Log ""

Log "Installed QFEs"
$qfes = Get-HotFix
foreach ($qfe in $qfes) {
    $link = "http://support.microsoft.com/?kbid={0}" -f ($qfe.HotFixID.Replace("KB", ""))
    LOG ("`t{0,-9} : {1, -15} : {2}" -f $qfe.HotFixID, $Qfe.Description, $link)
}
Log ""

Log "Installed Updates"
$updateSession = New-Object -ComObject Microsoft.Update.Session
$updateSearcher = $UpdateSession.CreateUpdateSearcher()
$updates = $updateSearcher.Search("IsInstalled=1").Updates
foreach ($update in $updates) {
    LOG ("`t{0}" -f $update.Title)
}
LOG ""

LOG "Windows Update Registry Settings"
LOG "`thttps://docs.microsoft.com/en-us/windows/deployment/update/waas-wu-settings"

$wuRegistryKeys = @(
    "HKLM:SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate",
    "HKLM:SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU"
)

foreach ($key in $wuRegistryKeys) {
    Log ("`t{0}" -f $key)
    Get-Item -Path $key |
    Select-Object -ExpandProperty property |
    ForEach-Object {
        Log ("`t`t{0} : {1}" -f $_, (Get-ItemProperty -Path $key -Name $_).$_)
    }
}
Log ""

if (Test-Path 'C:\Program Files\Docker\') {
    Log "Docker Info"
    $dockerVersion = (docker --version) | Out-String
    Log ("Version: {0}" -f $dockerVersion)
    Log "Images:"
    LOG (docker images --format='{{json .}}' | ConvertFrom-Json | Format-Table Repository, Tag, ID)
}