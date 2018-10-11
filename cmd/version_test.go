package cmd

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the version command", func() {
	It("should create a version command", func() {
		output := newVersionCmd()
		Expect(output.Use).Should(Equal(versionName))
		Expect(output.Short).Should(Equal(versionShortDescription))
		Expect(output.Long).Should(Equal(versionLongDescription))
		Expect(output.Flags().Lookup("output")).NotTo(BeNil())
	})

	It("should print a json version of AKS Engine", func() {
		output := getVersion("json")

		expectedOutput, _ := helpers.JSONMarshalIndent(version, "", "  ", false)

		Expect(output).Should(Equal(string(expectedOutput)))
	})
	It("should print a humanized version of AKS Engine", func() {
		output := getVersion("human")

		expectedOutput := fmt.Sprintf("Version: %s\nGitCommit: %s\nGitTreeState: %s",
			BuildTag,
			BuildSHA,
			GitTreeState)

		Expect(output).Should(Equal(expectedOutput))
	})

	It("should print a json version of AKS Engine", func() {
		output := getVersion("json")

		expectedOutput, _ := helpers.JSONMarshalIndent(version, "", "  ", false)

		Expect(output).Should(Equal(string(expectedOutput)))
	})
})
