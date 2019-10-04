# Azure Policy Add-on

This is the Azure Policy add-on. Add this add-on to your json file as shown below to automatically enable Azure Policy with OPA Gatekeeper in your new Kubernetes cluster.

> 🚨 Please note that this add-on is in alpha stage and is not production ready.

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "addons": [
          {
            "name": "azure-policy",
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

You can validate that the add-on is running as expected with the following commands:

You should see gatekeeper and azure-policy pods:

```bash
kubectl get pods
```

Plese follow the README here for further information: https://github.com/open-policy-agent/gatekeeper

## Supported Orchestrators

Kubernetes

## Contact

If you have any questions or feedback regarding Gatekeeper, please file an issue at https://github.com/open-policy-agent/gatekeeper/issues
If you have any questions or feedback regarding Azure Policy addon, please reach us out through this email.
