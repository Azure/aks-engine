#!/bin/bash

set -x
# todo add support for containerd validation
if sudo kill $(pgrep dockerd); then
  sleep 10
  if ! sudo systemctl is-active docker; then
    exit 1
  fi
fi
sudo kill $(pgrep kubelet)
sleep 10
if ! sudo systemctl is-active kubelet; then
  exit 1
fi
