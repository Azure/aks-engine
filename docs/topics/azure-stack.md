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
* [Known Issues](#KnownIssues)
* [Frequently Asked Questions](#FrequentlyAskedQuestions)
	* [Network Policies](#NetworkPolicies)

<!-- vscode-markdown-toc-config
	numbering=false
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

## <a name='Introduction'></a>Introduction

Starting from release 0.36.XX, AKS Engine can be used to provision self-managed Kubernetes clusters on Azure Stack. It is now possible to execute AKS Engine's `generate`, `deploy`, `upgrade`, and `scale` commands as if you were targeting Azure's public cloud. You are only required to sligthly update your cluster definition to provide some extra information about your Azure Stack instance.

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

This section details how to tailor your cluster definitions in order to make them compatible with Azure Stack. You can start off from this [template](../../examples/azure-stack/kubernetes-azurestack-azure-ad.json) if your identity provider is AAD or from this other [template](../../examples/azure-stack/kubernetes-azurestack-adfs.json) if you are using ADFS instead.

Unless otherwise specified down below, standard [cluster definition](../../docs/topics/clusterdefinitions.md) properties should also work with Azure Stack. Please create an [issue](https://github.com/Azure/aks-engine/issues/new) if you find that we missed a property that should be called out.

### <a name='location'></a>location

| Name       | Required | Description                                                   |
| ---------- | -------- | ------------------------------------------------------------- |
| location   | yes      | The region name of the target Azure Stack. |

### <a name='kubernetesConfig'></a>kubernetesConfig

`kubernetesConfig` describes Kubernetes specific configuration.

| Name                            | Required | Description                          |
| ------------------------------- | -------- | ------------------------------------ |
| addons                          | no       | A few addons are not supported on Azure Stack. See the [complete list](#UnsupportedAddons) down below.|
| kubernetesImageBase             | yes      | Specifies the default image base URL to be used for all Kubernetes-related containers such as hyperkube, cloud-controller-manager, pause, addon-manager, etc. This property should be set to `"msazurestackdocker/"`. **TODO: LOG ISSUE** |
| networkPlugin                   | yes      | Specifies the network plugin implementation for the cluster. Valid values are `"kubenet"` for Kubernetes software networking implementation, and `"flannel"` for using CoreOS Flannel. |
| networkPolicy                   | no       | Network policies can be enforced using [Canal](https://docs.projectcalico.org/v3.7/getting-started/kubernetes/installation/flannel). **TODO: LOG ISSUE** |
| useInstanceMetadata             | yes      | Use the Azure cloud provider instance metadata service for appropriate resource discovery operations. This property should be always set to `"false"`. **TODO: LOG ISSUE** |

### <a name='customCloudProfile'></a>customCloudProfile

`customCloudProfile` contains information specific to the target Azure Stack instance.

| Name                            | Required | Description|
| ------------------------------- | -------- | ---------- |
| environment                     | no       | The custom cloud type. This property should be always set to `"AzureStackCloud"`. **TODO: LOG ISSUE** |
| identitySystem                  | no       | Specifies the identity provider used by the Azure Stack instance. Valid values are `"azure_ad"` (default) and `"adfs"`. |
| portalUrl                       | yes      | The tenant portal URL. |
| dependenciesLocation            | no       | Specifies where to locate the dependencies required to during the provision/upgrade process. Valid values are `"public"` (default), `"china"`, `"german"` and `"usgovernment".` **TODO: LOG ISSUE** |

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
| acceleratedNetworkingEnabled    | yes      | Use `Azure Accelerated Networking` feature for Linux agents. This property should be always set to `"false"`. **TODO: LOG ISSUE** |

## <a name='AzureStackInstancesRegisteredwithAzuresChinacloud'></a>Azure Stack Instances Registered with Azure's China cloud

If your Azure Stack instance is located in China, then the `dependenciesLocation` property of your cluster definition should be set to `"china"`. This switch ensures that the provisioning process fetches software dependencies from reachable hosts within China's mainland.

## <a name='DisconnectedAzureStackInstances'></a>Disconnected Azure Stack Instances

By default, the AKS Engine provisioning process relies on an internet connection to download the software dependencies required to create or upgrade a cluster (Kubernetes images, etcd binaries, network plugins and so on).

If your Azure Stack instance is air-gapped or if network connectivity in your geographical location is not reliable, then the default approach will not work, take a long time or timeout due to transcient networking issues.

With these challenges in mind, you can choose to set the `distro` property of your cluster definition to `"aks-ubuntu-16.04"`. This change will instruct AKS Engine to deploy VM nodes using a base OS image call `AKS Base Image`. This custom image already contains the required software dependencies in its file system. Hence, internet connectivity is not longer required during the provisioning process.

The `AKS Base Image` gallery item has to be available in your Azure Stack's Marketplace before it could be used by AKS Engine. Follow this [guide](https://docs.microsoft.com/en-us/azure-stack/operator/azure-stack-download-azure-marketplace-item) for a general explanation about how to download gallery items from Azure. 

Each AKS Engine release is validated and tied to a specific version of the AKS Base Image. Therefore, you need to take note of the base image version required by the AKS Engine release that you plan to use, and then download exactly that base image version. New builds of the `AKS Base Image` are frequently released to ensure that your disconnected cluster can be upgraded to the latest supported version of each component.

**TODO: COMMAND TO PRINT ASK BASE IMAGE**

**TODO: LOG ISSUE TO FAIL FAST AND PRETTY MESSAGE IF BASE IMAGE IS MISSING**

## <a name='UnsupportedAddons'></a>Unsupported Addons

AKS Engine includes a number of optional [addons](../topics/clusterdefinitions.md#addons) that can be deployed as part of the cluster provisioning process.

The list below includes the addons currently unsupported on Azure Stack:

* AAD Pod Identity
* ACI Connector
* Cluster Autoscaler
* Blobfuse Flex Volume
* SMB Flex Volume
* KeyVault Flex Volume
* Rescheduler
* NVIDIA Device Plugin
* Container Monitoring

## <a name='KnownIssues'></a>Known Issues

## <a name='NodesInternetAccess'></a>Nodes losing Internet connectivity
It has been found that after we delete last kuberenetes service configured with LoadBalancer. Nodes are losing internet connnectivity.
This will cause any new application deployment failure which requires new images to download from internet.

Workaround:
Current workaround of above issue is to keep atleast one kuberenetes service configured with LoadBalancer.

## <a name='FrequentlyAskedQuestions'></a>Frequently Asked Questions
