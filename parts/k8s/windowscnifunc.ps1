function Get-HnsPsm1
{
    Param(
        [string]
        $HnsUrl = "https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/hns.psm1",
        [Parameter(Mandatory=$true)][string]
        $HNSModule
    )
    DownloadFileOverHttp -Url $HnsUrl -DestinationPath "$HNSModule"
}

function Update-WinCNI
{
    Param(
        [string]
        $WinCniUrl = "https://github.com/Microsoft/SDN/raw/master/Kubernetes/flannel/l2bridge/cni/win-bridge.exe",
        [Parameter(Mandatory=$true)][string]
        $CNIPath
    )
    $wincni = "win-bridge.exe"
    $wincniFile = [Io.path]::Combine($CNIPath, $wincni)
    DownloadFileOverHttp -Url $WinCniUrl -DestinationPath $wincniFile
}

function Install-SdnBridge
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $Url,
        [Parameter(Mandatory=$true)][string]
        $CNIPath
    )

    $cnizip = [Io.path]::Combine($CNIPath, "cni.zip")
    DownloadFileOverHttp -Url $Url -DestinationPath $cnizip
    Expand-Archive -path $cnizip -DestinationPath $CNIPath
    del $cnizip
}

# TODO: Move the code that creates the wincni configuration file out of windowskubeletfunc.ps1 and put it here