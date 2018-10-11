package node

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const (
	//ServerVersion is used to parse out the version of the API running
	ServerVersion = `(Server Version:\s)+(.*)`
)

// Node represents the kubernetes Node Resource
type Node struct {
	Status   Status   `json:"status"`
	Metadata Metadata `json:"metadata"`
}

// Metadata contains things like name and created at
type Metadata struct {
	Name        string            `json:"name"`
	CreatedAt   time.Time         `json:"creationTimestamp"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

// Status parses information from the status key
type Status struct {
	Info          Info        `json:"Info"`
	NodeAddresses []Address   `json:"addresses"`
	Conditions    []Condition `json:"conditions"`
}

// Address contains an address and a type
type Address struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

// Info contains information like what version the kubelet is running
type Info struct {
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	KubeletProxyVersion     string `json:"kubeletVersion"`
	OperatingSystem         string `json:"operatingSystem"`
}

// Condition contains various status information
type Condition struct {
	LastHeartbeatTime  time.Time `json:"lastHeartbeatTime"`
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	Message            string    `json:"message"`
	Reason             string    `json:"reason"`
	Status             string    `json:"status"`
	Type               string    `json:"type"`
}

// List is used to parse out Nodes from a list
type List struct {
	Nodes []Node `json:"items"`
}

// AreAllReady returns a bool depending on cluster state
func AreAllReady(nodeCount int) bool {
	list, _ := Get()
	if list != nil && len(list.Nodes) == nodeCount {
		for _, node := range list.Nodes {
			for _, condition := range node.Status.Conditions {
				if condition.Type == "KubeletReady" && condition.Status == "false" {
					return false
				}
			}
		}
		return true
	}
	return false
}

// WaitOnReady will block until all nodes are in ready state
func WaitOnReady(nodeCount int, sleep, duration time.Duration) bool {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Nodes to become ready", duration.String())
			default:
				if AreAllReady(nodeCount) {
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

// Get returns the current nodes for a given kubeconfig
func Get() (*List, error) {
	cmd := exec.Command("kubectl", "get", "nodes", "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl get nodes':%s", string(out))
		return nil, err
	}
	nl := List{}
	err = json.Unmarshal(out, &nl)
	if err != nil {
		log.Printf("Error unmarshalling nodes json:%s", err)
	}
	return &nl, nil
}

// Version get the version of the server
func Version() (string, error) {
	cmd := exec.Command("kubectl", "version", "--short")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl version':%s", string(out))
		return "", err
	}
	split := strings.Split(string(out), "\n")
	exp, err := regexp.Compile(ServerVersion)
	if err != nil {
		log.Printf("Error while compiling regexp:%s", ServerVersion)
	}
	s := exp.FindStringSubmatch(split[1])
	return s[2], nil
}

// GetAddressByType will return the Address object for a given Kubernetes node
func (ns *Status) GetAddressByType(t string) *Address {
	for _, a := range ns.NodeAddresses {
		if a.Type == t {
			return &a
		}
	}
	return nil
}

// GetByPrefix will return a []Node of all nodes that have a name that match the prefix
func GetByPrefix(prefix string) ([]Node, error) {
	list, err := Get()
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, 0)
	for _, n := range list.Nodes {
		exp, err := regexp.Compile(prefix)
		if err != nil {
			return nil, err
		}
		if exp.MatchString(n.Metadata.Name) {
			nodes = append(nodes, n)
		}
	}
	return nodes, nil
}
