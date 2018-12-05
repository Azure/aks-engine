// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"log"
	"reflect"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
)

const exampleCustomHyperkubeImage = `example.azurecr.io/example/hyperkube-amd64:custom`

const exampleAPIModel = `{
		"apiVersion": "vlabs",
	"properties": {
		"orchestratorProfile": {
			"orchestratorType": "Kubernetes",
			"kubernetesConfig": {
				"customHyperkubeImage": "` + exampleCustomHyperkubeImage + `"
			}
		},
		"masterProfile": { "count": 1, "dnsPrefix": "", "vmSize": "Standard_D2_v2" },
		"agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
		"windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
		"linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
		},
		"servicePrincipalProfile": { "clientId": "", "secret": "" }
	}
}
`

const exampleSystemMSIModel = `{
	"apiVersion": "vlabs",
"properties": {
	"orchestratorProfile": {
		"orchestratorType": "Kubernetes",
		"kubernetesConfig": {
			"useManagedIdentity": true
		}
	},
	"masterProfile": { "count": 1, "dnsPrefix": "", "vmSize": "Standard_D2_v2" },
	"agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
	"windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
	"linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	},
	"servicePrincipalProfile": { "clientId": "", "secret": "" }
}
}
`

const exampleUserMSI = "/subscriptions/<subscription>/resourcegroups/<rg_name>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identityName>"

const exampleUserMSIModel = `{
	"apiVersion": "vlabs",
"properties": {
	"orchestratorProfile": {
		"orchestratorType": "Kubernetes",
		"kubernetesConfig": {
			"useManagedIdentity": true,
			"userAssignedID": "` + exampleUserMSI + `"
		}
	},
	"masterProfile": { "count": 1, "dnsPrefix": "", "vmSize": "Standard_D2_v2" },
	"agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
	"windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
	"linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	},
	"servicePrincipalProfile": { "clientId": "", "secret": "" }
}
}
`

func TestOSType(t *testing.T) {
	p := Properties{
		MasterProfile: &MasterProfile{
			Distro: RHEL,
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				OSType: Linux,
			},
			{
				OSType: Linux,
				Distro: RHEL,
			},
		},
	}

	if p.HasWindows() {
		t.Fatalf("expected HasWindows() to return false but instead returned true")
	}
	if p.AgentPoolProfiles[0].IsWindows() {
		t.Fatalf("expected IsWindows() to return false but instead returned true")
	}

	if !p.AgentPoolProfiles[0].IsLinux() {
		t.Fatalf("expected IsLinux() to return true but instead returned false")
	}

	if p.AgentPoolProfiles[0].IsRHEL() {
		t.Fatalf("expected IsRHEL() to return false but instead returned true")
	}

	if p.AgentPoolProfiles[0].IsCoreOS() {
		t.Fatalf("expected IsCoreOS() to return false but instead returned true")
	}

	if !p.AgentPoolProfiles[1].IsRHEL() {
		t.Fatalf("expected IsRHEL() to return true but instead returned false")
	}

	if p.AgentPoolProfiles[1].IsCoreOS() {
		t.Fatalf("expected IsCoreOS() to return false but instead returned true")
	}

	if !p.MasterProfile.IsRHEL() {
		t.Fatalf("expected IsRHEL() to return true but instead returned false")
	}

	if p.MasterProfile.IsCoreOS() {
		t.Fatalf("expected IsCoreOS() to return false but instead returned true")
	}

	p.MasterProfile.Distro = CoreOS
	p.AgentPoolProfiles[0].OSType = Windows
	p.AgentPoolProfiles[1].Distro = CoreOS

	if !p.HasWindows() {
		t.Fatalf("expected HasWindows() to return true but instead returned false")
	}

	if !p.AgentPoolProfiles[0].IsWindows() {
		t.Fatalf("expected IsWindows() to return true but instead returned false")
	}

	if p.AgentPoolProfiles[0].IsLinux() {
		t.Fatalf("expected IsLinux() to return false but instead returned true")
	}

	if p.AgentPoolProfiles[0].IsRHEL() {
		t.Fatalf("expected IsRHEL() to return false but instead returned true")
	}

	if p.AgentPoolProfiles[0].IsCoreOS() {
		t.Fatalf("expected IsCoreOS() to return false but instead returned true")
	}

	if p.AgentPoolProfiles[1].IsRHEL() {
		t.Fatalf("expected IsRHEL() to return false but instead returned true")
	}

	if !p.AgentPoolProfiles[1].IsCoreOS() {
		t.Fatalf("expected IsCoreOS() to return true but instead returned false")
	}

	if p.MasterProfile.IsRHEL() {
		t.Fatalf("expected IsRHEL() to return false but instead returned true")
	}

	if !p.MasterProfile.IsCoreOS() {
		t.Fatalf("expected IsCoreOS() to return true but instead returned false")
	}
}

func TestHasStorageProfile(t *testing.T) {
	cases := []struct {
		p                 Properties
		expectedHasMD     bool
		expectedHasSA     bool
		expectedMasterMD  bool
		expectedAgent0MD  bool
		expectedPrivateJB bool
		expectedHasDisks  bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					StorageProfile: StorageAccount,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile: StorageAccount,
						DiskSizesGB:    []int{5},
					},
					{
						StorageProfile: StorageAccount,
					},
				},
			},
			expectedHasMD:    false,
			expectedHasSA:    true,
			expectedMasterMD: false,
			expectedAgent0MD: false,
			expectedHasDisks: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					StorageProfile: ManagedDisks,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile: StorageAccount,
					},
					{
						StorageProfile: StorageAccount,
					},
				},
			},
			expectedHasMD:    true,
			expectedHasSA:    true,
			expectedMasterMD: true,
			expectedAgent0MD: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					StorageProfile: StorageAccount,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile: ManagedDisks,
					},
					{
						StorageProfile: StorageAccount,
					},
				},
			},
			expectedHasMD:    true,
			expectedHasSA:    true,
			expectedMasterMD: false,
			expectedAgent0MD: true,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					StorageProfile: ManagedDisks,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile: ManagedDisks,
					},
					{
						StorageProfile: ManagedDisks,
					},
				},
			},
			expectedHasMD:     true,
			expectedHasSA:     false,
			expectedMasterMD:  true,
			expectedAgent0MD:  true,
			expectedPrivateJB: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						PrivateCluster: &PrivateCluster{
							Enabled: helpers.PointerToBool(true),
							JumpboxProfile: &PrivateJumpboxProfile{
								StorageProfile: ManagedDisks,
							},
						},
					},
				},
				MasterProfile: &MasterProfile{
					StorageProfile: StorageAccount,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile: StorageAccount,
					},
				},
			},
			expectedHasMD:     true,
			expectedHasSA:     true,
			expectedMasterMD:  false,
			expectedAgent0MD:  false,
			expectedPrivateJB: true,
		},

		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						PrivateCluster: &PrivateCluster{
							Enabled: helpers.PointerToBool(true),
							JumpboxProfile: &PrivateJumpboxProfile{
								StorageProfile: StorageAccount,
							},
						},
					},
				},
				MasterProfile: &MasterProfile{
					StorageProfile: ManagedDisks,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile: ManagedDisks,
					},
				},
			},
			expectedHasMD:     true,
			expectedHasSA:     true,
			expectedMasterMD:  true,
			expectedAgent0MD:  true,
			expectedPrivateJB: true,
		},
	}

	for _, c := range cases {
		if c.p.HasManagedDisks() != c.expectedHasMD {
			t.Fatalf("expected HasManagedDisks() to return %t but instead returned %t", c.expectedHasMD, c.p.HasManagedDisks())
		}
		if c.p.HasStorageAccountDisks() != c.expectedHasSA {
			t.Fatalf("expected HasStorageAccountDisks() to return %t but instead returned %t", c.expectedHasSA, c.p.HasStorageAccountDisks())
		}
		if c.p.MasterProfile.IsManagedDisks() != c.expectedMasterMD {
			t.Fatalf("expected IsManagedDisks() to return %t but instead returned %t", c.expectedMasterMD, c.p.MasterProfile.IsManagedDisks())
		}
		if c.p.MasterProfile.IsStorageAccount() == c.expectedMasterMD {
			t.Fatalf("expected IsStorageAccount() to return %t but instead returned %t", !c.expectedMasterMD, c.p.MasterProfile.IsStorageAccount())
		}
		if c.p.AgentPoolProfiles[0].IsManagedDisks() != c.expectedAgent0MD {
			t.Fatalf("expected IsManagedDisks() to return %t but instead returned %t", c.expectedAgent0MD, c.p.AgentPoolProfiles[0].IsManagedDisks())
		}
		if c.p.AgentPoolProfiles[0].IsStorageAccount() == c.expectedAgent0MD {
			t.Fatalf("expected IsStorageAccount() to return %t but instead returned %t", !c.expectedAgent0MD, c.p.AgentPoolProfiles[0].IsStorageAccount())
		}
		if c.p.OrchestratorProfile != nil && c.p.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() != c.expectedPrivateJB {
			t.Fatalf("expected PrivateJumpboxProvision() to return %t but instead returned %t", c.expectedPrivateJB, c.p.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision())
		}
		if c.p.AgentPoolProfiles[0].HasDisks() != c.expectedHasDisks {
			t.Fatalf("expected HasDisks() to return %t but instead returned %t", c.expectedHasDisks, c.p.AgentPoolProfiles[0].HasDisks())
		}
	}
}

func TestTotalNodes(t *testing.T) {
	cases := []struct {
		p        Properties
		expected int
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count: 1,
					},
				},
			},
			expected: 2,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count: 3,
					},
					{
						Count: 4,
					},
				},
			},
			expected: 7,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 5,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count: 6,
					},
				},
			},
			expected: 11,
		},
	}

	for _, c := range cases {
		if c.p.TotalNodes() != c.expected {
			t.Fatalf("expected TotalNodes() to return %d but instead returned %d", c.expected, c.p.TotalNodes())
		}
	}
}
func TestMasterAvailabilityProfile(t *testing.T) {
	cases := []struct {
		p              Properties
		expectedISVMSS bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{},
			},
			expectedISVMSS: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					AvailabilityProfile: AvailabilitySet,
				},
			},
			expectedISVMSS: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					AvailabilityProfile: VirtualMachineScaleSets,
				},
			},
			expectedISVMSS: true,
		},
	}

	for _, c := range cases {
		if c.p.MasterProfile.IsVirtualMachineScaleSets() != c.expectedISVMSS {
			t.Fatalf("expected MasterProfile.IsVirtualMachineScaleSets() to return %t but instead returned %t", c.expectedISVMSS, c.p.MasterProfile.IsVirtualMachineScaleSets())
		}
	}
}
func TestAvailabilityProfile(t *testing.T) {
	cases := []struct {
		p               Properties
		expectedHasVMSS bool
		expectedISVMSS  bool
		expectedIsAS    bool
		expectedLowPri  bool
	}{
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
						ScaleSetPriority:    ScaleSetPriorityLow,
					},
				},
			},
			expectedHasVMSS: true,
			expectedISVMSS:  true,
			expectedIsAS:    false,
			expectedLowPri:  true,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
						ScaleSetPriority:    ScaleSetPriorityRegular,
					},
					{
						AvailabilityProfile: AvailabilitySet,
					},
				},
			},
			expectedHasVMSS: true,
			expectedISVMSS:  true,
			expectedIsAS:    false,
			expectedLowPri:  false,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: AvailabilitySet,
					},
				},
			},
			expectedHasVMSS: false,
			expectedISVMSS:  false,
			expectedIsAS:    true,
			expectedLowPri:  false,
		},
	}

	for _, c := range cases {
		if c.p.HasVMSSAgentPool() != c.expectedHasVMSS {
			t.Fatalf("expected HasVMSSAgentPool() to return %t but instead returned %t", c.expectedHasVMSS, c.p.HasVMSSAgentPool())
		}
		if c.p.AgentPoolProfiles[0].IsVirtualMachineScaleSets() != c.expectedISVMSS {
			t.Fatalf("expected IsVirtualMachineScaleSets() to return %t but instead returned %t", c.expectedISVMSS, c.p.AgentPoolProfiles[0].IsVirtualMachineScaleSets())
		}
		if c.p.AgentPoolProfiles[0].IsAvailabilitySets() != c.expectedIsAS {
			t.Fatalf("expected IsAvailabilitySets() to return %t but instead returned %t", c.expectedIsAS, c.p.AgentPoolProfiles[0].IsAvailabilitySets())
		}
		if c.p.AgentPoolProfiles[0].IsLowPriorityScaleSet() != c.expectedLowPri {
			t.Fatalf("expected IsLowPriorityScaleSet() to return %t but instead returned %t", c.expectedLowPri, c.p.AgentPoolProfiles[0].IsLowPriorityScaleSet())
		}
	}
}

func TestIsCustomVNET(t *testing.T) {
	cases := []struct {
		p              Properties
		expectedMaster bool
		expectedAgent  bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					VnetSubnetID: "testSubnet",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						VnetSubnetID: "testSubnet",
					},
				},
			},
			expectedMaster: true,
			expectedAgent:  true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count: 1,
					},
					{
						Count: 1,
					},
				},
			},
			expectedMaster: false,
			expectedAgent:  false,
		},
	}

	for _, c := range cases {
		if c.p.MasterProfile.IsCustomVNET() != c.expectedMaster {
			t.Fatalf("expected IsCustomVnet() to return %t but instead returned %t", c.expectedMaster, c.p.MasterProfile.IsCustomVNET())
		}
		if c.p.AgentPoolProfiles[0].IsCustomVNET() != c.expectedAgent {
			t.Fatalf("expected IsCustomVnet() to return %t but instead returned %t", c.expectedAgent, c.p.AgentPoolProfiles[0].IsCustomVNET())
		}
	}

}

func TestHasAvailabilityZones(t *testing.T) {
	cases := []struct {
		p                Properties
		expectedMaster   bool
		expectedAgent    bool
		expectedAllZones bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:             1,
					AvailabilityZones: []string{"1", "2"},
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:             1,
						AvailabilityZones: []string{"1", "2"},
					},
					{
						Count:             1,
						AvailabilityZones: []string{"1", "2"},
					},
				},
			},
			expectedMaster:   true,
			expectedAgent:    true,
			expectedAllZones: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count: 1,
					},
					{
						Count:             1,
						AvailabilityZones: []string{"1", "2"},
					},
				},
			},
			expectedMaster:   false,
			expectedAgent:    false,
			expectedAllZones: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:             1,
						AvailabilityZones: []string{},
					},
					{
						Count:             1,
						AvailabilityZones: []string{"1", "2"},
					},
				},
			},
			expectedMaster:   false,
			expectedAgent:    false,
			expectedAllZones: false,
		},
	}

	for _, c := range cases {
		if c.p.MasterProfile.HasAvailabilityZones() != c.expectedMaster {
			t.Fatalf("expected HasAvailabilityZones() to return %t but instead returned %t", c.expectedMaster, c.p.MasterProfile.HasAvailabilityZones())
		}
		if c.p.AgentPoolProfiles[0].HasAvailabilityZones() != c.expectedAgent {
			t.Fatalf("expected HasAvailabilityZones() to return %t but instead returned %t", c.expectedAgent, c.p.AgentPoolProfiles[0].HasAvailabilityZones())
		}
		if c.p.HasZonesForAllAgentPools() != c.expectedAllZones {
			t.Fatalf("expected HasZonesForAllAgentPools() to return %t but instead returned %t", c.expectedAllZones, c.p.HasZonesForAllAgentPools())
		}
	}
}

func TestRequireRouteTable(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
				},
			},
			expected: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: "",
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						NetworkPlugin: "azure",
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: "cilium",
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.OrchestratorProfile.RequireRouteTable() != c.expected {
			t.Fatalf("expected RequireRouteTable() to return %t but instead got %t", c.expected, c.p.OrchestratorProfile.RequireRouteTable())
		}
	}
}

func TestIsAzureCNI(t *testing.T) {
	k := &KubernetesConfig{
		NetworkPlugin: "azure",
	}

	o := &OrchestratorProfile{
		KubernetesConfig: k,
	}
	if !o.IsAzureCNI() {
		t.Fatalf("unable to detect orchestrator profile is using Azure CNI from NetworkPlugin=%s", o.KubernetesConfig.NetworkPlugin)
	}

	k = &KubernetesConfig{
		NetworkPlugin: "none",
	}

	o = &OrchestratorProfile{
		KubernetesConfig: k,
	}
	if o.IsAzureCNI() {
		t.Fatalf("unable to detect orchestrator profile is not using Azure CNI from NetworkPlugin=%s", o.KubernetesConfig.NetworkPlugin)
	}

	o = &OrchestratorProfile{}
	if o.IsAzureCNI() {
		t.Fatalf("unable to detect orchestrator profile is not using Azure CNI from nil KubernetesConfig")
	}
}

func TestOrchestrator(t *testing.T) {
	cases := []struct {
		p                    Properties
		expectedIsDCOS       bool
		expectedIsKubernetes bool
		expectedIsSwarmMode  bool
	}{
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
				},
			},
			expectedIsDCOS:       true,
			expectedIsKubernetes: false,
			expectedIsSwarmMode:  false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
			},
			expectedIsDCOS:       false,
			expectedIsKubernetes: true,
			expectedIsSwarmMode:  false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: SwarmMode,
				},
			},
			expectedIsDCOS:       false,
			expectedIsKubernetes: false,
			expectedIsSwarmMode:  true,
		},
	}

	for _, c := range cases {
		if c.expectedIsDCOS != c.p.OrchestratorProfile.IsDCOS() {
			t.Fatalf("Expected IsDCOS() to be %t with OrchestratorType=%s", c.expectedIsDCOS, c.p.OrchestratorProfile.OrchestratorType)
		}
		if c.expectedIsKubernetes != c.p.OrchestratorProfile.IsKubernetes() {
			t.Fatalf("Expected IsKubernetes() to be %t with OrchestratorType=%s", c.expectedIsKubernetes, c.p.OrchestratorProfile.OrchestratorType)
		}
		if c.expectedIsSwarmMode != c.p.OrchestratorProfile.IsSwarmMode() {
			t.Fatalf("Expected IsSwarmMode() to be %t with OrchestratorType=%s", c.expectedIsSwarmMode, c.p.OrchestratorProfile.OrchestratorType)
		}
	}
}

func TestWindowsProfile(t *testing.T) {
	w := WindowsProfile{}

	if w.HasSecrets() || w.HasCustomImage() {
		t.Fatalf("Expected HasSecrets() and HasCustomImage() to return false when WindowsProfile is empty")
	}

	dv := w.GetWindowsDockerVersion()
	if dv != KubernetesWindowsDockerVersion {
		t.Fatalf("Expected GetWindowsDockerVersion() to equal default KubernetesWindowsDockerVersion, got %s", dv)
	}

	windowsSku := w.GetWindowsSku()
	if windowsSku != KubernetesDefaultWindowsSku {
		t.Fatalf("Expected GetWindowsSku() to equal default KubernetesDefaultWindowsSku, got %s", windowsSku)
	}

	w = WindowsProfile{
		Secrets: []KeyVaultSecrets{
			{
				SourceVault: &KeyVaultID{"testVault"},
				VaultCertificates: []KeyVaultCertificate{
					{
						CertificateURL:   "testURL",
						CertificateStore: "testStore",
					},
				},
			},
		},
		WindowsImageSourceURL: "testCustomImage",
	}

	if !(w.HasSecrets() && w.HasCustomImage()) {
		t.Fatalf("Expected HasSecrets() and HasCustomImage() to return true")
	}

	w = WindowsProfile{
		WindowsDockerVersion: "18.03.1-ee-3",
		WindowsSku:           "Datacenter-Core-1809-with-Containers-smalldisk",
	}

	dv = w.GetWindowsDockerVersion()
	if dv != "18.03.1-ee-3" {
		t.Fatalf("Expected GetWindowsDockerVersion() to equal 18.03.1-ee-3, got %s", dv)
	}

	windowsSku = w.GetWindowsSku()
	if windowsSku != "Datacenter-Core-1809-with-Containers-smalldisk" {
		t.Fatalf("Expected GetWindowsSku() to equal Datacenter-Core-1809-with-Containers-smalldisk, got %s", windowsSku)
	}
}

func TestLinuxProfile(t *testing.T) {
	l := LinuxProfile{}

	if l.HasSecrets() || l.HasSearchDomain() || l.HasCustomNodesDNS() {
		t.Fatalf("Expected HasSecrets(), HasSearchDomain() and HasCustomNodesDNS() to return false when LinuxProfile is empty")
	}

	l = LinuxProfile{
		Secrets: []KeyVaultSecrets{
			{
				SourceVault: &KeyVaultID{"testVault"},
				VaultCertificates: []KeyVaultCertificate{
					{
						CertificateURL:   "testURL",
						CertificateStore: "testStore",
					},
				},
			},
		},
		CustomNodesDNS: &CustomNodesDNS{
			DNSServer: "testDNSServer",
		},
		CustomSearchDomain: &CustomSearchDomain{
			Name:          "testName",
			RealmPassword: "testRealmPassword",
			RealmUser:     "testRealmUser",
		},
	}

	if !(l.HasSecrets() && l.HasSearchDomain() && l.HasCustomNodesDNS()) {
		t.Fatalf("Expected HasSecrets(), HasSearchDomain() and HasCustomNodesDNS() to return true")
	}
}

func TestGetAPIServerEtcdAPIVersion(t *testing.T) {
	o := OrchestratorProfile{}

	if o.GetAPIServerEtcdAPIVersion() != "" {
		t.Fatalf("Expected GetAPIServerEtcdAPIVersion() to return \"\" but instead got %s", o.GetAPIServerEtcdAPIVersion())
	}

	o.KubernetesConfig = &KubernetesConfig{
		EtcdVersion: "3.2.1",
	}

	if o.GetAPIServerEtcdAPIVersion() != "etcd3" {
		t.Fatalf("Expected GetAPIServerEtcdAPIVersion() to return \"etcd3\" but instead got %s", o.GetAPIServerEtcdAPIVersion())
	}

	// invalid version string
	o.KubernetesConfig.EtcdVersion = "2.3.8"
	if o.GetAPIServerEtcdAPIVersion() != "etcd2" {
		t.Fatalf("Expected GetAPIServerEtcdAPIVersion() to return \"etcd2\" but instead got %s", o.GetAPIServerEtcdAPIVersion())
	}
}

func TestHasAadProfile(t *testing.T) {
	p := Properties{}

	if p.HasAadProfile() {
		t.Fatalf("Expected HasAadProfile() to return false")
	}

	p.AADProfile = &AADProfile{
		ClientAppID: "test",
		ServerAppID: "test",
	}

	if !p.HasAadProfile() {
		t.Fatalf("Expected HasAadProfile() to return true")
	}

}

func TestCustomHyperkubeImageField(t *testing.T) {
	log.Println(exampleAPIModel)
	apiloader := &Apiloader{
		Translator: nil,
	}
	apimodel, _, err := apiloader.DeserializeContainerService([]byte(exampleAPIModel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpectedly error deserializing the example apimodel: %s", err)
	}

	actualCustomHyperkubeImage := apimodel.Properties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage
	if actualCustomHyperkubeImage != exampleCustomHyperkubeImage {
		t.Fatalf("kubernetesConfig->customHyperkubeImage field value was unexpected: got(%s), expected(%s)", actualCustomHyperkubeImage, exampleCustomHyperkubeImage)
	}
}

func TestUserAssignedMSI(t *testing.T) {
	// Test1: With just System MSI
	log.Println(exampleSystemMSIModel)
	apiloader := &Apiloader{
		Translator: nil,
	}
	apiModel, _, err := apiloader.DeserializeContainerService([]byte(exampleSystemMSIModel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserailizing the example user msi api model: %s", err)
	}
	systemMSI := apiModel.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	actualUserMSI := apiModel.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID
	if !systemMSI || actualUserMSI != "" {
		t.Fatalf("found user msi: %t and usermsi: %s", systemMSI, actualUserMSI)
	}

	// Test2: With user assigned MSI
	log.Println(exampleUserMSIModel)
	apiloader = &Apiloader{
		Translator: nil,
	}
	apiModel, _, err = apiloader.DeserializeContainerService([]byte(exampleUserMSIModel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserailizing the example user msi api model: %s", err)
	}
	systemMSI = apiModel.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	actualUserMSI = apiModel.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID
	if !systemMSI && actualUserMSI != exampleUserMSI {
		t.Fatalf("found user msi: %t and usermsi: %s", systemMSI, actualUserMSI)
	}
}

func TestKubernetesAddon(t *testing.T) {
	addon := getMockAddon("addon")
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

func TestIsTillerEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsTillerEnabled()
	enabledDefault := DefaultTillerAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsTillerEnabled() should return %t when no tiller addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(DefaultTillerAddonName))
	enabled = c.IsTillerEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsTillerEnabled() should return default when a custom tiller addon has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultTillerAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsTillerEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsTillerEnabled() should return true when a custom tiller addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultTillerAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsTillerEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsTillerEnabled() should return false when a custom tiller addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsAADPodIdentityEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsAADPodIdentityEnabled()
	enabledDefault := DefaultAADPodIdentityAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsAADPodIdentityEnabled() should return %t when no aad pod identity addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(DefaultAADPodIdentityAddonName))
	enabled = c.IsAADPodIdentityEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsAADPodIdentityEnabled() should return default when aad pod identity addon has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultAADPodIdentityAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsAADPodIdentityEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsAADPodIdentityEnabled() should return true when aad pod identity addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultAADPodIdentityAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsAADPodIdentityEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsAADPodIdentityEnabled() should return false when aad pod identity addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsACIConnectorEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsACIConnectorEnabled()
	enabledDefault := DefaultACIConnectorAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsACIConnectorEnabled() should return %t when no ACI connector addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(DefaultACIConnectorAddonName))
	enabled = c.IsACIConnectorEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsACIConnectorEnabled() should return default when ACI connector has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultACIConnectorAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsACIConnectorEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsACIConnectorEnabled() should return true when ACI connector addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultACIConnectorAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsACIConnectorEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsACIConnectorEnabled() should return false when ACI connector addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsClusterAutoscalerEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsClusterAutoscalerEnabled()
	enabledDefault := DefaultClusterAutoscalerAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsClusterAutoscalerEnabled() should return %t when no cluster autoscaler addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(DefaultClusterAutoscalerAddonName))
	enabled = c.IsClusterAutoscalerEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsClusterAutoscalerEnabled() should return default when cluster autoscaler has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultClusterAutoscalerAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsClusterAutoscalerEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsClusterAutoscalerEnabled() should return true when cluster autoscaler addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultClusterAutoscalerAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsClusterAutoscalerEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsClusterAutoscalerEnabled() should return false when cluster autoscaler addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsBlobfuseFlexVolumeEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsBlobfuseFlexVolumeEnabled()
	enabledDefault := DefaultBlobfuseFlexVolumeAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsBlobfuseFlexVolumeEnabled() should return %t when no blobfuse flexvolume addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(DefaultBlobfuseFlexVolumeAddonName))
	enabled = c.IsBlobfuseFlexVolumeEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsBlobfuseFlexVolumeEnabled() should return default when blobfuse flexvolume has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultBlobfuseFlexVolumeAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsBlobfuseFlexVolumeEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsBlobfuseFlexVolumeEnabled() should return true when blobfuse flexvolume addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultBlobfuseFlexVolumeAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsBlobfuseFlexVolumeEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsBlobfuseFlexVolumeEnabled() should return false when blobfuse flexvolume addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsSMBFlexVolumeEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsSMBFlexVolumeEnabled()
	enabledDefault := DefaultSMBFlexVolumeAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsSMBFlexVolumeEnabled() should return %t when no SMB flexvolume addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(DefaultSMBFlexVolumeAddonName))
	enabled = c.IsSMBFlexVolumeEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsSMBFlexVolumeEnabled() should return default when SMB flexvolume has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultSMBFlexVolumeAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsSMBFlexVolumeEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsSMBFlexVolumeEnabled() should return true when SMB flexvolume addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultSMBFlexVolumeAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsSMBFlexVolumeEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsSMBFlexVolumeEnabled() should return true when SMB flexvolume addon has been specified as enabled, instead returned %t", enabled)
	}
}

func TestIsKeyVaultFlexVolumeEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsKeyVaultFlexVolumeEnabled()
	enabledDefault := DefaultKeyVaultFlexVolumeAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsKeyVaultFlexVolumeEnabled() should return %t when no key vault flexvolume addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(DefaultKeyVaultFlexVolumeAddonName))
	enabled = c.IsKeyVaultFlexVolumeEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsKeyVaultFlexVolumeEnabled() should return default when no keyvault flexvolume has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultKeyVaultFlexVolumeAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsKeyVaultFlexVolumeEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsKeyVaultFlexVolumeEnabled() should return true when keyvault flexvolume addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultKeyVaultFlexVolumeAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsKeyVaultFlexVolumeEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsKeyVaultFlexVolumeEnabled() should return false when keyvault flexvolume addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsNVIDIADevicePluginEnabled(t *testing.T) {
	p := Properties{
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:   "agentpool",
				VMSize: "Standard_N",
				Count:  1,
			},
		},
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    Kubernetes,
			OrchestratorVersion: "1.9.0",
			KubernetesConfig: &KubernetesConfig{
				Addons: []KubernetesAddon{
					getMockAddon("addon"),
				},
			},
		},
	}

	if !p.HasNSeriesSKU() {
		t.Fatalf("HasNSeriesSKU should return true when explicitly using VM Size %s", p.AgentPoolProfiles[0].VMSize)
	}
	if p.IsNVIDIADevicePluginEnabled() {
		t.Fatalf("KubernetesConfig.IsNVIDIADevicePluginEnabled() should return false with N-series VMs with < k8s 1.10, instead returned %t", p.IsNVIDIADevicePluginEnabled())
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.10.0"
	if !p.IsNVIDIADevicePluginEnabled() {
		t.Fatalf("KubernetesConfig.IsNVIDIADevicePluginEnabled() should return true with N-series VMs with k8s >= 1.10, instead returned %t", p.IsNVIDIADevicePluginEnabled())
	}

	p.AgentPoolProfiles[0].VMSize = "Standard_D2_v2"
	p.OrchestratorProfile.KubernetesConfig.Addons = []KubernetesAddon{
		{
			Name:    NVIDIADevicePluginAddonName,
			Enabled: helpers.PointerToBool(false),
		},
	}

	if p.HasNSeriesSKU() {
		t.Fatalf("HasNSeriesSKU should return false when explicitly using VM Size %s", p.AgentPoolProfiles[0].VMSize)
	}
	if p.IsNVIDIADevicePluginEnabled() {
		t.Fatalf("KubernetesConfig.IsNVIDIADevicePluginEnabled() should return false when explicitly disabled")
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
				OrchestratorVersion: "1.12.2",
			},
		}
		ret := p.AgentPoolProfiles[0].IsNSeriesSKU()
		if ret != c.Expected {
			t.Fatalf("expected IsNvidiaEnabledSKU(%s) to return %t, but instead got %t", c.VMSKU, c.Expected, ret)
		}
	}
}

func TestIsContainerMonitoringEnabled(t *testing.T) {
	v := "1.9.0"
	o := OrchestratorProfile{
		OrchestratorType:    "Kubernetes",
		OrchestratorVersion: v,
		KubernetesConfig: &KubernetesConfig{Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
		},
	}
	enabled := o.KubernetesConfig.IsContainerMonitoringEnabled()
	enabledDefault := DefaultContainerMonitoringAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringEnabled() should return %t for kubernetes version %s when no container-monitoring addon has been specified, instead returned %t", enabledDefault, v, enabled)
	}

	b := true
	cm := getMockAddon(ContainerMonitoringAddonName)
	cm.Enabled = &b
	o.KubernetesConfig.Addons = append(o.KubernetesConfig.Addons, cm)
	enabled = o.KubernetesConfig.IsContainerMonitoringEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringEnabled() should return %t for kubernetes version %s when the container-monitoring addon has been specified, instead returned %t", true, v, enabled)
	}

	b = false
	o = OrchestratorProfile{
		OrchestratorType:    "Kubernetes",
		OrchestratorVersion: v,
		KubernetesConfig: &KubernetesConfig{Addons: []KubernetesAddon{
			{
				Name:    ContainerMonitoringAddonName,
				Enabled: &b,
			},
		},
		},
	}
	enabled = o.KubernetesConfig.IsContainerMonitoringEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringEnabled() should return false when a custom container monitoring addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsDashboardEnabled(t *testing.T) {
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsDashboardEnabled()
	enabledDefault := DefaultDashboardAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsDashboardEnabled() should return %t when no kubernetes-dashboard addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	c.Addons = append(c.Addons, getMockAddon(DefaultDashboardAddonName))
	enabled = c.IsDashboardEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsDashboardEnabled() should return true when a custom kubernetes-dashboard addon has been specified, instead returned %t", enabled)
	}
	b := false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultDashboardAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsDashboardEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsDashboardEnabled() should return false when a custom kubernetes-dashboard addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsReschedulerEnabled(t *testing.T) {
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsReschedulerEnabled()
	enabledDefault := DefaultReschedulerAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsReschedulerEnabled() should return %t when no rescheduler addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	c.Addons = append(c.Addons, getMockAddon(DefaultReschedulerAddonName))
	enabled = c.IsReschedulerEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsReschedulerEnabled() should return true when a custom rescheduler addon has been specified, instead returned %t", enabled)
	}
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    DefaultReschedulerAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsReschedulerEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsReschedulerEnabled() should return false when a custom rescheduler addon has been specified as enabled, instead returned %t", enabled)
	}
}

func TestIsMetricsServerEnabled(t *testing.T) {
	v := "1.8.0"
	o := OrchestratorProfile{
		OrchestratorType:    "Kubernetes",
		OrchestratorVersion: v,
		KubernetesConfig: &KubernetesConfig{Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
		},
	}
	enabled := o.IsMetricsServerEnabled()
	enabledDefault := DefaultMetricsServerAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsMetricsServerEnabled() should return %t for kubernetes version %s when no metrics-server addon has been specified, instead returned %t", enabledDefault, v, enabled)
	}

	o.KubernetesConfig.Addons = append(o.KubernetesConfig.Addons, getMockAddon(DefaultMetricsServerAddonName))
	enabled = o.IsMetricsServerEnabled()
	enabledDefault = DefaultMetricsServerAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsMetricsServerEnabled() should return %t for kubernetes version %s when the metrics-server addon has been specified, instead returned %t", enabledDefault, v, enabled)
	}

	b := true
	o = OrchestratorProfile{
		OrchestratorType:    "Kubernetes",
		OrchestratorVersion: v,
		KubernetesConfig: &KubernetesConfig{Addons: []KubernetesAddon{
			{
				Name:    DefaultMetricsServerAddonName,
				Enabled: &b,
			},
		},
		},
	}
	enabled = o.IsMetricsServerEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsMetricsServerEnabled() should return true for kubernetes version %s when the metrics-server addon has been specified as enabled, instead returned %t", v, enabled)
	}
}

func TestIsIPMasqAgentEnabled(t *testing.T) {
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsIPMasqAgentEnabled()
	enabledDefault := IPMasqAgentAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsIPMasqAgentEnabled() should return %t when no ip-masq-agent addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	c.Addons = append(c.Addons, getMockAddon(IPMASQAgentAddonName))
	enabled = c.IsIPMasqAgentEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsIPMasqAgentEnabled() should return true when ip-masq-agent adddon has been specified, instead returned %t", enabled)
	}
	b := false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    IPMASQAgentAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsIPMasqAgentEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsIPMasqAgentEnabled() should return false when ip-masq-agent addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestGetAzureCNIURLFuncs(t *testing.T) {
	// Default case
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 1, 3, false)
	cs.Location = "eastus"
	cloudSpecConfig := cs.GetCloudSpecConfig()

	o := OrchestratorProfile{
		OrchestratorType: "Kubernetes",
		KubernetesConfig: &KubernetesConfig{},
	}
	linuxURL := o.KubernetesConfig.GetAzureCNIURLLinux(cloudSpecConfig)
	windowsURL := o.KubernetesConfig.GetAzureCNIURLWindows(cloudSpecConfig)
	if linuxURL != cloudSpecConfig.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL {
		t.Fatalf("GetAzureCNIURLLinux() should return default %s, instead returned %s", cloudSpecConfig.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL, linuxURL)
	}
	if windowsURL != cloudSpecConfig.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL {
		t.Fatalf("GetAzureCNIURLWindows() should return default %s, instead returned %s", cloudSpecConfig.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL, windowsURL)
	}

	// User-configurable case
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 1, 3, false)
	cs.Location = "eastus"
	cloudSpecConfig = cs.GetCloudSpecConfig()

	customLinuxURL := "https://custom-url/azure-cni-linux.0.0.1.tgz"
	customWindowsURL := "https://custom-url/azure-cni-windows.0.0.1.tgz"
	o = OrchestratorProfile{
		OrchestratorType: "Kubernetes",
		KubernetesConfig: &KubernetesConfig{
			AzureCNIURLLinux:   customLinuxURL,
			AzureCNIURLWindows: customWindowsURL,
		},
	}

	linuxURL = o.KubernetesConfig.GetAzureCNIURLLinux(cloudSpecConfig)
	windowsURL = o.KubernetesConfig.GetAzureCNIURLWindows(cloudSpecConfig)
	if linuxURL != customLinuxURL {
		t.Fatalf("GetAzureCNIURLLinux() should return custom URL %s, instead returned %s", customLinuxURL, linuxURL)
	}
	if windowsURL != customWindowsURL {
		t.Fatalf("GetAzureCNIURLWindows() should return custom URL %s, instead returned %s", customWindowsURL, windowsURL)
	}
}

func TestCloudProviderDefaults(t *testing.T) {
	// Test cloudprovider defaults when no user-provided values
	v := "1.8.0"
	o := OrchestratorProfile{
		OrchestratorType:    "Kubernetes",
		OrchestratorVersion: v,
		KubernetesConfig:    &KubernetesConfig{},
	}
	o.KubernetesConfig.SetCloudProviderBackoffDefaults()
	o.KubernetesConfig.SetCloudProviderRateLimitDefaults()

	intCases := []struct {
		defaultVal  int
		computedVal int
	}{
		{
			defaultVal:  DefaultKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			defaultVal:  DefaultKubernetesCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			defaultVal:  DefaultKubernetesCloudProviderRateLimitBucket,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCases {
		if c.computedVal != c.defaultVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.defaultVal, c.computedVal)
		}
	}

	floatCases := []struct {
		defaultVal  float64
		computedVal float64
	}{
		{
			defaultVal:  DefaultKubernetesCloudProviderBackoffJitter,
			computedVal: o.KubernetesConfig.CloudProviderBackoffJitter,
		},
		{
			defaultVal:  DefaultKubernetesCloudProviderBackoffExponent,
			computedVal: o.KubernetesConfig.CloudProviderBackoffExponent,
		},
		{
			defaultVal:  DefaultKubernetesCloudProviderRateLimitQPS,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCases {
		if c.computedVal != c.defaultVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.defaultVal, c.computedVal)
		}
	}

	customCloudProviderBackoffDuration := 99
	customCloudProviderBackoffExponent := 10.0
	customCloudProviderBackoffJitter := 11.9
	customCloudProviderBackoffRetries := 9
	customCloudProviderRateLimitBucket := 37
	customCloudProviderRateLimitQPS := 9.9

	// Test cloudprovider defaults when user provides configuration
	v = "1.8.0"
	o = OrchestratorProfile{
		OrchestratorType:    "Kubernetes",
		OrchestratorVersion: v,
		KubernetesConfig: &KubernetesConfig{
			CloudProviderBackoffDuration: customCloudProviderBackoffDuration,
			CloudProviderBackoffExponent: customCloudProviderBackoffExponent,
			CloudProviderBackoffJitter:   customCloudProviderBackoffJitter,
			CloudProviderBackoffRetries:  customCloudProviderBackoffRetries,
			CloudProviderRateLimitBucket: customCloudProviderRateLimitBucket,
			CloudProviderRateLimitQPS:    customCloudProviderRateLimitQPS,
		},
	}
	o.KubernetesConfig.SetCloudProviderBackoffDefaults()
	o.KubernetesConfig.SetCloudProviderRateLimitDefaults()

	intCasesCustom := []struct {
		customVal   int
		computedVal int
	}{
		{
			customVal:   customCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			customVal:   customCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			customVal:   customCloudProviderRateLimitBucket,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCasesCustom {
		if c.computedVal != c.customVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.customVal, c.computedVal)
		}
	}

	floatCasesCustom := []struct {
		customVal   float64
		computedVal float64
	}{
		{
			customVal:   customCloudProviderBackoffJitter,
			computedVal: o.KubernetesConfig.CloudProviderBackoffJitter,
		},
		{
			customVal:   customCloudProviderBackoffExponent,
			computedVal: o.KubernetesConfig.CloudProviderBackoffExponent,
		},
		{
			customVal:   customCloudProviderRateLimitQPS,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesCustom {
		if c.computedVal != c.customVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.customVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults when user provides *some* config values
	v = "1.8.0"
	o = OrchestratorProfile{
		OrchestratorType:    "Kubernetes",
		OrchestratorVersion: v,
		KubernetesConfig: &KubernetesConfig{
			CloudProviderBackoffDuration: customCloudProviderBackoffDuration,
			CloudProviderRateLimitBucket: customCloudProviderRateLimitBucket,
			CloudProviderRateLimitQPS:    customCloudProviderRateLimitQPS,
		},
	}
	o.KubernetesConfig.SetCloudProviderBackoffDefaults()
	o.KubernetesConfig.SetCloudProviderRateLimitDefaults()

	intCasesMixed := []struct {
		expectedVal int
		computedVal int
	}{
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			expectedVal: customCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			expectedVal: customCloudProviderRateLimitBucket,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.expectedVal, c.computedVal)
		}
	}

	floatCasesMixed := []struct {
		expectedVal float64
		computedVal float64
	}{
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffJitter,
			computedVal: o.KubernetesConfig.CloudProviderBackoffJitter,
		},
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffExponent,
			computedVal: o.KubernetesConfig.CloudProviderBackoffExponent,
		},
		{
			expectedVal: customCloudProviderRateLimitQPS,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}
}

func getMockAddon(name string) KubernetesAddon {
	return KubernetesAddon{
		Name: name,
		Containers: []KubernetesContainerSpec{
			{
				Name:           name,
				CPURequests:    "50m",
				MemoryRequests: "150Mi",
				CPULimits:      "50m",
				MemoryLimits:   "150Mi",
			},
		},
	}
}

func TestAreAgentProfilesCustomVNET(t *testing.T) {
	p := Properties{}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			VnetSubnetID: "subnetlink1",
		},
		{
			VnetSubnetID: "subnetlink2",
		},
	}

	if !p.AreAgentProfilesCustomVNET() {
		t.Fatalf("Expected isCustomVNET to be true when subnet exists for all agent pool profile")
	}

	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			VnetSubnetID: "subnetlink1",
		},
		{
			VnetSubnetID: "",
		},
	}

	if p.AreAgentProfilesCustomVNET() {
		t.Fatalf("Expected isCustomVNET to be false when subnet exists for some agent pool profile")
	}

	p.AgentPoolProfiles = nil

	if p.AreAgentProfilesCustomVNET() {
		t.Fatalf("Expected isCustomVNET to be false when agent pool profiles is nil")
	}
}

func TestGenerateClusterID(t *testing.T) {
	tests := []struct {
		name              string
		properties        *Properties
		expectedClusterID string
	}{
		{
			name: "From Master Profile",
			properties: &Properties{
				MasterProfile: &MasterProfile{
					DNSPrefix: "foo_master",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo_agent0",
					},
				},
			},
			expectedClusterID: "24569115",
		},
		{
			name: "From Hosted Master Profile",
			properties: &Properties{
				HostedMasterProfile: &HostedMasterProfile{
					DNSPrefix: "foo_hosted_master",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo_agent1",
					},
				},
			},
			expectedClusterID: "42761241",
		},
		{
			name: "No Master Profile",
			properties: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo_agent2",
					},
				},
			},
			expectedClusterID: "11729301",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.properties.GetClusterID()

			if actual != test.expectedClusterID {
				t.Errorf("expected cluster ID %s, but got %s", test.expectedClusterID, actual)
			}
		})
	}
}

func TestGetPrimaryAvailabilitySetName(t *testing.T) {
	p := &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		MasterProfile: &MasterProfile{
			Count:     1,
			DNSPrefix: "foo",
			VMSize:    "Standard_DS2_v2",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: AvailabilitySet,
			},
		},
	}

	expected := "agentpool-availabilitySet-28513887"
	got := p.GetPrimaryAvailabilitySetName()
	if got != expected {
		t.Errorf("expected primary availability set name %s, but got %s", expected, got)
	}
}

func TestGetPrimaryScaleSetName(t *testing.T) {
	p := &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		MasterProfile: &MasterProfile{
			Count:     1,
			DNSPrefix: "foo",
			VMSize:    "Standard_DS2_v2",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
	}

	expected := "k8s-agentpool-28513887-vmss"
	got := p.GetPrimaryScaleSetName()
	if got != expected {
		t.Errorf("expected primary availability set name %s, but got %s", expected, got)
	}
}

func TestGetRouteTableName(t *testing.T) {
	p := &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		HostedMasterProfile: &HostedMasterProfile{
			FQDN:      "fqdn",
			DNSPrefix: "foo",
			Subnet:    "mastersubnet",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
	}

	actualRTName := p.GetRouteTableName()
	expectedRTName := "aks-agentpool-28513887-routetable"

	actualNSGName := p.GetNSGName()
	expectedNSGName := "aks-agentpool-28513887-nsg"

	if actualRTName != expectedRTName {
		t.Errorf("expected route table name %s, but got %s", expectedRTName, actualRTName)
	}

	if actualNSGName != expectedNSGName {
		t.Errorf("expected route table name %s, but got %s", expectedNSGName, actualNSGName)
	}

	p = &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		MasterProfile: &MasterProfile{
			Count:     1,
			DNSPrefix: "foo",
			VMSize:    "Standard_DS2_v2",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
	}

	actualRTName = p.GetRouteTableName()
	expectedRTName = "k8s-master-28513887-routetable"

	actualNSGName = p.GetNSGName()
	expectedNSGName = "k8s-master-28513887-nsg"

	if actualRTName != expectedRTName {
		t.Errorf("expected route table name %s, but got %s", actualRTName, expectedRTName)
	}

	if actualNSGName != expectedNSGName {
		t.Errorf("expected route table name %s, but got %s", actualNSGName, expectedNSGName)
	}
}

func TestGetSubnetName(t *testing.T) {
	tests := []struct {
		name               string
		properties         *Properties
		expectedSubnetName string
	}{
		{
			name: "Cluster with HosterMasterProfile",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				HostedMasterProfile: &HostedMasterProfile{
					FQDN:      "fqdn",
					DNSPrefix: "foo",
					Subnet:    "mastersubnet",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool",
						VMSize:              "Standard_D2_v2",
						Count:               1,
						AvailabilityProfile: VirtualMachineScaleSets,
					},
				},
			},
			expectedSubnetName: "aks-subnet",
		},
		{
			name: "Cluster with HosterMasterProfile and custom VNET",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				HostedMasterProfile: &HostedMasterProfile{
					FQDN:      "fqdn",
					DNSPrefix: "foo",
					Subnet:    "mastersubnet",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool",
						VMSize:              "Standard_D2_v2",
						Count:               1,
						AvailabilityProfile: VirtualMachineScaleSets,
						VnetSubnetID:        "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/ExampleCustomVNET/subnets/BazAgentSubnet",
					},
				},
			},
			expectedSubnetName: "BazAgentSubnet",
		},
		{
			name: "Cluster with MasterProfile",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:     1,
					DNSPrefix: "foo",
					VMSize:    "Standard_DS2_v2",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool",
						VMSize:              "Standard_D2_v2",
						Count:               1,
						AvailabilityProfile: VirtualMachineScaleSets,
					},
				},
			},
			expectedSubnetName: "k8s-subnet",
		},
		{
			name: "Cluster with MasterProfile and custom VNET",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:        1,
					DNSPrefix:    "foo",
					VMSize:       "Standard_DS2_v2",
					VnetSubnetID: "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/ExampleCustomVNET/subnets/BazAgentSubnet",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool",
						VMSize:              "Standard_D2_v2",
						Count:               1,
						AvailabilityProfile: VirtualMachineScaleSets,
					},
				},
			},
			expectedSubnetName: "BazAgentSubnet",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.properties.GetSubnetName()

			if actual != test.expectedSubnetName {
				t.Errorf("expected subnet name %s, but got %s", test.expectedSubnetName, actual)
			}
		})
	}
}

func TestProperties_GetVirtualNetworkName(t *testing.T) {
	tests := []struct {
		name                       string
		properties                 *Properties
		expectedVirtualNetworkName string
	}{
		{
			name: "Cluster with HostedMasterProfile and Custom VNET AgentProfiles",
			properties: &Properties{
				HostedMasterProfile: &HostedMasterProfile{
					FQDN:      "fqdn",
					DNSPrefix: "foo",
					Subnet:    "mastersubnet",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool",
						VMSize:              "Standard_D2_v2",
						Count:               1,
						AvailabilityProfile: VirtualMachineScaleSets,
						VnetSubnetID:        "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/ExampleCustomVNET/subnets/BazAgentSubnet",
					},
				},
			},
			expectedVirtualNetworkName: "ExampleCustomVNET",
		},
		{
			name: "Cluster with HostedMasterProfile and AgentProfiles",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				HostedMasterProfile: &HostedMasterProfile{
					FQDN:      "fqdn",
					DNSPrefix: "foo",
					Subnet:    "mastersubnet",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool",
						VMSize:              "Standard_D2_v2",
						Count:               1,
						AvailabilityProfile: VirtualMachineScaleSets,
					},
				},
			},
			expectedVirtualNetworkName: "aks-vnet-28513887",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.properties.GetVirtualNetworkName()

			if actual != test.expectedVirtualNetworkName {
				t.Errorf("expected virtual network name %s, but got %s", test.expectedVirtualNetworkName, actual)
			}
		})
	}
}

func TestProperties_GetVNetResourceGroupName(t *testing.T) {
	p := &Properties{
		HostedMasterProfile: &HostedMasterProfile{
			FQDN:      "fqdn",
			DNSPrefix: "foo",
			Subnet:    "mastersubnet",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: VirtualMachineScaleSets,
				VnetSubnetID:        "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/ExampleCustomVNET/subnets/BazAgentSubnet",
			},
		},
	}
	expectedVNETResourceGroupName := "RESOURCE_GROUP_NAME"

	actual := p.GetVNetResourceGroupName()

	if expectedVNETResourceGroupName != actual {
		t.Errorf("expected vnet resource group name name %s, but got %s", expectedVNETResourceGroupName, actual)
	}
}

func TestProperties_GetClusterMetadata(t *testing.T) {
	p := &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		MasterProfile: &MasterProfile{
			Count:        1,
			DNSPrefix:    "foo",
			VMSize:       "Standard_DS2_v2",
			VnetSubnetID: "/subscriptions/SUBSCRIPTION_ID/resourceGroups/SAMPLE_RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/ExampleCustomVNET/subnets/BazAgentSubnet",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: AvailabilitySet,
			},
		},
	}

	metadata := p.GetClusterMetadata()

	if metadata == nil {
		t.Error("did not expect cluster metadata to be nil")
	}

	expectedSubnetName := "BazAgentSubnet"
	if metadata.SubnetName != expectedSubnetName {
		t.Errorf("expected subnet name %s, but got %s", expectedSubnetName, metadata.SubnetName)
	}

	expectedVNetResourceGroupName := "SAMPLE_RESOURCE_GROUP_NAME"
	if metadata.VNetResourceGroupName != expectedVNetResourceGroupName {
		t.Errorf("expected vNetResourceGroupName name %s, but got %s", expectedVNetResourceGroupName, metadata.VNetResourceGroupName)
	}

	expectedVirtualNetworkName := "ExampleCustomVNET"
	if metadata.VirtualNetworkName != expectedVirtualNetworkName {
		t.Errorf("expected VirtualNetworkName name %s, but got %s", expectedVirtualNetworkName, metadata.VirtualNetworkName)
	}

	expectedRouteTableName := "k8s-master-28513887-routetable"
	if metadata.RouteTableName != expectedRouteTableName {
		t.Errorf("expected RouteTableName name %s, but got %s", expectedVirtualNetworkName, metadata.RouteTableName)
	}

	expectedSecurityGroupName := "k8s-master-28513887-nsg"
	if metadata.SecurityGroupName != expectedSecurityGroupName {
		t.Errorf("expected SecurityGroupName name %s, but got %s", expectedSecurityGroupName, metadata.SecurityGroupName)
	}

	expectedPrimaryAvailabilitySetName := "agentpool-availabilitySet-28513887"
	if metadata.PrimaryAvailabilitySetName != expectedPrimaryAvailabilitySetName {
		t.Errorf("expected PrimaryAvailabilitySetName name %s, but got %s", expectedPrimaryAvailabilitySetName, metadata.PrimaryAvailabilitySetName)
	}

	expectedPrimaryScaleSetName := "k8s-agentpool-28513887-vmss"
	if metadata.PrimaryScaleSetName != expectedPrimaryScaleSetName {
		t.Errorf("expected PrimaryScaleSetName name %s, but got %s", expectedPrimaryScaleSetName, metadata.PrimaryScaleSetName)
	}
}

func TestGetAddonContainersIndexByName(t *testing.T) {
	addonName := "testaddon"
	addon := getMockAddon(addonName)
	i := addon.GetAddonContainersIndexByName(addonName)
	if i != 0 {
		t.Fatalf("getAddonContainersIndexByName() did not return the expected index value 0, instead returned: %d", i)
	}
	i = addon.GetAddonContainersIndexByName("nonExistentContainerName")
	if i != -1 {
		t.Fatalf("getAddonContainersIndexByName() did not return the expected index value 0, instead returned: %d", i)
	}
}

func TestGetAgentPoolIndexByName(t *testing.T) {
	tests := []struct {
		name          string
		profileName   string
		properties    *Properties
		expectedIndex int
	}{
		{
			name:        "index 0",
			profileName: "myagentpool",
			properties: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "myagentpool",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  1,
					},
				},
			},
			expectedIndex: 0,
		},
		{
			name:        "index 3",
			profileName: "myagentpool",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:     1,
					DNSPrefix: "myprefix1",
					VMSize:    "Standard_DS2_v2",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
					{
						Name:   "agentpool2",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
					{
						Name:   "agentpool3",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
					{
						Name:   "myagentpool",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
				},
			},
			expectedIndex: 3,
		},
		{
			name:        "not found",
			profileName: "myagentpool",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:     1,
					DNSPrefix: "myprefix2",
					VMSize:    "Standard_DS2_v2",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agent1",
						VMSize: "Standard_D2_v2",
						Count:  1,
					},
				},
			},
			expectedIndex: -1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.properties.getAgentPoolIndexByName(test.profileName)

			if actual != test.expectedIndex {
				t.Errorf("expected agent pool index %d, but got %d", test.expectedIndex, actual)
			}
		})
	}
}

func TestGetAgentVMPrefix(t *testing.T) {
	tests := []struct {
		name             string
		profile          *AgentPoolProfile
		properties       *Properties
		expectedVMPrefix string
	}{
		{
			name: "Linux VMAS agent pool profile",
			profile: &AgentPoolProfile{
				Name:   "agentpool",
				VMSize: "Standard_D2_v2",
				Count:  1,
				OSType: "Linux",
			},
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:     1,
					DNSPrefix: "myprefix",
					VMSize:    "Standard_DS2_v2",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool",
						VMSize: "Standard_D2_v2",
						Count:  1,
						OSType: "Linux",
					},
				},
			},
			expectedVMPrefix: "k8s-agentpool-42378941-",
		},
		{
			name: "Linux VMSS agent pool profile",
			profile: &AgentPoolProfile{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: "VirtualMachineScaleSets",
				OSType:              "Linux",
			},
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:     1,
					DNSPrefix: "myprefix1",
					VMSize:    "Standard_DS2_v2",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool",
						VMSize:              "Standard_D2_v2",
						Count:               1,
						AvailabilityProfile: "VirtualMachineScaleSets",
						OSType:              "Linux",
					},
				},
			},
			expectedVMPrefix: "k8s-agentpool-30819786-vmss",
		},
		{
			name: "Windows agent pool profile",
			profile: &AgentPoolProfile{
				Name:   "agentpool",
				VMSize: "Standard_D2_v2",
				Count:  1,
				OSType: "Windows",
			},
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:     1,
					DNSPrefix: "myprefix2",
					VMSize:    "Standard_DS2_v2",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool",
						VMSize: "Standard_D2_v2",
						Count:  1,
						OSType: "Windows",
					},
				},
			},
			expectedVMPrefix: "2478k8s00",
		},
		{
			name: "agent profile doesn't exist",
			profile: &AgentPoolProfile{
				Name:   "something",
				VMSize: "Standard_D2_v2",
				Count:  1,
				OSType: "Windows",
			},
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:     1,
					DNSPrefix: "myprefix2",
					VMSize:    "Standard_DS2_v2",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool",
						VMSize: "Standard_D2_v2",
						Count:  1,
					},
				},
			},
			expectedVMPrefix: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			p := test.properties
			actual := p.GetAgentVMPrefix(test.profile)

			if actual != test.expectedVMPrefix {
				t.Errorf("expected agent VM name %s, but got %s", test.expectedVMPrefix, actual)
			}
		})
	}
}

func TestFormatAzureProdFQDN(t *testing.T) {
	dnsPrefix := "santest"
	var actual []string
	for _, location := range helpers.GetAzureLocations() {
		actual = append(actual, FormatAzureProdFQDNByLocation(dnsPrefix, location))
	}

	expected := []string{
		"santest.australiacentral.cloudapp.azure.com",
		"santest.australiacentral2.cloudapp.azure.com",
		"santest.australiaeast.cloudapp.azure.com",
		"santest.australiasoutheast.cloudapp.azure.com",
		"santest.brazilsouth.cloudapp.azure.com",
		"santest.canadacentral.cloudapp.azure.com",
		"santest.canadaeast.cloudapp.azure.com",
		"santest.centralindia.cloudapp.azure.com",
		"santest.centralus.cloudapp.azure.com",
		"santest.centraluseuap.cloudapp.azure.com",
		"santest.chinaeast.cloudapp.chinacloudapi.cn",
		"santest.chinaeast2.cloudapp.chinacloudapi.cn",
		"santest.chinanorth.cloudapp.chinacloudapi.cn",
		"santest.chinanorth2.cloudapp.chinacloudapi.cn",
		"santest.eastasia.cloudapp.azure.com",
		"santest.eastus.cloudapp.azure.com",
		"santest.eastus2.cloudapp.azure.com",
		"santest.eastus2euap.cloudapp.azure.com",
		"santest.francecentral.cloudapp.azure.com",
		"santest.francesouth.cloudapp.azure.com",
		"santest.japaneast.cloudapp.azure.com",
		"santest.japanwest.cloudapp.azure.com",
		"santest.koreacentral.cloudapp.azure.com",
		"santest.koreasouth.cloudapp.azure.com",
		"santest.northcentralus.cloudapp.azure.com",
		"santest.northeurope.cloudapp.azure.com",
		"santest.southcentralus.cloudapp.azure.com",
		"santest.southeastasia.cloudapp.azure.com",
		"santest.southindia.cloudapp.azure.com",
		"santest.uksouth.cloudapp.azure.com",
		"santest.ukwest.cloudapp.azure.com",
		"santest.westcentralus.cloudapp.azure.com",
		"santest.westeurope.cloudapp.azure.com",
		"santest.westindia.cloudapp.azure.com",
		"santest.westus.cloudapp.azure.com",
		"santest.westus2.cloudapp.azure.com",
		"santest.chinaeast.cloudapp.chinacloudapi.cn",
		"santest.chinanorth.cloudapp.chinacloudapi.cn",
		"santest.chinanorth2.cloudapp.chinacloudapi.cn",
		"santest.chinaeast2.cloudapp.chinacloudapi.cn",
		"santest.germanycentral.cloudapp.microsoftazure.de",
		"santest.germanynortheast.cloudapp.microsoftazure.de",
		"santest.usgovvirginia.cloudapp.usgovcloudapi.net",
		"santest.usgoviowa.cloudapp.usgovcloudapi.net",
		"santest.usgovarizona.cloudapp.usgovcloudapi.net",
		"santest.usgovtexas.cloudapp.usgovcloudapi.net",
		"santest.francecentral.cloudapp.azure.com",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected formatted fqdns %s, but got %s", expected, actual)
	}

}

func TestKubernetesConfig_GetAddonScript(t *testing.T) {
	addon := getMockAddon(IPMASQAgentAddonName)
	addon.Data = "foobarbazdata"
	k := &KubernetesConfig{
		Addons: []KubernetesAddon{
			addon,
		},
	}

	expected := "foobarbazdata"
	actual := k.GetAddonScript(IPMASQAgentAddonName)
	if actual != expected {
		t.Errorf("expected GetAddonScript to return %s, but got %s", expected, actual)
	}
}

func TestContainerService_GetAzureProdFQDN(t *testing.T) {
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 1, 3, false)
	expected := "testmaster.eastus.cloudapp.azure.com"
	actual := cs.GetAzureProdFQDN()

	if expected != actual {
		t.Errorf("expected GetAzureProdFQDN to return %s, but got %s", expected, actual)
	}
}

func TestKubernetesConfig_RequiresDocker(t *testing.T) {
	// k8sConfig with empty runtime string
	k := &KubernetesConfig{
		ContainerRuntime: "",
	}

	if !k.RequiresDocker() {
		t.Error("expected RequiresDocker to return true for empty runtime string")
	}

	// k8sConfig with empty runtime string
	k = &KubernetesConfig{
		ContainerRuntime: "docker",
	}

	if !k.RequiresDocker() {
		t.Error("expected RequiresDocker to return true for docker runtime")
	}
}

func TestProperties_GetMasterVMPrefix(t *testing.T) {
	p := &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		MasterProfile: &MasterProfile{
			Count:     1,
			DNSPrefix: "myprefix1",
			VMSize:    "Standard_DS2_v2",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: "VirtualMachineScaleSets",
				OSType:              "Linux",
			},
		},
	}

	actual := p.GetMasterVMPrefix()
	expected := "k8s-master-30819786-"

	if actual != expected {
		t.Errorf("expected master VM prefix %s, but got %s", expected, actual)
	}
}

func TestIsFeatureEnabled(t *testing.T) {
	tests := []struct {
		name     string
		feature  string
		flags    *FeatureFlags
		expected bool
	}{
		{
			name:     "nil flags",
			feature:  "BlockOutboundInternet",
			flags:    nil,
			expected: false,
		},
		{
			name:     "empty flags",
			feature:  "BlockOutboundInternet",
			flags:    &FeatureFlags{},
			expected: false,
		},
		{
			name:    "Enabled feature",
			feature: "CSERunInBackground",
			flags: &FeatureFlags{
				EnableCSERunInBackground: true,
				BlockOutboundInternet:    false,
			},
			expected: true,
		},
		{
			name:    "Disabled feature",
			feature: "CSERunInBackground",
			flags: &FeatureFlags{
				EnableCSERunInBackground: false,
				BlockOutboundInternet:    true,
			},
			expected: false,
		},
		{
			name:    "Non-existent feature",
			feature: "Foo",
			flags: &FeatureFlags{
				EnableCSERunInBackground: true,
				BlockOutboundInternet:    true,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.flags.IsFeatureEnabled(test.feature)
			if actual != test.expected {
				t.Errorf("expected feature %s to be enabled:%v, but got %v", test.feature, test.expected, actual)
			}
		})
	}
}
