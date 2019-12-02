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

func TestKubernetesAddonSettingsInit(t *testing.T) {
	mockAzureStackProperties := api.GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	cases := []struct {
		p                              *api.Properties
		expectedKubeDNS                bool
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
			expectedKubeDNS:                true,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
								Name:    ScheduledMaintenanceAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                true,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
			expectedKubeDNS:                false,
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
			case "azure-csi-storage-classes.yaml":
				if c.expectedAzureCSIStorageClasses != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", componentFileSpec.sourceFile, c.expectedAzureCSIStorageClasses)
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
