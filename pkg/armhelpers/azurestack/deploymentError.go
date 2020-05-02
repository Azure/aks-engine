// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/sirupsen/logrus"
)

// DeploymentError contains the root deployment error along with deployment operation errors
type DeploymentError struct {
	DeploymentName    string
	ResourceGroup     string
	TopError          error
	StatusCode        int
	Response          []byte
	ProvisioningState string
	OperationsLists   []resources.DeploymentOperationsListResult
}

// Error implements error interface
func (e *DeploymentError) Error() string {
	var str string
	if e.TopError != nil {
		str = e.TopError.Error()
	}
	var ops []string
	for _, operationsList := range e.OperationsLists {
		if operationsList.Value == nil {
			continue
		}
		for _, operation := range *operationsList.Value {
			if operation.Properties != nil && *operation.Properties.ProvisioningState == string(api.Failed) && operation.Properties.StatusMessage != nil {
				if b, err := json.MarshalIndent(operation.Properties.StatusMessage, "", "  "); err == nil {
					ops = append(ops, string(b))
				}
			}
		}
	}
	return fmt.Sprintf("DeploymentName[%s] ResourceGroup[%s] TopError[%s] StatusCode[%d] Response[%s] ProvisioningState[%s] Operations[%s]",
		e.DeploymentName, e.ResourceGroup, str, e.StatusCode, e.Response, e.ProvisioningState, strings.Join(ops, " | "))
}

// DeploymentValidationError contains validation error
type DeploymentValidationError struct {
	Err error
}

// Error implements error interface
func (e *DeploymentValidationError) Error() string {
	return e.Err.Error()
}

// DeployTemplateSync deploys the template and returns ArmError
func DeployTemplateSync(az armhelpers.AKSEngineClient, logger *logrus.Entry, resourceGroupName, deploymentName string, template map[string]interface{}, parameters map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultARMOperationTimeout)
	defer cancel()
	deploymentExtended, err := az.DeployTemplate(ctx, resourceGroupName, deploymentName, template, parameters)
	if err == nil {
		return nil
	}

	logger.Infof("Getting detailed deployment errors for %s", deploymentName)
	deploymentErr := &DeploymentError{
		DeploymentName: deploymentName,
		ResourceGroup:  resourceGroupName,
		TopError:       err,
	}

	// try to extract error from ARM Response
	if deploymentExtended.Response.Response != nil && deploymentExtended.Body != nil {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(deploymentExtended.Body)
		logger.Infof("StatusCode: %d, Error: %s", deploymentExtended.Response.StatusCode, buf.String())
		deploymentErr.Response = buf.Bytes()
		deploymentErr.StatusCode = deploymentExtended.Response.StatusCode
	} else {
		logger.Errorf("Got error from Azure SDK without response from ARM")
		// This is the failed sdk validation before calling ARM path
		return deploymentErr
	}

	if deploymentExtended.Properties == nil || deploymentExtended.Properties.ProvisioningState == nil {
		logger.Warn("No resources.DeploymentExtended.Properties")
		return deploymentErr
	}
	properties := deploymentExtended.Properties
	deploymentErr.ProvisioningState = *properties.ProvisioningState

	for page, err := az.ListDeploymentOperations(ctx, resourceGroupName, deploymentName, nil); page.NotDone(); err = page.Next() {
		if err != nil {
			logger.Errorf("unable to list deployment operations %s. error: %v", deploymentName, err)
			return deploymentErr
		}
		deploymentErr.OperationsLists = append(deploymentErr.OperationsLists, page.Response())
	}

	return deploymentErr
}
