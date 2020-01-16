<#
    .SYNOPSIS
        Produces a release notes file for a Windows VHD

    .DESCRIPTION
        Produces a release notes file for a Windows VHD
#>

$ErrorActionPreference = "Stop"

$releaseNotesFilePath = "c:\release-notes.txt"

function Log($Message) {
    # Write-Output $Message
    $Message | Tee-Object -FilePath $releaseNotesFilePath -Append
}

Log "Build Number: $env:BUILD_NUMBER"
Log "Build Id:     $env:BUILD_ID"
Log "Build Repo:   $env:BUILD_REPO"
Log "Build Branch: $env:BUILD_BRANCH"
Log "Commit:       $env:BUILD_COMMIT"
Log ""

$vhdId = Get-Content 'c:\vhd-id.txt'
LOG ("VHD ID:      $vhdId")
LOG ""

Log "System Info"
$systemInfo = Get-ItemProperty -Path 'HKLM:SOFTWARE\Microsoft\Windows NT\CurrentVersion'
Log ("`t{0,-14} : {1}" -f "OS Name", $systemInfo.ProductName)
LOG ("`t{0,-14} : {1}" -f "OS Version", "$($systemInfo.CurrentBuildNumber).$($systemInfo.UBR)")
LOG ("`t{0,-14} : {1}" -f "OS InstallType", $systemInfo.InstallationType)
Log ""

$allowedSecurityProtocols = [System.Net.ServicePointManager]::SecurityProtocol
Log "Allowed security protocols: $allowedSecurityProtocols"
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
Log ""

Log "Cached Files:"
$displayObjects = @()
foreach ($file in [IO.Directory]::GetFiles('c:\akse-cache', '*', [IO.SearchOption]::AllDirectories))
{
    $attributes = Get-Item $file
    $hash = Get-FileHash $file -Algorithm SHA256
    $displayObjects += New-Object psobject -property @{
        File = $file;
        SizeBytes = $attributes.Length;
        Sha256 = $hash.Hash
    }
}

Log ($displayObjects | Format-Table -Property File, Sha256, SizeBytes | Out-String -Width 4096)

# Ensure proper encoding is set for release notes file
[IO.File]::ReadAllText($releaseNotesFilePath) | Out-File -Encoding utf8 $releaseNotesFilePath
