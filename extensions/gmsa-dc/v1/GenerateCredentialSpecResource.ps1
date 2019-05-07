<#
.Synopsis
 Renders a GMSA kubernetes resource manifest.
#>
Param(
    [Parameter(Position = 0, Mandatory = $true)] [String] $AccountName,
    [Parameter(Position = 1, Mandatory = $true)] [String] $ResourceName,
    [Parameter(Position = 2, Mandatory = $false)] [String] $ManifestFile,
    [Parameter(Mandatory=$false)] $Domain,
    [Parameter(Mandatory=$false)] [string[]] $AdditionalAccounts = @()
)
Start-Transcript -Path "C:\gmsa\CredSpec.txt"
# exit on error
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
$PSDefaultParameterValues['*:ErrorAction'] = 'Stop'

# generate the name of the output file if not specified
if (-not $ManifestFile -or $ManifestFile.Length -eq 0) {
    $ManifestFile = "gmsa-cred-spec-$ResourceName.yml"
}
# check the out file doesn't exist
if ([System.IO.File]::Exists($ManifestFile)) {
    throw "Output file $ManifestFile already exists, refusing to overwrite it"
}

# install the dependencies we need
if (-not (Get-WindowsFeature rsat-ad-powershell).Installed) {
    Add-WindowsFeature rsat-ad-powershell
}
if (-not (Get-Command ConvertTo-Yaml -errorAction SilentlyContinue)) {
    Install-Module powershell-yaml -Force
}

# download the canonical helper script
Invoke-WebRequest "https://raw.githubusercontent.com/Microsoft/Virtualization-Documentation/live/windows-server-container-tools/ServiceAccounts/CredentialSpec.psm1" -UseBasicParsing -OutFile $env:TEMP\cred.psm1
Import-Module $env:temp\cred.psm1

# generate a unique docker cred spec name
$dockerCredSpecName = "tmp-k8s-cred-spec" + -join ((48..57) + (97..122) | Get-Random -Count 64 | ForEach-Object {[char]$_})

# have the upstream function perform its magic
if (-not $Domain) {
    $Domain = Get-ADDomain
}
New-CredentialSpec -Name $dockerCredSpecName -AccountName $AccountName -Domain $Domain.DnsRoot -AdditionalAccounts $AdditionalAccounts > C:\gmsa\create-account.txt

# parse the JSON file thus generated
$dockerCredSpecPath = (Get-CredentialSpec | Where-Object {$_.Name -like "$dockerCredSpecName*"}).Path
$credSpecContents = Get-Content $dockerCredSpecPath | ConvertFrom-Json
# and clean it up
Remove-Item $dockerCredSpecPath

# generate the k8s resource
$resource = [ordered]@{
    "apiVersion" = "windows.k8s.io/v1alpha1";
    "kind" = 'GMSACredentialSpec';
    "metadata" = @{
        "name" = $ResourceName
    };
    "credspec" = $credSpecContents
}

ConvertTo-Yaml $resource | Set-Content $ManifestFile

Write-Output "K8S manifest rendered at $ManifestFile"
