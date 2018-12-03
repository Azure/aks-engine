// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package unstable

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/Azure/aks-engine/pkg/openshift/certgen/unstable/templates"
	"github.com/Azure/aks-engine/pkg/openshift/filesystem"
	"golang.org/x/crypto/bcrypt"
)

type modeinfo struct {
	path  *regexp.Regexp
	mode  os.FileMode
	user  string
	group string
}

// These ownerships are processed in order
var ownerships = []modeinfo{
	// etc/etcd directory
	{path: regexp.MustCompile(`^etc/etcd$`), mode: 0755, user: "etcd", group: "etcd"},
	// tmp directory
	{path: regexp.MustCompile(`^tmp$`), mode: 01777, user: "root", group: "root"},
	//start files
	{path: regexp.MustCompile(`^etc/etcd/.*\.key$`), mode: 0600, user: "etcd", group: "etcd"},
	{path: regexp.MustCompile(`^etc/etcd/.*$`), user: "etcd", group: "etcd"},
	{path: regexp.MustCompile(`.*\.key$`), mode: 0600},
	{path: regexp.MustCompile(`.*\.kubeconfig$`), mode: 0600},
	{path: regexp.MustCompile(`^etc/origin/master/htpasswd$`), mode: 0600},
}

// GetFileInfo returns the permissions and ownership of the file if defined
func GetFileInfo(filename string) filesystem.Fileinfo {
	// If filename matches a specific path then set the correct User, Group, and Mode
	f := filesystem.Fileinfo{User: "root", Group: "root", Mode: 0644}
	for _, owner := range ownerships {

		if owner.path.MatchString(filename) {
			if owner.user != "" {
				f.User = owner.user
			}
			if owner.group != "" {
				f.Group = owner.group
			}
			if owner.mode != 0 {
				f.Mode = owner.mode
			}
			break
		}
	}

	return f
}

// PrepareMasterFiles creates the shared authentication and encryption secrets
func (c *Config) PrepareMasterFiles() error {
	b := make([]byte, 24)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}
	c.AuthSecret = base64.StdEncoding.EncodeToString(b)

	_, err = rand.Read(b)
	if err != nil {
		return err
	}
	c.EncSecret = base64.StdEncoding.EncodeToString(b)

	return nil
}

// WriteMasterFiles writes the templated master config
func (c *Config) WriteMasterFiles(fs filesystem.Writer) error {

	// create special case directories
	specialCaseDirs := map[string]filesystem.Fileinfo{
		"tmp": {
			User:  "root",
			Group: "root",
			Mode:  os.FileMode(01777),
		},
		"etc/etcd": {
			Mode:  os.FileMode(0755),
			User:  "etcd",
			Group: "etcd",
		},
	}

	for na, fi := range specialCaseDirs {
		err := fs.Mkdir(na, fi)
		if err != nil {
			return err
		}
	}

	for _, name := range templates.AssetNames() {
		if !strings.HasPrefix(name, "master/") {
			continue
		}
		tb := templates.MustAsset(name)

		t, err := template.New("template").Funcs(template.FuncMap{
			"QuoteMeta": regexp.QuoteMeta,
			"Bcrypt": func(password string) (string, error) {
				h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				return string(h), err
			},
			"quote": strconv.Quote,
			"shellQuote": func(s string) string {
				return `'` + strings.Replace(s, `'`, `'\''`, -1) + `'`
			},
		}).Parse(string(tb))
		if err != nil {
			return err
		}

		b := &bytes.Buffer{}
		err = t.Execute(b, c)
		if err != nil {
			return err
		}

		fname := strings.TrimPrefix(name, "master/")
		fi := GetFileInfo(fname)

		err = fs.WriteFile(fname, b.Bytes(), fi)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteNodeFiles writes the templated node config
func (c *Config) WriteNodeFiles(fs filesystem.Writer) error {
	for _, name := range templates.AssetNames() {
		if !strings.HasPrefix(name, "node/") {
			continue
		}

		tb := templates.MustAsset(name)

		t, err := template.New("template").Funcs(template.FuncMap{
			"QuoteMeta": regexp.QuoteMeta,
			"quote":     strconv.Quote,
			"shellQuote": func(s string) string {
				return `'` + strings.Replace(s, `'`, `'\''`, -1) + `'`
			},
		}).Parse(string(tb))
		if err != nil {
			return err
		}

		b := &bytes.Buffer{}
		err = t.Execute(b, c)
		if err != nil {
			return err
		}

		fname := strings.TrimPrefix(name, "node/")
		fi := GetFileInfo(fname)

		err = fs.WriteFile(fname, b.Bytes(), fi)
		if err != nil {
			return err
		}
	}

	return nil
}
