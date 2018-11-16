// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

import (
	"net"
	"testing"
)

type cidrTest struct {
	cidr     string
	expected string
}

type vnetSubnetIDTest struct {
	vnetSubnetID   string
	expectedSubID  string
	expectedRG     string
	expectedVnet   string
	expectedSubnet string
}

func Test_CidrFirstIP(t *testing.T) {
	scenarios := []cidrTest{
		{
			cidr:     "10.0.0.0/16",
			expected: "10.0.0.1",
		},
		{
			cidr:     "10.16.32.32/27",
			expected: "10.16.32.33",
		},
	}

	for _, scenario := range scenarios {
		if first, _ := CidrStringFirstIP(scenario.cidr); first.String() != scenario.expected {
			t.Errorf("expected first ip of subnet %v to be %v but was %v", scenario.cidr, scenario.expected, first)
		}
	}
}

func Test_IP4BroadcastAddress(t *testing.T) {
	scenarios := []cidrTest{
		{
			cidr:     "10.0.0.0/16",
			expected: "10.0.255.255",
		},
		{
			cidr:     "10.16.32.32/27",
			expected: "10.16.32.63",
		},
	}

	for _, scenario := range scenarios {
		_, cidr, _ := net.ParseCIDR(scenario.cidr)
		if broadcast := IP4BroadcastAddress(cidr); broadcast.String() != scenario.expected {
			t.Errorf("expected broadcast ip of subnet %v to be %v but was %v", scenario.cidr, scenario.expected, broadcast)
		}
	}
}

func Test_GetVNETSubnetIDComponents(t *testing.T) {
	scenarios := []vnetSubnetIDTest{
		{
			vnetSubnetID:   "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/SUBNET_NAME",
			expectedSubID:  "SUB_ID",
			expectedRG:     "RG_NAME",
			expectedVnet:   "VNET_NAME",
			expectedSubnet: "SUBNET_NAME",
		},
		{
			vnetSubnetID:   "/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/SUBNET_NAME",
			expectedSubID:  "",
			expectedRG:     "",
			expectedVnet:   "",
			expectedSubnet: "",
		},
		{
			vnetSubnetID:   "badVnetSubnetID",
			expectedSubID:  "",
			expectedRG:     "",
			expectedVnet:   "",
			expectedSubnet: "",
		},
	}

	for _, scenario := range scenarios {
		subID, rg, vnet, subnet, _ := GetVNETSubnetIDComponents(scenario.vnetSubnetID)
		if subID != scenario.expectedSubID || rg != scenario.expectedRG || vnet != scenario.expectedVnet || subnet != scenario.expectedSubnet {
			t.Errorf("expected subID %s, rg %s, vnet %s and subnet %s but instead got subID %s, rg %s, vnet %s and subnet %s", scenario.expectedSubID, scenario.expectedRG, scenario.expectedVnet, scenario.expectedSubnet, subID, rg, vnet, subnet)
		}
	}
}
