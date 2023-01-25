#!/bin/bash -ex

export WD=/etc/kubernetes/rotate-certs
export NEW_CERTS_DIR=${WD}/certs

# copied from cse_helpers.sh, sourcing that file not always works
systemctl_restart() {
  retries=$1; wait_sleep=$2; timeout=$3 svcname=$4
  for i in $(seq 1 $retries); do
    timeout $timeout systemctl daemon-reload
    timeout $timeout systemctl restart $svcname && break ||
      if [ $i -eq $retries ]; then
        return 1
      else
        sleep $wait_sleep
      fi
  done
}

backup() {
  if [ ! -d /etc/kubernetes/certs.bak ]; then
    cp -rp /etc/kubernetes/certs/ /etc/kubernetes/certs.bak
  fi
}

cp_certs() {
  cp -p ${NEW_CERTS_DIR}/etcdpeer* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/etcdclient* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/etcdserver* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/ca.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/client.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/apiserver.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/kubeconfig /home/$(logname)/.kube/config

  rm -f /var/lib/kubelet/pki/kubelet-client-current.pem
}

cp_proxy() {
  source /etc/environment
  local NODE_INDEX
  NODE_INDEX=$(hostname | tail -c 2)
  if [[ $NODE_INDEX == 0 ]]; then
    export OVERRIDE_PROXY_CERTS="true"
  fi
  /etc/kubernetes/generate-proxy-certs.sh
}

agent_certs() {
  cp -p ${NEW_CERTS_DIR}/ca.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/client.* /etc/kubernetes/certs/

  rm -f /var/lib/kubelet/pki/kubelet-client-current.pem
  sync
  sleep 5
  systemctl_restart 10 5 10 kubelet
}

cleanup() {
  rm -rf ${WD}
  rm -rf /etc/kubernetes/certs.bak
}

"$@"
