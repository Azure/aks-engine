// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-01-01-preview/authorization"
	"github.com/Azure/go-autorest/autorest/to"
)

func createVMASRoleAssignment() SystemRoleAssignmentARM {
	systemRoleAssignment := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]",
			},
			Copy: map[string]string{
				"count": "[variables('masterCount')]",
				"name":  "vmLoopNode",
			},
		},
	}

	systemRoleAssignment.Name = to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), 'vmidentity'))]")
	systemRoleAssignment.Type = to.StringPtr("Microsoft.Authorization/roleAssignments")
	systemRoleAssignment.RoleAssignmentPropertiesWithScope = &authorization.RoleAssignmentPropertiesWithScope{
		RoleDefinitionID: to.StringPtr("[variables('contributorRoleDefinitionId')]"),
		PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset'))), '2017-03-30', 'Full').identity.principalId]"),
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
		PrincipalID:      to.StringPtr(fmt.Sprintf("[reference(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('%[1]sVMNamePrefix')), '2017-03-30', 'Full').identity.principalId]", profile.Name)),
	}

	return systemRoleAssignment
}

func createKubernetesMasterRoleAssignmentForAgentPools(masterProfile *api.MasterProfile, agentPoolProfiles []*api.AgentPoolProfile) []DeploymentWithResourceGroupARM {
	var dependenciesToMasterVms = make([]string, masterProfile.Count)

	for masterIdx := 0; masterIdx < masterProfile.Count; masterIdx++ {
		dependenciesToMasterVms[masterIdx] = fmt.Sprintf("[concat(variables('masterVMNamePrefix'), %d)]", masterIdx)
	}

	var roleAssignmentsForAllAgentPools = make([]DeploymentWithResourceGroupARM, len(agentPoolProfiles))

	// The following is based on:
	// * https://github.com/MicrosoftDocs/azure-docs/blob/master/articles/role-based-access-control/role-assignments-template.md#create-a-role-assignment-at-a-resource-scope
	// * https://github.com/Azure/azure-quickstart-templates/blob/master/201-rbac-builtinrole-multipleVMs/azuredeploy.json#L79
	for agentPoolIdx, agentPool := range agentPoolProfiles {
		var roleAssignments = make([]interface{}, masterProfile.Count)

		for masterIdx := 0; masterIdx < masterProfile.Count; masterIdx++ {
			masterVMReference := fmt.Sprintf("reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), %d)), '2017-03-30', 'Full').identity.principalId", masterIdx)

			assignment := SystemRoleAssignmentARM{
				ARMResource: ARMResource{
					APIVersion: "[variables('apiVersionAuthorizationSystem')]",
					// TODO: Use `copy` in ARM template instead of Go loops:
					/*
						Copy: map[string]string{
							"count": "[variables('masterCount')]",
							"name":  "vmLoopNode",
						},
					*/
				},
				// Reference to the subnet of the worker VMs:
				RoleAssignment: authorization.RoleAssignment{
					Name: to.StringPtr(fmt.Sprintf("[concat(variables('%sVnet'), '/Microsoft.Authorization/', guid(uniqueString(%s)))]", agentPool.Name, masterVMReference)),
					Type: to.StringPtr("Microsoft.Network/virtualNetworks/providers/roleAssignments"),
					RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
						// Built-in role `network contributor`:
						RoleDefinitionID: to.StringPtr("[variables('networkContributorRoleDefinitionId')]"),
						// The MSI of the master VM:
						PrincipalID: to.StringPtr(fmt.Sprintf("[%s]", masterVMReference)),
					},
				},
			}

			roleAssignments[masterIdx] = assignment
		}

		roleAssignmentsForAgentPoolSubDeployment := DeploymentWithResourceGroupARM{
			DeploymentARMResource: DeploymentARMResource{
				APIVersion: "2017-05-10",
				DependsOn:  dependenciesToMasterVms,
			},
			ResourceGroup: to.StringPtr(fmt.Sprintf("[variables('%sSubnetResourceGroup')]", agentPool.Name)),
			DeploymentExtended: resources.DeploymentExtended{
				Name: to.StringPtr(fmt.Sprintf("[concat('masterMsiRoleAssignment-', variables('%sVMNamePrefix'))]", agentPool.Name)),
				Type: to.StringPtr("Microsoft.Resources/deployments"),
				Properties: &resources.DeploymentPropertiesExtended{
					Mode: "Incremental",
					Template: map[string]interface{}{
						"resources":      roleAssignments,
						"contentVersion": "1.0.0.0",
						"$schema":        "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
					},
				},
			},
		}

		roleAssignmentsForAllAgentPools[agentPoolIdx] = roleAssignmentsForAgentPoolSubDeployment
	}

	return roleAssignmentsForAllAgentPools
}
