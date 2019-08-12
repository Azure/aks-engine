// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"os"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func TestNewGenerateCmd(t *testing.T) {
	command := newGenerateCmd()
	if command.Use != generateName || command.Short != generateShortDescription || command.Long != generateLongDescription {
		t.Fatalf("generate command should have use %s equal %s, short %s equal %s and long %s equal to %s", command.Use, generateName, command.Short, generateShortDescription, command.Long, generateLongDescription)
	}

	expectedFlags := []string{"api-model", "output-directory", "ca-certificate-path", "ca-private-key-path", "set", "no-pretty-print", "parameters-only", "client-id", "client-secret"}
	for _, f := range expectedFlags {
		if command.Flags().Lookup(f) == nil {
			t.Fatalf("generate command should have flag %s", f)
		}
	}

	command.SetArgs([]string{})
	if err := command.Execute(); err == nil {
		t.Fatalf("expected an error when calling generate with no arguments")
	}
}

func TestGenerateCmdValidate(t *testing.T) {
	g := &generateCmd{}
	r := &cobra.Command{}

	// validate cmd with 1 arg
	err := g.validate(r, []string{"../pkg/engine/testdata/simple/kubernetes.json"})
	if err != nil {
		t.Fatalf("unexpected error validating 1 arg: %s", err.Error())
	}

	g = &generateCmd{}

	// validate cmd with 0 args
	err = g.validate(r, []string{})
	t.Logf(err.Error())
	if err == nil {
		t.Fatalf("expected error validating 0 args")
	}

	g = &generateCmd{}

	// validate cmd with more than 1 arg
	err = g.validate(r, []string{"../pkg/engine/testdata/simple/kubernetes.json", "arg1"})
	t.Logf(err.Error())
	if err == nil {
		t.Fatalf("expected error validating multiple args")
	}

}

func TestGenerateCmdMergeAPIModel(t *testing.T) {
	g := &generateCmd{}
	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	err := g.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with no --set flag defined: %s", err.Error())
	}

	g = &generateCmd{}
	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	g.set = []string{"masterProfile.count=3,linuxProfile.adminUsername=testuser"}
	err = g.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag: %s", err.Error())
	}

	g = &generateCmd{}
	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	g.set = []string{"masterProfile.count=3", "linuxProfile.adminUsername=testuser"}
	err = g.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with multiple --set flags: %s", err.Error())
	}

	g = &generateCmd{}
	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	g.set = []string{"agentPoolProfiles[0].count=1"}
	err = g.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
	}

	// test with an ssh key that contains == sign
	g = &generateCmd{}
	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	g.set = []string{"linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""}
	err = g.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
	}

	// test with simple quote
	g = &generateCmd{}
	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	g.set = []string{"servicePrincipalProfile.secret='=MsR0ck5!t='"}
	err = g.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
	}
}

func TestGenerateCmdMLoadAPIModel(t *testing.T) {
	g := &generateCmd{}
	r := &cobra.Command{}

	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	g.set = []string{"agentPoolProfiles[0].count=1"}

	g.validate(r, []string{"../pkg/engine/testdata/simple/kubernetes.json"})
	g.mergeAPIModel()
	err := g.loadAPIModel()
	if err != nil {
		t.Fatalf("unexpected error loading api model: %s", err.Error())
	}
}

func TestAPIModelWithoutServicePrincipalProfileAndClientIdAndSecretInGenerateCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(false)

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	cs.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData = "ssh test"

	clientID, _ := uuid.FromString("e810b868-afab-412d-98cc-ce7db5cc840b")
	clientSecret := "Test Client secret"
	generateCmd := &generateCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		containerService: cs,
		apiVersion:       ver,
	}
	err = generateCmd.autofillApimodel()
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(generateCmd.outputDirectory)

	if generateCmd.containerService.Properties.ServicePrincipalProfile == nil || generateCmd.containerService.Properties.ServicePrincipalProfile.ClientID == "" || generateCmd.containerService.Properties.ServicePrincipalProfile.Secret == "" {
		t.Fatalf("expected service principal profile to be populated from deployment command arguments")
	}

	if generateCmd.containerService.Properties.ServicePrincipalProfile.ClientID != clientID.String() {
		t.Fatalf("expected service principal profile client id to be %s but got %s", clientID.String(), generateCmd.containerService.Properties.ServicePrincipalProfile.ClientID)
	}

	if generateCmd.containerService.Properties.ServicePrincipalProfile.Secret != clientSecret {
		t.Fatalf("expected service principal profile client secret to be %s but got %s", clientSecret, generateCmd.containerService.Properties.ServicePrincipalProfile.Secret)
	}

	err = generateCmd.validateAPIModelAsVLabs()
	if err != nil {
		t.Fatalf("unexpected error validateAPIModelAsVLabs the example apimodel: %s", err)
	}
}

func TestAPIModelWithoutServicePrincipalProfileAndWithoutClientIdAndSecretInGenerateCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(false)

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	cs.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData = "ssh test"
	generateCmd := &generateCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		containerService: cs,
		apiVersion:       ver,
	}
	err = generateCmd.autofillApimodel()
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(generateCmd.outputDirectory)

	if generateCmd.containerService.Properties.ServicePrincipalProfile != nil {
		t.Fatalf("expected service principal profile to be nil for unmanaged identity, where client id and secret are not supplied in api model and deployment command")
	}

	err = generateCmd.validateAPIModelAsVLabs()
	expectedErr := errors.New("ServicePrincipalProfile must be specified with Orchestrator Kubernetes")

	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected validate generate command to return error %s, but instead got %s", expectedErr.Error(), err.Error())
	}
}

func TestAPIModelWithManagedIdentityWithoutServicePrincipalProfileAndClientIdAndSecretInGenerateCmd(t *testing.T) {
	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(true)

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	cs.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData = "ssh test"
	clientID, _ := uuid.FromString("e810b868-afab-412d-98cc-ce7db5cc840b")
	clientSecret := "Test Client secret"
	generateCmd := &generateCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  "_test_output",
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		containerService: cs,
		apiVersion:       ver,
	}
	err = generateCmd.autofillApimodel()
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	defer os.RemoveAll(generateCmd.outputDirectory)

	if generateCmd.containerService.Properties.ServicePrincipalProfile != nil {
		t.Fatalf("expected service principal profile to be nil for managed identity")
	}

	err = generateCmd.validateAPIModelAsVLabs()
	if err != nil {
		t.Fatalf("unexpected error validateAPIModelAsVLabs the example apimodel: %s", err)
	}
}
