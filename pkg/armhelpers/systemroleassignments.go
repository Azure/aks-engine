package armhelpers

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/to"
)

func createVMScaleSetRoleAssignment(cs *api.ContainerService, profile api.AgentPoolProfile) SystemRoleAssignmentARM {
	systemRoleAssignment := SystemRoleAssignmentARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionAuthorizationSystem')]",
			DependsOn: []string{
				fmt.Sprintf("[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('%sVMNamePrefix'))]", profile.Name),
			},
		},
	}

	systemRoleAssignment.Name = to.StringPtr(fmt.Sprintf("[guid(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('%sVMNamePrefix'), 'vmidentity'))]", profile.Name))
	systemRoleAssignment.Type = to.StringPtr("Microsoft.Authorization/roleAssignments")
	systemRoleAssignment.RoleDefinitionID = to.StringPtr("[variables('readerRoleDefinitionId')]")
	systemRoleAssignment.PrincipalID = to.StringPtr(fmt.Sprintf("[reference(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('%sVMNamePrefix')), '2017-03-30', 'Full').identity.principalId]", profile.Name))
	return systemRoleAssignment
}
