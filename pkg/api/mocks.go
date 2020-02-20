// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/uuid"

	"github.com/Azure/aks-engine/pkg/api/common"
)

// CreateMockAgentPoolProfile creates a mock AgentPoolResource for testing
func CreateMockAgentPoolProfile(agentPoolName, orchestratorVersion string, provisioningState ProvisioningState, agentCount int) *AgentPoolResource {
	agentPoolResource := AgentPoolResource{}
	agentPoolResource.ID = uuid.Must(uuid.NewRandom()).String()
	agentPoolResource.Location = "westus2"
	agentPoolResource.Name = agentPoolName

	agentPoolResource.Properties = &AgentPoolProfile{}
	// AgentPoolProfile needs to be remain same, so the name is repeated inside.
	agentPoolResource.Properties.Name = agentPoolName
	agentPoolResource.Properties.Count = agentCount
	agentPoolResource.Properties.OrchestratorVersion = orchestratorVersion
	agentPoolResource.Properties.ProvisioningState = provisioningState
	return &agentPoolResource
}

// CreateMockContainerService returns a mock container service for testing purposes
func CreateMockContainerService(containerServiceName, orchestratorVersion string, masterCount, agentCount int, certs bool) *ContainerService {
	cs := ContainerService{}
	cs.ID = uuid.Must(uuid.NewRandom()).String()
	cs.Location = "eastus"
	cs.Name = containerServiceName

	cs.Properties = &Properties{}

	cs.Properties.MasterProfile = &MasterProfile{}
	cs.Properties.MasterProfile.Count = masterCount
	cs.Properties.MasterProfile.DNSPrefix = "testmaster"
	cs.Properties.MasterProfile.VMSize = "Standard_D2_v2"

	cs.Properties.AgentPoolProfiles = []*AgentPoolProfile{}
	agentPool := &AgentPoolProfile{}
	agentPool.Count = agentCount
	agentPool.Name = "agentpool1"
	agentPool.VMSize = "Standard_D2_v2"
	agentPool.OSType = Linux
	agentPool.AvailabilityProfile = "AvailabilitySet"
	agentPool.StorageProfile = "StorageAccount"

	cs.Properties.AgentPoolProfiles = append(cs.Properties.AgentPoolProfiles, agentPool)

	cs.Properties.LinuxProfile = &LinuxProfile{
		AdminUsername: "azureuser",
		SSH: struct {
			PublicKeys []PublicKey `json:"publicKeys"`
		}{},
	}

	cs.Properties.LinuxProfile.AdminUsername = "azureuser"
	cs.Properties.LinuxProfile.SSH.PublicKeys = append(
		cs.Properties.LinuxProfile.SSH.PublicKeys, PublicKey{KeyData: "test"})

	cs.Properties.ServicePrincipalProfile = &ServicePrincipalProfile{}
	cs.Properties.ServicePrincipalProfile.ClientID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
	cs.Properties.ServicePrincipalProfile.Secret = "DEC923E3-1EF1-4745-9516-37906D56DEC4"

	cs.Properties.OrchestratorProfile = &OrchestratorProfile{}
	cs.Properties.OrchestratorProfile.OrchestratorType = Kubernetes
	cs.Properties.OrchestratorProfile.OrchestratorVersion = orchestratorVersion
	cs.Properties.OrchestratorProfile.KubernetesConfig = &KubernetesConfig{
		EnableSecureKubelet:     to.BoolPtr(DefaultSecureKubeletEnabled),
		EnableRbac:              to.BoolPtr(DefaultRBACEnabled),
		EtcdDiskSizeGB:          DefaultEtcdDiskSize,
		ServiceCIDR:             DefaultKubernetesServiceCIDR,
		DockerBridgeSubnet:      DefaultDockerBridgeSubnet,
		DNSServiceIP:            DefaultKubernetesDNSServiceIP,
		GCLowThreshold:          DefaultKubernetesGCLowThreshold,
		GCHighThreshold:         DefaultKubernetesGCHighThreshold,
		MaxPods:                 DefaultKubernetesMaxPodsVNETIntegrated,
		ClusterSubnet:           DefaultKubernetesSubnet,
		ContainerRuntime:        DefaultContainerRuntime,
		NetworkPlugin:           DefaultNetworkPlugin,
		NetworkPolicy:           DefaultNetworkPolicy,
		EtcdVersion:             DefaultEtcdVersion,
		MobyVersion:             DefaultMobyVersion,
		ContainerdVersion:       DefaultContainerdVersion,
		LoadBalancerSku:         DefaultLoadBalancerSku,
		KubeletConfig:           make(map[string]string),
		ControllerManagerConfig: make(map[string]string),
		KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
	}

	cs.Properties.CertificateProfile = &CertificateProfile{}
	if certs {
		cs.Properties.CertificateProfile.CaCertificate = "cacert"
		cs.Properties.CertificateProfile.CaPrivateKey = "cakey"
		cs.Properties.CertificateProfile.KubeConfigCertificate = "kubeconfigcert"
		cs.Properties.CertificateProfile.KubeConfigPrivateKey = "kubeconfigkey"
		cs.Properties.CertificateProfile.APIServerCertificate = "apiservercert"
		cs.Properties.CertificateProfile.APIServerPrivateKey = "apiserverkey"
		cs.Properties.CertificateProfile.ClientCertificate = "clientcert"
		cs.Properties.CertificateProfile.ClientPrivateKey = "clientkey"
		cs.Properties.CertificateProfile.EtcdServerCertificate = "etcdservercert"
		cs.Properties.CertificateProfile.EtcdServerPrivateKey = "etcdserverkey"
		cs.Properties.CertificateProfile.EtcdClientCertificate = "etcdclientcert"
		cs.Properties.CertificateProfile.EtcdClientPrivateKey = "etcdclientkey"
		cs.Properties.CertificateProfile.EtcdPeerCertificates = []string{"etcdpeercert1", "etcdpeercert2", "etcdpeercert3", "etcdpeercert4", "etcdpeercert5"}
		cs.Properties.CertificateProfile.EtcdPeerPrivateKeys = []string{"etcdpeerkey1", "etcdpeerkey2", "etcdpeerkey3", "etcdpeerkey4", "etcdpeerkey5"}

	}

	return &cs
}

// GetK8sDefaultProperties returns a struct of type api.Properties for testing purposes.
func GetK8sDefaultProperties(hasWindows bool) *Properties {
	p := &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
			KubernetesConfig: &KubernetesConfig{},
		},
		MasterProfile: &MasterProfile{
			Count:     1,
			DNSPrefix: "foo",
			VMSize:    "Standard_DS2_v2",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: AvailabilitySet,
			},
		},
		ServicePrincipalProfile: &ServicePrincipalProfile{
			ClientID: "clientID",
			Secret:   "clientSecret",
		},
	}

	if hasWindows {
		p.AgentPoolProfiles = []*AgentPoolProfile{
			{
				Name:                "agentpool",
				VMSize:              "Standard_D2_v2",
				Count:               1,
				AvailabilityProfile: AvailabilitySet,
				OSType:              Windows,
			},
		}
		p.WindowsProfile = &WindowsProfile{
			AdminUsername: "azureuser",
			AdminPassword: "replacepassword1234$",
		}
	}

	return p
}

// GetMockPropertiesWithCustomCloudProfile returns a Properties object w/ mock CustomCloudProfile data
func GetMockPropertiesWithCustomCloudProfile(name string, hasCustomCloudProfile, hasEnvironment, hasAzureEnvironmentSpecConfig bool) Properties {
	var (
		managementPortalURL          = "https://management.local.azurestack.external/"
		publishSettingsURL           = "https://management.local.azurestack.external/publishsettings/index"
		serviceManagementEndpoint    = "https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189"
		resourceManagerEndpoint      = "https://management.local.azurestack.external/"
		activeDirectoryEndpoint      = "https://login.windows.net/"
		galleryEndpoint              = "https://portal.local.azurestack.external=30015/"
		keyVaultEndpoint             = "https://vault.azurestack.external/"
		graphEndpoint                = "https://graph.windows.net/"
		serviceBusEndpoint           = "https://servicebus.azurestack.external/"
		batchManagementEndpoint      = "https://batch.azurestack.external/"
		storageEndpointSuffix        = "core.azurestack.external"
		sqlDatabaseDNSSuffix         = "database.azurestack.external"
		trafficManagerDNSSuffix      = "trafficmanager.cn"
		keyVaultDNSSuffix            = "vault.azurestack.external"
		serviceBusEndpointSuffix     = "servicebus.azurestack.external"
		serviceManagementVMDNSSuffix = "chinacloudapp.cn"
		resourceManagerVMDNSSuffix   = "cloudapp.azurestack.external"
		containerRegistryDNSSuffix   = "azurecr.io"
		tokenAudience                = "https://management.azurestack.external/"
	)

	p := Properties{}
	if hasCustomCloudProfile {
		p.CustomCloudProfile = &CustomCloudProfile{}
		if hasEnvironment {
			p.CustomCloudProfile.Environment = &azure.Environment{
				Name:                         name,
				ManagementPortalURL:          managementPortalURL,
				PublishSettingsURL:           publishSettingsURL,
				ServiceManagementEndpoint:    serviceManagementEndpoint,
				ResourceManagerEndpoint:      resourceManagerEndpoint,
				ActiveDirectoryEndpoint:      activeDirectoryEndpoint,
				GalleryEndpoint:              galleryEndpoint,
				KeyVaultEndpoint:             keyVaultEndpoint,
				GraphEndpoint:                graphEndpoint,
				ServiceBusEndpoint:           serviceBusEndpoint,
				BatchManagementEndpoint:      batchManagementEndpoint,
				StorageEndpointSuffix:        storageEndpointSuffix,
				SQLDatabaseDNSSuffix:         sqlDatabaseDNSSuffix,
				TrafficManagerDNSSuffix:      trafficManagerDNSSuffix,
				KeyVaultDNSSuffix:            keyVaultDNSSuffix,
				ServiceBusEndpointSuffix:     serviceBusEndpointSuffix,
				ServiceManagementVMDNSSuffix: serviceManagementVMDNSSuffix,
				ResourceManagerVMDNSSuffix:   resourceManagerVMDNSSuffix,
				ContainerRegistryDNSSuffix:   containerRegistryDNSSuffix,
				TokenAudience:                tokenAudience,
			}
		}
		if hasAzureEnvironmentSpecConfig {
			//azureStackCloudSpec is the default configurations for azure stack with public Azure.
			azureStackCloudSpec := AzureEnvironmentSpecConfig{
				CloudName: AzureStackCloud,
				//DockerSpecConfig specify the docker engine download repo
				DockerSpecConfig: DefaultDockerSpecConfig,
				//KubernetesSpecConfig is the default kubernetes container image url.
				KubernetesSpecConfig: DefaultKubernetesSpecConfig,
				DCOSSpecConfig:       DefaultDCOSSpecConfig,
				EndpointConfig: AzureEndpointConfig{
					ResourceManagerVMDNSSuffix: "",
				},
				OSImageConfig: map[Distro]AzureOSImageConfig{
					Ubuntu:        Ubuntu1604OSImageConfig,
					RHEL:          RHELOSImageConfig,
					CoreOS:        CoreOSImageConfig,
					AKSUbuntu1604: AKSUbuntu1604OSImageConfig,
				},
			}
			p.CustomCloudProfile.AzureEnvironmentSpecConfig = &azureStackCloudSpec
		}
		p.CustomCloudProfile.IdentitySystem = AzureADIdentitySystem
		p.CustomCloudProfile.AuthenticationMethod = ClientSecretAuthMethod
	}
	return p
}

func getMockAddon(name string) KubernetesAddon {
	return KubernetesAddon{
		Name: name,
		Containers: []KubernetesContainerSpec{
			{
				Name:           name,
				CPURequests:    "50m",
				MemoryRequests: "150Mi",
				CPULimits:      "50m",
				MemoryLimits:   "150Mi",
			},
		},
		Pools: []AddonNodePoolsConfig{
			{
				Name: "pool1",
				Config: map[string]string{
					"min-nodes": "3",
					"max-nodes": "3",
				},
			},
		},
	}
}
