// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"
)

func TestSchedulerDefaultConfig(t *testing.T) {
	cs := CreateMockContainerService("testcluster", "", 3, 2, false)
	cs.setSchedulerConfig()
	s := cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	for key, val := range staticSchedulerConfig {
		if val != s[key] {
			t.Fatalf("got unexpected kube-scheduler static config value for %s. Expected %s, got %s",
				key, val, s[key])
		}
	}
	for key, val := range defaultSchedulerConfig {
		if val != s[key] {
			t.Fatalf("got unexpected kube-scheduler default config value for %s. Expected %s, got %s",
				key, val, s[key])
		}
	}
}

func TestSchedulerUserConfig(t *testing.T) {
	cs := CreateMockContainerService("testcluster", "", 3, 2, true)
	assignmentMap := map[string]string{
		"--scheduler-name": "my-custom-name",
		"--feature-gates":  "APIListChunking=true,APIResponseCompression=true,Accelerators=true,AdvancedAuditing=true",
	}
	cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig = assignmentMap
	cs.setSchedulerConfig()
	for key, val := range assignmentMap {
		if val != cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig[key] {
			t.Fatalf("got unexpected kube-scheduler config value for %s. Expected %s, got %s",
				key, val, cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig[key])
		}
	}
}

func TestSchedulerStaticConfig(t *testing.T) {
	cs := CreateMockContainerService("testcluster", "", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig = map[string]string{
		"--kubeconfig":   "user-override",
		"--leader-elect": "user-override",
		"--profiling":    "user-override",
	}
	cs.setSchedulerConfig()
	for key, val := range staticSchedulerConfig {
		if val != cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig[key] {
			t.Fatalf("kube-scheduler static config did not override user values for %s. Expected %s, got %s",
				key, val, cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig)
		}
	}
}

func TestSchedulerConfigEnableProfiling(t *testing.T) {
	// Test
	// "schedulerConfig": {
	// 	"--profiling": "true"
	// },
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig = map[string]string{
		"--profiling": "true",
	}
	cs.setSchedulerConfig()
	s := cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	if s["--profiling"] != "true" {
		t.Fatalf("got unexpected '--profiling' Scheduler config value for \"--profiling\": \"true\": %s",
			s["--profiling"])
	}

	// Test default
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setSchedulerConfig()
	s = cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	if s["--profiling"] != DefaultKubernetesSchedulerEnableProfiling {
		t.Fatalf("got unexpected default value for '--profiling' Scheduler config: %s",
			s["--profiling"])
	}
}

func TestSchedulerFeatureGates(t *testing.T) {
	// test defaultTestClusterVer
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setSchedulerConfig()
	s := cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	if s["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Scheduler config value for k8s v%s: %s",
			defaultTestClusterVer, s["--feature-gates"])
	}

	// test 1.19.0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.19.0"
	cs.setSchedulerConfig()
	s = cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	if s["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Scheduler config value for k8s v%s: %s",
			"1.19.0", s["--feature-gates"])
	}

	// test 1.22.0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.22.0"
	cs.setSchedulerConfig()
	s = cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	if s["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Scheduler config value for k8s v%s: %s",
			"1.22.0", s["--feature-gates"])
	}

	// test user-overrides, removal of VolumeSnapshotDataSource for k8s versions >= 1.22
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.22.0"
	cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig = make(map[string]string)
	s = cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	s["--feature-gates"] = "VolumeSnapshotDataSource=true"
	cs.setSchedulerConfig()
	if s["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Scheduler config value for \"--feature-gates\": \"VolumeSnapshotDataSource=true\": %s for k8s v%s",
			s["--feature-gates"], "1.22.0")
	}

	// test user-overrides, no removal of VolumeSnapshotDataSource for k8s versions < 1.22
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.19.0"
	cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig = make(map[string]string)
	s = cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig
	s["--feature-gates"] = "VolumeSnapshotDataSource=true"
	cs.setSchedulerConfig()
	if s["--feature-gates"] != "VolumeSnapshotDataSource=true" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for \"--feature-gates\": \"VolumeSnapshotDataSource=true\": %s for k8s v%s",
			s["--feature-gates"], "1.19.0")
	}
}
