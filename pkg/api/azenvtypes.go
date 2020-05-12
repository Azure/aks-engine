// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import "fmt"

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
	AzureTelemetryPID string `json:"azureTelemetryPID,omitempty"`
	// KubernetesImageBase defines a base image URL substring to source images that originate from upstream k8s.gcr.io
	KubernetesImageBase   string `json:"kubernetesImageBase,omitempty"`
	TillerImageBase       string `json:"tillerImageBase,omitempty"`
	ACIConnectorImageBase string `json:"aciConnectorImageBase,omitempty"`
	// MCRKubernetesImageBase defines a base image URL substring to source MS-curated images that originate from MCR
	MCRKubernetesImageBase           string `json:"mcrKubernetesImageBase,omitempty"`
	NVIDIAImageBase                  string `json:"nvidiaImageBase,omitempty"`
	AzureCNIImageBase                string `json:"azureCNIImageBase,omitempty"`
	CalicoImageBase                  string `json:"CalicoImageBase,omitempty"`
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

// AzureTelemetryPID represents the current telemetry ID
// See more information here https://docs.microsoft.com/en-us/azure/marketplace/azure-partner-customer-usage-attribution
// PID is maintained to keep consistent with Azure Stack Telemetry Terminologies
type AzureTelemetryPID string

const (
	// DefaultAzureStackDeployTelemetryPID tracking ID for Deployment
	DefaultAzureStackDeployTelemetryPID = "pid-1bda96ec-adf4-4eea-bb9a-8462de5475c0"
	// DefaultAzureStackScaleTelemetryPID tracking ID for Scale
	DefaultAzureStackScaleTelemetryPID = "pid-bbbafa53-d6a7-4022-84a2-86fcbaec7030"
	// DefaultAzureStackUpgradeTelemetryPID tracking ID for Upgrade
	DefaultAzureStackUpgradeTelemetryPID = "pid-0d9b5198-7cd7-4252-a890-5658eaf874be"
)

var (
	//DefaultKubernetesSpecConfig is the default Docker image source of Kubernetes
	DefaultKubernetesSpecConfig = KubernetesSpecConfig{
		KubernetesImageBase:              "k8s.gcr.io/",
		TillerImageBase:                  "mcr.microsoft.com/",
		ACIConnectorImageBase:            "microsoft/",
		NVIDIAImageBase:                  "mcr.microsoft.com/",
		CalicoImageBase:                  "calico/",
		AzureCNIImageBase:                "mcr.microsoft.com/containernetworking/",
		MCRKubernetesImageBase:           "mcr.microsoft.com/",
		EtcdDownloadURLBase:              "mcr.microsoft.com/oss/etcd-io/",
		KubeBinariesSASURLBase:           "https://kubernetesartifacts.azureedge.net/kubernetes/",
		WindowsTelemetryGUID:             "fb801154-36b9-41bc-89c2-f4d4f05472b0",
		CNIPluginsDownloadURL:            "https://kubernetesartifacts.azureedge.net/cni-plugins/" + CNIPluginVer + "/binaries/cni-plugins-linux-amd64-" + CNIPluginVer + ".tgz",
		VnetCNILinuxPluginsDownloadURL:   "https://kubernetesartifacts.azureedge.net/azure-cni/" + AzureCniPluginVerLinux + "/binaries/azure-vnet-cni-linux-amd64-" + AzureCniPluginVerLinux + ".tgz",
		VnetCNIWindowsPluginsDownloadURL: "https://kubernetesartifacts.azureedge.net/azure-cni/" + AzureCniPluginVerWindows + "/binaries/azure-vnet-cni-singletenancy-windows-amd64-" + AzureCniPluginVerWindows + ".zip",
		ContainerdDownloadURLBase:        "https://storage.googleapis.com/cri-containerd-release/",
		CSIProxyDownloadURL:              "https://kubernetesartifacts.azureedge.net/csi-proxy/v0.1.0/binaries/csi-proxy.tar.gz",
	}

	//DefaultDCOSSpecConfig is the default DC/OS binary download URL.
	DefaultDCOSSpecConfig = DCOSSpecConfig{
		DCOS188BootstrapDownloadURL:     fmt.Sprintf(AzureEdgeDCOSBootstrapDownloadURL, "stable", "5df43052907c021eeb5de145419a3da1898c58a5"),
		DCOS190BootstrapDownloadURL:     fmt.Sprintf(AzureEdgeDCOSBootstrapDownloadURL, "stable", "58fd0833ce81b6244fc73bf65b5deb43217b0bd7"),
		DCOS198BootstrapDownloadURL:     fmt.Sprintf(AzureEdgeDCOSBootstrapDownloadURL, "stable/1.9.8", "f4ae0d20665fc68ee25282d6f78681b2773c6e10"),
		DCOS110BootstrapDownloadURL:     fmt.Sprintf(AzureEdgeDCOSBootstrapDownloadURL, "stable/1.10.0", "4d92536e7381176206e71ee15b5ffe454439920c"),
		DCOS111BootstrapDownloadURL:     fmt.Sprintf(AzureEdgeDCOSBootstrapDownloadURL, "stable/1.11.0", "a0654657903fb68dff60f6e522a7f241c1bfbf0f"),
		DCOSWindowsBootstrapDownloadURL: "http://dcos-win.westus.cloudapp.azure.com/dcos-windows/stable/",
		DcosRepositoryURL:               "https://dcosio.azureedge.net/dcos/stable/1.11.0",
		DcosClusterPackageListID:        "248a66388bba1adbcb14a52fd3b7b424ab06fa76",
	}

	//DefaultDockerSpecConfig is the default Docker engine repo.
	DefaultDockerSpecConfig = DockerSpecConfig{
		DockerEngineRepo:         "https://aptdocker.azureedge.net/repo",
		DockerComposeDownloadURL: "https://github.com/docker/compose/releases/download",
	}

	//Ubuntu1604OSImageConfig is the default Linux distribution.
	Ubuntu1604OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "UbuntuServer",
		ImageSku:       "16.04-LTS",
		ImagePublisher: "Canonical",
		ImageVersion:   "latest",
	}

	//Ubuntu1804OSImageConfig is the Ubunutu 18.04-LTS Linux distribution.
	Ubuntu1804OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "UbuntuServer",
		ImageSku:       "18.04-LTS",
		ImagePublisher: "Canonical",
		ImageVersion:   "latest",
	}

	//Ubuntu1804Gen2OSImageConfig is Gen2 flavor the Ubunutu 18.04-LTS Linux distribution.
	Ubuntu1804Gen2OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "UbuntuServer",
		ImageSku:       "18_04-lts-gen2",
		ImagePublisher: "Canonical",
		ImageVersion:   "latest",
	}

	//RHELOSImageConfig is the RHEL Linux distribution.
	RHELOSImageConfig = AzureOSImageConfig{
		ImageOffer:     "RHEL",
		ImageSku:       "7.3",
		ImagePublisher: "RedHat",
		ImageVersion:   "latest",
	}

	// AKSUbuntu1604OSImageConfig is the AKS image based on Ubuntu 16.04-LTS.
	AKSUbuntu1604OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "aks",
		ImageSku:       "aks-engine-ubuntu-1604-202004",
		ImagePublisher: "microsoft-aks",
		ImageVersion:   "2020.04.21",
	}

	// AKSUbuntu1804OSImageConfig is the AKS image based on Ubuntu 18.04-LTS.
	AKSUbuntu1804OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "aks",
		ImageSku:       "aks-engine-ubuntu-1804-202004",
		ImagePublisher: "microsoft-aks",
		ImageVersion:   "2020.04.21",
	}

	// AKSWindowsServer2019OSImageConfig is the AKS image based on Windows Server 2019
	AKSWindowsServer2019OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "aks-windows",
		ImageSku:       "2019-datacenter-core-smalldisk-2004",
		ImagePublisher: "microsoft-aks",
		ImageVersion:   "17763.1158.200421",
	}

	// WindowsServer2019OSImageConfig is the 'vanilla' Windows Server 2019 image
	WindowsServer2019OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "WindowsServer",
		ImageSku:       "2019-Datacenter-Core-with-Containers-smalldisk",
		ImagePublisher: "MicrosoftWindowsServer",
		ImageVersion:   "17763.1158.2004131759",
	}

	// ACC1604OSImageConfig is the ACC image based on Ubuntu 16.04.
	ACC1604OSImageConfig = AzureOSImageConfig{
		ImageOffer:     "confidential-compute-preview",
		ImageSku:       "16.04-LTS",
		ImagePublisher: "Canonical",
		ImageVersion:   "latest",
	}

	//AzureCloudSpec is the default configurations for global azure.
	AzureCloudSpec = AzureEnvironmentSpecConfig{
		CloudName: AzurePublicCloud,
		//DockerSpecConfig specify the docker engine download repo
		DockerSpecConfig: DefaultDockerSpecConfig,
		//KubernetesSpecConfig is the default kubernetes container image url.
		KubernetesSpecConfig: DefaultKubernetesSpecConfig,
		DCOSSpecConfig:       DefaultDCOSSpecConfig,

		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "cloudapp.azure.com",
		},

		OSImageConfig: map[Distro]AzureOSImageConfig{
			Ubuntu:            Ubuntu1604OSImageConfig,
			Ubuntu1804:        Ubuntu1804OSImageConfig,
			Ubuntu1804Gen2:    Ubuntu1804Gen2OSImageConfig,
			RHEL:              RHELOSImageConfig,
			AKSUbuntu1604:     AKSUbuntu1604OSImageConfig,
			AKS1604Deprecated: AKSUbuntu1604OSImageConfig, // for back-compat
			AKSUbuntu1804:     AKSUbuntu1804OSImageConfig,
			AKS1804Deprecated: AKSUbuntu1804OSImageConfig, // for back-compat
			ACC1604:           ACC1604OSImageConfig,
		},
	}

	//AzureGermanCloudSpec is the German cloud config.
	AzureGermanCloudSpec = AzureEnvironmentSpecConfig{
		CloudName:            AzureGermanCloud,
		DockerSpecConfig:     DefaultDockerSpecConfig,
		KubernetesSpecConfig: DefaultKubernetesSpecConfig,
		DCOSSpecConfig:       DefaultDCOSSpecConfig,
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "cloudapp.microsoftazure.de",
		},
		OSImageConfig: map[Distro]AzureOSImageConfig{
			Ubuntu:            Ubuntu1604OSImageConfig,
			Ubuntu1804:        Ubuntu1804OSImageConfig,
			Ubuntu1804Gen2:    Ubuntu1804Gen2OSImageConfig,
			RHEL:              RHELOSImageConfig,
			AKSUbuntu1604:     Ubuntu1604OSImageConfig,
			AKS1604Deprecated: Ubuntu1604OSImageConfig, // for back-compat
			AKSUbuntu1804:     Ubuntu1604OSImageConfig, // workaround for https://github.com/Azure/aks-engine/issues/761
			AKS1804Deprecated: Ubuntu1604OSImageConfig, // for back-compat
		},
	}

	//AzureUSGovernmentCloudSpec is the US government config.
	AzureUSGovernmentCloudSpec = AzureEnvironmentSpecConfig{
		CloudName:            AzureUSGovernmentCloud,
		DockerSpecConfig:     DefaultDockerSpecConfig,
		KubernetesSpecConfig: DefaultKubernetesSpecConfig,
		DCOSSpecConfig:       DefaultDCOSSpecConfig,
		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "cloudapp.usgovcloudapi.net",
		},
		OSImageConfig: map[Distro]AzureOSImageConfig{
			Ubuntu:            Ubuntu1604OSImageConfig,
			Ubuntu1804:        Ubuntu1804OSImageConfig,
			Ubuntu1804Gen2:    Ubuntu1804Gen2OSImageConfig,
			RHEL:              RHELOSImageConfig,
			AKSUbuntu1604:     AKSUbuntu1604OSImageConfig,
			AKS1604Deprecated: AKSUbuntu1604OSImageConfig, // for back-compat
			AKSUbuntu1804:     AKSUbuntu1804OSImageConfig,
			AKS1804Deprecated: AKSUbuntu1804OSImageConfig, // for back-compat
		},
	}

	//AzureChinaCloudSpec is the configurations for Azure China (Mooncake)
	AzureChinaCloudSpec = AzureEnvironmentSpecConfig{
		CloudName: AzureChinaCloud,
		//DockerSpecConfig specify the docker engine download repo
		DockerSpecConfig: DockerSpecConfig{
			DockerEngineRepo:         "https://mirror.azk8s.cn/docker-engine/apt/repo/",
			DockerComposeDownloadURL: "https://mirror.azk8s.cn/docker-toolbox/linux/compose",
		},
		//KubernetesSpecConfig - Due to Chinese firewall issue, the default containers from google is blocked, use the Chinese local mirror instead
		KubernetesSpecConfig: KubernetesSpecConfig{
			KubernetesImageBase:              "gcr.azk8s.cn/google_containers/",
			TillerImageBase:                  "mcr.microsoft.com/",
			ACIConnectorImageBase:            "dockerhub.azk8s.cn/microsoft/",
			NVIDIAImageBase:                  "dockerhub.azk8s.cn/nvidia/",
			AzureCNIImageBase:                "dockerhub.azk8s.cn/containernetworking/",
			MCRKubernetesImageBase:           "mcr.microsoft.com/",
			CalicoImageBase:                  "dockerhub.azk8s.cn/calico/",
			EtcdDownloadURLBase:              "mcr.microsoft.com/oss/etcd-io/",
			KubeBinariesSASURLBase:           DefaultKubernetesSpecConfig.KubeBinariesSASURLBase,
			WindowsTelemetryGUID:             DefaultKubernetesSpecConfig.WindowsTelemetryGUID,
			CNIPluginsDownloadURL:            "https://mirror.azk8s.cn/kubernetes/containernetworking-plugins/cni-plugins-linux-amd64-" + CNIPluginVer + ".tgz",
			VnetCNILinuxPluginsDownloadURL:   "https://mirror.azk8s.cn/kubernetes/azure-container-networking/azure-vnet-cni-linux-amd64-" + AzureCniPluginVerLinux + ".tgz",
			VnetCNIWindowsPluginsDownloadURL: "https://mirror.azk8s.cn/kubernetes/azure-container-networking/azure-vnet-cni-windows-amd64-" + AzureCniPluginVerWindows + ".zip",
			ContainerdDownloadURLBase:        "https://mirror.azk8s.cn/kubernetes/containerd/",
			CSIProxyDownloadURL:              "https://kubernetesartifacts.blob.core.chinacloudapi.cn/csi-proxy/v0.1.0/binaries/csi-proxy.tar.gz",
		},
		DCOSSpecConfig: DCOSSpecConfig{
			DCOS188BootstrapDownloadURL:     fmt.Sprintf(AzureChinaCloudDCOSBootstrapDownloadURL, "5df43052907c021eeb5de145419a3da1898c58a5"),
			DCOSWindowsBootstrapDownloadURL: "https://dcosdevstorage.blob.core.windows.net/dcos-windows",
			DCOS190BootstrapDownloadURL:     fmt.Sprintf(AzureChinaCloudDCOSBootstrapDownloadURL, "58fd0833ce81b6244fc73bf65b5deb43217b0bd7"),
			DCOS198BootstrapDownloadURL:     fmt.Sprintf(AzureChinaCloudDCOSBootstrapDownloadURL, "f4ae0d20665fc68ee25282d6f78681b2773c6e10"),
		},

		EndpointConfig: AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "cloudapp.chinacloudapi.cn",
		},
		OSImageConfig: map[Distro]AzureOSImageConfig{
			Ubuntu:            Ubuntu1604OSImageConfig,
			Ubuntu1804:        Ubuntu1804OSImageConfig,
			Ubuntu1804Gen2:    Ubuntu1804Gen2OSImageConfig,
			RHEL:              RHELOSImageConfig,
			AKSUbuntu1604:     AKSUbuntu1604OSImageConfig,
			AKS1604Deprecated: AKSUbuntu1604OSImageConfig, // for back-compat
			AKSUbuntu1804:     AKSUbuntu1804OSImageConfig,
			AKS1804Deprecated: AKSUbuntu1804OSImageConfig, // for back-compat
		},
	}

	// AzureCloudSpecEnvMap is the environment configuration map for all the Azure cloud environments.
	AzureCloudSpecEnvMap = map[string]AzureEnvironmentSpecConfig{
		AzureChinaCloud:        AzureChinaCloudSpec,
		AzureGermanCloud:       AzureGermanCloudSpec,
		AzureUSGovernmentCloud: AzureUSGovernmentCloudSpec,
		AzurePublicCloud:       AzureCloudSpec,
	}
)
