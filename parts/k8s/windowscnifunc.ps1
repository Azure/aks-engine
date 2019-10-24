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
function Get-DefaultGateway($CIDR) {
    return $CIDR.substring(0, $CIDR.lastIndexOf(".")) + ".1"
}

function Get-PodCIDR() {
    $podCIDR = c:\k\kubectl.exe --kubeconfig=c:\k\config get nodes/$($env:computername.ToLower()) -o custom-columns=podCidr:.spec.podCIDR --no-headers
    return $podCIDR
}

function Test-PodCIDR($podCIDR) {
    return $podCIDR.length -gt 0
}

function Write-WinCNIConfig {
    param(
        [string] $cniConfigPath,
        [string] $networkMode,
        [string] $kubeDnsServiceIp,
        [string] $kubeDnsSearchPath,
        [string] $kubeClusterCIDR,
        [string] $masterSubnet,
        [string] $kubeServiceCIDR
    )

    Write-Log "Writing CNI config for kubenet"

    $jsonSampleConfig =
    "{
    ""cniVersion"": ""0.2.0"",
    ""name"": ""<NetworkMode>"",
    ""type"": ""win-bridge"",
    ""master"": ""Ethernet"",
    ""dns"" : {
        ""Nameservers"" : [ ""<NameServers>"" ],
        ""Search"" : [ ""<Cluster DNS Suffix or Search Path>"" ]
    },
    ""policies"": [
    {
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""OutBoundNAT"", ""ExceptionList"": [ ""<ClusterCIDR>"", ""<MgmtSubnet>"" ] }
    },
    {
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""ROUTE"", ""DestinationPrefix"": ""<ServiceCIDR>"", ""NeedEncap"" : true }
    }
    ]
}"

    $configJson = ConvertFrom-Json $jsonSampleConfig
    $configJson.name = $networkMode.ToLower()
    $configJson.dns.Nameservers[0] = $kubeDnsServiceIp
    $configJson.dns.Search[0] = $kubeDnsSearchPath

    $configJson.policies[0].Value.ExceptionList[0] = $kubeClusterCIDR
    $configJson.policies[0].Value.ExceptionList[1] = $masterSubnet
    $configJson.policies[1].Value.DestinationPrefix = $kubeServiceCIDR

    if (Test-Path $cniConfigPath) {
        Clear-Content -Path $cniConfigPath
    }

    Write-Log "Generated CNI Config [$configJson]"

    Add-Content -Path $cniConfigPath -Value (ConvertTo-Json $configJson -Depth 20)
}

function Get-PodCIDRForNode {
    param(
        [string[]] $kubeletArgList
    )

    Write-Log "Attempting to get pod CIDR"
    $podCIDR = Get-PodCIDR
    $podCidrDiscovered = Test-PodCIDR($podCIDR)

    Write-Log "Staring kubelet with args: $kubeletArgList"

    # if the podCIDR has not yet been assigned to this node, start the kubelet process to get the podCIDR, and then promptly kill it.
    if (-not $podCidrDiscovered) {
        Write-Log "Staring kubelet with args: $kubeletArgList"

        $process = Start-Process -FilePath c:\k\kubelet.exe -PassThru -ArgumentList $kubeletArgList

        # run kubelet until podCidr is discovered
        Write-Log "waiting to discover pod CIDR"
        while (-not $podCidrDiscovered) {
            Write-Log "Sleeping for 10s, and then waiting to discover pod CIDR"
            Start-Sleep 10

            $podCIDR = Get-PodCIDR
            $podCidrDiscovered = Test-PodCIDR($podCIDR)
        }

        # stop the kubelet process now that we have our CIDR, discard the process output
        $process | Stop-Process | Out-Null
    }

    Write-Log "Pod CIDR: $podCIDR"
    return $podCIDR
}
