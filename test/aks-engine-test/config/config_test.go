// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package config

import "testing"

func TestConfigParse(t *testing.T) {

	testCfg := `
{"deployments":
  [
    {
      "cluster_definition":"examples/kubernetes.json",
      "location":"westus",
      "skip_validation":true
    },
    {
      "cluster_definition":"examples/dcos.json",
      "location":"eastus",
      "skip_validation":false
    },
    {
      "cluster_definition":"examples/swarm.json",
      "location":"southcentralus"
    },
    {
      "cluster_definition":"examples/swarmmode.json",
      "location":"westus2"
    }
  ]
}
`

	testConfig := TestConfig{}
	if err := testConfig.Read([]byte(testCfg)); err != nil {
		t.Fatal(err)
	}
	if err := testConfig.validate(); err != nil {
		t.Fatal(err)
	}
	if len(testConfig.Deployments) != 4 {
		t.Fatalf("Wrong number of deployments: %d instead of 4", len(testConfig.Deployments))
	}
}
