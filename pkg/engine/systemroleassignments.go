// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
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
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "vmLoopNode",
			},
		},
	}

	systemRoleAssignment.Name = to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), 'vmidentity'))]")
	systemRoleAssignment.Type = to.StringPtr("Microsoft.Authorization/roleAssignments")
	systemRoleAssignment.RoleAssignmentPropertiesWithScope = &authorization.RoleAssignmentPropertiesWithScope{
		RoleDefinitionID: to.StringPtr("[variables('contributorRoleDefinitionId')]"),
		PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset'))), '2017-03-30', 'Full').identity.principalId]"),
		PrincipalType:    authorization.ServicePrincipal,
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
		PrincipalType:    authorization.ServicePrincipal,
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
		PrincipalType:    authorization.ServicePrincipal,
	}

	return systemRoleAssignment
}

func createKubernetesMasterRoleAssignmentForAgentPools(masterProfile *api.MasterProfile, agentPoolProfiles []*api.AgentPoolProfile) []DeploymentWithResourceGroupARM {
	var dependenciesToMasterVms = make([]string, masterProfile.Count)

	for masterIdx := 0; masterIdx < masterProfile.Count; masterIdx++ {
		dependenciesToMasterVms[masterIdx] = fmt.Sprintf("[concat(variables('masterVMNamePrefix'), %d)]", masterIdx)
	}

	roleAssignmentsForAllAgentPools := []DeploymentWithResourceGroupARM{}

	// The following is based on:
	// * https://github.com/MicrosoftDocs/azure-docs/blob/master/articles/role-based-access-control/role-assignments-template.md#create-a-role-assignment-at-a-resource-scope
	// * https://github.com/Azure/azure-quickstart-templates/blob/master/201-rbac-builtinrole-multipleVMs/azuredeploy.json#L79

	// We're gonna keep track of distinct VNETs in use across all our node pools;
	//   we only want to define master VM --> VNET role assignments once per VNET.
	// If our cluster configuration includes more than one pool sharing a common VNET,
	//   we define the master VM --> VNET role assignments (one per master VM) just once for those pools
	var vnetInCluster = struct{}{}
	vnets := make(map[string]struct{})
	for _, agentPool := range agentPoolProfiles {
		var roleAssignments = make([]interface{}, masterProfile.Count)
		subnetElements := strings.Split(agentPool.VnetSubnetID, "/")
		// We expect a very specific string format for the VnetSubnetID property;
		//   if it can't be split into at least 9 "/"-delimited elements,
		//   then we should assume that our role assignment composition below will be malformed,
		//   and so we simply skip assigning a role assigment for this pool.
		// This should never happen, but this defensive posture ensures no code panic execution path
		//   // when we statically grab the first 9 elements (`subnetElements[:9]` below)
		if len(subnetElements) < 9 {
			continue
		}
		vnetResourceURI := strings.Join(subnetElements[:9], "/")
		if _, ok := vnets[vnetResourceURI]; ok {
			continue
		}
		vnets[vnetResourceURI] = vnetInCluster

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
				// Reference to the VNET of the worker VMs:
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

		roleAssignmentsForAllAgentPools = append(roleAssignmentsForAllAgentPools, roleAssignmentsForAgentPoolSubDeployment)
	}

	return roleAssignmentsForAllAgentPools
}
