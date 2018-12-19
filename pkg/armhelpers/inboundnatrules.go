package armhelpers

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

// InboundNatRule inbound NAT rule of the load balancer.
type InboundNatRule struct {
	// InboundNatRulePropertiesFormat - Properties of load balancer inbound nat rule.
	*network.InboundNatRulePropertiesFormat `json:"properties,omitempty"`
	// Name - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Location - Resource Location
	Location *string `json:"location,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}

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
		InboundNatRule: InboundNatRule{
			Name: to.StringPtr("[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]", ),
			InboundNatRulePropertiesFormat: &network.InboundNatRulePropertiesFormat{
				BackendPort:      to.Int32Ptr(22),
				EnableFloatingIP: to.BoolPtr(false),
				FrontendIPConfiguration: &network.SubResource{
					ID: to.StringPtr("[variables('masterLbIPConfigID')]"),
				},
				// TODO: directly resolve sshNatPorts here
				FrontendPort: to.Int32Ptr(123123123), //"[variables('sshNatPorts')[copyIndex(variables('masterOffset'))]]",
				Protocol:     network.TransportProtocolTCP,
			},

			Location: to.StringPtr("[variables('location')]"),
			Type:     to.StringPtr("Microsoft.Network/loadBalancers/inboundNatRules"),
		},
	}
}
