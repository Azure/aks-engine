// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"k8s.io/apimachinery/pkg/api/equality"

	"github.com/Azure/aks-engine/pkg/api/common"
	v20170701 "github.com/Azure/aks-engine/pkg/api/v20170701"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/go-autorest/autorest/azure"
)

func TestAddDCOSPublicAgentPool(t *testing.T) {
	expectedNumPools := 2
	for _, masterCount := range [2]int{1, 3} {
		profiles := []*AgentPoolProfile{}
		profile := makeAgentPoolProfile(1, "agentprivate", "test-dcos-pool", "Standard_D2_v2", "Linux")
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
		DNSPrefix: "test-dcos",
		VMSize:    "Standard_D2_v2",
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
	cs = ConvertVLabsContainerService(vlabscs, false)
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(false) {
		t.Fatalf("incorrect OrchestratorVersion '%s'", cs.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	vlabscs = &vlabs.ContainerService{
		Properties: &vlabs.Properties{
			OrchestratorProfile: &vlabs.OrchestratorProfile{
				OrchestratorType:    vlabs.Kubernetes,
				OrchestratorVersion: "1.7.15",
			},
		},
	}
	cs = ConvertVLabsContainerService(vlabscs, false)
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != "1.7.15" {
		t.Fatalf("incorrect OrchestratorVersion '%s'", cs.Properties.OrchestratorProfile.OrchestratorVersion)
	}
}

func TestKubernetesVlabsDefaults(t *testing.T) {
	vp := makeKubernetesPropertiesVlabs()
	ap := makeKubernetesProperties()
	setVlabsKubernetesDefaults(vp, ap.OrchestratorProfile)
	if ap.OrchestratorProfile.KubernetesConfig == nil {
		t.Fatalf("KubernetesConfig cannot be nil after vlabs default conversion")
	}
	if ap.OrchestratorProfile.KubernetesConfig.NetworkPlugin != vlabs.DefaultNetworkPlugin {
		t.Fatalf("vlabs defaults not applied, expected NetworkPlugin: %s, instead got: %s", vlabs.DefaultNetworkPlugin, ap.OrchestratorProfile.KubernetesConfig.NetworkPlugin)
	}
	if ap.OrchestratorProfile.KubernetesConfig.NetworkPolicy != vlabs.DefaultNetworkPolicy {
		t.Fatalf("vlabs defaults not applied, expected NetworkPolicy: %s, instead got: %s", vlabs.DefaultNetworkPolicy, ap.OrchestratorProfile.KubernetesConfig.NetworkPolicy)
	}

	vp = makeKubernetesPropertiesVlabs()
	vp.WindowsProfile = &vlabs.WindowsProfile{}
	vp.AgentPoolProfiles = append(vp.AgentPoolProfiles, &vlabs.AgentPoolProfile{OSType: "Windows"})
	ap = makeKubernetesProperties()
	setVlabsKubernetesDefaults(vp, ap.OrchestratorProfile)
	if ap.OrchestratorProfile.KubernetesConfig == nil {
		t.Fatalf("KubernetesConfig cannot be nil after vlabs default conversion")
	}
	if ap.OrchestratorProfile.KubernetesConfig.NetworkPlugin != vlabs.DefaultNetworkPluginWindows {
		t.Fatalf("vlabs defaults not applied, expected NetworkPlugin: %s, instead got: %s", vlabs.DefaultNetworkPluginWindows, ap.OrchestratorProfile.KubernetesConfig.NetworkPlugin)
	}
	if ap.OrchestratorProfile.KubernetesConfig.NetworkPolicy != vlabs.DefaultNetworkPolicy {
		t.Fatalf("vlabs defaults not applied, expected NetworkPolicy: %s, instead got: %s", vlabs.DefaultNetworkPolicy, ap.OrchestratorProfile.KubernetesConfig.NetworkPolicy)
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

func makeKubernetesProperties() *Properties {
	ap := &Properties{}
	ap.OrchestratorProfile = &OrchestratorProfile{}
	ap.OrchestratorProfile.OrchestratorType = "Kubernetes"
	return ap
}

func makeKubernetesPropertiesVlabs() *vlabs.Properties {
	vp := &vlabs.Properties{}
	vp.OrchestratorProfile = &vlabs.OrchestratorProfile{}
	vp.OrchestratorProfile.OrchestratorType = "Kubernetes"
	return vp
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
		managementPortalURL          = "https=//management.local.azurestack.external/"
		publishSettingsURL           = "https=//management.local.azurestack.external/publishsettings/index"
		serviceManagementEndpoint    = "https=//management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189"
		resourceManagerEndpoint      = "https=//management.local.azurestack.external/"
		activeDirectoryEndpoint      = "https=//login.windows.net/"
		galleryEndpoint              = "https=//portal.local.azurestack.external=30015/"
		keyVaultEndpoint             = "https=//vault.azurestack.external/"
		graphEndpoint                = "https=//graph.windows.net/"
		serviceBusEndpoint           = "https=//servicebus.azurestack.external/"
		batchManagementEndpoint      = "https=//batch.azurestack.external/"
		storageEndpointSuffix        = "core.azurestack.external"
		sqlDatabaseDNSSuffix         = "database.azurestack.external"
		trafficManagerDNSSuffix      = "trafficmanager.cn"
		keyVaultDNSSuffix            = "vault.azurestack.external"
		serviceBusEndpointSuffix     = "servicebus.azurestack.external"
		serviceManagementVMDNSSuffix = "chinacloudapp.cn"
		resourceManagerVMDNSSuffix   = "cloudapp.azurestack.external"
		containerRegistryDNSSuffix   = "azurecr.io"
		tokenAudience                = "https=//management.azurestack.external/"
	)

	vlabscs := &vlabs.ContainerService{
		Properties: &vlabs.Properties{
			CustomCloudProfile: &vlabs.CustomCloudProfile{
				Enviornment: &azure.Environment{
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

	cs := ConvertVLabsContainerService(vlabscs, false)
	if cs.Properties.CustomCloudProfile.Enviornment.Name != name {
		t.Fatalf("incorrect Name, expect: '%s', actual: '%s'", name, cs.Properties.CustomCloudProfile.Enviornment.Name)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ManagementPortalURL != managementPortalURL {
		t.Fatalf("incorrect ManagementPortalURL, expect: '%s', actual: '%s'", managementPortalURL, cs.Properties.CustomCloudProfile.Enviornment.ManagementPortalURL)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.PublishSettingsURL != publishSettingsURL {
		t.Fatalf("incorrect PublishSettingsURL, expect: '%s', actual: '%s'", publishSettingsURL, cs.Properties.CustomCloudProfile.Enviornment.PublishSettingsURL)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ServiceManagementEndpoint != serviceManagementEndpoint {
		t.Fatalf("incorrect ServiceManagementEndpoint, expect: '%s', actual: '%s'", serviceManagementEndpoint, cs.Properties.CustomCloudProfile.Enviornment.ServiceManagementEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ResourceManagerEndpoint != resourceManagerEndpoint {
		t.Fatalf("incorrect ResourceManagerEndpoint, expect: '%s', actual: '%s'", resourceManagerEndpoint, cs.Properties.CustomCloudProfile.Enviornment.ResourceManagerEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ActiveDirectoryEndpoint != activeDirectoryEndpoint {
		t.Fatalf("incorrect ActiveDirectoryEndpoint, expect: '%s', actual: '%s'", activeDirectoryEndpoint, cs.Properties.CustomCloudProfile.Enviornment.ActiveDirectoryEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.GalleryEndpoint != galleryEndpoint {
		t.Fatalf("incorrect GalleryEndpoint, expect: '%s', actual: '%s'", galleryEndpoint, cs.Properties.CustomCloudProfile.Enviornment.GalleryEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.KeyVaultEndpoint != keyVaultEndpoint {
		t.Fatalf("incorrect KeyVaultEndpoint, expect: '%s', actual: '%s'", keyVaultEndpoint, cs.Properties.CustomCloudProfile.Enviornment.KeyVaultEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.GraphEndpoint != graphEndpoint {
		t.Fatalf("incorrect GraphEndpoint, expect: '%s', actual: '%s'", graphEndpoint, cs.Properties.CustomCloudProfile.Enviornment.GraphEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ServiceBusEndpoint != serviceBusEndpoint {
		t.Fatalf("incorrect ServiceBusEndpoint, expect: '%s', actual: '%s'", serviceBusEndpoint, cs.Properties.CustomCloudProfile.Enviornment.ServiceBusEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.BatchManagementEndpoint != batchManagementEndpoint {
		t.Fatalf("incorrect BatchManagementEndpoint, expect: '%s', actual: '%s'", batchManagementEndpoint, cs.Properties.CustomCloudProfile.Enviornment.BatchManagementEndpoint)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.StorageEndpointSuffix != storageEndpointSuffix {
		t.Fatalf("incorrect StorageEndpointSuffix, expect: '%s', actual: '%s'", storageEndpointSuffix, cs.Properties.CustomCloudProfile.Enviornment.StorageEndpointSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.SQLDatabaseDNSSuffix != sqlDatabaseDNSSuffix {
		t.Fatalf("incorrect SQLDatabaseDNSSuffix, expect: '%s', actual: '%s'", sqlDatabaseDNSSuffix, cs.Properties.CustomCloudProfile.Enviornment.SQLDatabaseDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.TrafficManagerDNSSuffix != trafficManagerDNSSuffix {
		t.Fatalf("incorrect TrafficManagerDNSSuffix, expect: '%s', actual: '%s'", trafficManagerDNSSuffix, cs.Properties.CustomCloudProfile.Enviornment.TrafficManagerDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.KeyVaultDNSSuffix != keyVaultDNSSuffix {
		t.Fatalf("incorrect KeyVaultDNSSuffix, expect: '%s', actual: '%s'", keyVaultDNSSuffix, cs.Properties.CustomCloudProfile.Enviornment.KeyVaultDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ServiceBusEndpointSuffix != serviceBusEndpointSuffix {
		t.Fatalf("incorrect ServiceBusEndpointSuffix, expect: '%s', actual: '%s'", serviceBusEndpointSuffix, cs.Properties.CustomCloudProfile.Enviornment.ServiceBusEndpointSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ServiceManagementVMDNSSuffix != serviceManagementVMDNSSuffix {
		t.Fatalf("incorrect ServiceManagementVMDNSSuffix, expect: '%s', actual: '%s'", serviceManagementVMDNSSuffix, cs.Properties.CustomCloudProfile.Enviornment.ServiceManagementVMDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ResourceManagerVMDNSSuffix != resourceManagerVMDNSSuffix {
		t.Fatalf("incorrect ResourceManagerVMDNSSuffix, expect: '%s', actual: '%s'", resourceManagerVMDNSSuffix, cs.Properties.CustomCloudProfile.Enviornment.ResourceManagerVMDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.ContainerRegistryDNSSuffix != containerRegistryDNSSuffix {
		t.Fatalf("incorrect ContainerRegistryDNSSuffix, expect: '%s', actual: '%s'", containerRegistryDNSSuffix, cs.Properties.CustomCloudProfile.Enviornment.ContainerRegistryDNSSuffix)
	}
	if cs.Properties.CustomCloudProfile.Enviornment.TokenAudience != tokenAudience {
		t.Fatalf("incorrect TokenAudience, expect: '%s', actual: '%s'", tokenAudience, cs.Properties.CustomCloudProfile.Enviornment.TokenAudience)
	}
}
