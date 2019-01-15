// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestGetK8sOutputs(t *testing.T) {

	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:     1,
				DNSPrefix: "blueorange",
				VMSize:    "Standard_D2_v2",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: api.Kubernetes,
				KubernetesConfig: &api.KubernetesConfig{
					PrivateCluster: &api.PrivateCluster{
						Enabled: to.BoolPtr(false),
					},
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:   "agentpool1",
					VMSize: "Standard_D2_v2",
					Count:  2,
				},
			},
		},
	}

	outputMap := GetKubernetesOutputs(cs)
	jsonStr := map[string]interface{}{
		"outputs": outputMap,
	}
	jsonObj, _ := json.MarshalIndent(jsonStr, "", "   ")
	fmt.Println(string(jsonObj))
}
