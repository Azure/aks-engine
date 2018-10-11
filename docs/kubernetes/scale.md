# Scaling a Kubernetes Cluster

## Install Pre-requisites

All the commands in this guide require both the Azure CLI and `aks-engine`. Follow the [installation instructions to download aks-engine before continuing](../aksengine.md#install-aks-engine) or [compile from source](../aksengine.md#build-from-source).

For installation instructions see [the Azure CLI GitHub repository](https://github.com/Azure/azure-cli#installation) for the latest release.

This guide assumes you already have deployed a cluster using AKS Engine. For more details on how to do that see [deploy](./deploy.md).

## Scale
After a cluster has been deployed using AKS Engine the cluster can be interacted further by using the scale command. The scale command can add more nodes to an existing node pool or remove them. Nodes will always be added or removed from the end of the agent pool. Nodes will be cordoned and drained before deletion.

This guide will assume you have a cluster deployed and the output for the deployed cluster is stored at _output/mycluster. It will also assume there is a node pool named "agentpool1" in your cluster. AKS Engine will default to storing the output at ./_output/dns-prefix from where the aks-engine command was ran.

To scale the cluster you will run a command like:

```
$ aks-engine scale --subscription-id 51ac25de-afdg-9201-d923-8d8e8e8e8e8e \
    --resource-group mycluster  --location westus2 \
    --deployment-dir _output/mycluster --new-node-count 5 \
    --node-pool agentpool1 --master-FQDN mycluster.westus2.cloudapp.azure.com
```

This command will look the the deployment directory to find info about the cluster currently deployed. Then it will generate and deploy a template deployment to update the cluster and add the new nodes. When it is done it will update the cluster definition in the deployment directory's apimodel.json to reflect the new node count.

### Parameters
|Parameter|Required|Description|
|---|---|---|
|subscription-id|yes|The subscription id the cluster is deployed in.|
|resource-group|yes|The resource group the cluster is deployed in.|
|location|yes|The location the resource group is in.|
|deployment-dir|yes|Relative path to the folder location for the output from the aks-engine deploy/generate command.|
|node-pool|depends|Required if there is more than one node pool. Which node pool should be scaled.|
|new-node-count|yes|Desired number of nodes in the node pool.|
|master-FQDN|depends|When scaling down a kuberentes cluster this is required. The master FDQN so that the nodes can be cordoned and drained before removal. This should be output as part of the create template or it can be found by looking at the public ip addresses in the resource group.|