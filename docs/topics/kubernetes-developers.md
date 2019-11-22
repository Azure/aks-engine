# For Kubernetes Developers

If you are working on Kubernetes upstream, you can use AKS Engine to test your build of Kubernetes in the Azure environment.

## In a kubernetes/kubernetes PR

If you open a PR on [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes), you could enter the following commands in the comment section to build Kubernetes according to your PR changes, deploy it using AKS Engine, and run various tests:

* `/test pull-kubernetes-e2e-aks-engine-azure`: Run Kubernetes conformance tests
* `/test pull-kubernetes-e2e-azure-disk`: Run e2e tests on in-tree Azure Disk plugin
* `/test pull-kubernetes-e2e-azure-disk-vmss`:  Run e2e tests on in-tree Azure Disk plugin in a VMSS-based cluster

## Building Kubernetes From Source

The following instructions describe in more detail how to create the required Docker image and deploy it using AKS Engine (replace `dockerhubid` and `sometag` with your Docker Hub ID and a unique tag for your build):

> NOTE: This build is extremely memory intensive. If you're using Docker For Mac/For Windows, ensure that your docker daemon has access to atleast 4.5 GB of memory and 2 GB of swap size. For Mac OSX see: https://docs.docker.com/docker-for-mac/#advance. For Windows see: https://docs.docker.com/docker-for-windows/#advanced

The commands to build and publish custom-built Kubernetes are different depending on which Kubernetes version you have. Regardless, they need to be executed in the Kubernetes project root (typically https://github.com/kubernetes/kubernetes or a fork).

## Kubernetes 1.16 or earlier

* Build Kubernetes and [hyperkube](https://github.com/kubernetes/kubernetes/tree/master/cluster/images/hyperkube#hyperkube):

```bash
make clean
make all
make WHAT=cmd/hyperkube
```

* Run the following commands to push the local hyperkube image to a Docker registry (you would need to have access to this registry):

```bash
export VERSION=<your-custom-version>
export REGISTRY=<your-docker-registry>
export HYPERKUBE_BIN=$(pwd)/_output/bin/hyperkube
make -C cluster/images/hyperkube push
```

For `VERSION` environment variable, we recommend that you provide a value which would help you identify the build of kubernetes. It will be used as a tag for your custom hyperkube image.

> NOTE: It's convenient to put these steps into a script.

### AKS Engine API Model

* Open the AKS Engine API Model (e.g. a file from the examples directory). Set `orchestratorRelease` to 1.16 or lower so various defaults and configuration are properly applied to the ARM template and artifacts. Additionally, add the following to the `kubernetesConfig` section:

```
"orchestratorRelease": "1.16",
"kubernetesConfig": {
    ...
    "customHyperkubeImage": "<your-docker-registry>/hyperkube-amd64:<your-custom-version>",
    ...
}
```

* AKS Engine defaults to the `ubuntu` "distro" for `customHyperkubeImage` scenarios, which will build Linux VM nodes using the Ubuntu 16.04-LTS image SKU. You may also use Ubuntu 18.04-LTS-built VMs by explicitly setting the "distro" configuration accordingly. The "distro" configuration is applied distinctly to master VMs, and to the VMs in any configured agent pools (you may use different "distro" values in the masterProfile, and in each agentPoolProfile, according to your cluster requirements), for example:

```
...
"masterProfile": {
    ...
    "distro": "ubuntu-18.04",
    ...
},
...
"agentPoolProfiles": [
    {
        ...
        "distro": "ubuntu-18.04",
        ...
    }
]
...
```

* Run `aks-engine deploy` [as normal](../tutorials/deploy.md).

## Kubernetes 1.17+

Since Kubernetes has removed `make WHAT=cmd/hyperkube` command and AKS Engine has stopped supporting hyperkube for Kubernetes 1.17+, we will have to build individual Kubernetes components and deploy them separately. That includes [kube-apiserver](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/), [kube-controller-manager](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/), [kube-scheduler](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/), [kube-proxy](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/), [kubelet](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/), and [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/).

* Run the following commands to build Kubernetes and push images of individual components:

```bash
export KUBE_DOCKER_REGISTRY=<your-docker-registry>
export KUBE_DOCKER_IMAGE_TAG=<your-custom-tag>
make quick-release
docker push <your-docker-registry>/kube-apiserver-amd64:<your-custom-tag>
docker push <your-docker-registry>/kube-controller-manager-amd64:<your-custom-tag>
docker push <your-docker-registry>/kube-proxy-amd64:<your-custom-tag>
docker push <your-docker-registry>/kube-scheduler-amd64:<your-custom-tag>
```

After that, you will have to upload a tarball (`_output/release-tars/kubernetes-node-linux-amd64.tar.gz`) built from the previous step to online storage such as GitHub or Azure Storage. AKS Engine uses the tarball to extract kubelet and kubectl.

### AKS Engine API Model

* Open the AKS Engine API Model (e.g. a file from the examples directory). Set `orchestratorRelease` to 1.17 or higher so various defaults and configuration are properly applied to the ARM template and artifacts. Additionally, add the following to the `kubernetesConfig` section:

```
"orchestratorRelease": "1.17",
"kubernetesConfig": {
    ...
    "customKubeAPIServerImage": "<your-docker-registry>/kube-apiserver-amd64:<your-custom-tag>",
    "customKubeControllerManagerImage": "<your-docker-registry>/kube-controller-manager-amd64:<your-custom-tag>",
    "customKubeProxyImage": "<your-docker-registry>/kube-proxy-amd64:<your-custom-tag>",
    "customKubeSchedulerImage": "<your-docker-registry>/kube-scheduler-amd64:<your-custom-tag>",
    "customKubeBinaryURL": "<URL to uploaded kubernetes-node-linux-amd64.tar.gz>",
    ...
}
```

## Private Registry

If the container registry is private, for example Azure Container Registry, then provide the name of the private Azure registry along with the custom hyperkube image like this:

```
"kubernetesConfig": {
    ...
    "privateAzureRegistryServer": "<your-private-registry>",
    ...
}
```

> NOTE: Make sure the service principal provided to run `aks-engine deploy` has access to pull images from this private registry. https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-service-principal#use-an-existing-service-principal
