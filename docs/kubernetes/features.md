# Features

|Feature|Status|API Version|Example|Description|
|---|---|---|---|---|
|Managed Disks|Beta|`vlabs`|[kubernetes-vmas.json](../../examples/disks-managed/kubernetes-vmas.json)|[Description](#feat-managed-disks)|
|Calico Network Policy|Alpha|`vlabs`|[kubernetes-calico.json](../../examples/networkpolicy/kubernetes-calico.json)|[Description](#feat-calico)|
|Cilium Network Policy|Alpha|`vlabs`|[kubernetes-cilium.json](../../examples/networkpolicy/kubernetes-cilium.json)|[Description](#feat-cilium)|
|Custom VNET|Beta|`vlabs`|[kubernetesvnet-azure-cni.json](../../examples/vnet/kubernetesvnet-azure-cni.json)|[Description](#feat-custom-vnet)|
|Clear Containers Runtime|Alpha|`vlabs`|[kubernetes-clear-containers.json](../../examples/kubernetes-clear-containers.json)|[Description](#feat-clear-containers)|
|Kata Containers Runtime|Alpha|`vlabs`|[kubernetes-kata-containers.json](../../examples/kubernetes-kata-containers.json)|[Description](#feat-kata-containers)|
|Private Cluster|Alpha|`vlabs`|[kubernetes-private-cluster.json](../../examples/kubernetes-config/kubernetes-private-cluster.json)|[Description](#feat-private-cluster)|
|Azure Key Vault Encryption|Alpha|`vlabs`|[kubernetes-keyvault-encryption.json](../../examples/kubernetes-config/kubernetes-keyvault-encryption.json)|[Description](#feat-keyvault-encryption)|

<a name="feat-kubernetes-msi"></a>

## Managed Identity

Enabling Managed Identity configures aks-engine to include and use MSI identities for all interactions with the Azure Resource Manager (ARM) API.

Instead of using a static servic principal written to `/etc/kubernetes/azure.json`, Kubernetes will use a dynamic, time-limited token fetched from the MSI extension running on master and agent nodes. This support is currently alpha and requires Kubernetes v1.9.1 or newer.

Enable Managed Identity by adding `useManagedIdentity` in `kubernetesConfig`.

```
"kubernetesConfig": {
  "useManagedIdentity": true
}
```

<a name="feat-managed-disks"></a>

## Optional: Disable Kubernetes Role-Based Access Control (RBAC)

By default, the cluster will be provisioned with [Role-Based Access Control](https://kubernetes.io/docs/admin/authorization/rbac/) enabled. Disable RBAC by adding `enableRbac` in `kubernetesConfig` in the api model:

```console
      "kubernetesConfig": {
        "enableRbac": false
      }
```

See [cluster definition](https://github.com/Azure/aks-engine/blob/master/docs/clusterdefinition.md#kubernetesconfig) for further detail.

## Managed Disks

[Managed disks](../examples/disks-managed/README.md) are supported for both node OS disks and Kubernetes persistent volumes.

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

```
storageprofile=managed
storagetier=<Standard_LRS|Premium_LRS>
```

They are managed-premium and managed-standard and map to Standard_LRS and Premium_LRS managed disk types respectively.

In order to use these storage classes the following conditions must be met.

* The cluster must be running Kubernetes release 1.7 or greater. Refer to this [example](../examples/kubernetes-releases/kubernetes1.7.json) for how to provision a Kubernetes cluster of a specific version.
* The node must support managed disks. See this [example](../examples/disks-managed/kubernetes-vmas.json) to provision nodes with managed disks. You can also confirm if a node has managed disks using kubectl.

```console
kubectl get nodes -l storageprofile=managed
NAME                    STATUS    AGE       VERSION
k8s-agent1-23731866-0   Ready     24m       v1.7.2
```

* The VM size must support the type of managed disk type requested. For example, Premium VM sizes with managed OS disks support both managed-standard and managed-premium storage classes whereas Standard VM sizes with managed OS disks only support managed-standard storage class.

* If you have mixed node cluster (both non-managed and managed disk types). You must use [affinity or nodeSelectors](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/) on your resource definitions in order to ensure that workloads are scheduled to VMs that support the underlying disk requirements.

For example
```
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

Kubernetes clusters are configured by default to use the [Azure CNI plugin](https://github.com/Azure/azure-container-networking) which provides an Azure native networking experience. Pods will receive IP addresses directly from the vnet subnet on which they're hosted. If the api model doesn't specify explicitly, aks-engine will automatically provide the following `networkPlugin` configuration in `kubernetesConfig`:

```
      "kubernetesConfig": {
        "networkPlugin": "azure"
      }
```

### Additional Azure integrated networking configuration

In addition you can modify the following settings to change the networking behavior when using Azure integrated networking:

IP addresses are pre-allocated in the subnet. Using ipAddressCount you can specify how many you would like to pre-allocate. This number needs to account for number of pods you would like to run on that subnet.

```
    "masterProfile": {
      "ipAddressCount": 200
    },
```

Currently, the IP addresses that are pre-allocated aren't allowed by the default natter for Internet bound traffic. In order to work around this limitation we allow the user to specify the vnetCidr (eg. 10.0.0.0/8) to be EXCLUDED from the default masquerade rule that is applied. The result is that traffic destined for anything within that block will NOT be natted on the outbound VM interface. This field has been called vnetCidr but may be wider than the vnet cidr block if you would like POD IPs to be routable across vnets using vnet-peering or express-route.
```
    "masterProfile": {
      "vnetCidr": "10.0.0.0/8",
    },
```

When using Azure integrated networking the maxPods setting will be set to 30 by default. This number can be changed keeping in mind that there is a limit of 4,000 IPs per vnet.

```
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
Pods, a policy enforcement is needed. To enable policy enforcement using Calico refer to the [cluster definition](https://github.com/Azure/aks-engine/blob/master/docs/clusterdefinition.md#kubernetesconfig) document under networkPolicy. There is also a reference cluster definition available [here](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy/kubernetes-calico.json).

This will deploy a Calico node controller to every instance of the cluster
using a Kubernetes DaemonSet. After a successful deployment you should be able
to see these Pods running in your cluster:

```
kubectl get pods --namespace kube-system -l k8s-app=calico-node -o wide
NAME                READY     STATUS    RESTARTS   AGE       IP             NODE
calico-node-034zh   2/2       Running   0          2h        10.240.255.5   k8s-master-30179930-0
calico-node-qmr7n   2/2       Running   0          2h        10.240.0.4     k8s-agentpool1-30179930-1
calico-node-z3p02   2/2       Running   0          2h        10.240.0.5     k8s-agentpool1-30179930-0
```

Per default Calico still allows all communication within the cluster. Using Kubernetes' NetworkPolicy API, you can define stricter policies. Good resources to get information about that are:

* [NetworkPolicy User Guide](https://kubernetes.io/docs/user-guide/networkpolicies/)
* [NetworkPolicy Example Walkthrough](https://kubernetes.io/docs/getting-started-guides/network-policy/walkthrough/)
* [Calico Kubernetes](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy)

<a name="feat-cilium"></a>

## Network Policy Enforcement with Cilium

Using the default configuration, Kubernetes allows communication between all
Pods within a cluster. To ensure that Pods can only be accessed by authorized
Pods, a policy enforcement is needed. To enable policy enforcement using Cilium refer to the
[cluster definition](https://github.com/Azure/aks-engine/blob/master/docs/clusterdefinition.md#kubernetesconfig)
document under networkPolicy. There is also a reference cluster definition available
[here](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy/kubernetes-cilium.json).

This will deploy a Cilium agent to every instance of the cluster
using a Kubernetes DaemonSet. After a successful deployment you should be able
to see these Pods running in your cluster:

```
kubectl get pods --namespace kube-system -l k8s-app=cilium -o wide
NAME                READY     STATUS    RESTARTS   AGE       IP             NODE
cilium-034zh   2/2       Running   0          2h        10.240.255.5   k8s-master-30179930-0
cilium-qmr7n   2/2       Running   0          2h        10.240.0.4     k8s-agentpool1-30179930-1
cilium-z3p02   2/2       Running   0          2h        10.240.0.5     k8s-agentpool1-30179930-0
```

Per default Cilium still allows all communication within the cluster. Using Kubernetes' NetworkPolicy API,
you can define stricter policies. Good resources to get information about that are:

* [Cilum Network Policy Docs](https://cilium.readthedocs.io/en/latest/kubernetes/policy/#k8s-policy)
* [NetworkPolicy User Guide](https://kubernetes.io/docs/user-guide/networkpolicies/)
* [NetworkPolicy Example Walkthrough](https://kubernetes.io/docs/getting-started-guides/network-policy/walkthrough/)
* [Cilium Kubernetes](https://github.com/Azure/aks-engine/blob/master/examples/networkpolicy)

<a name="feat-custom-vnet"></a>

## Custom VNET

*Note: Custom VNET for Kubernetes Windows cluster has a [known issue](https://github.com/Azure/aks-engine/issues/1767).*

AKS Engine supports deploying into an existing VNET. Operators must specify the ARM path/id of Subnets for the `masterProfile` and  any `agentPoolProfiles`, as well as the first IP address to use for static IP allocation in `firstConsecutiveStaticIP`. Please note that in any azure subnet, the first four and the last ip address is reserved and can not be used. Additionally, each pod now gets the IP address from the Subnet. As a result, enough IP addresses (equal to `ipAddressCount` for each node) should be available beyond `firstConsecutiveStaticIP`. By default, the `ipAddressCount` has a value of 31, 1 for the node and 30 for pods, (note that the number of pods can be changed via `KubeletConfig["--max-pods"]`). `ipAddressCount` can be changed if desired. Furthermore, to prevent source address NAT'ing within the VNET, we assign to the `vnetCidr` property in `masterProfile` the CIDR block that represents the usable address space in the existing VNET. Therefore, it is recommended to use a large subnet size such as `/16`.

Depending upon the size of the VNET address space, during deployment, it is possible to experience IP address assignment collision between the required Kubernetes static IPs (one each per master and one for the API server load balancer, if more than one masters) and Azure CNI-assigned dynamic IPs (one for each NIC on the agent nodes). In practice, the larger the VNET the less likely this is to happen; some detail, and then a guideline.

First, the detail:

* Azure CNI assigns dynamic IP addresses from the "beginning" of the subnet IP address space (specifically, it looks for available addresses starting at ".4" ["10.0.0.4" in a "10.0.0.0/24" network])
* aks-engine will require a range of up to 16 unused IP addresses in multi-master scenarios (1 per master for up to 5 masters, and then the next 10 IP addresses immediately following the "last" master for headroom reservation, and finally 1 more for the load balancer immediately adjacent to the afore-described _n_ masters+10 sequence) to successfully scaffold the network stack for your cluster

A guideline that will remove the danger of IP address allocation collision during deployment:

* If possible, assign to the `firstConsecutiveStaticIP` configuration property an IP address that is near the "end" of the available IP address space in the desired  subnet.
  * For example, if the desired subnet is a `/24`, choose the "239" address in that network space

In larger subnets (e.g., `/16`) it's not as practically useful to push static IP assignment to the very "end" of large subnet, but as long as it's not in the "first" `/24` (for example) your deployment will be resilient to this edge case behavior.

Before provisioning, modify the `masterProfile` and `agentPoolProfiles` to match the above requirements, with the below being a representative example:

```
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
Modify `masterProfile` in the api model, `vnetSubnetId`, `agentVnetSubnetId` should be set to the values of the `master` subnet and the `agent` subnet in the existing vnet respectively.
Modify `agentPoolProfiles`, `vnetSubnetId` should be set to the value of the `agent` subnet in the existing vnet.

*NOTE: The `firstConsecutiveStaticIP` configuration should be empty and will be derived from an offset and the first IP in the vnetCidr.*
For example, if `vnetCidr` is `10.239.0.0/16`, `master` subnet is `10.239.0.0/17`, `agent` subnet is `10.239.128.0/17`, then `firstConsecutiveStaticIP` will be `10.239.0.4`.

```
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

If you're *not* using Azure CNI (e.g., `"networkPlugin": "kubenet"` in the `kubernetesConfig` api model configuration object): After a custom VNET-configured cluster finishes provisioning, fetch the id of the Route Table resource from `Microsoft.Network` provider in your new cluster's Resource Group.

The route table resource id is of the format: `/subscriptions/SUBSCRIPTIONID/resourceGroups/RESOURCEGROUPNAME/providers/Microsoft.Network/routeTables/ROUTETABLENAME`

Existing subnets will need to use the Kubernetes-based Route Table so that machines can route to Kubernetes-based workloads.

Update properties of all subnets in the existing VNET he route table resource by appending the following to subnet properties:

```
"routeTable": {
        "id": "/subscriptions/<SubscriptionId>/resourceGroups/<ResourceGroupName>/providers/Microsoft.Network/routeTables/k8s-master-<SOMEID>-routetable>"
      }
```

E.g.:
```
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

<a name="feat-clear-containers"></a>

## Clear Containers

You can designate kubernetes agents to use Intel's Clear Containers as the
container runtime by setting:

```
      "kubernetesConfig": {
        "containerRuntime": "clear-containers"
      }
```

You will need to make sure your agents are using a `vmSize` that [supports
nested virtualization](https://azure.microsoft.com/en-us/blog/nested-virtualization-in-azure/).
These are the `Dv3` or `Ev3` series nodes.

This should look like:

```
"agentPoolProfiles": [
      {
        "name": "agentpool1",
        "count": 3,
        "vmSize": "Standard_D4s_v3",
        "availabilityProfile": "AvailabilitySet",
        "diskSizesGB": [1023]
      }
    ],
```

<a name="feat-kata-containers"></a>

## Kata Containers

You can designate kubernetes agents to use Kata Containers as the
container runtime by setting:

```
      "kubernetesConfig": {
        "containerRuntime": "kata-containers"
      }
```

You will need to make sure your agents are using a `vmSize` that [supports
nested virtualization](https://azure.microsoft.com/en-us/blog/nested-virtualization-in-azure/).
These are the `Dv3` or `Ev3` series nodes.

This should look like:

```
"agentPoolProfiles": [
      {
        "name": "agentpool1",
        "count": 3,
        "vmSize": "Standard_D4s_v3",
        "availabilityProfile": "AvailabilitySet",
        "diskSizesGB": [1023]
      }
    ],
```

<a name="feat-private-cluster"></a>

## Private Cluster

You can build a private Kubernetes cluster with no public IP addresses assigned by setting:

```
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

```
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

To enable this feature, add `encryptionWithExternalKms` in `kubernetesConfig` and `objectId` in `servicePrincipalProfile`:

```
"kubernetesConfig": {
  "enableEncryptionWithExternalKms": true
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
