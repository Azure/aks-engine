// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import "context"

// ListResourceSkus lists Microsoft.Compute SKUs available for a subscription
func (az *AzureClient) ListResourceSkus(ctx context.Context, filter string) (ResourceSkusResultPage, error) {
	page, err := az.resourceSkusClient.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &page, nil
}
