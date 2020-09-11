#!/bin/bash

set -x
MAX_ATTEMPTS=10
SLEEP_BETWEEN_PHASES=10

wait_systemd() {
  local attempt=1
  until sudo systemctl is-active $1; do
  if ((attempt == MAX_ATTEMPTS)); then
    exit 1
  fi
  sleep "$((2 ** ATTEMPT++))"
done
}

service_started_only_once() {
  if ! [[ $(sudo journalctl -b -u ${1} | grep -c ": Started ") == 1 ]]; then
    exit 1
  fi
}

# Validate that the monitor services themselves are stable
service_started_only_once "kubelet-monitor"
if sudo kill $(pgrep etcd); then
  service_started_only_once "etcd-monitor"
fi
service_started_only_once "docker-monitor"

# manually kill kubelet
sudo kill $(pgrep kubelet)
sleep $SLEEP_BETWEEN_PHASES
wait_systemd "kubelet" # validate that systemd brought the kubelet service back online

if apt list --installed | grep moby-engine; then
  CONTAINER_RUNTIME=docker
else
  CONTAINER_RUNTIME=containerd
fi
wait_systemd $CONTAINER_RUNTIME

# manually kill the cri
if [ $CONTAINER_RUNTIME == "docker" ]; then
  sudo kill -9 $(pgrep docker)
else
  sudo kill -9 $(pgrep containerd)
fi
sleep $SLEEP_BETWEEN_PHASES
wait_systemd $CONTAINER_RUNTIME # validate that systemd brought the cri service back online

# if I have an etcd binary that means I'm running etcd (control plane VM)
if sudo kill $(pgrep etcd); then
  sleep $SLEEP_BETWEEN_PHASES
  wait_systemd "etcd" # validate that systemd brought the etcd service back online
fi
sleep $SLEEP_BETWEEN_PHASES

wait_systemd "kubelet" # after all this nonsense let's verify that the kubelet systemd service is online
