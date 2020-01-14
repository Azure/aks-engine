// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"path"
	"path/filepath"
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

func TestWindowsProfileWithOSImageSettings(t *testing.T) {

	var tests = []struct {
		name           string // test case name
		windowsProfile api.WindowsProfile
		expectedResult map[string]string
	}{
		{
			"specify CustomImage",
			api.WindowsProfile{
				AKSOSImageVersion:     "",
				AdminUsername:         "FooUsername",
				AdminPassword:         "FooPassword",
				WindowsImageSourceURL: "FooWindowsImageSourceURL",
				WindowsPublisher:      "FooWindowsPublisher",
				WindowsOffer:          "FooWindowsOffer",
				WindowsSku:            "FooWindowsSku",
				ImageVersion:          "FooImageVersion",
			},
			map[string]string{
				"windowsAdminUsername":           "FooUsername",
				"windowsAdminPassword":           "FooPassword",
				"agentWindowsSourceUrl":          "FooWindowsImageSourceURL",
				"agentWindowsImageName":          "nil",
				"agentWindowsImageResourceGroup": "nil",
				"agentWindowsPublisher":          "nil",
				"agentWindowsOffer":              "nil",
				"agentWindowsSku":                "nil",
				"agentWindowsVersion":            "nil",
			},
		},
		{
			"specify ImageRef",
			api.WindowsProfile{
				AKSOSImageVersion: "",
				AdminUsername:     "FooUsername",
				AdminPassword:     "FooPassword",
				ImageRef: &api.ImageReference{
					Name:          "FooImageRefName",
					ResourceGroup: "FooImageRefResourceGroup",
				},
				WindowsPublisher: "FooWindowsPublisher",
				WindowsOffer:     "FooWindowsOffer",
				WindowsSku:       "FooWindowsSku",
				ImageVersion:     "FooImageVersion",
			},
			map[string]string{
				"windowsAdminUsername":           "FooUsername",
				"windowsAdminPassword":           "FooPassword",
				"agentWindowsSourceUrl":          "nil",
				"agentWindowsImageName":          "FooImageRefName",
				"agentWindowsImageResourceGroup": "FooImageRefResourceGroup",
				"agentWindowsPublisher":          "nil",
				"agentWindowsOffer":              "nil",
				"agentWindowsSku":                "nil",
				"agentWindowsVersion":            "nil",
			},
		},
		{
			"specify vhd",
			api.WindowsProfile{
				AKSOSImageVersion:    "",
				AdminUsername:        "FooUsername",
				AdminPassword:        "FooPassword",
				WindowsDockerVersion: "",
				WindowsPublisher:     "FooWindowsPublisher",
				WindowsOffer:         "FooWindowsOffer",
				WindowsSku:           "FooWindowsSku",
				ImageVersion:         "FooImageVersion",
			},
			map[string]string{
				"windowsAdminUsername":           "FooUsername",
				"windowsAdminPassword":           "FooPassword",
				"agentWindowsImageName":          "nil",
				"agentWindowsImageResourceGroup": "nil",
				"agentWindowsSourceUrl":          "nil",
				"agentWindowsPublisher":          "FooWindowsPublisher",
				"agentWindowsOffer":              "FooWindowsOffer",
				"agentWindowsSku":                "FooWindowsSku",
				"agentWindowsVersion":            "FooImageVersion",
			},
		},
		{
			"specify vhd with distro",
			api.WindowsProfile{
				AKSOSImageVersion:    api.AKSWindowsServer2019,
				AdminUsername:        "FooUsername",
				AdminPassword:        "FooPassword",
				WindowsDockerVersion: "",
				WindowsPublisher:     "",
				WindowsOffer:         "",
				WindowsSku:           "",
				ImageVersion:         "",
			},
			map[string]string{
				"windowsAdminUsername":           "FooUsername",
				"windowsAdminPassword":           "FooPassword",
				"agentWindowsImageName":          "nil",
				"agentWindowsImageResourceGroup": "nil",
				"agentWindowsSourceUrl":          "nil",
				"agentWindowsPublisher":          api.AKSWindowsServer2019OSImageConfig.ImagePublisher,
				"agentWindowsOffer":              api.AKSWindowsServer2019OSImageConfig.ImageOffer,
				"agentWindowsSku":                api.AKSWindowsServer2019OSImageConfig.ImageSku,
				"agentWindowsVersion":            api.AKSWindowsServer2019OSImageConfig.ImageVersion,
			},
		},
		{
			"do not specify vhd",
			api.WindowsProfile{
				AKSOSImageVersion: "",
				AdminUsername:     "FooUsername",
				AdminPassword:     "FooPassword",
				WindowsPublisher:  "",
				WindowsOffer:      "",
				WindowsSku:        "",
				ImageVersion:      "",
			},
			map[string]string{
				"windowsAdminUsername":           "FooUsername",
				"windowsAdminPassword":           "FooPassword",
				"agentWindowsImageName":          "nil",
				"agentWindowsImageResourceGroup": "nil",
				"agentWindowsSourceUrl":          "nil",
				"agentWindowsPublisher":          "",
				"agentWindowsOffer":              "",
				"agentWindowsSku":                "",
				"agentWindowsVersion":            "",
			},
		},
	}

	// Initialize locale for translation
	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	i18n.Initialize(locale)

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	apiModelFileName := filepath.Join(TestDataDir, "windows", "kubernetes-vmss.json")
	containerService, _, err := apiloader.LoadContainerServiceFromFile(apiModelFileName, true, false, nil)
	if err != nil {
		t.Fatalf("Loading file %s got error: %s", apiModelFileName, err.Error())
	}

	containerService.Location = "eastus"
	containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {

			containerService.Properties.WindowsProfile = &test.windowsProfile
			parametersMap := getParameters(containerService, DefaultGeneratorCode, "testversion")

			for k, v := range test.expectedResult {
				if v == "nil" {
					if parametersMap[k] != nil {
						t.Errorf("unexpected key for %s: expect non-exist but actual %s", k, parametersMap[k].(paramsMap)["value"])
					}
				} else {
					if parametersMap[k] == nil {
						t.Errorf("unexpected value for %s: expect %s but actual nil", k, v)
					} else if parametersMap[k].(paramsMap)["value"] != v {
						t.Errorf("unexpected value for %s: expect %s but actual %s", k, v, parametersMap[k].(paramsMap)["value"])
					}
				}
			}
		})
	}
}
