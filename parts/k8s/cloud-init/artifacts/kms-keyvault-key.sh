#!/bin/bash

ensureKMSKeyvaultKey() {
  local kms_file_path="/etc/kubernetes/manifests/kube-azure-kms.yaml"

  if [[ -z "$KMS_PROVIDER_KEY_VERSION" ]]; then
    getKeyVersion
  fi

  # Set the keyvault name, key name and version as args in the kms static pod
  if [ -f ${kms_file_path} ]; then
    sed -i "s|<keyvaultName>|$KMS_PROVIDER_VAULT_NAME|g" $kms_file_path
    sed -i "s|<keyvaultKeyName>|$KMS_PROVIDER_KEY_NAME|g" $kms_file_path
    sed -i "s|<keyvaultKeyVersion>|$KMS_PROVIDER_KEY_VERSION|g" $kms_file_path
  fi
}

getKeyVersion() {
  local azure_cloud="{{GetTargetEnvironment}}"
  local active_directory_endpoint
  local keyvault_dns_suffix

  # get the required parameters specific for cloud
  if [[ ${azure_cloud} == "AzurePublicCloud" ]]; then
      active_directory_endpoint="https://login.microsoftonline.com/"
      keyvault_dns_suffix="vault.azure.net"
  elif [[ ${azure_cloud} == "AzureChinaCloud" ]]; then
      active_directory_endpoint="https://login.chinacloudapi.cn/"
      keyvault_dns_suffix="vault.azure.cn"
  elif [[ ${azure_cloud} == "AzureGermanCloud" ]]; then
      active_directory_endpoint="https://login.microsoftonline.de/"
      keyvault_dns_suffix="vault.microsoftazure.de"
  elif [[ ${azure_cloud} == "AzureUSGovernmentCloud" ]]; then
      active_directory_endpoint="https://login.microsoftonline.us/"
      keyvault_dns_suffix="vault.usgovcloudapi.net"
  elif [[ ${azure_cloud} == "AzureStackCloud" ]]; then
      local azurestack_environment_json_path="/etc/kubernetes/azurestackcloud.json"
      active_directory_endpoint=$(jq -r '.activeDirectoryEndpoint' ${azurestack_environment_json_path})
      keyvault_dns_suffix=$(jq -r '.keyVaultDNSSuffix' ${azurestack_environment_json_path})
  else
      echo "Invalid cloud name"
      exit 120
  fi

  local token_url="${active_directory_endpoint}${TENANT_ID}/oauth2/token"
  local keyvault_url="https://${KMS_PROVIDER_VAULT_NAME}.${keyvault_dns_suffix}/keys/${KMS_PROVIDER_KEY_NAME}/versions?maxresults=1&api-version=7.1"
  local keyvault_endpoint="https://${keyvault_dns_suffix}"

  echo "Generating token for Azure Key Vault"
  echo "------------------------------------------------------------------------"
  echo "Parameters"
  echo "------------------------------------------------------------------------"
  echo "SERVICE_PRINCIPAL_CLIENT_ID:     ..."
  echo "SERVICE_PRINCIPAL_CLIENT_SECRET: ..."
  echo "ACTIVE_DIRECTORY_ENDPOINT:       ${active_directory_endpoint}"
  echo "TENANT_ID:                       $TENANT_ID"
  echo "TOKEN_URL:                       ${token_url}"
  echo "SCOPE:                           ${keyvault_endpoint}"
  echo "------------------------------------------------------------------------"

  if [[ $SERVICE_PRINCIPAL_CLIENT_ID == "msi" ]] && [[ $SERVICE_PRINCIPAL_CLIENT_SECRET == "msi" ]]; then
      if [[ -z $USER_ASSIGNED_IDENTITY_ID ]]; then
          # using system-assigned identity to access keyvault
          TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 \
              -H Metadata:true \
              "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=${keyvault_endpoint}" | jq '.access_token' | xargs)
      else
          # using user-assigned managed identity to access keyvault
          TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 \
              -H Metadata:true \
              "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&client_id=$USER_ASSIGNED_IDENTITY_ID&resource=${keyvault_endpoint}" | jq '.access_token' | xargs)
      fi
  else
      # use service principal token to access keyvault
      TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X POST \
          -H "Content-Type: application/x-www-form-urlencoded" \
          -d "grant_type=client_credentials" \
          -d "client_id=$SERVICE_PRINCIPAL_CLIENT_ID" \
          --data-urlencode "client_secret=$SERVICE_PRINCIPAL_CLIENT_SECRET" \
          --data-urlencode "resource=${keyvault_endpoint}" \
          "${token_url}" | jq '.access_token' | xargs)
  fi


  if [[ -z $TOKEN ]]; then
      echo "Error generating token for Azure Keyvault"
      exit 120
  fi

  # Get the keyID for the kms key created as part of cluster bootstrap
  local key_id
  key_id=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f \
      "${keyvault_url}" -H "Authorization: Bearer ${TOKEN}" | jq '.value[0].kid' | xargs)

  if [[ -z "${key_id}" || "${key_id}" == "null" ]]; then
      echo "Error getting the kms key version"
      exit 120
  fi

  # KID format: https://<keyvault name>.vault.azure.net/keys/<key name>/<key version>
  # Example KID: "https://akv0112master.vault.azure.net/keys/k8s/128a3d9956bc44feb6a0e2c2f35b732c"
  KMS_PROVIDER_KEY_VERSION=${key_id##*/}
}
#EOF
