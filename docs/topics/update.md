# Updating Kubernetes Node Pools

## Prerequisites

All the commands in this guide require both the Azure `az` CLI tool and the `aks-engine` binary tool. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing if you're creating a Kubernetes cluster using AKS Engine for the first time.

This guide assumes you already have a running cluster deployed using the `aks-engine` CLI. For more details on how to do that see [deploy](deploy.md) or [generate](generate.md).

## Update

The `aks-engine update` command can update the VMSS model of a node pool according to a modified configuration of the aks-engine-generated `apimodel.json`. When used in combination with a newer version of the `aks-engine` CLI compared to the version used to build the cluster originally, node pools can be regularly refreshed so that as they scale over time, new nodes always run the latest, validated bits.

This command can *only* be used with VMSS-backed node pools (the default AKS Engine node pool type is VMSS).

The example below will assume you have a cluster deployed, and that the API model originally used to deploy that cluster is stored at `_output/<dnsPrefix>/apimodel.json`. It will also assume that there is a node pool named "agentpool1" in your cluster.

To update the cluster you will run a command like:

```sh
$ aks-engine update --subscription-id <subscription_id> \
    --resource-group mycluster --location <location> \
    --client-id '<service principal client ID>' \
    --client-secret '<service principal client secret>' \
    --api-model _output/mycluster/apimodel.json \
    --node-pool agentpool1
```

The above operation will complete rather quickly, as it is only updating the VMSS model; it is not actually modifying any existing VMSS instances.

### Parameters

|Parameter|Required|Description|
|-----------------|---|---|
|--subscription-id|yes|The subscription id the cluster is deployed in.|
|--resource-group|yes|The resource group the cluster is deployed in.|
|--location|yes|The location the resource group is in.|
|--api-model|yes|Relative path to the generated API model for the cluster.|
|--client-id|depends| The Service Principal Client ID. This is required if the auth-method is set to service_princpal/client_certificate|
|--client-secret|depends| The Service Principal Client secret. This is required if the auth-method is set to service_princpal|
|--certificate-path|depends| The path to the file which contains the client certificate. This is required if the auth-method is set to client_certificate|
|--node-pool|yes|RWhich node pool should be updated.|
|--auth-method|no|The authentication method used. Default value is `client_secret`. Other supported values are: `cli`, `client_certificate`, and `device`.|
|--language|no|Language to return error message in. Default value is "en-us").|

## Frequently Asked Questions

### Why would I use update instead of upgrade to upgrade a VMSS node pool?
