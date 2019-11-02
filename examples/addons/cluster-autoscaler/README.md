# Cluster Autoscaler addon for VMSS pools

[Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) is a tool that automatically adjusts the size of the Kubernetes cluster when one of the following conditions is true:

- there are pods that failed to run in the cluster due to insufficient resources,
- there are nodes in the cluster that have been underutilized for an extended period of time and their pods can be placed on other existing nodes.

More information on the Azure cloudprovider implementation of cluster-autoscaler can be found [here](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/cloudprovider/azure/README.md).

# Example

Here's a simple example of a cluster configuration (api model) that includes the cluster-autoscaler addon:

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "addons": [
          {
            "name": "cluster-autoscaler",
            "enabled": true,
            "pools": [
              {
                "name": "agentpool1",
                "config": {
                  "min-nodes": "3",
                  "max-nodes": "10"
                }
              }
            ],
            "config": {
              "scan-interval": "1m"
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
        "name": "agentpool1",
        "count": 3,
        "vmSize": "Standard_DS2_v2",
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

The above example enables the cluster-autoscaler addon on your cluster at cluster create time, and engages it for the pool `agentpool1` with the following pool-specific configuration:

- don't scale down below the minimum node count of 3
  - `"min-nodes": "3"`
- don't scale up beyond the maximum node count of 10
  - `"max-nodes": "10"`

And the following cluster-wide configuration:

- check for unscheduled pods across the cluster every 1 minute; check for node removal conditions every 1 minute
  - `"scan-interval": "1m"`

# Configure cluster-autoscaler addon

The AKS Engine cluster-autoscaler addon will implement the set of configurable cluster-autoscaler options appropriate for the version of Kubernetes the cluster is running. Because the Azure cloudprovider cluster-autoscaler implementation has a concept of "node pools", there is pool-specific configuration, as well as configuration that applies to the cluster-autoscaler addon generally (against all node pools in the cluster).

## Pool-specific configuration

Pool-specific configuration is declared inside the `pools` array in your "cluster-autoscaler" `addons` configuration. For example:

```
        "addons": [
          ...,
          {
            "name": "cluster-autoscaler",
            "enabled": true,
            "pools": [
              {
                "name": "agentpool1",
                "config": {
                  "min-nodes": "3",
                  "max-nodes": "10"
                }
              }
            ],
            ...
          },
          ...
        ]
        ...
    "agentPoolProfiles": [
      ...,
      {
        "name": "agentpool1",
        "count": 3,
        "vmSize": "Standard_DS2_v2",
        "storageProfile": "ManagedDisks"
      },
      ...
    ],
```

In the above cluster configuration snippet, we declare a pool-specific configuration for the node pool identified by `"name": "agentpool1"` in the `agentPoolProfiles` array. The name of the pool in the addon `pools` array must match the name of the pool in the `agentPoolProfiles` exactly.

Here is the complete set of pool-specific configuration:

| Name           | Required | Description                                    | Default Value                                              |
| -------------- | -------- | ---------------------------------------------- | ---------------------------------------------------------- |
| min-nodes      | no       | minimum node count                             | equal to the "count" in the equivalent node pool in the `agentPoolProfiles` array                                                         |
| max-nodes      | no       | maximum node count                             | equal to the "count" in the equivalent node pool in the `agentPoolProfiles` array                                                          |

Values in the `pools` config object are always strings.

## Cluster configuration

Cluster configuration that applies generally across all node pools is declared in the `config` object in your cluster-autoscaler `addons` configuration. For example:

```
        "addons": [
          ...,
          {
            "name": "cluster-autoscaler",
            "enabled": true,
            "pools": [
              ...
            ],
            "config": {
              "scan-interval": "1m",
              "scale-down-delay-after-add": "60m0s",
              "skip-nodes-with-local-storage": "true",
              "stderrthreshold": "3"
            }
          },
          ...
        ]
        ...
```

Values in the `config` object are always strings.

Here is the complete set of cluster configuration:

| Name           | Required | Description                                    | Default Value                                              |
| -------------- | -------- | ---------------------------------------------- | ---------------------------------------------------------- |
| namespace      | no       | Which namespace to run the cluster-autoscaler in?                             | "kube-system" |
| scan-interval      | no       | How often cluster is reevaluated for scale up or down                             | "1m" |
| expendable-pods-priority-cutoff      | no       | Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.                             | "-10" |
| ignore-daemonsets-utilization (>= k8s 1.13.0)     | no       | Should CA ignore DaemonSet pods when calculating resource utilization for scaling down                             | "false" |
| ignore-mirror-pods-utilization (>= k8s 1.13.0)      | no       | Should CA ignore Mirror pods when calculating resource utilization for scaling down                             | "false" |
| max-autoprovisioned-node-group-count      | no       | The maximum number of autoprovisioned groups in the cluster                             | "15" |
| max-empty-bulk-delete      | no       | Maximum number of empty nodes that can be deleted at the same time                             | "10" |
| max-failing-time      | no       | Maximum time from last recorded successful autoscaler run before automatic restart                             | "15m0s" |
| max-graceful-termination-sec      | no       | Maximum number of seconds CA waits for pod termination when trying to scale down a node                             | "600" |
| max-inactivity      | no       | Maximum time from last recorded autoscaler activity before automatic restart                             | "10m0s" |
| max-node-provision-time      | no       | Maximum time CA waits for node to be provisioned                             | "15m0s" |
| max-nodes-total      | no       | Maximum number of nodes in all node groups. Cluster autoscaler will not grow the cluster beyond this number                             | "0" (i.e., no maximum) |
| max-total-unready-percentage      | no       | Maximum percentage of unready nodes in the cluster.  After this is exceeded, CA halts operations                             | "45" |
| memory-total      | no       | Minimum and maximum number of gigabytes of memory in cluster, in the format <min>:<max>. Cluster autoscaler will not scale the cluster beyond these numbers                             | "0:6400000" |
| min-replica-count      | no       | Minimum number or replicas that a replica set or replication controller should have to allow their pods deletion in scale down                             | "0" |
| new-pod-scale-up-delay (>= k8s 1.13.0)     | no       | Pods less than this old will not be considered for scale-up                             | "0s" |
| node-autoprovisioning-enabled      | no       | Should CA autoprovision node groups when needed                             | "false" |
| ok-total-unready-count      | no       | Number of allowed unready nodes, irrespective of max-total-unready-percentage                             | "3" |
| scale-down-candidates-pool-ratio      | no       | A ratio of nodes that are considered as additional non empty candidates for scale down when some candidates from previous iteration are no longer valid. Lower value means better CA responsiveness but possible slower scale down latency. Higher value can affect CA performance with big clusters (hundreds of nodes). Set to 1.0 to turn this heuristics off - CA will take all nodes as additional candidates.                             | "0.1" |
| scale-down-candidates-pool-min-count      | no       | Minimum number of nodes that are considered as additional non empty candidates for scale down when some candidates from previous iteration are no longer valid. When calculating the pool size for additional candidates we take max(#nodes * scale-down-candidates-pool-ratio, scale-down-candidates-pool-min-count).                             | "50" |
| scale-down-delay-after-add      | no       | How long after scale up that scale down evaluation resumes                             | "10m0s" |
| scale-down-delay-after-delete      | no       | How long after node deletion that scale down evaluation resumes, defaults to the scan-interval value                             | "1m" |
| scale-down-delay-after-failure      | no       | How long after scale down failure that scale down evaluation resumes                             | "3m0s" |
| scale-down-enabled      | no       | Should CA scale down the cluster                             | "true" |
| scale-down-non-empty-candidates-count      | no       | Maximum number of non empty nodes considered in one iteration as candidates for scale down with drain. Lower value means better CA responsiveness but possible slower scale down latency. Higher value can affect CA performance with big clusters (hundreds of nodes). Set to non positive value to turn this heuristic off - CA will not limit the number of nodes it considers.                             | "30" |
| scale-down-unneeded-time      | no       | How long a node should be unneeded before it is eligible for scale down                             | "10m0s" |
| scale-down-unready-time      | no       | How long an unready node should be unneeded before it is eligible for scale down                             | "20m0s" |
| scale-down-utilization-threshold      | no       | Sum of cpu or memory of all pods running on the node divided by node's corresponding allocatable resource, below which a node can be considered for scale down                             | "0.5" |
| skip-nodes-with-local-storage      | no       | If true cluster autoscaler will never delete nodes with pods with local storage, e.g. EmptyDir or HostPath                             | "false" |
| skip-nodes-with-system-pods      | no       | If true cluster autoscaler will never delete nodes with pods from kube-system (except for DaemonSet or mirror pods)                            | "true" |
| unremovable-node-recheck-timeout      | no       | The timeout before we check again a node that couldn't be removed before                            | "5m0s" |
| v      | no       | log verbosity                            | "3" |
| write-status-configmap      | no       | Should CA write status information to a configmap                            | "true" |
| balance-similar-node-groups      | no       | Detect similar node groups and balance the number of nodes between them                            | "true" |

# Addon mode

You may set the desired `addonmanager.kubernetes.io/mode` value for the cluster-autoscaler addon by passing in a `"mode"` configuration, e.g.:

```
        "addons": [
          ...,
          {
            "name": "cluster-autoscaler",
            "enabled": true,
            "mode": "EnsureExists"
            "pools": [
              ...
            ],
            "config": {
              ...
            }
          },
          ...
        ]
        ...
```
By default we set the mode to `"EnsureExists"` so that you are able to continously manage the cluster-autoscaler configuration (`kubectl edit deployment cluster-autoscaler -n <namespace>`) without the `kube-addon-manager` component overwriting any applied changes after cluster creation time (this is the practical effect of the `"Reconcile"` mode). For more information about how addon-manager reconciles addon configuration, see references to `addonmanager.kubernetes.io/mode` here:

- https://github.com/kubernetes/kubernetes/tree/master/cluster/addons/addon-manager
