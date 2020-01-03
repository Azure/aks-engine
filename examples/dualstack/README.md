# AKS Engine - Dual Stack clusters

## Overview

AKS Engine enables you to create dual stack (IPv4 *and* IPv6) Kubernetes clusters on Microsoft Azure.

- Dual stack support is available for Kubernetes version 1.16.0 and later

> Official docs are available here - https://kubernetes.io/docs/concepts/services-networking/dual-stack/

In order to create IPv6 enabled Azure virtual networks and use standard loadbalancer with IPv6 you must first configure your subscription [as follows](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-ipv4-ipv6-dual-stack-cli#prerequisites).

This example shows you how to configure a dual stack cluster:

1. **kubernetes.json** - deploying and using [Kubernetes](kubernetes.json).

Things to try out after the cluster is deployed -

- Nodes are Kubernetes version 1.17.0 or later

```bash
➜ kubectl get nodes
NAME                        STATUS   ROLES    AGE     VERSION
k8s-linuxpool1-91898541-0   Ready    agent    3m25s   v1.17.0
k8s-linuxpool1-91898541-1   Ready    agent    3m14s   v1.17.0
k8s-master-91898541-0       Ready    master   3m25s   v1.17.0
```

- Nodes have 2 internal IPs, one from each ip family

```bash
$ kubectl get nodes k8s-linuxpool1-20403072-0 -o go-template --template='{{range .status.addresses}}{{printf "%s: %s \n" .type .address}}{{end}}'
Hostname: k8s-pool1-12324934-0
InternalIP: 10.240.0.5
InternalIP: 2001:1234:5678:9abc::6
```

- Nodes have 2 PodCIDRs, one from each ip family

```bash
➜ kubectl get nodes k8s-linuxpool1-91898541-0 -o go-template --template='{{range .status.addresses}}{{printf "%s: %s \n" .type .address}}{{end}}'
Hostname: k8s-linuxpool1-91898541-0
InternalIP: 10.240.0.4
InternalIP: 2001:1234:5678:9abc::4
```

- Pods have 2 PodIPs, one from each ip family

```bash
➜ kubectl get pods nginx-58fdbbdb65-c5hv9 -o go-template --template='{{range .status.podIPs}}{{printf "%s \n" .ip}}{{end}}'
10.244.1.6
fc00::1:0:6
```

- Able to reach other pods in cluster using IPv6

```bash
# inside the nginx-pod
# ifconfig eth0
eth0      Link encap:Ethernet  HWaddr 3A:4A:38:FB:75:C1
          inet addr:10.244.0.7  Bcast:0.0.0.0  Mask:255.255.255.0
          inet6 addr: fe80::384a:38ff:fefb:75c1/64 Scope:Link
          inet6 addr: fc00::7/96 Scope:Global
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:76 errors:0 dropped:0 overruns:0 frame:0
          TX packets:72 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1321257 (1.2 MiB)  TX bytes:5639 (5.5 KiB)

# ping fc00::1:0:6
/ # ping fc00::1:0:6
PING fc00::1:0:6 (fc00::1:0:6): 56 data bytes
64 bytes from fc00::1:0:6: seq=0 ttl=62 time=2.067 ms
64 bytes from fc00::1:0:6: seq=1 ttl=62 time=0.817 ms
64 bytes from fc00::1:0:6: seq=2 ttl=62 time=1.373 ms
```

- Able to create services with IPv6 using `spec.IPFamily=IPv6` in the service manifest -

```
azureuser@k8s-master-13083844-0:~$ kubectl get svc
NAME          TYPE           CLUSTER-IP       EXTERNAL-IP          PORT(S)        AGE
kubernetes    ClusterIP      10.0.0.1         <none>               443/TCP        58m
nginx-ipv6    LoadBalancer   fd00::6283       2603:1030:805:3::3   80:31140/TCP   32s
```

The default node CIDR mask size for IPv6 is 64 with 1.17+. This can be configured as desired by setting 
```json
                "controllerManagerConfig": {
                    "--node-cidr-mask-size-ipv6": "96"
                }
```

Docs: https://kubernetes.io/docs/concepts/services-networking/dual-stack/#enable-ipv4-ipv6-dual-stack

**Note: The difference between the node CIDR mask size and cluster CIDR mask size can't be >16**

## Limitations

- Dual stack clusters are supported only with kubenet.
- Dual stack clusters are supported only with Linux.
- Dual stack clusters are currently only supported with ipvs kube-proxy mode.
- Dual stack clusters are currently only supported with Availability sets.
- API model enables Azure Standard LB for dual stack clusters. Azure Basic LBs have a limitation of only 1 IPv6 frontend configurations while Standard LB supports up to 600 IPv6 frontend configurations.