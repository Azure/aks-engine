# Retrieving Node and Cluster Logs

## Prerequisites

All documentation in these guides assumes you have already downloaded both the Azure CLI and `aks-engine`. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing.

This guide assumes you already have deployed a cluster using `aks-engine`. For more details on how to do that see [deploy](../tutorials/quickstart.md#deploy).

## Retrieving Logs

The `aks-engine get-logs` command can be useful to troubleshoot issues with your cluster. It will produce, collect and download to your workstation a set of files that include node configuration, cluster state and configuration, and provision log files.

At a high level, it works by establishing a SSH session into each node, executing a [log collection](#log-collection-scripts) script that collects and zips relevant files, and downloading the zip file to your local computer.

### SSH Authentication

A valid SSH private key is always required to stablish a SSH session to the cluster Linux nodes. Windows credentials are stored in the API model and will be loaded from there. Make sure `windowsprofile.sshEnabled` is set to `true` to enable SSH in your Windows nodes.

### Log Collection Scripts

To collect Linux nodes logs, specify the path to the script-to-execute on each node by setting [parameter](#Parameters) `--linux-script` if the node distro is not `aks-ubuntu`. A sample script can be found [here](/scripts/collect-logs.sh).

To collect Windows nodes logs, specify the path to the script-to-execute on each node by setting [parameter](#Parameters) `--windows-script` if the node distro is not `aks-windows`. A sample script can be found [here](/scripts/collect-windows-logs.ps1).

If you choose to pass your own custom log collection script, make sure it zips all relevant files to file `"/tmp/logs.zip"` for Linux and `"%TEMP%\{NodeName}.zip"` for Windows. Needless to say, the custom script should only query for troubleshooting information and it should not change the cluster or node configuration.

### Upload logs to a Storage Account Container

Once the cluster logs were successfully retrieved, AKS Engine can persist them to an Azure Storage Account container if optional parameter `--storage-container-sas-url` is set. AKS Engine expects the container name to be part of the provided [SAS URL](https://docs.microsoft.com/azure/storage/common/storage-sas-overview). The expected format is `https://{blob-service-uri}/{container-name}?{sas-token}`.
Note: storage accounts on custom clouds using the `AD FS` identity provider are not yet supported

## Usage

Assuming that you have a cluster deployed and the API model originally used to deploy that cluster is stored at `_output/<dnsPrefix>/apimodel.json`, then you can collect logs running a command like:

```console
$ aks-engine get-logs \
    --location <location> \
    --api-model _output/<dnsPrefix>/apimodel.json \
    --ssh-host <dnsPrefix>.<location>.cloudapp.azure.com \
    --linux-ssh-private-key ~/.ssh/id_rsa \
    --linux-script scripts/collect-logs.sh \
    --windows-script scripts/collect-windows-logs.ps1 \
    --storage-container-sas-url <storageContainerSASUrl>
```

### Parameters

|Parameter|Required|Description|
|---|---|---|
|--location|yes|Azure location of the cluster's resource group.|
|--api-model|yes|Path to the generated API model for the cluster.|
|--ssh-host|yes|FQDN, or IP address, of an SSH listener that can reach all nodes in the cluster.|
|--linux-ssh-private-key|yes|Path to a SSH private key that can be use to create a remote session on the cluster Linux nodes.|
|--linux-script|no|Custom log collection bash script. It it required only when the Linux node distro is not `aks-ubuntu` and it should produce file `/tmp/logs.zip`.|
|--windows-script|no|Custom log collection powershell script. It is required only when the Windows node distro is not `aks-windows` and it should produce file `%TEMP%\{NodeName}.zip`.|
|--output-directory|no|Output directory, derived from `--api-model` if missing.|
|--control-plane-only|no|Only collect logs from master nodes.|
|--storage-container-sas-url|no|Upload collected logs to a storage container on Azure or custom cloud.|