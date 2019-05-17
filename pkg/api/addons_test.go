// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import "testing"

func TestAppendAddonIfNotPresent(t *testing.T) {
	names := []string{"AddonA", "AddonB", "AddonC", "AddonD", "AddonE"}
	addons := []KubernetesAddon{}

	for i, name := range names {
		addon := KubernetesAddon{Name: name}
		addons = appendAddonIfNotPresent(addons, addon)
		if len(addons) != i+1 {
			t.Errorf("incorrect length by appendAddonIfNotPresent, expect: '%d', actual: '%d'", i+1, len(addons))
		}
	}

	for _, name := range names {
		addon := KubernetesAddon{Name: name}
		addons = appendAddonIfNotPresent(addons, addon)
		if len(addons) != len(names) {
			t.Errorf("incorrect length by appendAddonIfNotPresent, expect: '%d', actual: '%d'", len(names), len(addons))
		}
	}
}

func TestGetAddonsIndexByName(t *testing.T) {
	includedNames := []string{"AddonA", "AddonB", "AddonC", "AddonD", "AddonE"}
	notIncludedNames := []string{"AddonF", "AddonG", "AddonH"}
	addons := []KubernetesAddon{}

	for _, name := range includedNames {
		addons = append(addons, KubernetesAddon{Name: name})
	}

	for i, addon := range addons {
		j := getAddonsIndexByName(addons, addon.Name)
		if j != i {
			t.Errorf("incorrect index by getAddonsIndexByName, expect: '%d', actual: '%d'", i, j)
		}
	}

	for _, name := range notIncludedNames {
		j := getAddonsIndexByName(addons, name)
		if j != -1 {
			t.Errorf("incorrect index by getAddonsIndexByName, expect: '%d', actual: '%d'", -1, j)
		}
	}
}
