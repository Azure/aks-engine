// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
)

func getParameters(cs *api.ContainerService, generatorCode string, acsengineVersion string) (paramsMap, error) {
	properties := cs.Properties
	location := cs.Location
	parametersMap := paramsMap{}
	cloudSpecConfig := cs.GetCloudSpecConfig()

	// acsengine Parameters
	addValue(parametersMap, "acsengineVersion", acsengineVersion)

	// Master Parameters
	addValue(parametersMap, "location", location)

	// Identify Master distro
	if properties.MasterProfile != nil {
		addValue(parametersMap, "osImageOffer", cloudSpecConfig.OSImageConfig[properties.MasterProfile.Distro].ImageOffer)
		addValue(parametersMap, "osImageSKU", cloudSpecConfig.OSImageConfig[properties.MasterProfile.Distro].ImageSku)
		addValue(parametersMap, "osImagePublisher", cloudSpecConfig.OSImageConfig[properties.MasterProfile.Distro].ImagePublisher)
		addValue(parametersMap, "osImageVersion", cloudSpecConfig.OSImageConfig[properties.MasterProfile.Distro].ImageVersion)
		if properties.MasterProfile.ImageRef != nil {
			addValue(parametersMap, "osImageName", properties.MasterProfile.ImageRef.Name)
			addValue(parametersMap, "osImageResourceGroup", properties.MasterProfile.ImageRef.ResourceGroup)
		}
	}

	addValue(parametersMap, "fqdnEndpointSuffix", cloudSpecConfig.EndpointConfig.ResourceManagerVMDNSSuffix)
	addValue(parametersMap, "targetEnvironment", helpers.GetCloudTargetEnv(cs.Location))
	addValue(parametersMap, "linuxAdminUsername", properties.LinuxProfile.AdminUsername)
	if properties.LinuxProfile.CustomSearchDomain != nil {
		addValue(parametersMap, "searchDomainName", properties.LinuxProfile.CustomSearchDomain.Name)
		addValue(parametersMap, "searchDomainRealmUser", properties.LinuxProfile.CustomSearchDomain.RealmUser)
		addValue(parametersMap, "searchDomainRealmPassword", properties.LinuxProfile.CustomSearchDomain.RealmPassword)
	}
	if properties.LinuxProfile.CustomNodesDNS != nil {
		addValue(parametersMap, "dnsServer", properties.LinuxProfile.CustomNodesDNS.DNSServer)
	}
	// masterEndpointDNSNamePrefix is the basis for storage account creation across dcos, swarm, and k8s
	if properties.MasterProfile != nil {
		// MasterProfile exists, uses master DNS prefix
		addValue(parametersMap, "masterEndpointDNSNamePrefix", properties.MasterProfile.DNSPrefix)
	} else if properties.HostedMasterProfile != nil {
		// Agents only, use cluster DNS prefix
		addValue(parametersMap, "masterEndpointDNSNamePrefix", properties.HostedMasterProfile.DNSPrefix)
	}
	if properties.MasterProfile != nil {
		if properties.MasterProfile.IsCustomVNET() {
			addValue(parametersMap, "masterVnetSubnetID", properties.MasterProfile.VnetSubnetID)
			if properties.MasterProfile.IsVirtualMachineScaleSets() {
				addValue(parametersMap, "agentVnetSubnetID", properties.MasterProfile.AgentVnetSubnetID)
			}
			if properties.OrchestratorProfile.IsKubernetes() {
				addValue(parametersMap, "vnetCidr", properties.MasterProfile.VnetCidr)
			}
		} else {
			addValue(parametersMap, "masterSubnet", properties.MasterProfile.Subnet)
			addValue(parametersMap, "agentSubnet", properties.MasterProfile.AgentSubnet)
		}
		addValue(parametersMap, "firstConsecutiveStaticIP", properties.MasterProfile.FirstConsecutiveStaticIP)
		addValue(parametersMap, "masterVMSize", properties.MasterProfile.VMSize)
		if properties.MasterProfile.HasAvailabilityZones() {
			addValue(parametersMap, "availabilityZones", properties.MasterProfile.AvailabilityZones)
		}
	}
	if properties.HostedMasterProfile != nil {
		addValue(parametersMap, "masterSubnet", properties.HostedMasterProfile.Subnet)
	}
	addValue(parametersMap, "sshRSAPublicKey", properties.LinuxProfile.SSH.PublicKeys[0].KeyData)
	for i, s := range properties.LinuxProfile.Secrets {
		addValue(parametersMap, fmt.Sprintf("linuxKeyVaultID%d", i), s.SourceVault.ID)
		for j, c := range s.VaultCertificates {
			addValue(parametersMap, fmt.Sprintf("linuxKeyVaultID%dCertificateURL%d", i, j), c.CertificateURL)
		}
	}

	//Swarm and SwarmMode Parameters
	if properties.OrchestratorProfile.OrchestratorType == api.Swarm || properties.OrchestratorProfile.OrchestratorType == api.SwarmMode {
		var dockerEngineRepo, dockerComposeDownloadURL string
		if cloudSpecConfig.DockerSpecConfig.DockerEngineRepo == "" {
			dockerEngineRepo = DefaultDockerEngineRepo
		} else {
			dockerEngineRepo = cloudSpecConfig.DockerSpecConfig.DockerEngineRepo
		}
		if cloudSpecConfig.DockerSpecConfig.DockerComposeDownloadURL == "" {
			dockerComposeDownloadURL = DefaultDockerComposeURL
		} else {
			dockerComposeDownloadURL = cloudSpecConfig.DockerSpecConfig.DockerComposeDownloadURL
		}
		addValue(parametersMap, "dockerEngineDownloadRepo", dockerEngineRepo)
		addValue(parametersMap, "dockerComposeDownloadURL", dockerComposeDownloadURL)
	}

	// Kubernetes Parameters
	if properties.OrchestratorProfile.IsKubernetes() {
		assignKubernetesParameters(properties, parametersMap, cloudSpecConfig, generatorCode)
	}

	if strings.HasPrefix(properties.OrchestratorProfile.OrchestratorType, api.DCOS) {
		dcosBootstrapURL := cloudSpecConfig.DCOSSpecConfig.DCOS188BootstrapDownloadURL
		dcosWindowsBootstrapURL := cloudSpecConfig.DCOSSpecConfig.DCOSWindowsBootstrapDownloadURL
		dcosRepositoryURL := cloudSpecConfig.DCOSSpecConfig.DcosRepositoryURL
		dcosClusterPackageListID := cloudSpecConfig.DCOSSpecConfig.DcosClusterPackageListID
		dcosProviderPackageID := cloudSpecConfig.DCOSSpecConfig.DcosProviderPackageID

		switch properties.OrchestratorProfile.OrchestratorType {
		case api.DCOS:
			switch properties.OrchestratorProfile.OrchestratorVersion {
			case common.DCOSVersion1Dot8Dot8:
				dcosBootstrapURL = cloudSpecConfig.DCOSSpecConfig.DCOS188BootstrapDownloadURL
			case common.DCOSVersion1Dot9Dot0:
				dcosBootstrapURL = cloudSpecConfig.DCOSSpecConfig.DCOS190BootstrapDownloadURL
			case common.DCOSVersion1Dot9Dot8:
				dcosBootstrapURL = cloudSpecConfig.DCOSSpecConfig.DCOS198BootstrapDownloadURL
			case common.DCOSVersion1Dot10Dot0:
				dcosBootstrapURL = cloudSpecConfig.DCOSSpecConfig.DCOS110BootstrapDownloadURL
			default:
				dcosBootstrapURL = getDCOSDefaultBootstrapInstallerURL(properties.OrchestratorProfile)
				dcosWindowsBootstrapURL = getDCOSDefaultWindowsBootstrapInstallerURL(properties.OrchestratorProfile)
			}
		}

		if properties.OrchestratorProfile.DcosConfig != nil {
			if properties.OrchestratorProfile.DcosConfig.DcosWindowsBootstrapURL != "" {
				dcosWindowsBootstrapURL = properties.OrchestratorProfile.DcosConfig.DcosWindowsBootstrapURL
			}
			if properties.OrchestratorProfile.DcosConfig.DcosBootstrapURL != "" {
				dcosBootstrapURL = properties.OrchestratorProfile.DcosConfig.DcosBootstrapURL
			}
			if len(properties.OrchestratorProfile.DcosConfig.Registry) > 0 {
				addValue(parametersMap, "registry", properties.OrchestratorProfile.DcosConfig.Registry)
				addValue(parametersMap, "registryKey", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", properties.OrchestratorProfile.DcosConfig.RegistryUser, properties.OrchestratorProfile.DcosConfig.RegistryPass))))
			}
			if properties.OrchestratorProfile.DcosConfig.DcosRepositoryURL != "" {
				dcosRepositoryURL = properties.OrchestratorProfile.DcosConfig.DcosRepositoryURL
			} else {
				dcosRepositoryURL = getDCOSDefaultRepositoryURL(
					properties.OrchestratorProfile.OrchestratorType,
					properties.OrchestratorProfile.OrchestratorVersion)
			}

			if properties.OrchestratorProfile.DcosConfig.DcosClusterPackageListID != "" {
				dcosClusterPackageListID = properties.OrchestratorProfile.DcosConfig.DcosClusterPackageListID
			}

			if properties.OrchestratorProfile.DcosConfig.DcosProviderPackageID != "" {
				dcosProviderPackageID = properties.OrchestratorProfile.DcosConfig.DcosProviderPackageID
			} else {
				dcosProviderPackageID = getDCOSDefaultProviderPackageGUID(
					properties.OrchestratorProfile.OrchestratorType,
					properties.OrchestratorProfile.OrchestratorVersion,
					properties.MasterProfile.Count)
			}
		}

		addValue(parametersMap, "dcosBootstrapURL", dcosBootstrapURL)
		addValue(parametersMap, "dcosWindowsBootstrapURL", dcosWindowsBootstrapURL)
		addValue(parametersMap, "dcosRepositoryURL", dcosRepositoryURL)
		addValue(parametersMap, "dcosClusterPackageListID", dcosClusterPackageListID)
		addValue(parametersMap, "dcosProviderPackageID", dcosProviderPackageID)

		if properties.OrchestratorProfile.DcosConfig != nil && properties.OrchestratorProfile.DcosConfig.BootstrapProfile != nil {
			addValue(parametersMap, "bootstrapStaticIP", properties.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP)
			addValue(parametersMap, "bootstrapVMSize", properties.OrchestratorProfile.DcosConfig.BootstrapProfile.VMSize)
		}
	}

	// Agent parameters
	for _, agentProfile := range properties.AgentPoolProfiles {
		addValue(parametersMap, fmt.Sprintf("%sCount", agentProfile.Name), agentProfile.Count)
		addValue(parametersMap, fmt.Sprintf("%sVMSize", agentProfile.Name), agentProfile.VMSize)
		if agentProfile.HasAvailabilityZones() {
			addValue(parametersMap, fmt.Sprintf("%sAvailabilityZones", agentProfile.Name), agentProfile.AvailabilityZones)
		}
		if agentProfile.IsCustomVNET() {
			addValue(parametersMap, fmt.Sprintf("%sVnetSubnetID", agentProfile.Name), agentProfile.VnetSubnetID)
		} else {
			addValue(parametersMap, fmt.Sprintf("%sSubnet", agentProfile.Name), agentProfile.Subnet)
		}
		if len(agentProfile.Ports) > 0 {
			addValue(parametersMap, fmt.Sprintf("%sEndpointDNSNamePrefix", agentProfile.Name), agentProfile.DNSPrefix)
		}

		// Unless distro is defined, default distro is configured by defaults#setAgentProfileDefaults
		//   Ignores Windows OS
		if !(agentProfile.OSType == api.Windows) {
			if agentProfile.ImageRef != nil {
				addValue(parametersMap, fmt.Sprintf("%sosImageName", agentProfile.Name), agentProfile.ImageRef.Name)
				addValue(parametersMap, fmt.Sprintf("%sosImageResourceGroup", agentProfile.Name), agentProfile.ImageRef.ResourceGroup)
			}
			addValue(parametersMap, fmt.Sprintf("%sosImageOffer", agentProfile.Name), cloudSpecConfig.OSImageConfig[agentProfile.Distro].ImageOffer)
			addValue(parametersMap, fmt.Sprintf("%sosImageSKU", agentProfile.Name), cloudSpecConfig.OSImageConfig[agentProfile.Distro].ImageSku)
			addValue(parametersMap, fmt.Sprintf("%sosImagePublisher", agentProfile.Name), cloudSpecConfig.OSImageConfig[agentProfile.Distro].ImagePublisher)
			addValue(parametersMap, fmt.Sprintf("%sosImageVersion", agentProfile.Name), cloudSpecConfig.OSImageConfig[agentProfile.Distro].ImageVersion)
		}
	}

	// Windows parameters
	if properties.HasWindows() {
		addValue(parametersMap, "windowsAdminUsername", properties.WindowsProfile.AdminUsername)
		addSecret(parametersMap, "windowsAdminPassword", properties.WindowsProfile.AdminPassword, false)
		if properties.WindowsProfile.ImageVersion != "" {
			addValue(parametersMap, "agentWindowsVersion", properties.WindowsProfile.ImageVersion)
		}
		if properties.WindowsProfile.WindowsImageSourceURL != "" {
			addValue(parametersMap, "agentWindowsSourceUrl", properties.WindowsProfile.WindowsImageSourceURL)
		}
		if properties.WindowsProfile.WindowsPublisher != "" {
			addValue(parametersMap, "agentWindowsPublisher", properties.WindowsProfile.WindowsPublisher)
		}
		if properties.WindowsProfile.WindowsOffer != "" {
			addValue(parametersMap, "agentWindowsOffer", properties.WindowsProfile.WindowsOffer)
		}

		addValue(parametersMap, "agentWindowsSku", properties.WindowsProfile.GetWindowsSku())
		addValue(parametersMap, "windowsDockerVersion", properties.WindowsProfile.GetWindowsDockerVersion())

		for i, s := range properties.WindowsProfile.Secrets {
			addValue(parametersMap, fmt.Sprintf("windowsKeyVaultID%d", i), s.SourceVault.ID)
			for j, c := range s.VaultCertificates {
				addValue(parametersMap, fmt.Sprintf("windowsKeyVaultID%dCertificateURL%d", i, j), c.CertificateURL)
				addValue(parametersMap, fmt.Sprintf("windowsKeyVaultID%dCertificateStore%d", i, j), c.CertificateStore)
			}
		}
	}

	for _, extension := range properties.ExtensionProfiles {
		if extension.ExtensionParametersKeyVaultRef != nil {
			addKeyvaultReference(parametersMap, fmt.Sprintf("%sParameters", extension.Name),
				extension.ExtensionParametersKeyVaultRef.VaultID,
				extension.ExtensionParametersKeyVaultRef.SecretName,
				extension.ExtensionParametersKeyVaultRef.SecretVersion)
		} else {
			addValue(parametersMap, fmt.Sprintf("%sParameters", extension.Name), extension.ExtensionParameters)
		}
	}

	return parametersMap, nil
}
