// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/to"
)

func createKubernetesMasterResources(cs *api.ContainerService) []interface{} {
	var masterResources []interface{}

	p := cs.Properties
	//TODO: Implement CosmosEtcd

	if p.HasManagedDisks() {
		if !p.HasAvailabilityZones() {
			masterResources = append(masterResources, CreateAvailabilitySet(cs, true))
		}
	} else if p.MasterProfile.IsStorageAccount() {
		masterResources = append(masterResources, CreateAvailabilitySet(cs, false))

		storageAccount := createStorageAccount(cs)
		masterResources = append(masterResources, storageAccount)
	}

	if !p.MasterProfile.IsCustomVNET() {
		virtualNetwork := CreateVirtualNetwork(cs)
		masterResources = append(masterResources, virtualNetwork)
	}

	masterNsg := CreateNetworkSecurityGroup(cs)
	masterResources = append(masterResources, masterNsg)

	if cs.Properties.OrchestratorProfile.RequireRouteTable() {
		masterResources = append(masterResources, createRouteTable())
	}

	isPrivateCluster := to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled)

	if !isPrivateCluster {
		publicIPAddress := CreatePublicIPAddress(cs)
		masterResources = append(masterResources, publicIPAddress)

		loadBalancer := CreateLoadBalancer()
		masterResources = append(masterResources, loadBalancer)

		inboundNatRules := createInboundNATRules()

		masterResources = append(masterResources, inboundNatRules)

		masterNic := CreateNetworkInterfaces(cs)
		masterResources = append(masterResources, masterNic)

	} else {
		masterNic := createPrivateClusterNetworkInterface(cs)
		masterResources = append(masterResources, masterNic)

		provisionJumpBox := cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision()

		if provisionJumpBox {
			jumpBoxVM := createJumpboxVirtualMachine(cs)
			masterResources = append(masterResources, jumpBoxVM)
			jumpboxIsManagedDisks :=
				cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() &&
					cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks
			if !jumpboxIsManagedDisks {
				jumpBoxStorage := createJumpboxStorageAccount()
				masterResources = append(masterResources, jumpBoxStorage)
			}
			jumpBoxNSG := createJumpboxNSG()
			masterResources = append(masterResources, jumpBoxNSG)
			jumpBoxNIC := createJumpboxNIC(cs)
			masterResources = append(masterResources, jumpBoxNIC)
		}

	}

	if p.MasterProfile.Count > 1 {
		internalLb := CreateMasterInternalLoadBalancer(cs)
		masterResources = append(masterResources, internalLb)
	}

	//TODO: Implement KMS

	masterVM := CreateVirtualMachine(cs)
	masterResources = append(masterResources, masterVM)

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	if useManagedIdentity && !userAssignedIDEnabled {
		vmasRoleAssignment := createVmasRoleAssignment()
		masterResources = append(masterResources, vmasRoleAssignment)
	}

	masterCSE := CreateCustomScriptExtension(cs)
	masterResources = append(masterResources, masterCSE)

	if cs.GetCloudSpecConfig().CloudName == api.AzurePublicCloud {
		aksBillingExtension := CreateAKSBillingExtension(cs)
		masterResources = append(masterResources, aksBillingExtension)
	}

	//TODO: Implement Extensions

	return masterResources
}

func createKubernetesMasterResourcesVmss(cs *api.ContainerService) []interface{} {
	var masterResources []interface{}

	//TODO: Implement CosmosEtcd
	//TODO: Implement KMS

	masterNsg := CreateNetworkSecurityGroup(cs)
	masterResources = append(masterResources, masterNsg)

	if cs.Properties.OrchestratorProfile.RequireRouteTable() {
		masterResources = append(masterResources, createRouteTable())
	}
	if !cs.Properties.MasterProfile.IsCustomVNET() {
		masterVnet := createVirtualNetworkVMSS(cs)
		masterResources = append(masterResources, masterVnet)
	}

	if cs.Properties.MasterProfile.Count > 1 {
		internalLb := CreateMasterInternalLoadBalancerVMSS(cs)
		masterResources = append(masterResources, internalLb)
	}

	publicIPAddress := CreatePublicIPAddress(cs)
	masterResources = append(masterResources, publicIPAddress)

	loadBalancer := CreateLoadBalancer()
	masterResources = append(masterResources, loadBalancer)

	masterVmss := CreateMasterVMSS(cs)
	masterResources = append(masterResources, masterVmss)

	return masterResources
}
