package deployment

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/pod"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
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
	var cmd *exec.Cmd
	overrides := `{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}}}`
	if miscOpts != "" {
		cmd = exec.Command("kubectl", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--overrides", overrides, miscOpts)
	} else {
		cmd = exec.Command("kubectl", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--overrides", overrides)
	}
	out, err := util.RunAndLogCommand(cmd)
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	d, err := Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch Deployment %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return d, nil
}

// CreateLinuxDeployIfNotExist first checks if a deployment already exists, and return it if so
// If not, we call CreateLinuxDeploy
func CreateLinuxDeployIfNotExist(image, name, namespace, miscOpts string) (*Deployment, error) {
	deployment, err := Get(name, namespace)
	if err != nil {
		return CreateLinuxDeploy(image, name, namespace, miscOpts)
	}
	return deployment, nil
}

// RunLinuxDeploy will create a deployment that runs a bash command in a pod
// --overrides=' "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}}}'
func RunLinuxDeploy(image, name, namespace, command string, replicas int) (*Deployment, error) {
	overrides := `{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"linux"}}}}}`
	cmd := exec.Command("kubectl", "run", name, "-n", namespace, "--image", image, "--image-pull-policy=IfNotPresent", "--replicas", strconv.Itoa(replicas), "--overrides", overrides, "--command", "--", "/bin/sh", "-c", command)
	out, err := util.RunAndLogCommand(cmd)
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	d, err := Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch Deployment %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return d, nil
}

// CreateWindowsDeploy will crete a deployment for a given image with a name in a namespace
func CreateWindowsDeploy(image, name, namespace string, port int, hostport int) (*Deployment, error) {
	overrides := `{ "spec":{"template":{"spec": {"nodeSelector":{"beta.kubernetes.io/os":"windows"}}}}}`
	cmd := exec.Command("kubectl", "run", name, "-n", namespace, "--image", image, "--port", strconv.Itoa(port), "--hostport", strconv.Itoa(hostport), "--overrides", overrides)
	out, err := util.RunAndLogCommand(cmd)
	if err != nil {
		log.Printf("Error trying to deploy %s [%s] in namespace %s:%s\n", name, image, namespace, string(out))
		return nil, err
	}
	d, err := Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch Deployment %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return d, nil
}

// Get returns a deployment from a name and namespace
func Get(name, namespace string) (*Deployment, error) {
	cmd := exec.Command("kubectl", "get", "deploy", "-o", "json", "-n", namespace, name)
	out, err := util.RunAndLogCommand(cmd)
	if err != nil {
		log.Printf("Error while trying to fetch deployment %s in namespace %s:%s\n", name, namespace, string(out))
		return nil, err
	}
	d := Deployment{}
	err = json.Unmarshal(out, &d)
	if err != nil {
		log.Printf("Error while trying to unmarshal deployment json:%s\n%s\n", err, string(out))
		return nil, err
	}
	return &d, nil
}

// Delete will delete a deployment in a given namespace
func (d *Deployment) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("kubectl", "delete", "deploy", "-n", d.Metadata.Namespace, d.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd)
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
			cmd := exec.Command("kubectl", "delete", "hpa", "-n", d.Metadata.Namespace, d.Metadata.Name)
			kubectlOutput, kubectlError = util.RunAndLogCommand(cmd)
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
	cmd := exec.Command("kubectl", "expose", "deployment", d.Metadata.Name, "--type", svcType, "-n", d.Metadata.Namespace, "--target-port", strconv.Itoa(targetPort), "--port", strconv.Itoa(exposedPort))
	out, err := util.RunAndLogCommand(cmd)
	if err != nil {
		log.Printf("Error while trying to expose (%s) target port (%v) for deployment %s in namespace %s on port %v:%s\n", svcType, targetPort, d.Metadata.Name, d.Metadata.Namespace, exposedPort, string(out))
		return err
	}
	return nil
}

// CreateDeploymentHPA applies autoscale characteristics to deployment
func (d *Deployment) CreateDeploymentHPA(cpuPercent, min, max int) error {
	cmd := exec.Command("kubectl", "autoscale", "deployment", d.Metadata.Name, fmt.Sprintf("--cpu-percent=%d", cpuPercent),
		fmt.Sprintf("--min=%d", min), fmt.Sprintf("--max=%d", max))
	out, err := util.RunAndLogCommand(cmd)
	if err != nil {
		log.Printf("Error while configuring autoscale against deployment %s:%s\n", d.Metadata.Name, string(out))
		return err
	}
	d.Metadata.HasHPA = true
	return nil
}

// Pods will return all pods related to a deployment
func (d *Deployment) Pods() ([]pod.Pod, error) {
	return pod.GetAllByPrefix(d.Metadata.Name, d.Metadata.Namespace)
}

// WaitForReplicas waits for a pod replica count between min and max
func (d *Deployment) WaitForReplicas(min, max int, sleep, duration time.Duration) ([]pod.Pod, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	var pods []pod.Pod
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for minimum %d and maximum %d Pod replicas from Deployment %s", duration.String(), min, max, d.Metadata.Name)
			default:
				pods, err := pod.GetAllByPrefix(d.Metadata.Name, d.Metadata.Namespace)
				if err != nil {
					errCh <- err
					return
				}
				if min == -1 {
					if len(pods) <= max {
						readyCh <- true
					}
				} else if max == -1 {
					if len(pods) >= min {
						readyCh <- true
					}
				} else {
					if len(pods) >= min && len(pods) <= max {
						readyCh <- true
					}
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return pods, err
		case <-readyCh:
			return pods, nil
		}
	}
}
