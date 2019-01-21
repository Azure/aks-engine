// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/cli/config"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	deployName             = "deploy"
	deployShortDescription = "Deploy an Azure Resource Manager template"
	deployLongDescription  = "Deploy an Azure Resource Manager template, parameters file and other assets for a cluster"
)

type deployCmd struct {
	authProvider

	// derived
	containerService *api.ContainerService
	apiVersion       string
	locale           *gotext.Locale

	client        armhelpers.AKSEngineClient
	resourceGroup string
	random        *rand.Rand
	location      string
}

func newDeployCmd() *cobra.Command {
	dc := deployCmd{}

	deployCmd := &cobra.Command{
		Use:   deployName,
		Short: deployShortDescription,
		Long:  deployLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := dc.validateArgs(cmd, args); err != nil {
				log.Fatalf("error validating deployCmd: %s", err.Error())
			}
			if err := dc.mergeAPIModel(); err != nil {
				log.Fatalf("error merging API model in deployCmd: %s", err.Error())
			}
			if err := dc.loadAPIModel(cmd, args); err != nil {
				log.Fatalf("failed to load apimodel: %s", err.Error())
			}
			if _, _, err := dc.validateApimodel(); err != nil {
				log.Fatalf("Failed to validate the apimodel after populating values: %s", err.Error())
			}
			return dc.run()
		},
	}

	cfg := &currentConfig.CLIConfig.Deploy
	defaultCfg := &defaultConfigValues.CLIConfig.Deploy
	f := deployCmd.Flags()
	f.StringVarP(&cfg.APIModel, "api-model", "m", defaultCfg.APIModel, "path to the apimodel file")
	f.StringVarP(&cfg.DNSPrefix, "dns-prefix", "p", defaultCfg.DNSPrefix, "dns prefix (unique name for the cluster)")
	f.BoolVar(&cfg.AutoSuffix, "auto-suffix", defaultCfg.AutoSuffix, "automatically append a compressed timestamp to the dnsPrefix to ensure unique cluster name automatically")
	f.StringVarP(&cfg.OutputDirectory, "output-directory", "o", defaultCfg.OutputDirectory, "output directory (derived from FQDN if absent)")
	f.StringVar(&cfg.CACertificatePath, "ca-certificate-path", defaultCfg.CACertificatePath, "path to the CA certificate to use for Kubernetes PKI assets")
	f.StringVar(&cfg.CAPrivateKeyPath, "ca-private-key-path", defaultCfg.CAPrivateKeyPath, "path to the CA private key to use for Kubernetes PKI assets")
	f.StringVarP(&cfg.ResourceGroup, "resource-group", "g", defaultCfg.ResourceGroup, "resource group to deploy to (will use the DNS prefix from the apimodel if not specified)")
	f.StringVarP(&cfg.Location, "location", "l", defaultCfg.Location, "location to deploy to (required)")
	f.BoolVarP(&cfg.ForceOverwrite, "force-overwrite", "f", defaultCfg.ForceOverwrite, "automatically overwrite existing files in the output directory")
	f.BoolVar(&cfg.ParametersOnly, "parameters-only", defaultCfg.ParametersOnly, "only output parameters files")
	f.StringArrayVar(&cfg.Set, "set", defaultCfg.Set, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")

	addAuthFlags(dc.getAuthArgs(), f)

	return deployCmd
}

func (dc *deployCmd) getAuthArgs() *config.AuthConfig {
	return &currentConfig.Auth
}

func (dc *deployCmd) getClient() (armhelpers.AKSEngineClient, error) {
	return currentConfig.Auth.NewClient()
}

func (dc *deployCmd) validateArgs(cmd *cobra.Command, args []string) error {
	var err error

	dc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	if currentConfig.CLIConfig.Deploy.APIModel == "" {
		if len(args) == 1 {
			currentConfig.CLIConfig.Deploy.APIModel = args[0]
		} else if len(args) > 1 {
			cmd.Usage()
			return errors.New("too many arguments were provided to 'deploy'")
		}
	}

	if currentConfig.CLIConfig.Deploy.APIModel != "" {
		if _, err := os.Stat(currentConfig.CLIConfig.Deploy.APIModel); os.IsNotExist(err) {
			return errors.Errorf("specified api model does not exist (%s)", currentConfig.CLIConfig.Deploy.APIModel)
		}
	}

	if currentConfig.CLIConfig.Deploy.Location == "" {
		return errors.New("--location must be specified")
	}
	currentConfig.CLIConfig.Deploy.Location = helpers.NormalizeAzureRegion(currentConfig.CLIConfig.Deploy.Location)

	return nil
}

func (dc *deployCmd) mergeAPIModel() error {
	var err error

	if currentConfig.CLIConfig.Deploy.APIModel == "" {
		log.Infoln("no --api-model was specified, using default model")
		f, err := ioutil.TempFile("", fmt.Sprintf("%s-default-api-model_%s-%s_", filepath.Base(os.Args[0]), BuildSHA, GitTreeState))
		if err != nil {
			return errors.Wrap(err, "error creating temp file for default API model")
		}
		log.Infoln("default api model generated at", f.Name())

		defer f.Close()
		if err := writeDefaultModel(f); err != nil {
			return err
		}
		currentConfig.CLIConfig.Deploy.APIModel = f.Name()
	}

	// if --set flag has been used
	if len(currentConfig.CLIConfig.Deploy.Set) > 0 {
		m := make(map[string]transform.APIModelValue)
		transform.MapValues(m, currentConfig.CLIConfig.Deploy.Set)

		// overrides the api model and generates a new file
		currentConfig.CLIConfig.Deploy.APIModel, err = transform.MergeValuesWithAPIModel(currentConfig.CLIConfig.Deploy.APIModel, m)
		if err != nil {
			return errors.Wrapf(err, "error merging --set values with the api model: %s", currentConfig.CLIConfig.Deploy.APIModel)
		}

		log.Infoln(fmt.Sprintf("new api model file has been generated during merge: %s", currentConfig.CLIConfig.Deploy.APIModel))
	}

	return nil
}

func (dc *deployCmd) loadAPIModel(cmd *cobra.Command, args []string) error {
	var caCertificateBytes []byte
	var caKeyBytes []byte
	var err error

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: dc.locale,
		},
	}

	// do not validate when initially loading the apimodel, validation is done later after autofilling values
	dc.containerService, dc.apiVersion, err = apiloader.LoadContainerServiceFromFile(currentConfig.CLIConfig.Deploy.APIModel, false, false, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	if currentConfig.CLIConfig.Deploy.OutputDirectory == "" {
		if dc.containerService.Properties.MasterProfile != nil {
			currentConfig.CLIConfig.Deploy.OutputDirectory = path.Join("_output", dc.containerService.Properties.MasterProfile.DNSPrefix)
		} else {
			currentConfig.CLIConfig.Deploy.OutputDirectory = path.Join("_output", dc.containerService.Properties.HostedMasterProfile.DNSPrefix)
		}
	}

	// consume currentConfig.CLIConfig.Deploy.CACertificatePath and currentConfig.CLIConfig.Deploy.CAPrivateKeyPath
	if (currentConfig.CLIConfig.Deploy.CACertificatePath != "" && currentConfig.CLIConfig.Deploy.CAPrivateKeyPath == "") || (currentConfig.CLIConfig.Deploy.CACertificatePath == "" && currentConfig.CLIConfig.Deploy.CAPrivateKeyPath != "") {
		return errors.New("--ca-certificate-path and --ca-private-key-path must be specified together")
	}

	if currentConfig.CLIConfig.Deploy.CACertificatePath != "" {
		if caCertificateBytes, err = ioutil.ReadFile(currentConfig.CLIConfig.Deploy.CACertificatePath); err != nil {
			return errors.Wrap(err, "failed to read CA certificate file")
		}
		if caKeyBytes, err = ioutil.ReadFile(currentConfig.CLIConfig.Deploy.CAPrivateKeyPath); err != nil {
			return errors.Wrap(err, "failed to read CA private key file")
		}

		prop := dc.containerService.Properties
		if prop.CertificateProfile == nil {
			prop.CertificateProfile = &api.CertificateProfile{}
		}
		prop.CertificateProfile.CaCertificate = string(caCertificateBytes)
		prop.CertificateProfile.CaPrivateKey = string(caKeyBytes)
	}

	if dc.containerService.Location == "" {
		dc.containerService.Location = currentConfig.CLIConfig.Deploy.Location
	} else if dc.containerService.Location != currentConfig.CLIConfig.Deploy.Location {
		return errors.New("--location does not match api model location")
	}

	if err = dc.getAuthArgs().Validate(); err != nil {
		return err
	}

	dc.client, err = dc.authProvider.getClient()
	if err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	if err = autofillApimodel(dc); err != nil {
		return err
	}

	dc.random = rand.New(rand.NewSource(time.Now().UnixNano()))

	return nil
}

func autofillApimodel(dc *deployCmd) error {
	var err error

	if dc.containerService.Properties.LinuxProfile != nil {
		if dc.containerService.Properties.LinuxProfile.AdminUsername == "" {
			log.Warnf("apimodel: no linuxProfile.adminUsername was specified. Will use 'azureuser'.")
			dc.containerService.Properties.LinuxProfile.AdminUsername = "azureuser"
		}
	}

	if currentConfig.CLIConfig.Deploy.DNSPrefix != "" && dc.containerService.Properties.MasterProfile.DNSPrefix != "" {
		return errors.New("invalid configuration: the apimodel masterProfile.dnsPrefix and --dns-prefix were both specified")
	}
	if dc.containerService.Properties.MasterProfile.DNSPrefix == "" {
		if currentConfig.CLIConfig.Deploy.DNSPrefix == "" {
			return errors.New("apimodel: missing masterProfile.dnsPrefix and --dns-prefix was not specified")
		}
		log.Warnf("apimodel: missing masterProfile.dnsPrefix will use %q", currentConfig.CLIConfig.Deploy.DNSPrefix)
		dc.containerService.Properties.MasterProfile.DNSPrefix = currentConfig.CLIConfig.Deploy.DNSPrefix
	}

	if currentConfig.CLIConfig.Deploy.AutoSuffix {
		suffix := strconv.FormatInt(time.Now().Unix(), 16)
		dc.containerService.Properties.MasterProfile.DNSPrefix += "-" + suffix
	}

	if currentConfig.CLIConfig.Deploy.OutputDirectory == "" {
		currentConfig.CLIConfig.Deploy.OutputDirectory = path.Join("_output", dc.containerService.Properties.MasterProfile.DNSPrefix)
	}

	if _, err := os.Stat(currentConfig.CLIConfig.Deploy.OutputDirectory); !currentConfig.CLIConfig.Deploy.ForceOverwrite && err == nil {
		return errors.Errorf("Output directory already exists and forceOverwrite flag is not set: %s", currentConfig.CLIConfig.Deploy.OutputDirectory)
	}

	if currentConfig.CLIConfig.Deploy.ResourceGroup == "" {
		dnsPrefix := dc.containerService.Properties.MasterProfile.DNSPrefix
		log.Warnf("--resource-group was not specified. Using the DNS prefix from the apimodel as the resource group name: %s", dnsPrefix)
		currentConfig.CLIConfig.Deploy.ResourceGroup = dnsPrefix
		if currentConfig.CLIConfig.Deploy.Location == "" {
			return errors.New("--resource-group was not specified. --location must be specified in case the resource group needs creation")
		}
	}

	if dc.containerService.Properties.LinuxProfile != nil && (dc.containerService.Properties.LinuxProfile.SSH.PublicKeys == nil ||
		len(dc.containerService.Properties.LinuxProfile.SSH.PublicKeys) == 0 ||
		dc.containerService.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData == "") {
		translator := &i18n.Translator{
			Locale: dc.locale,
		}
		_, publicKey, err := helpers.CreateSaveSSH(dc.containerService.Properties.LinuxProfile.AdminUsername, currentConfig.CLIConfig.Deploy.OutputDirectory, translator)
		if err != nil {
			return errors.Wrap(err, "Failed to generate SSH Key")
		}

		dc.containerService.Properties.LinuxProfile.SSH.PublicKeys = []api.PublicKey{{KeyData: publicKey}}
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	_, err = dc.client.EnsureResourceGroup(ctx, currentConfig.CLIConfig.Deploy.ResourceGroup, currentConfig.CLIConfig.Deploy.Location, nil)
	if err != nil {
		return err
	}

	k8sConfig := dc.containerService.Properties.OrchestratorProfile.KubernetesConfig

	useManagedIdentity := k8sConfig != nil && k8sConfig.UseManagedIdentity

	if !useManagedIdentity {
		spp := dc.containerService.Properties.ServicePrincipalProfile
		if spp != nil && spp.ClientID == "" && spp.Secret == "" && spp.KeyvaultSecretRef == nil && (dc.getAuthArgs().ClientID == "" || dc.getAuthArgs().ClientID == "00000000-0000-0000-0000-000000000000") && dc.getAuthArgs().ClientSecret == "" {
			log.Warnln("apimodel: ServicePrincipalProfile was missing or empty, creating application...")

			// TODO: consider caching the creds here so they persist between subsequent runs of 'deploy'
			appName := dc.containerService.Properties.MasterProfile.DNSPrefix
			appURL := fmt.Sprintf("https://%s/", appName)
			var replyURLs *[]string
			var requiredResourceAccess *[]graphrbac.RequiredResourceAccess
			applicationResp, servicePrincipalObjectID, secret, err := dc.client.CreateApp(ctx, appName, appURL, replyURLs, requiredResourceAccess)
			if err != nil {
				return errors.Wrap(err, "apimodel invalid: ServicePrincipalProfile was empty, and we failed to create valid credentials")
			}
			applicationID := to.String(applicationResp.AppID)
			log.Warnf("created application with applicationID (%s) and servicePrincipalObjectID (%s).", applicationID, servicePrincipalObjectID)

			log.Warnln("apimodel: ServicePrincipalProfile was empty, assigning role to application...")

			err = dc.client.CreateRoleAssignmentSimple(ctx, currentConfig.CLIConfig.Deploy.ResourceGroup, servicePrincipalObjectID)
			if err != nil {
				return errors.Wrap(err, "apimodel: could not create or assign ServicePrincipal")

			}

			dc.containerService.Properties.ServicePrincipalProfile = &api.ServicePrincipalProfile{
				ClientID: applicationID,
				Secret:   secret,
				ObjectID: servicePrincipalObjectID,
			}
		} else if (dc.containerService.Properties.ServicePrincipalProfile == nil || ((dc.containerService.Properties.ServicePrincipalProfile.ClientID == "" || dc.containerService.Properties.ServicePrincipalProfile.ClientID == "00000000-0000-0000-0000-000000000000") && dc.containerService.Properties.ServicePrincipalProfile.Secret == "")) && dc.getAuthArgs().ClientID != "" && dc.getAuthArgs().ClientSecret != "" {
			dc.containerService.Properties.ServicePrincipalProfile = &api.ServicePrincipalProfile{
				ClientID: dc.getAuthArgs().ClientID,
				Secret:   dc.getAuthArgs().ClientSecret,
			}
		}
	}
	return nil
}

func (dc *deployCmd) validateApimodel() (*api.ContainerService, string, error) {
	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: dc.locale,
		},
	}

	p := dc.containerService.Properties
	if strings.ToLower(p.OrchestratorProfile.OrchestratorType) == "kubernetes" {
		if p.ServicePrincipalProfile == nil || (p.ServicePrincipalProfile.ClientID == "" || (p.ServicePrincipalProfile.Secret == "" && p.ServicePrincipalProfile.KeyvaultSecretRef == nil)) {
			if p.OrchestratorProfile.KubernetesConfig != nil && !p.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
				return nil, "", errors.New("when using the kubernetes orchestrator, must either set useManagedIdentity in the kubernetes config or set --client-id and --client-secret or KeyvaultSecretRef of secret (also available in the API model)")
			}
		}
	}

	// This isn't terribly elegant, but it's the easiest way to go for now w/o duplicating a bunch of code
	rawVersionedAPIModel, err := apiloader.SerializeContainerService(dc.containerService, dc.apiVersion)
	if err != nil {
		return nil, "", err
	}
	return apiloader.DeserializeContainerService(rawVersionedAPIModel, true, false, nil)
}

func (dc *deployCmd) run() error {
	ctx := engine.Context{
		Translator: &i18n.Translator{
			Locale: dc.locale,
		},
	}

	templateGenerator, err := engine.InitializeTemplateGenerator(ctx)
	if err != nil {
		log.Fatalf("failed to initialize template generator: %s", err.Error())
	}

	certsgenerated, err := dc.containerService.SetPropertiesDefaults(false, false)
	if err != nil {
		log.Fatalf("error in SetPropertiesDefaults template %s: %s", currentConfig.CLIConfig.Deploy.APIModel, err.Error())
		os.Exit(1)
	}

	template, parameters, err := templateGenerator.GenerateTemplate(dc.containerService, engine.DefaultGeneratorCode, BuildTag)
	if err != nil {
		log.Fatalf("error generating template %s: %s", currentConfig.CLIConfig.Deploy.APIModel, err.Error())
		os.Exit(1)
	}

	if template, err = transform.PrettyPrintArmTemplate(template); err != nil {
		log.Fatalf("error pretty printing template: %s \n", err.Error())
	}
	var parametersFile string
	if parametersFile, err = transform.BuildAzureParametersFile(parameters); err != nil {
		log.Fatalf("error pretty printing template parameters: %s \n", err.Error())
	}

	writer := &engine.ArtifactWriter{
		Translator: &i18n.Translator{
			Locale: dc.locale,
		},
	}
	if err = writer.WriteTLSArtifacts(dc.containerService, dc.apiVersion, template, parametersFile, currentConfig.CLIConfig.Deploy.OutputDirectory, certsgenerated, currentConfig.CLIConfig.Deploy.ParametersOnly); err != nil {
		log.Fatalf("error writing artifacts: %s \n", err.Error())
	}

	templateJSON := make(map[string]interface{})
	parametersJSON := make(map[string]interface{})

	err = json.Unmarshal([]byte(template), &templateJSON)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(parameters), &parametersJSON)
	if err != nil {
		log.Fatalln(err)
	}

	deploymentSuffix := dc.random.Int31()
	cx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	if res, err := dc.client.DeployTemplate(
		cx,
		currentConfig.CLIConfig.Deploy.ResourceGroup,
		fmt.Sprintf("%s-%d", currentConfig.CLIConfig.Deploy.ResourceGroup, deploymentSuffix),
		templateJSON,
		parametersJSON,
	); err != nil {
		if res.Response.Response != nil && res.Body != nil {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			log.Errorf(string(body))
		}
		log.Fatalln(err)
	}

	return nil
}
