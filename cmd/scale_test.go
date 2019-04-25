// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func TestNewScaleCmd(t *testing.T) {
	command := newScaleCmd()
	if command.Use != scaleName || command.Short != scaleShortDescription || command.Long != scaleLongDescription {
		t.Fatalf("scale command should have use %s equal %s, short %s equal %s and long %s equal to %s", command.Use, scaleName, command.Short, scaleShortDescription, command.Long, scaleLongDescription)
	}

	expectedFlags := []string{"location", "resource-group", "api-model", "new-node-count", "node-pool", "master-FQDN"}
	for _, f := range expectedFlags {
		if command.Flags().Lookup(f) == nil {
			t.Fatalf("scale command should have flag %s", f)
		}
	}

	command.SetArgs([]string{})
	if err := command.Execute(); err == nil {
		t.Fatalf("expected an error when calling scale with no arguments")
	}
}

func TestScaleCmdValidate(t *testing.T) {
	r := &cobra.Command{}

	cases := []struct {
		sc          *scaleCmd
		expectedErr error
	}{
		{
			sc: &scaleCmd{
				apiModelPath:         "./not/used",
				deploymentDirectory:  "",
				location:             "centralus",
				resourceGroupName:    "",
				agentPoolToScale:     "agentpool1",
				newDesiredAgentCount: 5,
				masterFQDN:           "test",
			},
			expectedErr: errors.New("--resource-group must be specified"),
		},
		{
			sc: &scaleCmd{
				apiModelPath:         "./not/used",
				deploymentDirectory:  "",
				location:             "",
				resourceGroupName:    "testRG",
				agentPoolToScale:     "agentpool1",
				newDesiredAgentCount: 5,
				masterFQDN:           "test",
			},
			expectedErr: errors.New("--location must be specified"),
		},
		{
			sc: &scaleCmd{
				apiModelPath:        "./not/used",
				deploymentDirectory: "",
				location:            "centralus",
				resourceGroupName:   "testRG",
				agentPoolToScale:    "agentpool1",
				masterFQDN:          "test",
			},
			expectedErr: errors.New("--new-node-count must be specified"),
		},
		{
			sc: &scaleCmd{
				apiModelPath:         "",
				deploymentDirectory:  "",
				location:             "centralus",
				resourceGroupName:    "testRG",
				agentPoolToScale:     "agentpool1",
				newDesiredAgentCount: 5,
				masterFQDN:           "test",
			},
			expectedErr: errors.New("--api-model must be specified"),
		},
		{
			sc: &scaleCmd{
				apiModelPath:         "some/long/path",
				deploymentDirectory:  "someDir",
				location:             "centralus",
				resourceGroupName:    "testRG",
				agentPoolToScale:     "agentpool1",
				newDesiredAgentCount: 5,
				masterFQDN:           "test",
			},
			expectedErr: errors.New("ambiguous, please specify only one of --api-model and --deployment-dir"),
		},
		{
			sc: &scaleCmd{
				apiModelPath:         "./not/used",
				deploymentDirectory:  "",
				location:             "centralus",
				resourceGroupName:    "testRG",
				agentPoolToScale:     "agentpool1",
				newDesiredAgentCount: 5,
				masterFQDN:           "test",
			},
			expectedErr: nil,
		},
	}

	for _, c := range cases {
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
	}
}
