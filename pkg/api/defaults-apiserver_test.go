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

var defaultTestClusterVer = common.RationalizeReleaseAndVersion(Kubernetes, common.KubernetesDefaultRelease, "", false, false, false)

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
	cs = CreateMockContainerService("testcluster", "", 3, 2, false)
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

	cs.Location = "chinaeast3"
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

	cs.Location = "chinanorth3"
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

	// Test EnableRbac = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--authorization-mode"]; ok {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for EnableRbac=false: %s",
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

func TestAPIServerServiceAccountFlags(t *testing.T) {
	cs := CreateMockContainerService("testcluster", common.RationalizeReleaseAndVersion(Kubernetes, "1.23", "", false, false, false), 3, 2, false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--service-account-issuer"] != "https://kubernetes.default.svc.cluster.local" {
		t.Fatalf("got unexpected '--service-account-issuer' API server config value for Kubernetes v1.23: %s",
			a["--service-account-issuer"])
	}
	if a["--service-account-signing-key-file"] != "/etc/kubernetes/certs/apiserver.key" {
		t.Fatalf("got unexpected '--service-account-signing-key-file' API server config value for Kubernetes v1.23: %s",
			a["--service-account-signing-key-file"])
	}

	cs = CreateMockContainerService("testcluster", common.RationalizeReleaseAndVersion(Kubernetes, "1.23", "", false, false, false), 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{
		"--service-account-issuer": "kubernetes.default.svc",
	}
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--service-account-issuer"] != "https://kubernetes.default.svc.cluster.local" {
		t.Fatalf("got unexpected '--service-account-issuer' API server config value for Kubernetes v1.23: %s",
			a["--service-account-issuer"])
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
	cs := CreateMockContainerService("testcluster", "", 3, 2, false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--repair-malformed-updates"]; ok {
		t.Fatalf("got unexpected default value for '--repair-malformed-updates' API server config: %s",
			a["--repair-malformed-updates"])
	}
}

func TestAPIServerAuditPolicyBackCompatOverride(t *testing.T) {
	// Validate that we statically override "--audit-policy-file" values of "/etc/kubernetes/manifests/audit-policy.yaml" for back-compat
	auditPolicyKey := "--audit-policy-file"
	cs := CreateMockContainerService("testcluster", "", 3, 2, false)
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
	for _, version := range []string{"1.15.12", "1.16.9", "1.17.5", "1.18.2"} {
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
	for _, version := range []string{"1.15.12", "1.16.9", "1.17.5", "1.18.2"} {
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
	cs := CreateMockContainerService("testcluster", "", 3, 2, false)
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

	cs = CreateMockContainerService("testcluster", "", 3, 2, false)
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
	// test defaultTestClusterVer
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for k8s v%s: %s",
			defaultTestClusterVer, a["--feature-gates"])
	}

	// test 1.19.0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.19.0"
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for k8s v%s: %s",
			"1.19.0", a["--feature-gates"])
	}

	// test 1.22.0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.22.0"
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for k8s v%s: %s",
			"1.22.0", a["--feature-gates"])
	}

	// test user-overrides, removal of VolumeSnapshotDataSource for k8s versions >= 1.22
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.22.0"
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = make(map[string]string)
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	a["--feature-gates"] = "VolumeSnapshotDataSource=true"
	cs.setAPIServerConfig()
	if a["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for \"--feature-gates\": \"VolumeSnapshotDataSource=true\": %s for k8s v%s",
			a["--feature-gates"], "1.22.0")
	}

	// test user-overrides, no removal of VolumeSnapshotDataSource for k8s versions < 1.22
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.19.0"
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = make(map[string]string)
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	a["--feature-gates"] = "VolumeSnapshotDataSource=true"
	cs.setAPIServerConfig()
	if a["--feature-gates"] != "VolumeSnapshotDataSource=true" {
		t.Fatalf("got unexpected '--feature-gates' API server config value for \"--feature-gates\": \"VolumeSnapshotDataSource=true\": %s for k8s v%s",
			a["--feature-gates"], "1.19.0")
	}
}

func TestAPIServerInsecureFlag(t *testing.T) {
	type apiServerTest struct {
		version string
		found   bool
	}

	apiTests := []apiServerTest{
		{
			version: "1.19.16",
			found:   true,
		},
		{
			version: "1.20.0",
			found:   false,
		},
		{
			version: "1.21.0",
			found:   false,
		},
		{
			version: "1.22.0",
			found:   false,
		},
		{
			version: "1.23.0",
			found:   false,
		},
		{
			version: "1.24.0",
			found:   false,
		},
		{
			version: "1.24.0-alpha.0",
			found:   false,
		},
		{
			version: "1.24.0-alpha.1-24",
			found:   false,
		},
	}

	for _, tt := range apiTests {
		cs := CreateMockContainerService("testcluster", tt.version, 3, 2, false)
		cs.setAPIServerConfig()
		a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig

		v, found := a["--insecure-port"]
		if found != tt.found {
			t.Fatalf("got found %t want %t", found, tt.found)
		}

		if tt.found && v != "0" {
			t.Fatalf("got unexpected '--insecure-port' API server config value for k8s v%s: %s",
				defaultTestClusterVer, a["--insecure-port"])
		}
	}

}

func TestAPIServerIPv6Only(t *testing.T) {
	cs := CreateMockContainerService("testcluster", "", 3, 2, false)
	cs.Properties.FeatureFlags = &FeatureFlags{EnableIPv6Only: true}
	cs.setAPIServerConfig()

	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	// bind address should be :: for single stack IPv6 cluster
	if a["--bind-address"] != "::" {
		t.Fatalf("got unexpected default value for '--bind-address' API server config: %s",
			a["--bind-address"])
	}
	for _, key := range []string{"--advertise-address"} {
		if _, ok := a[key]; ok {
			t.Fatalf("got unexpected '%s' API server config value for '--advertise-address' %s",
				key, a[key])
		}
	}
}

func TestAPIServerAnonymousAuth(t *testing.T) {
	// Validate anonymous-auth default is false
	cs := CreateMockContainerService("testcluster", "", 3, 2, false)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--anonymous-auth"] != "false" {
		t.Fatalf("got unexpected '--anonymous-auth' API server config value: %s",
			a["--anonymous-auth"])
	}

	// Validate anonymous-auth enabled
	cs = CreateMockContainerService("testcluster", "", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{
		"--anonymous-auth": "true",
	}
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--anonymous-auth"] != "true" {
		t.Fatalf("got unexpected '--anonymous-auth' API server config value: %s",
			a["--anonymous-auth"])
	}
}

func TestAPIServerConfigChangeVerbosity(t *testing.T) {
	// Test
	// "apiServerConfig": {
	// 	"--v": "4"
	// },
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{
		"--v": "4",
	}
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--v"] != "4" {
		t.Fatalf("got unexpected '--v' API server config value for \"--v\": \"4\": %s",
			a["--v"])
	}

	// Test default
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--v"] != DefaultKubernetesAPIServerVerbosity {
		t.Fatalf("got unexpected default value for '--v' API server config: %s",
			a["--v"])
	}
}
