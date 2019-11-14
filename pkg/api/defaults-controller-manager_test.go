// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
)

func TestControllerManagerConfigEnableRbac(t *testing.T) {
	// Test EnableRbac = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(true)
	cs.setControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--use-service-account-credentials"] != "true" {
		t.Fatalf("got unexpected '--use-service-account-credentials' Controller Manager config value for EnableRbac=true: %s",
			cm["--use-service-account-credentials"])
	}

	// Test default
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(false)
	cs.setControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--use-service-account-credentials"] != DefaultKubernetesCtrlMgrUseSvcAccountCreds {
		t.Fatalf("got unexpected '--use-service-account-credentials' Controller Manager config value for EnableRbac=false: %s",
			cm["--use-service-account-credentials"])
	}
}

func TestControllerManagerConfigCloudProvider(t *testing.T) {
	// Test UseCloudControllerManager = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = to.BoolPtr(true)
	cs.setControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--cloud-provider"] != "external" {
		t.Fatalf("got unexpected '--cloud-provider' Controller Manager config value for UseCloudControllerManager=true: %s",
			cm["--cloud-provider"])
	}

	// Test UseCloudControllerManager = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = to.BoolPtr(false)
	cs.setControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--cloud-provider"] != "azure" {
		t.Fatalf("got unexpected '--cloud-provider' Controller Manager config value for UseCloudControllerManager=false: %s",
			cm["--cloud-provider"])
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

	// test 1.14
	cs = CreateMockContainerService("testcluster", "1.14.1", 3, 2, false)
	cs.setControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--feature-gates"] != "LocalStorageCapacityIsolation=true,ServiceNodeExclusion=true" {
		t.Fatalf("got unexpected '--feature-gates' Controller Manager config value for \"--feature-gates\": \"LocalStorageCapacityIsolation=true,ServiceNodeExclusion=true\": %s",
			cm["--feature-gates"])
	}

	// test user-overrides
	cs = CreateMockContainerService("testcluster", "1.14.1", 3, 2, false)
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	cm["--feature-gates"] = "TaintBasedEvictions=true"
	cs.setControllerManagerConfig()
	if cm["--feature-gates"] != "LocalStorageCapacityIsolation=true,ServiceNodeExclusion=true,TaintBasedEvictions=true" {
		t.Fatalf("got unexpected '--feature-gates' Controller Manager config value for \"--feature-gates\": \"LocalStorageCapacityIsolation=true,ServiceNodeExclusion=true\": %s",
			cm["--feature-gates"])
	}
}

func TestControllerManagerConfigHostedMasterProfile(t *testing.T) {
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.MasterProfile = nil
	cs.Properties.HostedMasterProfile = &HostedMasterProfile{
		DNSPrefix: "foodns",
	}
	cs.setControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig
	if cm["--cluster-name"] != "foodns" {
		t.Fatalf("expected controller-manager to have cluster-name foodns when using HostedMasterProfile")
	}
}
