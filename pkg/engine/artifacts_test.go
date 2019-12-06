// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestKubernetesAddonSettingsInit(t *testing.T) {
	mockAzureStackProperties := api.GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	cases := []struct {
		p                              *api.Properties
		expectedCoreDNS                bool
		expectedKubeProxy              bool
		expectedCilium                 bool
		expectedFlannel                bool
		expectedAADAdminGroup          bool
		expectedAzureCloudProvider     bool
		expectedAuditPolicy            bool
		expectedPodSecurityPolicy      bool
		expectedManagedStorageClass    bool
		expectedUnmanagedStorageClass  bool
		expectedScheduledMaintenance   bool
		expectedAzureCSIStorageClasses bool
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
			expectedCoreDNS:                false,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            false,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 true,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                true,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          true,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
								Name:    common.ScheduledMaintenanceAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   true,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      true,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                false,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            false,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    false,
			expectedUnmanagedStorageClass:  true,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    false,
			expectedUnmanagedStorageClass:  true,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
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
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      true,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
		},
		// CSI storage classes scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.13.0",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin:             NetworkPluginAzure,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						StorageProfile: api.StorageAccount,
					},
				},
			},
			expectedCoreDNS:                true,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    false,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: true,
		},
		// kube-dns addon enabled scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						NetworkPlugin: NetworkPluginAzure,
						Addons: []api.KubernetesAddon{
							{
								Name:    common.KubeDNSAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedCoreDNS:                false,
			expectedKubeProxy:              true,
			expectedCilium:                 false,
			expectedFlannel:                false,
			expectedAADAdminGroup:          false,
			expectedAzureCloudProvider:     true,
			expectedAuditPolicy:            true,
			expectedPodSecurityPolicy:      false,
			expectedManagedStorageClass:    true,
			expectedUnmanagedStorageClass:  false,
			expectedScheduledMaintenance:   false,
			expectedAzureCSIStorageClasses: false,
		},
	}

	for _, c := range cases {
		componentFileSpecArray := kubernetesAddonSettingsInit(c.p)
		for _, componentFileSpec := range componentFileSpecArray {
			switch componentFileSpec.destinationFile {
			case "coredns.yaml":
				if c.expectedCoreDNS != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", common.CoreDNSAddonName, c.expectedCoreDNS)
				}
			case "kube-proxy-daemonset.yaml":
				if c.expectedKubeProxy != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", common.KubeProxyAddonName, c.expectedKubeProxy)
				}
			case "cilium-daemonset.yaml":
				if c.expectedCilium != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", common.CiliumAddonName, c.expectedCilium)
				}
			case "flannel-daemonset.yaml":
				if c.expectedFlannel != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", common.FlannelAddonName, c.expectedFlannel)
				}
			case "aad-default-admin-group-rbac.yaml":
				if c.expectedAADAdminGroup != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", common.AADAdminGroupAddonName, c.expectedAADAdminGroup)
				}
			case "azure-cloud-provider-deployment.yaml":
				if c.expectedAzureCloudProvider != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", common.AzureCloudProviderAddonName, c.expectedAzureCloudProvider)
				}
			case "audit-policy.yaml":
				if c.expectedAuditPolicy != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", common.AuditPolicyAddonName, c.expectedAuditPolicy)
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
					t.Fatalf("Expected %s to be %t", common.ScheduledMaintenanceAddonName, c.expectedScheduledMaintenance)
				}
			case "azure-csi-storage-classes.yaml":
				if c.expectedAzureCSIStorageClasses != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", componentFileSpec.sourceFile, c.expectedAzureCSIStorageClasses)
				}
			}
		}
	}
}

func TestKubernetesContainerAddonSettingsInit(t *testing.T) {
	/*
		$ echo "Hello, World\!" | base64
		SGVsbG8sIFdvcmxkXCEK
	*/
	const base64Data = "SGVsbG8sIFdvcmxkXCEK"
	cases := []struct {
		name                           string
		p                              *api.Properties
		expectedHeapster               kubernetesComponentFileSpec
		expectedMetricsServer          kubernetesComponentFileSpec
		expectedTiller                 kubernetesComponentFileSpec
		expectedAADPodIdentity         kubernetesComponentFileSpec
		expectedACIConnector           kubernetesComponentFileSpec
		expectedAzureDiskCSIDriver     kubernetesComponentFileSpec
		expectedAzureFileCSIDriver     kubernetesComponentFileSpec
		expectedClusterAutoscaler      kubernetesComponentFileSpec
		expectedBlobFlexVolume         kubernetesComponentFileSpec
		expectedSMBFlexVolume          kubernetesComponentFileSpec
		expectedKeyVaultFlexVolume     kubernetesComponentFileSpec
		expectedDashboard              kubernetesComponentFileSpec
		expectedRescheduler            kubernetesComponentFileSpec
		expectedNvidia                 kubernetesComponentFileSpec
		expectedContainerMonitoring    kubernetesComponentFileSpec
		expectedIPMasqAgent            kubernetesComponentFileSpec
		expectedAzureCNINetworkMonitor kubernetesComponentFileSpec
		expectedDNSAutoscaler          kubernetesComponentFileSpec
		expectedCalico                 kubernetesComponentFileSpec
		expectedAzureNetworkPolicy     kubernetesComponentFileSpec
		expectedAzurePolicy            kubernetesComponentFileSpec
		expectedCloudNodeManager       kubernetesComponentFileSpec
		expectedNodeProblemDetector    kubernetesComponentFileSpec
		expectedKubeDNS                kubernetesComponentFileSpec
	}{
		{
			name: "addons with data",
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.16.1",
					KubernetesConfig: &api.KubernetesConfig{
						Addons: []api.KubernetesAddon{
							{
								Name: common.HeapsterAddonName,
								Data: base64Data,
							},
							{
								Name: common.MetricsServerAddonName,
								Data: base64Data,
							},
							{
								Name: common.TillerAddonName,
								Data: base64Data,
							},
							{
								Name: common.AADPodIdentityAddonName,
								Data: base64Data,
							},
							{
								Name: common.ACIConnectorAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureDiskCSIDriverAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureFileCSIDriverAddonName,
								Data: base64Data,
							},
							{
								Name: common.ClusterAutoscalerAddonName,
								Data: base64Data,
							},
							{
								Name: common.BlobfuseFlexVolumeAddonName,
								Data: base64Data,
							},
							{
								Name: common.SMBFlexVolumeAddonName,
								Data: base64Data,
							},
							{
								Name: common.KeyVaultFlexVolumeAddonName,
								Data: base64Data,
							},
							{
								Name: common.DashboardAddonName,
								Data: base64Data,
							},
							{
								Name: common.ReschedulerAddonName,
								Data: base64Data,
							},
							{
								Name: common.NVIDIADevicePluginAddonName,
								Data: base64Data,
							},
							{
								Name: common.ContainerMonitoringAddonName,
								Data: base64Data,
							},
							{
								Name: common.IPMASQAgentAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureCNINetworkMonitorAddonName,
								Data: base64Data,
							},
							{
								Name: common.DNSAutoscalerAddonName,
								Data: base64Data,
							},
							{
								Name: common.CalicoAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureNetworkPolicyAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzurePolicyAddonName,
								Data: base64Data,
							},
							{
								Name: common.CloudNodeManagerAddonName,
								Data: base64Data,
							},
							{
								Name: common.NodeProblemDetectorAddonName,
								Data: base64Data,
							},
							{
								Name: common.KubeDNSAddonName,
								Data: base64Data,
							},
						},
					},
				},
			},
			expectedHeapster: kubernetesComponentFileSpec{
				sourceFile:      heapsterAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: heapsterAddonDestinationFilename,
			},
			expectedMetricsServer: kubernetesComponentFileSpec{
				sourceFile:      metricsServerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: metricsServerAddonDestinationFilename,
			},
			expectedTiller: kubernetesComponentFileSpec{
				sourceFile:      tillerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: tillerAddonDestinationFilename,
			},
			expectedAADPodIdentity: kubernetesComponentFileSpec{
				sourceFile:      aadPodIdentityAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: aadPodIdentityAddonDestinationFilename,
			},
			expectedACIConnector: kubernetesComponentFileSpec{
				sourceFile:      aciConnectorAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: aciConnectorAddonDestinationFilename,
			},
			expectedAzureDiskCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureDiskCSIAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureDiskCSIAddonDestinationFilename,
			},
			expectedAzureFileCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureFileCSIAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureFileCSIAddonDestinationFilename,
			},
			expectedClusterAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      clusterAutoscalerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: clusterAutoscalerAddonDestinationFilename,
			},
			expectedBlobFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      blobfuseFlexVolumeAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: blobfuseFlexVolumeAddonDestinationFilename,
			},
			expectedSMBFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      smbFlexVolumeAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: smbFlexVolumeAddonDestinationFilename,
			},
			expectedKeyVaultFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      keyvaultFlexVolumeAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: keyvaultFlexVolumeAddonDestinationFilename,
			},
			expectedDashboard: kubernetesComponentFileSpec{
				sourceFile:      dashboardAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: dashboardAddonDestinationFilename,
			},
			expectedRescheduler: kubernetesComponentFileSpec{
				sourceFile:      reschedulerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: reschedulerAddonDestinationFilename,
			},
			expectedNvidia: kubernetesComponentFileSpec{
				sourceFile:      nvidiaAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: nvidiaAddonDestinationFilename,
			},
			expectedContainerMonitoring: kubernetesComponentFileSpec{
				sourceFile:      containerMonitoringAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: containerMonitoringAddonDestinationFilename,
			},
			expectedIPMasqAgent: kubernetesComponentFileSpec{
				sourceFile:      ipMasqAgentAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: ipMasqAgentAddonDestinationFilename,
			},
			expectedAzureCNINetworkMonitor: kubernetesComponentFileSpec{
				sourceFile:      azureCNINetworkMonitorAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureCNINetworkMonitorAddonDestinationFilename,
			},
			expectedDNSAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      dnsAutoscalerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: dnsAutoscalerAddonDestinationFilename,
			},
			expectedCalico: kubernetesComponentFileSpec{
				sourceFile:      calicoAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: calicoAddonDestinationFilename,
			},
			expectedAzureNetworkPolicy: kubernetesComponentFileSpec{
				sourceFile:      azureNetworkPolicyAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureNetworkPolicyAddonDestinationFilename,
			},
			expectedAzurePolicy: kubernetesComponentFileSpec{
				sourceFile:      azurePolicyAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azurePolicyAddonDestinationFilename,
			},
			expectedCloudNodeManager: kubernetesComponentFileSpec{
				sourceFile:      cloudNodeManagerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: cloudNodeManagerAddonDestinationFilename,
			},
			expectedNodeProblemDetector: kubernetesComponentFileSpec{
				sourceFile:      nodeProblemDetectorAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: nodeProblemDetectorAddonDestinationFilename,
			},
			expectedKubeDNS: kubernetesComponentFileSpec{
				sourceFile:      kubeDNSAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: kubeDNSAddonDestinationFilename,
			},
		},
		{
			name: "addons with no data",
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.16.1",
					KubernetesConfig: &api.KubernetesConfig{
						Addons: []api.KubernetesAddon{
							{
								Name: common.HeapsterAddonName,
							},
							{
								Name: common.MetricsServerAddonName,
							},
							{
								Name: common.TillerAddonName,
							},
							{
								Name: common.AADPodIdentityAddonName,
							},
							{
								Name: common.ACIConnectorAddonName,
							},
							{
								Name: common.AzureDiskCSIDriverAddonName,
							},
							{
								Name: common.AzureFileCSIDriverAddonName,
							},
							{
								Name: common.ClusterAutoscalerAddonName,
							},
							{
								Name: common.BlobfuseFlexVolumeAddonName,
							},
							{
								Name: common.SMBFlexVolumeAddonName,
							},
							{
								Name: common.KeyVaultFlexVolumeAddonName,
							},
							{
								Name: common.DashboardAddonName,
							},
							{
								Name: common.ReschedulerAddonName,
							},
							{
								Name: common.NVIDIADevicePluginAddonName,
							},
							{
								Name: common.ContainerMonitoringAddonName,
							},
							{
								Name: common.IPMASQAgentAddonName,
							},
							{
								Name: common.AzureCNINetworkMonitorAddonName,
							},
							{
								Name: common.DNSAutoscalerAddonName,
							},
							{
								Name: common.CalicoAddonName,
							},
							{
								Name: common.AzureNetworkPolicyAddonName,
							},
							{
								Name: common.AzurePolicyAddonName,
							},
							{
								Name: common.CloudNodeManagerAddonName,
							},
							{
								Name: common.NodeProblemDetectorAddonName,
							},
							{
								Name: common.KubeDNSAddonName,
							},
						},
					},
				},
			},
			expectedHeapster: kubernetesComponentFileSpec{
				sourceFile:      heapsterAddonSourceFilename,
				base64Data:      "",
				destinationFile: heapsterAddonDestinationFilename,
			},
			expectedMetricsServer: kubernetesComponentFileSpec{
				sourceFile:      metricsServerAddonSourceFilename,
				base64Data:      "",
				destinationFile: metricsServerAddonDestinationFilename,
			},
			expectedTiller: kubernetesComponentFileSpec{
				sourceFile:      tillerAddonSourceFilename,
				base64Data:      "",
				destinationFile: tillerAddonDestinationFilename,
			},
			expectedAADPodIdentity: kubernetesComponentFileSpec{
				sourceFile:      aadPodIdentityAddonSourceFilename,
				base64Data:      "",
				destinationFile: aadPodIdentityAddonDestinationFilename,
			},
			expectedACIConnector: kubernetesComponentFileSpec{
				sourceFile:      aciConnectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: aciConnectorAddonDestinationFilename,
			},
			expectedAzureDiskCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureDiskCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureDiskCSIAddonDestinationFilename,
			},
			expectedAzureFileCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureFileCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureFileCSIAddonDestinationFilename,
			},
			expectedClusterAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      clusterAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: clusterAutoscalerAddonDestinationFilename,
			},
			expectedBlobFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      blobfuseFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: blobfuseFlexVolumeAddonDestinationFilename,
			},
			expectedSMBFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      smbFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: smbFlexVolumeAddonDestinationFilename,
			},
			expectedKeyVaultFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      keyvaultFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: keyvaultFlexVolumeAddonDestinationFilename,
			},
			expectedDashboard: kubernetesComponentFileSpec{
				sourceFile:      dashboardAddonSourceFilename,
				base64Data:      "",
				destinationFile: dashboardAddonDestinationFilename,
			},
			expectedRescheduler: kubernetesComponentFileSpec{
				sourceFile:      reschedulerAddonSourceFilename,
				base64Data:      "",
				destinationFile: reschedulerAddonDestinationFilename,
			},
			expectedNvidia: kubernetesComponentFileSpec{
				sourceFile:      nvidiaAddonSourceFilename,
				base64Data:      "",
				destinationFile: nvidiaAddonDestinationFilename,
			},
			expectedContainerMonitoring: kubernetesComponentFileSpec{
				sourceFile:      containerMonitoringAddonSourceFilename,
				base64Data:      "",
				destinationFile: containerMonitoringAddonDestinationFilename,
			},
			expectedIPMasqAgent: kubernetesComponentFileSpec{
				sourceFile:      ipMasqAgentAddonSourceFilename,
				base64Data:      "",
				destinationFile: ipMasqAgentAddonDestinationFilename,
			},
			expectedAzureCNINetworkMonitor: kubernetesComponentFileSpec{
				sourceFile:      azureCNINetworkMonitorAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureCNINetworkMonitorAddonDestinationFilename,
			},
			expectedDNSAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      dnsAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: dnsAutoscalerAddonDestinationFilename,
			},
			expectedCalico: kubernetesComponentFileSpec{
				sourceFile:      calicoAddonSourceFilename,
				base64Data:      "",
				destinationFile: calicoAddonDestinationFilename,
			},
			expectedAzureNetworkPolicy: kubernetesComponentFileSpec{
				sourceFile:      azureNetworkPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureNetworkPolicyAddonDestinationFilename,
			},
			expectedAzurePolicy: kubernetesComponentFileSpec{
				sourceFile:      azurePolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azurePolicyAddonDestinationFilename,
			},
			expectedCloudNodeManager: kubernetesComponentFileSpec{
				sourceFile:      cloudNodeManagerAddonSourceFilename,
				base64Data:      "",
				destinationFile: cloudNodeManagerAddonDestinationFilename,
			},
			expectedNodeProblemDetector: kubernetesComponentFileSpec{
				sourceFile:      nodeProblemDetectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: nodeProblemDetectorAddonDestinationFilename,
			},
			expectedKubeDNS: kubernetesComponentFileSpec{
				sourceFile:      kubeDNSAddonSourceFilename,
				base64Data:      "",
				destinationFile: kubeDNSAddonDestinationFilename,
			},
		},
		{
			name: "no addons in ContainerService object",
			p:    &api.Properties{},
			expectedHeapster: kubernetesComponentFileSpec{
				sourceFile:      heapsterAddonSourceFilename,
				base64Data:      "",
				destinationFile: heapsterAddonDestinationFilename,
			},
			expectedMetricsServer: kubernetesComponentFileSpec{
				sourceFile:      metricsServerAddonSourceFilename,
				base64Data:      "",
				destinationFile: metricsServerAddonDestinationFilename,
			},
			expectedTiller: kubernetesComponentFileSpec{
				sourceFile:      tillerAddonSourceFilename,
				base64Data:      "",
				destinationFile: tillerAddonDestinationFilename,
			},
			expectedAADPodIdentity: kubernetesComponentFileSpec{
				sourceFile:      aadPodIdentityAddonSourceFilename,
				base64Data:      "",
				destinationFile: aadPodIdentityAddonDestinationFilename,
			},
			expectedACIConnector: kubernetesComponentFileSpec{
				sourceFile:      aciConnectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: aciConnectorAddonDestinationFilename,
			},
			expectedAzureDiskCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureDiskCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureDiskCSIAddonDestinationFilename,
			},
			expectedAzureFileCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureFileCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureFileCSIAddonDestinationFilename,
			},
			expectedClusterAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      clusterAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: clusterAutoscalerAddonDestinationFilename,
			},
			expectedBlobFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      blobfuseFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: blobfuseFlexVolumeAddonDestinationFilename,
			},
			expectedSMBFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      smbFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: smbFlexVolumeAddonDestinationFilename,
			},
			expectedKeyVaultFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      keyvaultFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: keyvaultFlexVolumeAddonDestinationFilename,
			},
			expectedDashboard: kubernetesComponentFileSpec{
				sourceFile:      dashboardAddonSourceFilename,
				base64Data:      "",
				destinationFile: dashboardAddonDestinationFilename,
			},
			expectedRescheduler: kubernetesComponentFileSpec{
				sourceFile:      reschedulerAddonSourceFilename,
				base64Data:      "",
				destinationFile: reschedulerAddonDestinationFilename,
			},
			expectedNvidia: kubernetesComponentFileSpec{
				sourceFile:      nvidiaAddonSourceFilename,
				base64Data:      "",
				destinationFile: nvidiaAddonDestinationFilename,
			},
			expectedContainerMonitoring: kubernetesComponentFileSpec{
				sourceFile:      containerMonitoringAddonSourceFilename,
				base64Data:      "",
				destinationFile: containerMonitoringAddonDestinationFilename,
			},
			expectedIPMasqAgent: kubernetesComponentFileSpec{
				sourceFile:      ipMasqAgentAddonSourceFilename,
				base64Data:      "",
				destinationFile: ipMasqAgentAddonDestinationFilename,
			},
			expectedAzureCNINetworkMonitor: kubernetesComponentFileSpec{
				sourceFile:      azureCNINetworkMonitorAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureCNINetworkMonitorAddonDestinationFilename,
			},
			expectedDNSAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      dnsAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: dnsAutoscalerAddonDestinationFilename,
			},
			expectedCalico: kubernetesComponentFileSpec{
				sourceFile:      calicoAddonSourceFilename,
				base64Data:      "",
				destinationFile: calicoAddonDestinationFilename,
			},
			expectedAzureNetworkPolicy: kubernetesComponentFileSpec{
				sourceFile:      azureNetworkPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureNetworkPolicyAddonDestinationFilename,
			},
			expectedAzurePolicy: kubernetesComponentFileSpec{
				sourceFile:      azurePolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azurePolicyAddonDestinationFilename,
			},
			expectedCloudNodeManager: kubernetesComponentFileSpec{
				sourceFile:      cloudNodeManagerAddonSourceFilename,
				base64Data:      "",
				destinationFile: cloudNodeManagerAddonDestinationFilename,
			},
			expectedNodeProblemDetector: kubernetesComponentFileSpec{
				sourceFile:      nodeProblemDetectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: nodeProblemDetectorAddonDestinationFilename,
			},
			expectedKubeDNS: kubernetesComponentFileSpec{
				sourceFile:      kubeDNSAddonSourceFilename,
				base64Data:      "",
				destinationFile: kubeDNSAddonDestinationFilename,
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			componentFileSpec := kubernetesContainerAddonSettingsInit(c.p)
			for addon := range componentFileSpec {
				switch addon {
				case common.HeapsterAddonName:
					if c.expectedHeapster.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedHeapster.sourceFile)
					}
					if c.expectedHeapster.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedHeapster.base64Data)
					}
					if c.expectedHeapster.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedHeapster.destinationFile)
					}
				case common.MetricsServerAddonName:
					if c.expectedMetricsServer.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedMetricsServer.sourceFile)
					}
					if c.expectedMetricsServer.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedMetricsServer.base64Data)
					}
					if c.expectedMetricsServer.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedMetricsServer.destinationFile)
					}
				case common.TillerAddonName:
					if c.expectedTiller.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedTiller.sourceFile)
					}
					if c.expectedTiller.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedTiller.base64Data)
					}
					if c.expectedTiller.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedTiller.destinationFile)
					}
				case common.AADPodIdentityAddonName:
					if c.expectedAADPodIdentity.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAADPodIdentity.sourceFile)
					}
					if c.expectedAADPodIdentity.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAADPodIdentity.base64Data)
					}
					if c.expectedAADPodIdentity.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAADPodIdentity.destinationFile)
					}
				case common.ACIConnectorAddonName:
					if c.expectedACIConnector.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedACIConnector.sourceFile)
					}
					if c.expectedACIConnector.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedACIConnector.base64Data)
					}
					if c.expectedACIConnector.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedACIConnector.destinationFile)
					}
				case common.AzureDiskCSIDriverAddonName:
					if c.expectedAzureDiskCSIDriver.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureDiskCSIDriver.sourceFile)
					}
					if c.expectedAzureDiskCSIDriver.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureDiskCSIDriver.base64Data)
					}
					if c.expectedAzureDiskCSIDriver.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureDiskCSIDriver.destinationFile)
					}
				case common.AzureFileCSIDriverAddonName:
					if c.expectedAzureFileCSIDriver.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureFileCSIDriver.sourceFile)
					}
					if c.expectedAzureFileCSIDriver.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureFileCSIDriver.base64Data)
					}
					if c.expectedAzureFileCSIDriver.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureFileCSIDriver.destinationFile)
					}
				case common.ClusterAutoscalerAddonName:
					if c.expectedClusterAutoscaler.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedClusterAutoscaler.sourceFile)
					}
					if c.expectedClusterAutoscaler.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedClusterAutoscaler.base64Data)
					}
					if c.expectedClusterAutoscaler.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedClusterAutoscaler.destinationFile)
					}
				case common.BlobfuseFlexVolumeAddonName:
					if c.expectedBlobFlexVolume.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedBlobFlexVolume.sourceFile)
					}
					if c.expectedBlobFlexVolume.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedBlobFlexVolume.base64Data)
					}
					if c.expectedBlobFlexVolume.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedBlobFlexVolume.destinationFile)
					}
				case common.SMBFlexVolumeAddonName:
					if c.expectedSMBFlexVolume.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedSMBFlexVolume.sourceFile)
					}
					if c.expectedSMBFlexVolume.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedSMBFlexVolume.base64Data)
					}
					if c.expectedSMBFlexVolume.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedSMBFlexVolume.destinationFile)
					}
				case common.KeyVaultFlexVolumeAddonName:
					if c.expectedKeyVaultFlexVolume.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedKeyVaultFlexVolume.sourceFile)
					}
					if c.expectedKeyVaultFlexVolume.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedKeyVaultFlexVolume.base64Data)
					}
					if c.expectedKeyVaultFlexVolume.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedKeyVaultFlexVolume.destinationFile)
					}
				case common.DashboardAddonName:
					if c.expectedDashboard.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedDashboard.sourceFile)
					}
					if c.expectedDashboard.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedDashboard.base64Data)
					}
					if c.expectedDashboard.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedDashboard.destinationFile)
					}
				case common.ReschedulerAddonName:
					if c.expectedRescheduler.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedRescheduler.sourceFile)
					}
					if c.expectedRescheduler.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedRescheduler.base64Data)
					}
					if c.expectedRescheduler.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedRescheduler.destinationFile)
					}
				case common.NVIDIADevicePluginAddonName:
					if c.expectedNvidia.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedNvidia.sourceFile)
					}
					if c.expectedNvidia.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedNvidia.base64Data)
					}
					if c.expectedNvidia.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedNvidia.destinationFile)
					}
				case common.ContainerMonitoringAddonName:
					if c.expectedContainerMonitoring.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedContainerMonitoring.sourceFile)
					}
					if c.expectedContainerMonitoring.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedContainerMonitoring.base64Data)
					}
					if c.expectedContainerMonitoring.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedContainerMonitoring.destinationFile)
					}
				case common.IPMASQAgentAddonName:
					if c.expectedIPMasqAgent.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedIPMasqAgent.sourceFile)
					}
					if c.expectedIPMasqAgent.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedIPMasqAgent.base64Data)
					}
					if c.expectedIPMasqAgent.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedIPMasqAgent.destinationFile)
					}
				case common.AzureCNINetworkMonitorAddonName:
					if c.expectedAzureCNINetworkMonitor.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureCNINetworkMonitor.sourceFile)
					}
					if c.expectedAzureCNINetworkMonitor.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureCNINetworkMonitor.base64Data)
					}
					if c.expectedAzureCNINetworkMonitor.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureCNINetworkMonitor.destinationFile)
					}
				case common.DNSAutoscalerAddonName:
					if c.expectedDNSAutoscaler.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedDNSAutoscaler.sourceFile)
					}
					if c.expectedDNSAutoscaler.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedDNSAutoscaler.base64Data)
					}
					if c.expectedDNSAutoscaler.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedDNSAutoscaler.destinationFile)
					}
				case common.CalicoAddonName:
					if c.expectedCalico.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedCalico.sourceFile)
					}
					if c.expectedCalico.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedCalico.base64Data)
					}
					if c.expectedCalico.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedCalico.destinationFile)
					}
				case common.AzureNetworkPolicyAddonName:
					if c.expectedAzureNetworkPolicy.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureNetworkPolicy.sourceFile)
					}
					if c.expectedAzureNetworkPolicy.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureNetworkPolicy.base64Data)
					}
					if c.expectedAzureNetworkPolicy.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureNetworkPolicy.destinationFile)
					}
				case common.AzurePolicyAddonName:
					if c.expectedAzurePolicy.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzurePolicy.sourceFile)
					}
					if c.expectedAzurePolicy.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzurePolicy.base64Data)
					}
					if c.expectedAzurePolicy.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzurePolicy.destinationFile)
					}
				case common.CloudNodeManagerAddonName:
					if c.expectedCloudNodeManager.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedCloudNodeManager.sourceFile)
					}
					if c.expectedCloudNodeManager.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedCloudNodeManager.base64Data)
					}
					if c.expectedCloudNodeManager.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedCloudNodeManager.destinationFile)
					}
				case common.NodeProblemDetectorAddonName:
					if c.expectedNodeProblemDetector.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedNodeProblemDetector.sourceFile)
					}
					if c.expectedNodeProblemDetector.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedNodeProblemDetector.base64Data)
					}
					if c.expectedNodeProblemDetector.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedNodeProblemDetector.destinationFile)
					}
				case common.KubeDNSAddonName:
					if c.expectedKubeDNS.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedKubeDNS.sourceFile)
					}
					if c.expectedKubeDNS.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedKubeDNS.base64Data)
					}
					if c.expectedKubeDNS.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedKubeDNS.destinationFile)
					}
				}
			}
		})
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
