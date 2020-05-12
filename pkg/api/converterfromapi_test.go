// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"net/url"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/google/go-cmp/cmp"
)

const ValidSSHPublicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEApD8+lRvLtUcyfO8N2Cwq0zY9DG1Un9d+tcmU3HgnAzBr6UR/dDT5M07NV7DN1lmu/0dt6Ay/ItjF9xK//nwVJL3ezEX32yhLKkCKFMB1LcANNzlhT++SB5tlRBx65CTL8z9FORe4UCWVJNafxu3as/BshQSrSaYt3hjSeYuzTpwd4+4xQutzbTXEUBDUr01zEfjjzfUu0HDrg1IFae62hnLm3ajG6b432IIdUhFUmgjZDljUt5bI3OEz5IWPsNOOlVTuo6fqU8lJHClAtAlZEZkyv0VotidC7ZSCfV153rRsEk9IWscwL2PQIQnCw7YyEYEffDeLjBwkH6MIdJ6OgQ== rsa-key-20170510"

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
						AzureTelemetryPID:                "AzureTelemetryPID",
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
						CSIProxyDownloadURL:              "CSIProxyDownloadURL",
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
						AKSUbuntu1604: AKSUbuntu1604OSImageConfig,
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
	if vlabscsSpec.KubernetesSpecConfig.AzureTelemetryPID != csSpec.KubernetesSpecConfig.AzureTelemetryPID {
		t.Errorf("incorrect AzureTelemetryPID, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.AzureTelemetryPID, vlabscsSpec.KubernetesSpecConfig.AzureTelemetryPID)
	}
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
	if vlabscsSpec.KubernetesSpecConfig.CSIProxyDownloadURL != csSpec.KubernetesSpecConfig.CSIProxyDownloadURL {
		t.Errorf("incorrect CSIProxyDownloadURL, expect: '%s', actual: '%s'", csSpec.KubernetesSpecConfig.CSIProxyDownloadURL, vlabscsSpec.KubernetesSpecConfig.CSIProxyDownloadURL)
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

func TestConvertContainerServiceToVLabs(t *testing.T) {
	cs := getDefaultContainerService()
	vlabsCS := ConvertContainerServiceToVLabs(cs)
	if vlabsCS == nil {
		t.Errorf("expected the converted containerService struct to be non-nil")
	}
}

func getDefaultContainerService() *ContainerService {
	u, _ := url.Parse("http://foobar.com/search")
	return &ContainerService{
		ID:       "sampleID",
		Location: "westus2",
		Name:     "sampleCS",
		Plan: &ResourcePurchasePlan{
			Name:          "sampleRPP",
			Product:       "sampleProduct",
			PromotionCode: "sampleCode",
			Publisher:     "samplePublisher",
		},
		Tags: map[string]string{
			"foo": "bar",
		},
		Type: "sampleType",
		Properties: &Properties{
			WindowsProfile: &WindowsProfile{
				AdminUsername: "sampleAdminUsername",
				AdminPassword: "sampleAdminPassword",
			},
			DiagnosticsProfile: &DiagnosticsProfile{
				VMDiagnostics: &VMDiagnostics{
					Enabled:    true,
					StorageURL: u,
				},
			},
			LinuxProfile: &LinuxProfile{
				AdminUsername: "azureuser",
				SSH: struct {
					PublicKeys []PublicKey `json:"publicKeys"`
				}{
					PublicKeys: []PublicKey{
						{
							KeyData: ValidSSHPublicKey,
						},
					},
				},
				Secrets: []KeyVaultSecrets{
					{
						SourceVault: &KeyVaultID{
							ID: "sampleKeyVaultID",
						},
						VaultCertificates: []KeyVaultCertificate{
							{
								CertificateURL:   "FooCertURL",
								CertificateStore: "BarCertStore",
							},
						},
					},
				},
				CustomNodesDNS: &CustomNodesDNS{
					DNSServer: "SampleDNSServer",
				},
				CustomSearchDomain: &CustomSearchDomain{
					Name:          "FooCustomSearchDomain",
					RealmUser:     "sampleRealmUser",
					RealmPassword: "sampleRealmPassword",
				},
			},
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: "fooClientID",
				Secret:   "fooSecret",
				ObjectID: "fooObjectID",
				KeyvaultSecretRef: &KeyvaultSecretRef{
					VaultID:       "fooVaultID",
					SecretName:    "fooSecretName",
					SecretVersion: "fooSecretVersion",
				},
			},
			ExtensionProfiles: []*ExtensionProfile{
				{
					Name:                "fooExtension",
					Version:             "fooVersion",
					ExtensionParameters: "fooExtensionParameters",
					ExtensionParametersKeyVaultRef: &KeyvaultSecretRef{
						VaultID:       "fooVaultID",
						SecretName:    "fooSecretName",
						SecretVersion: "fooSecretVersion",
					},
					RootURL:  "fooRootURL",
					Script:   "fooSsript",
					URLQuery: "fooURL",
				},
			},
			JumpboxProfile: &JumpboxProfile{
				OSType:    "Linux",
				DNSPrefix: "blueorange",
				FQDN:      "blueorange.westus2.com",
			},
			CertificateProfile: &CertificateProfile{
				CaCertificate:         "SampleCACert",
				CaPrivateKey:          "SampleCAPrivateKey",
				APIServerCertificate:  "SampleAPIServerCert",
				APIServerPrivateKey:   "SampleAPIServerPrivateKey",
				ClientCertificate:     "SampleClientCert",
				ClientPrivateKey:      "SampleClientPrivateKey",
				KubeConfigCertificate: "SampleKubeConfigCert",
				KubeConfigPrivateKey:  "SampleKubeConfigPrivateKey",
				EtcdClientCertificate: "SampleEtcdClientCert",
				EtcdClientPrivateKey:  "SampleEtcdClientPrivateKey",
				EtcdServerCertificate: "SampleEtcdServerCert",
				EtcdServerPrivateKey:  "SampleEtcdServerPrivateKey",
			},
			FeatureFlags: &FeatureFlags{
				EnableCSERunInBackground: true,
				BlockOutboundInternet:    false,
				EnableTelemetry:          false,
			},
			AADProfile: &AADProfile{
				ClientAppID:     "SampleClientAppID",
				ServerAppID:     "ServerAppID",
				ServerAppSecret: "ServerAppSecret",
				TenantID:        "SampleTenantID",
				AdminGroupID:    "SampleAdminGroupID",
				Authenticator:   Webhook,
			},
			CustomProfile: &CustomProfile{
				Orchestrator: "Kubernetes",
			},
			OrchestratorProfile: &OrchestratorProfile{
				OrchestratorType:    "Kubernetes",
				OrchestratorVersion: "1.11.6",
				DcosConfig: &DcosConfig{
					DcosBootstrapURL:         "SampleDcosBootstrapURL",
					DcosWindowsBootstrapURL:  "SampleWindowsDcosBootstrapURL",
					Registry:                 "SampleRegistry",
					RegistryPass:             "SampleRegistryPass",
					RegistryUser:             "SampleRegistryUser",
					DcosClusterPackageListID: "SampleDcosClusterPackageListID",
					DcosProviderPackageID:    "SampleDcosProviderPackageID",
					BootstrapProfile: &BootstrapProfile{
						VMSize:       "Standard_Ds1_v1",
						OSDiskSizeGB: 256,
						OAuthEnabled: true,
						StaticIP:     "172.0.0.1",
						Subnet:       "255.255.255.0",
					},
				},
				KubernetesConfig: &KubernetesConfig{},
			},
			MasterProfile: &MasterProfile{
				Count:     1,
				DNSPrefix: "blueorange",
				SubjectAltNames: []string{
					"fooSubjectAltName",
				},
				CustomFiles: &[]CustomFile{
					{
						Source: "sampleCustomFileSource",
						Dest:   "sampleCustomFileDest",
					},
				},
				VMSize:                   "Standard_DS1_v1",
				OSDiskSizeGB:             256,
				VnetSubnetID:             "sampleVnetSubnetID",
				Subnet:                   "sampleSubnet",
				VnetCidr:                 "10.240.0.0/8",
				AgentVnetSubnetID:        "sampleAgentVnetSubnetID",
				FirstConsecutiveStaticIP: "10.240.0.0",
				IPAddressCount:           5,
				StorageProfile:           StorageAccount,
				HTTPSourceAddressPrefix:  "fooHTTPSourceAddressPrefix",
				OAuthEnabled:             true,
				PreprovisionExtension: &Extension{
					Name:        "sampleExtension",
					SingleOrAll: "single",
					Template:    "{{foobar}}",
				},
				Extensions: []Extension{
					{
						Name:        "sampleExtension",
						SingleOrAll: "single",
						Template:    "{{foobar}}",
					},
				},
				Distro: Ubuntu,
				ImageRef: &ImageReference{
					Name:          "FooImageRef",
					ResourceGroup: "FooImageRefResourceGroup",
				},
				KubernetesConfig: &KubernetesConfig{
					KubernetesImageBase: "quay.io",
					ClusterSubnet:       "fooClusterSubnet",
					NetworkPolicy:       "calico",
					NetworkPlugin:       "azure-cni",
					ContainerRuntime:    "docker",
					ContainerRuntimeConfig: map[string]string{
						common.ContainerDataDirKey: "/mnt/docker",
					},
					KubeReservedCgroup:              "kubesystem.slice",
					MaxPods:                         3,
					DockerBridgeSubnet:              "sampleDockerSubnet",
					DNSServiceIP:                    "172.0.0.1",
					ServiceCIDR:                     "172.0.0.1/16",
					UseManagedIdentity:              true,
					UserAssignedID:                  "fooUserAssigneID",
					UserAssignedClientID:            "fooUserAssigneClientID",
					MobyVersion:                     "3.0.0",
					CustomHyperkubeImage:            "",
					ContainerdVersion:               "1.2.4",
					CustomCcmImage:                  "sampleCCMImage",
					UseCloudControllerManager:       to.BoolPtr(true),
					CustomWindowsPackageURL:         "https://deisartifacts.windows.net",
					WindowsNodeBinariesURL:          "https://deisartifacts.windows.net",
					UseInstanceMetadata:             to.BoolPtr(true),
					LoadBalancerSku:                 BasicLoadBalancerSku,
					ExcludeMasterFromStandardLB:     to.BoolPtr(false),
					EnableRbac:                      to.BoolPtr(true),
					EnableSecureKubelet:             to.BoolPtr(true),
					EnableAggregatedAPIs:            true,
					EnableDataEncryptionAtRest:      to.BoolPtr(true),
					EnablePodSecurityPolicy:         to.BoolPtr(true),
					EnableEncryptionWithExternalKms: to.BoolPtr(true),
					GCHighThreshold:                 85,
					GCLowThreshold:                  80,
					EtcdVersion:                     "3.0.0",
					EtcdDiskSizeGB:                  "256",
					EtcdEncryptionKey:               "sampleEncruptionKey",
					AzureCNIVersion:                 "1.1.2",
					AzureCNIURLLinux:                "https://mirror.azk8s.cn/kubernetes/azure-container-networking/linux",
					AzureCNIURLWindows:              "https://mirror.azk8s.cn/kubernetes/azure-container-networking/windows",
					KeyVaultSku:                     "Basic",
					MaximumLoadBalancerRuleCount:    3,
					ProxyMode:                       KubeProxyModeIPTables,
					PrivateAzureRegistryServer:      "sampleRegistryServerURL",
					KubeletConfig: map[string]string{
						"barKey": "bazValue",
					},
					Addons: []KubernetesAddon{
						{
							Name:    "sampleAddon",
							Enabled: to.BoolPtr(true),
							Containers: []KubernetesContainerSpec{
								{
									Name:           "sampleK8sContainer",
									Image:          "sampleK8sImage",
									MemoryRequests: "20Mi",
									CPURequests:    "10m",
								},
							},
							Config: map[string]string{
								"sampleKey": "sampleVal",
							},
						},
					},
					APIServerConfig: map[string]string{
						"sampleAPIServerKey": "sampleAPIServerVal",
					},
					ControllerManagerConfig: map[string]string{
						"sampleCMKey": "sampleCMVal",
					},
					CloudControllerManagerConfig: map[string]string{
						"sampleCCMKey": "sampleCCMVal",
					},
					SchedulerConfig: map[string]string{
						"sampleSchedulerKey": "sampleSchedulerVal",
					},
					PrivateCluster: &PrivateCluster{
						Enabled: to.BoolPtr(true),
						JumpboxProfile: &PrivateJumpboxProfile{
							Name:           "sampleJumpboxProfile",
							VMSize:         "Standard_DS1_v2",
							OSDiskSizeGB:   512,
							Username:       "userName",
							PublicKey:      ValidSSHPublicKey,
							StorageProfile: StorageAccount,
						},
					},
					PodSecurityPolicyConfig: map[string]string{
						"samplePSPConfigKey": "samplePSPConfigVal",
					},
				},
			},
			AgentPoolProfiles: []*AgentPoolProfile{
				{
					Name:      "sampleAgent",
					Count:     2,
					VMSize:    "sampleVM",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.com",
					OSType:    "Linux",
					Subnet:    "sampleSubnet",
				},
				{
					Name:      "sampleAgent-public",
					Count:     2,
					VMSize:    "sampleVM",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.com",
					OSType:    "Linux",
					Subnet:    "sampleSubnet",
					ImageRef: &ImageReference{
						Name:           "testImage",
						ResourceGroup:  "testRg",
						SubscriptionID: "testSub",
						Gallery:        "testGallery",
						Version:        "0.0.1",
					},
				},
			},
		},
	}
}

func TestConvertOrchestratorVersionProfileToVLabs(t *testing.T) {
	ovpK8s := &OrchestratorVersionProfile{
		OrchestratorProfile: OrchestratorProfile{
			OrchestratorType:    Kubernetes,
			OrchestratorVersion: "1.9.11",
		},
		Upgrades: []*OrchestratorProfile{
			{
				OrchestratorType:    Kubernetes,
				OrchestratorVersion: "1.10.13",
			},
			{
				OrchestratorType:    Kubernetes,
				OrchestratorVersion: "1.11.6",
			},
		},
	}

	vlabsOvp := ConvertOrchestratorVersionProfileToVLabs(ovpK8s)

	if vlabsOvp == nil {
		t.Errorf("expected the converted orchestratorVersionProfileToVLabs struct to be non-nil")
	}
}

func TestTelemetryEnabledToVLabs(t *testing.T) {
	cs := getDefaultContainerService()
	cs.Properties.FeatureFlags.EnableTelemetry = true
	vlabsCS := ConvertContainerServiceToVLabs(cs)
	if vlabsCS == nil {
		t.Errorf("expected the converted containerService struct to be non-nil")
	}
	if !vlabsCS.Properties.FeatureFlags.EnableTelemetry {
		t.Errorf("expected the EnableTelemetry feature flag to be true")
	}
}

func TestTelemetryDefaultToVLabs(t *testing.T) {
	cs := getDefaultContainerService()
	vlabsCS := ConvertContainerServiceToVLabs(cs)
	if vlabsCS == nil {
		t.Errorf("expected the converted containerService struct to be non-nil")
	}
	if vlabsCS.Properties.FeatureFlags.EnableTelemetry {
		t.Errorf("expected the EnableTelemetry feature flag to be false")
	}
}

func TestPPGToVLabs(t *testing.T) {
	ppgResourceID1 := "ppgResourceID1"
	ppgResourceID2 := "ppgResourceID2"
	cs := getDefaultContainerService()
	cs.Properties.MasterProfile.ProximityPlacementGroupID = ppgResourceID1
	cs.Properties.AgentPoolProfiles[0].ProximityPlacementGroupID = ppgResourceID2
	vlabsCS := ConvertContainerServiceToVLabs(cs)
	if vlabsCS == nil {
		t.Errorf("expected the converted containerService struct to be non-nil")
	}
	if vlabsCS.Properties.MasterProfile.ProximityPlacementGroupID != ppgResourceID1 {
		t.Errorf("expected the agent pool profile proximity placement group to be %s", ppgResourceID1)
	}

	if vlabsCS.Properties.AgentPoolProfiles[0].ProximityPlacementGroupID != ppgResourceID2 {
		t.Errorf("expected the agent pool profile proximity placement group to be %s", ppgResourceID2)
	}
}

func TestPlatformFaultDomainCountToVLabs(t *testing.T) {
	cs := getDefaultContainerService()
	cs.Properties.MasterProfile.PlatformFaultDomainCount = to.IntPtr(3)
	cs.Properties.AgentPoolProfiles[0].PlatformFaultDomainCount = to.IntPtr(5)
	vlabsCS := ConvertContainerServiceToVLabs(cs)
	if vlabsCS == nil {
		t.Errorf("expected the converted containerService struct to be non-nil")
	}
	if *vlabsCS.Properties.MasterProfile.PlatformFaultDomainCount != 3 {
		t.Errorf("expected the master profile platform FD to be 3")
	}
	if *vlabsCS.Properties.AgentPoolProfiles[0].PlatformFaultDomainCount != 5 {
		t.Errorf("expected the agent pool profile platform FD to be 5")
	}
}

func TestPlatformUpdateDomainCountToVLabs(t *testing.T) {
	cs := getDefaultContainerService()
	cs.Properties.MasterProfile.PlatformUpdateDomainCount = to.IntPtr(3)
	cs.Properties.AgentPoolProfiles[0].PlatformUpdateDomainCount = to.IntPtr(3)
	vlabsCS := ConvertContainerServiceToVLabs(cs)
	if vlabsCS == nil {
		t.Errorf("expected the converted containerService struct to be non-nil")
	}
	if *vlabsCS.Properties.MasterProfile.PlatformUpdateDomainCount != 3 {
		t.Errorf("expected the master profile platform FD to be 3")
	}
	if *vlabsCS.Properties.AgentPoolProfiles[0].PlatformUpdateDomainCount != 3 {
		t.Errorf("expected the agent pool profile platform FD to be 3")
	}
}

func TestConvertTelemetryProfileToVLabs(t *testing.T) {
	cs := getDefaultContainerService()
	cs.Properties.TelemetryProfile = &TelemetryProfile{
		ApplicationInsightsKey: "app_insights_key",
	}

	vlabsCS := ConvertContainerServiceToVLabs(cs)

	if vlabsCS.Properties.TelemetryProfile == nil {
		t.Error("expected ConvertContainerServiceToVLabs to set TelemtryProfile")
	}

	if vlabsCS.Properties.TelemetryProfile.ApplicationInsightsKey != "app_insights_key" {
		t.Error("TelemetryProfile.APplicationInsightsKey not converted")
	}
}

func TestConvertWindowsProfileToVlabs(t *testing.T) {
	falseVar := false

	cases := []struct {
		name     string
		w        WindowsProfile
		expected vlabs.WindowsProfile
	}{
		{
			name: "empty profile",
			w:    WindowsProfile{},
			expected: vlabs.WindowsProfile{
				Secrets: []vlabs.KeyVaultSecrets{},
			},
		},
		{
			name: "misc fields",
			w: WindowsProfile{
				AdminUsername:          "user",
				AdminPassword:          "password",
				EnableAutomaticUpdates: &falseVar,
				ImageVersion:           "17763.615.1907121548",
				SSHEnabled:             &falseVar,
				WindowsPublisher:       "MicrosoftWindowsServer",
				WindowsOffer:           "WindowsServer",
				WindowsSku:             "2019-Datacenter-Core-smalldisk",
				WindowsDockerVersion:   "18.09",
			},
			expected: vlabs.WindowsProfile{
				AdminUsername:          "user",
				AdminPassword:          "password",
				EnableAutomaticUpdates: &falseVar,
				ImageVersion:           "17763.615.1907121548",
				SSHEnabled:             &falseVar,
				WindowsPublisher:       "MicrosoftWindowsServer",
				WindowsOffer:           "WindowsServer",
				WindowsSku:             "2019-Datacenter-Core-smalldisk",
				WindowsDockerVersion:   "18.09",
				Secrets:                []vlabs.KeyVaultSecrets{},
			},
		},
		{
			name: "image reference",
			w: WindowsProfile{
				ImageRef: &ImageReference{
					Gallery:        "gallery",
					Name:           "name",
					ResourceGroup:  "rg",
					SubscriptionID: "dc6bd10c-110c-4134-88c5-4d5a039129c4",
					Version:        "1.25.6",
				},
			},
			expected: vlabs.WindowsProfile{
				ImageRef: &vlabs.ImageReference{
					Gallery:        "gallery",
					Name:           "name",
					ResourceGroup:  "rg",
					SubscriptionID: "dc6bd10c-110c-4134-88c5-4d5a039129c4",
					Version:        "1.25.6",
				},
				Secrets: []vlabs.KeyVaultSecrets{},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actual := vlabs.WindowsProfile{}
			convertWindowsProfileToVLabs(&c.w, &actual)

			diff := cmp.Diff(actual, c.expected)

			if diff != "" {
				t.Errorf("unexpected diff testing convertWindowsProfileToVLabs: %s", diff)
			}
		})
	}
}

func TestConvertComponentsToVlabs(t *testing.T) {
	k := &KubernetesConfig{
		Components: []KubernetesComponent{
			{
				Name:    "component-0",
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name:           "component-0-container-0",
						Image:          "baz",
						CPURequests:    "1",
						MemoryRequests: "200m",
						CPULimits:      "2",
						MemoryLimits:   "400m",
					},
					{
						Name:           "component-0-container-1",
						Image:          "baz-1",
						CPURequests:    "1-1",
						MemoryRequests: "200m-1",
						CPULimits:      "2-1",
						MemoryLimits:   "400m-1",
					},
				},
				Config: map[string]string{
					"foo":     "bar",
					"command": "my-command",
				},
				Data: "my-data",
			},
			{
				Name:    "component-1",
				Enabled: to.BoolPtr(false),
				Containers: []KubernetesContainerSpec{
					{
						Name:           "component-1-container-0",
						Image:          "baz",
						CPURequests:    "1",
						MemoryRequests: "200m",
						CPULimits:      "2",
						MemoryLimits:   "400m",
					},
					{
						Name:           "component-1-container-1",
						Image:          "baz-1",
						CPURequests:    "1-1",
						MemoryRequests: "200m-1",
						CPULimits:      "2-1",
						MemoryLimits:   "400m-1",
					},
				},
				Config: map[string]string{
					"foo":     "bar",
					"command": "my-command",
				},
				Data: "my-data",
			},
		},
	}
	vk := &vlabs.KubernetesConfig{}
	convertComponentsToVlabs(k, vk)
	for i, component := range k.Components {
		if vk.Components[i].Name != component.Name {
			t.Errorf("unexpected Component.Name property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Name, component.Name)
		}
		if to.Bool(vk.Components[i].Enabled) != to.Bool(component.Enabled) {
			t.Errorf("unexpected Component.Enabled property %t after convertComponentsToVlabs conversion, expected %t", to.Bool(vk.Components[i].Enabled), to.Bool(component.Enabled))
		}
		if vk.Components[i].Data != component.Data {
			t.Errorf("unexpected Component.Data property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Data, component.Data)
		}
		for j, container := range component.Containers {
			if vk.Components[i].Containers[j].Name != container.Name {
				t.Errorf("unexpected Container.Name property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Containers[j].Name, container.Name)
			}
			if vk.Components[i].Containers[j].Image != container.Image {
				t.Errorf("unexpected Container.Image property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Containers[j].Image, container.Image)
			}
			if vk.Components[i].Containers[j].CPURequests != container.CPURequests {
				t.Errorf("unexpected Container.CPURequests property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Containers[j].CPURequests, container.CPURequests)
			}
			if vk.Components[i].Containers[j].MemoryRequests != container.MemoryRequests {
				t.Errorf("unexpected Container.MemoryRequests property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Containers[j].MemoryRequests, container.MemoryRequests)
			}
			if vk.Components[i].Containers[j].CPULimits != container.CPULimits {
				t.Errorf("unexpected Container.CPULimits property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Containers[j].CPULimits, container.CPULimits)
			}
			if vk.Components[i].Containers[j].MemoryLimits != container.MemoryLimits {
				t.Errorf("unexpected Container.MemoryLimits property %s after convertComponentsToVlabs conversion, expected %s", vk.Components[i].Containers[j].MemoryLimits, container.MemoryLimits)
			}
		}
		for key, val := range component.Config {
			if vk.Components[i].Config[key] != val {
				t.Errorf("unexpected Component.Config %s=%s after convertComponentsToVlabs conversion, expected %s=%s", key, vk.Components[i].Config[key], key, val)
			}
		}
		for key, val := range vk.Components[i].Config {
			if component.Config[key] != val {
				t.Errorf("unexpected Component.Config %s=%s after convertComponentsToVlabs conversion, expected %s=%s", key, component.Config[key], key, val)
			}
		}
	}
}
