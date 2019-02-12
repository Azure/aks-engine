# Building Windows Kubernetes Binaries

## Background

Microsoft maintains a fork of the Kubernetes project at https://github.com/Azure/kubernetes which includes patches not yet included in upstream Kubernetes for release 1.7 and 1.8; these are needed for Windows containers to function. *From release 1.9, all Windows features are in upstream and Windows binaries no longer needs to be built from Azure fork.*

## Instructions

The following instructions show how to deploy the Windows Kubernetes Binaries and deploy them to an Azure Storage Account.

### Prerequisites

- Azure Storage Account and Azure Storage Container to store Windows binaries
- Access to [wincni.exe] and [hns.psm1](https://github.com/Microsoft/SDN/tree/master/Kubernetes/windows/). Windows CNI is a plugin that supports the Container Network Interface (CNI) network model and interfaces with the Windows Host Networking Service (HNS) to configure host networking and policy.
- Docker installed and running. MacOS users using Docker for Mac must have at [least 3GB of memory allocated to Docker](https://github.com/kubernetes/kubernetes/tree/master/build/#requirements) or building will likely fail.

[build-windows-k8s.sh](../../scripts/build-windows-k8s.sh) does the following:

- Checks out the fork of Azure/kubernetes (includes Windows fixes not yet in upstream Kubernetes, needed for Windows containers to function)
- Builds kubelet.exe and kube-proxy.exe from source in a Docker container
- Downloads kubectl.exe for desired release
- Downloads [NSSM](https://nssm.cc) which is used to start kubelet and kube-proxy on Windows
- Downloads [Windows CNI exe and script] (https://github.com/Microsoft/SDN/tree/master/Kubernetes/windows/)
- Creates an .zip archive of these Windows components
- Uploads archive to Azure Blob Storage

More information about building Kubernetes binaries from source here: https://github.com/kubernetes/kubernetes/tree/master/build/

### Set Azure Storage credentials and Container name

A storage container is used to upload the resulting archive artifact.

```console
$ export AZURE_STORAGE_CONNECTION_STRING="DefaultEndpointsProtocol=https;AccountName=MyStorageAccountName;AccountKey=..." \
AZURE_STORAGE_CONTAINER_NAME=MyStorageContainerName
```

### Run `make build-windows-k8s`

Usage: `make build-windows-k8s K8S_VERSION=1.7.4 PATCH_VERSION=1`

This will build Kubernetes binaries for Windows agents based on Kubernetes release version 1.7.4 and place the zip artifact in Azure Blob Storage.
