// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func createInboundNATRules(masterCount int) []InboundNATRuleARM {

	var natRules []InboundNATRuleARM
	sshNatPorts := []int32{
		22, 2201, 2202, 2203, 2204,
	}
	for i := 0; i < masterCount; i++ {
		natRule := InboundNATRuleARM{
			ARMResourceLocType: ARMResourceLocType{
				ARMResource: ARMResource{
					ApiVersion: "[variables('apiVersionNetwork')]",
					DependsOn: []string{
						"[variables('masterLbID')]",
					},
				},
				Location: "[variables('location')]",
				Type:     "Microsoft.Network/loadBalancers/inboundNatRules",
			},
			InboundNatRule: network.InboundNatRule{
				Name: to.StringPtr(fmt.Sprintf("[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), %d)]", i)),
				InboundNatRulePropertiesFormat: &network.InboundNatRulePropertiesFormat{
					BackendPort:      to.Int32Ptr(22),
					EnableFloatingIP: to.BoolPtr(false),
					FrontendIPConfiguration: &network.SubResource{
						ID: to.StringPtr("[variables('masterLbIPConfigID')]"),
					},
					// TODO: directly resolve sshNatPorts here
					FrontendPort: to.Int32Ptr(sshNatPorts[i]), //"[variables('sshNatPorts')[copyIndex(variables('masterOffset'))]]",
					Protocol:     network.TransportProtocolTCP,
				},
			},
		}

		natRules = append(natRules, natRule)
	}

	return natRules
}
