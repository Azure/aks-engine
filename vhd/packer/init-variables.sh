#!/bin/bash -e

CDIR=$(dirname "${BASH_SOURCE}")

SETTINGS_JSON="${SETTINGS_JSON:-./settings.json}"
SP_JSON="${SP_JSON:-./sp.json}"
SUBSCRIPTION_ID="${SUBSCRIPTION_ID:-$(az account show -o json --query="id" | tr -d '"')}"
CREATE_TIME="$(date +%s)"
STORAGE_ACCOUNT_NAME="aksimages${CREATE_TIME}${RANDOM}"


echo "Subscription ID: ${SUBSCRIPTION_ID}"
echo "Service Principal Path: ${SP_JSON}"

if [ -a "${SP_JSON}" ]; then
	echo "Existing credentials file found."
	exit 0
elif [ -z "${CLIENT_ID}" ]; then
	echo "Service principal not found! Generating one @ ${SP_JSON}"
	az ad sp create-for-rbac -n aks-images-packer${CREATE_TIME} -o json > ${SP_JSON}
	CLIENT_ID=$(jq -r .appId ${SP_JSON})
	CLIENT_SECRET=$(jq -r .password ${SP_JSON})
	TENANT_ID=$(jq -r .tenant ${SP_JSON})
fi

avail=$(az storage account check-name -n ${STORAGE_ACCOUNT_NAME} -o json | jq -r .nameAvailable)
if $avail ; then
	echo "creating new storage account ${STORAGE_ACCOUNT_NAME}"
	az storage account create -n $STORAGE_ACCOUNT_NAME -g $AZURE_RESOURCE_GROUP_NAME --sku "Standard_RAGRS" --tags "now=${CREATE_TIME}"
	echo "creating new container system"
	key=$(az storage account keys list -n $STORAGE_ACCOUNT_NAME -g $AZURE_RESOURCE_GROUP_NAME | jq -r '.[0].value')
	az storage container create --name system --account-key=$key --account-name=$STORAGE_ACCOUNT_NAME
else
	echo "storage account ${STORAGE_ACCOUNT_NAME} already exists."
fi

if [ -z "${CLIENT_ID}" ]; then
	echo "CLIENT_ID was not set! Something happened when generating the service principal or when trying to read the sp file!"
	exit 1
fi

if [ -z "${CLIENT_SECRET}" ]; then
	echo "CLIENT_SECRET was not set! Something happened when generating the service principal or when trying to read the sp file!"
	exit 1
fi

if [ -z "${TENANT_ID}" ]; then
	echo "TENANT_ID was not set! Something happened when generating the service principal or when trying to read the sp file!"
	exit 1
fi

echo "storage name: ${STORAGE_ACCOUNT_NAME}"

cat <<EOF > vhd/packer/settings.json
{
  "subscription_id":  "${SUBSCRIPTION_ID}",
  "client_id": "${CLIENT_ID}",
  "client_secret": "${CLIENT_SECRET}",
  "tenant_id":      "${TENANT_ID}",
  "resource_group_name": "${AZURE_RESOURCE_GROUP_NAME}",
  "location": "${AZURE_LOCATION}",
  "storage_account_name": "${STORAGE_ACCOUNT_NAME}",
  "vm_size": "${AZURE_VM_SIZE}",
  "create_time": "${CREATE_TIME}"
}
EOF

cat vhd/packer/settings.json
