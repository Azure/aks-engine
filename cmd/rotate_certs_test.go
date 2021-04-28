// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: nil,
			name:        "Valid input",
		},
		{
			rcc: &rotateCertsCmd{
				linuxSSHPrivateKeyPath: existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--api-model must be specified"),
			name:        "Missing api-model",
		},
		{
			rcc: &rotateCertsCmd{
				apiModelPath:           missingFile,
				linuxSSHPrivateKeyPath: existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.Errorf("specified --api-model does not exist (%s)", missingFile),
			name:        "Invalid api-model",
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
				newCertsPath:           existingFile,
				sshHostURI:             "server.example.com",
				location:               "southcentralus",
			},
			expectedErr: nil,
			assert: func(rcc *rotateCertsCmd) {
				g.Expect(rcc.generateCerts).To(Equal(false))
			},
			name: "Unset generateCerts if newCertsPath is set",
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			c.rcc.authProvider = &mockAuthProvider{
				authArgs:      &authArgs{},
				getClientMock: &armhelpers.MockAKSEngineClient{},
			}
			cmd := &cobra.Command{}
			f := cmd.Flags()
			addAuthFlags(c.rcc.getAuthArgs(), f)
			fakeRawSubscriptionID := "6dc93fae-9a76-421f-bbe5-cc6460ea81cb"
			fakeSubscriptionID, err := uuid.Parse(fakeRawSubscriptionID)
			fakeClientID := "b829b379-ca1f-4f1d-91a2-0d26b244680d"
			fakeClientSecret := "0se43bie-3zs5-303e-aav5-dcf231vb82ds"
			if err != nil {
				t.Fatalf("Invalid SubscriptionId in Test: %s", err)
			}
			c.rcc.getAuthArgs().SubscriptionID = fakeSubscriptionID
			c.rcc.getAuthArgs().rawSubscriptionID = fakeRawSubscriptionID
			c.rcc.getAuthArgs().rawClientID = fakeClientID
			c.rcc.getAuthArgs().ClientSecret = fakeClientSecret

			err = c.rcc.validateArgs()
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
