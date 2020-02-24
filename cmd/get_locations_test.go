// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
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
	g.Expect(err).To(HaveOccurred())
}

func TestGetLocationsCmd_Run(t *testing.T) {
	t.Parallel()

	d := &locationsCmd{
		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs:      &authArgs{},
			getClientMock: &armhelpers.MockAKSEngineClient{},
		},
	}

	r := &cobra.Command{}
	f := r.Flags()

	addAuthFlags(d.getAuthArgs(), f)

	fakeRawSubscriptionID := "6dc93fae-9a76-421f-bbe5-cc6460ea81cb"
	fakeSubscriptionID, err := uuid.Parse(fakeRawSubscriptionID)
	fakeClientID := "b829b379-ca1f-4f1d-91a2-0d26b244680d"
	fakeClientSecret := "0se43bie-3zs5-303e-aav5-dcf231vb82ds"
	if err != nil {
		t.Fatalf("Invalid SubscriptionId in Test: %s", err)
	}

	d.getAuthArgs().SubscriptionID = fakeSubscriptionID
	d.getAuthArgs().rawSubscriptionID = fakeRawSubscriptionID
	d.getAuthArgs().rawClientID = fakeClientID
	d.getAuthArgs().ClientSecret = fakeClientSecret

	args := []string{}

	d.output = "human"
	err = d.run(r, args)
	if err != nil {
		t.Fatalf("Failed to call get-locations:` %s", err)
	}

	d.output = "json"
	err = d.run(r, args)
	if err != nil {
		t.Fatalf("Failed to call get-locations:` %s", err)
	}

	d.output = "code"
	err = d.run(r, args)
	if err != nil {
		t.Fatalf("Failed to call get-locations:` %s", err)
	}
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
