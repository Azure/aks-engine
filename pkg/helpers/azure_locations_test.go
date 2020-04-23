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
		"germanycentral":     {},
		"germanynorth":       {},
		"germanynortheast":   {},
		"germanywestcentral": {},
		"japaneast":          {},
		"japanwest":          {},
		"koreacentral":       {},
		"koreasouth":         {},
		"northcentralus":     {},
		"northeurope":        {},
		"norwayeast":         {},
		"norwaywest":         {},
		"southafricanorth":   {},
		"southafricawest":    {},
		"southcentralus":     {},
		"southeastasia":      {},
		"southindia":         {},
		"switzerlandnorth":   {},
		"switzerlandwest":    {},
		"uaecentral":         {},
		"uaenorth":           {},
		"uksouth":            {},
		"ukwest":             {},
		"usdodcentral":       {},
		"usdodeast":          {},
		"usgovarizona":       {},
		"usgoviowa":          {},
		"usgovtexas":         {},
		"usgovvirginia":      {},
		"westcentralus":      {},
		"westeurope":         {},
		"westindia":          {},
		"westus":             {},
		"westus2":            {},
	}

	locations := GetAzureLocations()

	if len(locations) != len(expectedLocationMap) {
		t.Errorf("expected the GetAzureLocations slice to be of length %d",
			len(expectedLocationMap))
	}

	for _, location := range locations {
		if _, ok := expectedLocationMap[location]; !ok {
			t.Errorf("unexpected location %s found", location)
		}
	}
}
