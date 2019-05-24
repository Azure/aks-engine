// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"context"
	"fmt"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/armhelpers/utils"
	azStorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// AADRoleResourceGroupScopeTemplate is a template for a roleDefinition scope
	AADRoleResourceGroupScopeTemplate = "/subscriptions/%s/resourceGroups/%s"
)

// CleanDeleteVirtualMachine deletes a VM and any associated OS disk
func CleanDeleteVirtualMachine(az armhelpers.AKSEngineClient, logger *log.Entry, subscriptionID, resourceGroup, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	logger.Debugf("fetching VM %s in resource group %s", name, resourceGroup)
	vm, err := az.GetVirtualMachine(ctx, resourceGroup, name)
	if err != nil {
		logger.Errorf("failed to get VM %s in resource group %s: %s", name, resourceGroup, err.Error())
		return err
	}

	vhd := vm.VirtualMachineProperties.StorageProfile.OsDisk.Vhd
	managedDisk := vm.VirtualMachineProperties.StorageProfile.OsDisk.ManagedDisk
	if vhd == nil && managedDisk == nil {
		logger.Errorf("failed to get a valid OS disk URI for VM %s in resource group %s", name, resourceGroup)

		return errors.New("OS disk does not have a VHD URI")
	}

	osDiskName := vm.VirtualMachineProperties.StorageProfile.OsDisk.Name

	var nicName string
	nicID := (*vm.VirtualMachineProperties.NetworkProfile.NetworkInterfaces)[0].ID
	if nicID == nil {
		logger.Warnf("NIC ID is not set for VM %s in resource group %s)", name, resourceGroup)
	} else {
		nicName, err = utils.ResourceName(*nicID)
		if err != nil {
			return err
		}
		logger.Debugf("found NIC name \"%s\" for VM %s in resource group %s", nicName, name, resourceGroup)
	}
	logger.Infof("deleting VM %s in resource group %s ...", name, resourceGroup)
	if err = az.DeleteVirtualMachine(ctx, resourceGroup, name); err != nil {
		return err
	}

	if len(nicName) > 0 {
		logger.Infof("deleting NIC %s in resource group %s ...", nicName, resourceGroup)
		if err = az.DeleteNetworkInterface(ctx, resourceGroup, nicName); err != nil {
			return err
		}
	}

	if vhd != nil {
		var accountName, vhdContainer, vhdBlob string
		accountName, vhdContainer, vhdBlob, err = utils.SplitBlobURI(*vhd.URI)
		if err != nil {
			return err
		}

		logger.Debugf("found OS disk storage reference: %s:%s/%s", accountName, vhdContainer, vhdBlob)

		var as armhelpers.AKSStorageClient
		as, err = az.GetStorageClient(ctx, resourceGroup, accountName)
		if err != nil {
			return err
		}

		logger.Infof("deleting blob %s/%s ...", vhdContainer, vhdBlob)
		if err = as.DeleteBlob(vhdContainer, vhdBlob, &azStorage.DeleteBlobOptions{}); err != nil {
			return err
		}
	} else if managedDisk != nil {
		if osDiskName == nil {
			logger.Warnf("managed disk Name is not set for VM %s in resource group %s", name, resourceGroup)
		} else {
			logger.Infof("deleting managed disk %s in resource group %s ...", *osDiskName, resourceGroup)
			if err = az.DeleteManagedDisk(ctx, resourceGroup, *osDiskName); err != nil {
				return err
			}
		}
	}

	if vm.Identity != nil && vm.Identity.PrincipalID != nil {
		// Role assignments are not deleted if the VM is destroyed, so we must cleanup ourselves!
		// The role assignments should only be relevant if managed identities are used,
		// but always cleaning them up is easier than adding rule based logic here and there.
		scope := fmt.Sprintf(AADRoleResourceGroupScopeTemplate, subscriptionID, resourceGroup)
		logger.Debugf("fetching role assignments: %s with principal %s", scope, *vm.Identity.PrincipalID)
		for vmRoleAssignmentsPage, err := az.ListRoleAssignmentsForPrincipal(ctx, scope, *vm.Identity.PrincipalID); vmRoleAssignmentsPage.NotDone(); err = vmRoleAssignmentsPage.Next() {
			if err != nil {
				logger.Errorf("failed to list role assignments: %s/%s: %s", scope, *vm.Identity.PrincipalID, err)
				return err
			}

			for _, roleAssignment := range vmRoleAssignmentsPage.Values() {
				logger.Infof("deleting role assignment %s ...", *roleAssignment.ID)
				_, deleteRoleAssignmentErr := az.DeleteRoleAssignmentByID(ctx, *roleAssignment.ID)
				if deleteRoleAssignmentErr != nil {
					logger.Errorf("failed to delete role assignment: %s: %s", *roleAssignment.ID, deleteRoleAssignmentErr.Error())
					return deleteRoleAssignmentErr
				}
			}
		}
	}

	return nil
}
