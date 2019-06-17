# AKS Engine on Azure Stack - Public Preview

* [Introduction](#introduction)
* [Service Principals and Identity Providers](#service-principals-and-identity-providers)
* [CLI flags](#cli-flags)
* [Cluster Definition (aka API Model)](#cluster-definition-aka-api-model)
    * [location](#location)
    * [kubernetesConfig](#kubernetesConfig)
    * [customCloudProfile](#customCloudProfile)
    * [masterProfile](#masterProfile)
    * [agentPoolProfiles](#agentPoolProfiles)
* [Azure Stack Instances Registered with Azure's China cloud](#azure-stack-instances-registered-with-azures-china-cloud)
* [Disconnected Azure Stack Instances](#disconnected-azure-stack-instances)
* [Unsupported Addons](#unsupported-addons)
* [Known Issues and Limitations](#known-issues-and-limitations)
* [Frequently Asked Questions](#frequently-asked-questions)

## Introduction to AKS Engine for Azure Stack

Starting from [release](https://github.com/Azure/aks-engine/releases/) v0.36.2, AKS Engine can be used to provision self-managed Kubernetes clusters on Azure Stack. It is now possible to execute AKS Engine's `generate`, `deploy`, `upgrade`, and `scale` commands as if you were targeting Azure's public cloud. You are only required to slightly update your cluster definition to provide some extra information about your Azure Stack instance.

The goal of this guide is to explain how to provision Kubernetes clusters to Azure Stack using AKS Engine and to capture the differences between Azure and Azure Stack. Bear in mind as well that not every AKS Engine feature or configuration option is currently supported on Azure Stack. In most cases, these are not available because dependent Azure components are not part of Azure Stack.

## Service principals and identity providers

Kubernetes uses a `service principal` identity to talk to Azure Stack APIs to dynamically manage resources such as storage or load balancers. Therefore, you will need to create a service principal before you can provision a Kubernetes cluster using AKS Engine.

This [guide](https://docs.microsoft.com/azure-stack/operator/azure-stack-create-service-principals) explains how to create and manage service principals on Azure Stack for both Azure Active Directory (Azure AD) and Active Directory Federation Services (AD FS) identity providers. This other [guide](../../docs/topics/service-principals.md) is a good resource to understand the permissions that the service principal requires to deploy under your subscription.

Once you have created the required service principal, make sure to assign it the `contributor` role at the target subscription scope.

## CLI flags

To indicate to AKS Engine that your target platform is Azure Stack, all commands require the CLI flag `azure-env` to be set to `"AzureStackCloud"`.

If your Azure Stack instance uses AD FS to authenticate identities, then the flag `identity-system` is also required.

``` bash
aks-engine deploy \
    --location local \
    --api-model ./kubernetes.json \
    --resource-group aks-engine-rg \
    --output-directory aks-engine \
    --client-id $SPN_CLIENT_ID \
    --client-secret $SPN_CLIENT_SECRET \
    --subscription-id $TENANT_SUBSCRIPTION_ID \
    --azure-env AzureStackCloud
```

## Cluster definition (also known as the API model)

This section details how to tailor your cluster definitions in order to make them compatible with Azure Stack. You can start off from this [template](../../examples/azure-stack/kubernetes-azurestack-azure-ad.json) if your identity provider is Azure AD or from this other [template](../../examples/azure-stack/kubernetes-azurestack-adfs.json) if you are using AD FS instead.

Unless otherwise specified down below, standard [cluster definition](../../docs/topics/clusterdefinitions.md) properties should also work with Azure Stack. Please create an [issue](https://github.com/Azure/aks-engine/issues/new) if you find that we missed a property that should be called out.

### location

| Name       | Required | Description                                                   |
| ---------- | -------- | ------------------------------------------------------------- |
| location   | yes      | The region name of the target Azure Stack. |

### kubernetesConfig

`kubernetesConfig` describes Kubernetes specific configuration.

| Name                            | Required | Description                          |
| ------------------------------- | -------- | ------------------------------------ |
| addons                          | no       | A few addons are not supported on Azure Stack. See the [complete list](#unsupported-addons) down below.|
| kubernetesImageBase             | yes      | Specifies the default image base URL to be used for all Kubernetes-related containers such as hyperkube, cloud-controller-manager, pause, addon-manager, etc. This property should be set to `"msazurestackdocker/"`. |
| networkPlugin                   | yes      | Specifies the network plugin implementation for the cluster. Valid values are `"kubenet"` for Kubernetes software networking implementation, and `"flannel"` to use CoreOS Flannel. |
| networkPolicy                   | no       | Network policies can be enforced using [Canal](https://docs.projectcalico.org/v3.7/getting-started/kubernetes/installation/flannel). |
| useInstanceMetadata             | yes      | Use the Azure cloud provider instance metadata service for appropriate resource discovery operations. This property should be always set to `"false"`. |

### customCloudProfile

`customCloudProfile` contains information specific to the target Azure Stack instance.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| environment                     | no       | The custom cloud type. This property should be always set to `"AzureStackCloud"`. |
| identitySystem                  | no       | Specifies the identity provider used by the Azure Stack instance. Valid values are `"azure_ad"` (default) and `"adfs"`. |
| portalUrl                       | yes      | The tenant portal URL. |
| dependenciesLocation            | no       | Specifies where to locate the dependencies required to during the provision/upgrade process. Valid values are `"public"` (default), `"china"`, `"german"` and `"usgovernment".` |

### masterProfile

`masterProfile` describes the settings for master configuration.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Specifies a valid [Azure Stack VM size](https://docs.microsoft.com/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. Currently supported values are: `"ubuntu"` and `"aks"`. The latter is a custom image based on ubuntu-16.04 that comes with pre-installed software necessary for Kubernetes deployments. `"aks"` should be selected if you target a disconnected Azure Stack instance. |

### agentPoolProfiles

`agentPoolProfiles` are used to create agents with different capabilities.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Describes a valid [Azure Stack VM size](https://docs.microsoft.com/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. Currently supported values are: `"ubuntu"` and `"aks"`. The latter is a custom image based on ubuntu-16.04 that comes with pre-installed software necessary for Kubernetes deployments. `"aks"` should be selected if you target a disconnected Azure Stack instance. |
| availabilityProfile             | yes      | Only `"AvailabilitySet"` is currently supported. |
| acceleratedNetworkingEnabled    | yes      | Use `Azure Accelerated Networking` feature for Linux agents. This property should be always set to `"false"`. |

## Azure Stack instances registered with Azure's China cloud

If your Azure Stack instance is located in China, then the `dependenciesLocation` property of your cluster definition should be set to `"china"`. This switch ensures that the provisioning process fetches software dependencies from reachable hosts within China's mainland.

## Disconnected Azure Stack instances

_Note: AKS Engine on disconnected Azure Stack instances is a private preview feature_

By default, AKS Engine provisioning process relies on an internet connection to download the software dependencies required to create or upgrade a cluster (Kubernetes images, etcd binaries, network plugins and so on).

If your Azure Stack instance is air-gapped or if network connectivity in your geographical location is not reliable, then the default approach will not work due to the length of time for the download or a timeout due to transient networking issues.

With these challenges in mind, you can choose to set the `distro` property of your cluster definition to `"aks"`. This change will instruct AKS Engine to deploy VM nodes using a base OS image called `AKS Base Image`. This custom image, generally based on Ubuntu Server, already contains the required software dependencies in its file system. Hence, internet connectivity won’t be required during the provisioning process.

The `AKS Base Image` gallery item has to be available in your Azure Stack's Marketplace before it can be used by AKS Engine. Your Azure Stack administrator can follow this [guide](https://docs.microsoft.com/azure-stack/operator/azure-stack-download-azure-marketplace-item) for a general explanation about how to download gallery items from Azure.

Each AKS Engine release is validated and tied to a specific version of AKS Base Image. Therefore, you need to take note of the base image version required by AKS Engine release that you plan to use, and then download exactly that base image version. New builds of the `AKS Base Image` are frequently released to ensure that your disconnected cluster can be upgraded to the latest supported version of each component.

## Unsupported addons

AKS Engine includes a number of optional [addons](../topics/clusterdefinitions.md#addons) that can be deployed as part of the cluster provisioning process.

The list below includes the addons currently unsupported on Azure Stack:

* AAD Pod Identity
* ACI Connector
* Blobfuse Flex Volume
* Cluster Autoscaler
* Container Monitoring
* KeyVault Flex Volume
* NVIDIA Device Plugin
* Rescheduler
* SMB Flex Volume

## Known Issues and Limitations

This section lists all known issues you may find when you use the public preview version.

### Agent nodes internet connectivity

Your agent nodes may lose internet connectivity after all Kubernetes services of type `LoadBalancer` are deleted. You are not expected to experience this problem if no services of type `LoadBalancer` are ever created.

To work around this issue, do not delete `LoadBalancer` services as part of your release pipeline or always keep a dummy service.

### Limited number of frontend public IPs

The `Basic` load balancer SKU available on Azure Stack limits the number of frontend IPs to five (5). That implies that each cluster's agents pool is limited to five (5) public IPs.

If you need to expose more than five (5) services, then the recommendation is to route traffic to those services using an ingress controller.

## Frequently asked questions

### Supported Kubernetes versions

These are the Kubernetes versions that you can deploy to Azure Stack using AKS Engine:

<!-- - 1.14.3 build and push to mcr -->
<!-- - 1.13.7 build and push to mcr -->s
- 1.14.1
- 1.13.5
- 1.12.8
- 1.12.7
- 1.11.10
- 1.11.9


### Network policies

If the Azure Stack instance is air-gapped, then the syndicated [AKS Base Image](#disconnected-azure-stack-instances) will determine the Kubernetes version that you can deploy or the target upgrade version. The table below lists the pre-pulled versions on each AKS Base Image version.

_Note: AKS Engine on disconnected Azure Stack instances is a private preview feature_

| AKS Engine                 | AKS Base Image     | Kubernetes versions | Notes |  
|----------------------------|--------------------|---------------------|-------|
| from v0.36.2 to v0.36.5    | [AKS Base Ubuntu 16.04-LTS Image Distro, May 2019](../../releases/vhd-notes/aks-ubuntu-1604/aks-ubuntu-1604-201904_2019.05.16.txt) | 1.12.7 - 1.11.10 - 1.11.9 | Only network plugin `"kubenet"` is supported |


To enforce network policies, you are required to manually deploy the [Canal](https://docs.projectcalico.org/v3.7/getting-started/kubernetes/installation/flannel) daemonset.

### Sample extensions are not working

Extensions in AKS Engine provide an easy way to include your own customization at provisioning time.

Because Azure and Azure Stack currently rely on a different version of the Compute Resource Provider API, you may find that some of sample [extensions](https://github.com/Azure/aks-engine/tree/master/extensions) fail to deploy correctly.

This can be resolved by making a small modification to the extension `template.json` file. Replacing all usages of template parameter `apiVersionDeployments` by the hard-code value `2017-12-01` (or whatever API version Azure Stack runs at the time you try to deploy) should be all you need.

Once you are done updating the extension template, host the extension directory in your own Github repository or storage account. Finally, at deployment time, make sure that your cluster definition points to the new [rootURL](https://github.com/Azure/aks-engine/blob/master/docs/topics/extensions.md#rooturl).

### Troubleshoting

This [how-to guide](/docs/howto/troubleshooting.md) has a good high-level explanation of how AKS Engine interacts with the Azure Resource Manager (ARM) and lists a few potential issues that can cause AKS Engine commands to fail. Checkout the scripts in this [repository](https://github.com/msazurestackworkloads/azurestack-gallery/tree/master/diagnosis) to simplify the logs collection task.
