# AKS-Engine - Units of Kubernetes on Azure!

[![Coverage Status](https://codecov.io/gh/Azure/aks-engine/branch/master/graph/badge.svg)](https://codecov.io/gh/Azure/aks-engine)
[![CircleCI](https://circleci.com/gh/Azure/aks-engine/tree/master.svg?style=svg)](https://circleci.com/gh/Azure/aks-engine/tree/master)
[![GoDoc](https://godoc.org/github.com/Azure/aks-engine?status.svg)](https://godoc.org/github.com/Azure/aks-engine)

## Overview

AKS-Engine leverages ARM (Azure Resource Manager) to build Kubernetes IaaS in Azure. AKS-Engine provides convenient tooling to quickly bootstrap clusters, and implements cluster provisioning and lifecycle operations for [AKS](https://docs.microsoft.com/en-us/azure/aks), Azure's managed Kubernetes service offering.

## Install AKS-Engine

To install via homebrew, run `brew tap azure/aks-engine && brew install azure/aks-engine/aks-engine` or download the binary via the [github releases page](https://github.com/Azure/aks-engine/releases)

## Getting Started

Building a Kubernetes cluster on Azure using AKS-Engine is as simple as:

1. `aks-engine init` to set up aks-engine (after the binary has been installed in your `$PATH`)
  - This will walk you through setting default configurations
    - Azure subscription context, including the service principal configuration that aks-engine will use to scaffold new IaaS on behalf of your Azure subscription
    - Preferred region(s)
    - Preferred default cluster configuration
2. `aks-engine create` to create a new cluster using the default Kubernetes configuration.

Cluster creation time will take between 3-5 minutes, after which your console session's `KUBECONFIG` context will be automatically set to your new cluster.

```
$ kubectl get nodes
NAME                                           STATUS    ROLES     AGE       VERSION
k8s-Standard_D2_v2-10678025-vmss000000         Ready     agent     18s       v1.10.12
k8s-Standard_D2_v2-10678025-vmss000001         Ready     agent     17s       v1.10.12
k8s-Standard_D2_v2-10678025-vmss000002         Ready     agent     17s       v1.10.12
k8s-master-10678025-0                          Ready     master    11s       v1.10.12
```

Scale out your cluster:

```
$ aks-engine scale -n 5
```

Scale out will similarly take between 3-5 minutes.

```
$ kubectl get nodes
NAME                                           STATUS    ROLES     AGE       VERSION
k8s-Standard_D2_v2-10678025-vmss000000         Ready     agent     18s       v1.10.12
k8s-Standard_D2_v2-10678025-vmss000001         Ready     agent     17s       v1.10.12
k8s-Standard_D2_v2-10678025-vmss000002         Ready     agent     17s       v1.10.12
k8s-Standard_D2_v2-10678025-vmss000003         Ready     agent     17s       v1.10.12
k8s-Standard_D2_v2-10678025-vmss000004         Ready     agent     17s       v1.10.12
k8s-master-10678025-0                          Ready     master    11s       v1.10.12
```

Scale in:

```
$ aks-engine scale -n 1
...
$ kubectl get nodes
NAME                                           STATUS    ROLES     AGE       VERSION
k8s-Standard_D2_v2-10678025-vmss000000         Ready     agent     18s       v1.10.12
k8s-master-10678025-0                          Ready     master    11s       v1.10.12
```

Upgrade your cluster:

```
$ aks-engine update -v 1.11
```

Cluster upgrade time will take between 5-20 minutes *per node* (including master node(s) running the control plane), due to cordon/drain, the addition of new vms with the desired changes, and the deletion of vms with the previous configuration. And it does this according to a rolling, one-at-a-time strategy to minimize operational side-effects. Depending on the size of your cluster, brew some coffee in the French press style, walk your dog(s), or go do that open source contribution you've been putting off for a few months.

...
$ kubectl get nodes
NAME                                           STATUS    ROLES     AGE       VERSION
k8s-Standard_D2_v2-10678025-vmss000005         Ready     agent     18s       v1.11.6
k8s-master-10678025-0                          Ready     master    11s       v1.11.6
```

More info, including a thorough tour through the CLI is [here](docs/aksengine.md).

Please see the [FAQ](/docs/faq.md) for answers about AKS-Engine and its progenitor ACS-Engine.

## User guides

[These guides](docs/kubernetes.md) will walk you through some of the common cluster configurations supported by AKS-Engine.

## Contributing

Follow the [developers guide](docs/developers.md) to set up your environment.

To build aks-engine, run `make build`. If you are developing with a working [Docker environment](https://docs.docker.com/engine), you can also run `make dev` (or `makedev.ps1` on Windows) first to start a Docker container and run `make build` inside the container.

Please follow these instructions before submitting a PR:

1. Execute `make test` to run unit tests.
2. Manually test deployments if you are making modifications to the templates.
* For example, if you have to change the expected resulting templates then you should deploy the relevant example cluster definitions to ensure that you are not introducing any regressions.
3. Make sure that your changes are properly documented and include relevant unit tests.

## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
