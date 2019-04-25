// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// Package test provides a method to export Ginkgo test results to custom unit test reporters.
package test

import (
	"fmt"
	"testing"

	"path/filepath"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
)

var (
	// JUnitOutDir grabs the root of the git path from Makefile
	JUnitOutDir = ""
)

// RunSpecsWithReporters bootstraps Ginkgo/Gomega tests to function and results go to the /test/junit directory and log output
func RunSpecsWithReporters(t *testing.T, junitprefix string, suitename string) {

	gomega.RegisterFailHandler(ginkgo.Fail)
	if JUnitOutDir == "" {
		ginkgo.RunSpecs(t, suitename)
		return
	}
	junitReporter := reporters.NewJUnitReporter(filepath.Join(JUnitOutDir, fmt.Sprintf("%s-junit.xml", junitprefix)))
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, suitename, []ginkgo.Reporter{junitReporter})
}
