package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
)

//ListVirtualMachineImages returns the list of images available in the current environment
func (az *AzureClient) ListVirtualMachineImages(ctx context.Context, top int32, location, publisherName, offer, skus, filter, orderBy string) (result compute.ListVirtualMachineImageResource, err error) {
	list, err := az.virtualMachineImageClient.List(ctx, location, publisherName, offer, skus, filter, &top, orderBy)

	return list, err
}

//GetVirtualMachineImage returns an image or an error where there is no image
func (az *AzureClient) GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (result compute.VirtualMachineImage, err error) {
	image, err := az.virtualMachineImageClient.Get(ctx, location, publisherName, offer, skus, version)

	return image, err
}
