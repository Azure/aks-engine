package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/i18n"
)

// VlabsContainerService is the type we read and write from file
// needed because the json that is sent to ARM and aks-engine
// is different from the json that the ACS RP Api gets from ARM
type VlabsContainerService struct {
	api.TypeMeta
	*vlabs.ContainerService
}

//DockerSpecConfig is the configurations of docker
type DockerSpecConfig struct {
	DockerEngineRepo         string
	DockerComposeDownloadURL string
}

//KubernetesSpecConfig is the kubernetes container images used.
type KubernetesSpecConfig struct {
	KubernetesImageBase              string
	TillerImageBase                  string
	ACIConnectorImageBase            string
	NVIDIAImageBase                  string
	AzureCNIImageBase                string
	EtcdDownloadURLBase              string
	KubeBinariesSASURLBase           string
	WindowsTelemetryGUID             string
	CNIPluginsDownloadURL            string
	VnetCNILinuxPluginsDownloadURL   string
	VnetCNIWindowsPluginsDownloadURL string
	ContainerdDownloadURLBase        string
}

//AzureEndpointConfig describes an Azure endpoint
type AzureEndpointConfig struct {
	ResourceManagerVMDNSSuffix string
}

//AzureOSImageConfig describes an Azure OS image
type AzureOSImageConfig struct {
	ImageOffer     string
	ImageSku       string
	ImagePublisher string
	ImageVersion   string
}

// Context represents the object that is passed to the package
type Context struct {
	Translator *i18n.Translator
}

// KeyVaultID represents a KeyVault instance on Azure
type KeyVaultID struct {
	ID string `json:"id"`
}

// KeyVaultRef represents a reference to KeyVault instance on Azure
type KeyVaultRef struct {
	KeyVault      KeyVaultID `json:"keyVault"`
	SecretName    string     `json:"secretName"`
	SecretVersion string     `json:"secretVersion,omitempty"`
}

type paramsMap map[string]interface{}
