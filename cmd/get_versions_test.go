// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The get-versions command", func() {
	It("should create a get-versions command", func() {
		output := newGetVersionsCmd()

		Expect(output.Use).Should(Equal(getVersionsName))
		Expect(output.Short).Should(Equal(getVersionsShortDescription))
		Expect(output.Long).Should(Equal(getVersionsLongDescription))
		Expect(output.Flags().Lookup("orchestrator")).To(BeNil())
		Expect(output.Flags().Lookup("version")).NotTo(BeNil())
	})

	It("should support JSON output", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.13.3",
			output:       "json",
		}
		err := command.run(nil, nil)
		Expect(err).To(BeNil())
	})

	It("should support human-readable output", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.13.3",
			output:       "human",
		}
		err := command.run(nil, nil)
		Expect(err).To(BeNil())
	})

	It("should error on an invalid output option", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.13.3",
			output:       "yaml",
		}
		err := command.run(nil, nil)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("output format \"yaml\" is not supported"))
	})
})
