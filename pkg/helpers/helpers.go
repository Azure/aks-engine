// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import (
	// "fmt"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/Azure/aks-engine/pkg/i18n"
	"golang.org/x/crypto/ssh"
)

// NormalizeAzureRegion returns a normalized Azure region with white spaces removed and converted to lower case
func NormalizeAzureRegion(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "", -1))
}

// JSONMarshalIndent marshals formatted JSON w/ optional SetEscapeHTML
func JSONMarshalIndent(content interface{}, prefix, indent string, escape bool) ([]byte, error) {
	b, err := JSONMarshal(content, escape)
	if err != nil {
		return nil, err
	}

	var bufIndent bytes.Buffer
	if err := json.Indent(&bufIndent, b, prefix, indent); err != nil {
		return nil, err
	}

	return bufIndent.Bytes(), nil
}

// JSONMarshal marshals JSON w/ optional SetEscapeHTML
func JSONMarshal(content interface{}, escape bool) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(escape)
	if err := enc.Encode(content); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// IsTrueBoolPointer is a simple boolean helper function for boolean pointers
func IsTrueBoolPointer(b *bool) bool {
	if b != nil && *b {
		return true
	}
	return false
}

// IsFalseBoolPointer is a simple boolean helper function for boolean pointers
func IsFalseBoolPointer(b *bool) bool {
	if b != nil && !*b {
		return true
	}
	return false
}

// PointerToBool returns a pointer to a bool
func PointerToBool(b bool) *bool {
	p := b
	return &p
}

// PointerToString returns a pointer to a string
func PointerToString(s string) *string {
	p := s
	return &p
}

// PointerToInt returns a pointer to a int
func PointerToInt(i int) *int {
	p := i
	return &p
}

// EqualError is a nil-safe method which reports whether errors a and b are considered equal.
// They're equal if both are nil, or both are not nil and a.Error() == b.Error().
func EqualError(a, b error) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Error() == b.Error()
}

// CreateSSH creates an SSH key pair.
func CreateSSH(rg io.Reader, s *i18n.Translator) (privateKey *rsa.PrivateKey, publicKeyString string, err error) {
	privateKey, err = rsa.GenerateKey(rg, SSHKeySize)
	if err != nil {
		return nil, "", s.Errorf("failed to generate private key for ssh: %q", err)
	}

	publicKey := privateKey.PublicKey
	sshPublicKey, err := ssh.NewPublicKey(&publicKey)
	if err != nil {
		return nil, "", s.Errorf("failed to create openssh public key string: %q", err)
	}
	authorizedKeyBytes := ssh.MarshalAuthorizedKey(sshPublicKey)
	authorizedKey := string(authorizedKeyBytes)

	return privateKey, authorizedKey, nil
}

// GetHomeDir attempts to get the home dir from env
func GetHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// ShellQuote returns a string that is enclosed within single quotes. If the string already has single quotes, they will be escaped.
func ShellQuote(s string) string {
	return `'` + strings.Replace(s, `'`, `'\''`, -1) + `'`
}

// CreateSaveSSH generates and stashes an SSH key pair.
func CreateSaveSSH(username, outputDirectory string, s *i18n.Translator) (privateKey *rsa.PrivateKey, publicKeyString string, err error) {
	privateKey, publicKeyString, err = CreateSSH(rand.Reader, s)
	if err != nil {
		return nil, "", err
	}

	privateKeyPem := privateKeyToPem(privateKey)

	f := &FileSaver{
		Translator: s,
	}

	err = f.SaveFile(outputDirectory, fmt.Sprintf("%s_rsa", username), privateKeyPem)
	if err != nil {
		return nil, "", err
	}

	return privateKey, publicKeyString, nil
}

// GetCloudTargetEnv determines and returns whether the region is a sovereign cloud which
// have their own data compliance regulations (China/Germany/USGov) or standard
// Azure public cloud
func GetCloudTargetEnv(location string) string {
	loc := strings.ToLower(strings.Join(strings.Fields(location), ""))
	switch {
	case loc == "chinaeast" || loc == "chinanorth" || loc == "chinaeast2" || loc == "chinaeast3" || loc == "chinanorth2" || loc == "chinanorth3":
		return "AzureChinaCloud"
	case loc == "germanynortheast" || loc == "germanycentral":
		return "AzureGermanCloud"
	case strings.HasPrefix(loc, "usgov") || strings.HasPrefix(loc, "usdod"):
		return "AzureUSGovernmentCloud"
	default:
		return "AzurePublicCloud"
	}
}

// GetTargetEnv determines and returns whether the region is a sovereign cloud which
// have their own data compliance regulations (China/Germany/USGov) or standard
// Azure public cloud
// CustomCloudName is name of environment if customCloudProfile is provided, it will be empty string if customCloudProfile is empty.
// Because customCloudProfile is empty for deployment for AzurePublicCloud, AzureChinaCloud,AzureGermanCloud,AzureUSGovernmentCloud,
// The customCloudName value will be empty string for those clouds
// This value is also used in azure.json file, which is used by Kubernetes to read various endpoints specific to cloud
func GetTargetEnv(location, customCloudName string) string {
	switch {
	// Kubernetes only understand "AzureStackCloud" env right now to read from custom endpoints:
	// https://github.com/Azure/go-autorest/blob/master/autorest/azure/environments.go
	// We should always return "AzureStackCloud" when CustomCloudProfile was used
	case customCloudName != "":
		return "AzureStackCloud"
	default:
		return GetCloudTargetEnv(location)
	}
}

// EnsureString returns an string for the default value.
// If val is not empty, return val
// If val is empty, return defaultVal
func EnsureString(val, defaultVal string) string {
	if val != "" {
		return val
	}
	return defaultVal
}

// GetLogAnalyticsWorkspaceDomain gets log analytics workspace based on the cloud or azure stack dependenciesLocation
func GetLogAnalyticsWorkspaceDomain(cloudOrDependenciesLocation string) string {
	var workspaceDomain string
	switch strings.ToLower(strings.TrimSpace(cloudOrDependenciesLocation)) {
	case "azurepubliccloud", "public":
		workspaceDomain = "opinsights.azure.com"
	case "azurechinacloud", "china":
		workspaceDomain = "opinsights.azure.cn"
	case "azureusgovernmentcloud", "usgovernment":
		workspaceDomain = "opinsights.azure.us"
	case "azuregermancloud", "german":
		workspaceDomain = "opinsights.azure.de"
	default:
		workspaceDomain = "opinsights.azure.com"
	}
	return workspaceDomain
}

// GetEnglishOrderedQuotedListWithOxfordCommas returns a string that you can use in an English statement to report a list of things
// With each thing in quotes, e.g. "thing 1", "thing 2", and "thing 3"
// Oxford Comma not optional!
func GetEnglishOrderedQuotedListWithOxfordCommas(l []string) string {
	var ret string
	for i, item := range l {
		ret += fmt.Sprintf("\"%s\"", item)
		switch i {
		case len(l) - 2:
			if len(l) > 2 {
				ret += ", and "
			} else {
				ret += " and "
			}
		case len(l) - 2:
			ret += ", and "
		case len(l) - 1:
			break
		default:
			ret += ", "
		}
	}
	return ret
}
