package v20170131

import (
	"regexp"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/pkg/errors"
)

// Validate implements APIObject
func (o *OrchestratorProfile) Validate() error {
	switch o.OrchestratorType {
	case Kubernetes:
	default:
		return errors.Errorf("OrchestratorProfile has unknown orchestrator: %s", o.OrchestratorType)
	}

	return nil
}

// Validate implements APIObject
func (m *MasterProfile) Validate() error {
	if m.Count != 1 && m.Count != 3 && m.Count != 5 {
		return errors.New("MasterProfile count needs to be 1, 3, or 5")
	}
	if e := validateName(m.DNSPrefix, "MasterProfile.DNSPrefix"); e != nil {
		return e
	}
	return common.ValidateDNSPrefix(m.DNSPrefix)
}

// Validate implements APIObject
func (a *AgentPoolProfile) Validate(orchestratorType string) error {
	if e := validateName(a.Name, "AgentPoolProfile.Name"); e != nil {
		return e
	}
	if e := validatePoolName(a.Name); e != nil {
		return e
	}
	if a.Count < MinAgentCount || a.Count > MaxAgentCount {
		return errors.Errorf("AgentPoolProfile count needs to be in the range [%d,%d]", MinAgentCount, MaxAgentCount)
	}
	if e := validateName(a.VMSize, "AgentPoolProfile.VMSize"); e != nil {
		return e
	}
	// Kubernetes don't allow agent DNSPrefix
	if orchestratorType == Kubernetes {
		a.DNSPrefix = ""
		if e := validateNameEmpty(a.DNSPrefix, "AgentPoolProfile.DNSPrefix"); e != nil {
			return e
		}
	}
	if a.DNSPrefix != "" {
		if e := common.ValidateDNSPrefix(a.DNSPrefix); e != nil {
			return e
		}
	}
	return nil
}

// Validate implements APIObject
func (l *LinuxProfile) Validate() error {
	if e := validateName(l.AdminUsername, "LinuxProfile.AdminUsername"); e != nil {
		return e
	}
	if len(l.SSH.PublicKeys) != 1 {
		return errors.New("LinuxProfile.PublicKeys requires only 1 SSH Key")
	}
	return validateName(l.SSH.PublicKeys[0].KeyData, "LinuxProfile.PublicKeys.KeyData")
}

// Validate implements APIObject
func (a *Properties) Validate() error {
	if a.OrchestratorProfile == nil {
		return errors.New("missing OrchestratorProfile")
	}
	if a.MasterProfile == nil {
		return errors.New("missing MasterProfile")
	}
	if a.LinuxProfile == nil {
		return errors.New("missing LinuxProfile")
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

		if len(a.ServicePrincipalProfile.Secret) == 0 {
			return errors.Errorf("service principal client secret must be specified with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
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
	return a.OrchestratorProfile.Validate()
}

func validateNameEmpty(name string, label string) error {
	if name != "" {
		return errors.Errorf("%s must be an empty value", label)
	}
	return nil
}

func validateName(name string, label string) error {
	if name == "" {
		return errors.Errorf("%s must be a non-empty value", label)
	}
	return nil
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
