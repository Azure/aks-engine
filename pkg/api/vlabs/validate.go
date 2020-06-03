// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/blang/semver"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	validate                       *validator.Validate
	keyvaultIDRegex                *regexp.Regexp
	labelValueRegex                *regexp.Regexp
	labelKeyRegex                  *regexp.Regexp
	diskEncryptionSetIDRegex       *regexp.Regexp
	proximityPlacementGroupIDRegex *regexp.Regexp
	// Any version has to be available in a container image from mcr.microsoft.com/oss/etcd-io/etcd:v[Version]
	etcdValidVersions = [...]string{"2.2.5", "2.3.0", "2.3.1", "2.3.2", "2.3.3", "2.3.4", "2.3.5", "2.3.6", "2.3.7", "2.3.8",
		"3.0.0", "3.0.1", "3.0.2", "3.0.3", "3.0.4", "3.0.5", "3.0.6", "3.0.7", "3.0.8", "3.0.9", "3.0.10", "3.0.11", "3.0.12", "3.0.13", "3.0.14", "3.0.15", "3.0.16", "3.0.17",
		"3.1.0", "3.1.1", "3.1.2", "3.1.2", "3.1.3", "3.1.4", "3.1.5", "3.1.6", "3.1.7", "3.1.8", "3.1.9", "3.1.10",
		"3.2.0", "3.2.1", "3.2.2", "3.2.3", "3.2.4", "3.2.5", "3.2.6", "3.2.7", "3.2.8", "3.2.9", "3.2.11", "3.2.12",
		"3.2.13", "3.2.14", "3.2.15", "3.2.16", "3.2.23", "3.2.24", "3.2.25", "3.2.26", "3.3.0", "3.3.1", "3.3.8", "3.3.9", "3.3.10", "3.3.13", "3.3.15", "3.3.18"}
	containerdValidVersions        = [...]string{"1.1.5", "1.1.6", "1.2.4"}
	networkPluginPlusPolicyAllowed = []k8sNetworkConfig{
		{
			networkPlugin: "",
			networkPolicy: "",
		},
		{
			networkPlugin: "azure",
			networkPolicy: "",
		},
		{
			networkPlugin: "azure",
			networkPolicy: "azure",
		},
		{
			networkPlugin: "kubenet",
			networkPolicy: "",
		},
		{
			networkPlugin: "flannel",
			networkPolicy: "",
		},
		{
			networkPlugin: NetworkPluginCilium,
			networkPolicy: NetworkPolicyCilium,
		},
		{
			networkPlugin: "kubenet",
			networkPolicy: "calico",
		},
		{
			networkPlugin: "azure",
			networkPolicy: "calico",
		},
		{
			networkPlugin: "",
			networkPolicy: "calico",
		},
		{
			networkPlugin: "",
			networkPolicy: NetworkPolicyCilium,
		},
		{
			networkPlugin: NetworkPluginAntrea,
			networkPolicy: NetworkPolicyAntrea,
		},
		{
			networkPlugin: "",
			networkPolicy: NetworkPolicyAntrea,
		},
		{
			networkPlugin: "",
			networkPolicy: "azure", // for backwards-compatibility w/ prior networkPolicy usage
		},
		{
			networkPlugin: "",
			networkPolicy: "none", // for backwards-compatibility w/ prior networkPolicy usage
		},
	}
)

const (
	labelKeyPrefixMaxLength = 253
	labelValueFormat        = "^([A-Za-z0-9][-A-Za-z0-9_.]{0,61})?[A-Za-z0-9]$"
	labelKeyFormat          = "^(([a-zA-Z0-9-]+[.])*[a-zA-Z0-9-]+[/])?([A-Za-z0-9][-A-Za-z0-9_.]{0,61})?[A-Za-z0-9]$"
)

type k8sNetworkConfig struct {
	networkPlugin string
	networkPolicy string
}

func init() {
	validate = validator.New()
	keyvaultIDRegex = regexp.MustCompile(`^/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.KeyVault/vaults/[^/\s]+$`)
	labelValueRegex = regexp.MustCompile(labelValueFormat)
	labelKeyRegex = regexp.MustCompile(labelKeyFormat)
	diskEncryptionSetIDRegex = regexp.MustCompile(`^/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.Compute/diskEncryptionSets/[^/\s]+$`)
	proximityPlacementGroupIDRegex = regexp.MustCompile(`^/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.Compute/proximityPlacementGroups/[^/\s]+$`)
}

// Validate implements APIObject
func (a *Properties) validate(isUpdate bool) error {
	if e := validate.Struct(a); e != nil {
		return handleValidationErrors(e.(validator.ValidationErrors))
	}
	if e := a.ValidateOrchestratorProfile(isUpdate); e != nil {
		return e
	}
	if e := a.validateMasterProfile(isUpdate); e != nil {
		return e
	}
	if e := a.validateAgentPoolProfiles(isUpdate); e != nil {
		return e
	}
	if e := a.validateZones(); e != nil {
		return e
	}
	if e := a.validateLinuxProfile(); e != nil {
		return e
	}
	if e := a.validateAddons(); e != nil {
		return e
	}
	if e := a.validateExtensions(); e != nil {
		return e
	}
	if e := a.validateVNET(); e != nil {
		return e
	}
	if e := a.validateServicePrincipalProfile(); e != nil {
		return e
	}

	if e := a.validateManagedIdentity(); e != nil {
		return e
	}

	if e := a.validateAADProfile(); e != nil {
		return e
	}

	if e := a.validateCustomKubeComponent(); e != nil {
		return e
	}

	if e := a.validatePrivateAzureRegistryServer(); e != nil {
		return e
	}

	if e := a.validateAzureStackSupport(); e != nil {
		return e
	}

	if e := a.validateWindowsProfile(); e != nil {
		return e
	}
	return nil
}

func handleValidationErrors(e validator.ValidationErrors) error {
	// Override any version specific validation error message
	// common.HandleValidationErrors if the validation error message is general
	return common.HandleValidationErrors(e)
}

//ValidateOrchestratorProfile validates the orchestrator profile and the addons dependent on the version of the orchestrator
func (a *Properties) ValidateOrchestratorProfile(isUpdate bool) error {
	o := a.OrchestratorProfile
	// On updates we only need to make sure there is a supported patch version for the minor version
	if !isUpdate {
		switch o.OrchestratorType {
		case DCOS:
			version := common.RationalizeReleaseAndVersion(
				o.OrchestratorType,
				o.OrchestratorRelease,
				o.OrchestratorVersion,
				isUpdate,
				false)
			if version == "" {
				return errors.Errorf("the following OrchestratorProfile configuration is not supported: OrchestratorType: %s, OrchestratorRelease: %s, OrchestratorVersion: %s. Please check supported Release or Version for this build of aks-engine", o.OrchestratorType, o.OrchestratorRelease, o.OrchestratorVersion)
			}
			if o.DcosConfig != nil && o.DcosConfig.BootstrapProfile != nil {
				if len(o.DcosConfig.BootstrapProfile.StaticIP) > 0 {
					if net.ParseIP(o.DcosConfig.BootstrapProfile.StaticIP) == nil {
						return errors.Errorf("DcosConfig.BootstrapProfile.StaticIP '%s' is an invalid IP address",
							o.DcosConfig.BootstrapProfile.StaticIP)
					}
				}
			}
		case Swarm:
		case SwarmMode:
		case Kubernetes:
			version := common.RationalizeReleaseAndVersion(
				o.OrchestratorType,
				o.OrchestratorRelease,
				o.OrchestratorVersion,
				isUpdate,
				a.HasWindows())
			if version == "" && a.HasWindows() {
				return errors.Errorf("the following OrchestratorProfile configuration is not supported with OsType \"Windows\": OrchestratorType: \"%s\", OrchestratorRelease: \"%s\", OrchestratorVersion: \"%s\". Please use one of the following versions: %v", o.OrchestratorType, o.OrchestratorRelease, o.OrchestratorVersion, common.GetAllSupportedKubernetesVersions(false, true))
			} else if version == "" {
				return errors.Errorf("the following OrchestratorProfile configuration is not supported: OrchestratorType: \"%s\", OrchestratorRelease: \"%s\", OrchestratorVersion: \"%s\". Please use one of the following versions: %v", o.OrchestratorType, o.OrchestratorRelease, o.OrchestratorVersion, common.GetAllSupportedKubernetesVersions(false, false))
			}

			sv, err := semver.Make(version)
			if err != nil {
				return errors.Errorf("could not validate version %s", version)
			}

			if a.HasAvailabilityZones() {
				minVersion, err := semver.Make("1.12.0")
				if err != nil {
					return errors.New("could not validate version")
				}

				if sv.LT(minVersion) {
					return errors.New("availabilityZone is only available in Kubernetes version 1.12 or greater")
				}
			}

			if o.KubernetesConfig != nil {
				err := o.KubernetesConfig.Validate(version, a.HasWindows(), a.FeatureFlags.IsIPv6DualStackEnabled(), a.FeatureFlags.IsIPv6OnlyEnabled())
				if err != nil {
					return err
				}
				minVersion, err := semver.Make("1.7.0")
				if err != nil {
					return errors.New("could not validate version")
				}

				if o.KubernetesConfig.EnableAggregatedAPIs {
					if !o.KubernetesConfig.IsRBACEnabled() {
						return errors.New("enableAggregatedAPIs requires the enableRbac feature as a prerequisite")
					}
				}

				if to.Bool(o.KubernetesConfig.EnableDataEncryptionAtRest) {
					if sv.LT(minVersion) {
						return errors.Errorf("enableDataEncryptionAtRest is only available in Kubernetes version %s or greater; unable to validate for Kubernetes version %s",
							minVersion.String(), o.OrchestratorVersion)
					}
					if o.KubernetesConfig.EtcdEncryptionKey != "" {
						_, err = base64.StdEncoding.DecodeString(o.KubernetesConfig.EtcdEncryptionKey)
						if err != nil {
							return errors.New("etcdEncryptionKey must be base64 encoded. Please provide a valid base64 encoded value or leave the etcdEncryptionKey empty to auto-generate the value")
						}
					}
				}

				if to.Bool(o.KubernetesConfig.EnableEncryptionWithExternalKms) {
					minVersion, err := semver.Make("1.10.0")
					if err != nil {
						return errors.Errorf("could not validate version")
					}
					if sv.LT(minVersion) {
						return errors.Errorf("enableEncryptionWithExternalKms is only available in Kubernetes version %s or greater; unable to validate for Kubernetes version %s",
							minVersion.String(), o.OrchestratorVersion)
					}
				}

				if o.KubernetesConfig.EnableRbac != nil && !o.KubernetesConfig.IsRBACEnabled() {
					minVersionNotAllowed, err := semver.Make("1.15.0")
					if err != nil {
						return errors.Errorf("could not validate version")
					}
					if !sv.LT(minVersionNotAllowed) {
						return errors.Errorf("RBAC support is required for Kubernetes version %s or greater; unable to build Kubernetes v%s cluster with enableRbac=false",
							minVersionNotAllowed.String(), o.OrchestratorVersion)
					}
				}

				if to.Bool(o.KubernetesConfig.EnablePodSecurityPolicy) {
					log.Warnf("EnablePodSecurityPolicy is deprecated in favor of the addon pod-security-policy.")
					if !o.KubernetesConfig.IsRBACEnabled() {
						return errors.Errorf("enablePodSecurityPolicy requires the enableRbac feature as a prerequisite")
					}
					if len(o.KubernetesConfig.PodSecurityPolicyConfig) > 0 {
						log.Warnf("Raw manifest for PodSecurityPolicy using PodSecurityPolicyConfig is deprecated in favor of the addon pod-security-policy. This will be ignored.")
					}
				}

				if o.KubernetesConfig.LoadBalancerSku != "" {
					if strings.ToLower(o.KubernetesConfig.LoadBalancerSku) != strings.ToLower(StandardLoadBalancerSku) && strings.ToLower(o.KubernetesConfig.LoadBalancerSku) != strings.ToLower(BasicLoadBalancerSku) {
						return errors.Errorf("Invalid value for loadBalancerSku, only %s and %s are supported", StandardLoadBalancerSku, BasicLoadBalancerSku)
					}
				}

				if o.KubernetesConfig.LoadBalancerSku == StandardLoadBalancerSku {
					minVersion, err := semver.Make("1.11.0")
					if err != nil {
						return errors.Errorf("could not validate version")
					}
					if sv.LT(minVersion) {
						return errors.Errorf("loadBalancerSku is only available in Kubernetes version %s or greater; unable to validate for Kubernetes version %s",
							minVersion.String(), o.OrchestratorVersion)
					}
					if !to.Bool(a.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB) {
						return errors.Errorf("standard loadBalancerSku should exclude master nodes. Please set KubernetesConfig \"ExcludeMasterFromStandardLB\" to \"true\"")
					}
				}

				if o.KubernetesConfig.DockerEngineVersion != "" {
					log.Warnf("docker-engine is deprecated in favor of moby, but you passed in a dockerEngineVersion configuration. This will be ignored.")
				}

				if o.KubernetesConfig.MaximumLoadBalancerRuleCount < 0 {
					return errors.New("maximumLoadBalancerRuleCount shouldn't be less than 0")
				}
				// https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-outbound-rules-overview
				if o.KubernetesConfig.LoadBalancerSku == StandardLoadBalancerSku && o.KubernetesConfig.OutboundRuleIdleTimeoutInMinutes != 0 && (o.KubernetesConfig.OutboundRuleIdleTimeoutInMinutes < 4 || o.KubernetesConfig.OutboundRuleIdleTimeoutInMinutes > 120) {
					return errors.New("outboundRuleIdleTimeoutInMinutes shouldn't be less than 4 or greater than 120")
				}

				if a.IsAzureStackCloud() {
					if to.Bool(o.KubernetesConfig.UseInstanceMetadata) {
						return errors.New("useInstanceMetadata shouldn't be set to true as feature not yet supported on Azure Stack")
					}

					if o.KubernetesConfig.EtcdDiskSizeGB != "" {
						etcdDiskSizeGB, err := strconv.Atoi(o.KubernetesConfig.EtcdDiskSizeGB)
						if err != nil {
							return errors.Errorf("could not convert EtcdDiskSizeGB to int")
						}
						if etcdDiskSizeGB > MaxAzureStackManagedDiskSize {
							return errors.Errorf("EtcdDiskSizeGB max size supported on Azure Stack is %d", MaxAzureStackManagedDiskSize)
						}
					}
				}
			}
		default:
			return errors.Errorf("OrchestratorProfile has unknown orchestrator: %s", o.OrchestratorType)
		}
	} else {
		switch o.OrchestratorType {
		case DCOS, Kubernetes:

			version := common.RationalizeReleaseAndVersion(
				o.OrchestratorType,
				o.OrchestratorRelease,
				o.OrchestratorVersion,
				false,
				a.HasWindows())
			if version == "" {
				patchVersion := common.GetValidPatchVersion(o.OrchestratorType, o.OrchestratorVersion, isUpdate, a.HasWindows())
				// if there isn't a supported patch version for this version fail
				if patchVersion == "" {
					if a.HasWindows() {
						return errors.Errorf("the following OrchestratorProfile configuration is not supported with Windows agentpools: OrchestratorType: \"%s\", OrchestratorRelease: \"%s\", OrchestratorVersion: \"%s\". Please check supported Release or Version for this build of aks-engine", o.OrchestratorType, o.OrchestratorRelease, o.OrchestratorVersion)
					}
					return errors.Errorf("the following OrchestratorProfile configuration is not supported: OrchestratorType: \"%s\", OrchestratorRelease: \"%s\", OrchestratorVersion: \"%s\". Please check supported Release or Version for this build of aks-engine", o.OrchestratorType, o.OrchestratorRelease, o.OrchestratorVersion)
				}
			}

		}
	}

	if o.OrchestratorType != Kubernetes && o.KubernetesConfig != nil {
		return errors.Errorf("KubernetesConfig can be specified only when OrchestratorType is Kubernetes")
	}

	if o.OrchestratorType != DCOS && o.DcosConfig != nil && (*o.DcosConfig != DcosConfig{}) {
		return errors.Errorf("DcosConfig can be specified only when OrchestratorType is DCOS")
	}

	return a.validateContainerRuntime()
}

func (a *Properties) validateMasterProfile(isUpdate bool) error {
	m := a.MasterProfile

	if a.OrchestratorProfile.OrchestratorType == Kubernetes {
		if m.IsVirtualMachineScaleSets() && m.VnetSubnetID != "" && m.FirstConsecutiveStaticIP != "" {
			return errors.New("when masterProfile's availabilityProfile is VirtualMachineScaleSets and a vnetSubnetID is specified, the firstConsecutiveStaticIP should be empty and will be determined by an offset from the first IP in the vnetCidr")
		}
		// validate distro is ubuntu if dual stack or ipv6 only feature is enabled
		if a.FeatureFlags.IsIPv6DualStackEnabled() || a.FeatureFlags.IsIPv6OnlyEnabled() {
			if m.Distro == CoreOS {
				return errors.Errorf("Dual stack and single stack IPv6 feature is currently supported only with Ubuntu, but master is of distro type %s", m.Distro)
			}
		}
	}

	if m.ImageRef != nil {
		if m.Distro != "" {
			return errors.New("masterProfile includes a custom image configuration (imageRef) and an explicit distro configuration, you may use one of these but not both simultaneously")
		}
		if err := m.ImageRef.validateImageNameAndGroup(); err != nil {
			return err
		}
	}

	if m.IsVirtualMachineScaleSets() && a.OrchestratorProfile.OrchestratorType == Kubernetes {
		log.Warnf("Clusters with VMSS masters are not yet upgradable! You will not be able to upgrade your cluster until a future version of aks-engine!")
		e := validateVMSS(a.OrchestratorProfile, false, m.StorageProfile)
		if e != nil {
			return e
		}
		if !a.IsClusterAllVirtualMachineScaleSets() {
			return errors.New("VirtualMachineScaleSets for master profile must be used together with virtualMachineScaleSets for agent profiles. Set \"availabilityProfile\" to \"VirtualMachineScaleSets\" for agent profiles")
		}

		if a.OrchestratorProfile.KubernetesConfig != nil && a.OrchestratorProfile.KubernetesConfig.UseManagedIdentity && a.OrchestratorProfile.KubernetesConfig.UserAssignedID == "" {
			return errors.New("virtualMachineScaleSets for master profile can be used only with user assigned MSI ! Please specify \"userAssignedID\" in \"kubernetesConfig\"")
		}
	}
	if m.SinglePlacementGroup != nil && m.AvailabilityProfile == AvailabilitySet {
		return errors.New("singlePlacementGroup is only supported with VirtualMachineScaleSets")
	}

	if e := validateProximityPlacementGroupID(m.ProximityPlacementGroupID); e != nil {
		return e
	}

	distroValues := DistroValues
	if isUpdate {
		distroValues = append(distroValues, AKSDockerEngine, AKS1604Deprecated, AKS1804Deprecated)
	}
	if !validateDistro(m.Distro, distroValues) {
		switch m.Distro {
		case AKSDockerEngine, AKS1604Deprecated:
			return errors.Errorf("The %s distro is deprecated, please use %s instead", m.Distro, AKSUbuntu1604)
		case AKS1804Deprecated:
			return errors.Errorf("The %s distro is deprecated, please use %s instead", m.Distro, AKSUbuntu1804)
		default:
			return errors.Errorf("The %s distro is not supported", m.Distro)
		}
	}

	if to.Bool(m.AuditDEnabled) {
		if m.Distro != "" && !m.IsUbuntu() {
			return errors.Errorf("You have enabled auditd for master vms, but you did not specify an Ubuntu-based distro.")
		}
	}

	return common.ValidateDNSPrefix(m.DNSPrefix)
}

func (a *Properties) validateAgentPoolProfiles(isUpdate bool) error {

	profileNames := make(map[string]bool)
	for i, agentPoolProfile := range a.AgentPoolProfiles {
		if e := validatePoolName(agentPoolProfile.Name); e != nil {
			return e
		}

		// validate os type is linux if dual stack feature is enabled
		if a.FeatureFlags.IsIPv6DualStackEnabled() || a.FeatureFlags.IsIPv6OnlyEnabled() {
			if agentPoolProfile.OSType == Windows {
				return errors.Errorf("Dual stack and single stack IPv6 feature is supported only with Linux, but agent pool '%s' is of os type %s", agentPoolProfile.Name, agentPoolProfile.OSType)
			}
			if agentPoolProfile.Distro == CoreOS {
				return errors.Errorf("Dual stack and single stack IPv6 feature is currently supported only with Ubuntu, but agent pool '%s' is of distro type %s", agentPoolProfile.Name, agentPoolProfile.Distro)
			}
		}

		// validate that each AgentPoolProfile Name is unique
		if _, ok := profileNames[agentPoolProfile.Name]; ok {
			return errors.Errorf("profile name '%s' already exists, profile names must be unique across pools", agentPoolProfile.Name)
		}
		profileNames[agentPoolProfile.Name] = true

		if e := validatePoolOSType(agentPoolProfile.OSType); e != nil {
			return e
		}

		if to.Bool(agentPoolProfile.AcceleratedNetworkingEnabled) || to.Bool(agentPoolProfile.AcceleratedNetworkingEnabledWindows) {
			if a.IsAzureStackCloud() {
				return errors.Errorf("AcceleratedNetworkingEnabled or AcceleratedNetworkingEnabledWindows shouldn't be set to true as feature is not yet supported on Azure Stack")
			} else if to.Bool(agentPoolProfile.AcceleratedNetworkingEnabledWindows) {
				return errors.Errorf("Accelerated Networking is currently unstable for Windows + Kubernetes, please set acceleratedNetworkingEnabledWindows to false")
			} else if e := validatePoolAcceleratedNetworking(agentPoolProfile.VMSize); e != nil {
				return e
			}
		}

		if to.Bool(agentPoolProfile.VMSSOverProvisioningEnabled) {
			if agentPoolProfile.AvailabilityProfile != VirtualMachineScaleSets {
				return errors.Errorf("You have specified VMSS Overprovisioning in agent pool %s, but you did not specify VMSS", agentPoolProfile.Name)
			}
		}

		if to.Bool(agentPoolProfile.AuditDEnabled) {
			if agentPoolProfile.Distro != "" && !agentPoolProfile.IsUbuntu() {
				return errors.Errorf("You have enabled auditd in agent pool %s, but you did not specify an Ubuntu-based distro", agentPoolProfile.Name)
			}
		}

		if to.Bool(agentPoolProfile.EnableVMSSNodePublicIP) {
			if agentPoolProfile.AvailabilityProfile != VirtualMachineScaleSets {
				return errors.Errorf("You have enabled VMSS node public IP in agent pool %s, but you did not specify VMSS", agentPoolProfile.Name)
			}
		}

		if e := agentPoolProfile.validateOrchestratorSpecificProperties(a.OrchestratorProfile.OrchestratorType); e != nil {
			return e
		}

		if agentPoolProfile.ImageRef != nil {
			if agentPoolProfile.Distro != "" {
				return errors.Errorf("agentPoolProfile %s includes a custom image configuration (imageRef) and an explicit distro configuration, you may use one of these but not both simultaneously", agentPoolProfile.Name)
			}
			return agentPoolProfile.ImageRef.validateImageNameAndGroup()
		}

		if e := agentPoolProfile.validateAvailabilityProfile(); e != nil {
			return e
		}

		if e := agentPoolProfile.validateRoles(a.OrchestratorProfile.OrchestratorType); e != nil {
			return e
		}

		if e := agentPoolProfile.validateStorageProfile(a.OrchestratorProfile.OrchestratorType); e != nil {
			return e
		}

		if e := agentPoolProfile.validateCustomNodeLabels(a.OrchestratorProfile.OrchestratorType); e != nil {
			return e
		}

		if agentPoolProfile.AvailabilityProfile == VirtualMachineScaleSets {
			e := validateVMSS(a.OrchestratorProfile, isUpdate, agentPoolProfile.StorageProfile)
			if e != nil {
				return e
			}
		}

		if a.OrchestratorProfile.OrchestratorType == Kubernetes {
			if a.AgentPoolProfiles[i].AvailabilityProfile != a.AgentPoolProfiles[0].AvailabilityProfile {
				return errors.New("mixed mode availability profiles are not allowed. Please set either VirtualMachineScaleSets or AvailabilitySet in availabilityProfile for all agent pools")
			}

			if a.AgentPoolProfiles[i].SinglePlacementGroup != nil && a.AgentPoolProfiles[i].AvailabilityProfile == AvailabilitySet {
				return errors.New("singlePlacementGroup is only supported with VirtualMachineScaleSets")
			}

			distroValues := DistroValues
			if isUpdate {
				distroValues = append(distroValues, AKSDockerEngine, AKS1604Deprecated, AKS1804Deprecated)
			}
			if !validateDistro(agentPoolProfile.Distro, distroValues) {
				switch agentPoolProfile.Distro {
				case AKSDockerEngine, AKS1604Deprecated:
					return errors.Errorf("The %s distro is deprecated, please use %s instead", agentPoolProfile.Distro, AKSUbuntu1604)
				case AKS1804Deprecated:
					return errors.Errorf("The %s distro is deprecated, please use %s instead", agentPoolProfile.Distro, AKSUbuntu1804)
				default:
					return errors.Errorf("The %s distro is not supported", agentPoolProfile.Distro)
				}
			}
		}

		if e := agentPoolProfile.validateLoadBalancerBackendAddressPoolIDs(); e != nil {
			return e
		}

		if agentPoolProfile.IsEphemeral() {
			log.Warnf("Ephemeral disks are enabled for Agent Pool %s. This feature in AKS-Engine is experimental, and data could be lost in some cases.", agentPoolProfile.Name)
		}

		if e := validateProximityPlacementGroupID(agentPoolProfile.ProximityPlacementGroupID); e != nil {
			return e
		}
	}

	return nil
}

func (a *Properties) validateZones() error {
	if a.OrchestratorProfile.OrchestratorType == Kubernetes {
		// all zones or no zones should be defined for the cluster
		if a.HasAvailabilityZones() {
			if a.MastersAndAgentsUseAvailabilityZones() {
				// agent pool profiles
				for _, agentPoolProfile := range a.AgentPoolProfiles {
					if agentPoolProfile.AvailabilityProfile == AvailabilitySet {
						return errors.New("Availability Zones are not supported with an AvailabilitySet. Please either remove availabilityProfile or set availabilityProfile to VirtualMachineScaleSets")
					}
				}
				if a.OrchestratorProfile.KubernetesConfig != nil && a.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != "" && strings.ToLower(a.OrchestratorProfile.KubernetesConfig.LoadBalancerSku) != strings.ToLower(StandardLoadBalancerSku) {
					return errors.New("Availability Zones requires Standard LoadBalancer. Please set KubernetesConfig \"LoadBalancerSku\" to \"Standard\"")
				}
			} else {
				return errors.New("Availability Zones need to be defined for master profile and all agent pool profiles. Please set \"availabilityZones\" for all profiles")
			}
		}
	}
	return nil
}

func (a *Properties) validateLinuxProfile() error {
	for _, publicKey := range a.LinuxProfile.SSH.PublicKeys {
		if e := validate.Var(publicKey.KeyData, "required"); e != nil {
			return errors.New("KeyData in LinuxProfile.SSH.PublicKeys cannot be empty string")
		}
	}
	return validateKeyVaultSecrets(a.LinuxProfile.Secrets, false)
}

func (a *Properties) validateAddons() error {
	if a.OrchestratorProfile.KubernetesConfig != nil && a.OrchestratorProfile.KubernetesConfig.Addons != nil {
		var isAvailabilitySets bool
		var IsNSeriesSKU bool
		var kubeDNSEnabled bool
		var corednsEnabled bool

		for _, agentPool := range a.AgentPoolProfiles {
			if agentPool.IsAvailabilitySets() {
				isAvailabilitySets = true
			}

			if agentPool.IsNSeriesSKU() {
				IsNSeriesSKU = true
			}
		}
		for _, addon := range a.OrchestratorProfile.KubernetesConfig.Addons {
			if addon.Data != "" {
				if len(addon.Config) > 0 || len(addon.Containers) > 0 {
					return errors.New("Config and containers should be empty when addon.Data is specified")
				}
				if _, err := base64.StdEncoding.DecodeString(addon.Data); err != nil {
					return errors.Errorf("Addon %s's data should be base64 encoded", addon.Name)
				}
			}

			if addon.Mode != "" {
				if addon.Mode != AddonModeEnsureExists && addon.Mode != AddonModeReconcile {
					return errors.Errorf("addon %s has a mode configuration '%s', must be either %s or %s", addon.Name, addon.Mode, AddonModeEnsureExists, AddonModeReconcile)
				}
			}

			// Validation for addons if they are enabled
			if to.Bool(addon.Enabled) {
				switch addon.Name {
				case "cluster-autoscaler":
					if isAvailabilitySets {
						return errors.Errorf("cluster-autoscaler addon can only be used with VirtualMachineScaleSets. Please specify \"availabilityProfile\": \"%s\"", VirtualMachineScaleSets)
					}
					for _, pool := range addon.Pools {
						if pool.Name == "" {
							return errors.Errorf("cluster-autoscaler addon pools configuration must have a 'name' property that correlates with a pool name in the agentPoolProfiles array")
						}
						if a.GetAgentPoolByName(pool.Name) == nil {
							return errors.Errorf("cluster-autoscaler addon pool 'name' %s does not match any agentPoolProfiles nodepool name", pool.Name)
						}
						if pool.Config != nil {
							var min, max int
							var err error
							if pool.Config["min-nodes"] != "" {
								min, err = strconv.Atoi(pool.Config["min-nodes"])
								if err != nil {
									return errors.Errorf("cluster-autoscaler addon pool 'name' %s has invalid 'min-nodes' config, must be a string int, got %s", pool.Name, pool.Config["min-nodes"])
								}
							}
							if pool.Config["max-nodes"] != "" {
								max, err = strconv.Atoi(pool.Config["max-nodes"])
								if err != nil {
									return errors.Errorf("cluster-autoscaler addon pool 'name' %s has invalid 'max-nodes' config, must be a string int, got %s", pool.Name, pool.Config["max-nodes"])
								}
							}
							if min > max {
								return errors.Errorf("cluster-autoscaler addon pool 'name' %s has invalid config, 'max-nodes' %d must be greater than or equal to 'min-nodes' %d", pool.Name, max, min)
							}
						}
					}
				case "nvidia-device-plugin":
					isValidVersion, err := common.IsValidMinVersion(a.OrchestratorProfile.OrchestratorType, a.OrchestratorProfile.OrchestratorRelease, a.OrchestratorProfile.OrchestratorVersion, "1.10.0")
					if err != nil {
						return err
					}
					if IsNSeriesSKU && !isValidVersion {
						return errors.New("NVIDIA Device Plugin add-on can only be used Kubernetes 1.10 or above. Please specify \"orchestratorRelease\": \"1.10\"")
					}
					if a.HasCoreOS() {
						return errors.New("NVIDIA Device Plugin add-on not currently supported on coreos. Please use node pools with Ubuntu only")
					}
				case "aad":
					if !a.HasAADAdminGroupID() {
						return errors.New("aad addon can't be enabled without a valid aadProfile w/ adminGroupID")
					}
				case "blobfuse-flexvolume":
					if a.HasCoreOS() {
						return errors.New("flexvolume add-ons not currently supported on coreos distro. Please use Ubuntu")
					}
				case "smb-flexvolume":
					if a.HasCoreOS() {
						return errors.New("flexvolume add-ons not currently supported on coreos distro. Please use Ubuntu")
					}
				case "keyvault-flexvolume":
					if a.HasCoreOS() {
						return errors.New("flexvolume add-ons not currently supported on coreos distro. Please use Ubuntu")
					}
				case "appgw-ingress":
					if (a.ServicePrincipalProfile == nil || len(a.ServicePrincipalProfile.ObjectID) == 0) &&
						!a.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
						return errors.New("appgw-ingress add-ons requires 'objectID' to be specified or UseManagedIdentity to be true")
					}

					if a.OrchestratorProfile.KubernetesConfig.NetworkPlugin != "azure" {
						return errors.New("appgw-ingress add-ons can only be used with Network Plugin as 'azure'")
					}

					if len(addon.Config["appgw-subnet"]) == 0 {
						return errors.New("appgw-ingress add-ons requires 'appgw-subnet' in the Config. It is used to provision the subnet for Application Gateway in the vnet")
					}
				case common.AzureDiskCSIDriverAddonName, common.AzureFileCSIDriverAddonName:
					if !to.Bool(a.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager) {
						return errors.New(fmt.Sprintf("%s add-on requires useCloudControllerManager to be true", addon.Name))
					}
				case "cloud-node-manager":
					if !common.IsKubernetesVersionGe(a.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
						return errors.New(fmt.Sprintf("%s add-on can only be used Kubernetes 1.16 or above", addon.Name))
					}
					if !to.Bool(a.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager) {
						return errors.New(fmt.Sprintf("%s add-on requires useCloudControllerManager to be true", addon.Name))
					}
				case common.CiliumAddonName:
					if !common.IsKubernetesVersionGe(a.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
						if a.OrchestratorProfile.KubernetesConfig.NetworkPolicy != NetworkPolicyCilium {
							return errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.CiliumAddonName, NetworkPolicyCilium)
						}
					} else {
						return errors.Errorf("%s addon is not supported on Kubernetes v1.16.0 or greater", common.CiliumAddonName)
					}
				case common.AntreaAddonName:
					if a.OrchestratorProfile.KubernetesConfig.NetworkPolicy != NetworkPolicyAntrea {
						return errors.Errorf("%s addon may only be enabled if the networkPolicy=%s", common.AntreaAddonName, NetworkPolicyAntrea)
					}
				case common.FlannelAddonName:
					if a.OrchestratorProfile.KubernetesConfig.NetworkPolicy != "" {
						return errors.Errorf("%s addon does not support NetworkPolicy, replace %s with \"\"", common.FlannelAddonName, a.OrchestratorProfile.KubernetesConfig.NetworkPolicy)
					}
					networkPlugin := a.OrchestratorProfile.KubernetesConfig.NetworkPlugin
					if networkPlugin != "" {
						if networkPlugin != NetworkPluginFlannel {
							return errors.Errorf("%s addon is not supported with networkPlugin=%s, please use networkPlugin=%s", common.FlannelAddonName, networkPlugin, NetworkPluginFlannel)
						}
					}
				case "azure-policy":
					isValidVersion, err := common.IsValidMinVersion(a.OrchestratorProfile.OrchestratorType, a.OrchestratorProfile.OrchestratorRelease, a.OrchestratorProfile.OrchestratorVersion, "1.10.0")
					if err != nil {
						return err
					}
					if !isValidVersion {
						return errors.New("Azure Policy add-on can only be used with Kubernetes v1.10 and above. Please specify a compatible version")
					}
					if a.ServicePrincipalProfile == nil || a.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
						return errors.New("Azure Policy add-on requires service principal profile to be specified")
					}
				case "kube-dns":
					kubeDNSEnabled = true
				case common.CoreDNSAddonName:
					corednsEnabled = true
				}
			} else {
				// Validation for addons if they are disabled
				switch addon.Name {
				case "cloud-node-manager":
					if to.Bool(a.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager) &&
						common.IsKubernetesVersionGe(a.OrchestratorProfile.OrchestratorVersion, "1.16.0") {
						return errors.New(fmt.Sprintf("%s add-on is required when useCloudControllerManager is true in Kubernetes 1.16 or above", addon.Name))
					}
				case common.AzureCloudProviderAddonName:
					return errors.Errorf("%s add-on is required, it cannot be disabled", addon.Name)
				}
			}
		}
		if kubeDNSEnabled && corednsEnabled {
			return errors.New("Both kube-dns and coredns addons are enabled, only one of these may be enabled on a cluster")
		}
	}
	return nil
}

func (a *Properties) validateExtensions() error {
	for _, agentPool := range a.AgentPoolProfiles {
		if len(agentPool.Extensions) != 0 && (len(agentPool.AvailabilityProfile) == 0 || agentPool.IsVirtualMachineScaleSets()) {
			return errors.Errorf("Extensions are currently not supported with VirtualMachineScaleSets. Please specify \"availabilityProfile\": \"%s\"", AvailabilitySet)
		}

		if agentPool.OSType == Windows && len(agentPool.Extensions) != 0 {
			for _, e := range agentPool.Extensions {
				if e.Name == "prometheus-grafana-k8s" {
					return errors.Errorf("prometheus-grafana-k8s extension is currently not supported for Windows agents")
				}
			}
		}
	}

	for _, extension := range a.ExtensionProfiles {
		if extension.ExtensionParametersKeyVaultRef != nil {
			if e := validate.Var(extension.ExtensionParametersKeyVaultRef.VaultID, "required"); e != nil {
				return errors.Errorf("the Keyvault ID must be specified for Extension %s", extension.Name)
			}
			if e := validate.Var(extension.ExtensionParametersKeyVaultRef.SecretName, "required"); e != nil {
				return errors.Errorf("the Keyvault Secret must be specified for Extension %s", extension.Name)
			}
			if !keyvaultIDRegex.MatchString(extension.ExtensionParametersKeyVaultRef.VaultID) {
				return errors.Errorf("Extension %s's keyvault secret reference is of incorrect format", extension.Name)
			}
		}
	}
	return nil
}

func (a *Properties) validateVNET() error {
	isCustomVNET := a.MasterProfile.IsCustomVNET()
	for _, agentPool := range a.AgentPoolProfiles {
		if agentPool.IsCustomVNET() != isCustomVNET {
			return errors.New("Multiple VNET Subnet configurations specified.  The master profile and each agent pool profile must all specify a custom VNET Subnet, or none at all")
		}
	}
	if isCustomVNET {
		if a.MasterProfile.IsVirtualMachineScaleSets() && a.MasterProfile.AgentVnetSubnetID == "" {
			return errors.New("when master profile is using VirtualMachineScaleSets and is custom vnet, set \"vnetsubnetid\" and \"agentVnetSubnetID\" for master profile")
		}

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
		if masterFirstIP == nil && !a.MasterProfile.IsVirtualMachineScaleSets() {
			return errors.Errorf("MasterProfile.FirstConsecutiveStaticIP (with VNET Subnet specification) '%s' is an invalid IP address", a.MasterProfile.FirstConsecutiveStaticIP)
		}

		if a.MasterProfile.VnetCidr != "" {
			_, _, err := net.ParseCIDR(a.MasterProfile.VnetCidr)
			if err != nil {
				return errors.Errorf("MasterProfile.VnetCidr '%s' contains invalid cidr notation", a.MasterProfile.VnetCidr)
			}
		}
	}
	return nil
}

func (a *Properties) validateServicePrincipalProfile() error {
	if a.OrchestratorProfile.OrchestratorType == Kubernetes {
		useManagedIdentity := a.OrchestratorProfile.KubernetesConfig != nil &&
			a.OrchestratorProfile.KubernetesConfig.UseManagedIdentity

		if !useManagedIdentity {
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

			if a.OrchestratorProfile.KubernetesConfig != nil && to.Bool(a.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms) && len(a.ServicePrincipalProfile.ObjectID) == 0 {
				return errors.Errorf("the service principal object ID must be specified with Orchestrator %s when enableEncryptionWithExternalKms is true", a.OrchestratorProfile.OrchestratorType)
			}

			if a.ServicePrincipalProfile.KeyvaultSecretRef != nil {
				if e := validate.Var(a.ServicePrincipalProfile.KeyvaultSecretRef.VaultID, "required"); e != nil {
					return errors.Errorf("the Keyvault ID must be specified for the Service Principle with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
				}
				if e := validate.Var(a.ServicePrincipalProfile.KeyvaultSecretRef.SecretName, "required"); e != nil {
					return errors.Errorf("the Keyvault Secret must be specified for the Service Principle with Orchestrator %s", a.OrchestratorProfile.OrchestratorType)
				}
				if !keyvaultIDRegex.MatchString(a.ServicePrincipalProfile.KeyvaultSecretRef.VaultID) {
					return errors.Errorf("service principal client keyvault secret reference is of incorrect format")
				}
			}
		}
	}
	return nil
}

func (a *Properties) validateManagedIdentity() error {
	if a.OrchestratorProfile.OrchestratorType == Kubernetes {
		useManagedIdentity := a.OrchestratorProfile.KubernetesConfig != nil &&
			a.OrchestratorProfile.KubernetesConfig.UseManagedIdentity

		if useManagedIdentity {
			version := common.RationalizeReleaseAndVersion(
				a.OrchestratorProfile.OrchestratorType,
				a.OrchestratorProfile.OrchestratorRelease,
				a.OrchestratorProfile.OrchestratorVersion,
				false,
				false)
			if version == "" {
				return errors.Errorf("the following user supplied OrchestratorProfile configuration is not supported: OrchestratorType: %s, OrchestratorRelease: %s, OrchestratorVersion: %s. Please check supported Release or Version for this build of aks-engine", a.OrchestratorProfile.OrchestratorType, a.OrchestratorProfile.OrchestratorRelease, a.OrchestratorProfile.OrchestratorVersion)
			}
			sv, err := semver.Make(version)
			if err != nil {
				return errors.Errorf("could not validate version %s", version)
			}
			minVersion, err := semver.Make("1.12.0")
			if err != nil {
				return errors.New("could not validate version")
			}

			if a.MasterProfile.IsVirtualMachineScaleSets() {
				if sv.LT(minVersion) {
					return errors.New("managed identity and VMSS masters can only be used with Kubernetes 1.12.0 or above. Please specify \"orchestratorRelease\": \"1.12\"")
				}
			} else if a.OrchestratorProfile.KubernetesConfig.UserAssignedID != "" && sv.LT(minVersion) {
				return errors.New("user assigned identity can only be used with Kubernetes 1.12.0 or above. Please specify \"orchestratorRelease\": \"1.12\"")
			}

		}
	}
	return nil
}

func (a *Properties) validateAADProfile() error {
	if profile := a.AADProfile; profile != nil {
		if a.OrchestratorProfile.OrchestratorType != Kubernetes {
			return errors.Errorf("'aadProfile' is only supported by orchestrator '%v'", Kubernetes)
		}
		if _, err := uuid.Parse(profile.ClientAppID); err != nil {
			return errors.Errorf("clientAppID '%v' is invalid", profile.ClientAppID)
		}
		if _, err := uuid.Parse(profile.ServerAppID); err != nil {
			return errors.Errorf("serverAppID '%v' is invalid", profile.ServerAppID)
		}
		if len(profile.TenantID) > 0 {
			if _, err := uuid.Parse(profile.TenantID); err != nil {
				return errors.Errorf("tenantID '%v' is invalid", profile.TenantID)
			}
		}
		if len(profile.AdminGroupID) > 0 {
			if _, err := uuid.Parse(profile.AdminGroupID); err != nil {
				return errors.Errorf("adminGroupID '%v' is invalid", profile.AdminGroupID)
			}
		}
	}
	return nil
}

func (a *AgentPoolProfile) validateAvailabilityProfile() error {
	switch a.AvailabilityProfile {
	case AvailabilitySet:
	case VirtualMachineScaleSets:
	case "":
	default:
		{
			return errors.Errorf("unknown availability profile type '%s' for agent pool '%s'.  Specify either %s, or %s", a.AvailabilityProfile, a.Name, AvailabilitySet, VirtualMachineScaleSets)
		}
	}

	return nil
}

func (a *AgentPoolProfile) validateRoles(orchestratorType string) error {
	validRoles := []AgentPoolProfileRole{AgentPoolProfileRoleEmpty}
	var found bool
	for _, validRole := range validRoles {
		if a.Role == validRole {
			found = true
			break
		}
	}
	if !found {
		return errors.Errorf("Role %q is not supported for Orchestrator %s", a.Role, orchestratorType)
	}
	return nil
}

func (a *AgentPoolProfile) validateStorageProfile(orchestratorType string) error {
	/* this switch statement is left to protect newly added orchestrators until they support Managed Disks*/
	if a.StorageProfile == ManagedDisks {
		switch orchestratorType {
		case DCOS:
		case Swarm:
		case Kubernetes:
		case SwarmMode:
		default:
			return errors.Errorf("HA volumes are currently unsupported for Orchestrator %s", orchestratorType)
		}
	}

	if a.StorageProfile == Ephemeral {
		switch orchestratorType {
		case Kubernetes:
			break
		case DCOS:
		case Swarm:
		case SwarmMode:
		default:
			return errors.Errorf("Ephemeral volumes are currently unsupported for Orchestrator %s", orchestratorType)
		}
	}

	return nil
}

func (a *AgentPoolProfile) validateCustomNodeLabels(orchestratorType string) error {
	if len(a.CustomNodeLabels) > 0 {
		switch orchestratorType {
		case DCOS:
		case Kubernetes:
			for k, v := range a.CustomNodeLabels {
				if e := validateKubernetesLabelKey(k); e != nil {
					return e
				}
				if e := validateKubernetesLabelValue(v); e != nil {
					return e
				}
			}
		default:
			return errors.New("Agent CustomNodeLabels are only supported for DCOS and Kubernetes")
		}
	}
	return nil
}

func validateVMSS(o *OrchestratorProfile, isUpdate bool, storageProfile string) error {
	if o.OrchestratorType == Kubernetes {
		version := common.RationalizeReleaseAndVersion(
			o.OrchestratorType,
			o.OrchestratorRelease,
			o.OrchestratorVersion,
			isUpdate,
			false)
		if version == "" {
			return errors.Errorf("the following OrchestratorProfile configuration is not supported: OrchestratorType: %s, OrchestratorRelease: %s, OrchestratorVersion: %s. Please check supported Release or Version for this build of aks-engine", o.OrchestratorType, o.OrchestratorRelease, o.OrchestratorVersion)
		}

		sv, err := semver.Make(version)
		if err != nil {
			return errors.Errorf("could not validate version %s", version)
		}
		minVersion, err := semver.Make("1.10.0")
		if err != nil {
			return errors.New("could not validate version")
		}
		if sv.LT(minVersion) {
			return errors.Errorf("VirtualMachineScaleSets are only available in Kubernetes version %s or greater. Please set \"orchestratorVersion\" to %s or above", minVersion.String(), minVersion.String())
		}
		// validation for instanceMetadata using VMSS with Kubernetes
		minVersion, err = semver.Make("1.10.2")
		if err != nil {
			return errors.New("could not validate version")
		}
		if o.KubernetesConfig != nil && o.KubernetesConfig.UseInstanceMetadata != nil {
			if *o.KubernetesConfig.UseInstanceMetadata && sv.LT(minVersion) {
				return errors.Errorf("VirtualMachineScaleSets with instance metadata is supported for Kubernetes version %s or greater. Please set \"useInstanceMetadata\": false in \"kubernetesConfig\" or set \"orchestratorVersion\" to %s or above", minVersion.String(), minVersion.String())
			}
		}
		if storageProfile == StorageAccount {
			return errors.Errorf("VirtualMachineScaleSets does not support %s disks.  Please specify \"storageProfile\": \"%s\" (recommended) or \"availabilityProfile\": \"%s\"", StorageAccount, ManagedDisks, AvailabilitySet)
		}
	}
	return nil
}

func (a *Properties) validateWindowsProfile() error {
	hasWindowsAgentPools := false
	for _, profile := range a.AgentPoolProfiles {
		if profile.OSType == Windows {
			hasWindowsAgentPools = true
			break
		}
	}

	if !hasWindowsAgentPools {
		return nil
	}

	o := a.OrchestratorProfile
	version := ""
	// This logic is broken because golang cases do not fallthrough by default.
	// I am leaving this in because I cannot get a clear answer on if we need to continue supporting Swarm + Windows and
	// RationalizeReleaseAndVersion does not properly handle Swarm.
	switch o.OrchestratorType {
	case DCOS:
	case Swarm:
	case SwarmMode:
	case Kubernetes:
		version = common.RationalizeReleaseAndVersion(
			o.OrchestratorType,
			o.OrchestratorRelease,
			o.OrchestratorVersion,
			false,
			true)

		if version == "" {
			return errors.Errorf("Orchestrator %s version %s does not support Windows", o.OrchestratorType, o.OrchestratorVersion)
		}
	default:
		return errors.Errorf("Orchestrator %v does not support Windows", o.OrchestratorType)
	}

	w := a.WindowsProfile
	if w == nil {
		return errors.New("WindowsProfile is required when the cluster definition contains Windows agent pools")
	}
	if e := validate.Var(w.AdminUsername, "required"); e != nil {
		return errors.New("WindowsProfile.AdminUsername is required, when agent pool specifies Windows")
	}
	if e := validate.Var(w.AdminPassword, "required"); e != nil {
		return errors.New("WindowsProfile.AdminPassword is required, when agent pool specifies Windows")
	}
	if !validatePasswordComplexity(w.AdminUsername, w.AdminPassword) {
		return errors.New("WindowsProfile.AdminPassword complexity not met. Windows password should contain 3 of the following categories - uppercase letters(A-Z), lowercase(a-z) letters, digits(0-9), special characters (~!@#$%^&*_-+=`|\\(){}[]:;<>,.?/')")
	}
	if e := validateKeyVaultSecrets(w.Secrets, true); e != nil {
		return e
	}
	if e := validateCsiProxyWindowsProperties(w, version); e != nil {
		return e
	}

	return nil
}

func validateCsiProxyWindowsProperties(w *WindowsProfile, k8sVersion string) error {
	if w.IsCSIProxyEnabled() {
		k8sSemVer, err := semver.Make(k8sVersion)
		if err != nil {
			return errors.Errorf("could not validate orchestrator version %s", k8sVersion)
		}
		minSemVer, err := semver.Make("1.18.0-beta.1")
		if err != nil {
			return errors.New("could not validate orchestrator version 1.18.0")
		}

		if k8sSemVer.LT(minSemVer) {
			return errors.New("CSI proxy for Windows is only available in Kubernetes versions 1.18.0 or greater")
		}

		if len(w.CSIProxyURL) == 0 {
			return errors.New("windowsProfile.csiProxyURL must be specified if enableCSIProxy is set")
		}
	}
	return nil
}

func (a *AgentPoolProfile) validateOrchestratorSpecificProperties(orchestratorType string) error {

	// for Kubernetes, we don't support AgentPoolProfile.DNSPrefix
	if orchestratorType == Kubernetes {
		if e := validate.Var(a.DNSPrefix, "len=0"); e != nil {
			return errors.New("AgentPoolProfile.DNSPrefix must be empty for Kubernetes")
		}
		if e := validate.Var(a.Ports, "len=0"); e != nil {
			return errors.New("AgentPoolProfile.Ports must be empty for Kubernetes")
		}
		if validate.Var(a.ScaleSetPriority, "eq=Regular") == nil && validate.Var(a.ScaleSetEvictionPolicy, "len=0") != nil {
			return errors.New("property 'AgentPoolProfile.ScaleSetEvictionPolicy' must be empty for AgentPoolProfile.Priority of Regular")
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
	} else if e := validate.Var(a.Ports, "len=0"); e != nil {
		return errors.Errorf("AgentPoolProfile.Ports must be empty when AgentPoolProfile.DNSPrefix is empty for Orchestrator: %s", orchestratorType)
	}

	if len(a.DiskSizesGB) > 0 {
		if e := validate.Var(a.StorageProfile, "eq=StorageAccount|eq=ManagedDisks"); e != nil {
			return errors.Errorf("property 'StorageProfile' must be set to either '%s' or '%s' when attaching disks", StorageAccount, ManagedDisks)
		}
		if e := validate.Var(a.AvailabilityProfile, "eq=VirtualMachineScaleSets|eq=AvailabilitySet"); e != nil {
			return errors.Errorf("property 'AvailabilityProfile' must be set to either '%s' or '%s' when attaching disks", VirtualMachineScaleSets, AvailabilitySet)
		}
		if a.StorageProfile == StorageAccount && (a.AvailabilityProfile == VirtualMachineScaleSets) {
			return errors.Errorf("VirtualMachineScaleSets does not support storage account attached disks.  Instead specify 'StorageAccount': '%s' or specify AvailabilityProfile '%s'", ManagedDisks, AvailabilitySet)
		}
	}

	if a.DiskEncryptionSetID != "" {
		if !diskEncryptionSetIDRegex.MatchString(a.DiskEncryptionSetID) {
			return errors.Errorf("DiskEncryptionSetID(%s) is of incorrect format, correct format: %s", a.DiskEncryptionSetID, diskEncryptionSetIDRegex.String())
		}
	}
	return nil
}

func (a *AgentPoolProfile) validateLoadBalancerBackendAddressPoolIDs() error {

	if a.LoadBalancerBackendAddressPoolIDs != nil {
		for _, backendPoolID := range a.LoadBalancerBackendAddressPoolIDs {
			if len(backendPoolID) == 0 {
				return errors.Errorf("AgentPoolProfile.LoadBalancerBackendAddressPoolIDs can not contain empty string. Agent pool name: %s", a.Name)
			}
		}
	}

	return nil
}

func validateProximityPlacementGroupID(ppgID string) error {
	if ppgID != "" {
		if !proximityPlacementGroupIDRegex.MatchString(ppgID) {
			return errors.Errorf("ProximityPlacementGroupID(%s) is of incorrect format, correct format: %s", ppgID, proximityPlacementGroupIDRegex.String())
		}
	}
	return nil
}

func validateKeyVaultSecrets(secrets []KeyVaultSecrets, requireCertificateStore bool) error {
	for _, s := range secrets {
		if len(s.VaultCertificates) == 0 {
			return errors.New("Valid KeyVaultSecrets must have no empty VaultCertificates")
		}
		if s.SourceVault == nil {
			return errors.New("missing SourceVault in KeyVaultSecrets")
		}
		if s.SourceVault.ID == "" {
			return errors.New("KeyVaultSecrets must have a SourceVault.ID")
		}
		for _, c := range s.VaultCertificates {
			if _, e := url.Parse(c.CertificateURL); e != nil {
				return errors.Errorf("Certificate url was invalid. received error %s", e)
			}
			if e := validateName(c.CertificateStore, "KeyVaultCertificate.CertificateStore"); requireCertificateStore && e != nil {
				return errors.Errorf("%s for certificates in a WindowsProfile", e)
			}
		}
	}
	return nil
}

// Validate ensures that the WindowsProfile is valid
func (w *WindowsProfile) Validate(orchestratorType string) error {
	if w.WindowsImageSourceURL != "" {
		if orchestratorType != DCOS && orchestratorType != Kubernetes {
			return errors.New("Windows Custom Images are only supported if the Orchestrator Type is DCOS or Kubernetes")
		}
	}
	if e := validate.Var(w.AdminUsername, "required"); e != nil {
		return errors.New("WindowsProfile.AdminUsername is required, when agent pool specifies windows")
	}
	if e := validate.Var(w.AdminPassword, "required"); e != nil {
		return errors.New("WindowsProfile.AdminPassword is required, when agent pool specifies windows")
	}
	if !validatePasswordComplexity(w.AdminUsername, w.AdminPassword) {
		return errors.New("WindowsProfile.AdminPassword complexity not met. Windows password should contain 3 of the following categories - uppercase letters(A-Z), lowercase(a-z) letters, digits(0-9), special characters (~!@#$%^&*_-+=`|\\(){}[]:;<>,.?/')")
	}
	return validateKeyVaultSecrets(w.Secrets, true)
}

func validatePasswordComplexity(name string, password string) (out bool) {

	if strings.EqualFold(name, password) {
		return false
	}

	if len(password) == 0 {
		return false
	}

	hits := 0
	if regexp.MustCompile(`[0-9]+`).MatchString(password) {
		hits++
	}
	if regexp.MustCompile(`[A-Z]+`).MatchString(password) {
		hits++
	}
	if regexp.MustCompile(`[a-z]`).MatchString(password) {
		hits++
	}
	if regexp.MustCompile(`[~!@#\$%\^&\*_\-\+=\x60\|\(\){}\[\]:;"'<>,\.\?/]+`).MatchString(password) {
		hits++
	}
	return hits > 2
}

// Validate validates the KubernetesConfig
func (k *KubernetesConfig) Validate(k8sVersion string, hasWindows, ipv6DualStackEnabled, isIPv6 bool) error {
	// number of minimum retries allowed for kubelet to post node status
	const minKubeletRetries = 4

	// enableIPv6DualStack and enableIPv6Only are mutually exclusive feature flags
	if ipv6DualStackEnabled && isIPv6 {
		return errors.Errorf("featureFlags.EnableIPv6DualStack and featureFlags.EnableIPv6Only can't be enabled at the same time.")
	}

	sv, err := semver.Make(k8sVersion)
	if err != nil {
		return errors.Errorf("could not validate version %s", k8sVersion)
	}

	if ipv6DualStackEnabled {
		minVersion, err := semver.Make("1.16.0")
		if err != nil {
			return errors.New("could not validate version")
		}
		if sv.LT(minVersion) {
			return errors.Errorf("IPv6 dual stack not available in kubernetes version %s", k8sVersion)
		}
		// ipv6 dual stack feature is currently only supported with kubenet
		if k.NetworkPlugin != "kubenet" {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.NetworkPlugin '%s' is invalid. IPv6 dual stack supported only with kubenet.", k.NetworkPlugin)
		}
	}

	if isIPv6 {
		minVersion, err := semver.Make("1.18.0")
		if err != nil {
			return errors.New("could not validate version")
		}
		if sv.LT(minVersion) {
			return errors.Errorf("IPv6 single stack not available in kubernetes version %s", k8sVersion)
		}
		// single stack IPv6 feature is currently only supported with kubenet
		if k.NetworkPlugin != "kubenet" {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.NetworkPlugin '%s' is invalid. IPv6 single stack supported only with kubenet.", k.NetworkPlugin)
		}
	}

	if k.ClusterSubnet != "" {
		clusterSubnets := strings.Split(k.ClusterSubnet, ",")
		if !ipv6DualStackEnabled && len(clusterSubnets) > 1 {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.ClusterSubnet '%s' is an invalid subnet", k.ClusterSubnet)
		}
		if ipv6DualStackEnabled && len(clusterSubnets) > 2 {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.ClusterSubnet '%s' is an invalid subnet. Not more than 2 subnets for ipv6 dual stack.", k.ClusterSubnet)
		}

		for _, clusterSubnet := range clusterSubnets {
			_, subnet, err := net.ParseCIDR(clusterSubnet)
			if err != nil {
				return errors.Errorf("OrchestratorProfile.KubernetesConfig.ClusterSubnet '%s' is an invalid subnet", clusterSubnet)
			}

			if k.NetworkPlugin == "azure" {
				ones, bits := subnet.Mask.Size()
				if bits-ones <= 8 {
					return errors.Errorf("OrchestratorProfile.KubernetesConfig.ClusterSubnet '%s' must reserve at least 9 bits for nodes", clusterSubnet)
				}
			}
		}
	}

	if k.DockerBridgeSubnet != "" {
		_, _, err := net.ParseCIDR(k.DockerBridgeSubnet)
		if err != nil {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.DockerBridgeSubnet '%s' is an invalid subnet", k.DockerBridgeSubnet)
		}
	}

	if k.MaxPods != 0 {
		if k.MaxPods < KubernetesMinMaxPods {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.MaxPods '%v' must be at least %v", k.MaxPods, KubernetesMinMaxPods)
		}
	}

	if k.KubeletConfig != nil {
		if _, ok := k.KubeletConfig["--node-status-update-frequency"]; ok {
			val := k.KubeletConfig["--node-status-update-frequency"]
			_, err := time.ParseDuration(val)
			if err != nil {
				return errors.Errorf("--node-status-update-frequency '%s' is not a valid duration", val)
			}
		}
	}

	if _, ok := k.ControllerManagerConfig["--node-monitor-grace-period"]; ok {
		_, err := time.ParseDuration(k.ControllerManagerConfig["--node-monitor-grace-period"])
		if err != nil {
			return errors.Errorf("--node-monitor-grace-period '%s' is not a valid duration", k.ControllerManagerConfig["--node-monitor-grace-period"])
		}
	}

	if k.KubeletConfig != nil {
		if _, ok := k.KubeletConfig["--node-status-update-frequency"]; ok {
			if _, ok := k.ControllerManagerConfig["--node-monitor-grace-period"]; ok {
				nodeStatusUpdateFrequency, _ := time.ParseDuration(k.KubeletConfig["--node-status-update-frequency"])
				ctrlMgrNodeMonitorGracePeriod, _ := time.ParseDuration(k.ControllerManagerConfig["--node-monitor-grace-period"])
				kubeletRetries := ctrlMgrNodeMonitorGracePeriod.Seconds() / nodeStatusUpdateFrequency.Seconds()
				if kubeletRetries < minKubeletRetries {
					return errors.Errorf("aks-engine requires that --node-monitor-grace-period(%f)s be larger than nodeStatusUpdateFrequency(%f)s by at least a factor of %d; ", ctrlMgrNodeMonitorGracePeriod.Seconds(), nodeStatusUpdateFrequency.Seconds(), minKubeletRetries)
				}
			}
		}
		// Re-enable this unit test if --non-masquerade-cidr is re-introduced
		/*if _, ok := k.KubeletConfig["--non-masquerade-cidr"]; ok {
			if _, _, err := net.ParseCIDR(k.KubeletConfig["--non-masquerade-cidr"]); err != nil {
				return errors.Errorf("--non-masquerade-cidr kubelet config '%s' is an invalid CIDR string", k.KubeletConfig["--non-masquerade-cidr"])
			}
		}*/
	}

	if _, ok := k.ControllerManagerConfig["--pod-eviction-timeout"]; ok {
		_, err := time.ParseDuration(k.ControllerManagerConfig["--pod-eviction-timeout"])
		if err != nil {
			return errors.Errorf("--pod-eviction-timeout '%s' is not a valid duration", k.ControllerManagerConfig["--pod-eviction-timeout"])
		}
	}

	if _, ok := k.ControllerManagerConfig["--route-reconciliation-period"]; ok {
		_, err := time.ParseDuration(k.ControllerManagerConfig["--route-reconciliation-period"])
		if err != nil {
			return errors.Errorf("--route-reconciliation-period '%s' is not a valid duration", k.ControllerManagerConfig["--route-reconciliation-period"])
		}
	}

	if k.DNSServiceIP != "" || k.ServiceCidr != "" {
		if k.DNSServiceIP == "" {
			return errors.New("OrchestratorProfile.KubernetesConfig.DNSServiceIP must be specified when ServiceCidr is")
		}
		if k.ServiceCidr == "" {
			return errors.New("OrchestratorProfile.KubernetesConfig.ServiceCidr must be specified when DNSServiceIP is")
		}

		dnsIP := net.ParseIP(k.DNSServiceIP)
		if dnsIP == nil {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.DNSServiceIP '%s' is an invalid IP address", k.DNSServiceIP)
		}

		primaryServiceCIDR := k.ServiceCidr
		if ipv6DualStackEnabled {
			// split the service cidr to see if there are multiple cidrs
			serviceCidrs := strings.Split(k.ServiceCidr, ",")
			if len(serviceCidrs) > 2 {
				return errors.Errorf("OrchestratorProfile.KubernetesConfig.ServiceCidr '%s' is an invalid CIDR subnet. More than 2 CIDRs not allowed for dualstack", k.ServiceCidr)
			}
			if len(serviceCidrs) == 2 {
				firstServiceCIDR, secondServiceCIDR := serviceCidrs[0], serviceCidrs[1]
				_, _, err := net.ParseCIDR(secondServiceCIDR)
				if err != nil {
					return errors.Errorf("OrchestratorProfile.KubernetesConfig.ServiceCidr '%s' is an invalid CIDR subnet", secondServiceCIDR)
				}
				// use the primary service cidr for further validation
				primaryServiceCIDR = firstServiceCIDR
			}
			// if # of service cidrs is 1, then continues with the default validation
		}

		_, serviceCidr, err := net.ParseCIDR(primaryServiceCIDR)
		if err != nil {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.ServiceCidr '%s' is an invalid CIDR subnet", primaryServiceCIDR)
		}

		// Finally validate that the DNS ip is within the subnet
		if !serviceCidr.Contains(dnsIP) {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.DNSServiceIP '%s' is not within the ServiceCidr '%s'", k.DNSServiceIP, primaryServiceCIDR)
		}

		// and that the DNS IP is _not_ the subnet broadcast address
		broadcast := common.IP4BroadcastAddress(serviceCidr)
		if dnsIP.Equal(broadcast) {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.DNSServiceIP '%s' cannot be the broadcast address of ServiceCidr '%s'", k.DNSServiceIP, primaryServiceCIDR)
		}

		// and that the DNS IP is _not_ the first IP in the service subnet
		firstServiceIP := common.CidrFirstIP(serviceCidr.IP)
		if firstServiceIP.Equal(dnsIP) {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.DNSServiceIP '%s' cannot be the first IP of ServiceCidr '%s'", k.DNSServiceIP, primaryServiceCIDR)
		}
	}

	if k.ProxyMode != "" && k.ProxyMode != KubeProxyModeIPTables && k.ProxyMode != KubeProxyModeIPVS {
		return errors.Errorf("Invalid KubeProxyMode %v. Allowed modes are %v and %v", k.ProxyMode, KubeProxyModeIPTables, KubeProxyModeIPVS)
	}

	// dualstack IPVS mode supported from 1.16+
	// dualstack IPtables mode supported from 1.18+
	if ipv6DualStackEnabled && k.ProxyMode == KubeProxyModeIPTables {
		minVersion, err := semver.Make("1.18.0")
		if err != nil {
			return errors.New("could not validate version")
		}
		if sv.LT(minVersion) {
			return errors.Errorf("KubeProxyMode %v in dualstack not supported with %s version", k.ProxyMode, k8sVersion)
		}
	}

	// Validate that we have a valid etcd version
	if e := validateEtcdVersion(k.EtcdVersion); e != nil {
		return e
	}

	// Validate containerd scenarios
	if k.ContainerRuntime == Docker || k.ContainerRuntime == "" {
		if k.ContainerdVersion != "" {
			return errors.Errorf("containerdVersion is only valid in a non-docker context, use %s or %s containerRuntime values instead if you wish to provide a containerdVersion", Containerd, KataContainers)
		}
	} else {
		if e := validateContainerdVersion(k.ContainerdVersion); e != nil {
			return e
		}
	}

	if k.UseCloudControllerManager != nil && *k.UseCloudControllerManager || k.CustomCcmImage != "" {
		sv, err := semver.Make(k8sVersion)
		if err != nil {
			return errors.Errorf("could not validate version %s", k8sVersion)
		}
		minVersion, err := semver.Make("1.8.0")
		if err != nil {
			return errors.New("could not validate version")
		}
		if sv.LT(minVersion) {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.UseCloudControllerManager and OrchestratorProfile.KubernetesConfig.CustomCcmImage not available in kubernetes version %s", k8sVersion)
		}
	}

	if e := k.validateNetworkPlugin(hasWindows); e != nil {
		return e
	}
	if e := k.validateNetworkPolicy(k8sVersion, hasWindows); e != nil {
		return e
	}
	if e := k.validateNetworkPluginPlusPolicy(); e != nil {
		return e
	}
	if e := k.validateNetworkMode(); e != nil {
		return e
	}
	return k.validateContainerRuntimeConfig()
}

func (k *KubernetesConfig) validateContainerRuntimeConfig() error {
	if val, ok := k.ContainerRuntimeConfig[common.ContainerDataDirKey]; ok {
		if val == "" {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.ContainerRuntimeConfig.DataDir '%s' is invalid: must not be empty", val)
		}
		if !strings.HasPrefix(val, "/") {
			return errors.Errorf("OrchestratorProfile.KubernetesConfig.ContainerRuntimeConfig.DataDir '%s' is invalid: must be absolute path", val)
		}
	}

	// Validate base config here, and only allow predefined mutations to ensure invariant.
	if k.ContainerRuntime == Containerd {
		_, err := common.GetContainerdConfig(k.ContainerRuntimeConfig, nil)
		if err != nil {
			return err
		}
	} else {
		_, err := common.GetDockerConfig(k.ContainerRuntimeConfig, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k *KubernetesConfig) validateNetworkPlugin(hasWindows bool) error {

	networkPlugin := k.NetworkPlugin

	// Check NetworkPlugin has a valid value.
	valid := false
	for _, plugin := range NetworkPluginValues {
		if networkPlugin == plugin {
			valid = true
			break
		}
	}
	if !valid {
		return errors.Errorf("unknown networkPlugin '%s' specified", networkPlugin)
	}

	// Temporary safety check, to be removed when Windows support is added.
	if (networkPlugin == NetworkPluginAntrea) && hasWindows {
		return errors.Errorf("networkPlugin '%s' is not supporting windows agents", networkPlugin)
	}

	return nil
}

func (k *KubernetesConfig) validateNetworkPolicy(k8sVersion string, hasWindows bool) error {

	networkPolicy := k.NetworkPolicy
	networkPlugin := k.NetworkPlugin

	// Check NetworkPolicy has a valid value.
	valid := false
	for _, plugin := range NetworkPolicyValues {
		if networkPolicy == plugin {
			valid = true
			break
		}
	}
	if !valid {
		return errors.Errorf("unknown networkPolicy '%s' specified", networkPolicy)
	}

	if networkPolicy == "azure" && networkPlugin == "azure" && !common.IsKubernetesVersionGe(k8sVersion, "1.8.0") {
		return errors.New("networkPolicy azure requires kubernetes version of 1.8 or higher")
	}

	// Temporary safety check, to be removed when Windows support is added.
	if (networkPolicy == "calico" || networkPolicy == NetworkPolicyCilium ||
		networkPolicy == NetworkPolicyAntrea) && hasWindows {
		return errors.Errorf("networkPolicy '%s' is not supporting windows agents", networkPolicy)
	}

	return nil
}

func (k *KubernetesConfig) validateNetworkPluginPlusPolicy() error {
	var config k8sNetworkConfig

	config.networkPlugin = k.NetworkPlugin
	config.networkPolicy = k.NetworkPolicy

	for _, c := range networkPluginPlusPolicyAllowed {
		if c.networkPlugin == config.networkPlugin && c.networkPolicy == config.networkPolicy {
			return nil
		}
	}
	return errors.Errorf("networkPolicy '%s' is not supported with networkPlugin '%s'", config.networkPolicy, config.networkPlugin)
}

func (k *KubernetesConfig) validateNetworkMode() error {
	networkPlugin := k.NetworkPlugin
	networkPolicy := k.NetworkPolicy
	networkMode := k.NetworkMode

	// Check NetworkMode has a valid value.
	valid := false
	for _, mode := range NetworkModeValues {
		if networkMode == mode {
			valid = true
			break
		}
	}
	if !valid {
		return errors.Errorf("unknown networkMode '%s' specified", networkMode)
	}

	if networkMode != "" {
		if networkPlugin != "azure" {
			return errors.New("networkMode requires network plugin to be 'azure'")
		}

		if networkPolicy == "calico" && networkMode != NetworkModeTransparent {
			return errors.Errorf("networkMode '%s' is not supported by calico", networkMode)
		}
	}

	return nil
}

func (k *KubernetesConfig) isUsingCustomKubeComponent() bool {
	return k.CustomKubeAPIServerImage != "" || k.CustomKubeControllerManagerImage != "" || k.CustomKubeProxyImage != "" || k.CustomKubeSchedulerImage != "" || k.CustomKubeBinaryURL != ""
}

func (a *Properties) validateContainerRuntime() error {
	var containerRuntime string

	switch a.OrchestratorProfile.OrchestratorType {
	case Kubernetes:
		if a.OrchestratorProfile.KubernetesConfig != nil {
			containerRuntime = a.OrchestratorProfile.KubernetesConfig.ContainerRuntime
		}
	default:
		return nil
	}

	// Check ContainerRuntime has a valid value.
	valid := false
	for _, runtime := range ContainerRuntimeValues {
		if containerRuntime == runtime {
			valid = true
			break
		}
	}
	if !valid {
		return errors.Errorf("unknown containerRuntime %q specified", containerRuntime)
	}

	// Make sure we don't use unsupported container runtimes on windows.
	if (containerRuntime == KataContainers || containerRuntime == Containerd) && a.HasWindows() {
		return errors.Errorf("containerRuntime %q is not supporting windows agents", containerRuntime)
	}

	return nil
}

func (a *Properties) validateCustomKubeComponent() error {
	k := a.OrchestratorProfile.KubernetesConfig
	if k == nil {
		return nil
	}

	if common.IsKubernetesVersionGe(a.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
		if k.CustomHyperkubeImage != "" {
			return errors.New("customHyperkubeImage has no effect in Kubernetes version 1.17.0 or above")
		}
	} else {
		if k.isUsingCustomKubeComponent() {
			return errors.New("customKubeAPIServerImage, customKubeControllerManagerImage, customKubeProxyImage, customKubeSchedulerImage or customKubeBinaryURL have no effect in Kubernetes version 1.16 or earlier")
		}
	}

	return nil
}

func (a *Properties) validatePrivateAzureRegistryServer() error {
	k := a.OrchestratorProfile.KubernetesConfig
	if k == nil || k.PrivateAzureRegistryServer == "" {
		return nil
	}

	// Custom components must be provided if private azure registry server is not empty
	if common.IsKubernetesVersionGe(a.OrchestratorProfile.OrchestratorVersion, "1.17.0") {
		if !k.isUsingCustomKubeComponent() {
			return errors.Errorf("customKubeAPIServerImage, customKubeControllerManagerImage, customKubeProxyImage or customKubeSchedulerImage must be provided when privateAzureRegistryServer is provided")
		}
	} else {
		if k.CustomHyperkubeImage == "" {
			return errors.Errorf("customHyperkubeImage must be provided when privateAzureRegistryServer is provided")
		}
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

func validatePoolOSType(os OSType) error {
	if os != Linux && os != Windows && os != "" {
		return errors.New("AgentPoolProfile.osType must be either Linux or Windows")
	}
	return nil
}

func validatePoolAcceleratedNetworking(vmSize string) error {
	if !helpers.AcceleratedNetworkingSupported(vmSize) {
		return errors.Errorf("AgentPoolProfile.vmsize %s does not support AgentPoolProfile.acceleratedNetworking", vmSize)
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

func validateKubernetesLabelValue(v string) error {
	if !(len(v) == 0) && !labelValueRegex.MatchString(v) {
		return errors.Errorf("Label value '%s' is invalid. Valid label values must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between", v)
	}
	return nil
}

func validateKubernetesLabelKey(k string) error {
	if !labelKeyRegex.MatchString(k) {
		return errors.Errorf("Label key '%s' is invalid. Valid label keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/)", k)
	}
	prefix := strings.Split(k, "/")
	if len(prefix) != 1 && len(prefix[0]) > labelKeyPrefixMaxLength {
		return errors.Errorf("Label key prefix '%s' is invalid. If specified, the prefix must be no longer than 253 characters in total", k)
	}
	return nil
}

func validateEtcdVersion(etcdVersion string) error {
	// "" is a valid etcdVersion that maps to DefaultEtcdVersion
	if etcdVersion == "" {
		return nil
	}
	for _, ver := range etcdValidVersions {
		if ver == etcdVersion {
			return nil
		}
	}
	return errors.Errorf("Invalid etcd version \"%s\", please use one of the following versions: %s", etcdVersion, etcdValidVersions)
}

func validateContainerdVersion(containerdVersion string) error {
	// "" is a valid containerd that maps to DefaultContainerdVersion
	if containerdVersion == "" {
		return nil
	}
	for _, ver := range containerdValidVersions {
		if ver == containerdVersion {
			return nil
		}
	}
	return errors.Errorf("Invalid containerd version \"%s\", please use one of the following versions: %s", containerdVersion, containerdValidVersions)
}

// Check that distro has a valid value
func validateDistro(distro Distro, distroValues []Distro) bool {
	for _, d := range distroValues {
		if distro == d {
			return true
		}
	}
	return false
}

func (i *ImageReference) validateImageNameAndGroup() error {
	if i.Name == "" && i.ResourceGroup != "" {
		return errors.New("imageName needs to be specified when imageResourceGroup is provided")
	}
	if i.Name != "" && i.ResourceGroup == "" {
		return errors.New("imageResourceGroup needs to be specified when imageName is provided")
	}
	return nil
}

func (cs *ContainerService) validateCustomCloudProfile() error {
	a := cs.Properties
	if a.CustomCloudProfile != nil {
		if a.CustomCloudProfile.PortalURL == "" {
			return errors.New("portalURL needs to be specified when CustomCloudProfile is provided")
		}
		if !strings.HasPrefix(a.CustomCloudProfile.PortalURL, fmt.Sprintf("https://portal.%s.", cs.Location)) {
			return fmt.Errorf("portalURL needs to start with https://portal.%s. ", cs.Location)
		}
		if a.CustomCloudProfile.AuthenticationMethod != "" && !(a.CustomCloudProfile.AuthenticationMethod == ClientSecretAuthMethod || a.CustomCloudProfile.AuthenticationMethod == ClientCertificateAuthMethod) {
			return errors.Errorf("authenticationMethod allowed values are '%s' and '%s'", ClientCertificateAuthMethod, ClientSecretAuthMethod)
		}
		if a.CustomCloudProfile.IdentitySystem != "" && !(a.CustomCloudProfile.IdentitySystem == AzureADIdentitySystem || a.CustomCloudProfile.IdentitySystem == ADFSIdentitySystem) {
			return errors.Errorf("identitySystem allowed values are '%s' and '%s'", AzureADIdentitySystem, ADFSIdentitySystem)
		}

		dependenciesLocationValues := DependenciesLocationValues
		if !validateDependenciesLocation(a.CustomCloudProfile.DependenciesLocation, dependenciesLocationValues) {
			return errors.Errorf("The %s dependenciesLocation is not supported. The supported vaules are %s", a.CustomCloudProfile.DependenciesLocation, dependenciesLocationValues)
		}
	}
	return nil
}

// Validate implements validation for ContainerService
func (cs *ContainerService) Validate(isUpdate bool) error {
	if e := cs.validateProperties(); e != nil {
		return e
	}
	if e := cs.validateLocation(); e != nil {
		return e
	}
	if e := cs.validateCustomCloudProfile(); e != nil {
		return e
	}
	if e := cs.Properties.validate(isUpdate); e != nil {
		return e
	}
	return nil
}

func (cs *ContainerService) validateLocation() error {
	if cs.Properties != nil && cs.Properties.IsAzureStackCloud() && cs.Location == "" {
		return errors.New("missing ContainerService Location")
	}
	return nil
}

func (cs *ContainerService) validateProperties() error {
	if cs.Properties == nil {
		return errors.New("missing ContainerService Properties")
	}
	return nil
}

// Check that dependenciesLocation has a valid value
func validateDependenciesLocation(dependenciesLocation DependenciesLocation, dependenciesLocationValues []DependenciesLocation) bool {
	for _, d := range dependenciesLocationValues {
		if dependenciesLocation == d {
			return true
		}
	}
	return false
}

// validateAzureStackSupport logs a warning if apimodel contains preview features and returns an error if a property is not supported on Azure Stack clouds
func (a *Properties) validateAzureStackSupport() error {
	if a.OrchestratorProfile.OrchestratorType == Kubernetes && a.IsAzureStackCloud() {
		networkPlugin := a.OrchestratorProfile.KubernetesConfig.NetworkPlugin
		if networkPlugin == "azure" || networkPlugin == "" {
			log.Warnf("NetworkPlugin 'azure' is a private preview feature on Azure Stack clouds")
		}
		if networkPlugin != "azure" && networkPlugin != "kubenet" && networkPlugin != "" {
			return errors.Errorf("kubernetesConfig.networkPlugin '%s' is not supported on Azure Stack clouds", networkPlugin)
		}
		if a.MasterProfile.AvailabilityProfile != AvailabilitySet {
			return errors.Errorf("masterProfile.availabilityProfile should be set to '%s' on Azure Stack clouds", AvailabilitySet)
		}
		for _, agentPool := range a.AgentPoolProfiles {
			pool := agentPool
			if pool.AvailabilityProfile != AvailabilitySet {
				return errors.Errorf("agentPoolProfiles[%s].availabilityProfile should be set to '%s' on Azure Stack clouds", pool.Name, AvailabilitySet)
			}
		}
	}
	return nil
}
