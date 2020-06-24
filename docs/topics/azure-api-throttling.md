# Azure API Throttling

## Overview

Azure has hard limits on the number of read and write requests against Azure APIs *per subscription, per region*. Running lots of clusters in a single subscription, or running a single large, dynamic cluster in a subscription can produce side effects that exceed the number of calls permitted within a given time window for a particular category of requests. See the following documents for more detail on Azure API throttling in general:

- https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling
- https://docs.microsoft.com/en-us/azure/virtual-machines/troubleshooting/troubleshooting-throttling-errors

## Running older versions of the Azure cloudprovider runtime

Over time, the Azure cloudprovider runtime has optimized its behaviors to reconcile Azure resource requests (network, compute, storage) with a minimum number of calls to the Azure APIs in order to prevent Azure API throttling. In practice, running versions of Kubernetes prior to the below list has known issues with generating excess Azure API requests for clusters with VMSS node pools:

- 1.15.12
- 1.16.9
- 1.17.5
- 1.18.2

We will demonstrate a real-time, in place remediation that updates the Azure cloudprovider runtimes to a newer version. This requires that we stop all active Kubernetes components that interface with Azure APIs (e.g., controller-manager, cluster-autoscaler), update those component specs so that they use a newer, optimized version of the Azure cloudprovider, wait ~15-30 minutes, and then restart those components.

## Stop controller-manager

The easiest way to stop the controller-manager is to stop the kubelet systemd job on every control plane VM. Assuming that they all share the same SSH keypair, and that the `azureuser` system user permits that SSH keypair for login:

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

AKS Engine will set up your control plane VMs with a working `kubectl` context that can connect to the cluster. So, let's get the list of control plane VMs in the cluster (in this example we'll be updating a 1.15.7 cluster's cloudprovider to 1.15.12):

```
azureuser@k8s-master-31453872-0:~$ kubectl get nodes | grep k8s-master
k8s-master-31453872-0           Ready    master   48m   v1.15.7
k8s-master-31453872-1           Ready    master   48m   v1.15.7
k8s-master-31453872-2           Ready    master   48m   v1.15.7
```

Now, we'll stop the `kubelet` systemd job on each one. First, from the VM we're already logged onto:

```
azureuser@k8s-master-31453872-0:~$ sudo systemctl stop kubelet
azureuser@k8s-master-31453872-0:~$ sudo systemctl status kubelet
● kubelet.service - Kubelet
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: enabled)
   Active: inactive (dead) since Wed 2020-06-24 17:51:17 UTC; 8s ago
 Main PID: 5290 (code=exited, status=0/SUCCESS)

Jun 24 17:14:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:14:56.972492    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:19:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:19:56.973082    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:24:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:24:56.973619    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:29:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:29:56.981518    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:34:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:34:56.982118    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:39:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:39:56.982497    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:44:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:44:56.983098    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:49:56 k8s-master-31453872-0 kubelet[5290]: I0624 17:49:56.984258    5290 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:51:17 k8s-master-31453872-0 systemd[1]: Stopping Kubelet...
Jun 24 17:51:17 k8s-master-31453872-0 systemd[1]: Stopped Kubelet.
```

Now, from that VM, because we've forwarded our keychain and we're using a common keypair for interactive login, we can log into the remaining control plane VMs and stop kubelet:

```
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-1
The authenticity of host 'k8s-master-31453872-1 (10.239.255.240)' can't be established.
ECDSA key fingerprint is SHA256:wy4cLGaj0bc/KnO/S1huvl/elDNcucAycVmTXg8SFWk.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'k8s-master-31453872-1,10.239.255.240' (ECDSA) to the list of known hosts.

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

0 packages can be updated.
0 updates are security updates.



The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

azureuser@k8s-master-31453872-1:~$ sudo systemctl stop kubelet
azureuser@k8s-master-31453872-1:~$ sudo systemctl status kubelet
● kubelet.service - Kubelet
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: enabled)
   Active: inactive (dead) since Wed 2020-06-24 17:55:53 UTC; 4s ago
 Main PID: 4081 (code=exited, status=0/SUCCESS)

Jun 24 17:20:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:20:22.265860    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:25:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:25:22.266375    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:30:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:30:22.266763    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:35:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:35:22.267346    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:40:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:40:22.267762    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:45:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:45:22.268173    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:50:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:50:22.269530    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:55:22 k8s-master-31453872-1 kubelet[4081]: I0624 17:55:22.269912    4081 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:55:53 k8s-master-31453872-1 systemd[1]: Stopping Kubelet...
Jun 24 17:55:53 k8s-master-31453872-1 systemd[1]: Stopped Kubelet.
azureuser@k8s-master-31453872-1:~$ exit
logout
Connection to k8s-master-31453872-1 closed.
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-2
The authenticity of host 'k8s-master-31453872-2 (10.239.255.241)' can't be established.
ECDSA key fingerprint is SHA256:5JC6jfj6vImnFTYQx9LLOKMEkh+tQuWM8GkPWmeW9lU.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'k8s-master-31453872-2,10.239.255.241' (ECDSA) to the list of known hosts.

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

32 packages can be updated.
0 updates are security updates.


*** System restart required ***

The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

azureuser@k8s-master-31453872-2:~$ sudo systemctl stop kubelet
azureuser@k8s-master-31453872-2:~$ sudo systemctl status kubelet
● kubelet.service - Kubelet
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: enabled)
   Active: inactive (dead) since Wed 2020-06-24 17:56:30 UTC; 5s ago
 Main PID: 5243 (code=exited, status=0/SUCCESS)

Jun 24 17:19:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:19:50.498593    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:24:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:24:50.499897    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:29:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:29:50.500582    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:34:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:34:50.501119    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:39:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:39:50.501761    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:44:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:44:50.505077    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:49:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:49:50.505539    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:54:50 k8s-master-31453872-2 kubelet[5243]: I0624 17:54:50.505996    5243 container_manager_linux.go:457] [ContainerManager]: Discovered runtime cgroups name: /system.slice/docker.service
Jun 24 17:56:30 k8s-master-31453872-2 systemd[1]: Stopping Kubelet...
Jun 24 17:56:30 k8s-master-31453872-2 systemd[1]: Stopped Kubelet.
azureuser@k8s-master-31453872-2:~$ exit
logout
Connection to k8s-master-31453872-2 closed.
```

## Update controller-manager runtime

Depending on the time window (5 minute or one hour) of the API throttling violation, we may have to wait at least 30 minutes to reliably restart the control plane runtime without being throttled once again. If unsure about the nature of throttling violations, then waiting the full 30 minutes is the most conservative tactic to take.

So, assuming we've waited 30 minutes or so, let's update the controller-manager spec so that it refers to 1.15.12 instead of 1.15.7 (using our example), so that it gets the improvements to the Azure cloudprovider runtime:

```
azureuser@k8s-master-31453872-0:~$ grep 1.15.7 /etc/kubernetes/manifests/kube-controller-manager.yaml
      image: k8s.gcr.io/hyperkube-amd64:v1.15.7
azureuser@k8s-master-31453872-0:~$ sudo sed -i "s|v1.15.7|v1.15.12|g" /etc/kubernetes/manifests/kube-controller-manager.yaml
azureuser@k8s-master-31453872-0:~$ grep 1.15.12 /etc/kubernetes/manifests/kube-controller-manager.yaml
      image: k8s.gcr.io/hyperkube-amd64:v1.15.12
```

And then on the other control plane VMs:

```
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-1

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

0 packages can be updated.
0 updates are security updates.

New release '18.04.4 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Wed Jun 24 17:54:46 2020 from 10.239.255.239
azureuser@k8s-master-31453872-1:~$ grep 1.15.7 /etc/kubernetes/manifests/kube-controller-manager.yaml
      image: k8s.gcr.io/hyperkube-amd64:v1.15.7
azureuser@k8s-master-31453872-1:~$ sudo sed -i "s|v1.15.7|v1.15.12|g" /etc/kubernetes/manifests/kube-controller-manager.yaml
azureuser@k8s-master-31453872-1:~$ grep 1.15.12 /etc/kubernetes/manifests/kube-controller-manager.yaml
      image: k8s.gcr.io/hyperkube-amd64:v1.15.12
azureuser@k8s-master-31453872-1:~$ exit
logout
Connection to k8s-master-31453872-1 closed.
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-2

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

32 packages can be updated.
0 updates are security updates.

New release '18.04.4 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


*** System restart required ***
Last login: Wed Jun 24 17:56:18 2020 from 10.239.255.239
azureuser@k8s-master-31453872-2:~$ grep 1.15.7 /etc/kubernetes/manifests/kube-controller-manager.yaml
      image: k8s.gcr.io/hyperkube-amd64:v1.15.7
azureuser@k8s-master-31453872-2:~$ sudo sed -i "s|v1.15.7|v1.15.12|g" /etc/kubernetes/manifests/kube-controller-manager.yaml
azureuser@k8s-master-31453872-2:~$ grep 1.15.12 /etc/kubernetes/manifests/kube-controller-manager.yaml
      image: k8s.gcr.io/hyperkube-amd64:v1.15.12
azureuser@k8s-master-31453872-2:~$ exit
logout
Connection to k8s-master-31453872-2 closed.
```

## Update cluster-autoscaler runtime

Now, are we running the `cluster-autoscaler` addon on this cluster? If so, let's make sure it's using the latest release for 1.15 (at the time of this document, the latest `cluster-autoscaler` release for Kubernetes 1.15 is `1.15.5`):

```
azureuser@k8s-master-31453872-0:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: k8s.gcr.io/cluster-autoscaler:v1.15.3
azureuser@k8s-master-31453872-0:~$ sudo sed -i "s|v1.15.3|v1.15.5|g" /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
azureuser@k8s-master-31453872-0:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: k8s.gcr.io/cluster-autoscaler:v1.15.5
```

The above validated that we *weren't* using the latest `cluster-autoscaler`, and so we changed the addon spec so that we would load 1.15.5 instead.

And, again, we do the same to the remaining control plane VMs:

```
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-1

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

0 packages can be updated.
0 updates are security updates.

New release '18.04.4 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Wed Jun 24 18:20:58 2020 from 10.239.255.239
azureuser@k8s-master-31453872-1:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: k8s.gcr.io/cluster-autoscaler:v1.15.3
azureuser@k8s-master-31453872-1:~$ sudo sed -i "s|v1.15.3|v1.15.5|g" /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
azureuser@k8s-master-31453872-1:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: k8s.gcr.io/cluster-autoscaler:v1.15.5
azureuser@k8s-master-31453872-1:~$ exit
logout
Connection to k8s-master-31453872-1 closed.
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-2

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

32 packages can be updated.
0 updates are security updates.

New release '18.04.4 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


*** System restart required ***
Last login: Wed Jun 24 18:21:29 2020 from 10.239.255.239
azureuser@k8s-master-31453872-2:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: k8s.gcr.io/cluster-autoscaler:v1.15.3
azureuser@k8s-master-31453872-2:~$ sudo sed -i "s|v1.15.3|v1.15.5|g" /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
azureuser@k8s-master-31453872-2:~$ grep 'cluster-autoscaler:v' /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
      - image: k8s.gcr.io/cluster-autoscaler:v1.15.5
azureuser@k8s-master-31453872-2:~$ exit
logout
Connection to k8s-master-31453872-2 closed.
```

(Note: if managing a cluster-autoscaler installation outside of AKS Engine addons, for example using a Helm chart, do the equivalent version update steps as demonstrated above to the self-managed cluster-autoscaler spec.)

## Restart control plane runtime

Now, we should restart kubelet on all control plane VMs:

```
azureuser@k8s-master-31453872-0:~$ sudo systemctl start kubelet
azureuser@k8s-master-31453872-0:~$ systemctl status kubelet
● kubelet.service - Kubelet
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: enabled)
   Active: active (running) since Wed 2020-06-24 18:32:06 UTC; 9s ago
  Process: 36752 ExecStartPre=/sbin/iptables -t nat --numeric --list (code=exited, status=0/SUCCESS)
  Process: 36749 ExecStartPre=/sbin/ebtables -t nat --list (code=exited, status=0/SUCCESS)
  Process: 36743 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh3=16384; fi (code=exited, status=0/SUCCESS)
  Process: 36738 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh2=8192; fi (code=exited, status=0/SUCCESS)
  Process: 36733 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh1=4096; fi (code=exited, status=0/SUCCESS)
  Process: 36730 ExecStartPre=/sbin/sysctl -w net.core.message_burst=80 (code=exited, status=0/SUCCESS)
  Process: 36727 ExecStartPre=/sbin/sysctl -w net.core.message_cost=40 (code=exited, status=0/SUCCESS)
  Process: 36724 ExecStartPre=/sbin/sysctl -w net.ipv4.tcp_max_syn_backlog=16384 (code=exited, status=0/SUCCESS)
  Process: 36721 ExecStartPre=/sbin/sysctl -w net.core.somaxconn=16384 (code=exited, status=0/SUCCESS)
  Process: 36717 ExecStartPre=/sbin/sysctl -w net.ipv4.tcp_retries2=8 (code=exited, status=0/SUCCESS)
  Process: 36714 ExecStartPre=/bin/mount --make-shared /var/lib/kubelet (code=exited, status=0/SUCCESS)
  Process: 36703 ExecStartPre=/bin/bash -c if [ $(mount | grep "/var/lib/kubelet" | wc -l) -le 0 ] ; then /bin/mount --bind /var/lib/kubelet /var/lib/kubelet ; fi (code=exited, status=0/SUCCESS)
  Process: 36702 ExecStartPre=/bin/mkdir -p /var/lib/cni (code=exited, status=0/SUCCESS)
  Process: 36698 ExecStartPre=/bin/mkdir -p /var/lib/kubelet (code=exited, status=0/SUCCESS)
  Process: 36682 ExecStartPre=/bin/bash /opt/azure/containers/kubelet.sh (code=exited, status=0/SUCCESS)
 Main PID: 36756 (kubelet)
    Tasks: 16
   Memory: 40.1M
      CPU: 8.066s
   CGroup: /system.slice/kubelet.service
           └─36756 /usr/local/bin/kubelet --enable-server --node-labels=kubernetes.azure.com/role=master,kubernetes.io/role=master,node-role.kubernetes.io/master=,kubernetes.azure.com/cluster=kubernetes-wes
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-1

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

0 packages can be updated.
0 updates are security updates.

New release '18.04.4 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Wed Jun 24 18:27:49 2020 from 10.239.255.239
azureuser@k8s-master-31453872-1:~$ sudo systemctl start kubelet
azureuser@k8s-master-31453872-1:~$ systemctl status kubelet
● kubelet.service - Kubelet
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: enabled)
   Active: active (running) since Wed 2020-06-24 18:32:31 UTC; 5s ago
  Process: 36677 ExecStartPre=/sbin/iptables -t nat --numeric --list (code=exited, status=0/SUCCESS)
  Process: 36674 ExecStartPre=/sbin/ebtables -t nat --list (code=exited, status=0/SUCCESS)
  Process: 36669 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh3=16384; fi (code=exited, status=0/SUCCESS)
  Process: 36664 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh2=8192; fi (code=exited, status=0/SUCCESS)
  Process: 36659 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh1=4096; fi (code=exited, status=0/SUCCESS)
  Process: 36657 ExecStartPre=/sbin/sysctl -w net.core.message_burst=80 (code=exited, status=0/SUCCESS)
  Process: 36654 ExecStartPre=/sbin/sysctl -w net.core.message_cost=40 (code=exited, status=0/SUCCESS)
  Process: 36652 ExecStartPre=/sbin/sysctl -w net.ipv4.tcp_max_syn_backlog=16384 (code=exited, status=0/SUCCESS)
  Process: 36648 ExecStartPre=/sbin/sysctl -w net.core.somaxconn=16384 (code=exited, status=0/SUCCESS)
  Process: 36645 ExecStartPre=/sbin/sysctl -w net.ipv4.tcp_retries2=8 (code=exited, status=0/SUCCESS)
  Process: 36643 ExecStartPre=/bin/mount --make-shared /var/lib/kubelet (code=exited, status=0/SUCCESS)
  Process: 36636 ExecStartPre=/bin/bash -c if [ $(mount | grep "/var/lib/kubelet" | wc -l) -le 0 ] ; then /bin/mount --bind /var/lib/kubelet /var/lib/kubelet ; fi (code=exited, status=0/SUCCESS)
  Process: 36632 ExecStartPre=/bin/mkdir -p /var/lib/cni (code=exited, status=0/SUCCESS)
  Process: 36630 ExecStartPre=/bin/mkdir -p /var/lib/kubelet (code=exited, status=0/SUCCESS)
  Process: 36614 ExecStartPre=/bin/bash /opt/azure/containers/kubelet.sh (code=exited, status=0/SUCCESS)
 Main PID: 36681 (kubelet)
    Tasks: 22
   Memory: 49.9M
      CPU: 1.193s
   CGroup: /system.slice/kubelet.service
           ├─36681 /usr/local/bin/kubelet --enable-server --node-labels=kubernetes.azure.com/role=master,kubernetes.io/role=master,node-role.kubernetes.io/master=,kubernetes.azure.com/cluster=kubernetes-wes
           └─36980 /opt/cni/bin/azure-vnet-telemetry -d /opt/cni/bin
azureuser@k8s-master-31453872-1:~$ exit
logout
Connection to k8s-master-31453872-1 closed.
azureuser@k8s-master-31453872-0:~$ ssh k8s-master-31453872-2

Authorized uses only. All activity may be monitored and reported.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.15.0-1064-azure x86_64)

32 packages can be updated.
0 updates are security updates.

New release '18.04.4 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


*** System restart required ***
Last login: Wed Jun 24 18:28:38 2020 from 10.239.255.239
azureuser@k8s-master-31453872-2:~$ sudo systemctl start kubelet
azureuser@k8s-master-31453872-2:~$ systemctl status kubelet
● kubelet.service - Kubelet
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: enabled)
   Active: active (running) since Wed 2020-06-24 18:32:45 UTC; 5s ago
  Process: 73216 ExecStartPre=/sbin/iptables -t nat --numeric --list (code=exited, status=0/SUCCESS)
  Process: 73213 ExecStartPre=/sbin/ebtables -t nat --list (code=exited, status=0/SUCCESS)
  Process: 73209 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh3=16384; fi (code=exited, status=0/SUCCESS)
  Process: 73204 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh2=8192; fi (code=exited, status=0/SUCCESS)
  Process: 73198 ExecStartPre=/bin/bash -c if [ $(nproc) -gt 8 ]; then /sbin/sysctl -w net.ipv4.neigh.default.gc_thresh1=4096; fi (code=exited, status=0/SUCCESS)
  Process: 73196 ExecStartPre=/sbin/sysctl -w net.core.message_burst=80 (code=exited, status=0/SUCCESS)
  Process: 73194 ExecStartPre=/sbin/sysctl -w net.core.message_cost=40 (code=exited, status=0/SUCCESS)
  Process: 73189 ExecStartPre=/sbin/sysctl -w net.ipv4.tcp_max_syn_backlog=16384 (code=exited, status=0/SUCCESS)
  Process: 73188 ExecStartPre=/sbin/sysctl -w net.core.somaxconn=16384 (code=exited, status=0/SUCCESS)
  Process: 73185 ExecStartPre=/sbin/sysctl -w net.ipv4.tcp_retries2=8 (code=exited, status=0/SUCCESS)
  Process: 73182 ExecStartPre=/bin/mount --make-shared /var/lib/kubelet (code=exited, status=0/SUCCESS)
  Process: 73173 ExecStartPre=/bin/bash -c if [ $(mount | grep "/var/lib/kubelet" | wc -l) -le 0 ] ; then /bin/mount --bind /var/lib/kubelet /var/lib/kubelet ; fi (code=exited, status=0/SUCCESS)
  Process: 73170 ExecStartPre=/bin/mkdir -p /var/lib/cni (code=exited, status=0/SUCCESS)
  Process: 73167 ExecStartPre=/bin/mkdir -p /var/lib/kubelet (code=exited, status=0/SUCCESS)
  Process: 73151 ExecStartPre=/bin/bash /opt/azure/containers/kubelet.sh (code=exited, status=0/SUCCESS)
 Main PID: 73220 (kubelet)
    Tasks: 16
   Memory: 39.4M
      CPU: 1.030s
   CGroup: /system.slice/kubelet.service
           └─73220 /usr/local/bin/kubelet --enable-server --node-labels=kubernetes.azure.com/role=master,kubernetes.io/role=master,node-role.kubernetes.io/master=,kubernetes.azure.com/cluster=kubernetes-wes

Jun 24 18:32:47 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:47.332336   73220 kuberuntime_manager.go:404] No sandbox for pod "kube-controller-manager-k8s-master-31453872-2_kube-system(6c557ed8279f0808
Jun 24 18:32:47 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:47.513965   73220 operation_generator.go:713] MountVolume.SetUp succeeded for volume "azure-ip-masq-agent-config-volume" (UniqueName: "kuber
Jun 24 18:32:48 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:48.226853   73220 kubelet_pods.go:1090] Killing unwanted pod "kube-controller-manager-k8s-master-31453872-2"
Jun 24 18:32:48 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:48.230587   73220 kuberuntime_container.go:581] Killing container "docker://84f75fc07d76a4c42a3f27c5c30967e57a7dc863a04c43ec7ea7163158bc302a
Jun 24 18:32:48 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:48.393892   73220 provider.go:124] Refreshing cache for provider: *credentialprovider.defaultDockerConfigProvider
Jun 24 18:32:48 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:48.394917   73220 provider.go:124] Refreshing cache for provider: *azure.acrProvider
Jun 24 18:32:48 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:48.478356   73220 kubelet.go:1933] SyncLoop (PLEG): "kube-controller-manager-k8s-master-31453872-2_kube-system(6c557ed8279f08082ce8545b701ff
Jun 24 18:32:49 k8s-master-31453872-2 kubelet[73220]: W0624 18:32:49.489357   73220 pod_container_deletor.go:75] Container "97ef204fed2c5dc700b248203c310ebce748dc03703ccf535b7288e992b9bdf7" not found in pod
Jun 24 18:32:50 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:50.221701   73220 kubelet_pods.go:1090] Killing unwanted pod "kube-controller-manager-k8s-master-31453872-2"
Jun 24 18:32:50 k8s-master-31453872-2 kubelet[73220]: I0624 18:32:50.225143   73220 kuberuntime_container.go:581] Killing container "docker://84f75fc07d76a4c42a3f27c5c30967e57a7dc863a04c43ec7ea7163158bc302a
azureuser@k8s-master-31453872-2:~$ exit
logout
Connection to k8s-master-31453872-2 closed.
```

## Verify controller-manager is running the desired, upgraded version

Now, we can verify that v1.15.12 is the running `controller-manager` version:

```
azureuser@k8s-master-31453872-0:~$ for pod in $(kubectl get pods -n kube-system | grep kube-controller-manager | awk '{print $1}'); do kubectl logs $pod -n kube-system | grep "v1.15.12"; done
I0624 18:32:18.102332       1 controllermanager.go:164] Version: v1.15.12
I0624 18:32:45.174767       1 controllermanager.go:164] Version: v1.15.12
I0624 18:33:05.825467       1 controllermanager.go:164] Version: v1.15.12
```

## Reload cluster-autoscaler addon

If also running (and have updated) `cluster-autoscaler`, note that the prior version is still running:

```
azureuser@k8s-master-31453872-0:~$ for pod in $(kubectl get pods -n kube-system | grep cluster-autoscaler | awk '{print $1}'); do kubectl logs $pod -n kube-system | grep "1.15"; done
I0624 18:32:40.458155       1 main.go:354] Cluster Autoscaler 1.15.3
```

That's because the `kube-addon-manager` and `cluster-autoscaler` containers continued running all the while, they were not affected by stopping the systemd kubelet service. We want to stop all the running `kube-addon-manager` pods to force a reload of the `/etc/kubernetes/addons` directory:

```
azureuser@k8s-master-31453872-0:~$ for pod in $(kubectl get pods -n kube-system | grep kube-addon-manager | awk '{print $1}'); do kubectl delete pod $pod -n kube-system; done
pod "kube-addon-manager-k8s-master-31453872-0" deleted
pod "kube-addon-manager-k8s-master-31453872-1" deleted
pod "kube-addon-manager-k8s-master-31453872-2" deleted
```

And then delete the loaded cluster-autoscaler *deployment* to force `kube-addon-manager` to load a new one using the new spec:

```
azureuser@k8s-master-31453872-0:~$ kubectl delete deployment cluster-autoscaler -n kube-system
deployment.extensions "cluster-autoscaler" deleted
```

Again, the above restart process assumes we're running the AKS Engine-provided cluster-autoscaler. Reloading a new cluster-autoscaler runtime using helm would require a different process, but the same idea.

## Verify cluster-autoscaler is running the desired, upgraded version

Once the new cluster-autoscaler deployment has replicated a pod, and the new container image has been downloaded and run, we should be able to verify the new version:

```
azureuser@k8s-master-31453872-0:~$ for pod in $(kubectl get pods -n kube-system | grep cluster-autoscaler | awk '{print $1}'); do kubectl logs $pod -n kube-system | grep "1.15"; done
I0624 18:45:10.068529       1 main.go:354] Cluster Autoscaler 1.15.5
```
