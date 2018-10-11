package persistentvolume

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
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

// Get returns the current pvs for a given kubeconfig
func Get() (*List, error) {
	cmd := exec.Command("kubectl", "get", "pv", "-o", "json")
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
func WaitOnReady(pvCount int, sleep, duration time.Duration) bool {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for PVs to become Bound", duration.String())
			default:
				if AreAllReady(pvCount) {
					readyCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case <-errCh:
			return false
		case ready := <-readyCh:
			return ready
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
