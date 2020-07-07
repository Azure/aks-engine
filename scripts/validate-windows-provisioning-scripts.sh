#!/bin/bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.
set -euo pipefail

function validate_package() {
    local package_url=$1

    temp_dir=$(mktemp -d -t aks-engine-XXXXXXXX)
    echo "Downloading $package_version to $temp_dir/scripts.zip"
    curl -L "$package_url" -o "$temp_dir/scripts.zip"

    echo "Extracting files to $temp_dir/scripts"
    unzip "$temp_dir/scripts.zip" -d "$temp_dir/scripts"

    # Perform a recursive diff of the two directories but ignore comments since checked in powershell files
    # will not container signature blocks at the end of the files.
    diff -bBr --ignore-matching-lines='^#' "$temp_dir/scripts" "./staging/provisioning/windows/"

    echo "Files match"
}

echo "Validating content of Windows provisioning scripts..."
package_version=$(grep  "DefaultWindowsProvisioningScriptsPackageVersion" ./pkg/api/const.go  | cut -d ' ' -f 3 | cut -d '"' -f2)

echo "Package version: $package_version"
if [ -z "$package_version" ]; then
    echo "Could not find value for 'DefaultWindowsProvisioningScriptsPackageVersion' in /pkg/api/const.go"
    exit 1
fi

# urls from pkg/api/azenvtypes.go
validate_package "https://kubernetesartifacts.azureedge.net/aks-engine/windows/provisioning/signedscripts-${package_version}.zip"
validate_package "https://mirror.azk8s.cn/aks-engine/windows/provisioning/signedscripts-${package_version}.zip"

