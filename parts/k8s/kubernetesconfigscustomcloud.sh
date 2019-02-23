#!/bin/bash

ensureCertificates() {
    AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
    AZURESTACK_RESOURCE_MANAGER_ENDPOINT=`cat $AZURESTACK_ENVIRONMENT_JSON_PATH | jq .resourceManagerEndpoint | tr -d "\""`
    AZURESTACK_RESOURCE_METADATA_ENDPOINT="$AZURESTACK_RESOURCE_MANAGER_ENDPOINT/metadata/endpoints?api-version=2015-01-01"
    curl $AZURESTACK_RESOURCE_METADATA_ENDPOINT
    CURL_RETURNCODE=$?
    KUBE_CONTROLLER_MANAGER_FILE=/etc/kubernetes/manifests/kube-controller-manager.yaml
    if [ $CURL_RETURNCODE != 0 ]; then
        # Replace placeholder for ssl binding
        if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
            sed -i "s|<volumessl>|- name: ssl\n      hostPath:\n        path: \\/etc\\/ssl\\/certs|g" $KUBE_CONTROLLER_MANAGER_FILE 
            sed -i "s|<volumeMountssl>|- name: "ssl"\n          mountPath: \\/etc\\/ssl\\/certs\n          readOnly: true|g" $KUBE_CONTROLLER_MANAGER_FILE
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
    
    # ensureCertificates will be retried if the exit code is not 0
    curl $AZURESTACK_RESOURCE_METADATA_ENDPOINT
    exit $?
}

configureK8sCustomCloud() {
    export -f ensureCertificates
    retrycmd_if_failure 60 10 30 bash -c ensureCertificates
    set +x
    # When AUTHENTICATION_METHOD is client_certificate, the certificate is stored into key valut, 
    # And SERVICE_PRINCIPAL_CLIENT_SECRET will be the following json payload with based64 encode
    #{
    #    "data": "$pfxAsBase64EncodedString",
    #    "dataType" :"pfx",
    #    "password": "$password"
    #}
    if [[ "${AUTHENTICATION_METHOD,,}" == "client_certificate"  ]]; then
        SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED=`echo ${SERVICE_PRINCIPAL_CLIENT_SECRET} | base64 --decode`
        SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=`echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .data`
        SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=`echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .password`
        
        # trim the starting and ending "
        SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT#"\""} 
        SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT%"\""} 

        SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD#"\""} 
        SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD%"\""}

        KUBERNETES_FILE_DIR=$(dirname "${AZURE_JSON_PATH}")
        K8S_CLIENT_CERT_PATH="${KUBERNETES_FILE_DIR}/k8s_auth_certificate.pfx"
        echo $SERVICE_PRINCIPAL_CLIENT_SECRET_CERT | base64 --decode > $K8S_CLIENT_CERT_PATH
        echo `cat "${AZURE_JSON_PATH}" | \
            jq --arg K8S_CLIENT_CERT_PATH ${K8S_CLIENT_CERT_PATH} '. + {aadClientCertPath:($K8S_CLIENT_CERT_PATH)}' | \
            jq --arg SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD ${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD} '. + {aadClientCertPassword:($SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD)}' |\
            jq 'del(.aadClientSecret)'` > ${AZURE_JSON_PATH}
    fi

    if [[ "${IDENTITY_SYSTEM,,}" == "adfs"  ]]; then
        # update the tenent id for ADFS environment.
        echo `cat "${AZURE_JSON_PATH}" | jq '.tenantId = "adfs"' ` > ${AZURE_JSON_PATH}
    fi
    set -x
}
