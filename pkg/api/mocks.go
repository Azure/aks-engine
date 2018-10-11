package api

import (
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/satori/go.uuid"
)

// CreateMockContainerService returns a mock container service for testing purposes
func CreateMockContainerService(containerServiceName, orchestratorVersion string, masterCount, agentCount int, certs bool) *ContainerService {
	cs := ContainerService{}
	cs.ID = uuid.NewV4().String()
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
	agentPool.OSType = "Linux"
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
		EnableSecureKubelet: helpers.PointerToBool(DefaultSecureKubeletEnabled),
		EnableRbac:          helpers.PointerToBool(DefaultRBACEnabled),
		EtcdDiskSizeGB:      DefaultEtcdDiskSize,
		ServiceCIDR:         DefaultKubernetesServiceCIDR,
		DockerBridgeSubnet:  DefaultDockerBridgeSubnet,
		DNSServiceIP:        DefaultKubernetesDNSServiceIP,
		GCLowThreshold:      DefaultKubernetesGCLowThreshold,
		GCHighThreshold:     DefaultKubernetesGCHighThreshold,
		MaxPods:             DefaultKubernetesMaxPodsVNETIntegrated,
		ClusterSubnet:       DefaultKubernetesSubnet,
		ContainerRuntime:    DefaultContainerRuntime,
		NetworkPlugin:       DefaultNetworkPlugin,
		NetworkPolicy:       DefaultNetworkPolicy,
		EtcdVersion:         DefaultEtcdVersion,
		KubeletConfig:       make(map[string]string),
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
