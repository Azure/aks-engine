// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/blang/semver"
	"github.com/pkg/errors"
)

type orchestratorsFunc func(*OrchestratorProfile, bool, bool) ([]*OrchestratorVersionProfile, error)

var funcmap map[string]orchestratorsFunc
var versionsMap map[string][]string
var versionsMapAzureStack map[string][]string

func init() {
	funcmap = map[string]orchestratorsFunc{
		Kubernetes: kubernetesInfo,
		DCOS:       dcosInfo,
		Swarm:      swarmInfo,
	}
	versionsMap = map[string][]string{
		Kubernetes: common.GetAllSupportedKubernetesVersions(true, false, false),
		DCOS:       common.GetAllSupportedDCOSVersions(),
		Swarm:      common.GetAllSupportedSwarmVersions(),
	}
	versionsMapAzureStack = map[string][]string{
		Kubernetes: common.GetAllSupportedKubernetesVersions(true, false, true),
		DCOS:       common.GetAllSupportedDCOSVersions(),
		Swarm:      common.GetAllSupportedSwarmVersions(),
	}
}

func validate(orchestrator, version string) (string, error) {
	switch {
	case strings.EqualFold(orchestrator, Kubernetes):
		return Kubernetes, nil
	case strings.EqualFold(orchestrator, DCOS):
		return DCOS, nil
	case strings.EqualFold(orchestrator, Swarm):
		return Swarm, nil
	case orchestrator == "":
		if version != "" {
			return "", errors.Errorf("Must specify orchestrator for version '%s'", version)
		}
	default:
		return "", errors.Errorf("Unsupported orchestrator '%s'", orchestrator)
	}
	return "", nil
}

func isVersionSupported(csOrch *OrchestratorProfile, isAzureStackCloud bool) bool {
	supported := false
	versions := versionsMap[csOrch.OrchestratorType]
	if isAzureStackCloud {
		versions = versionsMapAzureStack[csOrch.OrchestratorType]
	}

	for _, version := range versions {
		if version == csOrch.OrchestratorVersion {
			supported = true
			break
		}
	}
	return supported
}

// GetOrchestratorVersionProfileListVLabs returns vlabs OrchestratorVersionProfileList object per (optionally) specified orchestrator and version
func GetOrchestratorVersionProfileListVLabs(orchestrator, version string, windows bool, azureEnv string) (*vlabs.OrchestratorVersionProfileList, error) {
	apiOrchs, err := GetOrchestratorVersionProfileList(orchestrator, version, windows, azureEnv)
	if err != nil {
		return nil, err
	}
	orchList := &vlabs.OrchestratorVersionProfileList{}
	orchList.Orchestrators = []*vlabs.OrchestratorVersionProfile{}
	for _, orch := range apiOrchs {
		orchList.Orchestrators = append(orchList.Orchestrators, ConvertOrchestratorVersionProfileToVLabs(orch))
	}
	return orchList, nil
}

// GetOrchestratorVersionProfileList returns a list of unversioned OrchestratorVersionProfile objects per (optionally) specified orchestrator and version
func GetOrchestratorVersionProfileList(orchestrator, version string, windows bool, azureEnv string) ([]*OrchestratorVersionProfile, error) {
	var err error
	isAzureStackCloud := (strings.EqualFold(azureEnv, AzureStackCloud))
	if orchestrator, err = validate(orchestrator, version); err != nil {
		return nil, err
	}
	orchs := []*OrchestratorVersionProfile{}
	if len(orchestrator) == 0 {
		// return all orchestrators
		for _, f := range funcmap {
			var arr []*OrchestratorVersionProfile
			arr, err = f(&OrchestratorProfile{}, false, isAzureStackCloud)
			if err != nil {
				return nil, err
			}
			orchs = append(orchs, arr...)
		}
	} else {
		if orchs, err = funcmap[orchestrator](&OrchestratorProfile{OrchestratorType: orchestrator, OrchestratorVersion: version}, windows, isAzureStackCloud); err != nil {
			return nil, err
		}
	}
	return orchs, nil
}

// GetOrchestratorVersionProfile returns orchestrator info for upgradable container service
func GetOrchestratorVersionProfile(orch *OrchestratorProfile, hasWindows bool, isAzureStackCloud bool) (*OrchestratorVersionProfile, error) {
	if orch.OrchestratorVersion == "" {
		return nil, errors.New("Missing Orchestrator Version")
	}
	switch orch.OrchestratorType {
	case Kubernetes, DCOS:
		arr, err := funcmap[orch.OrchestratorType](orch, hasWindows, isAzureStackCloud)
		if err != nil {
			return nil, err
		}
		// has to be exactly one element per specified orchestrator/version
		if len(arr) != 1 {
			return nil, errors.New("Ambiguous Orchestrator Versions")
		}
		return arr[0], nil
	default:
		return nil, errors.Errorf("Upgrade operation is not supported for '%s'", orch.OrchestratorType)
	}
}

func kubernetesInfo(csOrch *OrchestratorProfile, hasWindows bool, isAzureStackCloud bool) ([]*OrchestratorVersionProfile, error) {
	orchs := []*OrchestratorVersionProfile{}
	if csOrch.OrchestratorVersion == "" {
		// get info for all supported versions
		for _, ver := range common.GetAllSupportedKubernetesVersions(false, hasWindows, isAzureStackCloud) {
			upgrades, err := kubernetesUpgrades(&OrchestratorProfile{OrchestratorVersion: ver}, hasWindows, isAzureStackCloud)
			if err != nil {
				return nil, err
			}
			orchs = append(orchs,
				&OrchestratorVersionProfile{
					OrchestratorProfile: OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: ver,
					},
					Default:  ver == common.GetDefaultKubernetesVersion(hasWindows, isAzureStackCloud),
					Upgrades: upgrades,
				})
		}
	} else {
		if !isVersionSupported(csOrch, isAzureStackCloud) {
			return nil, errors.Errorf("Kubernetes version %s is not supported", csOrch.OrchestratorVersion)
		}

		upgrades, err := kubernetesUpgrades(csOrch, hasWindows, isAzureStackCloud)
		if err != nil {
			return nil, err
		}
		orchs = append(orchs,
			&OrchestratorVersionProfile{
				OrchestratorProfile: OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: csOrch.OrchestratorVersion,
				},
				Default:  csOrch.OrchestratorVersion == common.GetDefaultKubernetesVersion(hasWindows, isAzureStackCloud),
				Upgrades: upgrades,
			})
	}
	return orchs, nil
}

func kubernetesUpgrades(csOrch *OrchestratorProfile, hasWindows bool, isAzureStackCloud bool) ([]*OrchestratorProfile, error) {
	ret := []*OrchestratorProfile{}

	upgradeVersions, err := getKubernetesAvailableUpgradeVersions(csOrch.OrchestratorVersion, common.GetAllSupportedKubernetesVersions(false, hasWindows, isAzureStackCloud))
	if err != nil {
		return nil, err
	}
	for _, ver := range upgradeVersions {
		ret = append(ret, &OrchestratorProfile{
			OrchestratorType:    Kubernetes,
			OrchestratorVersion: ver,
		})
	}
	return ret, nil
}

func getKubernetesAvailableUpgradeVersions(orchestratorVersion string, supportedVersions []string) ([]string, error) {
	var skipUpgradeMinor string
	currentVer, err := semver.Make(orchestratorVersion)
	if err != nil {
		return nil, err
	}
	versionsGT := common.GetVersionsGt(supportedVersions, orchestratorVersion, false, true)
	if len(versionsGT) != 0 {
		min, err := semver.Make(common.GetMinVersion(versionsGT, true))
		if err != nil {
			return nil, err
		}

		if currentVer.Major >= min.Major && currentVer.Minor+1 < min.Minor {
			skipUpgradeMinor = strconv.FormatUint(min.Major, 10) + "." + strconv.FormatUint(min.Minor+1, 10) + ".0-alpha.0"
		} else {
			skipUpgradeMinor = strconv.FormatUint(currentVer.Major, 10) + "." + strconv.FormatUint(currentVer.Minor+2, 10) + ".0-alpha.0"
		}

		return common.GetVersionsBetween(supportedVersions, orchestratorVersion, skipUpgradeMinor, false, true), nil
	}
	return []string{}, nil

}

func dcosInfo(csOrch *OrchestratorProfile, hasWindows bool, isAzureStackCloud bool) ([]*OrchestratorVersionProfile, error) {
	orchs := []*OrchestratorVersionProfile{}
	if csOrch.OrchestratorVersion == "" {
		// get info for all supported versions
		for _, ver := range common.AllDCOSSupportedVersions {
			upgrades := dcosUpgrades(&OrchestratorProfile{OrchestratorVersion: ver})
			orchs = append(orchs,
				&OrchestratorVersionProfile{
					OrchestratorProfile: OrchestratorProfile{
						OrchestratorType:    DCOS,
						OrchestratorVersion: ver,
					},
					Default:  ver == common.DCOSDefaultVersion,
					Upgrades: upgrades,
				})
		}
	} else {
		if !isVersionSupported(csOrch, false) {
			return nil, errors.Errorf("DCOS version %s is not supported", csOrch.OrchestratorVersion)
		}

		// get info for the specified version
		upgrades := dcosUpgrades(csOrch)
		orchs = append(orchs,
			&OrchestratorVersionProfile{
				OrchestratorProfile: OrchestratorProfile{
					OrchestratorType:    DCOS,
					OrchestratorVersion: csOrch.OrchestratorVersion,
				},
				Default:  csOrch.OrchestratorVersion == common.DCOSDefaultVersion,
				Upgrades: upgrades,
			})
	}
	return orchs, nil
}

func dcosUpgrades(csOrch *OrchestratorProfile) []*OrchestratorProfile {
	ret := []*OrchestratorProfile{}

	if csOrch.OrchestratorVersion == common.DCOSVersion1Dot11Dot0 {
		ret = append(ret, &OrchestratorProfile{
			OrchestratorType:    DCOS,
			OrchestratorVersion: common.DCOSVersion1Dot11Dot2,
		})
	}
	return ret
}

func swarmInfo(csOrch *OrchestratorProfile, hasWindows bool, isAzureStackCloud bool) ([]*OrchestratorVersionProfile, error) {
	if csOrch.OrchestratorVersion == "" {
		return []*OrchestratorVersionProfile{
			{
				OrchestratorProfile: OrchestratorProfile{
					OrchestratorType:    Swarm,
					OrchestratorVersion: SwarmVersion,
				},
			},
		}, nil
	}

	if !isVersionSupported(csOrch, false) {
		return nil, errors.Errorf("Swarm version %s is not supported", csOrch.OrchestratorVersion)
	}
	return []*OrchestratorVersionProfile{
		{
			OrchestratorProfile: OrchestratorProfile{
				OrchestratorType:    Swarm,
				OrchestratorVersion: csOrch.OrchestratorVersion,
			},
		},
	}, nil
}
