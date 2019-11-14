// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"
)

func TestCloudControllerManagerConfig(t *testing.T) {
	k8sVersion := "1.16.1"
	cs := CreateMockContainerService("testcluster", k8sVersion, 3, 2, false)
	cs.setCloudControllerManagerConfig()
	cm := cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	if cm["--controllers"] != "*" {
		t.Fatalf("got unexpected '--controllers' Cloud Controller Manager config value for Kubernetes %s: %s",
			k8sVersion, cm["--controllers"])
	}

	k8sVersion = "1.15.4"
	cs = CreateMockContainerService("testcluster", k8sVersion, 3, 2, false)
	cs.setCloudControllerManagerConfig()
	cm = cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig
	if val, ok := cm["--controllers"]; ok {
		t.Fatalf("got unexpected '--controllers' Cloud Controller Manager config value for Kubernetes %s: %s",
			k8sVersion, val)
	}
}
