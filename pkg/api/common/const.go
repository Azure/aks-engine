// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

// the orchestrators supported
const (
	// Mesos is the string constant for MESOS orchestrator type
	Mesos string = "Mesos"
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS string = "DCOS"
	// Swarm is the string constant for the Swarm orchestrator type
	Swarm string = "Swarm"
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
	// SwarmMode is the string constant for the Swarm Mode orchestrator type
	SwarmMode string = "SwarmMode"
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
	// address relative to the first consecutive Kubernetes static IP
	DefaultInternalLbStaticIPOffset = 10
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
	// Ephemeral means that the node's os disk is ephemeral. This is not compatible with attached volumes.
	Ephemeral = "Ephemeral"
)

const (
	// KubernetesDefaultRelease is the default Kubernetes release
	KubernetesDefaultRelease string = "1.13"
	// KubernetesDefaultReleaseWindows is the default Kubernetes release
	KubernetesDefaultReleaseWindows string = "1.14"
)

const (
	// DCOSVersion1Dot11Dot2 is the major.minor.patch string for 1.11.0 versions of DCOS
	DCOSVersion1Dot11Dot2 string = "1.11.2"
	// DCOSVersion1Dot11Dot0 is the major.minor.patch string for 1.11.0 versions of DCOS
	DCOSVersion1Dot11Dot0 string = "1.11.0"
	// DCOSVersion1Dot10Dot0 is the major.minor.patch string for 1.10.0 versions of DCOS
	DCOSVersion1Dot10Dot0 string = "1.10.0"
	// DCOSVersion1Dot9Dot0 is the major.minor.patch string for 1.9.0 versions of DCOS
	DCOSVersion1Dot9Dot0 string = "1.9.0"
	// DCOSVersion1Dot9Dot8 is the major.minor.patch string for 1.9.8 versions of DCOS
	DCOSVersion1Dot9Dot8 string = "1.9.8"
	// DCOSVersion1Dot8Dot8 is the major.minor.patch string for 1.8.8 versions of DCOS
	DCOSVersion1Dot8Dot8 string = "1.8.8"
	// DCOSDefaultVersion is the default major.minor.patch version for DCOS
	DCOSDefaultVersion string = DCOSVersion1Dot11Dot0
)

// AllDCOSSupportedVersions maintain a list of available dcos versions in aks-engine
var AllDCOSSupportedVersions = []string{
	DCOSVersion1Dot11Dot2,
	DCOSVersion1Dot11Dot0,
	DCOSVersion1Dot10Dot0,
	DCOSVersion1Dot9Dot8,
	DCOSVersion1Dot9Dot0,
	DCOSVersion1Dot8Dot8,
}

const (
	// SwarmVersion is the Swarm orchestrator version
	SwarmVersion = "swarm:1.1.0"
	// DockerCEVersion is the DockerCE orchestrator version
	DockerCEVersion = "17.03.*"
)

// GetAllSupportedDCOSVersions returns a slice of all supported DCOS versions.
func GetAllSupportedDCOSVersions() []string {
	return AllDCOSSupportedVersions
}

// GetAllSupportedSwarmVersions returns a slice of all supported Swarm versions.
func GetAllSupportedSwarmVersions() []string {
	return []string{SwarmVersion}
}

// GetAllSupportedDockerCEVersions returns a slice of all supported Docker CE versions.
func GetAllSupportedDockerCEVersions() []string {
	return []string{DockerCEVersion}
}

// MinCloudProviderQPSToBucketFactor defines the minimum ratio between QPS and Bucket size for cloudprovider rate limiting
const MinCloudProviderQPSToBucketFactor float64 = 0.1

// Addon name consts
const (
	// HeapsterAddonName is the name of the heapster addon
	HeapsterAddonName = "heapster"
	// TillerAddonName is the name of the tiller addon deployment
	TillerAddonName = "tiller"
	// ACIConnectorAddonName is the name of the aci-connector addon deployment
	ACIConnectorAddonName = "aci-connector"
	// ClusterAutoscalerAddonName is the name of the cluster autoscaler addon deployment
	ClusterAutoscalerAddonName = "cluster-autoscaler"
	// BlobfuseFlexVolumeAddonName is the name of the blobfuse flexvolume addon
	BlobfuseFlexVolumeAddonName = "blobfuse-flexvolume"
	// SMBFlexVolumeAddonName is the name of the smb flexvolume addon
	SMBFlexVolumeAddonName = "smb-flexvolume"
	// KeyVaultFlexVolumeAddonName is the name of the key vault flexvolume addon deployment
	KeyVaultFlexVolumeAddonName = "keyvault-flexvolume"
	// DashboardAddonName is the name of the kubernetes-dashboard addon deployment
	DashboardAddonName = "kubernetes-dashboard"
	// ReschedulerAddonName is the name of the rescheduler addon deployment
	ReschedulerAddonName = "rescheduler"
	// MetricsServerAddonName is the name of the kubernetes metrics server addon deployment
	MetricsServerAddonName = "metrics-server"
	// NVIDIADevicePluginAddonName is the name of the NVIDIA device plugin addon deployment
	NVIDIADevicePluginAddonName = "nvidia-device-plugin"
	// ContainerMonitoringAddonName is the name of the kubernetes Container Monitoring addon deployment
	ContainerMonitoringAddonName = "container-monitoring"
	// IPMASQAgentAddonName is the name of the ip masq agent addon
	IPMASQAgentAddonName = "ip-masq-agent"
	// AzureCNINetworkMonitorAddonName is the name of the Azure CNI networkmonitor addon
	AzureCNINetworkMonitorAddonName = "azure-cni-networkmonitor"
	// AzureNetworkPolicyAddonName is the name of the Azure network policy manager addon
	AzureNetworkPolicyAddonName = "azure-npm-daemonset"
	// AzureVnetTelemetryContainerName is the name of the Azure vnet telemetry container in the azure-npm-daemonset addon
	AzureVnetTelemetryContainerName = "azure-vnet-telemetry-daemonset"
	// CloudNodeManagerAddonName is the name of the cloud node manager addon
	CloudNodeManagerAddonName = "cloud-node-manager"
	// DNSAutoscalerAddonName is the name of the dns-autoscaler addon
	DNSAutoscalerAddonName = "dns-autoscaler"
	// CalicoAddonName is the name of calico daemonset addon
	CalicoAddonName = "calico-daemonset"
	// AADPodIdentityAddonName is the name of the aad-pod-identity addon deployment
	AADPodIdentityAddonName = "aad-pod-identity"
	// AzurePolicyAddonName is the name of the Azure Policy addon
	AzurePolicyAddonName = "azure-policy"
	// AppGwIngressAddonName appgw addon
	AppGwIngressAddonName = "appgw-ingress"
	// AzureDiskCSIDriverAddonName is the name of Azure Disk CSI Driver addon
	AzureDiskCSIDriverAddonName = "azuredisk-csi-driver"
	// AzureFileCSIDriverAddonName is the name of Azure File CSI Driver addon
	AzureFileCSIDriverAddonName = "azurefile-csi-driver"
	// AzureStorageClassesAddonName is the name of the azure storage classes addon
	AzureStorageClassesAddonName = "azure-storage-classes"
	// KubeDNSAddonName is the name of the kube-dns-deployment addon
	KubeDNSAddonName = "kube-dns"
	// CoreDNSAddonName is the name of the coredns addon
	CoreDNSAddonName = "coredns"
	// KubeProxyAddonName is the name of the kube-proxy config addon
	KubeProxyAddonName = "kube-proxy"
	// CiliumAddonName is the name of cilium daemonset addon
	CiliumAddonName = "cilium-daemonset"
	// FlannelAddonName is the name of flannel plugin daemonset addon
	FlannelAddonName = "flannel-daemonset"
	// AADAdminGroupAddonName is the name of the default admin group RBAC addon
	AADAdminGroupAddonName = "aad-default-admin-group-rbac"
	// AzureCloudProviderAddonName is the name of the azure cloud provider deployment addon
	AzureCloudProviderAddonName = "azure-cloud-provider-deployment"
	// AzureCSIStorageClassesAddonName is the name of Azure CSI storage classes addon
	AzureCSIStorageClassesAddonName = "azure-csi-storage-classes"
	// AuditPolicyAddonName is the name of the audit policy addon
	AuditPolicyAddonName = "audit-policy"
	// ScheduledMaintenanceAddonName is the name of the scheduled maintenance addon deployment
	ScheduledMaintenanceAddonName = "scheduled-maintenance"
	// PodSecurityPolicyAddonName is the name of the PodSecurityPolicy addon
	PodSecurityPolicyAddonName = "pod-security-policy"
	// NodeProblemDetectorAddonName is the name of the node problem detector addon
	NodeProblemDetectorAddonName = "node-problem-detector"
)
