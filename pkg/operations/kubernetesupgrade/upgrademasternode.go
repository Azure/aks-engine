// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/kubernetes"
	"github.com/Azure/aks-engine/pkg/operations"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Compiler to verify QueueMessageProcessor implements OperationsProcessor
var _ UpgradeNode = &UpgradeMasterNode{}

// UpgradeMasterNode upgrades a Kubernetes 1.5 master node to 1.6
type UpgradeMasterNode struct {
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
// the node.
// The 'drain' flag is not used for deleting master nodes.
func (kmn *UpgradeMasterNode) DeleteNode(vmName *string, drain bool) error {
	return operations.CleanDeleteVirtualMachine(kmn.Client, kmn.logger, kmn.SubscriptionID, kmn.ResourceGroup, *vmName)
}

// CreateNode creates a new master/agent node with the targeted version of Kubernetes
func (kmn *UpgradeMasterNode) CreateNode(ctx context.Context, poolName string, masterNo int) error {
	templateVariables := kmn.TemplateMap["variables"].(map[string]interface{})

	templateVariables["masterOffset"] = masterNo
	masterOffsetVar := templateVariables["masterOffset"]
	kmn.logger.Infof("Master offset: %v", masterOffsetVar)

	templateVariables["masterCount"] = masterNo + 1
	masterOffset := templateVariables["masterCount"]
	kmn.logger.Infof("Master pool set count to: %v temporarily during upgrade...", masterOffset)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	deploymentSuffix := random.Int31()
	deploymentName := fmt.Sprintf("k8s-upgrade-master-%d-%s-%d", masterNo, time.Now().Format("06-01-02T15.04.05"), deploymentSuffix)

	_, err := kmn.Client.DeployTemplate(
		ctx,
		kmn.ResourceGroup,
		deploymentName,
		kmn.TemplateMap,
		kmn.ParametersMap)
	return err
}

// Validate will verify the that master node has been upgraded as expected.
func (kmn *UpgradeMasterNode) Validate(vmName *string) error {
	if vmName == nil || *vmName == "" {
		kmn.logger.Warningf("VM name was empty. Skipping node condition check")
		return nil
	}

	nodeName := strings.ToLower(*vmName)

	if kmn.UpgradeContainerService.Properties.MasterProfile == nil {
		kmn.logger.Warningf("Master profile was empty. Skipping node condition check")
		return nil
	}

	apiserverURL := kmn.UpgradeContainerService.Properties.MasterProfile.FQDN

	client, err := kmn.Client.GetKubernetesClient(apiserverURL, kmn.kubeConfig, interval, kmn.timeout)
	if err != nil {
		return err
	}

	ch := make(chan struct{}, 1)
	go func() {
		for {
			masterNode, err := client.GetNode(nodeName)
			if err != nil {
				kmn.logger.Infof("Master node: %s status error: %v", nodeName, err)
				time.Sleep(time.Second * 5)
			} else if kubernetes.IsNodeReady(masterNode) {
				kmn.logger.Infof("Master node: %s is ready", nodeName)
				ch <- struct{}{}
			} else {
				kmn.logger.Infof("Master node: %s not ready yet...", nodeName)
				time.Sleep(time.Second * 5)
			}
		}
	}()

	for {
		select {
		case <-ch:
			return nil
		case <-time.After(kmn.timeout):
			kmn.logger.Errorf("Node was not ready within %v", kmn.timeout)
			return errors.Errorf("Node was not ready within %v", kmn.timeout)
		}
	}
}
