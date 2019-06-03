// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

const (
	// APIVersion is the version of this API
	APIVersion = "vlabs"
)

// the orchestrators supported by vlabs
const (
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS string = "DCOS"
	// Swarm is the string constant for the Swarm orchestrator type
	Swarm string = "Swarm"
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
	// SwarmMode is the string constant for the Swarm Mode orchestrator type
	SwarmMode string = "SwarmMode"
)

// the OSTypes supported by vlabs
const (
	Windows OSType = "Windows"
	Linux   OSType = "Linux"
)

// the LinuxDistros supported by vlabs
const (
	Ubuntu            Distro = "ubuntu"
	Ubuntu1804        Distro = "ubuntu-18.04"
	RHEL              Distro = "rhel"
	CoreOS            Distro = "coreos"
	AKS1604Deprecated Distro = "aks"               // deprecated AKS 16.04 distro. Equivalent to aks-ubuntu-16.04.
	AKS1804Deprecated Distro = "aks-1804"          // deprecated AKS 18.04 distro. Equivalent to aks-ubuntu-18.04.
	AKSDockerEngine   Distro = "aks-docker-engine" // deprecated docker-engine distro.
	AKSUbuntu1604     Distro = "aks-ubuntu-16.04"
	AKSUbuntu1804     Distro = "aks-ubuntu-18.04"
	ACC1604           Distro = "acc-16.04"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents per agent pool
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents per agent pool
	MaxAgentCount = 100
	// MinPort specifies the minimum tcp port to open
	MinPort = 1
	// MaxPort specifies the maximum tcp port to open
	MaxPort = 65535
	// MaxDisks specifies the maximum attached disks to add to the cluster
	MaxDisks = 4
	// MinDiskSizeGB specifies the minimum attached disk size
	MinDiskSizeGB = 1
	// MaxDiskSizeGB specifies the maximum attached disk size
	MaxDiskSizeGB = 1023
	// MinIPAddressCount specifies the minimum number of IP addresses per network interface
	MinIPAddressCount = 1
	// MaxIPAddressCount specifies the maximum number of IP addresses per network interface
	MaxIPAddressCount = 256
)

// Availability profiles
const (
	// AvailabilitySet means that the vms are in an availability set
	AvailabilitySet = "AvailabilitySet"
	// VirtualMachineScaleSets means that the vms are in a virtual machine scaleset
	VirtualMachineScaleSets = "VirtualMachineScaleSets"
)

// storage profiles
const (
	// StorageAccount means that the nodes use raw storage accounts for their os and attached volumes
	StorageAccount = "StorageAccount"
	// ManagedDisks means that the nodes use managed disks for their os and attached volumes
	ManagedDisks = "ManagedDisks"
)

// Supported container runtimes
const (
	Docker          = "docker"
	ClearContainers = "clear-containers"
	KataContainers  = "kata-containers"
	Containerd      = "containerd"
)

var (
	// NetworkPluginValues holds the valid values for network plugin implementation
	NetworkPluginValues = [...]string{"", "kubenet", "azure", NetworkPluginCilium, "flannel"}

	// NetworkPolicyValues holds the valid values for a network policy
	// "azure" and "none" are there for backwards-compatibility
	NetworkPolicyValues = [...]string{"", "calico", NetworkPolicyCilium, "azure", "none"}

	// ContainerRuntimeValues holds the valid values for container runtimes
	ContainerRuntimeValues = [...]string{"", Docker, ClearContainers, KataContainers, Containerd}

	// DistroValues holds the valid values for OS distros
	DistroValues = []Distro{"", Ubuntu, Ubuntu1804, RHEL, CoreOS, AKSUbuntu1604, AKSUbuntu1804, ACC1604}

	// DependenciesLocationValues holds the valid values for dependencies location
	DependenciesLocationValues = []DependenciesLocation{"", AzureStackDependenciesLocationPublic, AzureStackDependenciesLocationChina, AzureStackDependenciesLocationGerman, AzureStackDependenciesLocationUSGovernment}

	// HeapsterAddonContainerValues holds the valid values for heapster addon container names
	HeapsterAddonContainerValues = []string{"heapster", "heapster-nanny"}

	// TillerAddonContainerValues holds the valid values for tiller addon container names
	TillerAddonContainerValues = []string{"tiller"}

	// TillerAddonConfigValues holds the valid values for tiller addon configuration overrides
	TillerAddonConfigValues = []string{"max-history"}

	// ACIConnectorAddonContainerValues holds the valid values for aci-connector addon container names
	ACIConnectorAddonContainerValues = []string{"aci-connector"}

	// ACIConnectorAddonConfigValues holds the valid values for aci-connector addon configuration overrides
	ACIConnectorAddonConfigValues = []string{"region", "nodeName", "os", "taint"}

	// DashboardAddonContainerValues holds the valid values for dashboard addon container names
	DashboardAddonContainerValues = []string{"kubernetes-dashboard"}

	// ReschedulerAddonContainerValues holds the valid values for rescheduler addon container names
	ReschedulerAddonContainerValues = []string{"rescheduler"}

	// ContainerMonitoringAddonContainerValues holds the valid values for container monitoring addon container names
	ContainerMonitoringAddonContainerValues = []string{"omsagent"}

	// ContainerMonitoringAddonConfigValues holds the valid values for container monitoring addon configuration overrides
	ContainerMonitoringAddonConfigValues = []string{"omsAgentVersion", "dockerProviderVersion", "workspaceGuid", "workspaceKey"}

	// IPMasqAgentAddonContainerValues holds the valid values for ip-masq-agent addon container names
	IPMasqAgentAddonContainerValues = []string{"ip-masq-agent"}

	// IPMasqAgentAddonConfigValues holds the valid values for ip-masq-agent addon configuration overrides
	IPMasqAgentAddonConfigValues = []string{"non-masquerade-cidr", "non-masq-cni-cidr"}

	// AADPodIdentityAddonContainerValues holds the valid values for aad-pod-identity addon container names
	AADPodIdentityAddonContainerValues = []string{"nmi", "mic"}

	// ClusterAutoscalerAddonContainerValues holds the valid values for cluster-autoscaler addon container names
	ClusterAutoscalerAddonContainerValues = []string{"cluster-autoscaler"}

	// ClusterAutoscalerAddonConfigValues holds the valid values for cluster-autoscaler addon configuration overrides
	ClusterAutoscalerAddonConfigValues = []string{"min-nodes", "max-nodes", "scan-interval"}

	// NvidiaAddonContainerValues holds the valid values for nvidia addon container names
	NvidiaAddonContainerValues = []string{"nvidia-device-plugin"}

	// BlobfuseFlexvolumeAddonContainerValues holds the valid values for blobfuse-flexvolume addon container names
	BlobfuseFlexvolumeAddonContainerValues = []string{"blobfuse-flexvolume"}

	// SMBFlexvolumeAddonContainerValues holds the valid values for smb-flexvolume addon container names
	SMBFlexvolumeAddonContainerValues = []string{"smb-flexvolume"}

	// KeyvaultFlexvolumeAddonContainerValues holds the valid values for keyvault-flexvolume addon container names
	KeyvaultFlexvolumeAddonContainerValues = []string{"keyvault-flexvolume"}

	// MetricsServerAddonContainerValues holds the valid values for metrics-server addon container names
	MetricsServerAddonContainerValues = []string{"metrics-server"}

	// AzureCNINetworkMonitorAddonContainerValues holds the valid values for azure-cni-networkmonitor addon container names
	AzureCNINetworkMonitorAddonContainerValues = []string{"azure-cni-networkmonitor"}

	// DNSAutoscalerAddonContainerValues holds the valid values for dns-autoscaler addon container names
	DNSAutoscalerAddonContainerValues = []string{"dns-autoscaler"}

	// CalicoAddonContainerValues holds the valid values for calico addon container names
	CalicoAddonContainerValues = []string{"calico-typha", "calico-cni", "calico-node", "calico-cluster-proportional-autoscaler"}

	// AzureNetworkPolicyAddonContainerValues holds the valid values for azure-npm-daemonset addon container names
	AzureNetworkPolicyAddonContainerValues = []string{"azure-npm-daemonset"}
)

// Kubernetes configuration
const (
	// KubernetesMinMaxPods is the minimum valid value for MaxPods, necessary for running kube-system pods
	KubernetesMinMaxPods = 5
)

// vlabs default configuration
const (
	// DefaultNetworkPlugin defines the network plugin to use by default
	DefaultNetworkPlugin = "azure"
	// DefaultNetworkPluginWindows defines the network plugin to use by default for clusters with Windows agent pools
	DefaultNetworkPluginWindows = "azure"
	// DefaultNetworkPolicy defines the network policy to use by default
	DefaultNetworkPolicy = ""
	// NetworkPolicyCilium is the string expression for cilium network policy config option
	NetworkPolicyCilium = "cilium"
	// NetworkPluginCilium is the string expression for cilium network policy config option
	NetworkPluginCilium = NetworkPolicyCilium
)

const (
	// AgentPoolProfileRoleEmpty is the empty role
	AgentPoolProfileRoleEmpty AgentPoolProfileRole = ""
	// AgentPoolProfileRoleInfra is the infra role
	AgentPoolProfileRoleInfra AgentPoolProfileRole = "infra"
)

const (
	// AzureStackCloud is a const string reference identifier for Azure Stack cloud
	AzureStackCloud = "AzureStackCloud"
)

const (
	// AzureADIdentitySystem is a const string reference identifier for Azure AD identity System
	AzureADIdentitySystem = "azure_ad"
	// ADFSIdentitySystem is a const string reference identifier for ADFS identity System
	ADFSIdentitySystem = "adfs"
)

const (
	// AzureStackDependenciesLocationPublic indicates to get dependencies from in AzurePublic cloud
	AzureStackDependenciesLocationPublic = "public"
	// AzureStackDependenciesLocationChina indicates to get dependencies from AzureChina cloud
	AzureStackDependenciesLocationChina = "china"
	// AzureStackDependenciesLocationGerman indicates to get dependencies from AzureGerman cloud
	AzureStackDependenciesLocationGerman = "german"
	// AzureStackDependenciesLocationUSGovernment indicates to get dependencies from AzureUSGovernment cloud
	AzureStackDependenciesLocationUSGovernment = "usgovernment"
)

const (
	// ClientSecretAuthMethod indicates to use client seret for authentication
	ClientSecretAuthMethod = "client_secret"
	// ClientCertificateAuthMethod indicates to use client certificate for authentication
	ClientCertificateAuthMethod = "client_certificate"
)

// StandardLoadBalancerSku is the string const for Azure Standard Load Balancer
const StandardLoadBalancerSku = "Standard"
