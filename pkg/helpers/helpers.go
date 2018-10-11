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

const (
	// SSHKeySize is the size (in bytes) of SSH key to create
	SSHKeySize = 4096
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
	switch sku {
	case "Standard_D3_v2", "Standard_D12_v2", "Standard_D3_v2_Promo", "Standard_D12_v2_Promo",
		"Standard_DS3_v2", "Standard_DS12_v2", "Standard_DS13-4_v2", "Standard_DS14-4_v2",
		"Standard_DS3_v2_Promo", "Standard_DS12_v2_Promo", "Standard_DS13-4_v2_Promo",
		"Standard_DS14-4_v2_Promo", "Standard_F4", "Standard_F4s", "Standard_D8_v3", "Standard_D8s_v3",
		"Standard_D32-8s_v3", "Standard_E8_v3", "Standard_E8s_v3", "Standard_D3_v2_ABC",
		"Standard_D12_v2_ABC", "Standard_F4_ABC", "Standard_F8s_v2", "Standard_D4_v2",
		"Standard_D13_v2", "Standard_D4_v2_Promo", "Standard_D13_v2_Promo", "Standard_DS4_v2",
		"Standard_DS13_v2", "Standard_DS14-8_v2", "Standard_DS4_v2_Promo", "Standard_DS13_v2_Promo",
		"Standard_DS14-8_v2_Promo", "Standard_F8", "Standard_F8s", "Standard_M64-16ms", "Standard_D16_v3",
		"Standard_D16s_v3", "Standard_D32-16s_v3", "Standard_D64-16s_v3", "Standard_E16_v3",
		"Standard_E16s_v3", "Standard_E32-16s_v3", "Standard_D4_v2_ABC", "Standard_D13_v2_ABC",
		"Standard_F8_ABC", "Standard_F16s_v2", "Standard_D5_v2", "Standard_D14_v2", "Standard_D5_v2_Promo",
		"Standard_D14_v2_Promo", "Standard_DS5_v2", "Standard_DS14_v2", "Standard_DS5_v2_Promo",
		"Standard_DS14_v2_Promo", "Standard_F16", "Standard_F16s", "Standard_M64-32ms",
		"Standard_M128-32ms", "Standard_D32_v3", "Standard_D32s_v3", "Standard_D64-32s_v3",
		"Standard_E32_v3", "Standard_E32s_v3", "Standard_E32-8s_v3", "Standard_E32-16_v3",
		"Standard_D5_v2_ABC", "Standard_D14_v2_ABC", "Standard_F16_ABC", "Standard_F32s_v2",
		"Standard_D15_v2", "Standard_D15_v2_Promo", "Standard_D15_v2_Nested", "Standard_DS15_v2",
		"Standard_DS15_v2_Promo", "Standard_DS15_v2_Nested", "Standard_D40_v3", "Standard_D40s_v3",
		"Standard_D15_v2_ABC", "Standard_M64ms", "Standard_M64s", "Standard_M128-64ms",
		"Standard_D64_v3", "Standard_D64s_v3", "Standard_E64_v3", "Standard_E64s_v3", "Standard_E64-16s_v3",
		"Standard_E64-32s_v3", "Standard_F64s_v2", "Standard_F72s_v2", "Standard_M128s", "Standard_M128ms",
		"Standard_L8s_v2", "Standard_L16s_v2", "Standard_L32s_v2", "Standard_L64s_v2", "Standard_L96s_v2",
		"SQLGL", "SQLGLCore", "Standard_D4_v3", "Standard_D4s_v3", "Standard_D2_v2", "Standard_DS2_v2",
		"Standard_E4_v3", "Standard_E4s_v3", "Standard_F2", "Standard_F2s", "Standard_F4s_v2",
		"Standard_D11_v2", "Standard_DS11_v2", "AZAP_Performance_ComputeV17C", "Standard_PB6s",
		"Standard_PB12s", "Standard_PB24s":
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
//  Azure public cloud
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
