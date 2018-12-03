// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/pkg/errors"
)

func printCmd(cmd *exec.Cmd) {
	fmt.Printf("\n$ %s\n", strings.Join(cmd.Args, " "))
}

// ApplyFromTemplate processes and creates the provided templateName/templateNamespace template
// in the provided namespace.
func ApplyFromTemplate(templateName, templateNamespace, namespace string) error {
	processCmd := exec.Command("oc", "process", templateName, "-n", templateNamespace)
	printCmd(processCmd)
	out, err := processCmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "cannot process template %s\noutput: %s", templateName, string(out))
	}
	if err := ioutil.WriteFile(templateName, out, 0644); err != nil {
		return errors.Wrapf(err, "cannot create tempfile for processed template %s", templateName)
	}
	defer os.Remove(templateName)
	createCmd := exec.Command("oc", "apply", "-n", namespace, "-f", templateName)
	printCmd(createCmd)
	out, err = createCmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "cannot apply processed template %s\noutput: %s", templateName, string(out))
	}
	return nil
}

// WaitForDeploymentConfig waits until the provided deploymentconfig namespace/name
// gets deployed.
func WaitForDeploymentConfig(name, namespace string) error {
	cmd := exec.Command("oc", "rollout", "status", fmt.Sprintf("dc/%s", name), "-n", namespace)
	printCmd(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to see the rollout status of dc/%s: %s", name, string(out))
		return err
	}
	return nil
}

// GetHost expects the name and namespace of a route in order to
// return its host.
func GetHost(name, namespace string) (string, error) {
	cmd := exec.Command("oc", "get", fmt.Sprintf("route/%s", name), "-n", namespace, "-o", "jsonpath={.spec.host}")
	printCmd(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to get the hostname of route/%s: %s", name, string(out))
		return "", err
	}
	return string(out), nil
}

// TestHost tries to access host and retries maxRetries times with a retryDelay
// that is doubled on every retry.
func TestHost(host string, maxRetries int, retryDelay time.Duration) error {
	backoff := retryDelay
	url := fmt.Sprintf("http://%s", host)

	resp, err := http.Get(url)
	if err == nil && resp.StatusCode == http.StatusOK {
		return nil
	}
	if err == nil {
		log.Printf("got status %q while trying to access %s", resp.Status, host)
		resp.Body.Close()
	} else {
		log.Printf("error while trying to access %s: %v", host, err)
	}
	for retries := 1; retries <= maxRetries; retries++ {
		log.Printf("Retry #%d to access %s", retries, host)
		resp, err = http.Get(url)
		if err != nil {
			log.Printf("error while trying to access %s: %v", host, err)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Printf("got status %q while trying to access %s", resp.Status, host)
		log.Printf("sleeping for %fs", backoff.Seconds())
		time.Sleep(backoff)
		backoff *= 2
	}
	if err != nil {
		return err
	}
	return errors.Errorf("unexpected response status: %v", resp.Status)
}

// DumpNodes dumps information about nodes.
func DumpNodes() (string, error) {
	cmd := exec.Command("oc", "get", "nodes", "-o", "wide")
	printCmd(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to list nodes: %s", string(out))
		return "", err
	}
	return string(out), nil
}

// DumpPods dumps the pods from all namespaces.
func DumpPods() (string, error) {
	cmd := exec.Command("oc", "get", "pods", "--all-namespaces", "-o", "wide")
	printCmd(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to list pods from all namespaces: %s", string(out))
		return "", err
	}
	return string(out), nil
}

// FetchLogs returns logs for the provided kind/name in namespace.
func FetchLogs(kind, namespace, name string) string {
	cmd := exec.Command("oc", "logs", fmt.Sprintf("%s/%s", kind, name), "-n", namespace)
	printCmd(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error trying to fetch logs from %s/%s in %s: %s", kind, name, namespace, string(out))
	}
	return string(out)
}

// FetchClusterInfo returns node and pod information about the cluster.
func FetchClusterInfo(logPath string) error {
	needsLog := map[string]func() (string, error){
		"node-info": DumpNodes,
		"pod-info":  DumpPods,
	}

	var errs []error
	for base, logFn := range needsLog {
		logs, err := logFn()
		if err != nil {
			errs = append(errs, err)
			continue
		}
		path := filepath.Join(logPath, base)
		if err := ioutil.WriteFile(path, []byte(logs), 0644); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

// FetchOpenShiftLogs returns logs for all OpenShift components
// (control plane and infra).
func FetchOpenShiftLogs(distro, version, sshKeyPath, adminName, name, location, logPath string) error {
	if err := fetchControlPlaneLogs(distro, version, sshKeyPath, adminName, name, location, logPath); err != nil {
		return errors.Wrap(err, "cannot fetch logs for control plane components: %v")
	}
	if err := fetchInfraLogs(logPath); err != nil {
		return errors.Wrap(err, "cannot fetch logs for infra components: %v")
	}
	return nil
}

// fetchControlPlaneLogs returns logs for Openshift control plane components.
func fetchControlPlaneLogs(distro, version, sshKeyPath, adminName, name, location, logPath string) error {
	sshAddress := fmt.Sprintf("%s@%s.%s.cloudapp.azure.com", adminName, name, location)

	switch version {
	case common.OpenShiftVersion3Dot9Dot0:
		return fetch39ControlPlaneLogs(distro, sshKeyPath, sshAddress, logPath)
	case common.OpenShiftVersionUnstable:
		return fetchUnstableControlPlaneLogs(distro, sshKeyPath, sshAddress, name, logPath)
	default:
		return errors.Errorf("invalid OpenShift version %q - won't gather logs from the control plane", version)
	}
}

func fetch39ControlPlaneLogs(distro, sshKeyPath, sshAddress, logPath string) error {
	var errs []error
	for _, service := range getSystemdServices(distro) {
		cmdToExec := fmt.Sprintf("sudo journalctl -u %s.service", service)
		out := remoteExec(sshKeyPath, sshAddress, cmdToExec)
		path := filepath.Join(logPath, service)
		if err := ioutil.WriteFile(path, out, 0644); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func getSystemdServices(distro string) []string {
	services := []string{"etcd"}
	switch api.Distro(distro) {
	case api.OpenShift39RHEL:
		services = append(services, "atomic-openshift-master-api", "atomic-openshift-master-controllers", "atomic-openshift-node")
	case api.OpenShiftCentOS:
		services = append(services, "origin-master-api", "origin-master-controllers", "origin-node")
	default:
		log.Printf("Will not gather journal for the control plane because invalid OpenShift distro was specified: %q", distro)
	}
	return services
}

func remoteExec(sshKeyPath, sshAddress, cmdToExec string) []byte {
	cmd := exec.Command("ssh", "-i", sshKeyPath, "-o", "ConnectTimeout=10", "-o", "StrictHostKeyChecking=no", "-o", "UserKnownHostsFile=/dev/null", sshAddress, cmdToExec)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Cannot execute remote command %q: %v", cmdToExec, err)
	}
	return out
}

type resource struct {
	kind      string
	namespace string
	name      string
}

func (r resource) String() string {
	return fmt.Sprintf("%s_%s_%s", r.namespace, r.kind, r.name)
}

// TODO: Promote to 3.10 when the time comes
func fetchUnstableControlPlaneLogs(distro, sshKeyPath, sshAddress, name, logPath string) error {
	controlPlane := []resource{
		{kind: "pod", namespace: "kube-system", name: fmt.Sprintf("master-api-ocp-master-%s-0", name)},
		{kind: "pod", namespace: "kube-system", name: fmt.Sprintf("master-controllers-ocp-master-%s-0", name)},
		{kind: "pod", namespace: "kube-system", name: fmt.Sprintf("master-etcd-ocp-master-%s-0", name)},
	}

	var errs []error
	for _, r := range controlPlane {
		log := FetchLogs(r.kind, r.namespace, r.name)
		path := filepath.Join(logPath, r.name)
		if err := ioutil.WriteFile(path, []byte(log), 0644); err != nil {
			errs = append(errs, err)
		}
	}

	for _, service := range getSystemdServices(distro) {
		// 3.10+ deployments run only the node process as a systemd service
		if service != "atomic-openshift-node" && service != "origin-node" {
			continue
		}
		cmdToExec := fmt.Sprintf("sudo journalctl -u %s.service", service)
		out := remoteExec(sshKeyPath, sshAddress, cmdToExec)
		path := filepath.Join(logPath, service)
		if err := ioutil.WriteFile(path, out, 0644); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

// fetchInfraLogs returns logs for Openshift infra components.
// TODO: Eventually we may need to version this too.
func fetchInfraLogs(logPath string) error {
	infraResources := []resource{
		// TODO: Maybe collapse this list and the actual readiness check tests
		// in openshift e2e.
		{kind: "deploymentconfig", namespace: "default", name: "router"},
		{kind: "deploymentconfig", namespace: "default", name: "docker-registry"},
		{kind: "deploymentconfig", namespace: "default", name: "registry-console"},
		{kind: "statefulset", namespace: "openshift-infra", name: "bootstrap-autoapprover"},
		{kind: "statefulset", namespace: "openshift-metrics", name: "prometheus"},
		{kind: "daemonset", namespace: "kube-service-catalog", name: "apiserver"},
		{kind: "daemonset", namespace: "kube-service-catalog", name: "controller-manager"},
		{kind: "deploymentconfig", namespace: "openshift-ansible-service-broker", name: "asb"},
		{kind: "deploymentconfig", namespace: "openshift-ansible-service-broker", name: "asb-etcd"},
		{kind: "daemonset", namespace: "openshift-template-service-broker", name: "apiserver"},
		{kind: "deployment", namespace: "openshift-web-console", name: "webconsole"},
	}

	var errs []error
	for _, r := range infraResources {
		log := FetchLogs(r.kind, r.namespace, r.name)
		path := filepath.Join(logPath, "infra-"+r.String())
		err := ioutil.WriteFile(path, []byte(log), 0644)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

// FetchOpenShiftMetrics gathers metrics from etcd and the control plane.
func FetchOpenShiftMetrics(logPath string) error {
	var errs []error

	// api server metrics
	cmd := exec.Command("oc", "get", "--raw", "https://localhost:8443/metrics")
	printCmd(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errs = append(errs, errors.Wrap(err, "cannot get api server metrics"))
	} else {
		path := filepath.Join(logPath, "api-server-metrics")
		err := ioutil.WriteFile(path, []byte(out), 0644)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "cannot write api server metrics: "))
		}
	}

	// controller manager metrics
	cmd = exec.Command("oc", "get", "--raw", "https://localhost:8444/metrics")
	printCmd(cmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		errs = append(errs, errors.Wrap(err, "cannot get controller manager metrics"))
	} else {
		path := filepath.Join(logPath, "controller-manager-metrics")
		err := ioutil.WriteFile(path, []byte(out), 0644)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "cannot write controller manager metrics"))
		}
	}

	// etcd metrics
	cmd = exec.Command("oc", "get", "--raw", "https://localhost:2380/metrics")
	printCmd(cmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		errs = append(errs, errors.Wrap(err, "cannot get etcd metrics"))
	} else {
		path := filepath.Join(logPath, "etcd-metrics")
		err := ioutil.WriteFile(path, []byte(out), 0644)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "cannot write etcd metrics"))
		}
	}

	return kerrors.NewAggregate(errs)
}

// FetchWaagentLogs returns stdout and stderr from waagent.
func FetchWaagentLogs(sshKeyPath, adminName, name, location, logPath string) error {
	sshAddress := fmt.Sprintf("%s@%s.%s.cloudapp.azure.com", adminName, name, location)

	paths := []string{
		"/var/lib/waagent/custom-script/download/0/stderr",
		"/var/lib/waagent/custom-script/download/0/stdout",
	}

	var errs []error
	for _, path := range paths {
		cmdToExec := fmt.Sprintf("sudo cat %s", path)
		out := remoteExec(sshKeyPath, sshAddress, cmdToExec)
		logPath := filepath.Join(logPath, fmt.Sprintf("waagent-%s", filepath.Base(path)))
		if err := ioutil.WriteFile(logPath, out, 0644); err != nil {
			errs = append(errs, errors.Wrapf(err, "Cannot write to path %s", logPath))
		}
	}

	return kerrors.NewAggregate(errs)
}
