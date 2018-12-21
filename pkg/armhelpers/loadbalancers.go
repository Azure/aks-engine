package armhelpers

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createLoadBalancer() LoadBalancerARM {
	return LoadBalancerARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]",
			},
		},
		LoadBalancer: network.LoadBalancer{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('masterLbName')]"),
			LoadBalancerPropertiesFormat: &network.LoadBalancerPropertiesFormat{
				BackendAddressPools: &[]network.BackendAddressPool{
					{
						Name: to.StringPtr("[variables('masterLbBackendPoolName')]"),
					},
				},
				FrontendIPConfigurations: &[]network.FrontendIPConfiguration{
					{
						Name: to.StringPtr("[variables('masterLbIPConfigName')]"),
						FrontendIPConfigurationPropertiesFormat: &network.FrontendIPConfigurationPropertiesFormat{
							PublicIPAddress: &network.PublicIPAddress{
								ID: to.StringPtr("[resourceId('Microsoft.Network/publicIPAddresses',variables('masterPublicIPAddressName'))]"),
							},
						},
					},
				},
				LoadBalancingRules: &[]network.LoadBalancingRule{
					{
						Name: to.StringPtr("LBRuleHTTPS"),
						LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
							FrontendIPConfiguration: &network.SubResource{
								ID: to.StringPtr("[variables('masterLbIPConfigID')]"),
							},
							BackendAddressPool: &network.SubResource{
								ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
							},
							Protocol:             network.TransportProtocolTCP,
							FrontendPort:         to.Int32Ptr(443),
							BackendPort:          to.Int32Ptr(443),
							EnableFloatingIP:     to.BoolPtr(false),
							IdleTimeoutInMinutes: to.Int32Ptr(5),
							LoadDistribution:     network.Default,
							Probe: &network.SubResource{
								ID: to.StringPtr("[concat(variables('masterLbID'),'/probes/tcpHTTPSProbe')]"),
							},
						},
					},
				},
				Probes: &[]network.Probe{
					{
						Name: to.StringPtr("LBRuleHTTPS"),
						ProbePropertiesFormat: &network.ProbePropertiesFormat{
							Protocol:          network.ProbeProtocolTCP,
							Port:              to.Int32Ptr(443),
							IntervalInSeconds: to.Int32Ptr(5),
							NumberOfProbes:    to.Int32Ptr(2),
						},
					},
				},
			},
			Sku: &network.LoadBalancerSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/loadBalancers"),
		},
	}
}

func createMasterInternalLoadBalancer(cs *api.ContainerService) LoadBalancerARM {

	var dependencies []string
	if cs.Properties.MasterProfile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	return LoadBalancerARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionNetwork')]",
			DependsOn:  dependencies,
		},
		LoadBalancer: network.LoadBalancer{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('masterInternalLbName')]"),
			LoadBalancerPropertiesFormat: &network.LoadBalancerPropertiesFormat{
				BackendAddressPools: &[]network.BackendAddressPool{
					{
						Name: to.StringPtr("[variables('masterLbBackendPoolName')]"),
					},
				},
				FrontendIPConfigurations: &[]network.FrontendIPConfiguration{
					{
						Name: to.StringPtr("[variables('masterInternalLbIPConfigName')]"),
						FrontendIPConfigurationPropertiesFormat: &network.FrontendIPConfigurationPropertiesFormat{
							PrivateIPAddress:          to.StringPtr("[variables('kubernetesAPIServerIP')]"),
							PrivateIPAllocationMethod: network.Static,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
				},
				LoadBalancingRules: &[]network.LoadBalancingRule{
					{
						Name: to.StringPtr("InternalLBRuleHTTPS"),
						LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
							BackendAddressPool: &network.SubResource{
								ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
							},
							BackendPort:      to.Int32Ptr(4443),
							EnableFloatingIP: to.BoolPtr(false),
							FrontendIPConfiguration: &network.SubResource{
								ID: to.StringPtr("[variables('masterInternalLbIPConfigID')]"),
							},
							FrontendPort:         to.Int32Ptr(443),
							IdleTimeoutInMinutes: to.Int32Ptr(5),
							Protocol:             network.TransportProtocolTCP,
							Probe: &network.SubResource{
								ID: to.StringPtr("[concat(variables('masterInternalLbID'),'/probes/tcpHTTPSProbe')]"),
							},
						},
					},
				},
				Probes: &[]network.Probe{
					{
						Name: to.StringPtr("tcpHTTPSProbe"),
						ProbePropertiesFormat: &network.ProbePropertiesFormat{
							IntervalInSeconds: to.Int32Ptr(5),
							NumberOfProbes:    to.Int32Ptr(2),
							Port:              to.Int32Ptr(4443),
							Protocol:          network.ProbeProtocolTCP,
						},
					},
				},
			},
			Sku: &network.LoadBalancerSku{
				Name: network.LoadBalancerSkuName("[variables('loadBalancerSku')]"),
			},
			Type: to.StringPtr("Microsoft.Network/loadBalancers"),
		},
	}
}
