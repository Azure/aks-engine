// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

import (
	"testing"
)

func Test_GetAllSupportedKubernetesVersions(t *testing.T) {
	responseFromGetter := GetAllSupportedKubernetesVersions(true, false)

	if len(AllKubernetesSupportedVersions) != len(responseFromGetter) {
		t.Errorf("GetAllSupportedKubernetesVersions(true, false) returned %d items, expected %d", len(responseFromGetter), len(AllKubernetesSupportedVersions))
	}

	responseFromGetter = GetAllSupportedKubernetesVersions(false, false)

	for _, version := range responseFromGetter {
		if !AllKubernetesSupportedVersions[version] {
			t.Errorf("GetAllSupportedKubernetesVersions(false, false) returned a version %s that was not in the definitive AllKubernetesSupportedVersions map", version)
		}
	}
}

func Test_GetSupportedKubernetesVersion(t *testing.T) {
	versions := GetAllSupportedKubernetesVersions(false, false)
	for _, version := range versions {
		supportedVersion := GetSupportedKubernetesVersion(version, false)
		if supportedVersion != version {
			t.Errorf("GetSupportedKubernetesVersion(%s) should return the same passed-in string, instead returned %s", version, supportedVersion)
		}
	}

	defaultVersion := GetSupportedKubernetesVersion("", false)
	if defaultVersion != GetDefaultKubernetesVersion(false) {
		t.Errorf("GetSupportedKubernetesVersion(\"\") should return the default version %s, instead returned %s", GetDefaultKubernetesVersion(false), defaultVersion)
	}

	winVersions := GetAllSupportedKubernetesVersions(false, true)
	for _, version := range winVersions {
		supportedVersion := GetSupportedKubernetesVersion(version, true)
		if supportedVersion != version {
			t.Errorf("GetSupportedKubernetesVersion(%s) should return the same passed-in string, instead returned %s", version, supportedVersion)
		}
	}

	defaultWinVersion := GetSupportedKubernetesVersion("", true)
	if defaultWinVersion != GetDefaultKubernetesVersion(true) {
		t.Errorf("GetSupportedKubernetesVersion(\"\") should return the default version for windows %s, instead returned %s", GetDefaultKubernetesVersion(true), defaultWinVersion)
	}
}

func TestGetVersionsGt(t *testing.T) {
	versions := []string{"1.1.0-rc.1", "1.2.0-rc.1", "1.2.0", "1.2.1"}
	expected := []string{"1.1.0-rc.1", "1.2.0-rc.1", "1.2.0", "1.2.1"}
	expectedMap := map[string]bool{
		"1.1.0-rc.1": true,
		"1.2.0-rc.1": true,
		"1.2.0":      true,
		"1.2.1":      true,
	}
	v := GetVersionsGt(versions, "1.1.0-alpha.1", false, true)
	errStr := "GetVersionsGt returned an unexpected list of strings"
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.1.0", "1.2.0", "1.2.1"}
	expected = []string{"1.1.0", "1.2.0", "1.2.1"}
	expectedMap = map[string]bool{
		"1.1.0": true,
		"1.2.0": true,
		"1.2.1": true,
	}
	v = GetVersionsGt(versions, "1.1.0", true, false)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}
}

func TestIsKubernetesVersionGe(t *testing.T) {
	cases := []struct {
		name           string
		version        string
		actualVersion  string
		expectedResult bool
	}{
		{
			name:           "is 1.11.0-alpha.1 >= 1.6.0?",
			version:        "1.6.0",
			actualVersion:  "1.11.0-alpha.1",
			expectedResult: true,
		},
		{
			name:           "is 1.7.12 >= 1.8.0?",
			version:        "1.8.0",
			actualVersion:  "1.7.12",
			expectedResult: false,
		},
		{
			name:           "is 1.9.6 >= 1.9.6?",
			version:        "1.9.6",
			actualVersion:  "1.9.6",
			expectedResult: true,
		},
		{
			name:           "is 1.10.0-beta.2 >= 1.9.0?",
			version:        "1.9.0",
			actualVersion:  "1.10.0-beta.2",
			expectedResult: true,
		},
		{
			name:           "is 1.8.7 >= 1.7.0?",
			version:        "1.7.0",
			actualVersion:  "1.8.7",
			expectedResult: true,
		},
		{
			name:           "is 1.10.0-beta.2 >= 1.10.0-beta.1?",
			version:        "1.10.0-beta.1",
			actualVersion:  "1.10.0-beta.2",
			expectedResult: true,
		},
		{
			name:           "is 1.11.0-beta.1 >= 1.11.0-alpha.1?",
			version:        "1.11.0-alpha.1",
			actualVersion:  "1.11.0-beta.1",
			expectedResult: true,
		},
		{
			name:           "is 1.10.0-alpha.1 >= 1.10.0-rc.1?",
			version:        "1.10.0-rc.1",
			actualVersion:  "1.10.0-alpha.1",
			expectedResult: false,
		},
		{
			name:           "is 1.15.1 >= 1.16.0-alpha.1?",
			version:        "1.16.0-alpha.1",
			actualVersion:  "1.15.1",
			expectedResult: false,
		},
		{
			name:           "is 1.19.0 >= 1.19.0-alpha.2?",
			version:        "1.19.0",
			actualVersion:  "1.19.0-alpha.2",
			expectedResult: false,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if c.expectedResult != IsKubernetesVersionGe(c.actualVersion, c.version) {
				if c.expectedResult {
					t.Errorf("Expected version %s to be greater or equal than version %s", c.actualVersion, c.version)
				} else {
					t.Errorf("Expected version %s to not be greater or equal than version %s", c.actualVersion, c.version)
				}

			}
		})
	}
}

func TestGetVersionsLt(t *testing.T) {
	versions := []string{"1.1.0", "1.2.0-rc.1", "1.2.0", "1.2.1"}
	expected := []string{"1.1.0", "1.2.0"}
	// less than comparisons exclude pre-release versions from the result
	expectedMap := map[string]bool{
		"1.1.0": true,
		"1.2.0": true,
	}
	v := GetVersionsLt(versions, "1.2.1", false, false)
	errStr := "GetVersionsLt returned an unexpected list of strings"
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.1.0", "1.2.0", "1.2.1"}
	expected = []string{"1.1.0", "1.2.0", "1.2.1"}
	expectedMap = map[string]bool{
		"1.1.0": true,
		"1.2.0": true,
		"1.2.1": true,
	}
	v = GetVersionsLt(versions, "1.2.1", true, false)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}
}

func TestGetVersionsBetween(t *testing.T) {
	versions := []string{"1.1.0", "1.2.0", "1.2.1"}
	expected := []string{"1.2.0"}
	expectedMap := map[string]bool{
		"1.2.0": true,
	}
	v := GetVersionsBetween(versions, "1.1.0", "1.2.1", false, false)
	errStr := "GetVersionsBetween returned an unexpected list of strings"
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.1.0", "1.2.0", "1.2.1"}
	expected = []string{"1.1.0", "1.2.0", "1.2.1"}
	expectedMap = map[string]bool{
		"1.1.0": true,
		"1.2.0": true,
		"1.2.1": true,
	}
	v = GetVersionsBetween(versions, "1.1.0", "1.2.1", true, false)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.9.6", "1.10.0-beta.2", "1.10.0-beta.4", "1.10.0-rc.1"}
	expected = []string{"1.10.0-beta.2", "1.10.0-beta.4", "1.10.0-rc.1"}
	expectedMap = map[string]bool{
		"1.10.0-beta.2": true,
		"1.10.0-beta.4": true,
		"1.10.0-rc.1":   true,
	}
	v = GetVersionsBetween(versions, "1.9.6", "1.11.0", false, true)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}
	v = GetVersionsBetween(versions, "1.9.6", "1.11.0", false, false)
	if len(v) != 0 {
		t.Errorf(errStr)
	}

	versions = []string{"1.9.6", "1.10.0-beta.2", "1.10.0-beta.4", "1.10.0-rc.1"}
	expected = []string{"1.10.0-beta.4", "1.10.0-rc.1"}
	expectedMap = map[string]bool{
		"1.10.0-beta.4": true,
		"1.10.0-rc.1":   true,
	}
	v = GetVersionsBetween(versions, "1.10.0-beta.2", "1.12.0", false, false)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.10.0", "1.10.0-beta.2", "1.10.0-beta.4", "1.10.0-rc.1"}
	v = GetVersionsBetween(versions, "1.10.0", "1.12.0", false, false)
	if len(v) != 0 {
		t.Errorf(errStr)
	}

	versions = []string{"1.9.6", "1.10.0-beta.2", "1.10.0-beta.4", "1.10.0-rc.1"}
	expectedMap = map[string]bool{
		"1.9.6":         true,
		"1.10.0-beta.2": true,
		"1.10.0-beta.4": true,
		"1.10.0-rc.1":   true,
	}
	v = GetVersionsBetween(versions, "1.9.5", "1.12.0", false, true)
	if len(v) != len(versions) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.9.6", "1.10.0", "1.10.1", "1.10.2"}
	expected = []string{"1.10.0", "1.10.1", "1.10.2"}
	expectedMap = map[string]bool{
		"1.10.0": true,
		"1.10.1": true,
		"1.10.2": true,
	}
	v = GetVersionsBetween(versions, "1.10.0-rc.1", "1.12.0", false, true)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.11.0-alpha.1", "1.11.0-alpha.2", "1.11.0-beta.1"}
	expected = []string{"1.11.0-alpha.2"}
	expectedMap = map[string]bool{
		"1.11.0-alpha.2": true,
	}
	v = GetVersionsBetween(versions, "1.11.0-alpha.1", "1.11.0-beta.1", false, true)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}

	versions = []string{"1.11.0-alpha.1", "1.11.0-alpha.2", "1.11.0-beta.1"}
	expected = []string{}
	expectedMap = map[string]bool{}
	v = GetVersionsBetween(versions, "1.11.0-beta.1", "1.12.0", false, true)
	if len(v) != len(expected) {
		t.Errorf(errStr)
	}
	for _, ver := range v {
		if !expectedMap[ver] {
			t.Errorf(errStr)
		}
	}
}

func Test_GetValidPatchVersion(t *testing.T) {
	v := GetValidPatchVersion(Kubernetes, "", false, false)
	if v != GetDefaultKubernetesVersion(false) {
		t.Errorf("It is not the default Kubernetes version")
	}

	for version, enabled := range AllKubernetesSupportedVersions {
		if enabled {
			v = GetValidPatchVersion(Kubernetes, version, false, false)
			if v != version {
				t.Errorf("Expected version %s, instead got version %s", version, v)
			}
		}
	}

	v = GetValidPatchVersion(Kubernetes, "", true, true)
	if v != GetDefaultKubernetesVersion(true) {
		t.Errorf("It is not the default Kubernetes version")
	}

	v = GetValidPatchVersion(Mesos, "1.6.0", false, false)
	if v != "" {
		t.Errorf("Expected empty version for unsupported orchType")
	}

	for version, enabled := range AllKubernetesWindowsSupportedVersions {
		if enabled {
			v = GetValidPatchVersion(Kubernetes, version, false, true)
			if v != version {
				t.Errorf("Expected version %s, instead got version %s", version, v)
			}
		}
	}
}

func TestGetLatestPatchVersion(t *testing.T) {
	expected := "1.1.2"
	version := GetLatestPatchVersion("1.1", []string{"1.1.1", expected})
	if version != expected {
		t.Errorf("GetLatestPatchVersion returned the wrong latest version, expected %s, got %s", expected, version)
	}

	expected = "1.1.2"
	version = GetLatestPatchVersion("1.1", []string{"1.1.0", expected})
	if version != expected {
		t.Errorf("GetLatestPatchVersion returned the wrong latest version, expected %s, got %s", expected, version)
	}

	expected = "1.2.0"
	version = GetLatestPatchVersion("1.2", []string{"1.1.0", "1.3.0", expected})
	if version != expected {
		t.Errorf("GetLatestPatchVersion returned the wrong latest version, expected %s, got %s", expected, version)
	}

	expected = "1.2.0-rc.3"
	version = GetLatestPatchVersion("1.2", []string{"1.2.0-alpha.1", "1.2.0-beta.1", "1.2.0-rc.3", expected})
	if version != expected {
		t.Errorf("GetLatestPatchVersion returned the wrong latest version, expected %s, got %s", expected, version)
	}

	expected = ""
	version = GetLatestPatchVersion("1.2", []string{"1.1.0", "1.1.1", "1.1.2", expected})
	if version != expected {
		t.Errorf("GetLatestPatchVersion returned the wrong latest version, expected %s, got %s", expected, version)
	}
}

func TestGetMinMaxVersion(t *testing.T) {
	cases := []struct {
		name        string
		expectedMin string
		expectedMax string
		versions    []string
		preRelease  bool
	}{
		{
			name:        "list of released versions",
			expectedMin: "1.0.0",
			expectedMax: "1.0.3",
			versions:    []string{"1.0.1", "1.0.2", "1.0.0", "1.0.3"},
			preRelease:  false,
		},
		{
			name:        "list of released versions alternate",
			expectedMin: "0.0.20",
			expectedMax: "1.3.1",
			versions:    []string{"1.0.1", "1.1.2", "1.3.1", "0.0.20"},
			preRelease:  false,
		},
		{
			name:        "list of versions with pre-release but pre-release ignored",
			expectedMin: "1.0.1",
			expectedMax: "1.1.2",
			versions:    []string{"1.0.1", "1.1.2", "1.2.3-alpha.1"},
			preRelease:  false,
		},
		{
			name:        "list of versions with pre-release include pre-release versions",
			expectedMin: "1.0.1",
			expectedMax: "1.2.3-alpha.1",
			versions:    []string{"1.0.1", "1.1.2", "1.2.3-alpha.1"},
			preRelease:  true,
		},
		{
			name:        "list of versions with pre-release include pre-release versions alternate",
			expectedMin: "0.1.3-beta.1",
			expectedMax: "1.1.2",
			versions:    []string{"1.0.1", "1.1.2", "0.1.3-beta.1", "1.0.0-alpha.1"},
			preRelease:  true,
		},
		{
			name:        "empty slice ignore pre-release versions",
			expectedMin: "",
			expectedMax: "",
			versions:    []string{},
			preRelease:  false,
		},
		{
			name:        "empty slice include pre-release versions",
			expectedMin: "",
			expectedMax: "",
			versions:    []string{},
			preRelease:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			min := GetMinVersion(c.versions, c.preRelease)
			if min != c.expectedMin {
				t.Errorf("GetMinVersion returned the wrong min version, expected %s, got %s", c.expectedMin, min)
			}
			max := GetMaxVersion(c.versions, c.preRelease)
			if max != c.expectedMax {
				t.Errorf("GetMaxVersion returned the wrong max version, expected %s, got %s", c.expectedMax, max)
			}
		})
	}
}

func Test_RationalizeReleaseAndVersion(t *testing.T) {
	version := RationalizeReleaseAndVersion(Kubernetes, "", "", false, false)
	if version != GetDefaultKubernetesVersion(false) {
		t.Errorf("It is not the default Kubernetes version")
	}

	latest1Dot6Version := GetLatestPatchVersion("1.6", GetAllSupportedKubernetesVersions(true, false))
	version = RationalizeReleaseAndVersion(Kubernetes, "1.6", "", true, false)
	if version != latest1Dot6Version {
		t.Errorf("It is not Kubernetes version %s", latest1Dot6Version)
	}

	expectedVersion := "1.7.16"
	version = RationalizeReleaseAndVersion(Kubernetes, "", expectedVersion, true, false)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "1.7", expectedVersion, true, false)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "", "v"+expectedVersion, true, false)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "v1.7", "v"+expectedVersion, true, false)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "1.1", "", true, false)
	if version != "" {
		t.Errorf("It is not empty string")
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "1.1", "1.6.6", true, false)
	if version != "" {
		t.Errorf("It is not empty string")
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "", "", true, true)
	if version != GetDefaultKubernetesVersion(true) {
		t.Errorf("It is not the default Windows Kubernetes version")
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "1.5", "", true, true)
	if version != "" {
		t.Errorf("It is not empty string")
	}

	expectedVersion = "1.8.12"
	version = RationalizeReleaseAndVersion(Kubernetes, "", expectedVersion, true, true)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "1.8", expectedVersion, true, true)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "", "v"+expectedVersion, true, true)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "v1.8", "v"+expectedVersion, true, true)
	if version != expectedVersion {
		t.Errorf("It is not Kubernetes version %s", expectedVersion)
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "1.1", "", true, true)
	if version != "" {
		t.Errorf("It is not empty string")
	}

	version = RationalizeReleaseAndVersion(Kubernetes, "1.1", "1.6.6", true, true)
	if version != "" {
		t.Errorf("It is not empty string")
	}
}

func Test_IsSupportedKubernetesVersion(t *testing.T) {
	for _, isUpdate := range []bool{true, false} {
		for _, hasWindows := range []bool{true, false} {
			for _, version := range GetAllSupportedKubernetesVersions(isUpdate, hasWindows) {
				if !IsSupportedKubernetesVersion(version, isUpdate, hasWindows) {
					t.Errorf("Expected version %s to be supported when isUpdate is %t and hasWindows is %t", version, isUpdate, hasWindows)
				}
			}
		}
	}
}

func Test_IsValidMinVersion(t *testing.T) {
	orchestratorRelease := "1.15"
	orchestratorVersion := ""

	t.Run("Minimum version is valid", func(t *testing.T) {
		minVersion := "1.15.0"
		_, err := IsValidMinVersion(Kubernetes, orchestratorRelease, orchestratorVersion, minVersion)
		if err != nil {
			t.Errorf("version should be valid: %v", err)
		}
	})

	t.Run("Minimum version is invalid", func(t *testing.T) {
		minVersion := "v1.15.0"
		_, err := IsValidMinVersion(Kubernetes, orchestratorRelease, orchestratorVersion, minVersion)
		if err == nil {
			t.Errorf("version should be invalid: %v", err)
		}
	})

	t.Run("Kubernetes release is higher than required version", func(t *testing.T) {
		minVersion := "1.13.0"
		isValidVersion, _ := IsValidMinVersion(Kubernetes, orchestratorRelease, orchestratorVersion, minVersion)
		if !isValidVersion {
			t.Errorf("minimum version should be valid")
		}
	})

	t.Run("Kubernetes release is lower than minimum required version", func(t *testing.T) {
		minVersion := "1.16.0"
		isValidVersion, _ := IsValidMinVersion(Kubernetes, orchestratorRelease, orchestratorVersion, minVersion)
		if isValidVersion {
			t.Errorf("version should be not valid")
		}
	})

	t.Run("Kubernetes version is higher than required version", func(t *testing.T) {
		orchestratorVersion = "1.15.11"
		minVersion := "1.13.0"
		isValidVersion, _ := IsValidMinVersion(Kubernetes, "", orchestratorVersion, minVersion)
		if !isValidVersion {
			t.Errorf("minimum version should be valid")
		}
	})

	t.Run("Kubernetes version is lower than minimum required version", func(t *testing.T) {
		orchestratorVersion = "1.15.2"
		minVersion := "1.15.4"
		isValidVersion, _ := IsValidMinVersion(Kubernetes, "", orchestratorVersion, minVersion)
		if isValidVersion {
			t.Errorf("version should be not valid")
		}
	})
}
