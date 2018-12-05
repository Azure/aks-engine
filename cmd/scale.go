// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/armhelpers/utils"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/operations"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type scaleCmd struct {
	authArgs

	// user input
	resourceGroupName    string
	deploymentDirectory  string
	newDesiredAgentCount int
	location             string
	agentPoolToScale     string
	masterFQDN           string

	// derived
	containerService *api.ContainerService
	apiVersion       string
	apiModelPath     string
	agentPool        *api.AgentPoolProfile
	client           armhelpers.AKSEngineClient
	locale           *gotext.Locale
	nameSuffix       string
	agentPoolIndex   int
	logger           *log.Entry
}

const (
	scaleName             = "scale"
	scaleShortDescription = "Scale an existing Kubernetes cluster"
	scaleLongDescription  = "Scale an existing Kubernetes cluster by specifying increasing or decreasing the node count of an agentpool"
	apiModelFilename      = "apimodel.json"
)

// NewScaleCmd run a command to upgrade a Kubernetes cluster
func newScaleCmd() *cobra.Command {
	sc := scaleCmd{}

	scaleCmd := &cobra.Command{
		Use:   scaleName,
		Short: scaleShortDescription,
		Long:  scaleLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sc.run(cmd, args)
		},
	}

	f := scaleCmd.Flags()
	f.StringVarP(&sc.location, "location", "l", "", "location the cluster is deployed in")
	f.StringVarP(&sc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed")
	f.StringVar(&sc.deploymentDirectory, "deployment-dir", "", "the location of the output from `generate`")
	f.IntVarP(&sc.newDesiredAgentCount, "new-node-count", "c", 0, "desired number of nodes")
	f.StringVar(&sc.agentPoolToScale, "node-pool", "", "node pool to scale")
	f.StringVar(&sc.masterFQDN, "master-FQDN", "", "FQDN for the master load balancer, Needed to scale down Kubernetes agent pools")

	addAuthFlags(&sc.authArgs, f)

	return scaleCmd
}

func (sc *scaleCmd) validate(cmd *cobra.Command) error {
	log.Infoln("validating...")
	var err error

	sc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	if sc.resourceGroupName == "" {
		cmd.Usage()
		return errors.New("--resource-group must be specified")
	}

	if sc.location == "" {
		cmd.Usage()
		return errors.New("--location must be specified")
	}

	sc.location = helpers.NormalizeAzureRegion(sc.location)

	if sc.newDesiredAgentCount == 0 {
		cmd.Usage()
		return errors.New("--new-node-count must be specified")
	}

	if sc.deploymentDirectory == "" {
		cmd.Usage()
		return errors.New("--deployment-dir must be specified")
	}

	return nil
}

func (sc *scaleCmd) load(cmd *cobra.Command) error {
	sc.logger = log.New().WithField("source", "scaling command line")
	var err error

	if err = sc.authArgs.validateAuthArgs(); err != nil {
		return err
	}

	if sc.client, err = sc.authArgs.getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	_, err = sc.client.EnsureResourceGroup(ctx, sc.resourceGroupName, sc.location, nil)
	if err != nil {
		return err
	}

	// load apimodel from the deployment directory
	sc.apiModelPath = path.Join(sc.deploymentDirectory, apiModelFilename)

	if _, err = os.Stat(sc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api model does not exist (%s)", sc.apiModelPath)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: sc.locale,
		},
	}
	sc.containerService, sc.apiVersion, err = apiloader.LoadContainerServiceFromFile(sc.apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	if sc.containerService.Location == "" {
		sc.containerService.Location = sc.location
	} else if sc.containerService.Location != sc.location {
		return errors.New("--location does not match api model location")
	}

	if sc.agentPoolToScale == "" {
		agentPoolCount := len(sc.containerService.Properties.AgentPoolProfiles)
		if agentPoolCount > 1 {
			return errors.New("--node-pool is required if more than one agent pool is defined in the container service")
		} else if agentPoolCount == 1 {
			sc.agentPool = sc.containerService.Properties.AgentPoolProfiles[0]
			sc.agentPoolIndex = 0
			sc.agentPoolToScale = sc.containerService.Properties.AgentPoolProfiles[0].Name
		} else {
			return errors.New("No node pools found to scale")
		}
	} else {
		agentPoolIndex := -1
		for i, pool := range sc.containerService.Properties.AgentPoolProfiles {
			if pool.Name == sc.agentPoolToScale {
				agentPoolIndex = i
				sc.agentPool = pool
				sc.agentPoolIndex = i
			}
		}
		if agentPoolIndex == -1 {
			return errors.Errorf("node pool %s was not found in the deployed api model", sc.agentPoolToScale)
		}
	}

	templatePath := path.Join(sc.deploymentDirectory, "azuredeploy.json")
	contents, _ := ioutil.ReadFile(templatePath)

	var template interface{}
	json.Unmarshal(contents, &template)

	templateMap := template.(map[string]interface{})
	templateParameters := templateMap["parameters"].(map[string]interface{})

	nameSuffixParam := templateParameters["nameSuffix"].(map[string]interface{})
	sc.nameSuffix = nameSuffixParam["defaultValue"].(string)
	log.Infof("Name suffix: %s", sc.nameSuffix)
	return nil
}

func (sc *scaleCmd) run(cmd *cobra.Command, args []string) error {
	if err := sc.validate(cmd); err != nil {
		return errors.Wrap(err, "failed to validate scale command")
	}
	if err := sc.load(cmd); err != nil {
		return errors.Wrap(err, "failed to load existing container service")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	orchestratorInfo := sc.containerService.Properties.OrchestratorProfile
	var currentNodeCount, highestUsedIndex, index, winPoolIndex int
	winPoolIndex = -1
	indexes := make([]int, 0)
	indexToVM := make(map[int]string)
	if sc.agentPool.IsAvailabilitySets() {
		for vmsListPage, err := sc.client.ListVirtualMachines(ctx, sc.resourceGroupName); vmsListPage.NotDone(); err = vmsListPage.Next() {
			if err != nil {
				return errors.Wrap(err, "failed to get vms in the resource group")
			} else if len(vmsListPage.Values()) < 1 {
				return errors.New("The provided resource group does not contain any vms")
			}
			for _, vm := range vmsListPage.Values() {
				vmName := *vm.Name
				if !sc.vmInAgentPool(vmName, vm.Tags) {
					continue
				}

				osPublisher := vm.StorageProfile.ImageReference.Publisher
				if osPublisher != nil && strings.EqualFold(*osPublisher, "MicrosoftWindowsServer") {
					_, _, winPoolIndex, index, err = utils.WindowsVMNameParts(vmName)
				} else {
					_, _, index, err = utils.K8sLinuxVMNameParts(vmName)
				}
				if err != nil {
					return err
				}

				indexToVM[index] = vmName
				indexes = append(indexes, index)
			}
		}
		sortedIndexes := sort.IntSlice(indexes)
		sortedIndexes.Sort()
		indexes = []int(sortedIndexes)
		currentNodeCount = len(indexes)

		if currentNodeCount == sc.newDesiredAgentCount {
			log.Info("Cluster is currently at the desired agent count.")
			return nil
		}
		highestUsedIndex = indexes[len(indexes)-1]

		// Scale down Scenario
		if currentNodeCount > sc.newDesiredAgentCount {
			if sc.masterFQDN == "" {
				cmd.Usage()
				return errors.New("master-FQDN is required to scale down a kubernetes cluster's agent pool")
			}

			vmsToDelete := make([]string, 0)
			for i := currentNodeCount - 1; i >= sc.newDesiredAgentCount; i-- {
				index = indexes[i]
				vmsToDelete = append(vmsToDelete, indexToVM[index])
			}

			switch orchestratorInfo.OrchestratorType {
			case api.Kubernetes:
				kubeConfig, err := engine.GenerateKubeConfig(sc.containerService.Properties, sc.location)
				if err != nil {
					return errors.Wrap(err, "failed to generate kube config")
				}
				err = sc.drainNodes(kubeConfig, vmsToDelete)
				if err != nil {
					return errors.Wrap(err, "Got error while draining the nodes to be deleted")
				}
			}

			errList := operations.ScaleDownVMs(sc.client, sc.logger, sc.SubscriptionID.String(), sc.resourceGroupName, vmsToDelete...)
			if errList != nil {
				var err error
				format := "Node '%s' failed to delete with error: '%s'"
				for element := errList.Front(); element != nil; element = element.Next() {
					vmError, ok := element.Value.(*operations.VMScalingErrorDetails)
					if ok {
						if err == nil {
							err = errors.Errorf(format, vmError.Name, vmError.Error.Error())
						} else {
							err = errors.Wrapf(err, format, vmError.Name, vmError.Error.Error())
						}
					}
				}
				return err
			}

			return sc.saveAPIModel()
		}
	} else {
		for vmssListPage, err := sc.client.ListVirtualMachineScaleSets(ctx, sc.resourceGroupName); vmssListPage.NotDone(); vmssListPage.Next() {
			if err != nil {
				return errors.Wrap(err, "failed to get vmss list in the resource group")
			}
			for _, vmss := range vmssListPage.Values() {
				vmName := *vmss.Name
				if !sc.vmInAgentPool(vmName, vmss.Tags) {
					continue
				}

				osPublisher := vmss.VirtualMachineProfile.StorageProfile.ImageReference.Publisher
				if osPublisher != nil && strings.EqualFold(*osPublisher, "MicrosoftWindowsServer") {
					_, _, winPoolIndex, _, err = utils.WindowsVMNameParts(vmName)
					log.Errorln(err)
				}

				currentNodeCount = int(*vmss.Sku.Capacity)
				highestUsedIndex = 0
			}
		}
	}

	translator := engine.Context{
		Translator: &i18n.Translator{
			Locale: sc.locale,
		},
	}
	templateGenerator, err := engine.InitializeTemplateGenerator(translator)
	if err != nil {
		return errors.Wrap(err, "failed to initialize template generator")
	}

	sc.containerService.Properties.AgentPoolProfiles = []*api.AgentPoolProfile{sc.agentPool}

	_, err = sc.containerService.SetPropertiesDefaults(false, true)
	if err != nil {
		log.Fatalf("error in SetPropertiesDefaults template %s: %s", sc.apiModelPath, err.Error())
		os.Exit(1)
	}
	template, parameters, err := templateGenerator.GenerateTemplate(sc.containerService, engine.DefaultGeneratorCode, BuildTag)
	if err != nil {
		return errors.Wrapf(err, "error generating template %s", sc.apiModelPath)
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
		return errors.Wrap(err, "errror unmarshalling parameters")
	}

	transformer := transform.Transformer{Translator: translator.Translator}
	// Our templates generate a range of nodes based on a count and offset, it is possible for there to be holes in the template
	// So we need to set the count in the template to get enough nodes for the range, if there are holes that number will be larger than the desired count
	countForTemplate := sc.newDesiredAgentCount
	if highestUsedIndex != 0 {
		countForTemplate += highestUsedIndex + 1 - currentNodeCount
	}
	addValue(parametersJSON, sc.agentPool.Name+"Count", countForTemplate)

	if winPoolIndex != -1 {
		templateJSON["variables"].(map[string]interface{})[sc.agentPool.Name+"Index"] = winPoolIndex
	}
	switch orchestratorInfo.OrchestratorType {
	case api.Kubernetes:
		err = transformer.NormalizeForK8sVMASScalingUp(sc.logger, templateJSON)
		if err != nil {
			return errors.Wrapf(err, "error tranforming the template for scaling template %s", sc.apiModelPath)
		}
		if sc.agentPool.IsAvailabilitySets() {
			addValue(parametersJSON, fmt.Sprintf("%sOffset", sc.agentPool.Name), highestUsedIndex+1)
		}
	case api.Swarm:
	case api.SwarmMode:
	case api.DCOS:
		if sc.agentPool.IsAvailabilitySets() {
			return errors.Errorf("scaling isn't supported for orchestrator %q, with availability sets", orchestratorInfo.OrchestratorType)
		}
		transformer.NormalizeForVMSSScaling(sc.logger, templateJSON)
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	deploymentSuffix := random.Int31()

	_, err = sc.client.DeployTemplate(
		ctx,
		sc.resourceGroupName,
		fmt.Sprintf("%s-%d", sc.resourceGroupName, deploymentSuffix),
		templateJSON,
		parametersJSON)
	if err != nil {
		return err
	}

	return sc.saveAPIModel()
}

func (sc *scaleCmd) saveAPIModel() error {
	var err error
	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: sc.locale,
		},
	}
	var apiVersion string
	sc.containerService, apiVersion, err = apiloader.LoadContainerServiceFromFile(sc.apiModelPath, false, true, nil)
	if err != nil {
		return err
	}
	sc.containerService.Properties.AgentPoolProfiles[sc.agentPoolIndex].Count = sc.newDesiredAgentCount

	b, err := apiloader.SerializeContainerService(sc.containerService, apiVersion)

	if err != nil {
		return err
	}

	f := helpers.FileSaver{
		Translator: &i18n.Translator{
			Locale: sc.locale,
		},
	}

	return f.SaveFile(sc.deploymentDirectory, apiModelFilename, b)
}

func (sc *scaleCmd) vmInAgentPool(vmName string, tags map[string]*string) bool {
	// Try to locate the VM's agent pool by expected tags.
	if tags != nil {
		if poolName, ok := tags["poolName"]; ok {
			if nameSuffix, ok := tags["resourceNameSuffix"]; ok {
				// Use strings.Contains for the nameSuffix as the Windows Agent Pools use only
				// a substring of the first 5 characters of the entire nameSuffix.
				if strings.EqualFold(*poolName, sc.agentPoolToScale) && strings.Contains(sc.nameSuffix, *nameSuffix) {
					return true
				}
			}
		}
	}

	// Fall back to checking the VM name to see if it fits the naming pattern.
	return strings.Contains(vmName, sc.nameSuffix[:5]) && strings.Contains(vmName, sc.agentPoolToScale)
}

type paramsMap map[string]interface{}

func addValue(m paramsMap, k string, v interface{}) {
	m[k] = paramsMap{
		"value": v,
	}
}

func (sc *scaleCmd) drainNodes(kubeConfig string, vmsToDelete []string) error {
	masterURL := sc.masterFQDN
	if !strings.HasPrefix(masterURL, "https://") {
		masterURL = fmt.Sprintf("https://%s", masterURL)
	}
	numVmsToDrain := len(vmsToDelete)
	errChan := make(chan *operations.VMScalingErrorDetails, numVmsToDrain)
	defer close(errChan)
	for _, vmName := range vmsToDelete {
		go func(vmName string) {
			err := operations.SafelyDrainNode(sc.client, sc.logger,
				masterURL, kubeConfig, vmName, time.Duration(60)*time.Minute)
			if err != nil {
				log.Errorf("Failed to drain node %s, got error %v", vmName, err)
				errChan <- &operations.VMScalingErrorDetails{Error: err, Name: vmName}
				return
			}
			errChan <- nil
		}(vmName)
	}

	for i := 0; i < numVmsToDrain; i++ {
		errDetails := <-errChan
		if errDetails != nil {
			return errors.Wrapf(errDetails.Error, "Node %q failed to drain with error", errDetails.Name)
		}
	}

	return nil
}
