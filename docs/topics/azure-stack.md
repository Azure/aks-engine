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
  * [agentPoolProfiles](#agentPoolProfiles)ß
* [Azure Stack Hub Instances Registered with Azure's China cloud](#azure-stack-instances-registered-with-azures-china-cloud)
* [Disconnected Azure Stack Hub Instances](#disconnected-azure-stack-instances)
* [Supported AKS Engine Versions](#supported-aks-engine-versions)
* [Azure Monitor for containers](#azure-Monitor-for-containers)
* [Known Issues and Limitations](#known-issues-and-limitations)
* [Frequently Asked Questions](#frequently-asked-questions)

## Introduction

Specific AKS Engine [versions](#supported-aks-engine-versions) can be used to provision self-managed Kubernetes clusters on [Azure Stack Hub](https://azure.microsoft.com/overview/azure-stack/). AKS Engine's `generate`, [deploy](../tutorials/deploy.md), [upgrade](upgrade.md), and [scale](scale.md) commands can be executed as if you were targeting Azure's public cloud. You are only required to slightly update your cluster definition to provide some extra information about your Azure Stack Hub instance.

The goal of this guide is to explain how to provision Kubernetes clusters to Azure Stack Hub using AKS Engine and to capture the differences between Azure and Azure Stack Hub. Bear in mind as well that not every AKS Engine feature or configuration option is currently supported on Azure Stack Hub. In most cases, these are not available because dependent Azure components are not part of Azure Stack Hub.

## Marketplace prerequisites

Because Azure Stack Hub instances do not have infinite storage available, Azure Stack Hub administrators are in charge of managing it by selecting which marketplace items are downloaded from Azure's marketplace. The Azure Stack Hub administrator can follow this [guide](https://docs.microsoft.com/azure-stack/operator/azure-stack-download-azure-marketplace-item) for a general explanation about how to download marketplace items from Azure.

Before you try to deploy the first Kubernetes cluster, make sure these marketplace items were made available to the target subscription by the Azure Stack Hub administrator.

* `Custom Script for Linux 2.0` virtual machine extension
* [Required](#supported-aks-engine-versions) `AKS Base Image` virtual machine

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
    --identity-system adfs \ # Optional if azure_ad
    --client-id $SPN_CLIENT_ID \
    --client-secret $SPN_CLIENT_SECRET \
    --subscription-id $TENANT_SUBSCRIPTION_ID \
    --output-directory kube-rg
```

## Cluster Definition (aka API Model)

This section details how to tailor your cluster definitions in order to make them compatible with Azure Stack Hub. You can start off from this [template](../../examples/azure-stack/kubernetes-azurestack-azure-ad.json) if your identity provider is AAD or from this other [template](../../examples/azure-stack/kubernetes-azurestack-adfs.json) if you are using ADFS instead.

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
| networkPlugin                   | yes      | Specifies the network plugin implementation for the cluster. Valid values are `"kubenet"` for Kubernetes software networking implementation. |
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

`masterProfile` describes the settings for master configuration.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Specifies a valid [Azure Stack Hub VM size](https://docs.microsoft.com/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. The supported value is `"aks-ubuntu-16.04"`. This is a custom image based on UbuntuServer 16.04 that comes with pre-installed software necessary for Kubernetes deployments. |

### agentPoolProfiles

`agentPoolProfiles` are used to create agents with different capabilities.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Describes a valid [Azure Stack Hub VM size](https://docs.microsoft.com/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. The supported value is `"aks-ubuntu-16.04"`. This is a custom image based on UbuntuServer 16.04 that comes with pre-installed software necessary for Kubernetes deployments. |
| availabilityProfile             | yes      | Only `"AvailabilitySet"` is currently supported. |
| acceleratedNetworkingEnabled    | yes      | Use `Azure Accelerated Networking` feature for Linux agents. This property should be always set to `"false"`. |

## Azure Stack Hub Instances Registered with Azure's China cloud

If your Azure Stack Hub instance is located in China, then the `dependenciesLocation` property of your cluster definition should be set to `"china"`. This switch ensures that the provisioning process fetches software dependencies from reachable hosts within China's mainland.

## Disconnected Azure Stack Hub Instances

By default, the AKS Engine provisioning process relies on an internet connection to download the software dependencies required to create or upgrade a cluster (Kubernetes images, etcd binaries, network plugins and so on).

If your Azure Stack Hub instance is air-gapped or if network connectivity in your geographical location is not reliable, then the default approach will not work, take a long time or timeout due to transient networking issues.

To overcome these issues, you should set the `distro` property of your cluster definition to `"aks-ubuntu-16.04"`. This will instruct AKS Engine to deploy VM nodes using a base OS image called `AKS Base Image`. This custom image, generally based on Ubuntu Server, already contains the required software dependencies in its file system. Hence, internet connectivity won't be required during the provisioning process.

The `AKS Base Image` marketplace item has to be available in your Azure Stack Hub's Marketplace before it could be used by AKS Engine. Your Azure Stack Hub administrator can follow this [guide](https://docs.microsoft.com/azure-stack/operator/azure-stack-download-azure-marketplace-item) for a general explanation about how to download marketplace items from Azure.

Each AKS Engine release is validated and tied to a specific version of the AKS Base Image. Therefore, you need to take note of the base image version required by the AKS Engine release that you plan to use, and then download exactly that base image version. New builds of the `AKS Base Image` are frequently released to ensure that your disconnected cluster can be upgraded to the latest supported version of each component.

## Supported AKS Engine Versions

| AKS Engine                 | AKS Base Image     | Kubernetes versions | Notes |
|----------------------------|--------------------|---------------------|-------|
| [v0.43.1](https://github.com/Azure/aks-engine/releases/tag/v0.43.1)   | [AKS Base Ubuntu 16.04-LTS Image Distro, October 2019 (2019.10.24)](https://github.com/Azure/aks-engine/blob/v0.43.0/releases/vhd-notes/aks-ubuntu-1604/aks-ubuntu-1604-201910_2019.10.24.txt) | 1.15.5, 1.15.4, 1.14.8, 1.14.7 |  |
| [v0.48.0](https://github.com/Azure/aks-engine/releases/tag/v0.48.0)   | [AKS Base Ubuntu 16.04-LTS Image Distro, March 2020 (2020.03.19)](https://github.com/Azure/aks-engine/blob/v0.48.0/vhd/release-notes/aks-engine-ubuntu-1604/aks-engine-ubuntu-1604-202003_2020.03.19.txt) | 1.15.10, 1.14.7 | [Template](../../examples/azure-stack/kubernetes-azurestack.json) |

## Azure Monitor for containers

Azure Monitor for containers can be deployed to AKS Engine clusters hosted in Azure Stack Hub Cloud Environments. Refer to [Azure Monitor for containers](../topics/monitoring.md#azure-monitor-for-containers) for more details on how to onboard and monitor clusters, nodes, pods, containers inventory, performance metrics and logs.

## Known Issues and Limitations

This section lists all known issues you may find when you use the GA version.

### Unsupported Addons

AKS Engine includes a number of optional [addons](../topics/clusterdefinitions.md#addons) that can be deployed as part of the cluster provisioning process.

The list below includes the addons currently unsupported on Azure Stack Hub:

* AAD Pod Identity
* ACI Connector
* Blobfuse Flex Volume
* Cluster Autoscaler
* KeyVault Flex Volume
* NVIDIA Device Plugin
* Rescheduler
* SMB Flex Volume

### Agent Nodes Internet Connectivity

Your agent nodes may lose internet connectivity after all Kubernetes services of type `LoadBalancer` are deleted. You are not expected to experience this problem if no services of type `LoadBalancer` are ever created.

To work around this issue, do not delete `LoadBalancer` services as part of your release pipeline or always keep a dummy service.

### Limited Number of Frontend Public IPs

The `Basic` load balancer SKU available on Azure Stack Hub limits the number of frontend IPs to 5. That implies that each cluster's agents pool is limited to 5 public IPs.

If you need to expose more than 5 services, then the recommendation is to route traffic to those services using an Ingress controller.

### Single agent pool

Currently, Azure Stack Hub only supports the `Basic` load balancer SKU.
This SKU [limits](https://docs.microsoft.com/en-us/azure/load-balancer/concepts-limitations#skus)
the backend pool endpoints to virtual machines in a single availability set (or virtual machine scale set).
This implies that all replicas of a `LoadBalancer` service should be deployed on the same agent pool
and it also implies that each individual cluster can either have a Linux `LoadBalancer` service or
a Windows `LoadBalancer` service.

You can force Kubernetes to create pods in a specific agent pool by adding [node selector](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/) "`agentpool: MY_POOL_NAME`" in your pod template.

```yaml
nodeSelector:
  agentpool: linuxpool
```

If a `LoadBalancer` service was already created in your cluster, you can find out which agent pool was selected as the load balancer's backend pool by inspecting the load balancer `backend pools` blade in the Azure Stack Hub portal. Once you have that information, you can specify the target agent pool by updating your deployment/pod yaml (as explained in the previous paragraph).

### get-versions command

The output of the `get-versions` command only pertains to Azure and not Azure Stack Hub clouds. The different upgrade paths can be found [here](https://docs.microsoft.com/azure-stack/user/azure-stack-kubernetes-aks-engine-upgrade#steps-to-upgrade-to-a-newer-kubernetes-version).

## Frequently Asked Questions

### Sample extensions are not working

Extensions in AKS Engine provide an easy way to include your own customization at provisioning time.

Because Azure and Azure Stack Hub currently rely on a different version of the Compute Resource Provider API, you may find that some of sample [extensions](https://github.com/Azure/aks-engine/tree/master/extensions) fail to deploy correctly.

This can be resolved by making a small modification to the extension `template.json` file. Replacing all usages of template parameter `apiVersionDeployments` by the hard-code value `2017-12-01` (or whatever API version Azure Stack Hub runs at the time you try to deploy) should be all you need.

Once you are done updating the extension template, host the extension directory in your own Github repository or storage account. Finally, at deployment time, make sure that your cluster definition points to the new [rootURL](https://github.com/Azure/aks-engine/blob/master/docs/topics/extensions.md#rooturl).

### Troubleshoting

This [how-to guide](/docs/howto/troubleshooting.md) has a good high-level explanation of how AKS Engine interacts with the Azure Resource Manager (ARM) and lists a few potential issues that can cause AKS Engine commands to fail. Checkout the scripts in this [repository](https://github.com/msazurestackworkloads/azurestack-gallery/tree/master/diagnosis) to simplify the logs collection task.

## Next Steps

* [What is the AKS engine on Azure Stack Hub?](https://docs.microsoft.com/azure-stack/user/azure-stack-kubernetes-aks-engine-overview)
