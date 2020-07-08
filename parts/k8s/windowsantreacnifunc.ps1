function Install-OpenvSwitch
{
    Param(
        [string]
        $VCUrl = "https://download.microsoft.com/download/1/6/5/165255E7-1014-4D0A-B094-B6A430A6BFFC/vcredist_x64.exe",
        [string]
        $OVSUrl = "https://raw.githubusercontent.com/vmware-tanzu/antrea/master/hack/windows/Install-OVS.ps1",
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )

    # TODO: Remove this with Antrea v0.10.0 release
    Write-Log "Installing Microsoft VC 2010"
    $VCPath = [Io.path]::Combine($KubeDir, "VC.exe")
    DownloadFileOverHttp -Url $VCUrl -DestinationPath $VCPath
    Start-Process -FilePath $VCPath -Args '/install /passive /norestart' -Verb RunAs -Wait

    # Enable TESTSIGNING so that OVS datapath can be enabled. This settings is not recommended for production
    & Bcdedit.exe -set TESTSIGNING ON

    # TODO: Discuss with AKS Folks, how to install Upstream OVS on Windows Node
    # and sign it.
    Write-Log "Installating Openvswitch"
    $OVSPsPath = [Io.path]::Combine($KubeDir, "Install-OVS.ps1")
    DownloadFileOverHttp -Url $OVSUrl -DestinationPath $OVSPsPath
    & $OVSPsPath

    $OVSVersion = ovs-vswitchd.exe --version | %{ $_.Split(' ')[-1]; }
    if ([string]::IsNullOrEmpty($OVSVersion)) {
        Write-Log "Create ovsdb file failed due to OVS version not found, exit"
        $OVSVersion = "2.13.1"
    }
    & ovs-vsctl.exe --no-wait set Open_vSwitch . ovs_version=$OVSVersion
}

$AntreaStartWrapper= '
Param(
    [string]
    $KubeConfig="c:\k\config",
    [string]
    $AntreaStartPs="c:\k\AntreaStart.ps1"
)
try {
   & $AntreaStartPs -KubeConfig $KubeConfig -StartKubeProxy $false
   Wait-Process antrea-agent
}
finally {
   Stop-Process -Name antrea-agent -Force
}
'

function Install-Antrea
{
    Param(
        [string]
        $AntreaUrl = "https://github.com/vmware-tanzu/antrea/releases/download/v0.9.3/Start.ps1",
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )

    Write-Log "Downloading Antrea Start Powershell script"
    $AntreaStartPs = [Io.path]::Combine($KubeDir, "AntreaStart.ps1")
    $KubeConfig = [Io.path]::Combine($KubeDir, "config")
    DownloadFileOverHttp -Url $AntreaUrl -DestinationPath $AntreaStartPs

    # Write Antrea Start Wrapper Powershell file
    Set-Content -Path $KubeDir\AntreaStartWrapper.ps1 -Value $AntreaStartWrapper

    Write-Log "Register Antrea Start as a service"
    mkdir $KubeDir\antrea
    & "$KubeDir\nssm.exe" install antrea-agent C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppDirectory $KubeDir | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppParameters "$KubeDir\AntreaStartWrapper.ps1 -kubeconfig $KubeConfig " | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppStdout "$KubeDir\antrea\stdout.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppStderr "$KubeDir\antrea\stderr.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set antrea-agent AppRotateBytes 10485760 | RemoveNulls
}
