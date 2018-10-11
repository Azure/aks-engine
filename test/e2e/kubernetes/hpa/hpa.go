package hpa

import (
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
)

// HPA represents a kubernetes HPA
type HPA struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
}

// Metadata holds information like name, namespace, and labels
type Metadata struct {
	CreatedAt time.Time `json:"creationTimestamp"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
}

// Spec holds information like clusterIP and port
type Spec struct {
	MinReplicas                    int `json:"minReplicas"`
	MaxReplicas                    int `json:"maxReplicas"`
	TargetCPUUtilizationPercentage int `json:"targetCPUUtilizationPercentage"`
}

// Status holds the load balancer definition
type Status struct {
	LoadBalancer LoadBalancer `json:"loadBalancer"`
}

// LoadBalancer holds the ingress definitions
type LoadBalancer struct {
	CurrentCPUUtilizationPercentage int `json:"currentCPUUtilizationPercentage"`
	CurrentReplicas                 int `json:"currentReplicas"`
	DesiredReplicas                 int `json:"desiredReplicas"`
}

// Get returns the HPA definition specified in a given namespace
func Get(name, namespace string) (*HPA, error) {
	cmd := exec.Command("kubectl", "get", "hpa", "-o", "json", "-n", namespace, name)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl get hpa':%s\n", string(out))
		return nil, err
	}
	h := HPA{}
	err = json.Unmarshal(out, &h)
	if err != nil {
		log.Printf("Error unmarshalling service json:%s\n", err)
		return nil, err
	}
	return &h, nil
}

// Delete will delete a HPA in a given namespace
func (h *HPA) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("kubectl", "delete", "hpa", "-n", h.Metadata.Namespace, h.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd)
		if kubectlError != nil {
			log.Printf("Error while trying to delete service %s in namespace %s:%s\n", h.Metadata.Namespace, h.Metadata.Name, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}
