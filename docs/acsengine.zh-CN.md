# 微软Azure容器服务引擎

微软容器服务引擎（`aks-engine`）用于将一个容器集群描述文件转化成一组ARM（Azure Resource Manager）模板，通过在Azure上部署这些模板，用户可以很方便地在Azure上建立一套基于Docker的容器服务集群。用户可以自由地选择集群编排引擎DC/OS, Kubernetes或者是Swarm/Swarm Mode。集群描述文件使用和ARM模板相同的语法，它们都可以用来部署Azure容器服务。

# 基于Docker的部署

最简单的开始使用`aks-engine`的方式是使用Docker。如果本地计算机安装了Docker或者windows、Mac版本的Docker的话，无需安装任何软件就可以直接使用`aks-engine`了。

* Windows (PowerShell): `.\scripts\devenv.ps1`
* Linux (bash): `./scripts/devenv.sh`

上面的这段脚本在Docker容器中挂载了`aks-engine`源目录。你可以在任何熟悉的编辑器上修改这些源代码，所做的修改可以直接在Docker容器中编译和测试（本项目的持续集成系统中也采用了同样的方式）。

```
make bootstrap
```

当`devenv.{ps1,sh}`执行完毕的时候，你可以在容器中查看对应的日志，最后执行下面的脚本就可以生成`aks-engine`工具了：

```
make build
```

当项目编译通过后，可以使用如下的命令来验证`aks-engine`是否正常运行：

```
# ./bin/aks-engine
AKS Engine deploys and manages Kubernetes, Swarm Mode, and DC/OS clusters in Azure

Usage:
  aks-engine [command]

Available Commands:
  deploy        Deploy an Azure Resource Manager template
  generate      Generate an Azure Resource Manager template
  help          Help about any command
  orchestrators Display info about supported orchestrators
  scale         Scale an existing Kubernetes cluster
  upgrade       Upgrade an existing Kubernetes cluster
  version       Print the version of AKS Engine

Flags:
      --debug   enable verbose debug logs
  -h, --help    help for aks-engine

Use "aks-engine [command] --help" for more information about a command.
```

[详细的开发，编译，测试过程和步骤可以参考这个视频](https://www.youtube.com/watch?v=lc6UZmqxQMs)

# 本地下载并编译ACS引擎

ACS引擎具有跨平台特性，可以在windows，OS X和Linux上运行。以下是对应不同平台的安装步骤：

## Windows

安装依赖软件：
- Git for Windows. [点击这里下载安装](https://git-scm.com/download/win)
- Go for Windows. [点击这里下载安装](https://golang.org/dl/), 缺省默认安装.
- Powershell

编译步骤:

1. 设置工作目录。 这里假设使用`c:\gopath`作为工作目录：
  1. 使用Windows + R组合键打开运行窗口
  2. 执行命令：`rundll32 sysdm.cpl,EditEnvironmentVariables`打开系统环境变量设置对话框
  3. 添加`c:\go\bin`到PATH环境变量
  4. 点击“新建”按钮并新建GOPATH环境变量，设置缺省值为`c:\gopath`
2. 编译ACS引擎:
  1. 使用Windows + R组合键打开运行窗口
  2. 运行`cmd`命令打开命令行窗口
  3. 运行命令mkdir %GOPATH%
  4. cd %GOPATH%
  5. 运行`go get github.com/Azure/aks-engine`命令获取ACS引擎在github上的最新代码
  6. 运行`go get all`命令安装ACS引擎需要的依赖组件
  7. `cd %GOPATH%\src\github.com\Azure\aks-engine`
  8. 运行`go build`编译项目
3. 运行`aks-engine`命令，如果能看到命令参数提示就说明已经正确编译成功了。

## OS X

安装依赖软件：:
- Go for OS X. [点击这里下载安装](https://golang.org/dl/)

安装步骤:

1. 打开命令行窗口并设置GOPATH环境变量：
  1. `mkdir $HOME/gopath`
  2. 打开`$HOME/.bash_profile`文件并添加以下内容：
  ```
  export PATH=$PATH:/usr/local/go/bin
  export GOPATH=$HOME/gopath
  ```
  3. `source $HOME/.sh_profile`使配置生效。
2. 编译ACS引擎:
  1. 运行`go get github.com/Azure/aks-engine`命令获取ACS引擎在github上的最新代码。
  2. 运行`go get all`命令安装ACS引擎需要的依赖组件
  3. `cd $GOPATH/src/github.com/Azure/aks-engine`
  4. `go build`编译项目
3. 运行`aks-engine`命令，如果能看到命令参数提示就说明已经正确编译成功了。

## Linux

安装依赖软件：
- Go for Linux
  - [点击这里下载并安装](https://golang.org/dl/)
  - 执行命令sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz解压并替换原有文件。
- `git`

编译步骤:

1. 设置GOPATH:
  1. 运行命令`mkdir $HOME/gopath`新建gopath目录
  2. 编辑`$HOME/.profile`文件增加如下的配置：
  ```
  export PATH=$PATH:/usr/local/go/bin
  export GOPATH=$HOME/gopath
  ```
  3. 运行命令`source $HOME/.profile`使配置生效。
2. 编译ACS引擎:
  1. 运行命令`go get github.com/Azure/aks-engine`获取ACS引擎在github上的最新代码。
  2. 运行`go get all`命令安装ACS引擎需要的依赖组件
  3. `cd $GOPATH/src/github.com/Azure/aks-engine`
  4. 运行`go build`命令编译项目
3. 运行`aks-engine`命令，如果能看到命令参数提示就说明已经正确编译成功了。


# 生成模板

ACS引擎使用json格式的[集群定义文件](clusterdefinition.md)作为输入参数，生成3个或者多个类似如下的模板：

1. **apimodel.json** - 集群配置文件
2. **azuredeploy.json** - 核心的ARM (Azure Resource Model)模板，用来部署Docker集群
3. **azuredeploy.parameters.json** - 部署参数文件，其中的参数可以自定义
4. **certificate and access config files** - 某些编排引擎例如kubernetes需要生成一些证书，这些证书文件和它依赖的kube config配置文件也存放在和ARM模板同级目录下面

需要注意的是，当修改已有的Docker容器集群的时候，应该修改`apimodel.json`文件来保证最新的部署不会影响到目前集群中已有的资源。举个例子，如果一个容器集群中的节点数量不够的时候，可以修改`apimodel.json`中的集群节点数量，然后重新运行`aks-engine`命令并将`apimodel.json`作为输入参数来生成新的ARM模板。这样部署以后，集群中的旧的节点就不会有变化，新的节点会自动加入。

# 演示

这里通过部署一个kubernetes容器集群来演示如何使用`aks-engine`。kubernetes集群定义文件使用[examples/kubernetes.json](../examples/kubernetes.json)。

1. 首先需要准备一个[SSH 公钥私钥对](ssh.md#ssh-key-generation).
2. 编辑[examples/kubernetes.json](../examples/kubernetes.json)将其需要的参数配置好.
3. 运行`./bin/aks-engine generate examples/kubernetes.json`命令在_output/Kubernetes-UNIQUEID目录中生成对应的模板。（UNIQUEID是master节点的FQDN前缀的hash值）
4. 按照README中指定的方式使用`azuredeploy.json`和`azuredeploy.parameters.json`部署容器集群 [deployment usage](../acsengine.md#deployment-usage).

# 部署方法

[部署方式请参考这里](../acsengine.md#deployment-usage).
