// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const commandTimeout = 1 * time.Minute

// List holds a list of services returned from kubectl get svc
type List struct {
	Services []Service `json:"items"`
}

// Service represents a kubernetes service
type Service struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
}

// Metadata holds information like name, namespace, and labels
type Metadata struct {
	CreatedAt time.Time         `json:"creationTimestamp"`
	Labels    map[string]string `json:"labels"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
}

// Spec holds information like clusterIP and port
type Spec struct {
	ClusterIP string `json:"clusterIP"`
	Ports     []Port `json:"ports"`
	Type      string `json:"type"`
}

// Port represents a service port definition
type Port struct {
	NodePort   int    `json:"nodePort"`
	Port       int    `json:"port"`
	Protocol   string `json:"protocol"`
	TargetPort int    `json:"targetPort"`
}

// Status holds the load balancer definition
type Status struct {
	LoadBalancer LoadBalancer `json:"loadBalancer"`
}

// LoadBalancer holds the ingress definitions
type LoadBalancer struct {
	Ingress []map[string]string `json:"ingress"`
}

// GetResult is a return struct for GetAsync
type GetResult struct {
	svc *Service
	err error
}

// GetAsync wraps Get with a struct response for goroutine + channel usage
func GetAsync(name, namespace string) GetResult {
	svc, err := Get(name, namespace)
	return GetResult{
		svc: svc,
		err: err,
	}
}

// Get returns the service definition specified in a given namespace
func Get(name, namespace string) (*Service, error) {
	cmd := exec.Command("k", "get", "svc", "-o", "json", "-n", namespace, name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error getting svc:\n")
		util.PrintCommand(cmd)
		return nil, err
	}
	s := Service{}
	err = json.Unmarshal(out, &s)
	if err != nil {
		log.Printf("Error unmarshalling service json:%s\n", err)
		return nil, err
	}
	return &s, nil
}

// GetAll will return all services in a given namespace
func GetAll(namespace string) (*List, error) {
	cmd := exec.Command("k", "get", "svc", "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error getting all services:\n")
		return nil, err
	}
	sl := List{}
	err = json.Unmarshal(out, &sl)
	if err != nil {
		log.Printf("Error unmarshalling services json:%s\n", err)
		return nil, err
	}
	return &sl, nil
}

// GetAllByPrefixResult is a return struct for GetAllByPrefixAsync
type GetAllByPrefixResult struct {
	svcs []Service
	err  error
}

// GetAllByPrefixAsync wraps Get with a struct response for goroutine + channel usage
func GetAllByPrefixAsync(prefix, namespace string) GetAllByPrefixResult {
	svcs, err := GetAllByPrefix(prefix, namespace)
	return GetAllByPrefixResult{
		svcs: svcs,
		err:  err,
	}
}

// GetAllByPrefix will return all services in a given namespace that match a prefix
func GetAllByPrefix(prefix, namespace string) ([]Service, error) {
	sl, err := GetAll(namespace)
	if err != nil {
		return nil, err
	}
	services := []Service{}
	for _, s := range sl.Services {
		matched, err := regexp.MatchString(prefix+"-.*", s.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match service name:%s\n", err)
			return nil, err
		}
		if matched {
			services = append(services, s)
		}
	}
	return services, nil
}

// Delete will delete a service in a given namespace
func (s *Service) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "delete", "svc", "-n", s.Metadata.Namespace, s.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd, commandTimeout)
		if kubectlError != nil {
			log.Printf("Error while trying to delete service %s in namespace %s:%s\n", s.Metadata.Namespace, s.Metadata.Name, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// GetNodePort will return the node port for a given pod
func (s *Service) GetNodePort(port int) int {
	for _, p := range s.Spec.Ports {
		if p.Port == port {
			return p.NodePort
		}
	}
	return 0
}

// WaitForIngress waits for an Ingress to be provisioned
func (s *Service) WaitForIngress(timeout, sleep time.Duration) (*Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var mostRecentWaitForIngressError error
	ch := make(chan GetResult)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- GetAsync(s.Metadata.Name, s.Metadata.Namespace):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitForIngressError = result.err
			svc := result.svc
			if mostRecentWaitForIngressError == nil {
				if svc != nil && svc.Status.LoadBalancer.Ingress != nil {
					return svc, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("WaitForIngress timed out: %s\n", mostRecentWaitForIngressError)
		}
	}
}

// WaitOnDeleted returns when a service resource is successfully deleted
func WaitOnDeleted(servicePrefix, namespace string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetAllByPrefixResult)
	var mostRecentWaitOnDeletedError error
	var svcs []Service
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- GetAllByPrefixAsync(servicePrefix, namespace):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitOnDeletedError = result.err
			svcs = result.svcs
			if mostRecentWaitOnDeletedError == nil {
				if len(svcs) == 0 {
					return true, nil
				}
			}
		case <-ctx.Done():
			return false, errors.Errorf("WaitOnDeleted timed out: %s\n", mostRecentWaitOnDeletedError)
		}
	}
}

// Validate will attempt to run an http.Get against the root service url
func (s *Service) Validate(check string, attempts int, sleep, wait time.Duration) bool {
	var err error
	var url string
	var i int
	var resp *http.Response
	svc, waitErr := s.WaitForIngress(wait, 5*time.Second)
	if waitErr != nil {
		log.Printf("Unable to verify external IP, cannot validate service:%s\n", waitErr)
		return false
	}
	if svc.Status.LoadBalancer.Ingress == nil || len(svc.Status.LoadBalancer.Ingress) == 0 {
		log.Printf("Service LB ingress is empty or nil: %#v\n", svc.Status.LoadBalancer.Ingress)
		return false
	}
	for i = 1; i <= attempts; i++ {
		url = fmt.Sprintf("http://%s", svc.Status.LoadBalancer.Ingress[0]["ip"])
		resp, err = http.Get(url)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			matched, _ := regexp.MatchString(check, string(body))
			if matched {
				defer resp.Body.Close()
				return true
			}
			log.Printf("Got unexpected URL body, expected to find %s, got:\n%s\n", check, string(body))
		}
		time.Sleep(sleep)
	}
	log.Printf("Unable to validate URL %s after %s, err: %#v\n", url, time.Duration(i)*wait, err)
	if resp != nil {
		defer resp.Body.Close()
	}
	return false
}

// CreateServiceFromFile will create a Service from file with a name
func CreateServiceFromFile(filename, name, namespace string) (*Service, error) {
	cmd := exec.Command("k", "create", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create Service %s:%s\n", name, string(out))
		return nil, err
	}
	svc, err := Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch Service %s:%s\n", name, err)
		return nil, err
	}
	return svc, nil
}

// CreateServiceFromFileDeleteIfExist will create a Service from file, deleting any pre-existing service with the same name
func CreateServiceFromFileDeleteIfExist(filename, name, namespace string) (*Service, error) {
	s, _ := Get(name, namespace)
	if s != nil {
		err := s.Delete(util.DefaultDeleteRetries)
		if err != nil {
			return nil, err
		}
		_, err = WaitOnDeleted(name, namespace, 10*time.Second, 1*time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return CreateServiceFromFile(filename, name, namespace)
}
