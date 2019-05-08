// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestKubernetesContainerAddonSettingsInit(t *testing.T) {
	cases := []struct {
		p                         *api.Properties
		expectedHeapster          bool
		expectedMetricsServer     bool
		expectedTiller            bool
		expectedAADPodIdentity    bool
		expectedACIConnector      bool
		expectedClusterAutoscaler bool
		// TODO add the remaining addons supported by kubernetesContainerAddonSettingsInit
	}{
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.8.15",
					KubernetesConfig:    &api.KubernetesConfig{},
				},
			},
			expectedHeapster:          true,
			expectedMetricsServer:     false,
			expectedTiller:            true,
			expectedAADPodIdentity:    false,
			expectedACIConnector:      false,
			expectedClusterAutoscaler: false,
		},
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						Addons: []api.KubernetesAddon{
							{
								Name:    TillerAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    AADPodIdentityAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    ACIConnectorAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    ClusterAutoscalerAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedHeapster:          false,
			expectedMetricsServer:     true,
			expectedTiller:            false,
			expectedAADPodIdentity:    true,
			expectedACIConnector:      true,
			expectedClusterAutoscaler: true,
		},
	}

	for _, c := range cases {
		componentFileSpec := kubernetesContainerAddonSettingsInit(c.p)
		if c.expectedHeapster != componentFileSpec[DefaultHeapsterAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", DefaultHeapsterAddonName, c.expectedHeapster)
		}
		if c.expectedMetricsServer != componentFileSpec[MetricsServerAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", MetricsServerAddonName, c.expectedMetricsServer)
		}
		if c.expectedTiller != componentFileSpec[TillerAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", TillerAddonName, c.expectedTiller)
		}
		if c.expectedAADPodIdentity != componentFileSpec[AADPodIdentityAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", AADPodIdentityAddonName, c.expectedAADPodIdentity)
		}
		if c.expectedACIConnector != componentFileSpec[ACIConnectorAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", ACIConnectorAddonName, c.expectedACIConnector)
		}
		if c.expectedClusterAutoscaler != componentFileSpec[ClusterAutoscalerAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", ClusterAutoscalerAddonName, c.expectedClusterAutoscaler)
		}
	}
}

func TestKubernetesAddonSettingsInit(t *testing.T) {
	// TODO add tests for kubernetesAddonSettingsInit
}

func TestKubernetesManifestSettingsInit(t *testing.T) {
	// TODO add tests for kubernetesManifestSettingsInit
}
