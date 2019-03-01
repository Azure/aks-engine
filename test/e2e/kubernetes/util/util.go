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
func RunAndLogCommand(cmd *exec.Cmd, timeout time.Duration) ([]byte, error) {
	var err error
	var out []byte
	cmdLine := fmt.Sprintf("$ %s", strings.Join(cmd.Args, " "))
	start := time.Now()
	log.Printf("%s", cmdLine)
	out, err = cmd.CombinedOutput()
	end := time.Now()
	total := time.Since(start)
	log.Printf("#### %s completed in %s", cmdLine, end.Sub(start).String())
	if total.Seconds() > timeout.Seconds() {
		err = fmt.Errorf(fmt.Sprintf("%s took too long!", cmdLine))
	}
	return out, err
}
