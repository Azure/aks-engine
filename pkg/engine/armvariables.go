// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/telemetry"

	"github.com/Azure/go-autorest/autorest/to"
)

func GetKubernetesVariables(cs *api.ContainerService) (map[string]interface{}, error) {
	k8sVars := map[string]interface{}{}
	profiles := cs.Properties.AgentPoolProfiles
	for i := 0; i < len(profiles); i++ {
		profile := profiles[i]
		k8sVars[fmt.Sprintf("%sIndex", profile.Name)] = i
		agentVars := getK8sAgentVars(cs, profile)

		for k, v := range agentVars {
			k8sVars[k] = v
		}

		if profile.IsStorageAccount() {
			if profile.HasDisks() {
				k8sVars[fmt.Sprintf("%sDataAccountName", profile.Name)] = fmt.Sprintf("[concat(variables('storageAccountBaseName'), 'data%d')]", i)
			}
			k8sVars[fmt.Sprintf("%sAccountName", profile.Name)] = fmt.Sprintf("[concat(variables('storageAccountBaseName'), 'agnt%d')]", i)
		}
	}

	masterVars, err := getK8sMasterVars(cs)
	if err != nil {
		return k8sVars, err
	}

	for k, v := range masterVars {
		k8sVars[k] = v
	}

	telemetryVars := getTelemetryVars(cs)
	for k, v := range telemetryVars {
		k8sVars[k] = v
	}

	windowsProfileVars := getWindowsProfileVars(cs.Properties.WindowsProfile)
	for k, v := range windowsProfileVars {
		k8sVars[k] = v
	}

	return k8sVars, nil
}

func getK8sMasterVars(cs *api.ContainerService) (map[string]interface{}, error) {
	orchProfile := cs.Properties.OrchestratorProfile
	kubernetesConfig := orchProfile.KubernetesConfig
	masterProfile := cs.Properties.MasterProfile
	profiles := cs.Properties.AgentPoolProfiles

	var useManagedIdentity, userAssignedID, userAssignedClientID, enableEncryptionWithExternalKms bool
	var excludeMasterFromStandardLB, provisionJumpbox bool
	var maxLoadBalancerCount int
	var useInstanceMetadata *bool
	var userAssignedIDReference string
	if kubernetesConfig != nil {
		useManagedIdentity = kubernetesConfig.UseManagedIdentity
		userAssignedID = kubernetesConfig.UserAssignedIDEnabled()
		userAssignedClientID = useManagedIdentity && kubernetesConfig.UserAssignedClientID != ""
		enableEncryptionWithExternalKms = to.Bool(kubernetesConfig.EnableEncryptionWithExternalKms)
		useInstanceMetadata = kubernetesConfig.UseInstanceMetadata
		excludeMasterFromStandardLB = to.Bool(kubernetesConfig.ExcludeMasterFromStandardLB)
		maxLoadBalancerCount = kubernetesConfig.MaximumLoadBalancerRuleCount
		provisionJumpbox = kubernetesConfig.PrivateJumpboxProvision()

		if kubernetesConfig.ShouldCreateNewUserAssignedIdentity() {
			userAssignedIDReference = "[resourceId('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))]"
		} else {
			userAssignedIDReference = "[variables('userAssignedID')]"
		}
	}
	isHostedMaster := cs.Properties.IsHostedMasterProfile()
	isMasterVMSS := masterProfile != nil && masterProfile.IsVirtualMachineScaleSets()
	hasStorageAccountDisks := cs.Properties.HasStorageAccountDisks()
	isCustomVnet := cs.Properties.AreAgentProfilesCustomVNET()
	hasAgentPool := len(profiles) > 0
	hasCosmosEtcd := masterProfile != nil && masterProfile.HasCosmosEtcd()
	scriptParamsInput := api.ProvisionScriptParametersInput{
		Location:             common.WrapAsARMVariable("location"),
		ResourceGroup:        common.WrapAsARMVariable("resourceGroup"),
		TenantID:             common.WrapAsARMVariable("tenantID"),
		SubscriptionID:       common.WrapAsARMVariable("subscriptionId"),
		ClientID:             common.WrapAsARMVariable("servicePrincipalClientId"),
		ClientSecret:         common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"),
		APIServerCertificate: common.WrapAsParameter("apiServerCertificate"),
		KubeletPrivateKey:    common.WrapAsParameter("clientPrivateKey"),
		ClusterKeyVaultName:  common.WrapAsARMVariable("clusterKeyVaultName"),
	}

	masterVars := map[string]interface{}{
		"maxVMsPerPool":                 100,
		"useManagedIdentityExtension":   strconv.FormatBool(useManagedIdentity),
		"userAssignedIDReference":       userAssignedIDReference,
		"useInstanceMetadata":           strconv.FormatBool(to.Bool(useInstanceMetadata)),
		"loadBalancerSku":               kubernetesConfig.LoadBalancerSku,
		"excludeMasterFromStandardLB":   strconv.FormatBool(excludeMasterFromStandardLB),
		"maximumLoadBalancerRuleCount":  maxLoadBalancerCount,
		"masterFqdnPrefix":              cs.Properties.GetDNSPrefix(),
		"apiVersionCompute":             api.APIVersionCompute,
		"apiVersionDeployments":         api.APIVersionDeployments,
		"apiVersionStorage":             api.APIVersionStorage,
		"apiVersionKeyVault":            api.APIVersionKeyVault,
		"apiVersionNetwork":             api.APIVersionNetwork,
		"apiVersionManagedIdentity":     api.APIVersionManagedIdentity,
		"apiVersionAuthorizationUser":   api.APIVersionAuthorizationUser,
		"apiVersionAuthorizationSystem": api.APIVersionAuthorizationSystem,
		"locations": []string{
			"[resourceGroup().location]",
			"[parameters('location')]",
		},
		"location":                                  "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"masterAvailabilitySet":                     "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"resourceGroup":                             "[resourceGroup().name]",
		"truncatedResourceGroup":                    "[take(replace(replace(resourceGroup().name, '(', '-'), ')', '-'), 63)]",
		"labelResourceGroup":                        "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"routeTableID":                              "[resourceId('Microsoft.Network/routeTables', variables('routeTableName'))]",
		"sshNatPorts":                               []int{22, 2201, 2202, 2203, 2204},
		"sshKeyPath":                                "[concat('/home/',parameters('linuxAdminUsername'),'/.ssh/authorized_keys')]",
		"provisionScriptParametersCommon":           "[concat('" + cs.GetProvisionScriptParametersCommon(scriptParamsInput) + "')]",
		"orchestratorNameVersionTag":                fmt.Sprintf("%s:%s", orchProfile.OrchestratorType, orchProfile.OrchestratorVersion),
		"vnetNameResourceSegmentIndex":              8,
		"vnetResourceGroupNameResourceSegmentIndex": 4,
	}

	cloudInitFiles := map[string]interface{}{
		"provisionScript":           getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
		"provisionSource":           getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
		"provisionInstalls":         getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
		"provisionConfigs":          getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
		"customSearchDomainsScript": getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
		"generateProxyCertsScript":  getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
		"mountEtcdScript":           getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
		"etcdSystemdService":        getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
		"dhcpv6SystemdService":      getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
		"dhcpv6ConfigurationScript": getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
		"kubeletSystemdService":     getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
	}

	if cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AADPodIdentityAddonName) {
		cloudInitFiles["untaintNodesScript"] = getBase64EncodedGzippedCustomScript(untaintNodesScript, cs)
		cloudInitFiles["untaintNodesSystemdService"] = getBase64EncodedGzippedCustomScript(untaintNodesSystemdService, cs)
	}

	if !cs.Properties.IsVHDDistroForAllNodes() {
		cloudInitFiles["provisionCIS"] = getBase64EncodedGzippedCustomScript(kubernetesCISScript, cs)
		cloudInitFiles["kmsSystemdService"] = getBase64EncodedGzippedCustomScript(kmsSystemdService, cs)
		cloudInitFiles["labelNodesScript"] = getBase64EncodedGzippedCustomScript(labelNodesScript, cs)
		cloudInitFiles["labelNodesSystemdService"] = getBase64EncodedGzippedCustomScript(labelNodesSystemdService, cs)
		cloudInitFiles["aptPreferences"] = getBase64EncodedGzippedCustomScript(aptPreferences, cs)
		cloudInitFiles["healthMonitorScript"] = getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript, cs)
		cloudInitFiles["kubeletMonitorSystemdService"] = getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService, cs)
		cloudInitFiles["dockerMonitorSystemdService"] = getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService, cs)
		cloudInitFiles["dockerMonitorSystemdTimer"] = getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer, cs)
		cloudInitFiles["kubeletSystemdService"] = getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs)
		cloudInitFiles["dockerClearMountPropagationFlags"] = getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags, cs)
		cloudInitFiles["auditdRules"] = getBase64EncodedGzippedCustomScript(auditdRules, cs)
	}

	if kubernetesConfig != nil {
		if kubernetesConfig.NetworkPlugin == NetworkPluginCilium {
			cloudInitFiles["systemdBPFMount"] = getBase64EncodedGzippedCustomScript(systemdBPFMount, cs)
		}
	}

	masterVars["cloudInitFiles"] = cloudInitFiles

	blockOutboundInternet := cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet")
	var cosmosEndPointURI string
	if hasCosmosEtcd {
		cosmosEndPointURI = fmt.Sprintf("%sk8s.etcd.cosmosdb.azure.com", masterProfile.DNSPrefix)
		masterVars["cosmosAccountName"] = fmt.Sprintf(etcdAccountNameFmt, cs.Properties.MasterProfile.DNSPrefix)
		masterVars["cosmosDBCertb64"] = base64.StdEncoding.EncodeToString([]byte(cs.Properties.CertificateProfile.EtcdClientCertificate))
		masterVars["apiVersionCosmos"] = "2015-04-08"
	} else {
		cosmosEndPointURI = ""
	}

	if cs.Properties.IsCustomCloudProfile() {
		if cs.Properties.IsAzureStackCloud() {
			masterVars["apiVersionCompute"] = "2017-03-30"
			masterVars["apiVersionStorage"] = "2017-10-01"
			masterVars["apiVersionNetwork"] = "2017-10-01"
			masterVars["apiVersionKeyVault"] = "2016-10-01"
		}

		environmentJSON, err := cs.Properties.GetCustomEnvironmentJSON(false)
		if err != nil {
			return masterVars, err
		}
		masterVars["environmentJSON"] = environmentJSON
		masterVars["provisionConfigsCustomCloud"] = getBase64EncodedGzippedCustomScript(kubernetesCSECustomCloud, cs)
	}

	masterVars["customCloudAuthenticationMethod"] = cs.Properties.GetCustomCloudAuthenticationMethod()
	masterVars["customCloudIdentifySystem"] = cs.Properties.GetCustomCloudIdentitySystem()

	auditDEnabled := "false"
	clusterAutoscalerEnabled := "false"

	if masterProfile != nil {
		auditDEnabled = strconv.FormatBool(to.Bool(masterProfile.AuditDEnabled))
		if kubernetesConfig != nil {
			clusterAutoscalerEnabled = strconv.FormatBool(kubernetesConfig.IsAddonEnabled(common.ClusterAutoscalerAddonName))
		}
	}
	if !isHostedMaster {
		if isMasterVMSS {
			masterVars["provisionScriptParametersMaster"] = fmt.Sprintf("[concat('COSMOS_URI=%s MASTER_NODE=true NO_OUTBOUND=%t AUDITD_ENABLED=%s CLUSTER_AUTOSCALER_ADDON=%s ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]", cosmosEndPointURI, blockOutboundInternet, auditDEnabled, clusterAutoscalerEnabled)
		} else {
			masterVars["provisionScriptParametersMaster"] = fmt.Sprintf("[concat('COSMOS_URI=%s MASTER_VM_NAME=',variables('masterVMNames')[variables('masterOffset')],' ETCD_PEER_URL=',variables('masterEtcdPeerURLs')[variables('masterOffset')],' ETCD_CLIENT_URL=',variables('masterEtcdClientURLs')[variables('masterOffset')],' MASTER_NODE=true NO_OUTBOUND=%t AUDITD_ENABLED=%s CLUSTER_AUTOSCALER_ADDON=%s ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]", cosmosEndPointURI, blockOutboundInternet, auditDEnabled, clusterAutoscalerEnabled)
		}
	}

	if userAssignedID {
		masterVars["userAssignedID"] = kubernetesConfig.UserAssignedID
	} else {
		masterVars["userAssignedID"] = ""
	}

	if userAssignedClientID {
		masterVars["userAssignedClientID"] = kubernetesConfig.UserAssignedClientID
	} else {
		masterVars["userAssignedClientID"] = ""
	}

	if !isHostedMaster {
		masterCount := masterProfile.Count

		if masterCount == 1 {
			masterVars["etcdPeerPrivateKeys"] = []string{"[parameters('etcdPeerPrivateKey0')]"}
			masterVars["etcdPeerCertificates"] = []string{"[parameters('etcdPeerCertificate0')]"}
		} else if masterCount == 3 {
			masterVars["etcdPeerPrivateKeys"] = []string{
				"[parameters('etcdPeerPrivateKey0')]",
				"[parameters('etcdPeerPrivateKey1')]",
				"[parameters('etcdPeerPrivateKey2')]",
			}
			masterVars["etcdPeerCertificates"] = []string{
				"[parameters('etcdPeerCertificate0')]",
				"[parameters('etcdPeerCertificate1')]",
				"[parameters('etcdPeerCertificate2')]",
			}
		} else if masterCount == 5 {
			masterVars["etcdPeerPrivateKeys"] = []string{
				"[parameters('etcdPeerPrivateKey0')]",
				"[parameters('etcdPeerPrivateKey1')]",
				"[parameters('etcdPeerPrivateKey2')]",
				"[parameters('etcdPeerPrivateKey3')]",
				"[parameters('etcdPeerPrivateKey4')]",
			}
			masterVars["etcdPeerCertificates"] = []string{
				"[parameters('etcdPeerCertificate0')]",
				"[parameters('etcdPeerCertificate1')]",
				"[parameters('etcdPeerCertificate2')]",
				"[parameters('etcdPeerCertificate3')]",
				"[parameters('etcdPeerCertificate4')]",
			}
		}
		masterVars["etcdPeerCertFilepath"] = []string{
			"/etc/kubernetes/certs/etcdpeer0.crt",
			"/etc/kubernetes/certs/etcdpeer1.crt",
			"/etc/kubernetes/certs/etcdpeer2.crt",
			"/etc/kubernetes/certs/etcdpeer3.crt",
			"/etc/kubernetes/certs/etcdpeer4.crt",
		}

		masterVars["etcdPeerKeyFilepath"] = []string{
			"/etc/kubernetes/certs/etcdpeer0.key",
			"/etc/kubernetes/certs/etcdpeer1.key",
			"/etc/kubernetes/certs/etcdpeer2.key",
			"/etc/kubernetes/certs/etcdpeer3.key",
			"/etc/kubernetes/certs/etcdpeer4.key",
		}

		masterVars["etcdCaFilepath"] = "/etc/kubernetes/certs/ca.crt"
		masterVars["etcdClientCertFilepath"] = "/etc/kubernetes/certs/etcdclient.crt"
		masterVars["etcdClientKeyFilepath"] = "/etc/kubernetes/certs/etcdclient.key"
		masterVars["etcdServerCertFilepath"] = "/etc/kubernetes/certs/etcdserver.crt"
		masterVars["etcdServerKeyFilepath"] = "/etc/kubernetes/certs/etcdserver.key"
	}
	if useManagedIdentity && !isHostedMaster {
		masterVars["servicePrincipalClientId"] = "msi"
		masterVars["servicePrincipalClientSecret"] = "msi"
	} else {
		masterVars["servicePrincipalClientId"] = "[parameters('servicePrincipalClientId')]"
		masterVars["servicePrincipalClientSecret"] = "[parameters('servicePrincipalClientSecret')]"
	}

	if !isHostedMaster {
		masterVars["masterCount"] = masterProfile.Count
		if isMasterVMSS {
			masterVars["masterOffset"] = ""
			masterVars["masterIpAddressCount"] = masterProfile.IPAddressCount
		} else {
			masterVars["masterOffset"] = "[parameters('masterOffset')]"
		}

		if masterProfile.IsCustomVNET() {
			if masterProfile.IsVirtualMachineScaleSets() {
				masterVars["vnetSubnetID"] = "[parameters('agentVnetSubnetID')]"
				masterVars["vnetSubnetIDMaster"] = "[parameters('masterVnetSubnetID')]"
			} else {
				masterVars["vnetSubnetID"] = "[parameters('masterVnetSubnetID')]"
			}
			masterVars["virtualNetworkName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetNameResourceSegmentIndex')]]"
			masterVars["virtualNetworkResourceGroupName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetResourceGroupNameResourceSegmentIndex')]]"
		} else {
			if masterProfile.IsVirtualMachineScaleSets() {
				masterVars["vnetSubnetID"] = "[concat(variables('vnetID'),'/subnets/subnetagent')]"
				masterVars["vnetSubnetIDMaster"] = "[concat(variables('vnetID'),'/subnets/subnetmaster')]"
			} else {
				masterVars["vnetSubnetID"] = "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]"

			}
			masterVars["virtualNetworkName"] = "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]"
			masterVars["vnetID"] = "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]"
			masterVars["virtualNetworkResourceGroupName"] = "''"
		}
		masterVars["routeTableName"] = "[concat(variables('masterVMNamePrefix'),'routetable')]"
		if masterProfile.IsStorageAccount() {
			masterVars["masterStorageAccountName"] = "[concat(variables('storageAccountBaseName'), 'mstr0')]"
		}
		masterVars["nsgName"] = "[concat(variables('masterVMNamePrefix'), 'nsg')]"

	} else {
		if isCustomVnet {
			if hasAgentPool {
				masterVars["vnetSubnetID"] = fmt.Sprintf("[parameters('%sVnetSubnetID')]", profiles[0].Name)
			}
			masterVars["virtualNetworkName"] = "[split(variables('vnetSubnetID'), '/')[variables('vnetNameResourceSegmentIndex')]]"
			masterVars["virtualNetworkResourceGroupName"] = "[split(variables('vnetSubnetID'), '/')[variables('vnetResourceGroupNameResourceSegmentIndex')]]"
		} else {
			masterVars["vnetID"] = "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]"
			masterVars["vnetSubnetID"] = "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]"
			masterVars["virtualNetworkName"] = "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]"
			masterVars["virtualNetworkResourceGroupName"] = ""
		}
		masterVars["nsgName"] = "[concat(variables('agentNamePrefix'), 'nsg')]"

		masterVars["routeTableName"] = "[concat(variables('agentNamePrefix'), 'routetable')]"
	}

	masterVars["nsgID"] = "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]"
	masterVars["subnetName"] = cs.Properties.GetSubnetName()

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

	masterVars["vmType"] = cs.Properties.GetVMType()

	if cs.Properties.HasVMSSAgentPool() {
		masterVars["primaryAvailabilitySetName"] = ""
	} else {
		if hasAgentPool {
			masterVars["primaryAvailabilitySetName"] = fmt.Sprintf("[concat('%s-availabilitySet-',parameters('nameSuffix'))]", profiles[0].Name)
		} else {
			masterVars["primaryAvailabilitySetName"] = ""
		}
	}
	masterVars["primaryScaleSetName"] = cs.Properties.GetPrimaryScaleSetName()

	if isHostedMaster {
		masterVars["kubernetesAPIServerIP"] = "[parameters('kubernetesEndpoint')]"
		masterVars["agentNamePrefix"] = "[concat(parameters('orchestratorName'), '-agentpool-', parameters('nameSuffix'), '-')]"
	} else {
		if cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku && hasAgentPool {
			masterVars["agentLbID"] = "[resourceId('Microsoft.Network/loadBalancers',variables('agentLbName'))]"
			masterVars["agentLbName"] = "[parameters('masterEndpointDNSNamePrefix')]"
			masterVars["agentLbBackendPoolName"] = "[parameters('masterEndpointDNSNamePrefix')]"
			numIps := 1
			if cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerOutboundIPs != nil {
				numIps = *cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerOutboundIPs
			}
			ipAddressNameVarPrefix := "agentPublicIPAddressName"
			outboundIPNamePrefix := "agent-ip-outbound"
			outboundConfigNamePrefix := "agent-outbound"
			agentLbIPConfigIDVarPrefix := "agentLbIPConfigID"
			agentLbIPConfigNameVarPrefix := "agentLbIPConfigName"
			for i := 1; i <= numIps; i++ {
				ipAddressNameVar := ipAddressNameVarPrefix
				outboundIPName := outboundIPNamePrefix
				agentLbIPConfigIDVar := agentLbIPConfigIDVarPrefix
				agentLbIPConfigNameVar := agentLbIPConfigNameVarPrefix
				outboundConfigName := outboundConfigNamePrefix
				if i > 1 {
					ipAddressNameVar += strconv.Itoa(i)
					outboundIPName += strconv.Itoa(i)
					outboundConfigName += strconv.Itoa(i)
					agentLbIPConfigIDVar += strconv.Itoa(i)
					agentLbIPConfigNameVar += strconv.Itoa(i)
				}
				masterVars[ipAddressNameVar] = fmt.Sprintf("[concat(parameters('orchestratorName'), '-%s')]", outboundIPName)
				masterVars[agentLbIPConfigIDVar] = fmt.Sprintf("[concat(variables('agentLbID'),'/frontendIPConfigurations/', variables('%s'))]", agentLbIPConfigNameVar)
				masterVars[agentLbIPConfigNameVar] = fmt.Sprintf("[concat(parameters('orchestratorName'), '-%s')]", outboundConfigName)
			}
		}
		// private cluster + basic LB configurations do not need these vars (which serve the master LB and public IP resources), because:
		// - private cluster + basic LB + 1 master uses NIC outbound rules for master outbound access
		// - private cluster + basic LB + multiple masters uses internal master LB for outbound access (doesn't need a public IP)
		if !(cs.Properties.OrchestratorProfile.IsPrivateCluster() && cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.BasicLoadBalancerSku) {
			masterVars["masterPublicIPAddressName"] = "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]"
			masterVars["masterLbID"] = "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]"
			masterVars["masterLbIPConfigID"] = "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]"
			masterVars["masterLbIPConfigName"] = "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]"
			masterVars["masterLbName"] = "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]"
		}
		if cs.Properties.OrchestratorProfile.IsPrivateCluster() {
			masterVars["kubeconfigServer"] = "[concat('https://', variables('kubernetesAPIServerIP'), ':443')]"
			if provisionJumpbox {
				masterVars["jumpboxOSDiskName"] = "[concat(parameters('jumpboxVMName'), '-osdisk')]"
				masterVars["jumpboxPublicIpAddressName"] = "[concat(parameters('jumpboxVMName'), '-ip')]"
				masterVars["jumpboxNetworkInterfaceName"] = "[concat(parameters('jumpboxVMName'), '-nic')]"
				masterVars["jumpboxNetworkSecurityGroupName"] = "[concat(parameters('jumpboxVMName'), '-nsg')]"

				kubeConfig, err := GenerateKubeConfig(cs.Properties, cs.Location)
				if err != nil {
					panic(err)
				}
				masterVars["kubeconfig"] = kubeConfig

				isJumpboxManagedDisks := kubernetesConfig.PrivateJumpboxProvision() &&
					kubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks

				if !isJumpboxManagedDisks {
					masterVars["jumpboxStorageAccountName"] = "[concat(variables('storageAccountBaseName'), 'jb')]"
				}

				if !cs.Properties.HasStorageAccountDisks() {
					masterVars["vmSizesMap"] = getSizeMap()["vmSizesMap"]
				}

			}
		} else {
			masterVars["kubeconfigServer"] = "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]"
		}

		if masterProfile.HasMultipleNodes() {
			masterVars["masterInternalLbName"] = "[concat(parameters('orchestratorName'), '-master-internal-lb-', parameters('nameSuffix'))]"
			masterVars["masterInternalLbID"] = "[resourceId('Microsoft.Network/loadBalancers',variables('masterInternalLbName'))]"
			masterVars["masterInternalLbIPConfigName"] = "[concat(parameters('orchestratorName'), '-master-internal-lbFrontEnd-', parameters('nameSuffix'))]"
			masterVars["masterInternalLbIPConfigID"] = "[concat(variables('masterInternalLbID'),'/frontendIPConfigurations/', variables('masterInternalLbIPConfigName'))]"
			masterVars["masterInternalLbIPOffset"] = DefaultInternalLbStaticIPOffset
			if isMasterVMSS {
				masterVars["kubernetesAPIServerIP"] = "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.255.', variables('masterInternalLbIPOffset'))]"
			} else {
				masterVars["kubernetesAPIServerIP"] = "[concat(variables('masterFirstAddrPrefix'), add(variables('masterInternalLbIPOffset'), int(variables('masterFirstAddrOctet4'))))]"

			}
		} else {
			masterVars["kubernetesAPIServerIP"] = "[parameters('firstConsecutiveStaticIP')]"
		}

		masterVars["masterLbBackendPoolName"] = "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]"
		masterVars["masterFirstAddrComment"] = "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space"
		masterVars["masterFirstAddrOctets"] = "[split(parameters('firstConsecutiveStaticIP'),'.')]"
		masterVars["masterFirstAddrOctet4"] = "[variables('masterFirstAddrOctets')[3]]"
		masterVars["masterFirstAddrPrefix"] = "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]"
		masterVars["masterEtcdServerPort"] = DefaultMasterEtcdServerPort
		masterVars["masterEtcdClientPort"] = DefaultMasterEtcdClientPort

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
			masterVars["masterEtcdMetricURLs"] = []string{
				"[concat('http://', variables('masterPrivateIpAddrs')[0], ':2480')]",
				"[concat('http://', variables('masterPrivateIpAddrs')[1], ':2480')]",
				"[concat('http://', variables('masterPrivateIpAddrs')[2], ':2480')]",
				"[concat('http://', variables('masterPrivateIpAddrs')[3], ':2480')]",
				"[concat('http://', variables('masterPrivateIpAddrs')[4], ':2480')]",
			}
		}
	}

	masterVars["subscriptionId"] = "[subscription().subscriptionId]"
	masterVars["contributorRoleDefinitionId"] = "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]"
	masterVars["readerRoleDefinitionId"] = "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'acdd72a7-3385-48ef-bd42-f606fba81ae7')]"
	if cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AppGwIngressAddonName) {
		masterVars["managedIdentityOperatorRoleDefinitionId"] = "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'f1a07417-d97a-45cb-824c-7a7467783830')]"
	}
	masterVars["networkContributorRoleDefinitionId"] = "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', '4d97b98b-1d4f-4787-a291-c67834d212e7')]"
	masterVars["scope"] = "[resourceGroup().id]"
	masterVars["tenantId"] = "[subscription().tenantId]"
	masterVars["singleQuote"] = "'"

	if cs.Properties.HasWindows() {
		masterVars["windowsCustomScriptSuffix"] = " $inputFile = '%SYSTEMDRIVE%\\AzureData\\CustomData.bin' ; $outputFile = '%SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.ps1' ; Copy-Item $inputFile $outputFile ; Invoke-Expression('{0} {1}' -f $outputFile, $arguments) ; "
	}

	if enableEncryptionWithExternalKms {
		masterVars["clusterKeyVaultName"] = "[take(concat('kv', tolower(uniqueString(concat(variables('masterFqdnPrefix'),variables('location'),parameters('nameSuffix'))))), 22)]"
	} else {
		masterVars["clusterKeyVaultName"] = ""
	}

	if cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AppGwIngressAddonName) {
		masterVars["appGwName"] = "[concat(parameters('orchestratorName'), '-appgw-', parameters('nameSuffix'))]"
		masterVars["appGwSubnetName"] = "appgw-subnet"
		masterVars["appGwPublicIPAddressName"] = "[concat(parameters('orchestratorName'), '-appgw-ip-', parameters('nameSuffix'))]"
		masterVars["appGwICIdentityName"] = "[concat(parameters('orchestratorName'), '-appgw-ic-identity-', parameters('nameSuffix'))]"
		masterVars["appGwId"] = "[resourceId('Microsoft.Network/applicationGateways',variables('appGwName'))]"
		masterVars["appGwICIdentityId"] = "[resourceId('Microsoft.ManagedIdentity/userAssignedIdentities', variables('appGwICIdentityName'))]"
	}

	return masterVars, nil
}

func getK8sAgentVars(cs *api.ContainerService, profile *api.AgentPoolProfile) map[string]interface{} {
	agentVars := map[string]interface{}{}
	agentName := profile.Name

	storageAccountOffset := fmt.Sprintf("%sStorageAccountOffset", agentName)
	storageAccountsCount := fmt.Sprintf("%sStorageAccountsCount", agentName)
	agentsCount := fmt.Sprintf("%sCount", agentName)
	agentsVMNamePrefix := fmt.Sprintf("%sVMNamePrefix", agentName)
	agentOffset := fmt.Sprintf("%sOffset", agentName)
	agentAvailabilitySet := fmt.Sprintf("%sAvailabilitySet", agentName)
	agentScaleSetPriority := fmt.Sprintf("%sScaleSetPriority", agentName)
	agentScaleSetEvictionPolicy := fmt.Sprintf("%sScaleSetEvictionPolicy", agentName)
	agentVMSize := fmt.Sprintf("%sVMSize", agentName)
	agentVnetSubnetID := fmt.Sprintf("%sVnetSubnetID", agentName)
	agentSubnetName := fmt.Sprintf("%sSubnetName", agentName)
	agentVnetParts := fmt.Sprintf("%sVnetParts", agentName)
	agentSubnetResourceGroup := fmt.Sprintf("%sSubnetResourceGroup", agentName)
	agentVnet := fmt.Sprintf("%sVnet", agentName)

	agentOsImageOffer := fmt.Sprintf("%sosImageOffer", agentName)
	agentOsImageSku := fmt.Sprintf("%sosImageSKU", agentName)
	agentOsImagePublisher := fmt.Sprintf("%sosImagePublisher", agentName)
	agentOsImageVersion := fmt.Sprintf("%sosImageVersion", agentName)
	agentOsImageName := fmt.Sprintf("%sosImageName", agentName)
	agentOsImageResourceGroup := fmt.Sprintf("%sosImageResourceGroup", agentName)

	if profile.IsStorageAccount() {
		agentVars[storageAccountOffset] = fmt.Sprintf("[mul(variables('maxStorageAccountsPerAgent'),variables('%sIndex'))]", agentName)
		agentVars[storageAccountsCount] = fmt.Sprintf("[add(div(variables('%[1]sCount'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('%[1]sCount'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('%[1]sCount'), variables('maxVMsPerStorageAccount')),1)))]", agentName)
	}

	agentVars[agentsCount] = fmt.Sprintf("[parameters('%s')]", agentsCount)
	agentVars[agentsVMNamePrefix] = cs.Properties.GetAgentVMPrefix(profile, cs.Properties.GetAgentPoolIndexByName(agentName))

	if profile.IsWindows() {
		agentVars["winResourceNamePrefix"] = "[substring(parameters('nameSuffix'), 0, 5)]"
	}

	if profile.IsAvailabilitySets() {
		agentVars[agentOffset] = fmt.Sprintf("[parameters('%s')]", agentOffset)
		agentVars[agentAvailabilitySet] = fmt.Sprintf("[concat('%s-availabilitySet-', parameters('nameSuffix'))]", agentName)
	} else if profile.IsLowPriorityScaleSet() || profile.IsSpotScaleSet() {
		agentVars[agentScaleSetPriority] = fmt.Sprintf("[parameters('%s')]", agentScaleSetPriority)
		agentVars[agentScaleSetEvictionPolicy] = fmt.Sprintf("[parameters('%s')]", agentScaleSetEvictionPolicy)
	}
	agentVars[agentVMSize] = fmt.Sprintf("[parameters('%s')]", agentVMSize)

	if profile.IsCustomVNET() {
		agentVars[agentVnetSubnetID] = fmt.Sprintf("[parameters('%s')]", agentVnetSubnetID)
		agentVars[agentSubnetName] = fmt.Sprintf("[parameters('%s')]", agentVnetSubnetID)
		agentVars[agentVnetParts] = fmt.Sprintf("[split(parameters('%sVnetSubnetID'),'/subnets/')]", agentName)
	} else {
		agentVars[agentVnetSubnetID] = "[variables('vnetSubnetID')]"
		agentVars[agentSubnetName] = "[variables('subnetName')]"
	}

	agentVars[agentSubnetResourceGroup] = fmt.Sprintf("[split(variables('%sVnetSubnetID'), '/')[4]]", agentName)
	agentVars[agentVnet] = fmt.Sprintf("[split(variables('%sVnetSubnetID'), '/')[8]]", agentName)

	agentVars[agentOsImageOffer] = fmt.Sprintf("[parameters('%sosImageOffer')]", agentName)
	agentVars[agentOsImageSku] = fmt.Sprintf("[parameters('%sosImageSKU')]", agentName)
	agentVars[agentOsImagePublisher] = fmt.Sprintf("[parameters('%sosImagePublisher')]", agentName)
	agentVars[agentOsImageVersion] = fmt.Sprintf("[parameters('%sosImageVersion')]", agentName)
	agentVars[agentOsImageName] = fmt.Sprintf("[parameters('%sosImageName')]", agentName)
	agentVars[agentOsImageResourceGroup] = fmt.Sprintf("[parameters('%sosImageResourceGroup')]", agentName)

	return agentVars
}

func getTelemetryVars(cs *api.ContainerService) map[string]interface{} {

	enableTelemetry := false
	if cs.Properties.FeatureFlags != nil {
		enableTelemetry = cs.Properties.FeatureFlags.EnableTelemetry
	}

	applicationInsightsKey := ""
	if cs.Properties.TelemetryProfile != nil {
		applicationInsightsKey = telemetry.AKSEngineAppInsightsKey
	}

	telemetryVars := map[string]interface{}{
		"enableTelemetry":        enableTelemetry,
		"applicationInsightsKey": applicationInsightsKey,
	}

	return telemetryVars
}

func getWindowsProfileVars(wp *api.WindowsProfile) map[string]interface{} {
	enableCSIProxy := common.DefaultEnableCSIProxyWindows
	CSIProxyURL := ""

	if wp != nil {
		enableCSIProxy = wp.IsCSIProxyEnabled()
		CSIProxyURL = wp.CSIProxyURL
	}
	vars := map[string]interface{}{
		"windowsEnableCSIProxy": enableCSIProxy,
		"windowsCSIProxyURL":    CSIProxyURL,
	}
	return vars
}

func getSizeMap() map[string]interface{} {
	var sizeMap map[string]interface{}
	sizeMapStr := fmt.Sprintf("{%s}", helpers.GetSizeMap())
	_ = json.Unmarshal([]byte(sizeMapStr), &sizeMap)
	return sizeMap
}
