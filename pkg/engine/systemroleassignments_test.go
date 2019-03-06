// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/google/go-cmp/cmp"

	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestCreateVmasRoleAssignment(t *testing.T) {

	actual := createVMASRoleAssignment()

	expected := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			Copy: map[string]string{
				"count": "[variables('masterCount')]",
				"name":  "vmLoopNode",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Name: to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(), 'vmidentity'))]"),
			Type: to.StringPtr("Microsoft.Authorization/roleAssignments"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr("[variables('readerRoleDefinitionId')]"),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex()), '2017-03-30', 'Full').identity.principalId]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateAgentVmasSysRoleAssignment(t *testing.T) {

	profile := &api.AgentPoolProfile{
		Name: "fooprofile",
	}
	actual := createAgentVMASSysRoleAssignment(profile)

	expected := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				"[concat('Microsoft.Compute/virtualMachines/', variables('fooprofileVMNamePrefix'), copyIndex(variables('fooprofileOffset')))]",
			},
			Copy: map[string]string{
				"count": "[sub(variables('fooprofileCount'), variables('fooprofileOffset'))]",
				"name":  "vmLoopNode",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Name: to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachines/', variables('fooprofileVMNamePrefix'), copyIndex(variables('fooprofileOffset')), 'vmidentity'))]"),
			Type: to.StringPtr("Microsoft.Authorization/roleAssignments"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr("[variables('readerRoleDefinitionId')]"),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachines/', variables('fooprofileVMNamePrefix'), copyIndex(variables('fooprofileOffset'))), '2017-03-30', 'Full').identity.principalId]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

}

func TestCreateAgentVmssSysRoleAssignment(t *testing.T) {

	profile := &api.AgentPoolProfile{
		Name: "fooprofile",
	}
	actual := createAgentVMSSSysRoleAssignment(profile)

	expected := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				"[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('fooprofileVMNamePrefix'))]",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Name: to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('fooprofileVMNamePrefix'), 'vmidentity'))]"),
			Type: to.StringPtr("Microsoft.Authorization/roleAssignments"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr("[variables('readerRoleDefinitionId')]"),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('fooprofileVMNamePrefix'), '2017-03-30', 'Full').identity.principalId]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

}
