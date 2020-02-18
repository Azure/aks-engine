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
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/armhelpers/azurestack"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	ini "gopkg.in/ini.v1"
)

const (
	rootName             = "aks-engine"
	rootShortDescription = "AKS Engine deploys and manages Kubernetes clusters in Azure"
	rootLongDescription  = "AKS Engine deploys and manages Kubernetes clusters in Azure"
)

var (
	debug            bool
	dumpDefaultModel bool
)

// NewRootCmd returns the root command for AKS Engine.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   rootName,
		Short: rootShortDescription,
		Long:  rootLongDescription,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				log.SetLevel(log.DebugLevel)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if dumpDefaultModel {
				return writeDefaultModel(cmd.OutOrStdout())
			}
			return cmd.Usage()
		},
	}

	p := rootCmd.PersistentFlags()
	p.BoolVar(&debug, "debug", false, "enable verbose debug logs")

	f := rootCmd.Flags()
	f.BoolVar(&dumpDefaultModel, "show-default-model", false, "Dump the default API model to stdout")

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newGenerateCmd())
	rootCmd.AddCommand(newDeployCmd())
	rootCmd.AddCommand(newGetVersionsCmd())
	rootCmd.AddCommand(newOrchestratorsCmd())
	rootCmd.AddCommand(newUpgradeCmd())
	rootCmd.AddCommand(newScaleCmd())
	rootCmd.AddCommand(newRotateCertsCmd())
	rootCmd.AddCommand(newAddPoolCmd())
	rootCmd.AddCommand(newGetLocationsCmd())
	rootCmd.AddCommand(getCompletionCmd(rootCmd))

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
	getAuthArgs() *authArgs
	getClient() (armhelpers.AKSEngineClient, error)
}

type authArgs struct {
	RawAzureEnvironment string
	rawSubscriptionID   string
	SubscriptionID      uuid.UUID
	AuthMethod          string
	rawClientID         string

	ClientID        uuid.UUID
	ClientSecret    string
	CertificatePath string
	PrivateKeyPath  string
	IdentitySystem  string
	language        string
}

func addAuthFlags(authArgs *authArgs, f *flag.FlagSet) {
	f.StringVar(&authArgs.RawAzureEnvironment, "azure-env", "AzurePublicCloud", "the target Azure cloud")
	f.StringVarP(&authArgs.rawSubscriptionID, "subscription-id", "s", "", "azure subscription id (required)")
	f.StringVar(&authArgs.AuthMethod, "auth-method", "client_secret", "auth method (default:`client_secret`, `cli`, `client_certificate`, `device`)")
	f.StringVar(&authArgs.rawClientID, "client-id", "", "client id (used with --auth-method=[client_secret|client_certificate])")
	f.StringVar(&authArgs.ClientSecret, "client-secret", "", "client secret (used with --auth-method=client_secret)")
	f.StringVar(&authArgs.CertificatePath, "certificate-path", "", "path to client certificate (used with --auth-method=client_certificate)")
	f.StringVar(&authArgs.PrivateKeyPath, "private-key-path", "", "path to private key (used with --auth-method=client_certificate)")
	f.StringVar(&authArgs.IdentitySystem, "identity-system", "azure_ad", "identity system (default:`azure_ad`, `adfs`)")
	f.StringVar(&authArgs.language, "language", "en-us", "language to return error messages in")
}

//this allows the authArgs to be stubbed behind the authProvider interface, and be its own provider when not in tests.
func (authArgs *authArgs) getAuthArgs() *authArgs {
	return authArgs
}

func (authArgs *authArgs) isAzureStackCloud() bool {
	return strings.EqualFold(authArgs.RawAzureEnvironment, api.AzureStackCloud)
}

func (authArgs *authArgs) validateAuthArgs() error {
	authArgs.ClientID, _ = uuid.Parse(authArgs.rawClientID)
	authArgs.SubscriptionID, _ = uuid.Parse(authArgs.rawSubscriptionID)

	if authArgs.AuthMethod == "client_secret" {
		if authArgs.ClientID.String() == "00000000-0000-0000-0000-000000000000" || authArgs.ClientSecret == "" {
			return errors.New(`--client-id and --client-secret must be specified when --auth-method="client_secret"`)
		}
		// try parse the UUID
	} else if authArgs.AuthMethod == "client_certificate" {
		if authArgs.ClientID.String() == "00000000-0000-0000-0000-000000000000" || authArgs.CertificatePath == "" || authArgs.PrivateKeyPath == "" {
			return errors.New(`--client-id and --certificate-path, and --private-key-path must be specified when --auth-method="client_certificate"`)
		}
	}

	if authArgs.SubscriptionID.String() == "00000000-0000-0000-0000-000000000000" {
		subID, err := getSubFromAzDir(filepath.Join(helpers.GetHomeDir(), ".azure"))
		if err != nil || subID.String() == "00000000-0000-0000-0000-000000000000" {
			return errors.New("--subscription-id is required (and must be a valid UUID)")
		}
		log.Infoln("No subscription provided, using selected subscription from azure CLI:", subID.String())
		authArgs.SubscriptionID = subID
	}

	_, err := azure.EnvironmentFromName(authArgs.RawAzureEnvironment)
	if err != nil {
		return errors.New("failed to parse --azure-env as a valid target Azure cloud environment")
	}
	return nil
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
	return uuid.Parse(sub.String())
}

func (authArgs *authArgs) getClient() (armhelpers.AKSEngineClient, error) {
	if authArgs.isAzureStackCloud() {
		return authArgs.getAzureStackClient()
	}
	return authArgs.getAzureClient()
}

func (authArgs *authArgs) getAzureClient() (armhelpers.AKSEngineClient, error) {
	var client *armhelpers.AzureClient
	env, err := azure.EnvironmentFromName(authArgs.RawAzureEnvironment)
	if err != nil {
		return nil, err
	}
	switch authArgs.AuthMethod {
	case "cli":
		client, err = armhelpers.NewAzureClientWithCLI(env, authArgs.SubscriptionID.String())
	case "device":
		client, err = armhelpers.NewAzureClientWithDeviceAuth(env, authArgs.SubscriptionID.String())
	case "client_secret":
		client, err = armhelpers.NewAzureClientWithClientSecret(env, authArgs.SubscriptionID.String(), authArgs.ClientID.String(), authArgs.ClientSecret)
	case "client_certificate":
		client, err = armhelpers.NewAzureClientWithClientCertificateFile(env, authArgs.SubscriptionID.String(), authArgs.ClientID.String(), authArgs.CertificatePath, authArgs.PrivateKeyPath)
	default:
		return nil, errors.Errorf("--auth-method: ERROR: method unsupported. method=%q", authArgs.AuthMethod)
	}
	if err != nil {
		return nil, err
	}
	err = client.EnsureProvidersRegistered(authArgs.SubscriptionID.String())
	if err != nil {
		return nil, err
	}
	client.AddAcceptLanguages([]string{authArgs.language})
	return client, nil
}

func (authArgs *authArgs) getAzureStackClient() (armhelpers.AKSEngineClient, error) {
	var client *azurestack.AzureClient
	env, err := azure.EnvironmentFromName(authArgs.RawAzureEnvironment)
	if err != nil {
		return nil, err
	}
	switch authArgs.AuthMethod {
	case "client_secret":
		if authArgs.IdentitySystem == "azure_ad" {
			client, err = azurestack.NewAzureClientWithClientSecret(env, authArgs.SubscriptionID.String(), authArgs.ClientID.String(), authArgs.ClientSecret)
		} else if authArgs.IdentitySystem == "adfs" {
			// for ADFS environment, it is single tenant environment and the tenant id is aways adfs
			client, err = azurestack.NewAzureClientWithClientSecretExternalTenant(env, authArgs.SubscriptionID.String(), "adfs", authArgs.ClientID.String(), authArgs.ClientSecret)
		} else {
			return nil, errors.Errorf("--auth-method: ERROR: method unsupported. method=%q identitysystem=%q", authArgs.AuthMethod, authArgs.IdentitySystem)
		}
	case "client_certificate":
		if authArgs.IdentitySystem == "azure_ad" {
			client, err = azurestack.NewAzureClientWithClientCertificateFile(env, authArgs.SubscriptionID.String(), authArgs.ClientID.String(), authArgs.CertificatePath, authArgs.PrivateKeyPath)
			break
		} else if authArgs.IdentitySystem == "adfs" {
			// for ADFS environment, it is single tenant environment and the tenant id is aways adfs
			client, err = azurestack.NewAzureClientWithClientCertificateFileExternalTenant(env, authArgs.SubscriptionID.String(), "adfs", authArgs.ClientID.String(), authArgs.CertificatePath, authArgs.PrivateKeyPath)
			break
		}
		fallthrough
	default:
		return nil, errors.Errorf("--auth-method: ERROR: method unsupported. method=%q identitysystem=%q", authArgs.AuthMethod, authArgs.IdentitySystem)
	}
	if err != nil {
		return nil, err
	}
	err = client.EnsureProvidersRegistered(authArgs.SubscriptionID.String())
	if err != nil {
		return nil, err
	}
	client.AddAcceptLanguages([]string{authArgs.language})
	return client, nil
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return root.GenBashCompletion(os.Stdout)
		},
	}
	return completionCmd
}

func writeCustomCloudProfile(cs *api.ContainerService) error {

	tmpFile, err := ioutil.TempFile("", "azurestackcloud.json")
	tmpFileName := tmpFile.Name()
	if err != nil {
		return err
	}
	log.Infoln(fmt.Sprintf("Writing cloud profile to: %s", tmpFileName))

	// Build content for the file
	content, err := cs.Properties.GetCustomEnvironmentJSON(false)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(tmpFileName, []byte(content), os.ModeAppend); err != nil {
		return err
	}

	os.Setenv("AZURE_ENVIRONMENT_FILEPATH", tmpFileName)

	return nil
}
