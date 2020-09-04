#!/usr/bin/env bash

# This script originated at https://github.com/kubernetes/kubernetes/blob/master/cluster/gce/gci/health-monitor.sh
# and has been modified for aks-engine.

set -o nounset
set -o pipefail

container_runtime_monitoring() {
  local -r container_runtime_name="${CONTAINER_RUNTIME:-docker}"
  local healthcheck_command="docker ps"
  if [[ ${CONTAINER_RUNTIME} == "containerd" ]]; then
    healthcheck_command="ctr -n k8s.io containers ls"
  fi

  while true; do
    if ! timeout 60 ${healthcheck_command} >/dev/null; then
      echo "Container runtime ${container_runtime_name} failed!"
      if [[ $container_runtime_name == "docker" ]]; then
        pkill -SIGUSR1 dockerd
      fi
      if [[ $container_runtime_name == "containerd" ]]; then
        pkill -SIGUSR1 containerd
      fi
      systemctl kill --kill-who=main "${container_runtime_name}"
      sleep 120
      if ! systemctl is-active ${container_runtime_name}; then
        systemctl start ${container_runtime_name}
      fi
    else
      sleep "${SLEEP_SECONDS}"
    fi
  done
}

kubelet_monitoring() {
  echo "Wait for 2 minutes for kubelet to be functional"
  sleep 120
  local -r max_seconds=10
  local output=""
  while true; do
    if ! output=$(curl -m "${max_seconds}" -f -s -S http://127.0.0.1:${HEALTHZPORT}/healthz 2>&1); then
      echo $output
      echo "Kubelet is unhealthy!"
      systemctl kill kubelet
      sleep 60
      if ! systemctl is-active kubelet; then
        systemctl start kubelet
      fi
    else
      sleep "${SLEEP_SECONDS}"
    fi
  done
}

etcd_monitoring() {
  local -r max_seconds=10
  local output=""
  local private_ip
  private_ip=$(hostname -i)
  local endpoint="https://${private_ip}:2379"
  while true; do
    if ! output=$(curl -s -S -m "${max_seconds}" --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${endpoint}/v2/machines); then
      echo $output
      echo "etcd is unhealthy!"
      systemctl kill etcd
      sleep 60
      if ! systemctl is-active etcd; then
        systemctl start etcd
      fi
    else
      sleep "${SLEEP_SECONDS}"
    fi
  done
}

if [[ $# -ne 1 ]]; then
  echo "Usage: health-monitor.sh <container-runtime/kubelet>"
  exit 1
fi

KUBE_ENV="/etc/default/kube-env"
if [[ -e ${KUBE_ENV} ]]; then
  source "${KUBE_ENV}"
fi

SLEEP_SECONDS=10
component=$1
echo "Start kubernetes health monitoring for ${component}"

if [[ ${component} == "container-runtime" ]]; then
  container_runtime_monitoring
elif [[ ${component} == "kubelet" ]]; then
  kubelet_monitoring
elif [[ ${component} == "etcd" ]]; then
  etcd_monitoring
else
  echo "Health monitoring for component ${component} is not supported!"
fi
