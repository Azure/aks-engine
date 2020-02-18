// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetLocationsCmd(t *testing.T) {
	t.Parallel()

	command := newGetLocationsCmd()

	g := NewGomegaWithT(t)
	g.Expect(command.Use).Should(Equal(locationsName))
	g.Expect(command.Short).Should(Equal(locationsShortDescription))
	g.Expect(command.Long).Should(Equal(locationsLongDescription))
	g.Expect(command.Flags().Lookup("output")).NotTo(BeNil())

	command.SetArgs([]string{})
	err := command.Execute()
	// TODO: mock out this execution path
	g.Expect(err).To(HaveOccurred())
}

func TestGetLocationsCmd_ShouldErrorIfInvalidOption(t *testing.T) {
	t.Parallel()

	command := &locationsCmd{
		output: "yaml",
	}
	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("invalid output format: \"yaml\". Allowed values: human, json, code.\n"))
}
