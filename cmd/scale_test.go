// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/cli/config"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func TestNewScaleCmd(t *testing.T) {
	output := newScaleCmd()
	if output.Use != scaleName || output.Short != scaleShortDescription || output.Long != scaleLongDescription {
		t.Fatalf("scale command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, scaleName, output.Short, scaleShortDescription, output.Long, scaleLongDescription)
	}

	expectedFlags := []string{"location", "resource-group", "deployment-dir", "new-node-count", "node-pool", "master-fqdn"}
	for _, f := range expectedFlags {
		if output.Flags().Lookup(f) == nil {
			t.Fatalf("scale command should have flag %s", f)
		}
	}
}

func TestScaleCmdValidate(t *testing.T) {
	r := &cobra.Command{}

	cases := []struct {
		conf        config.ScaleConfig
		sc          *scaleCmd
		expectedErr error
	}{
		{
			conf: config.ScaleConfig{
				Location:      "centralus",
				ResourceGroup: "",
				DeploymentDir: "_output/test",
				NodePool:      "agentpool1",
				NewNodeCount:  5,
				MasterFQDN:    "test",
			},
			sc:          &scaleCmd{},
			expectedErr: errors.New("--resource-group must be specified"),
		},
		{
			conf: config.ScaleConfig{
				Location:      "",
				ResourceGroup: "testRG",
				DeploymentDir: "_output/test",
				NodePool:      "agentpool1",
				NewNodeCount:  5,
				MasterFQDN:    "test",
			},
			sc:          &scaleCmd{},
			expectedErr: errors.New("--location must be specified"),
		},
		{
			conf: config.ScaleConfig{
				Location:      "centralus",
				ResourceGroup: "testRG",
				DeploymentDir: "_output/test",
				NodePool:      "agentpool1",
				MasterFQDN:    "test",
			},
			sc:          &scaleCmd{},
			expectedErr: errors.New("--new-node-count must be specified"),
		},
		{
			conf: config.ScaleConfig{
				Location:      "centralus",
				ResourceGroup: "testRG",
				DeploymentDir: "",
				NodePool:      "agentpool1",
				NewNodeCount:  5,
				MasterFQDN:    "test",
			},
			sc:          &scaleCmd{},
			expectedErr: errors.New("--deployment-dir must be specified"),
		},
		{
			conf: config.ScaleConfig{
				Location:      "centralus",
				ResourceGroup: "testRG",
				DeploymentDir: "_output/test",
				NodePool:      "agentpool1",
				NewNodeCount:  5,
				MasterFQDN:    "test",
			},
			sc:          &scaleCmd{},
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		if err := mergo.Merge(&currentConfig.CLIConfig.Scale, c.conf); err != nil {
			t.Fatal(err)
		}
		err := c.sc.validate(r)
		if err != nil && c.expectedErr != nil {
			if err.Error() != c.expectedErr.Error() {
				t.Fatalf("expected validate scale command to return error %s, but instead got %s", c.expectedErr.Error(), err.Error())
			}
		} else {
			if c.expectedErr != nil {
				t.Fatalf("expected validate scale command to return error %s, but instead got no error", c.expectedErr.Error())
			} else if err != nil {
				t.Fatalf("expected validate scale command to return no error, but instead got %s", err.Error())
			}
		}
		// reset config
		currentConfig = config.Config{}
	}
}
