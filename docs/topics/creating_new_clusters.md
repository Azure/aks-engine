# Creating Kubernetes Clusters

## Deploy

The `aks-engine deploy` command will create a new Kubernetes cluster from scratch into a pre-existing Azure resource group. You define an API model (cluster definition) as a JSON file, and then pass in a reference to it, as well as appropriate Azure credentials, to a command statement like this:

```sh
$ aks-engine deploy --subscription-id $SUBSCRIPTION_ID \
    --dns-prefix $CLUSTER_NAME \
    --resource-group $RESOURCE_GROUP \
    --location $LOCATION \
    --api-model examples/kubernetes.json
```

`aks-engine deploy` is a long-running operation that creates Azure resources (e.g., Virtual Machine and/or Virtual Machine Scale Set [VMSS], Disk, Network Interface, Network Security Group, Public IP Address, Virtual Network, Load Balancer, and others) that will underly a Kubernetes cluster. All deployed VMs will be configured to run Kubernetes bootstrap scripts appropriate for the desired cluster configuration. The outcome of a successful `aks-engine deploy` operation is a fully operational Kubernetes cluster, ready for use immediately.

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
|--client-id|depends| The Service Principal Client ID. This is required if the auth-method is set to client_secret or client_certificate|
|--client-secret|depends| The Service Principal Client secret. This is required if the auth-method is set to client_secret|
|--certificate-path|depends| The path to the file which contains the client certificate. This is required if the auth-method is set to client_certificate|
|--identity-system|no|Identity system (default is azure_ad)|
|--auth-method|no|The authentication method used. Default value is `client_secret`. Other supported values are: `cli`, `client_certificate`, and `device`.|
|--private-key-path|no|Path to private key (used with --auth-method=client_certificate).|
|--language|no|Language to return error message in. Default value is "en-us").|

## Generate

The `aks-engine generate` command will generate artifacts that you can use to implement your own cluster create workflows. Like `aks-engine deploy`, you define an API model (cluster definition) as a JSON file, and then pass in a reference to it, as well as appropriate Azure credentials, to a command statement like this:

```sh
$ aks-engine generate --api-model ./my-cluster-definition.json  \
    --output-directory ./cluster_artifacts
```

The above command assumes that the API model at the relative filepath `./my-cluster-definition.json` contains a minimally populated cluster definition. At a minimum is needed:

1. In order to grant the required service privileges to Kubernetes runtime processes, you need either:

Service Principal credentials in order to grant Azure privileges to the relevent Kubernetes runtime processes:
```json
{
...
  "properties": {
...
    "servicePrincipalProfile": {
      "clientId": "<service principal ID>",
      "secret": "<service principal password>"
    }
...
}
```

Or, system-assigned identity enabled:
```json
{
...
  "properties": {
...
    "orchestratorProfile": {
      "kubernetesConfig": {
        "useManagedIdentity": true
        ...
      }
      ...
    }
...
}
```

2. To uniquely identify the cluster, you need a cluster name:
```json
{
...
  "properties": {
...
    "masterProfile": {
      "dnsPrefix": "<name of cluster>"
      ...
    }
...
}
```

3. To enable interactive login to node VMs via ssh key exchange, you need to provide a public key:
```json
{
...
  "properties": {
...
    "linuxProfile": {
      "ssh": {
        "publicKeys": [
          {
            "keyData": "<public key data>"
          }
        ]
      }
      ...
    }
...
}
```

### Parameters

|Parameter|Required|Description|
|-----------------|---|---|
|--api-model|yes|Relative path to the API model (cluster definition) that declares the desired cluster configuration.|
|--output-directory|no|Output directory (derived from FQDN if absent) to persist cluster configuration artifacts to.|
|--set|no|Set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2).|
|--ca-certificate-path|no|Path to the CA certificate to use for Kubernetes PKI assets.|
|--ca-private-key-path|no|Path to the CA private key to use for Kubernetes PKI assets.|
|--client-id|depends| The Service Principal Client ID. This is required if the auth-method is set to service_principal/client_certificate|
|--client-secret|depends| The Service Principal Client secret. This is required if the auth-method is set to service_principal|
|--parameters-only|no|Only output parameters files.|
|--no-pretty-print|no|Skip pretty printing the output.|

As mentioned above, `aks-engine generate` expects all cluster definition data to be present in the API model JSON file. You may actually inject data into the API model at runtime by invoking the command and including that data in the `--set` argument interface. For example, this command will produce artifacts that can be used to deploy a fully functional Kubernetes cluster based on the AKS Engine defaults (the `examples/kubernetes.json` file will build a "default" single master, 2 node cluster):

```sh
$ bin/aks-engine generate --api-model ./examples/kubernetes.json \
  --output-directory ./cluster_artifacts \
  --set masterProfile.dnsPrefix=my-cluster,orchestratorProfile.kubernetesConfig.useManagedIdentity=true,linuxProfile.ssh.publicKeys[0].keyData=$(cat ~/.ssh/id_rsa.pub)
INFO[0000] new API model file has been generated during merge: /var/folders/jq/t_y8l4556rv__mzvjhkd61n00000gp/T/mergedApiModel831700038
WARN[0000] No "location" value was specified, AKS Engine will generate an ARM template configuration valid for regions in public cloud only
INFO[0000] Generating assets into ./cluster_artifacts...
WARN[0000] containerd will be upgraded to version 1.3.7
```

## Frequently Asked Questions

### Why would I run `aks-engine generate` vs `aks-engine deploy`?

Depending on any customization you want to do either (1) the AKS Engine-generated ARM template, or (2) the particular way that the ARM template is deployed to Azure, you may want to use `aks-engine generate` to produce an ARM template specification, and then implement your own `az deployment group create`-equivalent workflow to actually bootstrap the cluster. Especially if you plan to bootstrap multiple clusters in multiple regions from a common cluster configuration, it may make sense to re-use a single ARM template across a set of ARM deployments. `aks-engine deploy` is only able to build one cluster at a time, so especially if you're bootstrapping multiple clusters in parallel using a common config, a workflow like this is probably optimal:

1. `aks-engine generate --api-model ./common-cluster-definition.json --output-directory /path/to/re-usable-arm-template-directory`
2. For every desired cluster+location, execute in parallel:
  1. `az group create -n $RESOURCE_GROUP -l $LOCATION`; then
  2. `az deployment group create --name $RESOURCE_GROUP --resource-group $RESOURCE_GROUP --template-file /path/to/re-usable-arm-template-directory/azuredeploy.json --parameters /path/to/re-usable-arm-template-directory/azuredeploy.parameters.json`

In the above example we use name of the resource group as the name of the ARM deployment, following the guidance that only one cluster be built per resource group.

In summary, when creating single clusters, and especially when maintaining Kubernetes environments distinctly (i.e., not maintaining a fleet of clusters running a common config), relying upon `aks-engine deploy` as a full end-to-end convenience to bootstrap your clusters is appropriate. For more sophisticated cluster configuration re-use scenarios, and/or more sophisticated ARM deployment reconciliation (i.e., retry logic for certain failures), `aks-engine generate` + `az deployment group create` is the more appropriate choice.

### Can I re-run `aks-engine deploy` on an existing cluster to update the cluster configuration?

No. See [addpool](addpool.md), [update](update.md), [scale](scale.md), and [upgrade](upgrade.md) for documentation describing how to continue to use AKS Engine to maintain your cluster configuration over time.
