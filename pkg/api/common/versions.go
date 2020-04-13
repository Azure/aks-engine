// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

import (
	"fmt"
	"sort"
	"strings"

	"github.com/blang/semver"
	"github.com/pkg/errors"
)

// AllKubernetesSupportedVersions is a whitelist map of all supported Kubernetes version strings
// The bool value indicates if creating new clusters with this version is allowed
var AllKubernetesSupportedVersions = map[string]bool{
	"1.6.6":          false,
	"1.6.9":          true, // need to keep 1.6.9 version support for v20160930
	"1.6.11":         false,
	"1.6.12":         false,
	"1.6.13":         false,
	"1.7.0":          false,
	"1.7.1":          false,
	"1.7.2":          false,
	"1.7.4":          false,
	"1.7.5":          false,
	"1.7.7":          false,
	"1.7.9":          false,
	"1.7.10":         false,
	"1.7.12":         false,
	"1.7.13":         false,
	"1.7.14":         false,
	"1.7.15":         false,
	"1.7.16":         false,
	"1.8.0":          false,
	"1.8.1":          false,
	"1.8.2":          false,
	"1.8.4":          false,
	"1.8.6":          false,
	"1.8.7":          false,
	"1.8.8":          false,
	"1.8.9":          false,
	"1.8.10":         false,
	"1.8.11":         false,
	"1.8.12":         false,
	"1.8.13":         false,
	"1.8.14":         false,
	"1.8.15":         false,
	"1.9.0":          false,
	"1.9.1":          false,
	"1.9.2":          false,
	"1.9.3":          false,
	"1.9.4":          false,
	"1.9.5":          false,
	"1.9.6":          false,
	"1.9.7":          false,
	"1.9.8":          false,
	"1.9.9":          false,
	"1.9.10":         false,
	"1.9.11":         false,
	"1.10.0-beta.2":  false,
	"1.10.0-beta.4":  false,
	"1.10.0-rc.1":    false,
	"1.10.0":         false,
	"1.10.1":         false,
	"1.10.2":         false,
	"1.10.3":         false,
	"1.10.4":         false,
	"1.10.5":         false,
	"1.10.6":         false,
	"1.10.7":         false,
	"1.10.8":         false,
	"1.10.9":         false,
	"1.10.12":        false,
	"1.10.13":        false,
	"1.11.0-alpha.1": false,
	"1.11.0-alpha.2": false,
	"1.11.0-beta.1":  false,
	"1.11.0-beta.2":  false,
	"1.11.0-rc.1":    false,
	"1.11.0-rc.2":    false,
	"1.11.0-rc.3":    false,
	"1.11.0":         false,
	"1.11.1":         false,
	"1.11.2":         false,
	"1.11.3":         false,
	"1.11.4":         false,
	"1.11.5":         false,
	"1.11.6":         false,
	"1.11.7":         false,
	"1.11.8":         false,
	"1.11.9":         false,
	"1.11.10":        false,
	"1.12.0-alpha.1": false,
	"1.12.0-beta.0":  false,
	"1.12.0-beta.1":  false,
	"1.12.0-rc.1":    false,
	"1.12.0-rc.2":    false,
	"1.12.0":         false,
	"1.12.1":         false,
	"1.12.2":         false,
	"1.12.4":         false,
	"1.12.5":         false,
	"1.12.6":         false,
	"1.12.7":         false,
	"1.12.8":         false,
	"1.12.9":         false, // disabled because of https://github.com/Azure/aks-engine/issues/1421
	"1.13.0-alpha.1": false,
	"1.13.0-alpha.2": false,
	"1.13.1":         false,
	"1.13.2":         false,
	"1.13.3":         false,
	"1.13.4":         false,
	"1.13.5":         false,
	"1.13.6":         false, // disabled because of https://github.com/kubernetes/kubernetes/issues/78308
	"1.13.7":         false,
	"1.13.8":         false,
	"1.13.9":         false,
	"1.13.10":        false,
	"1.13.11":        false,
	"1.13.12":        false, // disabled because of https://github.com/Azure/aks-engine/issues/2312
	"1.14.0-alpha.1": false,
	"1.14.0-alpha.2": false,
	"1.14.0-beta.1":  false,
	"1.14.0-beta.2":  false,
	"1.14.0-rc.1":    false,
	"1.14.0":         false,
	"1.14.1":         false,
	"1.14.2":         false, // disabled because of https://github.com/kubernetes/kubernetes/issues/78308
	"1.14.3":         false,
	"1.14.4":         false,
	"1.14.5":         false,
	"1.14.6":         false,
	"1.14.7":         true,
	"1.14.8":         false, // disabled because of https://github.com/Azure/aks-engine/issues/2312
	"1.14.10":        false, // disabled because of https://github.com/Azure/aks-engine/issues/2312
	"1.15.0-alpha.1": false,
	"1.15.0-alpha.2": false,
	"1.15.0-alpha.3": false,
	"1.15.0-beta.1":  false,
	"1.15.0-beta.2":  false,
	"1.15.0-rc.1":    false,
	"1.15.0":         false,
	"1.15.1":         false,
	"1.15.2":         false,
	"1.15.3":         false,
	"1.15.4":         false,
	"1.15.5":         false, // disabled because of https://github.com/Azure/aks-engine/issues/2312
	"1.15.7":         false,
	"1.15.8":         false, // disabled because of https://github.com/kubernetes/release/issues/1020
	"1.15.9":         false,
	"1.15.10":        false,
	"1.15.11":        true,
	"1.15.12":        true,
	"1.16.0-alpha.1": false,
	"1.16.0-alpha.2": false,
	"1.16.0-alpha.3": false,
	"1.16.0-beta.1":  false,
	"1.16.0-beta.2":  false,
	"1.16.0-rc.1":    false,
	"1.16.0":         false,
	"1.16.1":         false,
	"1.16.2":         false, // disabled because of https://github.com/Azure/aks-engine/issues/2312
	"1.16.4":         false,
	"1.16.5":         false, // disabled because of https://github.com/kubernetes/release/issues/1020
	"1.16.6":         false,
	"1.16.7":         false,
	"1.16.8":         true,
	"1.16.9":         true,
	"1.17.0-alpha.1": false,
	"1.17.0-alpha.2": false,
	"1.17.0-alpha.3": false,
	"1.17.0-beta.1":  false,
	"1.17.0-beta.2":  false,
	"1.17.0-rc.1":    false,
	"1.17.0-rc.2":    false,
	"1.17.0":         false,
	"1.17.1":         false,
	"1.17.2":         false,
	"1.17.3":         false,
	"1.17.4":         true,
	"1.17.5":         true,
	"1.18.0-alpha.1": false,
	"1.18.0-alpha.2": false,
	"1.18.0-alpha.3": false,
	"1.18.0-alpha.5": false,
	"1.18.0-beta.1":  false,
	"1.18.0":         false,
	"1.18.1":         true,
	"1.18.2":         true,
	"1.19.0-alpha.1": true,
	"1.19.0-alpha.2": true,
	"1.19.0-alpha.3": true,
}

// GetDefaultKubernetesVersion returns the default Kubernetes version, that is the latest patch of the default release
func GetDefaultKubernetesVersion(hasWindows bool) string {
	defaultRelease := KubernetesDefaultRelease
	if hasWindows {
		defaultRelease = KubernetesDefaultReleaseWindows
	}
	return GetLatestPatchVersion(defaultRelease, GetAllSupportedKubernetesVersions(false, hasWindows))
}

// GetSupportedKubernetesVersion verifies that a passed-in version string is supported, or returns a default version string if not
func GetSupportedKubernetesVersion(version string, hasWindows bool) string {
	k8sVersion := GetDefaultKubernetesVersion(hasWindows)
	if hasWindows {
		if AllKubernetesWindowsSupportedVersions[version] {
			k8sVersion = version
		}
	} else {
		if AllKubernetesSupportedVersions[version] {
			k8sVersion = version
		}
	}
	return k8sVersion
}

// GetAllSupportedKubernetesVersions returns a slice of all supported Kubernetes versions
func GetAllSupportedKubernetesVersions(isUpdate, hasWindows bool) []string {
	var versions []string
	allSupportedVersions := AllKubernetesSupportedVersions
	if hasWindows {
		allSupportedVersions = AllKubernetesWindowsSupportedVersions
	}
	for ver, supported := range allSupportedVersions {
		if isUpdate || supported {
			versions = append(versions, ver)
		}
	}
	sort.Slice(versions, func(i, j int) bool {
		return IsKubernetesVersionGe(versions[j], versions[i])
	})
	return versions
}

// GetVersionsGt returns a list of versions greater than a semver string given a list of versions
// inclusive=true means that we test for equality as well
// preReleases=true means that we include pre-release versions in the list
func GetVersionsGt(versions []string, version string, inclusive, preReleases bool) []string {
	// Try to get latest version matching the release
	var ret []string
	minVersion, _ := semver.Make(version)
	for _, v := range versions {
		sv, _ := semver.Make(v)
		if !preReleases && len(sv.Pre) != 0 {
			continue
		}
		if (inclusive && sv.GTE(minVersion)) || (!inclusive && sv.GT(minVersion)) {
			ret = append(ret, v)
		}
	}
	return ret
}

// GetVersionsLt returns a list of versions less than than a semver string given a list of versions
// inclusive=true means that we test for equality as well
// preReleases=true means that we include pre-release versions in the list
func GetVersionsLt(versions []string, version string, inclusive, preReleases bool) []string {
	// Try to get latest version matching the release
	var ret []string
	minVersion, _ := semver.Make(version)
	for _, v := range versions {
		sv, _ := semver.Make(v)
		if !preReleases && len(sv.Pre) != 0 {
			continue
		}
		if (inclusive && sv.LTE(minVersion)) || (!inclusive && sv.LT(minVersion)) {
			ret = append(ret, v)
		}
	}
	return ret
}

// GetVersionsBetween returns a list of versions between a min and max
// inclusive=true means that we test for equality on both bounds
// preReleases=true means that we include pre-release versions in the list
func GetVersionsBetween(versions []string, versionMin, versionMax string, inclusive, preReleases bool) []string {
	var ret []string
	if minV, _ := semver.Make(versionMin); len(minV.Pre) != 0 {
		preReleases = true
	}
	greaterThan := GetVersionsGt(versions, versionMin, inclusive, preReleases)
	lessThan := GetVersionsLt(versions, versionMax, inclusive, preReleases)
	for _, lv := range lessThan {
		for _, gv := range greaterThan {
			if lv == gv {
				ret = append(ret, lv)
			}
		}
	}
	return ret
}

// GetMinVersion gets the lowest semver version
// preRelease=true means accept a pre-release version as a min value
func GetMinVersion(versions []string, preRelease bool) string {
	if len(versions) < 1 {
		return ""
	}
	semverVersions := getSortedSemverVersions(versions, preRelease)
	return semverVersions[0].String()
}

// GetMaxVersion gets the highest semver version
// preRelease=true means accept a pre-release version as a max value
func GetMaxVersion(versions []string, preRelease bool) string {
	if len(versions) < 1 {
		return ""
	}
	semverVersions := getSortedSemverVersions(versions, preRelease)
	return semverVersions[len(semverVersions)-1].String()
}

func getSortedSemverVersions(versions []string, preRelease bool) []semver.Version {
	var semverVersions []semver.Version
	for _, v := range versions {
		sv, _ := semver.Make(v)
		if len(sv.Pre) == 0 || preRelease {
			semverVersions = append(semverVersions, sv)
		}
	}
	semver.Sort(semverVersions)
	return semverVersions
}

// AllKubernetesWindowsSupportedVersions maintain a set of available k8s Windows versions in aks-engine
var AllKubernetesWindowsSupportedVersions = getAllKubernetesWindowsSupportedVersionsMap()

func getAllKubernetesWindowsSupportedVersionsMap() map[string]bool {
	ret := make(map[string]bool)
	for k, v := range AllKubernetesSupportedVersions {
		ret[k] = v
	}
	for _, version := range []string{
		"1.6.6",
		"1.6.9",
		"1.6.11",
		"1.6.12",
		"1.6.13",
		"1.7.0",
		"1.7.1",
		"1.8.13",
		"1.8.14",
		"1.8.15",
		"1.10.0-beta.2",
		"1.10.0-beta.4",
		"1.10.0-rc.1",
		"1.11.0-alpha.1",
		"1.11.0-alpha.2"} {
		delete(ret, version)
	}
	// 1.8.12 is the latest supported patch for Windows
	ret["1.8.12"] = true
	return ret
}

// GetSupportedVersions get supported version list for a certain orchestrator
func GetSupportedVersions(orchType string, isUpdate, hasWindows bool) (versions []string, defaultVersion string) {
	switch orchType {
	case Kubernetes:
		return GetAllSupportedKubernetesVersions(isUpdate, hasWindows), GetDefaultKubernetesVersion(hasWindows)
	case DCOS:
		return AllDCOSSupportedVersions, DCOSDefaultVersion
	default:
		return nil, ""
	}
}

//GetValidPatchVersion gets the current valid patch version for the minor version of the passed in version
func GetValidPatchVersion(orchType, orchVer string, isUpdate, hasWindows bool) string {
	if orchVer == "" {
		return RationalizeReleaseAndVersion(
			orchType,
			"",
			"",
			isUpdate,
			hasWindows)
	}

	// check if the current version is valid, this allows us to have multiple supported patch versions in the future if we need it
	version := RationalizeReleaseAndVersion(
		orchType,
		"",
		orchVer,
		isUpdate,
		hasWindows)

	if version == "" {
		sv, err := semver.Make(orchVer)
		if err != nil {
			return ""
		}
		sr := fmt.Sprintf("%d.%d", sv.Major, sv.Minor)

		version = RationalizeReleaseAndVersion(
			orchType,
			sr,
			"",
			isUpdate,
			hasWindows)
	}
	return version
}

// RationalizeReleaseAndVersion return a version when it can be rationalized from the input, otherwise ""
func RationalizeReleaseAndVersion(orchType, orchRel, orchVer string, isUpdate, hasWindows bool) (version string) {
	// ignore "v" prefix in orchestrator version and release: "v1.8.0" is equivalent to "1.8.0", "v1.9" is equivalent to "1.9"
	orchVer = strings.TrimPrefix(orchVer, "v")
	orchRel = strings.TrimPrefix(orchRel, "v")
	supportedVersions, defaultVersion := GetSupportedVersions(orchType, isUpdate, hasWindows)
	if supportedVersions == nil {
		return ""
	}

	if orchRel == "" && orchVer == "" {
		return defaultVersion
	}

	if orchVer == "" {
		// Try to get latest version matching the release
		version = GetLatestPatchVersion(orchRel, supportedVersions)
		return version
	} else if orchRel == "" {
		// Try to get version the same with orchVer
		version = ""
		for _, ver := range supportedVersions {
			if ver == orchVer {
				version = ver
				break
			}
		}
		return version
	}
	// Try to get latest version matching the release
	version = ""
	for _, ver := range supportedVersions {
		sv, _ := semver.Make(ver)
		sr := fmt.Sprintf("%d.%d", sv.Major, sv.Minor)
		if sr == orchRel && ver == orchVer {
			version = ver
			break
		}
	}
	return version
}

func IsValidMinVersion(orchType, orchRelease, orchVersion, minVersion string) (bool, error) {
	version := RationalizeReleaseAndVersion(
		orchType,
		orchRelease,
		orchVersion,
		false,
		false)
	if version == "" {
		return false, errors.Errorf("the following user supplied OrchestratorProfile configuration is not supported: OrchestratorType: %s, OrchestratorRelease: %s, OrchestratorVersion: %s. Please check supported Release or Version for this build of aks-engine",
			orchType,
			orchRelease,
			orchVersion)
	}
	sv, err := semver.Make(version)
	if err != nil {
		return false, errors.Errorf("could not validate version %s", version)
	}
	m, err := semver.Make(minVersion)
	if err != nil {
		return false, errors.New("could not validate version")
	}
	if sv.LT(m) {
		return false, nil
	}
	return true, nil
}

// IsKubernetesVersionGe returns true if actualVersion is greater than or equal to version
func IsKubernetesVersionGe(actualVersion, version string) bool {
	v1, _ := semver.Make(actualVersion)
	v2, _ := semver.Make(version)
	return v1.GE(v2)
}

// GetLatestPatchVersion gets the most recent patch version from a list of semver versions given a major.minor string
func GetLatestPatchVersion(majorMinor string, versionsList []string) (version string) {
	// Try to get latest version matching the release
	version = ""
	for _, ver := range versionsList {
		sv, err := semver.Make(ver)
		if err != nil {
			return
		}
		sr := fmt.Sprintf("%d.%d", sv.Major, sv.Minor)
		if sr == majorMinor {
			if version == "" {
				version = ver
			} else {
				current, _ := semver.Make(version)
				if sv.GT(current) {
					version = ver
				}
			}
		}
	}
	return version
}

// IsSupportedKubernetesVersion return true if the provided Kubernetes version is supported
func IsSupportedKubernetesVersion(version string, isUpdate, hasWindows bool) bool {
	for _, ver := range GetAllSupportedKubernetesVersions(isUpdate, hasWindows) {
		if ver == version {
			return true
		}
	}
	return false
}
