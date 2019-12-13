// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api/common"
	. "github.com/onsi/gomega"
)

func TestGetK8sVersionComponents(t *testing.T) {
	g := NewGomegaWithT(t)

	oneDotSeventeenDotZero := getK8sVersionComponents("1.17.0", nil)
	if oneDotSeventeenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent := k8sComponentVersions["1.17"]
	expected := map[string]string{
		"kube-scheduler":                       "kube-scheduler:v1.17.0",
		"kube-controller-manager":              "kube-controller-manager:v1.17.0",
		"kube-apiserver":                       "kube-apiserver:v1.17.0",
		"kube-proxy":                           "kube-proxy:v1.17.0",
		"ccm":                                  "azure-cloud-controller-manager:v0.3.0",
		common.CloudNodeManagerAddonName:       "azure-cloud-node-manager:v0.3.0",
		"windowszip":                           "v1.17.0-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		common.CoreDNSAddonName:                k8sComponent[common.CoreDNSAddonName],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
	}
	g.Expect(oneDotSeventeenDotZero).To(Equal(expected))

	oneDotSixteenDotZero := getK8sVersionComponents("1.16.0", nil)
	if oneDotSixteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.16"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.16.0",
		"kube-proxy":                           "hyperkube-amd64:v1.16.0",
		"ccm":                                  "azure-cloud-controller-manager:v0.3.0",
		common.CloudNodeManagerAddonName:       "azure-cloud-node-manager:v0.3.0",
		"windowszip":                           "v1.16.0-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		common.CoreDNSAddonName:                k8sComponent[common.CoreDNSAddonName],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
	}
	g.Expect(oneDotSixteenDotZero).To(Equal(expected))

	oneDotFifteenDotZero := getK8sVersionComponents("1.15.0", nil)
	if oneDotFifteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.15"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.15.0",
		"kube-proxy":                           "hyperkube-amd64:v1.15.0",
		"ccm":                                  "cloud-controller-manager-amd64:v1.15.0",
		"windowszip":                           "v1.15.0-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		common.CoreDNSAddonName:                k8sComponent[common.CoreDNSAddonName],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotFifteenDotZero).To(Equal(expected))

	oneDotFourteenDotZero := getK8sVersionComponents("1.14.0", nil)
	if oneDotFourteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.14"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.14.0",
		"kube-proxy":                           "hyperkube-amd64:v1.14.0",
		"ccm":                                  "cloud-controller-manager-amd64:v1.14.0",
		"windowszip":                           "v1.14.0-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		common.CoreDNSAddonName:                k8sComponent[common.CoreDNSAddonName],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotFourteenDotZero).To(Equal(expected))

	oneDotThirteenDotZero := getK8sVersionComponents("1.13.0", nil)
	if oneDotThirteenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.13"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.13.0",
		"kube-proxy":                           "hyperkube-amd64:v1.13.0",
		"ccm":                                  "cloud-controller-manager-amd64:v1.13.0",
		"windowszip":                           "v1.13.0-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		common.CoreDNSAddonName:                k8sComponent[common.CoreDNSAddonName],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotThirteenDotZero).To(Equal(expected))

	oneDotTwelveDotZero := getK8sVersionComponents("1.12.0", nil)
	if oneDotTwelveDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.12"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.12.0",
		"kube-proxy":                           "hyperkube-amd64:v1.12.0",
		"ccm":                                  "cloud-controller-manager-amd64:v1.12.0",
		"windowszip":                           "v1.12.0-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		common.CoreDNSAddonName:                k8sComponent[common.CoreDNSAddonName],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotTwelveDotZero).To(Equal(expected))

	oneDotElevenDotZero := getK8sVersionComponents("1.11.0-alpha.1", nil)
	if oneDotElevenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.11"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.11.0-alpha.1",
		"kube-proxy":                           "hyperkube-amd64:v1.11.0-alpha.1",
		"ccm":                                  "cloud-controller-manager-amd64:v1.11.0-alpha.1",
		"windowszip":                           "v1.11.0-alpha.1-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotElevenDotZero).To(Equal(expected))

	oneDotTenDotZero := getK8sVersionComponents("1.10.0", nil)
	if oneDotTenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.10"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.10.0",
		"kube-proxy":                           "hyperkube-amd64:v1.10.0",
		"ccm":                                  "cloud-controller-manager-amd64:v1.10.0",
		"windowszip":                           "v1.10.0-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		common.NVIDIADevicePluginAddonName:     k8sComponent[common.NVIDIADevicePluginAddonName],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotTenDotZero).To(Equal(expected))

	oneDotNineDotThree := getK8sVersionComponents("1.9.3", nil)
	if oneDotNineDotThree == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.9"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.9.3",
		"kube-proxy":                           "hyperkube-amd64:v1.9.3",
		"ccm":                                  "cloud-controller-manager-amd64:v1.9.3",
		"windowszip":                           "v1.9.3-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotNineDotThree).To(Equal(expected))

	oneDotEightDotEight := getK8sVersionComponents("1.8.8", nil)
	if oneDotEightDotEight == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.8"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.8.8",
		"kube-proxy":                           "hyperkube-amd64:v1.8.8",
		"ccm":                                  "cloud-controller-manager-amd64:v1.8.8",
		"windowszip":                           "v1.8.8-1int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotEightDotEight).To(Equal(expected))

	oneDotSevenDotZero := getK8sVersionComponents("1.7.13", nil)
	if oneDotSevenDotZero == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.7"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.7.13",
		"kube-proxy":                           "hyperkube-amd64:v1.7.13",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(oneDotSevenDotZero).To(Equal(expected))

	override := getK8sVersionComponents("1.9.3", map[string]string{"windowszip": "v1.9.3-2int.zip"})
	if override == nil {
		t.Fatalf("getK8sVersionComponents() should not return nil for valid version")
	}
	k8sComponent = k8sComponentVersions["1.9"]
	expected = map[string]string{
		"hyperkube":                            "hyperkube-amd64:v1.9.3",
		"kube-proxy":                           "hyperkube-amd64:v1.9.3",
		"ccm":                                  "cloud-controller-manager-amd64:v1.9.3",
		"windowszip":                           "v1.9.3-2int.zip",
		common.DashboardAddonName:              k8sComponent["dashboard"],
		"exechealthz":                          k8sComponent["exechealthz"],
		"addonresizer":                         k8sComponent["addon-resizer"],
		"heapster":                             k8sComponent["heapster"],
		common.MetricsServerAddonName:          k8sComponent["metrics-server"],
		"kube-dns":                             k8sComponent["kube-dns"],
		"addonmanager":                         k8sComponent["addon-manager"],
		"dnsmasq":                              k8sComponent["dnsmasq"],
		"pause":                                k8sComponent["pause"],
		common.TillerAddonName:                 k8sComponent["tiller"],
		common.ReschedulerAddonName:            k8sComponent["rescheduler"],
		common.ACIConnectorAddonName:           k8sComponent["aci-connector"],
		common.ContainerMonitoringAddonName:    k8sComponent[common.ContainerMonitoringAddonName],
		common.AzureCNINetworkMonitorAddonName: k8sComponent[common.AzureCNINetworkMonitorAddonName],
		common.ClusterAutoscalerAddonName:      k8sComponent["cluster-autoscaler"],
		"k8s-dns-sidecar":                      k8sComponent["k8s-dns-sidecar"],
		common.BlobfuseFlexVolumeAddonName:     k8sComponent[common.BlobfuseFlexVolumeAddonName],
		common.SMBFlexVolumeAddonName:          k8sComponent[common.SMBFlexVolumeAddonName],
		"nodestatusfreq":                       k8sComponent["nodestatusfreq"],
		"nodegraceperiod":                      k8sComponent["nodegraceperiod"],
		"podeviction":                          k8sComponent["podeviction"],
		"routeperiod":                          k8sComponent["routeperiod"],
		"backoffretries":                       k8sComponent["backoffretries"],
		"backoffjitter":                        k8sComponent["backoffjitter"],
		"backoffduration":                      k8sComponent["backoffduration"],
		"backoffexponent":                      k8sComponent["backoffexponent"],
		"ratelimitqps":                         k8sComponent["ratelimitqps"],
		"ratelimitqpswrite":                    k8sComponent["ratelimitqpswrite"],
		"ratelimitbucket":                      k8sComponent["ratelimitbucket"],
		"ratelimitbucketwrite":                 k8sComponent["ratelimitbucketwrite"],
		"gchighthreshold":                      k8sComponent["gchighthreshold"],
		"gclowthreshold":                       k8sComponent["gclowthreshold"],
	}
	g.Expect(override).To(Equal(expected))

	unknown := getK8sVersionComponents("1.0.0", nil)
	if unknown != nil {
		t.Fatalf("getK8sVersionComponents() should return nil for unknown k8s version")
	}
}
