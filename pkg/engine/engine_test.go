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

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	v20160330 "github.com/Azure/aks-engine/pkg/api/v20160330"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
)

const (
	TestDataDir          = "./testdata"
	TestAKSEngineVersion = "1.0.0"
)

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

		certsGenerated, err := containerService.SetPropertiesDefaults(false, false)

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

		if certsGenerated {
			t.Errorf("cert generation unexpected for %s", containerService.Properties.OrchestratorProfile.OrchestratorType)
		}

		for i := 0; i < 3; i++ {
			certsGenerated, err = containerService.SetPropertiesDefaults(false, false)
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

			if certsGenerated {
				t.Errorf("cert generation unexpected for %s", containerService.Properties.OrchestratorProfile.OrchestratorType)
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
func IterateTestFilesDirectory(directory string, APIModelTestFiles *[]APIModelTestFile) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if e := IterateTestFilesDirectory(filepath.Join(directory, file.Name()), APIModelTestFiles); e != nil {
				return e
			}
		} else {
			if !strings.Contains(file.Name(), "_expected") && strings.HasSuffix(file.Name(), ".json") {
				tuple := &APIModelTestFile{}
				tuple.APIModelFilename = filepath.Join(directory, file.Name())
				*APIModelTestFiles = append(*APIModelTestFiles, *tuple)
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

func TestGetStorageAccountType(t *testing.T) {
	validPremiumVMSize := "Standard_DS2_v2"
	validStandardVMSize := "Standard_D2_v2"
	expectedPremiumTier := "Premium_LRS"
	expectedStandardTier := "Standard_LRS"
	invalidVMSize := "D2v2"

	// test premium VMSize returns premium managed disk tier
	premiumTier, err := getStorageAccountType(validPremiumVMSize)
	if err != nil {
		t.Fatalf("Invalid sizeName: %s", err)
	}

	if premiumTier != expectedPremiumTier {
		t.Fatalf("premium VM did no match premium managed storage tier")
	}

	// test standard VMSize returns standard managed disk tier
	standardTier, err := getStorageAccountType(validStandardVMSize)
	if err != nil {
		t.Fatalf("Invalid sizeName: %s", err)
	}

	if standardTier != expectedStandardTier {
		t.Fatalf("standard VM did no match standard managed storage tier")
	}

	// test invalid VMSize
	result, err := getStorageAccountType(invalidVMSize)
	if err == nil {
		t.Errorf("getStorageAccountType() = (%s, nil), want error", result)
	}
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
	containerService.SetPropertiesDefaults(false, false)
	armTemplate, _, err := templateGenerator.GenerateTemplate(containerService, DefaultGeneratorCode, TestAKSEngineVersion)
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
		"Standard_D11_v2_Promo",
		"Standard_D12",
		"Standard_D12_v2",
		"Standard_D12_v2_Promo",
		"Standard_D13",
		"Standard_D13_v2",
		"Standard_D13_v2_Promo",
		"Standard_D14",
		"Standard_D14_v2",
		"Standard_D14_v2_Promo",
		"Standard_D15_v2",
		"Standard_D16_v3",
		"Standard_D16s_v3",
		"Standard_D2",
		"Standard_D2_v2",
		"Standard_D2_v2_Promo",
		"Standard_D2_v3",
		"Standard_D2s_v3",
		"Standard_D3",
		"Standard_D32_v3",
		"Standard_D32s_v3",
		"Standard_D3_v2",
		"Standard_D3_v2_Promo",
		"Standard_D4",
		"Standard_D4_v2",
		"Standard_D4_v2_Promo",
		"Standard_D4_v3",
		"Standard_D4s_v3",
		"Standard_D5_v2",
		"Standard_D5_v2_Promo",
		"Standard_D64_v3",
		"Standard_D64s_v3",
		"Standard_D8_v3",
		"Standard_D8s_v3",
		"Standard_DS11",
		"Standard_DS11_v2",
		"Standard_DS11_v2_Promo",
		"Standard_DS12",
		"Standard_DS12_v2",
		"Standard_DS12_v2_Promo",
		"Standard_DS13",
		"Standard_DS13-2_v2",
		"Standard_DS13-4_v2",
		"Standard_DS13_v2",
		"Standard_DS13_v2_Promo",
		"Standard_DS14",
		"Standard_DS14-4_v2",
		"Standard_DS14-8_v2",
		"Standard_DS14_v2",
		"Standard_DS14_v2_Promo",
		"Standard_DS15_v2",
		"Standard_DS3",
		"Standard_DS3_v2",
		"Standard_DS3_v2_Promo",
		"Standard_DS4",
		"Standard_DS4_v2",
		"Standard_DS4_v2_Promo",
		"Standard_DS5_v2",
		"Standard_DS5_v2_Promo",
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

func TestGenerateIpList(t *testing.T) {
	count := 3
	forth := 240
	ipList := generateIPList(count, fmt.Sprintf("10.0.0.%d", forth))
	if len(ipList) != 3 {
		t.Fatalf("IP list size should be %d", count)
	}
	for i, ip := range ipList {
		expected := fmt.Sprintf("10.0.0.%d", forth+i)
		if ip != expected {
			t.Fatalf("wrong IP %s. Expected %s", ip, expected)
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
		t.Fatalf("Failed to load container service from file: %v", err)
	}
	kubeConfig, err := GenerateKubeConfig(containerService.Properties, "westus2")
	// TODO add actual kubeconfig validation
	if len(kubeConfig) < 1 {
		t.Fatalf("Got unexpected kubeconfig payload: %v", kubeConfig)
	}
	if err != nil {
		t.Fatalf("Failed to call GenerateKubeConfig with simple Kubernetes config from file: %v", testData)
	}

	p := api.Properties{}
	_, err = GenerateKubeConfig(&p, "westus2")
	if err == nil {
		t.Fatalf("Expected an error result from nil Properties child properties")
	}

	_, err = GenerateKubeConfig(nil, "westus2")
	if err == nil {
		t.Fatalf("Expected an error result from nil Properties child properties")
	}
}

func TestGetDataDisks(t *testing.T) {
	a := &api.AgentPoolProfile{
		Name: "sampleAgent",
		DiskSizesGB: []int{
			128, 128, 128, 128, 128,
		},
		StorageProfile: api.ManagedDisks,
	}
	str := getDataDisks(a)
	fmt.Println(str)

	a = &api.AgentPoolProfile{
		Name: "sampleAgentStorage",
		DiskSizesGB: []int{
			128, 128, 128, 128, 128,
		},
		StorageProfile: api.StorageAccount,
	}
	str = getDataDisks(a)
	fmt.Println(str)
}
