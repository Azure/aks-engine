# Azure Policy Add-on

This is the Azure Policy add-on. Add this add-on to your json file as shown below to automatically enable Azure Policy with OPA Gatekeeper in your new Kubernetes cluster.

This add-on works with a service principal only at this time. Make sure to create a role assignment for service principal using:

```
CLIENT_ID=<your service principal id>
SCOPE=<fully qualified resource id for the cluster. for example, "/subscriptions/<subscriptionId>/resourceGroups/<clusterResourceGroup>">

az role assignment create --assignee $CLIENT_ID --scope $SCOPE --role "Policy Insights Data Writer (Preview)"`
```

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
            "enabled": true,
            "config": {
              "auditInterval": "30",
              "constraintViolationsLimit": "20"
            }
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
kubectl get pods -n kube-system
```

Plese follow the README here for further information: https://github.com/open-policy-agent/gatekeeper

## Configuration

| Name                      | Required | Description                 | Default Value |
| ------------------------- | -------- | --------------------------- | ------------- |
| auditInterval             | no       | audit interval (in seconds) | 30            |
| constraintViolationsLimit | no       | constraint violations limit | 20            |

### Azure Policy

| Name           | Required | Description                       | Default Value                                                                 |
| -------------- | -------- | --------------------------------- | ----------------------------------------------------------------------------- |
| name           | no       | container name                    | "azure-policy"                                                                |
| image          | no       | image                             | "mcr.microsoft.com/azure-policy/policy-kubernetes-addon-prod:prod_20191011.1" |
| cpuRequests    | no       | cpu requests for the container    | "30m"                                                                         |
| memoryRequests | no       | memory requests for the container | "50Mi"                                                                        |
| cpuLimits      | no       | cpu limits for the container      | "100m"                                                                        |
| memoryLimits   | no       | memory limits for the container   | "200Mi"                                                                       |

### Gatekeeper

| Name           | Required | Description                       | Default Value                                        |
| -------------- | -------- | --------------------------------- | ---------------------------------------------------- |
| name           | no       | container name                    | "gatekeeper"                                         |
| image          | no       | image                             | "quay.io/open-policy-agent/gatekeeper:v3.0.4-beta.2" |
| cpuRequests    | no       | cpu requests for the container    | "100m"                                               |
| memoryRequests | no       | memory requests for the container | "256Mi"                                              |
| cpuLimits      | no       | cpu limits for the container      | "100m"                                               |
| memoryLimits   | no       | memory limits for the container   | "512Mi"                                              |

## Supported Orchestrators

Kubernetes

## Contact

If you have any questions or feedback regarding Gatekeeper, please file an issue at https://github.com/open-policy-agent/gatekeeper/issues
If you have any questions or feedback regarding Azure Policy addon, please reach us out through this email.
