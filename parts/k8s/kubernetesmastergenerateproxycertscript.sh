#!/bin/bash

source /opt/azure/containers/provision_source.sh

PROXY_CA_KEY="${PROXY_CA_KEY:=/tmp/proxy-client-ca.key}"
PROXY_CRT="${PROXY_CRT:=/tmp/proxy-client-ca.crt}"
PROXY_CLIENT_KEY="${PROXY_CLIENT_KEY:=/tmp/proxy-client.key}"
PROXY_CLIENT_CSR="${PROXY_CLIENT_CSR:=/tmp/proxy-client.csr}"
PROXY_CLIENT_CRT="${PROXY_CLIENT_CRT:=/tmp/proxy-client.crt}"
ETCD_REQUESTHEADER_CLIENT_CA="${ETCD_REQUESTHEADER_CLIENT_CA:=/proxycerts/requestheader-client-ca-file}"
ETCD_PROXY_CERT="${ETCD_PROXY_CERT:=/proxycerts/proxy-client-cert-file}"
ETCD_PROXY_KEY="${ETCD_PROXY_KEY:=/proxycerts/proxy-client-key-file}"
K8S_PROXY_CA_CRT_FILEPATH="${K8S_PROXY_CA_CRT_FILEPATH:=/etc/kubernetes/certs/proxy-ca.crt}"
K8S_PROXY_KEY_FILEPATH="${K8S_PROXY_KEY_FILEPATH:=/etc/kubernetes/certs/proxy.key}"
K8S_PROXY_CRT_FILEPATH="${K8S_PROXY_CRT_FILEPATH:=/etc/kubernetes/certs/proxy.crt}"

if [[ -z "${COSMOS_URI}" ]]; then
  ETCDCTL_ENDPOINTS="${ETCDCTL_ENDPOINTS:=https://127.0.0.1:2379}"
  ETCDCTL_CA_FILE="${ETCDCTL_CA_FILE:=/etc/kubernetes/certs/ca.crt}"
  ETCD_CA_PARAM="--cacert=${ETCDCTL_CA_FILE}"
else
  ETCDCTL_ENDPOINTS="${ETCDCTL_ENDPOINTS:=https://${COSMOS_URI}:2379}"
  ETCD_CA_PARAM=""
  #we don't provide etcd ca cert for cosmos
fi
ETCDCTL_KEY_FILE="${ETCDCTL_KEY_FILE:=/etc/kubernetes/certs/etcdclient.key}"
ETCDCTL_CERT_FILE="${ETCDCTL_CERT_FILE:=/etc/kubernetes/certs/etcdclient.crt}"

ETCDCTL_PARAMS="--command-timeout=30s --cert=${ETCDCTL_CERT_FILE} --key=${ETCDCTL_KEY_FILE} ${ETCD_CA_PARAM} --endpoints=${ETCDCTL_ENDPOINTS}"
export RANDFILE=$(mktemp)

# generate root CA
openssl genrsa -out $PROXY_CA_KEY 2048
openssl req -new -x509 -days 1826 -key $PROXY_CA_KEY -out $PROXY_CRT -subj '/CN=proxyClientCA'
# generate new cert
openssl genrsa -out $PROXY_CLIENT_KEY 2048
openssl req -new -key $PROXY_CLIENT_KEY -out $PROXY_CLIENT_CSR -subj '/CN=aggregator/O=system:masters'
openssl x509 -req -days 730 -in $PROXY_CLIENT_CSR -CA $PROXY_CRT -CAkey $PROXY_CA_KEY -set_serial 02 -out $PROXY_CLIENT_CRT

write_certs_to_disk() {
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_REQUESTHEADER_CLIENT_CA | sed 1d > $K8S_PROXY_CA_CRT_FILEPATH
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_KEY | sed 1d > $K8S_PROXY_KEY_FILEPATH
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_CERT | sed 1d > $K8S_PROXY_CRT_FILEPATH
    # Remove whitespace padding at beginning of 1st line
    sed -i '1s/\s//' $K8S_PROXY_CA_CRT_FILEPATH $K8S_PROXY_CRT_FILEPATH $K8S_PROXY_KEY_FILEPATH
    chmod 600 $K8S_PROXY_KEY_FILEPATH
}

write_certs_to_disk_with_retry() {
    for i in 1 2 3 4 5 6 7 8 9 10 11 12; do
        write_certs_to_disk
        [ $? -eq 0  ] && break || sleep 5
    done
}
is_cluster_healthy(){
    for i in 1 2 3 4 5 6 7 8 9 10 11 12; do
        ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} endpoint health
        [ $? -eq 0  ] && break || sleep 5
    done
}
# block until all etcd is ready
is_cluster_healthy 
# Make etcd keys, adding a leading whitespace because etcd won't accept a val that begins with a '-' (hyphen)!
# etcdctl will output the data it's given, stdout is redirected to dev null to avoid capturing sensitive data in logs
if ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_REQUESTHEADER_CLIENT_CA " $(cat ${PROXY_CRT})" > /dev/null 2>&1; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_PROXY_KEY " $(cat ${PROXY_CLIENT_KEY})" > /dev/null 2>&1
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_PROXY_CERT " $(cat ${PROXY_CLIENT_CRT})" > /dev/null 2>&1
    sleep 5
    write_certs_to_disk_with_retry
# If the etcdtl mk command failed, that means the key already exists
else
    sleep 5
    write_certs_to_disk_with_retry
fi
