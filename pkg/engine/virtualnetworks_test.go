// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"

	"testing"
)

func TestCreateVirtualNetwork(t *testing.T) {

	// Test Master VNet without RouteTable
	cs := &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "cilium",
				},
			},
		},
	}

	vnet := CreateVirtualNetwork(cs)
	expectedVnet := VirtualNetworkARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
			},
		},
		VirtualNetwork: network.VirtualNetwork{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('virtualNetworkName')]"),
			Type:     to.StringPtr("Microsoft.Network/virtualNetworks"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{
						"[parameters('vnetCidr')]",
					},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("[variables('subnetName')]"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
						},
					},
				},
			},
		},
	}

	diff := cmp.Diff(vnet, expectedVnet)

	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}

	//Test MasterVnet with route Table
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "calico",
				},
			},
		},
	}

	vnet = CreateVirtualNetwork(cs)
	expectedVnet.DependsOn = []string{
		"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
		"[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]",
	}

	expectedVnet.Subnets = &[]network.Subnet{
		{
			Name: to.StringPtr("[variables('subnetName')]"),
			SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
				AddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
				NetworkSecurityGroup: &network.SecurityGroup{
					ID: to.StringPtr("[variables('nsgID')]"),
				},
				RouteTable: &network.RouteTable{
					ID: to.StringPtr("[variables('routeTableID')]"),
				},
			},
		},
	}

	diff = cmp.Diff(vnet, expectedVnet)

	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}

	// Test master vnet with ipv6 dual stack feature enabled
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "kubenet",
				},
			},
			FeatureFlags: &api.FeatureFlags{
				EnableIPv6DualStack: true,
			},
		},
	}

	vnet = CreateVirtualNetwork(cs)
	expectedVnet = VirtualNetworkARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
				"[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]",
			},
		},
		VirtualNetwork: network.VirtualNetwork{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('virtualNetworkName')]"),
			Type:     to.StringPtr("Microsoft.Network/virtualNetworks"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{
						"[parameters('vnetCidr')]",
						"[parameters('vnetCidrIPv6')]",
					},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("[variables('subnetName')]"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefixes: &[]string{
								"[parameters('masterSubnet')]",
								"[parameters('masterSubnetIPv6')]",
							},
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
							RouteTable: &network.RouteTable{
								ID: to.StringPtr("[variables('routeTableID')]"),
							},
						},
					},
				},
			},
		},
	}

	diff = cmp.Diff(vnet, expectedVnet)
	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}

	// Test master vnet with appgw-ingress addon
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "azure",
					Addons: []api.KubernetesAddon{
						{
							Name:    common.AppGwIngressAddonName,
							Enabled: to.BoolPtr(true),
							Config: map[string]string{
								"appgw-subnet": "10.0.0.1/16",
							},
						},
					},
				},
			},
		},
	}

	vnet = CreateVirtualNetwork(cs)
	expectedVnet = VirtualNetworkARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
			},
		},
		VirtualNetwork: network.VirtualNetwork{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('virtualNetworkName')]"),
			Type:     to.StringPtr("Microsoft.Network/virtualNetworks"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{
						"[parameters('vnetCidr')]",
					},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("[variables('subnetName')]"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("[variables('appGwSubnetName')]"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("[parameters('appGwSubnet')]"),
						},
					},
				},
			},
		},
	}

	diff = cmp.Diff(vnet, expectedVnet)
	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}
}

func TestCreateVirtualNetworkVMSS(t *testing.T) {

	//Test Create Virtual Network Master VMSS without routeTable
	cs := &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "cilium",
				},
			},
		},
	}

	vnet := createVirtualNetworkVMSS(cs)

	expectedVnet := VirtualNetworkARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
			},
		},
		VirtualNetwork: network.VirtualNetwork{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('virtualNetworkName')]"),
			Type:     to.StringPtr("Microsoft.Network/virtualNetworks"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{
						"[parameters('vnetCidr')]",
					},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("subnetmaster"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("subnetagent"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("[parameters('agentSubnet')]"),
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
						},
					},
				},
			},
		},
	}

	diff := cmp.Diff(vnet, expectedVnet)

	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}

	//Test Create Virtual Network Master VMSS with Route Table
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "calico",
				},
			},
		},
	}

	vnet = createVirtualNetworkVMSS(cs)
	expectedVnet.DependsOn = []string{
		"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
		"[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]",
	}

	expectedVnet.Subnets = &[]network.Subnet{
		{
			Name: to.StringPtr("subnetmaster"),
			SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
				AddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
				NetworkSecurityGroup: &network.SecurityGroup{
					ID: to.StringPtr("[variables('nsgID')]"),
				},
				RouteTable: &network.RouteTable{
					ID: to.StringPtr("[variables('routeTableID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("subnetagent"),
			SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
				AddressPrefix: to.StringPtr("[parameters('agentSubnet')]"),
				NetworkSecurityGroup: &network.SecurityGroup{
					ID: to.StringPtr("[variables('nsgID')]"),
				},
				RouteTable: &network.RouteTable{
					ID: to.StringPtr("[variables('routeTableID')]"),
				},
			},
		},
	}

	diff = cmp.Diff(vnet, expectedVnet)

	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}

	// Test with ipv6 dual stack feature enabled
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "kubenet",
				},
			},
			FeatureFlags: &api.FeatureFlags{
				EnableIPv6DualStack: true,
			},
		},
	}

	vnet = createVirtualNetworkVMSS(cs)
	expectedVnet = VirtualNetworkARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
				"[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]",
			},
		},
		VirtualNetwork: network.VirtualNetwork{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('virtualNetworkName')]"),
			Type:     to.StringPtr("Microsoft.Network/virtualNetworks"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{
						"[parameters('vnetCidr')]",
						"[parameters('vnetCidrIPv6')]",
					},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("subnetmaster"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefixes: &[]string{
								"[parameters('masterSubnet')]",
								"[parameters('masterSubnetIPv6')]",
							},
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
							RouteTable: &network.RouteTable{
								ID: to.StringPtr("[variables('routeTableID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("subnetagent"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("[parameters('agentSubnet')]"),
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
							RouteTable: &network.RouteTable{
								ID: to.StringPtr("[variables('routeTableID')]"),
							},
						},
					},
				},
			},
		},
	}

	diff = cmp.Diff(vnet, expectedVnet)
	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}
}

func TestCreateHostedMasterVirtualNetwork(t *testing.T) {
	// Test Hosted Master VNet without RouteTable
	cs := &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "cilium",
				},
			},
		},
	}

	vnet := createHostedMasterVirtualNetwork(cs)

	expectedVnet := VirtualNetworkARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
				"[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]",
			},
		},
		VirtualNetwork: network.VirtualNetwork{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('virtualNetworkName')]"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{
						"[parameters('vnetCidr')]",
					},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("[variables('subnetName')]"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
							RouteTable: &network.RouteTable{ID: to.StringPtr("[variables('routeTableID')]")},
						},
					},
				},
			},
			Type: to.StringPtr("Microsoft.Network/virtualNetworks"),
		},
	}

	diff := cmp.Diff(vnet, expectedVnet)

	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}

	// test with ipv6 dual stack feature enabled
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPolicy: "kubenet",
				},
			},
			FeatureFlags: &api.FeatureFlags{
				EnableIPv6DualStack: true,
			},
		},
	}

	vnet = createHostedMasterVirtualNetwork(cs)
	expectedVnet = VirtualNetworkARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
				"[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]",
			},
		},
		VirtualNetwork: network.VirtualNetwork{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('virtualNetworkName')]"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{
						"[parameters('vnetCidr')]",
						"[parameters('vnetCidrIPv6')]",
					},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("[variables('subnetName')]"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefixes: &[]string{
								("[parameters('masterSubnet')]"),
								("[parameters('masterSubnetIPv6')]"),
							},
							NetworkSecurityGroup: &network.SecurityGroup{
								ID: to.StringPtr("[variables('nsgID')]"),
							},
							RouteTable: &network.RouteTable{ID: to.StringPtr("[variables('routeTableID')]")},
						},
					},
				},
			},
			Type: to.StringPtr("Microsoft.Network/virtualNetworks"),
		},
	}

	diff = cmp.Diff(vnet, expectedVnet)

	if diff != "" {
		t.Errorf("Unexpected diff while comparing vnets: %s", diff)
	}
}
