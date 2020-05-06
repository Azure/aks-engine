# 微软Azure容器服务引擎 - 构建基于Docker的容器集群

## 概览

微软容器服务引擎（`aks-engine`）用于将一个容器集群描述文件转化成一组ARM（Azure Resource Manager）模板，通过在Azure上部署这些模板，用户可以很方便地在Azure上建立一套基于Docker的容器服务集群。用户可以自由地选择集群编排引擎DC/OS, Kubernetes或者是Swarm/Swarm Mode。集群描述文件使用和ARM模板相同的语法，它们都可以用来部署Azure容器服务。

集群描述文件提供了一下几个功能：
* 可以自由选择DC/OS, Kubernetes, Swarm Mode和Swarm等多种编排引擎
* 可以自由定制集群节点的规格，包括：
    * 虚机的规格
    * 节点的数量
    * 弹性虚拟机集，高可用服务集
    * 存储设备，托管存储
* 可创建高达1200的容器集群节点数量
* 自定义VNET

## 演示链接

* [AKS Engine](docs/aksengine.md) - 演示如何使用ACS引擎来生成基于Docker的容器集群
* [Cluster Definition](docs/clusterdefinition.md) - 详细介绍集群描述文件的格式
* [DC/OS Walkthrough](docs/dcos.md) - 演示如何使用ACS引擎在Azure上创建DC/OS集群
* [Kubernetes Walkthrough](docs/kubernetes.md) - 演示如何使用ACS引擎在Azure上创建Kubernetes集群
* [Swarm Walkthrough](docs/swarm.md) - 演示如何使用ACS引擎在Azure上创建Swarm集群
* [Swarm Mode Walkthrough](docs/swarmmode.md) - 演示如何使用ACS引擎在Azure上创建Swarm Mode集群
* [Custom VNET](examples/vnet) - 演示如何在用户自定义VNET上创建容器集群
* [Attached Disks](examples/disks-storageaccount) - 演示如何在一个集群节点上创建4个磁盘
* [Managed Disks](examples/disks-managed) - 演示如何管理托管磁盘
* [Large Clusters](examples/largeclusters) - 演示如何创建1200个节点的容器集群

## 提交代码

请按照以下流程提交您的PR：

1. 提交之前请运行`make ci`命令确保测试用例通过

2. 如果您改动了部署模板的话，请务必手动部署您的改动到Azure来确保项目正常运行。

## 使用步骤

运行完毕后，项目的根目录下就会产生一个`_output/`的文件夹，这个文件夹中包含了所有的ARM模板，通过部署这些模板就可以在Azure上创建对应的容器集群了。（对于kubernetes来说，_output文件夹中也会生成一些证书之类的文件来供ARM部署时的需要）

## 部署方法

可以使用如下几种方式来部署ARM模板：
[the Azure CLI 2.0](https://github.com/Azure/azure-cli)，
[Powershell](https://github.com/Azure/azure-powershell).

### 使用Azure CLI 2.0部署
**NOTE:** Azure CLI 2.0目前任处于测试阶段，中国地区尚且无法使用。如果部署到国际版的Azure的话可以使用以下流程：

```bash
$ az login

$ az account set --subscription "<SUBSCRIPTION NAME OR ID>"

$ az group create \
    --name "<RESOURCE_GROUP_NAME>" \
    --location "<LOCATION>"

$ az deployment group create \
    --name "<DEPLOYMENT NAME>" \
    --resource-group "<RESOURCE_GROUP_NAME>" \
    --template-file "./_output/<INSTANCE>/azuredeploy.json" \
    --parameters "@./_output/<INSTANCE>/azuredeploy.parameters.json"
```

### 使用Powershell部署

```powershell
Add-AzureRmAccount

Select-AzureRmSubscription -SubscriptionID <SUBSCRIPTION_ID>

New-AzureRmResourceGroup `
    -Name <RESOURCE_GROUP_NAME> `
    -Location <LOCATION>

New-AzureRmResourceGroupDeployment `
    -Name <DEPLOYMENT_NAME> `
    -ResourceGroupName <RESOURCE_GROUP_NAME> `
    -TemplateFile _output\<INSTANCE>\azuredeploy.json `
    -TemplateParameterFile _output\<INSTANCE>\azuredeploy.parameters.json
```

## 项目说明

本项目遵循[Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). 详细信息请参阅 [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq)或者联系[opencode@microsoft.com](mailto:opencode@microsoft.com)获取更多的技术支持。
