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
	// Name                    Accelerated Networking Support
	// Standard_A0             false
	// Standard_A1             false
	// Standard_A10            false
	// Standard_A11            false
	// Standard_A1_v2          false
	// Standard_A2             false
	// Standard_A2_v2          false
	// Standard_A2m_v2         false
	// Standard_A3             false
	// Standard_A4             false
	// Standard_A4_v2          false
	// Standard_A4m_v2         false
	// Standard_A5             false
	// Standard_A6             false
	// Standard_A7             false
	// Standard_A8             false
	// Standard_A8_v2          false
	// Standard_A8m_v2         false
	// Standard_A9             false
	// Standard_B12ms          true
	// Standard_B16ms          true
	// Standard_B1ls           false
	// Standard_B1ms           false
	// Standard_B1s            false
	// Standard_B20ms          true
	// Standard_B2ms           false
	// Standard_B2s            false
	// Standard_B4ms           false
	// Standard_B8ms           false
	// Standard_D1             false
	// Standard_D11            false
	// Standard_D11_v2         true
	// Standard_D11_v2_Promo   true
	// Standard_D12            false
	// Standard_D12_v2         false
	// Standard_D12_v2         true
	// Standard_D12_v2_Promo   false
	// Standard_D13            false
	// Standard_D13_v2         true
	// Standard_D13_v2_Promo   true
	// Standard_D14            false
	// Standard_D14_v2         true
	// Standard_D14_v2_Promo   true
	// Standard_D15_v2         true
	// Standard_D16_v3         true
	// Standard_D16_v4         true
	// Standard_D16a_v3        false
	// Standard_D16a_v4        true
	// Standard_D16as_v3       false
	// Standard_D16as_v4       true
	// Standard_D16d_v4        true
	// Standard_D16ds_v4       true
	// Standard_D16s_v3        true
	// Standard_D16s_v4        true
	// Standard_D1_v2          false
	// Standard_D2             false
	// Standard_D2_v2          true
	// Standard_D2_v2_Promo    true
	// Standard_D2_v3          false
	// Standard_D2_v4          false
	// Standard_D2a_v3         false
	// Standard_D2a_v4         false
	// Standard_D2as_v3        false
	// Standard_D2as_v4        false
	// Standard_D2d_v4         false
	// Standard_D2ds_v4        false
	// Standard_D2s_v3         false
	// Standard_D2s_v4         false
	// Standard_D3             false
	// Standard_D32_v3         true
	// Standard_D32_v4         true
	// Standard_D32a_v3        false
	// Standard_D32a_v4        true
	// Standard_D32as_v3       false
	// Standard_D32as_v4       true
	// Standard_D32d_v4        true
	// Standard_D32ds_v4       true
	// Standard_D32s_v3        true
	// Standard_D32s_v4        true
	// Standard_D3_v2          true
	// Standard_D3_v2_Promo    true
	// Standard_D4             false
	// Standard_D48_v3         true
	// Standard_D48_v4         true
	// Standard_D48a_v3        false
	// Standard_D48a_v4        true
	// Standard_D48as_v3       false
	// Standard_D48as_v4       true
	// Standard_D48d_v4        true
	// Standard_D48ds_v4       true
	// Standard_D48s_v3        true
	// Standard_D48s_v4        true
	// Standard_D4_v2          true
	// Standard_D4_v2_Promo    true
	// Standard_D4_v3          true
	// Standard_D4_v4          true
	// Standard_D4a_v3         false
	// Standard_D4a_v4         true
	// Standard_D4as_v3        false
	// Standard_D4as_v4        true
	// Standard_D4d_v4         true
	// Standard_D4ds_v4        true
	// Standard_D4s_v3         true
	// Standard_D4s_v4         true
	// Standard_D5_v2          true
	// Standard_D5_v2_Promo    true
	// Standard_D64_v3         true
	// Standard_D64_v4         true
	// Standard_D64a_v3        false
	// Standard_D64a_v4        true
	// Standard_D64as_v3       false
	// Standard_D64as_v4       true
	// Standard_D64d_v4        true
	// Standard_D64ds_v4       true
	// Standard_D64s_v3        true
	// Standard_D64s_v4        true
	// Standard_D8_v3          true
	// Standard_D8_v4          true
	// Standard_D8a_v3         false
	// Standard_D8a_v4         true
	// Standard_D8as_v3        false
	// Standard_D8as_v4        true
	// Standard_D8d_v4         true
	// Standard_D8ds_v4        true
	// Standard_D8s_v3         true
	// Standard_D8s_v4         true
	// Standard_D96a_v3        false
	// Standard_D96a_v4        true
	// Standard_D96as_v3       false
	// Standard_D96as_v4       true
	// Standard_DC1s_v2        false
	// Standard_DC2s           false
	// Standard_DC2s_v2        true
	// Standard_DC4s           false
	// Standard_DC4s_v2        true
	// Standard_DC8_v2         true
	// Standard_DC8s           false
	// Standard_DS1            false
	// Standard_DS11           false
	// Standard_DS11-1_v2      false
	// Standard_DS11_v2        true
	// Standard_DS11_v2_Promo  true
	// Standard_DS12           false
	// Standard_DS12-1_v2      true
	// Standard_DS12-2_v2      true
	// Standard_DS12_v2        true
	// Standard_DS12_v2_Promo  true
	// Standard_DS13           false
	// Standard_DS13-2_v2      true
	// Standard_DS13-4_v2      true
	// Standard_DS13_v2        true
	// Standard_DS13_v2_Promo  true
	// Standard_DS14           false
	// Standard_DS14-4_v2      true
	// Standard_DS14-8_v2      true
	// Standard_DS14_v2        true
	// Standard_DS14_v2_Promo  true
	// Standard_DS15_v2        true
	// Standard_DS1_v2         false
	// Standard_DS2            false
	// Standard_DS2_v2         true
	// Standard_DS2_v2_Promo   true
	// Standard_DS3            false
	// Standard_DS3_v2         true
	// Standard_DS3_v2_Promo   true
	// Standard_DS4            false
	// Standard_DS4_v2         true
	// Standard_DS4_v2_Promo   true
	// Standard_DS5_v2         false
	// Standard_DS5_v2_Promo   false
	// Standard_E16-4as_v4     true
	// Standard_E16-4ds_v4     true
	// Standard_E16-4s_v3      true
	// Standard_E16-4s_v4      true
	// Standard_E16-8as_v4     false
	// Standard_E16-8ds_v4     true
	// Standard_E16-8s_v3      true
	// Standard_E16-8s_v4      true
	// Standard_E16_v3         true
	// Standard_E16_v4         true
	// Standard_E16a_v3        false
	// Standard_E16a_v4        true
	// Standard_E16as_v3       false
	// Standard_E16as_v4       true
	// Standard_E16d_v4        true
	// Standard_E16ds_v4       true
	// Standard_E16s_v3        true
	// Standard_E16s_v4        true
	// Standard_E20_v3         true
	// Standard_E20_v4         true
	// Standard_E20a_v4        true
	// Standard_E20as_v4       true
	// Standard_E20d_v4        true
	// Standard_E20ds_v4       true
	// Standard_E20s_v3        true
	// Standard_E20s_v4        true
	// Standard_E2_v3          false
	// Standard_E2_v4          false
	// Standard_E2a_v3         false
	// Standard_E2a_v4         false
	// Standard_E2as_v3        false
	// Standard_E2as_v4        false
	// Standard_E2d_v4         false
	// Standard_E2ds_v4        false
	// Standard_E2s_v3         false
	// Standard_E2s_v4         false
	// Standard_E32-16as_v4    true
	// Standard_E32-16ds_v4    true
	// Standard_E32-16s_v3     true
	// Standard_E32-16s_v4     true
	// Standard_E32-8as_v4     true
	// Standard_E32-8ds_v4     true
	// Standard_E32-8s_v3      true
	// Standard_E32-8s_v4      true
	// Standard_E32_v3         true
	// Standard_E32_v4         true
	// Standard_E32a_v3        false
	// Standard_E32a_v4        true
	// Standard_E32as_v3       false
	// Standard_E32as_v4       true
	// Standard_E32d_v4        true
	// Standard_E32ds_v4       true
	// Standard_E32s_v3        true
	// Standard_E32s_v4        true
	// Standard_E4-2as_v4      true
	// Standard_E4-2ds_v4      true
	// Standard_E4-2s_v3       true
	// Standard_E4-2s_v4       true
	// Standard_E48_v3         true
	// Standard_E48_v4         true
	// Standard_E48a_v3        false
	// Standard_E48a_v4        true
	// Standard_E48as_v3       false
	// Standard_E48as_v4       true
	// Standard_E48d_v4        true
	// Standard_E48ds_v4       true
	// Standard_E48s_v3        true
	// Standard_E48s_v4        true
	// Standard_E4_v3          true
	// Standard_E4_v4          true
	// Standard_E4a_v3         false
	// Standard_E4a_v4         true
	// Standard_E4as_v3        false
	// Standard_E4as_v4        true
	// Standard_E4d_v4         true
	// Standard_E4ds_v4        true
	// Standard_E4s_v3         true
	// Standard_E4s_v4         true
	// Standard_E64-16as_v4    true
	// Standard_E64-16ds_v4    true
	// Standard_E64-16s_v3     true
	// Standard_E64-16s_v4     true
	// Standard_E64-32as_v4    true
	// Standard_E64-32ds_v4    true
	// Standard_E64-32s_v3     true
	// Standard_E64-32s_v4     true
	// Standard_E64_v3         true
	// Standard_E64_v4         true
	// Standard_E64a_v3        false
	// Standard_E64a_v4        true
	// Standard_E64as_v3       false
	// Standard_E64as_v4       true
	// Standard_E64d_v4        true
	// Standard_E64ds_v4       true
	// Standard_E64i_v3        true
	// Standard_E64is_v3       true
	// Standard_E64s_v3        true
	// Standard_E64s_v4        true
	// Standard_E8-2as_v4      true
	// Standard_E8-2ds_v4      true
	// Standard_E8-2s_v3       true
	// Standard_E8-2s_v4       true
	// Standard_E8-4as_v4      true
	// Standard_E8-4ds_v4      true
	// Standard_E8-4s_v3       true
	// Standard_E8-4s_v4       true
	// Standard_E80ids_v4      true
	// Standard_E80is_v4       true
	// Standard_E8_v3          true
	// Standard_E8_v4          true
	// Standard_E8a_v3         false
	// Standard_E8a_v4         true
	// Standard_E8as_v3        false
	// Standard_E8as_v4        true
	// Standard_E8d_v4         true
	// Standard_E8ds_v4        true
	// Standard_E8s_v3         true
	// Standard_E8s_v4         true
	// Standard_E96-24as_v4    true
	// Standard_E96-48as_v4    true
	// Standard_E96a_v4        true
	// Standard_E96as_v4       true
	// Standard_F1             false
	// Standard_F16            true
	// Standard_F16s           true
	// Standard_F16s_v2        true
	// Standard_F1s            false
	// Standard_F2             true
	// Standard_F2s            true
	// Standard_F2s_v2         false
	// Standard_F32s_v2        true
	// Standard_F4             true
	// Standard_F48s_v2        true
	// Standard_F4s            true
	// Standard_F4s_v2         true
	// Standard_F64s_v2        true
	// Standard_F72s_v2        true
	// Standard_F8             true
	// Standard_F8s            true
	// Standard_F8s_v2         true
	// Standard_G1             false
	// Standard_G2             false
	// Standard_G3             false
	// Standard_G4             false
	// Standard_G5             false
	// Standard_GS1            false
	// Standard_GS2            false
	// Standard_GS3            false
	// Standard_GS4            false
	// Standard_GS4-4          false
	// Standard_GS4-8          false
	// Standard_GS5            false
	// Standard_GS5-16         false
	// Standard_GS5-8          false
	// Standard_H16            false
	// Standard_H16_Promo      false
	// Standard_H16m           false
	// Standard_H16m_Promo     false
	// Standard_H16mr          false
	// Standard_H16mr_Promo    false
	// Standard_H16r           false
	// Standard_H16r_Promo     false
	// Standard_H8             false
	// Standard_H8_Promo       false
	// Standard_H8m            false
	// Standard_H8m_Promo      false
	// Standard_HB120rs_v2     false
	// Standard_HB60rs         false
	// Standard_HC44rs         false
	// Standard_L16s           false
	// Standard_L16s_v2        true
	// Standard_L32s           false
	// Standard_L32s_v2        true
	// Standard_L48s_v2        true
	// Standard_L4s            false
	// Standard_L64s_v2        true
	// Standard_L80s_v2        true
	// Standard_L8s            false
	// Standard_L8s_v2         true
	// Standard_LRS            false
	// Standard_M128           true
	// Standard_M128-32ms      true
	// Standard_M128-64ms      true
	// Standard_M128dms_v2     true
	// Standard_M128ds_v2      true
	// Standard_M128m          true
	// Standard_M128ms         true
	// Standard_M128ms_v2      true
	// Standard_M128s          true
	// Standard_M128s_v2       true
	// Standard_M16-4ms        true
	// Standard_M16-8ms        true
	// Standard_M16ms          true
	// Standard_M192idms_v2    true
	// Standard_M192ids_v2     true
	// Standard_M192ims_v2     true
	// Standard_M192is_v2      true
	// Standard_M192ms_v2      true
	// Standard_M192s_v2       true
	// Standard_M208ms_v2      true
	// Standard_M208s_v2       true
	// Standard_M24ms_v2       true
	// Standard_M24s_v2        true
	// Standard_M32-16ms       true
	// Standard_M32-8ms        true
	// Standard_M32dms_v2      true
	// Standard_M32ls          true
	// Standard_M32ms          true
	// Standard_M32ms_v2       true
	// Standard_M32ts          true
	// Standard_M416-208ms_v2  true
	// Standard_M416-208s_v2   true
	// Standard_M416ms_v2      true
	// Standard_M416s_v2       true
	// Standard_M48ms_v2       true
	// Standard_M48s_v2        true
	// Standard_M64            true
	// Standard_M64-16ms       true
	// Standard_M64-32ms       true
	// Standard_M64dms_v2      true
	// Standard_M64ds_v2       true
	// Standard_M64ls          true
	// Standard_M64m           true
	// Standard_M64ms          true
	// Standard_M64ms_v2       true
	// Standard_M64s           false
	// Standard_M64s_v2        true
	// Standard_M8-2ms         false
	// Standard_M8-4ms         false
	// Standard_M8ms           true
	// Standard_M96ms_v2       true
	// Standard_M96s_v2        true
	// Standard_NC12           false
	// Standard_NC12_Promo     false
	// Standard_NC12s_v2       false
	// Standard_NC12s_v3       false
	// Standard_NC16as_T4_v3   true
	// Standard_NC24           false
	// Standard_NC24_Promo     false
	// Standard_NC24r          false
	// Standard_NC24r_Promo    false
	// Standard_NC24rs_v2      false
	// Standard_NC24rs_v3      false
	// Standard_NC24s_v2       false
	// Standard_NC24s_v3       false
	// Standard_NC4as_T4_v3    true
	// Standard_NC6            false
	// Standard_NC64as_T4_v3   true
	// Standard_NC6_Promo      false
	// Standard_NC6s_v2        false
	// Standard_NC6s_v3        false
	// Standard_NC8as_T4_v3    true
	// Standard_ND12s          false
	// Standard_ND24rs         false
	// Standard_ND24s          false
	// Standard_ND40rs_v2      false
	// Standard_ND40s_v3       true
	// Standard_ND6s           false
	// Standard_NP10s          false
	// Standard_NP20s          false
	// Standard_NP40s          false
	// Standard_NV12           false
	// Standard_NV12_Promo     false
	// Standard_NV12s_v2       false
	// Standard_NV12s_v3       false
	// Standard_NV16as_v4      false
	// Standard_NV24           false
	// Standard_NV24_Promo     false
	// Standard_NV24s_v2       false
	// Standard_NV24s_v3       false
	// Standard_NV32as_v4      false
	// Standard_NV48s_v3       false
	// Standard_NV4as_v4       false
	// Standard_NV6            false
	// Standard_NV6_Promo      false
	// Standard_NV6s_v2        false
	// Standard_NV8as_v4       false
	// Standard_PB12s          false
	// Standard_PB24s          false
	// Standard_PB6s           false
	// Standard_ZRS            false
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
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A1",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A10",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A11",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A1_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A2_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A2m_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A4_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A4m_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A5",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A6",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A7",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A8",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A8_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A8m_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_A9",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B12ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_B16ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_B1ls",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B1ms",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B1s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B20ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_B2ms",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B2s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B4ms",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_B8ms",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D1",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D11",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D11_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D11_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D12",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D12_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D12_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D12_v2_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D13",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D13_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D13_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D14",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D14_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D14_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D15_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D16a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D16as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D16s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D1_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D2_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D2_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2a_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2as_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2d_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2ds_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D2s_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D32_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D32a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D32as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D32s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D3_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D3_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D48_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D48a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D48as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D48s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D4a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D4as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D4s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D5_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D5_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D64a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D64as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D64s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D8a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D8as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D8s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D96a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D96a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_D96as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_D96as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC1s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DC2s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DC2s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC4s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DC4s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC8_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DC8s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS1",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS11",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS11-1_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS11_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS11_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS12-1_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12-2_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS12_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS13-2_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13-4_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS13_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS14-4_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14-8_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS14_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS15_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS1_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS2_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS2_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS3_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS3_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS4_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS4_v2_Promo",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_DS5_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_DS5_v2_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E16-4as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16-4ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16-4s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16-4s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16-8as_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E16-8ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16-8s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16-8s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E16a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E16as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E16s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E20s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E2_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2a_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2as_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2d_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2ds_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E2s_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E32-16as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-16ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-16s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-16s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-8as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-8ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-8s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32-8s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E32a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E32as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E32s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4-2as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4-2ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4-2s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4-2s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E48a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E48as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E48s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E4a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E4as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E4s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-16as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-16ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-16s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-16s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-32as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-32ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-32s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64-32s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E64a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E64as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64i_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64is_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E64s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-2as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-2ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-2s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-2s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-4as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-4ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-4s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8-4s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E80ids_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E80is_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8a_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E8a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8as_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_E8as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8d_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8ds_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E8s_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E96-24as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E96-48as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E96a_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_E96as_v4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F1",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_F16",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F16s",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F16s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F1s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_F2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F2s",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F2s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_F32s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F4",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F48s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F4s",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F4s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F64s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F72s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F8",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F8s",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_F8s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_G1",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_G5",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS1",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS4-4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS4-8",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS5",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS5-16",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_GS5-8",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16m",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16m_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16mr",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16mr_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16r",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H16r_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8m",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_H8m_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_HB120rs_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_HB60rs",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_HC44rs",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L16s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L16s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L32s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L32s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L48s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L4s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L64s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L80s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_L8s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_L8s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_LRS",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M128",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128-32ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128-64ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128dms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128ds_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128m",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128s",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M128s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M16-4ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M16-8ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M16ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192idms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192ids_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192ims_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192is_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M192s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M208ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M208s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M24ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M24s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32-16ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32-8ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32dms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32ls",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M32ts",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M416-208ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M416-208s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M416ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M416s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M48ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M48s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64-16ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64-32ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64dms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64ds_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64ls",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64m",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M64s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M64s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M8-2ms",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M8-4ms",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_M8ms",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M96ms_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_M96s_v2",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_NC12",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC12_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC12s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC12s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC16as_T4_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_NC24",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24r",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24r_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24rs_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24rs_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC24s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC4as_T4_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_NC6",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC64as_T4_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_NC6_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC6s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC6s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NC8as_T4_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_ND12s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND24rs",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND24s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND40rs_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ND40s_v3",
	//     "AcceleratedNetworking": true
	//   },
	//   {
	//     "Name": "Standard_ND6s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NP10s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NP20s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NP40s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV12s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV16as_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV24s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV32as_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV48s_v3",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV4as_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV6",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV6_Promo",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV6s_v2",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_NV8as_v4",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_PB12s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_PB24s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_PB6s",
	//     "AcceleratedNetworking": false
	//   },
	//   {
	//     "Name": "Standard_ZRS",
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
	// 	AcceleratedNetworking bool
	// }

	// var VMSkus = []VMSku{
	// 	{
	// 		Name:                  "Standard_A0",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A1",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A10",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A11",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A1_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A2_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A2m_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A4_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A4m_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A5",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A6",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A7",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A8",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A8_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A8m_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_A9",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B12ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_B16ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_B1ls",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B1ms",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B1s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B20ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_B2ms",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B2s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B4ms",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_B8ms",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D1",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D11",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D11_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D11_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D12_v2_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D13",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D13_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D13_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D14",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D14_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D14_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D15_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D16s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D1_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2a_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2as_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2d_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2ds_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D2s_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D32s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D3_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D3_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48a_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D48s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D4s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D5_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D5_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D64s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D8s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_D96as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC1s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC2s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC2s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC4s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC4s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC8_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DC8s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS1",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11-1_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS11_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12-1_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12-2_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS12_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13-2_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13-4_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS13_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14-4_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14-8_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS14_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS15_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS1_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS2_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS2_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS3_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS3_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS4_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS4_v2_Promo",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS5_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_DS5_v2_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-4ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-4s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-4s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-8ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-8s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16-8s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E16s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E20s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2a_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2as_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2d_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2ds_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E2s_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-16ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-16s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-16s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-8ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-8s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32-8s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E32s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4-2ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4-2s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4-2s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E48s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E4s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-16ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-16s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-16s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-32ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-32s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64-32s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64i_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64is_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E64s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8-2ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 			Name:                  "Standard_E8-2s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8-2s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8-4ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8-4s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8-4s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8a_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8as_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8d_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8ds_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E8s_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E96a_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_E96as_v4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F1",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_F16",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F16s",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F16s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F1s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_F2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F2s",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F2s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_F32s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F4",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F48s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F4s",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F4s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F64s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F72s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F8",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F8s",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_F8s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_G1",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_G5",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS1",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS4-4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS4-8",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS5",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS5-16",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_GS5-8",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16m",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16m_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16mr",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16mr_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16r",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H16r_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8m",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_H8m_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_HB120rs_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_HB60rs",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_HC44rs",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L16s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L16s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L32s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L32s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L48s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L4s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L64s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L80s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_L8s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_L8s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_LRS",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128-32ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128-64ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128m",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M128s",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M16-4ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M16-8ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M16ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M192ms_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M192s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M208ms_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M208s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M24ms_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M24s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32-16ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32-8ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32ls",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M32ts",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M416-208ms_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M416-208s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M416ms_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M416s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M48ms_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M48s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64-16ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64-32ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64ls",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64m",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M64s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M8-2ms",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M8-4ms",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_M8ms",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M96ms_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_M96s_v2",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC12s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24r",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24r_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24rs_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24rs_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC24s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NC6s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND12s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND24rs",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND24s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND40rs_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND40s_v3",
	// 		AcceleratedNetworking: true,
	// 	},
	// 	{
	// 		Name:                  "Standard_ND6s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NP10s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NP20s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NP40s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV12s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV16as_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV24s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV32as_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV48s_v3",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV4as_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV6",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV6_Promo",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV6s_v2",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_NV8as_v4",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_PB12s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_PB24s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_PB6s",
	// 		AcceleratedNetworking: false,
	// 	},
	// 	{
	// 		Name:                  "Standard_ZRS",
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
