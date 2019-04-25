// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/google/go-cmp/cmp"
)

func TestListVirtualMachineScaleSets(t *testing.T) {

	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()

	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	mc.RegisterListVirtualMachineScaleSets()

	list := &VirtualMachineScaleSetListValues{}
	err = unmarshalFromString(mc.ResponseListVirtualMachineScaleSets, &list)
	if err != nil {
		t.Error(err)
	}

	listExpected := []azcompute.VirtualMachineScaleSet{}
	DeepCopy(&listExpected, list.Value)

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
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()

	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	mc.RegisterListVirtualMachineScaleSetVMs()

	list := &VirtualMachineScaleSetVMValues{}
	err = unmarshalFromString(mc.ResponseListVirtualMachineScaleSetVMs, &list)
	if err != nil {
		t.Error(err)
	}

	listExpected := []azcompute.VirtualMachineScaleSetVM{}
	DeepCopy(&listExpected, list.Value)

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
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	mc.RegisterListVirtualMachines()

	list := &VirtualMachineVMValues{}
	err = unmarshalFromString(mc.ResponseListVirtualMachines, &list)
	if err != nil {
		t.Error(err)
	}

	listExpected := []azcompute.VirtualMachine{}
	DeepCopy(&listExpected, list.Value)

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
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()

	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	mc.RegisterGetVirtualMachine()

	vm := compute.VirtualMachine{}
	err = unmarshalFromString(mc.ResponseGetVirtualMachine, &vm)
	if err != nil {
		t.Error(err)
	}

	vmExpected := azcompute.VirtualMachine{}
	DeepCopy(&vmExpected, vm)
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
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	mc.RegisterDeleteVirtualMachine()

	err = azureClient.DeleteVirtualMachine(context.Background(), resourceGroup, virtualMachineName)
	if err != nil {
		t.Error(err)
	}
}
