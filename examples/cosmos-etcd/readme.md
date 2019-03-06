# Kubernetes on Cosmos

The Azure Kubernetes Engine (aks-engine) generates [ARM](https://docs.microsoft.com/azure/azure-resource-manager/resource-group-overview) (Azure Resource Manager) templates for Kubernetes clusters on Microsoft Azure. The input to aks-engine is a cluster definition file which describes the desired cluster, including orchestrator, features, and agents. 

This example deployment allows you to use aks-engine to bootstrap a Kubernetes cluster on Azure that uses [Azure Cosmos DB](https://azure.microsoft.com/services/cosmos-db/) instead of locally installed and configured etcd.

## What is the Azure Cosmos DB etcd API?
[Azure Cosmos DB](https://azure.microsoft.com/services/cosmos-db/) is a globally distributed, multi-model database service. It offers turnkey global distribution, guarantees single-digit millisecond latencies at the 99th percentile, elastic scaling of throughput and storage, and comprehensive SLA’s.

The Azure Cosmos DB etcd API allows you to use Azure Cosmos DB as the backing store for Kubernetes. Azure Cosmos DB implements the wire level protocol of etcd, allowing master nodes’ API servers to communicate with and store data Azure Cosmos DB as if it were locally installed etcd. 

## Prerequisites 
1.	Install the [Azure CLI](https://docs.microsoft.com/cli/azure/install-azure-cli?view=azure-cli-latest)
2.	Install [AKS-Engine](https://github.com/Azure/aks-engine/blob/master/docs/tutorials/quickstart.md)
3.	Have a whitelisted subscription for the Azure Cosmos DB etcd API. Follow instructions at https://aka.ms/cosmosetcdapi-signup to enroll in the private preview. 

## Steps
Follow the instructions [here](https://github.com/Azure/aks-engine/blob/master/docs/tutorials/quickstart.md#deploy-your-first-cluster) to deploy your first Kubernetes cluster. 

Follow the instructions here to deploy your first Kubernetes cluster. 
Here is an example aks-engine command: 

```console
$ aks-engine deploy --subscription-id <SubscriptionId> \
    --client-id <servicePrincipalClientId> \
    --client-secret <ServicePrincipalSecret> \
    --dns-prefix <YourDNSPrefix> \
    --location centralus \
    --api-model <apiModel>.json
```

In the API model JSON (the cluster definition file), specify ```"cosmosEtcd" : true``` in the ```masterProfile``` properties to use Azure Cosmos DB etcd API instead of regular etcd. 

A sample cluster definition file using Azure Cosmos DB is available [here](https://github.com/Azure/aks-engine/blob/master/examples/cosmos-etcd/kubernetes-3-masters-cosmos.json).

## Verifying the deployment
After deployment is successfully completed, you will see a new Azure Cosmos DB account created in your resource group. The account name will match your specified DNS prefix appended with k8s. 

![Image of Azure Cosmos DB etcd API account](../static/img/cosmos-etcd-account.png)


## FAQ

**What is provisioned in my Azure Cosmos DB account?**

Your Azure Cosmos DB account will be provisioned with adatabase (EtcdDB) and Collection (EtcdData). The collection will store all the etcd related data. By default, it is provisioned with throughput 10000 RU/s. Read more about [Request Units and provisioned throughput](https://docs.microsoft.com/azure/cosmos-db/request-units).

**What is the difference between using regular etcd vs. the Azure Cosmos DB etcd API as the backing store for Kubernetes?**

With Azure Cosmos DB etcd API, benefits include:
* No need to manually configure and manage etcd
* High availability guarantees of Cosmos (99.99% in single region, 99.999% in 2+ regions). 

**Which Azure regions are Azure Cosmos DB etcd API available in?**

In the preview, Central US is the currently supported region. More regions will be available in the future. 

**When will this be available for Azure Kubernetes Service (AKS)?**
Integration with AKS is on the roadmap. Please upvote on [UserVoice](https://feedback.azure.com/forums/914020-azure-kubernetes-service-aks) and stay tuned.  

## Getting Support
To get support during the preview you can file GitHub issues on this repo. Please mention cosmosetcd team when filing the issue. 
You can also email askcosmosetcdapi@microsoft.com 
