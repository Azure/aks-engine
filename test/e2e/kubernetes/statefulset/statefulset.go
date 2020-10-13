//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package statefulset

import (
	"encoding/json"
	"log"
	"os/exec"
	"time"
)

// Statefulset is used to parse data from kubectl get statefulsets
type Statefulset struct {
	Status Status `json:"status"`
}

// Status holds information like readyReplicas and replicas
type Status struct {
	ReadyReplicas int `json:"readyReplicas"`
	Replicas      int `json:"replicas"`
}

// Get will return a statefulset with a given name and namespace
func Get(ssName, namespace string, retries int) (*Statefulset, error) {
	ss := Statefulset{}
	var out []byte
	var err error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "get", "statefulsets", ssName, "-n", namespace, "-o", "json")
		out, err = cmd.CombinedOutput()
		if err == nil {
			jsonErr := json.Unmarshal(out, &ss)
			if jsonErr != nil {
				log.Printf("Error unmarshalling statefulsets json:%s\n", jsonErr)
				err = jsonErr
			}
			if err == nil {
				return &ss, err
			}
		}
		time.Sleep(3 * time.Second)
	}
	return &ss, err
}
