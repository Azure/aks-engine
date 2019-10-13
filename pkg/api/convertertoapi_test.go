// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"net/url"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"

	v20160330 "github.com/Azure/aks-engine/pkg/api/v20160330"
	v20170131 "github.com/Azure/aks-engine/pkg/api/v20170131"

	"github.com/davecgh/go-spew/spew"
	"k8s.io/apimachinery/pkg/api/equality"

	"github.com/Azure/aks-engine/pkg/api/common"
	v20160930 "github.com/Azure/aks-engine/pkg/api/v20160930"
	v20170701 "github.com/Azure/aks-engine/pkg/api/v20170701"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/google/go-cmp/cmp"
)

func TestAddDCOSPublicAgentPool(t *testing.T) {
	expectedNumPools := 2
	for _, masterCount := range [2]int{1, 3} {
		profiles := []*AgentPoolProfile{}
		profile := makeAgentPoolProfile(1, "agentprivate", "test-dcos-pool", "Standard_D2_v2", Linux)
		profiles = append(profiles, profile)
		master := makeMasterProfile(masterCount, "test-dcos", "Standard_D2_v2")
		props := getProperties(profiles, master)
		expectedPublicPoolName := props.AgentPoolProfiles[0].Name + publicAgentPoolSuffix
		expectedPublicDNSPrefix := props.AgentPoolProfiles[0].DNSPrefix
		expectedPrivateDNSPrefix := ""
		expectedPublicOSType := props.AgentPoolProfiles[0].OSType
		expectedPublicVMSize := props.AgentPoolProfiles[0].VMSize
		addDCOSPublicAgentPool(props)
		if len(props.AgentPoolProfiles) != expectedNumPools {
			t.Fatalf("incorrect agent pools count. expected=%d actual=%d", expectedNumPools, len(props.AgentPoolProfiles))
		}
		if props.AgentPoolProfiles[1].Name != expectedPublicPoolName {
			t.Fatalf("incorrect public pool name. expected=%s actual=%s", expectedPublicPoolName, props.AgentPoolProfiles[1].Name)
		}
		if props.AgentPoolProfiles[1].DNSPrefix != expectedPublicDNSPrefix {
			t.Fatalf("incorrect public pool DNS prefix. expected=%s actual=%s", expectedPublicDNSPrefix, props.AgentPoolProfiles[1].DNSPrefix)
		}
		if props.AgentPoolProfiles[0].DNSPrefix != expectedPrivateDNSPrefix {
			t.Fatalf("incorrect private pool DNS prefix. expected=%s actual=%s", expectedPrivateDNSPrefix, props.AgentPoolProfiles[0].DNSPrefix)
		}
		if props.AgentPoolProfiles[1].OSType != expectedPublicOSType {
			t.Fatalf("incorrect public pool OS type. expected=%s actual=%s", expectedPublicOSType, props.AgentPoolProfiles[1].OSType)
		}
		if props.AgentPoolProfiles[1].VMSize != expectedPublicVMSize {
			t.Fatalf("incorrect public pool VM size. expected=%s actual=%s", expectedPublicVMSize, props.AgentPoolProfiles[1].VMSize)
		}
		for i, port := range [3]int{80, 443, 8080} {
			if props.AgentPoolProfiles[1].Ports[i] != port {
				t.Fatalf("incorrect public pool port assignment. expected=%d actual=%d", port, props.AgentPoolProfiles[1].Ports[i])
			}
		}
		if props.AgentPoolProfiles[1].Count != masterCount {
			t.Fatalf("incorrect public pool VM size. expected=%d actual=%d", masterCount, props.AgentPoolProfiles[1].Count)
		}
	}
}

func makeAgentPoolProfile(count int, name, dNSPrefix, vMSize string, oSType OSType) *AgentPoolProfile {
	return &AgentPoolProfile{
		Name:      name,
		Count:     count,
		DNSPrefix: dNSPrefix,
		OSType:    oSType,
		VMSize:    vMSize,
	}
}

func makeMasterProfile(count int, dNSPrefix, vMSize string) *MasterProfile {
	return &MasterProfile{
		Count:     count,
		DNSPrefix: dNSPrefix,
		VMSize:    vMSize,
	}
}

func getProperties(profiles []*AgentPoolProfile, master *MasterProfile) *Properties {
	return &Properties{
		AgentPoolProfiles: profiles,
		MasterProfile:     master,
	}
}

func TestOrchestratorVersion(t *testing.T) {
	// test v20170701
	v20170701cs := &v20170701.ContainerService{
		Properties: &v20170701.Properties{
			OrchestratorProfile: &v20170701.OrchestratorProfile{
				OrchestratorType: v20170701.Kubernetes,
			},
		},
	}
	cs := ConvertV20170701ContainerService(v20170701cs, false)
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(false) {
		t.Fatalf("incorrect OrchestratorVersion '%s'", cs.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	v20170701cs = &v20170701.ContainerService{
		Properties: &v20170701.Properties{
			OrchestratorProfile: &v20170701.OrchestratorProfile{
				OrchestratorType:    v20170701.Kubernetes,
				OrchestratorVersion: "1.7.14",
			},
		},
	}
	cs = ConvertV20170701ContainerService(v20170701cs, true)
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != "1.7.14" {
		t.Fatalf("incorrect OrchestratorVersion '%s'", cs.Properties.OrchestratorProfile.OrchestratorVersion)
	}
	// test vlabs
	vlabscs := &vlabs.ContainerService{
		Properties: &vlabs.Properties{
			OrchestratorProfile: &vlabs.OrchestratorProfile{
				OrchestratorType: vlabs.Kubernetes,
			},
		},
	}
	cs, err := ConvertVLabsContainerService(vlabscs, false)
	if err != nil {
		t.Fatalf("Failed to convert ContainerService, error: %s", err)
	}
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(false) {
		t.Fatalf("incorrect OrchestratorVersion '%s'", cs.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	vlabscs = &vlabs.ContainerService{
		Properties: &vlabs.Properties{
			OrchestratorProfile: &vlabs.OrchestratorProfile{
				OrchestratorType:    vlabs.Kubernetes,
				OrchestratorVersion: "1.10.13",
			},
		},
	}
	cs, err = ConvertVLabsContainerService(vlabscs, false)
	if err != nil {
		t.Fatalf("Failed to convert ContainerService, error: %s", err)
	}
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != "1.10.13" {
		t.Fatalf("incorrect OrchestratorVersion '%s'", cs.Properties.OrchestratorProfile.OrchestratorVersion)
	}
}

func TestKubernetesOrchestratorVersionFailWhenInvalid(t *testing.T) {
	vlabscs := &vlabs.ContainerService{
		Properties: &vlabs.Properties{
			OrchestratorProfile: &vlabs.OrchestratorProfile{
				OrchestratorType:    vlabs.Kubernetes,
				OrchestratorVersion: "1.10.8",
			},
		},
	}

	_, err := ConvertVLabsContainerService(vlabscs, false)
	if err == nil {
		t.Error("1.10.8 is not a valid version and should fail, but didn't")
	}

	vlabscs.Properties.OrchestratorProfile.OrchestratorRelease = "1.9"
	vlabscs.Properties.OrchestratorProfile.OrchestratorVersion = "1.10.7"
	_, err = ConvertVLabsContainerService(vlabscs, false)
	if err == nil {
		t.Fatalf("release 1.9 is incoherent with 1.10.7 and should fail, but didn't")
	}

	vlabscs.Properties.OrchestratorProfile.OrchestratorVersion = "whatever"
	vlabscs.Properties.OrchestratorProfile.OrchestratorRelease = "1.10.8"

	_, err = ConvertVLabsContainerService(vlabscs, false)
	if err == nil {
		t.Fatalf("garbage version string should fail, but didn't")
	}

}

func TestConvertVLabsKubernetesConfigProfile(t *testing.T) {
	tests := map[string]struct {
		props  *vlabs.KubernetesConfig
		expect *KubernetesConfig
	}{
		"WindowsNodeBinariesURL": {
			props: &vlabs.KubernetesConfig{
				WindowsNodeBinariesURL: "http://test/test.tar.gz",
			},
			expect: &KubernetesConfig{
				WindowsNodeBinariesURL: "http://test/test.tar.gz",
			},
		},
	}

	for name, test := range tests {
		t.Logf("running scenario %q", name)
		actual := &KubernetesConfig{}
		convertVLabsKubernetesConfig(test.props, actual)
		if !equality.Semantic.DeepEqual(test.expect, actual) {
			t.Errorf(spew.Sprintf("Expected:\n%+v\nGot:\n%+v", test.expect, actual))
		}
	}
}

func TestConvertCustomFilesToAPI(t *testing.T) {
	expectedAPICustomFiles := []CustomFile{
		{
			Source: "/test/source",
			Dest:   "/test/dest",
		},
	}
	masterProfile := MasterProfile{}

	vp := &vlabs.MasterProfile{}
	vp.CustomFiles = &[]vlabs.CustomFile{
		{
			Source: "/test/source",
			Dest:   "/test/dest",
		},
	}
	convertCustomFilesToAPI(vp, &masterProfile)
	if !equality.Semantic.DeepEqual(&expectedAPICustomFiles, masterProfile.CustomFiles) {
		t.Fatalf("convertCustomFilesToApi conversion of vlabs.MasterProfile did not convert correctly")
	}
}

func TestCustomCloudProfile(t *testing.T) {
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

	vlabscs := &vlabs.ContainerService{
		Properties: &vlabs.Properties{
			CustomCloudProfile: &vlabs.CustomCloudProfile{
				IdentitySystem:       ADFSIdentitySystem,
				AuthenticationMethod: ClientCertificateAuthMethod,
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

	cs, err := ConvertVLabsContainerService(vlabscs, false)
	if err != nil {
		t.Fatalf("Failed to convert ContainerService, error: %s", err)
	}
	if cs.Properties.CustomCloudProfile.AuthenticationMethod != ClientCertificateAuthMethod {
		t.Errorf("incorrect AuthenticationMethod, expect: '%s', actual: '%s'", ClientCertificateAuthMethod, cs.Properties.CustomCloudProfile.AuthenticationMethod)
	}
	if cs.Properties.CustomCloudProfile.IdentitySystem != ADFSIdentitySystem {
		t.Errorf("incorrect IdentitySystem, expect: '%s', actual: '%s'", ADFSIdentitySystem, cs.Properties.CustomCloudProfile.IdentitySystem)
	}
	if cs.Properties.CustomCloudProfile.Environment.Name != name {
		t.Errorf("incorrect Name, expect: '%s', actual: '%s'", name, cs.Properties.CustomCloudProfile.Environment.Name)
	}
	if cs.Properties.CustomCloudProfile.Environment.ManagementPortalURL != managementPortalURL {
		t.Errorf("incorrect ManagementPortalURL, expect: '%s', actual: '%s'", managementPortalURL, cs.Properties.CustomCloudProfile.Environment.ManagementPortalURL)
	}
	if cs.Properties.CustomCloudProfile.Environment.PublishSettingsURL != publishSettingsURL {
		t.Errorf("incorrect PublishSettingsURL, expect: '%s', actual: '%s'", publishSettingsURL, cs.Properties.CustomCloudProfile.Environment.PublishSettingsURL)
	}
	if cs.Properties.CustomCloudProfile.Environment.ServiceManagementEndpoint != serviceManagementEndpoint {
		t.Errorf("incorrect ServiceManagementEndpoint, expect: '%s', actual: '%s'", serviceManagementEndpoint, cs.Properties.CustomCloudProfile.Environment.ServiceManagementEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.ResourceManagerEndpoint != resourceManagerEndpoint {
		t.Errorf("incorrect ResourceManagerEndpoint, expect: '%s', actual: '%s'", resourceManagerEndpoint, cs.Properties.CustomCloudProfile.Environment.ResourceManagerEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.ActiveDirectoryEndpoint != activeDirectoryEndpoint {
		t.Errorf("incorrect ActiveDirectoryEndpoint, expect: '%s', actual: '%s'", activeDirectoryEndpoint, cs.Properties.CustomCloudProfile.Environment.ActiveDirectoryEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.GalleryEndpoint != galleryEndpoint {
		t.Errorf("incorrect GalleryEndpoint, expect: '%s', actual: '%s'", galleryEndpoint, cs.Properties.CustomCloudProfile.Environment.GalleryEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.KeyVaultEndpoint != keyVaultEndpoint {
		t.Errorf("incorrect KeyVaultEndpoint, expect: '%s', actual: '%s'", keyVaultEndpoint, cs.Properties.CustomCloudProfile.Environment.KeyVaultEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.GraphEndpoint != graphEndpoint {
		t.Errorf("incorrect GraphEndpoint, expect: '%s', actual: '%s'", graphEndpoint, cs.Properties.CustomCloudProfile.Environment.GraphEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.ServiceBusEndpoint != serviceBusEndpoint {
		t.Errorf("incorrect ServiceBusEndpoint, expect: '%s', actual: '%s'", serviceBusEndpoint, cs.Properties.CustomCloudProfile.Environment.ServiceBusEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.BatchManagementEndpoint != batchManagementEndpoint {
		t.Errorf("incorrect BatchManagementEndpoint, expect: '%s', actual: '%s'", batchManagementEndpoint, cs.Properties.CustomCloudProfile.Environment.BatchManagementEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Environment.StorageEndpointSuffix != storageEndpointSuffix {
		t.Errorf("incorrect StorageEndpointSuffix, expect: '%s', actual: '%s'", storageEndpointSuffix, cs.Properties.CustomCloudProfile.Environment.StorageEndpointSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.SQLDatabaseDNSSuffix != sqlDatabaseDNSSuffix {
		t.Errorf("incorrect SQLDatabaseDNSSuffix, expect: '%s', actual: '%s'", sqlDatabaseDNSSuffix, cs.Properties.CustomCloudProfile.Environment.SQLDatabaseDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.TrafficManagerDNSSuffix != trafficManagerDNSSuffix {
		t.Errorf("incorrect TrafficManagerDNSSuffix, expect: '%s', actual: '%s'", trafficManagerDNSSuffix, cs.Properties.CustomCloudProfile.Environment.TrafficManagerDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.KeyVaultDNSSuffix != keyVaultDNSSuffix {
		t.Errorf("incorrect KeyVaultDNSSuffix, expect: '%s', actual: '%s'", keyVaultDNSSuffix, cs.Properties.CustomCloudProfile.Environment.KeyVaultDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.ServiceBusEndpointSuffix != serviceBusEndpointSuffix {
		t.Errorf("incorrect ServiceBusEndpointSuffix, expect: '%s', actual: '%s'", serviceBusEndpointSuffix, cs.Properties.CustomCloudProfile.Environment.ServiceBusEndpointSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.ServiceManagementVMDNSSuffix != serviceManagementVMDNSSuffix {
		t.Errorf("incorrect ServiceManagementVMDNSSuffix, expect: '%s', actual: '%s'", serviceManagementVMDNSSuffix, cs.Properties.CustomCloudProfile.Environment.ServiceManagementVMDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix != resourceManagerVMDNSSuffix {
		t.Errorf("incorrect ResourceManagerVMDNSSuffix, expect: '%s', actual: '%s'", resourceManagerVMDNSSuffix, cs.Properties.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.ContainerRegistryDNSSuffix != containerRegistryDNSSuffix {
		t.Errorf("incorrect ContainerRegistryDNSSuffix, expect: '%s', actual: '%s'", containerRegistryDNSSuffix, cs.Properties.CustomCloudProfile.Environment.ContainerRegistryDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Environment.TokenAudience != tokenAudience {
		t.Errorf("incorrect TokenAudience, expect: '%s', actual: '%s'", tokenAudience, cs.Properties.CustomCloudProfile.Environment.TokenAudience)
	}
}

func TestConvertAzureEnvironmentSpecConfig(t *testing.T) {
	//Mock AzureEnvironmentSpecConfig
	vlabscs := &vlabs.ContainerService{
		Properties: &vlabs.Properties{
			CustomCloudProfile: &vlabs.CustomCloudProfile{
				IdentitySystem:       AzureADIdentitySystem,
				AuthenticationMethod: ClientSecretAuthMethod,
				AzureEnvironmentSpecConfig: &vlabs.AzureEnvironmentSpecConfig{
					CloudName: "AzureStackCloud",
					//DockerSpecConfig specify the docker engine download repo
					DockerSpecConfig: vlabs.DockerSpecConfig{
						DockerEngineRepo:         "DockerEngineRepo",
						DockerComposeDownloadURL: "DockerComposeDownloadURL",
					},
					//KubernetesSpecConfig - Due to Chinese firewall issue, the default containers from google is blocked, use the Chinese local mirror instead
					KubernetesSpecConfig: vlabs.KubernetesSpecConfig{
						AzureTelemetryPID:                "AzureTelemetryPID",
						KubernetesImageBase:              "KubernetesImageBase",
						MCRKubernetesImageBase:           "MCRKubernetesImageBase",
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
					},
					DCOSSpecConfig: vlabs.DCOSSpecConfig{
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
					EndpointConfig: vlabs.AzureEndpointConfig{
						ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
					},
					OSImageConfig: map[vlabs.Distro]vlabs.AzureOSImageConfig{
						vlabs.Distro("Test"): {
							ImageOffer:     "ImageOffer",
							ImageSku:       "ImageSku",
							ImagePublisher: "ImagePublisher",
							ImageVersion:   "ImageVersion",
						},
					},
				},
			},
		},
	}

	cs, err := ConvertVLabsContainerService(vlabscs, false)
	if err != nil {
		t.Fatalf("Failed to convert ContainerService, error: %s", err)
	}

	if cs.Properties.CustomCloudProfile.AuthenticationMethod != ClientSecretAuthMethod {
		t.Errorf("incorrect AuthenticationMethod, expect: '%s', actual: '%s'", ClientSecretAuthMethod, cs.Properties.CustomCloudProfile.AuthenticationMethod)
	}
	if cs.Properties.CustomCloudProfile.IdentitySystem != AzureADIdentitySystem {
		t.Errorf("incorrect IdentitySystem, expect: '%s', actual: '%s'", AzureADIdentitySystem, cs.Properties.CustomCloudProfile.IdentitySystem)
	}

	csSpec := cs.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig
	vlabscsSpec := vlabscs.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig
	if csSpec.CloudName != vlabscsSpec.CloudName {
		t.Errorf("incorrect CloudName, expect: '%s', actual: '%s'", vlabscsSpec.CloudName, csSpec.CloudName)
	}

	//KubernetesSpecConfig
	if csSpec.KubernetesSpecConfig.AzureTelemetryPID != vlabscsSpec.KubernetesSpecConfig.AzureTelemetryPID {
		t.Errorf("incorrect AzureTelemetryPID, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.AzureTelemetryPID, csSpec.KubernetesSpecConfig.AzureTelemetryPID)
	}
	if csSpec.KubernetesSpecConfig.KubernetesImageBase != vlabscsSpec.KubernetesSpecConfig.KubernetesImageBase {
		t.Errorf("incorrect KubernetesImageBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.KubernetesImageBase, csSpec.KubernetesSpecConfig.KubernetesImageBase)
	}
	if csSpec.KubernetesSpecConfig.TillerImageBase != vlabscsSpec.KubernetesSpecConfig.TillerImageBase {
		t.Errorf("incorrect TillerImageBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.TillerImageBase, csSpec.KubernetesSpecConfig.TillerImageBase)
	}
	if csSpec.KubernetesSpecConfig.ACIConnectorImageBase != vlabscsSpec.KubernetesSpecConfig.ACIConnectorImageBase {
		t.Errorf("incorrect ACIConnectorImageBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.ACIConnectorImageBase, csSpec.KubernetesSpecConfig.ACIConnectorImageBase)
	}
	if csSpec.KubernetesSpecConfig.NVIDIAImageBase != vlabscsSpec.KubernetesSpecConfig.NVIDIAImageBase {
		t.Errorf("incorrect NVIDIAImageBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.NVIDIAImageBase, csSpec.KubernetesSpecConfig.NVIDIAImageBase)
	}
	if csSpec.KubernetesSpecConfig.AzureCNIImageBase != vlabscsSpec.KubernetesSpecConfig.AzureCNIImageBase {
		t.Errorf("incorrect AzureCNIImageBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.AzureCNIImageBase, csSpec.KubernetesSpecConfig.AzureCNIImageBase)
	}
	if csSpec.KubernetesSpecConfig.EtcdDownloadURLBase != vlabscsSpec.KubernetesSpecConfig.EtcdDownloadURLBase {
		t.Errorf("incorrect EtcdDownloadURLBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.EtcdDownloadURLBase, csSpec.KubernetesSpecConfig.EtcdDownloadURLBase)
	}
	if csSpec.KubernetesSpecConfig.KubeBinariesSASURLBase != vlabscsSpec.KubernetesSpecConfig.KubeBinariesSASURLBase {
		t.Errorf("incorrect KubeBinariesSASURLBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.KubeBinariesSASURLBase, csSpec.KubernetesSpecConfig.KubeBinariesSASURLBase)
	}
	if csSpec.KubernetesSpecConfig.WindowsTelemetryGUID != vlabscsSpec.KubernetesSpecConfig.WindowsTelemetryGUID {
		t.Errorf("incorrect WindowsTelemetryGUID, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.WindowsTelemetryGUID, csSpec.KubernetesSpecConfig.WindowsTelemetryGUID)
	}
	if csSpec.KubernetesSpecConfig.CNIPluginsDownloadURL != vlabscsSpec.KubernetesSpecConfig.CNIPluginsDownloadURL {
		t.Errorf("incorrect CNIPluginsDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.CNIPluginsDownloadURL, csSpec.KubernetesSpecConfig.CNIPluginsDownloadURL)
	}
	if csSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL != vlabscsSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL {
		t.Errorf("incorrect VnetCNILinuxPluginsDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL, csSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL)
	}
	if csSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL != vlabscsSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL {
		t.Errorf("incorrect VnetCNIWindowsPluginsDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL, csSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL)
	}
	if csSpec.KubernetesSpecConfig.ContainerdDownloadURLBase != vlabscsSpec.KubernetesSpecConfig.ContainerdDownloadURLBase {
		t.Errorf("incorrect ContainerdDownloadURLBase, expect: '%s', actual: '%s'", vlabscsSpec.KubernetesSpecConfig.ContainerdDownloadURLBase, csSpec.KubernetesSpecConfig.ContainerdDownloadURLBase)
	}

	//DockerSpecConfig
	if csSpec.DockerSpecConfig.DockerComposeDownloadURL != vlabscsSpec.DockerSpecConfig.DockerComposeDownloadURL {
		t.Errorf("incorrect DockerComposeDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.DockerSpecConfig.DockerComposeDownloadURL, csSpec.DockerSpecConfig.DockerComposeDownloadURL)
	}
	if csSpec.DockerSpecConfig.DockerEngineRepo != vlabscsSpec.DockerSpecConfig.DockerEngineRepo {
		t.Errorf("incorrect DockerEngineRepo, expect: '%s', actual: '%s'", vlabscsSpec.DockerSpecConfig.DockerEngineRepo, csSpec.DockerSpecConfig.DockerEngineRepo)
	}

	//DCOSSpecConfig
	if csSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL != vlabscsSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL {
		t.Errorf("incorrect DCOS188BootstrapDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL, csSpec.DCOSSpecConfig.DCOS188BootstrapDownloadURL)
	}
	if csSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL != vlabscsSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL {
		t.Errorf("incorrect DCOS190BootstrapDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL, csSpec.DCOSSpecConfig.DCOS190BootstrapDownloadURL)
	}
	if csSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL != vlabscsSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL {
		t.Errorf("incorrect DCOS198BootstrapDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL, csSpec.DCOSSpecConfig.DCOS198BootstrapDownloadURL)
	}
	if csSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL != vlabscsSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL {
		t.Errorf("incorrect DCOS110BootstrapDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL, csSpec.DCOSSpecConfig.DCOS110BootstrapDownloadURL)
	}
	if csSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL != vlabscsSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL {
		t.Errorf("incorrect DCOS111BootstrapDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL, csSpec.DCOSSpecConfig.DCOS111BootstrapDownloadURL)
	}
	if csSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL != vlabscsSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL {
		t.Errorf("incorrect DCOSWindowsBootstrapDownloadURL, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL, csSpec.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL)
	}
	if csSpec.DCOSSpecConfig.DcosRepositoryURL != vlabscsSpec.DCOSSpecConfig.DcosRepositoryURL {
		t.Errorf("incorrect DcosRepositoryURL, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DcosRepositoryURL, csSpec.DCOSSpecConfig.DcosRepositoryURL)
	}
	if csSpec.DCOSSpecConfig.DcosClusterPackageListID != vlabscsSpec.DCOSSpecConfig.DcosClusterPackageListID {
		t.Errorf("incorrect DcosClusterPackageListID, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DcosClusterPackageListID, csSpec.DCOSSpecConfig.DcosClusterPackageListID)
	}
	if csSpec.DCOSSpecConfig.DcosProviderPackageID != vlabscsSpec.DCOSSpecConfig.DcosProviderPackageID {
		t.Errorf("incorrect DcosProviderPackageID, expect: '%s', actual: '%s'", vlabscsSpec.DCOSSpecConfig.DcosProviderPackageID, csSpec.DCOSSpecConfig.DcosProviderPackageID)
	}

	//EndpointConfig
	if csSpec.EndpointConfig.ResourceManagerVMDNSSuffix != vlabscsSpec.EndpointConfig.ResourceManagerVMDNSSuffix {
		t.Errorf("incorrect ResourceManagerVMDNSSuffix, expect: '%s', actual: '%s'", vlabscsSpec.EndpointConfig.ResourceManagerVMDNSSuffix, csSpec.EndpointConfig.ResourceManagerVMDNSSuffix)
	}

	//OSImageConfig
	for k, v := range vlabscsSpec.OSImageConfig {
		if actualValue, ok := csSpec.OSImageConfig[Distro(string(k))]; ok {
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

func TestConvertV20160930ContainerService(t *testing.T) {
	storageURL, _ := url.Parse("https://bing.com")
	v20160930CS := &v20160930.ContainerService{
		Location: "westus2",
		Plan: &v20160930.ResourcePurchasePlan{
			Name:          "fooPlan",
			PromotionCode: "fooPromoCode",
			Product:       "fooProduct",
			Publisher:     "fooPublisher",
		},

		Tags: map[string]string{
			"foo": "bar",
		},
		Properties: &v20160930.Properties{
			ProvisioningState: v20160930.Succeeded,
			OrchestratorProfile: &v20160930.OrchestratorProfile{
				OrchestratorType: DCOS,
			},
			WindowsProfile: &v20160930.WindowsProfile{
				AdminUsername: "sampleAdminUsername",
				AdminPassword: "sampleAdminPassword",
			},
			DiagnosticsProfile: &v20160930.DiagnosticsProfile{
				VMDiagnostics: &v20160930.VMDiagnostics{
					Enabled:    true,
					StorageURL: storageURL,
				},
			},
			JumpboxProfile: &v20160930.JumpboxProfile{
				OSType:    "Linux",
				DNSPrefix: "blueorange",
				FQDN:      "blueorange.westus2.azureapp.com",
			},
			AgentPoolProfiles: []*v20160930.AgentPoolProfile{
				{
					Name:      "sampleagent",
					Count:     2,
					VMSize:    "Standard_DS1_v1",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.azureapp.com",
					OSType:    "Linux",
				},
			},
			MasterProfile: &v20160930.MasterProfile{
				Count: 1,
			},
			CustomProfile: &v20160930.CustomProfile{
				Orchestrator: DCOS,
			},
		},
	}

	apiCs := ConvertV20160930ContainerService(v20160930CS)

	if apiCs == nil {
		t.Error("unexpected nil output while executing ConvertV20160930ContainerService")
	}
}

func TestConvertV20170701ContainerService(t *testing.T) {
	v20170701CS := &v20170701.ContainerService{
		Location: "westus2",
		Plan: &v20170701.ResourcePurchasePlan{
			Name:          "fooPlan",
			PromotionCode: "fooPromoCode",
			Product:       "fooProduct",
			Publisher:     "fooPublisher",
		},

		Tags: map[string]string{
			"foo": "bar",
		},
		Properties: &v20170701.Properties{
			ProvisioningState: v20170701.Succeeded,
			OrchestratorProfile: &v20170701.OrchestratorProfile{
				OrchestratorType: DCOS,
			},
			WindowsProfile: &v20170701.WindowsProfile{
				AdminUsername: "sampleAdminUsername",
				AdminPassword: "sampleAdminPassword",
			},
			AgentPoolProfiles: []*v20170701.AgentPoolProfile{
				{
					Name:      "sampleagent",
					Count:     2,
					VMSize:    "Standard_DS1_v1",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.azureapp.com",
					OSType:    "Linux",
				},
			},
			MasterProfile: &v20170701.MasterProfile{
				Count: 1,
			},
			CustomProfile: &v20170701.CustomProfile{
				Orchestrator: DCOS,
			},
		},
	}

	apiCs := ConvertV20170701ContainerService(v20170701CS, false)

	if apiCs == nil {
		t.Error("unexpected nil output while executing ConvertV20170701ContainerService")
	}
}

func TestConvertV20160330ContainerService(t *testing.T) {
	storageURL, _ := url.Parse("https://bing.com")
	v20160330CS := &v20160330.ContainerService{
		Location: "westus2",
		Plan: &v20160330.ResourcePurchasePlan{
			Name:          "fooPlan",
			PromotionCode: "fooPromoCode",
			Product:       "fooProduct",
			Publisher:     "fooPublisher",
		},

		Tags: map[string]string{
			"foo": "bar",
		},
		Properties: &v20160330.Properties{
			ProvisioningState: v20160330.Succeeded,
			OrchestratorProfile: &v20160330.OrchestratorProfile{
				OrchestratorType: DCOS,
			},
			WindowsProfile: &v20160330.WindowsProfile{
				AdminUsername: "sampleAdminUsername",
				AdminPassword: "sampleAdminPassword",
			},
			DiagnosticsProfile: &v20160330.DiagnosticsProfile{
				VMDiagnostics: &v20160330.VMDiagnostics{
					Enabled:    true,
					StorageURL: storageURL,
				},
			},
			JumpboxProfile: &v20160330.JumpboxProfile{
				OSType:    "Linux",
				DNSPrefix: "blueorange",
				FQDN:      "blueorange.westus2.azureapp.com",
			},
			AgentPoolProfiles: []*v20160330.AgentPoolProfile{
				{
					Name:      "sampleagent",
					Count:     2,
					VMSize:    "Standard_DS1_v1",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.azureapp.com",
					OSType:    "Linux",
				},
			},
			MasterProfile: &v20160330.MasterProfile{
				Count: 1,
			},
		},
	}

	apiCs := ConvertV20160330ContainerService(v20160330CS)

	if apiCs == nil {
		t.Error("unexpected nil output while executing ConvertV20160330ContainerService")
	}
}

func TestConvertV20170131ContainerService(t *testing.T) {
	storageURL, _ := url.Parse("https://bing.com")
	v20170131CS := &v20170131.ContainerService{
		Location: "westus2",
		Plan: &v20170131.ResourcePurchasePlan{
			Name:          "fooPlan",
			PromotionCode: "fooPromoCode",
			Product:       "fooProduct",
			Publisher:     "fooPublisher",
		},

		Tags: map[string]string{
			"foo": "bar",
		},
		Properties: &v20170131.Properties{
			ProvisioningState: v20170131.Succeeded,
			OrchestratorProfile: &v20170131.OrchestratorProfile{
				OrchestratorType: DCOS,
			},
			WindowsProfile: &v20170131.WindowsProfile{
				AdminUsername: "sampleAdminUsername",
				AdminPassword: "sampleAdminPassword",
			},
			DiagnosticsProfile: &v20170131.DiagnosticsProfile{
				VMDiagnostics: &v20170131.VMDiagnostics{
					Enabled:    true,
					StorageURL: storageURL,
				},
			},
			JumpboxProfile: &v20170131.JumpboxProfile{
				OSType:    "Linux",
				DNSPrefix: "blueorange",
				FQDN:      "blueorange.westus2.azureapp.com",
			},
			AgentPoolProfiles: []*v20170131.AgentPoolProfile{
				{
					Name:      "sampleagent",
					Count:     2,
					VMSize:    "Standard_DS1_v1",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.azureapp.com",
					OSType:    "Linux",
				},
			},
			MasterProfile: &v20170131.MasterProfile{
				Count: 1,
			},
			CustomProfile: &v20170131.CustomProfile{
				Orchestrator: DCOS,
			},
		},
	}

	apiCs := ConvertV20170131ContainerService(v20170131CS)
	if apiCs == nil {
		t.Error("unexpected nil output while executing ConvertV20170131ContainerService")
	}
}

func TestConvertVLabsContainerService(t *testing.T) {
	vlabsCS := &vlabs.ContainerService{
		Location: "westus2",
		Plan: &vlabs.ResourcePurchasePlan{
			Name:          "fooPlan",
			PromotionCode: "fooPromoCode",
			Product:       "fooProduct",
			Publisher:     "fooPublisher",
		},

		Tags: map[string]string{
			"foo": "bar",
		},
		Properties: &vlabs.Properties{
			ProvisioningState: vlabs.Succeeded,
			OrchestratorProfile: &vlabs.OrchestratorProfile{
				OrchestratorType: DCOS,
				DcosConfig: &vlabs.DcosConfig{
					DcosBootstrapURL:         "SampleDcosBootstrapURL",
					DcosWindowsBootstrapURL:  "SampleWindowsDcosBootstrapURL",
					Registry:                 "SampleRegistry",
					RegistryPass:             "SampleRegistryPass",
					RegistryUser:             "SampleRegistryUser",
					DcosClusterPackageListID: "SampleDcosClusterPackageListID",
					DcosProviderPackageID:    "SampleDcosProviderPackageID",
					BootstrapProfile: &vlabs.BootstrapProfile{
						VMSize:       "Standard_Ds1_v1",
						OSDiskSizeGB: 256,
						OAuthEnabled: true,
						StaticIP:     "172.0.0.1",
						Subnet:       "255.255.255.0",
					},
				},
			},
			WindowsProfile: &vlabs.WindowsProfile{
				AdminUsername: "sampleAdminUsername",
				AdminPassword: "sampleAdminPassword",
			},
			AgentPoolProfiles: []*vlabs.AgentPoolProfile{
				{
					Name:      "sampleagent",
					Count:     2,
					VMSize:    "Standard_DS1_v1",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.azureapp.com",
					OSType:    "Linux",
				},
				{
					Name:      "sampleAgent-public",
					Count:     2,
					VMSize:    "sampleVM",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.com",
					OSType:    "Linux",
					ImageRef: &vlabs.ImageReference{
						Name:           "testImage",
						ResourceGroup:  "testRg",
						SubscriptionID: "testSub",
						Gallery:        "testGallery",
						Version:        "0.0.1",
					},
				},
			},
			MasterProfile: &vlabs.MasterProfile{
				Count: 1,
				PreProvisionExtension: &vlabs.Extension{
					Name:        "fooExtension",
					SingleOrAll: "All",
					Template:    "{{foobar}}",
				},
				ImageRef: &vlabs.ImageReference{
					Name:          "FooImageRef",
					ResourceGroup: "FooImageRefResourceGroup",
				},
				Extensions: []vlabs.Extension{
					{
						Name:        "sampleExtension",
						SingleOrAll: "single",
						Template:    "{{foobar}}",
					},
				},
			},
			CertificateProfile: &vlabs.CertificateProfile{
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
			FeatureFlags: &vlabs.FeatureFlags{
				EnableCSERunInBackground: true,
				BlockOutboundInternet:    false,
				EnableTelemetry:          false,
			},
			AADProfile: &vlabs.AADProfile{
				ClientAppID:  "SampleClientAppID",
				ServerAppID:  "ServerAppID",
				TenantID:     "SampleTenantID",
				AdminGroupID: "SampleAdminGroupID",
			},
			ExtensionProfiles: []*vlabs.ExtensionProfile{
				{
					Name:                "fooExtension",
					Version:             "fooVersion",
					ExtensionParameters: "fooExtensionParameters",
					ExtensionParametersKeyVaultRef: &vlabs.KeyvaultSecretRef{
						VaultID:       "fooVaultID",
						SecretName:    "fooSecretName",
						SecretVersion: "fooSecretVersion",
					},
					RootURL:  "fooRootURL",
					Script:   "fooSsript",
					URLQuery: "fooURL",
				},
			},
			LinuxProfile: &vlabs.LinuxProfile{
				AdminUsername: "azureuser",
				Secrets: []vlabs.KeyVaultSecrets{
					{
						SourceVault: &vlabs.KeyVaultID{
							ID: "sampleKeyVaultID",
						},
						VaultCertificates: []vlabs.KeyVaultCertificate{
							{
								CertificateURL:   "FooCertURL",
								CertificateStore: "BarCertStore",
							},
						},
					},
				},
				CustomNodesDNS: &vlabs.CustomNodesDNS{
					DNSServer: "SampleDNSServer",
				},
				CustomSearchDomain: &vlabs.CustomSearchDomain{
					Name:          "FooCustomSearchDomain",
					RealmUser:     "sampleRealmUser",
					RealmPassword: "sampleRealmPassword",
				},
			},
		},
	}

	apiCs, err := ConvertVLabsContainerService(vlabsCS, false)
	if apiCs == nil {
		t.Error("unexpected nil output while executing ConvertVLabsContainerService")
	}

	if err != nil {
		t.Errorf("unexpected error while executing ConvertVLabsContainerService: %s", err.Error())
	}

	//Test Vlabs with Kubernetes Orchestrator
	vlabsCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	vlabsCS.Properties.OrchestratorProfile.DcosConfig = nil
	vlabsCS.Properties.OrchestratorProfile.KubernetesConfig = &vlabs.KubernetesConfig{
		Addons: []vlabs.KubernetesAddon{
			{
				Name:    "sampleAddon",
				Enabled: to.BoolPtr(true),
				Containers: []vlabs.KubernetesContainerSpec{
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
		PrivateCluster: &vlabs.PrivateCluster{
			Enabled: to.BoolPtr(true),
			JumpboxProfile: &vlabs.PrivateJumpboxProfile{
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
	}

	apiCs, err = ConvertVLabsContainerService(vlabsCS, false)
	if apiCs == nil {
		t.Error("unexpected nil output while executing ConvertVLabsContainerService")
	}

	if err != nil {
		t.Errorf("unexpected error while executing ConvertVLabsContainerService: %s", err.Error())
	}

}

func TestTelemetryEnabled(t *testing.T) {
	vlabsCS := &vlabs.ContainerService{
		Location: "westus2",
		Plan: &vlabs.ResourcePurchasePlan{
			Name:          "fooPlan",
			PromotionCode: "fooPromoCode",
			Product:       "fooProduct",
			Publisher:     "fooPublisher",
		},

		Tags: map[string]string{
			"foo": "bar",
		},
		Properties: &vlabs.Properties{
			ProvisioningState: vlabs.Succeeded,
			OrchestratorProfile: &vlabs.OrchestratorProfile{
				OrchestratorType: DCOS,
				DcosConfig: &vlabs.DcosConfig{
					DcosBootstrapURL:         "SampleDcosBootstrapURL",
					DcosWindowsBootstrapURL:  "SampleWindowsDcosBootstrapURL",
					Registry:                 "SampleRegistry",
					RegistryPass:             "SampleRegistryPass",
					RegistryUser:             "SampleRegistryUser",
					DcosClusterPackageListID: "SampleDcosClusterPackageListID",
					DcosProviderPackageID:    "SampleDcosProviderPackageID",
					BootstrapProfile: &vlabs.BootstrapProfile{
						VMSize:       "Standard_Ds1_v1",
						OSDiskSizeGB: 256,
						OAuthEnabled: true,
						StaticIP:     "172.0.0.1",
						Subnet:       "255.255.255.0",
					},
				},
			},
			WindowsProfile: &vlabs.WindowsProfile{
				AdminUsername: "sampleAdminUsername",
				AdminPassword: "sampleAdminPassword",
			},
			AgentPoolProfiles: []*vlabs.AgentPoolProfile{
				{
					Name:      "sampleagent",
					Count:     2,
					VMSize:    "Standard_DS1_v1",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.azureapp.com",
					OSType:    "Linux",
				},
				{
					Name:      "sampleAgent-public",
					Count:     2,
					VMSize:    "sampleVM",
					DNSPrefix: "blueorange",
					FQDN:      "blueorange.westus2.com",
					OSType:    "Linux",
					ImageRef: &vlabs.ImageReference{
						Name:           "testImage",
						ResourceGroup:  "testRg",
						SubscriptionID: "testSub",
						Gallery:        "testGallery",
						Version:        "0.0.1",
					},
				},
			},
			MasterProfile: &vlabs.MasterProfile{
				Count: 1,
				PreProvisionExtension: &vlabs.Extension{
					Name:        "fooExtension",
					SingleOrAll: "All",
					Template:    "{{foobar}}",
				},
				ImageRef: &vlabs.ImageReference{
					Name:          "FooImageRef",
					ResourceGroup: "FooImageRefResourceGroup",
				},
				Extensions: []vlabs.Extension{
					{
						Name:        "sampleExtension",
						SingleOrAll: "single",
						Template:    "{{foobar}}",
					},
				},
			},
			CertificateProfile: &vlabs.CertificateProfile{
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
			FeatureFlags: &vlabs.FeatureFlags{
				EnableCSERunInBackground: true,
				BlockOutboundInternet:    false,
				EnableTelemetry:          true,
			},
			AADProfile: &vlabs.AADProfile{
				ClientAppID:  "SampleClientAppID",
				ServerAppID:  "ServerAppID",
				TenantID:     "SampleTenantID",
				AdminGroupID: "SampleAdminGroupID",
			},
			ExtensionProfiles: []*vlabs.ExtensionProfile{
				{
					Name:                "fooExtension",
					Version:             "fooVersion",
					ExtensionParameters: "fooExtensionParameters",
					ExtensionParametersKeyVaultRef: &vlabs.KeyvaultSecretRef{
						VaultID:       "fooVaultID",
						SecretName:    "fooSecretName",
						SecretVersion: "fooSecretVersion",
					},
					RootURL:  "fooRootURL",
					Script:   "fooSsript",
					URLQuery: "fooURL",
				},
			},
			LinuxProfile: &vlabs.LinuxProfile{
				AdminUsername: "azureuser",
				Secrets: []vlabs.KeyVaultSecrets{
					{
						SourceVault: &vlabs.KeyVaultID{
							ID: "sampleKeyVaultID",
						},
						VaultCertificates: []vlabs.KeyVaultCertificate{
							{
								CertificateURL:   "FooCertURL",
								CertificateStore: "BarCertStore",
							},
						},
					},
				},
				CustomNodesDNS: &vlabs.CustomNodesDNS{
					DNSServer: "SampleDNSServer",
				},
				CustomSearchDomain: &vlabs.CustomSearchDomain{
					Name:          "FooCustomSearchDomain",
					RealmUser:     "sampleRealmUser",
					RealmPassword: "sampleRealmPassword",
				},
			},
		},
	}

	apiCs, err := ConvertVLabsContainerService(vlabsCS, false)
	if apiCs == nil {
		t.Error("unexpected nil output while executing ConvertVLabsContainerService")
	}

	if err != nil {
		t.Errorf("unexpected error while executing ConvertVLabsContainerService: %s", err.Error())
	}

	if !vlabsCS.Properties.FeatureFlags.EnableTelemetry {
		t.Error("unexpected false output while checking for EnableTelemetry")
	}
}
func TestConvertVLabsWindowsProfile(t *testing.T) {
	falseVar := false

	cases := []struct {
		name     string
		w        vlabs.WindowsProfile
		expected WindowsProfile
	}{
		{
			name: "empty profile",
			w:    vlabs.WindowsProfile{},
			expected: WindowsProfile{
				Secrets: []KeyVaultSecrets{},
			},
		},
		{
			name: "misc fields",
			w: vlabs.WindowsProfile{
				AdminUsername:          "user",
				AdminPassword:          "password",
				EnableAutomaticUpdates: &falseVar,
				ImageVersion:           "17763.615.1907121548",
				SSHEnabled:             false,
				WindowsPublisher:       "MicrosoftWindowsServer",
				WindowsOffer:           "WindowsServer",
				WindowsSku:             "2019-Datacenter-Core-smalldisk",
				WindowsDockerVersion:   "18.09",
			},
			expected: WindowsProfile{
				AdminUsername:          "user",
				AdminPassword:          "password",
				EnableAutomaticUpdates: &falseVar,
				ImageVersion:           "17763.615.1907121548",
				SSHEnabled:             false,
				WindowsPublisher:       "MicrosoftWindowsServer",
				WindowsOffer:           "WindowsServer",
				WindowsSku:             "2019-Datacenter-Core-smalldisk",
				WindowsDockerVersion:   "18.09",
				Secrets:                []KeyVaultSecrets{},
			},
		},
		{
			name: "image reference",
			w: vlabs.WindowsProfile{
				ImageRef: &vlabs.ImageReference{
					Gallery:        "gallery",
					Name:           "name",
					ResourceGroup:  "rg",
					SubscriptionID: "dc6bd10c-110c-4134-88c5-4d5a039129c4",
					Version:        "1.25.6",
				},
			},
			expected: WindowsProfile{
				ImageRef: &ImageReference{
					Gallery:        "gallery",
					Name:           "name",
					ResourceGroup:  "rg",
					SubscriptionID: "dc6bd10c-110c-4134-88c5-4d5a039129c4",
					Version:        "1.25.6",
				},
				Secrets: []KeyVaultSecrets{},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			actual := WindowsProfile{}
			convertVLabsWindowsProfile(&c.w, &actual)

			diff := cmp.Diff(actual, c.expected)
			if diff != "" {
				t.Errorf("unexpected diff testing convertVLabsWindowsProfile: %s", diff)
			}
		})
	}
}

func TestSetVlabsKubernetesDefaults(t *testing.T) {
	tests := []struct {
		name                  string
		p                     *vlabs.Properties
		expectedNetworkPlugin string
		expectedNetworkPolicy string
	}{
		{
			name: "default",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "",
						NetworkPolicy: "",
					},
				},
			},
			expectedNetworkPlugin: vlabs.DefaultNetworkPlugin,
			expectedNetworkPolicy: "",
		},
		{
			name: "default windows",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "",
						NetworkPolicy: "",
					},
				},
				AgentPoolProfiles: []*vlabs.AgentPoolProfile{
					{
						OSType: "Windows",
					},
				},
			},
			expectedNetworkPlugin: vlabs.DefaultNetworkPluginWindows,
			expectedNetworkPolicy: "",
		},
		{
			name: "azure networkPlugin",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "azure",
						NetworkPolicy: "",
					},
				},
			},
			expectedNetworkPlugin: vlabs.DefaultNetworkPlugin,
			expectedNetworkPolicy: "",
		},
		{
			name: "azure networkPolicy back-compat",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "",
						NetworkPolicy: "azure",
					},
				},
			},
			expectedNetworkPlugin: "azure",
			expectedNetworkPolicy: "",
		},
		{
			name: "none networkPolicy back-compat",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "",
						NetworkPolicy: "none",
					},
				},
			},
			expectedNetworkPlugin: "kubenet",
			expectedNetworkPolicy: "",
		},
		{
			name: "test literal string conversion",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "foo",
						NetworkPolicy: "bar",
					},
				},
			},
			expectedNetworkPlugin: "foo",
			expectedNetworkPolicy: "bar",
		},
		{
			name: "calico networkPlicy",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "",
						NetworkPolicy: "calico",
					},
				},
			},
			expectedNetworkPlugin: "azure",
			expectedNetworkPolicy: "calico",
		},
		{
			name: "cilium networkPlicy",
			p: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					KubernetesConfig: &vlabs.KubernetesConfig{
						NetworkPlugin: "",
						NetworkPolicy: "cilium",
					},
				},
			},
			expectedNetworkPlugin: "",
			expectedNetworkPolicy: "cilium",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			converted := &OrchestratorProfile{}
			setVlabsKubernetesDefaults(test.p, converted)
			if converted.KubernetesConfig.NetworkPlugin != test.expectedNetworkPlugin {
				t.Errorf("expected NetworkPlugin : %s, but got %s", test.expectedNetworkPlugin, converted.KubernetesConfig.NetworkPlugin)
			}
		})
	}
}
