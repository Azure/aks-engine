// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"

	"github.com/Azure/aks-engine/pkg/armhelpers"
)

// ListDeploymentOperations gets all deployments operations for a deployment.
func (az *AzureClient) ListDeploymentOperations(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (armhelpers.DeploymentOperationsListResultPage, error) {
	list, err := az.deploymentOperationsClient.List(ctx, resourceGroupName, deploymentName, top)
	return &list, err
}
