package armhelpers

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"testing"
)

func TestGenerateARMResources(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               2,
				},
			},
		},
	}

	cs.SetPropertiesDefaults(false, false)
	armResources := GenerateARMResources(cs)

	jsonStr := map[string]interface{} {
		"resources": armResources,
	}
	jsonObj, _ := json.MarshalIndent(jsonStr, "", "   ")
	fmt.Println(string(jsonObj))

}
