# Using the Container Monitoring addon (Azure Monitor for containers)

Container Monitoring addon gives you performance monitoring ability by collecting memory and processor metrics from controllers, nodes, and containers that are available in Kubernetes through the Metrics API. After you enable Container Monitoring addon, these metrics are automatically collected for you through a containerized version of the Log Analytics agent for Linux and stored in your [Log Analytics] workspace. The included pre-defined views display the residing container workloads and what affects the performance health of the Kubernetes cluster so that you can:

- Identify containers that are running on the node and their average processor and memory utilization. This knowledge can help you identify resource bottlenecks.
- Identify where the container resides in a controller or a pod. This knowledge can help you view the controller's or pod's overall performance.
- Review the resource utilization of workloads running on the host that are unrelated to the standard processes that support the pod.
- Understand the behavior of the cluster under average and heaviest loads. This knowledge can help you identify capacity needs and determine the maximum load that the cluster can sustain.
- Logs (stdout/stderr) of the Containers to troubleshoot the issues in containers
- Scrapping metrics from Prometheus. Refer [configuring prometheus scrapping settings](https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-agent-config#overview-of-configurable-prometheus-scraping-settings)

Refer (Azure Monitor for containers)[https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-overview] for more details.

## Components

Your ability to monitor performance relies on a containerized Log Analytics agent for Linux, which collects performance and event data from all nodes, pods and containers in the cluster.
The agent is automatically deployed and registered with the specified or defaut Log Analytics workspace after you enable container monitoring addon.

## Onboarding

Container Monitoring can be onboarded either through the monitoring add-on or the HELM chart.
Following are supported options to enable container-monitoring add-on during the cluster creation or post custer creation.

> Note: option 1) and 2) are supported only through `aks-engine deploy` command.

### 1. Using Default Log Analytics Workspace

Container Monitoring add-on uses default log analytics workspace if its available for the corresponding region of the cluster. If there is no default log analytics workspaces exists for the corresponding region of the cluster, then it will create new default log analytics workspace.
With following definition, container monitoring addon will be onboarded fully without need of any additional steps such as adding container insights solution and tags etc.

        {
            "name": "container-monitoring",
            "enabled": true
        }

Refer [Sample Kubernetes definition file with monitoringa addon using default log analytics workspace](../../examples/addons/container-monitoring/kubernetes-container-monitoring.json)

                     OR

###  2. Using existing Log Analytics Workspace

If you prefer to use existing log analytics workspace, then you can set  the fully qualified azure resource id of the existing log analytics workspace as in below config.
With following definition, container monitoring addon will be onboarded fully without need of any additional steps such as adding container insights solution and tags etc.
Azure Log analytics workspace can be in any Azure subscription which you have Log Analytics Contributor role permission on the specified Azure Log Analytics workspace.

      {
            "name": "container-monitoring",
            "enabled": true,
             "config": {
                 "logAnalyticsWorkspaceResourceId" : "/subscriptions/<WorkspaceSubscriptionId>/resourceGroups/<WorkspaceResourceGroup>/providers/Microsoft.OperationalInsights/workspaces/<workspaceName>"
             }
     }

Refer [Sample Kubernetes definition file with monitoringa addon using existing log analytics workspace](../../examples/addons/container-monitoring/kubernetes-container-monitoring_existing_log_analytics_workspace.json)

                 OR

### 3. Using Worksapce GUID or Key

you can also configure with workspace GUID and Key of the existing Log analytics workspace, but with this is legacy option and requires additional configuration steps

         {
            "name": "container-monitoring",
            "enabled": true,
            "config": {
              "workspaceGuid": "<Azure Log Analytics Workspace Guid in Base-64 encoded>",
              "workspaceKey": "<Azure Log Analytics Workspace Key in Base-64 encoded>"
            }
          }

Refer [Sample Kubernetes definition file with monitoringa addon using workspace GUID and key of the existing log analytics workspace](../../examples/addons/container-monitoring/kubernetes-container-monitoring_existing_workspace_id_and_key.json)

### 4. Using Azure Monitor for containers HELM chart

If you prefer HELM chart, refer [Azure Monitor – Containers HELM chart](https://github.com/helm/charts/tree/master/incubator/azuremonitor-containers) for onboarding instructions

For more details and instructions to [onboard the container monitoring addon for the AKS Engine cluster(s)](../../examples/addons/container-monitoring/README.md)

## UX

After successful onboarding, navigating to [Azure Monitor for containers](https://aka.ms/azmon-containers) to view and monitor, and analyze health of your onboarded AKS-Engine cluster, pods and containers etc.

## Required Roles and Permissions

- User requires the reader role permission on the Azure Log Analytics workspace and AKS-Engine cluster resource group to view and monitor, and analyze health of your onboarded AKS-Engine cluster, pods and containers etc.
- For onboarding monitoring addon
     -  If the existing Azure Log Analytics workspace used, then requires Log Analytics Contributor role on existing Azure Log Analytics
     -  For the new Azure Log Analytics workspace, user requires the contributor role on the Subscription or the Resource group where the AKS-Engine cluster resources will be deployed

## Contact

If you have any questions or feedback regarding the container monitoring addon, please reach us out through [this](mailto:askcoin@microsoft.com) email.