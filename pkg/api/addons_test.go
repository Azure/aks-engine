// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"reflect"
	"strconv"
	"testing"

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
	mockCS := getMockBaseContainerService("1.8.0")
	o := mockCS.Properties.OrchestratorProfile

	isUpgrade := true
	base64DataPSP := "cHNwQ3VzdG9tRGF0YQ=="
	o.OrchestratorType = Kubernetes
	o.KubernetesConfig.EnablePodSecurityPolicy = to.BoolPtr(true)
	o.KubernetesConfig.PodSecurityPolicyConfig = map[string]string{
		"data": base64DataPSP,
	}

	mockCS.setAddonsConfig(isUpgrade)

	i := getAddonsIndexByName(o.KubernetesConfig.Addons, PodSecurityPolicyAddonName)
	if i < 0 {
		t.Errorf("expected a positive index for the addon %s, instead got %d from getAddonsIndexByName", PodSecurityPolicyAddonName, i)
	}

	if o.KubernetesConfig.Addons[i].Name != PodSecurityPolicyAddonName {
		t.Errorf("expected addon %s name to be present, instead got %s", PodSecurityPolicyAddonName, o.KubernetesConfig.Addons[i].Name)
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
		},
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
		},
	}
	AzureCloudSpecEnvMap[AzureStackCloud] = azureStackCloudSpec
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
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "tiller addon is enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    TillerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           TillerAddonName,
							CPURequests:    "50m",
							MemoryRequests: "150Mi",
							CPULimits:      "50m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.TillerImageBase + K8sComponentsByVersionMap["1.12.8"][TillerAddonName],
						},
					},
					Config: map[string]string{
						"max-history": strconv.Itoa(0),
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "ACI Connector addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ACIConnectorAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"region":   "westus",
						"nodeName": "aci-connector",
						"os":       "Linux",
						"taint":    "azure.com/aci",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           ACIConnectorAddonName,
							CPURequests:    "50m",
							MemoryRequests: "150Mi",
							CPULimits:      "50m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.ACIConnectorImageBase + K8sComponentsByVersionMap["1.12.8"][ACIConnectorAddonName],
						},
					},
				},
				{
					Name:    ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled - 1.12",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "1m",
						"expendable-pods-priority-cutoff":       "-10",
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled - 1.11",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.11.10",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.11.10"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.11.10"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"scan-interval":                         "1m",
						"expendable-pods-priority-cutoff":       "-10",
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.11.10"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.11.10"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.11.10"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.11.10"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled on cluster with multiple pools",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled on cluster with multiple pools - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled with configuration",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled with configuration - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled with mixed configuration plus defaults",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "cluster-autoscaler addon enabled with mixed configuration plus defaults - update",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.11",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ClusterAutoscalerAddonName,
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
				},
				{
					Name:    ACIConnectorAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ClusterAutoscalerAddonName,
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
							Name:           ClusterAutoscalerAddonName,
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][ClusterAutoscalerAddonName],
						},
					},
				},
				{
					Name:    BlobfuseFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.11"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.11"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "smb-flexvolume addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    SMBFlexVolumeAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           SMBFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/smb-flexvolume:1.0.2",
						},
					},
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "rescheduler addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ReschedulerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           ReschedulerAddonName,
							CPURequests:    "10m",
							MemoryRequests: "100Mi",
							CPULimits:      "10m",
							MemoryLimits:   "100Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][ReschedulerAddonName],
						},
					},
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "nvidia addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
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
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
				},
				{
					Name:    NVIDIADevicePluginAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           NVIDIADevicePluginAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          specConfig.NVIDIAImageBase + K8sComponentsByVersionMap["1.12.8"][NVIDIADevicePluginAddonName],
						},
					},
				},
				{
					Name:    ContainerMonitoringAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    IPMASQAgentAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "container-monitoring addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    ContainerMonitoringAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
				},
				{
					Name:    NVIDIADevicePluginAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    ContainerMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"omsAgentVersion":       "1.10.0.1",
						"dockerProviderVersion": "7.0.0-6",
						"schema-versions":       "v1",
						"clusterName":           "aks-engine-cluster",
						"workspaceDomain":       "b3BpbnNpZ2h0cy5henVyZS5jb20=",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           "omsagent",
							CPURequests:    "110m",
							MemoryRequests: "250Mi",
							CPULimits:      "150m",
							MemoryLimits:   "600Mi",
							Image:          "mcr.microsoft.com/azuremonitor/containerinsights/ciprod:ciprod11012019",
						},
					},
				},
				{
					Name:    IPMASQAgentAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "Azure Network Policy addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyAzure,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureNetworkPolicyAddonName,
							Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.29",
						},
						{
							Name:  AzureVnetTelemetryAddonName,
							Image: "mcr.microsoft.com/containernetworking/azure-vnet-telemetry:v1.0.29",
						},
					},
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "Azure Network Policy addon enabled - 1.16",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.0",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyAzure,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.16.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureNetworkPolicyAddonName,
							Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.29",
						},
					},
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "Azure Network Policy addon enabled - 1.16 upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.0",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyAzure,
							Addons: []KubernetesAddon{
								{
									Name:    AzureNetworkPolicyAddonName,
									Enabled: to.BoolPtr(true),
									Containers: []KubernetesContainerSpec{
										{
											Name:  AzureNetworkPolicyAddonName,
											Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.29",
										},
										{
											Name:  AzureVnetTelemetryAddonName,
											Image: "mcr.microsoft.com/containernetworking/azure-vnet-telemetry:v1.0.29",
										},
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.16.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureNetworkPolicyAddonName,
							Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.29",
						},
					},
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "dns-autoscaler addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    DNSAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    DNSAutoscalerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DNSAutoscalerAddonName,
							Image:          specConfig.KubernetesImageBase + "cluster-proportional-autoscaler-amd64:1.1.1",
							CPURequests:    "20m",
							MemoryRequests: "100Mi",
						},
					},
				},
				{
					Name:    CalicoAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "calico addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyCalico,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    DNSAutoscalerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CalicoAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "calico-typha",
							Image: specConfig.CalicoImageBase + "typha:v3.8.0",
						},
						{
							Name:  "calico-cni",
							Image: specConfig.CalicoImageBase + "cni:v3.8.0",
						},
						{
							Name:  "calico-node",
							Image: specConfig.CalicoImageBase + "node:v3.8.0",
						},
						{
							Name:  "calico-pod2daemon",
							Image: specConfig.CalicoImageBase + "pod2daemon-flexvol:v3.8.0",
						},
						{
							Name:  "calico-cluster-proportional-autoscaler",
							Image: specConfig.KubernetesImageBase + "cluster-proportional-autoscaler-amd64:1.1.2-r2",
						},
					},
				},
				{
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "calico addon back-compat",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							NetworkPolicy: NetworkPolicyCalico,
							Addons: []KubernetesAddon{
								{
									Name:    CalicoAddonName,
									Enabled: to.BoolPtr(false),
									Containers: []KubernetesContainerSpec{
										{
											Name:  "calico-typha",
											Image: specConfig.CalicoImageBase + "typha:old", // confirm that upgrade will change this to default image
										},
										{
											Name:  "calico-cni",
											Image: specConfig.CalicoImageBase + "cni:v3.8.0",
										},
										{
											Name:  "calico-node",
											Image: specConfig.CalicoImageBase + "node:v3.8.0",
										},
										{
											Name:  "calico-pod2daemon",
											Image: specConfig.CalicoImageBase + "pod2daemon-flexvol:v3.8.0",
										},
										{
											Name:  "calico-cluster-proportional-autoscaler",
											Image: specConfig.KubernetesImageBase + "cluster-proportional-autoscaler-amd64:1.1.2-r2",
										},
									},
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureNetworkPolicyAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    DNSAutoscalerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CalicoAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "calico-typha",
							Image: specConfig.CalicoImageBase + "typha:v3.8.0",
						},
						{
							Name:  "calico-cni",
							Image: specConfig.CalicoImageBase + "cni:v3.8.0",
						},
						{
							Name:  "calico-node",
							Image: specConfig.CalicoImageBase + "node:v3.8.0",
						},
						{
							Name:  "calico-pod2daemon",
							Image: specConfig.CalicoImageBase + "pod2daemon-flexvol:v3.8.0",
						},
						{
							Name:  "calico-cluster-proportional-autoscaler",
							Image: specConfig.KubernetesImageBase + "cluster-proportional-autoscaler-amd64:1.1.2-r2",
						},
					},
				},
				{
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "aad-pod-identity enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    AADPodIdentityAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           "nmi",
							Image:          "mcr.microsoft.com/k8s/aad-pod-identity/nmi:1.2",
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
						},
						{
							Name:           "mic",
							Image:          "mcr.microsoft.com/k8s/aad-pod-identity/mic:1.2",
							CPURequests:    "100m",
							MemoryRequests: "300Mi",
							CPULimits:      "100m",
							MemoryLimits:   "300Mi",
						},
					},
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "azure-policy addon enabled",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
							Addons: []KubernetesAddon{
								{
									Name:    AzurePolicyAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.12.8"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(true),
					Config: map[string]string{
						"auditInterval":             "30",
						"constraintViolationsLimit": "20",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           "azure-policy",
							Image:          "mcr.microsoft.com/azure-policy/policy-kubernetes-addon-prod:prod_20191011.1",
							CPURequests:    "30m",
							MemoryRequests: "50Mi",
							CPULimits:      "100m",
							MemoryLimits:   "200Mi",
						},
						{
							Name:           "gatekeeper",
							Image:          "quay.io/open-policy-agent/gatekeeper:v3.0.4-beta.2",
							CPURequests:    "100m",
							MemoryRequests: "256Mi",
							CPULimits:      "100m",
							MemoryLimits:   "512Mi",
						},
					},
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "cilium addons",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.12.8",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginCilium,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           HeapsterAddonName,
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["heapster"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
						{
							Name:           "heapster-nanny",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"]["addonresizer"],
							CPURequests:    "88m",
							MemoryRequests: "204Mi",
							CPULimits:      "88m",
							MemoryLimits:   "204Mi",
						},
					},
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.12.8"][MetricsServerAddonName],
						},
					},
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
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "k8s >= 1.13 addons",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.0",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "Azure Stack addons",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.14.0",
						KubernetesConfig: &KubernetesConfig{
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
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          "KubernetesImageBase" + K8sComponentsByVersionMap["1.14.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: "KubernetesImageBase" + K8sComponentsByVersionMap["1.14.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          "KubernetesImageBase" + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: "AzureCNIImageBase" + K8sComponentsByVersionMap["1.14.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "CoreOS addons",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.0",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin: NetworkPluginAzure,
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Distro: CoreOS,
							VMSize: "Standard_NC6", // to validate that CoreOS distro does not get nvidia addon
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
			},
		},
		{
			name: "azure disk and azure file csi driver enabled for k8s >= 1.13.0 and UseCloudControllerManager is true",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.13.0",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.13.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.13.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azurefile-csi",
							Image: "mcr.microsoft.com/k8s/csi/azurefile-csi:v0.3.0",
						},
					},
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azuredisk-csi",
							Image: "mcr.microsoft.com/k8s/csi/azuredisk-csi:v0.4.0",
						},
					},
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "azure cloud-node-manager enabled for k8s == 1.16 and useCloudControllerManager is true",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.1",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.1"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.1"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.16.1"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azurefile-csi",
							Image: "mcr.microsoft.com/k8s/csi/azurefile-csi:v0.3.0",
						},
					},
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azuredisk-csi",
							Image: "mcr.microsoft.com/k8s/csi/azuredisk-csi:v0.4.0",
						},
					},
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "azure cloud-node-manager enabled for k8s >= 1.17.0",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.17.0",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
						},
					},
				},
			},
			isUpgrade: false,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.17.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.17.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.17.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azurefile-csi",
							Image: "mcr.microsoft.com/k8s/csi/azurefile-csi:v0.3.0",
						},
					},
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azuredisk-csi",
							Image: "mcr.microsoft.com/k8s/csi/azuredisk-csi:v0.4.0",
						},
					},
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "azure cloud-node-manager enabled for k8s >= 1.17.0 - upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.17.0",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
							Addons: []KubernetesAddon{
								{
									Name:    AzureDiskCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    AzureFileCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    CloudNodeManagerAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.17.0"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.17.0"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.17.0"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azurefile-csi",
							Image: "mcr.microsoft.com/k8s/csi/azurefile-csi:v0.3.0",
						},
					},
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azuredisk-csi",
							Image: "mcr.microsoft.com/k8s/csi/azuredisk-csi:v0.4.0",
						},
					},
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
		{
			name: "azure cloud-node-manager enabled for k8s == 1.16.1 upgrade",
			cs: &ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorVersion: "1.16.1",
						KubernetesConfig: &KubernetesConfig{
							NetworkPlugin:             NetworkPluginAzure,
							UseCloudControllerManager: to.BoolPtr(true),
							Addons: []KubernetesAddon{
								{
									Name:    AzureDiskCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    AzureFileCSIDriverAddonName,
									Enabled: to.BoolPtr(false),
								},
								{
									Name:    CloudNodeManagerAddonName,
									Enabled: to.BoolPtr(false),
								},
							},
						},
					},
				},
			},
			isUpgrade: true,
			expectedAddons: []KubernetesAddon{
				{
					Name:    HeapsterAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    TillerAddonName,
					Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           BlobfuseFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
						},
					},
				},
				{
					Name:    SMBFlexVolumeAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    KeyVaultFlexVolumeAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           KeyVaultFlexVolumeAddonName,
							CPURequests:    "50m",
							MemoryRequests: "100Mi",
							CPULimits:      "50m",
							MemoryLimits:   "100Mi",
							Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
						},
					},
				},
				{
					Name:    DashboardAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           DashboardAddonName,
							CPURequests:    "300m",
							MemoryRequests: "150Mi",
							CPULimits:      "300m",
							MemoryLimits:   "150Mi",
							Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.1"][DashboardAddonName],
						},
					},
				},
				{
					Name:    ReschedulerAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    MetricsServerAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  MetricsServerAddonName,
							Image: specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.16.1"][MetricsServerAddonName],
						},
					},
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
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:           IPMASQAgentAddonName,
							CPURequests:    "50m",
							MemoryRequests: "50Mi",
							CPULimits:      "50m",
							MemoryLimits:   "250Mi",
							Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.5.0",
						},
					},
					Config: map[string]string{
						"non-masquerade-cidr": DefaultVNETCIDR,
						"non-masq-cni-cidr":   DefaultCNICIDR,
						"enable-ipv6":         "false",
					},
				},
				{
					Name:    AzureCNINetworkMonitoringAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  AzureCNINetworkMonitoringAddonName,
							Image: specConfig.AzureCNIImageBase + K8sComponentsByVersionMap["1.16.1"][AzureCNINetworkMonitoringAddonName],
						},
					},
				},
				{
					Name:    AzureNetworkPolicyAddonName,
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
					Name:    AADPodIdentityAddonName,
					Enabled: to.BoolPtr(false),
				},
				{
					Name:    AzureFileCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azurefile-csi",
							Image: "mcr.microsoft.com/k8s/csi/azurefile-csi:v0.3.0",
						},
					},
				},
				{
					Name:    AzureDiskCSIDriverAddonName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  "csi-provisioner",
							Image: "quay.io/k8scsi/csi-provisioner:v1.0.1",
						},
						{
							Name:  "csi-attacher",
							Image: "quay.io/k8scsi/csi-attacher:v1.0.1",
						},
						{
							Name:  "csi-cluster-driver-registrar",
							Image: "quay.io/k8scsi/csi-cluster-driver-registrar:v1.0.1",
						},
						{
							Name:  "livenessprobe",
							Image: "quay.io/k8scsi/livenessprobe:v1.1.0",
						},
						{
							Name:  "csi-node-driver-registrar",
							Image: "quay.io/k8scsi/csi-node-driver-registrar:v1.1.0",
						},
						{
							Name:  "azuredisk-csi",
							Image: "mcr.microsoft.com/k8s/csi/azuredisk-csi:v0.4.0",
						},
					},
				},
				{
					Name:    CloudNodeManagerAddonName,
					Enabled: to.BoolPtr(true),
				},
				{
					Name:    AzurePolicyAddonName,
					Enabled: to.BoolPtr(DefaultAzurePolicyAddonEnabled),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.cs.setAddonsConfig(test.isUpgrade)
			for _, addonName := range []string{
				HeapsterAddonName,
				TillerAddonName,
				ACIConnectorAddonName,
				ClusterAutoscalerAddonName,
				BlobfuseFlexVolumeAddonName,
				SMBFlexVolumeAddonName,
				KeyVaultFlexVolumeAddonName,
				DashboardAddonName,
				ReschedulerAddonName,
				MetricsServerAddonName,
				NVIDIADevicePluginAddonName,
				ContainerMonitoringAddonName,
				IPMASQAgentAddonName,
				AzureCNINetworkMonitoringAddonName,
				AzureNetworkPolicyAddonName,
				DNSAutoscalerAddonName,
				CalicoAddonName,
				AADPodIdentityAddonName,
				AzurePolicyAddonName,
				AzureFileCSIDriverAddonName,
				AzureDiskCSIDriverAddonName,
				CloudNodeManagerAddonName,
			} {
				addon := test.cs.Properties.OrchestratorProfile.KubernetesConfig.Addons[getAddonsIndexByName(test.cs.Properties.OrchestratorProfile.KubernetesConfig.Addons, addonName)]
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
	cases := []struct {
		name                string
		addon               KubernetesAddon
		cs                  *ContainerService
		expectedNodesConfig string
	}{
		{
			name: "1 pool",
			addon: KubernetesAddon{
				Name:    ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []KubernetesContainerSpec{
					{
						Name:           ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.15.4"][ClusterAutoscalerAddonName],
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
									Name:    ClusterAutoscalerAddonName,
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
				Name:    ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []KubernetesContainerSpec{
					{
						Name:           ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.15.4"][ClusterAutoscalerAddonName],
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
									Name:    ClusterAutoscalerAddonName,
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
				Name:    ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []KubernetesContainerSpec{
					{
						Name:           ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + K8sComponentsByVersionMap["1.15.4"][ClusterAutoscalerAddonName],
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
									Name:    ClusterAutoscalerAddonName,
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
