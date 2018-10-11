package engine

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
)

func TestCustomFilesIntoReadersNonExistingFile(t *testing.T) {

	customFiles := []api.CustomFile{
		{
			Source: "no/path/doesnt/exist/nofile",
			Dest:   "/tmp/output",
		},
	}
	_, err := customfilesIntoReaders(customFiles)
	if err == nil {
		t.Fatalf("Error was not thrown when reading file in path: %s", customFiles[0].Source)
	}

}

//What the output should look like for a file with content "test"
var testFullStringSlice = []string{
	fmt.Sprintf("- path: %s", "/tmp/test"),
	"  permissions: \\\"0644\\\"",
	"  encoding: gzip",
	"  owner: \\\"root\\\"",
	"  content: !!binary |",
	fmt.Sprintf("    %s\\n\\n", "H4sIAAAAAAAA/ypJLS4BBAAA//8Mfn/YBAAAAA=="),
}

//What the output should look like for a file with content "filecontent"
var fileContentFullStringSlice = []string{
	fmt.Sprintf("- path: %s", "/tmp/test"),
	"  permissions: \\\"0644\\\"",
	"  encoding: gzip",
	"  owner: \\\"root\\\"",
	"  content: !!binary |",
	fmt.Sprintf("    %s\\n\\n", "H4sIAAAAAAAA/0rLzElNzs8rSc0rAQQAAP//lfHhvwsAAAA="),
}

func TestSubstituteConfigStringCustomFiles(t *testing.T) {
	//Set up string we are about to modify
	str := `
	some stuff

	MASTER_CUSTOM_FILES_PLACEHOLDER

	some more stuff
	`
	//Define the correct output string using the above defined slices
	preCorrectStr := `
	some stuff

	%s

	some more stuff
	`
	contents := fmt.Sprintf("%s%s", strings.Join(testFullStringSlice, "\\n"), strings.Join(fileContentFullStringSlice, "\\n"))
	correctStr := fmt.Sprintf(preCorrectStr, contents)

	//Add new readers with hard coded strings corresponding to correct output string
	customFilesReader := []CustomFileReader{
		{
			Source: strings.NewReader("test"),
			Dest:   "/tmp/test",
		},
		{
			Source: strings.NewReader("filecontent"),
			Dest:   "/tmp/test",
		},
	}

	str = substituteConfigStringCustomFiles(str,
		customFilesReader,
		"MASTER_CUSTOM_FILES_PLACEHOLDER")

	if str != correctStr {
		t.Fatalf("Parsed string was not correct from substituteConfigStringCustomFiles")
	}

}

func TestBuildConfigStringCustomFiles(t *testing.T) {
	configStrOutput := buildConfigStringCustomFiles(strings.NewReader("test"), "/tmp/test")
	correctOutput := strings.Join(testFullStringSlice, "\\n")
	if configStrOutput != correctOutput {
		t.Fatalf("Parsed string was not correct from buildConfigStringCustomFiles")
	}
}

func TestGetBase64CustomFile(t *testing.T) {
	b64outputStr := getBase64CustomFile(strings.NewReader("test"))
	correctOutput := "H4sIAAAAAAAA/ypJLS4BBAAA//8Mfn/YBAAAAA=="
	if b64outputStr != correctOutput {
		t.Fatalf("b64 encoded and zipped string: \"test\" from getBase64CustomFile is not correct ")
	}

}
