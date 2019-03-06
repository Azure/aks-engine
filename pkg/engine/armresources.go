// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func GenerateARMResources(cs *api.ContainerService) []interface{} {
	var armResources []interface{}

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	if userAssignedIDEnabled {
		userAssignedID := createUserAssignedIdentities()
		msiRoleAssignment := createMSIRoleAssignment()
		armResources = append(armResources, userAssignedID, msiRoleAssignment)
	}

	profiles := cs.Properties.AgentPoolProfiles

	for _, profile := range profiles {
		if profile.IsVirtualMachineScaleSets() {
			if useManagedIdentity && !userAssignedIDEnabled {
				armResources = append(armResources, createAgentVMSSSysRoleAssignment(profile))
			}
			armResources = append(armResources, CreateAgentVMSS(cs, profile))
		} else {
			agentVmasResources := createKubernetesAgentVMASResources(cs, profile)
			armResources = append(armResources, agentVmasResources...)
		}
	}

	isHostedMaster := cs.Properties.IsHostedMasterProfile()
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

	return armResources
}

func createKubernetesAgentVMASResources(cs *api.ContainerService, profile *api.AgentPoolProfile) []interface{} {
	var agentVMASResources []interface{}

	if profile.IsWindows() {
		if cs.Properties.WindowsProfile.HasCustomImage() {
			agentVMASResources = append(agentVMASResources, createWindowsImage(profile))
		}
	}

	agentVmasNic := createAgentVMASNetworkInterface(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVmasNic)

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
		agentVMASSysRoleAssignment := createAgentVMASSysRoleAssignment
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
