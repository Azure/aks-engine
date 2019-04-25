# AKS Engine - Key vault certificate deployment

## Overview

AKS Engine enables you to create customized Kubernetes cluster on Microsoft Azure with certs installed from key vault during deployment.

The examples show you how to configure installing a cert from keyvault. These certs are assumed to be in the secrets portion of your keyvault:

1. **dcos.json** - deploying and using [DC/OS](../../docs/dcos.md)
2. **kubernetes.json** - deploying and using [Kubernetes](../../docs/kubernetes.md)
3. **swarm.json** - deploying and using [Swarm](../../docs/swarm.md)
4. **swarm-windows.json** - deploying and using [Swarm](../../docs/swarm.md)
5. **swarmmode.json** - deploying and using [Swarm Mode](../../docs/swarmmode.md)

On windows machines certificates will be installed under the machine in the specified store.
On linux machines the certificates will be installed in the folder /var/lib/waagent/. There will be two files
1. {thumbprint}.prv - this will be the private key pem formatted
2. {thumbprint}.crt - this will be the full cert chain pem formatted
