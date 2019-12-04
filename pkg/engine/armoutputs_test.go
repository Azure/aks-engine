// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestGetK8sOutputs(t *testing.T) {
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
					PrivateCluster: &api.PrivateCluster{
						Enabled: to.BoolPtr(false),
					},
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

	outputMap := GetKubernetesOutputs(cs)

	expected := map[string]interface{}{
		"resourceGroup": map[string]interface{}{
			"type":  "string",
			"value": "[variables('resourceGroup')]",
		},
		"vnetResourceGroup": map[string]interface{}{
			"type":  "string",
			"value": "[variables('virtualNetworkResourceGroupName')]",
		},
		"subnetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('subnetName')]",
		},
		"securityGroupName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('nsgName')]",
		},
		"virtualNetworkName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('virtualNetworkName')]",
		},
		"routeTableName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('routeTableName')]",
		},
		"primaryAvailabilitySetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('primaryAvailabilitySetName')]",
		},
		"primaryScaleSetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('primaryScaleSetName')]",
		},
		"masterFQDN": map[string]interface{}{
			"type":  "string",
			"value": "[reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn]",
		},
	}
	diff := cmp.Diff(outputMap, expected)

	if diff != "" {
		t.Errorf("unexpected error while comparing output maps: %s", diff)
	}
}

func TestK8sOutputsWithAvSets(t *testing.T) {
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
					PrivateCluster: &api.PrivateCluster{
						Enabled: to.BoolPtr(false),
					},
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool1",
					VMSize:              "Standard_D2_v2",
					Count:               2,
					AvailabilityProfile: api.AvailabilitySet,
					StorageProfile:      api.StorageAccount,
				},
			},
		},
	}

	outputMap := GetKubernetesOutputs(cs)

	expected := map[string]interface{}{
		"resourceGroup": map[string]interface{}{
			"type":  "string",
			"value": "[variables('resourceGroup')]",
		},
		"vnetResourceGroup": map[string]interface{}{
			"type":  "string",
			"value": "[variables('virtualNetworkResourceGroupName')]",
		},
		"subnetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('subnetName')]",
		},
		"securityGroupName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('nsgName')]",
		},
		"virtualNetworkName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('virtualNetworkName')]",
		},
		"routeTableName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('routeTableName')]",
		},
		"primaryAvailabilitySetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('primaryAvailabilitySetName')]",
		},
		"primaryScaleSetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('primaryScaleSetName')]",
		},
		"masterFQDN": map[string]interface{}{
			"type":  "string",
			"value": "[reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn]",
		},
		"agentpool1StorageAccountOffset": map[string]interface{}{
			"type":  "int",
			"value": "[variables('agentpool1StorageAccountOffset')]",
		},
		"agentpool1StorageAccountCount": map[string]interface{}{
			"type":  "int",
			"value": "[variables('agentpool1StorageAccountsCount')]",
		},
		"agentpool1SubnetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('agentpool1SubnetName')]",
		},

		"agentStorageAccountSuffix": map[string]interface{}{
			"type":  "string",
			"value": "[variables('storageAccountBaseName')]",
		},
		"agentStorageAccountPrefixes": map[string]interface{}{
			"type":  "array",
			"value": "[variables('storageAccountPrefixes')]",
		},
	}

	diff := cmp.Diff(outputMap, expected)

	if diff != "" {
		t.Errorf("unexpected error while comparing output maps: %s", diff)
	}
}

func TestK8sOutputsWithAppGwIngressAddon(t *testing.T) {
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
					Addons: []api.KubernetesAddon{
						{
							Name:    common.AppGwIngressAddonName,
							Enabled: to.BoolPtr(true),
						},
					},
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool1",
					VMSize:              "Standard_D2_v2",
					Count:               2,
					AvailabilityProfile: api.AvailabilitySet,
					StorageProfile:      api.StorageAccount,
				},
			},
		},
	}

	outputMap := GetKubernetesOutputs(cs)

	expected := map[string]interface{}{
		"resourceGroup": map[string]interface{}{
			"type":  "string",
			"value": "[variables('resourceGroup')]",
		},
		"vnetResourceGroup": map[string]interface{}{
			"type":  "string",
			"value": "[variables('virtualNetworkResourceGroupName')]",
		},
		"subnetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('subnetName')]",
		},
		"securityGroupName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('nsgName')]",
		},
		"virtualNetworkName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('virtualNetworkName')]",
		},
		"routeTableName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('routeTableName')]",
		},
		"primaryAvailabilitySetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('primaryAvailabilitySetName')]",
		},
		"primaryScaleSetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('primaryScaleSetName')]",
		},
		"masterFQDN": map[string]interface{}{
			"type":  "string",
			"value": "[reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn]",
		},
		"agentpool1StorageAccountOffset": map[string]interface{}{
			"type":  "int",
			"value": "[variables('agentpool1StorageAccountOffset')]",
		},
		"agentpool1StorageAccountCount": map[string]interface{}{
			"type":  "int",
			"value": "[variables('agentpool1StorageAccountsCount')]",
		},
		"agentpool1SubnetName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('agentpool1SubnetName')]",
		},
		"agentStorageAccountSuffix": map[string]interface{}{
			"type":  "string",
			"value": "[variables('storageAccountBaseName')]",
		},
		"agentStorageAccountPrefixes": map[string]interface{}{
			"type":  "array",
			"value": "[variables('storageAccountPrefixes')]",
		},
		"applicationGatewayName": map[string]interface{}{
			"type":  "string",
			"value": "[variables('appGwName')]",
		},
		"appGwIdentityResourceId": map[string]interface{}{
			"type":  "string",
			"value": "[variables('appGwICIdentityId')]",
		},
		"appGwIdentityClientId": map[string]interface{}{
			"type":  "string",
			"value": "[reference(variables('appGwICIdentityId'), variables('apiVersionManagedIdentity')).clientId]",
		},
	}

	diff := cmp.Diff(outputMap, expected)

	if diff != "" {
		t.Errorf("unexpected error while comparing output maps: %s", diff)
	}
}
