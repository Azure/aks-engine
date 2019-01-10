package armhelpers

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"testing"
)

func TestSizeMap(t *testing.T) {
	sizeMap := getSizeMap()
	b, _ := json.MarshalIndent(sizeMap["vmSizesMap"], "", "   ")
	fmt.Println(string(b))
}

func TestK8sVars(t *testing.T) {
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
			},
			LinuxProfile: &api.LinuxProfile{
			},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:   "agentpool1",
					VMSize: "Standard_D2_v2",
					Count:  2,
				},
			},
		},
	}

	cs.SetPropertiesDefaults(false, false)

	varMap := getK8sVars(cs)
	jsonStr := map[string]interface{} {
		"variables": varMap,
	}
	jsonObj, _ := json.MarshalIndent(jsonStr, "", "   ")
	fmt.Println(string(jsonObj))
}
