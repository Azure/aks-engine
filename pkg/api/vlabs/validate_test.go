// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/blang/semver"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	ValidKubernetesNodeStatusUpdateFrequency        = "10s"
	ValidKubernetesCtrlMgrNodeMonitorGracePeriod    = "40s"
	ValidKubernetesCtrlMgrPodEvictionTimeout        = "5m0s"
	ValidKubernetesCtrlMgrRouteReconciliationPeriod = "10s"
	ValidKubernetesCloudProviderBackoff             = false
	ValidKubernetesCloudProviderBackoffRetries      = 6
	ValidKubernetesCloudProviderBackoffJitter       = 1
	ValidKubernetesCloudProviderBackoffDuration     = 5
	ValidKubernetesCloudProviderBackoffExponent     = 1.5
	ValidKubernetesCloudProviderRateLimit           = false
	ValidKubernetesCloudProviderRateLimitQPS        = 3
	ValidKubernetesCloudProviderRateLimitBucket     = 10
)

var falseVal = false
var trueVal = true

func Test_OrchestratorProfile_Validate(t *testing.T) {
	tests := map[string]struct {
		properties    *Properties
		expectedError string
		isUpdate      bool
	}{
		"should error when KubernetesConfig has invalid etcd version": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						EtcdVersion: "1.0.0",
					},
				},
			},
			expectedError: "Invalid etcd version \"1.0.0\", please use one of the following versions: [2.2.5 2.3.0 2.3.1 2.3.2 2.3.3 2.3.4 2.3.5 2.3.6 2.3.7 2.3.8 3.0.0 3.0.1 3.0.2 3.0.3 3.0.4 3.0.5 3.0.6 3.0.7 3.0.8 3.0.9 3.0.10 3.0.11 3.0.12 3.0.13 3.0.14 3.0.15 3.0.16 3.0.17 3.1.0 3.1.1 3.1.2 3.1.2 3.1.3 3.1.4 3.1.5 3.1.6 3.1.7 3.1.8 3.1.9 3.1.10 3.2.0 3.2.1 3.2.2 3.2.3 3.2.4 3.2.5 3.2.6 3.2.7 3.2.8 3.2.9 3.2.11 3.2.12 3.2.13 3.2.14 3.2.15 3.2.16 3.2.23 3.2.24 3.2.25 3.2.26 3.3.0 3.3.1 3.3.8 3.3.9 3.3.10 3.3.13 3.3.15 3.3.18 3.3.19 3.3.22]",
		},
		"should error when KubernetesConfig has invalid containerd version for containerd runtime": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:  Containerd,
						ContainerdVersion: "1.0.0",
					},
				},
			},
			expectedError: "Invalid containerd version \"1.0.0\", please use one of the following versions: [1.3.2 1.3.3 1.3.4 1.3.5 1.3.6 1.3.7 1.3.8 1.3.9]",
		},
		"should error when KubernetesConfig has containerdVersion value for docker container runtime": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:  Docker,
						MobyVersion:       "3.0.11",
						ContainerdVersion: "1.3.2",
					},
				},
			},
			expectedError: fmt.Sprintf("containerdVersion is only valid in a non-docker context, use %s containerRuntime value instead if you wish to provide a containerdVersion", Containerd),
		},
		"should error when KubernetesConfig has enableAggregatedAPIs enabled and enableRBAC disabled": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.16.14",
					KubernetesConfig: &KubernetesConfig{
						EnableAggregatedAPIs: true,
						EnableRbac:           &falseVal,
					},
				},
			},
			expectedError: "enableAggregatedAPIs requires the enableRbac feature as a prerequisite",
		},
		"should error when KubernetesConfig has enableRBAC disabled for >= 1.15": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false, false)),
					KubernetesConfig: &KubernetesConfig{
						EnableRbac: &falseVal,
					},
				},
			},
			expectedError: fmt.Sprintf("RBAC support is required for Kubernetes version 1.15.0 or greater; unable to build Kubernetes v%s cluster with enableRbac=false", common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false, false))),
		},
		"should error when KubernetesConfig has enableDataEncryptionAtRest enabled with invalid version": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.6.9",
					KubernetesConfig: &KubernetesConfig{
						EnableDataEncryptionAtRest: &trueVal,
					},
				},
			},
			expectedError: "enableDataEncryptionAtRest is only available in Kubernetes version 1.7.0 or greater; unable to validate for Kubernetes version 1.6.9",
		},
		"should error when KubernetesConfig has enableDataEncryptionAtRest enabled with invalid encryption key": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.16.14",
					KubernetesConfig: &KubernetesConfig{
						EnableDataEncryptionAtRest: &trueVal,
						EtcdEncryptionKey:          "fakeEncryptionKey",
					},
				},
			},
			expectedError: "etcdEncryptionKey must be base64 encoded. Please provide a valid base64 encoded value or leave the etcdEncryptionKey empty to auto-generate the value",
		},
		"should error when KubernetesConfig has enableEncryptionWithExternalKms enabled with invalid version": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.6.9",
					KubernetesConfig: &KubernetesConfig{
						EnableEncryptionWithExternalKms: &trueVal,
					},
				},
			},
			expectedError: "enableEncryptionWithExternalKms is only available in Kubernetes version 1.10.0 or greater; unable to validate for Kubernetes version 1.6.9",
		},
		"should error when KubernetesConfig has Standard loadBalancerSku with invalid version": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.6.9",
					KubernetesConfig: &KubernetesConfig{
						LoadBalancerSku: StandardLoadBalancerSku,
					},
				},
			},
			expectedError: "loadBalancerSku is only available in Kubernetes version 1.11.0 or greater; unable to validate for Kubernetes version 1.6.9",
		},
		"should error when KubernetesConfig has Basic loadBalancerSku with loadBalancerOutboundIPs config": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.18.13",
					KubernetesConfig: &KubernetesConfig{
						LoadBalancerSku:             BasicLoadBalancerSku,
						ExcludeMasterFromStandardLB: to.BoolPtr(true),
						LoadBalancerOutboundIPs:     to.IntPtr(3),
					},
				},
			},
			expectedError: "kubernetesConfig.loadBalancerOutboundIPs configuration only supported for Standard loadBalancerSku=Standard",
		},
		"should error when too many loadBalancerOutboundIPs are configured": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.18.13",
					KubernetesConfig: &KubernetesConfig{
						LoadBalancerOutboundIPs: to.IntPtr(17),
					},
				},
			},
			expectedError: fmt.Sprintf("kubernetesConfig.loadBalancerOutboundIPs was set to %d, the maximum allowed is %d", 17, common.MaxLoadBalancerOutboundIPs),
		},
		"should error when KubernetesConfig has enablePodSecurity enabled with invalid settings": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.16.14",
					KubernetesConfig: &KubernetesConfig{
						EnablePodSecurityPolicy: &trueVal,
					},
				},
			},
			expectedError: "enablePodSecurityPolicy requires the enableRbac feature as a prerequisite",
		},
		"should not error with empty object": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
				},
			},
		},
		"kubernetes should have failed on old patch version": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.6.0",
				},
			},
			expectedError: fmt.Sprint("the following OrchestratorProfile configuration is not supported: OrchestratorType: \"Kubernetes\", OrchestratorRelease: \"\", OrchestratorVersion: \"1.6.0\". Please use one of the following versions: ", common.GetAllSupportedKubernetesVersions(false, false, false)),
		},
		"kubernetes should not fail on old patch version if update": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.6.0",
				},
			},
			isUpdate: true,
		},
		"kubernetes should not have failed on version with v prefix": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "v1.16.14",
				},
			},
		},
		"should error when maximumLoadBalancerRuleCount populated with a negative integer": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						MaximumLoadBalancerRuleCount: -1,
					},
				},
			},
			expectedError: "maximumLoadBalancerRuleCount shouldn't be less than 0",
		},
		"should error when outboundRuleIdleTimeoutInMinutes populated is out of valid range": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						LoadBalancerSku:                  StandardLoadBalancerSku,
						ExcludeMasterFromStandardLB:      to.BoolPtr(true),
						OutboundRuleIdleTimeoutInMinutes: 3,
					},
				},
			},
			expectedError: "outboundRuleIdleTimeoutInMinutes shouldn't be less than 4 or greater than 120",
		},
		"should error when EtcdStorageLimitGB is too low": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						EtcdStorageLimitGB: 1,
					},
				},
			},
			expectedError: "EtcdStorageLimitGB value of 1 is too small, the minimum allowed is 2",
		},
		"should error when using flatcar + Azure CNI + bridge mode": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						NetworkPlugin: "azure",
						NetworkMode:   NetworkModeBridge,
					},
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "flatcarpool",
						Count:  10,
						Distro: Flatcar,
					},
				},
			},
			expectedError: "Flatcar node pools require 'transparent' networkMode with Azure CNI",
		},
		"should not error when using Azure CNI + bridge mode + ubuntu": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						NetworkPlugin: "azure",
						NetworkMode:   NetworkModeBridge,
					},
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "ubuntu1604pool",
						Count:  10,
						Distro: AKSUbuntu1604,
					},
					{
						Name:   "ubuntu1804pool",
						Count:  10,
						Distro: AKSUbuntu1804,
					},
				},
			},
		},
	}

	for testName, test := range tests {
		test := test
		testName := testName
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := test.properties.ValidateOrchestratorProfile(test.isUpdate)

			if test.expectedError == "" && err == nil {
				return
			}
			if test.expectedError == "" && err != nil {
				t.Errorf("%s expected no error but received: %s", testName, err.Error())
				return
			}
			if test.expectedError != "" && err == nil {
				t.Errorf("%s expected error: %s, but received no error", testName, test.expectedError)
				return
			}
			if !strings.Contains(err.Error(), test.expectedError) {
				t.Errorf("%s expected error: %s but received: %s", testName, test.expectedError, err.Error())
			}
		})
	}
}

func ExampleProperties_validateOrchestratorProfile() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	cs := getK8sDefaultContainerService(true)
	cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		EtcdStorageLimitGB: 9,
	}
	if err := cs.Properties.ValidateOrchestratorProfile(false); err != nil {
		log.Error(err)
	}

	cs = getK8sDefaultContainerService(true)
	cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		EnableEncryptionWithExternalKms: to.BoolPtr(true),
	}
	if err := cs.Properties.ValidateOrchestratorProfile(false); err != nil {
		log.Error(err)
	}

	// Output:
	// level=warning msg="EtcdStorageLimitGB of 9 is larger than the recommended maximum of 8"
	// level=warning msg="Clusters with enableEncryptionWithExternalKms=true are not upgradable! You will not be able to upgrade your cluster using `aks-engine upgrade`"
}

func Test_KubernetesConfig_Validate(t *testing.T) {
	// Tests that should pass across all versions
	for _, k8sVersion := range common.GetAllSupportedKubernetesVersions(true, false, false) {
		c := KubernetesConfig{}
		if err := c.Validate(k8sVersion, false, false, false); err != nil {
			t.Errorf("should not error on empty KubernetesConfig: %v, version %s", err, k8sVersion)
		}

		c = KubernetesConfig{
			ClusterSubnet:                "10.120.0.0/16",
			DockerBridgeSubnet:           "10.120.1.0/16",
			MaxPods:                      42,
			CloudProviderBackoff:         to.BoolPtr(ValidKubernetesCloudProviderBackoff),
			CloudProviderBackoffRetries:  ValidKubernetesCloudProviderBackoffRetries,
			CloudProviderBackoffJitter:   ValidKubernetesCloudProviderBackoffJitter,
			CloudProviderBackoffDuration: ValidKubernetesCloudProviderBackoffDuration,
			CloudProviderBackoffExponent: ValidKubernetesCloudProviderBackoffExponent,
			CloudProviderRateLimit:       to.BoolPtr(ValidKubernetesCloudProviderRateLimit),
			CloudProviderRateLimitQPS:    ValidKubernetesCloudProviderRateLimitQPS,
			CloudProviderRateLimitBucket: ValidKubernetesCloudProviderRateLimitBucket,
			KubeletConfig: map[string]string{
				"--node-status-update-frequency": ValidKubernetesNodeStatusUpdateFrequency,
			},
			ControllerManagerConfig: map[string]string{
				"--node-monitor-grace-period":   ValidKubernetesCtrlMgrNodeMonitorGracePeriod,
				"--pod-eviction-timeout":        ValidKubernetesCtrlMgrPodEvictionTimeout,
				"--route-reconciliation-period": ValidKubernetesCtrlMgrRouteReconciliationPeriod,
			},
		}
		if err := c.Validate(k8sVersion, false, false, false); err != nil {
			t.Errorf("should not error on a KubernetesConfig with valid param values: %v", err)
		}

		c = KubernetesConfig{
			ClusterSubnet: "10.16.x.0/invalid",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error on invalid ClusterSubnet")
		}

		c = KubernetesConfig{
			DockerBridgeSubnet: "10.120.1.0/invalid",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error on invalid DockerBridgeSubnet")
		}

		c = KubernetesConfig{
			KubeletConfig: map[string]string{
				"--non-masquerade-cidr": "10.120.1.0/24",
			},
		}
		if err := c.Validate(k8sVersion, false, false, false); err != nil {
			t.Error("should not error on valid --non-masquerade-cidr")
		}

		// Re-implement these tests if we re-introduce --ip-maquerade-cidr
		/*c = KubernetesConfig{
			KubeletConfig: map[string]string{
				"--non-masquerade-cidr": "10.120.1.0/invalid",
			},
		}
		if err := c.Validate(k8sVersion, false); err == nil {
			t.Error("should error on invalid --non-masquerade-cidr")
		}*/

		c = KubernetesConfig{
			MaxPods: KubernetesMinMaxPods - 1,
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error on invalid MaxPods")
		}

		c = KubernetesConfig{
			KubeletConfig: map[string]string{
				"--node-status-update-frequency": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error on invalid --node-status-update-frequency kubelet config")
		}

		c = KubernetesConfig{
			ControllerManagerConfig: map[string]string{
				"--node-monitor-grace-period": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error on invalid --node-monitor-grace-period")
		}

		c = KubernetesConfig{
			ControllerManagerConfig: map[string]string{
				"--node-monitor-grace-period": "30s",
			},
			KubeletConfig: map[string]string{
				"--node-status-update-frequency": "10s",
			},
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when --node-monitor-grace-period is not sufficiently larger than --node-status-update-frequency kubelet config")
		}

		c = KubernetesConfig{
			ControllerManagerConfig: map[string]string{
				"--pod-eviction-timeout": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error on invalid --pod-eviction-timeout")
		}

		c = KubernetesConfig{
			ControllerManagerConfig: map[string]string{
				"--route-reconciliation-period": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error on invalid --route-reconciliation-period")
		}

		c = KubernetesConfig{
			DNSServiceIP: "192.168.0.10",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when DNSServiceIP but not ServiceCidr")
		}

		c = KubernetesConfig{
			ServiceCidr: "192.168.0.10/24",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when ServiceCidr but not DNSServiceIP")
		}

		c = KubernetesConfig{
			DNSServiceIP: "invalid",
			ServiceCidr:  "192.168.0.0/24",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when DNSServiceIP is invalid")
		}

		c = KubernetesConfig{
			DNSServiceIP: "192.168.1.10",
			ServiceCidr:  "192.168.0.0/not-a-len",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when ServiceCidr is invalid")
		}

		c = KubernetesConfig{
			DNSServiceIP: "192.168.1.10",
			ServiceCidr:  "192.168.0.0/24",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when DNSServiceIP is outside of ServiceCidr")
		}

		c = KubernetesConfig{
			DNSServiceIP: "172.99.255.255",
			ServiceCidr:  "172.99.0.1/16",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when DNSServiceIP is broadcast address of ServiceCidr")
		}

		c = KubernetesConfig{
			DNSServiceIP: "172.99.0.1",
			ServiceCidr:  "172.99.0.1/16",
		}
		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when DNSServiceIP is first IP of ServiceCidr")
		}

		c = KubernetesConfig{
			DNSServiceIP: "172.99.255.10",
			ServiceCidr:  "172.99.0.1/16",
		}
		if err := c.Validate(k8sVersion, false, false, false); err != nil {
			t.Error("should not error when DNSServiceIP and ServiceCidr are valid")
		}

		c = KubernetesConfig{
			ClusterSubnet: "192.168.0.1/24",
			NetworkPlugin: "azure",
		}

		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when ClusterSubnet has a mask of 24 bits or higher")
		}

		c = KubernetesConfig{
			ClusterSubnet: "192.168.0.1/24",
			NetworkPlugin: "azure",
		}

		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when ClusterSubnet has a mask of 24 bits or higher")
		}

		c = KubernetesConfig{
			ProxyMode: KubeProxyMode("invalid"),
		}

		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when ProxyMode has an invalid string value")
		}

		for _, validProxyModeValue := range []KubeProxyMode{KubeProxyModeIPTables, KubeProxyModeIPVS} {
			c = KubernetesConfig{
				ProxyMode: validProxyModeValue,
			}

			if err := c.Validate(k8sVersion, false, false, false); err != nil {
				t.Error("should error when ProxyMode has a valid string value")
			}

			c = KubernetesConfig{
				ProxyMode: validProxyModeValue,
			}

			if err := c.Validate(k8sVersion, false, false, false); err != nil {
				t.Error("should error when ProxyMode has a valid string value")
			}
		}
	}

	// Tests that apply to 1.6 and later releases
	for _, k8sVersion := range common.GetAllSupportedKubernetesVersions(false, false, false) {
		c := KubernetesConfig{
			CloudProviderBackoff:   to.BoolPtr(true),
			CloudProviderRateLimit: to.BoolPtr(true),
		}
		if err := c.Validate(k8sVersion, false, false, false); err != nil {
			t.Error("should not error when basic backoff and rate limiting are set to true with no options")
		}
	}

	// Tests that apply to 1.8 and later releases
	for _, k8sVersion := range common.GetVersionsGt(common.GetAllSupportedKubernetesVersions(true, false, false), "1.8.0", true, true) {
		c := KubernetesConfig{
			UseCloudControllerManager: to.BoolPtr(true),
		}
		if err := c.Validate(k8sVersion, false, false, false); err != nil {
			t.Error("should not error because UseCloudControllerManager is available since v1.8")
		}
	}

	// Tests that apply to dualstack with 1.16 and later releases
	for _, k8sVersion := range common.GetVersionsGt(common.GetAllSupportedKubernetesVersions(false, false, false), "1.16.0", true, true) {
		c := KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
		}

		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when more than 1 cluster subnet provided with ipv6dualstack feature disabled")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8,fec0::/7",
		}

		if err := c.Validate(k8sVersion, false, true, false); err == nil {
			t.Error("should error when more than 2 cluster subnets provided")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "iptables",
		}

		if err := c.Validate(k8sVersion, false, true, false); err == nil && !common.IsKubernetesVersionGe(k8sVersion, "1.18.0") {
			t.Errorf("should error with ipv6 dual stack feature enabled as iptables mode not supported in %s", k8sVersion)
		}

		c = KubernetesConfig{
			ServiceCidr: "10.0.0.0/16,fe80:20d::/112",
		}

		if err := c.Validate(k8sVersion, false, false, false); err == nil {
			t.Error("should error when more than 1 service cidr provided with ipv6dualstack feature disabled")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,fe80:20d::/112,fec0::/7",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true, false); err == nil {
			t.Error("should error when more than 2 service cidr provided with ipv6dualstack feature enabled")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,2001:db8::/129",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true, false); err == nil {
			t.Error("should error when secondary cidr is invalid with ipv6dualstack feature enabled")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,fe80:20d::/112",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true, false); err != nil {
			t.Error("shouldn't have errored with ipv6 dual stack feature enabled")
		}

		// validate config with azure cni dual stack and network policy enabled.
		c = KubernetesConfig{
			NetworkPlugin: "azure",
			ClusterSubnet: "10.240.0.0/12,fe80:20d::/112",
			NetworkPolicy: "azure",
		}

		if err := c.Validate(k8sVersion, false, true, false); err == nil {
			t.Errorf("should error when network policy defined for azure cni dual stack: %v", err)
		}

		//validate azure cni dual stack enabled scenario
		c = KubernetesConfig{
			NetworkPlugin: "azure",
			NetworkMode:   "bridge",
			ClusterSubnet: "10.240.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,fe80:20d::/112",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true, false); err != nil {
			t.Errorf("shouldn't have errored with azure cni ipv6 dual stack feature enabled: %v", err)
		}

		// Azure CNI + dualstack requires bridge NetworkMode
		c = KubernetesConfig{
			NetworkPlugin: "azure",
			ClusterSubnet: "10.240.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,fe80:20d::/112",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true, false); err == nil {
			t.Errorf("should error when Azure CNI + dual stack without bridge network mode")
		}

		// Azure CNI + dualstack doesn't work with transparent NetworkMode
		c = KubernetesConfig{
			NetworkPlugin: "azure",
			NetworkMode:   "transparent",
			ClusterSubnet: "10.240.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,fe80:20d::/112",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true, false); err == nil {
			t.Errorf("should error when Azure CNI + dual stack without bridge network mode")
		}
	}

	// Tests that apply to single stack IPv6 with 1.18 and later releases
	for _, k8sVersion := range common.GetVersionsGt(common.GetAllSupportedKubernetesVersions(false, false, false), "1.18.0", true, true) {
		c := KubernetesConfig{
			NetworkPlugin: "azure",
		}
		if err := c.Validate(k8sVersion, false, false, true); err == nil {
			t.Error("should error when network plugin is not kubenet for single stack IPv6")
		}
		if err := c.Validate(k8sVersion, false, true, true); err == nil {
			t.Error("should error when dual stack and single stack IPv6 enabled simultaneously")
		}
	}
}

func Test_Properties_ValidateCustomKubeComponent(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}

	p.OrchestratorProfile.OrchestratorVersion = "1.17.0"
	p.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage = "example.azurecr.io/hyperkube-amd64:tag"
	err := p.validateCustomKubeComponent()
	expectedMsg := "customHyperkubeImage has no effect in Kubernetes version 1.17.0 or above"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}

	p.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeAPIServerImage = "example.azurecr.io/kube-apiserver-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeControllerManagerImage = "example.azurecr.io/kube-controller-manager-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeProxyImage = "example.azurecr.io/kube-proxy-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeSchedulerImage = "example.azurecr.io/kube-scheduler-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeBinaryURL = "https://example.blob.core.windows.net/kubernetes-node-linux-amd64.tar.gz"
	err = p.validateCustomKubeComponent()
	if err != nil {
		t.Errorf("should not error because custom kube components can be used in 1.17, got error : %s", err.Error())
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.16.0"
	err = p.validateCustomKubeComponent()
	expectedMsg = "customKubeAPIServerImage, customKubeControllerManagerImage, customKubeSchedulerImage or customKubeBinaryURL have no effect in Kubernetes version 1.16 or earlier"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.15.0"
	p.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage = "example.azurecr.io/hyperkube-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeAPIServerImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeControllerManagerImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeProxyImage = "example.azurecr.io/kube-proxy-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeSchedulerImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeBinaryURL = ""
	err = p.validateCustomKubeComponent()
	expectedMsg = "customKubeProxyImage has no effect in Kubernetes version 1.15 or earlier"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.16.0"
	err = p.validateCustomKubeComponent()
	if err != nil {
		t.Errorf("should not error because custom kube-proxy and hyperkube components can be used in 1.16, got error : %s", err.Error())
	}
}

func Test_Properties_ValidateDistro(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes
	p.MasterProfile = &MasterProfile{
		DNSPrefix: "foo",
	}

	// Should not error on valid distros in non-update scenarios
	for _, distro := range DistroValues {
		p.MasterProfile.Distro = distro
		p.AgentPoolProfiles = []*AgentPoolProfile{
			{
				Name:   "pool1",
				Distro: distro,
			},
		}
		if err := p.validateMasterProfile(false); err != nil {
			t.Errorf(
				"should not error on distro=\"%s\"",
				distro,
			)
		}
		if err := p.validateAgentPoolProfiles(false); err != nil {
			t.Errorf(
				"should not error on distro=\"%s\"",
				distro,
			)
		}
	}

	// Should not error on valid distros in update scenarios
	for _, distro := range DistroValues {
		p.MasterProfile.Distro = distro
		p.AgentPoolProfiles = []*AgentPoolProfile{
			{
				Name:   "pool1",
				Distro: distro,
			},
		}
		if err := p.validateMasterProfile(true); err != nil {
			t.Errorf(
				"should not error on distro=\"%s\"",
				distro,
			)
		}
		if err := p.validateAgentPoolProfiles(true); err != nil {
			t.Errorf(
				"should not error on distro=\"%s\"",
				distro,
			)
		}
	}

	// Should error for invalid distros on non-update scenarios
	bogusDistroValues := []Distro{AKSDockerEngine, AKS1604Deprecated, AKS1804Deprecated, "bogon"}
	for _, distro := range bogusDistroValues {
		p.MasterProfile.Distro = distro
		p.AgentPoolProfiles = []*AgentPoolProfile{
			{
				Name:   "pool1",
				Distro: distro,
			},
		}
		if err := p.validateMasterProfile(false); err == nil {
			t.Errorf(
				"should error on distro=\"%s\"",
				distro,
			)
		}
		if err := p.validateAgentPoolProfiles(false); err == nil {
			t.Errorf(
				"should error on distro=\"%s\"",
				distro,
			)
		}
	}

	// Should not error for deprecated distro on update scenarios
	oldDistros := []Distro{AKSDockerEngine, AKS1604Deprecated, AKS1804Deprecated}
	for _, distro := range oldDistros {
		p.MasterProfile.Distro = distro
		p.AgentPoolProfiles = []*AgentPoolProfile{
			{
				Name:   "pool1",
				Distro: distro,
			},
		}
		if err := p.validateMasterProfile(true); err != nil {
			t.Errorf(
				"should error on distro=\"%s\"",
				distro,
			)
		}
		if err := p.validateAgentPoolProfiles(true); err != nil {
			t.Errorf(
				"should error on distro=\"%s\"",
				distro,
			)
		}
	}
}

func Test_Properties_ValidateNetworkPolicy(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes

	k8sVersion := "1.8.0"
	for _, policy := range NetworkPolicyValues {
		p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}
		p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = policy
		if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPolicy(k8sVersion, false); err != nil {
			t.Errorf(
				"should not error on networkPolicy=\"%s\" on k8sVersion=\"%s\"",
				policy,
				k8sVersion,
			)
		}
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = "not-existing"
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPolicy(k8sVersion, false); err == nil {
		t.Errorf(
			"should error on invalid networkPolicy",
		)
	}

	k8sVersion = "1.7.9"
	p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = "azure"
	p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = "azure"
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPolicy(k8sVersion, false); err == nil {
		t.Errorf(
			"should error on azure networkPolicy + azure networkPlugin with k8s version < 1.8.0",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = "calico"
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPolicy(k8sVersion, true); err == nil {
		t.Errorf(
			"should error on calico for windows clusters",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyCilium
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPolicy(k8sVersion, true); err == nil {
		t.Errorf(
			"should error on cilium for windows clusters",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyAntrea
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPolicy(k8sVersion, true); err == nil {
		t.Errorf(
			"should error on antrea for windows clusters",
		)
	}
}

func ExampleKubernetesConfig_validateNetworkPlugin() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	cs := getK8sDefaultContainerService(true)

	cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
	if err := cs.Properties.OrchestratorProfile.KubernetesConfig.validateNetworkPlugin(true); err != nil {
		fmt.Printf("error in ValidateNetworkPlugin: %s", err)
	}

	// Output:
	// level=warning msg="Windows + Kubenet is for development and testing only, not recommended for production"
}

func Test_Properties_ValidateNetworkPlugin(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes

	for _, policy := range NetworkPluginValues {
		p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}
		p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = policy
		if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPlugin(false); err != nil {
			t.Errorf(
				"should not error on networkPolicy=\"%s\"",
				policy,
			)
		}
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = "not-existing"
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPlugin(false); err == nil {
		t.Errorf(
			"should error on invalid networkPlugin",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAntrea
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPlugin(true); err == nil {
		t.Errorf(
			"should error on antrea for windows clusters",
		)
	}
}

func Test_Properties_ValidateNetworkPluginPlusPolicy(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes

	for _, config := range networkPluginPlusPolicyAllowed {
		p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}
		p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = config.networkPlugin
		p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = config.networkPolicy
		if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPluginPlusPolicy(); err != nil {
			t.Errorf(
				"should not error on networkPolicy=\"%s\" + networkPlugin=\"%s\"",
				config.networkPolicy, config.networkPlugin,
			)
		}
	}

	for _, config := range []k8sNetworkConfig{
		{
			networkPlugin: "azure",
			networkPolicy: NetworkPolicyCilium,
		},
		{
			networkPlugin: "azure",
			networkPolicy: "flannel",
		},
		{
			networkPlugin: "flannel",
			networkPolicy: "flannel",
		},
		{
			networkPlugin: "flannel",
			networkPolicy: "calico",
		},
		{
			networkPlugin: "kubenet",
			networkPolicy: "none",
		},
		{
			networkPlugin: "azure",
			networkPolicy: "none",
		},
		{
			networkPlugin: "kubenet",
			networkPolicy: "kubenet",
		},
		{
			networkPlugin: "cilium",
		},
	} {
		p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}
		p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = config.networkPlugin
		p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = config.networkPolicy
		if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkPluginPlusPolicy(); err == nil {
			t.Errorf(
				"should error on networkPolicy=\"%s\" + networkPlugin=\"%s\"",
				config.networkPolicy, config.networkPlugin,
			)
		}
	}
}

func Test_Properties_ValidateNetworkMode(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}

	p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = "azure"
	p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = ""
	for _, mode := range NetworkModeValues {
		if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkMode(); err != nil {
			t.Errorf(
				"should not error on networkMode=\"%s\" + networkPlugin=\"azure\" + networkPolicy=\"\"",
				mode,
			)
		}
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkMode = "not-existing"
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkMode(); err == nil {
		t.Errorf(
			"should error on invalid networkMode",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPolicy = "calico"
	p.OrchestratorProfile.KubernetesConfig.NetworkMode = "bridge"
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkMode(); err == nil {
		t.Errorf(
			"should error on networkPolicy=\"calico\" + networkMode=\"bridge\"",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.NetworkPlugin = Kubernetes
	p.OrchestratorProfile.KubernetesConfig.NetworkMode = "bridge"
	if err := p.OrchestratorProfile.KubernetesConfig.validateNetworkMode(); err == nil {
		t.Errorf(
			"should error on networkPlugin=\"kubenet\"",
		)
	}
}

func TestProperties_ValidateLinuxProfile(t *testing.T) {
	cs := getK8sDefaultContainerService(true)
	cs.Properties.LinuxProfile.SSH = struct {
		PublicKeys []PublicKey `json:"publicKeys" validate:"required,min=1"`
	}{
		PublicKeys: []PublicKey{{}},
	}
	expectedMsg := "KeyData in LinuxProfile.SSH.PublicKeys cannot be empty string"
	err := cs.Validate(true)

	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}

	cs.Properties.LinuxProfile.SSH = struct {
		PublicKeys []PublicKey `json:"publicKeys" validate:"required,min=1"`
	}{
		PublicKeys: []PublicKey{
			{
				KeyData: "not empty",
			},
			{},
		},
	}
	expectedMsg = "KeyData in LinuxProfile.SSH.PublicKeys cannot be empty string"
	err = cs.Validate(true)

	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}
}

func TestProperties_ValidateWindowsProfile(t *testing.T) {
	var trueVar = true
	tests := []struct {
		name          string
		k8sVersion    string
		wp            *WindowsProfile
		isUpdate      bool
		expectedError error
	}{
		{
			name:       "Valid WindowsProfile",
			k8sVersion: common.RationalizeReleaseAndVersion(common.Kubernetes, "1.17", "", false, false, false),
			wp: &WindowsProfile{
				AdminUsername: "AzureUser",
				AdminPassword: "replacePassword1234$",
			},
			expectedError: nil,
		},
		{
			name:       "No username",
			k8sVersion: common.RationalizeReleaseAndVersion(common.Kubernetes, "1.17", "", false, false, false),
			wp: &WindowsProfile{
				AdminUsername: "",
				AdminPassword: "replacePassword1234$",
			},
			expectedError: errors.New("WindowsProfile.AdminUsername is required, when agent pool specifies Windows"),
		},
		{
			name:       "No password",
			k8sVersion: common.RationalizeReleaseAndVersion(common.Kubernetes, "1.17", "", false, false, false),
			wp: &WindowsProfile{
				AdminUsername: "AzureUser",
				AdminPassword: "",
			},
			expectedError: errors.New("WindowsProfile.AdminPassword is required, when agent pool specifies Windows"),
		},
		{
			name:       "CSI proxy enabled",
			k8sVersion: common.RationalizeReleaseAndVersion(common.Kubernetes, "1.18", "", false, false, false),
			wp: &WindowsProfile{
				AdminUsername:  "AzureUser",
				AdminPassword:  "replacePassword1234$",
				EnableCSIProxy: &trueVar,
				CSIProxyURL:    "http://some/url",
			},
			expectedError: nil,
		},
		{
			name:       "CSI Proxy unsupported version",
			k8sVersion: common.RationalizeReleaseAndVersion(common.Kubernetes, "1.17", "", false, false, false),
			wp: &WindowsProfile{
				AdminUsername:  "AzureUser",
				AdminPassword:  "replacePassword1234$",
				EnableCSIProxy: &trueVar,
				CSIProxyURL:    "http://some/url",
			},
			expectedError: errors.New("CSI proxy for Windows is only available in Kubernetes versions 1.18.0 or greater"),
		},
		{
			name:       "Invalid Windows version",
			k8sVersion: "1.15.7",
			wp: &WindowsProfile{
				AdminUsername: "AzureUser",
				AdminPassword: "replacePassword1234$",
			},
			expectedError: errors.New("Orchestrator Kubernetes version 1.15.7 does not support Windows"),
		},
		{
			name:       "Old Windows version during upgrade",
			k8sVersion: "1.15.7",
			wp: &WindowsProfile{
				AdminUsername: "AzureUser",
				AdminPassword: "replacePassword1234$",
			},
			isUpdate:      true,
			expectedError: nil,
		},
		{
			name: "wrong runtime",
			wp: &WindowsProfile{
				AdminUsername: "azure",
				AdminPassword: "replacePassword1234$",
				WindowsRuntimes: &WindowsRuntimes{
					Default: "something",
				},
			},
			expectedError: errors.New("Default runtime types are process or hyperv"),
		},
		{
			name: "process runtime",
			wp: &WindowsProfile{
				AdminUsername: "azure",
				AdminPassword: "replacePassword1234$",
				WindowsRuntimes: &WindowsRuntimes{
					Default: "process",
				},
			},
			expectedError: nil,
		},
		{
			name: "hyperv runtime",
			wp: &WindowsProfile{
				AdminUsername: "azure",
				AdminPassword: "replacePassword1234$",
				WindowsRuntimes: &WindowsRuntimes{
					Default: "process",
				},
			},
			expectedError: nil,
		},
		{
			name: "invalid runtime handler name",
			wp: &WindowsProfile{
				AdminUsername: "azure",
				AdminPassword: "replacePassword1234$",
				WindowsRuntimes: &WindowsRuntimes{
					Default: "process",
					HypervRuntimes: []RuntimeHandlers{
						{BuildNumber: "something"},
					},
				},
			},
			expectedError: errors.New("Current hyper-v build id values supported are 17763, 18362, 18363, 19041"),
		},
		{
			name: "valid handler names",
			wp: &WindowsProfile{
				AdminUsername: "azure",
				AdminPassword: "replacePassword1234$",
				WindowsRuntimes: &WindowsRuntimes{
					Default: "process",
					HypervRuntimes: []RuntimeHandlers{
						{BuildNumber: "17763"},
						{BuildNumber: "18362"},
						{BuildNumber: "18363"},
						{BuildNumber: "19041"},
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "some valid handlers some not",
			wp: &WindowsProfile{
				AdminUsername: "azure",
				AdminPassword: "replacePassword1234$",
				WindowsRuntimes: &WindowsRuntimes{
					Default: "process",
					HypervRuntimes: []RuntimeHandlers{
						{BuildNumber: "17763"},
						{BuildNumber: "18362"},
						{BuildNumber: "invalid"},
						{BuildNumber: "19041"},
					},
				},
			},
			expectedError: errors.New("Current hyper-v build id values supported are 17763, 18362, 18363, 19041"),
		},
		{
			name: "valid handlers must be unique",
			wp: &WindowsProfile{
				AdminUsername: "azure",
				AdminPassword: "replacePassword1234$",
				WindowsRuntimes: &WindowsRuntimes{
					Default: "process",
					HypervRuntimes: []RuntimeHandlers{
						{BuildNumber: "17763"},
						{BuildNumber: "18362"},
						{BuildNumber: "18363"},
						{BuildNumber: "17763"},
					},
				},
			},
			expectedError: errors.New("Hyper-v RuntimeHandlers have duplicate runtime with build number '17763', Windows Runtimes must be unique"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.OrchestratorProfile.OrchestratorVersion = test.k8sVersion
			cs.Properties.WindowsProfile = test.wp
			err := cs.Properties.validateWindowsProfile(test.isUpdate)
			if !helpers.EqualError(err, test.expectedError) {
				t.Errorf("expected error : '%v', but got '%v'", test.expectedError, err)
			}
		})
	}
}

func TestProperties_ValidateInvalidExtensions(t *testing.T) {
	tests := []struct {
		name              string
		agentPoolProfiles []*AgentPoolProfile
		expectedErr       error
	}{
		{
			name: "Extensions for VirtualMachineScaleSets",
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: VirtualMachineScaleSets,
					Extensions: []Extension{
						{
							Name:        "extensionName",
							SingleOrAll: "single",
							Template:    "fakeTemplate",
						},
					},
				},
			},
			expectedErr: errors.New("Extensions are currently not supported with VirtualMachineScaleSets. Please specify \"availabilityProfile\": \"AvailabilitySet\""),
		},
		{
			name: "prometheus-grafana-k8s extensions for Winows agents",
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					OSType:              Windows,
					Extensions: []Extension{
						{
							Name: "prometheus-grafana-k8s",
						},
					},
				},
			},
			expectedErr: errors.New("prometheus-grafana-k8s extension is currently not supported for Windows agents"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Validate(true)
			if !helpers.EqualError(err, test.expectedErr) {
				t.Errorf("expected error with message : %s, but got %s", test.expectedErr.Error(), err.Error())
			}
		})
	}

}

func TestProperties_ValidateInvalidExtensionProfiles(t *testing.T) {
	tests := []struct {
		name              string
		extensionProfiles []*ExtensionProfile
		expectedErr       error
	}{
		{
			name: "Extension Profile without Keyvault ID",
			extensionProfiles: []*ExtensionProfile{
				{
					Name: "FakeExtensionProfile",
					ExtensionParametersKeyVaultRef: &KeyvaultSecretRef{
						VaultID:    "",
						SecretName: "fakeSecret",
					},
				},
			},
			expectedErr: errors.New("the Keyvault ID must be specified for Extension FakeExtensionProfile"),
		},
		{
			name: "Extension Profile without Keyvault Secret",
			extensionProfiles: []*ExtensionProfile{
				{
					Name: "FakeExtensionProfile",
					ExtensionParametersKeyVaultRef: &KeyvaultSecretRef{
						VaultID:    "fakeVaultID",
						SecretName: "",
					},
				},
			},
			expectedErr: errors.New("the Keyvault Secret must be specified for Extension FakeExtensionProfile"),
		},
		{
			name: "Extension Profile with invalid secret format",
			extensionProfiles: []*ExtensionProfile{
				{
					Name: "FakeExtensionProfile",
					ExtensionParametersKeyVaultRef: &KeyvaultSecretRef{
						VaultID:    "fakeVaultID",
						SecretName: "fakeSecret",
					},
				},
			},
			expectedErr: errors.New("Extension FakeExtensionProfile's keyvault secret reference is of incorrect format"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.ExtensionProfiles = test.extensionProfiles
			err := cs.Validate(true)
			if !helpers.EqualError(err, test.expectedErr) {
				t.Errorf("expected error with message : %s, but got %s", test.expectedErr.Error(), err.Error())
			}
		})
	}
}

func Test_ServicePrincipalProfile_ValidateSecretOrKeyvaultSecretRef(t *testing.T) {

	t.Run("ServicePrincipalProfile with secret should pass", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
			UseManagedIdentity: to.BoolPtr(false),
		}

		if err := cs.Validate(false); err != nil {
			t.Errorf("should not error %v", err)
		}
	})

	t.Run("ServicePrincipalProfile with KeyvaultSecretRef (with version) should pass", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.ServicePrincipalProfile.Secret = ""
		cs.Properties.ServicePrincipalProfile.KeyvaultSecretRef = &KeyvaultSecretRef{
			VaultID:       "/subscriptions/SUB-ID/resourceGroups/RG-NAME/providers/Microsoft.KeyVault/vaults/KV-NAME",
			SecretName:    "secret-name",
			SecretVersion: "version",
		}
		cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
			UseManagedIdentity: to.BoolPtr(false),
		}
		if err := cs.Validate(false); err != nil {
			t.Errorf("should not error %v", err)
		}
	})

	t.Run("ServicePrincipalProfile with KeyvaultSecretRef (without version) should pass", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.ServicePrincipalProfile.Secret = ""
		cs.Properties.ServicePrincipalProfile.KeyvaultSecretRef = &KeyvaultSecretRef{
			VaultID:    "/subscriptions/SUB-ID/resourceGroups/RG-NAME/providers/Microsoft.KeyVault/vaults/KV-NAME",
			SecretName: "secret-name",
		}
		cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
			UseManagedIdentity: to.BoolPtr(false),
		}

		if err := cs.Validate(false); err != nil {
			t.Errorf("should not error %v", err)
		}
	})

	t.Run("ServicePrincipalProfile with Secret and KeyvaultSecretRef should NOT pass", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.ServicePrincipalProfile.Secret = "secret"
		cs.Properties.ServicePrincipalProfile.KeyvaultSecretRef = &KeyvaultSecretRef{
			VaultID:    "/subscriptions/SUB-ID/resourceGroups/RG-NAME/providers/Microsoft.KeyVault/vaults/KV-NAME",
			SecretName: "secret-name",
		}
		cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
			UseManagedIdentity: to.BoolPtr(false),
		}

		if err := cs.Validate(false); err == nil {
			t.Error("error should have occurred")
		}
	})

	t.Run("ServicePrincipalProfile with incorrect KeyvaultSecretRef format should NOT pass", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.ServicePrincipalProfile.Secret = ""
		cs.Properties.ServicePrincipalProfile.KeyvaultSecretRef = &KeyvaultSecretRef{
			VaultID:    "randomID",
			SecretName: "secret-name",
		}
		cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
			UseManagedIdentity: to.BoolPtr(false),
		}

		if err := cs.Validate(false); err == nil || err.Error() != "service principal client keyvault secret reference is of incorrect format" {
			t.Error("error should have occurred")
		}
	})
}

func TestValidateKubernetesLabelValue(t *testing.T) {

	validLabelValues := []string{"", "a", "a1", "this--valid--label--is--exactly--sixty--three--characters--long", "123456", "my-label_valid.com"}
	invalidLabelValues := []string{"a$$b", "-abc", "not.valid.", "This____long____label___is______sixty______four_____chararacters", "Label with spaces"}

	for _, l := range validLabelValues {
		if err := validateKubernetesLabelValue(l); err != nil {
			t.Fatalf("Label value %v should not return error: %v", l, err)
		}
	}

	for _, l := range invalidLabelValues {
		if err := validateKubernetesLabelValue(l); err == nil {
			t.Fatalf("Label value %v should return an error", l)
		}
	}
}

func TestValidateKubernetesLabelKey(t *testing.T) {

	validLabelKeys := []string{"a", "a1", "this--valid--label--is--exactly--sixty--three--characters--long", "123456", "my-label_valid.com", "foo.bar/name", "1.2321.324/key_name.foo", "valid.long.253.characters.label.key.prefix.12345678910.fooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo/my-key"}
	invalidLabelKeys := []string{"", "a/b/c", ".startswithdot", "spaces in key", "foo/", "/name", "$.$/com", "too-long-254-characters-key-prefix-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------123/name", "wrong-slash\\foo"}

	for _, l := range validLabelKeys {
		if err := validateKubernetesLabelKey(l); err != nil {
			t.Fatalf("Label key %v should not return error: %v", l, err)
		}
	}

	for _, l := range invalidLabelKeys {
		if err := validateKubernetesLabelKey(l); err == nil {
			t.Fatalf("Label key %v should return an error", l)
		}
	}
}

func Test_AadProfile_Validate(t *testing.T) {
	t.Run("Valid aadProfile should pass", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		for _, aadProfile := range []*AADProfile{
			{
				ClientAppID: "92444486-5bc3-4291-818b-d53ae480991b",
				ServerAppID: "403f018b-4d89-495b-b548-0cf9868cdb0a",
			},
			{
				ClientAppID: "92444486-5bc3-4291-818b-d53ae480991b",
				ServerAppID: "403f018b-4d89-495b-b548-0cf9868cdb0a",
				TenantID:    "feb784f6-7174-46da-aeae-da66e80c7a11",
			},
		} {
			cs.Properties.AADProfile = aadProfile
			if err := cs.Properties.validateAADProfile(); err != nil {
				t.Errorf("should not error %v", err)
			}
		}
	})

	t.Run("Invalid aadProfiles should NOT pass", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		for _, aadProfile := range []*AADProfile{
			{
				ClientAppID: "1",
				ServerAppID: "d",
			},
			{
				ClientAppID: "6a247d73-ae33-4559-8e5d-4001fdc17b15",
			},
			{
				ClientAppID: "92444486-5bc3-4291-818b-d53ae480991b",
				ServerAppID: "403f018b-4d89-495b-b548-0cf9868cdb0a",
				TenantID:    "1",
			},
			{
				ClientAppID:  "92444486-5bc3-4291-818b-d53ae480991b",
				ServerAppID:  "403f018b-4d89-495b-b548-0cf9868cdb0a",
				TenantID:     "feb784f6-7174-46da-aeae-da66e80c7a11",
				AdminGroupID: "1",
			},
			{},
		} {
			cs.Properties.AADProfile = aadProfile
			if err := cs.Validate(true); err == nil {
				t.Errorf("error should have occurred")
			}
		}
	})
}

func getK8sDefaultContainerService(hasWindows bool) *ContainerService {
	p := &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		MasterProfile: &MasterProfile{
			Count:     1,
			DNSPrefix: "foo",
			VMSize:    "Standard_DS2_v2",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: AvailabilitySet,
			},
		},
		LinuxProfile: &LinuxProfile{
			AdminUsername: "azureuser",
			SSH: struct {
				PublicKeys []PublicKey `json:"publicKeys" validate:"required,min=1"`
			}{
				PublicKeys: []PublicKey{{
					KeyData: "publickeydata",
				}},
			},
		},
		ServicePrincipalProfile: &ServicePrincipalProfile{
			ClientID: "clientID",
			Secret:   "clientSecret",
		},
	}

	if hasWindows {
		p.AgentPoolProfiles = []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: AvailabilitySet,
				OSType:              Windows,
			},
		}
		p.WindowsProfile = &WindowsProfile{
			AdminUsername: "azureuser",
			AdminPassword: "replacepassword1234$",
		}
	}
	cs := ContainerService{
		Location:   "westus",
		Properties: p,
	}
	return &cs
}

func Test_Properties_ValidateContainerRuntime(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes

	for _, runtime := range ContainerRuntimeValues {
		p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}
		p.OrchestratorProfile.KubernetesConfig.ContainerRuntime = runtime
		if err := p.validateContainerRuntime(false); err != nil {
			t.Errorf(
				"should not error on containerRuntime=\"%s\"",
				runtime,
			)
		}
	}

	p.OrchestratorProfile.KubernetesConfig.ContainerRuntime = "not-existing"
	if err := p.validateContainerRuntime(false); err == nil {
		t.Errorf(
			"should error on invalid containerRuntime",
		)
	}

	// Expect deprecated error on upgrading w/ kata-containers
	p.OrchestratorProfile.KubernetesConfig.ContainerRuntime = KataContainers
	if err := p.validateContainerRuntime(true); err == nil {
		t.Errorf(
			"%s containerRuntime has been deprecated, you will not be able to update this cluster with this version of aks-engine", KataContainers,
		)
	}
}

func TestProperties_ValidateContainerRuntime_Windows(t *testing.T) {
	tests := []struct {
		name        string
		p           *Properties
		expectedErr error
	}{
		{
			name: "Docker",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime: Docker,
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "ContainerD-AzureCNI",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:     Docker,
						NetworkPlugin:        "Azure",
						WindowsContainerdURL: "http://some/url",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "ContainerD-AzureCNI-NoWindowsContainerdURL",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:     Containerd,
						NetworkPlugin:        "Azure",
						WindowsContainerdURL: "",
					},
				},
			},
			expectedErr: errors.Errorf("WindowsContainerdURL must be provided when using Windows with ContainerRuntime=containerd"),
		},
		{
			name: "ContainerD-kubenet",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:     Containerd,
						NetworkPlugin:        "kubenet",
						WindowsContainerdURL: "http://some/url",
						WindowsSdnPluginURL:  "http://some/url",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "ContainerD-kubenet-NoWindowsContainerdURL",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:     Containerd,
						NetworkPlugin:        "kubenet",
						WindowsContainerdURL: "",
					},
				},
			},
			expectedErr: errors.Errorf("WindowsContainerdURL must be provided when using Windows with ContainerRuntime=containerd"),
		},
		{
			name: "ContainerD-kubenet-NoWindowsSdnPluginURL",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:     Containerd,
						NetworkPlugin:        "kubenet",
						WindowsContainerdURL: "http://some/url",
						WindowsSdnPluginURL:  "",
					},
				},
			},
			expectedErr: errors.Errorf("WindowsSdnPluginURL must be provided when using Windows with ContainerRuntime=containerd and networkPlugin=kubenet"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			//	t.Parallel()

			// Add a Windows agent pool
			test.p.AgentPoolProfiles = []*AgentPoolProfile{
				{
					OSType: Windows,
				},
			}

			returnedErr := test.p.validateContainerRuntime(false)

			if !helpers.EqualError(returnedErr, test.expectedErr) {
				t.Errorf("Expected error: %v, Got error: %v", test.expectedErr, returnedErr)
			}
		})
	}
}

func TestValidateAddons(t *testing.T) {
	tests := []struct {
		name        string
		p           *Properties
		expectedErr error
	}{
		{
			name: "aad addon enabled w/ no AADProfile",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    "aad",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.New("aad addon can't be enabled without a valid aadProfile w/ adminGroupID"),
		},
		{
			name: "aad addon enabled w/ no AADProfile.AdminGroupID",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    "aad",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
				AADProfile: &AADProfile{},
			},
			expectedErr: errors.New("aad addon can't be enabled without a valid aadProfile w/ adminGroupID"),
		},
		{
			name: "valid aad addon enabled spec",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    "aad",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
				AADProfile: &AADProfile{
					AdminGroupID: "7d04bcd3-3c48-49ab-a064-c0b7d69896da",
				},
			},
			expectedErr: nil,
		},
		{
			name: "cilium addon enabled w/ no networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    "cilium",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.CiliumAddonName, NetworkPolicyCilium),
		},
		{
			name: "cilium addon enabled w/ azure networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: "azure",
						Addons: []KubernetesAddon{
							{
								Name:    "cilium",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.CiliumAddonName, NetworkPolicyCilium),
		},
		{
			name: "cilium addon enabled w/ calico networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: "calico",
						Addons: []KubernetesAddon{
							{
								Name:    "cilium",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.CiliumAddonName, NetworkPolicyCilium),
		},
		{
			name: "cilium addon enabled w/ cilium networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: NetworkPolicyCilium,
						Addons: []KubernetesAddon{
							{
								Name:    "cilium",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "cilium addon enabled w/ cilium networkPolicy + networkPlugin",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: NetworkPolicyCilium,
						NetworkPlugin: NetworkPluginCilium,
						Addons: []KubernetesAddon{
							{
								Name:    "cilium",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "cilium addon enabled w/ k8s >= 1.16",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.0",
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: NetworkPolicyCilium,
						Addons: []KubernetesAddon{
							{
								Name:    "cilium",
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon is not supported on Kubernetes v1.16.0 or greater", common.CiliumAddonName),
		},
		{
			name: "antrea addon enabled w/ no networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.AntreaAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.AntreaAddonName, NetworkPolicyAntrea),
		},
		{
			name: "antrea addon enabled w/ azure networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: "azure",
						Addons: []KubernetesAddon{
							{
								Name:    common.AntreaAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.AntreaAddonName, NetworkPolicyAntrea),
		},
		{
			name: "antrea addon enabled w/ calico networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: "calico",
						Addons: []KubernetesAddon{
							{
								Name:    common.AntreaAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.AntreaAddonName, NetworkPolicyAntrea),
		},
		{
			name: "antrea addon enabled w/ antrea networkPolicy",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: NetworkPolicyAntrea,
						Addons: []KubernetesAddon{
							{
								Name:    common.AntreaAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "antrea addon enabled w/ antrea networkPolicy + networkPlugin",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: NetworkPolicyAntrea,
						NetworkPlugin: NetworkPluginAntrea,
						Addons: []KubernetesAddon{
							{
								Name:    common.AntreaAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "flannel addon enabled",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime: Containerd,
						Addons: []KubernetesAddon{
							{
								Name:    common.FlannelAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "flannel addon enabled but no containerRuntime",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.FlannelAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon is only supported with containerRuntime=%s", common.FlannelAddonName, Containerd),
		},
		{
			name: "flannel addon enabled with docker",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime: Docker,
						Addons: []KubernetesAddon{
							{
								Name:    common.FlannelAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon is only supported with containerRuntime=%s", common.FlannelAddonName, Containerd),
		},
		{
			name: "flannel addon enabled w/ NetworkPlugin=flannel",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime: Containerd,
						NetworkPlugin:    NetworkPluginFlannel,
						Addons: []KubernetesAddon{
							{
								Name:    common.FlannelAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "flannel addon enabled w/ NetworkPlugin=azure",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime: Containerd,
						NetworkPlugin:    DefaultNetworkPlugin,
						Addons: []KubernetesAddon{
							{
								Name:    common.FlannelAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon is not supported with networkPlugin=%s, please use networkPlugin=%s", common.FlannelAddonName, DefaultNetworkPlugin, NetworkPluginFlannel),
		},
		{
			name: "flannel addon enabled w/ NetworkPlugin=kubenet",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime: Containerd,
						NetworkPlugin:    "kubenet",
						Addons: []KubernetesAddon{
							{
								Name:    common.FlannelAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon is not supported with networkPlugin=%s, please use networkPlugin=%s", common.FlannelAddonName, "kubenet", NetworkPluginFlannel),
		},
		{
			name: "flannel addon enabled w/ NetworkPolicy=calico",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime: Containerd,
						NetworkPolicy:    "calico",
						Addons: []KubernetesAddon{
							{
								Name:    common.FlannelAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s addon does not support NetworkPolicy, replace %s with \"\"", common.FlannelAddonName, "calico"),
		},
		{
			name: "azure-cloud-provider addon disabled",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPolicy: NetworkPolicyAntrea,
						NetworkPlugin: NetworkPluginAntrea,
						Addons: []KubernetesAddon{
							{
								Name:    common.AzureCloudProviderAddonName,
								Enabled: to.BoolPtr(false),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s add-on is required, it cannot be disabled", common.AzureCloudProviderAddonName),
		},
		{
			name: "csi-secrets-store enabled with version < 1.16",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.SecretsStoreCSIDriverAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("%s add-on can only be used in 1.16+", common.SecretsStoreCSIDriverAddonName),
		},
		{
			name: "keyvault-flexvolume and csi-secrets-store addons enabled",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.0",
					KubernetesConfig: &KubernetesConfig{
						Addons: []KubernetesAddon{
							{
								Name:    common.KeyVaultFlexVolumeAddonName,
								Enabled: to.BoolPtr(true),
							},
							{
								Name:    common.SecretsStoreCSIDriverAddonName,
								Enabled: to.BoolPtr(true),
							},
						},
					},
				},
			},
			expectedErr: errors.Errorf("Both %s and %s addons are enabled, only one of these may be enabled on a cluster", common.KeyVaultFlexVolumeAddonName, common.SecretsStoreCSIDriverAddonName),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			gotErr := test.p.validateAddons()
			if !helpers.EqualError(gotErr, test.expectedErr) {
				t.Logf("scenario %q", test.name)
				t.Errorf("expected error: %v, got: %v", test.expectedErr, gotErr)
			}
		})
	}
}

// TODO move these to TestValidateAddons above
func Test_Properties_ValidateAddons(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "cluster-autoscaler",
				Enabled: to.BoolPtr(true),
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			AvailabilityProfile: AvailabilitySet,
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error on cluster-autoscaler with availability sets",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "azure-policy",
				Enabled: to.BoolPtr(true),
			},
		},
	}
	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error on azure-policy when ServicePrincipalProfile is empty",
		)
	}
	p.ServicePrincipalProfile = &ServicePrincipalProfile{
		ClientID: "123",
	}
	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error on azure-policy when ServicePrincipalProfile is not empty",
		)
	}

	p.OrchestratorProfile.OrchestratorRelease = "1.12"
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error on azure-policy with k8s < 1.14",
		)
	}

	p.OrchestratorProfile.OrchestratorRelease = "1.18"
	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error on azure-policy with k8s >= 1.14",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = to.BoolPtr(true)
	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error on azure-policy with managed identity",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "cluster-autoscaler",
				Enabled: to.BoolPtr(true),
				Pools: []AddonNodePoolsConfig{
					{
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "5",
						},
					},
				},
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			AvailabilityProfile: VirtualMachineScaleSets,
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"cluster-autoscaler addon pools configuration must have a 'name' property that correlates with a pool name in the agentPoolProfiles array",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "cluster-autoscaler",
				Enabled: to.BoolPtr(true),
				Pools: []AddonNodePoolsConfig{
					{
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "5",
						},
					},
				},
				Mode: "foo",
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			AvailabilityProfile: VirtualMachineScaleSets,
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"addon cluster-autoscaler has a mode configuration 'foo', must be either EnsureExists or Reconcile",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "cluster-autoscaler",
				Enabled: to.BoolPtr(true),
				Pools: []AddonNodePoolsConfig{
					{
						Name: "foo",
						Config: map[string]string{
							"min-nodes": "baz",
							"max-nodes": "5",
						},
					},
				},
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			AvailabilityProfile: VirtualMachineScaleSets,
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"cluster-autoscaler addon pool 'name' foo has invalid 'min-nodes' config, must be a string int, got baz",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "cluster-autoscaler",
				Enabled: to.BoolPtr(true),
				Pools: []AddonNodePoolsConfig{
					{
						Name: "foo",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "baz",
						},
					},
				},
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			AvailabilityProfile: VirtualMachineScaleSets,
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"cluster-autoscaler addon pool 'name' foo has invalid 'max-nodes' config, must be a string int, got baz",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "cluster-autoscaler",
				Enabled: to.BoolPtr(true),
				Pools: []AddonNodePoolsConfig{
					{
						Name: "foo",
						Config: map[string]string{
							"min-nodes": "5",
							"max-nodes": "1",
						},
					},
				},
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			AvailabilityProfile: VirtualMachineScaleSets,
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"cluster-autoscaler addon pool 'name' foo has invalid config, 'max-nodes' 1 must be greater than 'min-nodes' 5",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "cluster-autoscaler",
				Enabled: to.BoolPtr(true),
				Pools: []AddonNodePoolsConfig{
					{
						Name: "foo",
						Config: map[string]string{
							"min-nodes": "1",
							"max-nodes": "5",
						},
					},
				},
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			Name:                "bar",
			AvailabilityProfile: VirtualMachineScaleSets,
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"cluster-autoscaler addon pool 'name' foo does not match any agentPoolProfiles nodepool name",
		)
	}

	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			VMSize: "Standard_NC6",
		},
	}
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "nvidia-device-plugin",
				Enabled: to.BoolPtr(true),
			},
		},
	}
	p.OrchestratorProfile.OrchestratorRelease = "1.9"
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error on nvidia-device-plugin with k8s < 1.10",
		)
	}

	p.OrchestratorProfile.OrchestratorRelease = "1.16"
	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error on nvidia-device-plugin with k8s >= 1.12",
		)
	}
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name: common.KubeProxyAddonName,
				Data: "asdasdasdasda",
				Config: map[string]string{
					"foo": "bar",
				},
			},
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"expected error for non-empty Config with non-empty Data",
		)
	}
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name: common.KubeProxyAddonName,
				Data: "asdasdasdasda",
				Containers: []KubernetesContainerSpec{
					{
						Name: "FooContainerSpec",
					},
				},
			},
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"expected error for non-empty Containers with non-empty Data",
		)
	}
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name: common.KubeProxyAddonName,
				Data: "foodata",
			},
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"expected error for invalid base64",
		)
	}
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name: common.KubeProxyAddonName,
				Data: "Zm9vZGF0YQ==",
			},
		},
	}
	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error on providing valid addon.Data",
		)
	}

	// appgw-ingress add-on

	// Basic test with UseManagedIdentity
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		NetworkPlugin:      "azure",
		UseManagedIdentity: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "appgw-ingress",
				Enabled: to.BoolPtr(true),
				Config: map[string]string{
					"appgw-subnet": "10.0.0.0/16",
				},
			},
		},
	}

	if err := p.validateAddons(); err != nil {
		t.Error(
			"should not error for correct config.",
			err,
		)
	}

	// Basic test with ObjectID
	p.ServicePrincipalProfile = &ServicePrincipalProfile{
		ObjectID: "random",
	}
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		NetworkPlugin: "azure",
		Addons: []KubernetesAddon{
			{
				Name:    "appgw-ingress",
				Enabled: to.BoolPtr(true),
				Config: map[string]string{
					"appgw-subnet": "10.0.0.0/16",
				},
			},
		},
	}

	if err := p.validateAddons(); err != nil {
		t.Error(
			"should not error for correct config.",
			err,
		)
	}

	// Test with missing objectID and UseManagedIdentity false
	p.ServicePrincipalProfile = &ServicePrincipalProfile{}
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		NetworkPlugin: "azure",
		Addons: []KubernetesAddon{
			{
				Name:    "appgw-ingress",
				Enabled: to.BoolPtr(true),
				Config: map[string]string{
					"appgw-subnet": "10.0.0.0/16",
				},
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Error(
			"should error as objectID not provided or UseManagedIdentity not true",
			err,
		)
	}

	// Test with wrong Network Plugin
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		NetworkPlugin: "kubelet",
		Addons: []KubernetesAddon{
			{
				Name:    "appgw-ingress",
				Enabled: to.BoolPtr(true),
				Config: map[string]string{
					"appgw-subnet": "10.0.0.0/16",
				},
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error using when not using 'azure' for Network Plugin",
		)
	}

	// Test with missing appgw-subnet
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		NetworkPlugin: "azure",
		Addons: []KubernetesAddon{
			{
				Name:    "appgw-ingress",
				Enabled: to.BoolPtr(true),
				Config:  map[string]string{},
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when missing the subnet for Application Gateway",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.13.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(false),
		Addons: []KubernetesAddon{
			{
				Name:    "azuredisk-csi-driver",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when useCloudControllerManager is disabled for azuredisk-csi-driver",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.13.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "azuredisk-csi-driver",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error when useCloudControllerManager is enabled and k8s version is >= 1.13 for azuredisk-csi-driver",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.13.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(false),
		Addons: []KubernetesAddon{
			{
				Name:    "azurefile-csi-driver",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when useCloudControllerManager is disabled for azurefile-csi-driver",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.13.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "azurefile-csi-driver",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error when useCloudControllerManager is enabled and k8s version is >= 1.13 for azurefile-csi-driver",
		)
	}

	// Basic tests for cloud-node-manager
	p.OrchestratorProfile.OrchestratorVersion = "1.15.4"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "cloud-node-manager",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when the orchestrator version is less than 1.16.0 for cloud-node-manager",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.16.1"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(false),
		Addons: []KubernetesAddon{
			{
				Name:    "cloud-node-manager",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when useCloudControllerManager is disabled for cloud-node-manager",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.17.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "cloud-node-manager",
				Enabled: to.BoolPtr(false),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when useCloudControllerManager is enabled and cloud-node-manager isn't",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.17.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "cloud-node-manager",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error when useCloudControllerManager is enabled and k8s version is >= 1.16 for cloud-node-manager",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.17.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "cloud-node-manager",
				Enabled: to.BoolPtr(true),
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			OSType: Windows,
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when using Windows cluster and k8s version is < 1.18 for cloud-node-manager",
		)
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.18.0"
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		UseCloudControllerManager: to.BoolPtr(true),
		Addons: []KubernetesAddon{
			{
				Name:    "cloud-node-manager",
				Enabled: to.BoolPtr(false),
			},
		},
	}
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			OSType: Windows,
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when using Windows cluster and k8s version is >= 1.18 with cloud-node-manager disabled",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    common.KubeDNSAddonName,
				Enabled: to.BoolPtr(true),
			},
			{
				Name:    common.CoreDNSAddonName,
				Enabled: to.BoolPtr(true),
			},
		},
	}
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error when both kube-dns and coredns are enabled",
		)
	}
}

func TestWindowsVersions(t *testing.T) {
	for _, version := range common.GetAllSupportedKubernetesVersions(false, true, false) {
		cs := getK8sDefaultContainerService(true)
		cs.Properties.OrchestratorProfile.OrchestratorVersion = version
		if err := cs.Validate(false); err != nil {
			t.Errorf(
				"should not error on valid Windows version: %v", err,
			)
		}
		cs = getK8sDefaultContainerService(true)
		cs.Properties.WindowsProfile.AdminPassword = "Password"
		if err := cs.Validate(false); err == nil {
			t.Errorf(
				"should error on windows password complexity not match because no digits and special characters found in the password ",
			)
		}
		cs = getK8sDefaultContainerService(true)
		cs.Properties.WindowsProfile.AdminPassword = "123!@#"
		if err := cs.Validate(false); err == nil {
			t.Errorf(
				"should error on windows password complexity not match because uppercase and lowercase letters found in the password",
			)
		}
		cs = getK8sDefaultContainerService(true)
		cs.Properties.WindowsProfile.AdminPassword = ""
		if err := cs.Validate(false); err == nil {
			t.Errorf(
				"should error on windows password length is zero",
			)
		}
		cs = getK8sDefaultContainerService(true)
		cs.Properties.WindowsProfile.AdminUsername = "User@123"
		cs.Properties.WindowsProfile.AdminPassword = "User@123"
		if err := cs.Validate(false); err == nil {
			t.Errorf(
				"should error on windows password complexity not match because username and password are  same",
			)
		}
		sv, _ := semver.Make(version)
		cs = getK8sDefaultContainerService(true)
		cs.Properties.OrchestratorProfile.OrchestratorRelease = fmt.Sprintf("%d.%d", sv.Major, sv.Minor)
		if err := cs.Validate(false); err != nil {
			t.Errorf(
				"should not error on valid Windows version: %v", err,
			)
		}
	}
	cs := getK8sDefaultContainerService(true)
	cs.Properties.OrchestratorProfile.OrchestratorRelease = "1.4"
	if err := cs.Validate(false); err == nil {
		t.Errorf(
			"should error on invalid Windows version",
		)
	}

	cs = getK8sDefaultContainerService(true)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.4.0"
	if err := cs.Validate(false); err == nil {
		t.Errorf(
			"should error on invalid Windows version",
		)
	}
}

func TestLinuxVersions(t *testing.T) {
	for _, version := range common.GetAllSupportedKubernetesVersions(false, false, false) {
		cs := getK8sDefaultContainerService(false)
		cs.Properties.OrchestratorProfile.OrchestratorVersion = version
		if err := cs.Validate(false); err != nil {
			t.Errorf(
				"should not error on valid Linux version: %v", err,
			)
		}
		sv, _ := semver.Make(version)
		cs = getK8sDefaultContainerService(false)
		cs.Properties.OrchestratorProfile.OrchestratorRelease = fmt.Sprintf("%d.%d", sv.Major, sv.Minor)
		if err := cs.Validate(false); err != nil {
			t.Errorf(
				"should not error on valid Linux version: %v", err,
			)
		}
	}
	cs := getK8sDefaultContainerService(false)
	cs.Properties.OrchestratorProfile.OrchestratorRelease = "1.4"
	if err := cs.Validate(false); err == nil {
		t.Errorf(
			"should error on invalid Linux version",
		)
	}

	cs = getK8sDefaultContainerService(false)
	cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.4.0"
	if err := cs.Validate(false); err == nil {
		t.Errorf(
			"should error on invalid Linux version",
		)
	}
}

func TestValidateImageNameAndGroup(t *testing.T) {
	tests := []struct {
		name        string
		image       ImageReference
		expectedErr error
	}{
		{
			name: "valid run",
			image: ImageReference{
				Name:          "rhel9000",
				ResourceGroup: "club",
			},
			expectedErr: nil,
		},
		{
			name: "invalid: image name is missing",
			image: ImageReference{
				ResourceGroup: "club",
			},
			expectedErr: errors.New(`imageName needs to be specified when imageResourceGroup is provided`),
		},
		{
			name: "invalid: image resource group is missing",
			image: ImageReference{
				Name: "rhel9000",
			},
			expectedErr: errors.New(`imageResourceGroup needs to be specified when imageName is provided`),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.AgentPoolProfiles = []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					ImageRef:            &test.image,
				},
			}
			gotErr := cs.Properties.validateAgentPoolProfiles(true)
			if !helpers.EqualError(gotErr, test.expectedErr) {
				t.Logf("scenario %q", test.name)
				t.Errorf("expected error: %v, got: %v", test.expectedErr, gotErr)
			}
		})
	}
}
func TestProperties_ValidateManagedIdentity(t *testing.T) {
	tests := []struct {
		name                string
		orchestratorRelease string
		useManagedIdentity  bool
		userAssignedID      string
		masterProfile       MasterProfile
		expectedErr         string
		agentPoolProfiles   []*AgentPoolProfile
	}{
		{
			name:                "use managed identity with master vmas",
			orchestratorRelease: "1.16",
			useManagedIdentity:  true,
			masterProfile: MasterProfile{
				DNSPrefix: "dummy",
				Count:     3,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:   "agentpool",
					VMSize: "Standard_DS2_v2",
					Count:  1,
				},
			},
		},
		{
			name:                "use master VMSS with empty user assigned ID",
			orchestratorRelease: "1.16",
			useManagedIdentity:  true,
			masterProfile: MasterProfile{
				DNSPrefix:           "dummy",
				Count:               3,
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               1,
					AvailabilityProfile: VirtualMachineScaleSets,
				},
			},
			expectedErr: "virtualMachineScaleSets for master profile can be used only with user assigned MSI ! Please specify \"userAssignedID\" in \"kubernetesConfig\"",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.MasterProfile = &test.masterProfile
			cs.Properties.MasterProfile.VMSize = "Standard_DS2_v2"
			cs.Properties.OrchestratorProfile = &OrchestratorProfile{
				OrchestratorRelease: test.orchestratorRelease,
				OrchestratorType:    Kubernetes,
				KubernetesConfig: &KubernetesConfig{
					UseManagedIdentity: to.BoolPtr(test.useManagedIdentity),
					UserAssignedID:     test.userAssignedID,
				},
			}
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Validate(false)
			if test.expectedErr == "" && err != nil ||
				test.expectedErr != "" && (err == nil || test.expectedErr != err.Error()) {
				t.Errorf("test %s: unexpected error %q\n", test.name, err)
			}
		})
	}
}

func TestMasterProfileValidate(t *testing.T) {
	tests := []struct {
		name                string
		orchestratorType    string
		orchestratorVersion string
		orchestratorRelease string
		useInstanceMetadata bool
		masterProfile       MasterProfile
		agentPoolProfiles   []*AgentPoolProfile
		expectedErr         string
	}{
		{
			name: "Master Profile with Invalid DNS Prefix",
			masterProfile: MasterProfile{
				DNSPrefix: "bad!",
			},
			expectedErr: "DNSPrefix 'bad!' is invalid. The DNSPrefix must contain between 3 and 45 characters and can contain only letters, numbers, and hyphens.  It must start with a letter and must end with a letter or a number. (length was 4)",
		},
		{
			name: "Master Profile with valid DNS Prefix 1",
			masterProfile: MasterProfile{
				DNSPrefix: "dummy",
				Count:     1,
			},
		},
		{
			name: "Master Profile with valid DNS Prefix 2",
			masterProfile: MasterProfile{
				DNSPrefix: "dummy",
				Count:     3,
			},
		},
		{
			name:             "Master Profile with empty imageName and non-empty imageResourceGroup",
			orchestratorType: Kubernetes,
			masterProfile: MasterProfile{
				DNSPrefix: "dummy",
				Count:     3,
				ImageRef: &ImageReference{
					Name:          "",
					ResourceGroup: "rg",
				},
			},
			expectedErr: "imageName needs to be specified when imageResourceGroup is provided",
		},
		{
			name:                "Master Profile with VMSS and storage account",
			orchestratorType:    Kubernetes,
			orchestratorRelease: "1.16",
			masterProfile: MasterProfile{
				DNSPrefix:           "dummy",
				Count:               3,
				AvailabilityProfile: VirtualMachineScaleSets,
				StorageProfile:      StorageAccount,
			},
			expectedErr: "VirtualMachineScaleSets does not support StorageAccount disks.  Please specify \"storageProfile\": \"ManagedDisks\" (recommended) or \"availabilityProfile\": \"AvailabilitySet\"",
		},
		{
			name:                "Master Profile with VMSS and agent profiles with VMAS",
			orchestratorType:    Kubernetes,
			orchestratorRelease: "1.16",
			masterProfile: MasterProfile{
				DNSPrefix:           "dummy",
				Count:               3,
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
				},
			},
			expectedErr: "VirtualMachineScaleSets for master profile must be used together with virtualMachineScaleSets for agent profiles. Set \"availabilityProfile\" to \"VirtualMachineScaleSets\" for agent profiles",
		},
		{
			name: "Master Profile with valid OSDiskCachingType: None",
			masterProfile: MasterProfile{
				DNSPrefix:         "foo",
				Count:             3,
				OSDiskCachingType: string(compute.CachingTypesNone),
			},
		},
		{
			name: "Master Profile with valid OSDiskCachingType: ReadWrite",
			masterProfile: MasterProfile{
				DNSPrefix:         "bar",
				Count:             3,
				OSDiskCachingType: string(compute.CachingTypesReadWrite),
			},
		},
		{
			name: "Master Profile with valid OSDiskCachingType: ReadOnly",
			masterProfile: MasterProfile{
				DNSPrefix:         "baz",
				Count:             3,
				OSDiskCachingType: string(compute.CachingTypesReadOnly),
			},
		},
		{
			name: "Master Profile with invalid OSDiskCachingType",
			masterProfile: MasterProfile{
				DNSPrefix:         "whizbang",
				Count:             3,
				OSDiskCachingType: "NotExist",
			},
			expectedErr: fmt.Sprintf("Invalid masterProfile osDiskCachingType value \"%s\", please use one of the following versions: %s", "NotExist", cachingTypesValidValues),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.MasterProfile = &test.masterProfile
			cs.Properties.MasterProfile.StorageProfile = test.masterProfile.StorageProfile
			cs.Properties.OrchestratorProfile = &OrchestratorProfile{
				OrchestratorType:    test.orchestratorType,
				OrchestratorVersion: test.orchestratorVersion,
				OrchestratorRelease: test.orchestratorRelease,
				KubernetesConfig: &KubernetesConfig{
					UseInstanceMetadata: to.BoolPtr(test.useInstanceMetadata),
				},
			}
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Properties.validateMasterProfile(false)
			if test.expectedErr == "" && err != nil ||
				test.expectedErr != "" && (err == nil || test.expectedErr != err.Error()) {
				t.Errorf("test %s: unexpected error %q\n", test.name, err)
			}
		})
	}
}

func TestProperties_ValidateZones(t *testing.T) {
	tests := []struct {
		name                        string
		orchestratorRelease         string
		loadBalancerSku             string
		excludeMasterFromStandardLB bool
		masterProfile               *MasterProfile
		agentProfiles               []*AgentPoolProfile
		expectedErr                 bool
		expectedErrStr              string
	}{
		{
			name:                "Agent profile with zones vmas",
			orchestratorRelease: "1.16",
			masterProfile: &MasterProfile{
				Count:               5,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
				AvailabilityZones:   []string{"1", "2"},
			},
			agentProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               4,
					AvailabilityProfile: AvailabilitySet,
					AvailabilityZones:   []string{"1", "2"},
				},
			},
			expectedErr:    true,
			expectedErrStr: "VirtualMachineScaleSets for master profile must be used together with virtualMachineScaleSets for agent profiles. Set \"availabilityProfile\" to \"VirtualMachineScaleSets\" for agent profiles",
		},
		{
			name:                "Master profile with zones and Agent profile without zones",
			orchestratorRelease: "1.16",
			masterProfile: &MasterProfile{
				Count:               5,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
				AvailabilityZones:   []string{"1", "2"},
			},
			agentProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               4,
					AvailabilityProfile: VirtualMachineScaleSets,
				},
			},
			expectedErr:    false,
			expectedErrStr: "Availability Zones need to be defined for master profile and all agent pool profiles. Please set \"availabilityZones\" for all profiles",
		},
		{
			name:                "Master profile without zones and Agent profile with zones",
			orchestratorRelease: "1.16",
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			agentProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               4,
					AvailabilityProfile: VirtualMachineScaleSets,
					AvailabilityZones:   []string{"1", "2"},
				},
			},
			expectedErr:    false,
			expectedErrStr: "Availability Zones need to be defined for master profile and all agent pool profiles. Please set \"availabilityZones\" for all profiles",
		},
		{
			name:                "all zones and basic loadbalancer",
			orchestratorRelease: "1.16",
			loadBalancerSku:     BasicLoadBalancerSku,
			masterProfile: &MasterProfile{
				Count:               5,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
				AvailabilityZones:   []string{"1", "2"},
			},
			agentProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               4,
					AvailabilityProfile: VirtualMachineScaleSets,
					AvailabilityZones:   []string{"1", "2"},
				},
			},
			expectedErr:    true,
			expectedErrStr: "Availability Zones requires Standard LoadBalancer. Please set KubernetesConfig \"LoadBalancerSku\" to \"Standard\"",
		},
		{
			name:                        "all zones with standard loadbalancer and false excludeMasterFromStandardLB",
			orchestratorRelease:         "1.16",
			loadBalancerSku:             StandardLoadBalancerSku,
			excludeMasterFromStandardLB: false,
			masterProfile: &MasterProfile{
				Count:               5,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
				AvailabilityZones:   []string{"1", "2"},
			},
			agentProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               4,
					AvailabilityProfile: VirtualMachineScaleSets,
					AvailabilityZones:   []string{"1", "2"},
				},
			},
			expectedErr:    true,
			expectedErrStr: "standard loadBalancerSku should exclude master nodes. Please set KubernetesConfig \"ExcludeMasterFromStandardLB\" to \"true\"",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.MasterProfile = test.masterProfile
			cs.Properties.AgentPoolProfiles = test.agentProfiles
			cs.Properties.OrchestratorProfile.OrchestratorRelease = test.orchestratorRelease
			cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
				LoadBalancerSku:             test.loadBalancerSku,
				ExcludeMasterFromStandardLB: to.BoolPtr(test.excludeMasterFromStandardLB),
			}

			err := cs.Validate(false)
			if test.expectedErr {
				if err == nil {
					t.Errorf("error should have occurred")
				} else {
					if err.Error() != test.expectedErrStr {
						t.Errorf("expected error with message : %s, but got : %s", test.expectedErrStr, err.Error())
					}
				}
			}
		})
	}
}

func ExampleProperties_validateLocation() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	cs := getK8sDefaultContainerService(true)
	cs.Location = ""
	if err := cs.validateLocation(); err != nil {
		fmt.Printf("error in validateLocation: %s", err)
	}
	// Output:
	// level=warning msg="No \"location\" value was specified, AKS Engine will generate an ARM template configuration valid for regions in public cloud only"
}

func ExampleProperties_validateMasterProfile() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	cs := getK8sDefaultContainerService(false)
	cs.Properties.MasterProfile.Count = 1
	cs.Properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
	cs.Properties.AgentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
	if err := cs.Properties.validateMasterProfile(false); err != nil {
		log.Errorf("shouldn't error with 1 control plane VM, got %s", err.Error())
	}

	cs = getK8sDefaultContainerService(false)
	cs.Properties.MasterProfile.Count = 1
	if err := cs.Properties.validateMasterProfile(true); err != nil {
		log.Errorf("shouldn't error with 1 control plane VM, got %s", err.Error())
	}
	// Output:
	// level=warning msg="Running only 1 control plane VM not recommended for production clusters, use 3 or 5 for control plane redundancy"
	// level=warning msg="Clusters with a VMSS control plane are not upgradable! You will not be able to upgrade your cluster using `aks-engine upgrade`"
}

func ExampleProperties_validateZones() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	cs := getK8sDefaultContainerService(true)

	// Master VMs have Availability Zone configuration, but pool does not
	cs.Properties.MasterProfile = &MasterProfile{
		Count:               5,
		DNSPrefix:           "foo",
		VMSize:              "Standard_DS2_v2",
		AvailabilityProfile: VirtualMachineScaleSets,
		AvailabilityZones:   []string{"1", "2"},
	}
	cs.Properties.AgentPoolProfiles = []*AgentPoolProfile{
		{
			Name:                "agentpool",
			VMSize:              "Standard_DS2_v2",
			Count:               4,
			AvailabilityProfile: AvailabilitySet,
		},
	}
	if err := cs.Properties.validateZones(); err != nil {
		log.Error(err)
	}
	// Should yield:
	// level=warning msg="This cluster is using Availability Zones for master VMs, but not for pool \"agentpool\""

	// Pool has Availability Zone configuration, but Master VMs do not
	cs.Properties.MasterProfile = &MasterProfile{
		Count:               5,
		DNSPrefix:           "foo",
		VMSize:              "Standard_DS2_v2",
		AvailabilityProfile: VirtualMachineScaleSets,
	}
	cs.Properties.AgentPoolProfiles = []*AgentPoolProfile{
		{
			Name:                "anotherpool",
			VMSize:              "Standard_DS2_v2",
			Count:               4,
			AvailabilityProfile: AvailabilitySet,
			AvailabilityZones:   []string{"1", "2"},
		},
	}
	if err := cs.Properties.validateZones(); err != nil {
		log.Error(err)
	}
	// Should yield:
	// level=warning msg="This cluster is using Availability Zones for pool \"anotherpool\", but not for master VMs"

	// Some pools have Availability Zone configuration, Master VMs do not
	cs.Properties.AgentPoolProfiles = []*AgentPoolProfile{
		{
			Name:                "anotherpool",
			VMSize:              "Standard_DS2_v2",
			Count:               4,
			AvailabilityProfile: AvailabilitySet,
		},
		{
			Name:                "anotherpool2",
			VMSize:              "Standard_DS2_v2",
			Count:               4,
			AvailabilityProfile: AvailabilitySet,
			AvailabilityZones:   []string{"1", "2"},
		},
		{
			Name:                "anotherpool3",
			VMSize:              "Standard_DS2_v2",
			Count:               4,
			AvailabilityProfile: AvailabilitySet,
		},
		{
			Name:                "anotherpool4",
			VMSize:              "Standard_DS2_v2",
			Count:               4,
			AvailabilityProfile: AvailabilitySet,
			AvailabilityZones:   []string{"1", "2"},
		},
	}
	if err := cs.Properties.validateZones(); err != nil {
		log.Error(err)
	}
	// Should yield:
	// level=warning msg="This cluster is using Availability Zones for pools \"anotherpool2\" and \"anotherpool4\", but not for pools \"anotherpool\" and \"anotherpool3\", nor for master VMs"

	// Master VMs and some (but not all) pools have Availability Zone configuration
	cs.Properties.MasterProfile = &MasterProfile{
		Count:               5,
		DNSPrefix:           "foo",
		VMSize:              "Standard_DS2_v2",
		AvailabilityProfile: VirtualMachineScaleSets,
		AvailabilityZones:   []string{"1", "2"},
	}
	if err := cs.Properties.validateZones(); err != nil {
		log.Error(err)
	}
	// Should yield:
	// level=warning msg="This cluster is using Availability Zones for master VMs, but not for pools \"anotherpool\" and \"anotherpool3\""
	// The ordered collection of all output is validated below:

	// Output:
	// level=warning msg="This cluster is using Availability Zones for master VMs, but not for pool \"agentpool\""
	// level=warning msg="This cluster is using Availability Zones for pool \"anotherpool\", but not for master VMs"
	// level=warning msg="This cluster is using Availability Zones for pools \"anotherpool2\" and \"anotherpool4\", but not for pools \"anotherpool\" and \"anotherpool3\", nor for master VMs"
	// level=warning msg="This cluster is using Availability Zones for master VMs, but not for pools \"anotherpool\" and \"anotherpool3\""
}

func TestProperties_ValidateLoadBalancer(t *testing.T) {
	tests := []struct {
		name                        string
		orchestratorRelease         string
		loadBalancerSku             string
		masterProfile               *MasterProfile
		agentProfiles               []*AgentPoolProfile
		excludeMasterFromStandardLB bool
		expectedErr                 bool
		expectedErrStr              string
	}{
		{
			name:                "lowercase basic LB",
			orchestratorRelease: "1.16",
			loadBalancerSku:     "basic",
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			agentProfiles: []*AgentPoolProfile{
				{
					Name:                   "agentpool",
					VMSize:                 "Standard_DS2_v2",
					Count:                  4,
					AvailabilityProfile:    VirtualMachineScaleSets,
					EnableVMSSNodePublicIP: to.BoolPtr(true),
				},
			},
		},
		{
			name:                "Basic LB",
			orchestratorRelease: "1.16",
			loadBalancerSku:     BasicLoadBalancerSku,
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
		{
			name:                "lowercase standard LB",
			orchestratorRelease: "1.16",
			loadBalancerSku:     "standard",
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
		{
			name:                "Standard LB without master excluded",
			orchestratorRelease: "1.16",
			loadBalancerSku:     StandardLoadBalancerSku,
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			expectedErr:    true,
			expectedErrStr: "standard loadBalancerSku should exclude master nodes. Please set KubernetesConfig \"ExcludeMasterFromStandardLB\" to \"true\"",
		},
		{
			name:                "Standard LB",
			orchestratorRelease: "1.16",
			loadBalancerSku:     StandardLoadBalancerSku,
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			excludeMasterFromStandardLB: true,
		},
		{
			name:                "empty string LB value",
			orchestratorRelease: "1.16",
			loadBalancerSku:     "",
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
		{
			name:                "invalid LB string value",
			orchestratorRelease: "1.16",
			loadBalancerSku:     "foo",
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			expectedErr:    true,
			expectedErrStr: fmt.Sprintf("Invalid value for loadBalancerSku, only %s and %s are supported", StandardLoadBalancerSku, BasicLoadBalancerSku),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.MasterProfile = test.masterProfile
			cs.Properties.AgentPoolProfiles = test.agentProfiles
			cs.Properties.OrchestratorProfile.OrchestratorRelease = test.orchestratorRelease
			cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
				LoadBalancerSku:             test.loadBalancerSku,
				ExcludeMasterFromStandardLB: to.BoolPtr(test.excludeMasterFromStandardLB),
			}

			err := cs.Validate(false)
			if test.expectedErr {
				if err == nil {
					t.Errorf("error should have occurred")
				} else {
					if err.Error() != test.expectedErrStr {
						t.Errorf("expected error with message : %s, but got : %s", test.expectedErrStr, err.Error())
					}
				}
			} else if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestProperties_ValidateSinglePlacementGroup(t *testing.T) {

	tests := []struct {
		name              string
		masterProfile     *MasterProfile
		agentPoolProfiles []*AgentPoolProfile
		expectedMsg       string
	}{
		{
			name: "Master profile VMAS with SinglePlacementGroup",
			masterProfile: &MasterProfile{
				Count:                1,
				DNSPrefix:            "foo",
				VMSize:               "Standard_DS2_v2",
				AvailabilityProfile:  AvailabilitySet,
				SinglePlacementGroup: to.BoolPtr(false),
			},
			expectedMsg: "singlePlacementGroup is only supported with VirtualMachineScaleSets",
		},
		{
			name: "Agent profile VMAS with SinglePlacementGroup",
			masterProfile: &MasterProfile{
				Count:               1,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                 "agentpool",
					VMSize:               "Standard_DS2_v2",
					Count:                4,
					AvailabilityProfile:  AvailabilitySet,
					SinglePlacementGroup: to.BoolPtr(false),
				},
			},
			expectedMsg: `VirtualMachineScaleSets for master profile must be used together with virtualMachineScaleSets for agent profiles. Set "availabilityProfile" to "VirtualMachineScaleSets" for agent profiles`,
		},
		{
			name: "VMSS with SinglePlacementGroup false and StorageAccount storage",
			masterProfile: &MasterProfile{
				Count:                1,
				DNSPrefix:            "foo",
				VMSize:               "Standard_DS2_v2",
				AvailabilityProfile:  VirtualMachineScaleSets,
				SinglePlacementGroup: to.BoolPtr(false),
				StorageProfile:       StorageAccount,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               4,
					AvailabilityProfile: VirtualMachineScaleSets,
				},
			},
			expectedMsg: "VirtualMachineScaleSets does not support StorageAccount disks.  Please specify \"storageProfile\": \"ManagedDisks\" (recommended) or \"availabilityProfile\": \"AvailabilitySet\"",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.OrchestratorProfile.OrchestratorRelease = "1.16"
			cs.Properties.MasterProfile = test.masterProfile
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Validate(true)
			if err.Error() != test.expectedMsg {
				t.Errorf("expected error message : %s, but got %s", test.expectedMsg, err.Error())
			}
		})
	}
}

func TestProperties_ValidatePPGID(t *testing.T) {

	tests := []struct {
		name              string
		masterProfile     *MasterProfile
		agentPoolProfiles []*AgentPoolProfile
		expectedMsg       string
	}{
		{
			name: "Master profile VMAs with faulty PPG",
			masterProfile: &MasterProfile{
				Count:                     1,
				DNSPrefix:                 "foo",
				VMSize:                    "Standard_DS2_v2",
				AvailabilityProfile:       AvailabilitySet,
				ProximityPlacementGroupID: "faultyPPG",
			},
			expectedMsg: `ProximityPlacementGroupID(faultyPPG) is of incorrect format, correct format: ^/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.Compute/proximityPlacementGroups/[^/\s]+$`,
		},
		{
			name: "Agent profile VMSS with faulty PPG",
			masterProfile: &MasterProfile{
				Count:               1,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: AvailabilitySet,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                      "agentpool",
					VMSize:                    "Standard_DS2_v2",
					Count:                     4,
					AvailabilityProfile:       VirtualMachineScaleSets,
					ProximityPlacementGroupID: "faultyPPG",
				},
			},
			expectedMsg: `ProximityPlacementGroupID(faultyPPG) is of incorrect format, correct format: ^/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.Compute/proximityPlacementGroups/[^/\s]+$`,
		},
		{
			name: "Faulty PPG",
			masterProfile: &MasterProfile{
				Count:               1,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: AvailabilitySet,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                      "agentpool",
					VMSize:                    "Standard_DS2_v2",
					Count:                     4,
					AvailabilityProfile:       VirtualMachineScaleSets,
					ProximityPlacementGroupID: "/subscriptions/11111111-0000-1111-0000-111111111111/resourceGroups/test-nodepool-ppg-rg/providers/Microsoft.Compute/proximityPlacementGroups",
				},
			},
			expectedMsg: `ProximityPlacementGroupID(/subscriptions/11111111-0000-1111-0000-111111111111/resourceGroups/test-nodepool-ppg-rg/providers/Microsoft.Compute/proximityPlacementGroups) is of incorrect format, correct format: ^/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.Compute/proximityPlacementGroups/[^/\s]+$`,
		},
		{
			name: "Correct PPGs in both master and nodepool",
			masterProfile: &MasterProfile{
				Count:                     1,
				DNSPrefix:                 "foo",
				VMSize:                    "Standard_DS2_v2",
				AvailabilityProfile:       AvailabilitySet,
				ProximityPlacementGroupID: "/subscriptions/11111111-0000-1111-0000-111111111111/resourceGroups/test-master-ppg-rg/providers/Microsoft.Compute/proximityPlacementGroups/test-master-ppg",
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                      "agentpool",
					VMSize:                    "Standard_DS2_v2",
					Count:                     4,
					AvailabilityProfile:       VirtualMachineScaleSets,
					ProximityPlacementGroupID: "/subscriptions/11111111-0000-1111-0000-111111111111/resourceGroups/test-nodepool-ppg-rg/providers/Microsoft.Compute/proximityPlacementGroups/test-nodepool-ppg",
				},
			},
			expectedMsg: ``,
		},
		{
			name: "Without PPG settings",
			masterProfile: &MasterProfile{
				Count:               1,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: AvailabilitySet,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_DS2_v2",
					Count:               4,
					AvailabilityProfile: VirtualMachineScaleSets,
				},
			},
			expectedMsg: ``,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.OrchestratorProfile.OrchestratorRelease = "1.16"
			cs.Properties.MasterProfile = test.masterProfile
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Validate(true)
			if err != nil {
				if err.Error() != test.expectedMsg {
					t.Errorf("expected error message : %s, but got %s", test.expectedMsg, err.Error())
				}
			}
		})
	}
}

func TestProperties_ValidateVNET(t *testing.T) {
	validVNetSubnetID := "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/SUBNET_NAME"
	validVNetSubnetID2 := "/subscriptions/SUB_ID2/resourceGroups/RG_NAME2/providers/Microsoft.Network/virtualNetworks/VNET_NAME2/subnets/SUBNET_NAME"

	tests := []struct {
		name              string
		masterProfile     *MasterProfile
		agentPoolProfiles []*AgentPoolProfile
		expectedMsg       string
	}{
		{
			name: "Multiple VNET Subnet configs",
			masterProfile: &MasterProfile{
				VnetSubnetID: "testvnetstring",
				Count:        1,
				DNSPrefix:    "foo",
				VMSize:       "Standard_DS2_v2",
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					VnetSubnetID:        "",
				},
			},
			expectedMsg: "Multiple VNET Subnet configurations specified.  The master profile and each agent pool profile must all specify a custom VNET Subnet, or none at all",
		},
		{
			name: "Invalid vnet subnet ID",
			masterProfile: &MasterProfile{
				VnetSubnetID: "testvnetstring",
				Count:        1,
				DNSPrefix:    "foo",
				VMSize:       "Standard_DS2_v2",
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					VnetSubnetID:        "testvnetstring",
				},
			},
			expectedMsg: "Unable to parse vnetSubnetID. Please use a vnetSubnetID with format /subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/SUBNET_NAME",
		},
		{
			name: "Multiple VNETs",
			masterProfile: &MasterProfile{
				VnetSubnetID: validVNetSubnetID,
				Count:        1,
				DNSPrefix:    "foo",
				VMSize:       "Standard_DS2_v2",
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					VnetSubnetID:        validVNetSubnetID,
				},
				{
					Name:                "agentpool2",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					VnetSubnetID:        validVNetSubnetID2,
				},
			},
			expectedMsg: "Multiple VNETS specified.  The master profile and each agent pool must reference the same VNET (but it is ok to reference different subnets on that VNET)",
		},
		{
			name: "Invalid MasterProfile FirstConsecutiveStaticIP when master is VMAS",
			masterProfile: &MasterProfile{
				FirstConsecutiveStaticIP: "10.0.0.invalid",
				VnetSubnetID:             validVNetSubnetID,
				Count:                    1,
				DNSPrefix:                "foo",
				VMSize:                   "Standard_DS2_v2",
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					VnetSubnetID:        validVNetSubnetID,
				},
			},
			expectedMsg: "MasterProfile.FirstConsecutiveStaticIP (with VNET Subnet specification) '10.0.0.invalid' is an invalid IP address",
		},
		{
			name: "Empty MasterProfile FirstConsecutiveStaticIP and empty agentVnetSubnetID when master is VMSS",
			masterProfile: &MasterProfile{
				VnetSubnetID:        validVNetSubnetID,
				Count:               1,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: VirtualMachineScaleSets,
					VnetSubnetID:        validVNetSubnetID,
				},
			},
			expectedMsg: "when master profile is using VirtualMachineScaleSets and is custom vnet, set \"vnetsubnetid\" and \"agentVnetSubnetID\" for master profile",
		},
		{
			name: "User-provided MasterProfile FirstConsecutiveStaticIP when master is VMSS",
			masterProfile: &MasterProfile{
				VnetSubnetID:             validVNetSubnetID,
				Count:                    1,
				DNSPrefix:                "foo",
				VMSize:                   "Standard_DS2_v2",
				AvailabilityProfile:      VirtualMachineScaleSets,
				FirstConsecutiveStaticIP: "10.0.0.4",
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: VirtualMachineScaleSets,
					VnetSubnetID:        validVNetSubnetID,
				},
			},
			expectedMsg: "when masterProfile's availabilityProfile is VirtualMachineScaleSets and a vnetSubnetID is specified, the firstConsecutiveStaticIP should be empty and will be determined by an offset from the first IP in the vnetCidr",
		},
		{
			name: "Invalid vnetcidr",
			masterProfile: &MasterProfile{
				FirstConsecutiveStaticIP: "10.0.0.1",
				VnetSubnetID:             validVNetSubnetID,
				Count:                    1,
				DNSPrefix:                "foo",
				VMSize:                   "Standard_DS2_v2",
				VnetCidr:                 "10.1.0.0/invalid",
			},
			agentPoolProfiles: []*AgentPoolProfile{
				{
					Name:                "agentpool",
					VMSize:              "Standard_D2_v2",
					Count:               1,
					AvailabilityProfile: AvailabilitySet,
					VnetSubnetID:        validVNetSubnetID,
				},
			},
			expectedMsg: "MasterProfile.VnetCidr '10.1.0.0/invalid' contains invalid cidr notation",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Properties.OrchestratorProfile.OrchestratorRelease = "1.16"
			cs.Properties.MasterProfile = test.masterProfile
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Validate(true)
			if err.Error() != test.expectedMsg {
				t.Errorf("expected error message : %s, but got %s", test.expectedMsg, err.Error())
			}
		})
	}
}

func TestValidate_VaultKeySecrets(t *testing.T) {

	tests := []struct {
		name        string
		secrets     []KeyVaultSecrets
		expectedErr error
	}{
		{
			name: "Empty Vault Certificates",
			secrets: []KeyVaultSecrets{
				{
					SourceVault: &KeyVaultID{
						ID: "0a0b0c0d0e0f",
					},
					VaultCertificates: []KeyVaultCertificate{},
				},
			},
			expectedErr: errors.New("Valid KeyVaultSecrets must have no empty VaultCertificates"),
		},
		{
			name: "No SourceVault ID",
			secrets: []KeyVaultSecrets{
				{
					SourceVault: &KeyVaultID{},
					VaultCertificates: []KeyVaultCertificate{
						{
							CertificateURL:   "dummyURL",
							CertificateStore: "dummyCertStore",
						},
					},
				},
			},
			expectedErr: errors.New("KeyVaultSecrets must have a SourceVault.ID"),
		},
		{
			name: "Empty SourceVault",
			secrets: []KeyVaultSecrets{
				{
					VaultCertificates: []KeyVaultCertificate{
						{
							CertificateURL:   "dummyURL",
							CertificateStore: "dummyCertStore",
						},
					},
				},
			},
			expectedErr: errors.New("missing SourceVault in KeyVaultSecrets"),
		},
		{
			name: "Empty Certificate Store",
			secrets: []KeyVaultSecrets{
				{
					SourceVault: &KeyVaultID{
						ID: "0a0b0c0d0e0f",
					},
					VaultCertificates: []KeyVaultCertificate{
						{
							CertificateURL:   "dummyUrl",
							CertificateStore: "",
						},
					},
				},
			},
			expectedErr: errors.New("KeyVaultCertificate.CertificateStore must be a non-empty value for certificates in a WindowsProfile"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := validateKeyVaultSecrets(test.secrets, true)
			if err.Error() != test.expectedErr.Error() {
				t.Errorf("expected error to be thrown with msg : %s", test.expectedErr.Error())
			}
		})
	}
}

func TestValidateProperties_OrchestratorSpecificProperties(t *testing.T) {
	t.Run("Should not support DNS prefix for Kubernetes orchestrators", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].DNSPrefix = "sampleprefix"
		expectedMsg := "AgentPoolProfile.DNSPrefix must be empty for Kubernetes"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s", expectedMsg)
		}
	})

	t.Run("Should not contain agentPool ports for Kubernetes orchestrators", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].Ports = []int{80, 443, 8080}
		expectedMsg := "AgentPoolProfile.Ports must be empty for Kubernetes"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should not support ScaleSetEviction policies with regular priority", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].Ports = []int{}
		agentPoolProfiles[0].ScaleSetPriority = "Regular"
		agentPoolProfiles[0].ScaleSetEvictionPolicy = "Deallocate"
		expectedMsg := "property 'AgentPoolProfile.ScaleSetEvictionPolicy' must be empty for AgentPoolProfile.Priority of Regular"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should not support os type other than linux for single stack ipv6 and dual stack feature", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(true)
		for _, featureFlags := range []FeatureFlags{{EnableIPv6DualStack: true}, {EnableIPv6Only: true}} {
			cs.Properties.FeatureFlags = &featureFlags
			masterProfile := cs.Properties.MasterProfile

			masterProfile.Distro = Ubuntu
			agentPoolProfiles := cs.Properties.AgentPoolProfiles
			agentPoolProfiles[0].OSType = Windows
			expectedMsg := fmt.Sprintf("Dual stack and single stack IPv6 feature is supported only with Linux, but agent pool '%s' is of os type %s", agentPoolProfiles[0].Name, agentPoolProfiles[0].OSType)
			if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
				t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
			}

			agentPoolProfiles[0].OSType = Linux
			agentPoolProfiles[0].Distro = Flatcar
			expectedMsg = fmt.Sprintf("Dual stack and single stack IPv6 feature is currently supported only with Ubuntu, but agent pool '%s' is of distro type %s", agentPoolProfiles[0].Name, agentPoolProfiles[0].Distro)
			if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
				t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
			}
		}

	})
}

func TestValidateProperties_CustomNodeLabels(t *testing.T) {

	t.Run("Should throw error for invalid Kubernetes Label Keys", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].CustomNodeLabels = map[string]string{
			"a/b/c": "a",
		}
		expectedMsg := "Label key 'a/b/c' is invalid. Valid label keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/)"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should throw error for invalid Kubernetes Label Values", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].CustomNodeLabels = map[string]string{
			"fookey": "b$$a$$r",
		}
		expectedMsg := "Label value 'b$$a$$r' is invalid. Valid label values must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})
}

func TestAgentPoolProfile_ValidateAvailabilityProfile(t *testing.T) {
	t.Run("Should fail for invalid availability profile", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = "InvalidAvailabilityProfile"
		expectedMsg := "unknown availability profile type 'InvalidAvailabilityProfile' for agent pool 'agentpool'.  Specify either AvailabilitySet, or VirtualMachineScaleSets"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for AvailabilitySet + SinglePlacementGroup true", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].SinglePlacementGroup = to.BoolPtr(true)
		expectedMsg := "singlePlacementGroup is only supported with VirtualMachineScaleSets"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for AvailabilitySet + SinglePlacementGroup false", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].SinglePlacementGroup = to.BoolPtr(false)
		expectedMsg := "singlePlacementGroup is only supported with VirtualMachineScaleSets"
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for AvailabilitySet + invalid LoadBalancerBackendAddressPoolIDs", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = AvailabilitySet
		agentPoolProfiles[0].LoadBalancerBackendAddressPoolIDs = []string{"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Network/loadBalancers/myVMSSSLB/backendAddressPools/myVMSSSLBBEPool", ""}
		expectedMsg := fmt.Sprintf("AgentPoolProfile.LoadBalancerBackendAddressPoolIDs can not contain empty string. Agent pool name: %s", agentPoolProfiles[0].Name)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})
}

func TestAgentPoolProfile_ValidateVirtualMachineScaleSet(t *testing.T) {
	t.Run("Should fail for invalid VMSS + Overprovisioning config", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = AvailabilitySet
		agentPoolProfiles[0].VMSSOverProvisioningEnabled = to.BoolPtr(true)
		expectedMsg := fmt.Sprintf("You have specified VMSS Overprovisioning in agent pool %s, but you did not specify VMSS", agentPoolProfiles[0].Name)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for invalid VMSS + Enable VMSS node public IP config", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = AvailabilitySet
		agentPoolProfiles[0].EnableVMSSNodePublicIP = to.BoolPtr(true)
		expectedMsg := fmt.Sprintf("You have enabled VMSS node public IP in agent pool %s, but you did not specify VMSS", agentPoolProfiles[0].Name)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for invalid LB + Enable VMSS node public IP config", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
			LoadBalancerSku: StandardLoadBalancerSku,
		}
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
		agentPoolProfiles[0].EnableVMSSNodePublicIP = to.BoolPtr(true)
		expectedMsg := fmt.Sprintf("You have enabled VMSS node public IP in agent pool %s, but you did not specify Basic Load Balancer SKU", agentPoolProfiles[0].Name)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for invalid VMSS + VnetSubnetID + FirstConsecutiveStaticIP config", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
		cs.Properties.MasterProfile.VnetSubnetID = "vnet"
		cs.Properties.MasterProfile.FirstConsecutiveStaticIP = "10.10.10.240"
		expectedMsg := "when masterProfile's availabilityProfile is VirtualMachineScaleSets and a vnetSubnetID is specified, the firstConsecutiveStaticIP should be empty and will be determined by an offset from the first IP in the vnetCidr"
		if err := cs.Properties.validateMasterProfile(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for VMSS master + AvailabilitySet agent pool", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = AvailabilitySet
		expectedMsg := "VirtualMachineScaleSets for master profile must be used together with virtualMachineScaleSets for agent profiles. Set \"availabilityProfile\" to \"VirtualMachineScaleSets\" for agent profiles"
		if err := cs.Properties.validateMasterProfile(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for VMSS + StorageAccount", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
		agentPoolProfiles[0].StorageProfile = StorageAccount
		expectedMsg := fmt.Sprintf("VirtualMachineScaleSets does not support %s disks.  Please specify \"storageProfile\": \"%s\" (recommended) or \"availabilityProfile\": \"%s\"", StorageAccount, ManagedDisks, AvailabilitySet)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for VMSS + VMAS", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.AgentPoolProfiles = append(cs.Properties.AgentPoolProfiles, &AgentPoolProfile{
			Name:                "agentpool2",
			VMSize:              "Standard_D2_v2",
			Count:               1,
			AvailabilityProfile: AvailabilitySet,
		})
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
		agentPoolProfiles[1].AvailabilityProfile = AvailabilitySet
		expectedMsg := "mixed mode availability profiles are not allowed. Please set either VirtualMachineScaleSets or AvailabilitySet in availabilityProfile for all agent pools"
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for VMSS + invalid LoadBalancerBackendAddressPoolIDs", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].AvailabilityProfile = VirtualMachineScaleSets
		agentPoolProfiles[0].LoadBalancerBackendAddressPoolIDs = []string{"/subscriptions/123/resourceGroups/rg/providers/Microsoft.Network/loadBalancers/myVMSSSLB/backendAddressPools/myVMSSSLBBEPool", ""}
		expectedMsg := fmt.Sprintf("AgentPoolProfile.LoadBalancerBackendAddressPoolIDs can not contain empty string. Agent pool name: %s", agentPoolProfiles[0].Name)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})
}

func TestAgentPoolProfile_ValidateAuditDEnabled(t *testing.T) {
	t.Run("Should have proper validation for auditd + distro combinations", func(t *testing.T) {
		t.Parallel()
		for _, distro := range DistroValues {
			cs := getK8sDefaultContainerService(false)
			agentPoolProfiles := cs.Properties.AgentPoolProfiles
			agentPoolProfiles[0].Distro = distro
			agentPoolProfiles[0].AuditDEnabled = to.BoolPtr(true)
			switch distro {
			case Flatcar:
				expectedMsg := fmt.Sprintf("You have enabled auditd in agent pool %s, but you did not specify an Ubuntu-based distro", agentPoolProfiles[0].Name)
				if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
					t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
				}
			case Ubuntu, Ubuntu1804, Ubuntu1804Gen2, AKSUbuntu1604, AKSUbuntu1804, ACC1604:
				if err := cs.Properties.validateAgentPoolProfiles(false); err != nil {
					t.Errorf("AuditDEnabled should work with distro %s, got error %s", distro, err.Error())
				}
			}
		}
	})
}

func TestMasterProfile_ValidateAuditDEnabled(t *testing.T) {
	t.Run("Should have proper validation for auditd + distro combinations", func(t *testing.T) {
		t.Parallel()
		for _, distro := range DistroValues {
			cs := getK8sDefaultContainerService(false)
			masterProfile := cs.Properties.MasterProfile
			masterProfile.Distro = distro
			masterProfile.AuditDEnabled = to.BoolPtr(true)
			switch distro {
			case Flatcar:
				expectedMsg := "You have enabled auditd for master vms, but you did not specify an Ubuntu-based distro."
				if err := cs.Properties.validateMasterProfile(false); err.Error() != expectedMsg {
					t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
				}
			case Ubuntu, Ubuntu1804, Ubuntu1804Gen2, AKSUbuntu1604, AKSUbuntu1804, ACC1604:
				if err := cs.Properties.validateMasterProfile(false); err != nil {
					t.Errorf("AuditDEnabled should work with distro %s, got error %s", distro, err.Error())
				}
			}
		}
	})
}

func TestValidateCustomCloudProfile(t *testing.T) {
	tests := []struct {
		name        string
		cs          *ContainerService
		expectedErr error
	}{
		{
			name: "valid run",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "custom profile is nil",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					LinuxProfile: &LinuxProfile{
						AdminUsername: "abc",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "PortalURL is empty for non-AzureStack CustomCloudProfile",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						Environment: &azure.Environment{
							Name: "NonAzureStack",
						},
						AuthenticationMethod: "azure_ad",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "PortalURL is empty for AzureStack",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						AuthenticationMethod: "azure_ad",
					},
				},
			},
			expectedErr: errors.New("portalURL needs to be specified when AzureStackCloud CustomCloudProfile is provided"),
		},
		{
			name: "PortalURL is invalid for AzureStack",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.testlocationinvalid.cotoso.com",
					},
				},
			},
			expectedErr: fmt.Errorf("portalURL needs to start with https://portal.%s. ", "testlocation"),
		},

		{
			name: "authenticationMethod has invalid value",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						AuthenticationMethod: "invalidAuthMethod",
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: errors.Errorf("authenticationMethod allowed values are '%s' and '%s'", ClientCertificateAuthMethod, ClientSecretAuthMethod),
		},
		{
			name: "identitySystem has invalid value",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						IdentitySystem: "invalidIdentySytem",
						PortalURL:      "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: errors.Errorf("identitySystem allowed values are '%s' and '%s'", AzureADIdentitySystem, ADFSIdentitySystem),
		},
		{
			name: "Dependencies location: china",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						DependenciesLocation: DependenciesLocation("china"),
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Dependencies location: public",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						DependenciesLocation: DependenciesLocation("public"),
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Dependencies location: german",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						DependenciesLocation: DependenciesLocation("german"),
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Dependencies location: usgovernment",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						DependenciesLocation: DependenciesLocation("usgovernment"),
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Dependencies location: invalid",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						DependenciesLocation: DependenciesLocation("invalidDependenciesLocation"),
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: errors.New("The invalidDependenciesLocation dependenciesLocation is not supported. The supported vaules are [ public china german usgovernment]"),
		},
		{
			name: " valid AzureAD and ClientSecret",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						IdentitySystem:       "azure_ad",
						AuthenticationMethod: "client_secret",
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: " valid ADFS and ClientCertificate",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						IdentitySystem:       "adfs",
						AuthenticationMethod: "client_certificate",
						PortalURL:            "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			gotErr := test.cs.validateCustomCloudProfile()
			if !helpers.EqualError(gotErr, test.expectedErr) {
				t.Logf("scenario %q", test.name)
				t.Errorf("expected error: %v, got: %v", test.expectedErr, gotErr)
			}
		})
	}
}

func TestValidateLocation(t *testing.T) {

	tests := []struct {
		name          string
		location      string
		propertiesnil bool
		cs            *ContainerService
		expectedErr   error
	}{

		{
			name:          "AzureStack location is empty",
			location:      "",
			propertiesnil: false,
			cs: &ContainerService{
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.testlocation.cotoso.com",
					},
				},
			},
			expectedErr: errors.New("missing ContainerService Location"),
		},
		{
			name:          "AzureStack UseInstanceMetadata is true",
			location:      "local",
			propertiesnil: false,
			cs: &ContainerService{
				Location: "local",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.local.cotoso.com",
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.16.14",
						KubernetesConfig: &KubernetesConfig{
							UseInstanceMetadata: to.BoolPtr(trueVal),
						},
					},
				},
			},
			expectedErr: errors.New("useInstanceMetadata shouldn't be set to true as feature not yet supported on Azure Stack"),
		},
		{
			name:          "AzureStack EtcdDiskSizeGB is 1024",
			location:      "local",
			propertiesnil: false,
			cs: &ContainerService{
				Location: "local",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.local.cotoso.com",
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.16.14",
						KubernetesConfig: &KubernetesConfig{
							EtcdDiskSizeGB: "1024",
						},
					},
				},
			},
			expectedErr: errors.Errorf("EtcdDiskSizeGB max size supported on Azure Stack is %d", MaxAzureStackManagedDiskSize),
		},
		{
			name:          "AzureStack EtcdDiskSizeGB is 1024",
			location:      "local",
			propertiesnil: false,
			cs: &ContainerService{
				Location: "local",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.local.cotoso.com",
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.16.14",
						KubernetesConfig: &KubernetesConfig{
							EtcdDiskSizeGB: "1024GB",
						},
					},
				},
			},
			expectedErr: errors.New("could not convert EtcdDiskSizeGB to int"),
		},
		{
			name:          "AzureStack AcceleratedNetworking is true",
			location:      "local",
			propertiesnil: false,
			cs: &ContainerService{
				Location: "local",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.local.cotoso.com",
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.16.14",
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:                         "testpool",
							Count:                        1,
							VMSize:                       "Standard_D2_v2",
							AcceleratedNetworkingEnabled: to.BoolPtr(trueVal),
						},
					},
				},
			},
			expectedErr: errors.New("AcceleratedNetworkingEnabled or AcceleratedNetworkingEnabledWindows shouldn't be set to true as feature is not yet supported on Azure Stack"),
		},
		{
			name:          "AzureStack AcceleratedNetworking is true",
			location:      "local",
			propertiesnil: false,
			cs: &ContainerService{
				Location: "local",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						PortalURL: "https://portal.local.cotoso.com",
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.16.14",
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:                                "testpool",
							Count:                               1,
							VMSize:                              "Standard_D2_v2",
							AcceleratedNetworkingEnabledWindows: to.BoolPtr(trueVal),
						},
					},
				},
			},
			expectedErr: errors.New("AcceleratedNetworkingEnabled or AcceleratedNetworkingEnabledWindows shouldn't be set to true as feature is not yet supported on Azure Stack"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(true)
			cs.Location = test.cs.Location
			if test.cs.Properties != nil {
				if test.cs.Properties.CustomCloudProfile != nil {
					cs.Properties.CustomCloudProfile = test.cs.Properties.CustomCloudProfile
				}

				if test.cs.Properties.OrchestratorProfile != nil {
					cs.Properties.OrchestratorProfile = test.cs.Properties.OrchestratorProfile
				}

				if test.cs.Properties.AgentPoolProfiles != nil {
					cs.Properties.AgentPoolProfiles = test.cs.Properties.AgentPoolProfiles
				}
			}

			if test.propertiesnil {
				cs.Properties = nil
			}
			gotErr := cs.Validate(false)
			if !helpers.EqualError(gotErr, test.expectedErr) {
				t.Logf("scenario %q", test.name)
				t.Errorf("expected error: %v, got: %v", test.expectedErr, gotErr)
			}
		})
	}
}

func TestValidateAcceleratedNetworkingEnabledWindows(t *testing.T) {

	tests := []struct {
		name        string
		cs          *ContainerService
		expectedErr error
	}{
		{
			name: "AcceleratedNetworkingEnabledWindows enabled",
			cs: &ContainerService{
				Properties: &Properties{
					MasterProfile: &MasterProfile{
						DNSPrefix: "foo",
						Count:     1,
						VMSize:    "Standard_D2_v3",
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.16.14",
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:                                "testpool",
							Count:                               1,
							VMSize:                              "Standard_D2_v2",
							AcceleratedNetworkingEnabledWindows: to.BoolPtr(true),
						},
					},
					LinuxProfile: &LinuxProfile{
						AdminUsername: "azureuser",
						SSH: struct {
							PublicKeys []PublicKey `json:"publicKeys" validate:"required,min=1"`
						}{
							PublicKeys: []PublicKey{{
								KeyData: "publickeydata",
							}},
						},
					},
				},
			},
			expectedErr: errors.New("Accelerated Networking is currently unstable for Windows + Kubernetes, please set acceleratedNetworkingEnabledWindows to false"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			gotErr := test.cs.Validate(false)
			if !helpers.EqualError(gotErr, test.expectedErr) {
				t.Logf("scenario %q", test.name)
				t.Errorf("expected error: %v, got: %v", test.expectedErr, gotErr)
			}
		})
	}
}

func TestValidateMasterProfileImageRef(t *testing.T) {
	tests := map[string]struct {
		properties    *Properties
		isUpdate      bool
		expectedError error
	}{
		"should not error when masterProfile includes both an ImageRef and a Distro configuration": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Distro: AKSUbuntu1604,
					ImageRef: &ImageReference{
						Name:           "name",
						ResourceGroup:  "rg",
						SubscriptionID: "sub-id",
						Gallery:        "gallery",
						Version:        "version",
					},
					DNSPrefix: "foo",
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when masterProfile includes both an ImageRef and a Distro configuration in update context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Distro: AKSUbuntu1604,
					ImageRef: &ImageReference{
						Name:           "name",
						ResourceGroup:  "rg",
						SubscriptionID: "sub-id",
						Gallery:        "gallery",
						Version:        "version",
					},
					DNSPrefix: "foo",
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
		"should not error when masterProfile includes an ImageRef configuration only": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					ImageRef: &ImageReference{
						Name:           "name",
						ResourceGroup:  "rg",
						SubscriptionID: "sub-id",
						Gallery:        "gallery",
						Version:        "version",
					},
					DNSPrefix: "foo",
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when masterProfile includes an ImageRef configuration only in an upgrade context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					ImageRef: &ImageReference{
						Name:           "name",
						ResourceGroup:  "rg",
						SubscriptionID: "sub-id",
						Gallery:        "gallery",
						Version:        "version",
					},
					DNSPrefix: "foo",
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
		"should not error when masterProfile includes a Distro configuration only": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Distro:    AKSUbuntu1604,
					DNSPrefix: "foo",
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when masterProfile includes a Distro configuration only in an upgrade context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					Distro:    AKSUbuntu1604,
					DNSPrefix: "foo",
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
		"should not error when masterProfile includes neither an explicit Distro nor ImageRef configuration": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					DNSPrefix: "foo",
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when masterProfile includes neither an explicit Distro nor ImageRef configuration in an upgrade context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				MasterProfile: &MasterProfile{
					DNSPrefix: "foo",
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
	}

	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := test.properties.validateMasterProfile(test.isUpdate)
			if !helpers.EqualError(err, test.expectedError) {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

func TestValidateAgentPoolProfilesImageRef(t *testing.T) {
	tests := map[string]struct {
		properties    *Properties
		isUpdate      bool
		expectedError error
	}{
		"should not error when AgentPoolProfile includes both an ImageRef and a Distro configuration": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "foo",
						Distro: AKSUbuntu1604,
						ImageRef: &ImageReference{
							Name:           "name",
							ResourceGroup:  "rg",
							SubscriptionID: "sub-id",
							Gallery:        "gallery",
							Version:        "version",
						},
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when AgentPoolProfile includes both an ImageRef and a Distro configuration in update context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "foo",
						Distro: AKSUbuntu1604,
						ImageRef: &ImageReference{
							Name:           "name",
							ResourceGroup:  "rg",
							SubscriptionID: "sub-id",
							Gallery:        "gallery",
							Version:        "version",
						},
					},
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
		"should not error when AgentPoolProfile includes an ImageRef configuration only": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo",
						ImageRef: &ImageReference{
							Name:           "name",
							ResourceGroup:  "rg",
							SubscriptionID: "sub-id",
							Gallery:        "gallery",
							Version:        "version",
						},
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when AgentPoolProfile includes an ImageRef configuration only in an upgrade context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo",
						ImageRef: &ImageReference{
							Name:           "name",
							ResourceGroup:  "rg",
							SubscriptionID: "sub-id",
							Gallery:        "gallery",
							Version:        "version",
						},
					},
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
		"should not error when AgentPoolProfile includes a Distro configuration only": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "foo",
						Distro: AKSUbuntu1604,
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when AgentPoolProfile includes a Distro configuration only in an upgrade context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:   "foo",
						Distro: AKSUbuntu1604,
					},
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
		"should not error when AgentPoolProfile includes neither an explicit Distro nor ImageRef configuration": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo",
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"should not error when AgentPoolProfile includes neither an explicit Distro nor ImageRef configuration in an upgrade context": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name: "foo",
					},
				},
			},
			isUpdate:      true,
			expectedError: nil,
		},
	}

	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := test.properties.validateAgentPoolProfiles(test.isUpdate)
			if !helpers.EqualError(err, test.expectedError) {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

func TestValidateAgentPoolProfilesOSDiskCachingType(t *testing.T) {
	tests := map[string]struct {
		properties    *Properties
		isUpdate      bool
		expectedError error
	}{
		"AgentPoolProfile with valid OSDiskCachingType: None": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:              "foo",
						OSDiskCachingType: string(compute.CachingTypesNone),
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"AgentPoolProfile with valid DataDiskCachingType: None": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "foo",
						DataDiskCachingType: string(compute.CachingTypesNone),
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"AgentPoolProfile with valid OSDiskCachingType: ReadWrite": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:              "bar",
						OSDiskCachingType: string(compute.CachingTypesReadWrite),
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"AgentPoolProfile with valid DataDiskCachingType: ReadWrite": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "bar",
						DataDiskCachingType: string(compute.CachingTypesReadWrite),
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"AgentPoolProfile with valid OSDiskCachingType: ReadOnly": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:              "baz",
						OSDiskCachingType: string(compute.CachingTypesReadOnly),
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"AgentPoolProfile with valid DataDiskCachingType: ReadOnly": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "baz",
						DataDiskCachingType: string(compute.CachingTypesReadOnly),
					},
				},
			},
			isUpdate:      false,
			expectedError: nil,
		},
		"AgentPoolProfile with invalid OSDiskCachingType": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:              "frick",
						OSDiskCachingType: "The Magnificent Ambersons original cut",
					},
				},
			},
			isUpdate:      false,
			expectedError: errors.Errorf("Invalid osDiskCachingType value \"%s\" for agentPoolProfile \"%s\", please use one of the following versions: %s", "The Magnificent Ambersons original cut", "frick", cachingTypesValidValues),
		},
		"AgentPoolProfile with invalid DataDiskCachingType": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:                "frack",
						DataDiskCachingType: "Spear of Longinus",
					},
				},
			},
			isUpdate:      false,
			expectedError: errors.Errorf("Invalid dataDiskCachingType value \"%s\" for agentPoolProfile \"%s\", please use one of the following versions: %s", "Spear of Longinus", "frack", cachingTypesValidValues),
		},
		"AgentPoolProfile with invalid OSDiskCachingType for Ephemeral Disk": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: Kubernetes,
				},
				AgentPoolProfiles: []*AgentPoolProfile{
					{
						Name:              "foo",
						OSDiskCachingType: string(compute.CachingTypesReadWrite),
						StorageProfile:    Ephemeral,
					},
				},
			},
			isUpdate:      false,
			expectedError: errors.Errorf("Invalid osDiskCachingType value \"%s\" for agentPoolProfile \"%s\" using Ephemeral Disk, you must use: %s", string(compute.CachingTypesReadWrite), "foo", string(compute.CachingTypesReadOnly)),
		},
	}

	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := test.properties.validateAgentPoolProfiles(test.isUpdate)
			if !helpers.EqualError(err, test.expectedError) {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

func TestValidateAzureStackSupport(t *testing.T) {
	tests := []struct {
		name               string
		networkPlugin      string
		masterAvailability string
		agentAvailability  string
		expectedErr        error
	}{
		{
			name:               "AzureStack supports the kubenet network plugin",
			networkPlugin:      "kubenet",
			masterAvailability: AvailabilitySet,
			agentAvailability:  AvailabilitySet,
			expectedErr:        nil,
		},
		{
			name:               "AzureStack supports for the azure network plugin is in preview",
			networkPlugin:      "azure",
			masterAvailability: AvailabilitySet,
			agentAvailability:  AvailabilitySet,
			expectedErr:        nil,
		},
		{
			name:               "AzureStack only supports kubenet and azure network plugins",
			networkPlugin:      NetworkPluginFlannel,
			masterAvailability: AvailabilitySet,
			agentAvailability:  AvailabilitySet,
			expectedErr:        errors.New("kubernetesConfig.networkPlugin 'flannel' is not supported on Azure Stack clouds"),
		},
		{
			name:               "AzureStack does not support VMSS on the master pool",
			networkPlugin:      "",
			masterAvailability: VirtualMachineScaleSets,
			agentAvailability:  VirtualMachineScaleSets,
			expectedErr:        errors.New("masterProfile.availabilityProfile should be set to 'AvailabilitySet' on Azure Stack clouds"),
		},
		{
			name:               "AzureStack does not support VMSS on the agent pools",
			networkPlugin:      "kubenet",
			masterAvailability: AvailabilitySet,
			agentAvailability:  VirtualMachineScaleSets,
			expectedErr:        errors.New("agentPoolProfiles[agentpool].availabilityProfile should be set to 'AvailabilitySet' on Azure Stack clouds"),
		},
		{
			name:               "AzureStack defaults masterAvailabilityProfile to 'AvailabilitySet'",
			networkPlugin:      "kubenet",
			masterAvailability: "",
			agentAvailability:  AvailabilitySet,
			expectedErr:        nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cs := getK8sDefaultContainerService(false)
			cs.Properties.CustomCloudProfile = &CustomCloudProfile{
				PortalURL: "https://portal.westus.contoso.com",
			}
			cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}
			if test.networkPlugin != "" {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = test.networkPlugin
			}
			if test.masterAvailability != "" {
				cs.Properties.MasterProfile.AvailabilityProfile = test.masterAvailability
			}
			if test.agentAvailability != "" {
				for _, agentPool := range cs.Properties.AgentPoolProfiles {
					pool := agentPool
					pool.AvailabilityProfile = test.agentAvailability
					break
				}
			}
			cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.16.13"
			if err := cs.Validate(false); !helpers.EqualError(err, test.expectedErr) {
				t.Logf("scenario %q", test.name)
				t.Logf("FIXME: expected error: %v, got: %v", test.expectedErr, err)
			}
		})
	}
}

func TestValidateKubernetesImageBaseType(t *testing.T) {
	tests := map[string]struct {
		k             *KubernetesConfig
		expectedError error
	}{
		"should not error for zero-value kubernetesImageBaseType value": {
			k:             &KubernetesConfig{},
			expectedError: nil,
		},
		"should not error for empty kubernetesImageBaseType value": {
			k: &KubernetesConfig{
				KubernetesImageBaseType: "",
			},
			expectedError: nil,
		},
		"should not error for valid kubernetesImageBaseType value 'gcr'": {
			k: &KubernetesConfig{
				KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
			},
			expectedError: nil,
		},
		"should not error for valid kubernetesImageBaseType value 'mcr'": {
			k: &KubernetesConfig{
				KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
			},
			expectedError: nil,
		},
		"should error on unknown kubernetesImageBaseType value": {
			k: &KubernetesConfig{
				KubernetesImageBaseType: "quay",
			},
			expectedError: errors.Errorf("Invalid kubernetesImageBaseType value \"%s\", please use one of the following versions: %s", "quay", kubernetesImageBaseTypeValidVersions),
		},
	}

	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := test.k.validateKubernetesImageBaseType()
			if !helpers.EqualError(err, test.expectedError) {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

func TestValidateContainerRuntimeConfig(t *testing.T) {
	tests := map[string]struct {
		k             *KubernetesConfig
		expectedError error
	}{
		"should succeed if unspecified with docker": {
			k:             &KubernetesConfig{},
			expectedError: nil,
		},
		"should succeed if unspecified with containerd": {
			k: &KubernetesConfig{
				ContainerRuntime: Containerd,
			},
			expectedError: nil,
		},
		"should succeed if config is defined but key not present": {
			k: &KubernetesConfig{
				ContainerRuntimeConfig: map[string]string{},
			},
			expectedError: nil,
		},
		"should succeed with containerd if config is defined but key not present": {
			k: &KubernetesConfig{
				ContainerRuntime:       Containerd,
				ContainerRuntimeConfig: map[string]string{},
			},
			expectedError: nil,
		},
		"should fail on empty string": {
			k: &KubernetesConfig{
				ContainerRuntimeConfig: map[string]string{
					ContainerDataDirKey: "",
				},
			},
			expectedError: errors.Errorf("OrchestratorProfile.KubernetesConfig.ContainerRuntimeConfig.DataDir '' is invalid: must not be empty"),
		},
		"should fail on relative path": {
			k: &KubernetesConfig{
				ContainerRuntimeConfig: map[string]string{
					ContainerDataDirKey: "mnt/docker",
				},
			},
			expectedError: errors.Errorf("OrchestratorProfile.KubernetesConfig.ContainerRuntimeConfig.DataDir 'mnt/docker' is invalid: must be absolute path"),
		},
		"should pass with absolute path": {
			k: &KubernetesConfig{
				ContainerRuntimeConfig: map[string]string{
					ContainerDataDirKey: "/mnt/docker",
				},
			},
			expectedError: nil,
		},
	}

	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := test.k.validateContainerRuntimeConfig()
			if !helpers.EqualError(err, test.expectedError) {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

func TestValidateConnectedClusterProfile(t *testing.T) {
	addon := &KubernetesAddon{}

	t.Run("incomplete connected cluster profile", func(t *testing.T) {
		err := addon.validateArcAddonConfig()
		expected := errors.New("azure-arc-onboarding addon configuration must have a 'location' property; azure-arc-onboarding addon configuration must have a 'tenantID' property; azure-arc-onboarding addon configuration must have a 'subscriptionID' property; azure-arc-onboarding addon configuration must have a 'resourceGroup' property; azure-arc-onboarding addon configuration must have a 'clusterName' property; azure-arc-onboarding addon configuration must have a 'clientID' property; azure-arc-onboarding addon configuration must have a 'clientSecret' property")
		if !helpers.EqualError(err, expected) {
			t.Errorf("expected error: %v, got: %v", expected, err)
		}
	})

	addon.Config = make(map[string]string)
	addon.Config["location"] = "location"
	addon.Config["tenantID"] = "tenantID"
	addon.Config["subscriptionID"] = "subscriptionID"
	addon.Config["resourceGroup"] = "resourceGroup"
	addon.Config["clusterName"] = "clusterName"
	addon.Config["clientID"] = "clientID"
	addon.Config["clientSecret"] = "clientSecret"

	t.Run("complete connected cluster profile", func(t *testing.T) {
		err := addon.validateArcAddonConfig()
		if err != nil {
			t.Errorf("error not expected, got: %v", err)
		}
	})
}
