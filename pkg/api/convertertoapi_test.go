// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"k8s.io/apimachinery/pkg/api/equality"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/v20170701"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
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
	cs := ConvertV20170701ContainerService(v20170701cs)
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(false) {
		t.Fatalf("incorrect OrchestratorVersion '%s'", cs.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	v20170701cs = &v20170701.ContainerService{
		Properties: &v20170701.Properties{
			OrchestratorProfile: &v20170701.OrchestratorProfile{
				OrchestratorType:    v20170701.Kubernetes,
				OrchestratorVersion: "1.7.15",
			},
		},
	}
	cs = ConvertV20170701ContainerService(v20170701cs)
	if cs.Properties.OrchestratorProfile.OrchestratorVersion != "1.7.15" {
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
