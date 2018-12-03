// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package dcos

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/engine"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	cfg     config.Config
	eng     engine.Engine
	err     error
	cluster *Cluster
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
	cs, err := engine.ParseInput(engCfg.ClusterDefinitionTemplate)
	Expect(err).NotTo(HaveOccurred())
	eng = engine.Engine{
		Config:            engCfg,
		ClusterDefinition: cs,
	}

	cluster, err = NewCluster(&cfg, &eng)
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("Azure Container Cluster using the DCOS Orchestrator", func() {
	Context("regardless of agent pool type", func() {

		It("should have have the appropriate node count", func() {
			count, err := cluster.NodeCount()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(eng.NodeCount()))
		})

		It("should be running the expected version", func() {
			version, err := cluster.Version()
			Expect(err).NotTo(HaveOccurred())

			expectedVersion := common.RationalizeReleaseAndVersion(
				eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorType,
				eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorRelease,
				eng.ClusterDefinition.Properties.OrchestratorProfile.OrchestratorVersion,
				false,
				false)
			Expect(version).To(Equal(expectedVersion))
		})

		It("should be able to install marathon", func() {
			err = cluster.InstallMarathonLB()
			if err != nil {
				log.Printf("Error while installing Marathon LB: %s\n", err.Error())
			}
			Expect(err).NotTo(HaveOccurred())

			marathonPath := filepath.Join(cfg.CurrentWorkingDir, "/test/e2e/dcos/marathon.json")
			port, err := cluster.InstallMarathonApp(marathonPath, 5*time.Second, cfg.Timeout)
			if err != nil {
				log.Printf("Error while installing Marathon app: %s\n", err.Error())
			}
			Expect(err).NotTo(HaveOccurred())

			// Need to have a wait for ready check here
			cmd := fmt.Sprintf("curl -sI http://marathon-lb.marathon.mesos:%v/", port)
			out, err := cluster.Connection.ExecuteWithRetries(cmd, 5*time.Second, cfg.Timeout)
			if err != nil {
				log.Printf("Was not able to run %s, err: %s\n", cmd, err.Error())
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(out).To(MatchRegexp("^HTTP/1.1 200 OK"))
		})

	})
})
