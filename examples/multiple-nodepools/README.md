# AKS Engine - Multiple Node Pools

aks-engine supports creating a Kubernetes cluster with more than one node pool. These pools can have different configurations, such as VM size or storage profile.

A cluster with multiple node pools can help you schedule CPU-intensive jobs to VM nodes with high processing power, or I/O intensive jobs to VMs with the fastest storage. Use [nodeSelectors][] or [resource requests][] to ensure that Pods are scheduled to nodes in the appropriate pool.

A complete example is contained in the `multipool.json` API model in this directory. To add a node pool to an existing apimodel, just add another entry to the `agentPoolProfile` section:

```json
    "agentPoolProfiles": [
      {
        "name": "workerpool",
        "count": 2,
        "vmSize": "Standard_D2_v3"
      },
      {
        "name": "gpupool",
        "count": 5,
        "vmSize": "Standard_NC6"
      }
    ],
```

When scaling the size of a node pool up or down with aks-engine, make sure to specify which node pool:

```sh
$ aks-engine scale -g mymultipoolRG --api-model ./multicluster/cluster.json --new-node-count 10 --node-pool gpupool
```

[nodeSelectors]: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
[resource requests]: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
