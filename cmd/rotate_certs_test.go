// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
)

func TestNewRotateCertsCmd(t *testing.T) {
	output := newRotateCertsCmd()
	if output.Use != rotateCertsName || output.Short != rotateCertsShortDescription || output.Long != rotateCertsLongDescription {
		t.Fatalf("rotate-certs command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, rotateCertsName, output.Short, rotateCertsShortDescription, output.Long, otateCertsLongDescription)
	}

	expectedFlags := []string{"location", "resource-group", "master-FQDN", "apimodel", "ssh"}
	for _, f := range expectedFlags {
		if output.Flags().Lookup(f) == nil {
			t.Fatalf("rotate-certs command should have flag %s", f)
		}
	}
}

func TestGetClusterNodes(t *testing.T) {
	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.MockKubernetesClient.FailListNodes = true
	rcc := newRotateCertsCmd()
	rcc.containerService = api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	err := rcc.getClusterNodes()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to get cluster nodes"))

	mockClient.MockKubernetesClient.FailListNodes = false
	err = rcc.getClusterNodes()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(len(rcc.masterNodes)).To(Equal(1))
	g.Expect(len(rcc.agentNodes)).To(Equal(1))
}
