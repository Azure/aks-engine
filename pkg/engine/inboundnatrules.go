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
