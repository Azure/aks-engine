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

	if p.MasterProfile.HasCosmosEtcd() {
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

	if kubernetesConfig.PrivateJumpboxProvision() {
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

	var masterNic NetworkInterfaceARM
	if cs.Properties.OrchestratorProfile.IsPrivateCluster() {
		masterNic = createPrivateClusterMasterVMNetworkInterface(cs)
	} else {
		masterNic = CreateMasterVMNetworkInterfaces(cs)
	}
	masterResources = append(masterResources, masterNic)

	// We don't create a master load balancer in a private cluster + single master vm scenario
	if !(cs.Properties.OrchestratorProfile.IsPrivateCluster() && !p.MasterProfile.HasMultipleNodes()) &&
		// And we don't create a master load balancer in a private cluster + Basic LB scenario
		!(cs.Properties.OrchestratorProfile.IsPrivateCluster() && cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.BasicLoadBalancerSku) {
		isForMaster := true
		var includeDNS bool
		loadBalancer := CreateMasterLoadBalancer(cs.Properties, false)
		// In a private cluster scenario, the master NIC spec is different,
		// and the master LB is for outbound access only and doesn't require a DNS record for the public IP
		if cs.Properties.OrchestratorProfile.IsPrivateCluster() {
			includeDNS = false
		} else {
			includeDNS = true
		}
		publicIPAddress := CreatePublicIPAddress(isForMaster, includeDNS)
		masterResources = append(masterResources, publicIPAddress, loadBalancer)
	}

	if p.MasterProfile.HasMultipleNodes() {
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

	if cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack") {
		// for standard lb sku, the loadbalancer and ipv4 FE is already created
		if cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != api.StandardLoadBalancerSku {
			clusterIPv4PublicIPAddress := CreateClusterPublicIPAddress()
			clusterLB := CreateClusterLoadBalancerForIPv6()

			masterResources = append(masterResources, clusterIPv4PublicIPAddress, clusterLB)
		}
	}

	masterVM := CreateMasterVM(cs)
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

	// TODO: This is only necessary if the resource group of the masters is different from the RG of the node pool
	// subnet. But when we generate the template we don't know to which RG it will be deployed to. To solve this we
	// would have to add the necessary condition into the template. For the resources we can use the `condition` field
	// but how can we conditionally declare the dependencies? Perhaps by creating a variable for the dependency array
	// and conditionally adding more dependencies.
	if kubernetesConfig.SystemAssignedIDEnabled() &&
		// The fix for ticket 2373 is only available for individual VMs / AvailabilitySet.
		cs.Properties.MasterProfile.IsAvailabilitySet() {
		masterRoleAssignmentForAgentPools := createKubernetesMasterRoleAssignmentForAgentPools(cs.Properties.MasterProfile, cs.Properties.AgentPoolProfiles)

		for _, assignmentForAgentPool := range masterRoleAssignmentForAgentPools {
			masterResources = append(masterResources, assignmentForAgentPool)
			masterCSE.ARMResource.DependsOn = append(masterCSE.ARMResource.DependsOn, *assignmentForAgentPool.Name)
		}
	}

	masterResources = append(masterResources, masterCSE)

	if cs.IsAKSBillingEnabled() {
		aksBillingExtension := CreateAKSBillingExtension(cs)
		masterResources = append(masterResources, aksBillingExtension)
	}

	customExtensions := CreateCustomExtensions(cs.Properties)
	for _, ext := range customExtensions {
		masterResources = append(masterResources, ext)
	}

	return masterResources
}

func createKubernetesMasterResourcesVMSS(cs *api.ContainerService) []interface{} {
	var masterResources []interface{}

	if cs.Properties.MasterProfile.HasCosmosEtcd() {
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

	if cs.Properties.MasterProfile.HasMultipleNodes() {
		internalLb := CreateMasterInternalLoadBalancer(cs)
		masterResources = append(masterResources, internalLb)
	}

	isForMaster := true
	includeDNS := !cs.Properties.OrchestratorProfile.IsPrivateCluster()
	publicIPAddress := CreatePublicIPAddress(isForMaster, includeDNS)
	loadBalancer := CreateMasterLoadBalancer(cs.Properties, true)
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

	if cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack") {
		// for standard lb sku, the loadbalancer and ipv4 FE is already created
		if cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != api.StandardLoadBalancerSku {
			clusterIPv4PublicIPAddress := CreateClusterPublicIPAddress()
			clusterLB := CreateClusterLoadBalancerForIPv6()

			masterResources = append(masterResources, clusterIPv4PublicIPAddress, clusterLB)
		}
	}

	masterVmss := CreateMasterVMSS(cs)
	masterResources = append(masterResources, masterVmss)

	return masterResources
}
