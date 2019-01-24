# Using the Container Monitoring Add-on (Container Health)

Container health gives you performance monitoring ability by collecting memory and processor metrics from controllers, nodes, and containers that are available in Kubernetes through the Metrics API. After you enable container health, these metrics are automatically collected for you through a containerized version of the Log Analytics agent for Linux and stored in your [Log Analytics] workspace. The included pre-defined views display the residing container workloads and what affects the performance health of the Kubernetes cluster so that you can:

  - Identify containers that are running on the node and their average processor and memory utilization. This knowledge can help you identify resource bottlenecks.
  - Identify where the container resides in a controller or a pod. This knowledge can help you view the controller's or pod's overall performance.
  - Review the resource utilization of workloads running on the host that are unrelated to the standard processes that support the pod.
  - Understand the behavior of the cluster under average and heaviest loads. This knowledge can help you identify capacity needs and determine the maximum load that the cluster can sustain.

# Prerequisites!

  - A Log Analytics workspace. You can create it when you enable monitoring of your new AKS cluster. You can create it through [Azure Resource Manager], through [PowerShell], or in the [Azure portal].
  - You will need to add the "Container Insights" solution to your workspace if you do not have any. Please follow the steps here. - [Add Container Insights]
  - The Log Analytics contributor role, to enable container monitoring. For more information about how to control access to a Log Analytics workspace, see [Manage workspaces].

# Components
- Your ability to monitor performance relies on a containerized Log Analytics agent for Linux, which collects performance and event data from all nodes in the cluster. The agent is automatically deployed and registered with the specified Log Analytics workspace after you enable container monitoring addon and specify the right encoded workspaceid and workspace key in the addon config.

    "name": "container-monitoring",
    "enabled": true,
    "config": {
      "workspaceGuid": "Base-64 encoded workspace guid",
      "workspaceKey": "Base 64 encoded workspace key"
    }
    ### Obtain workspace ID and key
    - In the Azure portal, click All services. In the list of resources, type Log Analytics. As you begin typing, the list filters based on your input. Select Log Analytics.
    - In your list of Log Analytics workspaces, select the workspace you intend on configuring the agent to report to.
    - Select Advanced settings.
    - Select Connected Sources, and then select Linux Servers.
    - Copy and paste into your favorite editor, the Workspace ID and Primary Key.
- Please run this [script] which helps you set the tags to the AKS-engine cluster.

##### After the deployment is complete, you should be able to see all the cluster data here: [Link to Container Health]
##### Any feedback: askcoin@microsoft.com

   [Log Analytics]: <https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-overview>
   [Azure Resource Manager]: <https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-template-workspace-configuration>
   [PowerShell]: <https://docs.microsoft.com/azure/log-analytics/scripts/log-analytics-powershell-sample-create-workspace?toc=%2fpowershell%2fmodule%2ftoc.json>
   [Azure portal]: <https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-quick-create-workspace>
   [Manage workspaces]: <https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-manage-access>
   [Link to Container Health]: <https://aka.ms/azmon-containers>
   [Add Container Insights]: <https://github.com/Microsoft/OMS-docker/blob/ci_feature_prod/docs/solution-onboarding.md>
   [script]: <https://github.com/Microsoft/OMS-docker/blob/ci_feature/docs/attach-monitoring-tags.md>
