#!/bin/bash

source /opt/azure/containers/provision_source.sh

clusterInfo() {
    FIRST_MASTER_READY=$(kubectl get nodes | grep k8s-master | grep Ready | sort | head -n 1 | cut -d ' ' -f 1)
    if [[ "${FIRST_MASTER_READY}" == "${HOSTNAME}" ]]; then
        retrycmd_no_stats 3 5 120 kubectl cluster-info dump --namespace=kube-system --output-directory=${OUTDIR}/cluster-info
    fi
}

collectCloudProviderJson() {
    local DIR=${OUTDIR}/etc/kubernetes
    mkdir -p ${DIR}
    if [ -f /etc/kubernetes/azure.json ]; then
        jq . /etc/kubernetes/azure.json | grep -v aadClient > ${DIR}/azure.json
    fi
    if [ -f /etc/kubernetes/azurestackcloud.json ]; then
        jq . /etc/kubernetes/azurestackcloud.json > ${DIR}/azurestackcloud.json
    fi
}

collectDirLogs() {
    local DIR=${OUTDIR}${1}
    if [ -d ${1} ]; then
        mkdir -p ${DIR}
        cp ${1}/*.log ${DIR}
    fi
}

collectDir() {
    local DIR=${OUTDIR}${1}
    if [ -d ${1} ]; then
        mkdir -p ${DIR}
        cp ${1}/* ${DIR}
    fi
}

collectDaemonLogs() {
    local DIR=${OUTDIR}/daemons
    mkdir -p ${DIR}
    if systemctl list-units --no-pager | grep -q ${1}; then
        timeout 15 journalctl --utc -o short-iso --no-pager -u ${1} &>> ${DIR}/${1}.log
    fi
}

compressLogsDirectory() {
    sync
    ZIP="/tmp/logs.zip"
    rm -f ${ZIP}
    (cd ${OUTDIR}/.. && zip -q -r ${ZIP} ${HOSTNAME})
}

# AzS headers
stackfy() {
    # Azure Stack can store/ingest these files in a Edge Kusto DB.
    # Data team asked for metadata to simplify queries. It is added as a header to each file.
    RESOURCE_GROUP=$(sudo jq -r '.resourceGroup' /etc/kubernetes/azure.json)
    SUB_ID=$(sudo jq -r '.subscriptionId' /etc/kubernetes/azure.json)
    TENANT_ID=$(sudo jq -r '.tenantId' /etc/kubernetes/azure.json)
    if [ "${TENANT_ID}" == "adfs" ]; then
        TENANT_ID=$(sudo jq -r '.serviceManagementEndpoint' /etc/kubernetes/azurestackcloud.json | cut -d / -f4)
    fi

    stackfyKubeletLog
    stackfyMobyLog
    stackfyEtcdLog
    stackfyFileNames /var/log
    stackfyFileNames /var/log/azure
    stackfyClusterInfo
    stackfyNetwork
}

stackfyKubeletLog() {
    KUBELET_VERSION=$(kubelet --version | grep -oh -E 'v.*$')
    KUBELET_VERBOSITY=$(grep -e '--v=[0-9]' -oh | grep -e '[0-9]' -oh /etc/systemd/system/kubelet.service | head -n 1)
    KUBELET_LOG_FILE=${OUTDIR}/daemons/k8s-kubelet.log
    
    {
        echo "== BEGIN HEADER ==";
        echo "Type: Daemon";
        echo "TenantId: ${TENANT_ID}";
        echo "Name: kubelet";
        echo "Version: ${KUBELET_VERSION}";
        echo "Verbosity: ${KUBELET_VERBOSITY}";
        echo "Image: ";
        echo "Hostname: ${HOSTNAME}";
        echo "SubscriptionID: ${SUB_ID}";
        echo "ResourceGroup: ${RESOURCE_GROUP}";
        echo "== END HEADER =="
    } >> ${KUBELET_LOG_FILE}

    cat ${OUTDIR}/daemons/kubelet.service.log >> ${KUBELET_LOG_FILE}
    rm ${OUTDIR}/daemons/kubelet.service.log
}

stackfyMobyLog() {
    DOCKER_VERSION=$(docker version | grep -A 20 "Server:" | grep "Version:" | head -n 1 | cut -d ":" -f 2 | xargs)
    DOCKER_LOG_FILE=${OUTDIR}/daemons/k8s-docker.log
    
    {
        echo "== BEGIN HEADER =="
        echo "Type: Daemon"
        echo "TenantId: ${TENANT_ID}"
        echo "Name: docker"
        echo "Version: ${DOCKER_VERSION}"
        echo "Hostname: ${HOSTNAME}"
        echo "SubscriptionID: ${SUB_ID}"
        echo "ResourceGroup: ${RESOURCE_GROUP}"
        echo "== END HEADER =="
    } >> ${DOCKER_LOG_FILE}

    cat ${OUTDIR}/daemons/docker.service.log >> ${DOCKER_LOG_FILE}
    rm ${OUTDIR}/daemons/docker.service.log
}

stackfyEtcdLog() {
    if [ ! -f /usr/bin/etcd ]; then
        return
    fi

    ETCD_VERSION=$(/usr/bin/etcd --version | grep "etcd Version:" | cut -d ":" -f 2 | xargs)
    ETCD_LOG_FILE=${OUTDIR}/daemons/k8s-etcd.log
    
    {
        echo "== BEGIN HEADER =="
        echo "Type: Daemon"
        echo "TenantId: ${TENANT_ID}"
        echo "Name: etcd"
        echo "Version: ${ETCD_VERSION}"
        echo "Hostname: ${HOSTNAME}"
        echo "SubscriptionID: ${SUB_ID}"
        echo "ResourceGroup: ${RESOURCE_GROUP}"
        echo "== END HEADER =="
    } >> ${ETCD_LOG_FILE}

    cat ${OUTDIR}/daemons/etcd.service.log >> ${ETCD_LOG_FILE}
    rm ${OUTDIR}/daemons/etcd.service.log
}

stackfyClusterInfo() {
    if [ ! -f ${OUTDIR}/cluster-info/nodes.json ]; then
        return
    fi
    mkdir -p ${OUTDIR}/containers
    for SRC in "${OUTDIR}"/cluster-info/kube-system/kube-controller-manager-*/logs.txt; do
        KCM_VERBOSITY=$(grep -e "--v=[0-9]" -oh /etc/kubernetes/manifests/kube-controller-manager.yaml | grep -e "[0-9]" -oh | head -n 1)
        KCM_IMAGE=$(grep image: /etc/kubernetes/manifests/kube-controller-manager.yaml | xargs | cut -f 2 -d " ")
        KCM_DIR=$(dirname ${SRC})
        KCM_NAME=$(basename ${KCM_DIR})
        KCM_LOG_FILE=${OUTDIR}/containers/k8s-${KCM_NAME}.log
        
        {
            echo "== BEGIN HEADER =="
            echo "Type: Container"
            echo "TenantId: ${TENANT_ID}"
            echo "Name: ${KCM_NAME}"
            echo "Hostname: ${HOSTNAME}"
            echo "ContainerID: "
            echo "Image: ${KCM_IMAGE}"
            echo "Verbosity: ${KCM_VERBOSITY}"
            echo "SubscriptionID: ${SUB_ID}"
            echo "ResourceGroup: ${RESOURCE_GROUP}"
            echo "== END HEADER =="
        } >> ${KCM_LOG_FILE}
        
        cat ${SRC} >> ${KCM_LOG_FILE}
    done
}

stackfyNetwork() {
    local DIR=${OUTDIR}/network
    mkdir -p ${DIR}
    # basic name resolution test
    ping ${HOSTNAME} -c 3 &> ${DIR}/k8s-ping.txt
}

stackfyFileNames() {
    local DIR=${OUTDIR}/${1}
    for SRC in "${DIR}"/*.log; do
        NAME=$(basename ${SRC})
        mv ${SRC} ${DIR}/k8s-${NAME}
    done
}

OUTDIR="$(mktemp -d)/${HOSTNAME}"

collectCloudProviderJson
collectDirLogs /var/log
collectDirLogs /var/log/azure
collectDir /etc/kubernetes/manifests
collectDir /etc/kubernetes/addons
collectDaemonLogs kubelet.service
collectDaemonLogs etcd.service
collectDaemonLogs docker.service
clusterInfo

if [ "${AZURE_ENV}" = "AzureStackCloud" ]; then
    stackfy
fi

compressLogsDirectory
