# NVIDIA Device Plugin

This is the [NVIDIA Device Plugin](https://github.com/NVIDIA/k8s-device-plugin) add-on for Kubernetes. This add-on will be automatically enabled if you are using a Kubernetes cluster (v1.10+) with an N-series agent pool (which contains an NVIDIA GPU). You can use this add-on to your json file as shown below to enable or disable NVIDIA Device Plugin explicitly.

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "orchestratorRelease": "1.13",
      "kubernetesConfig": {
        "addons": [
          {
            "name": "nvidia-device-plugin",
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
        "vmSize": "Standard_NC6"
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

You can validate that the add-on is running as expected with the following command.

You should see NVIDIA Device Plugin pods as running after executing:

```bash
kubectl get pods -n kube-system
```

Follow the README at [NVIDIA/k8s-device-plugin](https://github.com/NVIDIA/k8s-device-plugin) for more information.

## Supported Orchestrators

* Kubernetes
