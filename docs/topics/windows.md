# Kubernetes Windows Walkthrough

<!-- TOC -->

- [Quick Start](#quick-start)
  - [Install Needed Tools](#install-needed-tools)
    - [Windows](#windows)
    - [Mac](#mac)
    - [Linux](#linux)
  - [Create a Resource Group and Service Principal](#create-a-resource-group-and-service-principal)
    - [Create a Resource Group and Service Principal (Windows)](#create-a-resource-group-and-service-principal-windows)
    - [Create a Resource Group and Service Principal (Mac+Linux)](#create-a-resource-group-and-service-principal-maclinux)
  - [Create an AKS Engine apimodel](#create-an-aks-engine-apimodel)
    - [Filling out API model (Windows)](#filling-out-apimodel-windows)
    - [Filling out API model (Mac & Linux)](#filling-out-apimodel-mac--linux)
  - [Generate Azure Resource Manager template](#generate-azure-resource-manager-template)
  - [Deploy the cluster](#deploy-the-cluster)
    - [Check that the cluster is up](#check-that-the-cluster-is-up)
  - [Deploy your first application](#deploy-your-first-application)
  - [What was deployed](#what-was-deployed)
- [Next Steps](#next-steps)

<!-- /TOC -->

## Quick Start

This guide will step through everything needed to build your first Kubernetes cluster and deploy a Windows web server on it. The steps include:

- Getting the right tools
- Completing an AKS Engine API model which describes what you want to deploy
- Running AKS Engine to generate Azure Resource Model templates
- Deploying your first Kubernetes cluster with Windows Server 2019 nodes
- Managing the cluster from your Windows machine
- Deploying your first app on the cluster

All of these steps can be done from any OS platform, so some sections are split out by Windows, Mac or Linux to provide the most relevant samples and scripts. If you have a Windows machine but want to use the Linux tools - no problem! Set up the [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/about) and you can follow the Linux instructions on this page.

When you're done, you will have a working cluster running Windows Server 2019 and the latest Kubernetes v1.13 release. Older releases are available but not recommended.

> Note: Windows support for Kubernetes is still in beta and under **active development**. If you run into problems, please be sure to check the [Troubleshooting](../howto/troubleshooting.md) page and [active Windows issues](https://github.com/azure/aks-engine/issues?&q=is:issue+is:open+label:windows) in this repo, then help us by filing new issues for things that aren't already covered.

### Install Needed Tools

This guide needs a few important tools, which are available on Windows, Mac, and Linux:

- AKS Engine - used to generate the Azure Resource Manager (ARM) template to automatically deploy a Kubernetes cluster
- Azure CLI - used to log into Azure, create resource groups, and deploy a Kubernetes cluster from a template
- Kubectl - "Kube control" tool used to manage Kubernetes clusters
- SSH - A SSH public key is needed when you deploy a cluster. It's used to connect to the Linux VMs running the cluster if you need to do more  management or troubleshooting later.

#### Windows

##### Azure CLI (Windows)

Click the [download](https://aka.ms/installazurecliwindows) link, and choose "Run". Click through the setup steps as needed.

Once it's installed, make sure you can connect to Azure with it. Open a new PowerShell window, then run `az login`. It will have you log in to Azure in your web browser, then return back to the command line and show "You have logged in. Now let us find all the subscriptions to which you have access..." along with the list of subscriptions.

> If you want other versions, check out the [official instructions](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest). For more help, check out the Azure CLI [getting started](https://docs.microsoft.com/en-us/cli/azure/get-started-with-azure-cli?view=azure-cli-latest) page.

##### AKS Engine (Windows)

Windows support is evolving rapidly, so be sure to use the latest AKS Engine  version (v0.20 or later).

1. Browse to the AKS Engine [releases page](https://github.com/Azure/aks-engine/releases) on GitHub.

2. Find the latest version, and download the file ending in `-windows-amd64.zip`.

3. Extract the `aks-engine...-windows-amd64.zip` file to a working folder such as `c:\tools`

4. Check that it runs with `.\aks-engine.exe version`

```console
PS C:\Users\patrick\aks-engine> .\aks-engine.exe version
Version: v0.27.0
GitCommit: f895db7f
GitTreeState: clean
```

5. Add the folder you created in step 3 to your path.

```powershell
$ENV:Path += ';c:\tools'
# If you want to save the setting permanently, then run
$oldPath = [Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::User)
[Environment]::SetEnvironmentVariable('Path', $oldPath + ';c:\tools', [EnvironmentVariableTarget]::User)
```

##### Kubectl (Windows)

The latest release of Kubernetes Control (kubectl) is available on the [Kubernetes release page](https://kubernetes.io/docs/setup/release/notes/). Look for `kubernetes-client-windows-amd64.tar.gz` and download it.

Windows 10 version 1803 and later already include `tar`, so extract the archive and move `kubectl.exe` to the same folder (such as `c:\tools`) that you put `aks-engine.exe`. If you don't already have `tar`, then [busybox-w32](https://frippery.org/busybox/) is a good alternative. Download [busybox.exe](https://frippery.org/files/busybox/busybox.exe), then copy it to `c:\tools\tar.exe`. It must be named to `tar.exe` for the next step to work.

```powershell
tar xvzf C:\Users\patrick\Downloads\kubernetes-client-windows-amd64.tar.gz
Move-Item .\kubernetes\client\bin\kubectl.exe c:\tools
```

##### SSH (Windows)

Windows 10 version 1803 and later come with the Secure Shell (SSH) client as an optional feature installed at `C:\Windows\system32\openssh`. If you have `ssh.exe` and `ssh-keygen.exe` there, skip forward to [Generate SSH key (Windows)](#generate-ssh-key-windows)

1. Download the latest OpenSSH-Win64.zip file from [Win32-OpenSSH releases](https://github.com/PowerShell/Win32-OpenSSH/releases)
2. Extract it to the same `c:\tools` folder or another folder in your path

###### Generate SSH key (Windows)

First, check if you already have a SSH key generated at `~\.ssh\id_rsa.pub`

```powershell
dir ~\.ssh\id_rsa.pub
dir : Cannot find path 'C:\Users\patrick\.ssh\id_rsa.pub' because it does not exist.
```

If the file already exists, then you can skip forward to [Create a Resource Group and Service Principal](#create-a-resource-group-and-service-principal).

If it does not exist, then run `ssh-keygen.exe`. Use the default file, and enter a passphrase if you wish to protect it. Be sure not to use a SSH key with blank passphrase in production.

```powershell
PS C:\Users\patrick\aks-engine> ssh-keygen.exe
Generating public/private rsa key pair.
Enter file in which to save the key (C:\Users\patrick/.ssh/id_rsa):
Created directory 'C:\Users\patrick/.ssh'.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in C:\Users\patrick/.ssh/id_rsa.
Your public key has been saved in C:\Users\patrick/.ssh/id_rsa.pub.
The key fingerprint is:
SHA256:... patrick@plang-g1
The key's randomart image is:
+---[RSA 2048]----+
...
+----[SHA256]-----+
```

#### Mac

Most of the needed tools are available with [Homebrew](https://brew.sh/). Use it or another package manager to install these:

- `jq` - helpful JSON processor
- `azure-cli` - for the `az` Azure command line tool
- `kubernetes-cli` - for the `kubectl` "Kube Control" management tool

Once you have those installed, make sure you can log into Azure. Open a new Terminal window, then run `az login`. It will have you log in to Azure in your web browser, then return back to the command line and show "You have logged in. Now let us find all the subscriptions to which you have access..." along with the list of subscriptions.

##### AKS Engine (Mac)

Windows support is evolving rapidly, so be sure to use the latest AKS Engine version (v0.20 or later).

1. Browse to the AKS Engine [releases page](https://github.com/Azure/aks-engine/releases) on GitHub.

2. Find the latest version, and download the file ending in `-darwin-amd64.zip`.

3. Extract the `aks-engine...-darwin-amd64.zip` file to a folder in your path such as `/usr/local/bin`

4. Check that it runs with `aks-engine version`

```console
$ aks-engine.exe version
Version: v0.27.0
GitCommit: f895db7f
GitTreeState: clean
```

##### SSH (Mac)

SSH is preinstalled, but you may need to generate an SSH key.

###### Generate SSH key (Mac)

Open up Terminal, and make sure you have a SSH public key

```console
$ ls ~/.ssh/id_rsa.pub
/home/patrick/.ssh/id_rsa.pub
```

If the file doesn't exist, run `ssh-keygen` to create one.

#### Linux

These tools are included in most distributions. Use your typical package manager to make sure they're installed:

- `jq` - helpful JSON processor
- `curl` - to download files
- `openssh` or another `ssh` client
- `tar`

##### Azure CLI (Linux)

Packages for the `az` cli are available for most distributions. Please follow the right link for your package manager:
[apt](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-apt?view=azure-cli-latest),
 [yum](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-yum?view=azure-cli-latest),
 [zypper](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-zypper?view=azure-cli-latest)

Now, make sure you can log into Azure. Open a new Terminal window, then run `az login`. It will have you log in to Azure in your web browser, then return back to the command line and show "You have logged in. Now let us find all the subscriptions to which you have access..." along with the list of subscriptions.

##### AKS Engine (Linux)

Windows support is evolving rapidly, so be sure to use the latest AKS Engine version (v0.20 or later).

1. Browse to the AKS Engine [releases page](https://github.com/Azure/aks-engine/releases) on GitHub.

2. Find the latest version, and download the file ending in `-linux-amd64.zip`.

3. Extract the `aks-engine...-linux-amd64.zip` file to a folder in your path such as `/usr/local/bin`

4. Check that it runs with `aks-engine version`

```console
$ aks-engine.exe version
Version: v0.27.0
GitCommit: f895db7f
GitTreeState: clean
```

##### Kubectl (Linux)

The latest release of Kubernetes Control (kubectl) is available on the [Kubernetes release page](https://kubernetes.io/docs/setup/release/notes/). Look for `kubernetes-client-linux-....tar.gz` and copy the link to it.

Download and extract it with curl & tar:

```console
$ curl -L https://dl.k8s.io/v1.13.1/kubernetes-client-linux-amd64.tar.gz | tar xvzf -
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    304      0 --:--:-- --:--:-- --:--:--   304
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0kubernetes/
kubernetes/client/
kubernetes/client/bin/
kubernetes/client/bin/kubectl
100 13.2M  100 13.2M    0     0  5608k      0  0:00:02  0:00:02 --:--:-- 8034k
```

Then copy it to `/usr/local/bin` or another directory in your `PATH`

```console
$ sudo cp kubernetes/client/bin/kubectl /usr/local/bin/
```

##### Generate SSH key (Linux)

From a terminal, make sure you have a SSH public key

```console
$ ls ~/.ssh/id_rsa.pub
/home/patrick/.ssh/id_rsa.pub
```

If the file doesn't exist, run `ssh-keygen` to create one.

### Create a Resource Group and Service Principal

Now that we have the Azure CLI configured and a SSH key generated, it's time to create a resource group to hold the deployment.

AKS Engine and Kubernetes also need access to deploy resources inside that resource group to build the cluster, as well as configure more resources such as Azure Load Balancers once the cluster is running. This is done using an Azure Service Principal. It's safest to create one with access just to the resource group so that once your deployment is deleted, the service principal can't be used to make other changes in your subscription.

#### Create a Resource Group and Service Principal (Windows)

`az group create --location <location> --name <name>` will create a group for you. Be sure to use a unique name for each cluster. If you need a list of available locations, run `az account list-locations -o table`.

```powershell
PS C:\Users\patrick\aks-engine> az group create --location westus2 --name k8s-win1
{
  "id": "/subscriptions/df392461-0000-1111-2222-cd3aa2d911a6/resourceGroups/k8s-win1",
  "location": "westus2",
  "managedBy": null,
  "name": "k8s-win1",
  "properties": {
    "provisioningState": "Succeeded"
  },
  "tags": null
}
```

Now that the group is created, create a service principal with Contributor access for that group only

```powershell
# Get the group id
$groupId = (az group show --resource-group k8s-win1 --query id --output tsv)

# Create the service principal
$sp = az ad sp create-for-rbac --role="Contributor" --scopes=$groupId | ConvertFrom-JSON
```

#### Create a Resource Group and Service Principal (Mac+Linux)

`az group create --location <location> --name <name>` will create a group for you. Be sure to use a unique name for each cluster. If you need a list of available locations, run `az account list-locations -o table`.

```console
$ export RESOURCEGROUP=k8s-win1
$ export LOCATION=westus2
$ az group create --location $LOCATION --name $RESOURCEGROUP
```

Now that the group is created, create a service principal with Contributor access for that group only

```console
# Get the group id
$ export RESOURCEGROUPID=$(az group show --resource-group $RESOURCEGROUP --query id --output tsv)

# Create the service principal
$ export SERVICEPRINCIPAL=$(az ad sp create-for-rbac --role="Contributor" --scopes=$RESOURCEGROUPID --output json)
```

### Create an aks-engine apimodel

Multiple samples are available in this repo under [examples/windows](../../examples/windows/). This guide will use the [windows/kubernetes.json](../../examples/windows/kubernetes.json) sample to deploy 1 Linux VM to run Kubernetes services, and 2 Windows nodes to run your Windows containers.

After downloading that file, you will need to

1. Set windowsProfile.adminUsername and adminPassword. Be sure to check the Azure Windows VM [username](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/faq?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json#what-are-the-username-requirements-when-creating-a-vm) and [password](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/faq?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json#what-are-the-password-requirements-when-creating-a-vm) requirements first.
1. Set a unique name for masterProfile.dnsPrefix. This will be the first part of the domain name you'll use to manage the Kubernetes cluster later
1. Set the ssh public key that will be used to log into the Linux VM
1. Set the Azure service principal for the deployments

#### Filling out API model (Windows)

You can use the same PowerShell window from earlier to run this next script to do all that for you. Be sure to replace `$dnsPrefix` with something unique and descriptive, `$windowsUser` and `$windowsPassword` to meet the requirements.

```powershell
# Be sure to change these next 3 lines for your deployment
$dnsPrefix = "wink8s1"
$windowsUser = "winuser"
$windowsPassword = "Cr4shOverride!"

# Download template
Invoke-WebRequest -UseBasicParsing https://raw.githubusercontent.com/Azure/aks-engine/master/examples/windows/kubernetes.json -OutFile kubernetes-windows.json

# Load template
$inJson = Get-Content .\kubernetes-windows.json | ConvertFrom-Json

# Set dnsPrefix
$inJson.properties.masterProfile.dnsPrefix = $dnsPrefix

# Set Windows username & password
$inJson.properties.windowsProfile.adminPassword = $windowsPassword
$inJson.properties.windowsProfile.adminUsername = $windowsUser

# Copy in your SSH public key from `~/.ssh/id_rsa.pub` to linuxProfile.ssh.publicKeys.keyData
$inJson.properties.linuxProfile.ssh.publicKeys[0].keyData = [string](Get-Content "~/.ssh/id_rsa.pub")

# Set servicePrincipalProfile
$inJson.properties.servicePrincipalProfile.clientId = $sp.appId
$inJson.properties.servicePrincipalProfile.secret = $sp.password

# Save file
$inJson | ConvertTo-Json -Depth 5 | Out-File -Encoding ascii -FilePath "kubernetes-windows-complete.json"
```

#### Filling out API model (Mac & Linux)

Using the same terminal as before, you can use this script to download the template and fill it out. Be sure to set DNSPREFIX, WINDOWSUSER, and WINDOWSPASSWORD to meet the requirements.

```console
$ export DNSPREFIX="wink8s1"
$ export WINDOWSUSER="winuser"
$ export WINDOWSPASSWORD="Cr4shOverride!"
$ curl -L https://raw.githubusercontent.com/Azure/aks-engine/master/examples/windows/kubernetes.json -o kubernetes.json
$ cat kubernetes.json | \
  jq ".properties.masterProfile.dnsPrefix = \"$DNSPREFIX\"" | \
  jq ".properties.linuxProfile.ssh.publicKeys[0].keyData = \"`cat ~/.ssh/id_rsa.pub`\"" | \
  jq ".properties.servicePrincipalProfile.clientId = `echo $SERVICEPRINCIPAL | jq .appId`" | \
  jq ".properties.servicePrincipalProfile.secret = `echo $SERVICEPRINCIPAL | jq .password`" | \
  jq ".properties.windowsProfile.adminPassword = \"$WINDOWSPASSWORD\"" | \
  jq ".properties.windowsProfile.adminUsername = \"$WINDOWSUSER\"" > kubernetes-windows-complete.json
```

### Generate Azure Resource Manager template

Now that the AKS Engine cluster definition is complete, generate the Azure templates with `aks-engine generate kubernetes-windows-complete.json`

```console
$ aks-engine.exe generate kubernetes-windows-complete.json
INFO[0000] Generating assets into _output/wink8s1...
```

This will generate a `_output` directory with a subdirectory named after the dnsPrefix you set above. In this example, it's `_output/wink8s1`.

It will also create a working Kubernetes client config file in `_output/<dnsprefix>/kubeconfig` folder. We'll come back to that in a bit.

### Deploy the cluster

Get the paths to `azuredeploy.json` and `azuredeploy.parameters.json` from the last step, and pass them into `az deployment group create --name <name for deployment> --resource-group <resource group name> --template-file <...azuredeploy.json> --parameters <...azuredeploy.parameters.json>`

```console
$ az deployment group create --name k8s-win1-deploy --resource-group k8s-win1 --template-file "./_output/wink8s1/azuredeploy.json" --parameters "./_output/wink8s1/azuredeploy.parameters.json"
```

After several minutes, it will return the list of resources created in JSON. Look for `masterFQDN`.

```json
      "masterFQDN": {
        "type": "String",
        "value": "wink8s1.westus2.cloudapp.azure.com"
      },
```

The DNS prefix provided is used to generate the hostname of the cluster (e.g. staging, prodwest, blueberry) and must be unique for each cluster deployment.

#### Check that the cluster is up

As mentioned earlier, `aks-engine generate` also creates Kubernetes configuration files under `_output/<dnsprefix>/kubeconfig`. There will be one per possible region, so find the one matching the region you deployed in.

In the example above with `dnsprefix`=`wink8s1` and the `westus2` region, the filename would be `_output/wink8s1/kubeconfig/kubeconfig.westus2.json`.

##### Setting KUBECONFIG on Windows

Set `$ENV:KUBECONFIG` to the full path to that file.

```powershell
$ENV:KUBECONFIG=(Get-Item _output\wink8s1\kubeconfig\kubeconfig.westus2.json).FullName
```

##### Setting KUBECONFIG on Mac or Linux

```console
$ export KUBECONFIG=$(PWD)/_output/$DNSPREFIX/kubeconfig/kubeconfig.westus2.json
```

Once you have `KUBECONFIG` set, you can verify the cluster is up with `kubectl get node -o wide`.

```console
$ kubectl get node -o wide
NAME                    STATUS    ROLES     AGE       VERSION   EXTERNAL-IP   OS-IMAGE                    KERNEL-VERSION   CONTAINER-RUNTIME
40336k8s9000            Ready     <none>    21m       v1.9.10   <none>        Windows Server Datacenter   10.0.17134.112
                        docker://17.6.2
40336k8s9001            Ready     <none>    20m       v1.9.10   <none>    Windows Server Datacenter   10.0.17134.112
                        docker://17.6.2
k8s-master-40336153-0   Ready     master    22m       v1.9.10   <none>    Ubuntu 16.04.5 LTS   4.15.0-1018-azure   docker://1.13.1
```

##### SSH to the Linux master (optional)

If you would like to manage the cluster over SSH, you can connect to the Linux master directly using the FQDN of the cluster:

```console
$ ssh azureuser@plangk8swin1.westus2.cloudapp.azure.com
```

### Deploy your first application

Kubernetes deployments are typically written in YAML files. This one will create a pod with a container running the IIS web server, and tell Kubernetes to expose it as a service with the Azure Load Balancer on an external IP.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: iis-2019
  labels:
    app: iis-2019
spec:
  replicas: 1
  template:
    metadata:
      name: iis-2019
      labels:
        app: iis-2019
    spec:
      containers:
      - name: iis
        image: mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019
        resources:
          limits:
            cpu: 1
            memory: 800m
          requests:
            cpu: .1
            memory: 300m
        ports:
          - containerPort: 80
      nodeSelector:
        "beta.kubernetes.io/os": windows
  selector:
    matchLabels:
      app: iis-2019
---
apiVersion: v1
kind: Service
metadata:
  name: iis
spec:
  type: LoadBalancer
  ports:
  - protocol: TCP
    port: 80
  selector:
    app: iis-2019
```

Copy and paste that into a file called `iis.yaml`, then run `kubectl apply -f iis.yaml`. kubectl will show the deployment and service were created:

```console
$ kubectl apply -f .\iis.yaml

deployment.apps/iis-1809 created
service/iis created
```

Now, you can check the status of the pod and service with `kubectl get pod` and `kubectl get service` respectively.

Initially, the pod will be in the `ContainerCreating` state, and eventually go to `Running`. The service will show `<pending>` under `EXTERNAL-IP`. Here's what the first progress will look like:

```console
$ kubectl get pods
NAME                        READY     STATUS              RESTARTS   AGE
iis-1809-6c49777598-h45cs   0/1       ContainerCreating   0          1m

$ kubectl get services
NAME         TYPE           CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
iis          LoadBalancer   10.0.9.47    <pending>     80:31240/TCP   1m
kubernetes   ClusterIP      10.0.0.1     <none>        443/TCP        46m
```

Since this is the first deployment, it will probably take several minutes for the Windows node to download and run the container. Later deployments will be faster because the large `microsoft/windowsservercore` container will already be on disk.

The service will eventually show an EXTERNAL-IP as well:

```console
$ kubectl get service
NAME         TYPE           CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
iis          LoadBalancer   10.0.9.47    13.66.203.178 80:31240/TCP   1m
kubernetes   ClusterIP      10.0.0.1     <none>        443/TCP        46m
```

Once the pod is in `Running` state, get the IP from `kubectl get service` then visit `http://<EXTERNAL-IP>` to test your web server.

### What was deployed

Once your Kubernetes cluster has been created you will have a resource group containing:

1. 1 master accessible by SSH on port 22 or kubectl on port 443

2. A set of Windows and/or Linux nodes.  The windows nodes can be accessed through an RDP SSH tunnel via the master node, following these steps [Connecting to Windows Nodes](../howto/troubleshooting.md#connecting-to-windows-nodes).

![Image of Kubernetes cluster on azure with Windows](../static/img/kubernetes-windows.png)

These parts were all automatically created using the Azure Resource Manager template created by AKS Engine:

1. **Master Components** - The master runs the Kubernetes scheduler, api server, and controller manager.  Port 443 is exposed for remote management with the kubectl cli.
2. **Linux Nodes** - the Kubernetes nodes run in an availability set.  Azure load balancers are dynamically added to the cluster depending on exposed services.
3. **Windows Nodes** - the Kubernetes windows nodes run in an availability set.
4. **Common Components** - All VMs run a kubelet, Docker, and a Proxy.
5. **Networking** - All VMs are assigned an ip address in the 10.240.0.0/16 network and are fully accessible to each other.

## Next Steps

For more resources on Windows and AKS Engine, continue reading:

- [Customizing Windows Deployments](windows-and-kubernetes.md#customizing-windows-deployments)
- [More Examples](windows-and-kubernetes.md#more-examples)
- [Troubleshooting](../howto/troubleshooting.md)
- [Using Kubernetes ingress](../howto/mixed-cluster-ingress.md) for more flexibility in http and https routing

If you'd like to learn more about Kubernetes in general, check out these guides:

1. [Kubernetes Bootcamp](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html) - shows you how to deploy, scale, update, and debug containerized applications.
2. [Kubernetes Userguide](http://kubernetes.io/docs/user-guide/) - provides information on running programs in an existing Kubernetes cluster.
