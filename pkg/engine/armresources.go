// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func GenerateARMResources(cs *api.ContainerService) []interface{} {
	var armResources []interface{}

	deploymentTelemetryEnabled := cs.Properties.FeatureFlags.IsFeatureEnabled("EnableTelemetry")
	isAzureStack := cs.Properties.IsAzureStackCloud()

	azureTelemetryPID := cs.GetCloudSpecConfig().KubernetesSpecConfig.AzureTelemetryPID

	if deploymentTelemetryEnabled {
		if isAzureStack {
			deploymentResource := createAzureStackTelemetry(azureTelemetryPID)
			armResources = append(armResources, deploymentResource)
		}
	}

	var useManagedIdentity, userAssignedIDEnabled bool
	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig

	if kubernetesConfig != nil {
		useManagedIdentity = kubernetesConfig.UseManagedIdentity
		userAssignedIDEnabled = useManagedIdentity && kubernetesConfig.UserAssignedID != ""
	}

	isHostedMaster := cs.Properties.IsHostedMasterProfile()
	if userAssignedIDEnabled {
		userAssignedID := createUserAssignedIdentities()
		var msiRoleAssignment RoleAssignmentARM
		if isHostedMaster {
			msiRoleAssignment = createMSIRoleAssignment(IdentityReaderRole)
		} else {
			msiRoleAssignment = createMSIRoleAssignment(IdentityContributorRole)
		}
		armResources = append(armResources, userAssignedID, msiRoleAssignment)
	}

	// Create the Standard Load Balancer resource spec, so long as:
	// - we are not in an AKS template generation flow
	// - there are no node pools configured with LoadBalancerBackendAddressPoolIDs
	//    - i.e., user-provided LoadBalancerBackendAddressPoolIDs is not compatible w/ this Standard LB spec,
	//      which assumes *all vms in all node pools* as backend pool members
	if cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku &&
		!isHostedMaster &&
		!cs.Properties.AnyAgentHasLoadBalancerBackendAddressPoolIDs() {
		isForMaster := false
		includeDNS := false
		publicIPAddress := CreatePublicIPAddress(isForMaster, includeDNS)
		loadBalancer := CreateStandardLoadBalancerForNodePools(cs.Properties, true)
		armResources = append(armResources, publicIPAddress, loadBalancer)
	}

	profiles := cs.Properties.AgentPoolProfiles

	for _, profile := range profiles {

		if profile.IsWindows() {
			if cs.Properties.WindowsProfile.HasCustomImage() {
				// Create Image resource from VHD if requestesd
				armResources = append(armResources, createWindowsImage(profile))
			}
		}

		if profile.IsVirtualMachineScaleSets() {
			if useManagedIdentity && !userAssignedIDEnabled {
				armResources = append(armResources, createAgentVMSSSysRoleAssignment(profile))
			}
			armResources = append(armResources, CreateAgentVMSS(cs, profile))
		} else {
			agentVMASResources := createKubernetesAgentVMASResources(cs, profile)
			armResources = append(armResources, agentVMASResources...)
		}
	}

	isCustomVnet := cs.Properties.AreAgentProfilesCustomVNET()
	isAzureCNI := cs.Properties.OrchestratorProfile.IsAzureCNI()

	if isHostedMaster {
		if !isCustomVnet {
			hostedMasterVnet := createHostedMasterVirtualNetwork(cs)
			armResources = append(armResources, hostedMasterVnet)
		}

		if !isAzureCNI {
			armResources = append(armResources, createRouteTable())
		}

		hostedMasterNsg := createHostedMasterNSG()
		armResources = append(armResources, hostedMasterNsg)
	} else {
		isMasterVMSS := cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.IsVirtualMachineScaleSets()
		var masterResources []interface{}
		if isMasterVMSS {
			masterResources = createKubernetesMasterResourcesVMSS(cs)
		} else {
			masterResources = createKubernetesMasterResourcesVMAS(cs)
		}

		armResources = append(armResources, masterResources...)
	}

	if cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AppGwIngressAddonName) {
		armResources = append(armResources, createAppGwPublicIPAddress())
		armResources = append(armResources, createAppGwUserAssignedIdentities())
		armResources = append(armResources, createApplicationGateway(cs.Properties))
		armResources = append(armResources, createAppGwIdentityApplicationGatewayWriteSysRoleAssignment())
		armResources = append(armResources, createKubernetesSpAppGIdentityOperatorAccessRoleAssignment(cs.Properties))
		armResources = append(armResources, createAppGwIdentityResourceGroupReadSysRoleAssignment())
	}

	return armResources
}

func createKubernetesAgentVMASResources(cs *api.ContainerService, profile *api.AgentPoolProfile) []interface{} {
	var agentVMASResources []interface{}

	agentVMASNIC := createAgentVMASNetworkInterface(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVMASNIC)

	if profile.IsManagedDisks() {
		agentAvSet := createAgentAvailabilitySets(profile)
		agentVMASResources = append(agentVMASResources, agentAvSet)
	} else if profile.IsStorageAccount() {
		agentStorageAccount := createAgentVMASStorageAccount(cs, profile, false)
		agentVMASResources = append(agentVMASResources, agentStorageAccount)
		if profile.HasDisks() {
			agentDataDiskStorageAccount := createAgentVMASStorageAccount(cs, profile, true)
			agentVMASResources = append(agentVMASResources, agentDataDiskStorageAccount)
		}

		avSet := AvailabilitySetARM{
			ARMResource: ARMResource{
				APIVersion: "[variables('apiVersionCompute')]",
			},
			AvailabilitySet: compute.AvailabilitySet{
				Location: to.StringPtr("[variables('location')]"),
				Name: to.StringPtr(fmt.Sprintf("[variables('%sAvailabilitySet')]",
					profile.Name)),
				AvailabilitySetProperties: &compute.AvailabilitySetProperties{},
				Type:                      to.StringPtr("Microsoft.Compute/availabilitySets"),
			},
		}

		agentVMASResources = append(agentVMASResources, avSet)
	}

	agentVMASVM := createAgentAvailabilitySetVM(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVMASVM)

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	if useManagedIdentity && !userAssignedIDEnabled {
		agentVMASSysRoleAssignment := createAgentVMASSysRoleAssignment(profile)
		agentVMASResources = append(agentVMASResources, agentVMASSysRoleAssignment)
	}

	agentVMASCSE := createAgentVMASCustomScriptExtension(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVMASCSE)

	if cs.IsAKSBillingEnabled() {
		agentVMASAKSBilling := CreateAgentVMASAKSBillingExtension(cs, profile)
		agentVMASResources = append(agentVMASResources, agentVMASAKSBilling)
	}

	return agentVMASResources
}
