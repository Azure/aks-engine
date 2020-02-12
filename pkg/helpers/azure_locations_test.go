// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import "testing"

func TestGetAzureLocations(t *testing.T) {
	expectedLocationMap := map[string]struct{}{
		"australiacentral":   {},
		"australiacentral2":  {},
		"australiaeast":      {},
		"australiasoutheast": {},
		"brazilsouth":        {},
		"canadacentral":      {},
		"canadaeast":         {},
		"centralindia":       {},
		"centralus":          {},
		"centraluseuap":      {},
		"chinaeast":          {},
		"chinaeast2":         {},
		"chinanorth":         {},
		"chinanorth2":        {},
		"eastasia":           {},
		"eastus":             {},
		"eastus2":            {},
		"eastus2euap":        {},
		"francecentral":      {},
		"francesouth":        {},
		"germanynorth":       {},
		"germanywestcentral": {},
		"japaneast":          {},
		"japanwest":          {},
		"koreacentral":       {},
		"koreasouth":         {},
		"northcentralus":     {},
		"northeurope":        {},
		"norwayeast":         {},
		"norwaywest":         {},
		"southcentralus":     {},
		"southafricanorth":   {},
		"southafricawest":    {},
		"southeastasia":      {},
		"southindia":         {},
		"switzerlandnorth":   {},
		"switzerlandwest":    {},
		"uksouth":            {},
		"ukwest":             {},
		"usdodcentral":       {},
		"usdodeast":          {},
		"westcentralus":      {},
		"westeurope":         {},
		"westindia":          {},
		"westus":             {},
		"westus2":            {},
		"germanycentral":     {},
		"germanynortheast":   {},
		"usgovvirginia":      {},
		"usgoviowa":          {},
		"usgovarizona":       {},
		"usgovtexas":         {},
		"uaenorth":           {},
		"uaecentral":         {},
	}

	locations := GetAzureLocations()

	if len(locations) == 0 {
		t.Errorf("expected the GetAzureLocations slice to be non-empty")
	}

	for _, location := range locations {
		if _, ok := expectedLocationMap[location]; !ok {
			t.Errorf("unexpected location %s found", location)
		}
	}
}
