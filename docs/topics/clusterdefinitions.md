# Cluster Definitions



## Cluster Defintions for apiVersion "vlabs"

Here are the cluster definitions for apiVersion "vlabs":

### apiVersion

| Name       | Required | Description                                                   |
| ---------- | -------- | ------------------------------------------------------------- |
| apiVersion | yes      | The version of the template. For "vlabs" the value is "vlabs" |

### orchestratorProfile

`orchestratorProfile` describes the orchestrator settings.

| Name                | Required | Description                                        |
| ------------------- | -------- | -------------------------------------------------- |
| orchestratorType    | yes      | Specifies the orchestrator type for the cluster    |
| orchestratorRelease | no       | Specifies the orchestrator release for the cluster |
| orchestratorVersion | no       | Specifies the orchestrator version for the cluster |

Here are the valid values for the orchestrator types:

1. `Kubernetes` - this represents the Kubernetes orchestrator.

To learn more about supported versions, run the get-versions command:

```console
$ aks-engine get-versions
```

### kubernetesConfig

`kubernetesConfig` describes Kubernetes specific configuration.

| Name                             | Required                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| -------------------------------- | ------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| addons                           | no                        | Configure various Kubernetes addons configuration. See `addons` configuration [below](#addons)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| components                       | no                        | Configure core Kubernetes components. See `components` configuration [below](#components)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| apiServerConfig                  | no                        | Configure various runtime configuration for apiserver. See `apiServerConfig` [below](#feat-apiserver-config)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| cloudControllerManagerConfig     | no                        | Configure various runtime configuration for cloud-controller-manager. See `cloudControllerManagerConfig` [below](#feat-cloud-controller-manager-config)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| clusterSubnet                    | no                        | The IP subnet used for allocating IP addresses for pod network interfaces. The subnet must be in the VNET address space. With Azure CNI enabled, the default value is 10.240.0.0/12. Without Azure CNI, the default value is 10.244.0.0/16.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| containerRuntime                 | no                        | The container runtime to use as a backend. The default is `docker`. Also supported is `containerd`. Windows support for `containerd` is **Experimental** - see [Windows ContainerD](features.md#windows-containerd) |
| containerRuntimeConfig                | no                        | A map of key-value pairs to drive configuration of the container runtime. Currently accepts a single key, "dataDir", which configures the root data directory for the container runtime. dataDir must be an absolute path. This is only implemented on Linux. See an [example](../../examples/kubernetes-config/kubernetes-docker-tmpdir.json) which places docker on the tmp disk of a Linux VM. |
| controllerManagerConfig          | no                        | Configure various runtime configuration for controller-manager. See `controllerManagerConfig` [below](#feat-controller-manager-config)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| customWindowsPackageURL          | no                        | Configure custom windows Kubernetes release package URL for deployment on Windows. The format of this file is a zip file with multiple items (binaries, cni, infra container) in it. This setting will be deprecated in a future release of `aks-engine` where the binaries will be pulled in the format of Kubernetes releases that only contain the kubernetes binaries.                                                                                                                                                                                                                                                                                                                 |
| WindowsNodeBinariesURL           | no                        | Windows Kubernetes Node binaries can be provided in the format of Kubernetes release (example: https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.11.md#node-binaries-1). This setting allows overriding the binaries for custom builds.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| WindowsContainerdURL             | no (for development only) | **Experimental** - see [Windows ContainerD](features.md#windows-containerd) |
| WindowsSdnPluginURL              | no (for development only) | **Experimental** - see [Windows ContainerD](features.md#windows-containerd)  |
| dnsServiceIP                     | no                        | IP address for coredns or kube-dns to listen on. If specified must be in the range of `serviceCidr`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| mobyVersion                      | no (for development only) | Enables an explicit moby version, e.g. `3.0.3`. Default is `3.0.5`. This `kubernetesConfig` property is for development only, and applies only to cluster creation: `aks-engine upgrade` will always statically set `mobyVersion` to the default version at the time of upgrade, to ensure that upgraded clusters have the most recent, validated version of moby.                                                                                                                                                                                                                                                                                                                                                                        |
| containerdVersion                | no (for development only) | Enables an explicit containerd version, e.g. `1.1.4`. Default is `1.1.5`. This `kubernetesConfig` property is for development only, and applies only to cluster creation: `aks-engine upgrade` will always statically set `containerdVersion` to the default version at the time of upgrade, to ensure that upgraded clusters have the most recent, validated version of containerd. This value is currently ignored for Windows.                                                                                                                                                                                                                                                                                                                                                     |
| dockerBridgeSubnet               | no                        | The specific IP and subnet used for allocating IP addresses for the docker bridge network created on the kubernetes master and agents. Default value is 172.17.0.1/16. This value is used to configure the docker daemon using the [--bip flag](https://docs.docker.com/engine/userguide/networking/default_network/custom-docker0)                                                                                                                                                                                                                                                                                                                                                                                                       |
| enableAggregatedAPIs             | no                        | Enable [Kubernetes Aggregated APIs](https://kubernetes.io/docs/concepts/api-extension/apiserver-aggregation/). enableRbac must be set to true to use aggregated APIs. Aggregated API functionality is required by [Service Catalog](https://github.com/kubernetes-incubator/service-catalog/blob/master/README.md). (boolean - default is true)                                                                                                                                                                                                                                                                                                                                                                                                               |
| enableDataEncryptionAtRest       | no                        | Enable [kubernetes data encryption at rest](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/).This is currently an alpha feature. (boolean - default == false)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| enableEncryptionWithExternalKms  | no                        | Enable [kubernetes data encryption at rest with external KMS](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/).This is currently an alpha feature. (boolean - default == false)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| enablePodSecurityPolicy          | no                        | Deprecated, see the pod-security-policy addon for a description of the AKS Engine-configured PodSecurityPolicy spec that is bootstrapped as a Kubernetes addon                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| enableRbac                       | no                        | Enable [Kubernetes RBAC](https://kubernetes.io/docs/admin/authorization/rbac/) (boolean - default == true) RBAC support is required for Kubernetes 1.15.0 and greater, so enableRbac=false is not an allowed configuration for clusters >= 1.15.0. If you upgrade a cluster to 1.15.0 or greater from a version less than 1.15, and RBAC is disabled, the cluster configuration will be statically modified to enable RBAC as a result of running `aks-engine upgrade`.                                                                                                                                                                                                                                                                   |
| etcdDiskSizeGB                   | no                        | Size in GB to assign to etcd data volume. Defaults (if no user value provided) are: 256 GB for clusters up to 3 nodes; 512 GB for clusters with between 4 and 10 nodes; 1024 GB for clusters with between 11 and 20 nodes; and 2048 GB for clusters with more than 20 nodes                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| etcdEncryptionKey                | no                        | Enryption key to be used if enableDataEncryptionAtRest is enabled. Defaults to a random, generated, key                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| etcdVersion                      | no (for development only) | Enables an explicit etcd version, e.g. `3.2.23`. Default is `3.3.19`. This `kubernetesConfig` property is for development only, and recommended only for ephemeral clusters. However, you may use `aks-engine upgrade` on a cluster with an API model that includes a user-modified `etcdVersion` value. If `aks-engine upgrade` determines that the user-modified version is greater than the current AKS Engine default, `aks-engine upgrade` will _not_ replace the newer version with an older version. However, if `aks-engine upgrade` determines that the user-modified version is older than the current AKS Engine default, it will build the newly upgraded master node VMs with the newer, AKS Engine default version of etcd. |
| gcHighThreshold                  | no                        | Sets the --image-gc-high-threshold value on the kublet configuration. Default is 85. [See kubelet Garbage Collection](https://kubernetes.io/docs/concepts/cluster-administration/kubelet-garbage-collection/)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| gcLowThreshold                   | no                        | Sets the --image-gc-low-threshold value on the kublet configuration. Default is 80. [See kubelet Garbage Collection](https://kubernetes.io/docs/concepts/cluster-administration/kubelet-garbage-collection/)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| kubeletConfig                    | no                        | Configure various runtime configuration for kubelet. See `kubeletConfig` [below](#feat-kubelet-config)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| kubeReservedCgroup               | no                        | The name of a systemd slice to create for containment of both kubelet and the container runtime. When this value is a non-empty string, a file will be dropped at `/etc/systemd/system/$KUBE_RESERVED_CGROUP.slice` creating a systemd slice. Both kubelet and docker will run in this slice. This should not point to an existing systemd slice. If this value is unspecified or specified as the empty string, kubelet and the container runtime will run in the system slice by default.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| kubernetesImageBase              | no                        | Specifies the default image base URL (everything preceding the actual image filename) to be used for all kubernetes-related containers such as hyperkube, cloud-controller-manager, pause, addon-manager, heapster, exechealthz etc. e.g., `k8s.gcr.io/`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| loadBalancerSku                  | no                        | Sku of Load Balancer and Public IP. Candidate values are: `basic` and `standard`. If not set, it will be default to "standard". NOTE: Because VMs behind standard SKU load balancer will not be able to access the internet without an outbound rule configured with at least one frontend IP, AKS Engine creates a Load Balancer with an outbound rule and with agent nodes added to the backend pool during cluster creation, as described in the [Outbound NAT for internal Standard Load Balancer scenarios doc](https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-outbound-rules-overview#outbound-nat-for-internal-standard-load-balancer-scenarios)                                                                     |
| loadBalancerOutboundIPs                  | no                        | Number of outbound IP addresses (e.g., 3) to use in Standard LoadBalancer configuration. If not set, AKS Engine will configure a single outbound IP address. You may want more than one outbound IP address if you are running a large cluster that is processing lots of connections. See [here](https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-outbound-connections#multifesnat) for more documentation about how adding more outbound IP addresses can increase the number of SNAT ports available for use by the Standard Load Balancer in your cluster.                                                                     |
| networkPlugin                    | no                        | Specifies the network plugin implementation for the cluster. Valid values are:<br>`"azure"` (default), which provides an Azure native networking experience <br>`"kubenet"` for k8s software networking implementation. <br> `"flannel"` for using CoreOS Flannel <br> `"cilium"` for using the default Cilium CNI IPAM (requires the `"cilium"` networkPolicy as well)<br> `"antrea"` for using the Antrea network plugin (requires the `"antrea"` networkPolicy as well)                                                                                                                                                                                                                                                                                                                                                                                                                   |
| networkPolicy                    | no                        | Specifies the network policy enforcement tool for the cluster (currently Linux-only). Valid values are:<br>`"calico"` for Calico network policy.<br>`"cilium"` for cilium network policy (uses the `"cilium"` networkPlugin exclusively).<br> `"antrea"` for Antrea network policy (uses the `"antrea"` networkPlugin exclusively).<br> `"azure"` (experimental) for Azure CNI-compliant network policy (note: Azure CNI-compliant network policy requires explicit `"networkPlugin": "azure"` configuration as well).<br>See [network policy examples](../../examples/networkpolicy) for more information.                                                                                                                                                                                                                                                                   |
| privateCluster                   | no                        | Build a cluster without public addresses assigned. See `privateClusters` [below](#feat-private-cluster).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| schedulerConfig                  | no                        | Configure various runtime configuration for scheduler. See `schedulerConfig` [below](#feat-scheduler-config)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| serviceCidr                      | no                        | IP range for Service IPs, Default is "10.0.0.0/16". This range is never routed outside of a node so does not need to lie within clusterSubnet or the VNET                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| useInstanceMetadata              | no                        | Use the Azure cloudprovider instance metadata service for appropriate resource discovery operations. Default is `true`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| useManagedIdentity               | no                        | Includes and uses MSI identities for all interactions with the Azure Resource Manager (ARM) API. Instead of using a static service principal written to /etc/kubernetes/azure.json, Kubernetes will use a dynamic, time-limited token fetched from the MSI extension running on master and agent nodes. Defaults to true w/ user assigned identity when MasterProfile is using `VirtualMachineScaleSets`.                                                                                                                                                                      |
| userAssignedID               | no                        | When `useManagedIdentity` is set to true, this string value indicates that user assigned identity will be the type of managed identity used for cluster nodes, and appropriate pods. If the string value of `"userAssignedID"` is a fully qualified resource ID (e.g., `"/subscriptions/7a8f2518-7462-11ea-bc55-0242ac130003/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/my-user-assigned-identity"`), then the cluster will re-use that pre-existing user assigned managed identity resource; if the string value of `"userAssignedID"` is a simple string (e.g., `"my-new-user-assigned-identity"`), then a new user assigned managed identity resource will be created in the cluster resource group, with a name that matches that string value.                                                                                                                                                                        |
| azureCNIURLLinux                 | no                        | Deploy a private build of Azure CNI on Linux nodes. This should be a full path to the .tar.gz                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| azureCNIURLWindows               | no                        | Deploy a private build of Azure CNI on Windows nodes. This should be a full path to the .tar.gz                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| maximumLoadBalancerRuleCount     | no                        | Maximum allowed LoadBalancer Rule Count is the limit enforced by Azure Load balancer. Default is 250                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| kubeProxyMode                    | no                        | kube-proxy --proxy-mode value, either "iptables" or "ipvs". Default is "iptables". See https://kubernetes.io/blog/2018/07/09/ipvs-based-in-cluster-load-balancing-deep-dive/ for further reference.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| outboundRuleIdleTimeoutInMinutes | no                        | Specifies a value for IdleTimeoutInMinutes to control the outbound flow idle timeout of the agent standard loadbalancer. This value is set greater than the default Linux idle timeout (15.4 min): https://pracucci.com/linux-tcp-rto-min-max-and-tcp-retries2.html                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| cloudProviderBackoff | no                        | Use the Azure cloudprovider exponential backoff implementation when encountering retry-able errors from the Azure API. Defaults to `true` for Kubernetes v1.14.0 and greater.                                                                                                                                                    |
| cloudProviderBackoffMode | no                        | Which version of the Azure cloudprovider backoff implementation to use: the options are `"v1"` or `"v2"` (Kubernetes v1.14.0 or greater only). `"v2"` is a more recent backoff implementation which better honors Azure API HTTP headers to align backoff timings with the Azure API. Defaults to `"v2"` for Kubernetes v1.14.0 and greater, and `"v1"` for earlier versions of Kubernetes.                                                                                                                                                    |
| cloudProviderBackoffRetries | no                        | How many backoff retries before terminally failing the original Azure API operation. Defaults to `6`.                                                                                                                                                    |
| cloudProviderBackoffJitter | no                        | Only available in the `"v1"` cloudProviderBackoffMode: how much random variation to inject in retry timings, to better distribute stacked retries. Defaults to `1` when using cloudProviderBackoffMode `"v1"`, which effectively disables jitter.                                                                                                                                                     |
| cloudProviderBackoffDuration | no                        | The base duration, in seconds, in between retry attempts. Defaults to `5`.                                                                                                                                                     |
| cloudProviderBackoffExponent | no                        | Only available in the `"v1"` cloudProviderBackoffMode: the factor to multiply cloudProviderBackoffDuration by for each retry iteration. Defaults to `1.5` when using cloudProviderBackoffMode `"v1"`.                                                                                                                                                     |
| cloudProviderRateLimit | no                        | Use the Azure cloudprovider rate limiter to reduce the rate of calls to Azure APIs. Defaults to `true`.                                                                                                                                                      |
| cloudProviderRateLimitBucket | no                        | The size of the overflow read request queue when cloudProviderRateLimit is enabled. Defaults to the calculation "`100` * _number of agent pools in cluster configuration_.                                                                                                                                                      |
| cloudProviderRateLimitQPS | no                        | QPS for Azure cloudprovider read request rate limiter enforcement. Defaults to a minimum factor of `0.1` (i.e., 10%) of the cloudProviderRateLimitBucket (e.g., given a cloudProviderRateLimitBucket of 100, cloudProviderRateLimitQPS defaults to `10`), or `3`, whichever is greater.                                                                                                                                                 |
| cloudProviderRateLimitBucketWrite | no                        | The size of the overflow write request queue when cloudProviderRateLimit is enabled. Follows the same defaults calculation as cloudProviderRateLimitBucket.                                                                                                                                                      |
| cloudProviderRateLimitQPSWrite | no                        | QPS for Azure cloudprovider write request rate limiter enforcement. Follows the same defaults calculation as cloudProviderRateLimitQPS.                                                                                                                                                 |
| cloudProviderDisableOutboundSNAT | no                        | For clusters w/ Standard LB only: enforces the disabling of outbound NAT for that load balancing rule. See [here](https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-outbound-rules-overview#disablesnat) for more details. Defaults to `false`.                                                                                                                                               |

#### addons

`addons` is an interface to define user-configurable Kubernetes componentry. It is a child property of `kubernetesConfig`. Below is a list of currently available `addons`:

| Name of addon                                                                                             | Enabled by default?                                                                  | How many pods                   | Description                                                                                                                                                                                                                    |
| --------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ | ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| tiller                                                                                                    | false                                                                                | 1                               | Delivers the Helm server-side component: tiller. See https://github.com/kubernetes/helm for more info                                                                                                                          |
| kubernetes-dashboard                                                                                      | true                                                                                 | 1                               | Delivers the Kubernetes Dashboard component. See https://github.com/kubernetes/dashboard for more info                                                                                                                         |
| rescheduler                                                                                               | false                                                                                | 1                               | Delivers the Kubernetes rescheduler component                                                                                                                                                                                  |
| [cluster-autoscaler](../../examples/addons/cluster-autoscaler/README.md)                                  | false                                                                                | 1                               | Delivers the Kubernetes cluster autoscaler component. See https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/cloudprovider/azure for more info; only supported for VMSS clusters on the first agent pool. |
| [nvidia-device-plugin](../../examples/addons/nvidia-device-plugin/README.md)                              | true if using a Kubernetes cluster (v1.10+) with an N-series agent pool              | 1                               | Delivers the Kubernetes NVIDIA device plugin component. See https://github.com/NVIDIA/k8s-device-plugin for more info                                                                                                          |
| container-monitoring                                                                                      | false                                                                                | 1                               | Delivers the Kubernetes container monitoring component                                                                                                                                                                         |
| [blobfuse-flexvolume](https://github.com/Azure/kubernetes-volume-drivers/tree/master/flexvolume/blobfuse) | true                                                                                 | as many as linux agent nodes    | Access virtual filesystem backed by the Azure Blob storage                                                                                                                                                                     |
| [smb-flexvolume](https://github.com/Azure/kubernetes-volume-drivers/tree/master/flexvolume/smb)           | false                                                                                | as many as linux agent nodes    | Access SMB server by using CIFS/SMB protocol                                                                                                                                                                                   |
| [keyvault-flexvolume](../../examples/addons/keyvault-flexvolume/README.md)                                | false (true for <=1.15 clusters)                                                                                 | as many as linux agent nodes    | Access secrets, keys, and certs in Azure Key Vault from pods. This solution will be **DEPRECATED** in favor of csi-secrets-store addon                                                                                                                                                                  |
| [aad-pod-identity](../../examples/addons/aad-pod-identity/README.md)                                      | false                                                                                | 1 + 1 on each linux agent nodes | Assign Azure Active Directory Identities to Kubernetes applications                                                                                                                                                            |
| [scheduled-maintenance](https://github.com/awesomenix/drainsafe)                                          | false                                                                                | 1 + 1 on each linux agent nodes | Cordon and drain node during planned/unplanned [azure maintenance](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/scheduled-events)                                                                           |
| [azuredisk-csi-driver](https://github.com/kubernetes-sigs/azuredisk-csi-driver)                           | true if using a Kubernetes cluster (v1.13+) with `useCloudControllerManager` enabled | 1 + 1 on each linux agent nodes | Allows Kubernetes to use [Azure Disk](https://azure.microsoft.com/en-us/services/storage/disks/) volume                                                                                                                        |
| [azurefile-csi-driver](https://github.com/kubernetes-sigs/azurefile-csi-driver)                           | true if using a Kubernetes cluster (v1.13+) with `useCloudControllerManager` enabled | 1 + 1 on each linux agent nodes | Allows Kubernetes to use [Azure File](https://docs.microsoft.com/en-us/azure/storage/files/storage-files-introduction) volume                                                                                                  |
| [azure-policy](../../examples/addons/azure-policy/README.md)                                              | false                                                                                | 2                               | Open Policy Agent Gatekeeper with Azure Policy integration                                                                                                                                                                     |
| [node-problem-detector](../../examples/addons/node-problem-detector/README.md)                            | false                                                                                | as many as linux agent nodes    | Reports problems on Kubernetes nodes to kube-apiserver                                                                                                                                                                         |
| [kube-dns](https://github.com/kubernetes/kubernetes/tree/master/cluster/addons/dns/kube-dns)                            | false; if set to true, coredns must be set to false, i.e., only one cluster DNS addon may be used on a given cluster                                                                               | 1    | Cluster DNS services                                                                                                                                                                         |
| [coredns](https://github.com/coredns/deployment/tree/master/kubernetes)                            | true; if set to false, kube-dns must be set to true, i.e., you need at least one (and only one) of kube-dns or coredns enabled. For more configuration info, see `coredns` configuration [below](#coredns)                                                                                | 1    | Cluster DNS services                                                                                                                                                                         |
| [kube-proxy](https://kubernetes.io/docs/concepts/overview/components/#kube-proxy)                            | true                                                                                | 1    | a network proxy that runs on each node in your cluster                                                                                                                                                                         |
| [pod-security-policy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)                            | required for Kubernetes v1.15+; defaults to false for Kubernetes 1.14 and earlier                                                                                | 0    | a cluster-level resource that controls security-sensitive aspects of the pod specification                                                                                                                                                                         |
| [audit-policy](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#audit-policy)                            | true                                                                                | 0    | defines rules about what events should be recorded and what data they should include                                                                                                                                                                         |
| azure-cloud-provider                            | true                                                                                | 0    | Delivers required ClusterRole, ClusterRoleBinding, and StorageClass resources required for running the Azure cloudprovider runtime. May not be disabled.                                                                                                                                                                         |
| aad                            | true if adminGroupID is specified in the aadProfile configuration                                                                                | 0     | ClusterRoleBinding specification that adds an admin group matching the adminGroupID                                                                                                                                                                   |
| [cilium](https://docs.cilium.io/en/v1.4/kubernetes/policy/#ciliumnetworkpolicy)                            | true if networkPolicy is "cilium"; currently validated against Kubernetes v1.13, v1.14, and v1.15                                                                                | 0     | A NetworkPolicy CRD implementation by the Cilium project (currently supports v1.4)                                                                                                                                                                  |
| [flannel](https://coreos.com/flannel/docs/0.8.0/index.html)                            | false                                                                                | 0     | An addon that delivers flannel: a virtual network that gives a subnet to each host for use with container runtimes. If `networkPlugin` is set to `"flannel"` this addon will be enabled automatically. Not compatible with any other `networkPlugin` or `networkPolicy`.                                                                                                                                                                 |
| [csi-secrets-store](../../examples/addons/csi-secrets-store/README.md)                                | true (for 1.16+ clusters)                                                                                | as many as linux agent nodes    | Integrates secrets stores (Azure keyvault) via a [Container Storage Interface (CSI)](https://kubernetes-csi.github.io/docs/) volume.                                                                                                                                                                    |

To give a bit more info on the `addons` property: We've tried to expose the basic bits of data that allow useful configuration of these cluster features. Here are some example usage patterns that will unpack what `addons` provide:

To enable an addon (using "tiller" as an example):

```json
"kubernetesConfig": {
    "addons": [
        {
            "name": "tiller",
            "enabled" : true
        }
    ]
}
```

As you can see above, `addons` is an array child property of `kubernetesConfig`. Each addon that you want to add custom configuration to would be represented as an object item in the array. For example, to disable both tiller and dashboard:

```json
"kubernetesConfig": {
    "addons": [
        {
            "name": "tiller",
            "enabled" : false
        },
        {
            "name": "kubernetes-dashboard",
            "enabled" : false
        }
    ]
}
```

More usefully, let's add some custom configuration to the above addons:

```json
"kubernetesConfig": {
    "addons": [
        {
            "name": "tiller",
            "enabled": true,
            "containers": [
                {
                  "name": "tiller",
                  "image": "myDockerHubUser/tiller:v3.0.0-alpha",
                  "cpuRequests": "1",
                  "memoryRequests": "1024Mi",
                  "cpuLimits": "1",
                  "memoryLimits": "1024Mi"
                }
              ]
        },
        {
            "name": "kubernetes-dashboard",
            "enabled": true,
            "containers": [
                {
                  "name": "kubernetes-dashboard",
                  "cpuRequests": "50m",
                  "memoryRequests": "512Mi",
                  "cpuLimits": "50m",
                  "memoryLimits": "512Mi"
                },
                {
                  "name": "kubernetes-dashboard-metrics-scraper",
                  "cpuRequests": "50m",
                  "memoryRequests": "512Mi",
                  "cpuLimits": "50m",
                  "memoryLimits": "512Mi"
                }
              ]
        },
        {
            "name": "cluster-autoscaler",
            "enabled": true,
            "containers": [
              {
                "name": "cluster-autoscaler",
                "cpuRequests": "100m",
                "memoryRequests": "300Mi",
                "cpuLimits": "100m",
                "memoryLimits": "300Mi"
              }
            ],
            "config": {
              "max-nodes": "5",
              "min-nodes": "1",
              "scan-interval": "10s"
            }
        }
    ]
}
```

Above you see custom configuration for both tiller and kubernetes-dashboard. Both include specific resource limit values across the following dimensions:

- cpuRequests
- memoryRequests
- cpuLimits
- memoryLimits

See https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ for more on Kubernetes resource limits.

Additionally above, we specified a custom docker image for tiller, let's say we want to build a cluster and test an alpha version of tiller in it. **Important note!** customizing the image is not sticky across upgrade/scale, to ensure that AKS Engine always delivers a version-curated, known-working addon when moving a cluster to a new version. Considering all that, providing a custom image reference for an addon configuration should be considered for testing/development, but not for a production cluster. If you'd like to entirely customize one of the addons available, including across scale/upgrade operations, you may include in an addon's spec a base64-encoded string of a Kubernetes yaml manifest. E.g.,

```json
"kubernetesConfig": {
    "addons": [
        {
            "name": "kube-proxy",
            "enabled": true,
            "data": "IGFwaVZlcnNpb246IHYxCmtpbmQ6IFNlcnZpY2UKbWV0YWRhdGE6CiAgbmFtZToga3ViZS1kbnMKICBuYW1lc3BhY2U6IGt1YmUtc3lzdGVtCiAgbGFiZWxzOgogICAgazhzLWFwcDoga3ViZS1kbnMKICAgIGt1YmVybmV0ZXMuaW8vY2x1c3Rlci1zZXJ2aWNlOiAidHJ1ZSIKICAgIGFkZG9ubWFuYWdlci5rdWJlcm5ldGVzLmlvL21vZGU6IFJlY29uY2lsZQogICAga3ViZXJuZXRlcy5pby9uYW1lOiAiS3ViZUROUyIKc3BlYzoKICBzZWxlY3RvcjoKICAgIGs4cy1hcHA6IGt1YmUtZG5zCiAgY2x1c3RlcklQOiA8a3ViZUROU1NlcnZpY2VJUD4KICBwb3J0czoKICAtIG5hbWU6IGRucwogICAgcG9ydDogNTMKICAgIHByb3RvY29sOiBVRFAKICAtIG5hbWU6IGRucy10Y3AKICAgIHBvcnQ6IDUzCiAgICBwcm90b2NvbDogVENQCi0tLQphcGlWZXJzaW9uOiB2MQpraW5kOiBTZXJ2aWNlQWNjb3VudAptZXRhZGF0YToKICBuYW1lOiBrdWJlLWRucwogIG5hbWVzcGFjZToga3ViZS1zeXN0ZW0KICBsYWJlbHM6CiAgICBrdWJlcm5ldGVzLmlvL2NsdXN0ZXItc2VydmljZTogInRydWUiCiAgICBhZGRvbm1hbmFnZXIua3ViZXJuZXRlcy5pby9tb2RlOiBSZWNvbmNpbGUKLS0tCmFwaVZlcnNpb246IHYxCmtpbmQ6IENvbmZpZ01hcAptZXRhZGF0YToKICBuYW1lOiBrdWJlLWRucwogIG5hbWVzcGFjZToga3ViZS1zeXN0ZW0KICBsYWJlbHM6CiAgICBhZGRvbm1hbmFnZXIua3ViZXJuZXRlcy5pby9tb2RlOiBFbnN1cmVFeGlzdHMKLS0tCmFwaVZlcnNpb246IGV4dGVuc2lvbnMvdjFiZXRhMQpraW5kOiBEZXBsb3ltZW50Cm1ldGFkYXRhOgogIG5hbWU6IGt1YmUtZG5zCiAgbmFtZXNwYWNlOiBrdWJlLXN5c3RlbQogIGxhYmVsczoKICAgIGs4cy1hcHA6IGt1YmUtZG5zCiAgICBrdWJlcm5ldGVzLmlvL2NsdXN0ZXItc2VydmljZTogInRydWUiCiAgICBhZGRvbm1hbmFnZXIua3ViZXJuZXRlcy5pby9tb2RlOiBSZWNvbmNpbGUKc3BlYzoKICAjIHJlcGxpY2FzOiBub3Qgc3BlY2lmaWVkIGhlcmU6CiAgIyAxLiBJbiBvcmRlciB0byBtYWtlIEFkZG9uIE1hbmFnZXIgZG8gbm90IHJlY29uY2lsZSB0aGlzIHJlcGxpY2FzIHBhcmFtZXRlci4KICAjIDIuIERlZmF1bHQgaXMgMS4KICAjIDMuIFdpbGwgYmUgdHVuZWQgaW4gcmVhbCB0aW1lIGlmIEROUyBob3Jpem9udGFsIGF1dG8tc2NhbGluZyBpcyB0dXJuZWQgb24uCiAgc3RyYXRlZ3k6CiAgICByb2xsaW5nVXBkYXRlOgogICAgICBtYXhTdXJnZTogMTAlCiAgICAgIG1heFVuYXZhaWxhYmxlOiAwCiAgc2VsZWN0b3I6CiAgICBtYXRjaExhYmVsczoKICAgICAgazhzLWFwcDoga3ViZS1kbnMKICB0ZW1wbGF0ZToKICAgIG1ldGFkYXRhOgogICAgICBsYWJlbHM6CiAgICAgICAgazhzLWFwcDoga3ViZS1kbnMKICAgICAgYW5ub3RhdGlvbnM6CiAgICAgICAgc2NoZWR1bGVyLmFscGhhLmt1YmVybmV0ZXMuaW8vY3JpdGljYWwtcG9kOiAnJwogICAgICAgIHNlY2NvbXAuc2VjdXJpdHkuYWxwaGEua3ViZXJuZXRlcy5pby9wb2Q6ICdkb2NrZXIvZGVmYXVsdCcKICAgIHNwZWM6CiAgICAgIHByaW9yaXR5Q2xhc3NOYW1lOiBzeXN0ZW0tY2x1c3Rlci1jcml0aWNhbAogICAgICB0b2xlcmF0aW9uczoKICAgICAgLSBrZXk6ICJDcml0aWNhbEFkZG9uc09ubHkiCiAgICAgICAgb3BlcmF0b3I6ICJFeGlzdHMiCiAgICAgIHZvbHVtZXM6CiAgICAgIC0gbmFtZToga3ViZS1kbnMtY29uZmlnCiAgICAgICAgY29uZmlnTWFwOgogICAgICAgICAgbmFtZToga3ViZS1kbnMKICAgICAgICAgIG9wdGlvbmFsOiB0cnVlCiAgICAgIGNvbnRhaW5lcnM6CiAgICAgIC0gbmFtZToga3ViZWRucwogICAgICAgIGltYWdlOiA8a3ViZXJuZXRlc0t1YmVETlNTcGVjPgogICAgICAgIHJlc291cmNlczoKICAgICAgICAgICMgVE9ETzogU2V0IG1lbW9yeSBsaW1pdHMgd2hlbiB3ZSd2ZSBwcm9maWxlZCB0aGUgY29udGFpbmVyIGZvciBsYXJnZQogICAgICAgICAgIyBjbHVzdGVycywgdGhlbiBzZXQgcmVxdWVzdCA9IGxpbWl0IHRvIGtlZXAgdGhpcyBjb250YWluZXIgaW4KICAgICAgICAgICMgZ3VhcmFudGVlZCBjbGFzcy4gQ3VycmVudGx5LCB0aGlzIGNvbnRhaW5lciBmYWxscyBpbnRvIHRoZQogICAgICAgICAgIyAiYnVyc3RhYmxlIiBjYXRlZ29yeSBzbyB0aGUga3ViZWxldCBkb2Vzbid0IGJhY2tvZmYgZnJvbSByZXN0YXJ0aW5nIGl0LgogICAgICAgICAgbGltaXRzOgogICAgICAgICAgICBtZW1vcnk6IDE3ME1pCiAgICAgICAgICByZXF1ZXN0czoKICAgICAgICAgICAgY3B1OiAxMDBtCiAgICAgICAgICAgIG1lbW9yeTogNzBNaQogICAgICAgIGxpdmVuZXNzUHJvYmU6CiAgICAgICAgICBodHRwR2V0OgogICAgICAgICAgICBwYXRoOiAvaGVhbHRoY2hlY2sva3ViZWRucwogICAgICAgICAgICBwb3J0OiAxMDA1NAogICAgICAgICAgICBzY2hlbWU6IEhUVFAKICAgICAgICAgIGluaXRpYWxEZWxheVNlY29uZHM6IDYwCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgICAgc3VjY2Vzc1RocmVzaG9sZDogMQogICAgICAgICAgZmFpbHVyZVRocmVzaG9sZDogNQogICAgICAgIHJlYWRpbmVzc1Byb2JlOgogICAgICAgICAgaHR0cEdldDoKICAgICAgICAgICAgcGF0aDogL3JlYWRpbmVzcwogICAgICAgICAgICBwb3J0OiA4MDgxCiAgICAgICAgICAgIHNjaGVtZTogSFRUUAogICAgICAgICAgIyB3ZSBwb2xsIG9uIHBvZCBzdGFydHVwIGZvciB0aGUgS3ViZXJuZXRlcyBtYXN0ZXIgc2VydmljZSBhbmQKICAgICAgICAgICMgb25seSBzZXR1cCB0aGUgL3JlYWRpbmVzcyBIVFRQIHNlcnZlciBvbmNlIHRoYXQncyBhdmFpbGFibGUuCiAgICAgICAgICBpbml0aWFsRGVsYXlTZWNvbmRzOiAzCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgIGFyZ3M6CiAgICAgICAgLSAtLWRvbWFpbj08a3ViZXJuZXRlc0t1YmVsZXRDbHVzdGVyRG9tYWluPi4KICAgICAgICAtIC0tZG5zLXBvcnQ9MTAwNTMKICAgICAgICAtIC0tY29uZmlnLWRpcj0va3ViZS1kbnMtY29uZmlnCiAgICAgICAgLSAtLXY9MgogICAgICAgIGVudjoKICAgICAgICAtIG5hbWU6IFBST01FVEhFVVNfUE9SVAogICAgICAgICAgdmFsdWU6ICIxMDA1NSIKICAgICAgICBwb3J0czoKICAgICAgICAtIGNvbnRhaW5lclBvcnQ6IDEwMDUzCiAgICAgICAgICBuYW1lOiBkbnMtbG9jYWwKICAgICAgICAgIHByb3RvY29sOiBVRFAKICAgICAgICAtIGNvbnRhaW5lclBvcnQ6IDEwMDUzCiAgICAgICAgICBuYW1lOiBkbnMtdGNwLWxvY2FsCiAgICAgICAgICBwcm90b2NvbDogVENQCiAgICAgICAgLSBjb250YWluZXJQb3J0OiAxMDA1NQogICAgICAgICAgbmFtZTogbWV0cmljcwogICAgICAgICAgcHJvdG9jb2w6IFRDUAogICAgICAgIHZvbHVtZU1vdW50czoKICAgICAgICAtIG5hbWU6IGt1YmUtZG5zLWNvbmZpZwogICAgICAgICAgbW91bnRQYXRoOiAva3ViZS1kbnMtY29uZmlnCiAgICAgIC0gbmFtZTogZG5zbWFzcQogICAgICAgIGltYWdlOiA8a3ViZXJuZXRlc0ROU01hc3FTcGVjPgogICAgICAgIGxpdmVuZXNzUHJvYmU6CiAgICAgICAgICBodHRwR2V0OgogICAgICAgICAgICBwYXRoOiAvaGVhbHRoY2hlY2svZG5zbWFzcQogICAgICAgICAgICBwb3J0OiAxMDA1NAogICAgICAgICAgICBzY2hlbWU6IEhUVFAKICAgICAgICAgIGluaXRpYWxEZWxheVNlY29uZHM6IDYwCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgICAgc3VjY2Vzc1RocmVzaG9sZDogMQogICAgICAgICAgZmFpbHVyZVRocmVzaG9sZDogNQogICAgICAgIGFyZ3M6CiAgICAgICAgLSAtdj0yCiAgICAgICAgLSAtbG9ndG9zdGRlcnIKICAgICAgICAtIC1jb25maWdEaXI9L2V0Yy9rOHMvZG5zL2Ruc21hc3EtbmFubnkKICAgICAgICAtIC1yZXN0YXJ0RG5zbWFzcT10cnVlCiAgICAgICAgLSAtLQogICAgICAgIC0gLWsKICAgICAgICAtIC0tY2FjaGUtc2l6ZT0xMDAwCiAgICAgICAgLSAtLW5vLW5lZ2NhY2hlCiAgICAgICAgLSAtLWxvZy1mYWNpbGl0eT0tCiAgICAgICAgLSAtLXNlcnZlcj0vY2x1c3Rlci5sb2NhbC8xMjcuMC4wLjEjMTAwNTMKICAgICAgICAtIC0tc2VydmVyPS9pbi1hZGRyLmFycGEvMTI3LjAuMC4xIzEwMDUzCiAgICAgICAgLSAtLXNlcnZlcj0vaXA2LmFycGEvMTI3LjAuMC4xIzEwMDUzCiAgICAgICAgcG9ydHM6CiAgICAgICAgLSBjb250YWluZXJQb3J0OiA1MwogICAgICAgICAgbmFtZTogZG5zCiAgICAgICAgICBwcm90b2NvbDogVURQCiAgICAgICAgLSBjb250YWluZXJQb3J0OiA1MwogICAgICAgICAgbmFtZTogZG5zLXRjcAogICAgICAgICAgcHJvdG9jb2w6IFRDUAogICAgICAgICMgc2VlOiBodHRwczovL2dpdGh1Yi5jb20va3ViZXJuZXRlcy9rdWJlcm5ldGVzL2lzc3Vlcy8yOTA1NSBmb3IgZGV0YWlscwogICAgICAgIHJlc291cmNlczoKICAgICAgICAgIHJlcXVlc3RzOgogICAgICAgICAgICBjcHU6IDE1MG0KICAgICAgICAgICAgbWVtb3J5OiAyME1pCiAgICAgICAgdm9sdW1lTW91bnRzOgogICAgICAgIC0gbmFtZToga3ViZS1kbnMtY29uZmlnCiAgICAgICAgICBtb3VudFBhdGg6IC9ldGMvazhzL2Rucy9kbnNtYXNxLW5hbm55CiAgICAgIC0gbmFtZTogc2lkZWNhcgogICAgICAgIGltYWdlOiA8a3ViZXJuZXRlc0ROU1NpZGVjYXJTcGVjPgogICAgICAgIGxpdmVuZXNzUHJvYmU6CiAgICAgICAgICBodHRwR2V0OgogICAgICAgICAgICBwYXRoOiAvbWV0cmljcwogICAgICAgICAgICBwb3J0OiAxMDA1NAogICAgICAgICAgICBzY2hlbWU6IEhUVFAKICAgICAgICAgIGluaXRpYWxEZWxheVNlY29uZHM6IDYwCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgICAgc3VjY2Vzc1RocmVzaG9sZDogMQogICAgICAgICAgZmFpbHVyZVRocmVzaG9sZDogNQogICAgICAgIGFyZ3M6CiAgICAgICAgLSAtLXY9MgogICAgICAgIC0gLS1sb2d0b3N0ZGVycgogICAgICAgIC0gLS1wcm9iZT1rdWJlZG5zLDEyNy4wLjAuMToxMDA1MyxrdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLjxrdWJlcm5ldGVzS3ViZWxldENsdXN0ZXJEb21haW4+LDUsU1JWCiAgICAgICAgLSAtLXByb2JlPWRuc21hc3EsMTI3LjAuMC4xOjUzLGt1YmVybmV0ZXMuZGVmYXVsdC5zdmMuPGt1YmVybmV0ZXNLdWJlbGV0Q2x1c3RlckRvbWFpbj4sNSxTUlYKICAgICAgICBwb3J0czoKICAgICAgICAtIGNvbnRhaW5lclBvcnQ6IDEwMDU0CiAgICAgICAgICBuYW1lOiBtZXRyaWNzCiAgICAgICAgICBwcm90b2NvbDogVENQCiAgICAgICAgcmVzb3VyY2VzOgogICAgICAgICAgcmVxdWVzdHM6CiAgICAgICAgICAgIG1lbW9yeTogMjBNaQogICAgICAgICAgICBjcHU6IDEwbQogICAgICBkbnNQb2xpY3k6IERlZmF1bHQgICMgRG9uJ3QgdXNlIGNsdXN0ZXIgRE5TLgogICAgICBzZXJ2aWNlQWNjb3VudE5hbWU6IGt1YmUtZG5zCiAgICAgIG5vZGVTZWxlY3RvcjoKICAgICAgICBiZXRhLmt1YmVybmV0ZXMuaW8vb3M6IGxpbnV4"
        }
    ]
}
```

The reason for the unsightly base64-encoded input type is to optimize delivery payload, and to squash a human-maintainable yaml file representation into something that can be tightly pasted into a JSON string value without the arguably more unsightly carriage returns / whitespace that would be delivered with a literal copy/paste of a Kubernetes manifest.

#### coredns

The `coredns` addon includes integration with the `cluster-proportional-autoscaler` project to automatically scale out coredns pod replicas according to node, or core count. More information at the official docs [here](https://kubernetes.io/docs/tasks/administer-cluster/dns-horizontal-autoscaling/). The AKS Engine default configuration tunes the autoscaler thresholds to "32" nodes, and "512" cores (whichever threshold is crossed first engages scaling behaviors), with a minimum replica count of "1". The scale thresholds are higher than those seen in example docs due to observed (not catastrophic) increases in per-DNS resolution response times. In other words, for smaller clusters that aren't coredns pod-constrained, a single coredns pod is more responsive. These configurations are entirely user-configurable: you may tune them according to the operational DNS characteristics of your environment. E.g.:

```
"kubernetesConfig": {
    "addons": [
        ...
        {
          "name": "coredns",
          "enabled": true,
          "config": {
            "cores-per-replica": "512",
            "min-replicas": "3",
            "nodes-per-replica": "32"
          }
        },
        ...
    ]
}
```

The above example configuration would ship a coredns configuration that has at least 3 pod replicas at all times, and then scales out to 4 when the 97th node, or the 1537th core (whichever comes frirst) is observed running in the cluster (and so on and so on).

#### components

`components` is an interface to allow for user-configurable core Kubernetes component implementations. Normally, you won't need to modify this configuration, as AKS Engine will use the best, known-working component implementations validated against Azure for all supported versions of Kubernetes. To support the rapid development of Azure + Kubernetes (e.g., Azure cloudprovider), this configuration vector may be useful for validating a custom build or configuration of the various core components on a running Azure Kubernetes cluster. Again, as with addons, this configurable vector is designed for *cluster creation only*. Using `aks-engine upgrade` on a cluster will override the original, user-configured settings during the upgrade operation, rendering an upgraded cluster with the AKS Engine defaults for `kube-controller-manager`, `cloud-controller-manager`, `kube-apiserver`, `kube-scheduler`, and `kube-addon-manager`.

`components` is a child property of `kubernetesConfig`. Below is a list of `components` you may provide.:

| Name of addon                                                                                             | Enabled by default?                                                                  | How many pods                   | Description                                                                                                                                                                                                                    |
| --------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ | ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| kube-controller-manager | true | 1 per master node | The Kubernetes controller manager is a daemon that embeds the core control loops shipped with Kubernetes. Official docs [here](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/). |
| cloud-controller-manager | false | 1 per master node | The Cloud controller manager is a daemon that embeds the cloud specific control loops shipped with Kubernetes. Official docs [here](https://kubernetes.io/docs/reference/command-line-tools-reference/cloud-controller-manager/). |
| kube-apiserver | true | 1 per master node | The Kubernetes API server validates and configures data for the api objects which include pods, services, replicationcontrollers, and others. The API Server services REST operations and provides the frontend to the cluster's shared state through which all other components interact. Official docs [here](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/). |
| kube-scheduler | true | 1 per master node | The Kubernetes scheduler is a policy-rich, topology-aware, workload-specific function that significantly impacts availability, performance, and capacity. The scheduler needs to take into account individual and collective resource requirements, quality of service requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, data locality, inter-workload interference, deadlines, and so on. Official docs [here](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/). |
| kube-addon-manager | true | 1 per master node | Addon manager provides a standard way to deliver additional Kubernetes componentry. The addons supported by AKS Engine (documented [above](#addons)) are installed into the cluster using Addon manager. Official docs [here](https://github.com/kubernetes/kubernetes/tree/master/cluster/addons/addon-manager). |
| cluster-init | false | n/a | This is an interface to deliver Kubernetes resource configuration that may be tightly coupled to an ARM deployment. For example, if you use AKS Engine in an automated pipeline and you can scale up a cluster quickly by loading a Kubernetes specification immediately after bootstrapping the control plane, give the base64-encoded YAML representation of that spec to `cluster-init`. That data is decoded at cluster creation time to the path `/opt/azure/containers/cluster-init.yaml` on the first master VM, and then loaded into the cluster via `kubectl apply -f`. See[`kubectl apply`](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#apply) for more documentation on how to specify the source YAML. Note: the `cluster-init` spec is ignored during `aks-engine upgrade` and `aks-engine scale`operations. |

To give a bit more info on the `components` property: Currently, there are two configuration vectors available for use. Firstly, you have the option to pass in a custom image reference for the container that the AKS Engine-provided specs implement (all components are implemented as single-container Pod resources); in addition you may provide an alternate command string to execute inside the container. E.g.:

```json
"kubernetesConfig": {
    "components": [
        {
            "name": "kube-controller-manager",
            "enabled" : true,
            "containers": [
                {
                    "name": "kube-controller-manager",
                    "image": "myDockerHubUser/custom-controller-manager-build:v1.17.2-dirty",
                }
            ],
            "config": {
                "command": "kube-controller-manager-custom"
            }
        }
    ]
}
```

In addition, a component may be disabled by setting `"enabled": false`. This only makes practical sense for `cloud-controller-manager`, which is an optional Kubernetes control plane implementation that runs the Azure-specific control plane runtime in a separate pod from `kube-controller-manager`. For all other components, if they are disabled, then AKS Engine will not by itself create a functional Kubernetes cluster. This configuration is made available for development purposes only, and is not recommended for users building functional Kubernetes clusters using AKS Engine.

Note: `kube-addon-manager` does not support a command string, as that configuration isn't appropriate for Addon Manager, which only defines a container image as configurable input.

For each Kubernetes component (with the exception of `kube-addon-manager`), the args to the command are provided via the equivalent "`<component>`Config" property. E.g.:

- `kube-controller-manager`
  - Command arguments map to [controllerManagerConfig](#feat-controller-manager-config) key=val pairs
- `cloud-controller-manager`
  - Command arguments map to [cloudControllerManagerConfig](#feat-cloud-controller-manager-config) key=val pairs
- `kube-apiserver`
  - Command arguments map to [apiServerConfig](#feat-apiserver-config) key=val pairs
- `kube-scheduler`
  - Command arguments map to [schedulerConfig](#feat-scheduler-config) key=val pairs

See https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ for more on Kubernetes resource limits.

As with addons, you may include an entirely custom component spec as a base64-encoded string of a Kubernetes yaml manifest. E.g.,

```json
"kubernetesConfig": {
    "components": [
        {
            "name": "kube-controller-manager",
            "enabled": true,
            "data": "YXBpVmVyc2lvbjogdjEKa2luZDogUG9kCm1ldGFkYXRhOgogIG5hbWU6IGt1YmUtY29udHJvbGxl
ci1tYW5hZ2VyCiAgbmFtZXNwYWNlOiBrdWJlLXN5c3RlbQogIGxhYmVsczoKICAgIHRpZXI6IGNv
bnRyb2wtcGxhbmUKICAgIGNvbXBvbmVudDoga3ViZS1jb250cm9sbGVyLW1hbmFnZXIKc3BlYzoK
ICBwcmlvcml0eUNsYXNzTmFtZTogc3lzdGVtLW5vZGUtY3JpdGljYWwKICBob3N0TmV0d29yazog
dHJ1ZQogIGNvbnRhaW5lcnM6CiAgICAtIG5hbWU6IGt1YmUtY29udHJvbGxlci1tYW5hZ2VyCiAg
ICAgIGltYWdlOiBrOHMuZ2NyLmlvL2h5cGVya3ViZS1hbWQ2NDp2MS4xNi42CiAgICAgIGltYWdl
UHVsbFBvbGljeTogSWZOb3RQcmVzZW50CiAgICAgIGNvbW1hbmQ6IFsiL2h5cGVya3ViZSIsICJr
dWJlLWNvbnRyb2xsZXItbWFuYWdlciJdCiAgICAgIGFyZ3M6IFsiLS1hbGxvY2F0ZS1ub2RlLWNp
ZHJzPWZhbHNlIiwgIi0tY2xvdWQtY29uZmlnPS9ldGMva3ViZXJuZXRlcy9henVyZS5qc29uIiwg
Ii0tY2xvdWQtcHJvdmlkZXI9YXp1cmUiLCAiLS1jbHVzdGVyLWNpZHI9MTAuMjQwLjAuMC8xMiIs
ICItLWNsdXN0ZXItbmFtZT1rdWJlcm5ldGVzLXdlc3R1czItMzMyODEiLCAiLS1jbHVzdGVyLXNp
Z25pbmctY2VydC1maWxlPS9ldGMva3ViZXJuZXRlcy9jZXJ0cy9jYS5jcnQiLCAiLS1jbHVzdGVy
LXNpZ25pbmcta2V5LWZpbGU9L2V0Yy9rdWJlcm5ldGVzL2NlcnRzL2NhLmtleSIsICItLWNvbmZp
Z3VyZS1jbG91ZC1yb3V0ZXM9ZmFsc2UiLCAiLS1jb250cm9sbGVycz0qLGJvb3RzdHJhcHNpZ25l
cix0b2tlbmNsZWFuZXIiLCAiLS1mZWF0dXJlLWdhdGVzPUxvY2FsU3RvcmFnZUNhcGFjaXR5SXNv
bGF0aW9uPXRydWUsU2VydmljZU5vZGVFeGNsdXNpb249dHJ1ZSIsICItLWhvcml6b250YWwtcG9k
LWF1dG9zY2FsZXItY3B1LWluaXRpYWxpemF0aW9uLXBlcmlvZD0zMHMiLCAiLS1ob3Jpem9udGFs
LXBvZC1hdXRvc2NhbGVyLWRvd25zY2FsZS1zdGFiaWxpemF0aW9uPTMwcyIsICItLWt1YmVjb25m
aWc9L3Zhci9saWIva3ViZWxldC9rdWJlY29uZmlnIiwgIi0tbGVhZGVyLWVsZWN0PXRydWUiLCAi
LS1ub2RlLW1vbml0b3ItZ3JhY2UtcGVyaW9kPTQwcyIsICItLXBvZC1ldmljdGlvbi10aW1lb3V0
PTVtMHMiLCAiLS1wcm9maWxpbmc9ZmFsc2UiLCAiLS1yb290LWNhLWZpbGU9L2V0Yy9rdWJlcm5l
dGVzL2NlcnRzL2NhLmNydCIsICItLXJvdXRlLXJlY29uY2lsaWF0aW9uLXBlcmlvZD0xMHMiLCAi
LS1zZXJ2aWNlLWFjY291bnQtcHJpdmF0ZS1rZXktZmlsZT0vZXRjL2t1YmVybmV0ZXMvY2VydHMv
YXBpc2VydmVyLmtleSIsICItLXRlcm1pbmF0ZWQtcG9kLWdjLXRocmVzaG9sZD01MDAwIiwgIi0t
dXNlLXNlcnZpY2UtYWNjb3VudC1jcmVkZW50aWFscz10cnVlIiwgIi0tdj0yIl0KICAgICAgdm9s
dW1lTW91bnRzOgogICAgICAgIC0gbmFtZTogZXRjLWt1YmVybmV0ZXMKICAgICAgICAgIG1vdW50
UGF0aDogL2V0Yy9rdWJlcm5ldGVzCiAgICAgICAgLSBuYW1lOiB2YXItbGliLWt1YmVsZXQKICAg
ICAgICAgIG1vdW50UGF0aDogL3Zhci9saWIva3ViZWxldAogICAgICAgIC0gbmFtZTogbXNpCiAg
ICAgICAgICBtb3VudFBhdGg6IC92YXIvbGliL3dhYWdlbnQvTWFuYWdlZElkZW50aXR5LVNldHRp
bmdzCiAgICAgICAgICByZWFkT25seTogdHJ1ZQogIHZvbHVtZXM6CiAgICAtIG5hbWU6IGV0Yy1r
dWJlcm5ldGVzCiAgICAgIGhvc3RQYXRoOgogICAgICAgIHBhdGg6IC9ldGMva3ViZXJuZXRlcwog
ICAgLSBuYW1lOiB2YXItbGliLWt1YmVsZXQKICAgICAgaG9zdFBhdGg6CiAgICAgICAgcGF0aDog
L3Zhci9saWIva3ViZWxldAogICAgLSBuYW1lOiBtc2kKICAgICAgaG9zdFBhdGg6CiAgICAgICAg
cGF0aDogL3Zhci9saWIvd2FhZ2VudC9NYW5hZ2VkSWRlbnRpdHktU2V0dGluZ3MK"
        }
    ]
}
```

The above is the pattern we use to pass in a `cluster-init` spec for loading at cluster bootstrap time. E.g.:

```json
"kubernetesConfig": {
    "components": [
        {
            "name": "cluster-init",
            "enabled": true,
            "data": "YXBpVmVyc2lvbjogdjEKa2luZDogUG9kCm1ldGFkYXRhOgogIG5hbWU6IGFrcy1lbmdpbmUtcG9kLWluaXQKc3BlYzoKICBjb250YWluZXJzOgogIC0gbmFtZTogYWtzLWVuZ2luZS1wb2QtaW5pdAogICAgaW1hZ2U6IGJ1c3lib3g6MS4zMS4xCiAgICBhcmdzOgogICAgLSAvYmluL3NoCiAgICAtIC1jCiAgICAtIHdoaWxlIHRydWU7IGRvIHNsZWVwIDYwMDsgZG9uZQogIG5vZGVTZWxlY3RvcjoKICAgIGJldGEua3ViZXJuZXRlcy5pby9vczogbGludXgKLS0tCmFwaVZlcnNpb246IGJhdGNoL3YxCmtpbmQ6IEpvYgptZXRhZGF0YToKICBuYW1lOiBha3MtZW5naW5lLWpvYi1pbml0CnNwZWM6CiAgdGVtcGxhdGU6CiAgICBzcGVjOgogICAgICBjb250YWluZXJzOgogICAgICAtIGltYWdlOiBidXN5Ym94OjEuMzEuMQogICAgICAgIG5hbWU6IGJ1c3lib3gtYWdlbnQKICAgICAgICBjb21tYW5kOiBbJ3NoJywgJy1jJywgJ1sgJChlY2hvICJIZWxsbywgV29ybGQhIiB8IHNoYTI1NnN1bSB8IGN1dCAtZCIgIiAtZjEpID0gImM5OGMyNGI2NzdlZmY0NDg2MGFmZWE2ZjQ5M2JiYWVjNWJiMWM0Y2JiMjA5YzZmYzJiYmI0N2Y2NmZmMmFkMzEiIF0nXQogICAgICByZXN0YXJ0UG9saWN5OiBOZXZlcgogICAgICBub2RlU2VsZWN0b3I6CiAgICAgICAgYmV0YS5rdWJlcm5ldGVzLmlvL29zOiBsaW51eAogIGJhY2tvZmZMaW1pdDogMAo="
        }
    ]
}
```

<a name="feat-kubelet-config"></a>

#### kubeletConfig

`kubeletConfig` declares runtime configuration for the kubelet running on all master and agent nodes. It is a generic key/value object, and a child property of `kubernetesConfig`. The `kubeletConfig` configuration under `kubernetesConfig` will be inherited by a similar `kubeletConfig` configuration under the `masterProfile` configuration object, and by each `agentPoolProfile` in the `agentPoolProfiles` array. Specific master and per-pool kubelet configurations should be applied there. An example custom kubelet config:

```
"kubernetesConfig": {
    "kubeletConfig": {
        "--eviction-hard": "memory.available<250Mi,nodefs.available<20%,nodefs.inodesFree<10%"
    }
}
```

See [here](https://kubernetes.io/docs/reference/generated/kubelet/) for a reference of supported kubelet options.

Below is a list of kubelet options that AKS Engine will configure by default:

| kubelet option                        | default value                                                                                                                                                                                                                                                                                             |
| ------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| "--cadvisor-port"                     | "0"                                                                                                                                                                                                                                                                                                       |
| "--cloud-config"                      | "/etc/kubernetes/azure.json"                                                                                                                                                                                                                                                                              |
| "--cloud-provider"                    | "azure"                                                                                                                                                                                                                                                                                                   |
| "--cluster-domain"                    | "cluster.local"                                                                                                                                                                                                                                                                                           |
| "--event-qps"                         | "0"                                                                                                                                                                                                                                                                                                       |
| "--pod-infra-container-image"         | "pause-amd64:_version_"                                                                                                                                                                                                                                                                                   |
| "--network-plugin"                    | "cni"                                                                                                                                                                                                                                                                                                     |
| "--max-pods"                          | "30", or "110" if using kubenet --network-plugin (i.e., `"networkPlugin": "kubenet"`)                                                                                                                                                                                                                     |
| "--eviction-hard"                     | "memory.available<750Mi,nodefs.available<10%,nodefs.inodesFree<5%"                                                                                                                                                                                                                                        |
| "--node-status-update-frequency"      | "10s"                                                                                                                                                                                                                                                                                                     |
| "--image-gc-high-threshold"           | "85"                                                                                                                                                                                                                                                                                                      |
| "--image-gc-low-threshold"            | "80"                                                                                                                                                                                                                                                                                                      |
| "--non-masquerade-cidr"               | "0.0.0.0/0" (unless ip-masq-agent is disabled, in which case this is set to the value of `kubernetesConfig.ClusterSubnet` )                                                                                                                                                                               |
| "--azure-container-registry-config"   | "/etc/kubernetes/azure.json"                                                                                                                                                                                                                                                                              |
| "--pod-max-pids"                      | "-1" (need to activate the feature in --feature-gates=SupportPodPidsLimit=true)                                                                                                                                                                                                                           |
| "--image-pull-progress-deadline"      | "30m"                                                                                                                                                                                                                                                                                                     |
| "--feature-gates"                     | No default (can be a comma-separated list). On agent nodes `Accelerators=true` will be applied in the `--feature-gates` option for k8s versions before 1.11.0                                                                                                                                             |
| "--enforce-node-allocatable"          | "pods"                                                                                                                                                                                                                                                                                                    |
| "--streaming-connection-idle-timeout" | "4h"                                                                                                                                                                                                                                                                                                      |
| "--rotate-certificates"               | "true" (this default is set for clusters >= 1.11.9 )                                                                                                                                                                                                                                                      |
| "--tls-cipher-suites"                 | "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_GCM_SHA256" |
| "--authentication-token-webhook"      | "true" (this default is set for clusters >= 1.16.0 )                                                                                                                                                                                                                                                      |
| "--read-only-port"                    | "0" (this default is set for clusters >= 1.16.0 )                                                                                                                                                                                                                                                         |
| "--register-with-taints" | "node-role.kubernetes.io/master=true:NoSchedule" (`masterProfile` only; Note: you may add your own master-specific taints in the `kubeletConfig` under `masterProfile`, which will augment the built-in "node-role.kubernetes.io/master=true:NoSchedule" taint, which will always be present.) |
| "--container-runtime" | "remote" if in a containerd configuration, otherwise this configuration is not passed to kubelet runtime |
| "--runtime-request-timeout" | "15m" if in a containerd configuration, otherwise this configuration is not passed to kubelet runtime |
| "--container-runtime-endpoint" | "unix:///run/containerd/containerd.sock" if in a containerd configuration, otherwise this configuration is not passed to kubelet runtime |

Below is a list of kubelet options that are _not_ currently user-configurable, either because a higher order configuration vector is available that enforces kubelet configuration, or because a static configuration is required to build a functional cluster:

| kubelet option                               | default value                                      |
| -------------------------------------------- | ---------------------------------------------------|
| "--address"                                  | "0.0.0.0"                                          |
| "--allow-privileged"                         | "true"                                             |
| "--anonymous-auth"                           | "false"                                            |
| "--authorization-mode"                       | "webhook"                                          |
| "--client-ca-file"                           | "/etc/kubernetes/certs/ca.crt"                     |
| "--cluster-dns"                              | "10.0.0.10" (if ipv4, or "fd00::10" if using ipv6) |
| "--pod-manifest-path"                        | "/etc/kubernetes/manifests"                        |
| "--node-labels"                              | (based on Azure node metadata)                     |
| "--cgroups-per-qos"                          | "true"                                             |
| "--kubeconfig"                               | "/var/lib/kubelet/kubeconfig"                      |
| "--keep-terminated-pod-volumes"              | "false"                                            |
| "--tls-cert-file"                            | "/etc/kubernetes/certs/kubeletserver.crt"          |
| "--tls-private-key-file"                     | "/etc/kubernetes/certs/kubeletserver.key"          |
| "--v"                                        | "2"                                                |
| "--volume-plugin-dir"                        | "/etc/kubernetes/volumeplugins"                    |

<a name="feat-controller-manager-config"></a>

#### controllerManagerConfig

`controllerManagerConfig` declares runtime configuration for the kube-controller-manager daemon running on all master nodes. Like `kubeletConfig` it is a generic key/value object, and a child property of `kubernetesConfig`. An example custom controller-manager config:

```js
"kubernetesConfig": {
    "controllerManagerConfig": {
        "--node-monitor-grace-period": "40s",
        "--pod-eviction-timeout": "5m0s",
        "--route-reconciliation-period": "10s"
        "--terminated-pod-gc-threshold": "5000"
    }
}
```

See [here](https://kubernetes.io/docs/reference/generated/kube-controller-manager/) for a reference of supported controller-manager options.

Below is a list of controller-manager options that AKS Engine will configure by default:

| controller-manager option       | default value                              |
| ------------------------------- | ------------------------------------------ |
| "--node-monitor-grace-period"   | "40s"                                      |
| "--pod-eviction-timeout"        | "5m0s"                                     |
| "--route-reconciliation-period" | "10s"                                      |
| "--terminated-pod-gc-threshold" | "5000"                                     |
| "--feature-gates"               | No default (can be a comma-separated list) |

Below is a list of controller-manager options that are _not_ currently user-configurable, either because a higher order configuration vector is available that enforces controller-manager configuration, or because a static configuration is required to build a functional cluster:

| controller-manager option            | default value                                           |
| ------------------------------------ | ------------------------------------------------------- |
| "--kubeconfig"                       | "/var/lib/kubelet/kubeconfig"                           |
| "--allocate-node-cidrs"              | "false"                                                 |
| "--cluster-cidr"                     | _uses clusterSubnet value_                              |
| "--cluster-name"                     | _auto-generated using API model properties_             |
| "--root-ca-file"                     | "/etc/kubernetes/certs/ca.crt"                          |
| "--cluster-signing-cert-file"        | "/etc/kubernetes/certs/ca.crt"                          |
| "--cluster-signing-key-file"         | "/etc/kubernetes/certs/ca.key"                          |
| "--service-account-private-key-file" | "/etc/kubernetes/certs/apiserver.key"                   |
| "--leader-elect"                     | "true"                                                  |
| "--v"                                | "2"                                                     |
| "--profiling"                        | "false"                                                 |
| "--use-service-account-credentials"  | "false" ("true" if kubernetesConfig.enableRbac is true) |

<a name="feat-cloud-controller-manager-config"></a>

#### cloudControllerManagerConfig

`cloudControllerManagerConfig` declares runtime configuration for the cloud-controller-manager daemon running on all master nodes in a Cloud Controller Manager configuration. Like `kubeletConfig` it is a generic key/value object, and a child property of `kubernetesConfig`. An example custom cloud-controller-manager config:

```json
"kubernetesConfig": {
    "cloudControllerManagerConfig": {
        "--route-reconciliation-period": "1m"
    }
}
```

See [here](https://kubernetes.io/docs/reference/generated/cloud-controller-manager/) for a reference of supported controller-manager options.

Below is a list of cloud-controller-manager options that AKS Engine will configure by default:

| controller-manager option       | default value |
| ------------------------------- | ------------- |
| "--route-reconciliation-period" | "10s"         |

Below is a list of cloud-controller-manager options that are _not_ currently user-configurable, either because a higher order configuration vector is available that enforces controller-manager configuration, or because a static configuration is required to build a functional cluster:

| controller-manager option | default value                               |
| ------------------------- | ------------------------------------------- |
| "--kubeconfig"            | "/var/lib/kubelet/kubeconfig"               |
| "--allocate-node-cidrs"   | "false"                                     |
| "--cluster-cidr"          | _uses clusterSubnet value_                  |
| "--cluster-name"          | _auto-generated using API model properties_ |
| "--cloud-provider"        | "azure"                                     |
| "--cloud-config"          | "/etc/kubernetes/azure.json"                |
| "--leader-elect"          | "true"                                      |
| "--v"                     | "2"                                         |

<a name="feat-apiserver-config"></a>

#### apiServerConfig

`apiServerConfig` declares runtime configuration for the kube-apiserver daemon running on all master nodes. Like `kubeletConfig` and `controllerManagerConfig` it is a generic key/value object, and a child property of `kubernetesConfig`. An example custom apiserver config:

```json
"kubernetesConfig": {
    "apiServerConfig": {
        "--request-timeout": "30s"
    }
}
```

Or perhaps you want to customize/override the set of admission-control flags passed to the API Server by default, you can omit the options you don't want and specify only the ones you need as follows:

```json
"orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "orchestratorRelease": "1.8",
      "kubernetesConfig": {
        "apiServerConfig": {
          "--admission-control":  "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota,AlwaysPullImages"
        }
      }
    }
```

See [here](https://kubernetes.io/docs/reference/generated/kube-apiserver/) for a reference of supported apiserver options.

Below is a list of apiserver options that AKS Engine will configure by default:

| apiserver option                | default value                                                                                                                                                                                                                   |
| ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| "--admission-control"           | "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota" (Kubernetes versions prior to 1.9.0)                                                                                                          |
| "--enable-admission-plugins"`*` | "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,ExtendedResourceToleration" (Kubernetes versions 1.9.0 and later) |
| "--authorization-mode"          | "Node", "RBAC" (_the latter if enabledRbac is true_)                                                                                                                                                                            |
| "--audit-log-maxage"            | "30"                                                                                                                                                                                                                            |
| "--audit-log-maxbackup"         | "10"                                                                                                                                                                                                                            |
| "--audit-log-maxsize"           | "100"                                                                                                                                                                                                                           |
| "--feature-gates"               | No default (can be a comma-separated list)                                                                                                                                                                                      |
| "--oidc-username-claim"         | "oid" (_if has AADProfile_)                                                                                                                                                                                                     |
| "--oidc-groups-claim"           | "groups" (_if has AADProfile_)                                                                                                                                                                                                  |
| "--oidc-client-id"              | _calculated value that represents OID client ID_ (_if has AADProfile_)                                                                                                                                                          |
| "--oidc-issuer-url"             | _calculated value that represents OID issuer URL_ (_if has AADProfile_)                                                                                                                                                         |

`*` In Kubernetes versions 1.10.0 and later the `--admission-control` flag is deprecated and `--enable-admission-plugins` is used instead.

Below is a list of apiserver options that are _not_ currently user-configurable, either because a higher order configuration vector is available that enforces apiserver configuration, or because a static configuration is required to build a functional cluster:

| apiserver option                            | default value                                                                           |
| ------------------------------------------- | --------------------------------------------------------------------------------------- |
| "--bind-address"                            | "0.0.0.0"                                                                               |
| "--advertise-address"                       | _calculated value that represents listening URI for API server_                         |
| "--allow-privileged"                        | "true"                                                                                  |
| "--anonymous-auth"                          | "false                                                                                  |
| "--audit-log-path"                          | "/var/log/apiserver/audit.log"                                                          |
| "--insecure-port"                           | "8080"                                                                                  |
| "--secure-port"                             | "443"                                                                                   |
| "--service-account-lookup"                  | "true"                                                                                  |
| "--etcd-cafile"                             | "/etc/kubernetes/certs/ca.crt"                                                          |
| "--etcd-certfile"                           | "/etc/kubernetes/certs/etcdclient.crt"                                                  |
| "--etcd-keyfile"                            | "/etc/kubernetes/certs/etcdclient.key"                                                  |
| "--etcd-servers"                            | _calculated value that represents etcd servers_                                         |
| "--profiling"                               | "false"                                                                                 |
| "--repair-malformed-updates"                | "false" (_deprecated in v1.14_)                                                         |
| "--tls-cert-file"                           | "/etc/kubernetes/certs/apiserver.crt"                                                   |
| "--tls-private-key-file"                    | "/etc/kubernetes/certs/apiserver.key"                                                   |
| "--client-ca-file"                          | "/etc/kubernetes/certs/ca.crt"                                                          |
| "--service-account-key-file"                | "/etc/kubernetes/certs/apiserver.key"                                                   |
| "--kubelet-client-certificate"              | "/etc/kubernetes/certs/client.crt"                                                      |
| "--kubelet-client-key"                      | "/etc/kubernetes/certs/client.key"                                                      |
| "--service-cluster-ip-range"                | _see serviceCIDR_                                                                       |
| "--storage-backend"                         | _calculated value that represents etcd version_                                         |
| "--v"                                       | "4"                                                                                     |
| "--encryption-provider-config"              | "/etc/kubernetes/encryption-config.yaml" (_if enableDataEncryptionAtRest is true_)      |
| "--encryption-provider-config"              | "/etc/kubernetes/encryption-config.yaml" (_if enableEncryptionWithExternalKms is true_) |
| "--requestheader-client-ca-file"            | "/etc/kubernetes/certs/proxy-ca.crt" (_if enableAggregatedAPIs is true_)                |
| "--proxy-client-cert-file"                  | "/etc/kubernetes/certs/proxy.crt" (_if enableAggregatedAPIs is true_)                   |
| "--proxy-client-key-file"                   | "/etc/kubernetes/certs/proxy.key" (_if enableAggregatedAPIs is true_)                   |
| "--requestheader-allowed-names"             | "" (_if enableAggregatedAPIs is true_)                                                  |
| "--requestheader-extra-headers-prefix"      | "X-Remote-Extra-" (_if enableAggregatedAPIs is true_)                                   |
| "--requestheader-group-headers"             | "X-Remote-Group" (_if enableAggregatedAPIs is true_)                                    |
| "--requestheader-username-headers"          | "X-Remote-User" (_if enableAggregatedAPIs is true_)                                     |
| "--cloud-provider"                          | "azure" (_unless useCloudControllerManager is true_)                                    |
| "--cloud-config"                            | "/etc/kubernetes/azure.json" (_unless useCloudControllerManager is true_)               |

<a name="feat-scheduler-config"></a>

#### schedulerConfig

`schedulerConfig` declares runtime configuration for the kube-scheduler daemon running on all master nodes. Like `kubeletConfig`, `controllerManagerConfig`, and `apiServerConfig` it is a generic key/value object, and a child property of `kubernetesConfig`. An example custom apiserver config:

```json
"kubernetesConfig": {
    "schedulerConfig": {
        "--v": "2"
    }
}
```

See [here](https://kubernetes.io/docs/reference/generated/kube-scheduler/) for a reference of supported kube-scheduler options.

Below is a list of scheduler options that AKS Engine will configure by default:

| kube-scheduler option | default value                              |
| --------------------- | ------------------------------------------ |
| "--v"                 | "2"                                        |
| "--feature-gates"     | No default (can be a comma-separated list) |

Below is a list of kube-scheduler options that are _not_ currently user-configurable, either because a higher order configuration vector is available that enforces kube-scheduler configuration, or because a static configuration is required to build a functional cluster:

| kube-scheduler option | default value                 |
| --------------------- | ----------------------------- |
| "--kubeconfig"        | "/var/lib/kubelet/kubeconfig" |
| "--leader-elect"      | "true"                        |
| "--profiling"         | "false"                       |

We consider `kubeletConfig`, `controllerManagerConfig`, `apiServerConfig`, and `schedulerConfig` to be generic conveniences that add power/flexibility to cluster deployments. Their usage comes with no operational guarantees! They are manual tuning features that enable low-level configuration of a kubernetes cluster.

#### Custom YAML for Kubernetes component manifests

Custom YAML specifications can be configured for kube-scheduler, kube-controller-manager, cloud-controller-manager and kube-apiserver in addition to the addons described [above](#addons). You will need to pass in a _base64-encoded_ string of the kubernetes manifest YAML file to _KubernetesComponentConfig["data"]_ . For example, to pass a custom kube-scheduler config, do the following:

```json
"kubernetesConfig": {
    "schedulerConfig": {
            "data" : "<base64-encoded string of your k8s manifest YAML>"
        }
}
```

> _**NOTE**_: Custom YAML for addons is an experimental feature. Since `Addons.Data` allows you to provide your own scripts, you are responsible for any undesirable consequences of their errors or failures. Use at your own risk.

<a name="feat-sysctld-config"></a>

#### sysctldConfig

The `sysctldConfig` configuration interface allows generic Linux kernel runtime configuration that will be delivered to sysctl. Use at your own risk! It is a generic key/value object, and a child property of `masterProfile` and node pool configurations under `agentPoolProfiles`, for tuning the Linux kernel parameters on master and/or node pool VMs, respectively. An example custom sysctl config:

```
"kubernetesConfig": {
    "sysctldConfig": {
        "net.ipv4.tcp_keepalive_time": "120",
        "net.ipv4.tcp_keepalive_intvl": "75",
        "net.ipv4.tcp_keepalive_probes": "9"
    }
}
```

Kubernetes kernel configuration varies by distro, so please validate that the kernel parameter and value works for the Linux flavor you are using in your cluster.

Below is a list of sysctl configuration that AKS Engine will configure by default for both Ubuntu 16.04-LTS and 18.04-LTS, for both master and node pool VMs:

| kernel parameter                      | default value                                                                                                                                                                                                                                                                                             |
| ------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| "net.ipv4.tcp_retries2"               | "8"                                                                                                                                                                                                                                                                                                       |
| "net.core.somaxconn"                  | "16384"                                                                                                                                                                                                                                                                                                   |
| "net.ipv4.tcp_max_syn_backlog"        | "16384"                                                                                                                                                                                                                                                                                                   |
| "net.core.message_cost"               | "40"                                                                                                                                                                                                                                                                                                      |
| "net.core.message_burst"              | "80"                                                                                                                                                                                                                                                                                                      |
| "net.ipv4.neigh.default.gc_thresh1"   | "4096"                                                                                                                                                                                                                                                                                                    |
| "net.ipv4.neigh.default.gc_thresh2"   | "8192"                                                                                                                                                                                                                                                                                                    |
| "net.ipv4.neigh.default.gc_thresh3"   | "16384"                                                                                                                                                                                                                                                                                                   |

<a name="feat-private-cluster"></a>

#### privateCluster

`privateCluster` defines a cluster without public addresses assigned. It is a child property of `kubernetesConfig`.

| Name           | Required | Description                                                                                                                                          |
| -------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| enabled        | no       | Enable [Private Cluster](./features.md#feat-private-cluster) (boolean - default == false)                                                            |
| jumpboxProfile | no       | Configure and auto-provision a jumpbox to access your private cluster. `jumpboxProfile` is ignored if enabled is `false`. See `jumpboxProfile` below |

#### jumpboxProfile

`jumpboxProfile` describes the settings for a jumpbox deployed via `aks-engine` to access a private cluster. It is a child property of `privateCluster`.

| Name           | Required | Description                                                                                                                                                                              |
| -------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name           | yes      | This is the unique name for the jumpbox VM. Some resources deployed with the jumpbox are derived from this name                                                                          |
| vmSize         | yes      | Describes a valid [Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/)                                                             |
| publicKey      | yes      | The public SSH key used for authenticating access to the jumpbox. Here are instructions for [generating a public/private key pair][ssh]                                                  |
| osDiskSizeGB   | no       | Describes the OS Disk Size in GB. Defaults to `30`                                                                                                                                       |
| storageProfile | no       | Specifies the storage profile to use. Valid values are [ManagedDisks](../../examples/disks-managed) or [StorageAccount](../../examples/disks-storageaccount). Defaults to `ManagedDisks` |
| username       | no       | Describes the admin username to be used on the jumpbox. Defaults to `azureuser`                                                                                                          |

### masterProfile

`masterProfile` describes the settings for master configuration.

| Name                                                           | Required                                                                                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| -------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| count                                                          | yes                                                                                               | Masters have count value of 1, 3, or 5 masters                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| dnsPrefix                                                      | yes                                                                                               | The dns prefix for the master FQDN. The master FQDN is used for SSH or commandline access. This must be a unique name. ([bring your own VNET examples](../../examples/vnet))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| subjectAltNames                                                | no                                                                                                | An array of fully qualified domain names using which a user can reach API server. These domains are added as Subject Alternative Names to the generated API server certificate. **NOTE**: These domains **will not** be automatically provisioned.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| firstConsecutiveStaticIP                                       | only required when vnetSubnetId specified and when MasterProfile is not `VirtualMachineScaleSets` | The IP address of the first master. IP Addresses will be assigned consecutively to additional master nodes. When MasterProfile is using `VirtualMachineScaleSets`, this value will be determined by an offset from the first IP in the `vnetCidr`. For example, if `vnetCidr` is `10.239.0.0/16`, then `firstConsecutiveStaticIP` will be `10.239.0.4`                                                                                                                                                                                                                                                                                                                                                                                          |
| vmsize                                                         | yes                                                                                               | Describes a valid [Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/). These are restricted to machines with at least 2 cores and 100GB of disk space                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| storageProfile                                                 | no                                                                                                | Specifies the storage profile to use. Valid values are [ManagedDisks](../../examples/disks-managed) or [StorageAccount](../../examples/disks-storageaccount). Defaults to `ManagedDisks`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| osDiskSizeGB                                                   | no                                                                                                | Describes the OS Disk Size in GB                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| osDiskCachingType                                                   | no                                                                                                | Which disk caching type configuration to use for the OS disk on master VMs. Allowed configurations are `"ReadWrite"`, `"ReadOnly"`, or `"None"`. Default is `"ReadWrite"`. |
| vnetSubnetId                                                   | only required when using custom VNET                                                              | Specifies the Id of an alternate VNET subnet. The subnet id must specify a valid VNET ID owned by the same subscription. ([bring your own VNET examples](../../examples/vnet)). When MasterProfile is set to `VirtualMachineScaleSets`, this value should be the subnetId of the master subnet. When MasterProfile is set to `AvailabilitySet`, this value should be the subnetId shared by both master and agent nodes.                                                                                                                                                                                                                                                                                                                        |
| extensions                                                     | no                                                                                                | This is an array of extensions. This indicates that the extension be run on a single master. The name in the extensions array must exactly match the extension name in the extensionProfiles                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| vnetCidr                                                       | no                                                                                                | Specifies the VNET cidr when using a custom VNET ([bring your own VNET examples](../../examples/vnet)). This VNET cidr should include both the master and the agent subnets.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| imageReference.name                                            | no                                                                                                | The name of the Linux OS image. Needs to be used in conjunction with resourceGroup, below. (For information on setting this for Windows nodes see [WindowsProfile](#windowsProfile))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| imageReference.resourceGroup                                   | no                                                                                                | Resource group that contains the Linux OS image. Needs to be used in conjunction with name, above                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| imageReference.subscriptionId                                  | no                                                                                                | ID of subscription containing the Linux OS image. Applies only to Shared Image Galleries. All of name, resourceGroup, subscription, gallery, image name, and version must be specified for this scenario.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| imageReference.gallery                                         | no                                                                                                | Name of Shared Image Gallery containing the Linux OS image. Applies only to Shared Image Galleries. All of name, resourceGroup, subscription, gallery, image name, and version must be specified for this scenario.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| imageReference.version                                         | no                                                                                                | Version containing the Linux OS image. Applies only to Shared Image Galleries. All of name, resourceGroup, subscription, gallery, image name, and version must be specified for this scenario.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| distro                                                         | no                                                                                                | Specifies the masters' Linux distribution. Currently supported values are: `ubuntu`, `ubuntu-18.04`, `ubuntu-18.04-gen2` (Ubuntu 18.04-LTS running on a [Generation 2 VM](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/generation-2)), `aks-ubuntu-16.04` (previously `aks`), and `aks-ubuntu-18.04`. For Azure Public Cloud, Azure US Government Cloud and Azure China Cloud, defaults to `aks-ubuntu-16.04`. For other Sovereign Clouds, the default is `ubuntu-16.04` (There is a [known issue](https://github.com/Azure/aks-engine/issues/761) with `ubuntu-18.04` + Azure CNI). `aks-ubuntu-16.04` is a custom image based on `ubuntu-16.04` that comes with pre-installed software necessary for Kubernetes deployments. |
| customFiles                                                    | no                                                                                                | The custom files to be provisioned to the master nodes. Defined as an array of JSON objects with each defined as `"source":"absolute-local-path", "dest":"absolute-path-on-masternodes"`.[See examples](../../examples/customfiles)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| availabilityProfile                                            | no                                                                                                | Supported values are `AvailabilitySet` (default) and `VirtualMachineScaleSets` (still under development: upgrade not supported; requires Kubernetes clusters version 1.10+ and agent pool availabilityProfile must also be `VirtualMachineScaleSets`). When MasterProfile is using `VirtualMachineScaleSets`, to SSH into a master node, you need to use `ssh -p 50001` instead of port 22.                                                                                                                                                                                                                                                                                                                                                     |
| agentVnetSubnetId                                              | only required when using custom VNET and when MasterProfile is using `VirtualMachineScaleSets`    | Specifies the Id of an alternate VNET subnet for all the agent pool nodes. The subnet id must specify a valid VNET ID owned by the same subscription. ([bring your own VNET examples](../../examples/vnet)). When MasterProfile is using `VirtualMachineScaleSets`, this value should be the subnetId of the subnet for all agent pool nodes.                                                                                                                                                                                                                                                                                                                                                                                                   |
| [availabilityZones](../../examples/kubernetes-zones/README.md) | no                                                                                                | To protect your cluster from datacenter-level failures, you can enable the Availability Zones feature for your master VMs. Check out [Availability Zones README](../../examples/kubernetes-zones/README.md) for more details.                                                                                                                                                                                                                                                                                                                                                                                              |
| cosmosEtcd                                                     | no                                                                                                | True: uses cosmos etcd endpoint instead of installing etcd on masters                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| auditDEnabled                                                  | no                                                                                                | Enable auditd enforcement at the OS layer for each node VM. This configuration is only valid on an agent pool with an Ubuntu-backed distro, i.e., the default "aks-ubuntu-16.04" distro, or the "aks-ubuntu-18.04", "ubuntu", "ubuntu-18.04", "ubuntu-18.04-gen2", or "acc-16.04" distro values. Defaults to `false`                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| ultraSSDEnabled                                            | no                                                             | Enable UltraSSD feature for each node VM. More details about [Ultra disk](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/disks-types#ultra-ssd-preview).                                                                                                                                                                                                                                                                                                                                                                |
| customVMTags                                                   | no                                                                                                | Specifies a list of custom tags to be added to the master VMs or Scale Sets. Each tag is a key/value pair (ie: `"myTagKey": "myTagValue"`).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| sysctldConfig                    | no                        | Configure Linux kernel parameters via /etc/sysctl.d/. See `sysctldConfig` [below](#feat-sysctld-config)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| proximityPlacementGroupID        | no                        | Specifies the resource id of the Proximity Placement Group (PPG) to be used for master VMs.  Please find more details about PPG in this [Azure blog](https://azure.microsoft.com/en-us/blog/introducing-proximity-placement-groups). Note that the PPG should be created in advance. The following [Azure CLI documentation](https://docs.microsoft.com/en-us/cli/azure/ppg?view=azure-cli-latest#az-ppg-create) explains how to create a PPG. |
| kubeletConfig                    | no                        | Configure various runtime configuration for kubelet running on master nodes. See `kubeletConfig` [above](#feat-kubelet-config) |

### agentPoolProfiles

A cluster can have 0 to 12 agent pool profiles. Agent Pool Profiles are used for creating agents with different capabilities such as VMSizes, VMSS or Availability Set, Public/Private access, user-defined OS Images, [attached storage disks](../../examples/disks-storageaccount), [attached managed disks](../../examples/disks-managed), or [Windows](../../examples/windows).

| Name                                                           | Required                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| -------------------------------------------------------------- | -------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| availabilityProfile                                            | no                                                                   | Supported values are `VirtualMachineScaleSets` (default, except for Kubernetes clusters before version 1.10) and `AvailabilitySet`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| count                                                          | yes                                                                  | Describes the node count                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| [availabilityZones](../../examples/kubernetes-zones/README.md) | no                                                                   | To protect your cluster from datacenter-level failures, you can enable the Availability Zones feature for your node pool. Check out [Availability Zones README](../../examples/kubernetes-zones/README.md) for more details.                                                                                                                                                                                                                                                                                                                                                                                        |
| singlePlacementGroup                                           | no                                                                   | Supported values are `true` (default) and `false`. A value of `true`: A VMSS with a single placement group and has a range of 0-100 VMs. A value of `false`: A VMSS with multiple placement groups and has a range of 0-1,000 VMs. For more information, check out [virtual machine scale sets placement groups](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups). This configuration is only valid on an agent pool with an `"availabilityProfile"` value of `"VirtualMachineScaleSets"`                                                                                                                                                                                   |
| scaleSetPriority                                               | no                                                                   | Supported values are `Regular` (default), `Low` (DEPRECATED) and `Spot`. This configuration is only valid on an agent pool with an `"availabilityProfile"` value of `"VirtualMachineScaleSets"`. Enables the usage of [Low-priority VMs on Scale Sets](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-use-low-priority) or [Azure Spot VMs](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/use-spot).                                                                                                                                                                                                                                                                                 |
| scaleSetEvictionPolicy                                         | no                                                                   | Supported values are `Delete` (default) and `Deallocate`. This configuration is only valid on an agent pool with an `"availabilityProfile"` value of `"VirtualMachineScaleSets"` and a `"scaleSetPriority"` value of `"Low"` or `"Spot"`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| spotMaxPrice                                                   | no                                                                   | Supported values are `-1` (default) or any decimal value greater than zero. Only valid on an agent pool with an `"availabilityProfile"` value of `"VirtualMachineScaleSets"` and `scaleSetPriority` value of `"Spot"`. Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS. This price is in US Dollars. This price will be compared with the current Azure Spot price for the VM size. Also, the prices are compared at the time of create/update of Azure Spot VM/VMSS and the operation will only succeed if the maxPrice is greater than the current Azure Spot price. The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the maxPrice after creation of VM/VMSS. You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons. Also, the default max price is -1 if it is not provided by you |
| diskSizesGB                                                    | no                                                                   | Describes an array of up to 4 attached disk sizes. Valid disk size values are between 1 and 1024                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| dataDiskCachingType                                                   | no                                                                                                | Which disk caching type configuration to use for the data disks on master VMs. Allowed configurations are `"ReadWrite"`, `"ReadOnly"`, or `"None"`. Default is `"ReadOnly"`. |
| dnsPrefix                                                      | Required if agents are to be exposed publically with a load balancer | The dns prefix that forms the FQDN to access the loadbalancer for this agent pool. This must be a unique name among all agent pools. Not supported for Kubernetes clusters                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| name                                                           | yes                                                                  | This is the unique name for the agent pool profile. The resources of the agent pool profile are derived from this name                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ports                                                          | only required if needed for exposing services publically             | Describes an array of ports need for exposing publically. A tcp probe is configured for each port and only opens to an agent node if the agent node is listening on that port. A maximum of 150 ports may be specified. Not supported for Kubernetes clusters                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| storageProfile                                                 | no                                                                   | Specifies the storage profile to use. Valid values are [ManagedDisks](../../examples/disks-managed), [StorageAccount](../../examples/disks-storageaccount), or [Ephemeral](../../examples/disks-ephemeral). Defaults to `ManagedDisks`. `Ephemeral` is an experimental feature - please read more on the [feature status page](features.md)                                                                                                                                                                                                                                                                                                                                                                                               |
| vmsize                                                         | yes                                                                  | Describes a valid [Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/). These are restricted to machines with at least 2 cores                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| osDiskSizeGB                                                   | no                                                                   | Describes the OS Disk Size in GB                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| osDiskCachingType                                                   | no                                                                                                | Which disk caching type configuration to use for the OS disk on master VMs. Allowed configurations are `"ReadWrite"`, `"ReadOnly"`, or `"None"`. Default is `"ReadWrite"`. |
| vnetSubnetId                                                   | no                                                                   | Specifies the Id of an alternate VNET subnet. The subnet id must specify a valid VNET ID owned by the same subscription. ([bring your own VNET examples](../../examples/vnet))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| imageReference.name                                            | no                                                                   | The name of a a Linux OS image. Needs to be used in conjunction with resourceGroup, below                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| imageReference.resourceGroup                                   | no                                                                   | Resource group that contains the Linux OS image. Needs to be used in conjunction with name, above                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| osType                                                         | no                                                                   | Specifies the agent pool's Operating System. Supported values are `Windows` and `Linux`. Defaults to `Linux`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| distro                                                         | no                                                                   | Specifies the masters' Linux distribution. Currently supported values are: `ubuntu`, `ubuntu-18.04`, `aks-ubuntu-16.04` (previously `aks`), `aks-ubuntu-18.04`, and `ubuntu-18.04-gen2` (Ubuntu 18.04-LTS running on a [Generation 2 VM](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/generation-2)). For Azure Public Cloud, Azure US Government Cloud and Azure China Cloud, defaults to `aks-ubuntu-16.04`. For Sovereign Clouds, the default is `ubuntu-16.04` (There is a [known issue](https://github.com/Azure/aks-engine/issues/761) with `ubuntu-18.04` + Azure CNI). `aks-ubuntu-16.04` is a custom image based on `ubuntu-16.04` that comes with pre-installed software necessary for Kubernetes deployments. |
| acceleratedNetworkingEnabled                                   | no                                                                   | Use [Azure Accelerated Networking](https://azure.microsoft.com/en-us/blog/maximize-your-vm-s-performance-with-accelerated-networking-now-generally-available-for-both-windows-and-linux/) feature for Linux agents (You must select a VM SKU that supports Accelerated Networking). Defaults to `true` if the VM SKU selected supports Accelerated Networking                                                                                                                                                                                                                                                                                                                                                                             |
| acceleratedNetworkingEnabledWindows                            | no                                                                   | Currently unstable, and disabled for new clusters!                                                                                                                                                                                                                                                                                                                                                                                                                               |
| vmssOverProvisioningEnabled                                    | no                                                                   | Use [Overprovisioning](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-design-overview#overprovisioning) with VMSS. This configuration is only valid on an agent pool with an `"availabilityProfile"` value of `"VirtualMachineScaleSets"`. Defaults to `false`                                                                                                                                                                                                                                                                                                                                                                                                                              |
| enableVMSSNodePublicIP                                         | no                                                                   | Enable creation of public IP on VMSS nodes. This configuration is only valid on an agent pool with an `"availabilityProfile"` value of `"VirtualMachineScaleSets"`. Defaults to `false`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| LoadBalancerBackendAddressPoolIDs                              | no                                                                   | Enables automatic placement of the agent pool nodes into existing load balancer's backend address pools. Each element value of this string array is the corresponding load balancer backend address pool's Azure Resource Manager(ARM) resource ID. By default this property is not included in the API model, which is equivalent to an empty string array.                                                                                                                                                                                                                                                                                                                                                                              |
| auditDEnabled                                                  | no                                                                   | Enable auditd enforcement at the OS layer for each node VM. This configuration is only valid on an agent pool with an Ubuntu-backed distro, i.e., the default "aks-ubuntu-16.04" distro, or the "aks-ubuntu-18.04", "ubuntu", "ubuntu-18.04", or "acc-16.04" distro values. Defaults to `false`                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| customVMTags                                                   | no                                                                   | Specifies a list of custom tags to be added to the agent VMs or Scale Sets. Each tag is a key/value pair (ie: `"myTagKey": "myTagValue"`).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| diskEncryptionSetID                                            | no                                                             | Specifies ResourceId of the disk encryption set to use for enabling encryption at rest (ie: `"/subscriptions/{subs-id}/resourceGroups/{rg-name}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSet-name}"`). More details about [Server side encryption of Azure managed disks](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/disk-encryption).                                                                                                                                                                                                                                                                                                                                                                |
| ultraSSDEnabled                                            | no                                                             | Enable UltraSSD feature for each node VM. More details about [Ultra disk](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/disks-types#ultra-ssd-preview).                                                                                                                                                                                                                                                                                                                                                                |
| extensions | no | Specifies list of extensions to enable for the agent profile. More details about [agentPoolProfiles extensions](extensions.md#agentpoolprofiles) |
| preProvisionExtension | no | Specifies an extension to be run before the cluster is brought up. More details about [agentPoolProfiles extensions](extensions.md#agentpoolprofiles) |
| sysctldConfig                    | no                        | Configure Linux kernel parameters via /etc/sysctl.d/. See `sysctldConfig` [below](#feat-sysctld-config) |
| proximityPlacementGroupID        | no                        | Specifies the resource id of the Proximity Placement Group (PPG) to be used for this agentpool.  Please find more details about PPG in this [Azure blog](https://azure.microsoft.com/en-us/blog/introducing-proximity-placement-groups). Note that the PPG should be created in advance. The following [Azure CLI documentation](https://docs.microsoft.com/en-us/cli/azure/ppg?view=azure-cli-latest#az-ppg-create) explains how to create a PPG. |
| kubeletConfig                    | no                        | Configure various runtime configuration for kubelet running on this node pool. See `kubeletConfig` [above](#feat-kubelet-config) |

### linuxProfile

`linuxProfile` provides the linux configuration for each linux node in the cluster

| Name                             | Required | Description                                                                              |
| -------------------------------- | -------- | ---------------------------------------------------------------------------------------- |
| adminUsername                    | yes      | Describes the username to be used on all linux clusters                                  |
| ssh.publicKeys[].keyData         | yes      | The public SSH key used for authenticating access to all Linux nodes in the cluster      |
| secrets                          | no       | Specifies an array of key vaults to pull secrets from and what secrets to pull from each |
| customSearchDomain.name          | no       | describes the search domain to be used on all linux clusters                             |
| customSearchDomain.realmUser     | no       | describes the realm user with permissions to update dns registries on Windows Server DNS |
| customSearchDomain.realmPassword | no       | describes the realm user password to update dns registries on Windows Server DNS         |
| customNodesDNS.dnsServer         | no       | describes the IP address of the DNS Server                                               |

Here are instructions for [generating a public/private key pair][ssh] for `ssh.publicKeys.keyData`.

#### Notes on SSH public keys

At least one SSH key is required, but multiple are supported when deploying Kubernetes.

Here's a minimal example using just one key:

```
    "linuxProfile": {
      "adminUsername": "azureuser",
      "ssh": {
        "publicKeys": [
          {
            "keyData": "ssh-rsa AAAA...w=="
          }
        ]
      }
    },
```

And an example using two keys.

```json
    "linuxProfile": {
      "adminUsername": "azureuser",
      "ssh": {
        "publicKeys": [
          {
            "keyData": "ssh-rsa AAAA...w=="
          },
          {
            "keyData": "ssh-rsa AAAA...w=="
          }
        ]
      }
    },
```

#### secrets

`secrets` details which certificates to install on the masters and nodes in the cluster.

A cluster can have a list of key vaults to install certs from.

On linux boxes the certs are saved on under the directory "/var/lib/waagent/". 2 files are saved per certificate:

1.  `{thumbprint}.crt` : this is the full cert chain saved in PEM format
2.  `{thumbprint}.prv` : this is the private key saved in PEM format

| Name                             | Required | Description                                                         |
| -------------------------------- | -------- | ------------------------------------------------------------------- |
| sourceVault.id                   | yes      | The azure resource manager id of the key vault to pull secrets from |
| vaultCertificates.certificateUrl | yes      | Keyvault URL to this cert including the version                     |

format for `sourceVault.id`, can be obtained in cli, or found in the portal: /subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.KeyVault/vaults/{keyvaultname}

format for `vaultCertificates.certificateUrl`, can be obtained in cli, or found in the portal:
https://{keyvaultname}.vault.azure.net:443/secrets/{secretName}/{version}

### windowsProfile

`windowsProfile` provides configuration specific to Windows nodes in the cluster

| Name                          | Required | Description                                                                                                                                                                                                                                                     |
| ----------------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| adminUsername                 | yes      | Username for the Windows adminstrator account created on each Windows node                                                                                                                                                                                      |
| adminPassword                 | yes      | Password for the Windows adminstrator account created on each Windows node                                                                                                                                                                                      |
| csiProxyURL                   | no       | Path to a package containing csi proxy binaries for Windows.                                                                                                                                                                                                    |
| enableAutomaticUpdates        | no       | If set to `true` Windows Update will be configured to automatically apply updates on Windows nodes. Default: `false`                                                                                                                                            |
| enableCSIProxy                | no       | If set to `true` the csi-proxy specified by `windowsProfile.csiProxyURL` will get installed during node provisioning. See [Windows Csi Proxy](csi-proxy-windows.md) for more details.                                                                           |
| windowsPublisher              | no       | Publisher used to find Windows VM to deploy from marketplace. Default: `microsoft-aks`                                                                                                                                                                          |
| windowsOffer                  | no       | Offer used to find Windows VM to deploy from marketplace. Default: `aks-windows`                                                                                                                                                                                |
| windowsSku                    | no       | SKU usedto find Windows VM to deploy from marketplace. Default: `2019-datacenter-core-smalldisk`                                                                                                                                                                |
| imageVersion                  | no       | Specific image version to deploy from marketplace. Default: `17763.737.190923`. This default is incremented to include the latest Windows patches after being validated by the AKS Engine team.                                                                 |
| windowsImageSourceURL         | no       | Path to an existing Azure storage blob with a sysprepped VHD. This is used to test pre-release or customized VHD files that you have uploaded to Azure. If provided, the above 4 parameters are ignored.                                                        |
| imageReference.name           | no       | Name of an Image.                                                                                                                                                                                                                                               |
| imageReference.resourceGroup  | no       | Resource group that contains the Image.                                                                                                                                                                                                                         |
| imageReference.subscriptionId | no       | ID of subscription containing a Shared Image Gallery.                                                                                                                                                                                                           |
| imageReference.gallery        | no       | Name of a Shared Image Gallery.                                                                                                                                                                                                                                 |
| imageReference.version        | no       | Version of an Image from a Shared Image Gallery.                                                                                                                                                                                                                |
| sshEnabled                    | no       | If set to `true`, OpenSSH will be installed on windows nodes to allow for ssh remoting. **Only for Windows version 1809/2019 or later**. The same SSH authorized public key(s) will be added from [linuxProfile.ssh.publicKeys](#linuxProfile). Default: `true` |

#### Windows Images

You can configure the image used for all Windows nodes one of the following ways:

##### Defaults

The AKS Engine team produces images that are optimized for and validated against `aks-engine`-created Kubernetes clusters during the regular development and release process. The latest version of these images at the time of a new release of the `aks-engine` binary are used as the default images for Windows nodes.

These images are published to the Azure Marketplace under the `microsoft-aks` publisher and `aks-windows` offer.
Release notes for these images can be found under [releases/vhd-notes/aks-windows](../../releases/vhd-notes/aks-windows).

##### Marketplace Images

Aks-engine also supports running 'vanilla' Windows Server images published by Microsoft.
These can be used by advanced users if a release other than Winders Server 2019 is needed.

If you want to choose a specific Windows image, but automatically use the latest - set `windowsPublisher`, `windowsOffer`, and `windowsSku`. If you need a specific version, then add `imageVersion` too.

You can find all available images with `az vm image list --all --publisher MicrosoftWindowsServer --offer WindowsServer --output table`, and the contents of these images are described in the knowledge base article [Windows Server release on Azure Marketplace update history](https://support.microsoft.com/en-us/help/4497947).

If you want to use a specific image then `windowsPublisher`, `windowsOffer`, `windowsSku`, and `imageVersion` must all be set:

```json
"windowsProfile": {
            "adminUsername": "...",
            "adminPassword": "...",
            "windowsPublisher": "MicrosoftWindowsServer",
            "windowsOffer": "WindowsServer",
            "windowsSku": "2019-Datacenter-Core-with-Containers-smalldisk",
            "imageVersion": "2019.0.20181107"
     },
```

##### Custom Images

Listed in order of precedence based on what is specified in the API model:

###### VHD

To use an image uploaded to an Azure storage account (or any other accessible location) specify `windowsImageSourceURL`.

**Note:** URLs containing SAS tokens are not allowed by Azure!

```json
"windowsProfile": {
            "adminUsername": "...",
            "adminPassword": "...",
            "windowsImageSourceURL": "https://images.blob.core.windows.net/vhds/custom_windows_image.vhd",
     },
```

###### Shared Image Gallery

To use an Image from a Shared Image Gallery specify `imageReference.name`, `imageReference.resourceGroup`, `imageReference.subscriptionId`, `imageReference.galllery`, and `imageReference.version`.

```json
"windowsProfile": {
            "adminUsername": "...",
            "adminPassword": "...",
            "imageReference": {
              "name": "custom-image",
              "resourceGroup": "windows-images",
              "subscriptionId": "00000000-0000-0000-0000-000000000000",
              "gallery": "image-gallery",
              "version": "0.1.0"
            }
     },
```

###### Azure Image

To use a pre-existing Azure Image specify `imageReference.name` and `imageReference.resourceGroup`.

```json
"windowsProfile": {
            "adminUsername": "...",
            "adminPassword": "...",
            "imageReference": {
              "name": "custom-image",
              "resourceGroup": "windows-images"
            }
     },
```

### servicePrincipalProfile

`servicePrincipalProfile` describes an Azure Service credentials to be used by the cluster for self-configuration. See [service principal](service-principals.md) for more details on creation.

| Name                         | Required                          | Description                                                                                                 |
| ---------------------------- | --------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| clientId                     | yes, for Kubernetes clusters      | describes the Azure client id. It is recommended to use a separate client ID per cluster                    |
| secret                       | yes, for Kubernetes clusters      | describes the Azure client secret. It is recommended to use a separate client secret per client id          |
| objectId                     | optional, for Kubernetes clusters | describes the Azure service principal object id. It is required if enableEncryptionWithExternalKms is true  |
| keyvaultSecretRef.vaultId    | no, for Kubernetes clusters       | describes the vault id of the keyvault to retrieve the service principal secret from. See below for format. |
| keyvaultSecretRef.secretName | no, for Kubernetes clusters       | describes the name of the service principal secret in keyvault                                              |
| keyvaultSecretRef.version    | no, for Kubernetes clusters       | describes the version of the secret to use                                                                  |

format for `keyvaultSecretRef.vaultId`, can be obtained in cli, or found in the portal:
`/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>`. See [keyvault params](../../examples/keyvault-params/README.md#service-principal-profile) for an example.
