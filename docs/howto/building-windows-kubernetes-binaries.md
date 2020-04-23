# Building Windows Kubernetes Binaries

## Background

Microsoft maintains a fork of the Kubernetes project at https://github.com/Azure/kubernetes which includes patches not yet included in upstream Kubernetes for release 1.7 and 1.8; these are needed for Windows containers to function. *From release 1.9, all Windows features are in upstream and Windows binaries no longer needs to be built from Azure fork.*

## Azure Pipelines Does the Work

When a new release of Kubernetes is finalized, Microsoft builds it and makes its artifacts available in Azure storage under `https://kubernetesartifacts.azureedge.net/kubernetes/`.

There are no build tasks for AKS Engine maintainers to perform. Just ensure Microsoft's build pipeline finishes, then validate that the Windows Kubernetes binaries were packaged for AKS Engine. The Windows archive will be named like this:

`https://kubernetesartifacts.azureedge.net/kubernetes/v1.17.2/windowszip/v1.17.2-1int.zip`

Don't forget to update the Windows VHD with this archive name in the pull request for the new Kubernetes release.
