// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestOrchestratorsCmd_ShouldCreate(t *testing.T) {
	command := newOrchestratorsCmd()

	g := NewGomegaWithT(t)
	g.Expect(command.Use).Should(Equal(orchestratorsName))
	g.Expect(command.Short).Should(Equal(orchestratorsShortDescription))
	g.Expect(command.Long).Should(Equal(orchestratorsLongDescription))
	g.Expect(command.Flags().Lookup("orchestrator")).NotTo(BeNil())
	g.Expect(command.Flags().Lookup("version")).NotTo(BeNil())

	command.SetArgs([]string{})
	err := command.Execute()
	g.Expect(err).NotTo(HaveOccurred())
}

func TestOrchestratorsCmd_UnsupportedOrchestrator(t *testing.T) {
	t.Parallel()

	command := &getVersionsCmd{
		orchestrator: "unsupported",
	}

	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("Unsupported orchestrator 'unsupported'"))
}

func TestOrchestratorsCmd_OrchestratorNotProvided(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	command := &getVersionsCmd{
		version: "1.1.1",
	}

	err := command.run(nil, nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("Must specify orchestrator for version '1.1.1'"))
}

func TestOrchestratorsCmd_UnsupportedVersion(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	command := &getVersionsCmd{
		orchestrator: "kubernetes",
		version:      "1.1.1",
	}

	err := command.run(nil, nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("Kubernetes version 1.1.1 is not supported"))
}

func TestOrchestratorsCmd_Valid(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	command := &getVersionsCmd{
		orchestrator: "kubernetes",
		version:      "1.7.14",
		output:       "json",
	}

	err := command.run(nil, nil)
	g.Expect(err).NotTo(HaveOccurred())
}
