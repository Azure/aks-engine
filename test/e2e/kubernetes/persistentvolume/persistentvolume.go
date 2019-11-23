//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package persistentvolume

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
)

// PersistentVolume is used to parse data from kubectl get pv
type PersistentVolume struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
}

// Metadata holds information like name, create time, and namespace
type Metadata struct {
	CreatedAt time.Time         `json:"creationTimestamp"`
	Labels    map[string]string `json:"labels"`
	Name      string            `json:"name"`
}

// Spec holds information like storageClassName, nodeAffinity
type Spec struct {
	StorageClassName string       `json:"storageClassName"`
	NodeAffinity     NodeAffinity `json:"nodeAffinity"`
}

// NodeAffinity holds information like required nodeselector
type NodeAffinity struct {
	Required *NodeSelector `json:"required"`
}

// NodeSelector represents the union of the results of one or more label queries
type NodeSelector struct {
	//Required. A list of node selector terms. The terms are ORed.
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// NodeSelectorTerm represents node selector requirements
type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchFields      []NodeSelectorRequirement `json:"matchFields,omitempty"`
}

// NodeSelectorRequirement is a selector that contains values, a key, and an operator
type NodeSelectorRequirement struct {
	Key    string   `json:"key"`
	Values []string `json:"values,omitempty"`
}

// Status holds information like phase
type Status struct {
	Phase string `json:"phase"`
}

// List is used to parse out PersistentVolume from a list
type List struct {
	PersistentVolumes []PersistentVolume `json:"items"`
}

// DescribePVs describes all persistent volume resources
func DescribePVs() {
	list, err := Get()
	if err != nil {
		log.Printf("Unable to get pvs: %s", err)
	}
	if list != nil {
		for _, pv := range list.PersistentVolumes {
			err := pv.Describe()
			if err != nil {
				log.Printf("Unable to describe pv %s: %s", pv.Metadata.Name, err)
			}
		}
	}
}

// Describe will describe a pv resource
func (pv *PersistentVolume) Describe() error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "describe", "pv", pv.Metadata.Name)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	log.Printf("\n%s\n", string(out))
	return err
}

// Get returns the current pvs for a given kubeconfig
func Get() (*List, error) {
	cmd := exec.Command("k", "get", "pv", "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl get pv':%s", string(out))
		return nil, err
	}
	pvl := List{}
	err = json.Unmarshal(out, &pvl)
	if err != nil {
		log.Printf("Error unmarshalling pvs json:%s", err)
	}
	return &pvl, nil
}

// WaitOnReady will block until all pvs are in ready state
func WaitOnReady(pvCount int, sleep, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- AreAllReady(pvCount)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case ready := <-ch:
			if ready {
				return ready
			}
		case <-ctx.Done():
			DescribePVs()
			return false
		}
	}
}

// AreAllReady returns a bool depending on cluster state
func AreAllReady(pvCount int) bool {
	list, _ := Get()
	if list != nil && len(list.PersistentVolumes) == pvCount {
		for _, pv := range list.PersistentVolumes {
			if pv.Status.Phase == "Bound" {
				return true
			}
		}
	}
	return false
}
