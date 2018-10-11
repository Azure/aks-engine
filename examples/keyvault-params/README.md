# Microsoft Azure Kubernetes Service Engine - Key vault referencing for k8s parameters

## Overview

AKS Engine enables you to retrieve the following k8s deployment parameters from Microsoft Azure KeyVault:

* certificateProfile
  * apiServerCertificate
  * apiServerPrivateKey
  * caCertificate
  * caPrivateKey
  * clientCertificate
  * clientPrivateKey
  * kubeConfigCertificate
  * kubeConfigPrivateKey
  * etcdServerCertificate
  * etcdServerPrivateKey
  * etcdClientCertificate
  * etcdClientPrivateKey
  * etcdPeerCertificates (length of array depends on number of master nodes)
  * etcdPeerPrivateKeys (length of array depends on number of master nodes)
* servicePrincipalProfile* (a special case)

## Certificate Profile

For parameters referenced in the `properties.certificateProfile` section of the api model file, the value of each field should be formatted as:

```json
{
  "<PARAMETER>": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<NAME>[/<VERSION>]"
}
```

where:

* `SUB_ID` - is the subscription ID of the keyvault
* `RG_NAME` - is the resource group of the keyvault
* `KV_NAME` - is the name of the keyvault
* `NAME` - is the name of the secret in the keyvault
* `VERSION` (optional) - is the version of the secret (default: the latest version)

## Service Principal Profile

For the service principal profile secret, the keyvault is referenced differently. If embedding the secret as plain text, the secret is set in `properties.servicePrincipalProfile.secret`.

If the secret is stored in a keyvault, it can be referenced as follows:

```json
{
  "servicePrincipalProfile": {
    "clientId": "97ffd212-b56b-430a-97bd-9d15cc01ed43",
      "secret": "",
      "keyvaultSecretRef": {
        "vaultID": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>",
        "secretName": "<NAME>",
        "version": "<VERSION>"
    }
  }
}
```

The version field is optional.

## Example

The example `kubernetes.json` shows you how to refer deployment parameter to a secret in a keyvault.

**Important** The secrets in the KeyVault for the Certificates and Private Keys must be Base64 encoded, and all on a single line -- this means you can't use the `--encoding base64` option of the Azure CLI. Instead you should use the `base64` command:

```sh
  # On OSX base64 will not wrap by default
  az keyvault secret set --vault-name KV_NAME --name NAME --value "$(cat ca.crt | base64 --break=0)"

  # On Linux it will wrap at 76 chars by default
  az keyvault secret set --vault-name KV_NAME --name NAME --value "$(cat ca.crt | base64 --wrap=0)"
```

## KeyVault Configuration

To enable Azure Resource Manager to retrieve the secrets from the KeyVault, template deployment must be enabled on the KeyVault:

```sh
az keyvault update -g $RG_NAME -n $KV_NAME --enabled-for-template-deployment
```
