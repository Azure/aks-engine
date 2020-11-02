// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package ssh

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

// ExecuteRemote executes a script in a remote host
func ExecuteRemote(host *RemoteHost, script string) (combinedOutput string, err error) {
	c, err := client(host)
	if err != nil {
		return "", errors.Wrap(err, "creating SSH client")
	}
	defer c.Close()
	s, err := c.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer s.Close()
	if co, err := s.CombinedOutput(script); err != nil {
		return string(co), errors.Wrapf(err, "executing script")
	}
	return "", nil
}

// PublicKeyAuth returns an AuthMethod that uses a ssh key pair
func PublicKeyAuth(sshPrivateKeyPath string) (ssh.AuthMethod, error) {
	b, err := ioutil.ReadFile(sshPrivateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "reading ssh private key file")
	}
	k, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return nil, errors.Wrap(err, "parsing ssh private key content")
	}
	return ssh.PublicKeys(k), nil
}

func client(host *RemoteHost) (*ssh.Client, error) {
	jbConfig, err := config(host.Jumpbox.AuthConfig)
	if err != nil {
		return nil, errors.Wrap(err, "creating jumpbox client config")
	}
	jbConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host.Jumpbox.URI, host.Jumpbox.Port), jbConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "dialing jumpbox (%s)", host.Jumpbox.URI)
	}
	hostConn, err := jbConn.Dial("tcp", fmt.Sprintf("%s:%d", host.URI, host.Port))
	if err != nil {
		return nil, errors.Wrapf(err, "dialing host (%s)", host.URI)
	}
	hostConfig, err := config(host.AuthConfig)
	if err != nil {
		return nil, errors.Wrap(err, "creating host client config")
	}
	ncc, chans, reqs, err := ssh.NewClientConn(hostConn, host.URI, hostConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "starting new client connection to host (%s)", host.URI)
	}
	return ssh.NewClient(ncc, chans, reqs), nil
}

func config(authConfig *AuthConfig) (*ssh.ClientConfig, error) {
	var err error
	var auth ssh.AuthMethod
	if authConfig.PrivateKeyPath != "" {
		auth, err = PublicKeyAuth(authConfig.PrivateKeyPath)
		if err != nil {
			return nil, err
		}
	} else {
		auth = ssh.Password(authConfig.Password)
	}
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            authConfig.User,
		Auth:            []ssh.AuthMethod{auth},
	}, nil
}
