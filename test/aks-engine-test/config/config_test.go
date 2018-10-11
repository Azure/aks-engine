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
	if len(testConfig.Deployments) != 1 {
		t.Fatalf("Wrong number of deployments: %d instead of 4", len(testConfig.Deployments))
	}
}
