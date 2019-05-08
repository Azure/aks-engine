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

type kubernetesComponentFileSpec struct {
	sourceFile      string // filename to source spec data from
	base64Data      string // if not "", this base64-encoded string will take precedent over sourceFile
	destinationFile string // the filename to write to disk on the destination OS
	isEnabled       bool   // is this spec enabled?
}

func kubernetesContainerAddonSettingsInit(profile *api.Properties) map[string]kubernetesComponentFileSpec {
	return map[string]kubernetesComponentFileSpec{
		DefaultHeapsterAddonName: {
			sourceFile:      "kubernetesmasteraddons-heapster-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(HeapsterAddonName),
			destinationFile: "kube-heapster-deployment.yaml",
			isEnabled:       !common.IsKubernetesVersionGe(profile.OrchestratorProfile.OrchestratorVersion, "1.13.0"),
		},
		MetricsServerAddonName: {
			sourceFile:      "kubernetesmasteraddons-metrics-server-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(MetricsServerAddonName),
			destinationFile: "kube-metrics-server-deployment.yaml",
			isEnabled:       profile.OrchestratorProfile.IsMetricsServerEnabled(),
		},
		TillerAddonName: {
			sourceFile:      "kubernetesmasteraddons-tiller-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(TillerAddonName),
			destinationFile: "kube-tiller-deployment.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.IsTillerEnabled(),
		},
		AADPodIdentityAddonName: {
			sourceFile:      "kubernetesmasteraddons-aad-pod-identity-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AADPodIdentityAddonName),
			destinationFile: "aad-pod-identity-deployment.yaml",
			isEnabled:       profile.IsAADPodIdentityEnabled(),
		},
		ACIConnectorAddonName: {
			sourceFile:      "kubernetesmasteraddons-aci-connector-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(ACIConnectorAddonName),
			destinationFile: "aci-connector-deployment.yaml",
			isEnabled:       profile.IsACIConnectorEnabled(),
		},
		ClusterAutoscalerAddonName: {
			sourceFile:      "kubernetesmasteraddons-cluster-autoscaler-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(ClusterAutoscalerAddonName),
			destinationFile: "cluster-autoscaler-deployment.yaml",
			isEnabled:       profile.IsClusterAutoscalerEnabled(),
		},
		BlobfuseFlexVolumeAddonName: {
			sourceFile:      "kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(BlobfuseFlexVolumeAddonName),
			destinationFile: "blobfuse-flexvolume-installer.yaml",
			isEnabled:       profile.IsBlobfuseFlexVolumeEnabled() && !profile.HasCoreOS(),
		},

		SMBFlexVolumeAddonName: {
			sourceFile:      "kubernetesmasteraddons-smb-flexvolume-installer.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(SMBFlexVolumeAddonName),
			destinationFile: "smb-flexvolume-installer.yaml",
			isEnabled:       profile.IsSMBFlexVolumeEnabled() && !profile.HasCoreOS(),
		},
		KeyVaultFlexVolumeAddonName: {
			sourceFile:      "kubernetesmasteraddons-keyvault-flexvolume-installer.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(KeyVaultFlexVolumeAddonName),
			destinationFile: "keyvault-flexvolume-installer.yaml",
			isEnabled:       profile.IsKeyVaultFlexVolumeEnabled() && !profile.HasCoreOS(),
		},
		DashboardAddonName: {
			sourceFile:      "kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(DashboardAddonName),
			destinationFile: "kubernetes-dashboard-deployment.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.IsDashboardEnabled(),
		},
		ReschedulerAddonName: {
			sourceFile:      "kubernetesmasteraddons-kube-rescheduler-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(ReschedulerAddonName),
			destinationFile: "kube-rescheduler-deployment.yaml",
			isEnabled:       profile.IsReschedulerEnabled(),
		},
		NVIDIADevicePluginAddonName: {
			sourceFile:      "kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(NVIDIADevicePluginAddonName),
			destinationFile: "nvidia-device-plugin.yaml",
			isEnabled:       profile.IsNVIDIADevicePluginEnabled(),
		},
		ContainerMonitoringAddonName: {
			sourceFile:      "kubernetesmasteraddons-omsagent-daemonset.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(ContainerMonitoringAddonName),
			destinationFile: "omsagent-daemonset.yaml",
			isEnabled:       profile.IsContainerMonitoringEnabled(),
		},
		IPMASQAgentAddonName: {
			sourceFile:      "ip-masq-agent.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(IPMASQAgentAddonName),
			destinationFile: "ip-masq-agent.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.IsIPMasqAgentEnabled(),
		},
		AzureCNINetworkMonitorAddonName: {
			sourceFile:      "azure-cni-networkmonitor.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AzureCNINetworkMonitorAddonName),
			destinationFile: "azure-cni-networkmonitor.yaml",
			isEnabled:       profile.OrchestratorProfile.IsAzureCNI() && profile.OrchestratorProfile.KubernetesConfig.IsAzureCNIMonitoringEnabled(),
		},
		DNSAutoscalerAddonName: {
			sourceFile:      "dns-autoscaler.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(DNSAutoscalerAddonName),
			destinationFile: "dns-autoscaler.yaml",
			// TODO enable this when it has been smoke tested
			//common.IsKubernetesVersionGe(profile.OrchestratorProfile.OrchestratorVersion, "1.12.0"),
			isEnabled: false,
		},
		CalicoAddonName: {
			sourceFile:      "kubernetesmasteraddons-calico-daemonset.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(CalicoAddonName),
			destinationFile: "calico-daemonset.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.NetworkPolicy == NetworkPolicyCalico,
		},
	}
}

func kubernetesAddonSettingsInit(profile *api.Properties) []kubernetesComponentFileSpec {
	kubernetesFeatureSettings := []kubernetesComponentFileSpec{
		{
			sourceFile:      "kubernetesmasteraddons-kube-dns-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(KubeDNSAddonName),
			destinationFile: "kube-dns-deployment.yaml",
			isEnabled:       !common.IsKubernetesVersionGe(profile.OrchestratorProfile.OrchestratorVersion, "1.12.0"),
		},
		{
			sourceFile:      "coredns.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(CoreDNSAddonName),
			destinationFile: "coredns.yaml",
			isEnabled:       common.IsKubernetesVersionGe(profile.OrchestratorProfile.OrchestratorVersion, "1.12.0"),
		},
		{
			sourceFile:      "kubernetesmasteraddons-kube-proxy-daemonset.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(KubeProxyAddonName),
			destinationFile: "kube-proxy-daemonset.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      "kubernetesmasteraddons-azure-npm-daemonset.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AzureNetworkPolicyAddonName),
			destinationFile: "azure-npm-daemonset.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.NetworkPolicy == NetworkPolicyAzure && profile.OrchestratorProfile.KubernetesConfig.NetworkPlugin == NetworkPluginAzure,
		},
		{
			sourceFile:      "kubernetesmasteraddons-cilium-daemonset.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(CiliumAddonName),
			destinationFile: "cilium-daemonset.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.NetworkPolicy == NetworkPolicyCilium,
		},
		{
			sourceFile:      "kubernetesmasteraddons-flannel-daemonset.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(FlannelAddonName),
			destinationFile: "flannel-daemonset.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.NetworkPlugin == NetworkPluginFlannel,
		},
		{
			sourceFile:      "kubernetesmasteraddons-aad-default-admin-group-rbac.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AADAdminGroupAddonName),
			destinationFile: "aad-default-admin-group-rbac.yaml",
			isEnabled:       profile.AADProfile != nil && profile.AADProfile.AdminGroupID != "",
		},
		{
			sourceFile:      "kubernetesmasteraddons-azure-cloud-provider-deployment.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AzureCloudProviderAddonName),
			destinationFile: "azure-cloud-provider-deployment.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      "kubernetesmaster-audit-policy.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AuditPolicyAddonName),
			destinationFile: "audit-policy.yaml",
			isEnabled:       common.IsKubernetesVersionGe(profile.OrchestratorProfile.OrchestratorVersion, "1.8.0"),
		},
		{
			sourceFile:      "kubernetesmasteraddons-elb-svc.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(ELBServiceAddonName),
			destinationFile: "elb-svc.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == "Standard" && !to.Bool(profile.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled),
		},
		{
			sourceFile:      "kubernetesmasteraddons-pod-security-policy.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.PodSecurityPolicyConfig["data"],
			destinationFile: "pod-security-policy.yaml",
			isEnabled:       to.Bool(profile.OrchestratorProfile.KubernetesConfig.EnablePodSecurityPolicy),
		},
	}

	unmanagedStorageClassesSourceYaml := "kubernetesmasteraddons-unmanaged-azure-storage-classes.yaml"
	managedStorageClassesSourceYaml := "kubernetesmasteraddons-managed-azure-storage-classes.yaml"

	if profile.IsAzureStackCloud() {
		unmanagedStorageClassesSourceYaml = "kubernetesmasteraddons-unmanaged-azure-storage-classes-custom.yaml"
		managedStorageClassesSourceYaml = "kubernetesmasteraddons-managed-azure-storage-classes-custom.yaml"
	}

	if len(profile.AgentPoolProfiles) > 0 {
		kubernetesFeatureSettings = append(kubernetesFeatureSettings,
			kubernetesComponentFileSpec{
				sourceFile:      unmanagedStorageClassesSourceYaml,
				base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AzureStorageClassesAddonName),
				destinationFile: "azure-storage-classes.yaml",
				isEnabled:       profile.AgentPoolProfiles[0].StorageProfile != api.ManagedDisks,
			})
		kubernetesFeatureSettings = append(kubernetesFeatureSettings,
			kubernetesComponentFileSpec{
				sourceFile:      managedStorageClassesSourceYaml,
				base64Data:      profile.OrchestratorProfile.KubernetesConfig.GetAddonScript(AzureStorageClassesAddonName),
				destinationFile: "azure-storage-classes.yaml",
				isEnabled:       profile.AgentPoolProfiles[0].StorageProfile == api.ManagedDisks,
			})
	}

	return kubernetesFeatureSettings
}

func kubernetesManifestSettingsInit(profile *api.Properties) []kubernetesComponentFileSpec {
	kubeControllerManagerYaml := "kubernetesmaster-kube-controller-manager.yaml"

	if profile.IsAzureStackCloud() {
		kubeControllerManagerYaml = "kubernetesmaster-kube-controller-manager-custom.yaml"
	}

	return []kubernetesComponentFileSpec{
		{
			sourceFile:      "kubernetesmaster-kube-scheduler.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.SchedulerConfig["data"],
			destinationFile: "kube-scheduler.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      kubeControllerManagerYaml,
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig["data"],
			destinationFile: "kube-controller-manager.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      "kubernetesmaster-cloud-controller-manager.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig["data"],
			destinationFile: "cloud-controller-manager.yaml",
			isEnabled:       profile.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager != nil && *profile.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager,
		},
		{
			sourceFile:      "kubernetesmaster-kube-apiserver.yaml",
			base64Data:      profile.OrchestratorProfile.KubernetesConfig.APIServerConfig["data"],
			destinationFile: "kube-apiserver.yaml",
			isEnabled:       true,
		},
		{
			sourceFile:      "kubernetesmaster-kube-addon-manager.yaml",
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

func substituteConfigString(input string, kubernetesFeatureSettings []kubernetesComponentFileSpec, sourcePath string, destinationPath string, placeholder string, orchestratorVersion string) string {
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
				config += getAddonString(cscript, setting.destinationFile, destinationPath)
			} else {
				cscript = getCustomScriptFromFile(setting.sourceFile,
					sourcePath,
					versions[0]+"."+versions[1])
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

func getCustomScriptFromFile(sourceFile, sourcePath, version string) string {
	customDataFilePath := getCustomDataFilePath(sourceFile, sourcePath, version)
	return getBase64EncodedGzippedCustomScript(customDataFilePath)
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
