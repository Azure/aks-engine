// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

// CreatePublicIPAddressForNodePools returns public ipv4 address resource for node pool Load Balancer
func CreatePublicIPAddressForNodePools(name string) PublicIPAddressARM {
	return PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('" + name + "')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}
}

// CreatePublicIPAddressForMaster returns public ipv4 address resource for master Load Balancer
// Includes optional DNS configuration for public clusters
func CreatePublicIPAddressForMaster(includeDNS bool) PublicIPAddressARM {
	var dnsSettings *network.PublicIPAddressDNSSettings

	if includeDNS {
		dnsSettings = &network.PublicIPAddressDNSSettings{
			DomainNameLabel: to.StringPtr("[variables('masterFqdnPrefix')]"),
		}
	}

	return PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('masterPublicIPAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				DNSSettings:              dnsSettings,
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}
}

func createAppGwPublicIPAddress() PublicIPAddressARM {
	return PublicIPAddressARM{
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
}

func createJumpboxPublicIPAddress() PublicIPAddressARM {
	return PublicIPAddressARM{
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
}

// CreateClusterPublicIPAddress returns public ipv4 address resource for cluster
// this public ip address is created and added to the loadbalancer that's created with
// fqdn as name. ARM does not allow creating a loadbalancer with only ipv6 FE which is
// why a ipv4 fe is created here and added to lb.
func CreateClusterPublicIPAddress() PublicIPAddressARM {
	return PublicIPAddressARM{
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
}
