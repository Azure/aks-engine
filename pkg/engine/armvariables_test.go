// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
)

var testK8sVersion = common.GetSupportedKubernetesVersion("1.12", false)

func TestSizeMap(t *testing.T) {
	sizeMap := getSizeMap()
	_, err := json.MarshalIndent(sizeMap["vmSizesMap"], "", "   ")
	if err != nil {
		t.Errorf("unexpected error while attempting to marshal the size map: %s", err.Error())
	}
}

func TestK8sVars(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:     1,
				DNSPrefix: "blueorange",
				VMSize:    "Standard_D2_v2",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: api.Kubernetes,
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:   "agentpool1",
					VMSize: "Standard_D2_v2",
					Count:  2,
				},
			},
		},
	}

	cs.SetPropertiesDefaults(false, false)

	varMap, err := GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap := map[string]interface{}{
		"agentpool1Count":                "[parameters('agentpool1Count')]",
		"agentpool1Index":                0,
		"agentpool1SubnetName":           "[variables('subnetName')]",
		"agentpool1VMNamePrefix":         "k8s-agentpool1-18280257-vmss",
		"agentpool1VMSize":               "[parameters('agentpool1VMSize')]",
		"agentpool1VnetSubnetID":         "[variables('vnetSubnetID')]",
		"agentpool1osImageName":          "[parameters('agentpool1osImageName')]",
		"agentpool1osImageOffer":         "[parameters('agentpool1osImageOffer')]",
		"agentpool1osImagePublisher":     "[parameters('agentpool1osImagePublisher')]",
		"agentpool1osImageResourceGroup": "[parameters('agentpool1osImageResourceGroup')]",
		"agentpool1osImageSKU":           "[parameters('agentpool1osImageSKU')]",
		"agentpool1osImageVersion":       "[parameters('agentpool1osImageVersion')]",
		"apiVersionAuthorizationSystem":  "2018-01-01-preview",
		"apiVersionAuthorizationUser":    "2018-09-01-preview",
		"apiVersionCompute":              "2018-10-01",
		"apiVersionDeployments":          "2018-06-01",
		"apiVersionKeyVault":             "2018-02-14",
		"apiVersionManagedIdentity":      "2015-08-31-preview",
		"apiVersionNetwork":              "2018-08-01",
		"apiVersionStorage":              "2018-07-01",
		"clusterKeyVaultName":            "",
		"contributorRoleDefinitionId":    "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]",
		"etcdCaFilepath":                 "/etc/kubernetes/certs/ca.crt",
		"etcdClientCertFilepath":         "/etc/kubernetes/certs/etcdclient.crt",
		"etcdClientKeyFilepath":          "/etc/kubernetes/certs/etcdclient.key",
		"etcdPeerCertFilepath":           []string{"/etc/kubernetes/certs/etcdpeer0.crt", "/etc/kubernetes/certs/etcdpeer1.crt", "/etc/kubernetes/certs/etcdpeer2.crt", "/etc/kubernetes/certs/etcdpeer3.crt", "/etc/kubernetes/certs/etcdpeer4.crt"},
		"etcdPeerCertificates":           []string{"[parameters('etcdPeerCertificate0')]"},
		"etcdPeerKeyFilepath":            []string{"/etc/kubernetes/certs/etcdpeer0.key", "/etc/kubernetes/certs/etcdpeer1.key", "/etc/kubernetes/certs/etcdpeer2.key", "/etc/kubernetes/certs/etcdpeer3.key", "/etc/kubernetes/certs/etcdpeer4.key"},
		"etcdPeerPrivateKeys":            []string{"[parameters('etcdPeerPrivateKey0')]"},
		"etcdServerCertFilepath":         "/etc/kubernetes/certs/etcdserver.crt",
		"etcdServerKeyFilepath":          "/etc/kubernetes/certs/etcdserver.key",
		"excludeMasterFromStandardLB":    "false",
		"kubeconfigServer":               "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]",
		"kubernetesAPIServerIP":          "[parameters('firstConsecutiveStaticIP')]",
		"labelResourceGroup":             "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"loadBalancerSku":                "Basic",
		"location":                       "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"locations":                      []string{"[resourceGroup().location]", "[parameters('location')]"},
		"masterAvailabilitySet":          "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"masterCount":                    1,
		"masterEtcdClientPort":           2379,
		"masterEtcdClientURLs":           []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdClientPort'))]"},
		"masterEtcdClusterStates":        []string{"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2], ',', variables('masterVMNames')[3], '=', variables('masterEtcdPeerURLs')[3], ',', variables('masterVMNames')[4], '=', variables('masterEtcdPeerURLs')[4])]"},
		"masterEtcdPeerURLs":             []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdServerPort'))]"},
		"masterEtcdServerPort":           2380,
		"masterFirstAddrComment":         "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space",
		"masterFirstAddrOctet4":          "[variables('masterFirstAddrOctets')[3]]",
		"masterFirstAddrOctets":          "[split(parameters('firstConsecutiveStaticIP'),'.')]",
		"masterFirstAddrPrefix":          "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
		"masterFqdnPrefix":               "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
		"masterLbBackendPoolName":        "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]",
		"masterLbID":                     "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
		"masterLbIPConfigID":             "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
		"masterLbIPConfigName":           "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterLbName":                   "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]",
		"masterOffset":                   "[parameters('masterOffset')]",
		"masterPrivateIpAddrs":           []string{"[concat(variables('masterFirstAddrPrefix'), add(0, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(1, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(2, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(3, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(4, int(variables('masterFirstAddrOctet4'))))]"},
		"masterPublicIPAddressName":      "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]",
		"masterVMNamePrefix":             "k8s-master-18280257-",
		"masterVMNames":                  []string{"[concat(variables('masterVMNamePrefix'), '0')]", "[concat(variables('masterVMNamePrefix'), '1')]", "[concat(variables('masterVMNamePrefix'), '2')]", "[concat(variables('masterVMNamePrefix'), '3')]", "[concat(variables('masterVMNamePrefix'), '4')]"},
		"maxVMsPerPool":                  100,
		"maximumLoadBalancerRuleCount":   250,
		"nsgID":                          "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]",
		"nsgName":                        "[concat(variables('masterVMNamePrefix'), 'nsg')]",
		"orchestratorNameVersionTag":     "Kubernetes:" + testK8sVersion,
		"primaryAvailabilitySetName":     "",
		"primaryScaleSetName":            cs.Properties.GetPrimaryScaleSetName(),
		"cloudInitFiles": map[string]interface{}{
			"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript),
			"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript),
			"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall),
			"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig),
			"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript),
			"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript),
			"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript),
			"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript),
			"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd),
			"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService),
			"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService),
			"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService),
			"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer),
			"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService),
			"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences),
			"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags),
			"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService),
		},
		"provisionScriptParametersCommon":           fmt.Sprintf("[concat('ADMINUSER=',parameters('linuxAdminUsername'),' ETCD_DOWNLOAD_URL=',parameters('etcdDownloadURLBase'),' ETCD_VERSION=',parameters('etcdVersion'),' CONTAINERD_VERSION=',parameters('containerdVersion'),' MOBY_VERSION=',parameters('mobyVersion'),' TENANT_ID=',variables('tenantID'),' KUBERNETES_VERSION=%s HYPERKUBE_URL=',parameters('kubernetesHyperkubeSpec'),' APISERVER_PUBLIC_KEY=',parameters('apiServerCertificate'),' SUBSCRIPTION_ID=',variables('subscriptionId'),' RESOURCE_GROUP=',variables('resourceGroup'),' LOCATION=',variables('location'),' VM_TYPE=',variables('vmType'),' SUBNET=',variables('subnetName'),' NETWORK_SECURITY_GROUP=',variables('nsgName'),' VIRTUAL_NETWORK=',variables('virtualNetworkName'),' VIRTUAL_NETWORK_RESOURCE_GROUP=',variables('virtualNetworkResourceGroupName'),' ROUTE_TABLE=',variables('routeTableName'),' PRIMARY_AVAILABILITY_SET=',variables('primaryAvailabilitySetName'),' PRIMARY_SCALE_SET=',variables('primaryScaleSetName'),' SERVICE_PRINCIPAL_CLIENT_ID=',variables('servicePrincipalClientId'),' SERVICE_PRINCIPAL_CLIENT_SECRET=',variables('singleQuote'),variables('servicePrincipalClientSecret'),variables('singleQuote'),' KUBELET_PRIVATE_KEY=',parameters('clientPrivateKey'),' TARGET_ENVIRONMENT=',parameters('targetEnvironment'),' NETWORK_PLUGIN=',parameters('networkPlugin'),' NETWORK_POLICY=',parameters('networkPolicy'),' VNET_CNI_PLUGINS_URL=',parameters('vnetCniLinuxPluginsURL'),' CNI_PLUGINS_URL=',parameters('cniPluginsURL'),' CLOUDPROVIDER_BACKOFF=',toLower(string(parameters('cloudproviderConfig').cloudProviderBackoff)),' CLOUDPROVIDER_BACKOFF_RETRIES=',parameters('cloudproviderConfig').cloudProviderBackoffRetries,' CLOUDPROVIDER_BACKOFF_EXPONENT=',parameters('cloudproviderConfig').cloudProviderBackoffExponent,' CLOUDPROVIDER_BACKOFF_DURATION=',parameters('cloudproviderConfig').cloudProviderBackoffDuration,' CLOUDPROVIDER_BACKOFF_JITTER=',parameters('cloudproviderConfig').cloudProviderBackoffJitter,' CLOUDPROVIDER_RATELIMIT=',toLower(string(parameters('cloudproviderConfig').cloudProviderRatelimit)),' CLOUDPROVIDER_RATELIMIT_QPS=',parameters('cloudproviderConfig').cloudProviderRatelimitQPS,' CLOUDPROVIDER_RATELIMIT_BUCKET=',parameters('cloudproviderConfig').cloudProviderRatelimitBucket,' USE_MANAGED_IDENTITY_EXTENSION=',variables('useManagedIdentityExtension'),' USE_INSTANCE_METADATA=',variables('useInstanceMetadata'),' LOAD_BALANCER_SKU=',variables('loadBalancerSku'),' EXCLUDE_MASTER_FROM_STANDARD_LB=',variables('excludeMasterFromStandardLB'),' MAXIMUM_LOADBALANCER_RULE_COUNT=',variables('maximumLoadBalancerRuleCount'),' CONTAINER_RUNTIME=',parameters('containerRuntime'),' CONTAINERD_DOWNLOAD_URL_BASE=',parameters('containerdDownloadURLBase'),' POD_INFRA_CONTAINER_SPEC=',parameters('kubernetesPodInfraContainerSpec'),' KMS_PROVIDER_VAULT_NAME=',variables('clusterKeyVaultName'),' IS_HOSTED_MASTER=false',' PRIVATE_AZURE_REGISTRY_SERVER=',parameters('privateAzureRegistryServer'),' AUTHENTICATION_METHOD=',variables('customCloudAuthenticationMethod'),' IDENTITY_SYSTEM=',variables('customCloudIdentifySystem'))]", testK8sVersion),
		"provisionScriptParametersMaster":           "[concat('COSMOS_URI= MASTER_VM_NAME=',variables('masterVMNames')[variables('masterOffset')],' ETCD_PEER_URL=',variables('masterEtcdPeerURLs')[variables('masterOffset')],' ETCD_CLIENT_URL=',variables('masterEtcdClientURLs')[variables('masterOffset')],' MASTER_NODE=true NO_OUTBOUND=false AUDITD_ENABLED=false CLUSTER_AUTOSCALER_ADDON=',parameters('kubernetesClusterAutoscalerEnabled'),' ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]",
		"readerRoleDefinitionId":                    "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'acdd72a7-3385-48ef-bd42-f606fba81ae7')]",
		"resourceGroup":                             "[resourceGroup().name]",
		"routeTableID":                              "[resourceId('Microsoft.Network/routeTables', variables('routeTableName'))]",
		"routeTableName":                            "[concat(variables('masterVMNamePrefix'),'routetable')]",
		"scope":                                     "[resourceGroup().id]",
		"servicePrincipalClientId":                  "[parameters('servicePrincipalClientId')]",
		"servicePrincipalClientSecret":              "[parameters('servicePrincipalClientSecret')]",
		"singleQuote":                               "'",
		"sshKeyPath":                                "[concat('/home/',parameters('linuxAdminUsername'),'/.ssh/authorized_keys')]",
		"sshNatPorts":                               []int{22, 2201, 2202, 2203, 2204},
		"storageAccountBaseName":                    "",
		"storageAccountPrefixes":                    []interface{}{},
		"subnetName":                                "[concat(parameters('orchestratorName'), '-subnet')]",
		"subnetNameResourceSegmentIndex":            10,
		"subscriptionId":                            "[subscription().subscriptionId]",
		"tenantId":                                  "[subscription().tenantId]",
		"truncatedResourceGroup":                    "[take(replace(replace(resourceGroup().name, '(', '-'), ')', '-'), 63)]",
		"useInstanceMetadata":                       "true",
		"useManagedIdentityExtension":               "false",
		"userAssignedClientID":                      "",
		"userAssignedID":                            "",
		"userAssignedIDReference":                   "[resourceId('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))]",
		"virtualNetworkName":                        "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]",
		"virtualNetworkResourceGroupName":           "''",
		"vmType":                                    "vmss",
		"vnetID":                                    "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]",
		"vnetNameResourceSegmentIndex":              8,
		"vnetResourceGroupNameResourceSegmentIndex": 4,
		"vnetSubnetID":                              "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]",
		"customCloudAuthenticationMethod":           cs.Properties.GetCustomCloudAuthenticationMethod(),
		"customCloudIdentifySystem":                 cs.Properties.GetCustomCloudIdentitySystem(),
	}

	diff := cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with ubuntu 16.04 distro
	cs.Properties.AgentPoolProfiles[0].Distro = api.Ubuntu
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript),
		"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript),
		"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall),
		"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig),
		"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript),
		"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript),
		"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript),
		"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript),
		"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd),
		"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService),
		"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService),
		"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService),
		"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer),
		"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService),
		"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences),
		"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags),
		"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService),
	}

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with ubuntu 18.04 distro
	cs.Properties.AgentPoolProfiles[0].Distro = api.Ubuntu1804
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript),
		"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript),
		"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall),
		"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig),
		"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript),
		"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript),
		"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript),
		"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript),
		"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd),
		"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService),
		"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService),
		"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService),
		"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer),
		"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService),
		"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences),
		"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags),
		"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService),
	}

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with CustomVnet enabled
	cs.Properties.MasterProfile.VnetSubnetID = "fooSubnetID"
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["subnetName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('subnetNameResourceSegmentIndex')]]"
	expectedMap["virtualNetworkName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetNameResourceSegmentIndex')]]"
	expectedMap["virtualNetworkResourceGroupName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetResourceGroupNameResourceSegmentIndex')]]"
	expectedMap["vnetSubnetID"] = "[parameters('masterVnetSubnetID')]"
	delete(expectedMap, "vnetID")

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with  3 Multiple Master Nodes
	cs.Properties.MasterProfile.Count = 3
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["etcdPeerCertificates"] = []string{
		"[parameters('etcdPeerCertificate0')]",
		"[parameters('etcdPeerCertificate1')]",
		"[parameters('etcdPeerCertificate2')]",
	}
	expectedMap["etcdPeerPrivateKeys"] = []string{
		"[parameters('etcdPeerPrivateKey0')]",
		"[parameters('etcdPeerPrivateKey1')]",
		"[parameters('etcdPeerPrivateKey2')]",
	}
	expectedMap["kubernetesAPIServerIP"] = "[concat(variables('masterFirstAddrPrefix'), add(variables('masterInternalLbIPOffset'), int(variables('masterFirstAddrOctet4'))))]"
	expectedMap["masterCount"] = 3
	expectedMap["masterInternalLbID"] = "[resourceId('Microsoft.Network/loadBalancers',variables('masterInternalLbName'))]"
	expectedMap["masterInternalLbIPConfigID"] = "[concat(variables('masterInternalLbID'),'/frontendIPConfigurations/', variables('masterInternalLbIPConfigName'))]"
	expectedMap["masterInternalLbIPConfigName"] = "[concat(parameters('orchestratorName'), '-master-internal-lbFrontEnd-', parameters('nameSuffix'))]"
	expectedMap["masterInternalLbIPOffset"] = 10
	expectedMap["masterInternalLbName"] = "[concat(parameters('orchestratorName'), '-master-internal-lb-', parameters('nameSuffix'))]"

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with  5 Multiple Master Nodes
	cs.Properties.MasterProfile.Count = 5
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["etcdPeerCertificates"] = []string{
		"[parameters('etcdPeerCertificate0')]",
		"[parameters('etcdPeerCertificate1')]",
		"[parameters('etcdPeerCertificate2')]",
		"[parameters('etcdPeerCertificate3')]",
		"[parameters('etcdPeerCertificate4')]",
	}
	expectedMap["etcdPeerPrivateKeys"] = []string{
		"[parameters('etcdPeerPrivateKey0')]",
		"[parameters('etcdPeerPrivateKey1')]",
		"[parameters('etcdPeerPrivateKey2')]",
		"[parameters('etcdPeerPrivateKey3')]",
		"[parameters('etcdPeerPrivateKey4')]",
	}
	expectedMap["masterCount"] = 5

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with HostedMasterProfile and StorageAccount
	cs.Properties.MasterProfile = nil
	cs.Properties.HostedMasterProfile = &api.HostedMasterProfile{
		DNSPrefix: "fooDNSPrefix",
	}
	cs.Properties.AgentPoolProfiles[0].StorageProfile = api.StorageAccount
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["agentNamePrefix"] = "[concat(parameters('orchestratorName'), '-agentpool-', parameters('nameSuffix'), '-')]"
	expectedMap["agentpool1AccountName"] = "[concat(variables('storageAccountBaseName'), 'agnt0')]"
	expectedMap["agentpool1StorageAccountOffset"] = "[mul(variables('maxStorageAccountsPerAgent'),variables('agentpool1Index'))]"
	expectedMap["agentpool1StorageAccountsCount"] = "[add(div(variables('agentpool1Count'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('agentpool1Count'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('agentpool1Count'), variables('maxVMsPerStorageAccount')),1)))]"
	expectedMap["agentpool1VMNamePrefix"] = "aks-agentpool1-18280257-vmss"
	expectedMap["dataStorageAccountPrefixSeed"] = 97
	expectedMap["kubernetesAPIServerIP"] = "[parameters('kubernetesEndpoint')]"
	expectedMap["masterVMNamePrefix"] = "aks-master-18280257-"
	expectedMap["maxStorageAccountsPerAgent"] = "[div(variables('maxVMsPerPool'),variables('maxVMsPerStorageAccount'))]"
	expectedMap["maxVMsPerStorageAccount"] = 20
	expectedMap["nsgName"] = "[concat(variables('agentNamePrefix'), 'nsg')]"
	expectedMap["provisionScriptParametersCommon"] = fmt.Sprintf("[concat('ADMINUSER=',parameters('linuxAdminUsername'),' ETCD_DOWNLOAD_URL=',parameters('etcdDownloadURLBase'),' ETCD_VERSION=',parameters('etcdVersion'),' CONTAINERD_VERSION=',parameters('containerdVersion'),' MOBY_VERSION=',parameters('mobyVersion'),' TENANT_ID=',variables('tenantID'),' KUBERNETES_VERSION=%s HYPERKUBE_URL=',parameters('kubernetesHyperkubeSpec'),' APISERVER_PUBLIC_KEY=',parameters('apiServerCertificate'),' SUBSCRIPTION_ID=',variables('subscriptionId'),' RESOURCE_GROUP=',variables('resourceGroup'),' LOCATION=',variables('location'),' VM_TYPE=',variables('vmType'),' SUBNET=',variables('subnetName'),' NETWORK_SECURITY_GROUP=',variables('nsgName'),' VIRTUAL_NETWORK=',variables('virtualNetworkName'),' VIRTUAL_NETWORK_RESOURCE_GROUP=',variables('virtualNetworkResourceGroupName'),' ROUTE_TABLE=',variables('routeTableName'),' PRIMARY_AVAILABILITY_SET=',variables('primaryAvailabilitySetName'),' PRIMARY_SCALE_SET=',variables('primaryScaleSetName'),' SERVICE_PRINCIPAL_CLIENT_ID=',variables('servicePrincipalClientId'),' SERVICE_PRINCIPAL_CLIENT_SECRET=',variables('singleQuote'),variables('servicePrincipalClientSecret'),variables('singleQuote'),' KUBELET_PRIVATE_KEY=',parameters('clientPrivateKey'),' TARGET_ENVIRONMENT=',parameters('targetEnvironment'),' NETWORK_PLUGIN=',parameters('networkPlugin'),' NETWORK_POLICY=',parameters('networkPolicy'),' VNET_CNI_PLUGINS_URL=',parameters('vnetCniLinuxPluginsURL'),' CNI_PLUGINS_URL=',parameters('cniPluginsURL'),' CLOUDPROVIDER_BACKOFF=',toLower(string(parameters('cloudproviderConfig').cloudProviderBackoff)),' CLOUDPROVIDER_BACKOFF_RETRIES=',parameters('cloudproviderConfig').cloudProviderBackoffRetries,' CLOUDPROVIDER_BACKOFF_EXPONENT=',parameters('cloudproviderConfig').cloudProviderBackoffExponent,' CLOUDPROVIDER_BACKOFF_DURATION=',parameters('cloudproviderConfig').cloudProviderBackoffDuration,' CLOUDPROVIDER_BACKOFF_JITTER=',parameters('cloudproviderConfig').cloudProviderBackoffJitter,' CLOUDPROVIDER_RATELIMIT=',toLower(string(parameters('cloudproviderConfig').cloudProviderRatelimit)),' CLOUDPROVIDER_RATELIMIT_QPS=',parameters('cloudproviderConfig').cloudProviderRatelimitQPS,' CLOUDPROVIDER_RATELIMIT_BUCKET=',parameters('cloudproviderConfig').cloudProviderRatelimitBucket,' USE_MANAGED_IDENTITY_EXTENSION=',variables('useManagedIdentityExtension'),' USE_INSTANCE_METADATA=',variables('useInstanceMetadata'),' LOAD_BALANCER_SKU=',variables('loadBalancerSku'),' EXCLUDE_MASTER_FROM_STANDARD_LB=',variables('excludeMasterFromStandardLB'),' MAXIMUM_LOADBALANCER_RULE_COUNT=',variables('maximumLoadBalancerRuleCount'),' CONTAINER_RUNTIME=',parameters('containerRuntime'),' CONTAINERD_DOWNLOAD_URL_BASE=',parameters('containerdDownloadURLBase'),' POD_INFRA_CONTAINER_SPEC=',parameters('kubernetesPodInfraContainerSpec'),' KMS_PROVIDER_VAULT_NAME=',variables('clusterKeyVaultName'),' IS_HOSTED_MASTER=true',' PRIVATE_AZURE_REGISTRY_SERVER=',parameters('privateAzureRegistryServer'),' AUTHENTICATION_METHOD=',variables('customCloudAuthenticationMethod'),' IDENTITY_SYSTEM=',variables('customCloudIdentifySystem'))]", testK8sVersion)
	expectedMap["routeTableName"] = "[concat(variables('agentNamePrefix'), 'routetable')]"
	expectedMap["storageAccountBaseName"] = "[uniqueString(concat(variables('masterFqdnPrefix'),variables('location')))]"
	expectedMap["storageAccountPrefixes"] = []string{"0", "6", "c", "i", "o", "u", "1", "7", "d", "j", "p", "v", "2", "8", "e", "k", "q", "w", "3", "9", "f", "l", "r", "x", "4", "a", "g", "m", "s", "y", "5", "b", "h", "n", "t", "z"}
	expectedMap["storageAccountPrefixesCount"] = "[length(variables('storageAccountPrefixes'))]"
	expectedMap["subnetName"] = "[concat(parameters('orchestratorName'), '-subnet')]"
	expectedMap["virtualNetworkName"] = "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]"
	expectedMap["virtualNetworkResourceGroupName"] = ""
	expectedMap["vmSizesMap"] = map[string]interface{}{"Standard_A0": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A1": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A10": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A11": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A1_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A2_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A2m_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A4": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A4_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A4m_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A5": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A6": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A7": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A8": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A8_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A8m_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_A9": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_B1ls": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_B1ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_B1s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_B2ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_B2s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_B4ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_B8ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_D1": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D11": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D11_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D12": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D12_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D13": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D13_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D14": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D14_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D15_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D16_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D16s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_D1_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D2_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D2_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D2s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_D3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D32_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D32s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_D3_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D4": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D4_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D4_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D4s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_D5_v2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D64_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D64s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_D8_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_D8s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DC2s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DC4s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS1": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS11": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS11-1_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS11_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS12": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS12-1_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS12-2_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS12_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS13": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS13-2_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS13-4_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS13_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS14": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS14-4_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS14-8_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS14_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS15_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS1_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS2_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS3_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS4": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS4_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_DS5_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E16-4s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E16-8s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E16_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E16s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E20_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E20s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E2_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E2s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E32-16s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E32-8s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E32_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E32s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E4-2s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E4_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E4s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E64-16s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E64-32s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E64_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E64i_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E64is_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E64s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E8-2s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E8-4s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_E8_v3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_E8s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F1": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_F16": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_F16s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F16s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F1s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_F2s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F2s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F32s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F4": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_F4s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F4s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F64s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F72s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F8": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_F8s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_F8s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_G1": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_G2": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_G3": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_G4": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_G5": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_GS1": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS4": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS4-4": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS4-8": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS5": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS5-16": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_GS5-8": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_H16": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_H16m": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_H16mr": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_H16r": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_H8": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_H8m": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_HB60rs": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_HC44rs": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L16s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L16s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L32s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L32s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L4s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L64s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L80s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L8s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_L8s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M128": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_M128-32ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M128-64ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M128m": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_M128ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M128s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M16-4ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M16-8ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M16ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M32-16ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M32-8ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M32ls": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M32ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M32ts": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M64": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_M64-16ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M64-32ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M64ls": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M64m": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_M64ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M64s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M8-2ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M8-4ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_M8ms": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC12": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_NC12s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC12s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC24": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_NC24r": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_NC24rs_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC24rs_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC24s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC24s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC6": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_NC6s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NC6s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_ND12s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_ND24rs": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_ND24s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_ND6s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NV12": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_NV12s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NV12s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NV24": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_NV24s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NV24s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NV48s_v3": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_NV6": map[string]interface{}{"storageAccountType": "Standard_LRS"}, "Standard_NV6s_v2": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_PB12s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_PB24s": map[string]interface{}{"storageAccountType": "Premium_LRS"}, "Standard_PB6s": map[string]interface{}{"storageAccountType": "Premium_LRS"}}
	expectedMap["vmsPerStorageAccount"] = 20
	expectedMap["vnetID"] = "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]"
	expectedMap["vnetSubnetID"] = "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]"
	expectedMap["primaryScaleSetName"] = cs.Properties.GetPrimaryScaleSetName()

	delete(expectedMap, "etcdCaFilepath")
	delete(expectedMap, "etcdClientCertFilepath")
	delete(expectedMap, "etcdClientKeyFilepath")
	delete(expectedMap, "etcdPeerCertFilepath")
	delete(expectedMap, "etcdPeerCertificates")
	delete(expectedMap, "etcdPeerKeyFilepath")
	delete(expectedMap, "etcdPeerPrivateKeys")
	delete(expectedMap, "etcdServerCertFilepath")
	delete(expectedMap, "etcdServerKeyFilepath")
	delete(expectedMap, "masterCount")
	delete(expectedMap, "masterEtcdClientPort")
	delete(expectedMap, "masterEtcdServerPort")
	delete(expectedMap, "masterFirstAddrComment")
	delete(expectedMap, "masterFirstAddrOctets")
	delete(expectedMap, "masterFirstAddrPrefix")
	delete(expectedMap, "masterInternalLbID")
	delete(expectedMap, "masterInternalLbIPConfigID")
	delete(expectedMap, "masterInternalLbIPConfigName")
	delete(expectedMap, "masterInternalLbIPOffset")
	delete(expectedMap, "masterInternalLbName")
	delete(expectedMap, "kubeConfigServer")
	delete(expectedMap, "masterFirstAddrOctet4")
	delete(expectedMap, "masterLbBackendPoolName")
	delete(expectedMap, "masterLbID")
	delete(expectedMap, "masterLbIPConfigID")
	delete(expectedMap, "masterLbIPConfigName")
	delete(expectedMap, "masterLbName")
	delete(expectedMap, "masterOffset")
	delete(expectedMap, "masterPublicIPAddressName")
	delete(expectedMap, "provisionScriptParametersMaster")
	delete(expectedMap, "kubeconfigServer")
	delete(expectedMap, "masterVMNamePrefix")
	delete(expectedMap, "masterVMNames")
	delete(expectedMap, "masterPrivateIpAddrs")
	delete(expectedMap, "masterEtcdPeerURLs")
	delete(expectedMap, "masterEtcdClusterStates")
	delete(expectedMap, "masterEtcdClientURLs")

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with Custom cloud

	const (
		name                         = "azurestackcloud"
		managementPortalURL          = "https://management.local.azurestack.external/"
		publishSettingsURL           = "https://management.local.azurestack.external/publishsettings/index"
		serviceManagementEndpoint    = "https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189"
		resourceManagerEndpoint      = "https://management.local.azurestack.external/"
		activeDirectoryEndpoint      = "https://login.windows.net/"
		galleryEndpoint              = "https://portal.local.azurestack.external=30015/"
		keyVaultEndpoint             = "https://vault.azurestack.external/"
		graphEndpoint                = "https://graph.windows.net/"
		serviceBusEndpoint           = "https://servicebus.azurestack.external/"
		batchManagementEndpoint      = "https://batch.azurestack.external/"
		storageEndpointSuffix        = "core.azurestack.external"
		sqlDatabaseDNSSuffix         = "database.azurestack.external"
		trafficManagerDNSSuffix      = "trafficmanager.cn"
		keyVaultDNSSuffix            = "vault.azurestack.external"
		serviceBusEndpointSuffix     = "servicebus.azurestack.external"
		serviceManagementVMDNSSuffix = "chinacloudapp.cn"
		resourceManagerVMDNSSuffix   = "cloudapp.azurestack.external"
		containerRegistryDNSSuffix   = "azurecr.io"
		tokenAudience                = "https://management.azurestack.external/"
	)

	customCloudK8sVersion := testK8sVersion
	if name == "azurestackcloud" {
		customCloudK8sVersion = testK8sVersion + AzureStackSuffix
	}

	cs = &api.ContainerService{
		Location: "local",
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:     1,
				DNSPrefix: "blueorange",
				VMSize:    "Standard_D2_v2",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: api.Kubernetes,
			},
			LinuxProfile: &api.LinuxProfile{},
			CustomCloudProfile: &api.CustomCloudProfile{
				IdentitySystem:       api.AzureADIdentitySystem,
				AuthenticationMethod: api.ClientSecretAuthMethod,
				Environment: &azure.Environment{
					Name:                         name,
					ManagementPortalURL:          managementPortalURL,
					PublishSettingsURL:           publishSettingsURL,
					ServiceManagementEndpoint:    serviceManagementEndpoint,
					ResourceManagerEndpoint:      resourceManagerEndpoint,
					ActiveDirectoryEndpoint:      activeDirectoryEndpoint,
					GalleryEndpoint:              galleryEndpoint,
					KeyVaultEndpoint:             keyVaultEndpoint,
					GraphEndpoint:                graphEndpoint,
					ServiceBusEndpoint:           serviceBusEndpoint,
					BatchManagementEndpoint:      batchManagementEndpoint,
					StorageEndpointSuffix:        storageEndpointSuffix,
					SQLDatabaseDNSSuffix:         sqlDatabaseDNSSuffix,
					TrafficManagerDNSSuffix:      trafficManagerDNSSuffix,
					KeyVaultDNSSuffix:            keyVaultDNSSuffix,
					ServiceBusEndpointSuffix:     serviceBusEndpointSuffix,
					ServiceManagementVMDNSSuffix: serviceManagementVMDNSSuffix,
					ResourceManagerVMDNSSuffix:   resourceManagerVMDNSSuffix,
					ContainerRegistryDNSSuffix:   containerRegistryDNSSuffix,
					TokenAudience:                tokenAudience,
				},
			},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:   "agentpool1",
					VMSize: "Standard_D2_v2",
					Count:  2,
				},
			},
		},
	}

	cs.SetPropertiesDefaults(false, false)

	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap = map[string]interface{}{
		"agentpool1Count":                 "[parameters('agentpool1Count')]",
		"agentpool1Index":                 0,
		"agentpool1SubnetName":            "[variables('subnetName')]",
		"agentpool1VMNamePrefix":          "k8s-agentpool1-18280257-vmss",
		"agentpool1VMSize":                "[parameters('agentpool1VMSize')]",
		"agentpool1VnetSubnetID":          "[variables('vnetSubnetID')]",
		"agentpool1osImageName":           "[parameters('agentpool1osImageName')]",
		"agentpool1osImageOffer":          "[parameters('agentpool1osImageOffer')]",
		"agentpool1osImagePublisher":      "[parameters('agentpool1osImagePublisher')]",
		"agentpool1osImageResourceGroup":  "[parameters('agentpool1osImageResourceGroup')]",
		"agentpool1osImageSKU":            "[parameters('agentpool1osImageSKU')]",
		"agentpool1osImageVersion":        "[parameters('agentpool1osImageVersion')]",
		"apiVersionAuthorizationSystem":   "2018-01-01-preview",
		"apiVersionAuthorizationUser":     "2018-09-01-preview",
		"apiVersionCompute":               "2017-03-30",
		"apiVersionDeployments":           "2018-06-01",
		"apiVersionKeyVault":              "2016-10-01",
		"environmentJSON":                 `{"name":"azurestackcloud","managementPortalURL":"https://management.local.azurestack.external/","publishSettingsURL":"https://management.local.azurestack.external/publishsettings/index","serviceManagementEndpoint":"https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189","resourceManagerEndpoint":"https://management.local.azurestack.external/","activeDirectoryEndpoint":"https://login.windows.net/","galleryEndpoint":"https://portal.local.azurestack.external=30015/","keyVaultEndpoint":"https://vault.azurestack.external/","graphEndpoint":"https://graph.windows.net/","serviceBusEndpoint":"https://servicebus.azurestack.external/","batchManagementEndpoint":"https://batch.azurestack.external/","storageEndpointSuffix":"core.azurestack.external","sqlDatabaseDNSSuffix":"database.azurestack.external","trafficManagerDNSSuffix":"trafficmanager.cn","keyVaultDNSSuffix":"vault.azurestack.external","serviceBusEndpointSuffix":"servicebus.azurestack.external","serviceManagementVMDNSSuffix":"chinacloudapp.cn","resourceManagerVMDNSSuffix":"cloudapp.azurestack.external","containerRegistryDNSSuffix":"azurecr.io","tokenAudience":"https://management.azurestack.external/"}`,
		"customCloudAuthenticationMethod": "client_secret",
		"customCloudIdentifySystem":       "azure_ad",
		"apiVersionManagedIdentity":       "2015-08-31-preview",
		"apiVersionNetwork":               "2017-10-01",
		"apiVersionStorage":               "2017-10-01",
		"clusterKeyVaultName":             "",
		"contributorRoleDefinitionId":     "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]",
		"etcdCaFilepath":                  "/etc/kubernetes/certs/ca.crt",
		"etcdClientCertFilepath":          "/etc/kubernetes/certs/etcdclient.crt",
		"etcdClientKeyFilepath":           "/etc/kubernetes/certs/etcdclient.key",
		"etcdPeerCertFilepath":            []string{"/etc/kubernetes/certs/etcdpeer0.crt", "/etc/kubernetes/certs/etcdpeer1.crt", "/etc/kubernetes/certs/etcdpeer2.crt", "/etc/kubernetes/certs/etcdpeer3.crt", "/etc/kubernetes/certs/etcdpeer4.crt"},
		"etcdPeerCertificates":            []string{"[parameters('etcdPeerCertificate0')]"},
		"etcdPeerKeyFilepath":             []string{"/etc/kubernetes/certs/etcdpeer0.key", "/etc/kubernetes/certs/etcdpeer1.key", "/etc/kubernetes/certs/etcdpeer2.key", "/etc/kubernetes/certs/etcdpeer3.key", "/etc/kubernetes/certs/etcdpeer4.key"},
		"etcdPeerPrivateKeys":             []string{"[parameters('etcdPeerPrivateKey0')]"},
		"etcdServerCertFilepath":          "/etc/kubernetes/certs/etcdserver.crt",
		"etcdServerKeyFilepath":           "/etc/kubernetes/certs/etcdserver.key",
		"excludeMasterFromStandardLB":     "false",
		"kubeconfigServer":                "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]",
		"kubernetesAPIServerIP":           "[parameters('firstConsecutiveStaticIP')]",
		"labelResourceGroup":              "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"loadBalancerSku":                 "Basic",
		"location":                        "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"locations":                       []string{"[resourceGroup().location]", "[parameters('location')]"},
		"masterAvailabilitySet":           "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"masterCount":                     1,
		"masterEtcdClientPort":            2379,
		"masterEtcdClientURLs":            []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdClientPort'))]"},
		"masterEtcdClusterStates":         []string{"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2], ',', variables('masterVMNames')[3], '=', variables('masterEtcdPeerURLs')[3], ',', variables('masterVMNames')[4], '=', variables('masterEtcdPeerURLs')[4])]"},
		"masterEtcdPeerURLs":              []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdServerPort'))]"},
		"masterEtcdServerPort":            2380,
		"masterFirstAddrComment":          "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space",
		"masterFirstAddrOctet4":           "[variables('masterFirstAddrOctets')[3]]",
		"masterFirstAddrOctets":           "[split(parameters('firstConsecutiveStaticIP'),'.')]",
		"masterFirstAddrPrefix":           "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
		"masterFqdnPrefix":                "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
		"masterLbBackendPoolName":         "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]",
		"masterLbID":                      "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
		"masterLbIPConfigID":              "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
		"masterLbIPConfigName":            "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterLbName":                    "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]",
		"masterOffset":                    "[parameters('masterOffset')]",
		"masterPrivateIpAddrs":            []string{"[concat(variables('masterFirstAddrPrefix'), add(0, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(1, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(2, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(3, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(4, int(variables('masterFirstAddrOctet4'))))]"},
		"masterPublicIPAddressName":       "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]",
		"masterVMNamePrefix":              "k8s-master-18280257-",
		"masterVMNames":                   []string{"[concat(variables('masterVMNamePrefix'), '0')]", "[concat(variables('masterVMNamePrefix'), '1')]", "[concat(variables('masterVMNamePrefix'), '2')]", "[concat(variables('masterVMNamePrefix'), '3')]", "[concat(variables('masterVMNamePrefix'), '4')]"},
		"maxVMsPerPool":                   100,
		"maximumLoadBalancerRuleCount":    250,
		"nsgID":                           "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]",
		"nsgName":                         "[concat(variables('masterVMNamePrefix'), 'nsg')]",
		"orchestratorNameVersionTag":      "Kubernetes:" + testK8sVersion,
		"primaryAvailabilitySetName":      "",
		"primaryScaleSetName":             cs.Properties.GetPrimaryScaleSetName(),
		"cloudInitFiles": map[string]interface{}{
			"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript),
			"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript),
			"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall),
			"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig),
			"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript),
			"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript),
			"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript),
			"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript),
			"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd),
			"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService),
			"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService),
			"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService),
			"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer),
			"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService),
			"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences),
			"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags),
			"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService),
		},
		"provisionConfigsCustomCloud":               "H4sIAAAAAAAA/7RXa3PaSBb9zq+4I7Nlu8ZCOJNxpbzLpDRC2bC2wSXJk8qGlKpRX0HHolvpbvmxNv99q5uHMWCbuDL6QIF0de7pc5/s/OINGPcGRI1qNeSqkhig1CxnGdGo9vbhrgYA4P/3IgrjxA9O0rD7Vyfqdc/CbpL+J+5103M/+dhyPNSZd1kNUHLUqDzyv0qi0iS7zApR0cY3JbizihWFce8iCsL0zO/6/w6jNOy2z3udbtKq7337Dg2JSlQywzPCyRBlyGkpGNdQf4kP3IOW4FJw+o6z/7TbMPHbfuI/+HXq29DzxqgJJZp4OKOk3pOSuVcoFRO89aZ5+LvbPHSbh9MjZ5UsYDPyKgNrH1xEp2kUJhdRN+i1w1b9vb19cvFnmAa9bhL1Tk/DaEHrQ+c0bK0GYEw4y1FpZW+6meBaiqJA6Y6najZuybiwuCyHL1BfcQq/tKAJX/8JeoTcmplrByIsC5Ih2M+RKChKyIUEpQoYME4ZHy6sLbCbQ/055ms+zKWQgsvAUff/uhJFNUalij/uXeBkjMfGV59PDUdC6XOiR8fzGwCl+Qn9vlGk3/eMcd/LUGp1P3Se5/I8hzNRcf0UEXONjcH5k+6XTSUS2uPF7TFoWeH2xHJWWwpGIMpbxodGP/BNxcWm4kAKoSF7KGTQwpqQspSilMzcUlpI+2CAUJWUaKSNBfJyqvZ6SRqEUdL50An8JExnmTsr+ysivYINvGtChsi1t9w+GiWOna0w03YYJ3PISkmvEBkpPDUiEr2MuEtnMa1FLXlpZFI/OMnKx3X2DPnnDR8YLbCnKrkrdOxjLBT+hKzfmQYyOoN554NZscK80cxr7FF4mTJZpDTSA5A4FldogbYp0eVE95ZqzaOvKRVvtVK2h8lZbTm/d2B9HME1KwqTrxK1ZEiN0uaceMM0ZIJaIbjQ0HxV17Uw9fe1Sa2WCZ6zYSXx5J0KKqXFODBTbDEP8aYUUpsgr7OsTctby9tsTFOWpzlhRSURjppw2ITfmmDGLbjZU+8q1PDrzUyFTyPk4F8kH8NuYjKz0+sa+h97bXPYrGDIdbqUCwdWkZXksLVOgXEt4BJv4YoUlT6YefA5hTiM/uqYuog63aBz7p+mwWnHTNQ4DKIwWQhvsHNRFOLapKCZ6VCS20IQCtdMj8zJkB69BeQmHFMHdzM/AOCYmekcg1Mv8xtf/UkUHr0NrS2NtWR86BysWCe3JTpw7JT5zaNnJVHqWkg6RZv/mBpMFlPtCzj1u43qHRxMHGi1wFmX0IGvq8X5gj5pOzQjs92q72E2ElC/e+GFCdxbrY7egutSNALsb+3MtKmFpy2ZwT2Yhcoour2jcz+OP/Wi9iudzaOyvzyxtGRjm0ZKE6lNFhFOTXczX50f1OAlma3ZjtkBJ38H8j+myK/Q80X4uekPkn+Fh9VDmGYddcMkjG1/TtudqFXfo0yajceWk2moD6v2xHlIqJN38dyL0Wc20Ot3GzAn3uU7lZJKj5ZLr2HKfAG3XdIZT+vlBH9AfROdpVRUIyyKbITZJVCmyKDAVhy8aTbfPGwTRG88MdxD/9H8+/YdXJfI4UYFoH636fYEdhvwK9wRQgPbg8wgsMvj3kbm+5Pdp/1uG+2XW9PC9Al+05o+3nsxLnMcy3uN9i7FYq+xwI4xk6j3d03Y1vR+tBgs2nqnbZp68jmNP8dJeLZo6ITmyoH1Hr4zW+Bs99HIkWtg1C5GfvtDDMivmBR8jFw3flaWmIM2NHLCdYfCnNwWp9yBtlGEKATUoyacJRdmVx8zzYbmDHbZB7vt7yroRufAlKqmE9eWjcNy8y/Nvsw4aqCjrDSUVEUFaERwCdj/jBz1tZCXHuMapXlJLaGYb6VQ2q1K8NSAcY/l0+VoCj3WFRz+9nvzR5BnYdwEU1ssQO5NbVL7fwAAAP//W0ZTyJ8QAAA=",
		"provisionScriptParametersCommon":           fmt.Sprintf("[concat('ADMINUSER=',parameters('linuxAdminUsername'),' ETCD_DOWNLOAD_URL=',parameters('etcdDownloadURLBase'),' ETCD_VERSION=',parameters('etcdVersion'),' CONTAINERD_VERSION=',parameters('containerdVersion'),' MOBY_VERSION=',parameters('mobyVersion'),' TENANT_ID=',variables('tenantID'),' KUBERNETES_VERSION=%s HYPERKUBE_URL=',parameters('kubernetesHyperkubeSpec'),' APISERVER_PUBLIC_KEY=',parameters('apiServerCertificate'),' SUBSCRIPTION_ID=',variables('subscriptionId'),' RESOURCE_GROUP=',variables('resourceGroup'),' LOCATION=',variables('location'),' VM_TYPE=',variables('vmType'),' SUBNET=',variables('subnetName'),' NETWORK_SECURITY_GROUP=',variables('nsgName'),' VIRTUAL_NETWORK=',variables('virtualNetworkName'),' VIRTUAL_NETWORK_RESOURCE_GROUP=',variables('virtualNetworkResourceGroupName'),' ROUTE_TABLE=',variables('routeTableName'),' PRIMARY_AVAILABILITY_SET=',variables('primaryAvailabilitySetName'),' PRIMARY_SCALE_SET=',variables('primaryScaleSetName'),' SERVICE_PRINCIPAL_CLIENT_ID=',variables('servicePrincipalClientId'),' SERVICE_PRINCIPAL_CLIENT_SECRET=',variables('singleQuote'),variables('servicePrincipalClientSecret'),variables('singleQuote'),' KUBELET_PRIVATE_KEY=',parameters('clientPrivateKey'),' TARGET_ENVIRONMENT=',parameters('targetEnvironment'),' NETWORK_PLUGIN=',parameters('networkPlugin'),' NETWORK_POLICY=',parameters('networkPolicy'),' VNET_CNI_PLUGINS_URL=',parameters('vnetCniLinuxPluginsURL'),' CNI_PLUGINS_URL=',parameters('cniPluginsURL'),' CLOUDPROVIDER_BACKOFF=',toLower(string(parameters('cloudproviderConfig').cloudProviderBackoff)),' CLOUDPROVIDER_BACKOFF_RETRIES=',parameters('cloudproviderConfig').cloudProviderBackoffRetries,' CLOUDPROVIDER_BACKOFF_EXPONENT=',parameters('cloudproviderConfig').cloudProviderBackoffExponent,' CLOUDPROVIDER_BACKOFF_DURATION=',parameters('cloudproviderConfig').cloudProviderBackoffDuration,' CLOUDPROVIDER_BACKOFF_JITTER=',parameters('cloudproviderConfig').cloudProviderBackoffJitter,' CLOUDPROVIDER_RATELIMIT=',toLower(string(parameters('cloudproviderConfig').cloudProviderRatelimit)),' CLOUDPROVIDER_RATELIMIT_QPS=',parameters('cloudproviderConfig').cloudProviderRatelimitQPS,' CLOUDPROVIDER_RATELIMIT_BUCKET=',parameters('cloudproviderConfig').cloudProviderRatelimitBucket,' USE_MANAGED_IDENTITY_EXTENSION=',variables('useManagedIdentityExtension'),' USE_INSTANCE_METADATA=',variables('useInstanceMetadata'),' LOAD_BALANCER_SKU=',variables('loadBalancerSku'),' EXCLUDE_MASTER_FROM_STANDARD_LB=',variables('excludeMasterFromStandardLB'),' MAXIMUM_LOADBALANCER_RULE_COUNT=',variables('maximumLoadBalancerRuleCount'),' CONTAINER_RUNTIME=',parameters('containerRuntime'),' CONTAINERD_DOWNLOAD_URL_BASE=',parameters('containerdDownloadURLBase'),' POD_INFRA_CONTAINER_SPEC=',parameters('kubernetesPodInfraContainerSpec'),' KMS_PROVIDER_VAULT_NAME=',variables('clusterKeyVaultName'),' IS_HOSTED_MASTER=false',' PRIVATE_AZURE_REGISTRY_SERVER=',parameters('privateAzureRegistryServer'),' AUTHENTICATION_METHOD=',variables('customCloudAuthenticationMethod'),' IDENTITY_SYSTEM=',variables('customCloudIdentifySystem'))]", customCloudK8sVersion),
		"provisionScriptParametersMaster":           "[concat('COSMOS_URI= MASTER_VM_NAME=',variables('masterVMNames')[variables('masterOffset')],' ETCD_PEER_URL=',variables('masterEtcdPeerURLs')[variables('masterOffset')],' ETCD_CLIENT_URL=',variables('masterEtcdClientURLs')[variables('masterOffset')],' MASTER_NODE=true NO_OUTBOUND=false AUDITD_ENABLED=false CLUSTER_AUTOSCALER_ADDON=',parameters('kubernetesClusterAutoscalerEnabled'),' ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]",
		"readerRoleDefinitionId":                    "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'acdd72a7-3385-48ef-bd42-f606fba81ae7')]",
		"resourceGroup":                             "[resourceGroup().name]",
		"routeTableID":                              "[resourceId('Microsoft.Network/routeTables', variables('routeTableName'))]",
		"routeTableName":                            "[concat(variables('masterVMNamePrefix'),'routetable')]",
		"scope":                                     "[resourceGroup().id]",
		"servicePrincipalClientId":                  "[parameters('servicePrincipalClientId')]",
		"servicePrincipalClientSecret":              "[parameters('servicePrincipalClientSecret')]",
		"singleQuote":                               "'",
		"sshKeyPath":                                "[concat('/home/',parameters('linuxAdminUsername'),'/.ssh/authorized_keys')]",
		"sshNatPorts":                               []int{22, 2201, 2202, 2203, 2204},
		"storageAccountBaseName":                    "",
		"storageAccountPrefixes":                    []interface{}{},
		"subnetName":                                "[concat(parameters('orchestratorName'), '-subnet')]",
		"subnetNameResourceSegmentIndex":            10,
		"subscriptionId":                            "[subscription().subscriptionId]",
		"tenantId":                                  "[subscription().tenantId]",
		"truncatedResourceGroup":                    "[take(replace(replace(resourceGroup().name, '(', '-'), ')', '-'), 63)]",
		"useInstanceMetadata":                       "true",
		"useManagedIdentityExtension":               "false",
		"userAssignedClientID":                      "",
		"userAssignedID":                            "",
		"userAssignedIDReference":                   "[resourceId('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))]",
		"virtualNetworkName":                        "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]",
		"virtualNetworkResourceGroupName":           "''",
		"vmType":                                    "vmss",
		"vnetID":                                    "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]",
		"vnetNameResourceSegmentIndex":              8,
		"vnetResourceGroupNameResourceSegmentIndex": 4,
		"vnetSubnetID":                              "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]",
	}
	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

}

func TestK8sVarsMastersOnly(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:     3,
				DNSPrefix: "blueorange",
				VMSize:    "Standard_D2_v2",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: api.Kubernetes,
				KubernetesConfig: &api.KubernetesConfig{
					LoadBalancerSku:             api.StandardLoadBalancerSku,
					ExcludeMasterFromStandardLB: to.BoolPtr(true),
					NetworkPlugin:               "azure",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
		},
	}

	cs.SetPropertiesDefaults(false, false)

	varMap, err := GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap := map[string]interface{}{
		"apiVersionAuthorizationSystem":   "2018-01-01-preview",
		"apiVersionAuthorizationUser":     "2018-09-01-preview",
		"apiVersionCompute":               "2018-10-01",
		"apiVersionDeployments":           "2018-06-01",
		"apiVersionKeyVault":              "2018-02-14",
		"apiVersionManagedIdentity":       "2015-08-31-preview",
		"apiVersionNetwork":               "2018-08-01",
		"apiVersionStorage":               "2018-07-01",
		"clusterKeyVaultName":             "",
		"contributorRoleDefinitionId":     "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]",
		"customCloudAuthenticationMethod": "client_secret",
		"customCloudIdentifySystem":       "azure_ad",
		"etcdCaFilepath":                  "/etc/kubernetes/certs/ca.crt",
		"etcdClientCertFilepath":          "/etc/kubernetes/certs/etcdclient.crt",
		"etcdClientKeyFilepath":           "/etc/kubernetes/certs/etcdclient.key",
		"etcdPeerCertFilepath":            []string{"/etc/kubernetes/certs/etcdpeer0.crt", "/etc/kubernetes/certs/etcdpeer1.crt", "/etc/kubernetes/certs/etcdpeer2.crt", "/etc/kubernetes/certs/etcdpeer3.crt", "/etc/kubernetes/certs/etcdpeer4.crt"},
		"etcdPeerCertificates":            []string{"[parameters('etcdPeerCertificate0')]", "[parameters('etcdPeerCertificate1')]", "[parameters('etcdPeerCertificate2')]"},
		"etcdPeerKeyFilepath":             []string{"/etc/kubernetes/certs/etcdpeer0.key", "/etc/kubernetes/certs/etcdpeer1.key", "/etc/kubernetes/certs/etcdpeer2.key", "/etc/kubernetes/certs/etcdpeer3.key", "/etc/kubernetes/certs/etcdpeer4.key"},
		"etcdPeerPrivateKeys":             []string{"[parameters('etcdPeerPrivateKey0')]", "[parameters('etcdPeerPrivateKey1')]", "[parameters('etcdPeerPrivateKey2')]"},
		"etcdServerCertFilepath":          "/etc/kubernetes/certs/etcdserver.crt",
		"etcdServerKeyFilepath":           "/etc/kubernetes/certs/etcdserver.key",
		"excludeMasterFromStandardLB":     "true",
		"kubeconfigServer":                "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]",
		"kubernetesAPIServerIP":           "[concat(variables('masterFirstAddrPrefix'), add(variables('masterInternalLbIPOffset'), int(variables('masterFirstAddrOctet4'))))]",
		"labelResourceGroup":              "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"loadBalancerSku":                 api.StandardLoadBalancerSku,
		"location":                        "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"locations":                       []string{"[resourceGroup().location]", "[parameters('location')]"},
		"masterAvailabilitySet":           "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"masterCount":                     3,
		"masterEtcdClientPort":            2379,
		"masterEtcdClientURLs":            []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdClientPort'))]"},
		"masterEtcdClusterStates":         []string{"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2], ',', variables('masterVMNames')[3], '=', variables('masterEtcdPeerURLs')[3], ',', variables('masterVMNames')[4], '=', variables('masterEtcdPeerURLs')[4])]"},
		"masterEtcdPeerURLs":              []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdServerPort'))]"},
		"masterEtcdServerPort":            2380,
		"masterFirstAddrComment":          "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space",
		"masterFirstAddrOctet4":           "[variables('masterFirstAddrOctets')[3]]",
		"masterFirstAddrOctets":           "[split(parameters('firstConsecutiveStaticIP'),'.')]",
		"masterFirstAddrPrefix":           "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
		"masterFqdnPrefix":                "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
		"masterInternalLbID":              "[resourceId('Microsoft.Network/loadBalancers',variables('masterInternalLbName'))]",
		"masterInternalLbIPConfigID":      "[concat(variables('masterInternalLbID'),'/frontendIPConfigurations/', variables('masterInternalLbIPConfigName'))]",
		"masterInternalLbIPConfigName":    "[concat(parameters('orchestratorName'), '-master-internal-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterInternalLbIPOffset":        10,
		"masterInternalLbName":            "[concat(parameters('orchestratorName'), '-master-internal-lb-', parameters('nameSuffix'))]",
		"masterLbBackendPoolName":         "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]",
		"masterLbID":                      "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
		"masterLbIPConfigID":              "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
		"masterLbIPConfigName":            "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterLbName":                    "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]",
		"masterOffset":                    "[parameters('masterOffset')]",
		"masterPrivateIpAddrs":            []string{"[concat(variables('masterFirstAddrPrefix'), add(0, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(1, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(2, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(3, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(4, int(variables('masterFirstAddrOctet4'))))]"},
		"masterPublicIPAddressName":       "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]",
		"masterVMNamePrefix":              "k8s-master-18280257-",
		"masterVMNames":                   []string{"[concat(variables('masterVMNamePrefix'), '0')]", "[concat(variables('masterVMNamePrefix'), '1')]", "[concat(variables('masterVMNamePrefix'), '2')]", "[concat(variables('masterVMNamePrefix'), '3')]", "[concat(variables('masterVMNamePrefix'), '4')]"},
		"maxVMsPerPool":                   100,
		"maximumLoadBalancerRuleCount":    250,
		"nsgID":                           "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]",
		"nsgName":                         "[concat(variables('masterVMNamePrefix'), 'nsg')]",
		"orchestratorNameVersionTag":      "Kubernetes:" + testK8sVersion,
		"primaryAvailabilitySetName":      "",
		"primaryScaleSetName":             "",
		"cloudInitFiles": map[string]interface{}{
			"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript),
			"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript),
			"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall),
			"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig),
			"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript),
			"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript),
			"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript),
			"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript),
			"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd),
			"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService),
			"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService),
			"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService),
			"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer),
			"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService),
			"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences),
			"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags),
			"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService),
		},
		"provisionScriptParametersCommon":           fmt.Sprintf("[concat('ADMINUSER=',parameters('linuxAdminUsername'),' ETCD_DOWNLOAD_URL=',parameters('etcdDownloadURLBase'),' ETCD_VERSION=',parameters('etcdVersion'),' CONTAINERD_VERSION=',parameters('containerdVersion'),' MOBY_VERSION=',parameters('mobyVersion'),' TENANT_ID=',variables('tenantID'),' KUBERNETES_VERSION=%s HYPERKUBE_URL=',parameters('kubernetesHyperkubeSpec'),' APISERVER_PUBLIC_KEY=',parameters('apiServerCertificate'),' SUBSCRIPTION_ID=',variables('subscriptionId'),' RESOURCE_GROUP=',variables('resourceGroup'),' LOCATION=',variables('location'),' VM_TYPE=',variables('vmType'),' SUBNET=',variables('subnetName'),' NETWORK_SECURITY_GROUP=',variables('nsgName'),' VIRTUAL_NETWORK=',variables('virtualNetworkName'),' VIRTUAL_NETWORK_RESOURCE_GROUP=',variables('virtualNetworkResourceGroupName'),' ROUTE_TABLE=',variables('routeTableName'),' PRIMARY_AVAILABILITY_SET=',variables('primaryAvailabilitySetName'),' PRIMARY_SCALE_SET=',variables('primaryScaleSetName'),' SERVICE_PRINCIPAL_CLIENT_ID=',variables('servicePrincipalClientId'),' SERVICE_PRINCIPAL_CLIENT_SECRET=',variables('singleQuote'),variables('servicePrincipalClientSecret'),variables('singleQuote'),' KUBELET_PRIVATE_KEY=',parameters('clientPrivateKey'),' TARGET_ENVIRONMENT=',parameters('targetEnvironment'),' NETWORK_PLUGIN=',parameters('networkPlugin'),' NETWORK_POLICY=',parameters('networkPolicy'),' VNET_CNI_PLUGINS_URL=',parameters('vnetCniLinuxPluginsURL'),' CNI_PLUGINS_URL=',parameters('cniPluginsURL'),' CLOUDPROVIDER_BACKOFF=',toLower(string(parameters('cloudproviderConfig').cloudProviderBackoff)),' CLOUDPROVIDER_BACKOFF_RETRIES=',parameters('cloudproviderConfig').cloudProviderBackoffRetries,' CLOUDPROVIDER_BACKOFF_EXPONENT=',parameters('cloudproviderConfig').cloudProviderBackoffExponent,' CLOUDPROVIDER_BACKOFF_DURATION=',parameters('cloudproviderConfig').cloudProviderBackoffDuration,' CLOUDPROVIDER_BACKOFF_JITTER=',parameters('cloudproviderConfig').cloudProviderBackoffJitter,' CLOUDPROVIDER_RATELIMIT=',toLower(string(parameters('cloudproviderConfig').cloudProviderRatelimit)),' CLOUDPROVIDER_RATELIMIT_QPS=',parameters('cloudproviderConfig').cloudProviderRatelimitQPS,' CLOUDPROVIDER_RATELIMIT_BUCKET=',parameters('cloudproviderConfig').cloudProviderRatelimitBucket,' USE_MANAGED_IDENTITY_EXTENSION=',variables('useManagedIdentityExtension'),' USE_INSTANCE_METADATA=',variables('useInstanceMetadata'),' LOAD_BALANCER_SKU=',variables('loadBalancerSku'),' EXCLUDE_MASTER_FROM_STANDARD_LB=',variables('excludeMasterFromStandardLB'),' MAXIMUM_LOADBALANCER_RULE_COUNT=',variables('maximumLoadBalancerRuleCount'),' CONTAINER_RUNTIME=',parameters('containerRuntime'),' CONTAINERD_DOWNLOAD_URL_BASE=',parameters('containerdDownloadURLBase'),' POD_INFRA_CONTAINER_SPEC=',parameters('kubernetesPodInfraContainerSpec'),' KMS_PROVIDER_VAULT_NAME=',variables('clusterKeyVaultName'),' IS_HOSTED_MASTER=false',' PRIVATE_AZURE_REGISTRY_SERVER=',parameters('privateAzureRegistryServer'),' AUTHENTICATION_METHOD=',variables('customCloudAuthenticationMethod'),' IDENTITY_SYSTEM=',variables('customCloudIdentifySystem'))]", testK8sVersion),
		"provisionScriptParametersMaster":           "[concat('COSMOS_URI= MASTER_VM_NAME=',variables('masterVMNames')[variables('masterOffset')],' ETCD_PEER_URL=',variables('masterEtcdPeerURLs')[variables('masterOffset')],' ETCD_CLIENT_URL=',variables('masterEtcdClientURLs')[variables('masterOffset')],' MASTER_NODE=true NO_OUTBOUND=false AUDITD_ENABLED=false CLUSTER_AUTOSCALER_ADDON=',parameters('kubernetesClusterAutoscalerEnabled'),' ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]",
		"readerRoleDefinitionId":                    "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'acdd72a7-3385-48ef-bd42-f606fba81ae7')]",
		"resourceGroup":                             "[resourceGroup().name]",
		"routeTableID":                              "[resourceId('Microsoft.Network/routeTables', variables('routeTableName'))]",
		"routeTableName":                            "[concat(variables('masterVMNamePrefix'),'routetable')]",
		"scope":                                     "[resourceGroup().id]",
		"servicePrincipalClientId":                  "[parameters('servicePrincipalClientId')]",
		"servicePrincipalClientSecret":              "[parameters('servicePrincipalClientSecret')]",
		"singleQuote":                               "'",
		"sshKeyPath":                                "[concat('/home/',parameters('linuxAdminUsername'),'/.ssh/authorized_keys')]",
		"sshNatPorts":                               []int{22, 2201, 2202, 2203, 2204},
		"storageAccountBaseName":                    "",
		"storageAccountPrefixes":                    []interface{}{},
		"subnetName":                                "[concat(parameters('orchestratorName'), '-subnet')]",
		"subnetNameResourceSegmentIndex":            10,
		"subscriptionId":                            "[subscription().subscriptionId]",
		"tenantId":                                  "[subscription().tenantId]",
		"truncatedResourceGroup":                    "[take(replace(replace(resourceGroup().name, '(', '-'), ')', '-'), 63)]",
		"useInstanceMetadata":                       "true",
		"useManagedIdentityExtension":               "false",
		"userAssignedClientID":                      "",
		"userAssignedID":                            "",
		"userAssignedIDReference":                   "[resourceId('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))]",
		"virtualNetworkName":                        "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]",
		"virtualNetworkResourceGroupName":           "''",
		"vmType":                                    "standard",
		"vnetID":                                    "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]",
		"vnetNameResourceSegmentIndex":              8,
		"vnetResourceGroupNameResourceSegmentIndex": 4,
		"vnetSubnetID":                              "[concat(variables('vnetID'),'/subnets/',variables('subnetName'))]",
	}
	diff := cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}
