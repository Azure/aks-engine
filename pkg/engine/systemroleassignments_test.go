// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/google/go-cmp/cmp"

	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-01-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestCreateVmasRoleAssignment(t *testing.T) {

	actual := createVMASRoleAssignment()

	expected := SystemRoleAssignmentARM{
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
		RoleAssignment: authorization.RoleAssignment{
			Name: to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), 'vmidentity'))]"),
			Type: to.StringPtr("Microsoft.Authorization/roleAssignments"),
			RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr("[variables('contributorRoleDefinitionId')]"),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset'))), '2017-03-30', 'Full').identity.principalId]"),
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
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('fooprofileVMNamePrefix')), '2017-03-30', 'Full').identity.principalId]"),
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

}

func TestCreateKubernetesMasterRoleAssignmentForAgentPools(t *testing.T) {
	masterProfile := &api.MasterProfile{
		Count: 2,
	}
	agentProfile1 := &api.AgentPoolProfile{
		Name: "agentProfile1",
	}
	agentProfile2 := &api.AgentPoolProfile{
		Name: "agentProfile2",
	}

	actual := createKubernetesMasterRoleAssignmentForAgentPools(masterProfile, []*api.AgentPoolProfile{agentProfile1, agentProfile2})

	expected := []DeploymentWithResourceGroupARM{
		{
			DeploymentARMResource: DeploymentARMResource{
				APIVersion: "2017-05-10",
				DependsOn: []string{
					"[concat(variables('masterVMNamePrefix'), 0)]",
					"[concat(variables('masterVMNamePrefix'), 1)]",
				},
			},
			ResourceGroup: to.StringPtr("[variables('agentProfile1SubnetResourceGroup')]"),
			DeploymentExtended: resources.DeploymentExtended{
				Name: to.StringPtr("[concat('masterMsiRoleAssignment-', variables('agentProfile1VMNamePrefix'))]"),
				Type: to.StringPtr("Microsoft.Resources/deployments"),
				Properties: &resources.DeploymentPropertiesExtended{
					Mode: "Incremental",
					Template: map[string]interface{}{
						"resources": []interface{}{
							SystemRoleAssignmentARM{
								ARMResource: ARMResource{
									APIVersion: "[variables('apiVersionAuthorizationSystem')]",
								},
								RoleAssignment: authorization.RoleAssignment{
									Name: to.StringPtr("[concat(variables('agentProfile1Vnet'), '/Microsoft.Authorization/', guid(uniqueString(reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 0)), '2017-03-30', 'Full').identity.principalId)))]"),
									Type: to.StringPtr("Microsoft.Network/virtualNetworks/providers/roleAssignments"),
									RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
										RoleDefinitionID: to.StringPtr("[variables('networkContributorRoleDefinitionId')]"),
										PrincipalID:      to.StringPtr("[reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 0)), '2017-03-30', 'Full').identity.principalId]"),
									},
								},
							},
							SystemRoleAssignmentARM{
								ARMResource: ARMResource{
									APIVersion: "[variables('apiVersionAuthorizationSystem')]",
								},
								RoleAssignment: authorization.RoleAssignment{
									Name: to.StringPtr("[concat(variables('agentProfile1Vnet'), '/Microsoft.Authorization/', guid(uniqueString(reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 1)), '2017-03-30', 'Full').identity.principalId)))]"),
									Type: to.StringPtr("Microsoft.Network/virtualNetworks/providers/roleAssignments"),
									RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
										RoleDefinitionID: to.StringPtr("[variables('networkContributorRoleDefinitionId')]"),
										PrincipalID:      to.StringPtr("[reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 1)), '2017-03-30', 'Full').identity.principalId]"),
									},
								},
							},
						},
						"contentVersion": "1.0.0.0",
						"$schema":        "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
					},
				},
			},
		},
		{
			DeploymentARMResource: DeploymentARMResource{
				APIVersion: "2017-05-10",
				DependsOn: []string{
					"[concat(variables('masterVMNamePrefix'), 0)]",
					"[concat(variables('masterVMNamePrefix'), 1)]",
				},
			},
			ResourceGroup: to.StringPtr("[variables('agentProfile2SubnetResourceGroup')]"),
			DeploymentExtended: resources.DeploymentExtended{
				Name: to.StringPtr("[concat('masterMsiRoleAssignment-', variables('agentProfile2VMNamePrefix'))]"),
				Type: to.StringPtr("Microsoft.Resources/deployments"),
				Properties: &resources.DeploymentPropertiesExtended{
					Mode: "Incremental",
					Template: map[string]interface{}{
						"resources": []interface{}{
							SystemRoleAssignmentARM{
								ARMResource: ARMResource{
									APIVersion: "[variables('apiVersionAuthorizationSystem')]",
								},
								RoleAssignment: authorization.RoleAssignment{
									Name: to.StringPtr("[concat(variables('agentProfile2Vnet'), '/Microsoft.Authorization/', guid(uniqueString(reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 0)), '2017-03-30', 'Full').identity.principalId)))]"),
									Type: to.StringPtr("Microsoft.Network/virtualNetworks/providers/roleAssignments"),
									RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
										RoleDefinitionID: to.StringPtr("[variables('networkContributorRoleDefinitionId')]"),
										PrincipalID:      to.StringPtr("[reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 0)), '2017-03-30', 'Full').identity.principalId]"),
									},
								},
							},
							SystemRoleAssignmentARM{
								ARMResource: ARMResource{
									APIVersion: "[variables('apiVersionAuthorizationSystem')]",
								},
								RoleAssignment: authorization.RoleAssignment{
									Name: to.StringPtr("[concat(variables('agentProfile2Vnet'), '/Microsoft.Authorization/', guid(uniqueString(reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 1)), '2017-03-30', 'Full').identity.principalId)))]"),
									Type: to.StringPtr("Microsoft.Network/virtualNetworks/providers/roleAssignments"),
									RoleAssignmentPropertiesWithScope: &authorization.RoleAssignmentPropertiesWithScope{
										RoleDefinitionID: to.StringPtr("[variables('networkContributorRoleDefinitionId')]"),
										PrincipalID:      to.StringPtr("[reference(resourceId(resourceGroup().name, 'Microsoft.Compute/virtualMachines', concat(variables('masterVMNamePrefix'), 1)), '2017-03-30', 'Full').identity.principalId]"),
									},
								},
							},
						},
						"contentVersion": "1.0.0.0",
						"$schema":        "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
					},
				},
			},
		},
	}

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}
