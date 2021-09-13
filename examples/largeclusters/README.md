# AKS Engine - Large Clusters

## Overview

[This example](kubernetes.json) shows some recommended cluster configuration for building Kubernetes clusters with thousands of nodes using AKS Engine.

Note: this guidance is applicable for Linux only.

## Configuration recommendations

### Larger etcd storage size

Larger clusters are more likely to produce lots of cluster resources, which can add additional data storage requirements. The current recommended maximum for etcd's storage size limit is 8GB. Reference:

- https://etcd.io/docs/v3.3/dev-guide/limit/

To configure your cluster to have 8 GB of etcd data:

```
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "kubernetesConfig": {
        "etcdStorageLimitGB": 8,
        ....
      }
      ...
    },
    ...
  }
  ...
}
```

### Multiple cluster Load Balancer IP addresses for increased outbound SNAT connections

AKS Engine routes node outbound connectivity routes through a single, cluster load balancer. Azure Standard Load Balancer limits outbound SNAT ports to 64,000 per IP address. Active larger clusters can require more than that for regular network I/O, and to accommodate that, you may distribute your outbound SNAT among more than one IP address. Reference:

- https://docs.microsoft.com/en-us/azure/load-balancer/outbound-rules

As an example, to configure your cluster to use 5 SNAT IP addresses:

```
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "kubernetesConfig": {
        "loadBalancerOutboundIPs": 5,
        ....
      }
      ...
    },
    ...
  }
  ...
}
```

As a note: if you have a requirement that all outbound traffic share a single source IP address, then you must use the default configuration, which will create a Standard Load Balancer with 1 IP address. In this scenario, your cluster will have a hard limit of 64,000 SNAT ports.

### Use the cloud-controller-manager implementation of the Azure cloud provider

Also referred to as "out-of-tree", cloud-provider-azure code development is carried out in its own code repository, according to a separate release velocity than upstream Kubernetes. Reference:

- https://github.com/kubernetes-sigs/cloud-provider-azure/releases

The cloud-controller-manager implementation of cloud-provider-azure produces many runtime optimizations that optimize cluster behavior for running at scale. You should always use cloud-controller-manager when running large clusters:

```
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "kubernetesConfig": {
        "useCloudControllerManager": 5,
        ....
      }
      ...
    },
    ...
  }
  ...
}
```

### Tune kubelet max-pods for Azure CNI

AKS Engine only supports Azure CNI for clusters with over 400 nodes. Azure CNI pre-allocates IP addresses during node bootstrapping in accordance with that node's `--max-pods` kubelet configuration. As the number of nodes in your cluster grows, so will the number of allocated IP addresses in your cluster VNET. We must, therefore, configure any large cluster to not exceed the Azure-enforced limit of 65,336 private IP addresses per VNET. Reference:

- https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/azure-subscription-service-limits?toc=/azure/virtual-network/toc.json#networking-limits

We can do an example calculation to aid your own large cluster configuration.

The AKS Engine default `--max-pods` configuration for Azure CNI-backed clusters is 30, which will produce 30 pre-allocated private IP addresses as soon as a new node comes online in the cluster VNET's IP address space, to be used as pod IP addresses in the CNI network CIDR. If we add 1 more IP address to account for the `eth0` interface running on the VM's host OS, then we can easily calculate the maximum number of nodes a cluster can support using the default `"--max-pods": "30"` configuration:

```
>>> 65336 / 30
2177
```

The above quotient of 2177 includes control plane (master) nodes, which are also configured to use the same CNI network address space, so if we assume we're running 5 control plane nodes, then we should expect the absolute node count upper boundary to be 2172.

If you want to anticipate that your cluster will grow beyond roughly 2000 nodes, then you will need to provide a lower `--max-pods` configuration so that the cluster VNET IP address space does not fill up as quickly as your cluster node count grows. For example, to target 5000 nodes, you would divide the limit (65336) by 5000, and then subtract 1 from the quotient:

```
>>> (65336 / 5000) - 1
12
```

To use that configuration before you build your AKS Engine cluster:

```
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "kubernetesConfig": {
        "kubeletConfig": {
          "--max-pods": "12"
        },
        ....
      }
      ...
    },
    "masterProfile": {
      "kubernetesConfig": {
        "kubeletConfig": {
          "--max-pods": "24"
        }
      }
    },
    ...
  }
  ...
}
```

A note on the above example: we have configured `12` as the `--max-pods` value for the cluster over all, which will be used for each node pool that you create; but we used a larger value for the `masterProfile`, which configures the control plane. The reason we want to do that is to leave reasonable overhead for system pods running on the control plane. If you consider the standard Kubernetes system components: `apiserver`, `controller-manager`, `cloud-controller-manager`, `scheduler`, `kube-proxy`, you are already approaching the limit of `12`. To see a real-world example, here's a cluster built with all the recommended large cluster configurations:

```
$ kubectl get pods -n kube-system -o wide | grep 'k8s-master-40237885-0'
azure-ip-masq-agent-slrtr                        1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
cloud-controller-manager-k8s-master-40237885-0   1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
cloud-node-manager-68jlj                         1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
coredns-787d7f6757-clm85                         1/1     Running   0               5m    10.240.0.102     k8s-master-40237885-0                <none>           <none>
csi-azuredisk-node-5ghtr                         3/3     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
csi-azurefile-node-rrtz5                         3/3     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
kube-addon-manager-k8s-master-40237885-0         1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
kube-apiserver-k8s-master-40237885-0             1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
kube-controller-manager-k8s-master-40237885-0    1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
kube-proxy-zl6xp                                 1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
kube-scheduler-k8s-master-40237885-0             1/1     Running   0               5m    10.255.255.5     k8s-master-40237885-0                <none>           <none>
$ kubectl get pods -n kube-system -o wide | grep 'k8s-master-40237885-0' | wc -l
      11
```

As you can see, the default cluster configuration requires that a control plane VM runs 11 pods right after the cluster is created! This means that if we configure `--max-pods` to be `12` for control plane VMs, then little to no pod scheduling availability will be possible throughout the lifecycle of the cluster. Instead we configure the control plane VMs with more pod overhead, using a `"--max-pods": "24"` configuration.

Another note on the integration of Azure CNI with Kubernetes: we mentioned above that Azure CNI requires knowledge of the maximum number of pods that will ever run on the node that it is being installed onto, in order to pre-allocate the appropriate number of possibly needed IP addresses. Unfortunately, this means two undesirable side-effects:

1. Azure CNI will always pre-allocate at least one more IP address than it needs (for `kube-proxy`), and for AKS Engine-created clusters that number is in fact 3:

- `kube-proxy`
- `azure-ip-masq-agent`
- `cloud-node-manager`

If you include CSI drivers that number jumps to 5!:

- `csi-azuredisk`
- `csi-azurefile`

The reason for classifying the above 5 pods as taking up `--max-pods` overhead without actually using private IP addresses pre-allocated by Azure CNI is because the above 5 pods are using the _host network_, which means that they share the host IP address, which is not one of the (in the above example) 12 IP addresses pre-allocated by Azure CNI. The lesson here is that if you are strictly trying to optimize for IP address density in your cluster VNET, you will never be able to achieve 100% density in your Azure CNI cluster due to the fact that `kube-scheduler` and Azure CNI share a common configuration upper bound to restrict the number of pods, and the number of IP addresses in the CNI network CIDR, respectively. In practice, because some pods do nut use the CNI network CIDR, we will always run out of space to schedule pods before we use all of the pre-allocated IP addresses in our cluster VNET, and thus there will always be "unused" IP addresses in our cluster that are, unfortunately, counting against our upper limit of 65,336 IP addresses. So, the guidance is to plan carefully, and be aware that you'll never be able to optimize IP address density 100%.

2. Achieving optimal pod density will require a "best effort" affinity between the resource overhead your pod workloads require, the number of pods that you want to fit on any one particular node, and an appropriate VM SKU that will provide the appropriate resource surface area (available CPU/memory/storage) to accommodate as closely as possible the number of pods that your node can run before the scheduler has to choose another node due to `--max-pods` being reached.

To make this a bit more concrete, using the above example (including CSI drivers), we know that we can schedule 7 workload pods on any given node before our `--max-pods` configuration prevents the scheduler from adding any more pods. What we want to do is to anticipate how much CPU, memory, and storage is required for the "operational average" of 7 of our workload pods on our cluster, and then pick a VM that will provide enough CPU, memory, and storage to accommodate those requirements. For example, if we calculated the the average CPU core count of our production workload pods was 0.8, then we would want to make sure that we use a VM SKU with at least 6 cores (0.8 * 7 equals 5.6). Then we would do the same for memory. At this point, you would choose the VM SKU that matched both values, and that provided any amount of additional overhead that makes sense for your operational environment (perhaps in your tests you want at least a single core for the kubelet daemon, and another core for running ad hoc scripts).

The larger point of all this is that if you want to maximize the size of your cluster, you will have to take into account the Azure VNET limits inherited by the Azure CNI solution, and tune your nodes accordingly.

### configure for more apiserver requests

If you're running a cluster with hundreds or thousands of nodes, you will want to tune your apiserver accordingly to deal with the expected increase in activity. Here are some recommended changes across various components. We recommend you read up on these configurations, and observe your large clusters in real time to determine the best values for your environment!


```
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "kubernetesConfig": {
        "schedulerConfig": {
          "--kube-api-burst": "800",
          "--kube-api-qps": "600"
        },
        "controllerManagerConfig": {
          "--kube-api-burst": "800",
          "--kube-api-qps": "600"
        },
        "apiServerConfig": {
          "--delete-collection-workers": "250"
        }
        ....
      }
      ...
    },
    ...
  }
  ...
}
```

### 1000 nodes per VMSS node pool

There is a hard limit of 1000 VMs in a single VMSS. Reference:

- https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/overview

AKS Engine uses VMSS for node pools, and so if you want to build more than 1000 nodes, you must use more than one node pool. For example, to anticipate as many as 5000 nodes using a common node configuration:

```
{
  "apiVersion": "vlabs",
  "properties": {
    "agentPoolProfiles": [
      {
        "name": "nodepool1",
        "count": 1,
        "vmSize": "Standard_D8s_v3"
      },
      {
        "name": "nodepool2",
        "count": 1,
        "vmSize": "Standard_D8s_v3"
      },
      {
        "name": "nodepool3",
        "count": 1,
        "vmSize": "Standard_D8s_v3"
      },
      {
        "name": "nodepool4",
        "count": 1,
        "vmSize": "Standard_D8s_v3"
      },
      {
        "name": "nodepool5",
        "count": 1,
        "vmSize": "Standard_D8s_v3"
      }
    ],
    ...
  }
  ...
}
```