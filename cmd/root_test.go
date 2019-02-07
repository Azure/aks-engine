// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/Azure/aks-engine/pkg/cli/config"
	"github.com/imdario/mergo"
)

func TestNewRootCmd(t *testing.T) {
	output := NewRootCmd()
	if output.Use != rootName || output.Short != rootShortDescription || output.Long != rootLongDescription {
		t.Fatalf("root command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, rootName, output.Short, rootShortDescription, output.Long, rootLongDescription)
	}
	expectedCommands := map[string]bool{
		newConfigureCmd().Name():        false,
		getCompletionCmd(output).Name(): false,
		newDeployCmd().Name():           false,
		newGenerateCmd().Name():         false,
		newOrchestratorsCmd().Name():    false,
		newScaleCmd().Name():            false,
		newUpgradeCmd().Name():          false,
		newVersionCmd().Name():          false,
	}
	for _, command := range output.Commands() {
		expectedCommands[command.Name()] = true
	}
	for name, exists := range expectedCommands {
		if !exists {
			t.Errorf("root command should have command %s", name)
		}
	}
}

func TestLoadSettings(t *testing.T) {
	knownEnvVars := []string{
		"AKS_ENGINE_AZURE_ENVIRONMENT",
		"AKS_ENGINE_AUTH_METHOD",
		"AKS_ENGINE_SUBSCRIPTION_ID",
		"AKS_ENGINE_CLIENT_ID",
		"AKS_ENGINE_CLIENT_SECRET",
		"AKS_ENGINE_CERTIFICATE_PATH",
		"AKS_ENGINE_PRIVATE_KEY_PATH",
		"AKS_ENGINE_LANGUAGE",
		"AKS_ENGINE_DEBUG",
		"AKS_ENGINE_PROFILE",
		"AKS_ENGINE_SHOW_DEFAULT_MODEL",
	}
	expectedConfig := defaultConfigValues
	expectedConfig.Auth.AuthMethod = "foo"
	expectedConfig.Auth.ClientID = "bar"
	expectedConfig.Auth.ClientSecret = "car"
	expectedConfig.Auth.Language = "star"
	expectedConfigWithProfile := expectedConfig
	expectedConfigWithProfile.CLIConfig.Profile = "production"
	envVars := map[string]string{
		"AKS_ENGINE_AUTH_METHOD":   "foo",
		"AKS_ENGINE_CLIENT_ID":     "bar",
		"AKS_ENGINE_CLIENT_SECRET": "car",
		"AKS_ENGINE_LANGUAGE":      "star",
	}
	var tests = []struct {
		Name                    string
		CLIFlagConfig           config.Config
		EnvironmentVariables    map[string]string
		SettingsFileContent     string
		ExpectedConfig          config.Config
		FromReaderShouldError   bool
		LoadSettingsShouldError bool
		ExpectedErrorMessage    string
	}{
		{
			Name:           "basic config, nothing configured",
			ExpectedConfig: defaultConfigValues,
		},
		{
			Name: "feature flags set from CLI",
			CLIFlagConfig: config.Config{
				Auth: config.AuthConfig{
					AuthMethod:   "foo",
					ClientID:     "bar",
					ClientSecret: "car",
					Language:     "star",
				},
			},
			ExpectedConfig: expectedConfig,
		},
		{
			Name:                 "feature flags set from environment variables",
			EnvironmentVariables: envVars,
			ExpectedConfig:       expectedConfig,
		},
		{
			Name:                "feature flags set from settings.json",
			SettingsFileContent: `{"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"}}`,
			ExpectedConfig:      expectedConfig,
		},
		{
			Name: "set from CLI and environment variables",
			CLIFlagConfig: config.Config{
				Auth: config.AuthConfig{
					AuthMethod: "foo",
					ClientID:   "bar",
				},
			},
			// testing to see if feature flags set on the CLI take precedence over environment variables
			EnvironmentVariables: map[string]string{
				"AKS_ENGINE_AUTH_METHOD":   "Fear is the path to the dark side. Fear leads to anger. Anger leads to hate. Hate leads to suffering.",
				"AKS_ENGINE_CLIENT_ID":     "Truly wonderful, the mind of a child is.",
				"AKS_ENGINE_CLIENT_SECRET": "car",
				"AKS_ENGINE_LANGUAGE":      "star",
			},
			ExpectedConfig: expectedConfig,
		},
		{
			Name: "feature flags set from CLI and settings.json",
			CLIFlagConfig: config.Config{
				Auth: config.AuthConfig{
					AuthMethod: "foo",
					ClientID:   "bar",
				},
			},
			// testing to see if feature flags set on the CLI take precedence over settings.json
			SettingsFileContent: `{"auth":{"auth-method":"hello","client-id":"world","client-secret":"car","language":"star"}}`,
			ExpectedConfig:      expectedConfig,
		},
		{
			Name:                 "feature flags set from environment variables and settings.json",
			EnvironmentVariables: envVars,
			// same here; envvars have higher precedence than settings.json
			SettingsFileContent: `{"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"}}`,
			ExpectedConfig:      expectedConfig,
		},
		{
			Name: "feature flags set from CLI, environment variables and settings.json",
			CLIFlagConfig: config.Config{
				Auth: config.AuthConfig{
					AuthMethod:   "foo",
					ClientID:     "bar",
					ClientSecret: "car",
					Language:     "star",
				},
			},
			EnvironmentVariables: map[string]string{
				"AKS_ENGINE_AUTH_METHOD":   "Fear is the path to the dark side. Fear leads to anger. Anger leads to hate. Hate leads to suffering.",
				"AKS_ENGINE_CLIENT_ID":     "Truly wonderful, the mind of a child is.",
				"AKS_ENGINE_CLIENT_SECRET": "When nine hundred years old you reach, look as good you will not.",
				"AKS_ENGINE_LANGUAGE":      "A Jedi uses the Force for knowledge and defense, never for attack.",
			},
			SettingsFileContent: `{"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"}}`,
			ExpectedConfig:      expectedConfig,
		},
		{
			Name: "feature flags set from settings.json, profile name set from CLI",
			CLIFlagConfig: config.Config{
				CLIConfig: config.CLIConfig{
					Profile: "production",
				},
			},
			SettingsFileContent: `{"profiles":{"production":{"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"}}},"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"}}`,
			ExpectedConfig:      expectedConfigWithProfile,
		},
		{
			Name: "feature flags set from settings.json, profile name set from environment variables",
			EnvironmentVariables: map[string]string{
				"AKS_ENGINE_PROFILE": "production",
			},
			SettingsFileContent: `{"profiles":{"production":{"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"}}},"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"}}`,
			ExpectedConfig:      expectedConfigWithProfile,
		},
		{
			Name:                "feature flags set from settings.json, profile name set from settings.json",
			SettingsFileContent: `{"profiles":{"production":{"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"}}},"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"},"aks-engine":{"profile":"production"}}`,
			ExpectedConfig:      expectedConfigWithProfile,
		},
		{
			Name: "feature flags set from settings.json, profile name set from CLI and environment variables",
			CLIFlagConfig: config.Config{
				CLIConfig: config.CLIConfig{
					Profile: "production",
				},
			},
			EnvironmentVariables: map[string]string{
				"AKS_ENGINE_PROFILE": "development",
			},
			SettingsFileContent: `{"profiles":{"production":{"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"}}},"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"}}`,
			ExpectedConfig:      expectedConfigWithProfile,
		},
		{
			Name: "feature flags set from settings.json, profile name set from CLI and settings.json",
			CLIFlagConfig: config.Config{
				CLIConfig: config.CLIConfig{
					Profile: "production",
				},
			},
			SettingsFileContent: `{"profiles":{"production":{"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"}}},"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"},"aks-engine":{"profile":"development"}}`,
			ExpectedConfig:      expectedConfigWithProfile,
		},
		{
			Name: "feature flags set from settings.json, profile name set from CLI, environment variables and settings.json",
			CLIFlagConfig: config.Config{
				CLIConfig: config.CLIConfig{
					Profile: "production",
				},
			},
			EnvironmentVariables: map[string]string{
				"AKS_ENGINE_PROFILE": "development",
			},
			SettingsFileContent: `{"profiles":{"production":{"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"}}},"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"},"aks-engine":{"profile":"staging"}}`,
			ExpectedConfig:      expectedConfigWithProfile,
		},
		{
			Name:                  "bad settings.json",
			SettingsFileContent:   "{]",
			FromReaderShouldError: true,
			ExpectedErrorMessage:  "invalid character ']' looking for beginning of object key string",
		},
		{
			Name:                "bad profile name, defaulting to settings in the 'default' profile in settings.json",
			SettingsFileContent: `{"profiles":{"staging":{"auth":{"auth-method":"yer","client-id":"a","client-secret":"wizard","language":"Harry"}}},"auth":{"auth-method":"foo","client-id":"bar","client-secret":"car","language":"star"},"aks-engine":{"profile":"production"}}`,
			ExpectedConfig:      expectedConfigWithProfile,
		},
	}

	// clear the environment to ensure a clean test
	for _, envVar := range knownEnvVars {
		os.Unsetenv(envVar)
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			for k, v := range tc.EnvironmentVariables {
				os.Setenv(k, v)
			}
			currentConfig = defaultConfigValues
			mergo.Merge(&tc.CLIFlagConfig, currentConfig)
			currentConfig = tc.CLIFlagConfig
			s, err := config.FromReader(bytes.NewBufferString(tc.SettingsFileContent))
			if tc.FromReaderShouldError {
				if err == nil {
					t.Error("expected an error to occur, got nil")
				}
				if err.Error() != tc.ExpectedErrorMessage {
					t.Errorf("expected error message '%s', got '%s'", tc.ExpectedErrorMessage, err.Error())
				}
				// we expected an error when parsing the JSON object, so we're gonna cut short here.
				return
			}
			err = loadSettings(s)
			if tc.LoadSettingsShouldError {
				if err == nil {
					t.Error("expected an error to occur, got nil")
				} else if err.Error() != tc.ExpectedErrorMessage {
					t.Error(err)
				}
				// we were just checking if an error would occur, so we can skip checking what was in the settings
				return
			} else if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tc.ExpectedConfig, currentConfig) {
				t.Errorf("expected config and current config differ.\nExpected: %v\n Actual: %v", tc.ExpectedConfig, currentConfig)
			}
			for k := range tc.EnvironmentVariables {
				os.Unsetenv(k)
			}
		})
	}
}
