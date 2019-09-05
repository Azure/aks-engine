// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
)

const (
	azurestackGenerateGUID = "pid-1bda96ec-adf4-4eea-bb9a-8462de5475c0"
)

//TODO create separate function to allow users to pass in deployment name for azure
func createAzurestackTelemetry() DeploymentARM {
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
			Name:       to.StringPtr(azurestackGenerateGUID),
			Type:       to.StringPtr("Microsoft.Resources/deployments"),
			Properties: &properties,
		},
	}
}
