// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

// ListVirtualMachineImages returns the list of images available in the current environment
func (az *AzureClient) ListVirtualMachineImages(ctx context.Context, location, publisherName, offer, skus string) (compute.ListVirtualMachineImageResource, error) {
	// random value
	top := int32(10)
	list, err := az.virtualMachineImagesClient.List(ctx, location, publisherName, offer, skus, "", &top, "")
	return list, err
}

// GetVirtualMachineImage returns an image or an error where there is no image
func (az *AzureClient) GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (compute.VirtualMachineImage, error) {
	image, err := az.virtualMachineImagesClient.Get(ctx, location, publisherName, offer, skus, version)
	return image, err
}
