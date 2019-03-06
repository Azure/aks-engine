// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"strings"
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
	}{
		{
			uc: &upgradeCmd{
				resourceGroupName:   "",
				deploymentDirectory: "_output/test",
				upgradeVersion:      "1.8.9",
				location:            "centralus",
				timeoutInMinutes:    60,
			},
			expectedErr: errors.New("--resource-group must be specified"),
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				deploymentDirectory: "_output/test",
				upgradeVersion:      "1.8.9",
				location:            "",
				timeoutInMinutes:    60,
			},
			expectedErr: errors.New("--location must be specified"),
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				deploymentDirectory: "_output/test",
				upgradeVersion:      "",
				location:            "southcentralus",
				timeoutInMinutes:    60,
			},
			expectedErr: errors.New("--upgrade-version must be specified"),
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				deploymentDirectory: "",
				upgradeVersion:      "1.9.0",
				location:            "southcentralus",
				timeoutInMinutes:    60,
			},
			expectedErr: errors.New("--deployment-dir must be specified"),
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				deploymentDirectory: "",
				upgradeVersion:      "1.9.0",
				location:            "southcentralus",
				timeoutInMinutes:    60,
			},
			expectedErr: errors.New("--deployment-dir must be specified"),
		},
		{
			uc: &upgradeCmd{
				resourceGroupName:   "test",
				deploymentDirectory: "_output/mydir",
				upgradeVersion:      "1.9.0",
				location:            "southcentralus",
			},
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		err := c.uc.validate(r)
		if c.expectedErr != nil && err != nil {
			g.Expect(err.Error()).To(Equal(c.expectedErr.Error()))
		} else {
			g.Expect(err).To(BeNil())
			g.Expect(c.expectedErr).To(BeNil())
		}
	}
}

func TestCreateUpgradeCommandSuccesfully(t *testing.T) {
	g := NewGomegaWithT(t)
	output := newUpgradeCmd()

	g.Expect(output.Use).Should(Equal(upgradeName))
	g.Expect(output.Short).Should(Equal(upgradeShortDescription))
	g.Expect(output.Long).Should(Equal(upgradeLongDescription))
	g.Expect(output.Flags().Lookup("location")).NotTo(BeNil())
	g.Expect(output.Flags().Lookup("resource-group")).NotTo(BeNil())
	g.Expect(output.Flags().Lookup("deployment-dir")).NotTo(BeNil())
	g.Expect(output.Flags().Lookup("upgrade-version")).NotTo(BeNil())
}

func TestUpgradeShouldFailForSameVersion(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": true,
	})
	fakeARMTemplateHandle := strings.NewReader(`{"parameters" : { "nameSuffix" : {"defaultValue" : "test"}}}`)
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:   "rg",
		deploymentDirectory: "_output/test",
		upgradeVersion:      "1.10.13",
		location:            "centralus",
		timeoutInMinutes:    60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initializeFromLocalState(fakeARMTemplateHandle)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("upgrading from Kubernetes version 1.10.13 to version 1.10.13 is not supported"))
	resetValidVersions()
}

func TestUpgradeShouldFailForInvalidUpgradePath(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": false,
		"1.10.12": true,
	})
	fakeARMTemplateHandle := strings.NewReader(`{"parameters" : { "nameSuffix" : {"defaultValue" : "test"}}}`)
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:   "rg",
		deploymentDirectory: "_output/test",
		upgradeVersion:      "1.10.13",
		location:            "centralus",
		timeoutInMinutes:    60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.12", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initializeFromLocalState(fakeARMTemplateHandle)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("upgrading from Kubernetes version 1.10.12 to version 1.10.13 is not supported"))
	resetValidVersions()
}
func TestUpgradeShouldSuceedForValidUpgradePath(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": true,
		"1.10.12": true,
	})
	fakeARMTemplateHandle := strings.NewReader(`{"parameters" : { "nameSuffix" : {"defaultValue" : "test"}}}`)
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:   "rg",
		deploymentDirectory: "_output/test",
		upgradeVersion:      "1.10.13",
		location:            "centralus",
		timeoutInMinutes:    60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.12", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initializeFromLocalState(fakeARMTemplateHandle)
	g.Expect(err).NotTo(HaveOccurred())
	resetValidVersions()
}

func TestUpgradeFailWithPathWhenAzureDeployJsonIsInvalid(t *testing.T) {
	fakeARMTemplateHandle := strings.NewReader(`{"parameters" : { "nameSuffix"}}`)
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:   "rg",
		deploymentDirectory: "_output/myspecialtestfolder",
		upgradeVersion:      "1.13.3",
		location:            "centralus",
		timeoutInMinutes:    60,
		force:               true,
		client:              &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.13.2", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	err := upgradeCmd.initializeFromLocalState(fakeARMTemplateHandle)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("_output/myspecialtestfolder/azuredeploy.json"))
}
func TestUpgradeForceSameVersionShouldSucceed(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.13": false,
	})
	fakeARMTemplateHandle := strings.NewReader(`{"parameters" : { "nameSuffix" : {"defaultValue" : "test"}}}`)
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:   "rg",
		deploymentDirectory: "_output/test",
		upgradeVersion:      "1.10.13",
		location:            "centralus",
		timeoutInMinutes:    60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	upgradeCmd.force = true
	err := upgradeCmd.initializeFromLocalState(fakeARMTemplateHandle)
	g.Expect(err).NotTo(HaveOccurred())
	resetValidVersions()
}

func TestUpgradeForceDowngradeShouldSetVersionOnContainerService(t *testing.T) {
	setupValidVersions(map[string]bool{
		"1.10.12": true,
		"1.10.13": true,
	})
	fakeARMTemplateHandle := strings.NewReader(`{"parameters" : { "nameSuffix" : {"defaultValue" : "test"}}}`)
	g := NewGomegaWithT(t)
	upgradeCmd := &upgradeCmd{
		resourceGroupName:   "rg",
		deploymentDirectory: "_output/test",
		upgradeVersion:      "1.10.12",
		location:            "centralus",
		timeoutInMinutes:    60,

		client: &armhelpers.MockAKSEngineClient{},
	}

	containerServiceMock := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	containerServiceMock.Location = "centralus"
	upgradeCmd.containerService = containerServiceMock
	upgradeCmd.force = true
	err := upgradeCmd.initializeFromLocalState(fakeARMTemplateHandle)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(upgradeCmd.containerService.Properties.OrchestratorProfile.OrchestratorVersion).To(Equal("1.10.12"))
	resetValidVersions()
}
