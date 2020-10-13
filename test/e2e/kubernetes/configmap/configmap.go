//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package configmap

import (
	"encoding/json"
	"log"
	"os/exec"
	"time"
)

// ConfigMap is used to parse data from kubectl get configmap
type ConfigMap struct {
	Data map[string]string `json:"data"`
}

// Get will return a configmap with a given name and namespace
func Get(cmName, namespace string, retries int) (*ConfigMap, error) {
	cm := ConfigMap{}
	var out []byte
	var err error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "get", "configmap", cmName, "-n", namespace, "-o", "json")
		out, err = cmd.CombinedOutput()
		if err == nil {
			jsonErr := json.Unmarshal(out, &cm)
			if jsonErr != nil {
				log.Printf("Error unmarshalling configmap json:%s\n", jsonErr)
				err = jsonErr
			}
			if err == nil {
				return &cm, err
			}
		}
		time.Sleep(3 * time.Second)
	}
	return &cm, err
}
