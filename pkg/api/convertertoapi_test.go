package api

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/v20170701"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/davecgh/go-spew/spew"
	"k8s.io/apimachinery/pkg/api/equality"
)

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

func TestConvertVLabsOrchestratorProfile(t *testing.T) {
	t.Skip("Should be refactored to expect Kubernetes orchestrator.")
	tests := map[string]struct {
		props  *vlabs.Properties
		expect *OrchestratorProfile
	}{
		"nilOpenShiftConfig": {
			props: &vlabs.Properties{
				OrchestratorProfile: &vlabs.OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
			},
			expect: &OrchestratorProfile{
				OrchestratorType:    Kubernetes,
				OrchestratorVersion: common.KubernetesDefaultRelease,
			},
		},
	}

	for name, test := range tests {
		t.Logf("running scenario %q", name)
		actual := &OrchestratorProfile{}
		convertVLabsOrchestratorProfile(test.props, actual, false)
		if !equality.Semantic.DeepEqual(test.expect, actual) {
			t.Errorf(spew.Sprintf("Expected:\n%+v\nGot:\n%+v", test.expect, actual))
		}
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
