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

	isUpdate := true
	base64DataPSP := "cHNwQ3VzdG9tRGF0YQ=="
	o.OrchestratorType = Kubernetes
	o.KubernetesConfig.EnablePodSecurityPolicy = to.BoolPtr(true)
	o.KubernetesConfig.PodSecurityPolicyConfig = map[string]string{
		"data": base64DataPSP,
	}

	mockCS.setAddonsConfig(isUpdate)

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
		isUpdate            bool
		expectedResultAddon KubernetesAddon
	}{
		{
			name: "default addon enabled",
			myAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(true),
			},
			isUpdate: false,
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
			name: "addon disabled, isUpdate=false",
			myAddon: KubernetesAddon{
				Name: "mockAddon",
			},
			isUpdate: false,
			expectedResultAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
		},
		{
			name: "addon disabled, isUpdate=true",
			myAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
			isUpdate: true,
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
			result := assignDefaultAddonVals(c.myAddon, defaultAddon, c.isUpdate)
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
		isUpdate       bool
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
			isUpdate: false,
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
			isUpdate: false,
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
			isUpdate: false,
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
			},
		},
		{
			name: "cluster-autoscaler addon enabled",
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
				},
			},
			isUpdate: false,
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
						"min-nodes":     "1",
						"max-nodes":     "5",
						"scan-interval": "10s",
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
			isUpdate: false,
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
			isUpdate: false,
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
			isUpdate: false,
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
			isUpdate: false,
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
						"dockerProviderVersion": "6.0.0-0",
						"schema-versions":       "v1",
						"clusterName":           "aks-engine-cluster",
					},
					Containers: []KubernetesContainerSpec{
						{
							Name:           "omsagent",
							CPURequests:    "75m",
							MemoryRequests: "225Mi",
							CPULimits:      "150m",
							MemoryLimits:   "600Mi",
							Image:          "mcr.microsoft.com/azuremonitor/containerinsights/ciprod:ciprod07092019",
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
			isUpdate: false,
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
							Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.28",
						},
						{
							Name:  AzureVnetTelemetryAddonName,
							Image: "mcr.microsoft.com/containernetworking/azure-vnet-telemetry:v1.0.28",
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
			isUpdate: false,
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
							Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.28",
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
											Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.28",
										},
										{
											Name:  AzureVnetTelemetryAddonName,
											Image: "mcr.microsoft.com/containernetworking/azure-vnet-telemetry:v1.0.28",
										},
									},
								},
							},
						},
					},
				},
			},
			isUpdate: true,
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
							Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.28",
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
			isUpdate: false,
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
			isUpdate: false,
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
			isUpdate: true,
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
			isUpdate: false,
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
			isUpdate: false,
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
			isUpdate: false,
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
			isUpdate: false,
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
			isUpdate: false,
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
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.cs.setAddonsConfig(test.isUpdate)
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
			}
		})
	}
}
