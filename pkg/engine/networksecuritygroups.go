// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateNetworkSecurityGroup(cs *api.ContainerService) NetworkSecurityGroupARM {
	armResource := ARMResource{
		APIVersion: "[variables('apiVersionNetwork')]",
	}

	sshRule := network.SecurityRule{
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
	}

	kubeTLSRule := network.SecurityRule{
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
	}

	securityRules := []network.SecurityRule{
		sshRule,
		kubeTLSRule,
	}

	if cs.Properties.HasWindows() {
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

		securityRules = append(securityRules, rdpRule)
	}

	if cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") {
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

		securityRules = append(securityRules, vnetRule)
		securityRules = append(securityRules, blockOutBoundRule)
	}

	nsg := network.SecurityGroup{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[variables('nsgName')]"),
		Type:     to.StringPtr("Microsoft.Network/networkSecurityGroups"),
		SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
			SecurityRules: &securityRules,
		},
	}

	return NetworkSecurityGroupARM{
		ARMResource:   armResource,
		SecurityGroup: nsg,
	}
}

func createJumpboxNSG() NetworkSecurityGroupARM {
	armResource := ARMResource{
		APIVersion: "[variables('apiVersionNetwork')]",
	}

	securityRules := []network.SecurityRule{
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
	}
	nsg := network.SecurityGroup{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[variables('nsgName')]"),
		Type:     to.StringPtr("Microsoft.Network/networkSecurityGroups"),
		SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
			SecurityRules: &securityRules,
		},
	}
	return NetworkSecurityGroupARM{
		ARMResource:   armResource,
		SecurityGroup: nsg,
	}
}

func createHostedMasterNSG() NetworkSecurityGroupARM {
	armResource := ARMResource{
		APIVersion: "[variables('apiVersionNetwork')]",
	}
	nsg := network.SecurityGroup{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[variables('nsgName')]"),
		Type:     to.StringPtr("Microsoft.Network/networkSecurityGroups"),
		SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
			SecurityRules: &[]network.SecurityRule{},
		},
	}

	return NetworkSecurityGroupARM{
		ARMResource:   armResource,
		SecurityGroup: nsg,
	}
}
