// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	sysauth "github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-01-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"

	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

// ARMResource contains the fields that are common to all Azure Resource Manager objects.
type ARMResource struct {
	APIVersion string            `json:"apiVersion,omitempty"`
	Copy       map[string]string `json:"copy,omitempty"`
	DependsOn  []string          `json:"dependsOn,omitempty"`
}

// DeploymentARMResource is an alias for the ARMResource type to avoid MarshalJSON override
type DeploymentARMResource ARMResource

// MarshalJSON is the custom marshaler for an ARMResource.
func (arm ARMResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(arm)
}

// VirtualMachineARM embeds the ARMResource type in compute.VirtualMachine.
type VirtualMachineARM struct {
	ARMResource
	compute.VirtualMachine
}

// VirtualMachineScaleSetARM embeds the ARMResource type in compute.VirtualMachineScaleSet.
type VirtualMachineScaleSetARM struct {
	ARMResource
	compute.VirtualMachineScaleSet
}

// VirtualMachineExtensionARM embeds the ARMResource type in compute.VirtualMachineExtension.
type VirtualMachineExtensionARM struct {
	ARMResource
	compute.VirtualMachineExtension
}

// RoleAssignmentARM embeds the ARMResource type in authorization.RoleAssignment.
type RoleAssignmentARM struct {
	ARMResource
	authorization.RoleAssignment
}

// AvailabilitySetARM embeds the ARMResource type in compute.AvailabilitySet.
type AvailabilitySetARM struct {
	ARMResource
	compute.AvailabilitySet
}

// MarshalJSON is the custom marshaler for an AvailabilitySetARM.
// It acts as a decorator by replacing the JSON field "platformFaultDomainCount"
// with an ARM expression if the value was not set.
func (a AvailabilitySetARM) MarshalJSON() ([]byte, error) {
	// alias the type to avoid infinite recursion in marshaling
	type Alias AvailabilitySetARM
	bytes, err := json.Marshal((Alias)(a))
	if err != nil {
		return nil, err
	}

	if a.AvailabilitySet.PlatformFaultDomainCount != nil {
		return bytes, nil
	}

	// armExpr is evaluated by Azure Resource Manager at deployment time:
	//   if location is in the three-fault-domain list, return 3
	//   else if location is "canary" (testing), return 1
	//   else return 2
	// NOTE: use fault_domains_expr.py to update this ARM expression.
	armExpr := `"[
	if( contains(
	      split('canadacentral,centralus,eastus,eastus2,northcentralus,northeurope,southcentralus,westeurope,westus', ','),
	        variables('location') ),
	  3,
	if( equals('centraluseuap', variables('location') ),
	  1,
	  2
	))]"`
	// strip all whitespace
	armExpr = strings.Join(strings.Fields(armExpr), "")

	// insert ARM expression for platformFaultDomainCount as the first JSON property
	// NOTE: this relies on this field being omitted in JSON when its value is nil.
	re := regexp.MustCompile(`"properties" *: *{ *"`)
	s := re.ReplaceAllLiteralString(string(bytes), `"properties":{"platformFaultDomainCount":`+armExpr+`,"`)

	return []byte(s), nil
}

// StorageAccountARM embeds the ARMResource type in storage.Account.
type StorageAccountARM struct {
	ARMResource
	storage.Account
}

// SystemRoleAssignmentARM embeds the ARMResource type in authorization.SystemRoleAssignment(2018-01-01-preview).
type SystemRoleAssignmentARM struct {
	ARMResource
	sysauth.RoleAssignment
}

// VirtualNetworkARM embeds the ARMResource type in network.VirtualNetwork.
type VirtualNetworkARM struct {
	ARMResource
	network.VirtualNetwork
}

// NetworkSecurityGroupARM embeds the ARMResource type in network.SecurityGroup.
type NetworkSecurityGroupARM struct {
	ARMResource
	network.SecurityGroup
}

// RouteTableARM embeds the ARMResource type in network.RouteTable.
type RouteTableARM struct {
	ARMResource
	network.RouteTable
}

// PublicIPAddressARM embeds the ARMResource type in network.PublicIPAddress.
type PublicIPAddressARM struct {
	ARMResource
	network.PublicIPAddress
}

// LoadBalancerARM embeds the ARMResource type in network.LoadBalancer.
type LoadBalancerARM struct {
	ARMResource
	network.LoadBalancer
}

// ApplicationGatewayARM embeds the ARMResource type in network.ApplicationGateway.
type ApplicationGatewayARM struct {
	ARMResource
	network.ApplicationGateway
}

// NetworkInterfaceARM embeds the ARMResource type in network.Interface.
type NetworkInterfaceARM struct {
	ARMResource
	network.Interface
}

// DocumentDBAccountARM embeds the ARMResource type in documentdb.DatabaseAccountCreateUpdateParameters.
type DocumentDBAccountARM struct {
	ARMResource
	documentdb.DatabaseAccountCreateUpdateParameters
}

// KeyVaultARM embeds the ARMResource type in keyvault.Vault.
type KeyVaultARM struct {
	ARMResource
	keyvault.Vault
}

// UserAssignedIdentitiesARM embeds the ARMResource type in msi.Identity.
type UserAssignedIdentitiesARM struct {
	ARMResource
	msi.Identity
}

// ImageARM embeds the ARMResource type in compute.Image.
type ImageARM struct {
	ARMResource
	compute.Image
}

// DeploymentARM embeds the ARMResource type in resources.DeploymentExtended.
type DeploymentARM struct {
	DeploymentARMResource
	resources.DeploymentExtended
}

// TODO: Should we skip this type and add the `ResourceGroup` field directly to `DeploymentARM` ?
// DeploymentWithResourceGroupARM is like `DeploymentARM` but includes `ResourceGroup`.
type DeploymentWithResourceGroupARM struct {
	DeploymentARMResource
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	resources.DeploymentExtended
}
