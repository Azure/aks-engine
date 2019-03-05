// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	. "github.com/Azure/aks-engine/pkg/test"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

const TestAKSEngineVersion = "1.0.0"

type fakeUpgradeWorkflow struct {
	RunUpgradeError error
	ValidateError   error
}

func (workflow fakeUpgradeWorkflow) RunUpgrade() error {
	if workflow.RunUpgradeError != nil {
		return workflow.RunUpgradeError
	}
	return nil
}

func (workflow fakeUpgradeWorkflow) Validate() error {
	if workflow.ValidateError != nil {
		return workflow.ValidateError
	}
	return nil
}

func TestUpgradeCluster(t *testing.T) {
	RunSpecsWithReporters(t, "kubernetesupgrade", "Server Suite")
}

var _ = Describe("Upgrade Kubernetes cluster tests", func() {
	AfterEach(func() {
		// delete temp template directory
		os.RemoveAll("_output")
	})

	It("Should succeed when cluster VMs are missing expected tags during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 1, 1, false)
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailListVirtualMachinesTags = true
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).To(BeNil())
		Expect(uc.ClusterTopology.AgentPools).NotTo(BeEmpty())

		// Clean up
		os.RemoveAll("./translations")
	})

	It("Should return error message when failing to list VMs during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 1, 1, false)
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailListVirtualMachines = true
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("Error while querying ARM for resources: ListVirtualMachines failed"))

		// Clean up
		os.RemoveAll("./translations")
	})

	It("Should return error message when failing to delete VMs during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 1, 1, false)
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailDeleteVirtualMachine = true
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("DeleteVirtualMachine failed"))
	})

	It("Should return error message when failing to deploy template during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 1, 1, false)
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailDeployTemplate = true
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("DeployTemplate failed"))
	})

	It("Should return error message when failing to get a virtual machine during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 1, 6, false)
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailGetVirtualMachine = true
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("GetVirtualMachine failed"))
	})

	It("Should return error message when failing to get storage client during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 5, 1, false)
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailGetStorageClient = true
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("GetStorageClient failed"))
	})

	It("Should return error message when failing to delete network interface during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 3, 2, false)
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailDeleteNetworkInterface = true

		uc.Client = &mockClient
		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("DeleteNetworkInterface failed"))
	})

	It("Should return error message when failing to delete role assignment during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 3, 2, false)
		cs.Properties.OrchestratorProfile.KubernetesConfig = &api.KubernetesConfig{}
		cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailDeleteRoleAssignment = true
		mockClient.ShouldSupportVMIdentity = true
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("DeleteRoleAssignmentByID failed"))
	})

	Context("When upgrading a cluster with VMSS VMs", func() {
		var (
			cs         *api.ContainerService
			uc         UpgradeCluster
			mockClient armhelpers.MockAKSEngineClient
		)

		BeforeEach(func() {
			mockClient = armhelpers.MockAKSEngineClient{}
			cs = api.CreateMockContainerService("testcluster", "1.9.10", 3, 3, false)
			uc = UpgradeCluster{
				Translator: &i18n.Translator{},
				Logger:     log.NewEntry(log.New()),
			}
			mockClient.FakeListVirtualMachineScaleSetsResult = func() []compute.VirtualMachineScaleSet {
				scalesetName := "scalesetName"
				sku := compute.Sku{}
				location := "eastus"
				return []compute.VirtualMachineScaleSet{
					{
						Name:     &scalesetName,
						Sku:      &sku,
						Location: &location,
					},
				}
			}
			uc.Client = &mockClient
			uc.ClusterTopology = ClusterTopology{}
			uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
			uc.ResourceGroup = "TestRg"
			uc.DataModel = cs
			uc.NameSuffix = "12345678"
			uc.UpgradeWorkFlow = fakeUpgradeWorkflow{}
		})
		It("Should skip VMs that are already on desired version", func() {
			mockClient.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
				return []compute.VirtualMachineScaleSetVM{
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.10"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.9"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.7"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.10"),
				}
			}
			uc.Force = false

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(uc.AgentPoolScaleSetsToUpgrade[0].VMsToUpgrade).To(HaveLen(2))
		})
		It("Should skip VMs that cannot determine version", func() {
			mockClient.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
				return []compute.VirtualMachineScaleSetVM{
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.7"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:"),
				}
			}
			uc.Force = false

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(uc.AgentPoolScaleSetsToUpgrade[0].VMsToUpgrade).To(HaveLen(1))
		})
		It("Should not skip VMs that cannot determine version when using Force", func() {
			mockClient.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
				return []compute.VirtualMachineScaleSetVM{
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.7"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:"),
				}
			}
			uc.Force = true

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(uc.AgentPoolScaleSetsToUpgrade[0].VMsToUpgrade).To(HaveLen(2))
		})
		It("Should not skip any VMs when using Force", func() {
			mockClient.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
				return []compute.VirtualMachineScaleSetVM{
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.10"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.9"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.7"),
					mockClient.MakeFakeVirtualMachineScaleSetVM("Kubernetes:1.9.10"),
				}
			}
			uc.Force = true

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(uc.AgentPoolScaleSetsToUpgrade[0].VMsToUpgrade).To(HaveLen(4))
		})
	})

	Context("When upgrading a cluster with AvailibilitySets VMs", func() {
		var (
			cs               *api.ContainerService
			uc               UpgradeCluster
			mockClient       armhelpers.MockAKSEngineClient
			versionMapBackup map[string]bool
		)

		AfterEach(func() {
			common.AllKubernetesSupportedVersions = versionMapBackup
		})

		BeforeEach(func() {
			versionMapBackup = common.AllKubernetesSupportedVersions
			mockClient = armhelpers.MockAKSEngineClient{}
			cs = api.CreateMockContainerService("testcluster", "1.9.10", 3, 3, false)
			uc = UpgradeCluster{
				Translator: &i18n.Translator{},
				Logger:     log.NewEntry(log.New()),
			}

			uc.Client = &mockClient
			uc.ClusterTopology = ClusterTopology{}
			uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
			uc.ResourceGroup = "TestRg"
			uc.DataModel = cs
			uc.NameSuffix = "12345678"
			uc.UpgradeWorkFlow = fakeUpgradeWorkflow{}
		})
		It("Should skip VMs that are already on desired version", func() {
			mockClient.FakeListVirtualMachineResult = func() []compute.VirtualMachine {
				return []compute.VirtualMachine{
					mockClient.MakeFakeVirtualMachine("k8s-agentpool1-12345678-0", "Kubernetes:1.9.10"),
				}
			}
			uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}
			uc.Force = false

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(*uc.AgentPools["agentpool1"].AgentVMs).To(HaveLen(0))
		})
		It("Should fail when desired version target is not supported", func() {
			desiredVersion := "1.9.10"
			common.AllKubernetesSupportedVersions = map[string]bool{
				"1.9.7":        true,
				desiredVersion: false,
			}
			mockClient.FakeListVirtualMachineResult = func() []compute.VirtualMachine {
				return []compute.VirtualMachine{
					mockClient.MakeFakeVirtualMachine("k8s-agentpool1-12345678-0", "Kubernetes:1.9.7"),
				}
			}
			uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}
			uc.Force = false
			uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion = desiredVersion
			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("1.9.7 cannot be upgraded to 1.9.10"))
		})
		It("Should not fail when desired version target is not supported and force true", func() {
			desiredVersion := "1.9.10"
			common.AllKubernetesSupportedVersions = map[string]bool{
				"1.9.7":        true,
				desiredVersion: false,
			}
			mockClient.FakeListVirtualMachineResult = func() []compute.VirtualMachine {
				return []compute.VirtualMachine{
					mockClient.MakeFakeVirtualMachine("k8s-agentpool1-12345678-0", "Kubernetes:1.9.7"),
				}
			}
			uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}
			uc.Force = true
			uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion = desiredVersion
			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(*uc.AgentPools["agentpool1"].AgentVMs).To(HaveLen(1))

		})
		It("Should not skip VMs that are already on desired version when Force true", func() {
			mockClient.FakeListVirtualMachineResult = func() []compute.VirtualMachine {
				return []compute.VirtualMachine{
					mockClient.MakeFakeVirtualMachine("k8s-agentpool1-12345678-0", "Kubernetes:1.9.10"),
				}
			}
			uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}
			uc.Force = true

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(*uc.AgentPools["agentpool1"].AgentVMs).To(HaveLen(1))
		})
		It("Should skip master VMS that are already on desired version", func() {
			mockClient.FakeListVirtualMachineResult = func() []compute.VirtualMachine {
				return []compute.VirtualMachine{
					mockClient.MakeFakeVirtualMachine("k8s-master-12345678-0", "Kubernetes:1.9.10"),
				}
			}
			uc.Force = false

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(*uc.MasterVMs).To(HaveLen(0))
			Expect(*uc.UpgradedMasterVMs).To(HaveLen(1))

		})
		It("Should not skip master VMS that are already on desired version when Force is true", func() {
			mockClient.FakeListVirtualMachineResult = func() []compute.VirtualMachine {
				return []compute.VirtualMachine{
					mockClient.MakeFakeVirtualMachine("k8s-master-12345678-0", "Kubernetes:1.9.10"),
				}
			}
			uc.Force = true

			err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
			Expect(err).NotTo(HaveOccurred())
			Expect(*uc.MasterVMs).To(HaveLen(1))
			Expect(*uc.UpgradedMasterVMs).To(HaveLen(0))
		})
	})

	It("Should not fail if no managed identity is returned by azure during upgrade operation", func() {
		cs := api.CreateMockContainerService("testcluster", "1.9.11", 3, 2, false)
		cs.Properties.OrchestratorProfile.KubernetesConfig = &api.KubernetesConfig{}
		cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
		uc := UpgradeCluster{
			Translator: &i18n.Translator{},
			Logger:     log.NewEntry(log.New()),
		}

		mockClient := armhelpers.MockAKSEngineClient{}
		uc.Client = &mockClient

		uc.ClusterTopology = ClusterTopology{}
		uc.SubscriptionID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
		uc.ResourceGroup = "TestRg"
		uc.DataModel = cs
		uc.NameSuffix = "12345678"
		uc.AgentPoolsToUpgrade = map[string]bool{"agentpool1": true}

		err := uc.UpgradeCluster(&mockClient, "kubeConfig", TestAKSEngineVersion)
		Expect(err).To(BeNil())
	})

	It("Tests GetLastVMNameInVMSS", func() {
		ctx := context.Background()

		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FakeListVirtualMachineScaleSetsResult = func() []compute.VirtualMachineScaleSet {
			scalesetName := "scalesetName"
			sku := compute.Sku{}
			location := "eastus"
			return []compute.VirtualMachineScaleSet{
				{
					Name:     &scalesetName,
					Sku:      &sku,
					Location: &location,
				},
			}
		}
		mockClient.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
			return []compute.VirtualMachineScaleSetVM{
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", "aks-agentnode1-123456-vmss000002"),
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", "aks-agentnode1-123456-vmss000003"),
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", "aks-agentnode1-123456-vmss000004"),
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", "aks-agentnode1-123456-vmss000005"),
			}
		}

		u := &Upgrader{}
		u.Init(&i18n.Translator{}, log.NewEntry(log.New()), ClusterTopology{}, &mockClient, "", nil, TestAKSEngineVersion)

		vmname, err := u.getLastVMNameInVMSS(ctx, "resourcegroup", "scalesetName")
		Expect(vmname).To(Equal("aks-agentnode1-123456-vmss000005"))
		Expect(err).NotTo(HaveOccurred())

		mockClient.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
			return []compute.VirtualMachineScaleSetVM{
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", "aks-agentnode1-123456-vmss000002"),
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", "aks-agentnode1-123456-vmss000003"),
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", "aks-agentnode1-123456-vmss000004"),
				mockClient.MakeFakeVirtualMachineScaleSetVMWithGivenName("Kubernetes:1.9.10", ""),
			}
		}
		u.Init(&i18n.Translator{}, log.NewEntry(log.New()), ClusterTopology{}, &mockClient, "", nil, TestAKSEngineVersion)

		vmname, err = u.getLastVMNameInVMSS(ctx, "resourcegroup", "scalesetName")
		Expect(vmname).To(Equal(""))
		Expect(err).To(HaveOccurred())

		mockClient.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
			return []compute.VirtualMachineScaleSetVM{}
		}
		u.Init(&i18n.Translator{}, log.NewEntry(log.New()), ClusterTopology{}, &mockClient, "", nil, TestAKSEngineVersion)

		vmname, err = u.getLastVMNameInVMSS(ctx, "resourcegroup", "scalesetName")
		Expect(vmname).To(Equal(""))
		Expect(err).To(HaveOccurred())
	})

	It("Tests CopyCustomPropertiesToNewNode", func() {
		u := &Upgrader{}

		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.FailGetNode = true

		u.Init(&i18n.Translator{}, log.NewEntry(log.New()), ClusterTopology{}, nil, "", nil, TestAKSEngineVersion)
		err := u.copyCustomPropertiesToNewNode(mockClient.MockKubernetesClient, "oldNodeName", "newNodeName")
		Expect(err).Should(HaveOccurred())

		oldNode := &v1.Node{}
		oldNode.Annotations = map[string]string{}
		oldNode.Annotations["ann1"] = "val1"
		oldNode.Annotations["ann2"] = "val2"
		oldNode.Annotations["customAnn"] = "customVal"

		oldNode.Labels = map[string]string{}
		oldNode.Labels["label1"] = "val1"
		oldNode.Labels["label2"] = "val2"
		oldNode.Labels["customLabel"] = "customVal"

		oldNode.Spec.Taints = []v1.Taint{
			{
				Key:    "key1",
				Value:  "val1",
				Effect: "NoSchedule",
			},
			{
				Key:    "key2",
				Value:  "val2",
				Effect: "NoSchedule",
			},
		}

		newNode := &v1.Node{}
		newNode.Annotations = map[string]string{}
		newNode.Annotations["ann1"] = "newval1"
		newNode.Annotations["ann2"] = "newval2"

		newNode.Labels = map[string]string{}
		newNode.Labels["label1"] = "newval1"
		newNode.Labels["label2"] = "newval2"

		mockClient.MockKubernetesClient.UpdateNodeFunc = func(node *v1.Node) (*v1.Node, error) {
			return node, nil
		}

		err = u.copyCustomNodeProperties(mockClient.MockKubernetesClient, "oldnode", oldNode, "newnode", newNode)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(newNode.Annotations)).To(Equal(3))
		Expect(newNode.Annotations["ann1"]).To(Equal("newval1"))
		Expect(newNode.Annotations["ann2"]).To(Equal("newval2"))
		Expect(newNode.Annotations["customAnn"]).To(Equal("customVal"))

		Expect(len(newNode.Labels)).To(Equal(3))
		Expect(newNode.Labels["label1"]).To(Equal("newval1"))
		Expect(newNode.Labels["label2"]).To(Equal("newval2"))
		Expect(newNode.Labels["customLabel"]).To(Equal("customVal"))

		Expect(len(newNode.Spec.Taints)).To(Equal(2))
	})
})
