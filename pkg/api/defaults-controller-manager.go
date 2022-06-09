// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func (cs *ContainerService) setControllerManagerConfig() {
	o := cs.Properties.OrchestratorProfile
	isAzureCNIDualStack := cs.Properties.IsAzureCNIDualStack()
	clusterCidr := o.KubernetesConfig.ClusterSubnet
	if isAzureCNIDualStack {
		clusterSubnets := strings.Split(clusterCidr, ",")
		if len(clusterSubnets) > 1 {
			clusterCidr = clusterSubnets[1]
		}
	}
	staticControllerManagerConfig := map[string]string{
		"--kubeconfig":                       "/var/lib/kubelet/kubeconfig",
		"--allocate-node-cidrs":              strconv.FormatBool(!o.IsAzureCNI() || isAzureCNIDualStack),
		"--configure-cloud-routes":           strconv.FormatBool(cs.Properties.RequireRouteTable()),
		"--cluster-cidr":                     clusterCidr,
		"--root-ca-file":                     "/etc/kubernetes/certs/ca.crt",
		"--cluster-signing-cert-file":        "/etc/kubernetes/certs/ca.crt",
		"--cluster-signing-key-file":         "/etc/kubernetes/certs/ca.key",
		"--service-account-private-key-file": "/etc/kubernetes/certs/apiserver.key",
		"--leader-elect":                     "true",
		"--v":                                "2",
		"--controllers":                      "*,bootstrapsigner,tokencleaner",
	}

	// Set --cluster-name based on appropriate DNS prefix
	if cs.Properties.MasterProfile != nil {
		staticControllerManagerConfig["--cluster-name"] = cs.Properties.MasterProfile.DNSPrefix
	}

	// Enable cloudprovider if we're not using cloud controller manager
	if !to.Bool(o.KubernetesConfig.UseCloudControllerManager) {
		staticControllerManagerConfig["--cloud-provider"] = "azure"
		staticControllerManagerConfig["--cloud-config"] = "/etc/kubernetes/azure.json"
	} else {
		staticControllerManagerConfig["--cloud-provider"] = "external"
	}

	ctrlMgrNodeMonitorGracePeriod := DefaultKubernetesCtrlMgrNodeMonitorGracePeriod
	ctrlMgrPodEvictionTimeout := DefaultKubernetesCtrlMgrPodEvictionTimeout
	ctrlMgrRouteReconciliationPeriod := DefaultKubernetesCtrlMgrRouteReconciliationPeriod

	if cs.Properties.IsAzureStackCloud() {
		ctrlMgrNodeMonitorGracePeriod = DefaultAzureStackKubernetesCtrlMgrNodeMonitorGracePeriod
		ctrlMgrPodEvictionTimeout = DefaultAzureStackKubernetesCtrlMgrPodEvictionTimeout
		ctrlMgrRouteReconciliationPeriod = DefaultAzureStackKubernetesCtrlMgrRouteReconciliationPeriod
	}

	// Default controller-manager config
	defaultControllerManagerConfig := map[string]string{
		"--node-monitor-grace-period":       ctrlMgrNodeMonitorGracePeriod,
		"--pod-eviction-timeout":            ctrlMgrPodEvictionTimeout,
		"--route-reconciliation-period":     ctrlMgrRouteReconciliationPeriod,
		"--terminated-pod-gc-threshold":     DefaultKubernetesCtrlMgrTerminatedPodGcThreshold,
		"--use-service-account-credentials": DefaultKubernetesCtrlMgrUseSvcAccountCreds,
		"--profiling":                       DefaultKubernetesCtrMgrEnableProfiling,
	}

	// If no user-configurable controller-manager config values exists, use the defaults
	if o.KubernetesConfig.ControllerManagerConfig == nil {
		o.KubernetesConfig.ControllerManagerConfig = defaultControllerManagerConfig
	} else {
		for key, val := range defaultControllerManagerConfig {
			// If we don't have a user-configurable controller-manager config for each option
			if _, ok := o.KubernetesConfig.ControllerManagerConfig[key]; !ok {
				// then assign the default value
				o.KubernetesConfig.ControllerManagerConfig[key] = val
			}
		}
	}

	// Enables Node Exclusion from Services (toggled on agent nodes by the alpha.service-controller.kubernetes.io/exclude-balancer label).
	// ServiceNodeExclusion feature gate is GA in 1.19, removed in 1.22 (xref: https://github.com/kubernetes/kubernetes/pull/100776)
	if !common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.19.0") {
		addDefaultFeatureGates(o.KubernetesConfig.ControllerManagerConfig, o.OrchestratorVersion, "1.9.0", "ServiceNodeExclusion=true")
	}

	// Enable the consumption of local ephemeral storage and also the sizeLimit property of an emptyDir volume.
	addDefaultFeatureGates(o.KubernetesConfig.ControllerManagerConfig, o.OrchestratorVersion, "1.10.0", "LocalStorageCapacityIsolation=true")

	// Enable legacy service account token autogeneration for v1.24.0
	addDefaultFeatureGates(o.KubernetesConfig.ControllerManagerConfig, o.OrchestratorVersion, "1.24.0", "LegacyServiceAccountTokenNoAutoGeneration=false")

	// We don't support user-configurable values for the following,
	// so any of the value assignments below will override user-provided values
	for key, val := range staticControllerManagerConfig {
		o.KubernetesConfig.ControllerManagerConfig[key] = val
	}

	if o.KubernetesConfig.IsRBACEnabled() {
		o.KubernetesConfig.ControllerManagerConfig["--use-service-account-credentials"] = "true"
	}

	invalidFeatureGates := []string{}
	// Remove --feature-gate VolumeSnapshotDataSource starting with 1.22
	if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.22.0-alpha.1") {
		invalidFeatureGates = append(invalidFeatureGates, "VolumeSnapshotDataSource")
	}
	removeInvalidFeatureGates(o.KubernetesConfig.ControllerManagerConfig, invalidFeatureGates)
}
