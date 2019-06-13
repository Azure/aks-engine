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
		// and KubernetesImagesConfig.ImageBaseConfig.PauseImageBase,
		// if KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase is not already present
		if k.KubernetesImageBase != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = k.KubernetesImageBase
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

func GetImageBaseConfigFromKubernetesSpecConfig(imageConfigFromCloud KubernetesSpecConfig) *ImageBaseConfig {
	return &ImageBaseConfig{
		KubernetesImageBase:   imageConfigFromCloud.KubernetesImageBase,
		HyperkubeImageBase:    imageConfigFromCloud.KubernetesImageBase,
		PauseImageBase:        imageConfigFromCloud.KubernetesImageBase,
		TillerImageBase:       imageConfigFromCloud.TillerImageBase,
		ACIConnectorImageBase: imageConfigFromCloud.ACIConnectorImageBase,
		NVIDIAImageBase:       imageConfigFromCloud.NVIDIAImageBase,
		CalicoImageBase:       imageConfigFromCloud.CalicoImageBase,
		AzureCNIImageBase:     imageConfigFromCloud.AzureCNIImageBase,
	}
}
