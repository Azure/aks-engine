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
	"sort"
	"strconv"
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
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	v1 "k8s.io/api/core/v1"
)

type scaleCmd struct {
	authArgs

	// user input
	apiModelPath         string
	resourceGroupName    string
	newDesiredAgentCount int
	deploymentDirectory  string
	location             string
	agentPoolToScale     string
	masterFQDN           string

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
	scaleName             = "scale"
	scaleShortDescription = "Scale an existing AKS Engine-created Kubernetes cluster"
	scaleLongDescription  = "Scale an existing AKS Engine-created Kubernetes cluster by specifying increasing or decreasing the number of nodes in a node pool"
	apiModelFilename      = "apimodel.json"
)

// NewScaleCmd run a command to upgrade a Kubernetes cluster
func newScaleCmd() *cobra.Command {
	sc := scaleCmd{}

	scaleCmd := &cobra.Command{
		Use:   scaleName,
		Short: scaleShortDescription,
		Long:  scaleLongDescription,
		RunE:  sc.run,
	}

	f := scaleCmd.Flags()
	f.StringVarP(&sc.location, "location", "l", "", "location the cluster is deployed in")
	f.StringVarP(&sc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed")
	f.StringVarP(&sc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file")
	f.StringVar(&sc.deploymentDirectory, "deployment-dir", "", "the location of the output from `generate`")
	f.IntVarP(&sc.newDesiredAgentCount, "new-node-count", "c", 0, "desired number of nodes")
	f.StringVar(&sc.agentPoolToScale, "node-pool", "", "node pool to scale")
	f.StringVar(&sc.masterFQDN, "master-FQDN", "", "FQDN for the master load balancer that maps to the apiserver endpoint")
	f.StringVar(&sc.masterFQDN, "apiserver", "", "apiserver endpoint (required to cordon and drain nodes)")

	_ = f.MarkDeprecated("deployment-dir", "--deployment-dir is no longer required for scale or upgrade. Please use --api-model.")
	_ = f.MarkDeprecated("master-FQDN", "--apiserver is preferred")

	addAuthFlags(&sc.authArgs, f)

	return scaleCmd
}

func (sc *scaleCmd) validate(cmd *cobra.Command) error {
	log.Debugln("validating scale command line arguments...")
	var err error

	sc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	if sc.resourceGroupName == "" {
		_ = cmd.Usage()
		return errors.New("--resource-group must be specified")
	}

	if sc.location == "" {
		_ = cmd.Usage()
		return errors.New("--location must be specified")
	}

	sc.location = helpers.NormalizeAzureRegion(sc.location)

	if sc.newDesiredAgentCount == 0 {
		_ = cmd.Usage()
		return errors.New("--new-node-count must be specified")
	}

	if sc.apiModelPath == "" && sc.deploymentDirectory == "" {
		_ = cmd.Usage()
		return errors.New("--api-model must be specified")
	}

	if sc.apiModelPath != "" && sc.deploymentDirectory != "" {
		_ = cmd.Usage()
		return errors.New("ambiguous, please specify only one of --api-model and --deployment-dir")
	}

	return nil
}

func (sc *scaleCmd) load() error {
	logger := log.New()
	logger.Formatter = new(prefixed.TextFormatter)
	sc.logger = log.NewEntry(log.New())
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	if sc.apiModelPath == "" {
		sc.apiModelPath = filepath.Join(sc.deploymentDirectory, apiModelFilename)
	}

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

	if sc.containerService.Properties.IsCustomCloudProfile() {
		if err = writeCustomCloudProfile(sc.containerService); err != nil {
			return errors.Wrap(err, "error writing custom cloud profile")
		}

		if err = sc.containerService.Properties.SetCustomCloudSpec(api.AzureCustomCloudSpecParams{IsUpgrade: false, IsScale: true}); err != nil {
			return errors.Wrap(err, "error parsing the api model")
		}
	}

	if err = sc.authArgs.validateAuthArgs(); err != nil {
		return err
	}

	if sc.client, err = sc.authArgs.getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	_, err = sc.client.EnsureResourceGroup(ctx, sc.resourceGroupName, sc.location, nil)
	if err != nil {
		return err
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

	//allows to identify VMs in the resource group that belong to this cluster.
	sc.nameSuffix = sc.containerService.Properties.GetClusterID()
	log.Debugf("Cluster ID used in all agent pools: %s", sc.nameSuffix)

	if sc.masterFQDN != "" {
		if strings.HasPrefix(sc.masterFQDN, "https://") {
			sc.apiserverURL = sc.masterFQDN
		} else if strings.HasPrefix(sc.masterFQDN, "http://") {
			return errors.New("apiserver URL cannot be insecure http://")
		} else {
			sc.apiserverURL = fmt.Sprintf("https://%s", sc.masterFQDN)
		}
	}

	sc.kubeconfig, err = engine.GenerateKubeConfig(sc.containerService.Properties, sc.location)
	if err != nil {
		return errors.New("Unable to derive kubeconfig from api model")
	}
	return nil
}

func (sc *scaleCmd) run(cmd *cobra.Command, args []string) error {
	if err := sc.validate(cmd); err != nil {
		return errors.Wrap(err, "failed to validate scale command")
	}
	if err := sc.load(); err != nil {
		return errors.Wrap(err, "failed to load existing container service")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	orchestratorInfo := sc.containerService.Properties.OrchestratorProfile
	var currentNodeCount, highestUsedIndex, index, winPoolIndex int
	winPoolIndex = -1
	indexes := make([]int, 0)
	indexToVM := make(map[int]string)

	// Get nodes list from the k8s API before scaling for the desired pool
	if sc.apiserverURL != "" && orchestratorInfo.OrchestratorType == api.Kubernetes {
		nodes, err := operations.GetNodes(sc.client, sc.logger, sc.apiserverURL, sc.kubeconfig, time.Duration(5)*time.Minute, sc.agentPoolToScale, -1)
		if err == nil && nodes != nil {
			sc.nodes = nodes
		}
	}

	if sc.agentPool.IsAvailabilitySets() {
		for vmsListPage, err := sc.client.ListVirtualMachines(ctx, sc.resourceGroupName); vmsListPage.NotDone(); err = vmsListPage.Next() {
			if err != nil {
				return errors.Wrap(err, "failed to get VMs in the resource group")
			} else if len(vmsListPage.Values()) < 1 {
				return errors.New("The provided resource group does not contain any VMs")
			}
			for _, vm := range vmsListPage.Values() {
				vmName := *vm.Name
				if !sc.vmInAgentPool(vmName, vm.Tags) {
					continue
				}

				if sc.agentPool.OSType == api.Windows {
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
		indexes = sortedIndexes
		currentNodeCount = len(indexes)

		if currentNodeCount == sc.newDesiredAgentCount {
			sc.printScaleTargetEqualsExisting(currentNodeCount)
			return nil
		}
		highestUsedIndex = indexes[len(indexes)-1]

		// VMAS Scale down Scenario
		if currentNodeCount > sc.newDesiredAgentCount {
			if sc.apiserverURL == "" {
				_ = cmd.Usage()
				return errors.New("--apiserver is required to scale down a kubernetes cluster's agent pool")
			}

			if sc.nodes != nil {
				if len(sc.nodes) == 1 {
					sc.logger.Infof("There is %d node in pool %s before scaling down to %d:\n", len(sc.nodes), sc.agentPoolToScale, sc.newDesiredAgentCount)
				} else {
					sc.logger.Infof("There are %d nodes in pool %s before scaling down to %d:\n", len(sc.nodes), sc.agentPoolToScale, sc.newDesiredAgentCount)
				}
				operations.PrintNodes(sc.nodes)
				numNodesFromK8sAPI := len(sc.nodes)
				if currentNodeCount != numNodesFromK8sAPI {
					sc.logger.Warnf("There are %d VMs named \"*%s*\" in the resource group %s, but there are %d nodes named \"*%s*\" in the Kubernetes cluster\n", currentNodeCount, sc.agentPoolToScale, sc.resourceGroupName, numNodesFromK8sAPI, sc.agentPoolToScale)
				} else {
					nodesToDelete := currentNodeCount - sc.newDesiredAgentCount
					if nodesToDelete > 1 {
						sc.logger.Infof("%d nodes will be deleted\n", nodesToDelete)
					} else {
						sc.logger.Infof("%d node will be deleted\n", nodesToDelete)
					}
				}
			}

			vmsToDelete := make([]string, 0)
			for i := currentNodeCount - 1; i >= sc.newDesiredAgentCount; i-- {
				index = indexes[i]
				vmsToDelete = append(vmsToDelete, indexToVM[index])
			}

			for _, node := range vmsToDelete {
				sc.logger.Infof("Node %s will be cordoned and drained\n", node)
			}
			if orchestratorInfo.OrchestratorType == api.Kubernetes {
				err := sc.drainNodes(vmsToDelete)
				if err != nil {
					return errors.Wrap(err, "Got error while draining the nodes to be deleted")
				}
			}

			for _, node := range vmsToDelete {
				sc.logger.Infof("Node %s's VM will be deleted\n", node)
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
			if sc.nodes != nil {
				nodes, err := operations.GetNodes(sc.client, sc.logger, sc.apiserverURL, sc.kubeconfig, time.Duration(5)*time.Minute, sc.agentPoolToScale, sc.newDesiredAgentCount)
				if err == nil && nodes != nil {
					sc.nodes = nodes
					sc.logger.Infof("Nodes in pool %s after scaling:\n", sc.agentPoolToScale)
					operations.PrintNodes(sc.nodes)
				} else {
					sc.logger.Warningf("Unable to get nodes in pool %s after scaling:\n", sc.agentPoolToScale)
				}
			}

			return sc.saveAPIModel()
		}
	} else {
		for vmssListPage, err := sc.client.ListVirtualMachineScaleSets(ctx, sc.resourceGroupName); vmssListPage.NotDone(); err = vmssListPage.NextWithContext(ctx) {
			if err != nil {
				return errors.Wrap(err, "failed to get VMSS list in the resource group")
			}
			for _, vmss := range vmssListPage.Values() {
				vmssName := *vmss.Name
				if !sc.vmInAgentPool(vmssName, vmss.Tags) {
					continue
				}

				if vmss.Sku != nil {
					currentNodeCount = int(*vmss.Sku.Capacity)
					if int(*vmss.Sku.Capacity) == sc.newDesiredAgentCount {
						sc.printScaleTargetEqualsExisting(currentNodeCount)
						return nil
					} else if int(*vmss.Sku.Capacity) > sc.newDesiredAgentCount {
						log.Warnf("VMSS scale down is an alpha feature: VMSS VM nodes will not be cordoned and drained before scaling down!")
					}
				}

				if sc.agentPool.OSType == api.Windows {
					winPoolIndexStr := vmssName[len(vmssName)-2:]
					var err error
					winPoolIndex, err = strconv.Atoi(winPoolIndexStr)
					if err != nil {
						return errors.Wrap(err, "failed to get Windows pool index from VMSS name")
					}
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

	// Our templates generate a range of nodes based on a count and offset, it is possible for there to be holes in the template
	// So we need to set the count in the template to get enough nodes for the range, if there are holes that number will be larger than the desired count
	countForTemplate := sc.newDesiredAgentCount
	if highestUsedIndex != 0 {
		countForTemplate += highestUsedIndex + 1 - currentNodeCount
	}
	sc.agentPool.Count = countForTemplate
	sc.containerService.Properties.AgentPoolProfiles = []*api.AgentPoolProfile{sc.agentPool}

	_, err = sc.containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		return errors.Wrapf(err, "error in SetPropertiesDefaults template %s", sc.apiModelPath)
	}
	template, parameters, err := templateGenerator.GenerateTemplateV2(sc.containerService, engine.DefaultGeneratorCode, BuildTag)
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
		return errors.Wrap(err, "error unmarshaling parameters")
	}

	transformer := transform.Transformer{Translator: translator.Translator}

	addValue(parametersJSON, sc.agentPool.Name+"Count", countForTemplate)

	// The agent pool is set to index 0 for the scale operation, we need to overwrite the template variables that rely on pool index.
	if winPoolIndex != -1 {
		templateJSON["variables"].(map[string]interface{})[sc.agentPool.Name+"Index"] = winPoolIndex
		templateJSON["variables"].(map[string]interface{})[sc.agentPool.Name+"VMNamePrefix"] = sc.containerService.Properties.GetAgentVMPrefix(sc.agentPool, winPoolIndex)
	}
	if orchestratorInfo.OrchestratorType == api.Kubernetes {
		if orchestratorInfo.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku {
			err = transformer.NormalizeForK8sSLBScalingOrUpgrade(sc.logger, templateJSON)
			if err != nil {
				return errors.Wrapf(err, "error transforming the template for scaling with SLB %s", sc.apiModelPath)
			}
		}
		err = transformer.NormalizeForK8sVMASScalingUp(sc.logger, templateJSON)
		if err != nil {
			return errors.Wrapf(err, "error transforming the template for scaling template %s", sc.apiModelPath)
		}

		transformer.RemoveImmutableResourceProperties(sc.logger, templateJSON)

		if sc.agentPool.IsAvailabilitySets() {
			addValue(parametersJSON, fmt.Sprintf("%sOffset", sc.agentPool.Name), highestUsedIndex+1)
		}
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	deploymentSuffix := random.Int31()

	if sc.nodes != nil {
		sc.logger.Infof("Nodes in pool '%s' before scaling:\n", sc.agentPoolToScale)
		operations.PrintNodes(sc.nodes)
	}
	_, err = sc.client.DeployTemplate(
		ctx,
		sc.resourceGroupName,
		fmt.Sprintf("%s-%d", sc.resourceGroupName, deploymentSuffix),
		templateJSON,
		parametersJSON)
	if err != nil {
		return err
	}
	if sc.nodes != nil {
		nodes, err := operations.GetNodes(sc.client, sc.logger, sc.apiserverURL, sc.kubeconfig, time.Duration(5)*time.Minute, sc.agentPoolToScale, sc.newDesiredAgentCount)
		if err == nil && nodes != nil {
			sc.nodes = nodes
			sc.logger.Infof("Nodes in pool '%s' after scaling:\n", sc.agentPoolToScale)
			operations.PrintNodes(sc.nodes)
		} else {
			sc.logger.Warningf("Unable to get nodes in pool %s after scaling:\n", sc.agentPoolToScale)
		}
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
	dir, file := filepath.Split(sc.apiModelPath)
	return f.SaveFile(dir, file, b)
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

func (sc *scaleCmd) drainNodes(vmsToDelete []string) error {
	numVmsToDrain := len(vmsToDelete)
	errChan := make(chan *operations.VMScalingErrorDetails, numVmsToDrain)
	defer close(errChan)
	for _, vmName := range vmsToDelete {
		go func(vmName string) {
			err := operations.SafelyDrainNode(sc.client, sc.logger,
				sc.apiserverURL, sc.kubeconfig, vmName, time.Duration(60)*time.Minute)
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

func (sc *scaleCmd) printScaleTargetEqualsExisting(currentNodeCount int) {
	var printNodes bool
	trailingChar := "."
	if sc.nodes != nil {
		printNodes = true
		trailingChar = ":"
	}
	log.Infof("Node pool %s is already at the desired count %d%s", sc.agentPoolToScale, sc.newDesiredAgentCount, trailingChar)
	if printNodes {
		operations.PrintNodes(sc.nodes)
	}
	numNodesFromK8sAPI := len(sc.nodes)
	if currentNodeCount != numNodesFromK8sAPI {
		sc.logger.Warnf("There are %d nodes named \"*%s*\" in the Kubernetes cluster, but there are %d VMs named \"*%s*\" in the resource group %s\n", numNodesFromK8sAPI, sc.agentPoolToScale, currentNodeCount, sc.agentPoolToScale, sc.resourceGroupName)
	}
}
