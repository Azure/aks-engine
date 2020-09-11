#!/usr/bin/env bash
source {{GetCSEHelpersScriptFilepath}}

{{- /* This script originated at https://github.com/kubernetes/kubernetes/blob/master/cluster/gce/gci/health-monitor.sh */}}
{{- /* and has been modified for aks-engine. */}}

set -o nounset
set -o pipefail

container_runtime_monitoring() {
  sleep 300 {{/* Wait for 5 minutes for CRI to be functional/stable */}}
  local cri_name="${CONTAINER_RUNTIME:-docker}" cmd="docker ps"
  if [[ ${CONTAINER_RUNTIME} == "containerd" ]]; then
    cmd="ctr -n k8s.io containers ls"
  fi

  while true; do
    if ! timeout 60 ${cmd} >/dev/null; then
      echo "Container runtime ${cri_name} failed!"
      sleep 10 {{/* Wait 10 more seconds, check again, because the systemd job itself may have already restarted things */}}
      if ! timeout 60 ${cmd} >/dev/null; then
        if [[ $cri_name == "docker" ]]; then
          pkill -SIGUSR1 dockerd
        fi
        if [[ $cri_name == "containerd" ]]; then
          pkill -SIGUSR1 containerd
        fi
        systemctl kill --kill-who=main "${cri_name}"
        sleep 60 {{/* Wait a minute to validate that the systemd job restarted itself after we manually killed the process */}}
        if ! systemctl is-active ${cri_name}; then
          systemctl start ${cri_name}
        fi
      fi
    else
      sleep "${SLEEP_TIME}"
    fi
  done
}

kubelet_monitoring() {
  sleep 300 {{/* Wait for 5 minutes for kubelet to be functional/stable */}}
  local max_seconds=10 output=""
  local monitor_cmd="curl -m ${max_seconds} -f -s -S http://127.0.0.1:${HEALTHZPORT}/healthz"
  while true; do
    if ! output=$(${monitor_cmd} 2>&1); then
      echo $output
      echo "Kubelet is unhealthy!"
      sleep 10 {{/* Wait 10 more seconds, check again, because the systemd job itself may have already restarted things */}}
      if ! output=$(${monitor_cmd} 2>&1); then
        systemctl kill kubelet
        sleep 60 {{/* Wait a minute to validate that the systemd job restarted itself after we manually killed the process */}}
        if ! systemctl is-active kubelet; then
          systemctl start kubelet
        fi
      fi
    else
      sleep "${SLEEP_TIME}"
    fi
  done
}

etcd_monitoring() {
  sleep 300 {{/* Wait for 5 minutes for etcd to be functional/stable */}}
  local max_seconds=10 output=""
  local endpoint="https://${PRIVATE_IP}:2379"
  local monitor_cmd="curl -s -S -m ${max_seconds} --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${endpoint}/v2/machines"
  while true; do
    if ! output=$(${monitor_cmd}); then
      echo $output
      echo "etcd is unhealthy!"
      sleep 10 {{/* Wait 10 more seconds, check again, because the systemd job itself may have already restarted things */}}
      if ! output=$(${monitor_cmd}); then
        systemctl kill etcd
        sleep 60 {{/* Wait a minute to validate that the systemd job restarted itself after we manually killed the process */}}
        if ! systemctl is-active etcd; then
          systemctl start etcd
        fi
      fi
    else
      sleep "${SLEEP_TIME}"
    fi
  done
}

if [[ $# -ne 1 ]]; then
  echo "Usage: health-monitor.sh <container-runtime/kubelet>"
  exit 1
fi

SLEEP_TIME=10
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
