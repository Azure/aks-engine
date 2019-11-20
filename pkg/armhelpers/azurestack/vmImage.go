package azurestack

import (
	"context"
	"fmt"

	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
)

//ListVirtualMachineImages returns the list of images available in the current environment
func (az *AzureClient) ListVirtualMachineImages(ctx context.Context, top int32, location, publisherName, offer, skus, filter, orderBy string) (result azcompute.ListVirtualMachineImageResource, err error) {
	vmImages, err := az.virtualMachineImageClient.List(ctx, location, publisherName, offer, skus, filter, &top, orderBy)
	azVMImages := azcompute.ListVirtualMachineImageResource{}

	if err != nil {
		return azVMImages, fmt.Errorf("fail to get virtual machine images, %s", err)
	}
	err = DeepCopy(&azVMImages, vmImages)
	if err != nil {
		return azVMImages, fmt.Errorf("fail to convert virtual machine images, %s", err)
	}
	return azVMImages, err
}

//GetVirtualMachineImage returns an image or an error where there is no image
func (az *AzureClient) GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (result azcompute.VirtualMachineImage, err error) {
	vmImage, err := az.virtualMachineImageClient.Get(ctx, location, publisherName, offer, skus, version)

	azVMImage := azcompute.VirtualMachineImage{}

	if err != nil {
		return azVMImage, fmt.Errorf("fail to get virtual machine image, %s", err)
	}
	err = DeepCopy(&azVMImage, vmImage)
	if err != nil {
		return azVMImage, fmt.Errorf("fail to convert virtual machine images, %s", err)
	}
	return azVMImage, err
}
