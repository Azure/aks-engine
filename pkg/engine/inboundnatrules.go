// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

func createInboundNATRules() map[string]interface{} {

	inboundNatRules := map[string]interface{}{
		"apiVersion": "[variables('apiVersionNetwork')]",
		"copy": map[string]interface{}{
			"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
			"name":  "masterLbLoopNode",
		},
		"dependsOn": []string{
			"[variables('masterLbID')]",
		},
		"location": "[variables('location')]",
		"name":     "[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]",
		"properties": map[string]interface{}{
			"backendPort":      22,
			"enableFloatingIP": false,
			"frontendIPConfiguration": map[string]string{
				"id": "[variables('masterLbIPConfigID')]",
			},
			"frontendPort": "[variables('sshNatPorts')[copyIndex(variables('masterOffset'))]]",
			"protocol":     "Tcp",
		},
		"type": "Microsoft.Network/loadBalancers/inboundNatRules",
	}

	return inboundNatRules
}

//func createInboundNATRules(masterCount int) []InboundNATRuleARM {
//
//	var natRules []InboundNATRuleARM
//	sshNatPorts := []int32{
//		22, 2201, 2202, 2203, 2204,
//	}
//	for i := 0; i < masterCount; i++ {
//		natRule := InboundNATRuleARM{
//			ARMResourceLocType: ARMResourceLocType{
//				ARMResource: ARMResource{
//					APIVersion: "[variables('apiVersionNetwork')]",
//					DependsOn: []string{
//						"[variables('masterLbID')]",
//					},
//				},
//				Location: "[variables('location')]",
//				Type:     "Microsoft.Network/loadBalancers/inboundNatRules",
//			},
//			InboundNatRule: network.InboundNatRule{
//				Name: to.StringPtr(fmt.Sprintf("[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), '%d')]", i)),
//				InboundNatRulePropertiesFormat: &network.InboundNatRulePropertiesFormat{
//					BackendPort:      to.Int32Ptr(22),
//					EnableFloatingIP: to.BoolPtr(false),
//					FrontendIPConfiguration: &network.SubResource{
//						ID: to.StringPtr("[variables('masterLbIPConfigID')]"),
//					},
//					// TODO: directly resolve sshNatPorts here
//					FrontendPort: to.Int32Ptr(sshNatPorts[i]), //"[variables('sshNatPorts')[copyIndex(variables('masterOffset'))]]",
//					Protocol:     network.TransportProtocolTCP,
//				},
//			},
//		}
//
//		natRules = append(natRules, natRule)
//	}
//
//	return natRules
//}
