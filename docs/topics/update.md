# Updating Kubernetes Node Pools

## Prerequisites

All documentation in these guides assumes you have already downloaded both the Azure `az` CLI tool and the `aks-engine` binary tool. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing if you're creating a Kubernetes cluster using AKS Engine for the first time.

This guide assumes you already have a running cluster deployed using the `aks-engine` CLI. For more details on how to do that see [deploy](creating_new_clusters.md#deploy) or [generate](generate.md).

## Update

The `aks-engine update` command can update the VMSS model of a node pool according to a modified configuration of the aks-engine-generated `apimodel.json`. When used in combination with a newer version of the `aks-engine` CLI compared to the version used to build the cluster originally, node pools can be regularly refreshed so that as they scale over time, new nodes always run the latest, validated bits, using your latest, validated node configuration.

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
|--client-id|depends| The Service Principal Client ID. This is required if the auth-method is set to service_principal/client_certificate|
|--client-secret|depends| The Service Principal Client secret. This is required if the auth-method is set to service_principal|
|--certificate-path|depends| The path to the file which contains the client certificate. This is required if the auth-method is set to client_certificate|
|--node-pool|yes|Which node pool should be updated.|
|--auth-method|no|The authentication method used. Default value is `client_secret`. Other supported values are: `cli`, `client_certificate`, and `device`.|
|--language|no|Language to return error message in. Default value is "en-us").|

## Frequently Asked Questions

### Why would I use update instead of upgrade to upgrade a VMSS node pool?

The `aks-engine upgrade` command actually replaces existing nodes with new nodes, one-at-a-time. Such an approach is appropriate if you are  confident that the outcome of such an operation will be successful. We recommend to attain that confidence by staging a series of full end-to-end operations simulating the series of operations in your production environment. In other words:

1. Create a cluster with a specific configuration in a specific cloud environment + region using a specific version of `aks-engine`.
  - All of the above must exactly match the original configuration + `aks-engine` version used to create your cluster initially.
2. Do something like the above for every `aks-engine` operation performed the time when your cluster was originally created and now
3. Run `aks-engine upgrade` with your desired upgrade configuration.

Because `aks-engine upgrade` is a destructive operation, and there is no definitive "undo" or "rollback", then if #3 above fails for any reason, in order to continue experimenting in your staging environment, you will have to re-stage the entire cluster + set of operations each time, until you land upon a repeatedly working `aks-engine upgrade` scenario that you confidently apply against your production scenario.

The above is a time consuming and imperfect workflow, and so `aks-engine update` is an alternative approach that allows more flexibility. For example:

- Because `aks-engine update` is merely a VMSS model update against a single node pool and not a "whole cluster", destructive operation, the viability of an updated node pool can be tested piecemeal, without affecting existing production traffic.
- In the event that the updated VMSS model produces undesirable new nodes, you may "undo" or "roll back" this model update change to the last known good VMSS model by running an `aks-engine update` operation using an older, known-working version of AKS Engine (for example, if you've never run `aks-engine update` before, you would use the version of AKS Engine you used to deploy your cluster originally) with an API model specification that has been tested as working.
