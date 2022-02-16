# AKS Engine on Azure Stack Hub

* [Introduction](#introduction)
* [Marketplace Prerequisites](#marketplace-prerequisites)
* [Service Principals and Identity Providers](#service-principals-and-identity-providers)
* [CLI flags](#cli-flags)
* [Cluster Definition (aka API Model)](#cluster-definition-aka-api-model)
  * [location](#location)
  * [kubernetesConfig](#kubernetesConfig)
  * [customCloudProfile](#customCloudProfile)
  * [masterProfile](#masterProfile)
  * [agentPoolProfiles](#agentPoolProfiles)
* [Azure Stack Hub Instances Registered with Azure's China cloud](#azure-stack-hub-instances-registered-with-azures-china-cloud)
* [Disconnected Azure Stack Hub Instances](#disconnected-azure-stack-hub-instances)
* [Volume Provisioner: Container Storage Interface Drivers (preview)](#volume-provisioner-container-storage-interface-drivers-preview)
* [AKS Engine Versions](#aks-engine-versions)
* [Azure Monitor for containers](#azure-Monitor-for-containers)
* [Known Issues and Limitations](#known-issues-and-limitations)
* [Frequently Asked Questions](#frequently-asked-questions)

## Introduction

Specific AKS Engine [versions](#aks-engine-versions) can be used to provision self-managed Kubernetes clusters on [Azure Stack Hub](https://azure.microsoft.com/overview/azure-stack/). AKS Engine's `generate`, [deploy](../tutorials/quickstart.md#deploy), [upgrade](upgrade.md), and [scale](scale.md) commands can be executed as if you were targeting Azure's public cloud. You are only required to slightly update your cluster definition to provide some extra information about your Azure Stack Hub instance.

The goal of this guide is to explain how to provision Kubernetes clusters to Azure Stack Hub using AKS Engine and to capture the differences between Azure and Azure Stack Hub. Bear in mind as well that not every AKS Engine feature or configuration option is currently supported on Azure Stack Hub. In most cases, these are not available because dependent Azure components are not part of Azure Stack Hub.

## Marketplace prerequisites

Because Azure Stack Hub instances do not have infinite storage available, Azure Stack Hub administrators are in charge of managing it by selecting which marketplace items are downloaded from Azure's marketplace. The Azure Stack Hub administrator can follow this [guide](https://docs.microsoft.com/azure-stack/operator/azure-stack-download-azure-marketplace-item) for a general explanation about how to download marketplace items from Azure.

Before you try to deploy the first Kubernetes cluster, make sure these marketplace items were made available to the target subscription by the Azure Stack Hub administrator.

* `Custom Script for Linux 2.0` virtual machine extension
* [Required](#aks-engine-versions) `AKS Base Image` virtual machine

## Service Principals and Identity Providers

Kubernetes uses a `service principal` identity to talk to Azure Stack Hub APIs to dynamically manage resources such as storage or load balancers. Therefore, you will need to create a service principal before you can provision a Kubernetes cluster using AKS Engine.

This [guide](https://docs.microsoft.com/azure-stack/operator/azure-stack-create-service-principals) explains how to create and manage service principals on Azure Stack Hub for both Azure Active Directory (AAD) and Active Directory Federation Services (ADFS) identity providers. This other [guide](../../docs/topics/service-principals.md) is a good resource to understand the permissions that the service principal requires to deploy under your subscription.

Once you have created the required service principal, make sure to assign it the `contributor` role at the target subscription scope.

## CLI flags

To indicate to AKS Engine that your target platform is Azure Stack Hub, all commands require CLI flag `azure-env` to be set to `"AzureStackCloud"`.

If your Azure Stack Hub instance uses ADFS to authenticate identities, then flag `identity-system` is also required.

``` bash
aks-engine deploy \
  --azure-env AzureStackCloud \
  --api-model kubernetes.json \
  --location local \
  --resource-group kube-rg \
  --identity-system adfs \
  --client-id $SPN_CLIENT_ID \
  --client-secret $SPN_CLIENT_SECRET \
  --subscription-id $TENANT_SUBSCRIPTION_ID \
  --output-directory kube-rg
```

## Cluster Definition (aka API Model)

This section details how to tailor your cluster definitions in order to make them compatible with Azure Stack Hub. You can start off from this [template](../../examples/azure-stack/kubernetes-azurestack.json).

Unless otherwise specified down below, standard [cluster definition](../../docs/topics/clusterdefinitions.md) properties should also work with Azure Stack Hub. Please create an [issue](https://github.com/Azure/aks-engine/issues/new) if you find that we missed a property that should be called out.

### location

| Name       | Required | Description                                                   |
| ---------- | -------- | ------------------------------------------------------------- |
| location   | yes      | The region name of the target Azure Stack Hub. |

### kubernetesConfig

`kubernetesConfig` describes Kubernetes specific configuration.

| Name                            | Required | Description                          |
| ------------------------------- | -------- | ------------------------------------ |
| addons                          | no       | A few addons are not supported on Azure Stack Hub. See the [complete list](#unsupported-addons) down below.|
| kubernetesImageBase             | no       | For AKS Engine versions lower than v0.48.0, this is a required field. It specifies the default image base URL to be used for all Kubernetes-related containers such as hyperkube, cloud-controller-manager, pause, addon-manager, etc. This property should be set to `"mcr.microsoft.com/k8s/azurestack/core/"`. |
| networkPlugin                   | yes      | Specifies the network plugin implementation for the cluster. Valid values are `"kubenet"` (default) for k8s software networking implementation and `"azure"`, which provides an Azure native networking experience. |
| networkPolicy                   | no      | Specifies the network policy enforcement tool for the cluster (currently Linux-only). Valid values are: `"azure"` (experimental) for Azure CNI-compliant network policy (note: Azure CNI-compliant network policy requires explicit `"networkPlugin": "azure"` configuration as well). |
| useInstanceMetadata             | no      | Use the Azure cloud provider instance metadata service for appropriate resource discovery operations. This property should be always set to `"false"`. |

### customCloudProfile

`customCloudProfile` contains information specific to the target Azure Stack Hub instance.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| environment                     | no       | The custom cloud type. This property should be always set to `"AzureStackCloud"`. |
| identitySystem                  | yes      | Specifies the identity provider used by the Azure Stack Hub instance. Valid values are `"azure_ad"` (default) and `"adfs"`. |
| portalUrl                       | yes      | The tenant portal URL. |
| dependenciesLocation            | no       | Specifies where to locate the dependencies required to during the provision/upgrade process. Valid values are `"public"` (default), `"china"`, `"german"` and `"usgovernment".` |

### masterProfile

`masterProfile` describes the settings for control plane configuration.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Specifies a valid [Azure Stack Hub VM size](https://docs.microsoft.com/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the control plane's Linux distribution. `"aks-ubuntu-18.04"` is supported. This is a custom image based on UbuntuServer that come with pre-installed software necessary for Kubernetes deployments. |

### agentPoolProfiles

`agentPoolProfiles` are used to create agents with different capabilities.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Describes a valid [Azure Stack Hub VM size](https://docs.microsoft.com/azure-stack/user/azure-stack-vm-sizes). |
| osType                          | no       | Specifies the agent pool's Operating System. Supported values are `"Windows"` and `"Linux"`. Defaults to `"Linux"`. |
| distro                          | yes      | Specifies the control plane's Linux distribution. `"aks-ubuntu-18.04"` is supported. This is a custom image based on UbuntuServer that come with pre-installed software necessary for Kubernetes deployments. |
| availabilityProfile             | yes      | Only `"AvailabilitySet"` is currently supported. |
| acceleratedNetworkingEnabled    | yes      | Use `Azure Accelerated Networking` feature for Linux agents. This property should be always set to `"false"`. |

`linuxProfile` provides the linux configuration for each linux node in the cluster
| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| runUnattendedUpgradesOnBootstrap| no       | Invoke an unattended-upgrade when each Linux node VM comes online for the first time. In practice this is accomplished by performing an `apt-get update`, followed by a manual invocation of `/usr/bin/unattended-upgrade`, to fetch updated apt configuration, and install all package updates provided by the unattended-upgrade facility, respectively. Defaults to `"false"`. |

## Azure Stack Hub Instances Registered with Azure's China cloud

If your Azure Stack Hub instance is located in China, then the `dependenciesLocation` property of your cluster definition should be set to `"china"`. This switch ensures that the provisioning process fetches software dependencies from reachable hosts within China's mainland.

## Disconnected Azure Stack Hub Instances

By default, the AKS Engine provisioning process relies on an internet connection to download the software dependencies required to create or upgrade a cluster (Kubernetes images, etcd binaries, network plugins and so on).

If your Azure Stack Hub instance is air-gapped or if network connectivity in your geographical location is not reliable, then the default approach will not work, take a long time or timeout due to transient networking issues.

To overcome these issues, you should set the `distro` property of your cluster definition to `"aks-ubuntu-18.04"`. This will instruct AKS Engine to deploy VM nodes using a base OS image called `AKS Base Image`. This custom image, generally based on Ubuntu Server, already contains the required software dependencies in its file system. Hence, internet connectivity won't be required during the provisioning process.

The `AKS Base Image` marketplace item has to be available in your Azure Stack Hub's Marketplace before it could be used by AKS Engine. Your Azure Stack Hub administrator can follow this [guide](https://docs.microsoft.com/azure-stack/operator/azure-stack-download-azure-marketplace-item) for a general explanation about how to download marketplace items from Azure.

Each AKS Engine release is validated and tied to a specific version of the AKS Base Image. Therefore, you need to take note of the base image version required by the AKS Engine release that you plan to use, and then download exactly that base image version. New builds of the `AKS Base Image` are frequently released to ensure that your disconnected cluster can be upgraded to the latest supported version of each component.

Make sure `runUnattendedUpgradesOnBootstrap` in the `linuxProfile` of your `apimodel.json` file is set to `"false"` when you create or upgrade cluster on disconnected Azure Stack Hub instances, otherwise the operation will fail.

## AKS Engine Versions

| AKS Engine                 | AKS Base Image     | Kubernetes versions | Notes |
|----------------------------|--------------------|---------------------|-------|
| [v0.43.1](https://github.com/Azure/aks-engine/releases/tag/v0.43.1)   | [AKS Base Ubuntu 16.04-LTS Image Distro, October 2019 (2019.10.24)](https://github.com/Azure/aks-engine/blob/v0.43.0/releases/vhd-notes/aks-ubuntu-1604/aks-ubuntu-1604-201910_2019.10.24.txt) | 1.15.5, 1.15.4, 1.14.8, 1.14.7 |  |
| [v0.48.0](https://github.com/Azure/aks-engine/releases/tag/v0.48.0)   | [AKS Base Ubuntu 16.04-LTS Image Distro, March 2020 (2020.03.19)](https://github.com/Azure/aks-engine/blob/v0.48.0/vhd/release-notes/aks-engine-ubuntu-1604/aks-engine-ubuntu-1604-202003_2020.03.19.txt) | 1.15.10, 1.14.7 |  |
| [v0.51.0](https://github.com/Azure/aks-engine/releases/tag/v0.51.0)   | [AKS Base Ubuntu 16.04-LTS Image Distro, May 2020 (2020.05.13)](https://github.com/Azure/aks-engine/blob/v0.51.0/vhd/release-notes/aks-engine-ubuntu-1604/aks-engine-ubuntu-1604-202005_2020.05.13.txt), [AKS Base Windows Image (17763.1217.200513)](https://github.com/Azure/aks-engine/blob/v0.51.0/vhd/release-notes/aks-windows/2019-datacenter-core-smalldisk-17763.1217.200513.txt)  | 1.15.12, 1.16.8, 1.16.9 | API Model Samples ([Linux](https://github.com/Azure/aks-engine/blob/v0.51.0/examples/azure-stack/kubernetes-azurestack.json), [Windows](https://github.com/Azure/aks-engine/blob/v0.51.0/examples/azure-stack/kubernetes-windows.json)) |
| [v0.55.0](https://github.com/Azure/aks-engine/releases/tag/v0.55.0)   | [AKS Base Ubuntu 16.04-LTS Image Distro, August 2020 (2020.08.24)](https://github.com/Azure/aks-engine/blob/v0.55.0/vhd/release-notes/aks-engine-ubuntu-1604/aks-engine-ubuntu-1604-202007_2020.08.24.txt), [AKS Base Windows Image (17763.1397.200820)](https://github.com/Azure/aks-engine/blob/v0.55.0/vhd/release-notes/aks-windows/2019-datacenter-core-smalldisk-17763.1397.200820.txt)  | 1.15.12, 1.16.14, 1.17.11 | API Model Samples ([Linux](https://github.com/Azure/aks-engine/blob/v0.55.0/examples/azure-stack/kubernetes-azurestack.json), [Windows](https://github.com/Azure/aks-engine/blob/v0.55.0/examples/azure-stack/kubernetes-windows.json)) |
| [v0.55.4](https://github.com/Azure/aks-engine/releases/tag/v0.55.4)   | [AKS Base Ubuntu 16.04-LTS Image Distro, September 2020 (2020.09.14)](https://github.com/Azure/aks-engine/blob/v0.55.0/vhd/release-notes/aks-engine-ubuntu-1604/aks-engine-ubuntu-1604-202007_2020.08.24.txt), [AKS Base Windows Image (17763.1397.200820)](https://github.com/Azure/aks-engine/blob/v0.55.0/vhd/release-notes/aks-windows/2019-datacenter-core-smalldisk-17763.1397.200820.txt)  | 1.15.12, 1.16.14, 1.17.11 | API Model Samples ([Linux](https://github.com/Azure/aks-engine/blob/v0.55.0/examples/azure-stack/kubernetes-azurestack.json), [Windows](https://github.com/Azure/aks-engine/blob/v0.55.0/examples/azure-stack/kubernetes-windows.json)) |
| [v0.60.1](https://github.com/Azure/aks-engine/releases/tag/v0.60.1)   | [AKS Base Ubuntu 18.04-LTS Image Distro, 2021 Q1 (2021.01.28)](https://github.com/Azure/aks-engine/blob/v0.60.1/vhd/release-notes/aks-engine-ubuntu-1804/aks-engine-ubuntu-1804-202007_2021.01.28.txt), [AKS Base Ubuntu 16.04-LTS Image Distro, January 2021 (2021.01.28)](https://github.com/Azure/aks-engine/blob/v0.60.1/vhd/release-notes/aks-engine-ubuntu-1604/aks-engine-ubuntu-1604-202007_2021.01.28.txt), [AKS Base Windows Image (17763.1697.210129)](https://github.com/Azure/aks-engine/blob/v0.60.1/vhd/release-notes/aks-windows/2109-datacenter-core-smalldisk-17763.1697.210129.txt)  | 1.16.14, 1.16.15, 1.17.17, 1.18.15 | API Model Samples ([Linux](https://github.com/Azure/aks-engine/blob/v0.60.1/examples/azure-stack/kubernetes-azurestack.json), [Windows](https://github.com/Azure/aks-engine/blob/v0.60.1/examples/azure-stack/kubernetes-windows.json)) |
| [v0.63.0](https://github.com/Azure/aks-engine/releases/tag/v0.63.0)   | [AKS Base Ubuntu 18.04-LTS Image Distro, 2021 Q2 (2021.05.24)](https://github.com/Azure/aks-engine/blob/v0.63.0/vhd/release-notes/aks-engine-ubuntu-1804/aks-engine-ubuntu-1804-202007_2021.05.24.txt), [AKS Base Windows Image (17763.1935.210520)](https://github.com/Azure/aks-engine/blob/v0.63.0/vhd/release-notes/aks-windows/2109-datacenter-core-smalldisk-17763.1935.210520.txt)  | 1.18.18, 1.19.10, 1.20.6 | API Model Samples ([Linux](https://github.com/Azure/aks-engine/blob/v0.65.0/examples/azure-stack/kubernetes-azurestack.json), [Windows](https://github.com/Azure/aks-engine/blob/v0.65.0/examples/azure-stack/kubernetes-windows.json)) |
| [v0.67.0](https://github.com/Azure/aks-engine/releases/tag/v0.67.0) | [AKS Base Ubuntu 18.04-LTS Image Distro, 2021 Q3 (2021.09.27)](https://github.com/Azure/aks-engine/blob/v0.67.0/vhd/release-notes/aks-engine-ubuntu-1804/aks-engine-ubuntu-1804-202007_2021.09.27.txt), [AKS Base Windows Image (17763.2213.210927)](https://github.com/Azure/aks-engine/blob/v0.67.0/vhd/release-notes/aks-windows/2019-datacenter-core-smalldisk-17763.2213.210927.txt) | 1.19.15, 1.20.11 | API Model Samples ([Linux](https://github.com/Azure/aks-engine/blob/master/examples/azure-stack/kubernetes-azurestack.json), [Windows](https://github.com/Azure/aks-engine/blob/master/examples/azure-stack/kubernetes-windows.json)) |
| [v0.67.3](https://github.com/Azure/aks-engine/releases/tag/v0.67.3)   | [AKS Base Ubuntu 18.04-LTS Image Distro, 2021 Q3 (2021.09.27)](https://github.com/Azure/aks-engine/blob/v0.67.3/vhd/release-notes/aks-engine-ubuntu-1804/aks-engine-ubuntu-1804-202007_2021.09.27.txt), [AKS Base Windows Image (17763.2213.210927)](https://github.com/Azure/aks-engine/blob/v0.67.3/vhd/release-notes/aks-windows/2019-datacenter-core-smalldisk-17763.2213.210927.txt)  | 1.19.15, 1.20.11 | API Model Samples ([Linux](https://github.com/Azure/aks-engine/blob/master/examples/azure-stack/kubernetes-azurestack.json), [Windows](https://github.com/Azure/aks-engine/blob/master/examples/azure-stack/kubernetes-windows.json)) |
## Azure Monitor for containers

Azure Monitor for containers can be deployed to AKS Engine clusters hosted in Azure Stack Hub Cloud Environments. Refer to [Azure Monitor for containers](../topics/monitoring.md#azure-monitor-for-containers) for more details on how to onboard and monitor clusters, nodes, pods, containers inventory, performance metrics and logs.

## Volume Provisioner: Container Storage Interface Drivers (preview)
As a [replacement of the current in-tree volume provisioner](https://kubernetes.io/blog/2019/12/09/kubernetes-1-17-feature-csi-migration-beta/), three Container Storage Interface (CSI) Drivers are avaiable on Azure Stack Hub. Please find details in the following table.

|                       | Azure Disk CSI Driver                                                                                                        | Azure Blob CSI Driver                                                                                                   | NFS CSI Driver                                                           |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------|
| Stage on Azure Stack  | Public Preview                                                                                                               | Private Preview                                                                                                         | Public Preview                                                           |
| Project Repository    | [azuredisk-csi-driver](https://github.com/kubernetes-sigs/azuredisk-csi-driver)                                              | [blob-csi-driver](https://github.com/kubernetes-sigs/blob-csi-driver)                                                   | [csi-driver-nfs](https://github.com/kubernetes-csi/csi-driver-nfs)       |
| CSI Driver Version    | v1.0.0+                                                                                                                      | v1.0.0+                                                                                                                 | v3.0.0+                                                                  |
| Access Mode           | ReadWriteOnce                                                                                                                | ReadWriteOnce<br/>ReadOnlyMany<br/>ReadWriteMany                                                                        | ReadWriteOnce<br/>ReadOnlyMany<br/>ReadWriteMany                         |
| Windows Agent Node    | Support                                                                                                                      | Not support and no plans                                                                                                | Not support and no plans                                                 |
| Dynamic Provisioning  | Support                                                                                                                      | Support                                                                                                                 | Support                                                                  |
| Considerations        | [Azure Disk CSI Driver Limitations](https://github.com/kubernetes-sigs/azuredisk-csi-driver/blob/master/docs/limitations.md) | [Azure Blob CSI Driver Limitations](https://github.com/kubernetes-sigs/blob-csi-driver/blob/master/docs/limitations.md) | Users will be responsible for setting up and maintaining the NFS server. |
| Slack Support Channel | [#provider-azure](https://kubernetes.slack.com/archives/C5HJXTT9Q)                                                           | [#provider-azure](https://kubernetes.slack.com/archives/C5HJXTT9Q)                                                      | [#sig-storage](https://kubernetes.slack.com/archives/C09QZFCE5)          |

> To deploy a CSI driver to an air-gapped cluster, make sure that your `helm` chart is referencing container images that are reachable from the cluster nodes.

### Requirements

- Azure Stack build 2011 and later.
- AKS Engine version v0.60.1 and later.
- Kubernetes version 1.18 and later.
- Since the Controller server of CSI Drivers requires 2 replicas, a single node master pool is not recommended.
- [Helm 3](https://helm.sh/docs/intro/install/)

### Install and Uninstall CSI Drivers
In this section, please follow the example commands to deploy a StatefulSet application consuming CSI Driver.

#### Azure Disk CSI Driver

``` powershell
# Install CSI Driver
helm repo add azuredisk-csi-driver https://raw.githubusercontent.com/kubernetes-sigs/azuredisk-csi-driver/master/charts
helm install azuredisk-csi-driver azuredisk-csi-driver/azuredisk-csi-driver --namespace kube-system --set cloud=AzureStackCloud --set controller.runOnMaster=true --version v1.0.0

# Deploy Storage Class
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/azuredisk-csi-driver/master/deploy/example/storageclass-azuredisk-csi-azurestack.yaml

# Deploy example StatefulSet application
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/azuredisk-csi-driver/master/deploy/example/statefulset.yaml

# Validate volumes and applications
# You should see a sequence of timestamps are persisted in the volume.
kubectl exec statefulset-azuredisk-0 -- tail /mnt/azuredisk/outfile

# Delete example StatefulSet application
kubectl delete -f https://raw.githubusercontent.com/kubernetes-sigs/azuredisk-csi-driver/master/deploy/example/statefulset.yaml

# Delete Storage Class
# Before delete the Storage Class, please make sure Pods that consume the Storage Class have been terminated.
kubectl delete -f https://raw.githubusercontent.com/kubernetes-sigs/azuredisk-csi-driver/master/deploy/example/storageclass-azuredisk-csi-azurestack.yaml

# Uninstall CSI Driver
helm uninstall azuredisk-csi-driver --namespace kube-system
helm repo remove azuredisk-csi-driver
```

#### Azure Blob CSI Driver

``` powershell
# Install CSI Driver
helm repo add blob-csi-driver https://raw.githubusercontent.com/kubernetes-sigs/blob-csi-driver/master/charts
helm install blob-csi-driver blob-csi-driver/blob-csi-driver --namespace kube-system --set cloud=AzureStackCloud --set controller.runOnMaster=true --version v1.0.0

# Deploy Storage Class
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/blob-csi-driver/master/deploy/example/storageclass-blobfuse.yaml

# Deploy example StatefulSet application
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/blob-csi-driver/master/deploy/example/statefulset.yaml

# Validate volumes and applications
# You should see a sequence of timestamps are persisted in the volume.
kubectl exec statefulset-blob-0 -- tail /mnt/blob/outfile

# Delete example StatefulSet application
kubectl delete -f https://raw.githubusercontent.com/kubernetes-sigs/blob-csi-driver/master/deploy/example/statefulset.yaml

# Delete Storage Class
# Before delete the Storage Class, please make sure Pods that consume the Storage Class have been terminated.
kubectl delete -f https://raw.githubusercontent.com/kubernetes-sigs/blob-csi-driver/master/deploy/example/storageclass-blobfuse.yaml

# Uninstall CSI Driver
helm uninstall blob-csi-driver --namespace kube-system
helm repo remove blob-csi-driver
```

#### NFS CSI Driver

``` powershell
# Install CSI Driver
helm repo add csi-driver-nfs https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/charts
helm install csi-driver-nfs csi-driver-nfs/csi-driver-nfs --namespace kube-system --set controller.runOnMaster=true --version v3.0.0

# Deploy NFS Server. Please note that this NFS Server is just for validation, please set up and maintain your NFS Server properly for production.
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/example/nfs-provisioner/nfs-server.yaml

# Deploy Storage Class
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/example/storageclass-nfs.yaml

# Deploy example StatefulSet application
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/example/statefulset.yaml

# Validate volumes and applications
# You should see a sequence of timestamps are persisted in the volume.
kubectl exec statefulset-nfs-0 -- tail /mnt/nfs/outfile

# Delete example StatefulSet application
kubectl delete -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/example/statefulset.yaml

# Delete Storage Class
# Before delete the Storage Class, please make sure Pods that consume the Storage Class have been terminated.
kubectl delete -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/example/storageclass-nfs.yaml

# Delete example NFS Server.
kubectl delete -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/example/nfs-provisioner/nfs-server.yaml

# Uninstall CSI Driver
helm uninstall csi-driver-nfs --namespace kube-system
helm repo remove csi-driver-nfs
```

## Known Issues and Limitations

This section lists all known issues you may find when you use the GA version.

### Unsupported Addons

AKS Engine includes a number of optional [addons](../topics/clusterdefinitions.md#addons) that can be deployed as part of the cluster provisioning process.

The list below includes the addons currently unsupported on Azure Stack Hub:

* AAD Pod Identity
* Blobfuse Flex Volume
* Cluster Autoscaler
* KeyVault Flex Volume
* SMB Flex Volume

### Limited Number of Frontend Public IPs

The `Basic` load balancer SKU available on Azure Stack Hub limits the number of frontend IPs to 5. That implies that each cluster's agents pool is limited to 5 public IPs.

If you need to expose more than 5 services, then the recommendation is to route traffic to those services using an Ingress controller.

### get-versions command

By default, `aks-engine get-versions` shows which Kubernetes versions are supported by each AKS Engine release on Azure's public cloud. Include flag `--azure-env` to get the list of supported Kubernetes versions on a custom cloud such as an Azure Stack Hub cloud (`aks-engine get-versions --azure-env AzureStackCloud`). Upgrade paths for Azure Stack Hub can also be found [here](https://docs.microsoft.com/azure-stack/user/kubernetes-aks-engine-release-notes).

### Upgrade from private-preview Kubernetes cluster with Windows nodes

There is no official support for private-preview Kubernetes cluster with Windows nodes created with AKS Engine v0.43.1 to upgrade with AKS Engine v0.55.0. Users are encouraged to deploy new Kubernetes cluster with Windows nodes with the latest AKS Engine version.

### Upgrading Kubernetes clusters created with the Ubuntu 16.04 distro

Starting with AKS Engine v0.63.0, the Ubuntu 16.04 distro is not longer a supported option as the OS reached its end-of-life. For AKS Engine v0.67.0 or later versions, aks-engine upgrade will automatically overwrite the unsupported `aks-ubuntu-16.04` distro value with with `aks-ubuntu-18.04`.

For AKS Engine v0.63.0, in order to upgrade a cluster, make sure to set the OS distro to `aks-ubuntu-18.04` in your input API model, the one generated by `aks-engine deploy`, and passed as input to `aks-engine upgrade`.

```json
    "masterProfile": {
      "distro": "aks-ubuntu-18.04"
    },

    "agentPoolProfiles": [
      {
        "distro": "aks-ubuntu-18.04"
      }
    ]
```

## Frequently Asked Questions

### Sample extensions are not working

Extensions in AKS Engine provide an easy way to include your own customization at provisioning time.

Because Azure and Azure Stack Hub currently rely on a different version of the Compute Resource Provider API, you may find that some of sample [extensions](https://github.com/Azure/aks-engine/tree/master/extensions) fail to deploy correctly.

This can be resolved by making a small modification to the extension `template.json` file. Replacing all usages of template parameter `apiVersionDeployments` by the hard-code value `2017-12-01` (or whatever API version Azure Stack Hub runs at the time you try to deploy) should be all you need.

Once you are done updating the extension template, host the extension directory in your own Github repository or storage account. Finally, at deployment time, make sure that your cluster definition points to the new [rootURL](https://github.com/Azure/aks-engine/blob/master/docs/topics/extensions.md#rooturl).

### Ubuntu VMs does not contains latest OS security fixes

The `aks-ubuntu-16.04` and `aks-ubuntu-18.04` base images provided in Azure Stack Hub's Marketplace are created and published prior to the release date of corresponding AKS-Engine versions that consumes it, and are not actively updated after published to Azure Stack Hub's Marketplace. Kubernetes clusters created with these images will not contain OS security fixes released after the AKS Base Images are published.

To obtain all latest OS security fixes for Ubunutu VMs in your cluster, you can set `runUnattendedUpgradesOnBootstrap` to `"true"` in the `linuxProfile` of your generated `apimodel.json` file, and run `aks-engine upgrade` command for the same Kubernetes version. Alternatively, you can manually run `apt-get update`, followed by a manual invocation of `/usr/bin/unattended-upgrade`, and then restart the VMs after these commands.

### Troubleshoting

This [how-to guide](/docs/howto/troubleshooting.md) has a good high-level explanation of how AKS Engine interacts with the Azure Resource Manager (ARM) and lists a few potential issues that can cause AKS Engine commands to fail.

Please refer to the [get-logs](../topics/get-logs.md) command documentation to simplify the logs collection task.

## Next Steps

* [What is the AKS engine on Azure Stack Hub?](https://docs.microsoft.com/azure-stack/user/azure-stack-kubernetes-aks-engine-overview)
