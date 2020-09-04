#!/bin/bash

set -x
MAX_ATTEMPTS=10
sudo kill $(pgrep kubelet)
sleep 10
ATTEMPT=1
until sudo systemctl is-active kubelet; do
  if ((ATTEMPT == MAX_ATTEMPTS)); then
    exit 1
  fi
  sleep "$((2 ** ATTEMPT++))"
done
CONTAINER_RUNTIME=containerd
if sudo systemctl is-active docker; then
  sudo kill $(pgrep docker)
  CONTAINER_RUNTIME=docker
fi
sudo kill $(pgrep containerd)
sleep 10
ATTEMPT=1
until sudo systemctl is-active $CONTAINER_RUNTIME; do
  if ((ATTEMPT == MAX_ATTEMPTS)); then
    exit 1
  fi
  sleep "$((2 ** ATTEMPT++))"
done
if sudo kill $(pgrep etcd); then
  sleep 10
  ATTEMPT=1
  until sudo systemctl is-active etcd; do
    if ((ATTEMPT == MAX_ATTEMPTS)); then
      exit 1
    fi
    sleep "$((2 ** ATTEMPT++))"
  done
fi
ATTEMPT=1
until sudo systemctl is-active kubelet; do
  if ((ATTEMPT == MAX_ATTEMPTS)); then
    exit 1
  fi
  sleep "$((2 ** ATTEMPT++))"
done
