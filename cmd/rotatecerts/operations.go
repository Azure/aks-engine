// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Azure/aks-engine/cmd/rotatecerts/internal"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PauseClusterAutoscaler scales to zero the replica count of the cluster autoscaler deployment
// and returns a function that scales back to the original replica count.
//
// It NOPs if the original replica count is zero.
func PauseClusterAutoscaler(client internal.KubeClient) (func() error, error) {
	name := common.ClusterAutoscalerAddonName

	deploy, err := client.GetDeployment(metav1.NamespaceSystem, name)
	if err != nil && !apierrors.IsNotFound(err) {
		e := errors.Wrapf(err, "getting %s deployment", name)
		return func() error { return e }, e
	}
	if apierrors.IsNotFound(err) || *deploy.Spec.Replicas == 0 {
		// autoscaler not present or no replicas, NOP
		return func() error { return nil }, nil
	}

	// autoscaler present
	patch := func(msg string, count int32) error {
		log.Infof(msg)
		json := fmt.Sprintf(`{"spec":{"replicas": %d}}`, count)
		if _, err = client.PatchDeployment(metav1.NamespaceSystem, name, json); err != nil {
			return errors.Wrapf(err, "applying patch to %s deployment", name)
		}
		return nil
	}

	// pause autoscaler
	if err := patch(fmt.Sprintf("Pausing %s, setting replica count to 0", name), 0); err != nil {
		return func() error { return err }, err
	}

	// resume autoscaler func
	return func() error {
		c := *deploy.Spec.Replicas
		err := patch(fmt.Sprintf("Resuming %s, setting replica count to %d", name, c), c)
		log.Warnln("Run \"aks-engine upgrade\" to refresh the cluster-autoscaler node template")
		if err != nil {
			return err
		}
		return nil
	}, nil
}

// RotateServiceAccountTokens deletes all service account tokens and
// triggers a forced rollout of all daemonsets and deployments.
//
// Service account tokens are signed by the cluster CA,
// deleting them after the CA is rotated ensures that KCM will regenerate tokens signed by the new CA.
func RotateServiceAccountTokens(client internal.KubeClient) error {
	if err := deleteSATokens(client); err != nil {
		return err
	}
	if err := rolloutDeployments(client); err != nil {
		return err
	}
	if err := rolloutDaemonSets(client); err != nil {
		return err
	}
	// TODO rolloutStatefulSets
	return nil
}

func rolloutDeployments(client internal.KubeClient) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"ca-rotation":"%d"}}}}}`, random.Int31())

	deployList, err := client.ListDeployments(metav1.NamespaceAll, metav1.ListOptions{})
	if err != nil {
		return errors.Wrapf(err, "listing cluster deployments")
	}
	for _, deploy := range deployList.Items {
		// trigger rollout so the deploy replicas mount the newly generated sa token
		if _, err := client.PatchDeployment(deploy.Namespace, deploy.Name, patch); err != nil {
			return errors.Wrapf(err, "patching %s deployment %s", deploy.Namespace, deploy.Name)
		}
	}
	return nil
}

func rolloutDaemonSets(client internal.KubeClient) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"ca-rotation":"%d"}}}}}`, random.Int31())

	dsList, err := client.ListDaemonSets(metav1.NamespaceAll, metav1.ListOptions{})
	if err != nil {
		return errors.Wrapf(err, "listing cluster daemonsets")
	}
	for _, ds := range dsList.Items {
		// trigger rollout so the ds replicas mount the newly generated sa token
		if _, err = client.PatchDaemonSet(ds.Namespace, ds.Name, patch); err != nil {
			return errors.Wrapf(err, "patching %s daemonset %s", ds.Namespace, ds.Name)
		}
	}
	return nil
}

func deleteSATokens(client internal.KubeClient) error {
	saList, err := client.ListServiceAccounts(metav1.NamespaceAll, metav1.ListOptions{})
	if err != nil {
		return errors.Wrapf(err, "listing cluster service accounts")
	}
	if len(saList.Items) == 0 {
		return nil
	}
	for _, sa := range saList.Items {
		for _, s := range sa.Secrets {
			err := client.DeleteSecret(&v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: sa.Namespace,
					Name:      s.Name,
				},
			})
			if err != nil && !apierrors.IsNotFound(err) {
				return errors.Wrapf(err, "deleting %s secret %s", s.Namespace, s.Name)
			}
		}
	}
	return nil
}
