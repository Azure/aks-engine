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
		// to act as the value for KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase
		// and KubernetesImagesConfig.ImageBaseConfig.PauseImageBase,
		// if KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase is not already present
		if k.KubernetesImageBase != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.KubernetesImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase = k.KubernetesImageBase
			k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = k.KubernetesImageBase
		}
		// ditto KubernetesConfig.AzureCNIURLLinux
		if k.AzureCNIURLLinux != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.VnetCNILinuxPluginsDownloadURL = k.AzureCNIURLLinux
		}
		// ditto KubernetesConfig.AzureCNIURLWindows
		if k.AzureCNIURLWindows != "" {
			k.KubernetesImagesConfig.ImageBaseConfig.VnetCNIWindowsPluginsDownloadURL = k.AzureCNIURLWindows
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
					k.KubernetesImagesConfig.ImageBaseConfig.HyperkubeImageBase = imageConfigFromCloud.HyperkubeImageBase
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase == "" {
				if k.KubernetesImageBase != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = k.KubernetesImageBase
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.PauseImageBase = imageConfigFromCloud.PauseImageBase
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
			if k.KubernetesImagesConfig.ImageBaseConfig.EtcdDownloadURLBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.EtcdDownloadURLBase = imageConfigFromCloud.EtcdDownloadURLBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.KubeBinariesSASURLBase == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.KubeBinariesSASURLBase = imageConfigFromCloud.KubeBinariesSASURLBase
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.WindowsTelemetryGUID == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.WindowsTelemetryGUID = imageConfigFromCloud.WindowsTelemetryGUID
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.CNIPluginsDownloadURL == "" {
				k.KubernetesImagesConfig.ImageBaseConfig.CNIPluginsDownloadURL = imageConfigFromCloud.CNIPluginsDownloadURL
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.VnetCNILinuxPluginsDownloadURL == "" {
				if k.AzureCNIURLLinux != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.VnetCNILinuxPluginsDownloadURL = k.AzureCNIURLLinux
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.VnetCNILinuxPluginsDownloadURL = imageConfigFromCloud.VnetCNILinuxPluginsDownloadURL
				}
			}
			if k.KubernetesImagesConfig.ImageBaseConfig.VnetCNIWindowsPluginsDownloadURL == "" {
				if k.AzureCNIURLWindows != "" {
					k.KubernetesImagesConfig.ImageBaseConfig.VnetCNIWindowsPluginsDownloadURL = k.AzureCNIURLWindows
				} else {
					k.KubernetesImagesConfig.ImageBaseConfig.VnetCNIWindowsPluginsDownloadURL = imageConfigFromCloud.VnetCNIWindowsPluginsDownloadURL
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
		KubernetesImageBase:              imageConfigFromCloud.KubernetesImageBase,
		HyperkubeImageBase:               imageConfigFromCloud.HyperkubeImageBase,
		PauseImageBase:                   imageConfigFromCloud.PauseImageBase,
		TillerImageBase:                  imageConfigFromCloud.TillerImageBase,
		ACIConnectorImageBase:            imageConfigFromCloud.ACIConnectorImageBase,
		NVIDIAImageBase:                  imageConfigFromCloud.NVIDIAImageBase,
		CalicoImageBase:                  imageConfigFromCloud.CalicoImageBase,
		AzureCNIImageBase:                imageConfigFromCloud.AzureCNIImageBase,
		EtcdDownloadURLBase:              imageConfigFromCloud.EtcdDownloadURLBase,
		KubeBinariesSASURLBase:           imageConfigFromCloud.KubeBinariesSASURLBase,
		WindowsTelemetryGUID:             imageConfigFromCloud.WindowsTelemetryGUID,
		CNIPluginsDownloadURL:            imageConfigFromCloud.CNIPluginsDownloadURL,
		VnetCNILinuxPluginsDownloadURL:   imageConfigFromCloud.VnetCNILinuxPluginsDownloadURL,
		VnetCNIWindowsPluginsDownloadURL: imageConfigFromCloud.VnetCNIWindowsPluginsDownloadURL,
		ContainerdDownloadURLBase:        imageConfigFromCloud.ContainerdDownloadURLBase,
	}
}
