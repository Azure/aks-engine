// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/helpers/ssh"
	"github.com/google/go-cmp/cmp"
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
	t.Parallel()
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
				uploadSASURL:           "https://blob-service-uri/?sas-token",
			},
			expectedErr: errors.New("invalid upload SAS URL format, expected 'https://{blob-service-uri}/{container-name}?{sas-token}'"),
			name:        "InvalidSASURLNoPath",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				uploadSASURL:           "https://blob-service-uri//?sas-token",
			},
			expectedErr: errors.New("invalid upload SAS URL format, expected 'https://{blob-service-uri}/{container-name}?{sas-token}'"),
			name:        "InvalidSASURLEmptyPath",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				uploadSASURL:           "https://blob-service-uri//folder-name?sas-token",
			},
			expectedErr: errors.New("invalid upload SAS URL format, expected 'https://{blob-service-uri}/{container-name}?{sas-token}'"),
			name:        "InvalidSASURLNoContainerName",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				uploadSASURL:           "https://blob-service-uri/container-name/folder-name?sas-token",
			},
			expectedErr: nil,
			name:        "ValidSASURLWithDirectory",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				nodeNames:              []string{},
			},
			expectedErr: errors.New("--vm-names cannot be empty"),
			name:        "EmptyVMList",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				nodeNames:              []string{"vm1,vm2"},
				controlPlaneOnly:       true,
			},
			expectedErr: errors.New("--control-plane-only and --vm-names are mutually exclusive"),
			name:        "ControlPlane+VMNames",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        existingFile,
				windowsScriptPath:      existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				uploadSASURL:           "https://blob-service-uri/container-name?sas-token",
				nodeNames:              []string{"vm1,vm2"},
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

func TestGetLogsInit(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	existingFile := "../main.go"
	missingFile := "./random/file"
	cases := []struct {
		glc          *getLogsCmd
		hasWindows   bool
		isSSHEnabled bool
		expectedErr  error
		name         string
	}{
		{
			glc: &getLogsCmd{
				linuxScriptPath:   "",
				windowsScriptPath: "",
				cs:                api.CreateMockContainerService("test", "", 1, 1, false),
			},
			hasWindows:   true,
			isSSHEnabled: true,
			expectedErr:  nil,
			name:         "use VHD scripts",
		},
		{
			glc: &getLogsCmd{
				linuxScriptPath:   existingFile,
				windowsScriptPath: existingFile,
				cs:                api.CreateMockContainerService("test", "", 1, 1, false),
			},
			hasWindows:   true,
			isSSHEnabled: true,
			expectedErr:  nil,
			name:         "use VHD scripts",
		},
		{
			glc: &getLogsCmd{
				linuxScriptPath:   existingFile,
				windowsScriptPath: existingFile,
				cs:                api.CreateMockContainerService("test", "", 1, 1, false),
			},
			hasWindows:   true,
			isSSHEnabled: false,
			expectedErr:  nil,
			name:         "windows ssh disabled",
		},
		{
			glc: &getLogsCmd{
				linuxScriptPath:   existingFile,
				windowsScriptPath: existingFile,
				cs:                api.CreateMockContainerService("test", "", 1, 1, false),
			},
			hasWindows:   false,
			isSSHEnabled: false,
			expectedErr:  nil,
			name:         "no windows pool",
		},
		{
			glc: &getLogsCmd{
				linuxScriptPath:   missingFile,
				windowsScriptPath: existingFile,
				cs:                api.CreateMockContainerService("test", "", 1, 1, false),
			},
			hasWindows:   false,
			isSSHEnabled: false,
			expectedErr:  errors.Errorf("error reading log collection script %s: open %s: no such file or directory", missingFile, missingFile),
			name:         "bad custom linux script",
		},
		{
			glc: &getLogsCmd{
				linuxScriptPath:   existingFile,
				windowsScriptPath: missingFile,
				cs:                api.CreateMockContainerService("test", "", 1, 1, false),
			},
			hasWindows:   false,
			isSSHEnabled: false,
			expectedErr:  errors.Errorf("error reading log collection script %s: open %s: no such file or directory", missingFile, missingFile),
			name:         "bad custom windows script",
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			if c.hasWindows {
				c.glc.cs.Properties.WindowsProfile = api.GetK8sDefaultProperties(c.hasWindows).WindowsProfile
				c.glc.cs.Properties.WindowsProfile.SSHEnabled = &c.isSSHEnabled
			}
			err := c.glc.init()
			if c.expectedErr != nil {
				g.Expect(err.Error()).To(Equal(c.expectedErr.Error()))
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(c.glc.jumpbox).ToNot(BeNil())
				g.Expect(c.glc.linuxAuthConfig).ToNot(BeNil())
				g.Expect(c.glc.linuxVHDScript).ToNot(BeNil())
				g.Expect(c.glc.linuxVHDScript.Path).To(Equal(getLogsLinuxVHDScriptPath))
				g.Expect(c.glc.linuxCustomScript != nil).To(Equal(c.glc.linuxScriptPath != ""))
				if c.glc.linuxScriptPath != "" {
					g.Expect(c.glc.linuxCustomScript.Path).To(Equal(getLogsCustomLinuxScriptPath))
				}
				g.Expect(c.glc.windowsAuthConfig != nil).To(Equal(c.hasWindows && c.isSSHEnabled))
				g.Expect(c.glc.windowsVHDScript).ToNot(BeNil())
				g.Expect(c.glc.windowsVHDScript.Path).To(Equal(getLogsWindowsVHDScriptPath))
				g.Expect(c.glc.windowsCustomScript != nil).To(Equal(c.glc.windowsScriptPath != ""))
				if c.glc.windowsScriptPath != "" {
					g.Expect(c.glc.windowsCustomScript.Path).To(Equal(getLogsCustomWindowsScriptPath))
				}
			}
		})
	}
}

func TestGetLogsGetClusterNodes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	master := &ssh.RemoteHost{URI: "k8s-master-22998975-0", OperatingSystem: api.Linux}
	linuxAgent := &ssh.RemoteHost{URI: "k8s-agentpool1-22998975-0", OperatingSystem: api.Linux}
	windowsAgent := &ssh.RemoteHost{URI: "windows10", OperatingSystem: api.Windows}
	cases := []struct {
		glc                 *getLogsCmd
		isWindowsSSHEnabled bool
		nodeList            []string
		failListNodes       bool
		expected            []*ssh.RemoteHost
		expectedErr         error
		name                string
	}{
		{
			glc: &getLogsCmd{
				controlPlaneOnly: false,
				cs:               api.CreateMockContainerService("test", "", 1, 1, false),
			},
			isWindowsSSHEnabled: true,
			nodeList:            []string{"k8s-master-22998975-0", "k8s-agentpool1-22998975-0", "windows10"},
			failListNodes:       true,
			expected:            []*ssh.RemoteHost{master},
			name:                "cannot retrieve node list from apiserver",
		},
		{
			glc: &getLogsCmd{
				controlPlaneOnly: true,
				cs:               api.CreateMockContainerService("test", "", 1, 1, false),
			},
			isWindowsSSHEnabled: true,
			nodeList:            []string{"k8s-master-22998975-0", "k8s-agentpool1-22998975-0", "windows10"},
			failListNodes:       false,
			expected:            []*ssh.RemoteHost{master},
			name:                "control plane only",
		},
		{
			glc: &getLogsCmd{
				controlPlaneOnly: false,
				cs:               api.CreateMockContainerService("test", "", 1, 1, false),
			},
			isWindowsSSHEnabled: false,
			nodeList:            []string{"k8s-master-22998975-0", "k8s-agentpool1-22998975-0", "windows10"},
			failListNodes:       false,
			expected:            []*ssh.RemoteHost{master, linuxAgent},
			name:                "windows ssh not enabled",
		},
		{
			glc: &getLogsCmd{
				controlPlaneOnly: false,
				cs:               api.CreateMockContainerService("test", "", 1, 1, false),
			},
			isWindowsSSHEnabled: true,
			nodeList:            []string{"k8s-master-22998975-0", "k8s-agentpool1-22998975-0", "windows10"},
			failListNodes:       false,
			expected:            []*ssh.RemoteHost{master, linuxAgent, windowsAgent},
			name:                "expect all nodes",
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			if c.isWindowsSSHEnabled {
				c.glc.windowsAuthConfig = &ssh.AuthConfig{}
			}
			nodes := getClusterNodes(c.glc, &mockNodeLister{
				failListNodes: c.failListNodes,
				nodeNameList:  c.nodeList,
			})
			g.Expect(nodes).ToNot(BeNil())
			g.Expect(len(nodes)).To(Equal(len(c.expected)))
			opt := cmp.Comparer(func(x, y *ssh.RemoteHost) bool {
				return x.URI == y.URI && x.OperatingSystem == y.OperatingSystem
			})
			g.Expect(cmp.Equal(nodes, c.expected, opt)).To(BeTrue())
		})
	}
}

func TestGetLogsGetClusterNodeScripts(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	linuxVHDScript := &ssh.RemoteFile{Path: "linuxVHD"}
	linuxCustomScript := &ssh.RemoteFile{Path: "linuxCustom"}
	windowsVHDScript := &ssh.RemoteFile{Path: "winVHD"}
	windowsCustomScript := &ssh.RemoteFile{Path: "winCustom"}
	master := &ssh.RemoteHost{URI: "k8s-master-22998975-0", OperatingSystem: api.Linux}
	linuxAgent := &ssh.RemoteHost{URI: "k8s-agentpool1-22998975-0", OperatingSystem: api.Linux}
	windowsAgent := &ssh.RemoteHost{URI: "windows10", OperatingSystem: api.Windows}
	cases := []struct {
		glc      *getLogsCmd
		isVHD    bool
		nodes    []*ssh.RemoteHost
		expected map[*ssh.RemoteHost]*ssh.RemoteFile
		name     string
	}{
		{
			glc: &getLogsCmd{
				cs:                  api.CreateMockContainerService("test", "", 1, 1, false),
				linuxVHDScript:      linuxVHDScript,
				linuxCustomScript:   linuxCustomScript,
				windowsVHDScript:    windowsVHDScript,
				windowsCustomScript: windowsCustomScript,
			},
			isVHD: true,
			nodes: []*ssh.RemoteHost{master, linuxAgent, windowsAgent},
			expected: map[*ssh.RemoteHost]*ssh.RemoteFile{
				master:       linuxCustomScript,
				linuxAgent:   linuxCustomScript,
				windowsAgent: windowsCustomScript,
			},
			name: "vhd and custom scripts",
		},
		{
			glc: &getLogsCmd{
				cs:                  api.CreateMockContainerService("test", "", 1, 1, false),
				linuxVHDScript:      linuxVHDScript,
				linuxCustomScript:   linuxCustomScript,
				windowsVHDScript:    windowsVHDScript,
				windowsCustomScript: windowsCustomScript,
			},
			isVHD: false,
			nodes: []*ssh.RemoteHost{master, linuxAgent, windowsAgent},
			expected: map[*ssh.RemoteHost]*ssh.RemoteFile{
				master:       linuxCustomScript,
				linuxAgent:   linuxCustomScript,
				windowsAgent: windowsCustomScript,
			},
			name: "not vhd and custom scripts",
		},
		{
			glc: &getLogsCmd{
				cs:               api.CreateMockContainerService("test", "", 1, 1, false),
				linuxVHDScript:   linuxVHDScript,
				windowsVHDScript: windowsVHDScript,
			},
			isVHD: true,
			nodes: []*ssh.RemoteHost{master, linuxAgent, windowsAgent},
			expected: map[*ssh.RemoteHost]*ssh.RemoteFile{
				master:       linuxVHDScript,
				linuxAgent:   linuxVHDScript,
				windowsAgent: windowsVHDScript,
			},
			name: "vhd and no custom scripts",
		},
		{
			glc: &getLogsCmd{
				cs:               api.CreateMockContainerService("test", "", 1, 1, false),
				linuxVHDScript:   linuxVHDScript,
				windowsVHDScript: windowsVHDScript,
			},
			isVHD:    false,
			nodes:    []*ssh.RemoteHost{master, linuxAgent, windowsAgent},
			expected: map[*ssh.RemoteHost]*ssh.RemoteFile{},
			name:     "not vhd and no custom scripts",
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			if c.isVHD {
				c.glc.cs.Properties.MasterProfile.Distro = api.AKSUbuntu1604
				c.glc.cs.Properties.AgentPoolProfiles[0].Distro = api.AKSUbuntu1604
				c.glc.cs.Properties.WindowsProfile = api.GetK8sDefaultProperties(c.isVHD).WindowsProfile
				c.glc.cs.Properties.WindowsProfile.WindowsPublisher = api.AKSWindowsServer2019OSImageConfig.ImagePublisher
				c.glc.cs.Properties.WindowsProfile.WindowsOffer = api.AKSWindowsServer2019OSImageConfig.ImageOffer
			}
			nodeScripts := getClusterNodeScripts(c.glc, c.nodes)
			g.Expect(nodeScripts).ToNot(BeNil())
			g.Expect(len(nodeScripts)).To(Equal(len(c.expected)))
			opt := cmp.Comparer(func(x, y *ssh.RemoteHost) bool {
				return x.URI == y.URI && x.OperatingSystem == y.OperatingSystem
			})
			g.Expect(cmp.Equal(nodeScripts, c.expected, opt)).To(BeTrue())
		})
	}
}

type mockNodeLister struct {
	nodeNameList  []string
	failListNodes bool
}

func (m *mockNodeLister) ListNodes() (*v1.NodeList, error) {
	if m.failListNodes {
		return nil, errors.New("error")
	}
	nodeList := &v1.NodeList{}
	for _, name := range m.nodeNameList {
		node := &v1.Node{}
		node.Name = name
		if strings.HasPrefix(name, "k8s-") {
			node.Status.NodeInfo.OperatingSystem = "linux"
		} else {
			node.Status.NodeInfo.OperatingSystem = "windows"
		}
		nodeList.Items = append(nodeList.Items, *node)
	}
	return nodeList, nil
}
