##################################################################################
# This script will automate the creation of a Group Managed Service Account,
# install the account locally, and produce the JSON file with gMSA info.
#
# This is only for automated e2e testing.  DO NOT use this for production.
# Jeremy Wood (JeremyWx)
# Version: 2019041400
##################################################################################

# Change to working directory
Set-Location -Path C:\gmsa

# Uncomment line below for troubleshooting
Start-Transcript -Path "C:\gmsa\Setup-gmsa.txt" -Append
Start-Sleep -Seconds 60
Import-Module ActiveDirectory

if (Get-AdGroupMember -Identity "Enterprise Admins" | Select-String -Pattern "gmsa-admin" -Quiet) {

    $KdsRootKey = Get-KdsRootKey
    if ($null -eq $KdsRootKey.KeyId) {
        Add-KdsRootKey –EffectiveTime ((get-date).addhours(-10))
        Start-Sleep -Seconds 15
    }
    
    # Directory for credspecs if not already created
    mkdir -Path C:\ProgramData\docker\credentialspecs

    # Get GenerateCredSpecResource PowerShell Module
    #Invoke-WebRequest -UseBasicParsing https://raw.githubusercontent.com/kubernetes-sigs/windows-gmsa/master/scripts/GenerateCredentialSpecResource.ps1 -OutFile GenerateCredentialSpecResource.ps1
    Invoke-WebRequest -UseBasicParsing https://raw.githubusercontent.com/JeremyWx/aks-engine/master/extensions/gmsa-dc/v1/GenerateCredentialSpecResource.ps1 -OutFile GenerateCredentialSpecResource.ps1

    # Import AD Module and Setup gMSA
    New-ADServiceAccount -Name gmsa-e2e -DNSHostName gmsa-e2e.k8sgmsa.lan -PrincipalsAllowedToRetrieveManagedPassword "Domain Controllers" -ServicePrincipalnames http/gmsa-e2e.k8sgmsa.lan

    # Run GenerateCredSpecResource to provide the Yaml CredSpec
    C:\gmsa\GenerateCredentialSpecResource.ps1 gmsa-e2e gmsa-e2e

} else {
    Add-ADGroupMember -Identity "Enterprise Admins" -Members gmsa-admin
    Add-ADGroupMember -Identity "Domain Admins" -Members gmsa-admin
    $RunOnce = "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce"
    Set-ItemProperty $RunOnce "gmsa" -Value "C:\Windows\system32\WindowsPowerShell\v1.0\powershell.exe -command C:\gmsa\Setup-gMSA.ps1" -Type String
    Restart-Computer -Force
}

