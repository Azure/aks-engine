# AKS Engine on Azure Stack

<!-- TODO: TOC -->

<!-- Intro -->
Since release 0.36.XX, AKS Engine can be used to provision self-managed Kubernetes clusters on Azure Stack. You only need to slighly update your cluster definition to provide some extra information about your Azure Stack instance.

Bear in mind as well that not every AKS Engine feature or configuration option is currently supported on Azure Stack. In most cases, these are not yet available because a required Azure service did not make it to Azure Stack at this point in time. **TODO LINK TO LIST**

The goal of this guide is to explain how to provision Kubernetes clusters to Azure Stack using AKS Engine and to capture the differences between Azure and Azure Stack.

## Service Principals and Identity Providers

Kubernetes uses a Service Principal to talk to Azure Stack APIs to dynamically manage resources such a storage or load balances. Therefore, you will need to create a Service Principal before you can provision a Kubernetes cluster using AKS Engine.

This [guide](https://docs.microsoft.com/en-us/azure-stack/operator/azure-stack-create-service-principals) explains how to create and manage Service Principals on Azure Stack for both Azure Active Directory (AAD) and Active Directory Federation Services (ADFS). This other [guide](https://github.com/Azure/aks-engine/blob/master/docs/topics/service-principals.md) is a good resource to understand the permissions that the Service Principal requires to deploy under your subscription.

## Cluster Definition

This section explains how your cluster definition should change to deploy to Azure Stack. These templates can be used as reference when deploying to an [AAD]() or an [ADFS]() instance.

### Location

| Name       | Required | Description                                                   |
| ---------- | -------- | ------------------------------------------------------------- |
| location | yes      | **TODO** |

### kubernetesConfig

`kubernetesConfig` describes Kubernetes specific configuration.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| addons                          | no       | A few addons are not supported on Azure Stack. See `addons` below.|
| kubernetesImageBase             | yes      | Specifies the default image base URL to be used for all kubernetes-related containers such as hyperkube, cloud-controller-manager, pause, addon-manager, etc. This property value should be set to `"msazurestackdocker"/`|
| networkPlugin                   | yes      | Specifies the network plugin implementation for the cluster. Valid values are:<br>`"kubenet"` for k8s software networking implementation <br> `"flannel"` for using CoreOS Flannel |
| networkPolicy                   | no       | Network policies can be enforced using Canal **TODO ADD MORE** |
| useInstanceMetadata             | yes      | Use the Azure cloudprovider instance metadata service for appropriate resource discovery operations. This property value should be set to `"true"` |

### customCloudProfile

`customCloudProfile` contains information specific to the target Azure Stack instance.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| environment                     | ??       | The custom cloud name. This property value should be set to `"AzureStackCloud"`.|
| identitySystem                  | yes      | Specifies the identity provider used by the Azure Stack instance. Valid values are `"azure_ad"` (default) and `"adfs"`. |
| authenticationMethod            | yes      | Specifies the method used by Kubernetes to authenticate the Service Principal. Valid values are: `"client_secret"` (default) and `"client_certificate"`. |
| dependenciesLocation            | no       | Network policies can be enforced using Canal **TODO ADD MORE** |
| portalUrl                       | yes      | ?? |
| dependenciesLocation            | no       | Specifies where to locate the dependencies required to during the provision/upgrade proess. Valid values are `"public"` (default), `"china"`, `"german"` and `"usgovernment".`|

### masterProfile

`masterProfile` describes the settings for master configuration.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Describes a valid [Azure Stack VM Sizes](https://docs.microsoft.com/en-us/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. Currently supported values are: `"ubuntu"` and `"aks-ubuntu-16.04"` (previously `"aks"`). `"aks-ubuntu-16.04"` is a custom image based on ubuntu-16.04 that comes with pre-installed software necessary for Kubernetes deployments. `"aks-ubuntu-16.04"` should be selected if you target a disconnected Azure Stack. |

### agentPoolProfiles

`agentPoolProfiles` are used to create agents with different capabilities.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Describes a valid [Azure Stack VM Sizes](https://docs.microsoft.com/en-us/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. Currently supported values are: `"ubuntu"` and `"aks-ubuntu-16.04"` (previously `"aks"`). `"aks-ubuntu-16.04"` is a custom image based on ubuntu-16.04 that comes with pre-installed software necessary for Kubernetes deployments. `"aks-ubuntu-16.04"` should be selected if you target a disconnected Azure Stack. |
| availabilityProfile             | yes      | Only `AvailabilitySet` is currently supported. |
| acceleratedNetworkingEnabled    | yes      | Use `Azure Accelerated Networking` feature for Linux agents. This property should be set to Defaults to `"true"` |

## Azure Stack instances registered with Mooncake AAD

## Disconnected or air-gapped Azure Stack instances

## Extra CLI parameters

## Addons
