package v20170701

import (
	"net"
	"regexp"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate        *validator.Validate
	keyvaultIDRegex *regexp.Regexp
)

func init() {
	validate = validator.New()
	keyvaultIDRegex = regexp.MustCompile(`^/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.KeyVault/vaults/[^/\s]+$`)
}

// Validate implements APIObject
func (o *OrchestratorProfile) Validate(isUpdate, hasWindows bool) error {
	// Don't need to call validate.Struct(o)
	// It is handled by Properties.Validate()
	// On updates we only need to make sure there is a supported patch version for the minor version
	if !isUpdate {
		switch o.OrchestratorType {
		case Kubernetes:
			if k8sVersion := o.OrchestratorVersion; !common.IsSupportedKubernetesVersion(k8sVersion, isUpdate, hasWindows) && o.OrchestratorVersion != "" {
				return errors.Errorf("OrchestratorProfile has unknown orchestrator version: %s", o.OrchestratorVersion)
			}

		default:
			return errors.Errorf("OrchestratorProfile has unknown orchestrator: %s", o.OrchestratorType)
		}
	} else {
		switch o.OrchestratorType {
		case Kubernetes:
			patchVersion := common.GetValidPatchVersion(o.OrchestratorType, o.OrchestratorVersion, isUpdate, hasWindows)
			// if there isn't a supported patch version for this version fail
			if patchVersion == "" {
				return errors.Errorf("OrchestratorProfile has unknown orchestrator version: %s", o.OrchestratorVersion)
			}
		}
	}

	return nil
}

// Validate implements APIObject
func (m *MasterProfile) Validate() error {
	// Don't need to call validate.Struct(m)
	// It is handled by Properties.Validate()
	return common.ValidateDNSPrefix(m.DNSPrefix)
}

// Validate implements APIObject
func (a *AgentPoolProfile) Validate(orchestratorType string) error {
	// Don't need to call validate.Struct(a)
	// It is handled by Properties.Validate()
	if e := validatePoolName(a.Name); e != nil {
		return e
	}
	// Kubernetes don't allow agent DNSPrefix and ports
	if orchestratorType == Kubernetes {
		// The two lines below need to be removed after August 2017
		a.DNSPrefix = ""
		a.Ports = []int{}
		if e := validate.Var(a.DNSPrefix, "len=0"); e != nil {
			return errors.Errorf("AgentPoolProfile.DNSPrefix must be empty for Kubernetes")
		}
		if e := validate.Var(a.Ports, "len=0"); e != nil {
			return errors.Errorf("AgentPoolProfile.Ports must be empty for Kubernetes")
		}
	}
	if a.DNSPrefix != "" {
		if e := common.ValidateDNSPrefix(a.DNSPrefix); e != nil {
			return e
		}
		if len(a.Ports) > 0 {
			if e := validateUniquePorts(a.Ports, a.Name); e != nil {
				return e
			}
		} else {
			a.Ports = []int{80, 443, 8080}
		}
	} else {
		if e := validate.Var(a.Ports, "len=0"); e != nil {
			return errors.Errorf("AgentPoolProfile.Ports must be empty when AgentPoolProfile.DNSPrefix is empty for Orchestrator: %s", string(orchestratorType))
		}
	}
	return nil
}

// Validate implements APIObject
func (l *LinuxProfile) Validate() error {
	// Don't need to call validate.Struct(l)
	// It is handled by Properties.Validate()
	if e := validate.Var(l.SSH.PublicKeys[0].KeyData, "required"); e != nil {
		return errors.New("KeyData in LinuxProfile.SSH.PublicKeys cannot be empty string")
	}
	return nil
}

func handleValidationErrors(e validator.ValidationErrors) error {
	// Override any version specific validation error message

	// common.HandleValidationErrors if the validation error message is general
	return common.HandleValidationErrors(e)
}

// Validate implements APIObject
func (a *Properties) Validate(isUpdate bool) error {
	if e := validate.Struct(a); e != nil {
		return handleValidationErrors(e.(validator.ValidationErrors))
	}
	if e := a.OrchestratorProfile.Validate(isUpdate, a.HasWindows()); e != nil {
		return e
	}
	if e := a.MasterProfile.Validate(); e != nil {
		return e
	}
	if e := validateUniqueProfileNames(a.AgentPoolProfiles); e != nil {
		return e
	}

	if a.OrchestratorProfile.OrchestratorType == Kubernetes {
		if a.ServicePrincipalProfile == nil {
			return errors.Errorf("ServicePrincipalProfile must be specified with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
		}
		if e := validate.Var(a.ServicePrincipalProfile.ClientID, "required"); e != nil {
			return errors.Errorf("the service principal client ID must be specified with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
		}
		if (len(a.ServicePrincipalProfile.Secret) == 0 && a.ServicePrincipalProfile.KeyvaultSecretRef == nil) ||
			(len(a.ServicePrincipalProfile.Secret) != 0 && a.ServicePrincipalProfile.KeyvaultSecretRef != nil) {
			return errors.Errorf("either the service principal client secret or keyvault secret reference must be specified with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
		}

		if a.ServicePrincipalProfile.KeyvaultSecretRef != nil {
			if e := validate.Var(a.ServicePrincipalProfile.KeyvaultSecretRef.VaultID, "required"); e != nil {
				return errors.Errorf("the Keyvault ID must be specified for the Service Principle with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
			}
			if e := validate.Var(a.ServicePrincipalProfile.KeyvaultSecretRef.SecretName, "required"); e != nil {
				return errors.Errorf("the Keyvault Secret must be specified for the Service Principle with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
			}
			if !keyvaultIDRegex.MatchString(a.ServicePrincipalProfile.KeyvaultSecretRef.VaultID) {
				return errors.New("service principal client keyvault secret reference is of incorrect format")
			}
		}
	}

	for _, agentPoolProfile := range a.AgentPoolProfiles {
		if e := agentPoolProfile.Validate(a.OrchestratorProfile.OrchestratorType); e != nil {
			return e
		}

		if agentPoolProfile.OSType == Windows {
			if a.WindowsProfile == nil {
				return errors.New("missing WindowsProfile")
			}
			switch a.OrchestratorProfile.OrchestratorType {
			case Kubernetes:
			default:
				return errors.Errorf("Orchestrator %s does not support Windows", a.OrchestratorProfile.OrchestratorType)
			}
			if a.WindowsProfile == nil {
				return errors.Errorf("WindowsProfile must not be empty since agent pool '%s' specifies windows", agentPoolProfile.Name)
			}
			if len(a.WindowsProfile.AdminUsername) == 0 {
				return errors.Errorf("WindowsProfile.AdminUsername must not be empty since agent pool '%s' specifies windows", agentPoolProfile.Name)
			}
			if len(a.WindowsProfile.AdminPassword) == 0 {
				return errors.Errorf("WindowsProfile.AdminPassword must not be empty since  agent pool '%s' specifies windows", agentPoolProfile.Name)
			}
		}
	}
	if e := a.LinuxProfile.Validate(); e != nil {
		return e
	}
	return validateVNET(a)
}

func validatePoolName(poolName string) error {
	// we will cap at length of 12 and all lowercase letters since this makes up the VMName
	poolNameRegex := `^([a-z][a-z0-9]{0,11})$`
	re, err := regexp.Compile(poolNameRegex)
	if err != nil {
		return err
	}
	submatches := re.FindStringSubmatch(poolName)
	if len(submatches) != 2 {
		return errors.Errorf("pool name '%s' is invalid. A pool name must start with a lowercase letter, have max length of 12, and only have characters a-z0-9", poolName)
	}
	return nil
}

func validateUniqueProfileNames(profiles []*AgentPoolProfile) error {
	profileNames := make(map[string]bool)
	for _, profile := range profiles {
		if _, ok := profileNames[profile.Name]; ok {
			return errors.Errorf("profile name '%s' already exists, profile names must be unique across pools", profile.Name)
		}
		profileNames[profile.Name] = true
	}
	return nil
}

func validateUniquePorts(ports []int, name string) error {
	portMap := make(map[int]bool)
	for _, port := range ports {
		if _, ok := portMap[port]; ok {
			return errors.Errorf("agent profile '%s' has duplicate port '%d', ports must be unique", name, port)
		}
		portMap[port] = true
	}
	return nil
}

func validateVNET(a *Properties) error {
	isCustomVNET := a.MasterProfile.IsCustomVNET()
	for _, agentPool := range a.AgentPoolProfiles {
		if agentPool.IsCustomVNET() != isCustomVNET {
			return errors.New("Multiple VNET Subnet configurations specified.  The master profile and each agent pool profile must all specify a custom VNET Subnet, or none at all")
		}
	}
	if isCustomVNET {
		subscription, resourcegroup, vnetname, _, e := common.GetVNETSubnetIDComponents(a.MasterProfile.VnetSubnetID)
		if e != nil {
			return e
		}

		for _, agentPool := range a.AgentPoolProfiles {
			agentSubID, agentRG, agentVNET, _, err := common.GetVNETSubnetIDComponents(agentPool.VnetSubnetID)
			if err != nil {
				return err
			}
			if agentSubID != subscription ||
				agentRG != resourcegroup ||
				agentVNET != vnetname {
				return errors.New("Multiple VNETS specified.  The master profile and each agent pool must reference the same VNET (but it is ok to reference different subnets on that VNET)")
			}
		}

		masterFirstIP := net.ParseIP(a.MasterProfile.FirstConsecutiveStaticIP)
		if masterFirstIP == nil {
			return errors.Errorf("MasterProfile.FirstConsecutiveStaticIP (with VNET Subnet specification) '%s' is an invalid IP address", a.MasterProfile.FirstConsecutiveStaticIP)
		}
	}
	return nil
}
