// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"
	"strings"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
)

// kubernetesComponentFileSpec defines a k8s component that we will deliver via file to a master node vm
type kubernetesComponentFileSpec struct {
	sourceFile      string // filename to source spec data from
	base64Data      string // if not "", this base64-encoded string will take precedent over sourceFile
	destinationFile string // the filename to write to disk on the destination OS
	isEnabled       bool   // is this spec enabled?
}

func kubernetesContainerAddonSettingsInit(p *api.Properties) map[string]kubernetesComponentFileSpec {
	if p.OrchestratorProfile == nil {
		p.OrchestratorProfile = &api.OrchestratorProfile{}
	}
	if p.OrchestratorProfile.KubernetesConfig == nil {
		p.OrchestratorProfile.KubernetesConfig = &api.KubernetesConfig{}
	}
	o := p.OrchestratorProfile
	k := o.KubernetesConfig
	// TODO validate that each of these addons are actually wired in to the conveniences in getAddonFuncMap
	return map[string]kubernetesComponentFileSpec{
		common.HeapsterAddonName: {
			sourceFile:      "kubernetesmasteraddons-heapster-deployment.yaml",
			base64Data:      k.GetAddonScript(common.HeapsterAddonName),
			destinationFile: "kube-heapster-deployment.yaml",
		},
		common.MetricsServerAddonName: {
			sourceFile:      "kubernetesmasteraddons-metrics-server-deployment.yaml",
			base64Data:      k.GetAddonScript(common.MetricsServerAddonName),
			destinationFile: "kube-metrics-server-deployment.yaml",
		},
		common.TillerAddonName: {
			sourceFile:      "kubernetesmasteraddons-tiller-deployment.yaml",
			base64Data:      k.GetAddonScript(common.TillerAddonName),
			destinationFile: "kube-tiller-deployment.yaml",
		},
		common.AADPodIdentityAddonName: {
			sourceFile:      "kubernetesmasteraddons-aad-pod-identity-deployment.yaml",
			base64Data:      k.GetAddonScript(common.AADPodIdentityAddonName),
			destinationFile: "aad-pod-identity-deployment.yaml",
		},
		common.ACIConnectorAddonName: {
			sourceFile:      "kubernetesmasteraddons-aci-connector-deployment.yaml",
			base64Data:      k.GetAddonScript(common.ACIConnectorAddonName),
			destinationFile: "aci-connector-deployment.yaml",
		},
		common.AzureDiskCSIDriverAddonName: {
			sourceFile:      "kubernetesmasteraddons-azuredisk-csi-driver-deployment.yaml",
			base64Data:      k.GetAddonScript(common.AzureDiskCSIDriverAddonName),
			destinationFile: "azuredisk-csi-driver-deployment.yaml",
		},
		common.AzureFileCSIDriverAddonName: {
			sourceFile:      "kubernetesmasteraddons-azurefile-csi-driver-deployment.yaml",
			base64Data:      k.GetAddonScript(common.AzureFileCSIDriverAddonName),
			destinationFile: "azurefile-csi-driver-deployment.yaml",
		},
		common.ClusterAutoscalerAddonName: {
			sourceFile:      "kubernetesmasteraddons-cluster-autoscaler-deployment.yaml",
			base64Data:      k.GetAddonScript(common.ClusterAutoscalerAddonName),
			destinationFile: "cluster-autoscaler-deployment.yaml",
		},
		common.BlobfuseFlexVolumeAddonName: {
			sourceFile:      "kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml",
			base64Data:      k.GetAddonScript(common.BlobfuseFlexVolumeAddonName),
			destinationFile: "blobfuse-flexvolume-installer.yaml",
		},
		common.SMBFlexVolumeAddonName: {
			sourceFile:      "kubernetesmasteraddons-smb-flexvolume-installer.yaml",
			base64Data:      k.GetAddonScript(common.SMBFlexVolumeAddonName),
			destinationFile: "smb-flexvolume-installer.yaml",
		},
		common.KeyVaultFlexVolumeAddonName: {
			sourceFile:      "kubernetesmasteraddons-keyvault-flexvolume-installer.yaml",
			base64Data:      k.GetAddonScript(common.KeyVaultFlexVolumeAddonName),
			destinationFile: "keyvault-flexvolume-installer.yaml",
		},
		common.DashboardAddonName: {
			sourceFile:      "kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml",
			base64Data:      k.GetAddonScript(common.DashboardAddonName),
			destinationFile: "kubernetes-dashboard-deployment.yaml",
		},
		common.ReschedulerAddonName: {
			sourceFile:      "kubernetesmasteraddons-kube-rescheduler-deployment.yaml",
			base64Data:      k.GetAddonScript(common.ReschedulerAddonName),
			destinationFile: "kube-rescheduler-deployment.yaml",
		},
		common.NVIDIADevicePluginAddonName: {
			sourceFile:      "kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml",
			base64Data:      k.GetAddonScript(common.NVIDIADevicePluginAddonName),
			destinationFile: "nvidia-device-plugin.yaml",
		},
		common.ContainerMonitoringAddonName: {
			sourceFile:      "kubernetesmasteraddons-omsagent-daemonset.yaml",
			base64Data:      k.GetAddonScript(common.ContainerMonitoringAddonName),
			destinationFile: "omsagent-daemonset.yaml",
		},
		common.IPMASQAgentAddonName: {
			sourceFile:      "ip-masq-agent.yaml",
			base64Data:      k.GetAddonScript(common.IPMASQAgentAddonName),
			destinationFile: "ip-masq-agent.yaml",
		},
		common.AzureCNINetworkMonitorAddonName: {
			sourceFile:      "azure-cni-networkmonitor.yaml",
			base64Data:      k.GetAddonScript(common.AzureCNINetworkMonitorAddonName),
			destinationFile: "azure-cni-networkmonitor.yaml",
		},
		common.DNSAutoscalerAddonName: {
			sourceFile:      "dns-autoscaler.yaml",
			base64Data:      k.GetAddonScript(common.DNSAutoscalerAddonName),
			destinationFile: "dns-autoscaler.yaml",
		},
		common.CalicoAddonName: {
			sourceFile:      "kubernetesmasteraddons-calico-daemonset.yaml",
			base64Data:      k.GetAddonScript(common.CalicoAddonName),
			destinationFile: "calico-daemonset.yaml",
		},
		common.AzureNetworkPolicyAddonName: {
			sourceFile:      "kubernetesmasteraddons-azure-npm-daemonset.yaml",
			base64Data:      k.GetAddonScript(common.AzureNetworkPolicyAddonName),
			destinationFile: "azure-npm-daemonset.yaml",
		},
		common.AzurePolicyAddonName: {
			sourceFile:      "azure-policy-deployment.yaml",
			base64Data:      k.GetAddonScript(common.AzurePolicyAddonName),
			destinationFile: "azure-policy-deployment.yaml",
		},
		common.CloudNodeManagerAddonName: {
			sourceFile:      "kubernetesmasteraddons-cloud-node-manager.yaml",
			base64Data:      "",
			destinationFile: "cloud-node-manager.yaml",
		},
		common.NodeProblemDetectorAddonName: {
			sourceFile:      "node-problem-detector.yaml",
			base64Data:      k.GetAddonScript(common.NodeProblemDetectorAddonName),
			destinationFile: "node-problem-detector.yaml",
			isEnabled:       k.IsAddonEnabled(common.NodeProblemDetectorAddonName),
		},
	}
}

func kubernetesAddonSettingsInit(p *api.Properties) []kubernetesComponentFileSpec {
	if p.OrchestratorProfile == nil {
		p.OrchestratorProfile = &api.OrchestratorProfile{}
	}
	if p.OrchestratorProfile.KubernetesConfig == nil {
		p.OrchestratorProfile.KubernetesConfig = &api.KubernetesConfig{}
	}
	o := p.OrchestratorProfile
	k := o.KubernetesConfig
	kubernetesComponentFileSpecs := []kubernetesComponentFileSpec{
		{
			sourceFile:      "kubernetesmasteraddons-kube-dns-deployment.yaml",
			base64Data:      k.GetAddonScript(common.KubeDNSAddonName),
			destinationFile: "kube-dns-deployment.yaml",
			isEnabled:       !common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.12.0"),
		},
		{
			sourceFile:      "coredns.yaml",
			base64Data:      k.GetAddonScript(common.CoreDNSAddonName),
			destinationFile: "coredns.yaml",
			isEnabled:       common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.12.0"),
		},
		{
			sourceFile:      "kubernetesmasteraddons-kube-proxy-daemonset.yaml",
			base64Data:      k.GetAddonScript(common.KubeProxyAddonName),
			destinationFile: "kube-proxy-daemonset.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      "kubernetesmasteraddons-cilium-daemonset.yaml",
			base64Data:      k.GetAddonScript(common.CiliumAddonName),
			destinationFile: "cilium-daemonset.yaml",
			isEnabled:       k.NetworkPolicy == NetworkPolicyCilium,
		},
		{
			sourceFile:      "kubernetesmasteraddons-flannel-daemonset.yaml",
			base64Data:      k.GetAddonScript(common.FlannelAddonName),
			destinationFile: "flannel-daemonset.yaml",
			isEnabled:       k.NetworkPlugin == NetworkPluginFlannel,
		},
		{
			sourceFile:      "kubernetesmasteraddons-aad-default-admin-group-rbac.yaml",
			base64Data:      k.GetAddonScript(common.AADAdminGroupAddonName),
			destinationFile: "aad-default-admin-group-rbac.yaml",
			isEnabled:       p.AADProfile != nil && p.AADProfile.AdminGroupID != "",
		},
		{
			sourceFile:      "kubernetesmasteraddons-azure-cloud-provider-deployment.yaml",
			base64Data:      k.GetAddonScript(common.AzureCloudProviderAddonName),
			destinationFile: "azure-cloud-provider-deployment.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      "kubernetesmaster-audit-policy.yaml",
			base64Data:      k.GetAddonScript(common.AuditPolicyAddonName),
			destinationFile: "audit-policy.yaml",
			isEnabled:       common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.8.0"),
		},
		{
			sourceFile:      "kubernetesmasteraddons-pod-security-policy.yaml",
			base64Data:      k.GetAddonScript(common.PodSecurityPolicyAddonName),
			destinationFile: "pod-security-policy.yaml",
			isEnabled:       to.Bool(p.OrchestratorProfile.KubernetesConfig.EnablePodSecurityPolicy) || common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.15.0-beta.1"),
		},
		{
			sourceFile:      "kubernetesmasteraddons-scheduled-maintenance-deployment.yaml",
			base64Data:      k.GetAddonScript(common.ScheduledMaintenanceAddonName),
			destinationFile: "scheduled-maintenance-deployment.yaml",
			isEnabled:       k.IsAddonEnabled(common.ScheduledMaintenanceAddonName),
		},
	}

	if len(p.AgentPoolProfiles) > 0 {
		if to.Bool(k.UseCloudControllerManager) {
			kubernetesComponentFileSpecs = append(kubernetesComponentFileSpecs,
				kubernetesComponentFileSpec{
					sourceFile:      "kubernetesmasteraddons-azure-csi-storage-classes.yaml",
					base64Data:      k.GetAddonScript(common.AzureCSIStorageClassesAddonName),
					destinationFile: "azure-csi-storage-classes.yaml",
					isEnabled:       true,
				})
		} else {
			// Use built-in storage classes if CCM is disabled
			unmanagedStorageClassesSourceYaml := "kubernetesmasteraddons-unmanaged-azure-storage-classes.yaml"
			managedStorageClassesSourceYaml := "kubernetesmasteraddons-managed-azure-storage-classes.yaml"
			if p.IsAzureStackCloud() {
				unmanagedStorageClassesSourceYaml = "kubernetesmasteraddons-unmanaged-azure-storage-classes-custom.yaml"
				managedStorageClassesSourceYaml = "kubernetesmasteraddons-managed-azure-storage-classes-custom.yaml"
			}

			kubernetesComponentFileSpecs = append(kubernetesComponentFileSpecs,
				kubernetesComponentFileSpec{
					sourceFile:      unmanagedStorageClassesSourceYaml,
					base64Data:      k.GetAddonScript(common.AzureStorageClassesAddonName),
					destinationFile: "azure-storage-classes.yaml",
					isEnabled:       p.AgentPoolProfiles[0].StorageProfile == api.StorageAccount,
				})
			kubernetesComponentFileSpecs = append(kubernetesComponentFileSpecs,
				kubernetesComponentFileSpec{
					sourceFile:      managedStorageClassesSourceYaml,
					base64Data:      k.GetAddonScript(common.AzureStorageClassesAddonName),
					destinationFile: "azure-storage-classes.yaml",
					isEnabled:       p.AgentPoolProfiles[0].StorageProfile == api.ManagedDisks,
				})
		}
	}

	return kubernetesComponentFileSpecs
}

func kubernetesManifestSettingsInit(p *api.Properties) []kubernetesComponentFileSpec {
	if p.OrchestratorProfile == nil {
		p.OrchestratorProfile = &api.OrchestratorProfile{}
	}
	if p.OrchestratorProfile.KubernetesConfig == nil {
		p.OrchestratorProfile.KubernetesConfig = &api.KubernetesConfig{}
	}
	o := p.OrchestratorProfile
	k := o.KubernetesConfig
	if k.SchedulerConfig == nil {
		k.SchedulerConfig = map[string]string{}
	}
	if k.ControllerManagerConfig == nil {
		k.ControllerManagerConfig = map[string]string{}
	}
	if k.CloudControllerManagerConfig == nil {
		k.CloudControllerManagerConfig = map[string]string{}
	}
	if k.APIServerConfig == nil {
		k.APIServerConfig = map[string]string{}
	}
	kubeControllerManagerYaml := kubeControllerManagerManifestFilename

	if p.IsAzureStackCloud() {
		kubeControllerManagerYaml = kubeControllerManagerCustomManifestFilename
	}

	return []kubernetesComponentFileSpec{
		{
			sourceFile:      kubeSchedulerManifestFilename,
			base64Data:      k.SchedulerConfig["data"],
			destinationFile: "kube-scheduler.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      kubeControllerManagerYaml,
			base64Data:      k.ControllerManagerConfig["data"],
			destinationFile: "kube-controller-manager.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      ccmManifestFilename,
			base64Data:      k.CloudControllerManagerConfig["data"],
			destinationFile: "cloud-controller-manager.yaml",
			isEnabled:       to.Bool(k.UseCloudControllerManager),
		},
		{
			sourceFile:      kubeAPIServerManifestFilename,
			base64Data:      k.APIServerConfig["data"],
			destinationFile: "kube-apiserver.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      kubeAddonManagerManifestFilename,
			base64Data:      "", // arbitrary user-provided data not enabled for kube-addon-manager spec
			destinationFile: "kube-addon-manager.yaml",
			isEnabled:       true,
		},
	}
}

func getAddonString(input, destinationPath, destinationFile string) string {
	addonString := getBase64EncodedGzippedCustomScriptFromStr(input)
	return buildConfigString(addonString, destinationFile, destinationPath)
}

func substituteConfigString(input string, kubernetesFeatureSettings []kubernetesComponentFileSpec, sourcePath string, destinationPath string, placeholder string, orchestratorVersion string, cs *api.ContainerService) string {
	var config string

	versions := strings.Split(orchestratorVersion, ".")
	for _, setting := range kubernetesFeatureSettings {
		if setting.isEnabled {
			var cscript string
			if setting.base64Data != "" {
				var err error
				cscript, err = getStringFromBase64(setting.base64Data)
				if err != nil {
					return ""
				}
				config += getAddonString(cscript, destinationPath, setting.destinationFile)
			} else {
				cscript = getCustomScriptFromFile(setting.sourceFile,
					sourcePath,
					versions[0]+"."+versions[1], cs)
				config += buildConfigString(
					cscript,
					setting.destinationFile,
					destinationPath)
			}
		}
	}

	return strings.Replace(input, placeholder, config, -1)
}

func buildConfigString(configString, destinationFile, destinationPath string) string {
	contents := []string{
		fmt.Sprintf("- path: %s/%s", destinationPath, destinationFile),
		"  permissions: \\\"0644\\\"",
		"  encoding: gzip",
		"  owner: \\\"root\\\"",
		"  content: !!binary |",
		fmt.Sprintf("    %s\\n\\n", configString),
	}

	return strings.Join(contents, "\\n")
}

func getCustomScriptFromFile(sourceFile, sourcePath, version string, cs *api.ContainerService) string {
	customDataFilePath := getCustomDataFilePath(sourceFile, sourcePath, version)
	return getBase64EncodedGzippedCustomScript(customDataFilePath, cs)
}

func getCustomDataFilePath(sourceFile, sourcePath, version string) string {
	sourceFileFullPath := sourcePath + "/" + sourceFile
	sourceFileFullPathVersioned := sourcePath + "/" + version + "/" + sourceFile

	// Test to check if the versioned file can be read.
	_, err := Asset(sourceFileFullPathVersioned)
	if err == nil {
		sourceFileFullPath = sourceFileFullPathVersioned
	}
	return sourceFileFullPath
}
