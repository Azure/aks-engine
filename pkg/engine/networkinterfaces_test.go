// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"

	"github.com/Azure/aks-engine/pkg/api"
)

func TestCreateNIC(t *testing.T) {

	// Test Master NIC
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
				IPAddressCount:      5,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin: "azure",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
		},
	}

	nic := CreateMasterVMNetworkInterfaces(cs)

	expected := NetworkInterfaceARM{

		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "nicLoopNode",
			},
			DependsOn: []string{
				"[variables('vnetID')]",
				"[variables('masterLbName')]",
			},
		},
		Interface: network.Interface{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex(variables('masterOffset')))]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &[]network.InterfaceIPConfiguration{
					{
						Name: to.StringPtr("ipconfig1"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
								{
									ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
								},
							},
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
					},
					{
						Name: to.StringPtr("ipconfig2"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("ipconfig3"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("ipconfig4"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("ipconfig5"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
				},
			},
			Type: to.StringPtr("Microsoft.Network/networkInterfaces"),
		},
	}

	diff := cmp.Diff(nic, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test Master NIC with custom Vnet
	cs.Properties.MasterProfile.VnetSubnetID = "fooSubnet"

	expected.DependsOn = []string{
		"[variables('nsgID')]",
		"[variables('masterLbName')]",
	}

	expected.NetworkSecurityGroup = &network.SecurityGroup{
		ID: to.StringPtr("[variables('nsgID')]"),
	}

	nic = CreateMasterVMNetworkInterfaces(cs)

	diff = cmp.Diff(nic, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test Master NIC with MultiMaster

	cs.Properties.MasterProfile.Count = 3

	nic = CreateMasterVMNetworkInterfaces(cs)

	expected.DependsOn = []string{
		"[variables('nsgID')]",
		"[variables('masterInternalLbName')]",
		"[variables('masterLbName')]",
	}

	expected.IPConfigurations = &[]network.InterfaceIPConfiguration{
		{
			Name: to.StringPtr("ipconfig1"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
					},
					{
						ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
					},
				},
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
		},
		{
			Name: to.StringPtr("ipconfig2"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig3"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig4"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig5"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
	}

	diff = cmp.Diff(nic, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test Master NIC with Cosmos etcd

	cs.Properties.MasterProfile.CosmosEtcd = to.BoolPtr(true)
	cs.Properties.MasterProfile.Count = 3

	nic = CreateMasterVMNetworkInterfaces(cs)
	expected.DependsOn = []string{
		"[variables('nsgID')]",
		"[variables('masterInternalLbName')]",
		"[resourceId('Microsoft.DocumentDB/databaseAccounts/', variables('cosmosAccountName'))]",
		"[variables('masterLbName')]",
	}
	diff = cmp.Diff(nic, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test Master NIC without AzureCNI and customNodes DNS

	cs.Properties.MasterProfile.IPAddressCount = 5
	cs.Properties.LinuxProfile = &api.LinuxProfile{
		CustomNodesDNS: &api.CustomNodesDNS{
			DNSServer: "barServer",
		},
	}
	nic = CreateMasterVMNetworkInterfaces(cs)
	expected.Interface.DNSSettings = &network.InterfaceDNSSettings{
		DNSServers: &[]string{
			"[parameters('dnsServer')]",
		},
	}
	diff = cmp.Diff(nic, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreatePrivateClusterNetworkInterface(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
				IPAddressCount:      5,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin: "azure",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
		},
	}

	actual := createPrivateClusterMasterVMNetworkInterface(cs)

	expected := NetworkInterfaceARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "nicLoopNode",
			},
			DependsOn: []string{
				"[variables('vnetID')]",
			},
		},
		Interface: network.Interface{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex(variables('masterOffset')))]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &[]network.InterfaceIPConfiguration{
					{
						Name: to.StringPtr("ipconfig1"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							PrivateIPAddress:          to.StringPtr("[variables('masterPrivateIpAddrs')[copyIndex(variables('masterOffset'))]]"),
							Primary:                   to.BoolPtr(true),
							PrivateIPAllocationMethod: network.Static,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("ipconfig2"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("ipconfig3"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("ipconfig4"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
					{
						Name: to.StringPtr("ipconfig5"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Primary:                   to.BoolPtr(false),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
				},
			},
			Type: to.StringPtr("Microsoft.Network/networkInterfaces"),
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test private cluster NIC customVnet and multiple masters

	cs.Properties.MasterProfile.VnetSubnetID = "fooSubnet"
	cs.Properties.MasterProfile.Count = 3

	actual = createPrivateClusterMasterVMNetworkInterface(cs)

	expected.DependsOn = []string{
		"[variables('nsgID')]",
		"[variables('masterInternalLbName')]",
	}

	expected.Interface.NetworkSecurityGroup = &network.SecurityGroup{
		ID: to.StringPtr("[variables('nsgID')]"),
	}

	expected.IPConfigurations = &[]network.InterfaceIPConfiguration{
		{
			Name: to.StringPtr("ipconfig1"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
					},
				},
				LoadBalancerInboundNatRules: &[]network.InboundNatRule{},
				PrivateIPAddress:            to.StringPtr("[variables('masterPrivateIpAddrs')[copyIndex(variables('masterOffset'))]]"),
				Primary:                     to.BoolPtr(true),
				PrivateIPAllocationMethod:   network.Static,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig2"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig3"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig4"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig5"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				Primary:                   to.BoolPtr(false),
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test Master NIC without AzureCNI and customNodes DNS

	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = "notazure"
	cs.Properties.LinuxProfile = &api.LinuxProfile{
		CustomNodesDNS: &api.CustomNodesDNS{
			DNSServer: "barServer",
		},
	}
	actual = createPrivateClusterMasterVMNetworkInterface(cs)
	expected.EnableIPForwarding = to.BoolPtr(true)
	expected.IPConfigurations = &[]network.InterfaceIPConfiguration{
		{
			Name: to.StringPtr("ipconfig1"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
					},
				},
				LoadBalancerInboundNatRules: &[]network.InboundNatRule{},
				PrivateIPAddress:            to.StringPtr("[variables('masterPrivateIpAddrs')[copyIndex(variables('masterOffset'))]]"),
				Primary:                     to.BoolPtr(true),
				PrivateIPAllocationMethod:   network.Static,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('vnetSubnetID')]"),
				},
			},
		},
	}

	expected.Interface.InterfacePropertiesFormat.DNSSettings = &network.InterfaceDNSSettings{
		DNSServers: &[]string{
			"[parameters('dnsServer')]",
		},
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateJumpboxNIC(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin: "azure",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
		},
	}

	actual := createJumpboxNetworkInterface(cs)

	expected := NetworkInterfaceARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/publicIpAddresses/', variables('jumpboxPublicIpAddressName'))]",
				"[concat('Microsoft.Network/networkSecurityGroups/', variables('jumpboxNetworkSecurityGroupName'))]",
				"[variables('vnetID')]",
			},
		},
		Interface: network.Interface{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('jumpboxNetworkInterfaceName')]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
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
			},
			Type: to.StringPtr("Microsoft.Network/networkInterfaces"),
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}
func TestCreateAgentVMASNICWithSLB(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin:   "azure",
					LoadBalancerSku: StandardLoadBalancerSku,
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
		},
	}

	profile := &api.AgentPoolProfile{
		Name:           "fooAgent",
		OSType:         "Linux",
		IPAddressCount: 1,
	}

	// Test AgentVMAS NIC with Standard LB, should add dependsOn for agentLbID and adds agentLbBackendPoolName as backendaddress pool

	actual := createAgentVMASNetworkInterface(cs, profile)

	expected := NetworkInterfaceARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('fooAgentCount'), variables('fooAgentOffset'))]",
				"name":  "loop",
			},
			DependsOn: []string{
				"[variables('vnetID')]",
				"[variables('agentLbID')]",
			},
		},
		Interface: network.Interface{
			Type:     to.StringPtr("Microsoft.Network/networkInterfaces"),
			Name:     to.StringPtr("[concat(variables('fooAgentVMNamePrefix'), 'nic-', copyIndex(variables('fooAgentOffset')))]"),
			Location: to.StringPtr("[variables('location')]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &[]network.InterfaceIPConfiguration{
					{
						Name: to.StringPtr("ipconfig1"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
								{
									ID: to.StringPtr("[concat(variables('agentLbID'), '/backendAddressPools/', variables('agentLbBackendPoolName'))]"),
								},
							},
							Primary:                   to.BoolPtr(true),
							PrivateIPAllocationMethod: network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
							},
						},
					},
				},
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateAgentVMASNICWithSLBHostedMaster(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			HostedMasterProfile: &api.HostedMasterProfile{
				FQDN: "foo",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin:   "azure",
					LoadBalancerSku: StandardLoadBalancerSku,
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
		},
	}

	profile := &api.AgentPoolProfile{
		Name:           "fooAgent",
		OSType:         "Linux",
		IPAddressCount: 1,
	}

	// Test AgentVMAS NIC with Standard LB
	actual := createAgentVMASNetworkInterface(cs, profile)

	expected := NetworkInterfaceARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('fooAgentCount'), variables('fooAgentOffset'))]",
				"name":  "loop",
			},
			DependsOn: []string{
				"[variables('vnetID')]",
			},
		},
		Interface: network.Interface{
			Type:     to.StringPtr("Microsoft.Network/networkInterfaces"),
			Name:     to.StringPtr("[concat(variables('fooAgentVMNamePrefix'), 'nic-', copyIndex(variables('fooAgentOffset')))]"),
			Location: to.StringPtr("[variables('location')]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &[]network.InterfaceIPConfiguration{
					{
						Name: to.StringPtr("ipconfig1"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{},
							Primary:                         to.BoolPtr(true),
							PrivateIPAllocationMethod:       network.Dynamic,
							Subnet: &network.Subnet{
								ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
							},
						},
					},
				},
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateAgentVMASNIC(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin: "azure",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
		},
	}

	profile := &api.AgentPoolProfile{
		Name:                              "fooAgent",
		OSType:                            "Linux",
		Role:                              "Infra",
		LoadBalancerBackendAddressPoolIDs: []string{"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Network/loadBalancers/mySLB/backendAddressPools/mySLBBEPool"},
	}

	actual := createAgentVMASNetworkInterface(cs, profile)

	var ipConfigurations []network.InterfaceIPConfiguration

	expected := NetworkInterfaceARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('fooAgentCount'), variables('fooAgentOffset'))]",
				"name":  "loop",
			},
			DependsOn: []string{
				"[variables('vnetID')]",
			},
		},
		Interface: network.Interface{
			Type:     to.StringPtr("Microsoft.Network/networkInterfaces"),
			Name:     to.StringPtr("[concat(variables('fooAgentVMNamePrefix'), 'nic-', copyIndex(variables('fooAgentOffset')))]"),
			Location: to.StringPtr("[variables('location')]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &ipConfigurations,
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test AgentVMAS NIC with Custom Vnet
	profile.VnetSubnetID = "fooSubnet"

	actual = createAgentVMASNetworkInterface(cs, profile)

	expected.DependsOn = []string{
		"[variables('nsgID')]",
	}

	expected.NetworkSecurityGroup = &network.SecurityGroup{
		ID: to.StringPtr("[variables('nsgID')]"),
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test AgentVMAS NIC with Custom Vnet and multipleIPAddresses
	profile.IPAddressCount = 5

	actual = createAgentVMASNetworkInterface(cs, profile)

	expected.IPConfigurations = &[]network.InterfaceIPConfiguration{
		{
			Name: to.StringPtr("ipconfig1"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers', variables('routerLBName')), '/backendAddressPools/backend')]"),
					},
				},
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('fooAgentVnetSubnetID')]"),
				},
				Primary: to.BoolPtr(true),
			},
		},
		{
			Name: to.StringPtr("ipconfig2"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers', variables('routerLBName')), '/backendAddressPools/backend')]"),
					},
				},
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('fooAgentVnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig3"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers', variables('routerLBName')), '/backendAddressPools/backend')]"),
					},
				},
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('fooAgentVnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig4"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers', variables('routerLBName')), '/backendAddressPools/backend')]"),
					},
				},
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('fooAgentVnetSubnetID')]"),
				},
			},
		},
		{
			Name: to.StringPtr("ipconfig5"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers', variables('routerLBName')), '/backendAddressPools/backend')]"),
					},
				},
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('fooAgentVnetSubnetID')]"),
				},
			},
		},
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateNICWithIPv6DualStackFeature(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
				IPAddressCount:      5,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.15.0-beta.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin: "kubenet",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
			FeatureFlags: &api.FeatureFlags{
				EnableIPv6DualStack: true,
			},
		},
	}

	nic := CreateMasterVMNetworkInterfaces(cs)
	expected := NetworkInterfaceARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "nicLoopNode",
			},
			DependsOn: []string{
				"[variables('vnetID')]",
				"[variables('masterLbName')]",
			},
		},
		Interface: network.Interface{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex(variables('masterOffset')))]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &[]network.InterfaceIPConfiguration{
					{
						Name: to.StringPtr("ipconfig1"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
								{
									ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
								},
							},
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
					},
					{
						Name: to.StringPtr("ipconfigv6"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							PrivateIPAddressVersion: "IPv6",
							Primary:                 to.BoolPtr(false),
							Subnet: &network.Subnet{
								ID: to.StringPtr("[variables('vnetSubnetID')]"),
							},
						},
					},
				},
			},
			Type: to.StringPtr("Microsoft.Network/networkInterfaces"),
		},
	}
	expected.EnableIPForwarding = to.BoolPtr(true)
	diff := cmp.Diff(nic, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateAgentVMASNICWithIPv6DualStackFeature(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:               1,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: api.VirtualMachineScaleSets,
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType:    api.Kubernetes,
				OrchestratorVersion: "1.15.0-beta.2",
				KubernetesConfig: &api.KubernetesConfig{
					NetworkPlugin: "kubenet",
				},
			},
			LinuxProfile: &api.LinuxProfile{},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: api.AvailabilitySet,
				},
			},
			FeatureFlags: &api.FeatureFlags{
				EnableIPv6DualStack: true,
			},
		},
	}

	profile := &api.AgentPoolProfile{
		Name:                              "fooAgent",
		OSType:                            "Linux",
		Role:                              "Infra",
		LoadBalancerBackendAddressPoolIDs: []string{"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Network/loadBalancers/mySLB/backendAddressPools/mySLBBEPool"},
		IPAddressCount:                    1,
	}

	actual := createAgentVMASNetworkInterface(cs, profile)

	var ipConfigurations []network.InterfaceIPConfiguration

	expected := NetworkInterfaceARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('fooAgentCount'), variables('fooAgentOffset'))]",
				"name":  "loop",
			},
			DependsOn: []string{
				"[variables('vnetID')]",
			},
		},
		Interface: network.Interface{
			Type:     to.StringPtr("Microsoft.Network/networkInterfaces"),
			Name:     to.StringPtr("[concat(variables('fooAgentVMNamePrefix'), 'nic-', copyIndex(variables('fooAgentOffset')))]"),
			Location: to.StringPtr("[variables('location')]"),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &ipConfigurations,
			},
		},
	}
	expected.IPConfigurations = &[]network.InterfaceIPConfiguration{
		{
			Name: to.StringPtr("ipconfig1"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers', variables('routerLBName')), '/backendAddressPools/backend')]"),
					},
					{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers',parameters('masterEndpointDNSNamePrefix')), '/backendAddressPools/', parameters('masterEndpointDNSNamePrefix'))]"),
					},
				},
				PrivateIPAllocationMethod: network.Dynamic,
				Subnet: &network.Subnet{
					ID: to.StringPtr("[variables('fooAgentVnetSubnetID')]"),
				},
				Primary: to.BoolPtr(true),
			},
		},
		{
			Name: to.StringPtr("ipconfigv6"),
			InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
				PrivateIPAddressVersion: "IPv6",
				Primary:                 to.BoolPtr(false),
				Subnet: &network.Subnet{
					ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
				},
			},
		},
	}
	expected.EnableIPForwarding = to.BoolPtr(true)

	diff := cmp.Diff(actual, expected)
	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}
