// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	"fmt"
	"testing"

	mock "github.com/Azure/aks-engine/cmd/rotatecerts/internal/mock_internal"
	"github.com/Azure/aks-engine/pkg/api/common"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errAPIGeneric  = errors.New("generic api error")
	errAPINotFound = &apierrors.StatusError{
		ErrStatus: metav1.Status{
			Reason: metav1.StatusReasonNotFound,
		},
	}
)

func TestPauseClusterAutoscaler(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("GetDeployment fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).Times(1)

		resume, err := PauseClusterAutoscaler(mock)
		g.Expect(resume).NotTo(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
		msg := fmt.Sprintf("getting %s deployment: %s", common.ClusterAutoscalerAddonName, errAPIGeneric)
		g.Expect(errors.Unwrap(resume()).Error()).To(Equal(msg))
	})

	t.Run("Deployment does not exist", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(nil, errAPINotFound).Times(1)
		mock.EXPECT().PatchDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

		resume, err := PauseClusterAutoscaler(mock)
		g.Expect(resume).ToNot(BeNil())
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resume()).ToNot(HaveOccurred())
	})

	t.Run("Deployment replica count is zero", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		var replicas int32
		deploy := appsv1.Deployment{}
		deploy.Spec.Replicas = &replicas
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(&deploy, nil).Times(1)
		mock.EXPECT().PatchDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

		resume, err := PauseClusterAutoscaler(mock)
		g.Expect(resume).ToNot(BeNil())
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resume()).ToNot(HaveOccurred())
	})

	t.Run("Deployment scale down fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		var replicas int32 = 1
		deploy := appsv1.Deployment{}
		deploy.Spec.Replicas = &replicas
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(&deploy, nil).Times(1)
		mock.EXPECT().PatchDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Return(&deploy, errAPIGeneric).Times(1)

		resume, err := PauseClusterAutoscaler(mock)
		g.Expect(resume).NotTo(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
		msg := fmt.Sprintf("applying patch to %s deployment: %s", common.ClusterAutoscalerAddonName, errAPIGeneric)
		g.Expect(errors.Unwrap(resume()).Error()).To(Equal(msg))
	})

	t.Run("Deployment scale ok", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		var replicas int32 = 1
		deploy := appsv1.Deployment{}
		deploy.Spec.Replicas = &replicas
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(&deploy, nil).Times(1)
		mock.EXPECT().PatchDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Return(&deploy, nil).Times(2)

		resume, err := PauseClusterAutoscaler(mock)
		g.Expect(resume).ToNot(BeNil())
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resume()).ToNot(HaveOccurred())
	})

	t.Run("Deployment scale up fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		var replicas int32 = 1
		deploy := appsv1.Deployment{}
		deploy.Spec.Replicas = &replicas
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(&deploy, nil).Times(1)
		gomock.InOrder(
			mock.EXPECT().PatchDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Return(&deploy, nil),
			mock.EXPECT().PatchDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Return(&deploy, errAPIGeneric),
		)

		resume, err := PauseClusterAutoscaler(mock)
		g.Expect(resume).ToNot(BeNil())
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resume()).To(HaveOccurred())
	})
}

func TestRolloutDeployments(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("ListDeployment fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDeployments(metav1.NamespaceAll, metav1.ListOptions{}).Return(nil, errAPIGeneric).Times(1)

		err := rolloutDeployments(mock)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("All deployments are patched", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		d1 := appsv1.Deployment{}
		d1.Namespace = "ns1"
		d1.Name = "d1"
		d2 := appsv1.Deployment{}
		d2.Namespace = "ns2"
		d2.Name = "d2"
		list := &appsv1.DeploymentList{Items: []appsv1.Deployment{d1, d2}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDeployments(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)
		mock.EXPECT().PatchDeployment("ns1", "d1", gomock.Any()).Return(nil, nil).Times(1)
		mock.EXPECT().PatchDeployment("ns2", "d2", gomock.Any()).Return(nil, nil).Times(1)

		err := rolloutDeployments(mock)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("Return error if patch deployment fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		d1 := appsv1.Deployment{}
		d1.Namespace = "ns1"
		d1.Name = "d1"
		list := &appsv1.DeploymentList{Items: []appsv1.Deployment{d1}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDeployments(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)
		mock.EXPECT().PatchDeployment("ns1", "d1", gomock.Any()).Return(nil, errAPIGeneric).Times(1)

		err := rolloutDeployments(mock)
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})
}

func TestRolloutDaemonSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("ListDaemonSets fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(metav1.NamespaceAll, metav1.ListOptions{}).Return(nil, errAPIGeneric).Times(1)

		err := rolloutDaemonSets(mock)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("All daemonsets are patched, SA deleted", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ds1 := appsv1.DaemonSet{}
		ds1.Namespace = "ns1"
		ds1.Name = "ds1"
		ds2 := appsv1.DaemonSet{}
		ds2.Namespace = "ns2"
		ds2.Name = "ds2"
		list := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{ds1, ds2}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)
		mock.EXPECT().PatchDaemonSet("ns1", "ds1", gomock.Any()).Return(nil, nil).Times(1)
		mock.EXPECT().PatchDaemonSet("ns2", "ds2", gomock.Any()).Return(nil, nil).Times(1)

		err := rolloutDaemonSets(mock)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("Return error if patch daemonset fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ds1 := appsv1.DaemonSet{}
		ds1.Namespace = "ns1"
		ds1.Name = "ds1"
		list := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{ds1}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)
		mock.EXPECT().PatchDaemonSet("ns1", "ds1", gomock.Any()).Return(nil, errAPIGeneric).Times(1)

		err := rolloutDaemonSets(mock)
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})
}

func TestDeleteSATokens(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("List ServiceAccounts fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.ServiceAccountList{}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, errAPIGeneric).Times(1)

		err := deleteSATokens(mock)
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})

	t.Run("No ServiceAccounts in namespace", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.ServiceAccountList{}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)

		err := deleteSATokens(mock)
		g.Expect(err).To(BeNil())
	})

	t.Run("Secret to delete not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}
		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)
		mock.EXPECT().DeleteSecret(gomock.Any()).Return(errAPINotFound).Times(2)

		err := deleteSATokens(mock)
		g.Expect(err).To(BeNil())
	})

	t.Run("Delete Secret fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}
		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)
		mock.EXPECT().DeleteSecret(gomock.Any()).Return(errAPIGeneric).Times(1)

		err := deleteSATokens(mock)
		g.Expect(err).ToNot(BeNil())
	})

	t.Run("Secrets deleted", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}

		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(metav1.NamespaceAll, metav1.ListOptions{}).Return(list, nil).Times(1)
		mock.EXPECT().DeleteSecret(gomock.Any()).Return(nil).Times(2)

		err := deleteSATokens(mock)
		g.Expect(err).To(BeNil())
	})
}
