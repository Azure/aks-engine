//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package config

import (
	"testing"
)

func TestSetRandomRegion(t *testing.T) {
	cases := []struct {
		config   Config
		expected []string
	}{
		{
			config: Config{
				Regions: []string{"westcentralus", "southeastasia", "westus2", "westeurope"},
			},
			expected: []string{"westcentralus", "southeastasia", "westus2", "westeurope"},
		},
		{
			config: Config{
				Regions: []string{"eastus"},
			},
			expected: []string{"eastus"},
		},
		{
			config: Config{
				Regions: []string{},
			},
			expected: []string{"eastus", "uksouth", "southeastasia", "westus2", "westeurope"},
		},
		{
			config: Config{
				Regions: nil,
			},
			expected: []string{"eastus", "uksouth", "southeastasia", "westus2", "westeurope"},
		},
		{
			config: Config{
				Regions: []string{"antarctica", "northpole"},
			},
			expected: []string{"antarctica", "northpole"},
		},
	}

	for _, c := range cases {
		c.config.SetRandomRegion()
		success := false
		for _, l := range c.expected {
			if c.config.Location == l {
				success = true
				break
			}
		}
		if !success {
			t.Fatalf("unexpected location: %s", c.config.Location)
		}
	}

}
