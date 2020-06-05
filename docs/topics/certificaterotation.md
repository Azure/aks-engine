# Cluster certificate rotation

Instructions on rotating TLS CA and certificates for an AKS Engine cluster.

## Prerequesites

- The etcd members MUST be in a healthy state before rotating the CA and certs (ie. `etcdctl cluster-health` shows all peers are healthy and cluster is healthy).
- The API model file reflecting the current cluster configuration and a working ssh private key that has root access to all nodes. The API model file is persisted at AKS Engine template generation time, by default to the _output/ child directory from the working parent directory at the time of the `aks-engine` invocation.

<a name="preparation"></a>

## Preparation

**CAUTION**: Rotating certificates can break component connectivity and leave the cluster in an unrecoverable state. Before performing any of these instructions on a live cluster, it is preferrable to backup your cluster state and migrate critical workloads to another cluster.

## Rotation

**CAUTION**: Rotating certificates will cause cluster downtime.

run `aks-engine rotate-certs`. For example:

```bash
CLUSTER="<CLUSTER_DNS_PREFIX>" && bin/aks-engine rotate-certs --api-model _output/${CLUSTER}/apimodel.json
--client-id "<YOUR_CLIENT_ID>" --client-secret "<YOUR_CLIENT_SECRET>" --location <CLUSTER_LOCATION>
--apiserver ${CLUSTER}.<CLUSTER_LOCATION>.cloudapp.azure.com --ssh _output/${CLUSTER}-ssh
--subscription-id "<YOUR_SUBSCRIPTION_ID>" -g ${CLUSTER} -o _output/${CLUSTER}
```

The above example command assumes a default `output/` directory with the resource group name being the same as the cluster's DNS prefix.

`aks-engine rotate-certs` will:

- Generate new certificates.
- Rotate apiserver certificates.
- Rotate the kubelet certificates.
- Rotate etcd CA and certificates and restart etcd in all of the master nodes.
- Update the kubeconfig.
- Reboot all the VMs in the resource group.
- Restart all the pods to ensure they refresh their service account.

## Verification

After the above steps, you can verify the success of the CA and certs rotation:

- The old  `kubeconfig`  should  **NOT**  be able to contact the API server, however the new `kubeconfig` should be able to talk to it.
- All nodes are expected to be  `Ready`, all pods are expected to be  `Running`.
- Try to fetch the logs of  `kube-apiserver`,  `kube-scheduler`  and  `kube-controller-namager`. They should all be running correctly without printing errors. E.g. `kubectl logs kube-apiserver-k8s-master-58431286-0 -n kube-system`.

## Known Limitations

The certificate rotation tool has not been tested on and is expected to fail with the following cluster configurations:

- Private clusters.
- Clusters using keyvault references in Certificate Profile.
- Clusters using Cosmos etcd.
- Clusters with already expired certificates with unhealthy etcd.

The rotation involves rebooting the nodes. ALL VMs in the resource group will be restarted as part of running the `rotate-certs` command. If the resource group contains any VMs that are not part of the cluster, they will be restarted as well.

The tool is not currently idempotent, meaning that if the rotation fails halfway though or is interrupted, you will most likely not be able to re-run the operation without manual intervention. There is a risk that your cluster will become unrecoverable which is why it is strongly recommended to follow the [preparation step](#preparation).
