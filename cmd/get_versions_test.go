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
})
