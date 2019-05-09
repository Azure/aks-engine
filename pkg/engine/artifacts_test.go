// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestKubernetesContainerAddonSettingsInit(t *testing.T) {
	mockAzureStackProperties := api.GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	cases := []struct {
		p                              *api.Properties
		expectedHeapster               bool
		expectedMetricsServer          bool
		expectedTiller                 bool
		expectedAADPodIdentity         bool
		expectedACIConnector           bool
		expectedClusterAutoscaler      bool
		expectedBlobfuseFlexvolume     bool
		expectedSMBFlexvolume          bool
		expectedKeyvaultFlexvolume     bool
		expectedDashboard              bool
		expectedRescheduler            bool
		expectedNvidia                 bool
		expectedContainerMonitoring    bool
		expectedIPMasqAgent            bool
		expectedAzureCNINetworkMonitor bool
		expectedDNSAutoscaler          bool
		expectedCalico                 bool
	}{
		// Legacy default scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.8.15",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
					},
				},
			},
			expectedHeapster:               true,
			expectedMetricsServer:          false,
			expectedTiller:                 true,
			expectedAADPodIdentity:         false,
			expectedACIConnector:           false,
			expectedClusterAutoscaler:      false,
			expectedBlobfuseFlexvolume:     true,
			expectedSMBFlexvolume:          false,
			expectedKeyvaultFlexvolume:     true,
			expectedDashboard:              true,
			expectedRescheduler:            false,
			expectedNvidia:                 false,
			expectedContainerMonitoring:    false,
			expectedIPMasqAgent:            true,
			expectedAzureCNINetworkMonitor: true,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 false,
		},
		// 1.14 scenario w/ explicit addons config
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
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
							{
								Name:    BlobfuseFlexVolumeAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    SMBFlexVolumeAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    KeyVaultFlexVolumeAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    DashboardAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    ReschedulerAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    NVIDIADevicePluginAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    ContainerMonitoringAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    IPMASQAgentAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    AzureCNINetworkMonitorAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    DNSAutoscalerAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    CalicoAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedHeapster:               false,
			expectedMetricsServer:          true,
			expectedTiller:                 false,
			expectedAADPodIdentity:         true,
			expectedACIConnector:           true,
			expectedClusterAutoscaler:      true,
			expectedBlobfuseFlexvolume:     false,
			expectedSMBFlexvolume:          true,
			expectedKeyvaultFlexvolume:     false,
			expectedDashboard:              false,
			expectedRescheduler:            true,
			expectedNvidia:                 false,
			expectedContainerMonitoring:    true,
			expectedIPMasqAgent:            false,
			expectedAzureCNINetworkMonitor: false,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 false,
		},
		// Azure Stack scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
						Addons: []api.KubernetesAddon{
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
							{
								Name:    BlobfuseFlexVolumeAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    SMBFlexVolumeAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    KeyVaultFlexVolumeAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    ReschedulerAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    NVIDIADevicePluginAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    ContainerMonitoringAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
				CustomCloudProfile: mockAzureStackProperties.CustomCloudProfile,
			},
			expectedHeapster:               false,
			expectedMetricsServer:          true,
			expectedTiller:                 true,
			expectedAADPodIdentity:         false,
			expectedACIConnector:           false,
			expectedClusterAutoscaler:      false,
			expectedBlobfuseFlexvolume:     false,
			expectedSMBFlexvolume:          false,
			expectedKeyvaultFlexvolume:     false,
			expectedDashboard:              true,
			expectedRescheduler:            false,
			expectedNvidia:                 false,
			expectedContainerMonitoring:    false,
			expectedIPMasqAgent:            true,
			expectedAzureCNINetworkMonitor: true,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 false,
		},
		// N Series SKU scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
					},
				},
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						VMSize: "Standard_NC6",
					},
				},
			},
			expectedHeapster:               false,
			expectedMetricsServer:          true,
			expectedTiller:                 true,
			expectedAADPodIdentity:         false,
			expectedACIConnector:           false,
			expectedClusterAutoscaler:      false,
			expectedBlobfuseFlexvolume:     true,
			expectedSMBFlexvolume:          false,
			expectedKeyvaultFlexvolume:     true,
			expectedDashboard:              true,
			expectedRescheduler:            false,
			expectedNvidia:                 true,
			expectedContainerMonitoring:    false,
			expectedIPMasqAgent:            true,
			expectedAzureCNINetworkMonitor: true,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 false,
		},
		// cilium scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginCilium,
					},
				},
			},
			expectedHeapster:               false,
			expectedMetricsServer:          true,
			expectedTiller:                 true,
			expectedAADPodIdentity:         false,
			expectedACIConnector:           false,
			expectedClusterAutoscaler:      false,
			expectedBlobfuseFlexvolume:     true,
			expectedSMBFlexvolume:          false,
			expectedKeyvaultFlexvolume:     true,
			expectedDashboard:              true,
			expectedRescheduler:            false,
			expectedNvidia:                 false,
			expectedContainerMonitoring:    false,
			expectedIPMasqAgent:            false,
			expectedAzureCNINetworkMonitor: false,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 false,
		},
		// kubenet scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginKubenet,
					},
				},
			},
			expectedHeapster:               false,
			expectedMetricsServer:          true,
			expectedTiller:                 true,
			expectedAADPodIdentity:         false,
			expectedACIConnector:           false,
			expectedClusterAutoscaler:      false,
			expectedBlobfuseFlexvolume:     true,
			expectedSMBFlexvolume:          false,
			expectedKeyvaultFlexvolume:     true,
			expectedDashboard:              true,
			expectedRescheduler:            false,
			expectedNvidia:                 false,
			expectedContainerMonitoring:    false,
			expectedIPMasqAgent:            true,
			expectedAzureCNINetworkMonitor: false,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 false,
		},
		// calico scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPolicy: NetworkPolicyCalico,
					},
				},
			},
			expectedHeapster:               false,
			expectedMetricsServer:          true,
			expectedTiller:                 true,
			expectedAADPodIdentity:         false,
			expectedACIConnector:           false,
			expectedClusterAutoscaler:      false,
			expectedBlobfuseFlexvolume:     true,
			expectedSMBFlexvolume:          false,
			expectedKeyvaultFlexvolume:     true,
			expectedDashboard:              true,
			expectedRescheduler:            false,
			expectedNvidia:                 false,
			expectedContainerMonitoring:    false,
			expectedIPMasqAgent:            true,
			expectedAzureCNINetworkMonitor: false,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 true,
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
		if c.expectedBlobfuseFlexvolume != componentFileSpec[BlobfuseFlexVolumeAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", BlobfuseFlexVolumeAddonName, c.expectedBlobfuseFlexvolume)
		}
		if c.expectedSMBFlexvolume != componentFileSpec[SMBFlexVolumeAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", SMBFlexVolumeAddonName, c.expectedSMBFlexvolume)
		}
		if c.expectedKeyvaultFlexvolume != componentFileSpec[KeyVaultFlexVolumeAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", KeyVaultFlexVolumeAddonName, c.expectedKeyvaultFlexvolume)
		}
		if c.expectedDashboard != componentFileSpec[DashboardAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", DashboardAddonName, c.expectedDashboard)
		}
		if c.expectedRescheduler != componentFileSpec[ReschedulerAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", ReschedulerAddonName, c.expectedRescheduler)
		}
		if c.expectedNvidia != componentFileSpec[NVIDIADevicePluginAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", NVIDIADevicePluginAddonName, c.expectedNvidia)
		}
		if c.expectedContainerMonitoring != componentFileSpec[ContainerMonitoringAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", ContainerMonitoringAddonName, c.expectedContainerMonitoring)
		}
		if c.expectedIPMasqAgent != componentFileSpec[IPMASQAgentAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", IPMASQAgentAddonName, c.expectedIPMasqAgent)
		}
		if c.expectedAzureCNINetworkMonitor != componentFileSpec[AzureCNINetworkMonitorAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", AzureCNINetworkMonitorAddonName, c.expectedAzureCNINetworkMonitor)
		}
		if c.expectedDNSAutoscaler != componentFileSpec[DNSAutoscalerAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", DNSAutoscalerAddonName, c.expectedDNSAutoscaler)
		}
		if c.expectedCalico != componentFileSpec[CalicoAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", CalicoAddonName, c.expectedCalico)
		}
	}
}

func TestKubernetesAddonSettingsInit(t *testing.T) {
	cases := []struct {
		p                          *api.Properties
		expectedKubeDNS            bool
		expectedCoreDNS            bool
		expectedKubeProxy          bool
		expectedAzureNetworkPolicy bool
		expectedCilium             bool
		expectedFlannel            bool
		expectedAADAdminGroup      bool
		expectedAzureCloudProvider bool
		expectedAuditPolicy        bool
		expectedELBService         bool
	}{
		// Legacy default scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.7.10",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
					},
				},
			},
			expectedKubeDNS:            true,
			expectedCoreDNS:            false,
			expectedKubeProxy:          true,
			expectedAzureNetworkPolicy: false,
			expectedCilium:             false,
			expectedFlannel:            false,
			expectedAADAdminGroup:      false,
			expectedAzureCloudProvider: true,
			expectedAuditPolicy:        false,
			expectedELBService:         false,
		},
		// 1.14 default scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
					},
				},
			},
			expectedKubeDNS:            false,
			expectedCoreDNS:            true,
			expectedKubeProxy:          true,
			expectedAzureNetworkPolicy: false,
			expectedCilium:             false,
			expectedFlannel:            false,
			expectedAADAdminGroup:      false,
			expectedAzureCloudProvider: true,
			expectedAuditPolicy:        true,
			expectedELBService:         false,
		},
		// Azure network policy scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
						NetworkPolicy: NetworkPolicyAzure,
					},
				},
			},
			expectedKubeDNS:            false,
			expectedCoreDNS:            true,
			expectedKubeProxy:          true,
			expectedAzureNetworkPolicy: true,
			expectedCilium:             false,
			expectedFlannel:            false,
			expectedAADAdminGroup:      false,
			expectedAzureCloudProvider: true,
			expectedAuditPolicy:        true,
			expectedELBService:         false,
		},
		// Cilium scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPolicy: NetworkPolicyCilium,
					},
				},
			},
			expectedKubeDNS:            false,
			expectedCoreDNS:            true,
			expectedKubeProxy:          true,
			expectedAzureNetworkPolicy: false,
			expectedCilium:             true,
			expectedFlannel:            false,
			expectedAADAdminGroup:      false,
			expectedAzureCloudProvider: true,
			expectedAuditPolicy:        true,
			expectedELBService:         false,
		},
		// Flannel scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginFlannel,
					},
				},
			},
			expectedKubeDNS:            false,
			expectedCoreDNS:            true,
			expectedKubeProxy:          true,
			expectedAzureNetworkPolicy: false,
			expectedCilium:             false,
			expectedFlannel:            true,
			expectedAADAdminGroup:      false,
			expectedAzureCloudProvider: true,
			expectedAuditPolicy:        true,
			expectedELBService:         false,
		},
		// AAD Admin Group scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
					},
				},
				AADProfile: &api.AADProfile{
					AdminGroupID: "1234-5",
				},
			},
			expectedKubeDNS:            false,
			expectedCoreDNS:            true,
			expectedKubeProxy:          true,
			expectedAzureNetworkPolicy: false,
			expectedCilium:             false,
			expectedFlannel:            false,
			expectedAADAdminGroup:      true,
			expectedAzureCloudProvider: true,
			expectedAuditPolicy:        true,
			expectedELBService:         false,
		},
		// ELB service scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin:   NetworkPluginAzure,
						LoadBalancerSku: api.StandardLoadBalancerSku,
					},
				},
			},
			expectedKubeDNS:            false,
			expectedCoreDNS:            true,
			expectedKubeProxy:          true,
			expectedAzureNetworkPolicy: false,
			expectedCilium:             false,
			expectedFlannel:            false,
			expectedAADAdminGroup:      false,
			expectedAzureCloudProvider: true,
			expectedAuditPolicy:        true,
			expectedELBService:         true,
		},
	}

	for _, c := range cases {
		componentFileSpecArray := kubernetesAddonSettingsInit(c.p)
		for _, componentFileSpec := range componentFileSpecArray {
			switch componentFileSpec.destinationFile {
			case "kube-dns-deployment.yaml":
				if c.expectedKubeDNS != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", KubeDNSAddonName, c.expectedKubeDNS)
				}
			case "coredns.yaml":
				if c.expectedCoreDNS != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", CoreDNSAddonName, c.expectedCoreDNS)
				}
			case "kube-proxy-daemonset.yaml":
				if c.expectedKubeProxy != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", KubeProxyAddonName, c.expectedKubeProxy)
				}
			case "azure-npm-daemonset.yaml":
				if c.expectedAzureNetworkPolicy != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", AzureNetworkPolicyAddonName, c.expectedAzureNetworkPolicy)
				}
			case "cilium-daemonset.yaml":
				if c.expectedCilium != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", CiliumAddonName, c.expectedCilium)
				}
			case "flannel-daemonset.yaml":
				if c.expectedFlannel != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", FlannelAddonName, c.expectedFlannel)
				}
			case "aad-default-admin-group-rbac.yaml":
				if c.expectedAADAdminGroup != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", AADAdminGroupAddonName, c.expectedAADAdminGroup)
				}
			case "azure-cloud-provider-deployment.yaml":
				if c.expectedAzureCloudProvider != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", AzureCloudProviderAddonName, c.expectedAzureCloudProvider)
				}
			case "audit-policy.yaml":
				if c.expectedAuditPolicy != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", AuditPolicyAddonName, c.expectedAuditPolicy)
				}
			case "elb-svc.yaml":
				if c.expectedELBService != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", ELBServiceAddonName, c.expectedELBService)
				}
			}
		}
	}
}

func TestKubernetesManifestSettingsInit(t *testing.T) {
	// TODO add tests for kubernetesManifestSettingsInit
}
