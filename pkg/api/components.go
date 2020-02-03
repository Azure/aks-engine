// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func (cs *ContainerService) setComponentsConfig(isUpgrade bool) {
	o := cs.Properties.OrchestratorProfile
	cloudSpecConfig := cs.GetCloudSpecConfig()
	k8sComponents := K8sComponentsByVersionMap[o.OrchestratorVersion]
	specConfig := cloudSpecConfig.KubernetesSpecConfig

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
		Enabled: to.BoolPtr(to.Bool(o.KubernetesConfig.UseCloudControllerManager)),
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
				Image: specConfig.KubernetesImageBase + k8sComponents[common.AddonManagerComponentName],
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
	o := cs.Properties.OrchestratorProfile
	cloudSpecConfig := cs.GetCloudSpecConfig()
	k8sComponents := K8sComponentsByVersionMap[o.OrchestratorVersion]
	specConfig := cloudSpecConfig.KubernetesSpecConfig
	hyperkubeImageBase := specConfig.KubernetesImageBase
	hyperkubeImage := hyperkubeImageBase + k8sComponents["hyperkube"]
	kubernetesImageBase := o.KubernetesConfig.KubernetesImageBase
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
	default:
		return ""
	}
}
