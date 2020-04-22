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

    & "$KubeDir\nssm.exe" install csi-proxy-server "$KubeDir\csi-proxy-server.exe" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppDirectory "$KubeDir" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRestartDekay 5000 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server Description csi-proxy-server | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStdout "$KubeDir\csi-proxy-server.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStderr "$KubeDir\csi-proxy-server.err.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStdoutCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppStderrCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy-server AppRotateBytes 10485760 | RemoveNulls
}