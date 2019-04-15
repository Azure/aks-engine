// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Azure/aks-engine/pkg/i18n"
)

func TestFileSaver(t *testing.T) {

	tmpDir, err := ioutil.TempDir("", "filesavertest")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tmpDir)
	tmpFile, err := ioutil.TempFile("", "filesavertest")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpFile.Name())

	type SaveFileEntry struct {
		Dir  string
		File string
		Data string
	}
	saveFileEntries := []SaveFileEntry{
		SaveFileEntry{
			Dir:  "", // test that an empty dir doesn't error
			File: tmpFile.Name(),
			Data: "",
		},
		SaveFileEntry{
			Dir:  tmpDir,
			File: "ubernetes.json",
			Data: "",
		},
	}

	f := FileSaver{Translator: &i18n.Translator{}}

	for _, sfe := range saveFileEntries {
		err := f.SaveFileString(sfe.Dir, sfe.File, sfe.Data)
		if err != nil {
			t.Error(err)
		}
	}
}
