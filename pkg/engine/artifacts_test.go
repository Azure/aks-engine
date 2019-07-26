// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestKubernetesContainerAddonSettingsInit(t *testing.T) {
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
		expectedAzureNetworkPolicy     bool
	}{
		// addons disabled scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
						Addons: []api.KubernetesAddon{
							{
								Name:    HeapsterAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    MetricsServerAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    TillerAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    AADPodIdentityAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    ACIConnectorAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    ClusterAutoscalerAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    BlobfuseFlexVolumeAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    SMBFlexVolumeAddonName,
								Enabled: to.BoolPtr(false),
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
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    NVIDIADevicePluginAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    ContainerMonitoringAddonName,
								Enabled: to.BoolPtr(false),
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
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    CalicoAddonName,
								Enabled: to.BoolPtr(false),
							},
							{
								Name:    AzureNetworkPolicyAddonName,
								Enabled: to.BoolPtr(false),
							},
						},
					},
				},
			},
			expectedHeapster:               false,
			expectedMetricsServer:          false,
			expectedTiller:                 false,
			expectedAADPodIdentity:         false,
			expectedACIConnector:           false,
			expectedClusterAutoscaler:      false,
			expectedBlobfuseFlexvolume:     false,
			expectedSMBFlexvolume:          false,
			expectedKeyvaultFlexvolume:     false,
			expectedDashboard:              false,
			expectedRescheduler:            false,
			expectedNvidia:                 false,
			expectedContainerMonitoring:    false,
			expectedIPMasqAgent:            false,
			expectedAzureCNINetworkMonitor: false,
			expectedDNSAutoscaler:          false,
			expectedCalico:                 false,
			expectedAzureNetworkPolicy:     false,
		},
		// addons enabled scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
						Addons: []api.KubernetesAddon{
							{
								Name:    HeapsterAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    MetricsServerAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    TillerAddonName,
								Enabled: to.BoolPtr(true),
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
								Name:    DashboardAddonName,
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
							{
								Name:    IPMASQAgentAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    AzureCNINetworkMonitorAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    DNSAutoscalerAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    CalicoAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    AzureNetworkPolicyAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedHeapster:               true,
			expectedMetricsServer:          true,
			expectedTiller:                 true,
			expectedAADPodIdentity:         true,
			expectedACIConnector:           true,
			expectedClusterAutoscaler:      true,
			expectedBlobfuseFlexvolume:     true,
			expectedSMBFlexvolume:          true,
			expectedKeyvaultFlexvolume:     true,
			expectedDashboard:              true,
			expectedRescheduler:            true,
			expectedNvidia:                 true,
			expectedContainerMonitoring:    true,
			expectedIPMasqAgent:            true,
			expectedAzureCNINetworkMonitor: true,
			expectedDNSAutoscaler:          true,
			expectedCalico:                 true,
			expectedAzureNetworkPolicy:     true,
		},
	}

	for _, c := range cases {
		componentFileSpec := kubernetesContainerAddonSettingsInit(c.p)
		if c.expectedHeapster != componentFileSpec[HeapsterAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", HeapsterAddonName, c.expectedHeapster)
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
		if c.expectedAzureNetworkPolicy != componentFileSpec[AzureNetworkPolicyAddonName].isEnabled {
			t.Fatalf("Expected componentFileSpec[%s] to be %t", AzureNetworkPolicyAddonName, c.expectedAzureNetworkPolicy)
		}
	}
}

func TestKubernetesAddonSettingsInit(t *testing.T) {
	mockAzureStackProperties := api.GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	cases := []struct {
		p                             *api.Properties
		expectedKubeDNS               bool
		expectedCoreDNS               bool
		expectedKubeProxy             bool
		expectedCilium                bool
		expectedFlannel               bool
		expectedAADAdminGroup         bool
		expectedAzureCloudProvider    bool
		expectedAuditPolicy           bool
		expectedPodSecurityPolicy     bool
		expectedManagedStorageClass   bool
		expectedUnmanagedStorageClass bool
		expectedScheduledMaintenance  bool
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
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						StorageProfile: api.ManagedDisks,
					},
				},
			},
			expectedKubeDNS:               true,
			expectedCoreDNS:               false,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           false,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
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
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
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
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                true,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
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
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               true,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
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
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         true,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
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
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
		},
		// Scheduled Maintenance Scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin:   NetworkPluginAzure,
						LoadBalancerSku: api.StandardLoadBalancerSku,
						Addons: []api.KubernetesAddon{
							{
								Name:    ScheduledMaintenanceAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  true,
		},
		// PodSecurityPolicy scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						EnablePodSecurityPolicy: to.BoolPtr(true),
					},
				},
			},
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     true,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
		},
		// non-Managed Disk scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.7.10",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
					},
				},
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						StorageProfile: api.StorageAccount,
					},
				},
			},
			expectedKubeDNS:               true,
			expectedCoreDNS:               false,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           false,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   false,
			expectedUnmanagedStorageClass: true,
			expectedScheduledMaintenance:  false,
		},
		// Azure Stack Managed Disk scenario
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
						StorageProfile: api.ManagedDisks,
					},
				},
				CustomCloudProfile: mockAzureStackProperties.CustomCloudProfile,
			},
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
		},
		// Azure Stack non-Managed Disk scenario
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
						StorageProfile: api.StorageAccount,
					},
				},
				CustomCloudProfile: mockAzureStackProperties.CustomCloudProfile,
			},
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     false,
			expectedManagedStorageClass:   false,
			expectedUnmanagedStorageClass: true,
			expectedScheduledMaintenance:  false,
		},
		// 1.15.0-beta.1 scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.15.0-beta.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
					},
				},
			},
			expectedKubeDNS:               false,
			expectedCoreDNS:               true,
			expectedKubeProxy:             true,
			expectedCilium:                false,
			expectedFlannel:               false,
			expectedAADAdminGroup:         false,
			expectedAzureCloudProvider:    true,
			expectedAuditPolicy:           true,
			expectedPodSecurityPolicy:     true,
			expectedManagedStorageClass:   true,
			expectedUnmanagedStorageClass: false,
			expectedScheduledMaintenance:  false,
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
			case "pod-security-policy.yaml":
				if c.expectedPodSecurityPolicy != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "PodSecurityPolicy", c.expectedPodSecurityPolicy)
				}
			case "azure-storage-classes.yaml":
				if strings.Contains(componentFileSpec.sourceFile, "unmanaged-azure-storage") {
					if c.expectedUnmanagedStorageClass != componentFileSpec.isEnabled {
						t.Fatalf("Expected %s to be %t", componentFileSpec.sourceFile, c.expectedUnmanagedStorageClass)
					}
					if c.p.CustomCloudProfile != nil {
						if !strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
							t.Fatalf("Expected an Azure Stack-specific unmanaged disk spec, got %s instead", componentFileSpec.sourceFile)
						}
					} else {
						if strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
							t.Fatalf("Got an unexpected Azure Stack-specific unmanaged disk spec in a non-Azure Stack cluster configuration")
						}
					}
				} else {
					if c.expectedManagedStorageClass != componentFileSpec.isEnabled {
						t.Fatalf("Expected %s to be %t", componentFileSpec.sourceFile, c.expectedManagedStorageClass)
					}
					if c.p.CustomCloudProfile != nil {
						if !strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
							t.Fatalf("Expected an Azure Stack-specific Managed disk spec, got %s instead", componentFileSpec.sourceFile)
						}
					} else {
						if strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
							t.Fatalf("Got an unexpected Azure Stack-specific Managed disk spec in a non-Azure Stack cluster configuration")
						}
					}
				}
			case "scheduled-maintenance-deployment.yaml":
				if c.expectedScheduledMaintenance != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", ScheduledMaintenanceAddonName, c.expectedScheduledMaintenance)
				}
			}
		}
	}
}

func TestKubernetesManifestSettingsInit(t *testing.T) {
	mockAzureStackProperties := api.GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	cases := []struct {
		p                             *api.Properties
		expectedKubeScheduler         bool
		expectedKubeControllerManager bool
		expectedKubeCCM               bool
		expectedKubeAPIServer         bool
		expectedKubeAddonManager      bool
	}{
		// Default scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig: map[string]string{},
					},
				},
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               false,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
		// CCM scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig:              map[string]string{},
						UseCloudControllerManager:    to.BoolPtr(true),
						CloudControllerManagerConfig: map[string]string{},
					},
				},
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               true,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
		// Azure Stack Scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig: map[string]string{},
					},
				},
				CustomCloudProfile: mockAzureStackProperties.CustomCloudProfile,
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               false,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
		// Custom data scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("foo")),
						},
						ControllerManagerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("bar")),
						},
						CloudControllerManagerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("baz")),
						},
						APIServerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("bam")),
						},
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               true,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
	}
	for _, c := range cases {
		componentFileSpecArray := kubernetesManifestSettingsInit(c.p)
		for _, componentFileSpec := range componentFileSpecArray {
			switch componentFileSpec.destinationFile {
			case "kube-scheduler.yaml":
				if c.expectedKubeScheduler != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-scheduler", c.expectedKubeScheduler)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.SchedulerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.SchedulerConfig["data"])
				}
			case "kube-controller-manager.yaml":
				if c.p.CustomCloudProfile != nil {
					if !strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
						t.Fatalf("Expected an Azure Stack-specific controller-manager spec, got %s instead", componentFileSpec.sourceFile)
					}
				} else {
					if strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
						t.Fatalf("Got an unexpected Azure Stack-specific controller-manager spec in a non-Azure Stack cluster configuration")
					}
				}
				if c.expectedKubeControllerManager != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-controller-manager", c.expectedKubeControllerManager)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig["data"])
				}
			case "cloud-controller-manager.yaml":
				if c.expectedKubeCCM != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "cloud-controller-manager", c.expectedKubeCCM)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig["data"])
				}
			case "kube-apiserver.yaml":
				if c.expectedKubeAPIServer != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-apiserver", c.expectedKubeAPIServer)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.APIServerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.APIServerConfig["data"])
				}
			case "kube-addon-manager.yaml":
				if c.expectedKubeAddonManager != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-addon-manager", c.expectedKubeAddonManager)
				}
				if componentFileSpec.base64Data != "" {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, "")
				}
			}
		}
	}
}
