// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/google/go-cmp/cmp"
)

func TestListVirtualMachineScaleSets(t *testing.T) {

	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterListVirtualMachineScaleSets()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()

	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	list := &VirtualMachineScaleSetListValues{}
	err = unmarshalFromString(mc.ResponseListVirtualMachineScaleSets, &list)
	if err != nil {
		t.Error(err)
	}

	listExpected := []azcompute.VirtualMachineScaleSet{}
	if err := DeepCopy(&listExpected, list.Value); err != nil {
		t.Error(err)
	}

	for page, err := azureClient.ListVirtualMachineScaleSets(context.Background(), resourceGroup); page.NotDone(); err = page.Next() {
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(page.Values(), listExpected); diff != "" {
			t.Errorf("Fail to compare, Virtual Machine Scale Set %q", diff)
		}
	}
}

func TestListVirtualMachineScaleSetVMs(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterListVirtualMachineScaleSetVMs()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()

	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	list := &VirtualMachineScaleSetVMValues{}
	err = unmarshalFromString(mc.ResponseListVirtualMachineScaleSetVMs, &list)
	if err != nil {
		t.Error(err)
	}

	listExpected := []azcompute.VirtualMachineScaleSetVM{}
	if err := DeepCopy(&listExpected, list.Value); err != nil {
		t.Error(err)
	}

	for page, err := azureClient.ListVirtualMachineScaleSetVMs(context.Background(), resourceGroup, virtualMachineScaleSetName); page.NotDone(); err = page.Next() {
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(page.Values(), listExpected); diff != "" {
			t.Errorf("Fail to compare, Virtual Machine Scale Set VMs %q", diff)
		}
	}
}

func TestListVirtualMachines(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterListVirtualMachines()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	list := &VirtualMachineVMValues{}
	err = unmarshalFromString(mc.ResponseListVirtualMachines, &list)
	if err != nil {
		t.Error(err)
	}

	listExpected := []azcompute.VirtualMachine{}
	if err := DeepCopy(&listExpected, list.Value); err != nil {
		t.Error(err)
	}

	for page, err := azureClient.ListVirtualMachines(context.Background(), resourceGroup); page.NotDone(); err = page.Next() {
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(page.Values(), listExpected); diff != "" {
			t.Errorf("Fail to compare, Virtual Machines %q", diff)
		}
	}
}

func TestGetVirtualMachine(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterVirtualMachineEndpoint()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	vm := compute.VirtualMachine{}
	err = unmarshalFromString(mc.ResponseGetVirtualMachine, &vm)
	if err != nil {
		t.Error(err)
	}

	vmExpected := azcompute.VirtualMachine{}
	if err = DeepCopy(&vmExpected, vm); err != nil {
		t.Error(err)
	}
	vmActual, err := azureClient.GetVirtualMachine(context.Background(), resourceGroup, virtualMachineName)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(vmActual.VirtualMachineProperties, vmExpected.VirtualMachineProperties); diff != "" {
		t.Errorf("Fail to compare, Virtual Machine VirtualMachineProperties %q", diff)
	}
	if diff := cmp.Diff(vmActual.Name, vmExpected.Name); diff != "" {
		t.Errorf("Fail to compare, Virtual Machine Name %q", diff)
	}
	if diff := cmp.Diff(vmActual.Tags, vmExpected.Tags); diff != "" {
		t.Errorf("Fail to compare, Virtual Machine Tags %q", diff)
	}
	if diff := cmp.Diff(vmActual.Location, vmExpected.Location); diff != "" {
		t.Errorf("Fail to compare, Virtual Machine Location %q", diff)
	}

}
func TestDeleteVirtualMachine(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterVirtualMachineEndpoint()
	mc.RegisterDeleteOperation()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	err = azureClient.DeleteVirtualMachine(context.Background(), resourceGroup, virtualMachineName)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAvailabilitySet(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterGetAvailabilitySet()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()

	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	vmas, err := azureClient.GetAvailabilitySet(context.Background(), resourceGroup, virtualMachineAvailabilitySetName)
	if err != nil {
		t.Fatalf("can't get availability set: %s", err)
	}

	var expected int32 = 3
	if *vmas.PlatformFaultDomainCount != expected {
		t.Fatalf("expected PlatformFaultDomainCount of %d but got %v", expected, *vmas.PlatformFaultDomainCount)
	}

	if vmas.ProximityPlacementGroup != nil && vmas.ProximityPlacementGroup.ID != nil {
		t.Fatalf("expected ProximityPlacementGroup of %q but got %v", "", *vmas.ProximityPlacementGroup.ID)
	}
	if *vmas.PlatformUpdateDomainCount != expected {
		t.Fatalf("expected PlatformUpdateDomainCount of %d but got %v", expected, *vmas.PlatformUpdateDomainCount)
	}
	l := "eastus"
	if *vmas.Location != l {
		t.Fatalf("expected Location of %s but got %v", l, *vmas.Location)
	}
}

func TestGetAvailabilitySetFaultDomainCount(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterGetAvailabilitySetFaultDomainCount()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()

	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	count, err := azureClient.GetAvailabilitySetFaultDomainCount(context.Background(), resourceGroup, []string{"id1", "id2"})
	if err != nil {
		t.Fatalf("can't get availability set platform fault domain count: %s", err)
	}

	expected := 3
	if count != expected {
		t.Fatalf("platform fault domain count: expected %d but got %d", expected, count)
	}
}
