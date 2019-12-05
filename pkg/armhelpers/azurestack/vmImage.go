package azurestack

import (
	"context"
	"fmt"

	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

//ListVirtualMachineImages returns the list of images available in the current environment
func (az *AzureClient) ListVirtualMachineImages(ctx context.Context, location, publisherName, offer, skus, filter string, top int32, orderBy string) (azcompute.ListVirtualMachineImageResource, error) {
	vmImages, err := az.virtualMachineImageClient.List(ctx, location, publisherName, offer, skus, filter, &top, orderBy)
	azVMImages := azcompute.ListVirtualMachineImageResource{}

	if err != nil {
		return azVMImages, fmt.Errorf("fail to get virtual machine images, %s", err)
	}

	if err := DeepCopy(&azVMImages, vmImages); err != nil {
		return azVMImages, fmt.Errorf("fail to convert virtual machine images, %s", err)
	}
	return azVMImages, err
}

//GetVirtualMachineImage returns an image or an error if the image is not found
func (az *AzureClient) GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (azcompute.VirtualMachineImage, error) {
	vmImage, err := az.virtualMachineImageClient.Get(ctx, location, publisherName, offer, skus, version)

	azVMImage := azcompute.VirtualMachineImage{}

	if err != nil {
		return azVMImage, fmt.Errorf("fail to get virtual machine image, %s", err)
	}
	if err := DeepCopy(&azVMImage, vmImage); err != nil {
		return azVMImage, fmt.Errorf("fail to convert virtual machine images, %s", err)
	}
	return azVMImage, err
}
