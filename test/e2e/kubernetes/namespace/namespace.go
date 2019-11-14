//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package namespace

import (
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
)

// Namespace holds namespace metadata
type Namespace struct {
	Metadata Metadata `json:"metadata"`
}

// Metadata holds information like name and created timestamp
type Metadata struct {
	CreatedAt time.Time `json:"creationTimestamp"`
	Name      string    `json:"name"`
}

// Create a namespace with the given name
func Create(name string) (*Namespace, error) {
	cmd := exec.Command("k", "create", "namespace", name)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create namespace (%s):%s\n", name, string(out))
		return nil, err
	}
	return Get(name)
}

// CreateIfNotExist a namespace with the given name if it doesn't exist already
func CreateIfNotExist(name string) (*Namespace, error) {
	n, err := Get(name)
	if err != nil {
		return Create(name)
	}
	return n, nil
}

// Get returns a namespace for with a given name
func Get(name string) (*Namespace, error) {
	cmd := exec.Command("k", "get", "namespace", name, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to get namespace (%s):%s\n", name, string(out))
		return nil, err
	}
	n := Namespace{}
	err = json.Unmarshal(out, &n)
	if err != nil {
		log.Printf("Error unmarshalling namespace json:%s\n", err)
	}
	return &n, nil
}

// Delete a namespace
func (n *Namespace) Delete() error {
	cmd := exec.Command("k", "delete", "namespace", n.Metadata.Name)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to delete namespace (%s):%s\n", n.Metadata.Name, out)
		return err
	}
	return nil
}

// Label a namespace
func (n *Namespace) Label(label string) error {
	cmd := exec.Command("k", "label", "--overwrite=true", "namespace/"+n.Metadata.Name, label)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to label namespace (%s) with label %s, error: %s\n", n.Metadata.Name, label, out)
		return err
	}
	return nil
}
