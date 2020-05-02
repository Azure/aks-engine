// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
)

// CustomFileReader takes represents the source text of a file as an io.Reader and
// the desired destination to add it to
type CustomFileReader struct {
	Source io.Reader
	Dest   string
}

func masterCustomFiles(profile *api.Properties) []api.CustomFile {
	if profile.MasterProfile.CustomFiles != nil {
		return *profile.MasterProfile.CustomFiles
	}
	return []api.CustomFile{}
}

func customfilesIntoReaders(customFiles []api.CustomFile) ([]CustomFileReader, error) {
	customFileReaders := make([]CustomFileReader, len(customFiles))
	for idx, customFile := range customFiles {
		file, err := os.Open(customFile.Source)
		if err != nil {
			return []CustomFileReader{}, err
		}
		customFileReaders[idx] = CustomFileReader{
			Source: file,
			Dest:   customFile.Dest,
		}
	}
	return customFileReaders, nil
}

func substituteConfigStringCustomFiles(input string, customFiles []CustomFileReader, placeholder string) string {

	var config string
	for _, customFile := range customFiles {
		config += buildConfigStringCustomFiles(
			customFile.Source,
			customFile.Dest)

	}
	return strings.Replace(input, placeholder, config, -1)
}

func buildConfigStringCustomFiles(source io.Reader, destinationFile string) string {
	contents := []string{
		fmt.Sprintf("- path: %s", destinationFile),
		"  permissions: \\\"0644\\\"",
		"  encoding: gzip",
		"  owner: \\\"root\\\"",
		"  content: !!binary |",
		fmt.Sprintf("    %s\\n\\n", getBase64CustomFile(source)),
	}

	return strings.Join(contents, "\\n")
}

func getBase64CustomFile(source io.Reader) string {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(source)
	cfStr := buf.String()
	cfStr = strings.Replace(cfStr, "\r\n", "\n", -1)
	return getBase64EncodedGzippedCustomScriptFromStr(cfStr)
}
