// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/imdario/mergo"

	"os"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/cli/config"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

const ExampleAPIModel = `{
  "apiVersion": "vlabs",
  "properties": {
		"orchestratorProfile": { "orchestratorType": "Kubernetes", "kubernetesConfig": { "useManagedIdentity": %s, "etcdVersion" : "2.3.8" } },
    "masterProfile": { "count": 1, "dnsPrefix": "", "vmSize": "Standard_D2_v2" },
    "agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
    "windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
    "linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
    },
    "servicePrincipalProfile": { "clientId": "%s", "secret": "%s" }
  }
}
`

const ExampleAPIModelWithDNSPrefix = `{
	"apiVersion": "vlabs",
	"properties": {
		  "orchestratorProfile": { "orchestratorType": "Kubernetes", "kubernetesConfig": { "useManagedIdentity": %s, "etcdVersion" : "2.3.8" } },
	  "masterProfile": { "count": 1, "dnsPrefix": "mytestcluster", "vmSize": "Standard_D2_v2" },
	  "agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
	  "windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
	  "linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	  },
	  "servicePrincipalProfile": { "clientId": "%s", "secret": "%s" }
	}
  }
  `

const ExampleAPIModelWithoutServicePrincipalProfile = `{
	"apiVersion": "vlabs",
	"properties": {
		  "orchestratorProfile": { "orchestratorType": "Kubernetes", "kubernetesConfig": { "useManagedIdentity": %s, "etcdVersion" : "2.3.8" } },
	  "masterProfile": { "count": 1, "dnsPrefix": "mytestcluster", "vmSize": "Standard_D2_v2" },
	  "agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
	  "windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
	  "linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	  }
	}
  }
  `

//mockAuthProvider implements AuthProvider and allows in particular to stub out getClient()
type mockAuthProvider struct {
	getClientMock armhelpers.AKSEngineClient
	authArgs      *config.AuthConfig
}

func (provider *mockAuthProvider) getClient() (armhelpers.AKSEngineClient, error) {
	if provider.getClientMock == nil {
		return &armhelpers.MockAKSEngineClient{}, nil
	}
	return provider.getClientMock, nil

}
func (provider *mockAuthProvider) getAuthArgs() *config.AuthConfig {
	return provider.authArgs
}

func getExampleAPIModel(useManagedIdentity bool, clientID, clientSecret string) string {
	return getAPIModel(ExampleAPIModel, useManagedIdentity, clientID, clientSecret)
}

func getAPIModel(baseAPIModel string, useManagedIdentity bool, clientID, clientSecret string) string {
	return fmt.Sprintf(
		baseAPIModel,
		strconv.FormatBool(useManagedIdentity),
		clientID,
		clientSecret)
}

func getAPIModelWithoutServicePrincipalProfile(baseAPIModel string, useManagedIdentity bool) string {
	return fmt.Sprintf(
		baseAPIModel,
		strconv.FormatBool(useManagedIdentity))
}

func TestNewDeployCmd(t *testing.T) {
	output := newDeployCmd()
	if output.Use != deployName || output.Short != deployShortDescription || output.Long != deployLongDescription {
		t.Fatalf("deploy command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, deployName, output.Short, deployShortDescription, output.Long, versionLongDescription)
	}

	expectedFlags := []string{"api-model", "dns-prefix", "auto-suffix", "output-directory", "ca-private-key-path", "resource-group", "location", "force-overwrite"}
	for _, f := range expectedFlags {
		if output.Flags().Lookup(f) == nil {
			t.Fatalf("deploy command should have flag %s", f)
		}
	}
}

func TestValidate(t *testing.T) {
	r := &cobra.Command{}
	apimodelPath := "apimodel-unit-test.json"

	_, err := os.Create(apimodelPath)
	if err != nil {
		t.Fatalf("unable to create test apimodel path: %s", err.Error())
	}
	defer os.Remove(apimodelPath)

	cases := []struct {
		config      config.DeployConfig
		dc          *deployCmd
		expectedErr error
		args        []string
	}{
		{
			config: config.DeployConfig{
				DNSPrefix:         "test",
				OutputDirectory:   "output/test",
				CACertificatePath: "test",
				CAPrivateKeyPath:  "test",
				Location:          "west europe",
			},
			dc:          &deployCmd{},
			args:        []string{},
			expectedErr: nil,
		},
		{
			config: config.DeployConfig{
				DNSPrefix:         "test",
				OutputDirectory:   "output/test",
				CACertificatePath: "test",
				CAPrivateKeyPath:  "test",
			},
			dc:          &deployCmd{},
			args:        []string{"wrong/path"},
			expectedErr: errors.New("specified api model does not exist (wrong/path)"),
		},
		{
			config: config.DeployConfig{
				DNSPrefix:         "test",
				OutputDirectory:   "output/test",
				CACertificatePath: "test",
				CAPrivateKeyPath:  "test",
			},
			dc:          &deployCmd{},
			args:        []string{"test/apimodel.json", "some_random_stuff"},
			expectedErr: errors.New("too many arguments were provided to 'deploy'"),
		},
		{
			config: config.DeployConfig{
				DNSPrefix:         "test",
				OutputDirectory:   "output/test",
				CACertificatePath: "test",
				CAPrivateKeyPath:  "test",
			},
			dc:          &deployCmd{},
			args:        []string{apimodelPath},
			expectedErr: errors.New("--location must be specified"),
		},
		{
			config: config.DeployConfig{
				DNSPrefix:         "test",
				OutputDirectory:   "output/test",
				CACertificatePath: "test",
				CAPrivateKeyPath:  "test",
				Location:          "west europe",
			},
			dc:          &deployCmd{},
			args:        []string{apimodelPath},
			expectedErr: nil,
		},
		{
			config: config.DeployConfig{
				APIModel:          apimodelPath,
				DNSPrefix:         "test",
				OutputDirectory:   "output/test",
				CACertificatePath: "test",
				CAPrivateKeyPath:  "test",
				Location:          "canadaeast",
			},
			dc:          &deployCmd{},
			args:        []string{},
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		currentConfig.CLIConfig.Deploy = c.config
		if err := mergo.Merge(&currentConfig.CLIConfig.Deploy, defaultConfigValues.CLIConfig.Deploy); err != nil {
			t.Fatal(err)
		}
		err = c.dc.validateArgs(r, c.args)
		if err != nil && c.expectedErr != nil {
			if err.Error() != c.expectedErr.Error() {
				t.Fatalf("expected validate deploy command to return error %s, but instead got %s", c.expectedErr.Error(), err.Error())
			}
		} else {
			if c.expectedErr != nil {
				t.Fatalf("expected validate deploy command to return error %s, but instead got no error", c.expectedErr.Error())
			} else if err != nil {
				t.Fatalf("expected validate deploy command to return no error, but instead got %s", err.Error())
			}
		}
	}
}

func TestAutofillApimodelWithoutManagedIdentityCreatesCreds(t *testing.T) {
	testAutodeployCredentialHandling(t, false, "", "")
}

func TestAutofillApimodelWithManagedIdentitySkipsCreds(t *testing.T) {
	testAutodeployCredentialHandling(t, true, "", "")
}

func TestAutofillApimodelAllowsPrespecifiedCreds(t *testing.T) {
	testAutodeployCredentialHandling(t, false, "clientID", "clientSecret")
}

func TestAutoSufixWithDnsPrefixInApiModel(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModel(ExampleAPIModelWithDNSPrefix, false, "clientID", "clientSecret")
	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	conf := config.DeployConfig{
		APIModel:        "./this/is/unused.json",
		OutputDirectory: "_test_output",
		ForceOverwrite:  true,
		Location:        "westus",
		AutoSuffix:      true,
	}
	deployCmd := &deployCmd{
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &config.AuthConfig{},
		},
	}

	currentConfig.CLIConfig.Deploy = conf
	defer func() {
		currentConfig = config.Config{}
	}()
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(conf.OutputDirectory)

	if deployCmd.containerService.Properties.MasterProfile.DNSPrefix == "mytestcluster" {
		t.Fatalf("expected %s-{timestampsuffix} but got %s", "mytestcluster", deployCmd.containerService.Properties.MasterProfile.DNSPrefix)
	}

}

func TestAPIModelWithoutServicePrincipalProfileAndClientIdAndSecretInCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(ExampleAPIModelWithoutServicePrincipalProfile, false)
	TestClientIDInCmd, err := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")
	if err != nil {
		t.Fatalf("Invalid ClientID in Test: %s", err)
	}

	TestClientSecretInCmd := "DEC923E3-1EF1-4745-9516-37906D56DEC4"

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	conf := config.DeployConfig{
		APIModel:        "./this/is/unused.json",
		OutputDirectory: "_test_output",
		ForceOverwrite:  true,
		Location:        "westus",
	}
	deployCmd := &deployCmd{
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &config.AuthConfig{},
		},
	}
	deployCmd.getAuthArgs().ClientID = TestClientIDInCmd.String()
	deployCmd.getAuthArgs().ClientSecret = TestClientSecretInCmd

	currentConfig.CLIConfig.Deploy = conf
	defer func() {
		currentConfig = config.Config{}
	}()
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(conf.OutputDirectory)

	if deployCmd.containerService.Properties.ServicePrincipalProfile == nil || deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID == "" || deployCmd.containerService.Properties.ServicePrincipalProfile.Secret == "" {
		t.Fatalf("expected service principal profile to be populated from deployment command arguments")
	}

	if deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID != TestClientIDInCmd.String() {
		t.Fatalf("expected service principal profile client id to be %s but got %s", TestClientIDInCmd.String(), deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID)
	}

	if deployCmd.containerService.Properties.ServicePrincipalProfile.Secret != TestClientSecretInCmd {
		t.Fatalf("expected service principal profile client secret to be %s but got %s", TestClientSecretInCmd, deployCmd.containerService.Properties.ServicePrincipalProfile.Secret)
	}
}

func TestAPIModelWithEmptyServicePrincipalProfileAndClientIdAndSecretInCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModel(ExampleAPIModelWithDNSPrefix, false, "", "")
	TestClientIDInCmd, err := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")
	if err != nil {
		t.Fatalf("Invalid ClientID in Test: %s", err)
	}

	TestClientSecretInCmd := "DEC923E3-1EF1-4745-9516-37906D56DEC4"

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	conf := config.DeployConfig{
		APIModel:        "./this/is/unused.json",
		OutputDirectory: "_test_output",
		ForceOverwrite:  true,
		Location:        "westus",
	}
	deployCmd := &deployCmd{
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &config.AuthConfig{},
		},
	}
	deployCmd.getAuthArgs().ClientID = TestClientIDInCmd.String()
	deployCmd.getAuthArgs().ClientSecret = TestClientSecretInCmd
	currentConfig.CLIConfig.Deploy = conf
	defer func() {
		currentConfig = config.Config{}
	}()
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(conf.OutputDirectory)

	if deployCmd.containerService.Properties.ServicePrincipalProfile == nil || deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID == "" || deployCmd.containerService.Properties.ServicePrincipalProfile.Secret == "" {
		t.Fatalf("expected service principal profile to be populated from deployment command arguments")
	}

	if deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID != TestClientIDInCmd.String() {
		t.Fatalf("expected service principal profile client id to be %s but got %s", TestClientIDInCmd.String(), deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID)
	}

	if deployCmd.containerService.Properties.ServicePrincipalProfile.Secret != TestClientSecretInCmd {
		t.Fatalf("expected service principal profile client secret to be %s but got %s", TestClientSecretInCmd, deployCmd.containerService.Properties.ServicePrincipalProfile.Secret)
	}
}

func TestAPIModelWithoutServicePrincipalProfileAndWithoutClientIdAndSecretInCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(ExampleAPIModelWithoutServicePrincipalProfile, false)

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	conf := config.DeployConfig{
		APIModel:        "./this/is/unused.json",
		OutputDirectory: "_test_output",
		ForceOverwrite:  true,
		Location:        "westus",
	}
	deployCmd := &deployCmd{
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &config.AuthConfig{},
		},
	}
	currentConfig.CLIConfig.Deploy = conf
	defer func() {
		currentConfig = config.Config{}
	}()
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(conf.OutputDirectory)

	if deployCmd.containerService.Properties.ServicePrincipalProfile != nil {
		t.Fatalf("expected service principal profile to be nil for unmanaged identity, where client id and secret are not supplied in api model and deployment command")
	}

}

func TestAPIModelWithEmptyServicePrincipalProfileAndWithoutClientIdAndSecretInCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModel(ExampleAPIModelWithDNSPrefix, false, "", "")

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	conf := config.DeployConfig{
		APIModel:        "./this/is/unused.json",
		OutputDirectory: "_test_output",
		ForceOverwrite:  true,
		Location:        "westus",
	}
	deployCmd := &deployCmd{
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &config.AuthConfig{},
		},
	}
	currentConfig.CLIConfig.Deploy = conf
	defer func() {
		currentConfig = config.Config{}
	}()
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(conf.OutputDirectory)

	if deployCmd.containerService.Properties.ServicePrincipalProfile == nil {
		t.Fatalf("expected service principal profile to be Empty and not nil for unmanaged identity, where client id and secret are not supplied in api model and deployment command")
	}

	// mockclient returns "app-id" for ClientID when empty
	if deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID != "app-id" {
		t.Fatalf("expected service principal profile client id to be empty but got %s", deployCmd.containerService.Properties.ServicePrincipalProfile.ClientID)
	}

	// mockcliet returns "client-secret" when empty
	if deployCmd.containerService.Properties.ServicePrincipalProfile.Secret != "client-secret" {
		t.Fatalf("expected service principal profile client secret to be empty but got %s", deployCmd.containerService.Properties.ServicePrincipalProfile.Secret)
	}

}

func testAutodeployCredentialHandling(t *testing.T, useManagedIdentity bool, clientID, clientSecret string) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getExampleAPIModel(useManagedIdentity, clientID, clientSecret)
	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}

	// deserialization happens in validate(), but we are testing just the default
	// setting that occurs in autofillApimodel (which is called from validate)
	// Thus, it assumes that containerService/apiVersion are already populated
	conf := config.DeployConfig{
		APIModel:        "./this/is/unused.json",
		DNSPrefix:       "dnsPrefix1",
		OutputDirectory: "_test_output",
		ForceOverwrite:  true,
		Location:        "westus",
	}
	deployCmd := &deployCmd{
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &config.AuthConfig{},
		},
	}
	currentConfig.CLIConfig.Deploy = conf
	defer func() {
		currentConfig = config.Config{}
	}()

	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	// cleanup, since auto-populations creates dirs and saves the SSH private key that it might create
	defer os.RemoveAll(conf.OutputDirectory)

	cs, _, err = deployCmd.validateApimodel()
	if err != nil {
		t.Fatalf("unexpected error validating apimodel after populating defaults: %s", err)
	}

	if useManagedIdentity {
		if cs.Properties.ServicePrincipalProfile != nil &&
			(cs.Properties.ServicePrincipalProfile.ClientID != "" || cs.Properties.ServicePrincipalProfile.Secret != "") {
			t.Fatalf("Unexpected credentials were populated even though MSI was active.")
		}
	} else {
		if cs.Properties.ServicePrincipalProfile == nil ||
			cs.Properties.ServicePrincipalProfile.ClientID == "" || cs.Properties.ServicePrincipalProfile.Secret == "" {
			t.Fatalf("Credentials were missing even though MSI was not active.")
		}
	}
}

func TestDeployCmdMergeAPIModel(t *testing.T) {
	d := &deployCmd{}
	currentConfig.CLIConfig.Deploy.APIModel = "../pkg/engine/testdata/simple/kubernetes.json"
	err := d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with no --set flag defined: %s", err.Error())
	}

	d = &deployCmd{}
	currentConfig.CLIConfig.Deploy.APIModel = "../pkg/engine/testdata/simple/kubernetes.json"
	currentConfig.CLIConfig.Deploy.Set = []string{"masterProfile.count=3,linuxProfile.adminUsername=testuser"}
	err = d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag: %s", err.Error())
	}

	d = &deployCmd{}
	currentConfig.CLIConfig.Deploy.APIModel = "../pkg/engine/testdata/simple/kubernetes.json"
	currentConfig.CLIConfig.Deploy.Set = []string{"masterProfile.count=3", "linuxProfile.adminUsername=testuser"}
	err = d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with multiple --set flags: %s", err.Error())
	}

	d = &deployCmd{}
	currentConfig.CLIConfig.Deploy.APIModel = "../pkg/engine/testdata/simple/kubernetes.json"
	currentConfig.CLIConfig.Deploy.Set = []string{"agentPoolProfiles[0].count=1"}
	err = d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
	}
}

func TestDeployCmdRun(t *testing.T) {
	conf := config.DeployConfig{
		APIModel:        "./this/is/unused.json",
		OutputDirectory: "_test_output",
		ForceOverwrite:  true,
		Location:        "westus",
	}
	d := &deployCmd{
		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs:      &config.AuthConfig{},
			getClientMock: &armhelpers.MockAKSEngineClient{},
		},
	}
	if err := mergo.Merge(&currentConfig.CLIConfig.Deploy, conf); err != nil {
		t.Fatal(err)
	}

	r := &cobra.Command{}
	f := r.Flags()

	addAuthFlags(d.getAuthArgs(), f)

	fakeSubscriptionID := "6dc93fae-9a76-421f-bbe5-cc6460ea81cb"
	fakeClientID := "b829b379-ca1f-4f1d-91a2-0d26b244680d"
	fakeClientSecret := "0se43bie-3zs5-303e-aav5-dcf231vb82ds"

	currentConfig.CLIConfig.Deploy.APIModel = "../pkg/engine/testdata/simple/kubernetes.json"
	d.getAuthArgs().SubscriptionID = fakeSubscriptionID
	d.getAuthArgs().ClientID = fakeClientID
	d.getAuthArgs().ClientSecret = fakeClientSecret

	if err := d.loadAPIModel(r, []string{}); err != nil {
		t.Fatalf("Failed to call LoadAPIModel: %s", err)
	}

	if err := d.run(); err != nil {
		t.Fatalf("Failed to call LoadAPIModel: %s", err)
	}
}
