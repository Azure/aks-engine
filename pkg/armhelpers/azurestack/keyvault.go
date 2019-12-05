// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"fmt"
)

// GetSecret gets the specified KeyVault secret.
func (az *AzureClient) GetKeyVaultSecret(ctx context.Context, vaultName string, secretName string, secretVersion string) (secret string, err error) {
	secretBundle, err := az.keyvaultClient.GetSecret(ctx, fmt.Sprintf("https://%s.vault.azure.net", vaultName), secretName, secretVersion)
	if err != nil {
		return "", err
	}
	return *secretBundle.Value, nil
}
