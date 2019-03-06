// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func createWindowsImage(profile *api.AgentPoolProfile) ImageARM {
	return ImageARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
		},
		Image: compute.Image{
			Type: to.StringPtr("Microsoft.Compute/images"),
			Name: to.StringPtr(fmt.Sprintf("%sCustomWindowsImage", profile.Name)),
			ImageProperties: &compute.ImageProperties{
				StorageProfile: &compute.ImageStorageProfile{
					OsDisk: &compute.ImageOSDisk{
						OsType:             "Windows",
						OsState:            compute.Generalized,
						BlobURI:            to.StringPtr("[parameters('agentWindowsSourceUrl')]"),
						StorageAccountType: compute.StorageAccountTypesStandardLRS,
					},
				},
			},
		},
	}
}
