package remote

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// Connection is
type Connection struct {
	Host           string
	Port           string
	User           string
	PrivateKeyPath string
	ClientConfig   *ssh.ClientConfig
	Client         *ssh.Client
}

// NewConnection will build and return a new Connection object
func NewConnection(host, port, user, keyPath string) (*Connection, error) {
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		log.Printf("unable to establish net connection $SSH_AUTH_SOCK has value %s\n", os.Getenv("SSH_AUTH_SOCK"))
		return nil, err
	}
	defer conn.Close()
	ag := agent.NewClient(conn)

	privateKeyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := ssh.ParseRawPrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}
	addKey := agent.AddedKey{
		PrivateKey: privateKey,
	}

	ag.Add(addKey)
	signers, err := ag.Signers()
	if err != nil {
		log.Println("unable to add key to agent")
		return nil, err
	}
	auths := []ssh.AuthMethod{ssh.PublicKeys(signers...)}

	cfg := &ssh.ClientConfig{
		User:            user,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	cnctStr := fmt.Sprintf("%s:%s", host, port)
	sshClient, err := ssh.Dial("tcp", cnctStr, cfg)
	if err != nil {
		return nil, err
	}

	return &Connection{
		Host:           host,
		Port:           port,
		User:           user,
		PrivateKeyPath: keyPath,
		ClientConfig:   cfg,
		Client:         sshClient,
	}, nil
}

// Execute will execute a given cmd on a remote host
func (c *Connection) Execute(cmd string) ([]byte, error) {
	session, err := c.Client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.CombinedOutput(cmd)
}

func (c *Connection) Write(data, path string) error {
	remoteCommand := fmt.Sprintf("echo %s > %s", data, path)
	connectString := fmt.Sprintf("%s@%s", c.User, c.Host)
	cmd := exec.Command("ssh", "-i", c.PrivateKeyPath, "-o", "ConnectTimeout=30", "-o", "StrictHostKeyChecking=no", connectString, "-p", c.Port, remoteCommand)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error output:%s\n", out)
		return err
	}
	return nil
}

func (c *Connection) Read(path string) ([]byte, error) {
	remoteCommand := fmt.Sprintf("cat %s", path)
	connectString := fmt.Sprintf("%s@%s", c.User, c.Host)
	cmd := exec.Command("ssh", "-i", c.PrivateKeyPath, "-o", "ConnectTimeout=30", "-o", "StrictHostKeyChecking=no", connectString, "-p", c.Port, remoteCommand)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error output:%s\n", out)
		return nil, err
	}
	return out, nil
}

// CopyRemote uses this ssh connection to scp remote files
func (c *Connection) CopyRemote(hostname, path string) error {
	cmd := exec.Command("ssh-add", c.PrivateKeyPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error output:%s\n", out)
		return err
	}
	remoteCommand := fmt.Sprintf("scp -o StrictHostKeyChecking=no %s:%s /tmp/%s-%s", hostname, path, hostname, filepath.Base(path))
	connectString := fmt.Sprintf("%s@%s", c.User, c.Host)
	cmd = exec.Command("ssh", "-A", "-i", c.PrivateKeyPath, "-o", "ConnectTimeout=30", "-o", "StrictHostKeyChecking=no", connectString, "-p", c.Port, remoteCommand)
	util.PrintCommand(cmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error output:%s\n", out)
		return err
	}
	return nil
}

// ExecuteWithRetries will keep retrying a command until it does not return an error or the duration is exceeded
func (c *Connection) ExecuteWithRetries(cmd string, sleep, duration time.Duration) ([]byte, error) {
	outCh := make(chan []byte, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for command to not return an error: %s", duration.String(), cmd)
			default:
				out, err := c.Execute(cmd)
				if err == nil {
					outCh <- out
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return nil, err
		case out := <-outCh:
			return out, nil
		}
	}
}
