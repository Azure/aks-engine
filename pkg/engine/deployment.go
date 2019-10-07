// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
)

func createAzureStackTelemetry(azureTelemetryPID string) DeploymentARM {
	properties := resources.DeploymentPropertiesExtended{
		Mode: "Incremental",
		Template: map[string]interface{}{
			"resources":      []interface{}{},
			"contentVersion": "1.0.0.0",
			"$schema":        "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
		},
	}

	return DeploymentARM{
		DeploymentARMResource: DeploymentARMResource{
			APIVersion: "2015-01-01",
		},
		DeploymentExtended: resources.DeploymentExtended{
			Name:       to.StringPtr(azureTelemetryPID),
			Type:       to.StringPtr("Microsoft.Resources/deployments"),
			Properties: &properties,
		},
	}
}
