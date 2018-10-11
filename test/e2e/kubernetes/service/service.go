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

// Get returns the service definition specified in a given namespace
func Get(name, namespace string) (*Service, error) {
	cmd := exec.Command("kubectl", "get", "svc", "-o", "json", "-n", namespace, name)
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

// Delete will delete a service in a given namespace
func (s *Service) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("kubectl", "delete", "svc", "-n", s.Metadata.Namespace, s.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd)
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

// WaitForExternalIP waits for an external ip to be provisioned
func (s *Service) WaitForExternalIP(wait, sleep time.Duration) (*Service, error) {
	svcCh := make(chan *Service)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.New("Timeout exceeded while waiting for External IP to be provisioned")
			default:
				svc, _ := Get(s.Metadata.Name, s.Metadata.Namespace)
				if svc != nil && svc.Status.LoadBalancer.Ingress != nil {
					svcCh <- svc
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return nil, err
		case svc := <-svcCh:
			return svc, nil
		}
	}
}

// Validate will attempt to run an http.Get against the root service url
func (s *Service) Validate(check string, attempts int, sleep, wait time.Duration) bool {
	var err error
	var url string
	var i int
	var resp *http.Response
	svc, waitErr := s.WaitForExternalIP(wait, 5*time.Second)
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
	svc, err := Get(name, namespace)
	if err == nil {
		log.Printf("Service %s already exists\n", name)
		return svc, nil
	}
	cmd := exec.Command("kubectl", "create", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create Service %s:%s\n", name, string(out))
		return nil, err
	}
	svc, err = Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch Service %s:%s\n", name, err)
		return nil, err
	}
	return svc, nil
}
