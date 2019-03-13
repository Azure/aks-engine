// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"
)

func TestNewRotateCertsCmd(t *testing.T) {
	output := newRotateCertsCmd()
	if output.Use != rotateCertsName || output.Short != rotateCertsShortDescription || output.Long != otateCertsLongDescription {
		t.Fatalf("rotate-certs command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, rotateCertsName, output.Short, rotateCertsShortDescription, output.Long, otateCertsLongDescription)
	}

	expectedFlags := []string{"location", "resource-group", "deployment-dir", "new-node-count", "node-pool", "master-FQDN"}
	for _, f := range expectedFlags {
		if output.Flags().Lookup(f) == nil {
			t.Fatalf("rotate-certs command should have flag %s", f)
		}
	}
}
