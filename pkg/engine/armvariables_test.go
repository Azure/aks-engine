// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
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
				KubernetesConfig: &api.KubernetesConfig{
					LoadBalancerSku: api.BasicLoadBalancerSku,
				},
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

	_, err := cs.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		t.Errorf("expected no error from SetPropertiesDefaults, instead got %s", err)
	}

	varMap, err := GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap := map[string]interface{}{
		"agentpool1Count":                    "[parameters('agentpool1Count')]",
		"agentpool1Index":                    0,
		"agentpool1SubnetName":               "[variables('subnetName')]",
		"agentpool1SubnetResourceGroup":      "[split(variables('agentpool1VnetSubnetID'), '/')[4]]",
		"agentpool1VMNamePrefix":             "k8s-agentpool1-18280257-vmss",
		"agentpool1VMSize":                   "[parameters('agentpool1VMSize')]",
		"agentpool1Vnet":                     "[split(variables('agentpool1VnetSubnetID'), '/')[8]]",
		"agentpool1VnetSubnetID":             "[variables('vnetSubnetID')]",
		"agentpool1osImageName":              "[parameters('agentpool1osImageName')]",
		"agentpool1osImageOffer":             "[parameters('agentpool1osImageOffer')]",
		"agentpool1osImagePublisher":         "[parameters('agentpool1osImagePublisher')]",
		"agentpool1osImageResourceGroup":     "[parameters('agentpool1osImageResourceGroup')]",
		"agentpool1osImageSKU":               "[parameters('agentpool1osImageSKU')]",
		"agentpool1osImageVersion":           "[parameters('agentpool1osImageVersion')]",
		"apiVersionAuthorizationSystem":      "2018-01-01-preview",
		"apiVersionAuthorizationUser":        "2018-09-01-preview",
		"apiVersionCompute":                  "2019-07-01",
		"apiVersionDeployments":              "2018-06-01",
		"apiVersionKeyVault":                 "2018-02-14",
		"apiVersionManagedIdentity":          "2018-11-30",
		"apiVersionNetwork":                  "2018-08-01",
		"apiVersionStorage":                  "2018-07-01",
		"applicationInsightsKey":             "c92d8284-b550-4b06-b7ba-e80fd7178faa", // should be DefaultApplicationInsightsKey,
		"clusterKeyVaultName":                "",
		"contributorRoleDefinitionId":        "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]",
		"enableTelemetry":                    false,
		"etcdCaFilepath":                     "/etc/kubernetes/certs/ca.crt",
		"etcdClientCertFilepath":             "/etc/kubernetes/certs/etcdclient.crt",
		"etcdClientKeyFilepath":              "/etc/kubernetes/certs/etcdclient.key",
		"etcdPeerCertFilepath":               []string{"/etc/kubernetes/certs/etcdpeer0.crt", "/etc/kubernetes/certs/etcdpeer1.crt", "/etc/kubernetes/certs/etcdpeer2.crt", "/etc/kubernetes/certs/etcdpeer3.crt", "/etc/kubernetes/certs/etcdpeer4.crt"},
		"etcdPeerCertificates":               []string{"[parameters('etcdPeerCertificate0')]"},
		"etcdPeerKeyFilepath":                []string{"/etc/kubernetes/certs/etcdpeer0.key", "/etc/kubernetes/certs/etcdpeer1.key", "/etc/kubernetes/certs/etcdpeer2.key", "/etc/kubernetes/certs/etcdpeer3.key", "/etc/kubernetes/certs/etcdpeer4.key"},
		"etcdPeerPrivateKeys":                []string{"[parameters('etcdPeerPrivateKey0')]"},
		"etcdServerCertFilepath":             "/etc/kubernetes/certs/etcdserver.crt",
		"etcdServerKeyFilepath":              "/etc/kubernetes/certs/etcdserver.key",
		"excludeMasterFromStandardLB":        "false",
		"kubeconfigServer":                   "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]",
		"kubernetesAPIServerIP":              "[parameters('firstConsecutiveStaticIP')]",
		"labelResourceGroup":                 "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"loadBalancerSku":                    BasicLoadBalancerSku,
		"location":                           "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"locations":                          []string{"[resourceGroup().location]", "[parameters('location')]"},
		"masterAvailabilitySet":              "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"masterCount":                        1,
		"masterEtcdClientPort":               2379,
		"masterEtcdClientURLs":               []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdClientPort'))]"},
		"masterEtcdClusterStates":            []string{"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2], ',', variables('masterVMNames')[3], '=', variables('masterEtcdPeerURLs')[3], ',', variables('masterVMNames')[4], '=', variables('masterEtcdPeerURLs')[4])]"},
		"masterEtcdPeerURLs":                 []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdServerPort'))]"},
		"masterEtcdServerPort":               2380,
		"masterEtcdMetricURLs":               []string{"[concat('http://', variables('masterPrivateIpAddrs')[0], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[1], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[2], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[3], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[4], ':2480')]"},
		"masterFirstAddrComment":             "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space",
		"masterFirstAddrOctet4":              "[variables('masterFirstAddrOctets')[3]]",
		"masterFirstAddrOctets":              "[split(parameters('firstConsecutiveStaticIP'),'.')]",
		"masterFirstAddrPrefix":              "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
		"masterFqdnPrefix":                   "blueorange",
		"masterLbBackendPoolName":            "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]",
		"masterLbID":                         "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
		"masterLbIPConfigID":                 "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
		"masterLbIPConfigName":               "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterLbName":                       "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]",
		"masterOffset":                       "[parameters('masterOffset')]",
		"masterPrivateIpAddrs":               []string{"[concat(variables('masterFirstAddrPrefix'), add(0, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(1, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(2, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(3, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(4, int(variables('masterFirstAddrOctet4'))))]"},
		"masterPublicIPAddressName":          "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]",
		"masterVMNamePrefix":                 "k8s-master-18280257-",
		"masterVMNames":                      []string{"[concat(variables('masterVMNamePrefix'), '0')]", "[concat(variables('masterVMNamePrefix'), '1')]", "[concat(variables('masterVMNamePrefix'), '2')]", "[concat(variables('masterVMNamePrefix'), '3')]", "[concat(variables('masterVMNamePrefix'), '4')]"},
		"maxVMsPerPool":                      100,
		"maximumLoadBalancerRuleCount":       250,
		"networkContributorRoleDefinitionId": "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', '4d97b98b-1d4f-4787-a291-c67834d212e7')]",
		"nsgID":                              "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]",
		"nsgName":                            "[concat(variables('masterVMNamePrefix'), 'nsg')]",
		"orchestratorNameVersionTag":         "Kubernetes:" + testK8sVersion,
		"primaryAvailabilitySetName":         "",
		"primaryScaleSetName":                cs.Properties.GetPrimaryScaleSetName(),
		"cloudInitFiles": map[string]interface{}{
			"provisionScript":           getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
			"provisionSource":           getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
			"provisionInstalls":         getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
			"provisionConfigs":          getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
			"customSearchDomainsScript": getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
			"generateProxyCertsScript":  getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
			"mountEtcdScript":           getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
			"etcdSystemdService":        getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
			"dhcpv6ConfigurationScript": getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
			"dhcpv6SystemdService":      getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
			"kubeletSystemdService":     getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		},
		"provisionScriptParametersCommon":           "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]",
		"provisionScriptParametersMaster":           "[concat('COSMOS_URI= MASTER_VM_NAME=',variables('masterVMNames')[variables('masterOffset')],' ETCD_PEER_URL=',variables('masterEtcdPeerURLs')[variables('masterOffset')],' ETCD_CLIENT_URL=',variables('masterEtcdClientURLs')[variables('masterOffset')],' MASTER_NODE=true NO_OUTBOUND=false AUDITD_ENABLED=false CLUSTER_AUTOSCALER_ADDON=false ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]",
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
		"subnetName":                                "k8s-subnet",
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
		"windowsCSIProxyURL":                        "",
		"windowsEnableCSIProxy":                     false,
	}

	diff := cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with AAD Pod Identity
	cs.Properties.OrchestratorProfile.KubernetesConfig.Addons = []api.KubernetesAddon{
		{
			Name:    common.AADPodIdentityAddonName,
			Enabled: to.BoolPtr(true),
		},
	}
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":            getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
		"provisionSource":            getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
		"provisionInstalls":          getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
		"provisionConfigs":           getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
		"customSearchDomainsScript":  getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
		"generateProxyCertsScript":   getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
		"mountEtcdScript":            getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
		"etcdSystemdService":         getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
		"dhcpv6ConfigurationScript":  getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
		"dhcpv6SystemdService":       getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
		"kubeletSystemdService":      getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		"untaintNodesScript":         getBase64EncodedGzippedCustomScript(untaintNodesScript, cs),
		"untaintNodesSystemdService": getBase64EncodedGzippedCustomScript(untaintNodesSystemdService, cs),
	}

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with MSI
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["servicePrincipalClientId"] = "msi"
	expectedMap["servicePrincipalClientSecret"] = "msi"
	expectedMap["useManagedIdentityExtension"] = "true"
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with ubuntu 16.04 distro
	cs.Properties.OrchestratorProfile.KubernetesConfig.Addons = []api.KubernetesAddon{}
	cs.Properties.AgentPoolProfiles[0].Distro = api.Ubuntu
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = false
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"
	expectedMap["servicePrincipalClientId"] = "[parameters('servicePrincipalClientId')]"
	expectedMap["servicePrincipalClientSecret"] = "[parameters('servicePrincipalClientSecret')]"
	expectedMap["useManagedIdentityExtension"] = "false"
	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
		"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
		"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
		"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
		"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript, cs),
		"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript, cs),
		"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
		"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
		"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
		"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService, cs),
		"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService, cs),
		"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer, cs),
		"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService, cs),
		"labelNodesScript":                 getBase64EncodedGzippedCustomScript(labelNodesScript, cs),
		"labelNodesSystemdService":         getBase64EncodedGzippedCustomScript(labelNodesSystemdService, cs),
		"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences, cs),
		"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags, cs),
		"auditdRules":                      getBase64EncodedGzippedCustomScript(auditdRules, cs),
		"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
		"dhcpv6ConfigurationScript":        getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
		"dhcpv6SystemdService":             getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
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

	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with ubuntu 18.04 gen2 distro
	cs.Properties.AgentPoolProfiles[0].Distro = api.Ubuntu1804Gen2
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with CustomVnet enabled
	cs.Properties.MasterProfile.VnetSubnetID = "/subscriptions/fakesubID/resourceGroups/myRG/providers/Microsoft.Network/virtualNetworks/fooSubnetID/subnets/myCustomSubnet"
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["subnetName"] = "myCustomSubnet"
	expectedMap["virtualNetworkName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetNameResourceSegmentIndex')]]"
	expectedMap["virtualNetworkResourceGroupName"] = "[split(parameters('masterVnetSubnetID'), '/')[variables('vnetResourceGroupNameResourceSegmentIndex')]]"
	expectedMap["vnetSubnetID"] = "[parameters('masterVnetSubnetID')]"
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"
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
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"

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
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"

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
	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
		"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
		"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
		"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
		"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript, cs),
		"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript, cs),
		"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
		"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
		"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
		"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService, cs),
		"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService, cs),
		"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer, cs),
		"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService, cs),
		"labelNodesScript":                 getBase64EncodedGzippedCustomScript(labelNodesScript, cs),
		"labelNodesSystemdService":         getBase64EncodedGzippedCustomScript(labelNodesSystemdService, cs),
		"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences, cs),
		"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags, cs),
		"auditdRules":                      getBase64EncodedGzippedCustomScript(auditdRules, cs),
		"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
		"dhcpv6ConfigurationScript":        getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
		"dhcpv6SystemdService":             getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
	}
	expectedMap["agentNamePrefix"] = "[concat(parameters('orchestratorName'), '-agentpool-', parameters('nameSuffix'), '-')]"
	expectedMap["agentpool1AccountName"] = "[concat(variables('storageAccountBaseName'), 'agnt0')]"
	expectedMap["agentpool1StorageAccountOffset"] = "[mul(variables('maxStorageAccountsPerAgent'),variables('agentpool1Index'))]"
	expectedMap["agentpool1StorageAccountsCount"] = "[add(div(variables('agentpool1Count'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('agentpool1Count'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('agentpool1Count'), variables('maxVMsPerStorageAccount')),1)))]"
	expectedMap["agentpool1VMNamePrefix"] = "aks-agentpool1-18280257-vmss"
	expectedMap["dataStorageAccountPrefixSeed"] = 97
	expectedMap["kubernetesAPIServerIP"] = "[parameters('kubernetesEndpoint')]"
	expectedMap["masterFqdnPrefix"] = "foodnsprefix"
	expectedMap["masterVMNamePrefix"] = "aks-master-18280257-"
	expectedMap["maxStorageAccountsPerAgent"] = "[div(variables('maxVMsPerPool'),variables('maxVMsPerStorageAccount'))]"
	expectedMap["maxVMsPerStorageAccount"] = 20
	expectedMap["nsgName"] = "[concat(variables('agentNamePrefix'), 'nsg')]"
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"
	expectedMap["routeTableName"] = "[concat(variables('agentNamePrefix'), 'routetable')]"
	expectedMap["storageAccountBaseName"] = "[uniqueString(concat(variables('masterFqdnPrefix'),variables('location')))]"
	expectedMap["storageAccountPrefixes"] = []string{"0", "6", "c", "i", "o", "u", "1", "7", "d", "j", "p", "v", "2", "8", "e", "k", "q", "w", "3", "9", "f", "l", "r", "x", "4", "a", "g", "m", "s", "y", "5", "b", "h", "n", "t", "z"}
	expectedMap["storageAccountPrefixesCount"] = "[length(variables('storageAccountPrefixes'))]"
	expectedMap["subnetName"] = "aks-subnet"
	expectedMap["virtualNetworkName"] = "[concat(parameters('orchestratorName'), '-vnet-', parameters('nameSuffix'))]"
	expectedMap["virtualNetworkResourceGroupName"] = ""
	expectedMap["vmSizesMap"] = getSizeMap()["vmSizesMap"]
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
	delete(expectedMap, "masterEtcdMetricURLs")

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with HostedMaster + MSI
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
		"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
		"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
		"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
		"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript, cs),
		"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript, cs),
		"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
		"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
		"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
		"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService, cs),
		"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService, cs),
		"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer, cs),
		"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService, cs),
		"labelNodesScript":                 getBase64EncodedGzippedCustomScript(labelNodesScript, cs),
		"labelNodesSystemdService":         getBase64EncodedGzippedCustomScript(labelNodesSystemdService, cs),
		"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences, cs),
		"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags, cs),
		"auditdRules":                      getBase64EncodedGzippedCustomScript(auditdRules, cs),
		"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
		"dhcpv6ConfigurationScript":        getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
		"dhcpv6SystemdService":             getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
	}
	expectedMap["useManagedIdentityExtension"] = "true"
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with IPv6 DualStack feature enabled
	cs.Properties.FeatureFlags = &api.FeatureFlags{EnableIPv6DualStack: true}
	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"
	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":                  getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
		"provisionSource":                  getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
		"provisionInstalls":                getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
		"provisionConfigs":                 getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
		"provisionCIS":                     getBase64EncodedGzippedCustomScript(kubernetesCISScript, cs),
		"healthMonitorScript":              getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript, cs),
		"customSearchDomainsScript":        getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
		"generateProxyCertsScript":         getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
		"mountEtcdScript":                  getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
		"kubeletSystemdService":            getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		"kmsSystemdService":                getBase64EncodedGzippedCustomScript(kmsSystemdService, cs),
		"kubeletMonitorSystemdService":     getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService, cs),
		"dockerMonitorSystemdTimer":        getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer, cs),
		"dockerMonitorSystemdService":      getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService, cs),
		"labelNodesScript":                 getBase64EncodedGzippedCustomScript(labelNodesScript, cs),
		"labelNodesSystemdService":         getBase64EncodedGzippedCustomScript(labelNodesSystemdService, cs),
		"aptPreferences":                   getBase64EncodedGzippedCustomScript(aptPreferences, cs),
		"dockerClearMountPropagationFlags": getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags, cs),
		"auditdRules":                      getBase64EncodedGzippedCustomScript(auditdRules, cs),
		"etcdSystemdService":               getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
		"dhcpv6ConfigurationScript":        getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
		"dhcpv6SystemdService":             getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
	}
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

	_, err = cs.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		t.Errorf("expected no error from SetPropertiesDefaults, instead got %s", err)
	}

	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap = map[string]interface{}{
		"agentpool1Count":                    "[parameters('agentpool1Count')]",
		"agentpool1Index":                    0,
		"agentpool1SubnetName":               "[variables('subnetName')]",
		"agentpool1SubnetResourceGroup":      "[split(variables('agentpool1VnetSubnetID'), '/')[4]]",
		"agentpool1VMNamePrefix":             "k8s-agentpool1-18280257-vmss",
		"agentpool1VMSize":                   "[parameters('agentpool1VMSize')]",
		"agentpool1Vnet":                     "[split(variables('agentpool1VnetSubnetID'), '/')[8]]",
		"agentpool1VnetSubnetID":             "[variables('vnetSubnetID')]",
		"agentpool1osImageName":              "[parameters('agentpool1osImageName')]",
		"agentpool1osImageOffer":             "[parameters('agentpool1osImageOffer')]",
		"agentpool1osImagePublisher":         "[parameters('agentpool1osImagePublisher')]",
		"agentpool1osImageResourceGroup":     "[parameters('agentpool1osImageResourceGroup')]",
		"agentpool1osImageSKU":               "[parameters('agentpool1osImageSKU')]",
		"agentpool1osImageVersion":           "[parameters('agentpool1osImageVersion')]",
		"apiVersionAuthorizationSystem":      "2018-01-01-preview",
		"apiVersionAuthorizationUser":        "2018-09-01-preview",
		"apiVersionCompute":                  "2017-03-30",
		"apiVersionDeployments":              "2018-06-01",
		"apiVersionKeyVault":                 "2016-10-01",
		"applicationInsightsKey":             "c92d8284-b550-4b06-b7ba-e80fd7178faa", // should be DefaultApplicationInsightsKey,
		"environmentJSON":                    `{"name":"azurestackcloud","managementPortalURL":"https://management.local.azurestack.external/","publishSettingsURL":"https://management.local.azurestack.external/publishsettings/index","serviceManagementEndpoint":"https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189","resourceManagerEndpoint":"https://management.local.azurestack.external/","activeDirectoryEndpoint":"https://login.windows.net/","galleryEndpoint":"https://portal.local.azurestack.external=30015/","keyVaultEndpoint":"https://vault.azurestack.external/","graphEndpoint":"https://graph.windows.net/","serviceBusEndpoint":"https://servicebus.azurestack.external/","batchManagementEndpoint":"https://batch.azurestack.external/","storageEndpointSuffix":"core.azurestack.external","sqlDatabaseDNSSuffix":"database.azurestack.external","trafficManagerDNSSuffix":"trafficmanager.cn","keyVaultDNSSuffix":"vault.azurestack.external","serviceBusEndpointSuffix":"servicebus.azurestack.external","serviceManagementVMDNSSuffix":"chinacloudapp.cn","resourceManagerVMDNSSuffix":"cloudapp.azurestack.external","containerRegistryDNSSuffix":"azurecr.io","cosmosDBDNSSuffix":"","tokenAudience":"https://management.azurestack.external/","resourceIdentifiers":{"graph":"","keyVault":"","datalake":"","batch":"","operationalInsights":"","storage":""}}`,
		"customCloudAuthenticationMethod":    "client_secret",
		"customCloudIdentifySystem":          "azure_ad",
		"apiVersionManagedIdentity":          "2018-11-30",
		"apiVersionNetwork":                  "2017-10-01",
		"apiVersionStorage":                  "2017-10-01",
		"clusterKeyVaultName":                "",
		"contributorRoleDefinitionId":        "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]",
		"enableTelemetry":                    false,
		"etcdCaFilepath":                     "/etc/kubernetes/certs/ca.crt",
		"etcdClientCertFilepath":             "/etc/kubernetes/certs/etcdclient.crt",
		"etcdClientKeyFilepath":              "/etc/kubernetes/certs/etcdclient.key",
		"etcdPeerCertFilepath":               []string{"/etc/kubernetes/certs/etcdpeer0.crt", "/etc/kubernetes/certs/etcdpeer1.crt", "/etc/kubernetes/certs/etcdpeer2.crt", "/etc/kubernetes/certs/etcdpeer3.crt", "/etc/kubernetes/certs/etcdpeer4.crt"},
		"etcdPeerCertificates":               []string{"[parameters('etcdPeerCertificate0')]"},
		"etcdPeerKeyFilepath":                []string{"/etc/kubernetes/certs/etcdpeer0.key", "/etc/kubernetes/certs/etcdpeer1.key", "/etc/kubernetes/certs/etcdpeer2.key", "/etc/kubernetes/certs/etcdpeer3.key", "/etc/kubernetes/certs/etcdpeer4.key"},
		"etcdPeerPrivateKeys":                []string{"[parameters('etcdPeerPrivateKey0')]"},
		"etcdServerCertFilepath":             "/etc/kubernetes/certs/etcdserver.crt",
		"etcdServerKeyFilepath":              "/etc/kubernetes/certs/etcdserver.key",
		"excludeMasterFromStandardLB":        "false",
		"kubeconfigServer":                   "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]",
		"kubernetesAPIServerIP":              "[parameters('firstConsecutiveStaticIP')]",
		"labelResourceGroup":                 "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"loadBalancerSku":                    BasicLoadBalancerSku,
		"location":                           "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"locations":                          []string{"[resourceGroup().location]", "[parameters('location')]"},
		"masterAvailabilitySet":              "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"masterCount":                        1,
		"masterEtcdClientPort":               2379,
		"masterEtcdClientURLs":               []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdClientPort'))]"},
		"masterEtcdClusterStates":            []string{"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2], ',', variables('masterVMNames')[3], '=', variables('masterEtcdPeerURLs')[3], ',', variables('masterVMNames')[4], '=', variables('masterEtcdPeerURLs')[4])]"},
		"masterEtcdPeerURLs":                 []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdServerPort'))]"},
		"masterEtcdServerPort":               2380,
		"masterEtcdMetricURLs":               []string{"[concat('http://', variables('masterPrivateIpAddrs')[0], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[1], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[2], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[3], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[4], ':2480')]"},
		"masterFirstAddrComment":             "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space",
		"masterFirstAddrOctet4":              "[variables('masterFirstAddrOctets')[3]]",
		"masterFirstAddrOctets":              "[split(parameters('firstConsecutiveStaticIP'),'.')]",
		"masterFirstAddrPrefix":              "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
		"masterFqdnPrefix":                   "blueorange",
		"masterLbBackendPoolName":            "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]",
		"masterLbID":                         "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
		"masterLbIPConfigID":                 "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
		"masterLbIPConfigName":               "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterLbName":                       "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]",
		"masterOffset":                       "[parameters('masterOffset')]",
		"masterPrivateIpAddrs":               []string{"[concat(variables('masterFirstAddrPrefix'), add(0, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(1, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(2, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(3, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(4, int(variables('masterFirstAddrOctet4'))))]"},
		"masterPublicIPAddressName":          "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]",
		"masterVMNamePrefix":                 "k8s-master-18280257-",
		"masterVMNames":                      []string{"[concat(variables('masterVMNamePrefix'), '0')]", "[concat(variables('masterVMNamePrefix'), '1')]", "[concat(variables('masterVMNamePrefix'), '2')]", "[concat(variables('masterVMNamePrefix'), '3')]", "[concat(variables('masterVMNamePrefix'), '4')]"},
		"maxVMsPerPool":                      100,
		"maximumLoadBalancerRuleCount":       250,
		"networkContributorRoleDefinitionId": "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', '4d97b98b-1d4f-4787-a291-c67834d212e7')]",
		"nsgID":                              "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]",
		"nsgName":                            "[concat(variables('masterVMNamePrefix'), 'nsg')]",
		"orchestratorNameVersionTag":         "Kubernetes:" + testK8sVersion,
		"primaryAvailabilitySetName":         "",
		"primaryScaleSetName":                cs.Properties.GetPrimaryScaleSetName(),
		"cloudInitFiles": map[string]interface{}{
			"provisionScript":           getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
			"provisionSource":           getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
			"provisionInstalls":         getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
			"provisionConfigs":          getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
			"customSearchDomainsScript": getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
			"generateProxyCertsScript":  getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
			"mountEtcdScript":           getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
			"etcdSystemdService":        getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
			"dhcpv6ConfigurationScript": getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
			"dhcpv6SystemdService":      getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
			"kubeletSystemdService":     getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		},
		"provisionConfigsCustomCloud":               getBase64EncodedGzippedCustomScript(kubernetesCSECustomCloud, cs),
		"provisionScriptParametersCommon":           "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]",
		"provisionScriptParametersMaster":           "[concat('COSMOS_URI= MASTER_VM_NAME=',variables('masterVMNames')[variables('masterOffset')],' ETCD_PEER_URL=',variables('masterEtcdPeerURLs')[variables('masterOffset')],' ETCD_CLIENT_URL=',variables('masterEtcdClientURLs')[variables('masterOffset')],' MASTER_NODE=true NO_OUTBOUND=false AUDITD_ENABLED=false CLUSTER_AUTOSCALER_ADDON=false ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]",
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
		"subnetName":                                "k8s-subnet",
		"subscriptionId":                            "[subscription().subscriptionId]",
		"tenantId":                                  "[subscription().tenantId]",
		"truncatedResourceGroup":                    "[take(replace(replace(resourceGroup().name, '(', '-'), ')', '-'), 63)]",
		"useInstanceMetadata":                       "false",
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
		"windowsCSIProxyURL":                        "",
		"windowsEnableCSIProxy":                     false,
	}
	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	cs.Properties.OrchestratorProfile.KubernetesConfig.Addons = []api.KubernetesAddon{
		{
			Name:    common.AppGwIngressAddonName,
			Enabled: to.BoolPtr(true),
			Config: map[string]string{
				"appgw-sku": "WAF_v2",
			},
		},
	}

	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["managedIdentityOperatorRoleDefinitionId"] = "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'f1a07417-d97a-45cb-824c-7a7467783830')]"
	expectedMap["appGwName"] = "[concat(parameters('orchestratorName'), '-appgw-', parameters('nameSuffix'))]"
	expectedMap["appGwSubnetName"] = "appgw-subnet"
	expectedMap["appGwPublicIPAddressName"] = "[concat(parameters('orchestratorName'), '-appgw-ip-', parameters('nameSuffix'))]"
	expectedMap["appGwICIdentityName"] = "[concat(parameters('orchestratorName'), '-appgw-ic-identity-', parameters('nameSuffix'))]"
	expectedMap["appGwId"] = "[resourceId('Microsoft.Network/applicationGateways',variables('appGwName'))]"
	expectedMap["appGwICIdentityId"] = "[resourceId('Microsoft.ManagedIdentity/userAssignedIdentities', variables('appGwICIdentityName'))]"
	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with SLB, should generate agentLb resource variables
	cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku = api.StandardLoadBalancerSku

	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}
	expectedMap["agentPublicIPAddressName"] = "[concat(parameters('orchestratorName'), '-agent-ip-outbound')]"
	expectedMap["agentLbID"] = "[resourceId('Microsoft.Network/loadBalancers',variables('agentLbName'))]"
	expectedMap["agentLbIPConfigID"] = "[concat(variables('agentLbID'),'/frontendIPConfigurations/', variables('agentLbIPConfigName'))]"
	expectedMap["agentLbIPConfigName"] = "[concat(parameters('orchestratorName'), '-agent-outbound')]"
	expectedMap["agentLbName"] = "[parameters('masterEndpointDNSNamePrefix')]"
	expectedMap["agentLbBackendPoolName"] = "[parameters('masterEndpointDNSNamePrefix')]"
	expectedMap["loadBalancerSku"] = api.StandardLoadBalancerSku
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"

	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with cilium
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginCilium

	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap["cloudInitFiles"] = map[string]interface{}{
		"provisionScript":           getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
		"provisionSource":           getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
		"provisionInstalls":         getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
		"provisionConfigs":          getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
		"customSearchDomainsScript": getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
		"generateProxyCertsScript":  getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
		"mountEtcdScript":           getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
		"etcdSystemdService":        getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
		"dhcpv6ConfigurationScript": getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
		"dhcpv6SystemdService":      getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
		"kubeletSystemdService":     getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		"systemdBPFMount":           getBase64EncodedGzippedCustomScript(systemdBPFMount, cs),
	}
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"
	diff = cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with Spot VMs
	cs.Properties.AgentPoolProfiles[0].AvailabilityProfile = api.VirtualMachineScaleSets
	cs.Properties.AgentPoolProfiles[0].ScaleSetPriority = api.ScaleSetPrioritySpot

	varMap, err = GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	agentPoolName := cs.Properties.AgentPoolProfiles[0].Name
	expectedMap[fmt.Sprintf("%sScaleSetPriority", agentPoolName)] = fmt.Sprintf("[parameters('%sScaleSetPriority')]", agentPoolName)
	expectedMap[fmt.Sprintf("%sScaleSetEvictionPolicy", agentPoolName)] = fmt.Sprintf("[parameters('%sScaleSetEvictionPolicy')]", agentPoolName)
	expectedMap["provisionScriptParametersCommon"] = "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]"
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

	_, err := cs.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		t.Fatal(err)
	}

	varMap, err := GetKubernetesVariables(cs)
	if err != nil {
		t.Fatal(err)
	}

	expectedMap := map[string]interface{}{
		"apiVersionAuthorizationSystem":      "2018-01-01-preview",
		"apiVersionAuthorizationUser":        "2018-09-01-preview",
		"apiVersionCompute":                  "2019-07-01",
		"apiVersionDeployments":              "2018-06-01",
		"apiVersionKeyVault":                 "2018-02-14",
		"apiVersionManagedIdentity":          "2018-11-30",
		"apiVersionNetwork":                  "2018-08-01",
		"apiVersionStorage":                  "2018-07-01",
		"applicationInsightsKey":             "c92d8284-b550-4b06-b7ba-e80fd7178faa", // should be DefaultApplicationInsightsKey,
		"clusterKeyVaultName":                "",
		"contributorRoleDefinitionId":        "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', 'b24988ac-6180-42a0-ab88-20f7382dd24c')]",
		"customCloudAuthenticationMethod":    "client_secret",
		"customCloudIdentifySystem":          "azure_ad",
		"enableTelemetry":                    false,
		"etcdCaFilepath":                     "/etc/kubernetes/certs/ca.crt",
		"etcdClientCertFilepath":             "/etc/kubernetes/certs/etcdclient.crt",
		"etcdClientKeyFilepath":              "/etc/kubernetes/certs/etcdclient.key",
		"etcdPeerCertFilepath":               []string{"/etc/kubernetes/certs/etcdpeer0.crt", "/etc/kubernetes/certs/etcdpeer1.crt", "/etc/kubernetes/certs/etcdpeer2.crt", "/etc/kubernetes/certs/etcdpeer3.crt", "/etc/kubernetes/certs/etcdpeer4.crt"},
		"etcdPeerCertificates":               []string{"[parameters('etcdPeerCertificate0')]", "[parameters('etcdPeerCertificate1')]", "[parameters('etcdPeerCertificate2')]"},
		"etcdPeerKeyFilepath":                []string{"/etc/kubernetes/certs/etcdpeer0.key", "/etc/kubernetes/certs/etcdpeer1.key", "/etc/kubernetes/certs/etcdpeer2.key", "/etc/kubernetes/certs/etcdpeer3.key", "/etc/kubernetes/certs/etcdpeer4.key"},
		"etcdPeerPrivateKeys":                []string{"[parameters('etcdPeerPrivateKey0')]", "[parameters('etcdPeerPrivateKey1')]", "[parameters('etcdPeerPrivateKey2')]"},
		"etcdServerCertFilepath":             "/etc/kubernetes/certs/etcdserver.crt",
		"etcdServerKeyFilepath":              "/etc/kubernetes/certs/etcdserver.key",
		"excludeMasterFromStandardLB":        "true",
		"kubeconfigServer":                   "[concat('https://', variables('masterFqdnPrefix'), '.', variables('location'), '.', parameters('fqdnEndpointSuffix'))]",
		"kubernetesAPIServerIP":              "[concat(variables('masterFirstAddrPrefix'), add(variables('masterInternalLbIPOffset'), int(variables('masterFirstAddrOctet4'))))]",
		"labelResourceGroup":                 "[if(or(or(endsWith(variables('truncatedResourceGroup'), '-'), endsWith(variables('truncatedResourceGroup'), '_')), endsWith(variables('truncatedResourceGroup'), '.')), concat(take(variables('truncatedResourceGroup'), 62), 'z'), variables('truncatedResourceGroup'))]",
		"loadBalancerSku":                    api.StandardLoadBalancerSku,
		"location":                           "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
		"locations":                          []string{"[resourceGroup().location]", "[parameters('location')]"},
		"masterAvailabilitySet":              "[concat('master-availabilityset-', parameters('nameSuffix'))]",
		"masterCount":                        3,
		"masterEtcdClientPort":               2379,
		"masterEtcdClientURLs":               []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdClientPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdClientPort'))]"},
		"masterEtcdClusterStates":            []string{"[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2])]", "[concat(variables('masterVMNames')[0], '=', variables('masterEtcdPeerURLs')[0], ',', variables('masterVMNames')[1], '=', variables('masterEtcdPeerURLs')[1], ',', variables('masterVMNames')[2], '=', variables('masterEtcdPeerURLs')[2], ',', variables('masterVMNames')[3], '=', variables('masterEtcdPeerURLs')[3], ',', variables('masterVMNames')[4], '=', variables('masterEtcdPeerURLs')[4])]"},
		"masterEtcdPeerURLs":                 []string{"[concat('https://', variables('masterPrivateIpAddrs')[0], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[1], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[2], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[3], ':', variables('masterEtcdServerPort'))]", "[concat('https://', variables('masterPrivateIpAddrs')[4], ':', variables('masterEtcdServerPort'))]"},
		"masterEtcdServerPort":               2380,
		"masterEtcdMetricURLs":               []string{"[concat('http://', variables('masterPrivateIpAddrs')[0], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[1], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[2], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[3], ':2480')]", "[concat('http://', variables('masterPrivateIpAddrs')[4], ':2480')]"},
		"masterFirstAddrComment":             "these MasterFirstAddrComment are used to place multiple masters consecutively in the address space",
		"masterFirstAddrOctet4":              "[variables('masterFirstAddrOctets')[3]]",
		"masterFirstAddrOctets":              "[split(parameters('firstConsecutiveStaticIP'),'.')]",
		"masterFirstAddrPrefix":              "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
		"masterFqdnPrefix":                   "blueorange",
		"masterInternalLbID":                 "[resourceId('Microsoft.Network/loadBalancers',variables('masterInternalLbName'))]",
		"masterInternalLbIPConfigID":         "[concat(variables('masterInternalLbID'),'/frontendIPConfigurations/', variables('masterInternalLbIPConfigName'))]",
		"masterInternalLbIPConfigName":       "[concat(parameters('orchestratorName'), '-master-internal-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterInternalLbIPOffset":           10,
		"masterInternalLbName":               "[concat(parameters('orchestratorName'), '-master-internal-lb-', parameters('nameSuffix'))]",
		"masterLbBackendPoolName":            "[concat(parameters('orchestratorName'), '-master-pool-', parameters('nameSuffix'))]",
		"masterLbID":                         "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
		"masterLbIPConfigID":                 "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
		"masterLbIPConfigName":               "[concat(parameters('orchestratorName'), '-master-lbFrontEnd-', parameters('nameSuffix'))]",
		"masterLbName":                       "[concat(parameters('orchestratorName'), '-master-lb-', parameters('nameSuffix'))]",
		"masterOffset":                       "[parameters('masterOffset')]",
		"masterPrivateIpAddrs":               []string{"[concat(variables('masterFirstAddrPrefix'), add(0, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(1, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(2, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(3, int(variables('masterFirstAddrOctet4'))))]", "[concat(variables('masterFirstAddrPrefix'), add(4, int(variables('masterFirstAddrOctet4'))))]"},
		"masterPublicIPAddressName":          "[concat(parameters('orchestratorName'), '-master-ip-', variables('masterFqdnPrefix'), '-', parameters('nameSuffix'))]",
		"masterVMNamePrefix":                 "k8s-master-18280257-",
		"masterVMNames":                      []string{"[concat(variables('masterVMNamePrefix'), '0')]", "[concat(variables('masterVMNamePrefix'), '1')]", "[concat(variables('masterVMNamePrefix'), '2')]", "[concat(variables('masterVMNamePrefix'), '3')]", "[concat(variables('masterVMNamePrefix'), '4')]"},
		"maxVMsPerPool":                      100,
		"maximumLoadBalancerRuleCount":       250,
		"networkContributorRoleDefinitionId": "[concat('/subscriptions/', subscription().subscriptionId, '/providers/Microsoft.Authorization/roleDefinitions/', '4d97b98b-1d4f-4787-a291-c67834d212e7')]",
		"nsgID":                              "[resourceId('Microsoft.Network/networkSecurityGroups',variables('nsgName'))]",
		"nsgName":                            "[concat(variables('masterVMNamePrefix'), 'nsg')]",
		"orchestratorNameVersionTag":         "Kubernetes:" + testK8sVersion,
		"primaryAvailabilitySetName":         "",
		"primaryScaleSetName":                "",
		"cloudInitFiles": map[string]interface{}{
			"provisionScript":           getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript, cs),
			"provisionSource":           getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript, cs),
			"provisionInstalls":         getBase64EncodedGzippedCustomScript(kubernetesCSEInstall, cs),
			"provisionConfigs":          getBase64EncodedGzippedCustomScript(kubernetesCSEConfig, cs),
			"customSearchDomainsScript": getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript, cs),
			"generateProxyCertsScript":  getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript, cs),
			"mountEtcdScript":           getBase64EncodedGzippedCustomScript(kubernetesMountEtcd, cs),
			"etcdSystemdService":        getBase64EncodedGzippedCustomScript(etcdSystemdService, cs),
			"dhcpv6ConfigurationScript": getBase64EncodedGzippedCustomScript(dhcpv6ConfigurationScript, cs),
			"dhcpv6SystemdService":      getBase64EncodedGzippedCustomScript(dhcpv6SystemdService, cs),
			"kubeletSystemdService":     getBase64EncodedGzippedCustomScript(kubeletSystemdService, cs),
		},
		"provisionScriptParametersCommon":           "[concat('" + cs.GetProvisionScriptParametersCommon(api.ProvisionScriptParametersInput{Location: common.WrapAsARMVariable("location"), ResourceGroup: common.WrapAsARMVariable("resourceGroup"), TenantID: common.WrapAsARMVariable("tenantID"), SubscriptionID: common.WrapAsARMVariable("subscriptionId"), ClientID: common.WrapAsARMVariable("servicePrincipalClientId"), ClientSecret: common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"), APIServerCertificate: common.WrapAsParameter("apiServerCertificate"), KubeletPrivateKey: common.WrapAsParameter("clientPrivateKey"), ClusterKeyVaultName: common.WrapAsARMVariable("clusterKeyVaultName")}) + "')]",
		"provisionScriptParametersMaster":           "[concat('COSMOS_URI= MASTER_VM_NAME=',variables('masterVMNames')[variables('masterOffset')],' ETCD_PEER_URL=',variables('masterEtcdPeerURLs')[variables('masterOffset')],' ETCD_CLIENT_URL=',variables('masterEtcdClientURLs')[variables('masterOffset')],' MASTER_NODE=true NO_OUTBOUND=false AUDITD_ENABLED=false CLUSTER_AUTOSCALER_ADDON=false ACI_CONNECTOR_ADDON=',parameters('kubernetesACIConnectorEnabled'),' APISERVER_PRIVATE_KEY=',parameters('apiServerPrivateKey'),' CA_CERTIFICATE=',parameters('caCertificate'),' CA_PRIVATE_KEY=',parameters('caPrivateKey'),' MASTER_FQDN=',variables('masterFqdnPrefix'),' KUBECONFIG_CERTIFICATE=',parameters('kubeConfigCertificate'),' KUBECONFIG_KEY=',parameters('kubeConfigPrivateKey'),' ETCD_SERVER_CERTIFICATE=',parameters('etcdServerCertificate'),' ETCD_CLIENT_CERTIFICATE=',parameters('etcdClientCertificate'),' ETCD_SERVER_PRIVATE_KEY=',parameters('etcdServerPrivateKey'),' ETCD_CLIENT_PRIVATE_KEY=',parameters('etcdClientPrivateKey'),' ETCD_PEER_CERTIFICATES=',string(variables('etcdPeerCertificates')),' ETCD_PEER_PRIVATE_KEYS=',string(variables('etcdPeerPrivateKeys')),' ENABLE_AGGREGATED_APIS=',string(parameters('enableAggregatedAPIs')),' KUBECONFIG_SERVER=',variables('kubeconfigServer'))]",
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
		"subnetName":                                "k8s-subnet",
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
		"windowsCSIProxyURL":                        "",
		"windowsEnableCSIProxy":                     false,
	}
	diff := cmp.Diff(varMap, expectedMap)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestK8sVarsWindowsProfile(t *testing.T) {
	var trueVar = true
	cases := []struct {
		name         string
		wp           *api.WindowsProfile
		expectedVars map[string]interface{}
	}{
		{
			name: "No windows profile",
			wp:   nil,
			expectedVars: map[string]interface{}{
				"windowsEnableCSIProxy": false,
				"windowsCSIProxyURL":    "",
			},
		},
		{
			name: "Defaults",
			wp:   &api.WindowsProfile{},
			expectedVars: map[string]interface{}{
				"windowsEnableCSIProxy": false,
				"windowsCSIProxyURL":    "",
			},
		},
		{
			name: "Non-defaults",
			wp: &api.WindowsProfile{
				EnableCSIProxy: &trueVar,
				CSIProxyURL:    "http://some/package.tar",
			},
			expectedVars: map[string]interface{}{
				"windowsEnableCSIProxy": true,
				"windowsCSIProxyURL":    "http://some/package.tar",
			},
		},
	}

	for _, c := range cases {
		test := c
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			vars := getWindowsProfileVars(test.wp)

			diff := cmp.Diff(test.expectedVars, vars)
			if diff != "" {
				t.Errorf("unexpected diff in vars: %s", diff)
			}
		})
	}
}
