package helpers

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/Azure/aks-engine/pkg/i18n"
	log "github.com/sirupsen/logrus"
)

// FileSaver represents the object that save string or byte data to file
type FileSaver struct {
	Translator *i18n.Translator
}

// SaveFileString saves string to file
func (f *FileSaver) SaveFileString(dir string, file string, data string) error {
	return f.SaveFile(dir, file, []byte(data))
}

// SaveFile saves binary data to file
func (f *FileSaver) SaveFile(dir string, file string, data []byte) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if e := os.MkdirAll(dir, 0700); e != nil {
			return f.Translator.Errorf("error creating directory '%s': %s", dir, e.Error())
		}
	}

	path := path.Join(dir, file)
	if err := ioutil.WriteFile(path, []byte(data), 0600); err != nil {
		return err
	}

	log.Debugf("output: wrote %s", path)

	return nil
}
