// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"encoding/json"

	v20170831 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20170831"
	v20180331 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20180331"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestLoadContainerServiceFromFile(t *testing.T) {
	existingContainerService := &ContainerService{Name: "test",
		Properties: &Properties{OrchestratorProfile: &OrchestratorProfile{OrchestratorType: Kubernetes, OrchestratorVersion: "1.7.16"}}}

	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	i18n.Initialize(locale)
	apiloader := &Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	containerService, _, err := apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170701/kubernetes.json", true, false, existingContainerService)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != "1.8.12" {
		t.Errorf("Failed to set orcherstator version when it is set in the json, expected 1.8.12 but got %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170701/kubernetes-default-version.json", true, false, existingContainerService)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != "1.7.16" {
		t.Errorf("Failed to set orcherstator version when it is not set in the json, got %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170131/kubernetes.json", true, false, existingContainerService)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != "1.7.16" {
		t.Errorf("Failed to set orcherstator version when it is not set in the json, got %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20160930/kubernetes.json", true, false, existingContainerService)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != "1.7.16" {
		t.Errorf("Failed to set orcherstator version when it is not set in the json, got %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170701/kubernetes-default-version.json", true, false, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(false) {
		t.Errorf("Failed to set orcherstator version when it is not set in the json API v20170701, got %s but expected %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion, common.GetDefaultKubernetesVersion(false))
	}

	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170701/kubernetes-win-default-version.json", true, false, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(true) {
		t.Errorf("Failed to set orcherstator version to windows default when it is not set in the json API v20170701, got %s but expected %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion, common.GetDefaultKubernetesVersion(true))
	}

	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170131/kubernetes.json", true, false, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(false) {
		t.Errorf("Failed to set orcherstator version when it is not set in the json API v20170131, got %s but expected %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion, common.GetDefaultKubernetesVersion(false))
	}

	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170131/kubernetes-win.json", true, false, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != common.GetDefaultKubernetesVersion(true) {
		t.Errorf("Failed to set orcherstator version to windows default when it is not set in the json API v20170131, got %s but expected %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion, common.GetDefaultKubernetesVersion(true))
	}

	// Test ACS scale scenario
	existingContainerService.Properties.OrchestratorProfile.OrchestratorVersion = "1.8.12"
	containerService, _, err = apiloader.LoadContainerServiceFromFile("../engine/testdata/v20170701/kubernetes.json", true, true, existingContainerService)
	if err != nil {
		t.Error(err.Error())
	}
	if containerService.Properties.OrchestratorProfile.OrchestratorVersion != "1.8.12" {
		t.Errorf("Failed to set orcherstator version when it is set in the json, expected 1.8.12 but got %s", containerService.Properties.OrchestratorProfile.OrchestratorVersion)
	}

	// Test error scenario
	containerService, _, err = apiloader.LoadContainerServiceFromFile("../this-file-doesnt-exist.json", true, false, nil)
	if err == nil {
		t.Error(err.Error())
	}
}

func TestLoadContainerServiceForAgentPoolOnlyCluster(t *testing.T) {
	var _ = Describe("create/update cluster operations", func() {
		locale := gotext.NewLocale(path.Join("../../..", "../../..", "translations"), "en_US")
		i18n.Initialize(locale)
		apiloader := &Apiloader{
			Translator: &i18n.Translator{
				Locale: locale,
			},
		}
		k8sVersions := common.GetAllSupportedKubernetesVersions(true, false)
		defaultK8sVersion := common.GetDefaultKubernetesVersion(false)

		Context("v20180331", func() {
			It("it should return error if managed cluster body is empty", func() {

				model := v20180331.ManagedCluster{}

				modelString, _ := json.Marshal(model)
				_, _, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString, "2018-03-31", false, false, defaultK8sVersion, nil)
				Expect(err).NotTo(BeNil())
			})

			It("it should merge if managed cluster body is empty and trying to update", func() {
				model := v20180331.ManagedCluster{
					Name: "myaks",
					Properties: &v20180331.Properties{
						DNSPrefix:         "myaks",
						KubernetesVersion: k8sVersions[0],
						AgentPoolProfiles: []*v20180331.AgentPoolProfile{
							{
								Name:           "agentpool1",
								Count:          3,
								VMSize:         "Standard_DS2_v2",
								OSDiskSizeGB:   0,
								StorageProfile: "ManagedDisk",
							},
						},
						ServicePrincipalProfile: &v20180331.ServicePrincipalProfile{
							ClientID: "clientID",
							Secret:   "clientSecret",
						},
					},
				}
				modelString, _ := json.Marshal(model)
				cs, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString, "2018-03-31", false, false, defaultK8sVersion, nil)
				Expect(err).To(BeNil())
				Expect(sshAutoGenerated).To(BeFalse())

				model2 := v20180331.ManagedCluster{}
				modelString2, _ := json.Marshal(model2)
				cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString2, "2018-03-31", false, true, defaultK8sVersion, cs)

				Expect(err).To(BeNil())
				// ssh key should not be re-generated
				Expect(sshAutoGenerated).To(BeFalse())
				Expect(cs2.Properties.AgentPoolProfiles).NotTo(BeNil())
				Expect(cs2.Properties.LinuxProfile).NotTo(BeNil())
				Expect(cs2.Properties.WindowsProfile).NotTo(BeNil())
				Expect(cs2.Properties.ServicePrincipalProfile).NotTo(BeNil())
				Expect(cs2.Properties.HostedMasterProfile).NotTo(BeNil())
				Expect(cs2.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
				Expect(cs2.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			})
		})

		Context("20170831", func() {
			It("it should return error if managed cluster body is empty", func() {

				model := v20170831.ManagedCluster{}

				modelString, _ := json.Marshal(model)
				_, _, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString, "2018-03-31", false, false, defaultK8sVersion, nil)
				Expect(err).NotTo(BeNil())
			})

			It("it should merge if managed cluster body is empty and trying to update", func() {
				model := v20170831.ManagedCluster{
					Name: "myaks",
					Properties: &v20170831.Properties{
						DNSPrefix:         "myaks",
						KubernetesVersion: k8sVersions[0],
						AgentPoolProfiles: []*v20170831.AgentPoolProfile{
							{
								Name:           "agentpool1",
								Count:          3,
								VMSize:         "Standard_DS2_v2",
								OSDiskSizeGB:   0,
								StorageProfile: "ManagedDisk",
							},
						},
						ServicePrincipalProfile: &v20170831.ServicePrincipalProfile{
							ClientID: "clientID",
							Secret:   "clientSecret",
						},
					},
				}
				modelString, _ := json.Marshal(model)
				cs, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString, "2018-03-31", false, false, defaultK8sVersion, nil)
				Expect(err).To(BeNil())
				Expect(sshAutoGenerated).To(BeFalse())

				model2 := v20170831.ManagedCluster{}
				modelString2, _ := json.Marshal(model2)
				cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString2, "2018-03-31", false, true, defaultK8sVersion, cs)

				Expect(err).To(BeNil())
				// ssh key should not be re-generated
				Expect(sshAutoGenerated).To(BeFalse())
				Expect(cs2.Properties.AgentPoolProfiles).NotTo(BeNil())
				Expect(cs2.Properties.LinuxProfile).NotTo(BeNil())
				Expect(cs2.Properties.WindowsProfile).NotTo(BeNil())
				Expect(cs2.Properties.ServicePrincipalProfile).NotTo(BeNil())
				Expect(cs2.Properties.HostedMasterProfile).NotTo(BeNil())
				Expect(cs2.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			})
		})
	})
}

func TestLoadContainerServiceWithNilProperties(t *testing.T) {
	jsonWithoutProperties := `{
        "type": "Microsoft.ContainerService/managedClusters",
        "name": "[parameters('clusterName')]",
        "apiVersion": "2017-07-01",
        "location": "[resourceGroup().location]"
        }`

	tmpFile, err := ioutil.TempFile("", "containerService-invalid")
	fileName := tmpFile.Name()
	defer os.Remove(fileName)

	err = ioutil.WriteFile(fileName, []byte(jsonWithoutProperties), os.ModeAppend)

	apiloader := &Apiloader{}
	existingContainerService := &ContainerService{Name: "test",
		Properties: &Properties{OrchestratorProfile: &OrchestratorProfile{OrchestratorType: Kubernetes, OrchestratorVersion: "1.7.16"}}}
	_, _, err = apiloader.LoadContainerServiceFromFile(fileName, true, false, existingContainerService)
	if err == nil {
		t.Errorf("Expected error to be thrown")
	}
	expectedMsg := "missing ContainerService Properties"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error with message %s but got %s", expectedMsg, err.Error())
	}
}

func TestLoadContainerServiceWithEmptyLocationCustomCloud(t *testing.T) {
	jsonWithoutlocationcustomcloud := `{
		"apiVersion": "vlabs",
		"properties": {
			"orchestratorProfile": {
				"orchestratorType": "Kubernetes",
				"orchestratorRelease": "1.11",
				"kubernetesConfig": {
					"kubernetesImageBase": "msazurestackqa/",
					"useInstanceMetadata": false,
					"networkPolicy": "none"
				}
			},
			"customCloudProfile": {
				"environment": {
					"name": "AzureStackCloud",
					"managementPortalURL": "",
					"publishSettingsURL": "",
					"serviceManagementEndpoint": "https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189",
					"resourceManagerEndpoint": "https://management.local.azurestack.external/",
					"activeDirectoryEndpoint": "https://login.windows.net/",
					"galleryEndpoint": "https://portal.local.azurestack.external:30015/",
					"keyVaultEndpoint": "",
					"graphEndpoint": "https://graph.windows.net/",
					"storageEndpointSuffix": "local.azurestack.external",
					"sqlDatabaseDNSSuffix": "",
					"trafficManagerDNSSuffix": "",
					"keyVaultDNSSuffix": "vault.local.azurestack.external",
					"serviceBusEndpointSuffix": "",
					"serviceManagementVMDNSSuffix": "cloudapp.net",
					"resourceManagerVMDNSSuffix": "cloudapp.azurestack.external",
					"containerRegistryDNSSuffix": ""
				}
			},
			"masterProfile": {
				"dnsPrefix": "k111006",
				"distro": "ubuntu",
				"osDiskSizeGB": 200,
				"count": 3,
				"vmSize": "Standard_D2_v2"
			},
			"agentPoolProfiles": [
				{
					"name": "linuxpool",
					"osDiskSizeGB": 200,
					"count": 3,
					"vmSize": "Standard_D2_v2",
					"distro": "ubuntu",
					"availabilityProfile": "AvailabilitySet",
					"AcceleratedNetworkingEnabled": false
				}
			],
			"linuxProfile": {
				"adminUsername": "azureuser",
				"ssh": {
					"publicKeys": [
						{
							"keyData": "ssh-rsa PblicKey"
						}
					]
				}
			},
			"servicePrincipalProfile": {
				"clientId": "clientId",
				"secret": "secret"
			}
		}
	}`

	tmpFile, err := ioutil.TempFile("", "containerService-nolocation")
	fileName := tmpFile.Name()
	defer os.Remove(fileName)

	err = ioutil.WriteFile(fileName, []byte(jsonWithoutlocationcustomcloud), os.ModeAppend)

	apiloader := &Apiloader{}
	_, _, err = apiloader.LoadContainerServiceFromFile(fileName, true, false, nil)
	if err == nil {
		t.Errorf("Expected error for missing loation to be thrown")
	}
	expectedMsg := "missing ContainerService Location"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error with message %s but got %s", expectedMsg, err.Error())
	}

	jsonWithoutlocationpubliccloud := `{
		"apiVersion": "vlabs",
		"properties": {
			"orchestratorProfile": {
				"orchestratorType": "Kubernetes",
				"orchestratorRelease": "1.11",
				"kubernetesConfig": {
					"kubernetesImageBase": "msazurestackqa/",
					"useInstanceMetadata": false,
					"networkPolicy": "none"
				}
			},
			"masterProfile": {
				"dnsPrefix": "k111006",
				"distro": "ubuntu",
				"osDiskSizeGB": 200,
				"count": 3,
				"vmSize": "Standard_D2_v2"
			},
			"agentPoolProfiles": [
				{
					"name": "linuxpool",
					"osDiskSizeGB": 200,
					"count": 3,
					"vmSize": "Standard_D2_v2",
					"distro": "ubuntu",
					"availabilityProfile": "AvailabilitySet",
					"AcceleratedNetworkingEnabled": false
				}
			],
			"linuxProfile": {
				"adminUsername": "azureuser",
				"ssh": {
					"publicKeys": [
						{
							"keyData": "ssh-rsa PblicKey"
						}
					]
				}
			},
			"servicePrincipalProfile": {
				"clientId": "clientId",
				"secret": "secret"
			}
		}
	}`

	tmpFilewithoutlocationpubliccloud, err := ioutil.TempFile("", "containerService-nolocationpubliccloud")
	fileNamewithoutlocationpubliccloud := tmpFilewithoutlocationpubliccloud.Name()
	defer os.Remove(fileNamewithoutlocationpubliccloud)

	err = ioutil.WriteFile(fileNamewithoutlocationpubliccloud, []byte(jsonWithoutlocationpubliccloud), os.ModeAppend)

	apiloaderwithoutlocationpubliccloud := &Apiloader{}
	_, _, err = apiloaderwithoutlocationpubliccloud.LoadContainerServiceFromFile(fileNamewithoutlocationpubliccloud, true, false, nil)
	if err != nil {
		t.Errorf("Expected no error for missing loation for public cloud to be thrown")
	}
}

func TestDeserializeContainerService(t *testing.T) {
	apiloader := &Apiloader{
		Translator: nil,
	}

	// Test AKS Engine api model
	cs, version, err := apiloader.DeserializeContainerService([]byte(exampleAPIModel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	if version != vlabs.APIVersion {
		t.Fatalf("expected apiVersion %s, instead got: %s", vlabs.APIVersion, version)
	}
	if cs.Properties.OrchestratorProfile.OrchestratorType != Kubernetes {
		t.Fatalf("expected cs.Properties.OrchestratorProfile.OrchestratorType %s, instead got: %s", Kubernetes, cs.Properties.OrchestratorProfile.OrchestratorType)
	}

	// Test AKS api model
	cs, version, err = apiloader.DeserializeContainerService([]byte(exampleAKSAPIModel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	if version != v20180331.APIVersion {
		t.Fatalf("expected apiVersion %s, instead got: %s", v20180331.APIVersion, version)
	}
	if cs.Properties.OrchestratorProfile.OrchestratorType != Kubernetes {
		t.Fatalf("expected cs.Properties.OrchestratorProfile.OrchestratorType %s, instead got: %s", Kubernetes, cs.Properties.OrchestratorProfile.OrchestratorType)
	}
	if cs.Properties.MasterProfile != nil {
		t.Fatalf("expected nil MasterProfile for AKS container service object")
	}

	// Test error case
	_, _, err = apiloader.DeserializeContainerService([]byte(`{thisisnotson}`), false, false, nil)
	if err == nil {
		t.Fatalf("expected error from malformed api model input")
	}
}

func TestLoadDefaultContainerServiceProperties(t *testing.T) {
	m, p := LoadDefaultContainerServiceProperties()

	if m.APIVersion != defaultAPIVersion {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return API version %s, instead got %s", defaultAPIVersion, m.APIVersion)
	}

	if p.OrchestratorProfile.OrchestratorType != defaultOrchestrator {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s OrchestratorProfile.OrchestratorType, instead got %s", Kubernetes, p.OrchestratorProfile.OrchestratorType)
	}

	if p.MasterProfile.Count != defaultMasterCount {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d MasterProfile.Count, instead got %d", defaultMasterCount, p.MasterProfile.Count)
	}

	if p.MasterProfile.VMSize != defaultVMSize {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s MasterProfile.VMSize, instead got %s", defaultVMSize, p.MasterProfile.VMSize)
	}

	if p.MasterProfile.OSDiskSizeGB != defaultOSDiskSizeGB {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d MasterProfile.OSDiskSizeGB, instead got %d", defaultOSDiskSizeGB, p.MasterProfile.OSDiskSizeGB)
	}

	if len(p.AgentPoolProfiles) != 1 {
		t.Errorf("Expected 1 agent pool, instead got %d", len(p.AgentPoolProfiles))
	}

	if p.AgentPoolProfiles[0].Name != defaultAgentPoolName {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s AgentPoolProfiles[0].Name, instead got %s", defaultAgentPoolName, p.AgentPoolProfiles[0].Name)
	}

	if p.AgentPoolProfiles[0].Count != defaultAgentCount {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d AgentPoolProfiles[0].Count, instead got %d", defaultAgentCount, p.AgentPoolProfiles[0].Count)
	}

	if p.AgentPoolProfiles[0].VMSize != defaultVMSize {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s AgentPoolProfiles[0].VMSize, instead got %s", defaultVMSize, p.AgentPoolProfiles[0].VMSize)
	}

	if p.AgentPoolProfiles[0].OSDiskSizeGB != defaultOSDiskSizeGB {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d AgentPoolProfiles[0].OSDiskSizeGB, instead got %d", defaultOSDiskSizeGB, p.AgentPoolProfiles[0].OSDiskSizeGB)
	}

	if p.LinuxProfile.AdminUsername != defaultAdminUser {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s LinuxProfile.AdminAdminUsernameUsername, instead got %s", defaultAdminUser, p.LinuxProfile.AdminUsername)
	}
}
