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
