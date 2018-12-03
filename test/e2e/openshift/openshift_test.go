// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package openshift

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/engine"
	knode "github.com/Azure/aks-engine/test/e2e/kubernetes/node"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/pod"
	"github.com/Azure/aks-engine/test/e2e/openshift/node"
	"github.com/Azure/aks-engine/test/e2e/openshift/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	cfg config.Config
	eng engine.Engine
)

var _ = BeforeSuite(func() {
	cwd, _ := os.Getwd()
	rootPath := filepath.Join(cwd, "../../..") // The current working dir of these tests is down a few levels from the root of the project. We should traverse up that path so we can find the _output dir
	c, err := config.ParseConfig()
	c.CurrentWorkingDir = rootPath
	Expect(err).NotTo(HaveOccurred())
	cfg = *c // We have to do this because golang anon functions and scoping and stuff

	engCfg, err := engine.ParseConfig(c.CurrentWorkingDir, c.ClusterDefinition, c.Name)
	Expect(err).NotTo(HaveOccurred())
	csInput, err := engine.ParseInput(engCfg.ClusterDefinitionTemplate)
	Expect(err).NotTo(HaveOccurred())
	csGenerated, err := engine.ParseOutput(engCfg.GeneratedDefinitionPath + "/apimodel.json")
	Expect(err).NotTo(HaveOccurred())
	eng = engine.Engine{
		Config:             engCfg,
		ClusterDefinition:  csInput,
		ExpandedDefinition: csGenerated,
	}
})

var _ = Describe("Azure Container Cluster using the OpenShift Orchestrator", func() {

	It("should have bootstrap autoapprover running", func() {
		running, err := pod.WaitOnReady("bootstrap-autoapprover", "openshift-infra", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should have have the appropriate node count", func() {
		ready := knode.WaitOnReady(eng.NodeCount(), 10*time.Second, cfg.Timeout)
		Expect(ready).To(Equal(true))
	})

	It("should label nodes correctly", func() {
		labels := map[string]map[string]string{
			"master": {
				"node-role.kubernetes.io/master": "true",
				"openshift-infra":                "apiserver",
			},
			"compute": {
				"node-role.kubernetes.io/compute": "true",
				"region":                          "primary",
			},
			"infra": {
				"region": "infra",
			},
		}
		list, err := knode.Get()
		Expect(err).NotTo(HaveOccurred())

		for _, node := range list.Nodes {
			kind := strings.Split(node.Metadata.Name, "-")[1]
			Expect(labels).To(HaveKey(kind))
			for k, v := range labels[kind] {
				Expect(node.Metadata.Labels).To(HaveKeyWithValue(k, v))
			}
		}
	})

	It("should be running the expected version", func() {
		version, err := node.Version()
		Expect(err).NotTo(HaveOccurred())
		// normalize patch version to zero so we can support testing
		// across centos and rhel deployments where patch versions diverge.
		version = strings.Join(append(strings.Split(version, ".")[:2], "0"), ".")

		var expectedVersion string
		if eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorRelease != "" ||
			eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorVersion != "" {
			expectedVersion = common.RationalizeReleaseAndVersion(
				common.OpenShift,
				eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorRelease,
				eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorVersion,
				false,
				false)
		} else {
			expectedVersion = common.RationalizeReleaseAndVersion(
				common.OpenShift,
				eng.Config.OrchestratorRelease,
				eng.Config.OrchestratorVersion,
				false,
				false)
		}
		expectedVersionRationalized := strings.Split(expectedVersion, "-")[0] // to account for -alpha and -beta suffixes

		// skip unstable test as the version will constantly be changing
		if expectedVersionRationalized != "unstable" {
			Expect(version).To(Equal("v" + expectedVersionRationalized))
		}
	})

	It("should have router running", func() {
		running, err := pod.WaitOnReady("router", "default", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should have docker-registry running", func() {
		running, err := pod.WaitOnReady("docker-registry", "default", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should have registry-console running", func() {
		running, err := pod.WaitOnReady("registry-console", "default", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should deploy a sample app and access it via a route", func() {
		err := util.ApplyFromTemplate("nginx-example", "openshift", "default")
		Expect(err).NotTo(HaveOccurred())
		Expect(util.WaitForDeploymentConfig("nginx-example", "default")).NotTo(HaveOccurred())
		host, err := util.GetHost("nginx-example", "default")
		Expect(err).NotTo(HaveOccurred())
		Expect(util.TestHost(host, 10, 200*time.Millisecond)).NotTo(HaveOccurred())
	})

	It("should have the openshift webconsole running", func() {
		running, err := pod.WaitOnReady("webconsole", "openshift-web-console", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should have prometheus running", func() {
		running, err := pod.WaitOnReady("prometheus", "openshift-metrics", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should have service catalog apiserver running", func() {
		running, err := pod.WaitOnReady("apiserver", "kube-service-catalog", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should have service catalog controller-manager running", func() {
		running, err := pod.WaitOnReady("controller-manager", "kube-service-catalog", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})

	It("should have template service broker running", func() {
		running, err := pod.WaitOnReady("asb", "openshift-ansible-service-broker", 3, 30*time.Second, cfg.Timeout)
		Expect(err).NotTo(HaveOccurred())
		Expect(running).To(Equal(true))
	})
})
