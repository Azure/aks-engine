# Retrieving Node and Cluster Logs

## Prerequisites

All the commands in this guide require both the Azure CLI and `aks-engine`. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing.

This guide assumes you already have deployed a cluster using `aks-engine`. For more details on how to do that see [deploy](../tutorials/deploy.md).

## Retrieving Logs

The `aks-engine get-logs` command can be useful to troubleshoot issues with your cluster. It will produce, collect and download to your workstation a set of files that include node configuration, cluster state and configuration, and provision log files.

At a high level, it works by establishing a SSH session into each node, executing a [log collection](#log-collection-scripts) script that collects and zips relevant files, and downloading the zip file to your local computer.

### SSH Authentication

A valid SSH private key is always required to stablish a SSH session to the cluster Linux nodes. Windows credentials are stored in the API model and will be loaded from there. Make sure `windowsprofile.sshEnabled` is set to `true` to enable SSH in your Windows nodes.

### Log Collection Scripts

To collect Linux nodes logs, specify the path to the script-to-execute on each node by setting [parameter](#Parameters) `--linux-script`. A sample script can be found [here](/scripts/collect-logs.sh).

If you choose to pass your own custom log collection script, make sure it zips all relevant files to file `/tmp/logs.zip`. Needless to say, the custom script should only query for troubleshooting information and it should not change the cluster or node configuration.

The default OS distro for Windows node pools already includes a [log collection script](./scripts/collect-windows-logs.ps1). There is no support to pass your own custom script at this point.

## Usage

Assuming that you have a cluster deployed and the API model originally used to deploy that cluster is stored at `_output/<dnsPrefix>/apimodel.json`, then you can collect logs running a command like:

```console
$ aks-engine get-logs \
    --location <location> \
    --api-model _output/<dnsPrefix>/apimodel.json \
    --ssh-host <dnsPrefix>.<location>.cloudapp.azure.com \
    --linux-ssh-private-key ~/.ssh/id_rsa \
    --linux-script scripts/collect-logs.sh
```

### Parameters

|Parameter|Required|Description|
|---|---|---|
|--location|yes|Azure location of the cluster's resource group.|
|--api-model|yes|Path to the generated API model for the cluster.|
|--ssh-host|yes|FQDN, or IP address, of an SSH listener that can reach all nodes in the cluster.|
|--linux-ssh-private-key|yes|Path to a SSH private key that can be use to create a remote session on the cluster Linux nodes.|
|--linux-script|yes|Custom log collection script. It should produce file `/tmp/logs.zip`.|
|--output-directory|no|Output directory, derived from `--api-model` if missing.|
|--control-plane-only|no|Only collect logs from master nodes.|
