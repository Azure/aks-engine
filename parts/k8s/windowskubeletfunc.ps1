function
Write-AzureConfig
{
    Param(
        
        [Parameter(Mandatory=$true)][string]
        $AADClientId,
        [Parameter(Mandatory=$true)][string]
        $AADClientSecret,
        [Parameter(Mandatory=$true)][string]
        $TenantId,
        [Parameter(Mandatory=$true)][string]
        $SubscriptionId,
        [Parameter(Mandatory=$true)][string]
        $ResourceGroup,
        [Parameter(Mandatory=$true)][string]
        $Location,
        [Parameter(Mandatory=$true)][string]
        $VmType,
        [Parameter(Mandatory=$true)][string]
        $SubnetName,
        [Parameter(Mandatory=$true)][string]
        $SecurityGroupName,
        [Parameter(Mandatory=$true)][string]
        $VNetName,
        [Parameter(Mandatory=$true)][string]
        $RouteTableName,
        [Parameter(Mandatory=$false)][string] # Need one of these configured
        $PrimaryAvailabilitySetName,
        [Parameter(Mandatory=$false)][string] # Need one of these configured
        $PrimaryScaleSetName,
        [Parameter(Mandatory=$true)][string]
        $UseManagedIdentityExtension,
        [string]
        $UserAssignedClientID,
        [Parameter(Mandatory=$true)][string]
        $UseInstanceMetadata,
        [Parameter(Mandatory=$true)][string]
        $LoadBalancerSku,
        [Parameter(Mandatory=$true)][string]
        $ExcludeMasterFromStandardLB,
        [Parameter(Mandatory=$true)][string]
        $KubeDir
    )

    if ( -Not $PrimaryAvailabilitySetName -And -Not $PrimaryScaleSetName )
    {
        throw "Either PrimaryAvailabilitySetName or PrimaryScaleSetName must be set"
    }

    $azureConfigFile = [io.path]::Combine($KubeDir, "azure.json")

    $azureConfig = @"
{
    "tenantId": "$TenantId",
    "subscriptionId": "$SubscriptionId",
    "aadClientId": "$AADClientId",
    "aadClientSecret": "$AADClientSecret",
    "resourceGroup": "$ResourceGroup",
    "location": "$Location",
    "vmType": "$VmType",
    "subnetName": "$SubnetName",
    "securityGroupName": "$SecurityGroupName",
    "vnetName": "$VNetName",
    "routeTableName": "$RouteTableName",
    "primaryAvailabilitySetName": "$PrimaryAvailabilitySetName",
    "primaryScaleSetName": "$PrimaryScaleSetName",
    "useManagedIdentityExtension": $UseManagedIdentityExtension,
    "userAssignedIdentityID": $UserAssignedClientID,
    "useInstanceMetadata": $UseInstanceMetadata,
    "loadBalancerSku": "$LoadBalancerSku",
    "excludeMasterFromStandardLB": $ExcludeMasterFromStandardLB
}
"@

    $azureConfig | Out-File -encoding ASCII -filepath "$azureConfigFile"
}


function
Write-CACert
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $CACertificate,
        [Parameter(Mandatory=$true)][string]
        $KubeDir
    )
    $caFile = [io.path]::Combine($KubeDir, "ca.crt")
    [System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($CACertificate)) | Out-File -Encoding ascii $caFile
}

function
Write-KubeConfig
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $CACertificate,
        [Parameter(Mandatory=$true)][string]
        $MasterFQDNPrefix,
        [Parameter(Mandatory=$true)][string]
        $MasterIP,
        [Parameter(Mandatory=$true)][string]
        $AgentKey,
        [Parameter(Mandatory=$true)][string]
        $AgentCertificate,
        [Parameter(Mandatory=$true)][string]
        $KubeDir
    )
    $kubeConfigFile = [io.path]::Combine($KubeDir, "config")

    $kubeConfig = @"
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: "$CACertificate"
    server: https://${MasterIP}:443
  name: "$MasterFQDNPrefix"
contexts:
- context:
    cluster: "$MasterFQDNPrefix"
    user: "$MasterFQDNPrefix-admin"
  name: "$MasterFQDNPrefix"
current-context: "$MasterFQDNPrefix"
kind: Config
users:
- name: "$MasterFQDNPrefix-admin"
  user:
    client-certificate-data: "$AgentCertificate"
    client-key-data: "$AgentKey"
"@

    $kubeConfig | Out-File -encoding ASCII -filepath "$kubeConfigFile"
}

function
New-InfraContainer
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $KubeDir
    )
    cd $KubeDir
    $computerInfo = Get-ComputerInfo
    $windowsBase = if ($computerInfo.WindowsVersion -eq "1709") {
        "microsoft/nanoserver:1709"
    } elseif ($computerInfo.WindowsVersion -eq "1803") {
        "microsoft/nanoserver:1803"
    } elseif ($computerInfo.WindowsVersion -eq "1809") {
        # TODO: unsure if 2019 will report 1809 or not
        "microsoft/nanoserver:1809"
    } else {
        "mcr.microsoft.com/nanoserver-insider"
    }

    "FROM $($windowsBase)" | Out-File -encoding ascii -FilePath Dockerfile
    "CMD cmd /c ping -t localhost" | Out-File -encoding ascii -FilePath Dockerfile -Append
    docker build -t kubletwin/pause .
}


# TODO: Deprecate this and replace with methods that get individual components instead of zip containing everything
# This expects the ZIP file to be created by scripts/build-windows-k8s.sh
function
Get-KubePackage
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $KubeBinariesSASURL
    )
    
    $zipfile = "c:\k.zip"
    for ($i=0; $i -le 10; $i++)
    {
        DownloadFileOverHttp -Url $KubeBinariesSASURL -DestinationPath $zipfile
        if ($?) {
            break
        } else {
            Write-Log $Error[0].Exception.Message
        }
    }
    Expand-Archive -path $zipfile -DestinationPath C:\
}

function
Get-KubeBinaries
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $KubeBinariesURL
    )

    if ($computerInfo.WindowsVersion -eq "1709") 
    {        
        Write-Log "Server version 1709 does not support using kubernetes binaries in tar file."    
        return
    }
    
    $tempdir = New-TemporaryDirectory
    $binaryPackage = "$tempdir\k.tar.gz"
    for ($i=0; $i -le 10; $i++)
    {
        DownloadFileOverHttp -Url $KubeBinariesURL -DestinationPath $binaryPackage
        if ($?) {
            break
        } else {
            Write-Log $Error[0].Exception.Message
        }
    }

    # using tar to minimize dependencies    
    # tar should be avalible on 1803+
    tar -xzf $binaryPackage -C $tempdir
     
    # copy binaries over to kube folder      
    $windowsbinariespath = "c:\k\"
    if(!(Test-path $windowsbinariespath)) {
        mkdir $windowsbinariespath
    }   
    cp $tempdir\kubernetes\node\bin\* $windowsbinariespath -Recurse
    
    #remove temp folder created when unzipping            
    del $tempdir -Recurse
}

# TODO: replace KubeletStartFile with a Kubelet config, remove NSSM, and use built-in service integration
function
New-NSSMService
{
    Param(
        [string]
        [Parameter(Mandatory=$true)]
        $KubeDir,
        [string]
        [Parameter(Mandatory=$true)]
        $KubeletStartFile,
        [string]
        [Parameter(Mandatory=$true)]
        $KubeProxyStartFile
    )

    # setup kubelet
    & "$KubeDir\nssm.exe" install Kubelet C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe
    & "$KubeDir\nssm.exe" set Kubelet AppDirectory $KubeDir
    & "$KubeDir\nssm.exe" set Kubelet AppParameters $KubeletStartFile
    & "$KubeDir\nssm.exe" set Kubelet DisplayName Kubelet
    & "$KubeDir\nssm.exe" set Kubelet Description Kubelet
    & "$KubeDir\nssm.exe" set Kubelet Start SERVICE_AUTO_START
    & "$KubeDir\nssm.exe" set Kubelet ObjectName LocalSystem
    & "$KubeDir\nssm.exe" set Kubelet Type SERVICE_WIN32_OWN_PROCESS
    & "$KubeDir\nssm.exe" set Kubelet AppThrottle 1500
    & "$KubeDir\nssm.exe" set Kubelet AppStdout C:\k\kubelet.log
    & "$KubeDir\nssm.exe" set Kubelet AppStderr C:\k\kubelet.err.log
    & "$KubeDir\nssm.exe" set Kubelet AppStdoutCreationDisposition 4
    & "$KubeDir\nssm.exe" set Kubelet AppStderrCreationDisposition 4
    & "$KubeDir\nssm.exe" set Kubelet AppRotateFiles 1
    & "$KubeDir\nssm.exe" set Kubelet AppRotateOnline 1
    & "$KubeDir\nssm.exe" set Kubelet AppRotateSeconds 86400
    & "$KubeDir\nssm.exe" set Kubelet AppRotateBytes 1048576

    # setup kubeproxy
    & "$KubeDir\nssm.exe" install Kubeproxy C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe
    & "$KubeDir\nssm.exe" set Kubeproxy AppDirectory $KubeDir
    & "$KubeDir\nssm.exe" set Kubeproxy AppParameters $KubeProxyStartFile
    & "$KubeDir\nssm.exe" set Kubeproxy DisplayName Kubeproxy
    & "$KubeDir\nssm.exe" set Kubeproxy DependOnService Kubelet
    & "$KubeDir\nssm.exe" set Kubeproxy Description Kubeproxy
    & "$KubeDir\nssm.exe" set Kubeproxy Start SERVICE_AUTO_START
    & "$KubeDir\nssm.exe" set Kubeproxy ObjectName LocalSystem
    & "$KubeDir\nssm.exe" set Kubeproxy Type SERVICE_WIN32_OWN_PROCESS
    & "$KubeDir\nssm.exe" set Kubeproxy AppThrottle 1500
    & "$KubeDir\nssm.exe" set Kubeproxy AppStdout C:\k\kubeproxy.log
    & "$KubeDir\nssm.exe" set Kubeproxy AppStderr C:\k\kubeproxy.err.log
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateFiles 1
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateOnline 1
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateSeconds 86400
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateBytes 1048576
}

# Renamed from Write-KubernetesStartFiles
function
Install-KubernetesServices
{
    param(
        [Parameter(Mandatory=$true)][string[]]
        $KubeletConfigArgs,
        [Parameter(Mandatory=$true)][string]
        $KubeBinariesVersion,
        [Parameter(Mandatory=$true)][string]
        $NetworkPlugin,
        [Parameter(Mandatory=$true)][string]
        $NetworkMode,
        [Parameter(Mandatory=$true)][string]
        $KubeDir,
        [Parameter(Mandatory=$true)][string]
        $AzureCNIBinDir,
        [Parameter(Mandatory=$true)][string]
        $AzureCNIConfDir,
        [Parameter(Mandatory=$true)][string]
        $CNIPath,
        [Parameter(Mandatory=$true)][string]
        $CNIConfig,
        [Parameter(Mandatory=$true)][string]
        $CNIConfigPath,
        [Parameter(Mandatory=$true)][string]
        $MasterIP,
        [Parameter(Mandatory=$true)][string]
        $KubeDnsServiceIp,
        [Parameter(Mandatory=$true)][string]
        $MasterSubnet,
        [Parameter(Mandatory=$true)][string]
        $KubeClusterCIDR,
        [Parameter(Mandatory=$true)][string]
        $KubeServiceCIDR,
        [Parameter(Mandatory=$true)][string]
        $HNSModule,
        [Parameter(Mandatory=$true)][string]
        $KubeletNodeLabels
    )

    # Calculate some local paths
    $VolumePluginDir = [Io.path]::Combine($KubeDir, "volumeplugins")
    $KubeletStartFile = [io.path]::Combine($KubeDir, "kubeletstart.ps1")
    $KubeProxyStartFile = [io.path]::Combine($KubeDir, "kubeproxystart.ps1")

    mkdir $VolumePluginDir
    $KubeletArgList = $KubeletConfigArgs # This is the initial list passed in from aks-engine
    $KubeletArgList += "--node-labels=`$global:KubeletNodeLabels"
    # $KubeletArgList += "--hostname-override=`$global:AzureHostname" TODO: remove - dead code?
    $KubeletArgList += "--volume-plugin-dir=`$global:VolumePluginDir"
    # If you are thinking about adding another arg here, you should be considering pkg/engine/defaults-kubelet.go first
    # Only args that need to be calculated or combined with other ones on the Windows agent should be added here.
    

    # Regex to strip version to Major.Minor.Build format such that the following check does not crash for version like x.y.z-alpha
    [regex]$regex = "^[0-9.]+"
    $KubeBinariesVersionStripped = $regex.Matches($KubeBinariesVersion).Value
    if ([System.Version]$KubeBinariesVersionStripped -lt [System.Version]"1.8.0")
    {
        # --api-server deprecates from 1.8.0
        $KubeletArgList += "--api-servers=https://`${global:MasterIP}:443"
    }

    # Configure kubelet to use CNI plugins if enabled.
    if ($NetworkPlugin -eq "azure") {
        $KubeletArgList += @("--cni-bin-dir=$AzureCNIBinDir", "--cni-conf-dir=$AzureCNIConfDir")
    } elseif ($NetworkPlugin -eq "kubenet") {
        $KubeletArgList += @("--cni-bin-dir=$CNIPath", "--cni-conf-dir=$CNIConfigPath")
        # handle difference in naming between Linux & Windows reference plugin
        $KubeletArgList = $KubeletArgList -replace "kubenet", "cni"
    } else {
        throw "Unknown network type $NetworkPlugin, can't configure kubelet"
    }

    # Used in WinCNI version of kubeletstart.ps1
    $KubeletArgListStr = ""
    $KubeletArgList | Foreach-Object {
        # Since generating new code to be written to a file, need to escape quotes again
        if ($KubeletArgListStr.length -gt 0)
        {
            $KubeletArgListStr = $KubeletArgListStr + ", "
        }
        $KubeletArgListStr = $KubeletArgListStr + "`"" + $_.Replace("`"`"","`"`"`"`"") + "`""
    }
    $KubeletArgListStr = "@`($KubeletArgListStr`)"

    # Used in Azure-CNI version of kubeletstart.ps1
    $KubeletCommandLine = "$KubeDir\kubelet.exe " + ($KubeletArgList -join " ")

    $kubeStartStr = @"
`$global:MasterIP = "$MasterIP"
`$global:KubeDnsSearchPath = "svc.cluster.local"
`$global:KubeDnsServiceIp = "$KubeDnsServiceIp"
`$global:MasterSubnet = "$MasterSubnet"
`$global:KubeClusterCIDR = "$KubeClusterCIDR"
`$global:KubeServiceCIDR = "$KubeServiceCIDR"
`$global:KubeBinariesVersion = "$KubeBinariesVersion"
`$global:CNIPath = "$CNIPath"
`$global:NetworkMode = "$NetworkMode"
`$global:ExternalNetwork = "ext"
`$global:CNIConfig = "$CNIConfig"
`$global:HNSModule = "$HNSModule"
`$global:VolumePluginDir = "$VolumePluginDir"
`$global:NetworkPlugin="$NetworkPlugin"
`$global:KubeletNodeLabels="$KubeletNodeLabels"

"@

    if ($NetworkPlugin -eq "azure") {
        $KubeNetwork = "azure"
        $kubeStartStr += @"
Write-Host "NetworkPlugin azure, starting kubelet."

# Turn off Firewall to enable pods to talk to service endpoints. (Kubelet should eventually do this)
netsh advfirewall set allprofiles state off
# startup the service

# Find if the primary external switch network exists. If not create one.
# This is done only once in the lifetime of the node
`$hnsNetwork = Get-HnsNetwork | ? Name -EQ `$global:ExternalNetwork
if (!`$hnsNetwork)
{
    Write-Host "Creating a new hns Network"
    ipmo `$global:HNSModule
    # Fixme : use a smallest range possible, that will not collide with any pod space
    New-HNSNetwork -Type `$global:NetworkMode -AddressPrefix "192.168.255.0/30" -Gateway "192.168.255.1" -Name `$global:ExternalNetwork -Verbose
}

# Find if network created by CNI exists, if yes, remove it
# This is required to keep the network non-persistent behavior
# Going forward, this would be done by HNS automatically during restart of the node

`$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
if (`$hnsNetwork)
{
    # Cleanup all containers
    docker ps -q | foreach {docker rm `$_ -f}

    Write-Host "Cleaning up old HNS network found"
    Remove-HnsNetwork `$hnsNetwork
    # Kill all cni instances & stale data left by cni
    # Cleanup all files related to cni
    `$cnijson = [io.path]::Combine("$KubeDir", "azure-vnet-ipam.json")
    if ((Test-Path `$cnijson))
    {
        Remove-Item `$cnijson
    }
    `$cnilock = [io.path]::Combine("$KubeDir", "azure-vnet-ipam.lock")
    if ((Test-Path `$cnilock))
    {
        Remove-Item `$cnilock
    }
    taskkill /IM azure-vnet-ipam.exe /f

    `$cnijson = [io.path]::Combine("$KubeDir", "azure-vnet.json")
    if ((Test-Path `$cnijson))
    {
        Remove-Item `$cnijson
    }
    `$cnilock = [io.path]::Combine("$KubeDir", "azure-vnet.lock")
    if ((Test-Path `$cnilock))
    {
        Remove-Item `$cnilock
    }
    taskkill /IM azure-vnet.exe /f
}

# Restart Kubeproxy, which would wait, until the network is created
Restart-Service Kubeproxy

$KubeletCommandLine

"@
    } 
    else  # using WinCNI. TODO: If WinCNI support is removed, then delete this as dead code later
    {
        $KubeNetwork = "l2bridge"
        $kubeStartStr += @"

function
Get-DefaultGateway(`$CIDR)
{
    return `$CIDR.substring(0,`$CIDR.lastIndexOf(".")) + ".1"
}

function
Get-PodCIDR()
{
    `$podCIDR = c:\k\kubectl.exe --kubeconfig=c:\k\config get nodes/`$(`$env:computername.ToLower()) -o custom-columns=podCidr:.spec.podCIDR --no-headers
    return `$podCIDR
}

function
Test-PodCIDR(`$podCIDR)
{
    return `$podCIDR.length -gt 0
}

function
Update-CNIConfig(`$podCIDR, `$masterSubnetGW)
{
    `$jsonSampleConfig =
"{
    ""cniVersion"": ""0.2.0"",
    ""name"": ""<NetworkMode>"",
    ""type"": ""wincni.exe"",
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
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""OutBoundNAT"", ""ExceptionList"": [ ""<ClusterCIDR>"", ""<MgmtSubnet>"" ] }
    },
    {
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""ROUTE"", ""DestinationPrefix"": ""<ServiceCIDR>"", ""NeedEncap"" : true }
    }
    ]
}"

    `$configJson = ConvertFrom-Json `$jsonSampleConfig
    `$configJson.name = `$global:NetworkMode.ToLower()
    `$configJson.ipam.subnet=`$podCIDR
    `$configJson.ipam.routes[0].GW = `$masterSubnetGW
    `$configJson.dns.Nameservers[0] = `$global:KubeDnsServiceIp
    `$configJson.dns.Search[0] = `$global:KubeDnsSearchPath

    `$configJson.AdditionalArgs[0].Value.ExceptionList[0] = `$global:KubeClusterCIDR
    `$configJson.AdditionalArgs[0].Value.ExceptionList[1] = `$global:MasterSubnet
    `$configJson.AdditionalArgs[1].Value.DestinationPrefix  = `$global:KubeServiceCIDR

    if (Test-Path `$global:CNIConfig)
    {
        Clear-Content -Path `$global:CNIConfig
    }

    Write-Host "Generated CNI Config [`$configJson]"

    Add-Content -Path `$global:CNIConfig -Value (ConvertTo-Json `$configJson -Depth 20)
}

try
{
    `$masterSubnetGW = Get-DefaultGateway `$global:MasterSubnet
    `$podCIDR=Get-PodCIDR
    `$podCidrDiscovered=Test-PodCIDR(`$podCIDR)

    # if the podCIDR has not yet been assigned to this node, start the kubelet process to get the podCIDR, and then promptly kill it.
    if (-not `$podCidrDiscovered)
    {
        `$argList = $KubeletArgListStr

        `$process = Start-Process -FilePath c:\k\kubelet.exe -PassThru -ArgumentList `$argList

        # run kubelet until podCidr is discovered
        Write-Host "waiting to discover pod CIDR"
        while (-not `$podCidrDiscovered)
        {
            Write-Host "Sleeping for 10s, and then waiting to discover pod CIDR"
            Start-Sleep 10

            `$podCIDR=Get-PodCIDR
            `$podCidrDiscovered=Test-PodCIDR(`$podCIDR)
        }

        # stop the kubelet process now that we have our CIDR, discard the process output
        `$process | Stop-Process | Out-Null
    }

    # Turn off Firewall to enable pods to talk to service endpoints. (Kubelet should eventually do this)
    netsh advfirewall set allprofiles state off

    # startup the service
    `$hnsNetwork = Get-HnsNetwork | ? Name -EQ `$global:NetworkMode.ToLower()

    if (`$hnsNetwork)
    {
        # Kubelet has been restarted with existing network.
        # Cleanup all containers
        docker ps -q | foreach {docker rm `$_ -f}
        # cleanup network
        Write-Host "Cleaning up old HNS network found"
        Remove-HnsNetwork `$hnsNetwork
        Start-Sleep 10
    }

    Write-Host "Creating a new hns Network"
    ipmo `$global:HNSModule

    `$hnsNetwork = New-HNSNetwork -Type `$global:NetworkMode -AddressPrefix `$podCIDR -Gateway `$masterSubnetGW -Name `$global:NetworkMode.ToLower() -Verbose
    # New network has been created, Kubeproxy service has to be restarted
    Restart-Service Kubeproxy

    Start-Sleep 10
    # Add route to all other POD networks
    Update-CNIConfig `$podCIDR `$masterSubnetGW

    $KubeletCommandLine
}
catch
{
    Write-Error `$_
}

"@
    } # end else using WinCNI.

    # Now that the script is generated, based on what CNI plugin and startup options are needed, write it to disk
    $kubeStartStr | Out-File -encoding ASCII -filepath $KubeletStartFile

    $kubeProxyStartStr = @"
`$env:KUBE_NETWORK = "$KubeNetwork"
`$global:NetworkMode = "$NetworkMode"
`$global:HNSModule = "$HNSModule"
`$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
while (!`$hnsNetwork)
{
    Write-Host "Waiting for Network [$KubeNetwork] to be created . . ."
    Start-Sleep 10
    `$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
}

#
# cleanup the persisted policy lists
#
ipmo `$global:HNSModule
Get-HnsPolicyList | Remove-HnsPolicyList

$KubeDir\kube-proxy.exe --v=3 --proxy-mode=kernelspace --hostname-override=$env:computername --kubeconfig=$KubeDir\config
"@

    $kubeProxyStartStr | Out-File -encoding ASCII -filepath $KubeProxyStartFile

    New-NSSMService -KubeDir $KubeDir `
                    -KubeletStartFile $KubeletStartFile `
                    -KubeProxyStartFile $KubeProxyStartFile
}