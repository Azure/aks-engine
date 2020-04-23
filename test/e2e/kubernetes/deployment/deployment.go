//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package deployment

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
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

// CreateLinuxDeployAsync wraps CreateLinuxDeploy with a struct response for goroutine + channel usage
func CreateLinuxDeployAsync(ctx context.Context, image, name, namespace, app, role string) GetResult {
	d, err := CreateLinuxDeploy(image, name, namespace, app, role)
	return GetResult{
		deployment: d,
		err:        err,
	}
}

// CreateLinuxDeployWithRetry will create a deployment for a given image with a name in a namespace, with retry
func CreateLinuxDeployWithRetry(image, name, namespace, app, role string, sleep, timeout time.Duration) (*Deployment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentCreateLinuxDeployWithRetryWithRetryError error
	var d *Deployment
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- CreateLinuxDeployAsync(ctx, image, name, namespace, app, role)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentCreateLinuxDeployWithRetryWithRetryError = result.err
			d = result.deployment
			if mostRecentCreateLinuxDeployWithRetryWithRetryError == nil {
				return d, nil
			}
		case <-ctx.Done():
			return d, errors.Errorf("CreateLinuxDeployWithRetry timed out: %s\n", mostRecentCreateLinuxDeployWithRetryWithRetryError)
		}
	}
}

const webDeploymentTmpl = `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: %s
    role: %s
  name: %s
spec:
  replicas: 1
  selector:
    matchLabels:
      app: %s
      role: %s
  template:
    metadata:
      labels:
        app: %s
        role: %s
    spec:
      containers:
      - image: %s
        name: %s
        resources:
          requests:
            cpu: 10m
            memory: 10M
      nodeSelector:
        beta.kubernetes.io/os: %s
`

// CreateLinuxDeploy will create a deployment for a given image with a name in a namespace
func CreateLinuxDeploy(image, name, namespace, app, role string) (*Deployment, error) {
	var commandTimeout time.Duration

	tmpFile, err := ioutil.TempFile("", "e2e-linux-deployment-*.yaml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if app == "" {
		app = "webapp"
	}
	if role == "" {
		role = "any"
	}
	manifest := fmt.Sprintf(webDeploymentTmpl,
		app, role, name, app, role, app, role, image, name, "linux")
	fmt.Fprintln(tmpFile, manifest)

	cmd := exec.Command("k", "apply", "-n", namespace, "-f", tmpFile.Name())

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
func CreateLinuxDeployIfNotExist(image, name, namespace, app, role string) (*Deployment, error) {
	deployment, err := Get(name, namespace, validateDeploymentNotExistRetries)
	if err != nil {
		return CreateLinuxDeploy(image, name, namespace, app, role)
	}
	return deployment, nil
}

// CreateLinuxDeployDeleteIfExists will create a deployment, deleting any pre-existing deployment with the same name
func CreateLinuxDeployDeleteIfExists(pattern, image, name, namespace, app, role string, timeout time.Duration) (*Deployment, error) {
	deployments, err := GetAllByPrefixWithRetry(pattern, namespace, 5*time.Second, timeout)
	if err != nil {
		return nil, err
	}
	for _, d := range deployments {
		d.Delete(util.DefaultDeleteRetries)
	}
	return CreateLinuxDeployWithRetry(image, name, namespace, app, role, 3*time.Minute, timeout)
}

const runDeploymentTmpl = `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    run: %s
  name: %s
spec:
  replicas: %d
  selector:
    matchLabels:
      run: %s
  template:
    metadata:
      labels:
        run: %s
    spec:
      containers:
      - image: %s
        name: %s
        command:
        - /bin/sh
        - -c
        - "%s"
      nodeSelector:
        beta.kubernetes.io/os: %s
`

// RunLinuxDeploy will create a deployment that runs a bash command in a pod
func RunLinuxDeploy(image, name, namespace, command string, replicas int) (*Deployment, error) {
	var commandTimeout time.Duration

	tmpFile, err := ioutil.TempFile("", "e2e-linux-deployment-*.yaml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	manifest := fmt.Sprintf(runDeploymentTmpl,
		name, name, replicas, name, name, image, name, command, "linux")
	fmt.Fprintln(tmpFile, manifest)

	cmd := exec.Command("k", "apply", "-n", namespace, "-f", tmpFile.Name())

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

type deployRunnerCmd func(string, string, string, string, int) (*Deployment, error)

// RunDeploymentMultipleTimes runs the same command 'desiredAttempts' times
func RunDeploymentMultipleTimes(deployRunnerCmd deployRunnerCmd, image, name, command string, replicas, desiredAttempts int, sleep, podTimeout, timeout time.Duration) (int, error) {
	var successfulAttempts int
	var actualAttempts int
	logResults := func() {
		log.Printf("Ran command on %d of %d desired attempts with %d successes\n\n", actualAttempts, desiredAttempts, successfulAttempts)
	}
	defer logResults()
	for i := 0; i < desiredAttempts; i++ {
		actualAttempts++
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		deploymentName := fmt.Sprintf("%s-%d", name, r.Intn(99999))
		var d *Deployment
		var err error
		d, err = deployRunnerCmd(image, deploymentName, "default", command, replicas)
		if err != nil {
			return successfulAttempts, err
		}
		pods, err := d.WaitForReplicas(replicas, replicas, sleep, timeout)
		if err != nil {
			log.Printf("deployment %s did not have the expected replica count %d in time\n", deploymentName, replicas)
			return successfulAttempts, err
		}
		var podsSucceeded int
		for _, p := range pods {
			running, err := pod.WaitOnSuccesses(p.Metadata.Name, p.Metadata.Namespace, 6, sleep, podTimeout)
			if err != nil {
				log.Printf("pod %s did not succeed in time\n", p.Metadata.Name)
				return successfulAttempts, err
			}
			if running {
				podsSucceeded++
			}
		}
		err = d.Delete(util.DefaultDeleteRetries)
		if err != nil {
			return successfulAttempts, err
		}
		if podsSucceeded == replicas {
			successfulAttempts++
		}
	}

	return successfulAttempts, nil
}

// CreateWindowsDeployAsync wraps CreateWindowsDeploy with a struct response for goroutine + channel usage
func CreateWindowsDeployAsync(image, name, namespace, app, role string) GetResult {
	d, err := CreateWindowsDeploy(image, name, namespace, app, role)
	return GetResult{
		deployment: d,
		err:        err,
	}
}

// CreateWindowsDeployWithRetry will return all deployments in a given namespace that match a prefix, retrying if error up to a timeout
func CreateWindowsDeployWithRetry(image, name, namespace, app, role string, sleep, timeout time.Duration) (*Deployment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentCreateWindowsDeployWithRetryError error
	var d *Deployment
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- CreateWindowsDeployAsync(image, name, namespace, app, role)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentCreateWindowsDeployWithRetryError = result.err
			d = result.deployment
			if mostRecentCreateWindowsDeployWithRetryError == nil {
				return d, nil
			}
		case <-ctx.Done():
			return d, errors.Errorf("GetAllByPrefixWithRetry timed out: %s\n", mostRecentCreateWindowsDeployWithRetryError)
		}
	}
}

// CreateWindowsDeploy will create a deployment for a given image with a name in a namespace and create a service mapping a hostPort
func CreateWindowsDeploy(image, name, namespace, app, role string) (*Deployment, error) {
	var commandTimeout time.Duration
	var cmd *exec.Cmd

	tmpFile, err := ioutil.TempFile("", "e2e-windows-deployment-*.yaml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if app == "" {
		app = "webapp"
	}
	if role == "" {
		role = "any"
	}
	manifest := fmt.Sprintf(webDeploymentTmpl,
		app, role, name, app, role, app, role, image, name, "windows")
	fmt.Fprintln(tmpFile, manifest)

	cmd = exec.Command("k", "apply", "-n", namespace, "-f", tmpFile.Name())

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

const hostportDeploymentTmpl = `---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    run: %s
  name: %s
spec:
  replicas: 1
  selector:
    matchLabels:
      run: %s
  template:
    metadata:
      labels:
        run: %s
    spec:
      containers:
      - image: %s
        name: %s
        ports:
        - containerPort: %d%s
      nodeSelector:
        beta.kubernetes.io/os: %s
`

// CreateWindowsDeployWithHostport will create a deployment for a given image with a name in a namespace and create a service mapping a hostPort
func CreateWindowsDeployWithHostport(image, name, namespace string, port int, hostport int) (*Deployment, error) {
	var commandTimeout time.Duration

	tmpFile, err := ioutil.TempFile("", "e2e-windows-deployment-*.yaml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	var hostportStr string
	if hostport != -1 {
		hostportStr = fmt.Sprintf("\n          hostPort: %d", hostport)
	}

	manifest := fmt.Sprintf(hostportDeploymentTmpl,
		name, name, name, name, image, name, port, hostportStr, "windows")
	fmt.Fprintln(tmpFile, manifest)

	cmd := exec.Command("k", "apply", "-n", namespace, "-f", tmpFile.Name())

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
func CreateWindowsDeployDeleteIfExist(pattern, image, name, namespace, app, role string, timeout time.Duration) (*Deployment, error) {
	deployments, err := GetAllByPrefixWithRetry(pattern, namespace, 5*time.Second, timeout)
	if err != nil {
		return nil, err
	}
	for _, d := range deployments {
		d.Delete(util.DefaultDeleteRetries)
	}
	return CreateWindowsDeployWithRetry(image, name, namespace, app, role, 3*time.Minute, timeout)
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

// GetAllByPrefixResult is the result type for GetAllByPrefixAsync
type GetAllByPrefixResult struct {
	Deployments []Deployment
	Err         error
}

// GetAllByPrefixAsync wraps GetAllByPrefix with a struct response for goroutine + channel usage
func GetAllByPrefixAsync(prefix, namespace string) GetAllByPrefixResult {
	deployments, err := GetAllByPrefix(prefix, namespace)
	return GetAllByPrefixResult{
		Deployments: deployments,
		Err:         err,
	}
}

// GetAllByPrefixWithRetry will return all deployments in a given namespace that match a prefix, retrying if error up to a timeout
func GetAllByPrefixWithRetry(prefix, namespace string, sleep, timeout time.Duration) ([]Deployment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetAllByPrefixResult)
	var mostRecentGetAllByPrefixWithRetryError error
	var deployments []Deployment
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAllByPrefixAsync(prefix, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetAllByPrefixWithRetryError = result.Err
			deployments = result.Deployments
			if mostRecentGetAllByPrefixWithRetryError == nil {
				return deployments, nil
			}
		case <-ctx.Done():
			return deployments, errors.Errorf("GetAllByPrefixWithRetry timed out: %s\n", mostRecentGetAllByPrefixWithRetryError)
		}
	}
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
			default:
				ch <- GetAsync(name, namespace)
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
			default:
				ch <- pod.GetAllRunningByPrefixAsync(d.Metadata.Name, d.Metadata.Namespace)
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
