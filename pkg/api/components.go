// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func (cs *ContainerService) setComponentsConfig(isUpgrade bool) {
	o := cs.Properties.OrchestratorProfile
	/*
		cloudSpecConfig := cs.GetCloudSpecConfig()
		k8sComponents := K8sComponentsByVersionMap[o.OrchestratorVersion]
		specConfig := cloudSpecConfig.KubernetesSpecConfig
	*/

	defaultSchedulerComponentConfig := KubernetesComponent{
		Name:    common.SchedulerComponentName,
		Enabled: to.BoolPtr(true),
	}

	defaultControllerManagerComponentConfig := KubernetesComponent{
		Name:    common.ControllerManagerComponentName,
		Enabled: to.BoolPtr(true),
	}

	defaultCloudControllerManagerComponentConfig := KubernetesComponent{
		Name:    common.CloudControllerManagerComponentName,
		Enabled: o.KubernetesConfig.UseCloudControllerManager,
	}

	defaultAPIServerComponentConfig := KubernetesComponent{
		Name:    common.APIServerComponentName,
		Enabled: to.BoolPtr(true),
	}

	defaultAddonManagerComponentConfig := KubernetesComponent{
		Name:    common.AddonManagerComponentName,
		Enabled: to.BoolPtr(true),
	}

	defaultComponents := []KubernetesComponent{
		defaultSchedulerComponentConfig,
		defaultControllerManagerComponentConfig,
		defaultCloudControllerManagerComponentConfig,
		defaultAPIServerComponentConfig,
		defaultAddonManagerComponentConfig,
	}
	// Add default component specification, if no user-provided spec exists
	if o.KubernetesConfig.Components == nil {
		o.KubernetesConfig.Components = defaultComponents
	} else {
		for _, component := range defaultComponents {
			o.KubernetesConfig.Components = appendComponentIfNotPresent(o.KubernetesConfig.Components, component)
		}
	}

	for _, component := range defaultComponents {
		synthesizeComponentsConfig(o.KubernetesConfig.Components, component, isUpgrade)
	}
}

func appendComponentIfNotPresent(components []KubernetesComponent, component KubernetesComponent) []KubernetesComponent {
	i := getComponentsIndexByName(components, component.Name)
	if i < 0 {
		return append(components, component)
	}
	return components
}

func getComponentsIndexByName(components []KubernetesComponent, name string) int {
	for i := range components {
		if components[i].Name == name {
			return i
		}
	}
	return -1
}

// assignDefaultComponentVals will assign default values to component from defaults, for each property in component that has a zero value
func assignDefaultComponentVals(component, defaults KubernetesComponent, isUpgrade bool) KubernetesComponent {
	if component.Enabled == nil {
		component.Enabled = defaults.Enabled
	}
	if !to.Bool(component.Enabled) {
		return KubernetesComponent{
			Name:    component.Name,
			Enabled: component.Enabled,
		}
	}
	if component.Data != "" {
		return KubernetesComponent{
			Name:    component.Name,
			Enabled: component.Enabled,
			Data:    component.Data,
		}
	}
	for i := range defaults.Containers {
		c := component.GetContainersIndexByName(defaults.Containers[i].Name)
		if c < 0 {
			component.Containers = append(component.Containers, defaults.Containers[i])
		} else {
			if component.Containers[c].Image == "" || isUpgrade {
				component.Containers[c].Image = defaults.Containers[i].Image
			}
			if component.Containers[c].CPURequests == "" {
				component.Containers[c].CPURequests = defaults.Containers[i].CPURequests
			}
			if component.Containers[c].MemoryRequests == "" {
				component.Containers[c].MemoryRequests = defaults.Containers[i].MemoryRequests
			}
			if component.Containers[c].CPULimits == "" {
				component.Containers[c].CPULimits = defaults.Containers[i].CPULimits
			}
			if component.Containers[c].MemoryLimits == "" {
				component.Containers[c].MemoryLimits = defaults.Containers[i].MemoryLimits
			}
		}
	}
	for key, val := range defaults.Config {
		if component.Config == nil {
			component.Config = make(map[string]string)
		}
		if v, ok := component.Config[key]; !ok || v == "" {
			component.Config[key] = val
		}
	}
	return component
}

func synthesizeComponentsConfig(components []KubernetesComponent, component KubernetesComponent, isUpgrade bool) {
	i := getComponentsIndexByName(components, component.Name)
	if i >= 0 {
		components[i] = assignDefaultComponentVals(components[i], component, isUpgrade)
	}
}
