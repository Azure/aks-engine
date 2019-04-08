// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/operations"
	"github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	interval           = time.Second * 1
	retry              = time.Second * 5
	cordonDrainTimeout = time.Minute * 20
)

// Compiler to verify QueueMessageProcessor implements OperationsProcessor
var _ UpgradeNode = &UpgradeAgentNode{}

// UpgradeAgentNode upgrades a Kubernetes 1.5 agent node to 1.6
type UpgradeAgentNode struct {
	Translator              *i18n.Translator
	logger                  *logrus.Entry
	TemplateMap             map[string]interface{}
	ParametersMap           map[string]interface{}
	UpgradeContainerService *api.ContainerService
	SubscriptionID          string
	ResourceGroup           string
	Client                  armhelpers.AKSEngineClient
	kubeConfig              string
	timeout                 time.Duration
}

// DeleteNode takes state/resources of the master/agent node from ListNodeResources
// backs up/preserves state as needed by a specific version of Kubernetes and then deletes
// the node
// The 'drain' flag is used to invoke 'cordon and drain' flow.
func (kan *UpgradeAgentNode) DeleteNode(vmName *string, drain bool) error {
	var kubeAPIServerURL string

	if kan.UpgradeContainerService.Properties.HostedMasterProfile != nil {
		kubeAPIServerURL = kan.UpgradeContainerService.Properties.HostedMasterProfile.FQDN
	} else {
		kubeAPIServerURL = kan.UpgradeContainerService.Properties.MasterProfile.FQDN
	}

	if vmName == nil || *vmName == "" {
		return errors.Errorf("Error deleting VM: VM name was empty")
	}

	nodeName := strings.ToLower(*vmName)

	client, err := kan.Client.GetKubernetesClient(kubeAPIServerURL, kan.kubeConfig, interval, kan.timeout)
	if err != nil {
		return err
	}
	// Cordon and drain the node
	if drain {
		err = operations.SafelyDrainNodeWithClient(client, kan.logger, nodeName, cordonDrainTimeout)
		if err != nil {
			kan.logger.Warningf("Error draining agent VM %s. Proceeding with deletion. Error: %v", *vmName, err)
			// Proceed with deletion anyways
		}
	}
	// Delete VM in ARM
	if err = operations.CleanDeleteVirtualMachine(kan.Client, kan.logger, kan.SubscriptionID, kan.ResourceGroup, *vmName); err != nil {
		return err
	}
	// Delete VM in api server
	if err = client.DeleteNode(nodeName); err != nil {
		statusErr, ok := err.(*apierrors.StatusError)
		if ok && statusErr.ErrStatus.Reason != v1.StatusReasonNotFound {
			kan.logger.Warnf("Node %s got an error while deregistering: %#v", *vmName, err)
		}
	}
	return nil
}

// CreateNode creates a new master/agent node with the targeted version of Kubernetes
func (kan *UpgradeAgentNode) CreateNode(ctx context.Context, poolName string, agentNo int) error {
	poolCountParameter := kan.ParametersMap[poolName+"Count"].(map[string]interface{})
	poolCountParameter["value"] = agentNo + 1
	agentCount := poolCountParameter["value"]
	kan.logger.Infof("Agent pool: %s, set count to: %d temporarily during upgrade. Upgrading agent: %d",
		poolName, agentCount, agentNo)

	poolOffsetVarName := poolName + "Offset"
	templateVariables := kan.TemplateMap["variables"].(map[string]interface{})
	templateVariables[poolOffsetVarName] = agentNo

	// Debug function - keep commented out
	// WriteTemplate(kan.Translator, kan.UpgradeContainerService, kan.TemplateMap, kan.ParametersMap)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	deploymentSuffix := random.Int31()
	deploymentName := fmt.Sprintf("agent-%s-%d", time.Now().Format("06-01-02T15.04.05"), deploymentSuffix)

	return armhelpers.DeployTemplateSync(kan.Client, kan.logger, kan.ResourceGroup, deploymentName, kan.TemplateMap, kan.ParametersMap)
}

// Validate will verify that agent node has been upgraded as expected.
func (kan *UpgradeAgentNode) Validate(vmName *string) error {
	if vmName == nil || *vmName == "" {
		kan.logger.Warningf("VM name was empty. Skipping node condition check")
		return nil
	}
	nodeName := strings.ToLower(*vmName)

	kan.logger.Infof("Validating %s", nodeName)
	var masterURL string
	if kan.UpgradeContainerService.Properties.HostedMasterProfile != nil {
		masterURL = kan.UpgradeContainerService.Properties.HostedMasterProfile.FQDN
	} else {
		masterURL = kan.UpgradeContainerService.Properties.MasterProfile.FQDN
	}

	client, err := kan.Client.GetKubernetesClient(masterURL, kan.kubeConfig, interval, kan.timeout)
	if err != nil {
		return &armhelpers.DeploymentValidationError{Err: err}
	}

	retryTimer := time.NewTimer(time.Millisecond)
	timeoutTimer := time.NewTimer(kan.timeout)
	for {
		select {
		case <-timeoutTimer.C:
			retryTimer.Stop()
			return &armhelpers.DeploymentValidationError{Err: kan.Translator.Errorf("Node was not ready within %v", kan.timeout)}
		case <-retryTimer.C:
			agentNode, err := client.GetNode(nodeName)
			if err != nil {
				kan.logger.Infof("Agent node: %s status error: %v", nodeName, err)
				retryTimer.Reset(retry)
			} else if isNodeReady(agentNode) {
				kan.logger.Infof("Agent node: %s is ready", nodeName)
				timeoutTimer.Stop()
				return nil
			} else {
				kan.logger.Infof("Agent node: %s not ready yet...", nodeName)
				retryTimer.Reset(retry)
			}
		}
	}
}
