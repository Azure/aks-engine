// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/cli/config"
	"github.com/bacongobbler/symdiff"
	xdg "github.com/casimir/xdg-go"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

const (
	envVarPrefix         = "AKS_ENGINE"
	rootName             = "aks-engine"
	rootShortDescription = "AKS Engine deploys and manages Kubernetes clusters in Azure"
	rootLongDescription  = "AKS Engine deploys and manages Kubernetes clusters in Azure"
)

var (
	configHome       = filepath.Join(xdg.ConfigHome(), "aks-engine")
	settingsFileName = "settings.json"
	// holds the configuration provided by the user.
	//
	// Priority: feature flags > environment variables > settings.json > defaultConfigValues
	currentConfig       = config.Config{}
	defaultConfigValues = config.Config{
		Auth: config.AuthConfig{
			AzureEnvironment: "AzurePublicCloud",
			AuthMethod:       "client_secret",
			Language:         "en-us",
		},
		CLIConfig: config.CLIConfig{
			Generate: config.GenerateConfig{
				PrettyPrint: true,
			},
			Upgrade: config.UpgradeConfig{
				VMTimeout: -1,
			},
			Version: config.VersionConfig{
				OutputFormat: "human",
			},
		},
	}

	// envMap maps flag names to envvars
	envMap = map[string]string{
		"azure-env":          fmt.Sprintf("%s_AZURE_ENVIRONMENT", envVarPrefix),
		"auth-method":        fmt.Sprintf("%s_AUTH_METHOD", envVarPrefix),
		"subscription-id":    fmt.Sprintf("%s_SUBSCRIPTION_ID", envVarPrefix),
		"client-id":          fmt.Sprintf("%s_CLIENT_ID", envVarPrefix),
		"client-secret":      fmt.Sprintf("%s_CLIENT_SECRET", envVarPrefix),
		"certificate-path":   fmt.Sprintf("%s_CERTIFICATE_PATH", envVarPrefix),
		"private-key-path":   fmt.Sprintf("%s_PRIVATE_KEY_PATH", envVarPrefix),
		"language":           fmt.Sprintf("%s_LANGUAGE", envVarPrefix),
		"debug":              fmt.Sprintf("%s_DEBUG", envVarPrefix),
		"profile":            fmt.Sprintf("%s_PROFILE", envVarPrefix),
		"show-default-model": fmt.Sprintf("%s_SHOW_DEFAULT_MODEL", envVarPrefix),
	}
)

// NewRootCmd returns the root command for AKS Engine.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   rootName,
		Short: rootShortDescription,
		Long:  rootLongDescription,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			settingsFilepath := filepath.Join(configHome, settingsFileName)
			_, err := os.Stat(settingsFilepath)
			var r io.Reader
			if err != nil {
				if !os.IsNotExist(err) {
					return fmt.Errorf("could not read settings from %s: %v", settingsFilepath, err)
				}
				// read in an empty settings file so we can still load settings from environment variables
				// FIXME: refactor this such that we can read settings from environment variables without
				// needing to provide an empty buffer
				r = bytes.NewBufferString("{}")
			} else {
				f, err := os.Open(settingsFilepath)
				if err != nil {
					return err
				}
				defer f.Close()
				r = f
			}
			s, err := config.FromReader(r)
			// skip if the settings could not be read due to it being an empty file
			if err != nil && err != io.EOF {
				return fmt.Errorf("could not read from JSON stream: %v", err)
			}
			if err := loadSettings(s); err != nil {
				return fmt.Errorf("could not load settings: %v", err)
			}
			if currentConfig.CLIConfig.Debug {
				log.SetLevel(log.DebugLevel)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if currentConfig.CLIConfig.ShowDefaultModel {
				return writeDefaultModel(cmd.OutOrStdout())
			}
			return cmd.Usage()
		},
	}

	p := rootCmd.PersistentFlags()
	p.BoolVar(&currentConfig.CLIConfig.Debug, "debug", defaultConfigValues.CLIConfig.Debug, "enable verbose debug logs")
	p.StringVarP(&currentConfig.CLIConfig.Profile, "profile", "p", defaultConfigValues.CLIConfig.Profile, "The name of the settings profile name to use")

	f := rootCmd.Flags()
	f.BoolVar(&currentConfig.CLIConfig.ShowDefaultModel, "show-default-model", defaultConfigValues.CLIConfig.ShowDefaultModel, "Dump the default API model to stdout")

	rootCmd.AddCommand(newConfigureCmd())
	rootCmd.AddCommand(getCompletionCmd(rootCmd))
	rootCmd.AddCommand(newDeployCmd())
	rootCmd.AddCommand(newGenerateCmd())
	rootCmd.AddCommand(newOrchestratorsCmd())
	rootCmd.AddCommand(newScaleCmd())
	rootCmd.AddCommand(newUpgradeCmd())
	rootCmd.AddCommand(newVersionCmd())

	return rootCmd
}

func writeDefaultModel(out io.Writer) error {
	meta, p := api.LoadDefaultContainerServiceProperties()
	type withMeta struct {
		APIVersion string            `json:"apiVersion"`
		Properties *vlabs.Properties `json:"properties"`
	}

	b, err := json.MarshalIndent(withMeta{APIVersion: meta.APIVersion, Properties: p}, "", "\t")
	if err != nil {
		return errors.Wrap(err, "error encoding model to json")
	}
	b = append(b, '\n')
	if _, err := out.Write(b); err != nil {
		return errors.Wrap(err, "error writing output")
	}
	return nil
}

type authProvider interface {
	getAuthArgs() *config.AuthConfig
	getClient() (armhelpers.AKSEngineClient, error)
}

func addAuthFlags(auth *config.AuthConfig, f *flag.FlagSet) {
	f.StringVar(&auth.AzureEnvironment, "azure-env", defaultConfigValues.Auth.AzureEnvironment, "the target Azure cloud")
	f.StringVarP(&auth.SubscriptionID, "subscription-id", "s", defaultConfigValues.Auth.SubscriptionID, "azure subscription id (required)")
	f.StringVar(&auth.AuthMethod, "auth-method", defaultConfigValues.Auth.AuthMethod, "auth method (default:`client_secret`, `cli`, `client_certificate`, `device`)")
	f.StringVar(&auth.ClientID, "client-id", defaultConfigValues.Auth.ClientID, "client id (used with --auth-method=[client_secret|client_certificate])")
	f.StringVar(&auth.ClientSecret, "client-secret", defaultConfigValues.Auth.ClientSecret, "client secret (used with --auth-mode=client_secret)")
	f.StringVar(&auth.CertificatePath, "certificate-path", defaultConfigValues.Auth.CertificatePath, "path to client certificate (used with --auth-method=client_certificate)")
	f.StringVar(&auth.PrivateKeyPath, "private-key-path", defaultConfigValues.Auth.PrivateKeyPath, "path to private key (used with --auth-method=client_certificate)")
	f.StringVar(&auth.Language, "language", defaultConfigValues.Auth.Language, "language to return error messages in")
}

func getCompletionCmd(root *cobra.Command) *cobra.Command {
	var completionCmd = &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
		Long: `To load completion run

	source <(aks-engine completion)

	To configure your bash shell to load completions for each session, add this to your bashrc

	# ~/.bashrc or ~/.profile
	source <(aks-engine completion)
	`,
		Run: func(cmd *cobra.Command, args []string) {
			root.GenBashCompletion(os.Stdout)
		},
	}
	return completionCmd
}

func loadSettings(s *config.Settings) error {
	envConfig := config.Config{
		Auth: config.AuthConfig{
			AzureEnvironment: os.Getenv(envMap["azure-env"]),
			AuthMethod:       os.Getenv(envMap["auth-method"]),
			SubscriptionID:   os.Getenv(envMap["subscription-id"]),
			ClientID:         os.Getenv(envMap["client-id"]),
			ClientSecret:     os.Getenv(envMap["client-secret"]),
			CertificatePath:  os.Getenv(envMap["certificate-path"]),
			PrivateKeyPath:   os.Getenv(envMap["private-key-path"]),
			Language:         os.Getenv(envMap["language"]),
		},
		CLIConfig: config.CLIConfig{
			Debug:            os.Getenv(envMap["debug"]) != "",
			Profile:          os.Getenv(envMap["profile"]),
			ShowDefaultModel: os.Getenv(envMap["show-default-model"]) != "",
		},
	}
	// first, we need to determine the profile we want to load from settings.json
	profile := s.CLIConfig.Profile
	if envConfig.CLIConfig.Profile != "" {
		profile = envConfig.CLIConfig.Profile
	}
	if currentConfig.CLIConfig.Profile != "" {
		profile = currentConfig.CLIConfig.Profile
	}
	// merge in flags from the environment as well as from settings.json
	if err := mergo.Merge(&envConfig, s.UsingProfile(profile)); err != nil {
		return fmt.Errorf("could not merge settings from the settings file into the environment: %v", err)
	}
	// now we need to merge in flags from outside the CLI into what the CLI parsed from the command line
	//
	// NOTE(bacongobbler): this is a little tricky; mergo only merges in values from src (external flags) into
	// dest (flags set using the CLI) if dest's corresponding value is empty. Because certain flags
	// in the CLI are set to a non-empty value by default (such as the `--vm-timeout` flag in `aks-engine upgrade`
	// and the `--output` flag in `aks-engine version`), we need to handle those particular values in a clunky order:
	// merge from dest to src (to fill in any empty values from settings.json/envvars with the defaults), convert
	// values in dest that are their original default values to empty values, then merge src back into dest.
	if err := mergo.Merge(&envConfig, currentConfig); err != nil {
		return fmt.Errorf("could not merge settings from the CLI into the environment: %v", err)
	}
	if err := symdiff.Diff(&currentConfig, defaultConfigValues); err != nil {
		return fmt.Errorf("could not find the symmetric differences between what was set in the CLI and what was set in settings.json and the environment: %v", err)
	}
	// now that we've found out the feature flags that have been set, those take precedence over what came from external sources
	if err := mergo.Merge(&currentConfig, envConfig); err != nil {
		return fmt.Errorf("could not merge settings from the environment into the CLI: %v", err)
	}
	return nil
}
