# Scaling Kubernetes Clusters

## Prerequisites

All documentation in these guides assumes you have already downloaded both the Azure `az` CLI tool and the `aks-engine` binary tool. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing if you're creating a Kubernetes cluster using AKS Engine for the first time.

This guide assumes you already have a running cluster deployed using the `aks-engine` CLI. For more details on how to do that see [deploy](create_new_clusters.md#deploy) or [generate](generate.md).

## Scale

The `aks-engine scale` command can increase or decrease the number of nodes in an existing agent pool in an AKS Engine-created Kubernetes cluster. The command takes a desired node count, which means that you don't have any control over the naming of any new nodes, if the desired count is greater than the current number of nodes in the target pool (though generally new nodes are named incrementally from the "last" node); and you don't have any control over which nodes will be removed, if the desired node count is less than the current number of nodes in the target pool. For clusters that are relatively "static", using `aks-engine scale` may be appropriate. For highly dynamic clusters that want to take advantage of real-time, load-derived scaling, we recommend running `cluster-autoscaler` in your cluster, which we document [here](../../examples/addons/cluster-autoscaler/README.md).

Also note that for VMSS-backed node pools (the current AKS Engine default), scale "in" operations will *not* cordon and drain nodes before they are removed. This is because for VMSS node pools `aks-engine scale` is simply a thin wrapper around the VMSS API, and the VMSS API doesn't have any awareness of the Kubernetes application layer in order to cordon an drain nodes prior to removing instances from the VMSS. For this reason, again, we recommend using `cluster-autoscaler` with VMSS node pools for clusters with regular, period scaling requirements in both directions (both "in" and "out").

The example below will assume you have a cluster deployed, and that the API model originally used to deploy that cluster is stored at `_output/<dnsPrefix>/apimodel.json`. It will also assume that there is a node pool named "agentpool1" in your cluster.

To scale the cluster you will run a command like:

```sh
$ aks-engine scale --subscription-id <subscription_id> \
    --resource-group mycluster --location <location> \
    --client-id '<service principal client ID>' \
    --client-secret '<service principal client secret>' \
    --api-model _output/mycluster/apimodel.json --new-node-count <desired node count> \
    --node-pool agentpool1 --apiserver mycluster.<location>.cloudapp.azure.com
```

This command will re-use the `apimodel.json` file inside the output directory as input for a new ARM template deployment that will execute the scaling operation against the desired agent pool. When the scaling operation is done it will update the cluster definition in that same `apimodel.json` file to reflect the new node count and thus the updated, current cluster configuration.

### Parameters

|Parameter|Required|Description|
|-----------------|---|---|
|--subscription-id|yes|The subscription id the cluster is deployed in.|
|--resource-group|yes|The resource group the cluster is deployed in.|
|--location|yes|The location the resource group is in.|
|--api-model|yes|Relative path to the generated API model for the cluster.|
|--client-id|depends| The Service Principal Client ID. This is required if the auth-method is set to service_princpal/client_certificate|
|--client-secret|depends| The Service Principal Client secret. This is required if the auth-method is set to service_princpal|
|--certificate-path|depends| The path to the file which contains the client certificate. This is required if the auth-method is set to client_certificate|
|--node-pool|depends|Required if there is more than one node pool. Which node pool should be scaled.|
|--new-node-count|yes|Desired number of nodes in the node pool.|
|--apiserver|when scaling down|apiserver endpoint (required to cordon and drain nodes). This should be output as part of the create template or it can be found by looking at the public ip addresses in the resource group.|
|--auth-method|no|The authentication method used. Default value is `client_secret`. Other supported values are: `cli`, `client_certificate`, and `device`.|
|--language|no|Language to return error message in. Default value is "en-us").|

## Frequently Asked Questions

### Is it possible to scale control plane VMs?

It is not possible to increase or decrease *the number* of VMs that run the control plane. However, you may increase or decrease the *size* of the VM by modifying the `"vmSize"` property of the `masterProfile` in your cluster API model, and then run `aks-engine upgrade --control-plane-only`. See [the upgrade documentation](upgrade.md) for more information.

### What version of aks-engine should I use to run `aks-engine scale` operations?

As a general rule, we recommend that the latest released version of AKS Engine be used to scale out node pools. This is because the latest released version will have recent security updates and bug fixes to the OS layer, as well as critical system components like the container runtime. This may yield a heterogeneous node pool, but those differences should not introduce functional regressions; rather, they will ensure that a higher proportion of nodes in that pool are running the latest, validated bits. For example, here's an overview of a cluster originally built with 2 nodes in the pool "agentpool1" from `aks-engine` version `v0.52.1`, and then scaled out to 10 nodes using `aks-engine` v0.56.0:

```
$ kubectl get nodes -o wide
NAME                                 STATUS   ROLES    AGE     VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
k8s-agentpool1-10367588-vmss000000   Ready    agent    8m23s   v1.18.3   10.240.0.34    <none>        Ubuntu 18.04.4 LTS   5.3.0-1022-azure   docker://3.0.12+azure
k8s-agentpool1-10367588-vmss000001   Ready    agent    8m23s   v1.18.3   10.240.0.65    <none>        Ubuntu 18.04.4 LTS   5.3.0-1022-azure   docker://3.0.12+azure
k8s-agentpool1-10367588-vmss000002   Ready    agent    2m15s   v1.18.3   10.240.0.96    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000003   Ready    agent    2m38s   v1.18.3   10.240.0.127   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000004   Ready    agent    2m50s   v1.18.3   10.240.0.158   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000005   Ready    agent    3m38s   v1.18.3   10.240.0.189   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000006   Ready    agent    3m34s   v1.18.3   10.240.0.220   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000007   Ready    agent    3m32s   v1.18.3   10.240.0.251   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000008   Ready    agent    3m20s   v1.18.3   10.240.1.26    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000009   Ready    agent    3m33s   v1.18.3   10.240.1.57    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-master-10367588-0                Ready    master   8m23s   v1.18.3   10.255.255.5   <none>        Ubuntu 18.04.4 LTS   5.3.0-1022-azure   docker://3.0.12+azure
```

As you can see, there are 2 nodes (the original 2 nodes) running a moby build identified as `docker://3.0.12+azure`, while 8 nodes (the recently added nodes) are running a moby build identified as `docker://19.3.12`. Both of those builds of moby are functionally equivalent in terms of being able to service a Kubernetes v1.18.3 node, but the latter is preferable as it will include more recent fixes (some critical) to the container runtime implementation. It is preferable to have 8 of 10 nodes running the latest bits, compared to all 10 running the older bits, despite the potential negative trade off of the loss of strict homogeneity across the nodes in the pool.

The above scale operation using a newer version of the `aks-engine` CLI also has the side effect of updating the VMSS model that underlies that node pool, which means that any future scale out operation using the VMSS API (via the `az` CLI or Azure portal web UI) will yield nodes running the latest bits.

### How do I remove nodes from my VMSS node pool without incurring production downtime?

As stated above, when scaling "in" nodes running `aks-engine scale` against a VMSS-backed node pool, the deleted nodes will not be cordon + drained prior to being deleted, which means any running workloads will be interrupted non-gracefully. For this reason, when manually scaling in, we recommend that you *not* use `aks-engine scale`, but instead manually re-balance your cluster by moving workloads off of the number of nodes you desire to remove, and then manually delete those VMSS instances.

We'll use the example cluster above and remove the original 2 nodes running the older build of moby. First, we mark those nodes as unschedulable so that no new workloads are scheduled onto them during this maintenance:

```sh
$ for node in "k8s-agentpool1-10367588-vmss000000 k8s-agentpool1-10367588-vmss000001"; do kubectl cordon $node; done
node/k8s-agentpool1-10367588-vmss000000 cordoned
node/k8s-agentpool1-10367588-vmss000001 cordoned
```

We can then instruct the Kubernetes control plane to, as gracefully as possible, move workloads off of those nodes:

```sh
$ for node in "k8s-agentpool1-10367588-vmss000000 k8s-agentpool1-10367588-vmss000001"; do kubectl drain $node; done
node/k8s-agentpool1-10367588-vmss000000 already cordoned
node/k8s-agentpool1-10367588-vmss000001 already cordoned
error: unable to drain node "k8s-agentpool1-10367588-vmss000000", aborting command...

There are pending nodes to be drained:
 k8s-agentpool1-10367588-vmss000000
 k8s-agentpool1-10367588-vmss000001
error: cannot delete DaemonSet-managed Pods (use --ignore-daemonsets to ignore): kube-system/azure-cni-networkmonitor-wvrg7, kube-system/azure-ip-masq-agent-qqlvf, kube-system/blobfuse-flexvol-installer-9q45x, kube-system/csi-secrets-store-provider-azure-jsgkh, kube-system/csi-secrets-store-q5wnw, kube-system/kube-proxy-cgh7g
```

It's always best to do a vanilla `kubectl drain` first to see the set of scheduled pods that require a little more forceful removal, so that you can be extra sure that you actually want to do this. In our case, we're O.K. with removing those daemonsets, so we proceed to add the `--ignore-daemonsets` option:

```sh
$ for node in "k8s-agentpool1-10367588-vmss000000 k8s-agentpool1-10367588-vmss000001"; do kubectl drain $node --ignore-daemonsets; done
node/k8s-agentpool1-10367588-vmss000000 already cordoned
node/k8s-agentpool1-10367588-vmss000001 already cordoned
WARNING: ignoring DaemonSet-managed Pods: kube-system/azure-cni-networkmonitor-wvrg7, kube-system/azure-ip-masq-agent-qqlvf, kube-system/blobfuse-flexvol-installer-9q45x, kube-system/csi-secrets-store-provider-azure-jsgkh, kube-system/csi-secrets-store-q5wnw, kube-system/kube-proxy-cgh7g
evicting pod "metrics-server-bb7db87bc-xzxld"
pod/metrics-server-bb7db87bc-xzxld evicted
node/k8s-agentpool1-10367588-vmss000000 evicted
WARNING: ignoring DaemonSet-managed Pods: kube-system/azure-cni-networkmonitor-cvfqs, kube-system/azure-ip-masq-agent-p755d, kube-system/blobfuse-flexvol-installer-stc2x, kube-system/csi-secrets-store-fs9xr, kube-system/csi-secrets-store-provider-azure-7qhqt, kube-system/kube-proxy-bpdvl
evicting pod "coredns-autoscaler-5c7db64899-kp64h"
pod/coredns-autoscaler-5c7db64899-kp64h evicted
node/k8s-agentpool1-10367588-vmss000001 evicted
```

Now, delete the two VMSS instances:

```sh
$ az vmss delete-instances -g kubernetes-westus2-95121 -n k8s-agentpool1-10367588-vmss --instance-ids 0 1
$ echo $?
0
```

Following that, we can observe that the remaining 8 nodes are the ones we want.

```sh
$ kubectl get nodes -o wide
NAME                                 STATUS   ROLES    AGE   VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
k8s-agentpool1-10367588-vmss000002   Ready    agent    25m   v1.18.3   10.240.0.96    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000003   Ready    agent    26m   v1.18.3   10.240.0.127   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000004   Ready    agent    26m   v1.18.3   10.240.0.158   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000005   Ready    agent    27m   v1.18.3   10.240.0.189   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000006   Ready    agent    27m   v1.18.3   10.240.0.220   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000007   Ready    agent    27m   v1.18.3   10.240.0.251   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000008   Ready    agent    26m   v1.18.3   10.240.1.26    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000009   Ready    agent    27m   v1.18.3   10.240.1.57    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-master-10367588-0                Ready    master   31m   v1.18.3   10.255.255.5   <none>        Ubuntu 18.04.4 LTS   5.3.0-1022-azure   docker://3.0.12+azure
```

Now the node pool is once again homogeneous, and all future VMSS scale operations against the VMSS API will render the nodes using the new model.

### My cluster is in a no egress (airgap) environment, using a newer version of AKS Engine to scale isn't working. What's wrong?

AKS Engine curates a VHD (publicly available OS image) for each released version which ensures that all required components are pre-installed onto the VM for all versions of Kubernetes that are supported for that particular AKS Engine release (as a rule AKS Engine supports the latest 2 known-working patch versions of any given supported Kubernetes minor version at the time of release). Because those required components are already present on the VHD, so long as you're installing an AKS Engine-supported version of Kubernetes, your cluster operation will not have to traverse out the public internet (or even traverse outside your VNET to Azure) to bootstrap the Kubernetes runtime.

However, if you're running an operation like `aks-engine` scale using a newer version of the `aks-engine` CLI compared to the version used to build your cluster originally, it is very likely that the Kubernetes version support will have evolved in the meanwhile. Using the above examples, we can observe that the original version of `aks-engine` (v0.52.1 in our example) delivered a `v1.18.3` version of Kubernetes 1.18, and that using the newer version of `aks-engine` (v0.56.0 in our example) respected that (although it *did*, by design update the container runtime, as well as various other OS-layer bits). tl;dr "We still have a `v1.18.3` node pool."

While the above outcome is fine for clusters built in VNETs with permissive egress, if your VNET does not permit general egress to the public internet, you may observe that a newer version of `aks-engine` is not able to successfully complete an operation like the above. To overcome this, we can try two things: (1) obtain the list of supported Kubernetes versions that the newer version of the `aks-engine` CLI uses, and then (2) manually update your API model to explicitly require that newer Kubernetes version. For example:

Let's get the list of supported Kubernetes versions in `v0.56.0` of `aks-engine`:

```sh
$ aks-engine get-versions
Version Upgrades
1.19.1
1.19.0  1.19.1
1.18.8  1.19.0, 1.19.1
1.18.6  1.18.8, 1.19.0, 1.19.1
1.17.11 1.18.6, 1.18.8
1.17.9  1.17.11, 1.18.6, 1.18.8
1.16.15 1.17.9, 1.17.11
1.16.14 1.16.15, 1.17.9, 1.17.11
1.15.12 1.16.14, 1.16.15
1.15.11 1.15.12, 1.16.14, 1.16.15
1.6.9   1.15.11, 1.15.12
```

We can see above that for Kubernetes 1.18, the `aks-engine` CLI being invoked now supports `v1.18.6` and `v1.18.8`. As we expect based on our observations, the API model requires `v1.18.3`:

```sh
$ grep orchestratorVersion _output/kubernetes-westus2-95121/apimodel.json
      "orchestratorVersion": "1.18.3",
```

So, let's manually update that file to `"1.18.8"` instead (using vim or your preferred editor), to declare that we want the most recent, AKS Engine-supported 1.18 version of Kuberentes. After we do that:

```sh
$ grep orchestratorVersion _output/kubernetes-westus2-95121/apimodel.json
      "orchestratorVersion": "1.18.8",
```

Now, let's try that scale operation again!

```sh
$ bin/aks-engine scale --subscription-id $AZURE_SUB_ID --client-id $AZURE_SP_ID --client-secret $AZURE_SP_PW --api-model _output/$RESOURCE_GROUP/apimodel.json --location westus2 --resource-group $RESOURCE_GROUP --apiserver $RESOURCE_GROUP.westus2.cloudapp.azure.com --node-pool agentpool1 --new-node-count 10 --auth-method client_secret --identity-system azure_ad
INFO[0004] found VMSS k8s-agentpool1-10367588-vmss in resource group kubernetes-westus2-95121 that correlates with node pool agentpool1
WARN[0004] Any new nodes will have Moby version 19.03.12
WARN[0004] containerd will be upgraded to version 1.3.7
INFO[0004] Removing singlePlacementGroup property from [variables('agentpool1VMNamePrefix')]
INFO[0004] Nodes in pool 'agentpool1' before scaling:
NODE                                  STATUS    VERSION    OS                    KERNEL
k8s-agentpool1-10367588-vmss000002    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000003    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000004    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000005    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000006    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000007    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000008    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000009    Ready     v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
INFO[0004] Starting ARM Deployment kubernetes-westus2-95121-1270661800 in resource group kubernetes-westus2-95121. This will take some time...
INFO[0174] Finished ARM Deployment (kubernetes-westus2-95121-1270661800). Succeeded
INFO[0174] Nodes in pool 'agentpool1' after scaling:
NODE                                  STATUS      VERSION    OS                    KERNEL
k8s-agentpool1-10367588-vmss000002    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000003    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000004    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000005    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000006    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000007    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000008    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss000009    Ready       v1.18.3    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss00000a    NotReady    v1.18.8    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-agentpool1-10367588-vmss00000b    NotReady    v1.18.8    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
```

Shortly after the new nodes are `Ready`, and running `v1.18.8`:

```
$ k get nodes -o wide
NAME                                 STATUS   ROLES    AGE   VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
k8s-agentpool1-10367588-vmss000002   Ready    agent    49m   v1.18.3   10.240.0.96    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000003   Ready    agent    49m   v1.18.3   10.240.0.127   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000004   Ready    agent    50m   v1.18.3   10.240.0.158   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000005   Ready    agent    50m   v1.18.3   10.240.0.189   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000006   Ready    agent    50m   v1.18.3   10.240.0.220   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000007   Ready    agent    50m   v1.18.3   10.240.0.251   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000008   Ready    agent    50m   v1.18.3   10.240.1.26    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss000009   Ready    agent    50m   v1.18.3   10.240.1.57    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss00000a   Ready    agent    65s   v1.18.8   10.240.0.34    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-agentpool1-10367588-vmss00000b   Ready    agent    68s   v1.18.8   10.240.0.65    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-master-10367588-0                Ready    master   55m   v1.18.3   10.255.255.5   <none>        Ubuntu 18.04.4 LTS   5.3.0-1022-azure   docker://3.0.12+azure
```

In summary, by updating your API model to require the latest Kubernetes version, we produce an ARM template deployment that is able to be executed successully without traversing outside the VNET. As before, we've lost strict homogeneity, but because Kubernetes guarantees functional compatibility within a minor release channel (no breaking changes with patch releases), we now have an arguably more operationally stable cluster running the latest validated bits.
