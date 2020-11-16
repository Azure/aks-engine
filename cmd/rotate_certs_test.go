// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

func TestNewRotateCertsCmd(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	command := newRotateCertsCmd()
	g.Expect(command.Use).Should(Equal(rotateCertsName))
	g.Expect(command.Short).Should(Equal(rotateCertsShortDescription))
	g.Expect(command.Long).Should(Equal(rotateCertsLongDescription))

	command.SetArgs([]string{})
	err := command.Execute()
	g.Expect(err).To(HaveOccurred())

	for _, f := range []string{"location", "ssh-host", "api-model", "linux-ssh-private-key"} {
		if command.Flags().Lookup(f) == nil {
			t.Fatalf("rotate-certs command should have flag %s", f)
		}
	}
}

func TestRotateCertsCmdValidateArgs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	existingFile := "../examples/kubernetes.json"
	missingFile := "./random/file"

	cases := []struct {
		rcc         *rotateCertsCmd
		expectedErr error
		assert      func(*rotateCertsCmd)
		name        string
	}{
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				newCertsPath:           existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: nil,
			name:        "Valid input",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				sshHostURI:             "server.example.com",
				location:               "",
			},
			expectedErr: errors.New("--location must be specified"),
			name:        "Missing location",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				sshHostURI:             "",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--ssh-host must be specified"),
			name:        "Missing SSH host",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: "",
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--linux-ssh-private-key must be specified"),
			name:        "Missing SSH private key",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: missingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.Errorf("specified --linux-ssh-private-key does not exist (%s)", missingFile),
			name:        "Invalid SSH private key",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				newCertsPath:           missingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.Errorf("specified --certificate-profile does not exist (%s)", missingFile),
			name:        "Invalid new certs profile path",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				newCertsPath:           "",
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				generateCerts:          false,
			},
			expectedErr: errors.New("either --generate-new-certificates or --certificate-profile should be specified"),
			name:        "Either generate certs or pass certs profile path",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				newCertsPath:           "",
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				generateCerts:          false,
			},
			expectedErr: errors.New("either --generate-new-certificates or --certificate-profile should be specified"),
			name:        "Either generate certs or pass certs profile path",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				newCertsPath:           existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
				generateCerts:          true,
			},
			expectedErr: nil,
			assert: func(rcc *rotateCertsCmd) {
				g.Expect(rcc.generateCerts).To(Equal(false), "cannot set both --generate-new-certificates and --certificate-profile")
			},
			name: "Ignore generate certs if certs profile path is also set",
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			err := c.rcc.validateArgs()
			if c.expectedErr != nil {
				g.Expect(err.Error()).To(Equal(c.expectedErr.Error()))
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
			if c.assert != nil {
				c.assert(c.rcc)
			}
		})
	}
}
