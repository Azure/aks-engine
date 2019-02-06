function
Install-OpenSSH {
    Param(
        [Parameter(Mandatory = $true)][string] 
        $SSHKey
    )

    $adminpath = "c:\ProgramData\ssh"
    $adminfile = "administrators_authorized_keys"

    Write-Host "Installing OpenSSH"
    $isAvailable = Get-WindowsCapability -Online | ? Name -like 'OpenSSH*'

    if (!$isAvailable) {
        Write-Error "OpenSSH is not avaliable on this machine"
        exit 1
    }

    Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0

    Start-Service sshd

    if (!(Test-Path "$adminpath")) {
        Write-Host "Created new file and text content added"
        New-Item -path $adminpath -name $adminfile -type "file" -value ""
    }

    Write-Host "$adminpath found."
    Write-Host "Adding key to: $adminpath\$adminfile ..."
    Add-Content $adminpath\$adminfile $SSHKey

    Write-Host "Setting required permissions..."
    icacls $adminpath\$adminfile /remove "NT AUTHORITY\Authenticated Users"
    icacls $adminpath\$adminfile /inheritance:r

    Write-Host "Restarting sshd service..."
    Restart-Service sshd
    # OPTIONAL but recommended:
    Set-Service -Name sshd -StartupType 'Automatic'

    # Confirm the Firewall rule is configured. It should be created automatically by setup. 
    $firewall = Get-NetFirewallRule -Name *ssh*

    if (!$firewall) {
        Write-Error "OpenSSH is firewall is not configured properly"
        exit 1
    }
    Write-Host "OpenSSH installed and configured successfully"
}