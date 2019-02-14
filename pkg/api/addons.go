// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api/common"
)

func (cs *ContainerService) setAddonsConfig(isUpdate bool) {
	o := cs.Properties.OrchestratorProfile
	cloudSpecConfig := cs.GetCloudSpecConfig()
	k8sComponents := K8sComponentsByVersionMap[o.OrchestratorVersion]
	specConfig := cloudSpecConfig.KubernetesSpecConfig
	defaultsHeapsterAddonsConfig := KubernetesAddon{
		Name:    DefaultHeapsterAddonName,
		Enabled: to.BoolPtr(DefaultHeapsterAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:  DefaultHeapsterAddonName,
				Image: specConfig.KubernetesImageBase + k8sComponents["heapster"],
			},
			{
				Name:  "heapster-nanny",
				Image: specConfig.KubernetesImageBase + k8sComponents["addonresizer"],
			},
		},
	}

	defaultTillerAddonsConfig := KubernetesAddon{
		Name:    DefaultTillerAddonName,
		Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultTillerAddonName,
				CPURequests:    "50m",
				MemoryRequests: "150Mi",
				CPULimits:      "50m",
				MemoryLimits:   "150Mi",
				Image:          specConfig.TillerImageBase + k8sComponents[DefaultTillerAddonName],
			},
		},
		Config: map[string]string{
			"max-history": strconv.Itoa(DefaultTillerMaxHistory),
		},
	}

	defaultACIConnectorAddonsConfig := KubernetesAddon{
		Name:    DefaultACIConnectorAddonName,
		Enabled: to.BoolPtr(DefaultACIConnectorAddonEnabled),
		Config: map[string]string{
			"region":   "westus",
			"nodeName": "aci-connector",
			"os":       "Linux",
			"taint":    "azure.com/aci",
		},
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultACIConnectorAddonName,
				CPURequests:    "50m",
				MemoryRequests: "150Mi",
				CPULimits:      "50m",
				MemoryLimits:   "150Mi",
				Image:          specConfig.ACIConnectorImageBase + k8sComponents[DefaultACIConnectorAddonName],
			},
		},
	}

	defaultClusterAutoscalerAddonsConfig := KubernetesAddon{
		Name:    DefaultClusterAutoscalerAddonName,
		Enabled: to.BoolPtr(DefaultClusterAutoscalerAddonEnabled),
		Config: map[string]string{
			"min-nodes":     "1",
			"max-nodes":     "5",
			"scan-interval": "10s",
		},
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultClusterAutoscalerAddonName,
				CPURequests:    "100m",
				MemoryRequests: "300Mi",
				CPULimits:      "100m",
				MemoryLimits:   "300Mi",
				Image:          specConfig.KubernetesImageBase + k8sComponents[DefaultClusterAutoscalerAddonName],
			},
		},
	}

	defaultBlobfuseFlexVolumeAddonsConfig := KubernetesAddon{
		Name:    DefaultBlobfuseFlexVolumeAddonName,
		Enabled: to.BoolPtr(common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.8.0") && DefaultBlobfuseFlexVolumeAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultBlobfuseFlexVolumeAddonName,
				CPURequests:    "50m",
				MemoryRequests: "10Mi",
				CPULimits:      "50m",
				MemoryLimits:   "10Mi",
				Image:          "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
			},
		},
	}

	defaultSMBFlexVolumeAddonsConfig := KubernetesAddon{
		Name:    DefaultSMBFlexVolumeAddonName,
		Enabled: to.BoolPtr(common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.8.0") && DefaultSMBFlexVolumeAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultSMBFlexVolumeAddonName,
				CPURequests:    "50m",
				MemoryRequests: "10Mi",
				CPULimits:      "50m",
				MemoryLimits:   "10Mi",
				Image:          "mcr.microsoft.com/k8s/flexvolume/smb-flexvolume:1.0.2",
			},
		},
	}

	defaultKeyVaultFlexVolumeAddonsConfig := KubernetesAddon{
		Name:    DefaultKeyVaultFlexVolumeAddonName,
		Enabled: to.BoolPtr(DefaultKeyVaultFlexVolumeAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultKeyVaultFlexVolumeAddonName,
				CPURequests:    "50m",
				MemoryRequests: "10Mi",
				CPULimits:      "50m",
				MemoryLimits:   "10Mi",
				Image:          "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.7",
			},
		},
	}

	defaultDashboardAddonsConfig := KubernetesAddon{
		Name:    DefaultDashboardAddonName,
		Enabled: to.BoolPtr(DefaultDashboardAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultDashboardAddonName,
				CPURequests:    "300m",
				MemoryRequests: "150Mi",
				CPULimits:      "300m",
				MemoryLimits:   "150Mi",
				Image:          specConfig.KubernetesImageBase + k8sComponents[DefaultDashboardAddonName],
			},
		},
	}

	defaultReschedulerAddonsConfig := KubernetesAddon{
		Name:    DefaultReschedulerAddonName,
		Enabled: to.BoolPtr(DefaultReschedulerAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultReschedulerAddonName,
				CPURequests:    "10m",
				MemoryRequests: "100Mi",
				CPULimits:      "10m",
				MemoryLimits:   "100Mi",
				Image:          specConfig.KubernetesImageBase + k8sComponents[DefaultReschedulerAddonName],
			},
		},
	}

	defaultMetricsServerAddonsConfig := KubernetesAddon{
		Name:    DefaultMetricsServerAddonName,
		Enabled: k8sVersionMetricsServerAddonEnabled(o),
		Containers: []KubernetesContainerSpec{
			{
				Name:  DefaultMetricsServerAddonName,
				Image: specConfig.KubernetesImageBase + k8sComponents[DefaultMetricsServerAddonName],
			},
		},
	}

	defaultNVIDIADevicePluginAddonsConfig := KubernetesAddon{
		Name:    NVIDIADevicePluginAddonName,
		Enabled: to.BoolPtr(cs.Properties.HasNSeriesSKU() && common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.10.0")),
		Containers: []KubernetesContainerSpec{
			{
				Name: NVIDIADevicePluginAddonName,
				// from https://github.com/kubernetes/kubernetes/blob/master/cluster/addons/device-plugins/nvidia-gpu/daemonset.yaml#L44
				CPURequests:    "50m",
				MemoryRequests: "10Mi",
				CPULimits:      "50m",
				MemoryLimits:   "10Mi",
				Image:          specConfig.NVIDIAImageBase + k8sComponents[NVIDIADevicePluginAddonName],
			},
		},
	}

	defaultContainerMonitoringAddonsConfig := KubernetesAddon{
		Name:    ContainerMonitoringAddonName,
		Enabled: to.BoolPtr(DefaultContainerMonitoringAddonEnabled),
		Config: map[string]string{
			"omsAgentVersion":       "1.8.1.256",
			"dockerProviderVersion": "3.0.0-3",
		},
		Containers: []KubernetesContainerSpec{
			{
				Name:           "omsagent",
				CPURequests:    "50m",
				MemoryRequests: "200Mi",
				CPULimits:      "150m",
				MemoryLimits:   "750Mi",
				Image:          "microsoft/oms:ciprod01092019",
			},
		},
	}

	defaultIPMasqAgentAddonsConfig := KubernetesAddon{
		Name:    IPMASQAgentAddonName,
		Enabled: to.BoolPtr(IPMasqAgentAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           IPMASQAgentAddonName,
				CPURequests:    "50m",
				MemoryRequests: "50Mi",
				CPULimits:      "50m",
				MemoryLimits:   "250Mi",
				Image:          specConfig.KubernetesImageBase + "ip-masq-agent-amd64:v2.0.0",
			},
		},
		Config: map[string]string{
			"non-masquerade-cidr": cs.Properties.GetNonMasqueradeCIDR(),
			"non-masq-cni-cidr":   cs.Properties.GetAzureCNICidr(),
		},
	}

	defaultAzureCNINetworkMonitorAddonsConfig := KubernetesAddon{
		Name:    AzureCNINetworkMonitoringAddonName,
		Enabled: azureCNINetworkMonitorAddonEnabled(o),
		Containers: []KubernetesContainerSpec{
			{
				Name:  AzureCNINetworkMonitoringAddonName,
				Image: specConfig.AzureCNIImageBase + k8sComponents[AzureCNINetworkMonitoringAddonName],
			},
		},
	}

	defaultAzureNetworkPolicyAddonsConfig := KubernetesAddon{
		Name:    AzureNetworkPolicyAddonName,
		Enabled: azureNetworkPolicyAddonEnabled(o),
		Containers: []KubernetesContainerSpec{
			{
				Name: AzureNetworkPolicyAddonName,
			},
		},
	}

	defaultDNSAutoScalerAddonsConfig := KubernetesAddon{
		Name:    DefaultDNSAutoscalerAddonName,
		Enabled: to.BoolPtr(DefaultDNSAutoscalerAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DefaultDNSAutoscalerAddonName,
				Image:          specConfig.KubernetesImageBase + "cluster-proportional-autoscaler-amd64:1.1.1",
				CPURequests:    "20m",
				MemoryRequests: "10Mi",
			},
		},
	}

	defaultAddons := []KubernetesAddon{
		defaultsHeapsterAddonsConfig,
		defaultTillerAddonsConfig,
		defaultACIConnectorAddonsConfig,
		defaultClusterAutoscalerAddonsConfig,
		defaultBlobfuseFlexVolumeAddonsConfig,
		defaultSMBFlexVolumeAddonsConfig,
		defaultKeyVaultFlexVolumeAddonsConfig,
		defaultDashboardAddonsConfig,
		defaultReschedulerAddonsConfig,
		defaultMetricsServerAddonsConfig,
		defaultNVIDIADevicePluginAddonsConfig,
		defaultContainerMonitoringAddonsConfig,
		defaultAzureCNINetworkMonitorAddonsConfig,
		defaultAzureNetworkPolicyAddonsConfig,
		defaultIPMasqAgentAddonsConfig,
		defaultDNSAutoScalerAddonsConfig,
	}
	// Add default addons specification, if no user-provided spec exists
	if o.KubernetesConfig.Addons == nil {
		o.KubernetesConfig.Addons = defaultAddons
	} else {
		for _, addon := range defaultAddons {
			i := getAddonsIndexByName(o.KubernetesConfig.Addons, addon.Name)
			if i < 0 {
				o.KubernetesConfig.Addons = append(o.KubernetesConfig.Addons, addon)
			}
		}
	}

	for _, addon := range defaultAddons {
		synthesizeAddonsConfig(o.KubernetesConfig.Addons, addon, false, isUpdate)
	}
}

func getAddonsIndexByName(addons []KubernetesAddon, name string) int {
	for i := range addons {
		if addons[i].Name == name {
			return i
		}
	}
	return -1
}

// assignDefaultAddonVals will assign default values to addon from defaults, for each property in addon that has a zero value
func assignDefaultAddonVals(addon, defaults KubernetesAddon, isUpdate bool) KubernetesAddon {
	if addon.Enabled == nil {
		addon.Enabled = defaults.Enabled
	}
	for i := range defaults.Containers {
		c := addon.GetAddonContainersIndexByName(defaults.Containers[i].Name)
		if c < 0 {
			addon.Containers = append(addon.Containers, defaults.Containers[i])
		} else {
			if addon.Containers[c].Image == "" || isUpdate {
				addon.Containers[c].Image = defaults.Containers[i].Image
			}
			if addon.Containers[c].CPURequests == "" {
				addon.Containers[c].CPURequests = defaults.Containers[i].CPURequests
			}
			if addon.Containers[c].MemoryRequests == "" {
				addon.Containers[c].MemoryRequests = defaults.Containers[i].MemoryRequests
			}
			if addon.Containers[c].CPULimits == "" {
				addon.Containers[c].CPULimits = defaults.Containers[i].CPULimits
			}
			if addon.Containers[c].MemoryLimits == "" {
				addon.Containers[c].MemoryLimits = defaults.Containers[i].MemoryLimits
			}
		}
	}
	for key, val := range defaults.Config {
		if addon.Config == nil {
			addon.Config = make(map[string]string)
		}
		if v, ok := addon.Config[key]; !ok || v == "" {
			addon.Config[key] = val
		}
	}
	return addon
}

func synthesizeAddonsConfig(addons []KubernetesAddon, addon KubernetesAddon, enableIfNil bool, isUpdate bool) {
	i := getAddonsIndexByName(addons, addon.Name)
	if i >= 0 {
		if addons[i].IsEnabled(enableIfNil) {
			addons[i] = assignDefaultAddonVals(addons[i], addon, isUpdate)
		}
	}
}

func k8sVersionMetricsServerAddonEnabled(o *OrchestratorProfile) *bool {
	return to.BoolPtr(common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.9.0"))
}

func azureNetworkPolicyAddonEnabled(o *OrchestratorProfile) *bool {
	return to.BoolPtr(o.KubernetesConfig.NetworkPlugin == NetworkPluginAzure && o.KubernetesConfig.NetworkPolicy == NetworkPolicyAzure)
}

func azureCNINetworkMonitorAddonEnabled(o *OrchestratorProfile) *bool {
	return to.BoolPtr(o.IsAzureCNI())
}
