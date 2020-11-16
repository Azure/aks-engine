// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Client interface {
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
	// ListSecrets returns a list of secrets in the provided namespace.
	ListSecrets(namespace string, opts metav1.ListOptions) (*v1.SecretList, error)
	// PatchDeployment applies a JSON patch to a deployment in the provided namespace.
	PatchDeployment(namespace, name, jsonPatch string) (*appsv1.Deployment, error)
	// PatchDaemonSet applies a JSON patch to a daemonset in the provided namespace.
	PatchDaemonSet(namespace, name, jsonPatch string) (*appsv1.DaemonSet, error)
	// DeletePods deletes all pods in a namespace that match the option filters.
	DeletePods(namespace string, opts metav1.ListOptions) error
	// DeleteSecret deletes the passed in secret.
	DeleteSecret(secret *v1.Secret) error
}
