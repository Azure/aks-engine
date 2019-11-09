// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"

	"github.com/Azure/aks-engine/pkg/armhelpers"
)

func TestGetNodes_ShouldReturnAResultSetWithNodes(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Minute * 1

	result := listNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout)
	g := NewGomegaWithT(t)
	g.Expect(result.err).To(BeNil())
	g.Expect(result.nodes).To(HaveLen(2))
	g.Expect(result.nodes[0].Name).To(Equal("k8s-master-1234"))
	g.Expect(result.nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeReady))
	g.Expect(result.nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(result.nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.10"))
	g.Expect(result.nodes[1].Name).To(Equal("k8s-agentpool3-1234"))
	g.Expect(result.nodes[1].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
	g.Expect(result.nodes[1].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(result.nodes[1].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))
}

func TestGetNodes_ShouldErrorWhenK8sClientGetterFails(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{
		FailGetKubernetesClient: true,
		MockKubernetesClient:    &armhelpers.MockKubernetesClient{}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Minute * 1

	result := listNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout)

	g := NewGomegaWithT(t)
	g.Expect(result.err).NotTo(BeNil())
	g.Expect(result.err.Error()).To(Equal("GetKubernetesClient failed"))
	g.Expect(result.nodes).To(HaveLen(0))
}

func TestGetNodes_ShouldErrorWhenTheK8sAPIFailsToListNodes(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{
		FailListNodes: true,
	}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Minute * 1

	result := listNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout)

	g := NewGomegaWithT(t)
	g.Expect(result.err).NotTo(BeNil())
	g.Expect(result.err.Error()).To(Equal("ListNodes failed"))
	g.Expect(result.nodes).To(HaveLen(0))
}

func TestGetNodes_ShouldReturnNodes(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Minute * 1

	nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", -1)

	g := NewGomegaWithT(t)
	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(2))
	g.Expect(nodes[0].Name).To(Equal("k8s-master-1234"))
	g.Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeReady))
	g.Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.10"))
	g.Expect(nodes[1].Name).To(Equal("k8s-agentpool3-1234"))
	g.Expect(nodes[1].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
	g.Expect(nodes[1].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(nodes[1].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))
}

func TestGetNodes_ShouldReturnNodesInAPoolWhenAPoolStringIsSpecified(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Minute * 1

	nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "agentpool3", -1)

	g := NewGomegaWithT(t)
	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(1))
	g.Expect(nodes[0].Name).To(Equal("k8s-agentpool3-1234"))
	g.Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
	g.Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))

	nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "nonexistent", -1)

	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(0))
}

func TestGetNodes_ShouldRespectTheWaitForNumNodesArg(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Second * 1

	nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", 2)

	g := NewGomegaWithT(t)
	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(2))
	g.Expect(nodes[0].Name).To(Equal("k8s-master-1234"))
	g.Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeReady))
	g.Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.10"))
	g.Expect(nodes[1].Name).To(Equal("k8s-agentpool3-1234"))
	g.Expect(nodes[1].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
	g.Expect(nodes[1].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(nodes[1].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))

	// waiting for more nodes than the API returns should timeout
	nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", 3)
	var mostRecentGetNodesErr error

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
	g.Expect(nodes).To(BeNil())

	// waiting for fewer nodes than the API returns should timeout
	nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", 1)

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
	g.Expect(nodes).To(BeNil())

	// filtering by pool name and and the waiting for the expected node count
	nodes, err = GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "agentpool3", 1)

	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(1))
	g.Expect(nodes[0].Name).To(Equal("k8s-agentpool3-1234"))
	g.Expect(nodes[0].Status.Conditions[0].Type).To(Equal(v1.NodeOutOfDisk))
	g.Expect(nodes[0].Status.Conditions[0].Status).To(Equal(v1.ConditionTrue))
	g.Expect(nodes[0].Status.NodeInfo.KubeletVersion).To(Equal("1.9.9"))
}

func TestGetNodes_ShouldReturnAMeaningfulTimeoutErrorWhenK8sAPIFailsToListNodes(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{
		FailListNodes: true,
	}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Second * 1 // set the timeout value high enough to allow for a single attempt

	nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", -1)
	mostRecentGetNodesErr := errors.New("ListNodes failed")
	g := NewGomegaWithT(t)
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
	g.Expect(nodes).To(BeNil())
}

func TestGetNodes_ShouldReturnAVanillaTimeoutErrorIfOccursBeforeASingleRequest(t *testing.T) {
	t.Parallel()
	mockClient := armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	logger := log.NewEntry(log.New())
	apiserverURL := "https://apiserver"
	kubeconfig := "kubeconfig"
	timeout := time.Second * 0 // by setting the timeout to 0 we time out immediately

	nodes, err := GetNodes(&mockClient, logger, apiserverURL, kubeconfig, timeout, "", -1)
	var mostRecentGetNodesErr error
	g := NewGomegaWithT(t)
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal(fmt.Sprintf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)))
	g.Expect(nodes).To(BeNil())
}

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
