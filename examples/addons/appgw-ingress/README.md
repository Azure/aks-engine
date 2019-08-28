# AppGW Ingress Add-on

This add-on will deploy an Application Gateway and dependency resources with your new Kubernetes cluster.

Following resources are deployed:

1) Azure Application Gateway v2
2) Public IP Address
3) AppGw Subnet in the shared vnet.
4) User Assigned Identity to initialize the aad-pod-identity service and ingress controller.
5) Set required RBACs.

Supported Add-on `Config` options:

| Option | Required | Default | Description |
|--|--|--|--|
| `appgw-subnet` | true | N/A | CIDR of the Application Gateway subnet. This should not overlap with master/agent subnets. |
| `appgw-sku` | false | `WAF_v2` | SKU of the Application Gateway. (`Standard_v2`/`WAF_v2`) |
| `appgw-private-ip` | false | null | Private IP assigned to the Application Gateway from subnet. |

Once, the infrastructure is deployed, please follow the instructions to deploy the [Application Gateway  Ingress controller](https://github.com/Azure/application-gateway-kubernetes-ingress/blob/master/docs/install-new.md#setting-up-application-gateway-ingress-controller-on-aks)

> Note: Adding the add-on multiple times will not create multiple applicaiton gateways.

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "networkPlugin": "azure",
        "addons": [
          {
            "name": "appgw-ingress",
            "enabled": true,
            "config": {
              "appgw-subnet": "<subnet CIDR>",
              "appgw-sku": "WAF_v2"
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
      "secret": "",
      "objectId": "<>"
    }
  }
}
```

## Supported Orchestrators

Kubernetes
