// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import v1 "k8s.io/api/core/v1"

// IsNodeReady returns true if the NodeReady condition of node is set to true.
//
// Copy of https://github.com/kubernetes/kubernetes/blob/886e04f1fffbb04faf8a9f9ee141143b2684ae68/pkg/api/v1/node/util.go#L40
func IsNodeReady(node *v1.Node) bool {
	for _, c := range node.Status.Conditions {
		if c.Type == v1.NodeReady {
			return c.Status == v1.ConditionTrue
		}
	}
	return false
}

func IsAnyContainerCrashing(pod v1.Pod) bool {
	for _, c := range pod.Status.ContainerStatuses {
		if pod.Status.Phase == v1.PodRunning && IsContainerCrashing(c) {
			return true
		}
	}
	return false
}

func IsContainerCrashing(cs v1.ContainerStatus) bool {
	if cs.State.Waiting != nil && cs.State.Waiting.Reason == "CrashLoopBackOff" {
		return true
	}
	return false
}

func IsMirrorPod(pod v1.Pod) bool {
	for k := range pod.ObjectMeta.Annotations {
		if k == v1.MirrorPodAnnotationKey {
			return true
		}
	}
	return false
}
