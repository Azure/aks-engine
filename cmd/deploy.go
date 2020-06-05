// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/leonelquinteros/gotext"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"encoding/json"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/engine/transform"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
)

const (
	deployName             = "deploy"
	deployShortDescription = "Deploy an Azure Resource Manager template"
	deployLongDescription  = "Deploy an Azure Resource Manager template, parameters file and other assets for a cluster"
)

type deployCmd struct {
	authProvider
	apimodelPath      string
	dnsPrefix         string
	autoSuffix        bool
	outputDirectory   string // can be auto-determined from clusterDefinition
	forceOverwrite    bool
	caCertificatePath string
	caPrivateKeyPath  string
	parametersOnly    bool
	set               []string

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
	dc := deployCmd{
		authProvider: &authArgs{},
	}

	deployCmd := &cobra.Command{
		Use:   deployName,
		Short: deployShortDescription,
		Long:  deployLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := dc.validateArgs(cmd, args); err != nil {
				return errors.Wrap(err, "validating deployCmd")
			}
			if err := dc.mergeAPIModel(); err != nil {
				return errors.Wrap(err, "merging API model in deployCmd")
			}
			if err := dc.loadAPIModel(); err != nil {
				return errors.Wrap(err, "loading API model")
			}
			if dc.apiVersion == "vlabs" {
				if err := dc.validateAPIModelAsVLabs(); err != nil {
					return errors.Wrap(err, "validating API model after populating values")
				}
			} else {
				log.Warnf("API model validation is only available for \"apiVersion\": \"vlabs\", skipping validation...")
			}
			return dc.run()
		},
	}

	f := deployCmd.Flags()
	f.StringVarP(&dc.apimodelPath, "api-model", "m", "", "path to your cluster definition file")
	f.StringVarP(&dc.dnsPrefix, "dns-prefix", "p", "", "dns prefix (unique name for the cluster)")
	f.BoolVar(&dc.autoSuffix, "auto-suffix", false, "automatically append a compressed timestamp to the dnsPrefix to ensure unique cluster name automatically")
	f.StringVarP(&dc.outputDirectory, "output-directory", "o", "", "output directory (derived from FQDN if absent)")
	f.StringVar(&dc.caCertificatePath, "ca-certificate-path", "", "path to the CA certificate to use for Kubernetes PKI assets")
	f.StringVar(&dc.caPrivateKeyPath, "ca-private-key-path", "", "path to the CA private key to use for Kubernetes PKI assets")
	f.StringVarP(&dc.resourceGroup, "resource-group", "g", "", "resource group to deploy to (will use the DNS prefix from the apimodel if not specified)")
	f.StringVarP(&dc.location, "location", "l", "", "location to deploy to (required)")
	f.BoolVarP(&dc.forceOverwrite, "force-overwrite", "f", false, "automatically overwrite existing files in the output directory")
	f.StringArrayVar(&dc.set, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")

	addAuthFlags(dc.getAuthArgs(), f)

	return deployCmd
}

func (dc *deployCmd) validateArgs(cmd *cobra.Command, args []string) error {
	var err error

	dc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "loading translation files")
	}

	if dc.apimodelPath == "" {
		if len(args) == 1 {
			dc.apimodelPath = args[0]
		} else if len(args) > 1 {
			_ = cmd.Usage()
			return errors.New("too many arguments were provided to 'deploy'")
		}
	}

	if dc.apimodelPath != "" {
		if _, err := os.Stat(dc.apimodelPath); os.IsNotExist(err) {
			return errors.Errorf("specified api model does not exist (%s)", dc.apimodelPath)
		}
	}

	if dc.location == "" {
		return errors.New("--location must be specified")
	}
	dc.location = helpers.NormalizeAzureRegion(dc.location)

	return nil
}

func (dc *deployCmd) mergeAPIModel() error {
	var err error

	if dc.apimodelPath == "" {
		log.Infoln("no --api-model was specified, using default model")
		var f *os.File
		f, err = ioutil.TempFile("", fmt.Sprintf("%s-default-api-model_%s-%s_", filepath.Base(os.Args[0]), BuildSHA, GitTreeState))
		if err != nil {
			return errors.Wrap(err, "error creating temp file for default API model")
		}
		log.Infoln("default api model generated at", f.Name())

		defer f.Close()
		if err = writeDefaultModel(f); err != nil {
			return err
		}
		dc.apimodelPath = f.Name()
	}

	// if --set flag has been used
	if len(dc.set) > 0 {
		m := make(map[string]transform.APIModelValue)
		transform.MapValues(m, dc.set)

		// overrides the api model and generates a new file
		dc.apimodelPath, err = transform.MergeValuesWithAPIModel(dc.apimodelPath, m)
		if err != nil {
			return errors.Wrapf(err, "error merging --set values with the api model: %s", dc.apimodelPath)
		}

		log.Infoln(fmt.Sprintf("new API model file has been generated during merge: %s", dc.apimodelPath))
	}

	return nil
}

func (dc *deployCmd) loadAPIModel() error {
	var caCertificateBytes []byte
	var caKeyBytes []byte
	var err error

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: dc.locale,
		},
	}

	// do not validate when initially loading the apimodel, validation is done later after autofilling values
	dc.containerService, dc.apiVersion, err = apiloader.LoadContainerServiceFromFile(dc.apimodelPath, false, false, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	// consume dc.caCertificatePath and dc.caPrivateKeyPath
	if (dc.caCertificatePath != "" && dc.caPrivateKeyPath == "") || (dc.caCertificatePath == "" && dc.caPrivateKeyPath != "") {
		return errors.New("--ca-certificate-path and --ca-private-key-path must be specified together")
	}

	if dc.caCertificatePath != "" {
		if caCertificateBytes, err = ioutil.ReadFile(dc.caCertificatePath); err != nil {
			return errors.Wrap(err, "failed to read CA certificate file")
		}
		if caKeyBytes, err = ioutil.ReadFile(dc.caPrivateKeyPath); err != nil {
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
		dc.containerService.Location = dc.location
	} else if dc.containerService.Location != dc.location {
		return errors.New("--location does not match api model location")
	}

	if dc.containerService.Properties.IsCustomCloudProfile() {
		err = dc.containerService.SetCustomCloudProfileEnvironment()
		if err != nil {
			return errors.Wrap(err, "error parsing the api model")
		}
		if err = writeCustomCloudProfile(dc.containerService); err != nil {
			return errors.Wrap(err, "error writing custom cloud profile")
		}

		if dc.containerService.Properties.CustomCloudProfile.IdentitySystem == "" || dc.containerService.Properties.CustomCloudProfile.IdentitySystem != dc.authProvider.getAuthArgs().IdentitySystem {
			if dc.authProvider != nil {
				dc.containerService.Properties.CustomCloudProfile.IdentitySystem = dc.authProvider.getAuthArgs().IdentitySystem
			}
		}
	}

	if err = dc.getAuthArgs().validateAuthArgs(); err != nil {
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

	if dc.containerService.Properties.LinuxProfile != nil {
		if dc.containerService.Properties.LinuxProfile.AdminUsername == "" {
			log.Warnf("apimodel: no linuxProfile.adminUsername was specified. Will use 'azureuser'.")
			dc.containerService.Properties.LinuxProfile.AdminUsername = "azureuser"
		}
	}

	if dc.dnsPrefix != "" && dc.containerService.Properties.MasterProfile.DNSPrefix != "" {
		return errors.New("invalid configuration: the apimodel masterProfile.dnsPrefix and --dns-prefix were both specified")
	}
	if dc.containerService.Properties.MasterProfile.DNSPrefix == "" {
		if dc.dnsPrefix == "" {
			return errors.New("apimodel: missing masterProfile.dnsPrefix and --dns-prefix was not specified")
		}
		log.Warnf("apimodel: missing masterProfile.dnsPrefix will use %q", dc.dnsPrefix)
		dc.containerService.Properties.MasterProfile.DNSPrefix = dc.dnsPrefix
	}

	if dc.autoSuffix {
		suffix := strconv.FormatInt(time.Now().Unix(), 16)
		dc.containerService.Properties.MasterProfile.DNSPrefix += "-" + suffix
	}

	if dc.outputDirectory == "" {
		if dc.containerService.Properties.MasterProfile != nil {
			dc.outputDirectory = path.Join("_output", dc.containerService.Properties.MasterProfile.DNSPrefix)
		} else {
			dc.outputDirectory = path.Join("_output", dc.containerService.Properties.HostedMasterProfile.DNSPrefix)
		}
	}

	if _, err := os.Stat(dc.outputDirectory); !dc.forceOverwrite && err == nil {
		return errors.Errorf("Output directory already exists and forceOverwrite flag is not set: %s", dc.outputDirectory)
	}

	if dc.resourceGroup == "" {
		dnsPrefix := dc.containerService.Properties.MasterProfile.DNSPrefix
		log.Warnf("--resource-group was not specified. Using the DNS prefix from the apimodel as the resource group name: %s", dnsPrefix)
		dc.resourceGroup = dnsPrefix
		if dc.location == "" {
			return errors.New("--resource-group was not specified. --location must be specified in case the resource group needs creation")
		}
	}

	if dc.containerService.Properties.LinuxProfile != nil && (dc.containerService.Properties.LinuxProfile.SSH.PublicKeys == nil ||
		len(dc.containerService.Properties.LinuxProfile.SSH.PublicKeys) == 0 ||
		dc.containerService.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData == "") {
		translator := &i18n.Translator{
			Locale: dc.locale,
		}
		var publicKey string
		_, publicKey, err := helpers.CreateSaveSSH(dc.containerService.Properties.LinuxProfile.AdminUsername, dc.outputDirectory, translator)
		if err != nil {
			return errors.Wrap(err, "Failed to generate SSH Key")
		}

		dc.containerService.Properties.LinuxProfile.SSH.PublicKeys = []api.PublicKey{{KeyData: publicKey}}
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	_, err := dc.client.EnsureResourceGroup(ctx, dc.resourceGroup, dc.location, nil)
	if err != nil {
		return err
	}

	k8sConfig := dc.containerService.Properties.OrchestratorProfile.KubernetesConfig

	useManagedIdentity := k8sConfig != nil && k8sConfig.UseManagedIdentity

	if !useManagedIdentity {
		spp := dc.containerService.Properties.ServicePrincipalProfile
		if spp != nil && spp.ClientID == "" && spp.Secret == "" && spp.KeyvaultSecretRef == nil && (dc.getAuthArgs().ClientID.String() == "" || dc.getAuthArgs().ClientID.String() == "00000000-0000-0000-0000-000000000000") && dc.getAuthArgs().ClientSecret == "" {
			log.Warnln("apimodel: ServicePrincipalProfile was missing or empty, creating application...")

			// TODO: consider caching the creds here so they persist between subsequent runs of 'deploy'
			appName := dc.containerService.Properties.MasterProfile.DNSPrefix
			appURL := fmt.Sprintf("https://%s/", appName)
			var replyURLs *[]string
			var requiredResourceAccess *[]graphrbac.RequiredResourceAccess
			applicationResp, servicePrincipalObjectID, secret, createErr := dc.client.CreateApp(ctx, appName, appURL, replyURLs, requiredResourceAccess)
			if createErr != nil {
				return errors.Wrap(createErr, "apimodel invalid: ServicePrincipalProfile was empty, and we failed to create valid credentials")
			}
			applicationID := to.String(applicationResp.AppID)
			log.Warnf("created application with applicationID (%s) and servicePrincipalObjectID (%s).", applicationID, servicePrincipalObjectID)

			log.Warnln("apimodel: ServicePrincipalProfile was empty, assigning role to application...")

			err = dc.client.CreateRoleAssignmentSimple(ctx, dc.resourceGroup, servicePrincipalObjectID)
			if err != nil {
				return errors.Wrap(err, "apimodel: could not create or assign ServicePrincipal")

			}

			dc.containerService.Properties.ServicePrincipalProfile = &api.ServicePrincipalProfile{
				ClientID: applicationID,
				Secret:   secret,
				ObjectID: servicePrincipalObjectID,
			}
		} else if (dc.containerService.Properties.ServicePrincipalProfile == nil || ((dc.containerService.Properties.ServicePrincipalProfile.ClientID == "" || dc.containerService.Properties.ServicePrincipalProfile.ClientID == "00000000-0000-0000-0000-000000000000") && dc.containerService.Properties.ServicePrincipalProfile.Secret == "")) && dc.getAuthArgs().ClientID.String() != "" && dc.getAuthArgs().ClientSecret != "" {
			dc.containerService.Properties.ServicePrincipalProfile = &api.ServicePrincipalProfile{
				ClientID: dc.getAuthArgs().ClientID.String(),
				Secret:   dc.getAuthArgs().ClientSecret,
			}
		}
	}

	if k8sConfig != nil && k8sConfig.Addons != nil && k8sConfig.IsContainerMonitoringAddonEnabled() {
		log.Infoln("container monitoring addon enabled")
		cloudOrDependenciesLocation := dc.containerService.GetCloudSpecConfig().CloudName
		if dc.containerService.Properties.IsCustomCloudProfile() {
			cloudOrDependenciesLocation = string(dc.containerService.Properties.CustomCloudProfile.DependenciesLocation)
		}
		workspaceDomain := helpers.GetLogAnalyticsWorkspaceDomain(cloudOrDependenciesLocation)
		err := dc.configureContainerMonitoringAddon(ctx, k8sConfig, workspaceDomain)
		if err != nil {
			return errors.Wrap(err, "Failed to configure container monitoring addon")
		}
	}

	return nil
}

// validateAPIModelAsVLabs converts the ContainerService object to a vlabs ContainerService object and validates it
func (dc *deployCmd) validateAPIModelAsVLabs() error {
	return api.ConvertContainerServiceToVLabs(dc.containerService).Validate(false)
}

func (dc *deployCmd) run() error {
	ctx := engine.Context{
		Translator: &i18n.Translator{
			Locale: dc.locale,
		},
	}

	templateGenerator, err := engine.InitializeTemplateGenerator(ctx)
	if err != nil {
		return errors.Wrap(err, "initializing template generator")
	}

	certsgenerated, err := dc.containerService.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		return errors.Wrapf(err, "in SetPropertiesDefaults template %s", dc.apimodelPath)
	}

	if err = dc.validateOSBaseImage(); err != nil {
		return errors.Wrapf(err, "validating OS base images required by %s", dc.apimodelPath)
	}

	template, parameters, err := templateGenerator.GenerateTemplateV2(dc.containerService, engine.DefaultGeneratorCode, BuildTag)
	if err != nil {
		return errors.Wrapf(err, "generating template %s", dc.apimodelPath)
	}

	if template, err = transform.PrettyPrintArmTemplate(template); err != nil {
		return errors.Wrap(err, "pretty-printing template")
	}
	var parametersFile string
	if parametersFile, err = transform.BuildAzureParametersFile(parameters); err != nil {
		return errors.Wrap(err, "pretty-printing template parameters")
	}

	writer := &engine.ArtifactWriter{
		Translator: &i18n.Translator{
			Locale: dc.locale,
		},
	}
	if err = writer.WriteTLSArtifacts(dc.containerService, dc.apiVersion, template, parametersFile, dc.outputDirectory, certsgenerated, dc.parametersOnly); err != nil {
		return errors.Wrap(err, "writing artifacts")
	}

	templateJSON := make(map[string]interface{})
	parametersJSON := make(map[string]interface{})

	if err = json.Unmarshal([]byte(template), &templateJSON); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(parameters), &parametersJSON); err != nil {
		return err
	}

	cx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()

	deploymentSuffix := dc.random.Int31()

	if res, err := dc.client.DeployTemplate(
		cx,
		dc.resourceGroup,
		fmt.Sprintf("%s-%d", dc.resourceGroup, deploymentSuffix),
		templateJSON,
		parametersJSON,
	); err != nil {
		if res.Response.Response != nil && res.Body != nil {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			log.Errorf(string(body))
		}
		return err
	}

	return nil
}

// configure api model addon config with container monitoring addon
func (dc *deployCmd) configureContainerMonitoringAddon(ctx context.Context, k8sConfig *api.KubernetesConfig, workspaceDomain string) error {
	log.Infoln("configuring container monitoring addon info")
	if k8sConfig == nil {
		return errors.New("KubernetesConfig either null or invalid")
	}

	var workspaceResourceID string
	var err error
	addon := k8sConfig.GetAddonByName("container-monitoring")
	if addon.Config == nil || len(addon.Config) == 0 || addon.Config["logAnalyticsWorkspaceResourceId"] != "" {
		if dc.containerService.Properties.IsAzureStackCloud() {
			return errors.New("This is not supported option for AzureStackCloud. Please provide config with workspaceGuid and workspaceKey")
		}
		workspaceResourceID = strings.TrimSpace(addon.Config["logAnalyticsWorkspaceResourceId"])
		if workspaceResourceID != "" {
			log.Infoln("using provided log analytics workspace resource id:", workspaceResourceID)
			if !strings.HasPrefix(workspaceResourceID, "/") {
				workspaceResourceID = "/" + workspaceResourceID
			}
			workspaceResourceID = strings.TrimSuffix(workspaceResourceID, "/")
		} else {
			log.Infoln("creating default log analytics workspace if not exists already")
			workspaceResourceID, err = dc.client.EnsureDefaultLogAnalyticsWorkspace(ctx, dc.resourceGroup, dc.location)
			if err != nil {
				return errors.Wrap(err, "apimodel: Failed to create default log analytics workspace for container monitoring addon")
			}
			log.Infoln("successfully created or fetched default log analytics workspace:", workspaceResourceID)
		}
		resourceParts := strings.Split(workspaceResourceID, "/")
		if len(resourceParts) != 9 {
			return errors.Errorf("%s is not a valid azure resource id", workspaceResourceID)
		}
		workspaceSubscriptionID := resourceParts[2]
		workspaceResourceGroup := resourceParts[4]
		workspaceName := resourceParts[8]
		log.Infoln("Retrieving log analytics workspace Guid, Key and location details for the workspace resource:", workspaceResourceID)
		wsID, wsKey, wsLocation, err := dc.client.GetLogAnalyticsWorkspaceInfo(ctx, workspaceSubscriptionID, workspaceResourceGroup, workspaceName)
		if err != nil {
			return errors.Wrap(err, "apimodel: Failed to get the workspace Guid, Key and location details ")
		}
		log.Infoln("successfully retrieved log analytics workspace details")
		log.Infoln("log analytics workspace id: ", wsID)

		log.Infoln("adding container insights solution to log analytics workspace: ", workspaceResourceID)
		_, err = dc.client.AddContainerInsightsSolution(ctx, workspaceSubscriptionID, workspaceResourceGroup, workspaceName, wsLocation)
		if err != nil {
			return errors.Wrap(err, "apimodel: Failed to get add container insights solution")
		}
		log.Infoln("successfully added container insights solution to log analytics workspace: ", workspaceResourceID)

		log.Infoln("Adding log analytics workspaceGuid and workspaceKey, workspaceResourceId to the container monitoring addon")
		for _, addon := range dc.containerService.Properties.OrchestratorProfile.KubernetesConfig.Addons {
			if addon.Name == "container-monitoring" {
				addon.Config["workspaceGuid"] = base64.StdEncoding.EncodeToString([]byte(wsID))
				addon.Config["workspaceKey"] = base64.StdEncoding.EncodeToString([]byte(wsKey))
				addon.Config["logAnalyticsWorkspaceResourceId"] = workspaceResourceID
				addon.Config["workspaceDomain"] = base64.StdEncoding.EncodeToString([]byte(workspaceDomain))
			}
		}

	} else {
		log.Infoln("using provided workspaceGuid, workspaceKey in the container addon config")
		workspaceGUID := addon.Config["workspaceGuid"]
		workspaceKey := addon.Config["workspaceKey"]
		addon.Config["workspaceDomain"] = base64.StdEncoding.EncodeToString([]byte(workspaceDomain))
		log.Infoln("workspaceGuid:", workspaceGUID)
		log.Infoln("workspaceKey:", workspaceKey)
		log.Infoln("workspaceDomain:", workspaceDomain)
	}
	return nil
}

// validateOSBaseImage checks if the OS image is available on the target cloud (ATM, Azure Stack only)
func (dc *deployCmd) validateOSBaseImage() error {
	if dc.containerService.Properties.IsAzureStackCloud() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := armhelpers.ValidateRequiredImages(ctx, dc.location, dc.containerService.Properties, dc.client); err != nil {
			return errors.Wrap(err, "OS base image not available in target cloud")
		}
	}
	return nil
}
