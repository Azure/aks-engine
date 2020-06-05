// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var validVersionsBackup map[string]bool

func setupValidVersions(validVersions map[string]bool) {
	validVersionsBackup = common.AllKubernetesSupportedVersions
	common.AllKubernetesSupportedVersions = validVersions
}

func resetValidVersions() {
	common.AllKubernetesSupportedVersions = validVersionsBackup
}

func TestUpgradeCommandShouldBeValidated(t *testing.T) {
	g := NewGomegaWithT(t)
	r := &cobra.Command{}

	cases := []struct {
		uc          *upgradeCmd
		expectedErr error
		name        string
	}{
		{
			uc: &upgradeCmd{
				resourceGroupName:           "",
				apiModelPath:                "./not/used",
				deploymentDirectory:         "",
				upgradeVersion:              "1.8.9",
				location:                    "centralus",
				timeoutInMinutes:            60,
				cordonDrainTimeoutInMinutes: 60,
			},
			expectedErr: errors.New("--resource-group must be specified"),
			name:        "NeedsResourceGroup",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:           "test",
				apiModelPath:                "./not/used",
				deploymentDirectory:         "",
				upgradeVersion:              "1.8.9",
				location:                    "",
				timeoutInMinutes:            60,
				cordonDrainTimeoutInMinutes: 60,
			},
			expectedErr: errors.New("--location must be specified"),
			name:        "NeedsLocation",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:           "test",
				apiModelPath:                "./not/used",
				deploymentDirectory:         "",
				upgradeVersion:              "",
				location:                    "southcentralus",
				timeoutInMinutes:            60,
				cordonDrainTimeoutInMinutes: 60,
			},
			expectedErr: errors.New("--upgrade-version must be specified"),
			name:        "NeedsUpgradeVersion",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:           "test",
				apiModelPath:                "",
				deploymentDirectory:         "",
				upgradeVersion:              "1.9.0",
				location:                    "southcentralus",
				timeoutInMinutes:            60,
				cordonDrainTimeoutInMinutes: 60,
			},
			expectedErr: errors.New("--api-model must be specified"),
			name:        "NeedsAPIModel",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:           "test",
				apiModelPath:                "./somefile",
				deploymentDirectory:         "aDir/anotherDir",
				upgradeVersion:              "1.9.0",
				location:                    "southcentralus",
				timeoutInMinutes:            60,
				cordonDrainTimeoutInMinutes: 60,
			},
			expectedErr: errors.New("ambiguous, please specify only one of --api-model and --deployment-dir"),
			name:        "NeedsNonAmbiguous",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				apiModelPath:        "./not/used",
				deploymentDirectory: "",
				upgradeVersion:      "1.9.0",
				location:            "southcentralus",
				controlPlaneOnly:    true,
				agentPoolToUpgrade:  "linuxpool1",
			},
			expectedErr: errors.New("flags --control-plane-only and --node-pool are not allowed together"),
			name:        "MasterAndSinglePool",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				apiModelPath:        "./not/used",
				deploymentDirectory: "",
				upgradeVersion:      "1.9.0",
				location:            "southcentralus",
				agentPoolToUpgrade:  "linuxpool1",
			},
			expectedErr: nil,
			name:        "SinglePool",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				apiModelPath:        "./not/used",
				deploymentDirectory: "",
				upgradeVersion:      "1.9.0",
				location:            "southcentralus",
				controlPlaneOnly:    true,
			},
			expectedErr: nil,
			name:        "ControlPlaneOnly",
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				apiModelPath:        "./not/used",
				deploymentDirectory: "",
				upgradeVersion:      "1.9.0",
				location:            "southcentralus",
			},
			expectedErr: nil,
			name:        "IsValid",
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			err := c.uc.validate(r)
			if c.expectedErr != nil && err != nil {
				g.Expect(err.Error()).To(Equal(c.expectedErr.Error()))
			} else {
				g.Expect(err).To(BeNil())
				g.Expect(c.expectedErr).To(BeNil())
			}
		})
	}
}

func TestCreateUpgradeCommand(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	command := newUpgradeCmd()

	g.Expect(command.Use).Should(Equal(upgradeName))
	g.Expect(command.Short).Should(Equal(upgradeShortDescription))
	g.Expect(command.Long).Should(Equal(upgradeLongDescription))
	g.Expect(command.Flags().Lookup("location")).NotTo(BeNil())
	g.Expect(command.Flags().Lookup("resource-group")).NotTo(BeNil())
	g.Expect(command.Flags().Lookup("api-model")).NotTo(BeNil())
	g.Expect(command.Flags().Lookup("upgrade-version")).NotTo(BeNil())

	command.SetArgs([]string{})
	if err := command.Execute(); err == nil {
		t.Fatalf("expected an error when calling upgrade with no arguments")
	}
}

func TestUpgradeShouldFailForSameVersion(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": true,
	})
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:           "rg",
		apiModelPath:                "./not/used",
		upgradeVersion:              "1.10.13",
		location:                    "centralus",
		timeoutInMinutes:            60,
		cordonDrainTimeoutInMinutes: 60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initialize()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("upgrading from Kubernetes version 1.10.13 to version 1.10.13 is not supported"))
	resetValidVersions()
}

func TestUpgradeShouldFailForInvalidUpgradePath(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": false,
		"1.10.12": true,
	})
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:           "rg",
		apiModelPath:                "./not/used",
		upgradeVersion:              "1.10.13",
		location:                    "centralus",
		timeoutInMinutes:            60,
		cordonDrainTimeoutInMinutes: 60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.12", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initialize()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("upgrading from Kubernetes version 1.10.12 to version 1.10.13 is not supported"))
	resetValidVersions()
}
func TestUpgradeShouldSuceedForValidUpgradePath(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": true,
		"1.10.12": true,
	})
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:           "rg",
		apiModelPath:                "./not/used",
		upgradeVersion:              "1.10.13",
		location:                    "centralus",
		timeoutInMinutes:            60,
		cordonDrainTimeoutInMinutes: 60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.12", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initialize()
	g.Expect(err).NotTo(HaveOccurred())
	resetValidVersions()
}

func TestUpgradeFailWithPathWhenAzureDeployJsonIsInvalid(t *testing.T) {
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:           "rg",
		apiModelPath:                "./not/used",
		upgradeVersion:              "1.13.3",
		location:                    "centralus",
		timeoutInMinutes:            60,
		cordonDrainTimeoutInMinutes: 60,
		force:                       true,
		client:                      &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.13.2", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initialize()
	g.Expect(err).NotTo(HaveOccurred())
	resetValidVersions()
}
func TestUpgradeForceSameVersionShouldSucceed(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": false,
	})
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:           "rg",
		apiModelPath:                "./not/used",
		upgradeVersion:              "1.10.13",
		location:                    "centralus",
		timeoutInMinutes:            60,
		cordonDrainTimeoutInMinutes: 60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	upgradeCmd.force = true
	err := upgradeCmd.initialize()
	g.Expect(err).NotTo(HaveOccurred())
	resetValidVersions()
}

func TestUpgradeForceDowngradeShouldSetVersionOnContainerService(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.12": true,
		"1.10.13": true,
	})
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:           "rg",
		apiModelPath:                "./not/used",
		upgradeVersion:              "1.10.12",
		location:                    "centralus",
		timeoutInMinutes:            60,
		cordonDrainTimeoutInMinutes: 60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	upgradeCmd.force = true
	err := upgradeCmd.initialize()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(upgradeCmd.containerService.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal("1.10.12"))
	resetValidVersions()
}

func TestIsVMSSNameInAgentPoolsArray(t *testing.T) {
	cases := []struct {
		vmssName       string
		cs             *api.ContainerService
		poolsToUpgrade map[string]bool
		expected       bool
		name           string
	}{
		{
			vmssName: "k8s-agentpool1-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "agentpool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master":     true,
				"agentpool1": true,
			},
			expected: true,
			name:     "vmss is in the api model spec",
		},
		{
			vmssName: "my-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "agentpool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master":     true,
				"agentpool1": true,
			},
			expected: false,
			name:     "vmss unrecognized",
		},
		{
			vmssName: "k8s-frontendpool-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "frontendpool",
							Count:               30,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "backendpool",
							Count:               7,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "canary",
							Count:               5,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master":       true,
				"frontendpool": true,
				"backendpool":  true,
				"canary":       true,
			},
			expected: true,
			name:     "multiple pools, frontendpool vmss is in spec",
		},
		{
			vmssName: "k8s-backendpool-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "frontendpool",
							Count:               30,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "backendpool",
							Count:               7,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "canary",
							Count:               5,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master":       true,
				"frontendpool": true,
				"backendpool":  true,
				"canary":       true,
			},
			expected: true,
			name:     "multiple pools, backendpool vmss is in spec",
		},
		{
			vmssName: "k8s-canary-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "frontendpool",
							Count:               30,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "backendpool",
							Count:               7,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "canary",
							Count:               5,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master":       true,
				"frontendpool": true,
				"backendpool":  true,
				"canary":       true,
			},
			expected: true,
			name:     "multiple pools, canary vmss is in spec",
		},
		{
			vmssName: "k8s-canary-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master": true,
			},
			expected: false,
			name:     "no pools",
		},
		{
			vmssName: "k8s-canary-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "canary",
							Count:               1,
							AvailabilityProfile: api.AvailabilitySet,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master": true,
				"canary": true,
			},
			expected: false,
			name:     "availability set",
		},
		{
			vmssName: "k8s-frontendpool-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "frontendpool",
							Count:               30,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "backendpool",
							Count:               7,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "canary",
							Count:               5,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master":       false,
				"frontendpool": true,
				"backendpool":  false,
				"canary":       false,
			},
			expected: true,
			name:     "multiple pools, input vmss is the node pool to upgrade",
		},
		{
			vmssName: "k8s-frontendpool-41325566-vmss",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime: api.Docker,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "frontendpool",
							Count:               30,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "backendpool",
							Count:               7,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
						{
							Name:                "canary",
							Count:               5,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			poolsToUpgrade: map[string]bool{
				"master":       false,
				"frontendpool": false,
				"backendpool":  true,
				"canary":       false,
			},
			expected: false,
			name:     "multiple pools, input vmss is NOT the node pool to upgrade",
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ret := isVMSSNameInAgentPoolsArray(c.vmssName, c.cs, c.poolsToUpgrade)
			if ret != c.expected {
				t.Errorf("expected %t to be %t", ret, c.expected)
			}
		})
	}
}

func TestValidateSinglePoolTargetVersion(t *testing.T) {
	g := NewGomegaWithT(t)
	cases := []struct {
		uc            *upgradeCmd
		cs            *api.ContainerService
		validVersions map[string]bool
		expected      error
		name          string
	}{
		{
			uc: &upgradeCmd{
				upgradeVersion:     "1.15.4",
				agentPoolToUpgrade: "linuxpool1",
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						OrchestratorProfile: &api.OrchestratorProfile{
							OrchestratorType:    api.Kubernetes,
							OrchestratorVersion: "1.15.4",
						},
					},
				},
			},
			validVersions: map[string]bool{
				"1.15.4": true,
			},
			expected: nil,
			name:     "UpgradePoolToMasterVersion",
		},
		{
			uc: &upgradeCmd{
				upgradeVersion:     "1.15.3",
				agentPoolToUpgrade: "linuxpool1",
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						OrchestratorProfile: &api.OrchestratorProfile{
							OrchestratorType:    api.Kubernetes,
							OrchestratorVersion: "1.15.4",
						},
					},
				},
			},
			validVersions: map[string]bool{
				"1.15.3": true,
			},
			expected: nil,
			name:     "UpgradePoolToLowerMasterVersion",
		},
		{
			uc: &upgradeCmd{
				upgradeVersion:     "1.15.5",
				agentPoolToUpgrade: "linuxpool1",
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						OrchestratorProfile: &api.OrchestratorProfile{
							OrchestratorType:    api.Kubernetes,
							OrchestratorVersion: "1.15.4",
						},
					},
				},
			},
			validVersions: map[string]bool{
				"1.15.5": true,
			},
			expected: nil,
			name:     "UpgradePoolToHigherMasterVersion",
		},
		{
			uc: &upgradeCmd{
				upgradeVersion:     "1.15.5",
				agentPoolToUpgrade: "linuxpool1",
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						OrchestratorProfile: &api.OrchestratorProfile{
							OrchestratorType:    api.Kubernetes,
							OrchestratorVersion: "1.15.4",
						},
					},
				},
			},
			validVersions: map[string]bool{
				"1.15.5": false,
			},
			expected: errors.New("upgrading from Kubernetes version 1.15.4 to version 1.15.5 is not supported. To see a list of available upgrades, use 'aks-engine get-versions --version 1.15.4'"),
			name:     "UpgradePoolToUnsupportedVersion",
		},
	}
	for _, tc := range cases {
		c := tc
		setupValidVersions(c.validVersions)
		err := c.uc.validateTargetVersion()
		if c.expected != nil {
			g.Expect(c.expected.Error()).To(Equal(err.Error()))
		} else {
			g.Expect(c.expected).To(Not(HaveOccurred()))
		}
		resetValidVersions()
	}
}

func TestComputeAgentPoolsToUpgrade(t *testing.T) {
	cases := []struct {
		uc             *upgradeCmd
		cs             *api.ContainerService
		expectedResult map[string]bool
		expectedError  error
		name           string
	}{
		{
			uc: &upgradeCmd{
				agentPoolToUpgrade: "",
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						OrchestratorProfile: &api.OrchestratorProfile{
							OrchestratorType: api.Kubernetes,
						},
						AgentPoolProfiles: []*api.AgentPoolProfile{
							{
								Name: "linuxpool1",
							},
							{
								Name: "linuxpool2",
							},
						},
					},
				},
			},
			expectedResult: map[string]bool{
				"master":     true,
				"linuxpool1": true,
				"linuxpool2": true,
			},
			expectedError: nil,
			name:          "UpgradeAll",
		},
		{
			uc: &upgradeCmd{
				agentPoolToUpgrade: "linuxpool1",
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						OrchestratorProfile: &api.OrchestratorProfile{
							OrchestratorType: api.Kubernetes,
						},
						AgentPoolProfiles: []*api.AgentPoolProfile{
							{
								Name: "linuxpool1",
							},
							{
								Name: "linuxpool2",
							},
						},
					},
				},
			},
			expectedResult: map[string]bool{
				"master":     false,
				"linuxpool1": true,
				"linuxpool2": false,
			},
			expectedError: nil,
			name:          "UpgradeSinglePool",
		},
		{
			uc: &upgradeCmd{
				agentPoolToUpgrade: "bad-pool",
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						OrchestratorProfile: &api.OrchestratorProfile{
							OrchestratorType: api.Kubernetes,
						},
						AgentPoolProfiles: []*api.AgentPoolProfile{
							{
								Name: "linuxpool1",
							},
							{
								Name: "linuxpool2",
							},
						},
					},
				},
			},
			expectedResult: nil,
			expectedError:  errors.New("node pool bad-pool is not part of the cluster"),
			name:           "InvalidPoolName",
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			err := c.uc.computeAgentPoolsToUpgrade()
			actual := fmt.Sprint(c.uc.agentPoolsToUpgrade)
			expected := fmt.Sprint(c.expectedResult)
			if err != nil && err.Error() != c.expectedError.Error() {
				t.Errorf("expected '%s' to be '%s'", err.Error(), c.expectedError.Error())
			}
			if err == nil && actual != expected {
				t.Errorf("expected %s to be %s", expected, actual)
			}
		})
	}
}
