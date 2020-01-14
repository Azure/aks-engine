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

// AcceleratedNetworkingSupported check if the VmSKU support the Accelerated Networking
func AcceleratedNetworkingSupported(sku string) bool {
	// Trim the optional _Promo suffix.
	sku = strings.TrimSuffix(sku, "_Promo")
	switch sku {
	// Regenerate this list with the following `az` CLI command:
	// az vm list-skus -o json \
	//  --query '[? starts_with(name, `Standard`) && !ends_with(name, `Promo`)].{name: name, caps: capabilities[? name==`AcceleratedNetworkingEnabled`]}[].{sku: name, acceleratedNetworking: caps[0].value}[? acceleratedNetworking == `True`].sku'
	// Then munge the string in Python with something like:
	// for i, sku in enumerate(sorted(set(skus))):
	//     print('"{}", '.format(sku), end=('\n' if i % 5 == 4 else ''))
	case "Standard_B12ms", "Standard_B16ms", "Standard_B20ms", "Standard_D11_v2", "Standard_D12_v2",
		"Standard_D13_v2", "Standard_D14_v2", "Standard_D15_v2", "Standard_D16_v3", "Standard_D16a_v4",
		"Standard_D16as_v4", "Standard_D16s_v3", "Standard_D2_v2", "Standard_D32_v3", "Standard_D32a_v4",
		"Standard_D32as_v4", "Standard_D32s_v3", "Standard_D3_v2", "Standard_D48_v3", "Standard_D48s_v3",
		"Standard_D4_v2", "Standard_D4_v3", "Standard_D4a_v4", "Standard_D4as_v4", "Standard_D4s_v3",
		"Standard_D5_v2", "Standard_D64_v3", "Standard_D64s_v3", "Standard_D8_v3", "Standard_D8a_v4",
		"Standard_D8as_v4", "Standard_D8s_v3", "Standard_DS11-1_v2", "Standard_DS11_v2", "Standard_DS12-1_v2",
		"Standard_DS12-2_v2", "Standard_DS12_v2", "Standard_DS13-2_v2", "Standard_DS13-4_v2", "Standard_DS13_v2",
		"Standard_DS14-4_v2", "Standard_DS14-8_v2", "Standard_DS14_v2", "Standard_DS15_v2", "Standard_DS2_v2",
		"Standard_DS3_v2", "Standard_DS4_v2", "Standard_DS5_v2", "Standard_E16-4s_v3", "Standard_E16-8s_v3",
		"Standard_E16_v3", "Standard_E16a_v4", "Standard_E16as_v4", "Standard_E16s_v3", "Standard_E20_v3",
		"Standard_E20a_v4", "Standard_E20as_v4", "Standard_E20s_v3", "Standard_E32-16s_v3", "Standard_E32-8s_v3",
		"Standard_E32_v3", "Standard_E32a_v4", "Standard_E32as_v4", "Standard_E32s_v3", "Standard_E4-2s_v3",
		"Standard_E48_v3", "Standard_E48as_v4", "Standard_E48s_v3", "Standard_E4_v3", "Standard_E4a_v4",
		"Standard_E4as_v4", "Standard_E4s_v3", "Standard_E64-16s_v3", "Standard_E64-32s_v3", "Standard_E64_v3",
		"Standard_E64as_v4", "Standard_E64i_v3", "Standard_E64is_v3", "Standard_E64s_v3", "Standard_E8-2s_v3",
		"Standard_E8-4s_v3", "Standard_E8_v3", "Standard_E8a_v4", "Standard_E8as_v4", "Standard_E8s_v3",
		"Standard_E96as_v4", "Standard_F16", "Standard_F16s", "Standard_F16s_v2", "Standard_F2",
		"Standard_F2s", "Standard_F32s_v2", "Standard_F4", "Standard_F48s_v2", "Standard_F4s",
		"Standard_F4s_v2", "Standard_F64s_v2", "Standard_F72s_v2", "Standard_F8", "Standard_F8s",
		"Standard_F8s_v2", "Standard_L16s_v2", "Standard_L32s_v2", "Standard_L48s_v2", "Standard_L64s_v2",
		"Standard_L80s_v2", "Standard_L8s_v2", "Standard_M128", "Standard_M128-32ms", "Standard_M128-64ms",
		"Standard_M128m", "Standard_M128ms", "Standard_M128s", "Standard_M16-4ms", "Standard_M16-8ms",
		"Standard_M16ms", "Standard_M208ms_v2", "Standard_M208s_v2", "Standard_M32-16ms", "Standard_M32-8ms",
		"Standard_M32ls", "Standard_M32ms", "Standard_M32ts", "Standard_M416ms_v2", "Standard_M416s_v2",
		"Standard_M64", "Standard_M64-16ms", "Standard_M64-32ms", "Standard_M64ls", "Standard_M64m",
		"Standard_M64ms", "Standard_M64s", "Standard_M8-2ms", "Standard_M8-4ms", "Standard_M8ms":
		return true
	// grandfathered SKUs with accelerated networking enabled
	case "AZAP_Performance_ComputeV17C", "SQLGL", "SQLGLCore", "Standard_D12_v2_ABC",
		"Standard_D13_v2_ABC", "Standard_D14_v2_ABC", "Standard_D15_v2_ABC", "Standard_D32-16s_v3",
		"Standard_D32-8s_v3", "Standard_D3_v2_ABC", "Standard_D40_v3", "Standard_D40s_v3",
		"Standard_D4_v2_ABC", "Standard_D5_v2_ABC", "Standard_D64-16s_v3", "Standard_D64-32s_v3",
		"Standard_E32-16_v3", "Standard_F16_ABC", "Standard_F4_ABC", "Standard_F8_ABC",
		"Standard_L96s_v2":
		return true
	default:
		return false
	}
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
	case loc == "chinaeast" || loc == "chinanorth" || loc == "chinaeast2" || loc == "chinanorth2":
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
func GetTargetEnv(location, customCloudName string) string {
	switch {
	case customCloudName != "" && strings.EqualFold(customCloudName, "AzureStackCloud"):
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

//get domain of azure log analytics workspace based on the cloud or azure stack dependenciesLocation
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
