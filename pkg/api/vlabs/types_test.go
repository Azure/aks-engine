// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"encoding/json"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/azure"
)

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
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != StorageAccount after unmarshal")
	}

	if ap.IsEphemeral() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile == Ephemeral after unmarshal")
	}

	if ap.DiskEncryptionSetID != "" {
		t.Fatalf("unexpectedly detected AgentPoolProfile.DiskEncryptionSetID is not empty after unmarshal")
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

	// With osType Windows and Ephemeral disks
	AgentPoolProfileText = `{ "name": "linuxpool1", "osType" : "Windows", "count": 1, "vmSize": "Standard_D2_v2",
"availabilityProfile": "AvailabilitySet", "storageProfile" : "Ephemeral", "vnetSubnetID" : "12345", "diskEncryptionSetID": "diskEncryptionSetID" }`
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

	if ap.IsManagedDisks() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile == ManagedDisks after unmarshal")
	}

	if !ap.IsEphemeral() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != Ephemeral after unmarshal")
	}

	if ap.DiskEncryptionSetID == "" {
		t.Fatalf("unexpectedly detected AgentPoolProfile.DiskEncryptionSetID is empty after unmarshal")
	}

	// With osType Linux and RHEL distro
	AgentPoolProfileText = `{ "name": "linuxpool1", "osType" : "Linux", "distro" : "rhel", "count": 1, "vmSize": "Standard_D2_v2",
"availabilityProfile": "AvailabilitySet", "storageProfile" : "ManagedDisks", "vnetSubnetID" : "12345", "diskEncryptionSetID": "diskEncryptionSetID" }`
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

	if ap.IsEphemeral() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile == Ephemeral after unmarshal")
	}

	if ap.DiskEncryptionSetID == "" {
		t.Fatalf("unexpectedly detected AgentPoolProfile.DiskEncryptionSetID is empty after unmarshal")
	}

	// With VMSS and Spot VMs
	AgentPoolProfileText = `{"name":"linuxpool1","osType":"Linux","distro":"rhel","count":1,"vmSize":"Standard_D2_v2",
"availabilityProfile":"VirtualMachineScaleSets","scaleSetPriority":"Spot","ScaleSetEvictionPolicy":"Delete","SpotMaxPrice":88}`
	ap = &AgentPoolProfile{}
	if e := json.Unmarshal([]byte(AgentPoolProfileText), ap); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for AgentPoolProfile, %+v", e)
	}

	if ap.ScaleSetPriority != "Spot" {
		t.Fatalf("unexpectedly detected AgentPoolProfile.ScaleSetPriority != ScaleSetPrioritySpot after unmarshal")
	}

	if ap.ScaleSetEvictionPolicy != "Delete" {
		t.Fatalf("unexpectedly detected AgentPoolProfile.ScaleSetEvictionPolicy != ScaleSetEvictionPolicyDelete after unmarshal")
	}

	if *ap.SpotMaxPrice != float64(88) {
		t.Fatalf("unexpectedly detected *AgentPoolProfile.SpotMaxPrice != float64(88) after unmarshal")
	}

	// With osType Linux and coreos distro
	AgentPoolProfileText = `{ "name": "linuxpool1", "osType" : "Linux", "distro" : "coreos", "count": 1, "vmSize": "Standard_D2_v2",
"availabilityProfile": "VirtualMachineScaleSets", "storageProfile" : "ManagedDisks", "diskSizesGB" : [750, 250, 600, 1000], "diskEncryptionSetID": "diskEncryptionSetID" }`
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

	if ap.IsEphemeral() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile == Ephemeral after unmarshal")
	}

	if !ap.HasDisks() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.DiskSizesGB < 0 after unmarshal")
	}

	if !ap.IsVirtualMachineScaleSets() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.AvailabilitySets != VirtualMachineScaleSets after unmarshal")
	}

	if ap.DiskEncryptionSetID == "" {
		t.Fatalf("unexpectedly detected AgentPoolProfile.DiskEncryptionSetID is empty after unmarshal")
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
			GetMockPropertiesWithCustomCloudProfile("", true, true, false),
			true,
		},
		{
			"Empty environment name with AzureEnvironmentSpecConfig",
			GetMockPropertiesWithCustomCloudProfile("", true, true, true),
			true,
		},
		{
			"lower case cloud name",
			GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true),
			true,
		},
		{
			"cammel case cloud name",
			GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, true, true),
			true,
		},
		{
			"incorrect cloud name",
			GetMockPropertiesWithCustomCloudProfile("NotAzureStackCloud", true, true, true),
			true,
		},
		{
			"empty cloud profile",
			GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", false, false, false),
			false,
		},
		{
			"empty environment ",
			GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, false, true),
			true,
		},
	}
	for _, testcase := range testcases {
		actual := testcase.properties.IsAzureStackCloud()
		if testcase.expected != actual {
			t.Errorf("Test \"%s\": expected IsAzureStackCloud() to return %t, but got %t . ", testcase.name, testcase.expected, actual)
		}
	}
}

func TestUbuntuVersion(t *testing.T) {
	cases := []struct {
		p                  Properties
		expectedMaster1604 bool
		expectedAgent1604  bool
		expectedMaster1804 bool
		expectedAgent1804  bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1604,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: AKSUbuntu1604,
						OSType: Linux,
					},
				},
			},
			expectedMaster1604: true,
			expectedAgent1604:  true,
			expectedMaster1804: false,
			expectedAgent1804:  false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: ACC1604,
					},
				},
			},
			expectedMaster1604: false,
			expectedAgent1604:  true,
			expectedMaster1804: true,
			expectedAgent1804:  false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: "",
						OSType: Windows,
					},
				},
			},
			expectedMaster1604: true,
			expectedAgent1604:  false,
			expectedMaster1804: false,
			expectedAgent1804:  false,
		},
	}

	for _, c := range cases {
		if c.p.MasterProfile.IsUbuntu1604() != c.expectedMaster1604 {
			t.Fatalf("expected IsUbuntu1604() for master to return %t but instead returned %t", c.expectedMaster1604, c.p.MasterProfile.IsUbuntu1604())
		}
		if c.p.AgentPoolProfiles[0].IsUbuntu1604() != c.expectedAgent1604 {
			t.Fatalf("expected IsUbuntu1604() for agent to return %t but instead returned %t", c.expectedAgent1604, c.p.AgentPoolProfiles[0].IsUbuntu1604())
		}
		if c.p.MasterProfile.IsUbuntu1804() != c.expectedMaster1804 {
			t.Fatalf("expected IsUbuntu1804() for master to return %t but instead returned %t", c.expectedMaster1804, c.p.MasterProfile.IsUbuntu1804())
		}
		if c.p.AgentPoolProfiles[0].IsUbuntu1804() != c.expectedAgent1804 {
			t.Fatalf("expected IsUbuntu1804() for agent to return %t but instead returned %t", c.expectedAgent1804, c.p.AgentPoolProfiles[0].IsUbuntu1804())
		}
	}
}

func TestMasterIsUbuntu(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1604,
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: ACC1604,
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1804,
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: CoreOS,
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: RHEL,
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: "foo",
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.MasterProfile.IsUbuntu() != c.expected {
			t.Fatalf("expected IsUbuntu() to return %t but instead returned %t", c.expected, c.p.MasterProfile.IsUbuntu())
		}
	}
}

func TestAgentPoolIsUbuntu(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: AKSUbuntu1604,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu1804,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: ACC1604,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: AKSUbuntu1804,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu1804,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: CoreOS,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: RHEL,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: "foo",
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.AgentPoolProfiles[0].IsUbuntu() != c.expected {
			t.Fatalf("expected IsUbuntu() to return %t but instead returned %t", c.expected, c.p.AgentPoolProfiles[0].IsUbuntu())
		}
	}
}

func GetMockPropertiesWithCustomCloudProfile(name string, hasCustomCloudProfile, hasEnvironment, hasAzureEnvironmentSpecConfig bool) Properties {
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

func TestAADAdminGroupIDMethods(t *testing.T) {
	tests := []struct {
		name                       string
		properties                 *Properties
		expectedHasAADAdminGroupID bool
		expectedGetAADAdminGroupID string
	}{
		{
			name:                       "default",
			properties:                 &Properties{},
			expectedHasAADAdminGroupID: false,
			expectedGetAADAdminGroupID: "",
		},
		{
			name: "no AdminGroupID",
			properties: &Properties{
				AADProfile: &AADProfile{
					ClientAppID: "",
				},
			},
			expectedHasAADAdminGroupID: false,
			expectedGetAADAdminGroupID: "",
		},
		{
			name: "AdminGroupID value",
			properties: &Properties{
				AADProfile: &AADProfile{
					AdminGroupID: "7d04bcd3-3c48-49ab-a064-c0b7d69896da",
				},
			},
			expectedHasAADAdminGroupID: true,
			expectedGetAADAdminGroupID: "7d04bcd3-3c48-49ab-a064-c0b7d69896da",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			hasAADAdminGroupID := test.properties.HasAADAdminGroupID()
			if hasAADAdminGroupID != test.expectedHasAADAdminGroupID {
				t.Errorf("expected HasAADAdminGroupID %t, but got %t", test.expectedHasAADAdminGroupID, hasAADAdminGroupID)
			}
			getAADAdminGroupID := test.properties.GetAADAdminGroupID()
			if getAADAdminGroupID != test.expectedGetAADAdminGroupID {
				t.Errorf("expected HasAADAdminGroupID %s, but got %s", test.expectedGetAADAdminGroupID, getAADAdminGroupID)
			}
		})
	}
}
