// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package ssh

import (
	"github.com/Azure/aks-engine/pkg/api"
)

type RemoteFile struct {
	Path        string
	Permissions string
	Owner       string
	Content     []byte
}

func NewRemoteFile(path, permissions, owner string, content []byte) *RemoteFile {
	return &RemoteFile{
		Path:        path,
		Permissions: permissions,
		Owner:       owner,
		Content:     content,
	}
}

type AuthConfig struct {
	User           string
	Password       string
	PrivateKeyPath string
}

type JumpBox struct {
	URI             string
	Port            int
	OperatingSystem api.OSType
	AuthConfig      *AuthConfig
}

type RemoteHost struct {
	URI             string
	Port            int
	OperatingSystem api.OSType
	AuthConfig      *AuthConfig
	Jumpbox         *JumpBox
}
