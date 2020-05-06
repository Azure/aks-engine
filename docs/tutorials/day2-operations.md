# Day 2 Operations

NOTE: The steps listed here were contributed by a community member, and is not officially supported. Please use at your own risk.

The steps listed on this page describe a way to modify a running Kubernetes cluster deployed with `aks-engine` on Azure. These steps are only tested with changes targeting actually Azure resources. Changes made to Kubernetes configuration are not tested yet.

## `generate` and `deploy`

These are the common steps (unless described otherwise) you'll have to run after modifying an existing `apimodel.json` file.

* Modify the apimodel.json file located in the `_output/<clustername>` folder
* Run `aks-engine generate --api-model _output/<clustername>/apimodel.json`. This will update the `azuredeploy*` files needed for the new ARM deployment. These files are also located in the `_output` folder.
* Apply the changes by manually starting an ARM deployment. From within the  `_output/<clustername>` run

        az deployment group create --template-file azuredeploy.json --parameters azuredeploy.parameters.json --resource-group "<my-resource-group>"

  To use the `az` CLI tools you have to login. More info can be found here: https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli?view=azure-cli-latest

  _Note: I use `az deployment group` instead of `aks-engine deploy` because the latter seems to assume you are deploying a new cluster and as a result overwriting your private ssh keys located in the _ouput folder_

* Grab a coffee
* Profit!

## Common scenarios (tested)

### Adding a node pool

Add (or copy) an entry in the `agentPoolProfiles` array.

### Removing a node pool

* Delete the related entry from `agentPoolProfiles` section in the `_output/<clustername>/api-model.json` file
* [Drain](https://kubernetes.io/docs/tasks/administer-cluster/safely-drain-node/) nodes from inside Kubernetes
* `generate` and `deploy` (see above)
* Delete VM's and related resources (disk, NIC, availability set) from Azure portal
* Remove the pool from the original `apimodel.json` file

### Resizing a node pool

Use the `aks-engine scale` command

    aks-engine scale   --location westeurope --subscription-id "xxx" --resource-group "<my-resource-group" \
    --api-model ./somedir/apimodel.json --node-pool <nodepool name> --new-node-count <desired number of nodes> --apiserver <apiserver endpoint FQDN or IP address>

**Remember to also update your original api-model.json file (used for 1st deployment) or else you would end up with the original number of VM's after using the `generate` command described above**

### Resize VM's in existing agent pool

* Modify the `vmSize` in the  `agentPoolProfiles` section
* `generate` and `deploy` (see above)

**Important: The default ARM deployment won't drain your Kubernetes nodes properly before 'rebooting' them. Please [drain](https://kubernetes.io/docs/tasks/administer-cluster/safely-drain-node/) them manually before deploying the change**
