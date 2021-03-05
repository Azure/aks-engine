// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import (
	"fmt"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
)

// GetKubernetesAllowedVMSKUs returns the allowed sizes for Kubernetes agent
func GetKubernetesAllowedVMSKUs() string {
	var b strings.Builder
	b.WriteString("      \"allowedValues\": [\n")
	for i, sku := range VMSkus {
		fmt.Fprintf(&b, "\"%s\"", sku.Name)
		if i < len(VMSkus)-1 {
			b.WriteByte(',')
		}
		b.WriteByte('\n')
	}
	b.WriteString("    ],\n")
	return b.String()
}

// GetSizeMap returns the size / storage map
func GetSizeMap() string {
	var b strings.Builder
	b.WriteString("    \"vmSizesMap\": {\n")
	for i, sku := range VMSkus {
		storageType, err := common.GetStorageAccountType(sku.Name)
		if err != nil {
			storageType = err.Error()
		}
		fmt.Fprintf(&b, "    \"%s\": {\n      \"storageAccountType\": \"%s\"\n    }",
			sku.Name, storageType)
		if i < len(VMSkus)-1 {
			b.WriteByte(',')
		}
		b.WriteByte('\n')
	}
	b.WriteString("   }\n  ")
	return b.String()
}

// AcceleratedNetworkingSupported checks if the VM SKU supports Accelerated Networking.
func AcceleratedNetworkingSupported(sku string) bool {
	name := strings.TrimSuffix(sku, "_Promo")
	for _, sku := range VMSkus {
		if name == sku.Name {
			return sku.AcceleratedNetworking
		}
	}
	return false
}
