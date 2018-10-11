package api

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/helpers"
)

const defaultTestClusterVer = "1.7.12"

func TestAPIServerConfigEnableDataEncryptionAtRest(t *testing.T) {
	// Test EnableDataEncryptionAtRest = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableDataEncryptionAtRest = helpers.PointerToBool(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--experimental-encryption-provider-config"] != "/etc/kubernetes/encryption-config.yaml" {
		t.Fatalf("got unexpected '--experimental-encryption-provider-config' API server config value for EnableDataEncryptionAtRest=true: %s",
			a["--experimental-encryption-provider-config"])
	}

	// Test EnableDataEncryptionAtRest = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableDataEncryptionAtRest = helpers.PointerToBool(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--experimental-encryption-provider-config"]; ok {
		t.Fatalf("got unexpected '--experimental-encryption-provider-config' API server config value for EnableDataEncryptionAtRest=false: %s",
			a["--experimental-encryption-provider-config"])
	}
}

func TestAPIServerConfigEnableEncryptionWithExternalKms(t *testing.T) {
	// Test EnableEncryptionWithExternalKms = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms = helpers.PointerToBool(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--experimental-encryption-provider-config"] != "/etc/kubernetes/encryption-config.yaml" {
		t.Fatalf("got unexpected '--experimental-encryption-provider-config' API server config value for EnableEncryptionWithExternalKms=true: %s",
			a["--experimental-encryption-provider-config"])
	}

	// Test EnableEncryptionWithExternalKms = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms = helpers.PointerToBool(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--experimental-encryption-provider-config"]; ok {
		t.Fatalf("got unexpected '--experimental-encryption-provider-config' API server config value for EnableEncryptionWithExternalKms=false: %s",
			a["--experimental-encryption-provider-config"])
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = helpers.PointerToBool(true)
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = helpers.PointerToBool(false)
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = helpers.PointerToBool(true)
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--authorization-mode"] != "Node,RBAC" {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for EnableRbac=true: %s",
			a["--authorization-mode"])
	}

	// Test EnableRbac = true with 1.6 cluster
	cs = CreateMockContainerService("testcluster", "1.6.11", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = helpers.PointerToBool(true)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if a["--authorization-mode"] != "RBAC" {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for 1.6 cluster with EnableRbac=true: %s",
			a["--authorization-mode"])
	}

	// Test EnableRbac = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = helpers.PointerToBool(false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig
	if _, ok := a["--authorization-mode"]; ok {
		t.Fatalf("got unexpected '--authorization-mode' API server config value for EnableRbac=false: %s",
			a["--authorization-mode"])
	}

	// Test EnableRbac = false with 1.6 cluster
	cs = CreateMockContainerService("testcluster", "1.6.11", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = helpers.PointerToBool(false)
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac = helpers.PointerToBool(false)
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = helpers.PointerToBool(true)
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = helpers.PointerToBool(false)
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
	// Test --enable-admission-plugins for v1.10 and above
	version := "1.10.0"
	enableAdmissionPluginsKey := "--enable-admission-plugins"
	admissonControlKey := "--admission-control"
	cs := CreateMockContainerService("testcluster", version, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig = map[string]string{}
	cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig[admissonControlKey] = "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,AlwaysPullImages,ExtendedResourceToleration"
	cs.setAPIServerConfig()
	a := cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig

	// --enable-admission-plugins should be set for v1.10 and above
	if _, found := a[enableAdmissionPluginsKey]; !found {
		t.Fatalf("Admission control key '%s' not set in API server config for version %s", enableAdmissionPluginsKey, version)
	}

	// --admission-control was deprecated in v1.10
	if _, found := a[admissonControlKey]; found {
		t.Fatalf("Deprecated admission control key '%s' set in API server config for version %s", admissonControlKey, version)
	}

	// Test --admission-control for v1.9 and below
	version = "1.9.0"
	cs = CreateMockContainerService("testcluster", version, 3, 2, false)
	cs.setAPIServerConfig()
	a = cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig

	// --enable-admission-plugins is available for v1.10 and above and should not be set here
	if _, found := a[enableAdmissionPluginsKey]; found {
		t.Fatalf("Unknown admission control key '%s' set in API server config for version %s", enableAdmissionPluginsKey, version)
	}

	// --admission-control is used for v1.9 and below
	if _, found := a[admissonControlKey]; !found {
		t.Fatalf("Admission control key '%s' not set in API server config for version %s", enableAdmissionPluginsKey, version)
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
