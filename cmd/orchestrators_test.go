// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The orchestrators command", func() {
	It("should create an orchestrators command", func() {
		output := newOrchestratorsCmd()

		Expect(output.Use).Should(Equal(orchestratorsName))
		Expect(output.Short).Should(Equal(orchestratorsShortDescription))
		Expect(output.Long).Should(Equal(orchestratorsLongDescription))
		Expect(output.Flags().Lookup("orchestrator")).NotTo(BeNil())
		Expect(output.Flags().Lookup("version")).NotTo(BeNil())
	})

	It("should fail on unsupported orchestrator", func() {
		command := &getVersionsCmd{
			orchestrator: "unsupported",
		}

		err := command.run(nil, nil)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("Unsupported orchestrator 'unsupported'"))
	})

	It("should fail on unprovided orchestrator", func() {
		command := &getVersionsCmd{
			version: "1.1.1",
		}

		err := command.run(nil, nil)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("Must specify orchestrator for version '1.1.1'"))
	})

	It("should fail on unsupported version", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.1.1",
		}

		err := command.run(nil, nil)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("Kubernetes version 1.1.1 is not supported"))
	})

	It("should succeed", func() {
		command := &getVersionsCmd{
			orchestrator: "kubernetes",
			version:      "1.7.14",
			output:       "json",
		}

		err := command.run(nil, nil)
		Expect(err).To(BeNil())
	})
})
