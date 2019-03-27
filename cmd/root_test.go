// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
	ini "gopkg.in/ini.v1"
)

func TestNewRootCmd(t *testing.T) {
	command := NewRootCmd()
	if command.Use != rootName || command.Short != rootShortDescription || command.Long != rootLongDescription {
		t.Fatalf("root command should have use %s equal %s, short %s equal %s and long %s equal to %s", command.Use, rootName, command.Short, rootShortDescription, command.Long, rootLongDescription)
	}
	expectedCommands := []*cobra.Command{getCompletionCmd(command), newDeployCmd(), newGenerateCmd(), newGetVersionsCmd(), newOrchestratorsCmd(), newScaleCmd(), newUpgradeCmd(), newVersionCmd()}
	rc := command.Commands()
	for i, c := range expectedCommands {
		if rc[i].Use != c.Use {
			t.Fatalf("root command should have command %s", c.Use)
		}
	}

	command.SetArgs([]string{"--debug"})
	err := command.Execute()
	if err != nil {
		t.Fatal(err)
	}
}

func TestShowDefaultModelArg(t *testing.T) {
	command := NewRootCmd()
	command.SetArgs([]string{"--show-default-model"})
	err := command.Execute()
	if err != nil {
		t.Fatal(err)
	}
	// TODO: examine command output
}

func TestDebugArg(t *testing.T) {
	command := NewRootCmd()
	command.SetArgs([]string{"--show-default-model"})
	err := command.Execute()
	if err != nil {
		t.Fatal(err)
	}
	// TODO: examine command output
}

func TestCompletionCommand(t *testing.T) {
	command := getCompletionCmd(NewRootCmd())
	command.SetArgs([]string{})
	err := command.Execute()
	if err != nil {
		t.Fatal(err)
	}
	// TODO: examine command output
}

func TestGetSelectedCloudFromAzConfig(t *testing.T) {
	for _, test := range []struct {
		desc   string
		data   []byte
		expect string
	}{
		{"nil file", nil, "AzureCloud"},
		{"empty file", []byte{}, "AzureCloud"},
		{"no cloud section", []byte(`
		[key]
		foo = bar
		`), "AzureCloud"},
		{"cloud section empty", []byte(`
		[cloud]
		[foo]
		foo = bar
		`), "AzureCloud"},
		{"AzureCloud selected", []byte(`
		[cloud]
		name = AzureCloud
		`), "AzureCloud"},
		{"custom cloud", []byte(`
		[cloud]
		name = myCloud
		`), "myCloud"},
	} {
		t.Run(test.desc, func(t *testing.T) {
			f, err := ini.Load(test.data)
			if err != nil {
				t.Fatal(err)
			}

			cloud := getSelectedCloudFromAzConfig(f)
			if cloud != test.expect {
				t.Fatalf("exepcted %q, got %q", test.expect, cloud)
			}
		})
	}
}

func TestGetCloudSubFromAzConfig(t *testing.T) {
	goodUUID, err := uuid.FromString("ccabad21-ea42-4ea1-affc-17ae73f9df66")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		desc   string
		data   []byte
		expect uuid.UUID
		err    bool
	}{
		{"empty file", []byte{}, uuid.UUID{}, true},
		{"no entry for cloud", []byte(`
		[SomeCloud]
		subscription = 00000000-0000-0000-0000-000000000000
		`), uuid.UUID{}, true},
		{"invalid UUID", []byte(`
		[AzureCloud]
		subscription = not-a-good-value
		`), uuid.UUID{}, true},
		{"real UUID", []byte(`
		[AzureCloud]
		subscription = ` + goodUUID.String() + `
		`), goodUUID, false},
	} {
		t.Run(test.desc, func(t *testing.T) {
			f, err := ini.Load(test.data)
			if err != nil {
				t.Fatal(err)
			}

			uuid, err := getCloudSubFromAzConfig("AzureCloud", f)
			if test.err != (err != nil) {
				t.Fatalf("expected err=%v, got: %v", test.err, err)
			}
			if test.err {
				return
			}
			if uuid.String() != test.expect.String() {
				t.Fatalf("expected %s, got %s", test.expect, uuid)
			}
		})
	}
}

func TestWriteCustomCloudProfile(t *testing.T) {
	const (
		name                         = "azurestackcloud"
		managementPortalURL          = "https://management.local.azurestack.external/"
		publishSettingsURL           = "https://management.local.azurestack.external/publishsettings/index"
		serviceManagementEndpoint    = "https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189"
		resourceManagerEndpoint      = "https://management.local.azurestack.external/"
		activeDirectoryEndpoint      = "https://login.windows.net/"
		galleryEndpoint              = "https://portal.local.azurestack.external=30015/"
		keyVaultEndpoint             = "https://vault.azurestack.external/"
		graphEndpoint                = "https://graph.windows.net/"
		serviceBusEndpoint           = "https://servicebus.azurestack.external/"
		batchManagementEndpoint      = "https://batch.azurestack.external/"
		storageEndpointSuffix        = "core.azurestack.external"
		sqlDatabaseDNSSuffix         = "database.azurestack.external"
		trafficManagerDNSSuffix      = "trafficmanager.cn"
		keyVaultDNSSuffix            = "vault.azurestack.external"
		serviceBusEndpointSuffix     = "servicebus.azurestack.external"
		serviceManagementVMDNSSuffix = "chinacloudapp.cn"
		resourceManagerVMDNSSuffix   = "cloudapp.azurestack.external"
		containerRegistryDNSSuffix   = "azurecr.io"
		tokenAudience                = "https://management.azurestack.external/"
	)
	cs := &api.ContainerService{
		Properties: &api.Properties{
			ServicePrincipalProfile: &api.ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &api.MasterProfile{
				Count:     1,
				DNSPrefix: "blueorange",
				VMSize:    "Standard_D2_v2",
			},
			OrchestratorProfile: &api.OrchestratorProfile{
				OrchestratorType: api.Kubernetes,
			},
			LinuxProfile: &api.LinuxProfile{},
			CustomCloudProfile: &api.CustomCloudProfile{
				IdentitySystem:       api.AzureADIdentitySystem,
				AuthenticationMethod: api.ClientSecretAuthMethod,
				Environment: &azure.Environment{
					Name:                         name,
					ManagementPortalURL:          managementPortalURL,
					PublishSettingsURL:           publishSettingsURL,
					ServiceManagementEndpoint:    serviceManagementEndpoint,
					ResourceManagerEndpoint:      resourceManagerEndpoint,
					ActiveDirectoryEndpoint:      activeDirectoryEndpoint,
					GalleryEndpoint:              galleryEndpoint,
					KeyVaultEndpoint:             keyVaultEndpoint,
					GraphEndpoint:                graphEndpoint,
					ServiceBusEndpoint:           serviceBusEndpoint,
					BatchManagementEndpoint:      batchManagementEndpoint,
					StorageEndpointSuffix:        storageEndpointSuffix,
					SQLDatabaseDNSSuffix:         sqlDatabaseDNSSuffix,
					TrafficManagerDNSSuffix:      trafficManagerDNSSuffix,
					KeyVaultDNSSuffix:            keyVaultDNSSuffix,
					ServiceBusEndpointSuffix:     serviceBusEndpointSuffix,
					ServiceManagementVMDNSSuffix: serviceManagementVMDNSSuffix,
					ResourceManagerVMDNSSuffix:   resourceManagerVMDNSSuffix,
					ContainerRegistryDNSSuffix:   containerRegistryDNSSuffix,
					TokenAudience:                tokenAudience,
				},
			},
			AgentPoolProfiles: []*api.AgentPoolProfile{
				{
					Name:   "agentpool1",
					VMSize: "Standard_D2_v2",
					Count:  2,
				},
			},
		},
	}

	cs.SetPropertiesDefaults(false, false)

	if err := writeCustomCloudProfile(cs); err != nil {
		t.Fatalf("failed to write custom cloud profile: err - %s", err)
	}

	environmentFilePath := os.Getenv("AZURE_ENVIRONMENT_FILEPATH")
	if environmentFilePath == "" {
		t.Fatal("failed to write custom cloud profile: err - AZURE_ENVIRONMENT_FILEPATH is empty")
	}

	if _, err := os.Stat(environmentFilePath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		t.Fatalf("failed to write custom cloud profile: file %s does not exist", environmentFilePath)
	}

	azurestackenvironment, err := ioutil.ReadFile(environmentFilePath)
	if err != nil {
		t.Fatalf("failed to write custom cloud profile: can not read file %s ", environmentFilePath)
	}
	azurestackenvironmentStr := string(azurestackenvironment)
	expectedResult := `{"name":"azurestackcloud","managementPortalURL":"https://management.local.azurestack.external/","publishSettingsURL":"https://management.local.azurestack.external/publishsettings/index","serviceManagementEndpoint":"https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189","resourceManagerEndpoint":"https://management.local.azurestack.external/","activeDirectoryEndpoint":"https://login.windows.net/","galleryEndpoint":"https://portal.local.azurestack.external=30015/","keyVaultEndpoint":"https://vault.azurestack.external/","graphEndpoint":"https://graph.windows.net/","serviceBusEndpoint":"https://servicebus.azurestack.external/","batchManagementEndpoint":"https://batch.azurestack.external/","storageEndpointSuffix":"core.azurestack.external","sqlDatabaseDNSSuffix":"database.azurestack.external","trafficManagerDNSSuffix":"trafficmanager.cn","keyVaultDNSSuffix":"vault.azurestack.external","serviceBusEndpointSuffix":"servicebus.azurestack.external","serviceManagementVMDNSSuffix":"chinacloudapp.cn","resourceManagerVMDNSSuffix":"cloudapp.azurestack.external","containerRegistryDNSSuffix":"azurecr.io","tokenAudience":"https://management.azurestack.external/"}`
	if azurestackenvironmentStr != expectedResult {
		t.Fatalf("failed to write custom cloud profile: expected %s , got %s ", expectedResult, azurestackenvironmentStr)
	}
}
