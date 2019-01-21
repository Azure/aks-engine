// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package config

import (
	"bytes"
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
	ini "gopkg.in/ini.v1"
)

const settingsFixture = `
{
	"profiles": {
		"development": {
			"aks-engine": {
				"debug": true,
				"deploy": {
					"location": "westus2"
				},
				"scale": {
					"location": "westus2"
				},
				"upgrade": {
					"location": "westus2"
				}
			}
		},
		"production": {
			"auth": {
				"auth-method": "client_secret",
				"client-id": "abc123",
				"client-secret": "def456"
			},
			"aks-engine": {
				"deploy": {
					"location": "eastus",
					"set": ["key1=val1", "key2=val2"]
				},
				"scale": {
					"location": "eastus"
				},
				"upgrade": {
					"location": "eastus"
				}
			}
		}
	},
	"auth": {
		"auth-method": "cli",
		"subscription-id": "ghi789",
		"azure-env": "AzurePublicCloud"
	}
}`

func TestFromReader(t *testing.T) {

	expected := &Settings{
		Profiles: ProfileMap{
			"development": Config{
				CLIConfig: CLIConfig{
					Debug: true,
					Deploy: DeployConfig{
						Location: "westus2",
					},
					Scale: ScaleConfig{
						Location: "westus2",
					},
					Upgrade: UpgradeConfig{
						Location: "westus2",
					},
				},
			},
			"production": Config{
				Auth: AuthConfig{
					AuthMethod:   "client_secret",
					ClientID:     "abc123",
					ClientSecret: "def456",
				},
				CLIConfig: CLIConfig{
					Deploy: DeployConfig{
						Location: "eastus",
						Set: []string{
							"key1=val1",
							"key2=val2",
						},
					},
					Scale: ScaleConfig{
						Location: "eastus",
					},
					Upgrade: UpgradeConfig{
						Location: "eastus",
					},
				},
			},
		},
		Auth: AuthConfig{
			AuthMethod:       "cli",
			SubscriptionID:   "ghi789",
			AzureEnvironment: "AzurePublicCloud",
		},
	}

	settings, err := FromReader(bytes.NewBufferString(settingsFixture))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(settings, expected) {
		t.Errorf("Settings are not identical.\nExpected: %v\nActual: %v", expected, settings)
	}
}

func TestUsingProfile(t *testing.T) {
	settings, err := FromReader(bytes.NewBufferString(settingsFixture))
	if err != nil {
		t.Error(err)
	}
	expectedProduction := &Config{
		Auth: AuthConfig{
			AuthMethod:       "client_secret",
			ClientID:         "abc123",
			ClientSecret:     "def456",
			SubscriptionID:   "ghi789",
			AzureEnvironment: "AzurePublicCloud",
		},
		CLIConfig: CLIConfig{
			Deploy: DeployConfig{
				Location: "eastus",
				Set: []string{
					"key1=val1",
					"key2=val2",
				},
			},
			Scale: ScaleConfig{
				Location: "eastus",
			},
			Upgrade: UpgradeConfig{
				Location: "eastus",
			},
		},
	}

	expectedDevelopment := &Config{
		Auth: AuthConfig{
			AuthMethod:       "cli",
			SubscriptionID:   "ghi789",
			AzureEnvironment: "AzurePublicCloud",
		},
		CLIConfig: CLIConfig{
			Debug: true,
			Deploy: DeployConfig{
				Location: "westus2",
			},
			Scale: ScaleConfig{
				Location: "westus2",
			},
			Upgrade: UpgradeConfig{
				Location: "westus2",
			},
		},
	}

	actualProduction := settings.UsingProfile("production")
	if !reflect.DeepEqual(expectedProduction, actualProduction) {
		t.Errorf("Configuration is not identical.\nExpected: %v\nActual: %v", expectedProduction, actualProduction)
	}

	actualDevelopment := settings.UsingProfile("development")
	if !reflect.DeepEqual(expectedDevelopment, actualDevelopment) {
		t.Errorf("Configuration is not identical.\nExpected: %v\nActual: %v", expectedDevelopment, actualDevelopment)
	}
}

func TestGetSelectedCloudFromAzConfig(t *testing.T) {
	for _, test := range []struct {
		desc   string
		data   []byte
		expect string
	}{
		{"nil file", nil, "AzureCloud"},
		{"empty file", []byte{}, "AzureCloud"},
		{"no cloud section", []byte(`
		[key]
		foo = bar
		`), "AzureCloud"},
		{"cloud section empty", []byte(`
		[cloud]
		[foo]
		foo = bar
		`), "AzureCloud"},
		{"AzureCloud selected", []byte(`
		[cloud]
		name = AzureCloud
		`), "AzureCloud"},
		{"custom cloud", []byte(`
		[cloud]
		name = myCloud
		`), "myCloud"},
	} {
		t.Run(test.desc, func(t *testing.T) {
			f, err := ini.Load(test.data)
			if err != nil {
				t.Fatal(err)
			}

			cloud := getSelectedCloudFromAzConfig(f)
			if cloud != test.expect {
				t.Fatalf("exepcted %q, got %q", test.expect, cloud)
			}
		})
	}
}

func TestGetCloudSubFromAzConfig(t *testing.T) {
	goodUUID, err := uuid.FromString("ccabad21-ea42-4ea1-affc-17ae73f9df66")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		desc   string
		data   []byte
		expect uuid.UUID
		err    bool
	}{
		{"empty file", []byte{}, uuid.UUID{}, true},
		{"no entry for cloud", []byte(`
		[SomeCloud]
		subscription = 00000000-0000-0000-0000-000000000000
		`), uuid.UUID{}, true},
		{"invalid UUID", []byte(`
		[AzureCloud]
		subscription = not-a-good-value
		`), uuid.UUID{}, true},
		{"real UUID", []byte(`
		[AzureCloud]
		subscription = ` + goodUUID.String() + `
		`), goodUUID, false},
	} {
		t.Run(test.desc, func(t *testing.T) {
			f, err := ini.Load(test.data)
			if err != nil {
				t.Fatal(err)
			}

			uuid, err := getCloudSubFromAzConfig("AzureCloud", f)
			if test.err != (err != nil) {
				t.Fatalf("expected err=%v, got: %v", test.err, err)
			}
			if test.err {
				return
			}
			if uuid.String() != test.expect.String() {
				t.Fatalf("expected %s, got %s", test.expect, uuid)
			}
		})
	}
}
