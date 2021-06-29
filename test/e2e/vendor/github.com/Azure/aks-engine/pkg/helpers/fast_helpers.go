//+build fast
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

const (
	// SSHKeySize is the size (in bytes) of SSH key to create (only used when `-tags fast` is used)
	SSHKeySize = 512

	// DefaultPkiKeySize is the default size in bytes of the PKI key (only used when `-tags fast` is used)
	DefaultPkiKeySize = 512
)
