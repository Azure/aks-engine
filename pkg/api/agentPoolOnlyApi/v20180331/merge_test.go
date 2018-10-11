package v20180331

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/helpers"
)

func TestMerge_DNSPrefix(t *testing.T) {
	newMC := &ManagedCluster{
		Properties: &Properties{
			DNSPrefix: "newprefix",
		},
	}

	existingMC := &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "oldprefix",
			EnableRBAC: helpers.PointerToBool(false),
		},
	}

	e := newMC.Merge(existingMC)
	if e == nil {
		t.Error("expect error to not be nil")
	}

	newMC = &ManagedCluster{
		Properties: &Properties{},
	}

	existingMC = &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "oldprefix",
			EnableRBAC: helpers.PointerToBool(false),
		},
	}

	e = newMC.Merge(existingMC)
	if e != nil {
		t.Error("expect error to be nil")
	}

	if newMC.Properties.DNSPrefix != "oldprefix" {
		t.Error("expect dnsPrefix to be oldprefix when update with empty input")
	}

	newMC = &ManagedCluster{
		Properties: &Properties{},
	}

	existingMC = &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "",
			EnableRBAC: helpers.PointerToBool(false),
		},
	}

	e = newMC.Merge(existingMC)
	if e == nil {
		t.Error("expect error to not be nil")
	}
}

func TestMerge_AAD(t *testing.T) {
	// Partial AAD profile was passed during update
	newMC := &ManagedCluster{
		Properties: &Properties{
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: &AADProfile{
				ClientAppID: "1234-5",
				ServerAppID: "1a34-5",
				TenantID:    "c234-5",
			},
		},
	}

	existingMC := &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "something",
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: &AADProfile{
				ClientAppID:     "1234-5",
				ServerAppID:     "1a34-5",
				ServerAppSecret: "ba34-5",
				TenantID:        "c234-5",
			},
		},
	}

	e := newMC.Merge(existingMC)
	if e != nil {
		t.Error("expect error to be nil")
	}

	if newMC.Properties.AADProfile == nil {
		t.Error("AADProfile should not be nil")
	}

	if newMC.Properties.AADProfile.ServerAppSecret == "" {
		t.Error("ServerAppSecret did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ServerAppID != "1a34-5" {
		t.Error("ServerAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ServerAppSecret != "ba34-5" {
		t.Error("ServerAppSecret did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ClientAppID != "1234-5" {
		t.Error("ClientAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.TenantID != "c234-5" {
		t.Error("TenantID did not have the expected value after merge")
	}

	// Nil AAD profile was passed during update but DM had AAD Profile
	newMC = &ManagedCluster{
		Properties: &Properties{
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: nil,
		},
	}

	existingMC = &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "something",
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: &AADProfile{
				ClientAppID:     "1234-5",
				ServerAppID:     "1a34-5",
				ServerAppSecret: "ba34-5",
				TenantID:        "c234-5",
			},
		},
	}

	e = newMC.Merge(existingMC)
	if e != nil {
		t.Error("expect error to be nil")
	}

	if newMC.Properties.AADProfile == nil {
		t.Error("AADProfile should not be nil")
	}

	if newMC.Properties.AADProfile.ServerAppSecret == "" {
		t.Error("ServerAppSecret did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ServerAppID != "1a34-5" {
		t.Error("ServerAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ServerAppSecret != "ba34-5" {
		t.Error("ServerAppSecret did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ClientAppID != "1234-5" {
		t.Error("ClientAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.TenantID != "c234-5" {
		t.Error("TenantID did not have the expected value after merge")
	}

	// No AAD profile set
	newMC = &ManagedCluster{
		Properties: &Properties{
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: nil,
		},
	}

	existingMC = &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "something",
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: nil,
		},
	}

	e = newMC.Merge(existingMC)
	if e != nil {
		t.Error("expect error to be nil")
	}

	if newMC.Properties.AADProfile != nil {
		t.Error("AADProfile should be nil")
	}

	// Empty field in AAD profile was passed during update but DM had AAD Profile
	newMC = &ManagedCluster{
		Properties: &Properties{
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: &AADProfile{
				ClientAppID:     "1234-5",
				ServerAppID:     "1a34-5",
				ServerAppSecret: "",
				TenantID:        "c234-5",
			},
		},
	}

	existingMC = &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "something",
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: &AADProfile{
				ClientAppID:     "1234-5",
				ServerAppID:     "1a34-5",
				ServerAppSecret: "ba34-5",
				TenantID:        "c234-5",
			},
		},
	}

	e = newMC.Merge(existingMC)
	if e != nil {
		t.Error("expect error to be nil")
	}

	if newMC.Properties.AADProfile == nil {
		t.Error("AADProfile should not be nil")
	}

	if newMC.Properties.AADProfile.ServerAppID != "1a34-5" {
		t.Error("ServerAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ServerAppSecret != "ba34-5" {
		t.Error("ServerAppSecret did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ClientAppID != "1234-5" {
		t.Error("ClientAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.TenantID != "c234-5" {
		t.Error("TenantID did not have the expected value after merge")
	}

	// Full AAD profile was passed during update
	newMC = &ManagedCluster{
		Properties: &Properties{
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: &AADProfile{
				ClientAppID:     "1234-5",
				ServerAppID:     "1a34-5",
				ServerAppSecret: "ba34-5",
				TenantID:        "c234-5",
			},
		},
	}

	existingMC = &ManagedCluster{
		Properties: &Properties{
			DNSPrefix:  "something",
			EnableRBAC: helpers.PointerToBool(true),
			AADProfile: &AADProfile{
				ClientAppID:     "1234-5",
				ServerAppID:     "1a34-5",
				ServerAppSecret: "ba34-5",
				TenantID:        "c234-5",
			},
		},
	}

	e = newMC.Merge(existingMC)
	if e != nil {
		t.Error("expect error to be nil")
	}

	if newMC.Properties.AADProfile == nil {
		t.Error("AADProfile should not be nil")
	}

	if newMC.Properties.AADProfile.ServerAppSecret == "" {
		t.Error("ServerAppSecret did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ServerAppID != "1a34-5" {
		t.Error("ServerAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ServerAppSecret != "ba34-5" {
		t.Error("ServerAppSecret did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.ClientAppID != "1234-5" {
		t.Error("ClientAppID did not have the expected value after merge")
	}

	if newMC.Properties.AADProfile.TenantID != "c234-5" {
		t.Error("TenantID did not have the expected value after merge")
	}
}
