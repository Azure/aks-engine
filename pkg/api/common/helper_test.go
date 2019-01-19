// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

import (
	"testing"

	"github.com/pkg/errors"
)

func TestValidateDNSPrefix(t *testing.T) {
	cases := []struct {
		dnsPrefix   string
		expectedErr error
	}{
		{
			"validDnsPrefix",
			nil,
		},
		{
			"",
			errors.New("DNSPrefix '' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 0)"),
		},
		{
			"a",
			errors.New("DNSPrefix 'a' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 1)"),
		},
		{
			"1234",
			errors.New("DNSPrefix '1234' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 4)"),
		},
		{
			"verylongdnsprefixthatismorethan45characterslong",
			errors.New("DNSPrefix 'verylongdnsprefixthatismorethan45characterslong' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 47)"),
		},
		{
			"dnswith_special?char",
			errors.New("DNSPrefix 'dnswith_special?char' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 20)"),
		},
		{
			"myDNS-1234",
			nil,
		},
	}

	for _, c := range cases {
		err := ValidateDNSPrefix(c.dnsPrefix)
		if err != nil && c.expectedErr != nil {
			if err.Error() != c.expectedErr.Error() {
				t.Fatalf("expected validateDNSPrefix to return error %s, but instead got %s", c.expectedErr.Error(), err.Error())
			}
		} else {
			if c.expectedErr != nil {
				t.Fatalf("expected validateDNSPrefix to return error %s, but instead got no error", c.expectedErr.Error())
			} else if err != nil {
				t.Fatalf("expected validateDNSPrefix to return no error, but instead got %s", err.Error())
			}
		}
	}
}

func TestIsNvidiaEnabledSKU(t *testing.T) {
	cases := GetNSeriesVMCasesForTesting()

	for _, c := range cases {
		ret := IsNvidiaEnabledSKU(c.VMSKU)
		if ret != c.Expected {
			t.Fatalf("expected IsNvidiaEnabledSKU(%s) to return %t, but instead got %t", c.VMSKU, c.Expected, ret)
		}
	}
}

func getCSeriesVMCasesForTesting() []struct {
	VMSKU    string
	Expected bool
} {
	cases := []struct {
		VMSKU    string
		Expected bool
	}{
		{
			"Standard_DC2s",
			true,
		},
		{
			"Standard_DC4s",
			true,
		},
		{
			"Standard_D2_v2",
			false,
		},
		{
			"gobledygook",
			false,
		},
		{
			"",
			false,
		},
	}
	return cases
}

func TestIsSGXEnabledSKU(t *testing.T) {
	cases := getCSeriesVMCasesForTesting()

	for _, c := range cases {
		ret := IsSgxEnabledSKU(c.VMSKU)
		if ret != c.Expected {
			t.Fatalf("expected IsSgxEnabledSKU(%s) to return %t, but instead got %t", c.VMSKU, c.Expected, ret)
		}
	}
}
