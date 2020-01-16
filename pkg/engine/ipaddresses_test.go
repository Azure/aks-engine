// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestCreatePublicIPAddress(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('masterPublicIPAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				DNSSettings: &network.PublicIPAddressDNSSettings{
					DomainNameLabel: to.StringPtr("[variables('masterFqdnPrefix')]"),
				},
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}
	isForMaster := true
	includeDNS := true
	actual := CreatePublicIPAddress(isForMaster, includeDNS)

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
	// Testing CreatePublicIPAddress when it's not for master

	expected = PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('agentPublicIPAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}
	isForMaster = false
	includeDNS = false
	actual = CreatePublicIPAddress(isForMaster, includeDNS)

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Testing CreatePublicIPAddress without DNS

	expected = PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('masterPublicIPAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}
	isForMaster = true
	includeDNS = false
	actual = CreatePublicIPAddress(isForMaster, includeDNS)

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateJumpboxPublicIPAddress(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('jumpboxPublicIpAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				DNSSettings: &network.PublicIPAddressDNSSettings{
					DomainNameLabel: to.StringPtr("[variables('masterFqdnPrefix')]"),
				},
				PublicIPAllocationMethod: network.Dynamic,
			},
			Sku: &network.PublicIPAddressSku{
				Name: network.PublicIPAddressSkuNameBasic,
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}

	actual := createJumpboxPublicIPAddress()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateClusterPublicIPAddress(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("fee-ipv4"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}

	actual := CreateClusterPublicIPAddress()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateAppGwPublicIPAddress(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('appGwPublicIPAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "Standard",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}

	actual := createAppGwPublicIPAddress()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}
