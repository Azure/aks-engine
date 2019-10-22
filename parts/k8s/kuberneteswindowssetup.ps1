<#
    .SYNOPSIS
        Provisions VM as a Kubernetes agent.

    .DESCRIPTION
        Provisions VM as a Kubernetes agent.

        The parameters passed in are required, and will vary per-deployment.

        Notes on modifying this file:
        - This file extension is PS1, but it is actually used as a template from pkg/engine/template_generator.go
        - All of the lines that have braces in them will be modified. Please do not change them here, change them in the Go sources
        - Single quotes are forbidden, they are reserved to delineate the different members for the ARM template concat() call
#>
[CmdletBinding(DefaultParameterSetName="Standard")]
param(
    [string]
    [ValidateNotNullOrEmpty()]
    $MasterIP,

    [parameter()]
    [ValidateNotNullOrEmpty()]
    $KubeDnsServiceIp,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $MasterFQDNPrefix,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $Location,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $AgentKey,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $AADClientId,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $AADClientSecret, # base64

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $NetworkAPIVersion,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $TargetEnvironment
)



# These globals will not change between nodes in the same cluster, so they are not
# passed as powershell parameters

## SSH public keys to add to authorized_keys
$global:SSHKeys = @( {{ GetSshPublicKeysPowerShell }} )

## Certificates generated by aks-engine
$global:CACertificate = "{{WrapAsParameter "caCertificate"}}"
$global:AgentCertificate = "{{WrapAsParameter "clientCertificate"}}"

## Download sources provided by aks-engine
$global:KubeBinariesPackageSASURL = "{{WrapAsParameter "kubeBinariesSASURL"}}"
$global:WindowsKubeBinariesURL = "{{WrapAsParameter "windowsKubeBinariesURL"}}"
$global:KubeBinariesVersion = "{{WrapAsParameter "kubeBinariesVersion"}}"

## Docker Version
$global:DockerVersion = "{{WrapAsParameter "windowsDockerVersion"}}"

## VM configuration passed by Azure
$global:WindowsTelemetryGUID = "{{WrapAsParameter "windowsTelemetryGUID"}}"
{{if eq GetIdentitySystem "adfs"}}
$global:TenantId = "adfs"
{{else}}
$global:TenantId = "{{WrapAsVariable "tenantID"}}"
{{end}}
$global:SubscriptionId = "{{WrapAsVariable "subscriptionId"}}"
$global:ResourceGroup = "{{WrapAsVariable "resourceGroup"}}"
$global:VmType = "{{WrapAsVariable "vmType"}}"
$global:SubnetName = "{{WrapAsVariable "subnetName"}}"
$global:MasterSubnet = "{{GetWindowsMasterSubnetARMParam}}"
$global:SecurityGroupName = "{{WrapAsVariable "nsgName"}}"
$global:VNetName = "{{WrapAsVariable "virtualNetworkName"}}"
$global:RouteTableName = "{{WrapAsVariable "routeTableName"}}"
$global:PrimaryAvailabilitySetName = "{{WrapAsVariable "primaryAvailabilitySetName"}}"
$global:PrimaryScaleSetName = "{{WrapAsVariable "primaryScaleSetName"}}"

$global:KubeClusterCIDR = "{{WrapAsParameter "kubeClusterCidr"}}"
$global:KubeServiceCIDR = "{{WrapAsParameter "kubeServiceCidr"}}"
$global:VNetCIDR = "{{WrapAsParameter "vnetCidr"}}"
{{if IsKubernetesVersionGe "1.16.0"}}
$global:KubeletNodeLabels = "{{GetAgentKubernetesLabels . "',variables('labelResourceGroup'),'"}}"
{{else}}
$global:KubeletNodeLabels = "{{GetAgentKubernetesLabelsDeprecated . "',variables('labelResourceGroup'),'"}}"
{{end}}
$global:KubeletConfigArgs = @( {{GetKubeletConfigKeyValsPsh .KubernetesConfig }} )

$global:UseManagedIdentityExtension = "{{WrapAsVariable "useManagedIdentityExtension"}}"
$global:UserAssignedClientID = "{{WrapAsVariable "userAssignedClientID"}}"
$global:UseInstanceMetadata = "{{WrapAsVariable "useInstanceMetadata"}}"

$global:LoadBalancerSku = "{{WrapAsVariable "loadBalancerSku"}}"
$global:ExcludeMasterFromStandardLB = "{{WrapAsVariable "excludeMasterFromStandardLB"}}"


# Windows defaults, not changed by aks-engine
$global:CacheDir = "c:\akse-cache"
$global:KubeDir = "c:\k"
$global:HNSModule = [Io.path]::Combine("$global:KubeDir", "hns.psm1")

$global:KubeDnsSearchPath = "svc.cluster.local"

$global:CNIPath = [Io.path]::Combine("$global:KubeDir", "cni")
$global:NetworkMode = "L2Bridge"
$global:CNIConfig = [Io.path]::Combine($global:CNIPath, "config", "`$global:NetworkMode.conf")
$global:CNIConfigPath = [Io.path]::Combine("$global:CNIPath", "config")


$global:AzureCNIDir = [Io.path]::Combine("$global:KubeDir", "azurecni")
$global:AzureCNIBinDir = [Io.path]::Combine("$global:AzureCNIDir", "bin")
$global:AzureCNIConfDir = [Io.path]::Combine("$global:AzureCNIDir", "netconf")

# Azure cni configuration
# $global:NetworkPolicy = "{{WrapAsParameter "networkPolicy"}}" # BUG: unused
$global:NetworkPlugin = "{{WrapAsParameter "networkPlugin"}}"
$global:VNetCNIPluginsURL = "{{WrapAsParameter "vnetCniWindowsPluginsURL"}}"

# Base64 representation of ZIP archive
$zippedFiles = "{{ GetKubernetesWindowsAgentFunctions }}"

# Extract ZIP from script
[io.file]::WriteAllBytes("scripts.zip", [System.Convert]::FromBase64String($zippedFiles))
Expand-Archive scripts.zip -DestinationPath "C:\\AzureData\\"

# Dot-source contents of zip. This should match the list in template_generator.go GetKubernetesWindowsAgentFunctions
. c:\AzureData\k8s\kuberneteswindowsfunctions.ps1
. c:\AzureData\k8s\windowsconfigfunc.ps1
. c:\AzureData\k8s\windowskubeletfunc.ps1
. c:\AzureData\k8s\windowscnifunc.ps1
. c:\AzureData\k8s\windowsazurecnifunc.ps1
. c:\AzureData\k8s\windowsinstallopensshfunc.ps1

function
Update-ServiceFailureActions()
{
    sc.exe failure "kubelet" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure "kubeproxy" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure "docker" actions= restart/60000/restart/60000/restart/60000 reset= 900
}

try
{
    # Set to false for debugging.  This will output the start script to
    # c:\AzureData\CustomDataSetupScript.log, and then you can RDP
    # to the windows machine, and run the script manually to watch
    # the output.
    if ($true) {
        Write-Log "Provisioning $global:DockerServiceName... with IP $MasterIP"

        # Install OpenSSH if SSH enabled
        $sshEnabled = [System.Convert]::ToBoolean("{{ WindowsSSHEnabled }}")

        if ( $sshEnabled ) {
            Install-OpenSSH -SSHKeys $SSHKeys
        }

        Write-Log "Apply telemetry data setting"
        Set-TelemetrySetting -WindowsTelemetryGUID $global:WindowsTelemetryGUID

        Write-Log "Resize os drive if possible"
        Resize-OSDrive

        Write-Log "Initialize data disks"
        Initialize-DataDisks

        Write-Log "Create required data directories as needed"
        Initialize-DataDirectories

        Write-Log "Install docker"
        Install-Docker -DockerVersion $global:DockerVersion

        Write-Log "Download kubelet binaries and unzip"
        Get-KubePackage -KubeBinariesSASURL $global:KubeBinariesPackageSASURL

        # this overwrite the binaries that are download from the custom packge with binaries
        # The custom package has a few files that are nessary for future steps (nssm.exe)
        # this is a temporary work around to get the binaries until we depreciate
        # custom package and nssm.exe as defined in #3851.
        if ($global:WindowsKubeBinariesURL){
            Write-Log "Overwriting kube node binaries from $global:WindowsKubeBinariesURL"
            Get-KubeBinaries -KubeBinariesURL $global:WindowsKubeBinariesURL
        }

        Write-Log "Write Azure cloud provider config"
        Write-AzureConfig `
            -KubeDir $global:KubeDir `
            -AADClientId $AADClientId `
            -AADClientSecret $([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($AADClientSecret))) `
            -TenantId $global:TenantId `
            -SubscriptionId $global:SubscriptionId `
            -ResourceGroup $global:ResourceGroup `
            -Location $Location `
            -VmType $global:VmType `
            -SubnetName $global:SubnetName `
            -SecurityGroupName $global:SecurityGroupName `
            -VNetName $global:VNetName `
            -RouteTableName $global:RouteTableName `
            -PrimaryAvailabilitySetName $global:PrimaryAvailabilitySetName `
            -PrimaryScaleSetName $global:PrimaryScaleSetName `
            -UseManagedIdentityExtension $global:UseManagedIdentityExtension `
            -UserAssignedClientID $global:UserAssignedClientID `
            -UseInstanceMetadata $global:UseInstanceMetadata `
            -LoadBalancerSku $global:LoadBalancerSku `
            -ExcludeMasterFromStandardLB $global:ExcludeMasterFromStandardLB `
            -TargetEnvironment $TargetEnvironment

        {{if IsAzureStackCloud}}
        $azureStackConfigFile = [io.path]::Combine($global:KubeDir, "azurestackcloud.json")
        $envJSON = "{{ GetBase64EncodedEnvironmentJSON }}"
        [io.file]::WriteAllBytes($azureStackConfigFile, [System.Convert]::FromBase64String($envJSON))
        {{end}}

        Write-Log "Write ca root"
        Write-CACert -CACertificate $global:CACertificate `
                     -KubeDir $global:KubeDir

        Write-Log "Write kube config"
        Write-KubeConfig -CACertificate $global:CACertificate `
                         -KubeDir $global:KubeDir `
                         -MasterFQDNPrefix $MasterFQDNPrefix `
                         -MasterIP $MasterIP `
                         -AgentKey $AgentKey `
                         -AgentCertificate $global:AgentCertificate

        Write-Log "Create the Pause Container kubletwin/pause"
        New-InfraContainer -KubeDir $global:KubeDir

        Write-Log "Configuring networking with NetworkPlugin:$global:NetworkPlugin"

        Write-Log "Disable firewall to enable pods to talk to service endpoints"
        #TODO: Kubelet should eventually do this
        netsh advfirewall set allprofiles state off

        # Configure network policy.
        if ($global:NetworkPlugin -eq "azure") {
            Write-Log "Installing Azure VNet plugins"
            Install-VnetPlugins -AzureCNIConfDir $global:AzureCNIConfDir `
                                -AzureCNIBinDir $global:AzureCNIBinDir `
                                -VNetCNIPluginsURL $global:VNetCNIPluginsURL
            Set-AzureCNIConfig -AzureCNIConfDir $global:AzureCNIConfDir `
                               -KubeDnsSearchPath $global:KubeDnsSearchPath `
                               -KubeClusterCIDR $global:KubeClusterCIDR `
                               -MasterSubnet $global:MasterSubnet `
                               -KubeServiceCIDR $global:KubeServiceCIDR `
                               -VNetCIDR $global:VNetCIDR `
                               -TargetEnvironment $TargetEnvironment

            if ($TargetEnvironment -ieq "AzureStackCloud") {
                GenerateAzureStackCNIConfig `
                    -TenantId $global:TenantId `
                    -SubscriptionId $global:SubscriptionId `
                    -ResourceGroup $global:ResourceGroup `
                    -AADClientId $AADClientId `
                    -KubeDir $global:KubeDir `
                    -AADClientSecret $([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($AADClientSecret))) `
                    -NetworkAPIVersion $NetworkAPIVersion `
                    -AzureEnvironmentFilePath $([io.path]::Combine($global:KubeDir, "azurestackcloud.json")) `
                    -IdentitySystem "{{ GetIdentitySystem }}"
            }
        } elseif ($global:NetworkPlugin -eq "kubenet") {
            Write-Log "Fetching additional files needed for kubenet"
            Update-WinCNI -CNIPath $global:CNIPath

            Write-Log "Creating WinCNI network config file"
            Write-WinCNIConfig `
                -cniConfigPath "c:\k\cni\config\$($NetworkMode).conf" `
                -networkMode $global:NetworkMode `
                -kubeDnsServiceIp $KubeDnsServiceIp `
                -kubeDnsSearchPath "svc.cluster.local" `
                -kubeClusterCIDR $KubeClusterCIDR `
                -masterSubnet $global:MasterSubnet `
                -kubeServiceCIDR $global:KubeServiceCIDR
        }

        Write-Log "Creating ext HNS network"
        # This is done so network connectivity to the node is not lost when other networks are added/removed
        Get-HnsPsm1 -HNSModule $global:HNSModule
        Import-Module $global:HNSModule

        $netAdapters = @(Get-NetAdapter -Physical)
        if ($netAdapters.Count -eq 0) {
            throw "Failed to find any physical network adapters"
        }

        $managementIP = (Get-NetIPAddress -ifIndex $netAdapters[0].ifIndex -AddressFamily IPv4).IPAddress
        $adapterName = $netAdapters[0].Name

        Write-Log "Using adapter $adapterName with IP address $managementIP"

        # Fixme : use a smallest range possible, that will not collide with any pod space
        New-HNSNetwork -Type $global:NetworkMode -AddressPrefix "192.168.255.0/30" -Gateway "192.168.255.1" -AdapterName $adapterName -Name "ext" -Verbose

        # Wait for switch to be created and ip address to be assigned
        for ($i = 0; $i -lt 60; $i++) {
            $mgmtIPAfterNetworkCreate = Get-NetIPAddress $managementIP -ErrorAction SilentlyContinue
            if ($mgmtIPAfterNetworkCreate) {
                break
            }
            Write-Log "Waiting for IP address to be assigned..."
            Start-Sleep -Milliseconds 1000
        }

        if (-not $mgmtIPAfterNetworkCreate) {
            throw "Failed to find $managementIP after creating $($global:ExternalNetwork) network"
        }
        else {
            Write-Log "IP address has been assigned"
        }

        Write-Log "Write kubelet startfile with pod CIDR of $podCIDR"
        Install-KubernetesServices `
            -KubeletConfigArgs $global:KubeletConfigArgs `
            -KubeBinariesVersion $global:KubeBinariesVersion `
            -NetworkPlugin $global:NetworkPlugin `
            -NetworkMode $global:NetworkMode `
            -KubeDir $global:KubeDir `
            -AzureCNIBinDir $global:AzureCNIBinDir `
            -AzureCNIConfDir $global:AzureCNIConfDir `
            -CNIPath $global:CNIPath `
            -CNIConfig $global:CNIConfig `
            -CNIConfigPath $global:CNIConfigPath `
            -MasterIP $MasterIP `
            -KubeDnsServiceIp $KubeDnsServiceIp `
            -MasterSubnet $global:MasterSubnet `
            -KubeClusterCIDR $global:KubeClusterCIDR `
            -KubeServiceCIDR $global:KubeServiceCIDR `
            -HNSModule $global:HNSModule `
            -KubeletNodeLabels $global:KubeletNodeLabels

        Get-NetworkLogCollectionScripts

        Write-Log "Disable Internet Explorer compat mode and set homepage"
        Set-Explorer

        Write-Log "Adjust pagefile size"
        Adjust-PageFileSize

        Write-Log "Start preProvisioning script"
        PREPROVISION_EXTENSION

        Write-Log "Update service failure actions"
        Update-ServiceFailureActions

        if (Test-Path $CacheDir)
        {
            Write-Log "Removing aks-engine bits cache directory"
            Remove-Item $CacheDir -Recurse -Force
        }

        Register-NodeCleanupScriptTask

        Write-Log "Setup Complete, reboot computer"
        Restart-Computer
    }
    else
    {
        # keep for debugging purposes
        Write-Log ".\CustomDataSetupScript.ps1 -MasterIP $MasterIP -KubeDnsServiceIp $KubeDnsServiceIp -MasterFQDNPrefix $MasterFQDNPrefix -Location $Location -AgentKey $AgentKey -AADClientId $AADClientId -AADClientSecret $AADClientSecret"
    }
}
catch
{
    Write-Error $_
    rruurnrn 1
}
