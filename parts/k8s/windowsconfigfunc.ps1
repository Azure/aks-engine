

# Set the service telemetry GUID. This is used with Windows Analytics https://docs.microsoft.com/en-us/sccm/core/clients/manage/monitor-windows-analytics
function Set-TelemetrySetting
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $WindowsTelemetryGUID
    )
    Set-ItemProperty -Path "HKLM:\Software\Microsoft\Windows\CurrentVersion\Policies\DataCollection" -Name "CommercialId" -Value $WindowsTelemetryGUID -Force
}

# Resize the system partition to the max available size. Azure can resize a managed disk, but the VM still needs to extend the partition boundary
function Resize-OSDrive
{
    $osDrive = ((Get-WmiObject Win32_OperatingSystem).SystemDrive).TrimEnd(":")
    $size = (Get-Partition -DriveLetter $osDrive).Size
    $maxSize = (Get-PartitionSupportedSize -DriveLetter $osDrive).SizeMax
    if ($size -lt $maxSize)
    {
        Resize-Partition -DriveLetter $osDrive -Size $maxSize
    }
}

# https://docs.microsoft.com/en-us/powershell/module/storage/new-partition
function Initialize-DataDisks
{
    Get-Disk | Where-Object PartitionStyle -eq 'raw' | Initialize-Disk -PartitionStyle MBR -PassThru | New-Partition -UseMaximumSize -AssignDriveLetter | Format-Volume -FileSystem NTFS -Force
}

# Set the Internet Explorer to use the latest rendering mode on all sites
# https://docs.microsoft.com/en-us/windows-hardware/customize/desktop/unattend/microsoft-windows-ie-internetexplorer-intranetcompatibilitymode
# (This only affects installations with UI)
function Set-Explorer
{
    New-Item -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer"
    New-Item -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\BrowserEmulation"
    New-ItemProperty -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\BrowserEmulation" -Name IntranetCompatibilityMode -Value 0 -Type DWord
    New-Item -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\Main"
    New-ItemProperty -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\Main" -Name "Start Page" -Type String -Value http://bing.com
}

function Install-Docker
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $DockerVersion
    )

    # DOCKER_API_VERSION needs to be set for Docker versions older than 18.09.0 EE
    # due to https://github.com/kubernetes/kubernetes/issues/69996
    # this issue was fixed by https://github.com/kubernetes/kubernetes/issues/69996#issuecomment-438499024
    # Note: to get a list of all versions, use this snippet
    # $versions = (curl.exe -L "https://go.microsoft.com/fwlink/?LinkID=825636&clcid=0x409" | ConvertFrom-Json).Versions | Get-Member -Type NoteProperty | Select-Object Name
    # Docker version to API version decoder: https://docs.docker.com/develop/sdk/#api-version-matrix

    switch ($DockerVersion.Substring(0,5))
    {
        "17.06" {
            Write-Log "Docker 17.06 found, setting DOCKER_API_VERSION to 1.30"
            [System.Environment]::SetEnvironmentVariable('DOCKER_API_VERSION', '1.30', [System.EnvironmentVariableTarget]::Machine)
        }

        "18.03" {
            Write-Log "Docker 18.03 found, setting DOCKER_API_VERSION to 1.37"
            [System.Environment]::SetEnvironmentVariable('DOCKER_API_VERSION', '1.37', [System.EnvironmentVariableTarget]::Machine)
        }

        default {
            Write-Log "Docker version $DockerVersion found, clearing DOCKER_API_VERSION"
            [System.Environment]::SetEnvironmentVariable('DOCKER_API_VERSION', $null, [System.EnvironmentVariableTarget]::Machine)
        }
    }

    try {
        $installDocker = $true
        $dockerService = Get-Service | ? Name -like 'docker'
        if ($dockerService.Count -eq 0) {
            Write-Log "Docker is not installed. Install docker version($DockerVersion)."
        }
        else {
            $dockerServerVersion = docker version --format '{{.Server.Version}}'
            Write-Log "Docker service is installed with docker version($dockerServerVersion)."
            if ($dockerServerVersion -eq $DockerVersion) {
                $installDocker = $false
                Write-Log "Same version docker installed will skip installing docker version($dockerServerVersion)."
            }
            else {
                Write-Log "Same version docker is not installed. Will install docker version($DockerVersion)."
            }
        }

        if ($installDocker) {
            Find-Package -Name Docker -ProviderName DockerMsftProvider -RequiredVersion $DockerVersion -ErrorAction Stop
            Write-Log "Found version $DockerVersion. Installing..."
            Install-Package -Name Docker -ProviderName DockerMsftProvider -Update -Force -RequiredVersion $DockerVersion
            net start docker
            Write-Log "Installed version $DockerVersion"
        }
    } catch {
        Write-Log "Error while installing package: $_.Exception.Message"
        $currentDockerVersion = (Get-Package -Name Docker -ProviderName DockerMsftProvider).Version
        Write-Log "Not able to install docker version. Using default version $currentDockerVersion"
    }
}

function Set-DockerLogFileOptions {
    Write-Log "Updating log file options in docker config"
    $dockerConfigPath = "C:\ProgramData\docker\config\daemon.json"

    if (-not (Test-Path $dockerConfigPath)) {
        "{}" | Out-File $dockerConfigPath
    }

    $dockerConfig = Get-Content $dockerConfigPath | ConvertFrom-Json
    $dockerConfig | Add-Member -Name "log-driver" -Value "json-file" -MemberType NoteProperty
    $logOpts = @{ "max-size" = "50m"; "max-file" = "5" }
    $dockerConfig | Add-Member -Name "log-opts" -Value $logOpts -MemberType NoteProperty
    $dockerConfig = $dockerConfig | ConvertTo-Json -Depth 10

    Write-Log "New docker config:"
    Write-Log $dockerConfig

    # daemon.json MUST be encoded as UTF8-no-BOM!
    Remove-Item $dockerConfigPath
    $fileEncoding = New-Object System.Text.UTF8Encoding $false
    [IO.File]::WriteAllLInes($dockerConfigPath, $dockerConfig, $fileEncoding)

    Restart-Service docker
}

# Pagefile adjustments
function Adjust-PageFileSize()
{
    wmic pagefileset set InitialSize=8096,MaximumSize=8096
}

function Adjust-DynamicPortRange()
{
    # Kube-proxy reserves 63 ports per service which limits clusters with Windows nodes
    # to ~225 services if default port reservations are used.
    # https://docs.microsoft.com/en-us/virtualization/windowscontainers/kubernetes/common-problems#load-balancers-are-plumbed-inconsistently-across-the-cluster-nodes
    # Kube-proxy load balancing should be set to DSR mode when it releases with future versions of the OS

    Invoke-Executable -Executable "netsh.exe" -ArgList @("int", "ipv4", "set", "dynamicportrange", "tcp", "16385", "49151")
}

# TODO: should this be in this PR?
# Service start actions. These should be split up later and included in each install step
function Update-ServiceFailureActions
{
    Param(
        [Parameter(Mandatory = $true)][string]
        $ContainerRuntime
    )
    sc.exe failure "kubelet" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure "kubeproxy" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure $ContainerRuntime actions= restart/60000/restart/60000/restart/60000 reset= 900
}

function Add-SystemPathEntry
{
    Param(
        [Parameter(Mandatory = $true)][string]
        $Directory
    )
    # update the path variable if it doesn't have the needed paths
    $path = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
    $updated = $false
    if(-not ($path -match $Directory.Replace("\","\\")+"(;|$)"))
    {
        $path += ";"+$Directory
        $updated = $true
    }
    if($updated)
    {
        Write-Output "Updating path, added $Directory"
        [Environment]::SetEnvironmentVariable("Path", $path, [EnvironmentVariableTarget]::Machine)
        $env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
    }
}