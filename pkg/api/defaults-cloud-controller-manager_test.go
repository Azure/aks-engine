// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
)

func TestCloudControllerManagerConfig(t *testing.T) {
	k8sVersion := common.RationalizeReleaseAndVersion(Kubernetes, "", "", false, false, true)
	cs := CreateMockContainerService("testcluster", k8sVersion, 3, 2, false)
	cs.setCloudControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	if cm["--controllers"] != "*,-cloud-node" {
		t.Fatalf("got unexpected '--controllers' Cloud Controller Manager config value for Kubernetes %s: %s",
			k8sVersion, cm["--controllers"])
	}
}

func TestCloudControllerManagerFeatureGates(t *testing.T) {
	// test defaultTestClusterVer
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setCloudControllerManagerConfig()
	ccm := cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	if ccm["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Cloud Controller Manager config value for k8s v%s: %s",
			defaultTestClusterVer, ccm["--feature-gates"])
	}

	// test 1.19.0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.19.0"
	cs.setCloudControllerManagerConfig()
	ccm = cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	if ccm["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Cloud Controller Manager config value for k8s v%s: %s",
			"1.19.0", ccm["--feature-gates"])
	}

	// test 1.22.0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.22.0"
	cs.setCloudControllerManagerConfig()
	ccm = cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	if ccm["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' Cloud Controller Manager config value for k8s v%s: %s",
			"1.22.0", ccm["--feature-gates"])
	}

	// test user-overrides, removal of VolumeSnapshotDataSource for k8s versions >= 1.22
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.22.0"
	cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig = make(map[string]string)
	ccm = cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	ccm["--feature-gates"] = "VolumeSnapshotDataSource=true"
	cs.setCloudControllerManagerConfig()
	if ccm["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for \"--feature-gates\": \"VolumeSnapshotDataSource=true\": %s for k8s v%s",
			ccm["--feature-gates"], "1.22.0")
	}

	// test user-overrides, no removal of VolumeSnapshotDataSource for k8s versions < 1.22
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.19.0"
	cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig = make(map[string]string)
	ccm = cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	ccm["--feature-gates"] = "VolumeSnapshotDataSource=true"
	cs.setCloudControllerManagerConfig()
	if ccm["--feature-gates"] != "VolumeSnapshotDataSource=true" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for \"--feature-gates\": \"VolumeSnapshotDataSource=true\": %s for k8s v%s",
			ccm["--feature-gates"], "1.19.0")
	}
}
