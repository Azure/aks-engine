# Microsoft Azure Kubernetes Service Engine - Using GPUs with Kubernetes

If you created a Kubernetes cluster with one or multiple agent pool(s) whose VM size is `Standard_NC*` or `Standard_NV*` you can schedule GPU workload on your cluster.
The NVIDIA drivers are automatically installed on every GPU agent in your cluster, so you don't need to do that manually, unless you require a specific version of the drivers. Currently, the installed driver is version 396.26.

To make sure everything is fine, run `kubectl describe node <name-of-a-gpu-node>`. You should see the correct number of GPU reported (in this example shows 2 GPU for a NC12 VM):

For Kubernetes v1.10+ clusters (using NVIDIA Device Plugin):

```
[...]
Capacity:
 nvidia.com/gpu:  2
 cpu:            12
[...]
```

For Kubernetes v1.6, v1.7, v1.8 and v1.9 clusters:

```
[...]
Capacity:
 alpha.kubernetes.io/nvidia-gpu:  2
 cpu:                            12
[...]
```

If `alpha.kubernetes.io/nvidia-gpu` or `nvidia.com/gpu` is `0` and you just created the cluster, you might have to wait a little bit. The driver installation takes about 12 minutes, and the node might join the cluster before the installation is completed. After a few minute the node should restart, and report the correct number of GPUs.

## Running a GPU-enabled container

When running a GPU container, you will need to specify how many GPU you want to use. If you don't specify a GPU count, kubernetes will asumme you don't require any, and will not map the device into the container.
You will also need to mount the drivers from the host (the kubernetes agent) into the container.

On the host, the drivers are installed under `/usr/local/nvidia`.

Here is an example template running TensorFlow:

For Kubernetes v1.10+ clusters (using NVIDIA Device Plugin):

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

For Kubernetes v1.6, v1.7, v1.8 and v1.9 clusters:

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
            alpha.kubernetes.io/nvidia-gpu: 1
        volumeMounts:
        - name: nvidia
          mountPath: /usr/local/nvidia
      volumes:
        - name: nvidia
          hostPath:
            path: /usr/local/nvidia
```

We specify `nvidia.com/gpu: 1` or `alpha.kubernetes.io/nvidia-gpu: 1` in the resources limits. For v1.6 to v1.9 clusters, we need to mount the drivers from the host into the container.
