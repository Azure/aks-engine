// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The get-versions command", func() {
	It("should create a get-versions command", func() {
		command := newGetVersionsCmd()

		Expect(command.Use).Should(Equal(getVersionsName))
		Expect(command.Short).Should(Equal(getVersionsShortDescription))
		Expect(command.Long).Should(Equal(getVersionsLongDescription))
		Expect(command.Flags().Lookup("orchestrator")).To(BeNil())
		Expect(command.Flags().Lookup("version")).NotTo(BeNil())

		command.SetArgs([]string{})
		err := command.Execute()
		Expect(err).NotTo(HaveOccurred())
	})

	It("should support JSON output", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.13.3",
			output:       "json",
		}
		err := command.run(nil, nil)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should support human-readable output", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.13.3",
			output:       "human",
		}
		err := command.run(nil, nil)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should error on an invalid output option", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.13.3",
			output:       "yaml",
		}
		err := command.run(nil, nil)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("output format \"yaml\" is not supported"))
	})
})
