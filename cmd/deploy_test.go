// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"testing"

	"os"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
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

const ExampleAPIModelWithoutDNSPrefix = `{
	"apiVersion": "vlabs",
	"properties": {
		  "orchestratorProfile": { "orchestratorType": "Kubernetes", "kubernetesConfig": { "useManagedIdentity": %s, "etcdVersion" : "2.3.8" } },
	  "masterProfile": { "count": 1, "vmSize": "Standard_D2_v2" },
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
const ExampleAPIModelWithContainerMonitoringAddonWithNoConfig = `{
	"apiVersion": "vlabs",
	"properties": {
		  "orchestratorProfile": { "orchestratorType": "Kubernetes", "kubernetesConfig": { "addons":[{"name": "container-monitoring","enabled": true }]}},
	  "masterProfile": { "count": 1, "dnsPrefix": "mytestcluster", "vmSize": "Standard_D2_v2" },
	  "agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
	  "windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
	  "linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	  }
	}
  }
  `

const ExampleAPIModelWithContainerMonitoringAddonWithExistingWorkspaceConfig = `{
	"apiVersion": "vlabs",
	"properties": {
		  "orchestratorProfile": { "orchestratorType": "Kubernetes", "kubernetesConfig": { "addons":[{"name": "container-monitoring","enabled": true, "config":{"logAnalyticsWorkspaceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-workspace-rg/providers/Microsoft.OperationalInsights/workspaces/test-workspace"} }]}},
	  "masterProfile": { "count": 1, "dnsPrefix": "mytestcluster", "vmSize": "Standard_D2_v2" },
	  "agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
	  "windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
	  "linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	  }
	}
  }
  `
const ExampleAPIModelWithContainerMonitoringAddonWithWorkspaceGUIDAndKeyConfig = `{
	"apiVersion": "vlabs",
	"properties": {
		  "orchestratorProfile": { "orchestratorType": "Kubernetes", "kubernetesConfig": { "addons":[{"name": "container-monitoring","enabled": true, "config":{"workspaceGuid": "MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAw", "workspaceKey": "NEQrdnlkNS9qU2NCbXNBd1pPRi8wR09CUTVrdUZRYzlKVmFXK0hsbko1OGN5ZVBKY3dUcGtzK3JWbXZnY1hHbW15dWpMRE5FVlBpVDhwQjI3NGE5WWc9PQ=="} }]}},
	  "masterProfile": { "count": 1, "dnsPrefix": "mytestcluster", "vmSize": "Standard_D2_v2" },
	  "agentPoolProfiles": [ { "name": "linuxpool1", "count": 2, "vmSize": "Standard_D2_v2", "availabilityProfile": "AvailabilitySet" } ],
	  "windowsProfile": { "adminUsername": "azureuser", "adminPassword": "replacepassword1234$" },
	  "linuxProfile": { "adminUsername": "azureuser", "ssh": { "publicKeys": [ { "keyData": "" } ] }
	  }
	}
  }
  `

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

func getAPIModelWithoutServicePrincipalProfile(useManagedIdentity bool) string {
	return fmt.Sprintf(
		ExampleAPIModelWithoutServicePrincipalProfile,
		strconv.FormatBool(useManagedIdentity))
}

func TestNewDeployCmd(t *testing.T) {
	command := newDeployCmd()
	if command.Use != deployName || command.Short != deployShortDescription || command.Long != deployLongDescription {
		t.Fatalf("deploy command should have use %s equal %s, short %s equal %s and long %s equal to %s", command.Use, deployName, command.Short, deployShortDescription, command.Long, versionLongDescription)
	}

	expectedFlags := []string{"api-model", "dns-prefix", "auto-suffix", "output-directory", "ca-private-key-path", "resource-group", "location", "force-overwrite"}
	for _, f := range expectedFlags {
		if command.Flags().Lookup(f) == nil {
			t.Fatalf("deploy command should have flag %s", f)
		}
	}

	command.SetArgs([]string{})
	if err := command.Execute(); err == nil {
		t.Fatalf("expected an error when calling deploy with no arguments")
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

	apimodel := getAPIModelWithoutServicePrincipalProfile(false)
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

	apimodel := getAPIModelWithoutServicePrincipalProfile(false)

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

	err = deployCmd.validateAPIModelAsVLabs()
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
		apimodelPath:    "../pkg/engine/testdata/simple/kubernetes.json",
		outputDirectory: "_test_output",
		forceOverwrite:  true,
		location:        "westus",
	}

	r := &cobra.Command{}
	f := r.Flags()

	addAuthFlags(d.getAuthArgs(), f)

	fakeRawSubscriptionID := "6dc93fae-9a76-421f-bbe5-cc6460ea81cb"
	fakeSubscriptionID, err := uuid.FromString(fakeRawSubscriptionID)
	fakeClientID := "b829b379-ca1f-4f1d-91a2-0d26b244680d"
	fakeClientSecret := "0se43bie-3zs5-303e-aav5-dcf231vb82ds"
	if err != nil {
		t.Fatalf("Invalid SubscriptionId in Test: %s", err)
	}

	d.getAuthArgs().SubscriptionID = fakeSubscriptionID
	d.getAuthArgs().rawSubscriptionID = fakeRawSubscriptionID
	d.getAuthArgs().rawClientID = fakeClientID
	d.getAuthArgs().ClientSecret = fakeClientSecret
	if err != nil {
		t.Fatalf("Invalid SubscriptionId in Test: %s", err)
	}

	err = d.loadAPIModel()
	if err != nil {
		t.Fatalf("Failed to call LoadAPIModel: %s", err)
	}

	err = d.run()
	if err != nil {
		t.Fatalf("Failed to call LoadAPIModel: %s", err)
	}
}

func TestOutputDirectoryWithDNSPrefix(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModel(ExampleAPIModelWithoutDNSPrefix, false, "clientID", "clientSecret")
	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}

	d := &deployCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "",
		dnsPrefix:        "dnsPrefix1",
		forceOverwrite:   true,
		location:         "westus",
		containerService: cs,
		apiVersion:       ver,
		client:           &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs: &authArgs{},
		},
	}

	err = autofillApimodel(d)
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(d.outputDirectory)

	if d.outputDirectory != path.Join("_output", d.dnsPrefix) {
		t.Fatalf("Calculated output directory should be %s, actual value %s", path.Join("_output", d.dnsPrefix), d.outputDirectory)
	}
}

func TestAPIModelWithContainerMonitoringAddonWithNoConfigInCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := ExampleAPIModelWithContainerMonitoringAddonWithNoConfig

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

	k8sConfig := deployCmd.containerService.Properties.OrchestratorProfile.KubernetesConfig

	if k8sConfig == nil {
		t.Fatalf("expected valid kubernetes config")
	}

	if len(k8sConfig.Addons) != 1 {
		t.Fatalf("expected one addon")
	}

	addon := k8sConfig.Addons[0]
	expectedAddonName := "container-monitoring"
	if addon.Name != expectedAddonName {
		t.Fatalf("expected addon name: %s but got: %s", expectedAddonName, addon.Name)
	}

	expectedWorkspaceGUIDInBase64 := "MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAw"
	if addon.Config["workspaceGuid"] != expectedWorkspaceGUIDInBase64 {
		t.Fatalf("expected workspaceGuid : %s but got : %s", expectedWorkspaceGUIDInBase64, addon.Config["workspaceGuid"])
	}

	expectedWorkspaceKeyInBase64 := "NEQrdnlkNS9qU2NCbXNBd1pPRi8wR09CUTVrdUZRYzlKVmFXK0hsbko1OGN5ZVBKY3dUcGtzK3JWbXZnY1hHbW15dWpMRE5FVlBpVDhwQjI3NGE5WWc9PQ=="
	if addon.Config["workspaceKey"] != expectedWorkspaceKeyInBase64 {
		t.Fatalf("unexpected workspaceKey : %s", addon.Config["workspaceKey"])
		t.Fatalf("expected workspaceKey : %s but got : %s", expectedWorkspaceKeyInBase64, addon.Config["workspaceKey"])
	}

	workspaceResourceID := addon.Config["logAnalyticsWorkspaceResourceId"]
	resourceParts := strings.Split(workspaceResourceID, "/")

	if len(resourceParts) != 9 {
		t.Fatalf("invalid workspaceResourceID : %s", workspaceResourceID)
	}

	workspaceSubscriptionID := resourceParts[2]
	workspaceResourceGroup := resourceParts[4]
	workspaceProvider := resourceParts[6]
	workspaceName := resourceParts[8]

	expectedworkspaceSubscriptionID := "00000000-0000-0000-0000-000000000000"
	if workspaceSubscriptionID != expectedworkspaceSubscriptionID {
		t.Fatalf("expected workspaceSubscriptionID: %s, but found : %s", expectedworkspaceSubscriptionID, workspaceSubscriptionID)
	}

	expectedworkspaceProvider := "Microsoft.OperationalInsights"
	if workspaceProvider != expectedworkspaceProvider {
		t.Fatalf("expected log analytics workspace provider name: %s, but got: %s", expectedworkspaceProvider, workspaceProvider)
	}

	expectedworkspaceResourceGroup := "test-workspace-rg"
	if workspaceResourceGroup != "test-workspace-rg" {
		t.Fatalf("expected workspaceResourceGroup : %s, but found : %s", expectedworkspaceResourceGroup, workspaceResourceGroup)
	}

	expectedworkspaceName := "test-workspace"
	if workspaceName != expectedworkspaceName {
		t.Fatalf("expected workspaceName : %s, but found : %s", expectedworkspaceName, workspaceName)
	}

}

func TestAPIModelWithContainerMonitoringAddonWithConfigInCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := ExampleAPIModelWithContainerMonitoringAddonWithExistingWorkspaceConfig

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

	k8sConfig := deployCmd.containerService.Properties.OrchestratorProfile.KubernetesConfig

	if k8sConfig == nil {
		t.Fatalf("expected valid kubernetes config")
	}

	if len(k8sConfig.Addons) != 1 {
		t.Fatalf("expected one addon")
	}

	addon := k8sConfig.Addons[0]

	expectedAddonName := "container-monitoring"
	if addon.Name != expectedAddonName {
		t.Fatalf("expected addon name: %s but got: %s", expectedAddonName, addon.Name)
	}

	expectedWorkspaceGUIDInBase64 := "MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAw"
	if addon.Config["workspaceGuid"] != expectedWorkspaceGUIDInBase64 {
		t.Fatalf("expected workspaceGuid : %s but got : %s", expectedWorkspaceGUIDInBase64, addon.Config["workspaceGuid"])
	}

	expectedWorkspaceKeyInBase64 := "NEQrdnlkNS9qU2NCbXNBd1pPRi8wR09CUTVrdUZRYzlKVmFXK0hsbko1OGN5ZVBKY3dUcGtzK3JWbXZnY1hHbW15dWpMRE5FVlBpVDhwQjI3NGE5WWc9PQ=="
	if addon.Config["workspaceKey"] != expectedWorkspaceKeyInBase64 {
		t.Fatalf("unexpected workspaceKey : %s", addon.Config["workspaceKey"])
		t.Fatalf("expected workspaceKey : %s but got : %s", expectedWorkspaceKeyInBase64, addon.Config["workspaceKey"])
	}

	workspaceResourceID := addon.Config["logAnalyticsWorkspaceResourceId"]
	resourceParts := strings.Split(workspaceResourceID, "/")

	if len(resourceParts) != 9 {
		t.Fatalf("invalid workspaceResourceID : %s", workspaceResourceID)
	}

	workspaceSubscriptionID := resourceParts[2]
	workspaceResourceGroup := resourceParts[4]
	workspaceProvider := resourceParts[6]
	workspaceName := resourceParts[8]

	expectedworkspaceSubscriptionID := "00000000-0000-0000-0000-000000000000"
	if workspaceSubscriptionID != expectedworkspaceSubscriptionID {
		t.Fatalf("expected workspaceSubscriptionID: %s, but found : %s", expectedworkspaceSubscriptionID, workspaceSubscriptionID)
	}

	expectedworkspaceProvider := "Microsoft.OperationalInsights"
	if workspaceProvider != expectedworkspaceProvider {
		t.Fatalf("expected log analytics workspace provider name: %s, but got: %s", expectedworkspaceProvider, workspaceProvider)
	}

	expectedworkspaceResourceGroup := "test-workspace-rg"
	if workspaceResourceGroup != "test-workspace-rg" {
		t.Fatalf("expected workspaceResourceGroup : %s, but found : %s", expectedworkspaceResourceGroup, workspaceResourceGroup)
	}

	expectedworkspaceName := "test-workspace"
	if workspaceName != expectedworkspaceName {
		t.Fatalf("expected workspaceName : %s, but found : %s", expectedworkspaceName, workspaceName)
	}
}

func TestAPIModelWithContainerMonitoringAddonWithWorkspaceGuidAndKeyConfigInCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := ExampleAPIModelWithContainerMonitoringAddonWithWorkspaceGUIDAndKeyConfig

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

	k8sConfig := deployCmd.containerService.Properties.OrchestratorProfile.KubernetesConfig

	if k8sConfig == nil {
		t.Fatalf("expected valid kubernetes config")
	}

	if len(k8sConfig.Addons) != 1 {
		t.Fatalf("expected one addon")
	}

	addon := k8sConfig.Addons[0]

	if addon.Name != "container-monitoring" {
		t.Fatalf("unexpected addon found : %s", addon.Name)
	}

	expectedWorkspaceGUIDInBase64 := "MDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAw"
	if addon.Config["workspaceGuid"] != expectedWorkspaceGUIDInBase64 {
		t.Fatalf("expected workspaceGuid : %s but got : %s", expectedWorkspaceGUIDInBase64, addon.Config["workspaceGuid"])
	}

	expectedWorkspaceKeyInBase64 := "NEQrdnlkNS9qU2NCbXNBd1pPRi8wR09CUTVrdUZRYzlKVmFXK0hsbko1OGN5ZVBKY3dUcGtzK3JWbXZnY1hHbW15dWpMRE5FVlBpVDhwQjI3NGE5WWc9PQ=="
	if addon.Config["workspaceKey"] != expectedWorkspaceKeyInBase64 {
		t.Fatalf("unexpected workspaceKey : %s", addon.Config["workspaceKey"])
		t.Fatalf("expected workspaceKey : %s but got : %s", expectedWorkspaceKeyInBase64, addon.Config["workspaceKey"])
	}
}
