// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/aks-engine/pkg/helpers/ssh"
	v1 "k8s.io/api/core/v1"
)

func TestHealPodsFunc(t *testing.T) {
	t.Parallel()

	alwaysTrue := func(_ v1.Pod) bool { return true }
	alwaysFalse := func(_ v1.Pod) bool { return false }

	t.Run("All conditions met, heal func called", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		healed := make(chan bool)
		defer close(healed)
		healFunc := func(_ v1.Pod, _ *ssh.RemoteHost) { healed <- true }

		podList := &v1.PodList{Items: []v1.Pod{v1.Pod{}}}
		podList.Items[0].Spec.NodeName = "node1"
		nodes := map[string]*ssh.RemoteHost{
			podList.Items[0].Spec.NodeName: nil,
		}

		healPodsFunc(ctx, nodes, healFunc, alwaysTrue, alwaysTrue)(podList)

		select {
		case <-healed:
			return
		case <-time.After(2 * time.Second):
			t.Fatal("Expected healFunc invocation did not happen")
		}
	})

	t.Run("No condition met, expect nothing", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		healed := make(chan bool)
		defer close(healed)
		healFunc := func(_ v1.Pod, _ *ssh.RemoteHost) { healed <- true }

		podList := &v1.PodList{Items: []v1.Pod{v1.Pod{}}}
		podList.Items[0].Spec.NodeName = "node1"
		nodes := map[string]*ssh.RemoteHost{
			podList.Items[0].Spec.NodeName: nil,
		}

		healPodsFunc(ctx, nodes, healFunc, alwaysFalse, alwaysFalse)(podList)

		select {
		case <-healed:
			t.Fatal("Unexpected healFunc invocation")
		case <-time.After(2 * time.Second):
		}
	})

	t.Run("No condition met, expect nothing", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		healed := make(chan bool)
		defer close(healed)
		healFunc := func(_ v1.Pod, _ *ssh.RemoteHost) { healed <- true }

		podList := &v1.PodList{Items: []v1.Pod{v1.Pod{}}}
		podList.Items[0].Spec.NodeName = "node1"
		nodes := map[string]*ssh.RemoteHost{
			podList.Items[0].Spec.NodeName: nil,
		}

		healPodsFunc(ctx, nodes, healFunc, alwaysTrue, alwaysFalse)(podList)

		select {
		case <-healed:
			t.Fatal("Unexpected healFunc invocation")
		case <-time.After(2 * time.Second):
		}
	})
}
