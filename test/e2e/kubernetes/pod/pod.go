package pod

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const (
	testDir string = "testdirectory"
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

// CreatePodFromFile will create a Pod from file with a name
func CreatePodFromFile(filename, name, namespace string, sleep, duration time.Duration) (*Pod, error) {
	cmd := exec.Command("kubectl", "apply", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create Pod %s:%s\n", name, string(out))
		return nil, err
	}
	pod, err := GetWithRetry(name, namespace, sleep, duration)
	if err != nil {
		log.Printf("Error while trying to fetch Pod %s:%s\n", name, err)
		return nil, err
	}
	return pod, nil
}

// RunLinuxPod will create a pod that runs a bash command
// --overrides := `"spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}`
func RunLinuxPod(image, name, namespace, command string, printOutput bool, sleep, duration time.Duration) (*Pod, error) {
	overrides := `{ "spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}`
	cmd := exec.Command("kubectl", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--restart=Never", "--overrides", overrides, "--command", "--", "/bin/sh", "-c", command)
	var out []byte
	var err error
	if printOutput {
		out, err = util.RunAndLogCommand(cmd)
	} else {
		out, err = cmd.CombinedOutput()
	}
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	p, err := GetWithRetry(name, namespace, sleep, duration)
	if err != nil {
		log.Printf("Error while trying to fetch Pod %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return p, nil
}

// RunWindowsPod will create a pod that runs a powershell command
// --overrides := `"spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}`
func RunWindowsPod(image, name, namespace, command string, printOutput bool, sleep, duration time.Duration) (*Pod, error) {
	overrides := `{ "spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}`
	cmd := exec.Command("kubectl", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--restart=Never", "--overrides", overrides, "--command", "--", "powershell", command)
	var out []byte
	var err error
	if printOutput {
		out, err = util.RunAndLogCommand(cmd)
	} else {
		out, err = cmd.CombinedOutput()
	}
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	p, err := GetWithRetry(name, namespace, sleep, duration)
	if err != nil {
		log.Printf("Error while trying to fetch Pod %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return p, nil
}

type podRunnerCmd func(string, string, string, string, bool, time.Duration, time.Duration) (*Pod, error)

// RunCommandMultipleTimes runs the same command 'desiredAttempts' times
func RunCommandMultipleTimes(podRunnerCmd podRunnerCmd, image, name, command string, desiredAttempts int, sleep, duration time.Duration) (int, error) {
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
		if i < 1 {
			// Print the first attempt
			p, err = podRunnerCmd(image, podName, "default", command, true, sleep, duration)
		} else {
			p, err = podRunnerCmd(image, podName, "default", command, false, sleep, duration)
		}

		if err != nil {
			return successfulAttempts, err
		}
		succeeded, _ := p.WaitOnSucceeded(sleep, duration)
		cmd := exec.Command("kubectl", "logs", podName, "-n", "default")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Unable to get logs from pod %s\n", podName)
		} else {
			log.Printf("%s\n", string(out[:]))
		}

		err = p.Delete(3)
		if err != nil {
			return successfulAttempts, err
		}

		if succeeded {
			successfulAttempts++
		}
	}

	return successfulAttempts, nil
}

// GetAll will return all pods in a given namespace
func GetAll(namespace string) (*List, error) {
	cmd := exec.Command("kubectl", "get", "pods", "-n", namespace, "-o", "json")
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
func GetWithRetry(podPrefix, namespace string, sleep, duration time.Duration) (*Pod, error) {
	podCh := make(chan *Pod, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pod (%s) in namespace (%s)", duration.String(), podPrefix, namespace)
			default:
				p, err := Get(podPrefix, namespace)
				if err != nil {
					log.Printf("Error getting pod %s in namespace %s: %s\n", podPrefix, namespace, err)
				} else if p != nil {
					podCh <- p
				}
				fmt.Print(".")
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			fmt.Print("\n")
			return nil, err
		case p := <-podCh:
			fmt.Print("\n")
			return p, nil
		}
	}
}

// Get will return a pod with a given name and namespace
func Get(podName, namespace string) (*Pod, error) {
	cmd := exec.Command("kubectl", "get", "pods", podName, "-n", namespace, "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error getting pod:\n")
		util.PrintCommand(cmd)
		return nil, err
	}
	p := Pod{}
	err = json.Unmarshal(out, &p)
	if err != nil {
		log.Printf("Error unmarshalling pods json:%s\n", err)
		return nil, err
	}
	return &p, nil
}

// GetTerminated will return a pod with a given name and namespace, including terminated pods
func GetTerminated(podName, namespace string) (*Pod, error) {
	cmd := exec.Command("kubectl", "get", "pods", podName, "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	p := Pod{}
	err = json.Unmarshal(out, &p)
	if err != nil {
		log.Printf("Error unmarshalling pods json:%s\n", err)
		return nil, err
	}
	return &p, nil
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

// AreAllPodsRunning will return true if all pods in a given namespace are in a Running State
func AreAllPodsRunning(podPrefix, namespace string) (bool, error) {
	pl, err := GetAll(namespace)
	if err != nil {
		return false, err
	}

	var status []bool
	for _, pod := range pl.Pods {
		matched, err := regexp.MatchString(podPrefix, pod.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return false, err
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

	return true, nil
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
			if pod.Status.Phase == "Failed" {
				return false, true, nil
			}
			if pod.Status.Phase != "Succeeded" {
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

// WaitOnReady is used when you dont have a handle on a pod but want to wait until its in a Ready state.
// successesNeeded is used to make sure we return the correct value even if the pod is in a CrashLoop
func WaitOnReady(podPrefix, namespace string, successesNeeded int, sleep, duration time.Duration) (bool, error) {
	successCount := 0
	failureCount := 0
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pods (%s) to become ready in namespace (%s), got %d of %d required successful pods ready results", duration.String(), podPrefix, namespace, successCount, successesNeeded)
			default:
				ready, err := AreAllPodsRunning(podPrefix, namespace)
				if err != nil {
					errCh <- err
					return
				}
				if ready {
					successCount = successCount + 1
					if successCount >= successesNeeded {
						readyCh <- true
					}
				} else {
					if successCount > 1 {
						failureCount = failureCount + 1
						if failureCount >= successesNeeded {
							errCh <- errors.Errorf("Pods from deployment (%s) in namespace (%s) have been checked out as all Ready %d times, but NotReady %d times. This behavior may mean it is in a crashloop", podPrefix, namespace, successCount, failureCount)
						}
					}
					time.Sleep(sleep)
				}
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// WaitOnSucceeded is used when you dont have a handle on a pod but want to wait until its in a Succeeded state.
func WaitOnSucceeded(podPrefix, namespace string, sleep, duration time.Duration) (bool, error) {
	succeededCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pods (%s) to succeed in namespace (%s)", duration.String(), podPrefix, namespace)
			default:
				succeeded, failed, err := AreAllPodsSucceeded(podPrefix, namespace)
				if err != nil {
					errCh <- err
					return
				}
				if failed {
					errCh <- errors.New("At least one pod in a Failed state")
				}
				if succeeded {
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
		case ready := <-succeededCh:
			return ready, nil
		}
	}
}

// WaitOnReady will call the static method WaitOnReady passing in p.Metadata.Name and p.Metadata.Namespace
func (p *Pod) WaitOnReady(sleep, duration time.Duration) (bool, error) {
	return WaitOnReady(p.Metadata.Name, p.Metadata.Namespace, 2, sleep, duration)
}

// WaitOnSucceeded will call the static method WaitOnSucceeded passing in p.Metadata.Name and p.Metadata.Namespace
func (p *Pod) WaitOnSucceeded(sleep, duration time.Duration) (bool, error) {
	return WaitOnSucceeded(p.Metadata.Name, p.Metadata.Namespace, sleep, duration)
}

// Exec will execute the given command in the pod
func (p *Pod) Exec(c ...string) ([]byte, error) {
	execCmd := []string{"exec", p.Metadata.Name, "-n", p.Metadata.Namespace}
	execCmd = append(execCmd, c...)
	cmd := exec.Command("kubectl", execCmd...)
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
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("kubectl", "delete", "po", "-n", p.Metadata.Namespace, p.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd)
		if kubectlError != nil {
			log.Printf("Error while trying to delete Pod %s in namespace %s:%s\n", p.Metadata.Namespace, p.Metadata.Name, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// CheckLinuxOutboundConnection will keep retrying the check if an error is received until the timeout occurs or it passes. This helps us when DNS may not be available for some time after a pod starts.
func (p *Pod) CheckLinuxOutboundConnection(sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	var installedCurl bool
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pod (%s) to check outbound internet connection", duration.String(), p.Metadata.Name)
			default:
				if !installedCurl {
					_, err := p.Exec("--", "/usr/bin/apt", "update")
					if err != nil {
						break
					}
					_, err = p.Exec("--", "/usr/bin/apt", "install", "-y", "curl")
					if err != nil {
						break
					}
					installedCurl = true
				}
				// if we can curl bing.com we have outbound internet access
				out, err := p.Exec("--", "curl", "bing.com")
				if err == nil {
					readyCh <- true
				} else {
					// in case bing.com is down let's hope google.com is also not down
					_, err := p.Exec("--", "curl", "google.com")
					if err == nil {
						readyCh <- true
					} else {
						// if both bing.com and google.com are down let's say we don't have outbound internet access
						log.Printf("Error:%s\n", err)
						log.Printf("Out:%s\n", out)
					}
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// ValidateCurlConnection connects to a URI on TCP 80
func (p *Pod) ValidateCurlConnection(uri string, sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	var installedCurl bool
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pod (%s) to curl uri %s", duration.String(), p.Metadata.Name, uri)
			default:
				if !installedCurl {
					_, err := p.Exec("--", "/usr/bin/apt", "update")
					if err != nil {
						break
					}
					_, err = p.Exec("--", "/usr/bin/apt", "install", "-y", "curl")
					if err != nil {
						break
					}
					installedCurl = true
				}
				_, err := p.Exec("--", "curl", uri)
				if err == nil {
					readyCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// ValidateOmsAgentLogs validates omsagent logs
func (p *Pod) ValidateOmsAgentLogs(execCmdString string, sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for logs to be written by omsagent", duration.String())
			default:
				_, err := p.Exec("grep", "-i", execCmdString, "/var/opt/microsoft/omsagent/log/omsagent.log")
				if err == nil {
					readyCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// CheckWindowsOutboundConnection will keep retrying the check if an error is received until the timeout occurs or it passes. This helps us when DNS may not be available for some time after a pod starts.
func (p *Pod) CheckWindowsOutboundConnection(url string, sleep, duration time.Duration) (bool, error) {
	exp, err := regexp.Compile(`(StatusCode\s*:\s*200)`)
	if err != nil {
		log.Printf("Error while trying to create regex for windows outbound check:%s\n", err)
		return false, err
	}
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pod (%s) to check outbound internet connection", duration.String(), p.Metadata.Name)
			default:
				out, err := p.Exec("--", "powershell", "iwr", "-UseBasicParsing", "-TimeoutSec", "60", url)
				if err == nil {
					matched := exp.MatchString(string(out))
					if matched {
						readyCh <- true
					} else {
						readyCh <- false
					}
				} else {
					log.Printf("Error:%s\n", err)
					log.Printf("Out:%s\n", out)
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// ValidateHostPort will attempt to run curl against the POD's hostIP and hostPort
func (p *Pod) ValidateHostPort(check string, attempts int, sleep time.Duration, master, sshKeyPath string) bool {
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
		out, err := util.RunAndLogCommand(cmd)
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

// ValidateAzureFile will keep retrying the check if azure file is mounted in Pod
func (p *Pod) ValidateAzureFile(mountPath string, sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pod (%s) to check azure file mounted", duration.String(), p.Metadata.Name)
			default:
				out, err := p.Exec("--", "powershell", "mkdir", "-force", mountPath+"\\"+testDir)
				if err == nil && strings.Contains(string(out), testDir) {
					out, err := p.Exec("--", "powershell", "ls", mountPath)
					if err == nil && strings.Contains(string(out), testDir) {
						readyCh <- true
					} else {
						log.Printf("Error:%s\n", err)
						log.Printf("Out:%s\n", out)
					}
				} else {
					log.Printf("Error:%s\n", err)
					log.Printf("Out:%s\n", out)
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// ValidatePVC will keep retrying the check if azure disk is mounted in Pod
func (p *Pod) ValidatePVC(mountPath string, sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Pod (%s) to check azure disk mounted", duration.String(), p.Metadata.Name)
			default:
				out, err := p.Exec("--", "mkdir", mountPath+"/"+testDir)
				if err == nil {
					out, err := p.Exec("--", "ls", mountPath)
					if err == nil && strings.Contains(string(out), testDir) {
						readyCh <- true
					} else {
						log.Printf("Error:%s\n", err)
						log.Printf("Out:%s\n", out)
					}
				} else {
					log.Printf("Error:%s\n", err)
					log.Printf("Out:%s\n", out)
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
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
	if expectedCPURequests != "" && expectedCPURequests != actualCPURequests {
		return errors.Errorf("expected CPU requests %s does not match %s", expectedCPURequests, actualCPURequests)
	} else if expectedCPULimits != "" && expectedCPULimits != actualCPULimits {
		return errors.Errorf("expected CPU limits %s does not match %s", expectedCPULimits, actualCPULimits)
	} else if expectedMemoryRequests != "" && expectedMemoryRequests != actualMemoryRequests {
		return errors.Errorf("expected Memory requests %s does not match %s", expectedMemoryRequests, actualMemoryRequests)
	} else if expectedMemoryLimits != "" && expectedMemoryLimits != actualLimits {
		return errors.Errorf("expected Memory limits %s does not match %s", expectedMemoryLimits, actualLimits)
	} else {
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
