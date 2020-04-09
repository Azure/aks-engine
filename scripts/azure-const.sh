#!/bin/bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

if [ -z "$CLIENT_ID" ]; then
    echo "must provide a CLIENT_ID env var"
    exit 1;
fi

if [ -z "$CLIENT_SECRET" ]; then
    echo "must provide a CLIENT_SECRET env var"
    exit 1;
fi

if [ -z "$TENANT_ID" ]; then
    echo "must provide a TENANT_ID env var"
    exit 1;
fi

if [ -z "$SUBSCRIPTION_ID" ]; then
    echo "must provide a SUBSCRIPTION_ID env var"
    exit 1;
fi

az login --service-principal \
		--username "${CLIENT_ID}" \
		--password "${CLIENT_SECRET}" \
		--tenant "${TENANT_ID}" &>/dev/null

# set to the sub id we want to cleanup
az account set -s "$SUBSCRIPTION_ID"

./bin/aks-engine get-locations -o code --client-id="${CLIENT_ID}" --client-secret="${CLIENT_SECRET}" --subscription-id="${SUBSCRIPTION_ID}" \
  > pkg/helpers/azure_locations.go
./bin/aks-engine get-skus -o code --client-id="${CLIENT_ID}" --client-secret="${CLIENT_SECRET}" --subscription-id="${SUBSCRIPTION_ID}" \
  > pkg/helpers/azure_skus_const.go
git status | grep pkg/helpers/azure
exit_code=$?
if [ $exit_code -gt "0" ]; then
  echo "No modifications found! Exiting 0"
  exit 0
else
  echo "File was modified! Exiting 1"
  exit 1
fi
