package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Deployment represents an AKS cluster deployment on Azure
type Deployment struct {
	ClusterDefinition string `json:"cluster_definition"`
	Location          string `json:"location"`
	TestCategory      string `json:"category,omitempty"`
	SkipValidation    bool   `json:"skip_validation,omitempty"`
}

// TestConfig represents a cluster config
type TestConfig struct {
	Deployments []Deployment `json:"deployments"`
}

func (c *TestConfig) Read(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *TestConfig) validate() error {
	for _, d := range c.Deployments {
		if d.ClusterDefinition == "" {
			return errors.New("Cluster definition is not set")
		}
	}
	return nil
}

// GetTestConfig parses a cluster config
func GetTestConfig(fname string) (*TestConfig, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	config := &TestConfig{}
	if err = config.Read(data); err != nil {
		return nil, err
	}
	if err = config.validate(); err != nil {
		return nil, err
	}
	return config, nil
}
