// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
)

// ListVirtualMachines returns (the first page of) the machines in the specified resource group.
func (az *AzureClient) ListVirtualMachines(ctx context.Context, resourceGroup string) (VirtualMachineListResultPage, error) {
	page, err := az.virtualMachinesClient.List(ctx, resourceGroup)
	return &page, err
}

// GetVirtualMachine returns the specified machine in the specified resource group.
func (az *AzureClient) GetVirtualMachine(ctx context.Context, resourceGroup, name string) (compute.VirtualMachine, error) {
	return az.virtualMachinesClient.Get(ctx, resourceGroup, name, "")
}

// RestartVirtualMachine restarts the specified virtual machine.
func (az *AzureClient) RestartVirtualMachine(ctx context.Context, resourceGroup, name string) error {
	future, err := az.virtualMachinesClient.Restart(ctx, resourceGroup, name)
	if err != nil {
		return err
	}

	if err = future.WaitForCompletionRef(ctx, az.virtualMachinesClient.Client); err != nil {
		return err
	}

	_, err = future.Result(az.virtualMachinesClient)
	return err
}

// DeleteVirtualMachine handles deletion of a CRP/VMAS VM (aka, not a VMSS VM).
func (az *AzureClient) DeleteVirtualMachine(ctx context.Context, resourceGroup, name string) error {
	future, err := az.virtualMachinesClient.Delete(ctx, resourceGroup, name)
	if err != nil {
		return err
	}

	if err = future.WaitForCompletionRef(ctx, az.virtualMachinesClient.Client); err != nil {
		return err
	}

	_, err = future.Result(az.virtualMachinesClient)
	return err
}

// ListVirtualMachineScaleSets returns (the first page of) the VMSS resources in the specified resource group.
func (az *AzureClient) ListVirtualMachineScaleSets(ctx context.Context, resourceGroup string) (VirtualMachineScaleSetListResultPage, error) {
	page, err := az.virtualMachineScaleSetsClient.List(ctx, resourceGroup)
	return &page, err
}

// RestartVirtualMachineScaleSets restarts the specified VMSS
func (az *AzureClient) RestartVirtualMachineScaleSets(ctx context.Context, resourceGroup string, virtualMachineScaleSet string, instanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) error {
	future, err := az.virtualMachineScaleSetsClient.Restart(ctx, resourceGroup, virtualMachineScaleSet, instanceIDs)
	if err != nil {
		return err
	}

	if err = future.WaitForCompletionRef(ctx, az.virtualMachineScaleSetsClient.Client); err != nil {
		return err
	}

	_, err = future.Result(az.virtualMachineScaleSetsClient)
	return err
}

// ListVirtualMachineScaleSetVMs returns the list of VMs per VMSS
func (az *AzureClient) ListVirtualMachineScaleSetVMs(ctx context.Context, resourceGroup, virtualMachineScaleSet string) (VirtualMachineScaleSetVMListResultPage, error) {
	page, err := az.virtualMachineScaleSetVMsClient.List(ctx, resourceGroup, virtualMachineScaleSet, "", "", "")
	return &page, err
}

// DeleteVirtualMachineScaleSetVM deletes a VM in a VMSS
func (az *AzureClient) DeleteVirtualMachineScaleSetVM(ctx context.Context, resourceGroup, virtualMachineScaleSet, instanceID string) error {
	future, err := az.virtualMachineScaleSetVMsClient.Delete(ctx, resourceGroup, virtualMachineScaleSet, instanceID)
	if err != nil {
		return err
	}

	if err = future.WaitForCompletionRef(ctx, az.virtualMachineScaleSetVMsClient.Client); err != nil {
		return err
	}

	_, err = future.Result(az.virtualMachineScaleSetVMsClient)
	return err
}

// DeleteVirtualMachineScaleSet deletes an entire VM Scale Set.
func (az *AzureClient) DeleteVirtualMachineScaleSet(ctx context.Context, resourceGroup, vmssName string) error {
	future, err := az.virtualMachineScaleSetsClient.Delete(ctx, resourceGroup, vmssName)
	if err != nil {
		return err
	}
	if err = future.WaitForCompletionRef(ctx, az.virtualMachineScaleSetsClient.Client); err != nil {
		return err
	}
	_, err = future.Result(az.virtualMachineScaleSetsClient)
	return err
}

// SetVirtualMachineScaleSetCapacity sets the VMSS capacity
func (az *AzureClient) SetVirtualMachineScaleSetCapacity(ctx context.Context, resourceGroup, virtualMachineScaleSet string, sku compute.Sku, location string) error {
	future, err := az.virtualMachineScaleSetsClient.CreateOrUpdate(
		ctx,
		resourceGroup,
		virtualMachineScaleSet,
		compute.VirtualMachineScaleSet{
			Location: &location,
			Sku:      &sku,
		})
	if err != nil {
		return err
	}

	if err = future.WaitForCompletionRef(ctx, az.virtualMachineScaleSetsClient.Client); err != nil {
		return err
	}

	_, err = future.Result(az.virtualMachineScaleSetsClient)
	return err
}

// GetAvailabilitySet retrieves the specified VM availability set.
func (az *AzureClient) GetAvailabilitySet(ctx context.Context, resourceGroup, availabilitySetName string) (compute.AvailabilitySet, error) {
	return az.availabilitySetsClient.Get(ctx, resourceGroup, availabilitySetName)
}

// GetAvailabilitySetFaultDomainCount returns the first existing fault domain count it finds from the IDs provided.
func (az *AzureClient) GetAvailabilitySetFaultDomainCount(ctx context.Context, resourceGroup string, vmasIDs []string) (int, error) {
	var count int
	if len(vmasIDs) > 0 {
		id := vmasIDs[0]
		// extract the last element of the id for VMAS name
		ss := strings.Split(id, "/")
		name := ss[len(ss)-1]
		vmas, err := az.GetAvailabilitySet(ctx, resourceGroup, name)
		if err != nil {
			return 0, err
		}
		// Assume that all VMASes in the cluster share a value for platformFaultDomainCount
		count = int(*vmas.AvailabilitySetProperties.PlatformFaultDomainCount)
	}
	return count, nil
}
