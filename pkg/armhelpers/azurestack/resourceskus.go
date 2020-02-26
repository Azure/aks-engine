// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"

	"github.com/Azure/aks-engine/pkg/armhelpers"
)

// ListResourceSkus lists Microsoft.Compute SKUs available for a subscription
func (az *AzureClient) ListResourceSkus(ctx context.Context, filter string) (armhelpers.ResourceSkusResultPage, error) {
	page, err := az.resourceSkusClient.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &page, nil
}
