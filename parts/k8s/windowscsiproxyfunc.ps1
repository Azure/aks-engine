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
    cp "$tempdir\bin\csi-proxy.exe" "$KubeDir\csi-proxy.exe"

    del $tempdir -Recurse

    & "$KubeDir\nssm.exe" install csi-proxy "$KubeDir\csi-proxy.exe" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppDirectory "$KubeDir" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRestartDelay 5000 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy Description csi-proxy | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStdout "$KubeDir\csi-proxy.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStderr "$KubeDir\csi-proxy.err.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStdoutCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStderrCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateBytes 10485760 | RemoveNulls
}