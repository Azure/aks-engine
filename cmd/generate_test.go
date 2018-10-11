package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestNewGenerateCmd(t *testing.T) {
	output := newGenerateCmd()
	if output.Use != generateName || output.Short != generateShortDescription || output.Long != generateLongDescription {
		t.Fatalf("generate command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, generateName, output.Short, generateShortDescription, output.Long, generateLongDescription)
	}

	expectedFlags := []string{"api-model", "output-directory", "ca-certificate-path", "ca-private-key-path", "set", "no-pretty-print", "parameters-only"}
	for _, f := range expectedFlags {
		if output.Flags().Lookup(f) == nil {
			t.Fatalf("generate command should have flag %s", f)
		}
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
	err := g.loadAPIModel(r, []string{"../pkg/engine/testdata/simple/kubernetes.json"})
	if err != nil {
		t.Fatalf("unexpected error loading api model: %s", err.Error())
	}
}
