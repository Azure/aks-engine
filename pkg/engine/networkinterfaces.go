// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateNetworkInterfaces(cs *api.ContainerService) NetworkInterfaceARM {

	var dependencies []string
	if cs.Properties.MasterProfile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	dependencies = append(dependencies, "[variables('masterLbName')]")

	if cs.Properties.MasterProfile.Count > 1 {
		dependencies = append(dependencies, "[variables('masterInternalLbName')]")
	}

	hasCosmosEtcd := nil != cs.Properties.MasterProfile && to.Bool(cs.Properties.MasterProfile.CosmosEtcd)

	if hasCosmosEtcd {
		dependencies = append(dependencies, "[resourceId('Microsoft.DocumentDB/databaseAccounts/', variables('cosmosAccountName'))]")
	}

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionNetwork')]",
		Copy: map[string]string{
			"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
			"name":  "nicLoopNode",
		},
		DependsOn: dependencies,
	}

	lbBackendAddressPools := []network.BackendAddressPool{
		{
			ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
		},
	}

	if cs.Properties.MasterProfile.Count > 1 {
		internalLbPool := network.BackendAddressPool{
			ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
		}
		lbBackendAddressPools = append(lbBackendAddressPools, internalLbPool)
	}

	loadBalancerIPConfig := network.InterfaceIPConfiguration{
		Name: to.StringPtr("ipconfig1"),
		InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
			LoadBalancerBackendAddressPools: &lbBackendAddressPools,
			LoadBalancerInboundNatRules: &[]network.InboundNatRule{
				{
					ID: to.StringPtr("[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'),copyIndex(variables('masterOffset')))]"),
				},
			},
			PrivateIPAddress:          to.StringPtr("[variables('masterPrivateIpAddrs')[copyIndex(variables('masterOffset'))]]"),
			Primary:                   to.BoolPtr(true),
			PrivateIPAllocationMethod: network.Static,
			Subnet: &network.Subnet{
				ID: to.StringPtr("[variables('vnetSubnetID')]"),
			},
		},
	}

	isAzureCNI := cs.Properties.OrchestratorProfile.IsAzureCNI()

	ipConfigurations := []network.InterfaceIPConfiguration{loadBalancerIPConfig}

	nicProperties := network.InterfacePropertiesFormat{
		IPConfigurations: &ipConfigurations,
	}

	if isAzureCNI {
		ipConfigurations = append(ipConfigurations, getSecondaryNICIPConfigs(cs.Properties.MasterProfile.IPAddressCount)...)
	} else {
		nicProperties.EnableIPForwarding = to.BoolPtr(true)
	}

	if cs.Properties.LinuxProfile.HasCustomNodesDNS() {
		nicProperties.DNSSettings = &network.InterfaceDNSSettings{
			DNSServers: &[]string{
				"[parameters('dnsServer')]",
			},
		}
	}

	if cs.Properties.MasterProfile.IsCustomVNET() {
		nicProperties.NetworkSecurityGroup = &network.SecurityGroup{
			ID: to.StringPtr("[variables('nsgID')]"),
		}
	}

	networkInterface := network.Interface{
		Location:                  to.StringPtr("[variables('location')]"),
		Name:                      to.StringPtr("[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex(variables('masterOffset')))]"),
		InterfacePropertiesFormat: &nicProperties,
		Type:                      to.StringPtr("Microsoft.Network/networkInterfaces"),
	}

	return NetworkInterfaceARM{
		ARMResource: armResource,
		Interface:   networkInterface,
	}
}

func createPrivateClusterNetworkInterface(cs *api.ContainerService) NetworkInterfaceARM {
	var dependencies []string
	if cs.Properties.MasterProfile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	if cs.Properties.MasterProfile.Count > 1 {
		dependencies = append(dependencies, "[variables('masterInternalLbName')]")
	}

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionNetwork')]",
		Copy: map[string]string{
			"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
			"name":  "nicLoopNode",
		},
		DependsOn: dependencies,
	}

	var lbBackendAddressPools []network.BackendAddressPool

	if cs.Properties.MasterProfile.Count > 1 {
		internalLbPool := network.BackendAddressPool{
			ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
		}
		lbBackendAddressPools = append(lbBackendAddressPools, internalLbPool)
	}

	loadBalancerIPConfig := network.InterfaceIPConfiguration{
		Name: to.StringPtr("ipconfig1"),
		InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
			LoadBalancerBackendAddressPools: &lbBackendAddressPools,
			LoadBalancerInboundNatRules:     &[]network.InboundNatRule{},
			PrivateIPAddress:                to.StringPtr("[variables('masterPrivateIpAddrs')[copyIndex(variables('masterOffset'))]]"),
			Primary:                         to.BoolPtr(true),
			PrivateIPAllocationMethod:       network.Static,
			Subnet: &network.Subnet{
				ID: to.StringPtr("[variables('vnetSubnetID')]"),
			},
		},
	}

	ipConfigurations := []network.InterfaceIPConfiguration{loadBalancerIPConfig}

	isAzureCNI := cs.Properties.OrchestratorProfile.IsAzureCNI()

	if isAzureCNI {
		for i := 2; i <= cs.Properties.MasterProfile.IPAddressCount; i++ {
			ipConfig := network.InterfaceIPConfiguration{
				Name: to.StringPtr(fmt.Sprintf("ipconfig%d", i)),
				InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
					Primary:                   to.BoolPtr(false),
					PrivateIPAllocationMethod: network.Dynamic,
					Subnet: &network.Subnet{
						ID: to.StringPtr("[variables('vnetSubnetID')]"),
					},
				},
			}
			ipConfigurations = append(ipConfigurations, ipConfig)
		}
	}

	nicProperties := network.InterfacePropertiesFormat{
		IPConfigurations: &ipConfigurations,
	}

	if !isAzureCNI {
		nicProperties.EnableIPForwarding = to.BoolPtr(true)
	}

	if cs.Properties.LinuxProfile.HasCustomNodesDNS() {
		nicProperties.DNSSettings = &network.InterfaceDNSSettings{
			DNSServers: &[]string{
				"[parameters('dnsServer')]",
			},
		}
	}

	if cs.Properties.MasterProfile.IsCustomVNET() {
		nicProperties.NetworkSecurityGroup = &network.SecurityGroup{
			ID: to.StringPtr("[variables('nsgID')]"),
		}
	}

	networkInterface := network.Interface{
		Location:                  to.StringPtr("[variables('location')]"),
		Name:                      to.StringPtr("[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex(variables('masterOffset')))]"),
		InterfacePropertiesFormat: &nicProperties,
		Type:                      to.StringPtr("Microsoft.Network/networkInterfaces"),
	}

	return NetworkInterfaceARM{
		ARMResource: armResource,
		Interface:   networkInterface,
	}
}

func createJumpboxNetworkInterface(cs *api.ContainerService) NetworkInterfaceARM {
	dependencies := []string{
		"[concat('Microsoft.Network/publicIpAddresses/', variables('jumpboxPublicIpAddressName'))]",
		"[concat('Microsoft.Network/networkSecurityGroups/', variables('jumpboxNetworkSecurityGroupName'))]",
	}

	if !cs.Properties.MasterProfile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionNetwork')]",
		Copy: map[string]string{
			"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
			"name":  "nicLoopNode",
		},
		DependsOn: dependencies,
	}

	nicProperties := network.InterfacePropertiesFormat{
		IPConfigurations: &[]network.InterfaceIPConfiguration{
			{
				Name: to.StringPtr("ipconfig1"),
				InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
					Subnet: &network.Subnet{
						ID: to.StringPtr("[variables('vnetSubnetID')]"),
					},
					Primary:                   to.BoolPtr(true),
					PrivateIPAllocationMethod: network.Dynamic,
					PublicIPAddress: &network.PublicIPAddress{
						ID: to.StringPtr("[resourceId('Microsoft.Network/publicIpAddresses', variables('jumpboxPublicIpAddressName'))]"),
					},
				},
			},
		},
		NetworkSecurityGroup: &network.SecurityGroup{
			ID: to.StringPtr("[resourceId('Microsoft.Network/networkSecurityGroups', variables('jumpboxNetworkSecurityGroupName'))]"),
		},
	}

	networkInterface := network.Interface{
		Location:                  to.StringPtr("[variables('location')]"),
		Name:                      to.StringPtr("[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex(variables('masterOffset')))]"),
		InterfacePropertiesFormat: &nicProperties,
		Type:                      to.StringPtr("Microsoft.Network/networkInterfaces"),
	}

	return NetworkInterfaceARM{
		ARMResource: armResource,
		Interface:   networkInterface,
	}
}

func createAgentVMASNetworkInterface(cs *api.ContainerService, profile *api.AgentPoolProfile) NetworkInterfaceARM {
	isWindows := profile.IsWindows()
	isCustomVNet := cs.Properties.MasterProfile.IsCustomVNET()
	isAzureCNI := cs.Properties.OrchestratorProfile.IsAzureCNI()

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionNetwork')]",
		Copy: map[string]string{
			"count": fmt.Sprintf("[sub(variables('%[1]sCount'), variables('%[1]sOffset'))]", profile.Name),
			"name":  "loop",
		},
	}

	var dependencies []string

	if isCustomVNet {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	armResource.DependsOn = dependencies

	networkInterface := network.Interface{
		Type:     to.StringPtr("Microsoft.Network/networkInterfaces"),
		Name:     to.StringPtr("[concat(variables('" + profile.Name + "VMNamePrefix'), 'nic-', copyIndex(variables('" + profile.Name + "Offset')))]"),
		Location: to.StringPtr("[variables('location')]"),
	}

	networkInterface.InterfacePropertiesFormat = &network.InterfacePropertiesFormat{}

	if isCustomVNet {
		networkInterface.NetworkSecurityGroup = &network.SecurityGroup{
			ID: to.StringPtr("[variables('nsgID')]"),
		}
	}

	if isWindows {
		networkInterface.EnableAcceleratedNetworking = profile.AcceleratedNetworkingEnabledWindows
	} else {
		networkInterface.EnableAcceleratedNetworking = profile.AcceleratedNetworkingEnabled
	}

	var ipConfigurations []network.InterfaceIPConfiguration
	for i := 1; i <= profile.IPAddressCount; i++ {
		ipConfig := network.InterfaceIPConfiguration{
			Name:                                     to.StringPtr(fmt.Sprintf("ipconfig%d", i)),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{},
		}
		if i == 1 {
			ipConfig.Primary = to.BoolPtr(true)
		}
		ipConfig.PrivateIPAllocationMethod = network.Dynamic
		ipConfig.Subnet = &network.Subnet{
			ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
		}
		if !isWindows {
			if profile.Role == "Infra" {
				ipConfig.LoadBalancerBackendAddressPools = &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers', variables('routerLBName')), '/backendAddressPools/backend')]"),
					},
				}
			}
		}
		ipConfigurations = append(ipConfigurations, ipConfig)
	}

	networkInterface.IPConfigurations = &ipConfigurations

	if !isAzureCNI {
		networkInterface.EnableIPForwarding = to.BoolPtr(true)
	}

	return NetworkInterfaceARM{
		ARMResource: armResource,
		Interface:   networkInterface,
	}
}

func getSecondaryNICIPConfigs(n int) []network.InterfaceIPConfiguration {
	var ipConfigurations []network.InterfaceIPConfiguration
	for i := 2; i <= n; i++ {
		ipConfig := network.InterfaceIPConfiguration{
			Name: to.StringPtr(fmt.Sprintf("ipconfig%d", i)),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		}
		ipConfigurations = append(ipConfigurations, ipConfig)
	}
	return ipConfigurations
}
