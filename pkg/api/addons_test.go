// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
)

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

func TestPodSecurityPolicyConfigUpgrade(t *testing.T) {
	mockCS := getMockBaseContainerService("1.8.0")
	o := mockCS.Properties.OrchestratorProfile

	isUpdate := true
	base64DataPSP := "cHNwQ3VzdG9tRGF0YQ=="
	o.OrchestratorType = Kubernetes
	o.KubernetesConfig.EnablePodSecurityPolicy = to.BoolPtr(true)
	o.KubernetesConfig.PodSecurityPolicyConfig = map[string]string{
		"data": base64DataPSP,
	}

	mockCS.setAddonsConfig(isUpdate)

	i := getAddonsIndexByName(o.KubernetesConfig.Addons, PodSecurityPolicyAddonName)
	if i < 0 {
		t.Errorf("expected a positive index for the addon %s, instead got %d from getAddonsIndexByName", PodSecurityPolicyAddonName, i)
	}

	if o.KubernetesConfig.Addons[i].Name != PodSecurityPolicyAddonName {
		t.Errorf("expected addon %s name to be present, instead got %s", PodSecurityPolicyAddonName, o.KubernetesConfig.Addons[i].Name)
	}

	if o.KubernetesConfig.Addons[i].Data != base64DataPSP {
		t.Errorf("expected %s data to be present, instead got %s", base64DataPSP, o.KubernetesConfig.Addons[i].Data)
	}
}

func TestDisabledAddons(t *testing.T) {
	defaultAddon := KubernetesAddon{
		Name:    "mockAddon",
		Enabled: to.BoolPtr(false),
		Containers: []KubernetesContainerSpec{
			{
				Name:           "mockAddon",
				CPURequests:    "50m",
				MemoryRequests: "50Mi",
				CPULimits:      "50m",
				MemoryLimits:   "250Mi",
				Image:          "mockImage",
			},
		},
		Config: map[string]string{
			"fake-config-1": "someValue1",
			"fake-config-2": "someValue2",
		},
	}

	cases := []struct {
		myAddon             KubernetesAddon
		isUpdate            bool
		expectedResultAddon KubernetesAddon
	}{
		{
			myAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(true),
			},
			isUpdate: false,
			expectedResultAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name:           "mockAddon",
						CPURequests:    "50m",
						MemoryRequests: "50Mi",
						CPULimits:      "50m",
						MemoryLimits:   "250Mi",
						Image:          "mockImage",
					},
				},
				Config: map[string]string{
					"fake-config-1": "someValue1",
					"fake-config-2": "someValue2",
				},
			},
		},
		{
			myAddon: KubernetesAddon{
				Name: "mockAddon",
			},
			isUpdate: false,
			expectedResultAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
		},
		{
			myAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
			isUpdate: true,
			expectedResultAddon: KubernetesAddon{
				Name:    "mockAddon",
				Enabled: to.BoolPtr(false),
			},
		},
	}

	for _, c := range cases {
		result := assignDefaultAddonVals(c.myAddon, defaultAddon, c.isUpdate)
		if !reflect.DeepEqual(result, c.expectedResultAddon) {
			t.Fatalf("expected result addon %v to be equal to %v", result, c.expectedResultAddon)
		}
	}

}
