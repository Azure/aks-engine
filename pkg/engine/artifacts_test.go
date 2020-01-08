// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestKubernetesAddonSettingsInit(t *testing.T) {
	/*
		$ echo "Hello, World\!" | base64
		SGVsbG8sIFdvcmxkXCEK
	*/
	const base64Data = "SGVsbG8sIFdvcmxkXCEK"
	cases := []struct {
		name                           string
		p                              *api.Properties
		expectedHeapster               kubernetesComponentFileSpec
		expectedMetricsServer          kubernetesComponentFileSpec
		expectedTiller                 kubernetesComponentFileSpec
		expectedAADPodIdentity         kubernetesComponentFileSpec
		expectedACIConnector           kubernetesComponentFileSpec
		expectedAzureDiskCSIDriver     kubernetesComponentFileSpec
		expectedAzureFileCSIDriver     kubernetesComponentFileSpec
		expectedClusterAutoscaler      kubernetesComponentFileSpec
		expectedBlobFlexVolume         kubernetesComponentFileSpec
		expectedSMBFlexVolume          kubernetesComponentFileSpec
		expectedKeyVaultFlexVolume     kubernetesComponentFileSpec
		expectedDashboard              kubernetesComponentFileSpec
		expectedRescheduler            kubernetesComponentFileSpec
		expectedNvidia                 kubernetesComponentFileSpec
		expectedContainerMonitoring    kubernetesComponentFileSpec
		expectedIPMasqAgent            kubernetesComponentFileSpec
		expectedAzureCNINetworkMonitor kubernetesComponentFileSpec
		expectedDNSAutoscaler          kubernetesComponentFileSpec
		expectedCalico                 kubernetesComponentFileSpec
		expectedCilium                 kubernetesComponentFileSpec
		expectedAzureNetworkPolicy     kubernetesComponentFileSpec
		expectedAzurePolicy            kubernetesComponentFileSpec
		expectedCloudNodeManager       kubernetesComponentFileSpec
		expectedNodeProblemDetector    kubernetesComponentFileSpec
		expectedKubeDNS                kubernetesComponentFileSpec
		expectedCoreDNS                kubernetesComponentFileSpec
		expectedKubeProxy              kubernetesComponentFileSpec
		expectedPodSecurityPolicy      kubernetesComponentFileSpec
		expectedAADDefaultAdminGroup   kubernetesComponentFileSpec
		expectedAntrea                 kubernetesComponentFileSpec
		expectedAuditPolicy            kubernetesComponentFileSpec
		expectedAzureCloudProvider     kubernetesComponentFileSpec
		expectedFlannel                kubernetesComponentFileSpec
	}{
		{
			name: "addons with data",
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.16.1",
					KubernetesConfig: &api.KubernetesConfig{
						Addons: []api.KubernetesAddon{
							{
								Name: common.HeapsterAddonName,
								Data: base64Data,
							},
							{
								Name: common.MetricsServerAddonName,
								Data: base64Data,
							},
							{
								Name: common.TillerAddonName,
								Data: base64Data,
							},
							{
								Name: common.AADPodIdentityAddonName,
								Data: base64Data,
							},
							{
								Name: common.ACIConnectorAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureDiskCSIDriverAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureFileCSIDriverAddonName,
								Data: base64Data,
							},
							{
								Name: common.ClusterAutoscalerAddonName,
								Data: base64Data,
							},
							{
								Name: common.BlobfuseFlexVolumeAddonName,
								Data: base64Data,
							},
							{
								Name: common.SMBFlexVolumeAddonName,
								Data: base64Data,
							},
							{
								Name: common.KeyVaultFlexVolumeAddonName,
								Data: base64Data,
							},
							{
								Name: common.DashboardAddonName,
								Data: base64Data,
							},
							{
								Name: common.ReschedulerAddonName,
								Data: base64Data,
							},
							{
								Name: common.NVIDIADevicePluginAddonName,
								Data: base64Data,
							},
							{
								Name: common.ContainerMonitoringAddonName,
								Data: base64Data,
							},
							{
								Name: common.IPMASQAgentAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureCNINetworkMonitorAddonName,
								Data: base64Data,
							},
							{
								Name: common.DNSAutoscalerAddonName,
								Data: base64Data,
							},
							{
								Name: common.CalicoAddonName,
								Data: base64Data,
							},
							{
								Name: common.CiliumAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureNetworkPolicyAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzurePolicyAddonName,
								Data: base64Data,
							},
							{
								Name: common.CloudNodeManagerAddonName,
								Data: base64Data,
							},
							{
								Name: common.NodeProblemDetectorAddonName,
								Data: base64Data,
							},
							{
								Name: common.KubeDNSAddonName,
								Data: base64Data,
							},
							{
								Name: common.CoreDNSAddonName,
								Data: base64Data,
							},
							{
								Name: common.KubeProxyAddonName,
								Data: base64Data,
							},
							{
								Name: common.PodSecurityPolicyAddonName,
								Data: base64Data,
							},
							{
								Name: common.AADAdminGroupAddonName,
								Data: base64Data,
							},
							{
								Name: common.AntreaAddonName,
								Data: base64Data,
							},
							{
								Name: common.AuditPolicyAddonName,
								Data: base64Data,
							},
							{
								Name: common.AzureCloudProviderAddonName,
								Data: base64Data,
							},
							{
								Name: common.FlannelAddonName,
								Data: base64Data,
							},
						},
					},
				},
			},
			expectedHeapster: kubernetesComponentFileSpec{
				sourceFile:      heapsterAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: heapsterAddonDestinationFilename,
			},
			expectedMetricsServer: kubernetesComponentFileSpec{
				sourceFile:      metricsServerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: metricsServerAddonDestinationFilename,
			},
			expectedTiller: kubernetesComponentFileSpec{
				sourceFile:      tillerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: tillerAddonDestinationFilename,
			},
			expectedAADPodIdentity: kubernetesComponentFileSpec{
				sourceFile:      aadPodIdentityAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: aadPodIdentityAddonDestinationFilename,
			},
			expectedACIConnector: kubernetesComponentFileSpec{
				sourceFile:      aciConnectorAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: aciConnectorAddonDestinationFilename,
			},
			expectedAzureDiskCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureDiskCSIAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureDiskCSIAddonDestinationFilename,
			},
			expectedAzureFileCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureFileCSIAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureFileCSIAddonDestinationFilename,
			},
			expectedClusterAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      clusterAutoscalerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: clusterAutoscalerAddonDestinationFilename,
			},
			expectedBlobFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      blobfuseFlexVolumeAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: blobfuseFlexVolumeAddonDestinationFilename,
			},
			expectedSMBFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      smbFlexVolumeAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: smbFlexVolumeAddonDestinationFilename,
			},
			expectedKeyVaultFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      keyvaultFlexVolumeAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: keyvaultFlexVolumeAddonDestinationFilename,
			},
			expectedDashboard: kubernetesComponentFileSpec{
				sourceFile:      dashboardAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: dashboardAddonDestinationFilename,
			},
			expectedRescheduler: kubernetesComponentFileSpec{
				sourceFile:      reschedulerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: reschedulerAddonDestinationFilename,
			},
			expectedNvidia: kubernetesComponentFileSpec{
				sourceFile:      nvidiaAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: nvidiaAddonDestinationFilename,
			},
			expectedContainerMonitoring: kubernetesComponentFileSpec{
				sourceFile:      containerMonitoringAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: containerMonitoringAddonDestinationFilename,
			},
			expectedIPMasqAgent: kubernetesComponentFileSpec{
				sourceFile:      ipMasqAgentAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: ipMasqAgentAddonDestinationFilename,
			},
			expectedAzureCNINetworkMonitor: kubernetesComponentFileSpec{
				sourceFile:      azureCNINetworkMonitorAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureCNINetworkMonitorAddonDestinationFilename,
			},
			expectedDNSAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      dnsAutoscalerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: dnsAutoscalerAddonDestinationFilename,
			},
			expectedCalico: kubernetesComponentFileSpec{
				sourceFile:      calicoAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: calicoAddonDestinationFilename,
			},
			expectedCilium: kubernetesComponentFileSpec{
				sourceFile:      ciliumAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: ciliumAddonDestinationFilename,
			},
			expectedAzureNetworkPolicy: kubernetesComponentFileSpec{
				sourceFile:      azureNetworkPolicyAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azureNetworkPolicyAddonDestinationFilename,
			},
			expectedAzurePolicy: kubernetesComponentFileSpec{
				sourceFile:      azurePolicyAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: azurePolicyAddonDestinationFilename,
			},
			expectedCloudNodeManager: kubernetesComponentFileSpec{
				sourceFile:      cloudNodeManagerAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: cloudNodeManagerAddonDestinationFilename,
			},
			expectedNodeProblemDetector: kubernetesComponentFileSpec{
				sourceFile:      nodeProblemDetectorAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: nodeProblemDetectorAddonDestinationFilename,
			},
			expectedKubeDNS: kubernetesComponentFileSpec{
				sourceFile:      kubeDNSAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: kubeDNSAddonDestinationFilename,
			},
			expectedCoreDNS: kubernetesComponentFileSpec{
				sourceFile:      corednsAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: corednsAddonDestinationFilename,
			},
			expectedKubeProxy: kubernetesComponentFileSpec{
				sourceFile:      kubeProxyAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: kubeProxyAddonDestinationFilename,
			},
			expectedPodSecurityPolicy: kubernetesComponentFileSpec{
				sourceFile:      podSecurityPolicyAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: podSecurityPolicyAddonDestinationFilename,
			},
			expectedAADDefaultAdminGroup: kubernetesComponentFileSpec{
				sourceFile:      aadDefaultAdminGroupAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: aadDefaultAdminGroupDestinationFilename,
			},
			expectedAntrea: kubernetesComponentFileSpec{
				sourceFile:      antreaAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: antreaAddonDestinationFilename,
			},
			expectedAuditPolicy: kubernetesComponentFileSpec{
				sourceFile:      auditPolicyAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: auditPolicyAddonDestinationFilename,
			},
			expectedAzureCloudProvider: kubernetesComponentFileSpec{
				sourceFile:      cloudProviderAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: cloudProviderAddonDestinationFilename,
			},
			expectedFlannel: kubernetesComponentFileSpec{
				sourceFile:      flannelAddonSourceFilename,
				base64Data:      base64Data,
				destinationFile: flannelAddonDestinationFilename,
			},
		},
		{
			name: "addons with no data",
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.16.1",
					KubernetesConfig: &api.KubernetesConfig{
						Addons: []api.KubernetesAddon{
							{
								Name: common.HeapsterAddonName,
							},
							{
								Name: common.MetricsServerAddonName,
							},
							{
								Name: common.TillerAddonName,
							},
							{
								Name: common.AADPodIdentityAddonName,
							},
							{
								Name: common.ACIConnectorAddonName,
							},
							{
								Name: common.AzureDiskCSIDriverAddonName,
							},
							{
								Name: common.AzureFileCSIDriverAddonName,
							},
							{
								Name: common.ClusterAutoscalerAddonName,
							},
							{
								Name: common.BlobfuseFlexVolumeAddonName,
							},
							{
								Name: common.SMBFlexVolumeAddonName,
							},
							{
								Name: common.KeyVaultFlexVolumeAddonName,
							},
							{
								Name: common.DashboardAddonName,
							},
							{
								Name: common.ReschedulerAddonName,
							},
							{
								Name: common.NVIDIADevicePluginAddonName,
							},
							{
								Name: common.ContainerMonitoringAddonName,
							},
							{
								Name: common.IPMASQAgentAddonName,
							},
							{
								Name: common.AzureCNINetworkMonitorAddonName,
							},
							{
								Name: common.DNSAutoscalerAddonName,
							},
							{
								Name: common.CalicoAddonName,
							},
							{
								Name: common.CiliumAddonName,
							},
							{
								Name: common.AzureNetworkPolicyAddonName,
							},
							{
								Name: common.AzurePolicyAddonName,
							},
							{
								Name: common.CloudNodeManagerAddonName,
							},
							{
								Name: common.NodeProblemDetectorAddonName,
							},
							{
								Name: common.KubeDNSAddonName,
							},
							{
								Name: common.CoreDNSAddonName,
							},
							{
								Name: common.KubeProxyAddonName,
							},
							{
								Name: common.PodSecurityPolicyAddonName,
							},
							{
								Name: common.AADAdminGroupAddonName,
							},
							{
								Name: common.AntreaAddonName,
							},
							{
								Name: common.AuditPolicyAddonName,
							},
							{
								Name: common.AzureCloudProviderAddonName,
							},
							{
								Name: common.FlannelAddonName,
							},
						},
					},
				},
			},
			expectedHeapster: kubernetesComponentFileSpec{
				sourceFile:      heapsterAddonSourceFilename,
				base64Data:      "",
				destinationFile: heapsterAddonDestinationFilename,
			},
			expectedMetricsServer: kubernetesComponentFileSpec{
				sourceFile:      metricsServerAddonSourceFilename,
				base64Data:      "",
				destinationFile: metricsServerAddonDestinationFilename,
			},
			expectedTiller: kubernetesComponentFileSpec{
				sourceFile:      tillerAddonSourceFilename,
				base64Data:      "",
				destinationFile: tillerAddonDestinationFilename,
			},
			expectedAADPodIdentity: kubernetesComponentFileSpec{
				sourceFile:      aadPodIdentityAddonSourceFilename,
				base64Data:      "",
				destinationFile: aadPodIdentityAddonDestinationFilename,
			},
			expectedACIConnector: kubernetesComponentFileSpec{
				sourceFile:      aciConnectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: aciConnectorAddonDestinationFilename,
			},
			expectedAzureDiskCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureDiskCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureDiskCSIAddonDestinationFilename,
			},
			expectedAzureFileCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureFileCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureFileCSIAddonDestinationFilename,
			},
			expectedClusterAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      clusterAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: clusterAutoscalerAddonDestinationFilename,
			},
			expectedBlobFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      blobfuseFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: blobfuseFlexVolumeAddonDestinationFilename,
			},
			expectedSMBFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      smbFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: smbFlexVolumeAddonDestinationFilename,
			},
			expectedKeyVaultFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      keyvaultFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: keyvaultFlexVolumeAddonDestinationFilename,
			},
			expectedDashboard: kubernetesComponentFileSpec{
				sourceFile:      dashboardAddonSourceFilename,
				base64Data:      "",
				destinationFile: dashboardAddonDestinationFilename,
			},
			expectedRescheduler: kubernetesComponentFileSpec{
				sourceFile:      reschedulerAddonSourceFilename,
				base64Data:      "",
				destinationFile: reschedulerAddonDestinationFilename,
			},
			expectedNvidia: kubernetesComponentFileSpec{
				sourceFile:      nvidiaAddonSourceFilename,
				base64Data:      "",
				destinationFile: nvidiaAddonDestinationFilename,
			},
			expectedContainerMonitoring: kubernetesComponentFileSpec{
				sourceFile:      containerMonitoringAddonSourceFilename,
				base64Data:      "",
				destinationFile: containerMonitoringAddonDestinationFilename,
			},
			expectedIPMasqAgent: kubernetesComponentFileSpec{
				sourceFile:      ipMasqAgentAddonSourceFilename,
				base64Data:      "",
				destinationFile: ipMasqAgentAddonDestinationFilename,
			},
			expectedAzureCNINetworkMonitor: kubernetesComponentFileSpec{
				sourceFile:      azureCNINetworkMonitorAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureCNINetworkMonitorAddonDestinationFilename,
			},
			expectedDNSAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      dnsAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: dnsAutoscalerAddonDestinationFilename,
			},
			expectedCalico: kubernetesComponentFileSpec{
				sourceFile:      calicoAddonSourceFilename,
				base64Data:      "",
				destinationFile: calicoAddonDestinationFilename,
			},
			expectedCilium: kubernetesComponentFileSpec{
				sourceFile:      ciliumAddonSourceFilename,
				base64Data:      "",
				destinationFile: ciliumAddonDestinationFilename,
			},
			expectedAzureNetworkPolicy: kubernetesComponentFileSpec{
				sourceFile:      azureNetworkPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureNetworkPolicyAddonDestinationFilename,
			},
			expectedAzurePolicy: kubernetesComponentFileSpec{
				sourceFile:      azurePolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azurePolicyAddonDestinationFilename,
			},
			expectedCloudNodeManager: kubernetesComponentFileSpec{
				sourceFile:      cloudNodeManagerAddonSourceFilename,
				base64Data:      "",
				destinationFile: cloudNodeManagerAddonDestinationFilename,
			},
			expectedNodeProblemDetector: kubernetesComponentFileSpec{
				sourceFile:      nodeProblemDetectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: nodeProblemDetectorAddonDestinationFilename,
			},
			expectedKubeDNS: kubernetesComponentFileSpec{
				sourceFile:      kubeDNSAddonSourceFilename,
				base64Data:      "",
				destinationFile: kubeDNSAddonDestinationFilename,
			},
			expectedCoreDNS: kubernetesComponentFileSpec{
				sourceFile:      corednsAddonSourceFilename,
				base64Data:      "",
				destinationFile: corednsAddonDestinationFilename,
			},
			expectedKubeProxy: kubernetesComponentFileSpec{
				sourceFile:      kubeProxyAddonSourceFilename,
				base64Data:      "",
				destinationFile: kubeProxyAddonDestinationFilename,
			},
			expectedPodSecurityPolicy: kubernetesComponentFileSpec{
				sourceFile:      podSecurityPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: podSecurityPolicyAddonDestinationFilename,
			},
			expectedAADDefaultAdminGroup: kubernetesComponentFileSpec{
				sourceFile:      aadDefaultAdminGroupAddonSourceFilename,
				base64Data:      "",
				destinationFile: aadDefaultAdminGroupDestinationFilename,
			},
			expectedAntrea: kubernetesComponentFileSpec{
				sourceFile:      antreaAddonSourceFilename,
				base64Data:      "",
				destinationFile: antreaAddonDestinationFilename,
			},
			expectedAuditPolicy: kubernetesComponentFileSpec{
				sourceFile:      auditPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: auditPolicyAddonDestinationFilename,
			},
			expectedAzureCloudProvider: kubernetesComponentFileSpec{
				sourceFile:      cloudProviderAddonSourceFilename,
				base64Data:      "",
				destinationFile: cloudProviderAddonDestinationFilename,
			},
			expectedFlannel: kubernetesComponentFileSpec{
				sourceFile:      flannelAddonSourceFilename,
				base64Data:      "",
				destinationFile: flannelAddonDestinationFilename,
			},
		},
		{
			name: "no addons in ContainerService object",
			p:    &api.Properties{},
			expectedHeapster: kubernetesComponentFileSpec{
				sourceFile:      heapsterAddonSourceFilename,
				base64Data:      "",
				destinationFile: heapsterAddonDestinationFilename,
			},
			expectedMetricsServer: kubernetesComponentFileSpec{
				sourceFile:      metricsServerAddonSourceFilename,
				base64Data:      "",
				destinationFile: metricsServerAddonDestinationFilename,
			},
			expectedTiller: kubernetesComponentFileSpec{
				sourceFile:      tillerAddonSourceFilename,
				base64Data:      "",
				destinationFile: tillerAddonDestinationFilename,
			},
			expectedAADPodIdentity: kubernetesComponentFileSpec{
				sourceFile:      aadPodIdentityAddonSourceFilename,
				base64Data:      "",
				destinationFile: aadPodIdentityAddonDestinationFilename,
			},
			expectedACIConnector: kubernetesComponentFileSpec{
				sourceFile:      aciConnectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: aciConnectorAddonDestinationFilename,
			},
			expectedAzureDiskCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureDiskCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureDiskCSIAddonDestinationFilename,
			},
			expectedAzureFileCSIDriver: kubernetesComponentFileSpec{
				sourceFile:      azureFileCSIAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureFileCSIAddonDestinationFilename,
			},
			expectedClusterAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      clusterAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: clusterAutoscalerAddonDestinationFilename,
			},
			expectedBlobFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      blobfuseFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: blobfuseFlexVolumeAddonDestinationFilename,
			},
			expectedSMBFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      smbFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: smbFlexVolumeAddonDestinationFilename,
			},
			expectedKeyVaultFlexVolume: kubernetesComponentFileSpec{
				sourceFile:      keyvaultFlexVolumeAddonSourceFilename,
				base64Data:      "",
				destinationFile: keyvaultFlexVolumeAddonDestinationFilename,
			},
			expectedDashboard: kubernetesComponentFileSpec{
				sourceFile:      dashboardAddonSourceFilename,
				base64Data:      "",
				destinationFile: dashboardAddonDestinationFilename,
			},
			expectedRescheduler: kubernetesComponentFileSpec{
				sourceFile:      reschedulerAddonSourceFilename,
				base64Data:      "",
				destinationFile: reschedulerAddonDestinationFilename,
			},
			expectedNvidia: kubernetesComponentFileSpec{
				sourceFile:      nvidiaAddonSourceFilename,
				base64Data:      "",
				destinationFile: nvidiaAddonDestinationFilename,
			},
			expectedContainerMonitoring: kubernetesComponentFileSpec{
				sourceFile:      containerMonitoringAddonSourceFilename,
				base64Data:      "",
				destinationFile: containerMonitoringAddonDestinationFilename,
			},
			expectedIPMasqAgent: kubernetesComponentFileSpec{
				sourceFile:      ipMasqAgentAddonSourceFilename,
				base64Data:      "",
				destinationFile: ipMasqAgentAddonDestinationFilename,
			},
			expectedAzureCNINetworkMonitor: kubernetesComponentFileSpec{
				sourceFile:      azureCNINetworkMonitorAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureCNINetworkMonitorAddonDestinationFilename,
			},
			expectedDNSAutoscaler: kubernetesComponentFileSpec{
				sourceFile:      dnsAutoscalerAddonSourceFilename,
				base64Data:      "",
				destinationFile: dnsAutoscalerAddonDestinationFilename,
			},
			expectedCalico: kubernetesComponentFileSpec{
				sourceFile:      calicoAddonSourceFilename,
				base64Data:      "",
				destinationFile: calicoAddonDestinationFilename,
			},
			expectedCilium: kubernetesComponentFileSpec{
				sourceFile:      ciliumAddonSourceFilename,
				base64Data:      "",
				destinationFile: ciliumAddonDestinationFilename,
			},
			expectedAzureNetworkPolicy: kubernetesComponentFileSpec{
				sourceFile:      azureNetworkPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azureNetworkPolicyAddonDestinationFilename,
			},
			expectedAzurePolicy: kubernetesComponentFileSpec{
				sourceFile:      azurePolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: azurePolicyAddonDestinationFilename,
			},
			expectedCloudNodeManager: kubernetesComponentFileSpec{
				sourceFile:      cloudNodeManagerAddonSourceFilename,
				base64Data:      "",
				destinationFile: cloudNodeManagerAddonDestinationFilename,
			},
			expectedNodeProblemDetector: kubernetesComponentFileSpec{
				sourceFile:      nodeProblemDetectorAddonSourceFilename,
				base64Data:      "",
				destinationFile: nodeProblemDetectorAddonDestinationFilename,
			},
			expectedKubeDNS: kubernetesComponentFileSpec{
				sourceFile:      kubeDNSAddonSourceFilename,
				base64Data:      "",
				destinationFile: kubeDNSAddonDestinationFilename,
			},
			expectedCoreDNS: kubernetesComponentFileSpec{
				sourceFile:      corednsAddonSourceFilename,
				base64Data:      "",
				destinationFile: corednsAddonDestinationFilename,
			},
			expectedKubeProxy: kubernetesComponentFileSpec{
				sourceFile:      kubeProxyAddonSourceFilename,
				base64Data:      "",
				destinationFile: kubeProxyAddonDestinationFilename,
			},
			expectedPodSecurityPolicy: kubernetesComponentFileSpec{
				sourceFile:      podSecurityPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: podSecurityPolicyAddonDestinationFilename,
			},
			expectedAADDefaultAdminGroup: kubernetesComponentFileSpec{
				sourceFile:      aadDefaultAdminGroupAddonSourceFilename,
				base64Data:      "",
				destinationFile: aadDefaultAdminGroupDestinationFilename,
			},
			expectedAntrea: kubernetesComponentFileSpec{
				sourceFile:      antreaAddonSourceFilename,
				base64Data:      "",
				destinationFile: antreaAddonDestinationFilename,
			},
			expectedAuditPolicy: kubernetesComponentFileSpec{
				sourceFile:      auditPolicyAddonSourceFilename,
				base64Data:      "",
				destinationFile: auditPolicyAddonDestinationFilename,
			},
			expectedAzureCloudProvider: kubernetesComponentFileSpec{
				sourceFile:      cloudProviderAddonSourceFilename,
				base64Data:      "",
				destinationFile: cloudProviderAddonDestinationFilename,
			},
			expectedFlannel: kubernetesComponentFileSpec{
				sourceFile:      flannelAddonSourceFilename,
				base64Data:      "",
				destinationFile: flannelAddonDestinationFilename,
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			componentFileSpec := kubernetesAddonSettingsInit(c.p)
			for addon := range componentFileSpec {
				switch addon {
				case common.HeapsterAddonName:
					if c.expectedHeapster.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedHeapster.sourceFile)
					}
					if c.expectedHeapster.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedHeapster.base64Data)
					}
					if c.expectedHeapster.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedHeapster.destinationFile)
					}
				case common.MetricsServerAddonName:
					if c.expectedMetricsServer.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedMetricsServer.sourceFile)
					}
					if c.expectedMetricsServer.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedMetricsServer.base64Data)
					}
					if c.expectedMetricsServer.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedMetricsServer.destinationFile)
					}
				case common.TillerAddonName:
					if c.expectedTiller.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedTiller.sourceFile)
					}
					if c.expectedTiller.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedTiller.base64Data)
					}
					if c.expectedTiller.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedTiller.destinationFile)
					}
				case common.AADPodIdentityAddonName:
					if c.expectedAADPodIdentity.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAADPodIdentity.sourceFile)
					}
					if c.expectedAADPodIdentity.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAADPodIdentity.base64Data)
					}
					if c.expectedAADPodIdentity.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAADPodIdentity.destinationFile)
					}
				case common.ACIConnectorAddonName:
					if c.expectedACIConnector.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedACIConnector.sourceFile)
					}
					if c.expectedACIConnector.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedACIConnector.base64Data)
					}
					if c.expectedACIConnector.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedACIConnector.destinationFile)
					}
				case common.AzureDiskCSIDriverAddonName:
					if c.expectedAzureDiskCSIDriver.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureDiskCSIDriver.sourceFile)
					}
					if c.expectedAzureDiskCSIDriver.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureDiskCSIDriver.base64Data)
					}
					if c.expectedAzureDiskCSIDriver.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureDiskCSIDriver.destinationFile)
					}
				case common.AzureFileCSIDriverAddonName:
					if c.expectedAzureFileCSIDriver.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureFileCSIDriver.sourceFile)
					}
					if c.expectedAzureFileCSIDriver.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureFileCSIDriver.base64Data)
					}
					if c.expectedAzureFileCSIDriver.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureFileCSIDriver.destinationFile)
					}
				case common.ClusterAutoscalerAddonName:
					if c.expectedClusterAutoscaler.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedClusterAutoscaler.sourceFile)
					}
					if c.expectedClusterAutoscaler.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedClusterAutoscaler.base64Data)
					}
					if c.expectedClusterAutoscaler.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedClusterAutoscaler.destinationFile)
					}
				case common.BlobfuseFlexVolumeAddonName:
					if c.expectedBlobFlexVolume.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedBlobFlexVolume.sourceFile)
					}
					if c.expectedBlobFlexVolume.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedBlobFlexVolume.base64Data)
					}
					if c.expectedBlobFlexVolume.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedBlobFlexVolume.destinationFile)
					}
				case common.SMBFlexVolumeAddonName:
					if c.expectedSMBFlexVolume.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedSMBFlexVolume.sourceFile)
					}
					if c.expectedSMBFlexVolume.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedSMBFlexVolume.base64Data)
					}
					if c.expectedSMBFlexVolume.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedSMBFlexVolume.destinationFile)
					}
				case common.KeyVaultFlexVolumeAddonName:
					if c.expectedKeyVaultFlexVolume.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedKeyVaultFlexVolume.sourceFile)
					}
					if c.expectedKeyVaultFlexVolume.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedKeyVaultFlexVolume.base64Data)
					}
					if c.expectedKeyVaultFlexVolume.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedKeyVaultFlexVolume.destinationFile)
					}
				case common.DashboardAddonName:
					if c.expectedDashboard.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedDashboard.sourceFile)
					}
					if c.expectedDashboard.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedDashboard.base64Data)
					}
					if c.expectedDashboard.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedDashboard.destinationFile)
					}
				case common.ReschedulerAddonName:
					if c.expectedRescheduler.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedRescheduler.sourceFile)
					}
					if c.expectedRescheduler.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedRescheduler.base64Data)
					}
					if c.expectedRescheduler.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedRescheduler.destinationFile)
					}
				case common.NVIDIADevicePluginAddonName:
					if c.expectedNvidia.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedNvidia.sourceFile)
					}
					if c.expectedNvidia.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedNvidia.base64Data)
					}
					if c.expectedNvidia.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedNvidia.destinationFile)
					}
				case common.ContainerMonitoringAddonName:
					if c.expectedContainerMonitoring.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedContainerMonitoring.sourceFile)
					}
					if c.expectedContainerMonitoring.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedContainerMonitoring.base64Data)
					}
					if c.expectedContainerMonitoring.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedContainerMonitoring.destinationFile)
					}
				case common.IPMASQAgentAddonName:
					if c.expectedIPMasqAgent.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedIPMasqAgent.sourceFile)
					}
					if c.expectedIPMasqAgent.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedIPMasqAgent.base64Data)
					}
					if c.expectedIPMasqAgent.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedIPMasqAgent.destinationFile)
					}
				case common.AzureCNINetworkMonitorAddonName:
					if c.expectedAzureCNINetworkMonitor.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureCNINetworkMonitor.sourceFile)
					}
					if c.expectedAzureCNINetworkMonitor.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureCNINetworkMonitor.base64Data)
					}
					if c.expectedAzureCNINetworkMonitor.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureCNINetworkMonitor.destinationFile)
					}
				case common.DNSAutoscalerAddonName:
					if c.expectedDNSAutoscaler.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedDNSAutoscaler.sourceFile)
					}
					if c.expectedDNSAutoscaler.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedDNSAutoscaler.base64Data)
					}
					if c.expectedDNSAutoscaler.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedDNSAutoscaler.destinationFile)
					}
				case common.CalicoAddonName:
					if c.expectedCalico.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedCalico.sourceFile)
					}
					if c.expectedCalico.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedCalico.base64Data)
					}
					if c.expectedCalico.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedCalico.destinationFile)
					}
				case common.CiliumAddonName:
					if c.expectedCilium.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedCilium.sourceFile)
					}
					if c.expectedCilium.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedCilium.base64Data)
					}
					if c.expectedCilium.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedCilium.destinationFile)
					}
				case common.AzureNetworkPolicyAddonName:
					if c.expectedAzureNetworkPolicy.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureNetworkPolicy.sourceFile)
					}
					if c.expectedAzureNetworkPolicy.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureNetworkPolicy.base64Data)
					}
					if c.expectedAzureNetworkPolicy.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureNetworkPolicy.destinationFile)
					}
				case common.AzurePolicyAddonName:
					if c.expectedAzurePolicy.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzurePolicy.sourceFile)
					}
					if c.expectedAzurePolicy.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzurePolicy.base64Data)
					}
					if c.expectedAzurePolicy.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzurePolicy.destinationFile)
					}
				case common.CloudNodeManagerAddonName:
					if c.expectedCloudNodeManager.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedCloudNodeManager.sourceFile)
					}
					if c.expectedCloudNodeManager.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedCloudNodeManager.base64Data)
					}
					if c.expectedCloudNodeManager.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedCloudNodeManager.destinationFile)
					}
				case common.NodeProblemDetectorAddonName:
					if c.expectedNodeProblemDetector.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedNodeProblemDetector.sourceFile)
					}
					if c.expectedNodeProblemDetector.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedNodeProblemDetector.base64Data)
					}
					if c.expectedNodeProblemDetector.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedNodeProblemDetector.destinationFile)
					}
				case common.KubeDNSAddonName:
					if c.expectedKubeDNS.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedKubeDNS.sourceFile)
					}
					if c.expectedKubeDNS.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedKubeDNS.base64Data)
					}
					if c.expectedKubeDNS.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedKubeDNS.destinationFile)
					}
				case common.CoreDNSAddonName:
					if c.expectedCoreDNS.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedCoreDNS.sourceFile)
					}
					if c.expectedCoreDNS.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedCoreDNS.base64Data)
					}
					if c.expectedCoreDNS.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedCoreDNS.destinationFile)
					}
				case common.KubeProxyAddonName:
					if c.expectedKubeProxy.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedKubeProxy.sourceFile)
					}
					if c.expectedKubeProxy.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedKubeProxy.base64Data)
					}
					if c.expectedKubeProxy.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedKubeProxy.destinationFile)
					}
				case common.PodSecurityPolicyAddonName:
					if c.expectedPodSecurityPolicy.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedPodSecurityPolicy.sourceFile)
					}
					if c.expectedPodSecurityPolicy.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedPodSecurityPolicy.base64Data)
					}
					if c.expectedPodSecurityPolicy.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedPodSecurityPolicy.destinationFile)
					}
				case common.AADAdminGroupAddonName:
					if c.expectedAADDefaultAdminGroup.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAADDefaultAdminGroup.sourceFile)
					}
					if c.expectedAADDefaultAdminGroup.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAADDefaultAdminGroup.base64Data)
					}
					if c.expectedAADDefaultAdminGroup.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAADDefaultAdminGroup.destinationFile)
					}
				case common.AntreaAddonName:
					if c.expectedAntrea.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAntrea.sourceFile)
					}
					if c.expectedAntrea.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAntrea.base64Data)
					}
					if c.expectedAntrea.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAntrea.destinationFile)
					}
				case common.AuditPolicyAddonName:
					if c.expectedAuditPolicy.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAuditPolicy.sourceFile)
					}
					if c.expectedAuditPolicy.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAuditPolicy.base64Data)
					}
					if c.expectedAuditPolicy.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAuditPolicy.destinationFile)
					}
				case common.AzureCloudProviderAddonName:
					if c.expectedAzureCloudProvider.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedAzureCloudProvider.sourceFile)
					}
					if c.expectedAzureCloudProvider.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedAzureCloudProvider.base64Data)
					}
					if c.expectedAzureCloudProvider.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedAzureCloudProvider.destinationFile)
					}
				case common.FlannelAddonName:
					if c.expectedFlannel.sourceFile != componentFileSpec[addon].sourceFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].sourceFile, c.expectedFlannel.sourceFile)
					}
					if c.expectedFlannel.base64Data != componentFileSpec[addon].base64Data {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].base64Data, c.expectedFlannel.base64Data)
					}
					if c.expectedFlannel.destinationFile != componentFileSpec[addon].destinationFile {
						t.Fatalf("Expected %s to be %s", componentFileSpec[addon].destinationFile, c.expectedFlannel.destinationFile)
					}
				}
			}
		})
	}
}

func TestKubernetesManifestSettingsInit(t *testing.T) {
	mockAzureStackProperties := api.GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	cases := []struct {
		p                             *api.Properties
		expectedKubeScheduler         bool
		expectedKubeControllerManager bool
		expectedKubeCCM               bool
		expectedKubeAPIServer         bool
		expectedKubeAddonManager      bool
	}{
		// Default scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig: map[string]string{},
					},
				},
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               false,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
		// CCM scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig:              map[string]string{},
						UseCloudControllerManager:    to.BoolPtr(true),
						CloudControllerManagerConfig: map[string]string{},
					},
				},
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               true,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
		// Azure Stack Scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig: map[string]string{},
					},
				},
				CustomCloudProfile: mockAzureStackProperties.CustomCloudProfile,
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               false,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
		// Custom data scenario
		{
			p: &api.Properties{
				OrchestratorProfile: &api.OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: "1.14.1",
					KubernetesConfig: &api.KubernetesConfig{
						SchedulerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("foo")),
						},
						ControllerManagerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("bar")),
						},
						CloudControllerManagerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("baz")),
						},
						APIServerConfig: map[string]string{
							"data": base64.StdEncoding.EncodeToString([]byte("bam")),
						},
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
			expectedKubeScheduler:         true,
			expectedKubeControllerManager: true,
			expectedKubeCCM:               true,
			expectedKubeAPIServer:         true,
			expectedKubeAddonManager:      true,
		},
	}
	for _, c := range cases {
		componentFileSpecArray := kubernetesManifestSettingsInit(c.p)
		for _, componentFileSpec := range componentFileSpecArray {
			switch componentFileSpec.destinationFile {
			case "kube-scheduler.yaml":
				if c.expectedKubeScheduler != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-scheduler", c.expectedKubeScheduler)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.SchedulerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.SchedulerConfig["data"])
				}
			case "kube-controller-manager.yaml":
				if c.p.CustomCloudProfile != nil {
					if !strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
						t.Fatalf("Expected an Azure Stack-specific controller-manager spec, got %s instead", componentFileSpec.sourceFile)
					}
				} else {
					if strings.Contains(componentFileSpec.sourceFile, "-custom.yaml") {
						t.Fatalf("Got an unexpected Azure Stack-specific controller-manager spec in a non-Azure Stack cluster configuration")
					}
				}
				if c.expectedKubeControllerManager != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-controller-manager", c.expectedKubeControllerManager)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig["data"])
				}
			case "cloud-controller-manager.yaml":
				if c.expectedKubeCCM != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "cloud-controller-manager", c.expectedKubeCCM)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig["data"])
				}
			case "kube-apiserver.yaml":
				if c.expectedKubeAPIServer != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-apiserver", c.expectedKubeAPIServer)
				}
				if c.p.OrchestratorProfile.KubernetesConfig.APIServerConfig["data"] != componentFileSpec.base64Data {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, c.p.OrchestratorProfile.KubernetesConfig.APIServerConfig["data"])
				}
			case "kube-addon-manager.yaml":
				if c.expectedKubeAddonManager != componentFileSpec.isEnabled {
					t.Fatalf("Expected %s to be %t", "kube-addon-manager", c.expectedKubeAddonManager)
				}
				if componentFileSpec.base64Data != "" {
					t.Fatalf("Expected %s to be %s", componentFileSpec.base64Data, "")
				}
			}
		}
	}
}
