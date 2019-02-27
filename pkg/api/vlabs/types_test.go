// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"encoding/json"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/azure"
)

func TestKubernetesAddon(t *testing.T) {
	addon := KubernetesAddon{
		Name: "addon",
		Containers: []KubernetesContainerSpec{
			{
				Name:           "addon",
				CPURequests:    "50m",
				MemoryRequests: "150Mi",
				CPULimits:      "50m",
				MemoryLimits:   "150Mi",
			},
		},
	}
	if !addon.IsEnabled(true) {
		t.Fatalf("KubernetesAddon.IsEnabled(true) should always return true when Enabled property is not specified")
	}

	if addon.IsEnabled(false) {
		t.Fatalf("KubernetesAddon.IsEnabled(false) should always return false when Enabled property is not specified")
	}
	e := true
	addon.Enabled = &e
	if !addon.IsEnabled(false) {
		t.Fatalf("KubernetesAddon.IsEnabled(false) should always return true when Enabled property is set to true")
	}
	if !addon.IsEnabled(true) {
		t.Fatalf("KubernetesAddon.IsEnabled(true) should always return true when Enabled property is set to true")
	}
	e = false
	addon.Enabled = &e
	if addon.IsEnabled(false) {
		t.Fatalf("KubernetesAddon.IsEnabled(false) should always return false when Enabled property is set to false")
	}
	if addon.IsEnabled(true) {
		t.Fatalf("KubernetesAddon.IsEnabled(true) should always return false when Enabled property is set to false")
	}
}

func TestOrchestratorProfile(t *testing.T) {
	OrchestratorProfileText := `{ "orchestratorType": "Mesos" }`
	op := &OrchestratorProfile{}
	if e := json.Unmarshal([]byte(OrchestratorProfileText), op); e == nil {
		t.Fatalf("expected unmarshal failure for OrchestratorProfile when passing an invalid orchestratorType")
	}

	OrchestratorProfileText = `{ "orchestratorType": "Swarm" }`
	op = &OrchestratorProfile{}
	if e := json.Unmarshal([]byte(OrchestratorProfileText), op); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for OrchestratorProfile, %+v", e)
	}

	OrchestratorProfileText = `{ "orchestratorType": "SwarmMode" }`
	op = &OrchestratorProfile{}
	if e := json.Unmarshal([]byte(OrchestratorProfileText), op); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for OrchestratorProfile, %+v", e)
	}

	if !op.IsSwarmMode() {
		t.Fatalf("unexpectedly detected OrchestratorProfile.Type != DockerCE after unmarshal")

	}

	OrchestratorProfileText = `{ "orchestratorType": "DCOS" }`
	op = &OrchestratorProfile{}
	if e := json.Unmarshal([]byte(OrchestratorProfileText), op); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for OrchestratorProfile, %+v", e)
	}

	OrchestratorProfileText = `{ "orchestratorType": "Kubernetes" }`
	op = &OrchestratorProfile{}
	if e := json.Unmarshal([]byte(OrchestratorProfileText), op); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for OrchestratorProfile, %+v", e)

	}
}

func TestMasterProfile(t *testing.T) {
	// With osType not specified
	MasterProfileText := `{"count" : 0, "storageProfile" : "StorageAccount", "vnetSubnetID" : "1234", "agentVnetSubnetID" : "5678"}`
	mp := &MasterProfile{}
	if e := json.Unmarshal([]byte(MasterProfileText), mp); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for MasterProfile, %+v", e)
	}

	if mp.Count != 0 {
		t.Fatalf("unexpectedly detected MasterProfile.Count != 1 after unmarshal")
	}

	if !mp.IsCustomVNET() {
		t.Fatalf("unexpectedly detected nil MasterProfile.VNetSubNetID after unmarshal")
	}

	if !mp.IsStorageAccount() {
		t.Fatalf("unexpectedly detected MasterProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	// With vmas
	MasterProfileText = `{  "count": 1, "vmSize": "Standard_D2_v2", "storageProfile" : "ManagedDisks", "diskSizesGB" : [750, 250, 600, 1000] }`
	mp = &MasterProfile{}
	if e := json.Unmarshal([]byte(MasterProfileText), mp); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for MasterProfile, %+v", e)
	}

	if mp.Count != 1 {
		t.Fatalf("unexpectedly detected MasterProfile.Count != 1 after unmarshal")
	}

	if !mp.IsManagedDisks() {
		t.Fatalf("unexpectedly detected MasterProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	if mp.IsVirtualMachineScaleSets() {
		t.Fatalf("unexpectedly detected MasterProfile.AvailabilitySets == VirtualMachineScaleSets after unmarshal")
	}

	// With VMSS and zones
	MasterProfileText = `{  "count": 3, "vmSize": "Standard_D2_v2", "availabilityProfile": "VirtualMachineScaleSets", "storageProfile" : "ManagedDisks", "diskSizesGB" : [750, 250, 600, 1000],  "AvailabilityZones": ["1","2"] }`
	mp = &MasterProfile{}
	if e := json.Unmarshal([]byte(MasterProfileText), mp); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for MasterProfile, %+v", e)
	}

	if mp.Count != 3 {
		t.Fatalf("unexpectedly detected MasterProfile.Count != 3 after unmarshal")
	}

	if !mp.IsManagedDisks() {
		t.Fatalf("unexpectedly detected MasterProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	if !mp.IsVirtualMachineScaleSets() {
		t.Fatalf("unexpectedly detected MasterProfile.AvailabilitySets != VirtualMachineScaleSets after unmarshal")
	}

	if !mp.HasAvailabilityZones() {
		t.Fatalf("unexpectedly detected MasterProfile.AvailabilityZones, HasAvailabilityZones returned false after unmarshal")
	}
}
func TestAgentPoolProfile(t *testing.T) {
	// With osType not specified
	AgentPoolProfileText := `{"count" : 0, "storageProfile" : "StorageAccount", "vnetSubnetID" : "1234"}`
	ap := &AgentPoolProfile{}
	if e := json.Unmarshal([]byte(AgentPoolProfileText), ap); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for AgentPoolProfile, %+v", e)
	}

	if ap.Count != 0 {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Count != 1 after unmarshal")
	}

	if !ap.IsCustomVNET() {
		t.Fatalf("unexpectedly detected nil AgentPoolProfile.VNetSubNetID after unmarshal")
	}

	if !ap.IsStorageAccount() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	// With osType Windows
	AgentPoolProfileText = `{ "name": "linuxpool1", "osType" : "Windows", "count": 1, "vmSize": "Standard_D2_v2",
"availabilityProfile": "AvailabilitySet", "storageProfile" : "ManagedDisks", "vnetSubnetID" : "12345" }`
	ap = &AgentPoolProfile{}
	if e := json.Unmarshal([]byte(AgentPoolProfileText), ap); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for AgentPoolProfile, %+v", e)
	}

	if ap.Count != 1 {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Count != 1 after unmarshal")
	}

	if !ap.IsWindows() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.OSType != Windows after unmarshal")
	}

	if !ap.IsManagedDisks() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	// With osType Linux and RHEL distro
	AgentPoolProfileText = `{ "name": "linuxpool1", "osType" : "Linux", "distro" : "rhel", "count": 1, "vmSize": "Standard_D2_v2",
"availabilityProfile": "AvailabilitySet", "storageProfile" : "ManagedDisks", "vnetSubnetID" : "12345" }`
	ap = &AgentPoolProfile{}
	if e := json.Unmarshal([]byte(AgentPoolProfileText), ap); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for AgentPoolProfile, %+v", e)
	}

	if ap.Count != 1 {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Count != 1 after unmarshal")
	}

	if !ap.IsLinux() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.OSType != Linux after unmarshal")
	}

	if !ap.IsRHEL() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Distro != RHEL after unmarshal")
	}

	if !ap.IsManagedDisks() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	// With osType Linux and coreos distro
	AgentPoolProfileText = `{ "name": "linuxpool1", "osType" : "Linux", "distro" : "coreos", "count": 1, "vmSize": "Standard_D2_v2",
"availabilityProfile": "VirtualMachineScaleSets", "storageProfile" : "ManagedDisks", "diskSizesGB" : [750, 250, 600, 1000] }`
	ap = &AgentPoolProfile{}
	if e := json.Unmarshal([]byte(AgentPoolProfileText), ap); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for AgentPoolProfile, %+v", e)
	}

	if ap.Count != 1 {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Count != 1 after unmarshal")
	}

	if !ap.IsLinux() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.OSType != Linux after unmarshal")
	}

	if !ap.IsCoreOS() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Distro != CoreOS after unmarshal")
	}

	if !ap.IsManagedDisks() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	if !ap.HasDisks() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.DiskSizesGB < 0 after unmarshal")
	}

	if !ap.IsVirtualMachineScaleSets() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.AvailabilitySets != VirtualMachineScaleSets after unmarshal")
	}
}

func TestContainerServiceProperties(t *testing.T) {
	// Agent pool with availability zones
	ContainerServicePropertiesText := `{"orchestratorProfile": {"orchestratorType": "Kubernetes","orchestratorRelease": "1.11"}, "agentPoolProfiles":[{ "name": "linuxpool1", "osType" : "Linux", "count": 1, "vmSize": "Standard_D2_v2",
		"availabilityProfile": "VirtualMachineScaleSets", "AvailabilityZones": ["1","2"]}]}`
	prop := &Properties{}
	if e := json.Unmarshal([]byte(ContainerServicePropertiesText), prop); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for ContainerServiceProperties, %+v", e)
	}

	if !prop.HasAvailabilityZones() {
		t.Fatalf("unexpectedly detected ContainerServiceProperties HasAvailabilityZones returns false  after unmarshal")
	}

	// master profile with availability zones
	ContainerServicePropertiesText = `{"orchestratorProfile": {"orchestratorType": "Kubernetes","orchestratorRelease": "1.12"}, "masterProfile":{"count": 4, "vmSize": "Standard_D2_v2", "availabilityProfile": "VirtualMachineScaleSets", "storageProfile": "ManagedDisks", "diskSizesGB": [750, 250, 600, 1000], "availabilityZones": ["1","2"] }, "agentPoolProfiles":[{ "name": "linuxpool1", "osType" : "Linux", "count": 1, "vmSize": "Standard_D2_v2",
		"availabilityProfile": "VirtualMachineScaleSets"}]}`
	prop = &Properties{}
	if e := json.Unmarshal([]byte(ContainerServicePropertiesText), prop); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for ContainerServiceProperties, %+v", e)
	}

	if !prop.HasAvailabilityZones() {
		t.Fatalf("unexpectedly detected ContainerServiceProperties HasAvailabilityZones returns false  after unmarshal")
	}

	if prop.MastersAndAgentsUseAvailabilityZones() {
		t.Fatalf("unexpectedly detected ContainerServiceProperties MastersAndAgentsUseAvailabilityZones returns true  after unmarshal")
	}
	// master profile and agent profile with availability zones
	ContainerServicePropertiesText = `{"orchestratorProfile": {"orchestratorType": "Kubernetes","orchestratorRelease": "1.12"}, "masterProfile":{"count": 4, "vmSize": "Standard_D2_v2", "availabilityProfile": "VirtualMachineScaleSets", "storageProfile": "ManagedDisks", "diskSizesGB": [750, 250, 600, 1000], "availabilityZones": ["1","2"] }, "agentPoolProfiles":[{ "name": "linuxpool1", "osType" : "Linux", "count": 1, "vmSize": "Standard_D2_v2",
		"availabilityProfile": "VirtualMachineScaleSets", "availabilityZones": ["1","2"] }]}`
	prop = &Properties{}
	if e := json.Unmarshal([]byte(ContainerServicePropertiesText), prop); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for ContainerServiceProperties, %+v", e)
	}

	if !prop.HasAvailabilityZones() {
		t.Fatalf("unexpectedly detected ContainerServiceProperties HasAvailabilityZones returns false  after unmarshal")
	}

	if !prop.MastersAndAgentsUseAvailabilityZones() {
		t.Fatalf("unexpectedly detected ContainerServiceProperties MastersAndAgentsUseAvailabilityZones returns false  after unmarshal")
	}
}

func TestAgentPoolIsNSeriesSKU(t *testing.T) {
	cases := common.GetNSeriesVMCasesForTesting()

	for _, c := range cases {
		p := Properties{
			AgentPoolProfiles: []*AgentPoolProfile{
				{
					Name:   "agentpool",
					VMSize: c.VMSKU,
					Count:  1,
				},
			},
			OrchestratorProfile: &OrchestratorProfile{
				OrchestratorType:    Kubernetes,
				OrchestratorRelease: "1.12",
			},
		}
		ret := p.AgentPoolProfiles[0].IsNSeriesSKU()
		if ret != c.Expected {
			t.Fatalf("expected IsNvidiaEnabledSKU(%s) to return %t, but instead got %t", c.VMSKU, c.Expected, ret)
		}
	}
}

func TestIsAzureStackCloud(t *testing.T) {
	testcases := []struct {
		name       string
		properties Properties
		expected   bool
	}{
		{
			"Empty environment name",
			getMockPropertiesWithCustomCloudProfile("", true, true, false),
			false,
		},
		{
			"Empty environment name with AzureEnvironmentSpecConfig",
			getMockPropertiesWithCustomCloudProfile("", true, true, true),
			false,
		},
		{
			"lower case cloud name",
			getMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true),
			true,
		},
		{
			"cammel case cloud name",
			getMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, true, true),
			true,
		},
		{
			"incorrect cloud name",
			getMockPropertiesWithCustomCloudProfile("NotAzureStackCloud", true, true, true),
			false,
		},
		{
			"empty cloud profile",
			getMockPropertiesWithCustomCloudProfile("AzureStackCloud", false, false, false),
			false,
		},
		{
			"empty environment ",
			getMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, false, true),
			false,
		},
	}
	for _, testcase := range testcases {
		actual := testcase.properties.IsAzureStackCloud()
		if testcase.expected != actual {
			t.Errorf("Test \"%s\": expected IsAzureStackCloud() to return %t, but got %t . ", testcase.name, testcase.expected, actual)
		}
	}
}

func getMockPropertiesWithCustomCloudProfile(name string, hasCustomCloudProfile, hasEnvironment, hasAzureEnvironmentSpecConfig bool) Properties {
	const (
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

	p := Properties{}
	if hasCustomCloudProfile {
		p.CustomCloudProfile = &CustomCloudProfile{}
		if hasEnvironment {
			p.CustomCloudProfile.Environment = &azure.Environment{
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
			}
		}
		if hasAzureEnvironmentSpecConfig {
			//azureStackCloudSpec is the default configurations for azure stack with public Azure.
			azureStackCloudSpec := AzureEnvironmentSpecConfig{
				CloudName: AzureStackCloud,
			}
			p.CustomCloudProfile.AzureEnvironmentSpecConfig = &azureStackCloudSpec
		}
	}
	return p
}
