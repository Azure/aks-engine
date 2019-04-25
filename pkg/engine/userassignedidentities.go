// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"
	"github.com/Azure/go-autorest/autorest/to"
)

func createUserAssignedIdentities() UserAssignedIdentitiesARM {
	return UserAssignedIdentitiesARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionManagedIdentity')]",
		},
		Identity: msi.Identity{
			Type:     "Microsoft.ManagedIdentity/userAssignedIdentities",
			Name:     to.StringPtr("[variables('userAssignedID')]"),
			Location: to.StringPtr("[variables('location')]"),
		},
	}
}
