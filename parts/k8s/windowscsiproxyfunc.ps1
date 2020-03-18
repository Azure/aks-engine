function New-CsiProxyService {
    Param(
        [Parameter(Mandatory = $true)][string]
        $CsiProxyPackageUrl,
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )

    $tempdir = New-TemporaryDirectory
    $binaryPackage = "$tempdir\csiproxy.tar"

    DownloadFileOverHttp -Url $CsiProxyPackageUrl -DestinationPath $binaryPackage

    tar -xzf $binaryPackage -C $tempdir
    cp "$tempdir\build\server.exe" "$KubeDir\csi-proxy-server.exe"

    del $tempdir -Recurse

    & "$KubeDir\nssm.exe" install csi-proxy-server "$KubeDir\csi-proxy-server.exe"
    & "$KubeDir\nssm.exe" set csi-proxy-server AppDirectory "$KubeDir"
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRestartDekay 5000
    & "$KubeDir\nssm.exe" set csi-proxy-server Description csi-proxy-server
    & "$KubeDir\nssm.exe" set csi-proxy-server Start SERVICE_DEMAND_START
    & "$KubeDir\nssm.exe" set csi-proxy-server ObjectName LocalSystem
    & "$KubeDir\nssm.exe" set csi-proxy-server Type SERVICE_WIN32_OWN_PROCESS
    & "$KubeDir\nssm.exe" set csi-proxy-server AppThrottle 1500
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStdout "$KubeDir\csi-proxy-server.log"
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStderr "$KubeDir\csi-proxy-server.err.log"
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStdoutCreationDisposition 4
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStderrCreationDisposition 4
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateFiles 1
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateOnline 1
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateSeconds 86400
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateBytes 10485760
}