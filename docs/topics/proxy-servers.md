# AKS-Engine and proxy servers

Using AKS-engine on [Azure Stack](azure-stack.md) in connected and disconnected environments often times requires the use of a non-transparent proxy server. Non-transparent means that they're not part of the default route and have to be configured to be used.

This document guides you through the different components and layers where you need to configure the proxy server.

- [AKS-engine](#aks-engine)
- [Cluster Nodes](#cluster-nodes)
- [Pods](#pods)

## AKS-engine

AKS-engine means in this context the VM (windows or linux) that is used to run AKS-engine to deploy, scale and upgrade your Kubernetes cluster.

**Linux**

You've to make sure that you export the proxy server configuration via environment variables. You can either set it in your current session, permanantly for a specific user in `~/.bashrc` or permanently for the whole system in `/etc/profile`.

```bash
export HTTP_PROXY=http://proxy:8888
export HTTPS_PROXY=http://proxy:8888
```

***Windows***

```cmd
netsh winhttp set proxy <proxy>:<port>
```

This is required to use tools like `wget`, `curl` etc.

## Cluster Nodes

In your cluster you've to configure a proxy server on both, your worker nodes as well as your master nodes.

You can use the same manual configuration for the proxy servers as described in the [AKS-engine](#aks-engine) section above. But that's a very static way to achieve that. A better way to dynamically configure your cluster nodes is using a Kubernetes [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/).

**Linux**

```bash
export HTTP_PROXY=http://proxy:8888
export HTTPS_PROXY=http://proxy:8888
```

Besides the system itself you might need to configure the proxy server for apt (aptitude) directly:

* [Setting up apt-get to use a http-proxy](https://help.ubuntu.com/community/AptGet/Howto#Setting_up_apt-get_to_use_a_http-proxy)

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
- [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/)
- [Setting up apt-get to use a http-proxy](https://help.ubuntu.com/community/AptGet/Howto#Setting_up_apt-get_to_use_a_http-proxy) (Ubuntu Linux)