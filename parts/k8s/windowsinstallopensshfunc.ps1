function
Install-OpenSSH {
    Param(
        [Parameter(Mandatory = $true)][string[]] 
        $SSHKeys
    )

    $adminpath = "c:\ProgramData\ssh"
    $adminfile = "administrators_authorized_keys"

    $sshdService = Get-Service | ? Name -like 'sshd'
    if ($sshdService.Count -eq 0)
    {
        Write-Log "Installing OpenSSH"
        $isAvailable = Get-WindowsCapability -Online | ? Name -like 'OpenSSH*'

        if (!$isAvailable) {
            throw "OpenSSH is not available on this machine"
        }

        Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
    }
    else
    {
        Write-Log "OpenSSH Server service detected - skipping online install..."
    }

    Start-Service sshd

    if (!(Test-Path "$adminpath")) {
        Write-Log "Created new file and text content added"
        New-Item -path $adminpath -name $adminfile -type "file" -value ""
    }

    Write-Log "$adminpath found."
    Write-Log "Adding keys to: $adminpath\$adminfile ..."
    $SSHKeys | foreach-object {
        Add-Content $adminpath\$adminfile $_
    }

    Write-Log "Setting required permissions..."
    icacls $adminpath\$adminfile /remove "NT AUTHORITY\Authenticated Users"
    icacls $adminpath\$adminfile /inheritance:r
    icacls $adminpath\$adminfile /grant SYSTEM:`(F`)
    icacls $adminpath\$adminfile /grant BUILTIN\Administrators:`(F`)

    Write-Log "Restarting sshd service..."
    Restart-Service sshd
    # OPTIONAL but recommended:
    Set-Service -Name sshd -StartupType 'Automatic'

    # Confirm the Firewall rule is configured. It should be created automatically by setup. 
    $firewall = Get-NetFirewallRule -Name *ssh*

    if (!$firewall) {
        throw "OpenSSH is firewall is not configured properly"
    }
    Write-Log "OpenSSH installed and configured successfully"
}
