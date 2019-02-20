// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/go-autorest/autorest/azure"
)

func TestConvertCloudProfileToVLabs(t *testing.T) {
	const (
		name                         = "AzureStackCloud"
		managementPortalURL          = "https://management.local.azurestack.external/"
		publishSettingsURL           = "https://management.local.azurestack.external/publishsettings/index"
		serviceManagementEndpoint    = "https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189"
		resourceManagerEndpoint      = "https://management.local.azurestack.external/"
		activeDirectoryEndpoint      = "https://login.windows.net/"
		galleryEndpoint              = "https://portal.local.azurestack.external=30015/"
		keyVaultEndpoint             = "https://vault.azurestack.external/"
		graphEndpoint                = "https://graph.windows.net/"
		serviceBusEndpoint           = "https://servicebus.azurestack.external/"
		batchManagementEndpoint      = "https://batch.azurestack.external/"
		storageEndpointSuffix        = "core.azurestack.external"
		sqlDatabaseDNSSuffix         = "database.azurestack.external"
		trafficManagerDNSSuffix      = "trafficmanager.cn"
		keyVaultDNSSuffix            = "vault.azurestack.external"
		serviceBusEndpointSuffix     = "servicebus.azurestack.external"
		serviceManagementVMDNSSuffix = "chinacloudapp.cn"
		resourceManagerVMDNSSuffix   = "cloudapp.azurestack.external"
		containerRegistryDNSSuffix   = "azurecr.io"
		tokenAudience                = "https://management.azurestack.external/"
	)

	cs := &ContainerService{
		Properties: &Properties{
			CustomCloudProfile: &CustomCloudProfile{
				IdentitySystem:       AzureADIdentitySystem,
				AuthenticationMethod: ClientSecretAuthMethod,
				Environment: &azure.Environment{
					Name:                         name,
					ManagementPortalURL:          managementPortalURL,
					PublishSettingsURL:           publishSettingsURL,
					ServiceManagementEndpoint:    serviceManagementEndpoint,
					ResourceManagerEndpoint:      resourceManagerEndpoint,
					ActiveDirectoryEndpoint:      activeDirectoryEndpoint,
					GalleryEndpoint:              galleryEndpoint,
					KeyVaultEndpoint:             keyVaultEndpoint,
					GraphEndpoint:                graphEndpoint,
					ServiceBusEndpoint:           serviceBusEndpoint,
					BatchManagementEndpoint:      batchManagementEndpoint,
					StorageEndpointSuffix:        storageEndpointSuffix,
					SQLDatabaseDNSSuffix:         sqlDatabaseDNSSuffix,
					TrafficManagerDNSSuffix:      trafficManagerDNSSuffix,
					KeyVaultDNSSuffix:            keyVaultDNSSuffix,
					ServiceBusEndpointSuffix:     serviceBusEndpointSuffix,
					ServiceManagementVMDNSSuffix: serviceManagementVMDNSSuffix,
					ResourceManagerVMDNSSuffix:   resourceManagerVMDNSSuffix,
					ContainerRegistryDNSSuffix:   containerRegistryDNSSuffix,
					TokenAudience:                tokenAudience,
				},
			},
		},
	}

	vlabscs := ConvertContainerServiceToVLabs(cs)

	if vlabscs.Properties.CustomCloudProfile.AuthenticationMethod != ClientSecretAuthMethod {
		t.Errorf("incorrect AuthenticationMethod, expect: '%s', actual: '%s'", ClientSecretAuthMethod, vlabscs.Properties.CustomCloudProfile.AuthenticationMethod)
	}
	if vlabscs.Properties.CustomCloudProfile.IdentitySystem != AzureADIdentitySystem {
		t.Errorf("incorrect IdentitySystem, expect: '%s', actual: '%s'", AzureADIdentitySystem, vlabscs.Properties.CustomCloudProfile.IdentitySystem)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.Name != name {
		t.Errorf("incorrect Name, expect: '%s', actual: '%s'", name, vlabscs.Properties.CustomCloudProfile.Environment.Name)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ManagementPortalURL != managementPortalURL {
		t.Errorf("incorrect ManagementPortalURL, expect: '%s', actual: '%s'", managementPortalURL, vlabscs.Properties.CustomCloudProfile.Environment.ManagementPortalURL)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.PublishSettingsURL != publishSettingsURL {
		t.Errorf("incorrect PublishSettingsURL, expect: '%s', actual: '%s'", publishSettingsURL, vlabscs.Properties.CustomCloudProfile.Environment.PublishSettingsURL)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ServiceManagementEndpoint != serviceManagementEndpoint {
		t.Errorf("incorrect ServiceManagementEndpoint, expect: '%s', actual: '%s'", serviceManagementEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.ServiceManagementEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ResourceManagerEndpoint != resourceManagerEndpoint {
		t.Errorf("incorrect ResourceManagerEndpoint, expect: '%s', actual: '%s'", resourceManagerEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.ResourceManagerEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ActiveDirectoryEndpoint != activeDirectoryEndpoint {
		t.Errorf("incorrect ActiveDirectoryEndpoint, expect: '%s', actual: '%s'", activeDirectoryEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.ActiveDirectoryEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.GalleryEndpoint != galleryEndpoint {
		t.Errorf("incorrect GalleryEndpoint, expect: '%s', actual: '%s'", galleryEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.GalleryEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.KeyVaultEndpoint != keyVaultEndpoint {
		t.Errorf("incorrect KeyVaultEndpoint, expect: '%s', actual: '%s'", keyVaultEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.KeyVaultEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.GraphEndpoint != graphEndpoint {
		t.Errorf("incorrect GraphEndpoint, expect: '%s', actual: '%s'", graphEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.GraphEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ServiceBusEndpoint != serviceBusEndpoint {
		t.Errorf("incorrect ServiceBusEndpoint, expect: '%s', actual: '%s'", serviceBusEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.ServiceBusEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.BatchManagementEndpoint != batchManagementEndpoint {
		t.Errorf("incorrect BatchManagementEndpoint, expect: '%s', actual: '%s'", batchManagementEndpoint, vlabscs.Properties.CustomCloudProfile.Environment.BatchManagementEndpoint)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.StorageEndpointSuffix != storageEndpointSuffix {
		t.Errorf("incorrect StorageEndpointSuffix, expect: '%s', actual: '%s'", storageEndpointSuffix, vlabscs.Properties.CustomCloudProfile.Environment.StorageEndpointSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.SQLDatabaseDNSSuffix != sqlDatabaseDNSSuffix {
		t.Errorf("incorrect SQLDatabaseDNSSuffix, expect: '%s', actual: '%s'", sqlDatabaseDNSSuffix, vlabscs.Properties.CustomCloudProfile.Environment.SQLDatabaseDNSSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.TrafficManagerDNSSuffix != trafficManagerDNSSuffix {
		t.Errorf("incorrect TrafficManagerDNSSuffix, expect: '%s', actual: '%s'", trafficManagerDNSSuffix, vlabscs.Properties.CustomCloudProfile.Environment.TrafficManagerDNSSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.KeyVaultDNSSuffix != keyVaultDNSSuffix {
		t.Errorf("incorrect KeyVaultDNSSuffix, expect: '%s', actual: '%s'", keyVaultDNSSuffix, vlabscs.Properties.CustomCloudProfile.Environment.KeyVaultDNSSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ServiceBusEndpointSuffix != serviceBusEndpointSuffix {
		t.Errorf("incorrect ServiceBusEndpointSuffix, expect: '%s', actual: '%s'", serviceBusEndpointSuffix, vlabscs.Properties.CustomCloudProfile.Environment.ServiceBusEndpointSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ServiceManagementVMDNSSuffix != serviceManagementVMDNSSuffix {
		t.Errorf("incorrect ServiceManagementVMDNSSuffix, expect: '%s', actual: '%s'", serviceManagementVMDNSSuffix, vlabscs.Properties.CustomCloudProfile.Environment.ServiceManagementVMDNSSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix != resourceManagerVMDNSSuffix {
		t.Errorf("incorrect ResourceManagerVMDNSSuffix, expect: '%s', actual: '%s'", resourceManagerVMDNSSuffix, vlabscs.Properties.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.ContainerRegistryDNSSuffix != containerRegistryDNSSuffix {
		t.Errorf("incorrect ContainerRegistryDNSSuffix, expect: '%s', actual: '%s'", containerRegistryDNSSuffix, vlabscs.Properties.CustomCloudProfile.Environment.ContainerRegistryDNSSuffix)
	}
	if vlabscs.Properties.CustomCloudProfile.Environment.TokenAudience != tokenAudience {
		t.Errorf("incorrect TokenAudience, expect: '%s', actual: '%s'", tokenAudience, vlabscs.Properties.CustomCloudProfile.Environment.TokenAudience)
	}
}

func TestConvertAzureEnvironmentSpecConfigToVLabs(t *testing.T) {
	//Mock AzureEnvironmentSpecConfig
	cs := &ContainerService{
		Properties: &Properties{
			CustomCloudProfile: &CustomCloudProfile{
				IdentitySystem:       ADFSIdentitySystem,
				AuthenticationMethod: ClientCertificateAuthMethod,
				AzureEnvironmentSpecConfig: &AzureEnvironmentSpecConfig{
					CloudName: "AzureStackCloud",
					//DockerSpecConfig specify the docker engine download repo
					DockerSpecConfig: DockerSpecConfig{
						DockerEngineRepo:         "DockerEngineRepo",
						DockerComposeDownloadURL: "DockerComposeDownloadURL",
					},
					//KubernetesSpecConfig - Due to Chinese firewall issue, the default containers from google is blocked, use the Chinese local mirror instead
					KubernetesSpecConfig: KubernetesSpecConfig{
						KubernetesImageBase:              "KubernetesImageBase",
						TillerImageBase:                  "TillerImageBase",
						ACIConnectorImageBase:            "ACIConnectorImageBase",
						NVIDIAImageBase:                  "NVIDIAImageBase",
						AzureCNIImageBase:                "AzureCNIImageBase",
						EtcdDownloadURLBase:              "EtcdDownloadURLBase",
						KubeBinariesSASURLBase:           "KubeBinariesSASURLBase",
						WindowsTelemetryGUID:             "WindowsTelemetryGUID",
						CNIPluginsDownloadURL:            "CNIPluginsDownloadURL",
						VnetCNILinuxPluginsDownloadURL:   "VnetCNILinuxPluginsDownloadURL",
						VnetCNIWindowsPluginsDownloadURL: "VnetCNIWindowsPluginsDownloadURL",
						ContainerdDownloadURLBase:        "ContainerdDownloadURLBase",
					},
					DCOSSpecConfig: DCOSSpecConfig{
						DCOS188BootstrapDownloadURL:     "DCOS188BootstrapDownloadURL",
						DCOS190BootstrapDownloadURL:     "DCOS190BootstrapDownloadURL",
						DCOS198BootstrapDownloadURL:     "DCOS198BootstrapDownloadURL",
						DCOS110BootstrapDownloadURL:     "DCOS110BootstrapDownloadURL",
						DCOS111BootstrapDownloadURL:     "DCOS111BootstrapDownloadURL",
						DCOSWindowsBootstrapDownloadURL: "DCOSWindowsBootstrapDownloadURL",
						DcosRepositoryURL:               "DcosRepositoryURL",
						DcosClusterPackageListID:        "DcosClusterPackageListID",
						DcosProviderPackageID:           "DcosProviderPackageID",
					},
					EndpointConfig: AzureEndpointConfig{
						ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
					},
					OSImageConfig: map[Distro]AzureOSImageConfig{
						Distro("Test"): {
							ImageOffer:     "ImageOffer",
							ImageSku:       "ImageSku",
							ImagePublisher: "ImagePublisher",
							ImageVersion:   "ImageVersion",
						},
						AKS:             DefaultAKSOSImageConfig,
						AKSDockerEngine: DefaultAKSDockerEngineOSImageConfig,
					},
				},
			},
		},
	}
	vlabscs := ConvertContainerServiceToVLabs(cs)
	if vlabscs.Properties.CustomCloudProfile.AuthenticationMethod != ClientCertificateAuthMethod {
		t.Errorf("incorrect AuthenticationMethod, expect: '%s', actual: '%s'", ClientCertificateAuthMethod, vlabscs.Properties.CustomCloudProfile.AuthenticationMethod)
	}
	if vlabscs.Properties.CustomCloudProfile.IdentitySystem != ADFSIdentitySystem {
		t.Errorf("incorrect IdentitySystem, expect: '%s', actual: '%s'", ADFSIdentitySystem, vlabscs.Properties.CustomCloudProfile.IdentitySystem)
	}
	csSpec := cs.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig
	vlabscsSpec := vlabscs.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig

	if vlabscsSpec.CloudName != csSpec.CloudName {
		t.Errorf("incorrect CloudName, expect: '%s', actual: '%s'", csSpec.CloudName, vlabscsSpec.CloudName)
	}

	//KubernetesSpecConfig
	if vlabscsSpec.KubernetesSpecConfig.KubernetesImageBase != csSpec.KubernetesSpecConfig.KubernetesImageBase {
		t.Errorf("incorrect KubernetesImageBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.KubernetesImageBase, vlabscsSpec.KubernetesSpecConfig.KubernetesImageBase)
	}
	if vlabscsSpec.KubernetesSpecConfig.TillerImageBase != csSpec.KubernetesSpecConfig.TillerImageBase {
		t.Errorf("incorrect TillerImageBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.TillerImageBase, vlabscsSpec.KubernetesSpecConfig.TillerImageBase)
	}
	if vlabscsSpec.KubernetesSpecConfig.ACIConnectorImageBase != csSpec.KubernetesSpecConfig.ACIConnectorImageBase {
		t.Errorf("incorrect ACIConnectorImageBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.ACIConnectorImageBase, vlabscsSpec.KubernetesSpecConfig.ACIConnectorImageBase)
	}
	if vlabscsSpec.KubernetesSpecConfig.NVIDIAImageBase != csSpec.KubernetesSpecConfig.NVIDIAImageBase {
		t.Errorf("incorrect NVIDIAImageBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.NVIDIAImageBase, vlabscsSpec.KubernetesSpecConfig.NVIDIAImageBase)
	}
	if vlabscsSpec.KubernetesSpecConfig.AzureCNIImageBase != csSpec.KubernetesSpecConfig.AzureCNIImageBase {
		t.Errorf("incorrect AzureCNIImageBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.AzureCNIImageBase, vlabscsSpec.KubernetesSpecConfig.AzureCNIImageBase)
	}
	if vlabscsSpec.KubernetesSpecConfig.EtcdDownloadURLBase != csSpec.KubernetesSpecConfig.EtcdDownloadURLBase {
		t.Errorf("incorrect EtcdDownloadURLBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.EtcdDownloadURLBase, vlabscsSpec.KubernetesSpecConfig.EtcdDownloadURLBase)
	}
	if vlabscsSpec.KubernetesSpecConfig.KubeBinariesSASURLBase != csSpec.KubernetesSpecConfig.KubeBinariesSASURLBase {
		t.Errorf("incorrect KubeBinariesSASURLBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.KubeBinariesSASURLBase, vlabscsSpec.KubernetesSpecConfig.KubeBinariesSASURLBase)
	}
	if vlabscsSpec.KubernetesSpecConfig.WindowsTelemetryGUID != csSpec.KubernetesSpecConfig.WindowsTelemetryGUID {
		t.Errorf("incorrect WindowsTelemetryGUID, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.WindowsTelemetryGUID, vlabscsSpec.KubernetesSpecConfig.WindowsTelemetryGUID)
	}
	if vlabscsSpec.KubernetesSpecConfig.CNIPluginsDownloadURL != csSpec.KubernetesSpecConfig.CNIPluginsDownloadURL {
		t.Errorf("incorrect CNIPluginsDownloadURL, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.CNIPluginsDownloadURL, vlabscsSpec.KubernetesSpecConfig.CNIPluginsDownloadURL)
	}
	if vlabscsSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL != csSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL {
		t.Errorf("incorrect VnetCNILinuxPluginsDownloadURL, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL, vlabscsSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL)
	}
	if vlabscsSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL != csSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL {
		t.Errorf("incorrect VnetCNIWindowsPluginsDownloadURL, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL, vlabscsSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL)
	}
	if vlabscsSpec.KubernetesSpecConfig.ContainerdDownloadURLBase != csSpec.KubernetesSpecConfig.ContainerdDownloadURLBase {
		t.Errorf("incorrect ContainerdDownloadURLBase, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.ContainerdDownloadURLBase, vlabscsSpec.KubernetesSpecConfig.ContainerdDownloadURLBase)
	}

	//DockerSpecConfig
	if vlabscsSpec.DockerSpecConfig.DockerComposeDownloadURL != csSpec.DockerSpecConfig.DockerComposeDownloadURL {
		t.Errorf("incorrect DockerComposeDownloadURL, expect: '%s', actual: '%s'", csSpec.DockerSpecConfig.DockerComposeDownloadURL, vlabscsSpec.DockerSpecConfig.DockerComposeDownloadURL)
	}
	if vlabscsSpec.DockerSpecConfig.DockerEngineRepo != csSpec.DockerSpecConfig.DockerEngineRepo {
		t.Errorf("incorrect DockerEngineRepo, expect: '%s', actual: '%s'", csSpec.DockerSpecConfig.DockerEngineRepo, vlabscsSpec.DockerSpecConfig.DockerEngineRepo)
	}

	//DCOSSpecConfig
	if vlabscsSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL != csSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL {
		t.Errorf("incorrect DCOS188BootstrapDownloadURL, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL, vlabscsSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL)
	}
	if vlabscsSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL != csSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL {
		t.Errorf("incorrect DCOS190BootstrapDownloadURL, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL, vlabscsSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL)
	}
	if vlabscsSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL != csSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL {
		t.Errorf("incorrect DCOS198BootstrapDownloadURL, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL, vlabscsSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL)
	}
	if vlabscsSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL != csSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL {
		t.Errorf("incorrect DCOS110BootstrapDownloadURL, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL, vlabscsSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL)
	}
	if vlabscsSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL != csSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL {
		t.Errorf("incorrect DCOS111BootstrapDownloadURL, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL, vlabscsSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL)
	}
	if vlabscsSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL != csSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL {
		t.Errorf("incorrect DCOSWindowsBootstrapDownloadURL, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL, vlabscsSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL)
	}
	if vlabscsSpec.DCOSSpecConfig.DcosRepositoryURL != csSpec.DCOSSpecConfig.DcosRepositoryURL {
		t.Errorf("incorrect DcosRepositoryURL, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DcosRepositoryURL, vlabscsSpec.DCOSSpecConfig.DcosRepositoryURL)
	}
	if vlabscsSpec.DCOSSpecConfig.DcosClusterPackageListID != csSpec.DCOSSpecConfig.DcosClusterPackageListID {
		t.Errorf("incorrect DcosClusterPackageListID, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DcosClusterPackageListID, vlabscsSpec.DCOSSpecConfig.DcosClusterPackageListID)
	}
	if vlabscsSpec.DCOSSpecConfig.DcosProviderPackageID != csSpec.DCOSSpecConfig.DcosProviderPackageID {
		t.Errorf("incorrect DcosProviderPackageID, expect: '%s', actual: '%s'", csSpec.DCOSSpecConfig.DcosProviderPackageID, vlabscsSpec.DCOSSpecConfig.DcosProviderPackageID)
	}

	//EndpointConfig
	if vlabscsSpec.EndpointConfig.ResourceManagerVMDNSSuffix != csSpec.EndpointConfig.ResourceManagerVMDNSSuffix {
		t.Errorf("incorrect ResourceManagerVMDNSSuffix, expect: '%s', actual: '%s'", csSpec.EndpointConfig.ResourceManagerVMDNSSuffix, vlabscsSpec.EndpointConfig.ResourceManagerVMDNSSuffix)
	}

	//OSImageConfig
	for k, v := range csSpec.OSImageConfig {
		if actualValue, ok := vlabscsSpec.OSImageConfig[vlabs.Distro(string(k))]; ok {
			if v.ImageOffer != actualValue.ImageOffer {
				t.Errorf("incorrect ImageOffer for '%s', expect: '%s', actual: '%s'", string(k), v.ImageOffer, actualValue.ImageOffer)
			}
			if v.ImagePublisher != actualValue.ImagePublisher {
				t.Errorf("incorrect ImagePublisher for '%s', expect: '%s', actual: '%s'", string(k), v.ImagePublisher, actualValue.ImagePublisher)
			}
			if v.ImageSku != actualValue.ImageSku {
				t.Errorf("incorrect ImageSku for '%s', expect: '%s', actual: '%s'", string(k), v.ImageSku, actualValue.ImageSku)
			}
			if v.ImageVersion != actualValue.ImageVersion {
				t.Errorf("incorrect ImageVersion for '%s', expect: '%s', actual: '%s'", string(k), v.ImageVersion, actualValue.ImageVersion)
			}
		} else {
			t.Errorf("incorrect OSImageConfig: '%s' is missing", string(k))
		}
	}
}
