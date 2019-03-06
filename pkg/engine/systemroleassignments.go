// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/go-autorest/autorest/to"
)

type SystemRoleAssignmentARM struct {
	ARMResource
	authorization.RoleAssignment
}

func createVMASRoleAssignment() SystemRoleAssignmentARM {
	systemRoleAssignment := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			Copy: map[string]string{
				"count": "[variables('masterCount')]",
				"name":  "vmLoopNode",
			},
		},
	}

	systemRoleAssignment.Name = to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(), 'vmidentity'))]")
	systemRoleAssignment.Type = to.StringPtr("Microsoft.Authorization/roleAssignments")
	systemRoleAssignment.RoleAssignmentPropertiesWithScope = &authorization.RoleAssignmentPropertiesWithScope{
		RoleDefinitionID: to.StringPtr("[variables('readerRoleDefinitionId')]"),
		PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex()), '2017-03-30', 'Full').identity.principalId]"),
	}
	return systemRoleAssignment
}

func createAgentVMASSysRoleAssignment(profile *api.AgentPoolProfile) SystemRoleAssignmentARM {
	systemRoleAssignment := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				fmt.Sprintf("[concat('Microsoft.Compute/virtualMachines/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name),
			},
			Copy: map[string]string{
				"count": fmt.Sprintf("[sub(variables('%[1]sCount'), variables('%[1]sOffset'))]", profile.Name),
				"name":  "vmLoopNode",
			},
		},
	}

	systemRoleAssignment.Name = to.StringPtr(fmt.Sprintf("[guid(concat('Microsoft.Compute/virtualMachines/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')), 'vmidentity'))]", profile.Name))
	systemRoleAssignment.Type = to.StringPtr("Microsoft.Authorization/roleAssignments")

	systemRoleAssignment.RoleAssignmentPropertiesWithScope = &authorization.RoleAssignmentPropertiesWithScope{
		RoleDefinitionID: to.StringPtr("[variables('readerRoleDefinitionId')]"),
		PrincipalID:      to.StringPtr(fmt.Sprintf("[reference(concat('Microsoft.Compute/virtualMachines/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset'))), '2017-03-30', 'Full').identity.principalId]", profile.Name)),
	}

	return systemRoleAssignment
}

func createAgentVMSSSysRoleAssignment(profile *api.AgentPoolProfile) SystemRoleAssignmentARM {
	systemRoleAssignment := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				fmt.Sprintf("[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('%[1]sVMNamePrefix'))]", profile.Name),
			},
		},
	}
	systemRoleAssignment.RoleAssignmentPropertiesWithScope = &authorization.RoleAssignmentPropertiesWithScope{}
	systemRoleAssignment.Name = to.StringPtr(fmt.Sprintf("[guid(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('%[1]sVMNamePrefix'), 'vmidentity'))]", profile.Name))
	systemRoleAssignment.Type = to.StringPtr("Microsoft.Authorization/roleAssignments")
	systemRoleAssignment.RoleAssignmentPropertiesWithScope = &authorization.RoleAssignmentPropertiesWithScope{
		RoleDefinitionID: to.StringPtr("[variables('readerRoleDefinitionId')]"),
		PrincipalID:      to.StringPtr(fmt.Sprintf("[reference(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('%[1]sVMNamePrefix'), '2017-03-30', 'Full').identity.principalId]", profile.Name)),
	}

	return systemRoleAssignment
}
