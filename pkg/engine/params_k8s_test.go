// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"path"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
)

func TestAssignKubernetesParameters(t *testing.T) {
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

		parametersMap := paramsMap{}
		containerService.Location = "eastus"
		cloudSpecConfig := containerService.GetCloudSpecConfig()
		containerService.SetPropertiesDefaults(false, false)
		if err := assignKubernetesParameters(containerService.Properties, parametersMap, cloudSpecConfig, DefaultGeneratorCode); err != nil {
			t.Error(err.Error())
			continue
		}
		for k, v := range parametersMap {
			switch val := v.(paramsMap)["value"].(type) {
			case *bool:
				t.Errorf("got a pointer to bool in paramsMap value, this is dangerous!: %s: %v", k, val)
			}
		}
	}
}

func TestAssignKubernetesComponentImageOverrideParams(t *testing.T) {
	p := &api.Properties{
		ServicePrincipalProfile: &api.ServicePrincipalProfile{},
		OrchestratorProfile: &api.OrchestratorProfile{
			OrchestratorType:    api.Kubernetes,
			OrchestratorVersion: "1.13.4",
			KubernetesConfig: &api.KubernetesConfig{
				KubernetesImageBase: "foo.com",
				ImageRepoOverrides: map[string]api.ImageRepoOverride{
					"foo.com/hyperkube-amd64": {Registry: "bar.com", Repo: "override/hyperkube"},
				},
			},
		},
	}

	// this is a manual mapping of overrides which listed in the config above
	// mapping to real component names, since there is no way to otherwise get
	// this mapping. Anything that is added or updated in the above config should
	// be added/modified accordingly here.
	overrideToComp := map[string]string{
		"foo.com/hyperkube-amd64": "hyperkube",
	}

	params := paramsMap{}
	if err := assignKubernetesParameters(p, params, api.AzureCloudSpec, DefaultGeneratorCode); err != nil {
		t.Fatal(err)
	}

	cloudSpecConfig := api.AzureEnvironmentSpecConfig{}

	for o := range p.OrchestratorProfile.KubernetesConfig.ImageRepoOverrides {
		name := overrideToComp[o]
		expect, err := api.GetKubernetesComponentImage(name, api.K8sComponentsByVersionMap[p.OrchestratorProfile.OrchestratorVersion], p.OrchestratorProfile.KubernetesConfig, false, cloudSpecConfig)
		if err != nil {
			t.Fatal(err)
		}

		actual := params[compToParamImage[name]].(paramsMap)["value"].(string)
		if actual != expect {
			t.Fatalf("expected %q, got %q", expect, actual)
		}
	}

	for comp, paramName := range compToParamImage {
		if comp == "hyperkube" {
			continue
		}

		pauseImageName := api.K8sComponentsByVersionMap[p.OrchestratorProfile.OrchestratorVersion][comp]
		param, ok := params[paramName]
		if !ok {
			t.Logf("skipping component %s", comp)
			continue
		}

		actual := param.(paramsMap)["value"].(string)
		expect := path.Join(p.OrchestratorProfile.KubernetesConfig.KubernetesImageBase, pauseImageName)
		if actual != expect {
			t.Errorf("expected %q, got %q", expect, actual)
		}
	}
}
