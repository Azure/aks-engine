###################################################################################
# This script will automate the creation of a Forest, Domain, and Domain Controller
#
# This is only for automated e2e testing.  DO NOT use this for production.
# Jeremy Wood (JeremyWx)
# Version: 2019041400
###################################################################################

# Create a working directory and change to it
mkdir C:\gmsa
Set-Location -Path C:\gmsa

# Uncomment line below for troubleshooting
Start-Transcript -Path "C:\gmsa\Promote.txt"

# Download gMSA Setup Script
Invoke-WebRequest -UseBasicParsing https://raw.githubusercontent.com/JeremyWx/aks-engine/master/extensions/gmsa-dc/v1/Setup-gMSA.ps1 -OutFile Setup-gMSA.ps1

#Install NuGet
Install-PackageProvider -Name NuGet -Force

#Create Local Administrator
$admpassword = ( "K8s" + -join ((48..57) + (97..122) | Get-Random -Count 64 | ForEach-Object {[char]$_}) )
$admpassword_secure = ( $admpassword | ConvertTo-SecureString -AsPlainText -Force)
New-LocalUser -Name gmsa-admin -Password $admpassword_secure
Add-LocalGroupMember -Group "Administrators" -Member gmsa-admin
Start-Sleep -Seconds 5

# Make Setup-gMSA run on next boot
$Logon = "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon"
$RunOnce = "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce"
Set-ItemProperty $Logon "AutoAdminLogon" -Value "1" -type String
Set-ItemProperty $Logon "AutoLogonCount" -Value "2" -type DWord
Set-ItemProperty $Logon "DefaultUsername" -Value "k8sgmsa\gmsa-admin" -type String
Set-ItemProperty $Logon "DefaultPassword" -Value "$admpassword" -type String
New-ItemProperty $RunOnce "gmsa" -Value "C:\Windows\system32\WindowsPowerShell\v1.0\powershell.exe -command C:\gmsa\Setup-gMSA.ps1" -Type String

# Install and Enable SSH Server
$SSHService = Get-WindowsCapability -Online | Where-Object Name -like "OpenSSH.Server*"
Add-WindowsCapability -Online -Name $SSHService.Name
Start-Service -Name "sshd"
Set-Service -Name "sshd" -StartupType Automatic 
# In case of firewall
New-NetFirewallRule -Name "SSH Server" -DisplayName "SSH Server" -Description "Allow SSH Inbound" -Profile Any -Direction Inbound -Action Allow -Protocol TCP -Program Any -LocalAddress Any -RemoteAddress Any -LocalPort 22 -RemotePort Any 

# Import ServerManager and install the bits for ADDS
Import-Module ServerManager
Add-WindowsFeature -Name AD-Domain-Services,DNS -IncludeManagementTools
# Create new Forest and Domain with new DC and DNS
Install-ADDSForest -DomainName k8sgmsa.lan -SafeModeAdministratorPassword ("#55p@$$w0rd" | ConvertTo-SecureString -AsPlainText -Force) -InstallDNS -DomainMode 6 -DomainNetbiosName k8sgmsa -ForestMode 6 -Confirm:$false

