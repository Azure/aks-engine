// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
)

///////////////////////////////////////////////////////////
// The converter exposes functions to convert the top level
// ContainerService resource
//
// All other functions are internal helper functions used
// for converting.
///////////////////////////////////////////////////////////

// ConvertVLabsContainerService converts a vlabs ContainerService to an unversioned ContainerService
func ConvertVLabsContainerService(vlabs *vlabs.ContainerService, isUpdate bool) (*ContainerService, error) {
	c := &ContainerService{}
	c.ID = vlabs.ID
	c.Location = helpers.NormalizeAzureRegion(vlabs.Location)
	c.Name = vlabs.Name
	if vlabs.Plan != nil {
		c.Plan = &ResourcePurchasePlan{}
		convertVLabsResourcePurchasePlan(vlabs.Plan, c.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range vlabs.Tags {
		c.Tags[k] = v
	}
	c.Type = vlabs.Type
	c.Properties = &Properties{}
	if err := convertVLabsProperties(vlabs.Properties, c.Properties, isUpdate); err != nil {
		return nil, err
	}
	return c, nil
}

// convertVLabsResourcePurchasePlan converts a vlabs ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertVLabsResourcePurchasePlan(vlabs *vlabs.ResourcePurchasePlan, api *ResourcePurchasePlan) {
	api.Name = vlabs.Name
	api.Product = vlabs.Product
	api.PromotionCode = vlabs.PromotionCode
	api.Publisher = vlabs.Publisher
}

func convertVLabsProperties(vlabs *vlabs.Properties, api *Properties, isUpdate bool) error {
	api.ProvisioningState = ProvisioningState(vlabs.ProvisioningState)
	if vlabs.OrchestratorProfile != nil {
		api.OrchestratorProfile = &OrchestratorProfile{}
		if err := convertVLabsOrchestratorProfile(vlabs, api.OrchestratorProfile, isUpdate); err != nil {
			return err
		}
	}
	if vlabs.MasterProfile != nil {
		api.MasterProfile = &MasterProfile{}
		convertVLabsMasterProfile(vlabs.MasterProfile, api.MasterProfile)
	}
	api.AgentPoolProfiles = []*AgentPoolProfile{}
	for _, p := range vlabs.AgentPoolProfiles {
		apiProfile := &AgentPoolProfile{}
		convertVLabsAgentPoolProfile(p, apiProfile)
		// by default vlabs will use managed disks for all orchestrators but kubernetes as it has encryption at rest.
		if !api.OrchestratorProfile.IsKubernetes() {
			if len(p.StorageProfile) == 0 {
				apiProfile.StorageProfile = ManagedDisks
			}
		}
		api.AgentPoolProfiles = append(api.AgentPoolProfiles, apiProfile)
	}
	if vlabs.LinuxProfile != nil {
		api.LinuxProfile = &LinuxProfile{}
		convertVLabsLinuxProfile(vlabs.LinuxProfile, api.LinuxProfile)
	}
	api.ExtensionProfiles = []*ExtensionProfile{}
	for _, p := range vlabs.ExtensionProfiles {
		apiExtensionProfile := &ExtensionProfile{}
		convertVLabsExtensionProfile(p, apiExtensionProfile)
		api.ExtensionProfiles = append(api.ExtensionProfiles, apiExtensionProfile)
	}
	if vlabs.WindowsProfile != nil {
		api.WindowsProfile = &WindowsProfile{}
		convertVLabsWindowsProfile(vlabs.WindowsProfile, api.WindowsProfile)
	}
	if vlabs.ServicePrincipalProfile != nil {
		api.ServicePrincipalProfile = &ServicePrincipalProfile{}
		convertVLabsServicePrincipalProfile(vlabs.ServicePrincipalProfile, api.ServicePrincipalProfile)
	}
	if vlabs.CertificateProfile != nil {
		api.CertificateProfile = &CertificateProfile{}
		convertVLabsCertificateProfile(vlabs.CertificateProfile, api.CertificateProfile)
	}

	if vlabs.AADProfile != nil {
		api.AADProfile = &AADProfile{}
		convertVLabsAADProfile(vlabs.AADProfile, api.AADProfile)
	}

	if vlabs.FeatureFlags != nil {
		api.FeatureFlags = &FeatureFlags{}
		convertVLabsFeatureFlags(vlabs.FeatureFlags, api.FeatureFlags)
	}

	if vlabs.CustomCloudProfile != nil {
		api.CustomCloudProfile = &CustomCloudProfile{}
		convertVLabsCustomCloudProfile(vlabs.CustomCloudProfile, api.CustomCloudProfile)
	}

	if vlabs.TelemetryProfile != nil {
		api.TelemetryProfile = &TelemetryProfile{}
		convertVLabsTelemetryProfile(vlabs.TelemetryProfile, api.TelemetryProfile)
	}

	return nil
}

func convertVLabsFeatureFlags(vlabs *vlabs.FeatureFlags, api *FeatureFlags) {
	api.EnableCSERunInBackground = vlabs.EnableCSERunInBackground
	api.BlockOutboundInternet = vlabs.BlockOutboundInternet
	api.EnableIPv6DualStack = vlabs.EnableIPv6DualStack
	api.EnableTelemetry = vlabs.EnableTelemetry
	api.EnableIPv6Only = vlabs.EnableIPv6Only
}

func convertVLabsExtensionProfile(vlabs *vlabs.ExtensionProfile, api *ExtensionProfile) {
	api.Name = vlabs.Name
	api.Version = vlabs.Version
	api.ExtensionParameters = vlabs.ExtensionParameters
	if vlabs.ExtensionParametersKeyVaultRef != nil {
		api.ExtensionParametersKeyVaultRef = &KeyvaultSecretRef{
			VaultID:       vlabs.ExtensionParametersKeyVaultRef.VaultID,
			SecretName:    vlabs.ExtensionParametersKeyVaultRef.SecretName,
			SecretVersion: vlabs.ExtensionParametersKeyVaultRef.SecretVersion,
		}
	}
	api.RootURL = vlabs.RootURL
	api.Script = vlabs.Script
	api.URLQuery = vlabs.URLQuery
}

func convertVLabsExtension(vlabs *vlabs.Extension, api *Extension) {
	api.Name = vlabs.Name
	api.SingleOrAll = vlabs.SingleOrAll
	api.Template = vlabs.Template
}

func convertVLabsLinuxProfile(vlabs *vlabs.LinuxProfile, api *LinuxProfile) {
	api.AdminUsername = vlabs.AdminUsername
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range vlabs.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys,
			PublicKey{KeyData: d.KeyData})
	}
	api.Secrets = []KeyVaultSecrets{}
	for _, s := range vlabs.Secrets {
		secret := &KeyVaultSecrets{}
		convertVLabsKeyVaultSecrets(&s, secret)
		api.Secrets = append(api.Secrets, *secret)
	}
	api.ScriptRootURL = vlabs.ScriptRootURL
	if vlabs.CustomSearchDomain != nil {
		api.CustomSearchDomain = &CustomSearchDomain{}
		api.CustomSearchDomain.Name = vlabs.CustomSearchDomain.Name
		api.CustomSearchDomain.RealmUser = vlabs.CustomSearchDomain.RealmUser
		api.CustomSearchDomain.RealmPassword = vlabs.CustomSearchDomain.RealmPassword
	}

	if vlabs.CustomNodesDNS != nil {
		api.CustomNodesDNS = &CustomNodesDNS{}
		api.CustomNodesDNS.DNSServer = vlabs.CustomNodesDNS.DNSServer
	}
}

func convertVLabsWindowsProfile(vlabs *vlabs.WindowsProfile, api *WindowsProfile) {
	api.AdminUsername = vlabs.AdminUsername
	api.AdminPassword = vlabs.AdminPassword
	api.CSIProxyURL = vlabs.CSIProxyURL
	api.EnableCSIProxy = vlabs.EnableCSIProxy
	if vlabs.ImageRef != nil {
		api.ImageRef = &ImageReference{}
		api.ImageRef.Gallery = vlabs.ImageRef.Gallery
		api.ImageRef.Name = vlabs.ImageRef.Name
		api.ImageRef.ResourceGroup = vlabs.ImageRef.ResourceGroup
		api.ImageRef.SubscriptionID = vlabs.ImageRef.SubscriptionID
		api.ImageRef.Version = vlabs.ImageRef.Version
	}
	api.ImageVersion = vlabs.ImageVersion
	api.WindowsImageSourceURL = vlabs.WindowsImageSourceURL
	api.WindowsPublisher = vlabs.WindowsPublisher
	api.WindowsOffer = vlabs.WindowsOffer
	api.WindowsSku = vlabs.WindowsSku
	api.WindowsDockerVersion = vlabs.WindowsDockerVersion
	api.Secrets = []KeyVaultSecrets{}
	for _, s := range vlabs.Secrets {
		secret := &KeyVaultSecrets{}
		convertVLabsKeyVaultSecrets(&s, secret)
		api.Secrets = append(api.Secrets, *secret)
	}
	if vlabs.SSHEnabled != nil {
		api.SSHEnabled = vlabs.SSHEnabled
	}
	api.EnableAutomaticUpdates = vlabs.EnableAutomaticUpdates
}

func convertVLabsOrchestratorProfile(vp *vlabs.Properties, api *OrchestratorProfile, isUpdate bool) error {
	vlabscs := vp.OrchestratorProfile
	api.OrchestratorType = vlabscs.OrchestratorType
	switch api.OrchestratorType {
	case Kubernetes:
		if vlabscs.KubernetesConfig != nil {
			api.KubernetesConfig = &KubernetesConfig{}
			convertVLabsKubernetesConfig(vlabscs.KubernetesConfig, api.KubernetesConfig)
		}
		setVlabsKubernetesDefaults(vp, api)

		// TODO (hack): this validation should be done as part of the main validation, but deploy does it only after loading the container.
		if !isUpdate {
			if err := vp.ValidateOrchestratorProfile(isUpdate); err != nil {
				return err
			}
		}

		api.OrchestratorVersion = common.RationalizeReleaseAndVersion(
			vlabscs.OrchestratorType,
			vlabscs.OrchestratorRelease,
			vlabscs.OrchestratorVersion,
			isUpdate,
			vp.HasWindows())

	case DCOS:
		if vlabscs.DcosConfig != nil {
			api.DcosConfig = &DcosConfig{}
			convertVLabsDcosConfig(vlabscs.DcosConfig, api.DcosConfig)
		}
		api.OrchestratorVersion = common.RationalizeReleaseAndVersion(
			vlabscs.OrchestratorType,
			vlabscs.OrchestratorRelease,
			vlabscs.OrchestratorVersion,
			isUpdate,
			false)
	}

	return nil
}

func convertVLabsDcosConfig(vlabs *vlabs.DcosConfig, api *DcosConfig) {
	api.DcosBootstrapURL = vlabs.DcosBootstrapURL
	api.DcosWindowsBootstrapURL = vlabs.DcosWindowsBootstrapURL

	if len(vlabs.Registry) > 0 {
		api.Registry = vlabs.Registry
	}

	if len(vlabs.RegistryUser) > 0 {
		api.RegistryUser = vlabs.RegistryUser
	}

	if len(vlabs.RegistryPass) > 0 {
		api.RegistryPass = vlabs.RegistryPass
	}
	api.DcosRepositoryURL = vlabs.DcosRepositoryURL
	api.DcosClusterPackageListID = vlabs.DcosClusterPackageListID
	api.DcosProviderPackageID = vlabs.DcosProviderPackageID

	if vlabs.BootstrapProfile != nil {
		api.BootstrapProfile = &BootstrapProfile{
			VMSize:       vlabs.BootstrapProfile.VMSize,
			OSDiskSizeGB: vlabs.BootstrapProfile.OSDiskSizeGB,
			OAuthEnabled: vlabs.BootstrapProfile.OAuthEnabled,
			StaticIP:     vlabs.BootstrapProfile.StaticIP,
			Subnet:       vlabs.BootstrapProfile.Subnet,
		}
	}
}

func convertVLabsKubernetesConfig(vlabs *vlabs.KubernetesConfig, api *KubernetesConfig) {
	api.KubernetesImageBase = vlabs.KubernetesImageBase
	api.KubernetesImageBaseType = vlabs.KubernetesImageBaseType
	api.MCRKubernetesImageBase = vlabs.MCRKubernetesImageBase
	api.ClusterSubnet = vlabs.ClusterSubnet
	api.DNSServiceIP = vlabs.DNSServiceIP
	api.ServiceCIDR = vlabs.ServiceCidr
	api.NetworkPlugin = vlabs.NetworkPlugin
	api.NetworkMode = vlabs.NetworkMode
	api.ContainerRuntime = vlabs.ContainerRuntime
	api.MaxPods = vlabs.MaxPods
	api.DockerBridgeSubnet = vlabs.DockerBridgeSubnet
	api.MobyVersion = vlabs.MobyVersion
	api.ContainerdVersion = vlabs.ContainerdVersion
	api.CloudProviderBackoff = vlabs.CloudProviderBackoff
	api.CloudProviderBackoffMode = vlabs.CloudProviderBackoffMode
	api.CloudProviderBackoffDuration = vlabs.CloudProviderBackoffDuration
	api.CloudProviderBackoffExponent = vlabs.CloudProviderBackoffExponent
	api.CloudProviderBackoffJitter = vlabs.CloudProviderBackoffJitter
	api.CloudProviderBackoffRetries = vlabs.CloudProviderBackoffRetries
	api.CloudProviderRateLimit = vlabs.CloudProviderRateLimit
	api.CloudProviderRateLimitBucket = vlabs.CloudProviderRateLimitBucket
	api.CloudProviderRateLimitBucketWrite = vlabs.CloudProviderRateLimitBucketWrite
	api.CloudProviderRateLimitQPS = vlabs.CloudProviderRateLimitQPS
	api.CloudProviderRateLimitQPSWrite = vlabs.CloudProviderRateLimitQPSWrite
	api.UseManagedIdentity = vlabs.UseManagedIdentity
	api.UserAssignedID = vlabs.UserAssignedID
	api.UserAssignedClientID = vlabs.UserAssignedClientID
	api.CustomHyperkubeImage = vlabs.CustomHyperkubeImage
	api.CustomKubeAPIServerImage = vlabs.CustomKubeAPIServerImage
	api.CustomKubeControllerManagerImage = vlabs.CustomKubeControllerManagerImage
	api.CustomKubeProxyImage = vlabs.CustomKubeProxyImage
	api.CustomKubeSchedulerImage = vlabs.CustomKubeSchedulerImage
	api.CustomKubeBinaryURL = vlabs.CustomKubeBinaryURL
	api.CustomCcmImage = vlabs.CustomCcmImage
	api.UseCloudControllerManager = vlabs.UseCloudControllerManager
	api.CustomWindowsPackageURL = vlabs.CustomWindowsPackageURL
	api.WindowsNodeBinariesURL = vlabs.WindowsNodeBinariesURL
	api.WindowsContainerdURL = vlabs.WindowsContainerdURL
	api.WindowsSdnPluginURL = vlabs.WindowsSdnPluginURL
	api.UseInstanceMetadata = vlabs.UseInstanceMetadata
	api.LoadBalancerSku = vlabs.LoadBalancerSku
	api.ExcludeMasterFromStandardLB = vlabs.ExcludeMasterFromStandardLB
	api.LoadBalancerOutboundIPs = vlabs.LoadBalancerOutboundIPs
	api.EnableRbac = vlabs.EnableRbac
	api.EnableSecureKubelet = vlabs.EnableSecureKubelet
	api.EnableAggregatedAPIs = vlabs.EnableAggregatedAPIs
	api.EnableDataEncryptionAtRest = vlabs.EnableDataEncryptionAtRest
	api.EnableEncryptionWithExternalKms = vlabs.EnableEncryptionWithExternalKms
	api.EnablePodSecurityPolicy = vlabs.EnablePodSecurityPolicy
	api.GCHighThreshold = vlabs.GCHighThreshold
	api.GCLowThreshold = vlabs.GCLowThreshold
	api.EtcdVersion = vlabs.EtcdVersion
	api.EtcdDiskSizeGB = vlabs.EtcdDiskSizeGB
	api.EtcdEncryptionKey = vlabs.EtcdEncryptionKey
	api.AzureCNIVersion = vlabs.AzureCNIVersion
	api.AzureCNIURLLinux = vlabs.AzureCNIURLLinux
	api.AzureCNIURLWindows = vlabs.AzureCNIURLWindows
	api.KeyVaultSku = vlabs.KeyVaultSku
	api.MaximumLoadBalancerRuleCount = vlabs.MaximumLoadBalancerRuleCount
	api.ProxyMode = KubeProxyMode(vlabs.ProxyMode)
	api.PrivateAzureRegistryServer = vlabs.PrivateAzureRegistryServer
	api.OutboundRuleIdleTimeoutInMinutes = vlabs.OutboundRuleIdleTimeoutInMinutes
	api.CloudProviderDisableOutboundSNAT = vlabs.CloudProviderDisableOutboundSNAT
	api.KubeReservedCgroup = vlabs.KubeReservedCgroup
	convertComponentsToAPI(vlabs, api)
	convertAddonsToAPI(vlabs, api)
	convertKubeletConfigToAPI(vlabs, api)
	convertControllerManagerConfigToAPI(vlabs, api)
	convertCloudControllerManagerConfigToAPI(vlabs, api)
	convertAPIServerConfigToAPI(vlabs, api)
	convertSchedulerConfigToAPI(vlabs, api)
	convertPrivateClusterToAPI(vlabs, api)
	convertPodSecurityPolicyConfigToAPI(vlabs, api)
	convertContainerRuntimeConfigToAPI(vlabs, api)
}

func setVlabsKubernetesDefaults(vp *vlabs.Properties, api *OrchestratorProfile) {
	if api.KubernetesConfig == nil {
		api.KubernetesConfig = &KubernetesConfig{}
	}

	if vp.OrchestratorProfile.KubernetesConfig != nil {
		// Included here for backwards compatibility with deprecated NetworkPolicy usage patterns
		if vp.OrchestratorProfile.KubernetesConfig.NetworkPlugin == "" &&
			vp.OrchestratorProfile.KubernetesConfig.NetworkPolicy == NetworkPolicyAzure {
			api.KubernetesConfig.NetworkPlugin = vp.OrchestratorProfile.KubernetesConfig.NetworkPolicy
			api.KubernetesConfig.NetworkPolicy = "" // no-op but included for emphasis
		} else if vp.OrchestratorProfile.KubernetesConfig.NetworkPolicy == NetworkPolicyNone {
			api.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
			api.KubernetesConfig.NetworkPolicy = "" // no-op but included for emphasis
		} else {
			api.KubernetesConfig.NetworkPlugin = vp.OrchestratorProfile.KubernetesConfig.NetworkPlugin
			api.KubernetesConfig.NetworkPolicy = vp.OrchestratorProfile.KubernetesConfig.NetworkPolicy
		}
	}
	if api.KubernetesConfig.NetworkPlugin == "" && (api.KubernetesConfig.NetworkPolicy == "" || api.KubernetesConfig.NetworkPolicy == NetworkPolicyCalico) {
		if vp.HasWindows() {
			api.KubernetesConfig.NetworkPlugin = vlabs.DefaultNetworkPluginWindows
		} else {
			if vp.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.FlannelAddonName) {
				api.KubernetesConfig.NetworkPlugin = NetworkPluginFlannel
			} else {
				api.KubernetesConfig.NetworkPlugin = vlabs.DefaultNetworkPlugin
			}
		}
	}
}

func convertComponentsToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.Components = []KubernetesComponent{}
	for i := range v.Components {
		a.Components = append(a.Components, KubernetesComponent{
			Name:    v.Components[i].Name,
			Enabled: v.Components[i].Enabled,
			Config:  map[string]string{},
			Data:    v.Components[i].Data,
		})
		for j := range v.Components[i].Containers {
			a.Components[i].Containers = append(a.Components[i].Containers, KubernetesContainerSpec{
				Name:           v.Components[i].Containers[j].Name,
				Image:          v.Components[i].Containers[j].Image,
				CPURequests:    v.Components[i].Containers[j].CPURequests,
				MemoryRequests: v.Components[i].Containers[j].MemoryRequests,
				CPULimits:      v.Components[i].Containers[j].CPULimits,
				MemoryLimits:   v.Components[i].Containers[j].MemoryLimits,
			})
		}
		if v.Components[i].Config != nil {
			for key, val := range v.Components[i].Config {
				a.Components[i].Config[key] = val
			}
		}
	}
}

func convertAddonsToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.Addons = []KubernetesAddon{}
	for i := range v.Addons {
		a.Addons = append(a.Addons, KubernetesAddon{
			Name:    v.Addons[i].Name,
			Enabled: v.Addons[i].Enabled,
			Mode:    v.Addons[i].Mode,
			Config:  map[string]string{},
			Data:    v.Addons[i].Data,
		})
		for j := range v.Addons[i].Containers {
			a.Addons[i].Containers = append(a.Addons[i].Containers, KubernetesContainerSpec{
				Name:           v.Addons[i].Containers[j].Name,
				Image:          v.Addons[i].Containers[j].Image,
				CPURequests:    v.Addons[i].Containers[j].CPURequests,
				MemoryRequests: v.Addons[i].Containers[j].MemoryRequests,
				CPULimits:      v.Addons[i].Containers[j].CPULimits,
				MemoryLimits:   v.Addons[i].Containers[j].MemoryLimits,
			})
		}
		for k := range v.Addons[i].Pools {
			a.Addons[i].Pools = append(a.Addons[i].Pools, AddonNodePoolsConfig{
				Name:   v.Addons[i].Pools[k].Name,
				Config: map[string]string{},
			})
			if v.Addons[i].Pools[k].Config != nil {
				for key, val := range v.Addons[i].Pools[k].Config {
					a.Addons[i].Pools[k].Config[key] = val
				}
			}
		}
		if v.Addons[i].Config != nil {
			for key, val := range v.Addons[i].Config {
				a.Addons[i].Config[key] = val
			}
		}
	}
}

func convertCustomFilesToAPI(v *vlabs.MasterProfile, a *MasterProfile) {
	if v.CustomFiles != nil {
		a.CustomFiles = &[]CustomFile{}
		for i := range *v.CustomFiles {
			*a.CustomFiles = append(*a.CustomFiles, CustomFile{
				Dest:   (*v.CustomFiles)[i].Dest,
				Source: (*v.CustomFiles)[i].Source,
			})
		}
	}
}

func convertContainerRuntimeConfigToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.ContainerRuntimeConfig = map[string]string{}
	for key, val := range v.ContainerRuntimeConfig {
		a.ContainerRuntimeConfig[key] = val
	}
}

func convertKubeletConfigToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.KubeletConfig = map[string]string{}
	for key, val := range v.KubeletConfig {
		a.KubeletConfig[key] = val
	}
}

func convertControllerManagerConfigToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.ControllerManagerConfig = map[string]string{}
	for key, val := range v.ControllerManagerConfig {
		a.ControllerManagerConfig[key] = val
	}
}

func convertCloudControllerManagerConfigToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.CloudControllerManagerConfig = map[string]string{}
	for key, val := range v.CloudControllerManagerConfig {
		a.CloudControllerManagerConfig[key] = val
	}
}

func convertAPIServerConfigToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.APIServerConfig = map[string]string{}
	for key, val := range v.APIServerConfig {
		a.APIServerConfig[key] = val
	}
}

func convertSchedulerConfigToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.SchedulerConfig = map[string]string{}
	for key, val := range v.SchedulerConfig {
		a.SchedulerConfig[key] = val
	}
}

func convertPodSecurityPolicyConfigToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.PodSecurityPolicyConfig = map[string]string{}
	for key, val := range v.PodSecurityPolicyConfig {
		a.PodSecurityPolicyConfig[key] = val
	}
}

func convertPrivateClusterToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	if v.PrivateCluster != nil {
		a.PrivateCluster = &PrivateCluster{}
		a.PrivateCluster.Enabled = v.PrivateCluster.Enabled
		if v.PrivateCluster.JumpboxProfile != nil {
			a.PrivateCluster.JumpboxProfile = &PrivateJumpboxProfile{}
			convertPrivateJumpboxProfileToAPI(v.PrivateCluster.JumpboxProfile, a.PrivateCluster.JumpboxProfile)
		}
	}
}

func convertPrivateJumpboxProfileToAPI(v *vlabs.PrivateJumpboxProfile, a *PrivateJumpboxProfile) {
	a.Name = v.Name
	a.OSDiskSizeGB = v.OSDiskSizeGB
	a.VMSize = v.VMSize
	a.PublicKey = v.PublicKey
	a.Username = v.Username
	a.StorageProfile = v.StorageProfile
}

func convertVLabsMasterProfile(vlabs *vlabs.MasterProfile, api *MasterProfile) {
	api.Count = vlabs.Count
	api.DNSPrefix = vlabs.DNSPrefix
	api.SubjectAltNames = vlabs.SubjectAltNames
	api.VMSize = vlabs.VMSize
	api.CustomVMTags = vlabs.CustomVMTags
	api.OSDiskSizeGB = vlabs.OSDiskSizeGB
	api.VnetSubnetID = vlabs.VnetSubnetID
	api.AgentVnetSubnetID = vlabs.AgentVnetSubnetID
	api.FirstConsecutiveStaticIP = vlabs.FirstConsecutiveStaticIP
	api.VnetCidr = vlabs.VnetCidr
	api.Subnet = vlabs.GetSubnet()
	api.SubnetIPv6 = vlabs.GetSubnetIPv6()
	api.IPAddressCount = vlabs.IPAddressCount
	api.FQDN = vlabs.FQDN
	api.StorageProfile = vlabs.StorageProfile
	api.HTTPSourceAddressPrefix = vlabs.HTTPSourceAddressPrefix
	api.OAuthEnabled = vlabs.OAuthEnabled
	// by default vlabs will use managed disks as it has encryption at rest
	if len(api.StorageProfile) == 0 {
		api.StorageProfile = ManagedDisks
	}

	if vlabs.PreProvisionExtension != nil {
		apiExtension := &Extension{}
		convertVLabsExtension(vlabs.PreProvisionExtension, apiExtension)
		api.PreprovisionExtension = apiExtension
	}

	api.Extensions = []Extension{}
	for _, extension := range vlabs.Extensions {
		apiExtension := &Extension{}
		convertVLabsExtension(&extension, apiExtension)
		api.Extensions = append(api.Extensions, *apiExtension)
	}

	api.Distro = Distro(vlabs.Distro)
	if vlabs.KubernetesConfig != nil {
		api.KubernetesConfig = &KubernetesConfig{}
		convertVLabsKubernetesConfig(vlabs.KubernetesConfig, api.KubernetesConfig)
	}
	if vlabs.ImageRef != nil {
		api.ImageRef = &ImageReference{}
		api.ImageRef.Name = vlabs.ImageRef.Name
		api.ImageRef.ResourceGroup = vlabs.ImageRef.ResourceGroup
		api.ImageRef.SubscriptionID = vlabs.ImageRef.SubscriptionID
		api.ImageRef.Gallery = vlabs.ImageRef.Gallery
		api.ImageRef.Version = vlabs.ImageRef.Version
	}

	api.AvailabilityProfile = vlabs.AvailabilityProfile
	api.AgentSubnet = vlabs.AgentSubnet
	api.AvailabilityZones = vlabs.AvailabilityZones
	api.PlatformFaultDomainCount = vlabs.PlatformFaultDomainCount
	api.PlatformUpdateDomainCount = vlabs.PlatformUpdateDomainCount
	api.SinglePlacementGroup = vlabs.SinglePlacementGroup
	api.CosmosEtcd = vlabs.CosmosEtcd
	api.UltraSSDEnabled = vlabs.UltraSSDEnabled
	api.EncryptionAtHost = vlabs.EncryptionAtHost
	api.AuditDEnabled = vlabs.AuditDEnabled
	api.ProximityPlacementGroupID = vlabs.ProximityPlacementGroupID
	api.OSDiskCachingType = vlabs.OSDiskCachingType
	convertCustomFilesToAPI(vlabs, api)
	api.SysctlDConfig = map[string]string{}
	for key, val := range vlabs.SysctlDConfig {
		api.SysctlDConfig[key] = val
	}
}

func convertVLabsAgentPoolProfile(vlabs *vlabs.AgentPoolProfile, api *AgentPoolProfile) {
	api.Name = vlabs.Name
	api.Count = vlabs.Count
	api.VMSize = vlabs.VMSize
	api.CustomVMTags = vlabs.CustomVMTags
	api.OSDiskSizeGB = vlabs.OSDiskSizeGB
	api.DNSPrefix = vlabs.DNSPrefix
	api.OSType = OSType(vlabs.OSType)
	api.Ports = []int{}
	api.Ports = append(api.Ports, vlabs.Ports...)
	api.AvailabilityProfile = vlabs.AvailabilityProfile
	api.ScaleSetPriority = vlabs.ScaleSetPriority
	api.ScaleSetEvictionPolicy = vlabs.ScaleSetEvictionPolicy
	api.SpotMaxPrice = vlabs.SpotMaxPrice
	api.StorageProfile = vlabs.StorageProfile
	api.DiskSizesGB = []int{}
	api.DiskSizesGB = append(api.DiskSizesGB, vlabs.DiskSizesGB...)
	api.VnetSubnetID = vlabs.VnetSubnetID
	api.Subnet = vlabs.GetSubnet()
	api.IPAddressCount = vlabs.IPAddressCount
	api.FQDN = vlabs.FQDN
	api.AcceleratedNetworkingEnabled = vlabs.AcceleratedNetworkingEnabled
	api.AcceleratedNetworkingEnabledWindows = vlabs.AcceleratedNetworkingEnabledWindows
	api.VMSSOverProvisioningEnabled = vlabs.VMSSOverProvisioningEnabled
	api.AvailabilityZones = vlabs.AvailabilityZones
	api.PlatformFaultDomainCount = vlabs.PlatformFaultDomainCount
	api.PlatformUpdateDomainCount = vlabs.PlatformUpdateDomainCount
	api.SinglePlacementGroup = vlabs.SinglePlacementGroup
	api.EnableVMSSNodePublicIP = vlabs.EnableVMSSNodePublicIP
	api.LoadBalancerBackendAddressPoolIDs = vlabs.LoadBalancerBackendAddressPoolIDs
	api.AuditDEnabled = vlabs.AuditDEnabled
	api.DiskEncryptionSetID = vlabs.DiskEncryptionSetID
	api.UltraSSDEnabled = vlabs.UltraSSDEnabled
	api.EncryptionAtHost = vlabs.EncryptionAtHost
	api.ProximityPlacementGroupID = vlabs.ProximityPlacementGroupID

	api.CustomNodeLabels = map[string]string{}
	for k, v := range vlabs.CustomNodeLabels {
		api.CustomNodeLabels[k] = v
	}

	if vlabs.PreProvisionExtension != nil {
		apiExtension := &Extension{}
		convertVLabsExtension(vlabs.PreProvisionExtension, apiExtension)
		api.PreprovisionExtension = apiExtension
	}

	api.Extensions = []Extension{}
	for _, extension := range vlabs.Extensions {
		apiExtension := &Extension{}
		convertVLabsExtension(&extension, apiExtension)
		api.Extensions = append(api.Extensions, *apiExtension)
	}
	api.Distro = Distro(vlabs.Distro)
	if vlabs.KubernetesConfig != nil {
		api.KubernetesConfig = &KubernetesConfig{}
		convertVLabsKubernetesConfig(vlabs.KubernetesConfig, api.KubernetesConfig)
	}
	if vlabs.ImageRef != nil {
		api.ImageRef = &ImageReference{}
		api.ImageRef.Name = vlabs.ImageRef.Name
		api.ImageRef.ResourceGroup = vlabs.ImageRef.ResourceGroup
		api.ImageRef.SubscriptionID = vlabs.ImageRef.SubscriptionID
		api.ImageRef.Gallery = vlabs.ImageRef.Gallery
		api.ImageRef.Version = vlabs.ImageRef.Version
	}
	api.Role = AgentPoolProfileRole(vlabs.Role)
	api.SysctlDConfig = map[string]string{}
	for key, val := range vlabs.SysctlDConfig {
		api.SysctlDConfig[key] = val
	}
	api.OSDiskCachingType = vlabs.OSDiskCachingType
	api.DataDiskCachingType = vlabs.DataDiskCachingType
}

func convertVLabsKeyVaultSecrets(vlabs *vlabs.KeyVaultSecrets, api *KeyVaultSecrets) {
	api.SourceVault = &KeyVaultID{ID: vlabs.SourceVault.ID}
	api.VaultCertificates = []KeyVaultCertificate{}
	for _, c := range vlabs.VaultCertificates {
		cert := KeyVaultCertificate{}
		cert.CertificateStore = c.CertificateStore
		cert.CertificateURL = c.CertificateURL
		api.VaultCertificates = append(api.VaultCertificates, cert)
	}
}

func convertVLabsServicePrincipalProfile(vlabs *vlabs.ServicePrincipalProfile, api *ServicePrincipalProfile) {
	api.ClientID = vlabs.ClientID
	api.Secret = vlabs.Secret
	api.ObjectID = vlabs.ObjectID
	if vlabs.KeyvaultSecretRef != nil {
		api.KeyvaultSecretRef = &KeyvaultSecretRef{
			VaultID:       vlabs.KeyvaultSecretRef.VaultID,
			SecretName:    vlabs.KeyvaultSecretRef.SecretName,
			SecretVersion: vlabs.KeyvaultSecretRef.SecretVersion,
		}
	}
}

func convertVLabsCertificateProfile(vlabs *vlabs.CertificateProfile, api *CertificateProfile) {
	api.CaCertificate = vlabs.CaCertificate
	api.CaPrivateKey = vlabs.CaPrivateKey
	api.APIServerCertificate = vlabs.APIServerCertificate
	api.APIServerPrivateKey = vlabs.APIServerPrivateKey
	api.ClientCertificate = vlabs.ClientCertificate
	api.ClientPrivateKey = vlabs.ClientPrivateKey
	api.KubeConfigCertificate = vlabs.KubeConfigCertificate
	api.KubeConfigPrivateKey = vlabs.KubeConfigPrivateKey
	api.EtcdServerCertificate = vlabs.EtcdServerCertificate
	api.EtcdServerPrivateKey = vlabs.EtcdServerPrivateKey
	api.EtcdClientCertificate = vlabs.EtcdClientCertificate
	api.EtcdClientPrivateKey = vlabs.EtcdClientPrivateKey
	api.EtcdPeerCertificates = vlabs.EtcdPeerCertificates
	api.EtcdPeerPrivateKeys = vlabs.EtcdPeerPrivateKeys
}

func convertVLabsAADProfile(vlabs *vlabs.AADProfile, api *AADProfile) {
	api.ClientAppID = vlabs.ClientAppID
	api.ServerAppID = vlabs.ServerAppID
	api.TenantID = vlabs.TenantID
	api.AdminGroupID = vlabs.AdminGroupID
	api.Authenticator = OIDC
}

func addDCOSPublicAgentPool(api *Properties) {
	publicPool := &AgentPoolProfile{}
	// tag this agent pool with a known suffix string
	publicPool.Name = api.AgentPoolProfiles[0].Name + publicAgentPoolSuffix
	// move DNS prefix to public pool
	publicPool.DNSPrefix = api.AgentPoolProfiles[0].DNSPrefix
	api.AgentPoolProfiles[0].DNSPrefix = ""
	publicPool.VMSize = api.AgentPoolProfiles[0].VMSize // - use same VMsize for public pool
	publicPool.OSType = api.AgentPoolProfiles[0].OSType // - use same OSType for public pool
	api.AgentPoolProfiles[0].Ports = nil
	for _, port := range [3]int{80, 443, 8080} {
		publicPool.Ports = append(publicPool.Ports, port)
	}
	// - VM Count for public agents is based on the following:
	// 1 master => 1 VM
	// 3, 5 master => 3 VMsize
	if api.MasterProfile.Count == 1 {
		publicPool.Count = 1
	} else {
		publicPool.Count = 3
	}
	api.AgentPoolProfiles = append(api.AgentPoolProfiles, publicPool)
}

func convertVLabsCustomCloudProfile(vlabs *vlabs.CustomCloudProfile, api *CustomCloudProfile) {
	if vlabs.Environment != nil {
		api.Environment = &azure.Environment{}
		api.Environment.Name = vlabs.Environment.Name
		api.Environment.ManagementPortalURL = vlabs.Environment.ManagementPortalURL
		api.Environment.PublishSettingsURL = vlabs.Environment.PublishSettingsURL
		api.Environment.ServiceManagementEndpoint = vlabs.Environment.ServiceManagementEndpoint
		api.Environment.ResourceManagerEndpoint = vlabs.Environment.ResourceManagerEndpoint
		api.Environment.ActiveDirectoryEndpoint = vlabs.Environment.ActiveDirectoryEndpoint
		api.Environment.GalleryEndpoint = vlabs.Environment.GalleryEndpoint
		api.Environment.KeyVaultEndpoint = vlabs.Environment.KeyVaultEndpoint
		api.Environment.GraphEndpoint = vlabs.Environment.GraphEndpoint
		api.Environment.ServiceBusEndpoint = vlabs.Environment.ServiceBusEndpoint
		api.Environment.BatchManagementEndpoint = vlabs.Environment.BatchManagementEndpoint
		api.Environment.StorageEndpointSuffix = vlabs.Environment.StorageEndpointSuffix
		api.Environment.SQLDatabaseDNSSuffix = vlabs.Environment.SQLDatabaseDNSSuffix
		api.Environment.TrafficManagerDNSSuffix = vlabs.Environment.TrafficManagerDNSSuffix
		api.Environment.KeyVaultDNSSuffix = vlabs.Environment.KeyVaultDNSSuffix
		api.Environment.ServiceBusEndpointSuffix = vlabs.Environment.ServiceBusEndpointSuffix
		api.Environment.ServiceManagementVMDNSSuffix = vlabs.Environment.ServiceManagementVMDNSSuffix
		api.Environment.ResourceManagerVMDNSSuffix = vlabs.Environment.ResourceManagerVMDNSSuffix
		api.Environment.ContainerRegistryDNSSuffix = vlabs.Environment.ContainerRegistryDNSSuffix
		api.Environment.TokenAudience = vlabs.Environment.TokenAudience
	}
	if vlabs.AzureEnvironmentSpecConfig != nil {
		api.AzureEnvironmentSpecConfig = &AzureEnvironmentSpecConfig{}
		convertAzureEnvironmentSpecConfig(vlabs.AzureEnvironmentSpecConfig, api.AzureEnvironmentSpecConfig)
	}

	api.IdentitySystem = vlabs.IdentitySystem
	api.AuthenticationMethod = vlabs.AuthenticationMethod
	api.DependenciesLocation = DependenciesLocation(vlabs.DependenciesLocation)
	api.PortalURL = vlabs.PortalURL
	api.CustomCloudRootCertificates = vlabs.CustomCloudRootCertificates
	api.CustomCloudSourcesList = vlabs.CustomCloudSourcesList
}

func convertVLabsTelemetryProfile(vlabs *vlabs.TelemetryProfile, api *TelemetryProfile) {
	api.ApplicationInsightsKey = vlabs.ApplicationInsightsKey
}

func convertAzureEnvironmentSpecConfig(vlabses *vlabs.AzureEnvironmentSpecConfig, api *AzureEnvironmentSpecConfig) {
	api.CloudName = vlabses.CloudName
	api.DCOSSpecConfig = DCOSSpecConfig{
		DCOS188BootstrapDownloadURL:     vlabses.DCOSSpecConfig.DCOS188BootstrapDownloadURL,
		DCOS190BootstrapDownloadURL:     vlabses.DCOSSpecConfig.DCOS190BootstrapDownloadURL,
		DCOS198BootstrapDownloadURL:     vlabses.DCOSSpecConfig.DCOS198BootstrapDownloadURL,
		DCOS110BootstrapDownloadURL:     vlabses.DCOSSpecConfig.DCOS110BootstrapDownloadURL,
		DCOS111BootstrapDownloadURL:     vlabses.DCOSSpecConfig.DCOS111BootstrapDownloadURL,
		DCOSWindowsBootstrapDownloadURL: vlabses.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL,
		DcosRepositoryURL:               vlabses.DCOSSpecConfig.DcosRepositoryURL,
		DcosClusterPackageListID:        vlabses.DCOSSpecConfig.DcosClusterPackageListID,
		DcosProviderPackageID:           vlabses.DCOSSpecConfig.DcosProviderPackageID,
	}

	api.DockerSpecConfig = DockerSpecConfig{
		DockerEngineRepo:         vlabses.DockerSpecConfig.DockerEngineRepo,
		DockerComposeDownloadURL: vlabses.DockerSpecConfig.DockerComposeDownloadURL,
	}
	api.EndpointConfig = AzureEndpointConfig{
		ResourceManagerVMDNSSuffix: vlabses.EndpointConfig.ResourceManagerVMDNSSuffix,
	}
	api.KubernetesSpecConfig = KubernetesSpecConfig{
		AzureTelemetryPID:                vlabses.KubernetesSpecConfig.AzureTelemetryPID,
		KubernetesImageBase:              vlabses.KubernetesSpecConfig.KubernetesImageBase,
		TillerImageBase:                  vlabses.KubernetesSpecConfig.TillerImageBase,
		ACIConnectorImageBase:            vlabses.KubernetesSpecConfig.ACIConnectorImageBase,
		NVIDIAImageBase:                  vlabses.KubernetesSpecConfig.NVIDIAImageBase,
		AzureCNIImageBase:                vlabses.KubernetesSpecConfig.AzureCNIImageBase,
		CalicoImageBase:                  vlabses.KubernetesSpecConfig.CalicoImageBase,
		EtcdDownloadURLBase:              vlabses.KubernetesSpecConfig.EtcdDownloadURLBase,
		KubeBinariesSASURLBase:           vlabses.KubernetesSpecConfig.KubeBinariesSASURLBase,
		WindowsTelemetryGUID:             vlabses.KubernetesSpecConfig.WindowsTelemetryGUID,
		CNIPluginsDownloadURL:            vlabses.KubernetesSpecConfig.CNIPluginsDownloadURL,
		VnetCNILinuxPluginsDownloadURL:   vlabses.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL,
		VnetCNIWindowsPluginsDownloadURL: vlabses.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL,
		ContainerdDownloadURLBase:        vlabses.KubernetesSpecConfig.ContainerdDownloadURLBase,
		CSIProxyDownloadURL:              vlabses.KubernetesSpecConfig.CSIProxyDownloadURL,
	}
	api.OSImageConfig = map[Distro]AzureOSImageConfig{}
	for k, v := range vlabses.OSImageConfig {
		api.OSImageConfig[Distro(string(k))] = AzureOSImageConfig{
			ImageOffer:     v.ImageOffer,
			ImageSku:       v.ImageSku,
			ImagePublisher: v.ImagePublisher,
			ImageVersion:   v.ImageVersion,
		}
	}
}
