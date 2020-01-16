// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func TestNewAddPoolCmd(t *testing.T) {
	command := newAddPoolCmd()
	if command.Use != addPoolName || command.Short != addPoolShortDescription || command.Long != addPoolLongDescription {
		t.Fatalf("addpool command should have use %s equal %s, short %s equal %s and long %s equal to %s", command.Use, addPoolName, command.Short, addPoolShortDescription, command.Long, addPoolLongDescription)
	}

	expectedFlags := []string{"location", "resource-group", "api-model", "node-pool"}
	for _, f := range expectedFlags {
		if command.Flags().Lookup(f) == nil {
			t.Fatalf("addpool command should have flag %s", f)
		}
	}

	command.SetArgs([]string{})
	if err := command.Execute(); err == nil {
		t.Fatalf("expected an error when calling addpool with no arguments")
	}
}

func TestAddPoolCmdValidate(t *testing.T) {
	r := &cobra.Command{}

	cases := []struct {
		sc          *addPoolCmd
		expectedErr error
		name        string
	}{
		{
			sc: &addPoolCmd{
				apiModelPath:      "./not/used",
				nodePoolPath:      "",
				location:          "centralus",
				resourceGroupName: "",
			},
			expectedErr: errors.New("--resource-group must be specified"),
			name:        "NoResourceGroup",
		},
		{
			sc: &addPoolCmd{
				apiModelPath:      "./not/used",
				nodePoolPath:      "",
				location:          "",
				resourceGroupName: "testRG",
			},
			expectedErr: errors.New("--location must be specified"),
			name:        "NoLocation",
		},
		{
			sc: &addPoolCmd{
				apiModelPath:      "./not/used",
				location:          "centralus",
				resourceGroupName: "testRG",
			},
			expectedErr: errors.New("--nodepool must be specified"),
			name:        "NoNodePool",
		},
		{
			sc: &addPoolCmd{
				apiModelPath:      "",
				nodePoolPath:      "",
				location:          "centralus",
				resourceGroupName: "testRG",
			},
			expectedErr: errors.New("--api-model must be specified"),
			name:        "NoAPIModel",
		},
		{
			sc: &addPoolCmd{
				apiModelPath:      "./not/used",
				nodePoolPath:      "./some/path",
				location:          "centralus",
				resourceGroupName: "testRG",
			},
			expectedErr: nil,
			name:        "IsValid",
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			err := c.sc.validate(r)
			if err != nil && c.expectedErr != nil {
				if err.Error() != c.expectedErr.Error() {
					t.Fatalf("expected validate addpool command to return error %s, but instead got %s", c.expectedErr.Error(), err.Error())
				}
			} else {
				if c.expectedErr != nil {
					t.Fatalf("expected validate addpool command to return error %s, but instead got no error", c.expectedErr.Error())
				} else if err != nil {
					t.Fatalf("expected validate addpool command to return no error, but instead got %s", err.Error())
				}
			}
		})
	}
}
