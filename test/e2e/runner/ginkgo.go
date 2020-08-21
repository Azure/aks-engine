//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/Azure/aks-engine/test/e2e/metrics"
	"github.com/kelseyhightower/envconfig"
)

// Ginkgo contains all of the information needed to run the ginkgo suite of tests
type Ginkgo struct {
	GinkgoNodes string `envconfig:"GINKGO_NODES" default:"6"`
	Config      *config.Config
	Point       *metrics.Point
}

// BuildGinkgoRunner creates a new Ginkgo object
func BuildGinkgoRunner(cfg *config.Config, pt *metrics.Point) (*Ginkgo, error) {
	g := new(Ginkgo)
	if err := envconfig.Process("ginkgo", g); err != nil {
		return nil, err
	}
	g.Config = cfg
	g.Point = pt
	return g, nil
}

// Run will execute an orchestrator suite of tests
func (g *Ginkgo) Run() error {
	g.Point.SetTestStart()
	// use the test bin rather than compile the directory b/c the compile will happen in a sub dir which is another module
	testFile := fmt.Sprintf("test/e2e/%s/%s.test", g.Config.Orchestrator, g.Config.Orchestrator)

	args := []string{"-slowSpecThreshold", "180", "-r", "-v"}
	if g.Config.GinkgoFailFast {
		args = append(args, "--failFast")
	}
	if g.Config.GinkgoFocus != "" {
		args = append(args, "--focus")
		args = append(args, g.Config.GinkgoFocus)
	}
	if g.Config.GinkgoSkip != "" {
		args = append(args, "--skip")
		args = append(args, g.Config.GinkgoSkip)
	}
	args = append(args, testFile)
	var cmd = exec.Command("ginkgo", args...)

	util.PrintCommand(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		g.Point.RecordTestError()
		log.Printf("Error while trying to start ginkgo:%s\n", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		g.Point.RecordTestError()
		if g.Config.IsKubernetes() {
			kubectl := exec.Command("k", "get", "all", "--all-namespaces", "-o", "wide")
			util.PrintCommand(kubectl)
			kubectl.CombinedOutput()
			kubectl = exec.Command("k", "get", "nodes", "-o", "wide")
			util.PrintCommand(kubectl)
			kubectl.CombinedOutput()
		}
		return err
	}
	g.Point.RecordTestSuccess()
	return nil
}
