# Using SGX with Kubernetes

<!-- TOC -->
- [Deploy a Kubernetes Cluster](#deploy-a-kubernetes-cluster)
- [Running a SGX-enabled container](#running-a-sgx-enabled-container)
- [OPTIONAL: Using oe-sgx device plugin (alpha)](#optional-using-oe-sgx-device-plugin-alpha)
   - [Device plugin installation](#device-plugin-installation)
   - [Scheduling Pods to TEE enabled Hardware](#scheduling-pods-to-tee-enabled-hardware)
<!-- /TOC -->

[Intel&reg; Secure Guard Extension](https://software.intel.com/en-us/sgx) (Intel&reg; SGX) is an architecture extension designed to increase the security of application code and data.
Developers may choose [Intel&reg; SGX SDK](https://software.intel.com/en-us/sgx-sdk) or [Open Enclave SDK](https://github.com/Microsoft/openenclave/) to create applications that leverage this technology.

Azure supports provisioning of SGX-enabled VMs under the umbrella of Azure Confidential Compute (ACC). You can create a Kubernetes cluster with one or multiple agent pool(s) running ACC VMs by specifying a [DC-series](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes-general#dc-series) VM size with a supported distro from the table below.

## Deploy a Kubernetes Cluster
Refer to the [Quickstart Guide](../tutorials/quickstart.md) for details on how to provision a cluster using AKS-Engine. In order to use SGX enabled hardware we suggest updating the cluster model to include an additional agentpool with the supported operating system and virtual machine size. See below for further detail.


| OS           | distro      | Notes |
| ------------ | ----------- |-------|
| Ubuntu 16.04 | `acc-16.04` | specially built image with UEFI BIOS support
| Ubuntu 18.04 | `aks-ubuntu-18.04` | AKS-maintained Ubuntu 18.04 image with preinstalled components

The following example is a fragment of a cluster definition (apimodel) file declaring two ACC agent pools, one running `Ubuntu 16.04` image on `2 vCPU` nodes, and another running `Ubuntu 18.04` image on `4 vCPU` nodes:

```
  "agentPoolProfiles": [
    {
      "name": "agentpool1",
      "count": 3,
      "distro": "acc-16.04",
      "vmSize": "Standard_DC2s"
    },
    {
      "name": "agentpool2",
      "count": 3,
      "distro": "aks-ubuntu-18.04",
      "vmSize": "Standard_DC4s"
    }
  ],
```

The SGX driver is automatically installed on every ACC node in your cluster, so you don't need to do that manually.

## Running a SGX-enabled container

When running an SGX container, you will need to mount the drivers from the host (the kubernetes node) into the container.

On the host, the drivers are installed under `/dev/sgx`.

Here is an example template of Pod YAML file:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: <POD NAME>
spec:
  containers:
  - name: <NAME>
    image: <IMAGE>
    command: <COMMAND>
    imagePullPolicy: IfNotPresent
    volumeMounts:
    - name: dev-sgx
      mountPath: /dev/sgx
    securityContext:
      privileged: true
  volumes:
  - name: dev-sgx
    hostPath:
      path: /dev/sgx
      type: CharDevice
```

## OPTIONAL: Using oe-sgx device plugin (alpha)

Alternatively, you can install the oe-sgx device plugin (alpha) which surfaces the usage of Intel SGX’s Encrypted Page Cache (EPC) RAM as a schedulable resource for Kubernetes. This allows you to schedule pods that use the Open Enclave SDK onto hardware which supports Trusted Execution Environments.

### Device plugin installation

We recommend labelling the nodes so that a nodeSelector can be used to run the device plugin only on the nodes that support Intel SGX based Trusted Execution Environments. Use the following command to apply the appropriate labels to the Intel SGX enabled nodes:

`kubectl label nodes <node-name> tee=sgx`

We also recommend tainting the nodes so that only pods that toleration that taint are scheduled to that specific node. Apply the following taints to all nodes that are Intel SGX enabled using the following command:

`kubectl taint nodes <node-name> openenclave.io/sgx_epc_MiB=true:NoSchedule`

Using kubectl, install the device plugin DaemonSet:

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: oe-sgx-device-plugin
  namespace: kube-system
  labels:
    app: oe-sgx-device-plugin
spec:
  selector:
    matchLabels:
      app: oe-sgx-device-plugin
  template:
    metadata:
      labels:
        app: oe-sgx-device-plugin
    spec:
      tolerations:
      - key: openenclave.io/sgx_epc_MiB
        operator: Exists
        effect: NoSchedule
      containers:
      - name: oe-sgx-device-plugin
        image: "mcr.microsoft.com/aks/acc/sgx-device-plugin:0.1"
        command: ["/usr/local/bin/oe-sgx-device-plugin"]
        imagePullPolicy: IfNotPresent
        volumeMounts:
        - name: device-plugin
          mountPath: /var/lib/kubelet/device-plugins
        - name: dev-sgx
          mountPath: /dev/sgx
        securityContext:
          privileged: true
      volumes:
      - name: device-plugin
        hostPath:
          path: /var/lib/kubelet/device-plugins
      - name: dev-sgx
        hostPath:
          path: /dev/sgx
      nodeSelector:
        tee: sgx
```

Confirm that the DaemonSet pods are running on each IntelSGX enabled node as follows:

```bash
$ kubectl get pods -n kube-system -l app=oe-sgx-device-plugin
```

```bash
NAME                         READY   STATUS    RESTARTS   AGE
oe-sgx-device-plugin-7d5l8   1/1     Running   0          12m
oe-sgx-device-plugin-jzhk9   1/1     Running   0          12m
```

Confirm that the device plugin is advertising the available EPC RAM to the Kubernetes scheduler by running the following command:

```bash
$ kubectl get nodes <node-name> -o yaml
```

Under the status field you should see the total allocable resources with a name of `openenclave.io/sgx_epc_MiB` 
```bash
<snip>
status:
  allocatable:
    openenclave.io/sgx_epc_MiB: "82"
<snip>
```

### Scheduling Pods to TEE enabled Hardware

The following pod specification demonstrates how you would schedule a pod to have access to a TEE by defining a limit on the specific EPC memory that is advertised to the Kubernetes scheduler by the device plugin available in alpha. 

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: oe-deployment
spec:
  selector:
    matchLabels:
      app: oe-app
  replicas: 1
  template:
    metadata:
      labels:
        app: oe-app
    spec:
      tolerations:
      - key: openenclave.io/sgx_epc_MiB
        operator: Exists
        effect: NoSchedule
      containers:
      - name: <image_name>
        image: <image_reference>
        command: <exec>
        resources:
          limits:
            openenclave.io/sgx_epc_MiB: 10
```

You can use the following test workload to confirm that your cluster is correctly configured:

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: sgx-test
  labels:
    app: sgx-test
spec:
  template:
    metadata:
      labels:
        app: sgx-test
    spec:
      containers:
      - name: sgxtest
        image: oeciteam/sgx-test
        command: ["/helloworld/host/helloworldhost", "/helloworld/enclave/helloworldenc.signed"]
        volumeMounts:
        - mountPath: /dev/sgx
          name: dev-sgx
        securityContext:
          privileged: true
        resources:
          limits:
            openenclave.io/sgx_epc_MiB: 10
      volumes:
      - name: dev-sgx
        hostPath:
          path: /dev/sgx
          type: CharDevice
      restartPolicy: Never
  backoffLimit: 0
  ```

  You can confirm that the workload successfully created a Trusted Execution Environment by running the following commands:

```bash
$ kubectl get jobs -l app=sgx-test
```

```bash
$ kubectl get jobs -l app=sgx-test
NAME       COMPLETIONS   DURATION   AGE
sgx-test   1/1           1s         23s
```

```bash
$ kubectl get pods -l app=sgx-test
```

```bash
$ kubectl get pods -l app=sgx-test
NAME             READY   STATUS      RESTARTS   AGE
sgx-test-rchvg   0/1     Completed   0          25s
```

```bash
$ kubectl logs -l app=sgx-test
```

```bash
$ kubectl logs -l app=sgx-test
Hello world from the enclave
Enclave called into host to print: Hello World!
```


