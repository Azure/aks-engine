// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/go-autorest/autorest/to"
)

type IdentityRoleDefinition string

const (
	// IdentityContributorRole means created user assigned identity will have "Contributor" role in created resource group
	IdentityContributorRole IdentityRoleDefinition = "[variables('contributorRoleDefinitionId')]"
	// IdentityReaderRole means created user assigned identity will have "Reader" role in created resource group
	IdentityReaderRole IdentityRoleDefinition = "[variables('readerRoleDefinitionId')]"
	// IdentityManagedIdentityOperatorRole means created user assigned identity or service principal will have operator access on a different managed identity
	IdentityManagedIdentityOperatorRole IdentityRoleDefinition = "[variables('managedIdentityOperatorRoleDefinitionId')]"
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

// createKubernetesSpAppGIdentityOperatorAccessRoleAssignment gives identity operator access on AGIC Identity to the cluster identity
func createKubernetesSpAppGIdentityOperatorAccessRoleAssignment(prop *api.Properties) RoleAssignmentARM {
	kubernetesSpObjectID := ""
	// determine objectId of the cluster identity used by the kubernetes cluster
	if prop.OrchestratorProfile != nil &&
		prop.OrchestratorProfile.KubernetesConfig != nil &&
		prop.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
		kubernetesSpObjectID = "[reference(concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))).principalId]"
	} else if prop.ServicePrincipalProfile.ObjectID != "" {
		kubernetesSpObjectID = prop.ServicePrincipalProfile.ObjectID
	}

	return RoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/applicationgateways/', variables('appGwName'))]",
				"[concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('appGwICIdentityName'))]",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Type: to.StringPtr("Microsoft.ManagedIdentity/userAssignedIdentities/providers/roleAssignments"),
			Name: to.StringPtr("[concat(variables('appGwICIdentityName'), '/Microsoft.Authorization/', guid(resourceGroup().id, 'aksidentityaccess'))]"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr(string(IdentityManagedIdentityOperatorRole)),
				PrincipalID:      to.StringPtr(kubernetesSpObjectID),
				PrincipalType:    authorization.ServicePrincipal,
				Scope:            to.StringPtr("[variables('appGwICIdentityId')]"),
			},
		},
	}
}

// createAppGwIdentityResourceGroupReadSysRoleAssignment gives read access to Resource Group for Identity used by AGIC
func createAppGwIdentityResourceGroupReadSysRoleAssignment() RoleAssignmentARM {
	return RoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/applicationgateways/', variables('appGwName'))]",
				"[concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('appGwICIdentityName'))]",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Type: to.StringPtr("Microsoft.Authorization/roleAssignments"),
			Name: to.StringPtr("[guid(resourceGroup().id, 'identityrgaccess')]"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr(string(IdentityReaderRole)),
				PrincipalID:      to.StringPtr("[reference(variables('appGwICIdentityId'), variables('apiVersionManagedIdentity')).principalId]"),
				Scope:            to.StringPtr("[resourceGroup().id]"),
			},
		},
	}
}

// createAppGwIdentityApplicationGatewayWriteSysRoleAssignment gives write access to Application Gateway for Identity used by AGIC
func createAppGwIdentityApplicationGatewayWriteSysRoleAssignment() RoleAssignmentARM {
	return RoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				"[concat('Microsoft.Network/applicationgateways/', variables('appGwName'))]",
				"[concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('appGwICIdentityName'))]",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Type: to.StringPtr("Microsoft.Network/applicationgateways/providers/roleAssignments"),
			Name: to.StringPtr("[concat(variables('appGwName'), '/Microsoft.Authorization/', guid(resourceGroup().id, 'identityappgwaccess'))]"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr(string(IdentityContributorRole)),
				PrincipalID:      to.StringPtr("[reference(variables('appGwICIdentityId'), variables('apiVersionManagedIdentity')).principalId]"),
				Scope:            to.StringPtr("[variables('appGwId')]"),
			},
		},
	}
}
