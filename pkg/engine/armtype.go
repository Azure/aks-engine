// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
)

type ARMResource struct {
	APIVersion string            `json:"apiVersion,omitempty"`
	Copy       map[string]string `json:"copy,omitempty"`
	DependsOn  []string          `json:"dependsOn,omitempty"`
}

func (arm ARMResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(arm)
}

type VirtualMachineARM struct {
	ARMResource
	compute.VirtualMachine
}

type VirtualMachineScaleSetARM struct {
	ARMResource
	compute.VirtualMachineScaleSet
}

type VirtualMachineExtensionARM struct {
	ARMResource
	compute.VirtualMachineExtension
}

type RoleAssignmentARM struct {
	ARMResource
	authorization.RoleAssignment
}

type AvailabilitySetARM struct {
	ARMResource
	compute.AvailabilitySet
}

type StorageAccountARM struct {
	ARMResource
	storage.Account
}

type VirtualNetworkARM struct {
	ARMResource
	network.VirtualNetwork
}

type NetworkSecurityGroupARM struct {
	ARMResource
	network.SecurityGroup
}

type RouteTableARM struct {
	ARMResource
	network.RouteTable
}

type PublicIPAddressARM struct {
	ARMResource
	network.PublicIPAddress
}

type LoadBalancerARM struct {
	ARMResource
	network.LoadBalancer
}

type NetworkInterfaceARM struct {
	ARMResource
	network.Interface
}

type DocumentDBAccountARM struct {
	ARMResource
	documentdb.DatabaseAccountCreateUpdateParameters
}

type KeyVaultARM struct {
	ARMResource
	keyvault.Vault
}

type UserAssignedIdentitiesARM struct {
	ARMResource
	msi.Identity
}

type ImageARM struct {
	ARMResource
	compute.Image
}
