// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import "github.com/Azure/aks-engine/pkg/helpers"

func (p *Properties) setCustomCloudProfileDefaults() {

	if p.IsAzureStackCloud() {

		//azureStackCloudSpec is the default configurations for azure stack with public Azure.
		azureStackCloudSpec := AzureEnvironmentSpecConfig{
			CloudName: AzureStackCloud,
			//DockerSpecConfig specify the docker engine download repo
			DockerSpecConfig: DefaultDockerSpecConfig,
			//KubernetesSpecConfig is the default kubernetes container image url.
			KubernetesSpecConfig: DefaultKubernetesSpecConfig,
			DCOSSpecConfig:       DefaultDCOSSpecConfig,
			EndpointConfig: AzureEndpointConfig{
				ResourceManagerVMDNSSuffix: "",
			},
			OSImageConfig: map[Distro]AzureOSImageConfig{
				Ubuntu:          DefaultUbuntuImageConfig,
				RHEL:            DefaultRHELOSImageConfig,
				CoreOS:          DefaultCoreOSImageConfig,
				AKS:             DefaultAKSOSImageConfig,
				AKSDockerEngine: DefaultAKSDockerEngineOSImageConfig,
			},
		}

		//Set default value for ResourceManagerVMDNSSuffix
		azureStackCloudSpec.EndpointConfig.ResourceManagerVMDNSSuffix = p.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix

		// Use the custom input to overwrite the default values in AzureStackCloudSpec
		if p.CustomCloudProfile.AzureEnvironmentSpecConfig != nil {
			ascc := p.CustomCloudProfile.AzureEnvironmentSpecConfig
			azureStackCloudSpec.CloudName = helpers.EnsureString(ascc.CloudName, azureStackCloudSpec.CloudName)

			// DockerSpecConfig
			asccDockerSpecConfig := ascc.DockerSpecConfig
			azsDockerSpecConfig := azureStackCloudSpec.DockerSpecConfig
			azureStackCloudSpec.DockerSpecConfig.DockerComposeDownloadURL = helpers.EnsureString(asccDockerSpecConfig.DockerComposeDownloadURL, azsDockerSpecConfig.DockerComposeDownloadURL)
			azureStackCloudSpec.DockerSpecConfig.DockerEngineRepo = helpers.EnsureString(asccDockerSpecConfig.DockerEngineRepo, azsDockerSpecConfig.DockerComposeDownloadURL)

			//KubernetesSpecConfig
			asccKubernetesSpecConfig := ascc.KubernetesSpecConfig
			azsKubernetesSpecConfig := azureStackCloudSpec.KubernetesSpecConfig
			azureStackCloudSpec.KubernetesSpecConfig.ACIConnectorImageBase = helpers.EnsureString(asccKubernetesSpecConfig.ACIConnectorImageBase, azsKubernetesSpecConfig.ACIConnectorImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.AzureCNIImageBase = helpers.EnsureString(asccKubernetesSpecConfig.AzureCNIImageBase, azsKubernetesSpecConfig.AzureCNIImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.CNIPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.CNIPluginsDownloadURL, azsKubernetesSpecConfig.CNIPluginsDownloadURL)
			azureStackCloudSpec.KubernetesSpecConfig.ContainerdDownloadURLBase = helpers.EnsureString(asccKubernetesSpecConfig.ContainerdDownloadURLBase, azsKubernetesSpecConfig.ContainerdDownloadURLBase)
			azureStackCloudSpec.KubernetesSpecConfig.EtcdDownloadURLBase = helpers.EnsureString(asccKubernetesSpecConfig.EtcdDownloadURLBase, azsKubernetesSpecConfig.EtcdDownloadURLBase)
			azureStackCloudSpec.KubernetesSpecConfig.KubeBinariesSASURLBase = helpers.EnsureString(asccKubernetesSpecConfig.KubeBinariesSASURLBase, azsKubernetesSpecConfig.KubeBinariesSASURLBase)
			azureStackCloudSpec.KubernetesSpecConfig.KubernetesImageBase = helpers.EnsureString(asccKubernetesSpecConfig.KubernetesImageBase, azsKubernetesSpecConfig.KubernetesImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.NVIDIAImageBase = helpers.EnsureString(asccKubernetesSpecConfig.NVIDIAImageBase, azsKubernetesSpecConfig.NVIDIAImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.TillerImageBase = helpers.EnsureString(asccKubernetesSpecConfig.TillerImageBase, azsKubernetesSpecConfig.TillerImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL, azsKubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL)
			azureStackCloudSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL, azsKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL)
			azureStackCloudSpec.KubernetesSpecConfig.WindowsTelemetryGUID = helpers.EnsureString(asccKubernetesSpecConfig.WindowsTelemetryGUID, azsKubernetesSpecConfig.WindowsTelemetryGUID)

			//EndpointConfig
			asccEndpointConfig := ascc.EndpointConfig
			azsEndpointConfig := azureStackCloudSpec.EndpointConfig
			azureStackCloudSpec.EndpointConfig.ResourceManagerVMDNSSuffix = helpers.EnsureString(asccEndpointConfig.ResourceManagerVMDNSSuffix, azsEndpointConfig.ResourceManagerVMDNSSuffix)

			//OSImageConfig
			azureStackCloudSpec.OSImageConfig = make(map[Distro]AzureOSImageConfig)
			for k, v := range ascc.OSImageConfig {
				azureStackCloudSpec.OSImageConfig[k] = v
			}
		}
		p.CustomCloudProfile.AzureEnvironmentSpecConfig = &azureStackCloudSpec
		AzureCloudSpecEnvMap[AzureStackCloud] = azureStackCloudSpec

		p.CustomCloudProfile.AuthenticationMethod = helpers.EnsureString(p.CustomCloudProfile.AuthenticationMethod, ClientSecretAuthMethod)
		p.CustomCloudProfile.IdentitySystem = helpers.EnsureString(p.CustomCloudProfile.IdentitySystem, AzureADIdentitySystem)
	}
}
