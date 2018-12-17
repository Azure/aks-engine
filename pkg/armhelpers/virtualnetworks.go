package armhelpers

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createVirtualNetwork(cs *api.ContainerService) VirtualNetworkARM {

	dependencies := []string{
		"[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]",
	}

	requireRouteTable := cs.Properties.OrchestratorProfile.RequireRouteTable()
	if requireRouteTable {
		dependencies = append(dependencies, "[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]")
	}

	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionNetwork')]",
		DependsOn:  dependencies,
	}

	subnet := network.Subnet{
		Name: to.StringPtr("[variables('subnetName')]"),
		SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
			AddressPrefix: to.StringPtr("[parameters('masterSubnet')]"),
			NetworkSecurityGroup: &network.SecurityGroup{
				ID: to.StringPtr("[variables('nsgID')]"),
			},
		},
	}

	if requireRouteTable {
		subnet.RouteTable = &network.RouteTable{
			ID: to.StringPtr("[variables('routeTableID')]"),
		}
	}

	virtualNetwork := network.VirtualNetwork{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[variables('virtualNetworkName')]"),
		Type:     to.StringPtr("Microsoft.Network/virtualNetworks"),
		VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
			Subnets: &[]network.Subnet{
				subnet,
			},
		},
	}

	return VirtualNetworkARM{
		ARMResource:    armResource,
		VirtualNetwork: virtualNetwork,
	}
}
