// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

//AzureEnvironmentSpecConfig is the overall configuration differences in different cloud environments.
type AzureEnvironmentSpecConfig struct {
	CloudName            string                        `json:"cloudName,omitempty"`
	DockerSpecConfig     DockerSpecConfig              `json:"dockerSpecConfig,omitempty"`
	KubernetesSpecConfig KubernetesSpecConfig          `json:"kubernetesSpecConfig,omitempty"`
	DCOSSpecConfig       DCOSSpecConfig                `json:"-"`
	EndpointConfig       AzureEndpointConfig           `json:"endpointConfig,omitempty"`
	OSImageConfig        map[Distro]AzureOSImageConfig `json:"osImageConfig,omitempty"`
}

//DockerSpecConfig is the configurations of docker
type DockerSpecConfig struct {
	DockerEngineRepo         string `json:"dockerEngineRepo,omitempty"`
	DockerComposeDownloadURL string `json:"dockerComposeDownloadURL,omitempty"`
}

//DCOSSpecConfig is the configurations of DCOS
type DCOSSpecConfig struct {
	DCOS188BootstrapDownloadURL     string
	DCOS190BootstrapDownloadURL     string
	DCOS198BootstrapDownloadURL     string
	DCOS110BootstrapDownloadURL     string
	DCOS111BootstrapDownloadURL     string
	DCOSWindowsBootstrapDownloadURL string
	DcosRepositoryURL               string // For custom install, for example CI, need these three addributes
	DcosClusterPackageListID        string // the id of the package list file
	DcosProviderPackageID           string // the id of the dcos-provider-xxx package
}

//KubernetesSpecConfig is the kubernetes container images used.
type KubernetesSpecConfig struct {
	AzureTelemetryPID                string `json:"azureTelemetryPID,omitempty"`
	KubernetesImageBase              string `json:"kubernetesImageBase,omitempty"`
	MCRKubernetesImageBase           string `json:"mcrKubernetesImageBase,omitempty"`
	TillerImageBase                  string `json:"tillerImageBase,omitempty"`
	ACIConnectorImageBase            string `json:"aciConnectorImageBase,omitempty"`
	NVIDIAImageBase                  string `json:"nvidiaImageBase,omitempty"`
	AzureCNIImageBase                string `json:"azureCNIImageBase,omitempty"`
	CalicoImageBase                  string `json:"calicoImageBase,omitempty"`
	EtcdDownloadURLBase              string `json:"etcdDownloadURLBase,omitempty"`
	KubeBinariesSASURLBase           string `json:"kubeBinariesSASURLBase,omitempty"`
	WindowsTelemetryGUID             string `json:"windowsTelemetryGUID,omitempty"`
	CNIPluginsDownloadURL            string `json:"cniPluginsDownloadURL,omitempty"`
	VnetCNILinuxPluginsDownloadURL   string `json:"vnetCNILinuxPluginsDownloadURL,omitempty"`
	VnetCNIWindowsPluginsDownloadURL string `json:"vnetCNIWindowsPluginsDownloadURL,omitempty"`
	ContainerdDownloadURLBase        string `json:"containerdDownloadURLBase,omitempty"`
	CSIProxyDownloadURL              string `json:"csiProxyDownloadURL,omitempty"`
}

//AzureEndpointConfig describes an Azure endpoint
type AzureEndpointConfig struct {
	ResourceManagerVMDNSSuffix string `json:"resourceManagerVMDNSSuffix,omitempty"`
}

//AzureOSImageConfig describes an Azure OS image
type AzureOSImageConfig struct {
	ImageOffer     string `json:"imageOffer,omitempty"`
	ImageSku       string `json:"imageSku,omitempty"`
	ImagePublisher string `json:"imagePublisher,omitempty"`
	ImageVersion   string `json:"imageVersion,omitempty"`
}
