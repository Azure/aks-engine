// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/blang/semver"
	"github.com/pkg/errors"
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
		"should error when KubernetesConfig populated for non-Kubernetes OrchestratorType": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "DCOS",
					KubernetesConfig: &KubernetesConfig{
						ClusterSubnet: "10.0.0.0/16",
					},
				},
			},
			expectedError: "KubernetesConfig can be specified only when OrchestratorType is Kubernetes",
		},
		"should error when KubernetesConfig has invalid etcd version": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						EtcdVersion: "1.0.0",
					},
				},
			},
			expectedError: "Invalid etcd version \"1.0.0\", please use one of the following versions: [2.2.5 2.3.0 2.3.1 2.3.2 2.3.3 2.3.4 2.3.5 2.3.6 2.3.7 2.3.8 3.0.0 3.0.1 3.0.2 3.0.3 3.0.4 3.0.5 3.0.6 3.0.7 3.0.8 3.0.9 3.0.10 3.0.11 3.0.12 3.0.13 3.0.14 3.0.15 3.0.16 3.0.17 3.1.0 3.1.1 3.1.2 3.1.2 3.1.3 3.1.4 3.1.5 3.1.6 3.1.7 3.1.8 3.1.9 3.1.10 3.2.0 3.2.1 3.2.2 3.2.3 3.2.4 3.2.5 3.2.6 3.2.7 3.2.8 3.2.9 3.2.11 3.2.12 3.2.13 3.2.14 3.2.15 3.2.16 3.2.23 3.2.24 3.2.25 3.2.26 3.3.0 3.3.1 3.3.8 3.3.9 3.3.10 3.3.13 3.3.15 3.3.18]",
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
			expectedError: "Invalid containerd version \"1.0.0\", please use one of the following versions: [1.3.2]",
		},
		"should error when KubernetesConfig has invalid containerd version for kata-containers runtime": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:  KataContainers,
						ContainerdVersion: "1.0.0",
					},
				},
			},
			expectedError: "Invalid containerd version \"1.0.0\", please use one of the following versions: [1.3.2]",
		},
		"should error when KubernetesConfig has containerdVersion value for docker container runtime": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						ContainerRuntime:  Docker,
						ContainerdVersion: "1.0.0",
					},
				},
			},
			expectedError: fmt.Sprintf("containerdVersion is only valid in a non-docker context, use %s or %s containerRuntime values instead if you wish to provide a containerdVersion", Containerd, KataContainers),
		},
		"should error when KubernetesConfig has containerdVersion value for default (empty string) container runtime": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					KubernetesConfig: &KubernetesConfig{
						ContainerdVersion: "1.0.0",
					},
				},
			},
			expectedError: fmt.Sprintf("containerdVersion is only valid in a non-docker context, use %s or %s containerRuntime values instead if you wish to provide a containerdVersion", Containerd, KataContainers),
		},
		"should error when KubernetesConfig has enableAggregatedAPIs enabled and enableRBAC disabled": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.13.11",
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
					OrchestratorVersion: common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false)),
					KubernetesConfig: &KubernetesConfig{
						EnableRbac: &falseVal,
					},
				},
			},
			expectedError: fmt.Sprintf("RBAC support is required for Kubernetes version 1.15.0 or greater; unable to build Kubernetes v%s cluster with enableRbac=false", common.GetLatestPatchVersion("1.15", common.GetAllSupportedKubernetesVersions(false, false))),
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
					OrchestratorVersion: "1.13.11",
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
		"should error when KubernetesConfig has enablePodSecurity enabled with invalid settings": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.13.11",
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
					DcosConfig:       &DcosConfig{},
				},
			},
		},
		"should error when DcosConfig orchestrator has invalid configuration": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "DCOS",
					OrchestratorVersion: "1.12.0",
				},
			},
			expectedError: "the following OrchestratorProfile configuration is not supported: OrchestratorType: DCOS, OrchestratorRelease: , OrchestratorVersion: 1.12.0. Please check supported Release or Version for this build of aks-engine",
		},
		"should error when DcosConfig orchestrator configuration has invalid static IP": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "DCOS",
					DcosConfig: &DcosConfig{
						BootstrapProfile: &BootstrapProfile{
							StaticIP: "0.0.0.0.0.0",
						},
					},
				},
			},
			expectedError: "DcosConfig.BootstrapProfile.StaticIP '0.0.0.0.0.0' is an invalid IP address",
		},
		"should error when DcosConfig populated for non-Kubernetes OrchestratorType 1": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					DcosConfig: &DcosConfig{
						DcosWindowsBootstrapURL: "http://www.microsoft.com",
					},
				},
			},
			expectedError: "DcosConfig can be specified only when OrchestratorType is DCOS",
		},
		"should error when DcosConfig populated for non-Kubernetes OrchestratorType 2": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType: "Kubernetes",
					DcosConfig: &DcosConfig{
						DcosWindowsBootstrapURL: "http://www.microsoft.com",
						DcosBootstrapURL:        "http://www.microsoft.com",
					},
				},
			},
			expectedError: "DcosConfig can be specified only when OrchestratorType is DCOS",
		},
		"kubernetes should have failed on old patch version": {
			properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorType:    "Kubernetes",
					OrchestratorVersion: "1.6.0",
				},
			},
			expectedError: fmt.Sprint("the following OrchestratorProfile configuration is not supported: OrchestratorType: \"Kubernetes\", OrchestratorRelease: \"\", OrchestratorVersion: \"1.6.0\". Please use one of the following versions: ", common.GetAllSupportedKubernetesVersions(false, false)),
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
					OrchestratorVersion: "v1.13.11",
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
	}

	for testName, test := range tests {
		test := test
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

func Test_KubernetesConfig_Validate(t *testing.T) {
	// Tests that should pass across all versions
	for _, k8sVersion := range common.GetAllSupportedKubernetesVersions(true, false) {
		c := KubernetesConfig{}
		if err := c.Validate(k8sVersion, false, false); err != nil {
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
		if err := c.Validate(k8sVersion, false, false); err != nil {
			t.Errorf("should not error on a KubernetesConfig with valid param values: %v", err)
		}

		c = KubernetesConfig{
			ClusterSubnet: "10.16.x.0/invalid",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error on invalid ClusterSubnet")
		}

		c = KubernetesConfig{
			DockerBridgeSubnet: "10.120.1.0/invalid",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error on invalid DockerBridgeSubnet")
		}

		c = KubernetesConfig{
			KubeletConfig: map[string]string{
				"--non-masquerade-cidr": "10.120.1.0/24",
			},
		}
		if err := c.Validate(k8sVersion, false, false); err != nil {
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
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error on invalid MaxPods")
		}

		c = KubernetesConfig{
			KubeletConfig: map[string]string{
				"--node-status-update-frequency": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error on invalid --node-status-update-frequency kubelet config")
		}

		c = KubernetesConfig{
			ControllerManagerConfig: map[string]string{
				"--node-monitor-grace-period": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
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
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when --node-monitor-grace-period is not sufficiently larger than --node-status-update-frequency kubelet config")
		}

		c = KubernetesConfig{
			ControllerManagerConfig: map[string]string{
				"--pod-eviction-timeout": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error on invalid --pod-eviction-timeout")
		}

		c = KubernetesConfig{
			ControllerManagerConfig: map[string]string{
				"--route-reconciliation-period": "invalid",
			},
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error on invalid --route-reconciliation-period")
		}

		c = KubernetesConfig{
			DNSServiceIP: "192.168.0.10",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when DNSServiceIP but not ServiceCidr")
		}

		c = KubernetesConfig{
			ServiceCidr: "192.168.0.10/24",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when ServiceCidr but not DNSServiceIP")
		}

		c = KubernetesConfig{
			DNSServiceIP: "invalid",
			ServiceCidr:  "192.168.0.0/24",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when DNSServiceIP is invalid")
		}

		c = KubernetesConfig{
			DNSServiceIP: "192.168.1.10",
			ServiceCidr:  "192.168.0.0/not-a-len",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when ServiceCidr is invalid")
		}

		c = KubernetesConfig{
			DNSServiceIP: "192.168.1.10",
			ServiceCidr:  "192.168.0.0/24",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when DNSServiceIP is outside of ServiceCidr")
		}

		c = KubernetesConfig{
			DNSServiceIP: "172.99.255.255",
			ServiceCidr:  "172.99.0.1/16",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when DNSServiceIP is broadcast address of ServiceCidr")
		}

		c = KubernetesConfig{
			DNSServiceIP: "172.99.0.1",
			ServiceCidr:  "172.99.0.1/16",
		}
		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when DNSServiceIP is first IP of ServiceCidr")
		}

		c = KubernetesConfig{
			DNSServiceIP: "172.99.255.10",
			ServiceCidr:  "172.99.0.1/16",
		}
		if err := c.Validate(k8sVersion, false, false); err != nil {
			t.Error("should not error when DNSServiceIP and ServiceCidr are valid")
		}

		c = KubernetesConfig{
			ClusterSubnet: "192.168.0.1/24",
			NetworkPlugin: "azure",
		}

		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when ClusterSubnet has a mask of 24 bits or higher")
		}

		c = KubernetesConfig{
			ProxyMode: KubeProxyMode("invalid"),
		}

		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when ProxyMode has an invalid string value")
		}

		for _, validProxyModeValue := range []KubeProxyMode{KubeProxyModeIPTables, KubeProxyModeIPVS} {
			c = KubernetesConfig{
				ProxyMode: validProxyModeValue,
			}

			if err := c.Validate(k8sVersion, false, false); err != nil {
				t.Error("should error when ProxyMode has a valid string value")
			}

			c = KubernetesConfig{
				ProxyMode: validProxyModeValue,
			}

			if err := c.Validate(k8sVersion, false, false); err != nil {
				t.Error("should error when ProxyMode has a valid string value")
			}
		}
	}

	// Tests that apply to 1.6 and later releases
	for _, k8sVersion := range common.GetAllSupportedKubernetesVersions(false, false) {
		c := KubernetesConfig{
			CloudProviderBackoff:   to.BoolPtr(true),
			CloudProviderRateLimit: to.BoolPtr(true),
		}
		if err := c.Validate(k8sVersion, false, false); err != nil {
			t.Error("should not error when basic backoff and rate limiting are set to true with no options")
		}
	}

	// Tests that apply to 1.8 and later releases
	for _, k8sVersion := range common.GetVersionsGt(common.GetAllSupportedKubernetesVersions(true, false), "1.8.0", true, true) {
		c := KubernetesConfig{
			UseCloudControllerManager: to.BoolPtr(true),
		}
		if err := c.Validate(k8sVersion, false, false); err != nil {
			t.Error("should not error because UseCloudControllerManager is available since v1.8")
		}
	}

	// Tests that apply to dualstack with 1.16 and later releases
	for _, k8sVersion := range common.GetVersionsGt(common.GetAllSupportedKubernetesVersions(false, false), "1.16.0", true, true) {
		c := KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
		}

		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when more than 1 cluster subnet provided with ipv6dualstack feature disabled")
		}

		// validate config when ipv6 dual stack feature is enabled
		c = KubernetesConfig{
			NetworkPlugin: "azure",
		}

		if err := c.Validate(k8sVersion, false, true); err == nil {
			t.Error("should error when network plugin is not kubenet")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8,fec0::/7",
		}

		if err := c.Validate(k8sVersion, false, true); err == nil {
			t.Error("should error when more than 2 cluster subnets provided")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
		}

		if err := c.Validate(k8sVersion, false, true); err == nil {
			t.Error("should error when proxy mode is not set to ipvs")
		}

		c = KubernetesConfig{
			ServiceCidr: "10.0.0.0/16,fe80:20d::/112",
		}

		if err := c.Validate(k8sVersion, false, false); err == nil {
			t.Error("should error when more than 1 service cidr provided with ipv6dualstack feature disabled")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,fe80:20d::/112,fec0::/7",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true); err == nil {
			t.Error("should error when more than 2 service cidr provided with ipv6dualstack feature enabled")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,2001:db8::/129",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true); err == nil {
			t.Error("should error when secondary cidr is invalid with ipv6dualstack feature enabled")
		}

		c = KubernetesConfig{
			NetworkPlugin: "kubenet",
			ClusterSubnet: "10.244.0.0/16,ace:cab:deca::/8",
			ProxyMode:     "ipvs",
			ServiceCidr:   "10.0.0.0/16,fe80:20d::/112",
			DNSServiceIP:  "10.0.0.10",
		}

		if err := c.Validate(k8sVersion, false, true); err != nil {
			t.Error("shouldn't have errored with ipv6 dual stack feature enabled")
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
	expectedMsg = "customKubeAPIServerImage, customKubeControllerManagerImage, customKubeProxyImage, customKubeSchedulerImage or customKubeBinaryURL have no effect in Kubernetes version 1.16 or earlier"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}

	p.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage = "example.azurecr.io/hyperkube-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeAPIServerImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeControllerManagerImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeProxyImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeSchedulerImage = ""
	p.OrchestratorProfile.KubernetesConfig.CustomKubeBinaryURL = ""
	err = p.validateCustomKubeComponent()
	if err != nil {
		t.Errorf("should not error because custom hyperkube image can be used in 1.16, got error : %s", err.Error())
	}
}

func Test_Properties_ValidatePrivateAzureRegistryServer(t *testing.T) {
	p := &Properties{}
	p.OrchestratorProfile = &OrchestratorProfile{}
	p.OrchestratorProfile.OrchestratorType = Kubernetes
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{}

	p.OrchestratorProfile.OrchestratorVersion = "1.16.0"
	p.OrchestratorProfile.KubernetesConfig.PrivateAzureRegistryServer = "example.azurecr.io"
	err := p.validatePrivateAzureRegistryServer()
	expectedMsg := "customHyperkubeImage must be provided when privateAzureRegistryServer is provided"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}

	p.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage = "example.azurecr.io/hyperkube-amd64:tag"
	err = p.validatePrivateAzureRegistryServer()
	if err != nil {
		t.Errorf("should not error because CustomHyperkubeImage is provided, got error : %s", err.Error())
	}

	p.OrchestratorProfile.OrchestratorVersion = "1.17.0"
	p.OrchestratorProfile.KubernetesConfig.PrivateAzureRegistryServer = "example.azurecr.io"
	err = p.validatePrivateAzureRegistryServer()
	expectedMsg = "customKubeAPIServerImage, customKubeControllerManagerImage, customKubeProxyImage or customKubeSchedulerImage must be provided when privateAzureRegistryServer is provided"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message : %s to be thrown, but got : %s", expectedMsg, err.Error())
	}

	p.OrchestratorProfile.KubernetesConfig.PrivateAzureRegistryServer = "example.azurecr.io"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeAPIServerImage = "example.azurecr.io/kube-apiserver-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeControllerManagerImage = "example.azurecr.io/kube-controller-manager-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeProxyImage = "example.azurecr.io/kube-proxy-amd64:tag"
	p.OrchestratorProfile.KubernetesConfig.CustomKubeSchedulerImage = "example.azurecr.io/kube-scheduler-amd64:tag"
	err = p.validatePrivateAzureRegistryServer()
	if err != nil {
		t.Errorf("should not error because CustomKubeAPIServerImage, CustomKubeControllerManagerImage, CustomKubeProxyImage, and CustomKubeSchedulerImage are provided, got error : %s", err.Error())
	}

	p.OrchestratorProfile.KubernetesConfig.PrivateAzureRegistryServer = ""
	err = p.validatePrivateAzureRegistryServer()
	if err != nil {
		t.Errorf("should not error because PrivateAzureRegistryServer is not provided, got error : %s", err.Error())
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
					OSType:              "Windows",
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

	t.Run("aadProfiles should not be supported non-Kubernetes orchestrators", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.OrchestratorProfile = &OrchestratorProfile{
			OrchestratorType: DCOS,
		}
		cs.Properties.AADProfile = &AADProfile{
			ClientAppID: "92444486-5bc3-4291-818b-d53ae480991b",
			ServerAppID: "403f018b-4d89-495b-b548-0cf9868cdb0a",
		}
		expectedMsg := "'aadProfile' is only supported by orchestrator 'Kubernetes'"
		if err := cs.Properties.validateAADProfile(); err == nil || err.Error() != expectedMsg {
			t.Errorf("error should have occurred with msg : %s, but got : %s", expectedMsg, err.Error())
		}
	})
}

func TestProperties_ValidateInvalidStruct(t *testing.T) {
	cs := getK8sDefaultContainerService(false)
	cs.Properties.OrchestratorProfile = &OrchestratorProfile{}
	expectedMsg := "missing Properties.OrchestratorProfile.OrchestratorType"
	if err := cs.Validate(false); err == nil || err.Error() != expectedMsg {
		t.Errorf("expected validation error with message : %s", err.Error())
	}
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
		if err := p.validateContainerRuntime(); err != nil {
			t.Errorf(
				"should not error on containerRuntime=\"%s\"",
				runtime,
			)
		}
	}

	p.OrchestratorProfile.KubernetesConfig.ContainerRuntime = "not-existing"
	if err := p.validateContainerRuntime(); err == nil {
		t.Errorf(
			"should error on invalid containerRuntime",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.ContainerRuntime = KataContainers
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			OSType: Windows,
		},
	}
	if err := p.validateContainerRuntime(); err == nil {
		t.Errorf(
			"should error on kata-containers for windows clusters",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.ContainerRuntime = Containerd
	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			OSType: Windows,
		},
	}
	if err := p.validateContainerRuntime(); err == nil {
		t.Errorf(
			"should error on containerd for windows clusters",
		)
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
			name: "flannel addon enabled w/ NetworkPlugin=flannel",
			p: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					KubernetesConfig: &KubernetesConfig{
						NetworkPlugin: NetworkPluginFlannel,
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
						NetworkPlugin: DefaultNetworkPlugin,
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
						NetworkPlugin: "kubenet",
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
						NetworkPolicy: "calico",
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
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error on azure-policy when ServicePrincipalProfile is empty",
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
	p.OrchestratorProfile.OrchestratorRelease = "1.13"
	if err := p.validateAddons(); err != nil {
		t.Errorf(
			"should not error on azure-policy with k8s >= 1.12",
		)
	}

	p.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error on azure-policy with managed identity",
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

	p.OrchestratorProfile.OrchestratorRelease = "1.13"
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

	p.AgentPoolProfiles = []*AgentPoolProfile{
		{
			AvailabilityProfile: AvailabilitySet,
			Distro:              CoreOS,
		},
	}

	p.MasterProfile = &MasterProfile{
		Distro: CoreOS,
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "smb-flexvolume",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error using incompatible addon with coreos (smb-flexvolume)",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "keyvault-flexvolume",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error using incompatible addon with coreos (keyvault-flexvolume)",
		)
	}

	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		Addons: []KubernetesAddon{
			{
				Name:    "blobfuse-flexvolume",
				Enabled: to.BoolPtr(true),
			},
		},
	}

	if err := p.validateAddons(); err == nil {
		t.Errorf(
			"should error using incompatible addon with coreos (blobfuse-flexvolume)",
		)
	}

	// appgw-ingress add-on

	// Basic test with UseManagedIdentity
	p.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		NetworkPlugin:      "azure",
		UseManagedIdentity: true,
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
	for _, version := range common.GetAllSupportedKubernetesVersions(false, true) {
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
	for _, version := range common.GetAllSupportedKubernetesVersions(false, false) {
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
			orchestratorRelease: "1.13",
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
			orchestratorRelease: "1.13",
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
					UseManagedIdentity: test.useManagedIdentity,
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
			orchestratorRelease: "1.13",
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
			orchestratorRelease: "1.13",
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
		expectedErr                 string
	}{
		{
			name:                "Agent profile with zones vmas",
			orchestratorRelease: "1.13",
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
			expectedErr: "VirtualMachineScaleSets for master profile must be used together with virtualMachineScaleSets for agent profiles. Set \"availabilityProfile\" to \"VirtualMachineScaleSets\" for agent profiles",
		},
		{
			name:                "Master profile with zones and Agent profile without zones",
			orchestratorRelease: "1.13",
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
			expectedErr: "Availability Zones need to be defined for master profile and all agent pool profiles. Please set \"availabilityZones\" for all profiles",
		},
		{
			name:                "Master profile without zones and Agent profile with zones",
			orchestratorRelease: "1.13",
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
			expectedErr: "Availability Zones need to be defined for master profile and all agent pool profiles. Please set \"availabilityZones\" for all profiles",
		},
		{
			name:                "all zones and basic loadbalancer",
			orchestratorRelease: "1.13",
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
			expectedErr: "Availability Zones requires Standard LoadBalancer. Please set KubernetesConfig \"LoadBalancerSku\" to \"Standard\"",
		},
		{
			name:                        "all zones with standard loadbalancer and false excludeMasterFromStandardLB",
			orchestratorRelease:         "1.13",
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
			expectedErr: "standard loadBalancerSku should exclude master nodes. Please set KubernetesConfig \"ExcludeMasterFromStandardLB\" to \"true\"",
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

			if err := cs.Validate(false); err != nil {
				expectedMsg := test.expectedErr
				if err.Error() != expectedMsg {
					t.Errorf("expected error with message : %s, but got : %s", expectedMsg, err.Error())
				}
			} else {
				t.Errorf("error should have occurred")
			}
		})
	}
}

func TestProperties_ValidateLoadBalancer(t *testing.T) {
	tests := []struct {
		name                string
		orchestratorRelease string
		loadBalancerSku     string
		masterProfile       *MasterProfile
		agentProfiles       []*AgentPoolProfile
		expectedErr         bool
		expectedErrStr      string
	}{
		{
			name:                "lowercase basic LB",
			orchestratorRelease: "1.13",
			loadBalancerSku:     "basic",
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
		{
			name:                "Basic LB",
			orchestratorRelease: "1.13",
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
			orchestratorRelease: "1.13",
			loadBalancerSku:     "standard",
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
		{
			name:                "Standard LB",
			orchestratorRelease: "1.13",
			loadBalancerSku:     StandardLoadBalancerSku,
			masterProfile: &MasterProfile{
				Count:               3,
				DNSPrefix:           "foo",
				VMSize:              "Standard_DS2_v2",
				AvailabilityProfile: VirtualMachineScaleSets,
			},
		},
		{
			name:                "empty string LB value",
			orchestratorRelease: "1.13",
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
			orchestratorRelease: "1.13",
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
				LoadBalancerSku: test.loadBalancerSku,
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
			cs.Properties.OrchestratorProfile.OrchestratorRelease = "1.13"
			cs.Properties.MasterProfile = test.masterProfile
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Validate(true)
			if err.Error() != test.expectedMsg {
				t.Errorf("expected error message : %s, but got %s", test.expectedMsg, err.Error())
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
				VnetSubnetID:             validVNetSubnetID,
				Count:                    1,
				DNSPrefix:                "foo",
				VMSize:                   "Standard_DS2_v2",
				FirstConsecutiveStaticIP: "10.0.0.invalid",
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
				VnetSubnetID:             validVNetSubnetID,
				Count:                    1,
				DNSPrefix:                "foo",
				VMSize:                   "Standard_DS2_v2",
				FirstConsecutiveStaticIP: "10.0.0.1",
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
			cs.Properties.OrchestratorProfile.OrchestratorRelease = "1.13"
			cs.Properties.MasterProfile = test.masterProfile
			cs.Properties.AgentPoolProfiles = test.agentPoolProfiles
			err := cs.Validate(true)
			if err.Error() != test.expectedMsg {
				t.Errorf("expected error message : %s, but got %s", test.expectedMsg, err.Error())
			}
		})
	}
}

func TestWindowsProfile_Validate(t *testing.T) {
	tests := []struct {
		name             string
		orchestratorType string
		w                *WindowsProfile
		expectedMsg      string
	}{
		{
			name:             "unsupported orchestrator",
			orchestratorType: "Mesos",
			w: &WindowsProfile{
				WindowsImageSourceURL: "http://fakeWindowsImageSourceURL",
			},
			expectedMsg: "Windows Custom Images are only supported if the Orchestrator Type is DCOS or Kubernetes",
		},
		{
			name:             "empty adminUsername",
			orchestratorType: "Kubernetes",
			w: &WindowsProfile{
				WindowsImageSourceURL: "http://fakeWindowsImageSourceURL",
				AdminUsername:         "",
				AdminPassword:         "password",
			},
			expectedMsg: "WindowsProfile.AdminUsername is required, when agent pool specifies windows",
		},
		{
			name:             "empty password",
			orchestratorType: "DCOS",
			w: &WindowsProfile{
				WindowsImageSourceURL: "http://fakeWindowsImageSourceURL",
				AdminUsername:         "azure",
				AdminPassword:         "",
			},
			expectedMsg: "WindowsProfile.AdminPassword is required, when agent pool specifies windows",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.w.Validate(test.orchestratorType)
			if err.Error() != test.expectedMsg {
				t.Errorf("should error on unsupported orchType with msg : %s, but got : %s", test.expectedMsg, err.Error())
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

	t.Run("Should not support os type other than linux for ipv6 dual stack feature", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(true)
		cs.Properties.FeatureFlags = &FeatureFlags{EnableIPv6DualStack: true}
		masterProfile := cs.Properties.MasterProfile
		masterProfile.Distro = CoreOS
		expectedMsg := fmt.Sprintf("Dual stack feature is currently supported only with Ubuntu, but master is of distro type %s", masterProfile.Distro)
		if err := cs.Properties.validateMasterProfile(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}

		masterProfile.Distro = Ubuntu
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].OSType = Windows
		expectedMsg = fmt.Sprintf("Dual stack feature is supported only with Linux, but agent pool '%s' is of os type %s", agentPoolProfiles[0].Name, agentPoolProfiles[0].OSType)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}

		agentPoolProfiles[0].OSType = Linux
		agentPoolProfiles[0].Distro = CoreOS
		expectedMsg = fmt.Sprintf("Dual stack feature is currently supported only with Ubuntu, but agent pool '%s' is of distro type %s", agentPoolProfiles[0].Name, agentPoolProfiles[0].Distro)
		if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
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

	t.Run("Should not support orchestratorTypes other than Kubernetes/DCOS", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.OrchestratorProfile.OrchestratorType = SwarmMode
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].CustomNodeLabels = map[string]string{
			"foo": "bar",
		}
		expectedMsg := "Agent CustomNodeLabels are only supported for DCOS and Kubernetes"
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
		expectedMsg := fmt.Sprintf("singlePlacementGroup is only supported with VirtualMachineScaleSets")
		if err := cs.Properties.validateAgentPoolProfiles(true); err.Error() != expectedMsg {
			t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
		}
	})

	t.Run("Should fail for AvailabilitySet + SinglePlacementGroup false", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		agentPoolProfiles := cs.Properties.AgentPoolProfiles
		agentPoolProfiles[0].SinglePlacementGroup = to.BoolPtr(false)
		expectedMsg := fmt.Sprintf("singlePlacementGroup is only supported with VirtualMachineScaleSets")
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

	t.Run("Should fail for invalid VMSS + VnetSubnetID + FirstConsecutiveStaticIP config", func(t *testing.T) {
		t.Parallel()
		cs := getK8sDefaultContainerService(false)
		cs.Properties.MasterProfile.AvailabilityProfile = VirtualMachineScaleSets
		cs.Properties.MasterProfile.VnetSubnetID = "vnet"
		cs.Properties.MasterProfile.FirstConsecutiveStaticIP = "10.10.10.240"
		expectedMsg := fmt.Sprintf("when masterProfile's availabilityProfile is VirtualMachineScaleSets and a vnetSubnetID is specified, the firstConsecutiveStaticIP should be empty and will be determined by an offset from the first IP in the vnetCidr")
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
		expectedMsg := fmt.Sprintf("VirtualMachineScaleSets for master profile must be used together with virtualMachineScaleSets for agent profiles. Set \"availabilityProfile\" to \"VirtualMachineScaleSets\" for agent profiles")
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
		expectedMsg := fmt.Sprintf("mixed mode availability profiles are not allowed. Please set either VirtualMachineScaleSets or AvailabilitySet in availabilityProfile for all agent pools")
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
			case RHEL, CoreOS:
				expectedMsg := fmt.Sprintf("You have enabled auditd in agent pool %s, but you did not specify an Ubuntu-based distro", agentPoolProfiles[0].Name)
				if err := cs.Properties.validateAgentPoolProfiles(false); err.Error() != expectedMsg {
					t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
				}
			case Ubuntu, Ubuntu1804, AKSUbuntu1604, AKSUbuntu1804, ACC1604:
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
			case RHEL, CoreOS:
				expectedMsg := fmt.Sprintf("You have enabled auditd for master vms, but you did not specify an Ubuntu-based distro.")
				if err := cs.Properties.validateMasterProfile(false); err.Error() != expectedMsg {
					t.Errorf("expected error with message : %s, but got %s", expectedMsg, err.Error())
				}
			case Ubuntu, Ubuntu1804, AKSUbuntu1604, AKSUbuntu1804, ACC1604:
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
			name: "PortalURL is empty",
			cs: &ContainerService{
				Location: "testlocation",
				Properties: &Properties{
					CustomCloudProfile: &CustomCloudProfile{
						AuthenticationMethod: "azure_ad",
					},
				},
			},
			expectedErr: errors.New("portalURL needs to be specified when CustomCloudProfile is provided"),
		},
		{
			name: "PortalURL is invalid",
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
						OrchestratorVersion: "1.13.11",
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
						OrchestratorVersion: "1.13.11",
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
						OrchestratorVersion: "1.13.11",
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
						OrchestratorVersion: "1.13.11",
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
						OrchestratorVersion: "1.13.11",
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
						OrchestratorVersion: "1.15.9",
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
		"should error when masterProfile includes both an ImageRef and a Distro configuration": {
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
			expectedError: errors.New("masterProfile includes a custom image configuration (imageRef) and an explicit distro configuration, you may use one of these but not both simultaneously"),
		},
		"should error when masterProfile includes both an ImageRef and a Distro configuration in update context": {
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
			expectedError: errors.New("masterProfile includes a custom image configuration (imageRef) and an explicit distro configuration, you may use one of these but not both simultaneously"),
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
		"should error when AgentPoolProfile includes both an ImageRef and a Distro configuration": {
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
			expectedError: errors.Errorf("agentPoolProfile %s includes a custom image configuration (imageRef) and an explicit distro configuration, you may use one of these but not both simultaneously", "foo"),
		},
		"should error when AgentPoolProfile includes both an ImageRef and a Distro configuration in update context": {
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
			expectedError: errors.Errorf("agentPoolProfile %s includes a custom image configuration (imageRef) and an explicit distro configuration, you may use one of these but not both simultaneously", "foo"),
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
			if err := cs.Validate(false); !helpers.EqualError(err, test.expectedErr) {
				t.Logf("scenario %q", test.name)
				t.Errorf("expected error: %v, got: %v", test.expectedErr, err)
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
