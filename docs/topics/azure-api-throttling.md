# Azure API Throttling

## Overview

Azure has hard limits on the number of read and write requests against Azure APIs *per subscription, per region*. Running lots of clusters in a single subscription, or running a single large, dynamic cluster in a subscription can produce side effects that exceed the number of calls permitted within a given time window for a particular category of requests. See the following documents for more detail on Azure API throttling in general:

- https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling
- https://docs.microsoft.com/en-us/azure/virtual-machines/troubleshooting/troubleshooting-throttling-errors

## Running older versions of the Azure cloud provider runtime

Over time, the Azure cloud provider runtime has optimized its behaviors to reconcile Azure resource requests (network, compute, storage) with a minimum number of calls to the Azure APIs in order to prevent Azure API throttling. In practice, running versions of Kubernetes prior to the below list has known issues with generating excess Azure API requests for clusters with VMSS node pools:

- 1.15.12
- 1.16.9
- 1.17.5
- 1.18.2

We will demonstrate a real-time, in place remediation that updates the Azure cloud provider runtimes to a newer version. This requires that we stop all active Kubernetes components that interface with Azure APIs (e.g., kube-controller-manager or cloud-controller-manager, cluster-autoscaler), update those component specs so that they use a newer, optimized version of the Azure cloud provider, wait ~15-30 minutes, and then restart those components.

## Stop kube-controller-manager (or cloud-controller-manager) component

The easiest way to stop the `kube-controller-manager` (or `cloud-controller-manager` if you're using the Azure cloud provider controller as a separate component) is to move its pod spec out of the directory where the kubelet process continually processes control plane manifests (`/etc/kubernetes/manifests`). The running kubelet on the control plane VM will detect the missing controller-manager spec, and automatically delete the pod scheduled on that node. This has the effect of removing that particular Azure API client runtime from operation, reducing calls that count against subscription-wide regional request limits. Below we will demonstrate how to do that from a control plane VM.

In this example the `azureuser` system user has SSH access:

```
$ ssh-add -D && ssh-add <path to SSH private key>
$ ssh -A -i <path to SSH private key> azureuser@kubernetes-westus2-84461.westus2.cloudapp.azure.com

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

124 packages can be updated.
102 updates are security updates.

New release '18.04.4 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Wed Jun 24 17:38:37 2020 from 75.164.224.176
azureuser@k8s-master-31453872-0:~$
```

AKS Engine will set up our control plane VMs with a working `kubectl` context that can connect to the cluster. So, let's get the list of control plane VMs in the cluster (in this example we'll be updating a 1.15.7 cluster's cloud provider to 1.15.12):

```
azureuser@k8s-master-31453872-0:~$ kubectl get nodes | grep k8s-master
k8s-master-31453872-0           Ready    master   48m   v1.15.7
k8s-master-31453872-1           Ready    master   48m   v1.15.7
k8s-master-31453872-2           Ready    master   48m   v1.15.7
```

Now, we'll move the `kube-controller-manager` spec from the `/etc/kubernetes/manifests/` file on each node VM:

```
azureuser@k8s-master-31453872-0:~$ for control_plane_vm in $(kubectl get nodes | grep k8s-master | awk '{print $1}'); do ssh $control_plane_vm "sudo mv /etc/kubernetes/manifests/kube-controller-manager.yaml /opt/azure/kube-controller-manager.yaml"; done

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.
```

Within a few seconds, we should notice that there are no active `kube-controller-manager` pods running in the cluster:

```
azureuser@k8s-master-31453872-0:~$ k get pods --all-namespaces -o wide | grep kube-controller-manager | wc -l
       0
```

If we're running the Azure `cloud-controller-manager` component (`"useCloudControllerManager": true` in the cluster api model), then we would run this command instead:

```
azureuser@k8s-master-31453872-0:~$ for control_plane_vm in $(kubectl get nodes | grep k8s-master | awk '{print $1}'); do ssh $control_plane_vm "sudo mv /etc/kubernetes/manifests/cloud-controller-manager.yaml /opt/azure/cloud-controller-manager.yaml"; done

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.
```

At this point, the Azure controller-manager component will no longer be running on any control plane nodes in the cluster, and no Azure API requests will be generated from this source.

## Stop cluster-autoscaler component

If `cluster-autoscaler` is installed on the cluster, then we may want to stop, or slow down, the rate of Azure API requests that it is responsible for. If using the AKS Engine `cluster-autoscaler` addon, by default we configure the `kube-addon-manager` mode to `"EnsureExists"`, which means we can simply edit the deployment:

```
azureuser@k8s-master-31453872-0:~$ kubectl edit deployment cluster-autoscaler -n kube-system
...
```

Using the `kubectl` editor (`vim` by default), replace the existing replicas spec with `replicas: 0` (by default the `cluster-autoscaler` addon spec is set to `replicas: 1`).

Essentially, you want to see the following result:

```
azureuser@k8s-master-31453872-0:~$ kubectl get pods -l app=cluster-autoscaler -n kube-system
No resources found in kube-system namespace.
```

## Update controller-manager runtime

Depending on the time window (5 minute or one hour) of the API throttling violation, we may have to wait at least 30 minutes to reliably restart the control plane runtime without being throttled once again. If unsure about the nature of throttling violations, then waiting the full 30 minutes is the most conservative tactic to take.

So, assuming we've waited 30 minutes or so, let's update the controller-manager spec so that it refers to 1.15.12 instead of 1.15.7 (using our example), so that it gets the improvements to the Azure cloud provider runtime. Note that we're updating the spec at the `/opt/azure/` location that we moved it to earlier:

```
azureuser@k8s-master-31453872-0:~$ grep 1.15.7 /opt/azure/kube-controller-manager.yaml
      image: registry.k8s.io/hyperkube-amd64:v1.15.7
```

Let's update the spec on all control plane VMs:

```
azureuser@k8s-master-31453872-0:~$ for control_plane_vm in $(kubectl get nodes | grep k8s-master | awk '{print $1}'); do ssh $control_plane_vm "sudo sed -i 's|v1.15.7|v1.15.12|g' /opt/azure/kube-controller-manager.yaml"; done

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.
azureuser@k8s-master-31453872-0:~$ grep 1.15.12 /opt/azure/kube-controller-manager.yaml
      image: registry.k8s.io/hyperkube-amd64:v1.15.12
```

(Again, if you're using `cloud-controller-manager`, substitute the correct `cloud-controller-manager.yaml` file name.)

## Update cluster-autoscaler runtime

Now, if we're running the `cluster-autoscaler` addon on this cluster let's make sure it's using the latest release for 1.15 (at the time of this document, the latest `cluster-autoscaler` release for Kubernetes 1.15 is `1.15.6`):

```
azureuser@k8s-master-31453872-0:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: registry.k8s.io/cluster-autoscaler:v1.15.3
azureuser@k8s-master-31453872-0:~$ for control_plane_vm in $(kubectl get nodes | grep k8s-master | awk '{print $1}'); do ssh $control_plane_vm "sudo sed -i 's|v1.15.3|v1.15.6|g' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml"; done

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.
azureuser@k8s-master-31453872-0:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: registry.k8s.io/cluster-autoscaler:v1.15.6
```

The above validated that we *weren't* using the latest `cluster-autoscaler`, and so we changed the addon spec on each control plane VM in the `/etc/kubernetes/addons/` directory so that we would load 1.15.6 instead.

(Note: if managing a cluster-autoscaler installation outside of AKS Engine addons, for example using a Helm chart, do the equivalent version update steps as demonstrated above to the self-managed cluster-autoscaler spec, if appropriate.)

## Restart kube-controller-manager (or cloud-controller-manager) component

Again, we assume our ~30 minute cool down window has passed, and we're ready to restart the controller-manager control plane component by moving the pod spec back to the original `/etc/kubernetes/manifests/` location:

```
azureuser@k8s-master-31453872-0:~$ for control_plane_vm in $(kubectl get nodes | grep k8s-master | awk '{print $1}'); do ssh $control_plane_vm "sudo mv /opt/azure/kube-controller-manager.yaml /etc/kubernetes/manifests/kube-controller-manager.yaml"; done

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.

Authorized uses only. All activity may be monitored and reported.
```

## Verify controller-manager is re-loaded and running the desired, upgraded version

Now, we can verify that v1.15.12 is the running `controller-manager` version:

```
azureuser@k8s-master-31453872-0:~$ for pod in $(kubectl get pods -n kube-system -l component=kube-controller-manager | awk '{print $1}'); do kubectl logs $pod -n kube-system | grep "v1.15.12"; done
I0624 18:32:18.102332       1 controllermanager.go:164] Version: v1.15.12
I0624 18:32:45.174767       1 controllermanager.go:164] Version: v1.15.12
I0624 18:33:05.825467       1 controllermanager.go:164] Version: v1.15.12
```

## Reload cluster-autoscaler addon

If we're also using (and have stopped, and/or updated) `cluster-autoscaler`, remember to reload it, and/or verify that it's running the desired version. For example, restore the replica count to its original value:

```
azureuser@k8s-master-31453872-0:~$ kubectl edit deployment cluster-autoscaler -n kube-system
...
```

Now you can validate that it's running with the desired version. For example:

```
azureuser@k8s-master-31453872-0:~$ for pod in $(kubectl get pods -n kube-system -l app=cluster-autoscaler | awk '{print $1}'); do kubectl logs $pod -n kube-system | grep "1.15"; done
I0624 18:32:40.458155       1 main.go:354] Cluster Autoscaler 1.15.6
```

Again, there are a variety of `cluster-autoscaler` implementations that our cluster may be using. Two examples are the user-configurable cluster-autoscaler addon provided by AKS Engine and a user-defined Helm chart that is installed and maintained _after_ the cluster has been created by AKS Engine.
