package v20180331

import "encoding/json"

// The validate tag is used for validation
// Reference to gopkg.in/go-playground/validator.v9

// ResourcePurchasePlan defines resource plan as required by ARM
// for billing purposes.
type ResourcePurchasePlan struct {
	Name          string `json:"name,omitempty"`
	Product       string `json:"product,omitempty"`
	PromotionCode string `json:"promotionCode,omitempty"`
	Publisher     string `json:"publisher,omitempty"`
}

// ManagedCluster complies with the ARM model of
// resource definition in a JSON template.
type ManagedCluster struct {
	ID       string                `json:"id,omitempty"`
	Location string                `json:"location,omitempty" validate:"required"`
	Name     string                `json:"name,omitempty"`
	Plan     *ResourcePurchasePlan `json:"plan,omitempty"`
	Tags     map[string]string     `json:"tags,omitempty"`
	Type     string                `json:"type,omitempty"`

	Properties *Properties `json:"properties"`
}

// Properties represents the AKS cluster definition
type Properties struct {
	ProvisioningState       ProvisioningState        `json:"provisioningState,omitempty"`
	KubernetesVersion       string                   `json:"kubernetesVersion"`
	DNSPrefix               string                   `json:"dnsPrefix,omitempty"`
	FQDN                    string                   `json:"fqdn,omitempty"`
	AgentPoolProfiles       []*AgentPoolProfile      `json:"agentPoolProfiles,omitempty" validate:"dive,required"`
	LinuxProfile            *LinuxProfile            `json:"linuxProfile,omitempty"`
	WindowsProfile          *WindowsProfile          `json:"windowsProfile,omitempty"`
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
	AccessProfiles          map[string]AccessProfile `json:"accessProfiles,omitempty"`
	AddonProfiles           map[string]AddonProfile  `json:"addonProfiles,omitempty"`
	NodeResourceGroup       string                   `json:"nodeResourceGroup,omitempty"`
	EnableRBAC              *bool                    `json:"enableRBAC,omitempty"`
	NetworkProfile          *NetworkProfile          `json:"networkProfile,omitempty"`
	AADProfile              *AADProfile              `json:"aadProfile,omitempty"`
}

// NetworkProfile represents network related definitions
type NetworkProfile struct {
	NetworkPlugin    NetworkPlugin `json:"networkPlugin,omitempty"`
	NetworkPolicy    NetworkPolicy `json:"networkPolicy,omitempty"`
	PodCidr          string        `json:"podCidr,omitempty"`
	ServiceCidr      string        `json:"serviceCidr,omitempty"`
	DNSServiceIP     string        `json:"dnsServiceIP,omitempty"`
	DockerBridgeCidr string        `json:"dockerBridgeCidr,omitempty"`
}

// NetworkPlugin represnets types of network plugin
type NetworkPlugin string

const (
	// Azure represents Azure CNI network plugin
	Azure NetworkPlugin = "azure"
	// Kubenet represents Kubenet network plugin
	Kubenet NetworkPlugin = "kubenet"
)

// NetworkPolicy represnets types of network policy
type NetworkPolicy string

const (
	// NetworkPolicyCalico represents Calico network policy
	NetworkPolicyCalico NetworkPolicy = "calico"
)

// AddonProfile represents an addon for managed cluster
type AddonProfile struct {
	Enabled bool              `json:"enabled"`
	Config  map[string]string `json:"config"`
}

// ManagedClusterAccessProfile represents the access profile definition for managed cluster
// The Id captures the Role Name e.g. clusterAdmin, clusterUser
type ManagedClusterAccessProfile struct {
	ID       string `json:"id,omitempty"`
	Location string `json:"location,omitempty" validate:"required"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`

	Properties *AccessProfile `json:"properties"`
}

// ServicePrincipalProfile contains the client and secret used by the cluster for Azure Resource CRUD
// The 'Secret' parameter could be either a plain text, or referenced to a secret in a keyvault.
// In the latter case, the format of the parameter's value should be
// "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<NAME>[/<VERSION>]"
// where:
//    <SUB_ID> is the subscription ID of the keyvault
//    <RG_NAME> is the resource group of the keyvault
//    <KV_NAME> is the name of the keyvault
//    <NAME> is the name of the secret.
//    <VERSION> (optional) is the version of the secret (default: the latest version)
type ServicePrincipalProfile struct {
	ClientID string `json:"clientId,omitempty" validate:"required"`
	Secret   string `json:"secret,omitempty" conform:"redact"`
}

// LinuxProfile represents the Linux configuration passed to the cluster
type LinuxProfile struct {
	AdminUsername string `json:"adminUsername" validate:"required"`

	SSH struct {
		PublicKeys []PublicKey `json:"publicKeys" validate:"required,len=1"`
	} `json:"ssh" validate:"required"`
}

// PublicKey represents an SSH key for LinuxProfile
type PublicKey struct {
	KeyData string `json:"keyData"`
}

// WindowsProfile represents the Windows configuration passed to the cluster
type WindowsProfile struct {
	AdminUsername string `json:"adminUsername,omitempty" validate:"required"`
	AdminPassword string `json:"adminPassword,omitempty" conform:"redact"`
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
	// Upgrading means an existing resource is being upgraded
	Upgrading ProvisioningState = "Upgrading"
)

// PoolUpgradeProfile contains pool properties:
//  - kubernetes version
//  - pool name (for agent pool)
//  - OS type of the VMs in the pool
//  - list of applicable upgrades
type PoolUpgradeProfile struct {
	KubernetesVersion string   `json:"kubernetesVersion"`
	Name              string   `json:"name,omitempty"`
	OSType            string   `json:"osType,omitempty"`
	Upgrades          []string `json:"upgrades,omitempty"`
}

// UpgradeProfileProperties contains properties of UpgradeProfile
type UpgradeProfileProperties struct {
	ControlPlaneProfile *PoolUpgradeProfile   `json:"controlPlaneProfile"`
	AgentPoolProfiles   []*PoolUpgradeProfile `json:"agentPoolProfiles"`
}

// UpgradeProfile contains controlPlane and agent pools upgrade profiles
type UpgradeProfile struct {
	ID         string                   `json:"id,omitempty"`
	Name       string                   `json:"name,omitempty"`
	Type       string                   `json:"type,omitempty"`
	Properties UpgradeProfileProperties `json:"properties"`
}

// AgentPoolProfile represents configuration of VMs running agent
// daemons that register with the master and offer resources to
// host applications in containers.
type AgentPoolProfile struct {
	Name           string `json:"name" validate:"required"`
	Count          int    `json:"count" validate:"required,min=1,max=100"`
	VMSize         string `json:"vmSize" validate:"required"`
	OSDiskSizeGB   int    `json:"osDiskSizeGB,omitempty" validate:"min=0,max=1023"`
	StorageProfile string `json:"storageProfile" validate:"eq=ManagedDisks|len=0"`
	VnetSubnetID   string `json:"vnetSubnetID,omitempty"`
	MaxPods        *int   `json:"maxPods,omitempty"`

	// OSType is the operating system type for agents
	// Set as nullable to support backward compat because
	// this property was added later.
	// If the value is null or not set, it defaulted to Linux.
	OSType OSType `json:"osType,omitempty"`

	// subnet is internal
	subnet string
}

// AccessProfile represents role name and kubeconfig
type AccessProfile struct {
	KubeConfig string `json:"kubeConfig"`
}

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
}

// UnmarshalJSON unmarshal json using the default behavior
// And do fields manipulation, such as populating default value
func (a *AgentPoolProfile) UnmarshalJSON(b []byte) error {
	// Need to have a alias type to avoid circular unmarshal
	type aliasAgentPoolProfile AgentPoolProfile
	aa := aliasAgentPoolProfile{}
	if e := json.Unmarshal(b, &aa); e != nil {
		return e
	}
	*a = AgentPoolProfile(aa)
	if a.Count == 0 {
		// if AgentPoolProfile.Count is missing or 0, set it to default 1
		a.Count = 1
	}

	if a.StorageProfile == "" {
		// if StorageProfile is missing, set to default ManagedDisks
		a.StorageProfile = ManagedDisks
	}

	if string(a.OSType) == "" {
		// OSType is the operating system type for agents
		// Set as nullable to support backward compat because
		// this property was added later.
		// If the value is null or not set, it defaulted to Linux.
		a.OSType = Linux
	}

	// OSDiskSizeGB is an override value. vm sizes have default OS disk sizes.
	// If it is not set. The user should get the default for the vm size
	return nil
}

// OSType represents OS types of agents
type OSType string

// HasWindows returns true if the cluster contains windows
func (a *Properties) HasWindows() bool {
	for _, agentPoolProfile := range a.AgentPoolProfiles {
		if agentPoolProfile.OSType == Windows {
			return true
		}
	}
	return false
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

// GetSubnet returns the read-only subnet for the agent pool
func (a *AgentPoolProfile) GetSubnet() string {
	return a.subnet
}

// SetSubnet sets the read-only subnet for the agent pool
func (a *AgentPoolProfile) SetSubnet(subnet string) {
	a.subnet = subnet
}

// IsManagedDisks returns true if the customer specified managed disks
func (a *AgentPoolProfile) IsManagedDisks() bool {
	return a.StorageProfile == ManagedDisks
}
