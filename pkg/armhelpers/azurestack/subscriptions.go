// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-06-01/subscriptions"
)

// ListLocations returns the Azure regions available to the subscription.
func (az *AzureClient) ListLocations(ctx context.Context) (*[]subscriptions.Location, error) {
	list, error := az.subscriptionsClient.ListLocations(ctx, az.subscriptionID)
	if error != nil {
		return nil, error
	}
	return list.Value, nil
}
