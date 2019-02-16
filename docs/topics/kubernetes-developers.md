# For Kubernetes Developers

If you're working on Kubernetes upstream, you can use AKS Engine to test your build of Kubernetes in the Azure environment.  The option that allows you to do this is `orchestratorProfile.kubernetesConfig.customHyperkubeImage`, which you should set to point to a Docker image containing your build of hyperkube.

The following instructions describe in more detail how to create the required Docker image and deploy it using AKS Engine (replace `dockerhubid` and `sometag` with your Docker Hub ID and a unique tag for your build):

## In the Kubernetes repo

NOTE: This build is extremely memory intensive. If you're using Docker For Mac/For Windows, ensure that your docker daemon has access to atleast 4.5 GB of memory and 2 GB of swap size.

For Mac OSX see: https://docs.docker.com/docker-for-mac/#advanced
For Windows see: https://docs.docker.com/docker-for-windows/#advanced

The following commands need to be executed in the Kubernetes project root (typically https://github.com/kubernetes/kubernetes or a fork).

* Build Kubernetes:

NOTE: Ensure that you have all the pre-requirements met to build the kubernetes project.

```
make clean
make all
```

* Run the following commands to build the local hyperkube and push it to a Docker registry (you would need to have access to this registry)

```
export VERSION=<your-custom-version>
export REGISTRY=<your-docker-registry>
hack/dev-push-hyperkube.sh
```
For VERSION, we recommend that you provide a value which would help you identify the build of kubernetes. The value is VERSION will be used as tag for your custom hyperkube image.

(It's convenient to put these steps into a script.)

## In the AKS Engine repo

* Open the AKS Engine input JSON (e.g. a file from the examples directory) and add the following to the `orchestratorProfile` section:

```
"kubernetesConfig": {
    "customHyperkubeImage": "docker.io/<your-docker-registry>/hyperkube-amd64:<your-custom-version>"
}
```

* If the container registry is private, for example Azure Container Registry, then provide the name of the private Azure registry along with the custom hyperkube image like this:

```
"kubernetesConfig": {
    "customHyperkubeImage": "<your-private-registry>/hyperkube-amd64:<your-custom-version>",
    "privateAzureRegistryServer": "<your-private-registry>"
}
```
NOTE: Make sure the service principal provided to run `aks-engine deploy` has access to pull images from this private registry. https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-service-principal#use-an-existing-service-principal

* AKS Engine uses the `aks-ubuntu-1604` image for the master and Linux agents. This image contains prebuilt versions of kubelet and kubectl, thus the provision scripts will not copy the desired kubelet binary ( and kubectl ) from the custom hyperkube image. In order to get over this limitation it is required to specify `ubuntu` for the master and Linux agents in "masterProfile" and each agent pool under `agentPoolProfiles`.

```
"masterProfile": {
    "distro": "ubuntu"
}

"agentPoolProfiles": [
    {
        "distro": "ubuntu"
    }
]
```

* Run `aks-engine deploy` [as normal](../tutorials/deploy.md).
