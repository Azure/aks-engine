// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

// AzureEnvironmentSpecConfig is the overall configuration differences in different cloud environments.
type AzureEnvironmentSpecConfig struct {
	CloudName            string                        `json:"cloudName,omitempty"`
	KubernetesSpecConfig KubernetesSpecConfig          `json:"kubernetesSpecConfig,omitempty"`
	EndpointConfig       AzureEndpointConfig           `json:"endpointConfig,omitempty"`
	OSImageConfig        map[Distro]AzureOSImageConfig `json:"osImageConfig,omitempty"`
}

// KubernetesSpecConfig is the kubernetes container images used.
type KubernetesSpecConfig struct {
	AzureTelemetryPID                    string `json:"azureTelemetryPID,omitempty"`
	KubernetesImageBase                  string `json:"kubernetesImageBase,omitempty"`
	MCRKubernetesImageBase               string `json:"mcrKubernetesImageBase,omitempty"`
	TillerImageBase                      string `json:"tillerImageBase,omitempty"`
	ACIConnectorImageBase                string `json:"aciConnectorImageBase,omitempty"` // Deprecated
	NVIDIAImageBase                      string `json:"nvidiaImageBase,omitempty"`
	AzureCNIImageBase                    string `json:"azureCNIImageBase,omitempty"`
	CalicoImageBase                      string `json:"calicoImageBase,omitempty"`
	EtcdDownloadURLBase                  string `json:"etcdDownloadURLBase,omitempty"`
	KubeBinariesSASURLBase               string `json:"kubeBinariesSASURLBase,omitempty"`
	WindowsTelemetryGUID                 string `json:"windowsTelemetryGUID,omitempty"`
	CNIPluginsDownloadURL                string `json:"cniPluginsDownloadURL,omitempty"`
	VnetCNILinuxPluginsDownloadURL       string `json:"vnetCNILinuxPluginsDownloadURL,omitempty"`
	VnetCNIWindowsPluginsDownloadURL     string `json:"vnetCNIWindowsPluginsDownloadURL,omitempty"`
	ContainerdDownloadURLBase            string `json:"containerdDownloadURLBase,omitempty"`
	CSIProxyDownloadURL                  string `json:"csiProxyDownloadURL,omitempty"`
	WindowsProvisioningScriptsPackageURL string `json:"windowsProvisioningScriptsPackageURL,omitempty"`
	WindowsPauseImageURL                 string `json:"windowsPauseImageURL,omitempty"`
	AlwaysPullWindowsPauseImage          bool   `json:"alwaysPullWindowsPauseImage,omitempty"`
}

// AzureEndpointConfig describes an Azure endpoint
type AzureEndpointConfig struct {
	ResourceManagerVMDNSSuffix string `json:"resourceManagerVMDNSSuffix,omitempty"`
}

// AzureOSImageConfig describes an Azure OS image
type AzureOSImageConfig struct {
	ImageOffer     string `json:"imageOffer,omitempty"`
	ImageSku       string `json:"imageSku,omitempty"`
	ImagePublisher string `json:"imagePublisher,omitempty"`
	ImageVersion   string `json:"imageVersion,omitempty"`
}
