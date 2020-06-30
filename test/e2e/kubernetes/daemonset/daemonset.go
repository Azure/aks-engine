//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package daemonset

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

// Daemonset is used to parse data from kubectl get daemonsets
type Daemonset struct {
	Metadata Metadata `json:"metadata"`
	Status   Status   `json:"status"`
}

// Metadata holds information like name, createdat, labels, and namespace
type Metadata struct {
	CreatedAt time.Time         `json:"creationTimestamp"`
	Labels    map[string]string `json:"labels"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
}

// Status holds information like hostIP and phase
type Status struct {
	CurrentNumberScheduled int `json:"currentNumberScheduled"`
	DesiredNumberScheduled int `json:"desiredNumberScheduled"`
	NumberAvailable        int `json:"numberAvailable"`
	NumberReady            int `json:"numberReady"`
}

// List is a container that holds all pods returned from doing a kubectl get daemonsets
type List struct {
	Daemonsets []Daemonset `json:"items"`
}

// GetResult is a return struct for GetAsync
type GetResult struct {
	ds  *Daemonset
	err error
}

// CreateDaemonsetFromFile will create a Pod from file with a name
func CreateDaemonsetFromFile(filename, name, namespace string, sleep, timeout time.Duration) (*Daemonset, error) {
	cmd := exec.Command("k", "apply", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create Daemonset %s:%s\n", name, string(out))
		return nil, err
	}
	d, err := GetWithRetry(name, namespace, sleep, timeout)
	if err != nil {
		log.Printf("Error while trying to fetch Pod %s:%s\n", name, err)
		return nil, err
	}
	return d, nil
}

// Get will return a daemonset with a given name and namespace
func Get(dsName, namespace string, retries int) (*Daemonset, error) {
	ds := Daemonset{}
	var out []byte
	var err error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "get", "daemonsets", dsName, "-n", namespace, "-o", "json")
		out, err = cmd.CombinedOutput()
		if err == nil {
			jsonErr := json.Unmarshal(out, &ds)
			if jsonErr != nil {
				log.Printf("Error unmarshalling pods json:%s\n", jsonErr)
				err = jsonErr
			}
		}
		time.Sleep(3 * time.Second)
	}
	return &ds, err
}

// GetAsync wraps Get with a struct response for goroutine + channel usage
func GetAsync(dsName, namespace string) GetResult {
	ds, err := Get(dsName, namespace, 1)
	return GetResult{
		ds:  ds,
		err: err,
	}
}

// GetWithRetry gets a daemonset, allowing for retries
func GetWithRetry(dsName, namespace string, sleep, timeout time.Duration) (*Daemonset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentGetWithRetryError error
	var ds *Daemonset
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAsync(dsName, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetWithRetryError = result.err
			ds = result.ds
			if mostRecentGetWithRetryError == nil {
				if ds != nil {
					return ds, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetWithRetry timed out: %s\n", mostRecentGetWithRetryError)
		}
	}
}
