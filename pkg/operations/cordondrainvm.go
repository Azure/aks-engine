// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	interval            = time.Second * 1
	mirrorPodAnnotation = "kubernetes.io/config.mirror"

	// This is checked into K8s code but I was getting into vendoring issues so I copied it here instead
	kubernetesOptimisticLockErrorMsg = "the object has been modified; please apply your changes to the latest version and try again"
	cordonMaxRetries                 = 5
)

type drainOperation struct {
	client  armhelpers.KubernetesClient
	node    *v1.Node
	logger  *log.Entry
	timeout time.Duration
}

type podFilter func(v1.Pod) bool

// SafelyDrainNode safely drains a node so that it can be deleted from the cluster
func SafelyDrainNode(az armhelpers.AKSEngineClient, logger *log.Entry, apiserverURL, kubeConfig, nodeName string, timeout time.Duration) error {
	//get client using kubeconfig
	client, err := az.GetKubernetesClient(apiserverURL, kubeConfig, interval, timeout)
	if err != nil {
		return err
	}
	_, err = SafelyDrainNodeWithClient(client, logger, nodeName, timeout)
	return err
}

// SafelyDrainNodeWithClient safely drains a node so that it can be deleted from the cluster
func SafelyDrainNodeWithClient(client armhelpers.KubernetesClient, logger *log.Entry, nodeName string, timeout time.Duration) ([]v1.Pod, error) {
	nodeName = strings.ToLower(nodeName)
	//Mark the node unschedulable
	var node *v1.Node
	var err error
	for i := 0; i < cordonMaxRetries; i++ {
		node, err = client.GetNode(nodeName)
		if err != nil {
			return nil, err
		}
		node.Spec.Unschedulable = true
		node, err = client.UpdateNode(node)
		if err != nil {
			// If this error is because of a concurrent modification get the update
			// and then apply the change
			if strings.Contains(err.Error(), kubernetesOptimisticLockErrorMsg) {
				logger.Infof("Node %s got an error suggesting a concurrent modification. Will retry to cordon", nodeName)
				continue
			}
			return nil, err
		}
		break
	}
	logger.Infof("Node %s has been marked unschedulable.", nodeName)

	// Sleep one minute so that the node update event can be logged.
	time.Sleep(time.Second * 60)

	//Evict pods in node
	drainOp := &drainOperation{client: client, node: node, logger: logger, timeout: timeout}
	pods, err := drainOp.getPodsForDeletion()
	if err != nil {
		return nil, err
	}
	return pods, drainOp.deleteOrEvictPodsSimple(pods)
}

func (o *drainOperation) deleteOrEvictPodsSimple(pods []v1.Pod) error {
	if len(pods) > 0 {
		o.logger.WithFields(log.Fields{
			"prefix": "drain",
			"node":   o.node.Name,
		}).Infof("%d pods will be deleted", len(pods))
	} else {
		o.logger.Infof("Node %s has no scheduled pods", o.node.Name)
	}

	err := o.deleteOrEvictPods(pods)
	if err != nil {
		pendingPods, newErr := o.getPodsForDeletion()
		if newErr != nil {
			return newErr
		}
		o.logger.Errorf("There are pending pods when an error occurred: %v\n", err)
		for _, pendingPod := range pendingPods {
			o.logger.Errorf("%s/%s\n", "pod", pendingPod.Name)
		}
	}
	return err
}

func mirrorPodFilter(pod v1.Pod) bool {
	if _, found := pod.ObjectMeta.Annotations[mirrorPodAnnotation]; found {
		return false
	}
	return true
}

func getControllerRef(pod *v1.Pod) *metav1.OwnerReference {
	for _, ref := range pod.ObjectMeta.OwnerReferences {
		if ref.Controller != nil && *ref.Controller {
			return &ref
		}
	}
	return nil
}

func daemonSetPodFilter(pod v1.Pod) bool {
	controllerRef := getControllerRef(&pod)
	// Kubectl goes and verifies this controller exists in the api server to make sure it isn't orphaned
	// we are deleting orphaned pods so we don't care and delete any that aren't a daemonset
	if controllerRef == nil || controllerRef.Kind != "DaemonSet" {
		return true
	}
	// Don't delete/evict daemonsets as they will just come back
	// and can deleting/evicting them can cause service disruptions
	return false
}

// getPodsForDeletion returns all the pods we're going to delete.  If there are
// any pods preventing us from deleting, we return that list in an error.
func (o *drainOperation) getPodsForDeletion() (pods []v1.Pod, err error) {
	podList, err := o.client.ListPods(o.node)
	if err != nil {
		return pods, err
	}

	for _, pod := range podList.Items {
		podOk := true
		for _, filt := range []podFilter{
			mirrorPodFilter,
			daemonSetPodFilter,
		} {
			podOk = podOk && filt(pod)
		}
		if podOk {
			pods = append(pods, pod)
		}
	}
	return pods, nil
}

// deleteOrEvictPods deletes or evicts the pods on the api server
func (o *drainOperation) deleteOrEvictPods(pods []v1.Pod) error {
	if len(pods) == 0 {
		return nil
	}

	policyGroupVersion, err := o.client.SupportEviction()
	if err != nil {
		return err
	}

	if len(policyGroupVersion) > 0 {
		return o.evictPods(pods, policyGroupVersion)
	}
	return o.deletePods(pods)

}

func (o *drainOperation) evictPods(pods []v1.Pod, policyGroupVersion string) error {
	doneCh := make(chan bool, len(pods))
	errCh := make(chan error, 1)

	for _, pod := range pods {
		go func(pod v1.Pod, doneCh chan bool, errCh chan error) {
			var err error
			for {
				err = o.client.EvictPod(&pod, policyGroupVersion)
				if err == nil {
					break
				} else if apierrors.IsNotFound(err) {
					doneCh <- true
					return
				} else if apierrors.IsTooManyRequests(err) {
					time.Sleep(5 * time.Second)
				} else {
					errCh <- errors.Wrapf(err, "error when evicting pod %q", pod.Name)
					return
				}
			}
			podArray := []v1.Pod{pod}
			_, err = o.client.WaitForDelete(o.logger, podArray, true)
			if err == nil {
				doneCh <- true
			} else {
				errCh <- errors.Wrapf(err, "error when waiting for pod %q terminating", pod.Name)
			}
		}(pod, doneCh, errCh)
	}

	doneCount := 0
	for {
		select {
		case err := <-errCh:
			return err
		case <-doneCh:
			doneCount++
			if doneCount == len(pods) {
				return nil
			}
		case <-time.After(o.timeout):
			return errors.Errorf("Drain did not complete within %v", o.timeout)
		}
	}
}

func (o *drainOperation) deletePods(pods []v1.Pod) error {
	for _, pod := range pods {
		err := o.client.DeletePod(&pod)
		if err != nil && !apierrors.IsNotFound(err) {
			return err
		}
	}
	_, err := o.client.WaitForDelete(o.logger, pods, false)
	return err
}

func podOwnedByStatefulSet(pod v1.Pod) bool {
	ownerReferences := pod.ObjectMeta.OwnerReferences

	if ownerReferences == nil {
		return false
	}

	for _, owenerReference := range ownerReferences {
		if owenerReference.Kind == "StatefulSet" {
			return true
		}
	}

	return false
}

func waitForVolumesAttachedForPod(ctx context.Context, cancel context.CancelFunc, doneCh chan struct{}, pod v1.Pod, client armhelpers.KubernetesClient, logger *log.Entry, interval time.Duration) {
	defer func() {
		doneCh <- struct{}{}
	}()

	namespace := pod.Namespace
	podName := pod.Name
	logger.Infof("Start monitoring volume attachment for Pod %s...", podName)
	for {
		select {
		case <-ctx.Done():
			logger.Infof("Volume checking for Pod %s is canceled.", podName)
			return
		default:
			time.Sleep(interval)
			logger.Debugf("Sleep %s...", interval)

			if !podOwnedByStatefulSet(pod) {
				logger.Infof("Pod %s is not owned by StatefulSet, skip the volume attachment monitoring.", podName)
				return
			}

			podAfterDrain, _ := client.GetPod(namespace, podName)

			if podAfterDrain == nil || len(podAfterDrain.Status.Conditions) == 0 {
				logger.Debugf("Pod %s is not scheduled yet.", podName)
				continue
			}

			firstCondition := podAfterDrain.Status.Conditions[0]
			if string(firstCondition.Type) == "PodScheduled" && firstCondition.Reason == "Unschedulable" {
				logger.Infof("Pod %s cannot not be scheduled: %s Stop monitoring volume attachment.", podName, firstCondition.Message)
				if match, _ := regexp.MatchString("0/[0-9]+ nodes are available:.*exceed max volume count", firstCondition.Message); match {
					logger.Info("Nodes exceed max volume count, cancelling the volume attachment monitoring.")
					cancel()
				}
				return
			}

			nodeName := podAfterDrain.Spec.NodeName
			node, err := client.GetNode(nodeName)
			if err != nil {
				logger.Errorf("Failed to get node '%s': %s", nodeName, err.Error())
				return
			}

			done := true
			provisioner := "kubernetes.io/azure-disk"
			logger.Debugf("Found new pod %s scheduled on node '%s'", podName, nodeName)
			for _, volume := range podAfterDrain.Spec.Volumes {
				pvcSpec := volume.VolumeSource.PersistentVolumeClaim
				if pvcSpec != nil {
					pvc, err := client.GetPersistentVolumeClaim(namespace, pvcSpec.ClaimName)
					if err != nil {
						logger.Errorf("Failed to get PVC %s in namespace %s: %s", pvcSpec.ClaimName, namespace, err.Error())
						return
					}

					pv, err := client.GetPersistentVolume(pvc.Spec.VolumeName)
					if err != nil {
						logger.Errorf("Failed to get PV %s: %s", pvc.Spec.VolumeName, err.Error())
						return
					}

					logger.Debugf("Pod %s is using PV %s on node '%s'", podName, pv.Name, nodeName)
					if pv.ObjectMeta.Annotations["pv.kubernetes.io/provisioned-by"] == provisioner {
						attached := false
						for _, volumeAttached := range node.Status.VolumesAttached {
							if string(volumeAttached.Name) == provisioner+"/"+pv.Spec.AzureDisk.DataDiskURI {
								logger.Debugf("Volume %s is attached to node %s.", pv.Spec.AzureDisk.DataDiskURI, nodeName)
								attached = true
								break
							}
						}

						if !attached {
							logger.Debugf("Volume %s is not attached to node %s yet.", pv.Spec.AzureDisk.DataDiskURI, nodeName)
							done = false
							break
						}
					}
				}
			}

			if !done {
				logger.Debugf("Not all volumes for Pod %s are attached to new node %s.", podName, nodeName)
				continue
			} else {
				logger.Infof("All volumes for Pod %s are attached to new node %s.", podName, nodeName)
				return
			}
		}
	}
}

// WaitForDisksAttached waits for disks are re-attached to new nodes and volumes are ready to be used by pods.
func WaitForDisksAttached(podsForDeletion []v1.Pod, client armhelpers.KubernetesClient, logger *log.Entry, args ...time.Duration) error {
	if len(podsForDeletion) == 0 {
		return nil
	}

	timeout := time.Hour * 1
	interval := time.Minute * 1
	if len(args) > 0 {
		timeout = args[0]
	}
	if len(args) > 1 {
		interval = args[1]
	}

	ctx, cancel := context.WithCancel(context.Background())
	doneCh := make(chan struct{}, len(podsForDeletion))

	for _, pod := range podsForDeletion {
		go waitForVolumesAttachedForPod(ctx, cancel, doneCh, pod, client, logger, interval)
	}

	doneCount := 0
	for {
		select {
		case <-doneCh:
			doneCount++
			if doneCount == len(podsForDeletion) {
				logger.Info("Volume attachment check is done.")
				cancel()
				return nil
			}
		case <-time.After(timeout):
			err := fmt.Errorf("Volume attachment check time out")
			logger.Errorf(err.Error())
			cancel()
		}
	}
}
