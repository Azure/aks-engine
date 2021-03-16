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
		return nil, errors.Wrapf(err, "getting %s deployment", name)
	}
	if apierrors.IsNotFound(err) || *deploy.Spec.Replicas == 0 {
		return func() error { return nil }, nil
	}

	patch := func(msg string, count int32) error {
		log.Infof(msg)
		json := fmt.Sprintf(`{"spec":{"replicas": %d}}`, count)
		if _, err = client.PatchDeployment(metav1.NamespaceSystem, name, json); err != nil {
			return errors.Wrapf(err, "applying patch to %s deployment", name)
		}
		return nil
	}

	if err := patch(fmt.Sprintf("Pausing %s, setting replica count to 0", name), 0); err != nil {
		return nil, err
	}

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

// RotateServiceAccountTokens deletes service account tokens referenced by daemonsets and deployments
// from the namespaces of interest and triggers a rollout once the tokens are deleted.
//
// Service account tokens are signed by the cluster CA,
// deleting them after the CA is rotated ensures that KCM will regenerate tokens signed by the new CA.
func RotateServiceAccountTokens(client internal.KubeClient, namespaces []string) error {
	for _, ns := range namespaces {
		deleteSATokens, err := deleteSATokensFunc(client, ns)
		if err != nil {
			return err
		}
		if deleteSATokens == nil {
			// no tokens to rotate in this namespace
			continue
		}
		if err = deleteDeploymentSATokensAndForceRollout(client, ns, deleteSATokens); err != nil {
			return err
		}
		if err = deleteDaemonSetSATokensAndForceRollout(client, ns, deleteSATokens); err != nil {
			return err
		}
	}
	return nil
}

func deleteDeploymentSATokensAndForceRollout(client internal.KubeClient, ns string, deleteSATokens func(string) error) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"ca-rotation":"%d"}}}}}`, random.Int31())

	deployList, err := client.ListDeployments(ns, metav1.ListOptions{})
	if err != nil {
		return errors.Wrapf(err, "listing %s deployments", ns)
	}
	for _, deploy := range deployList.Items {
		if deploy.Spec.Template.Spec.ServiceAccountName != "" {
			// delete SA tokens
			if err = deleteSATokens(deploy.Spec.Template.Spec.ServiceAccountName); err != nil {
				return err
			}
		}
		// trigger rollout so the deploy replicas mount the newly generated sa token
		if _, err := client.PatchDeployment(ns, deploy.Name, patch); err != nil {
			return errors.Wrapf(err, "patching %s deployment %s", ns, deploy.Name)
		}
	}
	return nil
}

func deleteDaemonSetSATokensAndForceRollout(client internal.KubeClient, ns string, deleteSATokens func(string) error) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"ca-rotation":"%d"}}}}}`, random.Int31())

	dsList, err := client.ListDaemonSets(ns, metav1.ListOptions{})
	if err != nil {
		return errors.Wrapf(err, "listing %s daemonsets", ns)
	}
	for _, ds := range dsList.Items {
		if ds.Spec.Template.Spec.ServiceAccountName != "" {
			// delete SA tokens
			if err = deleteSATokens(ds.Spec.Template.Spec.ServiceAccountName); err != nil {
				return err
			}
		}
		// trigger rollout so the ds replicas mount the newly generated sa token
		if _, err = client.PatchDaemonSet(ns, ds.Name, patch); err != nil {
			return errors.Wrapf(err, "patching %s daemonset %s", ns, ds.Name)
		}
	}
	return nil
}

func deleteSATokensFunc(client internal.KubeClient, ns string) (func(string) error, error) {
	saList, err := client.ListServiceAccounts(ns, metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "listing %s service accounts", ns)
	}
	if len(saList.Items) == 0 {
		return nil, nil
	}
	saMap := make(map[string]v1.ServiceAccount)
	for _, sa := range saList.Items {
		saMap[sa.Name] = sa
	}
	return func(name string) error {
		sa, ok := saMap[name]
		if !ok {
			return nil
		}
		for _, s := range sa.Secrets {
			err := client.DeleteSecret(&v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: ns,
					Name:      s.Name,
				},
			})
			if err != nil && !apierrors.IsNotFound(err) {
				return errors.Wrapf(err, "deleting %s secret %s", ns, s.Name)
			}
		}
		if err := client.DeleteServiceAccount(&sa); err != nil && !apierrors.IsNotFound(err) {
			return errors.Wrapf(err, "deleting %s service account %s", ns, sa.Name)
		}
		delete(saMap, name)
		return nil
	}, nil
}
