# Windows Provisioning Scripts

## Overview

Due to requirements around signing executable code on Windows nodes some powershell scripts are delivered via a .zip file during node provisioning.
This document details how to update and test changes to these files.

## Testing

Changes to the provisioning scripts packaged in the .zip can be tested by:

- Running [scripts/build-windows-provisioning-scripts.sh](../../scripts/build-windows-provisioning-scripts.sh) which will create a .zip and upload it to a storage account
- Create/update a cluster definition file and set `windowsProfile.provisioningScriptsPackageURL' to point to the uploaded back.
- Perform a cluster deployment with aks-engine.

## Updates

- Check in script updates to [staging/provisioning/windows](../../staging/provisioning/windows)
- Create a new **versioned** zip with the contents of [staging/provisioning/windows](../../staging/provisioning/windows) and upload it to a storage account
  - Optional - First sign the powershell scripts (out-of-scope for this documentation)
- Update `DefaultWindowsProvisioningScriptsPackageURL` in [pkg/api/const.go](../../pkg/api/const.go) with new URL
- Add new URL to [/vhd/packer/configure-windows-vhd.ps1::Get-FilesToCacheOnVHD](../../vhd/packer/configure-windows-vhd.ps1)

## Validation

Before creating a new aks-engine release

- Run [scripts/validate-windows-provisioning-scripts.sh](../../scripts/validate-windows-provisioning-scripts.sh) to verify the contents in the staging directory match the the contents of the zip used in default cluster deployments.
- Perform a no-outbound-connections deployment to verify the zip used in default cluster deployments is cached on the Windows VHD.
