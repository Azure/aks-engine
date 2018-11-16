// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v20170831

import (
	"github.com/pkg/errors"
)

// Merge existing ManagedCluster attribute into mc
func (mc *ManagedCluster) Merge(emc *ManagedCluster) error {

	// Merge properties.dnsPrefix
	if emc.Properties.DNSPrefix == "" {
		return errors.New("existing ManagedCluster expect properties.dnsPrefix not to be empty")
	}

	if mc.Properties.DNSPrefix == "" {
		mc.Properties.DNSPrefix = emc.Properties.DNSPrefix
	} else if mc.Properties.DNSPrefix != emc.Properties.DNSPrefix {
		return errors.Errorf("change dnsPrefix from %s to %s is not supported",
			emc.Properties.DNSPrefix,
			mc.Properties.DNSPrefix)
	}
	// Merge Properties.AgentPoolProfiles
	if mc.Properties.AgentPoolProfiles == nil {
		mc.Properties.AgentPoolProfiles = emc.Properties.AgentPoolProfiles
	}
	// Merge Properties.LinuxProfile
	if mc.Properties.LinuxProfile == nil {
		mc.Properties.LinuxProfile = emc.Properties.LinuxProfile
	}
	// Merge Properties.WindowsProfile
	if mc.Properties.WindowsProfile == nil {
		mc.Properties.WindowsProfile = emc.Properties.WindowsProfile
	}
	// Merge Properties.ServicePrincipalProfile
	if mc.Properties.ServicePrincipalProfile == nil {
		mc.Properties.ServicePrincipalProfile = emc.Properties.ServicePrincipalProfile
	}
	// Merge Properties.KubernetesVersion
	if mc.Properties.KubernetesVersion == "" {
		mc.Properties.KubernetesVersion = emc.Properties.KubernetesVersion
	}

	return nil
}
