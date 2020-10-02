<#
    .SYNOPSIS
        Used to produce Windows AKS images.

    .DESCRIPTION
        This script is used by packer to produce Windows AKS images.
#>

param()

$ErrorActionPreference = "Stop"

filter Timestamp { "$(Get-Date -Format o): $_" }

$global:containerdPackageUrl = "https://github.com/containerd/containerd/releases/download/v1.4.1/containerd-1.4.1-windows-amd64.tar.gz"

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
                "mcr.microsoft.com/oss/kubernetes/pause:1.4.0",
                "mcr.microsoft.com/oss/kubernetes-csi/livenessprobe:v2.0.1-alpha.1-windows-1809-amd64",
                "mcr.microsoft.com/oss/kubernetes-csi/csi-node-driver-registrar:v1.2.1-alpha.1-windows-1809-amd64")
        }
        '2004' {
            $imagesToPull = @(
                "mcr.microsoft.com/windows/servercore:2004",
                "mcr.microsoft.com/windows/nanoserver:2004",
                "mcr.microsoft.com/oss/kubernetes/pause:1.4.0")
        }
        default { 
            $imagesToPull = @()
        }
    }


    if ($containerRuntime -eq 'containerd') {
        foreach ($image in $imagesToPull) {
            & ctr.exe -n k8s.io images pull $image
        }
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
            "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-v0.0.3.zip"
        );
        "c:\akse-cache\containerd\"   = @(
            $global:containerdPackageUrl
        );

        "c:\akse-cache\win-k8s\"      = @(
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.11-azs/windowszip/v1.15.11-azs-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.12-azs/windowszip/v1.15.12-azs-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.16.13-azs/windowszip/v1.16.13-azs-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.16.14-azs/windowszip/v1.16.14-azs-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.9-azs/windowszip/v1.17.9-azs-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.11-azs/windowszip/v1.17.11-azs-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.10/windowszip/v1.15.10-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.11/windowszip/v1.15.11-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.11-hotfix.20200714/windowszip/v1.15.11-hotfix.20200714-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.11-hotfix.20200817/windowszip/v1.15.11-hotfix.20200817-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.12/windowszip/v1.15.12-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.12-hotfix.20200714/windowszip/v1.15.12-hotfix.20200714-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.12-hotfix.20200623/windowszip/v1.15.12-hotfix.20200623-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.15.12-hotfix.20200817/windowszip/v1.15.12-hotfix.20200817-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.16.13/windowszip/v1.16.13-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.16.13-hotfix.20200714/windowszip/v1.16.13-hotfix.20200714-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.16.13-hotfix.20200817/windowszip/v1.16.13-hotfix.20200817-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.16.14/windowszip/v1.16.14-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.16.15/windowszip/v1.16.15-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.9/windowszip/v1.17.9-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.9-hotfix.20200714/windowszip/v1.17.9-hotfix.20200714-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.9-hotfix.20200817/windowszip/v1.17.9-hotfix.20200817-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.11/windowszip/v1.17.11-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.12/windowszip/v1.17.12-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.18.6/windowszip/v1.18.6-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.18.6-hotfix.20200723/windowszip/v1.18.6-hotfix.20200723-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.18.8/windowszip/v1.18.8-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.18.9/windowszip/v1.18.9-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.19.1/windowszip/v1.19.1-1int.zip",
            "https://kubernetesartifacts.azureedge.net/kubernetes/v1.19.2/windowszip/v1.19.2-1int.zip"
        );
        "c:\akse-cache\win-vnet-cni\" = @(
            "https://kubernetesartifacts.azureedge.net/azure-cni/v1.1.2/binaries/azure-vnet-cni-singletenancy-windows-amd64-v1.1.2.zip",
            "https://kubernetesartifacts.azureedge.net/azure-cni/v1.1.3/binaries/azure-vnet-cni-singletenancy-windows-amd64-v1.1.3.zip"
            "https://kubernetesartifacts.azureedge.net/azure-cni/v1.1.6/binaries/azure-vnet-cni-singletenancy-windows-amd64-v1.1.6.zip"
        )
    }

    foreach ($dir in $map.Keys) {
        New-Item -ItemType Directory $dir -Force | Out-Null

        foreach ($URL in $map[$dir]) {
            $fileName = [IO.Path]::GetFileName($URL)
            $dest = [IO.Path]::Combine($dir, $fileName)

            Write-Log "Downloading $URL to $dest"
            Invoke-WebRequest -UseBasicParsing -Uri $URL -OutFile $dest
        }
    }
}

function Install-ContainerD {
    Write-Log "Getting containerD binaries from $global:containerdPackageUrl"

    $installDir = "c:\program files\containerd"
    $tarPath = [IO.Path]::Combine($installDir, "containerd.tar.gz")

    Write-Log "Installing containerd to $installDir"
    New-Item -ItemType Directory $installDir -Force | Out-Null
    Invoke-WebRequest -UseBasicParsing -Uri $global:containerdPackageUrl -OutFile $tarPath
    tar -xzf $tarPath --strip=1 -C $installDir
    Remove-Item -Path $tarPath | Out-Null

    $newPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";$installDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, [EnvironmentVariableTarget]::Machine)
    $env:Path += ";$installDir"

    Write-Log "Registering containerd as a service"
    & containerd.exe --register-service
    $svc = Get-Service -Name "containerd" -ErrorAction SilentlyContinue
    if ($null -eq $svc) {
        throw "containerd.exe did not get installed as a service correctly."
    }

    Write-Log "Starting containerd service"
    $svc | Start-Service
    if ($svc.Status -ne "Running") {
        throw "containerd service is not running"
    }
}

function Install-Docker {
    $defaultDockerVersion = "19.03.11"

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

            # KB4558998 contains August 11, 2020 cumulative updates for Windows Server 2019
            # https://www.catalog.update.microsoft.com/Search.aspx?q=KB4565349
            $patchUrls = @("http://download.windowsupdate.com/d/msdownload/update/software/secu/2020/08/windows10.0-kb4565349-x64_919b9f31d4ccfa91183fbb9bab8c2975529e66b6.msu")
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
        # TODO: make decision on if we want to install docker along with containerd (will need to update CSE too,)
        Install-Docker
        if ($containerRuntime -eq 'containerd') {
            Install-ContainerD
        }
        Get-ContainerImages -containerRuntime $containerRuntime -WindowsServerVersion $windowsServerVersion
        Get-FilesToCacheOnVHD
        (New-Guid).Guid | Out-File -FilePath 'c:\vhd-id.txt'
    }
    default {
        Write-Log "Unable to determine provisiong phase... exiting"
        exit 1
    }
}
