// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"
)

func TestGetK8sVersionComponents(t *testing.T) {
	oneDotFourteenDotZero := getK8sVersionComponents("1.14.0", nil)
	if oneDotFourteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected := map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.14.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.14.0",
		"windowszip":                       "v1.14.0-1int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.14"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.14"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.14"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.14"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.14"]["metrics-server"],
		"coredns":                          k8sComponentVersions["1.14"]["coredns"],
		"kube-dns":                         k8sComponentVersions["1.14"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.14"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.14"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.14"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.14"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.14"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.14"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.14"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.14"][AzureCNINetworkMonitoringAddonName],
		DefaultClusterAutoscalerAddonName:  k8sComponentVersions["1.14"]["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponentVersions["1.14"][NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponentVersions["1.14"]["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponentVersions["1.14"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.14"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.14"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.14"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.14"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.14"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.14"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.14"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.14"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.14"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.14"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.14"]["gclowthreshold"],
	}

	for k, v := range oneDotFourteenDotZero {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.14.0: %s = %s", k, oneDotFourteenDotZero[k])
		}
	}

	oneDotThirteenDotZero := getK8sVersionComponents("1.13.0", nil)
	if oneDotThirteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.13.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.13.0",
		"windowszip":                       "v1.13.0-1int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.13"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.13"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.13"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.13"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.13"]["metrics-server"],
		"coredns":                          k8sComponentVersions["1.13"]["coredns"],
		"kube-dns":                         k8sComponentVersions["1.13"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.13"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.13"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.13"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.13"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.13"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.13"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.13"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.13"][AzureCNINetworkMonitoringAddonName],
		DefaultClusterAutoscalerAddonName:  k8sComponentVersions["1.13"]["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponentVersions["1.13"][NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponentVersions["1.13"]["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponentVersions["1.13"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.13"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.13"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.13"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.13"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.13"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.13"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.13"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.13"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.13"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.13"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.13"]["gclowthreshold"],
	}

	for k, v := range oneDotThirteenDotZero {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.13.0: %s = %s", k, oneDotThirteenDotZero[k])
		}
	}

	oneDotTwelveDotZero := getK8sVersionComponents("1.12.0", nil)
	if oneDotTwelveDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.12.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.12.0",
		"windowszip":                       "v1.12.0-1int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.12"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.12"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.12"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.12"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.12"]["metrics-server"],
		"coredns":                          k8sComponentVersions["1.12"]["coredns"],
		"kube-dns":                         k8sComponentVersions["1.12"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.12"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.12"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.12"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.12"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.12"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.12"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.12"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.12"][AzureCNINetworkMonitoringAddonName],
		DefaultClusterAutoscalerAddonName:  k8sComponentVersions["1.12"]["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponentVersions["1.12"][NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponentVersions["1.12"]["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponentVersions["1.12"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.12"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.12"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.12"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.12"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.12"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.12"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.12"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.12"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.12"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.12"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.12"]["gclowthreshold"],
	}

	for k, v := range oneDotTwelveDotZero {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.12.0: %s = %s", k, oneDotTwelveDotZero[k])
		}
	}

	oneDotElevenDotZero := getK8sVersionComponents("1.11.0-alpha.1", nil)
	if oneDotElevenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.11.0-alpha.1",
		"ccm":                              "cloud-controller-manager-amd64:v1.11.0-alpha.1",
		"windowszip":                       "v1.11.0-alpha.1-1int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.11"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.11"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.11"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.11"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.11"]["metrics-server"],
		"kube-dns":                         k8sComponentVersions["1.11"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.11"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.11"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.11"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.11"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.11"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.11"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.11"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.11"][AzureCNINetworkMonitoringAddonName],
		DefaultClusterAutoscalerAddonName:  k8sComponentVersions["1.11"]["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponentVersions["1.11"][NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponentVersions["1.11"]["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponentVersions["1.11"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.11"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.11"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.11"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.11"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.11"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.11"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.11"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.11"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.11"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.11"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.11"]["gclowthreshold"],
	}

	for k, v := range oneDotElevenDotZero {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.11.0-alpha.1: %s = %s", k, oneDotElevenDotZero[k])
		}
	}

	oneDotTenDotZero := getK8sVersionComponents("1.10.0", nil)
	if oneDotTenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.10.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.10.0",
		"windowszip":                       "v1.10.0-1int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.10"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.10"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.10"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.10"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.10"]["metrics-server"],
		"kube-dns":                         k8sComponentVersions["1.10"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.10"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.10"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.10"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.10"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.10"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.10"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.10"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.10"][AzureCNINetworkMonitoringAddonName],
		DefaultClusterAutoscalerAddonName:  k8sComponentVersions["1.10"]["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponentVersions["1.10"][NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponentVersions["1.10"]["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponentVersions["1.10"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.10"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.10"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.10"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.10"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.10"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.10"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.10"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.10"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.10"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.10"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.10"]["gclowthreshold"],
	}

	for k, v := range oneDotTenDotZero {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.10.0: %s = %s", k, oneDotTenDotZero[k])
		}
	}

	oneDotNineDotThree := getK8sVersionComponents("1.9.3", nil)
	if oneDotNineDotThree == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.9.3",
		"ccm":                              "cloud-controller-manager-amd64:v1.9.3",
		"windowszip":                       "v1.9.3-1int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.9"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.9"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.9"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.9"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.9"]["metrics-server"],
		"kube-dns":                         k8sComponentVersions["1.9"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.9"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.9"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.9"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.9"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.9"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.9"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.9"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.9"][AzureCNINetworkMonitoringAddonName],
		DefaultClusterAutoscalerAddonName:  k8sComponentVersions["1.9"]["cluster-autoscaler"],
		"k8s-dns-sidecar":                  k8sComponentVersions["1.9"]["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponentVersions["1.9"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.9"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.9"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.9"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.9"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.9"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.9"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.9"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.9"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.9"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.9"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.9"]["gclowthreshold"],
	}

	for k, v := range oneDotNineDotThree {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.9.3: %s = %s", k, oneDotNineDotThree[k])
		}
	}

	oneDotEightDotEight := getK8sVersionComponents("1.8.8", nil)
	if oneDotEightDotEight == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.8.8",
		"ccm":                              "cloud-controller-manager-amd64:v1.8.8",
		"windowszip":                       "v1.8.8-1int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.8"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.8"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.8"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.8"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.8"]["metrics-server"],
		"kube-dns":                         k8sComponentVersions["1.8"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.8"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.8"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.8"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.8"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.8"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.8"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.11"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.8"][AzureCNINetworkMonitoringAddonName],
		"nodestatusfreq":                   k8sComponentVersions["1.8"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.8"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.8"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.8"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.8"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.8"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.8"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.8"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.8"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.8"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.8"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.8"]["gclowthreshold"],
	}
	for k, v := range oneDotEightDotEight {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.8.8: %s = %s", k, oneDotNineDotThree[k])
		}
	}

	oneDotSevenDotZero := getK8sVersionComponents("1.7.13", nil)
	if oneDotSevenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.7.13",
		DefaultDashboardAddonName:          k8sComponentVersions["1.7"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.7"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.7"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.7"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.7"]["metrics-server"],
		"kube-dns":                         k8sComponentVersions["1.7"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.7"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.7"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.7"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.7"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.7"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.7"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.11"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.7"][AzureCNINetworkMonitoringAddonName],
		"nodestatusfreq":                   k8sComponentVersions["1.7"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.7"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.7"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.7"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.7"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.7"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.7"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.7"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.7"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.7"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.7"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.7"]["gclowthreshold"],
	}
	for k, v := range oneDotSevenDotZero {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.7.0: %s = %s", k, oneDotSevenDotZero[k])
		}
	}

	override := getK8sVersionComponents("1.9.3", map[string]string{"windowszip": "v1.9.3-2int.zip"})
	if override == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.9.3",
		"ccm":                              "cloud-controller-manager-amd64:v1.9.3",
		"windowszip":                       "v1.9.3-2int.zip",
		DefaultDashboardAddonName:          k8sComponentVersions["1.9"]["dashboard"],
		"exechealthz":                      k8sComponentVersions["1.9"]["exechealthz"],
		"addonresizer":                     k8sComponentVersions["1.9"]["addon-resizer"],
		"heapster":                         k8sComponentVersions["1.9"]["heapster"],
		DefaultMetricsServerAddonName:      k8sComponentVersions["1.9"]["metrics-server"],
		"kube-dns":                         k8sComponentVersions["1.9"]["kube-dns"],
		"addonmanager":                     k8sComponentVersions["1.9"]["addon-manager"],
		"dnsmasq":                          k8sComponentVersions["1.9"]["dnsmasq"],
		"pause":                            k8sComponentVersions["1.9"]["pause"],
		DefaultTillerAddonName:             k8sComponentVersions["1.9"]["tiller"],
		DefaultReschedulerAddonName:        k8sComponentVersions["1.9"]["rescheduler"],
		DefaultACIConnectorAddonName:       k8sComponentVersions["1.9"]["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponentVersions["1.11"][ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponentVersions["1.9"][AzureCNINetworkMonitoringAddonName],
		DefaultClusterAutoscalerAddonName:  k8sComponentVersions["1.9"]["cluster-autoscaler"],
		"k8s-dns-sidecar":                  k8sComponentVersions["1.9"]["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponentVersions["1.9"]["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponentVersions["1.9"]["nodegraceperiod"],
		"podeviction":                      k8sComponentVersions["1.9"]["podeviction"],
		"routeperiod":                      k8sComponentVersions["1.9"]["routeperiod"],
		"backoffretries":                   k8sComponentVersions["1.9"]["backoffretries"],
		"backoffjitter":                    k8sComponentVersions["1.9"]["backoffjitter"],
		"backoffduration":                  k8sComponentVersions["1.9"]["backoffduration"],
		"backoffexponent":                  k8sComponentVersions["1.9"]["backoffexponent"],
		"ratelimitqps":                     k8sComponentVersions["1.9"]["ratelimitqps"],
		"ratelimitbucket":                  k8sComponentVersions["1.9"]["ratelimitbucket"],
		"gchighthreshold":                  k8sComponentVersions["1.9"]["gchighthreshold"],
		"gclowthreshold":                   k8sComponentVersions["1.9"]["gclowthreshold"],
	}
	for k, v := range override {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.9.3 w/ overrides: %s = %s", k, override[k])
		}
	}

	unknown := getK8sVersionComponents("1.0.0", nil)
	if unknown != nil {
		t.Fatalf("getK8sVersionComponents() should return nil for unknown k8s version")
	}
}
