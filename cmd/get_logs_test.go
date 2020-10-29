// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
)

func TestGetLogsCmd(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	command := newGetLogsCmd()
	g.Expect(command.Use).Should(Equal(getLogsName))
	g.Expect(command.Short).Should(Equal(getLogsShortDescription))
	g.Expect(command.Long).Should(Equal(getLogsLongDescription))

	command.SetArgs([]string{})
	err := command.Execute()
	g.Expect(err).To(HaveOccurred())
}

func TestGetLogsCmdValidateArgs(t *testing.T) {
	g := NewGomegaWithT(t)

	existingFile := "../examples/kubernetes.json"
	missingFile := "./random/file"

	cases := []struct {
		glc         *getLogsCmd
		expectedErr error
		name        string
	}{
		{
			glc: &getLogsCmd{
				apiModelPath:           missingFile,
				linuxSSHPrivateKeyPath: "",
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.Errorf("specified --api-model does not exist (%s)", missingFile),
			name:        "BadLinuxSSHPrivateKey",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: "",
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--linux-ssh-private-key must be specified"),
			name:        "NeedsLinuxSSHPrivateKey",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        missingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.Errorf("specified --linux-script does not exist (%s)", missingFile),
			name:        "BadLinuxScript",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      missingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.Errorf("specified --windows-script does not exist (%s)", missingFile),
			name:        "BadWindowsScript",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "",
			},
			expectedErr: errors.New("--location must be specified"),
			name:        "NeedsLocation",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--ssh-host must be specified"),
			name:        "NeedsSSHHost",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           "",
				linuxSSHPrivateKeyPath: missingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--api-model must be specified"),
			name:        "BadAPIModel",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           missingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.Errorf("specified --api-model does not exist (%s)", missingFile),
			name:        "NeedsAPIModel",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				storageContainerSASURL: "https://blob-service-uri/?sas-token",
			},
			expectedErr: errors.Errorf("invalid upload SAS URL format, expected 'https://{blob-service-uri}/{container-name}?{sas-token}'"),
			name:        "InvalidSASURL",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				storageContainerSASURL: "https://blob-service-uri//?sas-token",
			},
			expectedErr: errors.Errorf("invalid upload SAS URL format, expected 'https://{blob-service-uri}/{container-name}?{sas-token}'"),
			name:        "InvalidSASURL2",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				storageContainerSASURL: "https://blob-service-uri//folder-name?sas-token",
			},
			expectedErr: errors.Errorf("invalid upload SAS URL format, expected 'https://{blob-service-uri}/{container-name}?{sas-token}'"),
			name:        "InvalidSASURL3",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				storageContainerSASURL: "https://blob-service-uri/container-name/folder-name?sas-token",
			},
			expectedErr: nil,
			name:        "ValidSASURL",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				storageContainerSASURL: "https://blob-service-uri/container-name?sas-token",
			},
			expectedErr: nil,
			name:        "IsValid",
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			err := c.glc.validateArgs()
			if c.expectedErr != nil {
				g.Expect(err.Error()).To(Equal(c.expectedErr.Error()))
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
		})
	}
}

func TestComputeControlPlaneNodes(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	nodeList := computeControlPlaneNodes(3, "12345678")
	for i, node := range nodeList {
		g.Expect(node.Name).To(Equal(fmt.Sprintf("k8s-master-12345678-%d", i)))
		g.Expect(node.Status.NodeInfo.OperatingSystem).To(Equal("linux"))
	}
}

func TestFilterNodesFromPool(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	var nodeList []v1.Node
	for i := 0; i < 3; i++ {
		var node1, node2 v1.Node
		node1.Name = fmt.Sprintf("k8s-linuxpool-12345678-%d", i)
		node1.Status.NodeInfo.OperatingSystem = "linux"
		nodeList = append(nodeList, node1)
		node2.Name = fmt.Sprintf("k8s-linuxpoool-12345678-%d", i)
		node2.Status.NodeInfo.OperatingSystem = "linux"
		nodeList = append(nodeList, node2)
	}
	nodeListA := filterNodesFromPool(nodeList, "linuxpool")
	g.Expect(len(nodeListA)).To(Equal(3))
	nodeListB := filterNodesFromPool(nodeList, "linuxpoool")
	g.Expect(len(nodeListB)).To(Equal(3))
	nodeListC := filterNodesFromPool(nodeList, "linuxpol")
	g.Expect(len(nodeListC)).To(Equal(6))
}
