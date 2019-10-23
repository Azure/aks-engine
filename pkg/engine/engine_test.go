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
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	v20160330 "github.com/Azure/aks-engine/pkg/api/v20160330"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
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

			certsGenerated, err := containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
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
					certsGenerated, err = containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
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
				if version != vlabs.APIVersion && version != v20160330.APIVersion {
					// Set CertificateProfile here to avoid a new one generated.
					// Kubernetes template needs certificate profile to match expected template
					// API versions other than vlabs don't expose CertificateProfile
					// API versions after v20160330 supports Kubernetes
					containerService.Properties.CertificateProfile = &api.CertificateProfile{}
					addTestCertificateProfile(containerService.Properties.CertificateProfile)
				}
			}
		} else {
			if version != vlabs.APIVersion && version != v20160330.APIVersion {
				// Set CertificateProfile here to avoid a new one generated.
				// Kubernetes template needs certificate profile to match expected template
				// API versions other than vlabs don't expose CertificateProfile
				// API versions after v20160330 supports Kubernetes
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

			certsGenerated, err := containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
				IsScale:    false,
				IsUpgrade:  false,
				PkiKeySize: helpers.DefaultPkiKeySize,
			})
			if certsGenerated {
				t.Errorf("cert generation unexpected for %s", containerService.Properties.OrchestratorProfile.OrchestratorType)
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
					certsGenerated, err = containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
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
				if version != vlabs.APIVersion && version != v20160330.APIVersion {
					// Set CertificateProfile here to avoid a new one generated.
					// Kubernetes template needs certificate profile to match expected template
					// API versions other than vlabs don't expose CertificateProfile
					// API versions after v20160330 supports Kubernetes
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
	i18n.Initialize(locale)

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
	containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
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
	i18n.Initialize(locale)

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

	expected := `New-Item -ItemType Directory -Force -Path "$env:SystemDrive:/AzureData/extensions/fooExtension" ; Invoke-WebRequest -Uri "fooRootURLextensions/fooExtension/1.0/fooBar Script?fooURLQuery" -OutFile "$env:SystemDrive:/AzureData/extensions/fooExtension/fooBar Script" ; powershell "$env:SystemDrive:/AzureData/extensions/fooExtension/fooBar Script $preprovisionExtensionParams"
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
