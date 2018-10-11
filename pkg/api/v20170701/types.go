package v20170701

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

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

// ContainerService complies with the ARM model of
// resource definition in a JSON template.
type ContainerService struct {
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
	OrchestratorProfile     *OrchestratorProfile     `json:"orchestratorProfile,omitempty" validate:"required"`
	MasterProfile           *MasterProfile           `json:"masterProfile,omitempty" validate:"required"`
	AgentPoolProfiles       []*AgentPoolProfile      `json:"agentPoolProfiles,omitempty" validate:"dive,required"`
	LinuxProfile            *LinuxProfile            `json:"linuxProfile,omitempty" validate:"required"`
	WindowsProfile          *WindowsProfile          `json:"windowsProfile,omitempty"`
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
	CustomProfile           *CustomProfile           `json:"customProfile,omitempty"`
}

// ServicePrincipalProfile contains the client and secret used by the cluster for Azure Resource CRUD
// The 'Secret' parameter should be a secret in plain text.
// The 'KeyvaultSecretRef' parameter is a reference to a secret in a keyvault.
// The format of the parameter's value should be
// "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<NAME>[/<VERSION>]"
// where:
//    <SUB_ID> is the subscription ID of the keyvault
//    <RG_NAME> is the resource group of the keyvault
//    <KV_NAME> is the name of the keyvault
//    <NAME> is the name of the secret.
//    <VERSION> (optional) is the version of the secret (default: the latest version)
type ServicePrincipalProfile struct {
	ClientID          string             `json:"clientId,omitempty" validate:"required"`
	Secret            string             `json:"secret,omitempty"`
	ObjectID          string             `json:"objectId,omitempty"`
	KeyvaultSecretRef *KeyvaultSecretRef `json:"keyvaultSecretRef,omitempty"`
}

// KeyvaultSecretRef is a reference to a secret in a keyvault.
type KeyvaultSecretRef struct {
	VaultID       string `json:"vaultID" validate:"required"`
	SecretName    string `json:"secretName" validate:"required"`
	SecretVersion string `json:"version,omitempty"`
}

// CustomProfile specifies custom properties that are used for
// cluster instantiation.  Should not be used by most users.
type CustomProfile struct {
	Orchestrator string `json:"orchestrator,omitempty"`
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
	AdminPassword string `json:"adminPassword,omitempty" validate:"required"`
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
	OrchestratorType    string `json:"orchestratorType" validate:"required"`
	OrchestratorVersion string `json:"orchestratorVersion,omitempty"`
}

// MasterProfile represents the definition of master cluster
type MasterProfile struct {
	Count                    int    `json:"count" validate:"required,eq=1|eq=3|eq=5"`
	DNSPrefix                string `json:"dnsPrefix" validate:"required"`
	VMSize                   string `json:"vmSize" validate:"required"`
	OSDiskSizeGB             int    `json:"osDiskSizeGB,omitempty" validate:"min=0,max=1023"`
	VnetSubnetID             string `json:"vnetSubnetID,omitempty"`
	FirstConsecutiveStaticIP string `json:"firstConsecutiveStaticIP,omitempty"`
	StorageProfile           string `json:"storageProfile,omitempty" validate:"eq=StorageAccount|eq=ManagedDisks|len=0"`

	// subnet is internal
	subnet string
	// Master LB public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`
}

// UnmarshalJSON unmarshal json using the default behavior
// And do fields manipulation, such as populating default value
func (m *MasterProfile) UnmarshalJSON(b []byte) error {
	// Need to have a alias type to avoid circular unmarshal
	type aliasMasterProfile MasterProfile
	mm := aliasMasterProfile{}
	if e := json.Unmarshal(b, &mm); e != nil {
		return e
	}
	*m = MasterProfile(mm)
	if m.Count == 0 {
		// if MasterProfile.Count is missing or 0, set to default 1
		m.Count = 1
	}

	if m.FirstConsecutiveStaticIP == "" {
		// if FirstConsecutiveStaticIP is missing, set to default 10.240.255.5
		m.FirstConsecutiveStaticIP = "10.240.255.5"
	}

	// OSDiskSizeGB is an override value. vm sizes have default OS disk sizes.
	// If it is not set. The user should get the default for the vm size
	return nil
}

// AgentPoolProfile represents configuration of VMs running agent
// daemons that register with the master and offer resources to
// host applications in containers.
type AgentPoolProfile struct {
	Name           string `json:"name" validate:"required"`
	Count          int    `json:"count" validate:"required,min=1,max=100"`
	VMSize         string `json:"vmSize" validate:"required"`
	OSDiskSizeGB   int    `json:"osDiskSizeGB,omitempty" validate:"min=0,max=1023"`
	DNSPrefix      string `json:"dnsPrefix"`
	FQDN           string `json:"fqdn"`
	Ports          []int  `json:"ports,omitempty" validate:"dive,min=1,max=65535"`
	StorageProfile string `json:"storageProfile" validate:"eq=StorageAccount|eq=ManagedDisks|len=0"`
	VnetSubnetID   string `json:"vnetSubnetID,omitempty"`
	// OSType is the operating system type for agents
	// Set as nullable to support backward compat because
	// this property was added later.
	// If the value is null or not set, it defaulted to Linux.
	OSType OSType `json:"osType,omitempty"`

	// subnet is internal
	subnet string
}

// PoolUpgradeProfile contains pool properties:
//  - orchestrator type and version
//  - pool name (for agent pool)
//  - OS type of the VMs in the pool
//  - list of applicable upgrades
type PoolUpgradeProfile struct {
	OrchestratorProfile
	Name     string                 `json:"name,omitempty"`
	OSType   string                 `json:"osType,omitempty"`
	Upgrades []*OrchestratorProfile `json:"upgrades,omitempty"`
}

// UpgradeProfileProperties contains properties of UpgradeProfile
type UpgradeProfileProperties struct {
	MasterPoolProfile *PoolUpgradeProfile   `json:"masterPoolProfile"`
	AgentPoolProfiles []*PoolUpgradeProfile `json:"agentPoolProfiles"`
}

// UpgradeProfile contains master and agent pools upgrade profiles
type UpgradeProfile struct {
	ID         string                   `json:"id,omitempty"`
	Name       string                   `json:"name,omitempty"`
	Type       string                   `json:"type,omitempty"`
	Properties UpgradeProfileProperties `json:"properties"`
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
	case strings.EqualFold(orchestratorType, Kubernetes):
		o.OrchestratorType = Kubernetes
	default:
		return errors.Errorf("OrchestratorType has unknown orchestrator: %s", orchestratorType)
	}
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

// IsManagedDisks returns true if the customer specified managed disks
func (a *AgentPoolProfile) IsManagedDisks() bool {
	return a.StorageProfile == ManagedDisks
}

// IsStorageAccount returns true if the customer specified storage account
func (a *AgentPoolProfile) IsStorageAccount() bool {
	return a.StorageProfile == StorageAccount
}

// GetSubnet returns the read-only subnet for the agent pool
func (a *AgentPoolProfile) GetSubnet() string {
	return a.subnet
}

// SetSubnet sets the read-only subnet for the agent pool
func (a *AgentPoolProfile) SetSubnet(subnet string) {
	a.subnet = subnet
}
