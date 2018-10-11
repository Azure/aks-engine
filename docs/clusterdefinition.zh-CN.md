# 微软Azure容器服务引擎 - 集群定义文件

## apiVersion为"vlabs"的集群定义文件

以下是apiVersion为"vlabs"时对应的集群定义文件配置说明：

### apiVersion

|名称|是否必须|说明|
|---|---|---|
|apiVersion|是|集群定义文件的版本。比如“vlabs”版本的api对应的值就是“vlabs”|

### orchestratorProfile

`orchestratorProfile` 字段包含了编排引擎的各种设置。

|名称|是否必须|说明|
|---|---|---|
|orchestratorType|是|这个字段指定了ACS引擎使用的编排引擎的类型。|

可选的编排引擎如下所示：

1. `Kubernetes` - 指定编排引擎为 [Kubernetes编排引擎](kubernetes.md)。

### masterProfile
`masterProfile` 指定了集群中master节点的各种配置。

|名称|是否必须|说明|
|---|---|---|
|count|是|集群中master的节点可以指定为1，3，5。|
|dnsPrefix|是|指定master节点的FQDN的dns前缀。当使用ssh或者命令行工具连接master的时候就需要用到这个FQDN。这个字段必须是一个唯一值。([用户自定义VNET的例子](../examples/vnet))|
|firstConsecutiveStaticIP|只有当vnetSubnetId被设置是需要|指定第一个master节点的IP地址。后续的master节点的IP地址会根据这个值向后累加。|
|vmsize|是|具体的值请参考 [Azure虚机规格](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/).  所选的虚机的规格必须最少有2个CPU核心和100GB的磁盘空间。|
|vnetSubnetId|否|指定一个备用的VNET子网的ID。这个子网必须具有一个正确的VNET ID，并且处于同一个订阅中。([用户自定义VNET的例子](../examples/vnet))|

### agentPoolProfiles
一个集群可以拥有0到12个agent pool配置。agent pool配置用来指定创建各种资源比如虚机，虚机规模集或者高可用集，Public/Private access，[attached storage disks](../examples/disks-storageaccount), [attached managed disks](../examples/disks-managed), 或者 [Windows](../examples/windows).

|名称|是否必须|说明|
|---|---|---|
|availabilityProfile|否, 默认值为`VirtualMachineScaleSets`| 可以指定`VirtualMachineScaleSets`或者`AvailabilitySet`。一般来说可以选择`VirtualMachineScaleSets`，如果需要使用Kubernetes集群或者动态挂载磁盘等功能的话可以选择其他的配置。|
|count|是|指定集群中agent节点的数量|
|diskSizesGB|否|指定一个数组，其中包含了4个或4个以内的磁盘大小。每个磁盘的大小在1到1024之间（GB）。|
|dnsPrefix|当agent节点需要通过一个外部的负载均衡暴露给外部网络时使用。|这个值用来创建agent节点的FQDN，然后外部网络就可以通过这个FQDN远程连接到agent节点。这个字段必须是一个唯一值。|
|name|是|agent pool配置的唯一名称。agent pool中各种资源的名称都基于这个名称创建。|
|ports|当且仅当需要将服务暴露给外部网络时使用。|指定一个数组，其中包含了需要暴露给外部的端口号。每个端口都会配置一个tcp probe，并且只能被agent节点访问。最大可以指定150个端口。|
|storageProfile|否, 默认值为`StorageAccount`|指定存储配置。可选值为[StorageAccount](../examples/disks-storageaccount)或者[ManagedDisks](../examples/disks-managed)|
|vmsize|是|指定虚机的规格[Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/).虚机最小必须是2个CPU核心。|
|vnetSubnetId|否|指定备用的VNET子网的ID。这个子网必须具有一个正确的VNET ID，并且处于同一个订阅中。([用户自定义VNET的例子](../examples/vnet))|

### linuxProfile

`linuxProfile`用来指定集群中每个linux节点的配置。

|名称|是否必须|说明|
|---|---|---|
|adminUsername|是|指定linux节点的用户名。|
|ssh.publicKeys.keyData|是|指定授权认证linux节点的ssh公钥。如何创建ssh公钥请参考[生成ssh公钥私钥对](ssh.md#ssh-key-generation).|
|secrets|否|指定一个key vaults数组，通过这个数组从key vault中获取对应的密钥。|

#### secrets
`secrets`指定在集群中master和agent节点上安装的证书。

一个集群可以有多个key vaults用来安装证书。

在linux节点上，证书被安装在"/var/lib/waagent/"目录中。每个证书有两个文件：

1. `{thumbprint}.crt` : PEM格式的证书链。
2. `{thumbprint}.prv` : PEM格式的私钥。

|名称|是否必须|说明|
|---|---|---|
|sourceVault.id|是|key vault的ARM ID。通过这个ID从key vault中获取对应的密钥。|
|vaultCertificates.certificateUrl|是|证书对应的key vault的url以及版本。|
`sourceVault.id`的格式可以通过cli来获取，或者在Azure门户中的：/subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.KeyVault/vaults/{keyvaultname}中获取。

`vaultCertificates.certificateUrl`的格式可以通过cli获取，也可以通过Azure门户中的：https://{keyvaultname}.vault.azure.net:443/secrets/{secretName}/{version}获取。

### servicePrincipalProfile

`servicePrincipalProfile`指定了集群用来配置资源时需要的密钥。更多的信息请参考[service principal](serviceprincipal.md)。

|名称|是否必须|说明|
|---|---|---|
|clientId|当指定编排引擎为kubernetes时需要。|指定了Azure client id. 这里建议针对不同的集群使用不同的client ID|
|secret|当指定编排引擎为kubernetes时需要。|指定了Azure client secret.  这里建议针对不同的集群使用不同的client secret。|

## "2016-03-30"版本apiVersion的集群定义文件

以下是"2016-03-30"版本apiVersion的集群定义文件，这个版本的api和Azure Kubernetes Service Engine的一致。

### apiVersion

|名称|是否必须|说明|
|---|---|---|
|apiVersion|是|集群定义文件的版本。例如"2016-03-30"版本的api对应的值为"2016-03-30"。|

### orchestratorProfile
`orchestratorProfile` 字段包含了编排引擎的各种设置。

|名称|是否必须|说明|
|---|---|---|
|orchestratorType|是|这个字段指定了ACS引擎使用的编排引擎的类型。|

可选的编排引擎如下所示：

1. `Kubernetes` - 指定编排引擎为 [Kubernetes编排引擎](kubernetes.md)。

### masterProfile
`masterProfile` 指定了集群中master节点的各种配置。

|名称|是否必须|说明|
|---|---|---|
|count|是|集群中master的节点可以指定为1，3，5。|
|dnsPrefix|是|指定master节点的FQDN的dns前缀。当使用ssh或者命令行工具连接master的时候就需要用到这个FQDN。这个字段必须是一个唯一值。([用户自定义VNET的例子](../examples/vnet))|

### agentPoolProfiles
对于"2016-03-30"版本apiVersion的集群定义文件，一个集群中只能包含一个agent pool配置。

|名称|是否必须|说明|
|---|---|---|
|count|是|指定集群中agent节点的数量|
|dnsPrefix|当agent节点需要通过一个外部的负载均衡暴露给外部网络时使用。|这个值用来创建agent节点的FQDN，然后外部网络就可以通过这个FQDN远程连接到agent节点。这个字段必须是一个唯一值。|
|name|是|agent pool配置的唯一名称。agent pool中各种资源的名称都基于这个名称创建。|
|vmsize|是|指定虚机的规格[Azure VM Sizes](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-windows-sizes/).虚机最小必须是2个CPU核心。|

### linuxProfile
`linuxProfile`用来指定集群中每个linux节点的配置。

|名称|是否必须|说明|
|---|---|---|
|adminUsername|是|指定linux节点的用户名。|
|ssh.publicKeys[0].keyData|是|指定授权访问linux节点的ssh公钥。如何创建ssh公钥请参考[生成ssh公钥私钥对](ssh.md#ssh-key-generation).|
