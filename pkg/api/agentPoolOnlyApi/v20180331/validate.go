package v20180331

import (
	"net"
	"regexp"
	"strings"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/pkg/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	// KubernetesMinMaxPods is the minimum valid value for MaxPods, necessary for running kube-system pods
	KubernetesMinMaxPods = 5
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate implements APIObject
func (a *AgentPoolProfile) Validate() error {
	// Don't need to call validate.Struct(a)
	// It is handled by Properties.Validate()
	return validatePoolName(a.Name)
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

// Validate implements APIObject
func (a *AADProfile) Validate(rbacEnabled *bool) error {
	if !helpers.IsTrueBoolPointer(rbacEnabled) {
		return ErrorRBACNotEnabledForAAD
	}

	if e := validate.Var(a.ServerAppID, "required"); e != nil {
		return ErrorAADServerAppIDNotSet
	}

	// Don't need to call validate.Struct(l)
	// It is handled by Properties.Validate()
	if e := validate.Var(a.ServerAppSecret, "required"); e != nil {
		return ErrorAADServerAppSecretNotSet
	}

	if e := validate.Var(a.ClientAppID, "required"); e != nil {
		return ErrorAADClientAppIDNotSet
	}

	if e := validate.Var(a.TenantID, "required"); e != nil {
		return ErrorAADTenantIDNotSet
	}

	return nil
}

func handleValidationErrors(e validator.ValidationErrors) error {
	err := e[0]
	ns := err.Namespace()
	switch ns {
	// TODO: Add more validation here
	case "Properties.ServicePrincipalProfile.ClientID",
		"Properties.ServicePrincipalProfile.Secret", "Properties.WindowsProfile.AdminUsername",
		"Properties.WindowsProfile.AdminPassword":
		return errors.Errorf("missing %s", ns)
	default:
		if strings.HasPrefix(ns, "Properties.AgentPoolProfiles") {
			switch {
			case strings.HasSuffix(ns, ".Name") || strings.HasSuffix(ns, "VMSize"):
				return errors.Errorf("missing %s", ns)
			case strings.HasSuffix(ns, ".Count"):
				return errors.Errorf("AgentPoolProfile count needs to be in the range [%d,%d]", MinAgentCount, MaxAgentCount)
			case strings.HasSuffix(ns, ".OSDiskSizeGB"):
				return errors.Errorf("Invalid os disk size of %d specified.  The range of valid values are [%d, %d]", err.Value().(int), MinDiskSizeGB, MaxDiskSizeGB)
			case strings.HasSuffix(ns, ".StorageProfile"):
				return errors.Errorf("Unknown storageProfile '%s'. Must specify %s", err.Value().(string), ManagedDisks)
			default:
				break
			}
		}
	}
	return errors.Errorf("Namespace %s is not caught, %+v", ns, e)
}

// Validate implements APIObject
func (a *Properties) Validate() error {
	if e := validate.Struct(a); e != nil {
		return handleValidationErrors(e.(validator.ValidationErrors))
	}

	// Don't need to call validate.Struct(m)
	// It is handled by Properties.Validate()
	if e := common.ValidateDNSPrefix(a.DNSPrefix); e != nil {
		return e
	}

	if e := validateUniqueProfileNames(a.AgentPoolProfiles); e != nil {
		return e
	}

	for _, agentPoolProfile := range a.AgentPoolProfiles {
		if e := agentPoolProfile.Validate(); e != nil {
			return e
		}
	}

	if a.LinuxProfile != nil {
		if e := a.LinuxProfile.Validate(); e != nil {
			return e
		}
	}

	if e := validateVNET(a); e != nil {
		return e
	}

	if a.AADProfile != nil {
		if e := a.AADProfile.Validate(a.EnableRBAC); e != nil {
			return e
		}
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

// validateVNET validate network profile and custom VNET logic
func validateVNET(a *Properties) error {

	n := a.NetworkProfile

	// validate network profile settings
	if n != nil {
		switch n.NetworkPlugin {
		case Azure, Kubenet:
			if n.ServiceCidr != "" && n.DNSServiceIP != "" && n.DockerBridgeCidr != "" {
				// validate ServiceCidr
				_, serviceCidr, err := net.ParseCIDR(n.ServiceCidr)
				if err != nil {
					return ErrorInvalidServiceCidr
				}

				// validate ServiceCidr not too large
				var ones, bits = serviceCidr.Mask.Size()
				if bits-ones > 20 {
					return ErrorServiceCidrTooLarge
				}

				// validate DNSServiceIP
				dnsServiceIP := net.ParseIP(n.DNSServiceIP)
				if dnsServiceIP == nil {
					return ErrorInvalidDNSServiceIP
				}

				// validate DockerBridgeCidr
				_, _, err = net.ParseCIDR(n.DockerBridgeCidr)
				if err != nil {
					return ErrorInvalidDockerBridgeCidr
				}

				// validate DNSServiceIP is within ServiceCidr
				if !serviceCidr.Contains(dnsServiceIP) {
					return ErrorDNSServiceIPNotInServiceCidr
				}

				// validate DNSServiceIP is not the first IP in ServiceCidr. The first IP is reserved for redirect svc.
				kubernetesServiceIP, err := common.CidrStringFirstIP(n.ServiceCidr)
				if err != nil {
					return ErrorInvalidServiceCidr
				}
				if dnsServiceIP.String() == kubernetesServiceIP.String() {
					return ErrorDNSServiceIPAlreadyUsed
				}
			} else if n.ServiceCidr == "" && n.DNSServiceIP == "" && n.DockerBridgeCidr == "" {
				// this is a valid case, and no validation needed.
			} else {
				return ErrorInvalidNetworkProfile
			}

			// PodCidr should not be set for Azure CNI
			if n.NetworkPlugin == Azure && n.PodCidr != "" {
				return ErrorPodCidrNotSetableInAzureCNI
			}
		default:
			return ErrorInvalidNetworkPlugin
		}
	}

	// validate agent pool custom VNET settings
	if a.AgentPoolProfiles != nil {
		if e := validateAgentPoolVNET(a.AgentPoolProfiles); e != nil {
			return e
		}
	}

	return nil
}

func validateAgentPoolVNET(a []*AgentPoolProfile) error {

	// validate custom VNET logic at agent pool level
	if isCustomVNET(a) {
		var subscription string
		var resourceGroup string
		var vnet string

		for _, agentPool := range a {
			// validate each agent pool has a subnet
			if !agentPool.IsCustomVNET() {
				return ErrorAtLeastAgentPoolNoSubnet
			}

			if agentPool.MaxPods != nil && *agentPool.MaxPods < KubernetesMinMaxPods {
				return ErrorInvalidMaxPods
			}

			// validate subscription, resource group and vnet are the same among subnets
			subnetSubscription, subnetResourceGroup, subnetVnet, _, err := common.GetVNETSubnetIDComponents(agentPool.VnetSubnetID)
			if err != nil {
				return ErrorParsingSubnetID
			}

			if subscription == "" {
				subscription = subnetSubscription
			} else {
				if subscription != subnetSubscription {
					return ErrorSubscriptionNotMatch
				}
			}

			if resourceGroup == "" {
				resourceGroup = subnetResourceGroup
			} else {
				if resourceGroup != subnetResourceGroup {
					return ErrorResourceGroupNotMatch
				}
			}

			if vnet == "" {
				vnet = subnetVnet
			} else {
				if vnet != subnetVnet {
					return ErrorVnetNotMatch
				}
			}
		}
	}

	return nil
}

// check agent pool subnet, return true as long as one agent pool has a subnet defined.
func isCustomVNET(a []*AgentPoolProfile) bool {
	for _, agentPool := range a {
		if agentPool.IsCustomVNET() {
			return true
		}
	}

	return false
}
