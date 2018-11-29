# Microsoft Azure Container Service Engine - Cluster Definition

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

1.  `DCOS` - this represents the [DC/OS orchestrator](dcos.md). [Older releases of DCOS 1.8 may be specified](../examples/dcos-releases).
2.  `Kubernetes` - this represents the [Kubernetes orchestrator](kubernetes.md).
3.  `Swarm` - this represents the [Swarm orchestrator](swarm.md).
4.  `Swarm Mode` - this represents the [Swarm Mode orchestrator](swarmmode.md).

To learn more about supported orchestrators and versions, run the orchestrators command:

```/bin/aks-engine orchestrators```


### kubernetesConfig

`kubernetesConfig` describes Kubernetes specific configuration.

| Name                            | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                   |
| ------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| addons                          | no       | Configure various Kubernetes addons configuration (currently supported: tiller, kubernetes-dashboard). See `addons` configuration below                                                                                                                                                                                                                                                                       |
| apiServerConfig                 | no       | Configure various runtime configuration for apiserver. See `apiServerConfig` [below](#feat-apiserver-config)                                                                                                                                                                                                                                                                                                  |
| cloudControllerManagerConfig    | no       | Configure various runtime configuration for cloud-controller-manager. See `cloudControllerManagerConfig` [below](#feat-cloud-controller-manager-config)                                                                                                                                                                                                                                                       |
| clusterSubnet                   | no       | The IP subnet used for allocating IP addresses for pod network interfaces. The subnet must be in the VNET address space. With Azure CNI enabled, the default value is 10.240.0.0/12. Without Azure CNI, the default value is 10.244.0.0/16.                                            |
| containerRuntime                | no       | The container runtime to use as a backend. The default is `docker`. The other options are `clear-containers`, `kata-containers`, and `containerd`                                                                                                                                                                                                                                                             |
| controllerManagerConfig         | no       | Configure various runtime configuration for controller-manager. See `controllerManagerConfig` [below](#feat-controller-manager-config)                                                                                                                                                                                                                                                                        |
| customWindowsPackageURL         | no       | Configure custom windows Kubernetes release package URL for deployment on Windows that is generated by scripts/build-windows-k8s.sh.  The format of this file is a zip file with multiple items (binaries, cni, infra container) in it.  This setting will be depreciated in future release of aks-engine where the binaries will be pulled in the format of Kubernetes releases that only contain the kubernetes binaries.                                                                                                                                                                                                                                                                                         |
| WindowsNodeBinariesURL          | no       | Windows Kubernetes Node binaries can be provided in the format of Kubernetes release (example: https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.11.md#node-binaries-1). This setting allows overriding the binaries for custom builds.                                                                                                                                                                                                                                                                                         |
| dnsServiceIP                    | no       | IP address for kube-dns to listen on. If specified must be in the range of `serviceCidr`                                                                                                                                                                                                                                                                                                                      |
| dockerBridgeSubnet              | no       | The specific IP and subnet used for allocating IP addresses for the docker bridge network created on the kubernetes master and agents. Default value is 172.17.0.1/16. This value is used to configure the docker daemon using the [--bip flag](https://docs.docker.com/engine/userguide/networking/default_network/custom-docker0)                                                                           |
| enableAggregatedAPIs            | no       | Enable [Kubernetes Aggregated APIs](https://kubernetes.io/docs/concepts/api-extension/apiserver-aggregation/).This is required by [Service Catalog](https://github.com/kubernetes-incubator/service-catalog/blob/master/README.md). (boolean - default is true for k8s versions greater or equal to 1.9.0, false otherwise)                                                                                                                                              |
| enableDataEncryptionAtRest      | no       | Enable [kubernetes data encryption at rest](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/).This is currently an alpha feature. (boolean - default == false)                                                                                                                                                                                                                               |
| enableEncryptionWithExternalKms | no       | Enable [kubernetes data encryption at rest with external KMS](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/).This is currently an alpha feature. (boolean - default == false)                                                                                                                                                                                                             |
| enablePodSecurityPolicy         | no       | Enable [kubernetes pod security policy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/).This is currently a beta feature. (boolean - default == false)                                                                                                                                                                                                                                       |
| enableRbac                      | no       | Enable [Kubernetes RBAC](https://kubernetes.io/docs/admin/authorization/rbac/) (boolean - default == true)                                                                                                                                                                                                                                                                                                    |
| etcdDiskSizeGB                  | no       | Size in GB to assign to etcd data volume. Defaults (if no user value provided) are: 256 GB for clusters up to 3 nodes; 512 GB for clusters with between 4 and 10 nodes; 1024 GB for clusters with between 11 and 20 nodes; and 2048 GB for clusters with more than 20 nodes                                                                                                                                   |
| etcdEncryptionKey               | no       | Enryption key to be used if enableDataEncryptionAtRest is enabled. Defaults to a random, generated, key                                                                                                                                                                                                                                                                                                       |
| gcHighThreshold                 | no       | Sets the --image-gc-high-threshold value on the kublet configuration. Default is 85. [See kubelet Garbage Collection](https://kubernetes.io/docs/concepts/cluster-administration/kubelet-garbage-collection/)                                                                                                                                                                                                 |
| gcLowThreshold                  | no       | Sets the --image-gc-low-threshold value on the kublet configuration. Default is 80. [See kubelet Garbage Collection](https://kubernetes.io/docs/concepts/cluster-administration/kubelet-garbage-collection/)                                                                                                                                                                                                  |
| kubeletConfig                   | no       | Configure various runtime configuration for kubelet. See `kubeletConfig` [below](#feat-kubelet-config)                                                                                                                                                                                                                                                                                                        |
| kubernetesImageBase             | no       | Specifies the default image base URL (everything preceding the actual image filename) to be used for all kubernetes-related containers such as hyperkube, cloud-controller-manager, pause, addon-manager, heapster, exechealthz etc. e.g., `k8s.gcr.io/`                                                                                                                                                                                                                                     |
| loadBalancerSku                 | no       | Sku of Load Balancer and Public IP. Candidate values are: `basic` and `standard`. If not set, it will be default to basic. Requires Kubernetes 1.11 or newer. NOTE: VMs behind ILB standard SKU will not be able to access the internet without ELB configured with at least one frontend IP as described in the [standard loadbalancer outbound connectivity doc](https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-standard-overview#control-outbound-connectivity). For Kubernetes 1.11 and 1.12, We have created an external loadbalancer service in the kube-system namespace as a workaround to this issue. Starting k8s 1.13, instead of creating an ELB service, we will setup outbound rules in ARM template once the API is available.                                                                                                                                                                                                                                                                                                          |
| networkPlugin                   | no       | Specifies the network plugin implementation for the cluster. Valid values are:<br>`"azure"` (default), which provides an Azure native networking experience <br>`"kubenet"` for k8s software networking implementation. <br> `"flannel"` for using CoreOS Flannel <br> `"cilium"` for using the default Cilium CNI IPAM                                                                                       |
| networkPolicy                   | no       | Specifies the network policy enforcement tool for the cluster (currently Linux-only). Valid values are:<br>`"calico"` for Calico network policy.<br>`"cilium"` for cilium network policy (Lin), and `"azure"` (experimental) for Azure CNI-compliant network policy (note: Azure CNI-compliant network policy requires explicit `"networkPlugin": "azure"` configuration as well).<br>See [network policy examples](../examples/networkpolicy) for more information.                                                                                                                                  |
| privateCluster                  | no       | Build a cluster without public addresses assigned. See `privateClusters` [below](#feat-private-cluster).                                                                                                                                                                                                                                                                                                      |
| schedulerConfig                 | no       | Configure various runtime configuration for scheduler. See `schedulerConfig` [below](#feat-scheduler-config)                                                                                                                                                                                                                                                                                                  |
| serviceCidr                     | no       | IP range for Service IPs, Default is "10.0.0.0/16". This range is never routed outside of a node so does not need to lie within clusterSubnet or the VNET                                                                                                                                                                                                                                                     |
| useInstanceMetadata             | no       | Use the Azure cloudprovider instance metadata service for appropriate resource discovery operations. Default is `true`                                                                                                                                                                                                                                                                                        |
| useManagedIdentity              | no       | Includes and uses MSI identities for all interactions with the Azure Resource Manager (ARM) API. Instead of using a static service principal written to /etc/kubernetes/azure.json, Kubernetes will use a dynamic, time-limited token fetched from the MSI extension running on master and agent nodes. This support is currently alpha and requires Kubernetes v1.9.1 or newer. (boolean - default == false). When MasterProfile is using `VirtualMachineScaleSets`, this feature requires Kubernetes v1.12 or newer as we default to using user assigned identity. |
| azureCNIURLLinux                | no       | Deploy a private build of Azure CNI on Linux nodes. This should be a full path to the .tar.gz |
| azureCNIURLWindows              | no       | Deploy a private build of Azure CNI on Windows nodes. This should be a full path to the .tar.gz |

#### addons

`addons` describes various addons configuration. It is a child property of `kubernetesConfig`. Below is a list of currently available addons:

| Name of addon                                                         | Enabled by default? | How many containers | Description                                                                                                                                                         |
| --------------------------------------------------------------------- | ------------------- | ------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| tiller                                                                | true                | 1                   | Delivers the Helm server-side component: tiller. See https://github.com/kubernetes/helm for more info                                                               |
| kubernetes-dashboard                                                  | true                | 1                   | Delivers the Kubernetes dashboard component. See https://github.com/kubernetes/dashboard for more info                                                              |
| rescheduler                                                           | false               | 1                   | Delivers the Kubernetes rescheduler component                                                                                                                       |
| [cluster-autoscaler](../examples/addons/cluster-autoscaler/README.md) | false               | 1                   | Delivers the Kubernetes cluster autoscaler component. See https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/cloudprovider/azure for more info |
| [nvidia-device-plugin](../examples/addons/nvidia-device-plugin/README.md) | true if using a Kubernetes cluster (v1.10+) with an N-series agent pool               | 1                   | Delivers the Kubernetes NVIDIA device plugin component. See https://github.com/NVIDIA/k8s-device-plugin for more info |
| container-monitoring                       | false               | 1                   | Delivers the Kubernetes container monitoring component |
| [blobfuse-flexvolume](https://github.com/Azure/kubernetes-volume-drivers/tree/master/flexvolume/blobfuse)                        | true               | as many as linux agent nodes                   | Access virtual filesystem backed by the Azure Blob storage |
| [smb-flexvolume](https://github.com/Azure/kubernetes-volume-drivers/tree/master/flexvolume/smb)                        | true               | as many as linux agent nodes                   | Access SMB server by using CIFS/SMB protocol |
| [keyvault-flexvolume](../examples/addons/keyvault-flexvolume/README.md)                        | true               | as many as linux agent nodes                   | Access secrets, keys, and certs in Azure Key Vault from pods |
| [aad-pod-identity](../examples/addons/aad-pod-identity/README.md)                        | false               | 1 + 1 on each linux agent nodes | Assign Azure Active Directory Identities to Kubernetes applications |

To give a bit more info on the `addons` property: We've tried to expose the basic bits of data that allow useful configuration of these cluster features. Here are some example usage patterns that will unpack what `addons` provide:

To enable an addon (using "tiller" as an example):

```
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

```
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

```
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
              "maxNodes": "5",
              "minNodes": "1"
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

Additionally above, we specified a custom docker image for tiller, let's say we want to build a cluster and test an alpha version of tiller in it. **Important note!** customizing the image is not sticky across upgrade/scale, to ensure that aks-engine always delivers a version-curated, known-working addon when moving a cluster to a new version. Considering all that, providing a custom image reference for an addon configuration should be considered for testing/development, but not for a production cluster. If you'd like to entirely customize one of the addons available, including across scale/upgrade operations, you may include in an addon's spec a gzip+base64-encoded (in that order) string of a Kubernetes yaml manifest. E.g.,

```
"kubernetesConfig": {
    "addons": [
        {
            "name": "kube-proxy-daemonset",
            "enabled": true,
            "data": "IGFwaVZlcnNpb246IHYxCmtpbmQ6IFNlcnZpY2UKbWV0YWRhdGE6CiAgbmFtZToga3ViZS1kbnMKICBuYW1lc3BhY2U6IGt1YmUtc3lzdGVtCiAgbGFiZWxzOgogICAgazhzLWFwcDoga3ViZS1kbnMKICAgIGt1YmVybmV0ZXMuaW8vY2x1c3Rlci1zZXJ2aWNlOiAidHJ1ZSIKICAgIGFkZG9ubWFuYWdlci5rdWJlcm5ldGVzLmlvL21vZGU6IFJlY29uY2lsZQogICAga3ViZXJuZXRlcy5pby9uYW1lOiAiS3ViZUROUyIKc3BlYzoKICBzZWxlY3RvcjoKICAgIGs4cy1hcHA6IGt1YmUtZG5zCiAgY2x1c3RlcklQOiA8a3ViZUROU1NlcnZpY2VJUD4KICBwb3J0czoKICAtIG5hbWU6IGRucwogICAgcG9ydDogNTMKICAgIHByb3RvY29sOiBVRFAKICAtIG5hbWU6IGRucy10Y3AKICAgIHBvcnQ6IDUzCiAgICBwcm90b2NvbDogVENQCi0tLQphcGlWZXJzaW9uOiB2MQpraW5kOiBTZXJ2aWNlQWNjb3VudAptZXRhZGF0YToKICBuYW1lOiBrdWJlLWRucwogIG5hbWVzcGFjZToga3ViZS1zeXN0ZW0KICBsYWJlbHM6CiAgICBrdWJlcm5ldGVzLmlvL2NsdXN0ZXItc2VydmljZTogInRydWUiCiAgICBhZGRvbm1hbmFnZXIua3ViZXJuZXRlcy5pby9tb2RlOiBSZWNvbmNpbGUKLS0tCmFwaVZlcnNpb246IHYxCmtpbmQ6IENvbmZpZ01hcAptZXRhZGF0YToKICBuYW1lOiBrdWJlLWRucwogIG5hbWVzcGFjZToga3ViZS1zeXN0ZW0KICBsYWJlbHM6CiAgICBhZGRvbm1hbmFnZXIua3ViZXJuZXRlcy5pby9tb2RlOiBFbnN1cmVFeGlzdHMKLS0tCmFwaVZlcnNpb246IGV4dGVuc2lvbnMvdjFiZXRhMQpraW5kOiBEZXBsb3ltZW50Cm1ldGFkYXRhOgogIG5hbWU6IGt1YmUtZG5zCiAgbmFtZXNwYWNlOiBrdWJlLXN5c3RlbQogIGxhYmVsczoKICAgIGs4cy1hcHA6IGt1YmUtZG5zCiAgICBrdWJlcm5ldGVzLmlvL2NsdXN0ZXItc2VydmljZTogInRydWUiCiAgICBhZGRvbm1hbmFnZXIua3ViZXJuZXRlcy5pby9tb2RlOiBSZWNvbmNpbGUKc3BlYzoKICAjIHJlcGxpY2FzOiBub3Qgc3BlY2lmaWVkIGhlcmU6CiAgIyAxLiBJbiBvcmRlciB0byBtYWtlIEFkZG9uIE1hbmFnZXIgZG8gbm90IHJlY29uY2lsZSB0aGlzIHJlcGxpY2FzIHBhcmFtZXRlci4KICAjIDIuIERlZmF1bHQgaXMgMS4KICAjIDMuIFdpbGwgYmUgdHVuZWQgaW4gcmVhbCB0aW1lIGlmIEROUyBob3Jpem9udGFsIGF1dG8tc2NhbGluZyBpcyB0dXJuZWQgb24uCiAgc3RyYXRlZ3k6CiAgICByb2xsaW5nVXBkYXRlOgogICAgICBtYXhTdXJnZTogMTAlCiAgICAgIG1heFVuYXZhaWxhYmxlOiAwCiAgc2VsZWN0b3I6CiAgICBtYXRjaExhYmVsczoKICAgICAgazhzLWFwcDoga3ViZS1kbnMKICB0ZW1wbGF0ZToKICAgIG1ldGFkYXRhOgogICAgICBsYWJlbHM6CiAgICAgICAgazhzLWFwcDoga3ViZS1kbnMKICAgICAgYW5ub3RhdGlvbnM6CiAgICAgICAgc2NoZWR1bGVyLmFscGhhLmt1YmVybmV0ZXMuaW8vY3JpdGljYWwtcG9kOiAnJwogICAgICAgIHNlY2NvbXAuc2VjdXJpdHkuYWxwaGEua3ViZXJuZXRlcy5pby9wb2Q6ICdkb2NrZXIvZGVmYXVsdCcKICAgIHNwZWM6CiAgICAgIHByaW9yaXR5Q2xhc3NOYW1lOiBzeXN0ZW0tY2x1c3Rlci1jcml0aWNhbAogICAgICB0b2xlcmF0aW9uczoKICAgICAgLSBrZXk6ICJDcml0aWNhbEFkZG9uc09ubHkiCiAgICAgICAgb3BlcmF0b3I6ICJFeGlzdHMiCiAgICAgIHZvbHVtZXM6CiAgICAgIC0gbmFtZToga3ViZS1kbnMtY29uZmlnCiAgICAgICAgY29uZmlnTWFwOgogICAgICAgICAgbmFtZToga3ViZS1kbnMKICAgICAgICAgIG9wdGlvbmFsOiB0cnVlCiAgICAgIGNvbnRhaW5lcnM6CiAgICAgIC0gbmFtZToga3ViZWRucwogICAgICAgIGltYWdlOiA8a3ViZXJuZXRlc0t1YmVETlNTcGVjPgogICAgICAgIHJlc291cmNlczoKICAgICAgICAgICMgVE9ETzogU2V0IG1lbW9yeSBsaW1pdHMgd2hlbiB3ZSd2ZSBwcm9maWxlZCB0aGUgY29udGFpbmVyIGZvciBsYXJnZQogICAgICAgICAgIyBjbHVzdGVycywgdGhlbiBzZXQgcmVxdWVzdCA9IGxpbWl0IHRvIGtlZXAgdGhpcyBjb250YWluZXIgaW4KICAgICAgICAgICMgZ3VhcmFudGVlZCBjbGFzcy4gQ3VycmVudGx5LCB0aGlzIGNvbnRhaW5lciBmYWxscyBpbnRvIHRoZQogICAgICAgICAgIyAiYnVyc3RhYmxlIiBjYXRlZ29yeSBzbyB0aGUga3ViZWxldCBkb2Vzbid0IGJhY2tvZmYgZnJvbSByZXN0YXJ0aW5nIGl0LgogICAgICAgICAgbGltaXRzOgogICAgICAgICAgICBtZW1vcnk6IDE3ME1pCiAgICAgICAgICByZXF1ZXN0czoKICAgICAgICAgICAgY3B1OiAxMDBtCiAgICAgICAgICAgIG1lbW9yeTogNzBNaQogICAgICAgIGxpdmVuZXNzUHJvYmU6CiAgICAgICAgICBodHRwR2V0OgogICAgICAgICAgICBwYXRoOiAvaGVhbHRoY2hlY2sva3ViZWRucwogICAgICAgICAgICBwb3J0OiAxMDA1NAogICAgICAgICAgICBzY2hlbWU6IEhUVFAKICAgICAgICAgIGluaXRpYWxEZWxheVNlY29uZHM6IDYwCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgICAgc3VjY2Vzc1RocmVzaG9sZDogMQogICAgICAgICAgZmFpbHVyZVRocmVzaG9sZDogNQogICAgICAgIHJlYWRpbmVzc1Byb2JlOgogICAgICAgICAgaHR0cEdldDoKICAgICAgICAgICAgcGF0aDogL3JlYWRpbmVzcwogICAgICAgICAgICBwb3J0OiA4MDgxCiAgICAgICAgICAgIHNjaGVtZTogSFRUUAogICAgICAgICAgIyB3ZSBwb2xsIG9uIHBvZCBzdGFydHVwIGZvciB0aGUgS3ViZXJuZXRlcyBtYXN0ZXIgc2VydmljZSBhbmQKICAgICAgICAgICMgb25seSBzZXR1cCB0aGUgL3JlYWRpbmVzcyBIVFRQIHNlcnZlciBvbmNlIHRoYXQncyBhdmFpbGFibGUuCiAgICAgICAgICBpbml0aWFsRGVsYXlTZWNvbmRzOiAzCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgIGFyZ3M6CiAgICAgICAgLSAtLWRvbWFpbj08a3ViZXJuZXRlc0t1YmVsZXRDbHVzdGVyRG9tYWluPi4KICAgICAgICAtIC0tZG5zLXBvcnQ9MTAwNTMKICAgICAgICAtIC0tY29uZmlnLWRpcj0va3ViZS1kbnMtY29uZmlnCiAgICAgICAgLSAtLXY9MgogICAgICAgIGVudjoKICAgICAgICAtIG5hbWU6IFBST01FVEhFVVNfUE9SVAogICAgICAgICAgdmFsdWU6ICIxMDA1NSIKICAgICAgICBwb3J0czoKICAgICAgICAtIGNvbnRhaW5lclBvcnQ6IDEwMDUzCiAgICAgICAgICBuYW1lOiBkbnMtbG9jYWwKICAgICAgICAgIHByb3RvY29sOiBVRFAKICAgICAgICAtIGNvbnRhaW5lclBvcnQ6IDEwMDUzCiAgICAgICAgICBuYW1lOiBkbnMtdGNwLWxvY2FsCiAgICAgICAgICBwcm90b2NvbDogVENQCiAgICAgICAgLSBjb250YWluZXJQb3J0OiAxMDA1NQogICAgICAgICAgbmFtZTogbWV0cmljcwogICAgICAgICAgcHJvdG9jb2w6IFRDUAogICAgICAgIHZvbHVtZU1vdW50czoKICAgICAgICAtIG5hbWU6IGt1YmUtZG5zLWNvbmZpZwogICAgICAgICAgbW91bnRQYXRoOiAva3ViZS1kbnMtY29uZmlnCiAgICAgIC0gbmFtZTogZG5zbWFzcQogICAgICAgIGltYWdlOiA8a3ViZXJuZXRlc0ROU01hc3FTcGVjPgogICAgICAgIGxpdmVuZXNzUHJvYmU6CiAgICAgICAgICBodHRwR2V0OgogICAgICAgICAgICBwYXRoOiAvaGVhbHRoY2hlY2svZG5zbWFzcQogICAgICAgICAgICBwb3J0OiAxMDA1NAogICAgICAgICAgICBzY2hlbWU6IEhUVFAKICAgICAgICAgIGluaXRpYWxEZWxheVNlY29uZHM6IDYwCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgICAgc3VjY2Vzc1RocmVzaG9sZDogMQogICAgICAgICAgZmFpbHVyZVRocmVzaG9sZDogNQogICAgICAgIGFyZ3M6CiAgICAgICAgLSAtdj0yCiAgICAgICAgLSAtbG9ndG9zdGRlcnIKICAgICAgICAtIC1jb25maWdEaXI9L2V0Yy9rOHMvZG5zL2Ruc21hc3EtbmFubnkKICAgICAgICAtIC1yZXN0YXJ0RG5zbWFzcT10cnVlCiAgICAgICAgLSAtLQogICAgICAgIC0gLWsKICAgICAgICAtIC0tY2FjaGUtc2l6ZT0xMDAwCiAgICAgICAgLSAtLW5vLW5lZ2NhY2hlCiAgICAgICAgLSAtLWxvZy1mYWNpbGl0eT0tCiAgICAgICAgLSAtLXNlcnZlcj0vY2x1c3Rlci5sb2NhbC8xMjcuMC4wLjEjMTAwNTMKICAgICAgICAtIC0tc2VydmVyPS9pbi1hZGRyLmFycGEvMTI3LjAuMC4xIzEwMDUzCiAgICAgICAgLSAtLXNlcnZlcj0vaXA2LmFycGEvMTI3LjAuMC4xIzEwMDUzCiAgICAgICAgcG9ydHM6CiAgICAgICAgLSBjb250YWluZXJQb3J0OiA1MwogICAgICAgICAgbmFtZTogZG5zCiAgICAgICAgICBwcm90b2NvbDogVURQCiAgICAgICAgLSBjb250YWluZXJQb3J0OiA1MwogICAgICAgICAgbmFtZTogZG5zLXRjcAogICAgICAgICAgcHJvdG9jb2w6IFRDUAogICAgICAgICMgc2VlOiBodHRwczovL2dpdGh1Yi5jb20va3ViZXJuZXRlcy9rdWJlcm5ldGVzL2lzc3Vlcy8yOTA1NSBmb3IgZGV0YWlscwogICAgICAgIHJlc291cmNlczoKICAgICAgICAgIHJlcXVlc3RzOgogICAgICAgICAgICBjcHU6IDE1MG0KICAgICAgICAgICAgbWVtb3J5OiAyME1pCiAgICAgICAgdm9sdW1lTW91bnRzOgogICAgICAgIC0gbmFtZToga3ViZS1kbnMtY29uZmlnCiAgICAgICAgICBtb3VudFBhdGg6IC9ldGMvazhzL2Rucy9kbnNtYXNxLW5hbm55CiAgICAgIC0gbmFtZTogc2lkZWNhcgogICAgICAgIGltYWdlOiA8a3ViZXJuZXRlc0ROU1NpZGVjYXJTcGVjPgogICAgICAgIGxpdmVuZXNzUHJvYmU6CiAgICAgICAgICBodHRwR2V0OgogICAgICAgICAgICBwYXRoOiAvbWV0cmljcwogICAgICAgICAgICBwb3J0OiAxMDA1NAogICAgICAgICAgICBzY2hlbWU6IEhUVFAKICAgICAgICAgIGluaXRpYWxEZWxheVNlY29uZHM6IDYwCiAgICAgICAgICB0aW1lb3V0U2Vjb25kczogNQogICAgICAgICAgc3VjY2Vzc1RocmVzaG9sZDogMQogICAgICAgICAgZmFpbHVyZVRocmVzaG9sZDogNQogICAgICAgIGFyZ3M6CiAgICAgICAgLSAtLXY9MgogICAgICAgIC0gLS1sb2d0b3N0ZGVycgogICAgICAgIC0gLS1wcm9iZT1rdWJlZG5zLDEyNy4wLjAuMToxMDA1MyxrdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLjxrdWJlcm5ldGVzS3ViZWxldENsdXN0ZXJEb21haW4+LDUsU1JWCiAgICAgICAgLSAtLXByb2JlPWRuc21hc3EsMTI3LjAuMC4xOjUzLGt1YmVybmV0ZXMuZGVmYXVsdC5zdmMuPGt1YmVybmV0ZXNLdWJlbGV0Q2x1c3RlckRvbWFpbj4sNSxTUlYKICAgICAgICBwb3J0czoKICAgICAgICAtIGNvbnRhaW5lclBvcnQ6IDEwMDU0CiAgICAgICAgICBuYW1lOiBtZXRyaWNzCiAgICAgICAgICBwcm90b2NvbDogVENQCiAgICAgICAgcmVzb3VyY2VzOgogICAgICAgICAgcmVxdWVzdHM6CiAgICAgICAgICAgIG1lbW9yeTogMjBNaQogICAgICAgICAgICBjcHU6IDEwbQogICAgICBkbnNQb2xpY3k6IERlZmF1bHQgICMgRG9uJ3QgdXNlIGNsdXN0ZXIgRE5TLgogICAgICBzZXJ2aWNlQWNjb3VudE5hbWU6IGt1YmUtZG5zCiAgICAgIG5vZGVTZWxlY3RvcjoKICAgICAgICBiZXRhLmt1YmVybmV0ZXMuaW8vb3M6IGxpbnV4"
        }
    ]
}
```

The reason for the unsightly gzip+base64 encoded input type is to optimize delivery payload, and to squash a human-maintainable yaml file representation into something that can be tightly pasted into a JSON string value without the arguably more unsightly carriage returns / whitespace that would be delivered with a literal copy/paste of a Kubernetes manifest.

Finally, the `addons.enabled` boolean property was omitted above; that's by design. If you specify a `containers` configuration, aks-engine assumes you're enabling the addon. The very first example above demonstrates a simple "enable this addon with default configuration" declaration.

#### External Custom YAML scripts

External YAML scripts can be configured for these supported addons and the manifest files for kube-scheduler, kube-controller-manager, cloud-controller-manager, kube-apiserver and PodSecurityPolicy. For addons, you will need to pass in a _base64_ encoded string of the kubernetes addon YAML file that you wish to use to `addons.Data` property. When `addons.Data` is provided with a value, the `containers` and `config` are required to be empty.

CAVEAT: Please note that this is an experimental feature. Since Addons.Data allows you to provide your own scripts, you face the risk of any unintended/undesirable consequences of the errors and failures from running that script.

```
"kubernetesConfig": {
    "addons": [
        {
            "name": "kube-proxy-daemonset",
            "enabled" : true,
            "data" : <base64 encoded string of your k8s addon YAML>,
        }
    ]
}
```
For kubernetes component manifests, you will need to pass in a _base64_ encoded string of the kubernetes manifest YAML file to _KubernetesComponentConfig["data"]_ . For example, to pass a custom kube-scheduler config, do the following:

```
"kubernetesConfig": {
    "schedulerConfig": {
            "data" : "<base64 encoded string of your k8s manifest YAML>"
        }
}
```
and to pass a custom pod security policy config, do the following:

```
"kubernetesConfig": {
    "podSecurityPolicyConfig": {
            "data" : "<base64 encoded string of your k8s manifest YAML>"
        }
}
```
<a name="feat-kubelet-config"></a>

#### kubeletConfig

`kubeletConfig` declares runtime configuration for the kubelet running on all master and agent nodes. It is a generic key/value object, and a child property of `kubernetesConfig`. An example custom kubelet config:

```
"kubernetesConfig": {
    "kubeletConfig": {
        "--eviction-hard": "memory.available<250Mi,nodefs.available<20%,nodefs.inodesFree<10%"
    }
}
```

See [here](https://kubernetes.io/docs/reference/generated/kubelet/) for a reference of supported kubelet options.

Below is a list of kubelet options that aks-engine will configure by default:

| kubelet option                      | default value                                                                                                                                                 |
| ----------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| "--cloud-config"                    | "/etc/kubernetes/azure.json"                                                                                                                                  |
| "--cloud-provider"                  | "azure"                                                                                                                                                       |
| "--cluster-domain"                  | "cluster.local"                                                                                                                                               |
| "--pod-infra-container-image"       | "pause-amd64:_version_"                                                                                                                                       |
| "--max-pods"                        | "30", or "110" if using kubenet --network-plugin (i.e., `"networkPlugin": "kubenet"`)                                                                         |
| "--eviction-hard"                   | "memory.available<100Mi,nodefs.available<10%,nodefs.inodesFree<5%"                                                                                            |
| "--node-status-update-frequency"    | "10s"                                                                                                                                                         |
| "--image-gc-high-threshold"         | "85"                                                                                                                                                          |
| "--image-gc-low-threshold"          | "850"                                                                                                                                                         |
| "--non-masquerade-cidr"             | "10.0.0.0/8"                                                                                                                                                  |
| "--azure-container-registry-config" | "/etc/kubernetes/azure.json"                                                                                                                                  |
| "--pod-max-pids"                    | "100" (need to activate the feature in --feature-gates=SupportPodPidsLimit=true)                                                                              |
| "--image-pull-progress-deadline"    | "30m"                                                                                                                                                         |
| "--feature-gates"                   | No default (can be a comma-separated list). On agent nodes `Accelerators=true` will be applied in the `--feature-gates` option for k8s versions before 1.11.0 |

Below is a list of kubelet options that are _not_ currently user-configurable, either because a higher order configuration vector is available that enforces kubelet configuration, or because a static configuration is required to build a functional cluster:

| kubelet option                               | default value                                    |
| -------------------------------------------- | ------------------------------------------------ |
| "--address"                                  | "0.0.0.0"                                        |
| "--allow-privileged"                         | "true"                                           |
| "--pod-manifest-path"                        | "/etc/kubernetes/manifests"                      |
| "--network-plugin"                           | "cni"                                            |
| "--node-labels"                              | (based on Azure node metadata)                   |
| "--cgroups-per-qos"                          | "true"                                           |
| "--enforce-node-allocatable"                 | "pods"                                           |
| "--kubeconfig"                               | "/var/lib/kubelet/kubeconfig"                    |
| "--register-node" (master nodes only)        | "true"                                           |
| "--register-with-taints" (master nodes only) | "node-role.kubernetes.io/master=true:NoSchedule" |
| "--keep-terminated-pod-volumes"              | "false"                                          |

<a name="feat-controller-manager-config"></a>

#### controllerManagerConfig

`controllerManagerConfig` declares runtime configuration for the kube-controller-manager daemon running on all master nodes. Like `kubeletConfig` it is a generic key/value object, and a child property of `kubernetesConfig`. An example custom controller-manager config:

```
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

Below is a list of controller-manager options that aks-engine will configure by default:

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
| "--cluster-name"                     | _auto-generated using api model properties_             |
| "--cloud-provider"                   | "azure"                                                 |
| "--cloud-config"                     | "/etc/kubernetes/azure.json"                            |
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

```
"kubernetesConfig": {
    "cloudControllerManagerConfig": {
        "--route-reconciliation-period": "1m"
    }
}
```

See [here](https://kubernetes.io/docs/reference/generated/cloud-controller-manager/) for a reference of supported controller-manager options.

Below is a list of cloud-controller-manager options that aks-engine will configure by default:

| controller-manager option       | default value |
| ------------------------------- | ------------- |
| "--route-reconciliation-period" | "10s"         |

Below is a list of cloud-controller-manager options that are _not_ currently user-configurable, either because a higher order configuration vector is available that enforces controller-manager configuration, or because a static configuration is required to build a functional cluster:

| controller-manager option | default value                               |
| ------------------------- | ------------------------------------------- |
| "--kubeconfig"            | "/var/lib/kubelet/kubeconfig"               |
| "--allocate-node-cidrs"   | "false"                                     |
| "--cluster-cidr"          | _uses clusterSubnet value_                  |
| "--cluster-name"          | _auto-generated using api model properties_ |
| "--cloud-provider"        | "azure"                                     |
| "--cloud-config"          | "/etc/kubernetes/azure.json"                |
| "--leader-elect"          | "true"                                      |
| "--v"                     | "2"                                         |

<a name="feat-apiserver-config"></a>

#### apiServerConfig

`apiServerConfig` declares runtime configuration for the kube-apiserver daemon running on all master nodes. Like `kubeletConfig` and `controllerManagerConfig` it is a generic key/value object, and a child property of `kubernetesConfig`. An example custom apiserver config:

```
"kubernetesConfig": {
    "apiServerConfig": {
        "--request-timeout": "30s"
    }
}
```

Or perhaps you want to customize/override the set of admission-control flags passed to the API Server by default, you can omit the options you don't want and specify only the ones you need as follows:

```
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

Below is a list of apiserver options that aks-engine will configure by default:

| apiserver option                | default value                                                                                                                                                                                                                           |
| ------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| "--admission-control"           | "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota" (Kubernetes versions prior to 1.9.0)                                                                               |
| "--enable-admission-plugins"`*` | "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,ExtendedResourceToleration" (Kubernetes versions 1.9.0 and later) |
| "--authorization-mode"          | "Node", "RBAC" (_the latter if enabledRbac is true_)                                                                                                                                                                                    |
| "--audit-log-maxage"            | "30"                                                                                                                                                                                                                                    |
| "--audit-log-maxbackup"         | "10"                                                                                                                                                                                                                                    |
| "--audit-log-maxsize"           | "100"                                                                                                                                                                                                                                   |
| "--feature-gates"               | No default (can be a comma-separated list)                                                                                                                                                                                              |
| "--oidc-username-claim"         | "oid" (_if has AADProfile_)                                                                                                                                                                                                             |
| "--oidc-groups-claim"           | "groups" (_if has AADProfile_)                                                                                                                                                                                                          |
| "--oidc-client-id"              | _calculated value that represents OID client ID_ (_if has AADProfile_)                                                                                                                                                                  |
| "--oidc-issuer-url"             | _calculated value that represents OID issuer URL_ (_if has AADProfile_)                                                                                                                                                                 |

`*` In Kubernetes versions 1.10.0 and later the `--admission-control` flag is deprecated and `--enable-admission-plugins` is used in its stead.

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
| "--repair-malformed-updates"                | "false"                                                                                 |
| "--tls-cert-file"                           | "/etc/kubernetes/certs/apiserver.crt"                                                   |
| "--tls-private-key-file"                    | "/etc/kubernetes/certs/apiserver.key"                                                   |
| "--client-ca-file"                          | "/etc/kubernetes/certs/ca.crt"                                                          |
| "--service-account-key-file"                | "/etc/kubernetes/certs/apiserver.key"                                                   |
| "--kubelet-client-certificate"              | "/etc/kubernetes/certs/client.crt"                                                      |
| "--kubelet-client-key"                      | "/etc/kubernetes/certs/client.key"                                                      |
| "--service-cluster-ip-range"                | _see serviceCIDR_                                                                       |
| "--storage-backend"                         | _calculated value that represents etcd version_                                         |
| "--v"                                       | "4"                                                                                     |
| "--experimental-encryption-provider-config" | "/etc/kubernetes/encryption-config.yaml" (_if enableDataEncryptionAtRest is true_)      |
| "--experimental-encryption-provider-config" | "/etc/kubernetes/encryption-config.yaml" (_if enableEncryptionWithExternalKms is true_) |
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

```
"kubernetesConfig": {
    "schedulerConfig": {
        "--v": "2"
    }
}
```

See [here](https://kubernetes.io/docs/reference/generated/kube-scheduler/) for a reference of supported kube-scheduler options.

Below is a list of scheduler options that aks-engine will configure by default:

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

<a name="feat-private-cluster"></a>

#### privateCluster

`privateCluster` defines a cluster without public addresses assigned. It is a child property of `kubernetesConfig`.

| Name           | Required | Description                                                                                                                                          |
| -------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| enabled        | no       | Enable [Private Cluster](./kubernetes/features.md/#feat-private-cluster) (boolean - default == false)                                                |
| jumpboxProfile | no       | Configure and auto-provision a jumpbox to access your private cluster. `jumpboxProfile` is ignored if enabled is `false`. See `jumpboxProfile` below |

#### jumpboxProfile

`jumpboxProfile` describes the settings for a jumpbox deployed via aks-engine to access a private cluster. It is a child property of `privateCluster`.

| Name           | Required | Description                                                                                                                                                                        |
| -------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name           | yes      | This is the unique name for the jumpbox VM. Some resources deployed with the jumpbox are derived from this name                                                                    |
| vmSize         | yes      | Describes a valid [Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/)                                                       |
| publicKey      | yes      | The public SSH key used for authenticating access to the jumpbox. Here are instructions for [generating a public/private key pair](ssh.md#ssh-key-generation)                      |
| osDiskSizeGB   | no       | Describes the OS Disk Size in GB. Defaults to `30`                                                                                                                                 |
| storageProfile | no       | Specifies the storage profile to use. Valid values are [ManagedDisks](../examples/disks-managed) or [StorageAccount](../examples/disks-storageaccount). Defaults to `ManagedDisks` |
| username       | no       | Describes the admin username to be used on the jumpbox. Defaults to `azureuser`                                                                                                    |

### masterProfile

`masterProfile` describes the settings for master configuration.

| Name                         | Required                                  | Description                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---------------------------- | ----------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| count                        | yes                                       | Masters have count value of 1, 3, or 5 masters                                                                                                                                                                                                                                                                                                                                                                             |
| dnsPrefix                    | yes                                       | The dns prefix for the master FQDN. The master FQDN is used for SSH or commandline access. This must be a unique name. ([bring your own VNET examples](../examples/vnet))                                                                                                                                                                                                                                                  |
| subjectAltNames              | no                                        | An array of fully qualified domain names using which a user can reach API server. These domains are added as Subject Alternative Names to the generated API server certificate. **NOTE**: These domains **will not** be automatically provisioned.                                                                                                                                                                         |
| firstConsecutiveStaticIP     | only required when vnetSubnetId specified and when MasterProfile is not `VirtualMachineScaleSets`  | The IP address of the first master. IP Addresses will be assigned consecutively to additional master nodes. When MasterProfile is using `VirtualMachineScaleSets`, this value will be determined by an offset from the first IP in the `vnetCidr`. For example, if `vnetCidr` is `10.239.0.0/16`, then `firstConsecutiveStaticIP` will be `10.239.0.4`                                                                                                                                                                                                                                                                                                                 |
| vmsize                       | yes                                       | Describes a valid [Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/). These are restricted to machines with at least 2 cores and 100GB of ephemeral disk space                                                                                                                                                                                                     |
| osDiskSizeGB                 | no                                        | Describes the OS Disk Size in GB                                                                                                                                                                                                                                                                                                                                                                                           |
| vnetSubnetId                 | only required when using custom VNET                                        | Specifies the Id of an alternate VNET subnet. The subnet id must specify a valid VNET ID owned by the same subscription. ([bring your own VNET examples](../examples/vnet)). When MasterProfile is set to `VirtualMachineScaleSets`, this value should be the subnetId of the master subnet. When MasterProfile is set to `AvailabilitySet`, this value should be the subnetId shared by both master and agent nodes.                                                                                                                                                                                                                                               |
| extensions                   | no                                        | This is an array of extensions. This indicates that the extension be run on a single master. The name in the extensions array must exactly match the extension name in the extensionProfiles                                                                                                                                                                                                                               |
| vnetCidr                     | no                                        | Specifies the VNET cidr when using a custom VNET ([bring your own VNET examples](../examples/vnet)). This VNET cidr should include both the master and the agent subnets.                                                                                                                                                                                                                                                                                                                        |
| imageReference.name          | no                                        | The name of the Linux OS image. Needs to be used in conjunction with resourceGroup, below                                                                                                                                                                                                                                                                                                                                  |
| imageReference.resourceGroup | no                                        | Resource group that contains the Linux OS image. Needs to be used in conjunction with name, above                                                                                                                                                                                                                                                                                                                          |
| distro                       | no                                        | Specifies the masters' Linux distribution. Currently supported values are: `ubuntu`, `aks`, `aks-docker-engine` and `coreos` (CoreOS support is currently experimental - [Example of CoreOS Master with CoreOS Agents](../examples/coreos/kubernetes-coreos.json)). For Azure Public Cloud, defaults to `aks` if undefined, unless GPU nodes are present, in which case it will default to `aks-docker-engine`. For Sovereign Clouds, the default is `ubuntu`. `aks` is a custom image based on `ubuntu` that comes with pre-installed software necessary for Kubernetes deployments (Azure Public Cloud only for now). **NOTE**: GPU nodes are currently incompatible with the default Moby container runtime provided in the `aks` image. Clusters containing GPU nodes will be set to use the `aks-docker-engine` distro which is functionally equivalent to `aks` with the exception of the docker distribution (see [GPU support Walkthrough](kubernetes/gpu.md) for details). Currently supported OS and orchestrator configurations -- `ubuntu` and `aks`: DCOS, Docker Swarm, Kubernetes;  `coreos`: Kubernetes. [Example of CoreOS Master with CoreOS Agents](../examples/coreos/kubernetes-coreos.json) |
| customFiles                  | no                                        | The custom files to be provisioned to the master nodes. Defined as an array of json objects with each defined as `"source":"absolute-local-path", "dest":"absolute-path-on-masternodes"`.[See examples](../examples/customfiles)                                                                                                                                                                                           |
| availabilityProfile          | no                                                                   | Supported values are `AvailabilitySet` (default) and `VirtualMachineScaleSets` (still under development: upgrade not supported; requires Kubernetes clusters version 1.10+ and agent pool availabilityProfile must also be `VirtualMachineScaleSets`). When MasterProfile is using `VirtualMachineScaleSets`, to SSH into a master node, you need to use `ssh -p 50001` instead of port 22.                                                                                                                                                                                                                                                                                                                                                                                             |
| agentVnetSubnetId                 | only required when using custom VNET and when MasterProfile is using `VirtualMachineScaleSets`                                         | Specifies the Id of an alternate VNET subnet for all the agent pool nodes. The subnet id must specify a valid VNET ID owned by the same subscription. ([bring your own VNET examples](../examples/vnet)). When MasterProfile is using `VirtualMachineScaleSets`, this value should be the subnetId of the subnet for all agent pool nodes.                                                                                                                                                                                                                                                |
| [availabilityZones](../examples/kubernetes-zones/README.md)                    | no                                       | To protect your cluster from datacenter-level failures, you can enable the Availability Zones feature for your cluster by configuring `"availabilityZones"` for the master profile and all of the agentPool profiles in the cluster definition. Check out [Availability Zones README](../examples/kubernetes-zones/README.md) for more details.                                                                                                                                                                                                                                                   |

### agentPoolProfiles

A cluster can have 0 to 12 agent pool profiles. Agent Pool Profiles are used for creating agents with different capabilities such as VMSizes, VMSS or Availability Set, Public/Private access, user-defined OS Images, [attached storage disks](../examples/disks-storageaccount), [attached managed disks](../examples/disks-managed), or [Windows](../examples/windows).

| Name                         | Required                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ---------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| availabilityProfile          | no                                                                   | Supported values are `VirtualMachineScaleSets` (default, except for Kubernetes clusters before version 1.10) and `AvailabilitySet`.                                                                                                                                                                                                                                                                                                                                                                                              |
| count                        | yes                                                                  | Describes the node count                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| [availabilityZones](../examples/kubernetes-zones/README.md)                    | no                                       | To protect your cluster from datacenter-level failures, you can enable the Availability Zones feature for your cluster by configuring `"availabilityZones"` for the master profile and all of the agentPool profiles in the cluster definition. Check out [Availability Zones README](../examples/kubernetes-zones/README.md) for more details.                                                                                                                                                                                                                                                   |
| singlePlacementGroup             | no                                                                   | Supported values are `true` (default) and `false`. Only applies to clusters with availabilityProfile `VirtualMachineScaleSets`. `true`: A VMSS with a single placement group and has a range of 0-100 VMs. `false`: A VMSS with multiple placement groups and has a range of 0-1,000 VMs. For more information, check out [virtual machine scale sets placement groups](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups).                                                                                                                                                                                                                           |
| scaleSetPriority             | no                                                                   | Supported values are `Regular` (default) and `Low`. Only applies to clusters with availabilityProfile `VirtualMachineScaleSets`. Enables the usage of [Low-priority VMs on Scale Sets](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-use-low-priority).                                                                                                                                                                                                                           |
| scaleSetEvictionPolicy       | no                                                                   | Supported values are `Delete` (default) and `Deallocate`. Only applies to clusters with availabilityProfile of `VirtualMachineScaleSets` and scaleSetPriority of `Low`.                                                                                                                                                                                                                                                                                                                                                          |
| diskSizesGB                  | no                                                                   | Describes an array of up to 4 attached disk sizes. Valid disk size values are between 1 and 1024                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| dnsPrefix                    | Required if agents are to be exposed publically with a load balancer | The dns prefix that forms the FQDN to access the loadbalancer for this agent pool. This must be a unique name among all agent pools. Not supported for Kubernetes clusters                                                                                                                                                                                                                                                                                                                                                       |
| name                         | yes                                                                  | This is the unique name for the agent pool profile. The resources of the agent pool profile are derived from this name                                                                                                                                                                                                                                                                                                                                                                                                           |
| ports                        | only required if needed for exposing services publically             | Describes an array of ports need for exposing publically. A tcp probe is configured for each port and only opens to an agent node if the agent node is listening on that port. A maximum of 150 ports may be specified. Not supported for Kubernetes clusters                                                                                                                                                                                                                                                                    |
| storageProfile               | no                                                                   | Specifies the storage profile to use. Valid values are [ManagedDisks](../examples/disks-managed) or [StorageAccount](../examples/disks-storageaccount). Defaults to `ManagedDisks`                                                                                                                                                                                                                                                                                                                                               |
| vmsize                       | yes                                                                  | Describes a valid [Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/). These are restricted to machines with at least 2 cores                                                                                                                                                                                                                                                                                                                                             |
| osDiskSizeGB                 | no                                                                   | Describes the OS Disk Size in GB                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| vnetSubnetId                 | no                                                                   | Specifies the Id of an alternate VNET subnet. The subnet id must specify a valid VNET ID owned by the same subscription. ([bring your own VNET examples](../examples/vnet))                                                                                                                                                                                                                                                                                                                                                      |
| imageReference.name          | no                                                                   | The name of a a Linux OS image. Needs to be used in conjunction with resourceGroup, below                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| imageReference.resourceGroup | no                                                                   | Resource group that contains the Linux OS image. Needs to be used in conjunction with name, above                                                                                                                                                                                                                                                                                                                                                                                                                                |
| osType                       | no                                                                   | Specifies the agent pool's Operating System. Supported values are `Windows` and `Linux`. Defaults to `Linux`                                                                                                                                                                                                                                                                                                                                                                                                                     |
| distro                       | no                                                                   | Specifies the agent pool's Linux distribution. Currently supported values are: `ubuntu`, `aks`, `aks-docker-engine` and `coreos` (CoreOS support is currently experimental - [Example of CoreOS Master with CoreOS Agents](../examples/coreos/kubernetes-coreos.json)). For Azure Public Cloud, defaults to `aks` if undefined, unless GPU nodes are present, in which case it will default to `aks-docker-engine`. For Sovereign Clouds, the default is `ubuntu`. `aks` is a custom image based on `ubuntu` that comes with pre-installed software necessary for Kubernetes deployments (Azure Public Cloud only for now). **NOTE**: GPU nodes are currently incompatible with the default Moby container runtime provided in the `aks` image. Clusters containing GPU nodes will be set to use the `aks-docker-engine` distro which is functionally equivalent to `aks` with the exception of the docker distribution (see [GPU support Walkthrough](kubernetes/gpu.md) for details). Currently supported OS and orchestrator configurations -- `ubuntu`: DCOS, Docker Swarm, Kubernetes; `coreos`: Kubernetes. [Example of CoreOS Master with Windows and Linux (CoreOS and Ubuntu) Agents](../examples/coreos/kubernetes-coreos-hybrid.json) |
| acceleratedNetworkingEnabled | no                                                                   | Use [Azure Accelerated Networking](https://azure.microsoft.com/en-us/blog/maximize-your-vm-s-performance-with-accelerated-networking-now-generally-available-for-both-windows-and-linux/) feature for Linux agents (You must select a VM SKU that supports Accelerated Networking). Defaults to `true` if the VM SKU selected supports Accelerated Networking                                                                                                                                                                                                                                                      |
| acceleratedNetworkingEnabledWindows | no                                                                   | Use [Azure Accelerated Networking](https://azure.microsoft.com/en-us/blog/maximize-your-vm-s-performance-with-accelerated-networking-now-generally-available-for-both-windows-and-linux/) feature for Windows agents (You must select a VM SKU that supports Accelerated Networking). Defaults to `false`                                                                                                                                                                                                                                                      |

### linuxProfile

`linuxProfile` provides the linux configuration for each linux node in the cluster

| Name                             | Required | Description                                                                                                                                                                      |
| -------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| adminUsername                    | yes      | Describes the username to be used on all linux clusters                                                                                                                          |
| ssh.publicKeys.keyData           | yes      | The public SSH key used for authenticating access to all Linux nodes in the cluster. Here are instructions for [generating a public/private key pair](ssh.md#ssh-key-generation) |
| secrets                          | no       | Specifies an array of key vaults to pull secrets from and what secrets to pull from each                                                                                         |
| customSearchDomain.name          | no       | describes the search domain to be used on all linux clusters                                                                                                                     |
| customSearchDomain.realmUser     | no       | describes the realm user with permissions to update dns registries on Windows Server DNS                                                                                         |
| customSearchDomain.realmPassword | no       | describes the realm user password to update dns registries on Windows Server DNS                                                                                                 |
| customNodesDNS.dnsServer         | no       | describes the IP address of the DNS Server                                                                                                                                       |

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

| Name                             | Required | Description                                                              |
| -------------------------------- | -------- | ------------------------------------------------------------------------ |
| adminUsername                    | yes      | Username for the Windows adminstrator account created on each Windows node |
| adminPassword                    | yes      | Password for the Windows adminstrator account created on each Windows node |
| windowsPublisher                 | no       | Publisher used to find Windows VM to deploy from marketplace. Default: `MicrosoftWindowsServer` |
| windowsOffer                     | no       | Offer used to find Windows VM to deploy from marketplace. Default: `WindowsServerSemiAnnual` |
| windowsSku                       | no       | SKU usedto find Windows VM to deploy from marketplace. Default: `Datacenter-Core-1809-with-Containers-smalldisk` |
| imageVersion                     | no       | Specific image version to deploy from marketplace.  Default: `latest` |
| windowsImageSourceURL            | no       | Path to an existing Azure storage blob with a sysprepped VHD. This is used to test pre-release or customized VHD files that you have uploaded to Azure. If provided, the above 4 parameters are ignored. |

#### Choosing a Windows version

If you want to choose a specific Windows image, but automatically use the latest - set `windowsPublisher`, `windowsOffer`, and `windowsSku`. If you need a specific version, then add `agentWindowsVersion` too.

You can find all available images with `az vm image list`


```bash
$ az vm image list --publisher MicrosoftWindowsServer --all -o table

Offer                    Publisher                      Sku                                             Urn                                                                                                            Version
-----------------------  -----------------------------  ----------------------------------------------  -------------------------------------------------------------------------------------------------------------  -----------------
...
WindowsServerSemiAnnual  MicrosoftWindowsServer         Datacenter-Core-1709-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServerSemiAnnual:Datacenter-Core-1709-with-Containers-smalldisk:1709.0.20181017  1709.0.20181017
WindowsServerSemiAnnual  MicrosoftWindowsServer         Datacenter-Core-1803-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServerSemiAnnual:Datacenter-Core-1803-with-Containers-smalldisk:1803.0.20181017  1803.0.20181017
WindowsServerSemiAnnual  MicrosoftWindowsServer         Datacenter-Core-1809-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServerSemiAnnual:Datacenter-Core-1809-with-Containers-smalldisk:1809.0.20181107  1809.0.20181107
WindowsServer            MicrosoftWindowsServer         2019-Datacenter-Core-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServer:2019-Datacenter-Core-with-Containers-smalldisk:2019.0.20181107            2019.0.20181107
```

If you wanted to use the last one in the list above, then set:

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

### servicePrincipalProfile

`servicePrincipalProfile` describes an Azure Service credentials to be used by the cluster for self-configuration. See [service principal](serviceprincipal.md) for more details on creation.

| Name                         | Required                          | Description                                                                                                 |
| ---------------------------- | --------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| clientId                     | yes, for Kubernetes clusters      | describes the Azure client id. It is recommended to use a separate client ID per cluster                    |
| secret                       | yes, for Kubernetes clusters      | describes the Azure client secret. It is recommended to use a separate client secret per client id          |
| objectId                     | optional, for Kubernetes clusters | describes the Azure service principal object id. It is required if enableEncryptionWithExternalKms is true  |
| keyvaultSecretRef.vaultId    | no, for Kubernetes clusters       | describes the vault id of the keyvault to retrieve the service principal secret from. See below for format. |
| keyvaultSecretRef.secretName | no, for Kubernetes clusters       | describes the name of the service principal secret in keyvault                                              |
| keyvaultSecretRef.version    | no, for Kubernetes clusters       | describes the version of the secret to use                                                                  |

format for `keyvaultSecretRef.vaultId`, can be obtained in cli, or found in the portal:
`/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>`. See [keyvault params](../examples/keyvault-params/README.md#service-principal-profile) for an example.

## Cluster Defintions for apiVersion "2016-03-30"

Here are the cluster definitions for apiVersion "2016-03-30". This matches the api version of the Azure Container Service Engine.

### apiVersion

| Name       | Required | Description                                                             |
| ---------- | -------- | ----------------------------------------------------------------------- |
| apiVersion | yes      | The version of the template. For "2016-03-30" the value is "2016-03-30" |

### orchestratorProfile

`orchestratorProfile` describes the orchestrator settings.

| Name             | Required | Description                                     |
| ---------------- | -------- | ----------------------------------------------- |
| orchestratorType | yes      | Specifies the orchestrator type for the cluster |

Here are the valid values for the orchestrator types:

1.  `DCOS` - this represents the [DC/OS orchestrator](dcos.md).
2.  `Swarm` - this represents the [Swarm orchestrator](swarm.md).
3.  `Kubernetes` - this represents the [Kubernetes orchestrator](kubernetes.md).
4.  `Swarm Mode` - this represents the [Swarm Mode orchestrator](swarmmode.md).

### masterProfile

`masterProfile` describes the settings for master configuration.

| Name      | Required | Description                                                                                                                                                                |
| --------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| count     | yes      | Masters have count value of 1, 3, or 5 masters                                                                                                                             |
| dnsPrefix | yes      | The dns prefix for the masters FQDN. The master FQDN is used for SSH or commandline access. This must be a unique name. ([bring your own VNET examples](../examples/vnet)) |

### agentPoolProfiles

For apiVersion "2016-03-30", a cluster may have only 1 agent pool profiles.

| Name      | Required                                                             | Description                                                                                                                                                                          |
| --------- | -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| count     | yes                                                                  | Describes the node count                                                                                                                                                             |
| dnsPrefix | required if agents are to be exposed publically with a load balancer | The dns prefix that forms the FQDN to access the loadbalancer for this agent pool. This must be a unique name among all agent pools                                                  |
| name      | yes                                                                  | The unique name for the agent pool profile. The resources of the agent pool profile are derived from this name                                                                       |
| vmsize    | yes                                                                  | Describes a valid [Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/). These are restricted to machines with at least 2 cores |

### linuxProfile

`linuxProfile` provides the linux configuration for each linux node in the cluster

| Name                      | Required | Description                                                                                                                                                                      |
| ------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| adminUsername             | yes      | Describes the username to be used on all linux clusters                                                                                                                          |
| ssh.publicKeys[0].keyData | yes      | The public SSH key used for authenticating access to all Linux nodes in the cluster. Here are instructions for [generating a public/private key pair](ssh.md#ssh-key-generation) |

### aadProfile

`aadProfile` provides [Azure Active Directory integration](kubernetes.aad.md) configuration for the cluster, currently only available for Kubernetes orchestrator.

| Name         | Required | Description                                                                                                                 |
| ------------ | -------- | --------------------------------------------------------------------------------------------------------------------------- |
| clientAppID  | yes      | Describes the client AAD application ID                                                                                     |
| serverAppID  | yes      | Describes the server AAD application ID                                                                                     |
| adminGroupID | no       | Describes the AAD Group Object ID that will be assigned the cluster-admin RBAC role                                         |
| tenantID     | no       | Describes the AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription |

### extensionProfiles

A cluster can have 0 - N extensions in extension profiles. Extension profiles allow a user to easily add pre-packaged functionality into a cluster. An example would be configuring a monitoring solution on your cluster. You can think of extensions like a marketplace for acs clusters.

| Name                | Required | Description                                                                                                                                                                      |
| ------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name                | yes      | The name of the extension. This has to exactly match the name of a folder under the extensions folder                                                                            |
| version             | yes      | The version of the extension. This has to exactly match the name of the folder under the extension name folder                                                                   |
| extensionParameters | optional | Extension parameters may be required by extensions. The format of the parameters is also extension dependant                                                                     |
| rootURL             | optional | URL to the root location of extensions. The rootURL must have an extensions child folder that follows the extensions convention. The rootURL is mainly used for testing purposes |

You can find more information, as well as a list of extensions on the [extensions documentation](extensions.md).
