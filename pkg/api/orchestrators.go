// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/v20170930"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/blang/semver"
	"github.com/pkg/errors"
)

type orchestratorsFunc func(*OrchestratorProfile, bool) ([]*OrchestratorVersionProfile, error)

var funcmap map[string]orchestratorsFunc
var versionsMap map[string][]string

func init() {
	funcmap = map[string]orchestratorsFunc{
		Kubernetes: kubernetesInfo,
		DCOS:       dcosInfo,
		Swarm:      swarmInfo,
		SwarmMode:  dockerceInfo,
	}
	versionsMap = map[string][]string{
		Kubernetes: common.GetAllSupportedKubernetesVersions(true, false),
		DCOS:       common.GetAllSupportedDCOSVersions(),
		Swarm:      common.GetAllSupportedSwarmVersions(),
		SwarmMode:  common.GetAllSupportedDockerCEVersions(),
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
	case strings.EqualFold(orchestrator, SwarmMode):
		return SwarmMode, nil
	case orchestrator == "":
		if version != "" {
			return "", errors.Errorf("Must specify orchestrator for version '%s'", version)
		}
	default:
		return "", errors.Errorf("Unsupported orchestrator '%s'", orchestrator)
	}
	return "", nil
}

func isVersionSupported(csOrch *OrchestratorProfile) bool {
	supported := false
	for _, version := range versionsMap[csOrch.OrchestratorType] {

		if version == csOrch.OrchestratorVersion {
			supported = true
			break
		}
	}
	return supported
}

// GetOrchestratorVersionProfileListVLabs returns vlabs OrchestratorVersionProfileList object per (optionally) specified orchestrator and version
func GetOrchestratorVersionProfileListVLabs(orchestrator, version string, windows bool) (*vlabs.OrchestratorVersionProfileList, error) {
	apiOrchs, err := getOrchestratorVersionProfileList(orchestrator, version, windows)
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

// GetOrchestratorVersionProfileListV20170930 returns v20170930 OrchestratorVersionProfileList object per (optionally) specified orchestrator and version
func GetOrchestratorVersionProfileListV20170930(orchestrator, version string) (*v20170930.OrchestratorVersionProfileList, error) {
	apiOrchs, err := getOrchestratorVersionProfileList(orchestrator, version, false)
	if err != nil {
		return nil, err
	}
	orchList := &v20170930.OrchestratorVersionProfileList{}
	for _, orch := range apiOrchs {
		orchList.Properties.Orchestrators = append(orchList.Properties.Orchestrators, ConvertOrchestratorVersionProfileToV20170930(orch))
	}
	return orchList, nil
}

func getOrchestratorVersionProfileList(orchestrator, version string, windows bool) ([]*OrchestratorVersionProfile, error) {
	var err error
	if orchestrator, err = validate(orchestrator, version); err != nil {
		return nil, err
	}
	orchs := []*OrchestratorVersionProfile{}
	if len(orchestrator) == 0 {
		// return all orchestrators
		for _, f := range funcmap {
			arr, err := f(&OrchestratorProfile{}, false)
			if err != nil {
				return nil, err
			}
			orchs = append(orchs, arr...)
		}
	} else {
		if orchs, err = funcmap[orchestrator](&OrchestratorProfile{OrchestratorType: orchestrator, OrchestratorVersion: version}, windows); err != nil {
			return nil, err
		}
	}
	return orchs, nil
}

// GetOrchestratorVersionProfile returns orchestrator info for upgradable container service
func GetOrchestratorVersionProfile(orch *OrchestratorProfile, hasWindows bool) (*OrchestratorVersionProfile, error) {
	if orch.OrchestratorVersion == "" {
		return nil, errors.New("Missing Orchestrator Version")
	}
	switch orch.OrchestratorType {
	case Kubernetes, DCOS:
		arr, err := funcmap[orch.OrchestratorType](orch, hasWindows)
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

func kubernetesInfo(csOrch *OrchestratorProfile, hasWindows bool) ([]*OrchestratorVersionProfile, error) {
	orchs := []*OrchestratorVersionProfile{}
	if csOrch.OrchestratorVersion == "" {
		// get info for all supported versions
		for _, ver := range common.GetAllSupportedKubernetesVersions(false, hasWindows) {
			upgrades, err := kubernetesUpgrades(&OrchestratorProfile{OrchestratorVersion: ver}, hasWindows)
			if err != nil {
				return nil, err
			}
			orchs = append(orchs,
				&OrchestratorVersionProfile{
					OrchestratorProfile: OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: ver,
					},
					Default:  ver == common.GetDefaultKubernetesVersion(hasWindows),
					Upgrades: upgrades,
				})
		}
	} else {
		if !isVersionSupported(csOrch) {
			return nil, errors.Errorf("Kubernetes version %s is not supported", csOrch.OrchestratorVersion)
		}

		upgrades, err := kubernetesUpgrades(csOrch, hasWindows)
		if err != nil {
			return nil, err
		}
		orchs = append(orchs,
			&OrchestratorVersionProfile{
				OrchestratorProfile: OrchestratorProfile{
					OrchestratorType:    Kubernetes,
					OrchestratorVersion: csOrch.OrchestratorVersion,
				},
				Default:  csOrch.OrchestratorVersion == common.GetDefaultKubernetesVersion(hasWindows),
				Upgrades: upgrades,
			})
	}
	return orchs, nil
}

func kubernetesUpgrades(csOrch *OrchestratorProfile, hasWindows bool) ([]*OrchestratorProfile, error) {
	ret := []*OrchestratorProfile{}

	upgradeVersions, err := getKubernetesAvailableUpgradeVersions(csOrch.OrchestratorVersion, common.GetAllSupportedKubernetesVersions(false, hasWindows))
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

func dcosInfo(csOrch *OrchestratorProfile, hasWindows bool) ([]*OrchestratorVersionProfile, error) {
	orchs := []*OrchestratorVersionProfile{}
	if csOrch.OrchestratorVersion == "" {
		// get info for all supported versions
		for _, ver := range common.AllDCOSSupportedVersions {
			upgrades, err := dcosUpgrades(&OrchestratorProfile{OrchestratorVersion: ver})
			if err != nil {
				return nil, err
			}
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
		if !isVersionSupported(csOrch) {
			return nil, errors.Errorf("DCOS version %s is not supported", csOrch.OrchestratorVersion)
		}

		// get info for the specified version
		upgrades, err := dcosUpgrades(csOrch)
		if err != nil {
			return nil, err
		}
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

func dcosUpgrades(csOrch *OrchestratorProfile) ([]*OrchestratorProfile, error) {
	ret := []*OrchestratorProfile{}

	switch csOrch.OrchestratorVersion {
	case common.DCOSVersion1Dot11Dot0:
		ret = append(ret, &OrchestratorProfile{
			OrchestratorType:    DCOS,
			OrchestratorVersion: common.DCOSVersion1Dot11Dot2,
		})
	}
	return ret, nil
}

func swarmInfo(csOrch *OrchestratorProfile, hasWindows bool) ([]*OrchestratorVersionProfile, error) {
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

	if !isVersionSupported(csOrch) {
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

func dockerceInfo(csOrch *OrchestratorProfile, hasWindows bool) ([]*OrchestratorVersionProfile, error) {

	if csOrch.OrchestratorVersion == "" {
		return []*OrchestratorVersionProfile{
			{
				OrchestratorProfile: OrchestratorProfile{
					OrchestratorType:    SwarmMode,
					OrchestratorVersion: DockerCEVersion,
				},
			},
		}, nil
	}

	if !isVersionSupported(csOrch) {
		return nil, errors.Errorf("Docker CE version %s is not supported", csOrch.OrchestratorVersion)
	}
	return []*OrchestratorVersionProfile{
		{
			OrchestratorProfile: OrchestratorProfile{
				OrchestratorType:    SwarmMode,
				OrchestratorVersion: csOrch.OrchestratorVersion,
			},
		},
	}, nil
}
