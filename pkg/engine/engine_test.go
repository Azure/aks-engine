// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
)

const (
	TestDataDir          = "./testdata"
	TestAKSEngineVersion = "1.0.0"
)

// LBRuleBaseString is a raw string that represents a template that we will compose an LB rule from
const LBRuleBaseString string = `	          {
            "name": "LBRule%d",
            "properties": {
              "backendAddressPool": {
                "id": "[concat(variables('%sLbID'), '/backendAddressPools/', variables('%sLbBackendPoolName'))]"
              },
              "backendPort": %d,
              "enableFloatingIP": false,
              "frontendIPConfiguration": {
                "id": "[variables('%sLbIPConfigID')]"
              },
              "frontendPort": %d,
              "idleTimeoutInMinutes": 5,
              "loadDistribution": "Default",
              "probe": {
                "id": "[concat(variables('%sLbID'),'/probes/tcp%dProbe')]"
              },
              "protocol": "Tcp"
            }
          }`

const TCPProbeBaseString string = `          {
            "name": "tcp%dProbe",
            "properties": {
              "intervalInSeconds": 5,
              "numberOfProbes": 2,
              "port": %d,
              "protocol": "Tcp"
            }
          }`

const securityRuleBaseString string = `          {
            "name": "Allow_%d",
            "properties": {
              "access": "Allow",
              "description": "Allow traffic from the Internet to port %d",
              "destinationAddressPrefix": "*",
              "destinationPortRange": "%d",
              "direction": "Inbound",
              "priority": %d,
              "protocol": "*",
              "sourceAddressPrefix": "Internet",
              "sourcePortRange": "*"
            }
          }`

func TestExpected(t *testing.T) {
	// Initialize locale for translation
	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	if err := i18n.Initialize(locale); err != nil {
		t.Error(err)
	}

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
		containerService, version, err := apiloader.LoadContainerServiceFromFile(tuple.APIModelFilename, true, false, nil)
		if err != nil {
			t.Errorf("Loading file %s got error: %s", tuple.APIModelFilename, err.Error())
			continue
		}

		if containerService.Properties.OrchestratorProfile.OrchestratorType != Kubernetes {
			// test the output container service 3 times:
			// 1. first time tests loaded containerService
			// 2. second time tests generated containerService
			// 3. third time tests the generated containerService from the generated containerService
			ctx := Context{
				Translator: &i18n.Translator{
					Locale: locale,
				},
			}
			templateGenerator, e3 := InitializeTemplateGenerator(ctx)
			if e3 != nil {
				t.Error(e3.Error())
				continue
			}

			certsGenerated, _ := containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
				IsScale:    false,
				IsUpgrade:  false,
				PkiKeySize: helpers.DefaultPkiKeySize,
			})
			if certsGenerated {
				t.Errorf("cert generation unexpected for %s", containerService.Properties.OrchestratorProfile.OrchestratorType)
			}

			armTemplate, params, err := templateGenerator.GenerateTemplate(containerService, DefaultGeneratorCode, TestAKSEngineVersion)
			if err != nil {
				t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, err.Error()))
				continue
			}

			expectedPpArmTemplate, e1 := transform.PrettyPrintArmTemplate(armTemplate)
			if e1 != nil {
				t.Error(armTemplate)
				t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e1.Error()))
				break
			}

			expectedPpParams, e2 := transform.PrettyPrintJSON(params)
			if e2 != nil {
				t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e2.Error()))
				continue
			}

			for i := 0; i < 3; i++ {
				if i > 0 {
					certsGenerated, _ = containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
						IsScale:    false,
						IsUpgrade:  false,
						PkiKeySize: helpers.DefaultPkiKeySize,
					})
					if certsGenerated {
						t.Errorf("cert generation unexpected for %s", containerService.Properties.OrchestratorProfile.OrchestratorType)
					}
				}
				armTemplate, params, err := templateGenerator.GenerateTemplate(containerService, DefaultGeneratorCode, TestAKSEngineVersion)
				if err != nil {
					t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, err.Error()))
					continue
				}
				generatedPpArmTemplate, e1 := transform.PrettyPrintArmTemplate(armTemplate)
				if e1 != nil {
					t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e1.Error()))
					continue
				}

				generatedPpParams, e2 := transform.PrettyPrintJSON(params)
				if e2 != nil {
					t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e2.Error()))
					continue
				}

				if !bytes.Equal([]byte(expectedPpArmTemplate), []byte(generatedPpArmTemplate)) {
					diffstr, differr := tuple.WriteArmTemplateErrFilename([]byte(generatedPpArmTemplate))
					if differr != nil {
						diffstr += differr.Error()
					}
					t.Errorf("generated output different from expected for model %s: '%s'", tuple.APIModelFilename, diffstr)
				}

				if !bytes.Equal([]byte(expectedPpParams), []byte(generatedPpParams)) {
					diffstr, differr := tuple.WriteArmTemplateParamsErrFilename([]byte(generatedPpParams))
					if differr != nil {
						diffstr += differr.Error()
					}
					t.Errorf("generated parameters different from expected for model %s: '%s'", tuple.APIModelFilename, diffstr)
				}

				b, err := apiloader.SerializeContainerService(containerService, version)
				if err != nil {
					t.Error(err)
				}
				containerService, version, err = apiloader.DeserializeContainerService(b, true, false, nil)
				if err != nil {
					t.Error(err)
				}
				if version != vlabs.APIVersion {
					// Set CertificateProfile here to avoid a new one generated.
					// Kubernetes template needs certificate profile to match expected template
					// API versions other than vlabs don't expose CertificateProfile
					containerService.Properties.CertificateProfile = &api.CertificateProfile{}
					addTestCertificateProfile(containerService.Properties.CertificateProfile)
				}
			}
		} else {
			if version != vlabs.APIVersion {
				// Set CertificateProfile here to avoid a new one generated.
				// Kubernetes template needs certificate profile to match expected template
				// API versions other than vlabs don't expose CertificateProfile
				containerService.Properties.CertificateProfile = &api.CertificateProfile{}
				addTestCertificateProfile(containerService.Properties.CertificateProfile)
			}

			// test the output container service 3 times:
			// 1. first time tests loaded containerService
			// 2. second time tests generated containerService
			// 3. third time tests the generated containerService from the generated containerService
			ctx := Context{
				Translator: &i18n.Translator{
					Locale: locale,
				},
			}
			templateGenerator, e3 := InitializeTemplateGenerator(ctx)
			if e3 != nil {
				t.Error(e3.Error())
				continue
			}

			certsGenerated, _ := containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
				IsScale:    false,
				IsUpgrade:  false,
				PkiKeySize: helpers.DefaultPkiKeySize,
			})
			if certsGenerated {
				t.Errorf("cert generation unexpected for %s, apiversion: %s, path: %s ", containerService.Properties.OrchestratorProfile.OrchestratorType, version, tuple.APIModelFilename)
			}

			armTemplate, params, err := templateGenerator.GenerateTemplateV2(containerService, DefaultGeneratorCode, TestAKSEngineVersion)
			if err != nil {
				t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, err.Error()))
				continue
			}

			expectedPpArmTemplate, e1 := transform.PrettyPrintArmTemplate(armTemplate)
			if e1 != nil {
				t.Error(armTemplate)
				t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e1.Error()))
				break
			}

			expectedPpParams, e2 := transform.PrettyPrintJSON(params)
			if e2 != nil {
				t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e2.Error()))
				continue
			}

			for i := 0; i < 3; i++ {
				if i > 0 {
					certsGenerated, _ = containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
						IsScale:    false,
						IsUpgrade:  false,
						PkiKeySize: helpers.DefaultPkiKeySize,
					})
					if certsGenerated {
						t.Errorf("cert generation unexpected for %s", containerService.Properties.OrchestratorProfile.OrchestratorType)
					}
				}

				armTemplate, params, err := templateGenerator.GenerateTemplateV2(containerService, DefaultGeneratorCode, TestAKSEngineVersion)
				if err != nil {
					t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, err.Error()))
					continue
				}
				generatedPpArmTemplate, e1 := transform.PrettyPrintArmTemplate(armTemplate)
				if e1 != nil {
					t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e1.Error()))
					continue
				}

				generatedPpParams, e2 := transform.PrettyPrintJSON(params)
				if e2 != nil {
					t.Error(errors.Errorf("error in file %s: %s", tuple.APIModelFilename, e2.Error()))
					continue
				}

				if !bytes.Equal([]byte(expectedPpArmTemplate), []byte(generatedPpArmTemplate)) {
					diffstr, differr := tuple.WriteArmTemplateErrFilename([]byte(generatedPpArmTemplate))
					if differr != nil {
						diffstr += differr.Error()
					}
					t.Errorf("generated output different from expected for model %s: '%s'", tuple.APIModelFilename, diffstr)
				}

				if !bytes.Equal([]byte(expectedPpParams), []byte(generatedPpParams)) {
					diffstr, differr := tuple.WriteArmTemplateParamsErrFilename([]byte(generatedPpParams))
					if differr != nil {
						diffstr += differr.Error()
					}
					t.Errorf("generated parameters different from expected for model %s: '%s'", tuple.APIModelFilename, diffstr)
				}

				b, err := apiloader.SerializeContainerService(containerService, version)
				if err != nil {
					t.Error(err)
				}
				containerService, version, err = apiloader.DeserializeContainerService(b, true, false, nil)
				if err != nil {
					t.Error(err)
				}
				if version != vlabs.APIVersion {
					// Set CertificateProfile here to avoid a new one generated.
					// Kubernetes template needs certificate profile to match expected template
					// API versions other than vlabs don't expose CertificateProfile
					containerService.Properties.CertificateProfile = &api.CertificateProfile{}
					addTestCertificateProfile(containerService.Properties.CertificateProfile)
				}
			}
		}
	}
}

// APIModelTestFile holds the test file name and knows how to find the expected files
type APIModelTestFile struct {
	APIModelFilename string
}

// WriteArmTemplateErrFilename writes out an error file to sit parallel for comparison
func (a *APIModelTestFile) WriteArmTemplateErrFilename(contents []byte) (string, error) {
	filename := fmt.Sprintf("%s_expected.err", a.APIModelFilename)
	if err := ioutil.WriteFile(filename, contents, 0600); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s written for diff", filename), nil
}

// WriteArmTemplateParamsErrFilename writes out an error file to sit parallel for comparison
func (a *APIModelTestFile) WriteArmTemplateParamsErrFilename(contents []byte) (string, error) {
	filename := fmt.Sprintf("%s_expected_params.err", a.APIModelFilename)
	if err := ioutil.WriteFile(filename, contents, 0600); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s written for diff", filename), nil
}

// IterateTestFilesDirectory iterates the test data directory adding api model files to the test file slice.
func IterateTestFilesDirectory(directory string, apiModelTestFiles *[]APIModelTestFile) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if e := IterateTestFilesDirectory(filepath.Join(directory, file.Name()), apiModelTestFiles); e != nil {
				return e
			}
		} else {
			if !strings.Contains(file.Name(), "_expected") && strings.HasSuffix(file.Name(), ".json") {
				tuple := &APIModelTestFile{}
				tuple.APIModelFilename = filepath.Join(directory, file.Name())
				*apiModelTestFiles = append(*apiModelTestFiles, *tuple)
			}
		}
	}
	return nil
}

// addTestCertificateProfile add certificate artifacts for test purpose
func addTestCertificateProfile(api *api.CertificateProfile) {
	api.CaCertificate = "caCertificate"
	api.CaPrivateKey = "caPrivateKey"
	api.APIServerCertificate = "apiServerCertificate"
	api.APIServerPrivateKey = "apiServerPrivateKey"
	api.ClientCertificate = "clientCertificate"
	api.ClientPrivateKey = "clientPrivateKey"
	api.KubeConfigCertificate = "kubeConfigCertificate"
	api.KubeConfigPrivateKey = "kubeConfigPrivateKey"
	api.EtcdClientCertificate = "etcdClientCertificate"
	api.EtcdClientPrivateKey = "etcdClientPrivateKey"
	api.EtcdServerCertificate = "etcdServerCertificate"
	api.EtcdServerPrivateKey = "etcdServerPrivateKey"
	api.EtcdPeerCertificates = []string{"etcdPeerCertificate0"}
	api.EtcdPeerPrivateKeys = []string{"etcdPeerPrivateKey0"}
}

type TestARMTemplate struct {
	Outputs map[string]OutputElement `json:"outputs"`
	//Parameters *json.RawMessage `json:"parameters"`
	//Resources  *json.RawMessage `json:"resources"`
	//Variables  *json.RawMessage `json:"variables"`
}

type OutputElement struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func TestTemplateOutputPresence(t *testing.T) {
	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	if err := i18n.Initialize(locale); err != nil {
		t.Error(err)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	ctx := Context{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	templateGenerator, err := InitializeTemplateGenerator(ctx)

	if err != nil {
		t.Fatalf("Failed to initialize template generator: %v", err)
	}

	containerService, _, err := apiloader.LoadContainerServiceFromFile("./testdata/simple/kubernetes.json", true, false, nil)
	if err != nil {
		t.Fatalf("Failed to load container service from file: %v", err)
	}
	_, err = containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		t.Error(err)
	}
	armTemplate, _, err := templateGenerator.GenerateTemplateV2(containerService, DefaultGeneratorCode, TestAKSEngineVersion)
	if err != nil {
		t.Fatalf("Failed to generate arm template: %v", err)
	}

	var template TestARMTemplate
	err = json.Unmarshal([]byte(armTemplate), &template)
	if err != nil {
		t.Fatalf("couldn't unmarshall ARM template: %#v\n", err)
	}

	tt := []struct {
		key   string
		value string
	}{
		{key: "resourceGroup", value: "[variables('resourceGroup')]"},
		{key: "subnetName", value: "[variables('subnetName')]"},
		{key: "securityGroupName", value: "[variables('nsgName')]"},
		{key: "virtualNetworkName", value: "[variables('virtualNetworkName')]"},
		{key: "routeTableName", value: "[variables('routeTableName')]"},
		{key: "primaryAvailabilitySetName", value: "[variables('primaryAvailabilitySetName')]"},
	}

	for _, tc := range tt {
		element, found := template.Outputs[tc.key]
		if !found {
			t.Fatalf("Output key %v not found", tc.key)
		} else if element.Value != tc.value {
			t.Fatalf("Expected %q at key %v but got: %q", tc.value, tc.key, element.Value)
		}
	}
}

func TestIsNSeriesSKU(t *testing.T) {
	// VMSize with GPU
	validSkus := []string{
		"Standard_NC12",
		"Standard_NC12s_v2",
		"Standard_NC12s_v3",
		"Standard_NC24",
		"Standard_NC24r",
		"Standard_NC24rs_v2",
		"Standard_NC24rs_v3",
		"Standard_NC24s_v2",
		"Standard_NC24s_v3",
		"Standard_NC6",
		"Standard_NC6s_v2",
		"Standard_NC6s_v3",
		"Standard_ND12s",
		"Standard_ND24rs",
		"Standard_ND24s",
		"Standard_ND6s",
		"Standard_NV12",
		"Standard_NV24",
		"Standard_NV6",
		"Standard_NV24r",
	}

	invalidSkus := []string{
		"Standard_A10",
		"Standard_A11",
		"Standard_A2",
		"Standard_A2_v2",
		"Standard_A2m_v2",
		"Standard_A3",
		"Standard_A4",
		"Standard_A4_v2",
		"Standard_A4m_v2",
		"Standard_A5",
		"Standard_A6",
		"Standard_A7",
		"Standard_A8",
		"Standard_A8_v2",
		"Standard_A8m_v2",
		"Standard_A9",
		"Standard_B2ms",
		"Standard_B4ms",
		"Standard_B8ms",
		"Standard_D11",
		"Standard_D11_v2",
		"Standard_D12",
		"Standard_D12_v2",
		"Standard_D13",
		"Standard_D13_v2",
		"Standard_D14",
		"Standard_D14_v2",
		"Standard_D15_v2",
		"Standard_D16_v3",
		"Standard_D16s_v3",
		"Standard_D2",
		"Standard_D2_v2",
		"Standard_D2_v3",
		"Standard_D2s_v3",
		"Standard_D3",
		"Standard_D32_v3",
		"Standard_D32s_v3",
		"Standard_D3_v2",
		"Standard_D4",
		"Standard_D4_v2",
		"Standard_D4_v3",
		"Standard_D4s_v3",
		"Standard_D5_v2",
		"Standard_D64_v3",
		"Standard_D64s_v3",
		"Standard_D8_v3",
		"Standard_D8s_v3",
		"Standard_DS11",
		"Standard_DS11_v2",
		"Standard_DS12",
		"Standard_DS12_v2",
		"Standard_DS13",
		"Standard_DS13-2_v2",
		"Standard_DS13-4_v2",
		"Standard_DS13_v2",
		"Standard_DS14",
		"Standard_DS14-4_v2",
		"Standard_DS14-8_v2",
		"Standard_DS14_v2",
		"Standard_DS15_v2",
		"Standard_DS3",
		"Standard_DS3_v2",
		"Standard_DS4",
		"Standard_DS4_v2",
		"Standard_DS5_v2",
		"Standard_E16_v3",
		"Standard_E16s_v3",
		"Standard_E2_v3",
		"Standard_E2s_v3",
		"Standard_E32-16s_v3",
		"Standard_E32-8s_v3",
		"Standard_E32_v3",
		"Standard_E32s_v3",
		"Standard_E4_v3",
		"Standard_E4s_v3",
		"Standard_E64-16s_v3",
		"Standard_E64-32s_v3",
		"Standard_E64_v3",
		"Standard_E64s_v3",
		"Standard_E8_v3",
		"Standard_E8s_v3",
		"Standard_F16",
		"Standard_F16s",
		"Standard_F16s_v2",
		"Standard_F2",
		"Standard_F2s_v2",
		"Standard_F32s_v2",
		"Standard_F4",
		"Standard_F4s",
		"Standard_F4s_v2",
		"Standard_F64s_v2",
		"Standard_F72s_v2",
		"Standard_F8",
		"Standard_F8s",
		"Standard_F8s_v2",
		"Standard_G1",
		"Standard_G2",
		"Standard_G3",
		"Standard_G4",
		"Standard_G5",
		"Standard_GS1",
		"Standard_GS2",
		"Standard_GS3",
		"Standard_GS4",
		"Standard_GS4-4",
		"Standard_GS4-8",
		"Standard_GS5",
		"Standard_GS5-16",
		"Standard_GS5-8",
		"Standard_H16",
		"Standard_H16m",
		"Standard_H16mr",
		"Standard_H16r",
		"Standard_H8",
		"Standard_H8m",
		"Standard_L16s",
		"Standard_L32s",
		"Standard_L4s",
		"Standard_L8s",
		"Standard_M128-32ms",
		"Standard_M128-64ms",
		"Standard_M128ms",
		"Standard_M128s",
		"Standard_M64-16ms",
		"Standard_M64-32ms",
		"Standard_M64ms",
		"Standard_M64s",
	}

	for _, sku := range validSkus {
		if !common.IsNvidiaEnabledSKU(sku) {
			t.Fatalf("Expected common.IsNvidiaEnabledSKU(%s) to be true", sku)
		}
	}

	for _, sku := range invalidSkus {
		if common.IsNvidiaEnabledSKU(sku) {
			t.Fatalf("Expected common.IsNvidiaEnabledSKU(%s) to be false", sku)
		}
	}
}

func TestGenerateKubeConfig(t *testing.T) {
	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	if err := i18n.Initialize(locale); err != nil {
		t.Error(err)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	testData := "./testdata/simple/kubernetes.json"

	containerService, _, err := apiloader.LoadContainerServiceFromFile(testData, true, false, nil)
	if err != nil {
		t.Errorf("Failed to load container service from file: %v", err)
	}
	kubeConfig, err := GenerateKubeConfig(containerService.Properties, "westus2")
	// TODO add actual kubeconfig validation
	if len(kubeConfig) < 1 {
		t.Errorf("Got unexpected kubeconfig payload: %v", kubeConfig)
	}
	if err != nil {
		t.Errorf("Failed to call GenerateKubeConfig with simple Kubernetes config from file: %v", testData)
	}

	p := api.Properties{}
	_, err = GenerateKubeConfig(&p, "westus2")
	if err == nil {
		t.Errorf("Expected an error result from nil Properties child properties")
	}

	_, err = GenerateKubeConfig(nil, "westus2")
	if err == nil {
		t.Errorf("Expected an error result from nil Properties child properties")
	}

	containerService.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster = &api.PrivateCluster{
		Enabled: to.BoolPtr(true),
	}

	_, err = GenerateKubeConfig(containerService.Properties, "westus2")
	if err != nil {
		t.Errorf("Failed to call GenerateKubeConfig with simple Kubernetes config from file: %v", testData)
	}

	containerService.Properties.MasterProfile.Count = 3
	_, err = GenerateKubeConfig(containerService.Properties, "westus2")
	if err == nil {
		t.Errorf("expected an error result when Private Cluster is Enabled and no FirstConsecutiveStaticIP was specified")
	}

	containerService.Properties.MasterProfile.FirstConsecutiveStaticIP = "10.239.255.239"
	_, err = GenerateKubeConfig(containerService.Properties, "westus2")
	if err != nil {
		t.Errorf("Failed to call GenerateKubeConfig with simple Kubernetes config from file: %v", testData)
	}

	containerService.Properties.AADProfile = &api.AADProfile{
		ClientAppID: "fooClientAppID",
		TenantID:    "fooTenantID",
		ServerAppID: "fooServerAppID",
	}

	_, err = GenerateKubeConfig(containerService.Properties, "westus2")
	if err != nil {
		t.Errorf("Failed to call GenerateKubeConfig with simple Kubernetes config from file: %v", testData)
	}
}

func TestValidateDistro(t *testing.T) {
	// Test with Invalid Master Profile
	cs := &api.ContainerService{
		Properties: &api.Properties{
			MasterProfile: &api.MasterProfile{
				Distro: "rhel",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: "Kubernetes",
			},
		},
	}

	result := validateDistro(cs)

	if result {
		t.Errorf("expected validateDistro to return false for Kubernetes type")
	}

	// Test with invalid Agent Pool Profile
	cs.Properties.MasterProfile.Distro = "coreos"
	cs.Properties.AgentPoolProfiles = []*api.AgentPoolProfile{
		{
			Distro: "rhel",
		},
	}

	result = validateDistro(cs)

	if result {
		t.Errorf("expected validateDistro to return false for Kubernetes type")
	}
}

func TestMakeMasterExtensionScriptCommands(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ExtensionProfiles: []*api.ExtensionProfile{
				{
					Name:                "fooExtension",
					RootURL:             "fooRootURL",
					Version:             "1.0",
					Script:              "fooBar Script",
					URLQuery:            "fooURLQuery",
					ExtensionParameters: "fooExtensionParams",
				},
			},
		},
	}

	// Windows profile
	profile := &api.AgentPoolProfile{
		OSType: "Windows",
		PreprovisionExtension: &api.Extension{
			Name: "fooExtension",
		},
	}

	actual := makeAgentExtensionScriptCommands(cs, profile)

	expected := `New-Item -ItemType Directory -Force -Path "$env:SystemDrive:/AzureData/extensions/fooExtension" ; Invoke-WebRequest -Uri "fooRootURLextensions/fooExtension/1.0/fooBar Script?fooURLQuery" -OutFile "$env:SystemDrive:/AzureData/extensions/fooExtension/fooBar Script" ; powershell "$env:SystemDrive:/AzureData/extensions/fooExtension/fooBar Script ` + "`" + `"',parameters('fooExtensionParameters'),'` + "`" + `""
`

	if actual != expected {
		t.Errorf("expected to get %s, but got %s instead", expected, actual)
	}

	// Linux profile
	profile.OSType = "Linux"

	actual = makeAgentExtensionScriptCommands(cs, profile)

	expected = `- sudo /usr/bin/curl --retry 5 --retry-delay 10 --retry-max-time 30 -o /opt/azure/containers/extensions/fooExtension/fooBar Script --create-dirs  "fooRootURLextensions/fooExtension/1.0/fooBar Script?fooURLQuery" 
- sudo /bin/chmod 744 /opt/azure/containers/extensions/fooExtension/fooBar Script 
- sudo /opt/azure/containers/extensions/fooExtension/fooBar Script ',parameters('fooExtensionParameters'),' > /var/log/fooExtension-output.log`

	if actual != expected {
		t.Errorf("expected to get %s, but got %s instead", expected, actual)
	}

	// Azure Stack Linux profile
	cs.Properties.CustomCloudProfile = &api.CustomCloudProfile{}

	actual = makeAgentExtensionScriptCommands(cs, profile)

	expected = `- sudo /usr/bin/curl --retry 5 --retry-delay 10 --retry-max-time 30 -o /opt/azure/containers/extensions/fooExtension/fooBar Script --create-dirs --cacert /var/lib/waagent/Certificates.pem "fooRootURLextensions/fooExtension/1.0/fooBar Script?fooURLQuery" 
- sudo /bin/chmod 744 /opt/azure/containers/extensions/fooExtension/fooBar Script 
- sudo /opt/azure/containers/extensions/fooExtension/fooBar Script ',parameters('fooExtensionParameters'),' > /var/log/fooExtension-output.log`

	if actual != expected {
		t.Errorf("expected to get %s, but got %s instead", expected, actual)
	}
}

func TestGetDCOSWindowsAgentPreprovisionParameters(t *testing.T) {
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ExtensionProfiles: []*api.ExtensionProfile{
				{
					Name:                "fooExtension",
					ExtensionParameters: "fooExtensionParams",
				},
			},
		},
	}

	profile := &api.AgentPoolProfile{
		PreprovisionExtension: &api.Extension{
			Name: "fooExtension",
		},
	}

	actual := getDCOSWindowsAgentPreprovisionParameters(cs, profile)

	expected := "fooExtensionParams"

	if actual != expected {
		t.Errorf("expected to get %s, but got %s instead", expected, actual)
	}
}

func TestGetDCOSWindowsAgentCustomAttributes(t *testing.T) {
	profile := &api.AgentPoolProfile{
		OSType: api.Windows,
		Ports: []int{
			8000,
			8080,
		},
		CustomNodeLabels: map[string]string{
			"foo":   "bar",
			"abc":   "xyz",
			"lorem": "ipsum",
		},
	}

	actual := getDCOSWindowsAgentCustomAttributes(profile)

	if !strings.Contains(actual, "os:Windows;public_ip:yes;") {
		t.Errorf("expected output string of getDCOSWindowsAgentCustomAttributes %s to contain os:Windows;public_ip:yes;", actual)
	}

	for k, v := range profile.CustomNodeLabels {
		if !strings.Contains(actual, fmt.Sprintf("%s:%s", k, v)) {
			t.Errorf("expected output string of getDCOSWindowsAgentCustomAttributes %s to contain key-value pairs %s:%s", actual, k, v)
		}
	}
}

func TestGetKubernetesSubnets(t *testing.T) {
	props := &api.Properties{
		AgentPoolProfiles: []*api.AgentPoolProfile{
			{
				OSType: api.Windows,
				Count:  1,
			},
		},
		MasterProfile: &api.MasterProfile{
			Count: 1,
		},
	}

	actual := getKubernetesSubnets(props)

	expected := `,
{
            "name": "podCIDR2",
            "properties": {
              "addressPrefix": "10.244.2.0/24",
              "networkSecurityGroup": {
                "id": "[variables('nsgID')]"
              },
              "routeTable": {
                "id": "[variables('routeTableID')]"
              }
            }
          }`

	if actual != expected {
		t.Errorf("expected to get %s, but got %s instead", expected, actual)
	}
}

func TestGetVNETSubnetDependencies(t *testing.T) {
	baseString := `        "[concat('Microsoft.Network/networkSecurityGroups/', variables('%sNSGName'))]"`
	cases := []struct {
		p        *api.Properties
		expected string
	}{
		{
			p:        &api.Properties{},
			expected: "",
		},
		{
			p: &api.Properties{
				AgentPoolProfiles: []*api.AgentPoolProfile{},
			},
			expected: "",
		},
		{
			p: &api.Properties{
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						Name: "pool1",
					},
				},
			},
			expected: fmt.Sprintf(baseString, "pool1"),
		},
		{
			p: &api.Properties{
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						Name: "pool1",
					},
					{
						Name: "pool2",
					},
					{
						Name: "pool3",
					},
				},
			},
			expected: fmt.Sprintf(baseString, "pool1") + ",\n" + fmt.Sprintf(baseString, "pool2") + ",\n" + fmt.Sprintf(baseString, "pool3"),
		},
	}

	for _, c := range cases {
		if getVNETSubnetDependencies(c.p) != c.expected {
			t.Fatalf("expected getVNETSubnetDependencies() to return %s but instead got %s", c.expected, getVNETSubnetDependencies(c.p))
		}
	}
}

func TestGetLBRule(t *testing.T) {
	cases := []struct {
		name     string
		port     int
		expected string
	}{
		{
			name:     "foo",
			port:     80,
			expected: fmt.Sprintf(LBRuleBaseString, 80, "foo", "foo", 80, "foo", 80, "foo", 80),
		},
		{
			name:     "bar",
			port:     8080,
			expected: fmt.Sprintf(LBRuleBaseString, 8080, "bar", "bar", 8080, "bar", 8080, "bar", 8080),
		},
		{
			name:     "",
			port:     0,
			expected: fmt.Sprintf(LBRuleBaseString, 0, "", "", 0, "", 0, "", 0),
		},
	}

	for _, c := range cases {
		if getLBRule(c.name, c.port) != c.expected {
			t.Fatalf("expected getLBRule() to return %s but instead got %s", c.expected, getLBRule(c.name, c.port))
		}
	}
}

func TestGetLBRules(t *testing.T) {
	cases := []struct {
		name     string
		ports    []int
		expected string
	}{
		{
			name:     "foo",
			ports:    []int{80},
			expected: fmt.Sprintf(LBRuleBaseString, 80, "foo", "foo", 80, "foo", 80, "foo", 80),
		},
		{
			name:     "bar",
			ports:    []int{8080},
			expected: fmt.Sprintf(LBRuleBaseString, 8080, "bar", "bar", 8080, "bar", 8080, "bar", 8080),
		},
		{
			name:     "baz",
			ports:    []int{80, 8080},
			expected: fmt.Sprintf(LBRuleBaseString, 80, "baz", "baz", 80, "baz", 80, "baz", 80) + ",\n" + fmt.Sprintf(LBRuleBaseString, 8080, "baz", "baz", 8080, "baz", 8080, "baz", 8080),
		},
		{
			name:     "",
			ports:    []int{},
			expected: "",
		},
	}

	for _, c := range cases {
		if getLBRules(c.name, c.ports) != c.expected {
			t.Fatalf("expected getLBRules() to return %s but instead got %s", c.expected, getLBRules(c.name, c.ports))
		}
	}
}

func TestGetProbe(t *testing.T) {
	cases := []struct {
		port     int
		expected string
	}{
		{
			port:     80,
			expected: fmt.Sprintf(TCPProbeBaseString, 80, 80),
		},
		{
			port:     0,
			expected: fmt.Sprintf(TCPProbeBaseString, 0, 0),
		},
	}

	for _, c := range cases {
		if getProbe(c.port) != c.expected {
			t.Fatalf("expected getProbe() to return %s but instead got %s", c.expected, getProbe(c.port))
		}
	}
}

func TestGetProbes(t *testing.T) {
	cases := []struct {
		ports    []int
		expected string
	}{
		{
			ports:    []int{80, 81},
			expected: fmt.Sprintf(TCPProbeBaseString, 80, 80) + ",\n" + fmt.Sprintf(TCPProbeBaseString, 81, 81),
		},
		{
			ports:    []int{8080},
			expected: fmt.Sprintf(TCPProbeBaseString, 8080, 8080),
		},
	}

	for _, c := range cases {
		if getProbes(c.ports) != c.expected {
			t.Fatalf("expected getProbes() to return %s but instead got %s", c.expected, getProbes(c.ports))
		}
	}
}

func TestGetSecurityRule(t *testing.T) {
	cases := []struct {
		port      int
		portIndex int
		expected  string
	}{
		{
			port:      80,
			portIndex: 0,
			expected:  fmt.Sprintf(securityRuleBaseString, 80, 80, 80, 200),
		},
		{
			port:      0,
			portIndex: 1,
			expected:  fmt.Sprintf(securityRuleBaseString, 0, 0, 0, 201),
		},
	}

	for _, c := range cases {
		if getSecurityRule(c.port, c.portIndex) != c.expected {
			t.Fatalf("expected getSecurityRule() to return %s but instead got %s", c.expected, getSecurityRule(c.port, c.portIndex))
		}
	}
}

func TestGetSecurityRules(t *testing.T) {
	cases := []struct {
		ports    []int
		expected string
	}{
		{
			ports:    []int{80, 81},
			expected: fmt.Sprintf(securityRuleBaseString, 80, 80, 80, 200) + ",\n" + fmt.Sprintf(securityRuleBaseString, 81, 81, 81, 201),
		},
		{
			ports:    []int{80},
			expected: fmt.Sprintf(securityRuleBaseString, 80, 80, 80, 200),
		},
	}

	for _, c := range cases {
		if getSecurityRules(c.ports) != c.expected {
			t.Fatalf("expected getSecurityRules() to return %s but instead got %s", c.expected, getSecurityRules(c.ports))
		}
	}
}

func TestGetVNETAddressPrefixes(t *testing.T) {
	baseString := `"[variables('masterSubnet')]"`
	cases := []struct {
		p        *api.Properties
		expected string
	}{
		{
			p: &api.Properties{
				MasterProfile: &api.MasterProfile{},
			},
			expected: baseString,
		},
		{
			p: &api.Properties{
				MasterProfile: &api.MasterProfile{},
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						Name:   "foo",
						Subnet: "10.0.0.0/24",
					},
				},
			},
			expected: baseString + fmt.Sprintf(",\n            \"[variables('%sSubnet')]\"", "foo"),
		},
		{
			p: &api.Properties{
				MasterProfile: &api.MasterProfile{},
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						Name:   "foo",
						Subnet: "10.0.0.0/24",
					},
					{
						Name:   "bar",
						Subnet: "10.0.0.0/24",
					},
				},
			},
			expected: baseString + fmt.Sprintf(",\n            \"[variables('%sSubnet')]\"", "foo") + fmt.Sprintf(",\n            \"[variables('%sSubnet')]\"", "bar"),
		},
	}

	for _, c := range cases {
		if getVNETAddressPrefixes(c.p) != c.expected {
			t.Fatalf("expected getVNETAddressPrefixes() to return %s but instead got %s", c.expected, getVNETAddressPrefixes(c.p))
		}
	}
}

func TestGetVNETSubnets(t *testing.T) {
	masterString := `{
            "name": "[variables('masterSubnetName')]",
            "properties": {
              "addressPrefix": "[variables('masterSubnet')]"
            }
          }`
	agentString := `          {
            "name": "[variables('%sSubnetName')]",
            "properties": {
              "addressPrefix": "[variables('%sSubnet')]"
            }
          }`
	agentStringNSG := `          {
            "name": "[variables('%sSubnetName')]",
            "properties": {
              "addressPrefix": "[variables('%sSubnet')]",
              "networkSecurityGroup": {
                "id": "[resourceId('Microsoft.Network/networkSecurityGroups', variables('%sNSGName'))]"
              }
            }
          }`
	cases := []struct {
		p        *api.Properties
		addNSG   bool
		expected string
	}{
		{
			p: &api.Properties{
				AgentPoolProfiles: []*api.AgentPoolProfile{},
			},
			addNSG:   false,
			expected: masterString,
		},
		{
			p: &api.Properties{
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						Name: "foo",
					},
					{
						Name: "bar",
					},
				},
			},
			addNSG:   false,
			expected: masterString + ",\n" + fmt.Sprintf(agentString, "foo", "foo") + ",\n" + fmt.Sprintf(agentString, "bar", "bar"),
		},
		{
			p: &api.Properties{
				AgentPoolProfiles: []*api.AgentPoolProfile{
					{
						Name: "foo",
					},
					{
						Name: "bar",
					},
				},
			},
			addNSG:   true,
			expected: masterString + ",\n" + fmt.Sprintf(agentStringNSG, "foo", "foo", "foo") + ",\n" + fmt.Sprintf(agentStringNSG, "bar", "bar", "bar"),
		},
	}

	for _, c := range cases {
		if getVNETSubnets(c.p, c.addNSG) != c.expected {
			t.Fatalf("expected getVNETSubnets() to return %s but instead got %s", c.expected, getVNETSubnets(c.p, c.addNSG))
		}
	}
}

func TestGetDataDisks(t *testing.T) {
	baseString := "\"dataDisks\": [\n"
	dataDisks := `            {
              "createOption": "Empty",
              "diskSizeGB": "%d",
              "lun": %d,
              "caching": "ReadOnly",
              "name": "[concat(variables('%sVMNamePrefix'), copyIndex(),'-datadisk%d')]",
              "vhd": {
                "uri": "[concat('http://',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('%sDataAccountName'),'.blob.core.windows.net/vhds/',variables('%sVMNamePrefix'),copyIndex(), '--datadisk%d.vhd')]"
              }
            }`
	managedDataDisks := `            {
              "diskSizeGB": "%d",
              "lun": %d,
              "caching": "ReadOnly",
              "createOption": "Empty"
            }`
	cases := []struct {
		p        *api.AgentPoolProfile
		expected string
	}{
		{
			p:        &api.AgentPoolProfile{},
			expected: "",
		},
		{
			p: &api.AgentPoolProfile{
				Name: "foo",
				DiskSizesGB: []int{
					128,
				},
				StorageProfile: api.ManagedDisks,
			},
			expected: baseString + fmt.Sprintf(managedDataDisks, 128, 0) + "\n          ],",
		},
		{
			p: &api.AgentPoolProfile{
				Name: "bar",
				DiskSizesGB: []int{
					128,
				},
				StorageProfile: api.StorageAccount,
			},
			expected: baseString + fmt.Sprintf(dataDisks, 128, 0, "bar", 0, "bar", "bar", "bar", "bar", 0) + "\n          ],",
		},
		{
			p: &api.AgentPoolProfile{
				Name: "foo",
				DiskSizesGB: []int{
					128,
					384,
				},
				StorageProfile: api.ManagedDisks,
			},
			expected: baseString + fmt.Sprintf(managedDataDisks, 128, 0) + ",\n" + fmt.Sprintf(managedDataDisks, 384, 1) + "\n          ],",
		},
		{
			p: &api.AgentPoolProfile{
				Name: "bar",
				DiskSizesGB: []int{
					128,
					256,
				},
				StorageProfile: api.StorageAccount,
			},
			expected: baseString + fmt.Sprintf(dataDisks, 128, 0, "bar", 0, "bar", "bar", "bar", "bar", 0) + ",\n" + fmt.Sprintf(dataDisks, 256, 1, "bar", 1, "bar", "bar", "bar", "bar", 1) + "\n          ],",
		},
	}

	for _, c := range cases {
		if getDataDisks(c.p) != c.expected {
			t.Fatalf("expected getDataDisks() to return %s but instead got %s", c.expected, getDataDisks(c.p))
		}
	}
}

func TestGenerateConsecutiveIPsList(t *testing.T) {
	cases := []struct {
		count            int
		firstThreeOctets string
		fourthOctet      int
		expectError      bool
	}{
		{
			count:            3,
			firstThreeOctets: "10.0.0",
			fourthOctet:      240,
		},
		{
			count:            45,
			firstThreeOctets: "192.168.1",
			fourthOctet:      201,
		},
		{
			count:            45,
			firstThreeOctets: "192.168.1.5.6.7", // Too many octets
			fourthOctet:      1,
			expectError:      true,
		},
		{
			count:            10, // This will result in fourth octet overflow
			firstThreeOctets: "4.5.6",
			fourthOctet:      254,
			expectError:      true,
		},
		{
			count:            1, // This is the broadcast address
			firstThreeOctets: "192.168.1",
			fourthOctet:      254,
			expectError:      true,
		},
	}

	for _, c := range cases {
		firstIPAddress := fmt.Sprintf("%s.%d", c.firstThreeOctets, c.fourthOctet)
		ret, err := generateConsecutiveIPsList(c.count, fmt.Sprintf("%s.%d", c.firstThreeOctets, c.fourthOctet))
		if c.expectError {
			if err == nil {
				t.Fatalf("expected error from generateConsecutiveIPsList(%d, %s)!", c.count, firstIPAddress)
			}
		} else {
			if len(ret) != c.count {
				t.Fatalf("expected %d IP addresses from generateConsecutiveIPsList() response, but instead got %d", c.count, len(ret))
			}
			for i, ip := range ret {
				expected := fmt.Sprintf("%s.%d", c.firstThreeOctets, c.fourthOctet+i)
				if ip != expected {
					t.Fatalf("expected %s in generateConsecutiveIPsList() response set at index %d, but instead got %s", expected, i, ip)
				}
			}
		}
	}
}

func TestValidateProfileOptedForExtension(t *testing.T) {
	cases := []struct {
		extensions          []api.Extension
		name                string
		expectedEnabled     bool
		expectedSingleOrAll string
	}{
		{
			extensions: []api.Extension{
				{
					Name:        "foo",
					SingleOrAll: "single",
				},
			},
			name:                "foo",
			expectedEnabled:     true,
			expectedSingleOrAll: "single",
		},
		{
			extensions: []api.Extension{
				{
					Name:        "foo",
					SingleOrAll: "All",
				},
			},
			name:                "foo",
			expectedEnabled:     true,
			expectedSingleOrAll: "All",
		},
		{
			extensions: []api.Extension{
				{
					Name: "foo",
				},
			},
			name:                "foo",
			expectedEnabled:     true,
			expectedSingleOrAll: "",
		},
		{
			extensions: []api.Extension{
				{
					Name: "foo",
				},
			},
			name:                "bar",
			expectedEnabled:     false,
			expectedSingleOrAll: "",
		},
	}

	for _, c := range cases {
		enabled, singleOrAll := validateProfileOptedForExtension(c.name, c.extensions)
		if enabled != c.expectedEnabled {
			t.Fatalf("expected validateProfileOptedForExtension(%s, %v) to return %t but instead got %t", c.name, c.extensions, c.expectedEnabled, enabled)
		}
		if singleOrAll != c.expectedSingleOrAll {
			t.Fatalf("expected validateProfileOptedForExtension(%s, %v) to return %s but instead got %s", c.name, c.extensions, c.expectedSingleOrAll, singleOrAll)
		}
	}
}

func TestGetMasterLinkedTemplateText(t *testing.T) {
	cases := []struct {
		orchestrator     string
		extensionProfile *api.ExtensionProfile
		singleOrAll      string
		expected         string
		expectedErr      error
	}{
		{
			orchestrator: api.Kubernetes,
			extensionProfile: &api.ExtensionProfile{
				Name:    "winrm",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
			singleOrAll: "single",
			expected: `{
	"name": "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), 'winrm')]",
	"type": "Microsoft.Resources/deployments",
	"apiVersion": "[variables('apiVersionDeployments')]",
	"dependsOn": [
		"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/extensions/cse', '-master-', copyIndex(variables('masterOffset')))]"
	],
	"copy": {
		"count": 1,
		"name": "winrmExtensionLoop"
	},
	"properties": {
		"mode": "Incremental",
		"templateLink": {
			"uri": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json",
			"contentVersion": "1.0.0.0"
		},
		"parameters": {
			"artifactsLocation": {
				"value": "https://raw.githubusercontent.com/Azure/aks-engine/master/"
			},
			"apiVersionDeployments": {
				"value": "[variables('apiVersionDeployments')]"
			},
			"targetVMName": {
				"value": "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"
			},
			"targetVMType": {
				"value": "master"
			},
			"extensionParameters": {
				"value": "[parameters('winrmParameters')]"
			},
			"vmIndex":{
				"value": "[copyIndex(variables('masterOffset'))]"
			}
		}
	}
}`,
			expectedErr: nil,
		},
		{
			orchestrator: api.Kubernetes,
			extensionProfile: &api.ExtensionProfile{
				Name:    "winrm",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
			expected: `{
	"name": "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), 'winrm')]",
	"type": "Microsoft.Resources/deployments",
	"apiVersion": "[variables('apiVersionDeployments')]",
	"dependsOn": [
		"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/extensions/cse', '-master-', copyIndex(variables('masterOffset')))]"
	],
	"copy": {
		"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
		"name": "winrmExtensionLoop"
	},
	"properties": {
		"mode": "Incremental",
		"templateLink": {
			"uri": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json",
			"contentVersion": "1.0.0.0"
		},
		"parameters": {
			"artifactsLocation": {
				"value": "https://raw.githubusercontent.com/Azure/aks-engine/master/"
			},
			"apiVersionDeployments": {
				"value": "[variables('apiVersionDeployments')]"
			},
			"targetVMName": {
				"value": "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"
			},
			"targetVMType": {
				"value": "master"
			},
			"extensionParameters": {
				"value": "[parameters('winrmParameters')]"
			},
			"vmIndex":{
				"value": "[copyIndex(variables('masterOffset'))]"
			}
		}
	}
}`,
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		ret, err := getMasterLinkedTemplateText(c.orchestrator, c.extensionProfile, c.singleOrAll)
		ret = strings.Join(strings.Fields(ret), " ")
		expected := strings.Join(strings.Fields(c.expected), " ")
		if ret != expected {
			t.Fatalf("expected getMasterLinkedTemplateText(%s, %v, %s) to return %s but instead got %s", c.orchestrator, c.extensionProfile, c.singleOrAll, expected, ret)
		}
		if err != c.expectedErr {
			t.Fatalf("expected getMasterLinkedTemplateText(%s, %v, %s) to return %s but instead got %s", c.orchestrator, c.extensionProfile, c.singleOrAll, c.expectedErr, err)
		}
	}
}

func TestGetAgentPoolLinkedTemplateText(t *testing.T) {
	cases := []struct {
		orchestrator     string
		agentPoolProfile *api.AgentPoolProfile
		extensionProfile *api.ExtensionProfile
		singleOrAll      string
		expected         string
		expectedErr      error
	}{
		{
			orchestrator: api.Kubernetes,
			agentPoolProfile: &api.AgentPoolProfile{
				Name:                "foo",
				AvailabilityProfile: "AvailabilitySet",
			},
			extensionProfile: &api.ExtensionProfile{
				Name:    "winrm",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
			singleOrAll: "single",
			expected:    `{ "name": "[concat(variables('fooVMNamePrefix'), copyIndex(variables('fooOffset')), 'winrm')]", "type": "Microsoft.Resources/deployments", "apiVersion": "[variables('apiVersionDeployments')]", "dependsOn": [ "[concat('Microsoft.Compute/virtualMachines/', variables('fooVMNamePrefix'), copyIndex(variables('fooOffset')), '/extensions/cse', '-agent-', copyIndex(variables('fooOffset')))]" ], "copy": { "count": 1, "name": "winrmExtensionLoop" }, "properties": { "mode": "Incremental", "templateLink": { "uri": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json", "contentVersion": "1.0.0.0" }, "parameters": { "artifactsLocation": { "value": "https://raw.githubusercontent.com/Azure/aks-engine/master/" }, "apiVersionDeployments": { "value": "[variables('apiVersionDeployments')]" }, "targetVMName": { "value": "[concat(variables('fooVMNamePrefix'), copyIndex(variables('fooOffset')))]" }, "targetVMType": { "value": "agent" }, "extensionParameters": { "value": "[parameters('winrmParameters')]" }, "vmIndex":{ "value": "[copyIndex(variables('fooOffset'))]" } } } }`,
			expectedErr: nil,
		},
		{
			orchestrator: api.Kubernetes,
			agentPoolProfile: &api.AgentPoolProfile{
				Name:                "foo",
				AvailabilityProfile: "VirtualMachineScaleSets",
			},
			extensionProfile: &api.ExtensionProfile{
				Name:    "winrm",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
			singleOrAll: "single",
			expected:    `{ "name": "[concat(variables('fooVMNamePrefix'), copyIndex(), 'winrm')]", "type": "Microsoft.Resources/deployments", "apiVersion": "[variables('apiVersionDeployments')]", "dependsOn": [ "[concat('Microsoft.Compute/virtualMachines/', variables('fooVMNamePrefix'), copyIndex(), '/extensions/cse', '-agent-', copyIndex())]" ], "copy": { "count": 1, "name": "winrmExtensionLoop" }, "properties": { "mode": "Incremental", "templateLink": { "uri": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json", "contentVersion": "1.0.0.0" }, "parameters": { "artifactsLocation": { "value": "https://raw.githubusercontent.com/Azure/aks-engine/master/" }, "apiVersionDeployments": { "value": "[variables('apiVersionDeployments')]" }, "targetVMName": { "value": "[concat(variables('fooVMNamePrefix'), copyIndex())]" }, "targetVMType": { "value": "agent" }, "extensionParameters": { "value": "[parameters('winrmParameters')]" }, "vmIndex":{ "value": "[copyIndex()]" } } } }`,
			expectedErr: nil,
		},
		{
			orchestrator: api.Kubernetes,
			agentPoolProfile: &api.AgentPoolProfile{
				Name:                "foo",
				AvailabilityProfile: "VirtualMachineScaleSets",
			},
			extensionProfile: &api.ExtensionProfile{
				Name:    "winrm",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
			singleOrAll: "All",
			expected:    `{ "name": "[concat(variables('fooVMNamePrefix'), copyIndex(), 'winrm')]", "type": "Microsoft.Resources/deployments", "apiVersion": "[variables('apiVersionDeployments')]", "dependsOn": [ "[concat('Microsoft.Compute/virtualMachines/', variables('fooVMNamePrefix'), copyIndex(), '/extensions/cse', '-agent-', copyIndex())]" ], "copy": { "count": "[variables('fooCount'))]", "name": "winrmExtensionLoop" }, "properties": { "mode": "Incremental", "templateLink": { "uri": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json", "contentVersion": "1.0.0.0" }, "parameters": { "artifactsLocation": { "value": "https://raw.githubusercontent.com/Azure/aks-engine/master/" }, "apiVersionDeployments": { "value": "[variables('apiVersionDeployments')]" }, "targetVMName": { "value": "[concat(variables('fooVMNamePrefix'), copyIndex())]" }, "targetVMType": { "value": "agent" }, "extensionParameters": { "value": "[parameters('winrmParameters')]" }, "vmIndex":{ "value": "[copyIndex()]" } } } }`,
			expectedErr: nil,
		},
		{
			orchestrator: api.Kubernetes,
			agentPoolProfile: &api.AgentPoolProfile{
				Name:                "foo",
				AvailabilityProfile: "AvailabilitySet",
			},
			extensionProfile: &api.ExtensionProfile{
				Name:    "winrm",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
			singleOrAll: "All",
			expected:    `{ "name": "[concat(variables('fooVMNamePrefix'), copyIndex(variables('fooOffset')), 'winrm')]", "type": "Microsoft.Resources/deployments", "apiVersion": "[variables('apiVersionDeployments')]", "dependsOn": [ "[concat('Microsoft.Compute/virtualMachines/', variables('fooVMNamePrefix'), copyIndex(variables('fooOffset')), '/extensions/cse', '-agent-', copyIndex(variables('fooOffset')))]" ], "copy": { "count": "[sub(variables('fooCount'), variables('fooOffset'))]", "name": "winrmExtensionLoop" }, "properties": { "mode": "Incremental", "templateLink": { "uri": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json", "contentVersion": "1.0.0.0" }, "parameters": { "artifactsLocation": { "value": "https://raw.githubusercontent.com/Azure/aks-engine/master/" }, "apiVersionDeployments": { "value": "[variables('apiVersionDeployments')]" }, "targetVMName": { "value": "[concat(variables('fooVMNamePrefix'), copyIndex(variables('fooOffset')))]" }, "targetVMType": { "value": "agent" }, "extensionParameters": { "value": "[parameters('winrmParameters')]" }, "vmIndex":{ "value": "[copyIndex(variables('fooOffset'))]" } } } }`,
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		ret, err := getAgentPoolLinkedTemplateText(c.agentPoolProfile, c.orchestrator, c.extensionProfile, c.singleOrAll)
		ret = strings.Join(strings.Fields(ret), " ")
		expected := strings.Join(strings.Fields(c.expected), " ")
		if ret != expected {
			t.Fatalf("expected getAgentPoolLinkedTemplateText(%v, %s, %v, %s) to return %s but instead got %s", c.agentPoolProfile, c.orchestrator, c.extensionProfile, c.singleOrAll, expected, ret)
		}
		if err != c.expectedErr {
			t.Fatalf("expected getAgentPoolLinkedTemplateText(%v, %s, %v, %s) to return %s but instead got %s", c.agentPoolProfile, c.orchestrator, c.extensionProfile, c.singleOrAll, c.expectedErr, err)
		}
	}
}

func TestGetSSHPublicKeysPowerShell(t *testing.T) {
	cases := []struct {
		publicKeys []api.PublicKey
		expected   string
	}{
		{
			publicKeys: []api.PublicKey{
				{
					KeyData: "foo   ",
				},
			},
			expected: "\"foo\"",
		},
		{
			publicKeys: []api.PublicKey{
				{
					KeyData: "  foo",
				},
				{
					KeyData: " bar   ",
				},
			},
			expected: "\"foo\", \"bar\"",
		},
	}

	for _, c := range cases {
		linuxProfile := &api.LinuxProfile{}
		linuxProfile.SSH.PublicKeys = c.publicKeys
		ret := getSSHPublicKeysPowerShell(linuxProfile)
		if ret != c.expected {
			t.Fatalf("expected getSSHPublicKeysPowerShell(%v) to return %s but instead got %s", linuxProfile, c.expected, ret)
		}
	}
}

func TestGetWindowsMasterSubnetARMParam(t *testing.T) {
	cases := []struct {
		m        *api.MasterProfile
		expected string
	}{
		{
			m:        &api.MasterProfile{},
			expected: "',parameters('masterSubnet'),'",
		},
		{
			m:        nil,
			expected: "',parameters('masterSubnet'),'",
		},
		{
			m: &api.MasterProfile{
				VnetSubnetID: "/my/subnet",
			},
			expected: "',parameters('vnetCidr'),'",
		},
	}

	for _, c := range cases {
		ret := getWindowsMasterSubnetARMParam(c.m)
		if ret != c.expected {
			t.Fatalf("expected getWindowsMasterSubnetARMParam(%v) to return %s but instead got %s", c.m, c.expected, ret)
		}
	}
}

func TestVerifyGetBase64EncodedGzippedCustomScriptIsTransparent(t *testing.T) {
	cases := []struct {
		name string
		cs   *api.ContainerService
	}{
		{
			name: "zero value cs",
			cs:   &api.ContainerService{},
		},
		{
			name: "cs with stuff in it",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
							UseManagedIdentity: true,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			for _, file := range []string{kubernetesMountEtcd,
				etcdSystemdService,
				dhcpv6ConfigurationScript,
				kubernetesCISScript,
				kmsSystemdService,
				labelNodesScript,
				labelNodesSystemdService,
				aptPreferences,
				kubernetesHealthMonitorScript,
				kubernetesKubeletMonitorSystemdService,
				kubernetesDockerMonitorSystemdService,
				kubernetesDockerMonitorSystemdTimer,
				kubernetesDockerMonitorSystemdTimer,
				dockerClearMountPropagationFlags,
				auditdRules,
				systemdBPFMount,
			} {
				ret := getBase64EncodedGzippedCustomScript(file, c.cs)
				b, err := Asset(file)
				if err != nil {
					t.Fatalf("unable to load file")
				}
				if getBase64EncodedGzippedCustomScriptFromStr(string(b)) != ret {
					t.Fatalf("getBase64EncodedGzippedCustomScript returned an unexpected result for file %s, perhaps it was interpreted by golang templates unexpectedly", file)
				}
			}
		})
	}
}

func TestWrapAsVariableObject(t *testing.T) {
	tests := []struct {
		name     string
		o        string
		s        string
		expected string
	}{
		{
			name:     "just a string",
			o:        "cloudInitFiles",
			s:        "foo",
			expected: "',variables('cloudInitFiles').foo,'",
		},
		{
			name:     "empty string",
			o:        "cloudInitFiles",
			s:        "",
			expected: "',variables('cloudInitFiles').,'",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := wrapAsVariableObject(test.o, test.s)
			if test.expected != ret {
				t.Errorf("expected %s, instead got : %s", test.expected, ret)
			}
		})
	}
}

func TestGetClusterAutoscalerAddonFuncMap(t *testing.T) {
	specConfig := api.AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	k8sComponentsByVersionMap := api.GetK8sComponentsByVersionMap(&api.KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR})
	cases := []struct {
		name                                              string
		addon                                             api.KubernetesAddon
		cs                                                *api.ContainerService
		expectedImage                                     string
		expectedCPUReqs                                   string
		expectedCPULimits                                 string
		expectedMemReqs                                   string
		expectedMemLimits                                 string
		expectedScanInterval                              string
		expectedVersion                                   string
		expectedMode                                      string
		expectedNodesConfig                               string
		expectedVMType                                    string
		expectedVolumeMounts                              string
		expectedVolumes                                   string
		expectedHostNetwork                               string
		expectedCloud                                     string
		expectedUseManagedIdentity                        string
		expectedIsKubernetesVersionGeOneDotSixteenDotZero bool
	}{
		{
			name: "single pool",
			addon: api.KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.16.9"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []api.AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.16.9",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
							UseManagedIdentity: true,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedImage:              specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.16.9"][common.ClusterAutoscalerAddonName],
			expectedCPUReqs:            "100m",
			expectedCPULimits:          "100m",
			expectedMemReqs:            "300Mi",
			expectedMemLimits:          "300Mi",
			expectedScanInterval:       "1m",
			expectedVersion:            "3",
			expectedMode:               api.AddonModeEnsureExists,
			expectedNodesConfig:        "        - --nodes=1:10:k8s-pool1-49584119-vmss",
			expectedVMType:             "dm1zcw==", // base 64 encoding of vmss
			expectedVolumeMounts:       "\n        - mountPath: /var/lib/waagent/\n          name: waagent\n          readOnly: true",
			expectedVolumes:            "\n      - hostPath:\n          path: /var/lib/waagent/\n        name: waagent",
			expectedHostNetwork:        "\n      hostNetwork: true",
			expectedCloud:              "AzurePublicCloud",
			expectedUseManagedIdentity: "true",
			expectedIsKubernetesVersionGeOneDotSixteenDotZero: true,
		},
		{
			name: "multiple pools",
			addon: api.KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []api.AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
					{
						Name: "pool2",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
							UseManagedIdentity: true,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedImage:              specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
			expectedCPUReqs:            "100m",
			expectedCPULimits:          "100m",
			expectedMemReqs:            "300Mi",
			expectedMemLimits:          "300Mi",
			expectedScanInterval:       "1m",
			expectedVersion:            "3",
			expectedMode:               api.AddonModeEnsureExists,
			expectedNodesConfig:        "        - --nodes=1:10:k8s-pool1-49584119-vmss\n        - --nodes=1:10:k8s-pool2-49584119-vmss",
			expectedVMType:             "dm1zcw==", // base 64 encoding of vmss
			expectedVolumeMounts:       "\n        - mountPath: /var/lib/waagent/\n          name: waagent\n          readOnly: true",
			expectedVolumes:            "\n      - hostPath:\n          path: /var/lib/waagent/\n        name: waagent",
			expectedHostNetwork:        "\n      hostNetwork: true",
			expectedCloud:              "AzurePublicCloud",
			expectedUseManagedIdentity: "true",
		},
		{
			name: "no pools",
			addon: api.KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
							UseManagedIdentity: true,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedImage:              specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
			expectedCPUReqs:            "100m",
			expectedCPULimits:          "100m",
			expectedMemReqs:            "300Mi",
			expectedMemLimits:          "300Mi",
			expectedScanInterval:       "1m",
			expectedVersion:            "3",
			expectedMode:               api.AddonModeEnsureExists,
			expectedNodesConfig:        "",
			expectedVMType:             "dm1zcw==", // base 64 encoding of vmss
			expectedVolumeMounts:       "\n        - mountPath: /var/lib/waagent/\n          name: waagent\n          readOnly: true",
			expectedVolumes:            "\n      - hostPath:\n          path: /var/lib/waagent/\n        name: waagent",
			expectedHostNetwork:        "\n      hostNetwork: true",
			expectedCloud:              "AzurePublicCloud",
			expectedUseManagedIdentity: "true",
		},
		{
			name: "non-MSI scenario",
			addon: api.KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []api.AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
					{
						Name: "pool2",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "3",
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedImage:              specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
			expectedCPUReqs:            "100m",
			expectedCPULimits:          "100m",
			expectedMemReqs:            "300Mi",
			expectedMemLimits:          "300Mi",
			expectedScanInterval:       "1m",
			expectedVersion:            "3",
			expectedMode:               api.AddonModeEnsureExists,
			expectedNodesConfig:        "        - --nodes=1:10:k8s-pool1-49584119-vmss\n        - --nodes=1:3:k8s-pool2-49584119-vmss",
			expectedVMType:             "dm1zcw==", // base 64 encoding of vmss
			expectedVolumeMounts:       "",
			expectedVolumes:            "",
			expectedHostNetwork:        "",
			expectedCloud:              "AzurePublicCloud",
			expectedUseManagedIdentity: "false",
		},
		{
			name: "china scenario",
			addon: api.KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []api.AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
					{
						Name: "pool2",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "3",
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
				Location: "chinanorth",
			},
			expectedImage:              specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
			expectedCPUReqs:            "100m",
			expectedCPULimits:          "100m",
			expectedMemReqs:            "300Mi",
			expectedMemLimits:          "300Mi",
			expectedScanInterval:       "1m",
			expectedVersion:            "3",
			expectedMode:               api.AddonModeEnsureExists,
			expectedNodesConfig:        "        - --nodes=1:10:k8s-pool1-49584119-vmss\n        - --nodes=1:3:k8s-pool2-49584119-vmss",
			expectedVMType:             "dm1zcw==", // base 64 encoding of vmss
			expectedVolumeMounts:       "",
			expectedVolumes:            "",
			expectedHostNetwork:        "",
			expectedCloud:              "AzureChinaCloud",
			expectedUseManagedIdentity: "false",
		},
		{
			name: "german cloud scenario",
			addon: api.KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []api.AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
					{
						Name: "pool2",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "3",
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
				Location: "germanynortheast",
			},
			expectedImage:              specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
			expectedCPUReqs:            "100m",
			expectedCPULimits:          "100m",
			expectedMemReqs:            "300Mi",
			expectedMemLimits:          "300Mi",
			expectedScanInterval:       "1m",
			expectedVersion:            "3",
			expectedMode:               api.AddonModeEnsureExists,
			expectedNodesConfig:        "        - --nodes=1:10:k8s-pool1-49584119-vmss\n        - --nodes=1:3:k8s-pool2-49584119-vmss",
			expectedVMType:             "dm1zcw==", // base 64 encoding of vmss
			expectedVolumeMounts:       "",
			expectedVolumes:            "",
			expectedHostNetwork:        "",
			expectedCloud:              "AzureGermanCloud",
			expectedUseManagedIdentity: "false",
		},
		{
			name: "usgov cloud scenario",
			addon: api.KubernetesAddon{
				Name:    common.ClusterAutoscalerAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"scan-interval": "1m",
					"v":             "3",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.ClusterAutoscalerAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
					},
				},
				Pools: []api.AddonNodePoolsConfig{
					{
						Name: "pool1",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "10",
						},
					},
					{
						Name: "pool2",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "3",
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.ClusterAutoscalerAddonName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
				Location: "usgovnorth",
			},
			expectedImage:              specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.ClusterAutoscalerAddonName],
			expectedCPUReqs:            "100m",
			expectedCPULimits:          "100m",
			expectedMemReqs:            "300Mi",
			expectedMemLimits:          "300Mi",
			expectedScanInterval:       "1m",
			expectedVersion:            "3",
			expectedMode:               api.AddonModeEnsureExists,
			expectedNodesConfig:        "        - --nodes=1:10:k8s-pool1-49584119-vmss\n        - --nodes=1:3:k8s-pool2-49584119-vmss",
			expectedVMType:             "dm1zcw==", // base 64 encoding of vmss
			expectedVolumeMounts:       "",
			expectedVolumes:            "",
			expectedHostNetwork:        "",
			expectedCloud:              "AzureUSGovernmentCloud",
			expectedUseManagedIdentity: "false",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			funcMap := getClusterAutoscalerAddonFuncMap(c.addon, c.cs)
			v := reflect.ValueOf(funcMap["ContainerImage"])
			ret := v.Call([]reflect.Value{reflect.ValueOf(common.ClusterAutoscalerAddonName)})
			if ret[0].Interface() != c.expectedImage {
				t.Errorf("expected funcMap invocation of ContainerImage %s to return %s, instead got %s", common.ClusterAutoscalerAddonName, c.expectedImage, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerCPUReqs"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.ClusterAutoscalerAddonName)})
			if ret[0].Interface() != c.expectedCPUReqs {
				t.Errorf("expected funcMap invocation of ContainerCPUReqs %s to return %s, instead got %s", common.ClusterAutoscalerAddonName, c.expectedCPUReqs, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerCPULimits"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.ClusterAutoscalerAddonName)})
			if ret[0].Interface() != c.expectedCPULimits {
				t.Errorf("expected funcMap invocation of ContainerCPULimits %s to return %s, instead got %s", common.ClusterAutoscalerAddonName, c.expectedCPULimits, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerMemReqs"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.ClusterAutoscalerAddonName)})
			if ret[0].Interface() != c.expectedMemReqs {
				t.Errorf("expected funcMap invocation of ContainerMemReqs %s to return %s, instead got %s", common.ClusterAutoscalerAddonName, c.expectedMemReqs, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerMemLimits"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.ClusterAutoscalerAddonName)})
			if ret[0].Interface() != c.expectedMemLimits {
				t.Errorf("expected funcMap invocation of ContainerMemLimits %s to return %s, instead got %s", common.ClusterAutoscalerAddonName, c.expectedMemLimits, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerConfig"])
			ret = v.Call([]reflect.Value{reflect.ValueOf("scan-interval")})
			if ret[0].Interface() != c.expectedScanInterval {
				t.Errorf("expected funcMap invocation of ContainerConfig %s to return %s, instead got %s", "scan-interval", c.expectedScanInterval, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerConfig"])
			ret = v.Call([]reflect.Value{reflect.ValueOf("v")})
			if ret[0].Interface() != c.expectedVersion {
				t.Errorf("expected funcMap invocation of ContainerConfig %s to return %s, instead got %s", "v", c.expectedVersion, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerConfig"])
			ret = v.Call([]reflect.Value{reflect.ValueOf("non-existent")})
			if ret[0].Interface() != "" {
				t.Errorf("expected funcMap invocation of ContainerConfig %s to return \"\", instead got %s", "non-existent", ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetMode"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedMode {
				t.Errorf("expected funcMap invocation of GetMode to return %s, instead got %s", c.expectedMode, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetClusterAutoscalerNodesConfig"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedNodesConfig {
				t.Errorf("expected funcMap invocation of GetClusterAutoscalerNodesConfig to return %s, instead got %s", c.expectedNodesConfig, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetBase64EncodedVMType"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedVMType {
				t.Errorf("expected funcMap invocation of GetBase64EncodedVMType to return %s, instead got %s", c.expectedVMType, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetVolumeMounts"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedVolumeMounts {
				t.Errorf("expected funcMap invocation of GetVolumeMounts to return %s, instead got %s", c.expectedVolumeMounts, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetVolumes"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedVolumes {
				t.Errorf("expected funcMap invocation of GetVolumes to return %s, instead got %s", c.expectedVolumes, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetHostNetwork"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHostNetwork {
				t.Errorf("expected funcMap invocation of GetHostNetwork to return %s, instead got %s", c.expectedHostNetwork, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetCloud"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedCloud {
				t.Errorf("expected funcMap invocation of GetCloud to return %s, instead got %s", c.expectedCloud, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["UseManagedIdentity"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedUseManagedIdentity {
				t.Errorf("expected funcMap invocation of UseManagedIdentity to return %s, instead got %s", c.expectedUseManagedIdentity, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["IsKubernetesVersionGe"])
			ret = v.Call([]reflect.Value{reflect.ValueOf("1.16.0")})
			if ret[0].Interface() != c.expectedIsKubernetesVersionGeOneDotSixteenDotZero {
				t.Errorf("expected funcMap invocation of IsKubernetesVersionGe for 1.16.0 to return %t, instead got %t", c.expectedIsKubernetesVersionGeOneDotSixteenDotZero, ret[0].Interface())
			}
		})
	}
}

func TestGetAddonFuncMap(t *testing.T) {
	specConfig := api.AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	k8sComponentsByVersionMap := api.GetK8sComponentsByVersionMap(&api.KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR})
	cases := []struct {
		name                                              string
		addon                                             api.KubernetesAddon
		cs                                                *api.ContainerService
		expectedImage                                     string
		expectedCPUReqs                                   string
		expectedCPULimits                                 string
		expectedMemReqs                                   string
		expectedMemLimits                                 string
		expectedFoo                                       string
		expectedIsAzureStackCloud                         bool
		expectedNeedsStorageAccountStorageClasses         bool
		expectedNeedsManagedDiskStorageClasses            bool
		expectedUsesCloudControllerManager                bool
		expectedHasAvailabilityZones                      bool
		expectedGetZones                                  string
		expectedHasWindows                                bool
		expectedHasLinux                                  bool
		expectedCSIControllerReplicas                     string
		expectedShouldEnableAzureDiskCSISnapshotFeature   bool
		expectedShouldEnableAzureFileCSISnapshotFeature   bool
		expectedIsKubernetesVersionGeOneDotSixteenDotZero bool
		expectedMode                                      string
	}{
		{
			name: "coredns as an example",
			addon: api.KubernetesAddon{
				Name:    common.CoreDNSAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"foo": "bar",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.CoreDNSAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"foo": "bar",
									},
									Containers: []api.KubernetesContainerSpec{
										{
											Name:           common.CoreDNSAddonName,
											CPURequests:    "100m",
											MemoryRequests: "300Mi",
											CPULimits:      "100m",
											MemoryLimits:   "300Mi",
											Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							StorageProfile:      api.ManagedDisks,
							OSType:              api.Linux,
						},
					},
				},
			},
			expectedImage:             specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
			expectedCPUReqs:           "100m",
			expectedCPULimits:         "100m",
			expectedMemReqs:           "300Mi",
			expectedMemLimits:         "300Mi",
			expectedFoo:               "bar",
			expectedIsAzureStackCloud: false,
			expectedNeedsStorageAccountStorageClasses:         false,
			expectedNeedsManagedDiskStorageClasses:            true,
			expectedUsesCloudControllerManager:                false,
			expectedHasAvailabilityZones:                      false,
			expectedGetZones:                                  "",
			expectedHasWindows:                                false,
			expectedHasLinux:                                  true,
			expectedCSIControllerReplicas:                     "2",
			expectedShouldEnableAzureDiskCSISnapshotFeature:   false,
			expectedShouldEnableAzureFileCSISnapshotFeature:   true,
			expectedIsKubernetesVersionGeOneDotSixteenDotZero: false,
			expectedMode: api.AddonModeEnsureExists,
		},
		{
			name: "coredns as an example - Azure Stack",
			addon: api.KubernetesAddon{
				Name:    common.CoreDNSAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"foo": "bar",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.CoreDNSAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"foo": "bar",
									},
									Containers: []api.KubernetesContainerSpec{
										{
											Name:           common.CoreDNSAddonName,
											CPURequests:    "100m",
											MemoryRequests: "300Mi",
											CPULimits:      "100m",
											MemoryLimits:   "300Mi",
											Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							StorageProfile:      api.ManagedDisks,
							OSType:              api.Linux,
						},
					},
					CustomCloudProfile: &api.CustomCloudProfile{
						IdentitySystem: "adfs",
						PortalURL:      "https://portal.testlocation.contoso.com/",
					},
				},
			},
			expectedImage:             specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
			expectedCPUReqs:           "100m",
			expectedCPULimits:         "100m",
			expectedMemReqs:           "300Mi",
			expectedMemLimits:         "300Mi",
			expectedFoo:               "bar",
			expectedIsAzureStackCloud: true,
			expectedNeedsStorageAccountStorageClasses:         false,
			expectedNeedsManagedDiskStorageClasses:            true,
			expectedUsesCloudControllerManager:                false,
			expectedHasAvailabilityZones:                      false,
			expectedGetZones:                                  "",
			expectedHasWindows:                                false,
			expectedHasLinux:                                  true,
			expectedCSIControllerReplicas:                     "2",
			expectedShouldEnableAzureDiskCSISnapshotFeature:   false,
			expectedShouldEnableAzureFileCSISnapshotFeature:   true,
			expectedIsKubernetesVersionGeOneDotSixteenDotZero: false,
			expectedMode: api.AddonModeEnsureExists,
		},
		{
			name: "coredns as an example - StorageAccount",
			addon: api.KubernetesAddon{
				Name:    common.CoreDNSAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"foo": "bar",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.CoreDNSAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							NetworkPlugin: api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"foo": "bar",
									},
									Containers: []api.KubernetesContainerSpec{
										{
											Name:           common.CoreDNSAddonName,
											CPURequests:    "100m",
											MemoryRequests: "300Mi",
											CPULimits:      "100m",
											MemoryLimits:   "300Mi",
											Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							StorageProfile:      api.StorageAccount,
							OSType:              api.Linux,
						},
					},
				},
			},
			expectedImage:             specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
			expectedCPUReqs:           "100m",
			expectedCPULimits:         "100m",
			expectedMemReqs:           "300Mi",
			expectedMemLimits:         "300Mi",
			expectedFoo:               "bar",
			expectedIsAzureStackCloud: false,
			expectedNeedsStorageAccountStorageClasses:         true,
			expectedNeedsManagedDiskStorageClasses:            false,
			expectedUsesCloudControllerManager:                false,
			expectedHasAvailabilityZones:                      false,
			expectedGetZones:                                  "",
			expectedHasWindows:                                false,
			expectedHasLinux:                                  true,
			expectedCSIControllerReplicas:                     "2",
			expectedShouldEnableAzureDiskCSISnapshotFeature:   false,
			expectedShouldEnableAzureFileCSISnapshotFeature:   true,
			expectedIsKubernetesVersionGeOneDotSixteenDotZero: false,
			expectedMode: api.AddonModeEnsureExists,
		},
		{
			name: "coredns as an example - CCM",
			addon: api.KubernetesAddon{
				Name:    common.CoreDNSAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"foo": "bar",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.CoreDNSAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							UseCloudControllerManager: to.BoolPtr(true),
							NetworkPlugin:             api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"foo": "bar",
									},
									Containers: []api.KubernetesContainerSpec{
										{
											Name:           common.CoreDNSAddonName,
											CPURequests:    "100m",
											MemoryRequests: "300Mi",
											CPULimits:      "100m",
											MemoryLimits:   "300Mi",
											Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							StorageProfile:      api.ManagedDisks,
							OSType:              api.Linux,
						},
					},
				},
			},
			expectedImage:             specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
			expectedCPUReqs:           "100m",
			expectedCPULimits:         "100m",
			expectedMemReqs:           "300Mi",
			expectedMemLimits:         "300Mi",
			expectedFoo:               "bar",
			expectedIsAzureStackCloud: false,
			expectedNeedsStorageAccountStorageClasses:         false,
			expectedNeedsManagedDiskStorageClasses:            true,
			expectedUsesCloudControllerManager:                true,
			expectedHasAvailabilityZones:                      false,
			expectedGetZones:                                  "",
			expectedHasWindows:                                false,
			expectedHasLinux:                                  true,
			expectedCSIControllerReplicas:                     "2",
			expectedShouldEnableAzureDiskCSISnapshotFeature:   false,
			expectedShouldEnableAzureFileCSISnapshotFeature:   true,
			expectedIsKubernetesVersionGeOneDotSixteenDotZero: false,
			expectedMode: api.AddonModeEnsureExists,
		},
		{
			name: "coredns as an example - Availability Zones",
			addon: api.KubernetesAddon{
				Name:    common.CoreDNSAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeEnsureExists,
				Config: map[string]string{
					"foo": "bar",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.CoreDNSAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
					},
				},
			},
			cs: &api.ContainerService{
				Location: "eastus2",
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.17.0",
						KubernetesConfig: &api.KubernetesConfig{
							UseCloudControllerManager: to.BoolPtr(true),
							NetworkPlugin:             api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"foo": "bar",
									},
									Containers: []api.KubernetesContainerSpec{
										{
											Name:           common.CoreDNSAddonName,
											CPURequests:    "100m",
											MemoryRequests: "300Mi",
											CPULimits:      "100m",
											MemoryLimits:   "300Mi",
											Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							StorageProfile:      api.ManagedDisks,
							AvailabilityZones: []string{
								"1",
								"2",
							},
							OSType: api.Linux,
						},
					},
				},
			},
			expectedImage:             specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
			expectedCPUReqs:           "100m",
			expectedCPULimits:         "100m",
			expectedMemReqs:           "300Mi",
			expectedMemLimits:         "300Mi",
			expectedFoo:               "bar",
			expectedIsAzureStackCloud: false,
			expectedNeedsStorageAccountStorageClasses:         false,
			expectedNeedsManagedDiskStorageClasses:            true,
			expectedUsesCloudControllerManager:                true,
			expectedHasAvailabilityZones:                      true,
			expectedGetZones:                                  "\n    - eastus2-1\n    - eastus2-2",
			expectedHasWindows:                                false,
			expectedHasLinux:                                  true,
			expectedCSIControllerReplicas:                     "2",
			expectedShouldEnableAzureDiskCSISnapshotFeature:   true,
			expectedShouldEnableAzureFileCSISnapshotFeature:   false,
			expectedIsKubernetesVersionGeOneDotSixteenDotZero: true,
			expectedMode: api.AddonModeEnsureExists,
		},
		{
			name: "coredns as an example - hybrid cluster",
			addon: api.KubernetesAddon{
				Name:    common.CoreDNSAddonName,
				Enabled: to.BoolPtr(true),
				Mode:    api.AddonModeReconcile,
				Config: map[string]string{
					"foo": "bar",
				},
				Containers: []api.KubernetesContainerSpec{
					{
						Name:           common.CoreDNSAddonName,
						CPURequests:    "100m",
						MemoryRequests: "300Mi",
						CPULimits:      "100m",
						MemoryLimits:   "300Mi",
						Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
					},
				},
			},
			cs: &api.ContainerService{
				Location: "eastus2",
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.18.0",
						KubernetesConfig: &api.KubernetesConfig{
							UseCloudControllerManager: to.BoolPtr(true),
							NetworkPlugin:             api.NetworkPluginAzure,
							Addons: []api.KubernetesAddon{
								{
									Name:    common.CoreDNSAddonName,
									Enabled: to.BoolPtr(true),
									Config: map[string]string{
										"foo": "bar",
									},
									Containers: []api.KubernetesContainerSpec{
										{
											Name:           common.CoreDNSAddonName,
											CPURequests:    "100m",
											MemoryRequests: "300Mi",
											CPULimits:      "100m",
											MemoryLimits:   "300Mi",
											Image:          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
										},
									},
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							StorageProfile:      api.ManagedDisks,
							OSType:              api.Windows,
						},
						{
							Name:                "pool2",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							StorageProfile:      api.ManagedDisks,
							OSType:              api.Linux,
						},
					},
				},
			},
			expectedImage:                          specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.4"][common.CoreDNSAddonName],
			expectedCPUReqs:                        "100m",
			expectedCPULimits:                      "100m",
			expectedMemReqs:                        "300Mi",
			expectedMemLimits:                      "300Mi",
			expectedFoo:                            "bar",
			expectedNeedsManagedDiskStorageClasses: true,
			expectedUsesCloudControllerManager:     true,
			expectedHasWindows:                     true,
			expectedHasLinux:                       true,
			expectedCSIControllerReplicas:          "2",
			expectedShouldEnableAzureDiskCSISnapshotFeature:   true,
			expectedShouldEnableAzureFileCSISnapshotFeature:   false,
			expectedIsKubernetesVersionGeOneDotSixteenDotZero: true,
			expectedMode: api.AddonModeReconcile,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			funcMap := getAddonFuncMap(c.addon, c.cs)
			v := reflect.ValueOf(funcMap["ContainerImage"])
			ret := v.Call([]reflect.Value{reflect.ValueOf(common.CoreDNSAddonName)})
			if ret[0].Interface() != c.expectedImage {
				t.Errorf("expected funcMap invocation of ContainerImage %s to return %s, instead got %s", common.CoreDNSAddonName, c.expectedImage, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerCPUReqs"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.CoreDNSAddonName)})
			if ret[0].Interface() != c.expectedCPUReqs {
				t.Errorf("expected funcMap invocation of ContainerCPUReqs %s to return %s, instead got %s", common.CoreDNSAddonName, c.expectedCPUReqs, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerCPULimits"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.CoreDNSAddonName)})
			if ret[0].Interface() != c.expectedCPULimits {
				t.Errorf("expected funcMap invocation of ContainerCPULimits %s to return %s, instead got %s", common.CoreDNSAddonName, c.expectedCPULimits, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerMemReqs"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.CoreDNSAddonName)})
			if ret[0].Interface() != c.expectedMemReqs {
				t.Errorf("expected funcMap invocation of ContainerMemReqs %s to return %s, instead got %s", common.CoreDNSAddonName, c.expectedMemReqs, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerMemLimits"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.CoreDNSAddonName)})
			if ret[0].Interface() != c.expectedMemLimits {
				t.Errorf("expected funcMap invocation of ContainerMemLimits %s to return %s, instead got %s", common.CoreDNSAddonName, c.expectedMemLimits, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ContainerConfig"])
			ret = v.Call([]reflect.Value{reflect.ValueOf("foo")})
			if ret[0].Interface() != c.expectedFoo {
				t.Errorf("expected funcMap invocation of ContainerConfig %s to return %s, instead got %s", "foo", c.expectedFoo, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["IsAzureStackCloud"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedIsAzureStackCloud {
				t.Errorf("expected funcMap invocation of IsAzureStackCloud to return %t, instead got %t", c.expectedIsAzureStackCloud, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["NeedsStorageAccountStorageClasses"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedNeedsStorageAccountStorageClasses {
				t.Errorf("expected funcMap invocation of NeedsStorageAccountStorageClasses to return %t, instead got %t", c.expectedNeedsStorageAccountStorageClasses, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["NeedsManagedDiskStorageClasses"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedNeedsManagedDiskStorageClasses {
				t.Errorf("expected funcMap invocation of NeedsManagedDiskStorageClasses to return %t, instead got %t", c.expectedNeedsManagedDiskStorageClasses, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["UsesCloudControllerManager"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedUsesCloudControllerManager {
				t.Errorf("expected funcMap invocation of UsesCloudControllerManager to return %t, instead got %t", c.expectedUsesCloudControllerManager, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["HasAvailabilityZones"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasAvailabilityZones {
				t.Errorf("expected funcMap invocation of HasAvailabilityZones to return %t, instead got %t", c.expectedHasAvailabilityZones, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetZones"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetZones {
				t.Errorf("expected funcMap invocation of GetZones to return %s, instead got %s", c.expectedGetZones, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["HasWindows"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasWindows {
				t.Errorf("expected funcMap invocation of HasWindows to return %t, instead got %t", c.expectedHasWindows, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["HasLinux"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasLinux {
				t.Errorf("expected funcMap invocation of HasLinux to return %t, instead got %t", c.expectedHasLinux, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["CSIControllerReplicas"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedCSIControllerReplicas {
				t.Errorf("expected funcMap invocation of CSIControllerReplicas to return %s, instead got %s", c.expectedCSIControllerReplicas, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ShouldEnableCSISnapshotFeature"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.AzureDiskCSIDriverAddonName)})
			if ret[0].Interface() != c.expectedShouldEnableAzureDiskCSISnapshotFeature {
				t.Errorf("expected funcMap invocation of ShouldEnableCSISnapshotFeature for %s to return %t, instead got %t", common.AzureDiskCSIDriverAddonName, c.expectedShouldEnableAzureDiskCSISnapshotFeature, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["ShouldEnableCSISnapshotFeature"])
			ret = v.Call([]reflect.Value{reflect.ValueOf(common.AzureFileCSIDriverAddonName)})
			if ret[0].Interface() != c.expectedShouldEnableAzureFileCSISnapshotFeature {
				t.Errorf("expected funcMap invocation of ShouldEnableCSISnapshotFeature for %s to return %t, instead got %t", common.AzureFileCSIDriverAddonName, c.expectedShouldEnableAzureFileCSISnapshotFeature, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["IsKubernetesVersionGe"])
			ret = v.Call([]reflect.Value{reflect.ValueOf("1.16.0")})
			if ret[0].Interface() != c.expectedIsKubernetesVersionGeOneDotSixteenDotZero {
				t.Errorf("expected funcMap invocation of IsKubernetesVersionGe for 1.16.0 to return %t, instead got %t", c.expectedIsKubernetesVersionGeOneDotSixteenDotZero, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetAADPodIdentityTaintKey"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != common.AADPodIdentityTaintKey {
				t.Errorf("expected funcMap invocation of GetAADPodIdentityTaintKey to return %s, instead got %s", common.AADPodIdentityTaintKey, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetMode"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedMode {
				t.Errorf("expected funcMap invocation of GetMode to return %s, instead got %s", c.expectedMode, ret[0].Interface())
			}
		})
	}
}

func TestGetComponentFuncMap(t *testing.T) {
	specConfig := api.AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	k8sComponentsByVersionMap := api.GetK8sComponentsByVersionMap(&api.KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR})
	//specConfigAzureStack := api.AzureCloudSpecEnvMap["AzureStackCloud"].KubernetesSpecConfig
	cases := []struct {
		name                                              string
		components                                        []api.KubernetesComponent
		cs                                                *api.ContainerService
		expectedCPUReqs                                   string
		expectedCPULimits                                 string
		expectedMemReqs                                   string
		expectedMemLimits                                 string
		expectedIsAzureStackCloud                         bool
		expectedIsKubernetesVersionGeOneDotFifteenDotZero bool
		expectedAPIServerImage                            string
		expectedAPIServerCommand                          string
		expectedAPIServerArgs                             string
		expectedControllerManagerImage                    string
		expectedControllerManagerCommand                  string
		expectedControllerManagerArgs                     string
		expectedCloudControllerManagerImage               string
		expectedCloudControllerManagerCommand             string
		expectedCloudControllerManagerArgs                string
		expectedSchedulerImage                            string
		expectedSchedulerCommand                          string
		expectedSchedulerArgs                             string
		expectedAddonManagerImage                         string
	}{
		{
			name: "1.15",
			components: []api.KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-controller-manager\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.CloudControllerManagerComponentName],
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-scheduler\"",
					},
				},
				{
					Name:    common.AddonManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.AddonManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.AddonManagerComponentName],
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.7",
						KubernetesConfig: &api.KubernetesConfig{
							APIServerConfig: map[string]string{
								"foo": "bar",
								"baz": "bang",
							},
							ControllerManagerConfig: map[string]string{
								"this": "that",
								"quid": "ergo",
							},
							CloudControllerManagerConfig: map[string]string{
								"bugs": "bunny",
							},
							SchedulerConfig: map[string]string{
								"daffy": "duck",
								"porky": "pig",
								"elmer": "fudd",
							},
						},
					},
				},
			},
			expectedCPUReqs:           "",
			expectedCPULimits:         "",
			expectedMemReqs:           "",
			expectedMemLimits:         "",
			expectedIsAzureStackCloud: false,
			expectedIsKubernetesVersionGeOneDotFifteenDotZero: true,
			expectedAPIServerImage:                            specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
			expectedAPIServerCommand:                          "\"/hyperkube\", \"kube-apiserver\"",
			expectedAPIServerArgs:                             "\"baz=bang\", \"foo=bar\"",
			expectedControllerManagerImage:                    specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
			expectedControllerManagerCommand:                  "\"/hyperkube\", \"kube-controller-manager\"",
			expectedControllerManagerArgs:                     "\"quid=ergo\", \"this=that\"",
			expectedCloudControllerManagerImage:               specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.CloudControllerManagerComponentName],
			expectedCloudControllerManagerCommand:             "\"cloud-controller-manager\"",
			expectedCloudControllerManagerArgs:                "\"bugs=bunny\"",
			expectedSchedulerImage:                            specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
			expectedSchedulerCommand:                          "\"/hyperkube\", \"kube-scheduler\"",
			expectedSchedulerArgs:                             "\"daffy=duck\", \"elmer=fudd\", \"porky=pig\"",
			expectedAddonManagerImage:                         specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.AddonManagerComponentName],
		},
		{
			name: "Azure Stack",
			components: []api.KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube] + common.AzureStackSuffix,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube] + common.AzureStackSuffix,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-controller-manager\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.CloudControllerManagerComponentName],
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube] + common.AzureStackSuffix,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-scheduler\"",
					},
				},
				{
					Name:    common.AddonManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.AddonManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.AddonManagerComponentName],
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.7",
						KubernetesConfig: &api.KubernetesConfig{
							APIServerConfig: map[string]string{
								"foo": "bar",
								"baz": "bang",
							},
							ControllerManagerConfig: map[string]string{
								"this": "that",
								"quid": "ergo",
							},
							CloudControllerManagerConfig: map[string]string{
								"bugs": "bunny",
							},
							SchedulerConfig: map[string]string{
								"daffy": "duck",
								"porky": "pig",
								"elmer": "fudd",
							},
						},
					},
					CustomCloudProfile: &api.CustomCloudProfile{
						Environment: &azure.Environment{
							Name: "AzureStackCloud",
						},
					},
				},
			},
			expectedCPUReqs:           "",
			expectedCPULimits:         "",
			expectedMemReqs:           "",
			expectedMemLimits:         "",
			expectedIsAzureStackCloud: true,
			expectedIsKubernetesVersionGeOneDotFifteenDotZero: true,
			expectedAPIServerImage:                            specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube] + common.AzureStackSuffix,
			expectedAPIServerCommand:                          "\"/hyperkube\", \"kube-apiserver\"",
			expectedAPIServerArgs:                             "\"baz=bang\", \"foo=bar\"",
			expectedControllerManagerImage:                    specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube] + common.AzureStackSuffix,
			expectedControllerManagerCommand:                  "\"/hyperkube\", \"kube-controller-manager\"",
			expectedControllerManagerArgs:                     "\"quid=ergo\", \"this=that\"",
			expectedCloudControllerManagerImage:               specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.CloudControllerManagerComponentName],
			expectedCloudControllerManagerCommand:             "\"cloud-controller-manager\"",
			expectedCloudControllerManagerArgs:                "\"bugs=bunny\"",
			expectedSchedulerImage:                            specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube] + common.AzureStackSuffix,
			expectedSchedulerCommand:                          "\"/hyperkube\", \"kube-scheduler\"",
			expectedSchedulerArgs:                             "\"daffy=duck\", \"elmer=fudd\", \"porky=pig\"",
			expectedAddonManagerImage:                         specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.AddonManagerComponentName],
		},
		{
			name: "1.14",
			components: []api.KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-controller-manager\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.CloudControllerManagerComponentName],
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-scheduler\"",
					},
				},
				{
					Name:    common.AddonManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []api.KubernetesContainerSpec{
						{
							Name:  common.AddonManagerComponentName,
							Image: specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.AddonManagerComponentName],
						},
					},
				},
			},
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.14.7",
						KubernetesConfig: &api.KubernetesConfig{
							APIServerConfig: map[string]string{
								"foo": "bar",
								"baz": "bang",
							},
							ControllerManagerConfig: map[string]string{
								"this": "that",
								"quid": "ergo",
							},
							CloudControllerManagerConfig: map[string]string{
								"bugs": "bunny",
							},
							SchedulerConfig: map[string]string{
								"daffy": "duck",
								"porky": "pig",
								"elmer": "fudd",
							},
						},
					},
				},
			},
			expectedCPUReqs:           "",
			expectedCPULimits:         "",
			expectedMemReqs:           "",
			expectedMemLimits:         "",
			expectedIsAzureStackCloud: false,
			expectedIsKubernetesVersionGeOneDotFifteenDotZero: false,
			expectedAPIServerImage:                            specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
			expectedAPIServerCommand:                          "\"/hyperkube\", \"kube-apiserver\"",
			expectedAPIServerArgs:                             "\"baz=bang\", \"foo=bar\"",
			expectedControllerManagerImage:                    specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
			expectedControllerManagerCommand:                  "\"/hyperkube\", \"kube-controller-manager\"",
			expectedControllerManagerArgs:                     "\"quid=ergo\", \"this=that\"",
			expectedCloudControllerManagerImage:               specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.CloudControllerManagerComponentName],
			expectedCloudControllerManagerCommand:             "\"cloud-controller-manager\"",
			expectedCloudControllerManagerArgs:                "\"bugs=bunny\"",
			expectedSchedulerImage:                            specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.Hyperkube],
			expectedSchedulerCommand:                          "\"/hyperkube\", \"kube-scheduler\"",
			expectedSchedulerArgs:                             "\"daffy=duck\", \"elmer=fudd\", \"porky=pig\"",
			expectedAddonManagerImage:                         specConfig.KubernetesImageBase + k8sComponentsByVersionMap["1.15.7"][common.AddonManagerComponentName],
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			for _, component := range c.components {
				funcMap := getComponentFuncMap(component, c.cs)
				v := reflect.ValueOf(funcMap["ContainerCPUReqs"])
				ret := v.Call([]reflect.Value{reflect.ValueOf(component.Containers[0].Name)})
				if ret[0].Interface() != c.expectedCPUReqs {
					t.Errorf("expected funcMap invocation of ContainerCPUReqs %s to return %s, instead got %s", component.Containers[0].Name, c.expectedCPUReqs, ret[0].Interface())
				}
				v = reflect.ValueOf(funcMap["ContainerCPULimits"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(component.Containers[0].Name)})
				if ret[0].Interface() != c.expectedCPULimits {
					t.Errorf("expected funcMap invocation of ContainerCPULimits %s to return %s, instead got %s", component.Containers[0].Name, c.expectedCPULimits, ret[0].Interface())
				}
				v = reflect.ValueOf(funcMap["ContainerMemReqs"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(component.Containers[0].Name)})
				if ret[0].Interface() != c.expectedMemReqs {
					t.Errorf("expected funcMap invocation of ContainerMemReqs %s to return %s, instead got %s", component.Containers[0].Name, c.expectedMemReqs, ret[0].Interface())
				}
				v = reflect.ValueOf(funcMap["ContainerMemLimits"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(component.Containers[0].Name)})
				if ret[0].Interface() != c.expectedMemLimits {
					t.Errorf("expected funcMap invocation of ContainerMemLimits %s to return %s, instead got %s", component.Containers[0].Name, c.expectedMemLimits, ret[0].Interface())
				}
				v = reflect.ValueOf(funcMap["IsAzureStackCloud"])
				ret = v.Call(make([]reflect.Value, 0))
				if ret[0].Interface() != c.expectedIsAzureStackCloud {
					t.Errorf("expected funcMap invocation of IsAzureStackCloud to return %t, instead got %t", c.expectedIsAzureStackCloud, ret[0].Interface())
				}
				v = reflect.ValueOf(funcMap["IsKubernetesVersionGe"])
				ret = v.Call([]reflect.Value{reflect.ValueOf("1.15.0")})
				if ret[0].Interface() != c.expectedIsKubernetesVersionGeOneDotFifteenDotZero {
					t.Errorf("expected funcMap invocation of IsKubernetesVersionGe %s to return %t, instead got %t", c.cs.Properties.OrchestratorProfile.OrchestratorVersion, c.expectedIsKubernetesVersionGeOneDotFifteenDotZero, ret[0].Interface())
				}
				v = reflect.ValueOf(funcMap["ContainerConfig"])
				ret = v.Call([]reflect.Value{reflect.ValueOf("command")})
				command := ret[0].Interface()
				v = reflect.ValueOf(funcMap["ContainerImage"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(component.Containers[0].Name)})
				image := ret[0].Interface()
				if component.Name == common.APIServerComponentName {
					if command != c.expectedAPIServerCommand {
						t.Errorf("expected funcMap invocation of ContainerConfig %s to return %s, instead got %s", "command", c.expectedAPIServerCommand, ret[0].Interface())
					}
					if image != c.expectedAPIServerImage {
						t.Errorf("expected funcMap invocation of ContainerImage %s to return %s, instead got %s", component.Containers[0].Name, c.expectedAPIServerImage, ret[0].Interface())
					}
					v = reflect.ValueOf(funcMap["GetAPIServerArgs"])
					ret = v.Call(make([]reflect.Value, 0))
					if ret[0].Interface() != c.expectedAPIServerArgs {
						t.Errorf("expected funcMap invocation of GetAPIServerArgs to return %s, instead got %s", c.expectedAPIServerArgs, ret[0].Interface())
					}
				}
				if component.Name == common.ControllerManagerComponentName {
					if command != c.expectedControllerManagerCommand {
						t.Errorf("expected funcMap invocation of ContainerConfig %s to return %s, instead got %s", "command", c.expectedControllerManagerCommand, ret[0].Interface())
					}
					if image != c.expectedControllerManagerImage {
						t.Errorf("expected funcMap invocation of ContainerImage %s to return %s, instead got %s", component.Containers[0].Name, c.expectedControllerManagerImage, ret[0].Interface())
					}
					v = reflect.ValueOf(funcMap["GetControllerManagerArgs"])
					ret = v.Call(make([]reflect.Value, 0))
					if ret[0].Interface() != c.expectedControllerManagerArgs {
						t.Errorf("expected funcMap invocation of GetControllerManagerArgs to return %s, instead got %s", c.expectedControllerManagerArgs, ret[0].Interface())
					}
				}
				if component.Name == common.CloudControllerManagerComponentName {
					if command != c.expectedCloudControllerManagerCommand {
						t.Errorf("expected funcMap invocation of ContainerConfig %s to return %s, instead got %s", "command", c.expectedCloudControllerManagerCommand, ret[0].Interface())
					}
					if image != c.expectedCloudControllerManagerImage {
						t.Errorf("expected funcMap invocation of ContainerImage %s to return %s, instead got %s", component.Containers[0].Name, c.expectedCloudControllerManagerImage, ret[0].Interface())
					}
					v = reflect.ValueOf(funcMap["GetCloudControllerManagerArgs"])
					ret = v.Call(make([]reflect.Value, 0))
					if ret[0].Interface() != c.expectedCloudControllerManagerArgs {
						t.Errorf("expected funcMap invocation of GetCloudControllerManagerArgs to return %s, instead got %s", c.expectedCloudControllerManagerArgs, ret[0].Interface())
					}
				}
				if component.Name == common.SchedulerComponentName {
					if command != c.expectedSchedulerCommand {
						t.Errorf("expected funcMap invocation of ContainerConfig %s to return %s, instead got %s", "command", c.expectedSchedulerCommand, ret[0].Interface())
					}
					if image != c.expectedSchedulerImage {
						t.Errorf("expected funcMap invocation of ContainerImage %s to return %s, instead got %s", component.Containers[0].Name, c.expectedSchedulerImage, ret[0].Interface())
					}
					v = reflect.ValueOf(funcMap["GetSchedulerArgs"])
					ret = v.Call(make([]reflect.Value, 0))
					if ret[0].Interface() != c.expectedSchedulerArgs {
						t.Errorf("expected funcMap invocation of GetSchedulerArgs to return %s, instead got %s", c.expectedSchedulerArgs, ret[0].Interface())
					}
				}
				if component.Name == common.AddonManagerComponentName {
					if image != c.expectedAddonManagerImage {
						t.Errorf("expected funcMap invocation of ContainerImage %s to return %s, instead got %s", component.Containers[0].Name, c.expectedAddonManagerImage, ret[0].Interface())
					}
				}
			}
		})
	}
}
