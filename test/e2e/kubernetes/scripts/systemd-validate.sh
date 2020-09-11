#!/bin/bash

set -x
if ! [[ $(sudo journalctl -b -u kubelet-monitor | grep -c ": Started ") == 1 ]]; then
  exit 1
fi
if sudo kill $(pgrep etcd); then
  if ! [[ $(sudo journalctl -b -u etcd-monitor | grep -c ": Started ") == 1 ]]; then
    exit 1
  fi
fi
if ! [[ $(sudo journalctl -b -u docker-monitor | grep -c ": Started ") == 1 ]]; then
  exit 1
fi
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
if apt list --installed | grep moby-engine; then
  CONTAINER_RUNTIME=docker
else
  CONTAINER_RUNTIME=containerd
fi
ATTEMPT=1
until sudo systemctl is-active $CONTAINER_RUNTIME; do
  if ((ATTEMPT == MAX_ATTEMPTS)); then
    exit 1
  fi
  sleep "$((2 ** ATTEMPT++))"
done
if [ $CONTAINER_RUNTIME == "docker" ]; then
  sudo kill -9 $(pgrep docker)
else
  sudo kill -9 $(pgrep containerd)
fi
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
