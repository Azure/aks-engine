package v20170131

import (
	"encoding/json"
	neturl "net/url"
	"strings"

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
	OrchestratorProfile     *OrchestratorProfile     `json:"orchestratorProfile,omitempty"`
	MasterProfile           *MasterProfile           `json:"masterProfile,omitempty"`
	AgentPoolProfiles       []*AgentPoolProfile      `json:"agentPoolProfiles,omitempty"`
	LinuxProfile            *LinuxProfile            `json:"linuxProfile,omitempty"`
	WindowsProfile          *WindowsProfile          `json:"windowsProfile,omitempty"`
	DiagnosticsProfile      *DiagnosticsProfile      `json:"diagnosticsProfile,omitempty"`
	JumpboxProfile          *JumpboxProfile          `json:"jumpboxProfile,omitempty"`
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
	CustomProfile           *CustomProfile           `json:"customProfile,omitempty"`
}

// ServicePrincipalProfile contains the client and secret used by the cluster for Azure Resource CRUD
type ServicePrincipalProfile struct {
	ClientID string `json:"clientId,omitempty"`
	Secret   string `json:"secret,omitempty"`
	ObjectID string `json:"objectId,omitempty"`
}

// CustomProfile specifies custom properties that are used for
// cluster instantiation.  Should not be used by most users.
type CustomProfile struct {
	Orchestrator string `json:"orchestrator,omitempty"`
}

// LinuxProfile represents the Linux configuration passed to the cluster
type LinuxProfile struct {
	AdminUsername string `json:"adminUsername"`

	SSH struct {
		PublicKeys []PublicKey `json:"publicKeys"`
	} `json:"ssh"`
}

// PublicKey represents an SSH key for LinuxProfile
type PublicKey struct {
	KeyData string `json:"keyData"`
}

// WindowsProfile represents the Windows configuration passed to the cluster
type WindowsProfile struct {
	AdminUsername string `json:"adminUsername,omitempty"`
	AdminPassword string `json:"adminPassword,omitempty"`
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
	OrchestratorType string `json:"orchestratorType"`
}

// MasterProfile represents the definition of master cluster
type MasterProfile struct {
	Count     int    `json:"count"`
	DNSPrefix string `json:"dnsPrefix"`

	// Master LB public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`

	// subnet is internal
	subnet string
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
	return nil
}

// AgentPoolProfile represents configuration of VMs running agent
// daemons that register with the master and offer resources to
// host applications in containers.
type AgentPoolProfile struct {
	Name      string `json:"name"`
	Count     int    `json:"count"`
	VMSize    string `json:"vmSize"`
	DNSPrefix string `json:"dnsPrefix"`
	FQDN      string `json:"fqdn"`

	// OSType is the operating system type for agents
	// Set as nullable to support backward compat because
	// this property was added later.
	// If the value is null or not set, it defaulted to Linux.
	OSType OSType `json:"osType,omitempty"`

	// subnet is internal
	subnet string
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
	return nil
}

// JumpboxProfile dscribes properties of the jumpbox setup
// in the ACS container cluster.
type JumpboxProfile struct {
	OSType    OSType `json:"osType,omitempty"`
	DNSPrefix string `json:"dnsPrefix"`

	// Jumpbox public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`
}

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

// GetSubnet returns the read-only subnet for the master
func (m *MasterProfile) GetSubnet() string {
	return m.subnet
}

// SetSubnet sets the read-only subnet for the master
func (m *MasterProfile) SetSubnet(subnet string) {
	m.subnet = subnet
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
