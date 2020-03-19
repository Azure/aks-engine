# CSI Secrets Store Add-on

[The CSI Secrets Store](https://github.com/kubernetes-sigs/secrets-store-csi-driver) integrates secrets stores with Kubernetes via a [Container Storage Interface (CSI)](https://kubernetes-csi.github.io/docs/) volume. 

With the CSI secrets store and the azure provider installed, developers can access application-specific secrets, keys, and certs stored in Azure Key Vault directly from their pods.

This addon is available for 1.15+ clusters

**Note** This addon is enabled by default for 1.16+ clusters

To enable on 1.15 cluster, add this addon to the apimodel and also set the feature gate `CSIInlineVolume=true` for API server and kubelet config

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "orchestratorRelease": "1.15",
      "kubernetesConfig": {
        "addons": [
          {
            "name": "csi-secrets-store",
            "enabled": true
          }
        ],
        "apiServerConfig": {
          "--feature-gates": "CSIInlineVolume=true"
        },
        "kubeletConfig": {
          "--feature-gates": "CSIInlineVolume=true"
        }
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

To validate the add-on is running as expected, run the following commands:

You should see the csi-secrets-store driver pods and azure provider pods running on each agent node:

```bash
kubectl get pods -n kube-system 

csi-secrets-store-4vmbw                         3/3     Running   0          43m
csi-secrets-store-kmjcr                         3/3     Running   0          50m
csi-secrets-store-provider-azure-7ldqq          1/1     Running   0          43m
csi-secrets-store-provider-azure-h5xmh          1/1     Running   0          50m
```

Follow the README at https://github.com/Azure/secrets-store-csi-driver-provider-azure for get started steps.

## Supported Orchestrators

Kubernetes