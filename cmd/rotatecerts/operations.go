// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package rotatecerts

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RotateServiceAccountTokens deletes service account tokens referenced by kube-system daemonsets/deployments
// and triggers a rollout of daemonsets/deployments that mount service account tokens.
//
// Service account tokens are signed by the cluster CA, deleting them after the CA is rotated ensure that
// KCM will regenerate tokens signed by the new CA.
func RotateServiceAccountTokens(client Client, namespaces []string) error {
	for _, ns := range namespaces {
		err := rotateServiceAccountTokensInNamespace(client, ns)
		if err != nil {
			return err
		}
	}
	return nil
}

func rotateServiceAccountTokensInNamespace(client Client, ns string) error {
	saList, err := client.ListServiceAccounts(ns, metav1.ListOptions{})
	if err != nil {
		return errors.Errorf("listing %s service accounts", ns)
	}
	saMap := make(map[string]v1.ServiceAccount)
	for _, sa := range saList.Items {
		saMap[sa.Name] = sa
	}
	if len(saMap) == 0 {
		return nil
	}

	deployList, err := client.ListDeployments(ns, metav1.ListOptions{})
	if err != nil {
		return errors.Errorf("listing %s deployments", ns)
	}
	err = deleteDeploymentSATokensAndForceRollout(client, ns, deployList, saMap)
	if err != nil {
		return err
	}

	dsList, err := client.ListDaemonSets(ns, metav1.ListOptions{})
	if err != nil {
		return errors.Errorf("listing %s daemonsets", ns)
	}
	err = deleteDaemonSetSATokensAndForceRollout(client, ns, dsList, saMap)
	if err != nil {
		return err
	}
	return nil
}

func deleteDeploymentSATokensAndForceRollout(client Client, ns string, deployList *appsv1.DeploymentList, saMap map[string]v1.ServiceAccount) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"ca-rotation":"%d"}}}}}`, random.Int31())

	for _, deploy := range deployList.Items {
		if deploy.Spec.Template.Spec.ServiceAccountName == "" {
			continue
		}
		// delete SA tokens
		if sa, ok := saMap[deploy.Spec.Template.Spec.ServiceAccountName]; ok {
			for _, s := range sa.Secrets {
				err := client.DeleteSecret(&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: ns,
						Name:      s.Name,
					},
				})
				if err != nil && !apierrors.IsNotFound(err) {
					return errors.Errorf("deleting %s secret %s", ns, s.Name)
				}
			}
		}
		// trigger rollout so the deploy replicas mount the newly generated sa token
		_, err := client.PatchDeployment(ns, deploy.Name, patch)
		if err != nil {
			return errors.Errorf("patching %s deployment %s", ns, deploy.Name)
		}
	}
	return nil
}

func deleteDaemonSetSATokensAndForceRollout(client Client, ns string, dsList *appsv1.DaemonSetList, saMap map[string]v1.ServiceAccount) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"ca-rotation":"%d"}}}}}`, random.Int31())

	for _, ds := range dsList.Items {
		if ds.Spec.Template.Spec.ServiceAccountName == "" {
			continue
		}
		// delete SA tokens
		if sa, ok := saMap[ds.Spec.Template.Spec.ServiceAccountName]; ok {
			for _, s := range sa.Secrets {
				err := client.DeleteSecret(&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: ns,
						Name:      s.Name,
					},
				})
				if err != nil && !apierrors.IsNotFound(err) {
					return errors.Errorf("deleting %s secret %s", ns, s.Name)
				}
			}
		}
		// trigger rollout so the ds replicas mount the newly generated sa token
		_, err := client.PatchDaemonSet(ns, ds.Name, patch)
		if err != nil {
			return errors.Errorf("patching %s daemonset %s", ns, ds.Name)
		}
	}
	return nil
}
