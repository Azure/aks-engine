# this is $global to persist across all functions since this is dot-sourced
$global:ContainerdInstallLocation = "$Env:ProgramFiles\containerd"

function RegisterContainerDService {
  Assert-FileExists (Join-Path $global:ContainerdInstallLocation containerd.exe)

  Write-Host "Registering containerd as a service"
  $cdbinary = Join-Path $global:ContainerdInstallLocation containerd.exe
  $svc = Get-Service -Name containerd -ErrorAction SilentlyContinue
  if ($null -ne $svc) {
    & $cdbinary --unregister-service
  }
  & $cdbinary --register-service
  $svc = Get-Service -Name "containerd" -ErrorAction SilentlyContinue
  if ($null -eq $svc) {
    throw "containerd.exe did not get installed as a service correctly."
  }
  $svc | Start-Service
  if ($svc.Status -ne "Running") {
    throw "containerd service is not running"
  }
}


function Install-Containerd {
  Param(
    [Parameter(Mandatory = $true)][string]
    $ContainerdUrl,
    [Parameter(Mandatory = $true)][string]
    $CNIBinDir,
    [Parameter(Mandatory = $true)][string]
    $CNIConfDir
  )

  $svc = Get-Service -Name containerd -ErrorAction SilentlyContinue
  if ($null -ne $svc) {
    Write-Log "Stoping containerd service"
    $svc | Stop-Service
  }

  # TODO: check if containerd is already installed and is the same version before this.
  
  # Extract the package
  if ($ContainerdUrl.endswith(".zip")) {
    $zipfile = [Io.path]::Combine($ENV:TEMP, "containerd.zip")
    DownloadFileOverHttp -Url $ContainerdUrl -DestinationPath $zipfile
    Expand-Archive -path $zipfile -DestinationPath $global:ContainerdInstallLocation -Force
    del $zipfile
  }elseif ($ContainerdUrl.endswith(".tar.gz")) {
    # upstream containerd package is a tar 
    $tarfile = [Io.path]::Combine($ENV:TEMP, "containerd.tar.gz")
    DownloadFileOverHttp -Url $ContainerdUrl -DestinationPath $tarfile
    mkdir -Force "C:\Program Files\containerd"
    tar -xzf $tarfile -C $global:ContainerdInstallLocation
    mv $global:ContainerdInstallLocation\bin\* $global:ContainerdInstallLocation\
    del $tarfile
    del -Recurse -Force $global:ContainerdInstallLocation\bin
  }

  # get configuration options
  Add-SystemPathEntry $global:ContainerdInstallLocation
  $cdbinary = Join-Path $global:ContainerdInstallLocation containerd.exe
  $configFile = [Io.Path]::Combine($global:ContainerdInstallLocation, "config.toml")
  $clusterConfig = ConvertFrom-Json ((Get-Content $global:KubeClusterConfigPath -ErrorAction Stop) | Out-String)
  $pauseImage = $clusterConfig.Cri.Images.Pause
  $formatedbin=$(($CNIBinDir).Replace("\","/"))
  $formatedconf=$(($CNIConfDir).Replace("\","/"))

  # configure
  if (Test-Path -Path "$global:ContainerdInstallLocation\containerd-template.toml") {
    # Hyperv config set up
    Write-Host "Configuring containerd for hyperv"
    $windowsVersion = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").ReleaseId
    (Get-Content -Path "c:\AzureData\k8s\containerdhypervtemplate.toml" -Raw).
      Replace('{{pauseImage}}', $pauseImage).
      Replace('{{cnibin}}', $formatedbin).
      Replace('{{cniconf}}', $formatedconf).
      Replace('{{currentversion}}', $windowsVersion) | `
      Out-File -FilePath "$configFile" -Encoding ascii
  }else{
    # non hyperv configuration
    Write-Host "Configuring containerd for process isolated"
    & $cdbinary config dump | `
      % {$_ -replace "sandbox_image = ""(.*?[^\\])""", "sandbox_image = ""$pauseImage""" } | `
      % {$_ -replace "bin_dir = ""(.*?[^\\])""", "bin_dir = ""$formatedbin""" } | `
      % {$_ -replace "conf_dir = ""(.*?[^\\])""", "conf_dir = ""$formatedconf""" } | `
      Out-File $configFile -Encoding ascii
  }

  RegisterContainerDService
}

