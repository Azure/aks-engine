// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

func (cs *ContainerService) setKubernetesImagesConfig() {
	cloudSpecConfig := cs.GetCloudSpecConfig()
	imageConfigFromCloud := cloudSpecConfig.KubernetesSpecConfig
	k := cs.Properties.OrchestratorProfile.KubernetesConfig

	// Use the cloud-specific config if KubernetesImagesConfig, or if Azure Stack context
	if k.KubernetesImagesConfig == nil || cs.Properties.IsAzureStackCloud() {
		k.KubernetesImagesConfig = &KubernetesImagesConfig{
			ImageBaseConfig: GetImageBaseConfigFromKubernetesSpecConfig(imageConfigFromCloud),
			ImageConfig:     map[string]string{},
		}
		// if KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase is not user-configured, then
		// inherit the deprecated KubernetesConfig.KubernetesImageBase property for backwards compatibility
		if k.KubernetesImageBase != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.AddonManagerImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.CloudControllerManagerImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.K8sDNSSidecarImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.CoreDNSImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.KubeDNSImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.DNSMasqImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.HeapsterImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.AddonResizerImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.ClusterAutoscalerImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.DashboardImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.ReschedulerImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.MetricsServerImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.IPMasqAgentImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.ClusterProportionalAutoscalerImageBase = k.KubernetesImageBase
		}
	} else {
		if k.KubernetesImagesConfig.ImageBaseConfig != nil {
			if k.KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.AddonManagerImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.AddonManagerImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.AddonManagerImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.CloudControllerManagerImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.CloudControllerManagerImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.CloudControllerManagerImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.K8sDNSSidecarImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.K8sDNSSidecarImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.K8sDNSSidecarImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.CoreDNSImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.CoreDNSImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.CoreDNSImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.KubeDNSImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.KubeDNSImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.KubeDNSImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.DNSMasqImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.DNSMasqImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.DNSMasqImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.HeapsterImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.HeapsterImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.HeapsterImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.AddonResizerImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.AddonResizerImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.AddonResizerImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.ClusterAutoscalerImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.ClusterAutoscalerImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.ClusterAutoscalerImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.DashboardImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.DashboardImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.DashboardImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.ReschedulerImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.ReschedulerImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.ReschedulerImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.MetricsServerImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.MetricsServerImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.MetricsServerImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.IPMasqAgentImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.IPMasqAgentImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.IPMasqAgentImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.ClusterProportionalAutoscalerImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.ClusterProportionalAutoscalerImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.ClusterProportionalAutoscalerImageBase = imageConfigFromCloud.KubernetesImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.TillerImageBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.TillerImageBase = imageConfigFromCloud.TillerImageBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.ACIConnectorImageBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.ACIConnectorImageBase = imageConfigFromCloud.ACIConnectorImageBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.NVIDIAImageBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.NVIDIAImageBase = imageConfigFromCloud.NVIDIAImageBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.CalicoImageBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.CalicoImageBase = imageConfigFromCloud.CalicoImageBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIImageBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIImageBase = imageConfigFromCloud.AzureCNIImageBase
			}
		} else {
			k.KubernetesImagesConfig.ImageBaseConfig = GetImageBaseConfigFromKubernetesSpecConfig(imageConfigFromCloud)
		}
	}
}

// GetImageBaseConfigFromKubernetesSpecConfig converts a subset of KubernetesSpecConfig properties to a fully populated ImageBaseConfig
func GetImageBaseConfigFromKubernetesSpecConfig(imageConfigFromCloud KubernetesSpecConfig) *ImageBaseConfig {
	return &ImageBaseConfig{
		KubernetesImageBase:                    imageConfigFromCloud.KubernetesImageBase,
		HyperkubeImageBase:                     imageConfigFromCloud.KubernetesImageBase,
		PauseImageBase:                         imageConfigFromCloud.KubernetesImageBase,
		AddonManagerImageBase:                  imageConfigFromCloud.KubernetesImageBase,
		CloudControllerManagerImageBase:        imageConfigFromCloud.KubernetesImageBase,
		K8sDNSSidecarImageBase:                 imageConfigFromCloud.KubernetesImageBase,
		CoreDNSImageBase:                       imageConfigFromCloud.KubernetesImageBase,
		KubeDNSImageBase:                       imageConfigFromCloud.KubernetesImageBase,
		DNSMasqImageBase:                       imageConfigFromCloud.KubernetesImageBase,
		HeapsterImageBase:                      imageConfigFromCloud.KubernetesImageBase,
		AddonResizerImageBase:                  imageConfigFromCloud.KubernetesImageBase,
		ClusterAutoscalerImageBase:             imageConfigFromCloud.KubernetesImageBase,
		DashboardImageBase:                     imageConfigFromCloud.KubernetesImageBase,
		ReschedulerImageBase:                   imageConfigFromCloud.KubernetesImageBase,
		MetricsServerImageBase:                 imageConfigFromCloud.KubernetesImageBase,
		IPMasqAgentImageBase:                   imageConfigFromCloud.KubernetesImageBase,
		ClusterProportionalAutoscalerImageBase: imageConfigFromCloud.KubernetesImageBase,
		TillerImageBase:                        imageConfigFromCloud.TillerImageBase,
		ACIConnectorImageBase:                  imageConfigFromCloud.ACIConnectorImageBase,
		NVIDIAImageBase:                        imageConfigFromCloud.NVIDIAImageBase,
		CalicoImageBase:                        imageConfigFromCloud.CalicoImageBase,
		AzureCNIImageBase:                      imageConfigFromCloud.AzureCNIImageBase,
	}
}
