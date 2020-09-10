# ACI Connector Add-on


This is the ACI Connector add-on.  Add this add-on to your json file as shown below to automatically enable ACI Connector in your new Kubernetes cluster.

```json
{
    "apiVersion": "vlabs",
    "properties": {
      "orchestratorProfile": {
        "orchestratorType": "Kubernetes",
        "kubernetesConfig": {
          "addons": [
            {
              "name": "aci-connector",
              "enabled" : true,
              "config": {
                  "region": "",
                  "nodeName": "",
                  "os": "",
                  "taint": ""
              }
            }
          ]
        }
      },
      "masterProfile": {
        "count": 1,
        "dnsPrefix": "",
        "vmSize": "Standard_DS2_v2",
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

You should see ACI Connector as `Running` after executing:

```bash
kubectl get pods -n kube-system
```

You should see ACI Connector node after executing:

```bash
kubectl get nodes
```

Follow the README at https://github.com/virtual-kubelet/virtual-kubelet for examples.

## Configuration

|Name|Required|Description|Default Value|
|---|---|---|---|
|region|no|Azure region|"westus"|
|nodeName|no|node name|"aci-connector"|
|os|no|operating system (Linux/Windows)|"Linux"|
|taint|no|apply taint to node, making scheduling explicit|"azure.com/aci"|
|name|no|container name|"aci-connector"|
|image|no|image|"mcr.microsoft.com/oss/virtual-kubelet/virtual-kubelet:latest"|
|cpuRequests|no|cpu requests for the container|"50m"|
|memoryRequests|no|memory requests for the container|"150Mi"|
|cpuLimits|no|cpu limits for the container|"50m"|
|memoryLimits|no|memory limits for the container|"150Mi"|

## Supported Orchestrators

Kubernetes
