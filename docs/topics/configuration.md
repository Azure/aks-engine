# Configuring the CLI

There is a growing list of configuration options for AKS Engine. Many of these options can be passed to the CLI as feature flags such as the log verbosity level (`--debug`). Some options, such as the authentication method and its corresponding credentials, can also be read from environment variables (e.g. `AKS_ENGINE_AUTH_METHOD` and `AKS_ENGINE_SUBSCRIPTION_ID`). It is also possible to manage these options in a central place through the use of config files.

The order of operations for configuration is

1. Feature flags set directly on the command line
1. Environment variables
1. Settings provided in settings.json

When `aks-engine` is invoked, the CLI takes input from all three sources and merges it all together into one final configuration. That is, if you set both `$AKS_ENGINE_AUTH_METHOD` and `--auth-method` to different values, AKS Engine will take the value from `--auth-method` as its configuration.

## Using Environment Variables

Some of AKS Engine's global configuration is available through the use of environment variables. The following environment variables correspond to their feature flags:

+---------------------------------------+------------------------------+
| Environment Variable Name             | AKS Engine Feature Flag Name |
+---------------------------------------+------------------------------+
| AKS_ENGINE_AZURE_ENVIRONMENT          | --azure-env                  |
| AKS_ENGINE_AUTH_METHOD                | --auth-method                |
| AKS_ENGINE_SUBSCRIPTION_ID            | --subscription-id            |
| AKS_ENGINE_CLIENT_ID                  | --client-id                  |
| AKS_ENGINE_CLIENT_SECRET              | --client-secret              |
| AKS_ENGINE_CERTIFICATE_PATH           | --certificate-path           |
| AKS_ENGINE_PRIVATE_KEY_PATH           | --private-key-path           |
| AKS_ENGINE_PRIVATE_LANGUAGE           | --language                   |
| AKS_ENGINE_PRIVATE_DEBUG              | --debug                      |
| AKS_ENGINE_PRIVATE_PROFILE            | --profile                    |
| AKS_ENGINE_PRIVATE_SHOW_DEFAULT_MODEL | --show-default-model         |
+---------------------------------------+------------------------------+

## Using settings.json

An AKS Engine config file is a JSON configuration file that specifies values for options that control the behavior of the AKS Engine CLI.

In Unix/Linux systems, AKS Engine looks for configuration under `~/.config/aks-engine/settings.json`. On Windows, AKS Engine looks under `%LOCALAPPDATA%\aks-engine\settings.json`.

`aks-engine configure` will prompt you for the most basic required configuration values such as your auth method, enabling debug logging etc. You can configure a named profile using the `--profile` flag, for example `aks-engine configure --profile development`. If your config file does not exist, the AKS Engine CLI will create it for you. To keep an existing value, hit enter when prompted for the value. When you are prompted for information, the current value will be displayed in `(parentheses)`. If the config item has no value, no parentheses will be displayed.

An example configuration file generated with `aks-engine configure` might look like this:

```json
{
    "auth": {
        "auth-method": "cli",
        "subscription-id": "abc123",
        "azure-env": "AzurePublicCloud"
    },
    "aks-engine": {
        "debug": true,
    }
}
```

### Sections

The following sections and options are currently recognized within the config file.

#### profiles

If you have multiple Azure credentials that you use for different purposes, use the profile style. You can set an arbitrary number of profiles within your configuration files and then reference them by name when you instantiate your connection. If you specify a profile that does not exist in the configuration during `aks-engine configure`, it will be created for you.

Using the same example above, generating profile-specific configuration with `aks-engine configure --profile production` after calling `aks-engine configure` might look like this:

```json
{
    "profiles": {
        "production": {
            "auth": {
                "auth-method": "client_secret",
                "client-id": "abc123",
                "client-secret": "abc123"
            },
            "aks-engine": {
                "debug": false,
                "deploy": {
                    "location": "eastus",
                    "set": ["key1=val1", "key2=val2"]
                }
            }
        }
    },
    "auth": {
        "auth-method": "cli",
        "subscription-id": "abc123",
        "azure-env": "AzurePublicCloud"
    },
    "aks-engine": {
        "debug": true
    }
}
```

Which you can then reference in the CLI using `aks-engine --profile production deploy <...>`.

#### auth

The auth section is used to specify the Azure credentials used for all AKS Engine requests.

The order of precedence for authentication credentials is:

- credentials specified by feature flags on the CLI (`--auth-method`)
- credentials specified by environment variables (`AKS_ENGINE_AUTH_METHOD`)
- credentials specified as named profiles in the config file
- credentials specified by default in the config file

If you have multiple Azure credentials that you use for different purposes, use the profile style shown above. You can set an arbitrary number of profiles within your configuration files and then reference them by name when you instantiate your connection.

#### aks-engine

The aks-engine section is used to specify options that control the operation of the AKS Engine CLI itself. This section defines the following options:

+---------------------+----------+-----------------------------------------------------------------------------------------+
| Parameter Name      | Type     | Description                                                                             |
+---------------------+----------+-----------------------------------------------------------------------------------------+
| debug               | boolean  | Enable verbose debug logs                                                               |
| show-default-model  | boolean  | Dump the default API model to stdout                                                    |
| deploy              | object   | Parameters for `aks-engine deploy`. See below                                           |
| generate            | object   | Parameters for `aks-engine generate`. See below                                         |
| scale               | object   | Parameters for `aks-engine scale`. See below                                            |
| upgrade             | object   | Parameters for `aks-engine upgrade`. See below                                          |
| version             | object   | Parameters for `aks-engine version`. See below                                          |
+---------------------+----------+-----------------------------------------------------------------------------------------+

Parameters for the `deploy` object:

+---------------------+----------+-----------------------------------------------------------------------------------------+
| Parameter Name      | Type     | Description                                                                             |
+---------------------+----------+-----------------------------------------------------------------------------------------+
| api-model           | string   | Path to the apimodel file                                                               |
| auto-suffix         | boolean  | Append a compressed timestamp to the dns prefix                                         |
| ca-certificate-path | string   | Path to the CA certificate to use for Kubernetes PKI assets                             |
| ca-private-key-path | string   | Path to the CA private key to use for Kubernetes PKI assets                             |
| dns-prefix          | string   | DNS prefix (unique name for the cluster)                                                |
| force-overwrite     | boolean  | Automatically overwrite existing files in the output directory                          |
| location            | string   | Azure region to deploy to                                                               |
| output-directory    | string   | Output directory for all generated assets                                               |
| resource-group      | string   | Resource group to deploy to. Will use the DNS prefix from the apimodel if not specified |
| set                 | string[] | Set values on the command line                                                          |
+---------------------+----------+-----------------------------------------------------------------------------------------+

Parameters for the `generate` object:

+---------------------+----------+-----------------------------------------------------------------------------------------+
| Parameter Name      | Type     | Description                                                                             |
+---------------------+----------+-----------------------------------------------------------------------------------------+
| api-model           | string   | Path to the apimodel file                                                               |
| ca-certificate-path | string   | Path to the CA certificate to use for Kubernetes PKI assets                             |
| ca-private-key-path | string   | Path to the CA private key to use for Kubernetes PKI assets                             |
| pretty-print        | boolean  | Pretty print the output                                                                 |
| output-directory    | string   | Output directory for all generated assets                                               |
| parameters-only     | boolean  | Only output parameters files                                                            |
| set                 | string[] | Set values on the command line                                                          |
+---------------------+----------+-----------------------------------------------------------------------------------------+

Parameters for the `scale` object:

+---------------------+----------+-----------------------------------------------------------------------------------------+
| Parameter Name      | Type     | Description                                                                             |
+---------------------+----------+-----------------------------------------------------------------------------------------+
| deployment-dir      | string   | the location of the output from generate                                                |
| location            | string   | Azure region where the cluster is deployed                                              |
| master-fqdn         | string   | FQDN for the master load balancer. Needed to scale down Kubernetes agent pools          |
| new-node-count      | int      | The desired number of nodes                                                             |
| node-pool           | string   | The node pool to scale                                                                  |
| resource-group      | string   | The resource group where the cluster is deployed                                        |
+---------------------+----------+-----------------------------------------------------------------------------------------+

Parameters for the `upgrade` object:

+---------------------+----------+-----------------------------------------------------------------------------------------+
| Parameter Name      | Type     | Description                                                                             |
+---------------------+----------+-----------------------------------------------------------------------------------------+
| deployment-dir      | string   | the location of the output from generate                                                |
| location            | string   | Azure region where the cluster is deployed                                              |
| resource-group      | string   | The resource group where the cluster is deployed                                        |
| upgrade-version     | string   | Desired kubernetes version                                                              |
| vm-timeout          | int      | How long to wait for each VM to be upgraded in minutes                                  |
+---------------------+----------+-----------------------------------------------------------------------------------------+

Parameters for the `version` object:

+---------------------+----------+-----------------------------------------------------------------------------------------+
| Parameter Name      | Type     | Description                                                                             |
+---------------------+----------+-----------------------------------------------------------------------------------------+
| output              | string   | Output format to use                                                                    |
+---------------------+----------+-----------------------------------------------------------------------------------------+
