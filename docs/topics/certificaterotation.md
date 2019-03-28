# Cluster certificate rotation

Instructions on rotating TLS CA and certificates for an AKS Engine cluster.

## Prerequesites

- Kubernetes version 1.8.0 or later is required.
- The etcd members MUST be in a healthy state before rotating the CA and certs (ie. `etcdctl cluster-health` shows all peers are healthy and cluster is healthy).
- The apimodel file reflecting the current cluster configuration, a working ssh private key that has root access to all master nodes, and a working kubeconfig file in order to establish an authenticated connection to the Kubernetes cluster API. These files are persisted at AKS Engine template generation time, by default to the _output/ child directory from the working parent directory at the time of the aks-engine invocation.

## Preparation

**CAUTION**: Before rotation, it's preferrable to back up your cluster:

// TODO: add more detailed instructions on backing up the etcd cluster data.

## Rotation

**CAUTION**: Rotating certificates will cause cluster downtime.

run `aks-engine rotate-certs`

`aks-engine rotate-certs` will:

- Use the exisiting certificate generation code to generate new certificates.
- Rotate etcd CA and certificates in all of the master nodes.
- From the first master node, rotate apiserver certificates in the nodes.
- Update the kubeconfig and rotate the kubelet certificates.

## Reboot Cluster

After updating the CA and certs of the API server, we need to restart all the pods to ensure they refresh their service account.

This can be done by rebooting all the nodes in the cluster.

// TODO: add `aks-engine` reboot command

## Verification

After the above steps, you can verify the success of the CA and certs rotation:

- Old  `kubeconfig`  should  **NOT**  be able to contact the API server, however the new kubeconfig should be able to talk to it.
- All nodes are expected to be  `Ready`, all pods are expected to be  `Running`.
- Try to fetch the logs of  `kube-apiserver`,  `kube-scheduler`  and  `kube-controller-namager`, they should all be running correctly without spitting errors. E.g. ```kubectl logs -lk8s-app=kube-scheduler -n kube-system`.

## Known Limitations

- Private clusters
- Clusters using keyvault
- Clusters using Cosmos Etcd