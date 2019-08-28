// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

// Controller is the node labeler itself.
type Controller struct {
	kubeclientset kubernetes.Interface
	indexer       cache.Indexer
	queue         workqueue.RateLimitingInterface
	informer      cache.Controller
}

// NewController returns a new node labeler controller.
func NewController(kubeclientset kubernetes.Interface, queue workqueue.RateLimitingInterface,
	indexer cache.Indexer, informer cache.Controller) *Controller {
	return &Controller{
		kubeclientset: kubeclientset,
		informer:      informer,
		indexer:       indexer,
		queue:         queue,
	}
}

// processNextItem pulls items off the queue and dispatches them to business logic.
func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.patchLabelsIfNeeded(key.(string))
	c.handleErr(err, key)

	return true
}

// NodeLabelsType is a map of string to string, representing a set of Kubernetes labels.
type NodeLabelsType map[string]string

// patchLabelsIfNeeded examines a Node to see if it is missing one or more master or agent labels,
// and applies them via PATCH if necessary.
func (c *Controller) patchLabelsIfNeeded(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		klog.Errorf("getting object by key %s: %v", key, err)
		return err
	}

	if exists {
		node := obj.(*v1.Node)

		// NodeLabels is a list of NodeLabelsType containing the expected Kubernetes labels for
		// master or agent nodes.
		var NodeLabels = []NodeLabelsType{
			{
				"kubernetes.azure.com/role":      "master",
				"kubernetes.io/role":             "master",
				"node-role.kubernetes.io/master": "",
			},
			{
				"kubernetes.azure.com/role":     "agent",
				"kubernetes.io/role":            "agent",
				"node-role.kubernetes.io/agent": "",
			},
		}

		for _, nodeLabels := range NodeLabels {
			labels := node.GetLabels()
			if needsLabeling(labels, nodeLabels) {
				if err := c.labelNode(node, nodeLabels); err != nil {
					klog.Error(err)
					break
				}
			}
		}
	}

	return nil
}

// needsLabeling inspects a map of labels to see if it is missing any from the
// second set of labels provided.
func needsLabeling(labels, nodeLabels map[string]string) bool {
	matched := 0
	for key, value := range nodeLabels {
		for key1, value1 := range labels {
			if key == key1 && value == value1 {
				matched++
			}
		}
	}

	// return true if the node has any--but not all--of the specified labels
	return (matched > 0 && matched < len(nodeLabels))
}

func (c *Controller) labelNode(node *v1.Node, labels map[string]string) error {
	oldData, err := json.Marshal(node)
	if err != nil {
		return err
	}
	nodeLabels := node.GetLabels()
	for k, v := range nodeLabels {
		labels[k] = v
	}
	node.SetLabels(labels)
	newData, err := json.Marshal(node)
	if err != nil {
		return err
	}
	patchBytes, err := jsonpatch.CreateMergePatch(oldData, newData)
	if err != nil {
		return err
	}
	name := node.GetName()
	klog.Warningf("updating labels on Node %s:\n\t%s\n", name, string(patchBytes))
	if _, err := c.kubeclientset.CoreV1().Nodes().Patch(name, types.MergePatchType, patchBytes); err != nil {
		return err
	}
	return nil
}

// handleErr checks if an error happened and queues the failed event to be retried.
func (c *Controller) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	// Retry 5 times before giving up.
	if c.queue.NumRequeues(key) < 5 {
		klog.Infof("error syncing node %v: %v", key, err)
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)

	runtime.HandleError(err)
	klog.Infof("dropping node %q from the queue: %v", key, err)
}

// Run starts the node labeler controller in motion.
func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	klog.Info("starting Node controller")

	go c.informer.Run(stopCh)

	// Wait for all involved caches to be synced before processing.
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(errors.New("timeout waiting for cache sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	klog.Info("stopping Node controller")
}

// runWorker dispatches items to be processed in a loop.
func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func main() {
	// set klog to send to stderr
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	logtostderr := klogFlags.Lookup("logtostderr")
	if err := logtostderr.Value.Set("true"); err == nil {
		klog.Warning(err)
	}

	// use the in-cluster Kubernetes config
	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	// TODO: listen only for events where labels actually changed?
	nodeWatcher := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(), "nodes", v1.NamespaceAll, fields.Everything())

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	indexer, informer := cache.NewIndexerInformer(nodeWatcher, &v1.Node{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if key, err := cache.MetaNamespaceKeyFunc(obj); err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			if key, err := cache.MetaNamespaceKeyFunc(new); err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	controller := NewController(clientset, queue, indexer, informer)

	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)

	select {}
}
