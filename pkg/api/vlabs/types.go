// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"encoding/json"
	"strings"

	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"
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
	CustomCloudProfile      *CustomCloudProfile      `json:"customCloudProfile,omitempty"`
	TelemetryProfile        *TelemetryProfile        `json:"telemetryProfile,omitempty"`
}

// FeatureFlags defines feature-flag restricted functionality
type FeatureFlags struct {
	EnableCSERunInBackground bool `json:"enableCSERunInBackground,omitempty"`
	BlockOutboundInternet    bool `json:"blockOutboundInternet,omitempty"`
	EnableIPv6DualStack      bool `json:"enableIPv6DualStack,omitempty"`
	EnableTelemetry          bool `json:"enableTelemetry,omitempty"`
	EnableIPv6Only           bool `json:"enableIPv6Only,omitempty"`
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
		PublicKeys []PublicKey `json:"publicKeys" validate:"required,min=1"`
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
	AdminUsername          string            `json:"adminUsername,omitempty"`
	AdminPassword          string            `json:"adminPassword,omitempty"`
	CSIProxyURL            string            `json:"csiProxyURL,omitempty"`
	EnableCSIProxy         *bool             `json:"enableCSIProxy,omitempty"`
	ImageRef               *ImageReference   `json:"imageReference,omitempty"`
	ImageVersion           string            `json:"imageVersion,omitempty"`
	WindowsImageSourceURL  string            `json:"WindowsImageSourceUrl"`
	WindowsPublisher       string            `json:"WindowsPublisher"`
	WindowsOffer           string            `json:"WindowsOffer"`
	WindowsSku             string            `json:"WindowsSku"`
	WindowsDockerVersion   string            `json:"windowsDockerVersion"`
	Secrets                []KeyVaultSecrets `json:"secrets,omitempty"`
	SSHEnabled             *bool             `json:"sshEnabled,omitempty"`
	EnableAutomaticUpdates *bool             `json:"enableAutomaticUpdates,omitempty"`
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

// AddonNodePoolsConfig defines configuration for pool-specific cluster-autoscaler configuration
type AddonNodePoolsConfig struct {
	Name   string            `json:"name,omitempty"`
	Config map[string]string `json:"config,omitempty"`
}

// KubernetesAddon defines a list of addons w/ configuration to include with the cluster deployment
type KubernetesAddon struct {
	Name       string                    `json:"name,omitempty"`
	Enabled    *bool                     `json:"enabled,omitempty"`
	Mode       string                    `json:"mode,omitempty"`
	Containers []KubernetesContainerSpec `json:"containers,omitempty"`
	Config     map[string]string         `json:"config,omitempty"`
	Pools      []AddonNodePoolsConfig    `json:"pools,omitempty"`
	Data       string                    `json:"data,omitempty"`
}

// IsEnabled returns true if the addon is enabled
func (a *KubernetesAddon) IsEnabled() bool {
	if a.Enabled == nil {
		return false
	}
	return *a.Enabled
}

// KubernetesComponent defines a component w/ configuration to include with the cluster deployment
type KubernetesComponent struct {
	Name       string                    `json:"name,omitempty"`
	Enabled    *bool                     `json:"enabled,omitempty"`
	Containers []KubernetesContainerSpec `json:"containers,omitempty"`
	Config     map[string]string         `json:"config,omitempty"`
	Data       string                    `json:"data,omitempty"`
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
	OSDiskSizeGB   int    `json:"osDiskSizeGB,omitempty" validate:"min=0,max=2048"`
	Username       string `json:"username,omitempty"`
	PublicKey      string `json:"publicKey" validate:"required"`
	StorageProfile string `json:"storageProfile,omitempty"`
}

// KubeProxyMode is for iptables and ipvs (and future others)
type KubeProxyMode string

// We currently support ipvs and iptables
const (
	KubeProxyModeIPTables KubeProxyMode = "iptables"
	KubeProxyModeIPVS     KubeProxyMode = "ipvs"
)

// KubernetesConfig contains the Kubernetes config structure, containing
// Kubernetes specific configuration
type KubernetesConfig struct {
	KubernetesImageBase               string                `json:"kubernetesImageBase,omitempty"`
	KubernetesImageBaseType           string                `json:"kubernetesImageBaseType,omitempty"`
	MCRKubernetesImageBase            string                `json:"mcrKubernetesImageBase,omitempty"`
	ClusterSubnet                     string                `json:"clusterSubnet,omitempty"`
	DNSServiceIP                      string                `json:"dnsServiceIP,omitempty"`
	ServiceCidr                       string                `json:"serviceCidr,omitempty"`
	NetworkPolicy                     string                `json:"networkPolicy,omitempty"`
	NetworkPlugin                     string                `json:"networkPlugin,omitempty"`
	NetworkMode                       string                `json:"networkMode,omitempty"`
	ContainerRuntime                  string                `json:"containerRuntime,omitempty"`
	MaxPods                           int                   `json:"maxPods,omitempty"`
	DockerBridgeSubnet                string                `json:"dockerBridgeSubnet,omitempty"`
	UseManagedIdentity                bool                  `json:"useManagedIdentity,omitempty"`
	UserAssignedID                    string                `json:"userAssignedID,omitempty"`
	UserAssignedClientID              string                `json:"userAssignedClientID,omitempty"` //Note: cannot be provided in config. Used *only* for transferring this to azure.json.
	CustomHyperkubeImage              string                `json:"customHyperkubeImage,omitempty"`
	CustomKubeAPIServerImage          string                `json:"customKubeAPIServerImage,omitempty"`
	CustomKubeControllerManagerImage  string                `json:"customKubeControllerManagerImage,omitempty"`
	CustomKubeProxyImage              string                `json:"customKubeProxyImage,omitempty"`
	CustomKubeSchedulerImage          string                `json:"customKubeSchedulerImage,omitempty"`
	CustomKubeBinaryURL               string                `json:"customKubeBinaryURL,omitempty"`
	DockerEngineVersion               string                `json:"dockerEngineVersion,omitempty"` // Deprecated
	MobyVersion                       string                `json:"mobyVersion,omitempty"`
	ContainerdVersion                 string                `json:"containerdVersion,omitempty"`
	CustomCcmImage                    string                `json:"customCcmImage,omitempty"`
	UseCloudControllerManager         *bool                 `json:"useCloudControllerManager,omitempty"`
	CustomWindowsPackageURL           string                `json:"customWindowsPackageURL,omitempty"`
	WindowsNodeBinariesURL            string                `json:"windowsNodeBinariesURL,omitempty"`
	WindowsContainerdURL              string                `json:"windowsContainerdURL,omitempty"`
	WindowsSdnPluginURL               string                `json:"windowsSdnPluginURL,omitempty"`
	UseInstanceMetadata               *bool                 `json:"useInstanceMetadata,omitempty"`
	EnableRbac                        *bool                 `json:"enableRbac,omitempty"`
	EnableSecureKubelet               *bool                 `json:"enableSecureKubelet,omitempty"`
	EnableAggregatedAPIs              bool                  `json:"enableAggregatedAPIs,omitempty"`
	PrivateCluster                    *PrivateCluster       `json:"privateCluster,omitempty"`
	GCHighThreshold                   int                   `json:"gchighthreshold,omitempty"`
	GCLowThreshold                    int                   `json:"gclowthreshold,omitempty"`
	EtcdVersion                       string                `json:"etcdVersion,omitempty"`
	EtcdDiskSizeGB                    string                `json:"etcdDiskSizeGB,omitempty"`
	EtcdEncryptionKey                 string                `json:"etcdEncryptionKey,omitempty"`
	EnableDataEncryptionAtRest        *bool                 `json:"enableDataEncryptionAtRest,omitempty"`
	EnableEncryptionWithExternalKms   *bool                 `json:"enableEncryptionWithExternalKms,omitempty"`
	EnablePodSecurityPolicy           *bool                 `json:"enablePodSecurityPolicy,omitempty"`
	Addons                            []KubernetesAddon     `json:"addons,omitempty"`
	Components                        []KubernetesComponent `json:"components,omitempty"`
	ContainerRuntimeConfig            map[string]string     `json:"containerRuntimeConfig,omitempty"`
	KubeletConfig                     map[string]string     `json:"kubeletConfig,omitempty"`
	ControllerManagerConfig           map[string]string     `json:"controllerManagerConfig,omitempty"`
	CloudControllerManagerConfig      map[string]string     `json:"cloudControllerManagerConfig,omitempty"`
	APIServerConfig                   map[string]string     `json:"apiServerConfig,omitempty"`
	SchedulerConfig                   map[string]string     `json:"schedulerConfig,omitempty"`
	PodSecurityPolicyConfig           map[string]string     `json:"podSecurityPolicyConfig,omitempty"` // Deprecated
	KubeReservedCgroup                string                `json:"kubeReservedCgroup,omitempty"`
	CloudProviderBackoffMode          string                `json:"cloudProviderBackoffMode"`
	CloudProviderBackoff              *bool                 `json:"cloudProviderBackoff,omitempty"`
	CloudProviderBackoffRetries       int                   `json:"cloudProviderBackoffRetries,omitempty"`
	CloudProviderBackoffJitter        float64               `json:"cloudProviderBackoffJitter,omitempty"`
	CloudProviderBackoffDuration      int                   `json:"cloudProviderBackoffDuration,omitempty"`
	CloudProviderBackoffExponent      float64               `json:"cloudProviderBackoffExponent,omitempty"`
	CloudProviderRateLimit            *bool                 `json:"cloudProviderRateLimit,omitempty"`
	CloudProviderRateLimitQPS         float64               `json:"cloudProviderRateLimitQPS,omitempty"`
	CloudProviderRateLimitQPSWrite    float64               `json:"cloudProviderRateLimitQPSWrite,omitempty"`
	CloudProviderRateLimitBucket      int                   `json:"cloudProviderRateLimitBucket,omitempty"`
	CloudProviderRateLimitBucketWrite int                   `json:"cloudProviderRateLimitBucketWrite,omitempty"`
	CloudProviderDisableOutboundSNAT  *bool                 `json:"cloudProviderDisableOutboundSNAT,omitempty"`
	LoadBalancerSku                   string                `json:"loadBalancerSku,omitempty"`
	ExcludeMasterFromStandardLB       *bool                 `json:"excludeMasterFromStandardLB,omitempty"`
	LoadBalancerOutboundIPs           *int                  `json:"loadBalancerOutboundIPs,omitempty"`
	AzureCNIVersion                   string                `json:"azureCNIVersion,omitempty"`
	AzureCNIURLLinux                  string                `json:"azureCNIURLLinux,omitempty"`
	AzureCNIURLWindows                string                `json:"azureCNIURLWindows,omitempty"`
	KeyVaultSku                       string                `json:"keyVaultSku,omitempty"`
	MaximumLoadBalancerRuleCount      int                   `json:"maximumLoadBalancerRuleCount,omitempty"`
	ProxyMode                         KubeProxyMode         `json:"kubeProxyMode,omitempty"`
	PrivateAzureRegistryServer        string                `json:"privateAzureRegistryServer,omitempty"`
	OutboundRuleIdleTimeoutInMinutes  int32                 `json:"outboundRuleIdleTimeoutInMinutes,omitempty"`
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
	Count                     int               `json:"count" validate:"required,eq=1|eq=3|eq=5"`
	DNSPrefix                 string            `json:"dnsPrefix" validate:"required"`
	SubjectAltNames           []string          `json:"subjectAltNames"`
	VMSize                    string            `json:"vmSize" validate:"required"`
	OSDiskSizeGB              int               `json:"osDiskSizeGB,omitempty" validate:"min=0,max=2048"`
	VnetSubnetID              string            `json:"vnetSubnetID,omitempty"`
	VnetCidr                  string            `json:"vnetCidr,omitempty"`
	AgentVnetSubnetID         string            `json:"agentVnetSubnetID,omitempty"`
	FirstConsecutiveStaticIP  string            `json:"firstConsecutiveStaticIP,omitempty"`
	IPAddressCount            int               `json:"ipAddressCount,omitempty" validate:"min=0,max=256"`
	StorageProfile            string            `json:"storageProfile,omitempty" validate:"eq=StorageAccount|eq=ManagedDisks|len=0"`
	HTTPSourceAddressPrefix   string            `json:"HTTPSourceAddressPrefix,omitempty"`
	OAuthEnabled              bool              `json:"oauthEnabled"`
	PreProvisionExtension     *Extension        `json:"preProvisionExtension"`
	Extensions                []Extension       `json:"extensions"`
	Distro                    Distro            `json:"distro,omitempty"`
	KubernetesConfig          *KubernetesConfig `json:"kubernetesConfig,omitempty"`
	ImageRef                  *ImageReference   `json:"imageReference,omitempty"`
	CustomFiles               *[]CustomFile     `json:"customFiles,omitempty"`
	AvailabilityProfile       string            `json:"availabilityProfile"`
	AgentSubnet               string            `json:"agentSubnet,omitempty"`
	AvailabilityZones         []string          `json:"availabilityZones,omitempty"`
	SinglePlacementGroup      *bool             `json:"singlePlacementGroup,omitempty"`
	PlatformFaultDomainCount  *int              `json:"platformFaultDomainCount,omitempty"`
	PlatformUpdateDomainCount *int              `json:"platformUpdateDomainCount"`
	AuditDEnabled             *bool             `json:"auditDEnabled,omitempty"`
	CustomVMTags              map[string]string `json:"customVMTags,omitempty"`
	SysctlDConfig             map[string]string `json:"sysctldConfig,omitempty"`
	UltraSSDEnabled           *bool             `json:"ultraSSDEnabled,omitempty"`
	EncryptionAtHost          *bool             `json:"encryptionAtHost,omitempty"`

	// subnet is internal
	subnet string
	// subnetIPv6 is internal
	subnetIPv6 string

	// Master LB public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`

	// True: uses cosmos etcd endpoint instead of installing etcd on masters
	CosmosEtcd                *bool  `json:"cosmosEtcd,omitempty"`
	ProximityPlacementGroupID string `json:"proximityPlacementGroupID,omitempty"`
	OSDiskCachingType         string `json:"osDiskCachingType,omitempty"`
}

// ImageReference represents a reference to an Image resource in Azure.
type ImageReference struct {
	Name           string `json:"name,omitempty"`
	ResourceGroup  string `json:"resourceGroup,omitempty"`
	SubscriptionID string `json:"subscriptionId,omitempty"`
	Gallery        string `json:"gallery,omitempty"`
	Version        string `json:"version,omitempty"`
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
	OSDiskSizeGB                        int                  `json:"osDiskSizeGB,omitempty" validate:"min=0,max=2048"`
	DNSPrefix                           string               `json:"dnsPrefix,omitempty"`
	OSType                              OSType               `json:"osType,omitempty"`
	Ports                               []int                `json:"ports,omitempty" validate:"dive,min=1,max=65535"`
	AvailabilityProfile                 string               `json:"availabilityProfile"`
	ScaleSetPriority                    string               `json:"scaleSetPriority,omitempty" validate:"eq=Regular|eq=Low|eq=Spot|len=0"`
	ScaleSetEvictionPolicy              string               `json:"scaleSetEvictionPolicy,omitempty" validate:"eq=Delete|eq=Deallocate|len=0"`
	SpotMaxPrice                        *float64             `json:"spotMaxPrice,omitempty"`
	StorageProfile                      string               `json:"storageProfile" validate:"eq=StorageAccount|eq=ManagedDisks|eq=Ephemeral|len=0"`
	DiskSizesGB                         []int                `json:"diskSizesGB,omitempty" validate:"max=4,dive,min=1,max=32767"`
	VnetSubnetID                        string               `json:"vnetSubnetID,omitempty"`
	IPAddressCount                      int                  `json:"ipAddressCount,omitempty" validate:"min=0,max=256"`
	Distro                              Distro               `json:"distro,omitempty"`
	KubernetesConfig                    *KubernetesConfig    `json:"kubernetesConfig,omitempty"`
	ImageRef                            *ImageReference      `json:"imageReference,omitempty"`
	Role                                AgentPoolProfileRole `json:"role,omitempty"`
	AcceleratedNetworkingEnabled        *bool                `json:"acceleratedNetworkingEnabled,omitempty"`
	AcceleratedNetworkingEnabledWindows *bool                `json:"acceleratedNetworkingEnabledWindows,omitempty"`
	VMSSOverProvisioningEnabled         *bool                `json:"vmssOverProvisioningEnabled,omitempty"`
	AuditDEnabled                       *bool                `json:"auditDEnabled,omitempty"`
	CustomVMTags                        map[string]string    `json:"customVMTags,omitempty"`
	DiskEncryptionSetID                 string               `json:"diskEncryptionSetID,omitempty"`
	UltraSSDEnabled                     *bool                `json:"ultraSSDEnabled,omitempty"`
	EncryptionAtHost                    *bool                `json:"encryptionAtHost,omitempty"`
	// subnet is internal
	subnet string

	FQDN                              string            `json:"fqdn"`
	CustomNodeLabels                  map[string]string `json:"customNodeLabels,omitempty"`
	PreProvisionExtension             *Extension        `json:"preProvisionExtension"`
	Extensions                        []Extension       `json:"extensions"`
	SinglePlacementGroup              *bool             `json:"singlePlacementGroup,omitempty"`
	PlatformFaultDomainCount          *int              `json:"platformFaultDomainCount,omitempty"`
	PlatformUpdateDomainCount         *int              `json:"platformUpdateDomainCount"`
	AvailabilityZones                 []string          `json:"availabilityZones,omitempty"`
	EnableVMSSNodePublicIP            *bool             `json:"enableVMSSNodePublicIP,omitempty"`
	LoadBalancerBackendAddressPoolIDs []string          `json:"loadBalancerBackendAddressPoolIDs,omitempty"`
	SysctlDConfig                     map[string]string `json:"sysctldConfig,omitempty"`
	ProximityPlacementGroupID         string            `json:"proximityPlacementGroupID,omitempty"`
	OSDiskCachingType                 string            `json:"osDiskCachingType,omitempty"`
	DataDiskCachingType               string            `json:"dataDiskCachingType,omitempty"`
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

// DependenciesLocation represents location to retrieve the dependencies.
type DependenciesLocation string

// CustomCloudProfile represents the custom cloud profile
type CustomCloudProfile struct {
	Environment                 *azure.Environment          `json:"environment,omitempty"`
	AzureEnvironmentSpecConfig  *AzureEnvironmentSpecConfig `json:"azureEnvironmentSpecConfig,omitempty"`
	IdentitySystem              string                      `json:"identitySystem,omitempty"`
	AuthenticationMethod        string                      `json:"authenticationMethod,omitempty"`
	DependenciesLocation        DependenciesLocation        `json:"dependenciesLocation,omitempty"`
	PortalURL                   string                      `json:"portalURL,omitempty"`
	CustomCloudRootCertificates string                      `json:"customCloudRootCertificates,omitempty"`
	CustomCloudSourcesList      string                      `json:"customCloudSourcesList,omitempty"`
}

// TelemetryProfile contains settings for collecting telemtry.
// Note telemtry is currently enabled/disabled with the 'EnableTelemetry' feature flag.
type TelemetryProfile struct {
	ApplicationInsightsKey string `json:"applicationInsightsKey,omitempty"`
}

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

// IsCustomCloudProfile return true if user has provided a custom cloud profile
func (p *Properties) IsCustomCloudProfile() bool {
	return p.CustomCloudProfile != nil
}

// GetCustomCloudRootCertificates returns comma-separated list of base64-encoded custom root certificates
func (p *Properties) GetCustomCloudRootCertificates() string {
	if p.IsCustomCloudProfile() {
		return p.CustomCloudProfile.CustomCloudRootCertificates
	}
	return ""
}

// GetCustomCloudSourcesList returns a base64-encoded custom sources.list file
func (p *Properties) GetCustomCloudSourcesList() string {
	if p.IsCustomCloudProfile() {
		return p.CustomCloudProfile.CustomCloudSourcesList
	}
	return ""
}

// IsAzureStackCloud return true if the cloud is AzureStack
func (p *Properties) IsAzureStackCloud() bool {
	// For backward compatibility, treat nil Environment and empty Environment name as AzureStackCloud as well
	return p.IsCustomCloudProfile() && (p.CustomCloudProfile.Environment == nil || p.CustomCloudProfile.Environment.Name == "" || strings.EqualFold(p.CustomCloudProfile.Environment.Name, "AzureStackCloud"))
}

// HasAADAdminGroupID returns true if the cluster has an AADProfile w/ a valid AdminGroupID
func (p *Properties) HasAADAdminGroupID() bool {
	return p.AADProfile != nil && p.AADProfile.AdminGroupID != ""
}

// GetAADAdminGroupID returns AADProfile.AdminGroupID, or "" if no AADProfile
func (p *Properties) GetAADAdminGroupID() string {
	if p.AADProfile != nil {
		return p.AADProfile.AdminGroupID
	}
	return ""
}

// IsCustomVNET returns true if the customer brought their own VNET
func (m *MasterProfile) IsCustomVNET() bool {
	return len(m.VnetSubnetID) > 0
}

// GetSubnet returns the read-only subnet for the master
func (m *MasterProfile) GetSubnet() string {
	return m.subnet
}

// GetSubnetIPv6 returns the read-only ipv6 subnet for the master
func (m *MasterProfile) GetSubnetIPv6() string {
	return m.subnetIPv6
}

// SetSubnet sets the read-only subnet for the master
func (m *MasterProfile) SetSubnet(subnet string) {
	m.subnet = subnet
}

// SetSubnetIPv6 sets the read-only ipv6 subnet for the master
func (m *MasterProfile) SetSubnetIPv6(subnetIPv6 string) {
	m.subnetIPv6 = subnetIPv6
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

// IsUbuntu1604 returns true if the master profile distro is based on Ubuntu 16.04
func (m *MasterProfile) IsUbuntu1604() bool {
	switch m.Distro {
	case AKSUbuntu1604, Ubuntu, ACC1604:
		return true
	default:
		return false
	}
}

// IsUbuntu1804 returns true if the master profile distro is based on Ubuntu 18.04
func (m *MasterProfile) IsUbuntu1804() bool {
	switch m.Distro {
	case AKSUbuntu1804, Ubuntu1804, Ubuntu1804Gen2:
		return true
	default:
		return false
	}
}

// IsUbuntu returns true if the master profile distro is any ubuntu distro
func (m *MasterProfile) IsUbuntu() bool {
	return m.IsUbuntu1604() || m.IsUbuntu1804()
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

// GetAgentPoolByName returns the pool in the AgentPoolProfiles array that matches a name, nil if no match
func (p *Properties) GetAgentPoolByName(name string) *AgentPoolProfile {
	for _, profile := range p.AgentPoolProfiles {
		if profile.Name == name {
			return profile
		}
	}
	return nil
}

// ShouldEnableAzureCloudAddon determines whether or not we should enable the following addons:
// 1. cloud-node-manager,
// 2. azuredisk-csi-driver,
// 3. azurefile-csi-driver.
// For Linux clusters, we should enable CSI Drivers when using K8s 1.13+ and cloud-node-manager when using K8s 1.16+.
// For Windows clusters, we should enable them when using K8s 1.18+.
func (p *Properties) ShouldEnableAzureCloudAddon(addonName string) bool {
	o := p.OrchestratorProfile
	if !to.Bool(o.KubernetesConfig.UseCloudControllerManager) {
		return false
	}
	if !p.HasWindows() {
		switch addonName {
		case common.AzureDiskCSIDriverAddonName, common.AzureFileCSIDriverAddonName:
			return common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.13.0")
		case common.CloudNodeManagerAddonName:
			return common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.16.0")
		default:
			return false
		}
	}
	return common.IsKubernetesVersionGe(o.OrchestratorVersion, "1.18.0")
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

// IsEphemeral returns true if the customer specified ephemeral disks
func (a *AgentPoolProfile) IsEphemeral() bool {
	return a.StorageProfile == Ephemeral
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

// IsUbuntu1604 returns true if the agent pool profile distro is based on Ubuntu 16.04
func (a *AgentPoolProfile) IsUbuntu1604() bool {
	if a.OSType != Windows {
		switch a.Distro {
		case AKSUbuntu1604, Ubuntu, ACC1604:
			return true
		default:
			return false
		}
	}
	return false
}

// IsUbuntu1804 returns true if the agent pool profile distro is based on Ubuntu 16.04
func (a *AgentPoolProfile) IsUbuntu1804() bool {
	if a.OSType != Windows {
		switch a.Distro {
		case AKSUbuntu1804, Ubuntu1804, Ubuntu1804Gen2:
			return true
		default:
			return false
		}
	}
	return false
}

// IsUbuntu returns true if the master profile distro is any ubuntu distro
func (a *AgentPoolProfile) IsUbuntu() bool {
	return a.IsUbuntu1604() || a.IsUbuntu1804()
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

// IsCSIProxyEnabled returns true if CSI proxy service should be enable for Windows nodes
func (w *WindowsProfile) IsCSIProxyEnabled() bool {
	if w.EnableCSIProxy != nil {
		return *w.EnableCSIProxy
	}
	return common.DefaultEnableCSIProxyWindows
}

// IsSwarmMode returns true if this template is for Swarm Mode orchestrator
func (o *OrchestratorProfile) IsSwarmMode() bool {
	return o.OrchestratorType == SwarmMode
}

// RequiresDocker returns if the kubernetes settings require docker binary to be installed.
func (k *KubernetesConfig) RequiresDocker() bool {
	runtime := strings.ToLower(k.ContainerRuntime)
	return runtime == Docker || runtime == ""
}

// IsRBACEnabled checks if RBAC is enabled
func (k *KubernetesConfig) IsRBACEnabled() bool {
	if k.EnableRbac != nil {
		return to.Bool(k.EnableRbac)
	}
	return false
}

// GetAddonByName returns the KubernetesAddon instance with name `addonName`
func (k *KubernetesConfig) GetAddonByName(addonName string) KubernetesAddon {
	var kubeAddon KubernetesAddon
	for _, addon := range k.Addons {
		if addon.Name == addonName {
			kubeAddon = addon
			break
		}
	}
	return kubeAddon
}

// IsAddonEnabled checks whether a k8s addon with name "addonName" is enabled or not based on the Enabled field of KubernetesAddon.
// If the value of Enabled is nil, the "defaultValue" is returned.
func (k *KubernetesConfig) IsAddonEnabled(addonName string) bool {
	if k != nil {
		kubeAddon := k.GetAddonByName(addonName)
		return kubeAddon.IsEnabled()
	}
	return false
}

// IsIPv6DualStackEnabled checks if IPv6DualStack feature is enabled
func (f *FeatureFlags) IsIPv6DualStackEnabled() bool {
	return f != nil && f.EnableIPv6DualStack
}

// IsIPv6OnlyEnabled checks if IPv6Only feature is enabled
func (f *FeatureFlags) IsIPv6OnlyEnabled() bool {
	return f != nil && f.EnableIPv6Only
}
