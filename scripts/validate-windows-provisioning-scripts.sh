#!/bin/bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.
set -euo pipefail

echo "Validating content of Windows provisioning scripts..."
package_url=$(grep  "DefaultWindowsProvisioningScriptsPackageURL" ./pkg/api/const.go  | cut -d ' ' -f 3 | cut -d '"' -f2)

if [ -z "$package_url" ]; then
    echo "Could not find value for 'DefaultWindowsProvisioningScriptsPackageURL' in /pkg/api/const.go"
    exit 1
fi

temp_dir=$(mktemp -d -t aks-engine-XXXXXXXX)

echo "Downloading $package_url to $temp_dir/scripts.zip"
curl -L "$package_url" -o "$temp_dir/scripts.zip"

echo "Extracting files to $temp_dir/scripts"
unzip "$temp_dir/scripts.zip" -d "$temp_dir/scripts"

# Perform a recursive diff of the two directories but ignore comments since checked in powershell files
# will not container signature blocks at the end of the files.
diff -bBr --ignore-matching-lines='^#' "$temp_dir/scripts" "./staging/provisioning/windows/"

echo "Files match"
