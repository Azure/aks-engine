// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
)

func TestCreateAvailabilitySet(t *testing.T) {

	//Test AvSet without ManagedDisk
	cs := &api.ContainerService{
		Properties: &api.Properties{
			MasterProfile: &api.MasterProfile{
				AvailabilityZones: []string{},
			},
		},
	}

	avSet := CreateAvailabilitySet(cs, false)

	expectedAvSet := AvailabilitySetARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
		},
		AvailabilitySet: compute.AvailabilitySet{
			Name:     to.StringPtr("[variables('masterAvailabilitySet')]"),
			Location: to.StringPtr("[variables('location')]"),
			Type:     to.StringPtr("Microsoft.Compute/availabilitySets"),
		},
	}

	diff := cmp.Diff(avSet, expectedAvSet)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

	//Test AvSet with ManagedDisk

	cs = &api.ContainerService{
		Properties: &api.Properties{
			MasterProfile: &api.MasterProfile{},
		},
	}

	avSet = CreateAvailabilitySet(cs, true)

	expectedAvSet = AvailabilitySetARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
		},
		AvailabilitySet: compute.AvailabilitySet{
			Name:     to.StringPtr("[variables('masterAvailabilitySet')]"),
			Location: to.StringPtr("[variables('location')]"),
			Type:     to.StringPtr("Microsoft.Compute/availabilitySets"),
			Sku: &compute.Sku{
				Name: to.StringPtr("Aligned"),
			},
			AvailabilitySetProperties: &compute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(2),
				PlatformUpdateDomainCount: to.Int32Ptr(3),
			},
		},
	}

	diff = cmp.Diff(avSet, expectedAvSet)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

	//Test AvSet with StorageAccount
	cs = &api.ContainerService{
		Properties: &api.Properties{
			MasterProfile: &api.MasterProfile{
				StorageProfile: api.StorageAccount,
			},
		},
	}

	avSet = CreateAvailabilitySet(cs, false)

	expectedAvSet = AvailabilitySetARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
		},
		AvailabilitySet: compute.AvailabilitySet{
			Name:                      to.StringPtr("[variables('masterAvailabilitySet')]"),
			Location:                  to.StringPtr("[variables('location')]"),
			Type:                      to.StringPtr("Microsoft.Compute/availabilitySets"),
			AvailabilitySetProperties: &compute.AvailabilitySetProperties{},
		},
	}

	diff = cmp.Diff(avSet, expectedAvSet)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

}

func TestCreateAgentAvailabilitySets(t *testing.T) {
	//Test AvSet without ManagedDisk
	profile := &api.AgentPoolProfile{
		Name:           "foobar",
		StorageProfile: api.StorageAccount,
	}

	avSet := createAgentAvailabilitySets(profile)

	expectedAvSet := AvailabilitySetARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
		},
		AvailabilitySet: compute.AvailabilitySet{
			Name:                      to.StringPtr("[variables('foobarAvailabilitySet')]"),
			Location:                  to.StringPtr("[variables('location')]"),
			Type:                      to.StringPtr("Microsoft.Compute/availabilitySets"),
			AvailabilitySetProperties: &compute.AvailabilitySetProperties{},
		},
	}

	diff := cmp.Diff(avSet, expectedAvSet)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

	//Test AvSet wit ManagedDisk
	profile = &api.AgentPoolProfile{
		Name:           "foobar",
		StorageProfile: api.ManagedDisks,
	}

	avSet = createAgentAvailabilitySets(profile)

	expectedAvSet = AvailabilitySetARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
		},
		AvailabilitySet: compute.AvailabilitySet{
			Name:     to.StringPtr("[variables('foobarAvailabilitySet')]"),
			Location: to.StringPtr("[variables('location')]"),
			Type:     to.StringPtr("Microsoft.Compute/availabilitySets"),
			AvailabilitySetProperties: &compute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(2),
				PlatformUpdateDomainCount: to.Int32Ptr(3),
			},
			Sku: &compute.Sku{
				Name: to.StringPtr("Aligned"),
			},
		},
	}

	diff = cmp.Diff(avSet, expectedAvSet)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

}
