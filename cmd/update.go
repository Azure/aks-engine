// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	v1 "k8s.io/api/core/v1"
)

type updateCmd struct {
	authArgs

	// user input
	apiModelPath      string
	resourceGroupName string
	location          string
	agentPoolToUpdate string

	// derived
	containerService *api.ContainerService
	apiVersion       string
	agentPool        *api.AgentPoolProfile
	client           armhelpers.AKSEngineClient
	locale           *gotext.Locale
	nameSuffix       string
	agentPoolIndex   int
	logger           *log.Entry
	apiserverURL     string
	kubeconfig       string
	nodes            []v1.Node
}

const (
	updateName             = "update"
	updateShortDescription = "Update an existing AKS Engine-created VMSS node pool"
	updateLongDescription  = "Update an existing AKS Engine-created VMSS node pool in a Kubernetes cluster by updating its VMSS model"
)

// newUpdateCmd returns an instance reference of updateCmd
func newUpdateCmd() *cobra.Command {
	uc := updateCmd{}

	updateCmd := &cobra.Command{
		Use:   updateName,
		Short: updateShortDescription,
		Long:  updateLongDescription,
		RunE:  uc.run,
	}

	f := updateCmd.Flags()
	f.StringVarP(&uc.location, "location", "l", "", "location the cluster is deployed in")
	f.StringVarP(&uc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed")
	f.StringVarP(&uc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file")
	f.StringVar(&uc.agentPoolToUpdate, "node-pool", "", "node pool to scale")
	addAuthFlags(&uc.authArgs, f)

	return updateCmd
}

func (uc *updateCmd) validate(cmd *cobra.Command) error {
	log.Debugln("validating update command line arguments...")
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

	if uc.apiModelPath == "" {
		_ = cmd.Usage()
		return errors.New("--api-model must be specified")
	}

	if uc.agentPoolToUpdate == "" {
		_ = cmd.Usage()
		return errors.New("--node-pool must be specified")
	}

	return nil
}

func (uc *updateCmd) load() error {
	logger := log.New()
	logger.Formatter = new(prefixed.TextFormatter)
	uc.logger = log.NewEntry(log.New())
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	if _, err = os.Stat(uc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api model does not exist (%s)", uc.apiModelPath)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: uc.locale,
		},
	}
	uc.containerService, uc.apiVersion, err = apiloader.LoadContainerServiceFromFile(uc.apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	if uc.containerService.Properties.IsCustomCloudProfile() {
		if err = writeCustomCloudProfile(uc.containerService); err != nil {
			return errors.Wrap(err, "error writing custom cloud profile")
		}

		if err = uc.containerService.Properties.SetCustomCloudSpec(api.AzureCustomCloudSpecParams{IsUpgrade: false, IsScale: true}); err != nil {
			return errors.Wrap(err, "error parsing the api model")
		}
	}

	if err = uc.authArgs.validateAuthArgs(); err != nil {
		return err
	}

	if uc.client, err = uc.authArgs.getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	_, err = uc.client.EnsureResourceGroup(ctx, uc.resourceGroupName, uc.location, nil)
	if err != nil {
		return err
	}

	if uc.containerService.Location == "" {
		uc.containerService.Location = uc.location
	} else if uc.containerService.Location != uc.location {
		return errors.New("--location does not match api model location")
	}

	agentPoolIndex := -1
	for i, pool := range uc.containerService.Properties.AgentPoolProfiles {
		if pool.Name == uc.agentPoolToUpdate {
			agentPoolIndex = i
			uc.agentPool = pool
			uc.agentPoolIndex = i
		}
	}
	if agentPoolIndex == -1 {
		return errors.Errorf("node pool %s was not found in the deployed api model", uc.agentPoolToUpdate)
	}
	if uc.agentPool.AvailabilityProfile != api.VirtualMachineScaleSets {
		return errors.Errorf("aks-engine node pool update requires a VMSS node pool, %s is backed by a VM Availability Set", uc.agentPoolToUpdate)
	}

	//allows to identify VMs in the resource group that belong to this cluster.
	uc.nameSuffix = uc.containerService.Properties.GetClusterID()
	log.Debugf("Cluster ID used in all agent pools: %s", uc.nameSuffix)

	uc.kubeconfig, err = engine.GenerateKubeConfig(uc.containerService.Properties, uc.location)
	if err != nil {
		return errors.New("Unable to derive kubeconfig from api model")
	}
	return nil
}

func (uc *updateCmd) run(cmd *cobra.Command, args []string) error {
	if err := uc.validate(cmd); err != nil {
		return errors.Wrap(err, "failed to validate scale command")
	}
	if err := uc.load(); err != nil {
		return errors.Wrap(err, "failed to load existing container service")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	sc := scaleCmd{
		location:          uc.location,
		apiModelPath:      uc.apiModelPath,
		resourceGroupName: uc.resourceGroupName,
		agentPoolToScale:  uc.agentPoolToUpdate,
		validateCmd:       false,
		updateVMSSModel:   true,
		loadAPIModel:      false,
		persistAPIModel:   false,
	}
	sc.RawAzureEnvironment = uc.RawAzureEnvironment
	sc.rawSubscriptionID = uc.rawSubscriptionID
	sc.SubscriptionID = uc.SubscriptionID
	sc.AuthMethod = uc.AuthMethod
	sc.rawClientID = uc.rawClientID
	sc.ClientID = uc.ClientID
	sc.ClientSecret = uc.ClientSecret
	sc.CertificatePath = uc.CertificatePath
	sc.PrivateKeyPath = uc.PrivateKeyPath
	sc.IdentitySystem = uc.IdentitySystem
	sc.language = uc.language
	sc.logger = uc.logger
	sc.containerService = uc.containerService
	sc.apiVersion = uc.apiVersion
	sc.client = uc.client
	sc.nameSuffix = uc.nameSuffix
	sc.kubeconfig = uc.kubeconfig
	sc.agentPool = uc.agentPool
	sc.agentPoolIndex = uc.agentPoolIndex

	for vmssListPage, err := sc.client.ListVirtualMachineScaleSets(ctx, sc.resourceGroupName); vmssListPage.NotDone(); err = vmssListPage.NextWithContext(ctx) {
		if err != nil {
			return errors.Wrap(err, "failed to get VMSS list in the resource group")
		}
		for _, vmss := range vmssListPage.Values() {
			vmssName := *vmss.Name
			if sc.agentPool.OSType == api.Windows {
				possibleIndex, nameMungingErr := strconv.Atoi(vmssName[len(vmssName)-2:])
				if nameMungingErr != nil {
					continue
				}
				if !(sc.containerService.Properties.GetAgentVMPrefix(sc.agentPool, possibleIndex) == vmssName) {
					continue
				}
			} else {
				if !sc.vmInAgentPool(vmssName, vmss.Tags) {
					continue
				}
			}

			if vmss.Sku != nil {
				sc.newDesiredAgentCount = int(*vmss.Sku.Capacity)
				uc.agentPool.Count = sc.newDesiredAgentCount
			} else {
				return errors.Wrap(err, fmt.Sprintf("failed to detect find VMSS matching node pool %s in resource group %s", sc.agentPoolToScale, sc.resourceGroupName))
			}
		}
	}

	sc.run(cmd, args)

	return uc.saveAPIModel()
}

func (uc *updateCmd) saveAPIModel() error {
	var err error
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
