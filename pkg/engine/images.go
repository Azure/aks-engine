// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func createWindowsImageReference(agentPoolProfileName string, windowsProfile *api.WindowsProfile) *compute.ImageReference {
	var computeImageRef compute.ImageReference

	if windowsProfile.HasCustomImage() {
		computeImageRef = compute.ImageReference{
			ID: to.StringPtr(fmt.Sprintf("[resourceId('Microsoft.Compute/images', '%sCustomWindowsImage')]", agentPoolProfileName)),
		}
	} else if windowsProfile.HasImageRef() {
		imageRef := windowsProfile.ImageRef
		if windowsProfile.HasImageGallery() {
			computeImageRef = compute.ImageReference{
				ID: to.StringPtr(fmt.Sprintf("[concat('/subscriptions/', '%s', '/resourceGroups/', parameters('agentWindowsImageResourceGroup'), '/providers/Microsoft.Compute/galleries/', '%s', '/images/', parameters('agentWindowsImageName'), '/versions/', '%s')]", imageRef.SubscriptionID, imageRef.Gallery, imageRef.Version)),
			}
		} else {
			computeImageRef = compute.ImageReference{
				ID: to.StringPtr("[resourceId(parameters('agentWindowsImageResourceGroup'), 'Microsoft.Compute/images', parameters('agentWindowsImageName'))]"),
			}
		}
	} else {
		computeImageRef = compute.ImageReference{
			Offer:     to.StringPtr("[parameters('agentWindowsOffer')]"),
			Publisher: to.StringPtr("[parameters('agentWindowsPublisher')]"),
			Sku:       to.StringPtr("[parameters('agentWindowsSku')]"),
			Version:   to.StringPtr("[parameters('agentWindowsVersion')]"),
		}
	}

	return &computeImageRef
}

func createWindowsImage(profile *api.AgentPoolProfile) ImageARM {
	return ImageARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
		},
		Image: compute.Image{
			Type:     to.StringPtr("Microsoft.Compute/images"),
			Name:     to.StringPtr(fmt.Sprintf("%sCustomWindowsImage", profile.Name)),
			Location: to.StringPtr("[variables('location')]"),
			ImageProperties: &compute.ImageProperties{
				StorageProfile: &compute.ImageStorageProfile{
					OsDisk: &compute.ImageOSDisk{
						OsType:             "Windows",
						OsState:            compute.Generalized,
						BlobURI:            to.StringPtr("[parameters('agentWindowsSourceUrl')]"),
						StorageAccountType: compute.StorageAccountTypesStandardLRS,
					},
				},
				// TODO: Expose Hyper-V generation for VHD URL refs in apimodel
				HyperVGeneration: compute.HyperVGenerationTypesV1,
			},
		},
	}
}
