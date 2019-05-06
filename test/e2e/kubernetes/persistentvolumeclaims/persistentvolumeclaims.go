// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package persistentvolumeclaims

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const commandTimeout = 1 * time.Minute

type List struct {
	PersistentVolumeClaims []PersistentVolumeClaim `json:"items"`
}

// PersistentVolumeClaim is used to parse data from kubectl get pvc
type PersistentVolumeClaim struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
}

// Metadata holds information like name, create time, and namespace
type Metadata struct {
	CreatedAt time.Time `json:"creationTimestamp"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
}

// Spec holds information like storageClassName, volumeName
type Spec struct {
	StorageClassName string `json:"storageClassName"`
	VolumeName       string `json:"volumeName"`
}

// Status holds information like phase
type Status struct {
	Phase string `json:"phase"`
}

// CreatePersistentVolumeClaimsFromFile will create a PVC from file with a name
func CreatePersistentVolumeClaimsFromFile(filename, name, namespace string) (*PersistentVolumeClaim, error) {
	cmd := exec.Command("k", "apply", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create PersistentVolumeClaim %s in namespace %s:%s\n", name, namespace, string(out))
		return nil, err
	}
	pvc, err := Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch PersistentVolumeClaim %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return pvc, nil
}

// CreatePVCFromFileDeleteIfExist will create a PVC from file with a name
func CreatePVCFromFileDeleteIfExist(filename, name, namespace string) (*PersistentVolumeClaim, error) {
	pvc, _ := Get(name, namespace)
	if pvc != nil {
		err := pvc.Delete(util.DefaultDeleteRetries)
		if err != nil {
			return nil, err
		}
		_, err = WaitOnDeleted(pvc.Metadata.Name, pvc.Metadata.Namespace, 5*time.Second, 1*time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return CreatePersistentVolumeClaimsFromFile(filename, name, namespace)
}

// Get will return a PersistentVolumeClaim with a given name and namespace
func Get(pvcName, namespace string) (*PersistentVolumeClaim, error) {
	cmd := exec.Command("k", "get", "pvc", pvcName, "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	pvc := PersistentVolumeClaim{}
	err = json.Unmarshal(out, &pvc)
	if err != nil {
		log.Printf("Error unmarshalling PersistentVolumeClaim json:%s\n", err)
		return nil, err
	}
	return &pvc, nil
}

// GetAll will return all pvcs in a given namespace
func GetAll(namespace string) (*List, error) {
	cmd := exec.Command("k", "get", "pvc", "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	pvcl := List{}
	err = json.Unmarshal(out, &pvcl)
	if err != nil {
		log.Printf("Error unmarshalling pvc json:%s\n", err)
		return nil, err
	}
	return &pvcl, nil
}

// GetAllByPrefix will return all jobs in a given namespace that match a prefix
func GetAllByPrefix(prefix, namespace string) ([]PersistentVolumeClaim, error) {
	pvcl, err := GetAll(namespace)
	if err != nil {
		return nil, err
	}
	pvcs := []PersistentVolumeClaim{}
	for _, p := range pvcl.PersistentVolumeClaims {
		matched, err := regexp.MatchString(prefix+"-.*", p.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return nil, err
		}
		if matched {
			pvcs = append(pvcs, p)
		}
	}
	return pvcs, nil
}

// Describe gets the description for the given pvc and namespace.
func Describe(pvcName, namespace string) error {
	cmd := exec.Command("k", "describe", "pvc", pvcName, "-n", namespace)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("\n%s\n", string(out))
	return nil
}

// Delete will delete a PersistentVolumeClaim in a given namespace
func (pvc *PersistentVolumeClaim) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "delete", "pvc", "-n", pvc.Metadata.Namespace, pvc.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd, commandTimeout)
		if kubectlError != nil {
			log.Printf("Error while trying to delete PVC %s in namespace %s:%s\n", pvc.Metadata.Name, pvc.Metadata.Namespace, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// WaitOnReady will block until PersistentVolumeClaim is available
func (pvc *PersistentVolumeClaim) WaitOnReady(namespace string, sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for PersistentVolumeClaim (%s) to become ready", duration.String(), pvc.Metadata.Name)
			default:
				query, _ := Get(pvc.Metadata.Name, namespace)
				if query != nil && query.Status.Phase == "Bound" {
					readyCh <- true
				} else {
					Describe(pvc.Metadata.Name, namespace)
					time.Sleep(sleep)
				}
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// WaitOnDeleted returns when a pvc is successfully deleted
func WaitOnDeleted(pvcPrefix, namespace string, sleep, duration time.Duration) (bool, error) {
	succeededCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Jobs (%s) to be deleted in namespace (%s)", duration.String(), pvcPrefix, namespace)
			default:
				p, err := GetAllByPrefix(pvcPrefix, namespace)
				if err != nil {
					errCh <- errors.Errorf("Got error while getting Jobs with prefix \"%s\" in namespace \"%s\"", pvcPrefix, namespace)
				}
				if len(p) == 0 {
					succeededCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case deleted := <-succeededCh:
			return deleted, nil
		}
	}
}
