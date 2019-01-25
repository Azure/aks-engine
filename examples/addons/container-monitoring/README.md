# Container-monitoring add-on

This is the Container-monitoring add-on. This add-on requires Azure Log Analytics Workspace GUID and key in Base64 encoded form. If you don't have Azure Log Analytics workspace, please create one by following the instructions in https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-quick-create-workspace.
> Note: Workspace GUID and Key MUST be in Base-64 encoded.

Here are the instructions to obtain the Workspace Guid and Key of your Azure Log Analytics workspace

- In the Azure portal, click All services. In the list of resources, type Log Analytics. As you begin typing, the list filters based on your input. Select Log Analytics Workspaces.
- In your list of Log Analytics workspaces, select the workspace you intend on configuring the agent to report to.
- Select Advanced settings.
- Select Connected Sources, and then select Linux Servers.
- Copy and paste into your favorite editor, the Workspace ID and Primary Key.

Update the json file with your Base-64 encoded Workspace Guid and Key  along with other parameter values, and use this update json file to automatically enable container monitoring in your new Kubernetes cluster. 

Following additional steps required to view the health, metrics and logs of your AKS Engine Kubernetes cluster(s)

1. [Add the 'AzureMonitor-Containers' Solution to your Log Analytics workspace.](http://aka.ms/coinhelmdoc)

2. [For AKS-engine K8S cluster, add clusterName and Log Analytics workspace tag to cluster resources, to be able to use Azure Container monitoring User experience (aka.ms/azmon-containers)](http://aka.ms/coin-acs-tag-doc)

> Note: At the moment, if the AKS engine version is v0.29.1, then default clusterName is my_acs_cluster_name in the container monitoring add-on. 
In post AKS-engine versions v0.29.1, default cluster name is aks-engine-cluster. This name needs to be tagged as clusterName value in step #2 above.

Navigate to [azmon-containers](https://aka.ms/azmon-containers) to view the health, metrics and logs of your Kubernetes cluster(s).

For more details on how to use the product, see [Azure Monitor for containers](https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-overview)

![Image of Azure Monitor for containers](../../static/img/azure_monitor_aks_engine.png)

If you have any questions regarding the Container Monitoring add-on, please reach us out thru [this](mailto:askcoin@microsoft.com) email.

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "addons": [
          {
            "name": "container-monitoring",
            "enabled": true,
            "config": {
              "workspaceGuid": "<Azure Log Analytics Workspace Guid in Base-64 encoded>",
              "workspaceKey": "<Azure Log Analytics Workspace Key in Base-64 encoded>"
            }
          }
        ]
      }
    },
    "masterProfile": {
      "count": 1,
      "dnsPrefix": "",
      "vmSize": "Standard_DS2_v2"
    },
    "agentPoolProfiles": [
      {
        "name": "agentpool",
        "count": 3,
        "vmSize": "Standard_DS2_v2",
        "availabilityProfile": "VirtualMachineScaleSets"
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

You can validate that the add-on is running as expected with the following commands:

You should see two sets of omsagent pods 1). One omsagent replicat set in one of the node 2). Many omsagent demonset pods as there are agent nodes .
All these pods should be in 'Running' state after executing:

```bash
kubectl get pods -n kube-system
```

## Supported Matrix

 Refer to [azuremonitor-containers-aks-engine](https://github.com/Microsoft/OMS-docker/blob/aks-engine/README.md) for the supported matrix, troubleshooting and supportability etc.

