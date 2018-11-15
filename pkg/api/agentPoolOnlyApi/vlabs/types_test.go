// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"encoding/json"
	"testing"
)

func TestAgentPoolProfile(t *testing.T) {
	// With osType not specified
	AgentPoolProfileText := `{ "name": "linuxpool1", "count": 0, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" }`
	ap := &AgentPoolProfile{}
	if e := json.Unmarshal([]byte(AgentPoolProfileText), ap); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for AgentPoolProfile, %+v", e)
	}

	if ap.Count != 1 {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Count != 1 after unmarshal")
	}

	if !ap.IsLinux() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.OSType != Linux after unmarshal")
	}

	if !ap.IsStorageAccount() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != StorageAccount after unmarshal")
	}

	// With osType specified
	AgentPoolProfileText = `{ "name": "linuxpool1", "osType" : "Windows", "count": 1, "vmSize": "Standard_D2_v2", 
"availabilityProfile": "AvailabilitySet", "storageProfile" : "ManagedDisks", "vnetSubnetID" : "12345" }`
	ap = &AgentPoolProfile{}
	if e := json.Unmarshal([]byte(AgentPoolProfileText), ap); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for AgentPoolProfile, %+v", e)
	}

	if ap.Count != 1 {
		t.Fatalf("unexpectedly detected AgentPoolProfile.Count != 1 after unmarshal")
	}

	if !ap.IsWindows() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.OSType != Windows after unmarshal")
	}

	if !ap.IsManagedDisks() {
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != ManagedDisks after unmarshal")
	}

	if !ap.IsCustomVNET() {
		t.Fatalf("unexpectedly detected empty AgentPoolProfile.VNetSubnetID after unmarshal")
	}
}
