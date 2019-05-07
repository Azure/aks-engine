// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package hpa

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"regexp"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const commandTimeout = 1 * time.Minute

type List struct {
	HPAs []HPA `json:"items"`
}

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
	cmd := exec.Command("k", "get", "hpa", "-o", "json", "-n", namespace, name)
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

// GetAll will return all HPA resources in a given namespace
func GetAll(namespace string) (*List, error) {
	cmd := exec.Command("k", "get", "hpa", "-n", namespace, "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error getting hpa:\n")
		util.PrintCommand(cmd)
		return nil, err
	}
	hl := List{}
	err = json.Unmarshal(out, &hl)
	if err != nil {
		log.Printf("Error unmarshalling pods json:%s\n", err)
		return nil, err
	}
	return &hl, nil
}

// GetAllByPrefix will return all pods in a given namespace that match a prefix
func GetAllByPrefix(prefix, namespace string) ([]HPA, error) {
	hl, err := GetAll(namespace)
	if err != nil {
		return nil, err
	}
	hpas := []HPA{}
	for _, h := range hl.HPAs {
		matched, err := regexp.MatchString(prefix+"-.*", h.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return nil, err
		}
		if matched {
			hpas = append(hpas, h)
		}
	}
	return hpas, nil
}

// Delete will delete a HPA in a given namespace
func (h *HPA) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "delete", "hpa", "-n", h.Metadata.Namespace, h.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd, commandTimeout)
		if kubectlError != nil {
			log.Printf("Error while trying to delete service %s in namespace %s:%s\n", h.Metadata.Namespace, h.Metadata.Name, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// WaitOnDeleted returns when an hpa resource is successfully deleted
func WaitOnDeleted(hpaPrefix, namespace string, sleep, duration time.Duration) (bool, error) {
	succeededCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pods (%s) to be deleted in namespace (%s)", duration.String(), hpaPrefix, namespace)
			default:
				p, err := GetAllByPrefix(hpaPrefix, namespace)
				if err != nil {
					errCh <- errors.Errorf("Got error while getting Pods with prefix \"%s\" in namespace \"%s\"", hpaPrefix, namespace)
				}
				if len(p) == 0 {
					succeededCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case deleted := <-succeededCh:
			return deleted, nil
		}
	}
}
