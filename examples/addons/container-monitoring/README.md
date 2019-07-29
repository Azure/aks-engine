# Container-monitoring addon

This is sample API definition with Container-monitoring addon.

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
            "enabled": true
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
> Note: If the AKS-Engine version is v0.38.5 or higher, the clusterName will be the dnsPrefix of the cluster. If the AKS Engine version is v0.29.1 or ACS-Engine Kubernetes cluster, then default clusterName is my_acs_cluster_name in the container monitoring addon. In  AKS Engine version is less than v0.38.5, the default cluster name is aks-engine-cluster.

Navigate to [azmon-containers](https://aka.ms/azmon-containers) to view the health, metrics and logs of your Kubernetes cluster(s).


You can validate that the addon is running as expected with the following commands:

You should see two sets of omsagent pods 1). One omsagent replicat set in one of the node 2). Many omsagent demonset pods as there are agent nodes .
All these pods should be in 'Running' state after executing:

```bash
kubectl get pods -n kube-system
```

## UX

For more details on how to use the product, see [Azure Monitor for containers](https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-overview)

![Image of Azure Monitor for containers](../../../docs/static/img/azure_monitor_aks_engine.png)

## Supported Matrix

 Refer to [azuremonitor-containers-aks-engine](https://github.com/Microsoft/OMS-docker/blob/aks-engine/README.md) for the supported matrix, troubleshooting and supportability etc.

## Disable Monitoring

After you enable monitoring of your AKS Engine cluster, you can stop monitoring the cluster if you decide you no longer want to monitor it.

- If you have onboarded the monitoring using the HELM chart, then you can disable monitoring by uninstalling the chart. Refer Uninstalling the Chart section in [azuremonitor-containers](https://github.com/helm/charts/tree/master/incubator/azuremonitor-containers)

- If you have onboarded using the Container Monitoring addon, then you can stop monitoring either

      - kubectl delete -f <kubernetesmasteraddons-omsagent-daemonset.yaml>

        or

      - deleting omsagent daemonset and replicaset using Kubectl tool

## Contact

If you have any questions or feedback regarding the container monitoring addon, please reach us out through [this](mailto:askcoin@microsoft.com) email.

## References

- [Log Analytics](https://docs.microsoft.com/en-us/azure/azure-monitor/log-query/log-query-overview)
- [Azure Resource Manager](https://docs.microsoft.com/en-us/azure/azure-monitor/platform/template-workspace-configuration)
- [Azure portal](https://docs.microsoft.com/en-us/azure/azure-monitor/learn/quick-create-workspace)
- [Manage workspaces](https://docs.microsoft.com/en-us/azure/azure-monitor/platform/manage-access)
- [Link to Container Health](https://aka.ms/ci-dogfood)
