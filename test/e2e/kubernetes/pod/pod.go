//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package pod

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const (
	testDir                    string = "testdirectory"
	validatePodNotExistRetries        = 3
)

// List is a container that holds all pods returned from doing a kubectl get pods
type List struct {
	Pods []Pod `json:"items"`
}

// Pod is used to parse data from kubectl get pods
type Pod struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
}

// Metadata holds information like name, createdat, labels, and namespace
type Metadata struct {
	CreatedAt time.Time         `json:"creationTimestamp"`
	Labels    map[string]string `json:"labels"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
}

// Spec holds information like containers
type Spec struct {
	Containers []Container `json:"containers"`
	NodeName   string      `json:"nodeName"`
}

// Container holds information like image and ports
type Container struct {
	Image     string    `json:"image"`
	Ports     []Port    `json:"ports"`
	Env       []EnvVar  `json:"env"`
	Resources Resources `json:"resources"`
	Name      string    `json:"name"`
	Args      []string  `json:"args"`
}

// TerminatedContainerState shows terminated state of a container
type TerminatedContainerState struct {
	ContainerID string `json:"containerID"`
	ExitCode    int    `json:"exitCode"`
	FinishedAt  string `json:"finishedAt"`
	Reason      string `json:"reason"`
	StartedAt   string `json:"startedAt"`
}

// ContainerState has state of a container
type ContainerState struct {
	Terminated TerminatedContainerState `json:"terminated"`
}

// ContainerStatus has status of a container
type ContainerStatus struct {
	ContainerID  string         `json:"containerID"`
	Image        string         `json:"image"`
	ImageID      string         `json:"imageID"`
	Name         string         `json:"name"`
	Ready        bool           `json:"ready"`
	RestartCount int            `json:"restartCount"`
	State        ContainerState `json:"state"`
	LastState    ContainerState `json:"lastState"`
}

// EnvVar holds environment variables
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Port represents a container port definition
type Port struct {
	ContainerPort int `json:"containerPort"`
	HostPort      int `json:"hostPort"`
}

// Resources represents a container resources definition
type Resources struct {
	Requests Requests `json:"requests"`
	Limits   Limits   `json:"limits"`
}

// Requests represents container resource requests
type Requests struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// Limits represents container resource limits
type Limits struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// Status holds information like hostIP and phase
type Status struct {
	HostIP            string            `json:"hostIP"`
	Phase             string            `json:"phase"`
	PodIP             string            `json:"podIP"`
	StartTime         time.Time         `json:"startTime"`
	ContainerStatuses []ContainerStatus `json:"containerStatuses"`
}

// ReplaceContainerImageFromFile loads in a YAML, finds the image: line, and replaces it with the value of containerImage
func ReplaceContainerImageFromFile(filename, containerImage string) (string, error) {
	var outString string
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening source YAML file %s\n", filename)
		return "", err
	}
	defer file.Close()
	re := regexp.MustCompile("(image:) .*$")
	replacementString := "$1 " + containerImage
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		outString += re.ReplaceAllString(scanner.Text(), replacementString) + "\n"
	}
	err = scanner.Err()
	if err != nil {
		return "", err
	}
	_, filenameOnly := path.Split(filename)
	tmpFile, err := ioutil.TempFile(os.TempDir(), filenameOnly)
	if err != nil {
		return "", err
	}
	_, err = tmpFile.Write([]byte(outString))
	return tmpFile.Name(), err
}

// CreatePodFromFile will create a Pod from file with a name
func CreatePodFromFile(filename, name, namespace string, sleep, timeout time.Duration) (*Pod, error) {
	cmd := exec.Command("k", "apply", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create Pod %s:%s\n", name, string(out))
		return nil, err
	}
	p, err := GetWithRetry(name, namespace, sleep, timeout)
	if err != nil {
		log.Printf("Error while trying to fetch Pod %s:%s\n", name, err)
		return nil, err
	}
	return p, nil
}

// CreatePodFromFileIfNotExist will create a Pod from file with a name
func CreatePodFromFileIfNotExist(filename, name, namespace string, sleep, timeout time.Duration) (*Pod, error) {
	p, err := Get(name, namespace, validatePodNotExistRetries)
	if err != nil {
		return CreatePodFromFile(filename, name, namespace, sleep, timeout)
	}
	return p, nil
}

// RunLinuxPod will create a pod that runs a bash command
// --overrides := `"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}`
func RunLinuxPod(image, name, namespace, command string, printOutput bool, sleep, commandTimeout, podGetTimeout time.Duration) (*Pod, error) {
	overrides := `{ "spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}`
	cmd := exec.Command("k", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--restart=Never", "--overrides", overrides, "--command", "--", "/bin/sh", "-c", command)
	var out []byte
	var err error
	if printOutput {
		out, err = util.RunAndLogCommand(cmd, commandTimeout)
	} else {
		out, err = cmd.CombinedOutput()
	}
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	p, err := GetWithRetry(name, namespace, sleep, podGetTimeout)
	if err != nil {
		log.Printf("Error while trying to fetch Pod %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return p, nil
}

// RunWindowsPod will create a pod that runs a powershell command
// --overrides := `"spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}`
func RunWindowsPod(image, name, namespace, command string, printOutput bool, sleep, commandTimeout time.Duration, podGetTimeout time.Duration) (*Pod, error) {
	overrides := `{ "spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}`
	cmd := exec.Command("k", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--restart=Never", "--overrides", overrides, "--command", "--", "powershell", command)
	var out []byte
	var err error
	if printOutput {
		out, err = util.RunAndLogCommand(cmd, commandTimeout)
	} else {
		out, err = cmd.CombinedOutput()
	}
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	p, err := GetWithRetry(name, namespace, sleep, podGetTimeout)
	if err != nil {
		log.Printf("Error while trying to fetch Pod %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return p, nil
}

type podRunnerCmd func(string, string, string, string, bool, time.Duration, time.Duration, time.Duration) (*Pod, error)

// RunCommandMultipleTimes runs the same command 'desiredAttempts' times
func RunCommandMultipleTimes(podRunnerCmd podRunnerCmd, image, name, command string, desiredAttempts int, sleep, commandTimeout, timeout time.Duration) (int, error) {
	var successfulAttempts int
	var actualAttempts int
	logResults := func() {
		log.Printf("Ran command on %d of %d desired attempts with %d successes\n\n", actualAttempts, desiredAttempts, successfulAttempts)
	}
	defer logResults()
	for i := 0; i < desiredAttempts; i++ {
		actualAttempts++
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		podName := fmt.Sprintf("%s-%d", name, r.Intn(99999))
		var p *Pod
		var err error
		p, err = podRunnerCmd(image, podName, "default", command, true, sleep, timeout, timeout)
		if err != nil {
			return successfulAttempts, err
		}
		succeeded, err := p.WaitOnSucceeded(sleep, timeout)
		if err != nil {
			log.Printf("pod %s did not succeed in time\n", podName)
			return successfulAttempts, err
		}
		terminated, err := p.WaitOnTerminated(podName, sleep, commandTimeout, timeout)
		if err != nil {
			log.Printf("pod %s container %s did not reach a terminal exit 0 state in time\n", podName, podName)
			return successfulAttempts, err
		}
		err = p.Delete(util.DefaultDeleteRetries)
		if err != nil {
			return successfulAttempts, err
		}
		if succeeded && terminated {
			successfulAttempts++
		}
	}

	return successfulAttempts, nil
}

// GetAll will return all pods in a given namespace
func GetAll(namespace string) (*List, error) {
	cmd := exec.Command("k", "get", "pods", "-n", namespace, "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error getting pod:\n")
		util.PrintCommand(cmd)
		return nil, err
	}
	pl := List{}
	err = json.Unmarshal(out, &pl)
	if err != nil {
		log.Printf("Error unmarshalling pods json:%s\n", err)
		return nil, err
	}
	return &pl, nil
}

// GetWithRetry gets a pod, allowing for retries
func GetWithRetry(podName, namespace string, sleep, timeout time.Duration) (*Pod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentGetWithRetryError error
	var pod *Pod
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAsync(podName, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetWithRetryError = result.err
			pod = result.pod
			if mostRecentGetWithRetryError == nil {
				if pod != nil {
					return pod, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetWithRetry timed out: %s\n", mostRecentGetWithRetryError)
		}
	}
}

// RunLinuxWithRetry runs a command in a Linux pod, allowing for retries
func RunLinuxWithRetry(image, name, namespace, command string, printOutput bool, sleep, timeout time.Duration) (*Pod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentRunWithRetryError error
	var pod *Pod
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- RunLinuxAsyncDeleteIfExists(image, name, namespace, command, printOutput, sleep, timeout, timeout)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentRunWithRetryError = result.err
			pod = result.pod
			if mostRecentRunWithRetryError == nil {
				if pod != nil {
					return pod, nil
				}
			} else {
				if pod != nil {
					err := pod.Logs()
					if err != nil {
						log.Printf("Unable to print pod logs for pod %s: %s", pod.Metadata.Name, err)
					}
					err = pod.Describe()
					if err != nil {
						log.Printf("Unable to describe pod %s: %s", pod.Metadata.Name, err)
					}
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("RunLinuxWithRetry timed out: %s\n", mostRecentRunWithRetryError)
		}
	}
}

// RunWindowsWithRetry runs a command in a Windows pod, allowing for retries
func RunWindowsWithRetry(image, name, namespace, command string, printOutput bool, sleep, timeout time.Duration) (*Pod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentRunWithRetryError error
	var pod *Pod
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- RunWindowsAsyncDeleteIfExists(image, name, namespace, command, printOutput, sleep, timeout, timeout)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentRunWithRetryError = result.err
			pod = result.pod
			if mostRecentRunWithRetryError == nil {
				if pod != nil {
					return pod, nil
				}
			} else {
				if pod != nil {
					err := pod.Logs()
					if err != nil {
						log.Printf("Unable to print pod logs for pod %s: %s", pod.Metadata.Name, err)
					}
					err = pod.Describe()
					if err != nil {
						log.Printf("Unable to describe pod %s: %s", pod.Metadata.Name, err)
					}
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("RunWindowsWithRetry timed out: %s\n", mostRecentRunWithRetryError)
		}
	}
}

// GetResult is a return struct for GetAsync
type GetResult struct {
	pod *Pod
	err error
}

// GetAsync wraps Get with a struct response for goroutine + channel usage
func GetAsync(podName, namespace string) GetResult {
	pod, err := Get(podName, namespace, 1)
	return GetResult{
		pod: pod,
		err: err,
	}
}

// RunLinuxAsync wraps RunLinuxPod with a struct response for goroutine + channel usage
func RunLinuxAsync(image, name, namespace, command string, printOutput bool, sleep, commandTimeout, podGetTimeout time.Duration) GetResult {
	pod, err := RunLinuxPod(image, name, namespace, command, printOutput, sleep, commandTimeout, podGetTimeout)
	return GetResult{
		pod: pod,
		err: err,
	}
}

// RunLinuxAsyncDeleteIfExists wraps RunLinuxPod with a struct response for goroutine + channel usage
// If a pod with the name already exists, we delete it first and create a new one
func RunLinuxAsyncDeleteIfExists(image, name, namespace, command string, printOutput bool, sleep, commandTimeout, podGetTimeout time.Duration) GetResult {
	p, err := Get(name, namespace, 3)
	if err == nil {
		log.Printf("pod %s in namespace %s already exists, will delete\n", name, namespace)
		err := p.Delete(3)
		if err != nil {
			log.Printf("unable to delete pod %s in namespace %s\n", name, namespace)
			return GetResult{
				pod: p,
				err: err,
			}
		}
	}
	p, err = RunLinuxPod(image, name, namespace, command, printOutput, sleep, commandTimeout, podGetTimeout)
	return GetResult{
		pod: p,
		err: err,
	}
}

// RunWindowsAsync wraps RunWindowsPod with a struct response for goroutine + channel usage
func RunWindowsAsync(image, name, namespace, command string, printOutput bool, sleep, commandTimeout, podGetTimeout time.Duration) GetResult {
	pod, err := RunWindowsPod(image, name, namespace, command, printOutput, sleep, commandTimeout, podGetTimeout)
	return GetResult{
		pod: pod,
		err: err,
	}
}

// RunWindowsAsyncDeleteIfExists wraps RunWindowsPod with a struct response for goroutine + channel usage
// If a pod with the name already exists, we delete it first and create a new one
func RunWindowsAsyncDeleteIfExists(image, name, namespace, command string, printOutput bool, sleep, commandTimeout, podGetTimeout time.Duration) GetResult {
	p, err := Get(name, namespace, 3)
	if err == nil {
		log.Printf("pod %s in namespace %s already exists, will delete\n", name, namespace)
		err := p.Delete(3)
		if err != nil {
			log.Printf("unable to delete pod %s in namespace %s\n", name, namespace)
			return GetResult{
				pod: p,
				err: err,
			}
		}
	}
	p, err = RunWindowsPod(image, name, namespace, command, printOutput, sleep, commandTimeout, podGetTimeout)
	return GetResult{
		pod: p,
		err: err,
	}
}

// Get will return a pod with a given name and namespace
func Get(podName, namespace string, retries int) (*Pod, error) {
	p := Pod{}
	var out []byte
	var err error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "get", "pods", podName, "-n", namespace, "-o", "json")
		out, err = cmd.CombinedOutput()
		if err == nil {
			jsonErr := json.Unmarshal(out, &p)
			if jsonErr != nil {
				log.Printf("Error unmarshalling pods json:%s\n", jsonErr)
				err = jsonErr
			}
		}
		time.Sleep(3 * time.Second)
	}
	return &p, err
}

// PrintPodsLogs prints logs for all pods whose name matches a substring
func PrintPodsLogs(podPrefix, namespace string, sleep, timeout time.Duration) {
	pods, err := GetAllByPrefixWithRetry(podPrefix, namespace, sleep, timeout)
	if err != nil {
		log.Printf("Unable to print logs for pods matching prefix %s in namespace %s: %s", podPrefix, namespace, err)
	}
	for _, p := range pods {
		err := p.Logs()
		if err != nil {
			log.Printf("Unable to print pod logs for pod %s: %s", p.Metadata.Name, err)
		}
		err = p.Describe()
		if err != nil {
			log.Printf("Unable to describe pod %s: %s", p.Metadata.Name, err)
		}
	}
}

// GetAllByPrefixResult is the result type for GetAllByPrefixAsync
type GetAllByPrefixResult struct {
	Pods []Pod
	Err  error
}

// GetAllByPrefixAsync wraps GetAllByPrefix with a struct response for goroutine + channel usage
func GetAllByPrefixAsync(prefix, namespace string) GetAllByPrefixResult {
	pods, err := GetAllByPrefix(prefix, namespace)
	return GetAllByPrefixResult{
		Pods: pods,
		Err:  err,
	}
}

// GetAllByPrefixWithRetry will return all pods in a given namespace that match a prefix, retrying if error up to a timeout
func GetAllByPrefixWithRetry(prefix, namespace string, sleep, timeout time.Duration) ([]Pod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetAllByPrefixResult)
	var mostRecentGetAllByPrefixWithRetryError error
	var pods []Pod
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
			pods = result.Pods
			if mostRecentGetAllByPrefixWithRetryError == nil {
				return pods, nil
			}
		case <-ctx.Done():
			return pods, errors.Errorf("GetAllByPrefixWithRetry timed out: %s\n", mostRecentGetAllByPrefixWithRetryError)
		}
	}
}

// GetAllByPrefix will return all pods in a given namespace that match a prefix
func GetAllByPrefix(prefix, namespace string) ([]Pod, error) {
	pl, err := GetAll(namespace)
	if err != nil {
		return nil, err
	}
	pods := []Pod{}
	for _, p := range pl.Pods {
		matched, err := regexp.MatchString(prefix+"-.*", p.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return nil, err
		}
		if matched {
			pods = append(pods, p)
		}
	}
	return pods, nil
}

// GetAllRunningByPrefixAsync wraps GetAllRunningByPrefix with a struct response for goroutine + channel usage
func GetAllRunningByPrefixAsync(prefix, namespace string) GetAllByPrefixResult {
	pods, err := GetAllRunningByPrefix(prefix, namespace)
	return GetAllByPrefixResult{
		Pods: pods,
		Err:  err,
	}
}

// GetAllRunningByPrefix will return all Running pods in a given namespace that match a prefix
func GetAllRunningByPrefix(prefix, namespace string) ([]Pod, error) {
	pl, err := GetAll(namespace)
	if err != nil {
		return nil, err
	}
	pods := []Pod{}
	for _, p := range pl.Pods {
		matched, err := regexp.MatchString(prefix+"-.*", p.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return nil, err
		}
		if matched {
			if p.Status.Phase == "Running" {
				pods = append(pods, p)
			}
		}
	}
	return pods, nil
}

// GetAllRunningByPrefixWithRetry will return all Running pods in a given namespace that match a prefix, retrying if error up to a timeout
func GetAllRunningByPrefixWithRetry(prefix, namespace string, sleep, timeout time.Duration) ([]Pod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetAllByPrefixResult)
	var mostRecentGetAllRunningByPrefixWithRetryError error
	var pods []Pod
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAllRunningByPrefixAsync(prefix, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetAllRunningByPrefixWithRetryError = result.Err
			pods = result.Pods
			if mostRecentGetAllRunningByPrefixWithRetryError == nil {
				return pods, nil
			}
		case <-ctx.Done():
			return pods, errors.Errorf("GetAllRunningByPrefixWithRetry timed out: %s\n", mostRecentGetAllRunningByPrefixWithRetryError)
		}
	}
}

// AreAllPodsRunningResult is a return struct for AreAllPodsRunningAsync
type AreAllPodsRunningResult struct {
	ready bool
	err   error
}

// AreAllPodsRunningAsync wraps AreAllPodsRunning with a struct response for goroutine + channel usage
func AreAllPodsRunningAsync(podPrefix, namespace string) AreAllPodsRunningResult {
	ready, err := AreAllPodsRunning(podPrefix, namespace)
	return AreAllPodsRunningResult{
		ready: ready,
		err:   err,
	}
}

// AreAllPodsRunning will return true if all pods in a given namespace are in a Running State
func AreAllPodsRunning(podPrefix, namespace string) (bool, error) {
	pl, err := GetAll(namespace)
	if err != nil {
		return false, err
	}

	var status []bool
	for _, pod := range pl.Pods {
		matched, regexErr := regexp.MatchString(podPrefix, pod.Metadata.Name)
		if regexErr != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return false, regexErr
		}
		if matched {
			if pod.Status.Phase != "Running" {
				status = append(status, false)
			} else {
				status = append(status, true)
			}
		}
	}

	if len(status) == 0 {
		return false, nil
	}

	for _, s := range status {
		if !s {
			return false, nil
		}
	}

	return true, err
}

// AreAllPodsSucceededResult is a return struct for AreAllPodsSucceeded
type AreAllPodsSucceededResult struct {
	allPodsSucceeded bool
	anyPodsFailed    bool
	err              error
}

// AreAllPodsSucceededAsync wraps AreAllPodsSucceeded with a struct response for goroutine + channel usage
func AreAllPodsSucceededAsync(podPrefix, namespace string) AreAllPodsSucceededResult {
	allPodsSucceeded, anyPodsFailed, err := AreAllPodsSucceeded(podPrefix, namespace)
	return AreAllPodsSucceededResult{
		allPodsSucceeded: allPodsSucceeded,
		anyPodsFailed:    anyPodsFailed,
		err:              err,
	}
}

// AreAllPodsSucceeded returns true, false if all pods in a given namespace are in a Running State
// returns false, true if any one pod is in a Failed state
func AreAllPodsSucceeded(podPrefix, namespace string) (bool, bool, error) {
	pl, err := GetAll(namespace)
	if err != nil {
		return false, false, err
	}

	var status []bool
	for _, pod := range pl.Pods {
		matched, err := regexp.MatchString(podPrefix, pod.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return false, false, err
		}
		if matched {
			if pod.IsFailed() {
				return false, true, nil
			}
			if pod.IsSucceeded() {
				status = append(status, false)
			} else {
				status = append(status, true)
			}
		}
	}

	if len(status) == 0 {
		return false, false, nil
	}

	for _, s := range status {
		if !s {
			return false, false, nil
		}
	}

	return true, false, nil
}

// GetPodResult is a return struct for GetPodAsync
type GetPodResult struct {
	pod *Pod
	err error
}

// GetPodAsync wraps GetWithRetry with a struct response for goroutine + channel usage
func GetPodAsync(name, namespace string, timeout time.Duration) GetPodResult {
	p, err := GetWithRetry(name, namespace, 3*time.Second, timeout)
	return GetPodResult{
		pod: p,
		err: err,
	}
}

// WaitOnSuccesses returns true if all pods matching a prefix substring are in a succeeded state within a period of time
// successesNeeded is used to make sure we return the correct value even if the pod is in a CrashLoop
func WaitOnSuccesses(podPrefix, namespace string, successesNeeded int, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan AreAllPodsRunningResult)
	var mostRecentWaitOnSuccessesErr error
	successCount := 0
	flapCount := 0
	var lastResult bool
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- AreAllPodsRunningAsync(podPrefix, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitOnSuccessesErr = result.err
			if mostRecentWaitOnSuccessesErr == nil {
				if result.ready {
					lastResult = true
					successCount++
					if successCount >= successesNeeded {
						return true, nil
					}
				} else {
					if lastResult {
						flapCount++
						if flapCount >= (successesNeeded - 1) {
							PrintPodsLogs(podPrefix, namespace, 5*time.Second, 1*time.Minute)
							return false, errors.Errorf("Pods from deployment (%s) in namespace (%s) have been checked out as all Ready %d times, but included %d transitions away from a Ready state. This behavior may mean it is in a crashloop", podPrefix, namespace, successCount, flapCount)
						}
						lastResult = false
					}
				}
			}
		case <-ctx.Done():
			PrintPodsLogs(podPrefix, namespace, 5*time.Second, 1*time.Minute)
			return false, errors.Errorf("WaitOnReady timed out: %s\n", mostRecentWaitOnSuccessesErr)
		}
	}
}

// WaitOnSucceeded will return true if a pod is in a Succeeded State
func WaitOnSucceeded(name, namespace string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetPodResult)
	var mostRecentWaitOnSucceededPodError error
	var pod *Pod
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetPodAsync(name, namespace, timeout)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			pod = result.pod
			mostRecentWaitOnSucceededPodError = result.err
			if mostRecentWaitOnSucceededPodError == nil {
				if pod.IsSucceeded() {
					return true, nil
				}
			}
		case <-ctx.Done():
			if pod != nil {
				err := pod.Logs()
				if err != nil {
					log.Printf("Unable to print pod logs for pod %s: %s", pod.Metadata.Name, err)
				}
				err = pod.Describe()
				if err != nil {
					log.Printf("Unable to describe pod %s: %s", pod.Metadata.Name, err)
				}
			}
			return false, errors.Errorf("WaitOnSucceeded timed out: %s\n", mostRecentWaitOnSucceededPodError)
		}
	}
}

// WaitOnTerminated will return true if a pod's container is in a successful (0 exit code) Terminated State
// and completed prior to the passed in containerExecutionTimeout
func WaitOnTerminated(name, namespace, containerName string, sleep, containerExecutionTimeout, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetPodResult)
	var mostRecentWaitOnTerminatedPodError error
	var pod *Pod
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetPodAsync(name, namespace, timeout)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			pod = result.pod
			mostRecentWaitOnTerminatedPodError = result.err
			if mostRecentWaitOnTerminatedPodError == nil {
				status, err := pod.ContainerStatus(containerName)
				if err == nil {
					t1, err := time.Parse(time.RFC3339, status.StartTime())
					if err != nil {
						return false, err
					}
					t2, err := time.Parse(time.RFC3339, status.EndTime())
					if err != nil {
						return false, err
					}
					duration := t2.Sub(t1)
					if duration >= containerExecutionTimeout {
						return false, errors.Errorf("execution time %s is greater than timeout %s\n", duration.String(), containerExecutionTimeout.String())
					}
					return true, nil
				}
			}
		case <-ctx.Done():
			if pod != nil {
				err := pod.Logs()
				if err != nil {
					log.Printf("Unable to print pod logs for pod %s: %s", pod.Metadata.Name, err)
				}
				err = pod.Describe()
				if err != nil {
					log.Printf("Unable to describe pod %s: %s", pod.Metadata.Name, err)
				}
			}
			return false, errors.Errorf("WaitOnTerminated timed out: %s\n", mostRecentWaitOnTerminatedPodError)
		}
	}
}

func EnsureContainersRunningInAllPods(containers []string, podPrefix, namespace string, successesNeeded int, sleep, timeout time.Duration) error {
	running, err := WaitOnSuccesses(podPrefix, namespace, successesNeeded, sleep, timeout)
	if err != nil {
		return err
	}
	if !running {
		return errors.Errorf("%s is not in Running state", podPrefix)
	}

	pods, err := GetAllRunningByPrefixWithRetry(podPrefix, namespace, 3*time.Second, timeout)
	if err != nil {
		return err
	}

	for _, p := range pods {
		for _, c := range containers {
			if !p.HasContainer(c) {
				return errors.Errorf("%s is not running in %s", c, p.Metadata.Name)
			}
		}
	}

	return nil
}

// WaitOnReady will call the static method WaitOnReady passing in p.Metadata.Name and p.Metadata.Namespace
func (p *Pod) WaitOnReady(sleep, timeout time.Duration) (bool, error) {
	return WaitOnSuccesses(p.Metadata.Name, p.Metadata.Namespace, 6, sleep, timeout)
}

// WaitOnSucceeded will call the static method WaitOnSucceeded passing in p.Metadata.Name and p.Metadata.Namespace
func (p *Pod) WaitOnSucceeded(sleep, duration time.Duration) (bool, error) {
	return WaitOnSucceeded(p.Metadata.Name, p.Metadata.Namespace, sleep, duration)
}

// WaitOnTerminated will call the static method WaitOnTerminated passing in p.Metadata.Name and p.Metadata.Namespace
func (p *Pod) WaitOnTerminated(container string, sleep, containerExecutionTimeout, timeout time.Duration) (bool, error) {
	return WaitOnTerminated(p.Metadata.Name, p.Metadata.Namespace, container, sleep, containerExecutionTimeout, timeout)
}

func (p *Pod) attemptOutboundConn() error {
	var err error
	urls := getExternalURLs()
	for _, url := range urls {
		err = p.curlURL(url)
		if err != nil {
			return nil
		}
	}
	return err
}

func (p *Pod) curlURL(url string) error {
	_, err := p.Exec("--", "/usr/bin/apt", "update")
	if err != nil {
		return err
	}
	_, err = p.Exec("--", "/usr/bin/apt", "install", "-y", "curl")
	if err != nil {
		return err
	}
	_, err = p.Exec("--", "curl", url)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pod) mkdir(mountPath string) error {
	_, err := p.Exec("--", "mkdir", mountPath+"/"+testDir)
	if err != nil {
		return err
	}
	out, err := p.Exec("--", "ls", mountPath)
	if err != nil {
		return err
	}
	if !strings.Contains(string(out), testDir) {
		return errors.Errorf("Unexpected output from ls: %s", string(out))
	}
	return nil
}

func (p *Pod) powershellMkdir(mountPath string) error {
	out, err := p.Exec("--", "powershell", "mkdir", "-force", mountPath+"\\"+testDir)
	if err != nil {
		return err
	}
	if !strings.Contains(string(out), testDir) {
		return errors.Errorf("Unexpected output from mkdir: %s", string(out))
	}
	out, err = p.Exec("--", "powershell", "ls", mountPath)
	if err != nil {
		return err
	}
	if !strings.Contains(string(out), testDir) {
		return errors.Errorf("Unexpected output from ls: %s", string(out))
	}
	return nil
}

// ExecResult is a return struct for ExecAsync
type ExecResult struct {
	out []byte
	err error
}

// ExecAsync wraps Exec with a struct response for goroutine + channel usage
func (p *Pod) ExecAsync(c ...string) ExecResult {
	out, err := p.Exec(c...)
	return ExecResult{
		out: out,
		err: err,
	}
}

// Exec will execute the given command in the pod
func (p *Pod) Exec(c ...string) ([]byte, error) {
	execCmd := []string{"exec", p.Metadata.Name, "-n", p.Metadata.Namespace}
	execCmd = append(execCmd, c...)
	cmd := exec.Command("k", execCmd...)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl exec':%s\n", string(out))
		log.Printf("Command:kubectl exec %s -n %s %s \n", p.Metadata.Name, p.Metadata.Namespace, c)
		return nil, err
	}
	return out, nil
}

// Delete will delete a Pod in a given namespace
func (p *Pod) Delete(retries int) error {
	var zeroValueDuration time.Duration
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "delete", "po", "-n", p.Metadata.Namespace, p.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd, zeroValueDuration)
		if kubectlError != nil {
			log.Printf("Error while trying to delete Pod %s in namespace %s:%s\n", p.Metadata.Namespace, p.Metadata.Name, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// IsSucceeded returns if a pod is in a Succeeded state
func (p *Pod) IsSucceeded() bool {
	return p.Status.Phase == "Succeeded"
}

// IsFailed returns if a pod is in a Failed state
func (p *Pod) IsFailed() bool {
	return p.Status.Phase == "Failed"
}

// ContainerStatus returns a named pod ContainerStatus object
func (p *Pod) ContainerStatus(name string) (ContainerStatus, error) {
	for _, status := range p.Status.ContainerStatuses {
		if status.Name == name {
			return status, nil
		}
	}
	return ContainerStatus{}, errors.Errorf("no container status object found for name %s in pod %s", name, p.Metadata.Name)
}

func (p *Pod) HasContainer(containerName string) bool {
	for _, c := range p.Spec.Containers {
		if c.Name == containerName {
			return true
		}
	}
	return false
}

// ExitCode returns a ContainerStatus's terminal exit code
func (c *ContainerStatus) ExitCode() int {
	return c.State.Terminated.ExitCode
}

// StartTime returns a terminal ContainerStatus's terminal "started at" RFC3339 time
func (c *ContainerStatus) StartTime() string {
	return c.State.Terminated.StartedAt
}

// EndTime returns a ContainerStatus's terminal "finished at" RFC3339 time
func (c *ContainerStatus) EndTime() string {
	return c.State.Terminated.FinishedAt
}

// CheckOutboundConnection checks outbound connection for a list of pods.
func (l *List) CheckOutboundConnection(sleep, timeout time.Duration, osType api.OSType) (bool, error) {
	type isReady struct {
		pod   Pod
		ready bool
		err   error
	}
	ch := make(chan isReady)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, p := range l.Pods {
		localPod := p
		go func() {
			switch osType {
			case api.Linux:
				ready, err := localPod.CheckLinuxOutboundConnection(sleep, timeout)
				ch <- isReady{
					pod:   localPod,
					ready: ready,
					err:   err,
				}
			case api.Windows:
				ready, err := localPod.CheckWindowsOutboundConnection(sleep, timeout)
				ch <- isReady{
					pod:   localPod,
					ready: ready,
					err:   err,
				}
			}
		}()
	}

	readyCount := 0
	for {
		select {
		case <-ctx.Done():
			return false, errors.Errorf("Timeout exceeded (%s) while waiting for PodList to check outbound internet connection for OS type %s", timeout.String(), osType)
		case response := <-ch:
			ready := response.ready
			err := response.err
			pod := response.pod
			if err == nil {
				if ready {
					readyCount++
				}
			} else {
				err := pod.Logs()
				if err != nil {
					log.Printf("Unable to print pod logs\n: %s", err)
				}
				err = pod.Describe()
				if err != nil {
					log.Printf("Unable to describe pod\n: %s", err)
				}
				return false, errors.Errorf("CheckOutboundConnection returned error for pod %#v: %s\n", pod, err)
			}
			if readyCount == len(l.Pods) {
				return true, nil
			}
		}
	}
}

//ValidateCurlConnection checks curl connection for a list of Linux pods to a specified uri.
func (l *List) ValidateCurlConnection(uri string, sleep, timeout time.Duration) (bool, error) {
	type isReady struct {
		pod   Pod
		ready bool
		err   error
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan isReady)
	for _, p := range l.Pods {
		localPod := p
		go func() {
			ready, err := localPod.ValidateCurlConnection(uri, sleep, timeout)
			ch <- isReady{
				pod:   localPod,
				ready: ready,
				err:   err,
			}
		}()
	}

	readyCount := 0
	for {
		select {
		case <-ctx.Done():
			return false, errors.Errorf("Timeout exceeded (%s) while waiting for PodList to check outbound internet connection", timeout.String())
		case response := <-ch:
			ready := response.ready
			err := response.err
			pod := response.pod
			if err == nil {
				if ready {
					readyCount++
				}
			} else {
				err := pod.Logs()
				if err != nil {
					log.Printf("Unable to print pod logs\n: %s", err)
				}
				err = pod.Describe()
				if err != nil {
					log.Printf("Unable to describe pod\n: %s", err)
				}
				return false, errors.Errorf("ValidateCurlConnection returned error for pod %#v: %s\n", pod, err)
			}
			if readyCount == len(l.Pods) {
				return true, nil
			}
		}
	}
}

// CheckLinuxOutboundConnection will keep retrying the check if an error is received until the timeout occurs or it passes. This helps us when DNS may not be available for some time after a pod starts.
func (p *Pod) CheckLinuxOutboundConnection(sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan error)
	var mostRecentCheckLinuxOutboundConnectionError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- p.attemptOutboundConn()
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case mostRecentCheckLinuxOutboundConnectionError = <-ch:
			if mostRecentCheckLinuxOutboundConnectionError == nil {
				return true, nil
			}
		case <-ctx.Done():
			err := p.Logs()
			if err != nil {
				log.Printf("Unable to print pod logs\n: %s", err)
			}
			err = p.Describe()
			if err != nil {
				log.Printf("Unable to describe pod\n: %s", err)
			}
			return false, errors.Errorf("CheckLinuxOutboundConnection timed out: %s\n", mostRecentCheckLinuxOutboundConnectionError)
		}
	}
}

// ValidateCurlConnection connects to a URI on TCP 80
func (p *Pod) ValidateCurlConnection(uri string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan error)
	var mostRecentValidateCurlConnectionError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- p.curlURL(uri)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case mostRecentValidateCurlConnectionError = <-ch:
			if mostRecentValidateCurlConnectionError == nil {
				return true, nil
			}
		case <-ctx.Done():
			err := p.Logs()
			if err != nil {
				log.Printf("Unable to print pod logs\n: %s", err)
			}
			err = p.Describe()
			if err != nil {
				log.Printf("Unable to describe pod\n: %s", err)
			}
			return false, errors.Errorf("ValidateCurlConnection timed out: %s\n", mostRecentValidateCurlConnectionError)
		}
	}
}

// ValidateOmsAgentLogs validates omsagent logs
func (p *Pod) ValidateOmsAgentLogs(execCmdString string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan ExecResult)
	var mostRecentValidateOmsAgentLogsError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- p.ExecAsync("grep", "-i", execCmdString, "/var/opt/microsoft/omsagent/log/omsagent.log")
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentValidateOmsAgentLogsError = result.err
			if mostRecentValidateOmsAgentLogsError == nil {
				return true, nil
			}
		case <-ctx.Done():
			err := p.Logs()
			if err != nil {
				log.Printf("Unable to print pod logs\n: %s", err)
			}
			err = p.Describe()
			if err != nil {
				log.Printf("Unable to describe pod\n: %s", err)
			}
			return false, errors.Errorf("ValidateOmsAgentLogs timed out: %s\n", mostRecentValidateOmsAgentLogsError)
		}
	}
}

// CheckWindowsOutboundConnection will keep retrying the check if an error is received until the timeout occurs or it passes. This helps us when DNS may not be available for some time after a pod starts.
func (p *Pod) CheckWindowsOutboundConnection(sleep, timeout time.Duration) (bool, error) {
	exp, err := regexp.Compile(`(Connected\s*:\s*True)`)
	if err != nil {
		return false, errors.Errorf("Error while trying to create regex for windows outbound check:%s\n", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan ExecResult)
	var mostRecentCheckWindowsOutboundConnectionError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- p.ExecAsync("--", "powershell", "New-Object", "System.Net.Sockets.TcpClient('8.8.8.8', 443)")
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentCheckWindowsOutboundConnectionError = result.err
			out := result.out
			if mostRecentCheckWindowsOutboundConnectionError == nil {
				if exp.MatchString(string(out)) {
					return true, nil
				}
			}
		case <-ctx.Done():
			err := p.Logs()
			if err != nil {
				log.Printf("Unable to print pod logs\n: %s", err)
			}
			err = p.Describe()
			if err != nil {
				log.Printf("Unable to describe pod\n: %s", err)
			}
			return false, errors.Errorf("CheckWindowsOutboundConnection timed out: %s\n", mostRecentCheckWindowsOutboundConnectionError)
		}
	}
}

// ValidateHostPort will attempt to run curl against the POD's hostIP and hostPort
func (p *Pod) ValidateHostPort(check string, attempts int, sleep time.Duration, master, sshKeyPath string) bool {
	var commandTimeout time.Duration
	hostIP := p.Status.HostIP
	if len(p.Spec.Containers) == 0 || len(p.Spec.Containers[0].Ports) == 0 {
		log.Printf("Unexpected POD container spec: %v. Should have hostPort.\n", p.Spec)
		return false
	}
	hostPort := p.Spec.Containers[0].Ports[0].HostPort

	url := fmt.Sprintf("http://%s:%d", hostIP, hostPort)
	curlCMD := fmt.Sprintf("curl --max-time 60 %s", url)

	for i := 0; i < attempts; i++ {
		cmd := exec.Command("ssh", "-i", sshKeyPath, "-o", "ConnectTimeout=10", "-o", "StrictHostKeyChecking=no", "-o", "UserKnownHostsFile=/dev/null", master, curlCMD)
		out, err := util.RunAndLogCommand(cmd, commandTimeout)
		if err == nil {
			matched, _ := regexp.MatchString(check, string(out))
			if matched {
				return true
			}
		}
		time.Sleep(sleep)
	}
	return false
}

// Logs will get logs from all containers in a pod
func (p *Pod) Logs() error {
	var commandTimeout time.Duration
	for _, container := range p.Spec.Containers {
		cmd := exec.Command("k", "logs", p.Metadata.Name, "-c", container.Name, "-n", p.Metadata.Namespace)
		out, err := util.RunAndLogCommand(cmd, commandTimeout)
		log.Printf("\n%s\n", string(out))
		if err != nil {
			return err
		}
	}
	return nil
}

// Describe will describe a pod resource
func (p *Pod) Describe() error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "describe", "pod", p.Metadata.Name, "-n", p.Metadata.Namespace)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	log.Printf("\n%s\n", string(out))
	return err
}

// ValidateAzureFile will keep retrying the check if azure file is mounted in Pod
func (p *Pod) ValidateAzureFile(mountPath string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan error)
	var mostRecentValidateAzureFileError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- p.powershellMkdir(mountPath)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case mostRecentValidateAzureFileError = <-ch:
			if mostRecentValidateAzureFileError == nil {
				return true, nil
			}
		case <-ctx.Done():
			err := p.Logs()
			if err != nil {
				log.Printf("Unable to print pod logs\n: %s", err)
			}
			err = p.Describe()
			if err != nil {
				log.Printf("Unable to describe pod\n: %s", err)
			}
			return false, errors.Errorf("ValidateAzureFile timed out: %s\n", mostRecentValidateAzureFileError)
		}
	}
}

// ValidatePVC will keep retrying the check if azure disk is mounted in Pod
func (p *Pod) ValidatePVC(mountPath string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan error)
	var mostRecentValidatePVCError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- p.mkdir(mountPath)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case mostRecentValidatePVCError = <-ch:
			if mostRecentValidatePVCError == nil {
				return true, nil
			}
		case <-ctx.Done():
			err := p.Logs()
			if err != nil {
				log.Printf("Unable to print pod logs\n: %s", err)
			}
			err = p.Describe()
			if err != nil {
				log.Printf("Unable to describe pod\n: %s", err)
			}
			return false, errors.Errorf("ValidatePVC timed out: %s\n", mostRecentValidatePVCError)
		}
	}
}

// ValidateResources checks that an addon has the expected memory/cpu limits and requests
func (c *Container) ValidateResources(a api.KubernetesContainerSpec) error {
	expectedCPURequests := a.CPURequests
	expectedCPULimits := a.CPULimits
	expectedMemoryRequests := a.MemoryRequests
	expectedMemoryLimits := a.MemoryLimits
	actualCPURequests := c.getCPURequests()
	actualCPULimits := c.getCPULimits()
	actualMemoryRequests := c.getMemoryRequests()
	actualLimits := c.getMemoryLimits()
	switch {
	case expectedCPURequests != "" && expectedCPURequests != actualCPURequests:
		return errors.Errorf("expected CPU requests %s does not match %s", expectedCPURequests, actualCPURequests)
	case expectedCPULimits != "" && expectedCPULimits != actualCPULimits:
		return errors.Errorf("expected CPU limits %s does not match %s", expectedCPULimits, actualCPULimits)
	case expectedMemoryRequests != "" && expectedMemoryRequests != actualMemoryRequests:
		return errors.Errorf("expected Memory requests %s does not match %s", expectedMemoryRequests, actualMemoryRequests)
	case expectedMemoryLimits != "" && expectedMemoryLimits != actualLimits:
		return errors.Errorf("expected Memory limits %s does not match %s", expectedMemoryLimits, actualLimits)
	default:
		return nil
	}
}

// GetEnvironmentVariable returns an environment variable value from a container within a pod
func (c *Container) GetEnvironmentVariable(varName string) (string, error) {
	for _, envvar := range c.Env {
		if envvar.Name == varName {
			return envvar.Value, nil
		}
	}
	return "", errors.New("environment variable not found")
}

// GetArg returns an arg's value from a container within a pod
func (c *Container) GetArg(argKey string) (string, error) {
	for _, argvar := range c.Args {
		if strings.Contains(argvar, argKey) {
			value := strings.SplitAfter(argvar, "=")[1]
			return value, nil
		}
	}
	return "", errors.New("container argument not found")
}

// getCPURequests returns an the CPU Requests value from a container within a pod
func (c *Container) getCPURequests() string {
	return c.Resources.Requests.CPU
}

// getCPULimits returns an the CPU Requests value from a container within a pod
func (c *Container) getCPULimits() string {
	return c.Resources.Limits.CPU
}

// DashboardtMemoryRequests returns an the CPU Requests value from a container within a pod
func (c *Container) getMemoryRequests() string {
	return c.Resources.Requests.Memory
}

// getMemoryLimits returns an the CPU Requests value from a container within a pod
func (c *Container) getMemoryLimits() string {
	return c.Resources.Limits.Memory
}

// getExternalURLs returns a list of external URLs
func getExternalURLs() []string {
	return []string{"www.bing.com", "google.com"}
}
