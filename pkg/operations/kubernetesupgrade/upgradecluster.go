// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/armhelpers/utils"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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
}

// AgentPoolScaleSet contains necessary data required to upgrade a VMSS
type AgentPoolScaleSet struct {
	Name         string
	Sku          compute.Sku
	Location     string
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
	Client          armhelpers.AKSEngineClient
	StepTimeout     *time.Duration
	UpgradeWorkFlow UpgradeWorkFlow
	Force           bool
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

	if err := uc.getClusterNodeStatus(az, uc.ResourceGroup, kubeConfig); err != nil {
		return uc.Translator.Errorf("Error while querying ARM for resources: %+v", err)
	}

	upgradeVersion := uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion
	uc.Logger.Infof("Upgrading to Kubernetes version %s", upgradeVersion)

	if err := uc.getUpgradeWorkflow(kubeConfig, aksEngineVersion).RunUpgrade(); err != nil {
		return err
	}

	uc.Logger.Infof("Cluster upgraded successfully to Kubernetes version %s", upgradeVersion)
	return nil
}

func (uc *UpgradeCluster) getUpgradeWorkflow(kubeConfig string, aksEngineVersion string) UpgradeWorkFlow {
	if uc.UpgradeWorkFlow != nil {
		return uc.UpgradeWorkFlow
	}
	u := &Upgrader{}
	u.Init(uc.Translator, uc.Logger, uc.ClusterTopology, uc.Client, kubeConfig, uc.StepTimeout, aksEngineVersion)
	return u
}

func (uc *UpgradeCluster) getClusterNodeStatus(az armhelpers.AKSEngineClient, resourceGroup, kubeConfig string) error {
	goalVersion := uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	var kubeClient armhelpers.KubernetesClient
	if az != nil {
		timeout := time.Duration(60) * time.Minute
		k, err := az.GetKubernetesClient("", kubeConfig, interval, timeout)
		if err != nil {
			uc.Logger.Warnf("Failed to get a Kubernetes client: %v", err)
		}
		kubeClient = k
	}

	for vmScaleSetPage, err := uc.Client.ListVirtualMachineScaleSets(ctx, resourceGroup); vmScaleSetPage.NotDone(); err = vmScaleSetPage.NextWithContext(ctx) {
		if err != nil {
			return err
		}
		for _, vmScaleSet := range vmScaleSetPage.Values() {
			for vmScaleSetVMsPage, err := uc.Client.ListVirtualMachineScaleSetVMs(ctx, resourceGroup, *vmScaleSet.Name); vmScaleSetVMsPage.NotDone(); err = vmScaleSetVMsPage.NextWithContext(ctx) {
				if err != nil {
					return err
				}
				scaleSetToUpgrade := AgentPoolScaleSet{
					Name:     *vmScaleSet.Name,
					Sku:      *vmScaleSet.Sku,
					Location: *vmScaleSet.Location,
				}
				for _, vm := range vmScaleSetVMsPage.Values() {
					currentVersion := uc.getNodeVersion(kubeClient, *vm.Name, vm.Tags)
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
			currentVersion := uc.getNodeVersion(kubeClient, *vm.Name, vm.Tags)

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
func (uc *UpgradeCluster) getNodeVersion(client armhelpers.KubernetesClient, name string, tags map[string]*string) string {
	if tags != nil && tags["orchestrator"] != nil {
		parts := strings.Split(*tags["orchestrator"], ":")
		if len(parts) == 2 {
			return parts[1]
		}
	}
	uc.Logger.Warnf("Expected tag \"orchestrator\" not found for VM: %s. Using Kubernetes API to retrieve Kubernetes version.", name)
	if client != nil {
		node, err := client.GetNode(name)
		if err == nil {
			return strings.TrimPrefix(node.Status.NodeInfo.KubeletVersion, "v")
		}
		uc.Logger.Warnf("Failed to get node %s: %v", name, err)
		// If it's a VMSS cluster, generate the likely Kubernetes node name and try again.
		if strings.Contains(name, "vmss_") {
			parts := strings.Split(name, "_")
			if len(parts) == 2 {
				end := 28 // keep the overall node name at 34 chars or less
				if len(parts[0]) < end {
					end = len(parts[0])
				}
				vmssName := fmt.Sprintf("%s%06s", parts[0][0:end], parts[1])
				node, err := client.GetNode(vmssName)
				if err == nil {
					uc.Logger.Infof("Found VMSS node %s under the name %s", name, vmssName)
					return strings.TrimPrefix(node.Status.NodeInfo.KubeletVersion, "v")
				}
				uc.Logger.Warnf("Failed to get node %s: %v", vmssName, err)
			}
		}
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
		uc.addVMToAgentPool(vm, true)
	}
}

func (uc *UpgradeCluster) addVMToFinishedSets(vm compute.VirtualMachine, currentVersion string) {
	if strings.Contains(*(vm.Name), MasterVMNamePrefix) {
		uc.Logger.Infof("Master VM name: %s, orchestrator: %s (UpgradedMasterVMs)", *vm.Name, currentVersion)
		*uc.UpgradedMasterVMs = append(*uc.UpgradedMasterVMs, vm)
	} else {
		uc.addVMToAgentPool(vm, false)
	}
}
