package armhelpers

import (
	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/go-autorest/autorest/to"
)

func createRoleAssignment() RoleAssignmentARM {
	return RoleAssignmentARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionAuthorization')]",
			Copy: map[string]string{
				"count": "[variables('masterCount')]",
				"name":  "vmLoopNode",
			},
		},
		RoleAssignment: authorization.RoleAssignment{
			Name: to.StringPtr("[guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(),'vmidentity'))]"),
			Type: to.StringPtr("Microsoft.Authorization/roleAssignments"),
			Properties: &authorization.RoleAssignmentPropertiesWithScope{
				RoleDefinitionID: to.StringPtr("[variables('contributorRoleDefinitionId')]"),
				PrincipalID:      to.StringPtr("[reference(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex()), '2017-03-30', 'Full').identity.principalId]"),
			},
		},
	}
}
