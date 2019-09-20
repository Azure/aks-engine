# New K8s Release

This document outlines the steps needed to ingest new Kubernetes releases into AKS Engine.

## Version Support

In general the following K8s versions are supported in AKS Engine
- The three most recent released Minor versions
- The two most recently Patch versions for supported Minor versions
- All Prerelease versions for the current in-progress release 

## Required AKS Engine Repo Updates

In [`pkg\api\common\versions.go`](../../pkg/api/common/versions.go) 
- Add new version to to AllKubernetesSupportedVersions map.
- Update the is-install-allowed bool following the above [Version Support](#version-Support) guidance to drop install support for older versions.

In [docs\topics\azure-stack.md](../docs/topics/azure-stack.md)
- Update the `Supported Kubernetes Versions` section to match currently supported versions.

In [packer\install-dependencies.sh](../../packer/install-dependencies.sh)
- Update entries in K8S_VERSIONS to reflect currently supported versions (also include Azure Stack /-azs entries).

In [packer\configure-windows-vhd.ps1](../../packer/configure-windows-vhd)
- Update entries in  Get-FilesToCacheOnVHD to reflect currently supported versions (also include Azure Stack / azd- entries).

Update any template files under [examples](../../examples/) or [test\e2e\test_cluster_configs](../../test/e2e/test_cluster_configs) that reference a now unsupported K8s version to a supported one.
