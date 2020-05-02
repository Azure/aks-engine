// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"regexp"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/armhelpers/utils"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/operations/kubernetesupgrade"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	upgradeName             = "upgrade"
	upgradeShortDescription = "Upgrade an existing AKS Engine-created Kubernetes cluster"
	upgradeLongDescription  = "Upgrade an existing AKS Engine-created Kubernetes cluster, one node at a time"
)

type upgradeCmd struct {
	authProvider

	// user input
	resourceGroupName                        string
	apiModelPath                             string
	deploymentDirectory                      string
	upgradeVersion                           string
	location                                 string
	kubeconfigPath                           string
	timeoutInMinutes                         int
	cordonDrainTimeoutInMinutes              int
	force                                    bool
	controlPlaneOnly                         bool
	disableClusterInitComponentDuringUpgrade bool

	// derived
	containerService    *api.ContainerService
	apiVersion          string
	client              armhelpers.AKSEngineClient
	locale              *gotext.Locale
	nameSuffix          string
	agentPoolsToUpgrade map[string]bool
	timeout             *time.Duration
	cordonDrainTimeout  *time.Duration
}

func newUpgradeCmd() *cobra.Command {
	uc := upgradeCmd{
		authProvider: &authArgs{},
	}

	upgradeCmd := &cobra.Command{
		Use:   upgradeName,
		Short: upgradeShortDescription,
		Long:  upgradeLongDescription,
		RunE:  uc.run,
	}

	f := upgradeCmd.Flags()
	f.StringVarP(&uc.location, "location", "l", "", "location the cluster is deployed in (required)")
	f.StringVarP(&uc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed (required)")
	f.StringVarP(&uc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file")
	f.StringVar(&uc.deploymentDirectory, "deployment-dir", "", "the location of the output from `generate`")
	f.StringVarP(&uc.upgradeVersion, "upgrade-version", "k", "", "desired kubernetes version (required)")
	f.StringVarP(&uc.kubeconfigPath, "kubeconfig", "b", "", "the path of the kubeconfig file")
	f.IntVar(&uc.timeoutInMinutes, "vm-timeout", -1, "how long to wait for each vm to be upgraded in minutes")
	f.IntVar(&uc.cordonDrainTimeoutInMinutes, "cordon-drain-timeout", -1, "how long to wait for each vm to be cordoned in minutes")
	f.BoolVarP(&uc.force, "force", "f", false, "force upgrading the cluster to desired version. Allows same version upgrades and downgrades.")
	f.BoolVarP(&uc.controlPlaneOnly, "control-plane-only", "", false, "upgrade control plane VMs only, do not upgrade node pools")
	addAuthFlags(uc.getAuthArgs(), f)

	_ = f.MarkDeprecated("deployment-dir", "deployment-dir is no longer required for scale or upgrade. Please use --api-model.")

	return upgradeCmd
}

func (uc *upgradeCmd) validate(cmd *cobra.Command) error {
	var err error

	uc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	if uc.resourceGroupName == "" {
		_ = cmd.Usage()
		return errors.New("--resource-group must be specified")
	}

	if uc.location == "" {
		_ = cmd.Usage()
		return errors.New("--location must be specified")
	}
	uc.location = helpers.NormalizeAzureRegion(uc.location)

	if uc.timeoutInMinutes != -1 {
		timeout := time.Duration(uc.timeoutInMinutes) * time.Minute
		uc.timeout = &timeout
	}

	if uc.cordonDrainTimeoutInMinutes != -1 {
		cordonDrainTimeout := time.Duration(uc.cordonDrainTimeoutInMinutes) * time.Minute
		uc.cordonDrainTimeout = &cordonDrainTimeout
	}

	if uc.upgradeVersion == "" {
		_ = cmd.Usage()
		return errors.New("--upgrade-version must be specified")
	}

	if uc.apiModelPath == "" && uc.deploymentDirectory == "" {
		_ = cmd.Usage()
		return errors.New("--api-model must be specified")
	}

	if uc.apiModelPath != "" && uc.deploymentDirectory != "" {
		_ = cmd.Usage()
		return errors.New("ambiguous, please specify only one of --api-model and --deployment-dir")
	}

	return nil
}

func (uc *upgradeCmd) loadCluster() error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	// Load apimodel from the directory.
	if uc.apiModelPath == "" {
		uc.apiModelPath = filepath.Join(uc.deploymentDirectory, apiModelFilename)
	}

	if _, err = os.Stat(uc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api model does not exist (%s)", uc.apiModelPath)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
	}

	// Load the container service.
	uc.containerService, uc.apiVersion, err = apiloader.LoadContainerServiceFromFile(uc.apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	// The cluster-init component is a cluster create-only feature, temporarily disable if enabled
	if i := api.GetComponentsIndexByName(uc.containerService.Properties.OrchestratorProfile.KubernetesConfig.Components, common.ClusterInitComponentName); i > -1 {
		if uc.containerService.Properties.OrchestratorProfile.KubernetesConfig.Components[i].IsEnabled() {
			uc.disableClusterInitComponentDuringUpgrade = true
			uc.containerService.Properties.OrchestratorProfile.KubernetesConfig.Components[i].Enabled = to.BoolPtr(false)
		}
	}

	if uc.containerService.Properties.IsCustomCloudProfile() {
		if err = writeCustomCloudProfile(uc.containerService); err != nil {
			return errors.Wrap(err, "error writing custom cloud profile")
		}
		if err = uc.containerService.Properties.SetCustomCloudSpec(api.AzureCustomCloudSpecParams{
			IsUpgrade: true,
			IsScale:   false,
		}); err != nil {
			return errors.Wrap(err, "error parsing the api model")
		}
	}

	if err = uc.getAuthArgs().validateAuthArgs(); err != nil {
		return err
	}

	if uc.client, err = uc.getAuthArgs().getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	_, err = uc.client.EnsureResourceGroup(ctx, uc.resourceGroupName, uc.location, nil)
	if err != nil {
		return errors.Wrap(err, "error ensuring resource group")
	}

	err = uc.initialize()
	if err != nil {
		return errors.Wrap(err, "error validating the api model")
	}
	return nil
}

func (uc *upgradeCmd) validateTargetVersion() error {
	// Get available upgrades for container service.
	orchestratorInfo, err := api.GetOrchestratorVersionProfile(uc.containerService.Properties.OrchestratorProfile, uc.containerService.Properties.HasWindows())
	if err != nil {
		return errors.Wrap(err, "error getting list of available upgrades")
	}

	found := false
	for _, up := range orchestratorInfo.Upgrades {
		if up.OrchestratorVersion == uc.upgradeVersion {
			found = true
			break
		}
	}
	if !found {
		return errors.Errorf("upgrading from Kubernetes version %s to version %s is not supported. To see a list of available upgrades, use 'aks-engine get-versions --version %s'", uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion, uc.upgradeVersion, uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion)
	}
	return nil
}

func (uc *upgradeCmd) initialize() error {
	if uc.containerService.Location == "" {
		uc.containerService.Location = uc.location
	} else if uc.containerService.Location != uc.location {
		return errors.New("--location does not match api model location")
	}

	if !uc.force {
		err := uc.validateTargetVersion()
		if err != nil {
			return errors.Wrap(err, "Invalid upgrade target version. Consider using --force if you really want to proceed")
		}
	}
	uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion = uc.upgradeVersion

	//allows to identify VMs in the resource group that belong to this cluster.
	uc.nameSuffix = uc.containerService.Properties.GetClusterID()

	log.Infoln(fmt.Sprintf("Upgrading cluster with name suffix: %s", uc.nameSuffix))

	uc.agentPoolsToUpgrade = make(map[string]bool)
	uc.agentPoolsToUpgrade[kubernetesupgrade.MasterPoolName] = true
	for _, agentPool := range uc.containerService.Properties.AgentPoolProfiles {
		uc.agentPoolsToUpgrade[agentPool.Name] = true
	}
	return nil
}

func (uc *upgradeCmd) run(cmd *cobra.Command, args []string) error {
	err := uc.validate(cmd)
	if err != nil {
		return errors.Wrap(err, "validating upgrade command")
	}

	err = uc.loadCluster()
	if err != nil {
		return errors.Wrap(err, "loading existing cluster")
	}

	upgradeCluster := kubernetesupgrade.UpgradeCluster{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
		Logger:             log.NewEntry(log.New()),
		Client:             uc.client,
		StepTimeout:        uc.timeout,
		CordonDrainTimeout: uc.cordonDrainTimeout,
	}

	upgradeCluster.ClusterTopology = kubernetesupgrade.ClusterTopology{}
	upgradeCluster.SubscriptionID = uc.getAuthArgs().SubscriptionID.String()
	upgradeCluster.ResourceGroup = uc.resourceGroupName
	upgradeCluster.DataModel = uc.containerService
	upgradeCluster.NameSuffix = uc.nameSuffix
	upgradeCluster.AgentPoolsToUpgrade = uc.agentPoolsToUpgrade
	upgradeCluster.Force = uc.force
	upgradeCluster.ControlPlaneOnly = uc.controlPlaneOnly

	var kubeConfig string
	if uc.kubeconfigPath != "" {
		var path string
		var content []byte
		path, err = filepath.Abs(uc.kubeconfigPath)
		if err != nil {
			return errors.Wrap(err, "reading --kubeconfig")
		}
		content, err = ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrap(err, "reading --kubeconfig")
		}
		kubeConfig = string(content)
	} else {
		kubeConfig, err = engine.GenerateKubeConfig(uc.containerService.Properties, uc.location)
		if err != nil {
			return errors.Wrap(err, "generating kubeconfig")
		}
	}

	upgradeCluster.IsVMSSToBeUpgraded = isVMSSNameInAgentPoolsArray

	if err = upgradeCluster.UpgradeCluster(uc.client, kubeConfig, BuildTag); err != nil {
		return errors.Wrap(err, "upgrading cluster")
	}

	// Save the new apimodel to reflect the cluster's state.
	// Restore the original cluster-init component enabled value, if it was disabled during upgrade
	if uc.disableClusterInitComponentDuringUpgrade {
		if i := api.GetComponentsIndexByName(uc.containerService.Properties.OrchestratorProfile.KubernetesConfig.Components, common.ClusterInitComponentName); i > -1 {
			uc.containerService.Properties.OrchestratorProfile.KubernetesConfig.Components[i].Enabled = to.BoolPtr(true)
		}
	}
	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
	}
	b, err := apiloader.SerializeContainerService(uc.containerService, uc.apiVersion)
	if err != nil {
		return err
	}

	f := helpers.FileSaver{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
	}
	dir, file := filepath.Split(uc.apiModelPath)
	return f.SaveFile(dir, file, b)
}

// isVMSSNameInAgentPoolsArray is a helper func to filter out any VMSS in the cluster resource group
// that are not participating in the aks-engine-created Kubernetes cluster
func isVMSSNameInAgentPoolsArray(vmss string, cs *api.ContainerService) bool {
	for _, pool := range cs.Properties.AgentPoolProfiles {
		if pool.AvailabilityProfile == api.VirtualMachineScaleSets {
			if pool.OSType == api.Windows {
				re := regexp.MustCompile(`^[0-9]{4}k8s[0]+`)
				if re.FindString(vmss) != "" {
					return true
				}
			} else {
				if poolName, _, _ := utils.VmssNameParts(vmss); poolName == pool.Name {
					return true
				}
			}
		}
	}
	return false
}
