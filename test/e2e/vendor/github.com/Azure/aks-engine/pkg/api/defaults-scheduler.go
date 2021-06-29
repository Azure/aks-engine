// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

// staticSchedulerConfig is not user-overridable
var staticSchedulerConfig = map[string]string{
	"--kubeconfig":   "/var/lib/kubelet/kubeconfig",
	"--leader-elect": "true",
}

// defaultSchedulerConfig provides targeted defaults, but is user-overridable
var defaultSchedulerConfig = map[string]string{
	"--v":         "2",
	"--profiling": DefaultKubernetesSchedulerEnableProfiling,
}

func (cs *ContainerService) setSchedulerConfig() {
	o := cs.Properties.OrchestratorProfile

	// If no user-configurable scheduler config values exists, make an empty map, and fill in with defaults
	if o.KubernetesConfig.SchedulerConfig == nil {
		o.KubernetesConfig.SchedulerConfig = make(map[string]string)
	}

	for key, val := range defaultSchedulerConfig {
		// If we don't have a user-configurable scheduler config for each option
		if _, ok := o.KubernetesConfig.SchedulerConfig[key]; !ok {
			// then assign the default value
			o.KubernetesConfig.SchedulerConfig[key] = val
		}
	}

	// We don't support user-configurable values for the following,
	// so any of the value assignments below will override user-provided values
	for key, val := range staticSchedulerConfig {
		o.KubernetesConfig.SchedulerConfig[key] = val
	}
}
