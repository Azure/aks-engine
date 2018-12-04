# AKS-Engine - Units of Kubernetes on Azure!

[![Coverage Status](https://codecov.io/gh/Azure/aks-engine/branch/master/graph/badge.svg)](https://codecov.io/gh/Azure/aks-engine)
[![CircleCI](https://circleci.com/gh/Azure/aks-engine/tree/master.svg?style=svg)](https://circleci.com/gh/Azure/aks-engine/tree/master)
[![GoDoc](https://godoc.org/github.com/Azure/aks-engine?status.svg)](https://godoc.org/github.com/Azure/aks-engine)

## Overview

AKS-Engine leverages ARM (Azure Resource Manager) to build Kubernetes IaaS in Azure. AKS-Engine provides convenient tooling to quickly bootstrap clusters, and is the vm provisioning implementation for [AKS](https://docs.microsoft.com/en-us/azure/aks), Azure's managed Kubernetes service offering.

More info, including a thorough walkthrough is [here](docs/aksengine.md).

## User guides

[This guide](docs/kubernetes.md) walks you through your first cluster deployment.

These guides cover more advanced features to try out after you have built your first cluster:

* [Cluster Definition](docs/clusterdefinition.md) - describes the components of the cluster definition file
* [Custom VNET](examples/vnet) - shows how to use a custom VNET
* [Attached Disks](examples/disks-storageaccount) - shows how to attach up to 4 disks per node
* [Managed Disks](examples/disks-managed) - shows how to use managed disks
* [Large Clusters](examples/largeclusters) - shows how to create cluster sizes of up to 1200 nodes

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
