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

		var replicas int32 = 0
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

func TestDeleteDeploymentSATokensAndForceRollout(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	deleteSATokens := func(saMap map[string]bool) func(name string) error {
		return func(name string) error {
			saMap[name] = false
			return nil
		}
	}

	t.Run("ListDeployment fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).Times(1)

		err := deleteDeploymentSATokensAndForceRollout(mock, "ns", deleteSATokens(nil))
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("All deployments are patched, SA deleted", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		noSA := appsv1.Deployment{}
		noSA.Name = "noSA"
		noSA.Spec.Template.Spec.ServiceAccountName = "noSA"
		hasSA := appsv1.Deployment{}
		hasSA.Name = "hasSA"
		hasSA.Spec.Template.Spec.ServiceAccountName = "hasSA"
		list := &appsv1.DeploymentList{Items: []appsv1.Deployment{hasSA, noSA}}

		saMap := make(map[string]bool)
		saMap[hasSA.Spec.Template.Spec.ServiceAccountName] = true
		saMap["random"] = true

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().PatchDeployment(gomock.Any(), "noSA", gomock.Any()).Return(nil, nil).Times(1)
		mock.EXPECT().PatchDeployment(gomock.Any(), "hasSA", gomock.Any()).Return(nil, nil).Times(1)

		err := deleteDeploymentSATokensAndForceRollout(mock, "ns", deleteSATokens(saMap))
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(saMap["random"]).To(BeTrue())
		g.Expect(saMap["hasSA"]).To(BeFalse())
	})

	t.Run("Return error if delete SA fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		hasSA := appsv1.Deployment{}
		hasSA.Spec.Template.Spec.ServiceAccountName = "hasSA"
		list := &appsv1.DeploymentList{Items: []appsv1.Deployment{hasSA}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)

		err := deleteDeploymentSATokensAndForceRollout(mock, "ns", func(name string) error {
			return errors.New("Delete SA failed")
		})
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Return error if patch deployment fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		hasSA := appsv1.Deployment{}
		hasSA.Spec.Template.Spec.ServiceAccountName = "hasSA"
		list := &appsv1.DeploymentList{Items: []appsv1.Deployment{hasSA}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().PatchDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).Times(1)

		err := deleteDeploymentSATokensAndForceRollout(mock, "ns", func(_ string) error { return nil })
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})
}

func TestDeleteDaemonSetSATokensAndForceRollout(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	deleteSATokens := func(saMap map[string]bool) func(name string) error {
		return func(name string) error {
			saMap[name] = false
			return nil
		}
	}

	t.Run("ListDaemonSets fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).Times(1)

		err := deleteDaemonSetSATokensAndForceRollout(mock, "ns", deleteSATokens(nil))
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("All daemonsets are patched, SA deleted", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		noSA := appsv1.DaemonSet{}
		noSA.Name = "noSA"
		noSA.Spec.Template.Spec.ServiceAccountName = "noSA"
		hasSA := appsv1.DaemonSet{}
		hasSA.Name = "hasSA"
		hasSA.Spec.Template.Spec.ServiceAccountName = "hasSA"
		list := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{hasSA, noSA}}

		saMap := make(map[string]bool)
		saMap[hasSA.Spec.Template.Spec.ServiceAccountName] = true
		saMap["random"] = true

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().PatchDaemonSet(gomock.Any(), "noSA", gomock.Any()).Return(nil, nil).Times(1)
		mock.EXPECT().PatchDaemonSet(gomock.Any(), "hasSA", gomock.Any()).Return(nil, nil).Times(1)

		err := deleteDaemonSetSATokensAndForceRollout(mock, "ns", deleteSATokens(saMap))
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(saMap["random"]).To(BeTrue())
		g.Expect(saMap["hasSA"]).To(BeFalse())
	})

	t.Run("Return error if delete SA fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		hasSA := appsv1.DaemonSet{}
		hasSA.Spec.Template.Spec.ServiceAccountName = "hasSA"
		list := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{hasSA}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)

		err := deleteDaemonSetSATokensAndForceRollout(mock, "ns", func(name string) error {
			return errors.New("Delete SA failed")
		})
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("Return error if patch daemonset fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		hasSA := appsv1.DaemonSet{}
		hasSA.Spec.Template.Spec.ServiceAccountName = "hasSA"
		list := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{hasSA}}

		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListDaemonSets(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().PatchDaemonSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errAPIGeneric).Times(1)

		err := deleteDaemonSetSATokensAndForceRollout(mock, "ns", func(_ string) error { return nil })
		g.Expect(err).To(HaveOccurred())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})
}

func TestDeleteSATokensFunc(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("List ServiceAccounts fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.ServiceAccountList{}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(gomock.Any(), gomock.Any()).Return(list, errAPIGeneric).Times(1)

		deleteSATokens, err := deleteSATokensFunc(mock, "ns")
		g.Expect(deleteSATokens).To(BeNil())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})

	t.Run("No ServiceAccounts in namespace", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		list := &v1.ServiceAccountList{}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)

		deleteSATokens, err := deleteSATokensFunc(mock, "ns")
		g.Expect(deleteSATokens).To(BeNil())
		g.Expect(err).To(BeNil())
	})

	t.Run("Expected ServiceAccount not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			v1.ObjectReference{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			v1.ObjectReference{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}
		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)

		deleteSATokens, err := deleteSATokensFunc(mock, "ns")
		g.Expect(deleteSATokens).ToNot(BeNil())
		g.Expect(err).To(BeNil())
		err = deleteSATokens("404")
		g.Expect(err).To(BeNil())
	})

	t.Run("Secret to delete not found, service account deleted", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			v1.ObjectReference{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			v1.ObjectReference{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}
		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().DeleteSecret(gomock.Any()).Return(errAPINotFound).Times(2)
		mock.EXPECT().DeleteServiceAccount(&sa).Return(nil).Times(1)

		deleteSATokens, err := deleteSATokensFunc(mock, "ns")
		g.Expect(deleteSATokens).ToNot(BeNil())
		g.Expect(err).To(BeNil())
		err = deleteSATokens(sa.Name)
		g.Expect(err).To(BeNil())
	})

	t.Run("Delete Secret fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			v1.ObjectReference{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			v1.ObjectReference{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}
		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().DeleteSecret(gomock.Any()).Return(errAPIGeneric).Times(1)

		deleteSATokens, err := deleteSATokensFunc(mock, "ns")
		g.Expect(deleteSATokens).ToNot(BeNil())
		g.Expect(err).To(BeNil())
		err = deleteSATokens(sa.Name)
		g.Expect(err).ToNot(BeNil())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})

	t.Run("Secrets deleted, service account deleted", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			v1.ObjectReference{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			v1.ObjectReference{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}

		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().DeleteSecret(gomock.Any()).Return(nil).Times(2)
		mock.EXPECT().DeleteServiceAccount(&sa).Return(nil).Times(1)

		deleteSATokens, err := deleteSATokensFunc(mock, "ns")
		g.Expect(deleteSATokens).ToNot(BeNil())
		g.Expect(err).To(BeNil())
		err = deleteSATokens(sa.Name)
		g.Expect(err).To(BeNil())
		// check is only deleted once
		err = deleteSATokens(sa.Name)
		g.Expect(err).To(BeNil())
	})

	t.Run("Secrets deleted, delete service account fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		sa := v1.ServiceAccount{}
		sa.Name = "sa"
		sa.Secrets = []v1.ObjectReference{
			v1.ObjectReference{
				Name:      "sasecret1",
				Namespace: "ns",
			},
			v1.ObjectReference{
				Name:      "sasecret2",
				Namespace: "ns",
			},
		}

		list := &v1.ServiceAccountList{Items: []v1.ServiceAccount{sa}}
		mock := mock.NewMockKubeClient(mockCtrl)
		mock.EXPECT().ListServiceAccounts(gomock.Any(), gomock.Any()).Return(list, nil).Times(1)
		mock.EXPECT().DeleteSecret(gomock.Any()).Return(nil).Times(2)
		mock.EXPECT().DeleteServiceAccount(&sa).Return(errAPIGeneric).Times(1)

		deleteSATokens, err := deleteSATokensFunc(mock, "ns")
		g.Expect(deleteSATokens).ToNot(BeNil())
		g.Expect(err).To(BeNil())
		err = deleteSATokens(sa.Name)
		g.Expect(err).ToNot(BeNil())
		g.Expect(errors.Cause(err)).To(Equal(errAPIGeneric))
	})
}
