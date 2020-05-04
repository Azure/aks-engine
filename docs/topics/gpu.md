# Using GPUs with Kubernetes

If you created a Kubernetes cluster with one or multiple node pools whose VM size is `Standard_NC*` or `Standard_NV*` you can schedule GPU workloads on your cluster.
The NVIDIA drivers are automatically installed on every GPU agent in your cluster, so you don't need to do that manually, unless you require a specific version of the drivers. Currently, the installed driver is version 418.40.04.

To make sure everything is fine, run `kubectl describe node <name-of-a-gpu-node>`. You should see the correct number of GPU reported (in this example shows 2 GPU for a NC12 VM):

For Kubernetes v1.10+ clusters (using NVIDIA Device Plugin):

```
[...]
Capacity:
 nvidia.com/gpu:  2
 cpu:            12
[...]
```

If `nvidia.com/gpu` is `0` and you just created the cluster, you might have to wait a little bit. The driver installation can add a few minutes to the cluster deployment time compared to non-GPU agent pool scenarios, and the node might join the cluster before the installation is completed. After a few minutes the node should restart, and report the correct number of GPUs.

## Running a GPU-enabled container

When running a GPU container, you will need to specify how many GPU you want to use. If you don't specify a GPU count, kubernetes will asumme you don't require any, and will not map the device into the container.
You will also need to mount the drivers from the host (the kubernetes agent) into the container.

On the host, the drivers are installed under `/usr/local/nvidia`.

Here is an example template running TensorFlow:

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: tensorflow
  name: tensorflow
spec:
  template:
    metadata:
      labels:
        app: tensorflow
    spec:
      containers:
      - name: tensorflow
        image: <SOME_IMAGE>
        command: <SOME_COMMAND>
        imagePullPolicy: IfNotPresent
        resources:
          limits:
            nvidia.com/gpu: 1
```

We specify `nvidia.com/gpu: 1` in the resources limits.
