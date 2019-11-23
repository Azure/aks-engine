//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

// Config represents a kubernetes config object
type Config struct {
	Clusters []Cluster `json:"clusters"`
}

// Cluster contains the name and the cluster info
type Cluster struct {
	Name        string      `json:"name"`
	ClusterInfo ClusterInfo `json:"cluster"`
}

// ClusterInfo holds the server and cert
type ClusterInfo struct {
	Server string `json:"server"`
}

// GetConfigResult is the result type for GetConfigAsync
type GetConfigResult struct {
	Config *Config
	Err    error
}

// GetConfigAsync wraps GetConfig with a struct response for goroutine + channel usage
func GetConfigAsync() GetConfigResult {
	config, err := GetConfig()
	if config == nil {
		config = &Config{}
	}
	return GetConfigResult{
		Config: config,
		Err:    err,
	}
}

// GetConfig returns a Config value representing the current kubeconfig
func GetConfig() (*Config, error) {
	cmd := exec.Command("k", "config", "view", "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl config view':%s\n", err)
		return nil, err
	}
	c := Config{}
	err = json.Unmarshal(out, &c)
	if err != nil {
		log.Printf("Error unmarshalling config json:%s\n", err)
	}
	return &c, nil
}

// GetConfigWithRetry gets nodes, allowing for retries
func GetConfigWithRetry(sleep, timeout time.Duration) (*Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetConfigResult)
	var mostRecentGetConfigWithRetryError error
	var config *Config
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetConfigAsync()
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetConfigWithRetryError = result.Err
			config = result.Config
			if mostRecentGetConfigWithRetryError == nil {
				return config, nil
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetConfigWithRetry timed out: %s\n", mostRecentGetConfigWithRetryError)
		}
	}
}

// GetServerName returns the server for the given config in an sshable format
func (c *Config) GetServerName() string {
	s := c.Clusters[0].ClusterInfo.Server
	return strings.Split(s, "://")[1]
}
