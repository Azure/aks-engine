// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure"
)

// GetKeyVaultSecret gets the specified KeyVault secret.
func (az *AzureClient) GetKeyVaultSecret(ctx context.Context, vaultName string, secretName string, secretVersion string) (secret string, err error) {
	secretBundle, err := az.keyvaultClient.GetSecret(ctx, fmt.Sprintf("https://%s.%s", vaultName, azure.PublicCloud.KeyVaultDNSSuffix), secretName, secretVersion)
	if err != nil {
		return "", err
	}
	return *secretBundle.Value, nil
}
