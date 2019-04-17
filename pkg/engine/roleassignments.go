// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/go-autorest/autorest/to"
)

type IdentityRoleDefinition string

const (
	// IdentityContributorRole means created user assigned identity will have "Contributor" role in created resource group
	IdentityContributorRole IdentityRoleDefinition = "[variables('contributorRoleDefinitionId')]"
	// IdentityReaderRole means created user assigned identity will have "Reader" role in created resource group
	IdentityReaderRole IdentityRoleDefinition = "[variables('readerRoleDefinitionId')]"
)

func createMSIRoleAssignment(identityRoleDefinition IdentityRoleDefinition) RoleAssignmentARM {
	return RoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationUser')]",
			DependsOn: []string{
				"[concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))]",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Type: to.StringPtr("Microsoft.Authorization/roleAssignments"),
			Name: to.StringPtr("[guid(concat(variables('userAssignedID'), 'roleAssignment', resourceGroup().id))]"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr(string(identityRoleDefinition)),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))).principalId]"),
				PrincipalType:    authorization.ServicePrincipal,
				Scope:            to.StringPtr("[resourceGroup().id]"),
			},
		},
	}
}
