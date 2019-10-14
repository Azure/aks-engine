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

# TODO: Move the code that creates the wincni configuration file out of windowskubeletfunc.ps1 and put it here

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

function Update-CNIConfig($podCIDR, $masterSubnetGW) {
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
    $configJson.name = $global:NetworkMode.ToLower()
    $configJson.dns.Nameservers[0] = $global:KubeDnsServiceIp
    $configJson.dns.Search[0] = $global:KubeDnsSearchPath

    $configJson.policies[0].Value.ExceptionList[0] = $global:KubeClusterCIDR
    $configJson.policies[0].Value.ExceptionList[1] = $global:MasterSubnet
    $configJson.policies[1].Value.DestinationPrefix = $global:KubeServiceCIDR

    if (Test-Path $global:CNIConfig) {
        Clear-Content -Path $global:CNIConfig
    }

    Write-Host "Generated CNI Config [$configJson]"

    Add-Content -Path $global:CNIConfig -Value (ConvertTo-Json $configJson -Depth 20)
}

function Write-WinCNIConfig ($kubeletArgList) {
    $masterSubnetGW = Get-DefaultGateway $global:MasterSubnet
    $podCIDR = Get-PodCIDR
    $podCidrDiscovered = Test-PodCIDR($podCIDR)

    # if the podCIDR has not yet been assigned to this node, start the kubelet process to get the podCIDR, and then promptly kill it.
    if (-not $podCidrDiscovered) {
        Write-Host "Staring kubelet with args: $kubeletArgList"
        $process = Start-Process -FilePath c:\k\kubelet.exe -PassThru -ArgumentList $kubeletArgList

        # run kubelet until podCidr is discovered
        Write-Host "waiting to discover pod CIDR"
        while (-not $podCidrDiscovered) {
            Write-Host "Sleeping for 10s, and then waiting to discover pod CIDR"
            Start-Sleep 10

            $podCIDR = Get-PodCIDR
            $podCidrDiscovered = Test-PodCIDR($podCIDR)
        }

        # stop the kubelet process now that we have our CIDR, discard the process output
        $process | Stop-Process | Out-Null
    }

    Write-Host "Updating CNI config - PodCIRD: $podCIDR, MasterSubnetGW: $masterSubnetGW"
    Update-CNIConfig $podCIDR $masterSubnetGW
}

