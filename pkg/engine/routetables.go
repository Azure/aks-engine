// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createRouteTable() RouteTableARM {
	return RouteTableARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		RouteTable: network.RouteTable{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('routeTableName')]"),
			Type:     to.StringPtr("Microsoft.Network/routeTables"),
		},
	}
}
