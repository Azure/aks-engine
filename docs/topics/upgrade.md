# Upgrading Kubernetes Clusters

## Prerequisites

All the commands in this guide require both the Azure CLI and `aks-engine`. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing.

This guide assumes you already have deployed a cluster using `aks-engine`. For more details on how to do that see [deploy](../tutorials/deploy.md).

## Upgrade

This document provides guidance on how to upgrade the Kubernetes version for an existing `aks-engine` cluster and recommendations for adopting `aks-engine upgrade` as a production quality tool.

<a name="pre-requirements"></a>

### Know before you go

In order to ensure that your `aks-engine upgrade` operation runs smoothly, there are a few things you should be aware of before getting started.

1) You will need access to the original `_output/` directory files that were generated during cluster deployment in order to run `upgrade`.  `aks-engine` will use the `--deployment-dir` argument provided in the upgrade command to introspect the `apimodel.json` file and Azure Resource Manager (ARM) parameters in order to determine the cluster's current Kubernetes version, as well as all other cluster configuration data as defined by `aks-engine` during the last time that `aks-engine` was used to deploy, scale, or upgrade the cluster.

2) `aks-engine upgrade` expects a cluster configuration that conforms to the current state of the cluster. In other words, the Azure resources deployed by `aks-engine` should be in the same state that when they were originally created by `aks-engine`. If you perform manual operations on your Azure IaaS resources (other than `aks-engine scale` and `aks-engine upgrade`) you risk not being able to upgrade your cluster successfully, as the aks-engine-generated ARM template won't be reconcilable against the state of the Azure resources that reside in the resource group. This includes naming of resources; `aks-engine` relies on some resources (such as VMs) to be named in accordance with the original `aks-engine` deployment. In summary, the set of Azure resources in the resource group are mutually reconcilable only if they have been exclusively created and managed as the result of a series of successive ARM template deployments originating from `aks-engine`.

3) `aks-engine upgrade` allows upgrading the Kubernetes version to any aks-engine-supported patch release in the current minor release channel that is greater than the current version on the cluster (e.g., from `1.11.4` to `1.11.5`), or to the next aks-engine-supported minor version (e.g., from `1.11.4` to `1.12.5`). In practice, the next aks-engine-supported minor version will commonly be a single minor version ahead of the current cluster version. However, if the cluster has not been upgraded in a significant amount of time, the "next" minor version may have actually been deprecated by aks-engine. In such a case, your long-lived cluster will be upgradable to the nearest, supported minor version that `aks-engine` supports at the time of upgrade (e.g., from `1.7.16` to `1.9.11`).

    To get the list of all available Kubernetes versions and upgrades, run the `orchestrators` command and specify Kubernetes orchestrator type:

    ```bash
    ./bin/aks-engine orchestrators --orchestrator Kubernetes
    ```

    To get the versions of Kubernetes that your particular cluster version is upgradable to, provide its current Kubernetes version in the `version` arg:

    ```bash
    ./bin/aks-engine orchestrators --orchestrator Kubernetes --version 1.11.5
    ```

4) If using `aks-engine upgrade` in production, it is recommended to stage an upgrade test on an cluster that was built to the same specifications (built with the same cluster configuration + `aks-engine` version) as your production cluster before performing the upgrade, especially if the cluster configuration is "interesting", or in other words differs significantly from defaults. The reason for this is that `aks-engine` supports many different cluster configurations and the extent of E2E testing that the AKS-Engine team runs cannot practically cover every single configuration out there. Therefore, it is recommended that you make sure that your specific cluster configuration works with the existing upgrade implementation before starting this long-running operation.

5) `aks-engine upgrade` is backwards compatible. If you deployed with `aks-engine` version `0.27.x`, you can run upgrade with version `0.29.y`. In fact, it is recommended that you use the latest available `aks-engine` version when running an upgrade operation. This will ensure that you get the latest available software and bug fixes in your upgraded cluster.

### Under the hood

During the upgrade, *aks-engine* successively visits virtual machines that constitute the cluster (first the master nodes, then the agent nodes) and performs the following operations:

- cordon the node and drain existing workload
- delete the VM
- create new VM and install desired orchestrator version
- add the new VM to the cluster

### Simple steps to run upgrade

Once you have read all the [prerequirements](#pre-requirements), run `aks-engine upgrade` with the appropriate arguments:

```bash
./bin/aks-engine upgrade \
  --subscription-id <subscription id> \
  --deployment-dir <aks-engine output directory > \
  --location <resource group location> \
  --resource-group <resource group name> \
  --upgrade-version <desired Kubernetes version> \
  --auth-method client_secret \
  --client-id <service principal id> \
  --client-secret <service principal secret>
```

For example,

```bash
./bin/aks-engine upgrade \
  --subscription-id xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
  --deployment-dir ./_output/test \
  --location westus \
  --resource-group test-upgrade \
  --upgrade-version 1.8.7 \
  --client-id xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
  --client-secret xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

## Known Limitations

### Manual reconciliation

The upgrade operation is long running, and for large clusters, more susceptible to single operational failures. This is based on the design principle of upgrade enumerating, one-at-a-time, through each node in the cluster. A transient Azure resource allocation error could thus interrupt the successful progression of the overall transaction. At present, the upgrade operation is implemented to "fail fast"; and so, if a well formed upgrade operation fails before completing, it can be manually retried by invoking the exact same command line arguments as were sent originally. The upgrade operation will enumerate through the cluster nodes, skipping any nodes that have already been upgraded to the desired Kubernetes version. Those nodes that match the *original* Kubernetes version will then, one-at-a-time, be cordon and drained, and upgraded to the desired version. Put another way, an upgrade command is designed to be idempotent across retry scenarios.

### Cluster-autoscaler + VMSS

At present, the Azure cloudprovider cluster-autoscaler implementation for VMSS relies upon the original ARM template deployment specification to inform the Azure IaaS (VM, NIC, CustomScriptExtension, etc) configuration to scale out new nodes.

Because `aks-engine upgrade` also employs ARM template deployments to evolve the cluster state going forward, as soon as you upgrade your cluster, the cluster-autoscaler will no longer scale out your cluster according to the latest version. You will see cluster-autoscaler scale out scenarios that look like this one:

```
$ kubectl get nodes
NAME                                STATUS    ROLES     AGE       VERSION
k8s-agentpool-27988949-vmss00000a   Ready     agent     1h        v1.11.5
k8s-agentpool-27988949-vmss00000m   Ready     agent     1m        v1.10.12
k8s-agentpool-27988949-vmss00000n   Ready     agent     1m        v1.10.12
k8s-agentpool-27988949-vmss00000o   Ready     agent     1m        v1.10.12
k8s-agentpool-27988949-vmss00000p   Ready     agent     1m        v1.10.12
k8s-agentpool-27988949-vmss00000q   Ready     agent     1m        v1.10.12
k8s-agentpool-27988949-vmss00000r   Ready     agent     1m        v1.10.12
k8s-master-27988949-0               Ready     master    2h        v1.11.5
```

For this reason, we do not recommend incorporating `aks-engine upgrade` into your operational workflow if you also rely upon cluster-autoscaler + VMSS

### Cluster-autoscaler + VMAS

A similar scenario exists for VMAS as well, but because the cluster-autoscaler spec includes a configurable ARM template deployment reference, you may manually maintain that reference over time to be current with the ARM template deployment that `aks-engine upgrade` creates during an upgrade operation.

// TODO describe the above


