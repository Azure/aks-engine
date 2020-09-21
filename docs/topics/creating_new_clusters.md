# Creating Kubernetes Clusters

## Deploy

The `aks-engine deploy` command will create a new Kubernetes cluster from scratch into a pre-existing Azure resource group. You define an API model (cluster definition) as a JSON file, and then pass in a reference to it, as well as appropriate Azure credentials, to a command statement like this:

```sh
$ aks-engine deploy --subscription-id $SUBSCRIPTION_ID \
    --dns-prefix $CLUSTER_NAME \
    --resource-group $RESOURCE_GROUP \
    --location $LOCATION \
    --api-model examples/kubernetes.json \
    --client-id $SERVICE_PRINCIPAL_ID \
    --client-secret $SERVICE_PRINCIPAL_PASSWORD
```

`aks-engine deploy` is a long-running operation that creates Azure resources (e.g., Virtual Machine and/or Virtual Machine Scale Set (VMSS), Disk, Network Interface, Network Security Group, Public IP Address, Virtual Network, Load Balancer, and others) that will underly a Kubernetes cluster. All deployed VMs will be configured to run Kubernetes bootstrap scripts appropriate for the desired cluster configuration. The outcome of a successful `aks-engine deploy` operation is a fully operational Kubernetes cluster, ready for use immediately.

A more detailed walk-through of `aks-engine deploy` is in the [quickstart guide](../tutorials/quickstart.md#deploy)

### Parameters

|Parameter|Required|Description|
|-----------------|---|---|
|--api-model|yes|Relative path to the API model (cluster definition) that declares the desired cluster configuration.|
|--dns-prefix|no, if present in API model|Unique name for the cluster.|
|--auto-suffix|no|Automatically append a compressed timestamp to the dnsPrefix to ensure cluster name uniqueness.|
|--azure-env|no|The target Azure cloud (default "AzurePublicCloud") to deploy to.|
|--subscription-id|yes|The subscription id the cluster is deployed in.|
|--resource-group|yes|The resource group the cluster is deployed in.|
|--location|yes|The location to deploy to.|
|--force-overwrite|no|Automatically overwrite any existing files in the output directory (default is false).|
|--output-directory|no|Output directory (derived from FQDN if absent) to persist cluster configuration artifacts to.|
|--set|no|Set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2).|
|--ca-certificate-path|no|Path to the CA certificate to use for Kubernetes PKI assets.|
|--ca-private-key-path|no|Path to the CA private key to use for Kubernetes PKI assets.|
|--client-id|depends| The Service Principal Client ID. This is required if the auth-method is set to service_princpal/client_certificate|
|--client-secret|depends| The Service Principal Client secret. This is required if the auth-method is set to service_princpal|
|--certificate-path|depends| The path to the file which contains the client certificate. This is required if the auth-method is set to client_certificate|
|--identity-system|no|Identity system (default is azure_ad)|
|--auth-method|no|The authentication method used. Default value is `client_secret`. Other supported values are: `cli`, `client_certificate`, and `device`.|
|--private-key-path|no|Path to private key (used with --auth-method=client_certificate).|
|--language|no|Language to return error message in. Default value is "en-us").|

## Frequently Asked Questions

### Can I re-run `aks-engine deploy` on an existing cluster to update the cluster configuration?

No. See [addpool](addpool.md), [update](update.md), [scale](scale.md), and [upgrade](upgrade.md) for documentation describing how to continue to use AKS Engine to maintain your cluster configuration over time.
