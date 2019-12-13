// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
)

const (
	dashboardImageReference                           string = "kubernetes-dashboard-amd64:v1.10.1"
	execHealthZImageReference                         string = "exechealthz-amd64:1.2"
	heapsterImageReference                            string = "heapster-amd64:v1.5.4"
	coreDNSImageReference                             string = "coredns:1.6.5"
	kubeDNSImageReference                             string = "k8s-dns-kube-dns-amd64:1.15.4"
	kubeDNSMasqNannyImageReference                    string = "k8s-dns-dnsmasq-nanny-amd64:1.15.4"
	kubeDNSSidecarImageReference                      string = "k8s-dns-sidecar-amd64:1.14.10"
	pauseImageReference                               string = "pause:1.2.0"
	tillerImageReference                              string = "tiller:v2.13.1"
	reschedulerImageReference                         string = "rescheduler:v0.4.0"
	virtualKubeletImageReference                      string = "virtual-kubelet:latest"
	omsImageReference                                 string = "oms:ciprod11012019"
	azureCNINetworkMonitorImageReference              string = "networkmonitor:v0.0.6"
	nvidiaDevicePluginImageReference                  string = "k8s-device-plugin:1.11"
	blobfuseFlexVolumeImageReference                  string = "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8"
	smbFlexVolumeImageReference                       string = "mcr.microsoft.com/k8s/flexvolume/smb-flexvolume:1.0.2"
	keyvaultFlexVolumeImageReference                  string = "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13"
	ipMasqAgentImageReference                         string = "ip-masq-agent-amd64:v2.5.0"
	dnsAutoscalerImageReference                       string = "cluster-proportional-autoscaler-amd64:1.1.1"
	calicoTyphaImageReference                         string = "typha:v3.8.0"
	calicoCNIImageReference                           string = "cni:v3.8.0"
	calicoNodeImageReference                          string = "node:v3.8.0"
	calicoPod2DaemonImageReference                    string = "pod2daemon-flexvol:v3.8.0"
	calicoClusterProportionalAutoscalerImageReference string = "cluster-proportional-autoscaler-amd64:1.1.2-r2"
	azureNPMContainerImageReference                   string = "mcr.microsoft.com/containernetworking/azure-npm:v1.0.29"
	azureVNETTelemetryContainerImageReference         string = "mcr.microsoft.com/containernetworking/azure-vnet-telemetry:v1.0.29"
	aadPodIdentityNMIImageReference                   string = "mcr.microsoft.com/k8s/aad-pod-identity/nmi:1.2"
	aadPodIdentityMICImageReference                   string = "mcr.microsoft.com/k8s/aad-pod-identity/mic:1.2"
	azurePolicyImageReference                         string = "mcr.microsoft.com/azure-policy/policy-kubernetes-addon-prod:prod_20191011.1"
	gatekeeperImageReference                          string = "quay.io/open-policy-agent/gatekeeper:v3.0.4-beta.2"
	nodeProblemDetectorImageReference                 string = "k8s.gcr.io/node-problem-detector:v0.8.0"
)

// k8sComponentVersions is a convenience map to make UT maintenance easier,
// at the expense of some add'l indirection in getK8sVersionComponents below
var k8sComponentVersions = map[string]map[string]string{
	"1.17": {
		"addon-resizer":                   "addon-resizer:1.8.5",
		"metrics-server":                  "metrics-server-amd64:v0.3.5",
		"addon-manager":                   "kube-addon-manager-amd64:v9.0.2",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.17.0",
	},
	"1.16": {
		"addon-resizer":                   "addon-resizer:1.8.5",
		"metrics-server":                  "metrics-server-amd64:v0.3.4",
		"addon-manager":                   "kube-addon-manager-amd64:v9.0.2",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.16.3",
	},
	"1.15": {
		"addon-resizer":                   "addon-resizer:1.8.5",
		"metrics-server":                  "metrics-server-amd64:v0.2.1",
		"addon-manager":                   "kube-addon-manager-amd64:v9.0.2",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.15.4",
	},
	"1.14": {
		"addon-resizer":                   "addon-resizer:1.8.4",
		"metrics-server":                  "metrics-server-amd64:v0.2.1",
		"addon-manager":                   "kube-addon-manager-amd64:v9.0.2",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.14.7",
	},
	"1.13": {
		"addon-resizer":                   "addon-resizer:1.8.4",
		"metrics-server":                  "metrics-server-amd64:v0.2.1",
		"addon-manager":                   "kube-addon-manager-amd64:v8.9.1",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.13.9",
	},
	"1.12": {
		"addon-resizer":                   "addon-resizer:1.8.4",
		"metrics-server":                  "metrics-server-amd64:v0.2.1",
		"addon-manager":                   "kube-addon-manager-amd64:v8.9.1",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.12.8",
	},
	"1.11": {
		"addon-resizer":                   "addon-resizer:1.8.4",
		"metrics-server":                  "metrics-server-amd64:v0.2.1",
		"addon-manager":                   "kube-addon-manager-amd64:v8.9.1",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.3.9",
	},
	"1.10": {
		"addon-resizer":                    "addon-resizer:1.8.4",
		"metrics-server":                   "metrics-server-amd64:v0.2.1",
		"addon-manager":                    "kube-addon-manager-amd64:v8.9.1",
		"rescheduler":                      "rescheduler:v0.3.1",
		common.ClusterAutoscalerAddonName:  "cluster-autoscaler:v1.2.5",
		common.NVIDIADevicePluginAddonName: "k8s-device-plugin:1.10",
		"k8s-dns-sidecar":                  "k8s-dns-sidecar-amd64:1.14.8",
	},
	"1.9": {
		"addon-resizer":                   "addon-resizer:1.8.4",
		"metrics-server":                  "metrics-server-amd64:v0.2.1",
		"addon-manager":                   "kube-addon-manager-amd64:v8.6",
		"rescheduler":                     "rescheduler:v0.3.1",
		common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.1.2",
		"k8s-dns-sidecar":                 "k8s-dns-sidecar-amd64:1.14.7",
	},
	"1.8": {
		"addon-resizer":  "addon-resizer:1.7",
		"heapster":       "heapster-amd64:v1.5.1",
		"metrics-server": "metrics-server-amd64:v0.2.1",
		"kube-dns":       "k8s-dns-kube-dns-amd64:1.14.13",
		"addon-manager":  "kube-addon-manager-amd64:v8.6",
		"dnsmasq":        "k8s-dns-dnsmasq-nanny-amd64:1.14.8",
		"rescheduler":    "rescheduler:v0.3.1",
	},
	"1.7": {
		"dashboard":      "kubernetes-dashboard-amd64:v1.6.3",
		"addon-resizer":  "addon-resizer:1.7",
		"heapster":       "heapster-amd64:v1.5.1",
		"metrics-server": "metrics-server-amd64:v0.2.1",
		"kube-dns":       "k8s-dns-kube-dns-amd64:1.14.5",
		"addon-manager":  "kube-addon-manager-amd64:v8.6",
		"dnsmasq":        "k8s-dns-dnsmasq-nanny-amd64:1.14.5",
		"rescheduler":    "rescheduler:v0.3.1",
	},
	"1.6": {
		"dashboard":      "kubernetes-dashboard-amd64:v1.6.3",
		"addon-resizer":  "addon-resizer:1.7",
		"heapster":       "heapster-amd64:v1.3.0",
		"metrics-server": "metrics-server-amd64:v0.2.1",
		"kube-dns":       "k8s-dns-kube-dns-amd64:1.14.5",
		"addon-manager":  "kube-addon-manager-amd64:v6.5",
		"dnsmasq":        "k8s-dns-dnsmasq-nanny-amd64:1.14.5",
		"rescheduler":    "rescheduler:v0.3.1",
	},
}

// K8sComponentsByVersionMap represents Docker images used for Kubernetes components based on Kubernetes versions (major.minor.patch)
var K8sComponentsByVersionMap map[string]map[string]string

func init() {
	K8sComponentsByVersionMap = getKubeConfigs()
}

func getKubeConfigs() map[string]map[string]string {
	ret := make(map[string]map[string]string)
	for _, version := range common.GetAllSupportedKubernetesVersions(true, false) {
		ret[version] = getK8sVersionComponents(version, getVersionOverrides(version))
	}
	return ret
}

func getVersionOverrides(v string) map[string]string {
	switch v {
	case "1.8.11":
		return map[string]string{"kube-dns": "k8s-dns-kube-dns-amd64:1.14.9"}
	case "1.8.9":
		return map[string]string{"windowszip": "v1.8.9-2int.zip"}
	case "1.8.6":
		return map[string]string{"windowszip": "v1.8.6-2int.zip"}
	case "1.8.2":
		return map[string]string{"windowszip": "v1.8.2-2int.zip"}
	case "1.8.1":
		return map[string]string{"windowszip": "v1.8.1-2int.zip"}
	case "1.8.0":
		return map[string]string{"windowszip": "v1.8.0-2int.zip"}
	case "1.7.16":
		return map[string]string{"windowszip": "v1.7.16-1int.zip"}
	case "1.7.15":
		return map[string]string{"windowszip": "v1.7.15-1int.zip"}
	case "1.7.14":
		return map[string]string{"windowszip": "v1.7.14-1int.zip"}
	case "1.7.13":
		return map[string]string{"windowszip": "v1.7.13-1int.zip"}
	case "1.7.12":
		return map[string]string{"windowszip": "v1.7.12-2int.zip"}
	case "1.7.10":
		return map[string]string{"windowszip": "v1.7.10-1int.zip"}
	case "1.7.9":
		return map[string]string{"windowszip": "v1.7.9-2int.zip"}
	case "1.7.7":
		return map[string]string{"windowszip": "v1.7.7-2int.zip"}
	case "1.7.5":
		return map[string]string{"windowszip": "v1.7.5-4int.zip"}
	case "1.7.4":
		return map[string]string{"windowszip": "v1.7.4-2int.zip"}
	case "1.7.2":
		return map[string]string{"windowszip": "v1.7.2-1int.zip"}
	default:
		return nil
	}
}

func getK8sVersionComponents(version string, overrides map[string]string) map[string]string {
	s := strings.Split(version, ".")
	majorMinor := strings.Join(s[:2], ".")
	var ret map[string]string
	k8sComponent := k8sComponentVersions[majorMinor]
	switch majorMinor {
	case "1.17":
		ret = map[string]string{
			"kube-apiserver":                         "kube-apiserver:v" + version,
			"kube-controller-manager":                "kube-controller-manager:v" + version,
			common.KubeProxyAddonName:                "kube-proxy:v" + version,
			"kube-scheduler":                         "kube-scheduler:v" + version,
			"ccm":                                    "azure-cloud-controller-manager:v0.3.0",
			common.CloudNodeManagerAddonName:         "azure-cloud-node-manager:v0.3.0",
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			common.CoreDNSAddonName:                  coreDNSImageReference,
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              reschedulerImageReference,
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        kubeDNSSidecarImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       nvidiaDevicePluginImageReference,
		}
	case "1.16":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "azure-cloud-controller-manager:v0.3.0",
			common.CloudNodeManagerAddonName:         "azure-cloud-node-manager:v0.3.0",
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			common.CoreDNSAddonName:                  coreDNSImageReference,
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              reschedulerImageReference,
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        kubeDNSSidecarImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       nvidiaDevicePluginImageReference,
		}
	case "1.15":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			common.CoreDNSAddonName:                  coreDNSImageReference,
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              reschedulerImageReference,
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        kubeDNSSidecarImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       nvidiaDevicePluginImageReference,
		}
	case "1.14":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			common.CoreDNSAddonName:                  coreDNSImageReference,
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              reschedulerImageReference,
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        kubeDNSSidecarImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       nvidiaDevicePluginImageReference,
		}
	case "1.13":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			common.CoreDNSAddonName:                  coreDNSImageReference,
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              reschedulerImageReference,
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        kubeDNSSidecarImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       nvidiaDevicePluginImageReference,
		}
	case "1.12":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			common.CoreDNSAddonName:                  coreDNSImageReference,
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              reschedulerImageReference,
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        kubeDNSSidecarImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       nvidiaDevicePluginImageReference,
		}
	case "1.11":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              reschedulerImageReference,
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        kubeDNSSidecarImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       nvidiaDevicePluginImageReference,
		}
	case "1.10":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              k8sComponent["rescheduler"],
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        k8sComponent["k8s-dns-sidecar"],
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:       k8sComponent[common.NVIDIADevicePluginAddonName],
		}
	case "1.9":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               heapsterImageReference,
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			"kube-dns":                               kubeDNSImageReference,
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                kubeDNSMasqNannyImageReference,
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              k8sComponent["rescheduler"],
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:        k8sComponent[common.ClusterAutoscalerAddonName],
			"k8s-dns-sidecar":                        k8sComponent["k8s-dns-sidecar"],
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:      nodeProblemDetectorImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
		}
	case "1.8":
		ret = map[string]string{
			"hyperkube":                              "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:                "hyperkube-amd64:v" + version,
			"ccm":                                    "cloud-controller-manager-amd64:v" + version,
			"windowszip":                             "v" + version + "-1int.zip",
			common.DashboardAddonName:                dashboardImageReference,
			"exechealthz":                            execHealthZImageReference,
			"addonresizer":                           k8sComponent["addon-resizer"],
			"heapster":                               k8sComponent["heapster"],
			common.MetricsServerAddonName:            k8sComponent["metrics-server"],
			"kube-dns":                               k8sComponent["kube-dns"],
			"addonmanager":                           k8sComponent["addon-manager"],
			"dnsmasq":                                k8sComponent["dnsmasq"],
			"pause":                                  pauseImageReference,
			common.TillerAddonName:                   tillerImageReference,
			common.ReschedulerAddonName:              k8sComponent["rescheduler"],
			common.ACIConnectorAddonName:             virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:      omsImageReference,
			common.AzureCNINetworkMonitorAddonName:   azureCNINetworkMonitorImageReference,
			common.BlobfuseFlexVolumeAddonName:       blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:            smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:       keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:              ipMasqAgentImageReference,
			common.DNSAutoscalerAddonName:            dnsAutoscalerImageReference,
			common.AzureNetworkPolicyAddonName:       azureNPMContainerImageReference,
			common.AzureVnetTelemetryContainerName:   azureVNETTelemetryContainerImageReference,
			"calico-typha":                           calicoTyphaImageReference,
			"calico-cni":                             calicoCNIImageReference,
			"calico-node":                            calicoNodeImageReference,
			"calico-pod2daemon":                      calicoPod2DaemonImageReference,
			"calico-cluster-proportional-autoscaler": calicoClusterProportionalAutoscalerImageReference,
			common.NMIContainerName:                  aadPodIdentityNMIImageReference,
			common.MICContainerName:                  aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:              azurePolicyImageReference,
			common.GatekeeperContainerName:           gatekeeperImageReference,
			"nodestatusfreq":                         DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                        DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                            DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                            DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                         strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                          strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                        strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                           strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                      strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                        strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                   strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                        strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                         strconv.Itoa(DefaultKubernetesGCLowThreshold),
		}
	case "1.7":
		ret = map[string]string{
			"hyperkube":                            "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:              "hyperkube-amd64:v" + version,
			common.DashboardAddonName:              k8sComponent["dashboard"],
			"exechealthz":                          execHealthZImageReference,
			"addonresizer":                         k8sComponent["addon-resizer"],
			"heapster":                             k8sComponent["heapster"],
			common.MetricsServerAddonName:          k8sComponent["metrics-server"],
			"kube-dns":                             k8sComponent["kube-dns"],
			"addonmanager":                         k8sComponent["addon-manager"],
			"dnsmasq":                              k8sComponent["dnsmasq"],
			"pause":                                pauseImageReference,
			common.TillerAddonName:                 tillerImageReference,
			common.ReschedulerAddonName:            k8sComponent["rescheduler"],
			common.ACIConnectorAddonName:           virtualKubeletImageReference,
			common.ContainerMonitoringAddonName:    omsImageReference,
			common.AzureCNINetworkMonitorAddonName: azureCNINetworkMonitorImageReference,
			"nodestatusfreq":                       DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                      DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                          DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                          DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                       strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                      strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                      strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                         strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                      strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                      strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                       strconv.Itoa(DefaultKubernetesGCLowThreshold),
		}
	case "1.6":
		ret = map[string]string{
			"hyperkube":                            "hyperkube-amd64:v" + version,
			common.KubeProxyAddonName:              "hyperkube-amd64:v" + version,
			common.DashboardAddonName:              k8sComponent["dashboard"],
			"exechealthz":                          execHealthZImageReference,
			"addonresizer":                         k8sComponent["addon-resizer"],
			"heapster":                             k8sComponent["heapster"],
			common.MetricsServerAddonName:          k8sComponent["metrics-server"],
			"kube-dns":                             k8sComponent["kube-dns"],
			"addonmanager":                         k8sComponent["addon-manager"],
			"dnsmasq":                              k8sComponent["dnsmasq"],
			"pause":                                pauseImageReference,
			common.TillerAddonName:                 tillerImageReference,
			common.ReschedulerAddonName:            k8sComponent["rescheduler"],
			common.ACIConnectorAddonName:           virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName: azureCNINetworkMonitorImageReference,
			"nodestatusfreq":                       DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                      DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                          DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                          DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                       strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                        strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                      strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                      strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                         strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                      strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                      strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                       strconv.Itoa(DefaultKubernetesGCLowThreshold),
		}

	default:
		ret = nil
	}
	for k, v := range overrides {
		ret[k] = v
	}
	return ret
}
