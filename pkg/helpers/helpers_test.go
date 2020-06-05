// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/pkg/errors"
)

type ContainerService struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
}

func TestJSONMarshal(t *testing.T) {
	input := &ContainerService{}
	result, _ := JSONMarshal(input, false)
	expected := "{\"id\":\"\",\"location\":\"\",\"name\":\"\"}\n"
	if string(result) != expected {
		t.Fatalf("JSONMarshal returned unexpected result: expected %s but got %s", expected, string(result))
	}
	result, _ = JSONMarshalIndent(input, "", "", false)
	expected = "{\n\"id\": \"\",\n\"location\": \"\",\n\"name\": \"\"\n}\n"
	if string(result) != expected {
		t.Fatalf("JSONMarshal returned unexpected result: expected \n%sbut got \n%s", expected, result)
	}
}

func TestNormalizeAzureRegion(t *testing.T) {
	cases := []struct {
		input          string
		expectedResult string
	}{
		{
			input:          "westus",
			expectedResult: "westus",
		},
		{
			input:          "West US",
			expectedResult: "westus",
		},
		{
			input:          "Eastern Africa",
			expectedResult: "easternafrica",
		},
		{
			input:          "",
			expectedResult: "",
		},
	}

	for _, c := range cases {
		result := NormalizeAzureRegion(c.input)
		if c.expectedResult != result {
			t.Fatalf("NormalizeAzureRegion returned unexpected result: expected %s but got %s", c.expectedResult, result)
		}
	}
}

func TestGetEnglishOrderedListWithOxfordCommas(t *testing.T) {
	cases := []struct {
		input          []string
		expectedResult string
	}{
		{
			input:          []string{"foo"},
			expectedResult: "\"foo\"",
		},
		{
			input:          []string{"foo", "bar"},
			expectedResult: "\"foo\" and \"bar\"",
		},
		{
			input:          []string{"foo", "bar", "baz"},
			expectedResult: "\"foo\", \"bar\", and \"baz\"",
		},
		{
			input:          []string{"thing 1", "thing 2"},
			expectedResult: "\"thing 1\" and \"thing 2\"",
		},
	}

	for _, c := range cases {
		result := GetEnglishOrderedQuotedListWithOxfordCommas(c.input)
		if c.expectedResult != result {
			t.Fatalf("GetEnglishOrderedQuotedListWithOxfordCommas returned unexpected result: expected %s but got %s", c.expectedResult, result)
		}
	}
}

func TestPointerToBool(t *testing.T) {
	boolVar := true
	ret := PointerToBool(boolVar)
	if !*ret {
		t.Fatal("expected PointerToBool(true) to return true")
	}

	if !IsTrueBoolPointer(ret) {
		t.Fatal("expected IsTrueBoolPointer(*true) to return true")
	}
	ret = to.BoolPtr(boolVar)
	if !*ret {
		t.Fatal("expected to.BoolPtr(true) to return true")
	}

	if !to.Bool(ret) {
		t.Fatal("expected to.Bool(*true) to return true")
	}

	boolVar = false
	ret = PointerToBool(boolVar)
	if *ret {
		t.Fatal("expected PointerToBool(false) to return false")
	}

	if !IsFalseBoolPointer(ret) {
		t.Fatal("expected IsFalseBoolPointer(*false) to return true")
	}

	ret = to.BoolPtr(boolVar)
	if *ret {
		t.Fatal("expected to.BoolPtr(false) to return false")
	}

	if to.Bool(ret) {
		t.Fatal("expected to.Bool(*true) to return false")
	}

	boolVar = true
	ret = PointerToBool(boolVar)
	if IsFalseBoolPointer(ret) {
		t.Fatal("expected IsFalseBoolPointer(*true) to return false")
	}

	if !to.Bool(ret) {
		t.Fatal("expected !to.Bool(*true) to return false")
	}
}

func TestPointerToInt(t *testing.T) {
	int1 := 1
	int2 := 2
	ret1 := PointerToInt(int1)
	if *ret1 != int1 {
		t.Fatalf("expected PointerToInt(1) to return *1, instead returned %#v", ret1)
	}
	ret2 := PointerToInt(int2)
	if *ret2 != int2 {
		t.Fatalf("expected PointerToInt(2) to return *2, instead returned %#v", ret2)
	}

	if *ret2 <= *ret1 {
		t.Fatalf("Pointers to ints messed up their values and made 2 <= 1")
	}
}

func TestPointerToString(t *testing.T) {
	str := "foobar"
	ret := PointerToString(str)

	if *ret != str {
		t.Fatalf("expected PointerToString(foobar) to return *foobar, instead returned %#v", ret)
	}
}

func TestCreateSSH(t *testing.T) {
	rg := rand.New(rand.NewSource(42))

	translator := &i18n.Translator{
		Locale: nil,
	}

	privateKey, publicKey, err := CreateSSH(rg, translator)
	if err != nil {
		t.Fatalf("failed to generate SSH: %s", err)
	}
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	pemBuffer := bytes.Buffer{}
	if err = pem.Encode(&pemBuffer, pemBlock); err != nil {
		t.Fatalf("failed to encode private key: %s", err)
	}

	if !strings.HasPrefix(pemBuffer.String(), "-----BEGIN RSA PRIVATE KEY-----") {
		t.Fatalf("Private Key did not start with expected header")
	}

	if privateKey.N.BitLen() != SSHKeySize {
		t.Fatalf("Private Key was of length %d but %d was expected", privateKey.N.BitLen(), SSHKeySize)
	}

	if err := privateKey.Validate(); err != nil {
		t.Fatalf("Private Key failed validation: %v", err)
	}

	if !strings.HasPrefix(publicKey, "ssh-rsa ") {
		t.Fatalf("Public Key did not start with expected header")
	}
}

func TestAcceleratedNetworkingSupported(t *testing.T) {
	cases := []struct {
		input          string
		expectedResult bool
	}{
		{
			input:          "Standard_A1",
			expectedResult: false,
		},
		{
			input:          "Standard_G4",
			expectedResult: false,
		},
		{
			input:          "Standard_B3",
			expectedResult: false,
		},
		{
			input:          "Standard_D1_v2",
			expectedResult: false,
		},
		{
			input:          "Standard_L3",
			expectedResult: false,
		},
		{
			input:          "Standard_NC6",
			expectedResult: false,
		},
		{
			input:          "Standard_G4",
			expectedResult: false,
		},
		{
			input:          "Standard_D2_v2",
			expectedResult: true,
		},
		{
			input:          "Standard_DS2_v2",
			expectedResult: true,
		},
		{
			input:          "Standard_DS3_v2",
			expectedResult: true,
		},
		{
			input:          "Standard_M8ms",
			expectedResult: true,
		},
		{
			input:          "AZAP_Performance_ComputeV17C",
			expectedResult: false,
		},
		{
			input:          "SQLGL",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: false,
		},
	}

	for _, c := range cases {
		result := AcceleratedNetworkingSupported(c.input)
		if c.expectedResult != result {
			t.Fatalf("AcceleratedNetworkingSupported returned unexpected result for %s: expected %t but got %t", c.input, c.expectedResult, result)
		}
	}
}

func TestEqualError(t *testing.T) {
	testcases := []struct {
		errA     error
		errB     error
		expected bool
	}{
		{
			errA:     nil,
			errB:     nil,
			expected: true,
		},
		{
			errA:     errors.New("sample error"),
			errB:     nil,
			expected: false,
		},
		{
			errA:     nil,
			errB:     errors.New("sample error"),
			expected: false,
		},
		{
			errA:     errors.New("sample error"),
			errB:     errors.New("sample error"),
			expected: true,
		},
		{
			errA:     errors.New("sample error 1"),
			errB:     errors.New("sample error 2"),
			expected: false,
		},
	}

	for _, test := range testcases {
		if EqualError(test.errA, test.errB) != test.expected {
			t.Errorf("expected EqualError to return %t for errors %s and %s", test.expected, test.errA, test.errB)
		}
	}
}

func TestShellQuote(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{
			"Hel'lo'p\"la`ygr'ound'",
			`'Hel'\''lo'\''p"la` + "`" + `ygr'\''ound'\'''`,
		},
		{
			`"PwEV@QG7/PYt"re9`,
			`'"PwEV@QG7/PYt"re9'`,
		},
		{
			"",
			"''",
		},
		{
			"plaintext1234",
			`'plaintext1234'`,
		},
		{
			"Hel'lo'p\"la`ygr'ound",
			`'Hel'\''lo'\''p"la` + "`" + `ygr'\''ound'`,
		},
		{
			`conse''cutive`,
			`'conse'\'''\''cutive'`,
		},
		{
			"conse\\\\cutive",
			`'conse\\cutive'`,
		},
		{
			"consec\"\"utive",
			`'consec""utive'`,
		},
		{
			`PwEV@QG7/PYt"re9`,
			`'PwEV@QG7/PYt"re9'`,
		},
		{
			"Lnsr@191",
			"'Lnsr@191'",
		},
		{
			"Jach#321",
			"'Jach#321'",
		},
		{
			"Bgmo%219",
			"'Bgmo%219'",
		},
		{
			"@#$%^&*-_!+=[]{}|\\:,.?/~\"();" + "`",
			`'@#$%^&*-_!+=[]{}|\:,.?/~"();` + "`'",
		},
	}

	for _, test := range testcases {
		actual := ShellQuote(test.input)

		if actual != test.expected {
			t.Errorf("expected shellQuote to return %s, but got %s", test.expected, actual)
		}

		if runtime.GOOS != "windows" {
			out, err := exec.Command("/bin/bash", "-c", "testvar="+actual+"; echo -n $testvar").Output()
			if err != nil {
				t.Errorf("unexpected error : %s", err.Error())
			}

			if string(out) != test.input {
				t.Errorf("failed in Bash output test. Expected %s but got %s", test, out)
			}
		}
	}
}

func TestCreateSaveSSH(t *testing.T) {
	translator := &i18n.Translator{
		Locale: nil,
	}
	username := "test_user"
	outputDirectory := "unit_tests"
	expectedFile := outputDirectory + "/" + username + "_rsa"

	defer os.Remove(expectedFile)

	_, _, err := CreateSaveSSH(username, outputDirectory, translator)

	if err != nil {
		t.Fatalf("Unexpected error creating and saving ssh key: %s", err)
	}

	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Fatalf("ssh file was not created")
	}
}
func TestGetCloudTargetEnv(t *testing.T) {
	testcases := []struct {
		location string
		expected string
	}{
		{
			"chinaeast",
			"AzureChinaCloud",
		},
		{
			"chinanorth",
			"AzureChinaCloud",
		},
		{
			"chinaeast",
			"AzureChinaCloud",
		},
		{
			"chinaeast2",
			"AzureChinaCloud",
		},
		{
			"chinanorth2",
			"AzureChinaCloud",
		},
		{
			"germanycentral",
			"AzureGermanCloud",
		},
		{
			"germanynortheast",
			"AzureGermanCloud",
		},
		{
			"usgov123",
			"AzureUSGovernmentCloud",
		},
		{
			"usdod-123",
			"AzureUSGovernmentCloud",
		},
		{
			"sampleinput",
			"AzurePublicCloud",
		},
	}

	for _, testcase := range testcases {
		actual := GetCloudTargetEnv(testcase.location)
		if testcase.expected != actual {
			t.Errorf("expected GetCloudTargetEnv to return %s, but got %s", testcase.expected, actual)
		}
	}
}
func TestGetTargetEnv(t *testing.T) {
	testcases := []struct {
		location   string
		clouldName string
		expected   string
	}{
		{
			"chinaeast",
			"",
			"AzureChinaCloud",
		},
		{
			"chinanorth",
			"",
			"AzureChinaCloud",
		},
		{
			"chinaeast",
			"",
			"AzureChinaCloud",
		},
		{
			"chinaeast2",
			"",
			"AzureChinaCloud",
		},
		{
			"chinanorth2",
			"",
			"AzureChinaCloud",
		},
		{
			"germanycentral",
			"",
			"AzureGermanCloud",
		},
		{
			"germanynortheast",
			"",
			"AzureGermanCloud",
		},
		{
			"usgov123",
			"",
			"AzureUSGovernmentCloud",
		},
		{
			"usdod-123",
			"",
			"AzureUSGovernmentCloud",
		},
		{
			"sampleinput",
			"",
			"AzurePublicCloud",
		},
		{
			"azurestacklocation",
			"azurestackcloud",
			"AzureStackCloud",
		},
		{
			"azurestacklocation",
			"AzureStackcloud",
			"AzureStackCloud",
		},
		{
			"customlocation",
			"customcloud",
			"AzureStackCloud",
		},
	}

	for _, testcase := range testcases {
		actual := GetTargetEnv(testcase.location, testcase.clouldName)
		if testcase.expected != actual {
			t.Errorf("expected GetCloudTargetEnv to return %s, but got %s", testcase.expected, actual)
		}
	}
}

func TestEnsureString(t *testing.T) {
	testcases := []struct {
		defaultstring string
		overwrite     string
		expected      string
	}{
		{
			"",
			"randomstring",
			"randomstring",
		},
		{
			"srcisnotempty",
			"randomstring2",
			"srcisnotempty",
		},
	}

	for _, testcase := range testcases {
		result := EnsureString(testcase.defaultstring, testcase.overwrite)
		if testcase.expected != result {
			t.Errorf("expected EnsureString to return %s, but got %s", testcase.expected, result)
		}
	}

}

func TestGetLogAnalyticsWorkspaceDomain(t *testing.T) {
	testcases := []struct {
		cloudOrDependenciesLocation string
		expected                    string
	}{
		{
			"AzurePublicCloud",
			"opinsights.azure.com",
		},
		{
			"public",
			"opinsights.azure.com",
		},
		{
			"AzureChinaCloud",
			"opinsights.azure.cn",
		},
		{
			"china",
			"opinsights.azure.cn",
		},
		{
			"AzureUSGovernmentCloud",
			"opinsights.azure.us",
		},
		{
			"usgovernment",
			"opinsights.azure.us",
		},
		{
			"AzureGermanCloud",
			"opinsights.azure.de",
		},
		{
			"german",
			"opinsights.azure.de",
		},
		{
			"",
			"opinsights.azure.com",
		},
	}

	for _, testcase := range testcases {
		actual := GetLogAnalyticsWorkspaceDomain(testcase.cloudOrDependenciesLocation)
		if testcase.expected != actual {
			t.Errorf("expected GetCloudTargetEnv to return %s, but got %s", testcase.expected, actual)
		}
	}
}
