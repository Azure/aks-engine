// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/helpers/ssh"
	"github.com/Azure/aks-engine/pkg/kubernetes"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

type restartPodCondition func(v1.Pod) bool

func healPodsFunc(ctx context.Context, nodes map[string]*ssh.RemoteHost, heal func(v1.Pod, *ssh.RemoteHost), conditions ...restartPodCondition) func(*v1.PodList) {
	podStream := make(chan v1.Pod)
	go func() {
		defer close(podStream)
		for {
			select {
			case pod := <-podStream:
				if node, ok := nodes[pod.Spec.NodeName]; ok {
					heal(pod, node)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return func(pl *v1.PodList) {
		for _, pli := range pl.Items {
			needsRestart := true
			for _, cond := range conditions {
				needsRestart = needsRestart && cond(pli)
			}
			if needsRestart {
				podStream <- pli
			}
		}
	}
}

func restartContainers(pod v1.Pod, node *ssh.RemoteHost) {
	log.Debugf("Healing pod: %s", pod.Name)
	for _, cs := range pod.Status.ContainerStatuses {
		if !kubernetes.IsContainerCrashing(cs) {
			return
		}
		containerID := strings.Split(cs.ContainerID, "://")
		cri := containerID[0]
		arg := containerID[1]
		if len(containerID) != 2 || (cri != api.Docker && cri != api.Containerd) {
			log.Debugf("Restart aborted. Unexpected pod.Status.ContainerStatuses.ContainerID value: %s", cs.ContainerID)
			return
		}
		if cri == api.Containerd {
			for _, c := range pod.Spec.Containers {
				if c.Name == cs.Name {
					arg = c.Command[0]
					if arg == "/hyperkube" {
						arg = c.Command[1]
					}
				}
			}
		}
		// this recovers mirror pods (apiserver mostly) in CrashLoopBackOff from error:
		// "Error: failed to create listener: failed to listen on 0.0.0.0:443: listen tcp 0.0.0.0:443: bind: address already in use"
		log.Debugf("Healing container: %s", cs.Name)
		script := fmt.Sprintf("bash -euxo pipefail -c \"sudo /etc/kubernetes/rotate-certs/rotate-certs.sh restart_mirror_pod_%s %s\"", cri, arg)
		out, err := ssh.ExecuteRemote(node, script)
		if err != nil {
			log.Debugf("Remote command output: %s", out)
		}
	}
}
