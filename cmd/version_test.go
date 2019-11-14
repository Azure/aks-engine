// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/aks-engine/pkg/helpers"
)

func TestVersionCommand_ShouldCreateCommand(t *testing.T) {
	t.Parallel()

	command := newVersionCmd()
	g := NewGomegaWithT(t)
	g.Expect(command.Use).Should(Equal(versionName))
	g.Expect(command.Short).Should(Equal(versionShortDescription))
	g.Expect(command.Long).Should(Equal(versionLongDescription))
	g.Expect(command.Flags().Lookup("output")).NotTo(BeNil())

	command.SetArgs([]string{})
	err := command.Execute()
	g.Expect(err).NotTo(HaveOccurred())
}

func TestVersionCommand_ShouldPrintJsonVersionOfAKSEngine(t *testing.T) {
	t.Parallel()

	output, _ := getVersion("json")
	expectedOutput, _ := helpers.JSONMarshalIndent(version, "", "  ", false)
	g := NewGomegaWithT(t)
	g.Expect(output).Should(Equal(string(expectedOutput)))
}

func TestVersionCommand_ShouldPrintHumanizedVersionOfAKSEngine(t *testing.T) {
	t.Parallel()

	output, _ := getVersion("human")
	expectedOutput := fmt.Sprintf("Version: %s\nGitCommit: %s\nGitTreeState: %s",
		BuildTag,
		BuildSHA,
		GitTreeState)
	g := NewGomegaWithT(t)
	g.Expect(output).Should(Equal(expectedOutput))
}

func TestVersionCommand_ShouldErrorWhenAskedForYamlVersion(t *testing.T) {
	t.Parallel()

	output, err := getVersion("yaml")
	g := NewGomegaWithT(t)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("output format \"yaml\" is not supported"))
	g.Expect(output).To(BeEmpty())
}
