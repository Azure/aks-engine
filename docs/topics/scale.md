# Scaling Kubernetes Clusters

## Prerequisites

All the commands in this guide require both the Azure `az` CLI tool and the `aks-engine` binary tool. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing.

This guide assumes you already have deployed a cluster using `aks-engine`. For more details on how to do that see [deploy](../tutorials/deploy.md).

## Scale

The `aks-engine scale` command can increase or decrease the number of nodes in an existing agent pool in an AKS Engine-created Kubernetes cluster. Nodes will always be added or removed from the end of the agent pool. Nodes will be cordoned and drained before deletion.

This guide will assume you have a cluster deployed and the API model originally used to deploy that cluster is stored at `_output/<dnsPrefix>/apimodel.json`. It will also assume there is a node pool named "agentpool1" in your cluster.

To scale the cluster you will run a command like:

```console
$ aks-engine scale --subscription-id <subscription_id> \
    --resource-group mycluster --location <location> \
    --client-id '<service principal client ID>' \
    --client-secret '<service principal client secret>' \
    --api-model _output/mycluster/apimodel.json --new-node-count <desired node count> \
    --node-pool agentpool1 --apiserver mycluster.<location>.cloudapp.azure.com
```

This command will re-use the `apimodel.json` file inside the output directory as input for a new ARM template deployment that will execute the scaling operation against the desired agent pool. When the scaling operation is done it will update the cluster definition in that same `apimodel.json` file to reflect the new node count and thus the updated, current cluster configuration.

### Parameters

|Parameter|Required|Description|
|---|---|---|
|--subscription-id|yes|The subscription id the cluster is deployed in.|
|--resource-group|yes|The resource group the cluster is deployed in.|
|--location|yes|The location the resource group is in.|
|--api-model|yes|Relative path to the generated API model for the cluster.|
|--client-id|depends| The Service Principal Client ID. This is required if the auth-method is set to service_princpal/client_certificate|
|--client-secret|depends| The Service Principal Client secret. This is required if the auth-method is set to service_princpal|
|--certificate-path|depends| The path to the file which contains the client certificate. This is required if the auth-method is set to client_certificate|
|--node-pool|depends|Required if there is more than one node pool. Which node pool should be scaled.|
|--new-node-count|yes|Desired number of nodes in the node pool.|
|--apiserver|when scaling down|apiserver endpoint (required to cordon and drain nodes). This should be output as part of the create template or it can be found by looking at the public ip addresses in the resource group.|
|--auth-method|no|The authentication method used. Default value is `client_secret`. Other supported values are: `cli`, `client_certificate`, and `device`.|
|--language|no|Language to return error message in. Default value is "en-us").|
