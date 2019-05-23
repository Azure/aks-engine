# AKS Engine on Azure Stack

<!-- vscode-markdown-toc -->
* [Introduction](#Introduction)
* [Service Principals and Identity Providers](#ServicePrincipalsandIdentityProviders)
* [CLI flags](#CLIflags)
* [Cluster Definition (aka API Model)](#ClusterDefinitionakaAPIModel)
	* [location](#location)
	* [kubernetesConfig](#kubernetesConfig)
	* [customCloudProfile](#customCloudProfile)
	* [masterProfile](#masterProfile)
	* [agentPoolProfiles](#agentPoolProfiles)
* [Azure Stack Instances Registered with Azure's China cloud](#AzureStackInstancesRegisteredwithAzuresChinacloud)
* [Disconnected Azure Stack Instances](#DisconnectedAzureStackInstances)
* [Unsupported Addons](#UnsupportedAddons)

<!-- vscode-markdown-toc-config
	numbering=false
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

## <a name='Introduction'></a>Introduction

Starting from release 0.36.XX, AKS Engine can be used to provision self-managed Kubernetes clusters on Azure Stack. It is now possible to execute AKS Engine's `generate`, `deploy`, `upgrade`, and `scale` commands as if you were targeting Azure's public cloud. You are only required to slighly update your cluster definition to provide some extra information about your Azure Stack instance.

The goal of this guide is to explain how to provision Kubernetes clusters to Azure Stack using AKS Engine and to capture the differences between Azure and Azure Stack. Bear in mind as well that not every AKS Engine feature or configuration option is currently supported on Azure Stack. In most cases, these are not available because dependent Azure components are not part of Azure Stack.

## <a name='ServicePrincipalsandIdentityProviders'></a>Service Principals and Identity Providers

Kubernetes uses a `service principal` identity to talk to Azure Stack APIs to dynamically manage resources such a storage or load balances. Therefore, you will need to create a service principal before you can provision a Kubernetes cluster using AKS Engine.
This [guide](https://docs.microsoft.com/en-us/azure-stack/operator/azure-stack-create-service-principals) explains how to create and manage service principals on Azure Stack for both Azure Active Directory (AAD) and Active Directory Federation Services (ADFS) identity providers. This other [guide](../../docs/topics/service-principals.md) is a good resource to understand the permissions that the service principal requires to deploy under your subscription.


## <a name='CLIflags'></a>CLI flags

To indicate AKS Engine that your target platform is Azure Stack, all commands require CLI flag `azure-env` to be set to `"AzureStackCloud"`.

```
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

## <a name='ClusterDefinitionakaAPIModel'></a>Cluster Definition (aka API Model)

This section details how to tailor your cluster definitions in order to make them compatible with Azure Stack. You can start off from this [template](../../examples/azure-stack/kubernetes-azurestack-azure-ad.json) if your identity provier is AAD or this other [template](../../examples/azure-stack/kubernetes-azurestack-adfs.json) if you are using ADFS instead.

Unless otherwise specified down below, standard [cluster definition](../../docs/topics/clusterdefinitions.md) properties should also work with Azure Stack. Please create an [issue](https://github.com/Azure/aks-engine/issues/new) if you find that we missed something.

### <a name='location'></a>location

| Name       | Required | Description                                                   |
| ---------- | -------- | ------------------------------------------------------------- |
| location   | yes      | The region name of the target Azure Stack. |

### <a name='kubernetesConfig'></a>kubernetesConfig

`kubernetesConfig` describes Kubernetes specific configuration.

| Name                            | Required | Description                          |
| ------------------------------- | -------- | ------------------------------------ |
| addons                          | no       | A few addons are not supported on Azure Stack. See the [complete list]() down below.|
| kubernetesImageBase             | yes      | Specifies the default image base URL to be used for all kubernetes-related containers such as hyperkube, cloud-controller-manager, pause, addon-manager, etc. This property should be set to `"msazurestackdocker/"`.|
| networkPlugin                   | yes      | Specifies the network plugin implementation for the cluster. Valid values are:<br>`"kubenet"` for k8s software networking implementation <br> `"flannel"` for using CoreOS Flannel. |
| networkPolicy                   | no       | Network policies can be enforced using [Canal](). **TODO EXPAND** |
| useInstanceMetadata             | yes      | Use the Azure cloudprovider instance metadata service for appropriate resource discovery operations. This property should be always set to `"false"`. |

### <a name='customCloudProfile'></a>customCloudProfile

`customCloudProfile` contains information specific to the target Azure Stack instance.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| environment                     | no       | The custom cloud type. This property should be always set to `"AzureStackCloud"`.|
| identitySystem                  | no       | Specifies the identity provider used by the Azure Stack instance. Valid values are `"azure_ad"` (default) and `"adfs"`. |
| portalUrl                       | yes      | The tenant portal URL. |
| dependenciesLocation            | no       | Specifies where to locate the dependencies required to during the provision/upgrade process. Valid values are `"public"` (default), `"china"`, `"german"` and `"usgovernment".`|

### <a name='masterProfile'></a>masterProfile

`masterProfile` describes the settings for master configuration.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Specifies a valid [Azure Stack VM size](https://docs.microsoft.com/en-us/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. Currently supported values are: `"ubuntu"` and `"aks-ubuntu-16.04"` (previously `"aks"`). `"aks-ubuntu-16.04"` is a custom image based on ubuntu-16.04 that comes with pre-installed software necessary for Kubernetes deployments. `"aks-ubuntu-16.04"` should be selected if you target a disconnected Azure Stack instance. |

### <a name='agentPoolProfiles'></a>agentPoolProfiles

`agentPoolProfiles` are used to create agents with different capabilities.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| vmsize                          | yes      | Describes a valid [Azure Stack VM size](https://docs.microsoft.com/en-us/azure-stack/user/azure-stack-vm-sizes). |
| distro                          | yes      | Specifies the masters' Linux distribution. Currently supported values are: `"ubuntu"` and `"aks-ubuntu-16.04"` (previously `"aks"`). `"aks-ubuntu-16.04"` is a custom image based on ubuntu-16.04 that comes with pre-installed software necessary for Kubernetes deployments. `"aks-ubuntu-16.04"` should be selected if you target a disconnected Azure Stack instance. |
| availabilityProfile             | yes      | Only `"AvailabilitySet"` is currently supported. |
| acceleratedNetworkingEnabled    | yes      | Use `Azure Accelerated Networking` feature for Linux agents. This property should be always set to `"true"`. |

## <a name='AzureStackInstancesRegisteredwithAzuresChinacloud'></a>Azure Stack Instances Registered with Azure's China cloud

If your Azure Stack instance is located in China, then the `dependenciesLocation` property on your cluster definition should be set to `"china"`. This switch ensures that the provisioning process fetches software dependencies from reachable hosts.

## <a name='DisconnectedAzureStackInstances'></a>Disconnected Azure Stack Instances

By default, the AKS Engine provisioning process relies on an internet connection to download the software dependencies required to create or upgrade a cluster (kubernetes images, etcd binaries, network plugins and so on).

If your Azure Stack instance is air-gapped or if network connectivity in your geographical location is not reliable, then the default approach will not work, take a long time or timeout due to transcient networking issues.

With these challenges in mind, you can choose to set the `distro` property on your cluster definition to `"aks-ubuntu-16.04"`. This change will instruct AKS Engine to deploy nodes using a base OS image that already contains the required software dependencies in its file system. Therefore, avoiding any need of reaching out to the internet.

**TODO EXPLAIN HOW TO SYNDICATE THE IMAGE**

## <a name='UnsupportedAddons'></a>Unsupported Addons

AKS Engine includes a number of optional [addons](../topics/clusterdefinitions.md#addons) that can be deployed as part of the cluster provisioning process.

The list below includes the addons currently unsupported on Azure Stack.

* AAD Pod Identity
* ACI Connector
* Cluster Autoscaler
* Blobfuse Flex Volume
* SMB Flex Volume
* KeyVault Flex Volume
* Rescheduler
* NVIDIA Device Plugin
* Container Monitoring
