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

### `aks-engine deploy`

The `aks-engine deploy` command will create a new cluster from scratch, using an API model (cluster definition) file as input to define the desired cluster configuration and shape, into the subscription, region, and resource group you provide, using credentials that you provide.

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

### `aks-engine scale`

The `aks-engine scale` command will scale (in or out) a specific node pool participating in a Kubernetes cluster created by AKS Engine.

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
