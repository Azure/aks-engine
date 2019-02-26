// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
)

// DeleteManagedDisk deletes a managed disk.
func (az *AzureClient) DeleteManagedDisk(ctx context.Context, resourceGroupName string, diskName string) error {
	future, err := az.disksClient.Delete(ctx, resourceGroupName, diskName)
	if err != nil {
		return err
	}

	if err = future.WaitForCompletionRef(ctx, az.disksClient.Client); err != nil {
		return err
	}

	_, err = future.Result(az.disksClient)
	return err
}

// ListManagedDisksByResourceGroup lists managed disks in a resource group.
func (az *AzureClient) ListManagedDisksByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.DiskListPage, err error) {
	return az.disksClient.ListByResourceGroup(ctx, resourceGroupName)
}
