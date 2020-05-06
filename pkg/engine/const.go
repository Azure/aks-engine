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
	// NetworkPolicyAntrea is the string expression for antrea network policy config option
	NetworkPolicyAntrea = "antrea"
	// NetworkPluginAntrea is the string expression for antrea network plugin config option
	NetworkPluginAntrea = NetworkPolicyAntrea
	// NetworkPolicyAzure is the string expression for Azure CNI network policy manager
	NetworkPolicyAzure = "azure"
	// NetworkPluginAzure is the string expression for Azure CNI plugin
	NetworkPluginAzure = "azure"
	// NetworkPluginKubenet is the string expression for kubenet network plugin
	NetworkPluginKubenet = "kubenet"
	// NetworkPluginFlannel is the string expression for flannel network plugin
	NetworkPluginFlannel = "flannel"
	// DefaultGeneratorCode specifies the source generator of the cluster template.
	DefaultGeneratorCode = "aksengine"
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
	kubeConfigJSON = "k8s/kubeconfig.json"
	// Windows custom scripts. These should all be listed in template_generator.go:func GetKubernetesWindowsAgentFunctions
	kubernetesWindowsAgentCustomDataPS1     = "k8s/kuberneteswindowssetup.ps1"
	kubernetesWindowsAgentFunctionsPS1      = "k8s/kuberneteswindowsfunctions.ps1"
	kubernetesWindowsConfigFunctionsPS1     = "k8s/windowsconfigfunc.ps1"
	kubernetesWindowsContainerdFunctionsPS1 = "k8s/windowscontainerdfunc.ps1"
	kubernetesWindowsCsiProxyFunctionsPS1   = "k8s/windowscsiproxyfunc.ps1"
	kubernetesWindowsKubeletFunctionsPS1    = "k8s/windowskubeletfunc.ps1"
	kubernetesWindowsCniFunctionsPS1        = "k8s/windowscnifunc.ps1"
	kubernetesWindowsAzureCniFunctionsPS1   = "k8s/windowsazurecnifunc.ps1"
	kubernetesWindowsOpenSSHFunctionPS1     = "k8s/windowsinstallopensshfunc.ps1"
	kubernetesWindowsLogsCleanupPS1         = "k8s/windowslogscleanup.ps1"
	kubernetesWindowsNodeResetPS1           = "k8s/windowsnodereset.ps1"
	kubeletStartPS1                         = "k8s/kubeletstart.ps1"
	kubeproxyStartPS1                       = "k8s/kubeproxystart.ps1"
)

// cloud-init (i.e. ARM customData) source file references
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
	untaintNodesScript                       = "k8s/cloud-init/artifacts/untaint-nodes.sh"
	untaintNodesSystemdService               = "k8s/cloud-init/artifacts/untaint-nodes.service"
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

// cloud-init destination file references
const (
	customCloudConfigCSEScriptFilepath   = "/opt/azure/containers/provision_configs_custom_cloud.sh"
	cseHelpersScriptFilepath             = "/opt/azure/containers/provision_source.sh"
	cseInstallScriptFilepath             = "/opt/azure/containers/provision_installs.sh"
	cseConfigScriptFilepath              = "/opt/azure/containers/provision_configs.sh"
	customSearchDomainsCSEScriptFilepath = "/opt/azure/containers/setup-custom-search-domains.sh"
	dhcpV6ServiceCSEScriptFilepath       = "/etc/systemd/system/dhcpv6.service"
	dhcpV6ConfigCSEScriptFilepath        = "/opt/azure/containers/enable-dhcpv6.sh"
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

// addons source and destination file references
const (
	heapsterAddonSourceFilename                    string = "kubernetesmasteraddons-heapster-deployment.yaml"
	heapsterAddonDestinationFilename               string = "kube-heapster-deployment.yaml"
	metricsServerAddonSourceFilename               string = "metrics-server.yaml"
	metricsServerAddonDestinationFilename          string = "metrics-server.yaml"
	tillerAddonSourceFilename                      string = "kubernetesmasteraddons-tiller-deployment.yaml"
	tillerAddonDestinationFilename                 string = "kube-tiller-deployment.yaml"
	aadPodIdentityAddonSourceFilename              string = "kubernetesmasteraddons-aad-pod-identity-deployment.yaml"
	aadPodIdentityAddonDestinationFilename         string = "aad-pod-identity-deployment.yaml"
	aciConnectorAddonSourceFilename                string = "kubernetesmasteraddons-aci-connector-deployment.yaml"
	aciConnectorAddonDestinationFilename           string = "aci-connector-deployment.yaml"
	azureDiskCSIAddonSourceFilename                string = "azuredisk-csi-driver-deployment.yaml"
	azureDiskCSIAddonDestinationFilename           string = "azuredisk-csi-driver-deployment.yaml"
	azureFileCSIAddonSourceFilename                string = "azurefile-csi-driver-deployment.yaml"
	azureFileCSIAddonDestinationFilename           string = "azurefile-csi-driver-deployment.yaml"
	clusterAutoscalerAddonSourceFilename           string = "cluster-autoscaler.yaml"
	clusterAutoscalerAddonDestinationFilename      string = "cluster-autoscaler.yaml"
	blobfuseFlexVolumeAddonSourceFilename          string = "kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml"
	blobfuseFlexVolumeAddonDestinationFilename     string = "blobfuse-flexvolume-installer.yaml"
	smbFlexVolumeAddonSourceFilename               string = "kubernetesmasteraddons-smb-flexvolume-installer.yaml"
	smbFlexVolumeAddonDestinationFilename          string = "smb-flexvolume-installer.yaml"
	keyvaultFlexVolumeAddonSourceFilename          string = "kubernetesmasteraddons-keyvault-flexvolume-installer.yaml"
	keyvaultFlexVolumeAddonDestinationFilename     string = "keyvault-flexvolume-installer.yaml"
	dashboardAddonSourceFilename                   string = "kubernetes-dashboard.yaml"
	dashboardAddonDestinationFilename              string = "kubernetes-dashboard.yaml"
	reschedulerAddonSourceFilename                 string = "kubernetesmasteraddons-kube-rescheduler-deployment.yaml"
	reschedulerAddonDestinationFilename            string = "kube-rescheduler-deployment.yaml"
	nvidiaAddonSourceFilename                      string = "kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml"
	nvidiaAddonDestinationFilename                 string = "nvidia-device-plugin.yaml"
	containerMonitoringAddonSourceFilename         string = "container-monitoring.yaml"
	containerMonitoringAddonDestinationFilename    string = "container-monitoring.yaml"
	ipMasqAgentAddonSourceFilename                 string = "ip-masq-agent.yaml"
	ipMasqAgentAddonDestinationFilename            string = "ip-masq-agent.yaml"
	azureCNINetworkMonitorAddonSourceFilename      string = "azure-cni-networkmonitor.yaml"
	azureCNINetworkMonitorAddonDestinationFilename string = "azure-cni-networkmonitor.yaml"
	calicoAddonSourceFilename                      string = "kubernetesmasteraddons-calico-daemonset.yaml"
	calicoAddonDestinationFilename                 string = "calico-daemonset.yaml"
	azureNetworkPolicyAddonSourceFilename          string = "kubernetesmasteraddons-azure-npm-daemonset.yaml"
	azureNetworkPolicyAddonDestinationFilename     string = "azure-npm-daemonset.yaml"
	azurePolicyAddonSourceFilename                 string = "azure-policy-deployment.yaml"
	azurePolicyAddonDestinationFilename            string = "azure-policy-deployment.yaml"
	cloudNodeManagerAddonSourceFilename            string = "kubernetesmasteraddons-cloud-node-manager.yaml"
	cloudNodeManagerAddonDestinationFilename       string = "cloud-node-manager.yaml"
	nodeProblemDetectorAddonSourceFilename         string = "node-problem-detector.yaml"
	nodeProblemDetectorAddonDestinationFilename    string = "node-problem-detector.yaml"
	kubeDNSAddonSourceFilename                     string = "kubernetesmasteraddons-kube-dns-deployment.yaml"
	kubeDNSAddonDestinationFilename                string = "kube-dns-deployment.yaml"
	corednsAddonSourceFilename                     string = "coredns.yaml"
	corednsAddonDestinationFilename                string = "coredns.yaml"
	kubeProxyAddonSourceFilename                   string = "kubernetesmasteraddons-kube-proxy-daemonset.yaml"
	kubeProxyAddonDestinationFilename              string = "kube-proxy-daemonset.yaml"
	podSecurityPolicyAddonSourceFilename           string = "kubernetesmasteraddons-pod-security-policy.yaml"
	podSecurityPolicyAddonDestinationFilename      string = "pod-security-policy.yaml"
	aadDefaultAdminGroupAddonSourceFilename        string = "aad-default-admin-group-rbac.yaml"
	aadDefaultAdminGroupDestinationFilename        string = "aad-default-admin-group-rbac.yaml"
	ciliumAddonSourceFilename                      string = "kubernetesmasteraddons-cilium-daemonset.yaml"
	ciliumAddonDestinationFilename                 string = "cilium-daemonset.yaml"
	antreaAddonSourceFilename                      string = "antrea.yaml"
	antreaAddonDestinationFilename                 string = "antrea.yaml"
	auditPolicyAddonSourceFilename                 string = "kubernetesmaster-audit-policy.yaml"
	auditPolicyAddonDestinationFilename            string = "audit-policy.yaml"
	cloudProviderAddonSourceFilename               string = "kubernetesmasteraddons-azure-cloud-provider-deployment.yaml"
	cloudProviderAddonDestinationFilename          string = "azure-cloud-provider-deployment.yaml"
	flannelAddonSourceFilename                     string = "kubernetesmasteraddons-flannel-daemonset.yaml"
	flannelAddonDestinationFilename                string = "flannel-daemonset.yaml"
	scheduledMaintenanceAddonSourceFilename        string = "scheduled-maintenance-deployment.yaml"
	scheduledMaintenanceAddonDestinationFilename   string = "scheduled-maintenance-deployment.yaml"
	secretsStoreCSIDriverAddonSourceFileName       string = "secrets-store-csi-driver.yaml"
	secretsStoreCSIDriverAddonDestinationFileName  string = "secrets-store-csi-driver.yaml"
)

// components source and destination file references
const (
	schedulerComponentSourceFilename                   string = "kubernetesmaster-kube-scheduler.yaml"
	schedulerComponentDestinationFilename              string = "kube-scheduler.yaml"
	controllerManagerComponentSourceFilename           string = "kubernetesmaster-kube-controller-manager.yaml"
	controllerManagerComponentDestinationFilename      string = "kube-controller-manager.yaml"
	cloudControllerManagerComponentSourceFilename      string = "kubernetesmaster-cloud-controller-manager.yaml"
	cloudControllerManagerComponentDestinationFilename string = "cloud-controller-manager.yaml"
	apiServerComponentSourceFilename                   string = "kubernetesmaster-kube-apiserver.yaml"
	apiServerComponentDestinationFilename              string = "kube-apiserver.yaml"
	addonManagerComponentSourceFilename                string = "kubernetesmaster-kube-addon-manager.yaml"
	addonManagerComponentDestinationFilename           string = "kube-addon-manager.yaml"
	clusterInitComponentDestinationFilename            string = "cluster-init.yaml"
)
