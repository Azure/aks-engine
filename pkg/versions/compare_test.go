// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package versions

import (
	"testing"
)

func assertVersion(t *testing.T, a, b string, result int) {
	if r := compare(a, b); r != result {
		t.Fatalf("Unexpected version comparison result. Found %d, expected %d", r, result)
	}
}

func TestCompareVersion(t *testing.T) {
	assertVersion(t, "1.12", "1.12", 0)
	assertVersion(t, "1.0.0", "1", 0)
	assertVersion(t, "1", "1.0.0", 0)
	assertVersion(t, "1.05.00.0156", "1.0.221.9289", 1)
	assertVersion(t, "1", "1.0.1", -1)
	assertVersion(t, "1.0.1", "1", 1)
	assertVersion(t, "1.0.1", "1.0.2", -1)
	assertVersion(t, "1.0.2", "1.0.3", -1)
	assertVersion(t, "1.0.3", "1.1", -1)
	assertVersion(t, "1.1", "1.1.1", -1)
	assertVersion(t, "1.1.1", "1.1.2", -1)
	assertVersion(t, "1.1.2", "1.2", -1)
	// Check CalVer
	assertVersion(t, "19.03.12", "3.0.10", 1)
	assertVersion(t, "3.0.10", "19.03.12", -1)
	assertVersion(t, "19.03", "3.0.10", 1)
	assertVersion(t, "19.03.12", "19.03", 1)
	assertVersion(t, "19.03.12", "19.04", -1)
}
