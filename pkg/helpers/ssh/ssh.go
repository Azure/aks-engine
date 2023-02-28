// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package ssh

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

// ExecuteRemote executes a script in a remote host.
//
// Context ctx is only enforced during the process that stablishes
// the SSH connection and creates the SSH client.
func ExecuteRemote(ctx context.Context, host *RemoteHost, script string) (combinedOutput string, err error) {
	c, err := clientWithRetry(ctx, host)
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

// ValidateConfig checks the JumpBox configuration
func ValidateConfig(host *JumpBox) error {
	jbConfig, err := config(host.AuthConfig)
	if err != nil {
		return errors.Wrap(err, "creating ssh client config")
	}
	_, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", host.URI, host.Port), jbConfig)
	if err != nil {
		return errors.Wrapf(err, "dialing ssh (%s)", host.URI)
	}
	return nil
}

func clientWithRetry(ctx context.Context, host *RemoteHost) (*ssh.Client, error) {
	// TODO Granular retry func
	retryFunc := func(err error) bool {
		select {
		case <-ctx.Done():
			return false
		default:
			return true
		}
	}
	backoff := wait.Backoff{Steps: 300, Duration: 10 * time.Second}
	var c *ssh.Client
	var err error
	err = retry.OnError(backoff, retryFunc, func() error {
		c, err = client(host)
		return err
	})
	return c, err
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
	c, err := ssh.NewClient(ncc, chans, reqs), nil
	if err != nil {
		return nil, errors.Wrapf(err, "creating new ssh client for host (%s)", host.URI)
	}
	return c, nil
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
