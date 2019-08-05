// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"fmt"
	"testing"
	"time"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	. "github.com/Azure/aks-engine/pkg/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

func TestGetNodes(t *testing.T) {
	RunSpecsWithReporters(t, "Kubernetes API Operations conveniences library", "Server Suite")
}

var _ = Describe("GetNodes tests", func() {
	It("listNodes should return a result set with nodes", func() {
		mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Minute * 1

		result := listNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout)

		Expect(result.err).To(BeNil())
		Expect(result.nodes).To(HaveLen(2))
		Expect(result.nodes[0].Name).To(Equal("k8s-master-1234"))
		Expect(result.nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeReady))
		Expect(result.nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(result.nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.10"))
		Expect(result.nodes[1].Name).To(Equal("k8s-agentpool3-1234"))
		Expect(result.nodes[1].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
		Expect(result.nodes[1].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(result.nodes[1].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))
	})

	It("listNodes should error when the k8s client getter fails", func() {
		mockClient := armhelpers.MockAKSEngineClient{
			FailGetKubernetesClient: true,
			MockKubernetesClient:    &armhelpers.MockKubernetesClient{}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Minute * 1

		result := listNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout)

		Expect(result.err).NotTo(BeNil())
		Expect(result.err.Error()).To(Equal("GetKubernetesClient failed"))
		Expect(result.nodes).To(HaveLen(0))
	})

	It("listNodes should error when the k8s API fails to list nodes", func() {
		mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{
			FailListNodes: true,
		}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Minute * 1

		result := listNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout)

		Expect(result.err).NotTo(BeNil())
		Expect(result.err.Error()).To(Equal("ListNodes failed"))
		Expect(result.nodes).To(HaveLen(0))
	})

	It("GetNodes should return nodes", func() {
		mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Minute * 1

		nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", -1)

		Expect(err).To(BeNil())
		Expect(nodes).To(HaveLen(2))
		Expect(nodes[0].Name).To(Equal("k8s-master-1234"))
		Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeReady))
		Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.10"))
		Expect(nodes[1].Name).To(Equal("k8s-agentpool3-1234"))
		Expect(nodes[1].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
		Expect(nodes[1].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(nodes[1].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))
	})

	It("GetNodes should only return nodes in a pool when a pool string is specified", func() {
		mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Minute * 1

		nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "agentpool3", -1)

		Expect(err).To(BeNil())
		Expect(nodes).To(HaveLen(1))
		Expect(nodes[0].Name).To(Equal("k8s-agentpool3-1234"))
		Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
		Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))

		nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "nonexistent", -1)

		Expect(err).To(BeNil())
		Expect(nodes).To(HaveLen(0))
	})

	It("GetNodes should respect the waitForNumNodes arg", func() {
		mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Second * 1

		nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", 2)

		Expect(err).To(BeNil())
		Expect(nodes).To(HaveLen(2))
		Expect(nodes[0].Name).To(Equal("k8s-master-1234"))
		Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeReady))
		Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.10"))
		Expect(nodes[1].Name).To(Equal("k8s-agentpool3-1234"))
		Expect(nodes[1].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
		Expect(nodes[1].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(nodes[1].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))

		// waiting for more nodes than the API returns should timeout
		nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", 3)
		var mostRecentGetNodesErr error

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
		Expect(nodes).To(BeNil())

		// waiting for fewer nodes than the API returns should timeout
		nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", 1)

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
		Expect(nodes).To(BeNil())

		// filtering by pool name and and the waiting for the expected node count
		nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "agentpool3", 1)

		Expect(err).To(BeNil())
		Expect(nodes).To(HaveLen(1))
		Expect(nodes[0].Name).To(Equal("k8s-agentpool3-1234"))
		Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
		Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
		Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))
	})

	It("GetNodes should return a meaningful timeout error when the k8s API fails to list nodes", func() {
		mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{
			FailListNodes: true,
		}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Second * 1 // set the timeout value high enough to allow for a single attempt

		nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", -1)
		mostRecentGetNodesErr := errors.New("ListNodes failed")

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
		Expect(nodes).To(BeNil())
	})

	It("GetNodes should return a vanilla timeout error if timeout occurs before a single request", func() {
		mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		logger := log.NewEntry(log.New())
		apiserverURL := "https://apiserver"
		kubeconfig := "kubeconfig"
		timeout := time.Second * 0 // by setting the timeout to 0 we time out immediately

		nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", -1)
		var mostRecentGetNodesErr error

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
		Expect(nodes).To(BeNil())
	})
})

func ExamplePrintNodes() {
	var nodes []v1.Node
	node := v1.Node{}
	node.Name = "k8s-master-1234"
	node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionTrue})
	node.Status.NodeInfo.KubeletVersion = "1.10.0"
	node.Status.NodeInfo.OSImage = "my-os"
	node.Status.NodeInfo.KernelVersion = "3.1.4"
	nodes = append(nodes, node)
	PrintNodes(nodes)
	// Output: NODE               STATUS    VERSION    OS       KERNEL
	// k8s-master-1234    Ready     1.10.0     my-os    3.1.4
}
