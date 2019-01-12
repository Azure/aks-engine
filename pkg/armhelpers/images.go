package armhelpers

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func createWindowsImage(cs *api.ContainerService, profile *api.AgentPoolProfile) ImageARM {
	return ImageARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionCompute')]",
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
