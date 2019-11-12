// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"strconv"

	"github.com/Azure/aks-engine/pkg/api/common"
)

func (cs *ContainerService) setCloudControllerManagerConfig() {
	o := cs.Properties.OrchestratorProfile
	staticCloudControllerManagerConfig := map[string]string{
		"--allocate-node-cidrs":         strconv.FormatBool(!o.IsAzureCNI()),
		"--configure-cloud-routes":      strconv.FormatBool(o.RequireRouteTable()),
		"--cloud-provider":              "azure",
		"--cloud-config":                "/etc/kubernetes/azure.json",
		"--cluster-cidr":                o.KubernetesConfig.ClusterSubnet,
		"--kubeconfig":                  "/var/lib/kubelet/kubeconfig",
		"--leader-elect":                "true",
		"--route-reconciliation-period": "10s",
		"--v":                           "2",
	}

	// Add new arguments for Azure cloud-controller-manager component.
	if common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.16.0") {
		staticCloudControllerManagerConfig["--controllers"] = "*"
	}

	// Set --cluster-name based on appropriate DNS prefix
	if cs.Properties.MasterProfile != nil {
		staticCloudControllerManagerConfig["--cluster-name"] = cs.Properties.MasterProfile.DNSPrefix
	} else if cs.Properties.HostedMasterProfile != nil {
		staticCloudControllerManagerConfig["--cluster-name"] = cs.Properties.HostedMasterProfile.DNSPrefix
	}

	// Default cloud-controller-manager config
	defaultCloudControllerManagerConfig := map[string]string{
		"--route-reconciliation-period": DefaultKubernetesCtrlMgrRouteReconciliationPeriod,
	}

	// If no user-configurable cloud-controller-manager config values exists, use the defaults
	if o.KubernetesConfig.CloudControllerManagerConfig == nil {
		o.KubernetesConfig.CloudControllerManagerConfig = defaultCloudControllerManagerConfig
	} else {
		for key, val := range defaultCloudControllerManagerConfig {
			// If we don't have a user-configurable cloud-controller-manager config for each option
			if _, ok := o.KubernetesConfig.CloudControllerManagerConfig[key]; !ok {
				// then assign the default value
				o.KubernetesConfig.CloudControllerManagerConfig[key] = val
			}
		}
	}

	// We don't support user-configurable values for the following,
	// so any of the value assignments below will override user-provided values
	for key, val := range staticCloudControllerManagerConfig {
		o.KubernetesConfig.CloudControllerManagerConfig[key] = val
	}

	// TODO add RBAC support
	/*if *o.KubernetesConfig.EnableRbac {
		o.KubernetesConfig.CloudControllerManagerConfig["--use-service-account-credentials"] = "true"
	}*/
}
