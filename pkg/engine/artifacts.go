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

func kubernetesAddonSettingsInit(p *api.Properties) map[string]kubernetesComponentFileSpec {
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
			sourceFile:      heapsterAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.HeapsterAddonName),
			destinationFile: heapsterAddonDestinationFilename,
		},
		common.MetricsServerAddonName: {
			sourceFile:      metricsServerAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.MetricsServerAddonName),
			destinationFile: metricsServerAddonDestinationFilename,
		},
		common.TillerAddonName: {
			sourceFile:      tillerAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.TillerAddonName),
			destinationFile: tillerAddonDestinationFilename,
		},
		common.AADPodIdentityAddonName: {
			sourceFile:      aadPodIdentityAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AADPodIdentityAddonName),
			destinationFile: aadPodIdentityAddonDestinationFilename,
		},
		common.ACIConnectorAddonName: {
			sourceFile:      aciConnectorAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.ACIConnectorAddonName),
			destinationFile: aciConnectorAddonDestinationFilename,
		},
		common.AzureDiskCSIDriverAddonName: {
			sourceFile:      azureDiskCSIAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AzureDiskCSIDriverAddonName),
			destinationFile: azureDiskCSIAddonDestinationFilename,
		},
		common.AzureFileCSIDriverAddonName: {
			sourceFile:      azureFileCSIAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AzureFileCSIDriverAddonName),
			destinationFile: azureFileCSIAddonDestinationFilename,
		},
		common.ClusterAutoscalerAddonName: {
			sourceFile:      clusterAutoscalerAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.ClusterAutoscalerAddonName),
			destinationFile: clusterAutoscalerAddonDestinationFilename,
		},
		common.BlobfuseFlexVolumeAddonName: {
			sourceFile:      blobfuseFlexVolumeAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.BlobfuseFlexVolumeAddonName),
			destinationFile: blobfuseFlexVolumeAddonDestinationFilename,
		},
		common.SMBFlexVolumeAddonName: {
			sourceFile:      smbFlexVolumeAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.SMBFlexVolumeAddonName),
			destinationFile: smbFlexVolumeAddonDestinationFilename,
		},
		common.KeyVaultFlexVolumeAddonName: {
			sourceFile:      keyvaultFlexVolumeAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.KeyVaultFlexVolumeAddonName),
			destinationFile: keyvaultFlexVolumeAddonDestinationFilename,
		},
		common.DashboardAddonName: {
			sourceFile:      dashboardAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.DashboardAddonName),
			destinationFile: dashboardAddonDestinationFilename,
		},
		common.ReschedulerAddonName: {
			sourceFile:      reschedulerAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.ReschedulerAddonName),
			destinationFile: reschedulerAddonDestinationFilename,
		},
		common.NVIDIADevicePluginAddonName: {
			sourceFile:      nvidiaAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.NVIDIADevicePluginAddonName),
			destinationFile: nvidiaAddonDestinationFilename,
		},
		common.ContainerMonitoringAddonName: {
			sourceFile:      containerMonitoringAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.ContainerMonitoringAddonName),
			destinationFile: containerMonitoringAddonDestinationFilename,
		},
		common.IPMASQAgentAddonName: {
			sourceFile:      ipMasqAgentAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.IPMASQAgentAddonName),
			destinationFile: ipMasqAgentAddonDestinationFilename,
		},
		common.AzureCNINetworkMonitorAddonName: {
			sourceFile:      azureCNINetworkMonitorAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AzureCNINetworkMonitorAddonName),
			destinationFile: azureCNINetworkMonitorAddonDestinationFilename,
		},
		common.DNSAutoscalerAddonName: {
			sourceFile:      dnsAutoscalerAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.DNSAutoscalerAddonName),
			destinationFile: dnsAutoscalerAddonDestinationFilename,
		},
		common.CalicoAddonName: {
			sourceFile:      calicoAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.CalicoAddonName),
			destinationFile: calicoAddonDestinationFilename,
		},
		common.AzureNetworkPolicyAddonName: {
			sourceFile:      azureNetworkPolicyAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AzureNetworkPolicyAddonName),
			destinationFile: azureNetworkPolicyAddonDestinationFilename,
		},
		common.AzurePolicyAddonName: {
			sourceFile:      azurePolicyAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AzurePolicyAddonName),
			destinationFile: azurePolicyAddonDestinationFilename,
		},
		common.CloudNodeManagerAddonName: {
			sourceFile:      cloudNodeManagerAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.CloudNodeManagerAddonName),
			destinationFile: cloudNodeManagerAddonDestinationFilename,
		},
		common.NodeProblemDetectorAddonName: {
			sourceFile:      nodeProblemDetectorAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.NodeProblemDetectorAddonName),
			destinationFile: nodeProblemDetectorAddonDestinationFilename,
		},
		common.KubeDNSAddonName: {
			sourceFile:      kubeDNSAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.KubeDNSAddonName),
			destinationFile: kubeDNSAddonDestinationFilename,
		},
		common.CoreDNSAddonName: {
			sourceFile:      corednsAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.CoreDNSAddonName),
			destinationFile: corednsAddonDestinationFilename,
		},
		common.KubeProxyAddonName: {
			sourceFile:      kubeProxyAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.KubeProxyAddonName),
			destinationFile: kubeProxyAddonDestinationFilename,
		},
		common.PodSecurityPolicyAddonName: {
			sourceFile:      podSecurityPolicyAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.PodSecurityPolicyAddonName),
			destinationFile: podSecurityPolicyAddonDestinationFilename,
		},
		common.AADAdminGroupAddonName: {
			sourceFile:      aadDefaultAdminGroupAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AADAdminGroupAddonName),
			destinationFile: aadDefaultAdminGroupDestinationFilename,
		},
		common.CiliumAddonName: {
			sourceFile:      ciliumAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.CiliumAddonName),
			destinationFile: ciliumAddonDestinationFilename,
		},
		common.AntreaAddonName: {
			sourceFile:      antreaAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AntreaAddonName),
			destinationFile: antreaAddonDestinationFilename,
		},
		common.AuditPolicyAddonName: {
			sourceFile:      auditPolicyAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AuditPolicyAddonName),
			destinationFile: auditPolicyAddonDestinationFilename,
		},
		common.AzureCloudProviderAddonName: {
			sourceFile:      cloudProviderAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.AuditPolicyAddonName),
			destinationFile: cloudProviderAddonDestinationFilename,
		},
		common.FlannelAddonName: {
			sourceFile:      flannelAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.FlannelAddonName),
			destinationFile: flannelAddonDestinationFilename,
		},
		common.ScheduledMaintenanceAddonName: {
			sourceFile:      scheduledMaintenanceAddonSourceFilename,
			base64Data:      k.GetAddonScript(common.ScheduledMaintenanceAddonName),
			destinationFile: scheduledMaintenanceAddonSourceFilename,
		},
	}
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
