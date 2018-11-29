// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"encoding/json"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/pkg/errors"
)

// ResourcePurchasePlan defines resource plan as required by ARM
// for billing purposes.
type ResourcePurchasePlan struct {
	Name          string `json:"name,omitempty"`
	Product       string `json:"product,omitempty"`
	PromotionCode string `json:"promotionCode,omitempty"`
	Publisher     string `json:"publisher,omitempty"`
}

// ContainerService complies with the ARM model of
// resource definition in a JSON template.
type ContainerService struct {
	ID       string                `json:"id,omitempty"`
	Location string                `json:"location,omitempty"`
	Name     string                `json:"name,omitempty"`
	Plan     *ResourcePurchasePlan `json:"plan,omitempty"`
	Tags     map[string]string     `json:"tags,omitempty"`
	Type     string                `json:"type,omitempty"`

	Properties *Properties `json:"properties"`
}

// Properties represents the AKS cluster definition
type Properties struct {
	ProvisioningState       ProvisioningState        `json:"provisioningState,omitempty"`
	OrchestratorProfile     *OrchestratorProfile     `json:"orchestratorProfile,omitempty" validate:"required"`
	MasterProfile           *MasterProfile           `json:"masterProfile,omitempty" validate:"required"`
	AgentPoolProfiles       []*AgentPoolProfile      `json:"agentPoolProfiles,omitempty" validate:"dive,required"`
	LinuxProfile            *LinuxProfile            `json:"linuxProfile,omitempty" validate:"required"`
	ExtensionProfiles       []*ExtensionProfile      `json:"extensionProfiles,omitempty"`
	WindowsProfile          *WindowsProfile          `json:"windowsProfile,omitempty"`
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
	CertificateProfile      *CertificateProfile      `json:"certificateProfile,omitempty"`
	AADProfile              *AADProfile              `json:"aadProfile,omitempty"`
	FeatureFlags            *FeatureFlags            `json:"featureFlags,omitempty"`
}

// FeatureFlags defines feature-flag restricted functionality
type FeatureFlags struct {
	EnableCSERunInBackground bool `json:"enableCSERunInBackground,omitempty"`
	BlockOutboundInternet    bool `json:"blockOutboundInternet,omitempty"`
}

// ServicePrincipalProfile contains the client and secret used by the cluster for Azure Resource CRUD
// The 'Secret' and 'KeyvaultSecretRef' parameters are mutually exclusive
// The 'Secret' parameter should be a secret in plain text.
// The 'KeyvaultSecretRef' parameter is a reference to a secret in a keyvault.
type ServicePrincipalProfile struct {
	ClientID          string             `json:"clientId,omitempty"`
	Secret            string             `json:"secret,omitempty"`
	ObjectID          string             `json:"objectId,omitempty"`
	KeyvaultSecretRef *KeyvaultSecretRef `json:"keyvaultSecretRef,omitempty"`
}

// KeyvaultSecretRef is a reference to a secret in a keyvault.
// The format of 'VaultID' value should be
// "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>"
// where:
//    <SUB_ID> is the subscription ID of the keyvault
//    <RG_NAME> is the resource group of the keyvault
//    <KV_NAME> is the name of the keyvault
// The 'SecretName' is the name of the secret in the keyvault
// The 'SecretVersion' (optional) is the version of the secret (default: the latest version)
type KeyvaultSecretRef struct {
	VaultID       string `json:"vaultID" validate:"required"`
	SecretName    string `json:"secretName" validate:"required"`
	SecretVersion string `json:"version,omitempty"`
}

// CertificateProfile represents the definition of the master cluster
// The JSON parameters could be either a plain text, or referenced to a secret in a keyvault.
// In the latter case, the format of the parameter's value should be
// "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<NAME>[/<VERSION>]"
// where:
//    <SUB_ID> is the subscription ID of the keyvault
//    <RG_NAME> is the resource group of the keyvault
//    <KV_NAME> is the name of the keyvault
//    <NAME> is the name of the secret
//    <VERSION> (optional) is the version of the secret (default: the latest version)
type CertificateProfile struct {
	// CaCertificate is the certificate authority certificate.
	CaCertificate string `json:"caCertificate,omitempty"`
	// CaPrivateKey is the certificate authority key.
	CaPrivateKey string `json:"caPrivateKey,omitempty"`
	// ApiServerCertificate is the rest api server certificate, and signed by the CA
	APIServerCertificate string `json:"apiServerCertificate,omitempty"`
	// ApiServerPrivateKey is the rest api server private key, and signed by the CA
	APIServerPrivateKey string `json:"apiServerPrivateKey,omitempty"`
	// ClientCertificate is the certificate used by the client kubelet services and signed by the CA
	ClientCertificate string `json:"clientCertificate,omitempty"`
	// ClientPrivateKey is the private key used by the client kubelet services and signed by the CA
	ClientPrivateKey string `json:"clientPrivateKey,omitempty"`
	// KubeConfigCertificate is the client certificate used for kubectl cli and signed by the CA
	KubeConfigCertificate string `json:"kubeConfigCertificate,omitempty"`
	// KubeConfigPrivateKey is the client private key used for kubectl cli and signed by the CA
	KubeConfigPrivateKey string `json:"kubeConfigPrivateKey,omitempty"`
	// EtcdServerCertificate is the server certificate for etcd, and signed by the CA
	EtcdServerCertificate string `json:"etcdServerCertificate,omitempty"`
	// EtcdServerPrivateKey is the server private key for etcd, and signed by the CA
	EtcdServerPrivateKey string `json:"etcdServerPrivateKey,omitempty"`
	// EtcdClientCertificate is etcd client certificate, and signed by the CA
	EtcdClientCertificate string `json:"etcdClientCertificate,omitempty"`
	// EtcdClientPrivateKey is the etcd client private key, and signed by the CA
	EtcdClientPrivateKey string `json:"etcdClientPrivateKey,omitempty"`
	// EtcdPeerCertificates is list of etcd peer certificates, and signed by the CA
	EtcdPeerCertificates []string `json:"etcdPeerCertificates,omitempty"`
	// EtcdPeerPrivateKeys is list of etcd peer private keys, and signed by the CA
	EtcdPeerPrivateKeys []string `json:"etcdPeerPrivateKeys,omitempty"`
}

// LinuxProfile represents the linux parameters passed to the cluster
type LinuxProfile struct {
	AdminUsername string `json:"adminUsername" validate:"required"`
	SSH           struct {
		PublicKeys []PublicKey `json:"publicKeys" validate:"required,len=1"`
	} `json:"ssh" validate:"required"`
	Secrets            []KeyVaultSecrets   `json:"secrets,omitempty"`
	ScriptRootURL      string              `json:"scriptroot,omitempty"`
	CustomSearchDomain *CustomSearchDomain `json:"customSearchDomain,omitempty"`
	CustomNodesDNS     *CustomNodesDNS     `json:"customNodesDNS,omitempty"`
}

// PublicKey represents an SSH key for LinuxProfile
type PublicKey struct {
	KeyData string `json:"keyData"`
}

// CustomSearchDomain represents the Search Domain when the custom vnet has a windows server DNS as a nameserver.
type CustomSearchDomain struct {
	Name          string `json:"name,omitempty"`
	RealmUser     string `json:"realmUser,omitempty"`
	RealmPassword string `json:"realmPassword,omitempty"`
}

// CustomNodesDNS represents the Search Domain
type CustomNodesDNS struct {
	DNSServer string `json:"dnsServer,omitempty"`
}

// WindowsProfile represents the windows parameters passed to the cluster
type WindowsProfile struct {
	AdminUsername         string            `json:"adminUsername,omitempty"`
	AdminPassword         string            `json:"adminPassword,omitempty"`
	ImageVersion          string            `json:"imageVersion,omitempty"`
	WindowsImageSourceURL string            `json:"WindowsImageSourceUrl"`
	WindowsPublisher      string            `json:"WindowsPublisher"`
	WindowsOffer          string            `json:"WindowsOffer"`
	WindowsSku            string            `json:"WindowsSku"`
	WindowsDockerVersion  string            `json:"windowsDockerVersion"`
	Secrets               []KeyVaultSecrets `json:"secrets,omitempty"`
}

// ProvisioningState represents the current state of container service resource.
type ProvisioningState string

const (
	// Creating means ContainerService resource is being created.
	Creating ProvisioningState = "Creating"
	// Updating means an existing ContainerService resource is being updated
	Updating ProvisioningState = "Updating"
	// Failed means resource is in failed state
	Failed ProvisioningState = "Failed"
	// Succeeded means resource created succeeded during last create/update
	Succeeded ProvisioningState = "Succeeded"
	// Deleting means resource is in the process of being deleted
	Deleting ProvisioningState = "Deleting"
	// Migrating means resource is being migrated from one subscription or
	// resource group to another
	Migrating ProvisioningState = "Migrating"
)

// OrchestratorProfile contains Orchestrator properties
type OrchestratorProfile struct {
	OrchestratorType    string            `json:"orchestratorType" validate:"required"`
	OrchestratorRelease string            `json:"orchestratorRelease,omitempty"`
	OrchestratorVersion string            `json:"orchestratorVersion,omitempty"`
	KubernetesConfig    *KubernetesConfig `json:"kubernetesConfig,omitempty"`
	DcosConfig          *DcosConfig       `json:"dcosConfig,omitempty"`
}

// UnmarshalJSON unmarshal json using the default behavior
// And do fields manipulation, such as populating default value
func (o *OrchestratorProfile) UnmarshalJSON(b []byte) error {
	// Need to have a alias type to avoid circular unmarshal
	type aliasOrchestratorProfile OrchestratorProfile
	op := aliasOrchestratorProfile{}
	if e := json.Unmarshal(b, &op); e != nil {
		return e
	}
	*o = OrchestratorProfile(op)
	// Unmarshal OrchestratorType, format it as well
	orchestratorType := o.OrchestratorType
	switch {
	case strings.EqualFold(orchestratorType, DCOS):
		o.OrchestratorType = DCOS
	case strings.EqualFold(orchestratorType, Swarm):
		o.OrchestratorType = Swarm
	case strings.EqualFold(orchestratorType, Kubernetes):
		o.OrchestratorType = Kubernetes
	case strings.EqualFold(orchestratorType, SwarmMode):
		o.OrchestratorType = SwarmMode
	default:
		return errors.Errorf("OrchestratorType has unknown orchestrator: %s", orchestratorType)
	}
	return nil
}

// KubernetesContainerSpec defines configuration for a container spec
type KubernetesContainerSpec struct {
	Name           string `json:"name,omitempty"`
	Image          string `json:"image,omitempty"`
	CPURequests    string `json:"cpuRequests,omitempty"`
	MemoryRequests string `json:"memoryRequests,omitempty"`
	CPULimits      string `json:"cpuLimits,omitempty"`
	MemoryLimits   string `json:"memoryLimits,omitempty"`
}

// KubernetesAddon defines a list of addons w/ configuration to include with the cluster deployment
type KubernetesAddon struct {
	Name       string                    `json:"name,omitempty"`
	Enabled    *bool                     `json:"enabled,omitempty"`
	Containers []KubernetesContainerSpec `json:"containers,omitempty"`
	Config     map[string]string         `json:"config,omitempty"`
	Data       string                    `json:"data,omitempty"`
}

// IsEnabled returns if the addon is explicitly enabled, or the user-provided default if non explicitly enabled
func (a *KubernetesAddon) IsEnabled(ifNil bool) bool {
	if a.Enabled == nil {
		return ifNil
	}
	return *a.Enabled
}

// PrivateCluster defines the configuration for a private cluster
type PrivateCluster struct {
	Enabled        *bool                  `json:"enabled,omitempty"`
	JumpboxProfile *PrivateJumpboxProfile `json:"jumpboxProfile,omitempty"`
}

// PrivateJumpboxProfile represents a jumpbox definition
type PrivateJumpboxProfile struct {
	Name           string `json:"name" validate:"required"`
	VMSize         string `json:"vmSize" validate:"required"`
	OSDiskSizeGB   int    `json:"osDiskSizeGB,omitempty" validate:"min=0,max=1023"`
	Username       string `json:"username,omitempty"`
	PublicKey      string `json:"publicKey" validate:"required"`
	StorageProfile string `json:"storageProfile,omitempty"`
}

// KubernetesConfig contains the Kubernetes config structure, containing
// Kubernetes specific configuration
type KubernetesConfig struct {
	KubernetesImageBase             string            `json:"kubernetesImageBase,omitempty"`
	ClusterSubnet                   string            `json:"clusterSubnet,omitempty"`
	DNSServiceIP                    string            `json:"dnsServiceIP,omitempty"`
	ServiceCidr                     string            `json:"serviceCidr,omitempty"`
	NetworkPolicy                   string            `json:"networkPolicy,omitempty"`
	NetworkPlugin                   string            `json:"networkPlugin,omitempty"`
	ContainerRuntime                string            `json:"containerRuntime,omitempty"`
	MaxPods                         int               `json:"maxPods,omitempty"`
	DockerBridgeSubnet              string            `json:"dockerBridgeSubnet,omitempty"`
	UseManagedIdentity              bool              `json:"useManagedIdentity,omitempty"`
	UserAssignedID                  string            `json:"userAssignedID,omitempty"`
	UserAssignedClientID            string            `json:"userAssignedClientID,omitempty"` //Note: cannot be provided in config. Used *only* for transferring this to azure.json.
	CustomHyperkubeImage            string            `json:"customHyperkubeImage,omitempty"`
	DockerEngineVersion             string            `json:"dockerEngineVersion,omitempty"` // Deprecated
	CustomCcmImage                  string            `json:"customCcmImage,omitempty"`
	UseCloudControllerManager       *bool             `json:"useCloudControllerManager,omitempty"`
	CustomWindowsPackageURL         string            `json:"customWindowsPackageURL,omitempty"`
	WindowsNodeBinariesURL          string            `json:"windowsNodeBinariesURL,omitempty"`
	UseInstanceMetadata             *bool             `json:"useInstanceMetadata,omitempty"`
	EnableRbac                      *bool             `json:"enableRbac,omitempty"`
	EnableSecureKubelet             *bool             `json:"enableSecureKubelet,omitempty"`
	EnableAggregatedAPIs            bool              `json:"enableAggregatedAPIs,omitempty"`
	PrivateCluster                  *PrivateCluster   `json:"privateCluster,omitempty"`
	GCHighThreshold                 int               `json:"gchighthreshold,omitempty"`
	GCLowThreshold                  int               `json:"gclowthreshold,omitempty"`
	EtcdVersion                     string            `json:"etcdVersion,omitempty"`
	EtcdDiskSizeGB                  string            `json:"etcdDiskSizeGB,omitempty"`
	EtcdEncryptionKey               string            `json:"etcdEncryptionKey,omitempty"`
	EnableDataEncryptionAtRest      *bool             `json:"enableDataEncryptionAtRest,omitempty"`
	EnableEncryptionWithExternalKms *bool             `json:"enableEncryptionWithExternalKms,omitempty"`
	EnablePodSecurityPolicy         *bool             `json:"enablePodSecurityPolicy,omitempty"`
	Addons                          []KubernetesAddon `json:"addons,omitempty"`
	KubeletConfig                   map[string]string `json:"kubeletConfig,omitempty"`
	ControllerManagerConfig         map[string]string `json:"controllerManagerConfig,omitempty"`
	CloudControllerManagerConfig    map[string]string `json:"cloudControllerManagerConfig,omitempty"`
	APIServerConfig                 map[string]string `json:"apiServerConfig,omitempty"`
	SchedulerConfig                 map[string]string `json:"schedulerConfig,omitempty"`
	PodSecurityPolicyConfig         map[string]string `json:"podSecurityPolicyConfig,omitempty"`
	CloudProviderBackoff            *bool             `json:"cloudProviderBackoff,omitempty"`
	CloudProviderBackoffRetries     int               `json:"cloudProviderBackoffRetries,omitempty"`
	CloudProviderBackoffJitter      float64           `json:"cloudProviderBackoffJitter,omitempty"`
	CloudProviderBackoffDuration    int               `json:"cloudProviderBackoffDuration,omitempty"`
	CloudProviderBackoffExponent    float64           `json:"cloudProviderBackoffExponent,omitempty"`
	CloudProviderRateLimit          *bool             `json:"cloudProviderRateLimit,omitempty"`
	CloudProviderRateLimitQPS       float64           `json:"cloudProviderRateLimitQPS,omitempty"`
	CloudProviderRateLimitBucket    int               `json:"cloudProviderRateLimitBucket,omitempty"`
	LoadBalancerSku                 string            `json:"loadBalancerSku,omitempty"`
	ExcludeMasterFromStandardLB     *bool             `json:"excludeMasterFromStandardLB,omitempty"`
	AzureCNIVersion                 string            `json:"azureCNIVersion,omitempty"`
	AzureCNIURLLinux                string            `json:"azureCNIURLLinux,omitempty"`
	AzureCNIURLWindows              string            `json:"azureCNIURLWindows,omitempty"`
}

// CustomFile has source as the full absolute source path to a file and dest
// is the full absolute desired destination path to put the file on a master node
type CustomFile struct {
	Source string `json:"source,omitempty"`
	Dest   string `json:"dest,omitempty"`
}

// BootstrapProfile represents the definition of the DCOS bootstrap node used to deploy the cluster
type BootstrapProfile struct {
	VMSize       string `json:"vmSize,omitempty"`
	OSDiskSizeGB int    `json:"osDiskSizeGB,omitempty"`
	OAuthEnabled bool   `json:"oauthEnabled,omitempty"`
	StaticIP     string `json:"staticIP,omitempty"`
	Subnet       string `json:"subnet,omitempty"`
}

// DcosConfig Configuration for DC/OS
type DcosConfig struct {
	DcosBootstrapURL         string            `json:"dcosBootstrapURL,omitempty"`
	DcosWindowsBootstrapURL  string            `json:"dcosWindowsBootstrapURL,omitempty"`
	Registry                 string            `json:"registry,omitempty"`
	RegistryUser             string            `json:"registryUser,omitempty"`
	RegistryPass             string            `json:"registryPassword,omitempty"`
	DcosRepositoryURL        string            `json:"dcosRepositoryURL,omitempty"`        // For CI use, you need to specify
	DcosClusterPackageListID string            `json:"dcosClusterPackageListID,omitempty"` // all three of these items
	DcosProviderPackageID    string            `json:"dcosProviderPackageID,omitempty"`    // repo url is the location of the build,
	BootstrapProfile         *BootstrapProfile `json:"bootstrapProfile,omitempty"`
}

// MasterProfile represents the definition of the master cluster
type MasterProfile struct {
	Count                    int               `json:"count" validate:"required,eq=1|eq=3|eq=5"`
	DNSPrefix                string            `json:"dnsPrefix" validate:"required"`
	SubjectAltNames          []string          `json:"subjectAltNames"`
	VMSize                   string            `json:"vmSize" validate:"required"`
	OSDiskSizeGB             int               `json:"osDiskSizeGB,omitempty" validate:"min=0,max=1023"`
	VnetSubnetID             string            `json:"vnetSubnetID,omitempty"`
	VnetCidr                 string            `json:"vnetCidr,omitempty"`
	AgentVnetSubnetID        string            `json:"agentVnetSubnetID,omitempty"`
	FirstConsecutiveStaticIP string            `json:"firstConsecutiveStaticIP,omitempty"`
	IPAddressCount           int               `json:"ipAddressCount,omitempty" validate:"min=0,max=256"`
	StorageProfile           string            `json:"storageProfile,omitempty" validate:"eq=StorageAccount|eq=ManagedDisks|len=0"`
	HTTPSourceAddressPrefix  string            `json:"HTTPSourceAddressPrefix,omitempty"`
	OAuthEnabled             bool              `json:"oauthEnabled"`
	PreProvisionExtension    *Extension        `json:"preProvisionExtension"`
	Extensions               []Extension       `json:"extensions"`
	Distro                   Distro            `json:"distro,omitempty"`
	KubernetesConfig         *KubernetesConfig `json:"kubernetesConfig,omitempty"`
	ImageRef                 *ImageReference   `json:"imageReference,omitempty"`
	CustomFiles              *[]CustomFile     `json:"customFiles,omitempty"`
	AvailabilityProfile      string            `json:"availabilityProfile"`
	AgentSubnet              string            `json:"agentSubnet,omitempty"`
	AvailabilityZones        []string          `json:"availabilityZones,omitempty"`
	SinglePlacementGroup     *bool             `json:"singlePlacementGroup,omitempty"`

	// subnet is internal
	subnet string

	// Master LB public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`
}

// ImageReference represents a reference to an Image resource in Azure.
type ImageReference struct {
	Name          string `json:"name,omitempty"`
	ResourceGroup string `json:"resourceGroup,omitempty"`
}

// ExtensionProfile represents an extension definition
type ExtensionProfile struct {
	Name                           string             `json:"name"`
	Version                        string             `json:"version"`
	ExtensionParameters            string             `json:"extensionParameters,omitempty"`
	ExtensionParametersKeyVaultRef *KeyvaultSecretRef `json:"parametersKeyvaultSecretRef,omitempty"`
	RootURL                        string             `json:"rootURL,omitempty"`
	// This is only needed for preprovision extensions and it needs to be a bash script
	Script   string `json:"script,omitempty"`
	URLQuery string `json:"urlQuery,omitempty"`
}

// Extension represents an extension definition in the master or agentPoolProfile
type Extension struct {
	Name        string `json:"name"`
	SingleOrAll string `json:"singleOrAll"`
	Template    string `json:"template"`
}

// AgentPoolProfile represents an agent pool definition
type AgentPoolProfile struct {
	Name                                string               `json:"name" validate:"required"`
	Count                               int                  `json:"count" validate:"required,min=1,max=100"`
	VMSize                              string               `json:"vmSize" validate:"required"`
	OSDiskSizeGB                        int                  `json:"osDiskSizeGB,omitempty" validate:"min=0,max=1023"`
	DNSPrefix                           string               `json:"dnsPrefix,omitempty"`
	OSType                              OSType               `json:"osType,omitempty"`
	Ports                               []int                `json:"ports,omitempty" validate:"dive,min=1,max=65535"`
	AvailabilityProfile                 string               `json:"availabilityProfile"`
	ScaleSetPriority                    string               `json:"scaleSetPriority,omitempty" validate:"eq=Regular|eq=Low|len=0"`
	ScaleSetEvictionPolicy              string               `json:"scaleSetEvictionPolicy,omitempty" validate:"eq=Delete|eq=Deallocate|len=0"`
	StorageProfile                      string               `json:"storageProfile" validate:"eq=StorageAccount|eq=ManagedDisks|len=0"`
	DiskSizesGB                         []int                `json:"diskSizesGB,omitempty" validate:"max=4,dive,min=1,max=1023"`
	VnetSubnetID                        string               `json:"vnetSubnetID,omitempty"`
	IPAddressCount                      int                  `json:"ipAddressCount,omitempty" validate:"min=0,max=256"`
	Distro                              Distro               `json:"distro,omitempty"`
	KubernetesConfig                    *KubernetesConfig    `json:"kubernetesConfig,omitempty"`
	ImageRef                            *ImageReference      `json:"imageReference,omitempty"`
	Role                                AgentPoolProfileRole `json:"role,omitempty"`
	AcceleratedNetworkingEnabled        *bool                `json:"acceleratedNetworkingEnabled,omitempty"`
	AcceleratedNetworkingEnabledWindows *bool                `json:"acceleratedNetworkingEnabledWindows,omitempty"`

	// subnet is internal
	subnet string

	FQDN                  string            `json:"fqdn"`
	CustomNodeLabels      map[string]string `json:"customNodeLabels,omitempty"`
	PreProvisionExtension *Extension        `json:"preProvisionExtension"`
	Extensions            []Extension       `json:"extensions"`
	SinglePlacementGroup  *bool             `json:"singlePlacementGroup,omitempty"`
	AvailabilityZones     []string          `json:"availabilityZones,omitempty"`
}

// AgentPoolProfileRole represents an agent role
type AgentPoolProfileRole string

// AADProfile specifies attributes for AAD integration
type AADProfile struct {
	// The client AAD application ID.
	ClientAppID string `json:"clientAppID,omitempty"`
	// The server AAD application ID.
	ServerAppID string `json:"serverAppID,omitempty"`
	// The AAD tenant ID to use for authentication.
	// If not specified, will use the tenant of the deployment subscription.
	// Optional
	TenantID string `json:"tenantID,omitempty"`
	// The Azure Active Directory Group Object ID that will be assigned the
	// cluster-admin RBAC role.
	// Optional
	AdminGroupID string `json:"adminGroupID,omitempty"`
}

// KeyVaultSecrets specifies certificates to install on the pool
// of machines from a given key vault
// the key vault specified must have been granted read permissions to CRP
type KeyVaultSecrets struct {
	SourceVault       *KeyVaultID           `json:"sourceVault,omitempty"`
	VaultCertificates []KeyVaultCertificate `json:"vaultCertificates,omitempty"`
}

// KeyVaultID specifies a key vault
type KeyVaultID struct {
	ID string `json:"id,omitempty"`
}

// KeyVaultCertificate specifies a certificate to install
// On Linux, the certificate file is placed under the /var/lib/waagent directory
// with the file name <UppercaseThumbprint>.crt for the X509 certificate file
// and <UppercaseThumbprint>.prv for the private key. Both of these files are .pem formatted.
// On windows the certificate will be saved in the specified store.
type KeyVaultCertificate struct {
	CertificateURL   string `json:"certificateUrl,omitempty"`
	CertificateStore string `json:"certificateStore,omitempty"`
}

// OSType represents OS types of agents
type OSType string

// Distro represents Linux distro to use for Linux VMs
type Distro string

// HasWindows returns true if the cluster contains windows
func (p *Properties) HasWindows() bool {
	for _, agentPoolProfile := range p.AgentPoolProfiles {
		if agentPoolProfile.OSType == Windows {
			return true
		}
	}
	return false
}

// HasAvailabilityZones returns true if the cluster contains any profile with zones
func (p *Properties) HasAvailabilityZones() bool {
	hasZones := p.MasterProfile != nil && p.MasterProfile.HasAvailabilityZones()
	if !hasZones && p.AgentPoolProfiles != nil {
		for _, agentPoolProfile := range p.AgentPoolProfiles {
			if agentPoolProfile.HasAvailabilityZones() {
				hasZones = true
				break
			}
		}
	}
	return hasZones
}

// IsCustomVNET returns true if the customer brought their own VNET
func (m *MasterProfile) IsCustomVNET() bool {
	return len(m.VnetSubnetID) > 0
}

// GetSubnet returns the read-only subnet for the master
func (m *MasterProfile) GetSubnet() string {
	return m.subnet
}

// SetSubnet sets the read-only subnet for the master
func (m *MasterProfile) SetSubnet(subnet string) {
	m.subnet = subnet
}

// IsManagedDisks returns true if the master specified managed disks
func (m *MasterProfile) IsManagedDisks() bool {
	return m.StorageProfile == ManagedDisks
}

// IsStorageAccount returns true if the master specified storage account
func (m *MasterProfile) IsStorageAccount() bool {
	return m.StorageProfile == StorageAccount
}

// IsRHEL returns true if the master specified a RHEL distro
func (m *MasterProfile) IsRHEL() bool {
	return m.Distro == RHEL
}

// IsCoreOS returns true if the master specified a CoreOS distro
func (m *MasterProfile) IsCoreOS() bool {
	return m.Distro == CoreOS
}

// IsVirtualMachineScaleSets returns true if the master availability profile is VMSS
func (m *MasterProfile) IsVirtualMachineScaleSets() bool {
	return m.AvailabilityProfile == VirtualMachineScaleSets
}

// HasAvailabilityZones returns true if the master profile has availability zones
func (m *MasterProfile) HasAvailabilityZones() bool {
	return m.AvailabilityZones != nil && len(m.AvailabilityZones) > 0
}

// HasZonesForAllAgentPools returns true if all of the agent pools have zones
func (p *Properties) HasZonesForAllAgentPools() bool {
	for _, ap := range p.AgentPoolProfiles {
		if !ap.HasAvailabilityZones() {
			return false
		}
	}
	return true
}

// MastersAndAgentsUseAvailabilityZones returns true if the cluster contains AZs for all agents and masters profiles
func (p *Properties) MastersAndAgentsUseAvailabilityZones() bool {
	return (p.MasterProfile != nil && p.MasterProfile.HasAvailabilityZones()) && p.HasZonesForAllAgentPools()
}

// IsClusterAllVirtualMachineScaleSets returns true if the cluster contains only Virtual Machine Scale Sets
func (p *Properties) IsClusterAllVirtualMachineScaleSets() bool {
	isAll := p.MasterProfile != nil && p.MasterProfile.IsVirtualMachineScaleSets()
	if isAll && p.AgentPoolProfiles != nil {
		for _, agentPoolProfile := range p.AgentPoolProfiles {
			if agentPoolProfile.AvailabilityProfile == AvailabilitySet {
				isAll = false
				break
			}
		}
	}
	return isAll
}

// IsCustomVNET returns true if the customer brought their own VNET
func (a *AgentPoolProfile) IsCustomVNET() bool {
	return len(a.VnetSubnetID) > 0
}

// IsWindows returns true if the agent pool is windows
func (a *AgentPoolProfile) IsWindows() bool {
	return a.OSType == Windows
}

// IsLinux returns true if the agent pool is linux
func (a *AgentPoolProfile) IsLinux() bool {
	return a.OSType == Linux
}

// IsRHEL returns true if the agent pool specified a RHEL distro
func (a *AgentPoolProfile) IsRHEL() bool {
	return a.OSType == Linux && a.Distro == RHEL
}

// IsCoreOS returns true if the agent specified a CoreOS distro
func (a *AgentPoolProfile) IsCoreOS() bool {
	return a.OSType == Linux && a.Distro == CoreOS
}

// IsAvailabilitySets returns true if the customer specified disks
func (a *AgentPoolProfile) IsAvailabilitySets() bool {
	return a.AvailabilityProfile == AvailabilitySet
}

// IsVirtualMachineScaleSets returns true if the agent pool availability profile is VMSS
func (a *AgentPoolProfile) IsVirtualMachineScaleSets() bool {
	return a.AvailabilityProfile == VirtualMachineScaleSets
}

// IsNSeriesSKU returns true if the agent pool contains an N-series (NVIDIA GPU) VM
func (a *AgentPoolProfile) IsNSeriesSKU() bool {
	return common.IsNvidiaEnabledSKU(a.VMSize)
}

// IsManagedDisks returns true if the customer specified managed disks
func (a *AgentPoolProfile) IsManagedDisks() bool {
	return a.StorageProfile == ManagedDisks
}

// IsStorageAccount returns true if the customer specified storage account
func (a *AgentPoolProfile) IsStorageAccount() bool {
	return a.StorageProfile == StorageAccount
}

// HasDisks returns true if the customer specified disks
func (a *AgentPoolProfile) HasDisks() bool {
	return len(a.DiskSizesGB) > 0
}

// GetSubnet returns the read-only subnet for the agent pool
func (a *AgentPoolProfile) GetSubnet() string {
	return a.subnet
}

// SetSubnet sets the read-only subnet for the agent pool
func (a *AgentPoolProfile) SetSubnet(subnet string) {
	a.subnet = subnet
}

// HasAvailabilityZones returns true if the agent pool has availability zones
func (a *AgentPoolProfile) HasAvailabilityZones() bool {
	return a.AvailabilityZones != nil && len(a.AvailabilityZones) > 0
}

// HasSearchDomain returns true if the customer specified secrets to install
func (l *LinuxProfile) HasSearchDomain() bool {
	if l.CustomSearchDomain != nil {
		if l.CustomSearchDomain.Name != "" && l.CustomSearchDomain.RealmPassword != "" && l.CustomSearchDomain.RealmUser != "" {
			return true
		}
	}
	return false
}

// HasCustomNodesDNS returns true if the customer specified secrets to install
func (l *LinuxProfile) HasCustomNodesDNS() bool {
	if l.CustomNodesDNS != nil {
		if l.CustomNodesDNS.DNSServer != "" {
			return true
		}
	}
	return false
}

// IsSwarmMode returns true if this template is for Swarm Mode orchestrator
func (o *OrchestratorProfile) IsSwarmMode() bool {
	return o.OrchestratorType == SwarmMode
}

// RequiresDocker returns if the kubernetes settings require docker binary to be installed.
func (k *KubernetesConfig) RequiresDocker() bool {
	runtime := strings.ToLower(k.ContainerRuntime)
	return runtime == "docker" || runtime == ""
}
