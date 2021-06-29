// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createApplicationGateway(prop *api.Properties) ApplicationGatewayARM {

	applicationGateway := ApplicationGatewayARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/publicIPAddresses/', variables('appGwPublicIPAddressName'))]",
				"[concat('Microsoft.Network/virtualNetworks/', variables('virtualNetworkName'))]",
			},
		},
		ApplicationGateway: network.ApplicationGateway{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('appGwName')]"),
			ApplicationGatewayPropertiesFormat: &network.ApplicationGatewayPropertiesFormat{
				Sku: &network.ApplicationGatewaySku{
					Name:     network.ApplicationGatewaySkuName("[parameters('appGwSku')]"),
					Tier:     network.ApplicationGatewayTier("[parameters('appGwSku')]"),
					Capacity: to.Int32Ptr(2),
				},
				GatewayIPConfigurations: &[]network.ApplicationGatewayIPConfiguration{
					{
						Name: to.StringPtr("gatewayIP"),
						ApplicationGatewayIPConfigurationPropertiesFormat: &network.ApplicationGatewayIPConfigurationPropertiesFormat{
							Subnet: &network.SubResource{
								ID: to.StringPtr("[concat(variables('vnetID'),'/subnets/',variables('appGwSubnetName'))]"),
							},
						},
					},
				},
				FrontendIPConfigurations: &[]network.ApplicationGatewayFrontendIPConfiguration{
					{
						Name: to.StringPtr("frontendIP"),
						ApplicationGatewayFrontendIPConfigurationPropertiesFormat: &network.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
							PublicIPAddress: &network.SubResource{
								ID: to.StringPtr("[resourceId('Microsoft.Network/publicIpAddresses',variables('appGwPublicIPAddressName'))]"),
							},
						},
					},
				},
				FrontendPorts: &[]network.ApplicationGatewayFrontendPort{
					{
						Name: to.StringPtr("httpPort"),
						ApplicationGatewayFrontendPortPropertiesFormat: &network.ApplicationGatewayFrontendPortPropertiesFormat{
							Port: to.Int32Ptr(80),
						},
					},
				},
				BackendAddressPools: &[]network.ApplicationGatewayBackendAddressPool{
					{
						Name: to.StringPtr("pool"),
						ApplicationGatewayBackendAddressPoolPropertiesFormat: &network.ApplicationGatewayBackendAddressPoolPropertiesFormat{
							BackendAddresses: &[]network.ApplicationGatewayBackendAddress{},
						},
					},
				},
				HTTPListeners: &[]network.ApplicationGatewayHTTPListener{
					{
						Name: to.StringPtr("httpListener"),
						ApplicationGatewayHTTPListenerPropertiesFormat: &network.ApplicationGatewayHTTPListenerPropertiesFormat{
							Protocol: network.HTTP,
							FrontendPort: &network.SubResource{
								ID: to.StringPtr("[concat(variables('appGwId'), '/frontendPorts/httpPort')]"),
							},
							FrontendIPConfiguration: &network.SubResource{
								ID: to.StringPtr("[concat(variables('appGwId'), '/frontendIPConfigurations/frontendIP')]"),
							},
						},
					},
				},
				BackendHTTPSettingsCollection: &[]network.ApplicationGatewayBackendHTTPSettings{
					{
						Name: to.StringPtr("setting"),
						ApplicationGatewayBackendHTTPSettingsPropertiesFormat: &network.ApplicationGatewayBackendHTTPSettingsPropertiesFormat{
							Port:     to.Int32Ptr(80),
							Protocol: network.HTTP,
						},
					},
				},
				RequestRoutingRules: &[]network.ApplicationGatewayRequestRoutingRule{
					{
						Name: to.StringPtr("rule"),
						ApplicationGatewayRequestRoutingRulePropertiesFormat: &network.ApplicationGatewayRequestRoutingRulePropertiesFormat{
							HTTPListener: &network.SubResource{
								ID: to.StringPtr("[concat(variables('appGwId'), '/httpListeners/httpListener')]"),
							},
							BackendAddressPool: &network.SubResource{
								ID: to.StringPtr("[concat(variables('appGwId'), '/backendAddressPools/pool')]"),
							},
							BackendHTTPSettings: &network.SubResource{
								ID: to.StringPtr("[concat(variables('appGwId'), '/backendHttpSettingsCollection/setting')]"),
							},
						},
					},
				},
			},
			Type: to.StringPtr("Microsoft.Network/applicationGateways"),
		},
	}

	if prop.OrchestratorProfile.KubernetesConfig.GetAddonByName(common.AppGwIngressAddonName).Config["appgw-sku"] == "WAF_v2" {
		applicationGateway.ApplicationGateway.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration = &network.ApplicationGatewayWebApplicationFirewallConfiguration{
			Enabled:      to.BoolPtr(true),
			FirewallMode: network.Detection,
		}
	}

	privateIP := prop.OrchestratorProfile.KubernetesConfig.GetAddonByName(common.AppGwIngressAddonName).Config["appgw-private-ip"]
	if privateIP != "" {
		frontendIPConfigurations := append(
			*applicationGateway.ApplicationGateway.ApplicationGatewayPropertiesFormat.FrontendIPConfigurations,
			network.ApplicationGatewayFrontendIPConfiguration{
				Name: to.StringPtr("privateIp"),
				ApplicationGatewayFrontendIPConfigurationPropertiesFormat: &network.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
					PrivateIPAddress: to.StringPtr(privateIP),
				},
			})
		applicationGateway.ApplicationGateway.ApplicationGatewayPropertiesFormat.FrontendIPConfigurations = &frontendIPConfigurations
	}

	return applicationGateway
}
