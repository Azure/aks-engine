// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateInboundNatRules(t *testing.T) {
	ibnr := createInboundNATRules(1)
	jsonObj, _ := json.MarshalIndent(ibnr, "", "   ")
	fmt.Println(string(jsonObj))
}
