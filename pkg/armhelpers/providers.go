// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/go-autorest/autorest/to"
)

// ListProviders returns all the providers for a given AzureClient
func (az *AzureClient) ListProviders(ctx context.Context) (ProviderListResultPage, error) {
	page, err := az.providersClient.List(ctx, to.Int32Ptr(100), "")
	return &page, err
}
