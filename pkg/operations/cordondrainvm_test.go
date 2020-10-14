// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Safely Drain node operation tests", func() {
	It("Should return error messages for failure to create kubernetes client", func() {
		err := SafelyDrainNode(&armhelpers.MockAKSEngineClient{FailGetKubernetesClient: true}, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should return error messages for Failure to get node ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.FailGetNode = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should retry on resource conflict when updating node ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		i := 3
		mockClient.MockKubernetesClient.UpdateNodeFunc = func(node *v1.Node) (*v1.Node, error) {
			if i > 0 {
				i--
				return node, errors.New(kubernetesOptimisticLockErrorMsg)
			}
			return node, nil
		}
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should return error messages for Failure to update node ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.FailUpdateNode = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should return error messages for Failure to list pods ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.FailListPods = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should return error messages for Failure to check support eviction ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.PodsList = &v1.PodList{Items: []v1.Pod{{}}}
		mockClient.MockKubernetesClient.FailSupportEviction = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should return error messages for Failure to delete pod ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.PodsList = &v1.PodList{Items: []v1.Pod{{}}}
		mockClient.MockKubernetesClient.FailDeletePod = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should return error messages for Failure to Evict Pod ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.PodsList = &v1.PodList{Items: []v1.Pod{{}}}
		mockClient.MockKubernetesClient.ShouldSupportEviction = true
		mockClient.MockKubernetesClient.FailEvictPod = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should return error messages for Failure to wait for delete in delete path ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.PodsList = &v1.PodList{Items: []v1.Pod{{}}}
		mockClient.MockKubernetesClient.ShouldSupportEviction = true
		mockClient.MockKubernetesClient.FailWaitForDelete = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should return error messages for Failure to wait for delete in eviction path ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.PodsList = &v1.PodList{Items: []v1.Pod{{}}}
		mockClient.MockKubernetesClient.ShouldSupportEviction = false
		mockClient.MockKubernetesClient.FailWaitForDelete = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).Should(HaveOccurred())
	})
	It("Should not return error in valid eviction path ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.PodsList = &v1.PodList{Items: []v1.Pod{{}}}
		mockClient.MockKubernetesClient.ShouldSupportEviction = true
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should not return error in valid delete path ", func() {
		mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
		mockClient.MockKubernetesClient.PodsList = &v1.PodList{Items: []v1.Pod{{}}}
		mockClient.MockKubernetesClient.ShouldSupportEviction = false
		err := SafelyDrainNode(mockClient, log.NewEntry(log.New()), "http://bad.com/", "bad", "node", time.Minute)
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should not return daemonSet pods in the list of pods to delete/evict", func() {
		mockClient := &armhelpers.MockKubernetesClient{}
		truebool := true
		mockClient.PodsList = &v1.PodList{
			Items: []v1.Pod{
				{}, //unreplicated pod
				{
					ObjectMeta: metav1.ObjectMeta{
						OwnerReferences: []metav1.OwnerReference{
							{
								Kind:       "DaemonSet",
								Controller: &truebool,
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						OwnerReferences: []metav1.OwnerReference{
							{
								Kind:       "ReplicaSet",
								Controller: &truebool,
							},
						},
					},
				},
			},
		}
		mockClient.ShouldSupportEviction = true
		o := drainOperation{client: mockClient}
		pods, err := o.getPodsForDeletion()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(len(pods)).Should(Equal(2))
	})
})

var _ = Describe("Wait for disks attached operation tests", func() {
	var mockKubernetesClient *armhelpers.MockKubernetesClient
	var buffer bytes.Buffer
	var loggerEntry *log.Entry
	podStore := podStore()

	BeforeEach(func() {
		mockKubernetesClient = &armhelpers.MockKubernetesClient{}
		mockKubernetesClient.GetPodFunc = func(namespace, name string) (*v1.Pod, error) {
			return podStore[name], nil
		}
		mockKubernetesClient.GetPvcFunc = func(namespace, name string) (*v1.PersistentVolumeClaim, error) {
			pvc := &v1.PersistentVolumeClaim{}
			pvc.ObjectMeta.Name = name
			pvc.ObjectMeta.Namespace = namespace
			pvc.Spec.VolumeName = "pvc-" + name
			return pvc, nil
		}

		logger := log.New()
		logger.SetLevel(log.DebugLevel)
		logger.Out = &buffer
		logger.Formatter = new(log.JSONFormatter)
		loggerEntry = log.NewEntry(logger)
	})

	It("Should log error messages for Failure to get Node and complete quickly", func() {
		mockKubernetesClient.FailGetNode = true
		pod := *podStore["statefulset-pod-with-one-pvc"]
		podsForDeletion := []v1.Pod{pod}

		err := WaitForDisksAttached(podsForDeletion, mockKubernetesClient, loggerEntry, time.Second*3, time.Second*1)

		logRecords := readLogs(buffer)
		Expect(logRecords[len(logRecords)-2].Msg).To(Equal(fmt.Sprintf("Failed to get node '%s': %s", pod.Spec.NodeName, "GetNode failed")))
		Expect(logRecords[len(logRecords)-1].Msg).To(Equal("Volume attachment check is done."))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should log error messages for Failure to get PVC and complete quickly", func() {
		mockKubernetesClient.FailGetPvc = true
		pod := *podStore["statefulset-pod-with-one-pvc"]
		podsForDeletion := []v1.Pod{pod}

		err := WaitForDisksAttached(podsForDeletion, mockKubernetesClient, loggerEntry, time.Second*3, time.Second*1)

		logRecords := readLogs(buffer)
		Expect(logRecords[len(logRecords)-2].Msg).To(Equal(fmt.Sprintf("Failed to get PVC %s in namespace %s: %s", pod.Spec.Volumes[0].VolumeSource.PersistentVolumeClaim.ClaimName, pod.Namespace, "GetPersistentVolumeClaim failed")))
		Expect(logRecords[len(logRecords)-1].Msg).To(Equal("Volume attachment check is done."))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should log error messages for Failure to get PV and complete quickly", func() {
		mockKubernetesClient.FailGetPv = true
		pod := *podStore["statefulset-pod-with-one-pvc"]
		podsForDeletion := []v1.Pod{pod}

		err := WaitForDisksAttached(podsForDeletion, mockKubernetesClient, loggerEntry, time.Second*3, time.Second*1)

		logRecords := readLogs(buffer)
		Expect(logRecords[len(logRecords)-2].Msg).To(Equal(fmt.Sprintf("Failed to get PV %s: %s", "pvc-"+pod.Spec.Volumes[0].PersistentVolumeClaim.ClaimName, "GetPersistentVolume failed")))
		Expect(logRecords[len(logRecords)-1].Msg).To(Equal("Volume attachment check is done."))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should cancel the volume attachment check when time out", func() {
		mockKubernetesClient.FailGetPod = true
		pod := *podStore["statefulset-pod-with-one-pvc"]
		podsForDeletion := []v1.Pod{pod}

		err := WaitForDisksAttached(podsForDeletion, mockKubernetesClient, loggerEntry, time.Second*3, time.Second*1)

		logRecords := readLogs(buffer)
		var errMsg string
		for _, record := range logRecords {
			if record.Level == "error" {
				errMsg = record.Msg
			}
		}
		Expect(logRecords[2].Msg).To(Equal(fmt.Sprintf("Pod %s is not scheduled yet.", pod.Name)))
		Expect(errMsg).To(Equal("Volume attachment check time out"))
		Expect(logRecords[len(logRecords)-2].Msg).To(Equal(fmt.Sprintf("Volume checking for Pod %s is canceled.", pod.Name)))
		Expect(logRecords[len(logRecords)-1].Msg).To(Equal("Volume attachment check is done."))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should skip the volume attachment monitoring if the pod is not owned by a StatefulSet", func() {
		pod := *podStore["deployment-pod"]
		podsForDeletion := []v1.Pod{pod}

		err := WaitForDisksAttached(podsForDeletion, mockKubernetesClient, loggerEntry, time.Second*3, time.Second*1)

		logRecords := readLogs(buffer)
		Expect(logRecords[len(logRecords)-2].Msg).To(Equal(fmt.Sprintf("Pod %s is not owned by StatefulSet, skip the volume attachment monitoring.", pod.Name)))
		Expect(logRecords[len(logRecords)-1].Msg).To(Equal("Volume attachment check is done."))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should cancel the volume attachment monitoring if nodes exceed max volume count", func() {
		unschedulablePod := *podStore["statefulset-pod-with-unschedulable-condition"]
		normalPod := *podStore["statefulset-pod-with-one-pvc"]

		mockKubernetesClient.GetPodFunc = func(namespace, name string) (*v1.Pod, error) {
			if name == normalPod.Name {
				return nil, errors.New("GetNode failed")
			}
			return podStore[name], nil
		}

		podsForDeletion := []v1.Pod{unschedulablePod, normalPod}

		err := WaitForDisksAttached(podsForDeletion, mockKubernetesClient, loggerEntry, time.Second*3, time.Second*1)

		logRecords := readLogs(buffer)
		msgExist := false
		for _, record := range logRecords {
			if record.Msg == "Nodes exceed max volume count, cancelling the volume attachment monitoring." {
				msgExist = true
			}
		}
		Expect(msgExist).To(Equal(true))
		Expect(logRecords[len(logRecords)-2].Msg).To(Equal(fmt.Sprintf("Volume checking for Pod %s is canceled.", normalPod.Name)))
		Expect(logRecords[len(logRecords)-1].Msg).To(Equal("Volume attachment check is done."))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should succeed without error if all pods and functions are valid", func() {
		diskPrefix := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/k8scluster/providers/Microsoft.Compute/disks/k8scluster-dynamic-"
		mockKubernetesClient.GetPvFunc = func(name string) (*v1.PersistentVolume, error) {
			var pv *v1.PersistentVolume
			pvString := fmt.Sprintf(`{
				"metadata": {
					"name": "%s",
					"annotations": {
						"pv.kubernetes.io/provisioned-by": "kubernetes.io/azure-disk"
					}
				},
				"spec": {
					"azureDisk": {
						"diskURI": "%s%s"
					}
				}
			}`, name, diskPrefix, name)
			json.Unmarshal([]byte(pvString), &pv)
			return pv, nil
		}

		statefulsetPodWithOnePvc := *podStore["statefulset-pod-with-one-pvc"]
		statefulsetPodWithTwoPvc := *podStore["statefulset-pod-with-two-pvc"]
		deploymentPod := *podStore["deployment-pod"]

		state := make(map[string]int)
		provisioner := "kubernetes.io/azure-disk"
		mockKubernetesClient.GetNodeFunc = func(name string) (*v1.Node, error) {
			node := &v1.Node{}
			if count, ok := state[name]; ok {
				if count > 2 {
					node.Status = v1.NodeStatus{
						VolumesAttached: []v1.AttachedVolume{
							v1.AttachedVolume{Name: v1.UniqueVolumeName(provisioner + "/" + diskPrefix + "pvc-" + statefulsetPodWithOnePvc.Spec.Volumes[0].PersistentVolumeClaim.ClaimName)},
							v1.AttachedVolume{Name: v1.UniqueVolumeName(provisioner + "/" + diskPrefix + "pvc-" + statefulsetPodWithTwoPvc.Spec.Volumes[0].PersistentVolumeClaim.ClaimName)},
							v1.AttachedVolume{Name: v1.UniqueVolumeName(provisioner + "/" + diskPrefix + "pvc-" + statefulsetPodWithTwoPvc.Spec.Volumes[1].PersistentVolumeClaim.ClaimName)},
						},
					}
				} else {
					state[name]++
				}
			} else {
				state[name] = 1
			}
			return node, nil
		}

		podsForDeletion := []v1.Pod{statefulsetPodWithOnePvc, statefulsetPodWithTwoPvc, deploymentPod}
		err := WaitForDisksAttached(podsForDeletion, mockKubernetesClient, loggerEntry, time.Second*10, time.Second*1)

		logRecords := readLogs(buffer)

		var deploymentPodSucceed, statefulsetPod1Succeed, statefulsetPod2Succeed bool
		for _, record := range logRecords {
			switch record.Msg {
			case fmt.Sprintf("Pod %s is not owned by StatefulSet, skip the volume attachment monitoring.", deploymentPod.Name):
				deploymentPodSucceed = true
			case fmt.Sprintf("All volumes for Pod %s are attached to new node %s.", statefulsetPodWithOnePvc.Name, statefulsetPodWithOnePvc.Spec.NodeName):
				statefulsetPod1Succeed = true
			case fmt.Sprintf("All volumes for Pod %s are attached to new node %s.", statefulsetPodWithTwoPvc.Name, statefulsetPodWithTwoPvc.Spec.NodeName):
				statefulsetPod2Succeed = true
			}
		}

		Expect(deploymentPodSucceed).To(Equal(true))
		Expect(statefulsetPod1Succeed).To(Equal(true))
		Expect(statefulsetPod2Succeed).To(Equal(true))
		Expect(logRecords[len(logRecords)-1].Msg).To(Equal("Volume attachment check is done."))
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		fmt.Println(string(buffer.Bytes()))
		buffer.Reset()
	})
})

type Fields struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

func readLogs(buffer bytes.Buffer) (logRecords []Fields) {
	dec := json.NewDecoder(bytes.NewReader(buffer.Bytes()))
	for {
		var record Fields

		err := dec.Decode(&record)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		logRecords = append(logRecords, record)
	}
	return logRecords
}

func podStore() map[string]*v1.Pod {
	podsString := `{
		"statefulset-pod-with-one-pvc": {
			"metadata": {
				"name": "statefulset-pod-with-one-pvc",
				"ownerReferences": [
					{
						"kind": "StatefulSet"
					}
				]
			},
			"spec": {
				"nodeName": "k8s-linuxpool-12345678-0",
				"volumes": [
					{
						"name": "persistent-storage-statefulset",
						"persistentVolumeClaim": {
							"claimName": "persistent-storage"
						}
					}
				]
			},
			"status": {
				"conditions": [
					{
						"status": "True",
						"type": "PodScheduled"
					}
				]
			}
		},
		"statefulset-pod-with-two-pvc": {
			"metadata": {
				"name": "statefulset-pod-with-two-pvc",
				"ownerReferences": [
					{
						"kind": "StatefulSet"
					}
				]
			},
			"spec": {
				"nodeName": "k8s-linuxpool-12345678-0",
				"volumes": [
					{
						"name": "persistent-storage-statefulset1",
						"persistentVolumeClaim": {
							"claimName": "persistent-storage1"
						}
					},
					{
						"name": "persistent-storage-statefulset2",
						"persistentVolumeClaim": {
							"claimName": "persistent-storage2"
						}
					}
				]
			},
			"status": {
				"conditions": [
					{
						"status": "True",
						"type": "Initialized"
					}
				]
			}
		},
		"statefulset-pod-with-unschedulable-condition": {
			"metadata": {
				"name": "statefulset-pod-with-unschedulable-condition",
				"ownerReferences": [
					{
						"kind": "StatefulSet"
					}
				]
			},
			"status": {
				"conditions": [
					{
						"reason": "Unschedulable",
						"type": "PodScheduled",
						"message": "0/3 nodes are available: 3 node(s) exceed max volume count."
					}
				]
			}
		},
		"deployment-pod": {
			"metadata": {
				"name": "deployment-pod",
				"ownerReferences": [
					{
						"kind": "ReplicaSet"
					}
				]
			}
		}
	}`

	podStore := make(map[string]*v1.Pod)
	json.Unmarshal([]byte(podsString), &podStore)
	return podStore
}
