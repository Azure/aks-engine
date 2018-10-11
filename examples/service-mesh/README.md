# Kubernetes - Service Mesh

There are numerous implementations of a service mesh which integrate with kubernetes such as Istio, [Linkerd](http://linkerd.io), and [Conduit](https://conduit.io/).  [This is one blog post](https://medium.com/microservices-in-practice/service-mesh-for-microservices-2953109a3c9a) which explains some fundamentals behind what it is and why to use it.

Some service mesh implementations **may** benefit from or require additional [customizations to the kubernetes cluster itself](https://github.com/Azure/aks-engine/blob/master/docs/clusterdefinition.md).

## Istio

The `istio.json` file in this directory enables the kubernetes API server options to support automatic sidecar injection using [Isitio](https://istio.io/).  If automatic sidecar injection isn't enabled, then all services must then manually inject the sidecar configuration into every deployment, every time.

The main changes this configuration makes is adding these flags to the apiserver `MutatingAdmissionWebhook,ValidatingAdmissionWebhook`.

### Installation

#### Create Azure Resources

1. Create Resource Group

    ```
    az group create --name "<resourceGroupName>" --location "eastus"
    ```

2. Create Service Principal

    ```
    az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/<subscriptionId>/resourceGroups/<resourceGroupName>"
    ```

#### Update istio.json

3. Ensure `orchestratorRelease` is 1.9+. 
4. Update `--admission-control` to include `MutatingAdmissionWebhook,ValidatingAdmissionWebhook`

    **Note**: admission-controls need to be entered in the order defined on the kubernetes [docs](https://kubernetes.io/docs/admin/admission-controllers/#is-there-a-recommended-set-of-admission-controllers-to-use).

    Your updates should look like this.
    ```
    "orchestratorProfile": {
        "orchestratorType": "Kubernetes",
        "orchestratorRelease": "1.9",
        "kubernetesConfig": {
            "apiServerConfig": {
            "--admission-control":  "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,DenyEscalatingExec,AlwaysPullImages,ValidatingAdmissionWebhook,ResourceQuota",
            }
        }
    }
    ```

4. Add Service Principal ID and Secret
    ```
    "servicePrincipalProfile": {
        "clientId": "<Insert Service Principal Client ID>",
        "secret": "<Insert Service Principal Client Secret>"
    }
    ```


### Post installation

Once the template has been successfully deployed, then Istio can be installed via either:

1. Manual - follow the website [Installation steps](https://istio.io/docs/setup/kubernetes/quick-start.html#installation-steps).
1. Helm Chart - is maintained in the Istio repository itself (no longer hub.kubeapps.com).  [See these instructions on the Istio website](https://istio.io/docs/setup/kubernetes/helm.html).

> Note: So far it seems the manual steps are more well maintained and up-to-date than the helm chart.

After Istio has been installed, consider [walking through the various Tasks](https://istio.io/docs/tasks/) which use the Book info example application.