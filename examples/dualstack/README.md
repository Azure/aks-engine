# AKS Engine - Dual Stack clusters

## Overview

AKS Engine enables you to create dual stack (IPv4 *and* IPv6) Kubernetes clusters on Microsoft Azure.

- Dual stack support is available for Kubernetes version 1.16.0-alpha.1 and later

In order to create IPv6 enabled Azure virtual networks you must first configure your subscription [as follows](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-ipv4-ipv6-dual-stack-cli#prerequisites).

This example shows you how to configure a dual stack cluster:

1. **kubernetes.json** - deploying and using [Kubernetes](kubernetes.json).

Things to try out after the cluster is deployed -

- Nodes are Kubernetes version 1.16.0-alpha.1 or later

```bash
$ kubectl get nodes
NAME                        STATUS   ROLES    AGE    VERSION
k8s-linuxpool1-20403072-0   Ready    agent    116s   v1.16.0-alpha.1
k8s-linuxpool1-20403072-1   Ready    agent    50s    v1.16.0-alpha.1
k8s-master-20403072-0       Ready    master   2m7s   v1.16.0-alpha.1
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
$ kubectl get nodes k8s-linuxpool1-20403072-0 -o go-template --template='{{range .spec.podCIDRs}}{{printf "%s\n" .}}{{end}}'
10.244.2.0/24
fd00:200::/24
```

- Pods have 2 PodIPs, one from each ip family

```bash
kubectl get pods nginx-pod -o go-template --template='{{range .status.podIPs}}{{printf "%s \n" .ip}}{{end}}'
10.244.2.6
fd00:200::6
```

- Able to reach other pods in cluster using IPv6

```bash
# inside the nginx-pod
# ifconfig eth0
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 10.244.2.6  netmask 255.255.255.0  broadcast 0.0.0.0
        inet6 fd00:200::6  prefixlen 24  scopeid 0x0<global>
        inet6 fe80::8846:8cff:fe35:eaf0  prefixlen 64  scopeid 0x20<link>
        ether 8a:46:8c:35:ea:f0  txqueuelen 0  (Ethernet)
        RX packets 611  bytes 8685170 (8.2 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 415  bytes 35685 (34.8 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
# ping fd00:100::8
PING fd00:100::8(fd00:100::8) 56 data bytes
64 bytes from fd00:100::8: icmp_seq=1 ttl=62 time=0.798 ms
64 bytes from fd00:100::8: icmp_seq=2 ttl=62 time=0.762 ms
```

## Limitations

- Dual stack clusters are supported only with kubenet.
- Dual stack clusters are supported only with Linux.
- Egress pod internet routing will be available after this pending PR (https://github.com/kubernetes-incubator/ip-masq-agent/pull/45) will be merged.
