# AAD Pod Identity Add-on


This is the AAD Pod Identity add-on.  Add this add-on to your json file as shown below to automatically enable AAD Pod identity in your new Kubernetes cluster.
> Note: At the moment AAD Pod Identity supports only Availability Set and is tested only for Linux based clusters.

```json
{
    "apiVersion": "vlabs",
    "properties": {
      "orchestratorProfile": {
        "orchestratorType": "Kubernetes",
        "kubernetesConfig": {
        "addons": [
          {
            "name": "aad-pod-identity",
            "enabled": true
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
      }
    }
  }

```

You can validate that the add-on is running as expected with the following commands:

You should see two sets of pods - a single mic pod and as many nmi pods as there are agent nodes in 'Running' state after executing:

```bash
kubectl get pods
```

Plese follow the README here for further infromation: https://github.com/Azure/aad-pod-identity

## Supported Orchestrators

Kubernetes
