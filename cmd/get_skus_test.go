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

func TestGetSkusCmd(t *testing.T) {
	t.Parallel()

	command := newGetSkusCmd()

	g := NewGomegaWithT(t)
	g.Expect(command.Use).Should(Equal(skusName))
	g.Expect(command.Short).Should(Equal(skusShortDescription))
	g.Expect(command.Long).Should(Equal(skusLongDescription))
	g.Expect(command.Flags().Lookup("output")).NotTo(BeNil())

	command.SetArgs([]string{"--bogus"})
	err := command.Execute()
	g.Expect(err).To(HaveOccurred())
}

func TestGetSkusCmd_run(t *testing.T) {
	d := &SkusCmd{
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
	if err != nil {
		t.Fatalf("Invalid SubscriptionId in Test: %s", err)
	}
	fakeClientID := "b829b379-ca1f-4f1d-91a2-0d26b244680d"
	fakeClientSecret := "0se43bie-3zs5-303e-aav5-dcf231vb82ds"

	d.getAuthArgs().SubscriptionID = fakeSubscriptionID
	d.getAuthArgs().rawSubscriptionID = fakeRawSubscriptionID
	d.getAuthArgs().rawClientID = fakeClientID
	d.getAuthArgs().ClientSecret = fakeClientSecret

	args := []string{}

	d.output = "human"
	err = d.run(r, args)
	if err != nil {
		t.Fatalf("Failed to call get-skus:` %s", err)
	}
}

func TestSkusCmd_run_codeOutput(t *testing.T) {
	d := &SkusCmd{
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
	fakeSubscriptionID, _ := uuid.Parse(fakeRawSubscriptionID)
	fakeClientID := "b829b379-ca1f-4f1d-91a2-0d26b244680d"
	fakeClientSecret := "0se43bie-3zs5-303e-aav5-dcf231vb82ds"

	d.getAuthArgs().SubscriptionID = fakeSubscriptionID
	d.getAuthArgs().rawSubscriptionID = fakeRawSubscriptionID
	d.getAuthArgs().rawClientID = fakeClientID
	d.getAuthArgs().ClientSecret = fakeClientSecret

	args := []string{}

	d.output = "code"
	if err := d.run(r, args); err != nil {
		t.Fatalf("Failed to call get-skus -o code:` %s", err)
	}
}

func TestGetSkusCmd_ShouldErrorIfInvalidOption(t *testing.T) {
	t.Parallel()

	command := &SkusCmd{
		output: "yaml",
	}
	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("invalid output format: \"yaml\". Allowed values: human, json, code"))
}
