// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"fmt"
	"net"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/openshift/certgen/release39"
	"github.com/Azure/aks-engine/pkg/openshift/certgen/unstable"
)

// setOpenShiftSetDefaultCerts sets default certificate and configuration properties in the
// openshift orchestrator.
func setOpenShiftSetDefaultCerts(a *Properties, orchestratorName, clusterID string) (bool, []net.IP, error) {
	if len(a.OrchestratorProfile.OpenShiftConfig.ConfigBundles["master"]) > 0 &&
		len(a.OrchestratorProfile.OpenShiftConfig.ConfigBundles["bootstrap"]) > 0 {
		return true, nil, nil
	}
	if a.OrchestratorProfile.OpenShiftConfig.ConfigBundles == nil {
		a.OrchestratorProfile.OpenShiftConfig.ConfigBundles = make(map[string][]byte)
	}

	var err error

	var masterBundle, nodeBundle []byte

	switch a.OrchestratorProfile.OrchestratorVersion {
	case common.OpenShiftVersion3Dot9Dot0:
		c := createR39Config(a, orchestratorName, clusterID)
		masterBundle, nodeBundle, err = release39.OpenShiftSetDefaultCerts(c)
	default:
		c := createUnstableReleaseConfig(a, orchestratorName, clusterID)
		masterBundle, nodeBundle, err = unstable.OpenShiftSetDefaultCerts(c)
	}

	if err != nil {
		return false, nil, err
	}

	a.OrchestratorProfile.OpenShiftConfig.ConfigBundles["master"] = masterBundle
	a.OrchestratorProfile.OpenShiftConfig.ConfigBundles["bootstrap"] = nodeBundle

	return true, nil, nil
}

func createR39Config(a *Properties, orchestratorName, clusterID string) *release39.Config {
	return &release39.Config{
		Master: &release39.Master{
			Hostname: fmt.Sprintf("%s-master-%s-0", orchestratorName, clusterID),
			IPs: []net.IP{
				net.ParseIP(a.MasterProfile.FirstConsecutiveStaticIP),
			},
			Port: 8443,
		},
		ExternalMasterHostname:  fmt.Sprintf("%s.%s.cloudapp.azure.com", a.MasterProfile.DNSPrefix, a.AzProfile.Location),
		ClusterUsername:         a.OrchestratorProfile.OpenShiftConfig.ClusterUsername,
		ClusterPassword:         a.OrchestratorProfile.OpenShiftConfig.ClusterPassword,
		EnableAADAuthentication: a.OrchestratorProfile.OpenShiftConfig.EnableAADAuthentication,
		AzureConfig: release39.AzureConfig{
			TenantID:                   a.AzProfile.TenantID,
			SubscriptionID:             a.AzProfile.SubscriptionID,
			AADClientID:                a.ServicePrincipalProfile.ClientID,
			AADClientSecret:            a.ServicePrincipalProfile.Secret,
			ResourceGroup:              a.AzProfile.ResourceGroup,
			Location:                   a.AzProfile.Location,
			SecurityGroupName:          fmt.Sprintf("%s-master-%s-nsg", orchestratorName, clusterID),
			PrimaryAvailabilitySetName: fmt.Sprintf("compute-availabilityset-%s", clusterID),
		},
	}
}

func createUnstableReleaseConfig(a *Properties, orchestratorName, clusterID string) *unstable.Config {
	return &unstable.Config{
		Master: &unstable.Master{
			Hostname: fmt.Sprintf("%s-master-%s-0", orchestratorName, clusterID),
			IPs: []net.IP{
				net.ParseIP(a.MasterProfile.FirstConsecutiveStaticIP),
			},
			Port: 8443,
		},
		ExternalMasterHostname:  fmt.Sprintf("%s.%s.cloudapp.azure.com", a.MasterProfile.DNSPrefix, a.AzProfile.Location),
		ClusterUsername:         a.OrchestratorProfile.OpenShiftConfig.ClusterUsername,
		ClusterPassword:         a.OrchestratorProfile.OpenShiftConfig.ClusterPassword,
		EnableAADAuthentication: a.OrchestratorProfile.OpenShiftConfig.EnableAADAuthentication,
		AzureConfig: unstable.AzureConfig{
			TenantID:                   a.AzProfile.TenantID,
			SubscriptionID:             a.AzProfile.SubscriptionID,
			AADClientID:                a.ServicePrincipalProfile.ClientID,
			AADClientSecret:            a.ServicePrincipalProfile.Secret,
			ResourceGroup:              a.AzProfile.ResourceGroup,
			Location:                   a.AzProfile.Location,
			SecurityGroupName:          fmt.Sprintf("%s-master-%s-nsg", orchestratorName, clusterID),
			PrimaryAvailabilitySetName: fmt.Sprintf("compute-availabilityset-%s", clusterID),
		},
	}
}
