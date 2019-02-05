// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"

	"testing"

	"github.com/Azure/aks-engine/pkg/api"
)

func TestCreateNetworkSecurityGroup(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:   "fooAgent",
					OSType: "Linux",
				},
			},
			FeatureFlags: &api.FeatureFlags{},
		},
	}

	// Test create normal nsg

	actual := CreateNetworkSecurityGroup(cs)

	expected := NetworkSecurityGroupARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		SecurityGroup: network.SecurityGroup{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('nsgName')]"),
			Type:     to.StringPtr("Microsoft.Network/networkSecurityGroups"),
			SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
				SecurityRules: &[]network.SecurityRule{
					{
						Name: to.StringPtr("allow_ssh"),
						SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
							Access:                   network.SecurityRuleAccessAllow,
							Description:              to.StringPtr("Allow SSH traffic to master"),
							DestinationAddressPrefix: to.StringPtr("*"),
							DestinationPortRange:     to.StringPtr("22-22"),
							Direction:                network.SecurityRuleDirectionInbound,
							Priority:                 to.Int32Ptr(101),
							Protocol:                 network.SecurityRuleProtocolTCP,
							SourceAddressPrefix:      to.StringPtr("*"),
							SourcePortRange:          to.StringPtr("*"),
						},
					},
					{
						Name: to.StringPtr("allow_kube_tls"),
						SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
							Access:                   network.SecurityRuleAccessAllow,
							Description:              to.StringPtr("Allow kube-apiserver (tls) traffic to master"),
							DestinationAddressPrefix: to.StringPtr("*"),
							DestinationPortRange:     to.StringPtr("443-443"),
							Direction:                network.SecurityRuleDirectionInbound,
							Priority:                 to.Int32Ptr(100),
							Protocol:                 network.SecurityRuleProtocolTCP,
							SourceAddressPrefix:      to.StringPtr("*"),
							SourcePortRange:          to.StringPtr("*"),
						},
					},
				},
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing nsgs : %s", diff)
	}

	// Test Create NSG with windows and Block Outbound internet

	cs.Properties.AgentPoolProfiles = []*api.AgentPoolProfile{
		{
			Name:   "fooAgent",
			OSType: "Windows",
		},
	}

	cs.Properties.FeatureFlags.BlockOutboundInternet = true

	actual = CreateNetworkSecurityGroup(cs)

	rules := *expected.SecurityRules

	rdpRule := network.SecurityRule{
		Name: to.StringPtr("allow_rdp"),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Access:                   network.SecurityRuleAccessAllow,
			Description:              to.StringPtr("Allow RDP traffic to master"),
			DestinationAddressPrefix: to.StringPtr("*"),
			DestinationPortRange:     to.StringPtr("3389-3389"),
			Direction:                network.SecurityRuleDirectionInbound,
			Priority:                 to.Int32Ptr(102),
			Protocol:                 network.SecurityRuleProtocolTCP,
			SourceAddressPrefix:      to.StringPtr("*"),
			SourcePortRange:          to.StringPtr("*"),
		},
	}

	vnetRule := network.SecurityRule{
		Name: to.StringPtr("allow_vnet"),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Access:                   network.SecurityRuleAccessAllow,
			Description:              to.StringPtr("Allow outbound internet to vnet"),
			DestinationAddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
			DestinationPortRange:     to.StringPtr("*"),
			Direction:                network.SecurityRuleDirectionOutbound,
			Priority:                 to.Int32Ptr(110),
			Protocol:                 network.SecurityRuleProtocolAsterisk,
			SourceAddressPrefix:      to.StringPtr("VirtualNetwork"),
			SourcePortRange:          to.StringPtr("*"),
		},
	}

	blockOutBoundRule := network.SecurityRule{
		Name: to.StringPtr("block_outbound"),
		SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
			Access:                   network.SecurityRuleAccessDeny,
			Description:              to.StringPtr("Block outbound internet from master"),
			DestinationAddressPrefix: to.StringPtr("*"),
			DestinationPortRange:     to.StringPtr("*"),
			Direction:                network.SecurityRuleDirectionOutbound,
			Priority:                 to.Int32Ptr(120),
			Protocol:                 network.SecurityRuleProtocolAsterisk,
			SourceAddressPrefix:      to.StringPtr("*"),
			SourcePortRange:          to.StringPtr("*"),
		},
	}

	rules = append(rules, rdpRule, vnetRule, blockOutBoundRule)

	expected.SecurityRules = &rules

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing nsgs : %s", diff)
	}
}

func TestCreateJumpboxNSG(t *testing.T) {
	expected := NetworkSecurityGroupARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		SecurityGroup: network.SecurityGroup{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('nsgName')]"),
			Type:     to.StringPtr("Microsoft.Network/networkSecurityGroups"),
			SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
				SecurityRules: &[]network.SecurityRule{
					{
						Name: to.StringPtr("default-allow-ssh"),
						SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
							Priority:                 to.Int32Ptr(1000),
							Protocol:                 network.SecurityRuleProtocolTCP,
							Access:                   network.SecurityRuleAccessAllow,
							Direction:                network.SecurityRuleDirectionInbound,
							SourceAddressPrefix:      to.StringPtr("*"),
							SourcePortRange:          to.StringPtr("*"),
							DestinationAddressPrefix: to.StringPtr("*"),
							DestinationPortRange:     to.StringPtr("22"),
						},
					},
				},
			},
		},
	}

	actual := createJumpboxNSG()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing nsgs : %s", diff)
	}
}

func TestCreateHostedMasterNSG(t *testing.T) {
	expected := NetworkSecurityGroupARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		SecurityGroup: network.SecurityGroup{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('nsgName')]"),
			Type:     to.StringPtr("Microsoft.Network/networkSecurityGroups"),
			SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
				SecurityRules: &[]network.SecurityRule{},
			},
		},
	}

	actual := createHostedMasterNSG()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing nsgs : %s", diff)
	}
}
