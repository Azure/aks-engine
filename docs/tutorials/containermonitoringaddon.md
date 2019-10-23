# Using the Container Monitoring addon (Azure Monitor for containers)

Container Monitoring addon gives you performance monitoring ability by collecting memory and processor metrics from controllers, nodes, and containers that are available in Kubernetes through the Metrics API. After you enable Container Monitoring addon, these metrics are automatically collected for you through a containerized version of the Log Analytics agent for Linux and stored in your [Log Analytics] workspace. The included pre-defined views display the residing container workloads and what affects the performance health of the Kubernetes cluster so that you can:

- Identify containers that are running on the node and their average processor and memory utilization. This knowledge can help you identify resource bottlenecks.
- Identify where the container resides in a controller or a pod. This knowledge can help you view the controller's or pod's overall performance.
- Review the resource utilization of workloads running on the host that are unrelated to the standard processes that support the pod.
- Understand the behavior of the cluster under average and heaviest loads. This knowledge can help you identify capacity needs and determine the maximum load that the cluster can sustain.
- Logs (stdout/stderr) of the Containers to troubleshoot the issues in containers
- Scraping metrics from Prometheus. Refer [configuring prometheus scraping settings](https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-agent-config#overview-of-configurable-prometheus-scraping-settings)

Refer to [Azure Monitor for containers](https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-overview) for more details.

## Components

Your ability to monitor performance relies on a containerized Log Analytics agent for Linux, which collects performance and event data from all nodes, pods and containers in the cluster.
The agent is automatically deployed and registered with the specified or default Log Analytics workspace after you enable container monitoring addon.

## Onboarding

Container Monitoring can be onboarded either through the monitoring add-on or the Helm chart.
Following are supported options to enable container-monitoring add-on during the cluster creation or post cluster creation.

> Note: option 1) and 2) are supported only through `aks-engine deploy` command.
> Note: For Azure stack cloud environments, only option 3) and 4) are supported.

### 1. Using Default Log Analytics Workspace

Container Monitoring add-on uses default log analytics workspace if its available for the corresponding region of the cluster. If there exist no default log analytics workspaces for the corresponding region of the cluster, then it will create new default log analytics workspace.
With the following definition, container monitoring addon will be onboarded fully without need of any additional steps such as adding container insights solution and tags etc.

        {
            "name": "container-monitoring",
            "enabled": true
        }

Refer to [Sample Kubernetes definition file with monitoringa addon using default log analytics workspace](../../examples/addons/container-monitoring/kubernetes-container-monitoring.json)

###  2. Using existing Log Analytics Workspace

If you prefer to use an existing log analytics workspace, then you can set  the fully qualified azure resource id of the existing log analytics workspace as in the below config.
With this definition, container monitoring addon will be onboarded fully without need of any additional steps such as adding container insights solution and tags etc.
Azure Log analytics workspace can be in any Azure subscription in which you have Log Analytics Contributor role permission on the specified Azure Log Analytics workspace.

      {
            "name": "container-monitoring",
            "enabled": true,
             "config": {
                 "logAnalyticsWorkspaceResourceId" : "/subscriptions/<WorkspaceSubscriptionId>/resourceGroups/<WorkspaceResourceGroup>/providers/Microsoft.OperationalInsights/workspaces/<workspaceName>"
             }
     }

Refer to [Sample Kubernetes definition file with monitoring addon using existing log analytics workspace](../../examples/addons/container-monitoring/kubernetes-container-monitoring_existing_log_analytics_workspace.json)

### 3. Using Workspace GUID or Key

You can also configure with workspace GUID and Key of the existing Log analytics workspace, but this is a legacy option and requires additional configuration steps

         {
            "name": "container-monitoring",
            "enabled": true,
            "config": {
              "workspaceGuid": "<Azure Log Analytics Workspace Guid in Base-64 encoded>",
              "workspaceKey": "<Azure Log Analytics Workspace Key in Base-64 encoded>"
            }
          }

Refer to [Sample Kubernetes definition file with monitoring addon using workspace GUID and key of the existing log analytics workspace](../../examples/addons/container-monitoring/kubernetes-container-monitoring_existing_workspace_id_and_key.json)

### 4. Using Azure Monitor for containers Helm chart

If you prefer Helm chart, refer to [Azure Monitor – Containers Helm chart](https://github.com/Helm/charts/tree/master/incubator/azuremonitor-containers) for onboarding instructions

For more details and instructions to [onboard the container monitoring addon for the AKS Engine cluster(s)](../../examples/addons/container-monitoring/README.md)

## UX

After successful onboarding, navigating to [Azure Monitor for containers](https://aka.ms/azmon-containers) to view and monitor, and analyze health of your onboarded AKS Engine cluster, pods and containers etc.

## Required Roles and Permissions

- For onboarding monitoring addon
     -  If the existing Azure Log Analytics workspace is used, then the Log Analytics Contributor role on existing Azure Log Analytics is required
     -  For the new Azure Log Analytics workspace, user requires the contributor role on the Subscription or the Resource group where the AKS Engine cluster resources will be deployed

  - User requires the reader role permission on the Azure Log Analytics workspace
  - For Azure AKS-Engine clusters, user requires reader role permission on cluster resource group and resources under that

## Supported Azure Cloud Environment(s)

 -  Azure Public Cloud
 -  Azure China Cloud
 -  Azure US Government Cloud
 -  Azure Stack Cloud

### Disable Monitoring

After you enable monitoring of your AKS Engine cluster, you can stop container monitoring on the cluster if you decide you no longer want to monitor it.

- If you have onboarded the monitoring using the HELM chart, then you can disable monitoring by uninstalling the chart. Refer Uninstalling the Chart section in [azuremonitor-containers](https://github.com/helm/charts/tree/master/incubator/azuremonitor-containers)

- If you have onboarded using the Container Monitoring addon, then you can remove monitoring addon with below steps

      1. ssh to master node of your AKS Engine cluster master node and navigate to /etc/kubernetes/addons directory
      2. delete all the resources related to container monitoring addon with `kubectl delete -f omsagent-daemonset.yaml` command against your AKS Engine cluster
      3. delete the container monitoring addon manifest file omsagent-daemonset.yaml  under /etc/kubernetes/addons

### Upgrade Container Monitoring Addon

For upgrading the container monitoring addon, you can disable the monitoring addon as described in Disable Monitoring section and use the HELM chart to install and upgrade.

## Contact

If you have any questions or feedback regarding the container monitoring addon, please reach us out through [this](mailto:askcoin@microsoft.com) email.