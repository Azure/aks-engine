// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

func (cs *ContainerService) setSysctlDConfig() {
	// Default aks-engine-provided sysctl.d config
	defaultSysctlDConfig := map[string]string{
		"net.ipv4.tcp_retries2":             "8",
		"net.core.somaxconn":                "16384",
		"net.ipv4.tcp_max_syn_backlog":      "16384",
		"net.core.message_cost":             "40",
		"net.core.message_burst":            "80",
		"net.ipv4.neigh.default.gc_thresh1": "4096",
		"net.ipv4.neigh.default.gc_thresh2": "8192",
		"net.ipv4.neigh.default.gc_thresh3": "16384",
		"kernel.dmesg_restrict":             "0",
	}

	if cs.Properties.OrchestratorProfile.KubernetesConfig.NeedsContainerd() {
		defaultSysctlDConfig["net.ipv4.ip_forward"] = "1"
	}

	// Master-specific kubelet config changes go here
	if cs.Properties.MasterProfile != nil {
		if cs.Properties.MasterProfile.SysctlDConfig == nil {
			cs.Properties.MasterProfile.SysctlDConfig = make(map[string]string)
		}
		setMissingSysctlDConfigValues(cs.Properties.MasterProfile.SysctlDConfig, defaultSysctlDConfig)
	}

	// Agent-specific kubelet config changes go here
	for _, profile := range cs.Properties.AgentPoolProfiles {
		if profile.IsLinux() {
			if profile.SysctlDConfig == nil {
				profile.SysctlDConfig = make(map[string]string)
			}

			setMissingSysctlDConfigValues(profile.SysctlDConfig, defaultSysctlDConfig)
		}
	}
}

func setMissingSysctlDConfigValues(sysctlDConfig map[string]string, defaults map[string]string) {
	for key, val := range defaults {
		// If we don't have a user-configurable value for each option
		if _, ok := sysctlDConfig[key]; !ok {
			// then assign the default value
			sysctlDConfig[key] = val
		}
	}
}
