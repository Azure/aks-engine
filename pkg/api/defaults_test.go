// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestCertsAlreadyPresent(t *testing.T) {
	var cert *CertificateProfile

	result := certsAlreadyPresent(nil, 1)
	expected := map[string]bool{
		"ca":         false,
		"apiserver":  false,
		"client":     false,
		"kubeconfig": false,
		"etcd":       false,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("certsAlreadyPresent() did not return false for all certs for a non-existent CertificateProfile")
	}
	cert = &CertificateProfile{}
	result = certsAlreadyPresent(cert, 1)
	expected = map[string]bool{
		"ca":         false,
		"apiserver":  false,
		"client":     false,
		"kubeconfig": false,
		"etcd":       false,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("certsAlreadyPresent() did not return false for all certs for empty CertificateProfile")
	}
	cert = &CertificateProfile{
		APIServerCertificate: "a",
	}
	result = certsAlreadyPresent(cert, 1)
	expected = map[string]bool{
		"ca":         false,
		"apiserver":  false,
		"client":     false,
		"kubeconfig": false,
		"etcd":       false,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("certsAlreadyPresent() did not return false for all certs for 1 cert in CertificateProfile")
	}

	cert = &CertificateProfile{
		APIServerCertificate:  "a",
		CaCertificate:         "c",
		CaPrivateKey:          "d",
		ClientCertificate:     "e",
		ClientPrivateKey:      "f",
		KubeConfigCertificate: "g",
		KubeConfigPrivateKey:  "h",
		EtcdClientCertificate: "i",
		EtcdClientPrivateKey:  "j",
		EtcdServerCertificate: "k",
		EtcdServerPrivateKey:  "l",
	}
	result = certsAlreadyPresent(cert, 3)
	expected = map[string]bool{
		"ca":         true,
		"apiserver":  false,
		"client":     true,
		"kubeconfig": true,
		"etcd":       false,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("certsAlreadyPresent() did not return expected result for some certs in CertificateProfile")
	}
	cert = &CertificateProfile{
		APIServerCertificate:  "a",
		APIServerPrivateKey:   "b",
		CaCertificate:         "c",
		CaPrivateKey:          "d",
		ClientCertificate:     "e",
		ClientPrivateKey:      "f",
		KubeConfigCertificate: "g",
		KubeConfigPrivateKey:  "h",
		EtcdClientCertificate: "i",
		EtcdClientPrivateKey:  "j",
		EtcdServerCertificate: "k",
		EtcdServerPrivateKey:  "l",
		EtcdPeerCertificates:  []string{"0", "1", "2"},
		EtcdPeerPrivateKeys:   []string{"0", "1", "2"},
	}
	result = certsAlreadyPresent(cert, 3)
	expected = map[string]bool{
		"ca":         true,
		"apiserver":  true,
		"client":     true,
		"kubeconfig": true,
		"etcd":       true,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("certsAlreadyPresent() did not return expected result for all certs in CertificateProfile")
	}
}

func TestSetMissingKubeletValues(t *testing.T) {
	config := &KubernetesConfig{}
	defaultKubeletConfig := map[string]string{
		"--network-plugin":               "1",
		"--pod-infra-container-image":    "2",
		"--max-pods":                     "3",
		"--eviction-hard":                "4",
		"--node-status-update-frequency": "5",
		"--image-gc-high-threshold":      "6",
		"--image-gc-low-threshold":       "7",
		"--non-masquerade-cidr":          "8",
		"--cloud-provider":               "9",
		"--pod-max-pids":                 "10",
	}
	setMissingKubeletValues(config, defaultKubeletConfig)
	for key, val := range defaultKubeletConfig {
		if config.KubeletConfig[key] != val {
			t.Fatalf("setMissingKubeletValue() did not return the expected value %s for key %s, instead returned: %s", val, key, config.KubeletConfig[key])
		}
	}

	config = &KubernetesConfig{
		KubeletConfig: map[string]string{
			"--network-plugin":            "a",
			"--pod-infra-container-image": "b",
			"--cloud-provider":            "c",
		},
	}
	expectedResult := map[string]string{
		"--network-plugin":               "a",
		"--pod-infra-container-image":    "b",
		"--max-pods":                     "3",
		"--eviction-hard":                "4",
		"--node-status-update-frequency": "5",
		"--image-gc-high-threshold":      "6",
		"--image-gc-low-threshold":       "7",
		"--non-masquerade-cidr":          "8",
		"--cloud-provider":               "c",
		"--pod-max-pids":                 "10",
	}
	setMissingKubeletValues(config, defaultKubeletConfig)
	for key, val := range expectedResult {
		if config.KubeletConfig[key] != val {
			t.Fatalf("setMissingKubeletValue() did not return the expected value %s for key %s, instead returned: %s", val, key, config.KubeletConfig[key])
		}
	}
	config = &KubernetesConfig{
		KubeletConfig: map[string]string{},
	}
	setMissingKubeletValues(config, defaultKubeletConfig)
	for key, val := range defaultKubeletConfig {
		if config.KubeletConfig[key] != val {
			t.Fatalf("setMissingKubeletValue() did not return the expected value %s for key %s, instead returned: %s", val, key, config.KubeletConfig[key])
		}
	}
}

func TestAddonsIndexByName(t *testing.T) {
	addonName := "testaddon"
	addons := []KubernetesAddon{
		getMockAddon(addonName),
	}
	i := getAddonsIndexByName(addons, addonName)
	if i != 0 {
		t.Fatalf("addonsIndexByName() did not return the expected index value 0, instead returned: %d", i)
	}
	i = getAddonsIndexByName(addons, "nonExistentAddonName")
	if i != -1 {
		t.Fatalf("addonsIndexByName() did not return -1 for a non-existent addon, instead returned: %d", i)
	}
}

func TestAssignDefaultAddonImages(t *testing.T) {
	customImage := "myimage"
	defaultAddonImages := map[string]string{
		TillerAddonName:                    "gcr.io/kubernetes-helm/tiller:v2.13.1",
		ACIConnectorAddonName:              "microsoft/virtual-kubelet:latest",
		ClusterAutoscalerAddonName:         "k8s.gcr.io/cluster-autoscaler:v1.2.5",
		BlobfuseFlexVolumeAddonName:        "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.8",
		SMBFlexVolumeAddonName:             "mcr.microsoft.com/k8s/flexvolume/smb-flexvolume:1.0.2",
		KeyVaultFlexVolumeAddonName:        "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.13",
		DashboardAddonName:                 "k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1",
		ReschedulerAddonName:               "k8s.gcr.io/rescheduler:v0.3.1",
		MetricsServerAddonName:             "k8s.gcr.io/metrics-server-amd64:v0.2.1",
		NVIDIADevicePluginAddonName:        "nvidia/k8s-device-plugin:1.10",
		ContainerMonitoringAddonName:       "mcr.microsoft.com/azuremonitor/containerinsights/ciprod:ciprod07092019",
		IPMASQAgentAddonName:               "k8s.gcr.io/ip-masq-agent-amd64:v2.5.0",
		AzureCNINetworkMonitoringAddonName: "mcr.microsoft.com/containernetworking/networkmonitor:v0.0.6",
		DNSAutoscalerAddonName:             "k8s.gcr.io/cluster-proportional-autoscaler-amd64:1.1.1",
		HeapsterAddonName:                  "k8s.gcr.io/heapster-amd64:v1.5.4",
		CalicoAddonName:                    "calico/typha:v3.8.0",
		AzureNetworkPolicyAddonName:        "mcr.microsoft.com/containernetworking/azure-npm:v1.0.28",
		AADPodIdentityAddonName:            "mcr.microsoft.com/k8s/aad-pod-identity/nmi:1.2",
	}

	customAddonImages := make(map[string]string)
	for k := range defaultAddonImages {
		customAddonImages[k] = customImage
	}

	cases := []struct {
		name           string
		myAddons       []KubernetesAddon
		isUpdate       bool
		expectedImages map[string]string
	}{
		{
			name:           "default",
			myAddons:       getFakeAddons(defaultAddonImages, ""),
			isUpdate:       false,
			expectedImages: defaultAddonImages,
		},
		{
			name:           "create scenario",
			myAddons:       getFakeAddons(defaultAddonImages, customImage),
			isUpdate:       false,
			expectedImages: customAddonImages, // Image should not be overridden in create scenarios.
		},
		{
			name:           "upgrade + scale scenario",
			myAddons:       getFakeAddons(defaultAddonImages, customImage),
			isUpdate:       true,
			expectedImages: defaultAddonImages, // Image should be overridden in update scenarios.
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			mockCS := getMockBaseContainerService("1.10.8")
			mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
			mockCS.Properties.OrchestratorProfile.KubernetesConfig.Addons = c.myAddons
			mockCS.setOrchestratorDefaults(c.isUpdate, c.isUpdate)
			resultAddons := mockCS.Properties.OrchestratorProfile.KubernetesConfig.Addons
			for _, result := range resultAddons {
				if len(result.Containers) > 0 && result.Containers[0].Image != c.expectedImages[result.Name] {
					t.Errorf("expected setDefaults to set Image to \"%s\" in addon %s, but got \"%s\"", c.expectedImages[result.Name], result.Name, result.Containers[0].Image)
				}
			}
		})
	}
}

func getFakeAddons(defaultAddonMap map[string]string, customImage string) []KubernetesAddon {
	var fakeCustomAddons []KubernetesAddon
	for addonName := range defaultAddonMap {
		containerName := addonName
		if addonName == ContainerMonitoringAddonName {
			containerName = "omsagent"
		}
		if addonName == CalicoAddonName {
			containerName = "calico-typha"
		}
		if addonName == AADPodIdentityAddonName {
			containerName = "nmi"
		}
		customAddon := KubernetesAddon{
			Name:    addonName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:           containerName,
					CPURequests:    "50m",
					MemoryRequests: "150Mi",
					CPULimits:      "50m",
					MemoryLimits:   "150Mi",
				},
			},
		}
		if customImage != "" {
			customAddon.Containers[0].Image = customImage
		}
		fakeCustomAddons = append(fakeCustomAddons, customAddon)
	}
	return fakeCustomAddons
}

func TestAssignDefaultAddonVals(t *testing.T) {
	addonName := "testaddon"
	customImage := "myimage"
	customCPURequests := "60m"
	customMemoryRequests := "160Mi"
	customCPULimits := "40m"
	customMemoryLimits := "140Mi"
	// Verify that an addon with all custom values provided remains unmodified during default value assignment
	customAddon := KubernetesAddon{
		Name:    addonName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:           addonName,
				Image:          customImage,
				CPURequests:    customCPURequests,
				MemoryRequests: customMemoryRequests,
				CPULimits:      customCPULimits,
				MemoryLimits:   customMemoryLimits,
			},
		},
	}
	addonWithDefaults := getMockAddon(addonName)
	isUpdate := false
	modifiedAddon := assignDefaultAddonVals(customAddon, addonWithDefaults, isUpdate)
	if modifiedAddon.Containers[0].Name != customAddon.Containers[0].Name {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'Name' value %s to %s,", customAddon.Containers[0].Name, modifiedAddon.Containers[0].Name)
	}
	if modifiedAddon.Containers[0].Image != customAddon.Containers[0].Image {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'Image' value %s to %s,", customAddon.Containers[0].Image, modifiedAddon.Containers[0].Image)
	}
	if modifiedAddon.Containers[0].CPURequests != customAddon.Containers[0].CPURequests {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'CPURequests' value %s to %s,", customAddon.Containers[0].CPURequests, modifiedAddon.Containers[0].CPURequests)
	}
	if modifiedAddon.Containers[0].MemoryRequests != customAddon.Containers[0].MemoryRequests {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'MemoryRequests' value %s to %s,", customAddon.Containers[0].MemoryRequests, modifiedAddon.Containers[0].MemoryRequests)
	}
	if modifiedAddon.Containers[0].CPULimits != customAddon.Containers[0].CPULimits {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'CPULimits' value %s to %s,", customAddon.Containers[0].CPULimits, modifiedAddon.Containers[0].CPULimits)
	}
	if modifiedAddon.Containers[0].MemoryLimits != customAddon.Containers[0].MemoryLimits {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'MemoryLimits' value %s to %s,", customAddon.Containers[0].MemoryLimits, modifiedAddon.Containers[0].MemoryLimits)
	}

	// Verify that an addon with no custom values provided gets all the appropriate defaults
	customAddon = KubernetesAddon{
		Name:    addonName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name: addonName,
			},
		},
	}
	isUpdate = false
	modifiedAddon = assignDefaultAddonVals(customAddon, addonWithDefaults, isUpdate)
	if modifiedAddon.Containers[0].Image != addonWithDefaults.Containers[0].Image {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'Image' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].Image, modifiedAddon.Containers[0].Image)
	}
	if modifiedAddon.Containers[0].CPURequests != addonWithDefaults.Containers[0].CPURequests {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'CPURequests' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].CPURequests, modifiedAddon.Containers[0].CPURequests)
	}
	if modifiedAddon.Containers[0].MemoryRequests != addonWithDefaults.Containers[0].MemoryRequests {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'MemoryRequests' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].MemoryRequests, modifiedAddon.Containers[0].MemoryRequests)
	}
	if modifiedAddon.Containers[0].CPULimits != addonWithDefaults.Containers[0].CPULimits {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'CPULimits' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].CPULimits, modifiedAddon.Containers[0].CPULimits)
	}
	if modifiedAddon.Containers[0].MemoryLimits != addonWithDefaults.Containers[0].MemoryLimits {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'MemoryLimits' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].MemoryLimits, modifiedAddon.Containers[0].MemoryLimits)
	}

	// More checking to verify default interpolation
	customAddon = KubernetesAddon{
		Name:    addonName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:         addonName,
				CPURequests:  customCPURequests,
				MemoryLimits: customMemoryLimits,
			},
		},
	}
	isUpdate = false
	modifiedAddon = assignDefaultAddonVals(customAddon, addonWithDefaults, isUpdate)
	if modifiedAddon.Containers[0].Image != addonWithDefaults.Containers[0].Image {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'Image' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].Image, modifiedAddon.Containers[0].Image)
	}
	if modifiedAddon.Containers[0].Name != customAddon.Containers[0].Name {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'Name' value %s to %s,", customAddon.Containers[0].Name, modifiedAddon.Containers[0].Name)
	}
	if modifiedAddon.Containers[0].MemoryRequests != addonWithDefaults.Containers[0].MemoryRequests {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'MemoryRequests' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].MemoryRequests, modifiedAddon.Containers[0].MemoryRequests)
	}
	if modifiedAddon.Containers[0].CPULimits != addonWithDefaults.Containers[0].CPULimits {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'CPULimits' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].CPULimits, modifiedAddon.Containers[0].CPULimits)
	}
	if modifiedAddon.Containers[0].MemoryLimits != customAddon.Containers[0].MemoryLimits {
		t.Fatalf("assignDefaultAddonVals() should not have modified Containers 'MemoryLimits' value %s to %s,", customAddon.Containers[0].MemoryLimits, modifiedAddon.Containers[0].MemoryLimits)
	}

	// Verify that an addon with a custom image value will be overridden during upgrade/scale
	customAddon = KubernetesAddon{
		Name:    addonName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:  addonName,
				Image: customImage,
			},
		},
	}
	isUpdate = true
	modifiedAddon = assignDefaultAddonVals(customAddon, addonWithDefaults, isUpdate)
	if modifiedAddon.Containers[0].Image != addonWithDefaults.Containers[0].Image {
		t.Fatalf("assignDefaultAddonVals() should have assigned a default 'Image' value of %s, instead assigned %s,", addonWithDefaults.Containers[0].Image, modifiedAddon.Containers[0].Image)
	}

	addonWithDefaults.Config = map[string]string{
		"os":    "Linux",
		"taint": "node.kubernetes.io/memory-pressure",
	}
	isUpdate = false
	modifiedAddon = assignDefaultAddonVals(customAddon, addonWithDefaults, isUpdate)

	if modifiedAddon.Config["os"] != "Linux" {
		t.Error("assignDefaultAddonVals() should have added the default config property")
	}

	if modifiedAddon.Config["taint"] != "node.kubernetes.io/memory-pressure" {
		t.Error("assignDefaultAddonVals() should have added the default config property")
	}

	// Verify that an addon with a nil enabled inherits the default enabled value
	customAddon = KubernetesAddon{
		Name: addonName,
		Containers: []KubernetesContainerSpec{
			{
				Name:  addonName,
				Image: customImage,
			},
		},
	}
	isUpdate = false
	addonWithDefaults.Enabled = to.BoolPtr(true)
	modifiedAddon = assignDefaultAddonVals(customAddon, addonWithDefaults, isUpdate)
	if to.Bool(modifiedAddon.Enabled) != to.Bool(addonWithDefaults.Enabled) {
		t.Errorf("assignDefaultAddonVals() should have assigned a default 'Enabled' value of %t, instead assigned %t,", to.Bool(addonWithDefaults.Enabled), to.Bool(modifiedAddon.Enabled))
	}

	customAddon = KubernetesAddon{
		Name: addonName,
		Containers: []KubernetesContainerSpec{
			{
				Name:  addonName,
				Image: customImage,
			},
		},
	}
	isUpdate = false
	addonWithDefaults.Enabled = to.BoolPtr(false)
	modifiedAddon = assignDefaultAddonVals(customAddon, addonWithDefaults, isUpdate)
	if to.Bool(modifiedAddon.Enabled) != to.Bool(addonWithDefaults.Enabled) {
		t.Errorf("assignDefaultAddonVals() should have assigned a default 'Enabled' value of %t, instead assigned %t,", to.Bool(addonWithDefaults.Enabled), to.Bool(modifiedAddon.Enabled))
	}
}

func TestAcceleratedNetworking(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled = nil
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows = nil

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  true,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In upgrade scenario, nil AcceleratedNetworkingEnabled should always render as false (i.e., we never turn on this feature on an existing vm that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled) {
		t.Errorf("expected nil acceleratedNetworkingEnabled to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled))
	}
	// In upgrade scenario, nil AcceleratedNetworkingEnabledWindows should always render as false (i.e., we never turn on this feature on an existing vm that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows) {
		t.Errorf("expected nil acceleratedNetworkingEnabledWindows to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled = nil
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows = nil

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In scale scenario, nil AcceleratedNetworkingEnabled should always render as false (i.e., we never turn on this feature on an existing agent pool / VMSS that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled) {
		t.Errorf("expected nil acceleratedNetworkingEnabled to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled))
	}
	// In scale scenario, nil AcceleratedNetworkingEnabledWindows should always render as false (i.e., we never turn on this feature on an existing VM that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows) {
		t.Errorf("expected nil acceleratedNetworkingEnabledWindows to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled = nil
	mockCS.Properties.AgentPoolProfiles[0].VMSize = "Standard_D2_v2"
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows = nil
	mockCS.Properties.AgentPoolProfiles[0].VMSize = "Standard_D2_v2"

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In create scenario, nil AcceleratedNetworkingEnabled should be the defaults
	acceleratedNetworkingEnabled := DefaultAcceleratedNetworking
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled) != acceleratedNetworkingEnabled {
		t.Errorf("expected default acceleratedNetworkingEnabled to be %t, instead got %t", acceleratedNetworkingEnabled, to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled))
	}
	// In create scenario, nil AcceleratedNetworkingEnabledWindows should be the defaults
	acceleratedNetworkingEnabled = DefaultAcceleratedNetworkingWindowsEnabled
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows) != acceleratedNetworkingEnabled {
		t.Errorf("expected default acceleratedNetworkingEnabledWindows to be %t, instead got %t", acceleratedNetworkingEnabled, to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled = nil
	mockCS.Properties.AgentPoolProfiles[0].VMSize = "Standard_D666_v2"
	mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows = nil
	mockCS.Properties.AgentPoolProfiles[0].VMSize = "Standard_D666_v2"

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In non-supported VM SKU scenario, acceleratedNetworkingEnabled should always be false
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled) {
		t.Errorf("expected acceleratedNetworkingEnabled to be %t for an unsupported VM SKU, instead got %t", false, to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabled))
	}
	// In non-supported VM SKU scenario, acceleratedNetworkingEnabledWindows should always be false
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows) {
		t.Errorf("expected acceleratedNetworkingEnabledWindows to be %t for an unsupported VM SKU, instead got %t", false, to.Bool(mockCS.Properties.AgentPoolProfiles[0].AcceleratedNetworkingEnabledWindows))
	}
}

func TestVMSSOverProvisioning(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
	mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled = nil
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  true,
		PkiKeySize: 4096,
	})

	// In upgrade scenario, nil AcceleratedNetworkingEnabled should always render as false (i.e., we never turn on this feature on an existing vm that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled) {
		t.Errorf("expected nil VMSSOverProvisioningEnabled to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
	mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled = nil
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In scale scenario, nil VMSSOverProvisioningEnabled should always render as false (i.e., we never turn on this feature on an existing agent pool / VMSS that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled) {
		t.Errorf("expected nil VMSSOverProvisioningEnabled to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
	mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled = nil
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In create scenario, nil VMSSOverProvisioningEnabled should be the defaults
	vmssOverProvisioningEnabled := DefaultVMSSOverProvisioningEnabled
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled) != vmssOverProvisioningEnabled {
		t.Errorf("expected default VMSSOverProvisioningEnabled to be %t, instead got %t", vmssOverProvisioningEnabled, to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
	mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled = to.BoolPtr(true)
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In create scenario with explicit true, VMSSOverProvisioningEnabled should be true
	if !to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled) {
		t.Errorf("expected VMSSOverProvisioningEnabled to be true, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
	mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled = to.BoolPtr(false)
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	// In create scenario with explicit false, VMSSOverProvisioningEnabled should be false
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled) {
		t.Errorf("expected VMSSOverProvisioningEnabled to be false, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].VMSSOverProvisioningEnabled))
	}
}

func TestAuditDEnabled(t *testing.T) {
	mockCS := getMockBaseContainerService("1.12.7")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	isUpgrade := true
	mockCS.Properties.setAgentProfileDefaults(isUpgrade, false, AzurePublicCloud)

	// In upgrade scenario, nil AuditDEnabled should always render as false (i.e., we never turn on this feature on an existing vm that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled) {
		t.Errorf("expected nil AuditDEnabled to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled))
	}

	mockCS = getMockBaseContainerService("1.12.7")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	isScale := true
	mockCS.Properties.setAgentProfileDefaults(false, isScale, AzurePublicCloud)

	// In scale scenario, nil AuditDEnabled should always render as false (i.e., we never turn on this feature on an existing agent pool / vms that didn't have it before)
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled) {
		t.Errorf("expected nil AuditDEnabled to be false after upgrade, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled))
	}

	mockCS = getMockBaseContainerService("1.12.7")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.setAgentProfileDefaults(false, false, AzurePublicCloud)

	// In create scenario, nil AuditDEnabled should be the defaults
	auditDEnabledEnabled := DefaultAuditDEnabled
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled) != auditDEnabledEnabled {
		t.Errorf("expected default AuditDEnabled to be %t, instead got %t", auditDEnabledEnabled, to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled = to.BoolPtr(true)
	mockCS.Properties.setAgentProfileDefaults(false, false, AzurePublicCloud)

	// In create scenario with explicit true, AuditDEnabled should be true
	if !to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled) {
		t.Errorf("expected AuditDEnabled to be true, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled))
	}

	mockCS = getMockBaseContainerService("1.10.8")
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled = to.BoolPtr(false)
	mockCS.Properties.setAgentProfileDefaults(false, false, AzurePublicCloud)

	// In create scenario with explicit false, AuditDEnabled should be false
	if to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled) {
		t.Errorf("expected AuditDEnabled to be false, instead got %t", to.Bool(mockCS.Properties.AgentPoolProfiles[0].AuditDEnabled))
	}
}

func TestKubeletFeatureGatesEnsureFeatureGatesOnAgentsFor1_6_0(t *testing.T) {
	mockCS := getMockBaseContainerService("1.6.0")
	properties := mockCS.Properties

	// No KubernetesConfig.KubeletConfig set for MasterProfile or AgentProfile
	// so they will inherit the top-level config
	properties.OrchestratorProfile.KubernetesConfig = getKubernetesConfigWithFeatureGates("TopLevel=true")

	mockCS.setKubeletConfig(false)

	agentFeatureGates := properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig["--feature-gates"]
	if agentFeatureGates != "TopLevel=true" {
		t.Fatalf("setKubeletConfig did not add 'TopLevel=true' for agent profile: expected 'TopLevel=true' got '%s'", agentFeatureGates)
	}

	// Verify that the TopLevel feature gate override has only been applied to the agents
	masterFeatureFates := properties.MasterProfile.KubernetesConfig.KubeletConfig["--feature-gates"]
	if masterFeatureFates != "TopLevel=true" {
		t.Fatalf("setKubeletConfig modified feature gates for master profile: expected 'TopLevel=true' got '%s'", agentFeatureGates)
	}
}

func TestKubeletFeatureGatesEnsureMasterAndAgentConfigUsedFor1_6_0(t *testing.T) {
	mockCS := getMockBaseContainerService("1.6.0")
	properties := mockCS.Properties

	// Set MasterProfile and AgentProfiles KubernetesConfig.KubeletConfig values
	// Verify that they are used instead of the top-level config
	properties.OrchestratorProfile.KubernetesConfig = getKubernetesConfigWithFeatureGates("TopLevel=true")
	properties.MasterProfile = &MasterProfile{KubernetesConfig: getKubernetesConfigWithFeatureGates("MasterLevel=true")}
	properties.AgentPoolProfiles[0].KubernetesConfig = getKubernetesConfigWithFeatureGates("AgentLevel=true")

	mockCS.setKubeletConfig(false)

	agentFeatureGates := properties.AgentPoolProfiles[0].KubernetesConfig.KubeletConfig["--feature-gates"]
	if agentFeatureGates != "AgentLevel=true" {
		t.Fatalf("setKubeletConfig agent profile: expected 'AgentLevel=true' got '%s'", agentFeatureGates)
	}

	// Verify that the TopLevel feature gate override has only been applied to the agents
	masterFeatureFates := properties.MasterProfile.KubernetesConfig.KubeletConfig["--feature-gates"]
	if masterFeatureFates != "MasterLevel=true" {
		t.Fatalf("setKubeletConfig master profile: expected 'MasterLevel=true' got '%s'", agentFeatureGates)
	}
}

func TestEtcdDiskSize(t *testing.T) {
	mockCS := getMockBaseContainerService("1.8.10")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != DefaultEtcdDiskSize {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, DefaultEtcdDiskSize)
	}

	mockCS = getMockBaseContainerService("1.8.10")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 5
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != DefaultEtcdDiskSizeGT3Nodes {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, DefaultEtcdDiskSizeGT3Nodes)
	}

	mockCS = getMockBaseContainerService("1.8.10")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 5
	properties.AgentPoolProfiles[0].Count = 6
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != DefaultEtcdDiskSizeGT10Nodes {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, DefaultEtcdDiskSizeGT10Nodes)
	}

	mockCS = getMockBaseContainerService("1.8.10")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 5
	properties.AgentPoolProfiles[0].Count = 16
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != DefaultEtcdDiskSizeGT20Nodes {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, DefaultEtcdDiskSizeGT20Nodes)
	}

	mockCS = getMockBaseContainerService("1.8.10")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 5
	properties.AgentPoolProfiles[0].Count = 50
	customEtcdDiskSize := "512"
	properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB = customEtcdDiskSize
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != customEtcdDiskSize {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, customEtcdDiskSize)
	}
}

func TestGenerateEtcdEncryptionKey(t *testing.T) {
	key1 := generateEtcdEncryptionKey()
	key2 := generateEtcdEncryptionKey()
	if key1 == key2 {
		t.Fatalf("generateEtcdEncryptionKey should return a unique key each time, instead returned identical %s and %s", key1, key2)
	}
	for _, val := range []string{key1, key2} {
		_, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			t.Fatalf("generateEtcdEncryptionKey should return a base64 encoded key, instead returned %s", val)
		}
	}
}

func TestNetworkPolicyDefaults(t *testing.T) {
	mockCS := getMockBaseContainerService("1.8.10")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyCalico
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin != NetworkPluginKubenet {
		t.Fatalf("NetworkPlugin did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin, NetworkPluginKubenet)
	}

	mockCS = getMockBaseContainerService("1.8.10")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyCilium
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin != NetworkPluginCilium {
		t.Fatalf("NetworkPlugin did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin, NetworkPluginCilium)
	}

	mockCS = getMockBaseContainerService("1.8.10")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyAzure
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin != NetworkPluginAzure {
		t.Fatalf("NetworkPlugin did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin, NetworkPluginAzure)
	}
	if properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy != "" {
		t.Fatalf("NetworkPolicy did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy, "")
	}

	mockCS = getMockBaseContainerService("1.8.10")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyNone
	mockCS.setOrchestratorDefaults(true, true)
	if properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin != NetworkPluginKubenet {
		t.Fatalf("NetworkPlugin did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin, NetworkPluginKubenet)
	}
	if properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy != "" {
		t.Fatalf("NetworkPolicy did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy, "")
	}
}

func TestContainerRuntime(t *testing.T) {

	for _, mobyVersion := range []string{"3.0.1", "3.0.3", "3.0.4", "3.0.5", "3.0.6", "3.0.7"} {
		mockCS := getMockBaseContainerService("1.10.13")
		properties := mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.MobyVersion = mobyVersion
		mockCS.setOrchestratorDefaults(true, true)
		if properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime != Docker {
			t.Fatalf("ContainerRuntime did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime, Docker)
		}
		if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != DefaultMobyVersion {
			t.Fatalf("MobyVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.MobyVersion, DefaultMobyVersion)
		}
		if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != "" {
			t.Fatalf("Containerd did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion, "")
		}

		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.MobyVersion = mobyVersion
		mockCS.setOrchestratorDefaults(false, false)
		if properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime != Docker {
			t.Fatalf("ContainerRuntime did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime, Docker)
		}
		if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != mobyVersion {
			t.Fatalf("MobyVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.MobyVersion, mobyVersion)
		}
		if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != "" {
			t.Fatalf("Containerd did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion, "")
		}
	}

	mockCS := getMockBaseContainerService("1.10.13")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.setOrchestratorDefaults(false, false)
	if properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime != Docker {
		t.Fatalf("ContainerRuntime did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime, Docker)
	}
	if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != DefaultMobyVersion {
		t.Fatalf("MobyVersion did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.MobyVersion, DefaultMobyVersion)
	}
	if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != "" {
		t.Fatalf("Containerd did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion, "")
	}

	mockCS = getMockBaseContainerService("1.10.13")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime = Containerd
	mockCS.setOrchestratorDefaults(false, false)
	if properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime != Containerd {
		t.Fatalf("ContainerRuntime did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime, Containerd)
	}
	if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != "" {
		t.Fatalf("MobyVersion did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.MobyVersion, "")
	}
	if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != DefaultContainerdVersion {
		t.Fatalf("Containerd did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion, DefaultContainerdVersion)
	}

	mockCS = getMockBaseContainerService("1.10.13")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime = KataContainers
	mockCS.setOrchestratorDefaults(false, false)
	if properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime != KataContainers {
		t.Fatalf("ContainerRuntime did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime, KataContainers)
	}
	if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != "" {
		t.Fatalf("MobyVersion did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.MobyVersion, "")
	}
	if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != DefaultContainerdVersion {
		t.Fatalf("Containerd did not have the expected value, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion, DefaultContainerdVersion)
	}

	for _, containerdVersion := range []string{"1.1.2", "1.1.4", "1.1.5"} {

		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime = Containerd
		properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion = containerdVersion
		mockCS.setOrchestratorDefaults(true, true)
		if properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime != Containerd {
			t.Fatalf("ContainerRuntime did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime, Containerd)
		}
		if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != "" {
			t.Fatalf("MobyVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.MobyVersion, "")
		}
		if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != DefaultContainerdVersion {
			t.Fatalf("Containerd did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion, DefaultContainerdVersion)
		}

		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime = Containerd
		properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion = containerdVersion
		mockCS.setOrchestratorDefaults(false, false)
		if properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime != Containerd {
			t.Fatalf("ContainerRuntime did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime, Containerd)
		}
		if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != "" {
			t.Fatalf("MobyVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.MobyVersion, "")
		}
		if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != containerdVersion {
			t.Fatalf("Containerd did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion, containerdVersion)
		}
	}
}

func TestEtcdVersion(t *testing.T) {
	// Default (no value) scenario
	for _, etcdVersion := range []string{""} {
		// Upgrade scenario should always upgrade to newer, default etcd version
		// This sort of artificial (upgrade scenario should always have value), but strictly speaking this is what we want to do
		mockCS := getMockBaseContainerService("1.10.13")
		properties := mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(true, false)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != DefaultEtcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, DefaultEtcdVersion)
		}

		// Create scenario should always accept the provided value
		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(false, false)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != DefaultEtcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, DefaultEtcdVersion)
		}

		// Scale scenario should always accept the provided value
		// This sort of artificial (upgrade scenario should always have value), but strictly speaking this is what we want to do
		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(false, true)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != DefaultEtcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, DefaultEtcdVersion)
		}
	}

	// These versions are all less than or equal to default
	for _, etcdVersion := range []string{"2.2.5", "3.2.24", DefaultEtcdVersion} {
		// Upgrade scenario should always upgrade to newer, default etcd version
		mockCS := getMockBaseContainerService("1.10.13")
		properties := mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(true, false)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != DefaultEtcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, DefaultEtcdVersion)
		}

		// Create scenario should always accept the provided value
		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(false, false)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != etcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, etcdVersion)
		}

		// Scale scenario should always accept the provided value
		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(false, true)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != etcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, etcdVersion)
		}
	}

	// These versions are all greater than default
	for _, etcdVersion := range []string{"3.4.0", "99.99"} {
		// Upgrade scenario should always keep the user-configured etcd version if it is greater than default
		mockCS := getMockBaseContainerService("1.10.13")
		properties := mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(true, false)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != etcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, etcdVersion)
		}

		// Create scenario should always accept the provided value
		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(false, false)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != etcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, etcdVersion)
		}

		// Scale scenario should always accept the provided value
		mockCS = getMockBaseContainerService("1.10.13")
		properties = mockCS.Properties
		properties.OrchestratorProfile.OrchestratorType = Kubernetes
		properties.OrchestratorProfile.KubernetesConfig.EtcdVersion = etcdVersion
		mockCS.setOrchestratorDefaults(false, true)
		if properties.OrchestratorProfile.KubernetesConfig.EtcdVersion != etcdVersion {
			t.Fatalf("EtcdVersion did not have the expected value, got %s, expected %s",
				properties.OrchestratorProfile.KubernetesConfig.EtcdVersion, etcdVersion)
		}
	}
}

func TestStorageProfile(t *testing.T) {
	// Test ManagedDisks default configuration
	mockCS := getMockBaseContainerService("1.8.10")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	properties.OrchestratorProfile.KubernetesConfig.PrivateCluster = &PrivateCluster{
		Enabled:        to.BoolPtr(true),
		JumpboxProfile: &PrivateJumpboxProfile{},
	}
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.MasterProfile.StorageProfile != ManagedDisks {
		t.Fatalf("MasterProfile.StorageProfile did not have the expected configuration, got %s, expected %s",
			properties.MasterProfile.StorageProfile, ManagedDisks)
	}
	if !properties.MasterProfile.IsManagedDisks() {
		t.Fatalf("MasterProfile.StorageProfile did not have the expected configuration, got %t, expected %t",
			false, true)
	}
	if properties.AgentPoolProfiles[0].StorageProfile != ManagedDisks {
		t.Fatalf("AgentPoolProfile.StorageProfile did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].StorageProfile, ManagedDisks)
	}
	if !properties.AgentPoolProfiles[0].IsManagedDisks() {
		t.Fatalf("AgentPoolProfile.IsManagedDisks() did not have the expected configuration, got %t, expected %t",
			false, true)
	}
	if properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile != ManagedDisks {
		t.Fatalf("MasterProfile.StorageProfile did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile, ManagedDisks)
	}
	if !properties.AgentPoolProfiles[0].IsAvailabilitySets() {
		t.Fatalf("AgentPoolProfile[0].AvailabilityProfile did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].AvailabilityProfile, AvailabilitySet)
	}

	mockCS = getMockBaseContainerService("1.10.2")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if !properties.AgentPoolProfiles[0].IsVirtualMachineScaleSets() {
		t.Fatalf("AgentPoolProfile[0].AvailabilityProfile did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].AvailabilityProfile, VirtualMachineScaleSets)
	}

}

// TestMasterProfileDefaults covers tests for setMasterProfileDefaults
func TestMasterProfileDefaults(t *testing.T) {
	// this validates default masterProfile configuration
	mockCS := getMockBaseContainerService("1.10.3")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = ""
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	properties.MasterProfile.AvailabilityProfile = ""
	properties.MasterProfile.Count = 3
	mockCS.Properties = properties
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.MasterProfile.IsVirtualMachineScaleSets() {
		t.Fatalf("Master VMAS, AzureCNI: MasterProfile AvailabilityProfile did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.AvailabilityProfile, AvailabilitySet)
	}
	if properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet != DefaultKubernetesSubnet {
		t.Fatalf("Master VMAS, AzureCNI: MasterProfile ClusterSubnet did not have the expected default configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet, DefaultKubernetesSubnet)
	}
	if properties.MasterProfile.Subnet != properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet {
		t.Fatalf("Master VMAS, AzureCNI: MasterProfile Subnet did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.Subnet, properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet)
	}
	if properties.AgentPoolProfiles[0].Subnet != properties.MasterProfile.Subnet {
		t.Fatalf("Master VMAS, AzureCNI: AgentPoolProfiles Subnet did not have the expected default configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].Subnet, properties.MasterProfile.Subnet)
	}
	if properties.MasterProfile.FirstConsecutiveStaticIP != "10.255.255.5" {
		t.Fatalf("Master VMAS, AzureCNI: MasterProfile FirstConsecutiveStaticIP did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.FirstConsecutiveStaticIP, "10.255.255.5")
	}

	// this validates default VMSS masterProfile configuration
	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = ""
	properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  true,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if !properties.MasterProfile.IsVirtualMachineScaleSets() {
		t.Fatalf("Master VMSS, AzureCNI: MasterProfile AvailabilityProfile did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.AvailabilityProfile, VirtualMachineScaleSets)
	}
	if properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet != DefaultKubernetesSubnet {
		t.Fatalf("Master VMSS, AzureCNI: MasterProfile ClusterSubnet did not have the expected default configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet, DefaultKubernetesSubnet)
	}
	if properties.MasterProfile.FirstConsecutiveStaticIP != DefaultFirstConsecutiveKubernetesStaticIPVMSS {
		t.Fatalf("Master VMSS, AzureCNI: MasterProfile FirstConsecutiveStaticIP did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.FirstConsecutiveStaticIP, DefaultFirstConsecutiveKubernetesStaticIPVMSS)
	}
	if properties.MasterProfile.Subnet != DefaultKubernetesMasterSubnet {
		t.Fatalf("Master VMSS, AzureCNI: MasterProfile Subnet did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.Subnet, DefaultKubernetesMasterSubnet)
	}
	if properties.MasterProfile.AgentSubnet != DefaultKubernetesAgentSubnetVMSS {
		t.Fatalf("Master VMSS, AzureCNI: MasterProfile AgentSubnet did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.AgentSubnet, DefaultKubernetesAgentSubnetVMSS)
	}

	// this validates default masterProfile configuration and kubenet
	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = ""
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
	properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet != DefaultKubernetesClusterSubnet {
		t.Fatalf("Master VMSS, kubenet: MasterProfile ClusterSubnet did not have the expected default configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet, DefaultKubernetesClusterSubnet)
	}
	if properties.MasterProfile.Subnet != DefaultKubernetesMasterSubnet {
		t.Fatalf("Master VMSS, kubenet: MasterProfile Subnet did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.Subnet, DefaultKubernetesMasterSubnet)
	}
	if properties.MasterProfile.FirstConsecutiveStaticIP != DefaultFirstConsecutiveKubernetesStaticIPVMSS {
		t.Fatalf("Master VMSS, kubenet: MasterProfile FirstConsecutiveStaticIP did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.FirstConsecutiveStaticIP, DefaultFirstConsecutiveKubernetesStaticIPVMSS)
	}
	if properties.MasterProfile.AgentSubnet != DefaultKubernetesAgentSubnetVMSS {
		t.Fatalf("Master VMSS, kubenet: MasterProfile AgentSubnet did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.AgentSubnet, DefaultKubernetesAgentSubnetVMSS)
	}
	properties.MasterProfile.AvailabilityProfile = AvailabilitySet
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.MasterProfile.FirstConsecutiveStaticIP != DefaultFirstConsecutiveKubernetesStaticIP {
		t.Fatalf("Master VMAS, kubenet: MasterProfile FirstConsecutiveStaticIP did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.FirstConsecutiveStaticIP, DefaultFirstConsecutiveKubernetesStaticIP)
	}

	// this validates default vmas masterProfile configuration, AzureCNI, and custom vnet
	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.VnetSubnetID = "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/ExampleCustomVNET/subnets/ExampleMasterSubnet"
	properties.MasterProfile.VnetCidr = "10.239.0.0/16"
	properties.MasterProfile.FirstConsecutiveStaticIP = "10.239.255.239"
	properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = ""
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	properties.MasterProfile.AvailabilityProfile = AvailabilitySet
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if properties.MasterProfile.FirstConsecutiveStaticIP != "10.239.255.239" {
		t.Fatalf("Master VMAS, AzureCNI, customvnet: MasterProfile FirstConsecutiveStaticIP did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.FirstConsecutiveStaticIP, "10.239.255.239")
	}

	// this validates default VMSS masterProfile configuration, AzureCNI, and custom VNET
	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.VnetSubnetID = "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/ExampleCustomVNET/subnets/ExampleMasterSubnet"
	properties.MasterProfile.VnetCidr = "10.239.0.0/16"
	properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = ""
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.MasterProfile.FirstConsecutiveStaticIP != "10.239.0.4" {
		t.Fatalf("Master VMSS, AzureCNI, customvnet: MasterProfile FirstConsecutiveStaticIP did not have the expected default configuration, got %s, expected %s",
			properties.MasterProfile.FirstConsecutiveStaticIP, "10.239.0.4")
	}

	// this validates default configurations for LoadBalancerSku and ExcludeMasterFromStandardLB
	mockCS = getMockBaseContainerService("1.11.6")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku = StandardLoadBalancerSku
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	excludeMaster := DefaultExcludeMasterFromStandardLB
	if *properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB != excludeMaster {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB did not have the expected configuration, got %t, expected %t",
			*properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB, excludeMaster)
	}

	// this validates default configurations for MaximumLoadBalancerRuleCount.
	mockCS = getMockBaseContainerService("1.11.6")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.OrchestratorProfile.KubernetesConfig.MaximumLoadBalancerRuleCount != DefaultMaximumLoadBalancerRuleCount {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.MaximumLoadBalancerRuleCount did not have the expected configuration, got %d, expected %d",
			properties.OrchestratorProfile.KubernetesConfig.MaximumLoadBalancerRuleCount, DefaultMaximumLoadBalancerRuleCount)
	}

	// this validates cluster subnet default configuration for dual stack feature.
	mockCS = getMockBaseContainerService("1.15.0-beta.1")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.FeatureFlags = &FeatureFlags{EnableIPv6DualStack: true}
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	expectedClusterSubnet := strings.Join([]string{DefaultKubernetesClusterSubnet, DefaultKubernetesClusterSubnetIPv6}, ",")
	if properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet != expectedClusterSubnet {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.ClusterSubnet did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet, expectedClusterSubnet)
	}

	// this validates cluster subnet default configuration for dual stack feature when only ipv4 subnet provided
	mockCS = getMockBaseContainerService("1.15.0-beta.1")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = "10.244.0.0/16"
	properties.FeatureFlags = &FeatureFlags{EnableIPv6DualStack: true}
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	expectedClusterSubnet = strings.Join([]string{"10.244.0.0/16", DefaultKubernetesClusterSubnetIPv6}, ",")
	if properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet != expectedClusterSubnet {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.ClusterSubnet did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet, expectedClusterSubnet)
	}

	// this validates cluster subnet default configuration for dual stack feature when only ipv6 subnet provided
	mockCS = getMockBaseContainerService("1.15.0-beta.1")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet = "ace:cab:deca::/8"
	properties.FeatureFlags = &FeatureFlags{EnableIPv6DualStack: true}
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	expectedClusterSubnet = strings.Join([]string{DefaultKubernetesClusterSubnet, "ace:cab:deca::/8"}, ",")
	if properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet != expectedClusterSubnet {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.ClusterSubnet did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.ClusterSubnet, expectedClusterSubnet)
	}

	// this validates default configurations for OutboundRuleIdleTimeoutInMinutes.
	mockCS = getMockBaseContainerService("1.14.4")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku = StandardLoadBalancerSku
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.OrchestratorProfile.KubernetesConfig.OutboundRuleIdleTimeoutInMinutes != DefaultOutboundRuleIdleTimeoutInMinutes {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.OutboundRuleIdleTimeoutInMinutes did not have the expected configuration, got %d, expected %d",
			properties.OrchestratorProfile.KubernetesConfig.OutboundRuleIdleTimeoutInMinutes, DefaultOutboundRuleIdleTimeoutInMinutes)
	}
}

func TestAgentPoolProfile(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.AgentPoolProfiles[0].ScaleSetPriority != "" {
		t.Fatalf("AgentPoolProfiles[0].ScaleSetPriority did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].ScaleSetPriority, "")
	}
	if properties.AgentPoolProfiles[0].ScaleSetEvictionPolicy != "" {
		t.Fatalf("AgentPoolProfiles[0].ScaleSetEvictionPolicy did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].ScaleSetEvictionPolicy, "")
	}
	properties.AgentPoolProfiles[0].ScaleSetPriority = ScaleSetPriorityLow
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.AgentPoolProfiles[0].ScaleSetEvictionPolicy != ScaleSetEvictionPolicyDelete {
		t.Fatalf("AgentPoolProfile[0].ScaleSetEvictionPolicy did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].ScaleSetEvictionPolicy, ScaleSetEvictionPolicyDelete)
	}
}

// TestDistroDefaults covers tests for setMasterProfileDefaults and setAgentProfileDefaults
func TestDistroDefaults(t *testing.T) {

	var tests = []struct {
		name                   string              // test case name
		orchestratorProfile    OrchestratorProfile // orchestrator to be tested
		masterProfileDistro    Distro
		agentPoolProfileDistro Distro
		expectedAgentDistro    Distro // expected agent result default disto to be used
		expectedMasterDistro   Distro // expected master result default disto to be used
		isUpgrade              bool
		isScale                bool
		cloudName              string
	}{
		{
			"default_kubernetes",
			OrchestratorProfile{
				OrchestratorType: Kubernetes,
			},
			"",
			"",
			AKSUbuntu1604,
			AKSUbuntu1604,
			false,
			false,
			AzurePublicCloud,
		},
		{
			"default_kubernetes_usgov",
			OrchestratorProfile{
				OrchestratorType: Kubernetes,
			},
			"",
			"",
			AKSUbuntu1604,
			AKSUbuntu1604,
			false,
			false,
			AzureUSGovernmentCloud,
		},
		{
			"1804_upgrade_kubernetes",
			OrchestratorProfile{
				OrchestratorType: Kubernetes,
			},
			AKSUbuntu1804,
			AKSUbuntu1804,
			AKSUbuntu1804,
			AKSUbuntu1804,
			true,
			false,
			AzurePublicCloud,
		},
		{
			"default_kubernetes_usgov",
			OrchestratorProfile{
				OrchestratorType: Kubernetes,
			},
			AKS1604Deprecated,
			AKS1604Deprecated,
			Ubuntu,
			Ubuntu,
			true,
			false,
			AzureGermanCloud,
		},
		{
			"deprecated_distro_kubernetes",
			OrchestratorProfile{
				OrchestratorType: Kubernetes,
			},
			AKS1604Deprecated,
			AKS1604Deprecated,
			AKSUbuntu1604,
			AKSUbuntu1604,
			true,
			false,
			AzureChinaCloud,
		},
		{
			"docker_engine_kubernetes",
			OrchestratorProfile{
				OrchestratorType: Kubernetes,
			},
			AKS1604Deprecated,
			AKSDockerEngine,
			AKSUbuntu1604,
			AKSUbuntu1604,
			false,
			true,
			AzurePublicCloud,
		},
		{
			"default_swarm",
			OrchestratorProfile{
				OrchestratorType: Swarm,
			},
			"",
			"",
			Ubuntu,
			Ubuntu,
			false,
			false,
			AzurePublicCloud,
		},
		{
			"default_swarmmode",
			OrchestratorProfile{
				OrchestratorType: SwarmMode,
			},
			"",
			"",
			Ubuntu,
			Ubuntu,
			false,
			false,
			AzurePublicCloud,
		},
		{
			"default_dcos",
			OrchestratorProfile{
				OrchestratorType: DCOS,
			},
			"",
			"",
			Ubuntu,
			Ubuntu,
			false,
			false,
			AzurePublicCloud,
		},
	}

	for _, test := range tests {
		mockAPI := getMockAPIProperties("1.0.0")
		mockAPI.OrchestratorProfile = &test.orchestratorProfile
		mockAPI.MasterProfile.Distro = test.masterProfileDistro
		for _, agent := range mockAPI.AgentPoolProfiles {
			agent.Distro = test.agentPoolProfileDistro
		}
		mockAPI.setMasterProfileDefaults(test.isUpgrade, test.isScale, test.cloudName)
		mockAPI.setAgentProfileDefaults(test.isUpgrade, test.isScale, test.cloudName)
		if mockAPI.MasterProfile.Distro != test.expectedMasterDistro {
			t.Fatalf("setMasterProfileDefaults() test case %v did not return right Distro configurations %v != %v", test.name, mockAPI.MasterProfile.Distro, test.expectedMasterDistro)
		}
		for _, agent := range mockAPI.AgentPoolProfiles {
			if agent.Distro != test.expectedAgentDistro {
				t.Fatalf("setAgentProfileDefaults() test case %v did not return right Distro configurations %v != %v", test.name, agent.Distro, test.expectedAgentDistro)
			}
		}
	}
}

func TestWindowsProfileDefaults(t *testing.T) {

	var tests = []struct {
		name                   string // test case name
		windowsProfile         WindowsProfile
		expectedWindowsProfile WindowsProfile
		isAzureStack           bool
	}{
		{
			"defaults",
			WindowsProfile{},
			WindowsProfile{
				WindowsPublisher:      AKSWindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:          AKSWindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:            AKSWindowsServer2019OSImageConfig.ImageSku,
				ImageVersion:          AKSWindowsServer2019OSImageConfig.ImageVersion,
				AdminUsername:         "",
				AdminPassword:         "",
				WindowsImageSourceURL: "",
				WindowsDockerVersion:  "",
				SSHEnabled:            false,
			},
			false,
		},
		{
			"aks vhd current version",
			WindowsProfile{
				WindowsPublisher: AKSWindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:     AKSWindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:       AKSWindowsServer2019OSImageConfig.ImageSku,
			},
			WindowsProfile{
				WindowsPublisher:      AKSWindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:          AKSWindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:            AKSWindowsServer2019OSImageConfig.ImageSku,
				ImageVersion:          AKSWindowsServer2019OSImageConfig.ImageVersion,
				AdminUsername:         "",
				AdminPassword:         "",
				WindowsImageSourceURL: "",
				WindowsDockerVersion:  "",
				SSHEnabled:            false,
			},
			false,
		},
		{
			"aks vhd specific version",
			WindowsProfile{
				WindowsPublisher: AKSWindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:     AKSWindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:       AKSWindowsServer2019OSImageConfig.ImageSku,
				ImageVersion:     "override",
			},
			WindowsProfile{
				WindowsPublisher:      AKSWindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:          AKSWindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:            AKSWindowsServer2019OSImageConfig.ImageSku,
				ImageVersion:          "override",
				AdminUsername:         "",
				AdminPassword:         "",
				WindowsImageSourceURL: "",
				WindowsDockerVersion:  "",
				SSHEnabled:            false,
			},
			false,
		},
		{
			"vanilla vhd current version",
			WindowsProfile{
				WindowsPublisher: WindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:     WindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:       WindowsServer2019OSImageConfig.ImageSku,
			},
			WindowsProfile{
				WindowsPublisher:      WindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:          WindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:            WindowsServer2019OSImageConfig.ImageSku,
				ImageVersion:          WindowsServer2019OSImageConfig.ImageVersion,
				AdminUsername:         "",
				AdminPassword:         "",
				WindowsImageSourceURL: "",
				WindowsDockerVersion:  "",
				SSHEnabled:            false,
			},
			false,
		},
		{
			"vanilla vhd spepcific version",
			WindowsProfile{
				WindowsPublisher: WindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:     WindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:       WindowsServer2019OSImageConfig.ImageSku,
				ImageVersion:     "override",
			},
			WindowsProfile{
				WindowsPublisher:      WindowsServer2019OSImageConfig.ImagePublisher,
				WindowsOffer:          WindowsServer2019OSImageConfig.ImageOffer,
				WindowsSku:            WindowsServer2019OSImageConfig.ImageSku,
				ImageVersion:          "override",
				AdminUsername:         "",
				AdminPassword:         "",
				WindowsImageSourceURL: "",
				WindowsDockerVersion:  "",
				SSHEnabled:            false,
			},
			false,
		},
		{
			"user overrides latest version",
			WindowsProfile{
				WindowsPublisher: "override",
				WindowsOffer:     "override",
				WindowsSku:       "override",
			},
			WindowsProfile{
				WindowsPublisher:      "override",
				WindowsOffer:          "override",
				WindowsSku:            "override",
				ImageVersion:          "latest",
				AdminUsername:         "",
				AdminPassword:         "",
				WindowsImageSourceURL: "",
				WindowsDockerVersion:  "",
				SSHEnabled:            false,
			},
			false,
		},
		{
			"user overrides specific version",
			WindowsProfile{
				WindowsPublisher: "override",
				WindowsOffer:     "override",
				WindowsSku:       "override",
				ImageVersion:     "override",
			},
			WindowsProfile{
				WindowsPublisher:      "override",
				WindowsOffer:          "override",
				WindowsSku:            "override",
				ImageVersion:          "override",
				AdminUsername:         "",
				AdminPassword:         "",
				WindowsImageSourceURL: "",
				WindowsDockerVersion:  "",
				SSHEnabled:            false,
			},
			false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mockAPI := getMockAPIProperties("1.16.0")
			mockAPI.WindowsProfile = &test.windowsProfile
			if test.isAzureStack {
				mockAPI.CustomCloudProfile = &CustomCloudProfile{}
			}
			mockAPI.setWindowsProfileDefaults(false, false)

			actual := mockAPI.WindowsProfile
			expected := &test.expectedWindowsProfile

			diff := cmp.Diff(actual, expected)
			if diff != "" {
				t.Errorf("unexpected diff while comparing WindowsProfile: %s", diff)
			}
		})
	}
}

func TestIsAzureCNINetworkmonitorAddon(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10.3")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	properties.OrchestratorProfile.KubernetesConfig.Addons = []KubernetesAddon{
		{
			Name: AzureCNINetworkMonitoringAddonName,
			Containers: []KubernetesContainerSpec{
				{
					Name:           AzureCNINetworkMonitoringAddonName,
					CPURequests:    "50m",
					MemoryRequests: "150Mi",
					CPULimits:      "50m",
					MemoryLimits:   "150Mi",
				},
			},
			Enabled: to.BoolPtr(true),
		},
	}
	mockCS.setOrchestratorDefaults(true, true)

	i := getAddonsIndexByName(properties.OrchestratorProfile.KubernetesConfig.Addons, AzureCNINetworkMonitoringAddonName)
	if !to.Bool(properties.OrchestratorProfile.KubernetesConfig.Addons[i].Enabled) {
		t.Fatalf("Azure CNI networkmonitor addon should be present")
	}

	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	mockCS.setOrchestratorDefaults(true, true)

	i = getAddonsIndexByName(properties.OrchestratorProfile.KubernetesConfig.Addons, AzureCNINetworkMonitoringAddonName)
	if !to.Bool(properties.OrchestratorProfile.KubernetesConfig.Addons[i].Enabled) {
		t.Fatalf("Azure CNI networkmonitor addon should be present by default if Azure CNI is set")
	}
}

// TestSetVMSSDefaultsAndZones covers tests for setVMSSDefaultsForAgents and masters
func TestSetVMSSDefaultsAndZones(t *testing.T) {
	// masters with VMSS and no zones
	mockCS := getMockBaseContainerService("1.12.0")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.MasterProfile.HasAvailabilityZones() {
		t.Fatalf("MasterProfile.HasAvailabilityZones did not have the expected return, got %t, expected %t",
			properties.MasterProfile.HasAvailabilityZones(), false)
	}
	if properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != DefaultLoadBalancerSku {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.LoadBalancerSku did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku, DefaultLoadBalancerSku)
	}
	// masters with VMSS and zones
	mockCS = getMockBaseContainerService("1.12.0")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
	properties.MasterProfile.AvailabilityZones = []string{"1", "2"}
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	singlePlacementGroup := DefaultSinglePlacementGroup
	if *properties.MasterProfile.SinglePlacementGroup != singlePlacementGroup {
		t.Fatalf("MasterProfile.SinglePlacementGroup default did not have the expected configuration, got %t, expected %t",
			*properties.MasterProfile.SinglePlacementGroup, singlePlacementGroup)
	}
	if !properties.MasterProfile.HasAvailabilityZones() {
		t.Fatalf("MasterProfile.HasAvailabilityZones did not have the expected return, got %t, expected %t",
			properties.MasterProfile.HasAvailabilityZones(), true)
	}
	if properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != StandardLoadBalancerSku {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.LoadBalancerSku did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku, StandardLoadBalancerSku)
	}
	excludeMaster := DefaultExcludeMasterFromStandardLB
	if *properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB != excludeMaster {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB did not have the expected configuration, got %t, expected %t",
			*properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB, excludeMaster)
	}
	// agents with VMSS and no zones
	mockCS = getMockBaseContainerService("1.12.0")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.AgentPoolProfiles[0].Count = 4
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if properties.AgentPoolProfiles[0].HasAvailabilityZones() {
		t.Fatalf("AgentPoolProfiles[0].HasAvailabilityZones did not have the expected return, got %t, expected %t",
			properties.AgentPoolProfiles[0].HasAvailabilityZones(), false)
	}
	if properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != DefaultLoadBalancerSku {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.LoadBalancerSku did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku, DefaultLoadBalancerSku)
	}
	// agents with VMSS and zones
	mockCS = getMockBaseContainerService("1.12.0")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.AgentPoolProfiles[0].Count = 4
	properties.AgentPoolProfiles[0].AvailabilityZones = []string{"1", "2"}
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if !properties.AgentPoolProfiles[0].IsVirtualMachineScaleSets() {
		t.Fatalf("AgentPoolProfile[0].AvailabilityProfile did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].AvailabilityProfile, VirtualMachineScaleSets)
	}
	if !properties.AgentPoolProfiles[0].HasAvailabilityZones() {
		t.Fatalf("AgentPoolProfiles[0].HasAvailabilityZones did not have the expected return, got %t, expected %t",
			properties.AgentPoolProfiles[0].HasAvailabilityZones(), true)
	}
	singlePlacementGroup = DefaultSinglePlacementGroup
	if *properties.AgentPoolProfiles[0].SinglePlacementGroup != singlePlacementGroup {
		t.Fatalf("AgentPoolProfile[0].SinglePlacementGroup default did not have the expected configuration, got %t, expected %t",
			*properties.AgentPoolProfiles[0].SinglePlacementGroup, singlePlacementGroup)
	}
	if properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != StandardLoadBalancerSku {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.LoadBalancerSku did not have the expected configuration, got %s, expected %s",
			properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku, StandardLoadBalancerSku)
	}
	excludeMaster = DefaultExcludeMasterFromStandardLB
	if *properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB != excludeMaster {
		t.Fatalf("OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB did not have the expected configuration, got %t, expected %t",
			*properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB, excludeMaster)
	}

	properties.AgentPoolProfiles[0].Count = 110
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if to.Bool(properties.AgentPoolProfiles[0].SinglePlacementGroup) {
		t.Fatalf("AgentPoolProfile[0].SinglePlacementGroup did not have the expected configuration, got %t, expected %t",
			*properties.AgentPoolProfiles[0].SinglePlacementGroup, false)
	}

	if !*properties.AgentPoolProfiles[0].SinglePlacementGroup && properties.AgentPoolProfiles[0].StorageProfile != ManagedDisks {
		t.Fatalf("AgentPoolProfile[0].StorageProfile did not have the expected configuration, got %s, expected %s",
			properties.AgentPoolProfiles[0].StorageProfile, ManagedDisks)
	}

}

func TestAzureCNIVersionString(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10.3")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	mockCS.setOrchestratorDefaults(true, true)

	if properties.OrchestratorProfile.KubernetesConfig.AzureCNIVersion != AzureCniPluginVerLinux {
		t.Fatalf("Azure CNI Version string not the expected value, got %s, expected %s", properties.OrchestratorProfile.KubernetesConfig.AzureCNIVersion, AzureCniPluginVerLinux)
	}

	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	properties.AgentPoolProfiles[0].OSType = Windows
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
	mockCS.setOrchestratorDefaults(true, true)

	if properties.OrchestratorProfile.KubernetesConfig.AzureCNIVersion != AzureCniPluginVerWindows {
		t.Fatalf("Azure CNI Version string not the expected value, got %s, expected %s", properties.OrchestratorProfile.KubernetesConfig.AzureCNIVersion, AzureCniPluginVerWindows)
	}

	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
	mockCS.setOrchestratorDefaults(true, true)

	if properties.OrchestratorProfile.KubernetesConfig.AzureCNIVersion != "" {
		t.Fatalf("Azure CNI Version string not the expected value, got %s, expected %s", properties.OrchestratorProfile.KubernetesConfig.AzureCNIVersion, "")
	}
}

func TestEnableAggregatedAPIs(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10.3")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(false)
	mockCS.setOrchestratorDefaults(true, true)

	if properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs {
		t.Fatalf("got unexpected EnableAggregatedAPIs config value for EnableRbac=false: %t",
			properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs)
	}

	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.EnableRbac = to.BoolPtr(true)
	mockCS.setOrchestratorDefaults(true, true)

	if !properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs {
		t.Fatalf("got unexpected EnableAggregatedAPIs config value for EnableRbac=true: %t",
			properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs)
	}
}

func TestDefaultCloudProvider(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10.3")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.setOrchestratorDefaults(true, true)

	if !to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoff) {
		t.Fatalf("got unexpected CloudProviderBackoff expected true, got %t",
			to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoff))
	}

	if !to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimit) {
		t.Fatalf("got unexpected CloudProviderBackoff expected true, got %t",
			to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoff))
	}

	mockCS = getMockBaseContainerService("1.10.3")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoff = to.BoolPtr(false)
	properties.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimit = to.BoolPtr(false)
	mockCS.setOrchestratorDefaults(true, true)

	if to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoff) {
		t.Fatalf("got unexpected CloudProviderBackoff expected true, got %t",
			to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoff))
	}

	if to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimit) {
		t.Fatalf("got unexpected CloudProviderBackoff expected true, got %t",
			to.Bool(properties.OrchestratorProfile.KubernetesConfig.CloudProviderBackoff))
	}
}
func TestSetCertDefaults(t *testing.T) {
	cs := &ContainerService{
		Properties: &Properties{
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &MasterProfile{
				Count:     3,
				DNSPrefix: "myprefix1",
				VMSize:    "Standard_DS2_v2",
			},
			OrchestratorProfile: &OrchestratorProfile{
				OrchestratorType:    Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &KubernetesConfig{
					NetworkPlugin: NetworkPluginAzure,
				},
			},
		},
	}

	cs.setOrchestratorDefaults(false, false)
	cs.Properties.setMasterProfileDefaults(false, false, AzurePublicCloud)
	result, ips, err := cs.SetDefaultCerts(DefaultCertParams{
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if !result {
		t.Error("expected SetDefaultCerts to return true")
	}

	if err != nil {
		t.Errorf("unexpected error thrown while executing SetDefaultCerts %s", err.Error())
	}

	if ips == nil {
		t.Error("expected SetDefaultCerts to create a list of IPs")
	} else {

		if len(ips) != cs.Properties.MasterProfile.Count+3 {
			t.Errorf("expected length of IPs from SetDefaultCerts %d, actual length %d", cs.Properties.MasterProfile.Count+3, len(ips))
		}

		firstMasterIP := net.ParseIP(cs.Properties.MasterProfile.FirstConsecutiveStaticIP).To4()
		offsetMultiplier := 1
		addr := binary.BigEndian.Uint32(firstMasterIP)
		expectedNewAddr := getNewAddr(addr, cs.Properties.MasterProfile.Count-1, offsetMultiplier)
		actualLastIPAddr := binary.BigEndian.Uint32(ips[len(ips)-2])
		if actualLastIPAddr != expectedNewAddr {
			expectedLastIP := make(net.IP, 4)
			binary.BigEndian.PutUint32(expectedLastIP, expectedNewAddr)
			t.Errorf("expected last IP of master vm from SetDefaultCerts %d, actual %d", expectedLastIP, ips[len(ips)-2])
		}

		if cs.Properties.MasterProfile.HasMultipleNodes() {
			expectedILBIP := net.IP{firstMasterIP[0], firstMasterIP[1], firstMasterIP[2], firstMasterIP[3] + byte(DefaultInternalLbStaticIPOffset)}
			actualILBIPAddr := binary.BigEndian.Uint32(ips[2])
			expectedILBIPAddr := binary.BigEndian.Uint32(expectedILBIP)

			if actualILBIPAddr != expectedILBIPAddr {
				t.Errorf("expected IP of master ILB from SetDefaultCerts %d, actual %d", expectedILBIP, ips[2])
			}
		}
	}
}

func TestSetCertDefaultsVMSS(t *testing.T) {
	cs := &ContainerService{
		Properties: &Properties{
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: "barClientID",
				Secret:   "bazSecret",
			},
			MasterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "myprefix1",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			OrchestratorProfile: &OrchestratorProfile{
				OrchestratorType:    Kubernetes,
				OrchestratorVersion: "1.10.2",
				KubernetesConfig: &KubernetesConfig{
					NetworkPlugin: NetworkPluginAzure,
				},
			},
		},
	}

	cs.setOrchestratorDefaults(false, false)
	cs.Properties.setMasterProfileDefaults(false, false, AzurePublicCloud)
	result, ips, err := cs.SetDefaultCerts(DefaultCertParams{
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if !result {
		t.Error("expected SetDefaultCerts to return true")
	}

	if err != nil {
		t.Errorf("unexpected error thrown while executing SetDefaultCerts %s", err.Error())
	}

	if ips == nil {
		t.Error("expected SetDefaultCerts to create a list of IPs")
	} else {

		if len(ips) != cs.Properties.MasterProfile.Count+3 {
			t.Errorf("expected length of IPs from SetDefaultCerts %d, actual length %d", cs.Properties.MasterProfile.Count+3, len(ips))
		}

		firstMasterIP := net.ParseIP(cs.Properties.MasterProfile.FirstConsecutiveStaticIP).To4()
		offsetMultiplier := cs.Properties.MasterProfile.IPAddressCount
		addr := binary.BigEndian.Uint32(firstMasterIP)
		expectedNewAddr := getNewAddr(addr, cs.Properties.MasterProfile.Count-1, offsetMultiplier)
		actualLastIPAddr := binary.BigEndian.Uint32(ips[len(ips)-2])
		if actualLastIPAddr != expectedNewAddr {
			expectedLastIP := make(net.IP, 4)
			binary.BigEndian.PutUint32(expectedLastIP, expectedNewAddr)
			t.Errorf("expected last IP of master vm from SetDefaultCerts %d, actual %d", expectedLastIP, ips[len(ips)-2])
		}

		if cs.Properties.MasterProfile.HasMultipleNodes() {
			expectedILBIP := net.IP{firstMasterIP[0], firstMasterIP[1], byte(255), byte(DefaultInternalLbStaticIPOffset)}
			actualILBIPAddr := binary.BigEndian.Uint32(ips[2])
			expectedILBIPAddr := binary.BigEndian.Uint32(expectedILBIP)

			if actualILBIPAddr != expectedILBIPAddr {
				t.Errorf("expected IP of master ILB from SetDefaultCerts %d, actual %d", expectedILBIP, ips[2])
			}
		}
	}
}

func TestProxyModeDefaults(t *testing.T) {
	// Test that default is what we expect
	mockCS := getMockBaseContainerService("1.10.12")
	properties := mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.MasterProfile.Count = 1
	mockCS.setOrchestratorDefaults(true, true)

	if properties.OrchestratorProfile.KubernetesConfig.ProxyMode != DefaultKubeProxyMode {
		t.Fatalf("ProxyMode string not the expected default value, got %s, expected %s", properties.OrchestratorProfile.KubernetesConfig.ProxyMode, DefaultKubeProxyMode)
	}

	// Test that default assignment flow doesn't overwrite a user-provided config
	mockCS = getMockBaseContainerService("1.10.12")
	properties = mockCS.Properties
	properties.OrchestratorProfile.OrchestratorType = Kubernetes
	properties.OrchestratorProfile.KubernetesConfig.ProxyMode = KubeProxyModeIPVS
	properties.MasterProfile.Count = 1
	mockCS.setOrchestratorDefaults(true, true)

	if properties.OrchestratorProfile.KubernetesConfig.ProxyMode != KubeProxyModeIPVS {
		t.Fatalf("ProxyMode string not the expected default value, got %s, expected %s", properties.OrchestratorProfile.KubernetesConfig.ProxyMode, KubeProxyModeIPVS)
	}
}
func TestSetCustomCloudProfileDefaults(t *testing.T) {

	// Test that the ResourceManagerVMDNSSuffix is set in EndpointConfig
	mockCS := getMockBaseContainerService("1.11.6")
	mockCSP := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	vmDNSSuffix := "contoso.net"
	mockCSP.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix = vmDNSSuffix
	mockCS.Properties.CustomCloudProfile = mockCSP.CustomCloudProfile
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if AzureCloudSpecEnvMap[AzureStackCloud].EndpointConfig.ResourceManagerVMDNSSuffix != vmDNSSuffix {
		t.Errorf("setCustomCloudProfileDefaults(): ResourceManagerVMDNSSuffix string in AzureCloudSpecEnvMap[AzureStackCloud] not the expected default value, got %s, expected %s", AzureCloudSpecEnvMap[AzureStackCloud].EndpointConfig.ResourceManagerVMDNSSuffix, vmDNSSuffix)
	}

	// Test that the AzureStackCloudSpec is default when azureEnvironmentSpecConfig is empty in api model JSON file
	mockCSDefaultSpec := getMockBaseContainerService("1.11.6")
	mockCSPDefaultSpec := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSDefaultSpec.Properties.CustomCloudProfile = mockCSPDefaultSpec.CustomCloudProfile
	mockCSDefaultSpec.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	actualEnv := AzureCloudSpecEnvMap[AzureStackCloud]
	expectedEnv := AzureCloudSpecEnvMap[AzurePublicCloud]
	expectedEnv.EndpointConfig.ResourceManagerVMDNSSuffix = mockCSPDefaultSpec.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix
	expectedEnv.CloudName = AzureStackCloud
	expectedEnv.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackDeployTelemetryPID
	if diff := cmp.Diff(actualEnv, expectedEnv); diff != "" {
		t.Errorf("setCustomCloudProfileDefaults(): did not set AzureStackCloudSpec as default when azureEnvironmentSpecConfig is empty in api model JSON file. %s", diff)
	}

	modeToSpec := map[string]string{
		"public":       "AzurePublicCloud",
		"china":        "AzureChinaCloud",
		"german":       "AzureGermanCloud",
		"usgovernment": "AzureUSGovernmentCloud",
	}

	for key, value := range modeToSpec {
		mockCSAzureChinaSpec := getMockBaseContainerService("1.11.6")
		mockCSPAzureChinaSpec := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
		mockCSPAzureChinaSpec.CustomCloudProfile.DependenciesLocation = DependenciesLocation(key)
		mockCSAzureChinaSpec.Properties.CustomCloudProfile = mockCSPAzureChinaSpec.CustomCloudProfile

		mockCSAzureChinaSpec.SetPropertiesDefaults(PropertiesDefaultsParams{
			IsScale:    false,
			IsUpgrade:  false,
			PkiKeySize: helpers.DefaultPkiKeySize,
		})

		actualEnvAzureChinaSpec := AzureCloudSpecEnvMap[AzureStackCloud]
		expectedEnvAzureChinaSpec := AzureCloudSpecEnvMap[value]
		expectedEnvAzureChinaSpec.EndpointConfig.ResourceManagerVMDNSSuffix = mockCSPDefaultSpec.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix
		expectedEnvAzureChinaSpec.CloudName = AzureStackCloud
		expectedEnvAzureChinaSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackDeployTelemetryPID
		t.Logf("verifying dependenciesLocation: %s", key)
		if diff := cmp.Diff(actualEnvAzureChinaSpec, expectedEnvAzureChinaSpec); diff != "" {
			t.Errorf("setCustomCloudProfileDefaults(): did not set AzureStackCloudSpec as default when connection Mode is %s in api model JSON file. %s", key, diff)
		}
	}

	// Test that correct error message if ResourceManagerVMDNSSuffix is empty
	mockCSEmptyResourceManagerVMDNSSuffix := getMockBaseContainerService("1.11.6")
	mockCSPEmptyResourceManagerVMDNSSuffix := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSEmptyResourceManagerVMDNSSuffix.Properties.CustomCloudProfile = mockCSPEmptyResourceManagerVMDNSSuffix.CustomCloudProfile
	mockCSEmptyResourceManagerVMDNSSuffix.Properties.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix = ""

	acutalerr := mockCSEmptyResourceManagerVMDNSSuffix.Properties.SetAzureStackCloudSpec(AzureStackCloudSpecParams{
		IsUpgrade: false,
		IsScale:   false,
	})
	expectError := errors.New("Failed to set Cloud Spec for Azure Stack due to invalid environment")
	if !helpers.EqualError(acutalerr, expectError) {
		t.Errorf("verify ResourceManagerVMDNSSuffix empty: expected error: %s - got: %s", acutalerr, expectError)
	}

	// Test that correct error message if environment is nil
	mockCSNilEnvironment := getMockBaseContainerService("1.11.6")
	mockCSPNilEnvironment := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSNilEnvironment.Properties.CustomCloudProfile = mockCSPNilEnvironment.CustomCloudProfile
	mockCSNilEnvironment.Properties.CustomCloudProfile.Environment = nil
	acutalerr = mockCSEmptyResourceManagerVMDNSSuffix.Properties.SetAzureStackCloudSpec(AzureStackCloudSpecParams{
		IsUpgrade: false,
		IsScale:   false,
	})
	if !helpers.EqualError(acutalerr, expectError) {
		t.Errorf("verify environment nil: expected error: %s - got: %s", acutalerr, expectError)
	}

	// Test that default assignment flow doesn't overwrite a user-provided config
	mockCSCustom := getMockBaseContainerService("1.11.6")
	mockCSPCustom := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)

	//Mock AzureEnvironmentSpecConfig
	customCloudSpec := AzureEnvironmentSpecConfig{
		CloudName: "AzuReStackCloud",
		//DockerSpecConfig specify the docker engine download repo
		DockerSpecConfig: DockerSpecConfig{
			DockerEngineRepo:         "DockerEngineRepo",
			DockerComposeDownloadURL: "DockerComposeDownloadURL",
		},
		//KubernetesSpecConfig - Due to Chinese firewall issue, the default containers from google is blocked, use the Chinese local mirror instead
		KubernetesSpecConfig: KubernetesSpecConfig{
			AzureTelemetryPID:                "AzureTelemetryPID",
			KubernetesImageBase:              "KubernetesImageBase",
			TillerImageBase:                  "TillerImageBase",
			ACIConnectorImageBase:            "ACIConnectorImageBase",
			NVIDIAImageBase:                  "NVIDIAImageBase",
			AzureCNIImageBase:                "AzureCNIImageBase",
			CalicoImageBase:                  "CalicoImageBase",
			EtcdDownloadURLBase:              "EtcdDownloadURLBase",
			KubeBinariesSASURLBase:           "KubeBinariesSASURLBase",
			WindowsTelemetryGUID:             "WindowsTelemetryGUID",
			CNIPluginsDownloadURL:            "CNIPluginsDownloadURL",
			VnetCNILinuxPluginsDownloadURL:   "VnetCNILinuxPluginsDownloadURL",
			VnetCNIWindowsPluginsDownloadURL: "VnetCNIWindowsPluginsDownloadURL",
			ContainerdDownloadURLBase:        "ContainerdDownloadURLBase",
		},
		DCOSSpecConfig: DefaultDCOSSpecConfig,
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
		},
		OSImageConfig: map[Distro]AzureOSImageConfig{
			Distro("Test"): {
				ImageOffer:     "ImageOffer",
				ImageSku:       "ImageSku",
				ImagePublisher: "ImagePublisher",
				ImageVersion:   "ImageVersion",
			},
			AKSUbuntu1604: AKSUbuntu1604OSImageConfig,
		},
	}
	mockCSPCustom.CustomCloudProfile.AzureEnvironmentSpecConfig = &customCloudSpec
	mockCSCustom.Properties.CustomCloudProfile = mockCSPCustom.CustomCloudProfile
	mockCSCustom.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if diff := cmp.Diff(AzureCloudSpecEnvMap[AzureStackCloud], customCloudSpec); diff != "" {
		t.Errorf("setCustomCloudProfileDefaults(): did not set AzureStackCloudSpec as default when azureEnvironmentSpecConfig is empty in api model JSON file")
	}

	if diff := cmp.Diff(mockCSCustom.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig, &customCloudSpec); diff != "" {
		t.Errorf("setCustomCloudProfileDefaults(): did not set CustomCloudProfile.AzureEnvironmentSpecConfig with customer input")
	}

	// Test that default assignment flow set the value if the field is partially  missing in user-provided config
	mockCSCustomP := getMockBaseContainerService("1.11.6")
	mockCSPCustomP := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)

	//Mock AzureEnvironmentSpecConfig
	customCloudSpecP := AzureEnvironmentSpecConfig{
		CloudName: "AzureStackCloud",
		//DockerSpecConfig specify the docker engine download repo
		DockerSpecConfig: DockerSpecConfig{
			DockerEngineRepo: "DockerEngineRepo",
		},
		//KubernetesSpecConfig - Due to Chinese firewall issue, the default containers from google is blocked, use the Chinese local mirror instead
		KubernetesSpecConfig: KubernetesSpecConfig{
			KubernetesImageBase:            "KubernetesImageBase",
			TillerImageBase:                "TillerImageBase",
			NVIDIAImageBase:                "NVIDIAImageBase",
			AzureCNIImageBase:              "AzureCNIImageBase",
			CalicoImageBase:                "CalicoImageBase",
			EtcdDownloadURLBase:            "EtcdDownloadURLBase",
			WindowsTelemetryGUID:           "WindowsTelemetryGUID",
			CNIPluginsDownloadURL:          "CNIPluginsDownloadURL",
			VnetCNILinuxPluginsDownloadURL: "VnetCNILinuxPluginsDownloadURL",
			ContainerdDownloadURLBase:      "ContainerdDownloadURLBase",
		},
		DCOSSpecConfig: DefaultDCOSSpecConfig,
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
		},
		OSImageConfig: map[Distro]AzureOSImageConfig{
			Distro("Test"): {
				ImageOffer:     "ImageOffer",
				ImageSku:       "ImageSku",
				ImagePublisher: "ImagePublisher",
				ImageVersion:   "ImageVersion",
			},
			AKSUbuntu1604: AKSUbuntu1604OSImageConfig,
		},
	}
	mockCSPCustomP.CustomCloudProfile.AzureEnvironmentSpecConfig = &customCloudSpecP
	mockCSCustomP.Properties.CustomCloudProfile = mockCSPCustomP.CustomCloudProfile
	mockCSCustomP.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.DockerSpecConfig.DockerComposeDownloadURL != DefaultDockerSpecConfig.DockerComposeDownloadURL {
		t.Errorf("setCustomCloudProfileDefaults(): did not set DockerComposeDownloadURL with default Value, got '%s', expected %s", mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.DockerSpecConfig.DockerComposeDownloadURL, DefaultDockerSpecConfig.DockerComposeDownloadURL)
	}
	if mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.KubernetesSpecConfig.ACIConnectorImageBase != DefaultKubernetesSpecConfig.ACIConnectorImageBase {
		t.Errorf("setCustomCloudProfileDefaults(): did not set ACIConnectorImageBase with default Value, got '%s', expected %s", mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.KubernetesSpecConfig.ACIConnectorImageBase, DefaultKubernetesSpecConfig.ACIConnectorImageBase)
	}
	if mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.KubernetesSpecConfig.KubeBinariesSASURLBase != DefaultKubernetesSpecConfig.KubeBinariesSASURLBase {
		t.Errorf("setCustomCloudProfileDefaults(): did not set KubeBinariesSASURLBase with default Value, got '%s', expected %s", mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.KubernetesSpecConfig.KubeBinariesSASURLBase, DefaultKubernetesSpecConfig.KubeBinariesSASURLBase)
	}
	if mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL != DefaultKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL {
		t.Errorf("setCustomCloudProfileDefaults(): did not set VnetCNIWindowsPluginsDownloadURL with default Value, got '%s', expected %s", mockCSCustomP.Properties.CustomCloudProfile.AzureEnvironmentSpecConfig.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL, DefaultKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL)
	}
	// Test that the default values are set for IdentitySystem and AuthenticationMethod if they are not in the configuration
	mockCSAuth := getMockBaseContainerService("1.11.6")
	mockCSPAuth := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true)
	mockCSPAuth.CustomCloudProfile.IdentitySystem = ""
	mockCSPAuth.CustomCloudProfile.AuthenticationMethod = ""
	mockCSAuth.Properties.CustomCloudProfile = mockCSPAuth.CustomCloudProfile
	mockCSAuth.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if mockCSAuth.Properties.CustomCloudProfile.AuthenticationMethod != ClientSecretAuthMethod {
		t.Errorf("setCustomCloudProfileDefaults(): AuthenticationMethod string not the expected default value, got %s, expected %s", mockCSAuth.Properties.CustomCloudProfile.AuthenticationMethod, ClientSecretAuthMethod)
	}
	if mockCSAuth.Properties.CustomCloudProfile.IdentitySystem != AzureADIdentitySystem {
		t.Errorf("setCustomCloudProfileDefaults(): IdentitySystem string not the expected default value, got %s, expected %s", mockCSAuth.Properties.CustomCloudProfile.IdentitySystem, AzureADIdentitySystem)
	}

	// Test that the custom input values are not overiwrited if they are in the configuration
	mockCSI := getMockBaseContainerService("1.11.6")
	mockCSPI := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, true)
	mockCSPI.CustomCloudProfile.IdentitySystem = ADFSIdentitySystem
	mockCSPI.CustomCloudProfile.AuthenticationMethod = ClientCertificateAuthMethod
	mockCSI.Properties.CustomCloudProfile = mockCSPI.CustomCloudProfile
	mockCSI.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	if mockCSI.Properties.CustomCloudProfile.AuthenticationMethod != ClientCertificateAuthMethod {
		t.Errorf("setCustomCloudProfileDefaults(): AuthenticationMethod string from customer not the expected default value, got %s, expected %s", mockCSI.Properties.CustomCloudProfile.AuthenticationMethod, ClientCertificateAuthMethod)
	}
	if mockCSI.Properties.CustomCloudProfile.IdentitySystem != ADFSIdentitySystem {
		t.Errorf("setCustomCloudProfileDefaults(): IdentitySystem string from customer not the expected default value, got %s, expected %s", mockCSI.Properties.CustomCloudProfile.IdentitySystem, ADFSIdentitySystem)
	}
}

func TestCustomCloudLocation(t *testing.T) {

	// Test that the ResourceManagerVMDNSSuffix is set in EndpointConfig
	mockCS := getMockBaseContainerService("1.11.6")
	mockCSP := GetMockPropertiesWithCustomCloudProfile("AzureStackCloud", true, true, true)
	mockCS.Properties = &mockCSP
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	dnsPrefix := "santest"
	actual := []string{FormatProdFQDNByLocation(dnsPrefix, mockCS.Location, "AzureStackCloud")}

	expected := []string{fmt.Sprintf("%s.%s.%s", dnsPrefix, mockCS.Location, AzureCloudSpecEnvMap[AzureStackCloud].EndpointConfig.ResourceManagerVMDNSSuffix)}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected formatted fqdns %s, but got %s", expected, actual)
	}
}

func TestSetCustomCloudProfileEnvironmentDefaults(t *testing.T) {
	location := "testlocation"
	cs := ContainerService{
		Location: location,
		Properties: &Properties{
			CustomCloudProfile: &CustomCloudProfile{
				IdentitySystem: "adfs",
				PortalURL:      "https://portal.testlocation.contoso.com/",
			},
		},
	}

	csPortal := ContainerService{
		Location: location,
		Properties: &Properties{
			CustomCloudProfile: &CustomCloudProfile{
				IdentitySystem: "adfs",
				PortalURL:      "https://portal.testlocation.contoso.com",
			},
		},
	}

	expectedEnv := &azure.Environment{
		Name:                       "AzureStackCloud",
		ManagementPortalURL:        "https://portal.testlocation.contoso.com/",
		ServiceManagementEndpoint:  "https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7",
		ResourceManagerEndpoint:    fmt.Sprintf("https://management.%s.contoso.com/", location),
		ActiveDirectoryEndpoint:    "https://adfs.testlocation.contoso.com/",
		GalleryEndpoint:            "https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/",
		GraphEndpoint:              "https://graph.testlocation.contoso.com/",
		StorageEndpointSuffix:      "testlocation.contoso.com",
		KeyVaultDNSSuffix:          "vault.testlocation.contoso.com",
		ResourceManagerVMDNSSuffix: "cloudapp.contoso.com",
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/adfs","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	err := cs.SetCustomCloudProfileEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(cs.Properties.CustomCloudProfile.Environment, expectedEnv); diff != "" {
		t.Errorf("Fail to compare, Environment adfs %q", diff)
	}

	err = csPortal.SetCustomCloudProfileEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(csPortal.Properties.CustomCloudProfile.Environment, expectedEnv); diff != "" {
		t.Errorf("Fail to compare, Environment portal url adfs %q", diff)
	}

	csAzureAD := ContainerService{
		Location: location,
		Properties: &Properties{
			CustomCloudProfile: &CustomCloudProfile{
				IdentitySystem: "azure_ad",
				PortalURL:      "https://portal.testlocation.contoso.com/",
			},
		},
	}

	//test setCustomCloudProfileDefaults with portal url
	mockCS := getMockBaseContainerService("1.11.6")
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}

	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)
	mockCS.Location = location
	_, err = mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if err != nil {
		t.Errorf("Failed to test setCustomCloudProfileDefaults with portal url - %s", err)
	}
	if diff := cmp.Diff(mockCS.Properties.CustomCloudProfile.Environment, expectedEnv); diff != "" {
		t.Errorf("Fail to compare, Environment setCustomCloudProfileDefaults %q", diff)
	}

	cloudSpec := AzureCloudSpecEnvMap[AzurePublicCloud]
	cloudSpec.CloudName = AzureStackCloud
	cloudSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackDeployTelemetryPID
	cloudSpec.EndpointConfig.ResourceManagerVMDNSSuffix = mockCS.Properties.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix
	if diff := cmp.Diff(AzureCloudSpecEnvMap[AzureStackCloud], cloudSpec); diff != "" {
		t.Errorf("Fail to compare, AzureCloudSpec AzureStackCloud %q", diff)
	}

	// Test for azure_ad
	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	err = csAzureAD.SetCustomCloudProfileEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(csAzureAD.Properties.CustomCloudProfile.Environment, expectedEnv); diff != "" {
		t.Errorf("Fail to compare, Environment azure_ad %q", diff)
	}

	csError := ContainerService{
		Location: location,
		Properties: &Properties{
			CustomCloudProfile: &CustomCloudProfile{
				IdentitySystem: "azure_ad",
				PortalURL:      "https://portal.abc.contoso.com/",
			},
		},
	}

	err = csError.SetCustomCloudProfileEnvironment()
	expectedError := fmt.Errorf("portalURL needs to start with https://portal.%s. ", location)
	if !helpers.EqualError(err, expectedError) {
		t.Errorf("expected error %s, got %s", expectedError, err)
	}
}

func TestSetOrchestratorProfileDefaultsOnAzureStack(t *testing.T) {
	location := "testlocation"
	//Test setMasterProfileDefaults with portal url
	mockCS := getMockBaseContainerService("1.11.6")
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}
	mockCS.Location = location
	mockCS.Properties.OrchestratorProfile.OrchestratorType = "Kubernetes"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/adfs","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if (*mockCS.Properties.OrchestratorProfile.KubernetesConfig.UseInstanceMetadata) != DefaultAzureStackUseInstanceMetadata {
		t.Fatalf("DefaultAzureStackUseInstanceMetadata did not have the expected value, got %t, expected %t",
			(*mockCS.Properties.OrchestratorProfile.KubernetesConfig.UseInstanceMetadata), DefaultAzureStackUseInstanceMetadata)
	}
}

func TestSetMasterProfileDefaultsOnAzureStack(t *testing.T) {
	location := "testlocation"
	oldFaultDomainCount := 2
	//Test setMasterProfileDefaults with portal url
	mockCS := getMockBaseContainerService("1.11.6")
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}
	mockCS.Location = location
	mockCS.Properties.MasterProfile.AvailabilityProfile = ""
	mockCS.Properties.MasterProfile.Count = 1

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/adfs","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if (*mockCS.Properties.MasterProfile.PlatformFaultDomainCount) != DefaultAzureStackFaultDomainCount {
		t.Fatalf("PlatformFaultDomainCount did not have the expected value, got %d, expected %d",
			(*mockCS.Properties.MasterProfile.PlatformFaultDomainCount), DefaultAzureStackFaultDomainCount)
	}

	// Check scenario where value is already set.
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}
	mockCS.Properties.MasterProfile.AvailabilityProfile = ""
	mockCS.Properties.MasterProfile.Count = 1
	mockCS.Properties.MasterProfile.PlatformFaultDomainCount = &oldFaultDomainCount
	mockCS.Location = location
	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if (*mockCS.Properties.MasterProfile.PlatformFaultDomainCount) != oldFaultDomainCount {
		t.Fatalf("PlatformFaultDomainCount did not have the expected value, got %d, expected %d",
			(*mockCS.Properties.MasterProfile.PlatformFaultDomainCount), oldFaultDomainCount)
	}
}

func TestSetAgentProfileDefaultsOnAzureStack(t *testing.T) {
	location := "testlocation"
	oldFaultDomainCount := 2
	//Test setMasterProfileDefaults with portal url
	mockCS := getMockBaseContainerService("1.11.6")
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}
	mockCS.Location = location
	mockCS.Properties.MasterProfile.AvailabilityProfile = ""
	mockCS.Properties.MasterProfile.Count = 1

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/adfs","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	for _, pool := range mockCS.Properties.AgentPoolProfiles {
		if (*pool.PlatformFaultDomainCount) != DefaultAzureStackFaultDomainCount {
			t.Fatalf("PlatformFaultDomainCount did not have the expected value, got %d, expected %d",
				(*pool.PlatformFaultDomainCount), DefaultAzureStackFaultDomainCount)
		}

		if (*pool.AcceleratedNetworkingEnabled) != DefaultAzureStackAcceleratedNetworking {
			t.Fatalf("AcceleratedNetworkingEnabled did not have the expected value, got %t, expected %t",
				(*pool.AcceleratedNetworkingEnabled), DefaultAzureStackAcceleratedNetworking)
		}

		if (*pool.AcceleratedNetworkingEnabledWindows) != DefaultAzureStackAcceleratedNetworking {
			t.Fatalf("AcceleratedNetworkingEnabledWindows did not have the expected value, got %t, expected %t",
				(*pool.AcceleratedNetworkingEnabledWindows), DefaultAzureStackAcceleratedNetworking)
		}
	}
	// Check scenario where value is already set.
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}
	mockCS.Properties.MasterProfile.AvailabilityProfile = ""
	mockCS.Properties.MasterProfile.Count = 1
	for _, pool := range mockCS.Properties.AgentPoolProfiles {
		pool.PlatformFaultDomainCount = &oldFaultDomainCount
	}
	mockCS.Location = location

	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	for _, pool := range mockCS.Properties.AgentPoolProfiles {
		if (*pool.PlatformFaultDomainCount) != oldFaultDomainCount {
			t.Fatalf("PlatformFaultDomainCount did not have the expected value, got %d, expected %d",
				(*pool.PlatformFaultDomainCount), oldFaultDomainCount)
		}
	}
}

func TestEtcdDiskSizeOnAzureStack(t *testing.T) {
	location := "testlocation"
	mockCS := getMockBaseContainerService("1.11.6")
	mockCS.Location = location
	mockCS.Properties.MasterProfile.Count = 1
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/adfs","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != DefaultEtcdDiskSize {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, DefaultEtcdDiskSize)
	}

	// Case where total node count is 5.
	mockCS = getMockBaseContainerService("1.11.6")
	mockCS.Location = location
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.MasterProfile.Count = 5
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}

	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != DefaultEtcdDiskSizeGT3Nodes {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, DefaultEtcdDiskSizeGT3Nodes)
	}

	// Case where total node count is 11.
	mockCS = getMockBaseContainerService("1.11.6")
	mockCS.Location = location
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.MasterProfile.Count = 5
	mockCS.Properties.AgentPoolProfiles[0].Count = 6
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}

	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != MaxAzureStackManagedDiskSize {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, MaxAzureStackManagedDiskSize)
	}

	// Case where total node count is 21.
	mockCS = getMockBaseContainerService("1.11.6")
	mockCS.Location = location
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.MasterProfile.Count = 5
	mockCS.Properties.AgentPoolProfiles[0].Count = 16
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}

	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != MaxAzureStackManagedDiskSize {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, MaxAzureStackManagedDiskSize)
	}

	// Case where total node count is 55 but EtcdDiskSizeGB size is passed
	mockCS = getMockBaseContainerService("1.11.6")
	mockCS.Location = location
	mockCS.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	mockCS.Properties.MasterProfile.Count = 5
	mockCS.Properties.AgentPoolProfiles[0].Count = 50
	customEtcdDiskSize := "512"
	mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB = customEtcdDiskSize
	mockCS.Properties.CustomCloudProfile = &CustomCloudProfile{
		PortalURL: "https://portal.testlocation.contoso.com",
	}

	httpmock.DeactivateAndReset()
	httpmock.Activate()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%smetadata/endpoints?api-version=1.0", fmt.Sprintf("https://management.%s.contoso.com/", location)),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"galleryEndpoint":"https://galleryartifacts.hosting.testlocation.contoso.com/galleryartifacts/","graphEndpoint":"https://graph.testlocation.contoso.com/","portalEndpoint":"https://portal.testlocation.contoso.com/","authentication":{"loginEndpoint":"https://adfs.testlocation.contoso.com/","audiences":["https://management.adfs.azurestack.testlocation/ce080287-be51-42e5-b99e-9de760fecae7"]}}`)
			return resp, nil
		},
	)

	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB != customEtcdDiskSize {
		t.Fatalf("EtcdDiskSizeGB did not have the expected size, got %s, expected %s",
			mockCS.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB, customEtcdDiskSize)
	}
}
func TestPreserveNodesProperties(t *testing.T) {
	mockCS := getMockBaseContainerService("1.10.8")
	mockCS.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	if !to.Bool(mockCS.Properties.AgentPoolProfiles[0].PreserveNodesProperties) {
		t.Errorf("expected preserveNodesProperties to be %t instead got %t", true, to.Bool(mockCS.Properties.AgentPoolProfiles[0].PreserveNodesProperties))
	}
}

func TestUbuntu1804Flags(t *testing.T) {
	// Validate --resolv-conf is missing with 16.04 distro and present with 18.04
	cs := CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	cs.Properties.MasterProfile.Distro = AKSUbuntu1604
	cs.Properties.AgentPoolProfiles[0].Distro = AKSUbuntu1804
	cs.Properties.AgentPoolProfiles[0].OSType = Linux
	cs.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
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
	cs.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
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
	cs.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
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

func getMockBaseContainerService(orchestratorVersion string) ContainerService {
	mockAPIProperties := getMockAPIProperties(orchestratorVersion)
	return ContainerService{
		Properties: &mockAPIProperties,
	}
}

func getMockAPIProperties(orchestratorVersion string) Properties {
	return Properties{
		ProvisioningState: "",
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorVersion: orchestratorVersion,
			KubernetesConfig:    &KubernetesConfig{},
		},
		MasterProfile: &MasterProfile{},
		AgentPoolProfiles: []*AgentPoolProfile{
			{},
			{},
			{},
			{},
		}}
}

func getKubernetesConfigWithFeatureGates(featureGates string) *KubernetesConfig {
	return &KubernetesConfig{
		KubeletConfig: map[string]string{"--feature-gates": featureGates},
	}
}

func TestDefaultEnablePodSecurityPolicy(t *testing.T) {
	cases := []struct {
		name     string
		cs       ContainerService
		expected bool
	}{
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.14.0",
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: false,
		},
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.0-alpha.1",
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: true,
		},
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.0-beta.1",
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: true,
		},
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.0",
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.cs.setOrchestratorDefaults(false, false)
			if to.Bool(c.cs.Properties.OrchestratorProfile.KubernetesConfig.EnablePodSecurityPolicy) != c.expected {
				t.Errorf("expected  %t, but got %t", c.expected, to.Bool(c.cs.Properties.OrchestratorProfile.KubernetesConfig.EnablePodSecurityPolicy))
			}
		})
	}
}

func TestDefaultLoadBalancerSKU(t *testing.T) {
	cases := []struct {
		name     string
		cs       ContainerService
		expected string
	}{
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.14.0",
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: BasicLoadBalancerSku,
		},
		{
			name: "basic",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.14.0",
						KubernetesConfig: &KubernetesConfig{
							LoadBalancerSku: "basic",
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: BasicLoadBalancerSku,
		},
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.14.0",
						KubernetesConfig: &KubernetesConfig{
							LoadBalancerSku: BasicLoadBalancerSku,
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: BasicLoadBalancerSku,
		},
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.14.0",
						KubernetesConfig: &KubernetesConfig{
							LoadBalancerSku: "standard",
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: StandardLoadBalancerSku,
		},
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.14.0",
						KubernetesConfig: &KubernetesConfig{
							LoadBalancerSku: StandardLoadBalancerSku,
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: StandardLoadBalancerSku,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.cs.setOrchestratorDefaults(false, false)
			if c.cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != c.expected {
				t.Errorf("expected %s, but got %s", c.expected, c.cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku)
			}
		})
	}
}

func TestEnableRBAC(t *testing.T) {
	cases := []struct {
		name      string
		cs        ContainerService
		isUpgrade bool
		isScale   bool
		expected  bool
	}{
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType: Kubernetes,
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: true,
		},
		{
			name: "1.14 disabled",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: common.GetLatestPatchVersion("1.14", common.GetAllSupportedKubernetesVersions(false, false)),
						KubernetesConfig: &KubernetesConfig{
							EnableRbac: to.BoolPtr(false),
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: false,
		},
		{
			name: "1.14 disabled upgrade",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: common.GetLatestPatchVersion("1.14", common.GetAllSupportedKubernetesVersions(false, false)),
						KubernetesConfig: &KubernetesConfig{
							EnableRbac: to.BoolPtr(false),
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			isUpgrade: true,
			expected:  false,
		},
		{
			name: "1.15",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false)),
					},
					MasterProfile: &MasterProfile{},
				},
			},
			expected: true,
		},
		{
			name: "1.15 upgrade",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false)),
					},
					MasterProfile: &MasterProfile{},
				},
			},
			isUpgrade: true,
			expected:  true,
		},
		{
			name: "1.15 upgrade false--> true override",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false)),
						KubernetesConfig: &KubernetesConfig{
							EnableRbac: to.BoolPtr(false),
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			isUpgrade: true,
			expected:  true,
		},
		{
			name: "1.16 upgrade false--> true override",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: common.GetLatestPatchVersion("1.16", common.GetAllSupportedKubernetesVersions(false, false)),
						KubernetesConfig: &KubernetesConfig{
							EnableRbac: to.BoolPtr(false),
						},
					},
					MasterProfile: &MasterProfile{},
				},
			},
			isUpgrade: true,
			expected:  true,
		},
		{
			name: "1.15 upgrade no false--> true override in AKS scenario",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false)),
						KubernetesConfig: &KubernetesConfig{
							EnableRbac: to.BoolPtr(false),
						},
					},
					HostedMasterProfile: &HostedMasterProfile{
						FQDN: "foo",
					},
				},
			},
			isUpgrade: true,
			expected:  false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.cs.setOrchestratorDefaults(c.isUpgrade, c.isScale)
			if to.Bool(c.cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac) != c.expected {
				t.Errorf("expected %t, but got %t", c.expected, to.Bool(c.cs.Properties.OrchestratorProfile.KubernetesConfig.EnableRbac))
			}
		})
	}
}

func TestDefaultTelemetry(t *testing.T) {
	// Test that the AzureTelemetryPID is set to DefaultAzureStackDeployTelemetryPID  by default
	mockCSDefaultSpec := getMockBaseContainerService("1.11.6")
	mockCSPDefaultSpec := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSDefaultSpec.Properties.CustomCloudProfile = mockCSPDefaultSpec.CustomCloudProfile
	mockCSDefaultSpec.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	actualEnv := AzureCloudSpecEnvMap[AzureStackCloud]
	expectedEnv := AzureCloudSpecEnvMap[AzurePublicCloud]
	expectedEnv.EndpointConfig.ResourceManagerVMDNSSuffix = mockCSPDefaultSpec.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix
	expectedEnv.CloudName = AzureStackCloud
	expectedEnv.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackDeployTelemetryPID
	if diff := cmp.Diff(actualEnv, expectedEnv); diff != "" {
		t.Errorf("setCustomCloudProfileDefaults(): did not set AzureTelemetryPID as DefaultAzureStackDeployTelemetryPID. %s", diff)
	}

	// Test that the AzureTelemetryPID is set to DefaultAzureStackScaleTelemetryPID by in Scale scenario
	mockCSScaleSpec := getMockBaseContainerService("1.11.6")
	mockCSPScaleSpec := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSScaleSpec.Properties.CustomCloudProfile = mockCSPScaleSpec.CustomCloudProfile
	mockCSScaleSpec.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    true,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	actualScaleEnv := AzureCloudSpecEnvMap[AzureStackCloud]
	expectedScaleEnv := AzureCloudSpecEnvMap[AzurePublicCloud]
	expectedScaleEnv.EndpointConfig.ResourceManagerVMDNSSuffix = mockCSPDefaultSpec.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix
	expectedScaleEnv.CloudName = AzureStackCloud
	expectedScaleEnv.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackScaleTelemetryPID
	if diff := cmp.Diff(actualScaleEnv, expectedScaleEnv); diff != "" {
		t.Errorf("setCustomCloudProfileDefaults(): did not set AzureTelemetryPID as DefaultAzureStackDeployTelemetryPID. %s", diff)
	}

	// Test that the AzureTelemetryPID is set to DefaultAzureStackUpgradeTelemetryPID in Upgrade scenario
	mockCSSUpgradeSpec := getMockBaseContainerService("1.11.6")
	mockCSPSUpgradeSpec := GetMockPropertiesWithCustomCloudProfile("azurestackcloud", true, true, false)
	mockCSSUpgradeSpec.Properties.CustomCloudProfile = mockCSPSUpgradeSpec.CustomCloudProfile
	mockCSSUpgradeSpec.SetPropertiesDefaults(PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  true,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})

	actualSUpgradeEnv := AzureCloudSpecEnvMap[AzureStackCloud]
	expectedSUpgradeEnv := AzureCloudSpecEnvMap[AzurePublicCloud]
	expectedSUpgradeEnv.EndpointConfig.ResourceManagerVMDNSSuffix = mockCSPDefaultSpec.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix
	expectedSUpgradeEnv.CloudName = AzureStackCloud
	expectedSUpgradeEnv.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackUpgradeTelemetryPID
	if diff := cmp.Diff(actualSUpgradeEnv, expectedSUpgradeEnv); diff != "" {
		t.Errorf("setCustomCloudProfileDefaults(): did not set AzureTelemetryPID as DefaultAzureStackUpgradeTelemetryPID. %s", diff)
	}
}

func TestImageReference(t *testing.T) {
	cases := []struct {
		name                      string
		cs                        ContainerService
		isUpgrade                 bool
		isScale                   bool
		expectedMasterProfile     MasterProfile
		expectedAgentPoolProfiles []AgentPoolProfile
	}{
		{
			name: "default",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType: Kubernetes,
					},
					MasterProfile: &MasterProfile{},
					AgentPoolProfiles: []*AgentPoolProfile{
						{},
					},
				},
			},
			expectedMasterProfile: MasterProfile{
				Distro:   AKSUbuntu1604,
				ImageRef: nil,
			},
			expectedAgentPoolProfiles: []AgentPoolProfile{
				{
					Distro:   AKSUbuntu1604,
					ImageRef: nil,
				},
			},
		},
		{
			name: "image references",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType: Kubernetes,
					},
					MasterProfile: &MasterProfile{
						ImageRef: &ImageReference{
							Name:           "name",
							ResourceGroup:  "resource-group",
							SubscriptionID: "sub-id",
							Gallery:        "gallery",
							Version:        "version",
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							ImageRef: &ImageReference{
								Name:           "name",
								ResourceGroup:  "resource-group",
								SubscriptionID: "sub-id",
								Gallery:        "gallery",
								Version:        "version",
							},
						},
					},
				},
			},
			expectedMasterProfile: MasterProfile{
				Distro: "",
				ImageRef: &ImageReference{
					Name:           "name",
					ResourceGroup:  "resource-group",
					SubscriptionID: "sub-id",
					Gallery:        "gallery",
					Version:        "version",
				},
			},
			expectedAgentPoolProfiles: []AgentPoolProfile{
				{
					Distro: "",
					ImageRef: &ImageReference{
						Name:           "name",
						ResourceGroup:  "resource-group",
						SubscriptionID: "sub-id",
						Gallery:        "gallery",
						Version:        "version",
					},
				},
			},
		},
		{
			name: "mixed",
			cs: ContainerService{
				Properties: &Properties{
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType: Kubernetes,
					},
					MasterProfile: &MasterProfile{},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							ImageRef: &ImageReference{
								Name:           "name",
								ResourceGroup:  "resource-group",
								SubscriptionID: "sub-id",
								Gallery:        "gallery",
								Version:        "version",
							},
						},
						{},
					},
				},
			},
			expectedMasterProfile: MasterProfile{
				Distro:   AKSUbuntu1604,
				ImageRef: nil,
			},
			expectedAgentPoolProfiles: []AgentPoolProfile{
				{
					Distro: "",
					ImageRef: &ImageReference{
						Name:           "name",
						ResourceGroup:  "resource-group",
						SubscriptionID: "sub-id",
						Gallery:        "gallery",
						Version:        "version",
					},
				},
				{
					Distro:   AKSUbuntu1604,
					ImageRef: nil,
				},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.cs.SetPropertiesDefaults(PropertiesDefaultsParams{
				IsUpgrade:  c.isUpgrade,
				IsScale:    c.isScale,
				PkiKeySize: helpers.DefaultPkiKeySize,
			})
			if c.cs.Properties.MasterProfile.Distro != c.expectedMasterProfile.Distro {
				t.Errorf("expected %s, but got %s", c.expectedMasterProfile.Distro, c.cs.Properties.MasterProfile.Distro)
			}
			if c.expectedMasterProfile.ImageRef == nil {
				if c.cs.Properties.MasterProfile.ImageRef != nil {
					t.Errorf("expected nil, but got an ImageRef")
				}
			} else {
				if c.cs.Properties.MasterProfile.ImageRef == nil {
					t.Errorf("got unexpected nil MasterProfile.ImageRef")
				}
				if c.cs.Properties.MasterProfile.ImageRef.Name != c.expectedMasterProfile.ImageRef.Name {
					t.Errorf("expected %s, but got %s", c.expectedMasterProfile.ImageRef.Name, c.cs.Properties.MasterProfile.ImageRef.Name)
				}
				if c.cs.Properties.MasterProfile.ImageRef.ResourceGroup != c.expectedMasterProfile.ImageRef.ResourceGroup {
					t.Errorf("expected %s, but got %s", c.expectedMasterProfile.ImageRef.ResourceGroup, c.cs.Properties.MasterProfile.ImageRef.ResourceGroup)
				}
				if c.cs.Properties.MasterProfile.ImageRef.SubscriptionID != c.expectedMasterProfile.ImageRef.SubscriptionID {
					t.Errorf("expected %s, but got %s", c.expectedMasterProfile.ImageRef.SubscriptionID, c.cs.Properties.MasterProfile.ImageRef.SubscriptionID)
				}
				if c.cs.Properties.MasterProfile.ImageRef.Gallery != c.expectedMasterProfile.ImageRef.Gallery {
					t.Errorf("expected %s, but got %s", c.expectedMasterProfile.ImageRef.Gallery, c.cs.Properties.MasterProfile.ImageRef.Gallery)
				}
				if c.cs.Properties.MasterProfile.ImageRef.Version != c.expectedMasterProfile.ImageRef.Version {
					t.Errorf("expected %s, but got %s", c.expectedMasterProfile.ImageRef.Version, c.cs.Properties.MasterProfile.ImageRef.Version)
				}
			}
			for i, profile := range c.cs.Properties.AgentPoolProfiles {
				if profile.Distro != c.expectedAgentPoolProfiles[i].Distro {
					t.Errorf("expected %s, but got %s", c.expectedAgentPoolProfiles[i].Distro, profile.Distro)
				}
				if c.expectedAgentPoolProfiles[i].ImageRef == nil {
					if profile.ImageRef != nil {
						t.Errorf("expected nil, but got an ImageRef")
					}
				} else {
					if profile.ImageRef == nil {
						t.Errorf("got unexpected nil MasterProfile.ImageRef")
					}
					if profile.ImageRef.Name != c.expectedAgentPoolProfiles[i].ImageRef.Name {
						t.Errorf("expected %s, but got %s", c.expectedAgentPoolProfiles[i].ImageRef.Name, profile.ImageRef.Name)
					}
					if profile.ImageRef.ResourceGroup != c.expectedAgentPoolProfiles[i].ImageRef.ResourceGroup {
						t.Errorf("expected %s, but got %s", c.expectedAgentPoolProfiles[i].ImageRef.ResourceGroup, profile.ImageRef.ResourceGroup)
					}
					if profile.ImageRef.SubscriptionID != c.expectedAgentPoolProfiles[i].ImageRef.SubscriptionID {
						t.Errorf("expected %s, but got %s", c.expectedAgentPoolProfiles[i].ImageRef.SubscriptionID, profile.ImageRef.SubscriptionID)
					}
					if profile.ImageRef.Gallery != c.expectedAgentPoolProfiles[i].ImageRef.Gallery {
						t.Errorf("expected %s, but got %s", c.expectedAgentPoolProfiles[i].ImageRef.Gallery, profile.ImageRef.Gallery)
					}
					if profile.ImageRef.Version != c.expectedAgentPoolProfiles[i].ImageRef.Version {
						t.Errorf("expected %s, but got %s", c.expectedAgentPoolProfiles[i].ImageRef.Version, profile.ImageRef.Version)
					}
				}
			}
		})
	}
}
