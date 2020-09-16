#!/bin/bash

{{- if IsCustomCloudProfile}}
  {{- if not IsAzureStackCloud}}
ensureCustomCloudRootCertificates() {
    CUSTOM_CLOUD_ROOT_CERTIFICATES="{{GetCustomCloudRootCertificates}}"
    KUBE_CONTROLLER_MANAGER_FILE=/etc/kubernetes/manifests/kube-controller-manager.yaml

    if [ ! -z $CUSTOM_CLOUD_ROOT_CERTIFICATES ]; then
        # Replace placeholder for ssl binding
        if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
            sed -i "s|<volumessl>|- name: ssl\n      hostPath:\n        path: \\/etc\\/ssl\\/certs|g" $KUBE_CONTROLLER_MANAGER_FILE
            sed -i "s|<volumeMountssl>|- name: ssl\n          mountPath: \\/etc\\/ssl\\/certs\n          readOnly: true|g" $KUBE_CONTROLLER_MANAGER_FILE
        fi

        local i=1
        for cert in $(echo $CUSTOM_CLOUD_ROOT_CERTIFICATES | tr ',' '\n')
        do
            echo $cert | base64 -d > "/usr/local/share/ca-certificates/customCloudRootCertificate$i.crt"
            ((i++))
        done

        update-ca-certificates
    else
        if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
            # remove the placeholder for ssl binding
            sed -i "/<volumessl>/d" $KUBE_CONTROLLER_MANAGER_FILE
            sed -i "/<volumeMountssl>/d" $KUBE_CONTROLLER_MANAGER_FILE
        fi
    fi
}

ensureCustomCloudSourcesList() {
    CUSTOM_CLOUD_SOURCES_LIST="{{GetCustomCloudSourcesList}}"

    if [ ! -z $CUSTOM_CLOUD_SOURCES_LIST ]; then
        # Just in case, let's take a back up before we overwrite
        cp /etc/apt/sources.list /etc/apt/sources.list.backup
        echo $CUSTOM_CLOUD_SOURCES_LIST | base64 -d > /etc/apt/sources.list
    fi
}
  {{end}}

configureK8sCustomCloud() {
  {{- if IsAzureStackCloud}}
  export -f ensureAzureStackCertificates
  retrycmd 60 10 30 bash -c ensureAzureStackCertificates
  set +x
  # When AUTHENTICATION_METHOD is client_certificate, the certificate is stored into key valut,
  # And SERVICE_PRINCIPAL_CLIENT_SECRET will be the following json payload with based64 encode
  #{
  #    "data": "$pfxAsBase64EncodedString",
  #    "dataType" :"pfx",
  #    "password": "$password"
  #}
  if [[ ${AUTHENTICATION_METHOD,,} == "client_certificate" ]]; then
    SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED=$(echo ${SERVICE_PRINCIPAL_CLIENT_SECRET} | base64 --decode)
    SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .data)
    SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .password)

    # trim the starting and ending "
    SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT#'"'}
    SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT%'"'}

    SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD#'"'}
    SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD%'"'}

    KUBERNETES_FILE_DIR=$(dirname "${AZURE_JSON_PATH}")
    K8S_CLIENT_CERT_PATH="${KUBERNETES_FILE_DIR}/k8s_auth_certificate.pfx"
    echo $SERVICE_PRINCIPAL_CLIENT_SECRET_CERT | base64 --decode >$K8S_CLIENT_CERT_PATH
    # shellcheck disable=SC2002,SC2005
    echo $(cat "${AZURE_JSON_PATH}" |
      jq --arg K8S_CLIENT_CERT_PATH ${K8S_CLIENT_CERT_PATH} '. + {aadClientCertPath:($K8S_CLIENT_CERT_PATH)}' |
      jq --arg SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD ${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD} '. + {aadClientCertPassword:($SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD)}' |
      jq 'del(.aadClientSecret)') >${AZURE_JSON_PATH}
  fi

  if [[ ${IDENTITY_SYSTEM,,} == "adfs" ]]; then
    # update the tenent id for ADFS environment.
    # shellcheck disable=SC2002,SC2005
    echo $(cat "${AZURE_JSON_PATH}" | jq '.tenantId = "adfs"') >${AZURE_JSON_PATH}
  fi
  set -x

  {{- if not IsAzureCNI}}
  # Decrease eth0 MTU to mitigate Azure Stack's NRP issue
  echo "iface eth0 inet dhcp" | sudo tee -a /etc/network/interfaces
  echo "    post-up /sbin/ifconfig eth0 mtu 1350" | sudo tee -a /etc/network/interfaces
  ifconfig eth0 mtu 1350
  {{end}}

  {{else}}
  ensureCustomCloudRootCertificates
  ensureCustomCloudSourcesList
  {{end}}
}
{{end}}

{{- if IsAzureStackCloud}}
ensureAzureStackCertificates() {
  AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
  AZURESTACK_RESOURCE_MANAGER_ENDPOINT=$(jq .resourceManagerEndpoint $AZURESTACK_ENVIRONMENT_JSON_PATH | tr -d '"')
  AZURESTACK_RESOURCE_METADATA_ENDPOINT="$AZURESTACK_RESOURCE_MANAGER_ENDPOINT/metadata/endpoints?api-version=2015-01-01"
  curl $AZURESTACK_RESOURCE_METADATA_ENDPOINT
  CURL_RETURNCODE=$?
  KUBE_CONTROLLER_MANAGER_FILE=/etc/kubernetes/manifests/kube-controller-manager.yaml
  if [ $CURL_RETURNCODE != 0 ]; then
    # Replace placeholder for ssl binding
    if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
      sed -i "s|<volumessl>|- name: ssl\n      hostPath:\n        path: \\/etc\\/ssl\\/certs|g" $KUBE_CONTROLLER_MANAGER_FILE
      sed -i "s|<volumeMountssl>|- name: ssl\n          mountPath: \\/etc\\/ssl\\/certs\n          readOnly: true|g" $KUBE_CONTROLLER_MANAGER_FILE
    fi

    # Copying the AzureStack root certificate to the appropriate store to be updated.
    AZURESTACK_ROOT_CERTIFICATE_SOURCE_PATH="/var/lib/waagent/Certificates.pem"
    AZURESTACK_ROOT_CERTIFICATE__DEST_PATH="/usr/local/share/ca-certificates/azsCertificate.crt"
    cp $AZURESTACK_ROOT_CERTIFICATE_SOURCE_PATH $AZURESTACK_ROOT_CERTIFICATE__DEST_PATH
    update-ca-certificates
  else
    if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
      # the ARM resource manager endpoint binding certificate is trusted, remove the placeholder for ssl binding
      sed -i "/<volumessl>/d" $KUBE_CONTROLLER_MANAGER_FILE
      sed -i "/<volumeMountssl>/d" $KUBE_CONTROLLER_MANAGER_FILE
    fi
  fi

  # ensureAzureStackCertificates will be retried if the exit code is not 0
  curl $AZURESTACK_RESOURCE_METADATA_ENDPOINT
  exit $?
}

configureAzureStackInterfaces() {
  set +x

  NETWORK_INTERFACES_FILE="/etc/kubernetes/network_interfaces.json"
  AZURE_CNI_CONFIG_FILE="/etc/kubernetes/interfaces.json"
  AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
  SERVICE_MANAGEMENT_ENDPOINT=$(jq -r '.serviceManagementEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
  ACTIVE_DIRECTORY_ENDPOINT=$(jq -r '.activeDirectoryEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
  RESOURCE_MANAGER_ENDPOINT=$(jq -r '.resourceManagerEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})

  if [[ ${IDENTITY_SYSTEM,,} == "adfs" ]]; then
    TOKEN_URL="${ACTIVE_DIRECTORY_ENDPOINT}adfs/oauth2/token"
  else
    TOKEN_URL="${ACTIVE_DIRECTORY_ENDPOINT}${TENANT_ID}/oauth2/token"
  fi

  echo "Generating token for Azure Resource Manager"
  echo "------------------------------------------------------------------------"
  echo "Parameters"
  echo "------------------------------------------------------------------------"
  echo "SERVICE_PRINCIPAL_CLIENT_ID:     ..."
  echo "SERVICE_PRINCIPAL_CLIENT_SECRET: ..."
  echo "SERVICE_MANAGEMENT_ENDPOINT:     $SERVICE_MANAGEMENT_ENDPOINT"
  echo "ACTIVE_DIRECTORY_ENDPOINT:       $ACTIVE_DIRECTORY_ENDPOINT"
  echo "TENANT_ID:                       $TENANT_ID"
  echo "IDENTITY_SYSTEM:                 $IDENTITY_SYSTEM"
  echo "TOKEN_URL:                       $TOKEN_URL"
  echo "------------------------------------------------------------------------"

  TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X POST \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -d "grant_type=client_credentials" \
    -d "client_id=$SERVICE_PRINCIPAL_CLIENT_ID" \
    --data-urlencode "client_secret=$SERVICE_PRINCIPAL_CLIENT_SECRET" \
    --data-urlencode "resource=$SERVICE_MANAGEMENT_ENDPOINT" \
    ${TOKEN_URL} | jq '.access_token' | xargs)

  if [[ -z $TOKEN ]]; then
    echo "Error generating token for Azure Resource Manager"
    exit 120
  fi

  echo "Fetching network interface configuration for node"
  echo "------------------------------------------------------------------------"
  echo "Parameters"
  echo "------------------------------------------------------------------------"
  echo "RESOURCE_MANAGER_ENDPOINT: $RESOURCE_MANAGER_ENDPOINT"
  echo "SUBSCRIPTION_ID:           $SUBSCRIPTION_ID"
  echo "RESOURCE_GROUP:            $RESOURCE_GROUP"
  echo "NETWORK_API_VERSION:       $NETWORK_API_VERSION"
  echo "------------------------------------------------------------------------"

  curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X GET \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    "${RESOURCE_MANAGER_ENDPOINT}subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP/providers/Microsoft.Network/networkInterfaces?api-version=$NETWORK_API_VERSION" >${NETWORK_INTERFACES_FILE}

  if [[ ! -s ${NETWORK_INTERFACES_FILE} ]]; then
    echo "Error fetching network interface configuration for node"
    exit 121
  fi

  echo "Generating Azure CNI interface file"

  mapfile -t local_interfaces < <(cat /sys/class/net/*/address | tr -d : | sed 's/.*/\U&/g')

  SDN_INTERFACES=$(jq ".value | map(select(.properties != null) | select(.properties.macAddress != null) | select(.properties.macAddress | inside(\"${local_interfaces[*]}\"))) | map(select((.properties.ipConfigurations | length) > 0))" ${NETWORK_INTERFACES_FILE})

  if [[ -z $SDN_INTERFACES ]]; then
      echo "Error extracting the SDN interfaces from the network interfaces file"
      exit 123
  fi

  AZURE_CNI_CONFIG=$(echo ${SDN_INTERFACES} | jq "{Interfaces: [.[] | {MacAddress: .properties.macAddress, IsPrimary: .properties.primary, IPSubnets: [{Prefix: .properties.ipConfigurations[0].properties.subnet.id, IPAddresses: .properties.ipConfigurations | [.[] | {Address: .properties.privateIPAddress, IsPrimary: .properties.primary}]}]}]}")

  mapfile -t SUBNET_IDS < <(echo ${SDN_INTERFACES} | jq '[.[].properties.ipConfigurations[0].properties.subnet.id] | unique | .[]' -r)

  for SUBNET_ID in "${SUBNET_IDS[@]}"; do
    SUBNET_PREFIX=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X GET \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json" \
      "${RESOURCE_MANAGER_ENDPOINT}${SUBNET_ID:1}?api-version=$NETWORK_API_VERSION" |
      jq '.properties.addressPrefix' -r)

    if [[ -z $SUBNET_PREFIX ]]; then
      echo "Error fetching the subnet address prefix for a subnet ID"
      exit 122
    fi

    # shellcheck disable=SC2001
    AZURE_CNI_CONFIG=$(echo ${AZURE_CNI_CONFIG} | sed "s|$SUBNET_ID|$SUBNET_PREFIX|g")
  done

  echo ${AZURE_CNI_CONFIG} >${AZURE_CNI_CONFIG_FILE}

  chmod 0444 ${AZURE_CNI_CONFIG_FILE}

  set -x
}
{{end}}
#EOF
