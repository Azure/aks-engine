# Kubernetes on Cosmos
This example deployment allows you to use AKS Engine to bootstrap a Kubernetes cluster on Azure that uses [Azure Cosmos](https://azure.microsoft.com/services/cosmos-db/) instead of locally installed and configured etcd.

## What is the Azure Cosmos etcd API?
[Azure Cosmos](https://azure.microsoft.com/services/cosmos-db/) is a globally distributed, multi-model database service. It offers turnkey global distribution, guarantees single-digit millisecond latencies at the 99th percentile, elastic scaling of throughput and storage, and comprehensive SLA’s.

The Azure Cosmos etcd API allows you to use Azure Cosmos as the backing store for Kubernetes. Azure Cosmos implements the wire level protocol of etcd, allowing the master nodes’ API servers to use Azure Cosmos just like it would for a locally installed etcd. 

## Prerequisites 
1.	Install the [Azure CLI](https://docs.microsoft.com/cli/azure/install-azure-cli?view=azure-cli-latest)
2.	Install [AKS Engine](https://github.com/Azure/aks-engine/blob/master/docs/tutorials/quickstart.md)
3.	Have a whitelisted subscription for the Azure Cosmos etcd API. Follow instructions at https://aka.ms/cosmosetcdapi-signup to enroll in the private preview. 

## Steps
Follow the instructions [here](https://github.com/Azure/aks-engine/blob/master/docs/tutorials/quickstart.md#deploy-your-first-cluster) to deploy your first Kubernetes cluster. 

Here is an example `aks-engine` command: 

```console
$ aks-engine deploy --subscription-id <SubscriptionId> \
    --client-id <servicePrincipalClientId> \
    --client-secret <ServicePrincipalSecret> \
    --dns-prefix <YourDNSPrefix> \
    --location centralus \
    --api-model <apiModel>.json
```

In the API model JSON (the cluster definition file), specify ```"cosmosEtcd" : true``` in the ```masterProfile``` properties to use Azure Cosmos etcd API instead of regular etcd. 

A sample cluster definition file using Azure Cosmos is available [here](https://github.com/Azure/aks-engine/blob/master/examples/cosmos-etcd/kubernetes-3-masters-cosmos.json).

## Verifying the deployment
After the deployment is successfully completed, you will see a new Azure Cosmos account created in your resource group. This account name will match your specified DNS prefix appended with k8s. 

## FAQ

**What is provisioned in my Azure Cosmos account?**

Your Azure Cosmos account will be provisioned with a database (EtcdDB) and a container (EtcdData). The container will store all the etcd related data.

**What is the difference between using regular etcd vs. the Azure Cosmos etcd API as the backing store for Kubernetes?**

With Azure Cosmos etcd API, benefits include:
* No need to manually configure and manage etcd
* High availability of etcd, guaranteed by Cosmos (99.99% in single region, 99.999% in 2+ regions). 
* Elastic scalability of etcd
* Secure by default & enterprise ready
* Industry-leading, comprehensive SLAs

**Which Azure regions are Azure Cosmos etcd API available in?**

In the preview, Central US is the currently supported region. More regions will be available in the future. 

**When will this be available for Azure Kubernetes Service (AKS)?**
Integration with AKS is on the roadmap. Please upvote on [UserVoice](https://feedback.azure.com/forums/914020-azure-kubernetes-service-aks) and stay tuned.  

## Getting Support
To get support during the preview, you can file GitHub issues on this repo. Please mention @Azure/cosmosetcd when filing the issue. You can also email "askcosmosetcdapi@microsoft.com".
