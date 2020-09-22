# AKS Engine CLI Overview

AKS Engine is designed to be used as a CLI tool (`aks-engine`). This document outlines the functionality that `aks-engine` provides to create and maintain a Kubernetes cluster on Azure.

## `aks-engine` commands

To get a quick overview of the commands available via the `aks-engine` CLI tool, just run `aks-engine` with no arguments (or include the `--help` argument):

```sh
$ aks-engine
Usage:
  aks-engine [flags]
  aks-engine [command]

Available Commands:
  addpool       Add a node pool to an existing AKS Engine-created Kubernetes cluster
  completion    Generates bash completion scripts
  deploy        Deploy an Azure Resource Manager template
  generate      Generate an Azure Resource Manager template
  get-logs      Collect logs and current cluster nodes configuration.
  get-versions  Display info about supported Kubernetes versions
  help          Help about any command
  rotate-certs  Rotate certificates on an existing AKS Engine-created Kubernetes cluster
  scale         Scale an existing AKS Engine-created Kubernetes cluster
  update        Update an existing AKS Engine-created VMSS node pool
  upgrade       Upgrade an existing AKS Engine-created Kubernetes cluster
  version       Print the version of aks-engine

Flags:
      --debug                enable verbose debug logs
  -h, --help                 help for aks-engine
      --show-default-model   Dump the default API model to stdout

Use "aks-engine [command] --help" for more information about a command.
```

## Operational Cluster Commands

These commands are provided by AKS Engine in order to create and maintain Kubernetes clusters. Note: there is no `aks-engine` command to delete a cluster; to delete a Kubernetes cluster created by AKS Engine, you must delete the resource group that contains cluster resources. If the resource group can't be deleted because it contains other, non-Kubernetes-relate Azure resources, then you must manually delete the Virtual Machine and/or Virtual Machine Scale Set (VMSS), Disk, Network Interface, Network Security Group, Public IP Address, Virtual Network, Load Balancer, and all other resources specified in the aks-engine-generated ARM template. Because manually deleting resources is tedious and requires following serial dependencies in the correct order, it is recommended that you dedicate a resource group for the Azure resources that AKS Engine will create to run your Kubernetes cluster. If you're running more than one cluster, we recommend a dedicated resource group per cluster.

### `aks-engine deploy`

The `aks-engine deploy` command will create a new cluster from scratch, using an API model (cluster definition) file as input to define the desired cluster configuration and shape, into the subscription, region, and resource group you provide, using credentials that you provide. Use this command to create a new cluster.

```sh
$ aks-engine deploy --help
Deploy an Azure Resource Manager template, parameters file and other assets for a cluster

Usage:
  aks-engine deploy [flags]

Flags:
  -m, --api-model string             path to your cluster definition file
      --auth-method client_secret    auth method (default:client_secret, `cli`, `client_certificate`, `device`) (default "client_secret")
      --auto-suffix                  automatically append a compressed timestamp to the dnsPrefix to ensure unique cluster name automatically
      --azure-env string             the target Azure cloud (default "AzurePublicCloud")
      --ca-certificate-path string   path to the CA certificate to use for Kubernetes PKI assets
      --ca-private-key-path string   path to the CA private key to use for Kubernetes PKI assets
      --certificate-path string      path to client certificate (used with --auth-method=client_certificate)
      --client-id string             client id (used with --auth-method=[client_secret|client_certificate])
      --client-secret string         client secret (used with --auth-method=client_secret)
  -p, --dns-prefix string            dns prefix (unique name for the cluster)
  -f, --force-overwrite              automatically overwrite existing files in the output directory
  -h, --help                         help for deploy
      --identity-system azure_ad     identity system (default:azure_ad, `adfs`) (default "azure_ad")
      --language string              language to return error messages in (default "en-us")
  -l, --location string              location to deploy to (required)
  -o, --output-directory string      output directory (derived from FQDN if absent)
      --private-key-path string      path to private key (used with --auth-method=client_certificate)
  -g, --resource-group string        resource group to deploy to (will use the DNS prefix from the apimodel if not specified)
      --set stringArray              set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)
  -s, --subscription-id string       azure subscription id (required)

Global Flags:
      --debug   enable verbose debug logs
```

Detailed documentation on `aks-engine deploy` can be found [here](../topics/creating_new_clusters.md#deploy).

### `aks-engine scale`

The `aks-engine scale` command will scale (in or out) a specific node pool participating in a Kubernetes cluster created by AKS Engine. Use this command to manually scale a node pool to a specific number of nodes.

```sh
$ aks-engine scale --help
Scale an existing AKS Engine-created Kubernetes cluster by specifying a new desired number of nodes in a node pool

Usage:
  aks-engine scale [flags]

Flags:
  -m, --api-model string            path to the generated apimodel.json file
      --apiserver string            apiserver endpoint (required to cordon and drain nodes)
      --auth-method client_secret   auth method (default:client_secret, `cli`, `client_certificate`, `device`) (default "client_secret")
      --azure-env string            the target Azure cloud (default "AzurePublicCloud")
      --certificate-path string     path to client certificate (used with --auth-method=client_certificate)
      --client-id string            client id (used with --auth-method=[client_secret|client_certificate])
      --client-secret string        client secret (used with --auth-method=client_secret)
  -h, --help                        help for scale
      --identity-system azure_ad    identity system (default:azure_ad, `adfs`) (default "azure_ad")
      --language string             language to return error messages in (default "en-us")
  -l, --location string             location the cluster is deployed in
  -c, --new-node-count int          desired number of nodes
      --node-pool string            node pool to scale
      --private-key-path string     path to private key (used with --auth-method=client_certificate)
  -g, --resource-group string       the resource group where the cluster is deployed
  -s, --subscription-id string      azure subscription id (required)

Global Flags:
      --debug   enable verbose debug logs
```

The `scale` command has limitations for scaling in (reducing the number of nodes in a node pool):

- It accepts a new, desired node count; it does not accept a list of specific nodes to remove from the pool.
- For VMSS-backed node pools, the removed nodes will not be cordoned and drained prior to being removed, which means any running workloads on nodes-to-be-removed will be disrupted without warning, and temporary operational impact is to be expected.

We generally recommend that you manage node pool scaling dynamically using the `cluster-autoscaler` project. More documentation about `cluster-autoscaler` is [here](../../examples/addons/cluster-autoscaler/README.md), including how to automatically install and configure it at cluster creation time as an AKS Engine addon.

Detailed documentation on `aks-engine scale` can be found [here](../topics/scale.md).

### `aks-engine update`

The `aks-engine update` command will update the VMSS model of a node pool according to a modified configuration of the aks-engine-generated `apimodel.json`. The updated node configuration will not take affect on any existing nodes, but will be applied to all future, new nodes created by VMSS scale out operations. Use this command to update the node configuration (OS configuration, VM SKU, Kubernetes kubelet configuration) of an existing node pool.

```sh
$ aks-engine update --help
Update an existing AKS Engine-created VMSS node pool in a Kubernetes cluster by updating its VMSS model

Usage:
  aks-engine update [flags]

Flags:
  -m, --api-model string            path to the generated apimodel.json file
      --auth-method client_secret   auth method (default:client_secret, `cli`, `client_certificate`, `device`) (default "client_secret")
      --azure-env string            the target Azure cloud (default "AzurePublicCloud")
      --certificate-path string     path to client certificate (used with --auth-method=client_certificate)
      --client-id string            client id (used with --auth-method=[client_secret|client_certificate])
      --client-secret string        client secret (used with --auth-method=client_secret)
  -h, --help                        help for update
      --identity-system azure_ad    identity system (default:azure_ad, `adfs`) (default "azure_ad")
      --language string             language to return error messages in (default "en-us")
  -l, --location string             location the cluster is deployed in
      --node-pool string            node pool to scale
      --private-key-path string     path to private key (used with --auth-method=client_certificate)
  -g, --resource-group string       the resource group where the cluster is deployed
  -s, --subscription-id string      azure subscription id (required)

Global Flags:
      --debug   enable verbose debug logs
```

Detailed documentation on `aks-engine update` can be found [here](../topics/update.md).

### `aks-engine addpool`

The `aks-engine addpool` command will add a new node pool to an existing AKS Engine-created cluster. Using a JSON file to define a brand new node pool, and referencing the aks-engine-generated `apimodel.json`, you can add new nodes to your cluster. Use this command to add a specific number of brand new nodes using a discrete configuration compared to existing nodes participating in your cluster.

```sh
$ aks-engine addpool --help
Add a node pool to an existing AKS Engine-created Kubernetes cluster by referencing a new agentpoolProfile spec

Usage:
  aks-engine addpool [flags]

Flags:
  -m, --api-model string            path to the generated apimodel.json file
      --auth-method client_secret   auth method (default:client_secret, `cli`, `client_certificate`, `device`) (default "client_secret")
      --azure-env string            the target Azure cloud (default "AzurePublicCloud")
      --certificate-path string     path to client certificate (used with --auth-method=client_certificate)
      --client-id string            client id (used with --auth-method=[client_secret|client_certificate])
      --client-secret string        client secret (used with --auth-method=client_secret)
  -h, --help                        help for addpool
      --identity-system azure_ad    identity system (default:azure_ad, `adfs`) (default "azure_ad")
      --language string             language to return error messages in (default "en-us")
  -l, --location string             location the cluster is deployed in
  -p, --node-pool string            path to a JSON file that defines the new node pool spec
      --private-key-path string     path to private key (used with --auth-method=client_certificate)
  -g, --resource-group string       the resource group where the cluster is deployed
  -s, --subscription-id string      azure subscription id (required)

Global Flags:
      --debug   enable verbose debug logs
```

Detailed documentation on `aks-engine addpool` can be found [here](../topics/addpool.md).

### `aks-engine upgrade`

The `aks-engine upgrade` command orchestrates a Kubernetes version upgrade across your existing cluster nodes. Use this command to upgrade the Kubernetes version running your control plane, and optionally on all your nodes as well.

```sh
$ aks-engine upgrade --help
Upgrade an existing AKS Engine-created Kubernetes cluster, one node at a time

Usage:
  aks-engine upgrade [flags]

Flags:
  -m, --api-model string            path to the generated apimodel.json file
      --auth-method client_secret   auth method (default:client_secret, `cli`, `client_certificate`, `device`) (default "client_secret")
      --azure-env string            the target Azure cloud (default "AzurePublicCloud")
      --certificate-path string     path to client certificate (used with --auth-method=client_certificate)
      --client-id string            client id (used with --auth-method=[client_secret|client_certificate])
      --client-secret string        client secret (used with --auth-method=client_secret)
      --control-plane-only          upgrade control plane VMs only, do not upgrade node pools
      --cordon-drain-timeout int    how long to wait for each vm to be cordoned in minutes (default -1)
  -f, --force                       force upgrading the cluster to desired version. Allows same version upgrades and downgrades.
  -h, --help                        help for upgrade
      --identity-system azure_ad    identity system (default:azure_ad, `adfs`) (default "azure_ad")
  -b, --kubeconfig string           the path of the kubeconfig file
      --language string             language to return error messages in (default "en-us")
  -l, --location string             location the cluster is deployed in (required)
      --private-key-path string     path to private key (used with --auth-method=client_certificate)
  -g, --resource-group string       the resource group where the cluster is deployed (required)
  -s, --subscription-id string      azure subscription id (required)
  -k, --upgrade-version string      desired kubernetes version (required)
      --upgrade-windows-vhd         upgrade image reference of the Windows nodes (default true)
      --vm-timeout int              how long to wait for each vm to be upgraded in minutes (default -1)

Global Flags:
      --debug   enable verbose debug logs
```

Detailed documentation on `aks-engine upgrade` can be found [here](../topics/upgrade.md).

## Generate an ARM Template

AKS Engine also provides a command to generate a reusable ARM template only, without creating any actual Azure resources.

### `aks-engine generate`

The `aks-engine generate` command is similar to `aks-engine deploy`: it uses an API model (cluster definition) file as input to define the desired cluster configuration and shape of a new Kubernetes cluster. Unlike `deploy`, `aks-engine generate` does not actually submit any operational requests to Azure, but is instead used to generate a reusable ARM template which may be deployed at a later time. Use this command as a part of a workflow that creates one or more Kubernetes clusters via an ARM group deployment that takes an ARM template as input (e.g., `az group deployment create` using the standard `az` Azure CLI).

```sh
$ aks-engine generate --help
Generates an Azure Resource Manager template, parameters file and other assets for a cluster

Usage:
  aks-engine generate [flags]

Flags:
  -m, --api-model string             path to your cluster definition file
      --ca-certificate-path string   path to the CA certificate to use for Kubernetes PKI assets
      --ca-private-key-path string   path to the CA private key to use for Kubernetes PKI assets
      --client-id string             client id
      --client-secret string         client secret
  -h, --help                         help for generate
      --no-pretty-print              skip pretty printing the output
  -o, --output-directory string      output directory (derived from FQDN if absent)
      --parameters-only              only output parameters files
      --set stringArray              set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)

Global Flags:
      --debug   enable verbose debug logs
```

Detailed documentation on `aks-engine generate` can be found [here](../topics/creating_new_clusters.md#generate).
