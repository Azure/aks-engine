// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/v20160330"
	"github.com/Azure/aks-engine/pkg/api/v20160930"
	"github.com/Azure/aks-engine/pkg/api/v20170131"
	"github.com/Azure/aks-engine/pkg/api/v20170701"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/helpers"
)

///////////////////////////////////////////////////////////
// The converter exposes functions to convert the top level
// ContainerService resource
//
// All other functions are internal helper functions used
// for converting.
///////////////////////////////////////////////////////////

// ConvertV20160930ContainerService converts a v20160930 ContainerService to an unversioned ContainerService
func ConvertV20160930ContainerService(v20160930 *v20160930.ContainerService) *ContainerService {
	c := &ContainerService{}
	c.ID = v20160930.ID
	c.Location = helpers.NormalizeAzureRegion(v20160930.Location)
	c.Name = v20160930.Name
	if v20160930.Plan != nil {
		c.Plan = &ResourcePurchasePlan{}
		convertV20160930ResourcePurchasePlan(v20160930.Plan, c.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range v20160930.Tags {
		c.Tags[k] = v
	}
	c.Type = v20160930.Type
	c.Properties = &Properties{}
	convertV20160930Properties(v20160930.Properties, c.Properties)
	return c
}

// ConvertV20160330ContainerService converts a v20160330 ContainerService to an unversioned ContainerService
func ConvertV20160330ContainerService(v20160330 *v20160330.ContainerService) *ContainerService {
	c := &ContainerService{}
	c.ID = v20160330.ID
	c.Location = helpers.NormalizeAzureRegion(v20160330.Location)
	c.Name = v20160330.Name
	if v20160330.Plan != nil {
		c.Plan = &ResourcePurchasePlan{}
		convertV20160330ResourcePurchasePlan(v20160330.Plan, c.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range v20160330.Tags {
		c.Tags[k] = v
	}
	c.Type = v20160330.Type
	c.Properties = &Properties{}
	convertV20160330Properties(v20160330.Properties, c.Properties)
	return c
}

// ConvertV20170131ContainerService converts a v20170131 ContainerService to an unversioned ContainerService
func ConvertV20170131ContainerService(v20170131 *v20170131.ContainerService) *ContainerService {
	c := &ContainerService{}
	c.ID = v20170131.ID
	c.Location = helpers.NormalizeAzureRegion(v20170131.Location)
	c.Name = v20170131.Name
	if v20170131.Plan != nil {
		c.Plan = &ResourcePurchasePlan{}
		convertV20170131ResourcePurchasePlan(v20170131.Plan, c.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range v20170131.Tags {
		c.Tags[k] = v
	}
	c.Type = v20170131.Type
	c.Properties = &Properties{}
	convertV20170131Properties(v20170131.Properties, c.Properties)
	return c
}

// ConvertV20170701ContainerService converts a v20170701 ContainerService to an unversioned ContainerService
func ConvertV20170701ContainerService(v20170701 *v20170701.ContainerService) *ContainerService {
	c := &ContainerService{}
	c.ID = v20170701.ID
	c.Location = helpers.NormalizeAzureRegion(v20170701.Location)
	c.Name = v20170701.Name
	if v20170701.Plan != nil {
		c.Plan = &ResourcePurchasePlan{}
		convertV20170701ResourcePurchasePlan(v20170701.Plan, c.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range v20170701.Tags {
		c.Tags[k] = v
	}
	c.Type = v20170701.Type
	c.Properties = &Properties{}
	convertV20170701Properties(v20170701.Properties, c.Properties)
	return c
}

// ConvertVLabsContainerService converts a vlabs ContainerService to an unversioned ContainerService
func ConvertVLabsContainerService(vlabs *vlabs.ContainerService, isUpdate bool) *ContainerService {
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
	convertVLabsProperties(vlabs.Properties, c.Properties, isUpdate)
	return c
}

// convertV20160930ResourcePurchasePlan converts a v20160930 ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertV20160930ResourcePurchasePlan(v20160930 *v20160930.ResourcePurchasePlan, api *ResourcePurchasePlan) {
	api.Name = v20160930.Name
	api.Product = v20160930.Product
	api.PromotionCode = v20160930.PromotionCode
	api.Publisher = v20160930.Publisher
}

// convertV20160330ResourcePurchasePlan converts a v20160330 ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertV20160330ResourcePurchasePlan(v20160330 *v20160330.ResourcePurchasePlan, api *ResourcePurchasePlan) {
	api.Name = v20160330.Name
	api.Product = v20160330.Product
	api.PromotionCode = v20160330.PromotionCode
	api.Publisher = v20160330.Publisher
}

// convertV20170131ResourcePurchasePlan converts a v20170131 ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertV20170131ResourcePurchasePlan(v20170131 *v20170131.ResourcePurchasePlan, api *ResourcePurchasePlan) {
	api.Name = v20170131.Name
	api.Product = v20170131.Product
	api.PromotionCode = v20170131.PromotionCode
	api.Publisher = v20170131.Publisher
}

// convertV20170701ResourcePurchasePlan converts a v20170701 ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertV20170701ResourcePurchasePlan(v20170701 *v20170701.ResourcePurchasePlan, api *ResourcePurchasePlan) {
	api.Name = v20170701.Name
	api.Product = v20170701.Product
	api.PromotionCode = v20170701.PromotionCode
	api.Publisher = v20170701.Publisher
}

// convertVLabsResourcePurchasePlan converts a vlabs ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertVLabsResourcePurchasePlan(vlabs *vlabs.ResourcePurchasePlan, api *ResourcePurchasePlan) {
	api.Name = vlabs.Name
	api.Product = vlabs.Product
	api.PromotionCode = vlabs.PromotionCode
	api.Publisher = vlabs.Publisher
}

func convertV20160930Properties(v20160930 *v20160930.Properties, api *Properties) {
	api.ProvisioningState = ProvisioningState(v20160930.ProvisioningState)
	if v20160930.OrchestratorProfile != nil {
		api.OrchestratorProfile = &OrchestratorProfile{}
		convertV20160930OrchestratorProfile(v20160930.OrchestratorProfile, api.OrchestratorProfile)
	}
	if v20160930.MasterProfile != nil {
		api.MasterProfile = &MasterProfile{}
		convertV20160930MasterProfile(v20160930.MasterProfile, api.MasterProfile)
	}
	api.AgentPoolProfiles = []*AgentPoolProfile{}
	for _, p := range v20160930.AgentPoolProfiles {
		apiProfile := &AgentPoolProfile{}
		// api.OrchestratorProfile already be filled in correctly
		if api.OrchestratorProfile.IsKubernetes() {
			// we only allow AvailabilitySet for kubernetes's agentpool
			convertV20160930AgentPoolProfile(p, AvailabilitySet, apiProfile)
		} else {
			// other orchestrators all use VMSS
			convertV20160930AgentPoolProfile(p, VirtualMachineScaleSets, apiProfile)
		}
		api.AgentPoolProfiles = append(api.AgentPoolProfiles, apiProfile)
	}
	if v20160930.LinuxProfile != nil {
		api.LinuxProfile = &LinuxProfile{}
		convertV20160930LinuxProfile(v20160930.LinuxProfile, api.LinuxProfile)
	}
	if v20160930.WindowsProfile != nil {
		api.WindowsProfile = &WindowsProfile{}
		convertV20160930WindowsProfile(v20160930.WindowsProfile, api.WindowsProfile)
	}
	if v20160930.DiagnosticsProfile != nil {
		api.DiagnosticsProfile = &DiagnosticsProfile{}
		convertV20160930DiagnosticsProfile(v20160930.DiagnosticsProfile, api.DiagnosticsProfile)
	}
	if v20160930.JumpboxProfile != nil {
		api.JumpboxProfile = &JumpboxProfile{}
		convertV20160930JumpboxProfile(v20160930.JumpboxProfile, api.JumpboxProfile)
	}
	if v20160930.ServicePrincipalProfile != nil {
		api.ServicePrincipalProfile = &ServicePrincipalProfile{}
		convertV20160930ServicePrincipalProfile(v20160930.ServicePrincipalProfile, api.ServicePrincipalProfile)
	}
	if v20160930.CustomProfile != nil {
		api.CustomProfile = &CustomProfile{}
		convertV20160930CustomProfile(v20160930.CustomProfile, api.CustomProfile)
	}
	if api.OrchestratorProfile.IsDCOS() && len(api.AgentPoolProfiles) == 1 {
		addDCOSPublicAgentPool(api)
	}
}

func convertV20160330Properties(v20160330 *v20160330.Properties, api *Properties) {
	api.ProvisioningState = ProvisioningState(v20160330.ProvisioningState)
	if v20160330.OrchestratorProfile != nil {
		api.OrchestratorProfile = &OrchestratorProfile{}
		convertV20160330OrchestratorProfile(v20160330.OrchestratorProfile, api.OrchestratorProfile)
	}
	if v20160330.MasterProfile != nil {
		api.MasterProfile = &MasterProfile{}
		convertV20160330MasterProfile(v20160330.MasterProfile, api.MasterProfile)
	}
	api.AgentPoolProfiles = []*AgentPoolProfile{}
	for _, p := range v20160330.AgentPoolProfiles {
		apiProfile := &AgentPoolProfile{}
		convertV20160330AgentPoolProfile(p, apiProfile)
		api.AgentPoolProfiles = append(api.AgentPoolProfiles, apiProfile)

	}
	if v20160330.LinuxProfile != nil {
		api.LinuxProfile = &LinuxProfile{}
		convertV20160330LinuxProfile(v20160330.LinuxProfile, api.LinuxProfile)
	}
	if v20160330.WindowsProfile != nil {
		api.WindowsProfile = &WindowsProfile{}
		convertV20160330WindowsProfile(v20160330.WindowsProfile, api.WindowsProfile)
	}
	if v20160330.DiagnosticsProfile != nil {
		api.DiagnosticsProfile = &DiagnosticsProfile{}
		convertV20160330DiagnosticsProfile(v20160330.DiagnosticsProfile, api.DiagnosticsProfile)
	}
	if v20160330.JumpboxProfile != nil {
		api.JumpboxProfile = &JumpboxProfile{}
		convertV20160330JumpboxProfile(v20160330.JumpboxProfile, api.JumpboxProfile)
	}
	if api.OrchestratorProfile.IsDCOS() && len(api.AgentPoolProfiles) == 1 {
		addDCOSPublicAgentPool(api)
	}
}

func convertV20170131Properties(v20170131 *v20170131.Properties, api *Properties) {
	api.ProvisioningState = ProvisioningState(v20170131.ProvisioningState)
	if v20170131.OrchestratorProfile != nil {
		api.OrchestratorProfile = &OrchestratorProfile{}
		convertV20170131OrchestratorProfile(v20170131.OrchestratorProfile, api.OrchestratorProfile, v20170131.HasWindows())
	}
	if v20170131.MasterProfile != nil {
		api.MasterProfile = &MasterProfile{}
		convertV20170131MasterProfile(v20170131.MasterProfile, api.MasterProfile)
	}
	api.AgentPoolProfiles = []*AgentPoolProfile{}
	for _, p := range v20170131.AgentPoolProfiles {
		apiProfile := &AgentPoolProfile{}
		// api.OrchestratorProfile already be filled in correctly
		if api.OrchestratorProfile.IsKubernetes() {
			// we only allow AvailabilitySet for kubernetes's agentpool
			convertV20170131AgentPoolProfile(p, AvailabilitySet, apiProfile)
		} else {
			// other orchestrators all use VMSS
			convertV20170131AgentPoolProfile(p, VirtualMachineScaleSets, apiProfile)
		}
		api.AgentPoolProfiles = append(api.AgentPoolProfiles, apiProfile)
	}
	if v20170131.LinuxProfile != nil {
		api.LinuxProfile = &LinuxProfile{}
		convertV20170131LinuxProfile(v20170131.LinuxProfile, api.LinuxProfile)
	}
	if v20170131.WindowsProfile != nil {
		api.WindowsProfile = &WindowsProfile{}
		convertV20170131WindowsProfile(v20170131.WindowsProfile, api.WindowsProfile)
	}
	if v20170131.DiagnosticsProfile != nil {
		api.DiagnosticsProfile = &DiagnosticsProfile{}
		convertV20170131DiagnosticsProfile(v20170131.DiagnosticsProfile, api.DiagnosticsProfile)
	}
	if v20170131.JumpboxProfile != nil {
		api.JumpboxProfile = &JumpboxProfile{}
		convertV20170131JumpboxProfile(v20170131.JumpboxProfile, api.JumpboxProfile)
	}
	if v20170131.ServicePrincipalProfile != nil {
		api.ServicePrincipalProfile = &ServicePrincipalProfile{}
		convertV20170131ServicePrincipalProfile(v20170131.ServicePrincipalProfile, api.ServicePrincipalProfile)
	}
	if v20170131.CustomProfile != nil {
		api.CustomProfile = &CustomProfile{}
		convertV20170131CustomProfile(v20170131.CustomProfile, api.CustomProfile)
	}
	if api.OrchestratorProfile.IsDCOS() && len(api.AgentPoolProfiles) == 1 {
		addDCOSPublicAgentPool(api)
	}
}

func convertV20170701Properties(v20170701 *v20170701.Properties, api *Properties) {
	api.ProvisioningState = ProvisioningState(v20170701.ProvisioningState)

	if v20170701.OrchestratorProfile != nil {
		api.OrchestratorProfile = &OrchestratorProfile{}
		convertV20170701OrchestratorProfile(v20170701.OrchestratorProfile, api.OrchestratorProfile, v20170701.HasWindows())
	}
	if v20170701.MasterProfile != nil {
		api.MasterProfile = &MasterProfile{}
		convertV20170701MasterProfile(v20170701.MasterProfile, api.MasterProfile)
	}
	api.AgentPoolProfiles = []*AgentPoolProfile{}
	for _, p := range v20170701.AgentPoolProfiles {
		apiProfile := &AgentPoolProfile{}
		// api.OrchestratorProfile already be filled in correctly
		if api.OrchestratorProfile.IsKubernetes() {
			// we only allow AvailabilitySet for kubernetes's agentpool
			convertV20170701AgentPoolProfile(p, AvailabilitySet, apiProfile)
		} else {
			// other orchestrators all use VMSS
			convertV20170701AgentPoolProfile(p, VirtualMachineScaleSets, apiProfile)
			// by default vlabs will use managed disks for all orchestrators but kubernetes as it has encryption at rest.
			if len(p.StorageProfile) == 0 {
				apiProfile.StorageProfile = ManagedDisks
			}
		}
		api.AgentPoolProfiles = append(api.AgentPoolProfiles, apiProfile)
	}
	if v20170701.LinuxProfile != nil {
		api.LinuxProfile = &LinuxProfile{}
		convertV20170701LinuxProfile(v20170701.LinuxProfile, api.LinuxProfile)
	}
	if v20170701.WindowsProfile != nil {
		api.WindowsProfile = &WindowsProfile{}
		convertV20170701WindowsProfile(v20170701.WindowsProfile, api.WindowsProfile)
	}
	if v20170701.ServicePrincipalProfile != nil {
		api.ServicePrincipalProfile = &ServicePrincipalProfile{}
		convertV20170701ServicePrincipalProfile(v20170701.ServicePrincipalProfile, api.ServicePrincipalProfile)
	}
	if v20170701.CustomProfile != nil {
		api.CustomProfile = &CustomProfile{}
		convertV20170701CustomProfile(v20170701.CustomProfile, api.CustomProfile)
	}
}

func convertVLabsProperties(vlabs *vlabs.Properties, api *Properties, isUpdate bool) {
	api.ProvisioningState = ProvisioningState(vlabs.ProvisioningState)
	if vlabs.OrchestratorProfile != nil {
		api.OrchestratorProfile = &OrchestratorProfile{}
		convertVLabsOrchestratorProfile(vlabs, api.OrchestratorProfile, isUpdate)
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
}

func convertVLabsFeatureFlags(vlabs *vlabs.FeatureFlags, api *FeatureFlags) {
	api.EnableCSERunInBackground = vlabs.EnableCSERunInBackground
	api.BlockOutboundInternet = vlabs.BlockOutboundInternet
}

func convertV20160930LinuxProfile(obj *v20160930.LinuxProfile, api *LinuxProfile) {
	api.AdminUsername = obj.AdminUsername
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range obj.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys,
			PublicKey{KeyData: d.KeyData})
	}
}

func convertV20160330LinuxProfile(v20160330 *v20160330.LinuxProfile, api *LinuxProfile) {
	api.AdminUsername = v20160330.AdminUsername
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range v20160330.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys,
			PublicKey{KeyData: d.KeyData})
	}
}

func convertV20170131LinuxProfile(v20170131 *v20170131.LinuxProfile, api *LinuxProfile) {
	api.AdminUsername = v20170131.AdminUsername
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range v20170131.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys, PublicKey{KeyData: d.KeyData})
	}
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

func convertV20170701LinuxProfile(v20170701 *v20170701.LinuxProfile, api *LinuxProfile) {
	api.AdminUsername = v20170701.AdminUsername
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range v20170701.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys,
			PublicKey{KeyData: d.KeyData})
	}
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

func convertV20160930WindowsProfile(v20160930 *v20160930.WindowsProfile, api *WindowsProfile) {
	api.AdminUsername = v20160930.AdminUsername
	api.AdminPassword = v20160930.AdminPassword
}

func convertV20160330WindowsProfile(v20160330 *v20160330.WindowsProfile, api *WindowsProfile) {
	api.AdminUsername = v20160330.AdminUsername
	api.AdminPassword = v20160330.AdminPassword
}

func convertV20170131WindowsProfile(v20170131 *v20170131.WindowsProfile, api *WindowsProfile) {
	api.AdminUsername = v20170131.AdminUsername
	api.AdminPassword = v20170131.AdminPassword
}

func convertV20170701WindowsProfile(v20170701 *v20170701.WindowsProfile, api *WindowsProfile) {
	api.AdminUsername = v20170701.AdminUsername
	api.AdminPassword = v20170701.AdminPassword
}

func convertVLabsWindowsProfile(vlabs *vlabs.WindowsProfile, api *WindowsProfile) {
	api.AdminUsername = vlabs.AdminUsername
	api.AdminPassword = vlabs.AdminPassword
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
}

func convertV20160930OrchestratorProfile(v20160930 *v20160930.OrchestratorProfile, api *OrchestratorProfile) {
	api.OrchestratorType = v20160930.OrchestratorType
	if api.OrchestratorType == Kubernetes {
		api.OrchestratorVersion = "1.6.9"
	} else if api.OrchestratorType == DCOS {
		api.OrchestratorVersion = common.DCOSVersion1Dot9Dot0
	}
}

func convertV20160330OrchestratorProfile(v20160330 *v20160330.OrchestratorProfile, api *OrchestratorProfile) {
	api.OrchestratorType = v20160330.OrchestratorType
	if api.OrchestratorType == DCOS {
		api.OrchestratorVersion = common.DCOSVersion1Dot9Dot0
	}
}

func convertV20170131OrchestratorProfile(v20170131 *v20170131.OrchestratorProfile, api *OrchestratorProfile, hasWindows bool) {
	api.OrchestratorType = v20170131.OrchestratorType
	if api.OrchestratorType == Kubernetes {
		api.OrchestratorVersion = common.GetSupportedKubernetesVersion("", hasWindows)
	} else if api.OrchestratorType == DCOS {
		api.OrchestratorVersion = common.DCOSVersion1Dot9Dot0
	}
}

func convertV20170701OrchestratorProfile(v20170701cs *v20170701.OrchestratorProfile, api *OrchestratorProfile, hasWindows bool) {
	if v20170701cs.OrchestratorType == v20170701.DockerCE {
		api.OrchestratorType = SwarmMode
	} else {
		api.OrchestratorType = v20170701cs.OrchestratorType
	}

	switch api.OrchestratorType {
	case Kubernetes:
		api.OrchestratorVersion = common.GetSupportedKubernetesVersion(v20170701cs.OrchestratorVersion, hasWindows)
	case DCOS:
		switch v20170701cs.OrchestratorVersion {
		case common.DCOSVersion1Dot10Dot0, common.DCOSVersion1Dot9Dot0, common.DCOSVersion1Dot8Dot8:
			api.OrchestratorVersion = v20170701cs.OrchestratorVersion
		default:
			api.OrchestratorVersion = common.DCOSVersion1Dot9Dot0
		}
	default:
		break
	}
}

func convertVLabsOrchestratorProfile(vp *vlabs.Properties, api *OrchestratorProfile, isUpdate bool) {
	vlabscs := vp.OrchestratorProfile
	api.OrchestratorType = vlabscs.OrchestratorType
	switch api.OrchestratorType {
	case Kubernetes:
		if vlabscs.KubernetesConfig != nil {
			api.KubernetesConfig = &KubernetesConfig{}
			convertVLabsKubernetesConfig(vlabscs.KubernetesConfig, api.KubernetesConfig)
		}
		setVlabsKubernetesDefaults(vp, api)
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
	api.ClusterSubnet = vlabs.ClusterSubnet
	api.DNSServiceIP = vlabs.DNSServiceIP
	api.ServiceCIDR = vlabs.ServiceCidr
	api.NetworkPlugin = vlabs.NetworkPlugin
	api.ContainerRuntime = vlabs.ContainerRuntime
	api.MaxPods = vlabs.MaxPods
	api.DockerBridgeSubnet = vlabs.DockerBridgeSubnet
	api.CloudProviderBackoff = vlabs.CloudProviderBackoff
	api.CloudProviderBackoffDuration = vlabs.CloudProviderBackoffDuration
	api.CloudProviderBackoffExponent = vlabs.CloudProviderBackoffExponent
	api.CloudProviderBackoffJitter = vlabs.CloudProviderBackoffJitter
	api.CloudProviderBackoffRetries = vlabs.CloudProviderBackoffRetries
	api.CloudProviderRateLimit = vlabs.CloudProviderRateLimit
	api.CloudProviderRateLimitBucket = vlabs.CloudProviderRateLimitBucket
	api.CloudProviderRateLimitQPS = vlabs.CloudProviderRateLimitQPS
	api.UseManagedIdentity = vlabs.UseManagedIdentity
	api.UserAssignedID = vlabs.UserAssignedID
	api.UserAssignedClientID = vlabs.UserAssignedClientID
	api.CustomHyperkubeImage = vlabs.CustomHyperkubeImage
	api.CustomCcmImage = vlabs.CustomCcmImage
	api.UseCloudControllerManager = vlabs.UseCloudControllerManager
	api.CustomWindowsPackageURL = vlabs.CustomWindowsPackageURL
	api.WindowsNodeBinariesURL = vlabs.WindowsNodeBinariesURL
	api.UseInstanceMetadata = vlabs.UseInstanceMetadata
	api.LoadBalancerSku = vlabs.LoadBalancerSku
	api.ExcludeMasterFromStandardLB = vlabs.ExcludeMasterFromStandardLB
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
	convertAddonsToAPI(vlabs, api)
	convertKubeletConfigToAPI(vlabs, api)
	convertControllerManagerConfigToAPI(vlabs, api)
	convertCloudControllerManagerConfigToAPI(vlabs, api)
	convertAPIServerConfigToAPI(vlabs, api)
	convertSchedulerConfigToAPI(vlabs, api)
	convertPrivateClusterToAPI(vlabs, api)
	convertPodSecurityPolicyConfigToAPI(vlabs, api)
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
	if api.KubernetesConfig.NetworkPlugin == "" {
		if vp.HasWindows() {
			api.KubernetesConfig.NetworkPlugin = vlabs.DefaultNetworkPluginWindows
		} else {
			api.KubernetesConfig.NetworkPlugin = vlabs.DefaultNetworkPlugin
		}
	}
}

func convertAddonsToAPI(v *vlabs.KubernetesConfig, a *KubernetesConfig) {
	a.Addons = []KubernetesAddon{}
	for i := range v.Addons {
		a.Addons = append(a.Addons, KubernetesAddon{
			Name:    v.Addons[i].Name,
			Enabled: v.Addons[i].Enabled,
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

func convertV20160930MasterProfile(v20160930 *v20160930.MasterProfile, api *MasterProfile) {
	api.Count = v20160930.Count
	api.DNSPrefix = v20160930.DNSPrefix
	api.FQDN = v20160930.FQDN
	api.Subnet = v20160930.GetSubnet()
	// Set default VMSize
	api.VMSize = "Standard_D2_v2"
}

func convertV20160330MasterProfile(v20160330 *v20160330.MasterProfile, api *MasterProfile) {
	api.Count = v20160330.Count
	api.DNSPrefix = v20160330.DNSPrefix
	api.FQDN = v20160330.FQDN
	api.Subnet = v20160330.GetSubnet()
	// Set default VMSize
	api.VMSize = "Standard_D2_v2"
}

func convertV20170131MasterProfile(v20170131 *v20170131.MasterProfile, api *MasterProfile) {
	api.Count = v20170131.Count
	api.DNSPrefix = v20170131.DNSPrefix
	api.FQDN = v20170131.FQDN
	api.Subnet = v20170131.GetSubnet()
	// Set default VMSize
	api.VMSize = "Standard_D2_v2"
}

func convertV20170701MasterProfile(v20170701 *v20170701.MasterProfile, api *MasterProfile) {
	api.Count = v20170701.Count
	api.DNSPrefix = v20170701.DNSPrefix
	api.FQDN = v20170701.FQDN
	api.Subnet = v20170701.GetSubnet()
	api.VMSize = v20170701.VMSize
	api.OSDiskSizeGB = v20170701.OSDiskSizeGB
	api.VnetSubnetID = v20170701.VnetSubnetID
	api.FirstConsecutiveStaticIP = v20170701.FirstConsecutiveStaticIP
	api.StorageProfile = v20170701.StorageProfile
	// by default 20170701 will use managed disks as it has encryption at rest
	if len(api.StorageProfile) == 0 {
		api.StorageProfile = ManagedDisks
	}
}

func convertVLabsMasterProfile(vlabs *vlabs.MasterProfile, api *MasterProfile) {
	api.Count = vlabs.Count
	api.DNSPrefix = vlabs.DNSPrefix
	api.SubjectAltNames = vlabs.SubjectAltNames
	api.VMSize = vlabs.VMSize
	api.OSDiskSizeGB = vlabs.OSDiskSizeGB
	api.VnetSubnetID = vlabs.VnetSubnetID
	api.AgentVnetSubnetID = vlabs.AgentVnetSubnetID
	api.FirstConsecutiveStaticIP = vlabs.FirstConsecutiveStaticIP
	api.VnetCidr = vlabs.VnetCidr
	api.Subnet = vlabs.GetSubnet()
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
	}

	api.AvailabilityProfile = vlabs.AvailabilityProfile
	api.AgentSubnet = vlabs.AgentSubnet
	api.AvailabilityZones = vlabs.AvailabilityZones
	api.SinglePlacementGroup = vlabs.SinglePlacementGroup
	convertCustomFilesToAPI(vlabs, api)
}

func convertV20160930AgentPoolProfile(v20160930 *v20160930.AgentPoolProfile, availabilityProfile string, api *AgentPoolProfile) {
	api.Name = v20160930.Name
	api.Count = v20160930.Count
	api.VMSize = v20160930.VMSize
	api.DNSPrefix = v20160930.DNSPrefix
	if api.DNSPrefix != "" {
		// Set default Ports when DNSPrefix specified
		api.Ports = []int{80, 443, 8080}
	}
	api.FQDN = v20160930.FQDN
	api.OSType = OSType(v20160930.OSType)
	api.Subnet = v20160930.GetSubnet()
	api.AvailabilityProfile = availabilityProfile
}

func convertV20160330AgentPoolProfile(v20160330 *v20160330.AgentPoolProfile, api *AgentPoolProfile) {
	api.Name = v20160330.Name
	api.Count = v20160330.Count
	api.VMSize = v20160330.VMSize
	api.DNSPrefix = v20160330.DNSPrefix
	if api.DNSPrefix != "" {
		// Set default Ports when DNSPrefix specified
		api.Ports = []int{80, 443, 8080}
	}
	api.FQDN = v20160330.FQDN
	api.OSType = OSType(v20160330.OSType)
	api.Subnet = v20160330.GetSubnet()
}

func convertV20170131AgentPoolProfile(v20170131 *v20170131.AgentPoolProfile, availabilityProfile string, api *AgentPoolProfile) {
	api.Name = v20170131.Name
	api.Count = v20170131.Count
	api.VMSize = v20170131.VMSize
	api.DNSPrefix = v20170131.DNSPrefix
	if api.DNSPrefix != "" {
		// Set default Ports when DNSPrefix specified
		api.Ports = []int{80, 443, 8080}
	}
	api.FQDN = v20170131.FQDN
	api.OSType = OSType(v20170131.OSType)
	api.Subnet = v20170131.GetSubnet()
	api.AvailabilityProfile = availabilityProfile
}

func convertV20170701AgentPoolProfile(v20170701 *v20170701.AgentPoolProfile, availabilityProfile string, api *AgentPoolProfile) {
	api.Name = v20170701.Name
	api.Count = v20170701.Count
	api.VMSize = v20170701.VMSize
	api.OSDiskSizeGB = v20170701.OSDiskSizeGB
	api.DNSPrefix = v20170701.DNSPrefix
	api.OSType = OSType(v20170701.OSType)
	api.Ports = []int{}
	api.Ports = append(api.Ports, v20170701.Ports...)
	api.StorageProfile = v20170701.StorageProfile
	api.VnetSubnetID = v20170701.VnetSubnetID
	api.Subnet = v20170701.GetSubnet()
	api.FQDN = v20170701.FQDN
	api.AvailabilityProfile = availabilityProfile
}

func convertVLabsAgentPoolProfile(vlabs *vlabs.AgentPoolProfile, api *AgentPoolProfile) {
	api.Name = vlabs.Name
	api.Count = vlabs.Count
	api.VMSize = vlabs.VMSize
	api.OSDiskSizeGB = vlabs.OSDiskSizeGB
	api.DNSPrefix = vlabs.DNSPrefix
	api.OSType = OSType(vlabs.OSType)
	api.Ports = []int{}
	api.Ports = append(api.Ports, vlabs.Ports...)
	api.AvailabilityProfile = vlabs.AvailabilityProfile
	api.ScaleSetPriority = vlabs.ScaleSetPriority
	api.ScaleSetEvictionPolicy = vlabs.ScaleSetEvictionPolicy
	api.StorageProfile = vlabs.StorageProfile
	api.DiskSizesGB = []int{}
	api.DiskSizesGB = append(api.DiskSizesGB, vlabs.DiskSizesGB...)
	api.VnetSubnetID = vlabs.VnetSubnetID
	api.Subnet = vlabs.GetSubnet()
	api.IPAddressCount = vlabs.IPAddressCount
	api.FQDN = vlabs.FQDN
	api.AcceleratedNetworkingEnabled = vlabs.AcceleratedNetworkingEnabled
	api.AcceleratedNetworkingEnabledWindows = vlabs.AcceleratedNetworkingEnabledWindows
	api.AvailabilityZones = vlabs.AvailabilityZones
	api.SinglePlacementGroup = vlabs.SinglePlacementGroup

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
	}
	api.Role = AgentPoolProfileRole(vlabs.Role)
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

func convertV20160930DiagnosticsProfile(v20160930 *v20160930.DiagnosticsProfile, api *DiagnosticsProfile) {
	if v20160930.VMDiagnostics != nil {
		api.VMDiagnostics = &VMDiagnostics{}
		convertV20160930VMDiagnostics(v20160930.VMDiagnostics, api.VMDiagnostics)
	}
}

func convertV20160930VMDiagnostics(v20160930 *v20160930.VMDiagnostics, api *VMDiagnostics) {
	api.Enabled = v20160930.Enabled
	api.StorageURL = v20160930.StorageURL
}

func convertV20160330DiagnosticsProfile(v20160330 *v20160330.DiagnosticsProfile, api *DiagnosticsProfile) {
	if v20160330.VMDiagnostics != nil {
		api.VMDiagnostics = &VMDiagnostics{}
		convertV20160330VMDiagnostics(v20160330.VMDiagnostics, api.VMDiagnostics)
	}
}

func convertV20160330VMDiagnostics(v20160330 *v20160330.VMDiagnostics, api *VMDiagnostics) {
	api.Enabled = v20160330.Enabled
	api.StorageURL = v20160330.StorageURL
}

func convertV20170131DiagnosticsProfile(v20170131 *v20170131.DiagnosticsProfile, api *DiagnosticsProfile) {
	if v20170131.VMDiagnostics != nil {
		api.VMDiagnostics = &VMDiagnostics{}
		convertV20170131VMDiagnostics(v20170131.VMDiagnostics, api.VMDiagnostics)
	}
}

func convertV20170131VMDiagnostics(v20170131 *v20170131.VMDiagnostics, api *VMDiagnostics) {
	api.Enabled = v20170131.Enabled
	api.StorageURL = v20170131.StorageURL
}

func convertV20160930JumpboxProfile(v20160930 *v20160930.JumpboxProfile, api *JumpboxProfile) {
	api.OSType = OSType(v20160930.OSType)
	api.DNSPrefix = v20160930.DNSPrefix
	api.FQDN = v20160930.FQDN
}

func convertV20160330JumpboxProfile(v20160330 *v20160330.JumpboxProfile, api *JumpboxProfile) {
	api.OSType = OSType(v20160330.OSType)
	api.DNSPrefix = v20160330.DNSPrefix
	api.FQDN = v20160330.FQDN
}

func convertV20170131JumpboxProfile(v20170131 *v20170131.JumpboxProfile, api *JumpboxProfile) {
	api.OSType = OSType(v20170131.OSType)
	api.DNSPrefix = v20170131.DNSPrefix
	api.FQDN = v20170131.FQDN
}

func convertV20160930ServicePrincipalProfile(v20160930 *v20160930.ServicePrincipalProfile, api *ServicePrincipalProfile) {
	api.ClientID = v20160930.ClientID
	api.Secret = v20160930.Secret
	api.ObjectID = v20160930.ObjectID
}

func convertV20170131ServicePrincipalProfile(v20170131 *v20170131.ServicePrincipalProfile, api *ServicePrincipalProfile) {
	api.ClientID = v20170131.ClientID
	api.Secret = v20170131.Secret
	api.ObjectID = v20170131.ObjectID
}

func convertV20170701ServicePrincipalProfile(v20170701 *v20170701.ServicePrincipalProfile, api *ServicePrincipalProfile) {
	api.ClientID = v20170701.ClientID
	api.Secret = v20170701.Secret
	api.ObjectID = v20170701.ObjectID
	if v20170701.KeyvaultSecretRef != nil {
		api.KeyvaultSecretRef = &KeyvaultSecretRef{
			VaultID:       v20170701.KeyvaultSecretRef.VaultID,
			SecretName:    v20170701.KeyvaultSecretRef.SecretName,
			SecretVersion: v20170701.KeyvaultSecretRef.SecretVersion,
		}
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

func convertV20160930CustomProfile(v20160930 *v20160930.CustomProfile, api *CustomProfile) {
	api.Orchestrator = v20160930.Orchestrator
}

func convertV20170131CustomProfile(v20170131 *v20170131.CustomProfile, api *CustomProfile) {
	api.Orchestrator = v20170131.Orchestrator
}

func convertV20170701CustomProfile(v20170701 *v20170701.CustomProfile, api *CustomProfile) {
	api.Orchestrator = v20170701.Orchestrator
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
