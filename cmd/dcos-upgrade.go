// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/operations/dcosupgrade"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	dcosUpgradeName             = "dcos-upgrade"
	dcosUpgradeShortDescription = "Upgrade an existing DC/OS cluster"
	dcosUpgradeLongDescription  = "Upgrade an existing DC/OS cluster"
)

type dcosUpgradeCmd struct {
	authArgs

	// user input
	resourceGroupName   string
	deploymentDirectory string
	upgradeVersion      string
	location            string
	sshPrivateKeyPath   string

	// derived
	containerService   *api.ContainerService
	apiVersion         string
	currentDcosVersion string
	client             armhelpers.AKSEngineClient
	locale             *gotext.Locale
	nameSuffix         string
	sshPrivateKey      []byte
}

func newDcosUpgradeCmd() *cobra.Command {
	uc := dcosUpgradeCmd{}

	dcosUpgradeCmd := &cobra.Command{
		Use:   dcosUpgradeName,
		Short: dcosUpgradeShortDescription,
		Long:  dcosUpgradeLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return uc.run(cmd, args)
		},
	}

	f := dcosUpgradeCmd.Flags()
	f.StringVarP(&uc.location, "location", "l", "", "location the cluster is deployed in (required)")
	f.StringVarP(&uc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed (required)")
	f.StringVar(&uc.deploymentDirectory, "deployment-dir", "", "the location of the output from `generate` (required)")
	f.StringVar(&uc.sshPrivateKeyPath, "ssh-private-key-path", "", "ssh private key path (default: <deployment-dir>/id_rsa)")
	f.StringVar(&uc.upgradeVersion, "upgrade-version", "", "desired DC/OS version (required)")
	addAuthFlags(&uc.authArgs, f)

	return dcosUpgradeCmd
}

func (uc *dcosUpgradeCmd) validate(cmd *cobra.Command) error {
	log.Infoln("validating...")

	var err error

	uc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	if len(uc.resourceGroupName) == 0 {
		cmd.Usage()
		return errors.New("--resource-group must be specified")
	}

	if len(uc.location) == 0 {
		cmd.Usage()
		return errors.New("--location must be specified")
	}
	uc.location = helpers.NormalizeAzureRegion(uc.location)

	if len(uc.upgradeVersion) == 0 {
		cmd.Usage()
		return errors.New("--upgrade-version must be specified")
	}

	if len(uc.deploymentDirectory) == 0 {
		cmd.Usage()
		return errors.New("--deployment-dir must be specified")
	}

	if len(uc.sshPrivateKeyPath) == 0 {
		uc.sshPrivateKeyPath = filepath.Join(uc.deploymentDirectory, "id_rsa")
	}
	if uc.sshPrivateKey, err = ioutil.ReadFile(uc.sshPrivateKeyPath); err != nil {
		cmd.Usage()
		return errors.Wrap(err, "ssh-private-key-path must be specified")
	}

	if err = uc.authArgs.validateAuthArgs(); err != nil {
		return err
	}
	return nil
}

func (uc *dcosUpgradeCmd) loadCluster(cmd *cobra.Command) error {
	var err error

	if uc.client, err = uc.authArgs.getClient(); err != nil {
		return errors.Wrap(err, "Failed to get client")
	}

	ctx := context.Background()
	_, err = uc.client.EnsureResourceGroup(ctx, uc.resourceGroupName, uc.location, nil)
	if err != nil {
		return errors.Wrap(err, "Error ensuring resource group")
	}

	// load apimodel from the deployment directory
	apiModelPath := path.Join(uc.deploymentDirectory, "apimodel.json")

	if _, err = os.Stat(apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api model does not exist (%s)", apiModelPath)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
	}
	uc.containerService, uc.apiVersion, err = apiloader.LoadContainerServiceFromFile(apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}
	uc.currentDcosVersion = uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion

	if uc.currentDcosVersion == uc.upgradeVersion {
		return errors.Errorf("already running DCOS %s", uc.upgradeVersion)
	}

	if len(uc.containerService.Location) == 0 {
		uc.containerService.Location = uc.location
	} else if uc.containerService.Location != uc.location {
		return errors.New("--location does not match api model location")
	}

	// get available upgrades for container service
	orchestratorInfo, err := api.GetOrchestratorVersionProfile(uc.containerService.Properties.OrchestratorProfile, false)
	if err != nil {
		return errors.Wrap(err, "error getting list of available upgrades")
	}
	// add the current version if upgrade has failed
	orchestratorInfo.Upgrades = append(orchestratorInfo.Upgrades, &api.OrchestratorProfile{
		OrchestratorType:    uc.containerService.Properties.OrchestratorProfile.OrchestratorType,
		OrchestratorVersion: uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion})

	// validate desired upgrade version and set goal state
	found := false
	for _, up := range orchestratorInfo.Upgrades {
		if up.OrchestratorVersion == uc.upgradeVersion {
			uc.containerService.Properties.OrchestratorProfile.OrchestratorVersion = uc.upgradeVersion
			found = true
			break
		}
	}
	if !found {
		return errors.Errorf("upgrade to DCOS %s is not supported", uc.upgradeVersion)
	}

	// Read name suffix to identify nodes in the resource group that belong
	// to this cluster.
	templatePath := path.Join(uc.deploymentDirectory, "azuredeploy.json")
	contents, _ := ioutil.ReadFile(templatePath)

	var template interface{}
	json.Unmarshal(contents, &template)

	templateMap := template.(map[string]interface{})
	templateParameters := templateMap["parameters"].(map[string]interface{})

	nameSuffixParam := templateParameters["nameSuffix"].(map[string]interface{})
	uc.nameSuffix = nameSuffixParam["defaultValue"].(string)
	log.Infof("Name suffix: %s", uc.nameSuffix)
	return nil
}

func (uc *dcosUpgradeCmd) run(cmd *cobra.Command, args []string) error {
	err := uc.validate(cmd)
	if err != nil {
		log.Fatalf("error validating upgrade command: %v", err)
	}

	err = uc.loadCluster(cmd)
	if err != nil {
		log.Fatalf("error loading existing cluster: %v", err)
	}

	upgradeCluster := dcosupgrade.UpgradeCluster{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
		Logger: log.NewEntry(log.New()),
		Client: uc.client,
	}

	if err = upgradeCluster.UpgradeCluster(uc.authArgs.SubscriptionID, uc.resourceGroupName, uc.currentDcosVersion,
		uc.containerService, uc.nameSuffix, uc.sshPrivateKey); err != nil {
		log.Fatalf("Error upgrading cluster: %v", err)
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

	return f.SaveFile(uc.deploymentDirectory, "apimodel.json", b)
}
