// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/operations"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	v1 "k8s.io/api/core/v1"
)

type addPoolCmd struct {
	authArgs

	// user input
	apiModelPath      string
	resourceGroupName string
	nodePoolPath      string
	location          string

	// derived
	containerService *api.ContainerService
	apiVersion       string
	nodePool         *api.AgentPoolProfile
	client           armhelpers.AKSEngineClient
	locale           *gotext.Locale
	nameSuffix       string
	logger           *log.Entry
	apiserverURL     string
	kubeconfig       string
	nodes            []v1.Node
}

const (
	addPoolName             = "addpool"
	addPoolShortDescription = "Add a node pool to an existing AKS Engine-created Kubernetes cluster"
	addPoolLongDescription  = "Add a node pool to an existing AKS Engine-created Kubernetes cluster by referencing a new agentpoolProfile spec"
)

// newAddPoolCmd run a command to add an agent pool to a Kubernetes cluster
func newAddPoolCmd() *cobra.Command {
	apc := addPoolCmd{}

	addPoolCmd := &cobra.Command{
		Use:   addPoolName,
		Short: addPoolShortDescription,
		Long:  addPoolLongDescription,
		RunE:  apc.run,
	}

	f := addPoolCmd.Flags()
	f.StringVarP(&apc.location, "location", "l", "", "location the cluster is deployed in")
	f.StringVarP(&apc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed")
	f.StringVarP(&apc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file")
	f.StringVarP(&apc.nodePoolPath, "node-pool", "p", "", "path to the generated nodepool.json file")

	addAuthFlags(&apc.authArgs, f)

	return addPoolCmd
}

func (apc *addPoolCmd) validate(cmd *cobra.Command) error {
	log.Debugln("validating addpool command line arguments...")
	var err error

	apc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	if apc.resourceGroupName == "" {
		_ = cmd.Usage()
		return errors.New("--resource-group must be specified")
	}

	if apc.location == "" {
		_ = cmd.Usage()
		return errors.New("--location must be specified")
	}

	apc.location = helpers.NormalizeAzureRegion(apc.location)

	if apc.apiModelPath == "" {
		_ = cmd.Usage()
		return errors.New("--api-model must be specified")
	}

	if apc.nodePoolPath == "" {
		_ = cmd.Usage()
		return errors.New("--nodepool must be specified")
	}
	return nil
}

func (apc *addPoolCmd) load() error {
	logger := log.New()
	logger.Formatter = new(prefixed.TextFormatter)
	apc.logger = log.NewEntry(log.New())
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	if _, err = os.Stat(apc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api model does not exist (%s)", apc.apiModelPath)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: apc.locale,
		},
	}
	apc.containerService, apc.apiVersion, err = apiloader.LoadContainerServiceFromFile(apc.apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	if _, err = os.Stat(apc.nodePoolPath); os.IsNotExist(err) {
		return errors.Errorf("specified agent pool spec does not exist (%s)", apc.nodePoolPath)
	}

	apc.nodePool, err = apiloader.LoadAgentpoolProfileFromFile(apc.nodePoolPath)
	if err != nil {
		return errors.Wrap(err, "error parsing the agent pool")
	}

	if apc.containerService.Properties.IsCustomCloudProfile() {
		if err = writeCustomCloudProfile(apc.containerService); err != nil {
			return errors.Wrap(err, "error writing custom cloud profile")
		}
		if err = apc.containerService.Properties.SetCustomCloudSpec(api.AzureCustomCloudSpecParams{IsUpgrade: false, IsScale: true}); err != nil {
			return errors.Wrap(err, "error parsing the api model")
		}
	}

	for _, p := range apc.containerService.Properties.AgentPoolProfiles {
		if strings.EqualFold(p.Name, apc.nodePool.Name) {
			return errors.Errorf("node pool %s already exists", p.Name)
		}
		if !strings.EqualFold(p.AvailabilityProfile, apc.nodePool.AvailabilityProfile) {
			return errors.New("mixed mode availability profiles are not allowed, all node pools should have the same availabilityProfile")
		}
	}

	if err = apc.authArgs.validateAuthArgs(); err != nil {
		return err
	}

	if apc.client, err = apc.authArgs.getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	_, err = apc.client.EnsureResourceGroup(ctx, apc.resourceGroupName, apc.location, nil)
	if err != nil {
		return err
	}

	if apc.containerService.Location == "" {
		apc.containerService.Location = apc.location
	} else if apc.containerService.Location != apc.location {
		return errors.New("--location does not match api model location")
	}

	//allows to identify VMs in the resource group that belong to this cluster.
	apc.nameSuffix = apc.containerService.Properties.GetClusterID()
	log.Debugf("Cluster ID used in all agent pools: %s", apc.nameSuffix)

	apc.kubeconfig, err = engine.GenerateKubeConfig(apc.containerService.Properties, apc.location)
	if err != nil {
		return errors.New("Unable to derive kubeconfig from api model")
	}
	return nil
}

func (apc *addPoolCmd) run(cmd *cobra.Command, args []string) error {
	if err := apc.validate(cmd); err != nil {
		return errors.Wrap(err, "failed to validate addpool command")
	}
	if err := apc.load(); err != nil {
		return errors.Wrap(err, "failed to load existing container service")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	orchestratorInfo := apc.containerService.Properties.OrchestratorProfile
	winPoolIndex := -1

	if apc.nodePool.IsVirtualMachineScaleSets() {
		for vmssListPage, err := apc.client.ListVirtualMachineScaleSets(ctx, apc.resourceGroupName); vmssListPage.NotDone(); err = vmssListPage.NextWithContext(ctx) {
			if err != nil {
				return errors.Wrap(err, "failed to get VMSS list in the resource group")
			}
			for _, vmss := range vmssListPage.Values() {
				segments := strings.Split(*vmss.Name, "-")
				if len(segments) == 4 && segments[0] == "k8s" {
					vmssName := segments[1]
					if apc.nodePool.Name == vmssName {
						return errors.New("An agent pool with the given name already exists in the cluster")
					}
				}
			}
		}
	}
	translator := engine.Context{
		Translator: &i18n.Translator{
			Locale: apc.locale,
		},
	}
	templateGenerator, err := engine.InitializeTemplateGenerator(translator)
	if err != nil {
		return errors.Wrap(err, "failed to initialize template generator")
	}

	apc.containerService.Properties.AgentPoolProfiles = []*api.AgentPoolProfile{apc.nodePool}

	_, err = apc.containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		return errors.Wrapf(err, "error in SetPropertiesDefaults template %s", apc.apiModelPath)
	}
	template, parameters, err := templateGenerator.GenerateTemplateV2(apc.containerService, engine.DefaultGeneratorCode, BuildTag)
	if err != nil {
		return errors.Wrapf(err, "error generating template %s", apc.apiModelPath)
	}

	if template, err = transform.PrettyPrintArmTemplate(template); err != nil {
		return errors.Wrap(err, "error pretty printing template")
	}

	templateJSON := make(map[string]interface{})
	parametersJSON := make(map[string]interface{})

	err = json.Unmarshal([]byte(template), &templateJSON)
	if err != nil {
		return errors.Wrap(err, "error unmarshaling template")
	}

	err = json.Unmarshal([]byte(parameters), &parametersJSON)
	if err != nil {
		return errors.Wrap(err, "error unmarshaling parameters")
	}

	// The agent pool is set to index 0 for the scale operation, we need to overwrite the template variables that rely on pool index.
	if winPoolIndex != -1 {
		templateJSON["variables"].(map[string]interface{})[apc.nodePool.Name+"Index"] = winPoolIndex
		templateJSON["variables"].(map[string]interface{})[apc.nodePool.Name+"VMNamePrefix"] = apc.containerService.Properties.GetAgentVMPrefix(apc.nodePool, winPoolIndex)
	}
	if orchestratorInfo.OrchestratorType == api.Kubernetes {
		transformer := transform.Transformer{Translator: translator.Translator}

		if orchestratorInfo.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku {
			err = transformer.NormalizeForK8sSLBScalingOrUpgrade(apc.logger, templateJSON)
			if err != nil {
				return errors.Wrapf(err, "error transforming the template for scaling with SLB %s", apc.apiModelPath)
			}
		}

		if apc.nodePool.IsVirtualMachineScaleSets() {
			err = transformer.NormalizeForK8sVMASScalingUp(apc.logger, templateJSON)
			if err != nil {
				return errors.Wrapf(err, "error transforming the template for scaling template %s", apc.apiModelPath)
			}
			addValue(parametersJSON, apc.nodePool.Name+"Count", 0)
		} else {
			err = transformer.NormalizeForK8sAddVMASPool(apc.logger, templateJSON)
			if err != nil {
				return errors.Wrap(err, "error transforming the template to add a VMAS node pool")
			}
		}
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	deploymentSuffix := random.Int31()

	_, err = apc.client.DeployTemplate(
		ctx,
		apc.resourceGroupName,
		fmt.Sprintf("%s-%d", apc.resourceGroupName, deploymentSuffix),
		templateJSON,
		parametersJSON)
	if err != nil {
		return err
	}
	if apc.nodes != nil {
		nodes, err := operations.GetNodes(apc.client, apc.logger, apc.apiserverURL, apc.kubeconfig, time.Duration(5)*time.Minute, apc.nodePool.Name, apc.nodePool.Count)
		if err == nil && nodes != nil {
			apc.nodes = nodes
			apc.logger.Infof("Nodes in pool '%s' after scaling:\n", apc.nodePool.Name)
			operations.PrintNodes(apc.nodes)
		} else {
			apc.logger.Warningf("Unable to get nodes in pool %s after scaling:\n", apc.nodePool.Name)
		}
	}

	return apc.saveAPIModel()
}

func (apc *addPoolCmd) saveAPIModel() error {
	var err error
	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: apc.locale,
		},
	}
	var apiVersion string
	apc.containerService, apiVersion, err = apiloader.LoadContainerServiceFromFile(apc.apiModelPath, false, true, nil)
	if err != nil {
		return err
	}

	apc.containerService.Properties.AgentPoolProfiles = append(apc.containerService.Properties.AgentPoolProfiles, apc.nodePool)

	b, err := apiloader.SerializeContainerService(apc.containerService, apiVersion)

	if err != nil {
		return err
	}

	f := helpers.FileSaver{
		Translator: &i18n.Translator{
			Locale: apc.locale,
		},
	}
	dir, file := filepath.Split(apc.apiModelPath)
	return f.SaveFile(dir, file, b)
}
