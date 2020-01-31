// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
)

func GetKubernetesOutputs(cs *api.ContainerService) map[string]interface{} {
	outputs := map[string]interface{}{
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
	}

	isHostedMaster := cs.Properties.IsHostedMasterProfile()

	if !isHostedMaster {
		for k, v := range getMasterOutputs(cs) {
			outputs[k] = v
		}
	}

	for _, profile := range cs.Properties.AgentPoolProfiles {
		if profile.IsAvailabilitySets() && profile.IsStorageAccount() {
			agentName := profile.Name
			outputs[fmt.Sprintf("%sStorageAccountOffset", agentName)] = map[string]interface{}{
				"type":  "int",
				"value": fmt.Sprintf("[variables('%sStorageAccountOffset')]", agentName),
			}
			outputs[fmt.Sprintf("%sStorageAccountCount", agentName)] = map[string]interface{}{
				"type":  "int",
				"value": fmt.Sprintf("[variables('%sStorageAccountsCount')]", agentName),
			}
			outputs[fmt.Sprintf("%sSubnetName", agentName)] = map[string]interface{}{
				"type":  "string",
				"value": fmt.Sprintf("[variables('%sSubnetName')]", agentName),
			}
		}
	}

	if cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AppGwIngressAddonName) {
		outputs["applicationGatewayName"] = map[string]interface{}{
			"type":  "string",
			"value": "[variables('appGwName')]",
		}
		outputs["appGwIdentityResourceId"] = map[string]interface{}{
			"type":  "string",
			"value": "[variables('appGwICIdentityId')]",
		}
		outputs["appGwIdentityClientId"] = map[string]interface{}{
			"type":  "string",
			"value": "[reference(variables('appGwICIdentityId'), variables('apiVersionManagedIdentity')).clientId]",
		}
	}

	return outputs
}

func getMasterOutputs(cs *api.ContainerService) map[string]interface{} {
	outputs := map[string]interface{}{}
	masterFQDN := ""

	if !cs.Properties.OrchestratorProfile.IsPrivateCluster() {
		masterFQDN = "[reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn]"
	}

	outputs["masterFQDN"] = map[string]interface{}{
		"type":  "string",
		"value": masterFQDN,
	}

	if cs.Properties.HasVMASAgentPool() {
		outputs["agentStorageAccountSuffix"] = map[string]interface{}{
			"type":  "string",
			"value": "[variables('storageAccountBaseName')]",
		}
		outputs["agentStorageAccountPrefixes"] = map[string]interface{}{
			"type":  "array",
			"value": "[variables('storageAccountPrefixes')]",
		}
	}

	return outputs
}
