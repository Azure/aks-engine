// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetesupgrade

// Compiler to verify QueueMessageProcessor implements OperationsProcessor
var _ UpgradeWorkFlow = &Kubernetes17upgrader{}

// Kubernetes17upgrader upgrades a Kubernetes 1.6 cluster to 1.7
type Kubernetes17upgrader struct {
	Upgrader
}
