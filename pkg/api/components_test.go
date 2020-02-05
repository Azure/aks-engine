package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
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
			name:      "1.13 + CCM",
			cs:        containerServiceMap["1.13 + CCM"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: getContainerImage(common.CloudControllerManagerComponentName, containerServiceMap["1.13 + CCM"]),
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.13 + CCM"]),
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
			name:      "1.14 + CCM",
			cs:        containerServiceMap["1.14 + CCM"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: getContainerImage(common.CloudControllerManagerComponentName, containerServiceMap["1.14 + CCM"]),
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.14 + CCM"]),
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
			name:      "1.15 + CCM",
			cs:        containerServiceMap["1.15 + CCM"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: getContainerImage(common.CloudControllerManagerComponentName, containerServiceMap["1.15 + CCM"]),
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.15 + CCM"]),
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
			name:      "1.16 + CCM",
			cs:        containerServiceMap["1.16 + CCM"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: getContainerImage(common.CloudControllerManagerComponentName, containerServiceMap["1.16 + CCM"]),
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.16 + CCM"]),
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
			name:      "1.17 + CCM",
			cs:        containerServiceMap["1.17 + CCM"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: getContainerImage(common.CloudControllerManagerComponentName, containerServiceMap["1.17 + CCM"]),
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.17 + CCM"]),
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
			name:      "1.18 + CCM",
			cs:        containerServiceMap["1.18 + CCM"],
			isUpgrade: false,
			expectedComponents: concatenateDefaultComponents([]KubernetesComponent{
				{
					Name:    common.CloudControllerManagerComponentName,
					Enabled: to.BoolPtr(true),
					Containers: []KubernetesContainerSpec{
						{
							Name:  common.CloudControllerManagerComponentName,
							Image: getContainerImage(common.CloudControllerManagerComponentName, containerServiceMap["1.18 + CCM"]),
						},
					},
					Config: map[string]string{
						"command": "\"cloud-controller-manager\"",
					},
				},
			}, containerServiceMap["1.18 + CCM"]),
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
				component := test.cs.Properties.OrchestratorProfile.KubernetesConfig.Components[getComponentsIndexByName(test.cs.Properties.OrchestratorProfile.KubernetesConfig.Components, componentName)]
				if component.IsEnabled() {
					if i := getComponentsIndexByName(test.expectedComponents, componentName); i == -1 {
						t.Fatalf("got component %s that we weren't expecting", component.Name)
					}
					expectedComponent := test.expectedComponents[getComponentsIndexByName(test.expectedComponents, componentName)]
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
					if i := getComponentsIndexByName(test.expectedComponents, componentName); i > -1 {
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
			result := getComponentsIndexByName(c.components, c.componentName)
			if result != c.expected {
				t.Fatalf("expected getComponentsIndexByName() result %d to be equal to %d", result, c.expected)
			}
		})
	}
}

func TestAssignDefaultComponentVals(t *testing.T) {
	containerServiceMap := getContainerServicesMap()
	defaultOneDotFifteenComponents := getDefaultComponents(getContainerServicesMap()["1.15"])
	controllerManagerComponent := defaultOneDotFifteenComponents[getComponentsIndexByName(defaultOneDotFifteenComponents, common.ControllerManagerComponentName)]
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
						Image: getContainerImage(common.ControllerManagerComponentName, containerServiceMap["1.15"]),
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
						Image: getContainerImage(common.ControllerManagerComponentName, containerServiceMap["1.15"]),
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
					"foo": "bar",
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
						Image:          getContainerImage(common.ControllerManagerComponentName, containerServiceMap["1.15"]),
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
	i := getComponentsIndexByName(defaultOneDotFifteenComponents, common.ControllerManagerComponentName)
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
			i := getComponentsIndexByName(c.components, common.ControllerManagerComponentName)
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
			expectedAPIServerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-apiserver\""),
			expectedControllerManagerCommandString: fmt.Sprintf("\"/hyperkube\", \"kube-controller-manager\""),
			expectedSchedulerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-scheduler\""),
		},
		{
			name:                                   "1.14",
			cs:                                     getContainerServicesMap()["1.14"],
			expectedAPIServerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-apiserver\""),
			expectedControllerManagerCommandString: fmt.Sprintf("\"/hyperkube\", \"kube-controller-manager\""),
			expectedSchedulerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-scheduler\""),
		},
		{
			name:                                   "1.15",
			cs:                                     getContainerServicesMap()["1.15"],
			expectedAPIServerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-apiserver\""),
			expectedControllerManagerCommandString: fmt.Sprintf("\"/hyperkube\", \"kube-controller-manager\""),
			expectedSchedulerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-scheduler\""),
		},
		{
			name:                                   "1.16",
			cs:                                     getContainerServicesMap()["1.16"],
			expectedAPIServerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-apiserver\""),
			expectedControllerManagerCommandString: fmt.Sprintf("\"/hyperkube\", \"kube-controller-manager\""),
			expectedSchedulerCommandString:         fmt.Sprintf("\"/hyperkube\", \"kube-scheduler\""),
		},
		{
			name:                                   "1.17",
			cs:                                     getContainerServicesMap()["1.17"],
			expectedAPIServerCommandString:         fmt.Sprintf("\"kube-apiserver\""),
			expectedControllerManagerCommandString: fmt.Sprintf("\"kube-controller-manager\""),
			expectedSchedulerCommandString:         fmt.Sprintf("\"kube-scheduler\""),
		},
		{
			name:                                   "1.18",
			cs:                                     getContainerServicesMap()["1.18"],
			expectedAPIServerCommandString:         fmt.Sprintf("\"kube-apiserver\""),
			expectedControllerManagerCommandString: fmt.Sprintf("\"kube-controller-manager\""),
			expectedSchedulerCommandString:         fmt.Sprintf("\"kube-scheduler\""),
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
	azureStackCloudSpec := AzureEnvironmentSpecConfig{
		CloudName: "AzureStackCloud",
		KubernetesSpecConfig: KubernetesSpecConfig{
			KubernetesImageBase:              "KubernetesImageBase",
			TillerImageBase:                  "TillerImageBase",
			ACIConnectorImageBase:            "ACIConnectorImageBase",
			NVIDIAImageBase:                  "NVIDIAImageBase",
			AzureCNIImageBase:                "AzureCNIImageBase",
			CalicoImageBase:                  "CalicoImageBase",
			EtcdDownloadURLBase:              "EtcdDownloadURLBase",
			KubeBinariesSASURLBase:           "KubeBinariesSASURLBase",
			WindowsTelemetryGUID:             "WindowsTelemetryGUID",
			CNIPluginsDownloadURL:            "CNIPluginsDownloadURL",
			VnetCNILinuxPluginsDownloadURL:   "VnetCNILinuxPluginsDownloadURL",
			VnetCNIWindowsPluginsDownloadURL: "VnetCNIWindowsPluginsDownloadURL",
			ContainerdDownloadURLBase:        "ContainerdDownloadURLBase",
		},
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
		},
	}
	AzureCloudSpecEnvMap[AzureStackCloud] = azureStackCloudSpec
	csOneDotThirteen := getContainerServicesMap()["1.13"]
	orchestratorVersionOneDotThirteen := csOneDotThirteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotFourteen := getContainerServicesMap()["1.14"]
	orchestratorVersionOneDotFourteen := csOneDotFourteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotFifteen := getContainerServicesMap()["1.15"]
	orchestratorVersionOneDotFifteen := csOneDotFifteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotSixteen := getContainerServicesMap()["1.16"]
	orchestratorVersionOneDotSixteen := csOneDotSixteen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotSeventeen := getContainerServicesMap()["1.17"]
	orchestratorVersionOneDotSeventeen := csOneDotSeventeen.Properties.OrchestratorProfile.OrchestratorVersion
	csOneDotEighteen := getContainerServicesMap()["1.18"]
	orchestratorVersionOneDotEighteen := csOneDotEighteen.Properties.OrchestratorProfile.OrchestratorVersion
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
			cs:                                   getContainerServicesMap()["1.13"],
			expectedAPIServerImageString:         specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.Hyperkube],
			expectedAddonManagerImageString:           specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotThirteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.14",
			cs:                                   csOneDotFourteen,
			expectedAPIServerImageString:         specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.Hyperkube],
			expectedAddonManagerImageString:           specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFourteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.15",
			cs:                                   csOneDotFifteen,
			expectedAPIServerImageString:         specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.Hyperkube],
			expectedAddonManagerImageString:           specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotFifteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.16",
			cs:                                   csOneDotSixteen,
			expectedAPIServerImageString:         specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.Hyperkube],
			expectedControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.Hyperkube],
			expectedCloudControllerManagerImageString: specConfig.MCRKubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.Hyperkube],
			expectedAddonManagerImageString:           specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSixteen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.17",
			cs:                                   csOneDotSeventeen,
			expectedAPIServerImageString:         specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.APIServerComponentName],
			expectedControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.ControllerManagerComponentName],
			expectedCloudControllerManagerImageString: specConfig.MCRKubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.SchedulerComponentName],
			expectedAddonManagerImageString:           specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotSeventeen][common.AddonManagerComponentName],
		},
		{
			name:                                 "1.18",
			cs:                                   csOneDotEighteen,
			expectedAPIServerImageString:         specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.APIServerComponentName],
			expectedControllerManagerImageString: specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.ControllerManagerComponentName],
			expectedCloudControllerManagerImageString: specConfig.MCRKubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.CloudControllerManagerComponentName],
			expectedSchedulerImageString:              specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.SchedulerComponentName],
			expectedAddonManagerImageString:           specConfig.KubernetesImageBase + K8sComponentsByVersionMap[orchestratorVersionOneDotEighteen][common.AddonManagerComponentName],
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			apiServerImageString := getContainerImage(common.APIServerComponentName, c.cs)
			if apiServerImageString != c.expectedAPIServerImageString {
				t.Fatalf("expected getContainerImage(%s, cs) result %s to be equal to %s", common.APIServerComponentName, apiServerImageString, c.expectedAPIServerImageString)
			}
			controllerManagerImageString := getContainerImage(common.ControllerManagerComponentName, c.cs)
			if controllerManagerImageString != c.expectedControllerManagerImageString {
				t.Fatalf("expected getContainerImage(%s, cs) result %s to be equal to %s", common.ControllerManagerComponentName, controllerManagerImageString, c.expectedControllerManagerImageString)
			}
			cloudControllerManagerImageString := getContainerImage(common.CloudControllerManagerComponentName, c.cs)
			if cloudControllerManagerImageString != c.expectedCloudControllerManagerImageString {
				t.Fatalf("expected getContainerImage(%s, cs) result %s to be equal to %s", common.CloudControllerManagerComponentName, cloudControllerManagerImageString, c.expectedCloudControllerManagerImageString)
			}
			schedulerImageString := getContainerImage(common.SchedulerComponentName, c.cs)
			if schedulerImageString != c.expectedSchedulerImageString {
				t.Fatalf("expected getContainerImage(%s, cs) result %s to be equal to %s", common.SchedulerComponentName, schedulerImageString, c.expectedSchedulerImageString)
			}
			addonManagerImageString := getContainerImage(common.AddonManagerComponentName, c.cs)
			if addonManagerImageString != c.expectedAddonManagerImageString {
				t.Fatalf("expected getContainerImage(%s, cs) result %s to be equal to %s", common.AddonManagerComponentName, addonManagerImageString, c.expectedAddonManagerImageString)
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
					Image: getContainerImage(common.SchedulerComponentName, cs),
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
					Image: getContainerImage(common.ControllerManagerComponentName, cs),
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
					Image: getContainerImage(common.APIServerComponentName, cs),
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
					Image: getContainerImage(common.AddonManagerComponentName, cs),
				},
			},
		},
	}

	return components
}

func concatenateDefaultComponents(components []KubernetesComponent, cs *ContainerService) []KubernetesComponent {
	defaults := getDefaultComponents(cs)
	defaults = append(defaults, components...)
	return defaults
}

func overwriteDefaultComponents(components []KubernetesComponent, cs *ContainerService) []KubernetesComponent {
	var ret []KubernetesComponent
	defaults := getDefaultComponents(cs)
	for _, defaultComponent := range defaults {
		i := getComponentsIndexByName(components, defaultComponent.Name)
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
		"user-configured kube-scheduler component": KubernetesComponent{
			Name:    common.SchedulerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.SchedulerComponentName,
					Image: "my-custom-kube-scheduler-image",
				},
			},
			Config: map[string]string{
				"command": "my-custom-kube-scheduler-command",
				"foo":     "bar",
			},
		},
		"user-configured controller-manager component": KubernetesComponent{
			Name:    common.ControllerManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.ControllerManagerComponentName,
					Image: "my-custom-controller-manager-image",
				},
			},
			Config: map[string]string{
				"command": "my-custom-controller-manager-command",
				"foo":     "bar",
			},
		},
		"user-configured cloud-controller-manager component": KubernetesComponent{
			Name:    common.CloudControllerManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.CloudControllerManagerComponentName,
					Image: "my-custom-cloud-controller-manager-image",
				},
			},
			Config: map[string]string{
				"command": "my-custom-cloud-controller-manager-command",
				"foo":     "bar",
			},
		},
		"user-configured kube-apiserver component": KubernetesComponent{
			Name:    common.APIServerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.APIServerComponentName,
					Image: "my-custom-kube-apiserver-image",
				},
			},
			Config: map[string]string{
				"command": "my-custom-kube-apiserver-command",
				"foo":     "bar",
			},
		},
		"user-configured kube-addon-manager component": KubernetesComponent{
			Name:    common.AddonManagerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.AddonManagerComponentName,
					Image: "my-custom-kube-addon-manager-image",
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
		"1.13": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.13 user-configured": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
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
		"1.13 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:       specConfig.KubernetesImageBase,
						MCRKubernetesImageBase:    specConfig.MCRKubernetesImageBase,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.14": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.14 user-configured": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
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
		"1.14 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:       specConfig.KubernetesImageBase,
						MCRKubernetesImageBase:    specConfig.MCRKubernetesImageBase,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.15": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.15 user-configured": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
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
		"1.15 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:       specConfig.KubernetesImageBase,
						MCRKubernetesImageBase:    specConfig.MCRKubernetesImageBase,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.16": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.16 user-configured": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
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
		"1.16 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:       specConfig.KubernetesImageBase,
						MCRKubernetesImageBase:    specConfig.MCRKubernetesImageBase,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.17": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.17 user-configured": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
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
		"1.17 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:       specConfig.KubernetesImageBase,
						MCRKubernetesImageBase:    specConfig.MCRKubernetesImageBase,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.18": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0-alpha.1",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
					},
				},
			},
		},
		"1.18 user-configured": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0-alpha.1",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:    specConfig.KubernetesImageBase,
						MCRKubernetesImageBase: specConfig.MCRKubernetesImageBase,
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
		"1.18 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0-alpha.1",
					KubernetesConfig: &KubernetesConfig{
						KubernetesImageBase:       specConfig.KubernetesImageBase,
						MCRKubernetesImageBase:    specConfig.MCRKubernetesImageBase,
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
	}
}
