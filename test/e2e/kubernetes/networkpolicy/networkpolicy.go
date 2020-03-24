//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package networkpolicy

import (
	"log"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/deployment"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/pod"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	. "github.com/onsi/gomega"
)

// CreateNetworkPolicyFromFile will create a NetworkPolicy from file with a name
func CreateNetworkPolicyFromFile(filename, name, namespace string) error {
	cmd := exec.Command("k", "create", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create NetworkPolicy %s:%s\n", name, string(out))
		return err
	}
	return nil
}

// DeleteNetworkPolicy will create a NetworkPolicy from file with a name
func DeleteNetworkPolicy(name, namespace string) {
	cmd := exec.Command("k", "delete", "networkpolicy", "-n", namespace, name)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to delete NetworkPolicy %s in namespace %s:%s\n", name, namespace, string(out))
	}
	Expect(err).NotTo(HaveOccurred())
}

func EnsureRunningPodExists(deploymentName string, namespace string, successesNeeded int, sleepTime time.Duration, timeout time.Duration) {
	running, err := pod.WaitOnSuccesses(deploymentName, namespace, 4, sleepTime, timeout)
	Expect(err).NotTo(HaveOccurred())
	Expect(running).To(Equal(true))
}

func GetRunningPodsFromDeployment(dep *deployment.Deployment) []pod.Pod {
	podsRunning, err := dep.PodsRunning()
	Expect(err).NotTo(HaveOccurred())
	Expect(len(podsRunning)).ToNot(BeZero())
	return podsRunning
}

func EnsureOutboundInternetAccess(pods []pod.Pod, cfg config.Config) {
	pl := pod.List{Pods: pods}
	pass, err := pl.CheckOutboundConnection(5*time.Second, cfg.Timeout, api.Linux)
	Expect(err).NotTo(HaveOccurred())
	Expect(pass).To(BeTrue())
}

func EnsureConnectivityResultBetweenPods(fromPods []pod.Pod, toPods []pod.Pod, timeout time.Duration, true bool) {
	pl := pod.List{Pods: fromPods}
	for _, toPod := range toPods {
		pass, err := pl.ValidateCurlConnection(toPod.Status.PodIP, 30*time.Second, timeout)
		if true {
			if err != nil {
				e := toPod.Describe()
				if e != nil {
					log.Printf("Unable to describe pod %s\n: %s", toPod.Metadata.Name, e)
				}
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(pass).To(BeTrue())
		} else {
			Expect(err).Should(HaveOccurred())
			Expect(pass).To(BeFalse())
		}
	}
}

func ApplyNetworkPolicy(nwpolicyName string, namespace string, nwpolicyFileName string, policyDir string) {
	err := CreateNetworkPolicyFromFile(filepath.Join(policyDir, nwpolicyFileName), nwpolicyName, namespace)
	Expect(err).NotTo(HaveOccurred())
}
