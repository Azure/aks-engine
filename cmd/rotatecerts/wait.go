// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	"time"

	"github.com/Azure/aks-engine/pkg/helpers/ssh"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

// WaitForNodesReady returns true if all the cluster nodes reached the Ready state
func WaitForNodesReady(c Client, nodes []*ssh.RemoteHost, successesNeeded int, interval, timeout time.Duration) error {
	isNodeReady := func(node v1.Node) bool {
		for _, c := range node.Status.Conditions {
			if c.Type == v1.NodeReady {
				return c.Status == v1.ConditionTrue
			}
		}
		return false
	}

	successesCount := 0
	nodeReady := make(map[string]bool, len(nodes))
	err := wait.PollImmediate(interval, timeout, func() (bool, error) {
		nl, err := c.ListNodes()
		if err != nil {
			return false, err
		}
		for _, n := range nodes {
			nodeReady[n.URI] = false
		}
		for _, nli := range nl.Items {
			if _, ok := nodeReady[nli.ObjectMeta.Name]; ok {
				ready := isNodeReady(nli)
				if !ready {
					return false, nil
				}
				nodeReady[nli.ObjectMeta.Name] = ready
			}
		}
		for _, ready := range nodeReady {
			if !ready {
				return false, nil
			}
		}
		successesCount++
		if successesCount < successesNeeded {
			return false, nil
		}
		return true, nil
	})
	return err
}

// WaitForAllInNamespaceReady returns true if all containers in a given namespace reached the Ready state
func WaitForAllInNamespaceReady(c Client, namespace string, successesNeeded int, interval, timeout time.Duration) error {
	successesCount := 0
	err := wait.PollImmediate(interval, timeout, func() (bool, error) {
		pl, err := c.ListPods(namespace, metav1.ListOptions{})
		if err != nil {
			return false, err
		}
		for _, pli := range pl.Items {
			ready := pli.Status.Phase == v1.PodRunning
			for _, c := range pli.Status.ContainerStatuses {
				ready = ready && c.State.Running != nil && c.Ready
			}
			if !ready {
				return false, nil
			}
		}
		successesCount++
		if successesCount < successesNeeded {
			return false, nil
		}
		return true, nil
	})
	return err
}

// WaitForReady returns true if all containers in a given pod list reached the Ready state
func WaitForReady(c Client, namespace string, pods []string, successesNeeded int, interval, timeout time.Duration) error {
	successesCount := 0
	podReady := make(map[string]bool, len(pods))
	err := wait.PollImmediate(interval, timeout, func() (bool, error) {
		pl, err := c.ListPods(namespace, metav1.ListOptions{})
		if err != nil {
			return false, err
		}
		for _, n := range pods {
			podReady[n] = false
		}
		for _, pli := range pl.Items {
			if _, ok := podReady[pli.ObjectMeta.Name]; ok {
				ready := pli.Status.Phase == v1.PodRunning
				for _, c := range pli.Status.ContainerStatuses {
					ready = ready && c.State.Running != nil && c.Ready
				}
				if !ready {
					return false, nil
				}
				podReady[pli.ObjectMeta.Name] = ready
			}
		}
		for _, ready := range podReady {
			if !ready {
				return false, nil
			}
		}
		successesCount++
		if successesCount < successesNeeded {
			return false, nil
		}
		return true, nil
	})
	return err
}

// WaitForRestart returns true if all containers in a given pod list reached the Ready state and started after restartTime
func WaitForRestart(c Client, namespace string, pods []string, restartTime time.Time, interval, timeout time.Duration) error {
	v1RestartTime := metav1.NewTime(restartTime)
	podStartTime := make(map[string]metav1.Time)
	for _, pod := range pods {
		podStartTime[pod] = v1RestartTime
	}
	err := wait.PollImmediate(interval, timeout, func() (bool, error) {
		pl, err := c.ListPods(namespace, metav1.ListOptions{})
		if err != nil {
			return false, err
		}
		for _, pli := range pl.Items {
			if _, ok := podStartTime[pli.ObjectMeta.Name]; ok {
				earlier := v1RestartTime
				for _, c := range pli.Status.ContainerStatuses {
					if pli.Status.Phase == v1.PodRunning && c.State.Running != nil && c.Ready {
						if c.State.Running.StartedAt.Before(&earlier) {
							earlier = c.State.Running.StartedAt
						}
					}
				}
				podStartTime[pli.ObjectMeta.Name] = earlier
			}
		}
		for _, started := range podStartTime {
			if started.Before(&v1RestartTime) {
				return false, nil
			}
		}
		return true, nil
	})
	return err
}
