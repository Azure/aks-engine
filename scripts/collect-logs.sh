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
compressLogsDirectory
