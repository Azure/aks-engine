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
	// DefaultEnableCSIProxyWindows determines if CSI proxy should be enabled by default for Windows nodes
	DefaultEnableCSIProxyWindows = false
	// MaxLoadBalancerOutboundIPs is the maximum number of outbound IPs in a Standard LoadBalancer frontend configuration
	MaxLoadBalancerOutboundIPs = 16
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
	KubernetesDefaultRelease string = "1.15"
	// KubernetesDefaultReleaseWindows is the default Kubernetes release
	KubernetesDefaultReleaseWindows string = "1.15"
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
	// DashboardMetricsScraperContainerName is the name of the metrics-scraper container in the kubernetes-dashboard addon
	DashboardMetricsScraperContainerName = "kubernetes-dashboard-metrics-scraper"
	// ReschedulerAddonName is the name of the rescheduler addon deployment
	ReschedulerAddonName = "rescheduler"
	// ExecHealthZComponentName is the name of the exechealthz component
	ExecHealthZComponentName = "exechealthz"
	// MetricsServerAddonName is the name of the kubernetes metrics server addon deployment
	MetricsServerAddonName = "metrics-server"
	// AddonResizerComponentName is the name of the kubernetes addon-resizer component
	AddonResizerComponentName = "addon-resizer"
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
	// AzureVnetTelemetryContainerName is the name of the deprecated Azure vnet telemetry container in the azure-npm-daemonset addon
	AzureVnetTelemetryContainerName = "azure-vnet-telemetry-daemonset"
	// NMIContainerName is the name of the nmi container in the aad-pod-identity addon
	NMIContainerName = "nmi"
	// MICContainerName is the name of the mic container in the aad-pod-identity addon
	MICContainerName = "mic"
	// CiliumAgentContainerName is the name of the cilium-agent container in the cilium addon
	CiliumAgentContainerName = "cilium-agent"
	// CiliumCleanStateContainerName is the name of the clean-cilium-state container in the cilium addon
	CiliumCleanStateContainerName = "clean-cilium-state"
	// CiliumOperatorContainerName is the name of the cilium-operator container in the cilium addon
	CiliumOperatorContainerName = "cilium-operator"
	// CiliumEtcdOperatorContainerName is the name of the cilium-etcd-operator container in the cilium addon
	CiliumEtcdOperatorContainerName = "cilium-etcd-operator"
	// AntreaControllerContainerName is the name of the antrea-controller container in the antrea addon
	AntreaControllerContainerName = "antrea-controller"
	// AntreaAgentContainerName is the name of the antrea-agent container in the antrea addon
	AntreaAgentContainerName = "antrea-agent"
	// AntreaOVSContainerName is the name of the antrea-ovs container in the antrea addon
	AntreaOVSContainerName = "antrea-ovs"
	// AntreaInstallCNIContainerName is the name of the install-cni container in the antrea addon
	AntreaInstallCNIContainerName = "install-cni"
	// GatekeeperContainerName is the name of the gatekeeper container in the azure-policy addon
	GatekeeperContainerName = "gatekeeper"
	// CloudNodeManagerAddonName is the name of the cloud node manager addon
	CloudNodeManagerAddonName = "cloud-node-manager"
	// CalicoAddonName is the name of calico daemonset addon
	CalicoAddonName = "calico-daemonset"
	// CalicoTyphaComponentName is the name of calico-typha component
	CalicoTyphaComponentName = "calico-typha"
	// CalicoCNIComponentName is the name of calico-cni component
	CalicoCNIComponentName = "calico-cni"
	// CalicoNodeComponentName is the name of calico-node component
	CalicoNodeComponentName = "calico-node"
	// CalicoPod2DaemonComponentName is the name of calico-pod2daemon component
	CalicoPod2DaemonComponentName = "calico-pod2daemon"
	// CalicoClusterAutoscalerComponentName is the name of calico-cluster-proportional-autoscaler component
	CalicoClusterAutoscalerComponentName = "calico-cluster-proportional-autoscaler"
	// AADPodIdentityAddonName is the name of the aad-pod-identity addon deployment
	AADPodIdentityAddonName = "aad-pod-identity"
	// AzurePolicyAddonName is the name of the Azure Policy addon
	AzurePolicyAddonName = "azure-policy"
	// AppGwIngressAddonName appgw addon
	AppGwIngressAddonName = "appgw-ingress"
	// AzureDiskCSIDriverAddonName is the name of Azure Disk CSI Driver addon
	AzureDiskCSIDriverAddonName = "azuredisk-csi-driver"
	// CSIProvisionerContainerName is the name of the csi-provisioner container in the azuredisk-csi-driver and azurefile-csi-driver addons
	CSIProvisionerContainerName = "csi-provisioner"
	// CSIAttacherContainerName is the name of the csi-attacher container in the azuredisk-csi-driver and azurefile-csi-driver addons
	CSIAttacherContainerName = "csi-attacher"
	// CSILivenessProbeContainerName is the name of the livenessprobe container in the azuredisk-csi-driver, azurefile-csi-driver and secrets-store-csi-driver addons
	CSILivenessProbeContainerName = "livenessprobe"
	// CSILivenessProbeWindowsContainerName is the name of the livenessprobe-windows container in the azuredisk-csi-driver and azurefile-csi-driver addons
	CSILivenessProbeWindowsContainerName = "livenessprobe-windows"
	// CSISnapshotterContainerName is the name of the csi-snapshotter container in the azuredisk-csi-driver and azurefile-csi-driver addons
	CSISnapshotterContainerName = "csi-snapshotter"
	// CSISnapshotControllerContainerName is the name of the csi-snapshot-controller container
	CSISnapshotControllerContainerName = "csi-snapshot-controller"
	// CSIResizerContainerName is the name of the csi-resizer container in the azuredisk-csi-driver addon
	CSIResizerContainerName = "csi-resizer"
	// CSINodeDriverRegistrarContainerName is the name of the csi-node-driver-registrar container in the azuredisk-csi-driver, azurefile-csi-driver and secrets-store-csi-driver addons
	CSINodeDriverRegistrarContainerName = "csi-node-driver-registrar"
	// CSINodeDriverRegistrarWindowsContainerName is the name of the csi-node-driver-registrar-windows container in the azuredisk-csi-driver and azurefile-csi-driver addons
	CSINodeDriverRegistrarWindowsContainerName = "csi-node-driver-registrar-windows"
	// CSIAzureDiskContainerName is the name of the azuredisk-csi container in the azuredisk-csi-driver and azurefile-csi-driver addons
	CSIAzureDiskContainerName = "azuredisk-csi"
	// AzureFileCSIDriverAddonName is the name of Azure File CSI Driver addon
	AzureFileCSIDriverAddonName = "azurefile-csi-driver"
	// CSIAzureFileContainerName is the name of the azurefile-csi container in the azurefile-csi-driver addon
	CSIAzureFileContainerName = "azurefile-csi"
	// AzureStorageClassesAddonName is the name of the azure storage classes addon
	AzureStorageClassesAddonName = "azure-storage-classes"
	// Hyperkube is the common "hyperkube" string reference
	Hyperkube = "hyperkube"
	// KubeDNSAddonName is the name of the kube-dns-deployment addon
	KubeDNSAddonName = "kube-dns"
	// DNSMasqComponentName is the name of the dnsmasq component
	DNSMasqComponentName = "dnsmasq"
	// DNSSidecarComponentName is the name of the dnsmasq component
	DNSSidecarComponentName = "k8s-dns-sidecar"
	// PauseComponentName is the name of the pause component
	PauseComponentName = "pause"
	// CoreDNSAddonName is the name of the coredns addon
	CoreDNSAddonName = "coredns"
	// CoreDNSAutoscalerName is the name of the coredns-autoscaler container in the coredns addon
	CoreDNSAutoscalerName = "coredns-autoscaler"
	// KubeProxyAddonName is the name of the kube-proxy config addon
	KubeProxyAddonName = "kube-proxy"
	// CiliumAddonName is the name of cilium daemonset addon
	CiliumAddonName = "cilium"
	// AntreaAddonName is the name of antrea daemonset addon
	AntreaAddonName = "antrea"
	// FlannelAddonName is the name of flannel plugin daemonset addon
	FlannelAddonName = "flannel"
	// KubeFlannelContainerName is the name of the kube-flannel container in the flannel addon
	KubeFlannelContainerName = "kube-flannel"
	// FlannelInstallCNIContainerName is the name of the install-cni container in the flannel addon
	FlannelInstallCNIContainerName = "install-cni"
	// KubeRBACProxyContainerName is the name of the kube-rbac-proxy container in the scheduled-maintenance addon
	KubeRBACProxyContainerName = "kube-rbac-proxy"
	// ScheduledMaintenanceManagerContainerName is the name of the manager container in the scheduled-maintenance addon
	ScheduledMaintenanceManagerContainerName = "manager"
	// AADAdminGroupAddonName is the name of the default admin group RBAC addon
	AADAdminGroupAddonName = "aad"
	// AzureCloudProviderAddonName is the name of the azure-cloud-provider addon
	AzureCloudProviderAddonName = "azure-cloud-provider"
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
	// SecretsStoreCSIDriverAddonName is the name of the secrets-store-csi-driver addon
	SecretsStoreCSIDriverAddonName = "csi-secrets-store"
	// CSISecretsStoreDriverContainerName is the name of the secrets-store container in the csi-secrets-store addon
	CSISecretsStoreDriverContainerName = "secrets-store"
	// CSISecretsStoreProviderAzureContainerName is the name of the provider-azure-installer container in csi-secrets-store addon
	CSISecretsStoreProviderAzureContainerName = "provider-azure-installer"
)

// Component name consts
const (
	// SchedulerComponentName is the name of the kube-scheduler component
	SchedulerComponentName = "kube-scheduler"
	// ControllerManagerComponentName is the name of the kube-controller-manager component
	ControllerManagerComponentName = "kube-controller-manager"
	// CloudControllerManagerComponentName is the name of the cloud-controller-manager component
	CloudControllerManagerComponentName = "cloud-controller-manager"
	// APIServerComponentName is the name of the kube-apiserver component
	APIServerComponentName = "kube-apiserver"
	// AddonManagerComponentName is the name of the kube-addon-manager component
	AddonManagerComponentName = "kube-addon-manager"
	// ClusterInitComponentName is the name of the cluster-init component
	ClusterInitComponentName = "cluster-init"
)

const WindowsArtifactComponentName = "windowszip"
const WindowsArtifactAzureStackComponentName = "windowszip-azs"

const (
	// AzureStackSuffix is appended to kubernetes version on Azure Stack instances
	AzureStackSuffix = "-azs"
	// AzureStackPrefix is prepended to windows binary version for Azure Stack instances
	AzureStackPrefix = "azs-"
	// AzureStackCaCertLocation is where Azure Stack's CRP drops the stamp CA certificate
	AzureStackCaCertLocation = "/var/lib/waagent/Certificates.pem"
)

const (
	KubernetesImageBaseTypeGCR = "gcr"
	KubernetesImageBaseTypeMCR = "mcr"
)

var (
	// DefaultDockerConfig describes the default configuration of the docker daemon.
	DefaultDockerConfig = DockerConfig{
		LiveRestore: true,
		LogDriver:   "json-file",
		LogOpts: LogOpts{
			MaxSize: "50m",
			MaxFile: "5",
		},
	}

	// DefaultContainerdConfig describes the default configuration of the containerd daemon.
	DefaultContainerdConfig = ContainerdConfig{
		Version:  2,
		OomScore: 0,
		Plugins: Plugins{
			IoContainerdGrpcV1Cri: IoContainerdGrpcV1Cri{
				CNI: ContainerdCNIPlugin{},
				Containerd: ContainerdPlugin{
					DefaultRuntimeName: "runc",
					Runtimes: map[string]ContainerdRuntime{
						"runc": {
							RuntimeType: "io.containerd.runc.v2",
						},
						// note: runc really should not be used for untrusted workloads... should we remove this? This is here because it was here before
						"untrusted": {
							RuntimeType: "io.containerd.runc.v2",
						},
					},
				},
			},
		},
	}
)

// GetDefaultDockerConfig returns the default docker config for processing.
func GetDefaultDockerConfig() DockerConfig {
	return DefaultDockerConfig
}

// GetDefaultContainerdConfig returns the default containerd config for processing.
func GetDefaultContainerdConfig() ContainerdConfig {
	return DefaultContainerdConfig
}

// Known container runtime configuration keys
const (
	ContainerDataDirKey = "dataDir"
)

// Antrea Plugin Const
const (
	AntreaDefaultTrafficEncapMode = "Encap"
	AntreaDefaultInstallCniCmd    = "install_cni"
	AntreaInstallCniChainCmd      = "install_cni_chaining"
	AntreaNetworkPolicyOnlyMode   = "networkPolicyOnly"
)

// Node Taint consts
const (
	// MasterNodeTaint is the node taint we apply to all master nodes
	MasterNodeTaint string = "node-role.kubernetes.io/master=true:NoSchedule"
	// AADPodIdentityTaintKey is the node taint key for AAD Pod Identity-enabled clusters before NMI daemonset is ready
	AADPodIdentityTaintKey string = "node.kubernetes.io/aad-pod-identity-not-ready"
)
