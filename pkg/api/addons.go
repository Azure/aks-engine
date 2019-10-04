// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api/common"
)

func (cs *ContainerService) setAddonsConfig(isUpdate bool) {
	o := cs.Properties.OrchestratorProfile
	clusterDNSPrefix := "aks-engine-cluster"
	if cs != nil && cs.Properties != nil && cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.DNSPrefix != "" {
		clusterDNSPrefix = cs.Properties.MasterProfile.DNSPrefix
	}
	cloudSpecConfig := cs.GetCloudSpecConfig()
	k8sComponents := K8sComponentsByVersionMap[o.OrchestratorVersion]
	specConfig := cloudSpecConfig.KubernetesSpecConfig
	omsagentImage := "mcr.microsoft.com/azuremonitor/containerinsights/ciprod:ciprod07092019"
	if strings.EqualFold(cloudSpecConfig.CloudName, "AzureChinaCloud") {
		omsagentImage = "dockerhub.azk8s.cn/microsoft/oms:ciprod07092019"
	}
	workspaceDomain := getLogAnalyticsWorkspaceDomain(cloudSpecConfig.CloudName)
	workspaceDomain = base64.StdEncoding.EncodeToString([]byte(workspaceDomain))

	defaultsHeapsterAddonsConfig := KubernetesAddon{
		Name:    HeapsterAddonName,
		Enabled: to.BoolPtr(DefaultHeapsterAddonEnabled && !common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.13.0")),
		Containers: []KubernetesContainerSpec{
			{
				Name:           HeapsterAddonName,
				Image:          specConfig.KubernetesImageBase + k8sComponents["heapster"],
				CPURequests:    "88m",
				MemoryRequests: "204Mi",
				CPULimits:      "88m",
				MemoryLimits:   "204Mi",
			},
			{
				Name:           "heapster-nanny",
				Image:          specConfig.KubernetesImageBase + k8sComponents["addonresizer"],
				CPURequests:    "88m",
				MemoryRequests: "204Mi",
				CPULimits:      "88m",
				MemoryLimits:   "204Mi",
			},
		},
	}

	defaultTillerAddonsConfig := KubernetesAddon{
		Name:    TillerAddonName,
		Enabled: to.BoolPtr(DefaultTillerAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           TillerAddonName,
				CPURequests:    "50m",
				MemoryRequests: "150Mi",
				CPULimits:      "50m",
				MemoryLimits:   "150Mi",
				Image:          specConfig.TillerImageBase + k8sComponents[TillerAddonName],
			},
		},
		Config: map[string]string{
			"max-history": strconv.Itoa(DefaultTillerMaxHistory),
		},
	}

	defaultACIConnectorAddonsConfig := KubernetesAddon{
		Name:    ACIConnectorAddonName,
		Enabled: to.BoolPtr(DefaultACIConnectorAddonEnabled && !cs.Properties.IsAzureStackCloud()),
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
				Image:          specConfig.ACIConnectorImageBase + k8sComponents[ACIConnectorAddonName],
			},
		},
	}

	defaultClusterAutoscalerAddonsConfig := KubernetesAddon{
		Name:    ClusterAutoscalerAddonName,
		Enabled: to.BoolPtr(DefaultClusterAutoscalerAddonEnabled && !cs.Properties.IsAzureStackCloud()),
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
				Image:          specConfig.KubernetesImageBase + k8sComponents[ClusterAutoscalerAddonName],
			},
		},
	}

	defaultBlobfuseFlexVolumeAddonsConfig := KubernetesAddon{
		Name:    BlobfuseFlexVolumeAddonName,
		Enabled: to.BoolPtr(DefaultBlobfuseFlexVolumeAddonEnabled && common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.8.0") && !cs.Properties.HasCoreOS() && !cs.Properties.IsAzureStackCloud()),
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
	}

	defaultSMBFlexVolumeAddonsConfig := KubernetesAddon{
		Name:    SMBFlexVolumeAddonName,
		Enabled: to.BoolPtr(DefaultSMBFlexVolumeAddonEnabled && common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.8.0") && !cs.Properties.HasCoreOS() && !cs.Properties.IsAzureStackCloud()),
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
	}

	defaultKeyVaultFlexVolumeAddonsConfig := KubernetesAddon{
		Name:    KeyVaultFlexVolumeAddonName,
		Enabled: to.BoolPtr(DefaultKeyVaultFlexVolumeAddonEnabled && !cs.Properties.HasCoreOS() && !cs.Properties.IsAzureStackCloud()),
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
	}

	defaultDashboardAddonsConfig := KubernetesAddon{
		Name:    DashboardAddonName,
		Enabled: to.BoolPtr(DefaultDashboardAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DashboardAddonName,
				CPURequests:    "300m",
				MemoryRequests: "150Mi",
				CPULimits:      "300m",
				MemoryLimits:   "150Mi",
				Image:          specConfig.KubernetesImageBase + k8sComponents[DashboardAddonName],
			},
		},
	}

	defaultReschedulerAddonsConfig := KubernetesAddon{
		Name:    ReschedulerAddonName,
		Enabled: to.BoolPtr(DefaultReschedulerAddonEnabled && !cs.Properties.IsAzureStackCloud()),
		Containers: []KubernetesContainerSpec{
			{
				Name:           ReschedulerAddonName,
				CPURequests:    "10m",
				MemoryRequests: "100Mi",
				CPULimits:      "10m",
				MemoryLimits:   "100Mi",
				Image:          specConfig.KubernetesImageBase + k8sComponents[ReschedulerAddonName],
			},
		},
	}

	defaultMetricsServerAddonsConfig := KubernetesAddon{
		Name:    MetricsServerAddonName,
		Enabled: to.BoolPtr(DefaultMetricsServerAddonEnabled && common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.9.0")),
		Containers: []KubernetesContainerSpec{
			{
				Name:  MetricsServerAddonName,
				Image: specConfig.KubernetesImageBase + k8sComponents[MetricsServerAddonName],
			},
		},
	}

	defaultNVIDIADevicePluginAddonsConfig := KubernetesAddon{
		Name:    NVIDIADevicePluginAddonName,
		Enabled: to.BoolPtr(cs.Properties.IsNvidiaDevicePluginCapable() && !cs.Properties.HasCoreOS() && !cs.Properties.IsAzureStackCloud()),
		Containers: []KubernetesContainerSpec{
			{
				Name: NVIDIADevicePluginAddonName,
				// from https://github.com/kubernetes/kubernetes/blob/master/cluster/addons/device-plugins/nvidia-gpu/daemonset.yaml#L44
				CPURequests:    "50m",
				MemoryRequests: "100Mi",
				CPULimits:      "50m",
				MemoryLimits:   "100Mi",
				Image:          specConfig.NVIDIAImageBase + k8sComponents[NVIDIADevicePluginAddonName],
			},
		},
	}

	defaultContainerMonitoringAddonsConfig := KubernetesAddon{
		Name:    ContainerMonitoringAddonName,
		Enabled: to.BoolPtr(DefaultContainerMonitoringAddonEnabled && !cs.Properties.IsAzureStackCloud()),
		Config: map[string]string{
			"omsAgentVersion":       "1.10.0.1",
			"dockerProviderVersion": "6.0.0-0",
			"schema-versions":       "v1",
			"clusterName":           clusterDNSPrefix,
			"workspaceDomain":       workspaceDomain,
		},
		Containers: []KubernetesContainerSpec{
			{
				Name:           "omsagent",
				CPURequests:    "75m",
				MemoryRequests: "225Mi",
				CPULimits:      "150m",
				MemoryLimits:   "600Mi",
				Image:          omsagentImage,
			},
		},
	}

	defaultIPMasqAgentAddonsConfig := KubernetesAddon{
		Name:    IPMASQAgentAddonName,
		Enabled: to.BoolPtr(DefaultIPMasqAgentAddonEnabled && o.KubernetesConfig.NetworkPlugin != NetworkPluginCilium),
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
			"non-masquerade-cidr":           cs.Properties.GetNonMasqueradeCIDR(),
			"non-masq-cni-cidr":             cs.Properties.GetAzureCNICidr(),
			"secondary-non-masquerade-cidr": cs.Properties.GetSecondaryNonMasqueradeCIDR(),
			"enable-ipv6":                   strconv.FormatBool(cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack")),
		},
	}

	defaultAzureCNINetworkMonitorAddonsConfig := KubernetesAddon{
		Name:    AzureCNINetworkMonitoringAddonName,
		Enabled: to.BoolPtr(o.IsAzureCNI() && o.KubernetesConfig.NetworkPolicy != NetworkPolicyCalico),
		Containers: []KubernetesContainerSpec{
			{
				Name:  AzureCNINetworkMonitoringAddonName,
				Image: specConfig.AzureCNIImageBase + k8sComponents[AzureCNINetworkMonitoringAddonName],
			},
		},
	}

	defaultAzureNetworkPolicyAddonsConfig := KubernetesAddon{
		Name:    AzureNetworkPolicyAddonName,
		Enabled: to.BoolPtr(o.KubernetesConfig.NetworkPlugin == NetworkPluginAzure && o.KubernetesConfig.NetworkPolicy == NetworkPolicyAzure),
		Containers: []KubernetesContainerSpec{
			{
				Name:  AzureNetworkPolicyAddonName,
				Image: "mcr.microsoft.com/containernetworking/azure-npm:v1.0.28",
			},
		},
	}

	if !common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.16.0") {
		defaultAzureNetworkPolicyAddonsConfig.Containers = append(defaultAzureNetworkPolicyAddonsConfig.Containers, KubernetesContainerSpec{Name: AzureVnetTelemetryAddonName, Image: "mcr.microsoft.com/containernetworking/azure-vnet-telemetry:v1.0.28"})
	}

	defaultDNSAutoScalerAddonsConfig := KubernetesAddon{
		Name: DNSAutoscalerAddonName,
		// TODO enable this when it has been smoke tested
		//common.IsKubernetesVersionGe(p.OrchestratorProfile.OrchestratorVersion, "1.12.0"),
		Enabled: to.BoolPtr(DefaultDNSAutoscalerAddonEnabled),
		Containers: []KubernetesContainerSpec{
			{
				Name:           DNSAutoscalerAddonName,
				Image:          specConfig.KubernetesImageBase + "cluster-proportional-autoscaler-amd64:1.1.1",
				CPURequests:    "20m",
				MemoryRequests: "100Mi",
			},
		},
	}

	defaultsCalicoDaemonSetAddonsConfig := KubernetesAddon{
		Name:    CalicoAddonName,
		Enabled: to.BoolPtr(o.KubernetesConfig.NetworkPolicy == NetworkPolicyCalico),
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
	}

	defaultsAADPodIdentityAddonsConfig := KubernetesAddon{
		Name:    AADPodIdentityAddonName,
		Enabled: to.BoolPtr(DefaultAADPodIdentityAddonEnabled && !cs.Properties.IsAzureStackCloud()),
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
	}

	defaultAppGwAddonsConfig := KubernetesAddon{
		Name:    AppGwIngressAddonName,
		Enabled: to.BoolPtr(DefaultAppGwIngressAddonEnabled),
		Config: map[string]string{
			"appgw-subnet":     "",
			"appgw-sku":        "WAF_v2",
			"appgw-private-ip": "",
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
		defaultsCalicoDaemonSetAddonsConfig,
		defaultsAADPodIdentityAddonsConfig,
		defaultAppGwAddonsConfig,
	}
	// Add default addons specification, if no user-provided spec exists
	if o.KubernetesConfig.Addons == nil {
		o.KubernetesConfig.Addons = defaultAddons
	} else {
		for _, addon := range defaultAddons {
			o.KubernetesConfig.Addons = appendAddonIfNotPresent(o.KubernetesConfig.Addons, addon)
		}
	}

	for _, addon := range defaultAddons {
		synthesizeAddonsConfig(o.KubernetesConfig.Addons, addon, isUpdate)
	}

	if len(o.KubernetesConfig.PodSecurityPolicyConfig) > 0 && isUpdate {
		if base64Data, ok := o.KubernetesConfig.PodSecurityPolicyConfig["data"]; ok {
			pspAddonsConfig := KubernetesAddon{
				Name: PodSecurityPolicyAddonName,
				Data: base64Data,
			}
			o.KubernetesConfig.Addons = appendAddonIfNotPresent(o.KubernetesConfig.Addons, pspAddonsConfig)
		}
	}

	// Specific back-compat business logic for calico addon
	// Ensure addon is set to Enabled w/ proper containers config no matter what if NetworkPolicy == calico
	i := getAddonsIndexByName(o.KubernetesConfig.Addons, CalicoAddonName)
	if isUpdate && o.KubernetesConfig.NetworkPolicy == NetworkPolicyCalico && i > -1 && o.KubernetesConfig.Addons[i].Enabled != to.BoolPtr(true) {
		j := getAddonsIndexByName(defaultAddons, CalicoAddonName)
		// Ensure calico is statically set to enabled
		o.KubernetesConfig.Addons[i].Enabled = to.BoolPtr(true)
		// Assume addon configuration was pruned due to an inherited enabled=false, so re-apply default values
		o.KubernetesConfig.Addons[i] = assignDefaultAddonVals(o.KubernetesConfig.Addons[i], defaultAddons[j], isUpdate)
	}

	// Support back-compat configuration for Azure NetworkPolicy, which no longer ships with a "telemetry" container starting w/ 1.16.0
	if isUpdate && o.KubernetesConfig.NetworkPolicy == NetworkPolicyAzure && common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.16.0") {
		i = getAddonsIndexByName(o.KubernetesConfig.Addons, AzureNetworkPolicyAddonName)
		var hasTelemetryContainerConfig bool
		var prunedContainersConfig []KubernetesContainerSpec
		if i > -1 {
			for _, c := range o.KubernetesConfig.Addons[i].Containers {
				if c.Name == AzureVnetTelemetryAddonName {
					hasTelemetryContainerConfig = true
				} else {
					prunedContainersConfig = append(prunedContainersConfig, c)
				}
			}
			if hasTelemetryContainerConfig {
				o.KubernetesConfig.Addons[i].Containers = prunedContainersConfig
			}
		}
	}
}

func appendAddonIfNotPresent(addons []KubernetesAddon, addon KubernetesAddon) []KubernetesAddon {
	i := getAddonsIndexByName(addons, addon.Name)
	if i < 0 {
		return append(addons, addon)
	}
	return addons
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
	if !to.Bool(addon.Enabled) {
		return KubernetesAddon{
			Name:    addon.Name,
			Enabled: addon.Enabled,
		}
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

func synthesizeAddonsConfig(addons []KubernetesAddon, addon KubernetesAddon, isUpdate bool) {
	i := getAddonsIndexByName(addons, addon.Name)
	if i >= 0 {
		addons[i] = assignDefaultAddonVals(addons[i], addon, isUpdate)
	}
}

func getLogAnalyticsWorkspaceDomain(cloudName string) string {
	var workspaceDomain string
	switch cloudName {
	case "AzurePublicCloud":
		workspaceDomain = "opinsights.azure.com"
	case "AzureChinaCloud":
		workspaceDomain = "opinsights.azure.cn"
	case "AzureUSGovernmentCloud":
		workspaceDomain = "opinsights.azure.us"
	case "AzureGermanCloud":
		workspaceDomain = "opinsights.azure.de"
	default:
		workspaceDomain = "opinsights.azure.com"
	}
	return workspaceDomain
}
