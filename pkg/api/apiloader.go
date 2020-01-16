// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"reflect"

	v20170831 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20170831"
	v20180331 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20180331"
	apvlabs "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/vlabs"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
)

const (
	defaultOrchestrator  = Kubernetes
	defaultAPIVersion    = vlabs.APIVersion
	defaultMasterCount   = 3
	defaultVMSize        = "Standard_DS2_v2"
	defaultOSDiskSizeGB  = 200
	defaultAgentPoolName = "agent"
	defaultAgentCount    = 3
	defaultAdminUser     = "azureuser"
)

// Apiloader represents the object that loads api model
type Apiloader struct {
	Translator *i18n.Translator
}

// LoadContainerServiceFromFile loads an AKS Cluster API Model from a JSON file
func (a *Apiloader) LoadContainerServiceFromFile(jsonFile string, validate, isUpdate bool, existingContainerService *ContainerService) (*ContainerService, string, error) {
	contents, e := ioutil.ReadFile(jsonFile)
	if e != nil {
		return nil, "", a.Translator.Errorf("error reading file %s: %s", jsonFile, e.Error())
	}
	return a.DeserializeContainerService(contents, validate, isUpdate, existingContainerService)
}

// LoadDefaultContainerServiceProperties loads the default API model
func LoadDefaultContainerServiceProperties() (TypeMeta, *vlabs.Properties) {
	return TypeMeta{APIVersion: defaultAPIVersion}, &vlabs.Properties{
		OrchestratorProfile: &vlabs.OrchestratorProfile{
			OrchestratorType: defaultOrchestrator,
		},
		MasterProfile: &vlabs.MasterProfile{
			Count:        defaultMasterCount,
			VMSize:       defaultVMSize,
			OSDiskSizeGB: defaultOSDiskSizeGB,
		},
		AgentPoolProfiles: []*vlabs.AgentPoolProfile{
			{
				Name:         defaultAgentPoolName,
				Count:        defaultAgentCount,
				VMSize:       defaultVMSize,
				OSDiskSizeGB: defaultOSDiskSizeGB,
			},
		},
		LinuxProfile: &vlabs.LinuxProfile{AdminUsername: defaultAdminUser},
	}
}

// DeserializeContainerService loads an AKS Engine Cluster API Model, validates it, and returns the unversioned representation
func (a *Apiloader) DeserializeContainerService(contents []byte, validate, isUpdate bool, existingContainerService *ContainerService) (*ContainerService, string, error) {
	m := &TypeMeta{}
	if err := json.Unmarshal(contents, &m); err != nil {
		return nil, "", err
	}

	version := m.APIVersion
	var cs *ContainerService
	var err error
	switch version {
	case "2017-08-31", "2018-03-31":
		cs, _, err = a.LoadContainerServiceForAgentPoolOnlyCluster(contents, version, validate, isUpdate, "", existingContainerService)
	default:
		cs, err = a.LoadContainerService(contents, version, validate, isUpdate, existingContainerService)
	}
	return cs, version, err
}

// LoadContainerService loads an AKS Cluster API Model, validates it, and returns the unversioned representation
func (a *Apiloader) LoadContainerService(
	contents []byte,
	version string,
	validate, isUpdate bool,
	existingContainerService *ContainerService) (*ContainerService, error) {
	var curOrchVersion string
	hasExistingCS := existingContainerService != nil
	if hasExistingCS {
		curOrchVersion = existingContainerService.Properties.OrchestratorProfile.OrchestratorVersion
	}
	switch version {
	case vlabs.APIVersion:
		containerService := &vlabs.ContainerService{}
		if e := json.Unmarshal(contents, &containerService); e != nil {
			return nil, e
		}
		if e := checkJSONKeys(contents, reflect.TypeOf(*containerService), reflect.TypeOf(TypeMeta{})); e != nil {
			return nil, e
		}
		if hasExistingCS {
			vecs := ConvertContainerServiceToVLabs(existingContainerService)
			if e := containerService.Merge(vecs); e != nil {
				return nil, e
			}
		}
		if validate {
			if e := containerService.Validate(isUpdate); e != nil {
				return nil, e
			}
		}

		var unversioned *ContainerService
		var err error
		if unversioned, err = ConvertVLabsContainerService(containerService, isUpdate); err != nil {
			return nil, err
		}
		if curOrchVersion != "" &&
			(containerService.Properties.OrchestratorProfile == nil ||
				(containerService.Properties.OrchestratorProfile.OrchestratorVersion == "" &&
					containerService.Properties.OrchestratorProfile.OrchestratorRelease == "")) {
			unversioned.Properties.OrchestratorProfile.OrchestratorVersion = curOrchVersion
		}
		return unversioned, nil

	default:
		return nil, a.Translator.Errorf("unrecognized APIVersion '%s'", version)
	}
}

// LoadContainerServiceForAgentPoolOnlyCluster loads an AKS Cluster API Model, validates it, and returns the unversioned representation
func (a *Apiloader) LoadContainerServiceForAgentPoolOnlyCluster(
	contents []byte,
	version string,
	validate, isUpdate bool,
	defaultKubernetesVersion string,
	existingContainerService *ContainerService) (*ContainerService, bool, error) {
	hasExistingCS := existingContainerService != nil
	IsSSHAutoGenerated := false
	hasWindows := false
	switch version {
	case v20170831.APIVersion:
		managedCluster := &v20170831.ManagedCluster{}
		if e := json.Unmarshal(contents, &managedCluster); e != nil {
			return nil, IsSSHAutoGenerated, e
		}
		// verify managedCluster.Properties is not nil for creating case
		if managedCluster.Properties == nil {
			if !isUpdate {
				return nil, false, a.Translator.Errorf("properties object in managed cluster should not be nil")
			}
			managedCluster.Properties = &v20170831.Properties{}
		}

		if hasExistingCS {
			vemc := ConvertContainerServiceToV20170831AgentPoolOnly(existingContainerService)
			if e := managedCluster.Merge(vemc); e != nil {
				return nil, IsSSHAutoGenerated, e
			}
		}

		// use defaultKubernetesVersion arg if no version was supplied in the request contents
		if managedCluster.Properties.KubernetesVersion == "" && defaultKubernetesVersion != "" {
			if !common.IsSupportedKubernetesVersion(defaultKubernetesVersion, isUpdate, hasWindows) {
				return nil, IsSSHAutoGenerated, a.Translator.Errorf("The selected orchestrator version '%s' is not supported", defaultKubernetesVersion)
			}
			managedCluster.Properties.KubernetesVersion = defaultKubernetesVersion
		}

		// verify orchestrator version
		if len(managedCluster.Properties.KubernetesVersion) > 0 && !common.IsSupportedKubernetesVersion(managedCluster.Properties.KubernetesVersion, isUpdate, hasWindows) {
			return nil, IsSSHAutoGenerated, a.Translator.Errorf("The selected orchestrator version '%s' is not supported", managedCluster.Properties.KubernetesVersion)
		}

		if validate {
			if e := managedCluster.Properties.Validate(); e != nil {
				return nil, IsSSHAutoGenerated, e
			}
		}

		return ConvertV20170831AgentPoolOnly(managedCluster), false, nil
	case v20180331.APIVersion:
		managedCluster := &v20180331.ManagedCluster{}
		if e := json.Unmarshal(contents, &managedCluster); e != nil {
			return nil, IsSSHAutoGenerated, e
		}
		// verify managedCluster.Properties is not nil for creating case
		if managedCluster.Properties == nil {
			if !isUpdate {
				return nil, false, a.Translator.Errorf("properties object in managed cluster should not be nil")
			}
			managedCluster.Properties = &v20180331.Properties{}
		}

		if hasExistingCS {
			vemc := ConvertContainerServiceToV20180331AgentPoolOnly(existingContainerService)
			if e := managedCluster.Merge(vemc); e != nil {
				return nil, IsSSHAutoGenerated, e
			}
		}

		// use defaultKubernetesVersion arg if no version was supplied in the request contents
		if managedCluster.Properties.KubernetesVersion == "" && defaultKubernetesVersion != "" {
			if !common.IsSupportedKubernetesVersion(defaultKubernetesVersion, isUpdate, hasWindows) {
				return nil, IsSSHAutoGenerated, a.Translator.Errorf("The selected orchestrator version '%s' is not supported", defaultKubernetesVersion)
			}
			if hasExistingCS {
				managedCluster.Properties.KubernetesVersion = existingContainerService.Properties.OrchestratorProfile.OrchestratorVersion
			} else {
				managedCluster.Properties.KubernetesVersion = defaultKubernetesVersion
			}
		}

		// verify orchestrator version
		if len(managedCluster.Properties.KubernetesVersion) > 0 && !common.IsSupportedKubernetesVersion(managedCluster.Properties.KubernetesVersion, isUpdate, hasWindows) {
			return nil, IsSSHAutoGenerated, a.Translator.Errorf("The selected orchestrator version '%s' is not supported", managedCluster.Properties.KubernetesVersion)
		}

		if validate {
			if e := managedCluster.Properties.Validate(); e != nil {
				return nil, IsSSHAutoGenerated, e
			}
		}

		// only generate ssh key on new cluster
		if !hasExistingCS && managedCluster.Properties.LinuxProfile == nil {
			linuxProfile := &v20180331.LinuxProfile{}
			linuxProfile.AdminUsername = "azureuser"
			_, publicKey, err := helpers.CreateSSH(rand.Reader, a.Translator)
			if err != nil {
				return nil, IsSSHAutoGenerated, err
			}
			linuxProfile.SSH.PublicKeys = []v20180331.PublicKey{{KeyData: publicKey}}
			managedCluster.Properties.LinuxProfile = linuxProfile
			IsSSHAutoGenerated = true
		}
		return ConvertV20180331AgentPoolOnly(managedCluster), IsSSHAutoGenerated, nil
	case apvlabs.APIVersion:
		managedCluster := &apvlabs.ManagedCluster{}
		if e := json.Unmarshal(contents, &managedCluster); e != nil {
			return nil, IsSSHAutoGenerated, e
		}
		if validate {
			if e := managedCluster.Properties.Validate(); e != nil {
				return nil, IsSSHAutoGenerated, e
			}
		}
		return ConvertVLabsAgentPoolOnly(managedCluster), IsSSHAutoGenerated, nil
	default:
		return nil, IsSSHAutoGenerated, a.Translator.Errorf("unrecognized APIVersion in LoadContainerServiceForAgentPoolOnlyCluster '%s'", version)
	}
}

// SerializeContainerService takes an unversioned container service and returns the bytes
func (a *Apiloader) SerializeContainerService(containerService *ContainerService, version string) ([]byte, error) {
	if containerService.Properties != nil && containerService.Properties.HostedMasterProfile != nil {
		b, err := a.serializeHostedContainerService(containerService, version)
		if err == nil && b != nil {
			return b, nil
		}
	}
	switch version {
	case vlabs.APIVersion:
		vlabsContainerService := ConvertContainerServiceToVLabs(containerService)
		armContainerService := &VlabsARMContainerService{}
		armContainerService.ContainerService = vlabsContainerService
		armContainerService.APIVersion = version
		b, err := helpers.JSONMarshalIndent(armContainerService, "", "  ", false)
		if err != nil {
			return nil, err
		}
		return b, nil

	default:
		return nil, a.Translator.Errorf("invalid version %s for conversion back from unversioned object", version)
	}
}

func (a *Apiloader) serializeHostedContainerService(containerService *ContainerService, version string) ([]byte, error) {
	switch version {
	case v20170831.APIVersion:
		v20170831ContainerService := ConvertContainerServiceToV20170831AgentPoolOnly(containerService)
		armContainerService := &V20170831ARMManagedContainerService{}
		armContainerService.ManagedCluster = v20170831ContainerService
		armContainerService.APIVersion = version
		b, err := helpers.JSONMarshalIndent(armContainerService, "", "  ", false)
		if err != nil {
			return nil, err
		}
		return b, nil
	case v20180331.APIVersion:
		v20180331ContainerService := ConvertContainerServiceToV20180331AgentPoolOnly(containerService)
		armContainerService := &V20180331ARMManagedContainerService{}
		armContainerService.ManagedCluster = v20180331ContainerService
		armContainerService.APIVersion = version
		b, err := helpers.JSONMarshalIndent(armContainerService, "", "  ", false)
		if err != nil {
			return nil, err
		}
		return b, nil
	default:
		return nil, a.Translator.Errorf("invalid version %s for conversion back from unversioned object", version)
	}
}

// LoadAgentpoolProfileFromFile loads an an AgentPoolProfile object from a JSON file
func (a *Apiloader) LoadAgentpoolProfileFromFile(jsonFile string) (*AgentPoolProfile, error) {
	contents, e := ioutil.ReadFile(jsonFile)
	if e != nil {
		return nil, a.Translator.Errorf("error reading file %s: %s", jsonFile, e.Error())
	}
	return a.LoadAgentPoolProfile(contents)
}

// LoadAgentPoolProfile marshalls raw data into a strongly typed AgentPoolProfile return object
func (a *Apiloader) LoadAgentPoolProfile(contents []byte) (*AgentPoolProfile, error) {
	agentPoolProfile := &AgentPoolProfile{}
	if e := json.Unmarshal(contents, &agentPoolProfile); e != nil {
		return nil, e
	}
	if e := checkJSONKeys(contents, reflect.TypeOf(*agentPoolProfile), reflect.TypeOf(TypeMeta{})); e != nil {
		return nil, e
	}
	return agentPoolProfile, nil
}
