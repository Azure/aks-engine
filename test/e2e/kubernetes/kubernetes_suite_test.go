// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Kubernetes Suite", []Reporter{junitReporter})
}
