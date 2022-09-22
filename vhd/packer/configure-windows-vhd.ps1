<#
    .SYNOPSIS
        Used to produce Windows AKS images.

    .DESCRIPTION
        This script is used by packer to produce Windows AKS images.
#>

param()

$ErrorActionPreference = "Stop"

filter Timestamp { "$(Get-Date -Format o): $_" }

$global:containerdPackageUrl = "https://mobyartifacts.azureedge.net/moby/moby-containerd/1.5.8+azure/windows/windows_amd64/moby-containerd-1.5.8+azure-1.amd64.zip"

function Write-Log($Message) {
    $msg = $message | Timestamp
    Write-Output $msg
}

function Disable-WindowsUpdates {
    # See https://docs.microsoft.com/en-us/windows/deployment/update/waas-wu-settings
    # for additional information on WU related registry settings

    Write-Log "Disabling automatic windows upates"
    $WindowsUpdatePath = "HKLM:SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate"
    $AutoUpdatePath = "HKLM:SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU"

    if (Test-Path -Path $WindowsUpdatePath) {
        Remove-Item -Path $WindowsUpdatePath -Recurse
    }

    New-Item -Path $WindowsUpdatePath | Out-Null
    New-Item -Path $AutoUpdatePath | Out-Null
    Set-ItemProperty -Path $AutoUpdatePath -Name NoAutoUpdate -Value 1 | Out-Null
}

function Get-ContainerImages {
    param (
        $containerRuntime,
        $windowsServerVersion
    )

    switch ($windowsServerVersion) {
        '2019' {
            $imagesToPull = @(
                "mcr.microsoft.com/windows/servercore:ltsc2019",
                "mcr.microsoft.com/windows/nanoserver:1809",
                "mcr.microsoft.com/oss/kubernetes/pause:1.4.1",
                "mcr.microsoft.com/oss/kubernetes/pause:3.4.1",
                "mcr.microsoft.com/oss/kubernetes/azure-cloud-node-manager:v1.0.18",
                "mcr.microsoft.com/oss/kubernetes/azure-cloud-node-manager:v1.1.14",
                "mcr.microsoft.com/oss/kubernetes/azure-cloud-node-manager:v1.23.11",
                "mcr.microsoft.com/oss/kubernetes-csi/azuredisk-csi:v1.10.0",
                "mcr.microsoft.com/oss/kubernetes-csi/csi-node-driver-registrar:v2.4.0",
                "mcr.microsoft.com/oss/kubernetes-csi/livenessprobe:v2.5.0")
        }
        '2004' {
            $imagesToPull = @(
                "mcr.microsoft.com/windows/servercore:2004",
                "mcr.microsoft.com/windows/nanoserver:2004",
                "mcr.microsoft.com/oss/kubernetes/pause:1.4.1",
                "mcr.microsoft.com/oss/kubernetes/pause:3.4.1")
        }
        default {
            $imagesToPull = @()
        }
    }


    if ($containerRuntime -eq 'containerd') {
        # start containerd to pre-pull the images to disk on VHD
        # CSE will configure and register containerd as a service at deployment time
        Start-Job -Name containerd -ScriptBlock { containerd.exe }
        foreach ($image in $imagesToPull) {
            & ctr.exe -n k8s.io images pull $image
        }
        Stop-Job  -Name containerd
        Remove-Job -Name containerd
    }
    else {
        foreach ($image in $imagesToPull) {
            docker pull $image
        }
    }
}

function Get-FilesToCacheOnVHD {
    Write-Log "Caching misc files on VHD"

    $map = @{
        "c:\akse-cache\"              = @(
            "https://github.com/Azure/aks-engine/raw/master/scripts/collect-windows-logs.ps1",
            "https://github.com/Microsoft/SDN/raw/master/Kubernetes/flannel/l2bridge/cni/win-bridge.exe",
            "https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/collectlogs.ps1",
            "https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/dumpVfpPolicies.ps1",
            "https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/portReservationTest.ps1",
            "https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/starthnstrace.cmd",
            "https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/startpacketcapture.cmd",
            "https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/stoppacketcapture.cmd",
            "https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/debug/VFP.psm1",
            "https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/helper.psm1",
            "https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/hns.psm1",
            "https://globalcdn.nuget.org/packages/microsoft.applicationinsights.2.11.0.nupkg",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.3.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.4.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.8.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.9.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.10.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.11.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.12.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.13.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.14.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.15.zip",
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.16.zip"
        );
        "c:\akse-cache\containerd\"   = @(
            $global:containerdPackageUrl
        );
        "c:\akse-cache\csi-proxy\"    = @(
            "https://kubernetesartifacts.azureedge.net/csi-proxy/v0.2.2/binaries/csi-proxy-v0.2.2.tar.gz"
        );
        "c:\akse-cache\win-k8s\"      = @(
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.21.14/windowszip/v1.21.14-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.22.7/windowszip/v1.22.7-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.22.13/windowszip/v1.22.13-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.23.6/windowszip/v1.23.6-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.23.12/windowszip/v1.23.12-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.24.4/windowszip/v1.24.4-1int.zip"
        );
        "c:\akse-cache\win-vnet-cni\" = @(
            "https://kubernetesartifacts.azureedge.net/azure-cni/v1.4.13/binaries/azure-vnet-cni-singletenancy-windows-amd64-v1.4.13.zip",
            "https://kubernetesartifacts.azureedge.net/azure-cni/v1.4.14/binaries/azure-vnet-cni-singletenancy-windows-amd64-v1.4.14.zip",
            "https://kubernetesartifacts.azureedge.net/azure-cni/v1.4.16/binaries/azure-vnet-cni-singletenancy-windows-amd64-v1.4.16.zip"
        )
    }

    foreach ($dir in $map.Keys) {
        New-Item -ItemType Directory $dir -Force | Out-Null

        foreach ($URL in $map[$dir]) {
            $fileName = [IO.Path]::GetFileName($URL)
            $dest = [IO.Path]::Combine($dir, $fileName)

            Write-Log "Downloading $URL to $dest"
            curl.exe -f --retry 5 --retry-delay 0 -L $URL -o $dest
            if ($LASTEXITCODE) {
                throw "Curl exited with '$LASTEXITCODE' while attemping to downlaod '$URL'"
            }
        }
    }
}

function Install-ContainerD {
    Write-Log "Getting containerD binaries from $global:containerdPackageUrl"

    $installDir = "c:\program files\containerd"
    Write-Log "Installing containerd to $installDir"
    New-Item -ItemType Directory $installDir -Force | Out-Null

    if ($global:containerdPackageUrl.endswith(".zip")) {
        $zipPath = [IO.Path]::Combine($installDir, "containerd.zip")
        Invoke-WebRequest -UseBasicParsing -Uri $global:containerdPackageUrl -OutFile $zipPath
        Expand-Archive -path $zipPath -DestinationPath $installDir -Force
        Remove-Item -Path $zipPath | Out-Null
    } else {
        $tarPath = [IO.Path]::Combine($installDir, "containerd.tar.gz")
        Invoke-WebRequest -UseBasicParsing -Uri $global:containerdPackageUrl -OutFile $tarPath
        tar -xzf $tarPath --strip=1 -C $installDir
        Remove-Item -Path $tarPath | Out-Null
    }

    $newPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";$installDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, [EnvironmentVariableTarget]::Machine)
    $env:Path += ";$installDir"
}

function Install-Docker {
    $defaultDockerVersion = "20.10.9"

    Write-Log "Attempting to install Docker version $defaultDockerVersion"
    Install-PackageProvider -Name DockerMsftProvider -Force -ForceBootstrap | Out-Null
    $package = Find-Package -Name Docker -ProviderName DockerMsftProvider -RequiredVersion $defaultDockerVersion
    Write-Log "Installing Docker version $($package.Version)"
    $package | Install-Package -Force | Out-Null
    Start-Service docker
}

function Install-OpenSSH {
    Write-Log "Installing OpenSSH Server"
    Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
}

function Install-WindowsPatches {
    param (
        $windowsServerVersion
    )

    switch ($windowsServerVersion) {
        '2019' {
            # Windows Server 2019 update history can be found at https://support.microsoft.com/en-us/help/4464619
            # then you can get download links by searching for specific KBs at http://www.catalog.update.microsoft.com/home.aspx

            # Find a specific patch at https://www.catalog.update.microsoft.com/Search.aspx?q=kb5005625
            # KBKB5015880 Contains 2022 7C patches for Windows server 2019
            $patchUrls = @("https://catalog.s.download.windowsupdate.com/c/msdownload/update/software/updt/2022/07/windows10.0-kb5015880-x64_1b8eadaa8f12dacd77342353c257e236e167a802.msu")
        }
        '2004' {
            # Windows Server, Version 2004 update history can be found at https://support.microsoft.com/en-us/help/4555932
            # then you can get download links by searching for specific KBs at http://www.catalog.update.microsoft.com/home.aspx

            $patchUrls = @()
        }
        default {
            $patchUrls = @()
        }
    }

    foreach ($patchUrl in $patchUrls) {
        $pathOnly = $patchUrl.Split("?")[0]
        $fileName = Split-Path $pathOnly -Leaf
        $fileExtension = [IO.Path]::GetExtension($fileName)
        $fullPath = [IO.Path]::Combine($env:TEMP, $fileName)

        switch ($fileExtension) {
            ".msu" {
                Write-Log "Downloading windows patch from $pathOnly to $fullPath"
                Invoke-WebRequest -UseBasicParsing $patchUrl -OutFile $fullPath
                Write-Log "Starting install of $fileName"
                $proc = Start-Process -PassThru -FilePath wusa.exe -ArgumentList "$fullPath /quiet /norestart"
                Wait-Process -InputObject $proc
                switch ($proc.ExitCode) {
                    0 {
                        Write-Log "Finished install of $fileName"
                    }
                    3010 {
                        Write-Log "Finished install of $fileName. Reboot required"
                    }
                    default {
                        Write-Log "Error during install of $fileName. ExitCode: $($proc.ExitCode)"
                        exit 1
                    }
                }
            }
            default {
                Write-Log "Installing patches with extension $fileExtension is not currently supported."
                exit 1
            }
        }
    }
}

function Set-AllowedSecurityProtocols {
    $allowedProtocols = @()
    $insecureProtocols = @([System.Net.SecurityProtocolType]::SystemDefault, [System.Net.SecurityProtocolType]::Ssl3)

    foreach ($protocol in [System.Enum]::GetValues([System.Net.SecurityProtocolType])) {
        if ($insecureProtocols -notcontains $protocol) {
            $allowedProtocols += $protocol
        }
    }

    Write-Log "Settings allowed security protocols to: $allowedProtocols"
    [System.Net.ServicePointManager]::SecurityProtocol = $allowedProtocols
}

function Set-WinRmServiceAutoStart {
    Write-Log "Setting WinRM service start to auto"
    sc.exe config winrm start=auto
}

function Set-WinRmServiceDelayedStart {
    # Hyper-V messes with networking components on startup after the feature is enabled
    # causing issues with communication over winrm and setting winrm to delayed start
    # gives Hyper-V enough time to finish configuration before having packer continue.
    Write-Log "Setting WinRM service start to delayed-auto"
    sc.exe config winrm start=delayed-auto
}

function Update-DefenderSignatures {
    Write-Log "Updating windows defender signatures."
    Update-MpSignature
}

function Update-WindowsFeatures {
    $featuresToEnable = @(
        "Containers",
        "Hyper-V",
        "Hyper-V-PowerShell")

    foreach ($feature in $featuresToEnable) {
        Write-Log "Enabling Windows feature: $feature"
        Install-WindowsFeature $feature
    }
}

function Update-Registry {
    # if multple LB policies are included for same endpoint then HNS hangs.
    # this fix forces an error
    Write-Host "Enable a HNS fix in 2021-2C+"
    Set-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Services\hns\State" -Name HNSControlFlag -Value 1 -Type DWORD
}

# Disable progress writers for this session to greatly speed up operations such as Invoke-WebRequest
$ProgressPreference = 'SilentlyContinue'

$containerRuntime = $env:ContainerRuntime
$validContainerRuntimes = @('containerd', 'docker')
if (-not ($validContainerRuntimes -contains $containerRuntime)) {
    Write-Host "Unsupported container runtime: $containerRuntime"
    exit 1
}

$windowsServerVersion = $env:WindowsServerVersion
$validWindowsServerContainers = @('2019', '2004')
if (-not ($validWindowsServerContainers -contains $windowsServerVersion)) {
    Write-Host "Unsupported Windows Server version: $windowsServerVersion"
    exit 1
}

switch ($env:ProvisioningPhase) {
    "1" {
        Write-Log "Performing actions for provisioning phase 1"
        Set-WinRmServiceDelayedStart
        Set-AllowedSecurityProtocols
        Disable-WindowsUpdates
        Install-WindowsPatches -WindowsServerVersion $windowsServerVersion
        Update-DefenderSignatures
        Install-OpenSSH
        Update-WindowsFeatures
    }
    "2" {
        Write-Log "Performing actions for provisioning phase 2 for container runtime '$containerRuntime'"
        Set-WinRmServiceAutoStart
        if ($containerRuntime -eq 'containerd') {
            Install-ContainerD
        } else {
            Install-Docker
        }
        Update-Registry
        Get-ContainerImages -containerRuntime $containerRuntime -WindowsServerVersion $windowsServerVersion
        Get-FilesToCacheOnVHD
        (New-Guid).Guid | Out-File -FilePath 'c:\vhd-id.txt'
    }
    default {
        Write-Log "Unable to determine provisiong phase... exiting"
        exit 1
    }
}
