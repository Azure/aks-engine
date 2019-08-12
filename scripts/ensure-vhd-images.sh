#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

PUBLISHED_IMAGES=100

exit_code=0

echo "==> Using az CLI to validate published VHD images <=="

az login --service-principal --username "${CLIENT_ID}" --password "${CLIENT_SECRET}" --tenant "${TENANT_ID}" &>/dev/null
IMAGES=$(az vm image list -p AKS -f aks --all -l eastus | jq '. | length')

if [ "$IMAGES" != "$PUBLISHED_IMAGES" ]; then
  echo "Expected to find $PUBLISHED_IMAGES published VHD images, instead found $IMAGES"
  exit_code=1
fi

exit $exit_code
