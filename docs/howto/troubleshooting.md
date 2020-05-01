## Troubleshooting

Common issues or questions that users have run into when using AKS Engine are detailed below.

## VMExtensionProvisioningError or VMExtensionProvisioningTimeout

The two above VMExtensionProvisioning— errors tell us that a vm in the cluster failed installing required application prerequisites after CRP provisioned the VM into the resource group. When `aks-engine deploy` creates a new Kubernetes cluster, a series of shell scripts runs to install prereq's like docker, etcd, Kubernetes runtime, and various other host OS packages that support the Kubernetes application layer. *Usually* this indicates one of the following:

1. Something about the cluster configuration is pathological. For example, perhaps the cluster config includes a custom version of a particular software dependency that doesn't exist. Or, another example, for a cluster created inside a custom VNET (i.e., a user-provided, pre-existing VNET), perhaps that custom VNET does not have general outbound internet access, and so apt, docker pull, etc is not able to execute successfully.
2. A transient Azure environmental error caused the shell script operation to timeout, or exceed its retry count. For example, the shell script may attempt to download a required package (e.g., etcd), and if the Azure networking environment for the newly provisioned vm is flaky for a period of time, then the shell script may retry several times, but eventually timeout and fail.

For classification #1 above, the appropriate strategic response is to figure out what about the cluster configuration is incorrect, and to fix it. We expect such scenarios to always fail in the above way: cluster deployments will not be successful until the cluster configuration is made to be correct.

For classification #2 above, the appropriate strategic response is to retry a few times. If a 2nd or 3rd attempt succeeds, it is a hint that a transient environmental condition is the cause of the initial failure.

### What is CSE?

CSE stands for CustomScriptExtension, and is just a way of expressing: "a script that executes as part of the VM provisioning process, and that must exit 0 (i.e., successfully) in order for that VM provisioning process to succeed". Basically it's another way of expressing the VMExtensionProvisioning— concept above.

To summarize, the way that AKS Engine implements Kubernetes on Azure is a collection of (1) Azure VM configuration + (2) shell script execution. Both are implemented as a single operational unit, and when #2 fails, we consider the entire VM provisioning operation to be a failure; more importantly, if only one VM in the cluster deployment fails, we consider the entire cluster operation to be a failure.

### How To Debug CSE errors (Linux)

In order to troubleshoot a cluster that failed in the above way(s), we need to grab the CSE logs from the host VM itself.

From a vm node that did not provision successfully:

- grab the entire file at `/var/log/azure/cluster-provision.log`

- grab the entire file at `/var/log/cloud-init-output.log`

How to determine the above?

1. Look at the deployment error message. The error should include which VM extension failed the deployment. For example, `cse-master-0` means that the CSE extension of VM master 0 failed. It should also include the date that the extension started running at and the name of the VM it was running on.

2. From a master node: `kubectl get nodes`

- Are there any missing master or agent nodes?
  - if so, that node vm probably failed CSE: grab the log files above from that vm
- Are there no working nodes?
  - if so, grab the log files above from the master vm you are on

#### CSE Exit Codes

```
"code": "VMExtensionProvisioningError"
"message": "VM has reported a failure when processing extension 'cse1'. Error message: "Enable failed: failed to
execute command: command terminated with exit status=20\n[stdout]\n\n[stderr]\n"."
```

Look for the exit code. In the above example, the exit code is `20`. The list of exit codes and their meaning can be found [here](../../pkg/engine/cse.go).

If after following the above you are still unable to troubleshoot your deployment error, please open a Github issue with title "CSE error: exit code <INSERT_YOUR_EXIT_CODE>" and include the following in the description:

1. Relevant data from the cluster definition JSON file (API model) used to deploy the cluster. **Please make sure you remove all secrets and keys before posting it on GitHub.**

2. The output of `kubectl get nodes`

3. The content of `/var/log/azure/cluster-provision.log` and `/var/log/cloud-init-output.log`


### How To Debug CSE Errors (Windows)

There are two symptoms where you may need to debug Custom Script Extension errors on Windows:

- VMExtensionProvisioningError or VMExtensionProvisioningTimeout
- `kubectl node` doesn't list the Windows node(s)

To get more logs, you need to connect to the Windows nodes using Remote Desktop - see [Connecting to Windows Nodes](#connecting-to-windows-nodes)

Once connected, check the following logs for errors:

 - `c:\Azure\CustomDataSetupScript.log`

#### Connecting to Windows nodes

Since the nodes are on a private IP range, you will need to use SSH local port forwarding from a master node to the Windows node to use remote.



1. Get the IP of the Windows node with `az vm list` and `az vm show`

    ```
    $ az vm list --resource-group group1 -o table
    Name                      ResourceGroup    Location
    ------------------------  ---------------  ----------
    29442k8s9000              group1           westus2
    29442k8s9001              group1           westus2
    k8s-linuxpool-29442807-0  group1           westus2
    k8s-linuxpool-29442807-1  group1           westus2
    k8s-master-29442807-0     group1           westus2

    $ az vm show -g group1 -n 29442k8s9000 --show-details --query 'privateIps'
    "10.240.0.4"
    ```

2. Forward a local port to the Windows port 3389, such as `ssh -L 5500:10.240.0.4:3389 <masternode>.<region>.cloudapp.azure.com`
3. Run `mstsc.exe /v:localhost:5500`

Now, you can use the default CMD window or install other tools as needed with the GUI. If you would like to enable PowerShell remoting, continue on to step 4.

4. Ansible uses PowerShell remoting over HTTPS, and has a convenient script to enable it. Run `PowerShell` on the Windows node, then these two steps to enable remoting.

```
Start-BitsTransfer https://raw.githubusercontent.com/ansible/ansible/devel/examples/scripts/ConfigureRemotingForAnsible.ps1
.\ConfigureRemotingForAnsible.ps1
```

5. Now, you're ready to connect from the Linux master to the Windows node:

```
$ docker run -it mcr.microsoft.com/powershell
PowerShell v6.0.2
Copyright (c) Microsoft Corporation. All rights reserved.

https://aka.ms/pscore6-docs
Type 'help' to get help.

PS /> $cred = Get-Credential

PowerShell credential request
Enter your credentials.
User: azureuser
Password for user azureuser: ************

PS /> Enter-PSSession 20143k8s9000 -Credential $cred -Authentication Basic -UseSSL
[20143k8s9000]: PS C:\Users\azureuser\Documents>
```

## Windows kubelet & CNI errors

If the node is not showing up in `kubectl get node` or fails to schedule pods, check for failures from the kubelet and CNI logs.

Follow the same steps [above](#how-to-debug-cse-errors-windows) to connect to Remote Desktop to the node, then look for errors in these logs:

 - `c:\k\kubelet.log`
 - `c:\k\kubelet.err.log`
 - `c:\k\azure-vnet*.log`



## Misconfigured Service Principal

If your Service Principal is misconfigured, none of the Kubernetes components will come up in a healthy manner.
You can check to see if this the problem:

```shell
ssh -i ~/.ssh/id_rsa USER@MASTERFQDN sudo journalctl -u kubelet | grep --text autorest
```

If you see output that looks like the following, then you have **not** configured the Service Principal correctly.
You may need to check to ensure the credentials were provided accurately, and that the configured Service Principal has
read and **write** permissions to the target Subscription.

`Nov 10 16:35:22 k8s-master-43D6F832-0 docker[3177]: E1110 16:35:22.840688    3201 kubelet_node_status.go:69] Unable to construct api.Node object for kubelet: failed to get external ID from cloud provider: autorest#WithErrorUnlessStatusCode: POST https://login.microsoftonline.com/72f988bf-86f1-41af-91ab-2d7cd011db47/oauth2/token?api-version=1.0 failed with 400 Bad Request: StatusCode=400`

[This documentation](../topics/service-principals.md) explains how to create/configure a service principal for an AKS Engine-created Kubernetes cluster.

## Failed upgrade

Please review the [upgrade documentation](../topics/upgrade.md) for a guide on upgrading AKS Engine-created Kubernetes clusters.
