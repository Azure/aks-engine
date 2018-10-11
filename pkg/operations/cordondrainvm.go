package operations

import (
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
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
func SafelyDrainNode(az armhelpers.AKSEngineClient, logger *log.Entry, masterURL, kubeConfig, nodeName string, timeout time.Duration) error {
	//get client using kubeconfig
	client, err := az.GetKubernetesClient(masterURL, kubeConfig, interval, timeout)
	if err != nil {
		return err
	}
	return SafelyDrainNodeWithClient(client, logger, nodeName, timeout)
}

// SafelyDrainNodeWithClient safely drains a node so that it can be deleted from the cluster
func SafelyDrainNodeWithClient(client armhelpers.KubernetesClient, logger *log.Entry, nodeName string, timeout time.Duration) error {
	//Mark the node unschedulable
	var node *v1.Node
	var err error
	for i := 0; i < cordonMaxRetries; i++ {
		node, err = client.GetNode(nodeName)
		if err != nil {
			return err
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
			return err
		}
		break
	}
	logger.Infof("Node %s has been marked unschedulable.", nodeName)

	//Evict pods in node
	drainOp := &drainOperation{client: client, node: node, logger: logger, timeout: timeout}
	return drainOp.deleteOrEvictPodsSimple()
}

func (o *drainOperation) deleteOrEvictPodsSimple() error {
	pods, err := o.getPodsForDeletion()
	if err != nil {
		return err
	}
	o.logger.Infof("%d pods need to be removed/deleted", len(pods))

	err = o.deleteOrEvictPods(pods)
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
