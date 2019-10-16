// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package deployment

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/hpa"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/pod"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/service"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const (
	validateDeploymentNotExistRetries = 3
	deploymentGetAfterCreateTimeout   = 1 * time.Minute
)

// List holds a list of deployments returned from kubectl get deploy
type List struct {
	Deployments []Deployment `json:"items"`
}

// Deployment repesentes a kubernetes deployment
type Deployment struct {
	Metadata Metadata `json:"metadata"`
}

// Metadata holds information like labels, name, and namespace
type Metadata struct {
	CreatedAt time.Time         `json:"creationTimestamp"`
	Labels    map[string]string `json:"labels"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	HasHPA    bool              `json:"hasHPA"`
}

// Spec holds information the deployment strategy and number of replicas
type Spec struct {
	Replicas int      `json:"replicas"`
	Template Template `json:"template"`
}

// Template is used for fetching the deployment spec -> containers
type Template struct {
	TemplateSpec TemplateSpec `json:"spec"`
}

// TemplateSpec holds the list of containers for a deployment, the dns policy, and restart policy
type TemplateSpec struct {
	Containers    []Container `json:"containers"`
	DNSPolicy     string      `json:"dnsPolicy"`
	RestartPolicy string      `json:"restartPolicy"`
}

// Container holds information like image, pull policy, name, etc...
type Container struct {
	Image      string `json:"image"`
	PullPolicy string `json:"imagePullPolicy"`
	Name       string `json:"name"`
}

// CreateLinuxDeploy will create a deployment for a given image with a name in a namespace
// --overrides='{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}}}'
func CreateLinuxDeploy(image, name, namespace, miscOpts string) (*Deployment, error) {
	var commandTimeout time.Duration
	var cmd *exec.Cmd
	overrides := `{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}}}`
	if miscOpts != "" {
		cmd = exec.Command("k", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--overrides", overrides, miscOpts)
	} else {
		cmd = exec.Command("k", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--overrides", overrides)
	}
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	d, err := GetWithRetry(name, namespace, 3*time.Second, deploymentGetAfterCreateTimeout)
	if err != nil {
		log.Printf("Error while trying to fetch Deployment %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return d, nil
}

// CreateLinuxDeployIfNotExist first checks if a deployment already exists, and return it if so
// If not, we call CreateLinuxDeploy
func CreateLinuxDeployIfNotExist(image, name, namespace, miscOpts string) (*Deployment, error) {
	deployment, err := Get(name, namespace, validateDeploymentNotExistRetries)
	if err != nil {
		return CreateLinuxDeploy(image, name, namespace, miscOpts)
	}
	return deployment, nil
}

// CreateLinuxDeployDeleteIfExists will create a deployment, deleting any pre-existing deployment with the same name
func CreateLinuxDeployDeleteIfExists(pattern, image, name, namespace, miscOpts string) (*Deployment, error) {
	deployments, err := GetAllByPrefix(pattern, namespace)
	if err != nil {
		return nil, err
	}
	for _, d := range deployments {
		d.Delete(util.DefaultDeleteRetries)
	}
	return CreateLinuxDeploy(image, name, namespace, miscOpts)
}

// RunLinuxDeploy will create a deployment that runs a bash command in a pod
// --overrides=' "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}}}'
func RunLinuxDeploy(image, name, namespace, command string, replicas int) (*Deployment, error) {
	var commandTimeout time.Duration
	overrides := `{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}}}`
	cmd := exec.Command("k", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--replicas", strconv.Itoa(replicas), "--overrides", overrides, "--command", "--", "/bin/sh", "-c", command)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	d, err := GetWithRetry(name, namespace, 3*time.Second, deploymentGetAfterCreateTimeout)
	if err != nil {
		log.Printf("Error while trying to fetch Deployment %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return d, nil
}

// RunLinuxDeployDeleteIfExists will create a deployment that runs a bash command in a pod,
// deleting any pre-existing deployment with the same name
func RunLinuxDeployDeleteIfExists(pattern, image, name, namespace, command string, replicas int) (*Deployment, error) {
	deployments, err := GetAllByPrefix(pattern, namespace)
	if err != nil {
		return nil, err
	}
	for _, d := range deployments {
		d.Delete(util.DefaultDeleteRetries)
	}
	return RunLinuxDeploy(image, name, namespace, command, replicas)
}

// CreateWindowsDeploy will create a deployment for a given image with a name in a namespace and create a service mapping a hostPort
func CreateWindowsDeploy(pattern, image, name, namespace, miscOpts string) (*Deployment, error) {
	var commandTimeout time.Duration
	overrides := `{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}}}`
	var args []string
	args = append(args, "run", name)
	args = append(args, "-n", namespace)
	args = append(args, "--image", image, "--image-pull-policy=IfNotPresent")
	args = append(args, "--overrides", overrides)
	if miscOpts != "" {
		args = append(args, miscOpts)
	}
	cmd := exec.Command("k", args[:]...)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	d, err := GetWithRetry(name, namespace, 3*time.Second, deploymentGetAfterCreateTimeout)
	if err != nil {
		log.Printf("Error while trying to fetch Deployment %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return d, nil
}

// CreateWindowsDeployWithHostport will create a deployment for a given image with a name in a namespace and create a service mapping a hostPort
func CreateWindowsDeployWithHostport(image, name, namespace string, port int, hostport int) (*Deployment, error) {
	var commandTimeout time.Duration
	overrides := `{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}}}`
	cmd := exec.Command("k", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--port", strconv.Itoa(port), "--hostport", strconv.Itoa(hostport), "--overrides", overrides)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	d, err := GetWithRetry(name, namespace, 3*time.Second, deploymentGetAfterCreateTimeout)
	if err != nil {
		log.Printf("Error while trying to fetch Deployment %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return d, nil
}

// CreateWindowsDeployWithHostportIfNotExist first checks if a deployment already exists, and return it if so
// If not, we call CreateWindowsDeploy
func CreateWindowsDeployWithHostportIfNotExist(image, name, namespace string, port int, hostport int) (*Deployment, error) {
	deployment, err := Get(name, namespace, validateDeploymentNotExistRetries)
	if err != nil {
		return CreateWindowsDeployWithHostport(image, name, namespace, port, hostport)
	}
	return deployment, nil
}

// CreateWindowsDeployWithHostportDeleteIfExist first checks if a deployment already exists according to a naming pattern
// If a pre-existing deployment is found matching that pattern, it is deleted
func CreateWindowsDeployWithHostportDeleteIfExist(pattern, image, name, namespace string, port int, hostport int) (*Deployment, error) {
	deployments, err := GetAllByPrefix(pattern, namespace)
	if err != nil {
		return nil, err
	}
	for _, d := range deployments {
		d.Delete(util.DefaultDeleteRetries)
	}
	return CreateWindowsDeployWithHostport(image, name, namespace, port, hostport)
}

// CreateWindowsDeployDeleteIfExist first checks if a deployment already exists according to a naming pattern
// If a pre-existing deployment is found matching that pattern, it is deleted
func CreateWindowsDeployDeleteIfExist(pattern, image, name, namespace, miscOpts string) (*Deployment, error) {
	deployments, err := GetAllByPrefix(pattern, namespace)
	if err != nil {
		return nil, err
	}
	for _, d := range deployments {
		d.Delete(util.DefaultDeleteRetries)
	}
	return CreateWindowsDeploy(pattern, image, name, namespace, miscOpts)
}

// Get returns a deployment from a name and namespace
func Get(name, namespace string, retries int) (*Deployment, error) {
	d := Deployment{}
	var out []byte
	var err error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "get", "deploy", name, "-n", namespace, "-o", "json")
		out, err = cmd.CombinedOutput()
		if err != nil {
			util.PrintCommand(cmd)
			log.Printf("Error getting deployment: %s\n", err)
		} else {
			jsonErr := json.Unmarshal(out, &d)
			if jsonErr != nil {
				log.Printf("Error unmarshalling deployment json:%s\n", jsonErr)
				err = jsonErr
			}
		}
		time.Sleep(3 * time.Second)
	}
	return &d, err
}

// GetAll will return all deployments in a given namespace
func GetAll(namespace string) (*List, error) {
	cmd := exec.Command("k", "get", "deployments", "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error getting all deployments:\n")
		return nil, err
	}
	dl := List{}
	err = json.Unmarshal(out, &dl)
	if err != nil {
		log.Printf("Error unmarshalling deployments json:%s\n", err)
		return nil, err
	}
	return &dl, nil
}

// GetAllByPrefix will return all pods in a given namespace that match a prefix
func GetAllByPrefix(prefix, namespace string) ([]Deployment, error) {
	dl, err := GetAll(namespace)
	if err != nil {
		return nil, err
	}
	deployments := []Deployment{}
	for _, d := range dl.Deployments {
		matched, err := regexp.MatchString(prefix+"-.*", d.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match deployment name:%s\n", err)
			return nil, err
		}
		if matched {
			deployments = append(deployments, d)
		}
	}
	return deployments, nil
}

// Describe will describe a deployment resource
func (d *Deployment) Describe() error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "describe", "deployment", d.Metadata.Name, "-n", d.Metadata.Namespace)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	log.Printf("\n%s\n", string(out))
	return err
}

// Delete will delete a deployment in a given namespace
func (d *Deployment) Delete(retries int) error {
	var zeroValueDuration time.Duration
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "delete", "deploy", "-n", d.Metadata.Namespace, d.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd, zeroValueDuration)
		if kubectlError != nil {
			log.Printf("Error while trying to delete deployment %s in namespace %s:%s\n", d.Metadata.Namespace, d.Metadata.Name, string(kubectlOutput))
			continue
		}
		break
	}

	if kubectlError != nil {
		return kubectlError
	}

	if d.Metadata.HasHPA {
		for i := 0; i < retries; i++ {
			cmd := exec.Command("k", "delete", "hpa", "-n", d.Metadata.Namespace, d.Metadata.Name)
			kubectlOutput, kubectlError = util.RunAndLogCommand(cmd, zeroValueDuration)
			if kubectlError != nil {
				log.Printf("Deployment %s has associated HPA but unable to delete in namespace %s:%s\n", d.Metadata.Namespace, d.Metadata.Name, string(kubectlOutput))
				continue
			}
			break
		}
	}

	return kubectlError
}

// Expose will create a load balancer and expose the deployment on a given port
func (d *Deployment) Expose(svcType string, targetPort, exposedPort int) error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "expose", "deployment", d.Metadata.Name, "--type", svcType, "-n", d.Metadata.Namespace, "--target-port", strconv.Itoa(targetPort), "--port", strconv.Itoa(exposedPort))
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	if err != nil {
		log.Printf("Error while trying to expose (%s) target port (%v) for deployment %s in namespace %s on port %v:%s\n", svcType, targetPort, d.Metadata.Name, d.Metadata.Namespace, exposedPort, string(out))
		return err
	}
	return nil
}

// ExposeIfNotExist will create a load balancer and expose the deployment on a given port if the associated service doesn't already exist
func (d *Deployment) ExposeIfNotExist(svcType string, targetPort, exposedPort int) error {
	_, err := service.Get(d.Metadata.Name, d.Metadata.Namespace)
	if err != nil {
		return d.Expose(svcType, targetPort, exposedPort)
	}
	return nil
}

// ExposeDeleteIfExist will create a load balancer and expose the deployment on a given port
// If a service matching the passed in pattern already exists, we'll delete it first
func (d *Deployment) ExposeDeleteIfExist(pattern, namespace, svcType string, targetPort, exposedPort int) error {
	services, err := service.GetAllByPrefix(pattern, namespace)
	if err != nil {
		return err
	}
	for _, s := range services {
		s.Delete(util.DefaultDeleteRetries)
	}
	return d.Expose(svcType, targetPort, exposedPort)
}

// ScaleDeployment scales a deployment to n instancees
func (d *Deployment) ScaleDeployment(n int) error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "scale", fmt.Sprintf("--replicas=%d", n), "deployment", d.Metadata.Name)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	if err != nil {
		log.Printf("Error while scaling deployment %s to %d pods:%s\n", d.Metadata.Name, n, string(out))
		return err
	}
	return nil
}

// CreateDeploymentHPA applies autoscale characteristics to deployment
func (d *Deployment) CreateDeploymentHPA(cpuPercent, min, max int) error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "autoscale", "deployment", d.Metadata.Name, fmt.Sprintf("--cpu-percent=%d", cpuPercent),
		fmt.Sprintf("--min=%d", min), fmt.Sprintf("--max=%d", max))
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	if err != nil {
		log.Printf("Error while configuring autoscale against deployment %s:%s\n", d.Metadata.Name, string(out))
		return err
	}
	d.Metadata.HasHPA = true
	return nil
}

// CreateDeploymentHPADeleteIfExist applies autoscale characteristics to deployment, deleting any pre-existing HPA resource first
func (d *Deployment) CreateDeploymentHPADeleteIfExist(cpuPercent, min, max int) error {
	h, err := hpa.Get(d.Metadata.Name, d.Metadata.Namespace, 5)
	if err == nil {
		err := h.Delete(util.DefaultDeleteRetries)
		if err != nil {
			return err
		}
		_, err = hpa.WaitOnDeleted(d.Metadata.Name, d.Metadata.Namespace, 5*time.Second, 1*time.Minute)
		if err != nil {
			return err
		}
	}
	return d.CreateDeploymentHPA(cpuPercent, min, max)
}

// Pods will return all pods related to a deployment
func (d *Deployment) Pods() ([]pod.Pod, error) {
	return pod.GetAllByPrefixWithRetry(d.Metadata.Name, d.Metadata.Namespace, 3*time.Second, 20*time.Minute)
}

// PodsRunning will return all pods in a Running state related to a deployment
func (d *Deployment) PodsRunning() ([]pod.Pod, error) {
	return pod.GetAllRunningByPrefixWithRetry(d.Metadata.Name, d.Metadata.Namespace, 3*time.Second, 20*time.Minute)
}

// GetWithRetry gets a deployment, allowing for retries
func GetWithRetry(name, namespace string, sleep, timeout time.Duration) (*Deployment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentGetWithRetryError error
	var deployment *Deployment
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- GetAsync(name, namespace):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetWithRetryError = result.err
			deployment = result.deployment
			if mostRecentGetWithRetryError == nil {
				if deployment != nil {
					return deployment, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetWithRetry timed out: %s\n", mostRecentGetWithRetryError)
		}
	}
}

// GetResult is a return struct for GetAsync
type GetResult struct {
	deployment *Deployment
	err        error
}

// GetAsync wraps Get with a struct response for goroutine + channel usage
func GetAsync(name, namespace string) GetResult {
	deployment, err := Get(name, namespace, 1)
	return GetResult{
		deployment: deployment,
		err:        err,
	}
}

// WaitForReplicas waits for a pod replica count between min and max
func (d *Deployment) WaitForReplicas(min, max int, sleep, timeout time.Duration) ([]pod.Pod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan pod.GetAllByPrefixResult)
	var mostRecentWaitForReplicasError error
	var pods []pod.Pod
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- pod.GetAllRunningByPrefixAsync(d.Metadata.Name, d.Metadata.Namespace):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitForReplicasError = result.Err
			pods = result.Pods
			if mostRecentWaitForReplicasError == nil {
				if min == -1 {
					if len(pods) <= max {
						return pods, nil
					}
				} else if max == -1 {
					if len(pods) >= min {
						return pods, nil
					}
				} else {
					if len(pods) >= min && len(pods) <= max {
						return pods, nil
					}
				}
			}
		case <-ctx.Done():
			err := d.Describe()
			if err != nil {
				log.Printf("Unable to describe deployment %s: %s", d.Metadata.Name, err)
			}
			return pods, errors.Errorf("WaitForReplicas timed out: %s\n", mostRecentWaitForReplicasError)
		}
	}
}
