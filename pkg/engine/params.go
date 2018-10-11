package engine

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/helpers"
)

func getParameters(cs *api.ContainerService, generatorCode string, aksengineVersion string) (paramsMap, error) {
	properties := cs.Properties
	location := cs.Location
	parametersMap := paramsMap{}
	cloudSpecConfig := cs.GetCloudSpecConfig()

	// aksengine Parameters
	addValue(parametersMap, "aksengineVersion", aksengineVersion)

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
	// masterEndpointDNSNamePrefix is the basis for storage account creation in k8s
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

	// Kubernetes Parameters
	if properties.OrchestratorProfile.IsKubernetes() {
		assignKubernetesParameters(properties, parametersMap, cloudSpecConfig, generatorCode)
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
		if properties.WindowsProfile.WindowsSku != "" {
			addValue(parametersMap, "agentWindowsSku", properties.WindowsProfile.WindowsSku)
		}

		addValue(parametersMap, "windowsDockerVersion", properties.WindowsProfile.GetWindowsDockerVersion())

		if properties.OrchestratorProfile.IsKubernetes() {
			k8sVersion := properties.OrchestratorProfile.OrchestratorVersion

			if properties.OrchestratorProfile.KubernetesConfig != nil {

				// Kubernetes packages as zip file as created by scripts/build-windows-k8s.sh
				// will be removed in future release as if gets phased out (https://github.com/Azure/aks-engine/issues/3851)
				kubeBinariesSASURL := properties.OrchestratorProfile.KubernetesConfig.CustomWindowsPackageURL
				if kubeBinariesSASURL == "" {
					kubeBinariesSASURL = cloudSpecConfig.KubernetesSpecConfig.KubeBinariesSASURLBase + api.K8sComponentsByVersionMap[k8sVersion]["windowszip"]
				}
				addValue(parametersMap, "kubeBinariesSASURL", kubeBinariesSASURL)

				// Kubernetes node binaries as packaged by upstream kubernetes
				// example at https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.11.md#node-binaries-1
				addValue(parametersMap, "windowsKubeBinariesURL", properties.OrchestratorProfile.KubernetesConfig.WindowsNodeBinariesURL)

				addValue(parametersMap, "kubeBinariesVersion", k8sVersion)
				addValue(parametersMap, "windowsTelemetryGUID", cloudSpecConfig.KubernetesSpecConfig.WindowsTelemetryGUID)
			}
		}
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
