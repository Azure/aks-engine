# Logs Retrieval

## Motivation

People new to Kubernetes (or Linux) may find challenging the retrieval of relevant logs to troubleshoot an AKS Engine operation failure.

Hence, it would be convenient to that audience to embedded the log collection inside AKS Engine itself.

## High level description

Just for consistency, the command interface could be similar to `rotate-certs`’ interface:

```
aks-engine get-logs
    --location ${AZURE_LOCATION}
    --api-model _out/prefix/apimodel.json
    --resource-group prefix
    --output-directory _out/prefix/logs_${STAMP}
    --client-id ${AZURE_CLIENT_ID}
    --client-secret ${AZURE_CLIENT_SECRET}
    --subscription-id ${AZURE_CLIENT_SUBSCRIPTION}
    --ssh-private-key ~/.ssh/id_rsa
    --apiserver prefix.location.cloudapp.azure.com
```

After a successful execution, the output directory could contain this (non-exhaustive) list of files:

- Output of kubectl cluster-info dump
- State of the Azure resources in the resource group    
- For each host
  - Files in directory /var/log
  - Files in directory /var/log/azure
  - Files in directory /etc/kubernetes/manifests
  - Files in directory /etc/kubernetes/addons
  - apimodel.json (no creds)
  - azure.json (no creds)
  - kubelet journal
  - container runtime journal
  - etcd journal
  - test DNS resolution
  - test metadata endpoint response  

## Implementation

Just like `rotate-cert`, this new command would SSH into each cluster hosts (Linux and/or Windows) and execute a “collect-logs” script. Additionally, we will have to write a small function that downloads a tar/zip that contains the files produced by the log collection script (SCP).

Because this would be expected to work on an air-gapped environment, we should either:

- Bake the required script/s onto the VHD in a well-known location
  - This is already done for Windows nodes: `c:\k\debug\collect-windows-logs.ps1`
- Add an extra flag to pass the required script/s to the CLI command
  - `--log-collection-script ~/collectlogs.sh`

In connected environments, AKS Engine or each host can potentially download the latest and greatest version of the log collection script.

Files could be uploaded to a storage account container to simplify colaboration.

Additionally, for this to work on Windows nodes we have to setup a SSH server by default.
