// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"encoding/json"
	"io/ioutil"
	"reflect"

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
	cs, err := a.LoadContainerService(contents, version, validate, isUpdate, existingContainerService)
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
		if containerService.Properties.OrchestratorProfile == nil {
			containerService.Properties.OrchestratorProfile = &vlabs.OrchestratorProfile{}
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

// SerializeContainerService takes an unversioned container service and returns the bytes
func (a *Apiloader) SerializeContainerService(containerService *ContainerService, version string) ([]byte, error) {
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

// LoadCertificateProfileFromFile loads a CertificateProfile object from a JSON file
func (a *Apiloader) LoadCertificateProfileFromFile(jsonFile string) (*CertificateProfile, error) {
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, a.Translator.Errorf("error reading file %s: %s", jsonFile, err.Error())
	}
	return a.LoadCertificateProfile(content)
}

// LoadCertificateProfile marshalls raw data into a strongly typed CertificateProfile return object
func (a *Apiloader) LoadCertificateProfile(content []byte) (*CertificateProfile, error) {
	certificateProfile := &CertificateProfile{}
	if err := json.Unmarshal(content, &certificateProfile); err != nil {
		return nil, err
	}
	if err := checkJSONKeys(content, reflect.TypeOf(*certificateProfile), reflect.TypeOf(TypeMeta{})); err != nil {
		return nil, err
	}
	return certificateProfile, nil
}
