// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd_test

import (
	"testing"

	. "github.com/Azure/aks-engine/pkg/test"
)

func TestCmd(t *testing.T) {

	RunSpecsWithReporters(t, "cmd", "Cmd Suite")

}
