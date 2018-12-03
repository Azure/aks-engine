// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package unstable

import (
	"bytes"

	"github.com/Azure/aks-engine/pkg/openshift/filesystem"
)

// OpenShiftSetDefaultCerts sets default certificate and configuration properties in the
// openshift orchestrator.
func OpenShiftSetDefaultCerts(c *Config) ([]byte, []byte, error) {
	err := c.PrepareMasterCerts()
	if err != nil {
		return nil, nil, err
	}
	err = c.PrepareMasterKubeConfigs()
	if err != nil {
		return nil, nil, err
	}
	err = c.PrepareMasterFiles()
	if err != nil {
		return nil, nil, err
	}

	err = c.PrepareBootstrapKubeConfig()
	if err != nil {
		return nil, nil, err
	}

	masterBundle, err := getConfigBundle(c.WriteMaster)
	if err != nil {
		return nil, nil, err
	}

	nodeBundle, err := getConfigBundle(c.WriteNode)
	if err != nil {
		return nil, nil, err
	}

	return masterBundle, nodeBundle, nil
}

type writeFn func(filesystem.Writer) error

func getConfigBundle(write writeFn) ([]byte, error) {
	b := &bytes.Buffer{}

	fs, err := filesystem.NewTGZWriter(b)
	if err != nil {
		return nil, err
	}

	err = write(fs)
	if err != nil {
		return nil, err
	}

	err = fs.Close()
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
