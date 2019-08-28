# node-labeler

Node Labeler applies master- or agent-specific labels to nodes in an Azure
Kubernetes cluster.

In Kubernetes 1.16, the `node-role.kubernetes.io` and `kubernetes.io/role`
labels can no longer be set with `kubelet --node-labels`, so
`kubernetes.azure.com/role` will replace it. But the Azure cloud-provider and
other components still expect the deprecated labels.

## Getting Started

Build the node-labeler image:

```shell
export DOCKER_IMAGE="quay.io/mboersma/node-labeler:latest"  # for example
docker build -t $DOCKER_IMAGE .
docker push $DOCKER_IMAGE
```

Edit the `image` field in `node-labeler-pod.yaml`, then apply the manifests:

```shell
kubectl -n kube-system apply -f node-labeler-rbac.yaml
kubectl -n kube-system apply -f node-labeler-pod.yaml
```

Remove labels from a node so node-labeler will reapply them:

```shell
kubectl label node/k8s-master-26399701-0 --overwrite kubernetes.io/role-
```

Check the logs to see that the labeling hook ran:

```shell
kubectl -n kube-system logs node-labeler
```

### Removing node-labeler

```shell
kubectl -n kube-system delete pod azure-node-labeler
kubectl -n kube-system delete clusterrolebinding azure-node-labeler
kubectl -n kube-system delete clusterrole azure-node-labeler
kubectl -n kube-system delete serviceaccount azure-node-labeler
docker rmi $DOCKER_IMAGE
```
