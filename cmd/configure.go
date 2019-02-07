// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"

	"github.com/Azure/aks-engine/pkg/cli/config"
)

const (
	configureDesc = `
This command will prompt you for configuration values that will be written to a settings file.
This file will be read every time you invoke 'aks-engine', so it's useful for setting default
values you commonly use with 'aks-engine'.

Based on your operating system, the settings file will be created and read from the following
places:

Linux: $HOME/.config/aks-engine/settings.json
MacOS: $HOME/Library/Preferences/aks-engine/settings.json
Windows: %APPDATA%\aks-engine\settings.json

The base directory can also be changed by setting the 'XDG_CONFIG_HOME' environment variable.
If you set 'XDG_CONFIG_HOME' on Linux to '/etc', your settings file will be created in
'/etc/aks-engine/settings.json'.

You can configure a named profile using the '--profile' flag. For example,
'aks-engine configure --profile development' will set up the 'development' profile which can then
be used with 'aks-engine --profile development ...'.

If the settings file does not exist, the AKS Engine CLI will create it for you.

To keep an existing value, hit enter when prompted for the value.

When you are prompted for information, the current value will be displayed in '(parentheses)'.
`
	configureName             = "configure"
	configureShortDescription = "configure common settings"
	configureLongDescription  = "configure common settings for AKS Engine. Useful for setting common features such as the language"
)

type configureCmd struct{}

func newConfigureCmd() *cobra.Command {
	cc := configureCmd{}

	configureCmd := &cobra.Command{
		Use:   configureName,
		Short: configureShortDescription,
		Long:  configureLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cc.run()
		},
	}

	return configureCmd
}

func (cc *configureCmd) run() error {
	f, err := os.Open(settingsFilepath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		// create if missing
		if err := os.MkdirAll(filepath.Dir(settingsFilepath), 0755); err != nil {
			return err
		}
		f, err = os.Create(settingsFilepath)
		if err != nil {
			return err
		}
	}
	s, err := config.FromReader(f)
	if err != nil && err != io.EOF {
		// log a warning message if the settings could not be read; possibly due to it being an empty or malformed file
		// config.FromReader will return a non-nil settings object so it's safe to carry on as if we read from the file
		log.Warningf("could not read settings from %s: %v", settingsFilepath, err)
	}
	f.Close()
	// load current settings; we display these to the user
	if err := loadSettings(s); err != nil {
		return fmt.Errorf("could not read settings from %s: %v", settingsFilepath, err)
	}
	if err := prompt(&currentConfig); err != nil {
		return err
	}

	if currentConfig.CLIConfig.Profile != "" {
		s.Profiles[currentConfig.CLIConfig.Profile] = currentConfig
	} else {
		s.Auth = currentConfig.Auth
		s.CLIConfig = currentConfig.CLIConfig
	}
	// save current config to the settings file
	content, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return fmt.Errorf("could not marshal settings to JSON: %v", err)
	}

	return ioutil.WriteFile(settingsFilepath, content, 0644)
}

// prompt handles all of the prompts.
//
// Default values are read from the current settings. Values are then
// replaced on that object.
func prompt(conf *config.Config) error {
	qs := []*survey.Question{
		{
			Name: "debug",
			Prompt: &survey.Confirm{
				Message: "Enable verbose debug logging?",
				Default: conf.CLIConfig.Debug,
			},
		},
	}
	if err := survey.Ask(qs, &conf.CLIConfig); err != nil {
		return err
	}

	if err := promptAuth(&conf.Auth); err != nil {
		return err
	}

	var advancedDeploy bool
	survey.AskOne(&survey.Confirm{
		Message: "Setup advanced configuration for 'aks-engine deploy'?",
		Default: advancedDeploy,
	}, advancedDeploy, nil)
	if advancedDeploy {
		if err := promptDeploy(&conf.CLIConfig.Deploy); err != nil {
			return err
		}
	}

	var advancedGenerate bool
	survey.AskOne(&survey.Confirm{
		Message: "Setup advanced configuration for 'aks-engine generate'?",
		Default: advancedGenerate,
	}, advancedGenerate, nil)
	if advancedGenerate {
		if err := promptGenerate(&conf.CLIConfig.Generate); err != nil {
			return err
		}
	}

	var advancedScale bool
	survey.AskOne(&survey.Confirm{
		Message: "Setup advanced configuration for 'aks-engine scale'?",
		Default: advancedScale,
	}, advancedScale, nil)
	if advancedScale {
		if err := promptScale(&conf.CLIConfig.Scale); err != nil {
			return err
		}
	}

	var advancedUpgrade bool
	survey.AskOne(&survey.Confirm{
		Message: "Setup advanced configuration for 'aks-engine upgrade'?",
		Default: advancedUpgrade,
	}, advancedUpgrade, nil)
	if advancedUpgrade {
		if err := promptUpgrade(&conf.CLIConfig.Upgrade); err != nil {
			return err
		}
	}

	var advancedVersion bool
	survey.AskOne(&survey.Confirm{
		Message: "Setup advanced configuration for 'aks-engine version'?",
		Default: advancedVersion,
	}, advancedVersion, nil)
	if advancedVersion {
		if err := promptVersion(&conf.CLIConfig.Version); err != nil {
			return err
		}
	}

	return nil
}

func promptAuth(conf *config.AuthConfig) error {
	qs := []*survey.Question{
		{
			Name: "AuthMethod",
			Prompt: &survey.Select{
				Message: "What authentication method would you like to use?",
				Options: []string{"cli", "client_secret", "client_certificate", "device"},
				Default: conf.AuthMethod,
				Help:    "'cli' uses your credentials from the Azure CLI.\n'client_secret' uses a service principal's client ID and secret for authentication.\n'client_certificate' uses your CA certificate to generate a service principal token for authentication.\n'device' uses device authentication by having you complete a device authentication flow.",
			},
		},
		{
			Name: "AzureEnvironment",
			Prompt: &survey.Select{
				Message: "What Azure environment are you targeting?",
				Options: []string{"AzurePublicCloud", "AzureChinaCloud", "AzureGermanCloud", "AzureUSGovernmentCloud"},
				Default: conf.AzureEnvironment,
			},
		},
		{
			Name: "SubscriptionID",
			Prompt: &survey.Input{
				Message: "What subscription ID are you using?",
				Help:    "Use 'az account list -o table' to get a list of available subscription IDs",
				Default: conf.SubscriptionID,
			},
		},
		{
			Name: "Language",
			Prompt: &survey.Input{
				Message: "What language should Azure return error messages in?",
				Default: conf.Language,
			},
		},
	}
	if err := survey.Ask(qs, conf); err != nil {
		return err
	}

	switch conf.AuthMethod {
	case "client_secret":
		qs = []*survey.Question{
			{
				Name: "ClientID",
				Prompt: &survey.Input{
					Message: "What is your service principal's client ID?",
					Default: conf.ClientID,
				},
			},
			{
				Name: "ClientSecret",
				Prompt: &survey.Input{
					Message: "What is your service principal's client secret?",
					Default: conf.ClientSecret,
				},
			},
		}
	case "client_certificate":
		qs = []*survey.Question{
			{
				Name: "ClientID",
				Prompt: &survey.Input{
					Message: "What is your service principal's client ID?",
					Default: conf.ClientID,
				},
			},
			{
				Name: "ClientSecret",
				Prompt: &survey.Input{
					Message: "What is your service principal's client secret?",
					Default: conf.ClientSecret,
				},
			},
			{
				Name: "CertificatePath",
				Prompt: &survey.Input{
					Message: "What is the absolute filepath to your service principal's client certificate?",
					Default: conf.CertificatePath,
				},
			},
			{
				Name: "PrivateKeyPath",
				Prompt: &survey.Input{
					Message: "What is the absolute filepath to your certificate's private key?",
					Default: conf.PrivateKeyPath,
				},
			},
		}
	default:
		// no more questions to ask
		qs = []*survey.Question{}
	}

	return survey.Ask(qs, conf)
}

func promptDeploy(conf *config.DeployConfig) error {
	qs := []*survey.Question{
		{
			Name: "APIModel",
			Prompt: &survey.Input{
				Message: "What is the absolute filepath to the API model file?",
				Default: conf.APIModel,
			},
		},
		{
			Name: "AutoSuffix",
			Prompt: &survey.Confirm{
				Message: "Automatically append a compressed timestamp to the DNS prefix to ensure unique cluster name automatically?",
				Default: conf.AutoSuffix,
			},
		},
		{
			Name: "CACertificatePath",
			Prompt: &survey.Input{
				Message: "What is the absolute filepath to your CA certificate?",
				Default: conf.CACertificatePath,
			},
		},
		{
			Name: "PrivateKeyPath",
			Prompt: &survey.Input{
				Message: "What is the absolute filepath to your CA certificate's private key?",
				Default: conf.CAPrivateKeyPath,
			},
		},
		{
			Name: "DNSPrefix",
			Prompt: &survey.Input{
				Message: "What is the DNS prefix for your cluster?",
				Default: conf.DNSPrefix,
			},
		},
		{
			Name: "OutputDirectory",
			Prompt: &survey.Input{
				Message: "Output directory where 'aks-engine deploy' will write configuration?",
				Help:    "leave empty if you'd like 'aks-engine deploy' to choose a directory for you",
				Default: conf.OutputDirectory,
			},
		},
		{
			Name: "ForceOverwrite",
			Prompt: &survey.Confirm{
				Message: "Automatically overwrite existing files in the output directory on subsequent deploys?",
				Default: conf.ForceOverwrite,
			},
		},
		{
			Name: "Location",
			Prompt: &survey.Input{
				Message: "Azure region where your cluster will be deployed?",
				Default: conf.Location,
			},
		},
		{
			Name: "ParametersOnly",
			Prompt: &survey.Confirm{
				Message: "Only output parameters files?",
				Default: conf.ParametersOnly,
			},
		},
		{
			Name: "ResourceGroup",
			Prompt: &survey.Input{
				Message: "Azure Resource Group where your cluster will be deployed?",
				Default: conf.ResourceGroup,
			},
		},
	}
	return survey.Ask(qs, conf)
}

func promptGenerate(conf *config.GenerateConfig) error {
	qs := []*survey.Question{
		{
			Name: "APIModel",
			Prompt: &survey.Input{
				Message: "What is the absolute filepath to the API model file?",
				Default: conf.APIModel,
			},
		},
		{
			Name: "CACertificatePath",
			Prompt: &survey.Input{
				Message: "What is the absolute filepath to your CA certificate?",
				Default: conf.CACertificatePath,
			},
		},
		{
			Name: "PrivateKeyPath",
			Prompt: &survey.Input{
				Message: "What is the absolute filepath to your CA certificate's private key?",
				Default: conf.CAPrivateKeyPath,
			},
		},
		{
			Name: "PrettyPrint",
			Prompt: &survey.Confirm{
				Message: "Pretty print output?",
				Default: conf.PrettyPrint,
			},
		},
		{
			Name: "OutputDirectory",
			Prompt: &survey.Input{
				Message: "Output directory where 'aks-engine generate' will write configuration?",
				Help:    "leave empty if you'd like 'aks-engine generate' to choose a directory for you",
				Default: conf.OutputDirectory,
			},
		},
		{
			Name: "ParametersOnly",
			Prompt: &survey.Confirm{
				Message: "Only output parameters files?",
				Default: conf.ParametersOnly,
			},
		},
	}
	return survey.Ask(qs, conf)
}

func promptScale(conf *config.ScaleConfig) error {
	qs := []*survey.Question{
		{
			Name: "DeploymentDir",
			Prompt: &survey.Input{
				Message: "Where should 'aks-engine scale' write configuration?",
				Help:    "This should be the same location where 'aks-engine deploy' or 'aks-engine generate' will write configuration to",
				Default: conf.DeploymentDir,
			},
		},
		{
			Name: "Location",
			Prompt: &survey.Input{
				Message: "Azure region where your cluster will be deployed?",
				Default: conf.Location,
			},
		},
		{
			Name: "ResourceGroup",
			Prompt: &survey.Input{
				Message: "Azure Resource Group where your cluster will be deployed?",
				Default: conf.ResourceGroup,
			},
		},
		{
			Name: "MasterFQDN",
			Prompt: &survey.Input{
				Message: "FQDN for the master load balancer?",
				Help:    "Needed to scale down Kubernetes agent pools",
				Default: conf.MasterFQDN,
			},
		},
		{
			Name: "NodePool",
			Prompt: &survey.Input{
				Message: "Node pool to scale?",
				Default: conf.NodePool,
			},
		},
	}
	return survey.Ask(qs, conf)
}

func promptUpgrade(conf *config.UpgradeConfig) error {
	qs := []*survey.Question{
		{
			Name: "DeploymentDir",
			Prompt: &survey.Input{
				Message: "Where should 'aks-engine scale' write configuration?",
				Help:    "This should be the same location where 'aks-engine deploy' or 'aks-engine generate' will write configuration to",
				Default: conf.DeploymentDir,
			},
		},
		{
			Name: "Location",
			Prompt: &survey.Input{
				Message: "Azure region where your cluster will be deployed?",
				Default: conf.Location,
			},
		},
		{
			Name: "ResourceGroup",
			Prompt: &survey.Input{
				Message: "Azure Resource Group where your cluster will be deployed?",
				Default: conf.ResourceGroup,
			},
		},
	}
	return survey.Ask(qs, conf)
}

func promptVersion(conf *config.VersionConfig) error {
	qs := []*survey.Question{
		{
			Name: "OutputFormat",
			Prompt: &survey.Select{
				Message: "What is the preferred output format?",
				Options: []string{"human", "json"},
				Default: conf.OutputFormat,
			},
		},
	}
	return survey.Ask(qs, conf)
}
