// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO These interfaces do not follow best practices
// https://github.com/golang/go/wiki/CodeReviewComments#interfaces

// Client interface models client for interacting with kubernetes api server
type Client interface {
	// ListPods returns Pods based on the passed in list options.
	ListPods(namespace string, opts metav1.ListOptions) (*v1.PodList, error)
	// ListNodes returns a list of Nodes registered in the api server.
	ListNodes() (*v1.NodeList, error)
	// ListNodesByOptions returns a list of Nodes registered in the api server.
	ListNodesByOptions(opts metav1.ListOptions) (*v1.NodeList, error)
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
	// GetDaemonSet returns details about DaemonSet with passed in name.
	GetDaemonSet(namespace, name string) (*appsv1.DaemonSet, error)
	// GetDeployment returns a given deployment in a namespace.
	GetDeployment(namespace, name string) (*appsv1.Deployment, error)
	// GetNode returns details about node with passed in name.
	GetNode(name string) (*v1.Node, error)
	// UpdateNode updates the node in the api server with the passed in info.
	UpdateNode(node *v1.Node) (*v1.Node, error)
	// DeleteNode deregisters node in the api server.
	DeleteNode(name string) error
	// SupportEviction queries the api server to discover if it supports eviction, and returns supported type if it is supported.
	SupportEviction() (string, error)
	// DeleteClusterRole deletes the passed in ClusterRole.
	DeleteClusterRole(role *rbacv1.ClusterRole) error
	// DeleteDaemonSet deletes the passed in DaemonSet.
	DeleteDaemonSet(ds *appsv1.DaemonSet) error
	// DeleteDeployment deletes the passed in Deployment.
	DeleteDeployment(ds *appsv1.Deployment) error
	// DeletePod deletes the passed in pod.
	DeletePod(pod *v1.Pod) error
	// DeletePods deletes all pods in a namespace that match the option filters.
	DeletePods(namespace string, opts metav1.ListOptions) error
	// DeleteSecret deletes the passed in secret.
	DeleteSecret(secret *v1.Secret) error
	// DeleteServiceAccount deletes the passed in service account.
	DeleteServiceAccount(sa *v1.ServiceAccount) error
	// EvictPod evicts the passed in pod using the passed in api version.
	EvictPod(pod *v1.Pod, policyGroupVersion string) error
	// WaitForDelete waits until all pods are deleted. Returns all pods not deleted and an error on failure.
	WaitForDelete(logger *log.Entry, pods []v1.Pod, usingEviction bool) ([]v1.Pod, error)
	// UpdateDeployment updates a deployment to match the given specification.
	UpdateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error)
}

// NodeLister is an interface implemented by Kubernetes clients
// that are able to list cluster nodes
type NodeLister interface {
	ListNodes() (*v1.NodeList, error)
}
