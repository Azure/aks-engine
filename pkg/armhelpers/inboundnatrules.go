package armhelpers

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createInboundNATRules() InboundNATRuleARM {
	return InboundNATRuleARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionNetwork')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "masterLbLoopNode",
			},
			DependsOn: []string{
				"[variables('masterLbID')]",
			},
		},
		InboundNatRule: network.InboundNatRule{
			Name: to.StringPtr("[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]", ),
			InboundNatRulePropertiesFormat: &network.InboundNatRulePropertiesFormat{
				BackendPort:      to.Int32Ptr(22),
				EnableFloatingIP: to.BoolPtr(false),
				FrontendIPConfiguration: &network.SubResource{
					ID: to.StringPtr("[variables('masterLbIPConfigID')]"),
				},
				// TODO: directly resolve sshNatPorts here
				FrontendPort: to.Int32Ptr(22), //"[variables('sshNatPorts')[copyIndex(variables('masterOffset'))]]",
				Protocol:     network.TransportProtocolTCP,
			},

			//Location: to.StringPtr("[variables('location')]"),
			//Type:     to.StringPtr("Microsoft.Network/loadBalancers/inboundNatRules"),
		},
	}
}
