// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
)

const (
	pauseImageReference                               string = "oss/kubernetes/pause:1.4.0"
	blobfuseFlexVolumeImageReference                  string = "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8"
	smbFlexVolumeImageReference                       string = "mcr.microsoft.com/k8s/flexvolume/smb-flexvolume:1.0.2"
	keyvaultFlexVolumeImageReference                  string = "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.16"
	calicoTyphaImageReference                         string = "typha:v3.8.0"
	calicoCNIImageReference                           string = "cni:v3.8.0"
	calicoNodeImageReference                          string = "node:v3.8.0"
	calicoPod2DaemonImageReference                    string = "pod2daemon-flexvol:v3.8.0"
	calicoClusterProportionalAutoscalerImageReference string = "mcr.microsoft.com/oss/kubernetes/autoscaler/cluster-proportional-autoscaler:1.1.2-r2"
	ciliumAgentImageReference                         string = "docker.io/cilium/cilium:v1.4"
	ciliumCleanStateImageReference                    string = "docker.io/cilium/cilium-init:2018-10-16"
	ciliumOperatorImageReference                      string = "docker.io/cilium/operator:v1.4"
	ciliumEtcdOperatorImageReference                  string = "docker.io/cilium/cilium-etcd-operator:v2.0.5"
	antreaControllerImageReference                    string = "antrea/antrea-ubuntu:v0.6.0"
	antreaAgentImageReference                                = antreaControllerImageReference
	antreaOVSImageReference                                  = antreaControllerImageReference
	antreaInstallCNIImageReference                           = antreaControllerImageReference
	azureNPMContainerImageReference                   string = "mcr.microsoft.com/containernetworking/azure-npm:v1.1.7"
	aadPodIdentityNMIImageReference                   string = "mcr.microsoft.com/k8s/aad-pod-identity/nmi:1.6.1"
	aadPodIdentityMICImageReference                   string = "mcr.microsoft.com/k8s/aad-pod-identity/mic:1.6.1"
	azurePolicyImageReference                         string = "mcr.microsoft.com/azure-policy/policy-kubernetes-addon-prod:prod_20200505.1"
	gatekeeperImageReference                          string = "mcr.microsoft.com/oss/open-policy-agent/gatekeeper:v3.1.0-beta.8"
	nodeProblemDetectorImageReference                 string = "k8s.gcr.io/node-problem-detector/node-problem-detector:v0.8.4"
	csiProvisionerImageReference                      string = "oss/kubernetes-csi/csi-provisioner:v1.5.0"
	csiAttacherImageReference                         string = "oss/kubernetes-csi/csi-attacher:v1.2.0"
	csiLivenessProbeImageReference                    string = "oss/kubernetes-csi/livenessprobe:v2.0.0"
	csiLivenessProbeWindowsImageReference             string = "oss/kubernetes-csi/livenessprobe:v2.0.1-alpha.1-windows-1809-amd64"
	csiNodeDriverRegistrarImageReference              string = "oss/kubernetes-csi/csi-node-driver-registrar:v1.2.0"
	csiNodeDriverRegistrarWindowsImageReference       string = "oss/kubernetes-csi/csi-node-driver-registrar:v1.2.1-alpha.1-windows-1809-amd64"
	csiResizerImageReference                          string = "oss/kubernetes-csi/csi-resizer:v0.3.0"
	csiSnapshotterImageReference                      string = "oss/kubernetes-csi/csi-snapshotter:v2.0.0"
	csiSnapshotControllerImageReference               string = "oss/kubernetes-csi/snapshot-controller:v2.0.0"
	csiAzureDiskImageReference                        string = "k8s/csi/azuredisk-csi:v0.7.0"
	csiAzureFileImageReference                        string = "k8s/csi/azurefile-csi:v0.6.0"
	azureCloudControllerManagerImageReference         string = "oss/kubernetes/azure-cloud-controller-manager:v0.5.1"
	azureCloudNodeManagerImageReference               string = "oss/kubernetes/azure-cloud-node-manager:v0.5.1"
	dashboardImageReference                           string = "mcr.microsoft.com/oss/kubernetes/dashboard:v2.0.4"
	dashboardMetricsScraperImageReference             string = "mcr.microsoft.com/oss/kubernetes/metrics-scraper:v1.0.4"
	kubeFlannelImageReference                         string = "quay.io/coreos/flannel:v0.8.0-amd64"
	flannelInstallCNIImageReference                   string = "quay.io/coreos/flannel:v0.10.0-amd64"
	KubeRBACProxyImageReference                       string = "gcr.io/kubebuilder/kube-rbac-proxy:v0.4.0"
	ScheduledMaintenanceManagerImageReference         string = "quay.io/awesomenix/drainsafe-manager:latest"
	nvidiaDevicePluginImageReference                  string = "oss/nvidia/k8s-device-plugin:1.0.0-beta6"
	virtualKubeletImageReference                      string = "virtual-kubelet:1.2.1.2"
	azureCNINetworkMonitorImageReference              string = "networkmonitor:v0.0.8"
	tillerImageReference                              string = "oss/kubernetes/tiller:v2.13.1"
	csiSecretsStoreProviderAzureImageReference        string = "oss/azure/secrets-store/provider-azure:0.0.9"
	csiSecretsStoreDriverImageReference               string = "k8s/csi/secrets-store/driver:v0.0.14"
	clusterProportionalAutoscalerImageReference       string = "mcr.microsoft.com/oss/kubernetes/autoscaler/cluster-proportional-autoscaler:1.7.1"
	azureArcOnboardingImageReference                  string = "arck8sonboarding.azurecr.io/arck8sonboarding:v0.1.0"
	azureKMSProviderImageReference                    string = "k8s/kms/keyvault:v0.0.9"
)

var kubernetesImageBaseDefaultImages = map[string]map[string]string{
	common.KubernetesImageBaseTypeGCR: {
		common.DashboardAddonName:                   "kubernetes-dashboard-amd64:v1.10.1",
		common.DashboardMetricsScraperContainerName: "",
		common.ExecHealthZComponentName:             "exechealthz-amd64:1.2",
		common.CoreDNSAddonName:                     "coredns:1.6.7",
		common.KubeDNSAddonName:                     "k8s-dns-kube-dns-amd64:1.15.4",
		common.DNSMasqComponentName:                 "k8s-dns-dnsmasq-nanny-amd64:1.15.4",
		common.DNSSidecarComponentName:              "k8s-dns-sidecar-amd64:1.14.10",
		common.ReschedulerAddonName:                 "rescheduler:v0.4.0",
		common.IPMASQAgentAddonName:                 "ip-masq-agent-amd64:v2.5.0",
		common.KubeProxyAddonName:                   "kube-proxy",
		common.ControllerManagerComponentName:       "kube-controller-manager",
		common.APIServerComponentName:               "kube-apiserver",
		common.SchedulerComponentName:               "kube-scheduler",
		common.Hyperkube:                            "hyperkube-amd64",
	},
	common.KubernetesImageBaseTypeMCR: {
		common.DashboardAddonName:                   "oss/kubernetes/dashboard:v2.0.4",
		common.DashboardMetricsScraperContainerName: "oss/kubernetes/metrics-scraper:v1.0.4",
		common.ExecHealthZComponentName:             "oss/kubernetes/exechealthz:1.2",
		common.CoreDNSAddonName:                     "oss/kubernetes/coredns:1.7.0",
		common.KubeDNSAddonName:                     "oss/kubernetes/k8s-dns-kube-dns:1.15.4",
		common.DNSMasqComponentName:                 "oss/kubernetes/k8s-dns-dnsmasq-nanny:1.15.4",
		common.DNSSidecarComponentName:              "oss/kubernetes/k8s-dns-sidecar:1.14.10",
		common.ReschedulerAddonName:                 "oss/kubernetes/rescheduler:v0.4.0",
		common.IPMASQAgentAddonName:                 "oss/kubernetes/ip-masq-agent:v2.5.0",
		common.KubeProxyAddonName:                   "oss/kubernetes/kube-proxy",
		common.ControllerManagerComponentName:       "oss/kubernetes/kube-controller-manager",
		common.APIServerComponentName:               "oss/kubernetes/kube-apiserver",
		common.SchedulerComponentName:               "oss/kubernetes/kube-scheduler",
		common.Hyperkube:                            "oss/kubernetes/hyperkube",
	},
}

var csiSidecarComponentsOverrides = map[string]map[string]string{
	common.AzureFileCSIDriverAddonName: {
		common.CSIProvisionerContainerName: "oss/kubernetes-csi/csi-provisioner:v1.4.0",
		common.CSISnapshotterContainerName: "oss/kubernetes-csi/csi-snapshotter:v1.1.0",
	},
}

func getDefaultImage(image, kubernetesImageBaseType string) string {
	return kubernetesImageBaseDefaultImages[kubernetesImageBaseType][image]
}

// kubernetesImageBaseVersionedImages is a convenience map for "kubernetesImageBase" image version references that are distinct across versions of Kubernetes
// For example, cluster-autoscaler generally ships a per-Kubernetes-version build
// The map supports GCR or MCR image string flavors
var kubernetesImageBaseVersionedImages = map[string]map[string]map[string]string{
	common.KubernetesImageBaseTypeGCR: {
		"1.19": {
			common.AddonResizerComponentName:  "addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "metrics-server/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "kube-addon-manager-amd64:v9.1.1",
			common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.18.0",
		},
		"1.18": {
			common.AddonResizerComponentName:  "addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "metrics-server/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "kube-addon-manager-amd64:v9.1.1",
			common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.18.0",
		},
		"1.17": {
			common.AddonResizerComponentName:  "addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "metrics-server/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "kube-addon-manager-amd64:v9.1.1",
			common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.17.1",
		},
		"1.16": {
			common.AddonResizerComponentName:  "addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "metrics-server/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "kube-addon-manager-amd64:v9.1.1",
			common.ClusterAutoscalerAddonName: "cluster-autoscaler:v1.16.4",
		},
		"1.15": {
			common.AddonResizerComponentName:           "addon-resizer:1.8.7",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v9.1.1",
			common.ClusterAutoscalerAddonName:          "cluster-autoscaler:v1.15.5",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.14": {
			common.AddonResizerComponentName:           "addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v9.0.2",
			common.ClusterAutoscalerAddonName:          "cluster-autoscaler:v1.14.7",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.13": {
			common.AddonResizerComponentName:           "addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v8.9.1",
			common.ClusterAutoscalerAddonName:          "cluster-autoscaler:v1.13.9",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.12": {
			common.AddonResizerComponentName:           "addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v8.9.1",
			common.ClusterAutoscalerAddonName:          "cluster-autoscaler:v1.12.8",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.11": {
			common.AddonResizerComponentName:           "addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v8.9.1",
			common.ClusterAutoscalerAddonName:          "cluster-autoscaler:v1.3.9",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.10": {
			common.AddonResizerComponentName:           "addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v8.9.1",
			common.ReschedulerAddonName:                "rescheduler:v0.3.1",
			common.ClusterAutoscalerAddonName:          "cluster-autoscaler:v1.2.5",
			common.DNSSidecarComponentName:             "k8s-dns-sidecar-amd64:1.14.8",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.9": {
			common.AddonResizerComponentName:           "addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v8.6",
			common.ReschedulerAddonName:                "rescheduler:v0.3.1",
			common.ClusterAutoscalerAddonName:          "cluster-autoscaler:v1.1.2",
			common.DNSSidecarComponentName:             "k8s-dns-sidecar-amd64:1.14.7",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.8": {
			common.AddonResizerComponentName:           "addon-resizer:1.7",
			common.MetricsServerAddonName:              "metrics-server-amd64:v0.2.1",
			common.KubeDNSAddonName:                    "k8s-dns-kube-dns-amd64:1.14.13",
			common.AddonManagerComponentName:           "kube-addon-manager-amd64:v8.6",
			common.DNSMasqComponentName:                "k8s-dns-dnsmasq-nanny-amd64:1.14.8",
			common.ReschedulerAddonName:                "rescheduler:v0.3.1",
			common.CloudControllerManagerComponentName: "cloud-controller-manager-amd64",
		},
		"1.7": {
			"dashboard":                      "kubernetes-dashboard-amd64:v1.6.3",
			common.AddonResizerComponentName: "addon-resizer:1.7",
			common.MetricsServerAddonName:    "metrics-server-amd64:v0.2.1",
			common.KubeDNSAddonName:          "k8s-dns-kube-dns-amd64:1.14.5",
			common.AddonManagerComponentName: "kube-addon-manager-amd64:v8.6",
			common.DNSMasqComponentName:      "k8s-dns-dnsmasq-nanny-amd64:1.14.5",
			common.ReschedulerAddonName:      "rescheduler:v0.3.1",
		},
		"1.6": {
			"dashboard":                      "kubernetes-dashboard-amd64:v1.6.3",
			common.AddonResizerComponentName: "addon-resizer:1.7",
			common.MetricsServerAddonName:    "metrics-server-amd64:v0.2.1",
			common.KubeDNSAddonName:          "k8s-dns-kube-dns-amd64:1.14.5",
			common.AddonManagerComponentName: "kube-addon-manager-amd64:v6.5",
			common.DNSMasqComponentName:      "k8s-dns-dnsmasq-nanny-amd64:1.14.5",
			common.ReschedulerAddonName:      "rescheduler:v0.3.1",
		},
	},
	common.KubernetesImageBaseTypeMCR: {
		"1.19": {
			common.AddonResizerComponentName:  "oss/kubernetes/autoscaler/addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "oss/kubernetes/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "oss/kubernetes/kube-addon-manager:v9.1.1",
			common.ClusterAutoscalerAddonName: "oss/kubernetes/autoscaler/cluster-autoscaler:v1.19.0",
		},
		"1.18": {
			common.AddonResizerComponentName:  "oss/kubernetes/autoscaler/addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "oss/kubernetes/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "oss/kubernetes/kube-addon-manager:v9.1.1",
			common.ClusterAutoscalerAddonName: "oss/kubernetes/autoscaler/cluster-autoscaler:v1.18.2",
		},
		"1.17": {
			common.AddonResizerComponentName:  "oss/kubernetes/autoscaler/addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "oss/kubernetes/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "oss/kubernetes/kube-addon-manager:v9.1.1",
			common.ClusterAutoscalerAddonName: "oss/kubernetes/autoscaler/cluster-autoscaler:v1.17.3",
		},
		"1.16": {
			common.AddonResizerComponentName:  "oss/kubernetes/autoscaler/addon-resizer:1.8.7",
			common.MetricsServerAddonName:     "oss/kubernetes/metrics-server:v0.3.7",
			common.AddonManagerComponentName:  "oss/kubernetes/kube-addon-manager:v9.1.1",
			common.ClusterAutoscalerAddonName: "oss/kubernetes/autoscaler/cluster-autoscaler:v1.16.6",
		},
		"1.15": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.8.7",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v9.1.1",
			common.ClusterAutoscalerAddonName:          "oss/kubernetes/autoscaler/cluster-autoscaler:v1.15.7",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.14": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v9.0.2",
			common.ClusterAutoscalerAddonName:          "oss/kubernetes/autoscaler/cluster-autoscaler:v1.14.8",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.13": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v8.9.1",
			common.ClusterAutoscalerAddonName:          "oss/kubernetes/autoscaler/cluster-autoscaler:v1.13.9",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.12": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v8.9.1",
			common.ClusterAutoscalerAddonName:          "oss/kubernetes/autoscaler/cluster-autoscaler:v1.12.8",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.11": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v8.9.1",
			common.ClusterAutoscalerAddonName:          "oss/kubernetes/autoscaler/cluster-autoscaler:v1.3.9",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.10": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v8.9.1",
			common.ReschedulerAddonName:                "oss/kubernetes/rescheduler:v0.3.1",
			common.ClusterAutoscalerAddonName:          "oss/kubernetes/autoscaler/cluster-autoscaler:v1.2.5",
			common.DNSSidecarComponentName:             "oss/kubernetes/k8s-dns-sidecar:1.14.8",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.9": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.8.4",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v8.6",
			common.ReschedulerAddonName:                "oss/kubernetes/rescheduler:v0.3.1",
			common.ClusterAutoscalerAddonName:          "oss/kubernetes/autoscaler/cluster-autoscaler:v1.1.2",
			common.DNSSidecarComponentName:             "oss/kubernetes/k8s-dns-sidecar:1.14.7",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.8": {
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.7",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.KubeDNSAddonName:                    "oss/kubernetes/k8s-dns-kube-dns:1.14.13",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v8.6",
			common.DNSMasqComponentName:                "oss/kubernetes/k8s-dns-dnsmasq-nanny:1.14.8",
			common.ReschedulerAddonName:                "oss/kubernetes/rescheduler:v0.3.1",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.7": {
			"dashboard":                                "oss/kubernetes/kubernetes-dashboard:v1.6.3",
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.7",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.KubeDNSAddonName:                    "oss/kubernetes/k8s-dns-kube-dns:1.14.5",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v8.6",
			common.DNSMasqComponentName:                "oss/kubernetes/k8s-dns-dnsmasq-nanny:1.14.5",
			common.ReschedulerAddonName:                "oss/kubernetes/rescheduler:v0.3.1",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
		"1.6": {
			"dashboard":                                "oss/kubernetes/kubernetes-dashboard:v1.6.3",
			common.AddonResizerComponentName:           "oss/kubernetes/autoscaler/addon-resizer:1.7",
			common.MetricsServerAddonName:              "oss/kubernetes/metrics-server:v0.2.1",
			common.KubeDNSAddonName:                    "oss/kubernetes/k8s-dns-kube-dns:1.14.5",
			common.AddonManagerComponentName:           "oss/kubernetes/kube-addon-manager:v6.5",
			common.DNSMasqComponentName:                "oss/kubernetes/k8s-dns-dnsmasq-nanny:1.14.5",
			common.ReschedulerAddonName:                "oss/kubernetes/rescheduler:v0.3.1",
			common.CloudControllerManagerComponentName: "oss/kubernetes/cloud-controller-manager",
		},
	},
}

type getK8sVersionComponentsOverrides func(string) map[string]string

func GetK8sComponentsByVersionMap(k *KubernetesConfig) map[string]map[string]string {
	var overrides getK8sVersionComponentsOverrides
	switch k.KubernetesImageBaseType {
	case common.KubernetesImageBaseTypeGCR:
		overrides = getVersionOverridesGCR
	case common.KubernetesImageBaseTypeMCR:
		overrides = getVersionOverridesMCR
	default:
		overrides = getVersionOverridesGCR
	}
	ret := make(map[string]map[string]string)
	for _, version := range common.GetAllSupportedKubernetesVersions(true, false, false) {
		ret[version] = getK8sVersionComponents(version, k.KubernetesImageBaseType, overrides(version))
	}
	return ret
}

func getVersionOverridesMCR(v string) map[string]string {
	switch v {
	case "1.18.6":
		return map[string]string{common.WindowsArtifactComponentName: "v1.18.6-hotfix.20200723/windowszip/v1.18.6-hotfix.20200723-1int.zip"}
	case "1.18.4":
		return map[string]string{common.WindowsArtifactComponentName: "v1.18.4-hotfix.20200626/windowszip/v1.18.4-hotfix.20200626-1int.zip"}
	case "1.18.2":
		return map[string]string{common.WindowsArtifactComponentName: "v1.18.2-hotfix.20200624/windowszip/v1.18.2-hotfix.20200624-1int.zip"}
	case "1.17.9":
		return map[string]string{common.WindowsArtifactComponentName: "v1.17.9-hotfix.20200817/windowszip/v1.17.9-hotfix.20200817-1int.zip"}
	case "1.17.7":
		return map[string]string{common.WindowsArtifactComponentName: "v1.17.7-hotfix.20200817/windowszip/v1.17.7-hotfix.20200817-1int.zip"}
	case "1.16.13":
		return map[string]string{common.WindowsArtifactComponentName: "v1.16.13-hotfix.20200817/windowszip/v1.16.13-hotfix.20200817-1int.zip"}
	case "1.16.11":
		return map[string]string{common.WindowsArtifactComponentName: "v1.16.11-hotfix.20200617/windowszip/v1.16.11-hotfix.20200617-1int.zip"}
	case "1.16.10":
		return map[string]string{common.WindowsArtifactComponentName: "v1.16.10-hotfix.20200817/windowszip/v1.16.10-hotfix.20200817-1int.zip"}
	case "1.15.12":
		return map[string]string{common.WindowsArtifactComponentName: "v1.15.12-hotfix.20200817/windowszip/v1.15.12-hotfix.20200817-1int.zip"}
	case "1.15.11":
		return map[string]string{common.WindowsArtifactComponentName: "v1.15.11-hotfix.20200817/windowszip/v1.15.11-hotfix.20200817-1int.zip"}
	default:
		return nil
	}
}

func getVersionOverridesGCR(v string) map[string]string {
	switch v {
	case "1.18.6":
		return map[string]string{common.WindowsArtifactComponentName: "v1.18.6-hotfix.20200723/windowszip/v1.18.6-hotfix.20200723-1int.zip"}
	case "1.18.4":
		return map[string]string{common.WindowsArtifactComponentName: "v1.18.4-hotfix.20200626/windowszip/v1.18.4-hotfix.20200626-1int.zip"}
	case "1.18.2":
		return map[string]string{common.WindowsArtifactComponentName: "v1.18.2-hotfix.20200624/windowszip/v1.18.2-hotfix.20200624-1int.zip"}
	case "1.17.9":
		return map[string]string{common.WindowsArtifactComponentName: "v1.17.9-hotfix.20200817/windowszip/v1.17.9-hotfix.20200817-1int.zip"}
	case "1.17.7":
		return map[string]string{common.WindowsArtifactComponentName: "v1.17.7-hotfix.20200817/windowszip/v1.17.7-hotfix.20200817-1int.zip"}
	case "1.16.13":
		return map[string]string{common.WindowsArtifactComponentName: "v1.16.13-hotfix.20200817/windowszip/v1.16.13-hotfix.20200817-1int.zip"}
	case "1.16.11":
		return map[string]string{common.WindowsArtifactComponentName: "v1.16.11-hotfix.20200617/windowszip/v1.16.11-hotfix.20200617-1int.zip"}
	case "1.16.10":
		return map[string]string{common.WindowsArtifactComponentName: "v1.16.10-hotfix.20200817/windowszip/v1.16.10-hotfix.20200817-1int.zip"}
	case "1.15.12":
		return map[string]string{common.WindowsArtifactComponentName: "v1.15.12-hotfix.20200817/windowszip/v1.15.12-hotfix.20200817-1int.zip"}
	case "1.15.11":
		return map[string]string{common.WindowsArtifactComponentName: "v1.15.11-hotfix.20200817/windowszip/v1.15.11-hotfix.20200817-1int.zip"}
	case "1.8.11":
		return map[string]string{common.KubeDNSAddonName: "k8s-dns-kube-dns-amd64:1.14.9"}
	case "1.8.9":
		return map[string]string{common.WindowsArtifactComponentName: "v1.8.9-2int.zip"}
	case "1.8.6":
		return map[string]string{common.WindowsArtifactComponentName: "v1.8.6-2int.zip"}
	case "1.8.2":
		return map[string]string{common.WindowsArtifactComponentName: "v1.8.2-2int.zip"}
	case "1.8.1":
		return map[string]string{common.WindowsArtifactComponentName: "v1.8.1-2int.zip"}
	case "1.8.0":
		return map[string]string{common.WindowsArtifactComponentName: "v1.8.0-2int.zip"}
	case "1.7.16":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.16-1int.zip"}
	case "1.7.15":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.15-1int.zip"}
	case "1.7.14":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.14-1int.zip"}
	case "1.7.13":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.13-1int.zip"}
	case "1.7.12":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.12-2int.zip"}
	case "1.7.10":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.10-1int.zip"}
	case "1.7.9":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.9-2int.zip"}
	case "1.7.7":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.7-2int.zip"}
	case "1.7.5":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.5-4int.zip"}
	case "1.7.4":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.4-2int.zip"}
	case "1.7.2":
		return map[string]string{common.WindowsArtifactComponentName: "v1.7.2-1int.zip"}
	default:
		return nil
	}
}

func getK8sVersionComponents(version, kubernetesImageBaseType string, overrides map[string]string) map[string]string {
	s := strings.Split(version, ".")
	majorMinor := strings.Join(s[:2], ".")
	var ret map[string]string
	k8sComponent := kubernetesImageBaseVersionedImages[kubernetesImageBaseType][majorMinor]
	switch majorMinor {
	case "1.19":
		ret = map[string]string{
			common.APIServerComponentName:                     getDefaultImage(common.APIServerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.ControllerManagerComponentName:             getDefaultImage(common.ControllerManagerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.KubeProxyAddonName, kubernetesImageBaseType) + ":v" + version,
			common.SchedulerComponentName:                     getDefaultImage(common.SchedulerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        azureCloudControllerManagerImageReference,
			common.CloudNodeManagerAddonName:                  azureCloudNodeManagerImageReference,
			common.WindowsArtifactComponentName:               "v" + version + "/windowszip/v" + version + "-1int.zip",
			common.WindowsArtifactAzureStackComponentName:     "v" + version + common.AzureStackSuffix + "/windowszip/v" + version + common.AzureStackSuffix + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.DashboardMetricsScraperContainerName:       dashboardMetricsScraperImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.CSIProvisionerContainerName:                csiProvisionerImageReference,
			common.CSIAttacherContainerName:                   csiAttacherImageReference,
			common.CSILivenessProbeContainerName:              csiLivenessProbeImageReference,
			common.CSILivenessProbeWindowsContainerName:       csiLivenessProbeWindowsImageReference,
			common.CSINodeDriverRegistrarContainerName:        csiNodeDriverRegistrarImageReference,
			common.CSINodeDriverRegistrarWindowsContainerName: csiNodeDriverRegistrarWindowsImageReference,
			common.CSISnapshotterContainerName:                csiSnapshotterImageReference,
			common.CSISnapshotControllerContainerName:         csiSnapshotControllerImageReference,
			common.CSIResizerContainerName:                    csiResizerImageReference,
			common.CSIAzureDiskContainerName:                  csiAzureDiskImageReference,
			common.CSIAzureFileContainerName:                  csiAzureFileImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
			common.CSISecretsStoreProviderAzureContainerName:  csiSecretsStoreProviderAzureImageReference,
			common.CSISecretsStoreDriverContainerName:         csiSecretsStoreDriverImageReference,
			common.AzureArcOnboardingAddonName:                azureArcOnboardingImageReference,
			common.AzureKMSProviderComponentName:              azureKMSProviderImageReference,
		}
	case "1.18":
		ret = map[string]string{
			common.APIServerComponentName:                     getDefaultImage(common.APIServerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.ControllerManagerComponentName:             getDefaultImage(common.ControllerManagerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.KubeProxyAddonName, kubernetesImageBaseType) + ":v" + version,
			common.SchedulerComponentName:                     getDefaultImage(common.SchedulerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        azureCloudControllerManagerImageReference,
			common.CloudNodeManagerAddonName:                  azureCloudNodeManagerImageReference,
			common.WindowsArtifactComponentName:               "v" + version + "/windowszip/v" + version + "-1int.zip",
			common.WindowsArtifactAzureStackComponentName:     "v" + version + common.AzureStackSuffix + "/windowszip/v" + version + common.AzureStackSuffix + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.DashboardMetricsScraperContainerName:       dashboardMetricsScraperImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.CSIProvisionerContainerName:                csiProvisionerImageReference,
			common.CSIAttacherContainerName:                   csiAttacherImageReference,
			common.CSILivenessProbeContainerName:              csiLivenessProbeImageReference,
			common.CSILivenessProbeWindowsContainerName:       csiLivenessProbeWindowsImageReference,
			common.CSINodeDriverRegistrarContainerName:        csiNodeDriverRegistrarImageReference,
			common.CSINodeDriverRegistrarWindowsContainerName: csiNodeDriverRegistrarWindowsImageReference,
			common.CSISnapshotterContainerName:                csiSnapshotterImageReference,
			common.CSISnapshotControllerContainerName:         csiSnapshotControllerImageReference,
			common.CSIResizerContainerName:                    csiResizerImageReference,
			common.CSIAzureDiskContainerName:                  csiAzureDiskImageReference,
			common.CSIAzureFileContainerName:                  csiAzureFileImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
			common.CSISecretsStoreProviderAzureContainerName:  csiSecretsStoreProviderAzureImageReference,
			common.CSISecretsStoreDriverContainerName:         csiSecretsStoreDriverImageReference,
			common.AzureArcOnboardingAddonName:                azureArcOnboardingImageReference,
			common.AzureKMSProviderComponentName:              azureKMSProviderImageReference,
		}
	case "1.17":
		ret = map[string]string{
			common.APIServerComponentName:                     getDefaultImage(common.APIServerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.ControllerManagerComponentName:             getDefaultImage(common.ControllerManagerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.KubeProxyAddonName, kubernetesImageBaseType) + ":v" + version,
			common.SchedulerComponentName:                     getDefaultImage(common.SchedulerComponentName, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        azureCloudControllerManagerImageReference,
			common.CloudNodeManagerAddonName:                  azureCloudNodeManagerImageReference,
			common.WindowsArtifactComponentName:               "v" + version + "/windowszip/v" + version + "-1int.zip",
			common.WindowsArtifactAzureStackComponentName:     "v" + version + common.AzureStackSuffix + "/windowszip/v" + version + common.AzureStackSuffix + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.DashboardMetricsScraperContainerName:       dashboardMetricsScraperImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.CSIProvisionerContainerName:                csiProvisionerImageReference,
			common.CSIAttacherContainerName:                   csiAttacherImageReference,
			common.CSILivenessProbeContainerName:              csiLivenessProbeImageReference,
			common.CSINodeDriverRegistrarContainerName:        csiNodeDriverRegistrarImageReference,
			common.CSISnapshotterContainerName:                csiSnapshotterImageReference,
			common.CSISnapshotControllerContainerName:         csiSnapshotControllerImageReference,
			common.CSIResizerContainerName:                    csiResizerImageReference,
			common.CSIAzureDiskContainerName:                  csiAzureDiskImageReference,
			common.CSIAzureFileContainerName:                  csiAzureFileImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
			common.CSISecretsStoreProviderAzureContainerName:  csiSecretsStoreProviderAzureImageReference,
			common.CSISecretsStoreDriverContainerName:         csiSecretsStoreDriverImageReference,
			common.AzureArcOnboardingAddonName:                azureArcOnboardingImageReference,
			common.AzureKMSProviderComponentName:              azureKMSProviderImageReference,
		}
	case "1.16":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.KubeProxyAddonName, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        azureCloudControllerManagerImageReference,
			common.CloudNodeManagerAddonName:                  azureCloudNodeManagerImageReference,
			common.WindowsArtifactComponentName:               "v" + version + "/windowszip/v" + version + "-1int.zip",
			common.WindowsArtifactAzureStackComponentName:     "v" + version + common.AzureStackSuffix + "/windowszip/v" + version + common.AzureStackSuffix + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.DashboardMetricsScraperContainerName:       dashboardMetricsScraperImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.CSIProvisionerContainerName:                csiProvisionerImageReference,
			common.CSIAttacherContainerName:                   csiAttacherImageReference,
			common.CSILivenessProbeContainerName:              csiLivenessProbeImageReference,
			common.CSINodeDriverRegistrarContainerName:        csiNodeDriverRegistrarImageReference,
			common.CSIResizerContainerName:                    csiResizerImageReference,
			common.CSIAzureDiskContainerName:                  csiAzureDiskImageReference,
			common.CSIAzureFileContainerName:                  csiAzureFileImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
			common.CSISecretsStoreProviderAzureContainerName:  csiSecretsStoreProviderAzureImageReference,
			common.CSISecretsStoreDriverContainerName:         csiSecretsStoreDriverImageReference,
			common.AzureArcOnboardingAddonName:                azureArcOnboardingImageReference,
			common.AzureKMSProviderComponentName:              azureKMSProviderImageReference,
		}
	case "1.15":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "/windowszip/v" + version + "-1int.zip",
			common.WindowsArtifactAzureStackComponentName:     "v" + version + common.AzureStackSuffix + "/windowszip/v" + version + common.AzureStackSuffix + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.DashboardMetricsScraperContainerName:       dashboardMetricsScraperImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.CSIProvisionerContainerName:                csiProvisionerImageReference,
			common.CSIAttacherContainerName:                   csiAttacherImageReference,
			common.CSILivenessProbeContainerName:              csiLivenessProbeImageReference,
			common.CSINodeDriverRegistrarContainerName:        csiNodeDriverRegistrarImageReference,
			common.CSIAzureDiskContainerName:                  csiAzureDiskImageReference,
			common.CSIAzureFileContainerName:                  csiAzureFileImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
			common.AzureArcOnboardingAddonName:                azureArcOnboardingImageReference,
			common.AzureKMSProviderComponentName:              azureKMSProviderImageReference,
		}
	case "1.14":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "/windowszip/v" + version + "-1int.zip",
			common.WindowsArtifactAzureStackComponentName:     "v" + version + common.AzureStackSuffix + "/windowszip/v" + version + common.AzureStackSuffix + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.DashboardMetricsScraperContainerName:       dashboardMetricsScraperImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.CSIProvisionerContainerName:                csiProvisionerImageReference,
			common.CSIAttacherContainerName:                   csiAttacherImageReference,
			common.CSILivenessProbeContainerName:              csiLivenessProbeImageReference,
			common.CSINodeDriverRegistrarContainerName:        csiNodeDriverRegistrarImageReference,
			common.CSIAzureDiskContainerName:                  csiAzureDiskImageReference,
			common.CSIAzureFileContainerName:                  csiAzureFileImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
			common.AzureKMSProviderComponentName:              azureKMSProviderImageReference,
		}
	case "1.13":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "/windowszip/v" + version + "-1int.zip",
			common.WindowsArtifactAzureStackComponentName:     "v" + version + common.AzureStackSuffix + "/windowszip/v" + version + common.AzureStackSuffix + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.CSIProvisionerContainerName:                csiProvisionerImageReference,
			common.CSIAttacherContainerName:                   csiAttacherImageReference,
			common.CSILivenessProbeContainerName:              csiLivenessProbeImageReference,
			common.CSINodeDriverRegistrarContainerName:        csiNodeDriverRegistrarImageReference,
			common.CSIAzureDiskContainerName:                  csiAzureDiskImageReference,
			common.CSIAzureFileContainerName:                  csiAzureFileImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
			common.AzureKMSProviderComponentName:              azureKMSProviderImageReference,
		}
	case "1.12":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.CoreDNSAddonName:                           getDefaultImage(common.CoreDNSAddonName, kubernetesImageBaseType),
			common.CoreDNSAutoscalerName:                      clusterProportionalAutoscalerImageReference,
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
		}
	case "1.11":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       getDefaultImage(common.ReschedulerAddonName, kubernetesImageBaseType),
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    getDefaultImage(common.DNSSidecarComponentName, kubernetesImageBaseType),
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
		}
	case "1.10":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       k8sComponent[common.ReschedulerAddonName],
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    k8sComponent[common.DNSSidecarComponentName],
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
			common.NVIDIADevicePluginAddonName:                nvidiaDevicePluginImageReference,
		}
	case "1.9":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.KubeDNSAddonName:                           getDefaultImage(common.KubeDNSAddonName, kubernetesImageBaseType),
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       getDefaultImage(common.DNSMasqComponentName, kubernetesImageBaseType),
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       k8sComponent[common.ReschedulerAddonName],
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.ClusterAutoscalerAddonName:                 k8sComponent[common.ClusterAutoscalerAddonName],
			common.DNSSidecarComponentName:                    k8sComponent[common.DNSSidecarComponentName],
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.NodeProblemDetectorAddonName:               nodeProblemDetectorImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
		}
	case "1.8":
		ret = map[string]string{
			common.Hyperkube:                                  getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:                         getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.CloudControllerManagerComponentName:        k8sComponent[common.CloudControllerManagerComponentName] + ":v" + version,
			common.WindowsArtifactComponentName:               "v" + version + "-1int.zip",
			common.DashboardAddonName:                         dashboardImageReference,
			common.ExecHealthZComponentName:                   getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:                  k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:                     k8sComponent[common.MetricsServerAddonName],
			common.KubeDNSAddonName:                           k8sComponent[common.KubeDNSAddonName],
			common.AddonManagerComponentName:                  k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:                       k8sComponent[common.DNSMasqComponentName],
			common.PauseComponentName:                         pauseImageReference,
			common.TillerAddonName:                            tillerImageReference,
			common.ReschedulerAddonName:                       k8sComponent[common.ReschedulerAddonName],
			common.ACIConnectorAddonName:                      virtualKubeletImageReference,
			common.AzureCNINetworkMonitorAddonName:            azureCNINetworkMonitorImageReference,
			common.BlobfuseFlexVolumeAddonName:                blobfuseFlexVolumeImageReference,
			common.SMBFlexVolumeAddonName:                     smbFlexVolumeImageReference,
			common.KeyVaultFlexVolumeAddonName:                keyvaultFlexVolumeImageReference,
			common.IPMASQAgentAddonName:                       getDefaultImage(common.IPMASQAgentAddonName, kubernetesImageBaseType),
			common.AzureNetworkPolicyAddonName:                azureNPMContainerImageReference,
			common.CalicoTyphaComponentName:                   calicoTyphaImageReference,
			common.CalicoCNIComponentName:                     calicoCNIImageReference,
			common.CalicoNodeComponentName:                    calicoNodeImageReference,
			common.CalicoPod2DaemonComponentName:              calicoPod2DaemonImageReference,
			common.CalicoClusterAutoscalerComponentName:       calicoClusterProportionalAutoscalerImageReference,
			common.CiliumAgentContainerName:                   ciliumAgentImageReference,
			common.CiliumCleanStateContainerName:              ciliumCleanStateImageReference,
			common.CiliumOperatorContainerName:                ciliumOperatorImageReference,
			common.CiliumEtcdOperatorContainerName:            ciliumEtcdOperatorImageReference,
			common.AntreaControllerContainerName:              antreaControllerImageReference,
			common.AntreaAgentContainerName:                   antreaAgentImageReference,
			common.AntreaOVSContainerName:                     antreaOVSImageReference,
			"antrea" + common.AntreaInstallCNIContainerName:   antreaInstallCNIImageReference,
			common.NMIContainerName:                           aadPodIdentityNMIImageReference,
			common.MICContainerName:                           aadPodIdentityMICImageReference,
			common.AzurePolicyAddonName:                       azurePolicyImageReference,
			common.GatekeeperContainerName:                    gatekeeperImageReference,
			common.KubeFlannelContainerName:                   kubeFlannelImageReference,
			"flannel" + common.FlannelInstallCNIContainerName: flannelInstallCNIImageReference,
			common.KubeRBACProxyContainerName:                 KubeRBACProxyImageReference,
			common.ScheduledMaintenanceManagerContainerName:   ScheduledMaintenanceManagerImageReference,
			"nodestatusfreq":                                  DefaultKubernetesNodeStatusUpdateFrequency,
			"nodegraceperiod":                                 DefaultKubernetesCtrlMgrNodeMonitorGracePeriod,
			"podeviction":                                     DefaultKubernetesCtrlMgrPodEvictionTimeout,
			"routeperiod":                                     DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
			"backoffretries":                                  strconv.Itoa(DefaultKubernetesCloudProviderBackoffRetries),
			"backoffjitter":                                   strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffJitter, 'f', -1, 64),
			"backoffduration":                                 strconv.Itoa(DefaultKubernetesCloudProviderBackoffDuration),
			"backoffexponent":                                 strconv.FormatFloat(DefaultKubernetesCloudProviderBackoffExponent, 'f', -1, 64),
			"ratelimitqps":                                    strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPS, 'f', -1, 64),
			"ratelimitqpswrite":                               strconv.FormatFloat(DefaultKubernetesCloudProviderRateLimitQPSWrite, 'f', -1, 64),
			"ratelimitbucket":                                 strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucket),
			"ratelimitbucketwrite":                            strconv.Itoa(DefaultKubernetesCloudProviderRateLimitBucketWrite),
			"gchighthreshold":                                 strconv.Itoa(DefaultKubernetesGCHighThreshold),
			"gclowthreshold":                                  strconv.Itoa(DefaultKubernetesGCLowThreshold),
		}
	case "1.7":
		ret = map[string]string{
			common.Hyperkube:                       getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:              getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.DashboardAddonName:              k8sComponent["dashboard"],
			common.ExecHealthZComponentName:        getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:       k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:          k8sComponent[common.MetricsServerAddonName],
			common.KubeDNSAddonName:                k8sComponent[common.KubeDNSAddonName],
			common.AddonManagerComponentName:       k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:            k8sComponent[common.DNSMasqComponentName],
			common.PauseComponentName:              pauseImageReference,
			common.TillerAddonName:                 tillerImageReference,
			common.ReschedulerAddonName:            k8sComponent[common.ReschedulerAddonName],
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
	case "1.6":
		ret = map[string]string{
			common.Hyperkube:                       getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.KubeProxyAddonName:              getDefaultImage(common.Hyperkube, kubernetesImageBaseType) + ":v" + version,
			common.DashboardAddonName:              k8sComponent["dashboard"],
			common.ExecHealthZComponentName:        getDefaultImage(common.ExecHealthZComponentName, kubernetesImageBaseType),
			common.AddonResizerComponentName:       k8sComponent[common.AddonResizerComponentName],
			common.MetricsServerAddonName:          k8sComponent[common.MetricsServerAddonName],
			common.KubeDNSAddonName:                k8sComponent[common.KubeDNSAddonName],
			common.AddonManagerComponentName:       k8sComponent[common.AddonManagerComponentName],
			common.DNSMasqComponentName:            k8sComponent[common.DNSMasqComponentName],
			common.PauseComponentName:              pauseImageReference,
			common.TillerAddonName:                 tillerImageReference,
			common.ReschedulerAddonName:            k8sComponent[common.ReschedulerAddonName],
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
