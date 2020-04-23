// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// HandleValidationErrors is the helper function to catch validator.ValidationError
// based on Namespace of the error, and return customized error message.
func HandleValidationErrors(e validator.ValidationErrors) error {
	err := e[0]
	ns := err.Namespace()
	switch ns {
	case "Properties.OrchestratorProfile", "Properties.OrchestratorProfile.OrchestratorType",
		"Properties.MasterProfile", "Properties.MasterProfile.DNSPrefix", "Properties.MasterProfile.VMSize",
		"Properties.LinuxProfile", "Properties.ServicePrincipalProfile.ClientID",
		"Properties.WindowsProfile.AdminUsername",
		"Properties.WindowsProfile.AdminPassword":
		return errors.Errorf("missing %s", ns)
	case "Properties.MasterProfile.Count":
		return errors.New("MasterProfile count needs to be 1, 3, or 5")
	case "Properties.MasterProfile.OSDiskSizeGB":
		return errors.Errorf("Invalid os disk size of %d specified.  The range of valid values are [%d, %d]", err.Value().(int), MinDiskSizeGB, MaxDiskSizeGB)
	case "Properties.MasterProfile.IPAddressCount":
		return errors.Errorf("MasterProfile.IPAddressCount needs to be in the range [%d,%d]", MinIPAddressCount, MaxIPAddressCount)
	case "Properties.MasterProfile.StorageProfile":
		return errors.Errorf("Unknown storageProfile '%s'. Specify either %s or %s", err.Value().(string), StorageAccount, ManagedDisks)
	default:
		if strings.HasPrefix(ns, "Properties.AgentPoolProfiles") {
			switch {
			case strings.HasSuffix(ns, ".Name") || strings.HasSuffix(ns, "VMSize"):
				return errors.Errorf("missing %s", ns)
			case strings.HasSuffix(ns, ".Count"):
				return errors.Errorf("AgentPoolProfile count needs to be in the range [%d,%d]", MinAgentCount, MaxAgentCount)
			case strings.HasSuffix(ns, ".OSDiskSizeGB"):
				return errors.Errorf("Invalid os disk size of %d specified.  The range of valid values are [%d, %d]", err.Value().(int), MinDiskSizeGB, MaxDiskSizeGB)
			case strings.Contains(ns, ".Ports"):
				return errors.Errorf("AgentPoolProfile Ports must be in the range[%d, %d]", MinPort, MaxPort)
			case strings.HasSuffix(ns, ".StorageProfile"):
				return errors.Errorf("Unknown storageProfile '%s'. Specify %s, %s, or %s", err.Value().(string), StorageAccount, ManagedDisks, Ephemeral)
			case strings.Contains(ns, ".DiskSizesGB"):
				return errors.Errorf("A maximum of %d disks may be specified, The range of valid disk size values are [%d, %d]", MaxDisks, MinDiskSizeGB, MaxDiskSizeGB)
			case strings.HasSuffix(ns, ".IPAddressCount"):
				return errors.Errorf("AgentPoolProfile.IPAddressCount needs to be in the range [%d,%d]", MinIPAddressCount, MaxIPAddressCount)
			default:
				break
			}
		}
	}
	return errors.Errorf("Namespace %s is not caught, %+v", ns, e)
}

// ValidateDNSPrefix is a helper function to check that a DNS Prefix is valid
func ValidateDNSPrefix(dnsName string) error {
	dnsNameRegex := `^([A-Za-z][A-Za-z0-9-]{1,43}[A-Za-z0-9])$`
	re, err := regexp.Compile(dnsNameRegex)
	if err != nil {
		return err
	}
	if !re.MatchString(dnsName) {
		return errors.Errorf("DNSPrefix '%s' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was %d)", dnsName, len(dnsName))
	}
	return nil
}

// IsNvidiaEnabledSKU determines if an VM SKU has nvidia driver support
func IsNvidiaEnabledSKU(vmSize string) bool {
	/* If a new GPU sku becomes available, add a key to this map, but only if you have a confirmation
	   that we have an agreement with NVIDIA for this specific gpu.
	*/
	dm := map[string]bool{
		// K80
		"Standard_NC6":   true,
		"Standard_NC12":  true,
		"Standard_NC24":  true,
		"Standard_NC24r": true,
		// M60
		"Standard_NV6":      true,
		"Standard_NV12":     true,
		"Standard_NV12s_v3": true,
		"Standard_NV24":     true,
		"Standard_NV24s_v3": true,
		"Standard_NV24r":    true,
		"Standard_NV48s_v3": true,
		// P40
		"Standard_ND6s":   true,
		"Standard_ND12s":  true,
		"Standard_ND24s":  true,
		"Standard_ND24rs": true,
		// P100
		"Standard_NC6s_v2":   true,
		"Standard_NC12s_v2":  true,
		"Standard_NC24s_v2":  true,
		"Standard_NC24rs_v2": true,
		// V100
		"Standard_NC6s_v3":   true,
		"Standard_NC12s_v3":  true,
		"Standard_NC24s_v3":  true,
		"Standard_NC24rs_v3": true,
		"Standard_ND40s_v3":  true,
		"Standard_ND40rs_v2": true,
	}
	// Trim the optional _Promo suffix.
	vmSize = strings.TrimSuffix(vmSize, "_Promo")
	if _, ok := dm[vmSize]; ok {
		return dm[vmSize]
	}

	return false
}

// GetNSeriesVMCasesForTesting returns a struct w/ VM SKUs and whether or not we expect them to be nvidia-enabled
func GetNSeriesVMCasesForTesting() []struct {
	VMSKU    string
	Expected bool
} {
	cases := []struct {
		VMSKU    string
		Expected bool
	}{
		{
			"Standard_NC6",
			true,
		},
		{
			"Standard_NC6_Promo",
			true,
		},
		{
			"Standard_NC12",
			true,
		},
		{
			"Standard_NC24",
			true,
		},
		{
			"Standard_NC24r",
			true,
		},
		{
			"Standard_NV6",
			true,
		},
		{
			"Standard_NV12",
			true,
		},
		{
			"Standard_NV24",
			true,
		},
		{
			"Standard_NV24r",
			true,
		},
		{
			"Standard_ND6s",
			true,
		},
		{
			"Standard_ND12s",
			true,
		},
		{
			"Standard_ND24s",
			true,
		},
		{
			"Standard_ND24rs",
			true,
		},
		{
			"Standard_NC6s_v2",
			true,
		},
		{
			"Standard_NC12s_v2",
			true,
		},
		{
			"Standard_NC24s_v2",
			true,
		},
		{
			"Standard_NC24rs_v2",
			true,
		},
		{
			"Standard_NC24rs_v2",
			true,
		},
		{
			"Standard_NC6s_v3",
			true,
		},
		{
			"Standard_NC12s_v3",
			true,
		},
		{
			"Standard_NC24s_v3",
			true,
		},
		{
			"Standard_NC24rs_v3",
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

// GetDCSeriesVMCasesForTesting returns a struct w/ VM SKUs and whether or not we expect them to be SGX-enabled
func GetDCSeriesVMCasesForTesting() []struct {
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
			"Standard_NC12",
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

// IsSgxEnabledSKU determines if an VM SKU has SGX driver support
func IsSgxEnabledSKU(vmSize string) bool {
	switch vmSize {
	case "Standard_DC2s", "Standard_DC4s":
		return true
	}
	return false
}

// GetMasterKubernetesLabels returns a k8s API-compliant labels string.
// The `kubernetes.io/role` and `node-role.kubernetes.io` labels are disallowed
// by the kubelet `--node-labels` argument in Kubernetes 1.16 and later.
func GetMasterKubernetesLabels(rg string, deprecated bool) string {
	var buf bytes.Buffer
	buf.WriteString("kubernetes.azure.com/role=master")
	buf.WriteString(",node.kubernetes.io/exclude-from-external-load-balancers=true")
	buf.WriteString(",node.kubernetes.io/exclude-disruption=true")
	if deprecated {
		buf.WriteString(",kubernetes.io/role=master")
		buf.WriteString(",node-role.kubernetes.io/master=")
	}
	buf.WriteString(fmt.Sprintf(",kubernetes.azure.com/cluster=%s", rg))
	return buf.String()
}

// GetStorageAccountType returns the support managed disk storage tier for a give VM size
func GetStorageAccountType(sizeName string) (string, error) {
	spl := strings.Split(sizeName, "_")
	if len(spl) < 2 {
		return "", errors.Errorf("Invalid sizeName: %s", sizeName)
	}
	capability := spl[1]
	if strings.Contains(strings.ToLower(capability), "s") {
		return "Premium_LRS", nil
	}
	return "Standard_LRS", nil
}

// GetOrderedEscapedKeyValsString returns an ordered string of escaped, quoted key=val
func GetOrderedEscapedKeyValsString(config map[string]string) string {
	keys := []string{}
	for key := range config {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(fmt.Sprintf("\"%s=%s\", ", key, config[key]))
	}
	return strings.TrimSuffix(buf.String(), ", ")
}

// GetOrderedNewlinedKeyValsStringForCloudInit returns an ordered string of key = val, separated by newlines
func GetOrderedNewlinedKeyValsStringForCloudInit(config map[string]string) string {
	keys := []string{}
	for key := range config {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(fmt.Sprintf("%s = %s\n%4s", key, config[key], " "))
	}
	return strings.TrimSuffix(buf.String(), fmt.Sprintf("\n%4s", " "))
}

// SliceIntIsNonEmpty is a simple convenience to determine if a []int is non-empty
func SliceIntIsNonEmpty(s []int) bool {
	return len(s) > 0
}

// WrapAsARMVariable formats a string for inserting an ARM variable into an ARM expression
func WrapAsARMVariable(s string) string {
	return fmt.Sprintf("',variables('%s'),'", s)
}

// WrapAsParameter formats a string for inserting an ARM parameter into an ARM expression
func WrapAsParameter(s string) string {
	return fmt.Sprintf("',parameters('%s'),'", s)
}

// WrapAsVerbatim formats a string for inserting a literal string into an ARM expression
func WrapAsVerbatim(s string) string {
	return fmt.Sprintf("',%s,'", s)
}

// GetDockerConfig transforms the default docker config with overrides. Overrides may be nil.
func GetDockerConfig(opts map[string]string, overrides []func(*DockerConfig) error) (string, error) {
	config := GetDefaultDockerConfig()

	for i := range overrides {
		if err := overrides[i](&config); err != nil {
			return "", err
		}
	}

	dataDir, ok := opts[ContainerDataDirKey]
	if ok {
		config.DataRoot = dataDir
	}

	b, err := json.MarshalIndent(config, "", "    ")
	return string(b), err
}

// GetContainerdConfig transforms the default containerd config with overrides. Overrides may be nil.
func GetContainerdConfig(opts map[string]string, overrides []func(*ContainerdConfig) error) (string, error) {
	config := GetDefaultContainerdConfig()

	for i := range overrides {
		if err := overrides[i](&config); err != nil {
			return "", err
		}
	}

	dataDir, ok := opts[ContainerDataDirKey]
	if ok {
		config.Root = dataDir
	}

	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(config)
	return buf.String(), err
}

// ContainerdKubenetOverride transforms a containerd config to set details required when using kubenet.
func ContainerdKubenetOverride(config *ContainerdConfig) error {
	config.Plugins.IoContainerdGrpcV1Cri.CNI.ConfTemplate = "/etc/containerd/kubenet_template.conf"
	return nil
}

// ContainerdSandboxImageOverrider produces a function to transform containerd config by setting the SandboxImage.
func ContainerdSandboxImageOverrider(image string) func(*ContainerdConfig) error {
	return func(config *ContainerdConfig) error {
		config.Plugins.IoContainerdGrpcV1Cri.SandboxImage = image
		return nil
	}
}

// DockerNvidiaOverride transforms a docker config to supply nvidia runtime configuration.
func DockerNvidiaOverride(config *DockerConfig) error {
	if config.DockerDaemonRuntimes == nil {
		config.DockerDaemonRuntimes = make(map[string]DockerDaemonRuntime)
	}
	config.DefaultRuntime = "nvidia"
	config.DockerDaemonRuntimes["nvidia"] = DockerDaemonRuntime{
		Path:        "/usr/bin/nvidia-container-runtime",
		RuntimeArgs: []string{},
	}
	return nil
}

// IndentString pads each line of an original string with N spaces and returns the new value.
func IndentString(original string, spaces int) string {
	out := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(strings.NewReader(original))
	for scanner.Scan() {
		for i := 0; i < spaces; i++ {
			out.WriteString(" ")
		}
		out.WriteString(scanner.Text())
		out.WriteString("\n")
	}
	return out.String()
}

func GetDockerConfigTestCases() map[string]string {
	return map[string]string{
		"default": defaultDockerConfigString,
		"gpu":     dockerNvidiaConfigString,
		"reroot":  dockerRerootConfigString,
		"all":     dockerAllConfigString,
	}
}

func GetContainerdConfigTestCases() map[string]string {
	return map[string]string{
		"default": containerdImageConfigString,
		"kubenet": containerdImageKubenetConfigString,
		"reroot":  containerdImageRerootConfigString,
		"all":     containerdAllConfigString,
	}
}

var defaultContainerdConfigString = `oom_score = 0
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    [plugins."io.containerd.grpc.v1.cri".cni]
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
          runtime_type = "io.containerd.runc.v2"
`

var containerdRerootConfigString = `oom_score = 0
root = "/mnt/containerd"
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    [plugins."io.containerd.grpc.v1.cri".cni]
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
          runtime_type = "io.containerd.runc.v2"
`

var containerdKubenetConfigString = `oom_score = 0
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    [plugins."io.containerd.grpc.v1.cri".cni]
      conf_template = "/etc/containerd/kubenet_template.conf"
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
          runtime_type = "io.containerd.runc.v2"
`

var containerdImageConfigString = `oom_score = 0
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    sandbox_image = "foo/k8s/core/pause:1.2.0"
    [plugins."io.containerd.grpc.v1.cri".cni]
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
          runtime_type = "io.containerd.runc.v2"
`

var containerdImageRerootConfigString = `oom_score = 0
root = "/mnt/containerd"
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    sandbox_image = "foo/k8s/core/pause:1.2.0"
    [plugins."io.containerd.grpc.v1.cri".cni]
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
          runtime_type = "io.containerd.runc.v2"
`

var containerdImageKubenetConfigString = `oom_score = 0
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    sandbox_image = "foo/k8s/core/pause:1.2.0"
    [plugins."io.containerd.grpc.v1.cri".cni]
      conf_template = "/etc/containerd/kubenet_template.conf"
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
          runtime_type = "io.containerd.runc.v2"
`

var containerdAllConfigString = `oom_score = 0
root = "/mnt/containerd"
version = 2

[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    sandbox_image = "foo/k8s/core/pause:1.2.0"
    [plugins."io.containerd.grpc.v1.cri".cni]
      conf_template = "/etc/containerd/kubenet_template.conf"
    [plugins."io.containerd.grpc.v1.cri".containerd]
      default_runtime_name = "runc"
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.untrusted]
          runtime_type = "io.containerd.runc.v2"
`

var defaultDockerConfigString = `{
    "live-restore": true,
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "50m",
        "max-file": "5"
    }
}`

var dockerRerootConfigString = `{
    "data-root": "/mnt/docker",
    "live-restore": true,
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "50m",
        "max-file": "5"
    }
}`

var dockerNvidiaConfigString = `{
    "live-restore": true,
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "50m",
        "max-file": "5"
    },
    "default-runtime": "nvidia",
    "runtimes": {
        "nvidia": {
            "path": "/usr/bin/nvidia-container-runtime",
            "runtimeArgs": []
        }
    }
}`

var dockerAllConfigString = `{
    "data-root": "/mnt/docker",
    "live-restore": true,
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "50m",
        "max-file": "5"
    },
    "default-runtime": "nvidia",
    "runtimes": {
        "nvidia": {
            "path": "/usr/bin/nvidia-container-runtime",
            "runtimeArgs": []
        }
    }
}`
