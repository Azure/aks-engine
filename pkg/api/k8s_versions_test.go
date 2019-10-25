// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"
)

func TestGetK8sVersionComponents(t *testing.T) {
	oneDotFifteenDotZero := getK8sVersionComponents("1.15.0", nil)
	if oneDotFifteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent := k8sComponentVersions["1.15"]
	expected := map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.15.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.15.0",
		"windowszip":                       "v1.15.0-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"coredns":                          k8sComponent["coredns"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponent[NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
	}

	for k, v := range oneDotFifteenDotZero {
		if expected[k] != v {
			t.Fatalf("getK8sVersionComponents() returned an unexpected map[string]string value for k8s 1.15.0: %s = %s", k, oneDotFifteenDotZero[k])
		}
	}

	oneDotFourteenDotZero := getK8sVersionComponents("1.14.0", nil)
	if oneDotFourteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.14"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.14.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.14.0",
		"windowszip":                       "v1.14.0-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"coredns":                          k8sComponent["coredns"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponent[NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.13"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.13.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.13.0",
		"windowszip":                       "v1.13.0-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"coredns":                          k8sComponent["coredns"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponent[NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.12"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.12.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.12.0",
		"windowszip":                       "v1.12.0-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"coredns":                          k8sComponent["coredns"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponent[NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.11"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.11.0-alpha.1",
		"ccm":                              "cloud-controller-manager-amd64:v1.11.0-alpha.1",
		"windowszip":                       "v1.11.0-alpha.1-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponent[NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.10"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.10.0",
		"ccm":                              "cloud-controller-manager-amd64:v1.10.0",
		"windowszip":                       "v1.10.0-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		NVIDIADevicePluginAddonName:        k8sComponent[NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.9"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.9.3",
		"ccm":                              "cloud-controller-manager-amd64:v1.9.3",
		"windowszip":                       "v1.9.3-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.8"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.8.8",
		"ccm":                              "cloud-controller-manager-amd64:v1.8.8",
		"windowszip":                       "v1.8.8-1int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.7"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.7.13",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
	k8sComponent = k8sComponentVersions["1.9"]
	expected = map[string]string{
		"hyperkube":                        "hyperkube-amd64:v1.9.3",
		"ccm":                              "cloud-controller-manager-amd64:v1.9.3",
		"windowszip":                       "v1.9.3-2int.zip",
		DashboardAddonName:                 k8sComponent["dashboard"],
		"exechealthz":                      k8sComponent["exechealthz"],
		"addonresizer":                     k8sComponent["addon-resizer"],
		"heapster":                         k8sComponent["heapster"],
		MetricsServerAddonName:             k8sComponent["metrics-server"],
		"kube-dns":                         k8sComponent["kube-dns"],
		"addonmanager":                     k8sComponent["addon-manager"],
		"dnsmasq":                          k8sComponent["dnsmasq"],
		"pause":                            k8sComponent["pause"],
		TillerAddonName:                    k8sComponent["tiller"],
		ReschedulerAddonName:               k8sComponent["rescheduler"],
		ACIConnectorAddonName:              k8sComponent["aci-connector"],
		ContainerMonitoringAddonName:       k8sComponent[ContainerMonitoringAddonName],
		AzureCNINetworkMonitoringAddonName: k8sComponent[AzureCNINetworkMonitoringAddonName],
		ClusterAutoscalerAddonName:         k8sComponent["cluster-autoscaler"],
		"k8s-dns-sidecar":                  k8sComponent["k8s-dns-sidecar"],
		"nodestatusfreq":                   k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                  k8sComponent["nodegraceperiod"],
		"podeviction":                      k8sComponent["podeviction"],
		"routeperiod":                      k8sComponent["routeperiod"],
		"backoffretries":                   k8sComponent["backoffretries"],
		"backoffjitter":                    k8sComponent["backoffjitter"],
		"backoffduration":                  k8sComponent["backoffduration"],
		"backoffexponent":                  k8sComponent["backoffexponent"],
		"ratelimitqps":                     k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                  k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":             k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                  k8sComponent["gchighthreshold"],
		"gclowthreshold":                   k8sComponent["gclowthreshold"],
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
