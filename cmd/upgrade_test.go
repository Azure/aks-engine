// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"github.com/Azure/aks-engine/pkg/cli/config"
	"github.com/imdario/mergo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var _ = Describe("the upgrade command", func() {

	It("should create an upgrade command", func() {
		output := newUpgradeCmd()

		Expect(output.Use).Should(Equal(upgradeName))
		Expect(output.Short).Should(Equal(upgradeShortDescription))
		Expect(output.Long).Should(Equal(upgradeLongDescription))
		Expect(output.Flags().Lookup("location")).NotTo(BeNil())
		Expect(output.Flags().Lookup("resource-group")).NotTo(BeNil())
		Expect(output.Flags().Lookup("deployment-dir")).NotTo(BeNil())
		Expect(output.Flags().Lookup("upgrade-version")).NotTo(BeNil())
	})

	It("should validate an upgrade command", func() {
		r := &cobra.Command{}

		cases := []struct {
			conf        config.UpgradeConfig
			uc          *upgradeCmd
			expectedErr error
		}{
			{
				conf: config.UpgradeConfig{
					ResourceGroup:  "",
					DeploymentDir:  "_output/test",
					UpgradeVersion: "1.8.9",
					Location:       "centralus",
					VMTimeout:      60,
				},
				uc:          &upgradeCmd{},
				expectedErr: errors.New("--resource-group must be specified"),
			},
			{
				conf: config.UpgradeConfig{
					ResourceGroup:  "test",
					DeploymentDir:  "_output/test",
					UpgradeVersion: "1.8.9",
					Location:       "",
					VMTimeout:      60,
				},
				uc:          &upgradeCmd{},
				expectedErr: errors.New("--location must be specified"),
			},
			{
				conf: config.UpgradeConfig{
					ResourceGroup:  "test",
					DeploymentDir:  "_output/test",
					UpgradeVersion: "",
					Location:       "centralus",
					VMTimeout:      60,
				},
				uc:          &upgradeCmd{},
				expectedErr: errors.New("--upgrade-version must be specified"),
			},
			{
				conf: config.UpgradeConfig{
					ResourceGroup:  "test",
					DeploymentDir:  "",
					UpgradeVersion: "1.8.9",
					Location:       "centralus",
					VMTimeout:      60,
				},
				uc:          &upgradeCmd{},
				expectedErr: errors.New("--deployment-dir must be specified"),
			},
			{
				conf: config.UpgradeConfig{
					ResourceGroup:  "test",
					DeploymentDir:  "",
					UpgradeVersion: "1.9.0",
					Location:       "centralus",
					VMTimeout:      60,
				},
				uc:          &upgradeCmd{},
				expectedErr: errors.New("--deployment-dir must be specified"),
			},
			{
				conf: config.UpgradeConfig{
					ResourceGroup:  "test",
					DeploymentDir:  "_output/mydir",
					UpgradeVersion: "1.9.0",
					Location:       "southcentralus",
				},
				uc:          &upgradeCmd{},
				expectedErr: nil,
			},
		}

		for _, c := range cases {
			Expect(mergo.Merge(&currentConfig.CLIConfig.Upgrade, c.conf)).To(BeNil())
			err := c.uc.validate(r)
			if c.expectedErr != nil && err != nil {
				Expect(err.Error()).To(Equal(c.expectedErr.Error()))
			} else {
				Expect(err).To(BeNil())
				Expect(c.expectedErr).To(BeNil())
			}
			// reset config
			currentConfig = config.Config{}
		}

	})

})
