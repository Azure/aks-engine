// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package util

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// PrintCommand prints a command string
func PrintCommand(cmd *exec.Cmd) {
	fmt.Printf("\n$ %s\n", strings.Join(cmd.Args, " "))
}

// RunAndLogCommand logs the command with a timestamp when it's run, and the duration at end
func RunAndLogCommand(cmd *exec.Cmd) ([]byte, error) {
	cmdLine := fmt.Sprintf("$ %s", strings.Join(cmd.Args, " "))
	start := time.Now()
	log.Printf("%s", cmdLine)
	out, err := cmd.CombinedOutput()
	end := time.Now()
	log.Printf("#### %s completed in %s", cmdLine, end.Sub(start).String())
	return out, err
}
