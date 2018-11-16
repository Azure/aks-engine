// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

import (
	"net"
	"regexp"

	"github.com/pkg/errors"
)

// CidrFirstIP returns the first IP of the provided subnet.
func CidrFirstIP(cidr net.IP) net.IP {
	for j := len(cidr) - 1; j >= 0; j-- {
		cidr[j]++
		if cidr[j] > 0 {
			break
		}
	}
	return cidr
}

// CidrStringFirstIP returns the first IP of the provided subnet string. Returns an error
// if the string cannot be parsed.
func CidrStringFirstIP(ip string) (net.IP, error) {
	cidr, _, err := net.ParseCIDR(ip)
	if err != nil {
		return nil, err
	}
	return CidrFirstIP(cidr), nil
}

// IP4BroadcastAddress returns the broadcast address for the given IP subnet.
func IP4BroadcastAddress(n *net.IPNet) net.IP {
	// see https://groups.google.com/d/msg/golang-nuts/IrfXFTUavXE/8YwzIOBwJf0J
	ip4 := n.IP.To4()
	if ip4 == nil {
		return nil
	}
	last := make(net.IP, len(ip4))
	copy(last, ip4)
	for i := range ip4 {
		last[i] |= ^n.Mask[i]
	}
	return last
}

// GetVNETSubnetIDComponents extract subscription, resourcegroup, vnetname, subnetname from the vnetSubnetID
func GetVNETSubnetIDComponents(vnetSubnetID string) (string, string, string, string, error) {
	vnetSubnetIDRegex := `^\/subscriptions\/([^\/]*)\/resourceGroups\/([^\/]*)\/providers\/Microsoft.Network\/virtualNetworks\/([^\/]*)\/subnets\/([^\/]*)$`
	re, err := regexp.Compile(vnetSubnetIDRegex)
	if err != nil {
		return "", "", "", "", err
	}
	submatches := re.FindStringSubmatch(vnetSubnetID)
	if len(submatches) != 5 {
		return "", "", "", "", errors.New("Unable to parse vnetSubnetID. Please use a vnetSubnetID with format /subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/SUBNET_NAME")
	}
	return submatches[1], submatches[2], submatches[3], submatches[4], nil
}
