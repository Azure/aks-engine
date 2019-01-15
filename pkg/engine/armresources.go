// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func GenerateARMResources(cs *api.ContainerService) []interface{} {
	var armResources []interface{}

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	if userAssignedIDEnabled {
		userAssignedID := createUserAssignedIdentities()
		msiRoleAssignment := createMSIRoleAssignment()
		armResources = append(armResources, userAssignedID)
		armResources = append(armResources, msiRoleAssignment)
	}

	profiles := cs.Properties.AgentPoolProfiles

	for _, profile := range profiles {
		if profile.IsVirtualMachineScaleSets() {
			if useManagedIdentity && !userAssignedIDEnabled {
				armResources = append(armResources, createAgentVmssSysRoleAssignment(profile))
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
			masterResources = createKubernetesMasterResourcesVmss(cs)
		} else {
			masterResources = createKubernetesMasterResources(cs)
		}

		armResources = append(armResources, masterResources...)

	}

	//nsg := CreateNetworkSecurityGroup(cs)
	//virtualNetwork := CreateVirtualNetwork(cs)
	//
	//armResources = append(armResources, nsg)
	//armResources = append(armResources, virtualNetwork)

	//for _, profile := range cs.Properties.AgentPoolProfiles {
	//	vmss := CreateAgentVMSS(cs, profile)
	//	armResources = append(armResources, vmss)
	//}
	//
	//masterAVProfile := cs.Properties.MasterProfile.AvailabilityProfile
	//
	//if masterAVProfile == api.AvailabilitySet {
	//	masterAvSet := CreateAvailabilitySet(cs)
	//	armResources = append(armResources, masterAvSet)
	//	masterVM := CreateVirtualMachine(cs)
	//	armResources = append(armResources, masterVM)
	//} else {
	//	masterVmss := CreateMasterVMSS(cs)
	//	armResources = append(armResources, masterVmss)
	//}
	//
	//publicIp := CreatePublicIPAddress(cs)
	//armResources = append(armResources, publicIp)
	//loadBalancers := CreateLoadBalancer()
	//masterNIC := CreateNetworkInterfaces(cs)
	//armResources = append(armResources, loadBalancers)
	//armResources = append(armResources, masterNIC)
	//
	//if cs.GetCloudSpecConfig().CloudName == api.AzurePublicCloud {
	//	aksBillingExtension := CreateAKSBillingExtension(cs)
	//	armResources = append(armResources, aksBillingExtension)
	//}

	return armResources
}

func createKubernetesAgentVMASResources(cs *api.ContainerService, profile *api.AgentPoolProfile) []interface{} {
	var agentVMASResources []interface{}

	if profile.IsWindows() {
		if cs.Properties.WindowsProfile.HasCustomImage() {
			agentVMASResources = append(agentVMASResources, createWindowsImage(cs, profile))
		}
	}

	agentVmasNic := createAgentVMASNIC(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVmasNic)

	if profile.IsManagedDisks() {
		agentAvSet := createAgentAvailabilitySets(cs, profile)
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

	agentVmasVM := createAgentAvailabilitySetVM(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVmasVM)

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	if useManagedIdentity && !userAssignedIDEnabled {
		agentVmasSysRoleAssignment := createAgentVmasSysRoleAssignment
		agentVMASResources = append(agentVMASResources, agentVmasSysRoleAssignment)
	}

	agentVmasCSE := createAgentVMASCustomScriptExtension(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVmasCSE)

	agentVmasAksBilling := CreateAgentVMASAKSBillingExtension(cs, profile)
	agentVMASResources = append(agentVMASResources, agentVmasAksBilling)

	return agentVMASResources
}
