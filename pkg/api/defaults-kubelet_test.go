package api

import (
	"strconv"
	"testing"

	"github.com/Azure/aks-engine/pkg/helpers"
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = helpers.PointerToBool(true)
	cs.setKubeletConfig()
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--cloud-provider"] != "external" {
		t.Fatalf("got unexpected '--cloud-provider' kubelet config value for UseCloudControllerManager=true: %s",
			k["--cloud-provider"])
	}

	// Test UseCloudControllerManager = false
	cs = CreateMockContainerService("testcluster", defaultTestClusterVer, 3, 2, false)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = helpers.PointerToBool(false)
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = helpers.PointerToBool(true)
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
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = helpers.PointerToBool(false)
	cs.setKubeletConfig()
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	for _, key := range []string{"--anonymous-auth", "--client-ca-file"} {
		if _, ok := k[key]; ok {
			t.Fatalf("got unexpected '%s' kubelet config value for EnableSecureKubelet=false: %s",
				key, k[key])
		}
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
