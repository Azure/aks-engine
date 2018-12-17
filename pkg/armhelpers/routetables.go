package armhelpers

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createRouteTable(cs *api.ContainerService) RouteTableARM {
	return RouteTableARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionNetwork')]",
		},
		RouteTable: network.RouteTable{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("variables('routeTableName')"),
			Type:     to.StringPtr("Microsoft.Network/routeTables"),
		},
	}
}
