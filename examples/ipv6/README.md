# AKS Engine - IPv6 clusters

## Overview

AKS Engine enables you to create IPv6 Kubernetes clusters on Microsoft Azure.

- IPv6 support is available for Kubernetes version 1.18.0 and later on Azure.
- IPv6 support is in beta as of Kubernetes version 1.18 in Kubernetes community.

In order to create IPv6 enabled Azure virtual networks and use standard loadbalancer with IPv6 you must first configure your subscription [as follows](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-ipv4-ipv6-dual-stack-cli#prerequisites).

This example shows you how to configure a IPv6 cluster:

1. **kubernetes.json** - deploying and using [Kubernetes](kubernetes.json).

Things to try out after the cluster is deployed -

- Nodes are Kubernetes version 1.18.0 or later
- Nodes have an IPv6 Internal-IP

```bash
➜ kubectl get nodes -o wide
NAME                        STATUS   ROLES    AGE   VERSION   INTERNAL-IP              EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION      CONTAINER-RUNTIME
k8s-agentpool1-38564210-0   Ready    agent    23h   v1.18.6   2001:1234:5678:9abc::4   <none>        Ubuntu 16.04.7 LTS   4.15.0-1093-azure   docker://19.3.12
k8s-agentpool1-38564210-1   Ready    agent    23h   v1.18.6   2001:1234:5678:9abc::5   <none>        Ubuntu 16.04.7 LTS   4.15.0-1093-azure   docker://19.3.12
k8s-master-38564210-0       Ready    master   23h   v1.18.6   2001:1234:5678:9abc::6   <none>        Ubuntu 16.04.7 LTS   4.15.0-1093-azure   docker://19.3.12
```

- Nodes have 2 internal IPs, one from each ip family. IPv6 clusters on Azure are run on dual-stack hosts. The IPv6 is the primary IP.

```bash
➜ kubectl get nodes k8s-agentpool1-38564210-0 -o go-template --template='{{range .status.addresses}}{{printf "%s: %s \n" .type .address}}{{end}}'
Hostname: k8s-agentpool1-38564210-0
InternalIP: 2001:1234:5678:9abc::4
InternalIP: 10.240.0.4
```

- Nodes have IPv6 PodCIDR

```bash
➜ kubectl get nodes k8s-agentpool1-38564210-0 -o go-template --template='{{.spec.podCIDR}}'
fc00::/64
```

- Pods have IPv6 IP

```bash
kubectl get pods nginx-pod -o go-template --template='{{.status.podIP}}'
fc00:0:0:2::3
```

- Able to reach other pods in cluster using IPv6

```bash
# inside the nginx-pod
# ifconfig eth0
eth0      Link encap:Ethernet  HWaddr F2:20:2E:93:24:6E
          inet6 addr: fc00:0:0:2::3/64 Scope:Global
          inet6 addr: fe80::f020:2eff:fe93:246e/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:102 errors:0 dropped:0 overruns:0 frame:0
          TX packets:106 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:16731 (16.3 KiB)  TX bytes:10755 (10.5 KiB)
# ping fc00::6
PING fc00::6 (fc00::6): 56 data bytes
64 bytes from fc00::6: seq=0 ttl=62 time=0.731 ms
64 bytes from fc00::6: seq=1 ttl=62 time=1.060 ms
64 bytes from fc00::6: seq=2 ttl=62 time=1.165 ms
```

- Kubernetes services have IPv6 ClusterIP and ExternalIP

```bash
➜ kubectl get svc
NAME         TYPE           CLUSTER-IP   EXTERNAL-IP          PORT(S)        AGE
kubernetes   ClusterIP      fd00::1      <none>               443/TCP        23h
nginx-ipv6   LoadBalancer   fd00::e049   2603:1030:805:1::e   80:30955/TCP   7m51s
```

- Able to reach the workload on IPv6 ExternalIP

```bash
➜ curl http://\[2603:1030:805:1::e\] -v
*   Trying 2603:1030:805:1::e...
* TCP_NODELAY set
* Connected to 2603:1030:805:1::e (2603:1030:805:1::e) port 80 (#0)
> GET / HTTP/1.1
> Host: [2603:1030:805:1::e]
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: nginx/1.17.0
< Date: Wed, 02 Sep 2020 18:43:24 GMT
< Content-Type: text/html
< Content-Length: 612
< Last-Modified: Tue, 21 May 2019 15:33:12 GMT
< Connection: keep-alive
< ETag: "5ce41a38-264"
< Accept-Ranges: bytes
```

## Limitations

- IPv6 clusters are supported only with Kubernetes version 1.18+.
- IPv6 clusters are supported only with kubenet.
- IPv6 clusters are supported only with Linux.
- API model enables Azure Standard LB for IPv6 clusters. Azure Basic LBs have a limitation of only 1 IPv6 frontend configurations while Standard LB supports up to 600 IPv6 frontend configurations.
- Kubernetes service of type `LoadBalancer` needs to be created for egress to work after cluster is created.
- Currently, IPv6 clusters are supported only with Ubuntu 16.04 distro (`ubuntu`).
