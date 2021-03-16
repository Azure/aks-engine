// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package kubernetes

import (
	"crypto/x509"
	"errors"
	"net/url"
	"testing"
	"time"

	mock "github.com/Azure/aks-engine/pkg/kubernetes/internal/mock_internal"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errAPIGeneric         = errors.New("generic api error")
	errAPINotFound        = &apierrors.StatusError{ErrStatus: metav1.Status{Reason: metav1.StatusReasonNotFound}}
	unknownAuthorityError = &url.Error{Err: x509.UnknownAuthorityError{}}
)

func TestListPods(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, opts := "ns", metav1.ListOptions{}
	result := &v1.PodList{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().ListPodsByOptions(ns, opts).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().ListPodsByOptions(ns, opts).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListPodsByOptions(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListPods(ns, opts)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().ListPodsByOptions(ns, opts).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListPodsByOptions(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListPods(ns, opts)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestDeletePods(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, opts := "ns", metav1.ListOptions{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().DeletePods(ns, opts).Return(errAPIGeneric),
			oldCAClientMock.EXPECT().DeletePods(ns, opts).Return(nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeletePods(ns, opts).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeletePods(ns, opts)
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("good client success after a 404, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().DeletePods(ns, opts).Return(errAPINotFound).MaxTimes(1)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeletePods(ns, opts).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeletePods(ns, opts)
		g.Expect(err).To(HaveOccurred())
		g.Expect(apierrors.IsNotFound(err)).To(BeTrue())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().DeletePods(ns, opts).Return(errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeletePods(ns, opts).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeletePods(ns, opts)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestListNodes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	result := &v1.NodeList{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().ListNodes().Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().ListNodes().Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListNodes().Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListNodes()
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().ListNodes().Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListNodes().Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListNodes()
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestListServiceAccounts(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, opts := "ns", metav1.ListOptions{}
	result := &v1.ServiceAccountList{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().ListServiceAccountsByOptions(ns, opts).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().ListServiceAccountsByOptions(ns, opts).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListServiceAccountsByOptions(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListServiceAccounts(ns, opts)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().ListServiceAccountsByOptions(ns, opts).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListServiceAccountsByOptions(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListServiceAccounts(ns, opts)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestListDeployments(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, opts := "ns", metav1.ListOptions{}
	result := &appsv1.DeploymentList{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().ListDeployments(ns, opts).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().ListDeployments(ns, opts).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListDeployments(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListDeployments(ns, opts)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().ListDeployments(ns, opts).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListDeployments(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListDeployments(ns, opts)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestGetDeployment(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, name := "ns", "name"
	result := &appsv1.Deployment{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().GetDeployment(ns, name).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().GetDeployment(ns, name).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().GetDeployment(ns, name).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.GetDeployment(ns, name)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("good client success after a 404, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().GetDeployment(ns, name).Return(nil, errAPINotFound).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().GetDeployment(ns, name).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.GetDeployment(ns, name)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPINotFound))
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().GetDeployment(ns, name).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().GetDeployment(ns, name).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.GetDeployment(ns, name)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestPatchDeployment(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, name, json := "ns", "name", "patch"
	result := &appsv1.Deployment{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().PatchDeployment(ns, name, json).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().PatchDeployment(ns, name, json).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().PatchDeployment(ns, name, json).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.PatchDeployment(ns, name, json)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("good client success after a 404, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().PatchDeployment(ns, name, json).Return(nil, errAPINotFound).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().PatchDeployment(ns, name, json).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.PatchDeployment(ns, name, json)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPINotFound))
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().PatchDeployment(ns, name, json).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().PatchDeployment(ns, name, json).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.PatchDeployment(ns, name, json)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestListDaemonSets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, opts := "ns", metav1.ListOptions{}
	result := &appsv1.DaemonSetList{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().ListDaemonSets(ns, opts).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().ListDaemonSets(ns, opts).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListDaemonSets(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListDaemonSets(ns, opts)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().ListDaemonSets(ns, opts).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListDaemonSets(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListDaemonSets(ns, opts)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestPatchDaemonSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, name, json := "ns", "name", "patch"
	result := &appsv1.DaemonSet{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().PatchDaemonSet(ns, name, json).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().PatchDaemonSet(ns, name, json).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().PatchDaemonSet(ns, name, json).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.PatchDaemonSet(ns, name, json)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("good client success after a 404, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().PatchDaemonSet(ns, name, json).Return(nil, errAPINotFound).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().PatchDaemonSet(ns, name, json).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.PatchDaemonSet(ns, name, json)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPINotFound))
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().PatchDaemonSet(ns, name, json).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().PatchDaemonSet(ns, name, json).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.PatchDaemonSet(ns, name, json)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestListSecrets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ns, opts := "ns", metav1.ListOptions{}
	result := &v1.SecretList{}

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().ListSecrets(ns, opts).Return(nil, errAPIGeneric),
			oldCAClientMock.EXPECT().ListSecrets(ns, opts).Return(result, nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListSecrets(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListSecrets(ns, opts)
		g.Expect(x).To(Equal(result))
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().ListSecrets(ns, opts).Return(nil, errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().ListSecrets(ns, opts).Return(nil, unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		x, err := sut.ListSecrets(ns, opts)
		g.Expect(x).To(BeNil())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestDeleteServiceAccount(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().DeleteServiceAccount(gomock.Any()).Return(errAPIGeneric),
			oldCAClientMock.EXPECT().DeleteServiceAccount(gomock.Any()).Return(nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeleteServiceAccount(gomock.Any()).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeleteServiceAccount(&v1.ServiceAccount{})
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("good client success after a 404, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().DeleteServiceAccount(gomock.Any()).Return(errAPINotFound).MaxTimes(1)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeleteServiceAccount(gomock.Any()).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeleteServiceAccount(&v1.ServiceAccount{})
		g.Expect(err).To(HaveOccurred())
		g.Expect(apierrors.IsNotFound(err)).To(BeTrue())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().DeleteServiceAccount(gomock.Any()).Return(errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeleteServiceAccount(gomock.Any()).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeleteServiceAccount(&v1.ServiceAccount{})
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}

func TestDeleteSecret(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	t.Run("good client success after a retry, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		gomock.InOrder(
			oldCAClientMock.EXPECT().DeleteSecret(gomock.Any()).Return(errAPIGeneric),
			oldCAClientMock.EXPECT().DeleteSecret(gomock.Any()).Return(nil),
		)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeleteSecret(gomock.Any()).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeleteSecret(&v1.Secret{})
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("good client success after a 404, bad client tries only once", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().DeleteSecret(gomock.Any()).Return(errAPINotFound).MaxTimes(1)
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeleteSecret(gomock.Any()).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeleteSecret(&v1.Secret{})
		g.Expect(err).To(HaveOccurred())
		g.Expect(apierrors.IsNotFound(err)).To(BeTrue())
	})

	t.Run("both clients fail, process times out", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		oldCAClientMock := mock.NewMockClient(mockCtrl)
		oldCAClientMock.EXPECT().DeleteSecret(gomock.Any()).Return(errAPIGeneric).AnyTimes()
		newCAClientMock := mock.NewMockClient(mockCtrl)
		newCAClientMock.EXPECT().DeleteSecret(gomock.Any()).Return(unknownAuthorityError).MaxTimes(1)

		interval, timeout := 1*time.Second, 5*time.Second
		sut := NewCompositeClient(oldCAClientMock, newCAClientMock, interval, timeout)
		err := sut.DeleteSecret(&v1.Secret{})
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(Equal(errAPIGeneric))
	})
}
