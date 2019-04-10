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

	if to.Bool(p.MasterProfile.CosmosEtcd) {
		masterResources = append(masterResources, createCosmosDBAccount())
	}

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

	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig

	if !cs.Properties.OrchestratorProfile.IsPrivateCluster() {
		publicIPAddress := CreatePublicIPAddress()
		loadBalancer := CreateLoadBalancer(cs.Properties.MasterProfile.Count, false)
		masterNic := CreateNetworkInterfaces(cs)

		masterResources = append(masterResources, publicIPAddress, loadBalancer, masterNic)

	} else {
		masterNic := createPrivateClusterNetworkInterface(cs)
		masterResources = append(masterResources, masterNic)

		var provisionJumpbox bool

		if kubernetesConfig != nil {
			provisionJumpbox = kubernetesConfig.PrivateJumpboxProvision()
		}

		if provisionJumpbox {
			jumpboxVM := createJumpboxVirtualMachine(cs)
			masterResources = append(masterResources, jumpboxVM)
			jumpboxIsManagedDisks :=
				kubernetesConfig.PrivateJumpboxProvision() &&
					kubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks
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

	var isKMSEnabled bool
	if kubernetesConfig != nil {
		isKMSEnabled = to.Bool(kubernetesConfig.EnableEncryptionWithExternalKms)
	}

	if isKMSEnabled {
		keyVaultStorageAccount := createKeyVaultStorageAccount()
		keyVault := CreateKeyVaultVMAS(cs)
		masterResources = append(masterResources, keyVaultStorageAccount, keyVault)
	}

	masterVM := CreateVirtualMachine(cs)
	masterResources = append(masterResources, masterVM)

	var useManagedIdentity, userAssignedIDEnabled bool
	useManagedIdentity = kubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled = useManagedIdentity && kubernetesConfig.UserAssignedID != ""

	if useManagedIdentity && !userAssignedIDEnabled {
		vmasRoleAssignment := createVMASRoleAssignment()
		masterResources = append(masterResources, vmasRoleAssignment)
	}

	masterCSE := CreateCustomScriptExtension(cs)
	if isKMSEnabled {
		masterCSE.ARMResource.DependsOn = append(masterCSE.ARMResource.DependsOn, "[concat('Microsoft.KeyVault/vaults/', variables('clusterKeyVaultName'))]")
	}
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

	if to.Bool(cs.Properties.MasterProfile.CosmosEtcd) {
		masterResources = append(masterResources, createCosmosDBAccount())
	}

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
	loadBalancer := CreateLoadBalancer(cs.Properties.MasterProfile.Count, true)
	masterResources = append(masterResources, publicIPAddress, loadBalancer)

	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig

	var isKMSEnabled bool
	if kubernetesConfig != nil {
		isKMSEnabled = to.Bool(kubernetesConfig.EnableEncryptionWithExternalKms)
	}

	if isKMSEnabled {
		keyVaultStorageAccount := createKeyVaultStorageAccount()
		keyVault := CreateKeyVaultVMSS(cs)
		masterResources = append(masterResources, keyVaultStorageAccount, keyVault)
	}

	masterVmss := CreateMasterVMSS(cs)
	masterResources = append(masterResources, masterVmss)

	return masterResources
}
