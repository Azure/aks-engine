// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

// ListVirtualMachineImages returns the list of images available in the current environment
func (az *AzureClient) ListVirtualMachineImages(ctx context.Context, location, publisherName, offer, skus string) (compute.ListVirtualMachineImageResource, error) {
	// random value
	top := int32(10)
	list, err := az.virtualMachineImagesClient.List(ctx, location, publisherName, offer, skus, "", &top, "")
	if err != nil {
		return compute.ListVirtualMachineImageResource{}, fmt.Errorf("failed to list virtual machine images, %s", err)
	}
	r := compute.ListVirtualMachineImageResource{}
	if err = DeepCopy(&r, list); err != nil {
		return r, fmt.Errorf("failed to deep copy virtual machine images, %s", err)
	}
	return r, err
}

// GetVirtualMachineImage returns an image or an error if the image is not found
func (az *AzureClient) GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (compute.VirtualMachineImage, error) {
	image, err := az.virtualMachineImagesClient.Get(ctx, location, publisherName, offer, skus, version)
	if err != nil {
		return compute.VirtualMachineImage{}, fmt.Errorf("failed to get virtual machine image, %s", err)
	}
	r := compute.VirtualMachineImage{}
	if err = DeepCopy(&r, image); err != nil {
		return r, fmt.Errorf("failed to deep copy virtual machine image, %s", err)
	}
	return r, err
}
