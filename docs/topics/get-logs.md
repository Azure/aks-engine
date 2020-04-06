# Retrieving node and cluster logs

## Prerequisites

All the commands in this guide require both the Azure CLI and `aks-engine`. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing.

This guide assumes you already have deployed a cluster using `aks-engine`. For more details on how to do that see [deploy](../tutorials/deploy.md).

## Retrieving Logs

The `aks-engine get-logs` command can be useful to troubleshoot issues with your cluster. It will produce, collect and download to your workstation a set of files that include node configuration, cluster state and configuration, and provision log files.

At a high level, it works by stablishing a SSH session into each node, executing a `log collection` script that collects and zips relevant files, and downloading the zip file to your local computer.

### SSH Sessions

A valid SSH private key is always required to stablish a SSH session to the cluster Linux nodes. Windows credentials are stored in the apimodel and will be loaded from there.

### Log Collection Scripts

_TODO Specify starting AKSe version._ The [log collection script](/scripts/collect-logs.sh) will be available in Linux nodes that were provisioned using a flavor of the `aks-ubuntu-*` distro. If you chose a different distro for your linux pool or used an older version of AKS Engine for the initial deployment, then you can use your own script by setting [parameter](#Parameters) `--linux-script`.

The default distro for Windows node pools will also include the [log collection script](/scripts/collect-windows-logs.ps1). There is no support to pass your own custom script at this point.

## Usage

Assuming that you have a cluster deployed and the apimodel originally used to deploy that cluster is stored at `_output/<dnsPrefix>/apimodel.json`, then you can collect logs running a command like:

```console
$ aks-engine get-logs \
    --location <location> \
    --api-model _output/mycluster/apimodel.json \
    --apiserver mycluster.<location>.cloudapp.azure.com \
    --linux-ssh-private-key ~/.ssh/id_rsa
```

### Parameters

|Parameter|Required|Description|
|---|---|---|
|--location|yes|Azure location of the cluster's resource group.|
|--api-model|yes|Path to the generated api model for the cluster.|
|--apiserver|yes|Kubernetes apiserver endpoint.|
|--linux-ssh-private-key|yes|Path to a SSH private key that can be use to create a remote session on the cluster Linux nodes.|
|--linux-script|no|Custom log collection script. It should produce file `/tmp/logs.zip`.|
|--output-directory|no|Output directory, derived from `--api-model` if missing.|
|--control-plane-only|no|Only collect logs from master nodes.|
