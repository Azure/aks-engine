# Microsoft Azure Kubernetes Service Engine - Kubernetes Upgrade

## Overview

This document describes how to upgrade kubernetes version for an existing cluster.

*aks-engine* supports Kubernetes version upgrades starting from ``1.5`` release.
During the upgrade, *aks-engine* successively visits virtual machines that constitute the cluster (first the master nodes, then the agent nodes) and performs the following operations:
 - cordon the node and drain existing workload
 - delete the VM
 - create new VM and install desired orchestrator version
 - add the new VM to the cluster

*aks-engine* allows one subsequent minor version upgrade at a time, for example, from ``1.6.x`` to ``1.7.y``.

For upgrade that spans over more than a single minor version, this operation should be called several times, each time advancing the minor version by one. For example, to upgrade from ``1.6.x`` to ``1.8.z`` one should first upgrade the cluster to ``1.7.y``, followed by upgrading it to ``1.8.z``

To get the list of all available Kubernetes versions and upgrades, run the *orchestrators* command and specify Kubernetes orchestrator type. The output is a JSON object:
```bash
./bin/aks-engine orchestrators --orchestrator Kubernetes
```

```json
{
  "orchestrators": [
    {
      "orchestratorType": "Kubernetes",
      "orchestratorVersion": "1.7.9",
      "default": true,
      "upgrades": [
        {
          "orchestratorVersion": "1.7.10"
        },
        {
          "orchestratorVersion": "1.7.12"
        },
        {
          "orchestratorVersion": "1.7.13"
        },
        {
          "orchestratorVersion": "1.7.14"
        },
        {
          "orchestratorVersion": "1.8.1"
        },
        {
          "orchestratorVersion": "1.8.0"
        },
        {
          "orchestratorVersion": "1.8.2"
        },
        {
          "orchestratorVersion": "1.8.4"
        },
        {
          "orchestratorVersion": "1.8.6"
        },
        {
          "orchestratorVersion": "1.8.7"
        },
        {
          "orchestratorVersion": "1.8.8"
        },
        {
          "orchestratorVersion": "1.8.9"
        }
      ]
    },
    ...
    ...
    ...
  ]
}
```

To get the information specific to the cluster, provide its current orchestrator version:
```bash
./bin/aks-engine orchestrators --orchestrator Kubernetes --version 1.7.8
```

```json
{
  "orchestrators": [
    {
      "orchestratorType": "Kubernetes",
      "orchestratorVersion": "1.7.8",
      "upgrades": [
        {
          "orchestratorVersion": "1.7.9"
        },
        {
          "orchestratorVersion": "1.7.10"
        },
        {
          "orchestratorVersion": "1.7.12"
        },
        {
          "orchestratorVersion": "1.7.13"
        },
        {
          "orchestratorVersion": "1.7.14"
        },
        {
          "orchestratorVersion": "1.8.0"
        },
        {
          "orchestratorVersion": "1.8.1"
        },
        {
          "orchestratorVersion": "1.8.2"
        },
        {
          "orchestratorVersion": "1.8.4"
        },
        {
          "orchestratorVersion": "1.8.6"
        },
        {
          "orchestratorVersion": "1.8.7"
        },
        {
          "orchestratorVersion": "1.8.8"
        },
        {
          "orchestratorVersion": "1.8.9"
        }
      ]
    }
  ]
}
```

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
  --auth-method client_secret \
  --client-id xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
  --client-secret xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

By its nature, the upgrade operation is long running and potentially could fail for various reasons, such as temporary lack of resources, etc. In this case, rerun the command. The *upgrade* command is idempotent, and will pick up execution from the point it failed on. 

[This directory](https://github.com/Azure/aks-engine/tree/master/examples/k8s-upgrade) contains the following files:
- **README.md** - this file
- **k8s-upgrade.sh** - script invoking upgrade operation
- **\*.json** - cluster definition examples for various orchestrator versions and configurations: Linux clusters, Windows clusters, hybrid clusters.
- **\*.env** - files with environment variables per corresponding cluster definition **.json** file, to set desired kubernetes version passed over to **k8s-upgrade.sh** by the test framework.
