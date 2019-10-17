// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"
)

func TestCreateAzurestackTelemetry(t *testing.T) {
	properties := resources.DeploymentPropertiesExtended{
		Mode: "Incremental",
		Template: map[string]interface{}{
			"resources":      []interface{}{},
			"contentVersion": "1.0.0.0",
			"$schema":        "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
		},
	}

	pid := "pid-1bda96ec-adf4-4eea-bb9a-8462de5475c0"
	actual := createAzureStackTelemetry(pid)
	expected := DeploymentARM{
		DeploymentARMResource: DeploymentARMResource{
			APIVersion: "2015-01-01",
		},
		DeploymentExtended: resources.DeploymentExtended{
			Name:       to.StringPtr(pid),
			Type:       to.StringPtr("Microsoft.Resources/deployments"),
			Properties: &properties,
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}
