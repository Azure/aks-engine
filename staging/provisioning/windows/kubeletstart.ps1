$Global:ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string)

$global:MasterIP = $Global:ClusterConfiguration.Kubernetes.ControlPlane.IpAddress
$global:KubeDnsSearchPath = "svc.cluster.local"
$global:KubeDnsServiceIp = $Global:ClusterConfiguration.Kubernetes.Network.DnsIp
$global:MasterSubnet = $Global:ClusterConfiguration.Kubernetes.ControlPlane.MasterSubnet
$global:KubeClusterCIDR = $Global:ClusterConfiguration.Kubernetes.Network.ClusterCidr
$global:KubeServiceCIDR = $Global:ClusterConfiguration.Kubernetes.Network.ServiceCidr
$global:KubeBinariesVersion = $Global:ClusterConfiguration.Kubernetes.Source.Release
$global:KubeDir = $Global:ClusterConfiguration.Install.Destination
$global:NetworkMode = "L2Bridge"
$global:ExternalNetwork = "ext"
$global:CNIConfig = "$CNIConfig"
$global:HNSModule = "c:\k\hns.psm1"
$global:NetworkPlugin = $Global:ClusterConfiguration.Cni.Name
$global:KubeletNodeLabels = $Global:ClusterConfiguration.Kubernetes.Kubelet.NodeLabels
$global:ContainerRuntime = $Global:ClusterConfiguration.Cri.Name

$global:AzureCNIDir = [Io.path]::Combine("$global:KubeDir", "azurecni")
$global:AzureCNIBinDir = [Io.path]::Combine("$global:AzureCNIDir", "bin")
$global:AzureCNIConfDir = [Io.path]::Combine("$global:AzureCNIDir", "netconf")

$global:CNIPath = [Io.path]::Combine("$global:KubeDir", "cni")
$global:CNIConfig = [Io.path]::Combine($global:CNIPath, "config", "$global:NetworkMode.conf")
$global:CNIConfigPath = [Io.path]::Combine("$global:CNIPath", "config")

ipmo $global:HNSModule

#TODO ksbrmnn refactor to be sensical instead of if if if ...

# Calculate some local paths
$global:VolumePluginDir = [Io.path]::Combine($global:KubeDir, "volumeplugins")
mkdir $global:VolumePluginDir -Force

$KubeletArgList = $Global:ClusterConfiguration.Kubernetes.Kubelet.ConfigArgs # This is the initial list passed in from aks-engine
$KubeletArgList += "--node-labels=$global:KubeletNodeLabels"
# $KubeletArgList += "--hostname-override=$global:AzureHostname" TODO: remove - dead code?
$KubeletArgList += "--volume-plugin-dir=$global:VolumePluginDir"
# If you are thinking about adding another arg here, you should be considering pkg/engine/defaults-kubelet.go first
# Only args that need to be calculated or combined with other ones on the Windows agent should be added here.

# Configure kubelet to use CNI plugins if enabled.
if ($NetworkPlugin -eq "azure") {
    $KubeletArgList += @("--cni-bin-dir=$AzureCNIBinDir", "--cni-conf-dir=$AzureCNIConfDir")
}
elseif ($NetworkPlugin -eq "kubenet") {
    $KubeletArgList += @("--cni-bin-dir=$CNIPath", "--cni-conf-dir=$CNIConfigPath")
    # handle difference in naming between Linux & Windows reference plugin
    $KubeletArgList = $KubeletArgList -replace "kubenet", "cni"
}
else {
    throw "Unknown network type $NetworkPlugin, can't configure kubelet"
}

# Update args to use ContainerD if needed
if ($global:ContainerRuntime -eq "containerd") {
    $KubeletArgList += @("--container-runtime=remote", "--container-runtime-endpoint=npipe://./pipe/containerd-containerd")
}

# Used in WinCNI version of kubeletstart.ps1
$KubeletArgListStr = ""
$KubeletArgList | Foreach-Object {
    # Since generating new code to be written to a file, need to escape quotes again
    if ($KubeletArgListStr.length -gt 0) {
        $KubeletArgListStr = $KubeletArgListStr + ", "
    }
    # TODO ksbrmnn figure out what's going on here re tick marks
    $KubeletArgListStr = $KubeletArgListStr + "`"" + $_.Replace("`"`"", "`"`"`"`"") + "`""
}
$KubeletArgListStr = "@($KubeletArgListStr`)"

# Used in Azure-CNI version of kubeletstart.ps1
$KubeletCommandLine = "$global:KubeDir\kubelet.exe " + ($KubeletArgList -join " ")

# Turn off Firewall to enable pods to talk to service endpoints. (Kubelet should eventually do this)
# TODO move this to CSE
netsh advfirewall set allprofiles state off

function
Get-DefaultGateway($CIDR) {
    return $CIDR.substring(0, $CIDR.lastIndexOf(".")) + ".1"
}
function
Get-PodCIDR() {
    $podCIDR = c:\k\kubectl.exe --kubeconfig=c:\k\config get nodes/$($env:computername.ToLower()) -o custom-columns=podCidr:.spec.podCIDR --no-headers
    return $podCIDR
}

function
Test-PodCIDR($podCIDR) {
    return $podCIDR.length -gt 0
}

function
Update-CNIConfigKubenetDocker($podCIDR, $masterSubnetGW) {
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
function
Update-CNIConfigKubenetContainerD($podCIDR, $masterSubnetGW) {
    $jsonSampleConfig =
    "{
    ""cniVersion"": ""0.2.0"",
    ""name"": ""<NetworkMode>"",
    ""type"": ""sdnbridge.exe"",
    ""master"": ""Ethernet"",
    ""capabilities"": { ""portMappings"": true },
    ""ipam"": {
        ""environment"": ""azure"",
        ""subnet"":""<PODCIDR>"",
        ""routes"": [{
        ""GW"":""<PODGW>""
        }]
    },
    ""dns"" : {
    ""Nameservers"" : [ ""<NameServers>"" ],
    ""Search"" : [ ""<Cluster DNS Suffix or Search Path>"" ]
    },
    ""AdditionalArgs"" : [
    {
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""OutBoundNAT"", ""Settings"" : { ""Exceptions"": [ ""<ClusterCIDR>"", ""<MgmtSubnet>"" ] }}
    },
    {
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""SDNRoute"", ""Settings"" : { ""DestinationPrefix"": ""<ServiceCIDR>"", ""NeedEncap"" : true }}
    }
    ]
}"

    $configJson = ConvertFrom-Json $jsonSampleConfig
    $configJson.name = $global:NetworkMode.ToLower()
    $configJson.ipam.subnet = $podCIDR
    $configJson.ipam.routes[0].GW = $masterSubnetGW
    $configJson.dns.Nameservers[0] = $global:KubeDnsServiceIp
    $configJson.dns.Search[0] = $global:KubeDnsSearchPath


    $configJson.AdditionalArgs[0].Value.Settings.Exceptions[0] = $global:KubeClusterCIDR
    $configJson.AdditionalArgs[0].Value.Settings.Exceptions[1] = $global:MasterSubnet
    $configJson.AdditionalArgs[1].Value.Settings.DestinationPrefix = $global:KubeServiceCIDR

    if (Test-Path $global:CNIConfig) {
        Clear-Content -Path $global:CNIConfig
    }

    Write-Host "Generated CNI Config [$configJson]"

    Add-Content -Path $global:CNIConfig -Value (ConvertTo-Json $configJson -Depth 20)
}

# Required to clean up the HNS policy lists properly
Write-Host "Stopping kubeproxy service"
Stop-Service kubeproxy

if ($global:NetworkPlugin -eq "azure") {
    Write-Host "NetworkPlugin azure, starting kubelet."

    # Find if network created by CNI exists, if yes, remove it
    # This is required to keep the network non-persistent behavior
    # Going forward, this would be done by HNS automatically during restart of the node
    ./cleanupnetwork.ps1 
    
    # Restart Kubeproxy, which would wait, until the network is created
    # This was fixed in 1.15, workaround still needed for 1.14 https://github.com/kubernetes/kubernetes/pull/78612
    Restart-Service Kubeproxy

    # Set env file for Azure Stack
    $env:AZURE_ENVIRONMENT_FILEPATH = "c:\k\azurestackcloud.json"
}

if ($global:NetworkPlugin -eq "kubenet") {
    try {
        $env:AZURE_ENVIRONMENT_FILEPATH = "c:\k\azurestackcloud.json"

        $masterSubnetGW = Get-DefaultGateway $global:MasterSubnet
        $podCIDR = Get-PodCIDR
        $podCidrDiscovered = Test-PodCIDR($podCIDR)

        # if the podCIDR has not yet been assigned to this node, start the kubelet process to get the podCIDR, and then promptly kill it.
        if (-not $podCidrDiscovered) {
            $process = Start-Process -FilePath c:\k\kubelet.exe -PassThru -ArgumentList $KubeletArgList

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

        ./cleanupnetwork.ps1 

        Write-Host "Creating a new hns Network"
        $hnsNetwork = New-HNSNetwork -Type $global:NetworkMode -AddressPrefix $podCIDR -Gateway $masterSubnetGW -Name $global:NetworkMode.ToLower() -Verbose
        # New network has been created, Kubeproxy service has to be restarted
        # This was fixed in 1.15, workaround still needed for 1.14 https://github.com/kubernetes/kubernetes/pull/78612
        Restart-Service Kubeproxy

        Start-Sleep 10

        if  ($global:ContainerRuntime -eq "containerd") {
            Write-Host "Updating CNI config"
            Update-CNIConfigKubenetContainerD $podCIDR $masterSubnetGW
        }

        if  ($global:ContainerRuntime -eq "docker") {
            # Add route to all other POD networks
            Update-CNIConfigKubenetDocker $podCIDR $masterSubnetGW
        }
    }
    catch {
        Write-Error $_
    }
}

# Start the kubelet
# Use run-process.cs to set process priority class as 'High'
Add-Type -Path .\run-process.cs
$exe = "$global:KubeDir\kubelet.exe"
$args = ($KubeletArgList -join " ")
[RunProcess.exec]::RunProcess($exe, $args)