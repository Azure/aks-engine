// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"fmt"
	"strconv"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
)

func (cs *ContainerService) setAPIServerConfig() {
	o := cs.Properties.OrchestratorProfile
	staticAPIServerConfig := map[string]string{
		"--bind-address":                "0.0.0.0",
		"--advertise-address":           "<advertiseAddr>",
		"--allow-privileged":            "true",
		"--audit-log-path":              "/var/log/kubeaudit/audit.log",
		"--secure-port":                 "443",
		"--service-account-lookup":      "true",
		"--etcd-certfile":               "/etc/kubernetes/certs/etcdclient.crt",
		"--etcd-keyfile":                "/etc/kubernetes/certs/etcdclient.key",
		"--tls-cert-file":               "/etc/kubernetes/certs/apiserver.crt",
		"--tls-private-key-file":        "/etc/kubernetes/certs/apiserver.key",
		"--client-ca-file":              "/etc/kubernetes/certs/ca.crt",
		"--service-account-key-file":    "/etc/kubernetes/certs/apiserver.key",
		"--kubelet-client-certificate":  "/etc/kubernetes/certs/client.crt",
		"--kubelet-client-key":          "/etc/kubernetes/certs/client.key",
		"--service-cluster-ip-range":    o.KubernetesConfig.ServiceCIDR,
		"--storage-backend":             o.GetAPIServerEtcdAPIVersion(),
		"--enable-bootstrap-token-auth": "true",
	}

	if cs.Properties.MasterProfile != nil {
		if cs.Properties.MasterProfile.HasCosmosEtcd() {
			// Configuration for cosmos etcd
			staticAPIServerConfig["--etcd-servers"] = fmt.Sprintf("https://%s:%s", cs.Properties.MasterProfile.GetCosmosEndPointURI(), strconv.Itoa(DefaultMasterEtcdClientPort))
		} else {
			// Configuration for local etcd
			staticAPIServerConfig["--etcd-cafile"] = "/etc/kubernetes/certs/ca.crt"
			staticAPIServerConfig["--etcd-servers"] = fmt.Sprintf("https://127.0.0.1:%s", strconv.Itoa(DefaultMasterEtcdClientPort))
		}
	}

	// Default apiserver config
	defaultAPIServerConfig := map[string]string{
		"--anonymous-auth":      "false",
		"--audit-log-maxage":    "30",
		"--audit-log-maxbackup": "10",
		"--audit-log-maxsize":   "100",
		"--profiling":           DefaultKubernetesAPIServerEnableProfiling,
		"--tls-cipher-suites":   TLSStrongCipherSuitesAPIServer,
		"--v":                   DefaultKubernetesAPIServerVerbosity,
	}

	// Data Encryption at REST configuration conditions
	if to.Bool(o.KubernetesConfig.EnableDataEncryptionAtRest) || to.Bool(o.KubernetesConfig.EnableEncryptionWithExternalKms) {
		staticAPIServerConfig["--encryption-provider-config"] = "/etc/kubernetes/encryption-config.yaml"
	}

	// Aggregated API configuration
	if o.KubernetesConfig.EnableAggregatedAPIs {
		defaultAPIServerConfig["--requestheader-client-ca-file"] = "/etc/kubernetes/certs/proxy-ca.crt"
		defaultAPIServerConfig["--proxy-client-cert-file"] = "/etc/kubernetes/certs/proxy.crt"
		defaultAPIServerConfig["--proxy-client-key-file"] = "/etc/kubernetes/certs/proxy.key"
		defaultAPIServerConfig["--requestheader-allowed-names"] = ""
		defaultAPIServerConfig["--requestheader-extra-headers-prefix"] = "X-Remote-Extra-"
		defaultAPIServerConfig["--requestheader-group-headers"] = "X-Remote-Group"
		defaultAPIServerConfig["--requestheader-username-headers"] = "X-Remote-User"
	}

	// Enable cloudprovider if we're not using cloud controller manager
	if !to.Bool(o.KubernetesConfig.UseCloudControllerManager) {
		staticAPIServerConfig["--cloud-provider"] = "azure"
		staticAPIServerConfig["--cloud-config"] = "/etc/kubernetes/azure.json"
	}

	// AAD configuration
	if cs.Properties.HasAadProfile() {
		defaultAPIServerConfig["--oidc-username-claim"] = "oid"
		defaultAPIServerConfig["--oidc-groups-claim"] = "groups"
		defaultAPIServerConfig["--oidc-client-id"] = "spn:" + cs.Properties.AADProfile.ServerAppID
		issuerHost := "sts.windows.net"
		if helpers.GetTargetEnv(cs.Location, cs.Properties.GetCustomCloudName()) == "AzureChinaCloud" {
			issuerHost = "sts.chinacloudapi.cn"
		}
		defaultAPIServerConfig["--oidc-issuer-url"] = "https://" + issuerHost + "/" + cs.Properties.AADProfile.TenantID + "/"
	}

	// Audit Policy configuration
	defaultAPIServerConfig["--audit-policy-file"] = "/etc/kubernetes/addons/audit-policy.yaml"

	// RBAC configuration
	if to.Bool(o.KubernetesConfig.EnableRbac) {
		defaultAPIServerConfig["--authorization-mode"] = "Node,RBAC"
	}

	if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.20.0-alpha.1") {
		defaultAPIServerConfig["--service-account-issuer"] = "https://kubernetes.default.svc.cluster.local"
		defaultAPIServerConfig["--service-account-signing-key-file"] = "/etc/kubernetes/certs/apiserver.key"
	}

	if !common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.20.0-alpha.0") {
		defaultAPIServerConfig["--insecure-port"] = "0"
	}

	// Set default admission controllers
	admissionControlKey, admissionControlValues := getDefaultAdmissionControls(cs)
	defaultAPIServerConfig[admissionControlKey] = admissionControlValues

	// If no user-configurable apiserver config values exists, use the defaults
	if o.KubernetesConfig.APIServerConfig == nil {
		o.KubernetesConfig.APIServerConfig = defaultAPIServerConfig
	} else {
		for key, val := range defaultAPIServerConfig {
			// If we don't have a user-configurable apiserver config for each option
			if _, ok := o.KubernetesConfig.APIServerConfig[key]; !ok {
				// then assign the default value
				o.KubernetesConfig.APIServerConfig[key] = val
			} else {
				// Manual override of "--audit-policy-file" for back-compat
				if key == "--audit-policy-file" {
					if o.KubernetesConfig.APIServerConfig[key] == "/etc/kubernetes/manifests/audit-policy.yaml" {
						o.KubernetesConfig.APIServerConfig[key] = val
					}
				}
			}
		}
	}

	// We don't support user-configurable values for the following,
	// so any of the value assignments below will override user-provided values
	for key, val := range staticAPIServerConfig {
		o.KubernetesConfig.APIServerConfig[key] = val
	}

	// Remove flags for secure communication to kubelet, if configured
	if !to.Bool(o.KubernetesConfig.EnableSecureKubelet) {
		for _, key := range []string{"--kubelet-client-certificate", "--kubelet-client-key"} {
			delete(o.KubernetesConfig.APIServerConfig, key)
		}
	}

	// Enforce flags removal that don't work with specific versions, to accommodate upgrade
	// Remove flags that are not compatible with any supported versions
	for _, key := range []string{"--admission-control", "--repair-malformed-updates"} {
		delete(o.KubernetesConfig.APIServerConfig, key)
	}

	// Set bind address to prefer IPv6 address for single stack IPv6 cluster
	// Remove --advertise-address so that --bind-address will be used
	if cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6Only") {
		o.KubernetesConfig.APIServerConfig["--bind-address"] = "::"
		for _, key := range []string{"--advertise-address"} {
			delete(o.KubernetesConfig.APIServerConfig, key)
		}
	}

	// Manual override of "--service-account-issuer" starting with 1.20
	if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.20.0-alpha.1") && o.KubernetesConfig.APIServerConfig["--service-account-issuer"] == "kubernetes.default.svc" {
		o.KubernetesConfig.APIServerConfig["--service-account-issuer"] = "https://kubernetes.default.svc.cluster.local"
	}

	invalidFeatureGates := []string{}
	// Remove --feature-gate VolumeSnapshotDataSource starting with 1.22
	// Reference: https://github.com/kubernetes/kubernetes/pull/101531
	if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.22.0-alpha.1") {
		invalidFeatureGates = append(invalidFeatureGates, "VolumeSnapshotDataSource")
	}
	removeInvalidFeatureGates(o.KubernetesConfig.APIServerConfig, invalidFeatureGates)
}

func getDefaultAdmissionControls(cs *ContainerService) (string, string) {
	o := cs.Properties.OrchestratorProfile
	admissionControlKey := "--enable-admission-plugins"
	admissionControlValues := "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,ValidatingAdmissionWebhook,ResourceQuota,ExtendedResourceToleration"

	// Pod Security Policy configuration
	if o.KubernetesConfig.IsAddonEnabled(common.PodSecurityPolicyAddonName) {
		admissionControlValues += ",PodSecurityPolicy"
	}

	return admissionControlKey, admissionControlValues
}
