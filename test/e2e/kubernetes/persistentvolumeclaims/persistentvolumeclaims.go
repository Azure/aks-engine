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

// GetResult is a return struct for GetAsync
type GetResult struct {
	pvc *PersistentVolumeClaim
	err error
}

// GetAsync wraps Get with a struct response for goroutine + channel usage
func GetAsync(pvcName, namespace string) GetResult {
	pvc, err := Get(pvcName, namespace)
	return GetResult{
		pvc: pvc,
		err: err,
	}
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

// GetAllByPrefixResult is a return struct for GetAllByPrefixAsync
type GetAllByPrefixResult struct {
	pvcs []PersistentVolumeClaim
	err  error
}

// GetAllByPrefixAsync wraps Get with a struct response for goroutine + channel usage
func GetAllByPrefixAsync(prefix, namespace string) GetAllByPrefixResult {
	pvcs, err := GetAllByPrefix(prefix, namespace)
	return GetAllByPrefixResult{
		pvcs: pvcs,
		err:  err,
	}
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
func (pvc *PersistentVolumeClaim) WaitOnReady(namespace string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var mostRecentWaitOnReadyError error
	ch := make(chan GetResult)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- GetAsync(pvc.Metadata.Name, namespace):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitOnReadyError = result.err
			pvc := result.pvc
			if mostRecentWaitOnReadyError == nil {
				if pvc != nil && pvc.Status.Phase == "Bound" {
					return true, nil
				}
				Describe(pvc.Metadata.Name, namespace)
			}
		case <-ctx.Done():
			return false, errors.Errorf("WaitOnReady timed out: %s\n", mostRecentWaitOnReadyError)
		}
	}
}

// WaitOnDeleted returns when a pvc is successfully deleted
func WaitOnDeleted(pvcPrefix, namespace string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetAllByPrefixResult)
	var mostRecentWaitOnDeletedError error
	var pvcs []PersistentVolumeClaim
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- GetAllByPrefixAsync(pvcPrefix, namespace):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitOnDeletedError = result.err
			pvcs = result.pvcs
			if mostRecentWaitOnDeletedError == nil {
				if len(pvcs) == 0 {
					return true, nil
				}
			}
		case <-ctx.Done():
			return false, errors.Errorf("WaitOnDeleted timed out: %s\n", mostRecentWaitOnDeletedError)
		}
	}
}
