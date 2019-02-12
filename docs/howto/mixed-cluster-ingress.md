# Using HTTP Ingress Routing in a Mixed Cluster

## Prerequisites

1. First, deploy a cluster with both Windows & Linux nodes. See the [Kubernetes Windows Walkthrough](../topics/windows.md) for a step by step example.
1. Install [Helm](http://helm.sh), the Kubernetes package manager

## Steps

### Configure Helm

```console
helm init --upgrade --node-selectors "beta.kubernetes.io/os=linux"
```

### Set up NGINX

```console
helm install --name nginx-ingress `
    --set controller.nodeSelector."beta\.kubernetes\.io\/os"=linux `
    --set defaultBackend.nodeSelector."beta\.kubernetes\.io\/os"=linux `
     --set rbac.create=true `
    stable/nginx-ingress
```

```console
helm install --name nginx-ingress \
    --set controller.nodeSelector."beta\.kubernetes\.io\/os"=linux \
    --set defaultBackend.nodeSelector."beta\.kubernetes\.io\/os"=linux \
    --set rbac.create=true \
    stable/nginx-ingress
```

This will return output like this

```text
NAME:   nginx-ingress
LAST DEPLOYED: Thu Aug 23 11:51:11 2018
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/Pod(related)
NAME                                            READY  STATUS             RESTARTS  AGE
nginx-ingress-controller-76c4d5cf59-zj7vb       0/1    ContainerCreating  0         2s
nginx-ingress-default-backend-69c6b65b46-d64nd  0/1    ContainerCreating  0         2s

==> v1/ConfigMap
NAME                      DATA  AGE
nginx-ingress-controller  1     2s

==> v1/Service
NAME                           TYPE          CLUSTER-IP    EXTERNAL-IP  PORT(S)                     AGE
nginx-ingress-controller       LoadBalancer  10.0.165.193  <pending>    80:32186/TCP,443:31319/TCP  2s
nginx-ingress-default-backend  ClusterIP     10.0.219.85   <none>       80/TCP                      2s

==> v1beta1/Deployment
NAME                           DESIRED  CURRENT  UP-TO-DATE  AVAILABLE  AGE
nginx-ingress-controller       1        1        1           0          2s
nginx-ingress-default-backend  1        1        1           0          2s

==> v1beta1/PodDisruptionBudget
NAME                           MIN AVAILABLE  MAX UNAVAILABLE  ALLOWED DISRUPTIONS  AGE
nginx-ingress-controller       1              N/A              0                    2s
nginx-ingress-default-backend  1              N/A              0                    2s


NOTES:
The nginx-ingress controller has been installed.
It may take a few minutes for the LoadBalancer IP to be available.
You can watch the status by running 'kubectl --namespace default get services -o wide -w nginx-ingress-controller'

An example Ingress that makes use of the controller:

  apiVersion: extensions/v1beta1
  kind: Ingress
  metadata:
    annotations:
      kubernetes.io/ingress.class: nginx
    name: example
    namespace: foo
  spec:
    rules:
      - host: www.example.com
        http:
          paths:
            - backend:
                serviceName: exampleService
                servicePort: 80
              path: /
    # This section is only required if TLS is to be enabled for the Ingress
    tls:
        - hosts:
            - www.example.com
          secretName: example-tls

If TLS is enabled for the Ingress, a Secret containing the certificate and key must also be provided:

  apiVersion: v1
  kind: Secret
  metadata:
    name: example-tls
    namespace: foo
  data:
    tls.crt: <base64 encoded cert>
    tls.key: <base64 encoded key>
  type: kubernetes.io/tls
```

### Create a web server and service

Copy the YAML below into a file called `iis.yaml`.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: iis-1803
  labels:
    app: iis-1803
spec:
  replicas: 1
  template:
    metadata:
      name: iis-1803
      labels:
        app: iis-1803
    spec:
      containers:
      - name: iis
        image: microsoft/iis:windowsservercore-1803
        ports:
          - containerPort: 80
      nodeSelector:
        "beta.kubernetes.io/os": windows
  selector:
    matchLabels:
      app: iis-1803
---
apiVersion: v1
kind: Service
metadata:
  name: iis
spec:
  type: ClusterIP
  ports:
  - protocol: TCP
    port: 80
  selector:
    app: iis-1803
```

If you're using Windows Server version 1803, you can use it as-is. If you're using 1709, replace all the instances of 1803 with 1709.

Now, run `kubectl create -f iis.yaml`

Check that the web server is running with `kubectl get pod`, and look for `iis-1803-...`.

```console
$ kubectl get pods
NAME                                             READY     STATUS    RESTARTS   AGE
iis-1803-8b7fdd569-nvzx8                         1/1       Running   0          4m
nginx-ingress-controller-57cbbfcb7c-fgtdn        1/1       Running   0          17m
nginx-ingress-default-backend-69c6b65b46-zjc2c   1/1       Running   0          17m
```

If it's not ready, check again in a bit. The first time this is run, the container image pull may take up to 20 minutes. `kubectl describe pod ...` will give more details on progress.

It's also good to confirm the service is up with `kubectl describe svc iis`, and that there is at least one endpoint listed.

```console
$ kubectl describe svc iis
Name:              iis
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          app=iis-1803
Type:              ClusterIP
IP:                10.0.7.142
Port:              <unset>  80/TCP
TargetPort:        80/TCP
Endpoints:         10.240.0.143:80
Session Affinity:  None
Events:            <none>
```

### Create the ingress rule

Now that the pod and service are running, it's time to create the ingress.

Copy this YAML to a file called `ingress.yaml`

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: nginx
  name: iis-ingress
  namespace: default
spec:
  rules:
    - host: test.ogfg.link
      http:
        paths:
          - backend:
              serviceName: iis
              servicePort: 80
            path: /
```

Then edit to your needs. Be sure to set `host` to a DNS name that you are able to manage. Create the ingress rule with `kubectl create -f ingress.yaml`

Now, it's a good time to test the ingress rule from inside the cluster. Run `kubectl get svc nginx-ingress-controller` and look for `CLUSTER-IP`

```console
$ kubectl get svc nginx-ingress-controller
NAME                       TYPE           CLUSTER-IP    EXTERNAL-IP     PORT(S)                      AGE
nginx-ingress-controller   LoadBalancer   10.0.71.219   13.77.176.117   80:32040/TCP,443:30669/TCP   31m
```

SSH to a Linux node in the cluster, and run `curl -H 'host:<hostname for ingress rule>' http://<CLUSTER-IP>`

```console
$ curl -H 'Host: test.ogfg.link' http://10.0.71.219
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1" />
<title>IIS Windows Server</title>
```

This should return the web page contents, not an error.

### Update DNS

Now, you need to make sure that the external IP for the `nginx-ingress-controller` is registered in your DNS zone.

First, get the external IP for the ingress controller with `kubectl describe svc nginx-ingress-controller`

```console
$ kubectl describe svc nginx-ingress-controller
Name:                     nginx-ingress-controller
Namespace:                default
Labels:                   app=nginx-ingress
                          chart=nginx-ingress-0.11.1
                          component=controller
                          heritage=Tiller
                          release=nginx-ingress
Annotations:              <none>
Selector:                 app=nginx-ingress,component=controller,release=nginx-ingress
Type:                     LoadBalancer
IP:                       10.0.71.219
LoadBalancer Ingress:     13.77.176.117
Port:                     http  80/TCP
TargetPort:               80/TCP
NodePort:                 http  32040/TCP
Endpoints:                10.240.0.56:80
Port:                     https  443/TCP
TargetPort:               443/TCP
NodePort:                 https  30669/TCP
Endpoints:                10.240.0.56:443
Session Affinity:         None
External Traffic Policy:  Cluster
Events:
  Type    Reason                Age   From                Message
  ----    ------                ----  ----                -------
  Normal  EnsuringLoadBalancer  24m   service-controller  Ensuring load balancer
  Normal  EnsuredLoadBalancer   23m   service-controller  Ensured load balancer
  ```

`LoadBalancer Ingress` is the external IP. Create a DNS `A` record with a matching hostname (or wildcard) and that external IP.

If you're using Azure DNS, then you can set this in your DNS zone with:

`az network dns record-set a add-record -n test -g <resource group containing dns zone> --zone-name <DNS zone> --ipv4-address <IP of ingress service>`

Now, you should be able to access your Windows web server running at `http://hostname`
