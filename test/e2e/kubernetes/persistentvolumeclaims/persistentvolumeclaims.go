package persistentvolumeclaims

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

// PersistentVolumeClaims is used to parse data from kubectl get pvc
type PersistentVolumeClaims struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
}

// Metadata holds information like name, create time, and namespace
type Metadata struct {
	CreatedAt time.Time `json:"creationTimestamp"`
	Name      string    `json:"name"`
	NameSpace string    `json:"namespace"`
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

// CreatePersistentVolumeClaimsFromFile will create a StorageClass from file with a name
func CreatePersistentVolumeClaimsFromFile(filename, name, namespace string) (*PersistentVolumeClaims, error) {
	cmd := exec.Command("kubectl", "apply", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create PersistentVolumeClaims %s in namespace %s:%s\n", name, namespace, string(out))
		return nil, err
	}
	pvc, err := Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch PersistentVolumeClaims %s in namespace %s:%s\n", name, namespace, err)
		return nil, err
	}
	return pvc, nil
}

// Get will return a PersistentVolumeClaims with a given name and namespace
func Get(pvcName, namespace string) (*PersistentVolumeClaims, error) {
	cmd := exec.Command("kubectl", "get", "pvc", pvcName, "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	pvc := PersistentVolumeClaims{}
	err = json.Unmarshal(out, &pvc)
	if err != nil {
		log.Printf("Error unmarshalling PersistentVolumeClaims json:%s\n", err)
		return nil, err
	}
	return &pvc, nil
}

// Describe gets the description for the given pvc and namespace.
func Describe(pvcName, namespace string) error {
	cmd := exec.Command("kubectl", "describe", "pvc", pvcName, "-n", namespace)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("\n%s\n", string(out))
	return nil
}

// Delete will delete a PersistentVolumeClaims in a given namespace
func (pvc *PersistentVolumeClaims) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("kubectl", "delete", "pvc", "-n", pvc.Metadata.NameSpace, pvc.Metadata.Name)
		kubectlOutput, kubectlError = util.RunAndLogCommand(cmd)
		if kubectlError != nil {
			log.Printf("Error while trying to delete PVC %s in namespace %s:%s\n", pvc.Metadata.Name, pvc.Metadata.NameSpace, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// WaitOnReady will block until PersistentVolumeClaims is available
func (pvc *PersistentVolumeClaims) WaitOnReady(namespace string, sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for PersistentVolumeClaims (%s) to become ready", duration.String(), pvc.Metadata.Name)
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
