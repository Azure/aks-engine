// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"path"
	"testing"

	"github.com/leonelquinteros/gotext"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
)

func TestAssignParameters(t *testing.T) {
	// Initialize locale for translation
	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	i18n.Initialize(locale)

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}
	// iterate the test data directory
	apiModelTestFiles := &[]APIModelTestFile{}
	if e := IterateTestFilesDirectory(TestDataDir, apiModelTestFiles); e != nil {
		t.Error(e.Error())
		return
	}

	for _, tuple := range *apiModelTestFiles {
		containerService, _, err := apiloader.LoadContainerServiceFromFile(tuple.APIModelFilename, true, false, nil)
		if err != nil {
			t.Errorf("Loading file %s got error: %s", tuple.APIModelFilename, err.Error())
			continue
		}

		containerService.Location = "eastus"
		containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
			IsScale:    false,
			IsUpgrade:  false,
			PkiKeySize: helpers.DefaultPkiKeySize,
		})
		parametersMap := getParameters(containerService, DefaultGeneratorCode, "testversion")
		for k, v := range parametersMap {
			switch val := v.(paramsMap)["value"].(type) {
			case *bool:
				t.Errorf("got a pointer to bool in paramsMap value, this is dangerous!: %s: %v", k, val)
			}
		}
	}
}
func TestAssignVnetCidr(t *testing.T) {
	// Initialize locale for translation
	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	i18n.Initialize(locale)

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}
	// iterate the test data directory
	apiModelTestFiles := &[]APIModelTestFile{}
	if e := IterateTestFilesDirectory(TestDataDir, apiModelTestFiles); e != nil {
		t.Error(e.Error())
		return
	}

	for _, tuple := range *apiModelTestFiles {
		containerService, _, err := apiloader.LoadContainerServiceFromFile(tuple.APIModelFilename, true, false, nil)
		if err != nil {
			t.Errorf("Loading file %s got error: %s", tuple.APIModelFilename, err.Error())
			continue
		}

		containerService.Location = "eastus"
		containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
			IsScale:    false,
			IsUpgrade:  false,
			PkiKeySize: helpers.DefaultPkiKeySize,
		})
		containerService.Properties.AgentPoolProfiles = []*api.AgentPoolProfile{
			{
				Name: "default",
			},
			{
				Name: "second",
			},
		}

		// When HostedMasterProfile = nil
		containerService.Properties.HostedMasterProfile = nil
		containerService.Properties.AgentPoolProfiles[1].VnetCidrs = []string{"172.18.0.0/16", "10.240.0.0/16"}
		parametersMap := getParameters(containerService, DefaultGeneratorCode, "testversion")
		for k, v := range parametersMap {
			switch val := v.(paramsMap)["value"].(type) {
			case *bool:
				t.Errorf("got a pointer to bool in paramsMap value, this is dangerous!: %s: %v", k, val)
			}

			if k == "vnetCidr" {
				if containerService.Properties.MasterProfile != nil {
					if v.(paramsMap)["value"] != containerService.Properties.MasterProfile.VnetCidr {
						t.Errorf("vnetCidr %s is not set to the cidr %s.", v.(paramsMap)["value"], containerService.Properties.MasterProfile.VnetCidr)
					}
				} else {
					t.Errorf("vnetCidr should not be set. Value: %s", v.(paramsMap)["value"])
				}
			}
		}

		// When only second agent pool has VnetCidrs
		containerService.Properties.HostedMasterProfile = &api.HostedMasterProfile{}
		containerService.Properties.AgentPoolProfiles[1].VnetCidrs = []string{"172.18.0.0/16", "10.240.0.0/16"}
		isSetVnetCidr := false
		parametersMap = getParameters(containerService, DefaultGeneratorCode, "testversion")
		for k, v := range parametersMap {
			switch val := v.(paramsMap)["value"].(type) {
			case *bool:
				t.Errorf("got a pointer to bool in paramsMap value, this is dangerous!: %s: %v", k, val)
			}
			if k == "vnetCidr" {
				isSetVnetCidr = true
				if v.(paramsMap)["value"] != "172.18.0.0/16" {
					t.Errorf("vnetCidr is not set to the first cidr. %s", v.(paramsMap)["value"])
				}
			}
		}

		if !isSetVnetCidr {
			t.Error("vnetCidr is not added to parametersMap")
		}

		// When both agent pools have VnetCidrs. This is only used to validate that vnetCidr is only set once.
		// For AKS, all agent pools in the same cluster share a same VnetCidrs.
		containerService.Properties.AgentPoolProfiles[0].VnetCidrs = []string{"172.19.0.0/16", "10.241.0.0/16"}
		isSetVnetCidr = false
		parametersMap = getParameters(containerService, DefaultGeneratorCode, "testversion")
		for k, v := range parametersMap {
			switch val := v.(paramsMap)["value"].(type) {
			case *bool:
				t.Errorf("got a pointer to bool in paramsMap value, this is dangerous!: %s: %v", k, val)
			}
			if k == "vnetCidr" {
				isSetVnetCidr = true
				if v.(paramsMap)["value"] != "172.19.0.0/16" {
					t.Errorf("vnetCidr is not set to the first cidr of the first agent pool. %s", v.(paramsMap)["value"])
				}
			}
		}

		if !isSetVnetCidr {
			t.Error("vnetCidr is not added to parametersMap")
		}
	}
}
