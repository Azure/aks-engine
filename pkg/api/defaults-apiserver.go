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
		"--anonymous-auth":              "false",
		"--audit-log-path":              "/var/log/kubeaudit/audit.log",
		"--insecure-port":               "8080",
		"--secure-port":                 "443",
		"--service-account-lookup":      "true",
		"--etcd-certfile":               "/etc/kubernetes/certs/etcdclient.crt",
		"--etcd-keyfile":                "/etc/kubernetes/certs/etcdclient.key",
		"--tls-cert-file":               "/etc/kubernetes/certs/apiserver.crt",
		"--tls-private-key-file":        "/etc/kubernetes/certs/apiserver.key",
		"--client-ca-file":              "/etc/kubernetes/certs/ca.crt",
		"--repair-malformed-updates":    "false",
		"--service-account-key-file":    "/etc/kubernetes/certs/apiserver.key",
		"--kubelet-client-certificate":  "/etc/kubernetes/certs/client.crt",
		"--kubelet-client-key":          "/etc/kubernetes/certs/client.key",
		"--service-cluster-ip-range":    o.KubernetesConfig.ServiceCIDR,
		"--storage-backend":             o.GetAPIServerEtcdAPIVersion(),
		"--enable-bootstrap-token-auth": "true",
		"--v":                           "4",
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
		"--audit-log-maxage":    "30",
		"--audit-log-maxbackup": "10",
		"--audit-log-maxsize":   "100",
		"--profiling":           DefaultKubernetesAPIServerEnableProfiling,
		"--tls-cipher-suites":   TLSStrongCipherSuitesAPIServer,
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
		if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.7.0") {
			defaultAPIServerConfig["--authorization-mode"] = "Node,RBAC"
		} else {
			defaultAPIServerConfig["--authorization-mode"] = "RBAC"
		}
	}

	// Set default admission controllers
	admissionControlKey, admissionControlValues := getDefaultAdmissionControls(cs)
	defaultAPIServerConfig[admissionControlKey] = admissionControlValues

	// Enable VolumeSnapshotDataSource feature gate for Azure Disk CSI Driver
	// which is disabled from 1.13 to 1.16 by default
	if !common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.17.0") {
		addDefaultFeatureGates(defaultAPIServerConfig, o.OrchestratorVersion, "1.13.0", "VolumeSnapshotDataSource=true")
	}

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
	// Remove flags that are not compatible with 1.10
	if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.10.0") {
		for _, key := range []string{"--admission-control"} {
			delete(o.KubernetesConfig.APIServerConfig, key)
		}
	}
	// Enforce flags removal that don't work with specific versions, to accommodate upgrade
	// Remove flags that are not compatible with 1.14
	if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.14.0-alpha.1") {
		for _, key := range []string{"--repair-malformed-updates"} {
			delete(o.KubernetesConfig.APIServerConfig, key)
		}
	}
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
