// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	"crypto/x509"
	"net/url"
	"time"

	"github.com/Azure/aks-engine/pkg/kubernetes/internal"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

// CompositeClientSet wraps a pair of Kubernetes clients hooked up to a live api server.
//
// Prefer this client when the cluster CA is expected to change (ex.: secret rotation operations).
type CompositeClientSet struct {
	oldCAClient internal.Client
	newCAClient internal.Client
	timeout     time.Duration
	backoff     wait.Backoff
	retryFunc   func(err error) bool
}

// NewCompositeClient returns a KubernetesClient hooked up to the api server at the apiserverURL.
func NewCompositeClient(oldCAClient, newCAClient internal.Client, interval, timeout time.Duration) *CompositeClientSet {
	return &CompositeClientSet{
		oldCAClient: oldCAClient,
		newCAClient: newCAClient,
		timeout:     timeout,
		backoff: wait.Backoff{
			Steps:    int(int64(timeout/time.Millisecond) / int64(interval/time.Millisecond)),
			Duration: interval,
			Factor:   1.0,
			Jitter:   0.0,
		},
		retryFunc: retriable, // Inject if ever needed
	}
}

// retriable returns true unless err is an x509.UnknownAuthorityError instance
func retriable(err error) bool {
	switch err := err.(type) {
	case x509.UnknownAuthorityError:
		return false
	case *url.Error:
		return retriable(err.Unwrap())
	default:
		return true
	}
}

type listPodsResult struct {
	x   *v1.PodList
	err error
}

// ListPods returns Pods based on the passed in list options.
func (c *CompositeClientSet) ListPods(namespace string, opts metav1.ListOptions) (*v1.PodList, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan listPodsResult {
		stream := make(chan listPodsResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.ListPodsByOptions(namespace, opts)
				if err != nil {
					lastError = err
					return err
				}
				stream <- listPodsResult{x, err}
				return nil
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

type listNodesResult struct {
	x   *v1.NodeList
	err error
}

// ListNodes returns a list of Nodes registered in the api server.
func (c *CompositeClientSet) ListNodes() (x *v1.NodeList, err error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan listNodesResult {
		stream := make(chan listNodesResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.ListNodes()
				if err != nil {
					lastError = err
					return err
				}
				stream <- listNodesResult{x, err}
				return nil
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

type listServiceAccountsResult struct {
	x   *v1.ServiceAccountList
	err error
}

// ListServiceAccounts returns a list of Service Accounts in the provided namespace.
func (c *CompositeClientSet) ListServiceAccounts(namespace string, opts metav1.ListOptions) (*v1.ServiceAccountList, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan listServiceAccountsResult {
		stream := make(chan listServiceAccountsResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.ListServiceAccountsByOptions(namespace, opts)
				if err != nil {
					lastError = err
					return err
				}
				stream <- listServiceAccountsResult{x, err}
				return nil
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

type listDeploymentsResult struct {
	x   *appsv1.DeploymentList
	err error
}

// ListDeployments returns a list of deployments in the provided namespace.
func (c *CompositeClientSet) ListDeployments(namespace string, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan listDeploymentsResult {
		stream := make(chan listDeploymentsResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.ListDeployments(namespace, opts)
				if err != nil {
					lastError = err
					return err
				}
				stream <- listDeploymentsResult{x, err}
				return nil
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

type listDaemonSetsResult struct {
	x   *appsv1.DaemonSetList
	err error
}

// ListDaemonSets returns a list of daemonsets in the provided namespace.
func (c *CompositeClientSet) ListDaemonSets(namespace string, opts metav1.ListOptions) (*appsv1.DaemonSetList, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan listDaemonSetsResult {
		stream := make(chan listDaemonSetsResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.ListDaemonSets(namespace, opts)
				if err != nil {
					lastError = err
					return err
				}
				stream <- listDaemonSetsResult{x, err}
				return nil
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

type listSecretsResult struct {
	x   *v1.SecretList
	err error
}

// ListSecrets returns a list of secrets in the provided namespace.
func (c *CompositeClientSet) ListSecrets(namespace string, opts metav1.ListOptions) (*v1.SecretList, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan listSecretsResult {
		stream := make(chan listSecretsResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.ListSecrets(namespace, opts)
				if err != nil {
					lastError = err
					return err
				}
				stream <- listSecretsResult{x, err}
				return nil
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

type deploymentResult struct {
	x   *appsv1.Deployment
	err error
}

// GetDeployment blah.
func (c *CompositeClientSet) GetDeployment(namespace, name string) (*appsv1.Deployment, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan deploymentResult {
		stream := make(chan deploymentResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.GetDeployment(namespace, name)
				if err == nil || apierrors.IsNotFound(err) {
					stream <- deploymentResult{x, err}
					return nil
				}
				lastError = err
				return err
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// PatchDeployment applies a JSON patch to a deployment in the provided namespace.
func (c *CompositeClientSet) PatchDeployment(namespace, name, jsonPatch string) (*appsv1.Deployment, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan deploymentResult {
		stream := make(chan deploymentResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.PatchDeployment(namespace, name, jsonPatch)
				if err == nil || apierrors.IsNotFound(err) {
					stream <- deploymentResult{x, err}
					return nil
				}
				lastError = err
				return err
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

type daemonsetResult struct {
	x   *appsv1.DaemonSet
	err error
}

// PatchDaemonSet applies a JSON patch to a daemonset in the provided namespace.
func (c *CompositeClientSet) PatchDaemonSet(namespace, name, jsonPatch string) (*appsv1.DaemonSet, error) {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan daemonsetResult {
		stream := make(chan daemonsetResult)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				x, err := client.PatchDaemonSet(namespace, name, jsonPatch)
				if err == nil || apierrors.IsNotFound(err) {
					stream <- daemonsetResult{x, err}
					return nil
				}
				lastError = err
				return err
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case res := <-result:
			return res.x, res.err
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// DeletePods deletes all pods in a namespace that match the option filters.
func (c *CompositeClientSet) DeletePods(namespace string, opts metav1.ListOptions) error {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan error {
		stream := make(chan error)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				err := client.DeletePods(namespace, opts)
				if err == nil || apierrors.IsNotFound(err) {
					stream <- err
					return nil
				}
				lastError = err
				return err
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case err := <-result:
			return err
		case <-time.After(c.timeout):
			return lastError
		}
	}
}

// DeleteServiceAccount deletes the passed in service account.
func (c *CompositeClientSet) DeleteServiceAccount(serviceAccount *v1.ServiceAccount) error {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan error {
		stream := make(chan error)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				err := client.DeleteServiceAccount(serviceAccount)
				if err == nil || apierrors.IsNotFound(err) {
					stream <- err
					return nil
				}
				lastError = err
				return err
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case err := <-result:
			return err
		case <-time.After(c.timeout):
			return lastError
		}
	}
}

// DeleteSecret deletes the passed in secret.
func (c *CompositeClientSet) DeleteSecret(secret *v1.Secret) error {
	lastError := wait.ErrWaitTimeout
	result := func(oldCAClient, newCAClient internal.Client) <-chan error {
		stream := make(chan error)
		exec := func(client internal.Client) {
			_ = retry.OnError(c.backoff, c.retryFunc, func() error {
				err := client.DeleteSecret(secret)
				if err == nil || apierrors.IsNotFound(err) {
					stream <- err
					return nil
				}
				lastError = err
				return err
			})
		}
		go exec(oldCAClient)
		go exec(newCAClient)
		return stream
	}(c.oldCAClient, c.newCAClient)
	for {
		select {
		case err := <-result:
			return err
		case <-time.After(c.timeout):
			return lastError
		}
	}
}
