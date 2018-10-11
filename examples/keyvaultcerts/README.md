# Microsoft Azure Kubernetes Service Engine - Key vault certificate deployment

## Overview

AKS Engine enables you to create a customized Kubernetes cluster on Microsoft Azure with certs installed from key vault during deployment.

The examples show you how to configure installing a cert from keyvault. These certs are assumed to be in the secrets portion of your keyvault:

1. **kubernetes.json** - deploying and using [Kubernetes](../../docs/kubernetes.md)

On windows machines certificates will be installed under the machine in the specified store.
On linux machines the certificates will be installed in the folder /var/lib/waagent/. There will be two files
1. {thumbprint}.prv - this will be the private key pem formatted
2. {thumbprint}.crt - this will be the full cert chain pem formatted
