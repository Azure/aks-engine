// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	"time"

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
	oldCAClient Client
	newCAClient Client
	timeout     time.Duration
	backoff     wait.Backoff
	retryFunc   func(err error) bool
}

// NewCompositeClient returns a KubernetesClient hooked up to the api server at the apiserverURL.
func NewCompositeClient(oldCAClient, newCAClient Client, interval, timeout time.Duration) *CompositeClientSet {
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
		retryFunc: func(err error) bool { return true }, // Always retry. Inject if ever needed
	}
}

// ListPods returns Pods based on the passed in list options.
func (c *CompositeClientSet) ListPods(namespace string, opts metav1.ListOptions) (*v1.PodList, error) {
	var lastError error
	f := func(client Client, resChan chan *v1.PodList, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.ListPods(namespace, opts)
			if err != nil {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *v1.PodList)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// ListNodes returns a list of Nodes registered in the api server.
func (c *CompositeClientSet) ListNodes() (x *v1.NodeList, err error) {
	var lastError error
	f := func(client Client, resChan chan *v1.NodeList, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.ListNodes()
			if err != nil {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *v1.NodeList)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// ListServiceAccounts returns a list of Service Accounts in the provided namespace.
func (c *CompositeClientSet) ListServiceAccounts(namespace string, opts metav1.ListOptions) (*v1.ServiceAccountList, error) {
	var lastError error
	f := func(client Client, resChan chan *v1.ServiceAccountList, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.ListServiceAccounts(namespace, opts)
			if err != nil {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *v1.ServiceAccountList)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// ListDeployments returns a list of deployments in the provided namespace.
func (c *CompositeClientSet) ListDeployments(namespace string, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	var lastError error
	f := func(client Client, resChan chan *appsv1.DeploymentList, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.ListDeployments(namespace, opts)
			if err != nil {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *appsv1.DeploymentList)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// ListDaemonSets returns a list of daemonsets in the provided namespace.
func (c *CompositeClientSet) ListDaemonSets(namespace string, opts metav1.ListOptions) (*appsv1.DaemonSetList, error) {
	var lastError error
	f := func(client Client, resChan chan *appsv1.DaemonSetList, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.ListDaemonSets(namespace, opts)
			if err != nil {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *appsv1.DaemonSetList)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// ListSecrets returns a list of secrets in the provided namespace.
func (c *CompositeClientSet) ListSecrets(namespace string, opts metav1.ListOptions) (*v1.SecretList, error) {
	var lastError error
	f := func(client Client, resChan chan *v1.SecretList, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.ListSecrets(namespace, opts)
			if err != nil {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *v1.SecretList)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// PatchDeployment applies a JSON patch to a deployment in the provided namespace.
func (c *CompositeClientSet) PatchDeployment(namespace, name, jsonPatch string) (*appsv1.Deployment, error) {
	var lastError error
	f := func(client Client, resChan chan *appsv1.Deployment, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.PatchDeployment(namespace, name, jsonPatch)
			if err != nil && !apierrors.IsNotFound(err) {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *appsv1.Deployment)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// PatchDaemonSet applies a JSON patch to a daemonset in the provided namespace.
func (c *CompositeClientSet) PatchDaemonSet(namespace, name, jsonPatch string) (*appsv1.DaemonSet, error) {
	var lastError error
	f := func(client Client, resChan chan *appsv1.DaemonSet, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			x, err := client.PatchDaemonSet(namespace, name, jsonPatch)
			if err != nil && !apierrors.IsNotFound(err) {
				errChan <- err
				return err
			}
			resChan <- x
			return nil
		})
	}

	resChan := make(chan *appsv1.DaemonSet)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case x := <-resChan:
			return x, nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return nil, lastError
		}
	}
}

// DeletePods deletes all pods in a namespace that match the option filters.
func (c *CompositeClientSet) DeletePods(namespace string, opts metav1.ListOptions) error {
	var lastError error
	f := func(client Client, resChan chan bool, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			err := client.DeletePods(namespace, opts)
			if err != nil && !apierrors.IsNotFound(err) {
				errChan <- err
				return err
			}
			resChan <- true
			return nil
		})
	}

	resChan := make(chan bool)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case <-resChan:
			return nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return lastError
		}
	}
}

// DeleteSecret deletes the passed in secret.
func (c *CompositeClientSet) DeleteSecret(secret *v1.Secret) error {
	var lastError error
	f := func(client Client, resChan chan bool, errChan chan error) {
		lastError = retry.OnError(c.backoff, c.retryFunc, func() error {
			err := client.DeleteSecret(secret)
			if err != nil && !apierrors.IsNotFound(err) {
				errChan <- err
				return err
			}
			resChan <- true
			return nil
		})
	}

	resChan := make(chan bool)
	errChan := make(chan error)
	go f(c.oldCAClient, resChan, errChan)
	go f(c.newCAClient, resChan, errChan)

	for {
		select {
		case <-resChan:
			return nil
		case lastError = <-errChan:
		case <-time.After(c.timeout):
			return lastError
		}
	}
}
