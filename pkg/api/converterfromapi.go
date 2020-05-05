// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/blang/semver"
)

///////////////////////////////////////////////////////////
// The converter exposes functions to convert the top level
// ContainerService resource
//
// All other functions are internal helper functions used
// for converting.
///////////////////////////////////////////////////////////

// ConvertContainerServiceToVLabs converts an unversioned ContainerService to a vlabs ContainerService
func ConvertContainerServiceToVLabs(api *ContainerService) *vlabs.ContainerService {
	vlabsCS := &vlabs.ContainerService{}
	vlabsCS.ID = api.ID
	vlabsCS.Location = api.Location
	vlabsCS.Name = api.Name
	if api.Plan != nil {
		vlabsCS.Plan = &vlabs.ResourcePurchasePlan{}
		convertResourcePurchasePlanToVLabs(api.Plan, vlabsCS.Plan)
	}
	vlabsCS.Tags = map[string]string{}
	for k, v := range api.Tags {
		vlabsCS.Tags[k] = v
	}
	vlabsCS.Type = api.Type
	vlabsCS.Properties = &vlabs.Properties{}
	convertPropertiesToVLabs(api.Properties, vlabsCS.Properties)
	return vlabsCS
}

// ConvertOrchestratorVersionProfileToVLabs converts an unversioned OrchestratorVersionProfile to a vlabs OrchestratorVersionProfile
func ConvertOrchestratorVersionProfileToVLabs(api *OrchestratorVersionProfile) *vlabs.OrchestratorVersionProfile {
	vlabsProfile := &vlabs.OrchestratorVersionProfile{}
	switch api.OrchestratorType {
	case Kubernetes:
		vlabsProfile.OrchestratorType = vlabs.Kubernetes
	case DCOS:
		vlabsProfile.OrchestratorType = vlabs.DCOS
	case Swarm:
		vlabsProfile.OrchestratorType = vlabs.Swarm
	case SwarmMode:
		vlabsProfile.OrchestratorType = vlabs.SwarmMode
	}
	vlabsProfile.OrchestratorVersion = api.OrchestratorVersion
	vlabsProfile.Default = api.Default
	if api.Upgrades != nil {
		vlabsProfile.Upgrades = make([]*vlabs.OrchestratorProfile, len(api.Upgrades))
		for i, h := range api.Upgrades {
			vlabsProfile.Upgrades[i] = &vlabs.OrchestratorProfile{
				OrchestratorVersion: h.OrchestratorVersion,
			}
		}
	}
	return vlabsProfile
}

// convertResourcePurchasePlanToVLabs converts a vlabs ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertResourcePurchasePlanToVLabs(api *ResourcePurchasePlan, vlabs *vlabs.ResourcePurchasePlan) {
	vlabs.Name = api.Name
	vlabs.Product = api.Product
	vlabs.PromotionCode = api.PromotionCode
	vlabs.Publisher = api.Publisher
}

func convertPropertiesToVLabs(api *Properties, vlabsProps *vlabs.Properties) {
	vlabsProps.ProvisioningState = vlabs.ProvisioningState(api.ProvisioningState)
	if api.OrchestratorProfile != nil {
		vlabsProps.OrchestratorProfile = &vlabs.OrchestratorProfile{}
		convertOrchestratorProfileToVLabs(api.OrchestratorProfile, vlabsProps.OrchestratorProfile)
	}
	if api.MasterProfile != nil {
		vlabsProps.MasterProfile = &vlabs.MasterProfile{}
		convertMasterProfileToVLabs(api.MasterProfile, vlabsProps.MasterProfile)
	}
	vlabsProps.AgentPoolProfiles = []*vlabs.AgentPoolProfile{}
	for _, apiProfile := range api.AgentPoolProfiles {
		vlabsProfile := &vlabs.AgentPoolProfile{}
		convertAgentPoolProfileToVLabs(apiProfile, vlabsProfile)
		vlabsProps.AgentPoolProfiles = append(vlabsProps.AgentPoolProfiles, vlabsProfile)
	}
	if api.LinuxProfile != nil {
		vlabsProps.LinuxProfile = &vlabs.LinuxProfile{}
		convertLinuxProfileToVLabs(api.LinuxProfile, vlabsProps.LinuxProfile)
	}
	vlabsProps.ExtensionProfiles = []*vlabs.ExtensionProfile{}
	for _, extensionProfile := range api.ExtensionProfiles {
		vlabsExtensionProfile := &vlabs.ExtensionProfile{}
		convertExtensionProfileToVLabs(extensionProfile, vlabsExtensionProfile)
		vlabsProps.ExtensionProfiles = append(vlabsProps.ExtensionProfiles, vlabsExtensionProfile)
	}
	if api.WindowsProfile != nil {
		vlabsProps.WindowsProfile = &vlabs.WindowsProfile{}
		convertWindowsProfileToVLabs(api.WindowsProfile, vlabsProps.WindowsProfile)
	}
	if api.ServicePrincipalProfile != nil {
		vlabsProps.ServicePrincipalProfile = &vlabs.ServicePrincipalProfile{}
		convertServicePrincipalProfileToVLabs(api.ServicePrincipalProfile, vlabsProps.ServicePrincipalProfile)
	}
	if api.CertificateProfile != nil {
		vlabsProps.CertificateProfile = &vlabs.CertificateProfile{}
		convertCertificateProfileToVLabs(api.CertificateProfile, vlabsProps.CertificateProfile)
	}
	if api.AADProfile != nil {
		vlabsProps.AADProfile = &vlabs.AADProfile{}
		convertAADProfileToVLabs(api.AADProfile, vlabsProps.AADProfile)
	}

	if api.FeatureFlags != nil {
		vlabsProps.FeatureFlags = &vlabs.FeatureFlags{}
		convertFeatureFlagsToVLabs(api.FeatureFlags, vlabsProps.FeatureFlags)
	}

	if api.CustomCloudProfile != nil {
		vlabsProps.CustomCloudProfile = &vlabs.CustomCloudProfile{}
		convertCloudProfileToVLabs(api.CustomCloudProfile, vlabsProps.CustomCloudProfile)
	}

	if api.TelemetryProfile != nil {
		vlabsProps.TelemetryProfile = &vlabs.TelemetryProfile{}
		convertTelemetryProfileToVLabs(api.TelemetryProfile, vlabsProps.TelemetryProfile)
	}
}

func convertExtensionProfileToVLabs(api *ExtensionProfile, obj *vlabs.ExtensionProfile) {
	obj.Name = api.Name
	obj.Version = api.Version
	obj.ExtensionParameters = api.ExtensionParameters
	if api.ExtensionParametersKeyVaultRef != nil {
		obj.ExtensionParametersKeyVaultRef = &vlabs.KeyvaultSecretRef{
			VaultID:       api.ExtensionParametersKeyVaultRef.VaultID,
			SecretName:    api.ExtensionParametersKeyVaultRef.SecretName,
			SecretVersion: api.ExtensionParametersKeyVaultRef.SecretVersion,
		}
	}
	obj.RootURL = api.RootURL
	obj.Script = api.Script
	obj.URLQuery = api.URLQuery
}

func convertExtensionToVLabs(api *Extension, vlabs *vlabs.Extension) {
	vlabs.Name = api.Name
	vlabs.SingleOrAll = api.SingleOrAll
	vlabs.Template = api.Template
}

func convertLinuxProfileToVLabs(obj *LinuxProfile, vlabsProfile *vlabs.LinuxProfile) {
	vlabsProfile.AdminUsername = obj.AdminUsername
	vlabsProfile.SSH.PublicKeys = []vlabs.PublicKey{}
	for _, d := range obj.SSH.PublicKeys {
		vlabsProfile.SSH.PublicKeys = append(vlabsProfile.SSH.PublicKeys,
			vlabs.PublicKey{KeyData: d.KeyData})
	}
	vlabsProfile.Secrets = []vlabs.KeyVaultSecrets{}
	for _, s := range obj.Secrets {
		secret := &vlabs.KeyVaultSecrets{}
		convertKeyVaultSecretsToVlabs(&s, secret)
		vlabsProfile.Secrets = append(vlabsProfile.Secrets, *secret)
	}
	vlabsProfile.ScriptRootURL = obj.ScriptRootURL
	if obj.CustomSearchDomain != nil {
		vlabsProfile.CustomSearchDomain = &vlabs.CustomSearchDomain{}
		vlabsProfile.CustomSearchDomain.Name = obj.CustomSearchDomain.Name
		vlabsProfile.CustomSearchDomain.RealmUser = obj.CustomSearchDomain.RealmUser
		vlabsProfile.CustomSearchDomain.RealmPassword = obj.CustomSearchDomain.RealmPassword
	}

	if obj.CustomNodesDNS != nil {
		vlabsProfile.CustomNodesDNS = &vlabs.CustomNodesDNS{}
		vlabsProfile.CustomNodesDNS.DNSServer = obj.CustomNodesDNS.DNSServer
	}
}

func convertWindowsProfileToVLabs(api *WindowsProfile, vlabsProfile *vlabs.WindowsProfile) {
	vlabsProfile.AdminUsername = api.AdminUsername
	vlabsProfile.AdminPassword = api.AdminPassword
	vlabsProfile.CSIProxyURL = api.CSIProxyURL
	vlabsProfile.EnableCSIProxy = api.EnableCSIProxy
	if api.ImageRef != nil {
		vlabsProfile.ImageRef = &vlabs.ImageReference{}
		vlabsProfile.ImageRef.Gallery = api.ImageRef.Gallery
		vlabsProfile.ImageRef.Name = api.ImageRef.Name
		vlabsProfile.ImageRef.ResourceGroup = api.ImageRef.ResourceGroup
		vlabsProfile.ImageRef.SubscriptionID = api.ImageRef.SubscriptionID
		vlabsProfile.ImageRef.Version = api.ImageRef.Version
	}
	vlabsProfile.ImageVersion = api.ImageVersion
	vlabsProfile.WindowsImageSourceURL = api.WindowsImageSourceURL
	vlabsProfile.WindowsPublisher = api.WindowsPublisher
	vlabsProfile.WindowsOffer = api.WindowsOffer
	vlabsProfile.WindowsSku = api.WindowsSku
	vlabsProfile.WindowsDockerVersion = api.WindowsDockerVersion
	vlabsProfile.Secrets = []vlabs.KeyVaultSecrets{}
	for _, s := range api.Secrets {
		secret := &vlabs.KeyVaultSecrets{}
		convertKeyVaultSecretsToVlabs(&s, secret)
		vlabsProfile.Secrets = append(vlabsProfile.Secrets, *secret)
	}
	if api.SSHEnabled != nil {
		vlabsProfile.SSHEnabled = api.SSHEnabled
	}
	vlabsProfile.EnableAutomaticUpdates = api.EnableAutomaticUpdates
}

func convertOrchestratorProfileToVLabs(api *OrchestratorProfile, o *vlabs.OrchestratorProfile) {
	o.OrchestratorType = api.OrchestratorType

	if api.OrchestratorVersion != "" {
		o.OrchestratorVersion = api.OrchestratorVersion
		sv, _ := semver.Make(o.OrchestratorVersion)
		o.OrchestratorRelease = fmt.Sprintf("%d.%d", sv.Major, sv.Minor)
	}

	if api.KubernetesConfig != nil {
		o.KubernetesConfig = &vlabs.KubernetesConfig{}
		convertKubernetesConfigToVLabs(api.KubernetesConfig, o.KubernetesConfig)
	}

	if api.DcosConfig != nil {
		o.DcosConfig = &vlabs.DcosConfig{}
		convertDcosConfigToVLabs(api.DcosConfig, o.DcosConfig)
	}
}

func convertDcosConfigToVLabs(api *DcosConfig, vl *vlabs.DcosConfig) {
	vl.DcosBootstrapURL = api.DcosBootstrapURL
	vl.DcosWindowsBootstrapURL = api.DcosWindowsBootstrapURL

	if api.Registry != "" {
		vl.Registry = api.Registry
	}

	if api.RegistryUser != "" {
		vl.RegistryUser = api.RegistryUser
	}

	if api.RegistryPass != "" {
		vl.RegistryPass = api.RegistryPass
	}
	vl.DcosRepositoryURL = api.DcosRepositoryURL
	vl.DcosClusterPackageListID = api.DcosClusterPackageListID
	vl.DcosProviderPackageID = api.DcosProviderPackageID

	if api.BootstrapProfile != nil {
		vl.BootstrapProfile = &vlabs.BootstrapProfile{
			VMSize:       api.BootstrapProfile.VMSize,
			OSDiskSizeGB: api.BootstrapProfile.OSDiskSizeGB,
			OAuthEnabled: api.BootstrapProfile.OAuthEnabled,
			StaticIP:     api.BootstrapProfile.StaticIP,
			Subnet:       api.BootstrapProfile.Subnet,
		}
	}
}

func convertKubernetesConfigToVLabs(apiCfg *KubernetesConfig, vlabsCfg *vlabs.KubernetesConfig) {
	vlabsCfg.KubernetesImageBase = apiCfg.KubernetesImageBase
	vlabsCfg.KubernetesImageBaseType = apiCfg.KubernetesImageBaseType
	vlabsCfg.MCRKubernetesImageBase = apiCfg.MCRKubernetesImageBase
	vlabsCfg.ClusterSubnet = apiCfg.ClusterSubnet
	vlabsCfg.DNSServiceIP = apiCfg.DNSServiceIP
	vlabsCfg.ServiceCidr = apiCfg.ServiceCIDR
	vlabsCfg.NetworkPolicy = apiCfg.NetworkPolicy
	vlabsCfg.NetworkPlugin = apiCfg.NetworkPlugin
	vlabsCfg.NetworkMode = apiCfg.NetworkMode
	vlabsCfg.ContainerRuntime = apiCfg.ContainerRuntime
	vlabsCfg.MaxPods = apiCfg.MaxPods
	vlabsCfg.DockerBridgeSubnet = apiCfg.DockerBridgeSubnet
	vlabsCfg.MobyVersion = apiCfg.MobyVersion
	vlabsCfg.ContainerdVersion = apiCfg.ContainerdVersion
	vlabsCfg.CloudProviderBackoff = apiCfg.CloudProviderBackoff
	vlabsCfg.CloudProviderBackoffMode = apiCfg.CloudProviderBackoffMode
	vlabsCfg.CloudProviderBackoffDuration = apiCfg.CloudProviderBackoffDuration
	vlabsCfg.CloudProviderBackoffExponent = apiCfg.CloudProviderBackoffExponent
	vlabsCfg.CloudProviderBackoffJitter = apiCfg.CloudProviderBackoffJitter
	vlabsCfg.CloudProviderBackoffRetries = apiCfg.CloudProviderBackoffRetries
	vlabsCfg.CloudProviderRateLimit = apiCfg.CloudProviderRateLimit
	vlabsCfg.CloudProviderRateLimitBucket = apiCfg.CloudProviderRateLimitBucket
	vlabsCfg.CloudProviderRateLimitBucketWrite = apiCfg.CloudProviderRateLimitBucketWrite
	vlabsCfg.CloudProviderRateLimitQPS = apiCfg.CloudProviderRateLimitQPS
	vlabsCfg.CloudProviderRateLimitQPSWrite = apiCfg.CloudProviderRateLimitQPSWrite
	vlabsCfg.UseManagedIdentity = apiCfg.UseManagedIdentity
	vlabsCfg.UserAssignedID = apiCfg.UserAssignedID
	vlabsCfg.UserAssignedClientID = apiCfg.UserAssignedClientID
	vlabsCfg.CustomHyperkubeImage = apiCfg.CustomHyperkubeImage
	vlabsCfg.CustomKubeAPIServerImage = apiCfg.CustomKubeAPIServerImage
	vlabsCfg.CustomKubeControllerManagerImage = apiCfg.CustomKubeControllerManagerImage
	vlabsCfg.CustomKubeProxyImage = apiCfg.CustomKubeProxyImage
	vlabsCfg.CustomKubeSchedulerImage = apiCfg.CustomKubeSchedulerImage
	vlabsCfg.CustomKubeBinaryURL = apiCfg.CustomKubeBinaryURL
	vlabsCfg.CustomCcmImage = apiCfg.CustomCcmImage
	vlabsCfg.UseCloudControllerManager = apiCfg.UseCloudControllerManager
	vlabsCfg.CustomWindowsPackageURL = apiCfg.CustomWindowsPackageURL
	vlabsCfg.WindowsNodeBinariesURL = apiCfg.WindowsNodeBinariesURL
	vlabsCfg.WindowsContainerdURL = apiCfg.WindowsContainerdURL
	vlabsCfg.WindowsSdnPluginURL = apiCfg.WindowsSdnPluginURL
	vlabsCfg.UseInstanceMetadata = apiCfg.UseInstanceMetadata
	vlabsCfg.LoadBalancerSku = apiCfg.LoadBalancerSku
	vlabsCfg.ExcludeMasterFromStandardLB = apiCfg.ExcludeMasterFromStandardLB
	vlabsCfg.LoadBalancerOutboundIPs = apiCfg.LoadBalancerOutboundIPs
	vlabsCfg.EnableRbac = apiCfg.EnableRbac
	vlabsCfg.EnableSecureKubelet = apiCfg.EnableSecureKubelet
	vlabsCfg.EnableAggregatedAPIs = apiCfg.EnableAggregatedAPIs
	vlabsCfg.EnableDataEncryptionAtRest = apiCfg.EnableDataEncryptionAtRest
	vlabsCfg.EnableEncryptionWithExternalKms = apiCfg.EnableEncryptionWithExternalKms
	vlabsCfg.EnablePodSecurityPolicy = apiCfg.EnablePodSecurityPolicy
	vlabsCfg.GCHighThreshold = apiCfg.GCHighThreshold
	vlabsCfg.GCLowThreshold = apiCfg.GCLowThreshold
	vlabsCfg.EtcdVersion = apiCfg.EtcdVersion
	vlabsCfg.EtcdDiskSizeGB = apiCfg.EtcdDiskSizeGB
	vlabsCfg.EtcdEncryptionKey = apiCfg.EtcdEncryptionKey
	vlabsCfg.AzureCNIVersion = apiCfg.AzureCNIVersion
	vlabsCfg.AzureCNIURLLinux = apiCfg.AzureCNIURLLinux
	vlabsCfg.AzureCNIURLWindows = apiCfg.AzureCNIURLWindows
	vlabsCfg.KeyVaultSku = apiCfg.KeyVaultSku
	vlabsCfg.MaximumLoadBalancerRuleCount = apiCfg.MaximumLoadBalancerRuleCount
	vlabsCfg.ProxyMode = vlabs.KubeProxyMode(apiCfg.ProxyMode)
	vlabsCfg.PrivateAzureRegistryServer = apiCfg.PrivateAzureRegistryServer
	vlabsCfg.OutboundRuleIdleTimeoutInMinutes = apiCfg.OutboundRuleIdleTimeoutInMinutes
	vlabsCfg.CloudProviderDisableOutboundSNAT = apiCfg.CloudProviderDisableOutboundSNAT
	vlabsCfg.KubeReservedCgroup = apiCfg.KubeReservedCgroup
	convertComponentsToVlabs(apiCfg, vlabsCfg)
	convertAddonsToVlabs(apiCfg, vlabsCfg)
	convertKubeletConfigToVlabs(apiCfg, vlabsCfg)
	convertControllerManagerConfigToVlabs(apiCfg, vlabsCfg)
	convertCloudControllerManagerConfigToVlabs(apiCfg, vlabsCfg)
	convertAPIServerConfigToVlabs(apiCfg, vlabsCfg)
	convertSchedulerConfigToVlabs(apiCfg, vlabsCfg)
	convertPrivateClusterToVlabs(apiCfg, vlabsCfg)
	convertPodSecurityPolicyConfigToVlabs(apiCfg, vlabsCfg)
	convertContainerRuntimeConfigToVlabs(apiCfg, vlabsCfg)
}

func convertContainerRuntimeConfigToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.ContainerRuntimeConfig = map[string]string{}
	for key, val := range a.ContainerRuntimeConfig {
		v.ContainerRuntimeConfig[key] = val
	}
}

func convertKubeletConfigToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.KubeletConfig = map[string]string{}
	for key, val := range a.KubeletConfig {
		v.KubeletConfig[key] = val
	}
}

func convertCustomFilesToVlabs(a *MasterProfile, v *vlabs.MasterProfile) {
	if a.CustomFiles != nil {
		v.CustomFiles = &[]vlabs.CustomFile{}
		for i := range *a.CustomFiles {
			*v.CustomFiles = append(*v.CustomFiles, vlabs.CustomFile{
				Dest:   (*a.CustomFiles)[i].Dest,
				Source: (*a.CustomFiles)[i].Source,
			})
		}
	}
}

func convertControllerManagerConfigToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.ControllerManagerConfig = map[string]string{}
	for key, val := range a.ControllerManagerConfig {
		v.ControllerManagerConfig[key] = val
	}
}

func convertCloudControllerManagerConfigToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.CloudControllerManagerConfig = map[string]string{}
	for key, val := range a.CloudControllerManagerConfig {
		v.CloudControllerManagerConfig[key] = val
	}
}

func convertAPIServerConfigToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.APIServerConfig = map[string]string{}
	for key, val := range a.APIServerConfig {
		v.APIServerConfig[key] = val
	}
}

func convertSchedulerConfigToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.SchedulerConfig = map[string]string{}
	for key, val := range a.SchedulerConfig {
		v.SchedulerConfig[key] = val
	}
}

func convertPodSecurityPolicyConfigToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.PodSecurityPolicyConfig = map[string]string{}
	for key, val := range a.PodSecurityPolicyConfig {
		v.PodSecurityPolicyConfig[key] = val
	}
}

func convertPrivateClusterToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	if a.PrivateCluster != nil {
		v.PrivateCluster = &vlabs.PrivateCluster{}
		v.PrivateCluster.Enabled = a.PrivateCluster.Enabled
		if a.PrivateCluster.JumpboxProfile != nil {
			v.PrivateCluster.JumpboxProfile = &vlabs.PrivateJumpboxProfile{}
			convertPrivateJumpboxProfileToVlabs(a.PrivateCluster.JumpboxProfile, v.PrivateCluster.JumpboxProfile)
		}
	}
}

func convertPrivateJumpboxProfileToVlabs(api *PrivateJumpboxProfile, vlabsProfile *vlabs.PrivateJumpboxProfile) {
	vlabsProfile.Name = api.Name
	vlabsProfile.OSDiskSizeGB = api.OSDiskSizeGB
	vlabsProfile.VMSize = api.VMSize
	vlabsProfile.PublicKey = api.PublicKey
	vlabsProfile.Username = api.Username
	vlabsProfile.StorageProfile = api.StorageProfile
}

func convertComponentsToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.Components = []vlabs.KubernetesComponent{}
	for i := range a.Components {
		v.Components = append(v.Components, vlabs.KubernetesComponent{
			Name:    a.Components[i].Name,
			Enabled: a.Components[i].Enabled,
			Config:  map[string]string{},
			Data:    a.Components[i].Data,
		})
		for j := range a.Components[i].Containers {
			v.Components[i].Containers = append(v.Components[i].Containers, vlabs.KubernetesContainerSpec{
				Name:           a.Components[i].Containers[j].Name,
				Image:          a.Components[i].Containers[j].Image,
				CPURequests:    a.Components[i].Containers[j].CPURequests,
				MemoryRequests: a.Components[i].Containers[j].MemoryRequests,
				CPULimits:      a.Components[i].Containers[j].CPULimits,
				MemoryLimits:   a.Components[i].Containers[j].MemoryLimits,
			})
		}
		if a.Components[i].Config != nil {
			for key, val := range a.Components[i].Config {
				v.Components[i].Config[key] = val
			}
		}
	}
}

func convertAddonsToVlabs(a *KubernetesConfig, v *vlabs.KubernetesConfig) {
	v.Addons = []vlabs.KubernetesAddon{}
	for i := range a.Addons {
		v.Addons = append(v.Addons, vlabs.KubernetesAddon{
			Name:    a.Addons[i].Name,
			Enabled: a.Addons[i].Enabled,
			Mode:    a.Addons[i].Mode,
			Config:  map[string]string{},
			Data:    a.Addons[i].Data,
		})
		for j := range a.Addons[i].Containers {
			v.Addons[i].Containers = append(v.Addons[i].Containers, vlabs.KubernetesContainerSpec{
				Name:           a.Addons[i].Containers[j].Name,
				Image:          a.Addons[i].Containers[j].Image,
				CPURequests:    a.Addons[i].Containers[j].CPURequests,
				MemoryRequests: a.Addons[i].Containers[j].MemoryRequests,
				CPULimits:      a.Addons[i].Containers[j].CPULimits,
				MemoryLimits:   a.Addons[i].Containers[j].MemoryLimits,
			})
		}
		for k := range a.Addons[i].Pools {
			v.Addons[i].Pools = append(v.Addons[i].Pools, vlabs.AddonNodePoolsConfig{
				Name:   a.Addons[i].Pools[k].Name,
				Config: map[string]string{},
			})
			if a.Addons[i].Pools[k].Config != nil {
				for key, val := range a.Addons[i].Pools[k].Config {
					v.Addons[i].Pools[k].Config[key] = val
				}
			}
		}
		if a.Addons[i].Config != nil {
			for key, val := range a.Addons[i].Config {
				v.Addons[i].Config[key] = val
			}
		}
	}
}

func convertMasterProfileToVLabs(api *MasterProfile, vlabsProfile *vlabs.MasterProfile) {
	vlabsProfile.Count = api.Count
	vlabsProfile.DNSPrefix = api.DNSPrefix
	vlabsProfile.CustomVMTags = api.CustomVMTags
	vlabsProfile.SubjectAltNames = api.SubjectAltNames
	vlabsProfile.VMSize = api.VMSize
	vlabsProfile.OSDiskSizeGB = api.OSDiskSizeGB
	vlabsProfile.VnetSubnetID = api.VnetSubnetID
	vlabsProfile.AgentVnetSubnetID = api.AgentVnetSubnetID
	vlabsProfile.FirstConsecutiveStaticIP = api.FirstConsecutiveStaticIP
	vlabsProfile.VnetCidr = api.VnetCidr
	vlabsProfile.SetSubnet(api.Subnet)
	vlabsProfile.SetSubnetIPv6(api.SubnetIPv6)
	vlabsProfile.FQDN = api.FQDN
	vlabsProfile.StorageProfile = api.StorageProfile
	if api.PreprovisionExtension != nil {
		vlabsExtension := &vlabs.Extension{}
		convertExtensionToVLabs(api.PreprovisionExtension, vlabsExtension)
		vlabsProfile.PreProvisionExtension = vlabsExtension
	}
	vlabsProfile.Extensions = []vlabs.Extension{}
	for _, extension := range api.Extensions {
		vlabsExtension := &vlabs.Extension{}
		convertExtensionToVLabs(&extension, vlabsExtension)
		vlabsProfile.Extensions = append(vlabsProfile.Extensions, *vlabsExtension)
	}
	vlabsProfile.Distro = vlabs.Distro(api.Distro)
	if api.KubernetesConfig != nil {
		vlabsProfile.KubernetesConfig = &vlabs.KubernetesConfig{}
		convertKubernetesConfigToVLabs(api.KubernetesConfig, vlabsProfile.KubernetesConfig)
	}
	if api.ImageRef != nil {
		vlabsProfile.ImageRef = &vlabs.ImageReference{}
		vlabsProfile.ImageRef.Name = api.ImageRef.Name
		vlabsProfile.ImageRef.ResourceGroup = api.ImageRef.ResourceGroup
		vlabsProfile.ImageRef.SubscriptionID = api.ImageRef.SubscriptionID
		vlabsProfile.ImageRef.Gallery = api.ImageRef.Gallery
		vlabsProfile.ImageRef.Version = api.ImageRef.Version
	}
	vlabsProfile.AvailabilityProfile = api.AvailabilityProfile
	vlabsProfile.AgentSubnet = api.AgentSubnet
	vlabsProfile.AvailabilityZones = api.AvailabilityZones
	vlabsProfile.PlatformFaultDomainCount = api.PlatformFaultDomainCount
	vlabsProfile.PlatformUpdateDomainCount = api.PlatformUpdateDomainCount
	vlabsProfile.SinglePlacementGroup = api.SinglePlacementGroup
	vlabsProfile.CosmosEtcd = api.CosmosEtcd
	vlabsProfile.AuditDEnabled = api.AuditDEnabled
	vlabsProfile.UltraSSDEnabled = api.UltraSSDEnabled
	vlabsProfile.EncryptionAtHost = api.EncryptionAtHost
	vlabsProfile.ProximityPlacementGroupID = api.ProximityPlacementGroupID
	vlabsProfile.OSDiskCachingType = api.OSDiskCachingType
	convertCustomFilesToVlabs(api, vlabsProfile)
	vlabsProfile.SysctlDConfig = map[string]string{}
	for key, val := range api.SysctlDConfig {
		vlabsProfile.SysctlDConfig[key] = val
	}
}

func convertKeyVaultSecretsToVlabs(api *KeyVaultSecrets, vlabsSecrets *vlabs.KeyVaultSecrets) {
	vlabsSecrets.SourceVault = &vlabs.KeyVaultID{ID: api.SourceVault.ID}
	vlabsSecrets.VaultCertificates = []vlabs.KeyVaultCertificate{}
	for _, c := range api.VaultCertificates {
		cert := vlabs.KeyVaultCertificate{}
		cert.CertificateStore = c.CertificateStore
		cert.CertificateURL = c.CertificateURL
		vlabsSecrets.VaultCertificates = append(vlabsSecrets.VaultCertificates, cert)
	}
}

func convertAgentPoolProfileToVLabs(api *AgentPoolProfile, p *vlabs.AgentPoolProfile) {
	p.Name = api.Name
	p.Count = api.Count
	p.VMSize = api.VMSize
	p.CustomVMTags = api.CustomVMTags
	p.OSDiskSizeGB = api.OSDiskSizeGB
	p.DNSPrefix = api.DNSPrefix
	p.OSType = vlabs.OSType(api.OSType)
	p.Ports = []int{}
	p.Ports = append(p.Ports, api.Ports...)
	p.AvailabilityProfile = api.AvailabilityProfile
	p.ScaleSetPriority = api.ScaleSetPriority
	p.ScaleSetEvictionPolicy = api.ScaleSetEvictionPolicy
	p.SpotMaxPrice = api.SpotMaxPrice
	p.StorageProfile = api.StorageProfile
	p.DiskSizesGB = []int{}
	p.DiskSizesGB = append(p.DiskSizesGB, api.DiskSizesGB...)
	p.VnetSubnetID = api.VnetSubnetID
	p.SetSubnet(api.Subnet)
	p.FQDN = api.FQDN
	p.CustomNodeLabels = map[string]string{}
	p.AcceleratedNetworkingEnabled = api.AcceleratedNetworkingEnabled
	p.AcceleratedNetworkingEnabledWindows = api.AcceleratedNetworkingEnabledWindows
	p.VMSSOverProvisioningEnabled = api.VMSSOverProvisioningEnabled
	p.AvailabilityZones = api.AvailabilityZones
	p.SinglePlacementGroup = api.SinglePlacementGroup
	p.PlatformFaultDomainCount = api.PlatformFaultDomainCount
	p.PlatformUpdateDomainCount = api.PlatformUpdateDomainCount
	p.EnableVMSSNodePublicIP = api.EnableVMSSNodePublicIP
	p.LoadBalancerBackendAddressPoolIDs = api.LoadBalancerBackendAddressPoolIDs
	p.AuditDEnabled = api.AuditDEnabled
	p.UltraSSDEnabled = api.UltraSSDEnabled
	p.DiskEncryptionSetID = api.DiskEncryptionSetID
	p.EncryptionAtHost = api.EncryptionAtHost
	p.ProximityPlacementGroupID = api.ProximityPlacementGroupID

	for k, v := range api.CustomNodeLabels {
		p.CustomNodeLabels[k] = v
	}

	if api.PreprovisionExtension != nil {
		vlabsExtension := &vlabs.Extension{}
		convertExtensionToVLabs(api.PreprovisionExtension, vlabsExtension)
		p.PreProvisionExtension = vlabsExtension
	}

	p.Extensions = []vlabs.Extension{}
	for _, extension := range api.Extensions {
		vlabsExtension := &vlabs.Extension{}
		convertExtensionToVLabs(&extension, vlabsExtension)
		p.Extensions = append(p.Extensions, *vlabsExtension)
	}
	p.Distro = vlabs.Distro(api.Distro)
	if api.KubernetesConfig != nil {
		p.KubernetesConfig = &vlabs.KubernetesConfig{}
		convertKubernetesConfigToVLabs(api.KubernetesConfig, p.KubernetesConfig)
	}
	if api.ImageRef != nil {
		p.ImageRef = &vlabs.ImageReference{}
		p.ImageRef.Name = api.ImageRef.Name
		p.ImageRef.ResourceGroup = api.ImageRef.ResourceGroup
		p.ImageRef.SubscriptionID = api.ImageRef.SubscriptionID
		p.ImageRef.Gallery = api.ImageRef.Gallery
		p.ImageRef.Version = api.ImageRef.Version
	}
	p.Role = vlabs.AgentPoolProfileRole(api.Role)
	p.SysctlDConfig = map[string]string{}
	for key, val := range api.SysctlDConfig {
		p.SysctlDConfig[key] = val
	}
	p.OSDiskCachingType = api.OSDiskCachingType
	p.DataDiskCachingType = api.DataDiskCachingType
}

func convertServicePrincipalProfileToVLabs(api *ServicePrincipalProfile, v *vlabs.ServicePrincipalProfile) {
	v.ClientID = api.ClientID
	v.Secret = api.Secret
	v.ObjectID = api.ObjectID
	if api.KeyvaultSecretRef != nil {
		v.KeyvaultSecretRef = &vlabs.KeyvaultSecretRef{
			VaultID:       api.KeyvaultSecretRef.VaultID,
			SecretName:    api.KeyvaultSecretRef.SecretName,
			SecretVersion: api.KeyvaultSecretRef.SecretVersion,
		}
	}
}

func convertCertificateProfileToVLabs(api *CertificateProfile, vlabs *vlabs.CertificateProfile) {
	vlabs.CaCertificate = api.CaCertificate
	vlabs.CaPrivateKey = api.CaPrivateKey
	vlabs.APIServerCertificate = api.APIServerCertificate
	vlabs.APIServerPrivateKey = api.APIServerPrivateKey
	vlabs.ClientCertificate = api.ClientCertificate
	vlabs.ClientPrivateKey = api.ClientPrivateKey
	vlabs.KubeConfigCertificate = api.KubeConfigCertificate
	vlabs.KubeConfigPrivateKey = api.KubeConfigPrivateKey
	vlabs.EtcdServerCertificate = api.EtcdServerCertificate
	vlabs.EtcdServerPrivateKey = api.EtcdServerPrivateKey
	vlabs.EtcdClientCertificate = api.EtcdClientCertificate
	vlabs.EtcdClientPrivateKey = api.EtcdClientPrivateKey
	vlabs.EtcdPeerCertificates = api.EtcdPeerCertificates
	vlabs.EtcdPeerPrivateKeys = api.EtcdPeerPrivateKeys
}

func convertAADProfileToVLabs(api *AADProfile, vlabs *vlabs.AADProfile) {
	vlabs.ClientAppID = api.ClientAppID
	vlabs.ServerAppID = api.ServerAppID
	vlabs.TenantID = api.TenantID
	vlabs.AdminGroupID = api.AdminGroupID
}

func convertFeatureFlagsToVLabs(api *FeatureFlags, vlabs *vlabs.FeatureFlags) {
	vlabs.EnableCSERunInBackground = api.EnableCSERunInBackground
	vlabs.BlockOutboundInternet = api.BlockOutboundInternet
	vlabs.EnableIPv6DualStack = api.EnableIPv6DualStack
	vlabs.EnableTelemetry = api.EnableTelemetry
	vlabs.EnableIPv6Only = api.EnableIPv6Only
}

func convertCloudProfileToVLabs(api *CustomCloudProfile, vlabsccp *vlabs.CustomCloudProfile) {
	if api.Environment != nil {
		vlabsccp.Environment = &azure.Environment{}
		vlabsccp.Environment.Name = api.Environment.Name
		vlabsccp.Environment.ManagementPortalURL = api.Environment.ManagementPortalURL
		vlabsccp.Environment.PublishSettingsURL = api.Environment.PublishSettingsURL
		vlabsccp.Environment.ServiceManagementEndpoint = api.Environment.ServiceManagementEndpoint
		vlabsccp.Environment.ResourceManagerEndpoint = api.Environment.ResourceManagerEndpoint
		vlabsccp.Environment.ActiveDirectoryEndpoint = api.Environment.ActiveDirectoryEndpoint
		vlabsccp.Environment.GalleryEndpoint = api.Environment.GalleryEndpoint
		vlabsccp.Environment.KeyVaultEndpoint = api.Environment.KeyVaultEndpoint
		vlabsccp.Environment.GraphEndpoint = api.Environment.GraphEndpoint
		vlabsccp.Environment.ServiceBusEndpoint = api.Environment.ServiceBusEndpoint
		vlabsccp.Environment.BatchManagementEndpoint = api.Environment.BatchManagementEndpoint
		vlabsccp.Environment.StorageEndpointSuffix = api.Environment.StorageEndpointSuffix
		vlabsccp.Environment.SQLDatabaseDNSSuffix = api.Environment.SQLDatabaseDNSSuffix
		vlabsccp.Environment.TrafficManagerDNSSuffix = api.Environment.TrafficManagerDNSSuffix
		vlabsccp.Environment.KeyVaultDNSSuffix = api.Environment.KeyVaultDNSSuffix
		vlabsccp.Environment.ServiceBusEndpointSuffix = api.Environment.ServiceBusEndpointSuffix
		vlabsccp.Environment.ServiceManagementVMDNSSuffix = api.Environment.ServiceManagementVMDNSSuffix
		vlabsccp.Environment.ResourceManagerVMDNSSuffix = api.Environment.ResourceManagerVMDNSSuffix
		vlabsccp.Environment.ContainerRegistryDNSSuffix = api.Environment.ContainerRegistryDNSSuffix
		vlabsccp.Environment.TokenAudience = api.Environment.TokenAudience
	}

	if api.AzureEnvironmentSpecConfig != nil {
		vlabsccp.AzureEnvironmentSpecConfig = &vlabs.AzureEnvironmentSpecConfig{}
		convertAzureEnvironmentSpecConfigToVLabs(api.AzureEnvironmentSpecConfig, vlabsccp.AzureEnvironmentSpecConfig)
	}
	vlabsccp.IdentitySystem = api.IdentitySystem
	vlabsccp.AuthenticationMethod = api.AuthenticationMethod
	vlabsccp.DependenciesLocation = vlabs.DependenciesLocation(api.DependenciesLocation)
	vlabsccp.PortalURL = api.PortalURL
	vlabsccp.CustomCloudRootCertificates = api.CustomCloudRootCertificates
	vlabsccp.CustomCloudSourcesList = api.CustomCloudSourcesList
}

func convertTelemetryProfileToVLabs(api *TelemetryProfile, vlabstp *vlabs.TelemetryProfile) {
	vlabstp.ApplicationInsightsKey = api.ApplicationInsightsKey
}

func convertAzureEnvironmentSpecConfigToVLabs(api *AzureEnvironmentSpecConfig, vlabses *vlabs.AzureEnvironmentSpecConfig) {
	vlabses.CloudName = api.CloudName
	vlabses.DCOSSpecConfig = vlabs.DCOSSpecConfig{
		DCOS188BootstrapDownloadURL:     api.DCOSSpecConfig.DCOS188BootstrapDownloadURL,
		DCOS190BootstrapDownloadURL:     api.DCOSSpecConfig.DCOS190BootstrapDownloadURL,
		DCOS198BootstrapDownloadURL:     api.DCOSSpecConfig.DCOS198BootstrapDownloadURL,
		DCOS110BootstrapDownloadURL:     api.DCOSSpecConfig.DCOS110BootstrapDownloadURL,
		DCOS111BootstrapDownloadURL:     api.DCOSSpecConfig.DCOS111BootstrapDownloadURL,
		DCOSWindowsBootstrapDownloadURL: api.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL,
		DcosRepositoryURL:               api.DCOSSpecConfig.DcosRepositoryURL,
		DcosClusterPackageListID:        api.DCOSSpecConfig.DcosClusterPackageListID,
		DcosProviderPackageID:           api.DCOSSpecConfig.DcosProviderPackageID,
	}

	vlabses.DockerSpecConfig = vlabs.DockerSpecConfig{
		DockerEngineRepo:         api.DockerSpecConfig.DockerEngineRepo,
		DockerComposeDownloadURL: api.DockerSpecConfig.DockerComposeDownloadURL,
	}
	vlabses.EndpointConfig = vlabs.AzureEndpointConfig{
		ResourceManagerVMDNSSuffix: api.EndpointConfig.ResourceManagerVMDNSSuffix,
	}
	vlabses.KubernetesSpecConfig = vlabs.KubernetesSpecConfig{
		AzureTelemetryPID:                api.KubernetesSpecConfig.AzureTelemetryPID,
		KubernetesImageBase:              api.KubernetesSpecConfig.KubernetesImageBase,
		MCRKubernetesImageBase:           api.KubernetesSpecConfig.MCRKubernetesImageBase,
		TillerImageBase:                  api.KubernetesSpecConfig.TillerImageBase,
		ACIConnectorImageBase:            api.KubernetesSpecConfig.ACIConnectorImageBase,
		NVIDIAImageBase:                  api.KubernetesSpecConfig.NVIDIAImageBase,
		AzureCNIImageBase:                api.KubernetesSpecConfig.AzureCNIImageBase,
		CalicoImageBase:                  api.KubernetesSpecConfig.CalicoImageBase,
		EtcdDownloadURLBase:              api.KubernetesSpecConfig.EtcdDownloadURLBase,
		KubeBinariesSASURLBase:           api.KubernetesSpecConfig.KubeBinariesSASURLBase,
		WindowsTelemetryGUID:             api.KubernetesSpecConfig.WindowsTelemetryGUID,
		CNIPluginsDownloadURL:            api.KubernetesSpecConfig.CNIPluginsDownloadURL,
		VnetCNILinuxPluginsDownloadURL:   api.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL,
		VnetCNIWindowsPluginsDownloadURL: api.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL,
		ContainerdDownloadURLBase:        api.KubernetesSpecConfig.ContainerdDownloadURLBase,
		CSIProxyDownloadURL:              api.KubernetesSpecConfig.CSIProxyDownloadURL,
	}
	vlabses.OSImageConfig = map[vlabs.Distro]vlabs.AzureOSImageConfig{}
	for k, v := range api.OSImageConfig {
		vlabses.OSImageConfig[vlabs.Distro(string(k))] = vlabs.AzureOSImageConfig{
			ImageOffer:     v.ImageOffer,
			ImageSku:       v.ImageSku,
			ImagePublisher: v.ImagePublisher,
			ImageVersion:   v.ImageVersion,
		}
	}
}
