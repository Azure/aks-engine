//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package util

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/pkg/errors"
)

// DefaultDeleteRetries defines a default retry count for resource deletion operations
const DefaultDeleteRetries = 10

// PrintCommand prints a command string
func PrintCommand(cmd *exec.Cmd) {
	fmt.Printf("\n$ %s\n", strings.Join(cmd.Args, " "))
}

// RunAndLogCommand logs the command with a timestamp when it's run, and the duration at end
func RunAndLogCommand(cmd *exec.Cmd, timeout time.Duration) ([]byte, error) {
	var zeroValueDuration time.Duration
	var err error
	var out []byte
	cmdLine := fmt.Sprintf("$ %s", strings.Join(cmd.Args, " "))
	start := time.Now()
	log.Printf("%s", cmdLine)
	out, err = cmd.CombinedOutput()
	end := time.Now()
	total := time.Since(start)
	log.Printf("#### %s completed in %s", cmdLine, end.Sub(start).String())
	if zeroValueDuration != timeout {
		if total.Seconds() > timeout.Seconds() {
			err = errors.Errorf("%s took too long!", cmdLine)
		}
	}
	return out, err
}

// AddToSSHKeyChain is a helper func to setup ssh agent forwarding
func AddToSSHKeyChain(keyfile string) error {
	cmd := exec.Command("ssh-add", "-D")
	PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	log.Printf("%s\n", out)
	if err != nil {
		return fmt.Errorf("Error while cleaning ssh agent keychain: %s", err)
	}
	cmd = exec.Command("ssh-add", keyfile)
	PrintCommand(cmd)
	out, err = cmd.CombinedOutput()
	log.Printf("%s\n", out)
	if err != nil {
		return fmt.Errorf("Error while adding private key to ssh agent keychain for forwarding: %s", err)
	}
	return nil
}

// IsLargeVMSKU returns if the VM SKU is a known, > 8 core SKU
func IsLargeVMSKU(sku string) bool {
	switch sku {
	case "Standard_D16_v3":
		return true
	default:
		return false
	}
}

func IsUsingManagedDisks(agentPools []*api.AgentPoolProfile) bool {
	for _, a := range agentPools {
		if a.IsManagedDisks() {
			return true
		}
	}
	return false
}

func IsUsingEphemeralDisks(agentPools []*api.AgentPoolProfile) bool {
	for _, a := range agentPools {
		if a.IsEphemeral() {
			return true
		}
	}
	return false
}
