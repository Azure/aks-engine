# Upgrading Kubernetes Clusters

## Prerequisites

All the commands in this guide require both the Azure CLI and `aks-engine`. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing.

This guide assumes you already have deployed a cluster using `aks-engine`. For more details on how to do that see [deploy](../tutorials/deploy.md).

## Upgrade

This document provides guidance on how to upgrade the Kubernetes version for an existing `aks-engine` cluster and recommendations for adopting `aks-engine upgrade` as a production quality tool.

### Know before you go

In order to ensure that your `aks-engine upgrade` operation runs smoothly, there are a few things you should be aware of before getting started.

1) You will need access to the original `_output/` directory files that were generated during cluster deployment in order to run `upgrade`.  `aks-engine` will use the `--deployment-dir` argument provided in the upgrade command to introspect the Azure Resource Manager (ARM) template and parameters to determine the cluster's current Kubernetes version and its current state.

2) `aks-engine upgrade` expects a cluster configuration that conforms to the current state of the cluster. In other words, the Azure resources deployed by `aks-engine` should be in the same state that when they were originally created by `aks-engine`. If you perform manual operations on your Azure IaaS resources (other than `aks-engine scale` and `aks-engine upgrade`) you risk not being able to upgrade your cluster successfully as `aks-engine` won't be able to reconcile the cluster's current state. This includes naming of resources; `aks-engine` relies on some resources (such as VMs) to be named in accordance with the original `aks-engine` deployment.

3) `aks-engine upgrade` allows one subsequent minor version upgrade at a time, as long as the next minor version is supported. If the next minor version is deprecated, `aks-engine` allows upgrade to the next supported minor version.

To get the list of all available Kubernetes versions and upgrades, run the `orchestrators` command and specify Kubernetes orchestrator type:

```bash
./bin/aks-engine orchestrators --orchestrator Kubernetes
```

To get the information specific to the cluster, provide its current Kubernetes version in the `version` arg:
```bash
./bin/aks-engine orchestrators --orchestrator Kubernetes --version 1.11.5
```

4) If using `aks-engine upgrade` in production, it is recommended to run a dry run test on an identical cluster before performing the upgrade, especially if the cluster configuration is "unusual", or in other words differs significantly from defaults. The reason for this is that `aks-engine` supports many different cluster configurations and the extent of E2E testing that the AKS-Engine team runs cannot possibly cover every single configuration out there. Therefore, it is recommended that you make sure that your specific cluster configuration works with the existing upgrade implementation before starting this long-running operation.

5) `aks-engine upgrade` is backwards compatible. If you deployed with `aks-engine` version `0.27.x`, you can run upgrade with version `0.29.y`. In fact, it is recommended that you use the latest available `aks-engine` version when running an upgrade operation. This will ensure that you get the latest available software and bug fixes in your upgraded cluster.

### Under the hood

During the upgrade, *aks-engine* successively visits virtual machines that constitute the cluster (first the master nodes, then the agent nodes) and performs the following operations:
 - cordon the node and drain existing workload
 - delete the VM
 - create new VM and install desired orchestrator version
 - add the new VM to the cluster


### Simple steps to run upgrade

Once the desired Kubernetes version is finalized, call the *upgrade* command:
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



### What can go wrong

Definitive guidance to aid customers who want to use the above tooling to their advantage
what can/can't they do to their Azure IaaS resources
what can go wrong


By its nature, the upgrade operation is long running and potentially could fail for various reasons, such as temporary lack of resources, etc. In this case, rerun the command. The *upgrade* command is idempotent, and will pick up execution from the point it failed on.