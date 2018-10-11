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
		version        string
		actualVersion  string
		expectedResult bool
	}{
		{
			version:        "1.6.0",
			actualVersion:  "1.11.0-alpha.1",
			expectedResult: true,
		},
		{
			version:        "1.8.0",
			actualVersion:  "1.7.12",
			expectedResult: false,
		},
		{
			version:        "1.9.6",
			actualVersion:  "1.9.6",
			expectedResult: true,
		},
		{
			version:        "1.9.0",
			actualVersion:  "1.10.0-beta.2",
			expectedResult: true,
		},
		{
			version:        "1.7.0",
			actualVersion:  "1.8.7",
			expectedResult: true,
		},
		{
			version:        "1.10.0-beta.1",
			actualVersion:  "1.10.0-beta.2",
			expectedResult: true,
		},
		{
			version:        "1.11.0-alpha.1",
			actualVersion:  "1.11.0-beta.1",
			expectedResult: true,
		},
		{
			version:        "1.10.0-rc.1",
			actualVersion:  "1.10.0-alpha.1",
			expectedResult: false,
		},
	}
	for _, c := range cases {
		if c.expectedResult != IsKubernetesVersionGe(c.actualVersion, c.version) {
			if c.expectedResult {
				t.Errorf("Expected version %s to be greater or equal than version %s", c.actualVersion, c.version)
			} else {
				t.Errorf("Expected version %s to not be greater or equal than version %s", c.actualVersion, c.version)
			}

		}
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
		expectedMin string
		expectedMax string
		versions    []string
		preRelease  bool
	}{
		{
			expectedMin: "1.0.0",
			expectedMax: "1.0.3",
			versions:    []string{"1.0.1", "1.0.2", "1.0.0", "1.0.3"},
			preRelease:  false,
		},
		{
			expectedMin: "0.0.20",
			expectedMax: "1.3.1",
			versions:    []string{"1.0.1", "1.1.2", "1.3.1", "0.0.20"},
			preRelease:  false,
		},
		{
			expectedMin: "1.0.1",
			expectedMax: "1.1.2",
			versions:    []string{"1.0.1", "1.1.2", "1.2.3-alpha.1"},
			preRelease:  false,
		},
		{
			expectedMin: "1.0.1",
			expectedMax: "1.2.3-alpha.1",
			versions:    []string{"1.0.1", "1.1.2", "1.2.3-alpha.1"},
			preRelease:  true,
		},
		{
			expectedMin: "0.1.3-beta.1",
			expectedMax: "1.1.2",
			versions:    []string{"1.0.1", "1.1.2", "0.1.3-beta.1", "1.0.0-alpha.1"},
			preRelease:  true,
		},
		{
			expectedMin: "",
			expectedMax: "",
			versions:    []string{},
			preRelease:  false,
		},
		{
			expectedMin: "",
			expectedMax: "",
			versions:    []string{},
			preRelease:  true,
		},
	}

	for _, c := range cases {
		min := GetMinVersion(c.versions, c.preRelease)
		if min != c.expectedMin {
			t.Errorf("GetMinVersion returned the wrong min version, expected %s, got %s", c.expectedMin, min)
		}
		max := GetMaxVersion(c.versions, c.preRelease)
		if max != c.expectedMax {
			t.Errorf("GetMaxVersion returned the wrong max version, expected %s, got %s", c.expectedMax, max)
		}
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
