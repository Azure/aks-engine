// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

const defaultTestClusterVer = "1.7.12"

func TestAPIServerConfigEnableDataEncryptionAtRest(t *testing.T) {
	// Test EnableDataEncryptionAtRest = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableDataEncryptionAtRest = to.BoolPtr(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--encryption-provider-config"] != "/etc/kubernetes/encryption-config.yaml" {
		t.Fatalf("got unexpected '--encryption-provider-config' API server config value for EnableDataEncryptionAtRest=true: %s",
			a["--encryption-provider-config"])
	}

	// Test EnableDataEncryptionAtRest = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableDataEncryptionAtRest = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--encryption-provider-config"]; ok {
		t.Fatalf("got unexpected '--encryption-provider-config' API server config value for EnableDataEncryptionAtRest=false: %s",
			a["--encryption-provider-config"])
	}
}

func TestAPIServerConfigEnableEncryptionWithExternalKms(t *testing.T) {
	// Test EnableEncryptionWithExternalKms = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms = to.BoolPtr(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--encryption-provider-config"] != "/etc/kubernetes/encryption-config.yaml" {
		t.Fatalf("got unexpected '--encryption-provider-config' API server config value for EnableEncryptionWithExternalKms=true: %s",
			a["--encryption-provider-config"])
	}

	// Test EnableEncryptionWithExternalKms = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--encryption-provider-config"]; ok {
		t.Fatalf("got unexpected '--encryption-provider-config' API server config value for EnableEncryptionWithExternalKms=false: %s",
			a["--encryption-provider-config"])
	}
}

func TestAPIServerConfigEnableAggregatedAPIs(t *testing.T) {
	// Test EnableAggregatedAPIs = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs = true
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--requestheader-client-ca-file"] != "/etc/kubernetes/certs/proxy-ca.crt" {
		t.Fatalf("got unexpected '--requestheader-client-ca-file' API server config value for EnableAggregatedAPIs=true: %s",
			a["--requestheader-client-ca-file"])
	}
	if a["--proxy-client-cert-file"] != "/etc/kubernetes/certs/proxy.crt" {
		t.Fatalf("got unexpected '--proxy-client-cert-file' API server config value for EnableAggregatedAPIs=true: %s",
			a["--proxy-client-cert-file"])
	}
	if a["--proxy-client-key-file"] != "/etc/kubernetes/certs/proxy.key" {
		t.Fatalf("got unexpected '--proxy-client-key-file' API server config value for EnableAggregatedAPIs=true: %s",
			a["--proxy-client-key-file"])
	}
	if a["--requestheader-allowed-names"] != "" {
		t.Fatalf("got unexpected '--requestheader-allowed-names' API server config value for EnableAggregatedAPIs=true: %s",
			a["--requestheader-allowed-names"])
	}
	if a["--requestheader-extra-headers-prefix"] != "X-Remote-Extra-" {
		t.Fatalf("got unexpected '--requestheader-extra-headers-prefix' API server config value for EnableAggregatedAPIs=true: %s",
			a["--requestheader-extra-headers-prefix"])
	}
	if a["--requestheader-group-headers"] != "X-Remote-Group" {
		t.Fatalf("got unexpected '--requestheader-group-headers' API server config value for EnableAggregatedAPIs=true: %s",
			a["--requestheader-group-headers"])
	}
	if a["--requestheader-username-headers"] != "X-Remote-User" {
		t.Fatalf("got unexpected '--requestheader-username-headers' API server config value for EnableAggregatedAPIs=true: %s",
			a["--requestheader-username-headers"])
	}

	// Test EnableAggregatedAPIs = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs = false
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	for _, key := range []string{"--requestheader-client-ca-file", "--proxy-client-cert-file", "--proxy-client-key-file",
		"--requestheader-allowed-names", "--requestheader-extra-headers-prefix", "--requestheader-group-headers",
		"--requestheader-username-headers"} {
		if _, ok := a[key]; ok {
			t.Fatalf("got unexpected '%s' API server config value for EnableAggregatedAPIs=false: %s",
				key, a[key])
		}
	}
}

func TestAPIServerConfigUseCloudControllerManager(t *testing.T) {
	// Test UseCloudControllerManager = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = to.BoolPtr(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--cloud-provider"]; ok {
		t.Fatalf("got unexpected '--cloud-provider' API server config value for UseCloudControllerManager=false: %s",
			a["--cloud-provider"])
	}
	if _, ok := a["--cloud-config"]; ok {
		t.Fatalf("got unexpected '--cloud-config' API server config value for UseCloudControllerManager=false: %s",
			a["--cloud-config"])
	}

	// Test UseCloudControllerManager = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--cloud-provider"] != "azure" {
		t.Fatalf("got unexpected '--cloud-provider' API server config value for UseCloudControllerManager=true: %s",
			a["--cloud-provider"])
	}
	if a["--cloud-config"] != "/etc/kubernetes/azure.json" {
		t.Fatalf("got unexpected '--cloud-config' API server config value for UseCloudControllerManager=true: %s",
			a["--cloud-config"])
	}
}

func TestAPIServerConfigHasAadProfile(t *testing.T) {
	// Test HasAadProfile = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.AADProfile = &AADProfile{
		ServerAppID: "test-id",
		TenantID:    "test-tenant",
	}
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--oidc-username-claim"] != "oid" {
		t.Fatalf("got unexpected '--oidc-username-claim' API server config value for HasAadProfile=true: %s",
			a["--oidc-username-claim"])
	}
	if a["--oidc-groups-claim"] != "groups" {
		t.Fatalf("got unexpected '--oidc-groups-claim' API server config value for HasAadProfile=true: %s",
			a["--oidc-groups-claim"])
	}
	if a["--oidc-client-id"] != "spn:"+cs.Properties.AADProfile.ServerAppID {
		t.Fatalf("got unexpected '--oidc-client-id' API server config value for HasAadProfile=true: %s",
			a["--oidc-client-id"])
	}
	if a["--oidc-issuer-url"] != "https://sts.windows.net/"+cs.Properties.AADProfile.TenantID+"/" {
		t.Fatalf("got unexpected '--oidc-issuer-url' API server config value for HasAadProfile=true: %s",
			a["--oidc-issuer-url"])
	}

	// Test OIDC user overrides
	cs = CreateMockContainerService("testcluster", "1.7.12", 3, 2, false)
	cs.Properties.AADProfile = &AADProfile{
		ServerAppID: "test-id",
		TenantID:    "test-tenant",
	}
	usernameClaimOverride := "custom-username-claim"
	groupsClaimOverride := "custom-groups-claim"
	clientIDOverride := "custom-client-id"
	issuerURLOverride := "custom-issuer-url"
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{
		"--oidc-username-claim": usernameClaimOverride,
		"--oidc-groups-claim":   groupsClaimOverride,
		"--oidc-client-id":      clientIDOverride,
		"--oidc-issuer-url":     issuerURLOverride,
	}
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--oidc-username-claim"] != usernameClaimOverride {
		t.Fatalf("got unexpected '--oidc-username-claim' API server config value when user override provided: %s, expected: %s",
			a["--oidc-username-claim"], usernameClaimOverride)
	}
	if a["--oidc-groups-claim"] != groupsClaimOverride {
		t.Fatalf("got unexpected '--oidc-groups-claim' API server config value when user override provided: %s, expected: %s",
			a["--oidc-groups-claim"], groupsClaimOverride)
	}
	if a["--oidc-client-id"] != clientIDOverride {
		t.Fatalf("got unexpected '--oidc-client-id' API server config value when user override provided: %s, expected: %s",
			a["--oidc-client-id"], clientIDOverride)
	}
	if a["--oidc-issuer-url"] != issuerURLOverride {
		t.Fatalf("got unexpected '--oidc-issuer-url' API server config value when user override provided: %s, expected: %s",
			a["--oidc-issuer-url"], issuerURLOverride)
	}

	// Test China Cloud settings
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.AADProfile = &AADProfile{
		ServerAppID: "test-id",
		TenantID:    "test-tenant",
	}
	cs.Location = "chinaeast"
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--oidc-issuer-url"] != "https://sts.chinacloudapi.cn/"+cs.Properties.AADProfile.TenantID+"/" {
		t.Fatalf("got unexpected '--oidc-issuer-url' API server config value for HasAadProfile=true using China cloud: %s",
			a["--oidc-issuer-url"])
	}

	cs.Location = "chinaeast2"
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--oidc-issuer-url"] != "https://sts.chinacloudapi.cn/"+cs.Properties.AADProfile.TenantID+"/" {
		t.Fatalf("got unexpected '--oidc-issuer-url' API server config value for HasAadProfile=true using China cloud: %s",
			a["--oidc-issuer-url"])
	}

	cs.Location = "chinanorth"
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--oidc-issuer-url"] != "https://sts.chinacloudapi.cn/"+cs.Properties.AADProfile.TenantID+"/" {
		t.Fatalf("got unexpected '--oidc-issuer-url' API server config value for HasAadProfile=true using China cloud: %s",
			a["--oidc-issuer-url"])
	}

	cs.Location = "chinanorth2"
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--oidc-issuer-url"] != "https://sts.chinacloudapi.cn/"+cs.Properties.AADProfile.TenantID+"/" {
		t.Fatalf("got unexpected '--oidc-issuer-url' API server config value for HasAadProfile=true using China cloud: %s",
			a["--oidc-issuer-url"])
	}

	// Test HasAadProfile = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	for _, key := range []string{"--oidc-username-claim", "--oidc-groups-claim", "--oidc-client-id", "--oidc-issuer-url"} {
		if _, ok := a[key]; ok {
			t.Fatalf("got unexpected '%s' API server config value for HasAadProfile=false: %s",
				key, a[key])
		}
	}
}

func TestAPIServerConfigEnableRbac(t *testing.T) {
	// Test EnableRbac = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--authorization-mode"] != "Node,RBAC" {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for EnableRbac=true: %s",
			a["--authorization-mode"])
	}

	// Test EnableRbac = true with 1.6 cluster
	cs = CreateMockContainerService("testcluster", "1.6.11", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(true)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--authorization-mode"] != "RBAC" {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for 1.6 cluster with EnableRbac=true: %s",
			a["--authorization-mode"])
	}

	// Test EnableRbac = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--authorization-mode"]; ok {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for EnableRbac=false: %s",
			a["--authorization-mode"])
	}

	// Test EnableRbac = false with 1.6 cluster
	cs = CreateMockContainerService("testcluster", "1.6.11", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--authorization-mode"]; ok {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for 1.6 cluster with EnableRbac=false: %s",
			a["--authorization-mode"])
	}
}

func TestAPIServerConfigDisableRbac(t *testing.T) {
	// Test EnableRbac = false
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--authorization-mode"] != "" {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for EnableRbac=false: %s",
			a["--authorization-mode"])
	}
}

func TestAPIServerConfigEnableSecureKubelet(t *testing.T) {
	// Test EnableSecureKubelet = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = to.BoolPtr(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--kubelet-client-certificate"] != "/etc/kubernetes/certs/client.crt" {
		t.Fatalf("got unexpected '--kubelet-client-certificate' API server config value for EnableSecureKubelet=true: %s",
			a["--kubelet-client-certificate"])
	}
	if a["--kubelet-client-key"] != "/etc/kubernetes/certs/client.key" {
		t.Fatalf("got unexpected '--kubelet-client-key' API server config value for EnableSecureKubelet=true: %s",
			a["--kubelet-client-key"])
	}

	// Test EnableSecureKubelet = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	for _, key := range []string{"--kubelet-client-certificate", "--kubelet-client-key"} {
		if _, ok := a[key]; ok {
			t.Fatalf("got unexpected '%s' API server config value for EnableSecureKubelet=false: %s",
				key, a[key])
		}
	}
}

func TestAPIServerConfigDefaultAdmissionControls(t *testing.T) {
	version := "1.15.4"
	enableAdmissionPluginsKey := "--enable-admission-plugins"
	admissonControlKey := "--admission-control"
	cs := CreateMockContainerService("testcluster", version, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{}
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig[admissonControlKey] = "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,AlwaysPullImages,ExtendedResourceToleration"
	cs.Properties.OrchestratorProfile.KubernetesConfig.Addons = []KubernetesAddon{
		{
			Name:    common.PodSecurityPolicyAddonName,
			Enabled: to.BoolPtr(true),
		},
	}
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig

	if _, found := a[enableAdmissionPluginsKey]; !found {
		t.Fatalf("Admission control key '%s' not set in API server config for version %s", enableAdmissionPluginsKey, version)
	}

	// --admission-control was deprecated in v1.10
	if _, found := a[admissonControlKey]; found {
		t.Fatalf("Deprecated admission control key '%s' set in API server config for version %s", admissonControlKey, version)
	}

	// PodSecurityPolicy should be enabled in admission control
	admissionControlVal := a[enableAdmissionPluginsKey]
	if !strings.Contains(admissionControlVal, ",PodSecurityPolicy") {
		t.Fatalf("Admission control value '%s' expected to contain PodSecurityPolicy", admissionControlVal)
	}
}

func TestAPIServerConfigEnableProfiling(t *testing.T) {
	// Test
	// "apiServerConfig": {
	// 	"--profiling": "true"
	// },
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{
		"--profiling": "true",
	}
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--profiling"] != "true" {
		t.Fatalf("got unexpected '--profiling' API server config value for \"--profiling\": \"true\": %s",
			a["--profiling"])
	}

	// Test default
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--profiling"] != DefaultKubernetesAPIServerEnableProfiling {
		t.Fatalf("got unexpected default value for '--profiling' API server config: %s",
			a["--profiling"])
	}
}

func TestAPIServerConfigRepairMalformedUpdates(t *testing.T) {
	// Test default
	cs := CreateMockContainerService("testcluster", "1.13.0", 3, 2, false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--repair-malformed-updates"] != "false" {
		t.Fatalf("got unexpected default value for '--repair-malformed-updates' API server config: %s",
			a["--repair-malformed-updates"])
	}

	// Validate that 1.14.0 doesn't include --repair-malformed-updates at all
	cs = CreateMockContainerService("testcluster", "1.14.0", 3, 2, false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--repair-malformed-updates"]; ok {
		t.Fatalf("got a value for the deprecated '--repair-malformed-updates' API server config: %s",
			a["--repair-malformed-updates"])
	}
}

func TestAPIServerAuditPolicyBackCompatOverride(t *testing.T) {
	// Validate that we statically override "--audit-policy-file" values of "/etc/kubernetes/manifests/audit-policy.yaml" for back-compat
	auditPolicyKey := "--audit-policy-file"
	cs := CreateMockContainerService("testcluster", "1.10.8", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{}
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig[auditPolicyKey] = "/etc/kubernetes/manifests/audit-policy.yaml"
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a[auditPolicyKey] != "/etc/kubernetes/addons/audit-policy.yaml" {
		t.Fatalf("got unexpected default value for '%s' API server config: %s",
			auditPolicyKey, a[auditPolicyKey])
	}
}

func TestAPIServerWeakCipherSuites(t *testing.T) {
	// Test allowed versions
	for _, version := range []string{"1.13.0", "1.14.0"} {
		cs := CreateMockContainerService("testcluster", version, 3, 2, false)
		cs.setAPIServerConfig()
		a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
		if a["--tls-cipher-suites"] != TLSStrongCipherSuitesAPIServer {
			t.Fatalf("got unexpected default value for '--tls-cipher-suites' API server config for Kubernetes version %s: %s",
				version, a["--tls-cipher-suites"])
		}
	}

	allSuites := "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_RC4_128_SHA,TLS_RSA_WITH_3DES_EDE_CBC_SHA,TLS_RSA_WITH_AES_128_CBC_SHA,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_RSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_256_CBC_SHA,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_RC4_128_SHA"
	// Test user-override
	for _, version := range []string{"1.13.0", "1.14.0"} {
		cs := CreateMockContainerService("testcluster", version, 3, 2, false)
		cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{
			"--tls-cipher-suites": allSuites,
		}
		cs.setAPIServerConfig()
		a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
		if a["--tls-cipher-suites"] != allSuites {
			t.Fatalf("got unexpected default value for '--tls-cipher-suites' API server config for Kubernetes version %s: %s",
				version, a["--tls-cipher-suites"])
		}
	}
}

func TestAPIServerCosmosEtcd(t *testing.T) {
	// Test default
	cs := CreateMockContainerService("testcluster", "1.15.4", 3, 2, false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--etcd-cafile"] != "/etc/kubernetes/certs/ca.crt" {
		t.Fatalf("got unexpected default value for '--etcd-cafile' API server config: %s",
			a["--etcd-cafile"])
	}
	if a["--etcd-servers"] != fmt.Sprintf("https://127.0.0.1:%s", strconv.Itoa(DefaultMasterEtcdClientPort)) {
		t.Fatalf("got unexpected default value for '--etcd-servers' API server config: %s",
			a["--etcd-servers"])
	}

	// Validate that 1.14.0 doesn't include --repair-malformed-updates at all
	cs = CreateMockContainerService("testcluster", "1.14.0", 3, 2, false)
	cs.Properties.MasterProfile.CosmosEtcd = to.BoolPtr(true)
	cs.Properties.MasterProfile.DNSPrefix = "my-cosmos"
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--etcd-servers"] != fmt.Sprintf("https://%s:%s", cs.Properties.MasterProfile.GetCosmosEndPointURI(), strconv.Itoa(DefaultMasterEtcdClientPort)) {
		t.Fatalf("got unexpected default value for '--etcd-servers' API server config: %s",
			a["--etcd-servers"])
	}
}

func TestAPIServerFeatureGates(t *testing.T) {
	// Test k8s < 1.13
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for k8s v%s: %s",
			defaultTestClusterVer, a["--feature-gates"])
	}

	// Test 1.13 <= k8s <= 1.16
	cs = CreateMockContainerService("testcluster", "1.14.0", 3, 2, false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--feature-gates"] != "VolumeSnapshotDataSource=true" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for k8s v%s: %s",
			"1.14.0", a["--feature-gates"])
	}

	// Test k8s >= 1.17
	cs = CreateMockContainerService("testcluster", "1.17.0", 3, 2, false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for k8s v%s: %s",
			"1.17.0", a["--feature-gates"])
	}
}
