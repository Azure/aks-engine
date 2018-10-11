# For Kubernetes Developers

If you're working on Kubernetes upstream, you can use AKS Engine to test your build of Kubernetes in the Azure environment.  The option that allows you to do this is `orchestratorProfile/kubernetesConfig/customHyperkubeImage`, which you should set to point to a Docker image containing your build of hyperkube.

The following instructions describe in more detail how to create the required Docker image and deploy it using AKS Engine (replace `dockerhubid` and `sometag` with your Docker Hub ID and a unique tag for your build):

## In the Kubernetes repo

* Build Kubernetes:

```
bash build/run.sh make cross KUBE_FASTBUILD=true ARCH=amd64
```

* Build a Docker image containing your custom build of hyperkube:

```
cd cluster/images/hyperkube
make VERSION=sometag
cd ../../..
```

* Push your Docker image to Docker Hub:

```
docker tag k8s.gcr.io/hyperkube-amd64:sometag dockerhubid/hyperkube-amd64:sometag
docker push dockerhubid/hyperkube-amd64:sometag
```

(It's convenient to put these steps into a script.)

## In the ACS repo

* Open the AKS Engine input JSON (e.g. a file from the examples directory) and add the following to the `orchestratorProfile` section:

```
"kubernetesConfig": {
    "customHyperkubeImage": "docker.io/dockerhubid/hyperkube-amd64:sometag"
}
```

* Run `./bin/aks-engine deploy --api-model the_json_file_you_just_edited.json ...` [as normal](deploy.md).
