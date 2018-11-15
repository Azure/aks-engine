// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
)

// ListProviders returns all the providers for a given AzureClient
func (az *AzureClient) ListProviders(ctx context.Context) (resources.ProviderListResultPage, error) {
	return az.providersClient.List(ctx, to.Int32Ptr(100), "")
}
