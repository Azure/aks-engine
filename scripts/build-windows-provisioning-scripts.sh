#!/bin/bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.
set -o errexit
set -o nounset
set -o pipefail

required_env_vars=(
    "CLIENT_ID"
    "CLIENT_SECRET"
    "LOCATION"
    "RESOURCE_GROUP_NAME"
    "STORAGE_ACCOUNT_NAME"
    "STORAGE_CONTAINER_NAME"
    "SUBSCRIPTION_ID"
    "TENANT_ID"
)

unset_vars=()
for v in "${required_env_vars[@]}"
do
    if [ -z "${!v}" ]; then
        unset_vars+=(${v})
    fi
done

if (( ${#unset_vars[@]} )); then
    echo "Required environment variables: ${required_env_vars[@]}"
    echo "Unset environment variables: ${unset_vars[@]}"
    exit 1
fi

CREATE_TIME="$(date +%s)"
az login --service-principal -u $CLIENT_ID -p $CLIENT_SECRET --tenant $TENANT_ID

rg_exists=$(az group exists -g $RESOURCE_GROUP_NAME --subscription $SUBSCRIPTION_ID -o json)
if $rg_exists ; then
    echo "resource group '${RESOURCE_GROUP_NAME}' exists"
else
    echo "creating new resource group: ${RESOURCE_GROUP_NAME}"
    az group create -l $LOCATION -g $RESOURCE_GROUP_NAME --subscription $SUBSCRIPTION_ID
fi

avail=$(az storage account check-name -n ${STORAGE_ACCOUNT_NAME} -o json | jq -r .nameAvailable)
if $avail ; then
    echo "creating new storage account: ${STORAGE_ACCOUNT_NAME}..."
    az storage account create -n $STORAGE_ACCOUNT_NAME -g $RESOURCE_GROUP_NAME --sku "Standard_RAGRS" --tags "now=${CREATE_TIME}"
else
    echo "storage account '${STORAGE_ACCOUNT_NAME}' already exists."
fi

exists=$(az storage container exists -n ${STORAGE_CONTAINER_NAME} --account-name ${STORAGE_ACCOUNT_NAME} -o json | jq -r .exists)
if ! $exists ; then
    echo "creating new storage container: ${STORAGE_CONTAINER_NAME}..."
    key=$(az storage account keys list -n $STORAGE_ACCOUNT_NAME -g $RESOURCE_GROUP_NAME -o json | jq -r '.[0].value')
    az storage container create --name $STORAGE_CONTAINER_NAME --account-name $STORAGE_ACCOUNT_NAME --account-key=$key --public-access=blob
else
    echo "storage container '${STORAGE_CONTAINER_NAME}' already exists."
fi

temp_dir=$(mktemp -d -t aks-engine-XXXXXXXX)
zip_name=windows-provisioning-scripts-${CREATE_TIME}.zip
zip_path="${temp_dir}/$zip_name"
echo "creating provisioning ${zip_path}..."
zip -j $zip_path ./staging/provisioning/windows/*

echo "uploading zip to storage account..."
az storage copy --source-local-path $zip_path --destination-account-name $STORAGE_ACCOUNT_NAME --destination-container $STORAGE_CONTAINER_NAME

rm -r -f $temp_dir
echo ""
echo "https://${STORAGE_ACCOUNT_NAME}.blob.core.windows.net/${STORAGE_CONTAINER_NAME}/${zip_name}"