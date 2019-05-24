# Building Azure marketplace images for Windows AKS support

## Motivation

Initially, AKS-Engine's codebase on Linux installed all of the needed stuff from package repos or container registries. This provided the best flexibility as any packages could be easily updated and tested with a simple generate+deploy. However, it had a few major downsides. Network or repo downtime would cause failed deployments that were not recoverable. It also took additional time to do the downloads, making the overall deployments slower. To mitigate this while preserving deployment flexibility, AKS-engine switched to using custom VM images [built with Packer](../../packer). 

Windows is in a halfway state today were the Windows Server team publishes `Windows Server version ___ with Containers` that has Docker preinstalled and some container images `nanoserver` and `windowsservercore` pre-pulled. This helps save some deployment time, but means that making changes to it depends on another team, and not having automation makes it much more difficult to pre-test changes to what's installed in these images. We need the flexibility to validate new OS versions, patches, and other dependencies. Moving to an image-based solution would align with the Linux workflow and allow us to control and pre-test all changes. It also leaves the option to make other customizations specific to AKS that will help with reliability or deployment time.

The other challenge with the Windows deployment workflow today is that all of the Kubernetes node binaries, CNI plugin, and default configuration files are in a single ZIP published to the acsmirror CDN. It's built + packaged in [build-windows-k8s.sh](../../scripts/build-windows-k8s.sh). This was written when the upstream kubernetes/kubernetes code was in an alpha state, downstream patches were required, and public builds were not yet available. This is no longer needed and makes it harder to revise and test those components individually.

All of this needs to be cleaned up so that we can get avoid deployment time penalties due to downloads and keep all changes in dependencies in source code so that they can be thoroughly tested and reproduced. 

## Use Cases

In general, most code build for AKS-Engine deployments should go through a VHD-based workflow to match what will be used in production and avoid having extra implementations for the same scenario. Ones that are intentionally overlapping should be called out.

### Testing new Windows versions - VHD

Testing new Windows versions is particularly difficult for a few reasons:

- Windows Server Insider builds are published frequently, but not all are available on Azure. Starting from an ISO+Packer works around that limitation.
- The base containers are different -  `nanoserver-insider` vs `nanoserver`, and are incompatible with the released versions. This means that images Kubernetes depends on such as `pause` need to be built and used as well. You can't build these on an older OS, but they can be done within the VM that's running as part of the VHD build process.
  - As an aside, matching test containers are required too. If these aren't built from Kubernetes official builds, they could be built and pushed to a private registry using the same VM. Leaving them on the VHD will actually improve test time by avoiding pulls later.

With the VHD-based workflow, adding (or removing) support for a new version could be reduced to a few parameters: `WindowsIsoUrl`, `NanoServerImage`. If these are in a pull request, automation can run the full vhd build followed by AKS-Engine E2E test pass.

If a private build needed to be tested, `WindowsIsoUrl` could be a path requiring an Azure login, and the resulting image could be kept in the private store for that subscription.


### Testing new Windows patches - VHD

This is similar, but simpler than the new Windows version case above. It could use a new `WindowsIsoUrl`, `WindowsPatchList`, or `WindowsPrivatePatchUrl`. There are no changes needed to the container images. 

Teams needing to create and test prerelease patches will need to use a private branch and Azure DevOps Pipeline to avoid unintentional security disclosures. This should be kept in-sync with the public repo but use a different Azure subscription and accounts.

### AKS-Users testing new Windows patches

During the VHD build, `WindowsPatchList` and `WindowsPrivatePatchList` control what cumulative updates and/or patches are installed. However, there may be cases where a private patch has been given to a customer who cannot build a new VHD. The existing [windows-patches](C:\Users\patri\Source\aks-engine\extensions\windows-patches) extension should still work to do this.


### Testing new runtime component versions in AKS-Engine

Testing dependencies such as Docker, ContainerD, CNI plugins, ... that are released through AKS-Engine builds should follow the same process and cadence as the Linux VHD workflow.

### Testing upstream builds of Kubernetes

The main [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes) codebase is tested on Azure using AKS-Engine. The test results are not triggered by or reported to the AKS-Engine repo. This is done by setting the `kubernetesConfig.WindowsNodeBinariesURL` in the API model. That will be downloaded to each node and overwrite what was installed in the VHD.

### Testing upstream builds of CRI-ContainerD

Similar to kubernetes/kubernetes, containerd/containerd needs a working deployment on Azure to test new builds. That CI system can produce a ContainerD build, then set `kubernetesConfig.WindowsContainerdURL` in the API model to get a working deployment.

### Testing upstream builds of CNI plugins

Similar to kubernetes/kubernetes, [Microsoft/SDN](https://github.com/Microsoft/SDN) needs a working deployment on Azure to test new builds. That CI system can produce a ContainerD build, then set `kubernetesConfig.WindowsSdnPluginURL` in the API model to get a working deployment.

## Workflow

### VHD Build

These configurations drive the VHD build process. As such, they should be version controlled in source code for reproducability. They will be set as variables for a Packer build. Some of these have intentional overlap between the VHD build steps and node deployment to fulfill different use cases.

1. (optional, only if installing from ISO) Run Packer with an unattend file to get a working Windows setup from an ISO image
  * Install optional features: `Containers`, `Hyper-V`
  * Enable remote access with WinRM & Remote Desktop
  * Install the [Azure Agent](https://docs.microsoft.com/en-us/azure/virtual-machines/extensions/agent-windows#manual-installation)
  * Reboot
1. Build a cache of all needed dependencies - Docker EE, Moby, ContainerD, Azure-CNI, SdnBridge, and Kubernetes
1. Install the default version of Docker EE-basic, Moby, or ContainerD
1. Pull container images needed such as core-dns, and also list of containers in `ContainerTags`
1. [Sysprep](https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/sysprep--generalize--a-windows-installation) to remove identifiers and default users from the VM, then shutdown.

#### Parameters for all Windows versions

These parameters are used to control everything that's installed on the VHD, regardless of Windows version or installation method (ISO vs Azure Marketplace as base image)

Parameter  | Description
-----------|-------------
`ContainerdVersions` | Future: List of [ContainerD versions](https://github.com/containerd/containerd/releases) to cache
`MobyVersions` | Future: Use for Moby versions to cache, pulled from packages.microsoft.com to match Linux versions
`DockerVersions` | List of Docker EE-basic versions to cache, pulled with [DockerMsftProvider](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick-start/quick-start-windows-server#install-docker)
`SdnPluginVersions` | Future: List of versions of `sdnbridge/sdnoverlay` to cache. These aren't released yet
`AzureVnetVersions` | List of [Azure-CNI](https://github.com/Azure/azure-container-networking/releases) versions to cache
`ContainerTags` | List of containers to pull in canonical `registry.dns/name:tag` format.
`CoreDnsVersions` | List of CoreDNS versions to pull. Use same registry as Linux - see [install-dependencies](../../packer/install-dependencies.sh)
`KubeNodeVersions` | List of [Kubernetes node binary versions](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.14.md#node-binaries) to cache
`ContainerRuntime` | Pulled and built images are not shared between containerd and Moby/Docker. This sets which one to use.
`DefaultDockerVersion` | Version of Docker EE-basic to use on Windows if `ContainerRuntime == Docker`
`DefaultContainerdVersion` | Version of containerd to install if applicable
`DefaultMobyVersion` | Version of Moby to install if applicable
`WindowsPatchList` | List of public patches to install using Windows Update, by [KB number](https://support.microsoft.com/en-us/help/4464619)
`WindowsPrivatePatchList` | After public patches are applied, this list of private patches will be downloaded and applied. This will turn test signing on.


#### Parameters for shipped OS versions

These parameters can only be used when starting from an existing Windows OS image available from the Azure marketplace. It should not have any container runtime installed or it may cause conflicts.

Parameter  | Description
-----------|-------------
`WindowsBaseImage` | Existing Azure image to boot up and customize. Mutually exclusive with `WindowsIsoUrl`. Example: `MicrosoftWindowsServer:WindowsServerSemiAnnual:Datacenter-Core-1809-smalldisk:1809.0.20181107`

#### Parameters for prerelease OS versions

These parameters cannot be used when `WindowsBaseImage` is set. This requires running Packer to boot a new VM, and run an unattended Windows setup. Once the unattended setup is complete, the remaining steps will be the same.

Parameter  | Description
-----------|-------------
`WindowsIsoUrl` | ISO image to build from. Can be an Azure storage blob requiring login. Mutually exclusive with `WindowsIsoUrl`
`NanoServerName` | Full name `mcr.microsoft.com/.../nanoserver-insider`

ISO builds will need a few more files in this repo for Windows VHD creation:

- `unattend.xml` - needed to configure default user, drive layout, and trial license keys
- `windowsserverinsider.json` - Packer configuration for the VM build.
- Some additional scripts may be called from the packer JSON if needed to install the Azure agent and other necessary steps.

### Deployment Phase

Once a VM is deployed, the right configuration should be applied. This is handled by a custom script included in the ARM template for that VM.

It will:

1. Finish installing the desired Kubernetes, Moby/Docker/ContainerD, and network plugin versions. The user will provide these through the AKS-Engine [Cluster Definition](docs\topics\clusterdefinitions.md)
2. Create or update configuration files for the above. Many of these have interdependencies based on file paths, binary names, or other configurations such as kubelet parameters that aren't known until deployment time.
3. Start the necessary services
4. Join the Kubernetes cluster

## AKS-Engine Apimodel Changes

These parameters are defined in [Cluster Definition](docs\topics\clusterdefinitions.md), and will have impact on what needs to be put in the Windows images and the configuration steps needed as they're deployed. Some of these are not used as-is, but instead are used to compute other values such as a path to a tarball on the acsmirror CDN. These ones should be checked for cached downloads before trying the acsmirror CDN since they may be in the VHD already.

Apimodel   |  Description
-----------|--------------
`windowsProfile.windowsPublisher` | Old default: `MicrosoftWindowsServer`. After this change, only `Azure` will be supported.
`windowsProfile.windowsOffer` | Old default: `WindowsServerSemiAnnual`. New value `WindowsServerForAKS`
`windowsProfile.windowsSku` | Old default: `Datacenter-Core-1809-with-Containers-smalldisk`. New values `Datacenter-Core-2019` (default), additional values: `Datacenter-Core-19H1`. **1809 will be ignored** as there is no reason to use it instead of 2019. All images will be 30Gb and expanded at deployment time if needed.
`windowsProfile.imageVersion` | Default: `latest`. Specific values will follow format `YYYYMMDD.commithash`
`orchestratorProfile.orchestratorRelease` + `orchestratorProfile.orchestratorVersion` | These together make up the Kubernetes version. Multiple releases+versions need to be cached on the node.
`kubernetesConfig.containerRuntime` | Windows will need to support `docker` and `containerd` and download/cache each.
`WindowsProfile.WindowsDockerVersion` | Version of Docker EE-basic to use on Windows
`kubernetesConfig.mobyVersion` | For future use. Linux nodes use [moby](https://github.com/moby/moby) builds from `packages.microsoft.com`. If Windows builds are published too, they should be supported. As of May 2019, this isn't needed. Only Docker EE-basic and ContainerD are implemented.
`kubernetesConfig.containerdVersion` | For future use. Once stable containerd builds are available matching the Linux supported versions then this will be used. Until then, use `WindowsContainerdURL` to point to a dev build.
`kubernetesConfig.networkPlugin` | `kubenet` will map to `sdnbridge` on Windows. `Azure` is supported. `flannel` may be added in the future.
`kubernetesConfig.networkPolicy` | `Azure` is supported, requires networkPlugin = Azure. In the future, `calico` could be supported if public builds are available.


These properties override defaults computed by AKS-Engine to aid in testing. These are not built-in to the VHD.

Apimodel   |  Description
-----------|--------------
`kubernetesConfig.WindowsNodeBinariesURL` | 
`kubernetesConfig.WindowsContainerdURL` |
`kubernetesConfig.WindowsSdnPluginURL` | 
`kubernetesConfig.azureCNIURLWindows` | 
`windowsProfile.windowsImageSourceURL` | 

These values will be removed with this change:

Apimodel   |  Description
-----------|--------------
`kubernetesConfig.customWindowsPackageURL` | This will be deprecated. It's an old package that combines multiple things into one zip generated by scripts/build-windows-k8s.sh. It makes versioning harder.


