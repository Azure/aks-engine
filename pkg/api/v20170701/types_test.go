package v20170701

import (
	"encoding/json"
	"testing"
)

func TestMasterProfile(t *testing.T) {
	MasterProfileText := "{\"count\" : 0}"
	mp := &MasterProfile{}
	if e := json.Unmarshal([]byte(MasterProfileText), mp); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for MasterProfile, %+v", e)
	}

	if mp.Count != 1 {
		t.Fatalf("unexpectedly detected MasterProfile.Count != 1 after unmarshal")
	}

	if mp.FirstConsecutiveStaticIP != "10.240.255.5" {
		t.Fatalf("unexpectedly detected MasterProfile.FirstConsecutiveStaticIP != 10.240.255.5 after unmarshal")
	}
}

func TestAgentPoolProfile(t *testing.T) {
	// With osType not specified
	AgentPoolProfileText := `{"count" : 0, "storageProfile" : "StorageAccount"}`
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
		t.Fatalf("unexpectedly detected AgentPoolProfile.StorageProfile != ManagedDisks after unmarshal")
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
}

func TestOrchestratorProfile(t *testing.T) {
	OrchestratorProfileText := `{ "orchestratorType": "Mesos" }`
	op := &OrchestratorProfile{}
	if e := json.Unmarshal([]byte(OrchestratorProfileText), op); e == nil {
		t.Fatalf("expected unmarshal failure for OrchestratorProfile when passing an invalid orchestratorType")
	}

	OrchestratorProfileText = `{ "orchestratorType": "Kubernetes" }`
	op = &OrchestratorProfile{}
	if e := json.Unmarshal([]byte(OrchestratorProfileText), op); e != nil {
		t.Fatalf("unexpectedly detected unmarshal failure for OrchestratorProfile, %+v", e)
	}
}
