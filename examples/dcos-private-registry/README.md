# Private Registry Support

ACS can deploy credentials to private registries to agent nodes DC/OS clusters.

The credentials are specified in the orchestrator profile in the apimodel:
```
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "DCOS",
      "dcosConfig" : {
        "Registry" : "",
        "RegistryUser" : "",
        "RegistryPassword" : ""
      }
    },
```

The agent provisioning process will then create a tar archive containing a docker config as documented at: [Using a Private Docker Registry](https://docs.mesosphere.com/1.9/deploying-services/private-docker-registry/)

## Example
Let's provision a DC/OS cluster with credentials to an [Azure Container Registry](https://azure.microsoft.com/en-us/services/container-registry/) deployed to every agent node.

- First, [provision an Azure Container Registry](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-managed-get-started-portal).

- Enable Admin Access and note the registry credentials
<img src="../../docs/images/acrblade.png" alt="ACR Blade with Admin Access enabled" style="width: 50%; height: 50%;"/>

- Clone [aks-engine](http://github.com/azure/aks-engine) and [start the container with the dev environment](https://github.com/Azure/aks-engine/blob/master/docs/aksengine.md).

- Edit the API model to include the credentials
```
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "DCOS",
      "registry" : "xtophregistry.azurecr.io",
      "registryUser" : "xtophregistry",
      "registryPassword" : "aN//=+l==Z+/A=3hXhA+mSX=rXwB/UgW"
    },
```

- Run aks-engine to create ARM templates
```
./aks-engine generate examples/dcos-private-registry/dcos.json
```

- Deploy the cluster
```
az group create -l eastus -n cluster-rg
az group deployment create -g cluster-rg --template-file _output/dcoscluster/azuredeploy.json --parameters @_output/dcoscluster/azuredeploy.parameters.json
```

- Create a Service to deploy a container from the ACR
<img src="../../docs/images/dcos-create-service-from-reg.png" alt="Service Creation from Registry" style="width: 50%; height: 50%;"/>

- Add the credential path on the agent using the JSON editor
<img src="../../docs/images/dcos-create-service-json.png" alt="JSON editor with credential path" style="width: 50%; height: 50%;"/>

- See the Service running
<img src="../../docs/images/dcos-running-service-from-reg.png" alt="Running Service" style="width: 50%; height: 50%;"/>

- Check the credential deployment
<img src="../../docs/images/dcos-running-service-from-reg-files.png" alt="Running Service" style="width: 50%; height: 50%;"/>

## Limitations
- The API model currenlty only supports credentials to a single registry.
- Not tested with Kubernetes clusters
- Credentials have to be updated on each node
