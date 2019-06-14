// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

func (cs *ContainerService) setKubernetesImagesConfig() {
	cloudSpecConfig := cs.GetCloudSpecConfig()
	imageConfigFromCloud := cloudSpecConfig.KubernetesSpecConfig
	k := cs.Properties.OrchestratorProfile.KubernetesConfig

	if k.KubernetesImagesConfig == nil {
		k.KubernetesImagesConfig = &KubernetesImagesConfig{
			ImageBaseConfig: GetImageBaseConfigFromKubernetesSpecConfig(imageConfigFromCloud),
			ImageConfig:     map[string]string{},
		}
		// Permit the deprecated KubernetesConfig.KubernetesImageBase property
		// to act as the value for KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.PauseImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.AddonManagerImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.CloudControllerManagerImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.K8sDNSSidecarImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.CoreDNSImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.KubeDNSImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.DNSMasqImageBase,
		// if KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase is not already present
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
		KubernetesImageBase:             imageConfigFromCloud.KubernetesImageBase,
		HyperkubeImageBase:              imageConfigFromCloud.KubernetesImageBase,
		PauseImageBase:                  imageConfigFromCloud.KubernetesImageBase,
		AddonManagerImageBase:           imageConfigFromCloud.KubernetesImageBase,
		CloudControllerManagerImageBase: imageConfigFromCloud.KubernetesImageBase,
		K8sDNSSidecarImageBase:          imageConfigFromCloud.KubernetesImageBase,
		CoreDNSImageBase:                imageConfigFromCloud.KubernetesImageBase,
		KubeDNSImageBase:                imageConfigFromCloud.KubernetesImageBase,
		DNSMasqImageBase:                imageConfigFromCloud.KubernetesImageBase,
		TillerImageBase:                 imageConfigFromCloud.TillerImageBase,
		ACIConnectorImageBase:           imageConfigFromCloud.ACIConnectorImageBase,
		NVIDIAImageBase:                 imageConfigFromCloud.NVIDIAImageBase,
		CalicoImageBase:                 imageConfigFromCloud.CalicoImageBase,
		AzureCNIImageBase:               imageConfigFromCloud.AzureCNIImageBase,
	}
}
