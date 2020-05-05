// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"reflect"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
)

const (
	customCcmImage                   = "my-custom-cloud-controller-manager-image"
	customHyperkubeImage             = "my-custom-hyperkube-image"
	customKubeAddonManagerImage      = "my-custom-kube-addon-manager-image"
	customKubeAPIServerImage         = "my-custom-kube-apiserver-image"
	customKubeControllerManagerImage = "my-custom-kube-controller-manager-image"
	customKubeSchedulerImage         = "my-custom-kube-scheduler-image"
)

func TestSetComponentsConfig(t *testing.T) {
	userConfiguredComponentsMap := getUserConfiguredComponentMap()
	containerServiceMap := getContainerServicesMap()

	tests := []struct {
		name               string
		cs                 *ContainerService
		isUpgrade          bool
		expectedComponents []KubernetesComponent
	}{
		{
			name:               "1.13",
			cs:                 containerServiceMap["1.13"],
			isUpgrade:          false,
			expectedComponents: getDefaultComponents(containerServiceMap["1.13"]),
		},
		{
			name:      "1.13 user-configured",
			cs:        containerServiceMap["1.13 user-configured"],
			isUpgrade: false,
			expectedComponents: append(overwriteDefaultComponents([]KubernetesComponent{
				userConfiguredComponentsMap["user-configured kube-scheduler component"],
				userConfiguredComponentsMap["user-configured controller-manager component"],
				userConfiguredComponentsMap["user-configured kube-apiserver component"],
				userConfiguredComponentsMap["user-configured kube-addon-manager component"],
			}, containerServiceMap["1.13 user-configured"]), userConfiguredComponentsMap["user-configured cloud-controller-manager component"]),
		},
		{
			name:      "1.13 + customHyperkubeImage + customCcmImage",
			cs:        containerServiceMap["1.13 + customHyperkubeImage + customCcmImage"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-scheduler\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: customCcmImage,
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.13 + customHyperkubeImage + customCcmImage"]),
		},
		{
			name:               "1.14",
			cs:                 containerServiceMap["1.14"],
			isUpgrade:          false,
			expectedComponents: getDefaultComponents(containerServiceMap["1.14"]),
		},
		{
			name:      "1.14 user-configured",
			cs:        containerServiceMap["1.14 user-configured"],
			isUpgrade: false,
			expectedComponents: append(overwriteDefaultComponents([]KubernetesComponent{
				userConfiguredComponentsMap["user-configured kube-scheduler component"],
				userConfiguredComponentsMap["user-configured controller-manager component"],
				userConfiguredComponentsMap["user-configured kube-apiserver component"],
				userConfiguredComponentsMap["user-configured kube-addon-manager component"],
			}, containerServiceMap["1.14 user-configured"]), userConfiguredComponentsMap["user-configured cloud-controller-manager component"]),
		},
		{
			name:      "1.14 + customHyperkubeImage + customCcmImage",
			cs:        containerServiceMap["1.14 + customHyperkubeImage + customCcmImage"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-scheduler\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: customCcmImage,
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.14 + customHyperkubeImage + customCcmImage"]),
		},
		{
			name:               "1.15",
			cs:                 containerServiceMap["1.15"],
			isUpgrade:          false,
			expectedComponents: getDefaultComponents(containerServiceMap["1.15"]),
		},
		{
			name:      "1.15 user-configured",
			cs:        containerServiceMap["1.15 user-configured"],
			isUpgrade: false,
			expectedComponents: append(overwriteDefaultComponents([]KubernetesComponent{
				userConfiguredComponentsMap["user-configured kube-scheduler component"],
				userConfiguredComponentsMap["user-configured controller-manager component"],
				userConfiguredComponentsMap["user-configured kube-apiserver component"],
				userConfiguredComponentsMap["user-configured kube-addon-manager component"],
			}, containerServiceMap["1.15 user-configured"]), userConfiguredComponentsMap["user-configured cloud-controller-manager component"]),
		},
		{
			name:      "1.15 + customHyperkubeImage + customCcmImage",
			cs:        containerServiceMap["1.15 + customHyperkubeImage + customCcmImage"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-scheduler\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: customCcmImage,
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.15 + customHyperkubeImage + customCcmImage"]),
		},
		{
			name:               "1.16",
			cs:                 containerServiceMap["1.16"],
			isUpgrade:          false,
			expectedComponents: getDefaultComponents(containerServiceMap["1.16"]),
		},
		{
			name:      "1.16 user-configured",
			cs:        containerServiceMap["1.16 user-configured"],
			isUpgrade: false,
			expectedComponents: append(overwriteDefaultComponents([]KubernetesComponent{
				userConfiguredComponentsMap["user-configured kube-scheduler component"],
				userConfiguredComponentsMap["user-configured controller-manager component"],
				userConfiguredComponentsMap["user-configured kube-apiserver component"],
				userConfiguredComponentsMap["user-configured kube-addon-manager component"],
			}, containerServiceMap["1.16 user-configured"]), userConfiguredComponentsMap["user-configured cloud-controller-manager component"]),
		},
		{
			name:      "1.16 + customHyperkubeImage + customCcmImage",
			cs:        containerServiceMap["1.16 + customHyperkubeImage + customCcmImage"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: customHyperkubeImage,
						},
					},
					Config: map[string]string{
						"command": "\"/hyperkube\", \"kube-scheduler\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: customCcmImage,
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.16 + customHyperkubeImage + customCcmImage"]),
		},
		{
			name:               "1.17",
			cs:                 containerServiceMap["1.17"],
			isUpgrade:          false,
			expectedComponents: getDefaultComponents(containerServiceMap["1.17"]),
		},
		{
			name:      "1.17 user-configured",
			cs:        containerServiceMap["1.17 user-configured"],
			isUpgrade: false,
			expectedComponents: append(overwriteDefaultComponents([]KubernetesComponent{
				userConfiguredComponentsMap["user-configured kube-scheduler component"],
				userConfiguredComponentsMap["user-configured controller-manager component"],
				userConfiguredComponentsMap["user-configured kube-apiserver component"],
				userConfiguredComponentsMap["user-configured kube-addon-manager component"],
			}, containerServiceMap["1.17 user-configured"]), userConfiguredComponentsMap["user-configured cloud-controller-manager component"]),
		},
		{
			name:      "1.17 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage",
			cs:        containerServiceMap["1.17 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: customKubeAPIServerImage,
						},
					},
					Config: map[string]string{
						"command": "\"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: customKubeControllerManagerImage,
						},
					},
					Config: map[string]string{
						"command": "\"kube-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: customKubeSchedulerImage,
						},
					},
					Config: map[string]string{
						"command": "\"kube-scheduler\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: customCcmImage,
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.17 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage"]),
		},
		{
			name:               "1.18",
			cs:                 containerServiceMap["1.18"],
			isUpgrade:          false,
			expectedComponents: getDefaultComponents(containerServiceMap["1.18"]),
		},
		{
			name:      "1.18 user-configured",
			cs:        containerServiceMap["1.18 user-configured"],
			isUpgrade: false,
			expectedComponents: append(overwriteDefaultComponents([]KubernetesComponent{
				userConfiguredComponentsMap["user-configured kube-scheduler component"],
				userConfiguredComponentsMap["user-configured controller-manager component"],
				userConfiguredComponentsMap["user-configured kube-apiserver component"],
				userConfiguredComponentsMap["user-configured kube-addon-manager component"],
			}, containerServiceMap["1.18 user-configured"]), userConfiguredComponentsMap["user-configured cloud-controller-manager component"]),
		},
		{
			name:      "1.18 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage",
			cs:        containerServiceMap["1.18 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.APIServerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.APIServerComponentName,
							Image: customKubeAPIServerImage,
						},
					},
					Config: map[string]string{
						"command": "\"kube-apiserver\"",
					},
				},
				{
					Name:    common.ControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.ControllerManagerComponentName,
							Image: customKubeControllerManagerImage,
						},
					},
					Config: map[string]string{
						"command": "\"kube-controller-manager\"",
					},
				},
				{
					Name:    common.SchedulerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.SchedulerComponentName,
							Image: customKubeSchedulerImage,
						},
					},
					Config: map[string]string{
						"command": "\"kube-scheduler\"",
					},
				},
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: customCcmImage,
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.18 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage"]),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.cs.setComponentsConfig(test.isUpgrade)
			for _, componentName := range []string{
				common.SchedulerComponentName,
				common.ControllerManagerComponentName,
				common.CloudControllerManagerComponentName,
				common.APIServerComponentName,
				common.AddonManagerComponentName,
			} {
				component := test.cs.Properties.OrchestratorProfile.KubernetesConfig.Components[GetComponentsIndexByName(test.cs.Properties.OrchestratorProfile.KubernetesConfig.Components, componentName)]
				if component.IsEnabled() {
					if i := GetComponentsIndexByName(test.expectedComponents, componentName); i == -1 {
						t.Fatalf("got component %s that we weren't expecting", component.Name)
					}
					expectedComponent := test.expectedComponents[GetComponentsIndexByName(test.expectedComponents, componentName)]
					if to.Bool(component.Enabled) != to.Bool(expectedComponent.Enabled) {
						t.Fatalf("expected component %s to have Enabled value %t, instead got %t", expectedComponent.Name, to.Bool(expectedComponent.Enabled), to.Bool(component.Enabled))
					}
					if expectedComponent.Containers != nil {
						if len(expectedComponent.Containers) != len(component.Containers) {
							t.Fatalf("expected component %s to have %d containers , got %d", expectedComponent.Name, len(expectedComponent.Containers), len(component.Containers))
						}
						for i, container := range expectedComponent.Containers {
							if container.Name != component.Containers[i].Name {
								t.Fatalf("expected component %s to have container Name %s at at Containers index %d, got %s", expectedComponent.Name, container.Name, i, component.Containers[i].Name)
							}
							if container.Image != component.Containers[i].Image {
								t.Fatalf("expected component %s to have container Image %s at at Containers index %d, got %s", expectedComponent.Name, container.Image, i, component.Containers[i].Image)
							}
							if container.CPURequests != component.Containers[i].CPURequests {
								t.Fatalf("expected component %s to have container CPURequests %s at at Containers index %d, got %s", expectedComponent.Name, container.CPURequests, i, component.Containers[i].CPURequests)
							}
							if container.MemoryRequests != component.Containers[i].MemoryRequests {
								t.Fatalf("expected component %s to have container MemoryRequests %s at at Containers index %d, got %s", expectedComponent.Name, container.MemoryRequests, i, component.Containers[i].MemoryRequests)
							}
							if container.CPULimits != component.Containers[i].CPULimits {
								t.Fatalf("expected component %s to have container CPULimits %s at at Containers index %d, got %s", expectedComponent.Name, container.CPULimits, i, component.Containers[i].CPULimits)
							}
							if container.MemoryLimits != component.Containers[i].MemoryLimits {
								t.Fatalf("expected component %s to have container MemoryLimits %s at at Containers index %d, got %s", expectedComponent.Name, container.MemoryLimits, i, component.Containers[i].MemoryLimits)
							}
						}
					}
					if expectedComponent.Config != nil {
						for key, val := range expectedComponent.Config {
							if val != component.Config[key] {
								t.Fatalf("expected component %s to have config %s=%s, got %s=%s", expectedComponent.Name, key, val, key, component.Config[key])
							}
						}
					}
					if component.Config != nil {
						for key, val := range component.Config {
							if val != expectedComponent.Config[key] {
								t.Fatalf("expected component %s to have config %s=%s, got %s=%s", component.Name, key, val, key, expectedComponent.Config[key])
							}
						}
					}
				} else {
					if i := GetComponentsIndexByName(test.expectedComponents, componentName); i > -1 {
						if to.Bool(test.expectedComponents[i].Enabled) {
							t.Fatalf("expected component %s to be enabled, instead it was disabled", componentName)
						}
					}
				}
			}
		})
	}
}

func TestAppendComponentIfNotPresent(t *testing.T) {
	existingComponents := []KubernetesComponent{
		{
			Name:    "i-exist",
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  "i-exist-container",
					Image: "i-exist-image",
				},
			},
			Config: map[string]string{
				"foo": "bar",
			},
		},
	}
	newComponent := KubernetesComponent{
		Name:    "i-am-new",
		Enabled: to.BoolPtr(true),
		Containers: []KubernetesContainerSpec{
			{
				Name:  "new-container",
				Image: "new-image",
			},
		},
		Config: map[string]string{
			"baz": "bang",
		},
	}
	cases := []struct {
		name               string
		existingComponents []KubernetesComponent
		newComponent       KubernetesComponent
		expectedComponents []KubernetesComponent
	}{
		{
			name:               "component not present",
			existingComponents: existingComponents,
			newComponent:       newComponent,
			expectedComponents: append(existingComponents, newComponent),
		},
		{
			name:               "existing components is empty",
			existingComponents: []KubernetesComponent{},
			newComponent:       newComponent,
			expectedComponents: []KubernetesComponent{newComponent},
		},
		{
			name:               "component is present",
			existingComponents: existingComponents,
			newComponent:       existingComponents[0],
			expectedComponents: existingComponents,
		},
		{
			name:               "empty new component",
			existingComponents: existingComponents,
			newComponent:       KubernetesComponent{},
			expectedComponents: existingComponents,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := appendComponentIfNotPresent(c.existingComponents, c.newComponent)
			if !reflect.DeepEqual(result, c.expectedComponents) {
				t.Fatalf("expected result addon %v to be equal to %v", result, c.expectedComponents)
			}
		})
	}
}

func TestGetComponentsIndexByName(t *testing.T) {
	existingComponents := make([]KubernetesComponent, 3)
	existingComponents[0] = KubernetesComponent{Name: "i-exist"}
	existingComponents[1] = KubernetesComponent{Name: "i-also-exist"}
	existingComponents[2] = KubernetesComponent{Name: "and-me-too"}
	cases := []struct {
		name          string
		components    []KubernetesComponent
		componentName string
		expected      int
	}{
		{
			name:          "component not present",
			components:    existingComponents,
			componentName: "i-do-not-exist",
			expected:      -1,
		},
		{
			name:          "index 0",
			components:    existingComponents,
			componentName: "i-exist",
			expected:      0,
		},
		{
			name:          "index 1",
			components:    existingComponents,
			componentName: "i-also-exist",
			expected:      1,
		},
		{
			name:          "index 2",
			components:    existingComponents,
			componentName: "and-me-too",
			expected:      2,
		},
		{
			name:          "empty component",
			components:    []KubernetesComponent{},
			componentName: "does-not-matter",
			expected:      -1,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := GetComponentsIndexByName(c.components, c.componentName)
			if result != c.expected {
				t.Fatalf("expected GetComponentsIndexByName() result %d to be equal to %d", result, c.expected)
			}
		})
	}
}

func TestAssignDefaultComponentVals(t *testing.T) {
	containerServiceMap := getContainerServicesMap()
	defaultOneDotFifteenComponents := getDefaultComponents(getContainerServicesMap()["1.15"])
	controllerManagerComponent := defaultOneDotFifteenComponents[GetComponentsIndexByName(defaultOneDotFifteenComponents, common.ControllerManagerComponentName)]
	cases := []struct {
		name             string
		component        KubernetesComponent
		defaultComponent KubernetesComponent
		isUpgrade        bool
		expected         KubernetesComponent
	}{
		{
			name:             "empty component",
			component:        KubernetesComponent{},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        false,
			expected:         controllerManagerComponent,
		},
		{
			name: "nil Enabled",
			component: KubernetesComponent{
				Name: common.ControllerManagerComponentName,
				Containers: []KubernetesContainerSpec{
					{
						Name:  common.ControllerManagerComponentName,
						Image: getComponentDefaultContainerImage(common.ControllerManagerComponentName, containerServiceMap["1.15"]),
					},
				},
				Config: map[string]string{
					"command": getControllerManagerDefaultCommandString(containerServiceMap["1.15"]),
				},
			},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        false,
			expected:         controllerManagerComponent,
		},
		{
			name: "disabled",
			component: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(false),
				Containers: []KubernetesContainerSpec{
					{
						Name:  common.ControllerManagerComponentName,
						Image: getComponentDefaultContainerImage(common.ControllerManagerComponentName, containerServiceMap["1.15"]),
					},
				},
				Config: map[string]string{
					"command": getControllerManagerDefaultCommandString(containerServiceMap["1.15"]),
				},
			},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        false,
			expected: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(false),
			},
		},
		{
			name: "data present",
			component: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
				Data:    "foo",
			},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        false,
			expected: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
				Data:    "foo",
			},
		},
		{
			name: "no containers or config",
			component: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
			},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        false,
			expected:         controllerManagerComponent,
		},
		{
			name: "no containers data",
			component: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name: common.ControllerManagerComponentName,
					},
				},
			},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        false,
			expected:         controllerManagerComponent,
		},
		{
			name: "additional user config",
			component: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name:           common.ControllerManagerComponentName,
						Image:          "baz",
						CPURequests:    "1",
						MemoryRequests: "200m",
						CPULimits:      "2",
						MemoryLimits:   "400m",
					},
				},
				Config: map[string]string{
					"foo": "bar",
				},
			},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        false,
			expected: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name:           common.ControllerManagerComponentName,
						Image:          "baz",
						CPURequests:    "1",
						MemoryRequests: "200m",
						CPULimits:      "2",
						MemoryLimits:   "400m",
					},
				},
				Config: map[string]string{
					"foo":     "bar",
					"command": getControllerManagerDefaultCommandString(containerServiceMap["1.15"]),
				},
			},
		},
		{
			name: "upgrade",
			component: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name:           common.ControllerManagerComponentName,
						Image:          "baz",
						CPURequests:    "1",
						MemoryRequests: "200m",
						CPULimits:      "2",
						MemoryLimits:   "400m",
					},
				},
				Config: map[string]string{
					"foo":     "bar",
					"command": "/bin/custom-command/overwrite-me",
				},
			},
			defaultComponent: controllerManagerComponent,
			isUpgrade:        true,
			expected: KubernetesComponent{
				Name:    common.ControllerManagerComponentName,
				Enabled: to.BoolPtr(true),
				Containers: []KubernetesContainerSpec{
					{
						Name:           common.ControllerManagerComponentName,
						Image:          getComponentDefaultContainerImage(common.ControllerManagerComponentName, containerServiceMap["1.15"]),
						CPURequests:    "1",
						MemoryRequests: "200m",
						CPULimits:      "2",
						MemoryLimits:   "400m",
					},
				},
				Config: map[string]string{
					"foo":     "bar",
					"command": getControllerManagerDefaultCommandString(containerServiceMap["1.15"]),
				},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			result := assignDefaultComponentVals(c.component, c.defaultComponent, c.isUpgrade)
			if !reflect.DeepEqual(result, c.expected) {
				t.Fatalf("expected assignDefaultComponentVals() result %v to be equal to %v", result, c.expected)
			}
		})
	}
}

func TestSynthesizeComponentsConfig(t *testing.T) {
	defaultOneDotFifteenComponents := getDefaultComponents(getContainerServicesMap()["1.15"])
	i := GetComponentsIndexByName(defaultOneDotFifteenComponents, common.ControllerManagerComponentName)
	defaultControllerManagerComponent := defaultOneDotFifteenComponents[i]
	customOneDotFifteenComponents := defaultOneDotFifteenComponents
	customControllerManagerComponent := KubernetesComponent{
		Name: common.ControllerManagerComponentName,
		Containers: []KubernetesContainerSpec{
			{
				Name:  common.ControllerManagerComponentName,
				Image: "my-custom-image",
			},
		},
	}
	customOneDotFifteenComponents[i] = customControllerManagerComponent
	synthesizedControllerManagerComponent := defaultControllerManagerComponent
	synthesizedControllerManagerComponent.Containers[0].Image = "my-custom-image"
	cases := []struct {
		name             string
		components       []KubernetesComponent
		defaultComponent KubernetesComponent
		isUpgrade        bool
		expected         KubernetesComponent
	}{
		{
			name:             "user override",
			components:       customOneDotFifteenComponents,
			defaultComponent: defaultControllerManagerComponent,
			isUpgrade:        false,
			expected:         synthesizedControllerManagerComponent,
		},
		{
			name:             "no user override",
			components:       defaultOneDotFifteenComponents,
			defaultComponent: defaultControllerManagerComponent,
			isUpgrade:        false,
			expected:         defaultControllerManagerComponent,
		},
		{
			name:             "upgrade",
			components:       customOneDotFifteenComponents,
			defaultComponent: defaultControllerManagerComponent,
			isUpgrade:        true,
			expected:         defaultControllerManagerComponent,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			synthesizeComponentsConfig(c.components, c.defaultComponent, c.isUpgrade)
			i := GetComponentsIndexByName(c.components, common.ControllerManagerComponentName)
			if !reflect.DeepEqual(c.components[i], c.expected) {
				t.Fatalf("expected synthesizeComponentsConfig() result %v to be equal to %v", c.components[i], c.expected)
			}
		})
	}
}

func TestGetDefaultCommandStrings(t *testing.T) {
	cases := []struct {
		name                                   string
		cs                                     *ContainerService
		expectedAPIServerCommandString         string
		expectedControllerManagerCommandString string
		expectedSchedulerCommandString         string
	}{
		{
			name:                                   "1.13",
			cs:                                     getContainerServicesMap()["1.13"],
			expectedAPIServerCommandString:         "\"/hyperkube\", \"kube-apiserver\"",
			expectedControllerManagerCommandString: "\"/hyperkube\", \"kube-controller-manager\"",
			expectedSchedulerCommandString:         "\"/hyperkube\", \"kube-scheduler\"",
		},
		{
			name:                                   "1.14",
			cs:                                     getContainerServicesMap()["1.14"],
			expectedAPIServerCommandString:         "\"/hyperkube\", \"kube-apiserver\"",
			expectedControllerManagerCommandString: "\"/hyperkube\", \"kube-controller-manager\"",
			expectedSchedulerCommandString:         "\"/hyperkube\", \"kube-scheduler\"",
		},
		{
			name:                                   "1.15",
			cs:                                     getContainerServicesMap()["1.15"],
			expectedAPIServerCommandString:         "\"/hyperkube\", \"kube-apiserver\"",
			expectedControllerManagerCommandString: "\"/hyperkube\", \"kube-controller-manager\"",
			expectedSchedulerCommandString:         "\"/hyperkube\", \"kube-scheduler\"",
		},
		{
			name:                                   "1.16",
			cs:                                     getContainerServicesMap()["1.16"],
			expectedAPIServerCommandString:         "\"/hyperkube\", \"kube-apiserver\"",
			expectedControllerManagerCommandString: "\"/hyperkube\", \"kube-controller-manager\"",
			expectedSchedulerCommandString:         "\"/hyperkube\", \"kube-scheduler\"",
		},
		{
			name:                                   "1.17",
			cs:                                     getContainerServicesMap()["1.17"],
			expectedAPIServerCommandString:         "\"kube-apiserver\"",
			expectedControllerManagerCommandString: "\"kube-controller-manager\"",
			expectedSchedulerCommandString:         "\"kube-scheduler\"",
		},
		{
			name:                                   "1.18",
			cs:                                     getContainerServicesMap()["1.18"],
			expectedAPIServerCommandString:         "\"kube-apiserver\"",
			expectedControllerManagerCommandString: "\"kube-controller-manager\"",
			expectedSchedulerCommandString:         "\"kube-scheduler\"",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			apiServerDefaultCommandString := getAPIServerDefaultCommandString(c.cs)
			if apiServerDefaultCommandString != c.expectedAPIServerCommandString {
				t.Fatalf("expected getAPIServerDefaultCommandString() result %s to be equal to %s", apiServerDefaultCommandString, c.expectedAPIServerCommandString)
			}
			controllerManagerDefaultCommandString := getControllerManagerDefaultCommandString(c.cs)
			if controllerManagerDefaultCommandString != c.expectedControllerManagerCommandString {
				t.Fatalf("expected getControllerManagerDefaultCommandString() result %s to be equal to %s", controllerManagerDefaultCommandString, c.expectedControllerManagerCommandString)
			}
			schedulerDefaultCommandString := getSchedulerDefaultCommandString(c.cs)
			if schedulerDefaultCommandString != c.expectedSchedulerCommandString {
				t.Fatalf("expected getSchedulerDefaultCommandString() result %s to be equal to %s", schedulerDefaultCommandString, c.expectedSchedulerCommandString)
			}
		})
	}
}

func TestGetContainerImages(t *testing.T) {
	specConfig := AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	csOneDotThirteen := getContainerServicesMap()["1.13"]
	csOneDotThirteenCustomImagesComponents := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	csOneDotThirteenAzureStack := getContainerServicesMap()["1.13"]
	csOneDotThirteenAzureStack.Properties.CustomCloudProfile = &CustomCloudProfile{
		Environment: &azure.Environment{
			Name: "AzureStackCloud",
		},
	}
	csOneDotThirteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "mcr.microsoft.com/azurestack/"
	csOneDotThirteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeMCR
	orchestratorVersionOneDotThirteen := csOneDotThirteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotFourteen := getContainerServicesMap()["1.14"]
	csOneDotFourteenCustomImagesComponents := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	csOneDotFourteenAzureStack := getContainerServicesMap()["1.14"]
	csOneDotFourteenAzureStack.Properties.CustomCloudProfile = &CustomCloudProfile{
		Environment: &azure.Environment{
			Name: "AzureStackCloud",
		},
	}
	csOneDotFourteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "mcr.microsoft.com/azurestack/"
	csOneDotFourteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeMCR
	orchestratorVersionOneDotFourteen := csOneDotFourteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotFifteen := getContainerServicesMap()["1.15"]
	csOneDotFifteenCustomImagesComponents := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	csOneDotFifteenAzureStack := getContainerServicesMap()["1.15"]
	csOneDotFifteenAzureStack.Properties.CustomCloudProfile = &CustomCloudProfile{
		Environment: &azure.Environment{
			Name: "AzureStackCloud",
		},
	}
	csOneDotFifteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "mcr.microsoft.com/azurestack/"
	csOneDotFifteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeMCR
	orchestratorVersionOneDotFifteen := csOneDotFifteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotSixteen := getContainerServicesMap()["1.16"]
	csOneDotSixteenCustomImagesComponents := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	csOneDotSixteenAzureStack := getContainerServicesMap()["1.16"]
	csOneDotSixteenAzureStack.Properties.CustomCloudProfile = &CustomCloudProfile{
		Environment: &azure.Environment{
			Name: "AzureStackCloud",
		},
	}
	csOneDotSixteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "mcr.microsoft.com/azurestack/"
	csOneDotSixteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeMCR
	orchestratorVersionOneDotSixteen := csOneDotSixteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotSeventeen := getContainerServicesMap()["1.17"]
	csOneDotSeventeenCustomImagesComponents := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	csOneDotSeventeenAzureStack := getContainerServicesMap()["1.17"]
	csOneDotSeventeenAzureStack.Properties.CustomCloudProfile = &CustomCloudProfile{
		Environment: &azure.Environment{
			Name: "AzureStackCloud",
		},
	}
	csOneDotSeventeenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "mcr.microsoft.com/azurestack/"
	csOneDotSeventeenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeMCR
	orchestratorVersionOneDotSeventeen := csOneDotSeventeen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotEighteen := getContainerServicesMap()["1.18"]
	csOneDotEighteenCustomImagesComponents := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	csOneDotEighteenAzureStack := getContainerServicesMap()["1.18"]
	csOneDotEighteenAzureStack.Properties.CustomCloudProfile = &CustomCloudProfile{
		Environment: &azure.Environment{
			Name: "AzureStackCloud",
		},
	}
	csOneDotEighteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = "mcr.microsoft.com/azurestack/"
	csOneDotEighteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = common.KubernetesImageBaseTypeMCR
	orchestratorVersionOneDotEighteen := csOneDotEighteen.Properties.OrchestratorProfile.OrchestratorVersion
	k8sComponentsByVersionMap := GetK8sComponentsByVersionMap(&KubernetesConfig{KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR})
	cases := []struct {
		name                                      string
		cs                                        *ContainerService
		expectedAPIServerImageString              string
		expectedControllerManagerImageString      string
		expectedCloudControllerManagerImageString string
		expectedSchedulerImageString              string
		expectedAddonManagerImageString           string
	}{
		{
			name:                                 "1.13",
			cs:                                   csOneDotThirteen,
			expectedAPIServerImageString:         specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: csOneDotThirteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.Hyperkube],
			expectedAddonManagerImageString:           csOneDotThirteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.13 Azure Stack",
			cs:                                   csOneDotThirteenAzureStack,
			expectedAPIServerImageString:         csOneDotThirteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotThirteenCustomImagesComponents[orchestratorVersionOneDotThirteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedControllerManagerImageString: csOneDotThirteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotThirteenCustomImagesComponents[orchestratorVersionOneDotThirteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedCloudControllerManagerImageString: csOneDotThirteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotThirteenCustomImagesComponents[orchestratorVersionOneDotThirteen][common.CloudControllerManagerComponentName] + common.AzureStackSuffix,
			expectedSchedulerImageString:              csOneDotThirteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotThirteenCustomImagesComponents[orchestratorVersionOneDotThirteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedAddonManagerImageString:           csOneDotThirteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotThirteenCustomImagesComponents[orchestratorVersionOneDotThirteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.14",
			cs:                                   csOneDotFourteen,
			expectedAPIServerImageString:         specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: csOneDotFourteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.Hyperkube],
			expectedAddonManagerImageString:           csOneDotFourteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.14 Azure Stack",
			cs:                                   csOneDotFourteenAzureStack,
			expectedAPIServerImageString:         csOneDotFourteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFourteenCustomImagesComponents[orchestratorVersionOneDotFourteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedControllerManagerImageString: csOneDotFourteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFourteenCustomImagesComponents[orchestratorVersionOneDotFourteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedCloudControllerManagerImageString: csOneDotFourteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFourteenCustomImagesComponents[orchestratorVersionOneDotFourteen][common.CloudControllerManagerComponentName] + common.AzureStackSuffix,
			expectedSchedulerImageString:              csOneDotFourteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFourteenCustomImagesComponents[orchestratorVersionOneDotFourteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedAddonManagerImageString:           csOneDotFourteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFourteenCustomImagesComponents[orchestratorVersionOneDotFourteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.15",
			cs:                                   csOneDotFifteen,
			expectedAPIServerImageString:         specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: csOneDotFifteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.Hyperkube],
			expectedAddonManagerImageString:           csOneDotFifteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.15 Azure Stack",
			cs:                                   csOneDotFifteenAzureStack,
			expectedAPIServerImageString:         csOneDotFifteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFifteenCustomImagesComponents[orchestratorVersionOneDotFifteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedControllerManagerImageString: csOneDotFifteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFifteenCustomImagesComponents[orchestratorVersionOneDotFifteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedCloudControllerManagerImageString: csOneDotFifteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFifteenCustomImagesComponents[orchestratorVersionOneDotFifteen][common.CloudControllerManagerComponentName] + common.AzureStackSuffix,
			expectedSchedulerImageString:              csOneDotFifteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFifteenCustomImagesComponents[orchestratorVersionOneDotFifteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedAddonManagerImageString:           csOneDotFifteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotFifteenCustomImagesComponents[orchestratorVersionOneDotFifteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.16",
			cs:                                   csOneDotSixteen,
			expectedAPIServerImageString:         specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: csOneDotSixteen.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.Hyperkube],
			expectedAddonManagerImageString:           csOneDotSixteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.16 Azure Stack",
			cs:                                   csOneDotSixteenAzureStack,
			expectedAPIServerImageString:         csOneDotSixteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSixteenCustomImagesComponents[orchestratorVersionOneDotSixteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedControllerManagerImageString: csOneDotSixteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSixteenCustomImagesComponents[orchestratorVersionOneDotSixteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedCloudControllerManagerImageString: csOneDotSixteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase + csOneDotSixteenCustomImagesComponents[orchestratorVersionOneDotSixteen][common.CloudControllerManagerComponentName] + common.AzureStackSuffix,
			expectedSchedulerImageString:              csOneDotSixteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSixteenCustomImagesComponents[orchestratorVersionOneDotSixteen][common.Hyperkube] + common.AzureStackSuffix,
			expectedAddonManagerImageString:           csOneDotSixteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSixteenCustomImagesComponents[orchestratorVersionOneDotSixteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.17",
			cs:                                   csOneDotSeventeen,
			expectedAPIServerImageString:         csOneDotSeventeen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.APIServerComponentName],
			expectedControllerManagerImageString: csOneDotSeventeen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.ControllerManagerComponentName],
			expectedCloudControllerManagerImageString: csOneDotSeventeen.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              csOneDotSeventeen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.SchedulerComponentName],
			expectedAddonManagerImageString:           csOneDotSeventeen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.17 Azure Stack",
			cs:                                   csOneDotSeventeenAzureStack,
			expectedAPIServerImageString:         csOneDotSeventeenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSeventeenCustomImagesComponents[orchestratorVersionOneDotSeventeen][common.APIServerComponentName] + common.AzureStackSuffix,
			expectedControllerManagerImageString: csOneDotSeventeenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSeventeenCustomImagesComponents[orchestratorVersionOneDotSeventeen][common.ControllerManagerComponentName] + common.AzureStackSuffix,
			expectedCloudControllerManagerImageString: csOneDotSeventeenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase + csOneDotSeventeenCustomImagesComponents[orchestratorVersionOneDotSeventeen][common.CloudControllerManagerComponentName] + common.AzureStackSuffix,
			expectedSchedulerImageString:              csOneDotSeventeenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSeventeenCustomImagesComponents[orchestratorVersionOneDotSeventeen][common.SchedulerComponentName] + common.AzureStackSuffix,
			expectedAddonManagerImageString:           csOneDotSeventeenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotSeventeenCustomImagesComponents[orchestratorVersionOneDotSeventeen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.18",
			cs:                                   csOneDotEighteen,
			expectedAPIServerImageString:         csOneDotEighteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.APIServerComponentName],
			expectedControllerManagerImageString: csOneDotEighteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.ControllerManagerComponentName],
			expectedCloudControllerManagerImageString: csOneDotEighteen.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              csOneDotEighteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.SchedulerComponentName],
			expectedAddonManagerImageString:           csOneDotEighteen.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.18 Azure Stack",
			cs:                                   csOneDotEighteenAzureStack,
			expectedAPIServerImageString:         csOneDotEighteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotEighteenCustomImagesComponents[orchestratorVersionOneDotEighteen][common.APIServerComponentName] + common.AzureStackSuffix,
			expectedControllerManagerImageString: csOneDotEighteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotEighteenCustomImagesComponents[orchestratorVersionOneDotEighteen][common.ControllerManagerComponentName] + common.AzureStackSuffix,
			expectedCloudControllerManagerImageString: csOneDotEighteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase + csOneDotEighteenCustomImagesComponents[orchestratorVersionOneDotEighteen][common.CloudControllerManagerComponentName] + common.AzureStackSuffix,
			expectedSchedulerImageString:              csOneDotEighteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotEighteenCustomImagesComponents[orchestratorVersionOneDotEighteen][common.SchedulerComponentName] + common.AzureStackSuffix,
			expectedAddonManagerImageString:           csOneDotEighteenAzureStack.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + csOneDotEighteenCustomImagesComponents[orchestratorVersionOneDotEighteen][common.AddonManagerComponentName],
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			apiServerImageString := getComponentDefaultContainerImage(common.APIServerComponentName, c.cs)
			if apiServerImageString != c.expectedAPIServerImageString {
				t.Fatalf("expected getComponentDefaultContainerImage(%s, cs) result %s to be equal to %s", common.APIServerComponentName, apiServerImageString, c.expectedAPIServerImageString)
			}
			controllerManagerImageString := getComponentDefaultContainerImage(common.ControllerManagerComponentName, c.cs)
			if controllerManagerImageString != c.expectedControllerManagerImageString {
				t.Fatalf("expected getComponentDefaultContainerImage(%s, cs) result %s to be equal to %s", common.ControllerManagerComponentName, controllerManagerImageString, c.expectedControllerManagerImageString)
			}
			cloudControllerManagerImageString := getComponentDefaultContainerImage(common.CloudControllerManagerComponentName, c.cs)
			if cloudControllerManagerImageString != c.expectedCloudControllerManagerImageString {
				t.Fatalf("expected getComponentDefaultContainerImage(%s, cs) result %s to be equal to %s", common.CloudControllerManagerComponentName, cloudControllerManagerImageString, c.expectedCloudControllerManagerImageString)
			}
			schedulerImageString := getComponentDefaultContainerImage(common.SchedulerComponentName, c.cs)
			if schedulerImageString != c.expectedSchedulerImageString {
				t.Fatalf("expected getComponentDefaultContainerImage(%s, cs) result %s to be equal to %s", common.SchedulerComponentName, schedulerImageString, c.expectedSchedulerImageString)
			}
			addonManagerImageString := getComponentDefaultContainerImage(common.AddonManagerComponentName, c.cs)
			if addonManagerImageString != c.expectedAddonManagerImageString {
				t.Fatalf("expected getComponentDefaultContainerImage(%s, cs) result %s to be equal to %s", common.AddonManagerComponentName, addonManagerImageString, c.expectedAddonManagerImageString)
			}
		})
	}
}

func getDefaultComponents(cs *ContainerService) []KubernetesComponent {
	components := []KubernetesComponent{
		{
			Name:    common.SchedulerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.SchedulerComponentName,
					Image: getComponentDefaultContainerImage(common.SchedulerComponentName, cs),
				},
			},
			Config: map[string]string{
				"command": getSchedulerDefaultCommandString(cs),
			},
		},
		{
			Name:    common.ControllerManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.ControllerManagerComponentName,
					Image: getComponentDefaultContainerImage(common.ControllerManagerComponentName, cs),
				},
			},
			Config: map[string]string{
				"command": getControllerManagerDefaultCommandString(cs),
			},
		},
		{
			Name:    common.APIServerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.APIServerComponentName,
					Image: getComponentDefaultContainerImage(common.APIServerComponentName, cs),
				},
			},
			Config: map[string]string{
				"command": getAPIServerDefaultCommandString(cs),
			},
		},
		{
			Name:    common.AddonManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.AddonManagerComponentName,
					Image: getComponentDefaultContainerImage(common.AddonManagerComponentName, cs),
				},
			},
		},
	}

	return components
}

func concatenateDefaultComponents(components []KubernetesComponent, cs *ContainerService) []KubernetesComponent {
	defaults := getDefaultComponents(cs)
	defaults = append(components, defaults...)
	return defaults
}

func overwriteDefaultComponents(components []KubernetesComponent, cs *ContainerService) []KubernetesComponent {
	var ret []KubernetesComponent
	defaults := getDefaultComponents(cs)
	for _, defaultComponent := range defaults {
		i := GetComponentsIndexByName(components, defaultComponent.Name)
		if i > -1 {
			ret = append(ret, components[i])
		} else {
			ret = append(ret, defaultComponent)
		}
	}
	return ret
}

func getUserConfiguredComponentMap() map[string]KubernetesComponent {
	return map[string]KubernetesComponent{
		"user-configured kube-scheduler component": {
			Name:    common.SchedulerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.SchedulerComponentName,
					Image: customKubeSchedulerImage,
				},
			},
			Config: map[string]string{
				"command": "my-custom-kube-scheduler-command",
				"foo":     "bar",
			},
		},
		"user-configured controller-manager component": {
			Name:    common.ControllerManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.ControllerManagerComponentName,
					Image: customKubeControllerManagerImage,
				},
			},
			Config: map[string]string{
				"command": "my-custom-controller-manager-command",
				"foo":     "bar",
			},
		},
		"user-configured cloud-controller-manager component": {
			Name:    common.CloudControllerManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.CloudControllerManagerComponentName,
					Image: customCcmImage,
				},
			},
			Config: map[string]string{
				"command": "my-custom-cloud-controller-manager-command",
				"foo":     "bar",
			},
		},
		"user-configured kube-apiserver component": {
			Name:    common.APIServerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.APIServerComponentName,
					Image: customKubeAPIServerImage,
				},
			},
			Config: map[string]string{
				"command": "my-custom-kube-apiserver-command",
				"foo":     "bar",
			},
		},
		"user-configured kube-addon-manager component": {
			Name:    common.AddonManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.AddonManagerComponentName,
					Image: customKubeAddonManagerImage,
				},
			},
			Config: map[string]string{
				"command": "my-custom-kube-addon-manager-command",
				"foo":     "bar",
			},
		},
	}
}

func getContainerServicesMap() map[string]*ContainerService {
	specConfig := AzureCloudSpecEnvMap["AzurePublicCloud"].KubernetesSpecConfig
	return map[string]*ContainerService{
		"1.13": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.13 user-configured": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
						Components: []KubernetesComponent{
							getUserConfiguredComponentMap()["user-configured kube-scheduler component"],
							getUserConfiguredComponentMap()["user-configured controller-manager component"],
							getUserConfiguredComponentMap()["user-configured cloud-controller-manager component"],
							getUserConfiguredComponentMap()["user-configured kube-apiserver component"],
							getUserConfiguredComponentMap()["user-configured kube-addon-manager component"],
						},
					},
				},
			},
		},
		"1.13 + customHyperkubeImage + customCcmImage": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig: &KubernetesConfig{
						CustomHyperkubeImage:      customHyperkubeImage,
						CustomCcmImage:            customCcmImage,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.14": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.14 user-configured": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
						Components: []KubernetesComponent{
							getUserConfiguredComponentMap()["user-configured kube-scheduler component"],
							getUserConfiguredComponentMap()["user-configured controller-manager component"],
							getUserConfiguredComponentMap()["user-configured cloud-controller-manager component"],
							getUserConfiguredComponentMap()["user-configured kube-apiserver component"],
							getUserConfiguredComponentMap()["user-configured kube-addon-manager component"],
						},
					},
				},
			},
		},
		"1.14 + customHyperkubeImage + customCcmImage": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig: &KubernetesConfig{
						CustomHyperkubeImage:      customHyperkubeImage,
						CustomCcmImage:            customCcmImage,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.15": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.15 user-configured": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
						Components: []KubernetesComponent{
							getUserConfiguredComponentMap()["user-configured kube-scheduler component"],
							getUserConfiguredComponentMap()["user-configured controller-manager component"],
							getUserConfiguredComponentMap()["user-configured cloud-controller-manager component"],
							getUserConfiguredComponentMap()["user-configured kube-apiserver component"],
							getUserConfiguredComponentMap()["user-configured kube-addon-manager component"],
						},
					},
				},
			},
		},
		"1.15 + customHyperkubeImage + customCcmImage": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig: &KubernetesConfig{
						CustomHyperkubeImage:      customHyperkubeImage,
						CustomCcmImage:            customCcmImage,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.16": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.16 user-configured": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
						Components: []KubernetesComponent{
							getUserConfiguredComponentMap()["user-configured kube-scheduler component"],
							getUserConfiguredComponentMap()["user-configured controller-manager component"],
							getUserConfiguredComponentMap()["user-configured cloud-controller-manager component"],
							getUserConfiguredComponentMap()["user-configured kube-apiserver component"],
							getUserConfiguredComponentMap()["user-configured kube-addon-manager component"],
						},
					},
				},
			},
		},
		"1.16 + customHyperkubeImage + customCcmImage": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig: &KubernetesConfig{
						CustomHyperkubeImage:      customHyperkubeImage,
						CustomCcmImage:            customCcmImage,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.17": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.17 user-configured": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
						Components: []KubernetesComponent{
							getUserConfiguredComponentMap()["user-configured kube-scheduler component"],
							getUserConfiguredComponentMap()["user-configured controller-manager component"],
							getUserConfiguredComponentMap()["user-configured cloud-controller-manager component"],
							getUserConfiguredComponentMap()["user-configured kube-apiserver component"],
							getUserConfiguredComponentMap()["user-configured kube-addon-manager component"],
						},
					},
				},
			},
		},
		"1.17 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig: &KubernetesConfig{
						CustomCcmImage:                   customCcmImage,
						CustomKubeAPIServerImage:         customKubeAPIServerImage,
						CustomKubeControllerManagerImage: customKubeControllerManagerImage,
						CustomKubeSchedulerImage:         customKubeSchedulerImage,
						UseCloudControllerManager:        to.BoolPtr(true),
					},
				},
			},
		},
		"1.18": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.18 user-configured": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:     specConfig.MCRKubernetesImageBase,
						KubernetesImageBaseType: common.KubernetesImageBaseTypeMCR,
						MCRKubernetesImageBase:  specConfig.MCRKubernetesImageBase,
						Components: []KubernetesComponent{
							getUserConfiguredComponentMap()["user-configured kube-scheduler component"],
							getUserConfiguredComponentMap()["user-configured controller-manager component"],
							getUserConfiguredComponentMap()["user-configured cloud-controller-manager component"],
							getUserConfiguredComponentMap()["user-configured kube-apiserver component"],
							getUserConfiguredComponentMap()["user-configured kube-addon-manager component"],
						},
					},
				},
			},
		},
		"1.18 + customCcmImage + customKubeAPIServerImage + customKubeControllerManagerImage + customKubeSchedulerImage": {
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0",
					KubernetesConfig: &KubernetesConfig{
						CustomCcmImage:                   customCcmImage,
						CustomKubeAPIServerImage:         customKubeAPIServerImage,
						CustomKubeControllerManagerImage: customKubeControllerManagerImage,
						CustomKubeSchedulerImage:         customKubeSchedulerImage,
						UseCloudControllerManager:        to.BoolPtr(true),
					},
				},
			},
		},
	}
}
