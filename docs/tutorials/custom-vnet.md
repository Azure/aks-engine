# Using a custom virtual network with AKS Engine

In this tutorial you are going to learn how to use [AKS Engine](https://github.com/Azure/aks-engine) to deploy a brand new cluster into an existing or pre-created virtual network.
By doing this, you will be able to control the properties of the virtual network or integrate a new cluster into your existing infrastructure.

_The Kubernetes cluster should be deployed in the same resource group as the virtual network and the service principal you use for the cluster needs permissions on the VNET resource's group too._

## Prerequisites

You can run this walkthrough on OS X, Windows, or Linux.

- You need an Azure subscription. If you don't have one, you can [sign up for an account](https://azure.microsoft.com/).
- Install the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest).
- Install [AKS Engine](quickstart.md#install-aks-engine)

## Create the virtual network

_You need a virtual network before creating the new cluster. If you already have one, you can skip this step._

For this example, we deployed a virtual network that contains two subnets:

- 10.100.0.0/24
- 10.200.0.0/24

The first one will be used for the master nodes and the second one for the agent nodes.

The Azure Resource Manager template used to deploy this virtual network is:

```json
{
  "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
  "contentVersion": "1.0.0.0",
  "parameters": {},
  "variables": {},
  "resources": [
    {
      "apiVersion": "2018-06-01",
      "location": "[resourceGroup().location]",
      "name": "ExampleCustomVNET",
      "properties": {
        "addressSpace": {
          "addressPrefixes": ["10.100.0.0/24", "10.200.0.0/24"]
        },
        "subnets": [
          {
            "name": "ExampleMasterSubnet",
            "properties": {
              "addressPrefix": "10.100.0.0/24"
            }
          },
          {
            "name": "ExampleAgentSubnet",
            "properties": {
              "addressPrefix": "10.200.0.0/24"
            }
          }
        ]
      },
      "type": "Microsoft.Network/virtualNetworks"
    }
  ]
}
```

And you can deploy it using the Azure CLI 2.0. First, you need to create a new resource group:

```bash
az group create -n aks-custom-vnet -l "westeurope"
```

Then you can deploy the virtual network using the JSON description above and the following command:

```bash
az deployment group create -g aks-custom-vnet --name "CustomVNet" --template-file azuredeploy.vnet.json
```

Alternatively, you can use Azure CLI to create the vnet and the subnet directly:

```bash
az network vnet create --resource-group aks-custom-vnet --name CustomVNET --address-prefixes 10.100.0.0/24 10.200.0.0/24 --subnet-name ExampleMasterSubnet --subnet-prefixes 10.100.0.0/24

az network vnet subnet create --resource-group aks-custom-vnet --vnet-name CustomVNET -n ExampleAgentSubnet --address-prefixes 10.200.0.0/24
```

Once the deployment is completed you should see the virtual network in the resource group.

## Create the template for AKS Engine

AKS Engine uses a JSON template in input and generates the ARM template and ARM parameters files in output.

Depending on the orchestrator you want to deploy, the number of agent pools, the machine size you want (etc.) this input template could differ from the one we are going to detail here.

There are a lot of examples available on the [AKS Engine GitHub](https://github.com/Azure/aks-engine/tree/master/examples) and you can find [one dedicated for virtual network](https://github.com/Azure/aks-engine/blob/master/examples/vnet/README.md).

In this case, we are going to use the following template (this creates a cluster with the Azure CNI network plugin):

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes"
    },
    "masterProfile": {
      "count": 1,
      "dnsPrefix": "",
      "vmSize": "Standard_D2_v2",
      "vnetSubnetId": "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/CustomVNET/subnets/ExampleMasterSubnet",
      "firstConsecutiveStaticIP": "10.100.0.5"
    },
    "agentPoolProfiles": [
      {
        "name": "agents",
        "count": 3,
        "vmSize": "Standard_D2_v2",
        "vnetSubnetId": "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME/providers/Microsoft.Network/virtualNetworks/CustomVNET/subnets/ExampleAgentSubnet"
      }
    ],
    "linuxProfile": {
      "adminUsername": "azureuser",
      "ssh": {
        "publicKeys": [
          {
            "keyData": ""
          }
        ]
      }
    },
    "servicePrincipalProfile": {
      "clientId": "",
      "secret": ""
    }
  }
}
```

As you can see, for all node pools definition (master or agents) you can use the **vnetSubnetId** and **firstConsecutiveStaticIP** properties to defines the virtual network where you want to deploy the cluster and the first IP address that should be used by the first machine in the pool.

_Note: Make sure the the vnetSubnetId matches with your subnet, by giving your **SUBSCRIPTION_ID**, **RESOURCE_GROUP_NAME**, virtual network and subnet names. You also need to fill DNS prefix for all the public pools you want to create, give an SSH keys..._

## Deploy the AKS Engine template

Once you are ready with the cluster definition file, you can either use AKS engine to deploy the cluster on Azure or you can use AKS Engine to generate an ARM template to deploy using Azure CLI.

### Deploy using AKS Engine

 Follow the [instructions on how to deploy](deploy.md#deploy) using the cluster definition (API model) file you prepared.

### Generate the cluster Azure Resource Manager template

Once your are ready with the cluster definition file, you can use AKS Engine to generate the ARM template that will be used to deploy the cluster on Azure:

```bash
aks-engine generate --api-model clusterdefinition.json
```

This command will output the following files in `_output/test`:

```console
INFO[0000] Generating assets into _output/test...
DEBU[0011] pki: PKI asset creation took 7.9016753s
DEBU[0011] output: wrote _output/test/apimodel.json
DEBU[0011] output: wrote _output/test/azuredeploy.json
DEBU[0011] output: wrote _output/test/azuredeploy.parameters.json
DEBU[0011] output: wrote _output/test/kubeconfig/kubeconfig.southcentralus.json
DEBU[0011] output: wrote _output/test/ca.key
DEBU[0011] output: wrote _output/test/ca.crt
DEBU[0011] output: wrote _output/test/apiserver.key
DEBU[0011] output: wrote _output/test/apiserver.crt
DEBU[0011] output: wrote _output/test/client.key
DEBU[0011] output: wrote _output/test/client.crt
DEBU[0011] output: wrote _output/test/kubectlClient.key
DEBU[0011] output: wrote _output/test/kubectlClient.crt
DEBU[0011] output: wrote _output/test/etcdserver.key
DEBU[0011] output: wrote _output/test/etcdserver.crt
DEBU[0011] output: wrote _output/test/etcdclient.key
DEBU[0011] output: wrote _output/test/etcdclient.crt
DEBU[0011] output: wrote _output/test/etcdpeer0.key
DEBU[0011] output: wrote _output/test/etcdpeer0.crt
aksengine took 37.1384ms
```

- apimodel.json: this is the cluster definition file you gave to AKS Engine
- azuredeploy.json: this is the Azure Resource Manager JSON template that you are going to use to deploy the cluster
- azuredeploy.parameters.json: this is the parameters file that you are going to use to deploy the cluster

#### Deploy the Azure Container Service cluster

Now that you have generated the ARM templates and its parameters file using AKS Engine, you can use Azure CLI 2.0 to start the deployment of the cluster:

```bash
az deployment group create -g aks-custom-vnet --name "ClusterDeployment" --template-file azuredeploy.json --parameters "@azuredeploy.parameters.json"
```

Depending on the number of agents you have asked for the deployment can take a while.

## Post-Deployment: Attach Cluster Route Table to VNET

_NOTE: This section is applicable only to Kubernetes clusters that use Kubenet. If AzureCNI is enabled in your cluster, you may disregard._

For Kubernetes clusters, we need to update the VNET to attach to the route table created by the above `az deployment group create` command. An example in bash form if the VNET is in the same ResourceGroup as the Kubernetes Cluster:

```
#!/bin/bash
rt=$(az network route-table list -g aks-custom-vnet -o json | jq -r '.[].id')
az network vnet subnet update -n KubernetesSubnet \
-g aks-custom-vnet \
--vnet-name KubernetesCustomVNET \
--route-table $rt
```

... where `KubernetesSubnet` is the name of the vnet subnet, and `KubernetesCustomVNET` is the name of the custom VNET itself.

An example in bash form if the VNET is in a separate ResourceGroup:

```bash
#!/bin/bash
rt=$(az network route-table list -g RESOURCE_GROUP_NAME_KUBE -o json | jq -r '.[].id')
az network vnet subnet update \
-g RESOURCE_GROUP_NAME_VNET \
--route-table $rt \
--ids "/subscriptions/SUBSCRIPTION_ID/resourceGroups/RESOURCE_GROUP_NAME_VNET/providers/Microsoft.Network/VirtualNetworks/KUBERNETES_CUSTOM_VNET/subnets/KUBERNETES_SUBNET"
```

... where `RESOURCE_GROUP_NAME_KUBE` is the name of the Resource Group that contains the AKS Engine-created Kubernetes cluster, `SUBSCRIPTION_ID` is the id of the Azure subscription that both the VNET & Cluster are in, `RESOURCE_GROUP_NAME_VNET` is the name of the Resource Group that the VNET is in, `KUBERNETES_SUBNET` is the name of the vnet subnet, and `KUBERNETES_CUSTOM_VNET` is the name of the custom VNET itself.

## Connect to your new cluster

Once the deployment is completed, you can follow [this documentation](https://docs.microsoft.com/en-us/azure/container-service/container-service-connect) to connect to your new Azure Container Service cluster.
