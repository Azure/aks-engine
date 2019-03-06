// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Azure/aks-engine/pkg/api"
)

func TestCreateKeyVault(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{},
			},
			MasterProfile: &api.MasterProfile{
				Count: 1,
			},
		},
	}

	actual := CreateKeyVaultVMAS(cs)

	expected := map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"location":   "[variables('location')]",
		"properties": map[string]interface{}{
			"accessPolicies": []interface{}{
				map[string]interface{}{
					"objectId": "[parameters('servicePrincipalObjectId')]",
					"permissions": map[string]interface{}{
						"keys": []string{"create", "encrypt", "decrypt", "get", "list"}},
					"tenantId": "[variables('tenantID')]"}},
			"enabledForDeployment":         "false",
			"enabledForDiskEncryption":     "false",
			"enabledForTemplateDeployment": "false",
			"sku": map[string]interface{}{
				"family": "A",
				"name":   "[parameters('clusterKeyVaultSku')]",
			},
			"tenantId": "[variables('tenantID')]"},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected error while comparing ARM resources: %s", diff)
	}

	//Test with UseManagedIdentityEnabled
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true

	actual = CreateKeyVaultVMAS(cs)

	expected = map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"dependsOn": []string{
			"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), '0')]",
			"[concat('Microsoft.Authorization/roleAssignments/', guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), '0', 'vmidentity')))]"},

		"location": "[variables('location')]",
		"properties": map[string]interface{}{
			"accessPolicies": []interface{}{
				map[string]interface{}{
					"objectId": "[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), 0'), '2017-03-30', 'Full').identity.principalId]",
					"permissions": map[string]interface{}{
						"keys": []string{"create", "encrypt", "decrypt", "get", "list"}},
					"tenantId": "[variables('tenantID')]"}},
			"enabledForDeployment":         "false",
			"enabledForDiskEncryption":     "false",
			"enabledForTemplateDeployment": "false",
			"sku": map[string]interface{}{
				"family": "A",
				"name":   "[parameters('clusterKeyVaultSku')]",
			},
			"tenantId": "[variables('tenantID')]"},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected error while comparing ARM resources: %s", diff)
	}

	//Test with UserAssignedID
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
	cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID = "fooID"

	actual = CreateKeyVaultVMAS(cs)

	expected = map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"dependsOn": []string{
			"[variables('userAssignedIDReference')]",
		},

		"location": "[variables('location')]",
		"properties": map[string]interface{}{
			"accessPolicies": []interface{}{
				map[string]interface{}{
					"objectId": "[reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).principalId]",
					"permissions": map[string]interface{}{
						"keys": []string{"create", "encrypt", "decrypt", "get", "list"}},
					"tenantId": "[variables('tenantID')]"}},
			"enabledForDeployment":         "false",
			"enabledForDiskEncryption":     "false",
			"enabledForTemplateDeployment": "false",
			"sku": map[string]interface{}{
				"family": "A",
				"name":   "[parameters('clusterKeyVaultSku')]",
			},
			"tenantId": "[variables('tenantID')]"},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected error while comparing ARM resources: %s", diff)
	}
}

func TestCreateKeyVaultVMSS(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{},
			},
			MasterProfile: &api.MasterProfile{
				Count: 1,
			},
		},
	}

	actual := CreateKeyVaultVMSS(cs)

	expected := map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"location":   "[variables('location')]",
		"properties": map[string]interface{}{
			"accessPolicies": []interface{}{
				map[string]interface{}{
					"permissions": map[string]interface{}{
						"keys": []string{"create", "encrypt", "decrypt", "get", "list"}},
					"tenantId": "[variables('tenantID')]"}},
			"enabledForDeployment":         "false",
			"enabledForDiskEncryption":     "false",
			"enabledForTemplateDeployment": "false",
			"sku": map[string]interface{}{
				"family": "A",
				"name":   "[parameters('clusterKeyVaultSku')]",
			},
			"tenantId": "[variables('tenantID')]"},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected error while comparing ARM resources: %s", diff)
	}

	//Test with UseManagedIdentityEnabled
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true

	actual = CreateKeyVaultVMSS(cs)

	expected = map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"dependsOn": []string{
			"[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('masterVMNamePrefix'), 'vmss')]",
		},

		"location": "[variables('location')]",
		"properties": map[string]interface{}{
			"accessPolicies": []interface{}{
				map[string]interface{}{
					"objectId": "[parameters('servicePrincipalObjectId')]",
					"permissions": map[string]interface{}{
						"keys": []string{"create", "encrypt", "decrypt", "get", "list"}},
					"tenantId": "[variables('tenantID')]"}},
			"enabledForDeployment":         "false",
			"enabledForDiskEncryption":     "false",
			"enabledForTemplateDeployment": "false",
			"sku": map[string]interface{}{
				"family": "A",
				"name":   "[parameters('clusterKeyVaultSku')]",
			},
			"tenantId": "[variables('tenantID')]"},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected error while comparing ARM resources: %s", diff)
	}

	//Test with UserAssignedID
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
	cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID = "fooID"

	actual = CreateKeyVaultVMSS(cs)

	expected = map[string]interface{}{
		"type":       "Microsoft.KeyVault/vaults",
		"name":       "[variables('clusterKeyVaultName')]",
		"apiVersion": "[variables('apiVersionKeyVault')]",
		"dependsOn": []string{
			"[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('masterVMNamePrefix'), 'vmss')]",
			"[variables('userAssignedIDReference')]",
		},

		"location": "[variables('location')]",
		"properties": map[string]interface{}{
			"accessPolicies": []interface{}{
				map[string]interface{}{
					"objectId": "[reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).principalId]",
					"permissions": map[string]interface{}{
						"keys": []string{"create", "encrypt", "decrypt", "get", "list"}},
					"tenantId": "[variables('tenantID')]"}},
			"enabledForDeployment":         "false",
			"enabledForDiskEncryption":     "false",
			"enabledForTemplateDeployment": "false",
			"sku": map[string]interface{}{
				"family": "A",
				"name":   "[parameters('clusterKeyVaultSku')]",
			},
			"tenantId": "[variables('tenantID')]"},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected error while comparing ARM resources: %s", diff)
	}
}
