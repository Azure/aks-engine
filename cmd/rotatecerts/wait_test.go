// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	"fmt"
	"testing"
	"time"

	mock "github.com/Azure/aks-engine/cmd/rotatecerts/internal/mock_internal"
	"github.com/Azure/aks-engine/pkg/helpers/ssh"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestWaitForNodesCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	falseCond := func(*v1.NodeList) bool { return false }
	trueCond := func(*v1.NodeList) bool { return true }

	t.Run("ListNodes fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListNodes().Return(nil, errAPIGeneric).AnyTimes()

		_, err := waitForNodesCondition(mock, falseCond, 1, 1*time.Second, 1*time.Minute)
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})

	t.Run("Node condition met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.NodeList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListNodes().Return(list, nil).AnyTimes()

		nl, err := waitForNodesCondition(mock, trueCond, 2, 500*time.Millisecond, 2*time.Second)
		g.Expect(nl).NotTo(BeNil())
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("Node condition not met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.NodeList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListNodes().Return(list, nil).AnyTimes()

		_, err := waitForNodesCondition(mock, falseCond, 1, 500*time.Millisecond, 1*time.Second)
		g.Expect(err).To(HaveOccurred())
		g.Expect(fmt.Sprint(err)).To(Equal("timed out waiting for the condition"))
	})
}

func TestAllExpectedNodesReadyCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	nodeReadyCondition := v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionTrue}
	nodeNotReadyCondition := v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionFalse}

	t.Run("All expected nodes ready", func(t *testing.T) {
		expected := make(map[string]*ssh.RemoteHost)
		expected["n1"] = &ssh.RemoteHost{URI: "n1"}
		n1 := v1.Node{}
		n1.Name = "n1"
		n1.Status.Conditions = []v1.NodeCondition{nodeReadyCondition}
		nl := &v1.NodeList{Items: []v1.Node{n1}}
		cond := allExpectedNodesReadyCondition(expected)(nl)
		g.Expect(cond).To(BeTrue())
	})

	t.Run("Some expected nodes ready", func(t *testing.T) {
		expected := make(map[string]*ssh.RemoteHost)
		expected["n1"] = &ssh.RemoteHost{URI: "n1"}
		n1 := v1.Node{}
		n1.Name = "n1"
		n1.Status.Conditions = []v1.NodeCondition{nodeNotReadyCondition}
		nl := &v1.NodeList{Items: []v1.Node{n1}}
		cond := allExpectedNodesReadyCondition(expected)(nl)
		g.Expect(cond).To(BeFalse())
	})

	t.Run("Expected node missing", func(t *testing.T) {
		expected := make(map[string]*ssh.RemoteHost)
		expected["n1"] = &ssh.RemoteHost{URI: "n1"}
		n1 := v1.Node{}
		n1.Name = "n2"
		n1.Status.Conditions = []v1.NodeCondition{nodeNotReadyCondition}
		nl := &v1.NodeList{Items: []v1.Node{n1}}
		cond := allExpectedNodesReadyCondition(expected)(nl)
		g.Expect(cond).To(BeFalse())
	})
}

func TestWaitForPodsCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	falseCond := func(*v1.PodList) error { return nil }
	trueCond := func(*v1.PodList) error { return errors.New("condition not met") }
	nop := func(*v1.PodList) {}

	t.Run("ListPods fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListPods(gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).AnyTimes()

		err := waitForPodsCondition(mock, "ns", falseCond, 2, 500*time.Millisecond, 200*time.Second, nop)
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})

	t.Run("Pod condition met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.PodList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListPods(gomock.Any(), gomock.Any()).Return(list, nil).AnyTimes()

		err := waitForPodsCondition(mock, "ns", falseCond, 2, 500*time.Millisecond, 200*time.Second, nop)
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("Pod condition not met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.PodList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListPods(gomock.Any(), gomock.Any()).Return(list, nil).AnyTimes()

		err := waitForPodsCondition(mock, "ns", trueCond, 1, 500*time.Millisecond, 1*time.Second, nop)
		g.Expect(err).To(HaveOccurred())
		g.Expect(fmt.Sprint(err)).To(Equal("condition successesCount: 0: condition not met"))
	})
}

func TestAllPodsReadyCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	podRunningCondition := v1.ContainerStatus{
		State: v1.ContainerState{Running: &v1.ContainerStateRunning{}},
		Ready: true,
	}
	podNotRunningCondition := v1.ContainerStatus{
		State: v1.ContainerState{Waiting: &v1.ContainerStateWaiting{}},
	}

	t.Run("Pod not running", func(t *testing.T) {
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodPending
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allPodsReadyCondition(pl)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Container not running", func(t *testing.T) {
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podNotRunningCondition}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allPodsReadyCondition(pl)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Container running", func(t *testing.T) {
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podRunningCondition}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allPodsReadyCondition(pl)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("No pods", func(t *testing.T) {
		err := allPodsReadyCondition(&v1.PodList{})
		g.Expect(err).ToNot(HaveOccurred())
	})
}

func TestAllExpectedPodsReadyCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	podRunningCondition := v1.ContainerStatus{
		State: v1.ContainerState{Running: &v1.ContainerStateRunning{}},
		Ready: true,
	}
	podNotRunningCondition := v1.ContainerStatus{
		State: v1.ContainerState{Waiting: &v1.ContainerStateWaiting{}},
	}

	t.Run("All expected pods ready", func(t *testing.T) {
		expected := []string{"p1"}
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podRunningCondition}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsReadyCondition(expected)(pl)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("Some expected pods ready", func(t *testing.T) {
		expected := []string{"p1"}
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podNotRunningCondition}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsReadyCondition(expected)(pl)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Expected pod missing", func(t *testing.T) {
		expected := []string{"p2"}
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podRunningCondition}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsReadyCondition(expected)(pl)
		g.Expect(err).To(HaveOccurred())
	})
}

func TestAllExpectedPodsRestartedCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	restartTime := time.Now()
	podRunningBeforeRestartTime := v1.ContainerStatus{
		State: v1.ContainerState{
			Running: &v1.ContainerStateRunning{
				StartedAt: metav1.NewTime(restartTime.Add(-1 * time.Minute)),
			}},
		Ready: true,
	}
	fmt.Println(restartTime)
	fmt.Println(metav1.NewTime(restartTime.Add(-1 * time.Minute)))
	fmt.Println(metav1.NewTime(restartTime.Add(1 * time.Minute)))
	podRunningAfterRestartTime := v1.ContainerStatus{
		State: v1.ContainerState{
			Running: &v1.ContainerStateRunning{
				StartedAt: metav1.NewTime(restartTime.Add(1 * time.Minute)),
			}},
		Ready: true,
	}
	podNotRunningCondition := v1.ContainerStatus{
		State: v1.ContainerState{Waiting: &v1.ContainerStateWaiting{}},
	}

	t.Run("Expected pods restarted", func(t *testing.T) {
		expected := []string{"p1"}
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podRunningAfterRestartTime}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsRestartedCondition(expected, restartTime)(pl)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("Some expected pods not restarted", func(t *testing.T) {
		expected := []string{"p1"}
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podRunningBeforeRestartTime}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsRestartedCondition(expected, restartTime)(pl)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Some expected pods not ready", func(t *testing.T) {
		expected := []string{"p1"}
		p1 := v1.Pod{}
		p1.Name = "p1"
		p1.Status.Phase = v1.PodRunning
		p1.Status.ContainerStatuses = []v1.ContainerStatus{podNotRunningCondition}
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsRestartedCondition(expected, restartTime)(pl)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Expected pod missing", func(t *testing.T) {
		expected := []string{"p2"}
		p1 := v1.Pod{}
		p1.Name = "p1"
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsRestartedCondition(expected, restartTime)(pl)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Empty input list", func(t *testing.T) {
		expected := []string{}
		p1 := v1.Pod{}
		p1.Name = "p1"
		pl := &v1.PodList{Items: []v1.Pod{p1}}
		err := allExpectedPodsRestartedCondition(expected, restartTime)(pl)
		g.Expect(err).ToNot(HaveOccurred())
	})
}

func TestWaitForDaemonSetCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	falseCond := func(*appsv1.DaemonSetList) error { return nil }
	trueCond := func(*appsv1.DaemonSetList) error { return errors.New("condition not met") }

	t.Run("ListDaemonSets fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).AnyTimes()

		err := waitForDaemonSetCondition(mock, "ns", falseCond, 2, 500*time.Millisecond, 200*time.Second)
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})

	t.Run("DaemonSet condition met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &appsv1.DaemonSetList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(gomock.Any(), gomock.Any()).Return(list, nil).AnyTimes()

		err := waitForDaemonSetCondition(mock, "ns", falseCond, 2, 500*time.Millisecond, 200*time.Second)
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("DaemonSet condition not met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &appsv1.DaemonSetList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(gomock.Any(), gomock.Any()).Return(list, nil).AnyTimes()

		err := waitForDaemonSetCondition(mock, "ns", trueCond, 1, 500*time.Millisecond, 1*time.Second)
		g.Expect(err).To(HaveOccurred())
		g.Expect(fmt.Sprint(err)).To(Equal("condition successesCount: 0: condition not met"))
	})
}

func TestAllDaemontSetReplicasUpdatedCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("Desired replicas updated and available", func(t *testing.T) {
		d1 := appsv1.DaemonSet{}
		d1.Name = "p1"
		d1.Status.DesiredNumberScheduled = 2
		d1.Status.CurrentNumberScheduled = 2
		d1.Status.UpdatedNumberScheduled = 2
		dl := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{d1}}
		err := allDaemontSetReplicasUpdatedCondition(dl)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("Not all updated replicas are available", func(t *testing.T) {
		d1 := appsv1.DaemonSet{}
		d1.Name = "p1"
		d1.Status.DesiredNumberScheduled = 2
		d1.Status.CurrentNumberScheduled = 1
		d1.Status.UpdatedNumberScheduled = 2
		dl := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{d1}}
		err := allDaemontSetReplicasUpdatedCondition(dl)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("at least one daemonset is still updating replicas: [p1]"))
	})

	t.Run("Not all replicas updated their template", func(t *testing.T) {
		d1 := appsv1.DaemonSet{}
		d1.Name = "p1"
		d1.Status.DesiredNumberScheduled = 2
		d1.Status.CurrentNumberScheduled = 2
		d1.Status.UpdatedNumberScheduled = 1
		dl := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{d1}}
		err := allDaemontSetReplicasUpdatedCondition(dl)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("at least one daemonset is still updating replicas: [p1]"))
	})
}

func TestWaitForDeploymentCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	falseCond := func(*appsv1.DeploymentList) error { return nil }
	trueCond := func(*appsv1.DeploymentList) error { return errors.New("condition not met") }

	t.Run("ListDaemonSets fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).AnyTimes()

		err := waitForDeploymentCondition(mock, "ns", falseCond, 2, 500*time.Millisecond, 200*time.Second)
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})

	t.Run("Deployment condition met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &appsv1.DeploymentList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return(list, nil).AnyTimes()

		err := waitForDeploymentCondition(mock, "ns", falseCond, 2, 500*time.Millisecond, 200*time.Second)
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("Deployment condition not met within timeout period", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &appsv1.DeploymentList{}
		mock := mock.NewMockClient(mockCtrl)
		mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return(list, nil).AnyTimes()

		err := waitForDeploymentCondition(mock, "ns", trueCond, 1, 500*time.Millisecond, 1*time.Second)
		g.Expect(err).To(HaveOccurred())
		g.Expect(fmt.Sprint(err)).To(Equal("condition successesCount: 0: condition not met"))
	})
}

func TestAllDeploymentReplicasUpdatedCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("Desired replicas updated and available", func(t *testing.T) {
		d1 := appsv1.Deployment{}
		d1.Name = "p1"
		d1.Status.Replicas = 2
		d1.Status.AvailableReplicas = 2
		d1.Status.UpdatedReplicas = 2
		dl := &appsv1.DeploymentList{Items: []appsv1.Deployment{d1}}
		err := allDeploymentReplicasUpdatedCondition(dl)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("Not all updated replicas are available", func(t *testing.T) {
		d1 := appsv1.Deployment{}
		d1.Name = "p1"
		d1.Status.Replicas = 2
		d1.Status.AvailableReplicas = 1
		d1.Status.UpdatedReplicas = 2
		dl := &appsv1.DeploymentList{Items: []appsv1.Deployment{d1}}
		err := allDeploymentReplicasUpdatedCondition(dl)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("at least one deployment is still updating replicas: [p1]"))
	})

	t.Run("Not all replicas updated their template", func(t *testing.T) {
		d1 := appsv1.Deployment{}
		d1.Name = "p1"
		d1.Status.Replicas = 2
		d1.Status.AvailableReplicas = 2
		d1.Status.UpdatedReplicas = 1
		dl := &appsv1.DeploymentList{Items: []appsv1.Deployment{d1}}
		err := allDeploymentReplicasUpdatedCondition(dl)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("at least one deployment is still updating replicas: [p1]"))
	})
}
