// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
)

func TestCreateVirtualMachines(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               3,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin: "azure",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
		},
	}

	cs.SetPropertiesDefaults(false, false)

	vm := CreateVirtualMachine(cs)
	jsonObj, _ := json.MarshalIndent(vm, "", "   ")
	fmt.Println(string(jsonObj))
}
