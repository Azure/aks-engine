// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// RemoteRun executes remote command
func RemoteRun(user string, addr string, port int, sshKey []byte, cmd string) (string, error) {
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(sshKey)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	// Authentication
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error { return nil },
	}
	// Connect
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", addr, port), config)
	if err != nil {
		return "", err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b // get output

	err = session.Run(cmd)
	return b.String(), err
}
