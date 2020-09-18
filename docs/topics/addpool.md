# Adding New Node Pools

## Prerequisites

All the commands in this guide require both the Azure `az` CLI tool and the `aks-engine` binary tool. Follow the [quickstart guide](../tutorials/quickstart.md) before continuing if you're creating a Kubernetes cluster using AKS Engine for the first time.

This guide assumes you already have a running cluster deployed using the `aks-engine` CLI. For more details on how to do that see [deploy](deploy.md) or [generate](generate.md).

## Addpool

The `aks-engine addpool` command can add a new node pool to an existing cluster. By specifing a new `agentPoolProfile` configuration as a JSON file, `aks-engine addpool` will add a node pool according to that configuration, and merge it into the pre-existing aks-engine-generated `apimodel.json`. When used in combination with a newer version of the `aks-engine` CLI compared to the version used to build the cluster originally, new node pools can be regularly added with the latest bits.

The example below will assume you have a cluster deployed, and that the API model originally used to deploy that cluster is stored at `_output/<dnsPrefix>/apimodel.json`.

To add a new pool to the cluster you will run a command like:

```sh
$ aks-engine addpool --subscription-id <subscription_id> \
    --resource-group mycluster --location <location> \
    --client-id '<service principal client ID>' \
    --client-secret '<service principal client secret>' \
    --api-model _output/mycluster/apimodel.json \
    --node-pool ./pool.json
```

The above assumes that the new node pool specification is in the current working directory, and called `pool.json`. Here's an example of what that file might look like:

```json
{
	"name": "pooladded",
	"count": 5,
	"vmSize": "Standard_D4s_v3",
	"availabilityProfile": "VirtualMachineScaleSets",
	"kubernetesConfig": {
			"kubeletConfig": {
			"--cloud-provider": "",
			"--cloud-config": "",
			"--azure-container-registry-config": ""
		}
	}
}
```

The above is a JSON object that conforms to the `agentPoolProfile` specification, just like in the API model. The `agentPoolProfile` spec is documented (here)[clusterdefinitions.md#agentpoolprofiles].

Some important considerations:

- The `"name"` value in a new pool must be unique; it may not be the same value as an existing node pool.
- The `"availabilityProfile"` value in a new pool must match the value in the existing cluster node pools. That enforced homogeneity is an AKS Engine limitation with how its provisioned LoadBalancer resources manage backend pool membership across all nodes in the cluster for svc ingress routing.
- The resultant, new Kubernetes node provisioned in your cluster is not entirely configured via its `agentPoolProfile` specification. It will also inherit certain properties from other configuration in the API model. Specifically, the version of Kubernetes may be modified in the API model JSON (not the JSON file expressing the new pool), and the new pool will be built running that version of Kubernetes. This can support experimenting with new versions of Kubernetes on new nodes (perhaps tainted or excluded from the cluster LoadBalancer) before rolling out that new version cluster-wide.
- All new nodes in the adde pool will be added to the backend pool of the Azure LoadBalancer that serves cluster svc ingress traffic. In practice this means that these new nodes can run pods that support inbound svc traffic coming into the cluster.

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
|--node-pool|yes|Path to JSON file expressing the `agentPoolProfile` spec of the new node pool.|
|--auth-method|no|The authentication method used. Default value is `client_secret`. Other supported values are: `cli`, `client_certificate`, and `device`.|
|--language|no|Language to return error message in. Default value is "en-us").|

## Frequently Asked Questions

### Why would I use addpool instead of update to upgrade a VMSS node pool?

Similar to `aks-engine update`, you may use the `addpool` command to try out a new node configuration in your cluster without affecting existing nodes or production workloads (although if your new configuration is risky in any way you will want to taint those nodes so that no production workloads are scheduled, until you can validate the new configuration). The primary differences are:

- Use `addpool` when the configuration delta compared to an existing node pool is significant enough where it makes sense to organize discretely. Especially if the new pool will only serve a particular type of traffic (e.g., GPU or confidential compute), a dedicated pool should be used for easy, discrete scaling in response to the specific load requirements of the specific workloads it will run.
- Use `addpool` when you want to run operational tests immediately, and especially if you want know the specific number of net new nodes to add. The primary operational difference between `addpool` and `update` is that `addpool` actually adds new operational capacity to your cluster immediately, whereas `update` merely changes the VMSS model, so that *the next* scale out operation renders a node with the new configuration.

### Why would I use addpool instead of upgrade to install a newer version of Kubernetes on my cluster?

If you're running a very large Kubernetes cluster, the one-node-at-a-time operation of `aks-engine upgrade` will take many hours, even days, depending on the size of the cluster. Each one of those node deletions + node additions is subject to environmental failures, and so a deterministic upgrade can indeed take many days. Depending on your tolerance for temporary additional quota, you can upgrade your nodes more quickly, one pool at a time, and use your own validation criteria to inform progression through the entire node pool upgrade operation. Let's demonstrate how that might work using a cluster with 3 node pools:

```sh
$ kubectl get nodes -o wide
NAME                            STATUS   ROLES    AGE     VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
k8s-master-26196714-0           Ready    master   3m7s    v1.18.8   10.255.255.5   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool1-26196714-vmss000000   Ready    agent    3m7s    v1.18.8   10.240.0.34    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool1-26196714-vmss000001   Ready    agent    103s    v1.18.8   10.240.0.65    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool1-26196714-vmss000002   Ready    agent    3m7s    v1.18.8   10.240.0.96    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000000   Ready    agent    3m7s    v1.18.8   10.240.1.181   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000001   Ready    agent    3m      v1.18.8   10.240.1.212   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000002   Ready    agent    3m      v1.18.8   10.240.1.243   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000000   Ready    agent    3m7s    v1.18.8   10.240.0.127   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000001   Ready    agent    2m32s   v1.18.8   10.240.0.158   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000002   Ready    agent    3m7s    v1.18.8   10.240.0.189   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000003   Ready    agent    3m7s    v1.18.8   10.240.0.220   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000004   Ready    agent    3m7s    v1.18.8   10.240.0.251   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000005   Ready    agent    3m7s    v1.18.8   10.240.1.26    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000006   Ready    agent    3m7s    v1.18.8   10.240.1.57    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000007   Ready    agent    3m7s    v1.18.8   10.240.1.88    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000008   Ready    agent    3m7s    v1.18.8   10.240.1.119   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000009   Ready    agent    3m7s    v1.18.8   10.240.1.150   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
```

Above we have a `pool1` with 3 nodes, a `pool2` with 3 nodes, and a `pool3` with 10 nodes. Rather than run a single, continuous upgrade operation across all nodes in the cluster, let's add pools, then validate the new version, and then scale those new pools up so the original nodes can be cordoned, drained, and deleted.

Here we use the command line `jq` tool to create three new JSON files that we'll use to initiate 3 new `aks-engine addpool` operations, derived from the original `agentPoolProfile` specifications in the API model generated during cluster deployment:

```sh
$ jq -r '.properties.agentPoolProfiles[0] | .name = "newpool1"' < _output/kubernetes-westus2-1838/apimodel.json > newpool1.json
$ jq -r '.properties.agentPoolProfiles[1] | .name = "newpool2"' < _output/kubernetes-westus2-1838/apimodel.json > newpool2.json
$ jq -r '.properties.agentPoolProfiles[2] | .name = "newpool3"' < _output/kubernetes-westus2-1838/apimodel.json > newpool3.json
```

Because those were derived from the API, those new pools are configured with a count of 3, 3, and 10, respectively. Let's change all of the node counts to 1, because we don't necessarily need full node pool capacity to validate the new Kubernetes versions against:

```sh
$ jq -r '.count = 1' < newpool1.json > newpool1-1node.json && mv newpool1-1node.json newpool1.json
$ jq -r '.count = 1' < newpool2.json > newpool2-1node.json && mv newpool2-1node.json newpool2.json
$ jq -r '.count = 1' < newpool3.json > newpool3-1node.json && mv newpool3-1node.json newpool3.json
```

Our final configuration change before running `aks-engine addpool` is updating the Kubernetes in the API model

```sh
$ jq -r '. | .properties.orchestratorProfile.orchestratorRelease = "1.19"' < _output/kubernetes-westus2-1838/apimodel.json > apimodel-1dot19.json
FrancisBookMS:aks-engine jackfrancis$ jq -r '. | .properties.orchestratorProfile.orchestratorVersion = "1.19.1"' < apimodel-1dot19.json > _output/kubernetes-westus2-1838/apimodel.json
$ grep orchestratorRelease -A 1 _output/kubernetes-westus2-1838/apimodel.json
      "orchestratorRelease": "1.19",
      "orchestratorVersion": "1.19.1",
```

We can now run addpool once per new pool to begin the process of validating v1.19.1 across our existing v1.18.8 cluster:

```sh
$ aks-engine addpool --subscription-id $TEST_AZURE_SUB_ID --api-model _output/kubernetes-westus2-1838/apimodel.json --node-pool newpool1.json --location westus2 --resource-group kubernetes-westus2-1838 --auth-method client_secret --client-id $TEST_AZURE_SP_ID --client-secret $TEST_AZURE_SP_PW
WARN[0003] Any new nodes will have containerd version 1.3.7
INFO[0003] Starting ARM Deployment kubernetes-westus2-1838-1942811440 in resource group kubernetes-westus2-1838. This will take some time...
INFO[0158] Finished ARM Deployment (kubernetes-westus2-1838-1942811440). Succeeded
$ aks-engine addpool --subscription-id $TEST_AZURE_SUB_ID --api-model _output/kubernetes-westus2-1838/apimodel.json --node-pool newpool2.json --location westus2 --resource-group kubernetes-westus2-1838 --auth-method client_secret --client-id $TEST_AZURE_SP_ID --client-secret $TEST_AZURE_SP_PW
WARN[0008] Any new nodes will have containerd version 1.3.7
INFO[0008] Starting ARM Deployment kubernetes-westus2-1838-25937475 in resource group kubernetes-westus2-1838. This will take some time...
INFO[0163] Finished ARM Deployment (kubernetes-westus2-1838-25937475). Succeeded
$ aks-engine addpool --subscription-id $TEST_AZURE_SUB_ID --api-model _output/kubernetes-westus2-1838/apimodel.json --node-pool newpool3.json --location westus2 --resource-group kubernetes-westus2-1838 --auth-method client_secret --client-id $TEST_AZURE_SP_ID --client-secret $TEST_AZURE_SP_PW
WARN[0004] Any new nodes will have containerd version 1.3.7
INFO[0004] Starting ARM Deployment kubernetes-westus2-1838-1370618455 in resource group kubernetes-westus2-1838. This will take some time...
INFO[0174] Finished ARM Deployment (kubernetes-westus2-1838-1370618455). Succeeded
```

At this point we now have three new nodes running v1.19.1 on our cluster, one per new pool, which correlates with one new pool per pre-existing pool:

```sh
$ k get nodes -o wide
NAME                               STATUS   ROLES    AGE     VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
k8s-master-26196714-0              Ready    master   36m     v1.18.8   10.255.255.5   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool1-26196714-vmss000000   Ready    agent    8m35s   v1.19.1   10.240.2.18    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool2-26196714-vmss000000   Ready    agent    3m41s   v1.19.1   10.240.2.49    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool3-26196714-vmss000000   Ready    agent    21s     v1.19.1   10.240.2.80    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool1-26196714-vmss000000      Ready    agent    36m     v1.18.8   10.240.0.34    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool1-26196714-vmss000001      Ready    agent    35m     v1.18.8   10.240.0.65    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool1-26196714-vmss000002      Ready    agent    36m     v1.18.8   10.240.0.96    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000000      Ready    agent    36m     v1.18.8   10.240.1.181   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000001      Ready    agent    36m     v1.18.8   10.240.1.212   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000002      Ready    agent    36m     v1.18.8   10.240.1.243   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000000      Ready    agent    36m     v1.18.8   10.240.0.127   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000001      Ready    agent    36m     v1.18.8   10.240.0.158   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000002      Ready    agent    36m     v1.18.8   10.240.0.189   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000003      Ready    agent    36m     v1.18.8   10.240.0.220   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000004      Ready    agent    36m     v1.18.8   10.240.0.251   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000005      Ready    agent    36m     v1.18.8   10.240.1.26    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000006      Ready    agent    36m     v1.18.8   10.240.1.57    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000007      Ready    agent    36m     v1.18.8   10.240.1.88    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000008      Ready    agent    36m     v1.18.8   10.240.1.119   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000009      Ready    agent    36m     v1.18.8   10.240.1.150   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
```

At this point we would probably taint those three nodes, and then run validations against them (using the appropriate tolerations so that they were scheduled onto the desired nodes):

```sh
$ kubectl taint nodes k8s-newpool1-26196714-vmss000000 validating:NoSchedule
node/k8s-newpool1-26196714-vmss000000 tainted
$ kubectl taint nodes k8s-newpool2-26196714-vmss000000 validating:NoSchedule
node/k8s-newpool2-26196714-vmss000000 tainted
$ kubectl taint nodes k8s-newpool3-26196714-vmss000000 validating:NoSchedule
node/k8s-newpool3-26196714-vmss000000 tainted
```

Let's say we've validated the "pool1" replacement, which we've called "newpool1". Let's scale that pool out to match the original "pool1":

```sh
$ aks-engine scale --subscription-id $TEST_AZURE_SUB_ID --client-id $TEST_AZURE_SP_ID --client-secret $TEST_AZURE_SP_PW --api-model _output/kubernetes-westus2-1838/apimodel.json --location westus2 --resource-group kubernetes-westus2-1838 --apiserver kubernetes-westus2-1838.westus2.cloudapp.azure.com --node-pool newpool1 --new-node-count 3 --auth-method client_secret --identity-system azure_ad
INFO[0003] found VMSS k8s-newpool1-26196714-vmss in resource group kubernetes-westus2-1838 that correlates with node pool newpool1
WARN[0003] Any new nodes will have containerd version 1.3.7
INFO[0003] Removing singlePlacementGroup property from [variables('newpool1VMNamePrefix')]
INFO[0003] Nodes in pool 'newpool1' before scaling:
NODE                                STATUS    VERSION    OS                    KERNEL
k8s-newpool1-26196714-vmss000000    Ready     v1.19.1    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
INFO[0003] Starting ARM Deployment kubernetes-westus2-1838-360281667 in resource group kubernetes-westus2-1838. This will take some time...
INFO[0230] Finished ARM Deployment (kubernetes-westus2-1838-360281667). Succeeded
INFO[0230] Nodes in pool 'newpool1' after scaling:
NODE                                STATUS      VERSION    OS                    KERNEL
k8s-newpool1-26196714-vmss000000    Ready       v1.19.1    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-newpool1-26196714-vmss000001    Ready       v1.19.1    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
k8s-newpool1-26196714-vmss000002    NotReady    v1.19.1    Ubuntu 18.04.5 LTS    5.4.0-1025-azure
```

Note: you may also use the VMSS API directly (either via the `az` CLI or the Azure portal web UI) to scale out the new pools. The advantage of using `aks-engine scale` to do so is that you will get immediate signal if, for any reason, the new VMs did not come online successfully as Kubernetes nodes.

Now that we have equivalent node capacity for our new pool compared to our original pool (note: "capacity equivalence" may be a little more complicated if, as part of a process like this one, you change the VM SKU of the new pool as compared to the original pool; YMMV.) we can cordon + drain the original nodes and rely upon the Kubernetes layer to re-schedule workloads to the new nodes (note: this will require you to really ensure your workload scheduling configuration as pertains to the way your nodes are labeled, tainted, etc, makes sense and that your production workload specifications adhere to that configuration schema).

```sh
$ for node in "k8s-pool1-26196714-vmss000000 k8s-pool1-26196714-vmss000001 k8s-pool1-26196714-vmss000002"; do kubectl cordon $node; done
node/k8s-pool1-26196714-vmss000000 cordoned
node/k8s-pool1-26196714-vmss000001 cordoned
node/k8s-pool1-26196714-vmss000002 cordoned
$ for node in "k8s-pool1-26196714-vmss000000 k8s-pool1-26196714-vmss000001 k8s-pool1-26196714-vmss000002"; do kubectl drain --ignore-daemonsets $node; done
node/k8s-pool1-26196714-vmss000000 already cordoned
node/k8s-pool1-26196714-vmss000001 already cordoned
node/k8s-pool1-26196714-vmss000002 already cordoned
WARNING: ignoring DaemonSet-managed Pods: kube-system/azure-cni-networkmonitor-z4tcw, kube-system/azure-ip-masq-agent-nmlnv, kube-system/blobfuse-flexvol-installer-zgjxg, kube-system/csi-secrets-store-jdmql, kube-system/csi-secrets-store-provider-azure-9d4j9, kube-system/kube-proxy-glrm6
node/k8s-pool1-26196714-vmss000000 drained
WARNING: ignoring DaemonSet-managed Pods: kube-system/azure-cni-networkmonitor-xhk8d, kube-system/azure-ip-masq-agent-lhj9p, kube-system/blobfuse-flexvol-installer-zdc4w, kube-system/csi-secrets-store-6zbx9, kube-system/csi-secrets-store-provider-azure-q2h6n, kube-system/kube-proxy-728sx
node/k8s-pool1-26196714-vmss000001 drained
WARNING: ignoring DaemonSet-managed Pods: kube-system/azure-cni-networkmonitor-mtx7c, kube-system/azure-ip-masq-agent-5p9lw, kube-system/blobfuse-flexvol-installer-cl9ls, kube-system/csi-secrets-store-provider-azure-vv8rb, kube-system/csi-secrets-store-xnjxn, kube-system/kube-proxy-rpfjt
node/k8s-pool1-26196714-vmss000002 drained
```

Note: the above example is rather brute-force. Depending on your operational reality, you may want to add some delay between draining each node. (cordon'ing all nodes at once actually makes sense, as you indeed want to stop any future scheduling onto those nodes all at the same time, once you have the required standby capacity, which in our example is the new, validated v1.19.1 nodes)

After all workloads have been drained, and moved over to the new nodes, you may delete the VMSS entirely:

```sh
$ az vmss delete -n k8s-pool1-26196714-vmss -g kubernetes-westus2-1838
$ echo $?
0
```

Now, the original "pool1" nodes are no longer participating in the cluster:

```sh
$ k get nodes -o wide
NAME                               STATUS   ROLES    AGE   VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
k8s-master-26196714-0              Ready    master   64m   v1.18.8   10.255.255.5   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool1-26196714-vmss000000   Ready    agent    36m   v1.19.1   10.240.2.18    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool1-26196714-vmss000001   Ready    agent    17m   v1.19.1   10.240.2.111   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool1-26196714-vmss000002   Ready    agent    16m   v1.19.1   10.240.2.142   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool2-26196714-vmss000000   Ready    agent    31m   v1.19.1   10.240.2.49    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-newpool3-26196714-vmss000000   Ready    agent    28m   v1.19.1   10.240.2.80    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000000      Ready    agent    64m   v1.18.8   10.240.1.181   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000001      Ready    agent    64m   v1.18.8   10.240.1.212   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool2-26196714-vmss000002      Ready    agent    64m   v1.18.8   10.240.1.243   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000000      Ready    agent    64m   v1.18.8   10.240.0.127   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000001      Ready    agent    63m   v1.18.8   10.240.0.158   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000002      Ready    agent    64m   v1.18.8   10.240.0.189   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000003      Ready    agent    64m   v1.18.8   10.240.0.220   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000004      Ready    agent    64m   v1.18.8   10.240.0.251   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000005      Ready    agent    64m   v1.18.8   10.240.1.26    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000006      Ready    agent    64m   v1.18.8   10.240.1.57    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000007      Ready    agent    64m   v1.18.8   10.240.1.88    <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000008      Ready    agent    64m   v1.18.8   10.240.1.119   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
k8s-pool3-26196714-vmss000009      Ready    agent    64m   v1.18.8   10.240.1.150   <none>        Ubuntu 18.04.5 LTS   5.4.0-1025-azure   docker://19.3.12
```

Final note: don't forget to remove the "pool1" `agentPoolProfile` JSON object from your API model!

### How do I integrate any added node pools into an existing cluster-autoscaler configuration?
