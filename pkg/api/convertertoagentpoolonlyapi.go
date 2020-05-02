// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"encoding/json"
	"strconv"

	"github.com/Azure/go-autorest/autorest/to"

	v20170831 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20170831"
	v20180331 "github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/v20180331"
	"github.com/Azure/aks-engine/pkg/api/agentPoolOnlyApi/vlabs"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
)

///////////////////////////////////////////////////////////
// The converter exposes functions to convert the top level
// ContainerService resource
//
// All other functions are internal helper functions used
// for converting.
///////////////////////////////////////////////////////////

// ConvertV20170831AgentPoolOnly converts an AgentPoolOnly object into an in-memory container service
func ConvertV20170831AgentPoolOnly(v20170831 *v20170831.ManagedCluster) *ContainerService {
	c := &ContainerService{}
	c.ID = v20170831.ID
	c.Location = helpers.NormalizeAzureRegion(v20170831.Location)
	c.Name = v20170831.Name
	if v20170831.Plan != nil {
		c.Plan = convertv20170831AgentPoolOnlyResourcePurchasePlan(v20170831.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range v20170831.Tags {
		c.Tags[k] = v
	}
	c.Type = v20170831.Type
	c.Properties = convertV20170831AgentPoolOnlyProperties(v20170831.Properties)
	return c
}

// ConvertV20180331AgentPoolOnly converts an AgentPoolOnly object into an in-memory container service
func ConvertV20180331AgentPoolOnly(v20180331 *v20180331.ManagedCluster) *ContainerService {
	c := &ContainerService{}
	c.ID = v20180331.ID
	c.Location = helpers.NormalizeAzureRegion(v20180331.Location)
	c.Name = v20180331.Name
	if v20180331.Plan != nil {
		c.Plan = convertv20180331AgentPoolOnlyResourcePurchasePlan(v20180331.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range v20180331.Tags {
		c.Tags[k] = v
	}
	c.Type = v20180331.Type
	c.Properties = convertV20180331AgentPoolOnlyProperties(v20180331.Properties)
	return c
}

func convertv20170831AgentPoolOnlyResourcePurchasePlan(v20170831 *v20170831.ResourcePurchasePlan) *ResourcePurchasePlan {
	return &ResourcePurchasePlan{
		Name:          v20170831.Name,
		Product:       v20170831.Product,
		PromotionCode: v20170831.PromotionCode,
		Publisher:     v20170831.Publisher,
	}
}

func convertV20170831AgentPoolOnlyProperties(obj *v20170831.Properties) *Properties {
	properties := &Properties{
		ProvisioningState: ProvisioningState(obj.ProvisioningState),
		MasterProfile:     nil,
	}

	properties.HostedMasterProfile = &HostedMasterProfile{}
	properties.HostedMasterProfile.DNSPrefix = obj.DNSPrefix
	properties.HostedMasterProfile.FQDN = obj.FQDN

	properties.OrchestratorProfile = convertV20170831AgentPoolOnlyOrchestratorProfile(obj.KubernetesVersion)

	properties.AgentPoolProfiles = make([]*AgentPoolProfile, len(obj.AgentPoolProfiles))
	for i := range obj.AgentPoolProfiles {
		properties.AgentPoolProfiles[i] = convertV20170831AgentPoolOnlyAgentPoolProfile(obj.AgentPoolProfiles[i], AvailabilitySet)
	}
	if obj.LinuxProfile != nil {
		properties.LinuxProfile = convertV20170831AgentPoolOnlyLinuxProfile(obj.LinuxProfile)
	}
	if obj.WindowsProfile != nil {
		properties.WindowsProfile = convertV20170831AgentPoolOnlyWindowsProfile(obj.WindowsProfile)
	}

	if obj.ServicePrincipalProfile != nil {
		properties.ServicePrincipalProfile = convertV20170831AgentPoolOnlyServicePrincipalProfile(obj.ServicePrincipalProfile)
	}

	return properties
}

// ConvertVLabsAgentPoolOnly converts a vlabs ContainerService to an unversioned ContainerService
func ConvertVLabsAgentPoolOnly(vlabs *vlabs.ManagedCluster) *ContainerService {
	c := &ContainerService{}
	c.ID = vlabs.ID
	c.Location = helpers.NormalizeAzureRegion(vlabs.Location)
	c.Name = vlabs.Name
	if vlabs.Plan != nil {
		c.Plan = &ResourcePurchasePlan{}
		convertVLabsAgentPoolOnlyResourcePurchasePlan(vlabs.Plan, c.Plan)
	}
	c.Tags = map[string]string{}
	for k, v := range vlabs.Tags {
		c.Tags[k] = v
	}
	c.Type = vlabs.Type
	c.Properties = &Properties{}
	convertVLabsAgentPoolOnlyProperties(vlabs.Properties, c.Properties)
	return c
}

// convertVLabsResourcePurchasePlan converts a vlabs ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertVLabsAgentPoolOnlyResourcePurchasePlan(vlabs *vlabs.ResourcePurchasePlan, api *ResourcePurchasePlan) {
	api.Name = vlabs.Name
	api.Product = vlabs.Product
	api.PromotionCode = vlabs.PromotionCode
	api.Publisher = vlabs.Publisher
}

func convertVLabsAgentPoolOnlyProperties(vlabs *vlabs.Properties, api *Properties) {
	api.ProvisioningState = ProvisioningState(vlabs.ProvisioningState)
	api.OrchestratorProfile = convertVLabsAgentPoolOnlyOrchestratorProfile(vlabs.KubernetesVersion)
	api.MasterProfile = nil

	api.HostedMasterProfile = &HostedMasterProfile{}
	api.HostedMasterProfile.DNSPrefix = vlabs.DNSPrefix
	api.HostedMasterProfile.FQDN = vlabs.FQDN

	api.AgentPoolProfiles = []*AgentPoolProfile{}
	for _, p := range vlabs.AgentPoolProfiles {
		apiProfile := &AgentPoolProfile{}
		convertVLabsAgentPoolOnlyAgentPoolProfile(p, apiProfile)
		// by default vlabs will use managed disks for all orchestrators but kubernetes as it has encryption at rest.
		if !api.OrchestratorProfile.IsKubernetes() {
			// by default vlabs will use managed disks for all orchestrators but kubernetes as it has encryption at rest.
			if len(p.StorageProfile) == 0 {
				apiProfile.StorageProfile = ManagedDisks
			}
		}
		api.AgentPoolProfiles = append(api.AgentPoolProfiles, apiProfile)
	}
	if vlabs.LinuxProfile != nil {
		api.LinuxProfile = &LinuxProfile{}
		convertVLabsAgentPoolOnlyLinuxProfile(vlabs.LinuxProfile, api.LinuxProfile)
	}
	if vlabs.WindowsProfile != nil {
		api.WindowsProfile = &WindowsProfile{}
		convertVLabsAgentPoolOnlyWindowsProfile(vlabs.WindowsProfile, api.WindowsProfile)
	}
	if vlabs.ServicePrincipalProfile != nil {
		api.ServicePrincipalProfile = &ServicePrincipalProfile{}
		convertVLabsAgentPoolOnlyServicePrincipalProfile(vlabs.ServicePrincipalProfile, api.ServicePrincipalProfile)
	}
	if vlabs.CertificateProfile != nil {
		api.CertificateProfile = &CertificateProfile{}
		convertVLabsAgentPoolOnlyCertificateProfile(vlabs.CertificateProfile, api.CertificateProfile)
	}
}

func convertVLabsAgentPoolOnlyLinuxProfile(vlabs *vlabs.LinuxProfile, api *LinuxProfile) {
	api.AdminUsername = vlabs.AdminUsername
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range vlabs.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys,
			PublicKey{KeyData: d.KeyData})
	}
	// api.Secrets = []KeyVaultSecrets{}
	// for _, s := range vlabs.Secrets {
	// 	secret := &KeyVaultSecrets{}
	// 	convertVLabsKeyVaultSecrets(&s, secret)
	// 	api.Secrets = append(api.Secrets, *secret)
	// }
}

func convertV20170831AgentPoolOnlyLinuxProfile(obj *v20170831.LinuxProfile) *LinuxProfile {
	api := &LinuxProfile{
		AdminUsername: obj.AdminUsername,
	}
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range obj.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys, PublicKey{KeyData: d.KeyData})
	}
	return api
}

func convertV20170831AgentPoolOnlyWindowsProfile(obj *v20170831.WindowsProfile) *WindowsProfile {
	return &WindowsProfile{
		AdminUsername: obj.AdminUsername,
		AdminPassword: obj.AdminPassword,
	}
}

func convertVLabsAgentPoolOnlyWindowsProfile(vlabs *vlabs.WindowsProfile, api *WindowsProfile) {
	api.AdminUsername = vlabs.AdminUsername
	api.AdminPassword = vlabs.AdminPassword
	api.ImageVersion = vlabs.ImageVersion
	// api.Secrets = []KeyVaultSecrets{}
	// for _, s := range vlabs.Secrets {
	// 	secret := &KeyVaultSecrets{}
	// 	convertVLabsKeyVaultSecrets(&s, secret)
	// 	api.Secrets = append(api.Secrets, *secret)
	// }
}

func convertV20170831AgentPoolOnlyOrchestratorProfile(kubernetesVersion string) *OrchestratorProfile {
	return &OrchestratorProfile{
		OrchestratorType:    Kubernetes,
		OrchestratorVersion: common.GetSupportedKubernetesVersion(kubernetesVersion, false),
		KubernetesConfig: &KubernetesConfig{
			EnableRbac:          to.BoolPtr(false),
			EnableSecureKubelet: to.BoolPtr(false),
			// set network default for un-versioned model
			NetworkPlugin:      string(v20180331.Kubenet),
			ClusterSubnet:      DefaultKubernetesClusterSubnet,
			ServiceCIDR:        DefaultKubernetesServiceCIDR,
			DNSServiceIP:       DefaultKubernetesDNSServiceIP,
			DockerBridgeSubnet: DefaultDockerBridgeSubnet,
		},
	}
}

func convertVLabsAgentPoolOnlyOrchestratorProfile(kubernetesVersion string) *OrchestratorProfile {
	return &OrchestratorProfile{
		OrchestratorType:    Kubernetes,
		OrchestratorVersion: common.GetSupportedKubernetesVersion(kubernetesVersion, false),
	}
}

func convertV20170831AgentPoolOnlyAgentPoolProfile(v20170831 *v20170831.AgentPoolProfile, availabilityProfile string) *AgentPoolProfile {
	api := &AgentPoolProfile{}
	api.Name = v20170831.Name
	api.Count = v20170831.Count
	api.VMSize = v20170831.VMSize
	api.OSDiskSizeGB = v20170831.OSDiskSizeGB
	api.OSType = OSType(v20170831.OSType)
	api.StorageProfile = v20170831.StorageProfile
	api.VnetSubnetID = v20170831.VnetSubnetID
	api.Subnet = v20170831.GetSubnet()
	api.AvailabilityProfile = availabilityProfile
	return api
}

func convertVLabsAgentPoolOnlyAgentPoolProfile(vlabs *vlabs.AgentPoolProfile, api *AgentPoolProfile) {
	api.Name = vlabs.Name
	api.Count = vlabs.Count
	api.VMSize = vlabs.VMSize
	api.OSDiskSizeGB = vlabs.OSDiskSizeGB
	api.OSType = OSType(vlabs.OSType)
	api.StorageProfile = vlabs.StorageProfile
	api.AvailabilityProfile = vlabs.AvailabilityProfile
	api.VnetSubnetID = vlabs.VnetSubnetID
	api.Subnet = vlabs.GetSubnet()
}

func convertVLabsAgentPoolOnlyServicePrincipalProfile(vlabs *vlabs.ServicePrincipalProfile, api *ServicePrincipalProfile) {
	api.ClientID = vlabs.ClientID
	api.Secret = vlabs.Secret
	// api.KeyvaultSecretRef = vlabs.KeyvaultSecretRef
}

func convertV20170831AgentPoolOnlyServicePrincipalProfile(obj *v20170831.ServicePrincipalProfile) *ServicePrincipalProfile {
	return &ServicePrincipalProfile{
		ClientID: obj.ClientID,
		Secret:   obj.Secret,
	}
}

func convertVLabsAgentPoolOnlyCertificateProfile(vlabs *vlabs.CertificateProfile, api *CertificateProfile) {
	api.CaCertificate = vlabs.CaCertificate
	api.CaPrivateKey = vlabs.CaPrivateKey
	api.APIServerCertificate = vlabs.APIServerCertificate
	api.APIServerPrivateKey = vlabs.APIServerPrivateKey
	api.ClientCertificate = vlabs.ClientCertificate
	api.ClientPrivateKey = vlabs.ClientPrivateKey
	api.KubeConfigCertificate = vlabs.KubeConfigCertificate
	api.KubeConfigPrivateKey = vlabs.KubeConfigPrivateKey
}

func isAgentPoolOnlyClusterJSON(contents []byte) bool {
	properties, propertiesPresent := propertiesAsMap(contents)
	if !propertiesPresent {
		return false
	}
	_, masterProfilePresent := properties["masterProfile"]
	return !masterProfilePresent
}

func propertiesAsMap(contents []byte) (map[string]interface{}, bool) {
	var raw interface{}
	if err := json.Unmarshal(contents, &raw); err != nil {
		return nil, false
	}
	jsonMap := raw.(map[string]interface{})
	properties, propertiesPresent := jsonMap["properties"]
	if !propertiesPresent {
		return nil, false
	}
	return properties.(map[string]interface{}), true
}

func convertv20180331AgentPoolOnlyResourcePurchasePlan(v20180331 *v20180331.ResourcePurchasePlan) *ResourcePurchasePlan {
	return &ResourcePurchasePlan{
		Name:          v20180331.Name,
		Product:       v20180331.Product,
		PromotionCode: v20180331.PromotionCode,
		Publisher:     v20180331.Publisher,
	}
}

func convertV20180331AgentPoolOnlyProperties(obj *v20180331.Properties) *Properties {
	properties := &Properties{
		ProvisioningState: ProvisioningState(obj.ProvisioningState),
		MasterProfile:     nil,
	}

	properties.HostedMasterProfile = &HostedMasterProfile{}
	properties.HostedMasterProfile.DNSPrefix = obj.DNSPrefix
	properties.HostedMasterProfile.FQDN = obj.FQDN

	kubernetesConfig := convertV20180331AgentPoolOnlyKubernetesConfig(obj.EnableRBAC)
	properties.OrchestratorProfile = convertV20180331AgentPoolOnlyOrchestratorProfile(obj.KubernetesVersion, obj.NetworkProfile, kubernetesConfig)

	properties.AgentPoolProfiles = make([]*AgentPoolProfile, len(obj.AgentPoolProfiles))
	for i := range obj.AgentPoolProfiles {
		properties.AgentPoolProfiles[i] = convertV20180331AgentPoolOnlyAgentPoolProfile(obj.AgentPoolProfiles[i], AvailabilitySet, obj.NetworkProfile)
	}
	if obj.LinuxProfile != nil {
		properties.LinuxProfile = convertV20180331AgentPoolOnlyLinuxProfile(obj.LinuxProfile)
	}

	if obj.WindowsProfile != nil {
		properties.WindowsProfile = convertV20180331AgentPoolOnlyWindowsProfile(obj.WindowsProfile)
	}

	if obj.ServicePrincipalProfile != nil {
		properties.ServicePrincipalProfile = convertV20180331AgentPoolOnlyServicePrincipalProfile(obj.ServicePrincipalProfile)
	}

	if obj.AddonProfiles != nil {
		properties.AddonProfiles = convertV20180331AgentPoolOnlyAddonProfiles(obj.AddonProfiles)
	}

	if obj.AADProfile != nil {
		properties.AADProfile = convertV20180331AgentPoolOnlyAADProfile(obj.AADProfile)
	}

	return properties
}

func convertV20180331AgentPoolOnlyLinuxProfile(obj *v20180331.LinuxProfile) *LinuxProfile {
	api := &LinuxProfile{
		AdminUsername: obj.AdminUsername,
	}
	api.SSH.PublicKeys = []PublicKey{}
	for _, d := range obj.SSH.PublicKeys {
		api.SSH.PublicKeys = append(api.SSH.PublicKeys, PublicKey{KeyData: d.KeyData})
	}
	return api
}

func convertV20180331AgentPoolOnlyWindowsProfile(obj *v20180331.WindowsProfile) *WindowsProfile {
	return &WindowsProfile{
		AdminUsername: obj.AdminUsername,
		AdminPassword: obj.AdminPassword,
	}
}

func convertV20180331AgentPoolOnlyKubernetesConfig(enableRBAC *bool) *KubernetesConfig {
	if to.Bool(enableRBAC) {
		// We set default behavior to be false
		return &KubernetesConfig{
			EnableRbac:          to.BoolPtr(true),
			EnableSecureKubelet: to.BoolPtr(true),
		}
	}
	return &KubernetesConfig{
		EnableRbac:          to.BoolPtr(false),
		EnableSecureKubelet: to.BoolPtr(false),
	}
}

func convertV20180331AgentPoolOnlyOrchestratorProfile(kubernetesVersion string, networkProfile *v20180331.NetworkProfile, kubernetesConfig *KubernetesConfig) *OrchestratorProfile {
	if kubernetesConfig == nil {
		kubernetesConfig = &KubernetesConfig{}
	}

	if networkProfile != nil {
		switch networkProfile.NetworkPlugin {
		case v20180331.Azure:
			kubernetesConfig.NetworkPlugin = NetworkPluginAzure

			if networkProfile.NetworkPolicy != "" {
				kubernetesConfig.NetworkPolicy = string(networkProfile.NetworkPolicy)
			}

			if networkProfile.ServiceCidr != "" {
				kubernetesConfig.ServiceCIDR = networkProfile.ServiceCidr
			} else {
				kubernetesConfig.ServiceCIDR = DefaultKubernetesServiceCIDR
			}

			if networkProfile.DNSServiceIP != "" {
				kubernetesConfig.DNSServiceIP = networkProfile.DNSServiceIP
			} else {
				kubernetesConfig.DNSServiceIP = DefaultKubernetesDNSServiceIP
			}

			if networkProfile.DockerBridgeCidr != "" {
				kubernetesConfig.DockerBridgeSubnet = networkProfile.DockerBridgeCidr
			} else {
				kubernetesConfig.DockerBridgeSubnet = DefaultDockerBridgeSubnet
			}
		case v20180331.Kubenet:
			kubernetesConfig.NetworkPlugin = NetworkPluginKubenet

			if networkProfile.NetworkPolicy != "" {
				kubernetesConfig.NetworkPolicy = string(networkProfile.NetworkPolicy)
			}

			if networkProfile.PodCidr != "" {
				kubernetesConfig.ClusterSubnet = networkProfile.PodCidr
			} else {
				kubernetesConfig.ClusterSubnet = DefaultKubernetesClusterSubnet
			}

			if networkProfile.ServiceCidr != "" {
				kubernetesConfig.ServiceCIDR = networkProfile.ServiceCidr
			} else {
				kubernetesConfig.ServiceCIDR = DefaultKubernetesServiceCIDR
			}

			if networkProfile.DNSServiceIP != "" {
				kubernetesConfig.DNSServiceIP = networkProfile.DNSServiceIP
			} else {
				kubernetesConfig.DNSServiceIP = DefaultKubernetesDNSServiceIP
			}

			if networkProfile.DockerBridgeCidr != "" {
				kubernetesConfig.DockerBridgeSubnet = networkProfile.DockerBridgeCidr
			} else {
				kubernetesConfig.DockerBridgeSubnet = DefaultDockerBridgeSubnet
			}
		}
	} else {
		// set network default for un-versioned model
		kubernetesConfig.NetworkPlugin = string(v20180331.Kubenet)
		kubernetesConfig.ClusterSubnet = DefaultKubernetesClusterSubnet
		kubernetesConfig.ServiceCIDR = DefaultKubernetesServiceCIDR
		kubernetesConfig.DNSServiceIP = DefaultKubernetesDNSServiceIP
		kubernetesConfig.DockerBridgeSubnet = DefaultDockerBridgeSubnet
	}

	return &OrchestratorProfile{
		OrchestratorType:    Kubernetes,
		OrchestratorVersion: common.GetSupportedKubernetesVersion(kubernetesVersion, false),
		KubernetesConfig:    kubernetesConfig,
	}
}

func convertV20180331AgentPoolOnlyAgentPoolProfile(agentPoolProfile *v20180331.AgentPoolProfile, availabilityProfile string, networkProfile *v20180331.NetworkProfile) *AgentPoolProfile {
	api := &AgentPoolProfile{}
	api.Name = agentPoolProfile.Name
	api.Count = agentPoolProfile.Count
	api.VMSize = agentPoolProfile.VMSize
	api.OSDiskSizeGB = agentPoolProfile.OSDiskSizeGB
	api.OSType = OSType(agentPoolProfile.OSType)
	api.StorageProfile = agentPoolProfile.StorageProfile
	api.VnetSubnetID = agentPoolProfile.VnetSubnetID
	var maxPods string
	// agentPoolProfile.MaxPods is 0 if maxPods field is not provided in API model
	if agentPoolProfile.MaxPods == nil {
		if networkProfile != nil && networkProfile.NetworkPlugin == v20180331.Azure {
			maxPods = DefaultKubernetesMaxPodsAzureCNI
		} else {
			maxPods = DefaultKubernetesMaxPodsKubenet
		}
	} else {
		maxPods = strconv.Itoa(*agentPoolProfile.MaxPods)
	}
	kubernetesConfig := &KubernetesConfig{
		KubeletConfig: map[string]string{"--max-pods": maxPods},
	}
	api.KubernetesConfig = kubernetesConfig
	api.Subnet = agentPoolProfile.GetSubnet()
	api.AvailabilityProfile = availabilityProfile
	return api
}

func convertV20180331AgentPoolOnlyServicePrincipalProfile(obj *v20180331.ServicePrincipalProfile) *ServicePrincipalProfile {
	return &ServicePrincipalProfile{
		ClientID: obj.ClientID,
		Secret:   obj.Secret,
	}
}

func convertV20180331AgentPoolOnlyAddonProfiles(obj map[string]v20180331.AddonProfile) map[string]AddonProfile {
	api := make(map[string]AddonProfile)
	for k, v := range obj {
		api[k] = AddonProfile{
			Enabled: v.Enabled,
			Config:  v.Config,
		}
	}
	return api
}

func convertV20180331AgentPoolOnlyAADProfile(obj *v20180331.AADProfile) *AADProfile {
	return &AADProfile{
		ClientAppID:     obj.ClientAppID,
		ServerAppID:     obj.ServerAppID,
		ServerAppSecret: obj.ServerAppSecret,
		TenantID:        obj.TenantID,
		Authenticator:   Webhook,
	}
}
