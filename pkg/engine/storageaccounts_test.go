// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"

	"testing"
)

func TestCreateStorageAccount(t *testing.T) {

	cs := &api.ContainerService{

		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					PrivateCluster: &api.PrivateCluster{
						Enabled: to.BoolPtr(false),
					},
				},
			},
		},
	}

	actual := createStorageAccount(cs)

	expected := StorageAccountARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionStorage')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]",
			},
		},
		Account: storage.Account{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('masterStorageAccountName')]"),
			Type:     to.StringPtr("Microsoft.Storage/storageAccounts"),
			Sku: &storage.Sku{
				Name: storage.SkuName("[variables('vmSizesMap')[parameters('masterVMSize')].storageAccountType]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateJumpboxStorageAccount(t *testing.T) {
	actual := createJumpboxStorageAccount()

	expected := StorageAccountARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionStorage')]",
		},
		Account: storage.Account{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('jumpboxStorageAccountName')]"),
			Type:     to.StringPtr("Microsoft.Storage/storageAccounts"),
			Sku: &storage.Sku{
				Name: storage.SkuName("[variables('vmSizesMap')[parameters('jumpboxVMSize')].storageAccountType]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateAgentVMASStorageAccount(t *testing.T) {

	cs := &api.ContainerService{

		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					PrivateCluster: &api.PrivateCluster{
						Enabled: to.BoolPtr(false),
					},
				},
			},
		},
	}

	profile := &api.AgentPoolProfile{
		Name: "fooAgent",
	}

	// Test Create VMAS Storage agent with Data Disk true
	actual := createAgentVMASStorageAccount(cs, profile, true)

	expected := StorageAccountARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionStorage')]",
			Copy: map[string]string{
				"count": "[variables('fooAgentStorageAccountsCount')]",
				"name":  "datadiskLoop",
			},
			DependsOn: []string{
				"[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]",
			},
		},
		Account: storage.Account{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('fooAgentStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('fooAgentStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('fooAgentDataAccountName'))]"),
			Type:     to.StringPtr("Microsoft.Storage/storageAccounts"),
			Sku: &storage.Sku{
				Name: storage.SkuName("[variables('vmSizesMap')[variables('fooAgentVMSize')].storageAccountType]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test Create VMAS Storage agent with Data Disk false
	actual = createAgentVMASStorageAccount(cs, profile, false)

	expected.Copy = map[string]string{
		"count": "[variables('fooAgentStorageAccountsCount')]",
		"name":  "loop",
	}

	expected.Account.Name = to.StringPtr("[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('fooAgentStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('fooAgentStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('fooAgentAccountName'))]")

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateKeyVaultStorageAccount(t *testing.T) {
	actual := createKeyVaultStorageAccount()
	expected := StorageAccountARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionStorage')]",
		},
		Account: storage.Account{
			Type:     to.StringPtr("Microsoft.Storage/storageAccounts"),
			Name:     to.StringPtr("[variables('clusterKeyVaultName')]"),
			Location: to.StringPtr("[variables('location')]"),
			Sku: &storage.Sku{
				Name: storage.StandardLRS,
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}
