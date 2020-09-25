# AKS Engine FAQ

This page provides help with the most common questions about AKS Engine.

### What's the Difference Between AKS and AKS Engine?

Azure Kubernetes Service ([AKS][]) is a Microsoft Azure service that supports fully managed Kubernetes clusters. [AKS Engine][] is an Azure open source project that allows you to create your own Kubernetes clusters with lots of user-configurable options, but with no official Azure support.

AKS clusters can be created in the Azure portal or with `az aks create` in the [Azure command-line tool][]. AKS Engine clusters can be created with `aks-engine deploy` (`aks-engine` is the AKS Engine command-line tool), or by generating ARM templates with `aks-engine generate` and deploying them as a separate step using the `az` command-line tool (e.g., `az group deployement create`).

### Can I Build an AKS Cluster with `aks-engine`?

No, the `aks-engine` command line tool will not create an AKS cluster. To learn about AKS, you can [read the official docs](https://docs.microsoft.com/en-us/azure/aks/).

### What's the Difference Between `acs-engine` and `aks-engine`?

AKS Engine is the next version of the ACS-Engine project. AKS Engine supports current and future versions of [Kubernetes][], while ACS-Engine also supported the Docker Swarm and Mesos DC/OS container orchestrators.

### Can I Scale or Upgrade an `acs-engine`-created Kubernetes Cluster with `aks-engine`?

Yes.

### Is ACS-Engine Still Active?

No further development or releases in ACS-Engine are planned. AKS Engine is a backward-compatible continuation of ACS-Engine, so all fixes and new features will target AKS Engine.

### Should I use the latest `aks-engine` release if I was previously using `acs-engine`?

Yes. `aks-engine` released [v0.27.0][] as a continuation of the ACS-Engine project ([v0.26.2][] was the final `acs-engine` release) with all the Kubernetes fixes and features included in [v0.26.2][] and more.


[AKS]: https://azure.microsoft.com/en-us/services/kubernetes-service/
[AKS Engine]: https://github.com/Azure/aks-engine
[Azure command-line tool]: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest
[Kubernetes]: https://kubernetes.io/
[v0.27.0]: https://github.com/Azure/aks-engine/releases/tag/v0.27.0
[v0.26.2]: https://github.com/Azure/acs-engine/releases/tag/v0.26.2
