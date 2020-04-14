// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v20180331_test

import (
	"github.com/Azure/aks-engine/pkg/api"
	v20180331 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20180331"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/i18n"
	. "github.com/Azure/aks-engine/pkg/test"
	"github.com/leonelquinteros/gotext"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"path"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	//RegisterFailHandler(Fail)
	RunSpecsWithReporters(t, "apiloader", "v20180331 Suite")
}

var _ = Describe("v20180331 test suite", func() {
	locale := gotext.NewLocale(path.Join("../../..", "../../..", "translations"), "en_US")
	i18n.Initialize(locale)
	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}
	k8sVersions := common.GetAllSupportedKubernetesVersions(false, false)
	defaultK8sVersion := common.GetDefaultKubernetesVersion(false)

	Context("when networkprofile is nil, enable the addon profile", func() {
		It("should merge fields properly", func() {

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
			Expect(sshAutoGenerated).To(BeTrue())
			Expect(cs.Properties.MasterProfile).To(BeNil())
			Expect(cs.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("kubenet"))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(api.DefaultKubernetesServiceCIDR))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(api.DefaultKubernetesDNSServiceIP))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(api.DefaultDockerBridgeSubnet))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(cs.Properties.AgentPoolProfiles).NotTo(BeNil())
			Expect(len(cs.Properties.AgentPoolProfiles)).To(Equal(1))
			Expect(cs.Properties.AgentPoolProfiles[0].Name).To(Equal(model.Properties.AgentPoolProfiles[0].Name))
			Expect(cs.Properties.AgentPoolProfiles[0].Count).To(Equal(model.Properties.AgentPoolProfiles[0].Count))
			Expect(cs.Properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig["--max-pods"]).To(Equal("110"))

			model2 := v20180331.ManagedCluster{
				Properties: &v20180331.Properties{
					AddonProfiles: map[string]v20180331.AddonProfile{
						"omsagent": {
							Enabled: true,
						},
					},
				},
			}

			modelString2, _ := json.Marshal(model2)
			cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString2, "2018-03-31", false, true, defaultK8sVersion, cs)

			Expect(err).To(BeNil())
			// ssh key should not be re-generated
			Expect(sshAutoGenerated).To(BeFalse())
			Expect(cs2.Properties.MasterProfile).To(BeNil())
			Expect(cs2.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("kubenet"))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(api.DefaultKubernetesServiceCIDR))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(api.DefaultKubernetesDNSServiceIP))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(api.DefaultDockerBridgeSubnet))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs2.Properties.ServicePrincipalProfile).NotTo(BeNil())
			Expect(cs2.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs2.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(len(cs2.Properties.AgentPoolProfiles)).To(Equal(1))
		})
	})

	Context("when custom vnet is used, enable the addon profile", func() {
		It("should merge fields properly", func() {
			serviceCidr := "172.17.0.0/16"
			dnsServiceIP := "172.17.0.5"
			dockerBridgeCidr := "173.0.0.0/16"

			model := v20180331.ManagedCluster{
				Name: "myaks",
				Properties: &v20180331.Properties{
					DNSPrefix:         "myaks",
					KubernetesVersion: k8sVersions[0],
					NetworkProfile: &v20180331.NetworkProfile{
						NetworkPlugin:    v20180331.Azure,
						ServiceCidr:      serviceCidr,
						DNSServiceIP:     dnsServiceIP,
						DockerBridgeCidr: dockerBridgeCidr,
					},
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
			Expect(sshAutoGenerated).To(BeTrue())
			Expect(cs.Properties.MasterProfile).To(BeNil())
			Expect(cs.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("azure"))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(serviceCidr))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(dnsServiceIP))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(dockerBridgeCidr))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(cs.Properties.AgentPoolProfiles).NotTo(BeNil())
			Expect(len(cs.Properties.AgentPoolProfiles)).To(Equal(1))
			Expect(cs.Properties.AgentPoolProfiles[0].Name).To(Equal(model.Properties.AgentPoolProfiles[0].Name))
			Expect(cs.Properties.AgentPoolProfiles[0].Count).To(Equal(model.Properties.AgentPoolProfiles[0].Count))
			Expect(cs.Properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig["--max-pods"]).To(Equal("30"))

			model2 := v20180331.ManagedCluster{
				Properties: &v20180331.Properties{
					AddonProfiles: map[string]v20180331.AddonProfile{
						"omsagent": {
							Enabled: true,
						},
					},
				},
			}

			modelString2, _ := json.Marshal(model2)
			cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString2, "2018-03-31", false, true, defaultK8sVersion, cs)

			Expect(err).To(BeNil())
			// ssh key should not be re-generated
			Expect(sshAutoGenerated).To(BeFalse())
			Expect(cs2.Properties.MasterProfile).To(BeNil())
			Expect(cs2.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("azure"))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(serviceCidr))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(dnsServiceIP))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(dockerBridgeCidr))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs2.Properties.ServicePrincipalProfile).NotTo(BeNil())
			Expect(cs2.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs2.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(len(cs2.Properties.AgentPoolProfiles)).To(Equal(1))
		})
	})

	Context("when custom vnet is used, scale the cluster", func() {
		It("should merge fields properly", func() {
			serviceCidr := "172.17.0.0/16"
			dnsServiceIP := "172.17.0.5"
			dockerBridgeCidr := "173.0.0.0/16"

			model := v20180331.ManagedCluster{
				Name: "myaks",
				Properties: &v20180331.Properties{
					DNSPrefix:         "myaks",
					KubernetesVersion: k8sVersions[0],
					NetworkProfile: &v20180331.NetworkProfile{
						NetworkPlugin:    v20180331.Azure,
						ServiceCidr:      serviceCidr,
						DNSServiceIP:     dnsServiceIP,
						DockerBridgeCidr: dockerBridgeCidr,
					},
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
			Expect(sshAutoGenerated).To(BeTrue())
			Expect(cs.Properties.MasterProfile).To(BeNil())
			Expect(cs.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("azure"))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(serviceCidr))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(dnsServiceIP))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(dockerBridgeCidr))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(cs.Properties.AgentPoolProfiles).NotTo(BeNil())
			Expect(len(cs.Properties.AgentPoolProfiles)).To(Equal(1))
			Expect(cs.Properties.AgentPoolProfiles[0].Name).To(Equal(model.Properties.AgentPoolProfiles[0].Name))
			Expect(cs.Properties.AgentPoolProfiles[0].Count).To(Equal(model.Properties.AgentPoolProfiles[0].Count))
			Expect(cs.Properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig["--max-pods"]).To(Equal("30"))

			model2 := v20180331.ManagedCluster{
				Properties: &v20180331.Properties{
					AgentPoolProfiles: []*v20180331.AgentPoolProfile{
						{
							Name:           "agentpool1",
							Count:          6,
							VMSize:         "Standard_DS2_v2",
							OSDiskSizeGB:   0,
							StorageProfile: "ManagedDisk",
						},
					},
				},
			}

			modelString2, _ := json.Marshal(model2)
			cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString2, "2018-03-31", false, true, defaultK8sVersion, cs)

			Expect(err).To(BeNil())
			// ssh key should not be re-generated
			Expect(sshAutoGenerated).To(BeFalse())
			Expect(cs2.Properties.MasterProfile).To(BeNil())
			Expect(cs2.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("azure"))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(serviceCidr))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(dnsServiceIP))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(dockerBridgeCidr))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs2.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs2.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(len(cs2.Properties.AgentPoolProfiles)).To(Equal(1))
			Expect(cs2.Properties.AgentPoolProfiles[0].Count).To(Equal(model2.Properties.AgentPoolProfiles[0].Count))
		})
	})

	Context("when custom vnet is used, upgrade the cluster", func() {
		It("should merge fields properly", func() {
			serviceCidr := "172.17.0.0/16"
			dnsServiceIP := "172.17.0.5"
			dockerBridgeCidr := "173.0.0.0/16"

			model := v20180331.ManagedCluster{
				Name: "myaks",
				Properties: &v20180331.Properties{
					DNSPrefix:         "myaks",
					KubernetesVersion: k8sVersions[0],
					NetworkProfile: &v20180331.NetworkProfile{
						NetworkPlugin:    v20180331.Azure,
						ServiceCidr:      serviceCidr,
						DNSServiceIP:     dnsServiceIP,
						DockerBridgeCidr: dockerBridgeCidr,
					},
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
			Expect(sshAutoGenerated).To(BeTrue())
			Expect(cs.Properties.MasterProfile).To(BeNil())
			Expect(cs.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[0]))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("azure"))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(serviceCidr))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(dnsServiceIP))
			Expect(cs.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(dockerBridgeCidr))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(cs.Properties.AgentPoolProfiles).NotTo(BeNil())
			Expect(len(cs.Properties.AgentPoolProfiles)).To(Equal(1))
			Expect(cs.Properties.AgentPoolProfiles[0].Name).To(Equal(model.Properties.AgentPoolProfiles[0].Name))
			Expect(cs.Properties.AgentPoolProfiles[0].Count).To(Equal(model.Properties.AgentPoolProfiles[0].Count))
			Expect(cs.Properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig["--max-pods"]).To(Equal("30"))

			model2 := v20180331.ManagedCluster{
				Properties: &v20180331.Properties{
					KubernetesVersion: k8sVersions[1],
				},
			}

			modelString2, _ := json.Marshal(model2)
			cs2, sshAutoGenerated, err := apiloader.LoadContainerServiceForAgentPoolOnlyCluster(modelString2, "2018-03-31", false, true, defaultK8sVersion, cs)

			Expect(err).To(BeNil())
			// ssh key should not be re-generated
			Expect(sshAutoGenerated).To(BeFalse())
			Expect(cs2.Properties.MasterProfile).To(BeNil())
			Expect(cs2.Properties.LinuxProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile).NotTo(BeNil())
			Expect(cs2.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal(k8sVersions[1]))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin).To(Equal("azure"))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy).To(Equal(""))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR).To(Equal(serviceCidr))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP).To(Equal(dnsServiceIP))
			Expect(cs2.Properties.OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet).To(Equal(dockerBridgeCidr))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac).To(Equal(false))
			Expect(*cs2.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet).To(Equal(false))
			Expect(cs2.Properties.HostedMasterProfile).NotTo(BeNil())
			Expect(cs2.Properties.HostedMasterProfile.DNSPrefix).To(Equal(model.Properties.DNSPrefix))
			Expect(len(cs2.Properties.AgentPoolProfiles)).To(Equal(1))
		})
	})
})
