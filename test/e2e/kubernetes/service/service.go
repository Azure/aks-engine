//+build test
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
	var zeroValueDuration time.Duration
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "delete", "svc", "-n", s.Metadata.Namespace, s.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd, zeroValueDuration)
		if kubectlError != nil {
			log.Printf("Error while trying to delete service %s in namespace %s:%s\n", s.Metadata.Namespace, s.Metadata.Name, kubectlError)
			log.Printf("%s\n", string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// DescribeServices describes all service resources whose name matches a substring
func DescribeServices(svcPrefix, namespace string) {
	svcs, err := GetAllByPrefix(svcPrefix, namespace)
	if err != nil {
		log.Printf("Unable to get services matching prefix %s in namespace %s: %s", svcPrefix, namespace, err)
	}
	for _, svc := range svcs {
		err := svc.Describe()
		if err != nil {
			log.Printf("Unable to describe service %s: %s", svc.Metadata.Name, err)
		}
	}
}

// Describe will describe a service resource
func (s *Service) Describe() error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "describe", "svc", s.Metadata.Name, "-n", s.Metadata.Namespace)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	log.Printf("\n%s\n", string(out))
	return err
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
func (s *Service) WaitForIngress(timeout, sleep time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var mostRecentWaitForIngressError error
	ch := make(chan GetResult)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAsync(s.Metadata.Name, s.Metadata.Namespace)
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
					s.Status.LoadBalancer.Ingress = svc.Status.LoadBalancer.Ingress
					return nil
				}
			}
		case <-ctx.Done():
			err := s.Describe()
			if err != nil {
				log.Printf("Unable to describe service\n: %s", err)
			}
			return errors.Errorf("WaitForIngress timed out: %s\n", mostRecentWaitForIngressError)
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
			default:
				ch <- GetAllByPrefixAsync(servicePrefix, namespace)
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
			DescribeServices(servicePrefix, namespace)
			return false, errors.Errorf("WaitOnDeleted timed out: %s\n", mostRecentWaitOnDeletedError)
		}
	}
}

// ValidateWithRetry waits for an Ingress to be provisioned
func (s *Service) ValidateWithRetry(bodyResponseTextMatch string, sleep, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var mostRecentValidateWithRetryError error
	ch := make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- s.Validate(bodyResponseTextMatch)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentValidateWithRetryError = result
			if mostRecentValidateWithRetryError == nil {
				return nil
			}
		case <-ctx.Done():
			err := s.Describe()
			if err != nil {
				log.Printf("Unable to describe service\n: %s", err)
			}
			return errors.Errorf("ValidateWithRetry timed out: %s\n", mostRecentValidateWithRetryError)
		}
	}
}

// Validate will attempt to run an http.Get against the root service url
func (s *Service) Validate(bodyResponseTextMatch string) error {
	if len(s.Status.LoadBalancer.Ingress) < 1 {
		return errors.Errorf("No LB ingress IP for service %s", s.Metadata.Name)
	}
	var resp *http.Response
	url := fmt.Sprintf("http://%s", s.Status.LoadBalancer.Ingress[0]["ip"])
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return errors.Errorf("Unable to call service at URL %s: %s", url, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Errorf("Unable to parse response body: %s", err)
	}
	matched, err := regexp.MatchString(bodyResponseTextMatch, string(body))
	if err != nil {
		return errors.Errorf("Unable to evalute response body against a regular expression match: %s", err)
	}
	if matched {
		return nil
	}
	return errors.Errorf("Got unexpected URL body, expected to find %s, got:\n%s\n", bodyResponseTextMatch, string(body))
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
