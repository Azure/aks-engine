// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// DeployTemplate implements the TemplateDeployer interface for the AzureClient client
func (az *AzureClient) DeployTemplate(ctx context.Context, resourceGroupName, deploymentName string, template map[string]interface{}, parameters map[string]interface{}) (de resources.DeploymentExtended, err error) {
	deployment := resources.Deployment{
		Properties: &resources.DeploymentProperties{
			Template:   &template,
			Parameters: &parameters,
			Mode:       resources.Incremental,
		},
	}

	log.Infof("Starting ARM Deployment %s in resource group %s. This will take some time...", deploymentName, resourceGroupName)
	future, err := az.deploymentsClient.CreateOrUpdate(ctx, resourceGroupName, deploymentName, deployment)
	if err != nil {
		return de, err
	}

	outcomeText := "Succeeded"
	err = future.WaitForCompletionRef(ctx, az.deploymentsClient.Client)
	if err != nil {
		outcomeText = fmt.Sprintf("Error: %v", err)
		log.Infof("Finished ARM Deployment (%s). %s", deploymentName, outcomeText)
		return de, err
	}

	de, err = future.Result(az.deploymentsClient)
	if err != nil {
		outcomeText = fmt.Sprintf("Error: %v", err)
	}

	log.Infof("Finished ARM Deployment (%s). %s", deploymentName, outcomeText)
	return de, err
}

// DeployTemplateResult is a struct result type for use with DeployTemplateWithRetry
type DeployTemplateResult struct {
	DeploymentResources resources.DeploymentExtended
	Err                 error
}

// DeployTemplateWithRetry will submit an ARM deployment, retrying if error up to a timeout
func (az *AzureClient) DeployTemplateWithRetry(ctx context.Context, resourceGroupName, deploymentName string, template map[string]interface{}, parameters map[string]interface{}, sleep, timeout time.Duration) (de resources.DeploymentExtended, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan DeployTemplateResult)
	var mostRecentDeployTemplateWithRetryError error
	var deploymentResources resources.DeploymentExtended
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				de, err := az.DeployTemplate(ctx, resourceGroupName, deploymentName, template, parameters)
				ch <- DeployTemplateResult{
					DeploymentResources: de,
					Err:                 err,
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentDeployTemplateWithRetryError = result.Err
			deploymentResources = result.DeploymentResources
			if mostRecentDeployTemplateWithRetryError == nil {
				return deploymentResources, nil
			}
		case <-ctx.Done():
			return deploymentResources, errors.Errorf("DeployTemplateWithRetry timed out: %s\n", mostRecentDeployTemplateWithRetryError)
		}
	}
}

// ValidateTemplate validate the template and parameters
func (az *AzureClient) ValidateTemplate(
	ctx context.Context,
	resourceGroupName string,
	deploymentName string,
	template map[string]interface{},
	parameters map[string]interface{}) (result resources.DeploymentValidateResult, err error) {
	deployment := resources.Deployment{
		Properties: &resources.DeploymentProperties{
			Template:   &template,
			Parameters: &parameters,
			Mode:       resources.Incremental,
		},
	}
	return az.deploymentsClient.Validate(ctx, resourceGroupName, deploymentName, deployment)
}

// GetDeployment returns the template deployment
func (az *AzureClient) GetDeployment(ctx context.Context, resourceGroupName, deploymentName string) (result resources.DeploymentExtended, err error) {
	return az.deploymentsClient.Get(ctx, resourceGroupName, deploymentName)
}

// CheckDeploymentExistence returns if the deployment already exists
func (az *AzureClient) CheckDeploymentExistence(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error) {
	return az.deploymentsClient.CheckExistence(ctx, resourceGroupName, deploymentName)
}
