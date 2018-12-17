package armhelpers

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func createAvailabilitySet(cs *api.ContainerService) AvailabilitySetARM {

	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionCompute')]",
	}

	avSet := compute.AvailabilitySet{
		Name:     to.StringPtr("[variables('masterAvailabilitySet')]"),
		Location: to.StringPtr("[variables('location')]"),
		Type:     to.StringPtr("Microsoft.Compute/availabilitySets"),
	}

	if !cs.Properties.MasterProfile.HasAvailabilityZones() {
		if cs.Properties.MasterProfile.IsManagedDisks() {
			avSet.AvailabilitySetProperties = &compute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(2),
				PlatformUpdateDomainCount: to.Int32Ptr(3),
			}
			avSet.Sku = &compute.Sku{
				Name: to.StringPtr("Aligned"),
			}
		} else if cs.Properties.MasterProfile.IsStorageAccount() {
			avSet.AvailabilitySetProperties = &compute.AvailabilitySetProperties{}
		}
	}

	return AvailabilitySetARM{
		ARMResource:     armResource,
		AvailabilitySet: avSet,
	}
}
