// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewRotateCertsCmd(t *testing.T) {
	output := newRotateCertsCmd()
	if output.Use != rotateCertsName || output.Short != rotateCertsShortDescription || output.Long != rotateCertsLongDescription {
		t.Fatalf("rotate-certs command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, rotateCertsName, output.Short, rotateCertsShortDescription, output.Long, rotateCertsLongDescription)
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
	rcc := rotateCertsCmd{
		authProvider:     &authArgs{},
		client:           mockClient,
		containerService: api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
	}
	err := rcc.getClusterNodes()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to get cluster nodes"))

	mockClient.MockKubernetesClient.FailListNodes = false
	err = rcc.getClusterNodes()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(len(rcc.masterNodes)).To(Equal(1))
	g.Expect(len(rcc.agentNodes)).To(Equal(1))
}

func TestDeleteAllPods(t *testing.T) {
	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.MockKubernetesClient.FailListPods = true
	rcc := rotateCertsCmd{
		authProvider:     &authArgs{},
		client:           mockClient,
		containerService: api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
	}
	err := rcc.deleteAllPods()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to get pods"))

	mockClient.MockKubernetesClient.FailListPods = false
	mockClient.MockKubernetesClient.FailDeletePod = true
	mockClient.MockKubernetesClient.PodsList = &v1.PodList{
		Items: []v1.Pod{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "kube-dns",
					Namespace: "kube-system",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-pod",
					Namespace: "kube-system",
				},
			},
		},
	}
	err = rcc.deleteAllPods()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to delete pod"))

	mockClient.MockKubernetesClient.FailDeletePod = false
	err = rcc.deleteAllPods()
	g.Expect(err).NotTo(HaveOccurred())
}

func TestRebootAllNodes(t *testing.T) {
	ctx := context.Background()
	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.FailListVirtualMachines = true
	mockClient.FailListVirtualMachineScaleSets = false
	mockClient.FakeListVirtualMachineScaleSetsResult = func() []compute.VirtualMachineScaleSet {
		scalesetName := "scalesetName"
		sku := compute.Sku{}
		location := "eastus"
		return []compute.VirtualMachineScaleSet{
			{
				Name:     &scalesetName,
				Sku:      &sku,
				Location: &location,
			},
		}
	}
	rcc := rotateCertsCmd{
		authProvider:      &authArgs{},
		client:            mockClient,
		containerService:  api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
		resourceGroupName: "test-rg",
	}
	err := rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to list Virtual Machines in resource group test-rg"))

	mockClient.FailListVirtualMachines = false
	mockClient.FailListVirtualMachineScaleSets = true
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to list Virtual Machine Scale Sets in resource group test-rg"))

	mockClient.FailListVirtualMachines = false
	mockClient.FailListVirtualMachineScaleSets = false
	mockClient.FailRestartVirtualMachine = true
	mockClient.FailRestartVirtualMachineScaleSets = false
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to restart Virtual Machine"))

	mockClient.FailRestartVirtualMachine = false
	mockClient.FailRestartVirtualMachineScaleSets = true
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to restart Virtual Machine Scale Sets"))

	mockClient.FailRestartVirtualMachine = false
	mockClient.FailRestartVirtualMachineScaleSets = false
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).NotTo(HaveOccurred())
}

func TestDeleteServiceAccounts(t *testing.T) {
	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.MockKubernetesClient.FailListServiceAccounts = true
	mockClient.MockKubernetesClient.ServiceAccountList = &v1.ServiceAccountList{
		Items: []v1.ServiceAccount{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "kube-dns",
					Namespace: "kube-system",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-sa",
					Namespace: "kube-system",
				},
			},
		},
	}
	rcc := rotateCertsCmd{
		authProvider:     &authArgs{},
		client:           mockClient,
		containerService: api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
	}
	err := rcc.deleteServiceAccounts()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to get cluster service accounts in namespace kube-system"))

	mockClient.MockKubernetesClient.FailListServiceAccounts = false
	mockClient.MockKubernetesClient.FailDeleteServiceAccount = true
	err = rcc.deleteServiceAccounts()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to delete service account kube-dns"))

	mockClient.MockKubernetesClient.FailDeleteServiceAccount = false
	err = rcc.deleteServiceAccounts()
	g.Expect(err).NotTo(HaveOccurred())
}
