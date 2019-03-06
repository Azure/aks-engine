// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
)

func CreateKeyVaultVMAS(cs *api.ContainerService) map[string]interface{} {
	keyVaultMap := map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"location":   "[variables('location')]",
	}

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""
	masterCount := cs.Properties.MasterProfile.Count

	if useManagedIdentity {
		var dependencies []string

		if userAssignedIDEnabled {
			dependencies = append(dependencies, "[variables('userAssignedIDReference')]")
		} else {
			for i := 0; i < masterCount; i++ {
				dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), '%d')]", i))
				dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Authorization/roleAssignments/', guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), '%d', 'vmidentity')))]", i))
			}
		}
		keyVaultMap["dependsOn"] = dependencies
	}

	keyVaultProps := map[string]interface{}{
		"enabledForDeployment":         "false",
		"enabledForDiskEncryption":     "false",
		"enabledForTemplateDeployment": "false",
		"tenantId":                     "[variables('tenantID')]",
		"sku": map[string]interface{}{
			"name":   "[parameters('clusterKeyVaultSku')]",
			"family": "A",
		},
	}

	var accessPolicies []interface{}

	if useManagedIdentity {
		if userAssignedIDEnabled {
			accessPolicy := map[string]interface{}{
				"tenantId": "[variables('tenantID')]",
				"objectId": "[reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).principalId]",
				"permissions": map[string]interface{}{
					"keys": []string{"create", "encrypt", "decrypt", "get", "list"},
				},
			}
			accessPolicies = append(accessPolicies, accessPolicy)
		} else {
			for i := 0; i < masterCount; i++ {
				accessPolicy := map[string]interface{}{
					"objectId": fmt.Sprintf("[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), %d'), '2017-03-30', 'Full').identity.principalId]", i),
					"permissions": map[string]interface{}{
						"keys": []string{
							"create",
							"encrypt",
							"decrypt",
							"get",
							"list",
						},
					},
					"tenantId": "[variables('tenantID')]",
				}
				accessPolicies = append(accessPolicies, accessPolicy)
			}
		}
	} else {
		accessPolicy := map[string]interface{}{
			"tenantId": "[variables('tenantID')]",
			"objectId": "[parameters('servicePrincipalObjectId')]",
			"permissions": map[string]interface{}{
				"keys": []string{"create", "encrypt", "decrypt", "get", "list"},
			},
		}
		accessPolicies = append(accessPolicies, accessPolicy)
	}
	keyVaultProps["accessPolicies"] = accessPolicies
	keyVaultMap["properties"] = keyVaultProps

	return keyVaultMap
}

func CreateKeyVaultVMSS(cs *api.ContainerService) map[string]interface{} {
	keyVaultMap := map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"location":   "[variables('location')]",
	}

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	accessPolicy := map[string]interface{}{
		"tenantId": "[variables('tenantID')]",
		"permissions": map[string]interface{}{
			"keys": []string{"create", "encrypt", "decrypt", "get", "list"},
		},
	}
	if useManagedIdentity {
		dependencies := []string{
			"[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('masterVMNamePrefix'), 'vmss')]",
		}
		if userAssignedIDEnabled {
			dependencies = append(dependencies, "[variables('userAssignedIDReference')]")
			accessPolicy["objectId"] = "[reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).principalId]"
		} else {
			accessPolicy["objectId"] = "[parameters('servicePrincipalObjectId')]"
		}
		keyVaultMap["dependsOn"] = dependencies
	}

	keyVaultProps := map[string]interface{}{
		"enabledForDeployment":         "false",
		"enabledForDiskEncryption":     "false",
		"enabledForTemplateDeployment": "false",
		"tenantId":                     "[variables('tenantID')]",
		"sku": map[string]interface{}{
			"name":   "[parameters('clusterKeyVaultSku')]",
			"family": "A",
		},
		"accessPolicies": []interface{}{
			accessPolicy,
		},
	}

	keyVaultMap["properties"] = keyVaultProps

	return keyVaultMap
}
