<#
    .SYNOPSIS
        Used to produce Windows AKS images.

    .DESCRIPTION
        This script is used by packer to produce Windows AKS images.
#>

param()

$ErrorActionPreference = "Stop"

filter Timestamp {"$(Get-Date -Format o): $_"}

function Write-Log($Message)
{
    $msg = $message | Timestamp
    Write-Output $msg
}

function Disable-WindowsUpdates
{
    # See https://docs.microsoft.com/en-us/windows/deployment/update/waas-wu-settings 
    # for additional information on WU related registry settings

    Write-Log "Disabling automatic windows upates"
    $WindowsUpdatePath = "HKLM:SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate"
    $AutoUpdatePath = "HKLM:SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU"

    if (Test-Path -Path $WindowsUpdatePath)
    {
        Remove-Item -Path $WindowsUpdatePath -Recurse
    }

    New-Item -Path $WindowsUpdatePath | Out-Null
    New-Item -Path $AutoUpdatePath | Out-Null
    Set-ItemProperty -Path $AutoUpdatePath -Name NoAutoUpdate -Value 1 | Out-Null
}

function Get-ContainerImages
{
    $imagesToPull = @(
        "mcr.microsoft.com/windows/servercore:ltsc2019",
        "mcr.microsoft.com/windows/nanoserver:1809",
        "mcr.microsoft.com/k8s/core/pause:1.2.0")

    foreach ($image in $imagesToPull) {
        docker pull $image
    }
}

function Get-FilesToCacheOnVHD
{
    Write-Log "Caching misc files on VHD"

    $map = @{
        "c:\akse-cache\" = @(
            "https://github.com/Microsoft/SDN/raw/master/Kubernetes/flannel/l2bridge/cni/win-bridge.exe",
            "https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/hns.psm1"
        );
        "c:\akse-cache\win-k8s\" = @(
            "https://acs-mirror.azureedge.net/wink8s/azs-v1.14.6-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/azs-v1.14.7-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/azs-v1.15.3-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/azs-v1.15.4-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/azs-v1.16.0-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/azs-v1.16.1-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/v1.14.6-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/v1.14.7-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/v1.15.3-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/v1.15.4-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/v1.16.0-1int.zip",
            "https://acs-mirror.azureedge.net/wink8s/v1.16.1-1int.zip"
        );
        "c:\akse-cache\win-vnet-cni\" = @(
            "https://acs-mirror.azureedge.net/cni/azure-vnet-cni-windows-amd64-v1.0.27.zip"
        )
    }

    foreach ($dir in $map.Keys)
    {
        New-Item -ItemType Directory $dir -Force | Out-Null

        foreach ($URL in $map[$dir])
        {
            $fileName = [IO.Path]::GetFileName($URL)
            $dest = [IO.Path]::Combine($dir, $fileName)

            Write-Log "Downloading $URL to $dest"
            Invoke-WebRequest -UseBasicParsing -Uri $URL -OutFile $dest
        }
    }
}

function Install-Docker
{
    $defaultDockerVersion = "19.03.2"

    Write-Log "Attempting to install Docker version $defaultDockerVersion"
    Install-PackageProvider -Name DockerMsftProvider -Force -ForceBootstrap | Out-null
    $package = Find-Package -Name Docker -ProviderName DockerMsftProvider -RequiredVersion $defaultDockerVersion
    Write-Log "Installing Docker version $($package.Version)"
    $package | Install-Package -Force | Out-Null
    Start-Service docker
}

function Install-OpenSSH
{
    Write-Log "Installing OpenSSH Server"
    Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
}

function Install-WindowsPatches
{
    # Windows Server 2019 update history can be found at https://support.microsoft.com/en-us/help/4464619
    # then you can get download links by searching for specific KBs at http://www.catalog.update.microsoft.com/home.aspx

    $patchUrls = @()

    foreach ($patchUrl in $patchUrls)
    {
        $pathOnly = $patchUrl.Split("?")[0]
        $fileName = Split-Path $pathOnly -Leaf
        $fileExtension = [IO.Path]::GetExtension($fileName)
        $fullPath = [IO.Path]::Combine($env:TEMP, $fileName)

        switch ($fileExtension)
        {
            ".msu"
            {
                Write-Log "Downloading windows patch from $pathOnly to $fullPath"
                Invoke-WebRequest -UseBasicParsing $patchUrl -OutFile $fullPath
                Write-Log "Starting install of $fileName"
                $proc = Start-Process -Passthru -FilePath wusa.exe -ArgumentList "$fullPath /quiet /norestart"
                Wait-Process -InputObject $proc
                switch ($proc.ExitCode)
                {
                    0
                    {
                        Write-Log "Finished install of $fileName"
                    }
                    3010
                    {
                        WRite-Log "Finished install of $fileName. Reboot required"
                    }
                    default
                    {
                        Write-Log "Error during install of $fileName. ExitCode: $($proc.ExitCode)"
                        exit 1
                    }
                }
            }
            default
            {
                Write-Log "Installing patches with extension $fileExtension is not currently supported."
                exit 1
            }
        }
    }
}

function Set-AllowedSecurityProtocols
{
    $allowedProtocols = @()
    $insecureProtocols = @([System.Net.SecurityProtocolType]::SystemDefault, [System.Net.SecurityProtocolType]::Ssl3)

    foreach ($protocol in [System.Enum]::GetValues([System.Net.SecurityProtocolType]))
    {
        if ($insecureProtocols -notcontains $protocol)
        {
            $allowedProtocols += $protocol
        }
    }

    Write-Log "Settings allowed security protocols to: $allowedProtocols"
    [System.Net.ServicePointManager]::SecurityProtocol = $allowedProtocols
}

function Set-WinRmServiceAutoStart
{
    Write-Log "Setting WinRM service start to auto"
    sc.exe config winrm start=auto
}

function Set-WinRmServiceDelayedStart
{
    # Hyper-V messes with networking components on startup after the feature is enabled
    # causing issues with communication over winrm and setting winrm to delayed start
    # gives Hyper-V enough time to finish configuration before having packer continue.
    Write-Log "Setting WinRM service start to delayed-auto"
    sc.exe config winrm start=delayed-auto
}

function Update-DefenderSignatures
{
    Write-Log "Updating windows defender signatures."
    Update-MpSignature
}

function Update-WindowsFeatures
{
    $featuresToEnable = @(
        "Containers",
        "Hyper-V",
        "Hyper-V-PowerShell")

    foreach ($feature in $featuresToEnable)
    {
        Write-Log "Enabling Windows feature: $feature"
        Install-WindowsFeature $feature
    }
}

# Disable progress writers for this session to greatly speed up operations such as Invoke-WebRequest
$ProgressPreference = 'SilentlyContinue'

switch ($env:ProvisioningPhase)
{
    "1"
    {
        Write-Log "Performing actions for provisioning phase 1"
        Set-WinRmServiceDelayedStart
        Set-AllowedSecurityProtocols
        Disable-WindowsUpdates
        Install-WindowsPatches
        Update-DefenderSignatures
        Install-OpenSSH
        Update-WindowsFeatures
    }
    "2"
    {
        Write-Log "Performing actions for provisioning phase 2"
        Set-WinRmServiceAutoStart
        Install-Docker
        Get-ContainerImages
        Get-FilesToCacheOnVHD
    }
    default
    {
        Write-Log "Unable to determine provisiong phase... exiting"
        exit 1
    }
}
