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

function CreateHypervisorRuntime {
  Param(
    [Parameter(Mandatory = $true)][string]
    $image,
    [Parameter(Mandatory = $true)][string]
    $version,
    [Parameter(Mandatory = $true)][string]
    $buildNumber
  )

  return @"
        [plugins.cri.containerd.runtimes.runhcs-wcow-hypervisor-$buildnumber]
          runtime_type = "io.containerd.runhcs.v1"
          [plugins.cri.containerd.runtimes.runhcs-wcow-hypervisor-$buildnumber.options]
            Debug = true
            DebugType = 2
            SandboxImage = "$image-windows-$version-amd64"
            SandboxPlatform = "windows/amd64"
            SandboxIsolation = 1
            ScaleCPULimitsToSandbox = true
"@
}

function CreateHypervisorRuntimes {
  Param(
    [Parameter(Mandatory = $true)][string[]]
    $builds,
    [Parameter(Mandatory = $true)][string]
    $image
  )
  
  Write-Host "Adding hyperv runtimes $builds"
  $hypervRuntimes = ""
  ForEach ($buildNumber in $builds) {
    $windowsVersion = Select-Windows-Version -buildNumber $buildNumber
    $runtime = createHypervisorRuntime -image $pauseImage -version $windowsVersion -buildNumber $buildNumber
    if ($hypervRuntimes -eq "") {
      $hypervRuntimes = $runtime
    }
    else {
      $hypervRuntimes = $hypervRuntimes + "`r`n" + $runtime
    }
  }

  return $hypervRuntimes
}

function Select-Windows-Version {
  param (
    [Parameter()]
    [string]
    $buildNumber
  )

  switch ($buildNumber) {
    "17763" { return "1809" }
    "18362" { return "1903" }
    "18363" { return "1909" }
    "19041" { return "2004" }
    Default { return "" } 
  }
}

function Enable-Logging {
  if ((Test-Path "$global:ContainerdInstallLocation\diag.ps1") -And (Test-Path "$global:ContainerdInstallLocation\ContainerPlatform.wprp")) {
    $logs = Join-path $pwd.drive.Root logs
    Write-Log "Containerd hyperv logging enabled; temp location $logs"
    $diag = Join-Path $global:ContainerdInstallLocation diag.ps1
    mkdir -Force $logs
    # !ContainerPlatformPersistent profile is made to work with long term and boot tracing
    & $diag -Start -ProfilePath "$global:ContainerdInstallLocation\ContainerPlatform.wprp!ContainerPlatformPersistent" -TempPath $logs
  }
  else {
    Write-Log "Containerd hyperv logging script not avalaible"
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
  }
  elseif ($ContainerdUrl.endswith(".tar.gz")) {
    # upstream containerd package is a tar 
    $tarfile = [Io.path]::Combine($ENV:TEMP, "containerd.tar.gz")
    DownloadFileOverHttp -Url $ContainerdUrl -DestinationPath $tarfile
    mkdir -Force $global:ContainerdInstallLocation
    tar -xzf $tarfile -C $global:ContainerdInstallLocation
    mv -Force $global:ContainerdInstallLocation\bin\* $global:ContainerdInstallLocation\
    del $tarfile
    del -Recurse -Force $global:ContainerdInstallLocation\bin
  }

  # get configuration options
  Add-SystemPathEntry $global:ContainerdInstallLocation
  $cdbinary = Join-Path $global:ContainerdInstallLocation containerd.exe
  $configFile = [Io.Path]::Combine($global:ContainerdInstallLocation, "config.toml")
  $clusterConfig = ConvertFrom-Json ((Get-Content $global:KubeClusterConfigPath -ErrorAction Stop) | Out-String)
  $pauseImage = $clusterConfig.Cri.Images.Pause
  $formatedbin = $(($CNIBinDir).Replace("\", "/"))
  $formatedconf = $(($CNIConfDir).Replace("\", "/"))
  $sandboxIsolation = 0
  $windowsVersion = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").ReleaseId
  $hypervRuntimes = ""
  $hypervHandlers = $global:HypervRuntimeHandlers.split(",", [System.StringSplitOptions]::RemoveEmptyEntries)

  # configure
  if ($global:DefaultContainerdRuntimeHandler -eq "hyperv") {
    Write-Log "default runtime for containerd set to hyperv"
    $sandboxIsolation = 1
  }

  $template = Get-Content -Path "c:\AzureData\k8s\containerdtemplate.toml" 
  if ($sandboxIsolation -eq 0 -And $hypervHandlers.Count -eq 0) {
    # remove the value hypervisor place holder
    $template = $template | Select-String -Pattern 'hypervisors' -NotMatch | Out-String
  }
  else {
    $hypervRuntimes = CreateHypervisorRuntimes -builds @($hypervHandlers) -image $pauseImage
  }

  $template.Replace('{{sandboxIsolation}}', $sandboxIsolation).
  Replace('{{pauseImage}}', $pauseImage).
  Replace('{{hypervisors}}', $hypervRuntimes).
  Replace('{{cnibin}}', $formatedbin).
  Replace('{{cniconf}}', $formatedconf).
  Replace('{{currentversion}}', $windowsVersion) | `
    Out-File -FilePath "$configFile" -Encoding ascii

  RegisterContainerDService
  Enable-Logging
}
