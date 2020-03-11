// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"reflect"
	"testing"
)

func TestSetSysctlDConfig(t *testing.T) {
	cases := []struct {
		name           string
		cs             *ContainerService
		expectedMaster *MasterProfile
		expectedPools  []*AgentPoolProfile
	}{
		{
			name: "defaults",
			cs: &ContainerService{
				Properties: &Properties{
					MasterProfile: &MasterProfile{},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:   "foo",
							OSType: Linux,
						},
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.0",
						KubernetesConfig:    &KubernetesConfig{},
					},
				},
			},
			expectedMaster: &MasterProfile{
				SysctlDConfig: map[string]string{
					"net.ipv4.tcp_retries2":             "8",
					"net.core.somaxconn":                "16384",
					"net.ipv4.tcp_max_syn_backlog":      "16384",
					"net.core.message_cost":             "40",
					"net.core.message_burst":            "80",
					"net.ipv4.neigh.default.gc_thresh1": "4096",
					"net.ipv4.neigh.default.gc_thresh2": "8192",
					"net.ipv4.neigh.default.gc_thresh3": "16384",
				},
			},
			expectedPools: []*AgentPoolProfile{
				{
					Name: "foo",
					SysctlDConfig: map[string]string{
						"net.ipv4.tcp_retries2":             "8",
						"net.core.somaxconn":                "16384",
						"net.ipv4.tcp_max_syn_backlog":      "16384",
						"net.core.message_cost":             "40",
						"net.core.message_burst":            "80",
						"net.ipv4.neigh.default.gc_thresh1": "4096",
						"net.ipv4.neigh.default.gc_thresh2": "8192",
						"net.ipv4.neigh.default.gc_thresh3": "16384",
					},
				},
			},
		},
		{
			name: "defaults w/ containerd",
			cs: &ContainerService{
				Properties: &Properties{
					MasterProfile: &MasterProfile{},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:   "foo",
							OSType: Linux,
						},
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.0",
						KubernetesConfig: &KubernetesConfig{
							ContainerRuntime: Containerd,
						},
					},
				},
			},
			expectedMaster: &MasterProfile{
				SysctlDConfig: map[string]string{
					"net.ipv4.tcp_retries2":             "8",
					"net.core.somaxconn":                "16384",
					"net.ipv4.tcp_max_syn_backlog":      "16384",
					"net.core.message_cost":             "40",
					"net.core.message_burst":            "80",
					"net.ipv4.neigh.default.gc_thresh1": "4096",
					"net.ipv4.neigh.default.gc_thresh2": "8192",
					"net.ipv4.neigh.default.gc_thresh3": "16384",
					"net.ipv4.ip_forward":               "1",
				},
			},
			expectedPools: []*AgentPoolProfile{
				{
					Name: "foo",
					SysctlDConfig: map[string]string{
						"net.ipv4.tcp_retries2":             "8",
						"net.core.somaxconn":                "16384",
						"net.ipv4.tcp_max_syn_backlog":      "16384",
						"net.core.message_cost":             "40",
						"net.core.message_burst":            "80",
						"net.ipv4.neigh.default.gc_thresh1": "4096",
						"net.ipv4.neigh.default.gc_thresh2": "8192",
						"net.ipv4.neigh.default.gc_thresh3": "16384",
						"net.ipv4.ip_forward":               "1",
					},
				},
			},
		},
		{
			name: "user-configured",
			cs: &ContainerService{
				Properties: &Properties{
					MasterProfile: &MasterProfile{
						SysctlDConfig: map[string]string{
							"net.ipv4.tcp_keepalive_time": "120",
						},
					},
					AgentPoolProfiles: []*AgentPoolProfile{
						{
							Name:   "foo",
							OSType: Linux,
							SysctlDConfig: map[string]string{
								"net.ipv4.tcp_keepalive_time": "240",
							},
						},
						{
							Name:   "bar",
							OSType: Linux,
							SysctlDConfig: map[string]string{
								"net.ipv4.tcp_keepalive_time": "360",
							},
						},
					},
					OrchestratorProfile: &OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.15.0",
						KubernetesConfig: &KubernetesConfig{
							ContainerRuntime: Containerd,
						},
					},
				},
			},
			expectedMaster: &MasterProfile{
				SysctlDConfig: map[string]string{
					"net.ipv4.tcp_retries2":             "8",
					"net.core.somaxconn":                "16384",
					"net.ipv4.tcp_max_syn_backlog":      "16384",
					"net.core.message_cost":             "40",
					"net.core.message_burst":            "80",
					"net.ipv4.neigh.default.gc_thresh1": "4096",
					"net.ipv4.neigh.default.gc_thresh2": "8192",
					"net.ipv4.neigh.default.gc_thresh3": "16384",
					"net.ipv4.tcp_keepalive_time":       "120",
					"net.ipv4.ip_forward":               "1",
				},
			},
			expectedPools: []*AgentPoolProfile{
				{
					Name: "foo",
					SysctlDConfig: map[string]string{
						"net.ipv4.tcp_retries2":             "8",
						"net.core.somaxconn":                "16384",
						"net.ipv4.tcp_max_syn_backlog":      "16384",
						"net.core.message_cost":             "40",
						"net.core.message_burst":            "80",
						"net.ipv4.neigh.default.gc_thresh1": "4096",
						"net.ipv4.neigh.default.gc_thresh2": "8192",
						"net.ipv4.neigh.default.gc_thresh3": "16384",
						"net.ipv4.tcp_keepalive_time":       "240",
						"net.ipv4.ip_forward":               "1",
					},
				},
				{
					Name: "bar",
					SysctlDConfig: map[string]string{
						"net.ipv4.tcp_retries2":             "8",
						"net.core.somaxconn":                "16384",
						"net.ipv4.tcp_max_syn_backlog":      "16384",
						"net.core.message_cost":             "40",
						"net.core.message_burst":            "80",
						"net.ipv4.neigh.default.gc_thresh1": "4096",
						"net.ipv4.neigh.default.gc_thresh2": "8192",
						"net.ipv4.neigh.default.gc_thresh3": "16384",
						"net.ipv4.tcp_keepalive_time":       "360",
						"net.ipv4.ip_forward":               "1",
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.cs.setSysctlDConfig()
			if !reflect.DeepEqual(c.cs.Properties.MasterProfile.SysctlDConfig, c.expectedMaster.SysctlDConfig) {
				t.Fatalf("expected MasterProfile.SysctlDConfig %v to be equal to %v", c.cs.Properties.MasterProfile.SysctlDConfig, c.expectedMaster.SysctlDConfig)
			}
			for _, pool := range c.cs.Properties.AgentPoolProfiles {
				for _, expectedPool := range c.expectedPools {
					if pool.Name == expectedPool.Name {
						if !reflect.DeepEqual(pool.SysctlDConfig, expectedPool.SysctlDConfig) {
							t.Fatalf("expected AgentPoolProfile.SysctlDConfig %v to be equal to %v for pool Name=%s", pool.SysctlDConfig, expectedPool.SysctlDConfig, pool.Name)
						}
					}
				}
			}
		})
	}
}
