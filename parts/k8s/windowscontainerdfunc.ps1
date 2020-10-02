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
  $zipfile = [Io.path]::Combine($ENV:TEMP, "containerd.zip")
  DownloadFileOverHttp -Url $ContainerdUrl -DestinationPath $zipfile
  Expand-Archive -path $zipfile -DestinationPath $global:ContainerdInstallLocation -Force
  del $zipfile

  Add-SystemPathEntry $global:ContainerdInstallLocation

  # TODO: remove if the node comes up without this code
  # $configDir = [Io.Path]::Combine($ENV:ProgramData, "containerd")
  # if (-Not (Test-Path $configDir)) {
  #     mkdir $configDir
  # }

  # TODO: call containerd.exe dump config, then modify instead of starting with hardcoded
  $configFile = [Io.Path]::Combine($global:ContainerdInstallLocation, "config.toml")

  $clusterConfig = ConvertFrom-Json ((Get-Content $global:KubeClusterConfigPath -ErrorAction Stop) | Out-String)
  $pauseImage = $clusterConfig.Cri.Images.Pause

  @"
version = 2
root = "C:\\ProgramData\\containerd\\root"
state = "C:\\ProgramData\\containerd\\state"
plugin_dir = ""
disabled_plugins = []
required_plugins = []
oom_score = 0

[grpc]
  address = "\\\\.\\pipe\\containerd-containerd"
  tcp_address = ""
  tcp_tls_cert = ""
  tcp_tls_key = ""
  uid = 0
  gid = 0
  max_recv_message_size = 16777216
  max_send_message_size = 16777216

[ttrpc]
  address = ""
  uid = 0
  gid = 0

[debug]
  address = ""
  uid = 0
  gid = 0
  level = ""

[metrics]
  address = ""
  grpc_histogram = false

[cgroup]
  path = ""

[timeouts]
  "io.containerd.timeout.shim.cleanup" = "5s"
  "io.containerd.timeout.shim.load" = "5s"
  "io.containerd.timeout.shim.shutdown" = "3s"
  "io.containerd.timeout.task.state" = "2s"

[plugins]
  [plugins."io.containerd.gc.v1.scheduler"]
    pause_threshold = 0.02
    deletion_threshold = 0
    mutation_threshold = 100
    schedule_delay = "0s"
    startup_delay = "100ms"
  [plugins."io.containerd.grpc.v1.cri"]
    disable_tcp_service = true
    stream_server_address = "127.0.0.1"
    stream_server_port = "0"
    stream_idle_timeout = "4h0m0s"
    enable_selinux = false
    sandbox_image = "$pauseImage"
    stats_collect_period = 10
    systemd_cgroup = false
    enable_tls_streaming = false
    max_container_log_line_size = 16384
    disable_cgroup = false
    disable_apparmor = false
    restrict_oom_score_adj = false
    max_concurrent_downloads = 3
    disable_proc_mount = false
    [plugins."io.containerd.grpc.v1.cri".containerd]
      snapshotter = "windows"
      default_runtime_name = "runhcs-wcow-process"
      no_pivot = false
      [plugins."io.containerd.grpc.v1.cri".containerd.default_runtime]
        runtime_type = ""
        runtime_engine = ""
        runtime_root = ""
        privileged_without_host_devices = false
      [plugins."io.containerd.grpc.v1.cri".containerd.untrusted_workload_runtime]
        runtime_type = ""
        runtime_engine = ""
        runtime_root = ""
        privileged_without_host_devices = false
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runhcs-wcow-process]
          runtime_type = "io.containerd.runhcs.v1"
          runtime_engine = ""
          runtime_root = ""
          privileged_without_host_devices = false
    [plugins."io.containerd.grpc.v1.cri".cni]
      bin_dir = "$(($CNIBinDir).Replace("\","//"))"
      conf_dir = "$(($CNIConfDir).Replace("\","//"))"
      max_conf_num = 1
      conf_template = ""
    [plugins."io.containerd.grpc.v1.cri".registry]
      [plugins."io.containerd.grpc.v1.cri".registry.mirrors]
        [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
          endpoint = ["https://registry-1.docker.io"]
    [plugins."io.containerd.grpc.v1.cri".x509_key_pair_streaming]
      tls_cert_file = ""
      tls_key_file = ""
  [plugins."io.containerd.metadata.v1.bolt"]
    content_sharing_policy = "shared"
  [plugins."io.containerd.runtime.v2.task"]
    platforms = ["windows/amd64", "linux/amd64"]
  [plugins."io.containerd.service.v1.diff-service"]
    default = ["windows", "windows-lcow"]
"@ | Out-File -Encoding ascii $configFile

  RegisterContainerDService
}