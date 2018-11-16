// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

// Compiler to verify QueueMessageProcessor implements OperationsProcessor
var _ UpgradeWorkFlow = &Kubernetes16upgrader{}

// Kubernetes16upgrader upgrades a Kubernetes 1.5 cluster to 1.6
type Kubernetes16upgrader struct {
	Upgrader
}
