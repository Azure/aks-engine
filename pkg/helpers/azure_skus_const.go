// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

// GetKubernetesAllowedVMSKUs provides the allowed sizes for Kubernetes agent VMs.
//
// Code generated for package helpers by aks-engine DO NOT EDIT. (@generated)
//
// To generate this code, run the command:
//   aks-engine get-skus --output=code

type VMSku struct {
	Name                  string
	AcceleratedNetworking bool
}

var VMSkus = []VMSku{
	{
		Name:                  "Standard_A0",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A1",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A10",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A11",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A1_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A2_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A2m_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A4_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A4m_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A5",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A6",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A7",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A8",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A8_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A8m_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_A9",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_B12ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_B16ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_B1ls",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_B1ms",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_B1s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_B20ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_B2ms",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_B2s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_B4ms",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_B8ms",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D1",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D11",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D11_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D11_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D12",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D12_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D12_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D12_v2_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D13",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D13_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D13_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D14",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D14_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D14_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D15_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D16a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D16as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D16s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D1_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D2_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D2_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D2a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2a_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D2as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2as_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D2d_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D2ds_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D2s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2s_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D2s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D32_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D32a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D32as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D32s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D3_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D3_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D48_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D48a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D48as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D48s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D4a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D4as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D4s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D5_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D5_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D64a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D64as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D64s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D8a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D8as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D8s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D96a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_D96as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_D96s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC16ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC16as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC16ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC16s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC1ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC1s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DC1s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC24ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC24s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC2ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC2as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC2ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC2s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DC2s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC2s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC32ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC32as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC32ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC32s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC48ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC48as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC48ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC48s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC4ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC4as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC4ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC4s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DC4s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC4s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC64ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC64as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC8_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC8ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC8as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC8ds_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC8s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DC8s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC96ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DC96as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS1",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS11",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS11-1_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS11_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS11_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS12",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS12-1_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS12-2_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS12_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS12_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS13",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS13-2_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS13-4_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS13_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS13_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS14",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS14-4_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS14-8_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS14_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS14_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS15_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS1_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS2_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS2_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS3_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS3_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS4_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS4_v2_Promo",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_DS5_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_DS5_v2_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E104i_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E104id_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E104ids_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E104is_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E112iads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E112ias_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-4s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-8ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-8as_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E16-8as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-8ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-8ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-8s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-8s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16-8s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E16a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E16as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16bds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16bs_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E16s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E20s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2a_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2as_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2bds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2bs_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2d_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2ds_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E2s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2s_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E2s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-16s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32-8s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E32a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E32as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32bds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32bs_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E32s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4-2s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E48a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E48as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48bds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48bs_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E48s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E4a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E4as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4bds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4bs_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E4s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-16s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64-32s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E64a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E64as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64bds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64bs_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64i_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64is_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E64s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-2s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8-4s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E80ids_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E80is_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8a_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E8a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8as_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_E8as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8bds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8bs_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8d_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8ds_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8s_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E8s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-24ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-24as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-24as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-24ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-24s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-48ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-48as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-48as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-48ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96-48s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96a_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96as_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96d_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96ds_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96ias_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_E96s_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC16ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC16as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC20ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC20as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC2ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC2as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC32ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC32as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC48ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC48as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC4ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC4as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC64ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC64as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC8ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC8as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC96ads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC96as_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC96iads_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_EC96ias_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F1",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_F16",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F16s",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F16s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F1s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_F2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F2s",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F2s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_F32s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F48s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F4s",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F4s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F64s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F72s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F8",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F8s",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_F8s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_FX12mds",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_FX24mds",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_FX36mds",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_FX48mds",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_FX4mds",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_G1",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_G2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_G3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_G4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_G5",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS1",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS4-4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS4-8",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS5",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS5-16",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_GS5-8",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16m",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16m_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16mr",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16mr_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16r",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H16r_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H8",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H8_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H8m",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_H8m_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_HB120-16rs_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB120-32rs_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB120-64rs_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB120-96rs_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB120rs_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_HB120rs_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB60-15rs",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB60-30rs",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB60-45rs",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HB60rs",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_HC44-16rs",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HC44-32rs",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_HC44rs",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_L16s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_L16s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_L32s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_L32s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_L48s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_L4s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_L64s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_L80s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_L8s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_L8s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_LRS",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_M128",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128-32ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128-64ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128dms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128ds_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128m",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128s",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M128s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M16-4ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M16-8ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M16ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M192idms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M192ids_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M192ims_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M192is_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M192ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M192s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M208ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M208s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M24ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M24s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M32-16ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M32-8ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M32dms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M32ls",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M32ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M32ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M32ts",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M416-208ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M416-208s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M416ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M416s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M48ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M48s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64-16ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64-32ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64dms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64ds_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64ls",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64m",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M64s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_M64s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M8-2ms",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_M8-4ms",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_M832ixs",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M8ms",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M96ms_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_M96s_v2",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NC12",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC12_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC12s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC12s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC16ads_A10_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NC16as_T4_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NC24",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC24_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC24r",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC24r_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC24rs_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC24rs_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC24s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC24s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC32ads_A10_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NC4as_T4_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NC6",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC64as_T4_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NC6_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC6s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC6s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NC8ads_A10_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NC8as_T4_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_ND12s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_ND24rs",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_ND24s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_ND40rs_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_ND40s_v3",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_ND6s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_ND96amsr_A100_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_ND96asr_v4",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NP10s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NP20s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NP40s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV12",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV12_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV12ads_A10_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NV12s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV12s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV16as_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV18ads_A10_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NV24",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV24_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV24s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV24s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV32as_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV36adms_A10_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NV36ads_A10_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NV48s_v3",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV4as_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV6",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV6_Promo",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV6ads_A10_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NV6s_v2",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_NV72ads_A10_v5",
		AcceleratedNetworking: true,
	},
	{
		Name:                  "Standard_NV8as_v4",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_PB12s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_PB24s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_PB6s",
		AcceleratedNetworking: false,
	},
	{
		Name:                  "Standard_ZRS",
		AcceleratedNetworking: false,
	},
}
