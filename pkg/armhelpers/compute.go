// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
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
