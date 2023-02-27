// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func SSHClient(jumpboxHost, jumpboxPort, hostname string, jumpboxConfig, nodeConfig *ssh.ClientConfig) (*ssh.Client, error) {
	lbConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", jumpboxHost, jumpboxPort), jumpboxConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "dialing jumpbox (%s)", jumpboxHost)
	}
	conn, err := lbConn.Dial("tcp", fmt.Sprintf("%s:22", hostname))
	if err != nil {
		return nil, errors.Wrapf(err, "dialing host (%s)", hostname)
	}
	ncc, chans, reqs, err := ssh.NewClientConn(conn, hostname, nodeConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "starting new client connection to host (%s)", hostname)
	}
	return ssh.NewClient(ncc, chans, reqs), nil
}

func SSHClientConfig(user string, auth ssh.AuthMethod) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		// FixedHostKey instead?
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            user,
		Auth: []ssh.AuthMethod{
			auth,
		},
	}
}

func PublicKeyAuth(sshPrivateKeyPath string) (ssh.AuthMethod, error) {
	b, err := os.ReadFile(sshPrivateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "reading ssh private key file")
	}
	k, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return nil, errors.Wrap(err, "parsing ssh private key content")
	}
	return ssh.PublicKeys(k), nil
}
