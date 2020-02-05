// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func (cs *ContainerService) setComponentsConfig(isUpgrade bool) {
	if cs == nil || cs.Properties == nil || cs.Properties.OrchestratorProfile == nil || cs.Properties.OrchestratorProfile.KubernetesConfig == nil {
		return
	}
	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig
	defaultSchedulerComponentConfig := KubernetesComponent{
		Name:    common.SchedulerComponentName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:  common.SchedulerComponentName,
				Image: getContainerImage(common.SchedulerComponentName, cs),
			},
		},
		Config: map[string]string{
			"command": getSchedulerDefaultCommandString(cs),
		},
	}

	defaultControllerManagerComponentConfig := KubernetesComponent{
		Name:    common.ControllerManagerComponentName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:  common.ControllerManagerComponentName,
				Image: getContainerImage(common.ControllerManagerComponentName, cs),
			},
		},
		Config: map[string]string{
			"command": getControllerManagerDefaultCommandString(cs),
		},
	}

	defaultCloudControllerManagerComponentConfig := KubernetesComponent{
		Name:    common.CloudControllerManagerComponentName,
		Enabled: to.BoolPtr(to.Bool(kubernetesConfig.UseCloudControllerManager)),
		Containers: []KubernetesContainerSpec{
			{
				Name:  common.CloudControllerManagerComponentName,
				Image: getContainerImage(common.CloudControllerManagerComponentName, cs),
			},
		},
		Config: map[string]string{
			"command": "\"cloud-controller-manager\"",
		},
	}

	defaultAPIServerComponentConfig := KubernetesComponent{
		Name:    common.APIServerComponentName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:  common.APIServerComponentName,
				Image: getContainerImage(common.APIServerComponentName, cs),
			},
		},
		Config: map[string]string{
			"command": getAPIServerDefaultCommandString(cs),
		},
	}

	defaultAddonManagerComponentConfig := KubernetesComponent{
		Name:    common.AddonManagerComponentName,
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:  common.AddonManagerComponentName,
				Image: getContainerImage(common.AddonManagerComponentName, cs),
			},
		},
	}

	defaultComponents := []KubernetesComponent{
		defaultSchedulerComponentConfig,
		defaultControllerManagerComponentConfig,
		defaultCloudControllerManagerComponentConfig,
		defaultAPIServerComponentConfig,
		defaultAddonManagerComponentConfig,
	}
	// Add default component specification, if no user-provided spec exists
	if kubernetesConfig.Components == nil {
		kubernetesConfig.Components = defaultComponents
	} else {
		for _, component := range defaultComponents {
			kubernetesConfig.Components = appendComponentIfNotPresent(kubernetesConfig.Components, component)
		}
	}

	for _, component := range defaultComponents {
		synthesizeComponentsConfig(kubernetesConfig.Components, component, isUpgrade)
	}
}

func appendComponentIfNotPresent(components []KubernetesComponent, component KubernetesComponent) []KubernetesComponent {
	if component.Name != "" {
		i := getComponentsIndexByName(components, component.Name)
		if i < 0 {
			return append(components, component)
		}
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
func assignDefaultComponentVals(component, defaultComponent KubernetesComponent, isUpgrade bool) KubernetesComponent {
	if component.Name == "" {
		component.Name = defaultComponent.Name
	}
	if component.Enabled == nil {
		component.Enabled = defaultComponent.Enabled
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
	for i := range defaultComponent.Containers {
		c := component.GetContainersIndexByName(defaultComponent.Containers[i].Name)
		if c < 0 {
			component.Containers = append(component.Containers, defaultComponent.Containers[i])
		} else {
			if component.Containers[c].Image == "" || isUpgrade {
				component.Containers[c].Image = defaultComponent.Containers[i].Image
			}
			if component.Containers[c].CPURequests == "" {
				component.Containers[c].CPURequests = defaultComponent.Containers[i].CPURequests
			}
			if component.Containers[c].MemoryRequests == "" {
				component.Containers[c].MemoryRequests = defaultComponent.Containers[i].MemoryRequests
			}
			if component.Containers[c].CPULimits == "" {
				component.Containers[c].CPULimits = defaultComponent.Containers[i].CPULimits
			}
			if component.Containers[c].MemoryLimits == "" {
				component.Containers[c].MemoryLimits = defaultComponent.Containers[i].MemoryLimits
			}
		}
	}
	for key, val := range defaultComponent.Config {
		if component.Config == nil {
			component.Config = make(map[string]string)
		}
		if v, ok := component.Config[key]; !ok || v == "" {
			component.Config[key] = val
		}
	}
	return component
}

func synthesizeComponentsConfig(components []KubernetesComponent, defaultComponent KubernetesComponent, isUpgrade bool) {
	i := getComponentsIndexByName(components, defaultComponent.Name)
	if i >= 0 {
		components[i] = assignDefaultComponentVals(components[i], defaultComponent, isUpgrade)
	}
}

func getAPIServerDefaultCommandString(cs *ContainerService) string {
	if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
		return fmt.Sprintf("\"kube-apiserver\"")
	} else {
		return fmt.Sprintf("\"/hyperkube\", \"kube-apiserver\"")
	}
}

func getControllerManagerDefaultCommandString(cs *ContainerService) string {
	if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
		return fmt.Sprintf("\"kube-controller-manager\"")
	} else {
		return fmt.Sprintf("\"/hyperkube\", \"kube-controller-manager\"")
	}
}

func getSchedulerDefaultCommandString(cs *ContainerService) string {
	if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
		return fmt.Sprintf("\"kube-scheduler\"")
	} else {
		return fmt.Sprintf("\"/hyperkube\", \"kube-scheduler\"")
	}
}

func getContainerImage(component string, cs *ContainerService) string {
	if cs == nil || cs.Properties == nil || cs.Properties.OrchestratorProfile == nil || cs.Properties.OrchestratorProfile.KubernetesConfig == nil {
		return ""
	}
	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig
	cloudSpecConfig := cs.GetCloudSpecConfig()
	k8sComponents := K8sComponentsByVersionMap[cs.Properties.OrchestratorProfile.OrchestratorVersion]
	specConfig := cloudSpecConfig.KubernetesSpecConfig
	hyperkubeImageBase := specConfig.KubernetesImageBase
	hyperkubeImage := hyperkubeImageBase + k8sComponents[common.Hyperkube]
	kubernetesImageBase := kubernetesConfig.KubernetesImageBase
	if cs.Properties.IsAzureStackCloud() {
		kubernetesImageBase = cs.GetCloudSpecConfig().KubernetesSpecConfig.KubernetesImageBase
	}
	if cs.Properties.IsAzureStackCloud() {
		hyperkubeImage = hyperkubeImage + common.AzureStackSuffix
	}
	if cs.Properties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage != "" {
		hyperkubeImage = cs.Properties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage
	}
	controllerManagerBase := kubernetesImageBase
	if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
		controllerManagerBase = cs.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase
	}
	ccmImage := controllerManagerBase + k8sComponents[common.CloudControllerManagerComponentName]
	if cs.Properties.OrchestratorProfile.KubernetesConfig.CustomCcmImage != "" {
		ccmImage = cs.Properties.OrchestratorProfile.KubernetesConfig.CustomCcmImage
	}

	switch component {
	case common.APIServerComponentName:
		if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
			return kubernetesImageBase + k8sComponents[common.APIServerComponentName]
		} else {
			return hyperkubeImage
		}
	case common.ControllerManagerComponentName:
		if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
			return kubernetesImageBase + k8sComponents[common.ControllerManagerComponentName]
		} else {
			return hyperkubeImage
		}
	case common.CloudControllerManagerComponentName:
		return ccmImage
	case common.SchedulerComponentName:
		if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
			return kubernetesImageBase + k8sComponents[common.SchedulerComponentName]
		} else {
			return hyperkubeImage
		}
	case common.AddonManagerComponentName:
		return kubernetesImageBase + k8sComponents[common.AddonManagerComponentName]
	default:
		return ""
	}
}
