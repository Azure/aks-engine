// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	evictionKind        = "Eviction"
	evictionSubresource = "pods/eviction"
)

// ClientSetClient is a Kubernetes client hooked up to a live api server.
type ClientSetClient struct {
	clientset         *kubernetes.Clientset
	interval, timeout time.Duration
}

// NewClient returns a KubernetesClient hooked up to the api server at the apiserverURL.
func NewClient(apiserverURL, kubeConfig string, interval, timeout time.Duration) (*ClientSetClient, error) {
	config, err := clientcmd.BuildConfigFromKubeconfigGetter(apiserverURL, func() (*clientcmdapi.Config, error) {
		return clientcmd.Load([]byte(kubeConfig))
	})
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &ClientSetClient{clientset: clientset, interval: interval, timeout: timeout}, nil
}

// ListPods returns Pods running on the passed in node.
func (c *ClientSetClient) ListPods(node *v1.Node) (*v1.PodList, error) {
	return c.ListPodsByOptions(metav1.NamespaceAll, metav1.ListOptions{
		FieldSelector: fields.SelectorFromSet(fields.Set{"spec.nodeName": node.Name}).String()})
}

// ListAllPods returns all Pods running.
func (c *ClientSetClient) ListAllPods() (*v1.PodList, error) {
	return c.ListPodsByOptions(metav1.NamespaceAll, metav1.ListOptions{})
}

// ListPodsByOptions returns Pods based on the passed in list options.
func (c *ClientSetClient) ListPodsByOptions(namespace string, opts metav1.ListOptions) (*v1.PodList, error) {
	ctx := context.TODO()
	return c.clientset.CoreV1().Pods(namespace).List(ctx, opts)
}

// ListNodes returns a list of Nodes registered in the api server.
func (c *ClientSetClient) ListNodes() (*v1.NodeList, error) {
	return c.ListNodesByOptions(metav1.ListOptions{})
}

// ListNodesByOptions returns a list of Nodes registered in the api server.
func (c *ClientSetClient) ListNodesByOptions(opts metav1.ListOptions) (*v1.NodeList, error) {
	ctx := context.TODO()
	return c.clientset.CoreV1().Nodes().List(ctx, opts)
}

// ListServiceAccounts returns a list of Service Accounts in the provided namespace.
func (c *ClientSetClient) ListServiceAccounts(namespace string) (*v1.ServiceAccountList, error) {
	return c.ListServiceAccountsByOptions(namespace, metav1.ListOptions{})
}

// ListServiceAccountsByOptions returns a list of Service Accounts in the provided namespace.
func (c *ClientSetClient) ListServiceAccountsByOptions(namespace string, opts metav1.ListOptions) (*v1.ServiceAccountList, error) {
	ctx := context.TODO()
	return c.clientset.CoreV1().ServiceAccounts(namespace).List(ctx, opts)
}

// ListDeployments returns a list of deployments in the provided namespace.
func (c *ClientSetClient) ListDeployments(namespace string, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	ctx := context.TODO()
	return c.clientset.AppsV1().Deployments(namespace).List(ctx, opts)
}

// ListDaemonSets returns a list of daemonsets in the provided namespace.
func (c *ClientSetClient) ListDaemonSets(namespace string, opts metav1.ListOptions) (*appsv1.DaemonSetList, error) {
	ctx := context.TODO()
	return c.clientset.AppsV1().DaemonSets(namespace).List(ctx, opts)
}

// ListSecrets returns a list of secrets in the provided namespace.
func (c *ClientSetClient) ListSecrets(namespace string, opts metav1.ListOptions) (*v1.SecretList, error) {
	ctx := context.TODO()
	return c.clientset.CoreV1().Secrets(namespace).List(ctx, opts)
}

// PatchDeployment applies a JSON patch to a deployment in the provided namespace.
func (c *ClientSetClient) PatchDeployment(namespace, name, jsonPatch string) (*appsv1.Deployment, error) {
	ctx := context.TODO()
	return c.clientset.AppsV1().Deployments(namespace).Patch(ctx, name, types.StrategicMergePatchType, []byte(jsonPatch), metav1.PatchOptions{})
}

// PatchDaemonSet applies a JSON patch to a daemonset in the provided namespace.
func (c *ClientSetClient) PatchDaemonSet(namespace, name, jsonPatch string) (*appsv1.DaemonSet, error) {
	ctx := context.TODO()
	return c.clientset.AppsV1().DaemonSets(namespace).Patch(ctx, name, types.StrategicMergePatchType, []byte(jsonPatch), metav1.PatchOptions{})
}

// GetNode returns details about node with passed in name.
func (c *ClientSetClient) GetNode(name string) (*v1.Node, error) {
	ctx := context.TODO()
	return c.clientset.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})
}

// UpdateNode updates the node in the api server with the passed in info.
func (c *ClientSetClient) UpdateNode(node *v1.Node) (*v1.Node, error) {
	ctx := context.TODO()
	return c.clientset.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})
}

// DeleteNode deregisters the node in the api server.
func (c *ClientSetClient) DeleteNode(name string) error {
	ctx := context.TODO()
	return c.clientset.CoreV1().Nodes().Delete(ctx, name, metav1.DeleteOptions{})
}

// DeleteServiceAccount deletes the passed in service account.
func (c *ClientSetClient) DeleteServiceAccount(sa *v1.ServiceAccount) error {
	ctx := context.TODO()
	return c.clientset.CoreV1().ServiceAccounts(sa.Namespace).Delete(ctx, sa.Name, metav1.DeleteOptions{})
}

// SupportEviction queries the api server to discover if it supports eviction, and returns supported type if it is supported.
func (c *ClientSetClient) SupportEviction() (string, error) {
	discoveryClient := c.clientset.Discovery()
	groupList, err := discoveryClient.ServerGroups()
	if err != nil {
		return "", err
	}
	foundPolicyGroup := false
	var policyGroupVersion string
	for _, group := range groupList.Groups {
		if group.Name == "policy" {
			foundPolicyGroup = true
			policyGroupVersion = group.PreferredVersion.GroupVersion
			break
		}
	}
	if !foundPolicyGroup {
		return "", nil
	}
	resourceList, err := discoveryClient.ServerResourcesForGroupVersion("v1")
	if err != nil {
		return "", err
	}
	for _, resource := range resourceList.APIResources {
		if resource.Name == evictionSubresource && resource.Kind == evictionKind {
			return policyGroupVersion, nil
		}
	}
	return "", nil
}

// DeleteClusterRole deletes the passed in cluster role.
func (c *ClientSetClient) DeleteClusterRole(role *rbacv1.ClusterRole) error {
	ctx := context.TODO()
	return c.clientset.RbacV1().ClusterRoles().Delete(ctx, role.Name, metav1.DeleteOptions{})
}

// DeleteDaemonSet deletes the passed in daemonset.
func (c *ClientSetClient) DeleteDaemonSet(daemonset *appsv1.DaemonSet) error {
	ctx := context.TODO()
	return c.clientset.AppsV1().DaemonSets(daemonset.Namespace).Delete(ctx, daemonset.Name, metav1.DeleteOptions{})
}

// DeleteDeployment deletes the passed in daemonset.
func (c *ClientSetClient) DeleteDeployment(deployment *appsv1.Deployment) error {
	ctx := context.TODO()
	return c.clientset.AppsV1().Deployments(deployment.Namespace).Delete(ctx, deployment.Name, metav1.DeleteOptions{})
}

// DeletePod deletes the passed in pod.
func (c *ClientSetClient) DeletePod(pod *v1.Pod) error {
	ctx := context.TODO()
	return c.clientset.CoreV1().Pods(pod.Namespace).Delete(ctx, pod.Name, metav1.DeleteOptions{})
}

// DeletePods deletes all pods in a namespace that match the option filters.
func (c *ClientSetClient) DeletePods(namespace string, opts metav1.ListOptions) error {
	ctx := context.TODO()
	return c.clientset.CoreV1().Pods(namespace).DeleteCollection(ctx, metav1.DeleteOptions{}, opts)
}

// DeleteSecret deletes the passed in secret.
func (c *ClientSetClient) DeleteSecret(secret *v1.Secret) error {
	ctx := context.TODO()
	return c.clientset.CoreV1().Secrets(secret.Namespace).Delete(ctx, secret.Name, metav1.DeleteOptions{})
}

// EvictPod evicts the passed in pod using the passed in api version.
func (c *ClientSetClient) EvictPod(pod *v1.Pod, policyGroupVersion string) error {
	ctx := context.TODO()
	eviction := &policy.Eviction{
		TypeMeta: metav1.TypeMeta{
			APIVersion: policyGroupVersion,
			Kind:       evictionKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		},
	}
	return c.clientset.PolicyV1beta1().Evictions(eviction.Namespace).Evict(ctx, eviction)
}

// GetPod returns the pod.
func (c *ClientSetClient) getPod(namespace, name string) (*v1.Pod, error) {
	ctx := context.TODO()
	return c.clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
}

// WaitForDelete waits until all pods are deleted. Returns all pods not deleted and an error on failure.
func (c *ClientSetClient) WaitForDelete(logger *log.Entry, pods []v1.Pod, usingEviction bool) ([]v1.Pod, error) {
	verbStr := "deleted"
	if usingEviction {
		verbStr = "evicted"
	}
	err := wait.PollImmediate(c.interval, c.timeout, func() (bool, error) {
		pendingPods := []v1.Pod{}
		for i, pod := range pods {
			p, err := c.getPod(pod.Namespace, pod.Name)
			if apierrors.IsNotFound(err) || (p != nil && p.ObjectMeta.UID != pod.ObjectMeta.UID) {
				logger.Infof("%s pod successfully %s", pod.Name, verbStr)
				continue
			} else if err != nil {
				return false, err
			} else {
				pendingPods = append(pendingPods, pods[i])
			}
		}
		pods = pendingPods
		if len(pendingPods) > 0 {
			return false, nil
		}
		return true, nil
	})
	return pods, err
}

// GetDaemonSet returns a given daemonset in a namespace.
func (c *ClientSetClient) GetDaemonSet(namespace, name string) (*appsv1.DaemonSet, error) {
	ctx := context.TODO()
	return c.clientset.AppsV1().DaemonSets(namespace).Get(ctx, name, metav1.GetOptions{})
}

// GetDeployment returns a given deployment in a namespace.
func (c *ClientSetClient) GetDeployment(namespace, name string) (*appsv1.Deployment, error) {
	ctx := context.TODO()
	return c.clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
}

// UpdateDeployment updates a deployment to match the given specification.
func (c *ClientSetClient) UpdateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	ctx := context.TODO()
	return c.clientset.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
}
