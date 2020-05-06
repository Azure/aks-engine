// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	log "github.com/sirupsen/logrus"
)

// ListVirtualMachines returns (the first page of) the machines in the specified resource group.
func (az *AzureClient) ListVirtualMachines(ctx context.Context, resourceGroup string) (armhelpers.VirtualMachineListResultPage, error) {
	page, err := az.virtualMachinesClient.List(ctx, resourceGroup)
	c := VirtualMachineListResultPageClient{
		vmlrp: page,
		err:   err,
	}
	return &c, err
}

// GetVirtualMachine returns the specified machine in the specified resource group.
func (az *AzureClient) GetVirtualMachine(ctx context.Context, resourceGroup, name string) (azcompute.VirtualMachine, error) {
	vm, err := az.virtualMachinesClient.Get(ctx, resourceGroup, name, "")
	azVM := azcompute.VirtualMachine{}
	if err != nil {
		return azVM, fmt.Errorf("fail to get virtual machine, %s", err)
	}
	err = DeepCopy(&azVM, vm)
	if err != nil {
		return azVM, fmt.Errorf("fail to convert virtual machine, %s", err)
	}
	return azVM, err
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
func (az *AzureClient) ListVirtualMachineScaleSets(ctx context.Context, resourceGroup string) (armhelpers.VirtualMachineScaleSetListResultPage, error) {
	page, err := az.virtualMachineScaleSetsClient.List(ctx, resourceGroup)
	c := VirtualMachineScaleSetListResultPageClient{
		vmsslrp: page,
		err:     err,
	}
	return &c, err
}

// RestartVirtualMachineScaleSets restarts the specified VMSS
func (az *AzureClient) RestartVirtualMachineScaleSets(ctx context.Context, resourceGroup string, virtualMachineScaleSet string, instanceIDs *azcompute.VirtualMachineScaleSetVMInstanceIDs) error {
	ids := compute.VirtualMachineScaleSetVMInstanceIDs{}
	err := DeepCopy(&ids, instanceIDs)
	if err != nil {
		return fmt.Errorf("fail to convert instance IDs, %v", err)
	}
	future, err := az.virtualMachineScaleSetsClient.Restart(ctx, resourceGroup, virtualMachineScaleSet, &ids)
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
func (az *AzureClient) ListVirtualMachineScaleSetVMs(ctx context.Context, resourceGroup, virtualMachineScaleSet string) (armhelpers.VirtualMachineScaleSetVMListResultPage, error) {
	page, err := az.virtualMachineScaleSetVMsClient.List(ctx, resourceGroup, virtualMachineScaleSet, "", "", "")
	c := VirtualMachineScaleSetVMListResultPageClient{
		vmssvlrp: page,
		err:      err,
	}
	return &c, err
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
func (az *AzureClient) SetVirtualMachineScaleSetCapacity(ctx context.Context, resourceGroup, virtualMachineScaleSet string, sku azcompute.Sku, location string) error {
	s := compute.Sku{}
	err := DeepCopy(&s, sku)
	if err != nil {
		return fmt.Errorf("fail to convert SKU, %v", err)
	}
	future, err := az.virtualMachineScaleSetsClient.CreateOrUpdate(
		ctx,
		resourceGroup,
		virtualMachineScaleSet,
		compute.VirtualMachineScaleSet{
			Location: &location,
			Sku:      &s,
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
func (az *AzureClient) GetAvailabilitySet(ctx context.Context, resourceGroup, availabilitySetName string) (azcompute.AvailabilitySet, error) {
	azVMAS := azcompute.AvailabilitySet{}
	vmas, err := az.availabilitySetsClient.Get(ctx, resourceGroup, availabilitySetName)
	if err != nil {
		log.Printf("fail to get availability set, %v", err)
		return azVMAS, err
	}
	if err = DeepCopy(&azVMAS, vmas); err != nil {
		log.Printf("fail to convert availability set, %v", err)
		return azVMAS, err
	}
	return azVMAS, nil
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
