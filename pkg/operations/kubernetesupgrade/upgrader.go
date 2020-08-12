// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/armhelpers/utils"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/operations"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Upgrader holds information on upgrading an AKS cluster
type Upgrader struct {
	Translator *i18n.Translator
	logger     *logrus.Entry
	ClusterTopology
	Client             armhelpers.AKSEngineClient
	kubeConfig         string
	stepTimeout        *time.Duration
	cordonDrainTimeout *time.Duration
	AKSEngineVersion   string
	CurrentVersion     string
	ControlPlaneOnly   bool
}

type vmStatus int

const (
	defaultTimeout                     = time.Minute * 20
	defaultCordonDrainTimeout          = time.Minute * 20
	nodePropertiesCopyTimeout          = time.Minute * 5
	getResourceTimeout                 = time.Minute * 1
	clusterUpgradeTimeout              = time.Minute * 180
	vmStatusUpgraded          vmStatus = iota
	vmStatusNotUpgraded
	vmStatusIgnored
)

type vmInfo struct {
	name   string
	status vmStatus
}

// Init initializes an upgrader struct
func (ku *Upgrader) Init(translator *i18n.Translator, logger *logrus.Entry, clusterTopology ClusterTopology, client armhelpers.AKSEngineClient, kubeConfig string, stepTimeout *time.Duration, cordonDrainTimeout *time.Duration, aksEngineVersion string, controlPlaneOnly bool) {
	ku.Translator = translator
	ku.logger = logger
	ku.ClusterTopology = clusterTopology
	ku.Client = client
	ku.kubeConfig = kubeConfig
	ku.stepTimeout = stepTimeout
	ku.cordonDrainTimeout = cordonDrainTimeout
	ku.AKSEngineVersion = aksEngineVersion
	ku.ControlPlaneOnly = controlPlaneOnly
}

// RunUpgrade runs the upgrade pipeline
func (ku *Upgrader) RunUpgrade() error {
	ctx, cancel := context.WithTimeout(context.Background(), clusterUpgradeTimeout)
	defer cancel()
	if err := ku.upgradeMasterNodes(ctx); err != nil {
		return err
	}

	ku.handleUnreconcilableAddons()

	if ku.ControlPlaneOnly {
		return nil
	}

	if err := ku.upgradeAgentScaleSets(ctx); err != nil {
		return err
	}

	//This is handling VMAS VMs only, not VMSS
	return ku.upgradeAgentPools(ctx)
}

// handleUnreconcilableAddons ensures addon upgrades that addon-manager cannot handle by itself.
// This method fails silently otherwide it would break test "Should not fail if a Kubernetes client cannot be created" (upgradecluster_test.go)
func (ku *Upgrader) handleUnreconcilableAddons() {
	upgradeVersion := ku.DataModel.Properties.OrchestratorProfile.OrchestratorVersion
	// kube-proxy upgrade fails from v1.15 to 1.16: https://github.com/Azure/aks-engine/issues/3557
	// deleting daemonset so addon-manager recreates instead of patching
	if !common.IsKubernetesVersionGe(ku.CurrentVersion, "1.16.0") && common.IsKubernetesVersionGe(upgradeVersion, "1.16.0") {
		ku.logger.Infof("Attempting to delete kube-proxy daemonset.")
		client, err := ku.getKubernetesClient(getResourceTimeout)
		if err != nil {
			ku.logger.Errorf("Error getting Kubernetes client: %v", err)
			return
		}
		err = client.DeleteDaemonSet(&appsv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "kube-system",
				Name:      common.KubeProxyAddonName,
			},
		})
		if err != nil {
			ku.logger.Errorf("Error deleting kube-proxy daemonset: %v", err)
		}
		ku.logger.Infof("Deleted kube-proxy daemonset. Addon-manager will recreate it.")
	}
	// metrics-server upgrade fails from v1.15 to 1.16 as the addon mode is EnsureExists for pre-v1.16 cluster
	if !common.IsKubernetesVersionGe(ku.CurrentVersion, "1.16.0") && common.IsKubernetesVersionGe(upgradeVersion, "1.16.0") {
		ku.logger.Infof("Attempting to delete metrics-server deployment.")
		client, err := ku.getKubernetesClient(getResourceTimeout)
		if err != nil {
			ku.logger.Errorf("Error getting Kubernetes client: %v", err)
			return
		}
		err = client.DeleteClusterRole(&rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name: "system:metrics-server",
			},
		})
		if err != nil {
			ku.logger.Errorf("Error deleting metrics-server cluster role: %v", err)
		}
		err = client.DeleteDeployment(&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "kube-system",
				Name:      common.MetricsServerAddonName,
			},
		})
		if err != nil {
			ku.logger.Errorf("Error deleting metrics-server deployment: %v", err)
		}
		ku.logger.Infof("Deleted metrics-server deployment. Addon-manager will recreate it.")
	}
}

// Validate will run validation post upgrade
func (ku *Upgrader) Validate() error {
	return nil
}

func (ku *Upgrader) upgradeMasterNodes(ctx context.Context) error {
	if ku.ClusterTopology.DataModel.Properties.MasterProfile == nil {
		return nil
	}
	ku.logger.Infof("Master nodes StorageProfile: %s", ku.ClusterTopology.DataModel.Properties.MasterProfile.StorageProfile)
	// Upgrade Master VMs
	templateMap, parametersMap, err := ku.generateUpgradeTemplate(ku.ClusterTopology.DataModel, ku.AKSEngineVersion)
	if err != nil {
		return ku.Translator.Errorf("error generating upgrade template: %s", err.Error())
	}

	ku.logger.Infof("Prepping master nodes for upgrade...")

	transformer := &transform.Transformer{
		Translator: ku.Translator,
	}

	if ku.ClusterTopology.DataModel.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() {
		err = transformer.RemoveJumpboxResourcesFromTemplate(ku.logger, templateMap)
		if err != nil {
			return ku.Translator.Errorf("error removing jumpbox resources from template: %s", err.Error())
		}
	}

	if ku.DataModel.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku {
		err = transformer.NormalizeForK8sSLBScalingOrUpgrade(ku.logger, templateMap)
		if err != nil {
			return ku.Translator.Errorf("error normalizing upgrade template for SLB: %s", err.Error())
		}
	}
	//TODO: rename this as it's not only touching master resources
	if err = transformer.NormalizeResourcesForK8sMasterUpgrade(ku.logger, templateMap, ku.DataModel.Properties.MasterProfile.IsManagedDisks(), nil); err != nil {
		ku.logger.Error(err.Error())
		return err
	}

	transformer.RemoveImmutableResourceProperties(ku.logger, templateMap)

	upgradeMasterNode := UpgradeMasterNode{
		Translator: ku.Translator,
		logger:     ku.logger,
	}
	upgradeMasterNode.TemplateMap = templateMap
	upgradeMasterNode.ParametersMap = parametersMap
	upgradeMasterNode.UpgradeContainerService = ku.ClusterTopology.DataModel
	upgradeMasterNode.ResourceGroup = ku.ClusterTopology.ResourceGroup
	upgradeMasterNode.SubscriptionID = ku.ClusterTopology.SubscriptionID
	upgradeMasterNode.Client = ku.Client
	upgradeMasterNode.kubeConfig = ku.kubeConfig
	if ku.stepTimeout == nil {
		upgradeMasterNode.timeout = defaultTimeout
	} else {
		upgradeMasterNode.timeout = *ku.stepTimeout
	}

	expectedMasterCount := ku.ClusterTopology.DataModel.Properties.MasterProfile.Count
	mastersUpgradedCount := len(*ku.ClusterTopology.UpgradedMasterVMs)
	mastersToUgradeCount := expectedMasterCount - mastersUpgradedCount

	ku.logger.Infof("Total expected master count: %d", expectedMasterCount)
	ku.logger.Infof("Master nodes that need to be upgraded: %d", mastersToUgradeCount)
	ku.logger.Infof("Master nodes that have been upgraded: %d", mastersUpgradedCount)

	ku.logger.Infof("Starting upgrade of master nodes...")

	masterNodesInCluster := len(*ku.ClusterTopology.MasterVMs) + mastersUpgradedCount
	ku.logger.Infof("masterNodesInCluster: %d", masterNodesInCluster)
	if masterNodesInCluster > expectedMasterCount {
		return ku.Translator.Errorf("Total count of master VMs: %d exceeded expected count: %d", masterNodesInCluster, expectedMasterCount)
	}

	// This condition is possible if the previous upgrade operation failed during master
	// VM upgrade when a master VM was deleted but creation of upgraded master did not run.
	if masterNodesInCluster < expectedMasterCount {
		ku.logger.Infof(
			"Found missing master VMs in the cluster. Reconstructing names of missing master VMs for recreation during upgrade...")
	}

	upgradedMastersIndex := make(map[int]bool)
	mastersToCreate := expectedMasterCount - masterNodesInCluster
	ku.logger.Infof("Expected master count: %d, Creating %d more master VMs", expectedMasterCount, mastersToCreate)

	// NOTE: this is NOT completely idempotent because it assumes that
	// the OS disk has been deleted
	for i := 0; i < mastersToCreate; i++ {
		masterIndexToCreate := 0
		for upgradedMastersIndex[masterIndexToCreate] {
			masterIndexToCreate++
		}

		ku.logger.Infof("Creating upgraded master VM with index: %d", masterIndexToCreate)

		err = upgradeMasterNode.CreateNode(ctx, "master", masterIndexToCreate)
		if err != nil {
			ku.logger.Infof("Error creating upgraded master VM with index: %d", masterIndexToCreate)
			return err
		}

		tempVMName := ""
		err = upgradeMasterNode.Validate(&tempVMName)
		if err != nil {
			ku.logger.Infof("Error validating upgraded master VM with index: %d", masterIndexToCreate)
			return err
		}

		upgradedMastersIndex[masterIndexToCreate] = true
	}

	for _, vm := range *ku.ClusterTopology.UpgradedMasterVMs {
		ku.logger.Infof("Master VM: %s is upgraded to expected orchestrator version", *vm.Name)
		masterIndex, _ := utils.GetVMNameIndex(vm.StorageProfile.OsDisk.OsType, *vm.Name)
		upgradedMastersIndex[masterIndex] = true
	}

	for _, vm := range *ku.ClusterTopology.MasterVMs {
		ku.logger.Infof("Upgrading Master VM: %s", *vm.Name)

		masterIndex, _ := utils.GetVMNameIndex(vm.StorageProfile.OsDisk.OsType, *vm.Name)

		err = upgradeMasterNode.DeleteNode(vm.Name, false)
		if err != nil {
			ku.logger.Infof("Error deleting master VM: %s, err: %v", *vm.Name, err)
			return err
		}

		err = upgradeMasterNode.CreateNode(ctx, "master", masterIndex)
		if err != nil {
			ku.logger.Infof("Error creating upgraded master VM: %s", *vm.Name)
			return err
		}

		err = upgradeMasterNode.Validate(vm.Name)
		if err != nil {
			ku.logger.Infof("Error validating upgraded master VM: %s", *vm.Name)
			return err
		}

		upgradedMastersIndex[masterIndex] = true
	}

	return nil
}

func (ku *Upgrader) upgradeAgentPools(ctx context.Context) error {
	for _, agentPool := range ku.ClusterTopology.AgentPools {
		// Upgrade Agent VMs
		templateMap, parametersMap, err := ku.generateUpgradeTemplate(ku.ClusterTopology.DataModel, ku.AKSEngineVersion)
		if err != nil {
			ku.logger.Errorf("Error generating upgrade template: %v", err)
			return ku.Translator.Errorf("Error generating upgrade template: %s", err.Error())
		}

		ku.logger.Infof("Prepping agent pool '%s' for upgrade...", *agentPool.Name)

		preservePools := map[string]bool{*agentPool.Name: true}
		transformer := &transform.Transformer{
			Translator: ku.Translator,
		}

		if ku.ClusterTopology.DataModel.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() {
			err = transformer.RemoveJumpboxResourcesFromTemplate(ku.logger, templateMap)
			if err != nil {
				return ku.Translator.Errorf("error removing jumpbox resources from template: %s", err.Error())
			}
		}

		var isMasterManagedDisk bool
		if ku.DataModel.Properties.MasterProfile != nil {
			isMasterManagedDisk = ku.DataModel.Properties.MasterProfile.IsManagedDisks()
		}

		if ku.DataModel.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku {
			err = transformer.NormalizeForK8sSLBScalingOrUpgrade(ku.logger, templateMap)
			if err != nil {
				return ku.Translator.Errorf("error normalizing upgrade template for SLB: %s", err.Error())
			}
		}
		if err = transformer.NormalizeResourcesForK8sAgentUpgrade(ku.logger, templateMap, isMasterManagedDisk, preservePools); err != nil {
			ku.logger.Errorf(err.Error())
			return ku.Translator.Errorf("Error generating upgrade template: %s", err.Error())
		}

		transformer.RemoveImmutableResourceProperties(ku.logger, templateMap)

		var agentCount int
		var agentPoolProfile *api.AgentPoolProfile
		for _, app := range ku.ClusterTopology.DataModel.Properties.AgentPoolProfiles {
			if app.Name == *agentPool.Name {
				agentCount = app.Count
				agentPoolProfile = app
				break
			}
		}

		if agentCount == 0 {
			ku.logger.Infof("Agent pool '%s' is empty", *agentPool.Name)
			return nil
		}

		upgradeAgentNode := UpgradeAgentNode{
			Translator: ku.Translator,
			logger:     ku.logger,
		}
		upgradeAgentNode.TemplateMap = templateMap
		upgradeAgentNode.ParametersMap = parametersMap
		upgradeAgentNode.UpgradeContainerService = ku.ClusterTopology.DataModel
		upgradeAgentNode.SubscriptionID = ku.ClusterTopology.SubscriptionID
		upgradeAgentNode.ResourceGroup = ku.ClusterTopology.ResourceGroup
		upgradeAgentNode.Client = ku.Client
		upgradeAgentNode.kubeConfig = ku.kubeConfig
		if ku.stepTimeout == nil {
			upgradeAgentNode.timeout = defaultTimeout
		} else {
			upgradeAgentNode.timeout = *ku.stepTimeout
		}
		if ku.cordonDrainTimeout == nil {
			upgradeAgentNode.cordonDrainTimeout = defaultCordonDrainTimeout
		} else {
			upgradeAgentNode.cordonDrainTimeout = *ku.cordonDrainTimeout
		}

		agentVMs := make(map[int]*vmInfo)
		// Go over upgraded VMs and verify provisioning state
		// per https://docs.microsoft.com/en-us/rest/api/compute/virtualmachines/virtualmachines-state :
		//  - Creating: Indicates the virtual Machine is being created.
		//  - Updating: Indicates that there is an update operation in progress on the Virtual Machine.
		//  - Succeeded: Indicates that the operation executed on the virtual machine succeeded.
		//  - Deleting: Indicates that the virtual machine is being deleted.
		//  - Failed: Indicates that the update operation on the Virtual Machine failed.
		// Delete VMs in 'bad' state. Such VMs will be re-created later in this function.
		upgradedCount := 0
		for _, vm := range *agentPool.UpgradedAgentVMs {
			ku.logger.Infof("Agent VM: %s, pool name: %s on expected orchestrator version", *vm.Name, *agentPool.Name)
			var vmProvisioningState string
			if vm.VirtualMachineProperties != nil && vm.VirtualMachineProperties.ProvisioningState != nil {
				vmProvisioningState = *vm.VirtualMachineProperties.ProvisioningState
			}
			agentIndex, _ := utils.GetVMNameIndex(vm.StorageProfile.OsDisk.OsType, *vm.Name)

			switch vmProvisioningState {
			case "Creating", "Updating", "Succeeded":
				agentVMs[agentIndex] = &vmInfo{*vm.Name, vmStatusUpgraded}
				upgradedCount++

			case "Failed":
				ku.logger.Infof("Deleting agent VM %s in provisioning state %s", *vm.Name, vmProvisioningState)
				err = upgradeAgentNode.DeleteNode(vm.Name, false)
				if err != nil {
					ku.logger.Errorf("Error deleting agent VM %s: %v", *vm.Name, err)
					return err
				}

			case "Deleting":
				fallthrough
			default:
				ku.logger.Infof("Ignoring agent VM %s in provisioning state %s", *vm.Name, vmProvisioningState)
				agentVMs[agentIndex] = &vmInfo{*vm.Name, vmStatusIgnored}
			}
		}

		for _, vm := range *agentPool.AgentVMs {
			agentIndex, _ := utils.GetVMNameIndex(vm.StorageProfile.OsDisk.OsType, *vm.Name)
			agentVMs[agentIndex] = &vmInfo{*vm.Name, vmStatusNotUpgraded}
		}
		toBeUpgradedCount := len(*agentPool.AgentVMs)

		ku.logger.Infof("Starting upgrade of %d agent nodes (out of %d) in pool identifier: %s, name: %s...",
			toBeUpgradedCount, agentCount, *agentPool.Identifier, *agentPool.Name)

		// Create missing nodes to match agentCount. This could be due to previous upgrade failure
		// If there are nodes that need to be upgraded, create one extra node, which will be used to take on the load from upgrading nodes.
		if toBeUpgradedCount > 0 {
			agentCount++
		}

		newCreatedVMs := []string{}
		client, err := ku.getKubernetesClient(10 * time.Second)
		if err != nil {
			ku.logger.Errorf("Error getting Kubernetes client: %v", err)
			return err
		}

		for upgradedCount+toBeUpgradedCount < agentCount {
			agentIndex := getAvailableIndex(agentVMs)

			var vmName string
			vmName, err = utils.GetK8sVMName(ku.DataModel.Properties, agentPoolProfile, agentIndex)
			if err != nil {
				ku.logger.Errorf("Error reconstructing agent VM name with index %d: %v", agentIndex, err)
				return err
			}
			ku.logger.Infof("Creating new agent node %s (index %d)", vmName, agentIndex)

			err = upgradeAgentNode.CreateNode(ctx, *agentPool.Name, agentIndex)
			if err != nil {
				ku.logger.Errorf("Error creating agent node %s (index %d): %v", vmName, agentIndex, err)
				return err
			}

			err = upgradeAgentNode.Validate(&vmName)
			if err != nil {
				ku.logger.Infof("Error validating agent node %s (index %d): %v", vmName, agentIndex, err)
				return err
			}

			newCreatedVMs = append(newCreatedVMs, vmName)
			agentVMs[agentIndex] = &vmInfo{vmName, vmStatusUpgraded}
			upgradedCount++
		}

		if toBeUpgradedCount == 0 {
			ku.logger.Infof("No nodes to upgrade")
			continue
		}

		// Upgrade nodes in agent pool
		upgradedCount = 0
		for agentIndex, vm := range agentVMs {
			if vm.status != vmStatusNotUpgraded {
				continue
			}
			ku.logger.Infof("Upgrading Agent VM: %s, pool name: %s", vm.name, *agentPool.Name)

			// copy custom properties from old node to new node if the PreserveNodesProperties in AgentPoolProfile is not set to false explicitly.
			preserveNodesProperties := api.DefaultPreserveNodesProperties
			if agentPoolProfile != nil && agentPoolProfile.PreserveNodesProperties != nil {
				preserveNodesProperties = *agentPoolProfile.PreserveNodesProperties
			}

			if preserveNodesProperties {
				if len(newCreatedVMs) > 0 {
					newNodeName := newCreatedVMs[0]
					newCreatedVMs = newCreatedVMs[1:]
					ku.logger.Infof("Copying custom annotations, labels, taints from old node %s to new node %s...", vm.name, newNodeName)
					err = ku.copyCustomPropertiesToNewNode(client, strings.ToLower(vm.name), newNodeName)
					if err != nil {
						ku.logger.Warningf("Failed to copy custom annotations, labels, taints from old node %s to new node %s: %v", vm.name, newNodeName, err)
					}
				}
			}

			err := upgradeAgentNode.DeleteNode(&vm.name, true)
			if err != nil {
				ku.logger.Errorf("Error deleting agent VM %s: %v", vm.name, err)
				return err
			}

			vmName, err := utils.GetK8sVMName(ku.DataModel.Properties, agentPoolProfile, agentIndex)
			if err != nil {
				ku.logger.Errorf("Error fetching new VM name: %v", err)
				return err
			}

			// do not create last node in favor of already created extra node.
			if upgradedCount == toBeUpgradedCount-1 {
				ku.logger.Infof("Skipping creation of VM %s (index %d)", vmName, agentIndex)
				delete(agentVMs, agentIndex)
			} else {
				err = upgradeAgentNode.CreateNode(ctx, *agentPool.Name, agentIndex)
				if err != nil {
					ku.logger.Errorf("Error creating upgraded agent VM %s: %v", vmName, err)
					return err
				}

				err = upgradeAgentNode.Validate(&vmName)
				if err != nil {
					ku.logger.Errorf("Error validating upgraded agent VM %s: %v", vmName, err)
					return err
				}
				newCreatedVMs = append(newCreatedVMs, vmName)
				vm.status = vmStatusUpgraded
			}
			upgradedCount++
		}
	}

	return nil
}

func (ku *Upgrader) upgradeAgentScaleSets(ctx context.Context) error {
	agentPoolMap := make(map[string]*api.AgentPoolProfile)
	for _, app := range ku.ClusterTopology.DataModel.Properties.AgentPoolProfiles {
		agentPoolMap[app.Name] = app
	}

	if len(ku.ClusterTopology.AgentPoolScaleSetsToUpgrade) > 0 {
		// need to apply the ARM template with target Kubernetes version to the VMSS first in order that the new VMSS instances
		// created can get the expected Kubernetes version. Otherwise the new instances created still have old Kubernetes version
		// if the topology doesn't have master nodes (so there are no ARM deployments in previous upgradeMasterNodes step)
		templateMap, parametersMap, err := ku.generateUpgradeTemplate(ku.ClusterTopology.DataModel, ku.AKSEngineVersion)
		if err != nil {
			ku.logger.Errorf("error generating upgrade template in upgradeAgentScaleSets: %v", err)
			return err
		}

		transformer := &transform.Transformer{
			Translator: ku.Translator,
		}

		if ku.ClusterTopology.DataModel.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() {
			err = transformer.RemoveJumpboxResourcesFromTemplate(ku.logger, templateMap)
			if err != nil {
				return ku.Translator.Errorf("error removing jumpbox resources from template: %s", err.Error())
			}
		}

		// TODO: rename this!
		// This is not called in scaling scenarios. only in this upgrade scenario!
		if err = transformer.NormalizeMasterResourcesForScaling(ku.logger, templateMap); err != nil {
			return err
		}

		if ku.DataModel.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku {
			ku.logger.Infof("upgradeAgentScaleSets SLB...")
			err = transformer.NormalizeForK8sSLBScalingOrUpgrade(ku.logger, templateMap)
			if err != nil {
				return ku.Translator.Errorf("error normalizing upgrade template for SLB: %s", err.Error())
			}
		}

		transformer.RemoveImmutableResourceProperties(ku.logger, templateMap)

		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		deploymentSuffix := random.Int31()
		deploymentName := fmt.Sprintf("agentscaleset-%s-%d", time.Now().Format("06-01-02T15.04.05"), deploymentSuffix)

		ku.logger.Infof("Deploying the agent scale sets ARM template...")
		_, err = ku.Client.DeployTemplate(
			ctx,
			ku.ClusterTopology.ResourceGroup,
			deploymentName,
			templateMap,
			parametersMap)

		if err != nil {
			ku.logger.Errorf("error applying upgrade template in upgradeAgentScaleSets: %v", err)
			return err
		}
	}

	for _, vmssToUpgrade := range ku.ClusterTopology.AgentPoolScaleSetsToUpgrade {
		ku.logger.Infof("Upgrading VMSS %s", vmssToUpgrade.Name)

		if len(vmssToUpgrade.VMsToUpgrade) == 0 {
			ku.logger.Infof("No VMs to upgrade for VMSS %s, skipping", vmssToUpgrade.Name)
			continue
		}

		newCapacity := *vmssToUpgrade.Sku.Capacity + 1
		ku.logger.Infof(
			"VMSS %s current capacity is %d and new capacity will be %d while each node is swapped",
			vmssToUpgrade.Name,
			*vmssToUpgrade.Sku.Capacity,
			newCapacity,
		)

		*vmssToUpgrade.Sku.Capacity = newCapacity

		for _, vmToUpgrade := range vmssToUpgrade.VMsToUpgrade {
			if err := ku.Client.SetVirtualMachineScaleSetCapacity(
				ctx,
				ku.ClusterTopology.ResourceGroup,
				vmssToUpgrade.Name,
				vmssToUpgrade.Sku,
				vmssToUpgrade.Location,
			); err != nil {
				ku.logger.Errorf("Failure to set capacity for VMSS %s", vmssToUpgrade.Name)
				return err
			}

			ku.logger.Infof("Successfully set capacity for VMSS %s", vmssToUpgrade.Name)

			var cordonDrainTimeout time.Duration
			if ku.cordonDrainTimeout == nil {
				cordonDrainTimeout = defaultCordonDrainTimeout
			} else {
				cordonDrainTimeout = *ku.cordonDrainTimeout
			}

			// Before we can delete the node we should safely and responsibly drain it
			client, err := ku.getKubernetesClient(cordonDrainTimeout)
			if err != nil {
				ku.logger.Errorf("Error getting Kubernetes client: %v", err)
				return err
			}

			ku.logger.Infof("Draining node %s", vmToUpgrade.Name)
			err = operations.SafelyDrainNodeWithClient(
				client,
				ku.logger,
				vmToUpgrade.Name,
				cordonDrainTimeout,
			)
			if err != nil {
				ku.logger.Errorf("Error draining VM in VMSS: %v", err)
				// Continue even if there's an error in draining the node.
			}

			ku.logger.Infof(
				"Deleting VM %s in VMSS %s",
				vmToUpgrade.Name,
				vmssToUpgrade.Name,
			)

			// copy custom properties from old node to new node if the PreserveNodesProperties in AgentPoolProfile is not set to false explicitly.
			preserveNodesProperties := api.DefaultPreserveNodesProperties
			var poolName string
			if vmssToUpgrade.IsWindows {
				poolName, _ = utils.WindowsVmssNameParts(vmssToUpgrade.Name)
			} else {
				poolName, _, _ = utils.VmssNameParts(vmssToUpgrade.Name)
			}
			if agentPool, ok := agentPoolMap[poolName]; ok {
				if agentPool != nil && agentPool.PreserveNodesProperties != nil {
					preserveNodesProperties = *agentPool.PreserveNodesProperties
				}
			}

			if preserveNodesProperties {
				newNodeName, err := ku.getLastVMNameInVMSS(ctx, ku.ClusterTopology.ResourceGroup, vmssToUpgrade.Name)
				if err != nil {
					return err
				}

				ku.logger.Infof("Copying custom annotations, labels, taints from old node %s to new node %s...", vmToUpgrade.Name, newNodeName)
				err = ku.copyCustomPropertiesToNewNode(client, strings.ToLower(vmToUpgrade.Name), strings.ToLower(newNodeName))
				if err != nil {
					ku.logger.Warningf("Failed to copy custom annotations, labels, taints from old node %s to new node %s: %v", vmToUpgrade.Name, newNodeName, err)
				}
			}

			// At this point we have our buffer node that will replace the node to delete
			// so we can just remove this current node then
			if err := ku.Client.DeleteVirtualMachineScaleSetVM(
				ctx,
				ku.ClusterTopology.ResourceGroup,
				vmssToUpgrade.Name,
				vmToUpgrade.InstanceID,
			); err != nil {
				ku.logger.Errorf(
					"Failed to delete VM %s in VMSS %s",
					vmToUpgrade.Name,
					vmssToUpgrade.Name)
				return err
			}
			ku.logger.Infof(
				"Successfully deleted VM %s in VMSS %s",
				vmToUpgrade.Name,
				vmssToUpgrade.Name)
		}
		ku.logger.Infof("Completed upgrading VMSS %s", vmssToUpgrade.Name)
	}

	ku.logger.Infoln("Completed upgrading all VMSS")

	return nil
}

func (ku *Upgrader) generateUpgradeTemplate(upgradeContainerService *api.ContainerService, aksEngineVersion string) (map[string]interface{}, map[string]interface{}, error) {
	var err error
	ctx := engine.Context{
		Translator: ku.Translator,
	}
	templateGenerator, err := engine.InitializeTemplateGenerator(ctx)
	if err != nil {
		return nil, nil, ku.Translator.Errorf("failed to initialize template generator: %s", err.Error())
	}

	_, err = upgradeContainerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  true,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		return nil, nil, ku.Translator.Errorf("error in SetPropertiesDefaults: %s", err.Error())

	}

	var templateJSON string
	var parametersJSON string
	if templateJSON, parametersJSON, err = templateGenerator.GenerateTemplateV2(upgradeContainerService, engine.DefaultGeneratorCode, aksEngineVersion); err != nil {
		return nil, nil, ku.Translator.Errorf("error generating upgrade template: %s", err.Error())
	}

	var template interface{}
	var parameters interface{}

	err = json.Unmarshal([]byte(templateJSON), &template)
	if err != nil {
		return nil, nil, ku.Translator.Errorf("error while unmarshaling the ARM template JSON: %s", err.Error())
	}

	err = json.Unmarshal([]byte(parametersJSON), &parameters)
	if err != nil {
		return nil, nil, ku.Translator.Errorf("error while unmarshaling the ARM parameters JSON: %s", err.Error())
	}

	templateMap := template.(map[string]interface{})
	parametersMap := parameters.(map[string]interface{})

	return templateMap, parametersMap, nil
}

func (ku *Upgrader) getLastVMNameInVMSS(ctx context.Context, resourceGroup string, vmScaleSetName string) (string, error) {
	lastVMName := ""
	for vmScaleSetVMsPage, err := ku.Client.ListVirtualMachineScaleSetVMs(ctx, resourceGroup, vmScaleSetName); vmScaleSetVMsPage.NotDone(); err = vmScaleSetVMsPage.Next() {
		if err != nil {
			return "", err
		}

		vms := vmScaleSetVMsPage.Values()
		if len(vms) > 0 {
			vm := vms[len(vms)-1]
			lastVMName = *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName
		}
	}

	if lastVMName == "" {
		return "", errors.Errorf("failed to get the last VM name in Scale Set %s", vmScaleSetName)
	}

	return lastVMName, nil
}

func (ku *Upgrader) copyCustomPropertiesToNewNode(client armhelpers.KubernetesClient, oldNodeName string, newNodeName string) error {
	// The new node is created without any taints, Kubernetes might schedule some pods on this newly created node before the taints/annotations/labels
	// are copied over from corresponding old node. So drain the new node first before copying over the node properties.
	// Note: SafelyDrainNodeWithClient() sets the Unschedulable of the node to true, set Unschedulable to false in copyCustomNodeProperties
	var cordonDrainTimeout time.Duration
	if ku.cordonDrainTimeout == nil {
		cordonDrainTimeout = defaultCordonDrainTimeout
	} else {
		cordonDrainTimeout = *ku.cordonDrainTimeout
	}
	err := operations.SafelyDrainNodeWithClient(client, ku.logger, newNodeName, cordonDrainTimeout)
	if err != nil {
		ku.logger.Warningf("Error draining agent VM %s. Proceeding with copying node properties. Error: %v", newNodeName, err)
	}

	ch := make(chan struct{}, 1)
	go func() {
		for {
			oldNode, err := client.GetNode(oldNodeName)
			if err != nil {
				ku.logger.Debugf("Failed to get properties of the old node %s: %v", oldNodeName, err)
				time.Sleep(time.Second * 5)
				continue
			}

			newNode, err := client.GetNode(newNodeName)
			if err != nil {
				ku.logger.Debugf("Failed to get properties of the new node %s: %v", newNodeName, err)
				time.Sleep(time.Second * 5)
				continue
			}

			err = ku.copyCustomNodeProperties(client, oldNodeName, oldNode, newNodeName, newNode)
			if err != nil {
				ku.logger.Debugf("Failed to copy custom annotations, labels, taints from old node %s to new node %s: %v", oldNodeName, newNodeName, err)
				time.Sleep(time.Second * 5)
			} else {
				ch <- struct{}{}
			}
		}
	}()

	for {
		select {
		case <-ch:
			ku.logger.Infof("Successfully copied custom annotations, labels, taints from old node %s to new node %s.", oldNodeName, newNodeName)
			return nil
		case <-time.After(nodePropertiesCopyTimeout):
			err := fmt.Errorf("Copying custom annotations, labels, taints from old node %s to new node %s can't complete within %v", oldNodeName, newNodeName, nodePropertiesCopyTimeout)
			ku.logger.Errorf(err.Error())
			return err
		}
	}
}

func (ku *Upgrader) copyCustomNodeProperties(client armhelpers.KubernetesClient, oldNodeName string, oldNode *v1.Node, newNodeName string, newNode *v1.Node) error {
	// copy additional custom annotations from old node to new node
	if oldNode.Annotations != nil {
		if newNode.Annotations == nil {
			newNode.Annotations = map[string]string{}
		}

		for k, v := range oldNode.Annotations {
			if _, ok := newNode.Annotations[k]; !ok {
				newNode.Annotations[k] = strings.Replace(v, oldNodeName, newNodeName, -1)
			}
		}
	}

	// copy additional custom labels from old node to new node
	if oldNode.Labels != nil {
		if newNode.Labels == nil {
			newNode.Labels = map[string]string{}
		}

		for k, v := range oldNode.Labels {
			if _, ok := newNode.Labels[k]; !ok {
				newNode.Labels[k] = strings.Replace(v, oldNodeName, newNodeName, -1)
			}
		}
	}

	// copy Taints from old node to new node
	if oldNode.Spec.Taints != nil {
		newNode.Spec.Taints = append([]v1.Taint{}, oldNode.Spec.Taints...)
		for i := range newNode.Spec.Taints {
			newNode.Spec.Taints[i].Value = strings.Replace(newNode.Spec.Taints[i].Value, oldNodeName, newNodeName, -1)
		}
	}

	newNode, err := client.UpdateNode(newNode)
	if err != nil {
		ku.logger.Warningf("Failed to update the new node %s: %v", newNodeName, err)
		return err
	}

	newNode.Spec.Unschedulable = false
	_, err = client.UpdateNode(newNode)

	return err
}

func (ku *Upgrader) getKubernetesClient(timeout time.Duration) (armhelpers.KubernetesClient, error) {
	apiserverURL := ku.DataModel.Properties.GetMasterFQDN()
	if ku.DataModel.Properties.HostedMasterProfile != nil {
		apiServerListeningPort := 443
		apiserverURL = fmt.Sprintf("https://%s:%d", apiserverURL, apiServerListeningPort)
	}

	return ku.Client.GetKubernetesClient(
		apiserverURL,
		ku.kubeConfig,
		interval,
		timeout)
}

// return unused index within the range of agent indices, or subsequent index
func getAvailableIndex(vms map[int]*vmInfo) int {
	maxIndex := 0

	for indx := range vms {
		if indx > maxIndex {
			maxIndex = indx
		}
	}

	for indx := 0; indx < maxIndex; indx++ {
		if _, found := vms[indx]; !found {
			return indx
		}
	}

	return maxIndex + 1
}

// isNodeReady returns true if a node is ready; false otherwise.
// Copied from: https://github.com/kubernetes/kubernetes/blob/886e04f1fffbb04faf8a9f9ee141143b2684ae68/pkg/api/v1/node/util.go#L40
func isNodeReady(node *v1.Node) bool {
	for _, c := range node.Status.Conditions {
		if c.Type == v1.NodeReady {
			return c.Status == v1.ConditionTrue
		}
	}
	return false
}
