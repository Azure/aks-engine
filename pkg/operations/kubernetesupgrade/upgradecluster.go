// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/armhelpers/utils"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// IsVMSSToBeUpgradedCb - Call back for checking whether the given vmss is to be upgraded or not.
type IsVMSSToBeUpgradedCb func(vmss string, cs *api.ContainerService) bool

// ClusterTopology contains resources of the cluster the upgrade operation
// is targeting
type ClusterTopology struct {
	DataModel      *api.ContainerService
	SubscriptionID string
	Location       string
	ResourceGroup  string
	NameSuffix     string

	AgentPoolsToUpgrade map[string]bool
	AgentPools          map[string]*AgentPoolTopology

	AgentPoolScaleSetsToUpgrade []AgentPoolScaleSet

	MasterVMs         *[]compute.VirtualMachine
	UpgradedMasterVMs *[]compute.VirtualMachine

	IsVMSSToBeUpgraded IsVMSSToBeUpgradedCb
}

// AgentPoolScaleSet contains necessary data required to upgrade a VMSS
type AgentPoolScaleSet struct {
	Name         string
	Sku          compute.Sku
	Location     string
	IsWindows    bool
	VMsToUpgrade []AgentPoolScaleSetVM
}

// AgentPoolScaleSetVM represents a VM in a VMSS
type AgentPoolScaleSetVM struct {
	Name       string
	InstanceID string
}

// AgentPoolTopology contains agent VMs in a single pool
type AgentPoolTopology struct {
	Identifier       *string
	Name             *string
	AgentVMs         *[]compute.VirtualMachine
	UpgradedAgentVMs *[]compute.VirtualMachine
}

// UpgradeCluster upgrades a cluster with Orchestrator version X.X to version Y.Y.
// Right now upgrades are supported for Kubernetes cluster only.
type UpgradeCluster struct {
	Translator *i18n.Translator
	Logger     *logrus.Entry
	ClusterTopology
	Client             armhelpers.AKSEngineClient
	StepTimeout        *time.Duration
	CordonDrainTimeout *time.Duration
	UpgradeWorkFlow    UpgradeWorkFlow
	Force              bool
	ControlPlaneOnly   bool
}

// MasterVMNamePrefix is the prefix for all master VM names for Kubernetes clusters
const MasterVMNamePrefix = "k8s-master-"

// MasterPoolName pool name
const MasterPoolName = "master"

// UpgradeCluster runs the workflow to upgrade a Kubernetes cluster.
func (uc *UpgradeCluster) UpgradeCluster(az armhelpers.AKSEngineClient, kubeConfig string, aksEngineVersion string) error {
	uc.MasterVMs = &[]compute.VirtualMachine{}
	uc.UpgradedMasterVMs = &[]compute.VirtualMachine{}
	uc.AgentPools = make(map[string]*AgentPoolTopology)

	var kubeClient armhelpers.KubernetesClient
	if az != nil {
		timeout := time.Duration(60) * time.Minute
		k, err := az.GetKubernetesClient("", kubeConfig, interval, timeout)
		if err != nil {
			uc.Logger.Warnf("Failed to get a Kubernetes client: %v", err)
		}
		kubeClient = k
	}

	if err := uc.getClusterNodeStatus(kubeClient, uc.ResourceGroup); err != nil {
		return uc.Translator.Errorf("Error while querying ARM for resources: %+v", err)
	}

	kc := uc.DataModel.Properties.OrchestratorProfile.KubernetesConfig
	if kc != nil && kc.IsClusterAutoscalerEnabled() && !uc.ControlPlaneOnly {
		// pause the cluster-autoscaler before running upgrade and resume it afterward
		uc.Logger.Info("Pausing cluster autoscaler, replica count: 0")
		count, err := uc.SetClusterAutoscalerReplicaCount(kubeClient, 0)
		if err != nil {
			uc.Logger.Errorf("Failed to pause cluster-autoscaler: %v", err)
			if !uc.Force {
				return err
			}
		} else {
			if err == nil {
				defer func() {
					uc.Logger.Infof("Resuming cluster autoscaler, replica count: %d", count)
					if _, err = uc.SetClusterAutoscalerReplicaCount(kubeClient, count); err != nil {
						uc.Logger.Errorf("Failed to resume cluster-autoscaler: %v", err)
					}
				}()
			}
		}
	}

	upgradeVersion := uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion
	what := "control plane and all nodes"
	if uc.ControlPlaneOnly {
		what = "control plane nodes"
	}
	uc.Logger.Infof("Upgrading %s to Kubernetes version %s", what, upgradeVersion)

	if err := uc.getUpgradeWorkflow(kubeConfig, aksEngineVersion).RunUpgrade(); err != nil {
		return err
	}

	what = "Cluster"
	if uc.ControlPlaneOnly {
		what = "Control plane"
	}
	uc.Logger.Infof("%s upgraded successfully to Kubernetes version %s", what, upgradeVersion)
	return nil
}

// SetClusterAutoscalerReplicaCount changes the replica count of a cluster-autoscaler deployment.
func (uc *UpgradeCluster) SetClusterAutoscalerReplicaCount(kubeClient armhelpers.KubernetesClient, replicaCount int32) (int32, error) {
	if kubeClient == nil {
		return 0, errors.New("no kubernetes client")
	}
	var count int32
	var err error
	const namespace, name, retries = "kube-system", "cluster-autoscaler", 10
	for attempt := 0; attempt < retries; attempt++ {
		deployment, getErr := kubeClient.GetDeployment(namespace, name)
		err = getErr
		if getErr == nil {
			count = *deployment.Spec.Replicas
			deployment.Spec.Replicas = &replicaCount
			if _, err = kubeClient.UpdateDeployment(namespace, deployment); err == nil {
				break
			}
		}
		sleepTime := time.Duration(rand.Intn(5))
		uc.Logger.Warnf("Failed to update cluster-autoscaler deployment: %v", err)
		uc.Logger.Infof("Retry updating cluster-autoscaler after %d seconds", sleepTime)
		time.Sleep(sleepTime * time.Second)
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (uc *UpgradeCluster) getUpgradeWorkflow(kubeConfig string, aksEngineVersion string) UpgradeWorkFlow {
	if uc.UpgradeWorkFlow != nil {
		return uc.UpgradeWorkFlow
	}
	u := &Upgrader{}
	u.Init(uc.Translator, uc.Logger, uc.ClusterTopology, uc.Client, kubeConfig, uc.StepTimeout, uc.CordonDrainTimeout, aksEngineVersion, uc.ControlPlaneOnly)
	return u
}

func (uc *UpgradeCluster) getClusterNodeStatus(kubeClient armhelpers.KubernetesClient, resourceGroup string) error {
	goalVersion := uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	for vmScaleSetPage, err := uc.Client.ListVirtualMachineScaleSets(ctx, resourceGroup); vmScaleSetPage.NotDone(); err = vmScaleSetPage.NextWithContext(ctx) {
		if err != nil {
			return err
		}
		for _, vmScaleSet := range vmScaleSetPage.Values() {
			if uc.IsVMSSToBeUpgraded != nil && !uc.IsVMSSToBeUpgraded(*vmScaleSet.Name, uc.DataModel) {
				continue
			}
			for vmScaleSetVMsPage, err := uc.Client.ListVirtualMachineScaleSetVMs(ctx, resourceGroup, *vmScaleSet.Name); vmScaleSetVMsPage.NotDone(); err = vmScaleSetVMsPage.NextWithContext(ctx) {
				if err != nil {
					return err
				}
				// set agent pool node count to match VMSS capacity
				for _, pool := range uc.ClusterTopology.DataModel.Properties.AgentPoolProfiles {
					if poolName, _, _ := utils.VmssNameParts(*vmScaleSet.Name); poolName == pool.Name {
						pool.Count = int(*vmScaleSet.Sku.Capacity)
						break
					}
				}
				scaleSetToUpgrade := AgentPoolScaleSet{
					Name:     *vmScaleSet.Name,
					Sku:      *vmScaleSet.Sku,
					Location: *vmScaleSet.Location,
				}
				if vmScaleSet.VirtualMachineProfile != nil &&
					vmScaleSet.VirtualMachineProfile.OsProfile != nil &&
					vmScaleSet.VirtualMachineProfile.OsProfile.WindowsConfiguration != nil {
					scaleSetToUpgrade.IsWindows = true
					uc.Logger.Infof("Set isWindows flag for vmss %s.", *vmScaleSet.Name)
				}
				for _, vm := range vmScaleSetVMsPage.Values() {
					currentVersion := uc.getNodeVersion(kubeClient, strings.ToLower(*vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName), vm.Tags, *vm.VirtualMachineScaleSetVMProperties.LatestModelApplied)
					if uc.Force {
						if currentVersion == "" {
							currentVersion = "Unknown"
						}
					}

					if currentVersion == "" {
						uc.Logger.Infof("Skipping VM: %s for upgrade as the orchestrator version could not be determined.", *vm.Name)
						continue
					}
					if uc.Force || currentVersion != goalVersion {
						uc.Logger.Infof(
							"VM %s in VMSS %s has a current version of %s and a desired version of %s. Upgrading this node.",
							*vm.Name,
							*vmScaleSet.Name,
							currentVersion,
							goalVersion,
						)
						scaleSetToUpgrade.VMsToUpgrade = append(
							scaleSetToUpgrade.VMsToUpgrade,
							AgentPoolScaleSetVM{
								Name:       *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName,
								InstanceID: *vm.InstanceID,
							},
						)
					}
				}
				uc.AgentPoolScaleSetsToUpgrade = append(uc.AgentPoolScaleSetsToUpgrade, scaleSetToUpgrade)
			}
		}
	}

	for vmListPage, err := uc.Client.ListVirtualMachines(ctx, resourceGroup); vmListPage.NotDone(); err = vmListPage.Next() {
		if err != nil {
			return err
		}

		for _, vm := range vmListPage.Values() {
			// Windows VMs contain a substring of the name suffix
			if !strings.Contains(*(vm.Name), uc.NameSuffix) && !strings.Contains(*(vm.Name), uc.NameSuffix[:4]+"k8s") {
				uc.Logger.Infof("Skipping VM: %s for upgrade as it does not belong to cluster with expected name suffix: %s",
					*vm.Name, uc.NameSuffix)
				continue
			}
			currentVersion := uc.getNodeVersion(kubeClient, strings.ToLower(*vm.Name), vm.Tags, true)

			if uc.Force {
				if currentVersion == "" {
					currentVersion = "Unknown"
				}
				uc.addVMToUpgradeSets(vm, currentVersion)
			} else {
				if currentVersion == "" {
					uc.Logger.Infof("Skipping VM: %s for upgrade as the orchestrator version could not be determined.", *vm.Name)
					continue
				}
				// If the current version is different than the desired version then we add the VM to the list of VMs to upgrade.
				if currentVersion != goalVersion {
					if !uc.DataModel.Properties.IsHostedMasterProfile() {
						if err := uc.upgradable(currentVersion); err != nil {
							return err
						}
					}
					uc.addVMToUpgradeSets(vm, currentVersion)
				} else if currentVersion == goalVersion {
					uc.addVMToFinishedSets(vm, currentVersion)
				}
			}
		}
	}

	return nil
}

func (uc *UpgradeCluster) upgradable(currentVersion string) error {
	nodeVersion := &api.OrchestratorProfile{
		OrchestratorType:    api.Kubernetes,
		OrchestratorVersion: currentVersion,
	}
	targetVersion := uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion

	orch, err := api.GetOrchestratorVersionProfile(nodeVersion, uc.DataModel.Properties.HasWindows())
	if err != nil {
		return err
	}

	for _, up := range orch.Upgrades {
		if up.OrchestratorVersion == targetVersion {
			return nil
		}
	}
	return errors.Errorf("%s cannot be upgraded to %s", currentVersion, targetVersion)
}

// getNodeVersion returns a node's current Kubernetes version via Kubernetes API or VM tag.
// For VMSS nodes, make sure OsProfile.ComputerName instead of VM name is used as the name here
// because the former is used as the K8s node name.
// Also, if the latest VMSS model is applied, then we can get the version info from the tags.
// Otherwise, we have to get version via K8s API. This is because VMSS does not support tags
// for individual instances and old/new instances have the same tags.
func (uc *UpgradeCluster) getNodeVersion(client armhelpers.KubernetesClient, name string, tags map[string]*string, getVersionFromTags bool) string {
	if getVersionFromTags {
		if tags != nil && tags["orchestrator"] != nil {
			parts := strings.Split(*tags["orchestrator"], ":")
			if len(parts) == 2 {
				return parts[1]
			}
		}

		uc.Logger.Warnf("Expected tag \"orchestrator\" not found for VM: %s. Using Kubernetes API to retrieve Kubernetes version.", name)
	}

	if client != nil {
		node, err := client.GetNode(name)
		if err == nil {
			return strings.TrimPrefix(node.Status.NodeInfo.KubeletVersion, "v")
		}
		uc.Logger.Warnf("Failed to get node %s: %v", name, err)
	}
	return ""
}

func (uc *UpgradeCluster) addVMToAgentPool(vm compute.VirtualMachine, isUpgradableVM bool) error {
	var poolIdentifier string
	var poolPrefix string
	var err error
	var vmPoolName string

	if vm.Tags != nil && vm.Tags["poolName"] != nil {
		vmPoolName = *vm.Tags["poolName"]
	} else {
		uc.Logger.Infof("poolName tag not found for VM: %s.", *vm.Name)
		// If there's only one agent pool, assume this VM is a member.
		agentPools := []string{}
		for k := range uc.AgentPoolsToUpgrade {
			if !strings.HasPrefix(k, "master") {
				agentPools = append(agentPools, k)
			}
		}
		if len(agentPools) == 1 {
			vmPoolName = agentPools[0]
		}
	}
	if vmPoolName == "" {
		uc.Logger.Warnf("Couldn't determine agent pool membership for VM: %s.", *vm.Name)
		return nil
	}

	uc.Logger.Infof("Evaluating VM: %s in pool: %s...", *vm.Name, vmPoolName)
	if vmPoolName == "" {
		uc.Logger.Infof("VM: %s does not contain `poolName` tag, skipping.", *vm.Name)
		return nil
	} else if !uc.AgentPoolsToUpgrade[vmPoolName] {
		uc.Logger.Infof("Skipping upgrade of VM: %s in pool: %s.", *vm.Name, vmPoolName)
		return nil
	}

	if vm.StorageProfile.OsDisk.OsType == compute.Windows {
		poolPrefix, _, _, _, err = utils.WindowsVMNameParts(*vm.Name)
		if err != nil {
			uc.Logger.Errorf(err.Error())
			return err
		}
		if !strings.Contains(uc.NameSuffix, poolPrefix) {
			uc.Logger.Infof("Skipping VM: %s for upgrade as it does not belong to cluster with expected name suffix: %s",
				*vm.Name, uc.NameSuffix)
			return nil
		}

		// The k8s Windows VM Naming Format was previously "^([a-fA-F0-9]{5})([0-9a-zA-Z]{3})([a-zA-Z0-9]{4,6})$" (i.e.: 50621k8s9000)
		// The k8s Windows VM Naming Format is now "^([a-fA-F0-9]{4})([0-9a-zA-Z]{3})([0-9]{3,8})$" (i.e.: 1708k8s020)
		// The pool identifier is made of the first 11 or 9 characters
		if string((*vm.Name)[8]) == "9" {
			poolIdentifier = (*vm.Name)[:11]
		} else {
			poolIdentifier = (*vm.Name)[:9]
		}
	} else { // vm.StorageProfile.OsDisk.OsType == compute.Linux
		poolIdentifier, poolPrefix, _, err = utils.K8sLinuxVMNameParts(*vm.Name)
		if err != nil {
			uc.Logger.Errorf(err.Error())
			return err
		}

		if !strings.EqualFold(uc.NameSuffix, poolPrefix) {
			uc.Logger.Infof("Skipping VM: %s for upgrade as it does not belong to cluster with expected name suffix: %s",
				*vm.Name, uc.NameSuffix)
			return nil
		}
	}

	if uc.AgentPools[poolIdentifier] == nil {
		uc.AgentPools[poolIdentifier] =
			&AgentPoolTopology{&poolIdentifier, &vmPoolName, &[]compute.VirtualMachine{}, &[]compute.VirtualMachine{}}
	}

	orchestrator := "unknown"
	if vm.Tags != nil && vm.Tags["orchestrator"] != nil {
		orchestrator = *vm.Tags["orchestrator"]
	}
	//TODO(sterbrec): extract this from add to agentPool
	// separate the upgrade/skip decision from the agentpool composition
	if isUpgradableVM {
		uc.Logger.Infof("Adding Agent VM: %s, orchestrator: %s to pool: %s (AgentVMs)",
			*vm.Name, orchestrator, poolIdentifier)
		*uc.AgentPools[poolIdentifier].AgentVMs = append(*uc.AgentPools[poolIdentifier].AgentVMs, vm)
	} else {
		uc.Logger.Infof("Adding Agent VM: %s, orchestrator: %s to pool: %s (UpgradedAgentVMs)",
			*vm.Name, orchestrator, poolIdentifier)
		*uc.AgentPools[poolIdentifier].UpgradedAgentVMs = append(*uc.AgentPools[poolIdentifier].UpgradedAgentVMs, vm)
	}

	return nil
}

func (uc *UpgradeCluster) addVMToUpgradeSets(vm compute.VirtualMachine, currentVersion string) {
	if strings.Contains(*(vm.Name), MasterVMNamePrefix) {
		uc.Logger.Infof("Master VM name: %s, orchestrator: %s (MasterVMs)", *vm.Name, currentVersion)
		*uc.MasterVMs = append(*uc.MasterVMs, vm)
	} else {
		if err := uc.addVMToAgentPool(vm, true); err != nil {
			uc.Logger.Errorf("Failed to add VM %s to agent pool: %s", *vm.Name, err)
		}
	}
}

func (uc *UpgradeCluster) addVMToFinishedSets(vm compute.VirtualMachine, currentVersion string) {
	if strings.Contains(*(vm.Name), MasterVMNamePrefix) {
		uc.Logger.Infof("Master VM name: %s, orchestrator: %s (UpgradedMasterVMs)", *vm.Name, currentVersion)
		*uc.UpgradedMasterVMs = append(*uc.UpgradedMasterVMs, vm)
	} else {
		if err := uc.addVMToAgentPool(vm, false); err != nil {
			uc.Logger.Errorf("Failed to add VM %s to agent pool: %s", *vm.Name, err)
		}
	}
}
