// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
)

func TestKubeletConfigDefaults(t *testing.T) {
	cs := CreateMockContainerService("testcluster", "1.8.6", 3, 2, false)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	// TODO test all default config values
	for key, val := range map[string]string{
		"--azure-container-registry-config": "/etc/kubernetes/azure.json",
		"--image-pull-progress-deadline":    "30m",
	} {
		if k[key] != val {
			t.Fatalf("got unexpected kubelet config value for %s: %s, expected %s",
				key, k[key], val)
		}
	}

	cs = CreateMockContainerService("testcluster", "1.8.6", 3, 2, false)
	// TODO test all default overrides
	overrideVal := "/etc/override"
	cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig = map[string]string{
		"--azure-container-registry-config": overrideVal,
	}
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	for key, val := range map[string]string{"--azure-container-registry-config": overrideVal} {
		if k[key] != val {
			t.Fatalf("got unexpected kubelet config value for %s: %s, expected %s",
				key, k[key], val)
		}
	}
}

func TestKubeletConfigUseCloudControllerManager(t *testing.T) {
	// Test UseCloudControllerManager = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = to.BoolPtr(true)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--cloud-provider"] != "external" {
		t.Fatalf("got unexpected '--cloud-provider' kubelet config value for UseCloudControllerManager=true: %s",
			k["--cloud-provider"])
	}

	// Test UseCloudControllerManager = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = to.BoolPtr(false)
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--cloud-provider"] != "azure" {
		t.Fatalf("got unexpected '--cloud-provider' kubelet config value for UseCloudControllerManager=false: %s",
			k["--cloud-provider"])
	}

}

func TestKubeletConfigCloudConfig(t *testing.T) {
	// Test default value and custom value for --cloud-config
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--cloud-config"] != "/etc/kubernetes/azure.json" {
		t.Fatalf("got unexpected '--cloud-config' kubelet config default value: %s",
			k["--cloud-config"])
	}

	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--cloud-config"] = "custom.json"
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--cloud-config"] != "custom.json" {
		t.Fatalf("got unexpected '--cloud-config' kubelet config default value: %s",
			k["--cloud-config"])
	}
}

func TestKubeletConfigAzureContainerRegistryCofig(t *testing.T) {
	// Test default value and custom value for --azure-container-registry-config
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--azure-container-registry-config"] != "/etc/kubernetes/azure.json" {
		t.Fatalf("got unexpected '--azure-container-registry-config' kubelet config default value: %s",
			k["--azure-container-registry-config"])
	}

	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--azure-container-registry-config"] = "custom.json"
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--azure-container-registry-config"] != "custom.json" {
		t.Fatalf("got unexpected '--azure-container-registry-config' kubelet config default value: %s",
			k["--azure-container-registry-config"])
	}
}

func TestKubeletConfigNetworkPlugin(t *testing.T) {
	// Test NetworkPlugin = "kubenet"
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--network-plugin"] != NetworkPluginKubenet {
		t.Fatalf("got unexpected '--network-plugin' kubelet config value for NetworkPlugin=kubenet: %s",
			k["--network-plugin"])
	}

	// Test NetworkPlugin = "azure"
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--network-plugin"] != "cni" {
		t.Fatalf("got unexpected '--network-plugin' kubelet config value for NetworkPlugin=azure: %s",
			k["--network-plugin"])
	}

}

func TestKubeletConfigEnableSecureKubelet(t *testing.T) {
	// Test EnableSecureKubelet = true
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = to.BoolPtr(true)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--anonymous-auth"] != "false" {
		t.Fatalf("got unexpected '--anonymous-auth' kubelet config value for EnableSecureKubelet=true: %s",
			k["--anonymous-auth"])
	}
	if k["--authorization-mode"] != "Webhook" {
		t.Fatalf("got unexpected '--authorization-mode' kubelet config value for EnableSecureKubelet=true: %s",
			k["--authorization-mode"])
	}
	if k["--client-ca-file"] != "/etc/kubernetes/certs/ca.crt" {
		t.Fatalf("got unexpected '--client-ca-file' kubelet config value for EnableSecureKubelet=true: %s",
			k["--client-ca-file"])
	}

	// Test EnableSecureKubelet = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = to.BoolPtr(false)
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	for _, key := range []string{"--anonymous-auth", "--client-ca-file"} {
		if _, ok := k[key]; ok {
			t.Fatalf("got unexpected '%s' kubelet config value for EnableSecureKubelet=false: %s",
				key, k[key])
		}
	}

	// Test default (EnableSecureKubelet = false) for Windows
	cs = CreateMockContainerService("testcluster", "1.10.13", 3, 1, false)
	p := GetK8sDefaultProperties(true)
	cs.Properties = p
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	for _, key := range []string{"--anonymous-auth", "--client-ca-file"} {
		if _, ok := k[key]; ok {
			t.Fatalf("got unexpected '%s' kubelet config value for EnableSecureKubelet=false: %s",
				key, k[key])
		}
	}

	// Test explicit EnableSecureKubelet = false for Windows
	cs = CreateMockContainerService("testcluster", "1.10.13", 3, 1, false)
	p = GetK8sDefaultProperties(true)
	cs.Properties = p
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = to.BoolPtr(false)
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	for _, key := range []string{"--anonymous-auth", "--client-ca-file"} {
		if _, ok := k[key]; ok {
			t.Fatalf("got unexpected '%s' kubelet config value for EnableSecureKubelet=false: %s",
				key, k[key])
		}
	}

	// Test EnableSecureKubelet = true for Windows
	cs = CreateMockContainerService("testcluster", "1.10.13", 3, 1, false)
	p = GetK8sDefaultProperties(true)
	cs.Properties = p
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = to.BoolPtr(true)
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--anonymous-auth"] != "false" {
		t.Fatalf("got unexpected '--anonymous-auth' kubelet config value for EnableSecureKubelet=true: %s",
			k["--anonymous-auth"])
	}
	if k["--client-ca-file"] != "/etc/kubernetes/certs/ca.crt" {
		t.Fatalf("got unexpected '--client-ca-file' kubelet config value for EnableSecureKubelet=true: %s",
			k["--client-ca-file"])
	}

}

func TestKubeletMaxPods(t *testing.T) {
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--max-pods"] != strconv.Itoa(DefaultKubernetesMaxPodsVNETIntegrated) {
		t.Fatalf("got unexpected '--max-pods' kubelet config value for NetworkPolicy=%s: %s",
			NetworkPluginAzure, k["--max-pods"])
	}

	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--max-pods"] != strconv.Itoa(DefaultKubernetesMaxPods) {
		t.Fatalf("got unexpected '--max-pods' kubelet config value for NetworkPolicy=%s: %s",
			NetworkPluginKubenet, k["--max-pods"])
	}

	// Test that user-overrides for --max-pods work as intended
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
	cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--max-pods"] = "99"
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--max-pods"] != "99" {
		t.Fatalf("got unexpected '--max-pods' kubelet config value for NetworkPolicy=%s: %s",
			NetworkPluginKubenet, k["--max-pods"])
	}

	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--max-pods"] = "99"
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--max-pods"] != "99" {
		t.Fatalf("got unexpected '--max-pods' kubelet config value for NetworkPolicy=%s: %s",
			NetworkPluginKubenet, k["--max-pods"])
	}
}

func TestKubeletCalico(t *testing.T) {
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyCalico
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--network-plugin"] != "cni" {
		t.Fatalf("got unexpected '--network-plugin' kubelet config value for NetworkPolicy=%s: %s",
			NetworkPolicyCalico, k["--network-plugin"])
	}
}

func TestKubeletHostedMasterIPMasqAgentDisabled(t *testing.T) {
	subnet := "172.16.0.0/16"
	// MasterIPMasqAgent disabled, --non-masquerade-cidr should be subnet
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.HostedMasterProfile = &HostedMasterProfile{
		IPMasqAgent: false,
	}
	cs.Properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = subnet
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--non-masquerade-cidr"] != subnet {
		t.Fatalf("got unexpected '--non-masquerade-cidr' kubelet config value %s, the expected value is %s",
			k["--non-masquerade-cidr"], subnet)
	}

	// MasterIPMasqAgent enabled, --non-masquerade-cidr should be 0.0.0.0/0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.HostedMasterProfile = &HostedMasterProfile{
		IPMasqAgent: true,
	}
	cs.Properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = subnet
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--non-masquerade-cidr"] != DefaultNonMasqueradeCIDR {
		t.Fatalf("got unexpected '--non-masquerade-cidr' kubelet config value %s, the expected value is %s",
			k["--non-masquerade-cidr"], DefaultNonMasqueradeCIDR)
	}

	// no HostedMasterProfile, --non-masquerade-cidr should be 0.0.0.0/0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = subnet
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--non-masquerade-cidr"] != DefaultNonMasqueradeCIDR {
		t.Fatalf("got unexpected '--non-masquerade-cidr' kubelet config value %s, the expected value is %s",
			k["--non-masquerade-cidr"], DefaultNonMasqueradeCIDR)
	}
}

func TestKubeletIPMasqAgentEnabledOrDisabled(t *testing.T) {
	subnet := "172.16.0.0/16"
	// MasterIPMasqAgent disabled, --non-masquerade-cidr should be subnet
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	b := false
	cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    IPMASQAgentAddonName,
				Enabled: &b,
			},
		},
	}
	cs.Properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = subnet
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--non-masquerade-cidr"] != subnet {
		t.Fatalf("got unexpected '--non-masquerade-cidr' kubelet config value %s, the expected value is %s",
			k["--non-masquerade-cidr"], subnet)
	}

	// MasterIPMasqAgent enabled, --non-masquerade-cidr should be 0.0.0.0/0
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	b = true
	cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    IPMASQAgentAddonName,
				Enabled: &b,
			},
		},
	}
	cs.Properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = subnet
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--non-masquerade-cidr"] != DefaultNonMasqueradeCIDR {
		t.Fatalf("got unexpected '--non-masquerade-cidr' kubelet config value %s, the expected value is %s",
			k["--non-masquerade-cidr"], DefaultNonMasqueradeCIDR)
	}
}

func TestEnforceNodeAllocatable(t *testing.T) {
	// Validate default
	cs := CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--enforce-node-allocatable"] != "pods" {
		t.Fatalf("got unexpected '--enforce-node-allocatable' kubelet config value %s, the expected value is %s",
			k["--enforce-node-allocatable"], "pods")
	}

	// Validate that --enforce-node-allocatable is overridable
	cs = CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		KubeletConfig: map[string]string{
			"--enforce-node-allocatable": "kube-reserved/system-reserved",
		},
	}
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--enforce-node-allocatable"] != "kube-reserved/system-reserved" {
		t.Fatalf("got unexpected '--enforce-node-allocatable' kubelet config value %s, the expected value is %s",
			k["--enforce-node-allocatable"], "kube-reserved/system-reserved")
	}
}

func TestStaticWindowsConfig(t *testing.T) {
	cs := CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 1, false)
	p := GetK8sDefaultProperties(true)
	cs.Properties = p
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = to.BoolPtr(true)

	// Start with copy of Linux config
	staticLinuxKubeletConfig := map[string]string{
		"--address":                     "0.0.0.0",
		"--allow-privileged":            "true",
		"--anonymous-auth":              "false",
		"--authorization-mode":          "Webhook",
		"--client-ca-file":              "/etc/kubernetes/certs/ca.crt",
		"--pod-manifest-path":           "/etc/kubernetes/manifests",
		"--cluster-dns":                 cs.Properties.OrchestratorProfile.KubernetesConfig.DNSServiceIP,
		"--cgroups-per-qos":             "true",
		"--kubeconfig":                  "/var/lib/kubelet/kubeconfig",
		"--keep-terminated-pod-volumes": "false",
	}
	staticWindowsKubeletConfig := make(map[string]string)
	for key, val := range staticLinuxKubeletConfig {
		if key != "--pod-manifest-path" {
			staticWindowsKubeletConfig[key] = val
		}
	}

	// Add Windows-specific overrides
	// Eventually paths should not be hardcoded here. They should be relative to $global:KubeDir in the PowerShell script
	staticWindowsKubeletConfig["--azure-container-registry-config"] = "c:\\k\\azure.json"
	staticWindowsKubeletConfig["--pod-infra-container-image"] = "kubletwin/pause"
	staticWindowsKubeletConfig["--kubeconfig"] = "c:\\k\\config"
	staticWindowsKubeletConfig["--cloud-config"] = "c:\\k\\azure.json"
	staticWindowsKubeletConfig["--cgroups-per-qos"] = "false"
	staticWindowsKubeletConfig["--enforce-node-allocatable"] = "\"\"\"\""
	staticWindowsKubeletConfig["--system-reserved"] = "memory=2Gi"
	staticWindowsKubeletConfig["--client-ca-file"] = "c:\\k\\ca.crt"
	staticWindowsKubeletConfig["--hairpin-mode"] = "promiscuous-bridge"
	staticWindowsKubeletConfig["--image-pull-progress-deadline"] = "20m"
	staticWindowsKubeletConfig["--resolv-conf"] = "\"\"\"\""
	staticWindowsKubeletConfig["--eviction-hard"] = "\"\"\"\""

	cs.setKubeletConfig()
	for _, profile := range cs.Properties.AgentPoolProfiles {
		if profile.OSType == Windows {
			for key, val := range staticWindowsKubeletConfig {
				if val != profile.KubernetesConfig.KubeletConfig[key] {
					t.Fatalf("got unexpected '%s' kubelet config value, expected %s, got %s",
						key, val, profile.KubernetesConfig.KubeletConfig[key])
				}
			}
		}
	}
}

func TestUbuntu1804Flags(t *testing.T) {
	// Validate --resolv-conf is missing with 16.04 distro and present with 18.04
	cs := CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	cs.Properties.MasterProfile.Distro = AKS
	cs.Properties.AgentPoolProfiles[0].Distro = AKS1804
	cs.Properties.AgentPoolProfiles[0].OSType = Linux
	cs.setKubeletConfig()
	km := cs.Properties.MasterProfile.KubernetesConfig.KubeletConfig
	if _, ok := km["--resolv-conf"]; ok {
		t.Fatalf("got unexpected '--resolv-conf' kubelet config value '%s' with Ubuntu 16.04 ",
			km["--resolv-conf"])
	}
	ka := cs.Properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig
	if ka["--resolv-conf"] != "/run/systemd/resolve/resolv.conf" {
		t.Fatalf("got unexpected '--resolv-conf' kubelet config value %s with Ubuntu 18.04, the expected value is %s",
			ka["--resolv-conf"], "/run/systemd/resolve/resolv.conf")
	}

	cs = CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	cs.Properties.MasterProfile.Distro = Ubuntu1804
	cs.Properties.AgentPoolProfiles[0].Distro = Ubuntu
	cs.Properties.AgentPoolProfiles[0].OSType = Linux
	cs.setKubeletConfig()
	km = cs.Properties.MasterProfile.KubernetesConfig.KubeletConfig
	if km["--resolv-conf"] != "/run/systemd/resolve/resolv.conf" {
		t.Fatalf("got unexpected '--resolv-conf' kubelet config value %s with Ubuntu 18.04, the expected value is %s",
			km["--resolv-conf"], "/run/systemd/resolve/resolv.conf")
	}
	ka = cs.Properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig
	if _, ok := ka["--resolv-conf"]; ok {
		t.Fatalf("got unexpected '--resolv-conf' kubelet config value '%s' with Ubuntu 16.04 ",
			ka["--resolv-conf"])
	}

	cs = CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	cs.Properties.MasterProfile.Distro = Ubuntu
	cs.Properties.AgentPoolProfiles[0].Distro = ""
	cs.Properties.AgentPoolProfiles[0].OSType = Windows
	cs.setKubeletConfig()
	km = cs.Properties.MasterProfile.KubernetesConfig.KubeletConfig
	if _, ok := km["--resolv-conf"]; ok {
		t.Fatalf("got unexpected '--resolv-conf' kubelet config value '%s' with Ubuntu 16.04 ",
			km["--resolv-conf"])
	}
	ka = cs.Properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig
	if ka["--resolv-conf"] != "\"\"\"\"" {
		t.Fatalf("got unexpected '--resolv-conf' kubelet config value %s with Windows, the expected value is %s",
			ka["--resolv-conf"], "\"\"\"\"")
	}
}

func TestKubeletConfigDefaultFeatureGates(t *testing.T) {
	// test 1.7
	cs := CreateMockContainerService("testcluster", "1.7.12", 3, 2, false)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--feature-gates"] != "" {
		t.Fatalf("got unexpected '--feature-gates' kubelet config value for \"--feature-gates\": \"\": %s",
			k["--feature-gates"])
	}

	// test 1.8
	cs = CreateMockContainerService("testcluster", "1.8.15", 3, 2, false)
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--feature-gates"] != "PodPriority=true" {
		t.Fatalf("got unexpected '--feature-gates' kubelet config value for \"--feature-gates\": \"\": %s",
			k["--feature-gates"])
	}

	// test 1.14
	cs = CreateMockContainerService("testcluster", "1.14.1", 3, 2, false)
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--feature-gates"] != "PodPriority=true" {
		t.Fatalf("got unexpected '--feature-gates' kubelet config value for \"--feature-gates\": \"\": %s",
			k["--feature-gates"])
	}

	// test user-overrides
	cs = CreateMockContainerService("testcluster", "1.14.1", 3, 2, false)
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	k["--feature-gates"] = "DynamicKubeletConfig=true"
	cs.setKubeletConfig()
	if k["--feature-gates"] != "DynamicKubeletConfig=true,PodPriority=true" {
		t.Fatalf("got unexpected '--feature-gates' kubelet config value for \"--feature-gates\": \"\": %s",
			k["--feature-gates"])
	}
}
