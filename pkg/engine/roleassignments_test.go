// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"
)

func TestCreateMSIRoleAssignment(t *testing.T) {
	// Test create Contributor role assignment
	actual := createMSIRoleAssignment(IdentityContributorRole)
	expected := RoleAssignmentARM{
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
				RoleDefinitionID: to.StringPtr("[variables('contributorRoleDefinitionId')]"),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))).principalId]"),
				PrincipalType:    authorization.ServicePrincipal,
				Scope:            to.StringPtr("[resourceGroup().id]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test create Reader role assignment
	actual = createMSIRoleAssignment(IdentityReaderRole)
	expected = RoleAssignmentARM{
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
				RoleDefinitionID: to.StringPtr("[variables('readerRoleDefinitionId')]"),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))).principalId]"),
				PrincipalType:    authorization.ServicePrincipal,
				Scope:            to.StringPtr("[resourceGroup().id]"),
			},
		},
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateKubernetesSpAppGIdentityOperatorAccessRoleAssignment(t *testing.T) {
	// using service principal
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ObjectID: "xxxx",
			},
		},
	}

	actual := createKubernetesSpAppGIdentityOperatorAccessRoleAssignment(cs.Properties)
	expected := RoleAssignmentARM{
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
				PrincipalID:      to.StringPtr("xxxx"),
				PrincipalType:    authorization.ServicePrincipal,
				Scope:            to.StringPtr("[variables('appGwICIdentityId')]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// using managed identity
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					UseManagedIdentity: true,
				},
			},
		},
	}

	actual = createKubernetesSpAppGIdentityOperatorAccessRoleAssignment(cs.Properties)
	expected = RoleAssignmentARM{
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
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.ManagedIdentity/userAssignedIdentities/', variables('userAssignedID'))).principalId]"),
				PrincipalType:    authorization.ServicePrincipal,
				Scope:            to.StringPtr("[variables('appGwICIdentityId')]"),
			},
		},
	}

	diff = cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateAppGwIdentityResourceGroupReadSysRoleAssignment(t *testing.T) {
	actual := createAppGwIdentityResourceGroupReadSysRoleAssignment()
	expected := RoleAssignmentARM{
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

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateAppGwIdentityApplicationGatewayWriteSysRoleAssignment(t *testing.T) {
	actual := createAppGwIdentityApplicationGatewayWriteSysRoleAssignment()
	expected := RoleAssignmentARM{
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

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}
