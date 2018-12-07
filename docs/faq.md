# AKS-Engine FAQ

This page provides help with the most common questions about AKS-Engine.

### What's the Difference Between AKS and AKS-Engine?

Azure Kubernetes Service ([AKS][]) is a Microsoft Azure service that supports fully managed Kubernetes clusters. [AKS-Engine][] is an Azure open source project that creates Kubernetes clusters with your custom requirements. AKS uses AKS-Engine internally, but they are not the same.

AKS clusters can be created in the Azure portal or with `az aks create` in the [Azure command-line tool][]. AKS-Engine clusters can be created with `aks-engine deploy` in the AKS-Engine command-line tool, or by generating the ARM templates with `aks-engine generate` and deploying them as a separate step.

### What's the Difference Between `acs-engine` and `aks-engine`?

AKS-Engine is the next version of the ACS-Engine project. AKS-Engine supports current and future versions of [Kubernetes][], while ACS-Engine also supported the Docker Swarm and Mesos DC/OS container orchestrators.

### Is ACS-Engine Still Active?

Security fixes and other critical changes will still go into ACS-Engine and patch releases will be made as needed. But in general, improvements and releases with new features will only be made in AKS-Engine.

### Can I Build an AKS Cluster with `aks-engine`?

No, Azure Kubernetes Service itself is the way to create a supported, managed AKS cluster. AKS-Engine shares some code with AKS, but does not create managed clusters.

### Where is the `aks-engine` Release?

We don't have an official aks-engine release yet. Please use the most recent [acs-engine release][] for now.

[AKS]: https://azure.microsoft.com/en-us/services/kubernetes-service/
[AKS-Engine]: https://github.com/Azure/aks-engine
[Azure command-line tool]: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest
[acs-engine release]: https://github.com/Azure/acs-engine/releases
[Kubernetes]: https://kubernetes.io/
