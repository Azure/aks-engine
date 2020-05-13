# AKS-Engine and proxy servers

Using AKS-engine on [Azure Stack](azure-stack.md) in connected and disconnected environments often times requires the use of a non-transparent proxy server. Non-transparent means that they're not part of the default route and have to be configured to be used.

> Azure Stack Hub itself supports only _transparent_ proxy server setups today. This isn't relevant for the workload, as long as the workload (and it's components) can access the proxy server.

This document guides you through the different components and layers where you need to configure the proxy server. Considerations include:

- [AKS-engine](#aks-engine)
- [Cluster Nodes](#cluster-nodes)
- [Pods](#pods)

What kind of egress traffic you can expect is documented here: [Control egress traffic for cluster nodes in Azure Kubernetes Service (AKS)](https://docs.microsoft.com/en-us/azure/aks/limit-egress-traffic). Even though the document is referring to AKS (in Azure) it gives you an idea what kind of traffic and requests you can expect.

## AKS-engine

AKS-engine means in this context the VM ([Windows](https://docs.microsoft.com/azure-stack/user/azure-stack-kubernetes-aks-engine-deploy-windows) or [Linux](https://docs.microsoft.com/azure-stack/user/azure-stack-kubernetes-aks-engine-deploy-linux)) that is used to run AKS-engine to deploy, scale and upgrade your Kubernetes cluster.

**Linux**

On a Linux system you've to make sure that you export the proxy server configuration via environment variables. Most tools will leverage these environment variables and will automatically use them w/o additional configuration.

You can either set these environment variables in your current session, permanantly for a specific user in `~/.bashrc` or permanently for the whole system in `/etc/profile` (or in `/etc/environment`).

```bash
export HTTP_PROXY=http://proxy:8888
export HTTPS_PROXY=http://proxy:8888
```

In case your proxy servers require authentication:

```bash
export HTTP_PROXY="http://usrname:passwrd@host:port"
export HTTPS_PROXY="http://usrname:passwrd@host:port"
```

And in case you want to exclude specific URLs and IP-addresses that do not not need a proxy you can use:

```bash
export NO_PROXY=master.hostname.example.com,<docker_registry_ip>,docker-registry.default.svc.cluster.local
```

Setting these environment variables is required to use tools like `wget`, `curl` etc. and force them to use a proxy server.

To set the proxy server configuration permanently on a system it's recommended to write the configuration into a separate file in `/etc/profile.d`:

```bash
echo "export http_proxy=http://host:port/" > /etc/profile.d/http_proxy.sh
```

This will make the configuration persistent and will also survi

***Windows***

On Windows-based systems you have to use `netsh` to configure the proxy server:

```cmd
netsh winhttp set proxy <proxy>:<port>
```

## Cluster Nodes

In your cluster you've to configure a proxy server on both, your worker nodes as well as your master nodes. This is required for example to use `apt` and also for `docker pull` to download container images.

You can use the same manual configuration for the proxy servers as described in the [AKS-engine](#aks-engine) section above. But that's a very static way to achieve that. A better way to dynamically configure your cluster nodes is using a Kubernetes [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) this will also take care of newly added nodes while scaling out your cluster.

**DaemonSet**  
Here's an example how you can set the proxy server configuration to all nodes in your cluster via a Kubernetes DaemonSet:

```YAML
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: proxy-configuration
spec:
  selector:
    matchLabels:
      name: proxy-configuration
  template:
    metadata:
      labels:
        name: proxy-configuration
    spec:
      volumes:
        - name: hostfs
          hostPath:
            path: /
      initContainers:
        - name: init
          image: alpine
          command:
            - /bin/sh
            - -xc
            - |
              # Write Proxy Config to /etc/profile.d/
              echo "export http_proxy=http://host:port/" > /etc/profile.d/http_proxy.sh
          volumeMounts:
            - name: hostfs
              mountPath: /host
      containers:
        - name: sleep
          image: alpine
          command:
            - /bin/sh
            - -xc
            - |
              while ($true) do { sleep 60; } done;
```

**Linux**

```bash
export HTTP_PROXY=http://proxy:8888
export HTTPS_PROXY=http://proxy:8888
```

Besides the system itself you might need to configure the proxy server for apt (aptitude) directly:

- [Setting up apt-get to use a http-proxy](https://help.ubuntu.com/community/AptGet/Howto#Setting_up_apt-get_to_use_a_http-proxy)

**Windows**

```cmd
netsh winhttp set proxy <proxy>:<port>
```

## Pods

Even your Pods/Container Instances need in some cases to go through a proxy server to access services outside the cluster or outside your network. You can inject the proxy configuration directly within your YAML manifest:

```YAML
containers:
- env:
  - name: "HTTP_PROXY"
    value: "http://USER:PASSWORD@IPADDR:PORT"
```

## External Resources

- [Configure machine proxy and Internet connectivity settings](https://docs.microsoft.com/en-us/windows/security/threat-protection/microsoft-defender-atp/configure-proxy-internet) (Windows)
- [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/) (Docker)
- [Setting up apt-get to use a http-proxy](https://help.ubuntu.com/community/AptGet/Howto#Setting_up_apt-get_to_use_a_http-proxy) (Ubuntu Linux)
- [How to set up proxy using http_proxy & https_proxy environment variable in Linux?](https://www.golinuxcloud.com/set-up-proxy-http-proxy-environment-variable/) (Linux)
