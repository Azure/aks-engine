// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math/rand"
	"net"
	neturl "net/url"
	"sort"
	"strconv"
	"strings"

	v20170831 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20170831"
	v20180331 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20180331"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/blang/semver"
)

// TypeMeta describes an individual API model object
type TypeMeta struct {
	// APIVersion is on every object
	APIVersion string `json:"apiVersion"`
}

// ResourcePurchasePlan defines resource plan as required by ARM
// for billing purposes.
type ResourcePurchasePlan struct {
	Name          string `json:"name"`
	Product       string `json:"product"`
	PromotionCode string `json:"promotionCode"`
	Publisher     string `json:"publisher"`
}

// ContainerService complies with the ARM model of
// resource definition in a JSON template.
type ContainerService struct {
	ID       string                `json:"id"`
	Location string                `json:"location"`
	Name     string                `json:"name"`
	Plan     *ResourcePurchasePlan `json:"plan,omitempty"`
	Tags     map[string]string     `json:"tags"`
	Type     string                `json:"type"`

	Properties *Properties `json:"properties,omitempty"`
}

// AgentPoolResource complies with the ARM model of
// agentpool resource definition in a JSON template.
type AgentPoolResource struct {
	ID       string                `json:"id"`
	Location string                `json:"location"`
	Name     string                `json:"name"`
	Plan     *ResourcePurchasePlan `json:"plan,omitempty"`
	Tags     map[string]string     `json:"tags"`
	Type     string                `json:"type"`

	Properties *AgentPoolProfile `json:"properties,omitempty"`
}

// Properties represents the AKS cluster definition
type Properties struct {
	ClusterID               string
	ProvisioningState       ProvisioningState        `json:"provisioningState,omitempty"`
	OrchestratorProfile     *OrchestratorProfile     `json:"orchestratorProfile,omitempty"`
	MasterProfile           *MasterProfile           `json:"masterProfile,omitempty"`
	AgentPoolProfiles       []*AgentPoolProfile      `json:"agentPoolProfiles,omitempty"`
	LinuxProfile            *LinuxProfile            `json:"linuxProfile,omitempty"`
	WindowsProfile          *WindowsProfile          `json:"windowsProfile,omitempty"`
	ExtensionProfiles       []*ExtensionProfile      `json:"extensionProfiles"`
	DiagnosticsProfile      *DiagnosticsProfile      `json:"diagnosticsProfile,omitempty"`
	JumpboxProfile          *JumpboxProfile          `json:"jumpboxProfile,omitempty"`
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
	CertificateProfile      *CertificateProfile      `json:"certificateProfile,omitempty"`
	AADProfile              *AADProfile              `json:"aadProfile,omitempty"`
	CustomProfile           *CustomProfile           `json:"customProfile,omitempty"`
	HostedMasterProfile     *HostedMasterProfile     `json:"hostedMasterProfile,omitempty"`
	AddonProfiles           map[string]AddonProfile  `json:"addonProfiles,omitempty"`
	FeatureFlags            *FeatureFlags            `json:"featureFlags,omitempty"`
	CustomCloudProfile      *CustomCloudProfile      `json:"customCloudProfile,omitempty"`
	TelemetryProfile        *TelemetryProfile        `json:"telemetryProfile,omitempty"`
}

// ClusterMetadata represents the metadata of the AKS cluster.
type ClusterMetadata struct {
	SubnetName                 string `json:"subnetName,omitempty"`
	VNetResourceGroupName      string `json:"vnetResourceGroupName,omitempty"`
	VirtualNetworkName         string `json:"virtualNetworkName,omitempty"`
	SecurityGroupName          string `json:"securityGroupName,omitempty"`
	RouteTableName             string `json:"routeTableName,omitempty"`
	PrimaryAvailabilitySetName string `json:"primaryAvailabilitySetName,omitempty"`
	PrimaryScaleSetName        string `json:"primaryScaleSetName,omitempty"`
	ResourcePrefix             string `json:"resourcePrefix,omitempty"`
}

// AddonProfile represents an addon for managed cluster
type AddonProfile struct {
	Enabled bool              `json:"enabled"`
	Config  map[string]string `json:"config"`
	// Identity contains information of the identity associated with this addon.
	// This property will only appear in an MSI-enabled cluster.
	Identity *UserAssignedIdentity `json:"identity,omitempty"`
}

// UserAssignedIdentity contains information that uniquely identifies an identity
type UserAssignedIdentity struct {
	ResourceID string `json:"resourceId,omitempty"`
	ClientID   string `json:"clientId,omitempty"`
	ObjectID   string `json:"objectId,omitempty"`
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
type ServicePrincipalProfile struct {
	ClientID          string             `json:"clientId"`
	Secret            string             `json:"secret,omitempty" conform:"redact"`
	ObjectID          string             `json:"objectId,omitempty"`
	KeyvaultSecretRef *KeyvaultSecretRef `json:"keyvaultSecretRef,omitempty"`
}

// KeyvaultSecretRef specifies path to the Azure keyvault along with secret name and (optionaly) version
// for Service Principal's secret
type KeyvaultSecretRef struct {
	VaultID       string `json:"vaultID"`
	SecretName    string `json:"secretName"`
	SecretVersion string `json:"version,omitempty"`
}

// CertificateProfile represents the definition of the master cluster
type CertificateProfile struct {
	// CaCertificate is the certificate authority certificate.
	CaCertificate string `json:"caCertificate,omitempty" conform:"redact"`
	// CaPrivateKey is the certificate authority key.
	CaPrivateKey string `json:"caPrivateKey,omitempty" conform:"redact"`
	// ApiServerCertificate is the rest api server certificate, and signed by the CA
	APIServerCertificate string `json:"apiServerCertificate,omitempty" conform:"redact"`
	// ApiServerPrivateKey is the rest api server private key, and signed by the CA
	APIServerPrivateKey string `json:"apiServerPrivateKey,omitempty" conform:"redact"`
	// ClientCertificate is the certificate used by the client kubelet services and signed by the CA
	ClientCertificate string `json:"clientCertificate,omitempty" conform:"redact"`
	// ClientPrivateKey is the private key used by the client kubelet services and signed by the CA
	ClientPrivateKey string `json:"clientPrivateKey,omitempty" conform:"redact"`
	// KubeConfigCertificate is the client certificate used for kubectl cli and signed by the CA
	KubeConfigCertificate string `json:"kubeConfigCertificate,omitempty" conform:"redact"`
	// KubeConfigPrivateKey is the client private key used for kubectl cli and signed by the CA
	KubeConfigPrivateKey string `json:"kubeConfigPrivateKey,omitempty" conform:"redact"`
	// EtcdServerCertificate is the server certificate for etcd, and signed by the CA
	EtcdServerCertificate string `json:"etcdServerCertificate,omitempty" conform:"redact"`
	// EtcdServerPrivateKey is the server private key for etcd, and signed by the CA
	EtcdServerPrivateKey string `json:"etcdServerPrivateKey,omitempty" conform:"redact"`
	// EtcdClientCertificate is etcd client certificate, and signed by the CA
	EtcdClientCertificate string `json:"etcdClientCertificate,omitempty" conform:"redact"`
	// EtcdClientPrivateKey is the etcd client private key, and signed by the CA
	EtcdClientPrivateKey string `json:"etcdClientPrivateKey,omitempty" conform:"redact"`
	// EtcdPeerCertificates is list of etcd peer certificates, and signed by the CA
	EtcdPeerCertificates []string `json:"etcdPeerCertificates,omitempty" conform:"redact"`
	// EtcdPeerPrivateKeys is list of etcd peer private keys, and signed by the CA
	EtcdPeerPrivateKeys []string `json:"etcdPeerPrivateKeys,omitempty" conform:"redact"`
}

// LinuxProfile represents the linux parameters passed to the cluster
type LinuxProfile struct {
	AdminUsername string `json:"adminUsername"`
	SSH           struct {
		PublicKeys []PublicKey `json:"publicKeys"`
	} `json:"ssh"`
	Secrets               []KeyVaultSecrets   `json:"secrets,omitempty"`
	Distro                Distro              `json:"distro,omitempty"`
	ScriptRootURL         string              `json:"scriptroot,omitempty"`
	CustomSearchDomain    *CustomSearchDomain `json:"customSearchDomain,omitempty"`
	CustomNodesDNS        *CustomNodesDNS     `json:"CustomNodesDNS,omitempty"`
	IsSSHKeyAutoGenerated *bool               `json:"isSSHKeyAutoGenerated,omitempty"`
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

// CustomNodesDNS represents the Search Domain when the custom vnet for a custom DNS as a nameserver.
type CustomNodesDNS struct {
	DNSServer string `json:"dnsServer,omitempty"`
}

// WindowsProfile represents the windows parameters passed to the cluster
type WindowsProfile struct {
	AdminUsername             string            `json:"adminUsername"`
	AdminPassword             string            `json:"adminPassword" conform:"redact"`
	CSIProxyURL               string            `json:"csiProxyURL,omitempty"`
	EnableCSIProxy            *bool             `json:"enableCSIProxy,omitempty"`
	ImageRef                  *ImageReference   `json:"imageReference,omitempty"`
	ImageVersion              string            `json:"imageVersion"`
	WindowsImageSourceURL     string            `json:"windowsImageSourceURL"`
	WindowsPublisher          string            `json:"windowsPublisher"`
	WindowsOffer              string            `json:"windowsOffer"`
	WindowsSku                string            `json:"windowsSku"`
	WindowsDockerVersion      string            `json:"windowsDockerVersion"`
	Secrets                   []KeyVaultSecrets `json:"secrets,omitempty"`
	SSHEnabled                *bool             `json:"sshEnabled,omitempty"`
	EnableAutomaticUpdates    *bool             `json:"enableAutomaticUpdates,omitempty"`
	IsCredentialAutoGenerated *bool             `json:"isCredentialAutoGenerated,omitempty"`
}

// ProvisioningState represents the current state of container service resource.
type ProvisioningState string

const (
	// Creating means ContainerService resource is being created.
	Creating ProvisioningState = "Creating"
	// Updating means an existing ContainerService resource is being updated
	Updating ProvisioningState = "Updating"
	// Scaling means an existing ContainerService resource is being scaled only
	Scaling ProvisioningState = "Scaling"
	// Failed means resource is in failed state
	Failed ProvisioningState = "Failed"
	// Succeeded means resource created succeeded during last create/update
	Succeeded ProvisioningState = "Succeeded"
	// Deleting means resource is in the process of being deleted
	Deleting ProvisioningState = "Deleting"
	// Migrating means resource is being migrated from one subscription or
	// resource group to another
	Migrating ProvisioningState = "Migrating"
	// Upgrading means an existing ContainerService resource is being upgraded
	Upgrading ProvisioningState = "Upgrading"
)

// OrchestratorProfile contains Orchestrator properties
type OrchestratorProfile struct {
	OrchestratorType    string            `json:"orchestratorType"`
	OrchestratorVersion string            `json:"orchestratorVersion"`
	KubernetesConfig    *KubernetesConfig `json:"kubernetesConfig,omitempty"`
	DcosConfig          *DcosConfig       `json:"dcosConfig,omitempty"`
}

// OrchestratorVersionProfile contains information of a supported orchestrator version:
type OrchestratorVersionProfile struct {
	// Orchestrator type and version
	OrchestratorProfile
	// Whether this orchestrator version is deployed by default if orchestrator release is not specified
	Default bool `json:"default,omitempty"`
	// List of available upgrades for this orchestrator version
	Upgrades []*OrchestratorProfile `json:"upgrades,omitempty"`
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

// IsDisabled returns true if the addon is explicitly disabled
func (a *KubernetesAddon) IsDisabled() bool {
	if a.Enabled == nil {
		return false
	}
	return !*a.Enabled
}

// GetAddonContainersIndexByName returns the KubernetesAddon containers index with the name `containerName`
func (a KubernetesAddon) GetAddonContainersIndexByName(containerName string) int {
	for i := range a.Containers {
		if a.Containers[i].Name == containerName {
			return i
		}
	}
	return -1
}

// GetAddonPoolIndexByName returns the KubernetesAddon pools index with the name `poolName`
func (a KubernetesAddon) GetAddonPoolIndexByName(poolName string) int {
	for i := range a.Pools {
		if a.Pools[i].Name == poolName {
			return i
		}
	}
	return -1
}

// KubernetesComponent defines a component w/ configuration to include with the cluster deployment
type KubernetesComponent struct {
	Name       string                    `json:"name,omitempty"`
	Enabled    *bool                     `json:"enabled,omitempty"`
	Containers []KubernetesContainerSpec `json:"containers,omitempty"`
	Config     map[string]string         `json:"config,omitempty"`
	Data       string                    `json:"data,omitempty"`
}

// IsEnabled returns true if the component is enabled
func (c *KubernetesComponent) IsEnabled() bool {
	if c.Enabled == nil {
		return false
	}
	return *c.Enabled
}

// IsDisabled returns true if the component is explicitly disabled
func (c *KubernetesComponent) IsDisabled() bool {
	if c.Enabled == nil {
		return false
	}
	return !*c.Enabled
}

// GetContainersIndexByName returns the KubernetesAddon containers index with the name `containerName`
func (c KubernetesComponent) GetContainersIndexByName(containerName string) int {
	for i := range c.Containers {
		if c.Containers[i].Name == containerName {
			return i
		}
	}
	return -1
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

// CloudProviderConfig contains the KubernetesConfig properties specific to the Cloud Provider
type CloudProviderConfig struct {
	CloudProviderBackoffMode          string `json:"cloudProviderBackoffMode,omitempty"`
	CloudProviderBackoff              *bool  `json:"cloudProviderBackoff,omitempty"`
	CloudProviderBackoffRetries       int    `json:"cloudProviderBackoffRetries,omitempty"`
	CloudProviderBackoffJitter        string `json:"cloudProviderBackoffJitter,omitempty"`
	CloudProviderBackoffDuration      int    `json:"cloudProviderBackoffDuration,omitempty"`
	CloudProviderBackoffExponent      string `json:"cloudProviderBackoffExponent,omitempty"`
	CloudProviderRateLimit            *bool  `json:"cloudProviderRateLimit,omitempty"`
	CloudProviderRateLimitQPS         string `json:"cloudProviderRateLimitQPS,omitempty"`
	CloudProviderRateLimitQPSWrite    string `json:"cloudProviderRateLimitQPSWrite,omitempty"`
	CloudProviderRateLimitBucket      int    `json:"cloudProviderRateLimitBucket,omitempty"`
	CloudProviderRateLimitBucketWrite int    `json:"cloudProviderRateLimitBucketWrite,omitempty"`
	CloudProviderDisableOutboundSNAT  *bool  `json:"cloudProviderDisableOutboundSNAT,omitempty"`
}

// KubernetesConfigDeprecated are properties that are no longer operable and will be ignored
// TODO use this when strict JSON checking accommodates struct embedding
type KubernetesConfigDeprecated struct {
	NonMasqueradeCidr                string `json:"nonMasqueradeCidr,omitempty"`
	NodeStatusUpdateFrequency        string `json:"nodeStatusUpdateFrequency,omitempty"`
	HardEvictionThreshold            string `json:"hardEvictionThreshold,omitempty"`
	CtrlMgrNodeMonitorGracePeriod    string `json:"ctrlMgrNodeMonitorGracePeriod,omitempty"`
	CtrlMgrPodEvictionTimeout        string `json:"ctrlMgrPodEvictionTimeout,omitempty"`
	CtrlMgrRouteReconciliationPeriod string `json:"ctrlMgrRouteReconciliationPeriod,omitempty"`
}

// KubeProxyMode is for iptables and ipvs (and future others)
type KubeProxyMode string

// We currently support ipvs and iptables
const (
	// KubeProxyModeIPTables is used to set the kube-proxy to iptables mode
	KubeProxyModeIPTables KubeProxyMode = "iptables"
	// KubeProxyModeIPVS is used to set the kube-proxy to ipvs mode
	KubeProxyModeIPVS KubeProxyMode = "ipvs"
)

// KubernetesConfig contains the Kubernetes config structure, containing
// Kubernetes specific configuration
type KubernetesConfig struct {
	KubernetesImageBase               string                `json:"kubernetesImageBase,omitempty"`
	KubernetesImageBaseType           string                `json:"kubernetesImageBaseType,omitempty"`
	MCRKubernetesImageBase            string                `json:"mcrKubernetesImageBase,omitempty"`
	ClusterSubnet                     string                `json:"clusterSubnet,omitempty"`
	NetworkPolicy                     string                `json:"networkPolicy,omitempty"`
	NetworkPlugin                     string                `json:"networkPlugin,omitempty"`
	NetworkMode                       string                `json:"networkMode,omitempty"`
	ContainerRuntime                  string                `json:"containerRuntime,omitempty"`
	MaxPods                           int                   `json:"maxPods,omitempty"`
	DockerBridgeSubnet                string                `json:"dockerBridgeSubnet,omitempty"`
	DNSServiceIP                      string                `json:"dnsServiceIP,omitempty"`
	ServiceCIDR                       string                `json:"serviceCidr,omitempty"`
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
	CustomCcmImage                    string                `json:"customCcmImage,omitempty"` // Image for cloud-controller-manager
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
	KubeletConfig                     map[string]string     `json:"kubeletConfig,omitempty"`
	ContainerRuntimeConfig            map[string]string     `json:"containerRuntimeConfig"`
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
	NonMasqueradeCidr                 string                `json:"nonMasqueradeCidr,omitempty"`
	NodeStatusUpdateFrequency         string                `json:"nodeStatusUpdateFrequency,omitempty"`
	HardEvictionThreshold             string                `json:"hardEvictionThreshold,omitempty"`
	CtrlMgrNodeMonitorGracePeriod     string                `json:"ctrlMgrNodeMonitorGracePeriod,omitempty"`
	CtrlMgrPodEvictionTimeout         string                `json:"ctrlMgrPodEvictionTimeout,omitempty"`
	CtrlMgrRouteReconciliationPeriod  string                `json:"ctrlMgrRouteReconciliationPeriod,omitempty"`
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

// HasPrivateRegistry returns if a private registry is specified
func (d *DcosConfig) HasPrivateRegistry() bool {
	return len(d.Registry) > 0
}

// HasBootstrap returns if a bootstrap profile is specified
func (d *DcosConfig) HasBootstrap() bool {
	return d.BootstrapProfile != nil
}

// MasterProfile represents the definition of the master cluster
type MasterProfile struct {
	Count                     int               `json:"count"`
	DNSPrefix                 string            `json:"dnsPrefix"`
	SubjectAltNames           []string          `json:"subjectAltNames"`
	VMSize                    string            `json:"vmSize"`
	OSDiskSizeGB              int               `json:"osDiskSizeGB,omitempty"`
	VnetSubnetID              string            `json:"vnetSubnetID,omitempty"`
	VnetCidr                  string            `json:"vnetCidr,omitempty"`
	AgentVnetSubnetID         string            `json:"agentVnetSubnetID,omitempty"`
	FirstConsecutiveStaticIP  string            `json:"firstConsecutiveStaticIP,omitempty"`
	Subnet                    string            `json:"subnet"`
	SubnetIPv6                string            `json:"subnetIPv6"`
	IPAddressCount            int               `json:"ipAddressCount,omitempty"`
	StorageProfile            string            `json:"storageProfile,omitempty"`
	HTTPSourceAddressPrefix   string            `json:"HTTPSourceAddressPrefix,omitempty"`
	OAuthEnabled              bool              `json:"oauthEnabled"`
	PreprovisionExtension     *Extension        `json:"preProvisionExtension"`
	Extensions                []Extension       `json:"extensions"`
	Distro                    Distro            `json:"distro,omitempty"`
	KubernetesConfig          *KubernetesConfig `json:"kubernetesConfig,omitempty"`
	ImageRef                  *ImageReference   `json:"imageReference,omitempty"`
	CustomFiles               *[]CustomFile     `json:"customFiles,omitempty"`
	AvailabilityProfile       string            `json:"availabilityProfile"`
	PlatformFaultDomainCount  *int              `json:"platformFaultDomainCount"`
	PlatformUpdateDomainCount *int              `json:"platformUpdateDomainCount"`
	AgentSubnet               string            `json:"agentSubnet,omitempty"`
	AvailabilityZones         []string          `json:"availabilityZones,omitempty"`
	SinglePlacementGroup      *bool             `json:"singlePlacementGroup,omitempty"`
	AuditDEnabled             *bool             `json:"auditDEnabled,omitempty"`
	UltraSSDEnabled           *bool             `json:"ultraSSDEnabled,omitempty"`
	EncryptionAtHost          *bool             `json:"encryptionAtHost,omitempty"`
	CustomVMTags              map[string]string `json:"customVMTags,omitempty"`
	// Master LB public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`
	// True: uses cosmos etcd endpoint instead of installing etcd on masters
	CosmosEtcd                *bool             `json:"cosmosEtcd,omitempty"`
	SysctlDConfig             map[string]string `json:"sysctldConfig,omitempty"`
	ProximityPlacementGroupID string            `json:"proximityPlacementGroupID,omitempty"`
	OSDiskCachingType         string            `json:"osDiskCachingType,omitempty"`
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
	Name                                string               `json:"name"`
	Count                               int                  `json:"count"`
	VMSize                              string               `json:"vmSize"`
	OSDiskSizeGB                        int                  `json:"osDiskSizeGB,omitempty"`
	DNSPrefix                           string               `json:"dnsPrefix,omitempty"`
	OSType                              OSType               `json:"osType,omitempty"`
	Ports                               []int                `json:"ports,omitempty"`
	ProvisioningState                   ProvisioningState    `json:"provisioningState,omitempty"`
	AvailabilityProfile                 string               `json:"availabilityProfile"`
	ScaleSetPriority                    string               `json:"scaleSetPriority,omitempty"`
	ScaleSetEvictionPolicy              string               `json:"scaleSetEvictionPolicy,omitempty"`
	SpotMaxPrice                        *float64             `json:"spotMaxPrice,omitempty"`
	StorageProfile                      string               `json:"storageProfile,omitempty"`
	DiskSizesGB                         []int                `json:"diskSizesGB,omitempty"`
	VnetSubnetID                        string               `json:"vnetSubnetID,omitempty"`
	Subnet                              string               `json:"subnet"`
	IPAddressCount                      int                  `json:"ipAddressCount,omitempty"`
	Distro                              Distro               `json:"distro,omitempty"`
	Role                                AgentPoolProfileRole `json:"role,omitempty"`
	AcceleratedNetworkingEnabled        *bool                `json:"acceleratedNetworkingEnabled,omitempty"`
	AcceleratedNetworkingEnabledWindows *bool                `json:"acceleratedNetworkingEnabledWindows,omitempty"`
	VMSSOverProvisioningEnabled         *bool                `json:"vmssOverProvisioningEnabled,omitempty"`
	FQDN                                string               `json:"fqdn,omitempty"`
	CustomNodeLabels                    map[string]string    `json:"customNodeLabels,omitempty"`
	PreprovisionExtension               *Extension           `json:"preProvisionExtension"`
	Extensions                          []Extension          `json:"extensions"`
	KubernetesConfig                    *KubernetesConfig    `json:"kubernetesConfig,omitempty"`
	OrchestratorVersion                 string               `json:"orchestratorVersion"`
	ImageRef                            *ImageReference      `json:"imageReference,omitempty"`
	MaxCount                            *int                 `json:"maxCount,omitempty"`
	MinCount                            *int                 `json:"minCount,omitempty"`
	EnableAutoScaling                   *bool                `json:"enableAutoScaling,omitempty"`
	AvailabilityZones                   []string             `json:"availabilityZones,omitempty"`
	PlatformFaultDomainCount            *int                 `json:"platformFaultDomainCount"`
	PlatformUpdateDomainCount           *int                 `json:"platformUpdateDomainCount"`
	SinglePlacementGroup                *bool                `json:"singlePlacementGroup,omitempty"`
	VnetCidrs                           []string             `json:"vnetCidrs,omitempty"`
	PreserveNodesProperties             *bool                `json:"preserveNodesProperties,omitempty"`
	WindowsNameVersion                  string               `json:"windowsNameVersion,omitempty"`
	EnableVMSSNodePublicIP              *bool                `json:"enableVMSSNodePublicIP,omitempty"`
	LoadBalancerBackendAddressPoolIDs   []string             `json:"loadBalancerBackendAddressPoolIDs,omitempty"`
	AuditDEnabled                       *bool                `json:"auditDEnabled,omitempty"`
	CustomVMTags                        map[string]string    `json:"customVMTags,omitempty"`
	DiskEncryptionSetID                 string               `json:"diskEncryptionSetID,omitempty"`
	SysctlDConfig                       map[string]string    `json:"sysctldConfig,omitempty"`
	UltraSSDEnabled                     *bool                `json:"ultraSSDEnabled,omitempty"`
	EncryptionAtHost                    *bool                `json:"encryptionAtHost,omitempty"`
	ProximityPlacementGroupID           string               `json:"proximityPlacementGroupID,omitempty"`
	OSDiskCachingType                   string               `json:"osDiskCachingType,omitempty"`
	DataDiskCachingType                 string               `json:"dataDiskCachingType,omitempty"`
}

// AgentPoolProfileRole represents an agent role
type AgentPoolProfileRole string

// DiagnosticsProfile setting to enable/disable capturing
// diagnostics for VMs hosting container cluster.
type DiagnosticsProfile struct {
	VMDiagnostics *VMDiagnostics `json:"vmDiagnostics"`
}

// VMDiagnostics contains settings to on/off boot diagnostics collection
// in RD Host
type VMDiagnostics struct {
	Enabled bool `json:"enabled"`

	// Specifies storage account Uri where Boot Diagnostics (CRP &
	// VMSS BootDiagostics) and VM Diagnostics logs (using Linux
	// Diagnostics Extension) will be stored. Uri will be of standard
	// blob domain. i.e. https://storageaccount.blob.core.windows.net/
	// This field is readonly as ACS RP will create a storage account
	// for the customer.
	StorageURL *neturl.URL `json:"storageUrl"`
}

// JumpboxProfile describes properties of the jumpbox setup
// in the AKS container cluster.
type JumpboxProfile struct {
	OSType    OSType `json:"osType"`
	DNSPrefix string `json:"dnsPrefix"`

	// Jumpbox public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`
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

// HostedMasterProfile defines properties for a hosted master
type HostedMasterProfile struct {
	// Master public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GETFQDN
	FQDN      string `json:"fqdn,omitempty"`
	DNSPrefix string `json:"dnsPrefix"`
	// Subnet holds the CIDR which defines the Azure Subnet in which
	// Agents will be provisioned. This is stored on the HostedMasterProfile
	// and will become `masterSubnet` in the compiled template.
	Subnet string `json:"subnet"`
	// ApiServerWhiteListRange is a comma delimited CIDR which is whitelisted to AKS
	APIServerWhiteListRange *string `json:"apiServerWhiteListRange"`
	IPMasqAgent             bool    `json:"ipMasqAgent"`
}

// AuthenticatorType represents the authenticator type the cluster was
// set up with.
type AuthenticatorType string

const (
	// OIDC represent cluster setup in OIDC auth mode
	OIDC AuthenticatorType = "oidc"
	// Webhook represent cluster setup in wehhook auth mode
	Webhook AuthenticatorType = "webhook"
)

// AADProfile specifies attributes for AAD integration
type AADProfile struct {
	// The client AAD application ID.
	ClientAppID string `json:"clientAppID,omitempty"`
	// The server AAD application ID.
	ServerAppID string `json:"serverAppID,omitempty"`
	// The server AAD application secret
	ServerAppSecret string `json:"serverAppSecret,omitempty" conform:"redact"`
	// The AAD tenant ID to use for authentication.
	// If not specified, will use the tenant of the deployment subscription.
	// Optional
	TenantID string `json:"tenantID,omitempty"`
	// The Azure Active Directory Group Object ID that will be assigned the
	// cluster-admin RBAC role.
	// Optional
	AdminGroupID string `json:"adminGroupID,omitempty"`
	// The authenticator to use, either "oidc" or "webhook".
	Authenticator AuthenticatorType `json:"authenticator"`
}

// CustomProfile specifies custom properties that are used for
// cluster instantiation.  Should not be used by most users.
type CustomProfile struct {
	Orchestrator string `json:"orchestrator,omitempty"`
}

// VlabsARMContainerService is the type we read and write from file
// needed because the json that is sent to ARM and aks-engine
// is different from the json that the ACS RP Api gets from ARM
type VlabsARMContainerService struct {
	TypeMeta
	*vlabs.ContainerService
}

// V20170831ARMManagedContainerService is the type we read and write from file
// needed because the json that is sent to ARM and aks-engine
// is different from the json that the ACS RP Api gets from ARM
type V20170831ARMManagedContainerService struct {
	TypeMeta
	*v20170831.ManagedCluster
}

// V20180331ARMManagedContainerService is the type we read and write from file
// needed because the json that is sent to ARM and aks-engine
// is different from the json that the ACS RP Api gets from ARM
type V20180331ARMManagedContainerService struct {
	TypeMeta
	*v20180331.ManagedCluster
}

// AzureStackMetadataEndpoints is the type for Azure Stack metadata endpoints
type AzureStackMetadataEndpoints struct {
	GalleryEndpoint string                            `json:"galleryEndpoint,omitempty"`
	GraphEndpoint   string                            `json:"graphEndpoint,omitempty"`
	PortalEndpoint  string                            `json:"portalEndpoint,omitempty"`
	Authentication  *AzureStackMetadataAuthentication `json:"authentication,omitempty"`
}

// AzureStackMetadataAuthentication is the type for Azure Stack metadata authentication endpoints
type AzureStackMetadataAuthentication struct {
	LoginEndpoint string   `json:"loginEndpoint,omitempty"`
	Audiences     []string `json:"audiences,omitempty"`
}

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

// HasManagedDisks returns true if the cluster contains Managed Disks
func (p *Properties) HasManagedDisks() bool {
	if p.MasterProfile != nil && p.MasterProfile.StorageProfile == ManagedDisks {
		return true
	}
	for _, agentPoolProfile := range p.AgentPoolProfiles {
		if agentPoolProfile.StorageProfile == ManagedDisks {
			return true
		}
	}
	if p.OrchestratorProfile != nil && p.OrchestratorProfile.KubernetesConfig != nil && p.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() && p.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == ManagedDisks {
		return true
	}
	return false
}

// HasStorageAccountDisks returns true if the cluster contains Storage Account Disks
func (p *Properties) HasStorageAccountDisks() bool {
	if p.MasterProfile != nil && p.MasterProfile.StorageProfile == StorageAccount {
		return true
	}
	for _, agentPoolProfile := range p.AgentPoolProfiles {
		if agentPoolProfile.StorageProfile == StorageAccount {
			return true
		}
	}
	if p.OrchestratorProfile != nil && p.OrchestratorProfile.KubernetesConfig != nil && p.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() && p.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == StorageAccount {
		return true
	}
	return false
}

// HasStorageAccountDisks returns true if the cluster contains agent pools with Ephemeral Disks
func (p *Properties) HasEphemeralDisks() bool {
	for _, agentPoolProfile := range p.AgentPoolProfiles {
		if agentPoolProfile.StorageProfile == Ephemeral {
			return true
		}
	}
	return false
}

// TotalNodes returns the total number of nodes in the cluster configuration
func (p *Properties) TotalNodes() int {
	var totalNodes int
	if p.MasterProfile != nil {
		totalNodes = p.MasterProfile.Count
	}
	for _, pool := range p.AgentPoolProfiles {
		totalNodes += pool.Count
	}
	return totalNodes
}

// HasVMSSAgentPool returns true if the cluster contains Virtual Machine Scale Sets agent pools
func (p *Properties) HasVMSSAgentPool() bool {
	for _, agentPoolProfile := range p.AgentPoolProfiles {
		if agentPoolProfile.AvailabilityProfile == VirtualMachineScaleSets {
			return true
		}
	}
	return false
}

// K8sOrchestratorName returns the 3 character orchestrator code for kubernetes-based clusters.
func (p *Properties) K8sOrchestratorName() string {
	if p.OrchestratorProfile.IsKubernetes() {
		if p.HostedMasterProfile != nil {
			return DefaultHostedProfileMasterName
		}
		return DefaultOrchestratorName
	}
	return ""
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

// GetAgentPoolIndexByName returns the index of the provided agentpool.
func (p *Properties) GetAgentPoolIndexByName(name string) int {
	index := -1
	for i, profile := range p.AgentPoolProfiles {
		if profile.Name == name {
			index = i
			break
		}
	}
	return index
}

// GetAgentVMPrefix returns the VM prefix for an agentpool.
func (p *Properties) GetAgentVMPrefix(a *AgentPoolProfile, index int) string {
	nameSuffix := p.GetClusterID()
	vmPrefix := ""
	if index != -1 {
		if a.IsWindows() {
			if a.WindowsNameVersion == "v2" {
				vmPrefix = p.K8sOrchestratorName() + a.Name
			} else {
				vmPrefix = nameSuffix[:4] + p.K8sOrchestratorName() + fmt.Sprintf("%02d", index)
			}
		} else {
			vmPrefix = p.K8sOrchestratorName() + "-" + a.Name + "-" + nameSuffix + "-"
			if a.IsVirtualMachineScaleSets() {
				vmPrefix += "vmss"
			}
		}
	}
	return vmPrefix
}

// GetVMType returns the type of VM "vmss" or "standard" to be passed to the cloud provider
func (p *Properties) GetVMType() string {
	if p.HasVMSSAgentPool() {
		return VMSSVMType
	}
	return StandardVMType
}

// HasVMASAgentPool checks whether any of the agents in the AgentPool use VMAS or not
func (p *Properties) HasVMASAgentPool() bool {
	for _, agentProfile := range p.AgentPoolProfiles {
		if agentProfile.IsAvailabilitySets() {
			return true
		}
	}
	return false
}

// AnyAgentIsLinux checks whether any of the agents in the AgentPools are linux
func (p *Properties) AnyAgentIsLinux() bool {
	for _, agentProfile := range p.AgentPoolProfiles {
		if agentProfile.IsLinux() {
			return true
		}
	}
	return false
}

// GetMasterVMPrefix returns the prefix of master VMs
func (p *Properties) GetMasterVMPrefix() string {
	return p.K8sOrchestratorName() + "-master-" + p.GetClusterID() + "-"
}

// GetResourcePrefix returns the prefix to use for naming cluster resources
func (p *Properties) GetResourcePrefix() string {
	if p.IsHostedMasterProfile() {
		return p.K8sOrchestratorName() + "-agentpool-" + p.GetClusterID() + "-"
	}
	return p.K8sOrchestratorName() + "-master-" + p.GetClusterID() + "-"

}

// GetRouteTableName returns the route table name of the cluster.
func (p *Properties) GetRouteTableName() string {
	return p.GetResourcePrefix() + "routetable"
}

// GetNSGName returns the name of the network security group of the cluster.
func (p *Properties) GetNSGName() string {
	return p.GetResourcePrefix() + "nsg"
}

// GetPrimaryAvailabilitySetName returns the name of the primary availability set of the cluster
func (p *Properties) GetPrimaryAvailabilitySetName() string {
	if len(p.AgentPoolProfiles) > 0 {
		if p.AgentPoolProfiles[0].AvailabilityProfile == AvailabilitySet {
			return p.AgentPoolProfiles[0].Name + "-availabilitySet-" + p.GetClusterID()
		}
	}
	return ""
}

// GetPrimaryScaleSetName returns the name of the primary scale set node of the cluster
func (p *Properties) GetPrimaryScaleSetName() string {
	if len(p.AgentPoolProfiles) > 0 {
		if p.AgentPoolProfiles[0].AvailabilityProfile == VirtualMachineScaleSets {
			return p.GetAgentVMPrefix(p.AgentPoolProfiles[0], 0)
		}
	}
	return ""
}

// IsHostedMasterProfile returns true if the cluster has a hosted master
func (p *Properties) IsHostedMasterProfile() bool {
	return p.HostedMasterProfile != nil
}

// IsIPMasqAgentEnabled returns true if the cluster has a hosted master and IpMasqAgent is disabled
func (p *Properties) IsIPMasqAgentEnabled() bool {
	if p.HostedMasterProfile != nil {
		return p.HostedMasterProfile.IPMasqAgent
	}
	return p.OrchestratorProfile.KubernetesConfig.IsIPMasqAgentEnabled()
}

// IsIPMasqAgentDisabled returns true if the ip-masq-agent functionality is disabled
func (p *Properties) IsIPMasqAgentDisabled() bool {
	if p.HostedMasterProfile != nil {
		return !p.HostedMasterProfile.IPMasqAgent
	}
	if p.OrchestratorProfile != nil && p.OrchestratorProfile.KubernetesConfig != nil {
		return p.OrchestratorProfile.KubernetesConfig.IsIPMasqAgentDisabled()
	}
	return false
}

// GetVNetResourceGroupName returns the virtual network resource group name of the cluster
func (p *Properties) GetVNetResourceGroupName() string {
	var vnetResourceGroupName string
	if p.IsHostedMasterProfile() && p.AreAgentProfilesCustomVNET() {
		vnetResourceGroupName = strings.Split(p.AgentPoolProfiles[0].VnetSubnetID, "/")[DefaultVnetResourceGroupSegmentIndex]
	} else if !p.IsHostedMasterProfile() && p.MasterProfile.IsCustomVNET() {
		vnetResourceGroupName = strings.Split(p.MasterProfile.VnetSubnetID, "/")[DefaultVnetResourceGroupSegmentIndex]
	}
	return vnetResourceGroupName
}

// GetVirtualNetworkName returns the virtual network name of the cluster
func (p *Properties) GetVirtualNetworkName() string {
	var vnetName string
	if p.IsHostedMasterProfile() && p.AreAgentProfilesCustomVNET() {
		vnetName = strings.Split(p.AgentPoolProfiles[0].VnetSubnetID, "/")[DefaultVnetNameResourceSegmentIndex]
	} else if !p.IsHostedMasterProfile() && p.MasterProfile.IsCustomVNET() {
		vnetName = strings.Split(p.MasterProfile.VnetSubnetID, "/")[DefaultVnetNameResourceSegmentIndex]
	} else {
		vnetName = p.K8sOrchestratorName() + "-vnet-" + p.GetClusterID()
	}
	return vnetName
}

// GetSubnetName returns the subnet name of the cluster based on its current configuration.
func (p *Properties) GetSubnetName() string {
	var subnetName string

	if !p.IsHostedMasterProfile() {
		if p.MasterProfile.IsCustomVNET() {
			subnetName = strings.Split(p.MasterProfile.VnetSubnetID, "/")[DefaultSubnetNameResourceSegmentIndex]
		} else if p.MasterProfile.IsVirtualMachineScaleSets() {
			subnetName = "subnetmaster"
		} else {
			subnetName = p.K8sOrchestratorName() + "-subnet"
		}
	} else {
		if p.AreAgentProfilesCustomVNET() {
			subnetName = strings.Split(p.AgentPoolProfiles[0].VnetSubnetID, "/")[DefaultSubnetNameResourceSegmentIndex]
		} else {
			subnetName = p.K8sOrchestratorName() + "-subnet"
		}
	}
	return subnetName
}

// GetDNSPrefix returns the the string used as master FQDN prefix
func (p *Properties) GetDNSPrefix() string {
	if p.MasterProfile != nil {
		// MasterProfile exists, uses master DNS prefix
		return strings.ToLower(p.MasterProfile.DNSPrefix)
	} else if p.HostedMasterProfile != nil {
		return strings.ToLower(p.HostedMasterProfile.DNSPrefix)
	}
	return ""
}

// AreAgentProfilesCustomVNET returns true if all of the agent profiles in the clusters are configured with VNET.
func (p *Properties) AreAgentProfilesCustomVNET() bool {
	if p.AgentPoolProfiles != nil {
		for _, agentPoolProfile := range p.AgentPoolProfiles {
			if !agentPoolProfile.IsCustomVNET() {
				return false
			}
		}
		return true
	}
	return false
}

// GetClusterID creates a unique 8 string cluster ID.
func (p *Properties) GetClusterID() string {
	if p.ClusterID == "" {
		uniqueNameSuffixSize := 8
		// the name suffix uniquely identifies the cluster and is generated off a hash
		// from the master dns name
		h := fnv.New64a()
		if p.MasterProfile != nil {
			_, _ = h.Write([]byte(p.MasterProfile.DNSPrefix))
		} else if p.HostedMasterProfile != nil {
			_, _ = h.Write([]byte(p.HostedMasterProfile.DNSPrefix))
		} else if len(p.AgentPoolProfiles) > 0 {
			_, _ = h.Write([]byte(p.AgentPoolProfiles[0].Name))
		}
		r := rand.New(rand.NewSource(int64(h.Sum64())))
		p.ClusterID = fmt.Sprintf("%08d", r.Uint32())[:uniqueNameSuffixSize]
	}
	return p.ClusterID
}

// GetClusterMetadata returns a instance of the struct type api.ClusterMetadata.
func (p *Properties) GetClusterMetadata() *ClusterMetadata {
	return &ClusterMetadata{
		SubnetName:                 p.GetSubnetName(),
		VNetResourceGroupName:      p.GetVNetResourceGroupName(),
		VirtualNetworkName:         p.GetVirtualNetworkName(),
		SecurityGroupName:          p.GetNSGName(),
		RouteTableName:             p.GetRouteTableName(),
		PrimaryAvailabilitySetName: p.GetPrimaryAvailabilitySetName(),
		PrimaryScaleSetName:        p.GetPrimaryScaleSetName(),
		ResourcePrefix:             p.GetResourcePrefix(),
	}
}

// HasZonesForAllAgentPools returns true if all of the agent pools have zones
func (p *Properties) HasZonesForAllAgentPools() bool {
	if len(p.AgentPoolProfiles) > 0 {
		for _, ap := range p.AgentPoolProfiles {
			if !ap.HasAvailabilityZones() {
				return false
			}
		}
		return true
	}
	return false
}

// IsVHDDistroForAllNodes returns true if all of the agent pools plus masters are running the VHD image
func (p *Properties) IsVHDDistroForAllNodes() bool {
	if len(p.AgentPoolProfiles) > 0 {
		for _, ap := range p.AgentPoolProfiles {
			if !ap.IsVHDDistro() {
				return false
			}
		}
	}
	if p.MasterProfile != nil {
		return p.MasterProfile.IsVHDDistro()
	}
	return true
}

// HasVHDDistroNodes returns true if any one Linux node pool, including masters, are running a VHD image
func (p *Properties) HasVHDDistroNodes() bool {
	if len(p.AgentPoolProfiles) > 0 {
		for _, ap := range p.AgentPoolProfiles {
			if ap.IsVHDDistro() {
				return true
			}
		}
	}
	if p.MasterProfile != nil {
		return p.MasterProfile.IsVHDDistro()
	}
	return false
}

// IsUbuntuDistroForAllNodes returns true if all of the agent pools plus masters are running the base Ubuntu image
func (p *Properties) IsUbuntuDistroForAllNodes() bool {
	if len(p.AgentPoolProfiles) > 0 {
		for _, ap := range p.AgentPoolProfiles {
			if !ap.IsUbuntuNonVHD() {
				return false
			}
		}
	}
	if p.MasterProfile != nil {
		return p.MasterProfile.IsUbuntuNonVHD()
	}
	return true
}

// HasUbuntuDistroNodes returns true if any of the agent pools or masters are running the base Ubuntu image
func (p *Properties) HasUbuntuDistroNodes() bool {
	if len(p.AgentPoolProfiles) > 0 {
		for _, ap := range p.AgentPoolProfiles {
			if ap.IsUbuntuNonVHD() {
				return true
			}
		}
	}
	if p.MasterProfile != nil {
		return p.MasterProfile.IsUbuntuNonVHD()
	}
	return false
}

// HasUbuntu1604DistroNodes returns true if any of the agent pools or masters are running the base Ubuntu 16.04-LTS image
func (p *Properties) HasUbuntu1604DistroNodes() bool {
	if len(p.AgentPoolProfiles) > 0 {
		for _, ap := range p.AgentPoolProfiles {
			if ap.Distro == Ubuntu {
				return true
			}
		}
	}
	if p.MasterProfile != nil {
		return p.MasterProfile.Distro == Ubuntu
	}
	return false
}

// HasUbuntu1804DistroNodes returns true if any of the agent pools or masters are running the base Ubuntu 18.04-LTS image
func (p *Properties) HasUbuntu1804DistroNodes() bool {
	if len(p.AgentPoolProfiles) > 0 {
		for _, ap := range p.AgentPoolProfiles {
			switch ap.Distro {
			case Ubuntu1804, Ubuntu1804Gen2:
				return true
			}
		}
	}
	if p.MasterProfile != nil {
		switch p.MasterProfile.Distro {
		case Ubuntu1804, Ubuntu1804Gen2:
			return true
		}
	}
	return false
}

// HasAvailabilityZones returns true if the cluster contains a profile with zones
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

// HasNonRegularPriorityScaleset returns true if any one node pool has a low or spot priority scaleset configuration
func (p *Properties) HasNonRegularPriorityScaleset() bool {
	for _, agentPoolProfile := range p.AgentPoolProfiles {
		if agentPoolProfile.IsLowPriorityScaleSet() || agentPoolProfile.IsSpotScaleSet() {
			return true
		}
	}
	return false
}

// GetNonMasqueradeCIDR returns the non-masquerade CIDR for the ip-masq-agent.
func (p *Properties) GetNonMasqueradeCIDR() string {
	var nonMasqCidr string
	if !p.IsHostedMasterProfile() {
		if p.OrchestratorProfile.IsAzureCNI() {
			if p.MasterProfile != nil && p.MasterProfile.IsCustomVNET() {
				nonMasqCidr = p.MasterProfile.VnetCidr
			} else {
				nonMasqCidr = DefaultVNETCIDR
			}
		} else {
			if p.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack") {
				cidr := strings.Split(p.OrchestratorProfile.KubernetesConfig.ClusterSubnet, ",")[0]
				_, ipnet, _ := net.ParseCIDR(cidr)
				nonMasqCidr = ipnet.String()
			} else {
				nonMasqCidr = p.OrchestratorProfile.KubernetesConfig.ClusterSubnet
			}
		}
	}
	return nonMasqCidr
}

// GetSecondaryNonMasqueradeCIDR returns second cidr in case of dualstack clusters
func (p *Properties) GetSecondaryNonMasqueradeCIDR() string {
	var nonMasqCidr string
	if !p.IsHostedMasterProfile() {
		if p.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack") {
			cidr := strings.Split(p.OrchestratorProfile.KubernetesConfig.ClusterSubnet, ",")[1]
			_, ipnet, _ := net.ParseCIDR(cidr)
			nonMasqCidr = ipnet.String()
		}
	}
	return nonMasqCidr
}

// GetAzureCNICidr returns the default CNI Cidr if Azure CNI is enabled.
func (p *Properties) GetAzureCNICidr() string {
	var masqCNIIP string
	if p.OrchestratorProfile != nil && p.OrchestratorProfile.IsAzureCNI() {
		masqCNIIP = DefaultCNICIDR
	}
	return masqCNIIP
}

// GetMasterFQDN returns the master FQDN.
func (p *Properties) GetMasterFQDN() string {
	if p.IsHostedMasterProfile() {
		return p.HostedMasterProfile.FQDN
	}

	return p.MasterProfile.FQDN
}

// AnyAgentHasLoadBalancerBackendAddressPoolIDs returns true if any of the agent profiles contains LoadBalancerBackendAddressPoolIDs
func (p *Properties) AnyAgentHasLoadBalancerBackendAddressPoolIDs() bool {
	for _, agentPoolProfile := range p.AgentPoolProfiles {
		if agentPoolProfile.LoadBalancerBackendAddressPoolIDs != nil {
			return true
		}
	}
	return false
}

// GetKubeProxyFeatureGates returns the feature gates string for the kube-proxy yaml manifest
func (p *Properties) GetKubeProxyFeatureGates() string {
	if p.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack") {
		return "IPv6DualStack: true"
	}
	return "{}"
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

// IsValid returns true if ImageRefernce contains at least Name and ResourceGroup
func (i *ImageReference) IsValid() bool {
	return len(i.Name) > 0 && len(i.ResourceGroup) > 0
}

// IsGalleryImage returns true if ImageRefernce contains Gallry, Name, ResourceGroup, SubscriptionID, and Version
func (i *ImageReference) IsGalleryImage() bool {
	return len(i.Gallery) > 0 && len(i.Name) > 0 && len(i.ResourceGroup) > 0 && len(i.SubscriptionID) > 0 && len(i.Version) > 0
}

// HasImageRef returns true if the customer brought os image
func (m *MasterProfile) HasImageRef() bool {
	return m.ImageRef != nil && m.ImageRef.IsValid()
}

// HasImageGallery returns true if the customer brought os image from Shared Image Gallery
func (m *MasterProfile) HasImageGallery() bool {
	return m.ImageRef != nil && m.ImageRef.IsGalleryImage()
}

// IsCustomVNET returns true if the customer brought their own VNET
func (m *MasterProfile) IsCustomVNET() bool {
	return len(m.VnetSubnetID) > 0
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

// IsVHDDistro returns true if the distro uses VHD SKUs
func (m *MasterProfile) IsVHDDistro() bool {
	return m.Distro == AKSUbuntu1604 || m.Distro == AKSUbuntu1804
}

// IsAuditDEnabled returns true if the master profile is configured for auditd
func (m *MasterProfile) IsAuditDEnabled() bool {
	return to.Bool(m.AuditDEnabled)
}

// IsVirtualMachineScaleSets returns true if the master availability profile is VMSS
func (m *MasterProfile) IsVirtualMachineScaleSets() bool {
	return m.AvailabilityProfile == VirtualMachineScaleSets
}

// IsAvailabilitySet returns true if the master availability profile is availability set
func (m *MasterProfile) IsAvailabilitySet() bool {
	return m.AvailabilityProfile == AvailabilitySet
}

// GetFirstConsecutiveStaticIPAddress returns the first static IP address of the given subnet.
func (m *MasterProfile) GetFirstConsecutiveStaticIPAddress(subnetStr string) string {
	_, subnet, err := net.ParseCIDR(subnetStr)
	if err != nil {
		return DefaultFirstConsecutiveKubernetesStaticIP
	}

	// Find the first and last octet of the host bits.
	ones, bits := subnet.Mask.Size()
	firstOctet := ones / 8
	lastOctet := bits/8 - 1

	if m.IsVirtualMachineScaleSets() {
		subnet.IP[lastOctet] = DefaultKubernetesFirstConsecutiveStaticIPOffsetVMSS
	} else {
		// Set the remaining host bits in the first octet.
		subnet.IP[firstOctet] |= (1 << byte((8 - (ones % 8)))) - 1

		// Fill the intermediate octets with 1s and last octet with offset. This is done so to match
		// the existing behavior of allocating static IP addresses from the last /24 of the subnet.
		for i := firstOctet + 1; i < lastOctet; i++ {
			subnet.IP[i] = 255
		}
		subnet.IP[lastOctet] = DefaultKubernetesFirstConsecutiveStaticIPOffset
	}

	return subnet.IP.String()
}

// HasAvailabilityZones returns true if the master profile has availability zones
func (m *MasterProfile) HasAvailabilityZones() bool {
	return m.AvailabilityZones != nil && len(m.AvailabilityZones) > 0
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

// IsUbuntuNonVHD returns true if the distro uses a base Ubuntu image
func (m *MasterProfile) IsUbuntuNonVHD() bool {
	return m.IsUbuntu() && !m.IsVHDDistro()
}

// HasMultipleNodes returns true if there are more than one master nodes
func (m *MasterProfile) HasMultipleNodes() bool {
	return m.Count > 1
}

// HasCosmosEtcd returns true if cosmos etcd configuration is enabled
func (m *MasterProfile) HasCosmosEtcd() bool {
	return to.Bool(m.CosmosEtcd)
}

// GetCosmosEndPointURI returns the URI string for the cosmos etcd endpoint
func (m *MasterProfile) GetCosmosEndPointURI() string {
	if m.HasCosmosEtcd() {
		return fmt.Sprintf(etcdEndpointURIFmt, m.DNSPrefix)
	}
	return ""
}

// HasImageRef returns true if the customer brought os image
func (a *AgentPoolProfile) HasImageRef() bool {
	imageRef := a.ImageRef
	return imageRef != nil && imageRef.IsValid()
}

// HasImageGallery returns true if the customer brought os image from Shared Image Gallery
func (a *AgentPoolProfile) HasImageGallery() bool {
	imageRef := a.ImageRef
	return imageRef != nil && imageRef.IsGalleryImage()
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

// IsVHDDistro returns true if the distro uses VHD SKUs
func (a *AgentPoolProfile) IsVHDDistro() bool {
	return a.Distro == AKSUbuntu1604 || a.Distro == AKSUbuntu1804
}

// IsAuditDEnabled returns true if the master profile is configured for auditd
func (a *AgentPoolProfile) IsAuditDEnabled() bool {
	return to.Bool(a.AuditDEnabled)
}

// IsAvailabilitySets returns true if the customer specified disks
func (a *AgentPoolProfile) IsAvailabilitySets() bool {
	return a.AvailabilityProfile == AvailabilitySet
}

// IsVirtualMachineScaleSets returns true if the agent pool availability profile is VMSS
func (a *AgentPoolProfile) IsVirtualMachineScaleSets() bool {
	return a.AvailabilityProfile == VirtualMachineScaleSets
}

// IsLowPriorityScaleSet returns true if the VMSS is Low Priority
func (a *AgentPoolProfile) IsLowPriorityScaleSet() bool {
	return a.AvailabilityProfile == VirtualMachineScaleSets && a.ScaleSetPriority == ScaleSetPriorityLow
}

// IsSpotScaleSet returns true if the VMSS is Spot Scale Set
func (a *AgentPoolProfile) IsSpotScaleSet() bool {
	return a.AvailabilityProfile == VirtualMachineScaleSets && a.ScaleSetPriority == ScaleSetPrioritySpot
}

// IsManagedDisks returns true if the customer specified disks
func (a *AgentPoolProfile) IsManagedDisks() bool {
	return a.StorageProfile == ManagedDisks
}

// IsStorageAccount returns true if the customer specified storage account
func (a *AgentPoolProfile) IsStorageAccount() bool {
	return a.StorageProfile == StorageAccount
}

// IsStorageAccount returns true if the customer specified ephemeral disks
func (a *AgentPoolProfile) IsEphemeral() bool {
	return a.StorageProfile == Ephemeral
}

// HasDisks returns true if the customer specified disks
func (a *AgentPoolProfile) HasDisks() bool {
	return len(a.DiskSizesGB) > 0
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

// IsUbuntuNonVHD returns true if the distro uses a base Ubuntu image
func (a *AgentPoolProfile) IsUbuntuNonVHD() bool {
	return a.IsUbuntu() && !a.IsVHDDistro()
}

// RequiresCloudproviderConfig returns true if the azure.json cloudprovider config should be delivered to the nodes in this pool
func (a *AgentPoolProfile) RequiresCloudproviderConfig() bool {
	if a.KubernetesConfig != nil && a.KubernetesConfig.KubeletConfig != nil {
		if v, ok := a.KubernetesConfig.KubeletConfig["--cloud-provider"]; ok {
			if v != "" {
				return true
			}
		} else {
			return true
		}
		if v, ok := a.KubernetesConfig.KubeletConfig["--cloud-config"]; ok {
			if v != "" {
				return true
			}
		} else {
			return true
		}
		if v, ok := a.KubernetesConfig.KubeletConfig["--azure-container-registry-config"]; ok {
			if v != "" {
				return true
			}
		} else {
			return true
		}
	} else {
		return true
	}
	return false
}

// GetKubernetesLabels returns a k8s API-compliant labels string for nodes in this profile
func (a *AgentPoolProfile) GetKubernetesLabels(rg string, deprecated bool) string {
	var buf bytes.Buffer
	buf.WriteString("kubernetes.azure.com/role=agent")
	if deprecated {
		buf.WriteString(",node-role.kubernetes.io/agent=")
		buf.WriteString(",kubernetes.io/role=agent")
	}
	buf.WriteString(fmt.Sprintf(",agentpool=%s", a.Name))
	if a.StorageProfile == ManagedDisks {
		storagetier, _ := common.GetStorageAccountType(a.VMSize)
		buf.WriteString(fmt.Sprintf(",storageprofile=managed,storagetier=%s", storagetier))
	}
	if common.IsNvidiaEnabledSKU(a.VMSize) {
		accelerator := "nvidia"
		buf.WriteString(fmt.Sprintf(",accelerator=%s", accelerator))
	}
	buf.WriteString(fmt.Sprintf(",kubernetes.azure.com/cluster=%s", rg))
	keys := []string{}
	for key := range a.CustomNodeLabels {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		buf.WriteString(fmt.Sprintf(",%s=%s", key, a.CustomNodeLabels[key]))
	}
	return buf.String()
}

// IsCSIProxyEnabled returns true if csi proxy service should be enable for Windows nodes
func (w *WindowsProfile) IsCSIProxyEnabled() bool {
	if w.EnableCSIProxy != nil {
		return *w.EnableCSIProxy
	}
	return common.DefaultEnableCSIProxyWindows
}

// HasSecrets returns true if the customer specified secrets to install
func (w *WindowsProfile) HasSecrets() bool {
	return len(w.Secrets) > 0
}

// HasCustomImage returns true if there is a custom windows os image url specified
func (w *WindowsProfile) HasCustomImage() bool {
	return len(w.WindowsImageSourceURL) > 0
}

// HasImageRef returns true if the customer brought os image
func (w *WindowsProfile) HasImageRef() bool {
	return w.ImageRef != nil && w.ImageRef.IsValid()
}

// HasImageGallery returns true if the customer brought os image from Shared Image Gallery
func (w *WindowsProfile) HasImageGallery() bool {
	return w.ImageRef != nil && w.ImageRef.IsGalleryImage()
}

// GetWindowsDockerVersion gets the docker version specified or returns default value
func (w *WindowsProfile) GetWindowsDockerVersion() string {
	if w.WindowsDockerVersion != "" {
		return w.WindowsDockerVersion
	}
	return KubernetesWindowsDockerVersion
}

// GetWindowsSku gets the marketplace sku specified (such as Datacenter-Core-1809-with-Containers-smalldisk) or returns default value
func (w *WindowsProfile) GetWindowsSku() string {
	if w.WindowsSku != "" {
		return w.WindowsSku
	}
	return KubernetesDefaultWindowsSku
}

// GetSSHEnabled gets it ssh should be enabled for Windows nodes
func (w *WindowsProfile) GetSSHEnabled() bool {
	if w.SSHEnabled != nil {
		return *w.SSHEnabled
	}
	return DefaultWindowsSSHEnabled
}

// GetEnableWindowsUpdate gets the flag for enable windows update or returns the default value
func (w *WindowsProfile) GetEnableWindowsUpdate() bool {
	if w.EnableAutomaticUpdates != nil {
		return *w.EnableAutomaticUpdates
	}
	return DefaultEnableAutomaticUpdates
}

// GetIsCredentialAutoGenerated gets the flag to indicate whether the WindowsProfile is auto generated or returns the default value
func (w *WindowsProfile) GetIsCredentialAutoGenerated() bool {
	if w.IsCredentialAutoGenerated != nil {
		return *w.IsCredentialAutoGenerated
	}
	return false
}

// HasSecrets returns true if the customer specified secrets to install
func (l *LinuxProfile) HasSecrets() bool {
	return len(l.Secrets) > 0
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

// HasCustomNodesDNS returns true if the customer specified a dns server
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

// IsKubernetes returns true if this template is for Kubernetes orchestrator
func (o *OrchestratorProfile) IsKubernetes() bool {
	return o.OrchestratorType == Kubernetes
}

// IsDCOS returns true if this template is for DCOS orchestrator
func (o *OrchestratorProfile) IsDCOS() bool {
	return o.OrchestratorType == DCOS
}

// IsDCOS19 returns true if this is a DCOS 1.9 orchestrator using the latest version
func (o *OrchestratorProfile) IsDCOS19() bool {
	return o.OrchestratorType == DCOS &&
		(o.OrchestratorVersion == common.DCOSVersion1Dot9Dot0 ||
			o.OrchestratorVersion == common.DCOSVersion1Dot9Dot8)
}

// IsAzureCNI returns true if Azure CNI network plugin is enabled
func (o *OrchestratorProfile) IsAzureCNI() bool {
	if o.KubernetesConfig != nil {
		return o.KubernetesConfig.NetworkPlugin == NetworkPluginAzure
	}
	return false
}

// IsPrivateCluster returns true if this deployment is a private cluster
func (o *OrchestratorProfile) IsPrivateCluster() bool {
	if !o.IsKubernetes() {
		return false
	}
	return o.KubernetesConfig != nil && o.KubernetesConfig.PrivateCluster != nil && to.Bool(o.KubernetesConfig.PrivateCluster.Enabled)
}

// GetPodInfraContainerSpec returns the sandbox image as a string (ex: k8s.gcr.io/pause-amd64:3.1)
func (o *OrchestratorProfile) GetPodInfraContainerSpec() string {
	return o.KubernetesConfig.MCRKubernetesImageBase + GetK8sComponentsByVersionMap(o.KubernetesConfig)[o.OrchestratorVersion][common.PauseComponentName]
}

// HasAadProfile returns true if the has aad profile
func (p *Properties) HasAadProfile() bool {
	return p.AADProfile != nil
}

// GetAPIServerEtcdAPIVersion Used to set apiserver's etcdapi version
func (o *OrchestratorProfile) GetAPIServerEtcdAPIVersion() string {
	if o.KubernetesConfig != nil {
		// if we are here, version has already been validated..
		etcdVersion, _ := semver.Make(o.KubernetesConfig.EtcdVersion)
		return "etcd" + strconv.FormatUint(etcdVersion.Major, 10)
	}
	return ""
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

// GetAddonScript retrieves the raw script data specified as input for the k8s addon with name "addonName".
func (k *KubernetesConfig) GetAddonScript(addonName string) string {
	kubeAddon := k.GetAddonByName(addonName)
	return kubeAddon.Data
}

// IsAddonEnabled checks whether a k8s addon with name "addonName" is enabled or not based on the Enabled field of KubernetesAddon.
// If the value of Enabled is nil, the "defaultValue" is returned.
func (k *KubernetesConfig) IsAddonEnabled(addonName string) bool {
	kubeAddon := k.GetAddonByName(addonName)
	return kubeAddon.IsEnabled()
}

// IsAddonDisabled checks whether a k8s addon with name "addonName" is explicitly disabled based on the Enabled field of KubernetesAddon.
// If the value of Enabled is nil, we return false (not explicitly disabled)
func (k *KubernetesConfig) IsAddonDisabled(addonName string) bool {
	kubeAddon := k.GetAddonByName(addonName)
	return kubeAddon.IsDisabled()
}

// IsAADPodIdentityEnabled checks if the AAD pod identity addon is enabled
func (k *KubernetesConfig) IsAADPodIdentityEnabled() bool {
	return k.IsAddonEnabled(common.AADPodIdentityAddonName)
}

// IsContainerMonitoringAddonEnabled checks if the container monitoring addon is enabled
func (k *KubernetesConfig) IsContainerMonitoringAddonEnabled() bool {
	return k.IsAddonEnabled(common.ContainerMonitoringAddonName)
}

// IsClusterAutoscalerEnabled checks if the cluster autoscaler addon is enabled
func (k *KubernetesConfig) IsClusterAutoscalerEnabled() bool {
	return k.IsAddonEnabled(common.ClusterAutoscalerAddonName)
}

// IsAzurePolicyEnabled checks if the azure policy addon is enabled
func (k *KubernetesConfig) IsAzurePolicyEnabled() bool {
	return k.IsAddonEnabled(common.AzurePolicyAddonName)
}

// IsAppGWIngressEnabled checks if the appgw ingress addon is enabled
func (k *KubernetesConfig) IsAppGWIngressEnabled() bool {
	return k.IsAddonEnabled(common.AppGwIngressAddonName)
}

// IsIPMasqAgentEnabled checks if the ip-masq-agent addon is enabled
func (k *KubernetesConfig) IsIPMasqAgentEnabled() bool {
	return k.IsAddonEnabled(common.IPMASQAgentAddonName)
}

// IsIPMasqAgentDisabled checks if the ip-masq-agent addon is disabled
func (k *KubernetesConfig) IsIPMasqAgentDisabled() bool {
	return k.IsAddonDisabled(common.IPMASQAgentAddonName)
}

// GetComponentByName returns the KubernetesComponent object with name `componentName`
func (k *KubernetesConfig) GetComponentByName(componentName string) KubernetesComponent {
	var component KubernetesComponent
	for _, c := range k.Components {
		if c.Name == componentName {
			component = c
			break
		}
	}
	return component
}

// GetComponentData retrieves the raw data specified as input for a component with name "componentName".
func (k *KubernetesConfig) GetComponentData(componentName string) string {
	component := k.GetComponentByName(componentName)
	return component.Data
}

// IsComponentEnabled checks whether a component with name "componentName" is enabled or not based on the Enabled field of KubernetesComponent.
// If the value of Enabled is nil, the "defaultValue" is returned.
func (k *KubernetesConfig) IsComponentEnabled(componentName string) (KubernetesComponent, bool) {
	component := k.GetComponentByName(componentName)
	return component, component.IsEnabled()
}

// IsRBACEnabled checks if RBAC is enabled
func (k *KubernetesConfig) IsRBACEnabled() bool {
	if k.EnableRbac != nil {
		return to.Bool(k.EnableRbac)
	}
	return false
}

// UserAssignedIDEnabled checks if the user assigned ID is enabled or not.
func (k *KubernetesConfig) UserAssignedIDEnabled() bool {
	return k.UseManagedIdentity && k.UserAssignedID != ""
}

// SystemAssignedIDEnabled checks if system assigned IDs should be used.
func (k *KubernetesConfig) SystemAssignedIDEnabled() bool {
	return k.UseManagedIdentity && k.UserAssignedID == ""
}

func (k *KubernetesConfig) ShouldCreateNewUserAssignedIdentity() bool {
	return !(k.UserAssignedIDEnabled() && strings.Contains(k.UserAssignedID, "/"))
}

// GetOrderedKubeletConfigString returns an ordered string of key/val pairs
func (k *KubernetesConfig) GetOrderedKubeletConfigString() string {
	keys := []string{}
	for key := range k.KubeletConfig {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(fmt.Sprintf("%s=%s ", key, k.KubeletConfig[key]))
	}
	return buf.String()
}

// GetOrderedKubeletConfigStringForPowershell returns an ordered string of key/val pairs for Powershell script consumption
func (k *KubernetesConfig) GetOrderedKubeletConfigStringForPowershell() string {
	keys := []string{}
	for key := range k.KubeletConfig {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(fmt.Sprintf("\"%s=%s\", ", key, k.KubeletConfig[key]))
	}
	return strings.TrimSuffix(buf.String(), ", ")
}

// NeedsContainerd returns whether or not we need the containerd runtime configuration
func (k *KubernetesConfig) NeedsContainerd() bool {
	return k.ContainerRuntime == Containerd
}

// IsNSeriesSKU returns true if the agent pool contains an N-series (NVIDIA GPU) VM
func (a *AgentPoolProfile) IsNSeriesSKU() bool {
	return common.IsNvidiaEnabledSKU(a.VMSize)
}

// HasNSeriesSKU returns whether or not there is an N series SKU agent pool
func (p *Properties) HasNSeriesSKU() bool {
	for _, profile := range p.AgentPoolProfiles {
		if strings.Contains(profile.VMSize, "Standard_N") {
			return true
		}
	}
	return false
}

// HasDCSeriesSKU returns whether or not there is an DC series SKU agent pool
func (p *Properties) HasDCSeriesSKU() bool {
	for _, profile := range p.AgentPoolProfiles {
		if strings.Contains(profile.VMSize, "Standard_DC") {
			return true
		}
	}
	return false
}

// IsNVIDIADevicePluginEnabled checks if the NVIDIA Device Plugin addon is enabled
// It is enabled by default if agents contain a GPU and Kubernetes version is >= 1.10.0
func (p *Properties) IsNVIDIADevicePluginEnabled() bool {
	if p.OrchestratorProfile == nil || p.OrchestratorProfile.KubernetesConfig == nil {
		return false
	}
	return p.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.NVIDIADevicePluginAddonName)
}

// IsCustomCloudProfile returns true if user has provided a custom cloud profile
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

// GetKubernetesVersion returns the cluster Kubernetes version, with the Azure Stack suffix if Azure Stack Cloud.
func (p *Properties) GetKubernetesVersion() string {
	if p.IsAzureStackCloud() {
		return p.OrchestratorProfile.OrchestratorVersion + AzureStackSuffix
	}
	return p.OrchestratorProfile.OrchestratorVersion
}

// GetKubernetesHyperkubeSpec returns the string to use for the Kubernetes hyperkube image.
func (p *Properties) GetKubernetesHyperkubeSpec() string {
	var kubernetesHyperkubeSpec string
	k8sComponents := GetK8sComponentsByVersionMap(p.OrchestratorProfile.KubernetesConfig)[p.OrchestratorProfile.OrchestratorVersion]
	kubernetesHyperkubeSpec = p.OrchestratorProfile.KubernetesConfig.KubernetesImageBase + k8sComponents["hyperkube"]
	if p.IsAzureStackCloud() {
		kubernetesHyperkubeSpec = kubernetesHyperkubeSpec + AzureStackSuffix
	}
	if p.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage != "" {
		kubernetesHyperkubeSpec = p.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage
	}
	return kubernetesHyperkubeSpec
}

// IsAzureStackCloud return true if the cloud is AzureStack
func (p *Properties) IsAzureStackCloud() bool {
	// For backward compatibility, treat nil Environment and empty Environment name as AzureStackCloud as well
	return p.IsCustomCloudProfile() && (p.CustomCloudProfile.Environment == nil || p.CustomCloudProfile.Environment.Name == "" || strings.EqualFold(p.CustomCloudProfile.Environment.Name, "AzureStackCloud"))
}

// GetCustomEnvironmentJSON return the JSON format string for custom environment
func (p *Properties) GetCustomEnvironmentJSON(escape bool) (string, error) {
	var environmentJSON string
	if p.IsCustomCloudProfile() {
		bytes, err := json.Marshal(p.CustomCloudProfile.Environment)
		if err != nil {
			return "", fmt.Errorf("Could not serialize Environment object - %s", err.Error())
		}
		environmentJSON = string(bytes)
		if escape {
			environmentJSON = strings.Replace(environmentJSON, "\"", "\\\"", -1)
		}
	}
	return environmentJSON, nil
}

// GetCustomCloudName returns name of environment if customCloudProfile is provided, returns empty string if customCloudProfile is empty.
// Because customCloudProfile is empty for deployment is AzurePublicCloud, AzureChinaCloud,AzureGermanCloud,AzureUSGovernmentCloud,
// the return value will be empty string for those clouds
func (p *Properties) GetCustomCloudName() string {
	var cloudProfileName string
	if p.IsCustomCloudProfile() {
		cloudProfileName = p.CustomCloudProfile.Environment.Name
	}
	return cloudProfileName
}

// GetLocations returns all supported regions.
// If AzureStackCloud, GetLocations provides the location of container service
// If AzurePublicCloud, AzureChinaCloud,AzureGermanCloud or AzureUSGovernmentCloud, GetLocations provides all azure regions in prod.
func (cs *ContainerService) GetLocations() []string {
	var allLocations []string
	if cs.Properties.IsCustomCloudProfile() {
		allLocations = []string{cs.Location}
	} else {
		allLocations = helpers.GetAzureLocations()
	}
	return allLocations
}

// GetCustomCloudAuthenticationMethod returns authentication method which k8s azure cloud provider will use
// For AzurePublicCloud,AzureChinaCloud,azureGermanCloud,AzureUSGovernmentCloud, it will be always be client_secret
// For AzureStackCloud, if it is specified in configuration, the value will be used, if not ,the default value is client_secret.
func (p *Properties) GetCustomCloudAuthenticationMethod() string {
	if p.IsCustomCloudProfile() {
		return p.CustomCloudProfile.AuthenticationMethod
	}
	return ClientSecretAuthMethod
}

// GetCustomCloudIdentitySystem returns identity system method for azure stack.
// For AzurePublicCloud,AzureChinaCloud,azureGermanCloud,AzureUSGovernmentCloud, it will be always be AzureAD
// For AzureStackCloud, if it is specified in configuration, the value will be used, if not ,the default value is AzureAD.
func (p *Properties) GetCustomCloudIdentitySystem() string {
	if p.IsCustomCloudProfile() {
		return p.CustomCloudProfile.IdentitySystem
	}
	return AzureADIdentitySystem
}

// IsNvidiaDevicePluginCapable determines if the cluster definition is compatible with the nvidia-device-plugin daemonset
func (p *Properties) IsNvidiaDevicePluginCapable() bool {
	return p.HasNSeriesSKU()
}

// IsAzureCNIDualStack determines if azure cni dual stack is enabled
func (p *Properties) IsAzureCNIDualStack() bool {
	o := p.OrchestratorProfile
	f := p.FeatureFlags
	return o.IsAzureCNI() && f.IsFeatureEnabled("EnableIPv6DualStack")
}

// RequireRouteTable returns true if this deployment requires routing table
func (p *Properties) RequireRouteTable() bool {
	o := p.OrchestratorProfile
	f := p.FeatureFlags
	switch o.OrchestratorType {
	case Kubernetes:
		if o.IsAzureCNI() && !f.IsFeatureEnabled("EnableIPv6DualStack") ||
			NetworkPolicyCilium == o.KubernetesConfig.NetworkPolicy ||
			"flannel" == o.KubernetesConfig.NetworkPlugin ||
			NetworkPluginAntrea == o.KubernetesConfig.NetworkPlugin {
			return false
		}
		return true
	default:
		return false
	}
}

// SetCloudProviderRateLimitDefaults sets default cloudprovider rate limiter config
func (p *Properties) SetCloudProviderRateLimitDefaults() {
	if p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucket == 0 {
		var agentPoolProfilesCount = len(p.AgentPoolProfiles)
		if agentPoolProfilesCount == 0 {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucket = DefaultKubernetesCloudProviderRateLimitBucket
		} else {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucket = agentPoolProfilesCount * common.MaxAgentCount
		}
		if p.IsAzureStackCloud() {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucket = DefaultAzureStackKubernetesCloudProviderRateLimitBucket
		}
	}
	if p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPS == 0 {
		if (DefaultKubernetesCloudProviderRateLimitQPS / float64(p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucket)) < common.MinCloudProviderQPSToBucketFactor {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPS = float64(p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucket) * common.MinCloudProviderQPSToBucketFactor
		} else {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPS = DefaultKubernetesCloudProviderRateLimitQPS
		}
		if p.IsAzureStackCloud() {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPS = DefaultAzureStackKubernetesCloudProviderRateLimitQPS
		}
	}
	if p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucketWrite == 0 {
		var agentPoolProfilesCount = len(p.AgentPoolProfiles)
		if agentPoolProfilesCount == 0 {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucketWrite = DefaultKubernetesCloudProviderRateLimitBucketWrite
		} else {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucketWrite = agentPoolProfilesCount * common.MaxAgentCount
		}
		if p.IsAzureStackCloud() {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucketWrite = DefaultAzureStackKubernetesCloudProviderRateLimitBucketWrite
		}
	}
	if p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPSWrite == 0 {
		if (DefaultKubernetesCloudProviderRateLimitQPSWrite / float64(p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucketWrite)) < common.MinCloudProviderQPSToBucketFactor {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPSWrite = float64(p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitBucketWrite) * common.MinCloudProviderQPSToBucketFactor
		} else {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPSWrite = DefaultKubernetesCloudProviderRateLimitQPSWrite
		}
		if p.IsAzureStackCloud() {
			p.OrchestratorProfile.KubernetesConfig.CloudProviderRateLimitQPSWrite = DefaultAzureStackKubernetesCloudProviderRateLimitQPSWrite
		}
	}
}

// IsReschedulerEnabled checks if the rescheduler addon is enabled
func (k *KubernetesConfig) IsReschedulerEnabled() bool {
	return k.IsAddonEnabled(common.ReschedulerAddonName)
}

// PrivateJumpboxProvision checks if a private cluster has jumpbox auto-provisioning
func (k *KubernetesConfig) PrivateJumpboxProvision() bool {
	if k != nil && k.PrivateCluster != nil && *k.PrivateCluster.Enabled && k.PrivateCluster.JumpboxProfile != nil {
		return true
	}
	return false
}

// RequiresDocker returns if the kubernetes settings require docker binary to be installed.
func (k *KubernetesConfig) RequiresDocker() bool {
	if k == nil {
		return false
	}

	runtime := strings.ToLower(k.ContainerRuntime)
	return runtime == Docker || runtime == ""
}

// SetCloudProviderBackoffDefaults sets default cloudprovider backoff config
func (p *Properties) SetCloudProviderBackoffDefaults() {
	k := p.OrchestratorProfile.KubernetesConfig
	if k.CloudProviderBackoffDuration == 0 {
		k.CloudProviderBackoffDuration = DefaultKubernetesCloudProviderBackoffDuration
		if p.IsAzureStackCloud() {
			k.CloudProviderBackoffDuration = DefaultAzureStackKubernetesCloudProviderBackoffDuration
		}
	}
	if k.CloudProviderBackoffRetries == 0 {
		k.CloudProviderBackoffRetries = DefaultKubernetesCloudProviderBackoffRetries
		if p.IsAzureStackCloud() {
			k.CloudProviderBackoffRetries = DefaultAzureStackKubernetesCloudProviderBackoffRetries
		}
	}
	if k.CloudProviderBackoffMode != CloudProviderBackoffModeV2 {
		if k.CloudProviderBackoffExponent == 0 {
			k.CloudProviderBackoffExponent = DefaultKubernetesCloudProviderBackoffExponent
			if p.IsAzureStackCloud() {
				k.CloudProviderBackoffExponent = DefaultAzureStackKubernetesCloudProviderBackoffExponent
			}
		}
		if k.CloudProviderBackoffJitter == 0 {
			k.CloudProviderBackoffJitter = DefaultKubernetesCloudProviderBackoffJitter
			if p.IsAzureStackCloud() {
				k.CloudProviderBackoffJitter = DefaultAzureStackKubernetesCloudProviderBackoffJitter
			}
		}
	}
}

// GetAzureCNIURLLinux returns the full URL to source Azure CNI binaries from
func (k *KubernetesConfig) GetAzureCNIURLLinux(cloudSpecConfig AzureEnvironmentSpecConfig) string {
	if k.AzureCNIURLLinux != "" {
		return k.AzureCNIURLLinux
	}
	return cloudSpecConfig.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL
}

// GetAzureCNIURLWindows returns the full URL to source Azure CNI binaries from
func (k *KubernetesConfig) GetAzureCNIURLWindows(cloudSpecConfig AzureEnvironmentSpecConfig) string {
	if k.AzureCNIURLWindows != "" {
		return k.AzureCNIURLWindows
	}
	return cloudSpecConfig.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL
}

// IsFeatureEnabled returns true if a feature flag is on for the provided feature
func (f *FeatureFlags) IsFeatureEnabled(feature string) bool {
	if f != nil {
		switch feature {
		case "CSERunInBackground":
			return f.EnableCSERunInBackground
		case "BlockOutboundInternet":
			return f.BlockOutboundInternet
		case "EnableIPv6DualStack":
			return f.EnableIPv6DualStack
		case "EnableTelemetry":
			return f.EnableTelemetry
		case "EnableIPv6Only":
			return f.EnableIPv6Only
		default:
			return false
		}
	}
	return false
}

// GetCloudSpecConfig returns the Kubernetes container images URL configurations based on the deploy target environment.
//for example: if the target is the public azure, then the default container image url should be k8s.gcr.io/...
//if the target is azure china, then the default container image should be mirror.azure.cn:5000/google_container/...
func (cs *ContainerService) GetCloudSpecConfig() AzureEnvironmentSpecConfig {
	targetEnv := helpers.GetTargetEnv(cs.Location, cs.Properties.GetCustomCloudName())
	return AzureCloudSpecEnvMap[targetEnv]
}

// IsAKSBillingEnabled checks if the AKS Billing Extension should be enabled for a cloud environment.
func (cs *ContainerService) IsAKSBillingEnabled() bool {
	cloudSpecConfig := cs.GetCloudSpecConfig()
	return cloudSpecConfig.CloudName == AzurePublicCloud || cloudSpecConfig.CloudName == AzureChinaCloud || cloudSpecConfig.CloudName == AzureUSGovernmentCloud
}

// GetAzureProdFQDN returns the formatted FQDN string for a given apimodel.
func (cs *ContainerService) GetAzureProdFQDN() string {
	return FormatProdFQDNByLocation(cs.Properties.MasterProfile.DNSPrefix, cs.Location, cs.Properties.GetCustomCloudName())
}

// ProvisionScriptParametersInput is the struct used to pass in Azure environment variables and secrets
// as either values or ARM template variables when generating provision script parameters.
type ProvisionScriptParametersInput struct {
	Location             string
	ResourceGroup        string
	TenantID             string
	SubscriptionID       string
	ClientID             string
	ClientSecret         string
	APIServerCertificate string
	KubeletPrivateKey    string
	ClusterKeyVaultName  string
}

// GetProvisionScriptParametersCommon returns the environment variables needed to run the Linux bootstrap scripts
// Ensure that the clientSecret parameter is surrounded by single quotes to protect against special characters
func (cs *ContainerService) GetProvisionScriptParametersCommon(input ProvisionScriptParametersInput) string {
	cloudSpecConfig := cs.GetCloudSpecConfig()
	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig
	parameters := map[string]string{
		"ADMINUSER":                            cs.Properties.LinuxProfile.AdminUsername,
		"ETCD_DOWNLOAD_URL":                    cloudSpecConfig.KubernetesSpecConfig.EtcdDownloadURLBase,
		"ETCD_VERSION":                         kubernetesConfig.EtcdVersion,
		"CONTAINERD_VERSION":                   kubernetesConfig.ContainerdVersion,
		"MOBY_VERSION":                         kubernetesConfig.MobyVersion,
		"TENANT_ID":                            input.TenantID,
		"KUBERNETES_VERSION":                   cs.Properties.GetKubernetesVersion(),
		"HYPERKUBE_URL":                        cs.Properties.GetKubernetesHyperkubeSpec(),
		"APISERVER_PUBLIC_KEY":                 input.APIServerCertificate,
		"SUBSCRIPTION_ID":                      input.SubscriptionID,
		"RESOURCE_GROUP":                       input.ResourceGroup,
		"LOCATION":                             input.Location,
		"VM_TYPE":                              cs.Properties.GetVMType(),
		"SUBNET":                               cs.Properties.GetSubnetName(),
		"NETWORK_SECURITY_GROUP":               cs.Properties.GetNSGName(),
		"VIRTUAL_NETWORK":                      cs.Properties.GetVirtualNetworkName(),
		"VIRTUAL_NETWORK_RESOURCE_GROUP":       cs.Properties.GetVNetResourceGroupName(),
		"ROUTE_TABLE":                          cs.Properties.GetRouteTableName(),
		"PRIMARY_AVAILABILITY_SET":             cs.Properties.GetPrimaryAvailabilitySetName(),
		"PRIMARY_SCALE_SET":                    cs.Properties.GetPrimaryScaleSetName(),
		"SERVICE_PRINCIPAL_CLIENT_ID":          input.ClientID,
		"SERVICE_PRINCIPAL_CLIENT_SECRET":      input.ClientSecret,
		"KUBELET_PRIVATE_KEY":                  input.KubeletPrivateKey,
		"NETWORK_PLUGIN":                       kubernetesConfig.NetworkPlugin,
		"NETWORK_POLICY":                       kubernetesConfig.NetworkPolicy,
		"VNET_CNI_PLUGINS_URL":                 kubernetesConfig.GetAzureCNIURLLinux(cloudSpecConfig),
		"CNI_PLUGINS_URL":                      cloudSpecConfig.KubernetesSpecConfig.CNIPluginsDownloadURL,
		"CLOUDPROVIDER_BACKOFF":                strconv.FormatBool(to.Bool(kubernetesConfig.CloudProviderBackoff)),
		"CLOUDPROVIDER_BACKOFF_MODE":           kubernetesConfig.CloudProviderBackoffMode,
		"CLOUDPROVIDER_BACKOFF_RETRIES":        strconv.Itoa(kubernetesConfig.CloudProviderBackoffRetries),
		"CLOUDPROVIDER_BACKOFF_EXPONENT":       strconv.FormatFloat(kubernetesConfig.CloudProviderBackoffExponent, 'f', -1, 64),
		"CLOUDPROVIDER_BACKOFF_DURATION":       strconv.Itoa(kubernetesConfig.CloudProviderBackoffDuration),
		"CLOUDPROVIDER_BACKOFF_JITTER":         strconv.FormatFloat(kubernetesConfig.CloudProviderBackoffJitter, 'f', -1, 64),
		"CLOUDPROVIDER_RATELIMIT":              strconv.FormatBool(to.Bool(kubernetesConfig.CloudProviderRateLimit)),
		"CLOUDPROVIDER_RATELIMIT_QPS":          strconv.FormatFloat(kubernetesConfig.CloudProviderRateLimitQPS, 'f', -1, 64),
		"CLOUDPROVIDER_RATELIMIT_QPS_WRITE":    strconv.FormatFloat(kubernetesConfig.CloudProviderRateLimitQPSWrite, 'f', -1, 64),
		"CLOUDPROVIDER_RATELIMIT_BUCKET":       strconv.Itoa(kubernetesConfig.CloudProviderRateLimitBucket),
		"CLOUDPROVIDER_RATELIMIT_BUCKET_WRITE": strconv.Itoa(kubernetesConfig.CloudProviderRateLimitBucketWrite),
		"LOAD_BALANCER_DISABLE_OUTBOUND_SNAT":  strconv.FormatBool(to.Bool(kubernetesConfig.CloudProviderDisableOutboundSNAT)),
		"USE_MANAGED_IDENTITY_EXTENSION":       strconv.FormatBool(kubernetesConfig.UseManagedIdentity),
		"USE_INSTANCE_METADATA":                strconv.FormatBool(to.Bool(kubernetesConfig.UseInstanceMetadata)),
		"LOAD_BALANCER_SKU":                    kubernetesConfig.LoadBalancerSku,
		"EXCLUDE_MASTER_FROM_STANDARD_LB":      strconv.FormatBool(to.Bool(kubernetesConfig.ExcludeMasterFromStandardLB)),
		"MAXIMUM_LOADBALANCER_RULE_COUNT":      strconv.Itoa(kubernetesConfig.MaximumLoadBalancerRuleCount),
		"CONTAINER_RUNTIME":                    kubernetesConfig.ContainerRuntime,
		"CONTAINERD_DOWNLOAD_URL_BASE":         cloudSpecConfig.KubernetesSpecConfig.ContainerdDownloadURLBase,
		"KMS_PROVIDER_VAULT_NAME":              input.ClusterKeyVaultName,
		"IS_HOSTED_MASTER":                     strconv.FormatBool(cs.Properties.IsHostedMasterProfile()),
		"IS_IPV6_DUALSTACK_FEATURE_ENABLED":    strconv.FormatBool(cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack")),
		"IS_IPV6_ENABLED":                      strconv.FormatBool(cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6Only") || cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack")),
		"AUTHENTICATION_METHOD":                cs.Properties.GetCustomCloudAuthenticationMethod(),
		"IDENTITY_SYSTEM":                      cs.Properties.GetCustomCloudIdentitySystem(),
		"NETWORK_API_VERSION":                  APIVersionNetwork,
		"NETWORK_MODE":                         kubernetesConfig.NetworkMode,
		"KUBE_BINARY_URL":                      kubernetesConfig.CustomKubeBinaryURL,
	}

	if cs.Properties.IsHostedMasterProfile() && cs.Properties.HostedMasterProfile.FQDN != "" {
		parameters["API_SERVER_NAME"] = cs.Properties.HostedMasterProfile.FQDN
	}

	keys := make([]string, 0)
	for k := range parameters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var provisionScriptParametersCommon strings.Builder
	for _, name := range keys {
		provisionScriptParametersCommon.WriteString(name)
		provisionScriptParametersCommon.WriteString("=")
		provisionScriptParametersCommon.WriteString(parameters[name])
		provisionScriptParametersCommon.WriteString(" ")
	}

	return provisionScriptParametersCommon.String()
}

// FormatAzureProdFQDNByLocation constructs an Azure prod fqdn
func FormatAzureProdFQDNByLocation(fqdnPrefix string, location string) string {
	targetEnv := helpers.GetCloudTargetEnv(location)
	FQDNFormat := AzureCloudSpecEnvMap[targetEnv].EndpointConfig.ResourceManagerVMDNSSuffix
	return fmt.Sprintf("%s.%s."+FQDNFormat, fqdnPrefix, location)
}

// FormatProdFQDNByLocation constructs an Azure prod fqdn with custom cloud profile
// CustomCloudName is name of environment if customCloudProfile is provided, it will be empty string if customCloudProfile is empty.
// Because customCloudProfile is empty for deployment for AzurePublicCloud, AzureChinaCloud,AzureGermanCloud,AzureUSGovernmentCloud,
// The customCloudName value will be empty string for those clouds
func FormatProdFQDNByLocation(fqdnPrefix string, location string, cloudName string) string {
	targetEnv := helpers.GetTargetEnv(location, cloudName)
	FQDNFormat := AzureCloudSpecEnvMap[targetEnv].EndpointConfig.ResourceManagerVMDNSSuffix
	return fmt.Sprintf("%s.%s."+FQDNFormat, fqdnPrefix, location)
}
