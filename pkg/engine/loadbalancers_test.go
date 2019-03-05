// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"

	"github.com/google/go-cmp/cmp"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestCreateLoadBalancer(t *testing.T) {
	actual := CreateLoadBalancer(1)

	expected := LoadBalancerARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
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
				InboundNatRules: &[]network.InboundNatRule{
					{
						Name: to.StringPtr("[concat('SSH-', variables('masterVMNamePrefix'), 0)]"),
						InboundNatRulePropertiesFormat: &network.InboundNatRulePropertiesFormat{
							FrontendIPConfiguration: &network.SubResource{
								ID: to.StringPtr("[variables('masterLbIPConfigID')]"),
							},
							Protocol:         network.TransportProtocol("Tcp"),
							FrontendPort:     to.Int32Ptr(22),
							BackendPort:      to.Int32Ptr(22),
							EnableFloatingIP: to.BoolPtr(false),
						},
					},
				},
				Probes: &[]network.Probe{
					{
						Name: to.StringPtr("tcpHTTPSProbe"),
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

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

}

func TestCreateMasterInternalLoadBalancer(t *testing.T) {
	// Test without custom VNET
	cs := &api.ContainerService{
		Properties: &api.Properties{
			MasterProfile: &api.MasterProfile{},
		},
	}

	actual := CreateMasterInternalLoadBalancer(cs)

	expected := LoadBalancerARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[variables('vnetID')]",
			},
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

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

	// Test with custom Vnet
	cs = &api.ContainerService{
		Properties: &api.Properties{
			MasterProfile: &api.MasterProfile{
				VnetSubnetID: "fooSubnet",
			},
		},
	}

	actual = CreateMasterInternalLoadBalancer(cs)

	expected.DependsOn = []string{
		"[variables('nsgID')]",
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

	// Test with VMSS
	cs = &api.ContainerService{
		Properties: &api.Properties{
			MasterProfile: &api.MasterProfile{
				VnetSubnetID:        "fooSubnet",
				AvailabilityProfile: api.VirtualMachineScaleSets,
			},
		},
	}

	actual = CreateMasterInternalLoadBalancer(cs)

	expected.FrontendIPConfigurations = &[]network.FrontendIPConfiguration{
		{
			Name: to.StringPtr("[variables('masterInternalLbIPConfigName')]"),
			FrontendIPConfigurationPropertiesFormat: &network.FrontendIPConfigurationPropertiesFormat{
				PrivateIPAddress:          to.StringPtr("[variables('kubernetesAPIServerIP')]"),
				PrivateIPAllocationMethod: network.Static,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetIDMaster')]"),
				},
			},
		},
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected error while comparing availability sets: %s", diff)
	}

}
