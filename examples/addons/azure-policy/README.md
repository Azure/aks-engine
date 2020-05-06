# Azure Policy Add-on

Azure Policy integrates with the AKS Engine to apply at-scale enforcements and safeguards on your clusters in a centralized, consistent manner. By extending use of [Open Policy Agent](https://www.openpolicyagent.org/) (OPA) [Gatekeeper](https://github.com/open-policy-agent/gatekeeper) v3 (beta), an _admission controller webhook_ for Kubernetes, Azure Policy makes it possible to manage and report on the compliance state of your Azure resources and AKS Engine clusters from one place.


> [!NOTE]
> Azure Policy for AKS Engine is in Public Preview. The service only supports built-in policy definitions
> and a single AKS Engine cluster for each resource group configured with a Service Principal.
> Gatekeeper v3 is in Beta and is supported by the open source community.


For detailed instructions to enable and use Azure policy add-on for AKS Engine, please refer [Install Azure Policy Add-on on AKS Engine](https://aka.ms/kubepolicydoc).

The following is a sample API definition with azure-policy addon.

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
              "auditInterval": "60",
              "constraintViolationsLimit": "100"
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
> 🚨 Please note that this add-on is in alpha stage and is not production ready.

You can validate that the add-on is running as expected with the following command. You should see gatekeeper and azure-policy pods running.

```bash
kubectl get pods -n kube-system
```

## Configuration

| Name                      | Required | Description                 | Default Value |
| ------------------------- | -------- | --------------------------- | ------------- |
| auditInterval             | no       | audit interval (in seconds) | 60            |
| constraintViolationsLimit | no       | constraint violations limit | 100            |

### Azure Policy

| Name           | Required | Description                       | Default Value                                                                 |
| -------------- | -------- | --------------------------------- | ----------------------------------------------------------------------------- |
| name           | no       | container name                    | "azure-policy"                                                                |
| image          | no       | image                             | "mcr.microsoft.com/azure-policy/policy-kubernetes-addon-prod:prod_20200505.1" |
| cpuRequests    | no       | cpu requests for the container    | "30m"                                                                         |
| memoryRequests | no       | memory requests for the container | "50Mi"                                                                        |
| cpuLimits      | no       | cpu limits for the container      | "100m"                                                                        |
| memoryLimits   | no       | memory limits for the container   | "200Mi"                                                                       |

### Gatekeeper

| Name           | Required | Description                       | Default Value                                        |
| -------------- | -------- | --------------------------------- | ---------------------------------------------------- |
| name           | no       | container name                    | "gatekeeper"                                         |
| image          | no       | image                             | "mcr.microsoft.com/oss/open-policy-agent/gatekeeper:v3.1.0-beta.8" |
| cpuRequests    | no       | cpu requests for the container    | "100m"                                               |
| memoryRequests | no       | memory requests for the container | "256Mi"                                              |
| cpuLimits      | no       | cpu limits for the container      | "1000m"                                              |
| memoryLimits   | no       | memory limits for the container   | "512Mi"                                              |

## Disable Azure Policy Add-on

### Option 1

- Update `apimodel.json` and set `azure-policy` addon `enabled` to `false`.
- Run `aks-engine upgrade` with the updated `apimodel.json`
- Run `kubectl delete deployments.apps -n kube-system gatekeeper-controller-manager azure-policy` to remove running deployments and pods

### Option 2

- SSH into AKS-Engine Kubernetes master node
- Run `kubectl delete -f /etc/kubernetes/addons/azure-policy-deployment.yaml && sudo rm /etc/kubernetes/addons/azure-policy-deployment.yaml` to remove all resources created by addon

## Supported Orchestrators

Kubernetes

## Contact

- If you have any questions or feedback regarding Azure Policy addon, please file an issue at https://github.com/Azure/aks-engine/issues and tag with `azure-policy`

- If you have any questions or feedback regarding Gatekeeper, please file an issue at https://github.com/open-policy-agent/gatekeeper/issues
