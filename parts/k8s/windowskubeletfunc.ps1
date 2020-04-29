function
Write-AzureConfig {
    Param(

        [Parameter(Mandatory = $true)][string]
        $AADClientId,
        [Parameter(Mandatory = $true)][string]
        $AADClientSecret,
        [Parameter(Mandatory = $true)][string]
        $TenantId,
        [Parameter(Mandatory = $true)][string]
        $SubscriptionId,
        [Parameter(Mandatory = $true)][string]
        $ResourceGroup,
        [Parameter(Mandatory = $true)][string]
        $Location,
        [Parameter(Mandatory = $true)][string]
        $VmType,
        [Parameter(Mandatory = $true)][string]
        $SubnetName,
        [Parameter(Mandatory = $true)][string]
        $SecurityGroupName,
        [Parameter(Mandatory = $true)][string]
        $VNetName,
        [Parameter(Mandatory = $true)][string]
        $RouteTableName,
        [Parameter(Mandatory = $false)][string] # Need one of these configured
        $PrimaryAvailabilitySetName,
        [Parameter(Mandatory = $false)][string] # Need one of these configured
        $PrimaryScaleSetName,
        [Parameter(Mandatory = $true)][string]
        $UseManagedIdentityExtension,
        [string]
        $UserAssignedClientID,
        [Parameter(Mandatory = $true)][string]
        $UseInstanceMetadata,
        [Parameter(Mandatory = $true)][string]
        $LoadBalancerSku,
        [Parameter(Mandatory = $true)][string]
        $ExcludeMasterFromStandardLB,
        [Parameter(Mandatory = $true)][string]
        $KubeDir,
        [Parameter(Mandatory = $true)][string]
        $TargetEnvironment,
        [Parameter(Mandatory = $false)][bool]
        $UseContainerD = $false
    )

    if ( -Not $PrimaryAvailabilitySetName -And -Not $PrimaryScaleSetName ) {
        throw "Either PrimaryAvailabilitySetName or PrimaryScaleSetName must be set"
    }

    $azureConfigFile = [io.path]::Combine($KubeDir, "azure.json")

    $azureConfig = @"
{
    "cloud": "$TargetEnvironment",
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
    "userAssignedIdentityID": "$UserAssignedClientID",
    "useInstanceMetadata": $UseInstanceMetadata,
    "loadBalancerSku": "$LoadBalancerSku",
    "excludeMasterFromStandardLB": $ExcludeMasterFromStandardLB
}
"@

    $azureConfig | Out-File -encoding ASCII -filepath "$azureConfigFile"
}


function
Write-CACert {
    Param(
        [Parameter(Mandatory = $true)][string]
        $CACertificate,
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )
    $caFile = [io.path]::Combine($KubeDir, "ca.crt")
    [System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($CACertificate)) | Out-File -Encoding ascii $caFile
}

function
Write-KubeConfig {
    Param(
        [Parameter(Mandatory = $true)][string]
        $CACertificate,
        [Parameter(Mandatory = $true)][string]
        $MasterFQDNPrefix,
        [Parameter(Mandatory = $true)][string]
        $MasterIP,
        [Parameter(Mandatory = $true)][string]
        $AgentKey,
        [Parameter(Mandatory = $true)][string]
        $AgentCertificate,
        [Parameter(Mandatory = $true)][string]
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
Test-ContainerImageExists {
    Param(
        [Parameter(Mandatory = $true)][string]
        $Image,
        [Parameter(Mandatory = $false)][string]
        $Tag,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )

    $target = $Image
    if ($Tag) {
        $target += ":$Tag"
    }

    if ($ContainerRuntime -eq "docker") {
        $images = docker image list $target --format "{{json .}}"
        return $images.Count -gt 0
    }
    else {
        return ( (ctr.exe -n k8s.io images list) | Select-String $target) -ne $Null
    }
}

function
Build-PauseContainer {
    Param(
        [Parameter(Mandatory = $true)][string]
        $WindowsBase,
        $DestinationTag,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )
    # Future work: This needs to build wincat - see https://github.com/Azure/aks-engine/issues/1461
    # Otherwise, delete this code and require a prebuilt pause image (or override with one from an Azure Container Registry instance)
    # ContainerD can't build, so doing the builds outside of node deployment is probably the right long-term solution.
    "FROM $($WindowsBase)" | Out-File -encoding ascii -FilePath Dockerfile
    "CMD cmd /c ping -t localhost" | Out-File -encoding ascii -FilePath Dockerfile -Append
    if ($ContainerRuntime -eq "docker") {
        Invoke-Executable -Executable "docker" -ArgList @("build", "-t", "$DestinationTag", ".")
    }
    else {
        throw "Cannot build pause container without Docker"
    }
}

function
New-InfraContainer {
    Param(
        [Parameter(Mandatory = $true)][string]
        $KubeDir,
        $DestinationTag = "kubletwin/pause",
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )
    cd $KubeDir
    $windowsVersion = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").ReleaseId

    # Reference for these tags: curl -L https://mcr.microsoft.com/v2/k8s/core/pause/tags/list
    # Then docker run --rm mplatform/manifest-tool inspect mcr.microsoft.com/k8s/core/pause:<tag>

    $clusterConfig = ConvertFrom-Json ((Get-Content $global:KubeClusterConfigPath -ErrorAction Stop) | Out-String)
    $defaultPauseImage = $clusterConfig.Cri.Images.Pause

    $pauseImageVersions = @("1803", "1809", "1903", "1909")

    if ($pauseImageVersions -icontains $windowsVersion) {
        if ($ContainerRuntime -eq "docker") {
            if (-not (Test-ContainerImageExists -Image $defaultPauseImage -ContainerRuntime $ContainerRuntime)) {
                Invoke-Executable -Executable "docker" -ArgList @("pull", "$defaultPauseImage") -Retries 5 -RetryDelaySeconds 30
            }
            Invoke-Executable -Executable "docker" -ArgList @("tag", "$defaultPauseImage", "$DestinationTag")
        }
        else {
            # containerd
            if (-not (Test-ContainerImageExists -Image $defaultPauseImage -ContainerRuntime $ContainerRuntime)) {
                Invoke-Executable -Executable "ctr" -ArgList @("-n", "k8s.io", "image", "pull", "$defaultPauseImage") -Retries 5 -RetryDelaySeconds 30
            }
            Invoke-Executable -Executable "ctr" -ArgList @("-n", "k8s.io", "image", "tag", "$defaultPauseImage", "$DestinationTag")
        }
    }
    else {
        Build-PauseContainer -WindowsBase "mcr.microsoft.com/nanoserver-insider" -DestinationTag $DestinationTag -ContainerRuntime $ContainerRuntime
    }
}

function
Test-ContainerImageExists {
    Param(
        [Parameter(Mandatory = $true)][string]
        $Image,
        [Parameter(Mandatory = $false)][string]
        $Tag,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )

    $target = $Image
    if ($Tag) {
        $target += ":$Tag"
    }

    if ($ContainerRuntime -eq "docker") {
        $images = docker image list $target --format "{{json .}}"
        return $images.Count -gt 0
    }
    else {
        return ( (ctr.exe -n k8s.io images list) | Select-String $target) -ne $Null
    }
}

# TODO: Deprecate this and replace with methods that get individual components instead of zip containing everything
# This expects the ZIP file created by Azure Pipelines.
function
Get-KubePackage {
    Param(
        [Parameter(Mandatory = $true)][string]
        $KubeBinariesSASURL
    )

    $zipfile = "c:\k.zip"
    for ($i = 0; $i -le 10; $i++) {
        DownloadFileOverHttp -Url $KubeBinariesSASURL -DestinationPath $zipfile
        if ($?) {
            break
        }
        else {
            Write-Log $Error[0].Exception.Message
        }
    }
    Expand-Archive -path $zipfile -DestinationPath C:\
}

function
Get-KubeBinaries {
    Param(
        [Parameter(Mandatory = $true)][string]
        $KubeBinariesURL
    )

    $tempdir = New-TemporaryDirectory
    $binaryPackage = "$tempdir\k.tar.gz"
    for ($i = 0; $i -le 10; $i++) {
        DownloadFileOverHttp -Url $KubeBinariesURL -DestinationPath $binaryPackage
        if ($?) {
            break
        }
        else {
            Write-Log $Error[0].Exception.Message
        }
    }

    # using tar to minimize dependencies
    # tar should be avalible on 1803+
    tar -xzf $binaryPackage -C $tempdir

    # copy binaries over to kube folder
    $windowsbinariespath = "c:\k\"
    if (!(Test-path $windowsbinariespath)) {
        mkdir $windowsbinariespath
    }
    cp $tempdir\kubernetes\node\bin\* $windowsbinariespath -Recurse

    #remove temp folder created when unzipping
    del $tempdir -Recurse
}

# TODO: replace KubeletStartFile with a Kubelet config, remove NSSM, and use built-in service integration
function
New-NSSMService {
    Param(
        [string]
        [Parameter(Mandatory = $true)]
        $KubeDir,
        [string]
        [Parameter(Mandatory = $true)]
        $KubeletStartFile,
        [string]
        [Parameter(Mandatory = $true)]
        $KubeProxyStartFile
    )

    $kubeletDependOnServices = "docker"
    if ($global:EnableCsiProxy) {
        $kubeletDependOnServices += " csi-proxy"
    }

    # setup kubelet
    & "$KubeDir\nssm.exe" install Kubelet C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppDirectory $KubeDir | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppParameters $KubeletStartFile | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet DisplayName Kubelet | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRestartDelay 5000 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet Description Kubelet | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStdout C:\k\kubelet.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStderr C:\k\kubelet.err.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStdoutCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStderrCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateBytes 10485760 | RemoveNulls
    # Do not use & when calling DependOnService since 'docker csi-proxy'
    # is parsed as a single string instead of two separate strings
    Invoke-Expression "$KubeDir\nssm.exe set Kubelet DependOnService $kubeletDependOnServices | RemoveNulls"

    # setup kubeproxy
    & "$KubeDir\nssm.exe" install Kubeproxy C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppDirectory $KubeDir | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppParameters $KubeProxyStartFile | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy DisplayName Kubeproxy | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy DependOnService Kubelet | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy Description Kubeproxy | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppStdout C:\k\kubeproxy.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppStderr C:\k\kubeproxy.err.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateBytes 10485760 | RemoveNulls
}

# Renamed from Write-KubernetesStartFiles
function
Install-KubernetesServices {
    param(
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )

    # TODO ksbrmnn fix callers to this function

    $KubeletStartFile = [io.path]::Combine($KubeDir, "kubeletstart.ps1")
    $KubeProxyStartFile = [io.path]::Combine($KubeDir, "kubeproxystart.ps1")

    New-NSSMService -KubeDir $KubeDir `
        -KubeletStartFile $KubeletStartFile `
        -KubeProxyStartFile $KubeProxyStartFile
}
