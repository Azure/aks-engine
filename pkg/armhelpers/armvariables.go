package armhelpers

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	"strconv"
)

func getK8sMasterVars(cs *api.ContainerService) map[string]interface{} {
	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedID := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""
	userAssignedClientID := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedClientID != ""
	useInstanceMetadata := cs.Properties.OrchestratorProfile.KubernetesConfig.UseInstanceMetadata
	excludeMasterFromStandardLB := to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB)
	maxLoadBalancerCount := cs.Properties.OrchestratorProfile.KubernetesConfig.MaximumLoadBalancerRuleCount
	isHostedMaster := cs.Properties.IsHostedMasterProfile()
	isMasterVMSS := cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.IsVirtualMachineScaleSets()
	hasStorageAccountDisks := cs.Properties.HasStorageAccountDisks()
	isAzureCNI := cs.Properties.OrchestratorProfile.IsAzureCNI()
	isCustomVnet := cs.Properties.AreAgentProfilesCustomVNET()
	isPrivateCluster := to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled)
	provisionJumpbox := cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision()
	enableEncryptionWithExternalKms := to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms)

	masterVars := map[string]interface{}{
		"maxVMsPerPool":                 100,
		"useManagedIdentityExtension":   strconv.FormatBool(useManagedIdentity),
		"userAssignedID":                strconv.FormatBool(userAssignedID),
		"userAssignedClientID":          strconv.FormatBool(userAssignedClientID),
		"userAssignedIDReference":       "[resourceId('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))]",
		"useInstanceMetadata":           strconv.FormatBool(to.Bool(useInstanceMetadata)),
		"loadBalancerSku":               cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku,
		"excludeMasterFromStandardLB":   strconv.FormatBool(excludeMasterFromStandardLB),
		"maximumLoadBalancerRuleCount":  maxLoadBalancerCount,
		"masterFqdnPrefix":              "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
		"apiVersionCompute":             "2018-06-01",
		"apiVersionStorage":             "2018-07-01",
		"apiVersionKeyVault":            "2018-02-14",
		"apiVersionNetwork":             "2018-08-01",
		"apiVersionManagedIdentity":     "2015-08-31-preview",
		"apiVersionAuthorizationUser":   "2018-09-01-preview",
		"apiVersionAuthorizationSystem": "2018-01-01-preview",
		"locations": []string{
			"[resourceGroup().location]",
			"[parameters('location')]",
		},
		"location":                  "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"masterAvailabilitySet":     "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"resourceGroup":             "[resourceGroup().name]",
		"truncatedResourceGroup":    "[take(replace(replace(resourceGroup().name, '(', '-'), ')', '-'), 63)]",
		"labelResourceGroup":        "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"routeTableID":              "[resourceId('Microsoft.Network/routeTables', variables('routeTableName'))]",
		"sshNatPorts":               []int{22, 2201, 2202, 2203, 2204},
		"sshKeyPath":                "[concat('/home/',parameters('linuxAdminUsername'),'/.ssh/authorized_keys')]",
		"provisionScript":           engine.GetKubernetesB64Provision(),
		"provisionSource":           engine.GetKubernetesB64ProvisionSource(),
		"healthMonitorScript":       engine.GetKubernetesB64HealthMonitorScript(),
		"provisionInstalls":         engine.GetKubernetesB64Installs(),
		"provisionConfigs":          engine.GetKubernetesB64Configs(),
		"mountetcdScript":           engine.GetKubernetesB64Mountetcd(),
		"customSearchDomainsScript": engine.GetKubernetesB64CustomSearchDomainsScript(),
		"sshdConfig":                engine.GetB64sshdConfig(),
		"systemConf":                engine.GetB64systemConf(),
		"provisionScriptParametersCommon": fmt.Sprintf("[concat('ADMINUSER=',parameters('linuxAdminUsername'),' ETCD_DOWNLOAD_URL=',parameters('etcdDownloadURLBase'),' ETCD_VERSION=',parameters('etcdVersion'),' DOCKER_ENGINE_REPO=',parameters('dockerEngineDownloadRepo'),' TENANT_ID=',variables('tenantID'),' KUBERNETES_VERSION=%s HYPERKUBE_URL=',parameters('kubernetesHyperkubeSpec'),' APISERVER_PUBLIC_KEY=',parameters('apiserverCertificate'),' SUBSCRIPTION_ID=',variables('subscriptionId'),' RESOURCE_GROUP=',variables('resourceGroup'),' LOCATION=',variables('location'),' VM_TYPE=',variables('vmType'),' SUBNET=',variables('subnetName'),' NETWORK_SECURITY_GROUP=',variables('nsgName'),' VIRTUAL_NETWORK=',variables('virtualNetworkName'),' VIRTUAL_NETWORK_RESOURCE_GROUP=',variables('virtualNetworkResourceGroupName'),' ROUTE_TABLE=',variables('routeTableName'),' PRIMARY_AVAILABILITY_SET=',variables('primaryAvailabilitySetName'),' PRIMARY_SCALE_SET=',variables('primaryScaleSetName'),' SERVICE_PRINCIPAL_CLIENT_ID=',variables('servicePrincipalClientId'),' SERVICE_PRINCIPAL_CLIENT_SECRET=',variables('singleQuote'),variables('servicePrincipalClientSecret'),variables('singleQuote'),' KUBELET_PRIVATE_KEY=',parameters('clientPrivateKey'),' TARGET_ENVIRONMENT=',parameters('targetEnvironment'),' NETWORK_PLUGIN=',parameters('networkPlugin'),' NETWORK_POLICY=',parameters('networkPolicy'),' VNET_CNI_PLUGINS_URL=',parameters('vnetCniLinuxPluginsURL'),' CNI_PLUGINS_URL=',parameters('cniPluginsURL'),' CLOUDPROVIDER_BACKOFF=',toLower(string(parameters('cloudproviderConfig').cloudProviderBackoff)),' CLOUDPROVIDER_BACKOFF_RETRIES=',parameters('cloudproviderConfig').cloudProviderBackoffRetries,' CLOUDPROVIDER_BACKOFF_EXPONENT=',parameters('cloudproviderConfig').cloudProviderBackoffExponent,' CLOUDPROVIDER_BACKOFF_DURATION=',parameters('cloudproviderConfig').cloudProviderBackoffDuration,' CLOUDPROVIDER_BACKOFF_JITTER=',parameters('cloudproviderConfig').cloudProviderBackoffJitter,' CLOUDPROVIDER_RATELIMIT=',toLower(string(parameters('cloudproviderConfig').cloudProviderRatelimit)),' CLOUDPROVIDER_RATELIMIT_QPS=',parameters('cloudproviderConfig').cloudProviderRatelimitQPS,' CLOUDPROVIDER_RATELIMIT_BUCKET=',parameters('cloudproviderConfig').cloudProviderRatelimitBucket,' USE_MANAGED_IDENTITY_EXTENSION=',variables('useManagedIdentityExtension'),' USER_ASSIGNED_IDENTITY_ID=',variables('userAssignedClientID'),' USE_INSTANCE_METADATA=',variables('useInstanceMetadata'),' LOAD_BALANCER_SKU=',variables('loadBalancerSku'),' EXCLUDE_MASTER_FROM_STANDARD_LB=',variables('excludeMasterFromStandardLB'),' MAXIMUM_LOADBALANCER_RULE_COUNT=',variables('maximumLoadBalancerRuleCount'),' CONTAINER_RUNTIME=',parameters('containerRuntime'),' CONTAINERD_DOWNLOAD_URL_BASE=',parameters('containerdDownloadURLBase'),' POD_INFRA_CONTAINER_SPEC=',parameters('kubernetesPodInfraContainerSpec'),' KMS_PROVIDER_VAULT_NAME=',variables('clusterKeyVaultName'),' IS_HOSTED_MASTER=%t')]",
			cs.Properties.OrchestratorProfile.OrchestratorVersion, isHostedMaster),
		"generateProxyCertsScript":                  engine.GetKubernetesB64GenerateProxyCerts(),
		"orchestratorNameVersionTag":                fmt.Sprintf("%s:%s", cs.Properties.OrchestratorProfile.OrchestratorType, cs.Properties.OrchestratorProfile.OrchestratorVersion),
		"subnetNameResourceSegmentIndex":            10,
		"vnetNameResourceSegmentIndex":              8,
		"vnetResourceGroupNameResourceSegmentIndex": 4,
	}

	if useManagedIdentity {
		masterVars["servicePrincipalClientId"] = "msi"
		masterVars["servicePrincipalClientSecret"] = "msi"
	} else {
		masterVars["servicePrincipalClientId"] = "[parameters('servicePrincipalClientId')]"
		masterVars["servicePrincipalClientId"] = "[parameters('servicePrincipalClientSecret')]"
	}

	if !isHostedMaster {
		masterVars["masterCount"] = cs.Properties.MasterProfile.Count
		if isMasterVMSS {
			masterVars["masterOffset"] = ""
			masterVars["masterIpAddressCount"] = cs.Properties.MasterProfile.IPAddressCount
		} else {
			masterVars["masterOffset"] = "[parameters('masterOffset')]"
		}

		if cs.Properties.MasterProfile.IsCustomVNET() {
			if cs.Properties.MasterProfile.IsVirtualMachineScaleSets() {
				masterVars["vnetSubnetID"] = "[parameters('agentVnetSubnetID')]"
				masterVars["vnetSubnetIDMaster"] = "[parameters('masterVnetSubnetID')]"
			} else {
				masterVars["vnetSubnetID"] = "[parameters('masterVnetSubnetID')]"
			}
			masterVars["subnetName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('subnetNameResourceSegmentIndex')]]"
			masterVars["virtualNetworkName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetNameResourceSegmentIndex')]]"
			masterVars["virtualNetworkResourceGroupName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetResourceGroupNameResourceSegmentIndex')]]"
		} else {
			if cs.Properties.MasterProfile.IsVirtualMachineScaleSets() {
				masterVars["subnetName"] = "subnetmaster"
				masterVars["vnetSubnetID"] = "[concat(variables('vnetID'),'/subnets/subnetagent')]"
				masterVars["vnetSubnetIDMaster"] = "[concat(variables('vnetID'),'/subnets/subnetmaster')]"
			} else {
				masterVars["subnetName"] = "[concat(parameters('orchestratorName'), '-subnet')]"
				masterVars["vnetSubnetID"] = "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]"
				masterVars["virtualNetworkName"] = "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]"
				masterVars[    "vnetID"] = "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]"
				masterVars[    "virtualNetworkResourceGroupName"] = "''"
			}
		}
		masterVars["routeTableName"] = "[concat(variables('masterVMNamePrefix'),'routetable')]"
		if cs.Properties.MasterProfile.IsStorageAccount() {
			masterVars["masterStorageAccountName"] = "[concat(variables('storageAccountBaseName'), 'mstr0')]"
		}

		masterVars["nsgName"] = "[concat(variables('agentNamePrefix'), 'nsg')]"

	} else {
		if isCustomVnet {
			masterVars["vnetSubnetID"] = fmt.Sprintf("[parameters('%sVnetSubnetID')]", cs.Properties.AgentPoolProfiles[0].Name)
			masterVars["subnetName"] = "[split(variables('vnetSubnetID'), '/')[variables('subnetNameResourceSegmentIndex')]]"
			masterVars["virtualNetworkName"] = "[split(variables('vnetSubnetID'), '/')[variables('vnetNameResourceSegmentIndex')]]"
			masterVars["virtualNetworkResourceGroupName"] = "[split(variables('vnetSubnetID'), '/')[variables('vnetResourceGroupNameResourceSegmentIndex')]]"
		} else {
			masterVars["subnetName"] = "[concat(parameters('orchestratorName'), '-subnet')]"
			masterVars["vnetID"] = "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]"
			masterVars["vnetSubnetID"] = "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]"
			masterVars["virtualNetworkName"] = "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]"
			masterVars["virtualNetworkResourceGroupName"] = ""
		}
		masterVars["nsgName"] = "[concat(variables('masterVMNamePrefix'), 'nsg')]"
		masterVars["routeTableName"] = "[concat(variables('agentNamePrefix'), 'routetable')]"
	}

	masterVars["nsgID"] = "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]"

	if hasStorageAccountDisks {
		masterVars["maxVMsPerStorageAccount"] = 20
		masterVars["maxStorageAccountsPerAgent"] = "[div(variables('maxVMsPerPool'),variables('maxVMsPerStorageAccount'))]"
		masterVars["dataStorageAccountPrefixSeed"] = 97
		masterVars["storageAccountPrefixes"] = []string{"0", "6", "c", "i", "o", "u", "1", "7", "d", "j", "p", "v", "2", "8", "e", "k", "q", "w", "3", "9", "f", "l", "r", "x", "4", "a", "g", "m", "s", "y", "5", "b", "h", "n", "t", "z"}
		masterVars["storageAccountPrefixesCount"] = "[length(variables('storageAccountPrefixes'))]"
		masterVars["vmsPerStorageAccount"] = 20
		masterVars["storageAccountBaseName"] = "[uniqueString(concat(variables('masterFqdnPrefix'),variables('location')))]"
		masterVars["vmSizesMap"] = getSizeMap()["vmSizesMap"]
	} else {
		masterVars["storageAccountPrefixes"] = []interface{}{}
		masterVars["storageAccountBaseName"] = ""
	}

	if isAzureCNI {
		masterVars["allocateNodeCidrs"] = false
	} else {
		masterVars["allocateNodeCidrs"] = true
	}

	if cs.Properties.AnyAgentUsesVirtualMachineScaleSets() {
		masterVars["primaryScaleSetName"] = fmt.Sprintf("[concat(parameters('orchestratorName'), '-%s-',parameters('nameSuffix'), '-vmss')]", cs.Properties.AgentPoolProfiles[0].Name)
		masterVars["primaryAvailabilitySetName"] = ""
		masterVars["vmType"] = "vmss"
	} else {
		masterVars["primaryScaleSetName"] = fmt.Sprintf("[concat('%s-availabilitySet-',parameters('nameSuffix'))]", cs.Properties.AgentPoolProfiles[0].Name)
		masterVars["primaryAvailabilitySetName"] = ""
		masterVars["vmType"] = "standard"
	}

	if isHostedMaster {
		masterVars["kubernetesAPIServerIP"] = "[parameters('kubernetesEndpoint')]"
		masterVars["agentNamePrefix"] = "[concat(parameters('orchestratorName'), '-agentpool-', parameters('nameSuffix'), '-')]"
	} else {
		if isPrivateCluster {
			masterVars["kubeconfigServer"] = "[concat('https://', variables('kubernetesAPIServerIP'), ':443')]"
			if provisionJumpbox {
				masterVars["jumpboxOSDiskName"] = "[concat(parameters('jumpboxVMName'), '-osdisk')]"
				masterVars["jumpboxPublicIpAddressName"] = "[concat(parameters('jumpboxVMName'), '-ip')]"
				masterVars["jumpboxNetworkInterfaceName"] = "[concat(parameters('jumpboxVMName'), '-nic')]"
				masterVars["jumpboxNetworkSecurityGroupName"] = "[concat(parameters('jumpboxVMName'), '-nsg')]"

				kubeConfig, err := engine.GenerateKubeConfig(cs.Properties, cs.Location) //TODO: implement escape Char method of kubeConfig
				if err != nil {
					panic(err)
				}
				masterVars["kubeconfig"] = kubeConfig

				isJumpboxManagedDisks := cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() &&
					cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks

				if !isJumpboxManagedDisks {
					masterVars["jumpboxStorageAccountName"] = "[concat(variables('storageAccountBaseName'), 'jb')]"
				}

				if !cs.Properties.HasStorageAccountDisks() {
					masterVars["vmSizesMap"] = getSizeMap()["vmSizesMap"]
				}

			}
		} else {
			masterVars["masterPublicIPAddressName"] = "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]"
			masterVars["masterLbID"] = "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]"
			masterVars["masterLbIPConfigID"] = "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]"
			masterVars["masterLbIPConfigName"] = "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]"
			masterVars["masterLbName"] = "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]"
			masterVars["kubeconfigServer"] = "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]"
		}

		if cs.Properties.MasterProfile.Count > 1 {
			masterVars["masterInternalLbName"] = "[concat(parameters('orchestratorName'), '-master-internal-lb-', parameters('nameSuffix'))]"
			masterVars[    "masterInternalLbID"] = "[resourceId('Microsoft.Network/loadBalancers',variables('masterInternalLbName'))]"
			masterVars[    "masterInternalLbIPConfigName"] = "[concat(parameters('orchestratorName'), '-master-internal-lbFrontEnd-', parameters('nameSuffix'))]"
			masterVars[    "masterInternalLbIPConfigID"] = "[concat(variables('masterInternalLbID'),'/frontendIPConfigurations/', variables('masterInternalLbIPConfigName'))]"
			masterVars[    "masterInternalLbIPOffset"] = engine.DefaultInternalLbStaticIPOffset
			if isMasterVMSS {
				masterVars["kubernetesAPIServerIP"] = "[parameters('firstConsecutiveStaticIP')]"
			} else {
				masterVars["kubernetesAPIServerIP"] = "[concat(variables('masterFirstAddrPrefix'), add(variables('masterInternalLbIPOffset'), int(variables('masterFirstAddrOctet4'))))]"

			}
		} else {
			masterVars["kubernetesAPIServerIP"] = "[parameters('firstConsecutiveStaticIP')]"
		}

		masterVars["masterLbBackendPoolName"] = "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]"
		//masterVars["masterFirstAddrComment"] = "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space"
		masterVars["masterFirstAddrOctets"] = "[split(parameters('firstConsecutiveStaticIP'),'.')]"
		masterVars["masterFirstAddrOctet4"] = "[variables('masterFirstAddrOctets')[3]]"
		masterVars["masterFirstAddrPrefix"] = "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]"
		masterVars["masterEtcdServerPort"] = engine.DefaultMasterEtcdServerPort
		masterVars["masterEtcdClientPort"] = engine.DefaultMasterEtcdClientPort
	}

	if isMasterVMSS {
		masterVars["masterVMNamePrefix"] = "[concat(parameters('orchestratorName'), '-master-', parameters('nameSuffix'), '-')]"
	} else {
		masterVars["masterVMNamePrefix"] = cs.Properties.GetMasterVMPrefix()
		masterVars["masterVMNames"] = []string{
			"[concat(variables('masterVMNamePrefix'), '0')]",
			"[concat(variables('masterVMNamePrefix'), '1')]",
			"[concat(variables('masterVMNamePrefix'), '2')]",
			"[concat(variables('masterVMNamePrefix'), '3')]",
			"[concat(variables('masterVMNamePrefix'), '4')]",
		}
		masterVars["masterPrivateIpAddrs"] = []string{
			"[concat(variables('masterFirstAddrPrefix'), add(0, int(variables('masterFirstAddrOctet4'))))]",
			"[concat(variables('masterFirstAddrPrefix'), add(1, int(variables('masterFirstAddrOctet4'))))]",
			"[concat(variables('masterFirstAddrPrefix'), add(2, int(variables('masterFirstAddrOctet4'))))]",
			"[concat(variables('masterFirstAddrPrefix'), add(3, int(variables('masterFirstAddrOctet4'))))]",
			"[concat(variables('masterFirstAddrPrefix'), add(4, int(variables('masterFirstAddrOctet4'))))]",
		}
		masterVars["masterEtcdPeerURLs"] = []string{
			"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdServerPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdServerPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdServerPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdServerPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdServerPort'))]",
		}
		masterVars["masterEtcdClientURLs"] = []string{
			"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdClientPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdClientPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdClientPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdClientPort'))]",
			"[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdClientPort'))]",
		}
		masterVars["masterEtcdClusterStates"] = []string{
			"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0])]",
			"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2])]",
			"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2], ',', variables('masterVMNames')[3], '=', variables('masterEtcdPeerURLs')[3], ',', variables('masterVMNames')[4], '=', variables('masterEtcdPeerURLs')[4])]",
		}
	}

	masterVars["subscriptionId"] = "[subscription().subscriptionId]"
	masterVars["contributorRoleDefinitionId"] = "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]"
	masterVars["readerRoleDefinitionId"] = "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'acdd72a7-3385-48ef-bd42-f606fba81ae7')]"
	masterVars["scope"] = "[resourceGroup().id]"
	masterVars["tenantId"] = "[subscription().tenantId]"
	masterVars["singleQuote"] = "'"

	//TODO: Implement Linux HasSecrets

	if cs.Properties.HasWindows() {
		masterVars["windowsCustomScriptSuffix"] = " $inputFile = '%SYSTEMDRIVE%\\AzureData\\CustomData.bin' ; $outputFile = '%SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.ps1' ; Copy-Item $inputFile $outputFile ; Invoke-Expression('{0} {1}' -f $outputFile, $arguments) ; "
	}

	if enableEncryptionWithExternalKms {
		masterVars["clusterKeyVaultName"] = "[take(concat('kv', tolower(uniqueString(concat(variables('masterFqdnPrefix'),variables('location'),parameters('nameSuffix'))))), 22)]"

	}

	return masterVars
}

func getSizeMap() map[string]interface{} {
	var sizeMap map[string]interface{}
	sizeMapStr := fmt.Sprintf("{%s}", helpers.GetSizeMap())
	json.Unmarshal([]byte(sizeMapStr), &sizeMap)
	return sizeMap
}
