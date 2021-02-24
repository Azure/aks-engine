// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"time"

	"github.com/Azure/aks-engine/pkg/kubernetes"
	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-06-01/subscriptions"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"

	azStorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceSkusResultPage
type ResourceSkusResultPage interface {
	Next() error
	NextWithContext(ctx context.Context) (err error)
	NotDone() bool
	Response() compute.ResourceSkusResult
	Values() []compute.ResourceSku
}

// VirtualMachineListResultPage is an interface for compute.VirtualMachineListResultPage to aid in mocking
type VirtualMachineListResultPage interface {
	Next() error
	NotDone() bool
	Response() compute.VirtualMachineListResult
	Values() []compute.VirtualMachine
}

// VirtualMachineScaleSetListResultPage is an interface for compute.VirtualMachineScaleSetListResultPage to aid in mocking
type VirtualMachineScaleSetListResultPage interface {
	Next() error
	NextWithContext(ctx context.Context) (err error)
	NotDone() bool
	Response() compute.VirtualMachineScaleSetListResult
	Values() []compute.VirtualMachineScaleSet
}

// VirtualMachineScaleSetVMListResultPage is an interface for compute.VirtualMachineScaleSetListResultPage to aid in mocking
type VirtualMachineScaleSetVMListResultPage interface {
	Next() error
	NextWithContext(ctx context.Context) (err error)
	NotDone() bool
	Response() compute.VirtualMachineScaleSetVMListResult
	Values() []compute.VirtualMachineScaleSetVM
}

// ProviderListResultPage is an interface for resources.ProviderListResultPage to aid in mocking
type ProviderListResultPage interface {
	Next() error
	NextWithContext(ctx context.Context) (err error)
	NotDone() bool
	Response() resources.ProviderListResult
	Values() []resources.Provider
}

// DeploymentOperationsListResultPage is an interface for resources.DeploymentOperationsListResultPage to aid in mocking
type DeploymentOperationsListResultPage interface {
	Next() error
	NotDone() bool
	Response() resources.DeploymentOperationsListResult
	Values() []resources.DeploymentOperation
}

// RoleAssignmentListResultPage is an interface for authorization.RoleAssignmentListResultPage to aid in mocking
type RoleAssignmentListResultPage interface {
	Next() error
	NotDone() bool
	Response() authorization.RoleAssignmentListResult
	Values() []authorization.RoleAssignment
}

// DiskListPage is an interface for compute.DiskListPage to aid in mocking
type DiskListPage interface {
	Next() error
	NextWithContext(ctx context.Context) (err error)
	NotDone() bool
	Response() compute.DiskList
	Values() []compute.Disk
}

//VMImageFetcher is an extension of AKSEngine client allows us to operate on the virtual machine images in the environment
type VMImageFetcher interface {

	// ListVirtualMachineImages return a list of images
	ListVirtualMachineImages(ctx context.Context, location, publisherName, offer, skus string) (compute.ListVirtualMachineImageResource, error)

	// GetVirtualMachineImage return a virtual machine image
	GetVirtualMachineImage(ctx context.Context, location, publisherName, offer, skus, version string) (compute.VirtualMachineImage, error)
}

// AKSEngineClient is the interface used to talk to an Azure environment.
// This interface exposes just the subset of Azure APIs and clients needed for
// AKS Engine.
type AKSEngineClient interface {

	//AddAcceptLanguages sets the list of languages to accept on this request
	AddAcceptLanguages(languages []string)

	// AddAuxiliaryTokens sets the list of aux tokens to accept on this request
	AddAuxiliaryTokens(tokens []string)

	// RESOURCES

	// DeployTemplate can deploy a template into Azure ARM
	DeployTemplate(ctx context.Context, resourceGroup, name string, template, parameters map[string]interface{}) (resources.DeploymentExtended, error)

	// EnsureResourceGroup ensures the specified resource group exists in the specified location
	EnsureResourceGroup(ctx context.Context, resourceGroup, location string, managedBy *string) (*resources.Group, error)

	// ListLocations returns all the Azure locations to which AKS Engine can deploy
	ListLocations(ctx context.Context) (*[]subscriptions.Location, error)

	//
	// COMPUTE

	// ListResourceSkus lists Microsoft.Compute SKUs available for a subscription
	ListResourceSkus(ctx context.Context, filter string) (ResourceSkusResultPage, error)

	// ListVirtualMachines lists VM resources
	ListVirtualMachines(ctx context.Context, resourceGroup string) (VirtualMachineListResultPage, error)

	// GetVirtualMachine retrieves the specified virtual machine.
	GetVirtualMachine(ctx context.Context, resourceGroup, name string) (compute.VirtualMachine, error)

	// RestartVirtualMachine restarts the specified virtual machine.
	RestartVirtualMachine(ctx context.Context, resourceGroup, name string) error

	// DeleteVirtualMachine deletes the specified virtual machine.
	DeleteVirtualMachine(ctx context.Context, resourceGroup, name string) error

	// ListVirtualMachineScaleSets lists the VMSS resources in the resource group
	ListVirtualMachineScaleSets(ctx context.Context, resourceGroup string) (VirtualMachineScaleSetListResultPage, error)

	// RestartVirtualMachineScaleSets restarts the specified VMSS
	RestartVirtualMachineScaleSets(ctx context.Context, resourceGroup, virtualMachineScaleSet string, instanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) error

	// ListVirtualMachineScaleSetVMs lists the virtual machines contained in a VMSS
	ListVirtualMachineScaleSetVMs(ctx context.Context, resourceGroup, virtualMachineScaleSet string) (VirtualMachineScaleSetVMListResultPage, error)

	// DeleteVirtualMachineScaleSetVM deletes a VM in a VMSS
	DeleteVirtualMachineScaleSetVM(ctx context.Context, resourceGroup, virtualMachineScaleSet, instanceID string) error

	// SetVirtualMachineScaleSetCapacity sets the VMSS capacity
	SetVirtualMachineScaleSetCapacity(ctx context.Context, resourceGroup, virtualMachineScaleSet string, sku compute.Sku, location string) error

	// GetAvailabilitySet retrieves the specified VM availability set.
	GetAvailabilitySet(ctx context.Context, resourceGroup, availabilitySet string) (compute.AvailabilitySet, error)

	// GetAvailabilitySetFaultDomainCount returns the first platform fault domain count it finds from the
	// VM availability set IDs provided.
	GetAvailabilitySetFaultDomainCount(ctx context.Context, resourceGroup string, vmasIDs []string) (int, error)

	// GetVirtualMachinePowerState returns the virtual machine's PowerState status code
	GetVirtualMachinePowerState(ctx context.Context, resourceGroup, name string) (string, error)

	// GetVirtualMachineScaleSetInstancePowerState returns the virtual machine's PowerState status code
	GetVirtualMachineScaleSetInstancePowerState(ctx context.Context, resourceGroup, name, instanceID string) (string, error)

	//
	// STORAGE

	// GetStorageClient uses SRP to retrieve keys, and then an authenticated client for talking to the specified storage
	// account.
	GetStorageClient(ctx context.Context, resourceGroup, accountName string) (AKSStorageClient, error)

	//
	// NETWORK

	// DeleteNetworkInterface deletes the specified network interface.
	DeleteNetworkInterface(ctx context.Context, resourceGroup, nicName string) error

	//
	// GRAPH

	// CreateGraphAppliction creates an application via the graphrbac client
	CreateGraphApplication(ctx context.Context, applicationCreateParameters graphrbac.ApplicationCreateParameters) (graphrbac.Application, error)

	// CreateGraphPrincipal creates a service principal via the graphrbac client
	CreateGraphPrincipal(ctx context.Context, servicePrincipalCreateParameters graphrbac.ServicePrincipalCreateParameters) (graphrbac.ServicePrincipal, error)
	CreateApp(ctx context.Context, applicationName, applicationURL string, replyURLs *[]string, requiredResourceAccess *[]graphrbac.RequiredResourceAccess) (result graphrbac.Application, servicePrincipalObjectID, secret string, err error)
	DeleteApp(ctx context.Context, applicationName, applicationObjectID string) (autorest.Response, error)

	// User Assigned MSI
	//CreateUserAssignedID - Creates a user assigned msi.
	CreateUserAssignedID(location string, resourceGroup string, userAssignedID string) (*msi.Identity, error)

	// RBAC
	CreateRoleAssignment(ctx context.Context, scope string, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error)
	CreateRoleAssignmentSimple(ctx context.Context, applicationID, roleID string) error
	DeleteRoleAssignmentByID(ctx context.Context, roleAssignmentNameID string) (authorization.RoleAssignment, error)
	ListRoleAssignmentsForPrincipal(ctx context.Context, scope string, principalID string) (RoleAssignmentListResultPage, error)

	// MANAGED DISKS
	DeleteManagedDisk(ctx context.Context, resourceGroupName string, diskName string) error
	ListManagedDisksByResourceGroup(ctx context.Context, resourceGroupName string) (result DiskListPage, err error)

	GetKubernetesClient(apiserverURL, kubeConfig string, interval, timeout time.Duration) (kubernetes.Client, error)

	ListProviders(ctx context.Context) (ProviderListResultPage, error)

	// DEPLOYMENTS

	// ListDeploymentOperations gets all deployments operations for a deployment.
	ListDeploymentOperations(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result DeploymentOperationsListResultPage, err error)

	// Log Analytics

	// EnsureDefaultLogAnalyticsWorkspace ensures the default log analytics exists corresponding to specified location in current subscription
	EnsureDefaultLogAnalyticsWorkspace(ctx context.Context, resourceGroup, location string) (workspaceResourceID string, err error)

	// GetLogAnalyticsWorkspaceInfo gets the details about the workspace
	GetLogAnalyticsWorkspaceInfo(ctx context.Context, workspaceSubscriptionID, workspaceResourceGroup, workspaceName string) (workspaceID string, workspaceKey, workspaceLocation string, err error)

	// AddContainerInsightsSolution adds container insights solution for the specified log analytics workspace
	AddContainerInsightsSolution(ctx context.Context, workspaceSubscriptionID, workspaceResourceGroup, workspaceName, workspaceLocation string) (result bool, err error)
}

// AKSStorageClient interface models the azure storage client
type AKSStorageClient interface {
	// DeleteBlob deletes the specified blob in the specified container.
	DeleteBlob(containerName, blobName string, options *azStorage.DeleteBlobOptions) error
	// CreateContainer creates the CloudBlobContainer if it does not exist
	CreateContainer(containerName string, options *azStorage.CreateContainerOptions) (bool, error)
	// SaveBlockBlob initializes a block blob by taking the byte
	SaveBlockBlob(containerName, blobName string, b []byte, options *azStorage.PutBlobOptions) error
}
