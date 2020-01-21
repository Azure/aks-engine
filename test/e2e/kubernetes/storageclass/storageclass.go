//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package storageclass

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

// StorageClass is used to parse data from kubectl get storageclass
type StorageClass struct {
	Metadata             Metadata               `json:"metadata"`
	Parameters           Parameters             `json:"parameters"`
	Provisioner          string                 `json:"provisioner"`
	VolumeBindingMode    string                 `json:"volumeBindingMode"`
	AllowVolumeExpansion bool                   `json:"allowVolumeExpansion,omitempty"`
	AllowedTopologies    []TopologySelectorTerm `json:"allowedTopologies,omitempty"`
}

// TopologySelectorTerm is a list of topology selector requirements by labels
type TopologySelectorTerm struct {
	MatchLabelExpressions []TopologySelectorLabelRequirement `json:"matchLabelExpressions,omitempty"`
}

// TopologySelectorLabelRequirement holds information about a topology selector
type TopologySelectorLabelRequirement struct {
	Key    string   `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}

// Metadata holds information like name, create time
type Metadata struct {
	CreatedAt time.Time `json:"creationTimestamp"`
	Name      string    `json:"name"`
}

// Parameters holds information like skuName
type Parameters struct {
	SkuName string `json:"skuName"`
}

// CreateStorageClassFromFile will create a StorageClass from file with a name
func CreateStorageClassFromFile(filename, name string) (*StorageClass, error) {
	cmd := exec.Command("k", "apply", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create StorageClass %s:%s\n", name, string(out))
		return nil, err
	}
	sc, err := Get(name)
	if err != nil {
		log.Printf("Error while trying to fetch StorageClass %s:%s\n", name, err)
		return nil, err
	}
	return sc, nil
}

// GetResult is a return struct for GetAsync
type GetResult struct {
	storageClass *StorageClass
	err          error
}

// GetAsync wraps Get with a struct response for goroutine + channel usage
func GetAsync(scName string) GetResult {
	storageClass, err := Get(scName)
	return GetResult{
		storageClass: storageClass,
		err:          err,
	}
}

// Get will return a StorageClass with a given name and namespace
func Get(scName string) (*StorageClass, error) {
	cmd := exec.Command("k", "get", "storageclass", scName, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	sc := StorageClass{}
	err = json.Unmarshal(out, &sc)
	if err != nil {
		log.Printf("Error unmarshalling StorageClass json:%s\n", err)
		return nil, err
	}
	return &sc, nil
}

// Describe will describe a storageclass resource
func (sc *StorageClass) Describe() error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "describe", "storageclass", sc.Metadata.Name)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	log.Printf("\n%s\n", string(out))
	return err
}

// WaitOnReady will block until StorageClass is available
func (sc *StorageClass) WaitOnReady(sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var mostRecentWaitOnReadyError error
	ch := make(chan GetResult)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAsync(sc.Metadata.Name)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitOnReadyError = result.err
			storageClass := result.storageClass
			if mostRecentWaitOnReadyError == nil {
				if storageClass != nil {
					return true, nil
				}
			}
		case <-ctx.Done():
			err := sc.Describe()
			if err != nil {
				log.Printf("Unable to describe storageclass\n: %s", err)
			}
			return false, errors.Errorf("WaitOnReady timed out: %s\n", mostRecentWaitOnReadyError)
		}
	}
}
