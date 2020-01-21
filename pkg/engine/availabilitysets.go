// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateAvailabilitySet(cs *api.ContainerService, isManagedDisks bool) AvailabilitySetARM {

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
	}

	avSet := compute.AvailabilitySet{
		Name:     to.StringPtr("[variables('masterAvailabilitySet')]"),
		Location: to.StringPtr("[variables('location')]"),
		Type:     to.StringPtr("Microsoft.Compute/availabilitySets"),
	}

	if !cs.Properties.MasterProfile.HasAvailabilityZones() {
		if isManagedDisks {
			avSet.AvailabilitySetProperties = &compute.AvailabilitySetProperties{}
			if cs.Properties.MasterProfile.PlatformFaultDomainCount != nil {
				p := int32(*cs.Properties.MasterProfile.PlatformFaultDomainCount)
				avSet.PlatformFaultDomainCount = to.Int32Ptr(p)
			}
			if cs.Properties.MasterProfile.PlatformUpdateDomainCount != nil {
				p := int32(*cs.Properties.MasterProfile.PlatformUpdateDomainCount)
				avSet.PlatformUpdateDomainCount = to.Int32Ptr(p)
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

func createAgentAvailabilitySets(profile *api.AgentPoolProfile) AvailabilitySetARM {

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
	}

	avSet := compute.AvailabilitySet{
		Name:                      to.StringPtr(fmt.Sprintf("[variables('%sAvailabilitySet')]", profile.Name)),
		Location:                  to.StringPtr("[variables('location')]"),
		Type:                      to.StringPtr("Microsoft.Compute/availabilitySets"),
		AvailabilitySetProperties: &compute.AvailabilitySetProperties{},
	}

	if profile.IsManagedDisks() {
		if profile.PlatformFaultDomainCount != nil {
			p := int32(*profile.PlatformFaultDomainCount)
			avSet.PlatformFaultDomainCount = to.Int32Ptr(p)
		}
		if profile.PlatformUpdateDomainCount != nil {
			p := int32(*profile.PlatformUpdateDomainCount)
			avSet.PlatformUpdateDomainCount = to.Int32Ptr(p)
		}
		avSet.Sku = &compute.Sku{
			Name: to.StringPtr("Aligned"),
		}
	}

	return AvailabilitySetARM{
		ARMResource:     armResource,
		AvailabilitySet: avSet,
	}
}
