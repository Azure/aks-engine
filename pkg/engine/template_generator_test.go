package engine

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"testing"
)

func TestGenerateTemplateV2(t *testing.T) {
	tg, _ := InitializeTemplateGenerator(Context{})

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

	paramsMap, err := tg.getParameterDescMap(cs)

	if err != nil {
		t.Errorf("unexpected error while running getParameterDescMap: %s", err.Error())
	}

	fmt.Println(paramsMap)
}
