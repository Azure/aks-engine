# Cluster certificate rotation

Instructions on how to rotate TLS CA and certs for an AKS Engine cluster.

## Prerequesites

- Kubernetes version 1.8.0 or later is required
- The etcd members MUST be in healthy state before rotating the CA and certs.
- The `_output` directory for your cluster deployment, including working SSH key for accessing the master nodes and `kubeconfig`.

## Preparation

**CAUTION**: Before rotation, it's preferrable to back up your cluster:

1. On an etcd node, take a backup of the current state.
2. //TODO recommend etcd backup tool, extract the existing state from etcd. Copy this back to your working machine as a precaution, in case the control plane goes down.
3. Keep a copy of your current _output directory

## Rotation

**CAUTION**: Rotating certificates will cause cluster downtime

run `aks-engine rotate-certs`

`aks-engine rotate-certs` will:

- Use part of the AKS-Engine deploy code to generate the certificates we need, using the same cert properties and default expiration as new cluster certs
- From master rotate etcd CA and certificates in all masters
- From master rotate CA and certs in the nodes
- Update kubeconfig and rotate the kubelet certs

## Reboot Cluster

After updating the CA and certs of the API server, we need to restart all the pods to ensure they refresh their service account.

This can be done by rebooting all the nodes in the cluster.

// TODO: add `aks-engine` reboot command

## Verification

After the above steps, you can verify the success of the CA and certs rotation:

- Old  `kubeconfig`  should  **NOT**  be able to contact the API server, however the new kubeconfig should be able to talk to it.
- All nodes are expected to be  `Ready`, all pods are expected to be  `Running`
- Try to fetch the logs of  `kube-apiserver`,  `kube-scheduler`  and  `kube-controller-namager`, they should all be running correctly without spitting errors. E.g. ```kubectl logs -lk8s-app=kube-scheduler -n kube-system`

### Sources

https://kubernetes.io/docs/tasks/tls/certificate-rotation/
https://github.com/coreos/tls_rotate