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
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
	"github.com/blang/semver"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
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
	Client      armhelpers.AKSEngineClient
	StepTimeout *time.Duration
}

// MasterVMNamePrefix is the prefix for all master VM names for Kubernetes clusters
const MasterVMNamePrefix = "k8s-master-"

// MasterPoolName pool name
const MasterPoolName = "master"

// UpgradeCluster runs the workflow to upgrade a Kubernetes cluster.
func (uc *UpgradeCluster) UpgradeCluster(subscriptionID uuid.UUID, az armhelpers.AKSEngineClient, kubeConfig, resourceGroup string,
	cs *api.ContainerService, nameSuffix string, agentPoolsToUpgrade []string, aksengineVersion string) error {
	uc.ClusterTopology = ClusterTopology{}
	uc.SubscriptionID = subscriptionID.String()
	uc.ResourceGroup = resourceGroup
	uc.DataModel = cs
	uc.NameSuffix = nameSuffix
	uc.MasterVMs = &[]compute.VirtualMachine{}
	uc.UpgradedMasterVMs = &[]compute.VirtualMachine{}
	uc.AgentPools = make(map[string]*AgentPoolTopology)
	uc.AgentPoolsToUpgrade = make(map[string]bool)

	for _, poolName := range agentPoolsToUpgrade {
		uc.AgentPoolsToUpgrade[poolName] = true
	}
	uc.AgentPoolsToUpgrade[MasterPoolName] = true

	if err := uc.getClusterNodeStatus(subscriptionID, az, resourceGroup, kubeConfig); err != nil {
		return uc.Translator.Errorf("Error while querying ARM for resources: %+v", err)
	}

	var upgrader UpgradeWorkFlow
	upgradeVersion := uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion
	uc.Logger.Infof("Upgrading to Kubernetes version %s\n", upgradeVersion)
	switch {
	case strings.HasPrefix(upgradeVersion, "1.6."):
		upgrader16 := &Kubernetes16upgrader{}
		upgrader16.Init(uc.Translator, uc.Logger, uc.ClusterTopology, uc.Client, kubeConfig, uc.StepTimeout, aksengineVersion)
		upgrader = upgrader16

	case strings.HasPrefix(upgradeVersion, "1.7."):
		upgrader17 := &Kubernetes17upgrader{}
		upgrader17.Init(uc.Translator, uc.Logger, uc.ClusterTopology, uc.Client, kubeConfig, uc.StepTimeout, aksengineVersion)
		upgrader = upgrader17

	case strings.HasPrefix(upgradeVersion, "1.8."):
		upgrader18 := &Kubernetes18upgrader{}
		upgrader18.Init(uc.Translator, uc.Logger, uc.ClusterTopology, uc.Client, kubeConfig, uc.StepTimeout, aksengineVersion)
		upgrader = upgrader18

	case strings.HasPrefix(upgradeVersion, "1.9."),
		strings.HasPrefix(upgradeVersion, "1.10."),
		strings.HasPrefix(upgradeVersion, "1.11."),
		strings.HasPrefix(upgradeVersion, "1.12."),
		strings.HasPrefix(upgradeVersion, "1.13."):
		u := &Upgrader{}
		u.Init(uc.Translator, uc.Logger, uc.ClusterTopology, uc.Client, kubeConfig, uc.StepTimeout, aksengineVersion)
		upgrader = u

	default:
		return uc.Translator.Errorf("Upgrade to Kubernetes version %s is not supported", upgradeVersion)
	}

	if err := upgrader.RunUpgrade(); err != nil {
		return err
	}

	uc.Logger.Infof("Cluster upgraded successfully to Kubernetes version %s\n", upgradeVersion)
	return nil
}

func (uc *UpgradeCluster) getClusterNodeStatus(subscriptionID uuid.UUID, az armhelpers.AKSEngineClient, resourceGroup, kubeConfig string) error {
	targetOrchestratorTypeVersion := fmt.Sprintf("%s:%s", uc.DataModel.Properties.OrchestratorProfile.OrchestratorType, uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion)

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

	for vmScaleSetPage, err := uc.Client.ListVirtualMachineScaleSets(ctx, resourceGroup); vmScaleSetPage.NotDone(); err = vmScaleSetPage.Next() {
		if err != nil {
			return err
		}
		for _, vmScaleSet := range vmScaleSetPage.Values() {
			for vmScaleSetVMsPage, err := uc.Client.ListVirtualMachineScaleSetVMs(ctx, resourceGroup, *vmScaleSet.Name); vmScaleSetVMsPage.NotDone(); err = vmScaleSetVMsPage.Next() {
				if err != nil {
					return err
				}
				scaleSetToUpgrade := AgentPoolScaleSet{
					Name:     *vmScaleSet.Name,
					Sku:      *vmScaleSet.Sku,
					Location: *vmScaleSet.Location,
				}
				for _, vm := range vmScaleSetVMsPage.Values() {
					scaleSetVMOrchestratorTypeAndVersion := uc.getClusterNodeVersion(kubeClient, *vm.Name, vm.Tags)
					if scaleSetVMOrchestratorTypeAndVersion == "" {
						uc.Logger.Infof("Skipping VM: %s for upgrade as the orchestrator version could not be determined.", *vm.Name)
						continue
					}

					if scaleSetVMOrchestratorTypeAndVersion != targetOrchestratorTypeVersion {
						// This condition is a scale set VM that is an older version and should be handled
						uc.Logger.Infof(
							"VM %s in VMSS %s has a current version of %s and a desired version of %s. Upgrading this node.",
							*vm.Name,
							*vmScaleSet.Name,
							scaleSetVMOrchestratorTypeAndVersion,
							targetOrchestratorTypeVersion,
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
			vmOrchestratorTypeAndVersion := uc.getClusterNodeVersion(kubeClient, *vm.Name, vm.Tags)
			if vmOrchestratorTypeAndVersion == "" {
				uc.Logger.Infof("Skipping VM: %s for upgrade as the orchestrator version could not be determined.", *vm.Name)
				continue
			}

			if vmOrchestratorTypeAndVersion != targetOrchestratorTypeVersion {
				if strings.Contains(*(vm.Name), MasterVMNamePrefix) {
					if !strings.Contains(*(vm.Name), uc.NameSuffix) {
						uc.Logger.Infof("Skipping VM: %s for upgrade as it does not belong to cluster with expected name suffix: %s\n",
							*vm.Name, uc.NameSuffix)
						continue
					}
					if err := uc.upgradable(vmOrchestratorTypeAndVersion); err != nil {
						return err
					}
					uc.Logger.Infof("Master VM name: %s, orchestrator: %s (MasterVMs)\n", *vm.Name, vmOrchestratorTypeAndVersion)
					*uc.MasterVMs = append(*uc.MasterVMs, vm)
				} else {
					if err := uc.upgradable(vmOrchestratorTypeAndVersion); err != nil {
						return err
					}
					uc.addVMToAgentPool(vm, true)
				}
			} else if vmOrchestratorTypeAndVersion == targetOrchestratorTypeVersion {
				if strings.Contains(*(vm.Name), MasterVMNamePrefix) {
					if !strings.Contains(*(vm.Name), uc.NameSuffix) {
						uc.Logger.Infof("Not adding VM: %s to upgraded list as it does not belong to cluster with expected name suffix: %s\n",
							*vm.Name, uc.NameSuffix)
						continue
					}
					uc.Logger.Infof("Master VM name: %s, orchestrator: %s (UpgradedMasterVMs)\n", *vm.Name, vmOrchestratorTypeAndVersion)
					*uc.UpgradedMasterVMs = append(*uc.UpgradedMasterVMs, vm)
				} else {
					uc.addVMToAgentPool(vm, false)
				}
			}
		}
	}

	return nil
}

// getClusterNodeVersion returns a node's "orchestrator:version" via Kubernetes API or VM tag.
func (uc *UpgradeCluster) getClusterNodeVersion(client armhelpers.KubernetesClient, name string, tags map[string]*string) string {
	if tags != nil && tags["orchestrator"] != nil {
		return *tags["orchestrator"]
	}
	uc.Logger.Warnf("Expected tag \"orchestrator\" not found for VM: %s", name)
	if client != nil {
		node, err := client.GetNode(name)
		if err == nil {
			return api.Kubernetes + ":" + strings.TrimPrefix(node.Status.NodeInfo.KubeletVersion, "v")
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
					return api.Kubernetes + ":" + strings.TrimPrefix(node.Status.NodeInfo.KubeletVersion, "v")
				}
				uc.Logger.Warnf("Failed to get node %s: %v", vmssName, err)
			}
		}
	}
	return ""
}

func (uc *UpgradeCluster) upgradable(vmOrchestratorTypeAndVersion string) error {
	arr := strings.Split(vmOrchestratorTypeAndVersion, ":")
	if len(arr) != 2 {
		return errors.Errorf("Unsupported orchestrator tag format %s", vmOrchestratorTypeAndVersion)
	}
	currentVer, err := semver.Make(arr[1])
	if err != nil {
		return errors.Errorf("Unsupported orchestrator version format %s", currentVer.String())
	}
	csOrch := &api.OrchestratorProfile{
		OrchestratorType:    api.Kubernetes,
		OrchestratorVersion: currentVer.String(),
	}
	orch, err := api.GetOrchestratorVersionProfile(csOrch, uc.DataModel.Properties.HasWindows())
	if err != nil {
		return err
	}
	for _, up := range orch.Upgrades {
		if up.OrchestratorVersion == uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion {
			return nil
		}
	}
	return errors.Errorf("%s cannot be upgraded to %s", vmOrchestratorTypeAndVersion, uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion)
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
			uc.Logger.Infof("Skipping VM: %s for upgrade as it does not belong to cluster with expected name suffix: %s\n",
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
			uc.Logger.Infof("Skipping VM: %s for upgrade as it does not belong to cluster with expected name suffix: %s\n",
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
	if isUpgradableVM {
		uc.Logger.Infof("Adding Agent VM: %s, orchestrator: %s to pool: %s (AgentVMs)\n",
			*vm.Name, orchestrator, poolIdentifier)
		*uc.AgentPools[poolIdentifier].AgentVMs = append(*uc.AgentPools[poolIdentifier].AgentVMs, vm)
	} else {
		uc.Logger.Infof("Adding Agent VM: %s, orchestrator: %s to pool: %s (UpgradedAgentVMs)\n",
			*vm.Name, orchestrator, poolIdentifier)
		*uc.AgentPools[poolIdentifier].UpgradedAgentVMs = append(*uc.AgentPools[poolIdentifier].UpgradedAgentVMs, vm)
	}

	return nil
}

/* WriteTemplate writes upgrade template to a folder
func WriteTemplate(
	translator *i18n.Translator,
	upgradeContainerService *api.ContainerService,
	templateMap map[string]interface{}, parametersMap map[string]interface{}) {
	updatedTemplateJSON, _ := json.Marshal(templateMap)
	parametersJSON, _ := json.Marshal(parametersMap)

	templateapp, err := engine.PrettyPrintArmTemplate(string(updatedTemplateJSON))
	if err != nil {
		logrus.Fatalf("error pretty printing template: %s \n", err.Error())
	}
	parametersapp, e := engine.PrettyPrintJSON(string(parametersJSON))
	if e != nil {
		logrus.Fatalf("error pretty printing template parameters: %s \n", e.Error())
	}
	outputDirectory := path.Join("_output", upgradeContainerService.Properties.MasterProfile.DNSPrefix, "Upgrade")
	writer := &engine.ArtifactWriter{
		Translator: translator,
	}
	if err := writer.WriteTLSArtifacts(upgradeContainerService, "vlabs", templateapp, parametersapp, outputDirectory, false, false); err != nil {
		logrus.Fatalf("error writing artifacts: %s\n", err.Error())
	}
}*/
