package storageclass

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

// StorageClass is used to parse data from kubectl get storageclass
type StorageClass struct {
	Metadata   Metadata   `json:"metadata"`
	Parameters Parameters `json:"parameters"`
}

// Metadata holds information like name, create time
type Metadata struct {
	CreatedAt time.Time `json:"creationTimestamp"`
	Name      string    `json:"name"`
}

// Parameters holds information like skuName
type Parameters struct {
	SkuName string `json:"skuName"`
}

// CreateStorageClassFromFile will create a StorageClass from file with a name
func CreateStorageClassFromFile(filename, name string) (*StorageClass, error) {
	cmd := exec.Command("kubectl", "apply", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create StorageClass %s:%s\n", name, string(out))
		return nil, err
	}
	sc, err := Get(name)
	if err != nil {
		log.Printf("Error while trying to fetch StorageClass %s:%s\n", name, err)
		return nil, err
	}
	return sc, nil
}

// Get will return a StorageClass with a given name and namespace
func Get(scName string) (*StorageClass, error) {
	cmd := exec.Command("kubectl", "get", "storageclass", scName, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	sc := StorageClass{}
	err = json.Unmarshal(out, &sc)
	if err != nil {
		log.Printf("Error unmarshalling StorageClass json:%s\n", err)
		return nil, err
	}
	return &sc, nil
}

// WaitOnReady will block until StorageClass is available
func (sc *StorageClass) WaitOnReady(sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for StorageClass (%s) to become ready", duration.String(), sc.Metadata.Name)
			default:
				query, _ := Get(sc.Metadata.Name)
				if query != nil {
					readyCh <- true
				} else {
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
