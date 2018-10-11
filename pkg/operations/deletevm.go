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
	logger.Infof("fetching VM: %s/%s", resourceGroup, name)
	vm, err := az.GetVirtualMachine(ctx, resourceGroup, name)
	if err != nil {
		logger.Errorf("failed to get VM: %s/%s: %s", resourceGroup, name, err.Error())
		return err
	}

	vhd := vm.VirtualMachineProperties.StorageProfile.OsDisk.Vhd
	managedDisk := vm.VirtualMachineProperties.StorageProfile.OsDisk.ManagedDisk
	if vhd == nil && managedDisk == nil {
		logger.Errorf("failed to get a valid os disk URI for VM: %s/%s", resourceGroup, name)

		return errors.New("os disk does not have a VHD URI")
	}

	osDiskName := vm.VirtualMachineProperties.StorageProfile.OsDisk.Name

	var nicName string
	nicID := (*vm.VirtualMachineProperties.NetworkProfile.NetworkInterfaces)[0].ID
	if nicID == nil {
		logger.Warnf("NIC ID is not set for VM (%s/%s)", resourceGroup, name)
	} else {
		nicName, err = utils.ResourceName(*nicID)
		if err != nil {
			return err
		}
		logger.Infof("found nic name for VM (%s/%s): %s", resourceGroup, name, nicName)
	}
	logger.Infof("deleting VM: %s/%s", resourceGroup, name)
	logger.Infof("waiting for vm deletion: %s/%s", resourceGroup, name)
	if err = az.DeleteVirtualMachine(ctx, resourceGroup, name); err != nil {
		return err
	}

	if len(nicName) > 0 {
		logger.Infof("deleting nic: %s/%s", resourceGroup, nicName)
		logger.Infof("waiting for nic deletion: %s/%s", resourceGroup, nicName)
		if err := az.DeleteNetworkInterface(ctx, resourceGroup, nicName); err != nil {
			return err
		}
	}

	if vhd != nil {
		accountName, vhdContainer, vhdBlob, err := utils.SplitBlobURI(*vhd.URI)
		if err != nil {
			return err
		}

		logger.Infof("found os disk storage reference: %s %s %s", accountName, vhdContainer, vhdBlob)

		as, err := az.GetStorageClient(ctx, resourceGroup, accountName)
		if err != nil {
			return err
		}

		logger.Infof("deleting blob: %s/%s", vhdContainer, vhdBlob)
		if err = as.DeleteBlob(vhdContainer, vhdBlob, &azStorage.DeleteBlobOptions{}); err != nil {
			return err
		}
	} else if managedDisk != nil {
		if osDiskName == nil {
			logger.Warnf("osDisk is not set for VM %s/%s", resourceGroup, name)
		} else {
			logger.Infof("deleting managed disk: %s/%s", resourceGroup, *osDiskName)
			if err = az.DeleteManagedDisk(ctx, resourceGroup, *osDiskName); err != nil {
				return err
			}
		}
	}

	if vm.Identity != nil {
		// Role assignments are not deleted if the VM is destroyed, so we must cleanup ourselves!
		// The role assignments should only be relevant if managed identities are used,
		// but always cleaning them up is easier than adding rule based logic here and there.
		scope := fmt.Sprintf(AADRoleResourceGroupScopeTemplate, subscriptionID, resourceGroup)
		logger.Infof("fetching roleAssignments: %s with principal %s", scope, *vm.Identity.PrincipalID)
		for vmRoleAssignmentsPage, err := az.ListRoleAssignmentsForPrincipal(ctx, scope, *vm.Identity.PrincipalID); vmRoleAssignmentsPage.NotDone(); err = vmRoleAssignmentsPage.Next() {
			if err != nil {
				logger.Errorf("failed to list role assignments: %s/%s: %s", scope, *vm.Identity.PrincipalID, err)
				return err
			}

			for _, roleAssignment := range vmRoleAssignmentsPage.Values() {
				logger.Infof("deleting role assignment: %s", *roleAssignment.ID)
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
