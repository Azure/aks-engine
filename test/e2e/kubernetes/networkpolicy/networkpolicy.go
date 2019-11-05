//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package networkpolicy

import (
	"log"
	"os/exec"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
)

// CreateNetworkPolicyFromFile will create a NetworkPolicy from file with a name
func CreateNetworkPolicyFromFile(filename, name, namespace string) error {
	cmd := exec.Command("k", "create", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create NetworkPolicy %s:%s\n", name, string(out))
		return err
	}
	return nil
}

// DeleteNetworkPolicy will create a NetworkPolicy from file with a name
func DeleteNetworkPolicy(name, namespace string) error {
	cmd := exec.Command("k", "delete", "networkpolicy", "-n", namespace, name)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to delete NetworkPolicy %s in namespace %s:%s\n", name, namespace, string(out))
		return err
	}
	return nil
}
