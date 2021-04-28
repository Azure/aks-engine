# Windows Provisioning Scripts

## Overview

Microsoft requires that all executable code (including PowerShell) used by first party services running on Windows machines be signed by Microsoft.
The build process used to produce aks-engine releases does not have the ability to sign code on behalf of Microsoft so powershell code used to support running Kubernetes on Windows nodes are signed elsewhere and are delivered via a .zip file during node provisioning.

This document details how to update and test changes to these files.

## Testing

### Local Validation

Changes to the provisioning scripts packaged in the .zip can be tested by:

- Running [scripts/build-windows-provisioning-scripts.sh](../../scripts/build-windows-provisioning-scripts.sh) which will create a .zip and upload it to a storage account.
- Create/update a cluster definition file and set `windowsProfile.provisioningScriptsPackageURL' to point to the uploaded back.
- Perform a cluster deployment with aks-engine.

### CI Pipeline validation

Changes to the provisioning scripts are also tested with an Azure DevOps pipeline job enabled in Azure/aks-engine.
This pipeline triggers when files under [staging/provisioning/windows](../../staging/provisioning/windows) are detected in a PR.
Pipeline definition file: [pr-windows-signed-scripts.yaml](../../.pipelines/pr-windows-signed-scripts.yaml)

Note: By default this pipeline will use [examples/e2e-tests/kubernetes/windows/hybrid/definition.json](../../examples/e2e-tests/kubernetes/windows/hybrid/definition.json). This can be set as a schedule-time variable by aks-engine maintainers to validate other cluster configurations.

## Updates

- Check in script updates to [staging/provisioning/windows](../../staging/provisioning/windows).
- Create a new **versioned** zip with the contents of [staging/provisioning/windows](../../staging/provisioning/windows) and upload it to a storage account. (Versioning ensues compatability with file caching mecanisms used by aks-engine VHD backed installs)
  - Optional - First sign the powershell scripts (out-of-scope for this documentation).
- Update `DefaultWindowsProvisioningScriptsPackageVersion` in [pkg/api/const.go](../../pkg/api/const.go) with new URL.
- Add new URL to [/vhd/packer/configure-windows-vhd.ps1::Get-FilesToCacheOnVHD](../../vhd/packer/configure-windows-vhd.ps1).

## Validation

Before creating a new aks-engine release

- Run [scripts/validate-windows-provisioning-scripts.sh](../../scripts/validate-windows-provisioning-scripts.sh) to verify the contents in the staging directory match the the contents of the zip used in default cluster deployments.
- Perform a no-outbound-connections deployment to verify the zip used in default cluster deployments is cached on the Windows VHD.
