// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package config

import (
	"encoding/json"
	"io"
	"path/filepath"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	ini "gopkg.in/ini.v1"
)

// DeployConfig defines the configuration for the `aks-engine deploy` command.
type DeployConfig struct {
	APIModel          string   `json:"api-model"`
	AutoSuffix        bool     `json:"auto-suffix"`
	CACertificatePath string   `json:"ca-certificate-path"`
	CAPrivateKeyPath  string   `json:"ca-private-key-path"`
	DNSPrefix         string   `json:"dns-prefix"`
	ForceOverwrite    bool     `json:"force-overwrite"`
	Location          string   `json:"location"`
	OutputDirectory   string   `json:"output-directory"`
	ResourceGroup     string   `json:"resource-group"`
	Set               []string `json:"set"`
	ParametersOnly    bool     `json:"parameters-only"`
}

// GenerateConfig defines the configuration for the `aks-engine generate` command.
type GenerateConfig struct {
	APIModel          string   `json:"api-model"`
	CACertificatePath string   `json:"ca-certificate-path"`
	CAPrivateKeyPath  string   `json:"ca-private-key-path"`
	PrettyPrint       bool     `json:"pretty-print"`
	OutputDirectory   string   `json:"output-directory"`
	ParametersOnly    bool     `json:"parameters-only"`
	Set               []string `json:"set"`
}

// ScaleConfig defines the configuration for the `aks-engine scale` command.
type ScaleConfig struct {
	DeploymentDir string `json:"deployment-dir"`
	Location      string `json:"location"`
	MasterFQDN    string `json:"master-fqdn"`
	NewNodeCount  int    `json:"new-node-count"`
	NodePool      string `json:"node-pool"`
	ResourceGroup string `json:"resource-group"`
}

// UpgradeConfig defines the configuration for the `aks-engine upgrade` command.
type UpgradeConfig struct {
	DeploymentDir  string `json:"deployment-dir"`
	Location       string `json:"location"`
	ResourceGroup  string `json:"resource-group"`
	UpgradeVersion string `json:"upgrade-version"`
	VMTimeout      int    `json:"vm-timeout"`
}

// VersionConfig defines the configuration for the `aks-engine version` command.
type VersionConfig struct {
	OutputFormat string `json:"output"`
}

// AuthConfig defines the configuration that can be read in for authenticating against Microsoft Azure.
type AuthConfig struct {
	AzureEnvironment string `json:"azure-env"`
	AuthMethod       string `json:"auth-method"`
	SubscriptionID   string `json:"subscription-id"`
	ClientID         string `json:"client-id"`
	ClientSecret     string `json:"client-secret"`
	CertificatePath  string `json:"certificate-path"`
	PrivateKeyPath   string `json:"private-key-path"`
	Language         string `json:"language"`
}

// CLIConfig defines the configuration that can be read in for the AKS Engine CLI.
type CLIConfig struct {
	Debug            bool           `json:"debug"`
	Profile          string         `json:"profile"`
	ShowDefaultModel bool           `json:"show-default-model"`
	Deploy           DeployConfig   `json:"deploy"`
	Generate         GenerateConfig `json:"generate"`
	Scale            ScaleConfig    `json:"scale"`
	Upgrade          UpgradeConfig  `json:"upgrade"`
	Version          VersionConfig  `json:"version"`
}

// Config defines the configuration that is read in for the AKS Engine CLI.
type Config struct {
	Auth      AuthConfig `json:"auth"`
	CLIConfig CLIConfig  `json:"aks-engine"`
}

// ProfileMap maps a name to some configuration which is provided via the --profile flag.
type ProfileMap map[string]Config

// Settings represents a settings.json
type Settings struct {
	Profiles  ProfileMap `json:"profiles"`
	Auth      AuthConfig `json:"auth"`
	CLIConfig CLIConfig  `json:"aks-engine"`
}

// New creates a new settings file.
func New() *Settings {
	return &Settings{
		Profiles: ProfileMap{},
	}
}

// FromReader reads settings from a JSON stream and returns the resulting Settings object.
func FromReader(r io.Reader) (*Settings, error) {
	s := New()
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(s); err != nil {
		return s, err
	}
	return s, nil
}

// UsingProfile returns a Config object based on the requested profile name, merging in from the default settings.
//
// If the profile name does not exist in the settings, the default configuration
// from the settings file is returned.
func (s *Settings) UsingProfile(name string) *Config {
	c := &Config{
		Auth:      s.Auth,
		CLIConfig: s.CLIConfig,
	}

	if profile, ok := s.Profiles[name]; ok {
		mergo.Merge(&profile, c)
		c = &profile
	}

	return c
}

// Validate checks to make sure that the auth config is well-formed.
func (a *AuthConfig) Validate() error {
	clientID, _ := uuid.FromString(a.ClientID)
	subID, _ := uuid.FromString(a.SubscriptionID)

	if a.AuthMethod == "client_secret" {
		if clientID.String() == "00000000-0000-0000-0000-000000000000" || a.ClientSecret == "" {
			return errors.New(`--client-id and --client-secret must be specified when --auth-method="client_secret"`)
		}
		// try parse the UUID
	} else if a.AuthMethod == "client_certificate" {
		if clientID.String() == "00000000-0000-0000-0000-000000000000" || a.CertificatePath == "" || a.PrivateKeyPath == "" {
			return errors.New(`--client-id and --certificate-path, and --private-key-path must be specified when --auth-method="client_certificate"`)
		}
	}

	if subID.String() == "00000000-0000-0000-0000-000000000000" {
		s, err := getSubFromAzDir(filepath.Join(helpers.GetHomeDir(), ".azure"))
		if err != nil || s.String() == "00000000-0000-0000-0000-000000000000" {
			return errors.New("--subscription-id is required (and must be a valid UUID)")
		}
		log.Infoln("No subscription provided. Using selected subscription from azure CLI:", subID.String())
		subID = s
	}

	_, err := azure.EnvironmentFromName(a.AzureEnvironment)
	if err != nil {
		return errors.New("failed to parse --azure-env as a valid target Azure cloud environment")
	}
	return nil
}

// NewClient fetches an AKS Engine client with the given parameters
func (a *AuthConfig) NewClient() (armhelpers.AKSEngineClient, error) {
	var client *armhelpers.AzureClient
	env, err := azure.EnvironmentFromName(a.AzureEnvironment)
	if err != nil {
		return nil, err
	}
	switch a.AuthMethod {
	case "cli":
		client, err = armhelpers.NewAzureClientWithCLI(env, a.SubscriptionID)
	case "device":
		client, err = armhelpers.NewAzureClientWithDeviceAuth(env, a.SubscriptionID)
	case "client_secret":
		client, err = armhelpers.NewAzureClientWithClientSecret(env, a.SubscriptionID, a.ClientID, a.ClientSecret)
	case "client_certificate":
		client, err = armhelpers.NewAzureClientWithClientCertificateFile(env, a.SubscriptionID, a.ClientID, a.CertificatePath, a.PrivateKeyPath)
	default:
		return nil, errors.Errorf("--auth-method: ERROR: method unsupported. method=%q", a.AuthMethod)
	}
	if err != nil {
		return nil, err
	}
	err = client.EnsureProvidersRegistered(a.SubscriptionID)
	if err != nil {
		return nil, err
	}
	client.AddAcceptLanguages([]string{a.Language})
	return client, nil
}

func getSubFromAzDir(root string) (uuid.UUID, error) {
	subConfig, err := ini.Load(filepath.Join(root, "clouds.config"))
	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, "error decoding cloud subscription config")
	}

	cloudConfig, err := ini.Load(filepath.Join(root, "config"))
	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, "error decoding cloud config")
	}

	cloud := getSelectedCloudFromAzConfig(cloudConfig)
	return getCloudSubFromAzConfig(cloud, subConfig)
}

func getSelectedCloudFromAzConfig(f *ini.File) string {
	selectedCloud := "AzureCloud"
	if cloud, err := f.GetSection("cloud"); err == nil {
		if name, err := cloud.GetKey("name"); err == nil {
			if s := name.String(); s != "" {
				selectedCloud = s
			}
		}
	}
	return selectedCloud
}

func getCloudSubFromAzConfig(cloud string, f *ini.File) (uuid.UUID, error) {
	cfg, err := f.GetSection(cloud)
	if err != nil {
		return uuid.UUID{}, errors.New("could not find user defined subscription id")
	}
	sub, err := cfg.GetKey("subscription")
	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, "error reading subscription id from cloud config")
	}
	return uuid.FromString(sub.String())
}
