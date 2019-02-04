# Using the Container Monitoring addon (Azure Monitor for containers)

Container Monitoring addon gives you performance monitoring ability by collecting memory and processor metrics from controllers, nodes, and containers that are available in Kubernetes through the Metrics API. After you enable Container Monitoring addon, these metrics are automatically collected for you through a containerized version of the Log Analytics agent for Linux and stored in your [Log Analytics] workspace. The included pre-defined views display the residing container workloads and what affects the performance health of the Kubernetes cluster so that you can:

- Identify containers that are running on the node and their average processor and memory utilization. This knowledge can help you identify resource bottlenecks.
- Identify where the container resides in a controller or a pod. This knowledge can help you view the controller's or pod's overall performance.
- Review the resource utilization of workloads running on the host that are unrelated to the standard processes that support the pod.
- Understand the behavior of the cluster under average and heaviest loads. This knowledge can help you identify capacity needs and determine the maximum load that the cluster can sustain.
- Logs (stdout/stderr) of the Containers to troubleshoot the issues in containers

## Components

Your ability to monitor performance relies on a containerized Log Analytics agent for Linux, which collects performance and event data from all nodes in the cluster. The agent is automatically deployed and registered with the specified Log Analytics workspace after you enable container monitoring addon and specify the right encoded workspaceid and workspace key in the addon config.

    "name": "container-monitoring",
    "enabled": true,
    "config": {
      "workspaceGuid": "Base-64 encoded workspace guid",
      "workspaceKey": "Base 64 encoded workspace key"
    }

## Onboarding

Follow the instructions to [onboard the container monitoring addon for the AKS Engine cluster(s)](../../examples/addons/container-monitoring/README.md)
