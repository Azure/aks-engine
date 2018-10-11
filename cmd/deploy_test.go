package cmd

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
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
	*authArgs
}

func (provider *mockAuthProvider) getClient() (armhelpers.AKSEngineClient, error) {
	if provider.getClientMock == nil {
		return &armhelpers.MockAKSEngineClient{}, nil
	}
	return provider.getClientMock, nil

}
func (provider *mockAuthProvider) getAuthArgs() *authArgs {
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
		dc          *deployCmd
		expectedErr error
		args        []string
	}{
		{
			dc: &deployCmd{
				apimodelPath:      "",
				dnsPrefix:         "test",
				outputDirectory:   "output/test",
				forceOverwrite:    false,
				caCertificatePath: "test",
				caPrivateKeyPath:  "test",
				location:          "west europe",
			},
			args:        []string{},
			expectedErr: nil,
		},
		{
			dc: &deployCmd{
				apimodelPath:      "",
				dnsPrefix:         "test",
				outputDirectory:   "output/test",
				caCertificatePath: "test",
				caPrivateKeyPath:  "test",
			},
			args:        []string{"wrong/path"},
			expectedErr: errors.New("specified api model does not exist (wrong/path)"),
		},
		{
			dc: &deployCmd{
				apimodelPath:      "",
				dnsPrefix:         "test",
				outputDirectory:   "output/test",
				caCertificatePath: "test",
				caPrivateKeyPath:  "test",
			},
			args:        []string{"test/apimodel.json", "some_random_stuff"},
			expectedErr: errors.New("too many arguments were provided to 'deploy'"),
		},
		{
			dc: &deployCmd{
				apimodelPath:      "",
				dnsPrefix:         "test",
				outputDirectory:   "output/test",
				caCertificatePath: "test",
				caPrivateKeyPath:  "test",
			},
			args:        []string{apimodelPath},
			expectedErr: errors.New("--location must be specified"),
		},
		{
			dc: &deployCmd{
				apimodelPath:      "",
				dnsPrefix:         "test",
				outputDirectory:   "output/test",
				caCertificatePath: "test",
				caPrivateKeyPath:  "test",
				location:          "west europe",
			},
			args:        []string{apimodelPath},
			expectedErr: nil,
		},
		{
			dc: &deployCmd{
				apimodelPath:      apimodelPath,
				dnsPrefix:         "test",
				outputDirectory:   "output/test",
				caCertificatePath: "test",
				caPrivateKeyPath:  "test",
				location:          "canadaeast",
			},
			args:        []string{},
			expectedErr: nil,
		},
	}

	for _, c := range cases {
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
	deployCmd := &deployCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		forceOverwrite:   true,
		location:         "westus",
		autoSuffix:       true,
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &authArgs{},
		},
	}

	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(deployCmd.outputDirectory)

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
	deployCmd := &deployCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		forceOverwrite:   true,
		location:         "westus",
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &authArgs{},
		},
	}
	deployCmd.getAuthArgs().ClientID = TestClientIDInCmd
	deployCmd.getAuthArgs().ClientSecret = TestClientSecretInCmd

	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(deployCmd.outputDirectory)

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
	deployCmd := &deployCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		forceOverwrite:   true,
		location:         "westus",
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &authArgs{},
		},
	}
	deployCmd.getAuthArgs().ClientID = TestClientIDInCmd
	deployCmd.getAuthArgs().ClientSecret = TestClientSecretInCmd
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(deployCmd.outputDirectory)

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
	deployCmd := &deployCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		forceOverwrite:   true,
		location:         "westus",
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &authArgs{},
		},
	}
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(deployCmd.outputDirectory)

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
	deployCmd := &deployCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		forceOverwrite:   true,
		location:         "westus",
		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &authArgs{},
		},
	}
	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(deployCmd.outputDirectory)

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
	deployCmd := &deployCmd{
		apimodelPath:    "./this/is/unused.json",
		dnsPrefix:       "dnsPrefix1",
		outputDirectory: "_test_output",
		forceOverwrite:  true,
		location:        "westus",

		containerService: cs,
		apiVersion:       ver,

		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &authArgs{},
		},
	}

	err = autofillApimodel(deployCmd)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	// cleanup, since auto-populations creates dirs and saves the SSH private key that it might create
	defer os.RemoveAll(deployCmd.outputDirectory)

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
	d.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	err := d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with no --set flag defined: %s", err.Error())
	}

	d = &deployCmd{}
	d.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	d.set = []string{"masterProfile.count=3,linuxProfile.adminUsername=testuser"}
	err = d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag: %s", err.Error())
	}

	d = &deployCmd{}
	d.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	d.set = []string{"masterProfile.count=3", "linuxProfile.adminUsername=testuser"}
	err = d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with multiple --set flags: %s", err.Error())
	}

	d = &deployCmd{}
	d.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	d.set = []string{"agentPoolProfiles[0].count=1"}
	err = d.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
	}
}

func TestDeployCmdRun(t *testing.T) {
	d := &deployCmd{
		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs:      &authArgs{},
			getClientMock: &armhelpers.MockAKSEngineClient{},
		},
		apimodelPath:    "./this/is/unused.json",
		outputDirectory: "_test_output",
		forceOverwrite:  true,
		location:        "westus",
	}

	r := &cobra.Command{}
	f := r.Flags()

	addAuthFlags(d.getAuthArgs(), f)

	fakeRawSubscriptionID := "6dc93fae-9a76-421f-bbe5-cc6460ea81cb"
	fakeSubscriptionID, err := uuid.FromString(fakeRawSubscriptionID)
	if err != nil {
		t.Fatalf("Invalid SubscriptionId in Test: %s", err)
	}

	d.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	d.getAuthArgs().SubscriptionID = fakeSubscriptionID
	d.getAuthArgs().rawSubscriptionID = fakeRawSubscriptionID

	err = d.loadAPIModel(r, []string{})
	if err != nil {
		t.Fatalf("Failed to call LoadAPIModel: %s", err)
	}

	err = d.run()
	if err != nil {
		t.Fatalf("Failed to call LoadAPIModel: %s", err)
	}

}
