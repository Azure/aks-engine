package armhelpers

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	"github.com/Azure/go-autorest/autorest/to"
)

func createStorageAccount(cs *api.ContainerService) StorageAccountARM {
	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionStorage')]",
	}

	if !to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled) {
		armResource.DependsOn = []string{
			"[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]",
		}
	}

	storageAccount := storage.Account{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[variables('masterStorageAccountName')]"),
		Type:     to.StringPtr("Microsoft.Storage/storageAccounts"),
		Sku: &storage.Sku{
			Name: storage.SkuName("[variables('vmSizesMap')[parameters('masterVMSize')].storageAccountType]"),
		},
	}

	return StorageAccountARM{
		ARMResource: armResource,
		Account:     storageAccount,
	}
}

func createJumpboxStorageAccount(cs *api.ContainerService) StorageAccountARM {
	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionStorage')]",
	}

	storageAccount := storage.Account{
		Type:     to.StringPtr("Microsoft.Storage/storageAccounts"),
		Name:     to.StringPtr("[variables('jumpboxStorageAccountName')]"),
		Location: to.StringPtr("[variables('location')]"),
		Sku: &storage.Sku{
			Name: storage.SkuName("[variables('vmSizesMap')[parameters('jumpboxVMSize')].storageAccountType]"),
		},
	}

	return StorageAccountARM{
		ARMResource: armResource,
		Account:     storageAccount,
	}
}

func createKeyVaultStorageAccount(cs *api.ContainerService) StorageAccountARM {
	return StorageAccountARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionStorage')]",
		},
		Account: storage.Account{
			Type: to.StringPtr("Microsoft.Storage/storageAccounts"),
			Name: to.StringPtr("[variables('clusterKeyVaultName')]"),
			Location: to.StringPtr("[variables('location')]"),
			Sku: &storage.Sku{
				Name: storage.StandardLRS,
			},
		},
	}
}