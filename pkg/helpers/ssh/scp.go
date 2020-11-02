// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package ssh

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/pkg/errors"
)

// CopyToRemote copies a file to a remote host
func CopyToRemote(host *RemoteHost, file *RemoteFile) (combinedOutput string, err error) {
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
	// Make this configurable if we find that consumers need to update the command
	cmd := getUploadCommand(host.OperatingSystem)(file)
	s.Stdin = bytes.NewReader(file.Content)
	if co, err := s.CombinedOutput(cmd); err != nil {
		return string(co), errors.Wrap(err, "uploading to remote host")
	}
	return "", nil
}

// CopyFromRemote copies a remote file to the local host
func CopyFromRemote(host *RemoteHost, remoteFile *RemoteFile, destinationPath string) (stderr string, err error) {
	f, err := os.OpenFile(destinationPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", errors.Wrap(err, "opening destination file")
	}
	defer f.Close()
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
	stdout, err := s.StdoutPipe()
	if err != nil {
		return "", errors.Wrap(err, "opening SSH session stdout pipe")
	}
	// Make this configurable if we find that consumers need to update the command
	cmd := getDownloadCommand(host.OperatingSystem)(remoteFile)
	if err = s.Start(cmd); err != nil {
		return fmt.Sprintf("%s", s.Stderr), errors.Wrap(err, "downloading logs from remote host")
	}
	_, err = io.Copy(f, stdout)
	if err != nil {
		return "", errors.Wrap(err, "downloading logs")
	}
	return "", nil
}

type uploadCommandBuilder func(file *RemoteFile) string

func getUploadCommand(os api.OSType) uploadCommandBuilder {
	switch os {
	case api.Linux:
		return func(f *RemoteFile) string {
			return fmt.Sprintf("sudo bash -c \"mkdir -p $(dirname %s); cat /dev/stdin > %s; chmod %s %s; chown %s %s\"",
				f.Path, f.Path, f.Permissions, f.Path, f.Owner, f.Path)
		}
	case api.Windows:
		return func(f *RemoteFile) string {
			return fmt.Sprintf("powershell -noprofile -command \"$Input | Out-File -Encoding ASCII %s\"",
				f.Path)
		}
	default:
		return nil
	}
}

type downloadCommandBuilder func(file *RemoteFile) string

func getDownloadCommand(os api.OSType) downloadCommandBuilder {
	switch os {
	case api.Linux:
		return func(f *RemoteFile) string {
			return fmt.Sprintf("bash -c \"cat %s > /dev/stdout\"", f.Path)
		}
	case api.Windows:
		return func(f *RemoteFile) string {
			return fmt.Sprintf("type %s", f.Path)
		}
	default:
		return nil
	}
}
