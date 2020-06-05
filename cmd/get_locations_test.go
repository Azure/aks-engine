// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
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

func TestLocationsCmd_run(t *testing.T) {
	d := &LocationsCmd{
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
		t.Fatalf("Failed to call get-locations:` %s", err)
	}
}

func ExampleLocationsCmd_run_humanOutput() {
	d := &LocationsCmd{
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

	d.output = "human"
	if err := d.run(r, args); err != nil {
		fmt.Printf("error running command: %s\n", err)
	}

	// Output:
	// Location          Name                      Latitude  Longitude
	// centraluseuap     Central US EUAP (Canary)  N/A       N/A
	// chinaeast         China East                N/A       N/A
	// chinaeast2        China East 2              N/A       N/A
	// chinanorth        China North               N/A       N/A
	// chinanorth2       China North 2             N/A       N/A
	// eastus2euap       East US 2 EUAP (Canary)   N/A       N/A
	// germanycentral    Germany Central           N/A       N/A
	// germanynortheast  Germany Northeast         N/A       N/A
	// usdodcentral      US DoD Central            N/A       N/A
	// usdodeast         US Dod East               N/A       N/A
	// usgovarizona      US Gov Arizona            N/A       N/A
	// usgoviowa         US Gov Iowa               N/A       N/A
	// usgovtexas        US Gov Texas              N/A       N/A
	// usgovvirginia     US Gov Virginia           N/A       N/A
}

func ExampleLocationsCmd_run_jsonOutput() {
	d := &LocationsCmd{
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

	d.output = "json"
	if err := d.run(r, args); err != nil {
		fmt.Printf("error running command: %s\n", err)
	}

	// Output:
	// [
	//   {
	//     "id": "N/A",
	//     "name": "centraluseuap",
	//     "displayName": "Central US EUAP (Canary)",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "chinaeast",
	//     "displayName": "China East",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "chinaeast2",
	//     "displayName": "China East 2",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "chinanorth",
	//     "displayName": "China North",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "chinanorth2",
	//     "displayName": "China North 2",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "eastus2euap",
	//     "displayName": "East US 2 EUAP (Canary)",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "germanycentral",
	//     "displayName": "Germany Central",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "germanynortheast",
	//     "displayName": "Germany Northeast",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "usdodcentral",
	//     "displayName": "US DoD Central",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "usdodeast",
	//     "displayName": "US Dod East",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "usgovarizona",
	//     "displayName": "US Gov Arizona",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "usgoviowa",
	//     "displayName": "US Gov Iowa",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "usgovtexas",
	//     "displayName": "US Gov Texas",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   },
	//   {
	//     "id": "N/A",
	//     "name": "usgovvirginia",
	//     "displayName": "US Gov Virginia",
	//     "latitude": "N/A",
	//     "longitude": "N/A"
	//   }
	// ]
}

func ExampleLocationsCmd_run_codeOutput() {
	d := &LocationsCmd{
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
		fmt.Printf("error running command: %s\n", err)
	}

	// Output:
	// // Copyright (c) Microsoft Corporation. All rights reserved.
	// // Licensed under the MIT license.

	// package helpers

	// // GetAzureLocations provides all available Azure cloud locations.
	// //
	// // Code generated for package helpers by aks-engine DO NOT EDIT. (@generated)
	// //
	// // To generate this code, run the command:
	// //   aks-engine get-locations --output=code
	// func GetAzureLocations() []string {
	// 	return []string{
	// 		"centraluseuap",
	// 		"chinaeast",
	// 		"chinaeast2",
	// 		"chinanorth",
	// 		"chinanorth2",
	// 		"eastus2euap",
	// 		"germanycentral",
	// 		"germanynortheast",
	// 		"usdodcentral",
	// 		"usdodeast",
	// 		"usgovarizona",
	// 		"usgoviowa",
	// 		"usgovtexas",
	// 		"usgovvirginia",
	// 	 }
	// }
}

func TestGetLocationsCmd_ShouldErrorIfInvalidOption(t *testing.T) {
	t.Parallel()

	command := &LocationsCmd{
		output: "yaml",
	}
	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("invalid output format: \"yaml\". Allowed values: human, json, code.\n"))
}
