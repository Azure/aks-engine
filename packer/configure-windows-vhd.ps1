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


    foreach ($image in $imagesToPull)
    {
        docker pull $image
    }
}

function Install-Docker
{
    $defaultDockerVersion = "18.09"

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

    $patchUrls = @(
        # 7C SSU and LCU
        'http://download.windowsupdate.com/c/msdownload/update/software/secu/2019/07/windows10.0-kb4512937-x64_2a065a9ecfee76e3e457f3c596550e821358971c.msu',
        'http://download.windowsupdate.com/d/msdownload/update/software/updt/2019/07/windows10.0-kb4505658-x64_cb660f3191eba56217694635b48d50d36883c3f2.msu')

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

switch ($env:ProvisioningPhase)
{
    "1"
    {
        Write-Log "Performing actions for provisioning phase 1"
        Disable-WindowsUpdates
        Install-WindowsPatches
        Install-OpenSSH
        Update-WindowsFeatures
    }
    "2"
    {
        Write-Log "Performing actions for provisioning phase 2"
        Install-Docker
        Get-ContainerImages
    }
    default
    {
        Write-Log "Unable to determine provisiong phase... exiting"
        exit 1
    }
}
