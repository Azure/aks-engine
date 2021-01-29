#!/usr/bin/env bash

set +x
set -euo pipefail

AZURE_JSON_PATH="/etc/kubernetes/azure.json"
SERVICE_PRINCIPAL_CLIENT_ID=$(jq -r '.aadClientId' ${AZURE_JSON_PATH})
SERVICE_PRINCIPAL_CLIENT_SECRET=$(jq -r '.aadClientSecret' ${AZURE_JSON_PATH})
TENANT_ID=$(jq -r '.tenantId' ${AZURE_JSON_PATH})
KMS_KEYVAULT_NAME=$(jq -r '.providerVaultName' ${AZURE_JSON_PATH})
KMS_KEY_NAME=$(jq -r '.providerKeyName' ${AZURE_JSON_PATH})
USER_ASSIGNED_IDENTITY_ID=$(jq -r '.userAssignedIdentityID' ${AZURE_JSON_PATH})
PROVIDER_KEY_VERSION=$(jq -r '.providerKeyVersion' ${AZURE_JSON_PATH})
AZURE_CLOUD=$(jq -r '.cloud' ${AZURE_JSON_PATH})

# get the required parameters specific for cloud
if [[ $AZURE_CLOUD == "AzurePublicCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.microsoftonline.com/"
    KEYVAULT_DNS_SUFFIX="vault.azure.net"
elif [[ $AZURE_CLOUD == "AzureChinaCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.chinacloudapi.cn/"
    KEYVAULT_DNS_SUFFIX="vault.azure.cn"
elif [[ $AZURE_CLOUD == "AzureGermanCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.microsoftonline.de/"
    KEYVAULT_DNS_SUFFIX="vault.microsoftazure.de"
elif [[ $AZURE_CLOUD == "AzureUSGovernmentCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.microsoftonline.us/"
    KEYVAULT_DNS_SUFFIX="vault.usgovcloudapi.net"
elif [[ $AZURE_CLOUD == "AzureStackCloud" ]]; then
    AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
    ACTIVE_DIRECTORY_ENDPOINT=$(jq -r '.activeDirectoryEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
    KEYVAULT_DNS_SUFFIX=$(jq -r '.keyVaultDNSSuffix' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
else
    echo "Invalid cloud name"
    exit 120
fi

TOKEN_URL="${ACTIVE_DIRECTORY_ENDPOINT}${TENANT_ID}/oauth2/token"
KEYVAULT_URL="https://${KMS_KEYVAULT_NAME}.${KEYVAULT_DNS_SUFFIX}/keys/${KMS_KEY_NAME}/versions?maxresults=1&api-version=7.1"
KEYVAULT_ENDPOINT="https://${KEYVAULT_DNS_SUFFIX}"
KMS_KUBERNETES_FILE=/etc/kubernetes/manifests/kube-azure-kms.yaml

# provider key version already exists
# this will be the case for BYOK
if [[ -n $PROVIDER_KEY_VERSION ]]; then
    echo "KMS provider key version already exists"
    exit 0
fi

echo "Generating token for Azure Key Vault"
echo "------------------------------------------------------------------------"
echo "Parameters"
echo "------------------------------------------------------------------------"
echo "SERVICE_PRINCIPAL_CLIENT_ID:     ..."
echo "SERVICE_PRINCIPAL_CLIENT_SECRET: ..."
echo "ACTIVE_DIRECTORY_ENDPOINT:       $ACTIVE_DIRECTORY_ENDPOINT"
echo "TENANT_ID:                       $TENANT_ID"
echo "TOKEN_URL:                       $TOKEN_URL"
echo "SCOPE:                           $KEYVAULT_ENDPOINT"
echo "------------------------------------------------------------------------"

if [[ $SERVICE_PRINCIPAL_CLIENT_ID == "msi" ]] && [[ $SERVICE_PRINCIPAL_CLIENT_SECRET == "msi" ]]; then
    if [[ -z $USER_ASSIGNED_IDENTITY_ID ]]; then
        # using system-assigned identity to access keyvault
        TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 \
            -H Metadata:true \
            "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=$KEYVAULT_ENDPOINT" | jq '.access_token' | xargs)
    else
        # using user-assigned managed identity to access keyvault
        TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 \
            -H Metadata:true \
            "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&client_id=$USER_ASSIGNED_IDENTITY_ID&resource=$KEYVAULT_ENDPOINT" | jq '.access_token' | xargs)
    fi
else
    # use service principal token to access keyvault
    TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X POST \
        -H "Content-Type: application/x-www-form-urlencoded" \
        -d "grant_type=client_credentials" \
        -d "client_id=$SERVICE_PRINCIPAL_CLIENT_ID" \
        --data-urlencode "client_secret=$SERVICE_PRINCIPAL_CLIENT_SECRET" \
        --data-urlencode "resource=$KEYVAULT_ENDPOINT" \
        ${TOKEN_URL} | jq '.access_token' | xargs)
fi


if [[ -z $TOKEN ]]; then
    echo "Error generating token for Azure Keyvault"
    exit 120
fi

# Get the keyID for the kms key created as part of cluster bootstrap
KEY_ID=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f \
    ${KEYVAULT_URL} -H "Authorization: Bearer ${TOKEN}" | jq '.value[0].kid' | xargs)

if [[ -z "$KEY_ID" || "$KEY_ID" == "null" ]]; then
    echo "Error getting the kms key version"
    exit 120
fi

# KID format: https://<keyvault name>.vault.azure.net/keys/<key name>/<key version>
# Example KID: "https://akv0112master.vault.azure.net/keys/k8s/128a3d9956bc44feb6a0e2c2f35b732c"
KEY_VERSION=${KEY_ID##*/}

# Set the version in azure.json
if [ -f $AZURE_JSON_PATH ]; then
    # once the version is set in azure.json, kms plugin will just default to using the key
    # this will be changed in upcoming kms release to set the version as container args
    tmpDir=$(mktemp -d "$(pwd)/XXX")
    jq --arg KEY_VERSION ${KEY_VERSION} '.providerKeyVersion=($KEY_VERSION)' "$AZURE_JSON_PATH" > $tmpDir/tmp
    mv $tmpDir/tmp "$AZURE_JSON_PATH"
    # set the permissions for azure json
    chmod 0600 "$AZURE_JSON_PATH"
    chown root:root "$AZURE_JSON_PATH"
    rm -Rf $tmpDir
fi

set -x
#EOF
