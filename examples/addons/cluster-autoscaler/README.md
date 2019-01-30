# Cluster Autoscaler (VMSS) Add-on

[Cluster Autoscaler](https://github.com/kubernetes/autoscaler) is a tool that automatically adjusts the size of the Kubernetes cluster when:

- there are pods that failed to run in the cluster due to insufficient resources.
- some nodes in the cluster are so underutilized, for an extended period of time, that they can be deleted and their pods will be easily placed on some other, existing nodes.

This is the Kubernetes Cluster Autoscaler add-on for Virtual Machine Scale Sets. Add this add-on to your json file as shown below to automatically enable cluster autoscaler in your new Kubernetes cluster.

See [this document](../../../docs/topic/upgrade.md) for details on known limitiations with `aks-engine upgrade` and VMSS.

To use this add-on, make sure your cluster's Kubernetes version is 1.10 or above and your agent pool `availabilityProfile` is set to `VirtualMachineScaleSets`. By default, the first agent pool will autoscale the node count between 1 and 5. You can override these settings in `config` section of the `cluster-autoscaler` add-on.

> At this time, only the primaryScaleSet (the first agent pool) is monitored by the autoscaler. If you use this addon on your single agent pool cluster and later manually add more VMSS agent pools and want to use this cluster-autoscaler addon to monitor (and scale) other node pools, you must manually edit the autoscaler YAML at `/etc/kubernetes/addons` on each master node.

> If you are creating a cluster with multiple VMSS agent pools, or you plan to manually evolve your agent pool count over time, we recommend you don't use this addon. See the [cluster-autoscaler](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/cloudprovider/azure/README.md) docs for guidance on how to craft a cluster-autoscaler specification that can accommodate a multiple VMSS agent pool cluster.

The following is an example:

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "orchestratorRelease": "1.10",
      "kubernetesConfig": {
        "useManagedIdentity": true,
        "addons": [
          {
            "name": "cluster-autoscaler",
            "enabled": true,
            "config": {
              "min-nodes": "1",
              "max-nodes": "5",
              "scan-interval": "10s",
            }
          }
        ]
      }
    },
    "masterProfile": {
      "count": 1,
      "dnsPrefix": "",
      "vmSize": "Standard_DS2_v2"
    },
    "agentPoolProfiles": [
      {
        "name": "agentpool",
        "count": 1,
        "vmSize": "Standard_DS2_v2",
        "availabilityProfile": "VirtualMachineScaleSets",
        "storageProfile": "ManagedDisks"
      }
    ],
    "linuxProfile": {
      "adminUsername": "azureuser",
      "ssh": {
        "publicKeys": [
          {
            "keyData": ""
          }
        ]
      }
    }
  }
}
```

You should see cluster autoscaler as running after running:

```
$ kubectl get pods -n kube-system
```

Follow the README at https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler for more information.

## Configuration

| Name           | Required | Description                                    | Default Value                                              |
| -------------- | -------- | ---------------------------------------------- | ---------------------------------------------------------- |
| min-nodes      | no       | minimum node count                             | 1                                                          |
| max-nodes      | no       | maximum node count                             | 5                                                          |
| scan-interval  | no       | interval to evaluate scale up/down decision    | "10s"                                                      |
| name           | no       | container name                                 | "cluster-autoscaler"                                       |
| image          | no       | image                                          | "gcr.io/google-containers/cluster-autoscaler"              |
| cpuRequests    | no       | cpu requests for the container                 | "100m"                                                     |
| memoryRequests | no       | memory requests for the container              | "300Mi"                                                    |
| cpuLimits      | no       | cpu limits for the container                   | "100m"                                                     |
| memoryLimits   | no       | memory limits for the container                | "300Mi"                                                    |

## Supported Orchestrators

- Kubernetes
