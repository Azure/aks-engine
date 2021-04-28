// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package internal

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KubeClient interface {
	// ListPods returns Pods based on the passed in list options.
	ListPods(namespace string, opts metav1.ListOptions) (*v1.PodList, error)
	// ListNodes returns a list of Nodes registered in the api server.
	ListNodes() (*v1.NodeList, error)
	// ListServiceAccounts returns a list of Service Accounts in the provided namespace.
	ListServiceAccounts(namespace string, opts metav1.ListOptions) (*v1.ServiceAccountList, error)
	// ListDeployments returns a list of deployments in the provided namespace.
	ListDeployments(namespace string, opts metav1.ListOptions) (*appsv1.DeploymentList, error)
	// ListDaemonSets returns a list of daemonsets in the provided namespace.
	ListDaemonSets(namespace string, opts metav1.ListOptions) (*appsv1.DaemonSetList, error)
	// GetDeployment returns a given deployment in a namespace.
	GetDeployment(namespace, name string) (*appsv1.Deployment, error)
	// PatchDeployment applies a JSON patch to a deployment in the provided namespace.
	PatchDeployment(namespace, name, jsonPatch string) (*appsv1.Deployment, error)
	// PatchDaemonSet applies a JSON patch to a daemonset in the provided namespace.
	PatchDaemonSet(namespace, name, jsonPatch string) (*appsv1.DaemonSet, error)
	// DeletePods deletes all pods in a namespace that match the option filters.
	DeletePods(namespace string, opts metav1.ListOptions) error
	// DeleteServiceAccount deletes the passed in service account.
	DeleteServiceAccount(secret *v1.ServiceAccount) error
	// DeleteSecret deletes the passed in secret.
	DeleteSecret(secret *v1.Secret) error
}

type ARMClient interface {
	// RestartVirtualMachine restarts the specified virtual machine.
	RestartVirtualMachine(resourceGroup, vmName string) error

	// RestartVirtualMachineScaleSets restarts the specified virtual machine scale set.
	RestartVirtualMachineScaleSets(resourceGroup, vmssName string) error

	// GetVirtualMachinePowerState returns the virtual machine's Power state.
	GetVirtualMachinePowerState(resourceGroup, vmName string) (string, error)

	// GetVirtualMachineScaleSetInstancePowerState returns the virtual machine's Power state.
	GetVirtualMachineScaleSetInstancePowerState(resourceGroup, vmssName, instanceID string) (string, error)
}
