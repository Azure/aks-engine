# More on Windows and Kubernetes

If you're trying to deploy Kubernetes with Windows the first time, be sure to check out the [quick start](windows.md) first. If you're looking for more details on deployments, examples or troubleshooting &mdash; read on.

## Customizing Windows deployments

AKS Engine allows a lot more customizations available in the [docs](../), but here are a few important ones you should know for Windows deployments. Each of these are extra parameters you can add into the AKS Engine API model file (such as `kubernetes-windows.json` from the quick start) before running `aks-engine generate`.

### Changing the OS disk size

The Windows Server deployments default to 30GB for the OS drive (C:), which is not enough to pull multiple `microsoft/windowsservercore`-based containers. It's easiest to start with 128GB, then see what your usage is over time before shrinking it down. You can change this size by adding `osDiskSizeGB` under the `agentPoolProfiles`, such as:

```json
"agentPoolProfiles": [
      {
        "name": "windowspool2",
        "count": 2,
        "vmSize": "Standard_D2_v3",
        "availabilityProfile": "AvailabilitySet",
        "osType": "Windows",
        "osDiskSizeGB": 128
     }
]
```

### Choosing the Windows Server version

If you want to deploy a specific Windows Server version, you can. First, find available versions with `az vm image list` command:

```console
$ az vm image list --publisher MicrosoftWindowsServer --all -o table

Offer                    Publisher                      Sku                                             Urn                                                                                                            Version
-----------------------  -----------------------------  ----------------------------------------------  -------------------------------------------------------------------------------------------------------------  -----------------
...
WindowsServerSemiAnnual  MicrosoftWindowsServer         Datacenter-Core-1709-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServerSemiAnnual:Datacenter-Core-1709-with-Containers-smalldisk:1709.0.20181017  1709.0.20181017
WindowsServerSemiAnnual  MicrosoftWindowsServer         Datacenter-Core-1803-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServerSemiAnnual:Datacenter-Core-1803-with-Containers-smalldisk:1803.0.20181017  1803.0.20181017
WindowsServerSemiAnnual  MicrosoftWindowsServer         Datacenter-Core-1809-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServerSemiAnnual:Datacenter-Core-1809-with-Containers-smalldisk:1809.0.20181107  1809.0.20181107
WindowsServer            MicrosoftWindowsServer         2019-Datacenter-Core-with-Containers-smalldisk  MicrosoftWindowsServer:WindowsServer:2019-Datacenter-Core-with-Containers-smalldisk:2019.0.20181107            2019.0.20181107
```

You can use the Offer, Publisher and Sku to pick a specific version by adding `windowsOffer`, `windowsPublisher`, `windowsSku` and (optionally) `imageVersion` to the `windowsProfile` section. In this example, the latest Windows Server version 1809 image would be deployed.

```json
"windowsProfile": {
            "adminUsername": "azureuser",
            "adminPassword": "...",
            "windowsPublisher": "MicrosoftWindowsServer",
            "windowsOffer": "WindowsServerSemiAnnual",
            "windowsSku": "Datacenter-Core-1809-with-Containers-smalldisk"
     },
```

### Disabling automatic updates

If you want to disable automatic Windows updates, you can use the `enableAutomaticUpdates` option.

```json
"windowsProfile": {
            "adminUsername": "azureuser",
            "adminPassword": "...",
            "windowsPublisher": "MicrosoftWindowsServer",
            "windowsOffer": "WindowsServerSemiAnnual",
            "windowsSku": "Datacenter-Core-1809-with-Containers-smalldisk",
            "enableAutomaticUpdates": false
     },
```

## More Examples

### Using Azure Files

For more background information, please check out [Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) in the Kubernetes documentation.

1. Create an azure file storage class

```console
$ kubectl apply -f https://raw.githubusercontent.com/JiangtianLi/Examples/master/windows/azurefile/storageclass-azurefile.yaml
```

2. Make sure storageclass is created successfully

```console
$ kubectl get storageclass/azurefile -o wide
```

3. Create a pvc for azure file

```console
kubectl apply -f https://raw.githubusercontent.com/JiangtianLi/Examples/master/windows/azurefile/pvc-azurefile.yaml
```

4. Make sure pvc is created successfully

```console
$ kubectl get pvc/pvc-azurefile -o wide
```

5. Create a pod with azure file pvc

```console
$ kubectl apply -f https://raw.githubusercontent.com/JiangtianLi/Examples/master/windows/azurefile/iis-azurefile.yaml
```

6. Watch the status of pod until its `STATUS` is `Running`

```
kubectl get po/iis-azurefile -o wide -w
```

7. Enter the pod container to validate

```console
kubectl exec -it iis-azurefile -- cmd
```

```cmd
C:\>dir c:\mnt\azure
 Volume in drive C has no label.
 Volume Serial Number is F878-8D74

 Directory of c:\mnt\azure

11/16/2017  09:45 PM    <DIR>          .
11/16/2017  09:45 PM    <DIR>          ..
               0 File(s)              0 bytes
               2 Dir(s)   5,368,709,120 bytes free
```

### Using Azure Disks

1. Create an azure disk storage class

option #1: k8s agent pool is based on blob disk VM

```console
$ kubectl apply -f https://raw.githubusercontent.com/JiangtianLi/Examples/master/windows/azuredisk/storageclass-azuredisk.yaml
```

option #2: k8s agent pool is based on managed disk VM

```console
$ kubectl apply -f https://raw.githubusercontent.com/JiangtianLi/Examples/master/windows/azuredisk/storageclass-azuredisk-managed.yaml
```

2. make sure storageclass is created successfully

```console
$ kubectl get storageclass/azuredisk -o wide
```

3. Create a pvc for azure disk

```console
$ kubectl apply -f https://raw.githubusercontent.com/JiangtianLi/Examples/master/windows/azuredisk/pvc-azuredisk.yaml
```

4. Make sure pvc is created successfully

```console
$ kubectl get pvc/pvc-azuredisk -o wide
```

5. Create a pod with azure disk pvc

```console
$ kubectl apply -f https://raw.githubusercontent.com/JiangtianLi/Examples/master/windows/azuredisk/iis-azuredisk.yaml
```

6. Watch the status of pod until its `STATUS` is `Running`

```console
$ watch kubectl get po/iis-azuredisk -o wide
```

7. Enter the pod container to validate

```console
$ kubectl exec -it iis-azuredisk -- cmd
```

### Multiple containers in a POD

Copy this yaml to a file, then deploy it with `kubectl apply -f <filename.yaml>`

It will run 2 containers:

- iis-container: This is a basic static web server
- servercore-container: This will run a script that changes the index page every 10 seconds

Once deployed, you can load the web page and refresh it to see the contents changing.

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: two-containers
  name: two-containers
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: two-containers
      name: two-containers
    spec:
      volumes:
      - name: shared-data
        emptyDir: {}

      containers:

        - name: iis-container
          image: microsoft/iis:windowsservercore-1803
          volumeMounts:
          - name: shared-data
            mountPath: /wwwcache
          command:
          - powershell.exe
          - -command
          - "while ($true) { Start-Sleep -Seconds 10; Copy-Item -Path C:\\wwwcache\\iisstart.htm -Destination C:\\inetpub\\wwwroot\\iisstart.htm; }"

        - name: servercore-container
          image: microsoft/windowsservercore:1803
          volumeMounts:
          - name: shared-data
            mountPath: /poddata
          command:
          - powershell.exe
          - -command
          - "$i=0; while ($true) { Start-Sleep -Seconds 10; $msg = 'Hello from the servercore container, count is {0}' -f $i; Set-Content -Path C:\\poddata\\iisstart.htm -Value $msg; $i++; }"

      nodeSelector:
        beta.kubernetes.io/os: windows
```

## Troubleshooting

Windows support is still in **active development** with many changes each week. Read on for known per-version issues and for help troubleshooting if you run into problems.

### Finding logs

To connect to a Windows node using Remote Desktop and get logs, please read over this topic in the main [troubleshooting](../howto/troubleshooting.md#connecting-to-windows-nodes) page first.

### Checking versions

Please be sure to include this info with any Windows bug reports:

1. Basic version information:

```console
$ kubectl version
$ kubectl describe node <windows node>
```

Also note any IP Addresses for the next step, but you don't need to share it

1. The Azure CNI plugin version and configuration that is stored in `C:\k\azurecni\netconf\10-azure.conflist`
1. The Azure CNI build by running `C:\k\azurecni\bin\azure-vnet.exe --help`. It will dump some errors, but the version such as `v1.0.4-1-gf0f090e` will be listed.

```text
...
2018/05/23 01:28:57 "Start Flag false CniSucceeded false Name CNI Version v1.0.4-1-gf0f090e ErrorMessage required env variables missing vnet []
...
```

### Known Issues per Version

AKS Engine | Windows Server | Kubernetes | Azure CNI | Notes
-----------|----------------|------------|-----------|----------
V0.16.2 | Windows Server version 1709 (10.0.16299.____) | V1.9.7 | ? | DNS resolution is not configured
V0.17.0 | Windows Server version 1709 | V1.10.2 | v1.0.4 | Acs-engine version 0.17 defaults to Windows Server version 1803. You can override it to use 1709 instead [here](#choosing-the-windows-server-version). Manual workarounds needed on Windows for DNS Server list, DNS search suffix
V0.17.0 | Windows Server version 1803 (10.0.17134.1) | V1.10.2 | v1.0.4 | Manual workarounds needed on Windows for DNS Server list, DNS search suffix, and dropped packets
v0.17.1 | Windows Server version 1709 | v1.10.3 | v1.0.4-1-gf0f090e | Manual workarounds needed on Windows for DNS Server list and DNS search suffix. This AKS Engine version defaults to Windows Server version 1803, but you can override it to use 1709 instead [here](#choosing-the-windows-server-version)
v0.18.3 | Windows Server version 1803 | v1.10.3 | v1.0.6 | Pods cannot resolve cluster DNS names
v0.20.9 | Windows Server version 1803 | v1.10.6 | v1.0.11 | Pods cannot resolve cluster DNS names

### Known problems

#### Packets from Windows pods are dropped

Affects: Windows Server version 1803 (10.0.17134.1)

Issues: <https://github.com/Azure/acs-engine/issues/3037>

There is a problem with the “L2Tunnel” networking mode not forwarding packets correctly specific to Windows Server version 1803. Windows Server version 1709 is not affected.

Workarounds:
**Fixes are still in development.** A Windows hotfix is needed, and will be deployed by AKS Engine once it's ready. The hotfix will be removed later when it's in a future cumulative rollup.

#### Pods cannot resolve public DNS names

Affects: Some builds of Azure CNI

Issues: <https://github.com/Azure/azure-container-networking/issues/147>

Run `ipconfig /all` in a pod, and check that the first DNS server listed is within your cluster IP range (10.x.x.x). If it's not listed, or not the first in the list, then an azure-cni update is needed.

Workaround:

1. Get the kube-dns service IP with `kubectl get svc -n kube-system kube-dns`
2. Cordon & drain the node
3. Modify `C:\k\azurecni\netconf\10-azure.conflist` and make it the first entry under Nameservers
4. Uncordon the node

Example:

```js
{
    "cniVersion":  "0.3.0",
    "name":  "azure",
    "plugins":  [
                    {
                        "type":  "azure-vnet",
                        "mode":  "tunnel",
                        "bridge":  "azure0",
                        "ipam":  {
                                     "type":  "azure-vnet-ipam"
                                 },
                        "dns":  {
                                    "Nameservers":  [
                                                        "10.0.0.10",
                                                        "168.63.129.16"
                                                    ],
                                    "Search":  [
                                                   "default.svc.cluster.local"
                                               ]
                                },
    ...
```

#### Pods cannot resolve cluster DNS names

Affects: Azure CNI plugin <= 0.3.0

Issues: <https://github.com/Azure/azure-container-networking/issues/146>

If you can't resolve internal service names within the same namespace, run `ipconfig /all` in a pod, and check that the DNS Suffix Search List matches the form `<namespace>.svc.cluster.local`. An Azure CNI update is needed to set the right DNS suffix.

Workaround:

1. Use the FQDN in DNS lookups such as `kubernetes.kube-system.svc.cluster.local`
1. Instead of DNS, use environment variables `* _SERVICE_HOST` and `*_SERVICE_PORT` to find service IPs and ports in the same namespace

#### Pods cannot ping default route or internet IPs

Affects: All clusters created by AKS Engine

ICMP traffic is not routed between private Azure vNETs or to the internet.

Workaround: test network connections with another protocol (TCP/UDP). For example `Invoke-WebRequest -UseBasicParsing https://www.azure.com` or `curl https://www.azure.com`.

## Cluster Troubleshooting

If your cluster is not reachable, you can run the following command to check for common failures.

### Misconfigured Service Principal

If your Service Principal is misconfigured, none of the Kubernetes components will come up in a healthy manner.
You can check to see if this the problem:

```shell
ssh -i ~/.ssh/id_rsa USER@MASTERFQDN sudo journalctl -u kubelet | grep --text autorest
```

If you see output that looks like the following, then you have **not** configured the Service Principal correctly.
You may need to check to ensure the credentials were provided accurately, and that the configured Service Principal has
read and **write** permissions to the target Subscription.

`Nov 10 16:35:22 k8s-master-43D6F832-0 docker[3177]: E1110 16:35:22.840688    3201 kubelet_node_status.go:69] Unable to construct api.Node object for kubelet: failed to get external ID from cloud provider: autorest#WithErrorUnlessStatusCode: POST https://login.microsoftonline.com/72f988bf-86f1-41af-91ab-2d7cd011db47/oauth2/token?api-version=1.0 failed with 400 Bad Request: StatusCode=400`
