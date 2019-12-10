# Use Key Vault as the Source of Cluster Configuration Secrets

## Overview

AKS Engine enables you to source the following cluster configuration from Microsoft Azure Key Vault.

For official Azure Key Vault documentation go [here](https://docs.microsoft.com/en-us/azure/key-vault/basic-concepts).

In order to use Key Vault as the source of cluster configuration secrets, you pass in a reference to the secret URI in your api model:


```json
{
...
    "servicePrincipalProfile": {
        "clientId": "ServicePrincipalClientID",
        "keyvaultSecretRef": {
            "vaultID": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>",
            "secretName": "<NAME>",
            "version": "<VERSION>"
        }
    },
    "certificateProfile": {
        "caCertificate": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<CA_CRT_NAME>",
        "caPrivateKey": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<CA_KEY_NAME>",
        "apiServerCertificate": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<APISERVER_CRT_NAME>",
        "apiServerPrivateKey": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<APISERVER_KEYNAME>",
        "clientCertificate": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<CLIENT_CRT_NAME>",
        "clientPrivateKey": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<CLIENT_KEY_NAME>",
        "kubeConfigCertificate": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<KUBE_CRT_NAME>",
        "kubeConfigPrivateKey": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<KUBE_KEY_NAME>",
        "etcdServerCertificate": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDSERVER_CRT_NAME>",
        "etcdServerPrivateKey": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDSERVER_KEY_NAME>",
        "etcdClientCertificate": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDCLIENT_CRT_NAME>",
        "etcdClientPrivateKey": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDCLIENT_KEY_NAME>",
        "etcdPeerCertificates": [
            "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDPEER0_CRT_NAME>",
            "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDPEER1_CRT_NAME>",
            "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDPEER2_CRT_NAME>"
        ],
        "etcdPeerPrivateKeys": [
            "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDPEER0_KEY_NAME>",
            "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDPEER1_KEY_NAME>",
            "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<ETCDPEER2_KEY_NAME>"
        ]
    }
  }
}
```

## Certificate Profile

For parameters referenced in the `properties.certificateProfile` section of the api model file, the value of each field should be formatted as:

```json
{
  "<PARAMETER>": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<NAME>[/<VERSION>]"
}
```

where:

* `SUB_ID` - is the subscription ID of the Key Vault
* `RG_NAME` - is the resource group of the Key Vault
* `KV_NAME` - is the name of the Key Vault
* `NAME` - is the name of the secret in the Key Vault
* `VERSION` (optional) - is the version of the secret (default: the latest version)

## Service Principal Profile

Passing in a service principal secret via a Key Vault reference looks a bit different:

```json
{
  "servicePrincipalProfile": {
    "clientId": "97ffd212-b56b-430a-97bd-9d15cc01ed43",
    "keyvaultSecretRef": {
      "vaultID": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>",
      "secretName": "<NAME>",
      "version": "<VERSION>"
    }
  }
}
```

Note: the version field is optional.

**Important** The secrets in the Key Vault for the Certificates and Private Keys must be Base64 encoded, and all on a single line -- this means you can't use the `--encoding base64` option of the Azure CLI. Instead you should use the `base64` command:

#### The default _MacOS_ version of the base64 CLI tool will not wrap by default
```sh
az keyvault secret set --vault-name KV_NAME --name NAME --value "$(cat ca.crt | base64 --break=0)"
```

#### The default base64 version of the base64 utility _will wrap by at 76 chars by default on most Linux distros_
```sh
az keyvault secret set --vault-name KV_NAME --name NAME --value "$(cat ca.crt | base64 --wrap=0)"
```

## Key Vault Configuration

To enable Azure Resource Manager to retrieve the secrets from Key Vault, template deployment must be enabled for each Key Vault secret that is referenced:

```sh
az keyvault update -g $RG_NAME -n $KV_NAME --enabled-for-template-deployment
```

## Upgrade Considerations

**Important** As of now, there is no working cluster upgrade implementation for clusters that were built according to the Key Vault-derived secrets. If you expect to use `aks-engine upgrade` to maintain a long-lived Kubernetes cluster built by `aks-engine`, you will not be able to take advantage of this feature at this time.
