// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestValidateDNSPrefix(t *testing.T) {
	cases := []struct {
		name        string
		dnsPrefix   string
		expectedErr error
	}{
		{
			"valid DNS prefix",
			"validDnsPrefix",
			nil,
		},
		{
			"empty string",
			"",
			errors.New("DNSPrefix '' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 0)"),
		},
		{
			"one char",
			"a",
			errors.New("DNSPrefix 'a' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 1)"),
		},
		{
			"numbers",
			"1234",
			errors.New("DNSPrefix '1234' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 4)"),
		},
		{
			"too many chars",
			"verylongdnsprefixthatismorethan45characterslong",
			errors.New("DNSPrefix 'verylongdnsprefixthatismorethan45characterslong' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 47)"),
		},
		{
			"invalid special character",
			"dnswith_special?char",
			errors.New("DNSPrefix 'dnswith_special?char' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 20)"),
		},
		{
			"valid with numbers",
			"myDNS-1234",
			nil,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
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
		})
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
	name     string
	VMSKU    string
	Expected bool
} {
	cases := []struct {
		name     string
		VMSKU    string
		Expected bool
	}{
		{
			"Standard_DC2s",
			"Standard_DC2s",
			true,
		},
		{
			"Standard_DC4s",
			"Standard_DC4s",
			true,
		},
		{
			"Standard_D2_v2",
			"Standard_D2_v2",
			false,
		},
		{
			"gobledygook",
			"gobledygook",
			false,
		},
		{
			"empty string",
			"",
			false,
		},
	}
	return cases
}

func TestIsSGXEnabledSKU(t *testing.T) {
	cases := getCSeriesVMCasesForTesting()

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ret := IsSgxEnabledSKU(c.VMSKU)
			if ret != c.Expected {
				t.Fatalf("expected IsSgxEnabledSKU(%s) to return %t, but instead got %t", c.VMSKU, c.Expected, ret)
			}
		})
	}
}

func TestGetMasterKubernetesLabelsDeprecated(t *testing.T) {
	cases := []struct {
		name       string
		rg         string
		deprecated bool
		expected   string
	}{
		{
			"valid rg string",
			"my-resource-group",
			false,
			"kubernetes.azure.com/role=master,node.kubernetes.io/exclude-from-external-load-balancers=true,node.kubernetes.io/exclude-disruption=true,kubernetes.azure.com/cluster=my-resource-group",
		},
		{
			"valid rg string",
			"my-resource-group",
			true,
			"kubernetes.azure.com/role=master,node.kubernetes.io/exclude-from-external-load-balancers=true,node.kubernetes.io/exclude-disruption=true,kubernetes.io/role=master,node-role.kubernetes.io/master=,kubernetes.azure.com/cluster=my-resource-group",
		},
		{
			"empty string",
			"",
			false,
			"kubernetes.azure.com/role=master,node.kubernetes.io/exclude-from-external-load-balancers=true,node.kubernetes.io/exclude-disruption=true,kubernetes.azure.com/cluster=",
		},
		{
			"empty string",
			"",
			true,
			"kubernetes.azure.com/role=master,node.kubernetes.io/exclude-from-external-load-balancers=true,node.kubernetes.io/exclude-disruption=true,kubernetes.io/role=master,node-role.kubernetes.io/master=,kubernetes.azure.com/cluster=",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ret := GetMasterKubernetesLabels(c.rg, c.deprecated)
			if ret != c.expected {
				t.Fatalf("expected GetMasterKubernetesLabels(%s, %t) to return %s, but instead got %s", c.rg, c.deprecated, c.expected, ret)
			}
		})
	}
}

func TestGetOrderedEscapedKeyValsString(t *testing.T) {
	alphabetizedString := `"foo=bar", "yes=please"`
	cases := []struct {
		name     string
		input    map[string]string
		expected string
	}{
		{
			name:     "nil input",
			input:    map[string]string{},
			expected: "",
		},
		{
			name: "valid input",
			input: map[string]string{
				"foo": "bar",
				"yes": "please",
			},
			expected: alphabetizedString,
		},
		{
			name: "valid input re-ordered",
			input: map[string]string{
				"yes": "please",
				"foo": "bar",
			},
			expected: alphabetizedString,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ret := GetOrderedEscapedKeyValsString(c.input)
			if ret != c.expected {
				t.Fatalf("expected GetOrderedEscapedKeyValsString(%s) to return %s, but instead got %s", c.input, c.expected, ret)
			}
		})
	}
}

func TestGetOrderedNewlinedKeyValsStringForCloudInit(t *testing.T) {
	alphabetizedString := `foo = bar
    yes = please`
	cases := []struct {
		name     string
		input    map[string]string
		expected string
	}{
		{
			name:     "nil input",
			input:    map[string]string{},
			expected: "",
		},
		{
			name: "valid input",
			input: map[string]string{
				"foo": "bar",
				"yes": "please",
			},
			expected: alphabetizedString,
		},
		{
			name: "valid input re-ordered",
			input: map[string]string{
				"yes": "please",
				"foo": "bar",
			},
			expected: alphabetizedString,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ret := GetOrderedNewlinedKeyValsStringForCloudInit(c.input)
			if ret != c.expected {
				t.Fatalf("expected GetOrderedNewlinedKeyValsStringForCloudInit(%s) to return %s, but instead got %s", c.input, c.expected, ret)
			}
		})
	}
}

func TestGetStorageAccountType(t *testing.T) {
	validPremiumVMSize := "Standard_DS2_v2"
	validStandardVMSize := "Standard_D2_v2"
	expectedPremiumTier := "Premium_LRS"
	expectedStandardTier := "Standard_LRS"
	invalidVMSize := "D2v2"

	// test premium VMSize returns premium managed disk tier
	premiumTier, err := GetStorageAccountType(validPremiumVMSize)
	if err != nil {
		t.Fatalf("Invalid sizeName: %s", err)
	}

	if premiumTier != expectedPremiumTier {
		t.Fatalf("premium VM did no match premium managed storage tier")
	}

	// test standard VMSize returns standard managed disk tier
	standardTier, err := GetStorageAccountType(validStandardVMSize)
	if err != nil {
		t.Fatalf("Invalid sizeName: %s", err)
	}

	if standardTier != expectedStandardTier {
		t.Fatalf("standard VM did no match standard managed storage tier")
	}

	// test invalid VMSize
	result, err := GetStorageAccountType(invalidVMSize)
	if err == nil {
		t.Errorf("GetStorageAccountType() = (%s, nil), want error", result)
	}
}

func TestSliceIntIsNonEmpty(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name: "valid slice",
			input: []int{
				1, 2, 3,
			},
			expected: true,
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: false,
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ret := SliceIntIsNonEmpty(c.input)
			if ret != c.expected {
				t.Fatalf("expected SliceIntIsNonEmpty(%v) to return %t, but instead got %t", c.input, c.expected, ret)
			}
		})
	}
}

func TestWrapAsARMVariable(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "just a string",
			s:        "foo",
			expected: "',variables('foo'),'",
		},
		{
			name:     "empty string",
			s:        "",
			expected: "',variables(''),'",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := WrapAsARMVariable(test.s)
			if test.expected != ret {
				t.Errorf("expected %s, instead got : %s", test.expected, ret)
			}
		})
	}
}

func TestWrapAsParameter(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "just a string",
			s:        "foo",
			expected: "',parameters('foo'),'",
		},
		{
			name:     "empty string",
			s:        "",
			expected: "',parameters(''),'",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := WrapAsParameter(test.s)
			if test.expected != ret {
				t.Errorf("expected %s, instead got : %s", test.expected, ret)
			}
		})
	}
}

func TestWrapAsVerbatim(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "just a string",
			s:        "foo",
			expected: "',foo,'",
		},
		{
			name:     "empty string",
			s:        "",
			expected: "',,'",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ret := WrapAsVerbatim(test.s)
			if test.expected != ret {
				t.Errorf("expected %s, instead got : %s", test.expected, ret)
			}
		})
	}
}

func TestGetDockerConfig(t *testing.T) {
	tests := []struct {
		name      string
		options   map[string]string
		overrides []func(*DockerConfig) error
		fail      bool
		want      string
	}{
		{
			name: "docker default config",
			want: defaultDockerConfigString,
			fail: false,
		},
		{
			name: "docker reroot config",
			want: dockerRerootConfigString,
			fail: false,
			options: map[string]string{
				"dataDir": "/mnt/docker",
			},
		},
		{
			name: "docker nvidia config",
			want: dockerNvidiaConfigString,
			fail: false,
			overrides: []func(*DockerConfig) error{
				DockerNvidiaOverride,
			},
		},
		{
			name: "docker force error",
			want: "",
			fail: true,
			overrides: []func(*DockerConfig) error{
				func(_ *DockerConfig) error {
					return errors.New("foo")
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := GetDockerConfig(test.options, test.overrides)
			if err != nil && !test.fail {
				t.Fatalf("failed to get docker config: %v", err)
			}
			if test.fail {
				if err == nil {
					t.Fatalf("got docker config successfully while expecting failure")
				}
			} else {
				if err != nil {
					t.Fatalf("failed to get docker config: %v", err)
				}
				diff := cmp.Diff(test.want, got)
				if diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}

func TestGetContainerdConfig(t *testing.T) {
	tests := []struct {
		name      string
		options   map[string]string
		overrides []func(*ContainerdConfig) error
		fail      bool
		want      string
	}{
		{
			name: "container default config",
			want: defaultContainerdConfigString,
			fail: false,
		},

		{
			name: "container reroot config",
			want: containerdRerootConfigString,
			fail: false,
			options: map[string]string{
				"dataDir": "/mnt/containerd",
			},
		},
		{
			name: "container kubenet config",
			want: containerdKubenetConfigString,
			fail: false,
			overrides: []func(*ContainerdConfig) error{
				ContainerdKubenetOverride,
			},
		},
		{
			name: "container sandbox image config",
			want: containerdImageConfigString,
			fail: false,
			overrides: []func(*ContainerdConfig) error{
				ContainerdSandboxImageOverrider("foo/k8s/core/pause:1.2.0"),
			},
		},
		{
			name: "container force error",
			want: "",
			fail: true,
			overrides: []func(*ContainerdConfig) error{
				func(_ *ContainerdConfig) error {
					return errors.New("foo")
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := GetContainerdConfig(test.options, test.overrides)
			if err != nil && !test.fail {
				t.Fatalf("failed to get docker config: %v", err)
			}
			if test.fail {
				if err == nil {
					t.Fatalf("got docker config successfully while expecting failure")
				}
			} else {
				if err != nil {
					t.Fatalf("failed to get docker config: %v", err)
				}
				diff := cmp.Diff(test.want, got)
				if diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}

func TestIndentString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		count int
		want  string
	}{
		{
			name:  "should leave empty string alone",
			input: "",
			count: 4,
			want:  "",
		},
		{
			name:  "should indent single line string 4 spaces",
			input: "foo",
			count: 4,
			want:  "    foo\n",
		},
		{
			name:  "should indent multi-line string 4 spaces",
			input: "foo\nbar",
			count: 4,
			want:  "    foo\n    bar\n",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := IndentString(test.input, test.count)
			diff := cmp.Diff(test.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
