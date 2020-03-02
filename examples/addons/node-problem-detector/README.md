# Node Problem Detector Add-on

[node-problem-detector](https://github.com/kubernetes/node-problem-detector) makes common Kubernetes node problems visible to the cluster management stack. It is a daemon which runs on each node, detects problems, and reports them to the apiserver.

For more details on the Kubernetes node-problem-detector, see the [GitHub project](https://github.com/kubernetes/node-problem-detector).

The following is a sample API definition with the node-problem-detector addon enabled.

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "addons": [
          {
            "name": "node-problem-detector",
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
        "availabilityProfile": "AvailabilitySet"
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

You can validate that the add-on is running as expected with the following command. You should see a node-problem-detector pod running for each agent node in the cluster.

```bash
kubectl get pods -n kube-system
```

To test node-problem-detector in a running cluster, you can inject messages into the logs it is watching. See [Try It Out](https://github.com/kubernetes/node-problem-detector#try-it-out) in the node-problem-detector documentation for details.

## Configuration

| Name                | Required | Description                                                | Default Value                                                                        |
| ------------------- | -------- | ---------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| customPluginMonitor | no       | Comma-separated list of custom plugin monitor config files | /config/kernel-monitor-counter.json,/config/systemd-monitor-counter.json             |
| systemLogMonitor    | no       | Comma-separated list of system log monitor config files    | /config/kernel-monitor.json,/config/docker-monitor.json,/config/systemd-monitor.json |
| systemStatsMonitor  | no       | Comma-separated list of system stats monitor config files  | /config/system-stats-monitor.json                                                    |
| versionLabel        | no       | Version label used as DaemonSet selector                   | v0.8.1                                                                               |

### Node Problem Detector

| Name           | Required | Description                       | Default Value                             |
| -------------- | -------- | --------------------------------- | ----------------------------------------- |
| name           | no       | container name                    | "node-problem-detector"                   |
| image          | no       | image                             | "k8s.gcr.io/node-problem-detector:v0.8.1" |
| cpuRequests    | no       | cpu requests for the container    | "20m"                                     |
| memoryRequests | no       | memory requests for the container | "20Mi"                                    |
| cpuLimits      | no       | cpu limits for the container      | "200m"                                    |
| memoryLimits   | no       | memory limits for the container   | "100Mi"                                   |

## Supported Orchestrators

Kubernetes

## Contact

- If you have any questions or feedback regarding the node-problem-detector addon, please file an issue at https://github.com/Azure/aks-engine/issues
