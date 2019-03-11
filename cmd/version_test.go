// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the version command", func() {
	It("should create a version command", func() {
		command := newVersionCmd()
		Expect(command.Use).Should(Equal(versionName))
		Expect(command.Short).Should(Equal(versionShortDescription))
		Expect(command.Long).Should(Equal(versionLongDescription))
		Expect(command.Flags().Lookup("output")).NotTo(BeNil())

		command.SetArgs([]string{})
		err := command.Execute()
		Expect(err).NotTo(HaveOccurred())
	})

	It("should print a json version of AKS Engine", func() {
		output, _ := getVersion("json")

		expectedOutput, _ := helpers.JSONMarshalIndent(version, "", "  ", false)

		Expect(output).Should(Equal(string(expectedOutput)))
	})

	It("should print a humanized version of AKS Engine", func() {
		output, _ := getVersion("human")

		expectedOutput := fmt.Sprintf("Version: %s\nGitCommit: %s\nGitTreeState: %s",
			BuildTag,
			BuildSHA,
			GitTreeState)

		Expect(output).Should(Equal(expectedOutput))
	})

	It("should error when asked for a yaml version of AKS Engine", func() {
		output, err := getVersion("yaml")

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("output format \"yaml\" is not supported"))
		Expect(output).To(BeEmpty())
	})
})
