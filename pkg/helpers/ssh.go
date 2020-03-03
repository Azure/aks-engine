package helpers

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func SSHClient(jumpboxHost, jumpboxPort, hostname string, config *ssh.ClientConfig) (*ssh.Client, error) {
	lbConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", jumpboxHost, jumpboxPort), config)
	if err != nil {
		return nil, errors.Wrapf(err, "dialing load balancer (%s)", jumpboxHost)
	}
	conn, err := lbConn.Dial("tcp", fmt.Sprintf("%s:22", hostname))
	if err != nil {
		return nil, errors.Wrapf(err, "dialing host (%s)", hostname)
	}
	ncc, chans, reqs, err := ssh.NewClientConn(conn, hostname, config)
	if err != nil {
		return nil, errors.Wrapf(err, "starting new client connection to host (%s)", hostname)
	}
	return ssh.NewClient(ncc, chans, reqs), nil
}

func SSHClientConfig(user, sshPrivateKeyPath string) (*ssh.ClientConfig, error) {
	auth, err := publicKeyAuth(sshPrivateKeyPath)
	if err != nil {
		return nil, err
	}
	return &ssh.ClientConfig{
		// FixedHostKey instead?
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            user,
		Auth: []ssh.AuthMethod{
			auth,
		},
	}, nil
}

func publicKeyAuth(sshPrivateKeyPath string) (ssh.AuthMethod, error) {
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
