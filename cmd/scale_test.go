// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
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
		name        string
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
			name:        "NoResourceGroup",
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
			name:        "NoLocation",
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
			name:        "NoNewNodeCount",
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
			name:        "NoAPIModel",
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
			name:        "Ambiguous",
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
					t.Fatalf("expected validate scale command to return error %s, but instead got %s", c.expectedErr.Error(), err.Error())
				}
			} else {
				if c.expectedErr != nil {
					t.Fatalf("expected validate scale command to return error %s, but instead got no error", c.expectedErr.Error())
				} else if err != nil {
					t.Fatalf("expected validate scale command to return no error, but instead got %s", err.Error())
				}
			}
		})
	}
}

func TestVmInVMASAgentPool(t *testing.T) {
	tags := map[string]*string{}

	cases := []struct {
		sc       *scaleCmd
		expected bool
		name     string
		vmName   string
	}{
		{
			sc: &scaleCmd{
				nameSuffix:       "39573225",
				agentPoolIndex:   0,
				agentPoolToScale: "linuxpool",
				agentPool: &api.AgentPoolProfile{
					Name:                "linuxpool2",
					OSType:              "Linux",
					AvailabilityProfile: "AvailabilitySet",
				},
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						ClusterID: "39573225",
					},
				},
			},
			expected: false,
			name:     "linux agentpool mismatch scale pool",
			vmName:   "k8s-linuxpool2-39573225-0",
		},
		{
			sc: &scaleCmd{
				nameSuffix:       "39573225",
				agentPoolIndex:   1,
				agentPoolToScale: "linuxpool2",
				agentPool: &api.AgentPoolProfile{
					Name:                "linuxpool2",
					OSType:              "Linux",
					AvailabilityProfile: "AvailabilitySet",
				},
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						ClusterID: "39573225",
					},
				},
			},
			expected: true,
			name:     "linux agentpool matches scale pool",
			vmName:   "k8s-linuxpool2-39573225-1",
		},
		{
			sc: &scaleCmd{
				nameSuffix:       "39573225",
				agentPoolIndex:   2,
				agentPoolToScale: "windowspool",
				agentPool: &api.AgentPoolProfile{
					Name:                "windowspool2",
					OSType:              "Windows",
					AvailabilityProfile: "AvailabilitySet",
				},
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						ClusterID: "39573225",
					},
				},
			},
			expected: false,
			name:     "windows agentpool mismatch scale pool",
			vmName:   "3957k8s030",
		},
		{
			sc: &scaleCmd{
				nameSuffix:       "39573225",
				agentPoolIndex:   3,
				agentPoolToScale: "windowspool2",
				agentPool: &api.AgentPoolProfile{
					Name:                "windowspool2",
					OSType:              "Windows",
					AvailabilityProfile: "AvailabilitySet",
				},
				containerService: &api.ContainerService{
					Properties: &api.Properties{
						ClusterID: "39573225",
					},
				},
			},
			expected: true,
			name:     "windows agentpool matches scale pool",
			vmName:   "3957k8s031",
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ret := c.sc.vmInVMASAgentPool(c.vmName, tags)
			if ret != c.expected {
				t.Errorf("expected %t to be %t", ret, c.expected)
			}
		})
	}
}
