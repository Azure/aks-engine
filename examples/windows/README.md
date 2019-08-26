# AKS Engine

## Overview

These cluster definition examples demonstrate how to create customized Docker Enabled Cluster with Windows on Microsoft Azure.

## User Guides

* [Kubernetes Windows Walkthrough](../../docs/topics/windows.md) - shows how to create a hybrid Kubernetes Windows enabled Docker cluster on Azure.
* [Building Kubernetes Windows binaries](../../docs/howto/building-windows-kubernetes-binaries.md) - shows how to build kubernetes windows binaries for use in a Windows Kubernetes cluster.

## Sample Deployments

### Kubernetes

- kubernetes.json - this is the simplest case for a 2-node Windows Kubernetes cluster
- kubernetes-custom-image.json - example using an existing Azure image for Windows nodes.
- kubernetes-shared-image.json - exmple using an Azure image from a shared image gallery for Windows nodes.
- kubernetes-custom-vhd.json - exmaple using a custom VHD (uploaded to an Azure storage account or other accessible location) for Windows nodes.
- kubernetes-hybrid.json - example with both Windows & Linux nodes in the same cluster
- kubernetes-hyperv.json - example with 2 Windows nodes with the [alpha Hyper-V isolation support](https://kubernetes.io/docs/getting-started-guides/windows/#hyper-v-containers) enabled
- kubernetes-wincni.json - example using kubenet plugin on Linux nodes and WinCNI on Windows
- kubernetes-windows-version.json - example of how to build a cluster with a specific Windows patch version
