// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"time"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	evictionKind        = "Eviction"
	evictionSubresource = "pods/eviction"
)

//KubernetesClientSetClient is a Kubernetes client hooked up to a live api server.
type KubernetesClientSetClient struct {
	clientset         *kubernetes.Clientset
	interval, timeout time.Duration
}

//GetKubernetesClient returns a KubernetesClient hooked up to the api server at the apiserverURL.
func (az *AzureClient) GetKubernetesClient(apiserverURL, kubeConfig string, interval, timeout time.Duration) (armhelpers.KubernetesClient, error) {
	// creates the clientset
	config, err := clientcmd.BuildConfigFromKubeconfigGetter(apiserverURL, func() (*clientcmdapi.Config, error) { return clientcmd.Load([]byte(kubeConfig)) })
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &KubernetesClientSetClient{clientset: clientset, interval: interval, timeout: timeout}, nil
}

// ListPods returns all Pods running on the passed in node.
func (c *KubernetesClientSetClient) ListPods(node *v1.Node) (*v1.PodList, error) {
	return c.clientset.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{
		FieldSelector: fields.SelectorFromSet(fields.Set{"spec.nodeName": node.Name}).String()})
}

// ListAllPods returns all Pods running.
func (c *KubernetesClientSetClient) ListAllPods() (*v1.PodList, error) {
	return c.clientset.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{})
}

// ListNodes returns a list of Nodes registered in the api server.
func (c *KubernetesClientSetClient) ListNodes() (*v1.NodeList, error) {
	return c.clientset.CoreV1().Nodes().List(metav1.ListOptions{})
}

// ListServiceAccounts returns a list of Service Accounts in the provided namespace.
func (c *KubernetesClientSetClient) ListServiceAccounts(namespace string) (*v1.ServiceAccountList, error) {
	return c.clientset.CoreV1().ServiceAccounts(namespace).List(metav1.ListOptions{})
}

// GetNode returns details about node with passed in name.
func (c *KubernetesClientSetClient) GetNode(name string) (*v1.Node, error) {
	return c.clientset.CoreV1().Nodes().Get(name, metav1.GetOptions{})
}

// UpdateNode updates the node in the api server with the passed in info.
func (c *KubernetesClientSetClient) UpdateNode(node *v1.Node) (*v1.Node, error) {
	return c.clientset.CoreV1().Nodes().Update(node)
}

// DeleteNode deregisters the node in the api server.
func (c *KubernetesClientSetClient) DeleteNode(name string) error {
	return c.clientset.CoreV1().Nodes().Delete(name, &metav1.DeleteOptions{})
}

// DeleteServiceAccount deletes the passed in service account.
func (c *KubernetesClientSetClient) DeleteServiceAccount(sa *v1.ServiceAccount) error {
	return c.clientset.CoreV1().ServiceAccounts(sa.Namespace).Delete(sa.Name, &metav1.DeleteOptions{})
}

// SupportEviction queries the api server to discover if it supports eviction, and returns supported type if it is supported.
func (c *KubernetesClientSetClient) SupportEviction() (string, error) {
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
func (c *KubernetesClientSetClient) DeleteClusterRole(role *rbacv1.ClusterRole) error {
	return c.clientset.RbacV1().ClusterRoles().Delete(role.Name, &metav1.DeleteOptions{})
}

// DeleteDaemonSet deletes the passed in daemonset.
func (c *KubernetesClientSetClient) DeleteDaemonSet(daemonset *appsv1.DaemonSet) error {
	return c.clientset.AppsV1().DaemonSets(daemonset.Namespace).Delete(daemonset.Name, &metav1.DeleteOptions{})
}

// DeleteDeployment deletes the passed in daemonset.
func (c *KubernetesClientSetClient) DeleteDeployment(deployment *appsv1.Deployment) error {
	return c.clientset.AppsV1().Deployments(deployment.Namespace).Delete(deployment.Name, &metav1.DeleteOptions{})
}

// DeletePod deletes the passed in pod.
func (c *KubernetesClientSetClient) DeletePod(pod *v1.Pod) error {
	return c.clientset.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{})
}

// EvictPod evicts the passed in pod using the passed in api version.
func (c *KubernetesClientSetClient) EvictPod(pod *v1.Pod, policyGroupVersion string) error {
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
	return c.clientset.PolicyV1beta1().Evictions(eviction.Namespace).Evict(eviction)
}

// GetPod returns the pod with the provided name and namespace.
func (c *KubernetesClientSetClient) getPod(namespace, name string) (*v1.Pod, error) {
	return c.clientset.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
}

// WaitForDelete waits until all pods are deleted. Returns all pods not deleted and an error on failure.
func (c *KubernetesClientSetClient) WaitForDelete(logger *log.Entry, pods []v1.Pod, usingEviction bool) ([]v1.Pod, error) {
	var verbStr string
	if usingEviction {
		verbStr = "evicted"
	} else {
		verbStr = "deleted"
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
func (c *KubernetesClientSetClient) GetDaemonSet(namespace, name string) (*appsv1.DaemonSet, error) {
	return c.clientset.AppsV1().DaemonSets(namespace).Get(name, metav1.GetOptions{})
}

// GetDeployment returns a given deployment in a namespace.
func (c *KubernetesClientSetClient) GetDeployment(namespace, name string) (*appsv1.Deployment, error) {
	return c.clientset.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
}

// UpdateDeployment updates a deployment to match the given specification.
func (c *KubernetesClientSetClient) UpdateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	return c.clientset.AppsV1().Deployments(namespace).Update(deployment)
}
