// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"reflect"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/osa/vlabs"
)

var testOpenShiftCluster = &vlabs.OpenShiftCluster{
	ID:       "id",
	Location: "location",
	Name:     "name",
	Plan: &vlabs.ResourcePurchasePlan{
		Name:          "plan.name",
		Product:       "plan.product",
		PromotionCode: "plan.promotionCode",
		Publisher:     "plan.publisher",
	},
	Tags: map[string]string{
		"tags.k1": "v1",
		"tags.k2": "v2",
	},
	Type: "type",
	Properties: &vlabs.Properties{
		ProvisioningState: "properties.provisioningState",
		OpenShiftVersion:  "properties.openShiftVersion",
		PublicHostname:    "properties.publicHostname",
		FQDN:              "properties.fqdn",
		RouterProfiles: []vlabs.RouterProfile{
			{
				Name:            "properties.routerProfiles.0.name",
				PublicSubdomain: "properties.routerProfiles.0.publicSubdomain",
				FQDN:            "properties.routerProfiles.0.fqdn",
			},
			{
				Name:            "properties.routerProfiles.1.name",
				PublicSubdomain: "properties.routerProfiles.1.publicSubdomain",
				FQDN:            "properties.routerProfiles.1.fqdn",
			},
		},
		AgentPoolProfiles: []vlabs.AgentPoolProfile{
			{
				Name:         "properties.agentPoolProfiles.0.name",
				Role:         "properties.agentPoolProfiles.0.role",
				Count:        1,
				VMSize:       "properties.agentPoolProfiles.0.vmSize",
				VnetSubnetID: "properties.agentPoolProfiles.0.vnetSubnetID",
				OSType:       "properties.agentPoolProfiles.0.osType",
			},
			{
				Name:         "properties.agentPoolProfiles.0.name",
				Role:         "properties.agentPoolProfiles.0.role",
				Count:        2,
				VMSize:       "properties.agentPoolProfiles.0.vmSize",
				VnetSubnetID: "properties.agentPoolProfiles.0.vnetSubnetID",
				OSType:       "properties.agentPoolProfiles.0.osType",
			},
		},
		ServicePrincipalProfile: vlabs.ServicePrincipalProfile{
			ClientID: "properties.servicePrincipalProfile.clientID",
			Secret:   "properties.servicePrincipalProfile.secret",
		},
	},
}

var testContainerService = &ContainerService{
	ID:       "id",
	Location: "location",
	Name:     "name",
	Plan: &ResourcePurchasePlan{
		Name:          "plan.name",
		Product:       "plan.product",
		PromotionCode: "plan.promotionCode",
		Publisher:     "plan.publisher",
	},
	Tags: map[string]string{
		"tags.k1": "v1",
		"tags.k2": "v2",
	},
	Type: "type",
	Properties: &Properties{
		ProvisioningState: "properties.provisioningState",
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorVersion: "properties.openShiftVersion",
			OpenShiftConfig: &OpenShiftConfig{
				PublicHostname: "properties.publicHostname",
				RouterProfiles: []OpenShiftRouterProfile{
					{
						Name:            "properties.routerProfiles.0.name",
						PublicSubdomain: "properties.routerProfiles.0.publicSubdomain",
						FQDN:            "properties.routerProfiles.0.fqdn",
					},
					{
						Name:            "properties.routerProfiles.1.name",
						PublicSubdomain: "properties.routerProfiles.1.publicSubdomain",
						FQDN:            "properties.routerProfiles.1.fqdn",
					},
				},
			},
		},
		MasterProfile: &MasterProfile{
			FQDN: "properties.fqdn",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:         "properties.agentPoolProfiles.0.name",
				Count:        1,
				VMSize:       "properties.agentPoolProfiles.0.vmSize",
				OSType:       "properties.agentPoolProfiles.0.osType",
				VnetSubnetID: "properties.agentPoolProfiles.0.vnetSubnetID",
				Role:         "properties.agentPoolProfiles.0.role",
			},
			{
				Name:         "properties.agentPoolProfiles.0.name",
				Count:        2,
				VMSize:       "properties.agentPoolProfiles.0.vmSize",
				OSType:       "properties.agentPoolProfiles.0.osType",
				VnetSubnetID: "properties.agentPoolProfiles.0.vnetSubnetID",
				Role:         "properties.agentPoolProfiles.0.role",
			},
		},
		ServicePrincipalProfile: &ServicePrincipalProfile{
			ClientID: "properties.servicePrincipalProfile.clientID",
			Secret:   "properties.servicePrincipalProfile.secret",
		},
	},
}

func TestConvertVLabsOpenShiftClusterToContainerService(t *testing.T) {
	cs := ConvertVLabsOpenShiftClusterToContainerService(testOpenShiftCluster)
	if !reflect.DeepEqual(cs, testContainerService) {
		t.Errorf("ConvertVLabsOpenShiftClusterToContainerService returned unexpected result\n%#v\n", cs)
	}
}
