// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
)

func TestSetKubernetesImagesConfig(t *testing.T) {
	azureStackKubernetesSpecConfig := KubernetesSpecConfig{
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
	}
	azureStackCloudSpec := AzureEnvironmentSpecConfig{
		CloudName:            "AzurePublicCloud",
		KubernetesSpecConfig: azureStackKubernetesSpecConfig,
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
		},
	}
	AzureCloudSpecEnvMap[AzureStackCloud] = azureStackCloudSpec
	cases := []struct {
		name                    string
		cs                      *ContainerService
		expectedImageBaseConfig ImageBaseConfig
		expectedImageConfig     map[string]string
	}{
		{
			name: "default",
			cs: &ContainerService{
				Location:   "westus2",
				Properties: &Properties{},
			},
			expectedImageBaseConfig: ImageBaseConfig{
				KubernetesImageBase:                    DefaultExternalContainerImageRegistry,
				TillerImageBase:                        "gcr.io/kubernetes-helm/",
				ACIConnectorImageBase:                  "microsoft/",
				NVIDIAImageBase:                        "nvidia/",
				CalicoImageBase:                        "calico/",
				AzureCNIImageBase:                      "mcr.microsoft.com/containernetworking/",
				HyperkubeImageBase:                     DefaultExternalContainerImageRegistry,
				PauseImageBase:                         DefaultExternalContainerImageRegistry,
				AddonManagerImageBase:                  DefaultExternalContainerImageRegistry,
				CloudControllerManagerImageBase:        DefaultExternalContainerImageRegistry,
				K8sDNSSidecarImageBase:                 DefaultExternalContainerImageRegistry,
				CoreDNSImageBase:                       DefaultExternalContainerImageRegistry,
				KubeDNSImageBase:                       DefaultExternalContainerImageRegistry,
				DNSMasqImageBase:                       DefaultExternalContainerImageRegistry,
				HeapsterImageBase:                      DefaultExternalContainerImageRegistry,
				AddonResizerImageBase:                  DefaultExternalContainerImageRegistry,
				ClusterAutoscalerImageBase:             DefaultExternalContainerImageRegistry,
				DashboardImageBase:                     DefaultExternalContainerImageRegistry,
				ReschedulerImageBase:                   DefaultExternalContainerImageRegistry,
				MetricsServerImageBase:                 DefaultExternalContainerImageRegistry,
				IPMasqAgentImageBase:                   DefaultExternalContainerImageRegistry,
				ClusterProportionalAutoscalerImageBase: DefaultExternalContainerImageRegistry,
			},
			expectedImageConfig: map[string]string{},
		},
		{
			name: "KubernetesConfig.KubernetesImageBase override case",
			cs: &ContainerService{
				Location: "westus2",
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBase: "custom",
						},
					},
				},
			},
			expectedImageBaseConfig: ImageBaseConfig{
				KubernetesImageBase:                    "custom",
				TillerImageBase:                        "gcr.io/kubernetes-helm/",
				ACIConnectorImageBase:                  "microsoft/",
				NVIDIAImageBase:                        "nvidia/",
				CalicoImageBase:                        "calico/",
				AzureCNIImageBase:                      "mcr.microsoft.com/containernetworking/",
				HyperkubeImageBase:                     "custom",
				PauseImageBase:                         "custom",
				AddonManagerImageBase:                  "custom",
				CloudControllerManagerImageBase:        "custom",
				K8sDNSSidecarImageBase:                 "custom",
				CoreDNSImageBase:                       "custom",
				KubeDNSImageBase:                       "custom",
				DNSMasqImageBase:                       "custom",
				HeapsterImageBase:                      "custom",
				AddonResizerImageBase:                  "custom",
				ClusterAutoscalerImageBase:             "custom",
				DashboardImageBase:                     "custom",
				ReschedulerImageBase:                   "custom",
				MetricsServerImageBase:                 "custom",
				IPMasqAgentImageBase:                   "custom",
				ClusterProportionalAutoscalerImageBase: "custom",
			},
			expectedImageConfig: map[string]string{},
		},
		{
			name: "ImageBaseConfig.KubernetesImageBase overrides case",
			cs: &ContainerService{
				Location: "westus2",
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						KubernetesConfig: &KubernetesConfig{
							KubernetesImagesConfig: &KubernetesImagesConfig{
								ImageBaseConfig: &ImageBaseConfig{
									KubernetesImageBase:                    "customKubernetesImageBase",
									TillerImageBase:                        "customTillerImageBase",
									ACIConnectorImageBase:                  "customACIConnectorImageBase",
									NVIDIAImageBase:                        "customNVIDIAImageBase",
									CalicoImageBase:                        "customCalicoImageBase",
									AzureCNIImageBase:                      "customAzureCNIImageBase",
									HyperkubeImageBase:                     "customHyperkubeImageBase",
									PauseImageBase:                         "customPauseImageBase",
									AddonManagerImageBase:                  "customAddonManagerImageBase",
									CloudControllerManagerImageBase:        "customCloudControllerManagerImageBase",
									K8sDNSSidecarImageBase:                 "customK8sDNSSidecarImageBase",
									CoreDNSImageBase:                       "customCoreDNSImageBase",
									KubeDNSImageBase:                       "customKubeDNSImageBase",
									DNSMasqImageBase:                       "customDNSMasqImageBase",
									HeapsterImageBase:                      "customHeapsterImageBase",
									AddonResizerImageBase:                  "customAddonResizerImageBase",
									ClusterAutoscalerImageBase:             "customClusterAutoscalerImageBase",
									DashboardImageBase:                     "customDashboardImageBase",
									ReschedulerImageBase:                   "customReschedulerImageBase",
									MetricsServerImageBase:                 "customMetricsServerImageBase",
									IPMasqAgentImageBase:                   "customIPMasqAgentImageBase",
									ClusterProportionalAutoscalerImageBase: "customClusterProportionalAutoscalerImageBase",
								},
							},
						},
					},
				},
			},
			expectedImageBaseConfig: ImageBaseConfig{
				KubernetesImageBase:                    "customKubernetesImageBase",
				TillerImageBase:                        "customTillerImageBase",
				ACIConnectorImageBase:                  "customACIConnectorImageBase",
				NVIDIAImageBase:                        "customNVIDIAImageBase",
				CalicoImageBase:                        "customCalicoImageBase",
				AzureCNIImageBase:                      "customAzureCNIImageBase",
				HyperkubeImageBase:                     "customHyperkubeImageBase",
				PauseImageBase:                         "customPauseImageBase",
				AddonManagerImageBase:                  "customAddonManagerImageBase",
				CloudControllerManagerImageBase:        "customCloudControllerManagerImageBase",
				K8sDNSSidecarImageBase:                 "customK8sDNSSidecarImageBase",
				CoreDNSImageBase:                       "customCoreDNSImageBase",
				KubeDNSImageBase:                       "customKubeDNSImageBase",
				DNSMasqImageBase:                       "customDNSMasqImageBase",
				HeapsterImageBase:                      "customHeapsterImageBase",
				AddonResizerImageBase:                  "customAddonResizerImageBase",
				ClusterAutoscalerImageBase:             "customClusterAutoscalerImageBase",
				DashboardImageBase:                     "customDashboardImageBase",
				ReschedulerImageBase:                   "customReschedulerImageBase",
				MetricsServerImageBase:                 "customMetricsServerImageBase",
				IPMasqAgentImageBase:                   "customIPMasqAgentImageBase",
				ClusterProportionalAutoscalerImageBase: "customClusterProportionalAutoscalerImageBase",
			},
			expectedImageConfig: nil,
		},
		{
			name: "ImageBaseConfig.KubernetesImageBase partial overrides case",
			cs: &ContainerService{
				Location: "westus2",
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						KubernetesConfig: &KubernetesConfig{
							KubernetesImagesConfig: &KubernetesImagesConfig{
								ImageBaseConfig: &ImageBaseConfig{
									KubernetesImageBase: "customKubernetesImageBase",
								},
							},
						},
					},
				},
			},
			expectedImageBaseConfig: ImageBaseConfig{
				KubernetesImageBase:                    "customKubernetesImageBase",
				TillerImageBase:                        "gcr.io/kubernetes-helm/",
				ACIConnectorImageBase:                  "microsoft/",
				NVIDIAImageBase:                        "nvidia/",
				CalicoImageBase:                        "calico/",
				AzureCNIImageBase:                      "mcr.microsoft.com/containernetworking/",
				HyperkubeImageBase:                     DefaultExternalContainerImageRegistry,
				PauseImageBase:                         DefaultExternalContainerImageRegistry,
				AddonManagerImageBase:                  DefaultExternalContainerImageRegistry,
				CloudControllerManagerImageBase:        DefaultExternalContainerImageRegistry,
				K8sDNSSidecarImageBase:                 DefaultExternalContainerImageRegistry,
				CoreDNSImageBase:                       DefaultExternalContainerImageRegistry,
				KubeDNSImageBase:                       DefaultExternalContainerImageRegistry,
				DNSMasqImageBase:                       DefaultExternalContainerImageRegistry,
				HeapsterImageBase:                      DefaultExternalContainerImageRegistry,
				AddonResizerImageBase:                  DefaultExternalContainerImageRegistry,
				ClusterAutoscalerImageBase:             DefaultExternalContainerImageRegistry,
				DashboardImageBase:                     DefaultExternalContainerImageRegistry,
				ReschedulerImageBase:                   DefaultExternalContainerImageRegistry,
				MetricsServerImageBase:                 DefaultExternalContainerImageRegistry,
				IPMasqAgentImageBase:                   DefaultExternalContainerImageRegistry,
				ClusterProportionalAutoscalerImageBase: DefaultExternalContainerImageRegistry,
			},
			expectedImageConfig: nil,
		},
		{
			name: "ImageBaseConfig.KubernetesImageBase partial overrides w/ ImageBaseConfig.KubernetesImageBase case",
			cs: &ContainerService{
				Location: "westus2",
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBase: "custom",
							KubernetesImagesConfig: &KubernetesImagesConfig{
								ImageBaseConfig: &ImageBaseConfig{
									DNSMasqImageBase: "customDNSMasqImageBase",
								},
							},
						},
					},
				},
			},
			expectedImageBaseConfig: ImageBaseConfig{
				KubernetesImageBase:                    "custom",
				TillerImageBase:                        "gcr.io/kubernetes-helm/",
				ACIConnectorImageBase:                  "microsoft/",
				NVIDIAImageBase:                        "nvidia/",
				CalicoImageBase:                        "calico/",
				AzureCNIImageBase:                      "mcr.microsoft.com/containernetworking/",
				HyperkubeImageBase:                     "custom",
				PauseImageBase:                         "custom",
				AddonManagerImageBase:                  "custom",
				CloudControllerManagerImageBase:        "custom",
				K8sDNSSidecarImageBase:                 "custom",
				CoreDNSImageBase:                       "custom",
				KubeDNSImageBase:                       "custom",
				DNSMasqImageBase:                       "customDNSMasqImageBase",
				HeapsterImageBase:                      "custom",
				AddonResizerImageBase:                  "custom",
				ClusterAutoscalerImageBase:             "custom",
				DashboardImageBase:                     "custom",
				ReschedulerImageBase:                   "custom",
				MetricsServerImageBase:                 "custom",
				IPMasqAgentImageBase:                   "custom",
				ClusterProportionalAutoscalerImageBase: "custom",
			},
			expectedImageConfig: nil,
		},
		{
			name: "ImageConfig values case",
			cs: &ContainerService{
				Location: "westus2",
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						KubernetesConfig: &KubernetesConfig{
							KubernetesImageBase: "custom",
							KubernetesImagesConfig: &KubernetesImagesConfig{
								ImageConfig: map[string]string{
									"foo": "bar",
								},
							},
						},
					},
				},
			},
			expectedImageBaseConfig: ImageBaseConfig{
				KubernetesImageBase:                    DefaultExternalContainerImageRegistry,
				TillerImageBase:                        "gcr.io/kubernetes-helm/",
				ACIConnectorImageBase:                  "microsoft/",
				NVIDIAImageBase:                        "nvidia/",
				CalicoImageBase:                        "calico/",
				AzureCNIImageBase:                      "mcr.microsoft.com/containernetworking/",
				HyperkubeImageBase:                     DefaultExternalContainerImageRegistry,
				PauseImageBase:                         DefaultExternalContainerImageRegistry,
				AddonManagerImageBase:                  DefaultExternalContainerImageRegistry,
				CloudControllerManagerImageBase:        DefaultExternalContainerImageRegistry,
				K8sDNSSidecarImageBase:                 DefaultExternalContainerImageRegistry,
				CoreDNSImageBase:                       DefaultExternalContainerImageRegistry,
				KubeDNSImageBase:                       DefaultExternalContainerImageRegistry,
				DNSMasqImageBase:                       DefaultExternalContainerImageRegistry,
				HeapsterImageBase:                      DefaultExternalContainerImageRegistry,
				AddonResizerImageBase:                  DefaultExternalContainerImageRegistry,
				ClusterAutoscalerImageBase:             DefaultExternalContainerImageRegistry,
				DashboardImageBase:                     DefaultExternalContainerImageRegistry,
				ReschedulerImageBase:                   DefaultExternalContainerImageRegistry,
				MetricsServerImageBase:                 DefaultExternalContainerImageRegistry,
				IPMasqAgentImageBase:                   DefaultExternalContainerImageRegistry,
				ClusterProportionalAutoscalerImageBase: DefaultExternalContainerImageRegistry,
			},
			expectedImageConfig: map[string]string{
				"foo": "bar",
			},
		},
		{
			name: "Azure Stack case",
			cs: &ContainerService{
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						Environment: &azure.Environment{
							Name: "AzureStackCloud",
						},
					},
				},
			},
			expectedImageBaseConfig: ImageBaseConfig{
				KubernetesImageBase:                    azureStackKubernetesSpecConfig.KubernetesImageBase,
				TillerImageBase:                        azureStackKubernetesSpecConfig.TillerImageBase,
				ACIConnectorImageBase:                  azureStackKubernetesSpecConfig.ACIConnectorImageBase,
				NVIDIAImageBase:                        azureStackKubernetesSpecConfig.NVIDIAImageBase,
				CalicoImageBase:                        azureStackKubernetesSpecConfig.CalicoImageBase,
				AzureCNIImageBase:                      azureStackKubernetesSpecConfig.AzureCNIImageBase,
				HyperkubeImageBase:                     azureStackKubernetesSpecConfig.KubernetesImageBase,
				PauseImageBase:                         azureStackKubernetesSpecConfig.KubernetesImageBase,
				AddonManagerImageBase:                  azureStackKubernetesSpecConfig.KubernetesImageBase,
				CloudControllerManagerImageBase:        azureStackKubernetesSpecConfig.KubernetesImageBase,
				K8sDNSSidecarImageBase:                 azureStackKubernetesSpecConfig.KubernetesImageBase,
				CoreDNSImageBase:                       azureStackKubernetesSpecConfig.KubernetesImageBase,
				KubeDNSImageBase:                       azureStackKubernetesSpecConfig.KubernetesImageBase,
				DNSMasqImageBase:                       azureStackKubernetesSpecConfig.KubernetesImageBase,
				HeapsterImageBase:                      azureStackKubernetesSpecConfig.KubernetesImageBase,
				AddonResizerImageBase:                  azureStackKubernetesSpecConfig.KubernetesImageBase,
				ClusterAutoscalerImageBase:             azureStackKubernetesSpecConfig.KubernetesImageBase,
				DashboardImageBase:                     azureStackKubernetesSpecConfig.KubernetesImageBase,
				ReschedulerImageBase:                   azureStackKubernetesSpecConfig.KubernetesImageBase,
				MetricsServerImageBase:                 azureStackKubernetesSpecConfig.KubernetesImageBase,
				IPMasqAgentImageBase:                   azureStackKubernetesSpecConfig.KubernetesImageBase,
				ClusterProportionalAutoscalerImageBase: azureStackKubernetesSpecConfig.KubernetesImageBase,
			},
			expectedImageConfig: map[string]string{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.cs.setKubernetesImagesConfig()
			imageBaseConfig := *c.cs.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImagesConfig.ImageBaseConfig
			if !reflect.DeepEqual(imageBaseConfig, c.expectedImageBaseConfig) {
				t.Fatalf("expected result ImageBaseConfig %v to be equal to %v", imageBaseConfig, c.expectedImageBaseConfig)
			}
			imageConfig := c.cs.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImagesConfig.ImageConfig
			if !reflect.DeepEqual(imageConfig, c.expectedImageConfig) {
				t.Fatalf("expected result ImageBaseConfig %v to be equal to %v", imageConfig, c.expectedImageConfig)
			}
		})
	}

}

func TestGetImageBaseConfigFromKubernetesSpecConfig(t *testing.T) {
	cases := []struct {
		name                    string
		specConfig              KubernetesSpecConfig
		expectedImageBaseConfig *ImageBaseConfig
	}{
		{
			name: "default",
			specConfig: KubernetesSpecConfig{
				KubernetesImageBase:   "KubernetesImageBase",
				TillerImageBase:       "TillerImageBase",
				ACIConnectorImageBase: "ACIConnectorImageBase",
				NVIDIAImageBase:       "NVIDIAImageBase",
				AzureCNIImageBase:     "AzureCNIImageBase",
				CalicoImageBase:       "CalicoImageBase",
			},
			expectedImageBaseConfig: &ImageBaseConfig{
				KubernetesImageBase:                    "KubernetesImageBase",
				TillerImageBase:                        "TillerImageBase",
				ACIConnectorImageBase:                  "ACIConnectorImageBase",
				NVIDIAImageBase:                        "NVIDIAImageBase",
				CalicoImageBase:                        "CalicoImageBase",
				AzureCNIImageBase:                      "AzureCNIImageBase",
				HyperkubeImageBase:                     "KubernetesImageBase",
				PauseImageBase:                         "KubernetesImageBase",
				AddonManagerImageBase:                  "KubernetesImageBase",
				CloudControllerManagerImageBase:        "KubernetesImageBase",
				K8sDNSSidecarImageBase:                 "KubernetesImageBase",
				CoreDNSImageBase:                       "KubernetesImageBase",
				KubeDNSImageBase:                       "KubernetesImageBase",
				DNSMasqImageBase:                       "KubernetesImageBase",
				HeapsterImageBase:                      "KubernetesImageBase",
				AddonResizerImageBase:                  "KubernetesImageBase",
				ClusterAutoscalerImageBase:             "KubernetesImageBase",
				DashboardImageBase:                     "KubernetesImageBase",
				ReschedulerImageBase:                   "KubernetesImageBase",
				MetricsServerImageBase:                 "KubernetesImageBase",
				IPMasqAgentImageBase:                   "KubernetesImageBase",
				ClusterProportionalAutoscalerImageBase: "KubernetesImageBase",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			imageBaseConfig := GetImageBaseConfigFromKubernetesSpecConfig(c.specConfig)
			if !reflect.DeepEqual(imageBaseConfig, c.expectedImageBaseConfig) {
				t.Fatalf("expected result ImageBaseConfig %v to be equal to %v", imageBaseConfig, c.expectedImageBaseConfig)
			}
		})
	}

}
