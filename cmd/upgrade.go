// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/operations/kubernetesupgrade"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	upgradeName             = "upgrade"
	upgradeShortDescription = "Upgrade an existing Kubernetes cluster"
	upgradeLongDescription  = "Upgrade an existing Kubernetes cluster, one minor version at a time"
)

type upgradeCmd struct {
	authArgs

	// user input
	resourceGroupName   string
	deploymentDirectory string
	upgradeVersion      string
	location            string
	timeoutInMinutes    int

	// derived
	containerService    *api.ContainerService
	apiVersion          string
	client              armhelpers.AKSEngineClient
	locale              *gotext.Locale
	nameSuffix          string
	agentPoolsToUpgrade map[string]bool
	timeout             *time.Duration
}

func newUpgradeCmd() *cobra.Command {
	uc := upgradeCmd{}

	upgradeCmd := &cobra.Command{
		Use:   upgradeName,
		Short: upgradeShortDescription,
		Long:  upgradeLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return uc.run(cmd, args)
		},
	}

	f := upgradeCmd.Flags()
	f.StringVarP(&uc.location, "location", "l", "", "location the cluster is deployed in (required)")
	f.StringVarP(&uc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed (required)")
	f.StringVar(&uc.deploymentDirectory, "deployment-dir", "", "the location of the output from `generate` (required)")
	f.StringVarP(&uc.upgradeVersion, "upgrade-version", "k", "", "desired kubernetes version (required)")
	f.IntVar(&uc.timeoutInMinutes, "vm-timeout", -1, "how long to wait for each vm to be upgraded in minutes")
	addAuthFlags(&uc.authArgs, f)

	return upgradeCmd
}

func (uc *upgradeCmd) validate(cmd *cobra.Command) error {
	var err error

	uc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	if uc.resourceGroupName == "" {
		cmd.Usage()
		return errors.New("--resource-group must be specified")
	}

	if uc.location == "" {
		cmd.Usage()
		return errors.New("--location must be specified")
	}
	uc.location = helpers.NormalizeAzureRegion(uc.location)

	if uc.timeoutInMinutes != -1 {
		timeout := time.Duration(uc.timeoutInMinutes) * time.Minute
		uc.timeout = &timeout
	}

	if uc.upgradeVersion == "" {
		cmd.Usage()
		return errors.New("--upgrade-version must be specified")
	}

	if uc.deploymentDirectory == "" {
		cmd.Usage()
		return errors.New("--deployment-dir must be specified")
	}
	return nil
}

func (uc *upgradeCmd) loadCluster(cmd *cobra.Command) error {
	var err error

	if err = uc.authArgs.validateAuthArgs(); err != nil {
		return err
	}

	if uc.client, err = uc.authArgs.getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	_, err = uc.client.EnsureResourceGroup(ctx, uc.resourceGroupName, uc.location, nil)
	if err != nil {
		return errors.Wrap(err, "error ensuring resource group")
	}

	// Load apimodel from the deployment directory.
	apiModelPath := path.Join(uc.deploymentDirectory, "apimodel.json")

	if _, err = os.Stat(apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api model does not exist (%s)", apiModelPath)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
	}

	// Load the container service.
	uc.containerService, uc.apiVersion, err = apiloader.LoadContainerServiceFromFile(apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	if uc.containerService.Location == "" {
		uc.containerService.Location = uc.location
	} else if uc.containerService.Location != uc.location {
		return errors.New("--location does not match api model location")
	}

	// Get available upgrades for container service.
	orchestratorInfo, err := api.GetOrchestratorVersionProfile(uc.containerService.Properties.OrchestratorProfile, uc.containerService.Properties.HasWindows())
	if err != nil {
		return errors.Wrap(err, "error getting list of available upgrades")
	}

	// Add the current version to account for failed upgrades.
	orchestratorInfo.Upgrades = append(orchestratorInfo.Upgrades, &api.OrchestratorProfile{
		OrchestratorType:    uc.containerService.Properties.OrchestratorProfile.OrchestratorType,
		OrchestratorVersion: uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion})

	// Validate desired upgrade version and set goal state.
	found := false
	for _, up := range orchestratorInfo.Upgrades {
		if up.OrchestratorVersion == uc.upgradeVersion {
			uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion = uc.upgradeVersion
			found = true
			break
		}
	}
	if !found {
		return errors.Errorf("upgrading from Kubernetes version %s to version %s is not supported. To see a list of available upgrades, use 'aks-engine get-versions --version %s'", uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion, uc.upgradeVersion, uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	// Read the name suffix from the parameters to identify VMs in the resource group that belong to this cluster.
	templatePath := path.Join(uc.deploymentDirectory, "azuredeploy.json")
	contents, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return errors.Wrap(err, "error reading ARM file")
	}

	var template interface{}
	json.Unmarshal(contents, &template)

	var templateMap, templateParameters, nameSuffixParam map[string]interface{}
	var okType bool

	if templateMap, okType = template.(map[string]interface{}); !okType {
		return errors.Errorf("error asserting data from file %q", templatePath)
	}

	const (
		parametersKey   = "parameters"
		nameSuffixKey   = "nameSuffix"
		defaultValueKey = "defaultValue"
	)

	if templateParameters, okType = templateMap[parametersKey].(map[string]interface{}); !okType {
		return errors.Errorf("error asserting data from key \"%s\" in file %q",
			parametersKey, templatePath)
	}

	if nameSuffixParam, okType = templateParameters[nameSuffixKey].(map[string]interface{}); !okType {
		return errors.Errorf("error asserting data from key \"%s.%s\" in file %q",
			parametersKey, nameSuffixKey, templatePath)
	}

	if uc.nameSuffix, okType = nameSuffixParam[defaultValueKey].(string); !okType {
		return errors.Errorf("error asserting data from key \"%s.%s.%s\" in file %q",
			parametersKey, nameSuffixKey, defaultValueKey, templatePath)
	}

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
		log.Fatalf("Error validating upgrade command: %v", err)
	}

	err = uc.loadCluster(cmd)
	if err != nil {
		log.Fatalf("Error loading existing cluster: %v", err)
	}

	upgradeCluster := kubernetesupgrade.UpgradeCluster{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
		Logger:      log.NewEntry(log.New()),
		Client:      uc.client,
		StepTimeout: uc.timeout,
	}

	upgradeCluster.ClusterTopology = kubernetesupgrade.ClusterTopology{}
	upgradeCluster.SubscriptionID = uc.authArgs.SubscriptionID.String()
	upgradeCluster.ResourceGroup = uc.resourceGroupName
	upgradeCluster.DataModel = uc.containerService
	upgradeCluster.NameSuffix = uc.nameSuffix
	upgradeCluster.AgentPoolsToUpgrade = uc.agentPoolsToUpgrade

	kubeConfig, err := engine.GenerateKubeConfig(uc.containerService.Properties, uc.location)
	if err != nil {
		log.Fatalf("Failed to generate kubeconfig: %v", err)
	}

	if err = upgradeCluster.UpgradeCluster(uc.client, kubeConfig, BuildTag); err != nil {
		log.Fatalf("Error upgrading cluster: %v\n", err)
	}

	// Save the new apimodel to reflect the cluster's state.
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

	return f.SaveFile(uc.deploymentDirectory, "apimodel.json", b)
}
