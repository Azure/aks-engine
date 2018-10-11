package api

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/helpers"
)

func TestControllerManagerConfigEnableRbac(t *testing.T) {
	// Test EnableRbac = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = helpers.PointerToBool(true)
	cs.setControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--use-service-account-credentials"] != "true" {
		t.Fatalf("got unexpected '--use-service-account-credentials' Controller Manager config value for EnableRbac=true: %s",
			cm["--use-service-account-credentials"])
	}

	// Test default
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = helpers.PointerToBool(false)
	cs.setControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--use-service-account-credentials"] != DefaultKubernetesCtrlMgrUseSvcAccountCreds {
		t.Fatalf("got unexpected '--use-service-account-credentials' Controller Manager config value for EnableRbac=false: %s",
			cm["--use-service-account-credentials"])
	}

}
func TestControllerManagerConfigEnableProfiling(t *testing.T) {
	// Test
	// "controllerManagerConfig": {
	// 	"--profiling": "true"
	// },
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig = map[string]string{
		"--profiling": "true",
	}
	cs.setControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--profiling"] != "true" {
		t.Fatalf("got unexpected '--profiling' Controller Manager config value for \"--profiling\": \"true\": %s",
			cm["--profiling"])
	}

	// Test default
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--profiling"] != DefaultKubernetesCtrMgrEnableProfiling {
		t.Fatalf("got unexpected default value for '--profiling' Controller Manager config: %s",
			cm["--profiling"])
	}
}

func TestControllerManagerConfigDefaultFeatureGates(t *testing.T) {
	// test defaultTestClusterVer
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Controller Manager config value for \"--feature-gates\": \"\": %s",
			cm["--feature-gates"])
	}

	// test 1.9.0
	cs = CreateMockContainerService("testcluster", "1.9.0", 3, 2, false)
	cs.setControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--feature-gates"] != "ServiceNodeExclusion=true" {
		t.Fatalf("got unexpected '--feature-gates' Controller Manager config value for \"--feature-gates\": \"ServiceNodeExclusion=true\": %s",
			cm["--feature-gates"])
	}

	// test 1.10.0
	cs = CreateMockContainerService("testcluster", "1.10.0", 3, 2, false)
	cs.setControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--feature-gates"] != "LocalStorageCapacityIsolation=true,ServiceNodeExclusion=true" {
		t.Fatalf("got unexpected '--feature-gates' Controller Manager config value for \"--feature-gates\": \"LocalStorageCapacityIsolation=true,ServiceNodeExclusion=true\": %s",
			cm["--feature-gates"])
	}
}
