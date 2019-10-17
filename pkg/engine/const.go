// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

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

const (
	// DefaultVNETCIDR is the default CIDR block for the VNET
	DefaultVNETCIDR = "10.0.0.0/8"
	// DefaultVNETCIDRIPv6 is the default IPv6 CIDR block for the VNET
	DefaultVNETCIDRIPv6 = "2001:1234:5678:9a00::/56"
	// DefaultInternalLbStaticIPOffset specifies the offset of the internal LoadBalancer's IP
	// address relative to the first consecutive Kubernetes static IP
	DefaultInternalLbStaticIPOffset = 10
	// NetworkPolicyNone is the string expression for the deprecated NetworkPolicy usage pattern "none"
	NetworkPolicyNone = "none"
	// NetworkPolicyCalico is the string expression for calico network policy config option
	NetworkPolicyCalico = "calico"
	// NetworkPolicyCilium is the string expression for cilium network policy config option
	NetworkPolicyCilium = "cilium"
	// NetworkPluginCilium is the string expression for cilium network plugin config option
	NetworkPluginCilium = NetworkPolicyCilium
	// NetworkPolicyAzure is the string expression for Azure CNI network policy manager
	NetworkPolicyAzure = "azure"
	// NetworkPluginAzure is the string expression for Azure CNI plugin
	NetworkPluginAzure = "azure"
	// NetworkPluginKubenet is the string expression for kubenet network plugin
	NetworkPluginKubenet = "kubenet"
	// NetworkPluginFlannel is the string expression for flannel network plugin
	NetworkPluginFlannel = "flannel"
	// KubeDNSAddonName is the name of the kube-dns-deployment addon
	KubeDNSAddonName = "kube-dns-deployment"
	// CoreDNSAddonName is the name of the coredns addon
	CoreDNSAddonName = "coredns"
	// DNSAutoscalerAddonName is the name of the coredns addon
	DNSAutoscalerAddonName = "dns-autoscaler"
	// KubeProxyAddonName is the name of the kube-proxy config addon
	KubeProxyAddonName = "kube-proxy-daemonset"
	// AzureStorageClassesAddonName is the name of the azure storage classes addon
	AzureStorageClassesAddonName = "azure-storage-classes"
	// AzureNetworkPolicyAddonName is the name of the azure npm daemon set addon
	AzureNetworkPolicyAddonName = "azure-npm-daemonset"
	// AzureVnetTelemetryAddonName is the name of the Azure vnet telemetry addon
	AzureVnetTelemetryAddonName = "azure-vnet-telemetry-daemonset"
	// CalicoAddonName is the name of calico daemonset addon
	CalicoAddonName = "calico-daemonset"
	// CiliumAddonName is the name of cilium daemonset addon
	CiliumAddonName = "cilium-daemonset"
	// FlannelAddonName is the name of flannel plugin daemonset addon
	FlannelAddonName = "flannel-daemonset"
	// AADAdminGroupAddonName is the name of the default admin group RBAC addon
	AADAdminGroupAddonName = "aad-default-admin-group-rbac"
	// AzureCloudProviderAddonName is the name of the azure cloud provider deployment addon
	AzureCloudProviderAddonName = "azure-cloud-provider-deployment"
	// AzureCNINetworkMonitorAddonName is the name of the azure cni network monitor addon
	AzureCNINetworkMonitorAddonName = "azure-cni-networkmonitor"
	// AuditPolicyAddonName is the name of the audit policy addon
	AuditPolicyAddonName = "audit-policy"
	// TillerAddonName is the name of the tiller addon deployment
	TillerAddonName = "tiller"
	// AADPodIdentityAddonName is the name of the aad-pod-identity addon deployment
	AADPodIdentityAddonName = "aad-pod-identity"
	// ACIConnectorAddonName is the name of the aci-connector addon deployment
	ACIConnectorAddonName = "aci-connector"
	// AppGwIngressAddonName appgw addon
	AppGwIngressAddonName = "appgw-ingress"
	// DashboardAddonName is the name of the kubernetes-dashboard addon deployment
	DashboardAddonName = "kubernetes-dashboard"
	// ClusterAutoscalerAddonName is the name of the autoscaler addon deployment
	ClusterAutoscalerAddonName = "cluster-autoscaler"
	// BlobfuseFlexVolumeAddonName is the name of the blobfuse flexvolume addon
	BlobfuseFlexVolumeAddonName = "blobfuse-flexvolume"
	// SMBFlexVolumeAddonName is the name of the smb flexvolume addon
	SMBFlexVolumeAddonName = "smb-flexvolume"
	// KeyVaultFlexVolumeAddonName is the name of the keyvault flexvolume addon deployment
	KeyVaultFlexVolumeAddonName = "keyvault-flexvolume"
	// ScheduledMaintenanceAddonName is the name of the scheduled maintenance addon deployment
	ScheduledMaintenanceAddonName = "scheduled-maintenance"
	// DefaultGeneratorCode specifies the source generator of the cluster template.
	DefaultGeneratorCode = "aksengine"
	// ReschedulerAddonName is the name of the rescheduler addon deployment
	ReschedulerAddonName = "rescheduler"
	// HeapsterAddonName is the name of the heapster addon deployment
	HeapsterAddonName = "heapster"
	// MetricsServerAddonName is the name of the kubernetes Metrics server addon deployment
	MetricsServerAddonName = "metrics-server"
	// NVIDIADevicePluginAddonName is the name of the kubernetes NVIDIA Device Plugin daemon set
	NVIDIADevicePluginAddonName = "nvidia-device-plugin"
	// ContainerMonitoringAddonName is the name of the kubernetes Container Monitoring addon deployment
	ContainerMonitoringAddonName = "container-monitoring"
	// AzureCNINetworkMonitoringAddonName is the name of the Azure CNI networkmonitor addon
	AzureCNINetworkMonitoringAddonName = "azure-cni-networkmonitor"
	// IPMASQAgentAddonName is the name of the ip masq agent addon
	IPMASQAgentAddonName = "ip-masq-agent"
	// PodSecurityPolicyAddonName is the name of the PodSecurityPolicy addon
	PodSecurityPolicyAddonName = "pod-security-policy"
	// DefaultKubernetesKubeletMaxPods is the max pods per kubelet
	DefaultKubernetesKubeletMaxPods = 110
	// DefaultMasterEtcdServerPort is the default etcd server port for Kubernetes master nodes
	DefaultMasterEtcdServerPort = 2380
	// DefaultMasterEtcdClientPort is the default etcd client port for Kubernetes master nodes
	DefaultMasterEtcdClientPort = 2379
	// etcdAccountNameFmt is the name format for a typical etcd account on Cosmos
	etcdAccountNameFmt = "%sk8s"
	// BasicLoadBalancerSku is the string const for Azure Basic Load Balancer
	BasicLoadBalancerSku = "Basic"
	// StandardLoadBalancerSku is the string const for Azure Standard Load Balancer
	StandardLoadBalancerSku = "Standard"
)

const (
	//DefaultExtensionsRootURL  Root URL for extensions
	DefaultExtensionsRootURL = "https://raw.githubusercontent.com/Azure/aks-engine/master/"
	// DefaultDockerEngineRepo for grabbing docker engine packages
	DefaultDockerEngineRepo = "https://download.docker.com/linux/ubuntu"
	// DefaultDockerComposeURL for grabbing docker images
	DefaultDockerComposeURL = "https://github.com/docker/compose/releases/download"
)

const (
	//DefaultConfigurationScriptRootURL  Root URL for configuration script (used for script extension on RHEL)
	DefaultConfigurationScriptRootURL = "https://raw.githubusercontent.com/Azure/aks-engine/master/parts/"
)

const (
	// AzureStackSuffix is appended to kubernetes version on Azure Stack instances
	AzureStackSuffix = "-azs"
	// AzureStackPrefix is appended to windows binary version for Azure Stack instances
	AzureStackPrefix = "azs-"
	// AzureStackCaCertLocation is where Azure Stack's CRP drops the stamp CA certificate
	AzureStackCaCertLocation = "/var/lib/waagent/Certificates.pem"
)

const (
	kubeConfigJSON = "k8s/kubeconfig.json"
	// Windows custom scripts
	kubernetesWindowsAgentCustomDataPS1   = "k8s/kuberneteswindowssetup.ps1"
	kubernetesWindowsAgentFunctionsPS1    = "k8s/kuberneteswindowsfunctions.ps1"
	kubernetesWindowsConfigFunctionsPS1   = "k8s/windowsconfigfunc.ps1"
	kubernetesWindowsKubeletFunctionsPS1  = "k8s/windowskubeletfunc.ps1"
	kubernetesWindowsCniFunctionsPS1      = "k8s/windowscnifunc.ps1"
	kubernetesWindowsAzureCniFunctionsPS1 = "k8s/windowsazurecnifunc.ps1"
	kubernetesWindowsOpenSSHFunctionPS1   = "k8s/windowsinstallopensshfunc.ps1"
)

// cloud-init (i.e. ARM customData) file references
const (
	kubernetesMasterNodeCustomDataYaml = "k8s/cloud-init/masternodecustomdata.yml"
	kubernetesNodeCustomDataYaml       = "k8s/cloud-init/nodecustomdata.yml"
	kubernetesJumpboxCustomDataYaml    = "k8s/cloud-init/jumpboxcustomdata.yml"
	kubernetesCSEMainScript            = "k8s/cloud-init/artifacts/cse_main.sh"
	kubernetesCSEHelpersScript         = "k8s/cloud-init/artifacts/cse_helpers.sh"
	kubernetesCSEInstall               = "k8s/cloud-init/artifacts/cse_install.sh"
	kubernetesCSEConfig                = "k8s/cloud-init/artifacts/cse_config.sh"
	kubernetesCISScript                = "k8s/cloud-init/artifacts/cis.sh"
	kubernetesCSECustomCloud           = "k8s/cloud-init/artifacts/cse_customcloud.sh"
	kubernetesHealthMonitorScript      = "k8s/cloud-init/artifacts/health-monitor.sh"
	// kubernetesKubeletMonitorSystemdTimer     = "k8s/cloud-init/artifacts/kubelet-monitor.timer" // TODO enable
	kubernetesKubeletMonitorSystemdService   = "k8s/cloud-init/artifacts/kubelet-monitor.service"
	kubernetesDockerMonitorSystemdTimer      = "k8s/cloud-init/artifacts/docker-monitor.timer"
	kubernetesDockerMonitorSystemdService    = "k8s/cloud-init/artifacts/docker-monitor.service"
	labelNodesScript                         = "k8s/cloud-init/artifacts/label-nodes.sh"
	labelNodesSystemdService                 = "k8s/cloud-init/artifacts/label-nodes.service"
	kubernetesMountEtcd                      = "k8s/cloud-init/artifacts/mountetcd.sh"
	kubernetesMasterGenerateProxyCertsScript = "k8s/cloud-init/artifacts/generateproxycerts.sh"
	kubernetesCustomSearchDomainsScript      = "k8s/cloud-init/artifacts/setup-custom-search-domains.sh"
	kubeletSystemdService                    = "k8s/cloud-init/artifacts/kubelet.service"
	kmsSystemdService                        = "k8s/cloud-init/artifacts/kms.service"
	aptPreferences                           = "k8s/cloud-init/artifacts/apt-preferences"
	dockerClearMountPropagationFlags         = "k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf"
	systemdBPFMount                          = "k8s/cloud-init/artifacts/sys-fs-bpf.mount"
	etcdSystemdService                       = "k8s/cloud-init/artifacts/etcd.service"
	auditdRules                              = "k8s/cloud-init/artifacts/auditd-rules"
	// scripts and service for enabling ipv6 dual stack
	dhcpv6SystemdService      = "k8s/cloud-init/artifacts/dhcpv6.service"
	dhcpv6ConfigurationScript = "k8s/cloud-init/artifacts/enable-dhcpv6.sh"
)

const (
	dcosCustomData188       = "dcos/dcoscustomdata188.t"
	dcosCustomData190       = "dcos/dcoscustomdata190.t"
	dcosCustomData198       = "dcos/dcoscustomdata198.t"
	dcosCustomData110       = "dcos/dcoscustomdata110.t"
	dcosProvision           = "dcos/dcosprovision.sh"
	dcosWindowsProvision    = "dcos/dcosWindowsProvision.ps1"
	dcosProvisionSource     = "dcos/dcosprovisionsource.sh"
	dcos2Provision          = "dcos/bstrap/dcosprovision.sh"
	dcos2BootstrapProvision = "dcos/bstrap/bootstrapprovision.sh"
	dcos2CustomData1110     = "dcos/bstrap/dcos1.11.0.customdata.t"
	dcos2CustomData1112     = "dcos/bstrap/dcos1.11.2.customdata.t"
)

const (
	swarmProvision            = "swarm/configure-swarm-cluster.sh"
	swarmWindowsProvision     = "swarm/Install-ContainerHost-And-Join-Swarm.ps1"
	swarmModeProvision        = "swarm/configure-swarmmode-cluster.sh"
	swarmModeWindowsProvision = "swarm/Join-SwarmMode-cluster.ps1"
)

const (
	agentOutputs                  = "agentoutputs.t"
	agentParams                   = "agentparams.t"
	armParameters                 = "k8s/armparameters.t"
	dcosAgentResourcesVMAS        = "dcos/dcosagentresourcesvmas.t"
	dcosWindowsAgentResourcesVMAS = "dcos/dcosWindowsAgentResourcesVmas.t"
	dcosAgentResourcesVMSS        = "dcos/dcosagentresourcesvmss.t"
	dcosWindowsAgentResourcesVMSS = "dcos/dcosWindowsAgentResourcesVmss.t"
	dcosAgentVars                 = "dcos/dcosagentvars.t"
	dcosBaseFile                  = "dcos/dcosbase.t"
	dcosParams                    = "dcos/dcosparams.t"
	dcosMasterResources           = "dcos/dcosmasterresources.t"
	dcosMasterVars                = "dcos/dcosmastervars.t"
	dcos2BaseFile                 = "dcos/bstrap/dcosbase.t"
	dcos2BootstrapVars            = "dcos/bstrap/bootstrapvars.t"
	dcos2BootstrapParams          = "dcos/bstrap/bootstrapparams.t"
	dcos2BootstrapResources       = "dcos/bstrap/bootstrapresources.t"
	dcos2BootstrapCustomdata      = "dcos/bstrap/bootstrapcustomdata.yml"
	dcos2MasterVars               = "dcos/bstrap/dcosmastervars.t"
	dcos2MasterResources          = "dcos/bstrap/dcosmasterresources.t"
	iaasOutputs                   = "iaasoutputs.t"
	kubernetesParams              = "k8s/kubernetesparams.t"
	masterOutputs                 = "masteroutputs.t"
	masterParams                  = "masterparams.t"
	swarmBaseFile                 = "swarm/swarmbase.t"
	swarmParams                   = "swarm/swarmparams.t"
	swarmAgentResourcesVMAS       = "swarm/swarmagentresourcesvmas.t"
	swarmAgentResourcesVMSS       = "swarm/swarmagentresourcesvmss.t"
	swarmAgentVars                = "swarm/swarmagentvars.t"
	swarmMasterResources          = "swarm/swarmmasterresources.t"
	swarmMasterVars               = "swarm/swarmmastervars.t"
	swarmWinAgentResourcesVMAS    = "swarm/swarmwinagentresourcesvmas.t"
	swarmWinAgentResourcesVMSS    = "swarm/swarmwinagentresourcesvmss.t"
	windowsParams                 = "windowsparams.t"
)
