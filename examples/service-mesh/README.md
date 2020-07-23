# Kubernetes - Service Mesh

There are numerous implementations of a service mesh which integrate with Kubernetes such as Istio, [Linkerd](http://linkerd.io), and [Conduit](https://conduit.io/). This blog [post](https://medium.com/microservices-in-practice/service-mesh-for-microservices-2953109a3c9a) explains what a service mesh is and why to use one.

This page details the customizations required to successfully deploy some of these service mesh implementations.

## Istio

_The steps below were validated using AKS Engine v0.53.0 and Kubernetes v1.17_

### Secret Discovery Service

To accomodate to the changes in [Secret Discovery Service](https://istio.io/latest/blog/2019/trustworthy-jwt-sds/) since Istio 1.3,
please refer to sample apimodel [istio.json](./istio.json)

This sample shows the extra Kubernetes API server flags that are required to enable `Service Account Token Volume Projection`as indicated
[here](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#service-account-token-volume-projection).

You may have to adjust `--service-account-api-audiences` and `--service-account-issuer` to your specific use case.

```json
{
    "kubernetesConfig": {
        "apiServerConfig": {
            "--feature-gates": "TokenRequestProjection=true",
            "--service-account-api-audiences": "api,istio-ca",
            "--service-account-issuer": "kubernetes.default.svc",
            "--service-account-signing-key-file": "/etc/kubernetes/certs/apiserver.key"
        }
    }
}
```
