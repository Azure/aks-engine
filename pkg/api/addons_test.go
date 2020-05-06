// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestAppendAddonIfNotPresent(t *testing.T) {
	names := []string{"AddonA", "AddonB", "AddonC", "AddonD", "AddonE"}
	addons := []KubernetesAddon{}

	for i, name := range names {
		addon := KubernetesAddon{Name: name}
		addons = appendAddonIfNotPresent(addons, addon)
		if len(addons) != i+1 {
			t.Errorf("incorrect length by appendAddonIfNotPresent, expect: '%d', actual: '%d'", i+1, len(addons))
		}
	}

	for _, name := range names {
		addon := KubernetesAddon{Name: name}
		addons = appendAddonIfNotPresent(addons, addon)
		if len(addons) != len(names) {
			t.Errorf("incorrect length by appendAddonIfNotPresent, expect: '%d', actual: '%d'", len(names), len(addons))
		}
	}
}

func TestGetAddonsIndexByName(t *testing.T) {
	includedNames := []string{"AddonA", "AddonB", "AddonC", "AddonD", "AddonE"}
	notIncludedNames := []string{"AddonF", "AddonG", "AddonH"}
	addons := []KubernetesAddon{}

	for _, name := range includedNames {
		addons = append(addons, KubernetesAddon{Name: name})
	}

	for i, addon := range addons {
		j := getAddonsIndexByName(addons, addon.Name)
		if j != i {
			t.Errorf("incorrect index by getAddonsIndexByName, expect: '%d', actual: '%d'", i, j)
		}
	}

	for _, name := range notIncludedNames {
		j := getAddonsIndexByName(addons, name)
		if j != -1 {
			t.Errorf("incorrect index by getAddonsIndexByName, expect: '%d', actual: '%d'", -1, j)
		}
	}
}

func TestPodSecurityPolicyConfigUpgrade(t *testing.T) {
	mockCS := getMockBaseContainerService("1.18.1")
	o := mockCS.Properties.OrchestratorProfile

	isUpgrade := true
	base64DataPSP := "cHNwQ3VzdG9tRGF0YQ=="
	o.OrchestratorType = Kubernetes
	o.KubernetesConfig.Addons = []KubernetesAddon{
		{
			Name:    common.PodSecurityPolicyAddonName,
			Enabled: to.BoolPtr(true),
			Data:    base64DataPSP,
		},
	}

	mockCS.setAddonsConfig(isUpgrade)

	i := getAddonsIndexByName(o.KubernetesConfig.Addons, common.PodSecurityPolicyAddonName)
	if i < 0 {
		t.Errorf("expected a positive index for the addon %s, instead got %d from getAddonsIndexByName", common.PodSecurityPolicyAddonName, i)
	}

	if o.KubernetesConfig.Addons[i].Name != common.PodSecurityPolicyAddonName {
		t.Errorf("expected addon %s name to be present, instead got %s", common.PodSecurityPolicyAddonName, o.KubernetesConfig.Addons[i].Name)
	}

	if o.KubernetesConfig.Addons[i].Data != base64DataPSP {
		t.Errorf("expected %s data to be present, instead got %s", base64DataPSP, o.KubernetesConfig.Addons[i].Data)
	}
}

func TestDisabledAddons(t *testing.T) {
	defaultAddon := KubernetesAddon{
		Name:    "mockAddon",
		Enabled: to.BoolPtr(false),
		Containers: []KubernetesContainerSpec{
			{
				Name:           "mockAddon",
				CPURequests:    "50m",
				MemoryRequests: "50Mi",
				CPULimits:      "50m",
				MemoryLimits:   "250Mi",
				Image:          "mockImage",
			},
		},
		Config: map[string]string{
			"fake-config-1": "someValue1",
			"fake-config-2": "someValue2",
		},
	}

	cases := []struct {
		name                string
		myAddon             KubernetesAddon
		isUpgrade           bool
		expectedResultAddon KubernetesAddon
	}{
		{
			name: "default addon enabled",
			myAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(true),
			},
			isUpgrade: false,
			expectedResultAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name:           "mockAddon",
						CPURequests:    "50m",
						MemoryRequests: "50Mi",
						CPULimits:      "50m",
						MemoryLimits:   "250Mi",
						Image:          "mockImage",
					},
				},
				Config: map[string]string{
					"fake-config-1": "someValue1",
					"fake-config-2": "someValue2",
				},
			},
		},
		{
			name: "addon disabled, isUpgrade=false",
			myAddon: KubernetesAddon{
				Name: "mockAddon",
			},
			isUpgrade: false,
			expectedResultAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
		},
		{
			name: "addon disabled, isUpgrade=true",
			myAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
			isUpgrade: true,
			expectedResultAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := assignDefaultAddonVals(c.myAddon, defaultAddon, c.isUpgrade)
			if !reflect.DeepEqual(result, c.expectedResultAddon) {
				t.Fatalf("expected result addon %v to be equal to %v", result, c.expectedResultAddon)
			}
		})
	}

}

func TestSetAddonsConfig(t *testing.T) {
	specConfig := AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	azureStackCloudSpec := AzureEnvironmentSpecConfig{
		CloudName: "AzureStackCloud",
		KubernetesSpecConfig: KubernetesSpecConfig{
			KubernetesImageBase:              "KubernetesImageBase",
			MCRKubernetesImageBase:           "MCRKubernetesImageBase",
			TillerImageBase:                  "TillerImageBase",
			ACIConnectorImageBase:            "ACIConnectorImageBase",
			NVIDIAImageBase:                  "NVIDIAImageBase",
			AzureCNIImageBase:                "AzureCNIImageBase",
			CalicoImageBase:                  "CalicoImageBase",
			EtcdDownloadURLBase:              "EtcdDownloadURLBase",
			KubeBinariesSASURLBase:           "KubeBinariesSASURLBase",
			WindowsTelemetryGUID:             "WindowsTelemetryGUID",
			CNIPluginsDownloadURL:            "CNIPluginsDownloadURL",
			VnetCNILinuxPluginsDownloadURL:   "VnetCNILinuxPluginsDownloadURL",
			VnetCNIWindowsPluginsDownloadURL: "VnetCNIWindowsPluginsDownloadURL",
			ContainerdDownloadURLBase:        "ContainerdDownloadURLBase",
			CSIProxyDownloadURL:              "CSIProxyDownloadURL",
		},
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
		},
	}
	AzureCloudSpecEnvMap[AzureStackCloud] = azureStackCloudSpec
	k8sComponentsByVersionMap := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	tests := []struct {
		name           string
		cs             *ContainerService
		isUpgrade      bool
		expectedAddons []KubernetesAddon
	}{
		{
			name: "default addons",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
						},
					},
				},
			},
			isUpgrade:      false,
			expectedAddons: getDefaultAddons("1.15.4", "", common.KubernetesImageBaseTypeMCR),
		},
		{
			name: "default addons w/ custom MCR KubernetesImageBase",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBase:     "custommcr.microsoft.com/",
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
						},
					},
				},
			},
			isUpgrade:      false,
			expectedAddons: getDefaultAddons("1.15.4", "custommcr.microsoft.com/", common.KubernetesImageBaseTypeMCR),
		},
		{
			name: "default addons w/ custom GCR KubernetesImageBase",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBase:     "customgcr.example.com/",
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
						},
					},
				},
			},
			isUpgrade:      false,
			expectedAddons: getDefaultAddons("1.15.4", "customgcr.example.com/", common.KubernetesImageBaseTypeGCR),
		},
		{
			name: "tiller addon is enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.TillerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.TillerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.TillerAddonName,
							CPURequests:    "50m",
							MemoryRequests: "150Mi",
							CPULimits:      "50m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.TillerImageBase + k8sComponentsByVersionMap["1.15.4"][common.TillerAddonName],
						},
					},
					Config: map[string]string{
						"max-history": strconv.Itoa(0),
					},
				},
			}, "1.15.4"),
		},
		{
			name: "ACI Connector addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ACIConnectorAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ACIConnectorAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"region":   "westus",
						"nodeName": "aci-connector",
						"os":       "Linux",
						"taint":    "azure.com/aci",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ACIConnectorAddonName,
							CPURequests:    "50m",
							MemoryRequests: "150Mi",
							CPULimits:      "50m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.ACIConnectorImageBase + k8sComponentsByVersionMap["1.15.4"][common.ACIConnectorAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "1m",
						"expendable-pods-priority-cutoff":       "-10",
						"ignore-daemonsets-utilization":         "false",
						"ignore-mirror-pods-utilization":        "false",
						"max-autoprovisioned-node-group-count":  "15",
						"max-empty-bulk-delete":                 "10",
						"max-failing-time":                      "15m0s",
						"max-graceful-termination-sec":          "600",
						"max-inactivity":                        "10m0s",
						"max-node-provision-time":               "15m0s",
						"max-nodes-total":                       "0",
						"max-total-unready-percentage":          "45",
						"memory-total":                          "0:6400000",
						"min-replica-count":                     "0",
						"new-pod-scale-up-delay":                "0s",
						"node-autoprovisioning-enabled":         "false",
						"ok-total-unready-count":                "3",
						"scale-down-candidates-pool-min-count":  "50",
						"scale-down-candidates-pool-ratio":      "0.1",
						"scale-down-delay-after-add":            "10m0s",
						"scale-down-delay-after-delete":         "1m",
						"scale-down-delay-after-failure":        "3m0s",
						"scale-down-enabled":                    "true",
						"scale-down-non-empty-candidates-count": "30",
						"scale-down-unneeded-time":              "10m0s",
						"scale-down-unready-time":               "20m0s",
						"scale-down-utilization-threshold":      "0.5",
						"skip-nodes-with-local-storage":         "false",
						"skip-nodes-with-system-pods":           "true",
						"stderrthreshold":                       "2",
						"unremovable-node-recheck-timeout":      "5m0s",
						"v":                                     "3",
						"write-status-configmap":                "true",
						"balance-similar-node-groups":           "true",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "1",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"min-nodes": "1",
										"max-nodes": "3",
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "1m",
						"expendable-pods-priority-cutoff":       "-10",
						"ignore-daemonsets-utilization":         "false",
						"ignore-mirror-pods-utilization":        "false",
						"max-autoprovisioned-node-group-count":  "15",
						"max-empty-bulk-delete":                 "10",
						"max-failing-time":                      "15m0s",
						"max-graceful-termination-sec":          "600",
						"max-inactivity":                        "10m0s",
						"max-node-provision-time":               "15m0s",
						"max-nodes-total":                       "0",
						"max-total-unready-percentage":          "45",
						"memory-total":                          "0:6400000",
						"min-replica-count":                     "0",
						"new-pod-scale-up-delay":                "0s",
						"node-autoprovisioning-enabled":         "false",
						"ok-total-unready-count":                "3",
						"scale-down-candidates-pool-min-count":  "50",
						"scale-down-candidates-pool-ratio":      "0.1",
						"scale-down-delay-after-add":            "10m0s",
						"scale-down-delay-after-delete":         "1m",
						"scale-down-delay-after-failure":        "3m0s",
						"scale-down-enabled":                    "true",
						"scale-down-non-empty-candidates-count": "30",
						"scale-down-unneeded-time":              "10m0s",
						"scale-down-unready-time":               "20m0s",
						"scale-down-utilization-threshold":      "0.5",
						"skip-nodes-with-local-storage":         "false",
						"skip-nodes-with-system-pods":           "true",
						"stderrthreshold":                       "2",
						"unremovable-node-recheck-timeout":      "5m0s",
						"v":                                     "3",
						"write-status-configmap":                "true",
						"balance-similar-node-groups":           "true",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "3",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled on cluster with multiple pools",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 5,
						},
						{
							Name:  "pool2",
							Count: 10,
						},
						{
							Name:  "pool3",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "1m",
						"expendable-pods-priority-cutoff":       "-10",
						"ignore-daemonsets-utilization":         "false",
						"ignore-mirror-pods-utilization":        "false",
						"max-autoprovisioned-node-group-count":  "15",
						"max-empty-bulk-delete":                 "10",
						"max-failing-time":                      "15m0s",
						"max-graceful-termination-sec":          "600",
						"max-inactivity":                        "10m0s",
						"max-node-provision-time":               "15m0s",
						"max-nodes-total":                       "0",
						"max-total-unready-percentage":          "45",
						"memory-total":                          "0:6400000",
						"min-replica-count":                     "0",
						"new-pod-scale-up-delay":                "0s",
						"node-autoprovisioning-enabled":         "false",
						"ok-total-unready-count":                "3",
						"scale-down-candidates-pool-min-count":  "50",
						"scale-down-candidates-pool-ratio":      "0.1",
						"scale-down-delay-after-add":            "10m0s",
						"scale-down-delay-after-delete":         "1m",
						"scale-down-delay-after-failure":        "3m0s",
						"scale-down-enabled":                    "true",
						"scale-down-non-empty-candidates-count": "30",
						"scale-down-unneeded-time":              "10m0s",
						"scale-down-unready-time":               "20m0s",
						"scale-down-utilization-threshold":      "0.5",
						"skip-nodes-with-local-storage":         "false",
						"skip-nodes-with-system-pods":           "true",
						"stderrthreshold":                       "2",
						"unremovable-node-recheck-timeout":      "5m0s",
						"v":                                     "3",
						"write-status-configmap":                "true",
						"balance-similar-node-groups":           "true",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "5",
								"max-nodes": "5",
							},
						},
						{
							Name: "pool2",
							Config: map[string]string{
								"min-nodes": "10",
								"max-nodes": "10",
							},
						},
						{
							Name: "pool3",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "1",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled on cluster with multiple pools - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"min-nodes": "5",
										"max-nodes": "100",
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 5,
						},
						{
							Name:  "pool2",
							Count: 10,
						},
						{
							Name:  "pool3",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "1m",
						"expendable-pods-priority-cutoff":       "-10",
						"ignore-daemonsets-utilization":         "false",
						"ignore-mirror-pods-utilization":        "false",
						"max-autoprovisioned-node-group-count":  "15",
						"max-empty-bulk-delete":                 "10",
						"max-failing-time":                      "15m0s",
						"max-graceful-termination-sec":          "600",
						"max-inactivity":                        "10m0s",
						"max-node-provision-time":               "15m0s",
						"max-nodes-total":                       "0",
						"max-total-unready-percentage":          "45",
						"memory-total":                          "0:6400000",
						"min-replica-count":                     "0",
						"new-pod-scale-up-delay":                "0s",
						"node-autoprovisioning-enabled":         "false",
						"ok-total-unready-count":                "3",
						"scale-down-candidates-pool-min-count":  "50",
						"scale-down-candidates-pool-ratio":      "0.1",
						"scale-down-delay-after-add":            "10m0s",
						"scale-down-delay-after-delete":         "1m",
						"scale-down-delay-after-failure":        "3m0s",
						"scale-down-enabled":                    "true",
						"scale-down-non-empty-candidates-count": "30",
						"scale-down-unneeded-time":              "10m0s",
						"scale-down-unready-time":               "20m0s",
						"scale-down-utilization-threshold":      "0.5",
						"skip-nodes-with-local-storage":         "false",
						"skip-nodes-with-system-pods":           "true",
						"stderrthreshold":                       "2",
						"unremovable-node-recheck-timeout":      "5m0s",
						"v":                                     "3",
						"write-status-configmap":                "true",
						"balance-similar-node-groups":           "true",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "5",
								"max-nodes": "100",
							},
						},
						{
							Name: "pool2",
							Config: map[string]string{
								"min-nodes": "10",
								"max-nodes": "10",
							},
						},
						{
							Name: "pool3",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "1",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled with configuration",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"scan-interval":                         "30s",
										"expendable-pods-priority-cutoff":       "-20",
										"ignore-daemonsets-utilization":         "true",
										"ignore-mirror-pods-utilization":        "true",
										"max-autoprovisioned-node-group-count":  "25",
										"max-empty-bulk-delete":                 "12",
										"max-failing-time":                      "15m10s",
										"max-graceful-termination-sec":          "700",
										"max-inactivity":                        "11m0s",
										"max-node-provision-time":               "16m0s",
										"max-nodes-total":                       "1",
										"max-total-unready-percentage":          "46",
										"memory-total":                          "0:12800000",
										"min-replica-count":                     "1",
										"new-pod-scale-up-delay":                "10s",
										"node-autoprovisioning-enabled":         "true",
										"ok-total-unready-count":                "4",
										"scale-down-candidates-pool-min-count":  "51",
										"scale-down-candidates-pool-ratio":      "0.3",
										"scale-down-delay-after-add":            "20m0s",
										"scale-down-delay-after-delete":         "20s",
										"scale-down-delay-after-failure":        "4m0s",
										"scale-down-enabled":                    "false",
										"scale-down-non-empty-candidates-count": "50",
										"scale-down-unneeded-time":              "11m0s",
										"scale-down-unready-time":               "23m0s",
										"scale-down-utilization-threshold":      "0.8",
										"skip-nodes-with-local-storage":         "true",
										"skip-nodes-with-system-pods":           "false",
										"stderrthreshold":                       "7",
										"unremovable-node-recheck-timeout":      "9m0s",
										"v":                                     "6",
										"write-status-configmap":                "false",
										"balance-similar-node-groups":           "false",
									},
									Pools: []AddonNodePoolsConfig{
										{
											Name: "pool1",
											Config: map[string]string{
												"min-nodes": "1",
												"max-nodes": "100",
											},
										},
										{
											Name: "pool2",
											Config: map[string]string{
												"min-nodes": "3",
												"max-nodes": "10",
											},
										},
										{
											Name: "pool3",
											Config: map[string]string{
												"min-nodes": "1",
												"max-nodes": "6",
											},
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 5,
						},
						{
							Name:  "pool2",
							Count: 10,
						},
						{
							Name:  "pool3",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "30s",
						"expendable-pods-priority-cutoff":       "-20",
						"ignore-daemonsets-utilization":         "true",
						"ignore-mirror-pods-utilization":        "true",
						"max-autoprovisioned-node-group-count":  "25",
						"max-empty-bulk-delete":                 "12",
						"max-failing-time":                      "15m10s",
						"max-graceful-termination-sec":          "700",
						"max-inactivity":                        "11m0s",
						"max-node-provision-time":               "16m0s",
						"max-nodes-total":                       "1",
						"max-total-unready-percentage":          "46",
						"memory-total":                          "0:12800000",
						"min-replica-count":                     "1",
						"new-pod-scale-up-delay":                "10s",
						"node-autoprovisioning-enabled":         "true",
						"ok-total-unready-count":                "4",
						"scale-down-candidates-pool-min-count":  "51",
						"scale-down-candidates-pool-ratio":      "0.3",
						"scale-down-delay-after-add":            "20m0s",
						"scale-down-delay-after-delete":         "20s",
						"scale-down-delay-after-failure":        "4m0s",
						"scale-down-enabled":                    "false",
						"scale-down-non-empty-candidates-count": "50",
						"scale-down-unneeded-time":              "11m0s",
						"scale-down-unready-time":               "23m0s",
						"scale-down-utilization-threshold":      "0.8",
						"skip-nodes-with-local-storage":         "true",
						"skip-nodes-with-system-pods":           "false",
						"stderrthreshold":                       "7",
						"unremovable-node-recheck-timeout":      "9m0s",
						"v":                                     "6",
						"write-status-configmap":                "false",
						"balance-similar-node-groups":           "false",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "100",
							},
						},
						{
							Name: "pool2",
							Config: map[string]string{
								"min-nodes": "3",
								"max-nodes": "10",
							},
						},
						{
							Name: "pool3",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "6",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled with configuration - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"scan-interval":                         "30s",
										"expendable-pods-priority-cutoff":       "-20",
										"ignore-daemonsets-utilization":         "true",
										"ignore-mirror-pods-utilization":        "true",
										"max-autoprovisioned-node-group-count":  "25",
										"max-empty-bulk-delete":                 "12",
										"max-failing-time":                      "15m10s",
										"max-graceful-termination-sec":          "700",
										"max-inactivity":                        "11m0s",
										"max-node-provision-time":               "16m0s",
										"max-nodes-total":                       "1",
										"max-total-unready-percentage":          "46",
										"memory-total":                          "0:12800000",
										"min-replica-count":                     "1",
										"new-pod-scale-up-delay":                "10s",
										"node-autoprovisioning-enabled":         "true",
										"ok-total-unready-count":                "4",
										"scale-down-candidates-pool-min-count":  "51",
										"scale-down-candidates-pool-ratio":      "0.3",
										"scale-down-delay-after-add":            "20m0s",
										"scale-down-delay-after-delete":         "20s",
										"scale-down-delay-after-failure":        "4m0s",
										"scale-down-enabled":                    "false",
										"scale-down-non-empty-candidates-count": "50",
										"scale-down-unneeded-time":              "11m0s",
										"scale-down-unready-time":               "23m0s",
										"scale-down-utilization-threshold":      "0.8",
										"skip-nodes-with-local-storage":         "true",
										"skip-nodes-with-system-pods":           "false",
										"stderrthreshold":                       "7",
										"unremovable-node-recheck-timeout":      "9m0s",
										"v":                                     "6",
										"write-status-configmap":                "false",
										"balance-similar-node-groups":           "false",
									},
									Pools: []AddonNodePoolsConfig{
										{
											Name: "pool1",
											Config: map[string]string{
												"min-nodes": "1",
												"max-nodes": "100",
											},
										},
										{
											Name: "pool2",
											Config: map[string]string{
												"min-nodes": "3",
												"max-nodes": "10",
											},
										},
										{
											Name: "pool3",
											Config: map[string]string{
												"min-nodes": "1",
												"max-nodes": "6",
											},
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 5,
						},
						{
							Name:  "pool2",
							Count: 10,
						},
						{
							Name:  "pool3",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "30s",
						"expendable-pods-priority-cutoff":       "-20",
						"ignore-daemonsets-utilization":         "true",
						"ignore-mirror-pods-utilization":        "true",
						"max-autoprovisioned-node-group-count":  "25",
						"max-empty-bulk-delete":                 "12",
						"max-failing-time":                      "15m10s",
						"max-graceful-termination-sec":          "700",
						"max-inactivity":                        "11m0s",
						"max-node-provision-time":               "16m0s",
						"max-nodes-total":                       "1",
						"max-total-unready-percentage":          "46",
						"memory-total":                          "0:12800000",
						"min-replica-count":                     "1",
						"new-pod-scale-up-delay":                "10s",
						"node-autoprovisioning-enabled":         "true",
						"ok-total-unready-count":                "4",
						"scale-down-candidates-pool-min-count":  "51",
						"scale-down-candidates-pool-ratio":      "0.3",
						"scale-down-delay-after-add":            "20m0s",
						"scale-down-delay-after-delete":         "20s",
						"scale-down-delay-after-failure":        "4m0s",
						"scale-down-enabled":                    "false",
						"scale-down-non-empty-candidates-count": "50",
						"scale-down-unneeded-time":              "11m0s",
						"scale-down-unready-time":               "23m0s",
						"scale-down-utilization-threshold":      "0.8",
						"skip-nodes-with-local-storage":         "true",
						"skip-nodes-with-system-pods":           "false",
						"stderrthreshold":                       "7",
						"unremovable-node-recheck-timeout":      "9m0s",
						"v":                                     "6",
						"write-status-configmap":                "false",
						"balance-similar-node-groups":           "false",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "100",
							},
						},
						{
							Name: "pool2",
							Config: map[string]string{
								"min-nodes": "3",
								"max-nodes": "10",
							},
						},
						{
							Name: "pool3",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "6",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled with mixed configuration plus defaults",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"scan-interval":                        "30s",
										"expendable-pods-priority-cutoff":      "-20",
										"ignore-daemonsets-utilization":        "true",
										"ignore-mirror-pods-utilization":       "true",
										"max-autoprovisioned-node-group-count": "25",
										"max-empty-bulk-delete":                "12",
										"max-failing-time":                     "15m10s",
										"max-graceful-termination-sec":         "700",
										"max-inactivity":                       "11m0s",
										"max-node-provision-time":              "16m0s",
										"max-nodes-total":                      "1",
										"max-total-unready-percentage":         "46",
									},
									Pools: []AddonNodePoolsConfig{
										{
											Name: "pool1",
											Config: map[string]string{
												"min-nodes": "1",
												"max-nodes": "100",
											},
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 5,
						},
						{
							Name:  "pool2",
							Count: 10,
						},
						{
							Name:  "pool3",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "30s",
						"expendable-pods-priority-cutoff":       "-20",
						"ignore-daemonsets-utilization":         "true",
						"ignore-mirror-pods-utilization":        "true",
						"max-autoprovisioned-node-group-count":  "25",
						"max-empty-bulk-delete":                 "12",
						"max-failing-time":                      "15m10s",
						"max-graceful-termination-sec":          "700",
						"max-inactivity":                        "11m0s",
						"max-node-provision-time":               "16m0s",
						"max-nodes-total":                       "1",
						"max-total-unready-percentage":          "46",
						"memory-total":                          "0:6400000",
						"min-replica-count":                     "0",
						"new-pod-scale-up-delay":                "0s",
						"node-autoprovisioning-enabled":         "false",
						"ok-total-unready-count":                "3",
						"scale-down-candidates-pool-min-count":  "50",
						"scale-down-candidates-pool-ratio":      "0.1",
						"scale-down-delay-after-add":            "10m0s",
						"scale-down-delay-after-delete":         "1m",
						"scale-down-delay-after-failure":        "3m0s",
						"scale-down-enabled":                    "true",
						"scale-down-non-empty-candidates-count": "30",
						"scale-down-unneeded-time":              "10m0s",
						"scale-down-unready-time":               "20m0s",
						"scale-down-utilization-threshold":      "0.5",
						"skip-nodes-with-local-storage":         "false",
						"skip-nodes-with-system-pods":           "true",
						"stderrthreshold":                       "2",
						"unremovable-node-recheck-timeout":      "5m0s",
						"v":                                     "3",
						"write-status-configmap":                "true",
						"balance-similar-node-groups":           "true",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "100",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cluster-autoscaler addon enabled with mixed configuration plus defaults - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"scan-interval":                        "30s",
										"expendable-pods-priority-cutoff":      "-20",
										"ignore-daemonsets-utilization":        "true",
										"ignore-mirror-pods-utilization":       "true",
										"max-autoprovisioned-node-group-count": "25",
										"max-empty-bulk-delete":                "12",
										"max-failing-time":                     "15m10s",
										"max-graceful-termination-sec":         "700",
										"max-inactivity":                       "11m0s",
										"max-node-provision-time":              "16m0s",
										"max-nodes-total":                      "1",
										"max-total-unready-percentage":         "46",
									},
									Pools: []AddonNodePoolsConfig{
										{
											Name: "pool1",
											Config: map[string]string{
												"min-nodes": "1",
												"max-nodes": "100",
											},
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 5,
						},
						{
							Name:  "pool2",
							Count: 10,
						},
						{
							Name:  "pool3",
							Count: 1,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "30s",
						"expendable-pods-priority-cutoff":       "-20",
						"ignore-daemonsets-utilization":         "true",
						"ignore-mirror-pods-utilization":        "true",
						"max-autoprovisioned-node-group-count":  "25",
						"max-empty-bulk-delete":                 "12",
						"max-failing-time":                      "15m10s",
						"max-graceful-termination-sec":          "700",
						"max-inactivity":                        "11m0s",
						"max-node-provision-time":               "16m0s",
						"max-nodes-total":                       "1",
						"max-total-unready-percentage":          "46",
						"memory-total":                          "0:6400000",
						"min-replica-count":                     "0",
						"new-pod-scale-up-delay":                "0s",
						"node-autoprovisioning-enabled":         "false",
						"ok-total-unready-count":                "3",
						"scale-down-candidates-pool-min-count":  "50",
						"scale-down-candidates-pool-ratio":      "0.1",
						"scale-down-delay-after-add":            "10m0s",
						"scale-down-delay-after-delete":         "1m",
						"scale-down-delay-after-failure":        "3m0s",
						"scale-down-enabled":                    "true",
						"scale-down-non-empty-candidates-count": "30",
						"scale-down-unneeded-time":              "10m0s",
						"scale-down-unready-time":               "20m0s",
						"scale-down-utilization-threshold":      "0.5",
						"skip-nodes-with-local-storage":         "false",
						"skip-nodes-with-system-pods":           "true",
						"stderrthreshold":                       "2",
						"unremovable-node-recheck-timeout":      "5m0s",
						"v":                                     "3",
						"write-status-configmap":                "true",
						"balance-similar-node-groups":           "true",
					},
					Pools: []AddonNodePoolsConfig{
						{
							Name: "pool1",
							Config: map[string]string{
								"min-nodes": "1",
								"max-nodes": "100",
							},
						},
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "smb-flexvolume addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.SMBFlexVolumeAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.SMBFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          k8sComponentsByVersionMap["1.15.4"][common.SMBFlexVolumeAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "rescheduler addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ReschedulerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ReschedulerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.ReschedulerAddonName,
							CPURequests:    "10m",
							MemoryRequests: "100Mi",
							CPULimits:      "10m",
							MemoryLimits:   "100Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ReschedulerAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "nvidia addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							VMSize: "Standard_NC6",
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.NVIDIADevicePluginAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.NVIDIADevicePluginAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          specConfig.NVIDIAImageBase + k8sComponentsByVersionMap["1.15.4"][common.NVIDIADevicePluginAddonName],
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "container-monitoring addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ContainerMonitoringAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.ContainerMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"omsAgentVersion":       "1.10.0.1",
						"dockerProviderVersion": "8.0.0-3",
						"schema-versions":       "v1",
						"clusterName":           "aks-engine-cluster",
						"workspaceDomain":       "b3BpbnNpZ2h0cy5henVyZS5jb20=",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           "omsagent",
							CPURequests:    "150m",
							MemoryRequests: "250Mi",
							CPULimits:      "1",
							MemoryLimits:   "750Mi",
							Image:          "mcr.microsoft.com/azuremonitor/containerinsights/ciprod:ciprod03022020",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "Azure Network Policy addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyAzure,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.AzureNetworkPolicyAddonName,
							Image:          k8sComponentsByVersionMap["1.15.4"][common.AzureNetworkPolicyAddonName],
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "100m",
							MemoryLimits:   "200Mi",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "Azure Network Policy addon enabled - 1.16 upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.AzureNetworkPolicyAddonName,
									Enabled: to.BoolPtr(true),
									Containers: []KubernetesContainerSpec{
										{
											Name:  common.AzureNetworkPolicyAddonName,
											Image: k8sComponentsByVersionMap["1.16.0"][common.AzureNetworkPolicyAddonName],
										},
										{
											Name:  common.AzureVnetTelemetryContainerName,
											Image: k8sComponentsByVersionMap["1.16.0"][common.AzureVnetTelemetryContainerName],
										},
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.AzureNetworkPolicyAddonName,
							Image:          k8sComponentsByVersionMap["1.16.0"][common.AzureNetworkPolicyAddonName],
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "100m",
							MemoryLimits:   "200Mi",
						},
					},
				},
			}, "1.16.0"),
		},
		{
			name: "calico addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyCalico,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.AzureCNINetworkMonitorAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.CalicoAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CalicoTyphaComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoTyphaComponentName],
						},
						{
							Name:  common.CalicoCNIComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoCNIComponentName],
						},
						{
							Name:  common.CalicoNodeComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoNodeComponentName],
						},
						{
							Name:  common.CalicoPod2DaemonComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoPod2DaemonComponentName],
						},
						{
							Name:  common.CalicoClusterAutoscalerComponentName,
							Image: k8sComponentsByVersionMap["1.15.4"][common.CalicoClusterAutoscalerComponentName],
						},
					},
				},
			}, "1.15.4")),
		},
		{
			name: "calico addon back-compat",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyCalico,
							Addons: []KubernetesAddon{
								{
									Name:    common.CalicoAddonName,
									Enabled: to.BoolPtr(false),
									Containers: []KubernetesContainerSpec{
										{
											Name:  common.CalicoTyphaComponentName,
											Image: specConfig.CalicoImageBase + "typha:old", // confirm that upgrade will change this to default image
										},
										{
											Name:  common.CalicoCNIComponentName,
											Image: specConfig.CalicoImageBase + "cni:v3.8.0",
										},
										{
											Name:  common.CalicoNodeComponentName,
											Image: specConfig.CalicoImageBase + "node:v3.8.0",
										},
										{
											Name:  common.CalicoPod2DaemonComponentName,
											Image: specConfig.CalicoImageBase + "pod2daemon-flexvol:v3.8.0",
										},
										{
											Name:  common.CalicoClusterAutoscalerComponentName,
											Image: specConfig.MCRKubernetesImageBase + "cluster-proportional-autoscaler-amd64:1.1.2-r2",
										},
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: omitFromAddons([]string{common.AzureCNINetworkMonitorAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.CalicoAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CalicoTyphaComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoTyphaComponentName],
						},
						{
							Name:  common.CalicoCNIComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoCNIComponentName],
						},
						{
							Name:  common.CalicoNodeComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoNodeComponentName],
						},
						{
							Name:  common.CalicoPod2DaemonComponentName,
							Image: specConfig.CalicoImageBase + k8sComponentsByVersionMap["1.15.4"][common.CalicoPod2DaemonComponentName],
						},
						{
							Name:  common.CalicoClusterAutoscalerComponentName,
							Image: k8sComponentsByVersionMap["1.15.4"][common.CalicoClusterAutoscalerComponentName],
						},
					},
				},
			}, "1.15.4")),
		},
		{
			name: "aad-pod-identity enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.AADPodIdentityAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AADPodIdentityAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.NMIContainerName,
							Image:          k8sComponentsByVersionMap["1.15.4"][common.NMIContainerName],
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
						},
						{
							Name:           common.MICContainerName,
							Image:          k8sComponentsByVersionMap["1.15.4"][common.MICContainerName],
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "azure-policy addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.AzurePolicyAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzurePolicyAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"auditInterval":             "60",
						"constraintViolationsLimit": "100",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.AzurePolicyAddonName,
							Image:          k8sComponentsByVersionMap["1.15.4"][common.AzurePolicyAddonName],
							CPURequests:    "30m",
							MemoryRequests: "50Mi",
							CPULimits:      "100m",
							MemoryLimits:   "200Mi",
						},
						{
							Name:           common.GatekeeperContainerName,
							Image:          k8sComponentsByVersionMap["1.15.4"][common.GatekeeperContainerName],
							CPURequests:    "100m",
							MemoryRequests: "256Mi",
							CPULimits:      "1000m",
							MemoryLimits:   "512Mi",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "cilium networkPolicy",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPolicy: NetworkPolicyCilium,
							NetworkPlugin: NetworkPluginCilium,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.IPMASQAgentAddonName, common.AzureCNINetworkMonitorAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.CiliumAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CiliumAgentContainerName,
							Image: k8sComponentsByVersionMap["1.15.4"][common.CiliumAgentContainerName],
						},
						{
							Name:  common.CiliumCleanStateContainerName,
							Image: k8sComponentsByVersionMap["1.15.4"][common.CiliumCleanStateContainerName],
						},
						{
							Name:  common.CiliumOperatorContainerName,
							Image: k8sComponentsByVersionMap["1.15.4"][common.CiliumOperatorContainerName],
						},
						{
							Name:  common.CiliumEtcdOperatorContainerName,
							Image: k8sComponentsByVersionMap["1.15.4"][common.CiliumEtcdOperatorContainerName],
						},
					},
				},
			}, "1.15.4")),
		},
		{
			name: "Azure Stack addons",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.14.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							VMSize: "Standard_NC6", // to validate that Azure Stack cluster config does not get nvidia addon
						},
					},
					CustomCloudProfile: &CustomCloudProfile{
						Environment: &azure.Environment{
							Name: "AzureStackCloud",
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    common.DashboardAddonName,
					Enabled: to.BoolPtr(false),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          "MCRKubernetesImageBase" + k8sComponentsByVersionMap["1.14.0"][common.DashboardAddonName],
						},
					},
				},
				{
					Name:    common.MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.MetricsServerAddonName,
							Image: "MCRKubernetesImageBase" + k8sComponentsByVersionMap["1.14.0"][common.MetricsServerAddonName],
						},
					},
				},
				{
					Name:    common.IPMASQAgentAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          "MCRKubernetesImageBase" + k8sComponentsByVersionMap["1.14.0"][common.IPMASQAgentAddonName],
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    common.AzureCNINetworkMonitorAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.AzureCNINetworkMonitorAddonName,
							Image:          "AzureCNIImageBase" + k8sComponentsByVersionMap["1.14.0"][common.AzureCNINetworkMonitorAddonName],
							CPURequests:    "30m",
							MemoryRequests: "25Mi",
							CPULimits:      "200m",
							MemoryLimits:   "256Mi",
						},
					},
				},
				{
					Name:    common.CoreDNSAddonName,
					Enabled: to.BoolPtr(DefaultCoreDNSAddonEnabled),
					Config: map[string]string{
						"domain":            "cluster.local",
						"clusterIP":         DefaultKubernetesDNSServiceIP,
						"cores-per-replica": "512",
						"nodes-per-replica": "32",
						"min-replicas":      "1",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CoreDNSAddonName,
							Image: "MCRKubernetesImageBase" + k8sComponentsByVersionMap["1.14.0"][common.CoreDNSAddonName],
						},
						{
							Name:  common.CoreDNSAutoscalerName,
							Image: k8sComponentsByVersionMap["1.14.0"][common.CoreDNSAutoscalerName],
						},
					},
				},
				{
					Name:    common.KubeProxyAddonName,
					Enabled: to.BoolPtr(DefaultKubeProxyAddonEnabled),
					Config: map[string]string{
						"cluster-cidr": DefaultKubernetesSubnet,
						"proxy-mode":   string(KubeProxyModeIPTables),
						"featureGates": "{}",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.KubeProxyAddonName,
							Image: "MCRKubernetesImageBase" + k8sComponentsByVersionMap["1.14.0"][common.KubeProxyAddonName] + common.AzureStackSuffix,
						},
					},
				},
			},
		},
		{
			name: "azure disk and azure file csi driver enabled for k8s >= 1.13.0 and UseCloudControllerManager is true",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureFileContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAzureFileContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotControllerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotControllerContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureDiskContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAzureDiskContainerName, k8sComponentsByVersionMap["1.15.4"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "azure cloud-node-manager enabled for k8s == 1.16 and useCloudControllerManager is true",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.1",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureFileContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAzureFileContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotControllerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotControllerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureDiskContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAzureDiskContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
			}, "1.16.1"),
		},
		{
			name: "azure cloud-node-manager disabled for Windows cluster with k8s < 1.18 and useCloudControllerManager is true",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.17.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							OSType: Windows,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    common.CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			}, "1.17.0"),
		},
		{
			name: "azure cloud-node-manager enabled for Windows cluster with k8s == 1.18 and useCloudControllerManager is true",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							OSType: Windows,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureFileContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAzureFileContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotControllerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotControllerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureDiskContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAzureDiskContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
			}, "1.18.0"),
		},
		{
			name: "azure cloud-node-manager enabled for k8s == 1.17.0 and useCloudControllerManager is true",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.17.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureFileContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAzureFileContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotControllerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotControllerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureDiskContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAzureDiskContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
			}, "1.17.0"),
		},
		{
			name: "azure cloud-node-manager enabled for k8s >= 1.17.0 - upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.17.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
							Addons: []KubernetesAddon{
								{
									Name:    common.AzureDiskCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    common.AzureFileCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    common.CloudNodeManagerAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureFileContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAzureFileContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotControllerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotControllerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureDiskContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAzureDiskContainerName, k8sComponentsByVersionMap["1.17.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
			}, "1.17.0"),
		},
		{
			name: "azure cloud-node-manager enabled for Windows cluster with k8s >= 1.18.0 - upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
							Addons: []KubernetesAddon{
								{
									Name:    common.AzureDiskCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    common.AzureFileCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    common.CloudNodeManagerAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							OSType: Windows,
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureFileContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAzureFileContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotControllerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotControllerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureDiskContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAzureDiskContainerName, k8sComponentsByVersionMap["1.18.0"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
			}, "1.18.0"),
		},
		{
			name: "azure cloud-node-manager enabled for k8s == 1.16.1 upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.1",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
							Addons: []KubernetesAddon{
								{
									Name:    common.AzureDiskCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    common.AzureFileCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    common.CloudNodeManagerAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureFileContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureFileCSIDriverAddonName, common.CSIAzureFileContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.CSIProvisionerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIProvisionerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAttacherContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAttacherContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSILivenessProbeWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSILivenessProbeWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSINodeDriverRegistrarWindowsContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSINodeDriverRegistrarWindowsContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotterContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotterContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSISnapshotControllerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSISnapshotControllerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIResizerContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIResizerContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
						{
							Name:           common.CSIAzureDiskContainerName,
							Image:          specConfig.MCRKubernetesImageBase + getCSISidecarComponent(common.AzureDiskCSIDriverAddonName, common.CSIAzureDiskContainerName, k8sComponentsByVersionMap["1.16.1"]),
							CPURequests:    "10m",
							MemoryRequests: "20Mi",
							CPULimits:      "2",
							MemoryLimits:   "2Gi",
						},
					},
				},
				{
					Name:    common.CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
			}, "1.16.1"),
		},
		{
			name: "upgrade w/ no kube-dns or coredns specified", // back-compat support for clusters configured prior to user-configurable coredns and kube-dns addons
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.1",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons:        []KubernetesAddon{},
						},
					},
				},
			},
			isUpgrade:      true,
			expectedAddons: getDefaultAddons("1.18.1", "", common.KubernetesImageBaseTypeMCR),
		},
		{
			name: "upgrade w/ manual kube-dns enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.1",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.KubeDNSAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: omitFromAddons([]string{common.CoreDNSAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.KubeDNSAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"domain":    "cluster.local",
						"clusterIP": DefaultKubernetesDNSServiceIP,
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  "kubedns",
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.1"][common.KubeDNSAddonName],
						},
						{
							Name:  common.DNSMasqComponentName,
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.1"][common.DNSMasqComponentName],
						},
						{
							Name:  "sidecar",
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.1"][common.DNSSidecarComponentName],
						},
					},
				},
			}, "1.18.1")),
		},
		{
			name: "upgrade w/ manual coredns enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.1",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade:      true,
			expectedAddons: getDefaultAddons("1.18.1", "", common.KubernetesImageBaseTypeMCR),
		},
		{
			name: "kube-dns enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.KubeDNSAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.CoreDNSAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.KubeDNSAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"domain":    "cluster.local",
						"clusterIP": DefaultKubernetesDNSServiceIP,
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  "kubedns",
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.KubeDNSAddonName],
						},
						{
							Name:  common.DNSMasqComponentName,
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.DNSMasqComponentName],
						},
						{
							Name:  "sidecar",
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.DNSSidecarComponentName],
						},
					},
				},
			}, "1.15.4")),
		},
		{
			name: "coredns enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade:      false,
			expectedAddons: getDefaultAddons("1.15.4", "", common.KubernetesImageBaseTypeMCR),
		},
		{
			name: "coredns w/ user configuration",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.1",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"min-replicas": "3",
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: overwriteDefaultAddons([]KubernetesAddon{
				{
					Name:    common.CoreDNSAddonName,
					Enabled: to.BoolPtr(DefaultCoreDNSAddonEnabled),
					Config: map[string]string{
						"domain":            "cluster.local",
						"clusterIP":         DefaultKubernetesDNSServiceIP,
						"cores-per-replica": "512",
						"nodes-per-replica": "32",
						"min-replicas":      "3",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CoreDNSAddonName,
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.1"][common.CoreDNSAddonName],
						},
						{
							Name:  common.CoreDNSAutoscalerName,
							Image: k8sComponentsByVersionMap["1.18.1"][common.CoreDNSAutoscalerName],
						},
					},
				},
			}, "1.18.1"),
		},
		{
			name: "kube-proxy w/ user configuration",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.KubeProxyAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"cluster-cidr": "foo",
										"proxy-mode":   "bar",
										"featureGates": "baz",
									},
									Containers: []KubernetesContainerSpec{
										{
											Name:  common.KubeProxyAddonName,
											Image: "bam",
										},
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: overwriteDefaultAddons([]KubernetesAddon{
				{
					Name:    common.KubeProxyAddonName,
					Enabled: to.BoolPtr(DefaultKubeProxyAddonEnabled),
					Config: map[string]string{
						"cluster-cidr": "foo",
						"proxy-mode":   "bar",
						"featureGates": "baz",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.KubeProxyAddonName,
							Image: "bam",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "kube-proxy disabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.KubeProxyAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade:      false,
			expectedAddons: omitFromAddons([]string{common.KubeProxyAddonName}, getDefaultAddons("1.15.4", "", common.KubernetesImageBaseTypeMCR)),
		},
		{
			name: "node-problem-detector addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.NodeProblemDetectorAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.NodeProblemDetectorAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"customPluginMonitor": "/config/kernel-monitor-counter.json,/config/systemd-monitor-counter.json",
						"systemLogMonitor":    "/config/kernel-monitor.json,/config/docker-monitor.json,/config/systemd-monitor.json",
						"systemStatsMonitor":  "/config/system-stats-monitor.json",
						"versionLabel":        "v0.8.1",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.NodeProblemDetectorAddonName,
							Image:          k8sComponentsByVersionMap["1.15.4"][common.NodeProblemDetectorAddonName],
							CPURequests:    "20m",
							MemoryRequests: "20Mi",
							CPULimits:      "200m",
							MemoryLimits:   "100Mi",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "pod-security-policy upgrade to 1.15",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
						},
					},
				},
			},
			isUpgrade:      true,
			expectedAddons: getDefaultAddons("1.15.4", "", common.KubernetesImageBaseTypeMCR),
		},
		{
			name: "pod-security-policy disabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.PodSecurityPolicyAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade:      false,
			expectedAddons: omitFromAddons([]string{common.PodSecurityPolicyAddonName}, getDefaultAddons("1.15.4", "", common.KubernetesImageBaseTypeMCR)),
		},
		{
			name: "pod-security-policy disabled during upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.PodSecurityPolicyAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade:      true,
			expectedAddons: omitFromAddons([]string{common.PodSecurityPolicyAddonName}, getDefaultAddons("1.15.4", "", common.KubernetesImageBaseTypeMCR)),
		},
		{
			name: "audit-policy disabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.AuditPolicyAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade:      false,
			expectedAddons: omitFromAddons([]string{common.AuditPolicyAddonName}, getDefaultAddons("1.15.4", "", common.KubernetesImageBaseTypeMCR)),
		},
		{
			name: "aad-default-aad-admin-group addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
						},
					},
					AADProfile: &AADProfile{
						AdminGroupID: "7d04bcd3-3c48-49ab-a064-c0b7d69896da",
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AADAdminGroupAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"adminGroupID": "7d04bcd3-3c48-49ab-a064-c0b7d69896da",
					},
				},
			}, "1.15.4"),
		},
		{
			name: "antrea addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAntrea,
							Addons: []KubernetesAddon{
								{
									Name:    common.AntreaAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"serviceCidr":      DefaultKubernetesServiceCIDR,
										"trafficEncapMode": common.AntreaDefaultTrafficEncapMode,
										"installCniCmd":    common.AntreaDefaultInstallCniCmd,
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.IPMASQAgentAddonName, common.AzureCNINetworkMonitorAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AntreaAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"serviceCidr":      DefaultKubernetesServiceCIDR,
						"trafficEncapMode": common.AntreaDefaultTrafficEncapMode,
						"installCniCmd":    common.AntreaDefaultInstallCniCmd,
					},
				},
			}, "1.15.4")),
		},
		{
			name: "antrea addon enabled with azure",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPolicy: NetworkPolicyAntrea,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.AntreaAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"serviceCidr": DefaultKubernetesServiceCIDR,
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.AzureCNINetworkMonitorAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.AntreaAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"trafficEncapMode": common.AntreaNetworkPolicyOnlyMode,
						"installCniCmd":    common.AntreaInstallCniChainCmd,
						"serviceCidr":      DefaultKubernetesServiceCIDR,
					},
				},
			}, "1.15.4")),
		},
		{
			name: "addons with IPv6 single stack",
			cs: &ContainerService{
				Properties: &Properties{
					FeatureFlags: &FeatureFlags{
						EnableIPv6Only: true,
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIPv6,
							NetworkPlugin:           NetworkPluginKubenet,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
								"--node-ip":        "::",
							},
							ClusterSubnet: DefaultKubernetesClusterSubnetIPv6,
							ProxyMode:     KubeProxyModeIPTables,
							APIServerConfig: map[string]string{
								"--bind-address": "::",
							},
							ControllerManagerConfig: map[string]string{
								"--bind-address": "::",
							},
							SchedulerConfig: map[string]string{
								"--bind-address": "::",
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.AzureCNINetworkMonitorAddonName}, overwriteDefaultAddons([]KubernetesAddon{
				{
					Name:    common.CoreDNSAddonName,
					Enabled: to.BoolPtr(DefaultCoreDNSAddonEnabled),
					Config: map[string]string{
						"domain":            "cluster.local",
						"clusterIP":         DefaultKubernetesDNSServiceIPv6,
						"use-host-network":  "true",
						"cores-per-replica": "512",
						"nodes-per-replica": "32",
						"min-replicas":      "1",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CoreDNSAddonName,
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.0"][common.CoreDNSAddonName],
						},
						{
							Name:  common.CoreDNSAutoscalerName,
							Image: k8sComponentsByVersionMap["1.18.0"][common.CoreDNSAutoscalerName],
						},
					},
				},
				{
					Name:    common.IPMASQAgentAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.0"][common.IPMASQAgentAddonName],
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr":           DefaultKubernetesClusterSubnetIPv6,
						"enable-ipv6":                   "true",
						"non-masq-cni-cidr":             "",
						"secondary-non-masquerade-cidr": "",
					},
				},
				{
					Name:    common.KubeProxyAddonName,
					Enabled: to.BoolPtr(DefaultKubeProxyAddonEnabled),
					Config: map[string]string{
						"cluster-cidr":         DefaultKubernetesClusterSubnetIPv6,
						"proxy-mode":           string(KubeProxyModeIPTables),
						"featureGates":         "{}",
						"bind-address":         "::",
						"healthz-bind-address": "::",
						"metrics-bind-address": "::1",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.KubeProxyAddonName,
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.0"][common.KubeProxyAddonName],
						},
					},
				},
			}, "1.18.0")),
		},
		{
			name: "addons with dual stack",
			cs: &ContainerService{
				Properties: &Properties{
					FeatureFlags: &FeatureFlags{
						EnableIPv6DualStack: true,
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.18.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							NetworkPlugin:           NetworkPluginKubenet,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
								"--feature-gates":  "IPv6DualStack=true",
							},
							ClusterSubnet: DefaultKubernetesClusterSubnet + "," + DefaultKubernetesClusterSubnetIPv6,
							ServiceCIDR:   DefaultKubernetesServiceCIDR + "," + DefaultKubernetesServiceCIDRIPv6,
							ProxyMode:     KubeProxyModeIPVS,
							APIServerConfig: map[string]string{
								"--feature-gates": "IPv6DualStack=true",
							},
							ControllerManagerConfig: map[string]string{
								"--feature-gates": "IPv6DualStack=true",
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.AzureCNINetworkMonitorAddonName}, overwriteDefaultAddons([]KubernetesAddon{
				{
					Name:    common.IPMASQAgentAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.0"][common.IPMASQAgentAddonName],
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr":           DefaultKubernetesClusterSubnet,
						"enable-ipv6":                   "true",
						"non-masq-cni-cidr":             "",
						"secondary-non-masquerade-cidr": DefaultKubernetesClusterSubnetIPv6,
					},
				},
				{
					Name:    common.KubeProxyAddonName,
					Enabled: to.BoolPtr(DefaultKubeProxyAddonEnabled),
					Config: map[string]string{
						"cluster-cidr": DefaultKubernetesClusterSubnet + "," + DefaultKubernetesClusterSubnetIPv6,
						"proxy-mode":   string(KubeProxyModeIPVS),
						"featureGates": "IPv6DualStack: true",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.KubeProxyAddonName,
							Image: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.18.0"][common.KubeProxyAddonName],
						},
					},
				},
			}, "1.18.0")),
		},
		{
			name: "kube proxy w/ customKubeProxyImage",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:        DefaultKubernetesSubnet,
							ProxyMode:            KubeProxyModeIPTables,
							NetworkPlugin:        NetworkPluginAzure,
							CustomKubeProxyImage: "my-custom-kube-proxy-image",
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: overwriteDefaultAddons([]KubernetesAddon{
				{
					Name:    common.KubeProxyAddonName,
					Enabled: to.BoolPtr(DefaultKubeProxyAddonEnabled),
					Config: map[string]string{
						"cluster-cidr": DefaultKubernetesSubnet,
						"proxy-mode":   string(KubeProxyModeIPTables),
						"featureGates": "{}",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.KubeProxyAddonName,
							Image: "my-custom-kube-proxy-image",
						},
					},
				},
			}, "1.15.4"),
		},
		{
			name: "csi-secrets-store addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.SecretsStoreCSIDriverAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.SecretsStoreCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
				},
			}, "1.15.4"),
		},
		{
			name: "keyvault-flexvolume enabled for k8s >= 1.16.0 - upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet:             DefaultKubernetesSubnet,
							ProxyMode:                 KubeProxyModeIPTables,
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(false),
							Addons: []KubernetesAddon{
								{
									Name:    common.KeyVaultFlexVolumeAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: omitFromAddons([]string{common.SecretsStoreCSIDriverAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          k8sComponentsByVersionMap["1.16.0"][common.KeyVaultFlexVolumeAddonName],
						},
					},
				},
			}, "1.16.0")),
		},
		{
			name: "keyvault-flexvolume enabled for 1.16+",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.0",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.KeyVaultFlexVolumeAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: omitFromAddons([]string{common.SecretsStoreCSIDriverAddonName}, concatenateDefaultAddons([]KubernetesAddon{
				{
					Name:    common.KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           common.KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          k8sComponentsByVersionMap["1.16.0"][common.KeyVaultFlexVolumeAddonName],
						},
					},
				},
			}, "1.16.0")),
		},
		{
			name: "ip-masq-agent upgrade to 1.1.16 overrides custom image",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.8",
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
							DNSServiceIP:            DefaultKubernetesDNSServiceIP,
							KubeletConfig: map[string]string{
								"--cluster-domain": "cluster.local",
							},
							ClusterSubnet: DefaultKubernetesSubnet,
							ProxyMode:     KubeProxyModeIPTables,
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.IPMASQAgentAddonName,
									Enabled: to.BoolPtr(true),
									Containers: []KubernetesContainerSpec{
										{
											Name:  common.IPMASQAgentAddonName,
											Image: "my-custom-image",
										},
									},
								},
							},
						},
					},
				},
			},
			isUpgrade:      true,
			expectedAddons: getDefaultAddons("1.16.8", "", common.KubernetesImageBaseTypeMCR), // the default addons will include the default image
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.cs.setAddonsConfig(test.isUpgrade)
			for _, addonName := range []string{
				common.HeapsterAddonName,
				common.TillerAddonName,
				common.ACIConnectorAddonName,
				common.ClusterAutoscalerAddonName,
				common.BlobfuseFlexVolumeAddonName,
				common.SMBFlexVolumeAddonName,
				common.KeyVaultFlexVolumeAddonName,
				common.DashboardAddonName,
				common.ReschedulerAddonName,
				common.MetricsServerAddonName,
				common.NVIDIADevicePluginAddonName,
				common.ContainerMonitoringAddonName,
				common.IPMASQAgentAddonName,
				common.AzureCNINetworkMonitorAddonName,
				common.AzureNetworkPolicyAddonName,
				common.CalicoAddonName,
				common.AADPodIdentityAddonName,
				common.AzurePolicyAddonName,
				common.AzureFileCSIDriverAddonName,
				common.AzureDiskCSIDriverAddonName,
				common.CloudNodeManagerAddonName,
				common.CoreDNSAddonName,
				common.KubeDNSAddonName,
				common.KubeProxyAddonName,
				common.NodeProblemDetectorAddonName,
				common.PodSecurityPolicyAddonName,
				common.AADAdminGroupAddonName,
				common.SecretsStoreCSIDriverAddonName,
				common.AntreaAddonName,
			} {
				addon := test.cs.Properties.OrchestratorProfile.KubernetesConfig.Addons[getAddonsIndexByName(test.cs.Properties.OrchestratorProfile.KubernetesConfig.Addons, addonName)]
				if addon.IsEnabled() {
					if i := getAddonsIndexByName(test.expectedAddons, addonName); i == -1 {
						t.Fatalf("got addon %s that we weren't expecting", addon.Name)
					}
					expectedAddon := test.expectedAddons[getAddonsIndexByName(test.expectedAddons, addonName)]
					if to.Bool(addon.Enabled) != to.Bool(expectedAddon.Enabled) {
						t.Fatalf("expected addon %s to have Enabled value %t, instead got %t", expectedAddon.Name, to.Bool(expectedAddon.Enabled), to.Bool(addon.Enabled))
					}
					if expectedAddon.Containers != nil {
						if len(expectedAddon.Containers) != len(addon.Containers) {
							t.Fatalf("expected addon %s to have %d containers , got %d", expectedAddon.Name, len(expectedAddon.Containers), len(addon.Containers))
						}
						for i, container := range expectedAddon.Containers {
							if container.Name != addon.Containers[i].Name {
								t.Fatalf("expected addon %s to have container Name %s at at Containers index %d, got %s", expectedAddon.Name, container.Name, i, addon.Containers[i].Name)
							}
							if container.Image != addon.Containers[i].Image {
								t.Fatalf("expected addon %s to have container Image %s at at Containers index %d, got %s", expectedAddon.Name, container.Image, i, addon.Containers[i].Image)
							}
							if container.CPURequests != addon.Containers[i].CPURequests {
								t.Fatalf("expected addon %s to have container CPURequests %s at at Containers index %d, got %s", expectedAddon.Name, container.CPURequests, i, addon.Containers[i].CPURequests)
							}
							if container.MemoryRequests != addon.Containers[i].MemoryRequests {
								t.Fatalf("expected addon %s to have container MemoryRequests %s at at Containers index %d, got %s", expectedAddon.Name, container.MemoryRequests, i, addon.Containers[i].MemoryRequests)
							}
							if container.CPULimits != addon.Containers[i].CPULimits {
								t.Fatalf("expected addon %s to have container CPULimits %s at at Containers index %d, got %s", expectedAddon.Name, container.CPULimits, i, addon.Containers[i].CPULimits)
							}
							if container.MemoryLimits != addon.Containers[i].MemoryLimits {
								t.Fatalf("expected addon %s to have container MemoryLimits %s at at Containers index %d, got %s", expectedAddon.Name, container.MemoryLimits, i, addon.Containers[i].MemoryLimits)
							}
						}
					}
					if expectedAddon.Config != nil {
						for key, val := range expectedAddon.Config {
							if val != addon.Config[key] {
								t.Fatalf("expected addon %s to have config %s=%s, got %s=%s", expectedAddon.Name, key, val, key, addon.Config[key])
							}
						}
					}
					if addon.Config != nil {
						for key, val := range addon.Config {
							if val != expectedAddon.Config[key] {
								t.Fatalf("expected addon %s to have config %s=%s, got %s=%s", addon.Name, key, val, key, expectedAddon.Config[key])
							}
						}
					}
					if expectedAddon.Pools != nil {
						if len(expectedAddon.Pools) != len(addon.Pools) {
							t.Fatalf("expected addon %s to have %d pools , got %d", expectedAddon.Name, len(expectedAddon.Pools), len(addon.Pools))
						}
						for i, expectedPool := range expectedAddon.Pools {
							if expectedPool.Name != addon.Pools[i].Name {
								t.Fatalf("expected addon %s to have pool Name %s at Pools index %d, got %s", expectedAddon.Name, expectedPool.Name, i, addon.Pools[i].Name)
							}
							if expectedPool.Config != nil {
								for key, val := range expectedPool.Config {
									if val != addon.Pools[i].Config[key] {
										t.Fatalf("expected addon %s to have pool config %s=%s for pool name %s, got %s=%s", expectedAddon.Name, key, val, expectedPool.Name, key, addon.Pools[i].Config[key])
									}
								}
							}
						}
					}
				} else {
					if i := getAddonsIndexByName(test.expectedAddons, addonName); i > -1 {
						if to.Bool(test.expectedAddons[i].Enabled) {
							t.Fatalf("expected addon %s to be enabled, instead it was disabled", addonName)
						}
					}
				}
			}
		})
	}
}
func TestMakeDefaultClusterAutoscalerAddonPoolsConfig(t *testing.T) {
	cases := []struct {
		name                     string
		cs                       *ContainerService
		expectedAddonPoolsConfig []AddonNodePoolsConfig
	}{
		{
			name: "1 pool",
			cs: &ContainerService{
				Properties: &Properties{
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 3,
						},
					},
				},
			},
			expectedAddonPoolsConfig: []AddonNodePoolsConfig{
				{
					Name: "pool1",
					Config: map[string]string{
						"min-nodes": "3",
						"max-nodes": "3",
					},
				},
			},
		},
		{
			name: "5 pools",
			cs: &ContainerService{
				Properties: &Properties{
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:  "pool1",
							Count: 3,
						},
						{
							Name:  "pool2",
							Count: 10,
						},
						{
							Name:  "pool3",
							Count: 1,
						},
						{
							Name:  "pool4",
							Count: 33,
						},
						{
							Name:  "pool5",
							Count: 5,
						},
					},
				},
			},
			expectedAddonPoolsConfig: []AddonNodePoolsConfig{
				{
					Name: "pool1",
					Config: map[string]string{
						"min-nodes": "3",
						"max-nodes": "3",
					},
				},
				{
					Name: "pool2",
					Config: map[string]string{
						"min-nodes": "10",
						"max-nodes": "10",
					},
				},
				{
					Name: "pool3",
					Config: map[string]string{
						"min-nodes": "1",
						"max-nodes": "1",
					},
				},
				{
					Name: "pool4",
					Config: map[string]string{
						"min-nodes": "33",
						"max-nodes": "33",
					},
				},
				{
					Name: "pool5",
					Config: map[string]string{
						"min-nodes": "5",
						"max-nodes": "5",
					},
				},
			},
		},
		{
			name: "0 pools",
			cs: &ContainerService{
				Properties: &Properties{},
			},
			expectedAddonPoolsConfig: []AddonNodePoolsConfig{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := makeDefaultClusterAutoscalerAddonPoolsConfig(c.cs)
			if c.expectedAddonPoolsConfig != nil {
				if len(c.expectedAddonPoolsConfig) != len(result) {
					t.Fatalf("expected to have %d pools, got %d", len(result), len(c.expectedAddonPoolsConfig))
				}
				for i, pool := range c.expectedAddonPoolsConfig {
					if pool.Name != result[i].Name {
						t.Fatalf("expected to have pool Name %s at Pools index %d, got %s", pool.Name, i, result[i].Name)
					}
					if pool.Config != nil {
						for key, val := range pool.Config {
							if val != result[i].Config[key] {
								t.Fatalf("expected to have pool config %s=%s for pool name %s, got %s=%s", key, val, pool.Name, key, result[i].Config[key])
							}
						}
					}
				}
			} else {
				if result != nil {
					t.Fatalf("expected a nil slice, got %#v", result)
				}
			}
		})
	}
}

func TestGetClusterAutoscalerNodesConfig(t *testing.T) {
	specConfig := AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	k8sComponentsByVersionMap := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	cases := []struct {
		name                string
		addon               KubernetesAddon
		cs                  *ContainerService
		expectedNodesConfig string
	}{
		{
			name: "1 pool",
			addon: KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
				},
			},
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
							UseManagedIdentity: true,
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedNodesConfig: "        - --nodes=1:10:k8s-pool1-49584119-vmss",
		},
		{
			name: "multiple pools",
			addon: KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
					{
						Name: "pool2",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
				},
			},
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
							UseManagedIdentity: true,
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedNodesConfig: "        - --nodes=1:10:k8s-pool1-49584119-vmss\n        - --nodes=1:10:k8s-pool2-49584119-vmss",
		},
		{
			name: "no pools",
			addon: KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
			},
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
							UseManagedIdentity: true,
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedNodesConfig: "",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := GetClusterAutoscalerNodesConfig(c.addon, c.cs)
			if c.expectedNodesConfig != result {
				t.Errorf("expected GetClusterAutoscalerNodesConfig to return %s, instead got %s", c.expectedNodesConfig, result)
			}
		})
	}
}

func concatenateDefaultAddons(addons []KubernetesAddon, version string) []KubernetesAddon {
	defaults := getDefaultAddons(version, "", common.KubernetesImageBaseTypeMCR)
	defaults = append(defaults, addons...)
	return defaults
}

func overwriteDefaultAddons(addons []KubernetesAddon, version string) []KubernetesAddon {
	overrideAddons := make(map[string]KubernetesAddon)
	for _, addonOverride := range addons {
		overrideAddons[addonOverride.Name] = addonOverride
	}

	var ret []KubernetesAddon
	defaults := getDefaultAddons(version, "", common.KubernetesImageBaseTypeMCR)

	for _, addon := range defaults {
		if _, exists := overrideAddons[addon.Name]; exists {
			ret = append(ret, overrideAddons[addon.Name])
			continue
		}
		ret = append(ret, addon)
	}

	return ret
}

func omitFromAddons(addons []string, completeSet []KubernetesAddon) []KubernetesAddon {
	var ret []KubernetesAddon
	for _, addon := range completeSet {
		if !isInStrSlice(addon.Name, addons) {
			ret = append(ret, addon)
		}
	}
	return ret
}

func isInStrSlice(name string, names []string) bool {
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}

func getDefaultAddons(version, kubernetesImageBase, kubernetesImageBaseType string) []KubernetesAddon {
	specConfig := AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	imageBase := specConfig.MCRKubernetesImageBase
	if kubernetesImageBase != "" {
		imageBase = kubernetesImageBase
	}
	k8sComponentsByVersionMap := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: kubernetesImageBaseType})
	metricsServerAddon := KubernetesAddon{
		Name:    common.MetricsServerAddonName,
		Enabled: to.BoolPtr(true),
		Mode:    AddonModeEnsureExists,
		Containers: []KubernetesContainerSpec{
			{
				Name:  common.MetricsServerAddonName,
				Image: imageBase + k8sComponentsByVersionMap[version][common.MetricsServerAddonName],
			},
		},
	}
	if common.IsKubernetesVersionGe(version, "1.16.0") {
		metricsServerAddon.Mode = AddonModeReconcile
	}
	addons := []KubernetesAddon{
		{
			Name:    common.BlobfuseFlexVolumeAddonName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:           common.BlobfuseFlexVolumeAddonName,
					CPURequests:    "50m",
					MemoryRequests: "100Mi",
					CPULimits:      "50m",
					MemoryLimits:   "100Mi",
					Image:          k8sComponentsByVersionMap[version][common.BlobfuseFlexVolumeAddonName],
				},
			},
		},
		{
			Name:    common.DashboardAddonName,
			Enabled: to.BoolPtr(false),
			Containers: []KubernetesContainerSpec{
				{
					Name:           common.DashboardAddonName,
					CPURequests:    "300m",
					MemoryRequests: "150Mi",
					CPULimits:      "300m",
					MemoryLimits:   "150Mi",
					Image:          imageBase + k8sComponentsByVersionMap[version][common.DashboardAddonName],
				},
			},
		},
		metricsServerAddon,
		{
			Name:    common.IPMASQAgentAddonName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:           common.IPMASQAgentAddonName,
					CPURequests:    "50m",
					MemoryRequests: "50Mi",
					CPULimits:      "50m",
					MemoryLimits:   "250Mi",
					Image:          imageBase + k8sComponentsByVersionMap[version][common.IPMASQAgentAddonName],
				},
			},
			Config: map[string]string{
				"non-masquerade-cidr": DefaultVNETCIDR,
				"non-masq-cni-cidr":   DefaultCNICIDR,
				"enable-ipv6":         "false",
			},
		},
		{
			Name:    common.AzureCNINetworkMonitorAddonName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:           common.AzureCNINetworkMonitorAddonName,
					Image:          specConfig.AzureCNIImageBase + k8sComponentsByVersionMap[version][common.AzureCNINetworkMonitorAddonName],
					CPURequests:    "30m",
					MemoryRequests: "25Mi",
					CPULimits:      "200m",
					MemoryLimits:   "256Mi",
				},
			},
		},
		{
			Name:    common.AuditPolicyAddonName,
			Enabled: to.BoolPtr(true),
		},
		{
			Name:    common.AzureCloudProviderAddonName,
			Enabled: to.BoolPtr(true),
		},
		{
			Name:    common.CoreDNSAddonName,
			Enabled: to.BoolPtr(DefaultCoreDNSAddonEnabled),
			Config: map[string]string{
				"domain":            "cluster.local",
				"clusterIP":         DefaultKubernetesDNSServiceIP,
				"cores-per-replica": "512",
				"nodes-per-replica": "32",
				"min-replicas":      "1",
			},
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.CoreDNSAddonName,
					Image: imageBase + k8sComponentsByVersionMap[version][common.CoreDNSAddonName],
				},
				{
					Name:  common.CoreDNSAutoscalerName,
					Image: k8sComponentsByVersionMap[version][common.CoreDNSAutoscalerName],
				},
			},
		},
		{
			Name:    common.KubeProxyAddonName,
			Enabled: to.BoolPtr(DefaultKubeProxyAddonEnabled),
			Config: map[string]string{
				"cluster-cidr": DefaultKubernetesSubnet,
				"proxy-mode":   string(KubeProxyModeIPTables),
				"featureGates": "{}",
			},
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.KubeProxyAddonName,
					Image: imageBase + k8sComponentsByVersionMap[version][common.KubeProxyAddonName],
				},
			},
		},
	}

	if common.IsKubernetesVersionGe(version, "1.15.0") {
		addons = append(addons, KubernetesAddon{
			Name:    common.PodSecurityPolicyAddonName,
			Enabled: to.BoolPtr(true),
		})
	}

	if !common.IsKubernetesVersionGe(version, "1.16.0") {
		addons = append(addons, KubernetesAddon{
			Name:    common.KeyVaultFlexVolumeAddonName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:           common.KeyVaultFlexVolumeAddonName,
					CPURequests:    "50m",
					MemoryRequests: "100Mi",
					CPULimits:      "50m",
					MemoryLimits:   "100Mi",
					Image:          k8sComponentsByVersionMap[version][common.KeyVaultFlexVolumeAddonName],
				},
			},
		})
	}

	if common.IsKubernetesVersionGe(version, "1.16.0") {
		addons = append(addons, KubernetesAddon{
			Name:    common.SecretsStoreCSIDriverAddonName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:           common.CSILivenessProbeContainerName,
					Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[version][common.CSILivenessProbeContainerName],
					CPURequests:    "10m",
					MemoryRequests: "20Mi",
					CPULimits:      "200m",
					MemoryLimits:   "200Mi",
				},
				{
					Name:           common.CSINodeDriverRegistrarContainerName,
					Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[version][common.CSINodeDriverRegistrarContainerName],
					CPURequests:    "10m",
					MemoryRequests: "20Mi",
					CPULimits:      "200m",
					MemoryLimits:   "200Mi",
				},
				{
					Name:           common.CSISecretsStoreDriverContainerName,
					Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[version][common.CSISecretsStoreDriverContainerName],
					CPURequests:    "50m",
					MemoryRequests: "100Mi",
					CPULimits:      "200m",
					MemoryLimits:   "200Mi",
				},
				{
					Name:           common.CSISecretsStoreProviderAzureContainerName,
					Image:          specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[version][common.CSISecretsStoreProviderAzureContainerName],
					CPURequests:    "50m",
					MemoryRequests: "100Mi",
					CPULimits:      "200m",
					MemoryLimits:   "200Mi",
				},
			},
		})
	}

	return addons
}

func TestKubeProxyImageSuffix(t *testing.T) {
	cases := []struct {
		name       string
		cs         ContainerService
		azurestack bool
		expected   string
	}{
		{
			name:       "return empty string if target cloud is NOT Azure Stack",
			cs:         getMockBaseContainerService("1.17.0"),
			azurestack: false,
			expected:   "",
		},
		{
			name:       "return '-azs' if target cloud is Azure Stack",
			cs:         getMockBaseContainerService("1.17.0"),
			azurestack: true,
			expected:   common.AzureStackSuffix,
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.azurestack {
				c.cs.Properties.CustomCloudProfile = &CustomCloudProfile{}
			}
			actual := kubeProxyImageSuffix(c.cs)
			if !strings.EqualFold(actual, c.expected) {
				t.Errorf("expected %s to be %s", actual, c.expected)
			}
		})
	}
}

func TestGetCSISidecarComponent(t *testing.T) {
	k8sComponentsByVersionMap := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	k8sComponents := k8sComponentsByVersionMap["1.18.0"]
	cases := []struct {
		name                        string
		csiDriverName               string
		csiSidecarName              string
		expectedCSISidecarComponent string
	}{
		{
			name:                        "unknown csi driver name",
			csiDriverName:               "unknown",
			csiSidecarName:              "unknown",
			expectedCSISidecarComponent: "",
		},
		{
			name:                        "unknown csi sidecar name",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              "unknown",
			expectedCSISidecarComponent: "",
		},
		{
			name:                        "get csi-provisioner image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSIProvisionerContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSIProvisionerContainerName],
		},
		{
			name:                        "get csi-attacher image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSIAttacherContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSIAttacherContainerName],
		},
		{
			name:                        "get livenessprobe image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSILivenessProbeContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSILivenessProbeContainerName],
		},
		{
			name:                        "get livenessprobe-windows image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSILivenessProbeWindowsContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSILivenessProbeWindowsContainerName],
		},
		{
			name:                        "get node-driver-registrar image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSINodeDriverRegistrarContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSINodeDriverRegistrarContainerName],
		},
		{
			name:                        "get node-driver-registrar-windows image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSINodeDriverRegistrarWindowsContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSINodeDriverRegistrarWindowsContainerName],
		},
		{
			name:                        "get csi-resizer image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSIResizerContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSIResizerContainerName],
		},
		{
			name:                        "get csi-snapshotter image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSISnapshotterContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSISnapshotterContainerName],
		},
		{
			name:                        "get snapshot-controller image for azuredisk-csi-driver",
			csiDriverName:               common.AzureDiskCSIDriverAddonName,
			csiSidecarName:              common.CSISnapshotControllerContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSISnapshotControllerContainerName],
		},
		{
			name:                        "get csi-provisioner image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSIProvisionerContainerName,
			expectedCSISidecarComponent: csiSidecarComponentsOverrides[common.AzureFileCSIDriverAddonName][common.CSIProvisionerContainerName],
		},
		{
			name:                        "get csi-attacher image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSIAttacherContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSIAttacherContainerName],
		},
		{
			name:                        "get livenessprobe image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSILivenessProbeContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSILivenessProbeContainerName],
		},
		{
			name:                        "get livenessprobe-windows image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSILivenessProbeWindowsContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSILivenessProbeWindowsContainerName],
		},
		{
			name:                        "get node-driver-registrar image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSINodeDriverRegistrarContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSINodeDriverRegistrarContainerName],
		},
		{
			name:                        "get node-driver-registrar-windows image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSINodeDriverRegistrarWindowsContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSINodeDriverRegistrarWindowsContainerName],
		},
		{
			name:                        "get csi-resizer image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSIResizerContainerName,
			expectedCSISidecarComponent: k8sComponents[common.CSIResizerContainerName],
		},
		{
			name:                        "get csi-snapshotter image for azurefile-csi-driver",
			csiDriverName:               common.AzureFileCSIDriverAddonName,
			csiSidecarName:              common.CSISnapshotterContainerName,
			expectedCSISidecarComponent: csiSidecarComponentsOverrides[common.AzureFileCSIDriverAddonName][common.CSISnapshotterContainerName],
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := getCSISidecarComponent(c.csiDriverName, c.csiSidecarName, k8sComponents)
			if c.expectedCSISidecarComponent != actual {
				t.Errorf("expected %s to be %s", c.expectedCSISidecarComponent, actual)
			}
		})
	}
}
