package api

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestSetComponentsConfig(t *testing.T) {
	userConfiguredComponentsMap := map[string]KubernetesComponent{
		"user-configured scheduler component": KubernetesComponent{
			Name:    common.SchedulerComponentName,
			Enabled: to.BoolPtr(true),
			Containers: []KubernetesContainerSpec{
				{
					Name:  common.SchedulerComponentName,
					Image: "my-custom-image",
				},
			},
			Config: map[string]string{
				"command": "my-custom-command",
				"foo":     "bar",
			},
		},
	}
	containerServiceMap := map[string]*ContainerService{
		"1.13": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig:    &KubernetesConfig{},
				},
			},
		},
		"1.13 user-configured": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.13.11",
					KubernetesConfig: &KubernetesConfig{
						Components: []KubernetesComponent{
							userConfiguredComponentsMap["user-configured scheduler component"],
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
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.14": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig:    &KubernetesConfig{},
				},
			},
		},
		"1.14 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.14.7",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.15": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig:    &KubernetesConfig{},
				},
			},
		},
		"1.15 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.15.9",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.16": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig:    &KubernetesConfig{},
				},
			},
		},
		"1.16 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.16.6",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.17": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig:    &KubernetesConfig{},
				},
			},
		},
		"1.17 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.17.2",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
		"1.18": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0-alpha.1",
					KubernetesConfig:    &KubernetesConfig{},
				},
			},
		},
		"1.18 + CCM": &ContainerService{
			Properties: &Properties{
				OrchestratorProfile: &OrchestratorProfile{
					OrchestratorVersion: "1.18.0-alpha.1",
					KubernetesConfig: &KubernetesConfig{
						UseCloudControllerManager: to.BoolPtr(true),
					},
				},
			},
		},
	}

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
			expectedComponents: overwriteDefaultComponents([]KubernetesComponent{
				userConfiguredComponentsMap["user-configured scheduler component"],
			}, containerServiceMap["1.13 user-configured"]),
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
	for _, componentOverride := range components {
		for _, component := range defaults {
			if component.Name == componentOverride.Name {
				ret = append(ret, componentOverride)
			} else {
				ret = append(ret, component)
			}
		}
	}
	return ret
}
