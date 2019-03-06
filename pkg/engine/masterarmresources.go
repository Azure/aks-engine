// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/to"
)

func createKubernetesMasterResourcesVMAS(cs *api.ContainerService) []interface{} {
	var masterResources []interface{}

	p := cs.Properties
	//TODO: Implement CosmosEtcd

	if p.HasManagedDisks() {
		if !p.HasAvailabilityZones() {
			masterResources = append(masterResources, CreateAvailabilitySet(cs, true))
		}
	} else if p.MasterProfile.IsStorageAccount() {
		availabilitySet := CreateAvailabilitySet(cs, false)
		storageAccount := createStorageAccount(cs)
		masterResources = append(masterResources, availabilitySet, storageAccount)
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
		publicIPAddress := CreatePublicIPAddress()
		loadBalancer := CreateLoadBalancer(cs.Properties.MasterProfile.Count)
		masterNic := CreateNetworkInterfaces(cs)

		masterResources = append(masterResources, publicIPAddress, loadBalancer, masterNic)

	} else {
		masterNic := createPrivateClusterNetworkInterface(cs)
		masterResources = append(masterResources, masterNic)

		provisionJumpbox := cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision()

		if provisionJumpbox {
			jumpboxVM := createJumpboxVirtualMachine(cs)
			masterResources = append(masterResources, jumpboxVM)
			jumpboxIsManagedDisks :=
				cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() &&
					cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks
			if !jumpboxIsManagedDisks {
				jumpBoxStorage := createJumpboxStorageAccount()
				masterResources = append(masterResources, jumpBoxStorage)
			}
			jumpboxNSG := createJumpboxNSG()
			jumpboxNIC := createJumpboxNetworkInterface(cs)
			jumpboxPublicIP := createJumpboxPublicIPAddress()
			masterResources = append(masterResources, jumpboxNSG, jumpboxNIC, jumpboxPublicIP)
		}

	}

	if p.MasterProfile.Count > 1 {
		internalLB := CreateMasterInternalLoadBalancer(cs)
		masterResources = append(masterResources, internalLB)
	}

	isKMSEnabled := to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms)

	if isKMSEnabled {
		keyVaultStorageAccount := createKeyVaultStorageAccount()
		keyVault := CreateKeyVaultVMAS(cs)
		masterResources = append(masterResources, keyVaultStorageAccount, keyVault)
	}

	masterVM := CreateVirtualMachine(cs)
	masterResources = append(masterResources, masterVM)

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := useManagedIdentity && cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	if useManagedIdentity && !userAssignedIDEnabled {
		vmasRoleAssignment := createVMASRoleAssignment()
		masterResources = append(masterResources, vmasRoleAssignment)
	}

	masterCSE := CreateCustomScriptExtension(cs)
	masterResources = append(masterResources, masterCSE)

	if cs.IsAKSBillingEnabled() {
		aksBillingExtension := CreateAKSBillingExtension(cs)
		masterResources = append(masterResources, aksBillingExtension)
	}

	//TODO: Implement Extensions

	return masterResources
}

func createKubernetesMasterResourcesVMSS(cs *api.ContainerService) []interface{} {
	var masterResources []interface{}

	//TODO: Implement CosmosEtcd

	masterNSG := CreateNetworkSecurityGroup(cs)
	masterResources = append(masterResources, masterNSG)

	if cs.Properties.OrchestratorProfile.RequireRouteTable() {
		masterResources = append(masterResources, createRouteTable())
	}
	if !cs.Properties.MasterProfile.IsCustomVNET() {
		masterVNET := createVirtualNetworkVMSS(cs)
		masterResources = append(masterResources, masterVNET)
	}

	if cs.Properties.MasterProfile.Count > 1 {
		internalLb := CreateMasterInternalLoadBalancer(cs)
		masterResources = append(masterResources, internalLb)
	}

	publicIPAddress := CreatePublicIPAddress()
	loadBalancer := CreateLoadBalancer(cs.Properties.MasterProfile.Count)
	masterResources = append(masterResources, publicIPAddress, loadBalancer)

	isKMSEnabled := to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms)

	if isKMSEnabled {
		keyVaultStorageAccount := createKeyVaultStorageAccount()
		keyVault := CreateKeyVaultVMSS(cs)
		masterResources = append(masterResources, keyVaultStorageAccount, keyVault)
	}

	masterVmss := CreateMasterVMSS(cs)
	masterResources = append(masterResources, masterVmss)

	return masterResources
}
