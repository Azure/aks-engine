// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
)

//func createKeyVaults(cs *api.ContainerService) KeyVaultARM {
//
//	armResource := ARMResource{
//		APIVersion: "[variables('apiVersionKeyVault')]",
//	}
//
//	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
//	userAssignedIDEnabled := cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""
//	if useManagedIdentity {
//		var dependencies []string
//		if userAssignedIDEnabled {
//			dependencies = append(dependencies, "[variables('userAssignedIDReference')]")
//		} else {
//			masterProfileCount := cs.Properties.MasterProfile.Count
//			for i := 0; i < masterProfileCount; i++ {
//				dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), '%d')]", i))
//				dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Authorization/roleAssignments/', guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), '%d', 'vmidentity')))]", i))
//			}
//		}
//		armResource.DependsOn = dependencies
//	}
//
//	vaultProps := keyvault.VaultProperties{
//		EnabledForDeployment:         to.BoolPtr(false),
//		EnabledForDiskEncryption:     to.BoolPtr(false),
//		EnabledForTemplateDeployment: to.BoolPtr(false),
//		//TODO: Need to find a way to set tenantID
//		//TenantID: to.StringPtr("[variables('tenantID')]"),
//		Sku: &keyvault.Sku{
//			Name:   keyvault.SkuName("[parameters('clusterKeyVaultSku')]"),
//			Family: to.StringPtr("A"),
//		},
//	}
//
//	var accessPolicies []keyvault.AccessPolicyEntry
//	if useManagedIdentity {
//		if userAssignedIDEnabled {
//			accessPolicy := keyvault.AccessPolicyEntry{
//				//TODO: Need to find a way to set tenantID
//				//"tenantId": "[variables('tenantID')]",
//				//TODO: Need to find a way to set objectID
//				//"objectId": "[reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).principalId]",
//				Permissions: &keyvault.Permissions{
//					Keys: &[]keyvault.KeyPermissions{"create", "encrypt", "decrypt", "get", "list"},
//				},
//			}
//			accessPolicies = append(accessPolicies, accessPolicy)
//		} else {
//			masterProfileCount := cs.Properties.MasterProfile.Count
//			for i := 0; i < masterProfileCount; i++ {
//				accessPolicy := keyvault.AccessPolicyEntry{
//					//"objectId": "[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), '{{$i}}'), '2017-03-30', 'Full').identity.principalId]",
//					Permissions: &keyvault.Permissions{
//						Keys: &[]keyvault.KeyPermissions{
//							"create",
//							"encrypt",
//							"decrypt",
//							"get",
//							"list",
//						},
//					},
//					//"tenantId": "[variables('tenantID')]"
//				}
//				accessPolicies = append(accessPolicies, accessPolicy)
//			}
//		}
//	} else {
//		accessPolicy := keyvault.AccessPolicyEntry{
//			//"tenantId": "[variables('tenantID')]",
//			//"objectId": "[parameters('servicePrincipalObjectId')]",
//			Permissions: &keyvault.Permissions{
//				Keys: &[]keyvault.KeyPermissions{"create", "encrypt", "decrypt", "get", "list"},
//			},
//		}
//		accessPolicies = append(accessPolicies, accessPolicy)
//	}
//
//	vaultProps.AccessPolicies = &accessPolicies
//
//	vault := keyvault.Vault{
//		Type:       to.StringPtr("Microsoft.KeyVault/vaults"),
//		Name:       to.StringPtr("[variables('clusterKeyVaultName')]"),
//		Location:   to.StringPtr("[variables('location')]"),
//		Properties: &vaultProps,
//	}
//
//	return KeyVaultARM{
//		ARMResource: armResource,
//		Vault:       vault,
//	}
//}

func CreateKeyVault(cs *api.ContainerService) map[string]interface{} {
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
