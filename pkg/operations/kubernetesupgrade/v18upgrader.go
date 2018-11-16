// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

// Compiler to verify QueueMessageProcessor implements OperationsProcessor
var _ UpgradeWorkFlow = &Kubernetes18upgrader{}

// Kubernetes18upgrader upgrades a Kubernetes 1.7.x cluster to 1.8.x
type Kubernetes18upgrader struct {
	Upgrader
}
