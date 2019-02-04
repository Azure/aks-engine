# AKS Engine - Large Clusters

## Overview

AKS Engine enables you to create customized Kubernetes cluster on Microsoft Azure with 1200 nodes.

The examples show you how to configure up to 12 agent pools with 100 nodes each:

1. **dcos.json** - deploying and using [DC/OS](../../docs/dcos.md)
2. **dcos-vmas.json** - this provides an example using availability sets instead of the default virtual machine scale sets.  You will want to use availability sets if you want to dynamically attach/detach disks.
3. **kubernetes.json** - deploying and using [Kubernetes](../../docs/kubernetes.md)
4. **swarm.json** - deploying and using [Swarm](../../docs/swarm.md)
5. **swarm-vmas.json** - this provides an example using availability sets instead of the default virtual machine scale sets.  You will want to use availability sets if you want to dynamically attach/detach disks.
6. **swarmmode.json** - deploying and using [Swarm Mode](../../docs/swarmmode.md)
7. **swarmmode-vmas.json** - this provides an example using availability sets instead of the default virtual machine scale sets.  You will want to use availability sets if you want to dynamically attach/detach disks.
