package api

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20170831"
	"github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20180331"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
				_, _, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster([]byte(modelString), "2018-03-31", false, false, defaultK8sVersion, nil)
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
				cs, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster([]byte(modelString), "2018-03-31", false, false, defaultK8sVersion, nil)
				Expect(err).To(BeNil())
				Expect(sshAutoGenerated).To(BeFalse())

				model2 := v20180331.ManagedCluster{}
				modelString2, _ := json.Marshal(model2)
				cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster([]byte(modelString2), "2018-03-31", false, true, defaultK8sVersion, cs)

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
				_, _, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster([]byte(modelString), "2018-03-31", false, false, defaultK8sVersion, nil)
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
				cs, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster([]byte(modelString), "2018-03-31", false, false, defaultK8sVersion, nil)
				Expect(err).To(BeNil())
				Expect(sshAutoGenerated).To(BeFalse())

				model2 := v20170831.ManagedCluster{}
				modelString2, _ := json.Marshal(model2)
				cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster([]byte(modelString2), "2018-03-31", false, true, defaultK8sVersion, cs)

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
