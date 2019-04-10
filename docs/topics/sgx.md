# Using SGX with Kubernetes

[Intel&reg; Secure Guard Extension](https://software.intel.com/en-us/sgx) (Intel&reg; SGX) is an architecture extension designed to increase the security of application code and data.
Developers may choose [Intel&reg; SGX SDK](https://software.intel.com/en-us/sgx-sdk) or [Open Enclave SDK](https://github.com/Microsoft/openenclave/) to create applications that leverage this technology.

Azure supports provisioning of SGX-enabled VMs under the umbrella of Azure Confidential Compute (ACC). You can create a Kubernetes cluster with one or multiple agent pool(s) running ACC VMs by specifying a [DC-series](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes-general#dc-series) VM size with a supported distro from the table below.

| OS           | distro      | Notes |
| ------------ | ----------- |-------|
| Ubuntu 16.04 | `acc-16.04` | specially built image with UEFI BIOS support
| Ubuntu 18.04 | `aks-18.04` | AKS-maintained Ubuntu 18.04 image with preinstalled components

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
      "distro": "aks-18.04",
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
    volumeDevices:
    - devicePath: /dev/sgx
      name: dev-sgx
    securityContext:
      privileged: true
  volumes:
  - name: dev-sgx
    hostPath:
      path: /dev/sgx
```
