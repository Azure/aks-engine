// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

func (cs *ContainerService) setKubernetesImagesConfig() {
	cloudSpecConfig := cs.GetCloudSpecConfig()
	imageConfigFromCloud := cloudSpecConfig.KubernetesSpecConfig
	k := cs.Properties.OrchestratorProfile.KubernetesConfig

	if k.KubernetesImagesConfig == nil {
		k.KubernetesImagesConfig = &KubernetesImagesConfig{
			ImageBaseConfig: getImageBaseConfigFromCloudSpecConfig(imageConfigFromCloud),
			ImageConfig:     map[string]string{},
		}
		// Permit the deprecated KubernetesConfig.KubernetesImageBase property
		// to act as the value for KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase,
		// If KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase is not already present
		if k.KubernetesImageBase != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase = k.KubernetesImageBase
		}
		// ditto KubernetesConfig.AzureCNIURLLinux
		if k.AzureCNIURLLinux != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLLinux = k.AzureCNIURLLinux
		}
		// ditto KubernetesConfig.AzureCNIURLWindows
		if k.AzureCNIURLWindows != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLWindows = k.AzureCNIURLWindows
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
				k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase = imageConfigFromCloud.HyperkubeImageBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = imageConfigFromCloud.PauseImageBase
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
			if k.KubernetesImagesConfig.ImageBaseConfig.EtcdDownloadURLBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.EtcdDownloadURLBase = imageConfigFromCloud.EtcdDownloadURLBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.WindowsBinariesBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.WindowsBinariesBase = imageConfigFromCloud.WindowsBinariesBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.WindowsTelemetryGUID == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.WindowsTelemetryGUID = imageConfigFromCloud.WindowsTelemetryGUID
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.CNIPluginsDownloadURL == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.CNIPluginsDownloadURL = imageConfigFromCloud.CNIPluginsDownloadURL
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLLinux == "" {
				if k.AzureCNIURLLinux != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLLinux = k.AzureCNIURLLinux
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLLinux = imageConfigFromCloud.AzureCNIURLLinux
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLWindows == "" {
				if k.AzureCNIURLWindows != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLWindows = k.AzureCNIURLWindows
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.AzureCNIURLWindows = imageConfigFromCloud.AzureCNIURLWindows
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.ContainerdDownloadURLBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.ContainerdDownloadURLBase = imageConfigFromCloud.ContainerdDownloadURLBase
			}
		} else {
			k.KubernetesImagesConfig.ImageBaseConfig = getImageBaseConfigFromCloudSpecConfig(imageConfigFromCloud)
		}
	}
}

func getImageBaseConfigFromCloudSpecConfig(imageConfigFromCloud KubernetesSpecConfig) *KubernetesSpecConfig {
	return &KubernetesSpecConfig{
		KubernetesImageBase:       imageConfigFromCloud.KubernetesImageBase,
		HyperkubeImageBase:        imageConfigFromCloud.HyperkubeImageBase,
		PauseImageBase:            imageConfigFromCloud.PauseImageBase,
		TillerImageBase:           imageConfigFromCloud.TillerImageBase,
		ACIConnectorImageBase:     imageConfigFromCloud.ACIConnectorImageBase,
		NVIDIAImageBase:           imageConfigFromCloud.NVIDIAImageBase,
		CalicoImageBase:           imageConfigFromCloud.CalicoImageBase,
		AzureCNIImageBase:         imageConfigFromCloud.AzureCNIImageBase,
		EtcdDownloadURLBase:       imageConfigFromCloud.EtcdDownloadURLBase,
		WindowsBinariesBase:       imageConfigFromCloud.WindowsBinariesBase,
		WindowsTelemetryGUID:      imageConfigFromCloud.WindowsTelemetryGUID,
		CNIPluginsDownloadURL:     imageConfigFromCloud.CNIPluginsDownloadURL,
		AzureCNIURLLinux:          imageConfigFromCloud.AzureCNIURLLinux,
		AzureCNIURLWindows:        imageConfigFromCloud.AzureCNIURLWindows,
		ContainerdDownloadURLBase: imageConfigFromCloud.ContainerdDownloadURLBase,
	}
}
