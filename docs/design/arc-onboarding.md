# Azure Arc for Kubernetes integration

## Motivation

Simplify Arc agents deploymet by making it an optional last step of the cluster deployment process.

## High level implementation

The goal is to reuse the existing `cluster-init` component.

- First, add `connectedClusterProfile` to apimodel

```json
{
    "connectedClusterProfile": {
      "tenantID": "azure-spn-tenant",
      "subscriptionID": "azure-sub",
      "resourceGroup": "connected-cluster-rg-azure",
      "clusterName": "friendly-name",
      "clientID": "azure-spn",
      "clientSecret": "azure-spn-secret",
      "region": "azure-region"
    }
}
```

- AKSe replaces the placeholders in [arc-onboarding.yml](https://github.com/Azure/azure-arc-kubernetes-onboarding/blob/master/arc-onboarding.yml), base64 encodes the yaml and appends the encoded string to the `cluster-init` component (enabling it if disabled).
