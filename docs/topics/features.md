# Features

|Feature|Status|API Version|Example|Description|
|---|---|---|---|---|
|Antrea Network Policy|Alpha|`vlabs`|[kubernetes-antrea.json](../../examples/networkpolicy/kubernetes-antrea.json)|[Description](#feat-antrea)|
|Azure Key Vault Encryption|Alpha|`vlabs`|[kubernetes-keyvault-encryption.json](../../examples/kubernetes-config/kubernetes-keyvault-encryption.json)|[Description](#feat-keyvault-encryption)|
|Calico Network Policy|Alpha|`vlabs`|[kubernetes-calico.json](../../examples/networkpolicy/kubernetes-calico-azure.json)|[Description](#feat-calico)|
|Cilium Network Policy|Alpha|`vlabs`|[kubernetes-cilium.json](../../examples/networkpolicy/kubernetes-cilium.json)|[Description](#feat-cilium)|
|ContainerD Runtime for Windows|Experimental|`vlabs`|[kubernetes-hybrid.containerd.json](../../examples/windows/kubernetes-hybrid.containerd.json)|[Description](#windows-containerd)|
|Custom VNET|Beta|`vlabs`|[kubernetesvnet-azure-cni.json](../../examples/vnet/kubernetesvnet-azure-cni.json)|[Description](#feat-custom-vnet)|
|Ephemeral OS Disks|Experimental|`vlabs`|[ephmeral-disk.json](../../examples/disks-ephemeral/ephemeral-disks.json)|[Description](#ephemeral-os-disks)|
|Managed Disks|Beta|`vlabs`|[kubernetes-vmas.json](../../examples/disks-managed/kubernetes-vmas.json)|[Description](#feat-managed-disks)|
|Private Cluster|Alpha|`vlabs`|[kubernetes-private-cluster.json](../../examples/kubernetes-config/kubernetes-private-cluster.json)|[Description](#feat-private-cluster)|
|Shared Image Gallery images|Alpha|`vlabs`|[custom-shared-image.json](../../examples/custom-shared-image.json)|[Description](#feat-shared-image-gallery)|

<a name="feat-kubernetes-msi"></a>

## Managed Identity

Enabling Managed Identity configures aks-engine to include and use MSI identities for all interactions with the Azure Resource Manager (ARM) API.

Instead of using a static service principal written to `/etc/kubernetes/azure.json`, Kubernetes will use a dynamic, time-limited token fetched from the MSI extension running on master and agent nodes. This support is currently alpha and requires Kubernetes v1.9.1 or newer.

Enable Managed Identity by adding `useManagedIdentity` in `kubernetesConfig`.

```json
"kubernetesConfig": {
  "useManagedIdentity": true
}
```

<a name="feat-managed-disks"></a>

## Optional: Disable Kubernetes Role-Based Access Control (RBAC) (for clusters running Kubernetes versions before 1.15.0)

By default, the cluster will be provisioned with [Role-Based Access Control](https://kubernetes.io/docs/admin/authorization/rbac/) enabled. Disable RBAC by adding `enableRbac` in `kubernetesConfig` in the API model:

```json
"kubernetesConfig": {
  "enableRbac": false
}
```

To emphasize: RBAC support is required for all Kubernetes clusters >= 1.15.0

See [cluster definitions](clusterdefinitions.md#kubernetesconfig) for further detail.

## Managed Disks

[Managed disks](../../examples/disks-managed/README.md) are supported for both node OS disks and Kubernetes persistent volumes.

Related [upstream PR](https://github.com/kubernetes/kubernetes/pull/46360) for details.

### Using Kubernetes Persistent Volumes

By default, each AKS Engine cluster is bootstrapped with several StorageClass resources. This bootstrapping is handled by the addon-manager pod that creates resources defined under /etc/kubernetes/addons directory on master VMs.

#### Non-managed Disks

The default storage class has been set via the Kubernetes admission controller `DefaultStorageClass`.

The default storage class will be used if persistent volume resources don't specify a storage class as part of the resource definition.

The default storage class uses non-managed blob storage and will provision the blob within an existing storage account present in the resource group or provision a new storage account.

Non-managed persistent volume types are available on all VM sizes.

#### Managed Disks

As part of cluster bootstrapping, two storage classes will be created to provide access to create Kubernetes persistent volumes using Azure managed disks.

Nodes will be labelled as follows if they support managed disks:

```text
storageprofile=managed
storagetier=<Standard_LRS|Premium_LRS>
```

They are managed-premium and managed-standard and map to Standard_LRS and Premium_LRS managed disk types respectively.

```console
kubectl get nodes -l storageprofile=managed
NAME                    STATUS    AGE       VERSION
k8s-agent1-23731866-0   Ready     24m       v1.12.8
```

- The VM size must support the type of managed disk type requested. For example, Premium VM sizes with managed OS disks support both managed-standard and managed-premium storage classes whereas Standard VM sizes with managed OS disks only support managed-standard storage class.

- If you have mixed node cluster (both non-managed and managed disk types). You must use [affinity or nodeSelectors](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/) on your resource definitions in order to ensure that workloads are scheduled to VMs that support the underlying disk requirements.

For example:

```yaml
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: storageprofile
            operator: In
            values:
            - managed
```

## Using Azure integrated networking (CNI)

Kubernetes clusters are configured by default to use the [Azure CNI plugin](https://github.com/Azure/azure-container-networking) which provides an Azure native networking experience. Pods will receive IP addresses directly from the vnet subnet on which they're hosted. If the API model doesn't specify explicitly, aks-engine will automatically provide the following `networkPlugin` configuration in `kubernetesConfig`:

```json
"kubernetesConfig": {
  "networkPlugin": "azure"
}
```

### Additional Azure integrated networking configuration

In addition you can modify the following settings to change the networking behavior when using Azure integrated networking:

IP addresses are pre-allocated in the subnet. Using ipAddressCount you can specify how many you would like to pre-allocate. This number needs to account for number of pods you would like to run on that subnet.

```json
"masterProfile": {
  "ipAddressCount": 200
},
```

Currently, the IP addresses that are pre-allocated aren't allowed by the default natter for Internet bound traffic. In order to work around this limitation we allow the user to specify the vnetCidr (eg. 10.0.0.0/8) to be EXCLUDED from the default masquerade rule that is applied. The result is that traffic destined for anything within that block will NOT be natted on the outbound VM interface. This field has been called vnetCidr but may be wider than the vnet cidr block if you would like POD IPs to be routable across vnets using vnet-peering or express-route.

```json
"masterProfile": {
  "vnetCidr": "10.0.0.0/8",
},
```

When using Azure integrated networking the maxPods setting will be set to 30 by default. This number can be changed keeping in mind that there is a limit of 65,536 IPs per vnet.

```json
"kubernetesConfig": {
  "kubeletConfig": {
    "--max-pods": "50"
  }
}
```

<a name="feat-calico"></a>

## Network Policy Enforcement with Calico

Using the default configuration, Kubernetes allows communication between all
Pods within a cluster. To ensure that Pods can only be accessed by authorized
Pods, a policy enforcement is needed. To enable policy enforcement using Calico refer to the [cluster definitions](clusterdefinitions.md#kubernetesconfig) document under networkPolicy. There is also a reference cluster definition available [here](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy/kubernetes-calico-azure.json).

This will deploy a Calico node controller to every instance of the cluster
using a Kubernetes DaemonSet. After a successful deployment you should be able
to see these Pods running in your cluster:

```console
$ kubectl get pods --namespace kube-system -l k8s-app=calico-node -o wide
NAME                READY     STATUS    RESTARTS   AGE       IP             NODE
calico-node-034zh   2/2       Running   0          2h        10.240.255.5   k8s-master-30179930-0
calico-node-qmr7n   2/2       Running   0          2h        10.240.0.4     k8s-agentpool1-30179930-1
calico-node-z3p02   2/2       Running   0          2h        10.240.0.5     k8s-agentpool1-30179930-0
```

Per default Calico still allows all communication within the cluster. Using Kubernetes' NetworkPolicy API, you can define stricter policies. Good resources to get information about that are:

- [NetworkPolicy User Guide](https://kubernetes.io/docs/user-guide/networkpolicies/)
- [NetworkPolicy Example Walkthrough](https://kubernetes.io/docs/getting-started-guides/network-policy/walkthrough/)
- [Calico Kubernetes](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy)

### Calico 3.3 cleanup after upgrading to 3.5 or greater

Because Calico 3.3 is using Calico CNI, while Calico 3.5 or greater moves to Azure CNI, if the cluster is upgraded from calico 3.3 to 3.5 or greater, then some manual cluster resource cleanup will be required to successfully complete the upgrade. We've provided a sample resource spec here that can be used as an example:

https://github.com/Azure/aks-engine/raw/master/docs/topics/calico-3.3.1-cleanup-after-upgrade.yaml

There are some placeholder tokens in the above `yaml` file, so please reconcile those with the actual values in your cluster. Look for these placeholder strings in the spec, then compare with the running spec of the comparable pre-3.5 resource in your cluster, and modify the cleanup spec accordingly:

- `<calicoIPAMConfig>`
- `<kubeClusterCidr>`

And then using your modified file, do something like this:

```sh
kubectl delete -f calico-3.3.1-cleanup-after-upgrade-modified-with-my-cluster-configuration.yaml
```

After this, addon-manager would enforce the correct spec for Calico 3.5 or greater.

<a name="feat-cilium"></a>

## Network Policy Enforcement with Cilium

Using the default configuration, Kubernetes allows communication between all
Pods within a cluster. To ensure that Pods can only be accessed by authorized
Pods, a policy enforcement is needed. To enable policy enforcement using Cilium refer to the
[cluster definitions](clusterdefinitions.md#kubernetesconfig)
document under networkPolicy. There is also a reference cluster definition available
[here](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy/kubernetes-cilium.json).

This will deploy a Cilium agent to every instance of the cluster
using a Kubernetes DaemonSet. After a successful deployment you should be able
to see these Pods running in your cluster:

```console
$ kubectl get pods --namespace kube-system -l k8s-app=cilium -o wide
NAME                READY     STATUS    RESTARTS   AGE       IP             NODE
cilium-034zh   2/2       Running   0          2h        10.240.255.5   k8s-master-30179930-0
cilium-qmr7n   2/2       Running   0          2h        10.240.0.4     k8s-agentpool1-30179930-1
cilium-z3p02   2/2       Running   0          2h        10.240.0.5     k8s-agentpool1-30179930-0
```

Per default Cilium still allows all communication within the cluster. Using Kubernetes' NetworkPolicy API,
you can define stricter policies. Good resources to get information about that are:

- [Cilum Network Policy Docs](https://cilium.readthedocs.io/en/latest/kubernetes/policy/#k8s-policy)
- [NetworkPolicy User Guide](https://kubernetes.io/docs/user-guide/networkpolicies/)
- [NetworkPolicy Example Walkthrough](https://kubernetes.io/docs/getting-started-guides/network-policy/walkthrough/)
- [Cilium Kubernetes](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy)

<a name="feat-antrea"></a>

## Network Policy Enforcement with Antrea

Using the default configuration, Kubernetes allows communication between all
Pods within a cluster. To ensure that Pods can only be accessed by authorized
Pods, a policy enforcement is needed. To enable policy enforcement using Antrea refer to the
[cluster definitions](clusterdefinitions.md#kubernetesconfig)
document under networkPolicy. There is also a reference cluster definition available
[here](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy/kubernetes-antrea.json).

This will deploy single replica of Antrea controller and Antrea agent to every
instance of the cluster using a Kubernetes DaemonSet. After a successful deployment
you should be able to see these Pods running in your cluster:

```console
kubectl get pods --namespace kube-system  -l app=antrea -o wide
NAME                                 READY   STATUS    RESTARTS   AGE     IP             NODE
antrea-agent-67t9z                   2/2     Running   1          7m38s   10.240.0.5     k8s-agentpool1-14956401-vmss000001
antrea-agent-87nm2                   2/2     Running   0          11m     10.240.0.4     k8s-agentpool1-14956401-vmss000000
antrea-agent-fhbsg                   2/2     Running   0          11m     10.240.0.6     k8s-agentpool1-14956401-vmss000002
antrea-agent-jjhxt                   2/2     Running   0          11m     10.240.255.5   k8s-master-14956401-0
antrea-controller-685c8c6f64-zk4jh   1/1     Running   0          11m     10.240.0.4     k8s-agentpool1-14956401-vmss000000
```

Per default Antrea still allows all communication within the cluster. Using Kubernetes' NetworkPolicy API,
you can define stricter policies. Good resources to get information about that are:

- [Antrea Architecture Docs](https://github.com/vmware-tanzu/antrea/blob/master/docs/architecture.md)
- [NetworkPolicy User Guide](https://kubernetes.io/docs/user-guide/networkpolicies/)
- [NetworkPolicy Example Walkthrough](https://kubernetes.io/docs/getting-started-guides/network-policy/walkthrough/)
- [Antrea Kubernetes](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy)

<a name="feat-custom-vnet"></a>

## Custom VNET

*Note: Custom VNET for Kubernetes Windows cluster has a [known issue](https://github.com/Azure/aks-engine/issues/371).*

AKS Engine supports deploying into an existing VNET. Operators must specify the ARM path/id of Subnets for the `masterProfile` and  any `agentPoolProfiles`, as well as the first IP address to use for static IP allocation in `firstConsecutiveStaticIP`. Please note that in any azure subnet, the first four and the last ip address is reserved and can not be used. Additionally, each pod now gets the IP address from the Subnet. As a result, enough IP addresses (equal to `ipAddressCount` for each node) should be available beyond `firstConsecutiveStaticIP`. By default, the `ipAddressCount` has a value of 31, 1 for the node and 30 for pods, (note that the number of pods can be changed via `KubeletConfig["--max-pods"]`). `ipAddressCount` can be changed if desired. Furthermore, to prevent source address NAT'ing within the VNET, we assign to the `vnetCidr` property in `masterProfile` the CIDR block that represents the usable address space in the existing VNET. Therefore, it is recommended to use a large subnet size such as `/16`.

Depending upon the size of the VNET address space, during deployment, it is possible to experience IP address assignment collision between the required Kubernetes static IPs (one each per master and one for the API server load balancer, if more than one masters) and Azure CNI-assigned dynamic IPs (one for each NIC on the agent nodes). In practice, the larger the VNET the less likely this is to happen; some detail, and then a guideline.

First, the detail:

- Azure CNI assigns dynamic IP addresses from the "beginning" of the subnet IP address space (specifically, it looks for available addresses starting at ".4" ["10.0.0.4" in a "10.0.0.0/24" network])
- AKS Engine will require a range of up to 16 unused IP addresses in multi-master scenarios (1 per master for up to 5 masters, and then the next 10 IP addresses immediately following the "last" master for headroom reservation, and finally 1 more for the load balancer immediately adjacent to the afore-described _n_ masters+10 sequence) to successfully scaffold the network stack for your cluster

A guideline that will remove the danger of IP address allocation collision during deployment:

- If possible, assign to the `firstConsecutiveStaticIP` configuration property an IP address that is near the "end" of the available IP address space in the desired  subnet.
  - For example, if the desired subnet is a `/24`, choose the "239" address in that network space

In larger subnets (e.g., `/16`) it's not as practically useful to push static IP assignment to the very "end" of large subnet, but as long as it's not in the "first" `/24` (for example) your deployment will be resilient to this edge case behavior.

Before provisioning, modify the `masterProfile` and `agentPoolProfiles` to match the above requirements, with the below being a representative example:

```js
"masterProfile": {
  ...
  "vnetSubnetId": "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/MASTER_SUBNET_NAME",
  "firstConsecutiveStaticIP": "10.239.255.239",
  "vnetCidr": "10.239.0.0/16",
  ...
},
...
"agentPoolProfiles": [
  {
    ...
    "name": "agentpri",
    "vnetSubnetId": "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/AGENT_SUBNET_NAME",
    ...
  },
```

### VirtualMachineScaleSets Masters Custom VNET

When using custom VNET with `VirtualMachineScaleSets` MasterProfile, make sure to create two subnets within the vnet: `master` and `agent`.
Modify `masterProfile` in the API model, `vnetSubnetId`, `agentVnetSubnetId` should be set to the values of the `master` subnet and the `agent` subnet in the existing vnet respectively.
Modify `agentPoolProfiles`, `vnetSubnetId` should be set to the value of the `agent` subnet in the existing vnet.

*NOTE: The `firstConsecutiveStaticIP` configuration should be empty and will be derived from an offset and the first IP in the vnetCidr.*
For example, if `vnetCidr` is `10.239.0.0/16`, `master` subnet is `10.239.0.0/17`, `agent` subnet is `10.239.128.0/17`, then `firstConsecutiveStaticIP` will be `10.239.0.4`.

```js
"masterProfile": {
  ...
  "vnetSubnetId": "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/MASTER_SUBNET_NAME",
  "agentVnetSubnetId": "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/AGENT_SUBNET_NAME",
  "vnetCidr": "10.239.0.0/16",
  ...
},
...
"agentPoolProfiles": [
  {
    ...
    "name": "agentpri",
    "vnetSubnetId": "/subscriptions/SUB_ID/resourceGroups/RG_NAME/providers/Microsoft.Network/virtualNetworks/VNET_NAME/subnets/AGENT_SUBNET_NAME",
    ...
  },
```

### Kubenet Networking Custom VNET

If you're *not- using Azure CNI (e.g., `"networkPlugin": "kubenet"` in the `kubernetesConfig` API model configuration object): After a custom VNET-configured cluster finishes provisioning, fetch the id of the Route Table resource from `Microsoft.Network` provider in your new cluster's Resource Group.

The route table resource id is of the format: `/subscriptions/SUBSCRIPTIONID/resourceGroups/RESOURCEGROUPNAME/providers/Microsoft.Network/routeTables/ROUTETABLENAME`

Existing subnets will need to use the Kubernetes-based Route Table so that machines can route to Kubernetes-based workloads.

Update properties of all subnets in the existing VNET route table resource by appending the following to subnet properties:

```json
"routeTable": {
        "id": "/subscriptions/<SubscriptionId>/resourceGroups/<ResourceGroupName>/providers/Microsoft.Network/routeTables/k8s-master-<SOMEID>-routetable>"
}
```

E.g.:

```js
"subnets": [
    {
      "name": "subnetname",
      "id": "/subscriptions/<SubscriptionId>/resourceGroups/<ResourceGroupName>/providers/Microsoft.Network/virtualNetworks/<VirtualNetworkName>/subnets/<SubnetName>",
      "properties": {
        "provisioningState": "Succeeded",
        "addressPrefix": "10.240.0.0/16",
        "routeTable": {
          "id": "/subscriptions/<SubscriptionId>/resourceGroups/<ResourceGroupName>/providers/Microsoft.Network/routeTables/k8s-master-<SOMEID>-routetable"
        }
      ...
      }
      ...
    }
]
```

<a name="feat-private-cluster"></a>

## Private Cluster

You can build a private Kubernetes cluster with no public IP addresses assigned by setting:

```json
"kubernetesConfig": {
  "privateCluster": {
    "enabled": true
}
```

In order to access this cluster using kubectl commands, you will need a jumpbox in the same VNET (or onto a peer VNET that routes to the VNET). If you do not already have a jumpbox, you can use aks-engine to provision your jumpbox (see below) or create it manually. You can create a new jumpbox manually in the Azure Portal under "Create a resource > Compute > Ubuntu Server 16.04 LTS VM" or using the [az cli](https://docs.microsoft.com/en-us/cli/azure/vm?view=azure-cli-latest#az_vm_create). You will then be able to:

- install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) on the jumpbox
- copy the kubeconfig artifact for the right region from the deployment directory to the jumpbox
- run `export KUBECONFIG=<path to your kubeconfig>`
- run `kubectl` commands directly on the jumpbox

Alternatively, you may also ssh into your nodes (given that your ssh key is on the jumpbox) and use the admin user kubeconfig on the cluster to run `kubectl` commands directly on the cluster. However, in the case of a multi-master private cluster, the connection will be refused when running commands on a master every time that master gets picked by the load balancer as it will be routing to itself (1 in 3 times for a 3 master cluster, 1 in 5 for 5 masters). This is expected behavior and therefore the method aforementioned of accessing nodes on the jumpbox using the `_output` directory kubeconfig is preferred.

To auto-provision a jumpbox with your aks-engine deployment use:

```json
"kubernetesConfig": {
  "privateCluster": {
    "enabled": true,
    "jumpboxProfile": {
      "name": "my-jb",
      "vmSize": "Standard_D4s_v3",
      "osDiskSizeGB": 30,
      "username": "azureuser",
      "publicKey": "xxx"
    }
}
```

<a name="feat-keyvault-encryption"></a>

## Azure Key Vault Data Encryption

Enabling Azure Key Vault Encryption configures aks-engine to create an Azure Key Vault in the same resource group as the Kubernetes cluster and configures Kubernetes to use a key from this Key Vault to encrypt and decrypt etcd data for the Kubernetes cluster.

To enable this feature, add `"enableEncryptionWithExternalKms": true` in `kubernetesConfig` and `objectId` in `servicePrincipalProfile`. Optional, if you want to create Hardware Security Modules (HSM) type keys, then add `"keyVaultSku": "Premium"` to enable creation of Premium SKU Key Vault and RSA-HSM type key. Otherwise, by default `keyVaultSku` can be omitted and a Standard SKU Key Vault and a RSA type key will be created.

```json
"kubernetesConfig": {
  "enableEncryptionWithExternalKms": true,
  "keyVaultSku": "Premium",
}
...

"servicePrincipalProfile": {
  "clientId": "",
  "secret": "",
  "objectId": ""
}
```

> Note: `objectId` is the objectId of the service principal used to create the key vault and to be granted access to keys in this key vault.

To get `objectId` of the service principal:

```console
az ad sp list --spn <YOUR SERVICE PRINCIPAL appId>
```

<a name="feat-shared-image-gallery"></a>

## Use a Shared Image Gallery image

This is possible by specifying `imageReference` under `masterProfile` or on a given `agentPoolProfile`. It also requires setting the distro to an appropriate value (e.g., `ubuntu`). When using `imageReference` with Shared Image Galleries, provide an image name and version, as well as the resource group, subscription, and name of the gallery. Example:

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes"
    },
    "masterProfile": {
      "imageReference": {
        "name": "linuxvm",
        "resourceGroup": "sig",
        "subscriptionID": "00000000-0000-0000-0000-000000000000",
        "gallery": "siggallery",
        "version": "0.0.1"
      },
      "count": 1,
      "dnsPrefix": "",
      "vmSize": "Standard_D2_v3"
    },
    "agentPoolProfiles": [
      {
        "name": "agentpool1",
        "count": 3,
        "imageReference": {
          "name": "linuxvm",
          "resourceGroup": "sig",
          "subscriptionID": "00000000-0000-0000-0000-000000000000",
          "gallery": "siggallery",
          "version": "0.0.1"
        },
        "vmSize": "Standard_D2_v3",
        "availabilityProfile": "AvailabilitySet"
      }
    ],
    "linuxProfile": {
      "adminUsername": "azureuser",
      "ssh": {
        "publicKeys": [
          {
            "keyData": ""
          }
        ]
      }
    },
    "servicePrincipalProfile": {
      "clientId": "",
      "secret": ""
    }
  }
}
```

## Ephemeral OS Disks

> This feature is considered experimental, and you may lose data. We're still evaluating what risks exist and how to mitigate them.

[Ephemeral OS Disks] is a new feature in Azure that allows the OS disk to use local SSD storage, with no writes to Azure storage. If a VM is stopped or deprovisioned, it's local storage is lost. If the same VM is restarted, it starts from the original OS disk and reapplies the custom script extension from AKS-Engine to join the cluster.

Benefits - VMs deploy faster, and have better local storage performance. The OS disk will perform at the _Max cached storage throughput_ for the VM size. For example with a `Standard_D2s_v3` size VM using a 50 GiB OS disk - it can achieve 4000 IOPs with ephemeral disks enabled, or 240 IOPs using a `Premium P6` [Premium SSD](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/disks-types#premium-ssd) at 50GiB. Apps will get faster container and `emptydir` performance. Container pull times are also improved.

Requirements:

- Be sure you are using a VM size that supports cache for the local disk
- The OS disk size must be set to <= the VM's _cache size in GiB_

These are fully explained in the [Ephemeral OS Disks] docs.


We are investigating possible risks & mitigations for when VMs are deprovisioned or moved for Azure maintenance:

- Logs for containers on those nodes are lost.
- Containers cannot be restarted on the same node, as their container directory and any emptydir volumes will be missing.


[Ephemeral OS Disks]: https://docs.microsoft.com/en-us/azure/virtual-machines/windows/ephemeral-os-disks


## Windows ContainerD

> This feature is currently experimental, and has open issues.

Kubernetes 1.18 introduces alpha support for the ContainerD runtime on Windows Server 2019. This is still a work-in-progress tracked in [kubernetes/enhancements#1001](https://github.com/kubernetes/enhancements/issues/1001). This feature in AKS-Engine is for testing the in-development versions of ContainerD and Kubernetes, and is not for production use. Be sure to review [open issues](https://github.com/azure/aks-engine/issues?q=containerd+label%3Awindows+is%3Aopen) if you want to test or contribute to this effort.

Currently it only supports the `kubenet` networking model, and requires URLs to custom ContainerD and CNI plugin builds.

### Deploying multi-OS clusters with ContainerD

If you want to test or develop with Windows & ContainerD in AKS-Engine, see this sample
[kubernetes-hybrid.containerd.json](../../examples/windows/kubernetes-hybrid.containerd.json)

These parameters are all required.

```json
      "kubernetesConfig": {
        "networkPlugin": "kubenet",
        "containerRuntime": "containerd",
        "windowsContainerdURL": "...",
        "windowsSdnPluginURL": "..."
      }
```

### Building ContainerD

As of March 3, 2020, the ContainerD and network plugin repos don't have public builds available. This repo has a script that will build them from source and create two ZIP files: [build-windows-containerd.sh](../../scripts/build-windows-containerd.sh)

Upload these ZIP files to a location that your cluster will be able to reach, then put those URLs in `windowsContainerdURL` and `windowsSdnPluginURL` in the AKS-Engine API model shown above.
