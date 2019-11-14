// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetVersionsCmd_ShouldCreate(t *testing.T) {
	t.Parallel()

	command := newGetVersionsCmd()

	g := NewGomegaWithT(t)
	g.Expect(command.Use).Should(Equal(getVersionsName))
	g.Expect(command.Short).Should(Equal(getVersionsShortDescription))
	g.Expect(command.Long).Should(Equal(getVersionsLongDescription))
	g.Expect(command.Flags().Lookup("orchestrator")).To(BeNil())
	g.Expect(command.Flags().Lookup("version")).NotTo(BeNil())

	command.SetArgs([]string{})
	err := command.Execute()
	g.Expect(err).NotTo(HaveOccurred())
}

func TestGetVersionsCmd_ShouldSupportJSON(t *testing.T) {
	t.Parallel()

	command := &getVersionsCmd{
		orchestrator: "kubernetes",
		version:      "1.13.3",
		output:       "json",
	}
	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).NotTo(HaveOccurred())
}

func TestGetVersionsCmd_ShouldBeHumanReadable(t *testing.T) {
	t.Parallel()

	command := &getVersionsCmd{
		orchestrator: "kubernetes",
		version:      "1.13.3",
		output:       "human",
	}
	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).NotTo(HaveOccurred())
}

func TestGetVersionsCmd_ShouldErrorIfInvalidOption(t *testing.T) {
	t.Parallel()

	command := &getVersionsCmd{
		orchestrator: "kubernetes",
		version:      "1.13.3",
		output:       "yaml",
	}
	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("output format \"yaml\" is not supported"))
}
