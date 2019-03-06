# AKS Engine - Units of Kubernetes on Azure!

[![Build Status](https://msazure.visualstudio.com/One/_apis/build/status/Custom/Compute/ContainerService/AKS%20Engine%20CI%20E2E?branchName=master)](https://msazure.visualstudio.com/One/_build/latest?definitionId=50661&branchName=master)
[![Coverage Status](https://codecov.io/gh/Azure/aks-engine/branch/master/graph/badge.svg)](https://codecov.io/gh/Azure/aks-engine)
[![GoDoc](https://godoc.org/github.com/Azure/aks-engine?status.svg)](https://godoc.org/github.com/Azure/aks-engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/Azure/aks-engine)](https://goreportcard.com/report/github.com/Azure/aks-engine)

AKS Engine is the easiest way to provision a self-managed Kubernetes cluster on Azure.

## Overview

AKS Engine provides convenient tooling to quickly bootstrap Kubernetes clusters on Azure. By leveraging [ARM (Azure Resource Manager)][ARM], AKS Engine helps you create, destroy and maintain clusters provisioned with basic IaaS resources in Azure. AKS Engine is also the library used by AKS for performing these operations to provide managed service implementations.

## Getting started

Depending on how new you are to AKS Engine, you can try [a tutorial][tutorials], or just dive straight into the [documentation][docs].

Please see the [FAQ][] for answers about AKS Engine and its progenitor ACS-Engine.

## Sharpen your skills

The official [AKS Engine documentation][docs] covers everything you need to know about AKS Engine (and then some).

## Join the community

Want to get involved? The [community guide][community] covers everything you need to know about the AKS Engine community and how you can contribute.

## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

[ARM]: https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-overview
[community]: docs/community/README.md
[docs]: docs/README.md
[FAQ]: docs/faq.md
[tutorials]: docs/tutorials/README.md
