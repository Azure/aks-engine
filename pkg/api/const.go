// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

// the orchestrators supported by vlabs
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

// the OSTypes supported by vlabs
const (
	Windows OSType = "Windows"
	Linux   OSType = "Linux"
)

// Distro string consts
const (
	Ubuntu            Distro = "ubuntu"
	Ubuntu1804        Distro = "ubuntu-18.04"
	Ubuntu1804Gen2    Distro = "ubuntu-18.04-gen2"
	RHEL              Distro = "rhel"
	CoreOS            Distro = "coreos"            // deprecated
	AKS1604Deprecated Distro = "aks"               // deprecated AKS 16.04 distro. Equivalent to aks-ubuntu-16.04.
	AKS1804Deprecated Distro = "aks-1804"          // deprecated AKS 18.04 distro. Equivalent to aks-ubuntu-18.04.
	AKSDockerEngine   Distro = "aks-docker-engine" // deprecated docker-engine distro.
	AKSUbuntu1604     Distro = "aks-ubuntu-16.04"
	AKSUbuntu1804     Distro = "aks-ubuntu-18.04"
	ACC1604           Distro = "acc-16.04"
)

const (
	// SwarmVersion is the Swarm orchestrator version
	SwarmVersion = "swarm:1.1.0"
	// SwarmDockerComposeVersion is the Docker Compose version
	SwarmDockerComposeVersion = "1.6.2"
	// DockerCEVersion is the DockerCE orchestrator version
	DockerCEVersion = "17.03.*"
	// DockerCEDockerComposeVersion is the Docker Compose version
	DockerCEDockerComposeVersion = "1.14.0"
	// KubernetesWindowsDockerVersion is the default version for docker on Windows nodes in kubernetes
	KubernetesWindowsDockerVersion = "19.03.5"
	// KubernetesDefaultWindowsSku is the default SKU for Windows VMs in kubernetes
	KubernetesDefaultWindowsSku = "Datacenter-Core-1809-with-Containers-smalldisk"
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
)

// Availability profiles
const (
	// AvailabilitySet means that the vms are in an availability set
	AvailabilitySet = "AvailabilitySet"
	// DefaultOrchestratorName specifies the 3 character orchestrator code of the cluster template and affects resource naming.
	DefaultOrchestratorName = "k8s"
	// DefaultHostedProfileMasterName specifies the 3 character orchestrator code of the clusters with hosted master profiles.
	DefaultHostedProfileMasterName = "aks"
	// DefaultFirstConsecutiveKubernetesStaticIP specifies the static IP address on Kubernetes master 0
	DefaultFirstConsecutiveKubernetesStaticIP = "10.240.255.5"
	// DefaultFirstConsecutiveKubernetesStaticIPVMSS specifies the static IP address on Kubernetes master 0 of VMSS
	DefaultFirstConsecutiveKubernetesStaticIPVMSS = "10.240.0.4"
	//DefaultCNICIDR specifies the default value for
	DefaultCNICIDR = "168.63.129.16/32"
	// DefaultKubernetesFirstConsecutiveStaticIPOffset specifies the IP address offset of master 0
	// when VNET integration is enabled.
	DefaultKubernetesFirstConsecutiveStaticIPOffset = 5
	// DefaultKubernetesFirstConsecutiveStaticIPOffsetVMSS specifies the IP address offset of master 0 in VMSS
	// when VNET integration is enabled.
	DefaultKubernetesFirstConsecutiveStaticIPOffsetVMSS = 4
	// DefaultSubnetNameResourceSegmentIndex specifies the default subnet name resource segment index.
	DefaultSubnetNameResourceSegmentIndex = 10
	// DefaultVnetResourceGroupSegmentIndex specifies the default virtual network resource segment index.
	DefaultVnetResourceGroupSegmentIndex = 4
	// DefaultVnetNameResourceSegmentIndex specifies the default virtual network name segment index.
	DefaultVnetNameResourceSegmentIndex = 8
	// VirtualMachineScaleSets means that the vms are in a virtual machine scaleset
	VirtualMachineScaleSets = "VirtualMachineScaleSets"
	// ScaleSetPriorityRegular is the default ScaleSet Priority
	ScaleSetPriorityRegular = "Regular"
	// ScaleSetPriorityLow means the ScaleSet will use Low-priority VMs
	ScaleSetPriorityLow = "Low"
	// ScaleSetPrioritySpot means the ScaleSet will use Spot VMs
	ScaleSetPrioritySpot = "Spot"
	// ScaleSetEvictionPolicyDelete is the default Eviction Policy for Low-priority VM ScaleSets
	ScaleSetEvictionPolicyDelete = "Delete"
	// ScaleSetEvictionPolicyDeallocate means a Low-priority VM ScaleSet will deallocate, rather than delete, VMs.
	ScaleSetEvictionPolicyDeallocate = "Deallocate"
)

// Supported container runtimes
const (
	Docker         = "docker"
	KataContainers = "kata-containers" // Deprecated
	Containerd     = "containerd"
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

// To identify programmatically generated public agent pools
const publicAgentPoolSuffix = "-public"

const (
	// DefaultHeapsterAddonEnabled determines the aks-engine provided default for enabling heapster addon
	DefaultHeapsterAddonEnabled = false
	// DefaultTillerAddonEnabled determines the aks-engine provided default for enabling tiller addon
	DefaultTillerAddonEnabled = false
	// DefaultAADPodIdentityAddonEnabled determines the aks-engine provided default for enabling aad-pod-identity addon
	DefaultAADPodIdentityAddonEnabled = false
	// DefaultAzurePolicyAddonEnabled determines the aks-engine provided default for enabling azure policy addon
	DefaultAzurePolicyAddonEnabled = false
	// DefaultNodeProblemDetectorAddonEnabled determines the aks-engine provided default for enabling the node problem detector addon
	DefaultNodeProblemDetectorAddonEnabled = false
	// DefaultACIConnectorAddonEnabled determines the aks-engine provided default for enabling aci connector addon
	DefaultACIConnectorAddonEnabled = false
	// DefaultAppGwIngressAddonEnabled determines the aks-engine provided default for enabling appgw ingress addon
	DefaultAppGwIngressAddonEnabled = false
	// DefaultAzureDiskCSIDriverAddonEnabled determines the aks-engine provided default for enabling Azure Disk CSI Driver
	DefaultAzureDiskCSIDriverAddonEnabled = true
	// DefaultAzureFileCSIDriverAddonEnabled determines the aks-engine provided default for enabling Azure File CSI Driver
	DefaultAzureFileCSIDriverAddonEnabled = true
	// DefaultClusterAutoscalerAddonEnabled determines the aks-engine provided default for enabling cluster autoscaler addon
	DefaultClusterAutoscalerAddonEnabled = false
	// DefaultBlobfuseFlexVolumeAddonEnabled determines the aks-engine provided default for enabling blobfuse flexvolume addon
	DefaultBlobfuseFlexVolumeAddonEnabled = true
	// DefaultSMBFlexVolumeAddonEnabled determines the aks-engine provided default for enabling smb flexvolume addon
	DefaultSMBFlexVolumeAddonEnabled = false
	// DefaultKeyVaultFlexVolumeAddonEnabled determines the aks-engine provided default for enabling key vault flexvolume addon
	DefaultKeyVaultFlexVolumeAddonEnabled = true
	// DefaultDashboardAddonEnabled determines the aks-engine provided default for enabling kubernetes-dashboard addon
	DefaultDashboardAddonEnabled = false
	// DefaultReschedulerAddonEnabled determines the aks-engine provided default for enabling kubernetes-rescheduler addon
	DefaultReschedulerAddonEnabled = false
	// DefaultAzureCNIMonitoringAddonEnabled determines the aks-engine provided default for enabling azurecni-network monitoring addon
	DefaultAzureCNIMonitoringAddonEnabled = true
	// DefaultKubeDNSAddonEnabled determines the aks-engine provided default for enabling coredns addon
	DefaultKubeDNSAddonEnabled = false
	// DefaultCoreDNSAddonEnabled determines the aks-engine provided default for enabling coredns addon
	DefaultCoreDNSAddonEnabled = true
	// DefaultKubeProxyAddonEnabled determines the aks-engine provided default for enabling kube-proxy addon
	DefaultKubeProxyAddonEnabled = true
	// DefaultSecretStoreCSIDriverAddonEnabled determines the aks-engine provided default for enabling secrets-store-csi-driver addon
	DefaultSecretStoreCSIDriverAddonEnabled = true
	// DefaultRBACEnabled determines the aks-engine provided default for enabling kubernetes RBAC
	DefaultRBACEnabled = true
	// DefaultUseInstanceMetadata determines the aks-engine provided default for enabling Azure cloudprovider instance metadata service
	DefaultUseInstanceMetadata = true
	// BasicLoadBalancerSku is the string const for Azure Basic Load Balancer
	BasicLoadBalancerSku = "Basic"
	// StandardLoadBalancerSku is the string const for Azure Standard Load Balancer
	StandardLoadBalancerSku = "Standard"
	// DefaultExcludeMasterFromStandardLB determines the aks-engine provided default for excluding master nodes from standard load balancer.
	DefaultExcludeMasterFromStandardLB = true
	// DefaultSecureKubeletEnabled determines the aks-engine provided default for securing kubelet communications
	DefaultSecureKubeletEnabled = true
	// DefaultMetricsServerAddonEnabled determines the aks-engine provided default for enabling kubernetes metrics-server addon
	DefaultMetricsServerAddonEnabled = true
	// DefaultNVIDIADevicePluginAddonEnabled determines the aks-engine provided default for enabling NVIDIA Device Plugin
	DefaultNVIDIADevicePluginAddonEnabled = false
	// DefaultContainerMonitoringAddonEnabled determines the aks-engine provided default for enabling kubernetes container monitoring addon
	DefaultContainerMonitoringAddonEnabled = false
	// DefaultIPMasqAgentAddonEnabled enables the ip-masq-agent addon
	DefaultIPMasqAgentAddonEnabled = true
	// DefaultPrivateClusterEnabled determines the aks-engine provided default for enabling kubernetes Private Cluster
	DefaultPrivateClusterEnabled = false
	// NetworkPolicyAzure is the string expression for Azure CNI network policy manager
	NetworkPolicyAzure = "azure"
	// NetworkPolicyNone is the string expression for the deprecated NetworkPolicy usage pattern "none"
	NetworkPolicyNone = "none"
	// NetworkPluginKubenet is the string expression for the kubenet NetworkPlugin config
	NetworkPluginKubenet = "kubenet"
	// NetworkPluginAzure is the string expression for Azure CNI plugin.
	NetworkPluginAzure = "azure"
	// DefaultSinglePlacementGroup determines the aks-engine provided default for supporting large VMSS
	// (true = single placement group 0-100 VMs, false = multiple placement group 0-1000 VMs)
	DefaultSinglePlacementGroup = true
	// ARMNetworkNamespace is the ARM-specific namespace for ARM's network providers.
	ARMNetworkNamespace = "Microsoft.Networks"
	// ARMVirtualNetworksResourceType is the ARM resource type for virtual network resources of ARM.
	ARMVirtualNetworksResourceType = "virtualNetworks"
	// DefaultAcceleratedNetworkingWindowsEnabled determines the aks-engine provided default for enabling accelerated networking on Windows nodes
	DefaultAcceleratedNetworkingWindowsEnabled = false
	// DefaultAcceleratedNetworking determines the aks-engine provided default for enabling accelerated networking on Linux nodes
	DefaultAcceleratedNetworking = true
	// DefaultVMSSOverProvisioningEnabled determines the aks-engine provided default for enabling VMSS Overprovisioning
	DefaultVMSSOverProvisioningEnabled = false
	// DefaultAuditDEnabled determines the aks-engine provided default for enabling auditd
	DefaultAuditDEnabled = false
	// DefaultUseCosmos determines if the cluster will use cosmos as etcd storage
	DefaultUseCosmos = false
	// etcdEndpointURIFmt is the name format for a typical etcd account uri
	etcdEndpointURIFmt = "%sk8s.etcd.cosmosdb.azure.com"
	// DefaultMaximumLoadBalancerRuleCount determines the default value of maximum allowed loadBalancer rule count according to
	// https://docs.microsoft.com/en-us/azure/azure-subscription-service-limits#load-balancer.
	DefaultMaximumLoadBalancerRuleCount = 250
	// DefaultEnableAutomaticUpdates determines the aks-engine provided default for enabling automatic updates
	DefaultEnableAutomaticUpdates = false
	// DefaultPreserveNodesProperties determines the aks-engine provided default for preserving nodes properties
	DefaultPreserveNodesProperties = true
	// DefaultEnableVMSSNodePublicIP determines the aks-engine provided default for enable VMSS node public IP
	DefaultEnableVMSSNodePublicIP = false
	// DefaultOutboundRuleIdleTimeoutInMinutes determines the aks-engine provided default for IdleTimeoutInMinutes of the OutboundRule of the agent loadbalancer
	// This value is set greater than the default Linux idle timeout (15.4 min): https://pracucci.com/linux-tcp-rto-min-max-and-tcp-retries2.html
	DefaultOutboundRuleIdleTimeoutInMinutes = 30
	// AddonModeEnsureExists
	AddonModeEnsureExists = "EnsureExists"
	// AddonModeReconcile
	AddonModeReconcile = "Reconcile"
	// VMSSVMType is the string const for the vmss VM Type
	VMSSVMType = "vmss"
	// StandardVMType is the string const for the standard VM Type
	StandardVMType = "standard"
)

// Azure API Versions
const (
	APIVersionAuthorizationUser   = "2018-09-01-preview"
	APIVersionAuthorizationSystem = "2018-01-01-preview"
	APIVersionCompute             = "2019-07-01"
	APIVersionDeployments         = "2018-06-01"
	APIVersionKeyVault            = "2018-02-14"
	APIVersionManagedIdentity     = "2018-11-30"
	APIVersionNetwork             = "2018-08-01"
	APIVersionStorage             = "2018-07-01"
)

// AzureStackCloud Specific Defaults
const (
	// DefaultUseInstanceMetadata set to false as Azure Stack today doesn't support instance metadata service
	DefaultAzureStackUseInstanceMetadata = false
	// DefaultAzureStackAcceleratedNetworking set to false as Azure Stack today doesn't support accelerated networking
	DefaultAzureStackAcceleratedNetworking = false
	// DefaultAzureStackAvailabilityProfile set to AvailabilitySet as VMSS clusters are not suppored on Azure Stack
	DefaultAzureStackAvailabilityProfile = AvailabilitySet
	// DefaultAzureStackFaultDomainCount set to 3 as Azure Stack today has minimum 4 node deployment
	DefaultAzureStackFaultDomainCount = 3
	// MaxAzureStackManagedDiskSize is the size in GB of the etcd disk volumes when total nodes count is greater than 10
	MaxAzureStackManagedDiskSize = "1023"
	// AzureStackSuffix is appended to kubernetes version on Azure Stack instances
	AzureStackSuffix = "-azs"
	// DefaultAzureStackLoadBalancerSku determines the aks-engine provided default for enabling Azure cloudprovider load balancer SKU on Azure Stack
	DefaultAzureStackLoadBalancerSku = BasicLoadBalancerSku
)

const (
	// AgentPoolProfileRoleEmpty is the empty role.  Deprecated; only used in
	// aks-engine.
	AgentPoolProfileRoleEmpty AgentPoolProfileRole = ""
	// AgentPoolProfileRoleCompute is the compute role
	AgentPoolProfileRoleCompute AgentPoolProfileRole = "compute"
	// AgentPoolProfileRoleInfra is the infra role
	AgentPoolProfileRoleInfra AgentPoolProfileRole = "infra"
	// AgentPoolProfileRoleMaster is the master role
	AgentPoolProfileRoleMaster AgentPoolProfileRole = "master"
)

const (
	// VHDDiskSizeAKS maps to the OSDiskSizeGB for AKS VHD image
	VHDDiskSizeAKS = 30
)

const (
	CloudProviderBackoffModeV2 = "v2"
	// DefaultKubernetesCloudProviderBackoffRetries is 6, takes effect if DefaultKubernetesCloudProviderBackoff is true
	DefaultKubernetesCloudProviderBackoffRetries = 6
	// DefaultKubernetesCloudProviderBackoffJitter is 1, takes effect if DefaultKubernetesCloudProviderBackoff is true
	DefaultKubernetesCloudProviderBackoffJitter = 1.0
	// DefaultKubernetesCloudProviderBackoffDuration is 5, takes effect if DefaultKubernetesCloudProviderBackoff is true
	DefaultKubernetesCloudProviderBackoffDuration = 5
	// DefaultKubernetesCloudProviderBackoffExponent is 1.5, takes effect if DefaultKubernetesCloudProviderBackoff is true
	DefaultKubernetesCloudProviderBackoffExponent = 1.5
	// DefaultKubernetesCloudProviderRateLimitQPS is 3, takes effect if DefaultKubernetesCloudProviderRateLimit is true
	DefaultKubernetesCloudProviderRateLimitQPS = 3.0
	// DefaultKubernetesCloudProviderRateLimitQPSWrite is 1, takes effect if DefaultKubernetesCloudProviderRateLimit is true
	DefaultKubernetesCloudProviderRateLimitQPSWrite = 1.0
	// DefaultKubernetesCloudProviderRateLimitBucket is 10, takes effect if DefaultKubernetesCloudProviderRateLimit is true
	DefaultKubernetesCloudProviderRateLimitBucket = 10
	// DefaultKubernetesCloudProviderRateLimitBucketWrite is 10, takes effect if DefaultKubernetesCloudProviderRateLimit is true
	DefaultKubernetesCloudProviderRateLimitBucketWrite = DefaultKubernetesCloudProviderRateLimitBucket
)

// Azure Stack configures all clusters as if they were large clusters.
const (
	DefaultAzureStackKubernetesCloudProviderBackoffRetries       = 1
	DefaultAzureStackKubernetesCloudProviderBackoffJitter        = 1.0
	DefaultAzureStackKubernetesCloudProviderBackoffDuration      = 30
	DefaultAzureStackKubernetesCloudProviderBackoffExponent      = 1.5
	DefaultAzureStackKubernetesCloudProviderRateLimitQPS         = 3.0
	DefaultAzureStackKubernetesCloudProviderRateLimitQPSWrite    = 3.0
	DefaultAzureStackKubernetesCloudProviderRateLimitBucket      = 10
	DefaultAzureStackKubernetesCloudProviderRateLimitBucketWrite = 10
	DefaultAzureStackKubernetesNodeStatusUpdateFrequency         = "1m"
	DefaultAzureStackKubernetesCtrlMgrRouteReconciliationPeriod  = "1m"
	DefaultAzureStackKubernetesCtrlMgrNodeMonitorGracePeriod     = "5m"
	DefaultAzureStackKubernetesCtrlMgrPodEvictionTimeout         = "5m"
)

const (
	//AzureEdgeDCOSBootstrapDownloadURL is the azure edge CDN download url
	AzureEdgeDCOSBootstrapDownloadURL = "https://dcosio.azureedge.net/dcos/%s/bootstrap/%s.bootstrap.tar.xz"
	//AzureChinaCloudDCOSBootstrapDownloadURL is the China specific DCOS package download url.
	AzureChinaCloudDCOSBootstrapDownloadURL = "https://acsengine.blob.core.chinacloudapi.cn/dcos/%s.bootstrap.tar.xz"
	//AzureEdgeDCOSWindowsBootstrapDownloadURL
)

const (
	// AzureCniPluginVerLinux specifies version of Azure CNI plugin, which has been mirrored from
	// https://github.com/Azure/azure-container-networking/releases/download/${AZURE_PLUGIN_VER}/azure-vnet-cni-linux-amd64-${AZURE_PLUGIN_VER}.tgz
	// to https://kubernetesartifacts.azureedge.net/azure-cni
	AzureCniPluginVerLinux = "v1.1.2"
	// AzureCniPluginVerWindows specifies version of Azure CNI plugin, which has been mirrored from
	// https://github.com/Azure/azure-container-networking/releases/download/${AZURE_PLUGIN_VER}/azure-vnet-cni-windows-amd64-${AZURE_PLUGIN_VER}.zip
	// to https://kubernetesartifacts.azureedge.net/azure-cni
	AzureCniPluginVerWindows = "v1.1.2"
	// CNIPluginVer specifies the version of CNI implementation
	// https://github.com/containernetworking/plugins
	CNIPluginVer = "v0.8.5"
)

const (
	// DefaultMasterSubnet specifies the default master subnet for DCOS or Swarm
	DefaultMasterSubnet = "172.16.0.0/24"
	// DefaultFirstConsecutiveStaticIP specifies the static IP address on master 0 for DCOS or Swarm
	DefaultFirstConsecutiveStaticIP = "172.16.0.5"
	// DefaultSwarmWindowsMasterSubnet specifies the default master subnet for a Swarm Windows cluster
	DefaultSwarmWindowsMasterSubnet = "192.168.255.0/24"
	// DefaultSwarmWindowsFirstConsecutiveStaticIP specifies the static IP address on master 0 for a Swarm WIndows cluster
	DefaultSwarmWindowsFirstConsecutiveStaticIP = "192.168.255.5"
	// DefaultDCOSMasterSubnet specifies the default master subnet for a DCOS cluster
	DefaultDCOSMasterSubnet = "192.168.255.0/24"
	// DefaultDCOSFirstConsecutiveStaticIP  specifies the static IP address on master 0 for a DCOS cluster
	DefaultDCOSFirstConsecutiveStaticIP = "192.168.255.5"
	// DefaultDCOSBootstrapStaticIP specifies the static IP address on bootstrap for a DCOS cluster
	DefaultDCOSBootstrapStaticIP = "192.168.255.240"
	// DefaultKubernetesMasterSubnet specifies the default subnet for masters and agents.
	// Except when master VMSS is used, this specifies the default subnet for masters.
	DefaultKubernetesMasterSubnet = "10.240.0.0/16"
	// DefaultKubernetesMasterSubnetIPv6 specifies the default IPv6 subnet for masters and agents.
	// Except when master VMSS is used, this specifies the default subnet for masters.
	DefaultKubernetesMasterSubnetIPv6 = "2001:1234:5678:9abc::/64"
	// DefaultAgentSubnetTemplate specifies a default agent subnet
	DefaultAgentSubnetTemplate = "10.%d.0.0/16"
	// DefaultKubernetesSubnet specifies the default subnet used for all masters, agents and pods
	// when VNET integration is enabled.
	DefaultKubernetesSubnet = "10.240.0.0/12"
	// DefaultVNETCIDR is the default CIDR block for the VNET
	DefaultVNETCIDR = "10.0.0.0/8"
	// DefaultVNETCIDRIPv6 is the default IPv6 CIDR block for the VNET
	DefaultVNETCIDRIPv6 = "2001:1234:5678:9a00::/56"
	// DefaultKubernetesMaxPods is the maximum number of pods to run on a node.
	DefaultKubernetesMaxPods = 110
	// DefaultKubernetesMaxPodsVNETIntegrated is the maximum number of pods to run on a node when VNET integration is enabled.
	DefaultKubernetesMaxPodsVNETIntegrated = 30
	// DefaultKubernetesClusterDomain is the dns suffix used in the cluster (used as a SAN in the PKI generation)
	DefaultKubernetesClusterDomain = "cluster.local"
	// DefaultInternalLbStaticIPOffset specifies the offset of the internal LoadBalancer's IP
	// address relative to the first consecutive Kubernetes static IP
	DefaultInternalLbStaticIPOffset = 10
	// NetworkPolicyCalico is the string expression for calico network policy config option
	NetworkPolicyCalico = "calico"
	// NetworkPolicyCilium is the string expression for cilium network policy config option
	NetworkPolicyCilium = "cilium"
	// NetworkPluginCilium is the string expression for cilium network plugin config option
	NetworkPluginCilium = NetworkPolicyCilium
	// NetworkPluginFlannel is the string expression for flannel network policy config option
	NetworkPluginFlannel = "flannel"
	// NetworkPluginAntrea is the string expression for antrea network plugin config option
	NetworkPluginAntrea = "antrea"
	// NetworkPolicyAntrea is the string expression for antrea network policy config option
	NetworkPolicyAntrea = NetworkPluginAntrea
	// DefaultNetworkPlugin defines the network plugin to use by default
	DefaultNetworkPlugin = NetworkPluginKubenet
	// DefaultNetworkPolicy defines the network policy implementation to use by default
	DefaultNetworkPolicy = ""
	// DefaultNetworkPluginWindows defines the network plugin implementation to use by default for clusters with Windows agent pools
	DefaultNetworkPluginWindows = NetworkPluginKubenet
	// DefaultNetworkPolicyWindows defines the network policy implementation to use by default for clusters with Windows agent pools
	DefaultNetworkPolicyWindows = ""
	// DefaultContainerRuntime is docker
	DefaultContainerRuntime = Docker
	// DefaultKubernetesNodeStatusUpdateFrequency is 10s, see --node-status-update-frequency at https://kubernetes.io/docs/admin/kubelet/
	DefaultKubernetesNodeStatusUpdateFrequency = "10s"
	// DefaultKubernetesHardEvictionThreshold is memory.available<100Mi,nodefs.available<10%,nodefs.inodesFree<5%, see --eviction-hard at https://kubernetes.io/docs/admin/kubelet/
	DefaultKubernetesHardEvictionThreshold = "memory.available<750Mi,nodefs.available<10%,nodefs.inodesFree<5%"
	// DefaultKubernetesCtrlMgrNodeMonitorGracePeriod is 40s, see --node-monitor-grace-period at https://kubernetes.io/docs/admin/kube-controller-manager/
	DefaultKubernetesCtrlMgrNodeMonitorGracePeriod = "40s"
	// DefaultKubernetesCtrlMgrPodEvictionTimeout is 5m0s, see --pod-eviction-timeout at https://kubernetes.io/docs/admin/kube-controller-manager/
	DefaultKubernetesCtrlMgrPodEvictionTimeout = "5m0s"
	// DefaultKubernetesCtrlMgrRouteReconciliationPeriod is 10s, see --route-reconciliation-period at https://kubernetes.io/docs/admin/kube-controller-manager/
	DefaultKubernetesCtrlMgrRouteReconciliationPeriod = "10s"
	// DefaultKubernetesCtrlMgrTerminatedPodGcThreshold is set to 5000, see --terminated-pod-gc-threshold at https://kubernetes.io/docs/admin/kube-controller-manager/ and https://github.com/kubernetes/kubernetes/issues/22680
	DefaultKubernetesCtrlMgrTerminatedPodGcThreshold = "5000"
	// DefaultKubernetesCtrlMgrUseSvcAccountCreds is "true", see --use-service-account-credentials at https://kubernetes.io/docs/admin/kube-controller-manager/
	DefaultKubernetesCtrlMgrUseSvcAccountCreds = "false"
	// DefaultKubernetesCloudProviderRateLimit is false to disable cloudprovider rate limiting implementation for API calls
	DefaultKubernetesCloudProviderRateLimit = true
	// DefaultTillerMaxHistory limits the maximum number of revisions saved per release. Use 0 for no limit.
	DefaultTillerMaxHistory = 0
	//DefaultKubernetesGCHighThreshold specifies the value for  for the image-gc-high-threshold kubelet flag
	DefaultKubernetesGCHighThreshold = 85
	//DefaultKubernetesGCLowThreshold specifies the value for the image-gc-low-threshold kubelet flag
	DefaultKubernetesGCLowThreshold = 80
	// DefaultEtcdVersion specifies the default etcd version to install
	DefaultEtcdVersion = "3.3.19"
	// DefaultEtcdDiskSize specifies the default size for Kubernetes master etcd disk volumes in GB
	DefaultEtcdDiskSize = "256"
	// DefaultEtcdDiskSizeGT3Nodes = size for Kubernetes master etcd disk volumes in GB if > 3 nodes
	DefaultEtcdDiskSizeGT3Nodes = "512"
	// DefaultEtcdDiskSizeGT10Nodes = size for Kubernetes master etcd disk volumes in GB if > 10 nodes
	DefaultEtcdDiskSizeGT10Nodes = "1024"
	// DefaultEtcdDiskSizeGT20Nodes = size for Kubernetes master etcd disk volumes in GB if > 20 nodes
	DefaultEtcdDiskSizeGT20Nodes = "2048"
	// DefaultMasterEtcdClientPort is the default etcd client port for Kubernetes master nodes
	DefaultMasterEtcdClientPort = 2379
	// DefaultKubeletEventQPS is 0, see --event-qps at https://kubernetes.io/docs/reference/generated/kubelet/
	DefaultKubeletEventQPS = "0"
	// DefaultKubeletCadvisorPort is 0, see --cadvisor-port at https://kubernetes.io/docs/reference/generated/kubelet/
	DefaultKubeletCadvisorPort = "0"
	// DefaultJumpboxDiskSize specifies the default size for private cluster jumpbox OS disk in GB
	DefaultJumpboxDiskSize = 30
	// DefaultJumpboxUsername specifies the default admin username for the private cluster jumpbox
	DefaultJumpboxUsername = "azureuser"
	// DefaultKubeletPodMaxPIDs specifies the default max pid authorized by pods
	DefaultKubeletPodMaxPIDs = -1
	// DefaultKubernetesAgentSubnetVMSS specifies the default subnet for agents when master is VMSS
	DefaultKubernetesAgentSubnetVMSS = "10.248.0.0/13"
	// DefaultKubernetesClusterSubnet specifies the default subnet for pods.
	DefaultKubernetesClusterSubnet = "10.244.0.0/16"
	// DefaultKubernetesClusterSubnetIPv6 specifies the IPv6 default subnet for pods.
	DefaultKubernetesClusterSubnetIPv6 = "fc00::/48"
	// DefaultKubernetesServiceCIDR specifies the IP subnet that kubernetes will create Service IPs within.
	DefaultKubernetesServiceCIDR = "10.0.0.0/16"
	// DefaultKubernetesDNSServiceIP specifies the IP address that kube-dns listens on by default. must by in the default Service CIDR range.
	DefaultKubernetesDNSServiceIP = "10.0.0.10"
	// DefaultKubernetesServiceCIDRIPv6 specifies the IPv6 subnet that kubernetes will create Service IPs within.
	DefaultKubernetesServiceCIDRIPv6 = "fd00::/108"
	// DefaultKubernetesDNSServiceIPv6 specifies the IPv6 address that kube-dns listens on by default. must by in the default Service CIDR range.
	DefaultKubernetesDNSServiceIPv6 = "fd00::10"
	// DefaultMobyVersion specifies the default Azure build version of Moby to install.
	DefaultMobyVersion = "3.0.11"
	// DefaultContainerdVersion specifies the default containerd version to install.
	DefaultContainerdVersion = "1.3.2"
	// DefaultDockerBridgeSubnet specifies the default subnet for the docker bridge network for masters and agents.
	DefaultDockerBridgeSubnet = "172.17.0.1/16"
	// DefaultKubernetesMaxPodsKubenet is the maximum number of pods to run on a node for Kubenet.
	DefaultKubernetesMaxPodsKubenet = "110"
	// DefaultKubernetesMaxPodsAzureCNI is the maximum number of pods to run on a node for Azure CNI.
	DefaultKubernetesMaxPodsAzureCNI = "30"
	// DefaultKubernetesAPIServerEnableProfiling is the config that enables profiling via web interface host:port/debug/pprof/
	DefaultKubernetesAPIServerEnableProfiling = "false"
	// DefaultKubernetesCtrMgrEnableProfiling is the config that enables profiling via web interface host:port/debug/pprof/
	DefaultKubernetesCtrMgrEnableProfiling = "false"
	// DefaultKubernetesSchedulerEnableProfiling is the config that enables profiling via web interface host:port/debug/pprof/
	DefaultKubernetesSchedulerEnableProfiling = "false"
	// DefaultNonMasqueradeCIDR is the default --non-masquerade-cidr value for kubelet
	DefaultNonMasqueradeCIDR = "0.0.0.0/0"
	// DefaultKubeProxyMode is the default KubeProxyMode value
	DefaultKubeProxyMode KubeProxyMode = KubeProxyModeIPTables
	// DefaultWindowsSSHEnabled is the default windowsProfile.sshEnabled value
	DefaultWindowsSSHEnabled = true
)

const (
	//DefaultExtensionsRootURL  Root URL for extensions
	DefaultExtensionsRootURL = "https://raw.githubusercontent.com/Azure/aks-engine/master/"
)

const (
	// AzurePublicCloud is a const string reference identifier for public cloud
	AzurePublicCloud = "AzurePublicCloud"
	// AzureChinaCloud is a const string reference identifier for china cloud
	AzureChinaCloud = "AzureChinaCloud"
	// AzureGermanCloud is a const string reference identifier for german cloud
	AzureGermanCloud = "AzureGermanCloud"
	// AzureUSGovernmentCloud is a const string reference identifier for us government cloud
	AzureUSGovernmentCloud = "AzureUSGovernmentCloud"
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
	// AzureCustomCloudDependenciesLocationPublic indicates to get dependencies from in AzurePublic cloud
	AzureCustomCloudDependenciesLocationPublic = "public"
	// AzureCustomCloudDependenciesLocationChina indicates to get dependencies from AzureChina cloud
	AzureCustomCloudDependenciesLocationChina = "china"
	// AzureCustomCloudDependenciesLocationGerman indicates to get dependencies from AzureGerman cloud
	AzureCustomCloudDependenciesLocationGerman = "german"
	// AzureCustomCloudDependenciesLocationUSGovernment indicates to get dependencies from AzureUSGovernment cloud
	AzureCustomCloudDependenciesLocationUSGovernment = "usgovernment"
)

const (
	// ClientSecretAuthMethod indicates to use client seret for authentication
	ClientSecretAuthMethod = "client_secret"
	// ClientCertificateAuthMethod indicates to use client certificate for authentication
	ClientCertificateAuthMethod = "client_certificate"
)

// TLSStrongCipherSuitesAPIServer is a kube-bench-recommended allowed cipher suites for apiserver
const TLSStrongCipherSuitesAPIServer = "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"

// TLSStrongCipherSuitesKubelet is a kube-bench-recommended allowed cipher suites for kubelet
const TLSStrongCipherSuitesKubelet = "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_GCM_SHA256"
