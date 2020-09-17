// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func TestNewUpdateCmd(t *testing.T) {
	command := newUpdateCmd()
	if command.Use != updateName || command.Short != updateShortDescription || command.Long != updateLongDescription {
		t.Fatalf("update command should have use %s equal %s, short %s equal %s and long %s equal to %s", command.Use, updateName, command.Short, updateShortDescription, command.Long, updateLongDescription)
	}

	expectedFlags := []string{"location", "resource-group", "api-model", "node-pool"}
	for _, f := range expectedFlags {
		if command.Flags().Lookup(f) == nil {
			t.Fatalf("update command should have flag %s", f)
		}
	}

	command.SetArgs([]string{})
	if err := command.Execute(); err == nil {
		t.Fatalf("expected an error when calling update with no arguments")
	}
}

func TestUpdateCmdValidate(t *testing.T) {
	r := &cobra.Command{}

	cases := []struct {
		uc          *updateCmd
		expectedErr error
		name        string
	}{
		{
			uc: &updateCmd{
				apiModelPath:      "./not/used",
				location:          "centralus",
				resourceGroupName: "",
				agentPoolToUpdate: "agentpool1",
			},
			expectedErr: errors.New("--resource-group must be specified"),
			name:        "NoResourceGroup",
		},
		{
			uc: &updateCmd{
				apiModelPath:      "./not/used",
				location:          "",
				resourceGroupName: "rgname",
				agentPoolToUpdate: "agentpool1",
			},
			expectedErr: errors.New("--location must be specified"),
			name:        "NoLocation",
		},
		{
			uc: &updateCmd{
				apiModelPath:      "",
				location:          "centralus",
				resourceGroupName: "rgname",
				agentPoolToUpdate: "agentpool1",
			},
			expectedErr: errors.New("--api-model must be specified"),
			name:        "NoApiModel",
		},
		{
			uc: &updateCmd{
				apiModelPath:      "./not/used",
				location:          "centralus",
				resourceGroupName: "rgname",
				agentPoolToUpdate: "",
			},
			expectedErr: errors.New("--node-pool must be specified"),
			name:        "NoNodePool",
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			err := c.uc.validate(r)
			if err != nil && c.expectedErr != nil {
				if err.Error() != c.expectedErr.Error() {
					t.Fatalf("expected validate update command to return error %s, but instead got %s", c.expectedErr.Error(), err.Error())
				}
			} else {
				if c.expectedErr != nil {
					t.Fatalf("expected validate update command to return error %s, but instead got no error", c.expectedErr.Error())
				} else if err != nil {
					t.Fatalf("expected validate update command to return no error, but instead got %s", err.Error())
				}
			}
		})
	}
}
