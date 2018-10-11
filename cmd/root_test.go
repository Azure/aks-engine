package cmd

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	ini "gopkg.in/ini.v1"
)

func TestNewRootCmd(t *testing.T) {
	output := NewRootCmd()
	if output.Use != rootName || output.Short != rootShortDescription || output.Long != rootLongDescription {
		t.Fatalf("root command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, rootName, output.Short, rootShortDescription, output.Long, rootLongDescription)
	}
	expectedCommands := []*cobra.Command{getCompletionCmd(output), newDeployCmd(), newGenerateCmd(), newOrchestratorsCmd(), newScaleCmd(), newUpgradeCmd(), newVersionCmd()}
	rc := output.Commands()
	for i, c := range expectedCommands {
		if rc[i].Use != c.Use {
			t.Fatalf("root command should have command %s", c.Use)
		}
	}
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
