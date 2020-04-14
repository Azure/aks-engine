// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
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
				apiserverURI:           "servier.example.com",
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
				apiserverURI:           "servier.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--linux-ssh-private-key must be specified"),
			name:        "NeedsLinuxSSHPrivateKey",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        "",
				apiserverURI:           "servier.example.com",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--linux-script must be specified"),
			name:        "NeedsLinuxScript",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           existingFile,
				linuxSSHPrivateKeyPath: existingFile,
				linuxScriptPath:        missingFile,
				apiserverURI:           "servier.example.com",
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
				apiserverURI:           "servier.example.com",
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
				apiserverURI:           "",
				location:               "southcentralus",
			},
			expectedErr: errors.New("--apiserver must be specified"),
			name:        "NeedsAPIServer",
		},
		{
			glc: &getLogsCmd{
				apiModelPath:           "",
				linuxSSHPrivateKeyPath: missingFile,
				linuxScriptPath:        existingFile,
				apiserverURI:           "servier.example.com",
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
				apiserverURI:           "servier.example.com",
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
				apiserverURI:           "servier.example.com",
				location:               "southcentralus",
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
