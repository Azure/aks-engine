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

func TestGetSkusCmd(t *testing.T) {
	t.Parallel()

	command := newGetSkusCmd()

	g := NewGomegaWithT(t)
	g.Expect(command.Use).Should(Equal(skusName))
	g.Expect(command.Short).Should(Equal(skusShortDescription))
	g.Expect(command.Long).Should(Equal(skusLongDescription))
	g.Expect(command.Flags().Lookup("output")).NotTo(BeNil())

	command.SetArgs([]string{})
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

func ExampleSkusCmd_run_humanOutput() {
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

	d.output = "human"
	if err := d.run(r, args); err != nil {
		fmt.Printf("error running command: %s\n", err)
	}

	// Output:
	// Name                    Storage Account Type  Accelerated Networking Support
	// Standard_A0             Standard_LRS          false
	// Standard_A1             Standard_LRS          false
	// Standard_A10            Standard_LRS          false
	// Standard_A11            Standard_LRS          false
	// Standard_A1_v2          Standard_LRS          false
	// Standard_A2             Standard_LRS          false
	// Standard_A2_v2          Standard_LRS          false
	// Standard_A2m_v2         Standard_LRS          false
	// Standard_A3             Standard_LRS          false
	// Standard_A4             Standard_LRS          false
	// Standard_A4_v2          Standard_LRS          false
	// Standard_A4m_v2         Standard_LRS          false
	// Standard_A5             Standard_LRS          false
	// Standard_A6             Standard_LRS          false
	// Standard_A7             Standard_LRS          false
	// Standard_A8             Standard_LRS          false
	// Standard_A8_v2          Standard_LRS          false
	// Standard_A8m_v2         Standard_LRS          false
	// Standard_A9             Standard_LRS          false
	// Standard_B12ms          Premium_LRS           true
	// Standard_B16ms          Premium_LRS           true
	// Standard_B1ls           Premium_LRS           false
	// Standard_B1ms           Premium_LRS           false
	// Standard_B1s            Premium_LRS           false
	// Standard_B20ms          Premium_LRS           true
	// Standard_B2ms           Premium_LRS           false
	// Standard_B2s            Premium_LRS           false
	// Standard_B4ms           Premium_LRS           false
	// Standard_B8ms           Premium_LRS           false
	// Standard_D1             Standard_LRS          false
	// Standard_D11            Standard_LRS          false
	// Standard_D11_v2         Standard_LRS          true
	// Standard_D11_v2_Promo   Standard_LRS          true
	// Standard_D12            Standard_LRS          false
	// Standard_D12_v2         Standard_LRS          false
	// Standard_D12_v2         Standard_LRS          true
	// Standard_D12_v2_Promo   Standard_LRS          false
	// Standard_D13            Standard_LRS          false
	// Standard_D13_v2         Standard_LRS          true
	// Standard_D13_v2_Promo   Standard_LRS          true
	// Standard_D14            Standard_LRS          false
	// Standard_D14_v2         Standard_LRS          true
	// Standard_D14_v2_Promo   Standard_LRS          true
	// Standard_D15_v2         Standard_LRS          true
	// Standard_D16_v3         Standard_LRS          true
	// Standard_D16a_v3        Standard_LRS          false
	// Standard_D16a_v4        Standard_LRS          true
	// Standard_D16as_v3       Premium_LRS           false
	// Standard_D16as_v4       Premium_LRS           true
	// Standard_D16s_v3        Premium_LRS           true
	// Standard_D1_v2          Standard_LRS          false
	// Standard_D2             Standard_LRS          false
	// Standard_D2_v2          Standard_LRS          true
	// Standard_D2_v2_Promo    Standard_LRS          true
	// Standard_D2_v3          Standard_LRS          false
	// Standard_D2a_v3         Standard_LRS          false
	// Standard_D2a_v4         Standard_LRS          false
	// Standard_D2as_v3        Premium_LRS           false
	// Standard_D2as_v4        Premium_LRS           false
	// Standard_D2s_v3         Premium_LRS           false
	// Standard_D3             Standard_LRS          false
	// Standard_D32_v3         Standard_LRS          true
	// Standard_D32a_v3        Standard_LRS          false
	// Standard_D32a_v4        Standard_LRS          true
	// Standard_D32as_v3       Premium_LRS           false
	// Standard_D32as_v4       Premium_LRS           true
	// Standard_D32s_v3        Premium_LRS           true
	// Standard_D3_v2          Standard_LRS          true
	// Standard_D3_v2_Promo    Standard_LRS          true
	// Standard_D4             Standard_LRS          false
	// Standard_D48_v3         Standard_LRS          true
	// Standard_D48a_v3        Standard_LRS          false
	// Standard_D48a_v4        Standard_LRS          true
	// Standard_D48as_v3       Premium_LRS           false
	// Standard_D48as_v4       Premium_LRS           true
	// Standard_D48s_v3        Premium_LRS           true
	// Standard_D4_v2          Standard_LRS          true
	// Standard_D4_v2_Promo    Standard_LRS          true
	// Standard_D4_v3          Standard_LRS          true
	// Standard_D4a_v3         Standard_LRS          false
	// Standard_D4a_v4         Standard_LRS          true
	// Standard_D4as_v3        Premium_LRS           false
	// Standard_D4as_v4        Premium_LRS           true
	// Standard_D4s_v3         Premium_LRS           true
	// Standard_D5_v2          Standard_LRS          true
	// Standard_D5_v2_Promo    Standard_LRS          true
	// Standard_D64_v3         Standard_LRS          true
	// Standard_D64a_v3        Standard_LRS          false
	// Standard_D64a_v4        Standard_LRS          true
	// Standard_D64as_v3       Premium_LRS           false
	// Standard_D64as_v4       Premium_LRS           true
	// Standard_D64s_v3        Premium_LRS           true
	// Standard_D8_v3          Standard_LRS          true
	// Standard_D8a_v3         Standard_LRS          false
	// Standard_D8a_v4         Standard_LRS          true
	// Standard_D8as_v3        Premium_LRS           false
	// Standard_D8as_v4        Premium_LRS           true
	// Standard_D8s_v3         Premium_LRS           true
	// Standard_D96a_v3        Standard_LRS          false
	// Standard_D96a_v4        Standard_LRS          true
	// Standard_D96as_v3       Premium_LRS           false
	// Standard_D96as_v4       Premium_LRS           true
	// Standard_DC1s_v2        Premium_LRS           false
	// Standard_DC2s           Premium_LRS           false
	// Standard_DC2s_v2        Premium_LRS           true
	// Standard_DC4s           Premium_LRS           false
	// Standard_DC4s_v2        Premium_LRS           true
	// Standard_DC8_v2         Standard_LRS          true
	// Standard_DC8s           Premium_LRS           false
	// Standard_DS1            Premium_LRS           false
	// Standard_DS11           Premium_LRS           false
	// Standard_DS11-1_v2      Premium_LRS           false
	// Standard_DS11_v2        Premium_LRS           true
	// Standard_DS11_v2_Promo  Premium_LRS           true
	// Standard_DS12           Premium_LRS           false
	// Standard_DS12-1_v2      Premium_LRS           true
	// Standard_DS12-2_v2      Premium_LRS           true
	// Standard_DS12_v2        Premium_LRS           true
	// Standard_DS12_v2_Promo  Premium_LRS           true
	// Standard_DS13           Premium_LRS           false
	// Standard_DS13-2_v2      Premium_LRS           true
	// Standard_DS13-4_v2      Premium_LRS           true
	// Standard_DS13_v2        Premium_LRS           true
	// Standard_DS13_v2_Promo  Premium_LRS           true
	// Standard_DS14           Premium_LRS           false
	// Standard_DS14-4_v2      Premium_LRS           true
	// Standard_DS14-8_v2      Premium_LRS           true
	// Standard_DS14_v2        Premium_LRS           true
	// Standard_DS14_v2_Promo  Premium_LRS           true
	// Standard_DS15_v2        Premium_LRS           true
	// Standard_DS1_v2         Premium_LRS           false
	// Standard_DS2            Premium_LRS           false
	// Standard_DS2_v2         Premium_LRS           true
	// Standard_DS2_v2_Promo   Premium_LRS           true
	// Standard_DS3            Premium_LRS           false
	// Standard_DS3_v2         Premium_LRS           true
	// Standard_DS3_v2_Promo   Premium_LRS           true
	// Standard_DS4            Premium_LRS           false
	// Standard_DS4_v2         Premium_LRS           true
	// Standard_DS4_v2_Promo   Premium_LRS           true
	// Standard_DS5_v2         Premium_LRS           false
	// Standard_DS5_v2_Promo   Premium_LRS           false
	// Standard_E16-4s_v3      Premium_LRS           true
	// Standard_E16-8s_v3      Premium_LRS           true
	// Standard_E16_v3         Standard_LRS          true
	// Standard_E16a_v3        Standard_LRS          false
	// Standard_E16a_v4        Standard_LRS          true
	// Standard_E16as_v3       Premium_LRS           false
	// Standard_E16as_v4       Premium_LRS           true
	// Standard_E16s_v3        Premium_LRS           true
	// Standard_E20_v3         Standard_LRS          true
	// Standard_E20a_v4        Standard_LRS          true
	// Standard_E20as_v4       Premium_LRS           true
	// Standard_E20s_v3        Premium_LRS           true
	// Standard_E2_v3          Standard_LRS          false
	// Standard_E2a_v3         Standard_LRS          false
	// Standard_E2a_v4         Standard_LRS          false
	// Standard_E2as_v3        Premium_LRS           false
	// Standard_E2as_v4        Premium_LRS           false
	// Standard_E2s_v3         Premium_LRS           false
	// Standard_E32-16s_v3     Premium_LRS           true
	// Standard_E32-8s_v3      Premium_LRS           true
	// Standard_E32_v3         Standard_LRS          true
	// Standard_E32a_v3        Standard_LRS          false
	// Standard_E32a_v4        Standard_LRS          true
	// Standard_E32as_v3       Premium_LRS           false
	// Standard_E32as_v4       Premium_LRS           true
	// Standard_E32s_v3        Premium_LRS           true
	// Standard_E4-2s_v3       Premium_LRS           true
	// Standard_E48_v3         Standard_LRS          true
	// Standard_E48a_v3        Standard_LRS          false
	// Standard_E48a_v4        Standard_LRS          true
	// Standard_E48as_v3       Premium_LRS           false
	// Standard_E48as_v4       Premium_LRS           true
	// Standard_E48s_v3        Premium_LRS           true
	// Standard_E4_v3          Standard_LRS          true
	// Standard_E4a_v3         Standard_LRS          false
	// Standard_E4a_v4         Standard_LRS          true
	// Standard_E4as_v3        Premium_LRS           false
	// Standard_E4as_v4        Premium_LRS           true
	// Standard_E4s_v3         Premium_LRS           true
	// Standard_E64-16s_v3     Premium_LRS           true
	// Standard_E64-32s_v3     Premium_LRS           true
	// Standard_E64_v3         Standard_LRS          true
	// Standard_E64a_v3        Standard_LRS          false
	// Standard_E64a_v4        Standard_LRS          true
	// Standard_E64as_v3       Premium_LRS           false
	// Standard_E64as_v4       Premium_LRS           true
	// Standard_E64i_v3        Standard_LRS          true
	// Standard_E64is_v3       Premium_LRS           true
	// Standard_E64s_v3        Premium_LRS           true
	// Standard_E8-2s_v3       Premium_LRS           true
	// Standard_E8-4s_v3       Premium_LRS           true
	// Standard_E8_v3          Standard_LRS          true
	// Standard_E8a_v3         Standard_LRS          false
	// Standard_E8a_v4         Standard_LRS          true
	// Standard_E8as_v3        Premium_LRS           false
	// Standard_E8as_v4        Premium_LRS           true
	// Standard_E8s_v3         Premium_LRS           true
	// Standard_E96a_v4        Standard_LRS          true
	// Standard_E96as_v4       Premium_LRS           true
	// Standard_F1             Standard_LRS          false
	// Standard_F16            Standard_LRS          true
	// Standard_F16s           Premium_LRS           true
	// Standard_F16s_v2        Premium_LRS           true
	// Standard_F1s            Premium_LRS           false
	// Standard_F2             Standard_LRS          true
	// Standard_F2s            Premium_LRS           true
	// Standard_F2s_v2         Premium_LRS           false
	// Standard_F32s_v2        Premium_LRS           true
	// Standard_F4             Standard_LRS          true
	// Standard_F48s_v2        Premium_LRS           true
	// Standard_F4s            Premium_LRS           true
	// Standard_F4s_v2         Premium_LRS           true
	// Standard_F64s_v2        Premium_LRS           true
	// Standard_F72s_v2        Premium_LRS           true
	// Standard_F8             Standard_LRS          true
	// Standard_F8s            Premium_LRS           true
	// Standard_F8s_v2         Premium_LRS           true
	// Standard_G1             Standard_LRS          false
	// Standard_G2             Standard_LRS          false
	// Standard_G3             Standard_LRS          false
	// Standard_G4             Standard_LRS          false
	// Standard_G5             Standard_LRS          false
	// Standard_GS1            Premium_LRS           false
	// Standard_GS2            Premium_LRS           false
	// Standard_GS3            Premium_LRS           false
	// Standard_GS4            Premium_LRS           false
	// Standard_GS4-4          Premium_LRS           false
	// Standard_GS4-8          Premium_LRS           false
	// Standard_GS5            Premium_LRS           false
	// Standard_GS5-16         Premium_LRS           false
	// Standard_GS5-8          Premium_LRS           false
	// Standard_H16            Standard_LRS          false
	// Standard_H16_Promo      Standard_LRS          false
	// Standard_H16m           Standard_LRS          false
	// Standard_H16m_Promo     Standard_LRS          false
	// Standard_H16mr          Standard_LRS          false
	// Standard_H16mr_Promo    Standard_LRS          false
	// Standard_H16r           Standard_LRS          false
	// Standard_H16r_Promo     Standard_LRS          false
	// Standard_H8             Standard_LRS          false
	// Standard_H8_Promo       Standard_LRS          false
	// Standard_H8m            Premium_LRS           false
	// Standard_H8m_Promo      Standard_LRS          false
	// Standard_HB120rs_v2     Premium_LRS           false
	// Standard_HB60rs         Premium_LRS           false
	// Standard_HC44rs         Standard_LRS          false
	// Standard_L16s           Premium_LRS           false
	// Standard_L16s_v2        Premium_LRS           true
	// Standard_L32s           Premium_LRS           false
	// Standard_L32s_v2        Premium_LRS           true
	// Standard_L48s_v2        Premium_LRS           true
	// Standard_L4s            Premium_LRS           false
	// Standard_L64s_v2        Premium_LRS           true
	// Standard_L80s_v2        Premium_LRS           true
	// Standard_L8s            Premium_LRS           false
	// Standard_L8s_v2         Premium_LRS           true
	// Standard_LRS            Premium_LRS           false
	// Standard_M128           Standard_LRS          true
	// Standard_M128-32ms      Premium_LRS           true
	// Standard_M128-64ms      Premium_LRS           true
	// Standard_M128m          Standard_LRS          true
	// Standard_M128ms         Premium_LRS           true
	// Standard_M128s          Premium_LRS           true
	// Standard_M16-4ms        Premium_LRS           true
	// Standard_M16-8ms        Premium_LRS           true
	// Standard_M16ms          Premium_LRS           true
	// Standard_M192ms_v2      Premium_LRS           true
	// Standard_M192s_v2       Premium_LRS           true
	// Standard_M208ms_v2      Premium_LRS           true
	// Standard_M208s_v2       Premium_LRS           true
	// Standard_M24ms_v2       Premium_LRS           true
	// Standard_M24s_v2        Premium_LRS           true
	// Standard_M32-16ms       Premium_LRS           true
	// Standard_M32-8ms        Premium_LRS           true
	// Standard_M32ls          Premium_LRS           true
	// Standard_M32ms          Premium_LRS           true
	// Standard_M32ts          Premium_LRS           true
	// Standard_M416ms_v2      Premium_LRS           true
	// Standard_M416s_v2       Premium_LRS           true
	// Standard_M48ms_v2       Premium_LRS           true
	// Standard_M48s_v2        Premium_LRS           true
	// Standard_M64            Standard_LRS          true
	// Standard_M64-16ms       Premium_LRS           true
	// Standard_M64-32ms       Premium_LRS           true
	// Standard_M64ls          Premium_LRS           true
	// Standard_M64m           Standard_LRS          true
	// Standard_M64ms          Premium_LRS           true
	// Standard_M64s           Premium_LRS           false
	// Standard_M8-2ms         Premium_LRS           false
	// Standard_M8-4ms         Premium_LRS           false
	// Standard_M8ms           Premium_LRS           true
	// Standard_M96ms_v2       Premium_LRS           true
	// Standard_M96s_v2        Premium_LRS           true
	// Standard_NC12           Standard_LRS          false
	// Standard_NC12_Promo     Standard_LRS          false
	// Standard_NC12s_v2       Premium_LRS           false
	// Standard_NC12s_v3       Premium_LRS           false
	// Standard_NC24           Standard_LRS          false
	// Standard_NC24_Promo     Standard_LRS          false
	// Standard_NC24r          Standard_LRS          false
	// Standard_NC24r_Promo    Standard_LRS          false
	// Standard_NC24rs_v2      Premium_LRS           false
	// Standard_NC24rs_v3      Premium_LRS           false
	// Standard_NC24s_v2       Premium_LRS           false
	// Standard_NC24s_v3       Premium_LRS           false
	// Standard_NC6            Standard_LRS          false
	// Standard_NC6_Promo      Standard_LRS          false
	// Standard_NC6s_v2        Premium_LRS           false
	// Standard_NC6s_v3        Premium_LRS           false
	// Standard_ND12s          Premium_LRS           false
	// Standard_ND24rs         Premium_LRS           false
	// Standard_ND24s          Premium_LRS           false
	// Standard_ND40rs_v2      Premium_LRS           false
	// Standard_ND40s_v3       Premium_LRS           true
	// Standard_ND6s           Premium_LRS           false
	// Standard_NP10s          Premium_LRS           false
	// Standard_NP20s          Premium_LRS           false
	// Standard_NP40s          Premium_LRS           false
	// Standard_NV12           Standard_LRS          false
	// Standard_NV12_Promo     Standard_LRS          false
	// Standard_NV12s_v2       Premium_LRS           false
	// Standard_NV12s_v3       Premium_LRS           false
	// Standard_NV16as_v4      Premium_LRS           false
	// Standard_NV24           Standard_LRS          false
	// Standard_NV24_Promo     Standard_LRS          false
	// Standard_NV24s_v2       Premium_LRS           false
	// Standard_NV24s_v3       Premium_LRS           false
	// Standard_NV32as_v4      Premium_LRS           false
	// Standard_NV48s_v3       Premium_LRS           false
	// Standard_NV4as_v4       Premium_LRS           false
	// Standard_NV6            Standard_LRS          false
	// Standard_NV6_Promo      Standard_LRS          false
	// Standard_NV6s_v2        Premium_LRS           false
	// Standard_NV8as_v4       Premium_LRS           false
	// Standard_PB12s          Premium_LRS           false
	// Standard_PB24s          Premium_LRS           false
	// Standard_PB6s           Premium_LRS           false
	// Standard_ZRS            Premium_LRS           false
}

func ExampleSkusCmd_run_jsonOutput() {
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

	d.output = "json"
	if err := d.run(r, args); err != nil {
		fmt.Printf("error running command: %s\n", err)
	}

	// Output:
	// [
	//   {
	//     "Name": "Standard_A0",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A1",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A10",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A11",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A1_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A2_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A2m_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A4_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A4m_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A5",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A6",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A7",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A8",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A8_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A8m_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A9",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B12ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_B16ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_B1ls",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B1ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B1s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B20ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_B2ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B2s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B4ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B8ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D1",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D11",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D11_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D11_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D12",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D12_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D12_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D12_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D13",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D13_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D13_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D14",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D14_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D14_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D15_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D16a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D16as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D1_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D2_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D2_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D32_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D32a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D32as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D3_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D3_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D48_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D48a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D48as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D4a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D4as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D5_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D5_v2_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D64a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D64as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D8a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D8as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D96a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D96a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D96as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D96as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC1s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DC2s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DC2s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC4s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DC4s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC8_v2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC8s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS1",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS11",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS11-1_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS11_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS11_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS12-1_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12-2_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS13-2_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13-4_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS14-4_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14-8_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS15_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS1_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS2_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS2_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS3_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS3_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS4_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS4_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS5_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS5_v2_Promo",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E16-4s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16-8s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E16a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E16as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E2_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E32-16s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-8s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E32a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E32as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4-2s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E48a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E48as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E4a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E4as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-16s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-32s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E64a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E64as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64i_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64is_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-2s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-4s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8a_v3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E8a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8as_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E8as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E96a_v4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E96as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F1",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_F16",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F16s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F16s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F1s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_F2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F2s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F2s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_F32s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F48s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F4s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F4s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F64s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F72s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F8",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F8s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F8s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_G1",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G2",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G3",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G4",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G5",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS1",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS4-4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS4-8",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS5",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS5-16",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS5-8",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16m",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16m_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16mr",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16mr_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16r",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16r_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8m",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8m_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_HB120rs_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_HB60rs",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_HC44rs",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L16s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L16s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L32s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L32s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L48s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L4s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L64s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L80s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L8s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L8s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_LRS",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M128",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128-32ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128-64ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128m",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M16-4ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M16-8ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M16ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192ms_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M208ms_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M208s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M24ms_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M24s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32-16ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32-8ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32ls",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32ts",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M416ms_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M416s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M48ms_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M48s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64-16ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64-32ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64ls",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64m",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M8-2ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M8-4ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M8ms",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M96ms_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M96s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_NC12",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC12_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC12s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC12s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24r",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24r_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24rs_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24rs_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC6",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC6_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC6s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC6s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND12s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND24rs",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND24s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND40rs_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND40s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_ND6s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NP10s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NP20s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NP40s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV16as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV32as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV48s_v3",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV4as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV6",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV6_Promo",
	//     "StorageAccountType": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV6s_v2",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV8as_v4",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_PB12s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_PB24s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_PB6s",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ZRS",
	//     "StorageAccountType": "Premium_LRS",
	//     "AcceleratedNetworking": false
	//   }
	// ]
}

func ExampleSkusCmd_run_codeOutput() {
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
		fmt.Printf("error running command: %s\n", err)
	}

	// Output:
	// // Copyright (c) Microsoft Corporation. All rights reserved.
	// // Licensed under the MIT license.

	// package helpers

	// // GetKubernetesAllowedVMSKUs provides the allowed sizes for Kubernetes agent VMs.
	// //
	// // Code generated for package helpers by aks-engine DO NOT EDIT. (@generated)
	// //
	// // To generate this code, run the command:
	// //   aks-engine get-skus --output=code

	// type VMSku struct {
	// 	Name                  string
	// 	StorageAccountType    string
	// 	AcceleratedNetworking bool
	// }

	// var VMSkus = []VMSku{
	// 	{
	// 		Name:                  "Standard_A0",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A1",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A10",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A11",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A1_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A2_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A2m_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A4_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A4m_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A5",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 			Name:                  "Standard_A6",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A7",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A8",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A8_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A8m_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A9",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B12ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_B16ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_B1ls",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B1ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B1s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B20ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_B2ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B2s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B4ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B8ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D1",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D11",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D11_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D11_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D13",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D13_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D13_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D14",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D14_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D14_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D15_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D1_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D3_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D3_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D5_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D5_v2_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC1s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC2s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC2s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC4s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC4s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC8_v2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC8s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS1",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11-1_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12-1_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12-2_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13-2_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13-4_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14-4_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14-8_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS15_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS1_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS2_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS2_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS3_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS3_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS4_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS4_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS5_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS5_v2_Promo",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-4s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-8s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-16s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-8s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4-2s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-16s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-32s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64i_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64is_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8-2s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8-4s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8a_v3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8as_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E96a_v4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E96as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F1",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_F16",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F16s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F16s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F1s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_F2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F2s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F2s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_F32s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F48s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F4s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F4s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F64s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F72s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F8",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F8s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F8s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_G1",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G2",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G3",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G4",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G5",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS1",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS4-4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS4-8",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS5",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS5-16",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS5-8",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16m",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16m_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16mr",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16mr_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16r",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16r_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8m",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8m_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_HB120rs_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_HB60rs",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_HC44rs",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L16s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L16s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L32s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L32s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L48s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L4s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L64s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L80s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L8s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L8s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_LRS",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128-32ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128-64ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128m",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M16-4ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M16-8ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M16ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M192ms_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M192s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M208ms_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M208s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M24ms_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M24s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32-16ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32-8ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32ls",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32ts",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M416ms_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M416s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M48ms_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M48s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64-16ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64-32ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64ls",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64m",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M8-2ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M8-4ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M8ms",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M96ms_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M96s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24r",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24r_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24rs_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24rs_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND12s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND24rs",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND24s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND40rs_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND40s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND6s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NP10s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NP20s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NP40s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV16as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV32as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV48s_v3",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV4as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV6",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV6_Promo",
	// 		StorageAccountType:    "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV6s_v2",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV8as_v4",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_PB12s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_PB24s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_PB6s",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ZRS",
	// 		StorageAccountType:    "Premium_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// }

}

func TestGetSkusCmd_ShouldErrorIfInvalidOption(t *testing.T) {
	t.Parallel()

	command := &SkusCmd{
		output: "yaml",
	}
	err := command.run(nil, nil)
	g := NewGomegaWithT(t)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("invalid output format: \"yaml\". Allowed values: human, json, code.\n"))
}
