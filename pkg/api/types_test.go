// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
)

const exampleCustomHyperkubeImage = `example.azurecr.io/example/hyperkube-amd64:custom`
const examplePrivateAzureRegistryServer = `example.azurecr.io`

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

const exampleAKSAPIModel = `{
		"apiVersion": "2018-03-31",
	"properties": {
		"dnsPrefix": "agents006",
		"fqdn": "agents006.azmk8s.io",
		"kubernetesVersion": "1.15.11",
		"agentPoolProfiles": [ { "name": "agentpool1", "count": 2, "vmSize": "Standard_D2_v2" } ],
		"linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	},
	"servicePrincipalProfile": { "clientId": "", "secret": "" }
	}
}
`

const exampleAPIModelWithPrivateAzureRegistry = `{
	"apiVersion": "vlabs",
"properties": {
	"orchestratorProfile": {
		"orchestratorType": "Kubernetes",
		"kubernetesConfig": {
			"customHyperkubeImage": "` + exampleCustomHyperkubeImage + `",
			"privateAzureRegistryServer": "` + examplePrivateAzureRegistryServer + `"
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

func TestOrchestratorProfile_GetPodInfraContainerSpec(t *testing.T) {
	o := OrchestratorProfile{
		KubernetesConfig: &KubernetesConfig{
			MCRKubernetesImageBase: "foo/",
		},
		OrchestratorVersion: "1.16.0",
	}
	expected := "foo/k8s/core/pause:1.2.0"
	actual := o.GetPodInfraContainerSpec()
	if actual != expected {
		t.Fatalf("expected GetPodInfraContainerSpec to return %s, but got %s", expected, actual)
	}
}

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

	if !p.AgentPoolProfiles[1].IsRHEL() {
		t.Fatalf("expected IsRHEL() to return true but instead returned false")
	}

	if !p.MasterProfile.IsRHEL() {
		t.Fatalf("expected IsRHEL() to return true but instead returned false")
	}
}

func TestAgentPoolProfileIsVHDDistro(t *testing.T) {
	cases := []struct {
		name     string
		ap       AgentPoolProfile
		expected bool
	}{
		{
			name: "16.04 VHD distro",
			ap: AgentPoolProfile{
				Distro: AKSUbuntu1604,
			},
			expected: true,
		},
		{
			name: "18.04 VHD distro",
			ap: AgentPoolProfile{
				Distro: AKSUbuntu1804,
			},
			expected: true,
		},
		{
			name: "ubuntu distro",
			ap: AgentPoolProfile{
				Distro: Ubuntu,
			},
			expected: false,
		},
		{
			name: "ubuntu 18.04 non-VHD distro",
			ap: AgentPoolProfile{
				Distro: Ubuntu1804,
			},
			expected: false,
		},
		{
			name: "ubuntu 18.04 gen2 non-VHD distro",
			ap: AgentPoolProfile{
				Distro: Ubuntu1804Gen2,
			},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.ap.IsVHDDistro() {
				t.Fatalf("Got unexpected AgentPoolProfile.IsVHDDistro() result. Expected: %t. Got: %t.", c.expected, c.ap.IsVHDDistro())
			}
		})
	}
}

func TestAgentPoolProfileIsAuditDEnabled(t *testing.T) {
	cases := []struct {
		name     string
		ap       AgentPoolProfile
		expected bool
	}{
		{
			name:     "default",
			ap:       AgentPoolProfile{},
			expected: false,
		},
		{
			name: "true",
			ap: AgentPoolProfile{
				AuditDEnabled: to.BoolPtr(true),
			},
			expected: true,
		},
		{
			name: "false",
			ap: AgentPoolProfile{
				AuditDEnabled: to.BoolPtr(false),
			},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.ap.IsAuditDEnabled() {
				t.Fatalf("Got unexpected AgentPoolProfile.IsAuditDEnabled() result. Expected: %t. Got: %t.", c.expected, c.ap.IsAuditDEnabled())
			}
		})
	}
}

func TestMasterProfileIsAuditDEnabled(t *testing.T) {
	cases := []struct {
		name     string
		mp       MasterProfile
		expected bool
	}{
		{
			name:     "default",
			mp:       MasterProfile{},
			expected: false,
		},
		{
			name: "true",
			mp: MasterProfile{
				AuditDEnabled: to.BoolPtr(true),
			},
			expected: true,
		},
		{
			name: "false",
			mp: MasterProfile{
				AuditDEnabled: to.BoolPtr(false),
			},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.mp.IsAuditDEnabled() {
				t.Fatalf("Got unexpected AgentPoolProfile.IsAuditDEnabled() result. Expected: %t. Got: %t.", c.expected, c.mp.IsAuditDEnabled())
			}
		})
	}
}

func TestAgentPoolProfileIsUbuntuNonVHD(t *testing.T) {
	cases := []struct {
		name     string
		ap       AgentPoolProfile
		expected bool
	}{
		{
			name: "ubuntu 16.04 VHD distro",
			ap: AgentPoolProfile{
				Distro: AKSUbuntu1604,
			},
			expected: false,
		},
		{
			name: "ubuntu 18.04 VHD distro",
			ap: AgentPoolProfile{
				Distro: AKSUbuntu1804,
			},
			expected: false,
		},
		{
			name: "ubuntu distro",
			ap: AgentPoolProfile{
				Distro: Ubuntu,
			},
			expected: true,
		},
		{
			name: "ubuntu 18.04 non-VHD distro",
			ap: AgentPoolProfile{
				Distro: Ubuntu1804,
			},
			expected: true,
		},
		{
			name: "ubuntu 18.04 gen2 non-VHD distro",
			ap: AgentPoolProfile{
				Distro: Ubuntu1804Gen2,
			},
			expected: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.ap.IsUbuntuNonVHD() {
				t.Fatalf("Got unexpected AgentPoolProfile.IsUbuntuNonVHD() result. Expected: %t. Got: %t.", c.expected, c.ap.IsUbuntuNonVHD())
			}
		})
	}
}

func TestRequiresCloudproviderConfig(t *testing.T) {
	cases := []struct {
		name     string
		ap       AgentPoolProfile
		expected bool
	}{
		{
			name:     "nil",
			ap:       AgentPoolProfile{},
			expected: true,
		},
		{
			name: "default",
			ap: AgentPoolProfile{
				KubernetesConfig: &KubernetesConfig{
					KubeletConfig: map[string]string{},
				},
			},
			expected: true,
		},
		{
			name: "--cloud-provider provided",
			ap: AgentPoolProfile{
				KubernetesConfig: &KubernetesConfig{
					KubeletConfig: map[string]string{
						"--cloud-provider":                  "azure",
						"--cloud-config":                    "",
						"--azure-container-registry-config": "",
					},
				},
			},
			expected: true,
		},
		{
			name: "--cloud-config provided",
			ap: AgentPoolProfile{
				KubernetesConfig: &KubernetesConfig{
					KubeletConfig: map[string]string{
						"--cloud-provider":                  "",
						"--cloud-config":                    "/etc/kubernetes/azure.json",
						"--azure-container-registry-config": "",
					},
				},
			},
			expected: true,
		},
		{
			name: "--azure-container-registry-config provided",
			ap: AgentPoolProfile{
				KubernetesConfig: &KubernetesConfig{
					KubeletConfig: map[string]string{
						"--cloud-provider":                  "",
						"--cloud-config":                    "",
						"--azure-container-registry-config": "/etc/kubernetes/azure.json",
					},
				},
			},
			expected: true,
		},
		{
			name: "all 3 flags set explicitly to empty string",
			ap: AgentPoolProfile{
				KubernetesConfig: &KubernetesConfig{
					KubeletConfig: map[string]string{
						"--cloud-provider":                  "",
						"--cloud-config":                    "",
						"--azure-container-registry-config": "",
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.ap.RequiresCloudproviderConfig() {
				t.Fatalf("Got unexpected AgentPoolProfile.RequiresCloudproviderConfig() result. Expected: %t. Got: %t.", c.expected, c.ap.RequiresCloudproviderConfig())
			}
		})
	}
}

func TestMasterProfileIsVHDDistro(t *testing.T) {
	cases := []struct {
		name     string
		m        MasterProfile
		expected bool
	}{
		{
			name: "ubuntu 16.04 VHD distro",
			m: MasterProfile{
				Distro: AKSUbuntu1604,
			},
			expected: true,
		},
		{
			name: "ubuntu 18.04 VHD distro",
			m: MasterProfile{
				Distro: AKSUbuntu1804,
			},
			expected: true,
		},
		{
			name: "ubuntu 16.04 non-VHD distro",
			m: MasterProfile{
				Distro: Ubuntu,
			},
			expected: false,
		},
		{
			name: "ubuntu 18.04 non-VHD distro",
			m: MasterProfile{
				Distro: Ubuntu1804,
			},
			expected: false,
		},
		{
			name: "ubuntu 18.04 gen2 non-VHD distro",
			m: MasterProfile{
				Distro: Ubuntu1804Gen2,
			},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.m.IsVHDDistro() {
				t.Fatalf("Got unexpected MasterProfile.IsVHDDistro() result. Expected: %t. Got: %t.", c.expected, c.m.IsVHDDistro())
			}
		})
	}
}

func TestMasterProfileIsUbuntuNonVHD(t *testing.T) {
	cases := []struct {
		name     string
		m        MasterProfile
		expected bool
	}{
		{
			name: "ubuntu 16.04 VHD distro",
			m: MasterProfile{
				Distro: AKSUbuntu1604,
			},
			expected: false,
		},
		{
			name: "ubuntu 18.04 VHD distro",
			m: MasterProfile{
				Distro: AKSUbuntu1804,
			},
			expected: false,
		},
		{
			name: "ubuntu 16.04 non-VHD distro",
			m: MasterProfile{
				Distro: Ubuntu,
			},
			expected: true,
		},
		{
			name: "ubuntu 18.04 non-VHD distro",
			m: MasterProfile{
				Distro: Ubuntu1804,
			},
			expected: true,
		},
		{
			name: "ubuntu 18.04 gen2 non-VHD distro",
			m: MasterProfile{
				Distro: Ubuntu1804Gen2,
			},
			expected: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.m.IsUbuntuNonVHD() {
				t.Fatalf("Got unexpected MasterProfile.IsUbuntuNonVHD() result. Expected: %t. Got: %t.", c.expected, c.m.IsUbuntuNonVHD())
			}
		})
	}
}

func TestMasterProfileHasMultipleNodes(t *testing.T) {
	cases := []struct {
		name     string
		m        MasterProfile
		expected bool
	}{
		{
			name: "1",
			m: MasterProfile{
				Count: 1,
			},
			expected: false,
		},
		{
			name: "2",
			m: MasterProfile{
				Count: 2,
			},
			expected: true,
		},
		{
			name: "3",
			m: MasterProfile{
				Count: 3,
			},
			expected: true,
		},
		{
			name: "0",
			m: MasterProfile{
				Count: 0,
			},
			expected: false,
		},
		{
			name: "-1",
			m: MasterProfile{
				Count: -1,
			},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.m.HasMultipleNodes() {
				t.Fatalf("Got unexpected MasterProfile.HasMultipleNodes() result. Expected: %t. Got: %t.", c.expected, c.m.HasMultipleNodes())
			}
		})
	}
}

func TestMasterProfileHasCosmosEtcd(t *testing.T) {
	cases := []struct {
		name     string
		m        MasterProfile
		expected bool
	}{
		{
			name: "enabled",
			m: MasterProfile{
				CosmosEtcd: to.BoolPtr(true),
			},
			expected: true,
		},
		{
			name: "disabled",
			m: MasterProfile{
				CosmosEtcd: to.BoolPtr(false),
			},
			expected: false,
		},
		{
			name:     "zero value master profile",
			m:        MasterProfile{},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.m.HasCosmosEtcd() {
				t.Fatalf("Got unexpected MasterProfile.HasCosmosEtcd() result. Expected: %t. Got: %t.", c.expected, c.m.HasCosmosEtcd())
			}
		})
	}
}

func TestMasterProfileGetCosmosEndPointURI(t *testing.T) {
	dnsPrefix := "my-prefix"
	cases := []struct {
		name     string
		m        MasterProfile
		expected string
	}{
		{
			name: "valid DNS prefix",
			m: MasterProfile{
				CosmosEtcd: to.BoolPtr(true),
				DNSPrefix:  dnsPrefix,
			},
			expected: fmt.Sprintf(etcdEndpointURIFmt, dnsPrefix),
		},
		{
			name: "no DNS prefix",
			m: MasterProfile{
				CosmosEtcd: to.BoolPtr(true),
			},
			expected: fmt.Sprintf(etcdEndpointURIFmt, ""),
		},
		{
			name: "cosmos etcd disabled",
			m: MasterProfile{
				CosmosEtcd: to.BoolPtr(false),
			},
			expected: "",
		},
		{
			name:     "zero value master profile",
			m:        MasterProfile{},
			expected: "",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.m.GetCosmosEndPointURI() {
				t.Fatalf("Got unexpected MasterProfile.GetCosmosEndPointURI() result. Expected: %s. Got: %s.", c.expected, c.m.GetCosmosEndPointURI())
			}
		})
	}
}

func TestHasStorageProfile(t *testing.T) {
	cases := []struct {
		name                     string
		p                        Properties
		expectedHasMD            bool
		expectedHasSA            bool
		expectedMasterMD         bool
		expectedAgent0E          bool
		expectedAgent0MD         bool
		expectedPrivateJB        bool
		expectedHasDisks         bool
		expectedDesID            string
		expectedUltraSSDEnabled  bool
		expectedEncryptionAtHost bool
	}{
		{
			name: "Storage Account",
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
			expectedAgent0E:  false,
			expectedHasDisks: true,
		},
		{
			name: "Managed Disk",
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
			expectedAgent0E:  false,
		},
		{
			name: "both",
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
			expectedAgent0E:  false,
		},
		{
			name: "Managed Disk everywhere",
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
			expectedAgent0E:   false,
			expectedPrivateJB: false,
		},
		{
			name: "Managed disk master with ephemeral agent",
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					StorageProfile: ManagedDisks,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile: Ephemeral,
					},
				},
			},
			expectedHasMD:     true,
			expectedHasSA:     false,
			expectedMasterMD:  true,
			expectedAgent0MD:  false,
			expectedAgent0E:   true,
			expectedPrivateJB: false,
		},
		{
			name: "Mixed with jumpbox",
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						PrivateCluster: &PrivateCluster{
							Enabled: to.BoolPtr(true),
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
			expectedAgent0E:   false,
			expectedPrivateJB: true,
		},
		{
			name: "Mixed with jumpbox alternate",
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						PrivateCluster: &PrivateCluster{
							Enabled: to.BoolPtr(true),
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
			expectedAgent0E:   false,
			expectedPrivateJB: true,
		},
		{
			name: "Managed Disk with DiskEncryptionSetID setting",
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					StorageProfile: ManagedDisks,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile:      ManagedDisks,
						DiskEncryptionSetID: "DiskEncryptionSetID",
					},
					{
						StorageProfile:      ManagedDisks,
						DiskEncryptionSetID: "DiskEncryptionSetID",
					},
				},
			},
			expectedHasMD:     true,
			expectedHasSA:     false,
			expectedMasterMD:  true,
			expectedAgent0MD:  true,
			expectedAgent0E:   false,
			expectedPrivateJB: false,
			expectedDesID:     "DiskEncryptionSetID",
		},
		{
			name: "UltraSSDEnabled setting",
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					StorageProfile:  ManagedDisks,
					UltraSSDEnabled: to.BoolPtr(true),
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile:  ManagedDisks,
						UltraSSDEnabled: to.BoolPtr(true),
					},
					{
						StorageProfile:  ManagedDisks,
						UltraSSDEnabled: to.BoolPtr(true),
					},
				},
			},
			expectedHasMD:           true,
			expectedHasSA:           false,
			expectedMasterMD:        true,
			expectedAgent0MD:        true,
			expectedAgent0E:         false,
			expectedPrivateJB:       false,
			expectedUltraSSDEnabled: true,
		},
		{
			name: "EncryptionAtHost setting",
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					StorageProfile:   ManagedDisks,
					EncryptionAtHost: to.BoolPtr(true),
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						StorageProfile:   ManagedDisks,
						EncryptionAtHost: to.BoolPtr(true),
					},
					{
						StorageProfile:   ManagedDisks,
						EncryptionAtHost: to.BoolPtr(true),
					},
				},
			},
			expectedHasMD:            true,
			expectedHasSA:            false,
			expectedMasterMD:         true,
			expectedAgent0MD:         true,
			expectedAgent0E:          false,
			expectedPrivateJB:        false,
			expectedEncryptionAtHost: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
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
			if to.Bool(c.p.MasterProfile.UltraSSDEnabled) != c.expectedUltraSSDEnabled {
				t.Fatalf("expected UltraSSDEnabled to return %v but instead returned %v", c.expectedUltraSSDEnabled, to.Bool(c.p.MasterProfile.UltraSSDEnabled))
			}
			if to.Bool(c.p.MasterProfile.EncryptionAtHost) != c.expectedEncryptionAtHost {
				t.Fatalf("expected EncryptionAtHost to return %v but instead returned %v", c.expectedEncryptionAtHost, to.Bool(c.p.MasterProfile.EncryptionAtHost))
			}
			if c.p.AgentPoolProfiles[0].IsManagedDisks() != c.expectedAgent0MD {
				t.Fatalf("expected IsManagedDisks() to return %t but instead returned %t", c.expectedAgent0MD, c.p.AgentPoolProfiles[0].IsManagedDisks())
			}
			expectedAgentStorageAccount := !(c.expectedAgent0MD || c.expectedAgent0E)
			if c.p.AgentPoolProfiles[0].IsStorageAccount() != expectedAgentStorageAccount {
				t.Fatalf("expected IsStorageAccount() to return %t but instead returned %t", expectedAgentStorageAccount, c.p.AgentPoolProfiles[0].IsStorageAccount())
			}
			if c.p.AgentPoolProfiles[0].IsEphemeral() != c.expectedAgent0E {
				t.Fatalf("expected IsEphemeral() to return %t but instead returned %t", c.expectedAgent0E, c.p.AgentPoolProfiles[0].IsEphemeral())
			}
			if c.p.OrchestratorProfile != nil && c.p.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() != c.expectedPrivateJB {
				t.Fatalf("expected PrivateJumpboxProvision() to return %t but instead returned %t", c.expectedPrivateJB, c.p.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision())
			}
			if c.p.AgentPoolProfiles[0].HasDisks() != c.expectedHasDisks {
				t.Fatalf("expected HasDisks() to return %t but instead returned %t", c.expectedHasDisks, c.p.AgentPoolProfiles[0].HasDisks())
			}
			if c.p.AgentPoolProfiles[0].DiskEncryptionSetID != c.expectedDesID {
				t.Fatalf("expected DiskEncryptionSetID to return %s but instead returned %s", c.expectedDesID, c.p.AgentPoolProfiles[0].DiskEncryptionSetID)
			}
			if to.Bool(c.p.AgentPoolProfiles[0].UltraSSDEnabled) != c.expectedUltraSSDEnabled {
				t.Fatalf("expected UltraSSDEnabled to return %v but instead returned %v", c.expectedUltraSSDEnabled, to.Bool(c.p.AgentPoolProfiles[0].UltraSSDEnabled))
			}
			if to.Bool(c.p.AgentPoolProfiles[0].EncryptionAtHost) != c.expectedEncryptionAtHost {
				t.Fatalf("expected EncryptionAtHost to return %v but instead returned %v", c.expectedUltraSSDEnabled, to.Bool(c.p.AgentPoolProfiles[0].UltraSSDEnabled))
			}
		})
	}
}

func TestAgentPoolProfileGetKubernetesLabels(t *testing.T) {
	cases := []struct {
		name       string
		ap         AgentPoolProfile
		rg         string
		deprecated bool
		expected   string
	}{
		{
			name:       "vanilla pool profile",
			ap:         AgentPoolProfile{},
			rg:         "my-resource-group",
			deprecated: true,
			expected:   "kubernetes.azure.com/role=agent,node-role.kubernetes.io/agent=,kubernetes.io/role=agent,agentpool=,kubernetes.azure.com/cluster=my-resource-group",
		},
		{
			name:       "vanilla pool profile, no deprecated labels",
			ap:         AgentPoolProfile{},
			rg:         "my-resource-group",
			deprecated: false,
			expected:   "kubernetes.azure.com/role=agent,agentpool=,kubernetes.azure.com/cluster=my-resource-group",
		},
		{
			name: "with managed disk",
			ap: AgentPoolProfile{
				StorageProfile: ManagedDisks,
			},
			rg:         "my-resource-group",
			deprecated: true,
			expected:   "kubernetes.azure.com/role=agent,node-role.kubernetes.io/agent=,kubernetes.io/role=agent,agentpool=,storageprofile=managed,storagetier=,kubernetes.azure.com/cluster=my-resource-group",
		},
		{
			name: "N series",
			ap: AgentPoolProfile{
				VMSize: "Standard_NC6",
			},
			rg:         "my-resource-group",
			deprecated: true,
			expected:   "kubernetes.azure.com/role=agent,node-role.kubernetes.io/agent=,kubernetes.io/role=agent,agentpool=,accelerator=nvidia,kubernetes.azure.com/cluster=my-resource-group",
		},
		{
			name: "with custom labels",
			ap: AgentPoolProfile{
				CustomNodeLabels: map[string]string{
					"mycustomlabel1": "foo",
					"mycustomlabel2": "bar",
				},
			},
			rg:         "my-resource-group",
			deprecated: true,
			expected:   "kubernetes.azure.com/role=agent,node-role.kubernetes.io/agent=,kubernetes.io/role=agent,agentpool=,kubernetes.azure.com/cluster=my-resource-group,mycustomlabel1=foo,mycustomlabel2=bar",
		},
		{
			name: "with custom labels, no deprecated labels",
			ap: AgentPoolProfile{
				CustomNodeLabels: map[string]string{
					"mycustomlabel1": "foo",
					"mycustomlabel2": "bar",
				},
			},
			rg:         "my-resource-group",
			deprecated: false,
			expected:   "kubernetes.azure.com/role=agent,agentpool=,kubernetes.azure.com/cluster=my-resource-group,mycustomlabel1=foo,mycustomlabel2=bar",
		},
		{
			name: "N series and managed disk with custom labels",
			ap: AgentPoolProfile{
				StorageProfile: ManagedDisks,
				VMSize:         "Standard_NC6",
				CustomNodeLabels: map[string]string{
					"mycustomlabel1": "foo",
					"mycustomlabel2": "bar",
				},
			},
			rg:         "my-resource-group",
			deprecated: true,
			expected:   "kubernetes.azure.com/role=agent,node-role.kubernetes.io/agent=,kubernetes.io/role=agent,agentpool=,storageprofile=managed,storagetier=Standard_LRS,accelerator=nvidia,kubernetes.azure.com/cluster=my-resource-group,mycustomlabel1=foo,mycustomlabel2=bar",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.ap.GetKubernetesLabels(c.rg, c.deprecated) {
				t.Fatalf("Got unexpected AgentPoolProfile.GetKubernetesLabels(%s, %t) result. Expected: %s. Got: %s.",
					c.rg, c.deprecated, c.expected, c.ap.GetKubernetesLabels(c.rg, c.deprecated))
			}
		})
	}
}

func TestKubernetesConfigGetOrderedKubeletConfigString(t *testing.T) {
	alphabetizedString := "--address=0.0.0.0 --allow-privileged=true --anonymous-auth=false --authorization-mode=Webhook --cgroups-per-qos=true --client-ca-file=/etc/kubernetes/certs/ca.crt --keep-terminated-pod-volumes=false --kubeconfig=/var/lib/kubelet/kubeconfig --pod-manifest-path=/etc/kubernetes/manifests "
	alphabetizedStringForPowershell := `"--address=0.0.0.0", "--allow-privileged=true", "--anonymous-auth=false", "--authorization-mode=Webhook", "--cgroups-per-qos=true", "--client-ca-file=/etc/kubernetes/certs/ca.crt", "--keep-terminated-pod-volumes=false", "--kubeconfig=/var/lib/kubelet/kubeconfig", "--pod-manifest-path=/etc/kubernetes/manifests"`
	cases := []struct {
		name                  string
		kc                    KubernetesConfig
		expected              string
		expectedForPowershell string
	}{
		{
			name:                  "zero value kubernetesConfig",
			kc:                    KubernetesConfig{},
			expected:              "",
			expectedForPowershell: "",
		},
		// Some values
		{
			name: "expected values",
			kc: KubernetesConfig{
				KubeletConfig: map[string]string{
					"--address":                     "0.0.0.0",
					"--allow-privileged":            "true",
					"--anonymous-auth":              "false",
					"--authorization-mode":          "Webhook",
					"--client-ca-file":              "/etc/kubernetes/certs/ca.crt",
					"--pod-manifest-path":           "/etc/kubernetes/manifests",
					"--cgroups-per-qos":             "true",
					"--kubeconfig":                  "/var/lib/kubelet/kubeconfig",
					"--keep-terminated-pod-volumes": "false",
				},
			},
			expected:              alphabetizedString,
			expectedForPowershell: alphabetizedStringForPowershell,
		},
		// Switch the "order" in the map, validate the same return string
		{
			name: "expected values re-ordered",
			kc: KubernetesConfig{
				KubeletConfig: map[string]string{
					"--address":                     "0.0.0.0",
					"--allow-privileged":            "true",
					"--kubeconfig":                  "/var/lib/kubelet/kubeconfig",
					"--client-ca-file":              "/etc/kubernetes/certs/ca.crt",
					"--authorization-mode":          "Webhook",
					"--pod-manifest-path":           "/etc/kubernetes/manifests",
					"--cgroups-per-qos":             "true",
					"--keep-terminated-pod-volumes": "false",
					"--anonymous-auth":              "false",
				},
			},
			expected:              alphabetizedString,
			expectedForPowershell: alphabetizedStringForPowershell,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expected != c.kc.GetOrderedKubeletConfigString() {
				t.Fatalf("Got unexpected AgentPoolProfile.GetOrderedKubeletConfigString() result. Expected: %s. Got: %s.", c.expected, c.kc.GetOrderedKubeletConfigString())
			}
			if c.expectedForPowershell != c.kc.GetOrderedKubeletConfigStringForPowershell() {
				t.Fatalf("Got unexpected AgentPoolProfile.GetOrderedKubeletConfigStringForPowershell() result. Expected: %s. Got: %s.", c.expectedForPowershell, c.kc.GetOrderedKubeletConfigStringForPowershell())
			}
		})
	}
}

func TestTotalNodes(t *testing.T) {
	cases := []struct {
		name     string
		p        Properties
		expected int
	}{
		{
			name: "2 total nodes between master and pool",
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
			name: "7 total nodes between 2 pools",
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
			name: "11 total nodes between master and pool",
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
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.p.TotalNodes() != c.expected {
				t.Fatalf("expected TotalNodes() to return %d but instead returned %d", c.expected, c.p.TotalNodes())
			}
		})
	}
}

func TestPropertiesIsHostedMasterProfile(t *testing.T) {
	cases := []struct {
		name     string
		p        Properties
		expected bool
	}{
		{
			name: "valid master 1 node",
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
			},
			expected: false,
		},
		{
			name: "valid master 3 nodes",
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 3,
				},
			},
			expected: false,
		},
		{
			name: "valid master 5 nodes",
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 5,
				},
			},
			expected: false,
		},
		{
			name: "zero value hosted master",
			p: Properties{
				HostedMasterProfile: &HostedMasterProfile{},
			},
			expected: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.p.IsHostedMasterProfile() != c.expected {
				t.Fatalf("expected IsHostedMasterProfile() to return %t but instead returned %t", c.expected, c.p.IsHostedMasterProfile())
			}
		})
	}
}

func TestPropertiesMasterCustomOS(t *testing.T) {
	cases := []struct {
		name            string
		p               Properties
		expectedRef     bool
		expectedGallery bool
	}{
		{
			name: "valid shared gallery image",
			p: Properties{
				MasterProfile: &MasterProfile{
					ImageRef: &ImageReference{
						Name:           "testImage",
						ResourceGroup:  "testRg",
						SubscriptionID: "testSub",
						Gallery:        "testGallery",
						Version:        "0.0.1",
					},
				},
			},
			expectedRef:     true,
			expectedGallery: true,
		},
		{
			name: "valid resource group image",
			p: Properties{
				MasterProfile: &MasterProfile{
					ImageRef: &ImageReference{
						Name:          "testImage",
						ResourceGroup: "testRg",
					},
				},
			},
			expectedRef:     true,
			expectedGallery: false,
		},
		{
			name: "valid no custom image",
			p: Properties{
				MasterProfile: &MasterProfile{
					ImageRef: nil,
				},
			},
			expectedRef:     false,
			expectedGallery: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.p.MasterProfile.HasImageRef() != c.expectedRef || c.p.MasterProfile.HasImageGallery() != c.expectedGallery {
				t.Fatalf("expected HasImageRef() to return %t but instead returned %t, HasImageGallery() expected: %t but actual: %t", c.expectedRef, c.p.MasterProfile.HasImageRef(), c.p.MasterProfile.HasImageGallery(), c.expectedGallery)
			}
		})
	}
}

func TestPropertiesAgentCustomOS(t *testing.T) {
	cases := []struct {
		name            string
		p               Properties
		expectedRef     bool
		expectedGallery bool
	}{
		{
			name: "valid shared gallery image",
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
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
			expectedRef:     true,
			expectedGallery: true,
		},
		{
			name: "valid resource group image",
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						ImageRef: &ImageReference{
							Name:          "testImage",
							ResourceGroup: "testRg",
						},
					},
				},
			},
			expectedRef:     true,
			expectedGallery: false,
		},
		{
			name: "valid no custom image",
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						ImageRef: nil,
					},
				},
			},
			expectedRef:     false,
			expectedGallery: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.p.AgentPoolProfiles[0].HasImageRef() != c.expectedRef || c.p.AgentPoolProfiles[0].HasImageGallery() != c.expectedGallery {
				t.Fatalf("expected HasImageRef() to return %t but instead returned %t, HasImageGallery() expected: %t but actual: %t", c.expectedRef, c.p.AgentPoolProfiles[0].HasImageRef(), c.expectedGallery, c.p.AgentPoolProfiles[0].HasImageGallery())
			}
		})
	}
}

func TestMasterAvailabilityProfile(t *testing.T) {
	cases := []struct {
		name           string
		p              Properties
		expectedISVMSS bool
		expectedIsVMAS bool
	}{
		{
			name: "zero value master profile",
			p: Properties{
				MasterProfile: &MasterProfile{},
			},
			expectedISVMSS: false,
			expectedIsVMAS: false,
		},
		{
			name: "master profile w/ AS",
			p: Properties{
				MasterProfile: &MasterProfile{
					AvailabilityProfile: AvailabilitySet,
				},
			},
			expectedISVMSS: false,
			expectedIsVMAS: true,
		},
		{
			name: "master profile w/ VMSS",
			p: Properties{
				MasterProfile: &MasterProfile{
					AvailabilityProfile: VirtualMachineScaleSets,
				},
			},
			expectedISVMSS: true,
			expectedIsVMAS: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.p.MasterProfile.IsVirtualMachineScaleSets() != c.expectedISVMSS {
				t.Fatalf("expected MasterProfile.IsVirtualMachineScaleSets() to return %t but instead returned %t", c.expectedISVMSS, c.p.MasterProfile.IsVirtualMachineScaleSets())
			}
			if c.p.MasterProfile.IsAvailabilitySet() != c.expectedIsVMAS {
				t.Fatalf("expected MasterProfile.IsAvailabilitySet() to return %t but instead returned %t", c.expectedIsVMAS, c.p.MasterProfile.IsAvailabilitySet())
			}
		})
	}
}
func TestAvailabilityProfile(t *testing.T) {
	cases := []struct {
		p               Properties
		expectedHasVMSS bool
		expectedISVMSS  bool
		expectedIsAS    bool
		expectedLowPri  bool
		expectedSpot    bool
		expectedVMType  string
	}{
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
						ScaleSetPriority:    ScaleSetPrioritySpot,
					},
				},
			},
			expectedHasVMSS: true,
			expectedISVMSS:  true,
			expectedIsAS:    false,
			expectedLowPri:  false,
			expectedSpot:    true,
			expectedVMType:  VMSSVMType,
		},
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
			expectedSpot:    false,
			expectedVMType:  VMSSVMType,
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
			expectedSpot:    false,
			expectedVMType:  VMSSVMType,
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
			expectedSpot:    false,
			expectedVMType:  StandardVMType,
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
		if c.p.AgentPoolProfiles[0].IsSpotScaleSet() != c.expectedSpot {
			t.Fatalf("expected IsSpotScaleSet() to return %t but instead returned %t", c.expectedSpot, c.p.AgentPoolProfiles[0].IsSpotScaleSet())
		}
		if c.p.GetVMType() != c.expectedVMType {
			t.Fatalf("expected GetVMType() to return %s but instead returned %s", c.expectedVMType, c.p.GetVMType())
		}
	}
}

func TestPerAgentPoolVersionAndState(t *testing.T) {
	cases := []struct {
		ap              AgentPoolProfile
		expectedVersion string
		expectedState   ProvisioningState
	}{
		{
			ap: AgentPoolProfile{
				Name:                "agentpool1",
				OrchestratorVersion: "1.12.0",
				ProvisioningState:   Creating,
			},
			expectedVersion: "1.12.0",
			expectedState:   Creating,
		},
		{
			ap: AgentPoolProfile{
				Name:                "agentpool2",
				OrchestratorVersion: "",
				ProvisioningState:   "",
			},
			expectedVersion: "",
			expectedState:   "",
		},
	}

	for _, c := range cases {
		if c.ap.OrchestratorVersion != c.expectedVersion {
			t.Fatalf("Orchestrator profile mismatch. Expected: %s. Got: %s.", c.expectedVersion, c.ap.OrchestratorVersion)
		}
		if c.ap.ProvisioningState != c.expectedState {
			t.Fatalf("Provisioning state mismatch. Expected: %s. Got: %s.", c.expectedState, c.ap.ProvisioningState)
		}
	}
}

func TestPerAgentPoolWindowsNameVersion(t *testing.T) {
	cases := []struct {
		ap                         AgentPoolProfile
		expectedWindowsNameVersion string
	}{
		{
			ap: AgentPoolProfile{
				Name:               "agentpool1",
				WindowsNameVersion: "v2",
			},
			expectedWindowsNameVersion: "v2",
		},
		{
			ap: AgentPoolProfile{
				Name: "agentpool2",
			},
			expectedWindowsNameVersion: "",
		},
	}

	for _, c := range cases {
		if c.expectedWindowsNameVersion != c.ap.WindowsNameVersion {
			t.Fatalf("WindowsNameVersion flag mismatch. Expected: %v. Got: %v.", &c.expectedWindowsNameVersion, &c.ap.WindowsNameVersion)
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

func TestHasNonRegularPriorityScaleset(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
						ScaleSetPriority:    ScaleSetPriorityLow,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
						ScaleSetPriority:    ScaleSetPriorityLow,
					},
					{
						AvailabilityProfile: VirtualMachineScaleSets,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
						ScaleSetPriority:    ScaleSetPriorityLow,
					},
					{
						AvailabilityProfile: AvailabilitySet,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
						ScaleSetPriority:    ScaleSetPrioritySpot,
					},
					{
						AvailabilityProfile: AvailabilitySet,
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
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
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: VirtualMachineScaleSets,
					},
					{
						AvailabilityProfile: AvailabilitySet,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						AvailabilityProfile: AvailabilitySet,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count: 1,
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.HasNonRegularPriorityScaleset() != c.expected {
			t.Fatalf("expected HasNonRegularPriorityScaleset() to return %t but instead returned %t", c.expected, c.p.HasNonRegularPriorityScaleset())
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
					Distro: Ubuntu1804Gen2,
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
					Distro: Ubuntu1804Gen2,
				},
			},
			expected: true,
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
						Distro: Ubuntu1804Gen2,
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

func TestIsUbuntuDistroForAllNodes(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: AKSUbuntu1604,
					},
				},
			},
			expected: false,
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
					Distro: Ubuntu1804Gen2,
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: Ubuntu1804Gen2,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
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
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.IsUbuntuDistroForAllNodes() != c.expected {
			t.Fatalf("expected IsUbuntuDistroForAllNodes() to return %t but instead returned %t", c.expected, c.p.IsUbuntuDistroForAllNodes())
		}
	}
}

func TestIsVHDDistroForAllNodes(t *testing.T) {
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: AKSUbuntu1604,
					},
				},
			},
			expected: false,
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
			expected: false,
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
						Distro: AKSUbuntu1804,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: Ubuntu1804Gen2,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu1804,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1604,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
					},
				},
			},
			expected: false,
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
						OSType: Windows,
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
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.IsVHDDistroForAllNodes() != c.expected {
			t.Fatalf("expected IsVHDDistroForAllNodes() to return %t but instead returned %t", c.expected, c.p.IsVHDDistroForAllNodes())
		}
	}
}

func TestHasUbuntuDistroNodes(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
					Distro: Ubuntu1804Gen2,
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804Gen2,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1604,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: AKSUbuntu1604,
					},
				},
			},
			expected: false,
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
						OSType: Windows,
					},
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
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
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.HasUbuntuDistroNodes() != c.expected {
			t.Fatalf("expected HasUbuntuDistroNodes() to return %t but instead returned %t", c.expected, c.p.HasUbuntuDistroNodes())
		}
	}
}

func TestHasVHDDistroNodes(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804Gen2,
				},
			},
			expected: false,
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
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: Ubuntu,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804Gen2,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: Ubuntu1804,
					},
				},
			},
			expected: false,
		},
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
					},
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
					},
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
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
						Distro: Ubuntu,
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
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.HasVHDDistroNodes() != c.expected {
			t.Fatalf("expected HasVHDDistroNodes() to return %t but instead returned %t", c.expected, c.p.HasVHDDistroNodes())
		}
	}
}

func TestHasUbuntu1604DistroNodes(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804Gen2,
				},
			},
			expected: false,
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
						Distro: Ubuntu,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1604,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: AKSUbuntu1604,
					},
				},
			},
			expected: false,
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
						OSType: Windows,
					},
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
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
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.HasUbuntu1604DistroNodes() != c.expected {
			t.Fatalf("expected HasUbuntu1604DistroNodes() to return %t but instead returned %t", c.expected, c.p.HasUbuntu1604DistroNodes())
		}
	}
}

func TestHasUbuntu1804DistroNodes(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: AKSUbuntu1604,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu,
				},
			},
			expected: false,
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
					Distro: Ubuntu1804Gen2,
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
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
					{
						Count:  1,
						Distro: Ubuntu,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: Ubuntu,
					},
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
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: AKSUbuntu1604,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						Distro: AKSUbuntu1604,
					},
				},
			},
			expected: false,
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
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				MasterProfile: &MasterProfile{
					Count:  1,
					Distro: Ubuntu1804,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
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
			expected: false,
		},
		{
			p: Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Count:  1,
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.HasUbuntu1804DistroNodes() != c.expected {
			t.Fatalf("expected HasUbuntu1804DistroNodes() to return %t but instead returned %t", c.expected, c.p.HasUbuntu1804DistroNodes())
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
						NetworkPlugin: NetworkPluginAzure,
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
						NetworkPolicy: NetworkPolicyCilium,
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
						NetworkPlugin: NetworkPolicyAntrea,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.RequireRouteTable() != c.expected {
			t.Fatalf("expected RequireRouteTable() to return %t but instead got %t", c.expected, c.p.RequireRouteTable())
		}
	}
}

func TestIsPrivateCluster(t *testing.T) {
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
				},
			},
			expected: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						PrivateCluster: &PrivateCluster{
							Enabled: to.BoolPtr(true),
						},
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
						PrivateCluster: &PrivateCluster{
							Enabled: to.BoolPtr(false),
						},
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
						PrivateCluster: &PrivateCluster{},
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.OrchestratorProfile.IsPrivateCluster() != c.expected {
			t.Fatalf("expected IsPrivateCluster() to return %t but instead got %t", c.expected, c.p.OrchestratorProfile.IsPrivateCluster())
		}
	}
}

func TestIsAzureCNI(t *testing.T) {
	k := &KubernetesConfig{
		NetworkPlugin: NetworkPluginAzure,
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

func TestIsDCOS19(t *testing.T) {
	cases := []struct {
		p                Properties
		expectedIsDCOS19 bool
	}{
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    DCOS,
					OrchestratorVersion: common.DCOSVersion1Dot9Dot8,
				},
			},
			expectedIsDCOS19: true,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    DCOS,
					OrchestratorVersion: "1.9.7",
				},
			},
			expectedIsDCOS19: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
				},
			},
			expectedIsDCOS19: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{},
			},
			expectedIsDCOS19: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
			},
			expectedIsDCOS19: false,
		},
	}

	for _, c := range cases {
		if c.expectedIsDCOS19 != c.p.OrchestratorProfile.IsDCOS19() {
			t.Fatalf("Expected IsDCOS19() to be %t got %t", c.expectedIsDCOS19, c.p.OrchestratorProfile.IsDCOS19())
		}
	}
}

func TestWindowsProfile(t *testing.T) {
	trueVar := true
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

	update := w.GetEnableWindowsUpdate()
	if update {
		t.Fatalf("Expected GetEnableWindowsUpdate() to equal default 'false', got %t", update)
	}

	autoGenerated := w.GetIsCredentialAutoGenerated()
	if autoGenerated {
		t.Fatalf("Expected GetIsCredentialAutoGenerated() to equal default 'false', got %t", autoGenerated)
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
		WindowsImageSourceURL:     "testCustomImage",
		IsCredentialAutoGenerated: to.BoolPtr(true),
	}

	if !(w.HasSecrets() && w.HasCustomImage()) {
		t.Fatalf("Expected HasSecrets() and HasCustomImage() to return true")
	}

	autoGenerated = w.GetIsCredentialAutoGenerated()
	if !autoGenerated {
		t.Fatalf("Expected GetIsCredentialAutoGenerated() to equal default 'true', got %t", autoGenerated)
	}

	w = WindowsProfile{
		WindowsDockerVersion:      "18.03.1-ee-3",
		WindowsSku:                "Datacenter-Core-1809-with-Containers-smalldisk",
		SSHEnabled:                &trueVar,
		IsCredentialAutoGenerated: to.BoolPtr(false),
	}

	dv = w.GetWindowsDockerVersion()
	if dv != "18.03.1-ee-3" {
		t.Fatalf("Expected GetWindowsDockerVersion() to equal 18.03.1-ee-3, got %s", dv)
	}

	windowsSku = w.GetWindowsSku()
	if windowsSku != "Datacenter-Core-1809-with-Containers-smalldisk" {
		t.Fatalf("Expected GetWindowsSku() to equal Datacenter-Core-1809-with-Containers-smalldisk, got %s", windowsSku)
	}

	se := w.GetSSHEnabled()
	if !se {
		t.Fatalf("Expected SSHEnabled to return true, got %v", se)
	}

	autoGenerated = w.GetIsCredentialAutoGenerated()
	if autoGenerated {
		t.Fatalf("Expected GetIsCredentialAutoGenerated() to equal default 'false', got %t", autoGenerated)
	}
}

func TestWindowsProfileCustomOS(t *testing.T) {
	cases := []struct {
		name            string
		w               WindowsProfile
		expectedRef     bool
		expectedGallery bool
		expectedURL     bool
	}{
		{
			name: "valid shared gallery image",
			w: WindowsProfile{
				ImageRef: &ImageReference{
					Name:           "test",
					ResourceGroup:  "testRG",
					SubscriptionID: "testSub",
					Gallery:        "testGallery",
					Version:        "0.1.0",
				},
			},
			expectedRef:     true,
			expectedGallery: true,
			expectedURL:     false,
		},
		{
			name: "valid non-shared image",
			w: WindowsProfile{
				ImageRef: &ImageReference{
					Name:          "test",
					ResourceGroup: "testRG",
				},
			},
			expectedRef:     true,
			expectedGallery: false,
			expectedURL:     false,
		},
		{
			name: "valid image URL",
			w: WindowsProfile{
				WindowsImageSourceURL: "https://some/image.vhd",
			},
			expectedRef:     false,
			expectedGallery: false,
			expectedURL:     true,
		},
		{
			name:            "valid no custom image",
			w:               WindowsProfile{},
			expectedRef:     false,
			expectedGallery: false,
			expectedURL:     false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			if c.w.HasCustomImage() != c.expectedURL {
				t.Errorf("expected HasCustomImage() to return %t but instead returned %t", c.expectedURL, c.w.HasCustomImage())
			}
			if c.w.HasImageRef() != c.expectedRef {
				t.Errorf("expected HasImageRef() to return %t but instead returned %t", c.expectedRef, c.w.HasImageRef())
			}
			if c.w.HasImageGallery() != c.expectedGallery {
				t.Errorf("expected HasImageGallery() to return %t but instead returned %t", c.expectedGallery, c.w.HasImageGallery())
			}
		})
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

func TestPrivateAzureRegistryServerField(t *testing.T) {
	apiloader := &Apiloader{
		Translator: nil,
	}
	apimodel, _, err := apiloader.DeserializeContainerService([]byte(exampleAPIModelWithPrivateAzureRegistry), false, false, nil)
	if err != nil {
		t.Fatalf("unexpectedly error deserializing the example apimodel: %s", err)
	}

	actualPrivateAzureRegistryServer := apimodel.Properties.OrchestratorProfile.KubernetesConfig.PrivateAzureRegistryServer
	if actualPrivateAzureRegistryServer != examplePrivateAzureRegistryServer {
		t.Fatalf("kubernetesConfig->privateAzureRegistryServer field value was unexpected: got(%s), expected(%s)", actualPrivateAzureRegistryServer, examplePrivateAzureRegistryServer)
	}
}

func TestUserAssignedMSI(t *testing.T) {
	// Test1: With just System MSI
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
	c.Addons = append(c.Addons, getMockAddon(common.AADPodIdentityAddonName))
	enabled = c.IsAADPodIdentityEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsAADPodIdentityEnabled() should return default when aad pod identity addon has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.AADPodIdentityAddonName,
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
				Name:    common.AADPodIdentityAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsAADPodIdentityEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsAADPodIdentityEnabled() should return false when aad pod identity addon has been specified as disabled, instead returned %t", enabled)
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
	if enabled {
		t.Fatalf("KubernetesConfig.IsClusterAutoscalerEnabled() should return %t when no cluster autoscaler addon has been specified, instead returned %t", false, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(common.ClusterAutoscalerAddonName))
	enabled = c.IsClusterAutoscalerEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsClusterAutoscalerEnabled() should return false when cluster autoscaler has been specified w/ no enabled value, instead returned %t", enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.ClusterAutoscalerAddonName,
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
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsClusterAutoscalerEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsClusterAutoscalerEnabled() should return false when cluster autoscaler addon has been specified as disabled, instead returned %t", enabled)
	}
}

func TestIsContainerMonitoringEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsContainerMonitoringAddonEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should return %t when no container monitoring addon has been specified, instead returned %t", false, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(common.ContainerMonitoringAddonName))
	enabled = c.IsContainerMonitoringAddonEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should return false when container monitoring addon has been specified w/ no enabled value, instead returned %t", enabled)
	}
	// Addon present and enabled with config
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.ContainerMonitoringAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsContainerMonitoringAddonEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should return true when container monitoring addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.ContainerMonitoringAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsContainerMonitoringAddonEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should return false when container monitoring addon has been specified as disabled, instead returned %t", enabled)
	}

	// Addon present and enabled with logAnalyticsWorkspaceResourceId in config
	b = true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.ContainerMonitoringAddonName,
				Enabled: &b,
				Config: map[string]string{
					"logAnalyticsWorkspaceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-workspace-rg/providers/Microsoft.OperationalInsights/workspaces/test-workspace",
				},
			},
		},
	}
	enabled = c.IsContainerMonitoringAddonEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should return true when container monitoring addon has been specified as enabled, instead returned %t", enabled)
	}

	addon := c.GetAddonByName(common.ContainerMonitoringAddonName)
	if addon.Config == nil || len(addon.Config) == 0 {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should have addon config instead returned null or empty")
	}

	if addon.Config["logAnalyticsWorkspaceResourceId"] == "" {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should have addon config with logAnalyticsWorkspaceResourceId, instead returned null or empty")
	}

	workspaceResourceID := addon.Config["logAnalyticsWorkspaceResourceId"]
	if workspaceResourceID == "" {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should have addon config with non empty azure logAnalyticsWorkspaceResourceId")
	}

	resourceParts := strings.Split(workspaceResourceID, "/")
	if len(resourceParts) != 9 {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should have addon config with valid Azure logAnalyticsWorkspaceResourceId, instead returned %s", workspaceResourceID)
	}

	// Addon present and enabled with legacy config
	b = true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.ContainerMonitoringAddonName,
				Enabled: &b,
				Config: map[string]string{
					"workspaceGuid": "MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAw",
					"workspaceKey":  "NEQrdnlkNS9qU2NCbXNBd1pPRi8wR09CUTVrdUZRYzlKVmFXK0hsbko1OGN5ZVBKY3dUcGtzK3JWbXZnY1hHbW15dWpMRE5FVlBpVDhwQjI3NGE5WWc9PQ==",
				},
			},
		},
	}
	enabled = c.IsContainerMonitoringAddonEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should return true when container monitoring addon has been specified as enabled, instead returned %t", enabled)
	}

	addon = c.GetAddonByName(common.ContainerMonitoringAddonName)
	if addon.Config == nil || len(addon.Config) == 0 {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should have addon config instead returned null or empty")
	}

	if addon.Config["workspaceGuid"] == "" {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should have addon config with non empty workspaceGuid")
	}

	if addon.Config["workspaceKey"] == "" {
		t.Fatalf("KubernetesConfig.IsContainerMonitoringAddonEnabled() should have addon config with non empty workspaceKey")
	}
}

func TestIsNVIDIADevicePluginEnabled(t *testing.T) {
	p := Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
			KubernetesConfig: &KubernetesConfig{
				Addons: []KubernetesAddon{},
			},
		},
	}
	p.OrchestratorProfile.KubernetesConfig.Addons = []KubernetesAddon{
		{
			Name:    common.NVIDIADevicePluginAddonName,
			Enabled: to.BoolPtr(true),
		},
	}

	if !p.IsNVIDIADevicePluginEnabled() {
		t.Fatalf("Properties.IsNVIDIADevicePluginEnabled() should return true with addon enabled, instead returned %t", p.IsNVIDIADevicePluginEnabled())
	}

	p.OrchestratorProfile.KubernetesConfig.Addons = []KubernetesAddon{
		{
			Name:    common.NVIDIADevicePluginAddonName,
			Enabled: to.BoolPtr(false),
		},
	}
	if p.IsNVIDIADevicePluginEnabled() {
		t.Fatalf("Properties.IsNVIDIADevicePluginEnabled() should return false when explicitly disabled")
	}
}

func TestIsAzurePolicyEnabled(t *testing.T) {
	// Default case
	c := KubernetesConfig{
		Addons: []KubernetesAddon{
			getMockAddon("addon"),
		},
	}
	enabled := c.IsAzurePolicyEnabled()
	enabledDefault := DefaultAzurePolicyAddonEnabled
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsAzurePolicyEnabled() should return %t when no azure policy addon has been specified, instead returned %t", enabledDefault, enabled)
	}
	// Addon present, but enabled not specified
	c.Addons = append(c.Addons, getMockAddon(common.AzurePolicyAddonName))
	enabled = c.IsAzurePolicyEnabled()
	if enabled != enabledDefault {
		t.Fatalf("KubernetesConfig.IsAzurePolicyEnabled() should return default when azure policy addon addon has been specified w/ no enabled value, expected %t, instead returned %t", enabledDefault, enabled)
	}
	// Addon present and enabled
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.AzurePolicyAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsAzurePolicyEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsAzurePolicyEnabled() should return true when azure policy addon has been specified as enabled, instead returned %t", enabled)
	}
	// Addon present and disabled
	b = false
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.AzurePolicyAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsAzurePolicyEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsAzurePolicyEnabled() should return false when azure policy addon has been specified as disabled, instead returned %t", enabled)
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
	c.Addons = append(c.Addons, getMockAddon(common.ReschedulerAddonName))
	enabled = c.IsReschedulerEnabled()
	if enabled {
		t.Fatalf("KubernetesConfig.IsReschedulerEnabled() should return true when a custom rescheduler addon has been specified, instead returned %t", enabled)
	}
	b := true
	c = KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.ReschedulerAddonName,
				Enabled: &b,
			},
		},
	}
	enabled = c.IsReschedulerEnabled()
	if !enabled {
		t.Fatalf("KubernetesConfig.IsReschedulerEnabled() should return false when a custom rescheduler addon has been specified as enabled, instead returned %t", enabled)
	}
}

func TestIsIPMasqAgentEnabled(t *testing.T) {
	cases := []struct {
		p                                            Properties
		expectedPropertiesIsIPMasqAgentEnabled       bool
		expectedKubernetesConfigIsIPMasqAgentEnabled bool
	}{
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							getMockAddon(common.IPMASQAgentAddonName),
						},
					},
				},
			},
			expectedPropertiesIsIPMasqAgentEnabled:       false,
			expectedKubernetesConfigIsIPMasqAgentEnabled: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{},
					},
				},
			},
			expectedPropertiesIsIPMasqAgentEnabled:       false,
			expectedKubernetesConfigIsIPMasqAgentEnabled: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name: common.IPMASQAgentAddonName,
								Containers: []KubernetesContainerSpec{
									{
										Name: common.IPMASQAgentAddonName,
									},
								},
							},
						},
					},
				},
			},
			expectedPropertiesIsIPMasqAgentEnabled:       false,
			expectedKubernetesConfigIsIPMasqAgentEnabled: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.IPMASQAgentAddonName,
								Enabled: to.BoolPtr(false),
								Containers: []KubernetesContainerSpec{
									{
										Name: common.IPMASQAgentAddonName,
									},
								},
							},
						},
					},
				},
			},
			expectedPropertiesIsIPMasqAgentEnabled:       false,
			expectedKubernetesConfigIsIPMasqAgentEnabled: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.IPMASQAgentAddonName,
								Enabled: to.BoolPtr(false),
								Containers: []KubernetesContainerSpec{
									{
										Name: common.IPMASQAgentAddonName,
									},
								},
							},
						},
					},
				},
				HostedMasterProfile: &HostedMasterProfile{
					IPMasqAgent: true,
				},
			},
			expectedPropertiesIsIPMasqAgentEnabled:       true,
			expectedKubernetesConfigIsIPMasqAgentEnabled: false, // unsure of the validity of this case, but because it's possible we unit test it
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.IPMASQAgentAddonName,
								Enabled: to.BoolPtr(true),
								Containers: []KubernetesContainerSpec{
									{
										Name: common.IPMASQAgentAddonName,
									},
								},
							},
						},
					},
				},
				HostedMasterProfile: &HostedMasterProfile{
					IPMasqAgent: true,
				},
			},
			expectedPropertiesIsIPMasqAgentEnabled:       true,
			expectedKubernetesConfigIsIPMasqAgentEnabled: true,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.IPMASQAgentAddonName,
								Enabled: to.BoolPtr(true),
								Containers: []KubernetesContainerSpec{
									{
										Name: common.IPMASQAgentAddonName,
									},
								},
							},
						},
					},
				},
				HostedMasterProfile: &HostedMasterProfile{
					IPMasqAgent: false,
				},
			},
			expectedPropertiesIsIPMasqAgentEnabled:       false,
			expectedKubernetesConfigIsIPMasqAgentEnabled: true, // unsure of the validity of this case, but because it's possible we unit test it
		},
	}

	for _, c := range cases {
		if c.p.IsIPMasqAgentEnabled() != c.expectedPropertiesIsIPMasqAgentEnabled {
			t.Fatalf("expected Properties.IsIPMasqAgentEnabled() to return %t but instead returned %t", c.expectedPropertiesIsIPMasqAgentEnabled, c.p.IsIPMasqAgentEnabled())
		}
		if c.p.OrchestratorProfile.KubernetesConfig.IsIPMasqAgentEnabled() != c.expectedKubernetesConfigIsIPMasqAgentEnabled {
			t.Fatalf("expected KubernetesConfig.IsIPMasqAgentEnabled() to return %t but instead returned %t", c.expectedKubernetesConfigIsIPMasqAgentEnabled, c.p.OrchestratorProfile.KubernetesConfig.IsIPMasqAgentEnabled())
		}
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
	p := Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig:    &KubernetesConfig{},
		},
	}
	o := p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

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
		{
			defaultVal:  DefaultKubernetesCloudProviderRateLimitBucketWrite,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucketWrite,
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
		{
			defaultVal:  DefaultKubernetesCloudProviderRateLimitQPSWrite,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPSWrite,
		},
	}

	for _, c := range floatCases {
		if c.computedVal != c.defaultVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.defaultVal, c.computedVal)
		}
	}

	// Test cloudprovider Azure Stack defaults when no user-provided values
	v = "1.14.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig:    &KubernetesConfig{},
		},
		CustomCloudProfile: &CustomCloudProfile{},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

	intCases = []struct {
		defaultVal  int
		computedVal int
	}{
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderRateLimitBucket,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderRateLimitBucketWrite,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucketWrite,
		},
	}

	for _, c := range intCases {
		if c.computedVal != c.defaultVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect Azure Stack default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.defaultVal, c.computedVal)
		}
	}

	floatCases = []struct {
		defaultVal  float64
		computedVal float64
	}{
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderBackoffJitter,
			computedVal: o.KubernetesConfig.CloudProviderBackoffJitter,
		},
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderBackoffExponent,
			computedVal: o.KubernetesConfig.CloudProviderBackoffExponent,
		},
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderRateLimitQPS,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
		{
			defaultVal:  DefaultAzureStackKubernetesCloudProviderRateLimitQPSWrite,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPSWrite,
		},
	}

	for _, c := range floatCases {
		if c.computedVal != c.defaultVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect Azure Stack default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.defaultVal, c.computedVal)
		}
	}

	customCloudProviderBackoffDuration := 99
	customCloudProviderBackoffExponent := 10.0
	customCloudProviderBackoffJitter := 11.9
	customCloudProviderBackoffRetries := 9
	customCloudProviderRateLimitBucket := 37
	customCloudProviderRateLimitQPS := 9.9
	customCloudProviderRateLimitQPSWrite := 100.1
	customCloudProviderRateLimitBucketWrite := 42

	// Test cloudprovider defaults when user provides configuration
	v = "1.8.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig: &KubernetesConfig{
				CloudProviderBackoffDuration:      customCloudProviderBackoffDuration,
				CloudProviderBackoffExponent:      customCloudProviderBackoffExponent,
				CloudProviderBackoffJitter:        customCloudProviderBackoffJitter,
				CloudProviderBackoffRetries:       customCloudProviderBackoffRetries,
				CloudProviderRateLimitBucket:      customCloudProviderRateLimitBucket,
				CloudProviderRateLimitQPS:         customCloudProviderRateLimitQPS,
				CloudProviderRateLimitQPSWrite:    customCloudProviderRateLimitQPSWrite,
				CloudProviderRateLimitBucketWrite: customCloudProviderRateLimitBucketWrite,
			},
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

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
		{
			customVal:   customCloudProviderRateLimitBucketWrite,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucketWrite,
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
		{
			customVal:   customCloudProviderRateLimitQPSWrite,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPSWrite,
		},
	}

	for _, c := range floatCasesCustom {
		if c.computedVal != c.customVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.customVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults when user provides *some* config values
	v = "1.8.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig: &KubernetesConfig{
				CloudProviderBackoffDuration: customCloudProviderBackoffDuration,
				CloudProviderRateLimitBucket: customCloudProviderRateLimitBucket,
				CloudProviderRateLimitQPS:    customCloudProviderRateLimitQPS,
			},
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

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
			expectedVal: customCloudProviderRateLimitQPS,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults for VMSS scenario
	v = "1.14.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig:    &KubernetesConfig{},
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

	intCasesMixed = []struct {
		expectedVal int
		computedVal int
	}{
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			expectedVal: common.MaxAgentCount,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.expectedVal, c.computedVal)
		}
	}

	floatCasesMixed = []struct {
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
			expectedVal: float64(common.MaxAgentCount) * common.MinCloudProviderQPSToBucketFactor,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults for VMSS scenario with 3 pools
	v = "1.14.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig:    &KubernetesConfig{},
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			{
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			{
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

	intCasesMixed = []struct {
		expectedVal int
		computedVal int
	}{
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			expectedVal: common.MaxAgentCount * 3,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.expectedVal, c.computedVal)
		}
	}

	floatCasesMixed = []struct {
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
			expectedVal: float64(common.MaxAgentCount*3) * common.MinCloudProviderQPSToBucketFactor,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults for VMSS scenario + AKS
	v = "1.14.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig:    &KubernetesConfig{},
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
		HostedMasterProfile: &HostedMasterProfile{
			FQDN: "my-cluster",
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

	intCasesMixed = []struct {
		expectedVal int
		computedVal int
	}{
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			expectedVal: common.MaxAgentCount,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.expectedVal, c.computedVal)
		}
	}

	floatCasesMixed = []struct {
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
			expectedVal: float64(common.MaxAgentCount) * common.MinCloudProviderQPSToBucketFactor,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults for VMAS scenario
	v = "1.14.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig:    &KubernetesConfig{},
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				AvailabilityProfile: AvailabilitySet,
			},
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

	intCasesMixed = []struct {
		expectedVal int
		computedVal int
	}{
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			expectedVal: common.MaxAgentCount,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.expectedVal, c.computedVal)
		}
	}

	floatCasesMixed = []struct {
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
			expectedVal: float64(common.MaxAgentCount) * common.MinCloudProviderQPSToBucketFactor,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults for VMAS + VMSS scenario
	v = "1.14.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig:    &KubernetesConfig{},
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				AvailabilityProfile: AvailabilitySet,
			},
			{
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()
	p.SetCloudProviderRateLimitDefaults()

	intCasesMixed = []struct {
		expectedVal int
		computedVal int
	}{
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffRetries,
			computedVal: o.KubernetesConfig.CloudProviderBackoffRetries,
		},
		{
			expectedVal: DefaultKubernetesCloudProviderBackoffDuration,
			computedVal: o.KubernetesConfig.CloudProviderBackoffDuration,
		},
		{
			expectedVal: 2 * common.MaxAgentCount,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitBucket,
		},
	}

	for _, c := range intCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %d, got %d", c.expectedVal, c.computedVal)
		}
	}

	floatCasesMixed = []struct {
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
			expectedVal: float64(common.MaxAgentCount*2) * common.MinCloudProviderQPSToBucketFactor,
			computedVal: o.KubernetesConfig.CloudProviderRateLimitQPS,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig empty cloudprovider configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}

	// Test cloudprovider defaults for backoff mode v2
	v = "1.14.0"
	p = Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorVersion: v,
			KubernetesConfig: &KubernetesConfig{
				CloudProviderBackoffMode: CloudProviderBackoffModeV2,
			},
		},
	}
	o = p.OrchestratorProfile
	p.SetCloudProviderBackoffDefaults()

	floatCasesMixed = []struct {
		expectedVal float64
		computedVal float64
	}{
		{
			expectedVal: 0,
			computedVal: o.KubernetesConfig.CloudProviderBackoffJitter,
		},
		{
			expectedVal: 0,
			computedVal: o.KubernetesConfig.CloudProviderBackoffExponent,
		},
	}

	for _, c := range floatCasesMixed {
		if c.computedVal != c.expectedVal {
			t.Fatalf("KubernetesConfig cloudprovider backoff v2 configs should reflect default values after SetCloudProviderBackoffDefaults(), expected %f, got %f", c.expectedVal, c.computedVal)
		}
	}
}

func TestGetKubernetesVersion(t *testing.T) {
	K8s1dot13dot11 := getMockAPIProperties("1.13.11")
	K8s1dot15dot6 := getMockAPIProperties("1.15.5")
	azureStackProperties := GetMockPropertiesWithCustomCloudProfile(AzureStackCloud, true, true, true)
	azureStackProperties.OrchestratorProfile = K8s1dot13dot11.OrchestratorProfile
	tests := []struct {
		name               string
		properties         *Properties
		expectedK8sVersion string
	}{
		{
			name:               "1.13.11 not Azure Stack",
			properties:         &K8s1dot13dot11,
			expectedK8sVersion: "1.13.11",
		},
		{
			name:               "1.15.5 not Azure Stack",
			properties:         &K8s1dot15dot6,
			expectedK8sVersion: "1.15.5",
		},
		{
			name:               "Azure Stack",
			properties:         &azureStackProperties,
			expectedK8sVersion: "1.13.11" + AzureStackSuffix,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.properties.GetKubernetesVersion()
			if actual != test.expectedK8sVersion {
				t.Errorf("expected k8s version string \"%s\", but got \"%s\"", test.expectedK8sVersion, actual)
			}
		})
	}
}

func TestGetKubernetesHyperkubeSpec(t *testing.T) {
	mock1dot13dot11 := getMockAPIProperties("1.13.11")
	mock1dot13dot11.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "k8s.gcr.io/"
	mock1dot13dot11.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeGCR
	mock1dot16dot3 := getMockAPIProperties("1.16.0")
	mock1dot16dot3.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "k8s.gcr.io/"
	mock1dot16dot3.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeGCR
	mock1dot15dot4azs := GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, true, true)
	mock1dot15dot4azs.OrchestratorProfile = &OrchestratorProfile{
		OrchestratorType:    Kubernetes,
		OrchestratorVersion: "1.15.4",
		KubernetesConfig: &KubernetesConfig{
			KubernetesImageBase:     "mcr.io/",
			KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
		},
	}
	mockcustomproperties := getMockAPIProperties("1.16.0")
	mockcustomproperties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "k8s.gcr.io/"
	mockcustomproperties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage = "mcr.io/my-custom-image"

	tests := []struct {
		name                  string
		properties            *Properties
		expectedHyperkubeSpec string
	}{
		{
			name:                  "1.13.11 Azure public cloud",
			properties:            &mock1dot13dot11,
			expectedHyperkubeSpec: "k8s.gcr.io/hyperkube-amd64:v1.13.11",
		},
		{
			name:                  "1.16.0 Azure public cloud",
			properties:            &mock1dot16dot3,
			expectedHyperkubeSpec: "k8s.gcr.io/hyperkube-amd64:v1.16.0",
		},
		{
			name:                  "1.15.4 Azure Stack",
			properties:            &mock1dot15dot4azs,
			expectedHyperkubeSpec: "mcr.io/oss/kubernetes/hyperkube:v1.15.4-azs",
		},
		{
			name:                  "Custom image",
			properties:            &mockcustomproperties,
			expectedHyperkubeSpec: "mcr.io/my-custom-image",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.properties.GetKubernetesHyperkubeSpec()
			if actual != test.expectedHyperkubeSpec {
				t.Errorf("expected Hyperkube Spec %s, but got %s", test.expectedHyperkubeSpec, actual)
			}
		})
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

	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			Name:                "agentpool",
			VMSize:              "Standard_D2_v2",
			Count:               1,
			AvailabilityProfile: VirtualMachineScaleSets,
		},
	}
	expected = ""
	got = p.GetPrimaryAvailabilitySetName()
	if got != expected {
		t.Errorf("expected primary availability set name %s, but got %s", expected, got)
	}

	p.AgentPoolProfiles = nil
	expected = ""
	got = p.GetPrimaryAvailabilitySetName()
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
		t.Errorf("expected primary scaleset name %s, but got %s", expected, got)
	}

	// Test with Windows agentpool
	p.AgentPoolProfiles[0].OSType = "Windows"

	expected = "2851k8s00"
	got = p.GetPrimaryScaleSetName()
	if got != expected {
		t.Errorf("expected primary scaleset name %s, but got %s", expected, got)
	}

	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			Name:                "agentpool",
			VMSize:              "Standard_D2_v2",
			Count:               1,
			AvailabilityProfile: AvailabilitySet,
		},
	}
	expected = ""
	got = p.GetPrimaryScaleSetName()
	if got != expected {
		t.Errorf("expected primary availability set name %s, but got %s", expected, got)
	}

	p.AgentPoolProfiles = nil
	expected = ""
	got = p.GetPrimaryScaleSetName()
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
		{
			name: "Cluster with VMSS MasterProfile",
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Count:               1,
					DNSPrefix:           "foo",
					VMSize:              "Standard_DS2_v2",
					AvailabilityProfile: VirtualMachineScaleSets,
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
			expectedSubnetName: "subnetmaster",
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

	expectedPrimaryScaleSetName := ""
	if metadata.PrimaryScaleSetName != expectedPrimaryScaleSetName {
		t.Errorf("expected PrimaryScaleSetName name %s, but got %s", expectedPrimaryScaleSetName, metadata.PrimaryScaleSetName)
	}
}

func TestGetAddonPoolIndexByName(t *testing.T) {
	addonName := "testaddon"
	addon := getMockAddon(addonName)
	i := addon.GetAddonPoolIndexByName("pool1")
	if i != 0 {
		t.Fatalf("GetAddonPoolIndexByName() did not return the expected index value 0, instead returned: %d", i)
	}
	i = addon.GetAddonPoolIndexByName("nonExistentContainerName")
	if i != -1 {
		t.Fatalf("GetAddonPoolIndexByName() did not return the expected index value -1, instead returned: %d", i)
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
		t.Fatalf("getAddonContainersIndexByName() did not return the expected index value -1, instead returned: %d", i)
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
			actual := test.properties.GetAgentPoolIndexByName(test.profileName)

			if actual != test.expectedIndex {
				t.Errorf("expected agent pool index %d, but got %d", test.expectedIndex, actual)
			}
		})
	}
}

func TestGetAgentPoolByName(t *testing.T) {
	cases := []struct {
		name          string
		profileName   string
		properties    *Properties
		expectedIndex int
	}{
		{
			name:        "1 pool",
			profileName: "myagentpool",
			properties: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "myagentpool",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
				},
			},
			expectedIndex: 0,
		},
		{
			name:        "3 pool",
			profileName: "myagentpool2",
			properties: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "myagentpool1",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
					{
						Name:   "myagentpool2",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
					{
						Name:   "myagentpool3",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
				},
			},
			expectedIndex: 1,
		},
		{
			name:        "non-existent pool name",
			profileName: "bogon",
			properties: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "myagentpool1",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
					{
						Name:   "myagentpool2",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
					{
						Name:   "myagentpool3",
						VMSize: "Standard_D2_v2",
						Count:  3,
					},
				},
			},
			expectedIndex: -1,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := c.properties.GetAgentPoolByName(c.profileName)
			if c.expectedIndex != -1 {
				if actual != c.properties.AgentPoolProfiles[c.expectedIndex] {
					t.Errorf("expected agent pool %#v, but got %#v", c.properties.AgentPoolProfiles[c.expectedIndex], actual)
				}
			} else {
				if actual != nil {
					t.Errorf("expected nil response, got %#v", actual)
				}
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
				OSType: Linux,
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
						OSType: Linux,
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
				OSType:              Linux,
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
						OSType:              Linux,
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
				OSType: Windows,
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
						OSType: Windows,
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
				OSType: Windows,
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
			actual := p.GetAgentVMPrefix(test.profile, p.GetAgentPoolIndexByName(test.profile.Name))

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
		"santest.germanycentral.cloudapp.microsoftazure.de",
		"santest.germanynorth.cloudapp.azure.com",
		"santest.germanynortheast.cloudapp.microsoftazure.de",
		"santest.germanywestcentral.cloudapp.azure.com",
		"santest.japaneast.cloudapp.azure.com",
		"santest.japanwest.cloudapp.azure.com",
		"santest.koreacentral.cloudapp.azure.com",
		"santest.koreasouth.cloudapp.azure.com",
		"santest.northcentralus.cloudapp.azure.com",
		"santest.northeurope.cloudapp.azure.com",
		"santest.norwayeast.cloudapp.azure.com",
		"santest.norwaywest.cloudapp.azure.com",
		"santest.southafricanorth.cloudapp.azure.com",
		"santest.southafricawest.cloudapp.azure.com",
		"santest.southcentralus.cloudapp.azure.com",
		"santest.southeastasia.cloudapp.azure.com",
		"santest.southindia.cloudapp.azure.com",
		"santest.switzerlandnorth.cloudapp.azure.com",
		"santest.switzerlandwest.cloudapp.azure.com",
		"santest.uaecentral.cloudapp.azure.com",
		"santest.uaenorth.cloudapp.azure.com",
		"santest.uksouth.cloudapp.azure.com",
		"santest.ukwest.cloudapp.azure.com",
		"santest.usdodcentral.cloudapp.usgovcloudapi.net",
		"santest.usdodeast.cloudapp.usgovcloudapi.net",
		"santest.usgovarizona.cloudapp.usgovcloudapi.net",
		"santest.usgoviowa.cloudapp.usgovcloudapi.net",
		"santest.usgovtexas.cloudapp.usgovcloudapi.net",
		"santest.usgovvirginia.cloudapp.usgovcloudapi.net",
		"santest.westcentralus.cloudapp.azure.com",
		"santest.westeurope.cloudapp.azure.com",
		"santest.westindia.cloudapp.azure.com",
		"santest.westus.cloudapp.azure.com",
		"santest.westus2.cloudapp.azure.com",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected formatted fqdns %s, but got %s", expected, actual)
	}

}

func TestFormatProdFQDNByLocation(t *testing.T) {
	// Test locations for Azure
	mockCSDefault := getMockBaseContainerService("1.11.6")
	mockCSDefault.Location = "eastus"
	dnsPrefix := "santest"
	var actual []string
	for _, location := range mockCSDefault.GetLocations() {
		actual = append(actual, FormatProdFQDNByLocation(dnsPrefix, location, mockCSDefault.Properties.GetCustomCloudName()))
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
		"santest.germanycentral.cloudapp.microsoftazure.de",
		"santest.germanynorth.cloudapp.azure.com",
		"santest.germanynortheast.cloudapp.microsoftazure.de",
		"santest.germanywestcentral.cloudapp.azure.com",
		"santest.japaneast.cloudapp.azure.com",
		"santest.japanwest.cloudapp.azure.com",
		"santest.koreacentral.cloudapp.azure.com",
		"santest.koreasouth.cloudapp.azure.com",
		"santest.northcentralus.cloudapp.azure.com",
		"santest.northeurope.cloudapp.azure.com",
		"santest.norwayeast.cloudapp.azure.com",
		"santest.norwaywest.cloudapp.azure.com",
		"santest.southafricanorth.cloudapp.azure.com",
		"santest.southafricawest.cloudapp.azure.com",
		"santest.southcentralus.cloudapp.azure.com",
		"santest.southeastasia.cloudapp.azure.com",
		"santest.southindia.cloudapp.azure.com",
		"santest.switzerlandnorth.cloudapp.azure.com",
		"santest.switzerlandwest.cloudapp.azure.com",
		"santest.uaecentral.cloudapp.azure.com",
		"santest.uaenorth.cloudapp.azure.com",
		"santest.uksouth.cloudapp.azure.com",
		"santest.ukwest.cloudapp.azure.com",
		"santest.usdodcentral.cloudapp.usgovcloudapi.net",
		"santest.usdodeast.cloudapp.usgovcloudapi.net",
		"santest.usgovarizona.cloudapp.usgovcloudapi.net",
		"santest.usgoviowa.cloudapp.usgovcloudapi.net",
		"santest.usgovtexas.cloudapp.usgovcloudapi.net",
		"santest.usgovvirginia.cloudapp.usgovcloudapi.net",
		"santest.westcentralus.cloudapp.azure.com",
		"santest.westeurope.cloudapp.azure.com",
		"santest.westindia.cloudapp.azure.com",
		"santest.westus.cloudapp.azure.com",
		"santest.westus2.cloudapp.azure.com",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected formatted fqdns %s, but got %s", expected, actual)
	}

	// Test location for Azure Stack Cloud
	mockCSDefaultSpec := getMockBaseContainerService("1.11.6")
	mockCSPDefaultSpec := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSDefaultSpec.Properties.CustomCloudProfile = mockCSPDefaultSpec.CustomCloudProfile
	mockCSDefaultSpec.Location = "randomlocation"
	mockCSDefaultSpec.Properties.MasterProfile.DNSPrefix = "azurestackprefix"
	_, err := mockCSDefaultSpec.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		t.Error(err)
	}
	var actualResult []string
	for _, location := range mockCSDefaultSpec.GetLocations() {
		actualResult = append(actualResult, FormatProdFQDNByLocation("azurestackprefix", location, mockCSDefaultSpec.Properties.GetCustomCloudName()))
	}
	expectedResult := []string{"azurestackprefix.randomlocation.cloudapp.azurestack.external"}
	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("Test TestGetLocations() : expected to return %s, but got %s . ", expectedResult, actualResult)
	}
}

func TestKubernetesConfig_GetAddonScript(t *testing.T) {
	addon := getMockAddon(common.IPMASQAgentAddonName)
	addon.Data = "foobarbazdata"
	k := &KubernetesConfig{
		Addons: []KubernetesAddon{
			addon,
		},
	}

	expected := "foobarbazdata"
	actual := k.GetAddonScript(common.IPMASQAgentAddonName)
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

func TestAgentPoolResource(t *testing.T) {
	expectedName := "TestAgentPool"
	expectedVersion := "1.13.0"
	expectedCount := 100

	agentPoolResource := CreateMockAgentPoolProfile(expectedName, expectedVersion, Succeeded, expectedCount)

	gotName := agentPoolResource.Properties.Name
	gotVervsion := agentPoolResource.Properties.OrchestratorVersion
	gotCount := agentPoolResource.Properties.Count

	if gotName != expectedName || gotVervsion != expectedVersion || gotCount != expectedCount {
		t.Fatalf("Expected values - name: %s, version: %s, count: %d. Got - name: %s, version: %s, count: %d", expectedName, expectedVersion, expectedCount,
			gotName, gotVervsion, gotCount)
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
		ContainerRuntime: Docker,
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
				OSType:              Linux,
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
			name:     "telemetry",
			feature:  "EnableTelemetry",
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

func TestKubernetesConfig_UserAssignedIDEnabled(t *testing.T) {
	k := KubernetesConfig{
		UseManagedIdentity: true,
		UserAssignedID:     "fooID",
	}
	if !k.UserAssignedIDEnabled() {
		t.Errorf("expected userAssignedIDEnabled to be true when UseManagedIdentity is true and UserAssignedID is non-empty")
	}

	k = KubernetesConfig{
		UseManagedIdentity: false,
		UserAssignedID:     "fooID",
	}

	if k.UserAssignedIDEnabled() {
		t.Errorf("expected userAssignedIDEnabled to be false when useManagedIdentity is set to false")
	}
}

func TestKubernetesConfig_ShouldCreateNewUserAssignedIdentity(t *testing.T) {
	k := KubernetesConfig{
		UseManagedIdentity: true,
		UserAssignedID:     "fooID",
	}
	if !k.ShouldCreateNewUserAssignedIdentity() {
		t.Errorf("expected ShouldCreateNewUserAssignedIdentity to be false when UserAssignedID does not have a resource id")
	}

	k = KubernetesConfig{
		UseManagedIdentity: true,
		UserAssignedID:     exampleUserMSI,
	}

	if !k.UserAssignedIDEnabled() {
		t.Errorf("expected ShouldCreateNewUserAssignedIdentity to be true when UserAssignedID has a resource id")
	}
}

func TestKubernetesConfig_SystemAssignedIDEnabled(t *testing.T) {
	k := KubernetesConfig{
		UseManagedIdentity: true,
		UserAssignedID:     "",
	}
	if !k.SystemAssignedIDEnabled() {
		t.Errorf("expected SystemAssignedIDEnabled to be true when UseManagedIdentity is true and UserAssignedID is empty")
	}

	k = KubernetesConfig{
		UseManagedIdentity: true,
		UserAssignedID:     "foo",
	}

	if k.SystemAssignedIDEnabled() {
		t.Errorf("expected SystemAssignedIDEnabled to be true when UseManagedIdentity is true and UserAssignedID is non-empty")
	}

	k = KubernetesConfig{
		UseManagedIdentity: false,
		UserAssignedID:     "",
	}

	if k.SystemAssignedIDEnabled() {
		t.Errorf("expected SystemAssignedIDEnabled to be false when UseManagedIdentity is set to false")
	}
}

func TestGetCustomCloudRootCertificates(t *testing.T) {
	var emptyCustomRootCertsProfile = GetMockPropertiesWithCustomCloudProfile("CustomCloud", true, true, true)
	emptyCustomRootCertsProfile.CustomCloudProfile.CustomCloudRootCertificates = ""

	var nonEmptyCustomRootCertsProfile = GetMockPropertiesWithCustomCloudProfile("CustomCloud", true, true, true)
	nonEmptyCustomRootCertsProfile.CustomCloudProfile.CustomCloudRootCertificates = "cert1,cert2,cert3"

	testcases := []struct {
		name       string
		properties Properties
		expected   string
	}{
		{
			"Empty CustomCloudRootCertificates",
			emptyCustomRootCertsProfile,
			"",
		},
		{
			"Non-Empty CustomCloudRootCertificates",
			nonEmptyCustomRootCertsProfile,
			"cert1,cert2,cert3",
		},
	}
	for _, testcase := range testcases {
		actual := testcase.properties.GetCustomCloudRootCertificates()
		if testcase.expected != actual {
			t.Errorf("Test \"%s\": expected GetCustomCloudRootCertificates() to return %s, but got %s . ", testcase.name, testcase.expected, actual)
		}
	}
}

func TestGetCustomCloudSourcesList(t *testing.T) {
	var emptyCustomCloudSourcesListProfile = GetMockPropertiesWithCustomCloudProfile("CustomCloud", true, true, true)
	emptyCustomCloudSourcesListProfile.CustomCloudProfile.CustomCloudSourcesList = ""

	var nonEmptyCustomCloudSourcesListProfile = GetMockPropertiesWithCustomCloudProfile("CustomCloud", true, true, true)
	nonEmptyCustomCloudSourcesListProfile.CustomCloudProfile.CustomCloudSourcesList = "test"

	testcases := []struct {
		name       string
		properties Properties
		expected   string
	}{
		{
			"Empty CustomCloudSourcesList",
			emptyCustomCloudSourcesListProfile,
			"",
		},
		{
			"Non-Empty CustomCloudSourcesList",
			nonEmptyCustomCloudSourcesListProfile,
			"test",
		},
	}
	for _, testcase := range testcases {
		actual := testcase.properties.GetCustomCloudSourcesList()
		if testcase.expected != actual {
			t.Errorf("Test \"%s\": expected GetCustomCloudSourcesList() to return %s, but got %s . ", testcase.name, testcase.expected, actual)
		}
	}
}

func TestIsCustomCloudProfile(t *testing.T) {
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
		actual := testcase.properties.IsCustomCloudProfile()
		if testcase.expected != actual {
			t.Errorf("Test \"%s\": expected IsCustomCloudProfile() to return %t, but got %t . ", testcase.name, testcase.expected, actual)
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
			"Empty environment name should be treated as AzureStackCloud",
			GetMockPropertiesWithCustomCloudProfile("", true, true, false),
			true,
		},
		{
			"Empty environment name (with AzureEnvironmentSpecConfig) should be treated as AzureStackCloud",
			GetMockPropertiesWithCustomCloudProfile("", true, true, true),
			true,
		},
		{
			"lower case AzureStackCloud name",
			GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true),
			true,
		},
		{
			"cammel case AzureStackCloud name",
			GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, true, true),
			true,
		},
		{
			"incorrect AzureStackCloud name",
			GetMockPropertiesWithCustomCloudProfile("NotAzureStackCloud", true, true, true),
			false,
		},
		{
			"empty cloud profile",
			GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", false, false, false),
			false,
		},
		{
			"empty environment should be treated as AzureStackCloud",
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

func TestGetCustomCloudName(t *testing.T) {
	testcases := []struct {
		name       string
		properties Properties
		expected   string
	}{
		{
			"lower case cloud name",
			GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true),
			"azurestackcloud",
		},
		{
			"cammel case cloud name",
			GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, true, true),
			"AzureStackCloud",
		},
	}
	for _, testcase := range testcases {
		actual := testcase.properties.GetCustomCloudName()
		if testcase.expected != actual {
			t.Errorf("Test \"%s\": expected GetCustomCloudName() to return %s, but got %s . ", testcase.name, testcase.expected, actual)
		}
	}
}

func TestGetCustomEnvironmentJSON(t *testing.T) {
	expectedAzureStackResult := `{"name":"azurestackcloud","managementPortalURL":"https://management.local.azurestack.external/","publishSettingsURL":"https://management.local.azurestack.external/publishsettings/index","serviceManagementEndpoint":"https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189","resourceManagerEndpoint":"https://management.local.azurestack.external/","activeDirectoryEndpoint":"https://login.windows.net/","galleryEndpoint":"https://portal.local.azurestack.external=30015/","keyVaultEndpoint":"https://vault.azurestack.external/","graphEndpoint":"https://graph.windows.net/","serviceBusEndpoint":"https://servicebus.azurestack.external/","batchManagementEndpoint":"https://batch.azurestack.external/","storageEndpointSuffix":"core.azurestack.external","sqlDatabaseDNSSuffix":"database.azurestack.external","trafficManagerDNSSuffix":"trafficmanager.cn","keyVaultDNSSuffix":"vault.azurestack.external","serviceBusEndpointSuffix":"servicebus.azurestack.external","serviceManagementVMDNSSuffix":"chinacloudapp.cn","resourceManagerVMDNSSuffix":"cloudapp.azurestack.external","containerRegistryDNSSuffix":"azurecr.io","cosmosDBDNSSuffix":"","tokenAudience":"https://management.azurestack.external/","resourceIdentifiers":{"graph":"","keyVault":"","datalake":"","batch":"","operationalInsights":"","storage":""}}`
	expectedNonAzureStackResult := `{"name":"nonazurestackcloud","managementPortalURL":"https://management.local.azurestack.external/","publishSettingsURL":"https://management.local.azurestack.external/publishsettings/index","serviceManagementEndpoint":"https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189","resourceManagerEndpoint":"https://management.local.azurestack.external/","activeDirectoryEndpoint":"https://login.windows.net/","galleryEndpoint":"https://portal.local.azurestack.external=30015/","keyVaultEndpoint":"https://vault.azurestack.external/","graphEndpoint":"https://graph.windows.net/","serviceBusEndpoint":"https://servicebus.azurestack.external/","batchManagementEndpoint":"https://batch.azurestack.external/","storageEndpointSuffix":"core.azurestack.external","sqlDatabaseDNSSuffix":"database.azurestack.external","trafficManagerDNSSuffix":"trafficmanager.cn","keyVaultDNSSuffix":"vault.azurestack.external","serviceBusEndpointSuffix":"servicebus.azurestack.external","serviceManagementVMDNSSuffix":"chinacloudapp.cn","resourceManagerVMDNSSuffix":"cloudapp.azurestack.external","containerRegistryDNSSuffix":"azurecr.io","cosmosDBDNSSuffix":"","tokenAudience":"https://management.azurestack.external/","resourceIdentifiers":{"graph":"","keyVault":"","datalake":"","batch":"","operationalInsights":"","storage":""}}`

	testcases := []struct {
		name       string
		properties Properties
		escape     bool
		expected   string
	}{
		{
			"escape",
			GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true),
			true,
			strings.Replace(expectedAzureStackResult, "\"", "\\\"", -1),
		},
		{
			"no escape",
			GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true),
			false,
			expectedAzureStackResult,
		},
		{
			"escape",
			GetMockPropertiesWithCustomCloudProfile("nonazurestackcloud", true, true, true),
			true,
			strings.Replace(expectedNonAzureStackResult, "\"", "\\\"", -1),
		},
		{
			"no escape",
			GetMockPropertiesWithCustomCloudProfile("nonazurestackcloud", true, true, true),
			false,
			expectedNonAzureStackResult,
		},
	}
	for _, testcase := range testcases {
		actual, err := testcase.properties.GetCustomEnvironmentJSON(testcase.escape)
		if err != nil {
			t.Error(err)
		}
		if testcase.expected != actual {
			t.Errorf("Test \"%s\": expected GetCustomEnvironmentJSON() to return %s, but got %s . ", testcase.name, testcase.expected, actual)
		}
	}
}

func TestGetLocations(t *testing.T) {

	// Test location for Azure Stack Cloud
	mockCSDefaultSpec := getMockBaseContainerService("1.11.6")
	mockCSPDefaultSpec := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSDefaultSpec.Properties.CustomCloudProfile = mockCSPDefaultSpec.CustomCloudProfile
	mockCSDefaultSpec.Location = "randomlocation"

	expectedResult := []string{"randomlocation"}
	actualResult := mockCSDefaultSpec.GetLocations()
	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("Test TestGetLocations() : expected to return %s, but got %s . ", expectedResult, actualResult)
	}

	// Test locations for Azure
	mockCSDefault := getMockBaseContainerService("1.11.6")
	mockCSDefault.Location = "eastus"

	expected := []string{
		"australiacentral",
		"australiacentral2",
		"australiaeast",
		"australiasoutheast",
		"brazilsouth",
		"canadacentral",
		"canadaeast",
		"centralindia",
		"centralus",
		"centraluseuap",
		"chinaeast",
		"chinaeast2",
		"chinanorth",
		"chinanorth2",
		"eastasia",
		"eastus",
		"eastus2",
		"eastus2euap",
		"francecentral",
		"francesouth",
		"germanycentral",
		"germanynorth",
		"germanynortheast",
		"germanywestcentral",
		"japaneast",
		"japanwest",
		"koreacentral",
		"koreasouth",
		"northcentralus",
		"northeurope",
		"norwayeast",
		"norwaywest",
		"southafricanorth",
		"southafricawest",
		"southcentralus",
		"southeastasia",
		"southindia",
		"switzerlandnorth",
		"switzerlandwest",
		"uaecentral",
		"uaenorth",
		"uksouth",
		"ukwest",
		"usdodcentral",
		"usdodeast",
		"usgovarizona",
		"usgoviowa",
		"usgovtexas",
		"usgovvirginia",
		"westcentralus",
		"westeurope",
		"westindia",
		"westus",
		"westus2",
	}
	actual := mockCSDefault.GetLocations()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Test TestGetLocations() : expected to return %s, but got %s . ", expected, actual)
	}
}

func TestGetMasterFQDN(t *testing.T) {
	tests := []struct {
		name              string
		properties        *Properties
		expectedFQDN      string
		expectedDNSPrefix string
	}{
		{
			name: "From Master Profile",
			properties: &Properties{
				MasterProfile: &MasterProfile{
					DNSPrefix: "foo_master",
					FQDN:      "FQDNFromMasterProfile",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo_agent0",
					},
				},
			},
			expectedFQDN:      "FQDNFromMasterProfile",
			expectedDNSPrefix: "foo_master",
		},
		{
			name: "From Hosted Master Profile",
			properties: &Properties{
				HostedMasterProfile: &HostedMasterProfile{
					DNSPrefix: "foo_hosted_master",
					FQDN:      "FQDNFromHostedMasterProfile",
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo_agent1",
					},
				},
			},
			expectedFQDN:      "FQDNFromHostedMasterProfile",
			expectedDNSPrefix: "foo_hosted_master",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actualFQDN := test.properties.GetMasterFQDN()
			if actualFQDN != test.expectedFQDN {
				t.Errorf("expected fqdn %s, but got %s", test.expectedFQDN, actualFQDN)
			}
			actualDNSPrefix := test.properties.GetDNSPrefix()
			if actualDNSPrefix != test.expectedDNSPrefix {
				t.Errorf("expected DNS prefix %s, but got %s", test.expectedDNSPrefix, actualDNSPrefix)
			}
		})
	}
}

func TestGetKubeProxyFeatureGates(t *testing.T) {
	tests := []struct {
		name                 string
		properties           *Properties
		expectedFeatureGates string
	}{
		{
			name: "default",
			properties: &Properties{
				FeatureFlags: &FeatureFlags{},
			},
			expectedFeatureGates: "{}",
		},
		{
			name: "IPV6 enabled",
			properties: &Properties{
				FeatureFlags: &FeatureFlags{
					EnableIPv6DualStack: true,
				},
			},
			expectedFeatureGates: "IPv6DualStack: true",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := test.properties.GetKubeProxyFeatureGates()
			if actual != test.expectedFeatureGates {
				t.Errorf("expected featureGates %s, but got %s", test.expectedFeatureGates, actual)
			}
		})
	}
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

func TestDcosConfigHasPrivateRegistry(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
					DcosConfig: &DcosConfig{
						Registry: "my-custom-registry",
					},
				},
			},
			expected: true,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
					DcosConfig: &DcosConfig{
						Registry: "",
					},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
					DcosConfig:       &DcosConfig{},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.p.OrchestratorProfile.DcosConfig.HasPrivateRegistry() != c.expected {
			t.Fatalf("expected HasPrivateRegistry() to return %t but instead got %t", c.expected, c.p.OrchestratorProfile.DcosConfig.HasPrivateRegistry())
		}
	}
}

func TestDcosConfigHasBootstrap(t *testing.T) {
	cases := []struct {
		p        Properties
		expected bool
	}{
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
					DcosConfig:       &DcosConfig{},
				},
			},
			expected: false,
		},
		{
			p: Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: DCOS,
					DcosConfig: &DcosConfig{
						BootstrapProfile: &BootstrapProfile{},
					},
				},
			},
			expected: true,
		},
	}

	for _, c := range cases {
		if c.p.OrchestratorProfile.DcosConfig.HasBootstrap() != c.expected {
			t.Fatalf("expected HasBootstrap() to return %t but instead got %t", c.expected, c.p.OrchestratorProfile.DcosConfig.HasBootstrap())
		}
	}
}

func TestKubernetesAddonIsEnabled(t *testing.T) {
	cases := []struct {
		a        *KubernetesAddon
		expected bool
	}{
		{
			a:        &KubernetesAddon{},
			expected: false,
		},
		{
			a: &KubernetesAddon{
				Enabled: to.BoolPtr(false),
			},
			expected: false,
		},
		{
			a: &KubernetesAddon{
				Enabled: to.BoolPtr(true),
			},
			expected: true,
		},
	}

	for _, c := range cases {
		if c.a.IsEnabled() != c.expected {
			t.Fatalf("expected IsEnabled() to return %t but instead returned %t", c.expected, c.a.IsEnabled())
		}
	}
}

func TestKubernetesAddonIsDisabled(t *testing.T) {
	cases := []struct {
		a        *KubernetesAddon
		expected bool
	}{
		{
			a:        &KubernetesAddon{},
			expected: false,
		},
		{
			a: &KubernetesAddon{
				Enabled: to.BoolPtr(false),
			},
			expected: true,
		},
		{
			a: &KubernetesAddon{
				Enabled: to.BoolPtr(true),
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.a.IsDisabled() != c.expected {
			t.Fatalf("expected IsDisabled() to return %t but instead returned %t", c.expected, c.a.IsDisabled())
		}
	}
}

func TestKubernetesConfigIsAddonEnabled(t *testing.T) {
	cases := []struct {
		k         *KubernetesConfig
		addonName string
		expected  bool
	}{
		{
			k:         &KubernetesConfig{},
			addonName: "foo",
			expected:  false,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name: "foo",
					},
				},
			},
			addonName: "foo",
			expected:  false,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    "foo",
						Enabled: to.BoolPtr(false),
					},
				},
			},
			addonName: "foo",
			expected:  false,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    "foo",
						Enabled: to.BoolPtr(true),
					},
				},
			},
			addonName: "foo",
			expected:  true,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    "bar",
						Enabled: to.BoolPtr(true),
					},
				},
			},
			addonName: "foo",
			expected:  false,
		},
	}

	for _, c := range cases {
		if c.k.IsAddonEnabled(c.addonName) != c.expected {
			t.Fatalf("expected KubernetesConfig.IsAddonEnabled(%s) to return %t but instead returned %t", c.addonName, c.expected, c.k.IsAddonEnabled(c.addonName))
		}
	}
}

func TestKubernetesConfigIsAddonDisabled(t *testing.T) {
	cases := []struct {
		k         *KubernetesConfig
		addonName string
		expected  bool
	}{
		{
			k:         &KubernetesConfig{},
			addonName: "foo",
			expected:  false,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name: "foo",
					},
				},
			},
			addonName: "foo",
			expected:  false,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    "foo",
						Enabled: to.BoolPtr(false),
					},
				},
			},
			addonName: "foo",
			expected:  true,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    "foo",
						Enabled: to.BoolPtr(true),
					},
				},
			},
			addonName: "foo",
			expected:  false,
		},
		{
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    "bar",
						Enabled: to.BoolPtr(true),
					},
				},
			},
			addonName: "foo",
			expected:  false,
		},
	}

	for _, c := range cases {
		if c.k.IsAddonDisabled(c.addonName) != c.expected {
			t.Fatalf("expected KubernetesConfig.IsAddonDisabled(%s) to return %t but instead returned %t", c.addonName, c.expected, c.k.IsAddonDisabled(c.addonName))
		}
	}
}

func TestKubernetesComponentIsEnabled(t *testing.T) {
	cases := []struct {
		c        *KubernetesComponent
		expected bool
	}{
		{
			c:        &KubernetesComponent{},
			expected: false,
		},
		{
			c: &KubernetesComponent{
				Enabled: to.BoolPtr(false),
			},
			expected: false,
		},
		{
			c: &KubernetesComponent{
				Enabled: to.BoolPtr(true),
			},
			expected: true,
		},
	}

	for _, c := range cases {
		if c.c.IsEnabled() != c.expected {
			t.Fatalf("expected IsEnabled() to return %t but instead returned %t", c.expected, c.c.IsEnabled())
		}
	}
}

func TestKubernetesComponentIsDisabled(t *testing.T) {
	cases := []struct {
		c        *KubernetesComponent
		expected bool
	}{
		{
			c:        &KubernetesComponent{},
			expected: false,
		},
		{
			c: &KubernetesComponent{
				Enabled: to.BoolPtr(false),
			},
			expected: true,
		},
		{
			c: &KubernetesComponent{
				Enabled: to.BoolPtr(true),
			},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.c.IsDisabled() != c.expected {
			t.Fatalf("expected IsDisabled() to return %t but instead returned %t", c.expected, c.c.IsDisabled())
		}
	}
}

func TestKubernetesConfigIsComponentEnabled(t *testing.T) {
	cases := []struct {
		k             *KubernetesConfig
		componentName string
		expected      bool
	}{
		{
			k:             &KubernetesConfig{},
			componentName: "foo",
			expected:      false,
		},
		{
			k: &KubernetesConfig{
				Components: []KubernetesComponent{
					{
						Name: "foo",
					},
				},
			},
			componentName: "foo",
			expected:      false,
		},
		{
			k: &KubernetesConfig{
				Components: []KubernetesComponent{
					{
						Name:    "foo",
						Enabled: to.BoolPtr(false),
					},
				},
			},
			componentName: "foo",
			expected:      false,
		},
		{
			k: &KubernetesConfig{
				Components: []KubernetesComponent{
					{
						Name:    "foo",
						Enabled: to.BoolPtr(true),
					},
				},
			},
			componentName: "foo",
			expected:      true,
		},
		{
			k: &KubernetesConfig{
				Components: []KubernetesComponent{
					{
						Name:    "bar",
						Enabled: to.BoolPtr(true),
					},
				},
			},
			componentName: "foo",
			expected:      false,
		},
	}

	for _, c := range cases {
		component, isEnabled := c.k.IsComponentEnabled(c.componentName)
		if isEnabled != c.expected {
			t.Fatalf("expected KubernetesConfig.IsComponentEnabled(%s) to return %t but instead returned %t", c.componentName, c.expected, isEnabled)
		}
		if !reflect.DeepEqual(component, c.k.GetComponentByName(c.componentName)) {
			t.Fatalf("expected result component %v to be equal to %v", component, c.k.GetComponentByName(c.componentName))
		}
	}
}

func TestKubernetesConfigComponentGetters(t *testing.T) {
	cases := []struct {
		k            *KubernetesConfig
		expectedData string
	}{
		{
			k:            &KubernetesConfig{},
			expectedData: "",
		},
		{
			k: &KubernetesConfig{
				Components: []KubernetesComponent{
					{
						Name: "foo",
						Data: "bar",
					},
				},
			},
			expectedData: "bar",
		},
		{
			k: &KubernetesConfig{
				Components: []KubernetesComponent{
					{
						Name:    "foo",
						Enabled: to.BoolPtr(false),
					},
				},
			},
			expectedData: "",
		},
	}

	for _, c := range cases {
		for _, component := range c.k.Components {
			got := c.k.GetComponentByName(component.Name)
			if !reflect.DeepEqual(got, component) {
				t.Fatalf("expected result component %v to be equal to %v", got, component)
			}
			data := c.k.GetComponentData(component.Name)
			if data != c.expectedData {
				t.Fatalf("expected KubernetesConfig.GetComponentData(%s) to return %s but instead returned %s", component.Name, c.expectedData, data)
			}
		}
	}
}

func TestHasVMASAgentPool(t *testing.T) {
	tests := []struct {
		name     string
		p        *Properties
		expected bool
	}{
		{
			name: "one agent pool w/ AvailabilitySet",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool1",
						VMSize:              "Standard_D2_v2",
						Count:               2,
						AvailabilityProfile: AvailabilitySet,
					},
				},
			},
			expected: true,
		},
		{
			name: "two agent pools, one w/ AvailabilitySet",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "agentpool1",
						VMSize:              "Standard_D2_v2",
						Count:               2,
						AvailabilityProfile: AvailabilitySet,
					},
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  100,
					},
				},
			},
			expected: true,
		},
		{
			name: "two agent pools",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  100,
					},
				},
			},
			expected: false,
		},
		{
			name: "two agent pools, one w/ VirtualMachineScaleSets",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
					{
						Name:                "agentpool1",
						VMSize:              "Standard_D2_v2",
						Count:               100,
						AvailabilityProfile: VirtualMachineScaleSets,
					},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := test.p.HasVMASAgentPool()
			if test.expected != ret {
				t.Errorf("expected %t, instead got : %t", test.expected, ret)
			}
		})
	}
}

func TestAnyAgentIsLinux(t *testing.T) {
	tests := []struct {
		name     string
		p        *Properties
		expected bool
	}{
		{
			name: "one agent pool w/ Linux",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  2,
						OSType: Linux,
					},
				},
			},
			expected: true,
		},
		{
			name: "two agent pools, one w/ Linux",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  2,
						OSType: Windows,
					},
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						OSType: Linux,
					},
				},
			},
			expected: true,
		},
		{
			name: "two agent pools",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  100,
					},
				},
			},
			expected: false,
		},
		{
			name: "two agent pools, one w/ Windows",
			p: &Properties{
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  2,
					},
					{
						Name:   "agentpool1",
						VMSize: "Standard_D2_v2",
						Count:  100,
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := test.p.AnyAgentIsLinux()
			if test.expected != ret {
				t.Errorf("expected %t, instead got : %t", test.expected, ret)
			}
		})
	}
}

func TestHasContainerd(t *testing.T) {
	tests := []struct {
		name     string
		k        *KubernetesConfig
		expected bool
	}{
		{
			name: "docker",
			k: &KubernetesConfig{
				ContainerRuntime: Docker,
			},
			expected: false,
		},
		{
			name: "empty string",
			k: &KubernetesConfig{
				ContainerRuntime: "",
			},
			expected: false,
		},
		{
			name: "unexpected string",
			k: &KubernetesConfig{
				ContainerRuntime: "foo",
			},
			expected: false,
		},
		{
			name: "containerd",
			k: &KubernetesConfig{
				ContainerRuntime: Containerd,
			},
			expected: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := test.k.NeedsContainerd()
			if test.expected != ret {
				t.Errorf("expected %t, instead got : %t", test.expected, ret)
			}
		})
	}
}

func TestGetNonMasqueradeCIDR(t *testing.T) {
	tests := []struct {
		name     string
		p        *Properties
		expected string
	}{
		{
			name: "single cluster cidr, no dualstack",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ClusterSubnet: "10.244.0.0/16",
					},
				},
				FeatureFlags: &FeatureFlags{},
			},
			expected: "10.244.0.0/16",
		},
		{
			name: "two cluster cidr v4v6, dualstack",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ClusterSubnet: "10.244.0.0/16,fd00:101::/8",
					},
				},
				FeatureFlags: &FeatureFlags{
					EnableIPv6DualStack: true,
				},
			},
			expected: "10.244.0.0/16",
		},
		{
			name: "two cluster cidr v6v4, dualstack",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ClusterSubnet: "fd00:101::/8,10.244.0.0/16",
					},
				},
				FeatureFlags: &FeatureFlags{
					EnableIPv6DualStack: true,
				},
			},
			expected: "fd00::/8",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := test.p.GetNonMasqueradeCIDR()
			if test.expected != ret {
				t.Errorf("expected %s, instead got : %s", test.expected, ret)
			}
		})
	}
}

func TestGetSecondaryNonMasqueradeCIDR(t *testing.T) {
	tests := []struct {
		name     string
		p        *Properties
		expected string
	}{
		{
			name: "single cluster cidr, no dualstack",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ClusterSubnet: "10.244.0.0/16",
					},
				},
				FeatureFlags: &FeatureFlags{},
			},
			expected: "",
		},
		{
			name: "two cluster cidr v4v6, dualstack",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ClusterSubnet: "10.244.0.0/16,fd00:101::/8",
					},
				},
				FeatureFlags: &FeatureFlags{
					EnableIPv6DualStack: true,
				},
			},
			expected: "fd00::/8",
		},
		{
			name: "two cluster cidr v6v4, dualstack",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ClusterSubnet: "fd00:101::/8,10.244.0.0/16",
					},
				},
				FeatureFlags: &FeatureFlags{
					EnableIPv6DualStack: true,
				},
			},
			expected: "10.244.0.0/16",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := test.p.GetSecondaryNonMasqueradeCIDR()
			if test.expected != ret {
				t.Errorf("expected %s, instead got : %s", test.expected, ret)
			}
		})
	}
}

func TestPropertiesHasDCSeriesSKU(t *testing.T) {
	cases := common.GetDCSeriesVMCasesForTesting()

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
				OrchestratorVersion: "1.16.0",
			},
		}
		ret := p.HasDCSeriesSKU()
		if ret != c.Expected {
			t.Fatalf("expected HasDCSeriesSKU(%s) to return %t, but instead got %t", c.VMSKU, c.Expected, ret)
		}
	}
}

func TestPropertiesIsIPMasqAgentDisabled(t *testing.T) {
	cases := []struct {
		name             string
		p                *Properties
		expectedDisabled bool
	}{
		{
			name:             "default",
			p:                &Properties{},
			expectedDisabled: false,
		},
		{
			name: "hostedMasterProfile disabled",
			p: &Properties{
				HostedMasterProfile: &HostedMasterProfile{
					IPMasqAgent: false,
				},
			},
			expectedDisabled: true,
		},
		{
			name: "hostedMasterProfile enabled",
			p: &Properties{
				HostedMasterProfile: &HostedMasterProfile{
					IPMasqAgent: true,
				},
			},
			expectedDisabled: false,
		},
		{
			name: "nil KubernetesConfig",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{},
			},
			expectedDisabled: false,
		},
		{
			name: "default KubernetesConfig",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{},
				},
			},
			expectedDisabled: false,
		},
		{
			name: "addons configured but no ip-masq-agent configuration",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.CoreDNSAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedDisabled: false,
		},
		{
			name: "ip-masq-agent explicitly disabled",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.IPMASQAgentAddonName,
								Enabled: to.BoolPtr(false),
							},
						},
					},
				},
			},
			expectedDisabled: true,
		},
		{
			name: "ip-masq-agent present but no configuration",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name: common.IPMASQAgentAddonName,
							},
						},
					},
				},
			},
			expectedDisabled: false,
		},
		{
			name: "ip-masq-agent explicitly enabled",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.IPMASQAgentAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedDisabled: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.p.IsIPMasqAgentDisabled() != c.expectedDisabled {
				t.Fatalf("expected Properties.IsIPMasqAgentDisabled() to return %t but instead returned %t", c.expectedDisabled, c.p.IsIPMasqAgentDisabled())
			}
		})
	}
}

func TestKubernetesConfigIsIPMasqAgentDisabled(t *testing.T) {
	cases := []struct {
		name             string
		k                *KubernetesConfig
		expectedDisabled bool
	}{
		{
			name:             "default",
			k:                &KubernetesConfig{},
			expectedDisabled: false,
		},
		{
			name: "ip-masq-agent present but no configuration",
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name: common.IPMASQAgentAddonName,
					},
				},
			},
			expectedDisabled: false,
		},
		{
			name: "ip-masq-agent explicitly disabled",
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    common.IPMASQAgentAddonName,
						Enabled: to.BoolPtr(false),
					},
				},
			},
			expectedDisabled: true,
		},
		{
			name: "ip-masq-agent explicitly enabled",
			k: &KubernetesConfig{
				Addons: []KubernetesAddon{
					{
						Name:    common.IPMASQAgentAddonName,
						Enabled: to.BoolPtr(true),
					},
				},
			},
			expectedDisabled: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.k.IsIPMasqAgentDisabled() != c.expectedDisabled {
				t.Fatalf("expected KubernetesConfig.IsIPMasqAgentDisabled() to return %t but instead returned %t", c.expectedDisabled, c.k.IsIPMasqAgentDisabled())
			}
		})
	}
}

func TestGetProvisionScriptParametersCommon(t *testing.T) {
	cases := []struct {
		name     string
		cs       *ContainerService
		input    ProvisionScriptParametersInput
		expected string
	}{
		{
			name: "Default container service with no ARM variables",
			cs:   CreateMockContainerService("testcluster", "1.16.8", 1, 3, true),
			input: ProvisionScriptParametersInput{
				Location:             "westus",
				ResourceGroup:        "fakerg",
				TenantID:             "faketenantID",
				SubscriptionID:       "fakesubID",
				ClientID:             "fakeclientID",
				ClientSecret:         "fakeclientSecret",
				APIServerCertificate: "fakecert",
				KubeletPrivateKey:    "fakekubeletkey",
				ClusterKeyVaultName:  "",
			},
			expected: "ADMINUSER=azureuser APISERVER_PUBLIC_KEY=fakecert AUTHENTICATION_METHOD=client_secret CLOUDPROVIDER_BACKOFF=false CLOUDPROVIDER_BACKOFF_DURATION=0 CLOUDPROVIDER_BACKOFF_EXPONENT=0 CLOUDPROVIDER_BACKOFF_JITTER=0 CLOUDPROVIDER_BACKOFF_MODE= CLOUDPROVIDER_BACKOFF_RETRIES=0 CLOUDPROVIDER_RATELIMIT=false CLOUDPROVIDER_RATELIMIT_BUCKET=0 CLOUDPROVIDER_RATELIMIT_BUCKET_WRITE=0 CLOUDPROVIDER_RATELIMIT_QPS=0 CLOUDPROVIDER_RATELIMIT_QPS_WRITE=0 CNI_PLUGINS_URL=https://kubernetesartifacts.azureedge.net/cni-plugins/" + CNIPluginVer + "/binaries/cni-plugins-linux-amd64-" + CNIPluginVer + ".tgz CONTAINERD_DOWNLOAD_URL_BASE=https://storage.googleapis.com/cri-containerd-release/ CONTAINERD_VERSION=" + DefaultContainerdVersion + " CONTAINER_RUNTIME=docker ETCD_DOWNLOAD_URL=mcr.microsoft.com/oss/etcd-io/ ETCD_VERSION=" + DefaultEtcdVersion + " EXCLUDE_MASTER_FROM_STANDARD_LB=false HYPERKUBE_URL=hyperkube-amd64:v1.16.8 IDENTITY_SYSTEM=azure_ad IS_HOSTED_MASTER=false IS_IPV6_DUALSTACK_FEATURE_ENABLED=false IS_IPV6_ENABLED=false KMS_PROVIDER_VAULT_NAME= KUBELET_PRIVATE_KEY=fakekubeletkey KUBERNETES_VERSION=1.16.8 KUBE_BINARY_URL= LOAD_BALANCER_DISABLE_OUTBOUND_SNAT=false LOAD_BALANCER_SKU=Basic LOCATION=westus MAXIMUM_LOADBALANCER_RULE_COUNT=0 MOBY_VERSION=" + DefaultMobyVersion + " NETWORK_API_VERSION=2018-08-01 NETWORK_MODE= NETWORK_PLUGIN=kubenet NETWORK_POLICY= NETWORK_SECURITY_GROUP=k8s-master-22998975-nsg PRIMARY_AVAILABILITY_SET=agentpool1-availabilitySet-22998975 PRIMARY_SCALE_SET= RESOURCE_GROUP=fakerg ROUTE_TABLE=k8s-master-22998975-routetable SERVICE_PRINCIPAL_CLIENT_ID=fakeclientID SERVICE_PRINCIPAL_CLIENT_SECRET=fakeclientSecret SUBNET=k8s-subnet SUBSCRIPTION_ID=fakesubID TENANT_ID=faketenantID USE_INSTANCE_METADATA=false USE_MANAGED_IDENTITY_EXTENSION=false VIRTUAL_NETWORK=k8s-vnet-22998975 VIRTUAL_NETWORK_RESOURCE_GROUP= VM_TYPE=standard VNET_CNI_PLUGINS_URL=https://kubernetesartifacts.azureedge.net/azure-cni/" + AzureCniPluginVerLinux + "/binaries/azure-vnet-cni-linux-amd64-" + AzureCniPluginVerLinux + ".tgz ",
		},
		{
			name: "With ARM variables",
			cs:   CreateMockContainerService("testcluster", "1.16.8", 1, 3, true),
			input: ProvisionScriptParametersInput{
				Location:             common.WrapAsARMVariable("location"),
				ResourceGroup:        common.WrapAsARMVariable("resourceGroup"),
				TenantID:             common.WrapAsARMVariable("tenantID"),
				SubscriptionID:       common.WrapAsARMVariable("subscriptionId"),
				ClientID:             common.WrapAsARMVariable("servicePrincipalClientId"),
				ClientSecret:         common.WrapAsARMVariable("singleQuote") + common.WrapAsARMVariable("servicePrincipalClientSecret") + common.WrapAsARMVariable("singleQuote"),
				APIServerCertificate: common.WrapAsParameter("apiServerCertificate"),
				KubeletPrivateKey:    common.WrapAsParameter("clientPrivateKey"),
				ClusterKeyVaultName:  common.WrapAsARMVariable("clusterKeyvaultName"),
			},
			expected: "ADMINUSER=azureuser APISERVER_PUBLIC_KEY=',parameters('apiServerCertificate'),' AUTHENTICATION_METHOD=client_secret CLOUDPROVIDER_BACKOFF=false CLOUDPROVIDER_BACKOFF_DURATION=0 CLOUDPROVIDER_BACKOFF_EXPONENT=0 CLOUDPROVIDER_BACKOFF_JITTER=0 CLOUDPROVIDER_BACKOFF_MODE= CLOUDPROVIDER_BACKOFF_RETRIES=0 CLOUDPROVIDER_RATELIMIT=false CLOUDPROVIDER_RATELIMIT_BUCKET=0 CLOUDPROVIDER_RATELIMIT_BUCKET_WRITE=0 CLOUDPROVIDER_RATELIMIT_QPS=0 CLOUDPROVIDER_RATELIMIT_QPS_WRITE=0 CNI_PLUGINS_URL=https://kubernetesartifacts.azureedge.net/cni-plugins/" + CNIPluginVer + "/binaries/cni-plugins-linux-amd64-" + CNIPluginVer + ".tgz CONTAINERD_DOWNLOAD_URL_BASE=https://storage.googleapis.com/cri-containerd-release/ CONTAINERD_VERSION=" + DefaultContainerdVersion + " CONTAINER_RUNTIME=docker ETCD_DOWNLOAD_URL=mcr.microsoft.com/oss/etcd-io/ ETCD_VERSION=" + DefaultEtcdVersion + " EXCLUDE_MASTER_FROM_STANDARD_LB=false HYPERKUBE_URL=hyperkube-amd64:v1.16.8 IDENTITY_SYSTEM=azure_ad IS_HOSTED_MASTER=false IS_IPV6_DUALSTACK_FEATURE_ENABLED=false IS_IPV6_ENABLED=false KMS_PROVIDER_VAULT_NAME=',variables('clusterKeyvaultName'),' KUBELET_PRIVATE_KEY=',parameters('clientPrivateKey'),' KUBERNETES_VERSION=1.16.8 KUBE_BINARY_URL= LOAD_BALANCER_DISABLE_OUTBOUND_SNAT=false LOAD_BALANCER_SKU=Basic LOCATION=',variables('location'),' MAXIMUM_LOADBALANCER_RULE_COUNT=0 MOBY_VERSION=" + DefaultMobyVersion + " NETWORK_API_VERSION=2018-08-01 NETWORK_MODE= NETWORK_PLUGIN=kubenet NETWORK_POLICY= NETWORK_SECURITY_GROUP=k8s-master-22998975-nsg PRIMARY_AVAILABILITY_SET=agentpool1-availabilitySet-22998975 PRIMARY_SCALE_SET= RESOURCE_GROUP=',variables('resourceGroup'),' ROUTE_TABLE=k8s-master-22998975-routetable SERVICE_PRINCIPAL_CLIENT_ID=',variables('servicePrincipalClientId'),' SERVICE_PRINCIPAL_CLIENT_SECRET=',variables('singleQuote'),'',variables('servicePrincipalClientSecret'),'',variables('singleQuote'),' SUBNET=k8s-subnet SUBSCRIPTION_ID=',variables('subscriptionId'),' TENANT_ID=',variables('tenantID'),' USE_INSTANCE_METADATA=false USE_MANAGED_IDENTITY_EXTENSION=false VIRTUAL_NETWORK=k8s-vnet-22998975 VIRTUAL_NETWORK_RESOURCE_GROUP= VM_TYPE=standard VNET_CNI_PLUGINS_URL=https://kubernetesartifacts.azureedge.net/azure-cni/" + AzureCniPluginVerLinux + "/binaries/azure-vnet-cni-linux-amd64-" + AzureCniPluginVerLinux + ".tgz ",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := c.cs.GetProvisionScriptParametersCommon(c.input)
			if actual != c.expected {
				t.Fatalf("expected cs.GetProvisionScriptParametersCommon() to return %s but instead returned %s", c.expected, actual)
			}
		})
	}
}

func TestShouldEnableAzureCloudAddon(t *testing.T) {
	cases := []struct {
		name      string
		addonName string
		p         *Properties
		expected  bool
	}{
		{
			name: "useCloudControllerManager disabled",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(false),
					},
				},
			},
			expected: false,
		},
		{
			name:      "unknown addon",
			addonName: "unknown addon",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
			expected: false,
		},
		{
			name:      "[Linux] azuredisk-csi-driver with K8s 1.13",
			addonName: common.AzureDiskCSIDriverAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
			expected: true,
		},
		{
			name:      "[Linux] azurefile-csi-driver with K8s 1.13",
			addonName: common.AzureFileCSIDriverAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
			expected: true,
		},
		{
			name:      "[Linux] cloud-node-manager with K8s 1.15",
			addonName: common.CloudNodeManagerAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
			expected: false,
		},
		{
			name:      "[Linux] cloud-node-manager with K8s 1.16",
			addonName: common.CloudNodeManagerAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
			expected: true,
		},
		{
			name:      "[Windows] azuredisk-csi-driver with K8s 1.18",
			addonName: common.AzureDiskCSIDriverAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						OSType: Windows,
					},
				},
			},
			expected: true,
		},
		{
			name:      "[Windows] azurefile-csi-driver with K8s 1.18",
			addonName: common.AzureFileCSIDriverAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						OSType: Windows,
					},
				},
			},
			expected: true,
		},
		{
			name:      "[Windows] cloud-node-manager with K8s 1.18",
			addonName: common.CloudNodeManagerAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						OSType: Windows,
					},
				},
			},
			expected: true,
		},
		{
			name:      "[Windows] cloud-node-manager with K8s 1.17",
			addonName: common.CloudNodeManagerAddonName,
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.0",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						OSType: Windows,
					},
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := c.p.ShouldEnableAzureCloudAddon(c.addonName)
			if actual != c.expected {
				t.Fatalf("expected p.ShouldEnableAzureCloudAddon(\"%s\") to return %t but instead returned %t", c.addonName, c.expected, actual)
			}
		})
	}
}
