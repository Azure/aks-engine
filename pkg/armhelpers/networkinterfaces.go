package armhelpers

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createNetworkInterfaces(cs *api.ContainerService) NetworkInterfaceARM {

	var dependencies []string
	if cs.Properties.MasterProfile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	dependencies = append(dependencies, "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'),copyIndex(variables('masterOffset')))]")

	if cs.Properties.MasterProfile.Count > 1 {
		dependencies = append(dependencies, "[variables('masterInternalLbName')]")
	}

	hasCosmosEtcd := nil != cs.Properties.MasterProfile && to.Bool(cs.Properties.MasterProfile.CosmosEtcd)

	if hasCosmosEtcd {
		dependencies = append(dependencies, "[resourceId('Microsoft.DocumentDB/databaseAccounts/', variables('cosmosAccountName'))]")
	}

	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionNetwork')]",
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

	loadBalancerIpConfig := network.InterfaceIPConfiguration{
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

	ipConfigurations := []network.InterfaceIPConfiguration{loadBalancerIpConfig}

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
		ApiVersion: "[variables('apiVersionNetwork')]",
		Copy: map[string]string{
			"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
			"name":  "nicLoopNode",
		},
		DependsOn: dependencies,
	}

	return NetworkInterfaceARM{
		ARMResource: armResource,
	}
}
