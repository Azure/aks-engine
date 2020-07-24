# Azure Arc enabled Kubernetes

You can attach and configure Kubernetes clusters by using [Azure Arc enabled Kubernetes](https://docs.microsoft.com/azure/azure-arc/kubernetes/overview).
When a Kubernetes cluster is attached to Azure Arc, it will appear in the Azure portal. It will have an Azure Resource Manager ID and a managed identity.
Clusters are attached to standard Azure subscriptions, are located in a resource group, and can receive tags just like any other Azure resource.

To connect a Kubernetes cluster to Azure, the cluster administrator needs to deploy agents. These agents run in a Kubernetes namespace named `azure-arc` and are standard Kubernetes deployments. The agents are responsible for connectivity to Azure, collecting Azure Arc logs and metrics, and watching for configuration requests.

You can deploy the Azure Arc agents either as part of the cluster creation process (by including the `azure-arc-onboarding` addon spec in your input `apimodel.json`) or manually using [azure-cli](https://docs.microsoft.com/en-us/azure/azure-arc/kubernetes/connect-cluster).

## Azure Arc enabled Kubernetes Addon

The `azure-arc-onboarding` addon creates a Kubernetes job (in namespace `azure-arc-onboarding`) in charge of deploying the Azure Arc agents.
The following information is required in order to successfully onboard the new cluster.

| Name             | Required | Description                                                                            |
| ---------------- | -------- | -------------------------------------------------------------------------------------- |
| location         | yes      | Azure region where the `connectedCluster` ARM resource will be created                 |
| subscriptionID   | yes      | Subscription ID where the `connectedCluster` ARM resource will be created              |
| tenantID         | yes      | Tenant ID that owns the specified Subscription                                         |
| resourceGroup    | yes      | Existing resource group name where the `connectedCluster` ARM resource will be created |
| clusterName      | yes      | Unique cluster friendly name                                                           |
| clientID         | yes      | Service principal ID with permissions to create resources in target subscription/group |
| clientSecret     | yes      | Service principal secret                                                               |

Example:

```json
{
    "name": "azure-arc-onboarding",
    "enabled": true,
    "config": {
        "tenantID": "88e66958-71dd-48b9-8fed-99e13b5c0a59",
        "subscriptionID": "88e66958-71dd-48b9-8fed-99e13b5c0a59",
        "resourceGroup": "connectedClusters",
        "clusterName": "clusterName",
        "clientID": "88e66958-71dd-48b9-8fed-99e13b5c0a59",
        "clientSecret": "88e66958-71dd-48b9-8fed-99e13b5c0a59",
        "location": "eastus"
    }
}
```

### Validation / Troubleshooting

To make sure that the onboarding process succeded, you can either look for the new `connectedCluster` resource in the Azure portal
(ARM ID: `/subscriptions/{subscriptionID}/resourceGroups/{resourceGroup}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}`)
or check the status of the agent pods in the `azure-arc` namespace.

```bash
kubectl get pods -n azure-arc
```

If you notice something wrong, the first troubleshooting step would be to inspect the logs produced by the onboarding process

```bash
kubectl logs -l job-name=azure-arc-onboarding -n azure-arc-onboarding
```

#### Frequent issues

Potential issues you may find by inspecting the job logs include:

- Target resource group does not exit
- Cluster name is not unique
- Invalid service principal credentials
- Service principal does not have enough permissions to create resources in target subscription or resource group
- Azure Arc is not available in the desired Azure region

### Clean up

You are free to delete the resources created in namespace `azure-arc` created by job `azure-arc-onboarding`.

However, you won't be able to permanently delete the resources created in namespace `azure-arc-onboarding`
until file `arc-onboarding.yaml` is moved out of directory `/etc/kubernetes/addons` (control plane nodes' file system)
as `addon-manager` will re-create the resources in namespace `azure-arc-onboarding`.

### Addon Reconfiguration

There are two different ways to reconfigure the `azure-arc-onboarding` addon the cluster is deployed.

The safer and recommended approach is to update, on every control plane node,
the secret resource declared in the addon manifest (`/etc/kubernetes/addons/arc-onboarding.yaml`)
and re-trigger the onboarding process by deleting the `azure-arc-onboarding` namespace.

A faster and more fragile alternative is to edit the secret using kubectl
(`kubectl edit secret azure-arc-onboarding -n azure-arc-onboarding`) and
and re-trigger the onboarding process by deleting the onboarding job
(`kubectl delete job azure-arc-onboarding -n azure-arc-onboarding`).
Keep in mind that your changes will be lost if the secret resource is deleted at any point in the future
as `addon-manager` will recreate it using the data in `arc-onboarding.yaml`.

More information on how to edit a Kubernetes secret can be found [here](https://kubernetes.io/docs/concepts/configuration/secret/#creating-a-secret-manually).
