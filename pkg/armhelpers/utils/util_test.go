// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package utils

import (
	"fmt"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

func Test_SplitBlobURI(t *testing.T) {
	expectedAccountName := "vhdstorage8h8pjybi9hbsl6"
	expectedContainerName := "vhds"
	expectedBlobPath := "osdisks/disk1234.vhd"
	accountName, containerName, blobPath, err := SplitBlobURI("https://vhdstorage8h8pjybi9hbsl6.blob.core.windows.net/vhds/osdisks/disk1234.vhd")
	if accountName != expectedAccountName {
		t.Fatalf("incorrect account name. expected=%s actual=%s", expectedAccountName, accountName)
	}
	if containerName != expectedContainerName {
		t.Fatalf("incorrect account name. expected=%s actual=%s", expectedContainerName, containerName)
	}
	if blobPath != expectedBlobPath {
		t.Fatalf("incorrect account name. expected=%s actual=%s", expectedBlobPath, blobPath)
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func Test_LinuxVMNameParts(t *testing.T) {
	data := []struct {
		poolIdentifier, nameSuffix string
		agentIndex                 int
	}{
		{"agentpool1", "38988164", 10},
		{"agent-pool1", "38988164", 8},
		{"agent-pool-1", "38988164", 0},
	}

	for _, el := range data {
		vmName := fmt.Sprintf("k8s-%s-%s-%d", el.poolIdentifier, el.nameSuffix, el.agentIndex)
		poolIdentifier, nameSuffix, agentIndex, err := K8sLinuxVMNameParts(vmName)
		if poolIdentifier != el.poolIdentifier {
			t.Fatalf("incorrect poolIdentifier. expected=%s actual=%s", el.poolIdentifier, poolIdentifier)
		}
		if nameSuffix != el.nameSuffix {
			t.Fatalf("incorrect nameSuffix. expected=%s actual=%s", el.nameSuffix, nameSuffix)
		}
		if agentIndex != el.agentIndex {
			t.Fatalf("incorrect agentIndex. expected=%d actual=%d", el.agentIndex, agentIndex)
		}
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}
}

func Test_VMSSNameParts(t *testing.T) {
	data := []struct {
		poolIdentifier, nameSuffix string
	}{
		{"agentpool1", "38988164"},
		{"agent-pool1", "38988164"},
		{"agent-pool-1", "38988164"},
	}

	for _, el := range data {
		vmssName := fmt.Sprintf("swarmm-%s-%s-vmss", el.poolIdentifier, el.nameSuffix)
		poolIdentifier, nameSuffix, err := VmssNameParts(vmssName)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if poolIdentifier != el.poolIdentifier {
			t.Fatalf("incorrect poolIdentifier. expected=%s actual=%s", el.poolIdentifier, poolIdentifier)
		}
		if nameSuffix != el.nameSuffix {
			t.Fatalf("incorrect nameSuffix. expected=%s actual=%s", el.nameSuffix, nameSuffix)
		}
	}
}

func Test_WindowsVMSSNameParts(t *testing.T) {
	data := []struct {
		poolIdentifier string
	}{
		{"winpo1"},
		{"win2po"},
	}

	for _, el := range data {
		vmssName := fmt.Sprintf("aks%s", el.poolIdentifier)
		poolIdentifier, err := WindowsVmssNameParts(vmssName)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if poolIdentifier != el.poolIdentifier {
			t.Fatalf("incorrect poolIdentifier. expected=%s actual=%s", el.poolIdentifier, poolIdentifier)
		}
	}
}

func Test_WindowsVMNameParts(t *testing.T) {
	data := []struct {
		VMName, expectedPoolPrefix, expectedOrch string
		expectedPoolIndex, expectedAgentIndex    int
	}{
		{"38988k8s90312", "38988", "k8s", 3, 12},
		{"4506k8s010", "4506", "k8s", 1, 0},
		{"2314k8s03000001", "2314", "k8s", 3, 1},
		{"2314k8s0310", "2314", "k8s", 3, 10},
	}

	for _, d := range data {
		poolPrefix, orch, poolIndex, agentIndex, err := WindowsVMNameParts(d.VMName)
		if poolPrefix != d.expectedPoolPrefix {
			t.Fatalf("incorrect poolPrefix. expected=%s actual=%s", d.expectedPoolPrefix, poolPrefix)
		}
		if orch != d.expectedOrch {
			t.Fatalf("incorrect acs string. expected=%s actual=%s", d.expectedOrch, orch)
		}
		if poolIndex != d.expectedPoolIndex {
			t.Fatalf("incorrect poolIndex. expected=%d actual=%d", d.expectedPoolIndex, poolIndex)
		}
		if agentIndex != d.expectedAgentIndex {
			t.Fatalf("incorrect agentIndex. expected=%d actual=%d", d.expectedAgentIndex, agentIndex)
		}
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}
}

func Test_GetVMNameIndexUnknown(t *testing.T) {
	var ostype compute.OperatingSystemTypes
	_, err := GetVMNameIndex(ostype, "k8s-agentpool1-38988164-65")
	if err == nil {
		t.Fatalf("unexpected error but got none")
	}
}

func Test_GetVMNameIndexLinux(t *testing.T) {
	expectedAgentIndex := 65

	agentIndex, err := GetVMNameIndex(compute.Linux, "k8s-agentpool1-38988164-65")

	if agentIndex != expectedAgentIndex {
		t.Fatalf("incorrect agentIndex. expected=%d actual=%d", expectedAgentIndex, agentIndex)
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func Test_GetVMNameIndexWindows(t *testing.T) {
	expectedAgentIndex := 20

	agentIndex, err := GetVMNameIndex(compute.Windows, "38988k8s90320")

	if agentIndex != expectedAgentIndex {
		t.Fatalf("incorrect agentIndex. expected=%d actual=%d", expectedAgentIndex, agentIndex)
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func Test_GetK8sVMName(t *testing.T) {
	p := &api.Properties{
		OrchestratorProfile: &api.OrchestratorProfile{
			OrchestratorType: api.Kubernetes,
		},
		HostedMasterProfile: &api.HostedMasterProfile{
			DNSPrefix: "foo",
		},
		AgentPoolProfiles: []*api.AgentPoolProfile{
			{
				Name:   "linux1",
				VMSize: "Standard_D2_v2",
				Count:  3,
				OSType: "Linux",
			},
			{
				Name:   "windows2",
				VMSize: "Standard_D2_v2",
				Count:  2,
				OSType: "Windows",
			},
			{
				Name:   "someotherpool",
				VMSize: "Standard_D2_v2",
				Count:  5,
				OSType: "Linux",
			},
		},
	}

	for _, s := range []struct {
		properties                 *api.Properties
		agentPoolIndex, agentIndex int
		expected                   string
		expectedErr                bool
	}{
		{properties: p, agentPoolIndex: 0, agentIndex: 2, expected: "aks-linux1-28513887-2", expectedErr: false},
		{properties: p, agentPoolIndex: 1, agentIndex: 1, expected: "2851aks011", expectedErr: false},
	} {
		vmName, err := GetK8sVMName(s.properties, p.AgentPoolProfiles[s.agentPoolIndex], s.agentIndex)

		if !s.expectedErr {
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
		}
		if vmName != s.expected {
			t.Fatalf("Got vmName %s, expected %s", vmName, s.expected)
		}
	}
}

func Test_ResourceName(t *testing.T) {
	s := "https://vhdstorage8h8pjybi9hbsl6.blob.core.windows.net/vhds/osdisks/disk1234.vhd"
	expected := "disk1234.vhd"
	r, err := ResourceName(s)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if r != expected {
		t.Fatalf("resourceName %s, expected %s", r, expected)
	}
}

func Test_ResourceNameInvalid(t *testing.T) {
	s := "https://vhdstorage8h8pjybi9hbsl6.blob.core.windows.net/vhds/osdisks/"
	expectedMsg := "resource name was missing from identifier"
	_, err := ResourceName(s)
	if err == nil || err.Error() != expectedMsg {
		t.Fatalf("expected error with message: %s", expectedMsg)
	}
}
